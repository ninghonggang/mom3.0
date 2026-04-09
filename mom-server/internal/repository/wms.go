package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type WarehouseRepository struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) *WarehouseRepository {
	return &WarehouseRepository{db: db}
}

func (r *WarehouseRepository) List(ctx context.Context, tenantID int64) ([]model.Warehouse, int64, error) {
	var list []model.Warehouse
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Warehouse{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *WarehouseRepository) GetByID(ctx context.Context, id uint) (*model.Warehouse, error) {
	var warehouse model.Warehouse
	err := r.db.WithContext(ctx).First(&warehouse, id).Error
	return &warehouse, err
}

func (r *WarehouseRepository) Create(ctx context.Context, warehouse *model.Warehouse) error {
	return r.db.WithContext(ctx).Create(warehouse).Error
}

func (r *WarehouseRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Warehouse{}).Where("id = ?", id).Updates(updates).Error
}

func (r *WarehouseRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Warehouse{}, id).Error
}

type LocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {
	return &LocationRepository{db: db}
}

func (r *LocationRepository) List(ctx context.Context, tenantID int64) ([]model.Location, int64, error) {
	var list []model.Location
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Location{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *LocationRepository) GetByID(ctx context.Context, id uint) (*model.Location, error) {
	var location model.Location
	err := r.db.WithContext(ctx).First(&location, id).Error
	return &location, err
}

func (r *LocationRepository) Create(ctx context.Context, location *model.Location) error {
	return r.db.WithContext(ctx).Create(location).Error
}

func (r *LocationRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Location{}).Where("id = ?", id).Updates(updates).Error
}

func (r *LocationRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Location{}, id).Error
}

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{db: db}
}

func (r *InventoryRepository) List(ctx context.Context, tenantID int64) ([]model.Inventory, int64, error) {
	var list []model.Inventory
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Inventory{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *InventoryRepository) GetByID(ctx context.Context, id uint) (*model.Inventory, error) {
	var inventory model.Inventory
	err := r.db.WithContext(ctx).First(&inventory, id).Error
	return &inventory, err
}

func (r *InventoryRepository) Create(ctx context.Context, inventory *model.Inventory) error {
	return r.db.WithContext(ctx).Create(inventory).Error
}

func (r *InventoryRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Inventory{}).Where("id = ?", id).Updates(updates).Error
}

func (r *InventoryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Inventory{}, id).Error
}

// ListByMaterialID 根据物料ID获取库存列表
func (r *InventoryRepository) ListByMaterialID(ctx context.Context, materialID int64) ([]model.Inventory, int64, error) {
	var list []model.Inventory
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Inventory{}).Where("material_id = ?", materialID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Where("material_id = ?", materialID).Order("id DESC").Find(&list).Error
	return list, total, err
}

// GetTotalAvailableByMaterialID 获取物料可用库存总量
func (r *InventoryRepository) GetTotalAvailableByMaterialID(ctx context.Context, materialID int64) (float64, error) {
	var total float64
	err := r.db.WithContext(ctx).Model(&model.Inventory{}).
		Where("material_id = ?", materialID).
		Select("COALESCE(SUM(available_qty), 0)").
		Scan(&total).Error
	return total, err
}

// ========== 收货单 ==========

type ReceiveOrderRepository struct {
	db *gorm.DB
}

func NewReceiveOrderRepository(db *gorm.DB) *ReceiveOrderRepository {
	return &ReceiveOrderRepository{db: db}
}

