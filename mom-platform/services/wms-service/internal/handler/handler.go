package handler

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common "github.com/ninghonggang/mom-platform/gen/common"
	wms "github.com/ninghonggang/mom-platform/gen/wms"

	"mom-platform/services/wms-service/internal/model"
	"mom-platform/services/wms-service/internal/service"
)

// Handler implements the wms.WmsServiceServer interface.
type Handler struct {
	wms.UnimplementedWmsServiceServer
	svc *service.Service
	log *zap.Logger
}

// New creates a new Handler.
func New(svc *service.Service, log *zap.Logger) *Handler {
	return &Handler{svc: svc, log: log}
}

// defaultTenantID returns a default tenant ID (in production, extract from gRPC metadata/auth).
const defaultTenantID uint = 1

// ---------------------------------------------------------------------------
// Warehouse methods
// ---------------------------------------------------------------------------

func (h *Handler) CreateWarehouse(ctx context.Context, req *wms.CreateWarehouseRequest) (*wms.Warehouse, error) {
	w := &model.Warehouse{
		TenantID:      defaultTenantID,
		WarehouseCode: req.GetWarehouseCode(),
		WarehouseName: req.GetWarehouseName(),
		WarehouseType: req.GetWarehouseType(),
		Status:        model.WarehouseStatusActive,
	}
	result, err := h.svc.CreateWarehouse(ctx, w)
	if err != nil {
		h.log.Error("create warehouse failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "create warehouse: %v", err)
	}
	return warehouseToProto(result), nil
}

func (h *Handler) ListWarehouses(ctx context.Context, req *wms.ListWarehousesRequest) (*wms.ListWarehousesResponse, error) {
	list, err := h.svc.ListWarehouses(ctx, defaultTenantID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list warehouses: %v", err)
	}
	return &wms.ListWarehousesResponse{
		Items: warehousesToProto(list),
		Page:  pageResponseFromTotal(req.GetPage(), len(list)),
	}, nil
}

// ---------------------------------------------------------------------------
// Location methods
// ---------------------------------------------------------------------------

func (h *Handler) CreateLocation(ctx context.Context, req *wms.CreateLocationRequest) (*wms.Location, error) {
	l := &model.Location{
		WarehouseID:  uint(req.GetWarehouseId()),
		AreaID:       uint(req.GetAreaId()),
		LocationCode: req.GetLocationCode(),
		LocationType: protoToLocationType(req.GetLocationType()),
		Capacity:     strToFloat64(req.GetCapacity()),
		Status:       model.LocationStatusActive,
	}
	result, err := h.svc.CreateLocation(ctx, l)
	if err != nil {
		h.log.Error("create location failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "create location: %v", err)
	}
	return locationToProto(result), nil
}

func (h *Handler) ListLocations(ctx context.Context, req *wms.ListLocationsRequest) (*wms.ListLocationsResponse, error) {
	list, err := h.svc.ListLocations(ctx, uint(req.GetWarehouseId()), uint(req.GetAreaId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list locations: %v", err)
	}
	return &wms.ListLocationsResponse{
		Items: locationsToProto(list),
		Page:  pageResponseFromTotal(req.GetPage(), len(list)),
	}, nil
}

// ---------------------------------------------------------------------------
// Inventory methods
// ---------------------------------------------------------------------------

func (h *Handler) GetBalance(ctx context.Context, req *wms.GetBalanceRequest) (*wms.InventoryBalance, error) {
	ib, err := h.svc.GetBalance(ctx, defaultTenantID, uint(req.GetMaterialId()), uint(req.GetLocationId()), req.GetBatchNo())
	if err != nil {
		h.log.Error("get balance failed", zap.Error(err))
		return nil, status.Errorf(codes.NotFound, "get balance: %v", err)
	}
	return inventoryBalanceToProto(ib), nil
}

func (h *Handler) ListBalances(ctx context.Context, req *wms.ListBalancesRequest) (*wms.ListBalancesResponse, error) {
	list, err := h.svc.ListBalances(ctx, defaultTenantID, uint(req.GetMaterialId()), balanceStatusFromProto(req.GetStatus()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list balances: %v", err)
	}
	return &wms.ListBalancesResponse{
		Items: inventoryBalancesToProto(list),
		Page:  pageResponseFromTotal(req.GetPage(), len(list)),
	}, nil
}

func (h *Handler) LockInventory(ctx context.Context, req *wms.LockInventoryRequest) (*wms.InventoryBalance, error) {
	ib, err := h.svc.LockInventory(ctx, uint(req.GetBalanceId()), strToFloat64(req.GetLockQty()))
	if err != nil {
		h.log.Error("lock inventory failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "lock inventory: %v", err)
	}
	return inventoryBalanceToProto(ib), nil
}

func (h *Handler) UnlockInventory(ctx context.Context, req *wms.UnlockInventoryRequest) (*wms.InventoryBalance, error) {
	ib, err := h.svc.UnlockInventory(ctx, uint(req.GetLockId()), 0)
	if err != nil {
		h.log.Error("unlock inventory failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "unlock inventory: %v", err)
	}
	return inventoryBalanceToProto(ib), nil
}

// ---------------------------------------------------------------------------
// Receive order methods
// ---------------------------------------------------------------------------

func (h *Handler) CreateReceiveOrder(ctx context.Context, req *wms.CreateReceiveOrderRequest) (*wms.ReceiveOrder, error) {
	lines := make([]service.ReceiveLineInput, 0, len(req.GetLines()))
	for _, l := range req.GetLines() {
		lines = append(lines, service.ReceiveLineInput{
			MaterialID:  uint(l.GetMaterialId()),
			ExpectedQty: strToFloat64(l.GetExpectedQty()),
			UnitPrice:   strToFloat64(l.GetUnitPrice()),
		})
	}
	ro, err := h.svc.CreateReceiveOrder(ctx, &service.CreateReceiveOrderReq{
		TenantID:   defaultTenantID,
		PoID:       strconv.FormatInt(req.GetPoId(), 10),
		SupplierID: uint(req.GetSupplierId()),
		Lines:      lines,
	})
	if err != nil {
		h.log.Error("create receive order failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "create receive order: %v", err)
	}
	return receiveOrderToProto(ro, nil), nil
}

func (h *Handler) ReceiveConfirm(ctx context.Context, req *wms.ReceiveConfirmRequest) (*wms.ReceiveOrder, error) {
	lines := make([]service.ReceiveLineConfirm, 0, len(req.GetLines()))
	for _, l := range req.GetLines() {
		lines = append(lines, service.ReceiveLineConfirm{
			LineID:      uint(l.GetLineId()),
			ReceivedQty: strToFloat64(l.GetReceivedQty()),
			BatchNo:     l.GetBatchNo(),
			ExpireDate:  parseExpireDate(l.GetExpireDate()),
		})
	}
	ro, err := h.svc.ReceiveConfirm(ctx, uint(req.GetId()), lines)
	if err != nil {
		h.log.Error("receive confirm failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "receive confirm: %v", err)
	}
	return receiveOrderToProto(ro, nil), nil
}

func (h *Handler) Putaway(ctx context.Context, req *wms.PutawayRequest) (*wms.ReceiveOrder, error) {
	items := make([]service.PutawayItem, 0, len(req.GetLines()))
	for _, l := range req.GetLines() {
		items = append(items, service.PutawayItem{
			LineID:     uint(l.GetLineId()),
			LocationID: uint(l.GetLocationId()),
			Quantity:   strToFloat64(l.GetQuantity()),
		})
	}
	ro, err := h.svc.Putaway(ctx, uint(req.GetReceiveOrderId()), items)
	if err != nil {
		h.log.Error("putaway failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "putaway: %v", err)
	}
	return receiveOrderToProto(ro, nil), nil
}

func (h *Handler) GetReceiveOrder(ctx context.Context, req *common.IdRequest) (*wms.ReceiveOrder, error) {
	ro, lines, err := h.svc.GetReceiveOrder(ctx, uint(req.GetId()))
	if err != nil {
		h.log.Error("get receive order failed", zap.Error(err))
		return nil, status.Errorf(codes.NotFound, "get receive order: %v", err)
	}
	return receiveOrderToProto(ro, lines), nil
}

func (h *Handler) ListReceiveOrders(ctx context.Context, req *wms.ListReceiveOrdersRequest) (*wms.ListReceiveOrdersResponse, error) {
	list, err := h.svc.ListReceiveOrders(ctx, receiveOrderStatusFromProto(req.GetStatus()), uint(req.GetSupplierId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list receive orders: %v", err)
	}
	items := make([]*wms.ReceiveOrder, 0, len(list))
	for _, ro := range list {
		items = append(items, receiveOrderToProto(ro, nil))
	}
	return &wms.ListReceiveOrdersResponse{
		Items: items,
		Page:  pageResponseFromTotal(req.GetPage(), len(list)),
	}, nil
}

// ---------------------------------------------------------------------------
// Delivery order methods
// ---------------------------------------------------------------------------

func (h *Handler) CreateDeliveryOrder(ctx context.Context, req *wms.CreateDeliveryOrderRequest) (*wms.DeliveryOrder, error) {
	lines := make([]service.DeliveryLineInput, 0, len(req.GetLines()))
	for _, l := range req.GetLines() {
		lines = append(lines, service.DeliveryLineInput{
			MaterialID: uint(l.GetMaterialId()),
			OrderedQty: strToFloat64(l.GetOrderedQty()),
		})
	}
	do, err := h.svc.CreateDeliveryOrder(ctx, &service.CreateDeliveryOrderReq{
		TenantID:   defaultTenantID,
		SoID:       strconv.FormatInt(req.GetSoId(), 10),
		CustomerID: uint(req.GetCustomerId()),
		Lines:      lines,
	})
	if err != nil {
		h.log.Error("create delivery order failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "create delivery order: %v", err)
	}
	return deliveryOrderToProto(do, nil), nil
}

func (h *Handler) PickItems(ctx context.Context, req *wms.PickRequest) (*wms.DeliveryOrder, error) {
	picks := make([]service.PickItem, 0, len(req.GetLines()))
	for _, l := range req.GetLines() {
		picks = append(picks, service.PickItem{
			LineID:     uint(l.GetLineId()),
			LocationID: uint(l.GetLocationId()),
			Quantity:   strToFloat64(l.GetPickedQty()),
			PickerID:   uint(l.GetPickerId()),
		})
	}
	do, err := h.svc.PickItems(ctx, uint(req.GetDeliveryOrderId()), picks)
	if err != nil {
		h.log.Error("pick items failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "pick items: %v", err)
	}
	return deliveryOrderToProto(do, nil), nil
}

func (h *Handler) ShipOrder(ctx context.Context, req *wms.ShipRequest) (*wms.DeliveryOrder, error) {
	do, err := h.svc.ShipOrder(ctx, uint(req.GetId()))
	if err != nil {
		h.log.Error("ship order failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "ship order: %v", err)
	}
	return deliveryOrderToProto(do, nil), nil
}

func (h *Handler) GetDeliveryOrder(ctx context.Context, req *common.IdRequest) (*wms.DeliveryOrder, error) {
	do, lines, err := h.svc.GetDeliveryOrder(ctx, uint(req.GetId()))
	if err != nil {
		h.log.Error("get delivery order failed", zap.Error(err))
		return nil, status.Errorf(codes.NotFound, "get delivery order: %v", err)
	}
	return deliveryOrderToProto(do, lines), nil
}

func (h *Handler) ListDeliveryOrders(ctx context.Context, req *wms.ListDeliveryOrdersRequest) (*wms.ListDeliveryOrdersResponse, error) {
	list, err := h.svc.ListDeliveryOrders(ctx, deliveryOrderStatusFromProto(req.GetStatus()), uint(req.GetCustomerId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list delivery orders: %v", err)
	}
	items := make([]*wms.DeliveryOrder, 0, len(list))
	for _, do := range list {
		items = append(items, deliveryOrderToProto(do, nil))
	}
	return &wms.ListDeliveryOrdersResponse{
		Items: items,
		Page:  pageResponseFromTotal(req.GetPage(), len(list)),
	}, nil
}

// ---------------------------------------------------------------------------
// Count plan methods
// ---------------------------------------------------------------------------

func (h *Handler) CreateCountPlan(ctx context.Context, req *wms.CreateCountPlanRequest) (*wms.CountPlan, error) {
	cp, err := h.svc.CreateCountPlan(ctx, uint(req.GetWarehouseId()), req.GetPlanType())
	if err != nil {
		h.log.Error("create count plan failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "create count plan: %v", err)
	}
	return countPlanToProto(cp), nil
}

func (h *Handler) SubmitCountRecord(ctx context.Context, req *wms.SubmitCountRecordRequest) (*wms.CountPlan, error) {
	items := make([]service.CountRecordItem, 0, len(req.GetItems()))
	for _, item := range req.GetItems() {
		items = append(items, service.CountRecordItem{
			MaterialID: uint(item.GetMaterialId()),
			LocationID: uint(item.GetLocationId()),
			ActualQty:  strToFloat64(item.GetActualQty()),
		})
	}
	cp, err := h.svc.SubmitCountRecord(ctx, &service.SubmitCountRecordReq{
		PlanID: uint(req.GetPlanId()),
		Items:  items,
	})
	if err != nil {
		h.log.Error("submit count record failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "submit count record: %v", err)
	}
	return countPlanToProto(cp), nil
}
