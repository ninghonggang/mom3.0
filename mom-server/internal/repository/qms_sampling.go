package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// QMSSamplingPlanRepository 抽样方案仓储
type QMSSamplingPlanRepository struct {
	db *gorm.DB
}

func NewQMSSamplingPlanRepository(db *gorm.DB) *QMSSamplingPlanRepository {
	return &QMSSamplingPlanRepository{db: db}
}

func (r *QMSSamplingPlanRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.QMSSamplingPlan, int64, error) {
	var list []model.QMSSamplingPlan
	var total int64

	q := r.db.WithContext(ctx).Model(&model.QMSSamplingPlan{}).Where("tenant_id = ?", tenantID)

	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}
	if planCode, ok := query["plan_code"]; ok && planCode != "" {
		q = q.Where("plan_code LIKE ?", "%"+planCode.(string)+"%")
	}
	if planName, ok := query["plan_name"]; ok && planName != "" {
		q = q.Where("plan_name LIKE ?", "%"+planName.(string)+"%")
	}

	q.Count(&total)

	page := 1
	limit := 20
	if p, ok := query["page"].(int); ok && p > 0 {
		page = p
	}
	if l, ok := query["limit"].(int); ok && l > 0 {
		limit = l
	}
	q = q.Offset((page - 1) * limit).Limit(limit).Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *QMSSamplingPlanRepository) GetByID(ctx context.Context, id uint) (*model.QMSSamplingPlan, error) {
	var plan model.QMSSamplingPlan
	err := r.db.WithContext(ctx).First(&plan, id).Error
	return &plan, err
}

func (r *QMSSamplingPlanRepository) GetByPlanCode(ctx context.Context, tenantID int64, planCode string) (*model.QMSSamplingPlan, error) {
	var plan model.QMSSamplingPlan
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND plan_code = ?", tenantID, planCode).First(&plan).Error
	return &plan, err
}

func (r *QMSSamplingPlanRepository) Create(ctx context.Context, plan *model.QMSSamplingPlan) error {
	return r.db.WithContext(ctx).Create(plan).Error
}

func (r *QMSSamplingPlanRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.QMSSamplingPlan{}).Where("id = ?", id).Updates(updates).Error
}

func (r *QMSSamplingPlanRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.QMSSamplingPlan{}, id).Error
}

// QMSSamplingRuleRepository 抽样规则仓储
type QMSSamplingRuleRepository struct {
	db *gorm.DB
}

func NewQMSSamplingRuleRepository(db *gorm.DB) *QMSSamplingRuleRepository {
	return &QMSSamplingRuleRepository{db: db}
}

func (r *QMSSamplingRuleRepository) ListByPlanID(ctx context.Context, planID int64) ([]model.QMSSamplingRule, error) {
	var rules []model.QMSSamplingRule
	err := r.db.WithContext(ctx).Where("plan_id = ?", planID).Order("batch_qty_from ASC").Find(&rules).Error
	return rules, err
}

func (r *QMSSamplingRuleRepository) DeleteByPlanID(ctx context.Context, planID int64) error {
	return r.db.WithContext(ctx).Where("plan_id = ?", planID).Delete(&model.QMSSamplingRule{}).Error
}

func (r *QMSSamplingRuleRepository) CreateBatch(ctx context.Context, rules []model.QMSSamplingRule) error {
	return r.db.WithContext(ctx).Create(&rules).Error
}

func (r *QMSSamplingRuleRepository) GetRuleForBatchQty(ctx context.Context, planID int64, batchQty float64) (*model.QMSSamplingRule, error) {
	var rule model.QMSSamplingRule
	err := r.db.WithContext(ctx).Where("plan_id = ? AND batch_qty_from <= ? AND batch_qty_to >= ?", planID, batchQty, batchQty).First(&rule).Error
	return &rule, err
}

// QMSSamplingRecordRepository 抽样记录仓储
type QMSSamplingRecordRepository struct {
	db *gorm.DB
}

func NewQMSSamplingRecordRepository(db *gorm.DB) *QMSSamplingRecordRepository {
	return &QMSSamplingRecordRepository{db: db}
}

func (r *QMSSamplingRecordRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.QMSSamplingRecord, int64, error) {
	var list []model.QMSSamplingRecord
	var total int64

	q := r.db.WithContext(ctx).Model(&model.QMSSamplingRecord{}).Where("tenant_id = ?", tenantID)

	if planID, ok := query["plan_id"]; ok && planID.(int64) > 0 {
		q = q.Where("plan_id = ?", planID)
	}
	if inspectionID, ok := query["inspection_id"]; ok && inspectionID.(int64) > 0 {
		q = q.Where("inspection_id = ?", inspectionID)
	}

	q.Count(&total)

	page := 1
	limit := 20
	if p, ok := query["page"].(int); ok && p > 0 {
		page = p
	}
	if l, ok := query["limit"].(int); ok && l > 0 {
		limit = l
	}
	q = q.Offset((page - 1) * limit).Limit(limit).Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *QMSSamplingRecordRepository) Create(ctx context.Context, record *model.QMSSamplingRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *QMSSamplingRecordRepository) GetByID(ctx context.Context, id uint) (*model.QMSSamplingRecord, error) {
	var record model.QMSSamplingRecord
	err := r.db.WithContext(ctx).First(&record, id).Error
	return &record, err
}
