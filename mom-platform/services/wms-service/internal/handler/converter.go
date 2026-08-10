package handler

import (
	"strconv"
	"time"

	common "github.com/ninghonggang/mom-platform/gen/common"
	wms "github.com/ninghonggang/mom-platform/gen/wms"

	"mom-platform/services/wms-service/internal/model"
)

// --- Warehouse conversions ---

func warehouseToProto(w *model.Warehouse) *wms.Warehouse {
	if w == nil {
		return nil
	}
	return &wms.Warehouse{
		Id:            int64(w.ID),
		TenantId:      int64(w.TenantID),
		WarehouseCode: w.WarehouseCode,
		WarehouseName: w.WarehouseName,
		WarehouseType: w.WarehouseType,
		Status:        string(w.Status),
	}
}

func warehousesToProto(list []*model.Warehouse) []*wms.Warehouse {
	out := make([]*wms.Warehouse, 0, len(list))
	for _, w := range list {
		out = append(out, warehouseToProto(w))
	}
	return out
}

// --- Location conversions ---

func locationTypeToProto(t model.LocationType) wms.LocationType {
	switch t {
	case model.LocationTypePick:
		return wms.LocationType_LOCATION_TYPE_PICK
	case model.LocationTypeStorage:
		return wms.LocationType_LOCATION_TYPE_STORAGE
	case model.LocationTypeInbound:
		return wms.LocationType_LOCATION_TYPE_INBOUND
	case model.LocationTypeOutbound:
		return wms.LocationType_LOCATION_TYPE_OUTBOUND
	default:
		return wms.LocationType_LOCATION_TYPE_UNSPECIFIED
	}
}

// protoToLocationType maps a proto enum to the short form persisted in the
// database. The generated enum's String() carries the full
// "LOCATION_TYPE_STORAGE" prefix, which overflows the varchar(20) column —
// so the prefix must be stripped here rather than stored verbatim.
func protoToLocationType(t wms.LocationType) model.LocationType {
	switch t {
	case wms.LocationType_LOCATION_TYPE_PICK:
		return model.LocationTypePick
	case wms.LocationType_LOCATION_TYPE_INBOUND:
		return model.LocationTypeInbound
	case wms.LocationType_LOCATION_TYPE_OUTBOUND:
		return model.LocationTypeOutbound
	case wms.LocationType_LOCATION_TYPE_STORAGE:
		return model.LocationTypeStorage
	default:
		// Unspecified defaults to a normal storage bin.
		return model.LocationTypeStorage
	}
}

func locationStatusToProto(s model.LocationStatus) wms.LocationStatus {
	switch s {
	case model.LocationStatusActive:
		return wms.LocationStatus_LOCATION_STATUS_ACTIVE
	case model.LocationStatusInactive:
		return wms.LocationStatus_LOCATION_STATUS_INACTIVE
	case model.LocationStatusFull:
		return wms.LocationStatus_LOCATION_STATUS_FULL
	default:
		return wms.LocationStatus_LOCATION_STATUS_UNSPECIFIED
	}
}

func locationToProto(l *model.Location) *wms.Location {
	if l == nil {
		return nil
	}
	return &wms.Location{
		Id:           int64(l.ID),
		WarehouseId:  int64(l.WarehouseID),
		AreaId:       int64(l.AreaID),
		LocationCode: l.LocationCode,
		LocationType: locationTypeToProto(l.LocationType),
		Capacity:     floatToStr(l.Capacity),
		UsedCapacity: floatToStr(l.UsedCapacity),
		Status:       locationStatusToProto(l.Status),
	}
}

func locationsToProto(list []*model.Location) []*wms.Location {
	out := make([]*wms.Location, 0, len(list))
	for _, l := range list {
		out = append(out, locationToProto(l))
	}
	return out
}

// --- InventoryBalance conversions ---

func balanceStatusToProto(s model.InventoryStatus) wms.BalanceStatus {
	switch s {
	case model.InventoryStatusNormal:
		return wms.BalanceStatus_BALANCE_STATUS_NORMAL
	case model.InventoryStatusLocked:
		return wms.BalanceStatus_BALANCE_STATUS_LOCKED
	case model.InventoryStatusExpired:
		return wms.BalanceStatus_BALANCE_STATUS_EXPIRED
	default:
		return wms.BalanceStatus_BALANCE_STATUS_UNSPECIFIED
	}
}

func balanceStatusFromProto(s wms.BalanceStatus) model.InventoryStatus {
	switch s {
	case wms.BalanceStatus_BALANCE_STATUS_NORMAL:
		return model.InventoryStatusNormal
	case wms.BalanceStatus_BALANCE_STATUS_LOCKED:
		return model.InventoryStatusLocked
	case wms.BalanceStatus_BALANCE_STATUS_EXPIRED:
		return model.InventoryStatusExpired
	default:
		return ""
	}
}

