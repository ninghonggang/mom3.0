package service

import (
	"context"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// ========== 调拨管理 ==========

type TransferOrderService struct {
	repo *repository.TransferOrderRepository
	itemRepo *repository.TransferOrderItemRepository
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

// ========== 盘点管理 ==========

type StockCheckService struct {
	repo *repository.StockCheckRepository
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
