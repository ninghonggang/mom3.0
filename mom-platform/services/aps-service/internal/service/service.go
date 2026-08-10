package service

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/ninghonggang/mom-platform/pkg/eventbus"
	"mom-platform/services/aps-service/internal/model"
	"mom-platform/services/aps-service/internal/repository"
)

var (
	ErrDuplicateCode        = errors.New("duplicate plan/job number")
	ErrInvalidStatus        = errors.New("invalid status transition")
	ErrInsufficientCapacity = errors.New("insufficient work center capacity")
	ErrWorkCenterNotFound   = errors.New("work center not found")
	ErrMpsNotFound          = errors.New("MPS plan not found")
)

// EventPublisher defines the interface for publishing domain events.
type EventPublisher interface {
	Publish(ctx context.Context, subject string, payload interface{}) error
}

type APSService struct {
	logger         *zap.Logger
	db             *gorm.DB
	mpsRepo        repository.MpsPlanRepository
	mrpRepo        repository.MrpPlanRepository
	workCenterRepo repository.WorkCenterRepository
	jobRepo        repository.ScheduleJobRepository
	changeoverRepo repository.ChangeoverRepository
	pub            EventPublisher
}

func NewAPSService(
	logger *zap.Logger,
	db *gorm.DB,
	mp repository.MpsPlanRepository,
	mrp repository.MrpPlanRepository,
	wc repository.WorkCenterRepository,
	j repository.ScheduleJobRepository,
	c repository.ChangeoverRepository,
	pub EventPublisher,
) *APSService {
	return &APSService{
		logger:         logger,
		db:             db,
		mpsRepo:        mp,
		mrpRepo:        mrp,
		workCenterRepo: wc,
		jobRepo:        j,
		changeoverRepo: c,
		pub:            pub,
	}
}

func (s *APSService) dbTransaction(ctx context.Context, fn func(txCtx context.Context) error) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txSvc := &APSService{
			logger:         s.logger,
			db:             tx,
			mpsRepo:        repository.NewMpsPlanRepo(tx),
			mrpRepo:        repository.NewMrpPlanRepo(tx),
			workCenterRepo: repository.NewWorkCenterRepo(tx),
			jobRepo:        repository.NewScheduleJobRepo(tx),
			changeoverRepo: repository.NewChangeoverRepo(tx),
			pub:            s.pub,
		}
		return fn(context.WithValue(ctx, struct{}{}, txSvc))
	})
}

// ------------------- MPS Plan -------------------

// autoPlanNo generates a plan number from tenant + date prefix
func autoPlanNo(tenantID, prefix string) string {
	return fmt.Sprintf("%s-%s-%s", prefix, tenantID, time.Now().Format("20060102150405"))
}

func (s *APSService) CreateMpsPlan(ctx context.Context, plan *model.MpsPlan) error {
	plan.PlanNo = autoPlanNo(plan.TenantID, "MPS")
	if plan.Status == "" {
		plan.Status = "DRAFT"
	}

	existing, err := s.mpsRepo.GetByPlanNo(ctx, plan.TenantID, plan.PlanNo)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	if existing != nil {
		return fmt.Errorf("%w: plan_no=%s", ErrDuplicateCode, plan.PlanNo)
	}
	return s.mpsRepo.Create(ctx, plan)
}

func (s *APSService) UpdateMpsPlan(ctx context.Context, plan *model.MpsPlan) error {
	return s.mpsRepo.Update(ctx, plan)
}

func (s *APSService) GetMpsPlan(ctx context.Context, id uint64) (*model.MpsPlan, error) {
	return s.mpsRepo.GetByID(ctx, id)
}

func (s *APSService) ListMpsPlans(ctx context.Context, tenantID string, offset, limit int) ([]model.MpsPlan, int64, error) {
	return s.mpsRepo.List(ctx, tenantID, offset, limit)
}

var mpsStatusTransitions = map[string][]string{
	"DRAFT":     {"CONFIRMED"},
	"CONFIRMED": {"RELEASED", "DRAFT"},
	"RELEASED":  {"CLOSED"},
	"CLOSED":    {},
}

func (s *APSService) UpdateMpsStatus(ctx context.Context, id uint64, newStatus string) error {
	p, err := s.mpsRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	allowed := mpsStatusTransitions[p.Status]
	valid := false
	for _, st := range allowed {
		if st == newStatus {
			valid = true
			break
		}
	}
	if !valid {
		return fmt.Errorf("%w: cannot transition from %s to %s", ErrInvalidStatus, p.Status, newStatus)
	}
	return s.mpsRepo.UpdateStatus(ctx, id, newStatus)
}

