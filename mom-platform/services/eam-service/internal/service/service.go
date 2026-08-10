package service

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"mom-platform/services/eam-service/internal/model"
	"mom-platform/services/eam-service/internal/repository"
)

// Service provides EAM business logic.
type Service struct {
	repo  repository.Repository
	db    *gorm.DB
	logger *zap.Logger
}

// New creates a new Service.
func New(repo repository.Repository, db *gorm.DB, logger *zap.Logger) *Service {
	return &Service{repo: repo, db: db, logger: logger}
}

// =====================================================================
// Equipment
// =====================================================================

// CreateEquipment creates a new equipment record.
func (s *Service) CreateEquipment(ctx context.Context, e *model.Equipment) error {
	if e.Status == "" {
		e.Status = model.EquipmentStatusIdle
	}
	if err := s.repo.Equipment().Create(ctx, e); err != nil {
		return fmt.Errorf("create equipment: %w", err)
	}
	return nil
}

// GetEquipment retrieves equipment by ID.
func (s *Service) GetEquipment(ctx context.Context, id int64) (*model.Equipment, error) {
	e, err := s.repo.Equipment().GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get equipment %d: %w", id, err)
	}
	return e, nil
}

// ListEquipment lists equipment with filtering and pagination.
func (s *Service) ListEquipment(ctx context.Context, filter repository.EquipmentFilter, page model.PageQuery) ([]model.Equipment, int64, error) {
	return s.repo.Equipment().List(ctx, filter, page)
}

// UpdateEquipmentStatus updates only the status field of an equipment.
func (s *Service) UpdateEquipmentStatus(ctx context.Context, id int64, status model.EquipmentStatus) error {
	e, err := s.repo.Equipment().GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("get equipment for status update: %w", err)
	}
	e.Status = status
	if err := s.repo.Equipment().Update(ctx, e); err != nil {
		return fmt.Errorf("update equipment status: %w", err)
	}
	return nil
}

// =====================================================================
// RepairOrder
// =====================================================================

// repairStateTransitions defines valid status transitions for repair orders.
var repairStateTransitions = map[model.RepairOrderStatus][]model.RepairOrderStatus{
	model.RepairStatusReported:    {model.RepairStatusAssigned, model.RepairStatusCancelled},
	model.RepairStatusAssigned:    {model.RepairStatusInProgress, model.RepairStatusCancelled},
	model.RepairStatusInProgress:  {model.RepairStatusPendingParts, model.RepairStatusCompleted, model.RepairStatusCancelled},
	model.RepairStatusPendingParts: {model.RepairStatusInProgress, model.RepairStatusCancelled},
	model.RepairStatusCompleted:   {model.RepairStatusVerified},
	model.RepairStatusVerified:    {},
	model.RepairStatusCancelled:   {},
}

// isValidTransition checks whether transitioning from one status to another is allowed.
func isValidTransition(from, to model.RepairOrderStatus) bool {
	allowed, ok := repairStateTransitions[from]
	if !ok {
		return false
	}
	for _, s := range allowed {
		if s == to {
			return true
		}
	}
	return false
}

// CreateRepairOrder creates a new repair order with auto-generated repair_no.
func (s *Service) CreateRepairOrder(ctx context.Context, ro *model.RepairOrder) (*model.RepairOrder, error) {
	// Validate equipment exists
	equip, err := s.repo.Equipment().GetByID(ctx, ro.EquipmentID)
	if err != nil {
		return nil, fmt.Errorf("validate equipment %d: %w", ro.EquipmentID, err)
	}
	if equip.Status == model.EquipmentStatusScrapped {
		return nil, fmt.Errorf("equipment %d is scrapped, cannot create repair order", ro.EquipmentID)
	}

	// Auto-generate repair_no: RP-yyyyMMddHHmmss
	now := time.Now()
	ro.RepairNo = fmt.Sprintf("RP-%s", now.Format("20060102150405"))
	ro.Status = model.RepairStatusReported
	ro.ReportedAt = now

	if err := s.repo.RepairOrder().Create(ctx, ro); err != nil {
		return nil, fmt.Errorf("create repair order: %w", err)
	}

	s.logger.Info("repair order created",
		zap.Int64("id", ro.ID),
		zap.String("repair_no", ro.RepairNo),
		zap.Int64("equipment_id", ro.EquipmentID),
	)
	return ro, nil
}

