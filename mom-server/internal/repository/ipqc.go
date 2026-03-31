package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type IPQCRepository struct {
	db *gorm.DB
}

func NewIPQCRepository(db *gorm.DB) *IPQCRepository {
	return &IPQCRepository{db: db}
}

func (r *IPQCRepository) List(ctx context.Context, tenantID int64) ([]model.IPQC, int64, error) {
	var list []model.IPQC
	var total int64

	err := r.db.WithContext(ctx).Model(&model.IPQC{}).Where("tenant_id = ?", tenantID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *IPQCRepository) GetByID(ctx context.Context, id uint) (*model.IPQC, error) {
	var ipqc model.IPQC
	err := r.db.WithContext(ctx).First(&ipqc, id).Error
	return &ipqc, err
}

func (r *IPQCRepository) Create(ctx context.Context, ipqc *model.IPQC) error {
	return r.db.WithContext(ctx).Create(ipqc).Error
}

func (r *IPQCRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.IPQC{}).Where("id = ?", id).Updates(updates).Error
}

func (r *IPQCRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.IPQC{}, id).Error
}
