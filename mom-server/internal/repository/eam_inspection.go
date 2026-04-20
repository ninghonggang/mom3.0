package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// EAMInspectionPlanRepository 巡检计划仓储
type EAMInspectionPlanRepository struct {
	db *gorm.DB
}

func NewEAMInspectionPlanRepository(db *gorm.DB) *EAMInspectionPlanRepository {
	return &EAMInspectionPlanRepository{db: db}
}

func (r *EAMInspectionPlanRepository) Create(ctx context.Context, m *model.EAMInspectionPlan) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *EAMInspectionPlanRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.EAMInspectionPlan{}).Where("id = ?", id).Updates(updates).Error
}

func (r *EAMInspectionPlanRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 先删除关联的巡检项目
		if err := tx.Where("plan_id = ?", id).Delete(&model.EAMInspectionItem{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.EAMInspectionPlan{}, id).Error
	})
}

func (r *EAMInspectionPlanRepository) GetByID(ctx context.Context, id int64) (*model.EAMInspectionPlan, error) {
	var m model.EAMInspectionPlan
	err := r.db.WithContext(ctx).First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *EAMInspectionPlanRepository) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.EAMInspectionPlan, int64, error) {
	var list []model.EAMInspectionPlan
	var total int64
	query := r.db.WithContext(ctx).Model(&model.EAMInspectionPlan{})
	if planNo, ok := filters["plan_no"].(string); ok && planNo != "" {
		query = query.Where("plan_no LIKE ?", "%"+planNo+"%")
	}
	if planName, ok := filters["plan_name"].(string); ok && planName != "" {
		query = query.Where("plan_name LIKE ?", "%"+planName+"%")
	}
	if status, ok := filters["status"].(string); ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if limit > 0 {
		query = query.Offset(offset).Limit(limit).Order("id desc")
	}
	if err := query.Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// EAMInspectionItemRepository 巡检项目仓储
type EAMInspectionItemRepository struct {
	db *gorm.DB
}

func NewEAMInspectionItemRepository(db *gorm.DB) *EAMInspectionItemRepository {
	return &EAMInspectionItemRepository{db: db}
}

func (r *EAMInspectionItemRepository) CreateBatch(ctx context.Context, items []model.EAMInspectionItem) error {
	return r.db.WithContext(ctx).Create(&items).Error
}

func (r *EAMInspectionItemRepository) UpdateByPlanID(ctx context.Context, planID int64, items []model.EAMInspectionItem) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 删除原有项目
		if err := tx.Where("plan_id = ?", planID).Delete(&model.EAMInspectionItem{}).Error; err != nil {
			return err
		}
		// 插入新项目
		if len(items) > 0 {
			if err := tx.Create(&items).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *EAMInspectionItemRepository) ListByPlanID(ctx context.Context, planID int64) ([]model.EAMInspectionItem, error) {
	var list []model.EAMInspectionItem
	err := r.db.WithContext(ctx).Where("plan_id = ?", planID).Order("sort_order").Find(&list).Error
	return list, err
}

func (r *EAMInspectionItemRepository) DeleteByPlanID(ctx context.Context, planID int64) error {
	return r.db.WithContext(ctx).Where("plan_id = ?", planID).Delete(&model.EAMInspectionItem{}).Error
}

// EAMInspectionSchemeRepository 巡检方案仓储
type EAMInspectionSchemeRepository struct {
	db *gorm.DB
}

func NewEAMInspectionSchemeRepository(db *gorm.DB) *EAMInspectionSchemeRepository {
	return &EAMInspectionSchemeRepository{db: db}
}

func (r *EAMInspectionSchemeRepository) Create(ctx context.Context, m *model.EAMInspectionScheme) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *EAMInspectionSchemeRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.EAMInspectionScheme{}).Where("id = ?", id).Updates(updates).Error
}

func (r *EAMInspectionSchemeRepository) GetByID(ctx context.Context, id int64) (*model.EAMInspectionScheme, error) {
	var m model.EAMInspectionScheme
	err := r.db.WithContext(ctx).First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *EAMInspectionSchemeRepository) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.EAMInspectionScheme, int64, error) {
	var list []model.EAMInspectionScheme
	var total int64
	query := r.db.WithContext(ctx).Model(&model.EAMInspectionScheme{})
	if planNo, ok := filters["plan_no"].(string); ok && planNo != "" {
		query = query.Where("plan_no LIKE ?", "%"+planNo+"%")
	}
	if status, ok := filters["status"].(string); ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if planID, ok := filters["plan_id"].(int64); ok && planID > 0 {
		query = query.Where("plan_id = ?", planID)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if limit > 0 {
		query = query.Offset(offset).Limit(limit).Order("id desc")
	}
	if err := query.Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// EAMInspectionResultRepository 巡检结果仓储
type EAMInspectionResultRepository struct {
	db *gorm.DB
}

func NewEAMInspectionResultRepository(db *gorm.DB) *EAMInspectionResultRepository {
	return &EAMInspectionResultRepository{db: db}
}

func (r *EAMInspectionResultRepository) CreateBatch(ctx context.Context, items []model.EAMInspectionResult) error {
	return r.db.WithContext(ctx).Create(&items).Error
}

func (r *EAMInspectionResultRepository) ListBySchemeID(ctx context.Context, schemeID int64) ([]model.EAMInspectionResult, error) {
	var list []model.EAMInspectionResult
	err := r.db.WithContext(ctx).Where("scheme_id = ?", schemeID).Find(&list).Error
	return list, err
}

func (r *EAMInspectionResultRepository) DeleteBySchemeID(ctx context.Context, schemeID int64) error {
	return r.db.WithContext(ctx).Where("scheme_id = ?", schemeID).Delete(&model.EAMInspectionResult{}).Error
}