package model

import (
	"time"
)

// ========== 仓储管理模块 ==========

// Warehouse 仓库
type Warehouse struct {
	BaseModel
	TenantID    int64   `json:"tenant_id" gorm:"index;not null"`
	WarehouseCode string `json:"warehouse_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_warehouse"`
	WarehouseName string `json:"warehouse_name" gorm:"size:100;not null"`
	WarehouseType string `json:"warehouse_type" gorm:"size:20"` // 原料仓/成品仓/线边仓
	Address      *string `json:"address" gorm:"size:200"`
	Manager      *string `json:"manager" gorm:"size:50"`
	Phone        *string `json:"phone" gorm:"size:20"`
	Status       int     `json:"status" gorm:"default:1"`
}

func (Warehouse) TableName() string {
	return "wms_warehouse"
}

// Location 库位
type Location struct {
	BaseModel
	TenantID   int64   `json:"tenant_id" gorm:"index;not null"`
	LocationCode string `json:"location_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_location"`
	LocationName string `json:"location_name" gorm:"size:100"`
	WarehouseID int64   `json:"warehouse_id"`
	ZoneCode   *string `json:"zone_code" gorm:"size:20"` // A区/B区
	Row         *int    `json:"row"` // 排
	Col         *int    `json:"col"` // 列
	Layer       *int    `json:"layer"` // 层
	LocationType string `json:"location_type" gorm:"size:20"` // 存储/备货/发货
	Capacity    *int    `json:"capacity"` // 容量
	Status      int     `json:"status" gorm:"default:1"`
}

func (Location) TableName() string {
	return "wms_location"
}

// Inventory 库存
type Inventory struct {
	BaseModel
	TenantID     int64   `json:"tenant_id" gorm:"index;not null"`
	MaterialID   int64   `json:"material_id" gorm:"index"`
	MaterialCode string   `json:"material_code" gorm:"size:50"`
	MaterialName string   `json:"material_name" gorm:"size:100"`
	WarehouseID  int64   `json:"warehouse_id" gorm:"index"`
	LocationID   int64   `json:"location_id" gorm:"index"`
	Quantity     float64 `json:"quantity" gorm:"type:decimal(18,4);default:0"` // 库存数量
	AvailableQty float64 `json:"available_qty" gorm:"type:decimal(18,4);default:0"` // 可用数量
	AllocatedQty float64 `json:"allocated_qty" gorm:"type:decimal(18,4);default:0"` // 已分配数量
	LockedQty    float64 `json:"locked_qty" gorm:"type:decimal(18,4);default:0"` // 冻结数量
	BatchNo      *string `json:"batch_no" gorm:"size:50"` // 批次号
}

func (Inventory) TableName() string {
	return "wms_inventory"
}

// InventoryRecord 库存记录
type InventoryRecord struct {
	BaseModel
	TenantID     int64      `json:"tenant_id" gorm:"index;not null"`
	RecordNo     string     `json:"record_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_inv_rec"`
	RecordType   string     `json:"record_type" gorm:"size:20"` // 入库/出库/调整
	MaterialID   int64      `json:"material_id"`
	MaterialCode string     `json:"material_code" gorm:"size:50"`
	MaterialName string     `json:"material_name" gorm:"size:100"`
	WarehouseID  int64      `json:"warehouse_id"`
	LocationID   int64      `json:"location_id"`
	Quantity     float64    `json:"quantity" gorm:"type:decimal(18,4)"`
	BatchNo      *string    `json:"batch_no" gorm:"size:50"`
	SourceNo     *string    `json:"source_no" gorm:"size:50"` // 来源单号
	OperatorID   int64      `json:"operator_id"`
	OperatorName *string    `json:"operator_name" gorm:"size:50"`
	OperateTime  *time.Time `json:"operate_time"`
	Remark       *string    `json:"remark" gorm:"size:500"`
}

func (InventoryRecord) TableName() string {
	return "wms_inventory_record"
}

// ReceiveOrder 收货单
type ReceiveOrder struct {
	BaseModel
	TenantID     int64      `json:"tenant_id" gorm:"index;not null"`
	ReceiveNo    string     `json:"receive_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_recv"`
	SupplierID   int64      `json:"supplier_id"`
	SupplierName *string    `json:"supplier_name" gorm:"size:100"`
	WarehouseID  int64      `json:"warehouse_id"`
	ReceiveDate  *time.Time `json:"receive_date"`
	ReceiveUserID int64     `json:"receive_user_id"`
	Status       int        `json:"status" gorm:"default:1"` // 1待收货/2收货中/3已完成
	Remark       *string    `json:"remark" gorm:"size:500"`
}

