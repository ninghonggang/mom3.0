package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// ========== 调拨管理 ==========

type TransferOrderRepository struct {
	db *gorm.DB
}

func NewTransferOrderRepository(db *gorm.DB) *TransferOrderRepository {
	return &TransferOrderRepository{db: db}
}

func (r *TransferOrderRepository) List(ctx context.Context, tenantID int64, query string) ([]model.TransferOrder, int64, error) {
	var list []model.TransferOrder
	var total int64

	db := r.db.WithContext(ctx).Model(&model.TransferOrder{}).Where("tenant_id = ?", tenantID)
	if query != "" {
		db = db.Where("transfer_no LIKE ?", "%"+query+"%")
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *TransferOrderRepository) GetByID(ctx context.Context, id uint) (*model.TransferOrder, error) {
	var order model.TransferOrder
	err := r.db.WithContext(ctx).First(&order, id).Error
	return &order, err
}

func (r *TransferOrderRepository) Create(ctx context.Context, order *model.TransferOrder) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *TransferOrderRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.TransferOrder{}).Where("id = ?", id).Updates(updates).Error
}

func (r *TransferOrderRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.TransferOrder{}, id).Error
}

type TransferOrderItemRepository struct {
	db *gorm.DB
}

func NewTransferOrderItemRepository(db *gorm.DB) *TransferOrderItemRepository {
	return &TransferOrderItemRepository{db: db}
}

func (r *TransferOrderItemRepository) ListByTransferID(ctx context.Context, transferID int64) ([]model.TransferOrderItem, error) {
	var list []model.TransferOrderItem
	err := r.db.WithContext(ctx).Where("transfer_id = ?", transferID).Find(&list).Error
	return list, err
}

func (r *TransferOrderItemRepository) Create(ctx context.Context, item *model.TransferOrderItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

// ========== 盘点管理 ==========

type StockCheckRepository struct {
	db *gorm.DB
}

func NewStockCheckRepository(db *gorm.DB) *StockCheckRepository {
	return &StockCheckRepository{db: db}
}

func (r *StockCheckRepository) List(ctx context.Context, tenantID int64, query string) ([]model.StockCheck, int64, error) {
	var list []model.StockCheck
	var total int64

	db := r.db.WithContext(ctx).Model(&model.StockCheck{}).Where("tenant_id = ?", tenantID)
	if query != "" {
		db = db.Where("check_no LIKE ?", "%"+query+"%")
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *StockCheckRepository) GetByID(ctx context.Context, id uint) (*model.StockCheck, error) {
	var check model.StockCheck
	err := r.db.WithContext(ctx).First(&check, id).Error
	return &check, err
}

func (r *StockCheckRepository) Create(ctx context.Context, check *model.StockCheck) error {
	return r.db.WithContext(ctx).Create(check).Error
}

func (r *StockCheckRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.StockCheck{}).Where("id = ?", id).Updates(updates).Error
}

type StockCheckItemRepository struct {
	db *gorm.DB
}

func NewStockCheckItemRepository(db *gorm.DB) *StockCheckItemRepository {
	return &StockCheckItemRepository{db: db}
}

func (r *StockCheckItemRepository) ListByCheckID(ctx context.Context, checkID int64) ([]model.StockCheckItem, error) {
	var list []model.StockCheckItem
	err := r.db.WithContext(ctx).Where("check_id = ?", checkID).Find(&list).Error
	return list, err
}

func (r *StockCheckItemRepository) Create(ctx context.Context, item *model.StockCheckItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *StockCheckItemRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.StockCheckItem{}).Where("id = ?", id).Updates(updates).Error
}

// ========== 线边库位 ==========

type SideLocationRepository struct {
	db *gorm.DB
}

func NewSideLocationRepository(db *gorm.DB) *SideLocationRepository {
	return &SideLocationRepository{db: db}
}

func (r *SideLocationRepository) List(ctx context.Context, tenantID int64, query string) ([]model.SideLocation, int64, error) {
	var list []model.SideLocation
	var total int64

	db := r.db.WithContext(ctx).Model(&model.SideLocation{}).Where("tenant_id = ?", tenantID)
	if query != "" {
		db = db.Where("location_code LIKE ? OR location_name LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *SideLocationRepository) GetByID(ctx context.Context, id uint) (*model.SideLocation, error) {
	var loc model.SideLocation
	err := r.db.WithContext(ctx).First(&loc, id).Error
	return &loc, err
}

func (r *SideLocationRepository) Create(ctx context.Context, loc *model.SideLocation) error {
	return r.db.WithContext(ctx).Create(loc).Error
}

func (r *SideLocationRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.SideLocation{}).Where("id = ?", id).Updates(updates).Error
}

func (r *SideLocationRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.SideLocation{}, id).Error
}

// ========== 看板拉动 ==========

type KanbanPullRepository struct {
	db *gorm.DB
}

func NewKanbanPullRepository(db *gorm.DB) *KanbanPullRepository {
	return &KanbanPullRepository{db: db}
}

func (r *KanbanPullRepository) List(ctx context.Context, tenantID int64, query string) ([]model.KanbanPull, int64, error) {
	var list []model.KanbanPull
	var total int64

	db := r.db.WithContext(ctx).Model(&model.KanbanPull{}).Where("tenant_id = ?", tenantID)
	if query != "" {
		db = db.Where("kanban_no LIKE ? OR material_code LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Order("priority DESC").Find(&list).Error
	return list, total, err
}

func (r *KanbanPullRepository) GetByID(ctx context.Context, id uint) (*model.KanbanPull, error) {
	var k model.KanbanPull
	err := r.db.WithContext(ctx).First(&k, id).Error
	return &k, err
}

func (r *KanbanPullRepository) Create(ctx context.Context, k *model.KanbanPull) error {
	return r.db.WithContext(ctx).Create(k).Error
}

func (r *KanbanPullRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.KanbanPull{}).Where("id = ?", id).Updates(updates).Error
}

func (r *KanbanPullRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.KanbanPull{}, id).Error
}
