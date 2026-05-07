package model

import (
	"time"
)

// EquipmentSpare 备件
type EquipmentSpare struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	TenantID      int64     `gorm:"index;not null;default:1" json:"tenant_id"`
	SpareCode     string    `gorm:"column:spare_part_code;size:50;uniqueIndex:idx_tenant_spare" json:"spare_code"`    // 备件编码
	SpareName     string    `gorm:"column:spare_part_name;size:100" json:"spare_name"`                              // 备件名称
	Spec          string    `gorm:"column:spec;size:100" json:"spec"`                                               // 规格型号
	Unit          string    `gorm:"column:unit;size:20" json:"unit"`                                                 // 单位
	Quantity      float64   `gorm:"column:quantity;type:decimal(18,2);default:0" json:"quantity"`                   // 当前库存
	MinQuantity   float64   `gorm:"column:min_quantity;type:decimal(18,2)" json:"min_quantity"`                      // 最小库存
	UnitPrice     float64   `gorm:"column:price;type:decimal(18,2)" json:"unit_price"`                              // 单价
	Supplier      string    `gorm:"column:supplier;size:100" json:"supplier"`                                        // 供应商
	Status        int       `gorm:"column:status;default:1" json:"status"`                                           // 状态
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName 表名
func (EquipmentSpare) TableName() string {
	return "equ_spare_part"
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
	SpareCode    string  `json:"spare_code" binding:"required"`
	SpareName    string  `json:"spare_name" binding:"required"`
	Spec         string  `json:"spec"`
	Unit         string  `json:"unit"`
	Quantity     float64 `json:"quantity"`
	MinQuantity  float64 `json:"min_quantity"`
	UnitPrice    float64 `json:"unit_price"`
}

// SpareUpdateReq 备件更新请求
type SpareUpdateReq struct {
	ID           uint    `json:"id" binding:"required"`
	SpareName    string  `json:"spare_name"`
	Spec         string  `json:"spec"`
	Unit         string  `json:"unit"`
	MinQuantity  float64 `json:"min_quantity"`
	UnitPrice    float64 `json:"unit_price"`
}

// SpareTransactionReq 备件事务请求
type SpareTransactionReq struct {
	SpareID         uint    `json:"spare_id" binding:"required"`
	TransactionType string  `json:"transaction_type" binding:"required"` // IN/OUT/APPLY
	Quantity        float64 `json:"quantity" binding:"required"`
	OrderNo         string  `json:"order_no"`
	Remark          string  `json:"remark"`
}
