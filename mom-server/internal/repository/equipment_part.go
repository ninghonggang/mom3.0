package repository

import (
	"context"

	"mom-server/internal/model"
	"gorm.io/gorm"
)

// EquipmentPartRepository 设备部件仓储
type EquipmentPartRepository struct {
	db *gorm.DB
}

func NewEquipmentPartRepository(db *gorm.DB) *EquipmentPartRepository {
	return &EquipmentPartRepository{db: db}
}

func (r *EquipmentPartRepository) List(ctx context.Context, tenantID int64, query map[string]any) ([]model.EquipmentPart, int64, error) {
	var list []model.EquipmentPart
	var total int64

	q := r.db.WithContext(ctx).Model(&model.EquipmentPart{}).Where("tenant_id = ?", tenantID)

	if equipmentID, ok := query["equipment_id"]; ok && equipmentID.(uint) > 0 {
		q = q.Where("equipment_id = ?", equipmentID)
	}
	if status, ok := query["status"]; ok && status.(int) > 0 {
		q = q.Where("status = ?", status)
	}

	q.Count(&total)
	err := q.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *EquipmentPartRepository) GetByID(ctx context.Context, id uint) (*model.EquipmentPart, error) {
	var item model.EquipmentPart
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *EquipmentPartRepository) Create(ctx context.Context, item *model.EquipmentPart) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *EquipmentPartRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.EquipmentPart{}).Where("id = ?", id).Updates(updates).Error
}

func (r *EquipmentPartRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.EquipmentPart{}).Error
}

func (r *EquipmentPartRepository) ListByEquipmentID(ctx context.Context, equipmentID int64) ([]model.EquipmentPart, error) {
	var list []model.EquipmentPart
	err := r.db.WithContext(ctx).Where("equipment_id = ?", equipmentID).Order("id DESC").Find(&list).Error
	return list, err
}
