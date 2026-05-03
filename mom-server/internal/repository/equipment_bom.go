package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// EquipmentBomRepository 设备BOM仓储
type EquipmentBomRepository struct {
	db *gorm.DB
}

func NewEquipmentBomRepository(db *gorm.DB) *EquipmentBomRepository {
	return &EquipmentBomRepository{db: db}
}

func (r *EquipmentBomRepository) Create(ctx context.Context, m *model.EquipmentBom) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *EquipmentBomRepository) GetByID(ctx context.Context, id int64) (*model.EquipmentBom, error) {
	var m model.EquipmentBom
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	return &m, err
}

func (r *EquipmentBomRepository) ListByEquipmentID(ctx context.Context, equipmentID int64) ([]model.EquipmentBom, error) {
	var list []model.EquipmentBom
	err := r.db.WithContext(ctx).Where("equipment_id = ? AND status = 1", equipmentID).Order("is_critical DESC, id ASC").Find(&list).Error
	return list, err
}

func (r *EquipmentBomRepository) Page(ctx context.Context, tenantID int64, req *model.EquipmentBomQuery) ([]model.EquipmentBom, int64, error) {
	var list []model.EquipmentBom
	query := r.db.WithContext(ctx).Model(&model.EquipmentBom{}).Where("tenant_id = ?", tenantID)
	if req.EquipmentID > 0 {
		query = query.Where("equipment_id = ?", req.EquipmentID)
	}
	if req.MaterialCode != "" {
		query = query.Where("material_code LIKE ?", "%"+req.MaterialCode+"%")
	}
	if req.MaterialName != "" {
		query = query.Where("material_name LIKE ?", "%"+req.MaterialName+"%")
	}
	if req.IsCritical > 0 {
		query = query.Where("is_critical = ?", req.IsCritical)
	}
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	err = query.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *EquipmentBomRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.EquipmentBom{}).Where("id = ?", id).Updates(updates).Error
}

func (r *EquipmentBomRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.EquipmentBom{}, id).Error
}