package model

import (
	"time"

	"mom-server/internal/pkg/status"
)

// ProductionComplete 完工记录
type ProductionComplete struct {
	ID                uint      `json:"id" gorm:"primarykey"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	CompleteNo        string    `json:"complete_no" gorm:"size:50;uniqueIndex;not null"`
	ProductionOrderID int64     `json:"production_order_id" gorm:"index"`
	OrderNo          *string   `json:"order_no" gorm:"size:50"`
	WorkshopID       *int64    `json:"workshop_id" gorm:"index"`
	WorkstationID    *int64    `json:"workstation_id" gorm:"index"`
	CompleteQty       float64   `json:"complete_qty" gorm:"type:decimal(18,3);not null;default:0"`
	QualifiedQty      float64   `json:"qualified_qty" gorm:"type:decimal(18,3);default:0"`
	Status            string    `json:"status" gorm:"size:20;not null;default:PENDING"` // PENDING/INSPECTING/QUALIFIED/STORED(legacy,保留双轨)
	StatusV2          string    `json:"status_v2" gorm:"size:30;index"`               // 双轨:varchar - MOM 3.0 V2.1(P2 扩展 batch 2)
	CompleteTime      *time.Time `json:"complete_time"`
	OperatorID        *int64    `json:"operator_id"`
	OperatorName      *string   `json:"operator_name" gorm:"size:50"`
	Remark            *string   `json:"remark" gorm:"type:text"`
	TenantID          int64     `json:"tenant_id" gorm:"index;not null"`
	Items             []ProductionCompleteItem `json:"items" gorm:"foreignKey:CompleteID"`
}

func (ProductionComplete) TableName() string {
	return "wms_production_complete"
}

// StatusCode 返回有效状态码(优先 status_v2, fallback status → 字典码)
// P2 扩展 batch 2,与 MdmBOM.StatusCode() 同模式
func (c *ProductionComplete) StatusCode() string {
	if c.StatusV2 != "" {
		return c.StatusV2
	}
	return string(status.ProductionCompleteFromLegacyVarchar(c.Status))
}

// SetStatus 同时写 status + status_v2(双轨制)
func (c *ProductionComplete) SetStatus(code string) {
	cc := status.Code(code)
	if !cc.IsValid(status.ProductionCompleteAll) {
		code = string(status.ProductionCompletePending)
	}
	c.Status = code
	c.StatusV2 = code
}

// ProductionCompleteItem 完工明细
type ProductionCompleteItem struct {
	ID           uint    `json:"id" gorm:"primarykey"`
	CompleteID   int64   `json:"complete_id" gorm:"not null;index"`
	LineNo       int     `json:"line_no" gorm:"not null"`
	MaterialID   int64   `json:"material_id" gorm:"not null;index"`
	MaterialCode *string `json:"material_code" gorm:"size:50"`
	MaterialName *string `json:"material_name" gorm:"size:100"`
	Unit         *string `json:"unit" gorm:"size:20"`
	CompleteQty  float64 `json:"complete_qty" gorm:"type:decimal(18,3);not null;default:0"`
	QualifiedQty float64 `json:"qualified_qty" gorm:"type:decimal(18,3);default:0"`
	WarehouseID  *int64  `json:"warehouse_id"`
	LocationID   *int64  `json:"location_id"`
	BatchNo      *string `json:"batch_no" gorm:"size:50"`
	TenantID     int64   `json:"tenant_id" gorm:"index;not null"`
}

func (ProductionCompleteItem) TableName() string {
	return "wms_production_complete_item"
}

// ProductionStockIn 入库记录
type ProductionStockIn struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	StockInNo  string    `json:"stock_in_no" gorm:"size:50;uniqueIndex;not null"`
	CompleteID  int64     `json:"complete_id" gorm:"index;not null"`
	CompleteNo  *string   `json:"complete_no" gorm:"size:50"`
	WarehouseID int64     `json:"warehouse_id" gorm:"index"`
	LocationID  *int64    `json:"location_id" gorm:"index"`
	Status      string    `json:"status" gorm:"size:20;not null;default:PENDING"` // PENDING/STORED
	StockInTime *time.Time `json:"stock_in_time"`
	OperatorID  *int64    `json:"operator_id"`
	OperatorName *string   `json:"operator_name" gorm:"size:50"`
	TenantID    int64     `json:"tenant_id" gorm:"index;not null"`
	Items       []ProductionStockInItem `json:"items" gorm:"foreignKey:StockInID"`
}

func (ProductionStockIn) TableName() string {
	return "wms_production_stock_in"
}

// ProductionStockInItem 入库明细
type ProductionStockInItem struct {
	ID           uint    `json:"id" gorm:"primarykey"`
	StockInID    int64   `json:"stock_in_id" gorm:"not null;index"`
	LineNo       int     `json:"line_no" gorm:"not null"`
	MaterialID   int64   `json:"material_id" gorm:"not null;index"`
	MaterialCode *string `json:"material_code" gorm:"size:50"`
	MaterialName *string `json:"material_name" gorm:"size:100"`
	Unit         *string `json:"unit" gorm:"size:20"`
	StockInQty   float64 `json:"stock_in_qty" gorm:"type:decimal(18,3);not null;default:0"`
	WarehouseID  *int64  `json:"warehouse_id"`
	LocationID   *int64  `json:"location_id"`
	BatchNo      *string `json:"batch_no" gorm:"size:50"`
	TenantID     int64   `json:"tenant_id" gorm:"index;not null"`
}

func (ProductionStockInItem) TableName() string {
	return "wms_production_stock_in_item"
}

// ProductionCompleteCreate 创建完工单请求
type ProductionCompleteCreate struct {
	ProductionOrderID int64                               `json:"production_order_id" binding:"required"`
	WorkstationID     *int64                              `json:"workstation_id"`
	WorkshopID        *int64                               `json:"workshop_id"`
	CompleteQty       float64                              `json:"complete_qty" binding:"required"`
	QualifiedQty      float64                              `json:"qualified_qty"`
	Items            []ProductionCompleteItemCreate        `json:"items" binding:"required,min=1"`
}

// ProductionCompleteItemCreate 完工明细创建请求
type ProductionCompleteItemCreate struct {
	MaterialID   int64   `json:"material_id" binding:"required"`
	MaterialCode *string `json:"material_code"`
	MaterialName *string `json:"material_name"`
	Unit         *string `json:"unit"`
	CompleteQty  float64 `json:"complete_qty" binding:"required"`
	QualifiedQty float64 `json:"qualified_qty"`
	WarehouseID  *int64  `json:"warehouse_id"`
	LocationID   *int64  `json:"location_id"`
	BatchNo      *string `json:"batch_no"`
}

// ProductionCompleteUpdate 更新完工单请求
type ProductionCompleteUpdate struct {
	WorkstationID *int64  `json:"workstation_id"`
	WorkshopID   *int64  `json:"workshop_id"`
	Remark       *string `json:"remark"`
}

// ProductionCompleteSubmitForInspect 提交质检请求
type ProductionCompleteSubmitForInspect struct {
	Items []ProductionCompleteItemInspect `json:"items"`
}

// ProductionCompleteItemInspect 完工明细质检请求
type ProductionCompleteItemInspect struct {
	MaterialID   int64   `json:"material_id" binding:"required"`
	QualifiedQty float64 `json:"qualified_qty" binding:"required"`
	WarehouseID  *int64  `json:"warehouse_id"`
	LocationID   *int64  `json:"location_id"`
	BatchNo      *string `json:"batch_no"`
}

// ProductionStockInCreate 创建入库单请求
type ProductionStockInCreate struct {
	CompleteID  int64                              `json:"complete_id" binding:"required"`
	WarehouseID int64                              `json:"warehouse_id" binding:"required"`
	LocationID  *int64                             `json:"location_id"`
	Items       []ProductionStockInItemCreate      `json:"items" binding:"required,min=1"`
}

// ProductionStockInItemCreate 入库明细创建请求
type ProductionStockInItemCreate struct {
	MaterialID   int64   `json:"material_id" binding:"required"`
	MaterialCode *string `json:"material_code"`
	MaterialName *string `json:"material_name"`
	Unit         *string `json:"unit"`
	StockInQty   float64 `json:"stock_in_qty" binding:"required"`
	WarehouseID  *int64  `json:"warehouse_id"`
	LocationID   *int64  `json:"location_id"`
	BatchNo      *string `json:"batch_no"`
}
