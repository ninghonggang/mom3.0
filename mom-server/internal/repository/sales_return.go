package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// SalesReturnRepository 销售退货仓储
type SalesReturnRepository struct {
	db *gorm.DB
}

func NewSalesReturnRepository(db *gorm.DB) *SalesReturnRepository {
	return &SalesReturnRepository{db: db}
}

func (r *SalesReturnRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.SalesReturn, int64, error) {
	var list []model.SalesReturn
	var total int64

	q := r.db.WithContext(ctx).Model(&model.SalesReturn{}).Where("tenant_id = ?", tenantID)

	if salesOrderID, ok := query["sales_order_id"]; ok && salesOrderID.(int64) > 0 {
		q = q.Where("sales_order_id = ?", salesOrderID)
	}
	if customerID, ok := query["customer_id"]; ok && customerID.(int64) > 0 {
		q = q.Where("customer_id = ?", customerID)
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

func (r *SalesReturnRepository) GetByID(ctx context.Context, id uint) (*model.SalesReturn, error) {
	var ret model.SalesReturn
	err := r.db.WithContext(ctx).Preload("Items").First(&ret, id).Error
	return &ret, err
}

func (r *SalesReturnRepository) GetByNo(ctx context.Context, tenantID int64, no string) (*model.SalesReturn, error) {
	var ret model.SalesReturn
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND return_no = ?", tenantID, no).First(&ret).Error
	return &ret, err
}

func (r *SalesReturnRepository) Create(ctx context.Context, ret *model.SalesReturn) error {
	return r.db.WithContext(ctx).Create(ret).Error
}

func (r *SalesReturnRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.SalesReturn{}).Where("id = ?", id).Updates(updates).Error
}

func (r *SalesReturnRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		tx.Where("return_id = ?", id).Delete(&model.SalesReturnItem{})
		return tx.Delete(&model.SalesReturn{}, id).Error
	})
}

// SalesReturnItemRepository 销售退货明细仓储
type SalesReturnItemRepository struct {
	db *gorm.DB
}

func NewSalesReturnItemRepository(db *gorm.DB) *SalesReturnItemRepository {
	return &SalesReturnItemRepository{db: db}
}

func (r *SalesReturnItemRepository) ListByReturnID(ctx context.Context, returnID int64) ([]model.SalesReturnItem, error) {
	var items []model.SalesReturnItem
	err := r.db.WithContext(ctx).Where("return_id = ?", returnID).Order("line_no").Find(&items).Error
	return items, err
}

func (r *SalesReturnItemRepository) CreateBatch(ctx context.Context, items []model.SalesReturnItem) error {
	return r.db.WithContext(ctx).Create(&items).Error
}

func (r *SalesReturnItemRepository) DeleteByReturnID(ctx context.Context, returnID int64) error {
	return r.db.WithContext(ctx).Where("return_id = ?", returnID).Delete(&model.SalesReturnItem{}).Error
}

func (r *SalesReturnItemRepository) UpdateReturnQty(ctx context.Context, id uint, returnQty float64) error {
	return r.db.WithContext(ctx).Model(&model.SalesReturnItem{}).Where("id = ?", id).Updates(map[string]interface{}{
		"return_qty": returnQty,
	}).Error
}