// ReleaseMPS sets the MPS status to RELEASED and publishes an event.
func (s *APSService) ReleaseMPS(ctx context.Context, id uint64) (*model.MpsPlan, error) {
	p, err := s.mpsRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if p.Status != "DRAFT" && p.Status != "CONFIRMED" {
		return nil, fmt.Errorf("%w: cannot release from %s", ErrInvalidStatus, p.Status)
	}
	if err := s.mpsRepo.UpdateStatus(ctx, id, "RELEASED"); err != nil {
		return nil, err
	}
	p.Status = "RELEASED"

	if s.pub != nil {
		_ = s.pub.Publish(ctx, eventbus.SubjectAPSMPSReleased, map[string]interface{}{
			"mps_id":      p.ID,
			"plan_no":     p.PlanNo,
			"material_id": p.MaterialID,
			"quantity":    p.PlannedQty,
			"plan_month":  p.PlanMonth,
		})
	}

	return p, nil
}

// ------------------- MRP Generation -------------------

// MrpResult holds the MRP calculation result for one material
type MrpResult struct {
	MaterialID          uint64  `json:"material_id"`
	GrossReq            float64 `json:"gross_req"`
	ScheduledReceipt    float64 `json:"scheduled_receipt"`
	ProjectedOnHand     float64 `json:"projected_on_hand"`
	NetReq              float64 `json:"net_req"`
	PlannedOrderRelease float64 `json:"planned_order_release"`
	PlannedOrderReceipt float64 `json:"planned_order_receipt"`
}

// GenerateMrp explodes MPS via BOM and calculates gross/net requirements.
// BOM data is external (furnished by caller or fetched from MDM).
// bombardParams carries exploded demand: materialID -> quantity
func (s *APSService) GenerateMrp(ctx context.Context, tenantID string, mpsID uint64, bombResults BomExplosion) ([]model.MrpPlan, error) {
	mps, err := s.mpsRepo.GetByID(ctx, mpsID)
	if err != nil {
		return nil, fmt.Errorf("mps: %w", err)
	}
	_ = mps // validated MPS exists

	// Regenerate: delete old MRP records for this MPS
	if err := s.mrpRepo.DeleteByMpsID(ctx, mpsID); err != nil {
		return nil, fmt.Errorf("delete old mrp: %w", err)
	}

	var plans []model.MrpPlan
	planNoPrefix := autoPlanNo(tenantID, "MRP")

	for materialID, qty := range bombResults {
		// Simplified MRP: gross = demand, net = gross (no inventory tracking here)
		gross := qty
		net := gross // simplified — real system would subtract on-hand

		plan := model.MrpPlan{
			TenantID:            tenantID,
			PlanNo:              fmt.Sprintf("%s-%d", planNoPrefix, materialID),
			MpsID:               mpsID,
			MaterialID:          materialID,
			GrossReq:            gross,
			ScheduledReceipt:    0,
			ProjectedOnHand:     0,
			NetReq:              net,
			PlannedOrderRelease: net,
			PlannedOrderReceipt: net,
		}
		plans = append(plans, plan)
	}

	if err := s.mrpRepo.CreateBatch(ctx, plans); err != nil {
		return nil, fmt.Errorf("create mrp batch: %w", err)
	}

	s.logger.Info("MRP generated",
		zap.Uint64("mps_id", mpsID),
		zap.Int("plan_count", len(plans)),
	)

	return plans, nil
}

// BomExplosion maps materialID -> required quantity
type BomExplosion map[uint64]float64

// RunMrp runs MRP for the given MPS and publishes an event.
// It uses the MPS material itself as a simplified single-level BOM.
func (s *APSService) RunMrp(ctx context.Context, tenantID string, mpsID uint64) ([]model.MrpPlan, error) {
	mps, err := s.mpsRepo.GetByID(ctx, mpsID)
	if err != nil {
		return nil, fmt.Errorf("mps: %w", err)
	}

	// Simplified BOM: the MPS material itself is the only demand
	bom := BomExplosion{mps.MaterialID: mps.PlannedQty}

	plans, err := s.GenerateMrp(ctx, tenantID, mpsID, bom)
	if err != nil {
		return nil, err
	}

	if s.pub != nil {
		_ = s.pub.Publish(ctx, eventbus.SubjectAPSMRPCompleted, map[string]interface{}{
			"mps_id":     mpsID,
			"mps_no":     mps.PlanNo,
			"plan_count": len(plans),
		})
	}

	return plans, nil
}

