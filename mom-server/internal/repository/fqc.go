package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type FQCRepository struct {
	db *gorm.DB
}

func NewFQCRepository(db *gorm.DB) *FQCRepository {
	return &FQCRepository{db: db}
}

func (r *FQCRepository) List(ctx context.Context, tenantID int64) ([]model.FQC, int64, error) {
	var list []model.FQC
	var total int64

	err := r.db.WithContext(ctx).Model(&model.FQC{}).Where("tenant_id = ?", tenantID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *FQCRepository) GetByID(ctx context.Context, id uint) (*model.FQC, error) {
	var fqc model.FQC
	err := r.db.WithContext(ctx).First(&fqc, id).Error
	return &fqc, err
}

func (r *FQCRepository) Create(ctx context.Context, fqc *model.FQC) error {
	return r.db.WithContext(ctx).Create(fqc).Error
}

func (r *FQCRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.FQC{}).Where("id = ?", id).Updates(updates).Error
}

func (r *FQCRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.FQC{}, id).Error
}
