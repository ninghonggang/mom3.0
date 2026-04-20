package model

import "time"

// ========== 调拨管理扩展 ==========

// TransferOrder 调拨单
type TransferOrder struct {
	BaseModel
	TenantID            int64      `json:"tenant_id" gorm:"index;not null"`
	TransferNo          string     `json:"transfer_no" gorm:"size:50;uniqueIndex:idx_tenant_transfer"`
	TransferType        string     `json:"transfer_type" gorm:"size:20"` // TRANSFER/ADJUSTMENT/TRANSFER_IN/TRANSFER_OUT
	FromWarehouseID     int64      `json:"from_warehouse_id"`
	FromWarehouseName   string     `json:"from_warehouse_name" gorm:"size:100"`
	FromWarehouseType   string     `json:"from_warehouse_type" gorm:"size:20"`
	ToWarehouseID       int64      `json:"to_warehouse_id"`
	ToWarehouseName     string     `json:"to_warehouse_name" gorm:"size:100"`
	ToWarehouseType     string     `json:"to_warehouse_type" gorm:"size:20"`
	FromWorkstationID   *int64     `json:"from_workstation_id"`
	FromWorkstationName *string    `json:"from_workstation_name" gorm:"size:100"`
	ToWorkstationID     *int64     `json:"to_workstation_id"`
	ToWorkstationName   *string    `json:"to_workstation_name" gorm:"size:100"`
	TransferReason      string     `json:"transfer_reason" gorm:"size:100"` // 生产领料/退料/调拨申请/盘盈调账
	SourceOrderID       *int64     `json:"source_order_id"`
	SourceOrderNo       *string    `json:"source_order_no" gorm:"size:50"`
	Status              string     `json:"status" gorm:"size:20;default:DRAFT"` // DRAFT/SUBMITTED/APPROVED/IN_TRANSIT/COMPLETED/CANCELLED
	RequesterID         *int64     `json:"requester_id"`
	RequesterName       *string    `json:"requester_name" gorm:"size:50"`
	RequestTime         *time.Time `json:"request_time"`
	ApproverID          *int64     `json:"approver_id"`
	ApprovedTime        *time.Time `json:"approved_time"`
	ApprovalComment     *string    `json:"approval_comment" gorm:"type:text"`
	ActualStartTime     *time.Time `json:"actual_start_time"`
	ActualCompleteTime  *time.Time `json:"actual_complete_time"`
	OperatorID          *int64     `json:"operator_id"`
	OperatorName        *string    `json:"operator_name" gorm:"size:50"`
	LogisticsProvider   *string    `json:"logistics_provider" gorm:"size:100"`
	TrackingNo          *string    `json:"tracking_no" gorm:"size:100"`
	TotalAmount         float64    `json:"total_amount" gorm:"type:decimal(18,2);default:0"`
	Currency            string     `json:"currency" gorm:"size:10;default:CNY"`
	Remark              *string    `json:"remark" gorm:"type:text"`
	WorkshopID          *int64     `json:"workshop_id"`
}

func (TransferOrder) TableName() string {
	return "wms_transfer_order"
}

// TransferOrderItem 调拨单明细
type TransferOrderItem struct {
	BaseModel
	TenantID          int64   `json:"tenant_id" gorm:"index;not null"`
	TransferID        int64   `json:"transfer_id"`
	LineNo            int     `json:"line_no"`
	MaterialID        int64   `json:"material_id"`
	MaterialCode      string  `json:"material_code" gorm:"size:50"`
	MaterialName      string  `json:"material_name" gorm:"size:100"`
	Specification     *string `json:"specification" gorm:"size:200"`
	Unit              *string `json:"unit" gorm:"size:20"`
	RequestQty        float64 `json:"request_qty" gorm:"type:decimal(18,3)"`     // 申请数量
	ApprovedQty       float64 `json:"approved_qty" gorm:"type:decimal(18,3)"`    // 审批数量
	TransferQty       float64 `json:"transfer_qty" gorm:"type:decimal(18,3)"`     // 调拨数量
	TransferredQty    float64 `json:"transferred_qty" gorm:"type:decimal(18,3);default:0"` // 已调拨数量
	BatchNo           *string `json:"batch_no" gorm:"size:50"`
	ProductionDate    *string `json:"production_date" gorm:"size:10"`
	ExpiryDate        *string `json:"expiry_date" gorm:"size:10"`
	FromLocationID    *int64  `json:"from_location_id"`
	FromLocationName  *string `json:"from_location_name" gorm:"size:100"`
	ToLocationID      *int64  `json:"to_location_id"`
	ToLocationName    *string `json:"to_location_name" gorm:"size:100"`
	RequireQC         int     `json:"require_qc" gorm:"default:0"` // 是否需要质检
	QCStatus          *string `json:"qc_status" gorm:"size:20"`    // PENDING/PASS/FAIL
	UnitCost          float64 `json:"unit_cost" gorm:"type:decimal(18,4)"`
	TotalCost         float64 `json:"total_cost" gorm:"type:decimal(18,2)"`
	Remark            *string `json:"remark" gorm:"size:500"`
}

func (TransferOrderItem) TableName() string {
	return "wms_transfer_order_item"
}

// TransferLot 调拨批次跟踪
type TransferLot struct {
	BaseModel
	TenantID        int64      `json:"tenant_id" gorm:"index;not null"`
	TransferItemID  int64      `json:"transfer_item_id"`
	LotNo           string     `json:"lot_no" gorm:"size:50"`             // 调拨批次号
	TransferQty     float64    `json:"transfer_qty" gorm:"type:decimal(18,3)"`
	TransferredQty  float64    `json:"transferred_qty" gorm:"type:decimal(18,3);default:0"`
	Status          string     `json:"status" gorm:"size:20;default:PENDING"` // PENDING/IN_TRANSIT/ARRIVED/COMPLETED
	ShipTime        *time.Time `json:"ship_time"`
	ArriveTime      *time.Time `json:"arrive_time"`
	Remark          *string    `json:"remark" gorm:"size:200"`
}

func (TransferLot) TableName() string {
	return "wms_transfer_lot"
}

// TransferTrace 调拨跟踪记录
type TransferTrace struct {
	BaseModel
	TenantID        int64      `json:"tenant_id" gorm:"index;not null"`
	TransferOrderID int64      `json:"transfer_order_id"`
	TransferItemID  *int64     `json:"transfer_item_id"`
	Action          string     `json:"action" gorm:"size:50"` // SUBMIT/APPROVE/TRANSFER_OUT/IN_TRANSIT/TRANSFER_IN/COMPLETE/CANCEL
	OperatorID      int64      `json:"operator_id"`
	OperatorName    string     `json:"operator_name" gorm:"size:50"`
	OperateTime     *time.Time `json:"operate_time"`
	LocationID      *int64     `json:"location_id"`
	LocationName    *string    `json:"location_name" gorm:"size:100"`
	Remark          *string    `json:"remark" gorm:"type:text"`
}

func (TransferTrace) TableName() string {
	return "wms_transfer_trace"
}
