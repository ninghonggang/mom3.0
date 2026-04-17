package model

import (
	"time"
)

// PickList 配料单
type PickList struct {
	BaseModel
	TenantID       int64      `json:"tenant_id" gorm:"index;not null"`
	PickNo        string     `json:"pick_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_pick"`
	OrderID       int64      `json:"order_id"` // 关联工单ID
	OrderNo       string     `json:"order_no" gorm:"size:50"` // 工单号
	WorkshopID    int64      `json:"workshop_id"`
	WorkshopName  *string    `json:"workshop_name" gorm:"size:100"`
	Status        int        `json:"status" gorm:"default:1"` // 1待配料/2配料中/3已完成
	PickerID      *int64     `json:"picker_id"`
	PickerName    *string    `json:"picker_name" gorm:"size:50"`
	PickTime     *time.Time `json:"pick_time"`
	Remark       *string    `json:"remark" gorm:"size:500"`
}

func (PickList) TableName() string {
	return "wms_pick_list"
}

// PickListItem 配料单明细
type PickListItem struct {
	BaseModel
	PickID       int64   `json:"pick_id" gorm:"index"`
	MaterialID   int64   `json:"material_id"`
	MaterialCode string  `json:"material_code" gorm:"size:50"`
	MaterialName string  `json:"material_name" gorm:"size:100"`
	Unit         string  `json:"unit" gorm:"size:20"`
	RequiredQty  float64 `json:"required_qty" gorm:"type:decimal(18,4)"` // 需求数量
	PickedQty    float64 `json:"picked_qty" gorm:"type:decimal(18,4);default:0"` // 已配数量
	WarehouseID  int64   `json:"warehouse_id"`
	WarehouseName *string `json:"warehouse_name" gorm:"size:100"`
	LocationID   int64   `json:"location_id"`
	LocationCode *string `json:"location_code" gorm:"size:50"`
	BatchNo      *string `json:"batch_no" gorm:"size:50"`
	Remark       *string `json:"remark" gorm:"size:500"`
}

func (PickListItem) TableName() string {
	return "wms_pick_list_item"
}

// GoodsReceipt 完工收货单（MES工单完工后WMS自动生成）
type GoodsReceipt struct {
	BaseModel
	TenantID       int64      `json:"tenant_id" gorm:"index;not null"`
	ReceiptNo      string     `json:"receipt_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_receipt"`
	SourceType     string     `json:"source_type" gorm:"size:20"` // MES_ORDER
	SourceID       int64      `json:"source_id"` // 来源单ID（工单ID）
	SourceNo       string     `json:"source_no" gorm:"size:50"` // 来源单号（工单号）
	MaterialID     int64      `json:"material_id"`
	MaterialCode   string     `json:"material_code" gorm:"size:50"`
	MaterialName   string     `json:"material_name" gorm:"size:100"`
	WarehouseID    int64      `json:"warehouse_id"`
	WarehouseName  *string    `json:"warehouse_name" gorm:"size:100"`
	LocationID     int64      `json:"location_id"`
	LocationCode   *string    `json:"location_code" gorm:"size:50"`
	Quantity       float64    `json:"quantity" gorm:"type:decimal(18,4)"` // 收货数量
	QualifiedQty   float64    `json:"qualified_qty" gorm:"type:decimal(18,4);default:0"` // 合格数量
	RejectedQty    float64    `json:"rejected_qty" gorm:"type:decimal(18,4);default:0"` // 不良数量
	ReceiptDate    *time.Time `json:"receipt_date"`
	ReceiverID     *int64     `json:"receiver_id"`
	ReceiverName   *string    `json:"receiver_name" gorm:"size:50"`
	Status         int        `json:"status" gorm:"default:1"` // 1待入库/2入库中/3已完成
	InspectStatus  int        `json:"inspect_status" gorm:"default:0"` // 0无需检验/1待检验/2检验中/3已检验
	Remark         *string    `json:"remark" gorm:"size:500"`
}

func (GoodsReceipt) TableName() string {
	return "wms_goods_receipt"
}

// EventLogPayload 事件载荷基础结构
type EventLogPayload struct {
	TenantID    int64                  `json:"tenant_id"`
	OperatorID  int64                  `json:"operator_id"`
	OperatorName string                `json:"operator_name"`
	Remark      string                 `json:"remark"`
}

// MESOrderStartedPayload MES工单开工事件载荷
type MESOrderStartedPayload struct {
	EventLogPayload
	OrderID     int64   `json:"order_id"`
	OrderNo     string  `json:"order_no"`
	MaterialID  int64   `json:"material_id"`
	MaterialCode string `json:"material_code"`
	MaterialName string `json:"material_name"`
	Quantity    float64 `json:"quantity"`
	WorkshopID  int64   `json:"workshop_id"`
	WorkshopName string `json:"workshop_name"`
}

// MESOrderReportedPayload MES工单报工完成事件载荷
type MESOrderReportedPayload struct {
	EventLogPayload
	OrderID     int64   `json:"order_id"`
	OrderNo     string  `json:"order_no"`
	ReportQty   float64 `json:"report_qty"`
	QualifiedQty float64 `json:"qualified_qty"`
	RejectedQty float64 `json:"rejected_qty"`
}

// WMSPurchaseReceivedPayload WMS采购收货事件载荷
type WMSPurchaseReceivedPayload struct {
	EventLogPayload
	ReceiveID   int64   `json:"receive_id"`
	ReceiveNo   string  `json:"receive_no"`
	SupplierID  int64   `json:"supplier_id"`
	SupplierName string `json:"supplier_name"`
	Items       []struct {
		MaterialID   int64   `json:"material_id"`
		MaterialCode string  `json:"material_code"`
		MaterialName string  `json:"material_name"`
		Quantity     float64 `json:"quantity"`
		BatchNo      string  `json:"batch_no"`
	} `json:"items"`
}

// QMSInspectionCompletedPayload QMS检验完成事件载荷
type QMSInspectionCompletedPayload struct {
	EventLogPayload
	IQCID       int64   `json:"iqc_id"`
	IQCNo       string  `json:"iqc_no"`
	ReceiveNo   string  `json:"receive_no"` // 关联收货单号
	Result      int     `json:"result"` // 2合格/3不合格
	QualifiedQty float64 `json:"qualified_qty"`
	RejectedQty float64 `json:"rejected_qty"`
}
