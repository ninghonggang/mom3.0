package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type InspectionFeatureRepository struct {
	db *gorm.DB
}

func NewInspectionFeatureRepository(db *gorm.DB) *InspectionFeatureRepository {
	return &InspectionFeatureRepository{db: db}
}

// List 查询检验特性列表
func (r *InspectionFeatureRepository) List(ctx context.Context, tenantID uint64, query map[string]interface{}) ([]model.InspectionFeature, int64, error) {
	var list []model.InspectionFeature
	var total int64

	db := r.db.WithContext(ctx).Model(&model.InspectionFeature{}).Where("tenant_id = ?", tenantID)

	if productID, ok := query["product_id"].(uint64); ok && productID > 0 {
		db = db.Where("product_id = ?", productID)
	}
	if inspectionType, ok := query["inspection_type"].(string); ok && inspectionType != "" {
		db = db.Where("inspection_type = ?", inspectionType)
	}
	if featureType, ok := query["feature_type"].(string); ok && featureType != "" {
		db = db.Where("feature_type = ?", featureType)
	}
	if status, ok := query["status"].(string); ok && status != "" {
		db = db.Where("status = ?", status)
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Order("id DESC").Find(&list).Error
	return list, total, err
}

// GetByID 根据ID获取检验特性
func (r *InspectionFeatureRepository) GetByID(ctx context.Context, id uint64) (*model.InspectionFeature, error) {
	var feature model.InspectionFeature
	err := r.db.WithContext(ctx).First(&feature, id).Error
	if err != nil {
		return nil, err
	}
	return &feature, nil
}

// Create 创建检验特性
func (r *InspectionFeatureRepository) Create(ctx context.Context, feature *model.InspectionFeature) error {
	return r.db.WithContext(ctx).Create(feature).Error
}

// Update 更新检验特性
func (r *InspectionFeatureRepository) Update(ctx context.Context, feature *model.InspectionFeature) error {
	return r.db.WithContext(ctx).Save(feature).Error
}

// Delete 删除检验特性
func (r *InspectionFeatureRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.InspectionFeature{}, id).Error
}

// GetByProductID 获取产品的所有检验特性
func (r *InspectionFeatureRepository) GetByProductID(ctx context.Context, tenantID uint64, productID uint64) ([]model.InspectionFeature, error) {
	var list []model.InspectionFeature
	err := r.db.WithContext(ctx).
		Where("tenant_id = ? AND product_id = ? AND status = ?", tenantID, productID, "ACTIVE").
		Order("id ASC").
		Find(&list).Error
	return list, err
}

// GetByFeatureCode 根据特性编码获取
func (r *InspectionFeatureRepository) GetByFeatureCode(ctx context.Context, tenantID uint64, featureCode string) (*model.InspectionFeature, error) {
	var feature model.InspectionFeature
	err := r.db.WithContext(ctx).
		Where("tenant_id = ? AND feature_code = ?", tenantID, featureCode).
		First(&feature).Error
	if err != nil {
		return nil, err
	}
	return &feature, nil
}

// BatchCreate 批量创建
func (r *InspectionFeatureRepository) BatchCreate(ctx context.Context, features []model.InspectionFeature) error {
	return r.db.WithContext(ctx).CreateInBatches(features, 100).Error
}