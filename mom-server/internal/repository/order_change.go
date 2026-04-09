package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type ProductionOrderChangeLogRepository struct {
	db *gorm.DB
}

func NewProductionOrderChangeLogRepository(db *gorm.DB) *ProductionOrderChangeLogRepository {
	return &ProductionOrderChangeLogRepository{db: db}
}

func (r *ProductionOrderChangeLogRepository) Create(ctx context.Context, log *model.ProductionOrderChangeLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *ProductionOrderChangeLogRepository) ListByOrderID(ctx context.Context, orderID int64) ([]model.ProductionOrderChangeLog, error) {
	var list []model.ProductionOrderChangeLog
	err := r.db.WithContext(ctx).Where("order_id = ?", orderID).Order("created_at DESC").Find(&list).Error
	return list, err
}
