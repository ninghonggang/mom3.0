package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type InspectionTemplateRepository struct {
	db *gorm.DB
}

func NewInspectionTemplateRepository(db *gorm.DB) *InspectionTemplateRepository {
	return &InspectionTemplateRepository{db: db}
}

func (r *InspectionTemplateRepository) Create(ctx context.Context, m *model.InspectionTemplate) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *InspectionTemplateRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.InspectionTemplate{}).Where("id = ?", id).Updates(updates).Error
}

func (r *InspectionTemplateRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("template_id = ?", id).Delete(&model.InspectionItem{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.InspectionTemplate{}, id).Error
	})
}

func (r *InspectionTemplateRepository) GetByID(ctx context.Context, id uint) (*model.InspectionTemplate, error) {
	var m model.InspectionTemplate
	err := r.db.WithContext(ctx).First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *InspectionTemplateRepository) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.InspectionTemplate, int64, error) {
	var list []model.InspectionTemplate
	var total int64
	query := r.db.WithContext(ctx).Model(&model.InspectionTemplate{})
	for k, v := range filters {
		if v != "" {
			query = query.Where(k+" = ?", v)
		}
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

type InspectionItemRepository struct {
	db *gorm.DB
}

func NewInspectionItemRepository(db *gorm.DB) *InspectionItemRepository {
	return &InspectionItemRepository{db: db}
}

func (r *InspectionItemRepository) Create(ctx context.Context, m *model.InspectionItem) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *InspectionItemRepository) CreateBatch(ctx context.Context, items []model.InspectionItem) error {
	return r.db.WithContext(ctx).Create(&items).Error
}

func (r *InspectionItemRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.InspectionItem{}).Where("id = ?", id).Updates(updates).Error
}

func (r *InspectionItemRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.InspectionItem{}, id).Error
}

func (r *InspectionItemRepository) GetByID(ctx context.Context, id uint) (*model.InspectionItem, error) {
	var m model.InspectionItem
	err := r.db.WithContext(ctx).First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *InspectionItemRepository) ListByTemplate(ctx context.Context, templateID uint) ([]model.InspectionItem, error) {
	var list []model.InspectionItem
	err := r.db.WithContext(ctx).Where("template_id = ?", templateID).Order("sort_order").Find(&list).Error
	return list, err
}

type InspectionPlanRepository struct {
	db *gorm.DB
}

func NewInspectionPlanRepository(db *gorm.DB) *InspectionPlanRepository {
	return &InspectionPlanRepository{db: db}
}

func (r *InspectionPlanRepository) Create(ctx context.Context, m *model.InspectionPlan) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *InspectionPlanRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.InspectionPlan{}).Where("id = ?", id).Updates(updates).Error
}

func (r *InspectionPlanRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.InspectionPlan{}, id).Error
}

func (r *InspectionPlanRepository) GetByID(ctx context.Context, id uint) (*model.InspectionPlan, error) {
	var m model.InspectionPlan
	err := r.db.WithContext(ctx).First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *InspectionPlanRepository) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.InspectionPlan, int64, error) {
	var list []model.InspectionPlan
	var total int64
	query := r.db.WithContext(ctx).Model(&model.InspectionPlan{})
	for k, v := range filters {
		if v != "" {
			query = query.Where(k+" = ?", v)
		}
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

type InspectionRecordRepository struct {
	db *gorm.DB
}

func NewInspectionRecordRepository(db *gorm.DB) *InspectionRecordRepository {
	return &InspectionRecordRepository{db: db}
}

func (r *InspectionRecordRepository) Create(ctx context.Context, m *model.InspectionRecord) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *InspectionRecordRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.InspectionRecord{}).Where("id = ?", id).Updates(updates).Error
}

func (r *InspectionRecordRepository) GetByID(ctx context.Context, id uint) (*model.InspectionRecord, error) {
	var m model.InspectionRecord
	err := r.db.WithContext(ctx).First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *InspectionRecordRepository) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.InspectionRecord, int64, error) {
	var list []model.InspectionRecord
	var total int64
	query := r.db.WithContext(ctx).Model(&model.InspectionRecord{})
	for k, v := range filters {
		if v != "" {
			query = query.Where(k+" = ?", v)
		}
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

type InspectionResultRepository struct {
	db *gorm.DB
}

func NewInspectionResultRepository(db *gorm.DB) *InspectionResultRepository {
	return &InspectionResultRepository{db: db}
}

func (r *InspectionResultRepository) Create(ctx context.Context, m *model.InspectionResult) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *InspectionResultRepository) CreateBatch(ctx context.Context, items []model.InspectionResult) error {
	return r.db.WithContext(ctx).Create(&items).Error
}

func (r *InspectionResultRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.InspectionResult{}).Where("id = ?", id).Updates(updates).Error
}

func (r *InspectionResultRepository) ListByRecord(ctx context.Context, recordID uint) ([]model.InspectionResult, error) {
	var list []model.InspectionResult
	err := r.db.WithContext(ctx).Where("record_id = ?", recordID).Find(&list).Error
	return list, err
}

type InspectionDefectRepository struct {
	db *gorm.DB
}

func NewInspectionDefectRepository(db *gorm.DB) *InspectionDefectRepository {
	return &InspectionDefectRepository{db: db}
}

func (r *InspectionDefectRepository) Create(ctx context.Context, m *model.InspectionDefect) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *InspectionDefectRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.InspectionDefect{}).Where("id = ?", id).Updates(updates).Error
}

func (r *InspectionDefectRepository) GetByID(ctx context.Context, id uint) (*model.InspectionDefect, error) {
	var m model.InspectionDefect
	err := r.db.WithContext(ctx).First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *InspectionDefectRepository) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.InspectionDefect, int64, error) {
	var list []model.InspectionDefect
	var total int64
	query := r.db.WithContext(ctx).Model(&model.InspectionDefect{})
	for k, v := range filters {
		if v != "" {
			query = query.Where(k+" = ?", v)
		}
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
