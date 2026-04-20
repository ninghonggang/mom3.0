package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// PurchaseReturnRepository 采购退货仓储
type PurchaseReturnRepository struct {
	db *gorm.DB
}

func NewPurchaseReturnRepository(db *gorm.DB) *PurchaseReturnRepository {
	return &PurchaseReturnRepository{db: db}
}

func (r *PurchaseReturnRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.PurchaseReturn, int64, error) {
	var list []model.PurchaseReturn
	var total int64

	q := r.db.WithContext(ctx).Model(&model.PurchaseReturn{}).Where("tenant_id = ?", tenantID)

	if purchaseOrderID, ok := query["purchase_order_id"]; ok && purchaseOrderID.(int64) > 0 {
		q = q.Where("purchase_order_id = ?", purchaseOrderID)
	}
	if supplierID, ok := query["supplier_id"]; ok && supplierID.(int64) > 0 {
		q = q.Where("supplier_id = ?", supplierID)
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

func (r *PurchaseReturnRepository) GetByID(ctx context.Context, id uint) (*model.PurchaseReturn, error) {
	var ret model.PurchaseReturn
	err := r.db.WithContext(ctx).Preload("Items").First(&ret, id).Error
	return &ret, err
}

func (r *PurchaseReturnRepository) GetByNo(ctx context.Context, tenantID int64, no string) (*model.PurchaseReturn, error) {
	var ret model.PurchaseReturn
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND return_no = ?", tenantID, no).First(&ret).Error
	return &ret, err
}

func (r *PurchaseReturnRepository) Create(ctx context.Context, ret *model.PurchaseReturn) error {
	return r.db.WithContext(ctx).Create(ret).Error
}

func (r *PurchaseReturnRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.PurchaseReturn{}).Where("id = ?", id).Updates(updates).Error
}

func (r *PurchaseReturnRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		tx.Where("return_id = ?", id).Delete(&model.PurchaseReturnItem{})
		return tx.Delete(&model.PurchaseReturn{}, id).Error
	})
}

// PurchaseReturnItemRepository 采购退货明细仓储
type PurchaseReturnItemRepository struct {
	db *gorm.DB
}

func NewPurchaseReturnItemRepository(db *gorm.DB) *PurchaseReturnItemRepository {
	return &PurchaseReturnItemRepository{db: db}
}

func (r *PurchaseReturnItemRepository) ListByReturnID(ctx context.Context, returnID int64) ([]model.PurchaseReturnItem, error) {
	var items []model.PurchaseReturnItem
	err := r.db.WithContext(ctx).Where("return_id = ?", returnID).Order("line_no").Find(&items).Error
	return items, err
}

func (r *PurchaseReturnItemRepository) CreateBatch(ctx context.Context, items []model.PurchaseReturnItem) error {
	return r.db.WithContext(ctx).Create(&items).Error
}

func (r *PurchaseReturnItemRepository) DeleteByReturnID(ctx context.Context, returnID int64) error {
	return r.db.WithContext(ctx).Where("return_id = ?", returnID).Delete(&model.PurchaseReturnItem{}).Error
}

func (r *PurchaseReturnItemRepository) UpdateReturnQty(ctx context.Context, id uint, returnQty float64) error {
	return r.db.WithContext(ctx).Model(&model.PurchaseReturnItem{}).Where("id = ?", id).Updates(map[string]interface{}{
		"return_qty": returnQty,
	}).Error
}