package service

import (
	"context"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"time"
)

type ProductionCostService struct {
	repo         *repository.ProductionCostRepository
	orderRepo    *repository.ProductionOrderRepository
}

func NewProductionCostService(repo *repository.ProductionCostRepository, orderRepo *repository.ProductionOrderRepository) *ProductionCostService {
	return &ProductionCostService{repo: repo, orderRepo: orderRepo}
}

func (s *ProductionCostService) Create(ctx context.Context, tenantID int64, req *model.ProductionCostCreateReq) (*model.ProductionCost, error) {
	costDate := time.Now()
	if req.CostDate != nil && *req.CostDate != "" {
		if parsed, err := time.Parse("2006-01-02", *req.CostDate); err == nil {
			costDate = parsed
		}
	}

	cost := &model.ProductionCost{
		TenantID:       tenantID,
		OrderID:        req.OrderID,
		OrderNo:        req.OrderNo,
		CostType:       req.CostType,
		CostItem:       req.CostItem,
		Quantity:       req.Quantity,
		UnitPrice:      req.UnitPrice,
		Amount:         req.Amount,
		ProcessID:      req.ProcessID,
		ProcessName:    req.ProcessName,
		DepartmentID:   req.DepartmentID,
		DepartmentName: req.DepartmentName,
		WorkerID:       req.WorkerID,
		WorkerName:     req.WorkerName,
		EquipmentID:    req.EquipmentID,
		EquipmentName:  req.EquipmentName,
		CostDate:       &costDate,
		Remark:         req.Remark,
	}

	if err := s.repo.Create(ctx, cost); err != nil {
		return nil, err
	}
	return cost, nil
}

func (s *ProductionCostService) List(ctx context.Context, tenantID int64, req *model.ProductionCostQuery) ([]model.ProductionCost, int64, error) {
	return s.repo.Page(ctx, tenantID, req)
}

func (s *ProductionCostService) GetSummary(ctx context.Context, orderID int64) (*model.ProductionCostSummary, error) {
	return s.repo.GetSummaryByOrderID(ctx, orderID)
}

func (s *ProductionCostService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}