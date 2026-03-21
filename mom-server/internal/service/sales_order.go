package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type SalesOrderService struct {
	repo *repository.SalesOrderRepository
}

func NewSalesOrderService(repo *repository.SalesOrderRepository) *SalesOrderService {
	return &SalesOrderService{repo: repo}
}

func (s *SalesOrderService) List(ctx context.Context) ([]model.SalesOrder, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *SalesOrderService) GetByID(ctx context.Context, id string) (*model.SalesOrder, error) {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, orderID)
}

func (s *SalesOrderService) Create(ctx context.Context, order *model.SalesOrder) error {
	return s.repo.Create(ctx, order)
}

func (s *SalesOrderService) Update(ctx context.Context, id string, order *model.SalesOrder) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, orderID, map[string]interface{}{
		"customer_id":   order.CustomerID,
		"customer_name": order.CustomerName,
		"order_date":    order.OrderDate,
		"delivery_date": order.DeliveryDate,
		"order_type":    order.OrderType,
		"priority":      order.Priority,
		"status":        order.Status,
		"remark":        order.Remark,
	})
}

func (s *SalesOrderService) Delete(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, orderID)
}

func (s *SalesOrderService) Confirm(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, orderID, map[string]interface{}{
		"status": 2,
	})
}
