package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type SPCDataRepository struct {
	db *gorm.DB
}

func NewSPCDataRepository(db *gorm.DB) *SPCDataRepository {
	return &SPCDataRepository{db: db}
}

func (r *SPCDataRepository) List(ctx context.Context, tenantID int64) ([]model.SPCData, int64, error) {
	var list []model.SPCData
	var total int64

	err := r.db.WithContext(ctx).Model(&model.SPCData{}).Where("tenant_id = ?", tenantID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("check_time DESC").Find(&list).Error
	return list, total, err
}

func (r *SPCDataRepository) GetByID(ctx context.Context, id uint) (*model.SPCData, error) {
	var spcData model.SPCData
	err := r.db.WithContext(ctx).First(&spcData, id).Error
	return &spcData, err
}

func (r *SPCDataRepository) Create(ctx context.Context, spcData *model.SPCData) error {
	return r.db.WithContext(ctx).Create(spcData).Error
}

func (r *SPCDataRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.SPCData{}).Where("id = ?", id).Updates(updates).Error
}

func (r *SPCDataRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.SPCData{}, id).Error
}

func (r *SPCDataRepository) GetChartData(ctx context.Context, tenantID int64, equipmentID, processID, stationID int64, checkItem string, limit int) ([]model.SPCData, error) {
	var list []model.SPCData
	query := r.db.WithContext(ctx).Model(&model.SPCData{}).Where("tenant_id = ?", tenantID)

	if equipmentID > 0 {
		query = query.Where("equipment_id = ?", equipmentID)
	}
	if processID > 0 {
		query = query.Where("process_id = ?", processID)
	}
	if stationID > 0 {
		query = query.Where("station_id = ?", stationID)
	}
	if checkItem != "" {
		query = query.Where("check_item = ?", checkItem)
	}

	err := query.Order("check_time ASC").Limit(limit).Find(&list).Error
	return list, err
}
