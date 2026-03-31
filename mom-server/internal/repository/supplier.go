package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type SupplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	return &SupplierRepository{db: db}
}

func (r *SupplierRepository) List(ctx context.Context, tenantID int64) ([]model.Supplier, int64, error) {
	var list []model.Supplier
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Supplier{}).Where("tenant_id = ?", tenantID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *SupplierRepository) GetByID(ctx context.Context, id uint) (*model.Supplier, error) {
	var supplier model.Supplier
	err := r.db.WithContext(ctx).First(&supplier, id).Error
	return &supplier, err
}

func (r *SupplierRepository) GetByCode(ctx context.Context, tenantID int64, code string) (*model.Supplier, error) {
	var supplier model.Supplier
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND code = ?", tenantID, code).First(&supplier).Error
	return &supplier, err
}

func (r *SupplierRepository) Create(ctx context.Context, supplier *model.Supplier) error {
	return r.db.WithContext(ctx).Create(supplier).Error
}

func (r *SupplierRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Supplier{}).Where("id = ?", id).Updates(updates).Error
}

func (r *SupplierRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Supplier{}, id).Error
}
