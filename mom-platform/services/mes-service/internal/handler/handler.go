package handler

import (
	"context"
	"errors"
	"strings"

	"github.com/ninghonggang/mom-platform/gen/common"
	mes "github.com/ninghonggang/mom-platform/gen/mes"
	"github.com/ninghonggang/mom-platform/services/mes-service/internal/model"
	"github.com/ninghonggang/mom-platform/services/mes-service/internal/repository"
	"github.com/ninghonggang/mom-platform/services/mes-service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

// MESHandler implements all 4 MES gRPC server interfaces.
// Embeds Unimplemented*Server types for forward compatibility.
type MESHandler struct {
	mes.UnimplementedProductionOrderServiceServer
	mes.UnimplementedDispatchServiceServer
	mes.UnimplementedJobReportServiceServer
	mes.UnimplementedProductionCompleteServiceServer

	orderService    *service.OrderService
	reportService   *service.ReportService
	dispatchService *service.DispatchService
	completeService *service.CompleteService
}

func NewMESHandler(
	orderService *service.OrderService,
	reportService *service.ReportService,
	dispatchService *service.DispatchService,
	completeService *service.CompleteService,
) *MESHandler {
	return &MESHandler{
		orderService:    orderService,
		reportService:   reportService,
		dispatchService: dispatchService,
		completeService: completeService,
	}
}

// --- ProductionOrderService ---

func (h *MESHandler) GetOrder(ctx context.Context, req *mes.GetOrderRequest) (*mes.GetOrderResponse, error) {
	order, err := h.orderService.Get(ctx, req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "order %d not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "get order: %v", err)
	}
	return &mes.GetOrderResponse{Order: orderToProto(order)}, nil
}

func (h *MESHandler) ListOrders(ctx context.Context, req *mes.ListOrdersRequest) (*mes.ListOrdersResponse, error) {
	page, pageSize := paginationFromProto(req.Pagination)
	var statusFilter *model.OrderStatus
	if req.Status != common.ProductionOrderStatus_PRODUCTION_ORDER_STATUS_UNSPECIFIED {
		s := model.OrderStatus(req.Status)
		statusFilter = &s
	}
	orders, total, err := h.orderService.List(ctx, repository.OrderFilter{
		TenantID:   req.TenantId,
		Keyword:    req.Keyword,
		WorkshopID: req.WorkshopId,
		LineID:     req.LineId,
		Status:     statusFilter,
		DateFrom:   req.DateFrom,
		DateTo:     req.DateTo,
		Page:       page,
		PageSize:   pageSize,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list orders: %v", err)
	}
	items := make([]*mes.ProductionOrder, 0, len(orders))
	for i := range orders {
		items = append(items, orderToProto(&orders[i]))
	}
	return &mes.ListOrdersResponse{
		Items:      items,
		Pagination: paginationToProto(total, page, pageSize),
	}, nil
}

func (h *MESHandler) CreateOrder(ctx context.Context, req *mes.CreateOrderRequest) (*mes.CreateOrderResponse, error) {
	if req.Order == nil {
		return nil, status.Error(codes.InvalidArgument, "order is required")
	}
	order := orderFromProto(req.Order)
	created, err := h.orderService.Create(ctx, order)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "create order: %v", err)
	}
	return &mes.CreateOrderResponse{Order: orderToProto(created)}, nil
}

func (h *MESHandler) UpdateOrderStatus(ctx context.Context, req *mes.UpdateOrderStatusRequest) (*mes.UpdateOrderStatusResponse, error) {
	if req.Status == common.ProductionOrderStatus_PRODUCTION_ORDER_STATUS_UNSPECIFIED {
		return nil, status.Error(codes.InvalidArgument, "status is required")
	}
	if err := h.orderService.UpdateStatus(ctx, req.Id, model.OrderStatus(req.Status)); err != nil {
		return nil, status.Errorf(codes.Internal, "update order status: %v", err)
	}
	order, err := h.orderService.Get(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get order after update: %v", err)
	}
	return &mes.UpdateOrderStatusResponse{Order: orderToProto(order)}, nil
}

func (h *MESHandler) HoldOrder(ctx context.Context, req *mes.HoldOrderRequest) (*mes.UpdateOrderStatusResponse, error) {
	if err := h.orderService.Hold(ctx, req.Id); err != nil {
		return nil, status.Errorf(codes.Internal, "hold order: %v", err)
	}
	order, err := h.orderService.Get(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get order after hold: %v", err)
	}
	return &mes.UpdateOrderStatusResponse{Order: orderToProto(order)}, nil
}

func (h *MESHandler) ResumeOrder(ctx context.Context, req *mes.ResumeOrderRequest) (*mes.UpdateOrderStatusResponse, error) {
	if err := h.orderService.Resume(ctx, req.Id); err != nil {
		return nil, status.Errorf(codes.Internal, "resume order: %v", err)
	}
	order, err := h.orderService.Get(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get order after resume: %v", err)
	}
	return &mes.UpdateOrderStatusResponse{Order: orderToProto(order)}, nil
}

// --- DispatchService ---

