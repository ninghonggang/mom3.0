package model

// TransferOrderItem 调拨单明细
type TransferOrderItem struct {
	BaseModel
	TenantID       int64   `json:"tenant_id" gorm:"index;not null"`
	TransferID    int64   `json:"transfer_id"`
	MaterialID    int64   `json:"material_id"`
	MaterialCode  string  `json:"material_code" gorm:"size:50"`
	MaterialName  string  `json:"material_name" gorm:"size:100"`
	Quantity      float64 `json:"quantity"` // 调拨数量
	OutQuantity   float64 `json:"out_quantity"` // 已出库数量
	InQuantity    float64 `json:"in_quantity"` // 已入库数量
	Unit          *string `json:"unit" gorm:"size:20"`
	BatchNo       *string `json:"batch_no" gorm:"size:50"`
	Remark        *string `json:"remark" gorm:"size:500"`
}

func (TransferOrderItem) TableName() string {
	return "wms_transfer_order_item"
}

// StockCheckItem 盘点明细
type StockCheckItem struct {
	BaseModel
	TenantID     int64   `json:"tenant_id" gorm:"index;not null"`
	CheckID      int64   `json:"check_id"`
	MaterialID   int64   `json:"material_id"`
	MaterialCode string  `json:"material_code" gorm:"size:50"`
	MaterialName string  `json:"material_name" gorm:"size:100"`
	LocationID   *int64  `json:"location_id"`
	LocationName *string `json:"location_name" gorm:"size:100"`
	BatchNo      *string `json:"batch_no" gorm:"size:50"`
	StockQty     float64 `json:"stock_qty"` // 账面数量
	CheckQty     float64 `json:"check_qty"` // 盘点数量
	DiffQty      float64 `json:"diff_qty"` // 差异数量
	Unit         *string `json:"unit" gorm:"size:20"`
	HandleResult *string `json:"handle_result" gorm:"size:100"` // 处理方式
	Remark       *string `json:"remark" gorm:"size:500"`
}

func (StockCheckItem) TableName() string {
	return "wms_stock_check_item"
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
