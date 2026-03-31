package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type OQCRepository struct {
	db *gorm.DB
}

func NewOQCRepository(db *gorm.DB) *OQCRepository {
	return &OQCRepository{db: db}
}

func (r *OQCRepository) List(ctx context.Context, tenantID int64) ([]model.OQC, int64, error) {
	var list []model.OQC
	var total int64

	err := r.db.WithContext(ctx).Model(&model.OQC{}).Where("tenant_id = ?", tenantID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *OQCRepository) GetByID(ctx context.Context, id uint) (*model.OQC, error) {
	var oqc model.OQC
	err := r.db.WithContext(ctx).First(&oqc, id).Error
	return &oqc, err
}

func (r *OQCRepository) Create(ctx context.Context, oqc *model.OQC) error {
	return r.db.WithContext(ctx).Create(oqc).Error
}

func (r *OQCRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.OQC{}).Where("id = ?", id).Updates(updates).Error
}

func (r *OQCRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.OQC{}, id).Error
}
