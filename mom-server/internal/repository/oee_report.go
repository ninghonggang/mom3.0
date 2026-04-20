package repository

import (
	"context"
	"time"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

type OEEReportRepository struct {
	db *gorm.DB
}

func NewOEEReportRepository(db *gorm.DB) *OEEReportRepository {
	return &OEEReportRepository{db: db}
}

func (r *OEEReportRepository) List(ctx context.Context, tenantID int64, startDate, endDate *time.Time, workshopID, lineID int64, page, pageSize int) ([]model.OEEReport, int64, error) {
	var list []model.OEEReport
	var total int64

	query := r.db.WithContext(ctx).Model(&model.OEEReport{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if workshopID > 0 {
		query = query.Where("workshop_id = ?", workshopID)
	}
	if lineID > 0 {
		query = query.Where("line_id = ?", lineID)
	}
	if startDate != nil {
		query = query.Where("report_date >= ?", startDate)
	}
	if endDate != nil {
		query = query.Where("report_date <= ?", endDate)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		query = query.Offset(offset).Limit(pageSize)
	}

	err = query.Order("report_date DESC").Find(&list).Error
	return list, total, err
}

func (r *OEEReportRepository) GetByID(ctx context.Context, id int64) (*model.OEEReport, error) {
	var report model.OEEReport
	err := r.db.WithContext(ctx).First(&report, id).Error
	return &report, err
}

func (r *OEEReportRepository) GetByDateAndLine(ctx context.Context, reportDate time.Time, lineID, tenantID int64) (*model.OEEReport, error) {
	var report model.OEEReport
	err := r.db.WithContext(ctx).
		Where("report_date = ? AND line_id = ? AND tenant_id = ?", reportDate, lineID, tenantID).
		First(&report).Error
	return &report, err
}

func (r *OEEReportRepository) Create(ctx context.Context, report *model.OEEReport) error {
	return r.db.WithContext(ctx).Create(report).Error
}

func (r *OEEReportRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.OEEReport{}).Where("id = ?", id).Updates(updates).Error
}

func (r *OEEReportRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.OEEReport{}, id).Error
}