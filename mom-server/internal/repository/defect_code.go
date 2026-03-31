package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type DefectCodeRepository struct {
	db *gorm.DB
}

func NewDefectCodeRepository(db *gorm.DB) *DefectCodeRepository {
	return &DefectCodeRepository{db: db}
}

func (r *DefectCodeRepository) List(ctx context.Context, tenantID int64) ([]model.DefectCode, int64, error) {
	var list []model.DefectCode
	var total int64

	err := r.db.WithContext(ctx).Model(&model.DefectCode{}).Where("tenant_id = ?", tenantID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *DefectCodeRepository) GetByID(ctx context.Context, id uint) (*model.DefectCode, error) {
	var defectCode model.DefectCode
	err := r.db.WithContext(ctx).First(&defectCode, id).Error
	return &defectCode, err
}

func (r *DefectCodeRepository) Create(ctx context.Context, defectCode *model.DefectCode) error {
	return r.db.WithContext(ctx).Create(defectCode).Error
}

func (r *DefectCodeRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.DefectCode{}).Where("id = ?", id).Updates(updates).Error
}

func (r *DefectCodeRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.DefectCode{}, id).Error
}
