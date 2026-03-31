package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type DefectRecordRepository struct {
	db *gorm.DB
}

func NewDefectRecordRepository(db *gorm.DB) *DefectRecordRepository {
	return &DefectRecordRepository{db: db}
}

func (r *DefectRecordRepository) List(ctx context.Context, tenantID int64) ([]model.DefectRecord, int64, error) {
	var list []model.DefectRecord
	var total int64

	err := r.db.WithContext(ctx).Model(&model.DefectRecord{}).Where("tenant_id = ?", tenantID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *DefectRecordRepository) GetByID(ctx context.Context, id uint) (*model.DefectRecord, error) {
	var record model.DefectRecord
	err := r.db.WithContext(ctx).First(&record, id).Error
	return &record, err
}

func (r *DefectRecordRepository) Create(ctx context.Context, record *model.DefectRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *DefectRecordRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.DefectRecord{}).Where("id = ?", id).Updates(updates).Error
}

func (r *DefectRecordRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.DefectRecord{}, id).Error
}
