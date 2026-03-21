package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type ProductionOrderRepository struct {
	db *gorm.DB
}

func NewProductionOrderRepository(db *gorm.DB) *ProductionOrderRepository {
	return &ProductionOrderRepository{db: db}
}

func (r *ProductionOrderRepository) List(ctx context.Context, tenantID int64) ([]model.ProductionOrder, int64, error) {
	var list []model.ProductionOrder
	var total int64

	err := r.db.WithContext(ctx).Model(&model.ProductionOrder{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *ProductionOrderRepository) GetByID(ctx context.Context, id uint) (*model.ProductionOrder, error) {
	var order model.ProductionOrder
	err := r.db.WithContext(ctx).First(&order, id).Error
	return &order, err
}

func (r *ProductionOrderRepository) Create(ctx context.Context, order *model.ProductionOrder) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *ProductionOrderRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ProductionOrder{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ProductionOrderRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.ProductionOrder{}, id).Error
}

type ProductionReportRepository struct {
	db *gorm.DB
}

func NewProductionReportRepository(db *gorm.DB) *ProductionReportRepository {
	return &ProductionReportRepository{db: db}
}

func (r *ProductionReportRepository) List(ctx context.Context, tenantID int64) ([]model.ProductionReport, int64, error) {
	var list []model.ProductionReport
	var total int64

	err := r.db.WithContext(ctx).Model(&model.ProductionReport{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *ProductionReportRepository) Create(ctx context.Context, report *model.ProductionReport) error {
	return r.db.WithContext(ctx).Create(report).Error
}

type DispatchRepository struct {
	db *gorm.DB
}

func NewDispatchRepository(db *gorm.DB) *DispatchRepository {
	return &DispatchRepository{db: db}
}

func (r *DispatchRepository) List(ctx context.Context, tenantID int64) ([]model.Dispatch, int64, error) {
	var list []model.Dispatch
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Dispatch{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *DispatchRepository) Create(ctx context.Context, dispatch *model.Dispatch) error {
	return r.db.WithContext(ctx).Create(dispatch).Error
}

func (r *DispatchRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Dispatch{}).Where("id = ?", id).Updates(updates).Error
}
