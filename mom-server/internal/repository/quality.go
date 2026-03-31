package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type IQCRepository struct {
	db *gorm.DB
}

func NewIQCRepository(db *gorm.DB) *IQCRepository {
	return &IQCRepository{db: db}
}

func (r *IQCRepository) List(ctx context.Context, tenantID int64) ([]model.IQC, int64, error) {
	var list []model.IQC
	var total int64

	err := r.db.WithContext(ctx).Model(&model.IQC{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *IQCRepository) GetByID(ctx context.Context, id uint) (*model.IQC, error) {
	var iqc model.IQC
	err := r.db.WithContext(ctx).First(&iqc, id).Error
	return &iqc, err
}

func (r *IQCRepository) Create(ctx context.Context, iqc *model.IQC) error {
	return r.db.WithContext(ctx).Create(iqc).Error
}

func (r *IQCRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.IQC{}).Where("id = ?", id).Updates(updates).Error
}

func (r *IQCRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.IQC{}, id).Error
}
