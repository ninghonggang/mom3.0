package repository

import (
	"context"
	"time"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

type DeliveryReportRepository struct {
	db *gorm.DB
}

func NewDeliveryReportRepository(db *gorm.DB) *DeliveryReportRepository {
	return &DeliveryReportRepository{db: db}
}

func (r *DeliveryReportRepository) List(ctx context.Context, tenantID int64, startMonth, endMonth *time.Time, customerID int64, page, pageSize int) ([]model.DeliveryReport, int64, error) {
	var list []model.DeliveryReport
	var total int64

	query := r.db.WithContext(ctx).Model(&model.DeliveryReport{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if customerID > 0 {
		query = query.Where("customer_id = ?", customerID)
	}
	if startMonth != nil {
		query = query.Where("report_month >= ?", startMonth)
	}
	if endMonth != nil {
		query = query.Where("report_month <= ?", endMonth)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		query = query.Offset(offset).Limit(pageSize)
	}

	err = query.Order("report_month DESC").Find(&list).Error
	return list, total, err
}

func (r *DeliveryReportRepository) GetByID(ctx context.Context, id int64) (*model.DeliveryReport, error) {
	var report model.DeliveryReport
	err := r.db.WithContext(ctx).First(&report, id).Error
	return &report, err
}

func (r *DeliveryReportRepository) GetByMonthAndCustomer(ctx context.Context, reportMonth time.Time, customerID, tenantID int64) (*model.DeliveryReport, error) {
	var report model.DeliveryReport
	err := r.db.WithContext(ctx).
		Where("report_month = ? AND customer_id = ? AND tenant_id = ?", reportMonth, customerID, tenantID).
		First(&report).Error
	return &report, err
}

func (r *DeliveryReportRepository) Create(ctx context.Context, report *model.DeliveryReport) error {
	return r.db.WithContext(ctx).Create(report).Error
}

func (r *DeliveryReportRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.DeliveryReport{}).Where("id = ?", id).Updates(updates).Error
}

func (r *DeliveryReportRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.DeliveryReport{}, id).Error
}