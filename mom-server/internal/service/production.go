package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type ProductionOrderService struct {
	repo *repository.ProductionOrderRepository
}

func NewProductionOrderService(repo *repository.ProductionOrderRepository) *ProductionOrderService {
	return &ProductionOrderService{repo: repo}
}

func (s *ProductionOrderService) List(ctx context.Context) ([]model.ProductionOrder, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *ProductionOrderService) GetByID(ctx context.Context, id string) (*model.ProductionOrder, error) {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, orderID)
}

func (s *ProductionOrderService) Create(ctx context.Context, order *model.ProductionOrder) error {
	return s.repo.Create(ctx, order)
}

func (s *ProductionOrderService) Update(ctx context.Context, id string, order *model.ProductionOrder) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, orderID, map[string]interface{}{
		"order_no":      order.OrderNo,
		"material_id":   order.MaterialID,
		"quantity":      order.Quantity,
		"workshop_id":   order.WorkshopID,
		"line_id":       order.LineID,
		"status":        order.Status,
	})
}

func (s *ProductionOrderService) Delete(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, orderID)
}

func (s *ProductionOrderService) Start(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, orderID, map[string]interface{}{
		"status": 2, // 生产中
	})
}

func (s *ProductionOrderService) Complete(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, orderID, map[string]interface{}{
		"status": 3, // 已完成
	})
}
