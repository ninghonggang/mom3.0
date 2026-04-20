package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// ProductionReturnRepository 生产退料仓储
type ProductionReturnRepository struct {
	db *gorm.DB
}

func NewProductionReturnRepository(db *gorm.DB) *ProductionReturnRepository {
	return &ProductionReturnRepository{db: db}
}

func (r *ProductionReturnRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.ProductionReturn, int64, error) {
	var list []model.ProductionReturn
	var total int64

	q := r.db.WithContext(ctx).Model(&model.ProductionReturn{}).Where("tenant_id = ?", tenantID)

	if orderID, ok := query["production_order_id"]; ok && orderID.(int64) > 0 {
		q = q.Where("production_order_id = ?", orderID)
	}
	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}
	if returnType, ok := query["return_type"]; ok && returnType != "" {
		q = q.Where("return_type = ?", returnType)
	}

	q.Count(&total)

	page := 1
	limit := 20
	if p, ok := query["page"].(int); ok && p > 0 {
		page = p
	}
	if l, ok := query["limit"].(int); ok && l > 0 {
		limit = l
	}
	q = q.Preload("Items").Offset((page-1)*limit).Limit(limit).Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *ProductionReturnRepository) GetByID(ctx context.Context, id uint) (*model.ProductionReturn, error) {
	var ret model.ProductionReturn
	err := r.db.WithContext(ctx).Preload("Items").First(&ret, id).Error
	return &ret, err
}

func (r *ProductionReturnRepository) GetByNo(ctx context.Context, tenantID int64, no string) (*model.ProductionReturn, error) {
	var ret model.ProductionReturn
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND return_no = ?", tenantID, no).First(&ret).Error
	return &ret, err
}

func (r *ProductionReturnRepository) Create(ctx context.Context, ret *model.ProductionReturn) error {
	return r.db.WithContext(ctx).Create(ret).Error
}

func (r *ProductionReturnRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ProductionReturn{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ProductionReturnRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		tx.Where("return_id = ?", id).Delete(&model.ProductionReturnItem{})
		return tx.Delete(&model.ProductionReturn{}, id).Error
	})
}

// ProductionReturnItemRepository 生产退料明细仓储
type ProductionReturnItemRepository struct {
	db *gorm.DB
}

func NewProductionReturnItemRepository(db *gorm.DB) *ProductionReturnItemRepository {
	return &ProductionReturnItemRepository{db: db}
}

func (r *ProductionReturnItemRepository) ListByReturnID(ctx context.Context, returnID int64) ([]model.ProductionReturnItem, error) {
	var items []model.ProductionReturnItem
	err := r.db.WithContext(ctx).Where("return_id = ?", returnID).Order("line_no").Find(&items).Error
	return items, err
}

func (r *ProductionReturnItemRepository) CreateBatch(ctx context.Context, items []model.ProductionReturnItem) error {
	return r.db.WithContext(ctx).Create(&items).Error
}

func (r *ProductionReturnItemRepository) DeleteByReturnID(ctx context.Context, returnID int64) error {
	return r.db.WithContext(ctx).Where("return_id = ?", returnID).Delete(&model.ProductionReturnItem{}).Error
}

func (r *ProductionReturnItemRepository) UpdateReturnQty(ctx context.Context, id uint, returnQty float64) error {
	return r.db.WithContext(ctx).Model(&model.ProductionReturnItem{}).Where("id = ?", id).Updates(map[string]interface{}{
		"return_qty": returnQty,
	}).Error
}
