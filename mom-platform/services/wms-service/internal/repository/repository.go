package repository

import (
	"context"
	"errors"

	"mom-platform/services/wms-service/internal/model"
)

// Common errors returned by the repository layer.
var (
	ErrNotFound              = errors.New("record not found")
	ErrInsufficientInventory = errors.New("insufficient available inventory")
	ErrInsufficientLockedQty = errors.New("insufficient locked quantity")
	ErrConflict              = errors.New("concurrent update conflict")
)

// TransactionManager manages database transactions. The callback receives a
// context that carries the transaction; repository methods called with that
// context will execute within the same transaction.
type TransactionManager interface {
	WithTx(ctx context.Context, fn func(ctx context.Context) error) error
}

// Warehouse operations.
type Warehouse interface {
	CreateWarehouse(ctx context.Context, w *model.Warehouse) error
	GetWarehouse(ctx context.Context, id uint) (*model.Warehouse, error)
	GetWarehouseByCode(ctx context.Context, tenantID uint, code string) (*model.Warehouse, error)
	ListWarehouses(ctx context.Context, tenantID uint) ([]*model.Warehouse, error)
}

// Location operations.
type Location interface {
	CreateLocation(ctx context.Context, l *model.Location) error
	GetLocation(ctx context.Context, id uint) (*model.Location, error)
	GetLocationByCode(ctx context.Context, warehouseID uint, code string) (*model.Location, error)
	ListLocations(ctx context.Context, warehouseID, areaID uint) ([]*model.Location, error)
	UpdateLocationUsedCapacity(ctx context.Context, id uint, delta float64) error
}

// InventoryBalance operations.
type Inventory interface {
	CreateInventoryBalance(ctx context.Context, ib *model.InventoryBalance) error
	GetInventoryBalance(ctx context.Context, id uint) (*model.InventoryBalance, error)
	// GetInventoryBalanceForUpdate reads with SELECT ... FOR UPDATE (must be in a tx).
	GetInventoryBalanceForUpdate(ctx context.Context, id uint) (*model.InventoryBalance, error)
	GetInventoryBalanceByMaterialLocation(
		ctx context.Context, tenantID, materialID, locationID uint, batchNo string,
	) (*model.InventoryBalance, error)
	ListInventoryByMaterial(ctx context.Context, tenantID, materialID uint) ([]*model.InventoryBalance, error)
	ListBalances(ctx context.Context, tenantID, materialID uint, status model.InventoryStatus) ([]*model.InventoryBalance, error)
	// IncrementInventoryQuantity adds qty to both quantity and available_qty atomically.
	IncrementInventoryQuantity(ctx context.Context, id uint, qty float64) error
	// IncrementInventoryWithCost adds qty and recomputes unit_cost by moving
	// weighted average. A unitPrice <= 0 leaves the existing cost untouched.
	IncrementInventoryWithCost(ctx context.Context, id uint, qty, unitPrice float64) error
	// LockInventory atomically decreases available_qty and increases locked_qty.
	LockInventory(ctx context.Context, id uint, qty float64) error
	// UnlockInventory atomically increases available_qty and decreases locked_qty.
	UnlockInventory(ctx context.Context, id uint, qty float64) error
	// DeductLockedInventory removes qty from both quantity and locked_qty (on shipment).
	DeductLockedInventory(ctx context.Context, id uint, qty float64) error
}

// ReceiveOrder operations.
type ReceiveOrder interface {
	CreateReceiveOrder(ctx context.Context, ro *model.ReceiveOrder) error
	GetReceiveOrder(ctx context.Context, id uint) (*model.ReceiveOrder, error)
	GetReceiveOrderByNo(ctx context.Context, tenantID uint, no string) (*model.ReceiveOrder, error)
	UpdateReceiveOrder(ctx context.Context, ro *model.ReceiveOrder) error
	UpdateReceiveOrderStatus(ctx context.Context, id uint, status model.ReceiveOrderStatus) error
	ListReceiveOrders(ctx context.Context, status model.ReceiveOrderStatus, supplierID uint) ([]*model.ReceiveOrder, error)
}

// ReceiveOrderLine operations.
type ReceiveOrderLine interface {
	CreateReceiveOrderLine(ctx context.Context, rol *model.ReceiveOrderLine) error
	GetReceiveOrderLine(ctx context.Context, id uint) (*model.ReceiveOrderLine, error)
	ListReceiveOrderLines(ctx context.Context, receiveOrderID uint) ([]*model.ReceiveOrderLine, error)
	UpdateReceiveOrderLineReceivedQty(ctx context.Context, id uint, receivedQty float64) error
}

// DeliveryOrder operations.
type DeliveryOrder interface {
	CreateDeliveryOrder(ctx context.Context, do *model.DeliveryOrder) error
	GetDeliveryOrder(ctx context.Context, id uint) (*model.DeliveryOrder, error)
	GetDeliveryOrderByNo(ctx context.Context, tenantID uint, no string) (*model.DeliveryOrder, error)
	UpdateDeliveryOrder(ctx context.Context, do *model.DeliveryOrder) error
	UpdateDeliveryOrderStatus(ctx context.Context, id uint, status model.DeliveryOrderStatus) error
	ListDeliveryOrders(ctx context.Context, status model.DeliveryOrderStatus, customerID uint) ([]*model.DeliveryOrder, error)
}

// DeliveryOrderLine operations.
type DeliveryOrderLine interface {
	CreateDeliveryOrderLine(ctx context.Context, dol *model.DeliveryOrderLine) error
	GetDeliveryOrderLine(ctx context.Context, id uint) (*model.DeliveryOrderLine, error)
	ListDeliveryOrderLines(ctx context.Context, deliveryOrderID uint) ([]*model.DeliveryOrderLine, error)
	UpdateDeliveryOrderLinePickedQty(ctx context.Context, id uint, pickedQty float64) error
}

// PutawayRecord operations.
type PutawayRecord interface {
	CreatePutawayRecord(ctx context.Context, pr *model.PutawayRecord) error
	ListPutawayRecords(ctx context.Context, receiveOrderID uint) ([]*model.PutawayRecord, error)
}

// PickRecord operations.
type PickRecord interface {
	CreatePickRecord(ctx context.Context, pr *model.PickRecord) error
	ListPickRecords(ctx context.Context, deliveryOrderID uint) ([]*model.PickRecord, error)
}

// CountPlan operations.
type CountPlan interface {
	CreateCountPlan(ctx context.Context, cp *model.CountPlan) error
	GetCountPlan(ctx context.Context, id uint) (*model.CountPlan, error)
	UpdateCountPlan(ctx context.Context, cp *model.CountPlan) error
}

// CountRecord operations.
type CountRecord interface {
	CreateCountRecord(ctx context.Context, cr *model.CountRecord) error
	UpdateCountRecord(ctx context.Context, cr *model.CountRecord) error
	GetCountRecord(ctx context.Context, id uint) (*model.CountRecord, error)
}

// Repository is the aggregate repository combining all domain sub-repositories.
type Repository interface {
	TransactionManager
	Warehouse
	Location
	Inventory
	ReceiveOrder
	ReceiveOrderLine
	DeliveryOrder
	DeliveryOrderLine
	PutawayRecord
	PickRecord
	CountPlan
	CountRecord
}
