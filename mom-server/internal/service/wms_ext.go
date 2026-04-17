package service

import (
	"context"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"time"
)

// ========== 调拨管理 ==========

type TransferOrderService struct {
	repo       *repository.TransferOrderRepository
	itemRepo   *repository.TransferOrderItemRepository
	traceRepo  *repository.TransferTraceRepository
	lotRepo    *repository.TransferLotRepository
}

func NewTransferOrderService(repo *repository.TransferOrderRepository, itemRepo *repository.TransferOrderItemRepository) *TransferOrderService {
	return &TransferOrderService{repo: repo, itemRepo: itemRepo}
}

func (s *TransferOrderService) List(ctx context.Context, query string) ([]model.TransferOrder, int64, error) {
	return s.repo.List(ctx, 1, query)
}

func (s *TransferOrderService) GetByID(ctx context.Context, id uint) (*model.TransferOrder, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TransferOrderService) Create(ctx context.Context, order *model.TransferOrder) error {
	if order.TenantID == 0 {
		order.TenantID = 1
	}
	return s.repo.Create(ctx, order)
}

func (s *TransferOrderService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *TransferOrderService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *TransferOrderService) GetItems(ctx context.Context, transferID int64) ([]model.TransferOrderItem, error) {
	return s.itemRepo.ListByTransferID(ctx, transferID)
}

func (s *TransferOrderService) AddItem(ctx context.Context, item *model.TransferOrderItem) error {
	if item.TenantID == 0 {
		item.TenantID = 1
	}
	return s.itemRepo.Create(ctx, item)
}

func (s *TransferOrderService) Submit(ctx context.Context, id uint) error {
	return s.repo.Update(ctx, id, map[string]interface{}{
		"status":       "SUBMITTED",
		"request_time": time.Now(),
	})
}

func (s *TransferOrderService) Approve(ctx context.Context, id uint, approved bool, comment *string) error {
	status := "APPROVED"
	if !approved {
		status = "REJECTED"
	}
	updates := map[string]interface{}{
		"status":          status,
		"approval_comment": comment,
	}
	if approved {
		now := time.Now()
		updates["approved_time"] = now
	}
	return s.repo.Update(ctx, id, updates)
}

func (s *TransferOrderService) StartTransfer(ctx context.Context, id uint) error {
	now := time.Now()
	return s.repo.Update(ctx, id, map[string]interface{}{
		"status":            "IN_TRANSIT",
		"actual_start_time": now,
	})
}

func (s *TransferOrderService) Ship(ctx context.Context, id uint, operatorID int64, operatorName string) error {
	return s.repo.Update(ctx, id, map[string]interface{}{
		"status":       "IN_TRANSIT",
		"operator_id":  operatorID,
		"operator_name": operatorName,
	})
}

func (s *TransferOrderService) Receive(ctx context.Context, id uint, operatorID int64, operatorName string) error {
	return s.repo.Update(ctx, id, map[string]interface{}{
		"operator_id":   operatorID,
		"operator_name": operatorName,
	})
}

func (s *TransferOrderService) Complete(ctx context.Context, id uint) error {
	now := time.Now()
	return s.repo.Update(ctx, id, map[string]interface{}{
		"status":               "COMPLETED",
		"actual_complete_time": now,
	})
}

func (s *TransferOrderService) Cancel(ctx context.Context, id uint, reason string) error {
	return s.repo.Update(ctx, id, map[string]interface{}{
		"status":  "CANCELLED",
		"remark":  reason,
	})
}

func (s *TransferOrderService) GetTraces(ctx context.Context, orderID uint) ([]model.TransferTrace, error) {
	if s.traceRepo == nil {
		return []model.TransferTrace{}, nil
	}
	return s.traceRepo.ListByOrderID(ctx, orderID)
}

// ========== 盘点管理 ==========

type StockCheckService struct {
	repo    *repository.StockCheckRepository
	itemRepo *repository.StockCheckItemRepository
}

func NewStockCheckService(repo *repository.StockCheckRepository, itemRepo *repository.StockCheckItemRepository) *StockCheckService {
	return &StockCheckService{repo: repo, itemRepo: itemRepo}
}

func (s *StockCheckService) List(ctx context.Context, query string) ([]model.StockCheck, int64, error) {
	return s.repo.List(ctx, 1, query)
}

func (s *StockCheckService) GetByID(ctx context.Context, id uint) (*model.StockCheck, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *StockCheckService) Create(ctx context.Context, check *model.StockCheck) error {
	if check.TenantID == 0 {
		check.TenantID = 1
	}
	return s.repo.Create(ctx, check)
}

func (s *StockCheckService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *StockCheckService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *StockCheckService) GetItems(ctx context.Context, checkID int64) ([]model.StockCheckItem, error) {
	return s.itemRepo.ListByCheckID(ctx, checkID)
}

func (s *StockCheckService) AddItem(ctx context.Context, item *model.StockCheckItem) error {
	if item.TenantID == 0 {
		item.TenantID = 1
	}
	return s.itemRepo.Create(ctx, item)
}

func (s *StockCheckService) UpdateItem(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.itemRepo.Update(ctx, id, updates)
}

func (s *StockCheckService) Submit(ctx context.Context, id uint) error {
	return s.repo.Update(ctx, id, map[string]interface{}{
		"status": "IN_PROGRESS",
	})
}

