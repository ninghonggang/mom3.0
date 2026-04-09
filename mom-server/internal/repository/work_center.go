package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type WorkCenterRepository struct {
	db *gorm.DB
}

func NewWorkCenterRepository(db *gorm.DB) *WorkCenterRepository {
	return &WorkCenterRepository{db: db}
}

func (r *WorkCenterRepository) List(ctx context.Context, tenantID int64) ([]model.WorkCenter, int64, error) {
	var list []model.WorkCenter
	var total int64
	query := r.db.WithContext(ctx).Model(&model.WorkCenter{})
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

func (r *WorkCenterRepository) GetByID(ctx context.Context, id uint) (*model.WorkCenter, error) {
	var wc model.WorkCenter
	err := r.db.WithContext(ctx).First(&wc, id).Error
	return &wc, err
}

func (r *WorkCenterRepository) GetByCode(ctx context.Context, code string) (*model.WorkCenter, error) {
	var wc model.WorkCenter
	err := r.db.WithContext(ctx).Where("work_center_code = ?", code).First(&wc).Error
	return &wc, err
}

func (r *WorkCenterRepository) Create(ctx context.Context, wc *model.WorkCenter) error {
	return r.db.WithContext(ctx).Create(wc).Error
}

func (r *WorkCenterRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.WorkCenter{}).Where("id = ?", id).Updates(updates).Error
}

func (r *WorkCenterRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.WorkCenter{}, id).Error
}

func (r *WorkCenterRepository) ListByWorkshop(ctx context.Context, workshopID int64) ([]model.WorkCenter, error) {
	var list []model.WorkCenter
	err := r.db.WithContext(ctx).Where("workshop_id = ?", workshopID).Order("id DESC").Find(&list).Error
	return list, err
}
