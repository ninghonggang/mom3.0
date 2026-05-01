package repository

import (
	"context"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

// AQLLevelRepository AQL级别仓储
type AQLLevelRepository struct {
	db *gorm.DB
}

func NewAQLLevelRepository(db *gorm.DB) *AQLLevelRepository {
	return &AQLLevelRepository{db: db}
}

func (r *AQLLevelRepository) List(ctx context.Context, tenantID int64) ([]model.AQLLevel, int64, error) {
	var list []model.AQLLevel
	var total int64
	q := r.db.WithContext(ctx).Model(&model.AQLLevel{}).Where("tenant_id = ?", tenantID)
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := q.Order(`"order" ASC, id ASC`).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *AQLLevelRepository) GetByID(ctx context.Context, id uint) (*model.AQLLevel, error) {
	var level model.AQLLevel
	err := r.db.WithContext(ctx).First(&level, id).Error
	return &level, err
}

func (r *AQLLevelRepository) Create(ctx context.Context, level *model.AQLLevel) error {
	return r.db.WithContext(ctx).Create(level).Error
}

func (r *AQLLevelRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.AQLLevel{}).Where("id = ?", id).Updates(updates).Error
}

func (r *AQLLevelRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.AQLLevel{}, id).Error
}

// AQLTableRowRepository AQL表行仓储
type AQLTableRowRepository struct {
	db *gorm.DB
}

func NewAQLTableRowRepository(db *gorm.DB) *AQLTableRowRepository {
	return &AQLTableRowRepository{db: db}
}

func (r *AQLTableRowRepository) ListByLevelID(ctx context.Context, aqlLevelID int64) ([]model.AQLTableRow, error) {
	var list []model.AQLTableRow
	err := r.db.WithContext(ctx).Where("aql_level_id = ?", aqlLevelID).Order("batch_min ASC").Find(&list).Error
	return list, err
}

func (r *AQLTableRowRepository) GetByBatchAndAQL(ctx context.Context, tenantID int64, batchSize int, aqlValue string) (*model.AQLTableRow, error) {
	var row model.AQLTableRow
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND aql_value = ? AND batch_min <= ? AND (batch_max = 0 OR batch_max >= ?)", tenantID, aqlValue, batchSize, batchSize).First(&row).Error
	return &row, err
}

func (r *AQLTableRowRepository) Create(ctx context.Context, row *model.AQLTableRow) error {
	return r.db.WithContext(ctx).Create(row).Error
}

func (r *AQLTableRowRepository) DeleteByLevelID(ctx context.Context, aqlLevelID uint) error {
	return r.db.WithContext(ctx).Where("aql_level_id = ?", aqlLevelID).Delete(&model.AQLTableRow{}).Error
}

// SamplingPlanRepository 抽样方案仓储
type SamplingPlanRepository struct {
	db *gorm.DB
}

func NewSamplingPlanRepository(db *gorm.DB) *SamplingPlanRepository {
	return &SamplingPlanRepository{db: db}
}

func (r *SamplingPlanRepository) List(ctx context.Context, tenantID int64, query string) ([]model.SamplingPlan, int64, error) {
	var list []model.SamplingPlan
	var total int64
	q := r.db.WithContext(ctx).Model(&model.SamplingPlan{}).Where("tenant_id = ?", tenantID)
	if query != "" {
		q = q.Where("code LIKE ? OR name LIKE ?", "%"+query+"%", "%"+query+"%")
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := q.Order("id DESC").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *SamplingPlanRepository) GetByID(ctx context.Context, id uint) (*model.SamplingPlan, error) {
	var plan model.SamplingPlan
	err := r.db.WithContext(ctx).First(&plan, id).Error
	return &plan, err
}

func (r *SamplingPlanRepository) GetByCode(ctx context.Context, tenantID int64, code string) (*model.SamplingPlan, error) {
	var plan model.SamplingPlan
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND code = ?", tenantID, code).First(&plan).Error
	return &plan, err
}

func (r *SamplingPlanRepository) Create(ctx context.Context, plan *model.SamplingPlan) error {
	return r.db.WithContext(ctx).Create(plan).Error
}

func (r *SamplingPlanRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.SamplingPlan{}).Where("id = ?", id).Updates(updates).Error
}

func (r *SamplingPlanRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.SamplingPlan{}, id).Error
}
