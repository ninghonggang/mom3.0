package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"

	"go.uber.org/zap"

	"github.com/ninghonggang/mom-platform/pkg/eventbus"
	"mom-platform/services/wms-service/internal/model"
	"mom-platform/services/wms-service/internal/repository"
)

// EventPublisher is the interface for publishing domain events.
// The eventbus.EventPublisher from pkg/eventbus satisfies this interface.
type EventPublisher interface {
	Publish(ctx context.Context, subject string, payload interface{}) error
}

// Service implements the WMS business logic.
type Service struct {
	repo repository.Repository
	log  *zap.Logger
	pub  EventPublisher
}

// New creates a new WMS Service.
func New(repo repository.Repository, log *zap.Logger, pub EventPublisher) *Service {
	if pub == nil {
		pub = &noopEventPublisher{log: log}
	}
	return &Service{repo: repo, log: log, pub: pub}
}

// ---------------------------------------------------------------------------
// Request / Response types
// ---------------------------------------------------------------------------

// CreateReceiveOrderReq is the input for creating a receive order.
type CreateReceiveOrderReq struct {
	TenantID   uint
	PoID       string
	SupplierID uint
	Lines      []ReceiveLineInput
}

// ReceiveLineInput represents one expected line on a receive order.
type ReceiveLineInput struct {
	MaterialID  uint
	ExpectedQty float64
	// UnitPrice 为采购单价，上架时驱动库存单位成本的移动加权平均计算。
	UnitPrice  float64
	BatchNo    string
	ExpireDate *time.Time
}

// ReceiveLineConfirm represents confirmed received quantities for a line.
type ReceiveLineConfirm struct {
	LineID      uint
	ReceivedQty float64
	BatchNo     string
	ExpireDate  *time.Time
}

// PutawayItem represents one putaway action.
type PutawayItem struct {
	LineID     uint
	MaterialID uint
	LocationID uint
	Quantity   float64
}

// CreateDeliveryOrderReq is the input for creating a delivery order.
type CreateDeliveryOrderReq struct {
	TenantID   uint
	SoID       string
	CustomerID uint
	Lines      []DeliveryLineInput
}

// DeliveryLineInput represents one ordered line on a delivery order.
type DeliveryLineInput struct {
	MaterialID uint
	OrderedQty float64
}

// PickItem represents one pick action.
type PickItem struct {
	LineID     uint
	LocationID uint
	MaterialID uint
	Quantity   float64
	PickerID   uint
}

// SubmitCountRecordReq is the input for submitting count records.
type SubmitCountRecordReq struct {
	PlanID uint
	Items  []CountRecordItem
}

// CountRecordItem represents one count record line.
type CountRecordItem struct {
	MaterialID uint
	LocationID uint
	ActualQty  float64
}

// ---------------------------------------------------------------------------
// Errors
// ---------------------------------------------------------------------------

var (
	ErrInvalidStatus       = errors.New("invalid order status for operation")
	ErrInvalidQuantity     = errors.New("invalid quantity")
	ErrEmptyLines          = errors.New("order must have at least one line")
	ErrLineNotFound        = errors.New("receive order line not found")
	ErrLocationNotFound    = errors.New("location not found")
	ErrWarehouseNotFound   = errors.New("warehouse not found")
	ErrOrderNotFound       = errors.New("order not found")
	ErrReceiveOrderNotRecv = errors.New("receive order not in RECEIVED status")
)

// ---------------------------------------------------------------------------
// Warehouse operations
// ---------------------------------------------------------------------------

// CreateWarehouse creates a new warehouse.
func (s *Service) CreateWarehouse(ctx context.Context, w *model.Warehouse) (*model.Warehouse, error) {
	if w.WarehouseCode == "" || w.WarehouseName == "" {
		return nil, fmt.Errorf("%w: warehouse_code and warehouse_name are required", ErrInvalidQuantity)
	}
	if w.Status == "" {
		w.Status = model.WarehouseStatusActive
	}
	if w.TenantID == 0 {
		w.TenantID = 1
	}
	if err := s.repo.CreateWarehouse(ctx, w); err != nil {
		return nil, fmt.Errorf("create warehouse: %w", err)
	}
	s.log.Info("warehouse created", zap.Uint("id", w.ID), zap.String("code", w.WarehouseCode))
	return w, nil
}

