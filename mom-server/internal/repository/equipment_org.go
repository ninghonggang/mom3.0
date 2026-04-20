package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// EquipmentOrgRepository 设备组织仓储
type EquipmentOrgRepository struct {
	db *gorm.DB
}

// NewEquipmentOrgRepository 创建设备组织仓储
func NewEquipmentOrgRepository(db *gorm.DB) *EquipmentOrgRepository {
	return &EquipmentOrgRepository{db: db}
}

// List 获取设备组织列表
func (r *EquipmentOrgRepository) List(ctx context.Context, tenantID int64, query *model.EquipmentOrgQuery) ([]model.EquipmentOrg, int64, error) {
	var list []model.EquipmentOrg
	var total int64

	q := r.db.WithContext(ctx).Model(&model.EquipmentOrg{})
	if tenantID > 0 {
		q = q.Where("tenant_id = ?", tenantID)
	}
	if query != nil {
		if query.FactoryID > 0 {
			q = q.Where("factory_id = ?", query.FactoryID)
		}
		if query.WorkshopID > 0 {
			q = q.Where("workshop_id = ?", query.WorkshopID)
		}
		if query.Status > 0 {
			q = q.Where("status = ?", query.Status)
		}
	}

	err := q.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = q.Order("factory_id, workshop_id, line_id").Find(&list).Error
	return list, total, err
}

// GetByID 根据ID获取设备组织
func (r *EquipmentOrgRepository) GetByID(ctx context.Context, id uint) (*model.EquipmentOrg, error) {
	var org model.EquipmentOrg
	err := r.db.WithContext(ctx).First(&org, id).Error
	return &org, err
}

// CreateSync 创建或同步设备组织关系
func (r *EquipmentOrgRepository) CreateSync(ctx context.Context, org *model.EquipmentOrg) error {
	return r.db.WithContext(ctx).Where(
		"tenant_id = ? AND factory_id = ? AND workshop_id = ? AND line_id = ?",
		org.TenantID, org.FactoryID, org.WorkshopID, org.LineID,
	).Assign(model.EquipmentOrg{
		FactoryCode:  org.FactoryCode,
		FactoryName:  org.FactoryName,
		WorkshopCode: org.WorkshopCode,
		WorkshopName: org.WorkshopName,
		LineCode:     org.LineCode,
		LineName:     org.LineName,
		Status:       org.Status,
	}).FirstOrCreate(org).Error
}

// Update 更新设备组织
func (r *EquipmentOrgRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.EquipmentOrg{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除设备组织
func (r *EquipmentOrgRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.EquipmentOrg{}, id).Error
}

// DeleteByLine 根据产线ID删除设备组织关系
func (r *EquipmentOrgRepository) DeleteByLine(ctx context.Context, tenantID, lineID int64) error {
	return r.db.WithContext(ctx).Where("tenant_id = ? AND line_id = ?", tenantID, lineID).Delete(&model.EquipmentOrg{}).Error
}

// GetFactoryList 获取厂区列表
func (r *EquipmentOrgRepository) GetFactoryList(ctx context.Context, tenantID int64) ([]model.Factory, error) {
	var list []model.Factory
	err := r.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Order("id").Find(&list).Error
	return list, err
}

// CreateFactory 创建厂区
func (r *EquipmentOrgRepository) CreateFactory(ctx context.Context, factory *model.Factory) error {
	return r.db.WithContext(ctx).Create(factory).Error
}

// GetFactoryByID 根据ID获取厂区
func (r *EquipmentOrgRepository) GetFactoryByID(ctx context.Context, id uint) (*model.Factory, error) {
	var factory model.Factory
	err := r.db.WithContext(ctx).First(&factory, id).Error
	return &factory, err
}

// GetFactoryByCode 根据编码获取厂区
func (r *EquipmentOrgRepository) GetFactoryByCode(ctx context.Context, tenantID int64, code string) (*model.Factory, error) {
	var factory model.Factory
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND factory_code = ?", tenantID, code).First(&factory).Error
	return &factory, err
}

// UpdateFactory 更新厂区
func (r *EquipmentOrgRepository) UpdateFactory(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Factory{}).Where("id = ?", id).Updates(updates).Error
}

// DeleteFactory 删除厂区
func (r *EquipmentOrgRepository) DeleteFactory(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Factory{}, id).Error
}
