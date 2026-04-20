package model

import (
	"time"
)

// ProductionReturn 生产退料单
type ProductionReturn struct {
	ID                  uint      `json:"id" gorm:"primarykey"`
	CreatedAt           time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt           time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	ReturnNo            string    `json:"return_no" gorm:"size:50;uniqueIndex;not null"`
	ProductionOrderID   int64     `json:"production_order_id" gorm:"index"`
	OrderNo             *string   `json:"order_no" gorm:"size:50"`
	WorkstationID       *int64    `json:"workstation_id" gorm:"index"`
	WorkshopID          *int64    `json:"workshop_id" gorm:"index"`
	Status              string    `json:"status" gorm:"size:20;not null;default:PENDING"` // PENDING/APPROVED/RETURNING/RETURNED
	ReturnType          string    `json:"return_type" gorm:"size:20;not null"`             // NORMAL/EMERGENCY
	RequestBy           *int64    `json:"request_by"`
	RequestTime         *time.Time `json:"request_time"`
	ApprovedBy          *int64    `json:"approved_by"`
	ApprovedTime        *time.Time `json:"approved_time"`
	ReturnedBy          *int64    `json:"returned_by"`
	ReturnedTime        *time.Time `json:"returned_time"`
	Remark              *string   `json:"remark" gorm:"type:text"`
	TenantID            int64     `json:"tenant_id" gorm:"index;not null"`
	CreatedBy           *string   `json:"created_by" gorm:"size:50"`
	Items               []ProductionReturnItem `json:"items" gorm:"foreignKey:ReturnID"`
}

func (ProductionReturn) TableName() string {
	return "wms_production_return"
}

// ProductionReturnItem 生产退料明细
type ProductionReturnItem struct {
	ID            uint    `json:"id" gorm:"primarykey"`
	ReturnID      uint   `json:"return_id" gorm:"not null;index"`
	LineNo        int     `json:"line_no" gorm:"not null"`
	MaterialID    int64   `json:"material_id" gorm:"not null;index"`
	MaterialCode  *string `json:"material_code" gorm:"size:50"`
	MaterialName  *string `json:"material_name" gorm:"size:100"`
	Unit          *string `json:"unit" gorm:"size:20"`
	IssuedQty     float64 `json:"issued_qty" gorm:"type:decimal(18,3);not null"` // 发料数量
	ReturnQty     float64 `json:"return_qty" gorm:"type:decimal(18,3);default:0"` // 退料数量
	WarehouseID   *int64  `json:"warehouse_id"`
	LocationID    *int64  `json:"location_id"`
	BatchNo       *string `json:"batch_no" gorm:"size:50"`
	Reason        *string `json:"reason" gorm:"size:200"` // 退料原因
	Remark        *string `json:"remark" gorm:"type:text"`
	TenantID      int64   `json:"tenant_id" gorm:"index;not null"`
}

func (ProductionReturnItem) TableName() string {
	return "wms_production_return_item"
}

// ProductionReturnCreate 创建请求
type ProductionReturnCreate struct {
	ProductionOrderID int64 `json:"production_order_id" binding:"required"`
	ReturnType        string `json:"return_type" binding:"required"`
	WorkstationID    *int64 `json:"workstation_id"`
	WorkshopID       *int64 `json:"workshop_id"`
	Items            []ProductionReturnItemCreate `json:"items" binding:"required,min=1"`
}

// ProductionReturnItemCreate 明细创建请求
type ProductionReturnItemCreate struct {
	MaterialID  int64   `json:"material_id" binding:"required"`
	MaterialCode *string `json:"material_code"`
	MaterialName *string `json:"material_name"`
	Unit       *string `json:"unit"`
	IssuedQty  float64 `json:"issued_qty" binding:"required"`
	ReturnQty  float64 `json:"return_qty"`
	WarehouseID *int64  `json:"warehouse_id"`
	LocationID *int64  `json:"location_id"`
	BatchNo    *string `json:"batch_no"`
	Reason     *string `json:"reason"`
	Remark     *string `json:"remark"`
}

// ProductionReturnUpdate 更新请求
type ProductionReturnUpdate struct {
	WorkstationID *int64 `json:"workstation_id"`
	WorkshopID    *int64 `json:"workshop_id"`
	Remark        *string `json:"remark"`
}

// ProductionReturnSubmit 提交退料单
type ProductionReturnSubmit struct {
	Items []ProductionReturnItemSubmit `json:"items"`
}

// ProductionReturnItemSubmit 明细提交
type ProductionReturnItemSubmit struct {
	MaterialID  int64   `json:"material_id" binding:"required"`
	ReturnQty  float64 `json:"return_qty" binding:"required"`
	WarehouseID *int64 `json:"warehouse_id"`
	LocationID *int64  `json:"location_id"`
	BatchNo    *string `json:"batch_no"`
	Reason     *string `json:"reason"`
}

// ProductionReturnConfirm 确认退料
type ProductionReturnConfirm struct {
	Items []ProductionReturnItemConfirm `json:"items"`
}

// ProductionReturnItemConfirm 明细确认退料
type ProductionReturnItemConfirm struct {
	MaterialID  int64   `json:"material_id" binding:"required"`
	ReturnQty   float64 `json:"return_qty" binding:"required"`
	WarehouseID *int64 `json:"warehouse_id"`
	LocationID  *int64  `json:"location_id"`
	BatchNo     *string `json:"batch_no"`
}
