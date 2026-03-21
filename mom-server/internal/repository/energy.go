package repository

import (
	"context"
	"mom-server/internal/model"
	"time"

	"gorm.io/gorm"
)

type EnergyRepository struct {
	db *gorm.DB
}

func NewEnergyRepository(db *gorm.DB) *EnergyRepository {
	return &EnergyRepository{db: db}
}

func (r *EnergyRepository) List(ctx context.Context, tenantID int64, energyType string, startDate, endDate time.Time) ([]model.EnergyRecord, int64, error) {
	var list []model.EnergyRecord
	var total int64
	query := r.db.WithContext(ctx).Model(&model.EnergyRecord{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if energyType != "" {
		query = query.Where("energy_type = ?", energyType)
	}
	if !startDate.IsZero() {
		query = query.Where("record_date >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("record_date <= ?", endDate)
	}
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *EnergyRepository) GetStats(ctx context.Context, tenantID int64, startDate, endDate time.Time) (map[string]interface{}, error) {
	var result struct {
		TotalQty   float64
		TotalAmount float64
	}
	err := r.db.WithContext(ctx).Model(&model.EnergyRecord{}).
		Where("tenant_id = ?", tenantID).
		Where("record_date >= ? AND record_date <= ?", startDate, endDate).
		Select("COALESCE(SUM(quantity), 0) as total_qty, COALESCE(SUM(amount), 0) as total_amount").
		Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"total_qty":    result.TotalQty,
		"total_amount":  result.TotalAmount,
	}, nil
}

func (r *EnergyRepository) GetTrend(ctx context.Context, tenantID int64, energyType string, startDate, endDate time.Time) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := r.db.WithContext(ctx).Model(&model.EnergyRecord{}).
		Where("tenant_id = ?", tenantID).
		Where("record_date >= ? AND record_date <= ?", startDate, endDate).
		Select("record_date, SUM(quantity) as quantity, SUM(amount) as amount").
		Group("record_date").
		Order("record_date ASC").
		Scan(&results).Error
	return results, err
}

func (r *EnergyRepository) Create(ctx context.Context, record *model.EnergyRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}
