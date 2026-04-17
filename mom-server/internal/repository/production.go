package repository

import (
	"context"
	"time"

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

// GetStatsByDateRange 获取日期范围内的工单统计数据
func (r *ProductionOrderRepository) GetStatsByDateRange(ctx context.Context, tenantID int64, startDate, endDate time.Time, workshopID int64) (map[string]interface{}, error) {
	type Stats struct {
		OrderCount        int
		CompletedCount    int
		TotalOutput       float64
		QualifiedQty      float64
		DefectQty         float64
		FirstPassQty      float64
	}
	var stats Stats

	query := r.db.WithContext(ctx).Model(&model.ProductionOrder{}).
		Where("tenant_id = ?", tenantID).
		Where("DATE(actual_end_date) >= ? AND DATE(actual_end_date) <= ?", startDate, endDate).
		Where("status = 3") // 已完成
	if workshopID > 0 {
		query = query.Where("workshop_id = ?", workshopID)
	}

	err := query.Select("COUNT(*) as order_count, "+
		"SUM(CASE WHEN status = 3 THEN 1 ELSE 0 END) as completed_count, "+
		"SUM(completed_qty) as total_output, "+
		"SUM(completed_qty - rejected_qty) as qualified_qty, "+
		"SUM(rejected_qty) as defect_qty, "+
		"0 as first_pass_qty").
		Scan(&stats).Error
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"order_count":     stats.OrderCount,
		"completed_count": stats.CompletedCount,
		"total_output":    stats.TotalOutput,
		"qualified_qty":   stats.QualifiedQty,
		"defect_qty":      stats.DefectQty,
		"first_pass_qty":  stats.FirstPassQty,
	}, nil
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

type SalesOrderRepository struct {
	db *gorm.DB
}

func NewSalesOrderRepository(db *gorm.DB) *SalesOrderRepository {
	return &SalesOrderRepository{db: db}
}

func (r *SalesOrderRepository) List(ctx context.Context, tenantID int64) ([]model.SalesOrder, int64, error) {
	var list []model.SalesOrder
	var total int64
	err := r.db.WithContext(ctx).Model(&model.SalesOrder{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *SalesOrderRepository) GetByID(ctx context.Context, id uint) (*model.SalesOrder, error) {
	var order model.SalesOrder
	err := r.db.WithContext(ctx).First(&order, id).Error
	return &order, err
}

func (r *SalesOrderRepository) Create(ctx context.Context, order *model.SalesOrder) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *SalesOrderRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.SalesOrder{}).Where("id = ?", id).Updates(updates).Error
}

func (r *SalesOrderRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.SalesOrder{}, id).Error
}
