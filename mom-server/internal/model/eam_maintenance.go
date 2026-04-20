package model

import (
	"time"
)

// EquipmentSpare 备件
type EquipmentSpare struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	TenantID      int64     `gorm:"index;not null;default:1" json:"tenant_id"`
	SpareCode     string    `gorm:"size:50;uniqueIndex" json:"spare_code"` // 备件编码
	SpareName     string    `gorm:"size:200" json:"spare_name"`           // 备件名称
	Category      string    `gorm:"size:50" json:"category"`              // 备件类别
	Specification string    `gorm:"size:200" json:"specification"`       // 规格型号
	Unit          string    `gorm:"size:20" json:"unit"`                 // 单位
	Quantity      float64   `gorm:"default:0" json:"quantity"`           // 当前库存
	MinQuantity   float64   `gorm:"default:0" json:"min_quantity"`      // 最小库存
	MaxQuantity   float64   `gorm:"default:0" json:"max_quantity"`      // 最大库存
	Location      string    `gorm:"size:200" json:"location"`            // 存放位置
	UnitPrice     float64   `json:"unit_price"`                          // 单价
	Status        string    `gorm:"size:20;default:'AVAILABLE'" json:"status"` // AVAILABLE/RESERVED/USED/SCRAPPED
	Remark        string    `gorm:"type:text" json:"remark"`             // 备注
	CreatedBy     string    `gorm:"size:100" json:"created_by"`
	UpdatedBy     string    `gorm:"size:100" json:"updated_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// TableName 表名
func (EquipmentSpare) TableName() string {
	return "eam_spare"
}

// EquipmentSpareTransaction 备件事务
type EquipmentSpareTransaction struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	TenantID        int64     `gorm:"index;not null;default:1" json:"tenant_id"`
	SpareID         uint      `gorm:"index;not null" json:"spare_id"`       // 备件ID
	SpareCode       string    `gorm:"size:50" json:"spare_code"`           // 备件编码
	TransactionType string    `gorm:"size:20;not null" json:"transaction_type"` // IN/OUT/APPLY/RESERVE
	Quantity        float64   `gorm:"not null" json:"quantity"`           // 数量
	BeforeQty       float64   `json:"before_qty"`                         // 变动前库存
	AfterQty        float64   `json:"after_qty"`                          // 变动后库存
	OrderNo         string    `gorm:"size:50" json:"order_no"`           // 相关单号
	HandlerID       uint      `json:"handler_id"`                         // 经办人ID
	HandlerName     string    `gorm:"size:100" json:"handler_name"`      // 经办人
	Remark          string    `gorm:"type:text" json:"remark"`           // 备注
	CreatedAt       time.Time `json:"created_at"`
}

// TableName 表名
func (EquipmentSpareTransaction) TableName() string {
	return "eam_spare_transaction"
}

// SpareCreateReq 备件创建请求
type SpareCreateReq struct {
	SpareCode     string  `json:"spare_code" binding:"required"`
	SpareName     string  `json:"spare_name" binding:"required"`
	Category      string  `json:"category"`
	Specification string  `json:"specification"`
	Unit          string  `json:"unit"`
	Quantity      float64 `json:"quantity"`
	MinQuantity   float64 `json:"min_quantity"`
	MaxQuantity   float64 `json:"max_quantity"`
	Location      string  `json:"location"`
	UnitPrice     float64 `json:"unit_price"`
	Remark        string  `json:"remark"`
}

// SpareUpdateReq 备件更新请求
type SpareUpdateReq struct {
	ID            uint    `json:"id" binding:"required"`
	SpareName     string  `json:"spare_name"`
	Category      string  `json:"category"`
	Specification string  `json:"specification"`
	Unit          string  `json:"unit"`
	MinQuantity   float64 `json:"min_quantity"`
	MaxQuantity   float64 `json:"max_quantity"`
	Location      string  `json:"location"`
	UnitPrice     float64 `json:"unit_price"`
	Remark        string  `json:"remark"`
}

// SpareTransactionReq 备件事务请求
type SpareTransactionReq struct {
	SpareID         uint    `json:"spare_id" binding:"required"`
	TransactionType string  `json:"transaction_type" binding:"required"` // IN/OUT/APPLY
	Quantity        float64 `json:"quantity" binding:"required"`
	OrderNo         string  `json:"order_no"`
	Remark          string  `json:"remark"`
}
