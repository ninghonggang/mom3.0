package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"time"
)

// WMSPickService 拣货作业服务
type WMSPickService struct {
	jobRepo    *repository.WMSPickJobRepository
	recordRepo *repository.WMSPickRecordRepository
}

func NewWMSPickService(jobRepo *repository.WMSPickJobRepository, recordRepo *repository.WMSPickRecordRepository) *WMSPickService {
	return &WMSPickService{jobRepo: jobRepo, recordRepo: recordRepo}
}

func (s *WMSPickService) List(ctx context.Context, tenantID int64, query string) ([]model.WMSPickJob, int64, error) {
	return s.jobRepo.List(ctx, tenantID, query)
}

func (s *WMSPickService) GetByID(ctx context.Context, id uint) (*model.WMSPickJob, error) {
	return s.jobRepo.GetByID(ctx, id)
}

func (s *WMSPickService) GetWithRecords(ctx context.Context, id uint) (*model.WMSPickJobRespVO, error) {
	job, err := s.jobRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	records, err := s.recordRepo.ListByPickJobID(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	var recordVOs []model.WMSPickRecordRespVO
	for _, r := range records {
		recordVOs = append(recordVOs, model.WMSPickRecordRespVO{
			Id:           r.ID,
			ItemID:       r.ItemID,
			ItemCode:     r.ItemCode,
			ItemName:     r.ItemName,
			LocationID:   r.LocationID,
			LocationCode: r.LocationCode,
			PickQty:      r.PickQty,
			PickedQty:    r.PickedQty,
			Status:       r.Status,
		})
	}

	return &model.WMSPickJobRespVO{
		Id:           job.ID,
		PickNo:       job.PickNo,
		PickType:     job.PickType,
		SourceType:   job.SourceType,
		SourceNo:     job.SourceNo,
		WarehouseID:  job.WarehouseID,
		WarehouseName: job.WarehouseName,
		Status:       job.Status,
		PickerID:     job.PickerID,
		PickerName:   job.PickerName,
		AssignTime:   job.AssignTime,
		PickedTime:   job.PickedTime,
		Remark:       job.Remark,
		Records:      recordVOs,
	}, nil
}

func (s *WMSPickService) Create(ctx context.Context, tenantID int64, req *model.WMSPickJobCreateReqVO) (*model.WMSPickJob, error) {
	// 生成拣货单号
	pickNo := fmt.Sprintf("PICK%d%d", time.Now().UnixNano(), tenantID%1000)

	job := &model.WMSPickJob{
		TenantID:    tenantID,
		PickNo:      pickNo,
		PickType:    "PICK",
		SourceType:  req.SourceType,
		SourceNo:    req.SourceNo,
		WarehouseID: req.WarehouseID,
		Status:      "PENDING",
	}

	if err := s.jobRepo.Create(ctx, job); err != nil {
		return nil, err
	}

	return job, nil
}

func (s *WMSPickService) Assign(ctx context.Context, id uint, pickerID int64, pickerName string) error {
	now := time.Now()
	return s.jobRepo.Update(ctx, id, map[string]interface{}{
		"status":      "ASSIGNED",
		"picker_id":   pickerID,
		"picker_name": pickerName,
		"assign_time": now,
	})
}

func (s *WMSPickService) Start(ctx context.Context, id uint) error {
	return s.jobRepo.Update(ctx, id, map[string]interface{}{
		"status": "PICKING",
	})
}

func (s *WMSPickService) Complete(ctx context.Context, id uint) error {
	now := time.Now()
	// 更新作业状态
	if err := s.jobRepo.Update(ctx, id, map[string]interface{}{
		"status":      "COMPLETED",
		"picked_time": now,
	}); err != nil {
		return err
	}
	// 更新明细状态
	return s.recordRepo.UpdateByPickJobID(ctx, int64(id), map[string]interface{}{
		"status": "COMPLETED",
	})
}

func (s *WMSPickService) Cancel(ctx context.Context, id uint, reason string) error {
	return s.jobRepo.Update(ctx, id, map[string]interface{}{
		"status": "CANCELLED",
		"remark": reason,
	})
}

func (s *WMSPickService) AddRecords(ctx context.Context, jobID int64, pickNo string, records []model.WMSPickRecord) error {
	for i := range records {
		records[i].PickJobID = jobID
		records[i].PickNo = pickNo
	}
	return s.recordRepo.CreateBatch(ctx, records)
}

func (s *WMSPickService) GetRecords(ctx context.Context, pickJobID int64) ([]model.WMSPickRecord, error) {
	return s.recordRepo.ListByPickJobID(ctx, pickJobID)
}