func (r *ReceiveOrderRepository) List(ctx context.Context, tenantID int64) ([]model.ReceiveOrder, int64, error) {
	var list []model.ReceiveOrder
	var total int64

	query := r.db.WithContext(ctx).Model(&model.ReceiveOrder{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *ReceiveOrderRepository) GetByID(ctx context.Context, id uint) (*model.ReceiveOrder, error) {
	var order model.ReceiveOrder
	err := r.db.WithContext(ctx).First(&order, id).Error
	return &order, err
}

func (r *ReceiveOrderRepository) GetByReceiveNo(ctx context.Context, receiveNo string) (*model.ReceiveOrder, error) {
	var order model.ReceiveOrder
	err := r.db.WithContext(ctx).Where("receive_no = ?", receiveNo).First(&order).Error
	return &order, err
}

func (r *ReceiveOrderRepository) Create(ctx context.Context, order *model.ReceiveOrder) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *ReceiveOrderRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ReceiveOrder{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ReceiveOrderRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.ReceiveOrder{}, id).Error
}

type ReceiveOrderItemRepository struct {
	db *gorm.DB
}

func NewReceiveOrderItemRepository(db *gorm.DB) *ReceiveOrderItemRepository {
	return &ReceiveOrderItemRepository{db: db}
}

func (r *ReceiveOrderItemRepository) ListByReceiveID(ctx context.Context, receiveID uint) ([]model.ReceiveOrderItem, error) {
	var list []model.ReceiveOrderItem
	err := r.db.WithContext(ctx).Where("receive_id = ?", receiveID).Order("id ASC").Find(&list).Error
	return list, err
}

func (r *ReceiveOrderItemRepository) Create(ctx context.Context, item *model.ReceiveOrderItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *ReceiveOrderItemRepository) DeleteByReceiveID(ctx context.Context, receiveID uint) error {
	return r.db.WithContext(ctx).Where("receive_id = ?", receiveID).Delete(&model.ReceiveOrderItem{}).Error
}

// ========== 发货单 ==========

type DeliveryOrderRepository struct {
	db *gorm.DB
}

func NewDeliveryOrderRepository(db *gorm.DB) *DeliveryOrderRepository {
	return &DeliveryOrderRepository{db: db}
}

func (r *DeliveryOrderRepository) List(ctx context.Context, tenantID int64) ([]model.DeliveryOrder, int64, error) {
	var list []model.DeliveryOrder
	var total int64

	query := r.db.WithContext(ctx).Model(&model.DeliveryOrder{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *DeliveryOrderRepository) GetByID(ctx context.Context, id uint) (*model.DeliveryOrder, error) {
	var order model.DeliveryOrder
	err := r.db.WithContext(ctx).First(&order, id).Error
	return &order, err
}

func (r *DeliveryOrderRepository) GetByDeliveryNo(ctx context.Context, deliveryNo string) (*model.DeliveryOrder, error) {
	var order model.DeliveryOrder
	err := r.db.WithContext(ctx).Where("delivery_no = ?", deliveryNo).First(&order).Error
	return &order, err
}

func (r *DeliveryOrderRepository) Create(ctx context.Context, order *model.DeliveryOrder) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *DeliveryOrderRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.DeliveryOrder{}).Where("id = ?", id).Updates(updates).Error
}

func (r *DeliveryOrderRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.DeliveryOrder{}, id).Error
}

type DeliveryOrderItemRepository struct {
	db *gorm.DB
}

func NewDeliveryOrderItemRepository(db *gorm.DB) *DeliveryOrderItemRepository {
	return &DeliveryOrderItemRepository{db: db}
}

func (r *DeliveryOrderItemRepository) ListByDeliveryID(ctx context.Context, deliveryID uint) ([]model.DeliveryOrderItem, error) {
	var list []model.DeliveryOrderItem
	err := r.db.WithContext(ctx).Where("delivery_id = ?", deliveryID).Order("id ASC").Find(&list).Error
	return list, err
}

func (r *DeliveryOrderItemRepository) Create(ctx context.Context, item *model.DeliveryOrderItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *DeliveryOrderItemRepository) DeleteByDeliveryID(ctx context.Context, deliveryID uint) error {
	return r.db.WithContext(ctx).Where("delivery_id = ?", deliveryID).Delete(&model.DeliveryOrderItem{}).Error
}
