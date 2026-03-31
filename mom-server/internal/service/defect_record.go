package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type DefectRecordService struct {
	repo *repository.DefectRecordRepository
}

func NewDefectRecordService(repo *repository.DefectRecordRepository) *DefectRecordService {
	return &DefectRecordService{repo: repo}
}

func (s *DefectRecordService) List(ctx context.Context) ([]model.DefectRecord, int64, error) {
	return s.repo.List(ctx, 1)
}

func (s *DefectRecordService) GetByID(ctx context.Context, id string) (*model.DefectRecord, error) {
	var recordID uint
	_, err := fmt.Sscanf(id, "%d", &recordID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, recordID)
}

func (s *DefectRecordService) Create(ctx context.Context, record *model.DefectRecord) error {
	record.TenantID = 1
	return s.repo.Create(ctx, record)
}

func (s *DefectRecordService) Update(ctx context.Context, id string, record *model.DefectRecord) error {
	var recordID uint
	_, err := fmt.Sscanf(id, "%d", &recordID)
	if err != nil {
		return err
	}
	updates := map[string]interface{}{
		"process_id":    record.ProcessID,
		"process_name":  record.ProcessName,
		"defect_code_id": record.DefectCodeID,
		"defect_code":   record.DefectCode,
		"defect_name":   record.DefectName,
		"quantity":      record.Quantity,
		"handle_method": record.HandleMethod,
		"status":        record.Status,
	}
	return s.repo.Update(ctx, recordID, updates)
}

func (s *DefectRecordService) Delete(ctx context.Context, id string) error {
	var recordID uint
	_, err := fmt.Sscanf(id, "%d", &recordID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, recordID)
}

func (s *DefectRecordService) Handle(ctx context.Context, id string, handleMethod int, handleUserID int64) error {
	var recordID uint
	_, err := fmt.Sscanf(id, "%d", &recordID)
	if err != nil {
		return err
	}
	updates := map[string]interface{}{
		"handle_method": handleMethod,
		"handle_user_id": handleUserID,
		"status":        3, // 3已处理
	}
	return s.repo.Update(ctx, recordID, updates)
}