func inventoryBalanceToProto(ib *model.InventoryBalance) *wms.InventoryBalance {
	if ib == nil {
		return nil
	}
	pb := &wms.InventoryBalance{
		Id:           int64(ib.ID),
		TenantId:     int64(ib.TenantID),
		MaterialId:   int64(ib.MaterialID),
		LocationId:   int64(ib.LocationID),
		BatchNo:      ib.BatchNo,
		Quantity:     floatToStr(ib.Quantity),
		LockedQty:    floatToStr(ib.LockedQty),
		AvailableQty: floatToStr(ib.AvailableQty),
		Status:       balanceStatusToProto(ib.Status),
		UnitCost:     floatToStr(ib.UnitCost),
	}
	if ib.ExpireDate != nil {
		pb.ExpireDate = ib.ExpireDate.Format("2006-01-02")
	}
	return pb
}

func inventoryBalancesToProto(list []*model.InventoryBalance) []*wms.InventoryBalance {
	out := make([]*wms.InventoryBalance, 0, len(list))
	for _, ib := range list {
		out = append(out, inventoryBalanceToProto(ib))
	}
	return out
}

// --- ReceiveOrder conversions ---

func receiveOrderStatusToProto(s model.ReceiveOrderStatus) wms.ReceiveOrderStatus {
	switch s {
	case model.ReceiveStatusDraft:
		return wms.ReceiveOrderStatus_RECEIVE_STATUS_DRAFT
	case model.ReceiveStatusReceiving:
		return wms.ReceiveOrderStatus_RECEIVE_STATUS_RECEIVING
	case model.ReceiveStatusReceived:
		return wms.ReceiveOrderStatus_RECEIVE_STATUS_RECEIVED
	case model.ReceiveStatusPutaway:
		return wms.ReceiveOrderStatus_RECEIVE_STATUS_PUTAWAY
	case model.ReceiveStatusCompleted:
		return wms.ReceiveOrderStatus_RECEIVE_STATUS_COMPLETED
	case model.ReceiveStatusCancelled:
		return wms.ReceiveOrderStatus_RECEIVE_STATUS_CANCELLED
	default:
		return wms.ReceiveOrderStatus_RECEIVE_STATUS_UNSPECIFIED
	}
}

func receiveOrderStatusFromProto(s wms.ReceiveOrderStatus) model.ReceiveOrderStatus {
	switch s {
	case wms.ReceiveOrderStatus_RECEIVE_STATUS_DRAFT:
		return model.ReceiveStatusDraft
	case wms.ReceiveOrderStatus_RECEIVE_STATUS_RECEIVING:
		return model.ReceiveStatusReceiving
	case wms.ReceiveOrderStatus_RECEIVE_STATUS_RECEIVED:
		return model.ReceiveStatusReceived
	case wms.ReceiveOrderStatus_RECEIVE_STATUS_PUTAWAY:
		return model.ReceiveStatusPutaway
	case wms.ReceiveOrderStatus_RECEIVE_STATUS_COMPLETED:
		return model.ReceiveStatusCompleted
	case wms.ReceiveOrderStatus_RECEIVE_STATUS_CANCELLED:
		return model.ReceiveStatusCancelled
	default:
		return ""
	}
}

func receiveOrderLineToProto(l *model.ReceiveOrderLine) *wms.ReceiveOrderLine {
	if l == nil {
		return nil
	}
	pb := &wms.ReceiveOrderLine{
		Id:          int64(l.ID),
		MaterialId:  int64(l.MaterialID),
		ExpectedQty: floatToStr(l.ExpectedQty),
		ReceivedQty: floatToStr(l.ReceivedQty),
		UnitPrice:   floatToStr(l.UnitPrice),
		BatchNo:     l.BatchNo,
	}
	if l.ExpireDate != nil {
		pb.ExpireDate = l.ExpireDate.Format("2006-01-02")
	}
	return pb
}

func receiveOrderToProto(ro *model.ReceiveOrder, lines []*model.ReceiveOrderLine) *wms.ReceiveOrder {
	if ro == nil {
		return nil
	}
	pb := &wms.ReceiveOrder{
		Id:         int64(ro.ID),
		TenantId:   int64(ro.TenantID),
		ReceiveNo:  ro.ReceiveNo,
		PoId:       parseStrToInt64(ro.PoID),
		SupplierId: int64(ro.SupplierID),
		Status:     receiveOrderStatusToProto(ro.Status),
	}
	if ro.ReceivedAt != nil {
		pb.ReceivedAt = ro.ReceivedAt.Unix()
	}
	if ro.PutawayAt != nil {
		pb.PutawayAt = ro.PutawayAt.Unix()
	}
	if ro.CompletedAt != nil {
		pb.CompletedAt = ro.CompletedAt.Unix()
	}
	for _, l := range lines {
		pb.Lines = append(pb.Lines, receiveOrderLineToProto(l))
	}
	return pb
}

