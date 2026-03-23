package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type MdmShiftRepository struct {
	db *gorm.DB
}

func NewMdmShiftRepository(db *gorm.DB) *MdmShiftRepository {
	return &MdmShiftRepository{db: db}
}

func (r *MdmShiftRepository) List(ctx context.Context, tenantID int64) ([]model.MdmShift, int64, error) {
	var list []model.MdmShift
	var total int64

	query := r.db.WithContext(ctx).Model(&model.MdmShift{})
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

func (r *MdmShiftRepository) GetByID(ctx context.Context, id uint) (*model.MdmShift, error) {
	var shift model.MdmShift
	err := r.db.WithContext(ctx).First(&shift, id).Error
	return &shift, err
}

func (r *MdmShiftRepository) Create(ctx context.Context, shift *model.MdmShift) error {
	return r.db.WithContext(ctx).Create(shift).Error
}

func (r *MdmShiftRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.MdmShift{}).Where("id = ?", id).Updates(updates).Error
}

func (r *MdmShiftRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.MdmShift{}, id).Error
}