// UpdateRepairOrder updates a repair order with state machine validation.
func (s *Service) UpdateRepairOrder(ctx context.Context, id int64, newStatus model.RepairOrderStatus, repairmanID *int64) (*model.RepairOrder, error) {
	ro, err := s.repo.RepairOrder().GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get repair order %d: %w", id, err)
	}

	if !isValidTransition(ro.Status, newStatus) {
		return nil, fmt.Errorf("invalid status transition: %s -> %s", ro.Status, newStatus)
	}

	ro.Status = newStatus
	if repairmanID != nil {
		ro.RepairmanID = repairmanID
	}

	if newStatus == model.RepairStatusCompleted {
		now := time.Now()
		ro.CompletedAt = &now
	}

	if err := s.repo.RepairOrder().Update(ctx, ro); err != nil {
		return nil, fmt.Errorf("update repair order: %w", err)
	}

	s.logger.Info("repair order updated",
		zap.Int64("id", ro.ID),
		zap.String("new_status", string(newStatus)),
	)
	return ro, nil
}

// GetRepairOrder retrieves a repair order by ID.
func (s *Service) GetRepairOrder(ctx context.Context, id int64) (*model.RepairOrder, error) {
	ro, err := s.repo.RepairOrder().GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get repair order %d: %w", id, err)
	}
	return ro, nil
}

// ListRepairOrders lists repair orders with filtering and pagination.
func (s *Service) ListRepairOrders(ctx context.Context, filter repository.RepairOrderFilter, page model.PageQuery) ([]model.RepairOrder, int64, error) {
	return s.repo.RepairOrder().List(ctx, filter, page)
}

// =====================================================================
// EquipmentDowntime
// =====================================================================