// ListWarehouses returns all warehouses for a tenant.
func (s *Service) ListWarehouses(ctx context.Context, tenantID uint) ([]*model.Warehouse, error) {
	return s.repo.ListWarehouses(ctx, tenantID)
}

// ---------------------------------------------------------------------------
// Location operations
// ---------------------------------------------------------------------------

// CreateLocation creates a new storage location.
func (s *Service) CreateLocation(ctx context.Context, l *model.Location) (*model.Location, error) {
	if l.LocationCode == "" {
		return nil, fmt.Errorf("%w: location_code is required", ErrInvalidQuantity)
	}
	if l.Status == "" {
		l.Status = model.LocationStatusActive
	}
	if err := s.repo.CreateLocation(ctx, l); err != nil {
		return nil, fmt.Errorf("create location: %w", err)
	}
	s.log.Info("location created", zap.Uint("id", l.ID), zap.String("code", l.LocationCode))
	return l, nil
}

// ListLocations returns locations filtered by warehouse and area.
func (s *Service) ListLocations(ctx context.Context, warehouseID, areaID uint) ([]*model.Location, error) {
	return s.repo.ListLocations(ctx, warehouseID, areaID)
}

// ---------------------------------------------------------------------------
// Inventory operations
// ---------------------------------------------------------------------------

// GetBalance returns the inventory balance for a material/location/batch.
func (s *Service) GetBalance(ctx context.Context, tenantID, materialID, locationID uint, batchNo string) (*model.InventoryBalance, error) {
	ib, err := s.repo.GetInventoryBalanceByMaterialLocation(ctx, tenantID, materialID, locationID, batchNo)
	if err != nil {
		return nil, fmt.Errorf("get balance: %w", err)
	}
	return ib, nil
}

// ListBalances returns inventory balances filtered by material and status.
func (s *Service) ListBalances(ctx context.Context, tenantID, materialID uint, status model.InventoryStatus) ([]*model.InventoryBalance, error) {
	return s.repo.ListBalances(ctx, tenantID, materialID, status)
}

// LockInventory locks the given quantity on an inventory balance.
func (s *Service) LockInventory(ctx context.Context, inventoryID uint, qty float64) (*model.InventoryBalance, error) {
	if qty <= 0 {
		return nil, fmt.Errorf("%w: qty must be positive", ErrInvalidQuantity)
	}
	if err := s.repo.LockInventory(ctx, inventoryID, qty); err != nil {
		return nil, err
	}
	return s.repo.GetInventoryBalance(ctx, inventoryID)
}

// UnlockInventory releases a previously locked quantity.
func (s *Service) UnlockInventory(ctx context.Context, inventoryID uint, qty float64) (*model.InventoryBalance, error) {
	if qty <= 0 {
		return nil, fmt.Errorf("%w: qty must be positive", ErrInvalidQuantity)
	}
	if err := s.repo.UnlockInventory(ctx, inventoryID, qty); err != nil {
		return nil, err
	}
	return s.repo.GetInventoryBalance(ctx, inventoryID)
}

// ---------------------------------------------------------------------------
// Receive order operations
// ---------------------------------------------------------------------------

// autoSeq is a simple in-process counter for order number uniqueness.
var autoSeq uint64

// generateOrderNo creates an order number like RO-20260807-00001-A3F2.
func generateOrderNo(prefix string) string {
	n := atomic.AddUint64(&autoSeq, 1)
	randSuffix := rand.Intn(0xFFFF)
	return fmt.Sprintf("%s-%s-%05d-%04X", prefix, time.Now().Format("20060102"), n, randSuffix)
}

