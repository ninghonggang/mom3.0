package service

import (
	"context"
	"encoding/json"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type ProductionOrderChangeLogService struct {
	repo *repository.ProductionOrderChangeLogRepository
}

func NewProductionOrderChangeLogService(repo *repository.ProductionOrderChangeLogRepository) *ProductionOrderChangeLogService {
	return &ProductionOrderChangeLogService{repo: repo}
}

// ChangeType constants
const (
	ChangeTypeQuantityChange  = "QUANTITY_CHANGE"
	ChangeTypeDateChange      = "DATE_CHANGE"
	ChangeTypePriorityChange  = "PRIORITY_CHANGE"
	ChangeTypeLineChange      = "LINE_CHANGE"
	ChangeTypeStatusChange    = "STATUS_CHANGE"
)

// RecordChange 记录变更
func (s *ProductionOrderChangeLogService) RecordChange(ctx context.Context, orderID int64, orderNo string, changeType string, oldValue interface{}, newValue interface{}, changeReason string, changedBy string) error {
	oldJSON, _ := json.Marshal(oldValue)
	newJSON, _ := json.Marshal(newValue)

	log := &model.ProductionOrderChangeLog{
		OrderID:      orderID,
		OrderNo:      orderNo,
		ChangeType:   changeType,
		OldValue:     string(oldJSON),
		NewValue:     string(newJSON),
		ChangeReason: changeReason,
		ChangedBy:    changedBy,
	}

	return s.repo.Create(ctx, log)
}

// GetOrderChanges 获取订单变更历史
func (s *ProductionOrderChangeLogService) GetOrderChanges(ctx context.Context, orderID int64) ([]model.ProductionOrderChangeLog, error) {
	return s.repo.ListByOrderID(ctx, orderID)
}