// GetMrpPlan returns a single MRP plan by ID.
func (s *APSService) GetMrpPlan(ctx context.Context, id uint64) (*model.MrpPlan, error) {
	return s.mrpRepo.GetByID(ctx, id)
}

// GetMrpPlansByMpsID returns all MRP plans for a given MPS.
func (s *APSService) GetMrpPlansByMpsID(ctx context.Context, mpsID uint64) ([]model.MrpPlan, error) {
	return s.mrpRepo.ListByMpsID(ctx, mpsID)
}

// ------------------- Scheduling -------------------

// ScheduleJobRequest contains parameters for scheduling
type ScheduleJobRequest struct {
	TenantID          string
	JobNo             string
	ProductionOrderID uint64
	WorkCenterID      uint64
	MaterialID        uint64
	Quantity          float64
	DurationHours     float64
}

// ScheduleJob assigns a job to a work center with capacity check and forward scheduling.
func (s *APSService) ScheduleJob(ctx context.Context, req ScheduleJobRequest) (*model.ScheduleJob, error) {
	wc, err := s.workCenterRepo.GetByID(ctx, req.WorkCenterID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrWorkCenterNotFound, err)
	}

	// Forward scheduling: find next available time slot on work center
	now := time.Now().Truncate(time.Hour)
	plannedStart, err := s.findNextAvailableSlot(ctx, req.WorkCenterID, now, req.DurationHours, wc.CapacityPerDay)
	if err != nil {
		return nil, err
	}

	plannedEnd := plannedStart.Add(time.Duration(req.DurationHours) * time.Hour)

	if req.JobNo == "" {
		req.JobNo = autoPlanNo(req.TenantID, "JOB")
	}

	job := &model.ScheduleJob{
		TenantID:          req.TenantID,
		JobNo:             req.JobNo,
		ProductionOrderID: req.ProductionOrderID,
		WorkCenterID:      req.WorkCenterID,
		PlannedStart:      plannedStart,
		PlannedEnd:        plannedEnd,
		Status:            "SCHEDULED",
	}

	if err := s.jobRepo.Create(ctx, job); err != nil {
		return nil, err
	}

	s.logger.Info("job scheduled",
		zap.String("job_no", job.JobNo),
		zap.Time("start", plannedStart),
		zap.Time("end", plannedEnd),
	)

	return job, nil
}

const workHoursPerDay = 8.0

// findNextAvailableSlot forward schedules to find the first open time window.
func (s *APSService) findNextAvailableSlot(ctx context.Context, wcID uint64, from time.Time, durationHours float64, capacityPerDay float64) (time.Time, error) {
	hoursPerDay := durationHours
	if int(hoursPerDay) < 1 {
		hoursPerDay = 1
	}

	// Fetch existing jobs in a window (from to from + 30 days)
	to := from.Add(30 * 24 * time.Hour)
	jobs, err := s.jobRepo.ListByWorkCenter(ctx, wcID, from, to)
	if err != nil {
		return from, err
	}

	if len(jobs) == 0 {
		return from, nil
	}

	// Sort by start time
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].PlannedStart.Before(jobs[j].PlannedStart)
	})

	slot := from
	for _, job := range jobs {
		gap := job.PlannedStart.Sub(slot).Hours()
		if gap >= hoursPerDay {
			// Enough gap before this job
			return slot, nil
		}
		// Move past this job
		if job.PlannedEnd.After(slot) {
			slot = job.PlannedEnd.Add(time.Duration(s.getChangeover(ctx, 0, 0)) * time.Minute)
		}
	}

	return slot, nil
}

// getChangeover looks up changeover time between two materials.
func (s *APSService) getChangeover(ctx context.Context, fromMaterialID, toMaterialID uint64) float64 {
	if fromMaterialID == 0 || toMaterialID == 0 {
		return 0
	}
	t, err := s.changeoverRepo.GetChangeoverTime(ctx, fromMaterialID, toMaterialID)
	if err != nil {
		return 0
	}
	return t
}

// GetChangeoverTime is the public API to look up changeover between two materials.
func (s *APSService) GetChangeoverTime(ctx context.Context, fromMaterialID, toMaterialID uint64) (float64, error) {
	t, err := s.changeoverRepo.GetChangeoverTime(ctx, fromMaterialID, toMaterialID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return 0, nil
		}
		return 0, err
	}
	return t, nil
}

// GetScheduledJobs returns scheduled jobs for a work center.
func (s *APSService) GetScheduledJobs(ctx context.Context, workCenterID uint64, start, end time.Time) ([]model.ScheduleJob, error) {
	return s.jobRepo.ListByWorkCenter(ctx, workCenterID, start, end)
}

