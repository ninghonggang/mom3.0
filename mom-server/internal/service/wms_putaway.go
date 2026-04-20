package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"

	"gorm.io/gorm"
)

type PutawayService struct {
	db             *gorm.DB
	jobRepo        *repository.PutawayJobRepository
	recordRepo     *repository.PutawayRecordRepository
	codeRuleSvc    *CodeRuleService
}

func NewPutawayService(db *gorm.DB, jobRepo *repository.PutawayJobRepository, recordRepo *repository.PutawayRecordRepository, codeRuleSvc *CodeRuleService) *PutawayService {
	return &PutawayService{
		db:         db,
		jobRepo:    jobRepo,
		recordRepo: recordRepo,
		codeRuleSvc: codeRuleSvc,
	}
}

func (s *PutawayService) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.WMSPutawayJob, int64, error) {
	return s.jobRepo.List(ctx, tenantID, query)
}

func (s *PutawayService) GetByID(ctx context.Context, id uint) (*model.WMSPutawayJob, error) {
	return s.jobRepo.GetByID(ctx, id)
}

func (s *PutawayService) GetWithRecords(ctx context.Context, id uint) (*model.WMSPutawayJobRespVO, error) {
	job, err := s.jobRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	records, err := s.recordRepo.ListByJobID(ctx, id)
	if err != nil {
		return nil, err
	}

	recordVOs := make([]model.WMSPutawayRecordRespVO, len(records))
	for i, r := range records {
		recordVOs[i] = model.WMSPutawayRecordRespVO{
			ID:             r.ID,
			ItemID:         r.ItemID,
			ItemCode:       r.ItemCode,
			ItemName:       r.ItemName,
			FromLocationID: r.FromLocationID,
			ToLocationID:   r.ToLocationID,
			PutawayQty:     r.PutawayQty,
			PutawardedQty:  r.PutawardedQty,
			Status:         r.Status,
		}
	}

	return &model.WMSPutawayJobRespVO{
		ID:           job.ID,
		PutawayNo:    job.PutawayNo,
		SourceType:   job.SourceType,
		SourceNo:     job.SourceNo,
		Status:       job.Status,
		OperatorID:   job.OperatorID,
		OperatorName: job.OperatorName,
		Records:      recordVOs,
	}, nil
}

// Create 创建上架作业单
func (s *PutawayService) Create(ctx context.Context, tenantID int64, req *model.WMSPutawayJobCreateReqVO, username string) (*model.WMSPutawayJob, error) {
	// 生成上架单号
	putawayNo := fmt.Sprintf("PA%s%d", time.Now().Format("20060102150405"), tenantID%1000)

	job := &model.WMSPutawayJob{
		TenantID:    tenantID,
		PutawayNo:   putawayNo,
		SourceType:  req.SourceType,
		SourceNo:    req.SourceNo,
		WarehouseID: req.WarehouseID,
		Status:      "PENDING",
	}

	if err := s.jobRepo.Create(ctx, job); err != nil {
		return nil, fmt.Errorf("failed to create putaway job: %w", err)
	}

	return job, nil
}

// Assign 分配操作员
func (s *PutawayService) Assign(ctx context.Context, id uint, req *model.WMSPutawayJobAssignReqVO) error {
	job, err := s.jobRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("putaway job not found: %w", err)
	}

	if job.Status != "PENDING" {
		return fmt.Errorf("only PENDING status job can be assigned")
	}

	now := time.Now()
	updates := map[string]interface{}{
		"status":        "ASSIGNED",
		"operator_id":    req.OperatorID,
		"operator_name": req.OperatorName,
		"assign_time":    &now,
	}

	return s.jobRepo.Update(ctx, id, updates)
}

// Start 开始上架
func (s *PutawayService) Start(ctx context.Context, id uint) error {
	job, err := s.jobRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("putaway job not found: %w", err)
	}

	if job.Status != "ASSIGNED" {
		return fmt.Errorf("only ASSIGNED status job can be started")
	}

	updates := map[string]interface{}{
		"status": "PUTAWAYING",
	}

	if err := s.jobRepo.Update(ctx, id, updates); err != nil {
		return fmt.Errorf("failed to start putaway: %w", err)
	}

	// 更新明细状态
	return s.recordRepo.UpdateByJobID(ctx, id, map[string]interface{}{"status": "PUTAWAYING"})
}

// Complete 完成上架
func (s *PutawayService) Complete(ctx context.Context, id uint) error {
	job, err := s.jobRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("putaway job not found: %w", err)
	}

	if job.Status != "PUTAWAYING" {
		return fmt.Errorf("only PUTAWAYING status job can be completed")
	}

	now := time.Now()
	updates := map[string]interface{}{
		"status":       "COMPLETED",
		"putaway_time": &now,
	}

	if err := s.jobRepo.Update(ctx, id, updates); err != nil {
		return fmt.Errorf("failed to complete putaway: %w", err)
	}

	// 更新明细状态
	return s.recordRepo.UpdateByJobID(ctx, id, map[string]interface{}{"status": "COMPLETED"})
}

// Cancel 取消上架
func (s *PutawayService) Cancel(ctx context.Context, id uint, reason string) error {
	job, err := s.jobRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("putaway job not found: %w", err)
	}

	if job.Status == "COMPLETED" || job.Status == "CANCELLED" {
		return fmt.Errorf("completed or cancelled job cannot be cancelled")
	}

	updates := map[string]interface{}{
		"status": "CANCELLED",
		"remark":  reason,
	}

	return s.jobRepo.Update(ctx, id, updates)
}

// AddRecord 添加上架明细
func (s *PutawayService) AddRecord(ctx context.Context, jobID uint, record *model.WMSPutawayRecord) error {
	job, err := s.jobRepo.GetByID(ctx, jobID)
	if err != nil {
		return fmt.Errorf("putaway job not found: %w", err)
	}

	if job.Status != "PENDING" {
		return fmt.Errorf("only PENDING status job can add records")
	}

	record.PutawayJobID = int64(jobID)
	record.PutawayNo = job.PutawayNo
	record.Status = "PENDING"

	return s.recordRepo.Create(ctx, record)
}

// UpdateRecord 更新上架明细
func (s *PutawayService) UpdateRecord(ctx context.Context, recordID uint, updates map[string]interface{}) error {
	return s.recordRepo.Update(ctx, recordID, updates)
}