func (ReceiveOrder) TableName() string {
	return "wms_receive_order"
}

// ReceiveOrderItem 收货单明细
type ReceiveOrderItem struct {
	BaseModel
	ReceiveID   int64   `json:"receive_id" gorm:"index"`
	MaterialID  int64   `json:"material_id"`
	MaterialCode string  `json:"material_code" gorm:"size:50"`
	MaterialName string  `json:"material_name" gorm:"size:100"`
	Quantity    float64 `json:"quantity" gorm:"type:decimal(18,4)"`
	ReceivedQty float64 `json:"received_qty" gorm:"type:decimal(18,4);default:0"`
	Unit       string  `json:"unit" gorm:"size:20"`
	BatchNo    *string `json:"batch_no" gorm:"size:50"`
}

func (ReceiveOrderItem) TableName() string {
	return "wms_receive_order_item"
}

// DeliveryOrder 发货单
type DeliveryOrder struct {
	BaseModel
	TenantID      int64      `json:"tenant_id" gorm:"index;not null"`
	DeliveryNo     string     `json:"delivery_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_delivery"`
	CustomerID    int64      `json:"customer_id"`
	CustomerName  *string    `json:"customer_name" gorm:"size:100"`
	WarehouseID   int64      `json:"warehouse_id"`
	DeliveryDate  *time.Time `json:"delivery_date"`
	DeliveryUserID int64     `json:"delivery_user_id"`
	Status        int        `json:"status" gorm:"default:1"` // 1待发货/2发货中/3已完成
	Remark        *string    `json:"remark" gorm:"size:500"`
}

func (DeliveryOrder) TableName() string {
	return "wms_delivery_order"
}

// DeliveryOrderItem 发货单明细
type DeliveryOrderItem struct {
	BaseModel
	DeliveryID   int64   `json:"delivery_id" gorm:"index"`
	MaterialID   int64   `json:"material_id"`
	MaterialCode string  `json:"material_code" gorm:"size:50"`
	MaterialName string  `json:"material_name" gorm:"size:100"`
	Quantity     float64 `json:"quantity" gorm:"type:decimal(18,4)"`
	ShippedQty   float64 `json:"shipped_qty" gorm:"type:decimal(18,4);default:0"`
	Unit        string  `json:"unit" gorm:"size:20"`
	BatchNo     *string `json:"batch_no" gorm:"size:50"`
}

func (DeliveryOrderItem) TableName() string {
	return "wms_delivery_order_item"
}

// SideLocation 线边库位
type SideLocation struct {
	BaseModel
	TenantID     int64   `json:"tenant_id" gorm:"index;not null"`
	LocationCode string  `json:"location_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_sideloc"`
	LocationName string  `json:"location_name" gorm:"size:100"`
	WorkshopID   int64   `json:"workshop_id"`
	WorkshopName *string `json:"workshop_name" gorm:"size:100"`
	LineID      *int64  `json:"line_id"`
	LineName    *string `json:"line_name" gorm:"size:100"`
	StationID   *int64  `json:"station_id"`
	StationName *string `json:"station_name" gorm:"size:100"`
	LocationType string  `json:"location_type" gorm:"size:20"` // 原料/成品/工装
	MaxCapacity *float64 `json:"max_capacity"` // 最大容量
	CurrentQty  float64  `json:"current_qty"` // 当前数量
	Status      int       `json:"status" gorm:"default:1"` // 1启用/2停用
	Remark      *string  `json:"remark" gorm:"size:500"`
}

func (SideLocation) TableName() string {
	return "wms_side_location"
}

// KanbanPull 看板拉动
type KanbanPull struct {
	BaseModel
	TenantID      int64     `json:"tenant_id" gorm:"index;not null"`
	KanbanNo     string    `json:"kanban_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_kanban"`
	MaterialID   int64     `json:"material_id"`
	MaterialCode string    `json:"material_code" gorm:"size:50"`
	MaterialName string    `json:"material_name" gorm:"size:100"`
	FromLocationID int64   `json:"from_location_id"` // 物料来源(线边库位)
	ToLocationID  int64    `json:"to_location_id"` // 物料去向(工位)
	KanbanQty    float64   `json:"kanban_qty"` // 看板数量
	TriggerQty   float64   `json:"trigger_qty"` // 触发数量
	CurrentQty   float64  `json:"current_qty"` // 当前库存
	LeadTime     int       `json:"lead_time"` // 提前期(分钟)
	Status       int       `json:"status" gorm:"default:1"` // 1启用/2停用
	Remark       *string   `json:"remark" gorm:"size:500"`
}

func (KanbanPull) TableName() string {
	return "wms_kanban_pull"
}
