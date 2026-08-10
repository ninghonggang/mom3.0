package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// --- Status type definitions ---

// WarehouseStatus represents the status of a warehouse.
type WarehouseStatus string

const (
	WarehouseStatusActive   WarehouseStatus = "ACTIVE"
	WarehouseStatusInactive WarehouseStatus = "INACTIVE"
)

// LocationType represents the type of a location.
type LocationType string

const (
	LocationTypePick    LocationType = "PICK"
	LocationTypeStorage LocationType = "STORAGE"
	LocationTypeInbound LocationType = "INBOUND"
	LocationTypeOutbound LocationType = "OUTBOUND"
)

// LocationStatus represents the status of a location.
type LocationStatus string

const (
	LocationStatusActive   LocationStatus = "ACTIVE"
	LocationStatusInactive LocationStatus = "INACTIVE"
	LocationStatusFull     LocationStatus = "FULL"
)

// InventoryStatus represents the status of an inventory balance.
type InventoryStatus string

const (
	InventoryStatusNormal  InventoryStatus = "NORMAL"
	InventoryStatusLocked  InventoryStatus = "LOCKED"
	InventoryStatusExpired InventoryStatus = "EXPIRED"
)

// ReceiveOrderStatus represents the status of a receive order.
type ReceiveOrderStatus string

const (
	ReceiveStatusDraft     ReceiveOrderStatus = "DRAFT"
	ReceiveStatusReceiving ReceiveOrderStatus = "RECEIVING"
	ReceiveStatusReceived  ReceiveOrderStatus = "RECEIVED"
	ReceiveStatusPutaway   ReceiveOrderStatus = "PUTAWAY"
	ReceiveStatusCompleted ReceiveOrderStatus = "COMPLETED"
	ReceiveStatusCancelled ReceiveOrderStatus = "CANCELLED"
)

// DeliveryOrderStatus represents the status of a delivery order.
type DeliveryOrderStatus string

const (
	DeliveryStatusDraft   DeliveryOrderStatus = "DRAFT"
	DeliveryStatusPicking DeliveryOrderStatus = "PICKING"
	DeliveryStatusPicked  DeliveryOrderStatus = "PICKED"
	DeliveryStatusPacking DeliveryOrderStatus = "PACKING"
	DeliveryStatusOnHold  DeliveryOrderStatus = "ON_HOLD"
	DeliveryStatusShipped DeliveryOrderStatus = "SHIPPED"
	DeliveryStatusCancelled DeliveryOrderStatus = "CANCELLED"
)

// CountPlanStatus represents the status of a count plan.
type CountPlanStatus string

const (
	CountPlanStatusDraft    CountPlanStatus = "DRAFT"
	CountPlanStatusActive   CountPlanStatus = "ACTIVE"
	CountPlanStatusCompleted CountPlanStatus = "COMPLETED"
	CountPlanStatusCancelled CountPlanStatus = "CANCELLED"
)

// --- Models ---

