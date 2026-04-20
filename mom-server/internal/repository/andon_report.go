package repository

import (
	"context"
	"time"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

type AndonReportRepository struct {
	db *gorm.DB
}

func NewAndonReportRepository(db *gorm.DB) *AndonReportRepository {
	return &AndonReportRepository{db: db}
}

func (r *AndonReportRepository) List(ctx context.Context, tenantID int64, startDate, endDate *time.Time, workshopID, lineID, stationID int64, page, pageSize int) ([]model.AndonReport, int64, error) {
	var list []model.AndonReport
	var total int64

	query := r.db.WithContext(ctx).Model(&model.AndonReport{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if workshopID > 0 {
		query = query.Where("workshop_id = ?", workshopID)
	}
	if lineID > 0 {
		query = query.Where("line_id = ?", lineID)
	}
	if stationID > 0 {
		query = query.Where("station_id = ?", stationID)
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

func (r *AndonReportRepository) GetByID(ctx context.Context, id int64) (*model.AndonReport, error) {
	var report model.AndonReport
	err := r.db.WithContext(ctx).First(&report, id).Error
	return &report, err
}

func (r *AndonReportRepository) GetByDateAndStation(ctx context.Context, reportDate time.Time, stationID, tenantID int64) (*model.AndonReport, error) {
	var report model.AndonReport
	err := r.db.WithContext(ctx).
		Where("report_date = ? AND station_id = ? AND tenant_id = ?", reportDate, stationID, tenantID).
		First(&report).Error
	return &report, err
}

func (r *AndonReportRepository) Create(ctx context.Context, report *model.AndonReport) error {
	return r.db.WithContext(ctx).Create(report).Error
}

func (r *AndonReportRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.AndonReport{}).Where("id = ?", id).Updates(updates).Error
}

func (r *AndonReportRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.AndonReport{}, id).Error
}