// StartDowntime creates a downtime record and updates equipment status to REPAIR.
func (s *Service) StartDowntime(ctx context.Context, equipmentID int64, dtType model.DowntimeType, reason string) (*model.EquipmentDowntime, error) {
	// Validate equipment exists
	equip, err := s.repo.Equipment().GetByID(ctx, equipmentID)
	if err != nil {
		return nil, fmt.Errorf("validate equipment %d: %w", equipmentID, err)
	}
	if equip.Status == model.EquipmentStatusScrapped {
		return nil, fmt.Errorf("equipment %d is scrapped", equipmentID)
	}

	now := time.Now()
	downtime := &model.EquipmentDowntime{
		EquipmentID:  equipmentID,
		DowntimeType: dtType,
		StartTime:    now,
		Reason:       reason,
		Status:       model.DowntimeStatusActive,
	}

	// Use transaction to ensure consistency
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(downtime).Error; err != nil {
			return fmt.Errorf("create downtime: %w", err)
		}
		// Update equipment status to REPAIR for unplanned downtime
		if dtType == model.DowntimeTypeUnplanned {
			if err := tx.Model(&model.Equipment{}).Where("id = ?", equipmentID).
				Update("status", model.EquipmentStatusRepair).Error; err != nil {
				return fmt.Errorf("update equipment status: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	s.logger.Info("downtime started",
		zap.Int64("id", downtime.ID),
		zap.Int64("equipment_id", equipmentID),
		zap.String("type", string(dtType)),
	)
	return downtime, nil
}

// ResolveDowntime resolves a downtime record, calculates duration, and restores equipment status.
// resolution/resolverID 为空时保留原值，避免二次调用把已有留痕抹掉。
func (s *Service) ResolveDowntime(ctx context.Context, id int64, resolution, resolverID string) (*model.EquipmentDowntime, error) {
	downtime, err := s.repo.Downtime().GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get downtime %d: %w", id, err)
	}
	if downtime.Status != model.DowntimeStatusActive {
		return nil, fmt.Errorf("downtime %d is not active (status=%s)", id, downtime.Status)
	}

	now := time.Now()
	downtime.EndTime = &now
	downtime.DurationSeconds = int32(now.Sub(downtime.StartTime).Seconds())
	downtime.Status = model.DowntimeStatusResolved
	if resolution != "" {
		downtime.Resolution = resolution
	}
	if resolverID != "" {
		downtime.ResolverID = resolverID
	}

	// Restore equipment status to IDLE (or RUNNING) using transaction
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(downtime).Error; err != nil {
			return fmt.Errorf("update downtime: %w", err)
		}
		// Restore equipment status for unplanned downtime
		if downtime.DowntimeType == model.DowntimeTypeUnplanned {
			if err := tx.Model(&model.Equipment{}).Where("id = ?", downtime.EquipmentID).
				Update("status", model.EquipmentStatusIdle).Error; err != nil {
				return fmt.Errorf("restore equipment status: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	s.logger.Info("downtime resolved",
		zap.Int64("id", downtime.ID),
		zap.Int32("duration_seconds", downtime.DurationSeconds),
	)
	return downtime, nil
}

// ListDowntimes lists downtime records with filtering and pagination.
func (s *Service) ListDowntimes(ctx context.Context, filter repository.DowntimeFilter, page model.PageQuery) ([]model.EquipmentDowntime, int64, error) {
	return s.repo.Downtime().List(ctx, filter, page)
}

// =====================================================================
// OEE
// =====================================================================

// SaveOee saves or updates an OEE record, auto-calculating OEE = A * P * Q.
func (s *Service) SaveOee(ctx context.Context, o *model.EquipmentOee) (*model.EquipmentOee, error) {
	o.OEE = calcOEE(o.Availability, o.Performance, o.Quality)
	if err := s.repo.Oee().Create(ctx, o); err != nil {
		return nil, fmt.Errorf("save oee: %w", err)
	}
	s.logger.Info("oee saved",
		zap.Int64("equipment_id", o.EquipmentID),
		zap.String("date", o.CalcDate),
		zap.Float64("oee", o.OEE),
	)
	return o, nil
}

// ListOee lists OEE records.
func (s *Service) ListOee(ctx context.Context, filter repository.OeeFilter) ([]model.EquipmentOee, error) {
	return s.repo.Oee().List(ctx, filter)
}

// calcOEE computes OEE = availability * performance * quality.
func calcOEE(availability, performance, quality float64) float64 {
	oee := availability * performance * quality
	// Round to 4 decimal places
	return float64(int(oee*10000)) / 10000
}

// =====================================================================
// MaintenancePlan
// =====================================================================

// CreateMaintenancePlan creates a new maintenance plan with auto-generated plan_no.
func (s *Service) CreateMaintenancePlan(ctx context.Context, p *model.MaintenancePlan) (*model.MaintenancePlan, error) {
	// Validate equipment exists
	if _, err := s.repo.Equipment().GetByID(ctx, p.EquipmentID); err != nil {
		return nil, fmt.Errorf("validate equipment %d: %w", p.EquipmentID, err)
	}

	p.PlanNo = fmt.Sprintf("MP-%s", time.Now().Format("20060102150405"))
	p.Status = model.MaintenanceStatusScheduled

	if err := s.repo.MaintenancePlan().Create(ctx, p); err != nil {
		return nil, fmt.Errorf("create maintenance plan: %w", err)
	}

	s.logger.Info("maintenance plan created",
		zap.Int64("id", p.ID),
		zap.String("plan_no", p.PlanNo),
		zap.Int64("equipment_id", p.EquipmentID),
	)
	return p, nil
}

// ListMaintenancePlans lists maintenance plans with filtering and pagination.
func (s *Service) ListMaintenancePlans(ctx context.Context, filter repository.MaintenancePlanFilter, page model.PageQuery) ([]model.MaintenancePlan, int64, error) {
	return s.repo.MaintenancePlan().List(ctx, filter, page)
}

// =====================================================================
// EquipmentCheck
// =====================================================================

// CreateCheck creates a new check record with auto-generated check_no.
func (s *Service) CreateCheck(ctx context.Context, c *model.EquipmentCheck) (*model.EquipmentCheck, error) {
	// Validate equipment exists
	if _, err := s.repo.Equipment().GetByID(ctx, c.EquipmentID); err != nil {
		return nil, fmt.Errorf("validate equipment %d: %w", c.EquipmentID, err)
	}

	c.CheckNo = fmt.Sprintf("CK-%s", time.Now().Format("20060102150405"))
	c.CheckTime = time.Now()

	if err := s.repo.Check().Create(ctx, c); err != nil {
		return nil, fmt.Errorf("create check: %w", err)
	}

	s.logger.Info("check created",
		zap.Int64("id", c.ID),
		zap.String("check_no", c.CheckNo),
		zap.Int64("equipment_id", c.EquipmentID),
	)
	return c, nil
}

// ListChecks lists check records with filtering and pagination.
func (s *Service) ListChecks(ctx context.Context, filter repository.CheckFilter, page model.PageQuery) ([]model.EquipmentCheck, int64, error) {
	return s.repo.Check().List(ctx, filter, page)
}
