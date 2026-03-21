package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type MaterialRepository struct {
	db *gorm.DB
}

func NewMaterialRepository(db *gorm.DB) *MaterialRepository {
	return &MaterialRepository{db: db}
}

func (r *MaterialRepository) List(ctx context.Context, tenantID int64) ([]model.Material, int64, error) {
	var list []model.Material
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Material{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *MaterialRepository) GetByID(ctx context.Context, id uint) (*model.Material, error) {
	var material model.Material
	err := r.db.WithContext(ctx).First(&material, id).Error
	return &material, err
}

func (r *MaterialRepository) Create(ctx context.Context, material *model.Material) error {
	return r.db.WithContext(ctx).Create(material).Error
}

func (r *MaterialRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Material{}).Where("id = ?", id).Updates(updates).Error
}

func (r *MaterialRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Material{}, id).Error
}