func (h *MESHandler) ListDispatches(ctx context.Context, req *mes.ListDispatchesRequest) (*mes.ListDispatchesResponse, error) {
	dispatches, err := h.dispatchService.List(ctx, repository.DispatchFilter{
		TenantID:      req.TenantId,
		OrderID:       req.OrderId,
		LineID:        req.LineId,
		WorkstationID: req.WorkstationId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list dispatches: %v", err)
	}
	items := make([]*mes.Dispatch, 0, len(dispatches))
	for i := range dispatches {
		items = append(items, dispatchToProto(&dispatches[i]))
	}
	return &mes.ListDispatchesResponse{Items: items}, nil
}

func (h *MESHandler) CreateDispatch(ctx context.Context, req *mes.CreateDispatchRequest) (*mes.CreateDispatchResponse, error) {
	if len(req.Dispatches) == 0 {
		return nil, status.Error(codes.InvalidArgument, "dispatches is required")
	}
	dispatches := make([]model.Dispatch, 0, len(req.Dispatches))
	for _, d := range req.Dispatches {
		dispatches = append(dispatches, *dispatchFromProto(d))
	}
	created, err := h.dispatchService.CreateBatch(ctx, req.OrderId, dispatches)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "create dispatches: %v", err)
	}
	items := make([]*mes.Dispatch, 0, len(created))
	for i := range created {
		items = append(items, dispatchToProto(&created[i]))
	}
	return &mes.CreateDispatchResponse{Dispatches: items}, nil
}

// --- JobReportService ---

func (h *MESHandler) CreateReport(ctx context.Context, req *mes.CreateJobReportRequest) (*mes.CreateJobReportResponse, error) {
	if req.Report == nil {
		return nil, status.Error(codes.InvalidArgument, "report is required")
	}
	report := reportFromProto(req.Report)
	created, err := h.reportService.Create(ctx, report)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "create report: %v", err)
	}
	return &mes.CreateJobReportResponse{Report: reportToProto(created)}, nil
}

func (h *MESHandler) ListReports(ctx context.Context, req *mes.ListJobReportsRequest) (*mes.ListJobReportsResponse, error) {
	page, pageSize := paginationFromProto(req.Pagination)
	reports, total, err := h.reportService.List(ctx, repository.ReportFilter{
		TenantID:   req.TenantId,
		OrderID:    req.OrderId,
		EmployeeID: req.EmployeeId,
		DateFrom:   req.DateFrom,
		DateTo:     req.DateTo,
		Page:       page,
		PageSize:   pageSize,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list reports: %v", err)
	}
	items := make([]*mes.MobileJobReport, 0, len(reports))
	for i := range reports {
		items = append(items, reportToProto(&reports[i]))
	}
	return &mes.ListJobReportsResponse{
		Items:      items,
		Pagination: paginationToProto(total, page, pageSize),
	}, nil
}

func (h *MESHandler) ConfirmReport(ctx context.Context, req *mes.ConfirmReportRequest) (*mes.ConfirmReportResponse, error) {
	if err := h.reportService.Confirm(ctx, req.ReportId); err != nil {
		return nil, status.Errorf(codes.Internal, "confirm report: %v", err)
	}
	report, err := h.reportService.Get(ctx, req.ReportId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get report after confirm: %v", err)
	}
	return &mes.ConfirmReportResponse{Report: reportToProto(report)}, nil
}

func (h *MESHandler) AuditReport(ctx context.Context, req *mes.AuditReportRequest) (*mes.AuditReportResponse, error) {
	if err := h.reportService.Audit(ctx, req.ReportId); err != nil {
		return nil, status.Errorf(codes.Internal, "audit report: %v", err)
	}
	report, err := h.reportService.Get(ctx, req.ReportId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get report after audit: %v", err)
	}
	return &mes.AuditReportResponse{Report: reportToProto(report)}, nil
}

// --- ProductionCompleteService ---

// CreateComplete records a finished-goods completion for a production order.
// It validates the quantity against the order plan, persists the completion,
// rolls the accumulated quantity back onto the order and — once the planned
// quantity is reached — closes the order and emits mes.order.completed.
func (h *MESHandler) CreateComplete(ctx context.Context, req *mes.CreateCompleteRequest) (*mes.CreateCompleteResponse, error) {
	in := req.GetComplete()
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "complete is required")
	}

	c := &model.ProductionComplete{
		OrderID:     in.GetOrderId(),
		WarehouseID: in.GetWarehouseId(),
		LocationID:  in.GetLocationId(),
		Quantity:    in.GetQuantity(),
		BatchNo:     in.GetBatchNo(),
	}
	if ts := in.GetCompleteTime(); ts != nil && ts.IsValid() {
		t := ts.AsTime()
		c.CompleteTime = &t
	}

	created, err := h.completeService.Create(ctx, c)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "order %d not found", in.GetOrderId())
		}
		if strings.Contains(err.Error(), "is required") ||
			strings.Contains(err.Error(), "must be positive") {
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		}
		if strings.Contains(err.Error(), "exceeds planned qty") ||
			strings.Contains(err.Error(), "is closed") {
			return nil, status.Errorf(codes.FailedPrecondition, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "create complete: %v", err)
	}
	return &mes.CreateCompleteResponse{Complete: completeToProto(created)}, nil
}

// completeToProto maps the persistence model onto the wire representation.
func completeToProto(c *model.ProductionComplete) *mes.ProductionComplete {
	out := &mes.ProductionComplete{
		Base: &common.BaseModel{
			Id:       c.ID,
			TenantId: c.TenantID,
		},
		OrderId:     c.OrderID,
		OrderNo:     c.OrderNo,
		WarehouseId: c.WarehouseID,
		LocationId:  c.LocationID,
		Quantity:    c.Quantity,
		BatchNo:     c.BatchNo,
	}
	if c.CompleteTime != nil {
		out.CompleteTime = timestamppb.New(*c.CompleteTime)
	}
	return out
}
