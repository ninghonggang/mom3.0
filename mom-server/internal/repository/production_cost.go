package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// ProductionCostRepository 生产成本仓储
type ProductionCostRepository struct {
	db *gorm.DB
}

func NewProductionCostRepository(db *gorm.DB) *ProductionCostRepository {
	return &ProductionCostRepository{db: db}
}

func (r *ProductionCostRepository) Create(ctx context.Context, m *model.ProductionCost) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *ProductionCostRepository) GetByID(ctx context.Context, id int64) (*model.ProductionCost, error) {
	var m model.ProductionCost
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	return &m, err
}

func (r *ProductionCostRepository) ListByOrderID(ctx context.Context, orderID int64) ([]model.ProductionCost, error) {
	var list []model.ProductionCost
	err := r.db.WithContext(ctx).Where("order_id = ?", orderID).Order("id DESC").Find(&list).Error
	return list, err
}

func (r *ProductionCostRepository) Page(ctx context.Context, tenantID int64, req *model.ProductionCostQuery) ([]model.ProductionCost, int64, error) {
	var list []model.ProductionCost
	query := r.db.WithContext(ctx).Model(&model.ProductionCost{}).Where("tenant_id = ?", tenantID)
	if req.OrderID > 0 {
		query = query.Where("order_id = ?", req.OrderID)
	}
	if req.OrderNo != "" {
		query = query.Where("order_no LIKE ?", "%"+req.OrderNo+"%")
	}
	if req.CostType != "" {
		query = query.Where("cost_type = ?", req.CostType)
	}
	if req.StartDate != "" {
		query = query.Where("cost_date >= ?", req.StartDate)
	}
	if req.EndDate != "" {
		query = query.Where("cost_date <= ?", req.EndDate)
	}
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	err = query.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *ProductionCostRepository) GetSummaryByOrderID(ctx context.Context, orderID int64) (*model.ProductionCostSummary, error) {
	var summary model.ProductionCostSummary
	summary.OrderID = orderID

	// 获取工单信息
	var order model.ProductionOrder
	if err := r.db.WithContext(ctx).Where("id = ?", orderID).First(&order).Error; err != nil {
		return nil, err
	}
	summary.OrderNo = order.OrderNo
	summary.CompletedQty = order.CompletedQty

	// 材料成本
	var materialAmount float64
	r.db.WithContext(ctx).Model(&model.ProductionCost{}).
		Where("order_id = ? AND cost_type = 'material'", orderID).
		Select("COALESCE(SUM(amount), 0)").Scan(&materialAmount)
	summary.MaterialCost = materialAmount

	// 人工成本
	var laborAmount float64
	r.db.WithContext(ctx).Model(&model.ProductionCost{}).
		Where("order_id = ? AND cost_type = 'labor'", orderID).
		Select("COALESCE(SUM(amount), 0)").Scan(&laborAmount)
	summary.LaborCost = laborAmount

	// 制造费用
	var overheadAmount float64
	r.db.WithContext(ctx).Model(&model.ProductionCost{}).
		Where("order_id = ? AND cost_type = 'overhead'", orderID).
		Select("COALESCE(SUM(amount), 0)").Scan(&overheadAmount)
	summary.OverheadCost = overheadAmount

	// 总成本
	summary.TotalCost = summary.MaterialCost + summary.LaborCost + summary.OverheadCost

	// 单位成本
	if summary.CompletedQty > 0 {
		summary.UnitCost = summary.TotalCost / summary.CompletedQty
	}

	return &summary, nil
}

func (r *ProductionCostRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.ProductionCost{}, id).Error
}