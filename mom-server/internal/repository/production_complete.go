package repository

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"time"

	"gorm.io/gorm"
)

type ProductionCompleteRepository struct {
	db *gorm.DB
}

func NewProductionCompleteRepository(db *gorm.DB) *ProductionCompleteRepository {
	return &ProductionCompleteRepository{db: db}
}

func (r *ProductionCompleteRepository) Create(ctx context.Context, item *model.ProductionComplete) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *ProductionCompleteRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.ProductionComplete{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ProductionCompleteRepository) GetByID(ctx context.Context, id uint) (*model.ProductionComplete, error) {
	var item model.ProductionComplete
	err := r.db.WithContext(ctx).Preload("Items").First(&item, id).Error
	return &item, err
}

func (r *ProductionCompleteRepository) List(ctx context.Context, tenantID int64, query string, page, pageSize int) ([]model.ProductionComplete, int64, error) {
	var list []model.ProductionComplete
	var total int64

	q := r.db.WithContext(ctx).Model(&model.ProductionComplete{}).Where("tenant_id = ?", tenantID)
	if query != "" {
		q = q.Where("complete_no LIKE ?", "%"+query+"%")
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * pageSize
	if err := q.Preload("Items").Order("id DESC").Offset(offset).Limit(pageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *ProductionCompleteRepository) GetByProductionOrderID(ctx context.Context, orderID int64) ([]model.ProductionComplete, error) {
	var list []model.ProductionComplete
	err := r.db.WithContext(ctx).Preload("Items").Where("production_order_id = ?", orderID).Order("id DESC").Find(&list).Error
	return list, err
}

// ProductionStockInRepository
type ProductionStockInRepository struct {
	db *gorm.DB
}

func NewProductionStockInRepository(db *gorm.DB) *ProductionStockInRepository {
	return &ProductionStockInRepository{db: db}
}

func (r *ProductionStockInRepository) Create(ctx context.Context, item *model.ProductionStockIn) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *ProductionStockInRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.ProductionStockIn{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ProductionStockInRepository) GetByID(ctx context.Context, id uint) (*model.ProductionStockIn, error) {
	var item model.ProductionStockIn
	err := r.db.WithContext(ctx).Preload("Items").First(&item, id).Error
	return &item, err
}

func (r *ProductionStockInRepository) GetByCompleteID(ctx context.Context, completeID uint) (*model.ProductionStockIn, error) {
	var item model.ProductionStockIn
	err := r.db.WithContext(ctx).Where("complete_id = ?", completeID).First(&item).Error
	return &item, err
}

func (r *ProductionStockInRepository) List(ctx context.Context, tenantID int64, query string, page, pageSize int) ([]model.ProductionStockIn, int64, error) {
	var list []model.ProductionStockIn
	var total int64

	q := r.db.WithContext(ctx).Model(&model.ProductionStockIn{}).Where("tenant_id = ?", tenantID)
	if query != "" {
		q = q.Where("stock_in_no LIKE ?", "%"+query+"%")
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * pageSize
	if err := q.Preload("Items").Order("id DESC").Offset(offset).Limit(pageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *ProductionStockInRepository) GenerateStockInNo(ctx context.Context, tenantID int64) (string, error) {
	var count int64
	today := time.Now().Format("2006-01-02")
	r.db.WithContext(ctx).Model(&model.ProductionStockIn{}).Where("DATE(created_at) = ?", today).Count(&count)
	return fmt.Sprintf("PSI%s%04d", today[2:10], count+1), nil
}
