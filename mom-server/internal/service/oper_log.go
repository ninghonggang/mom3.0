package service

import (
	"context"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type OperLogService struct {
	operLogRepo *repository.OperLogRepository
}

func NewOperLogService(operLogRepo *repository.OperLogRepository) *OperLogService {
	return &OperLogService{operLogRepo: operLogRepo}
}

func (s *OperLogService) RecordOper(ctx context.Context, log *model.OperLog) error {
	return s.operLogRepo.Create(ctx, log)
}

func (s *OperLogService) GetList(ctx context.Context, tenantID int64, title, operName, businessType, status string, page, pageSize int) ([]model.OperLog, int64, error) {
	return s.operLogRepo.FindByPage(ctx, tenantID, title, operName, businessType, status, page, pageSize)
}
