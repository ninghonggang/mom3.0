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

// InspectionCharacteristicRepository 检验特性仓储
type InspectionCharacteristicRepository struct {
	db *gorm.DB
}

func NewInspectionCharacteristicRepository(db *gorm.DB) *InspectionCharacteristicRepository {
	return &InspectionCharacteristicRepository{db: db}
}

func (r *InspectionCharacteristicRepository) List(ctx context.Context, tenantID int64, query string) ([]model.InspectionCharacteristic, int64, error) {
	var list []model.InspectionCharacteristic
	var total int64

	q := r.db.WithContext(ctx).Model(&model.InspectionCharacteristic{}).Where("tenant_id = ?", tenantID)
	if query != "" {
		q = q.Where("code LIKE ? OR name LIKE ?", "%"+query+"%", "%"+query+"%")
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := q.Order("id DESC").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *InspectionCharacteristicRepository) GetByID(ctx context.Context, id uint) (*model.InspectionCharacteristic, error) {
	var char model.InspectionCharacteristic
	err := r.db.WithContext(ctx).First(&char, id).Error
	return &char, err
}

func (r *InspectionCharacteristicRepository) Create(ctx context.Context, char *model.InspectionCharacteristic) error {
	return r.db.WithContext(ctx).Create(char).Error
}

func (r *InspectionCharacteristicRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.InspectionCharacteristic{}).Where("id = ?", id).Updates(updates).Error
}

func (r *InspectionCharacteristicRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.InspectionCharacteristic{}, id).Error
}
