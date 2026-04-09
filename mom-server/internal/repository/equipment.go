package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type EquipmentRepository struct {
	db *gorm.DB
}

func NewEquipmentRepository(db *gorm.DB) *EquipmentRepository {
	return &EquipmentRepository{db: db}
}

func (r *EquipmentRepository) List(ctx context.Context, tenantID int64) ([]model.Equipment, int64, error) {
	var list []model.Equipment
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Equipment{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *EquipmentRepository) GetByID(ctx context.Context, id uint) (*model.Equipment, error) {
	var equipment model.Equipment
	err := r.db.WithContext(ctx).First(&equipment, id).Error
	return &equipment, err
}

func (r *EquipmentRepository) Create(ctx context.Context, equipment *model.Equipment) error {
	return r.db.WithContext(ctx).Create(equipment).Error
}

func (r *EquipmentRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Equipment{}).Where("id = ?", id).Updates(updates).Error
}

func (r *EquipmentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Equipment{}, id).Error
}

type EquipmentCheckRepository struct {
	db *gorm.DB
}

func NewEquipmentCheckRepository(db *gorm.DB) *EquipmentCheckRepository {
	return &EquipmentCheckRepository{db: db}
}

func (r *EquipmentCheckRepository) List(ctx context.Context, tenantID int64) ([]model.EquipmentCheck, int64, error) {
	var list []model.EquipmentCheck
	var total int64
	query := r.db.WithContext(ctx).Model(&model.EquipmentCheck{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *EquipmentCheckRepository) Create(ctx context.Context, check *model.EquipmentCheck) error {
	return r.db.WithContext(ctx).Create(check).Error
}

func (r *EquipmentCheckRepository) GetByID(ctx context.Context, id uint) (*model.EquipmentCheck, error) {
	var check model.EquipmentCheck
	err := r.db.WithContext(ctx).First(&check, id).Error
	return &check, err
}

func (r *EquipmentCheckRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.EquipmentCheck{}).Where("id = ?", id).Updates(updates).Error
}

func (r *EquipmentCheckRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.EquipmentCheck{}, id).Error
}

type EquipmentMaintenanceRepository struct {
	db *gorm.DB
}

func NewEquipmentMaintenanceRepository(db *gorm.DB) *EquipmentMaintenanceRepository {
	return &EquipmentMaintenanceRepository{db: db}
}

func (r *EquipmentMaintenanceRepository) List(ctx context.Context, tenantID int64) ([]model.EquipmentMaintenance, int64, error) {
	var list []model.EquipmentMaintenance
	var total int64

	err := r.db.WithContext(ctx).Model(&model.EquipmentMaintenance{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *EquipmentMaintenanceRepository) Create(ctx context.Context, m *model.EquipmentMaintenance) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *EquipmentMaintenanceRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.EquipmentMaintenance{}).Where("id = ?", id).Updates(updates).Error
}

func (r *EquipmentMaintenanceRepository) GetByID(ctx context.Context, id uint) (*model.EquipmentMaintenance, error) {
	var m model.EquipmentMaintenance
	err := r.db.WithContext(ctx).First(&m, id).Error
	return &m, err
}

func (r *EquipmentMaintenanceRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.EquipmentMaintenance{}, id).Error
}

type EquipmentRepairRepository struct {
	db *gorm.DB
}

func NewEquipmentRepairRepository(db *gorm.DB) *EquipmentRepairRepository {
	return &EquipmentRepairRepository{db: db}
}

func (r *EquipmentRepairRepository) List(ctx context.Context, tenantID int64) ([]model.EquipmentRepair, int64, error) {
	var list []model.EquipmentRepair
	var total int64

	err := r.db.WithContext(ctx).Model(&model.EquipmentRepair{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *EquipmentRepairRepository) Create(ctx context.Context, m *model.EquipmentRepair) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *EquipmentRepairRepository) GetByID(ctx context.Context, id uint) (*model.EquipmentRepair, error) {
	var repair model.EquipmentRepair
	err := r.db.WithContext(ctx).First(&repair, id).Error
	return &repair, err
}

func (r *EquipmentRepairRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.EquipmentRepair{}).Where("id = ?", id).Updates(updates).Error
}

func (r *EquipmentRepairRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.EquipmentRepair{}, id).Error
}

type SparePartRepository struct {
	db *gorm.DB
}

func NewSparePartRepository(db *gorm.DB) *SparePartRepository {
	return &SparePartRepository{db: db}
}

func (r *SparePartRepository) List(ctx context.Context, tenantID int64) ([]model.SparePart, int64, error) {
	var list []model.SparePart
	var total int64
	query := r.db.WithContext(ctx).Model(&model.SparePart{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *SparePartRepository) Create(ctx context.Context, sp *model.SparePart) error {
	return r.db.WithContext(ctx).Create(sp).Error
}

func (r *SparePartRepository) GetByID(ctx context.Context, id uint) (*model.SparePart, error) {
	var sp model.SparePart
	err := r.db.WithContext(ctx).First(&sp, id).Error
	return &sp, err
}

func (r *SparePartRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.SparePart{}).Where("id = ?", id).Updates(updates).Error
}

func (r *SparePartRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.SparePart{}, id).Error
}