// Warehouse represents a warehouse (仓库).
type Warehouse struct {
	ID            uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID      uint            `gorm:"index;not null" json:"tenant_id"`
	WarehouseCode string          `gorm:"size:50;uniqueIndex;not null" json:"warehouse_code"`
	WarehouseName string          `gorm:"size:200;not null" json:"warehouse_name"`
	WarehouseType string          `gorm:"size:50" json:"warehouse_type"`
	Status        WarehouseStatus `gorm:"size:20;default:ACTIVE" json:"status"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

func (Warehouse) TableName() string { return "wms_warehouses" }

// Location represents a storage location (库位).
type Location struct {
	ID           uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	WarehouseID  uint           `gorm:"index;not null" json:"warehouse_id"`
	AreaID       uint           `gorm:"index" json:"area_id"`
	LocationCode string        `gorm:"size:50;uniqueIndex;not null" json:"location_code"`
	LocationType LocationType  `gorm:"size:20;not null" json:"location_type"`
	Capacity     float64        `gorm:"default:0" json:"capacity"`
	UsedCapacity float64        `gorm:"default:0" json:"used_capacity"`
	Status       LocationStatus `gorm:"size:20;default:ACTIVE" json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

func (Location) TableName() string { return "wms_locations" }

// InventoryBalance represents the inventory ledger (库存台账).
type InventoryBalance struct {
	ID           uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID     uint            `gorm:"index;not null" json:"tenant_id"`
	MaterialID   uint            `gorm:"index;not null" json:"material_id"`
	LocationID   uint            `gorm:"index;not null" json:"location_id"`
	BatchNo      string          `gorm:"size:100;index" json:"batch_no"`
	Quantity     float64         `gorm:"not null;default:0" json:"quantity"`
	LockedQty    float64         `gorm:"not null;default:0" json:"locked_qty"`
	AvailableQty float64         `gorm:"not null;default:0" json:"available_qty"`
	ExpireDate   *time.Time      `json:"expire_date"`
	Status       InventoryStatus `gorm:"size:20;default:NORMAL" json:"status"`
	UnitCost     float64         `gorm:"default:0" json:"unit_cost"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

func (InventoryBalance) TableName() string { return "wms_inventory_balances" }

// BeforeSave enforces the balance invariant: available_qty = quantity - locked_qty.
// This hook fires on struct-based Create/Save/Updates. Atomic SQL-expression
// updates (LockInventory/UnlockInventory) skip hooks via Session(SkipHooks).
func (ib *InventoryBalance) BeforeSave(tx *gorm.DB) error {
	ib.AvailableQty = ib.Quantity - ib.LockedQty
	if ib.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	if ib.LockedQty < 0 {
		return errors.New("locked_qty cannot be negative")
	}
	if ib.AvailableQty < 0 {
		return errors.New("available_qty cannot be negative: locked_qty exceeds quantity")
	}
	return nil
}

// ReceiveOrder represents an inbound receive order (入库单).
type ReceiveOrder struct {
	ID          uint               `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID    uint               `gorm:"index;not null" json:"tenant_id"`
	ReceiveNo   string             `gorm:"size:50;uniqueIndex;not null" json:"receive_no"`
	PoID        string             `gorm:"size:50" json:"po_id"`
	SupplierID  uint               `json:"supplier_id"`
	Status      ReceiveOrderStatus `gorm:"size:20;default:DRAFT" json:"status"`
	ReceivedAt  *time.Time         `json:"received_at"`
	PutawayAt   *time.Time         `json:"putaway_at"`
	CompletedAt *time.Time         `json:"completed_at"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

func (ReceiveOrder) TableName() string { return "wms_receive_orders" }

// ReceiveOrderLine represents a line item on a receive order (入库单行).
type ReceiveOrderLine struct {
	ID            uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	ReceiveOrderID uint      `gorm:"index;not null" json:"receive_order_id"`
	MaterialID    uint       `gorm:"not null" json:"material_id"`
	ExpectedQty   float64    `gorm:"not null" json:"expected_qty"`
	ReceivedQty   float64    `gorm:"default:0" json:"received_qty"`
	// UnitPrice 为采购单价，上架时按移动加权平均并入 InventoryBalance.UnitCost。
	UnitPrice     float64    `gorm:"default:0" json:"unit_price"`
	BatchNo       string     `gorm:"size:100" json:"batch_no"`
	ExpireDate    *time.Time `json:"expire_date"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

func (ReceiveOrderLine) TableName() string { return "wms_receive_order_lines" }

// DeliveryOrder represents an outbound delivery order (出库单).
type DeliveryOrder struct {
	ID         uint                `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID   uint                `gorm:"index;not null" json:"tenant_id"`
	DeliveryNo string              `gorm:"size:50;uniqueIndex;not null" json:"delivery_no"`
	SoID       string              `gorm:"size:50" json:"so_id"`
	CustomerID uint                `json:"customer_id"`
	Status     DeliveryOrderStatus `gorm:"size:20;default:DRAFT" json:"status"`
	ShippedAt  *time.Time          `json:"shipped_at"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
}

func (DeliveryOrder) TableName() string { return "wms_delivery_orders" }

// DeliveryOrderLine represents a line item on a delivery order (出库单行).
type DeliveryOrderLine struct {
	ID              uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	DeliveryOrderID uint    `gorm:"index;not null" json:"delivery_order_id"`
	MaterialID      uint    `gorm:"not null" json:"material_id"`
	OrderedQty      float64 `gorm:"not null" json:"ordered_qty"`
	PickedQty       float64 `gorm:"default:0" json:"picked_qty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (DeliveryOrderLine) TableName() string { return "wms_delivery_order_lines" }

// PutawayRecord records the placement of received goods onto shelves (上架记录).
type PutawayRecord struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ReceiveOrderID uint     `gorm:"index;not null" json:"receive_order_id"`
	LocationID    uint      `gorm:"not null" json:"location_id"`
	MaterialID    uint      `gorm:"not null" json:"material_id"`
	Quantity      float64   `gorm:"not null" json:"quantity"`
	BatchNo       string    `gorm:"size:100" json:"batch_no"`
	PutawayTime   time.Time `gorm:"not null" json:"putaway_time"`
	CreatedAt     time.Time `json:"created_at"`
}

func (PutawayRecord) TableName() string { return "wms_putaway_records" }

// PickRecord records a picking action for a delivery order (拣货记录).
type PickRecord struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	DeliveryOrderID uint      `gorm:"index;not null" json:"delivery_order_id"`
	LocationID      uint      `gorm:"not null" json:"location_id"`
	MaterialID      uint      `gorm:"not null" json:"material_id"`
	PickedQty       float64   `gorm:"not null" json:"picked_qty"`
	PickerID        uint      `json:"picker_id"`
	BatchNo         string    `gorm:"size:100" json:"batch_no"`
	PickTime        time.Time `gorm:"not null" json:"pick_time"`
	CreatedAt       time.Time `json:"created_at"`
}

func (PickRecord) TableName() string { return "wms_pick_records" }

// CountPlan represents an inventory count plan (盘点计划).
type CountPlan struct {
	ID         uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	PlanNo     string          `gorm:"size:50;uniqueIndex;not null" json:"plan_no"`
	WarehouseID uint           `gorm:"index;not null" json:"warehouse_id"`
	PlanType   string          `gorm:"size:50" json:"plan_type"`
	Status     CountPlanStatus `gorm:"size:20;default:DRAFT" json:"status"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}

func (CountPlan) TableName() string { return "wms_count_plans" }

// CountRecord records the result of a count action (盘点记录).
type CountRecord struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	JobID      uint      `gorm:"index;not null" json:"job_id"`
	MaterialID uint      `gorm:"not null" json:"material_id"`
	LocationID uint      `gorm:"not null" json:"location_id"`
	SystemQty  float64   `json:"system_qty"`
	ActualQty  float64   `json:"actual_qty"`
	DiffQty    float64   `json:"diff_qty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (CountRecord) TableName() string { return "wms_count_records" }

// AllModels returns all model structs for auto-migration.
func AllModels() []interface{} {
	return []interface{}{
		&Warehouse{},
		&Location{},
		&InventoryBalance{},
		&ReceiveOrder{},
		&ReceiveOrderLine{},
		&DeliveryOrder{},
		&DeliveryOrderLine{},
		&PutawayRecord{},
		&PickRecord{},
		&CountPlan{},
		&CountRecord{},
	}
}
