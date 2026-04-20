package repository

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

// SupplierASNRepository ASN仓库
type SupplierASNRepository struct {
	db *gorm.DB
}

func NewSupplierASNRepository(db *gorm.DB) *SupplierASNRepository {
	return &SupplierASNRepository{db: db}
}

// Create 创建ASN
func (r *SupplierASNRepository) Create(ctx context.Context, asn *model.SupplierASN) error {
	return r.db.WithContext(ctx).Create(asn).Error
}

// GetByID 根据ID查询
func (r *SupplierASNRepository) GetByID(ctx context.Context, id int64) (*model.SupplierASN, error) {
	var asn model.SupplierASN
	err := r.db.WithContext(ctx).First(&asn, id).Error
	if err != nil {
		return nil, err
	}
	return &asn, nil
}

// GetByASNNo 根据ASN编号查询
func (r *SupplierASNRepository) GetByASNNo(ctx context.Context, asnNo string) (*model.SupplierASN, error) {
	var asn model.SupplierASN
	err := r.db.WithContext(ctx).Where("asn_no = ?", asnNo).First(&asn).Error
	if err != nil {
		return nil, err
	}
	return &asn, nil
}

// Update 更新ASN
func (r *SupplierASNRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.SupplierASN{}).Where("id = ?", id).Updates(updates).Error
}

// UpdateStatus 更新状态
func (r *SupplierASNRepository) UpdateStatus(ctx context.Context, id int64, status string) error {
	return r.db.WithContext(ctx).Model(&model.SupplierASN{}).Where("id = ?", id).Update("status", status).Error
}

// List 查询ASN列表
func (r *SupplierASNRepository) List(ctx context.Context, q *model.SupplierASNQuery) ([]model.SupplierASN, int64, error) {
	var list []model.SupplierASN
	var total int64

	query := r.db.WithContext(ctx).Model(&model.SupplierASN{})

	if q.TenantID > 0 {
		query = query.Where("tenant_id = ?", q.TenantID)
	}
	if q.SupplierCode != "" {
		query = query.Where("supplier_code LIKE ?", "%"+q.SupplierCode+"%")
	}
	if q.Status != "" {
		query = query.Where("status = ?", q.Status)
	}
	if q.StartDate != "" {
		query = query.Where("delivery_date >= ?", q.StartDate)
	}
	if q.EndDate != "" {
		query = query.Where("delivery_date <= ?", q.EndDate+" 23:59:59")
	}

	query.Count(&total)

	page := q.Page
	if page < 1 {
		page = 1
	}
	pageSize := q.PageSize
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&list).Error
	return list, total, err
}

// ListBySupplier 按供应商查询
func (r *SupplierASNRepository) ListBySupplier(ctx context.Context, tenantID, supplierID int64, status string) ([]model.SupplierASN, error) {
	var list []model.SupplierASN
	query := r.db.WithContext(ctx).Where("tenant_id = ? AND supplier_id = ?", tenantID, supplierID)
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Order("created_at DESC").Find(&list).Error
	return list, err
}

// Delete 删除ASN
func (r *SupplierASNRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 删除明细
		if err := tx.Where("asn_id = ?", id).Delete(&model.SupplierASNItem{}).Error; err != nil {
			return err
		}
		// 删除主单
		return tx.Delete(&model.SupplierASN{}, id).Error
	})
}

// GenerateASNNo 生成ASN编号
func (r *SupplierASNRepository) GenerateASNNo(ctx context.Context) (string, error) {
	now := time.Now()
	prefix := fmt.Sprintf("ASN-%s-", now.Format("20060102"))
	var count int64
	r.db.WithContext(ctx).Model(&model.SupplierASN{}).Where("asn_no LIKE ?", prefix+"%").Count(&count)
	return fmt.Sprintf("%s%04d", prefix, count+1), nil
}

// CreateItem 创建ASN明细
func (r *SupplierASNRepository) CreateItem(ctx context.Context, item *model.SupplierASNItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

// GetItemsByASNID 获取ASN明细
func (r *SupplierASNRepository) GetItemsByASNID(ctx context.Context, asnID int64) ([]model.SupplierASNItem, error) {
	var items []model.SupplierASNItem
	err := r.db.WithContext(ctx).Where("asn_id = ?", asnID).Order("line_no ASC").Find(&items).Error
	return items, err
}

// UpdateItem 更新ASN明细
func (r *SupplierASNRepository) UpdateItem(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.SupplierASNItem{}).Where("id = ?", id).Updates(updates).Error
}

// DeleteItems 删除ASN明细
func (r *SupplierASNRepository) DeleteItems(ctx context.Context, asnID int64) error {
	return r.db.WithContext(ctx).Where("asn_id = ?", asnID).Delete(&model.SupplierASNItem{}).Error
}

// UpdateItemsQty 更新ASN明细的已收货数量
func (r *SupplierASNRepository) UpdateItemsQty(ctx context.Context, asnID int64, itemID int64, receivedQty float64) error {
	return r.db.WithContext(ctx).Model(&model.SupplierASNItem{}).
		Where("id = ? AND asn_id = ?", itemID, asnID).
		Update("received_qty", receivedQty).Error
}
