package service

import (
	"context"
	"fmt"
	"time"

	"github.com/ninghonggang/mom-platform/pkg/eventbus"
	"github.com/ninghonggang/mom-platform/services/mes-service/internal/model"
	"github.com/ninghonggang/mom-platform/services/mes-service/internal/repository"
)

// OrderService — 工单业务逻辑
type OrderService struct {
	repo      repository.OrderRepository
	publisher *eventbus.EventPublisher
}

func NewOrderService(repo repository.OrderRepository, publisher *eventbus.EventPublisher) *OrderService {
	return &OrderService{repo: repo, publisher: publisher}
}

func (s *OrderService) Get(ctx context.Context, id int64) (*model.ProductionOrder, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *OrderService) List(ctx context.Context, filter repository.OrderFilter) ([]model.ProductionOrder, int64, error) {
	return s.repo.List(ctx, filter)
}

func (s *OrderService) Create(ctx context.Context, order *model.ProductionOrder) (*model.ProductionOrder, error) {
	if order.OrderNo == "" {
		return nil, fmt.Errorf("order_no is required")
	}
	if order.MaterialID == 0 {
		return nil, fmt.Errorf("material_id is required")
	}
	if order.Quantity <= 0 {
		return nil, fmt.Errorf("quantity must be positive")
	}
	order.Status = model.OrderPending
	if err := s.repo.Create(ctx, order); err != nil {
		return nil, fmt.Errorf("create order: %w", err)
	}
	if s.publisher != nil {
		_ = s.publisher.Publish(ctx, eventbus.SubjectMESOrderCreated, order)
	}
	return order, nil
}

func (s *OrderService) UpdateStatus(ctx context.Context, id int64, status model.OrderStatus) error {
	// Validate state transition
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("get order: %w", err)
	}
	if !isValidTransition(existing.Status, status) {
		return fmt.Errorf("invalid transition from %d to %d", existing.Status, status)
	}
	if err := s.repo.UpdateStatus(ctx, id, status); err != nil {
		return fmt.Errorf("update status: %w", err)
	}
	if s.publisher != nil && status == model.OrderCompleted {
		_ = s.publisher.Publish(ctx, eventbus.SubjectMESOrderCompleted, map[string]interface{}{
			"order_id": id,
			"order_no": existing.OrderNo,
		})
	}
	return nil
}

func (s *OrderService) Hold(ctx context.Context, id int64) error {
	if err := s.UpdateStatus(ctx, id, model.OrderSuspended); err != nil {
		return err
	}
	if s.publisher != nil {
		_ = s.publisher.Publish(ctx, eventbus.SubjectMESOrderHold, map[string]interface{}{
			"order_id": id,
		})
	}
	return nil
}

func (s *OrderService) Resume(ctx context.Context, id int64) error {
	return s.UpdateStatus(ctx, id, model.OrderInProgress)
}

func isValidTransition(from, to model.OrderStatus) bool {
	valid := map[model.OrderStatus][]model.OrderStatus{
		model.OrderPending:    {model.OrderInProgress, model.OrderClosed, model.OrderSuspended},
		model.OrderInProgress: {model.OrderCompleted, model.OrderSuspended, model.OrderClosed},
		model.OrderSuspended:  {model.OrderInProgress, model.OrderClosed},
		model.OrderCompleted:  {},
		model.OrderClosed:     {},
	}
	for _, t := range valid[from] {
		if t == to {
			return true
		}
	}
	return false
}

// ReportService — 报工业务逻辑
type ReportService struct {
	repo      repository.JobReportRepository
	orderRepo repository.OrderRepository
}

func NewReportService(repo repository.JobReportRepository, orderRepo repository.OrderRepository) *ReportService {
	return &ReportService{repo: repo, orderRepo: orderRepo}
}

func (s *ReportService) Get(ctx context.Context, id int64) (*model.MobileJobReport, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ReportService) Create(ctx context.Context, report *model.MobileJobReport) (*model.MobileJobReport, error) {
	if report.OrderID == 0 {
		return nil, fmt.Errorf("order_id is required")
	}
	report.Status = model.ReportSubmitted
	if err := s.repo.Create(ctx, report); err != nil {
		return nil, fmt.Errorf("create report: %w", err)
	}
	return report, nil
}

func (s *ReportService) List(ctx context.Context, filter repository.ReportFilter) ([]model.MobileJobReport, int64, error) {
	return s.repo.List(ctx, filter)
}