func (s *StockCheckService) StartCheck(ctx context.Context, id uint) error {
	now := time.Now().Format("2006-01-02")
	return s.repo.Update(ctx, id, map[string]interface{}{
		"status":             "COUNTING",
		"actual_start_date": now,
	})
}

func (s *StockCheckService) CompleteCheck(ctx context.Context, id uint) error {
	now := time.Now().Format("2006-01-02")
	// 计算差异
	items, _ := s.itemRepo.ListByCheckID(ctx, int64(id))
	var varianceCount int
	var totalCounted int
	for _, item := range items {
		totalCounted++
		if item.VarianceQty != 0 {
			varianceCount++
		}
	}
	varianceRate := float64(0)
	if totalCounted > 0 {
		varianceRate = float64(varianceCount) / float64(totalCounted) * 100
	}
	return s.repo.Update(ctx, id, map[string]interface{}{
		"status":             "COMPLETED",
		"actual_end_date":    now,
		"variance_count":    varianceCount,
		"variance_rate":      varianceRate,
		"counted_locations":  totalCounted,
	})
}

func (s *StockCheckService) ApproveCheck(ctx context.Context, id uint, approved bool, comment *string) error {
	status := "APPROVED"
	if !approved {
		status = "REJECTED"
	}
	updates := map[string]interface{}{
		"approval_status":   status,
		"approval_comment":   comment,
	}
	if approved {
		now := time.Now()
		updates["approved_time"] = now
	}
	return s.repo.Update(ctx, id, updates)
}

func (s *StockCheckService) CountItem(ctx context.Context, checkID, itemID uint, countedQty float64, counterID int64, counterName string, countTime *time.Time) error {
	// 获取当前明细计算差异
	item, err := s.itemRepo.GetByID(ctx, itemID)
	if err != nil {
		return err
	}
	varianceQty := countedQty - item.SystemQty
	varianceAmount := varianceQty * 0 // 实际应按单价计算
	updates := map[string]interface{}{
		"counted_qty":     countedQty,
		"variance_qty":    varianceQty,
		"variance_amount": varianceAmount,
		"count_status":    "COUNTED",
		"counter_id":      counterID,
		"counter_name":    counterName,
		"count_time":      countTime,
	}
	return s.itemRepo.Update(ctx, itemID, updates)
}

func (s *StockCheckService) HandleVariance(ctx context.Context, checkID, itemID uint, handleMethod string, handleQty float64, handlerID int64, handlerName string) error {
	now := time.Now()
	updates := map[string]interface{}{
		"handle_status": "PROCESSED",
		"handle_method":  handleMethod,
		"handle_qty":     handleQty,
		"handler_id":     handlerID,
		"handler_name":   handlerName,
		"handle_time":    now,
	}
	return s.itemRepo.Update(ctx, itemID, updates)
}

func (s *StockCheckService) Recount(ctx context.Context, checkID, itemID uint, recountQty float64, recountBy int64, recountName string, recountTime *time.Time) error {
	updates := map[string]interface{}{
		"recount_qty":   recountQty,
		"recount_by":    recountBy,
		"recount_by_name": recountName,
		"recount_time":  recountTime,
	}
	return s.itemRepo.Update(ctx, itemID, updates)
}

func (s *StockCheckService) GetVariances(ctx context.Context, checkID uint) ([]model.StockCheckItem, error) {
	return s.itemRepo.ListVariancesByCheckID(ctx, int64(checkID))
}

// ========== 线边库位 ==========

type SideLocationService struct {
	repo *repository.SideLocationRepository
}

func NewSideLocationService(repo *repository.SideLocationRepository) *SideLocationService {
	return &SideLocationService{repo: repo}
}

func (s *SideLocationService) List(ctx context.Context, query string) ([]model.SideLocation, int64, error) {
	return s.repo.List(ctx, 1, query)
}

func (s *SideLocationService) GetByID(ctx context.Context, id uint) (*model.SideLocation, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *SideLocationService) Create(ctx context.Context, loc *model.SideLocation) error {
	if loc.TenantID == 0 {
		loc.TenantID = 1
	}
	return s.repo.Create(ctx, loc)
}

func (s *SideLocationService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *SideLocationService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

// ========== 看板拉动 ==========

type KanbanPullService struct {
	repo *repository.KanbanPullRepository
}

func NewKanbanPullService(repo *repository.KanbanPullRepository) *KanbanPullService {
	return &KanbanPullService{repo: repo}
}

func (s *KanbanPullService) List(ctx context.Context, query string) ([]model.KanbanPull, int64, error) {
	return s.repo.List(ctx, 1, query)
}

func (s *KanbanPullService) GetByID(ctx context.Context, id uint) (*model.KanbanPull, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *KanbanPullService) Create(ctx context.Context, k *model.KanbanPull) error {
	if k.TenantID == 0 {
		k.TenantID = 1
	}
	return s.repo.Create(ctx, k)
}

func (s *KanbanPullService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *KanbanPullService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *KanbanPullService) Trigger(ctx context.Context, id uint) error {
	// 看板触发后减少当前库存
	k, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	newQty := k.CurrentQty - k.TriggerQty
	if newQty < 0 {
		newQty = 0
	}
	return s.repo.Update(ctx, id, map[string]interface{}{
		"current_qty": newQty,
	})
}
