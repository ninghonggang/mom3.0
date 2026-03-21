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
