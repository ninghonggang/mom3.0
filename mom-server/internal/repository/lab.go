package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// LabSampleRepository 检测样品仓储
type LabSampleRepository struct {
	db *gorm.DB
}

func NewLabSampleRepository(db *gorm.DB) *LabSampleRepository {
	return &LabSampleRepository{db: db}
}

func (r *LabSampleRepository) List(ctx context.Context, query *model.LabSampleQuery) ([]model.LabSample, int64, error) {
	var list []model.LabSample
	var total int64

	db := r.db.WithContext(ctx).Model(&model.LabSample{})
	if query.SampleCode != "" {
		db = db.Where("sample_code LIKE ?", "%"+query.SampleCode+"%")
	}
	if query.SampleName != "" {
		db = db.Where("sample_name LIKE ?", "%"+query.SampleName+"%")
	}
	if query.InspectionType != "" {
		db = db.Where("inspection_type = ?", query.InspectionType)
	}
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (query.Page - 1) * query.PageSize
	err = db.Order("id DESC").Offset(offset).Limit(query.PageSize).Find(&list).Error
	return list, total, err
}

func (r *LabSampleRepository) GetByID(ctx context.Context, id uint64) (*model.LabSample, error) {
	var sample model.LabSample
	err := r.db.WithContext(ctx).First(&sample, id).Error
	if err != nil {
		return nil, err
	}
	return &sample, nil
}

func (r *LabSampleRepository) Create(ctx context.Context, sample *model.LabSample) error {
	return r.db.WithContext(ctx).Create(sample).Error
}

func (r *LabSampleRepository) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.LabSample{}).Where("id = ?", id).Updates(updates).Error
}

func (r *LabSampleRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.LabSample{}, id).Error
}

// LabTestItemRepository 检测项目仓储
type LabTestItemRepository struct {
	db *gorm.DB
}

func NewLabTestItemRepository(db *gorm.DB) *LabTestItemRepository {
	return &LabTestItemRepository{db: db}
}

func (r *LabTestItemRepository) ListBySampleID(ctx context.Context, sampleID uint64) ([]model.LabTestItem, error) {
	var items []model.LabTestItem
	err := r.db.WithContext(ctx).Where("sample_id = ?", sampleID).Order("id ASC").Find(&items).Error
	return items, err
}

func (r *LabTestItemRepository) GetByID(ctx context.Context, id uint64) (*model.LabTestItem, error) {
	var item model.LabTestItem
	err := r.db.WithContext(ctx).First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *LabTestItemRepository) Create(ctx context.Context, item *model.LabTestItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *LabTestItemRepository) BatchCreate(ctx context.Context, items []model.LabTestItem) error {
	if len(items) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Create(&items).Error
}

func (r *LabTestItemRepository) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.LabTestItem{}).Where("id = ?", id).Updates(updates).Error
}

func (r *LabTestItemRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.LabTestItem{}, id).Error
}

func (r *LabTestItemRepository) DeleteBySampleID(ctx context.Context, sampleID uint64) error {
	return r.db.WithContext(ctx).Where("sample_id = ?", sampleID).Delete(&model.LabTestItem{}).Error
}

// LabReportRepository 检测报告仓储
type LabReportRepository struct {
	db *gorm.DB
}

func NewLabReportRepository(db *gorm.DB) *LabReportRepository {
	return &LabReportRepository{db: db}
}

func (r *LabReportRepository) List(ctx context.Context, query *model.LabReportQuery) ([]model.LabReport, int64, error) {
	var list []model.LabReport
	var total int64

	err := r.db.WithContext(ctx).Model(&model.LabReport{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (query.Page - 1) * query.PageSize
	err = r.db.WithContext(ctx).Order("id DESC").Offset(offset).Limit(query.PageSize).Find(&list).Error
	return list, total, err
}

func (r *LabReportRepository) GetByID(ctx context.Context, id uint64) (*model.LabReport, error) {
	var report model.LabReport
	err := r.db.WithContext(ctx).First(&report, id).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *LabReportRepository) GetBySampleID(ctx context.Context, sampleID uint64) (*model.LabReport, error) {
	var report model.LabReport
	err := r.db.WithContext(ctx).Where("sample_id = ?", sampleID).First(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *LabReportRepository) Create(ctx context.Context, report *model.LabReport) error {
	return r.db.WithContext(ctx).Create(report).Error
}

func (r *LabReportRepository) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.LabReport{}).Where("id = ?", id).Updates(updates).Error
}

func (r *LabReportRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.LabReport{}, id).Error
}
