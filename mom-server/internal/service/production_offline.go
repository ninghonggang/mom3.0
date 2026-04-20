package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"time"
)

type ProductionOfflineService struct {
	offlineRepo *repository.ProductionOfflineRepository
}

func NewProductionOfflineService(offlineRepo *repository.ProductionOfflineRepository) *ProductionOfflineService {
	return &ProductionOfflineService{offlineRepo: offlineRepo}
}

func (s *ProductionOfflineService) List(ctx context.Context, tenantID uint64, query map[string]interface{}) ([]model.ProductionOffline, int64, error) {
	return s.offlineRepo.List(nil, tenantID, query)
}

func (s *ProductionOfflineService) GetByID(ctx context.Context, id uint64) (*model.ProductionOffline, error) {
	return s.offlineRepo.GetByID(id)
}

func (s *ProductionOfflineService) Create(ctx context.Context, tenantID uint64, req *model.ProductionOfflineCreateRequest, username string) (*model.ProductionOffline, error) {
	// 生成离线单号
	offlineCode, err := s.GenerateOfflineCode(tenantID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate offline code: %w", err)
	}

	offline := &model.ProductionOffline{
		TenantID:       tenantID,
		OfflineCode:    offlineCode,
		WorkOrderID:    req.WorkOrderID,
		WorkOrderCode:  req.WorkOrderCode,
		ProductID:      req.ProductID,
		ProductCode:    req.ProductCode,
		ProductName:    req.ProductName,
		OfflineType:    req.OfflineType,
		OfflineReason:  req.OfflineReason,
		OfflineQty:     req.OfflineQty,
		ProcessRouteID: req.ProcessRouteID,
		CurrentOpID:    req.CurrentOpID,
		CurrentOpName:  req.CurrentOpName,
		Status:         "OPEN",
		HandleResult:   "PENDING",
		OperatorID:     req.OperatorID,
		OperatorName:   req.OperatorName,
		Remark:         req.Remark,
		CreatedBy:      username,
		UpdatedBy:      username,
	}

	// 转换明细
	var items []model.ProductionOfflineItem
	for _, item := range req.Items {
		items = append(items, model.ProductionOfflineItem{
			TenantID:   tenantID,
			OfflineID:  0,
			SerialNo:   item.SerialNo,
			BatchNo:    item.BatchNo,
			OfflineQty: item.OfflineQty,
			Remark:     item.Remark,
		})
	}

	if err := s.offlineRepo.CreateWithItems(offline, items); err != nil {
		return nil, fmt.Errorf("failed to create offline record: %w", err)
	}

	return offline, nil
}

func (s *ProductionOfflineService) Update(ctx context.Context, id uint64, req *model.ProductionOfflineUpdateRequest, username string) error {
	updates := map[string]interface{}{
		"updated_by": username,
	}

	if req.OfflineType != "" {
		updates["offline_type"] = req.OfflineType
	}
	if req.OfflineReason != "" {
		updates["offline_reason"] = req.OfflineReason
	}
	if req.OfflineQty > 0 {
		updates["offline_qty"] = req.OfflineQty
	}
	if req.CurrentOpID > 0 {
		updates["current_op_id"] = req.CurrentOpID
	}
	if req.CurrentOpName != "" {
		updates["current_op_name"] = req.CurrentOpName
	}
	if req.OperatorID > 0 {
		updates["operator_id"] = req.OperatorID
	}
	if req.OperatorName != "" {
		updates["operator_name"] = req.OperatorName
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	return s.offlineRepo.UpdateWithContext(ctx, id, updates)
}

func (s *ProductionOfflineService) Delete(ctx context.Context, id uint64) error {
	return s.offlineRepo.Delete(id)
}

func (s *ProductionOfflineService) HandleOffline(ctx context.Context, id uint64, req *model.ProductionOfflineHandleRequest, username string) error {
	offline, err := s.offlineRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("offline record not found: %w", err)
	}

	updates := map[string]interface{}{
		"handle_method": req.HandleMethod,
		"handle_qty":    req.HandleQty,
		"handle_result": "COMPLETED",
		"status":        "HANDLED",
		"updated_by":    username,
	}

	// 根据处理方式更新对应数量
	switch req.HandleMethod {
	case "REWORK":
		updates["rework_order_id"] = req.ReworkOrderID
	case "SCRAP":
		updates["scrap_qty"] = req.HandleQty
	case "DOWNGRADE":
		updates["downgrade_qty"] = req.HandleQty
	}

	if err := s.offlineRepo.UpdateWithContext(ctx, id, updates); err != nil {
		return fmt.Errorf("failed to update offline record: %w", err)
	}

	// 处理明细
	if len(req.Items) > 0 {
		var items []model.ProductionOfflineItem
		for _, item := range req.Items {
			items = append(items, model.ProductionOfflineItem{
				ID:           item.ID,
				OfflineID:    id,
				TenantID:     offline.TenantID,
				HandleMethod: item.HandleMethod,
				HandleQty:    item.HandleQty,
				HandleResult: item.HandleResult,
				Remark:       item.Remark,
			})
		}
		if err := s.offlineRepo.UpdateItems(id, items); err != nil {
			return fmt.Errorf("failed to update items: %w", err)
		}
	}

	return nil
}

func (s *ProductionOfflineService) GetItems(ctx context.Context, offlineID uint64) ([]model.ProductionOfflineItem, error) {
	return s.offlineRepo.GetItemsByOfflineID(offlineID)
}

func (s *ProductionOfflineService) GenerateOfflineCode(tenantID uint64) (string, error) {
	now := time.Now()
	prefix := fmt.Sprintf("OFF-%s-", now.Format("20060102"))

	// 查询当天最大编号
	dateStr := now.Format("2006-01-02")
	_, count, err := s.offlineRepo.List(nil, tenantID, map[string]interface{}{
		"start_date": dateStr,
		"end_date":   dateStr + " 23:59:59",
	})
	if err != nil {
		// 忽略错误，使用默认编号
		return fmt.Sprintf("%s0001", prefix), nil
	}

	return fmt.Sprintf("%s%04d", prefix, count+1), nil
}
