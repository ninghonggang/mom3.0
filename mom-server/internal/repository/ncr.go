package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type NCRRepository struct {
	db *gorm.DB
}

func NewNCRRepository(db *gorm.DB) *NCRRepository {
	return &NCRRepository{db: db}
}

func (r *NCRRepository) List(ctx context.Context, tenantID int64) ([]model.NCR, int64, error) {
	var list []model.NCR
	var total int64

	err := r.db.WithContext(ctx).Model(&model.NCR{}).Where("tenant_id = ?", tenantID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *NCRRepository) GetByID(ctx context.Context, id uint) (*model.NCR, error) {
	var ncr model.NCR
	err := r.db.WithContext(ctx).First(&ncr, id).Error
	return &ncr, err
}

func (r *NCRRepository) Create(ctx context.Context, ncr *model.NCR) error {
	return r.db.WithContext(ctx).Create(ncr).Error
}

func (r *NCRRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.NCR{}).Where("id = ?", id).Updates(updates).Error
}

func (r *NCRRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.NCR{}, id).Error
}