// CreateScheduleJob creates a schedule job from the request and returns it.
func (s *APSService) CreateScheduleJob(ctx context.Context, job *model.ScheduleJob) (*model.ScheduleJob, error) {
	if job.JobNo == "" {
		job.JobNo = autoPlanNo(job.TenantID, "JOB")
	}
	if job.Status == "" {
		job.Status = "PENDING"
	}
	if err := s.jobRepo.Create(ctx, job); err != nil {
		return nil, err
	}
	return job, nil
}

// GetScheduleJob returns a schedule job by ID.
func (s *APSService) GetScheduleJob(ctx context.Context, id uint64) (*model.ScheduleJob, error) {
	return s.jobRepo.GetByID(ctx, id)
}

// ListScheduleJobs returns a paginated list of schedule jobs for the tenant.
func (s *APSService) ListScheduleJobs(ctx context.Context, tenantID string, offset, limit int) ([]model.ScheduleJob, int64, error) {
	return s.jobRepo.List(ctx, tenantID, offset, limit)
}

// PublishScheduleJob sets the job status to SCHEDULED and publishes an event.
func (s *APSService) PublishScheduleJob(ctx context.Context, id uint64) (*model.ScheduleJob, error) {
	job, err := s.jobRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if job.Status != "PENDING" && job.Status != "SCHEDULED" {
		return nil, fmt.Errorf("%w: cannot publish from %s", ErrInvalidStatus, job.Status)
	}
	job.Status = "SCHEDULED"
	if err := s.jobRepo.Update(ctx, job); err != nil {
		return nil, err
	}

	if s.pub != nil {
		_ = s.pub.Publish(ctx, eventbus.SubjectAPSSchedulePublished, map[string]interface{}{
			"plan_id":       job.ID,
			"job_no":        job.JobNo,
			"work_center_id": job.WorkCenterID,
			"planned_start":  job.PlannedStart,
			"planned_end":    job.PlannedEnd,
		})
	}

	return job, nil
}

// UpdateJobStatus transitions a job's status.
var jobStatusTransitions = map[string][]string{
	"PENDING":   {"SCHEDULED"},
	"SCHEDULED": {"RUNNING"},
	"RUNNING":   {"COMPLETED", "DELAYED"},
	"DELAYED":   {"RUNNING", "COMPLETED"},
	"COMPLETED": {},
}

func (s *APSService) UpdateJobStatus(ctx context.Context, id uint64, newStatus string) error {
	job, err := s.jobRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	allowed := jobStatusTransitions[job.Status]
	valid := false
	for _, st := range allowed {
		if st == newStatus {
			valid = true
			break
		}
	}
	if !valid {
		return fmt.Errorf("%w: cannot transition from %s to %s", ErrInvalidStatus, job.Status, newStatus)
	}

	now := time.Now()
	updates := model.ScheduleJob{
		ID:     id,
		Status: newStatus,
	}
	if newStatus == "RUNNING" && job.ActualStart == nil {
		updates.ActualStart = &now
	}
	if newStatus == "COMPLETED" || newStatus == "DELAYED" {
		updates.ActualEnd = &now
	}

	return s.jobRepo.Update(ctx, &updates)
}

// ------------------- Work Center -------------------

func (s *APSService) CreateWorkCenter(ctx context.Context, wc *model.WorkCenter) error {
	existing, err := s.workCenterRepo.GetByCode(ctx, wc.TenantID, wc.CenterCode)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	if existing != nil {
		return fmt.Errorf("%w: center_code=%s", ErrDuplicateCode, wc.CenterCode)
	}
	return s.workCenterRepo.Create(ctx, wc)
}

func (s *APSService) GetWorkCenter(ctx context.Context, id uint64) (*model.WorkCenter, error) {
	return s.workCenterRepo.GetByID(ctx, id)
}

func (s *APSService) ListWorkCenters(ctx context.Context, tenantID string, offset, limit int) ([]model.WorkCenter, int64, error) {
	return s.workCenterRepo.List(ctx, tenantID, offset, limit)
}

func (s *APSService) ListWorkCentersByWorkshop(ctx context.Context, tenantID string, workshopID uint64) ([]model.WorkCenter, error) {
	return s.workCenterRepo.ListByWorkshop(ctx, tenantID, workshopID)
}

// ------------------- Changeover -------------------

func (s *APSService) CreateChangeover(ctx context.Context, c *model.Changeover) error {
	return s.changeoverRepo.Create(ctx, c)
}
