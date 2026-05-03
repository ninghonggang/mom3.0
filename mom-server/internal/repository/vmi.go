package repository

import (
	"context"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

type VmiRepository struct {
	db *gorm.DB
}

func NewVmiRepository(db *gorm.DB) *VmiRepository {
	return &VmiRepository{db: db}
}

// VmiVendor
func (r *VmiRepository) ListVendors(ctx context.Context, tenantID int64, query *model.VmiVendorQuery) ([]model.VmiVendor, int64, error) {
	var list []model.VmiVendor
	var total int64

	db := r.db.WithContext(ctx).Model(&model.VmiVendor{}).Where("tenant_id = ?", tenantID)

	if query.VendorCode != "" {
		db = db.Where("vendor_code LIKE ?", "%"+query.VendorCode+"%")
	}
	if query.VendorName != "" {
		db = db.Where("vendor_name LIKE ?", "%"+query.VendorName+"%")
	}
	if query.WarehouseID > 0 {
		db = db.Where("warehouse_id = ?", query.WarehouseID)
	}
	if query.IsActive > 0 {
		db = db.Where("is_active = ?", query.IsActive)
	}
	if query.Status > 0 {
		db = db.Where("status = ?", query.Status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 20
	}
	offset := (query.Page - 1) * query.PageSize

	if err := db.Order("id DESC").Offset(offset).Limit(query.PageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *VmiRepository) GetVendorByID(ctx context.Context, id int64) (*model.VmiVendor, error) {
	var vendor model.VmiVendor
	if err := r.db.WithContext(ctx).First(&vendor, id).Error; err != nil {
		return nil, err
	}
	return &vendor, nil
}

func (r *VmiRepository) CreateVendor(ctx context.Context, vendor *model.VmiVendor) error {
	return r.db.WithContext(ctx).Create(vendor).Error
}

func (r *VmiRepository) UpdateVendor(ctx context.Context, vendor *model.VmiVendor) error {
	return r.db.WithContext(ctx).Where("id = ?", vendor.ID).Updates(vendor).Error
}

func (r *VmiRepository) DeleteVendor(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.VmiVendor{}, id).Error
}

// VmiMaterial
func (r *VmiRepository) ListMaterials(ctx context.Context, tenantID int64, query *model.VmiMaterialQuery) ([]model.VmiMaterial, int64, error) {
	var list []model.VmiMaterial
	var total int64

	db := r.db.WithContext(ctx).Model(&model.VmiMaterial{}).Where("tenant_id = ?", tenantID)

	if query.VendorID > 0 {
		db = db.Where("vendor_id = ?", query.VendorID)
	}
	if query.MaterialCode != "" {
		db = db.Where("material_code LIKE ?", "%"+query.MaterialCode+"%")
	}
	if query.MaterialName != "" {
		db = db.Where("material_name LIKE ?", "%"+query.MaterialName+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 20
	}
	offset := (query.Page - 1) * query.PageSize

	if err := db.Order("id DESC").Offset(offset).Limit(query.PageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *VmiRepository) GetMaterialByID(ctx context.Context, id int64) (*model.VmiMaterial, error) {
	var material model.VmiMaterial
	if err := r.db.WithContext(ctx).First(&material, id).Error; err != nil {
		return nil, err
	}
	return &material, nil
}

func (r *VmiRepository) CreateMaterial(ctx context.Context, material *model.VmiMaterial) error {
	return r.db.WithContext(ctx).Create(material).Error
}

func (r *VmiRepository) UpdateMaterial(ctx context.Context, material *model.VmiMaterial) error {
	return r.db.WithContext(ctx).Where("id = ?", material.ID).Updates(material).Error
}

func (r *VmiRepository) UpdateMaterialStock(ctx context.Context, id int64, currentStock, availableStock float64) error {
	return r.db.WithContext(ctx).Model(&model.VmiMaterial{}).Where("id = ?", id).Updates(map[string]interface{}{
		"current_stock": currentStock,
		"available_stock": availableStock,
	}).Error
}

func (r *VmiRepository) DeleteMaterial(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.VmiMaterial{}, id).Error
}

// VmiTransaction
func (r *VmiRepository) ListTransactions(ctx context.Context, tenantID int64, query *model.VmiTransactionQuery) ([]model.VmiTransaction, int64, error) {
	var list []model.VmiTransaction
	var total int64

	db := r.db.WithContext(ctx).Model(&model.VmiTransaction{}).Where("tenant_id = ?", tenantID)

	if query.VendorID > 0 {
		db = db.Where("vendor_id = ?", query.VendorID)
	}
	if query.MaterialID > 0 {
		db = db.Where("material_id = ?", query.MaterialID)
	}
	if query.TransactionType > 0 {
		db = db.Where("transaction_type = ?", query.TransactionType)
	}
	if query.StartDate != "" {
		db = db.Where("created_at >= ?", query.StartDate)
	}
	if query.EndDate != "" {
		db = db.Where("created_at <= ?", query.EndDate)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 20
	}
	offset := (query.Page - 1) * query.PageSize

	if err := db.Order("id DESC").Offset(offset).Limit(query.PageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *VmiRepository) CreateTransaction(ctx context.Context, tx *model.VmiTransaction) error {
	return r.db.WithContext(ctx).Create(tx).Error
}