func (s *ReportService) Confirm(ctx context.Context, id int64) error {
	return s.repo.UpdateStatus(ctx, id, model.ReportConfirmed)
}

func (s *ReportService) Audit(ctx context.Context, id int64) error {
	return s.repo.UpdateStatus(ctx, id, model.ReportAudited)
}

// DispatchService — 派工业务逻辑
type DispatchService struct {
	repo      repository.DispatchRepository
	orderRepo repository.OrderRepository
}

func NewDispatchService(repo repository.DispatchRepository, orderRepo repository.OrderRepository) *DispatchService {
	return &DispatchService{repo: repo, orderRepo: orderRepo}
}

func (s *DispatchService) List(ctx context.Context, filter repository.DispatchFilter) ([]model.Dispatch, error) {
	return s.repo.List(ctx, filter)
}

func (s *DispatchService) CreateBatch(ctx context.Context, orderID int64, dispatches []model.Dispatch) ([]model.Dispatch, error) {
	for i := range dispatches {
		dispatches[i].OrderID = orderID
		dispatches[i].Status = model.DispatchDispatched
	}
	if err := s.repo.CreateBatch(ctx, dispatches); err != nil {
		return nil, fmt.Errorf("create dispatches: %w", err)
	}
	return dispatches, nil
}

// CompleteService — 完工入库业务逻辑
//
// 完工是工单生命周期的收口动作：校验数量不超过工单计划量，
// 落库完工记录，回写工单已完成数量，达到计划量时把工单流转为 COMPLETED
// 并发出 mes.order.completed 事件供 WMS/TRACE 等下游消费。
type CompleteService struct {
	repo      repository.CompleteRepository
	orderRepo repository.OrderRepository
	publisher *eventbus.EventPublisher
}

func NewCompleteService(
	repo repository.CompleteRepository,
	orderRepo repository.OrderRepository,
	publisher *eventbus.EventPublisher,
) *CompleteService {
	return &CompleteService{repo: repo, orderRepo: orderRepo, publisher: publisher}
}

func (s *CompleteService) Get(ctx context.Context, id int64) (*model.ProductionComplete, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CompleteService) ListByOrder(ctx context.Context, orderID int64) ([]model.ProductionComplete, error) {
	return s.repo.ListByOrder(ctx, orderID)
}

// Create records a production completion (finished-goods receipt) for an order.
func (s *CompleteService) Create(ctx context.Context, c *model.ProductionComplete) (*model.ProductionComplete, error) {
	if c.OrderID == 0 {
		return nil, fmt.Errorf("order_id is required")
	}
	if c.Quantity <= 0 {
		return nil, fmt.Errorf("quantity must be positive")
	}

	order, err := s.orderRepo.GetByID(ctx, c.OrderID)
	if err != nil {
		return nil, fmt.Errorf("get order: %w", err)
	}
	if order.Status == model.OrderClosed {
		return nil, fmt.Errorf("order %d is closed", c.OrderID)
	}

	done, err := s.repo.SumQtyByOrder(ctx, c.OrderID)
	if err != nil {
		return nil, fmt.Errorf("sum completed qty: %w", err)
	}
	if done+c.Quantity > order.Quantity {
		return nil, fmt.Errorf(
			"completed qty %.4f exceeds planned qty %.4f (already completed %.4f)",
			done+c.Quantity, order.Quantity, done)
	}

	c.TenantID = order.TenantID
	c.OrderNo = order.OrderNo
	if c.CompleteTime == nil {
		now := time.Now()
		c.CompleteTime = &now
	}
	if err := s.repo.Create(ctx, c); err != nil {
		return nil, fmt.Errorf("create complete: %w", err)
	}

	// Write back the accumulated completed quantity onto the order.
	total := done + c.Quantity
	order.CompletedQty = total
	if total >= order.Quantity {
		order.Status = model.OrderCompleted
		order.ActualEndDate = c.CompleteTime
	}
	if err := s.orderRepo.Update(ctx, order); err != nil {
		return nil, fmt.Errorf("update order after complete: %w", err)
	}

	if s.publisher != nil && order.Status == model.OrderCompleted {
		_ = s.publisher.Publish(ctx, eventbus.SubjectMESOrderCompleted, map[string]interface{}{
			"order_id":     order.ID,
			"order_no":     order.OrderNo,
			"material_id":  order.MaterialID,
			"quantity":     total,
			"warehouse_id": c.WarehouseID,
			"location_id":  c.LocationID,
			"batch_no":     c.BatchNo,
		})
	}
	return c, nil
}
