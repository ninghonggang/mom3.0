package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// ScpSupplierContactRepository 供应商联系人仓库
type ScpSupplierContactRepository struct {
	db *gorm.DB
}

func NewScpSupplierContactRepository(db *gorm.DB) *ScpSupplierContactRepository {
	return &ScpSupplierContactRepository{db: db}
}

func (r *ScpSupplierContactRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.ScpSupplierContact, int64, error) {
	var list []model.ScpSupplierContact
	var total int64

	q := r.db.WithContext(ctx).Model(&model.ScpSupplierContact{}).Where("tenant_id = ?", tenantID)

	if supplierID, ok := query["supplier_id"]; ok && supplierID != nil {
		q = q.Where("supplier_id = ?", supplierID)
	}
	if isActive, ok := query["is_active"]; ok && isActive != nil {
		q = q.Where("is_active = ?", isActive)
	}

	q.Count(&total)
	q = q.Order("is_primary DESC, id DESC")

	if page, ok := query["page"].(int); ok && page > 0 {
		limit := 20
		if limitVal, ok := query["limit"].(int); ok && limitVal > 0 {
			limit = limitVal
		}
		q = q.Offset((page - 1) * limit).Limit(limit)
	}

	err := q.Find(&list).Error
	return list, total, err
}

func (r *ScpSupplierContactRepository) GetByID(ctx context.Context, id uint) (*model.ScpSupplierContact, error) {
	var contact model.ScpSupplierContact
	err := r.db.WithContext(ctx).First(&contact, id).Error
	return &contact, err
}

func (r *ScpSupplierContactRepository) ListBySupplier(ctx context.Context, supplierID int64) ([]model.ScpSupplierContact, error) {
	var list []model.ScpSupplierContact
	err := r.db.WithContext(ctx).Where("supplier_id = ? AND is_active = ?", supplierID, true).Order("is_primary DESC, id DESC").Find(&list).Error
	return list, err
}

func (r *ScpSupplierContactRepository) Create(ctx context.Context, contact *model.ScpSupplierContact) error {
	return r.db.WithContext(ctx).Create(contact).Error
}

func (r *ScpSupplierContactRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ScpSupplierContact{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ScpSupplierContactRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.ScpSupplierContact{}, id).Error
}

// ScpSupplierBankRepository 供应商银行账户仓库
type ScpSupplierBankRepository struct {
	db *gorm.DB
}

func NewScpSupplierBankRepository(db *gorm.DB) *ScpSupplierBankRepository {
	return &ScpSupplierBankRepository{db: db}
}

func (r *ScpSupplierBankRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.ScpSupplierBank, int64, error) {
	var list []model.ScpSupplierBank
	var total int64

	q := r.db.WithContext(ctx).Model(&model.ScpSupplierBank{}).Where("tenant_id = ?", tenantID)

	if supplierID, ok := query["supplier_id"]; ok && supplierID != nil {
		q = q.Where("supplier_id = ?", supplierID)
	}
	if isActive, ok := query["is_active"]; ok && isActive != nil {
		q = q.Where("is_active = ?", isActive)
	}

	q.Count(&total)
	q = q.Order("is_primary DESC, id DESC")

	if page, ok := query["page"].(int); ok && page > 0 {
		limit := 20
		if limitVal, ok := query["limit"].(int); ok && limitVal > 0 {
			limit = limitVal
		}
		q = q.Offset((page - 1) * limit).Limit(limit)
	}

	err := q.Find(&list).Error
	return list, total, err
}

func (r *ScpSupplierBankRepository) GetByID(ctx context.Context, id uint) (*model.ScpSupplierBank, error) {
	var bank model.ScpSupplierBank
	err := r.db.WithContext(ctx).First(&bank, id).Error
	return &bank, err
}

func (r *ScpSupplierBankRepository) ListBySupplier(ctx context.Context, supplierID int64) ([]model.ScpSupplierBank, error) {
	var list []model.ScpSupplierBank
	err := r.db.WithContext(ctx).Where("supplier_id = ? AND is_active = ?", supplierID, true).Order("is_primary DESC, id DESC").Find(&list).Error
	return list, err
}

func (r *ScpSupplierBankRepository) Create(ctx context.Context, bank *model.ScpSupplierBank) error {
	return r.db.WithContext(ctx).Create(bank).Error
}

func (r *ScpSupplierBankRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ScpSupplierBank{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ScpSupplierBankRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.ScpSupplierBank{}, id).Error
}
