package repository

import (
	"context"
	"mom-server/internal/model"
	"gorm.io/gorm"
)

type TenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) *TenantRepository {
	return &TenantRepository{db: db}
}

func (r *TenantRepository) Create(ctx context.Context, tenant *model.Tenant) error {
	return r.db.WithContext(ctx).Create(tenant).Error
}

func (r *TenantRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Tenant{}).Where("id = ?", id).Updates(updates).Error
}

func (r *TenantRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Tenant{}, id).Error
}

func (r *TenantRepository) FindByID(ctx context.Context, id int64) (*model.Tenant, error) {
	var tenant model.Tenant
	err := r.db.WithContext(ctx).First(&tenant, id).Error
	return &tenant, err
}

func (r *TenantRepository) FindByKey(ctx context.Context, key string) (*model.Tenant, error) {
	var tenant model.Tenant
	err := r.db.WithContext(ctx).Where("tenant_key = ?", key).First(&tenant).Error
	return &tenant, err
}

func (r *TenantRepository) FindByPage(ctx context.Context, tenantName, tenantKey, status string, page, pageSize int) ([]model.Tenant, int64, error) {
	var tenants []model.Tenant
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Tenant{})
	if tenantName != "" {
		query = query.Where("tenant_name LIKE ?", "%"+tenantName+"%")
	}
	if tenantKey != "" {
		query = query.Where("tenant_key LIKE ?", "%"+tenantKey+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("id DESC").Find(&tenants).Error
	return tenants, total, err
}
