package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type OperationRepository struct {
	db *gorm.DB
}

func NewOperationRepository(db *gorm.DB) *OperationRepository {
	return &OperationRepository{db: db}
}

func (r *OperationRepository) List(ctx context.Context, tenantID int64) ([]model.MdmOperation, int64, error) {
	var list []model.MdmOperation
	var total int64

	query := r.db.WithContext(ctx).Model(&model.MdmOperation{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("sequence ASC, id DESC").Find(&list).Error
	return list, total, err
}

func (r *OperationRepository) GetByID(ctx context.Context, id uint) (*model.MdmOperation, error) {
	var op model.MdmOperation
	err := r.db.WithContext(ctx).First(&op, id).Error
	return &op, err
}

func (r *OperationRepository) Create(ctx context.Context, op *model.MdmOperation) error {
	return r.db.WithContext(ctx).Create(op).Error
}

func (r *OperationRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.MdmOperation{}).Where("id = ?", id).Updates(updates).Error
}

func (r *OperationRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.MdmOperation{}, id).Error
}
