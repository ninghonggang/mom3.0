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
