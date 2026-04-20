package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type SupplierMaterialRepository struct {
	db *gorm.DB
}

func NewSupplierMaterialRepository(db *gorm.DB) *SupplierMaterialRepository {
	return &SupplierMaterialRepository{db: db}
}

func (r *SupplierMaterialRepository) List(ctx context.Context, tenantID int64, query string) ([]model.SupplierMaterial, int64, error) {
	var list []model.SupplierMaterial
	var total int64

	q := r.db.WithContext(ctx).Model(&model.SupplierMaterial{}).Where("tenant_id = ?", tenantID)
	if query != "" {
		q = q.Where("material_code LIKE ? OR material_name LIKE ? OR supplier_part_no LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%")
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := q.Order("id DESC").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *SupplierMaterialRepository) GetByID(ctx context.Context, id uint) (*model.SupplierMaterial, error) {
	var item model.SupplierMaterial
	err := r.db.WithContext(ctx).First(&item, id).Error
	return &item, err
}

func (r *SupplierMaterialRepository) ListBySupplier(ctx context.Context, supplierID int64) ([]model.SupplierMaterial, error) {
	var list []model.SupplierMaterial
	err := r.db.WithContext(ctx).Where("supplier_id = ? AND status = 1", supplierID).Order("is_preferred DESC, id DESC").Find(&list).Error
	return list, err
}

func (r *SupplierMaterialRepository) ListByMaterial(ctx context.Context, materialID int64) ([]model.SupplierMaterial, error) {
	var list []model.SupplierMaterial
	err := r.db.WithContext(ctx).Where("material_id = ? AND status = 1", materialID).Order("is_preferred DESC, id DESC").Find(&list).Error
	return list, err
}

func (r *SupplierMaterialRepository) Create(ctx context.Context, item *model.SupplierMaterial) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *SupplierMaterialRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.SupplierMaterial{}).Where("id = ?", id).Updates(updates).Error
}

func (r *SupplierMaterialRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.SupplierMaterial{}, id).Error
}

func (r *SupplierMaterialRepository) ClearPreferredBySupplier(ctx context.Context, supplierID int64) error {
	return r.db.WithContext(ctx).Model(&model.SupplierMaterial{}).Where("supplier_id = ?", supplierID).Update("is_preferred", 0).Error
}