// CreateReceiveOrder validates the request, auto-generates the receive number,
// and persists the order with all lines in status DRAFT.
func (s *Service) CreateReceiveOrder(ctx context.Context, req *CreateReceiveOrderReq) (*model.ReceiveOrder, error) {
	if req.TenantID == 0 {
		return nil, ErrInvalidQuantity
	}
	if len(req.Lines) == 0 {
		return nil, ErrEmptyLines
	}
	for i, l := range req.Lines {
		if l.ExpectedQty <= 0 {
			return nil, fmt.Errorf("line %d: %w: expected_qty must be positive", i, ErrInvalidQuantity)
		}
	}

	ro := &model.ReceiveOrder{
		TenantID:   req.TenantID,
		ReceiveNo:  generateOrderNo("RO"),
		PoID:       req.PoID,
		SupplierID: req.SupplierID,
		Status:     model.ReceiveStatusDraft,
	}

	err := s.repo.WithTx(ctx, func(ctx context.Context) error {
		if err := s.repo.CreateReceiveOrder(ctx, ro); err != nil {
			return fmt.Errorf("create receive order: %w", err)
		}
		for _, l := range req.Lines {
			line := &model.ReceiveOrderLine{
				ReceiveOrderID: ro.ID,
				MaterialID:     l.MaterialID,
				ExpectedQty:    l.ExpectedQty,
				ReceivedQty:    0,
				UnitPrice:      l.UnitPrice,
				BatchNo:        l.BatchNo,
				ExpireDate:     l.ExpireDate,
			}
			if err := s.repo.CreateReceiveOrderLine(ctx, line); err != nil {
				return fmt.Errorf("create receive order line: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	s.log.Info("receive order created",
		zap.Uint("id", ro.ID), zap.String("receive_no", ro.ReceiveNo))
	return ro, nil
}

// ReceiveConfirm validates received quantities, updates each line's received
// quantity, and transitions the order to RECEIVED status.
func (s *Service) ReceiveConfirm(ctx context.Context, receiveOrderID uint, lines []ReceiveLineConfirm) (*model.ReceiveOrder, error) {
	if len(lines) == 0 {
		return nil, ErrEmptyLines
	}

	var result *model.ReceiveOrder
	err := s.repo.WithTx(ctx, func(ctx context.Context) error {
		ro, err := s.repo.GetReceiveOrder(ctx, receiveOrderID)
		if err != nil {
			return fmt.Errorf("get receive order: %w", err)
		}
		if ro.Status != model.ReceiveStatusDraft && ro.Status != model.ReceiveStatusReceiving {
			return fmt.Errorf("current status %s: %w", ro.Status, ErrInvalidStatus)
		}

		// Validate all line IDs exist before mutating.
		existingLines, err := s.repo.ListReceiveOrderLines(ctx, receiveOrderID)
		if err != nil {
			return fmt.Errorf("list receive lines: %w", err)
		}
		lineMap := make(map[uint]*model.ReceiveOrderLine, len(existingLines))
		for _, l := range existingLines {
			lineMap[l.ID] = l
		}

		for _, c := range lines {
			l, ok := lineMap[c.LineID]
			if !ok {
				return fmt.Errorf("line %d: %w", c.LineID, ErrLineNotFound)
			}
			if c.ReceivedQty < 0 {
				return fmt.Errorf("line %d: %w: received_qty cannot be negative", l.ID, ErrInvalidQuantity)
			}
			if c.ReceivedQty > l.ExpectedQty {
				return fmt.Errorf("line %d: %w: received_qty exceeds expected_qty", l.ID, ErrInvalidQuantity)
			}
		}

		// Apply updates.
		for _, c := range lines {
			if err := s.repo.UpdateReceiveOrderLineReceivedQty(ctx, c.LineID, c.ReceivedQty); err != nil {
				return fmt.Errorf("update line %d: %w", c.LineID, err)
			}
		}

		// Transition status to RECEIVING then RECEIVED.
		if ro.Status == model.ReceiveStatusDraft {
			ro.Status = model.ReceiveStatusReceiving
			if err := s.repo.UpdateReceiveOrderStatus(ctx, ro.ID, ro.Status); err != nil {
				return fmt.Errorf("update receive order status: %w", err)
			}
		}

		now := time.Now()
		ro.Status = model.ReceiveStatusReceived
		ro.ReceivedAt = &now
		if err := s.repo.UpdateReceiveOrder(ctx, ro); err != nil {
			return fmt.Errorf("update receive order status: %w", err)
		}

		s.log.Info("receive order confirmed",
			zap.Uint("id", ro.ID), zap.String("receive_no", ro.ReceiveNo))
		result = ro
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Putaway creates putaway records, creates or updates inventory balances,
// updates location used capacity, and transitions the order to COMPLETED.
// After completion it publishes a wms.receive.completed event.
func (s *Service) Putaway(ctx context.Context, receiveOrderID uint, items []PutawayItem) (*model.ReceiveOrder, error) {
	if len(items) == 0 {
		return nil, ErrEmptyLines
	}

	var result *model.ReceiveOrder
	err := s.repo.WithTx(ctx, func(ctx context.Context) error {
		ro, err := s.repo.GetReceiveOrder(ctx, receiveOrderID)
		if err != nil {
			return fmt.Errorf("get receive order: %w", err)
		}
		if ro.Status != model.ReceiveStatusReceived && ro.Status != model.ReceiveStatusPutaway {
			return fmt.Errorf("current status %s: %w", ro.Status, ErrInvalidStatus)
		}

		// PutawayLine 只带 line_id，物料和批次必须从收货单行反查，
		// 否则库存余额会以 material_id=0 落库，后续拣货永远找不到。
		roLines, err := s.repo.ListReceiveOrderLines(ctx, receiveOrderID)
		if err != nil {
			return fmt.Errorf("list receive order lines: %w", err)
		}
		roLineByID := make(map[uint]*model.ReceiveOrderLine, len(roLines))
		for _, l := range roLines {
			roLineByID[l.ID] = l
		}

		for i := range items {
			item := &items[i]
			var batchNo string
			var unitPrice float64
			if line, ok := roLineByID[item.LineID]; ok {
				if item.MaterialID == 0 {
					item.MaterialID = line.MaterialID
				}
				batchNo = line.BatchNo
				unitPrice = line.UnitPrice
			} else if item.MaterialID == 0 {
				return fmt.Errorf("line %d: %w: not a line of receive order %d",
					item.LineID, ErrLineNotFound, receiveOrderID)
			}
			if item.Quantity <= 0 {
				return fmt.Errorf("material %d: %w: quantity must be positive", item.MaterialID, ErrInvalidQuantity)
			}
			// Validate location exists.
			loc, err := s.repo.GetLocation(ctx, item.LocationID)
			if err != nil {
				return fmt.Errorf("location %d: %w", item.LocationID, err)
			}

			// Check capacity.
			if loc.Capacity > 0 && loc.UsedCapacity+item.Quantity > loc.Capacity {
				return fmt.Errorf("location %d: %w: exceeds capacity", item.LocationID, ErrInvalidQuantity)
			}

			now := time.Now()
			// Create putaway record.
			pr := &model.PutawayRecord{
				ReceiveOrderID: receiveOrderID,
				LocationID:     item.LocationID,
				MaterialID:     item.MaterialID,
				Quantity:       item.Quantity,
				BatchNo:        batchNo,
				PutawayTime:    now,
			}
			if err := s.repo.CreatePutawayRecord(ctx, pr); err != nil {
				return fmt.Errorf("create putaway record: %w", err)
			}

			// Find or create inventory balance.
			ib, err := s.repo.GetInventoryBalanceByMaterialLocation(
				ctx, ro.TenantID, item.MaterialID, item.LocationID, "",
			)
			if err != nil && !errors.Is(err, repository.ErrNotFound) {
				return fmt.Errorf("get inventory balance: %w", err)
			}
			if errors.Is(err, repository.ErrNotFound) {
				ib := &model.InventoryBalance{
					TenantID:   ro.TenantID,
					MaterialID: item.MaterialID,
					LocationID: item.LocationID,
					BatchNo:    "",
					Quantity:   item.Quantity,
					LockedQty:  0,
					Status:     model.InventoryStatusNormal,
					// 首次建账，采购单价即为期初单位成本。
					UnitCost: unitPrice,
				}
				// BeforeSave hook sets AvailableQty = Quantity - LockedQty.
				if err := s.repo.CreateInventoryBalance(ctx, ib); err != nil {
					return fmt.Errorf("create inventory balance: %w", err)
				}
			} else {
				// 已有台账按移动加权平均滚动单位成本，保证库存估值随进价变化收敛。
				if err := s.repo.IncrementInventoryWithCost(ctx, ib.ID, item.Quantity, unitPrice); err != nil {
					return fmt.Errorf("increment inventory: %w", err)
				}
			}

			// Update location used capacity.
			if err := s.repo.UpdateLocationUsedCapacity(ctx, item.LocationID, item.Quantity); err != nil {
				return fmt.Errorf("update location capacity: %w", err)
			}
		}

		// Transition order to COMPLETED.
		now := time.Now()
		ro.Status = model.ReceiveStatusCompleted
		ro.CompletedAt = &now
		if ro.ReceivedAt == nil {
			ro.ReceivedAt = &now
		}
		if err := s.repo.UpdateReceiveOrder(ctx, ro); err != nil {
			return fmt.Errorf("update receive order: %w", err)
		}

		s.log.Info("putaway completed",
			zap.Uint("id", ro.ID), zap.String("receive_no", ro.ReceiveNo))

		result = ro
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Publish event after tx commits.
	_ = s.pub.Publish(ctx, eventbus.SubjectWMSReceiveCompleted, map[string]interface{}{
		"receive_order_id": result.ID,
		"receive_no":       result.ReceiveNo,
	})
	return result, nil
}

// GetReceiveOrder returns a receive order with its lines.
func (s *Service) GetReceiveOrder(ctx context.Context, id uint) (*model.ReceiveOrder, []*model.ReceiveOrderLine, error) {
	ro, err := s.repo.GetReceiveOrder(ctx, id)
	if err != nil {
		return nil, nil, fmt.Errorf("get receive order: %w", err)
	}
	lines, err := s.repo.ListReceiveOrderLines(ctx, id)
	if err != nil {
		return nil, nil, fmt.Errorf("list receive order lines: %w", err)
	}
	return ro, lines, nil
}

// ListReceiveOrders returns receive orders filtered by status and supplier.
func (s *Service) ListReceiveOrders(ctx context.Context, status model.ReceiveOrderStatus, supplierID uint) ([]*model.ReceiveOrder, error) {
	return s.repo.ListReceiveOrders(ctx, status, supplierID)
}

// ---------------------------------------------------------------------------
// Delivery order operations
// ---------------------------------------------------------------------------

// CreateDeliveryOrder validates the request, auto-generates the delivery
// number, and persists the order with all lines in status DRAFT.
func (s *Service) CreateDeliveryOrder(ctx context.Context, req *CreateDeliveryOrderReq) (*model.DeliveryOrder, error) {
	if req.TenantID == 0 {
		return nil, ErrInvalidQuantity
	}
	if len(req.Lines) == 0 {
		return nil, ErrEmptyLines
	}
	for i, l := range req.Lines {
		if l.OrderedQty <= 0 {
			return nil, fmt.Errorf("line %d: %w: ordered_qty must be positive", i, ErrInvalidQuantity)
		}
	}

	do := &model.DeliveryOrder{
		TenantID:   req.TenantID,
		DeliveryNo: generateOrderNo("DO"),
		SoID:       req.SoID,
		CustomerID: req.CustomerID,
		Status:     model.DeliveryStatusDraft,
	}

	err := s.repo.WithTx(ctx, func(ctx context.Context) error {
		if err := s.repo.CreateDeliveryOrder(ctx, do); err != nil {
			return fmt.Errorf("create delivery order: %w", err)
		}
		for _, l := range req.Lines {
			line := &model.DeliveryOrderLine{
				DeliveryOrderID: do.ID,
				MaterialID:      l.MaterialID,
				OrderedQty:      l.OrderedQty,
				PickedQty:       0,
			}
			if err := s.repo.CreateDeliveryOrderLine(ctx, line); err != nil {
				return fmt.Errorf("create delivery order line: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	s.log.Info("delivery order created",
		zap.Uint("id", do.ID), zap.String("delivery_no", do.DeliveryNo))
	return do, nil
}

// PickItems validates that sufficient available inventory exists, atomically
// locks the inventory, creates pick records, updates delivery line picked
// quantities, and transitions the order to PICKED status.
func (s *Service) PickItems(ctx context.Context, deliveryOrderID uint, picks []PickItem) (*model.DeliveryOrder, error) {
	if len(picks) == 0 {
		return nil, ErrEmptyLines
	}

	var result *model.DeliveryOrder
	err := s.repo.WithTx(ctx, func(ctx context.Context) error {
		do, err := s.repo.GetDeliveryOrder(ctx, deliveryOrderID)
		if err != nil {
			return fmt.Errorf("get delivery order: %w", err)
		}
		if do.Status != model.DeliveryStatusDraft && do.Status != model.DeliveryStatusPicking {
			return fmt.Errorf("current status %s: %w", do.Status, ErrInvalidStatus)
		}

		// Aggregate picks by delivery line (material).
		lines, err := s.repo.ListDeliveryOrderLines(ctx, deliveryOrderID)
		if err != nil {
			return fmt.Errorf("list delivery lines: %w", err)
		}
		lineByMaterial := make(map[uint]*model.DeliveryOrderLine, len(lines))
		lineByID := make(map[uint]*model.DeliveryOrderLine, len(lines))
		for _, l := range lines {
			lineByMaterial[l.MaterialID] = l
			lineByID[l.ID] = l
		}

		pickedByMaterial := make(map[uint]float64)
		now := time.Now()

		for i := range picks {
			pick := &picks[i]
			// 拣货行以 line_id 标识（PickLine 里没有 material_id），
			// 物料必须从订单行反查，否则后续按物料聚合时会全部落到 0。
			if pick.MaterialID == 0 {
				line, ok := lineByID[pick.LineID]
				if !ok {
					return fmt.Errorf("line %d: %w: not a line of delivery order %d",
						pick.LineID, ErrLineNotFound, deliveryOrderID)
				}
				pick.MaterialID = line.MaterialID
			}
			if pick.Quantity <= 0 {
				return fmt.Errorf("material %d: %w: pick quantity must be positive", pick.MaterialID, ErrInvalidQuantity)
			}
			// Find the inventory balance to pick from.
			ib, err := s.repo.GetInventoryBalanceByMaterialLocation(
				ctx, do.TenantID, pick.MaterialID, pick.LocationID, "",
			)
			if err != nil {
				return fmt.Errorf("find inventory for material %d at location %d: %w",
					pick.MaterialID, pick.LocationID, err)
			}
			// Lock inventory atomically (decreases available_qty, increases locked_qty).
			if err := s.repo.LockInventory(ctx, ib.ID, pick.Quantity); err != nil {
				return fmt.Errorf("lock inventory %d for qty %f: %w", ib.ID, pick.Quantity, err)
			}

			// Create pick record.
			pr := &model.PickRecord{
				DeliveryOrderID: deliveryOrderID,
				LocationID:      pick.LocationID,
				MaterialID:      pick.MaterialID,
				PickedQty:       pick.Quantity,
				PickerID:        pick.PickerID,
				BatchNo:         "",
				PickTime:        now,
			}
			if err := s.repo.CreatePickRecord(ctx, pr); err != nil {
				return fmt.Errorf("create pick record: %w", err)
			}

			pickedByMaterial[pick.MaterialID] += pick.Quantity
		}

		// Update delivery line picked quantities.
		for materialID, qty := range pickedByMaterial {
			line, ok := lineByMaterial[materialID]
			if !ok {
				return fmt.Errorf("material %d: %w: no matching delivery line", materialID, ErrLineNotFound)
			}
			newPicked := line.PickedQty + qty
			if newPicked > line.OrderedQty {
				return fmt.Errorf("material %d: %w: picked qty exceeds ordered", materialID, ErrInvalidQuantity)
			}
			if err := s.repo.UpdateDeliveryOrderLinePickedQty(ctx, line.ID, newPicked); err != nil {
				return fmt.Errorf("update delivery line: %w", err)
			}
		}

		// Transition to PICKED.
		if err := s.repo.UpdateDeliveryOrderStatus(ctx, deliveryOrderID, model.DeliveryStatusPicked); err != nil {
			return fmt.Errorf("update delivery status: %w", err)
		}

		do.Status = model.DeliveryStatusPicked
		result = do
		s.log.Info("pick completed",
			zap.Uint("id", deliveryOrderID), zap.Int("picks", len(picks)))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShipOrder deducts locked inventory, transitions the delivery order to
// SHIPPED, and publishes a wms.delivery.shipped event.
func (s *Service) ShipOrder(ctx context.Context, deliveryOrderID uint) (*model.DeliveryOrder, error) {
	var result *model.DeliveryOrder
	err := s.repo.WithTx(ctx, func(ctx context.Context) error {
		do, err := s.repo.GetDeliveryOrder(ctx, deliveryOrderID)
		if err != nil {
			return fmt.Errorf("get delivery order: %w", err)
		}
		if do.Status != model.DeliveryStatusPicked &&
			do.Status != model.DeliveryStatusPacking &&
			do.Status != model.DeliveryStatusOnHold {
			return fmt.Errorf("current status %s: %w", do.Status, ErrInvalidStatus)
		}

		// Deduct picked inventory from each pick record's source location.
		picks, err := s.repo.ListPickRecords(ctx, deliveryOrderID)
		if err != nil {
			return fmt.Errorf("list pick records: %w", err)
		}
		// Aggregate by inventory balance to deduct in bulk.
		type invKey struct {
			materialID uint
			locationID uint
		}
		type deductEntry struct {
			inventoryID uint
			qty         float64
		}
		deduct := make(map[invKey]*deductEntry)
		for _, p := range picks {
			ib, err := s.repo.GetInventoryBalanceByMaterialLocation(
				ctx, do.TenantID, p.MaterialID, p.LocationID, "",
			)
			if err != nil {
				return fmt.Errorf("find inventory for pick material %d: %w", p.MaterialID, err)
			}
			k := invKey{p.MaterialID, p.LocationID}
			if entry, ok := deduct[k]; ok {
				entry.qty += p.PickedQty
			} else {
				deduct[k] = &deductEntry{inventoryID: ib.ID, qty: p.PickedQty}
			}
		}
		for _, entry := range deduct {
			if err := s.repo.DeductLockedInventory(ctx, entry.inventoryID, entry.qty); err != nil {
				return fmt.Errorf("deduct locked inventory %d qty %f: %w", entry.inventoryID, entry.qty, err)
			}
		}

		// Transition to SHIPPED.
		now := time.Now()
		do.Status = model.DeliveryStatusShipped
		do.ShippedAt = &now
		if err := s.repo.UpdateDeliveryOrder(ctx, do); err != nil {
			return fmt.Errorf("update delivery order: %w", err)
		}

		s.log.Info("order shipped",
			zap.Uint("id", do.ID), zap.String("delivery_no", do.DeliveryNo))

		result = do
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Publish event after tx commits.
	_ = s.pub.Publish(ctx, eventbus.SubjectWMSShipped, map[string]interface{}{
		"delivery_order_id": result.ID,
		"delivery_no":       result.DeliveryNo,
	})
	return result, nil
}

// GetDeliveryOrder returns a delivery order with its lines.
func (s *Service) GetDeliveryOrder(ctx context.Context, id uint) (*model.DeliveryOrder, []*model.DeliveryOrderLine, error) {
	do, err := s.repo.GetDeliveryOrder(ctx, id)
	if err != nil {
		return nil, nil, fmt.Errorf("get delivery order: %w", err)
	}
	lines, err := s.repo.ListDeliveryOrderLines(ctx, id)
	if err != nil {
		return nil, nil, fmt.Errorf("list delivery order lines: %w", err)
	}
	return do, lines, nil
}

// ListDeliveryOrders returns delivery orders filtered by status and customer.
func (s *Service) ListDeliveryOrders(ctx context.Context, status model.DeliveryOrderStatus, customerID uint) ([]*model.DeliveryOrder, error) {
	return s.repo.ListDeliveryOrders(ctx, status, customerID)
}

// ---------------------------------------------------------------------------
// Count operations
// ---------------------------------------------------------------------------

// CreateCountPlan creates a new inventory count plan.
func (s *Service) CreateCountPlan(ctx context.Context, warehouseID uint, planType string) (*model.CountPlan, error) {
	if warehouseID == 0 {
		return nil, fmt.Errorf("%w: warehouse_id is required", ErrInvalidQuantity)
	}
	cp := &model.CountPlan{
		PlanNo:     generateOrderNo("CP"),
		WarehouseID: warehouseID,
		PlanType:   planType,
		Status:     model.CountPlanStatusDraft,
	}
	if err := s.repo.CreateCountPlan(ctx, cp); err != nil {
		return nil, fmt.Errorf("create count plan: %w", err)
	}
	s.log.Info("count plan created", zap.Uint("id", cp.ID), zap.String("plan_no", cp.PlanNo))
	return cp, nil
}

// SubmitCountRecord submits count records for a plan and returns the updated plan.
func (s *Service) SubmitCountRecord(ctx context.Context, req *SubmitCountRecordReq) (*model.CountPlan, error) {
	if req.PlanID == 0 {
		return nil, fmt.Errorf("%w: plan_id is required", ErrInvalidQuantity)
	}
	if len(req.Items) == 0 {
		return nil, ErrEmptyLines
	}

	cp, err := s.repo.GetCountPlan(ctx, req.PlanID)
	if err != nil {
		return nil, fmt.Errorf("get count plan: %w", err)
	}

	for _, item := range req.Items {
		// Determine system quantity.
		balances, err := s.repo.ListInventoryByMaterial(ctx, 0, item.MaterialID)
		if err != nil {
			return nil, fmt.Errorf("list inventory: %w", err)
		}
		var systemQty float64
		for _, b := range balances {
			if b.LocationID == item.LocationID {
				systemQty += b.Quantity
			}
		}

		cr := &model.CountRecord{
			JobID:      cp.ID,
			MaterialID: item.MaterialID,
			LocationID: item.LocationID,
			SystemQty:  systemQty,
			ActualQty:  item.ActualQty,
		}
		cr.DiffQty = item.ActualQty - systemQty

		if err := s.repo.CreateCountRecord(ctx, cr); err != nil {
			return nil, fmt.Errorf("create count record: %w", err)
		}
	}

	// Transition plan to COMPLETED.
	cp.Status = model.CountPlanStatusCompleted
	if err := s.repo.UpdateCountPlan(ctx, cp); err != nil {
		return nil, fmt.Errorf("update count plan: %w", err)
	}

	s.log.Info("count records submitted", zap.Uint("plan_id", cp.ID), zap.Int("items", len(req.Items)))
	return cp, nil
}

// ---------------------------------------------------------------------------
// noopEventPublisher
// ---------------------------------------------------------------------------

type noopEventPublisher struct {
	log *zap.Logger
}

func (n *noopEventPublisher) Publish(_ context.Context, subject string, payload interface{}) error {
	n.log.Info("event published (stub)", zap.String("subject", subject), zap.Any("payload", payload))
	return nil
}

// floatToStr converts a float64 to a string for proto quantity fields.
func floatToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
