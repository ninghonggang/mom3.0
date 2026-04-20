package model

import (
	"time"
)

// SalesReturn 销售退货单
type SalesReturn struct {
	ID            uint      `json:"id" gorm:"primarykey"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	ReturnNo      string    `json:"return_no" gorm:"size:50;uniqueIndex;not null"`
	SalesOrderID  int64     `json:"sales_order_id" gorm:"index"`
	CustomerID    int64     `json:"customer_id" gorm:"index"`
	CustomerName  *string   `json:"customer_name" gorm:"size:100"`
	WarehouseID   int64     `json:"warehouse_id" gorm:"index"`
	ReturnDate    time.Time `json:"return_date" gorm:"type:date"`
	Status        string    `json:"status" gorm:"size:20;not null;default:PENDING"` // PENDING/APPROVED/RETURNING/RETURNED
	ReturnType    string    `json:"return_type" gorm:"size:20;not null"`           // NORMAL/EMERGENCY/GOODWILL
	RequestBy     *int64    `json:"request_by"`
	RequestTime   *time.Time `json:"request_time"`
	ApprovedBy    *int64    `json:"approved_by"`
	ApprovedTime  *time.Time `json:"approved_time"`
	ReturnedBy    *int64    `json:"returned_by"`
	ReturnedTime  *time.Time `json:"returned_time"`
	Remark        *string   `json:"remark" gorm:"type:text"`
	TenantID      int64     `json:"tenant_id" gorm:"index;not null"`
	CreatedBy     *string   `json:"created_by" gorm:"size:50"`
	Items         []SalesReturnItem `json:"items" gorm:"foreignKey:ReturnID"`
}

func (SalesReturn) TableName() string {
	return "wms_sales_return"
}

// SalesReturnItem 销售退货明细
type SalesReturnItem struct {
	ID           uint    `json:"id" gorm:"primarykey"`
	ReturnID     uint    `json:"return_id" gorm:"not null;index"`
	LineNo       int     `json:"line_no" gorm:"not null"`
	MaterialID   int64   `json:"material_id" gorm:"not null;index"`
	MaterialCode *string `json:"material_code" gorm:"size:50"`
	MaterialName *string `json:"material_name" gorm:"size:100"`
	Unit         *string `json:"unit" gorm:"size:20"`
	ReturnQty    float64 `json:"return_qty" gorm:"type:decimal(18,3);not null;default:0"` // 退货数量
	Reason       *string `json:"reason" gorm:"size:200"` // 退货原因
	TenantID     int64   `json:"tenant_id" gorm:"index;not null"`
}

func (SalesReturnItem) TableName() string {
	return "wms_sales_return_item"
}

// SalesReturnCreate 创建请求
type SalesReturnCreate struct {
	SalesOrderID  int64                      `json:"sales_order_id" binding:"required"`
	CustomerID    int64                      `json:"customer_id" binding:"required"`
	CustomerName  *string                    `json:"customer_name"`
	WarehouseID   int64                      `json:"warehouse_id" binding:"required"`
	ReturnDate    time.Time                  `json:"return_date" binding:"required"`
	ReturnType    string                     `json:"return_type" binding:"required"`
	Remark        *string                    `json:"remark"`
	Items         []SalesReturnItemCreate   `json:"items" binding:"required,min=1"`
}

// SalesReturnItemCreate 明细创建请求
type SalesReturnItemCreate struct {
	MaterialID   int64   `json:"material_id" binding:"required"`
	MaterialCode *string `json:"material_code"`
	MaterialName *string `json:"material_name"`
	Unit         *string `json:"unit"`
	ReturnQty   float64 `json:"return_qty" binding:"required"`
	Reason       *string `json:"reason"`
}

// SalesReturnUpdate 更新请求
type SalesReturnUpdate struct {
	Remark *string `json:"remark"`
}

// SalesReturnSubmit 提交退货单
type SalesReturnSubmit struct {
	Items []SalesReturnItemSubmit `json:"items"`
}

// SalesReturnItemSubmit 明细提交
type SalesReturnItemSubmit struct {
	MaterialID int64   `json:"material_id" binding:"required"`
	ReturnQty float64 `json:"return_qty" binding:"required"`
	Reason    *string `json:"reason"`
}

// SalesReturnConfirm 确认退货
type SalesReturnConfirm struct {
	Items []SalesReturnItemConfirm `json:"items"`
}

// SalesReturnItemConfirm 明细确认退货
type SalesReturnItemConfirm struct {
	MaterialID int64   `json:"material_id" binding:"required"`
	ReturnQty float64 `json:"return_qty" binding:"required"`
}