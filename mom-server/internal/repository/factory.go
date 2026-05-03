package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// FactoryRepository 工厂仓储
type FactoryRepository struct {
	db *gorm.DB
}

func NewFactoryRepository(db *gorm.DB) *FactoryRepository {
	return &FactoryRepository{db: db}
}

func (r *FactoryRepository) Create(ctx context.Context, m *model.MdmFactory) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *FactoryRepository) GetByID(ctx context.Context, id int64) (*model.MdmFactory, error) {
	var m model.MdmFactory
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	return &m, err
}

func (r *FactoryRepository) GetByCode(ctx context.Context, tenantID int64, code string) (*model.MdmFactory, error) {
	var m model.MdmFactory
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND factory_code = ?", tenantID, code).First(&m).Error
	return &m, err
}

func (r *FactoryRepository) Page(ctx context.Context, tenantID int64, req *model.FactoryQuery) ([]model.MdmFactory, int64, error) {
	var list []model.MdmFactory
	query := r.db.WithContext(ctx).Model(&model.MdmFactory{}).Where("tenant_id = ?", tenantID)
	if req.FactoryCode != "" {
		query = query.Where("factory_code LIKE ?", "%"+req.FactoryCode+"%")
	}
	if req.FactoryName != "" {
		query = query.Where("factory_name LIKE ?", "%"+req.FactoryName+"%")
	}
	if req.Status > 0 {
		query = query.Where("status = ?", req.Status)
	}
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	err = query.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *FactoryRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Factory{}).Where("id = ?", id).Updates(updates).Error
}

func (r *FactoryRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Factory{}, id).Error
}

func (r *FactoryRepository) ListByTenant(ctx context.Context, tenantID int64) ([]model.MdmFactory, error) {
	var list []model.MdmFactory
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND status = 1", tenantID).Order("is_default DESC, id ASC").Find(&list).Error
	return list, err
}

// SetDefault 设置默认工厂
func (r *FactoryRepository) SetDefault(ctx context.Context, tenantID int64, factoryID int64) error {
	// 先清除所有默认标记
	if err := r.db.WithContext(ctx).Model(&model.MdmFactory{}).Where("tenant_id = ?", tenantID).Update("is_default", 0).Error; err != nil {
		return err
	}
	// 设置新默认
	return r.db.WithContext(ctx).Model(&model.MdmFactory{}).Where("id = ?", factoryID).Update("is_default", 1).Error
}

// TenantFactory 当前工厂设置
func (r *FactoryRepository) SetCurrentFactory(ctx context.Context, tenantID, userID, factoryID int64) error {
	// 清除当前设置
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND user_id = ?", tenantID, userID).Delete(&model.TenantFactory{}).Error
	if err != nil {
		return err
	}
	// 创建新当前工厂记录
	return r.db.WithContext(ctx).Create(&model.TenantFactory{
		TenantID:  tenantID,
		UserID:   userID,
		FactoryID: factoryID,
		IsCurrent: 1,
	}).Error
}

func (r *FactoryRepository) GetCurrentFactory(ctx context.Context, tenantID, userID int64) (*model.MdmFactory, error) {
	var tf model.TenantFactory
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND user_id = ? AND is_current = 1", tenantID, userID).First(&tf).Error
	if err != nil {
		// 没有设置，取默认工厂
		var f model.MdmFactory
		err := r.db.WithContext(ctx).Where("tenant_id = ? AND is_default = 1", tenantID).First(&f).Error
		if err == nil {
			return &f, nil
		}
		// 没有默认工厂，取第一个
		err = r.db.WithContext(ctx).Where("tenant_id = ? AND status = 1", tenantID).Order("id ASC").First(&f).Error
		return &f, err
	}
	return r.GetByID(ctx, tf.FactoryID)
}