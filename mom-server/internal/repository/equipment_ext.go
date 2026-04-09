package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// ========== TEEP分析 ==========

type TEEPDataRepository struct {
	db *gorm.DB
}

func NewTEEPDataRepository(db *gorm.DB) *TEEPDataRepository {
	return &TEEPDataRepository{db: db}
}

func (r *TEEPDataRepository) List(ctx context.Context, tenantID int64, equipmentID int64, startDate, endDate string) ([]model.TEEPData, int64, error) {
	var list []model.TEEPData
	var total int64

	query := r.db.WithContext(ctx).Model(&model.TEEPData{}).Where("tenant_id = ?", tenantID)
	if equipmentID > 0 {
		query = query.Where("equipment_id = ?", equipmentID)
	}
	if startDate != "" {
		query = query.Where("report_date >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("report_date <= ?", endDate)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Order("report_date DESC").Find(&list).Error
	return list, total, err
}

func (r *TEEPDataRepository) GetByID(ctx context.Context, id uint) (*model.TEEPData, error) {
	var data model.TEEPData
	err := r.db.WithContext(ctx).First(&data, id).Error
	return &data, err
}

func (r *TEEPDataRepository) Create(ctx context.Context, data *model.TEEPData) error {
	return r.db.WithContext(ctx).Create(data).Error
}

func (r *TEEPDataRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.TEEPData{}).Where("id = ?", id).Updates(updates).Error
}

func (r *TEEPDataRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.TEEPData{}, id).Error
}

// ========== 模具管理 ==========

type MoldRepository struct {
	db *gorm.DB
}

func NewMoldRepository(db *gorm.DB) *MoldRepository {
	return &MoldRepository{db: db}
}

func (r *MoldRepository) List(ctx context.Context, tenantID int64, query string) ([]model.Mold, int64, error) {
	var list []model.Mold
	var total int64

	db := r.db.WithContext(ctx).Model(&model.Mold{}).Where("tenant_id = ?", tenantID)
	if query != "" {
		db = db.Where("mold_code LIKE ? OR mold_name LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *MoldRepository) GetByID(ctx context.Context, id uint) (*model.Mold, error) {
	var mold model.Mold
	err := r.db.WithContext(ctx).First(&mold, id).Error
	return &mold, err
}

func (r *MoldRepository) GetByCode(ctx context.Context, tenantID int64, code string) (*model.Mold, error) {
	var mold model.Mold
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND mold_code = ?", tenantID, code).First(&mold).Error
	return &mold, err
}

func (r *MoldRepository) Create(ctx context.Context, mold *model.Mold) error {
	return r.db.WithContext(ctx).Create(mold).Error
}

func (r *MoldRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Mold{}).Where("id = ?", id).Updates(updates).Error
}

func (r *MoldRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Mold{}, id).Error
}

type MoldMaintenanceRepository struct {
	db *gorm.DB
}

func NewMoldMaintenanceRepository(db *gorm.DB) *MoldMaintenanceRepository {
	return &MoldMaintenanceRepository{db: db}
}

func (r *MoldMaintenanceRepository) List(ctx context.Context, tenantID int64, moldID int64) ([]model.MoldMaintenance, int64, error) {
	var list []model.MoldMaintenance
	var total int64

	query := r.db.WithContext(ctx).Model(&model.MoldMaintenance{}).Where("tenant_id = ?", tenantID)
	if moldID > 0 {
		query = query.Where("mold_id = ?", moldID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Order("maint_date DESC").Find(&list).Error
	return list, total, err
}

func (r *MoldMaintenanceRepository) Create(ctx context.Context, m *model.MoldMaintenance) error {
	return r.db.WithContext(ctx).Create(m).Error
}

type MoldRepairRepository struct {
	db *gorm.DB
}

func NewMoldRepairRepository(db *gorm.DB) *MoldRepairRepository {
	return &MoldRepairRepository{db: db}
}

func (r *MoldRepairRepository) List(ctx context.Context, tenantID int64, moldID int64) ([]model.MoldRepair, int64, error) {
	var list []model.MoldRepair
	var total int64

	query := r.db.WithContext(ctx).Model(&model.MoldRepair{}).Where("tenant_id = ?", tenantID)
	if moldID > 0 {
		query = query.Where("mold_id = ?", moldID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Order("repair_date DESC").Find(&list).Error
	return list, total, err
}

func (r *MoldRepairRepository) Create(ctx context.Context, m *model.MoldRepair) error {
	return r.db.WithContext(ctx).Create(m).Error
}

// ========== 量检具管理 ==========

type GaugeRepository struct {
	db *gorm.DB
}

func NewGaugeRepository(db *gorm.DB) *GaugeRepository {
	return &GaugeRepository{db: db}
}

func (r *GaugeRepository) List(ctx context.Context, tenantID int64, query string) ([]model.Gauge, int64, error) {
	var list []model.Gauge
	var total int64

	db := r.db.WithContext(ctx).Model(&model.Gauge{}).Where("tenant_id = ?", tenantID)
	if query != "" {
		db = db.Where("gauge_code LIKE ? OR gauge_name LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *GaugeRepository) GetByID(ctx context.Context, id uint) (*model.Gauge, error) {
	var gauge model.Gauge
	err := r.db.WithContext(ctx).First(&gauge, id).Error
	return &gauge, err
}

func (r *GaugeRepository) GetByCode(ctx context.Context, tenantID int64, code string) (*model.Gauge, error) {
	var gauge model.Gauge
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND gauge_code = ?", tenantID, code).First(&gauge).Error
	return &gauge, err
}

func (r *GaugeRepository) Create(ctx context.Context, gauge *model.Gauge) error {
	return r.db.WithContext(ctx).Create(gauge).Error
}

func (r *GaugeRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Gauge{}).Where("id = ?", id).Updates(updates).Error
}

func (r *GaugeRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Gauge{}, id).Error
}

type GaugeCalibrationRepository struct {
	db *gorm.DB
}

func NewGaugeCalibrationRepository(db *gorm.DB) *GaugeCalibrationRepository {
	return &GaugeCalibrationRepository{db: db}
}

func (r *GaugeCalibrationRepository) List(ctx context.Context, tenantID int64, gaugeID int64) ([]model.GaugeCalibration, int64, error) {
	var list []model.GaugeCalibration
	var total int64

	query := r.db.WithContext(ctx).Model(&model.GaugeCalibration{}).Where("tenant_id = ?", tenantID)
	if gaugeID > 0 {
		query = query.Where("gauge_id = ?", gaugeID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Order("cal_date DESC").Find(&list).Error
	return list, total, err
}

func (r *GaugeCalibrationRepository) Create(ctx context.Context, c *model.GaugeCalibration) error {
	return r.db.WithContext(ctx).Create(c).Error
}