// --- DeliveryOrder conversions ---

func deliveryOrderStatusToProto(s model.DeliveryOrderStatus) wms.DeliveryOrderStatus {
	switch s {
	case model.DeliveryStatusDraft:
		return wms.DeliveryOrderStatus_DELIVERY_STATUS_DRAFT
	case model.DeliveryStatusPicking:
		return wms.DeliveryOrderStatus_DELIVERY_STATUS_PICKING
	case model.DeliveryStatusPicked:
		return wms.DeliveryOrderStatus_DELIVERY_STATUS_PICKED
	case model.DeliveryStatusPacking:
		return wms.DeliveryOrderStatus_DELIVERY_STATUS_PACKING
	case model.DeliveryStatusOnHold:
		return wms.DeliveryOrderStatus_DELIVERY_STATUS_ON_HOLD
	case model.DeliveryStatusShipped:
		return wms.DeliveryOrderStatus_DELIVERY_STATUS_SHIPPED
	case model.DeliveryStatusCancelled:
		return wms.DeliveryOrderStatus_DELIVERY_STATUS_CANCELLED
	default:
		return wms.DeliveryOrderStatus_DELIVERY_STATUS_UNSPECIFIED
	}
}

func deliveryOrderStatusFromProto(s wms.DeliveryOrderStatus) model.DeliveryOrderStatus {
	switch s {
	case wms.DeliveryOrderStatus_DELIVERY_STATUS_DRAFT:
		return model.DeliveryStatusDraft
	case wms.DeliveryOrderStatus_DELIVERY_STATUS_PICKING:
		return model.DeliveryStatusPicking
	case wms.DeliveryOrderStatus_DELIVERY_STATUS_PICKED:
		return model.DeliveryStatusPicked
	case wms.DeliveryOrderStatus_DELIVERY_STATUS_PACKING:
		return model.DeliveryStatusPacking
	case wms.DeliveryOrderStatus_DELIVERY_STATUS_ON_HOLD:
		return model.DeliveryStatusOnHold
	case wms.DeliveryOrderStatus_DELIVERY_STATUS_SHIPPED:
		return model.DeliveryStatusShipped
	case wms.DeliveryOrderStatus_DELIVERY_STATUS_CANCELLED:
		return model.DeliveryStatusCancelled
	default:
		return ""
	}
}

func deliveryOrderLineToProto(l *model.DeliveryOrderLine) *wms.DeliveryOrderLine {
	if l == nil {
		return nil
	}
	return &wms.DeliveryOrderLine{
		Id:         int64(l.ID),
		MaterialId: int64(l.MaterialID),
		OrderedQty: floatToStr(l.OrderedQty),
		PickedQty:  floatToStr(l.PickedQty),
	}
}

func deliveryOrderToProto(do *model.DeliveryOrder, lines []*model.DeliveryOrderLine) *wms.DeliveryOrder {
	if do == nil {
		return nil
	}
	pb := &wms.DeliveryOrder{
		Id:         int64(do.ID),
		TenantId:   int64(do.TenantID),
		DeliveryNo: do.DeliveryNo,
		SoId:       parseStrToInt64(do.SoID),
		CustomerId: int64(do.CustomerID),
		Status:     deliveryOrderStatusToProto(do.Status),
	}
	if do.ShippedAt != nil {
		pb.ShippedAt = do.ShippedAt.Unix()
	}
	for _, l := range lines {
		pb.Lines = append(pb.Lines, deliveryOrderLineToProto(l))
	}
	return pb
}

// --- CountPlan conversions ---

func countPlanToProto(cp *model.CountPlan) *wms.CountPlan {
	if cp == nil {
		return nil
	}
	return &wms.CountPlan{
		Id:          int64(cp.ID),
		PlanNo:      cp.PlanNo,
		WarehouseId: int64(cp.WarehouseID),
		PlanType:    cp.PlanType,
		Status:      string(cp.Status),
	}
}

// --- Pagination helpers ---

func pageResponseFromTotal(pr *common.PageRequest, total int) *common.PageResponse {
	page := int32(1)
	pageSize := int32(20)
	if pr != nil {
		if pr.Page > 0 {
			page = pr.Page
		}
		if pr.PageSize > 0 {
			pageSize = pr.PageSize
		}
	}
	totalPages := int32(0)
	if pageSize > 0 {
		totalPages = int32((total + int(pageSize) - 1) / int(pageSize))
	}
	return &common.PageResponse{
		Page:       page,
		PageSize:   pageSize,
		Total:      int32(total),
		TotalPages: totalPages,
	}
}

// --- Utility ---

func floatToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func parseStrToInt64(s string) int64 {
	if s == "" {
		return 0
	}
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return n
}

func strToFloat64(s string) float64 {
	if s == "" {
		return 0
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

func parseExpireDate(s string) *time.Time {
	if s == "" {
		return nil
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return nil
	}
	return &t
}
