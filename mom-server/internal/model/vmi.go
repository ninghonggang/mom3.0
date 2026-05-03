package model

import (
	"time"
)

// VmiVendor VMI供应商
type VmiVendor struct {
	ID           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID     int64     `json:"tenant_id" gorm:"index;not null;default:1"`
	VendorID     int64     `json:"vendor_id" gorm:"index;not null"`    // 供应商ID
	VendorCode   string    `json:"vendor_code" gorm:"size:64"`          // 供应商编码
	VendorName   string    `json:"vendor_name" gorm:"size:200"`        // 供应商名称
	WarehouseID  int64     `json:"warehouse_id" gorm:"index"`           // 仓库ID
	WarehouseName string   `json:"warehouse_name" gorm:"size:100"`      // 仓库名称
	Contact      string    `json:"contact" gorm:"size:64"`             // 联系人
	Phone        string    `json:"phone" gorm:"size:20"`              // 电话
	MinStock     float64   `json:"min_stock"`                         // 最小库存
	MaxStock     float64   `json:"max_stock"`                         // 最大库存
	ReplenishCycle int     `json:"replenish_cycle"`                   // 补货周期(天)
	IsActive     int       `json:"is_active" gorm:"default:1"`        // 是否启用
	Status       int       `json:"status" gorm:"default:1"`            // 状态：1正常 0暂停
	Remarks      string    `json:"remarks" gorm:"type:text"`           // 备注
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (VmiVendor) TableName() string {
	return "vmi_vendor"
}

// VmiMaterial VMI物料
type VmiMaterial struct {
	ID            int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID      int64     `json:"tenant_id" gorm:"index;not null;default:1"`
	VendorID      int64     `json:"vendor_id" gorm:"index;not null"`   // 供应商ID
	MaterialID    int64     `json:"material_id" gorm:"index;not null"` // 物料ID
	MaterialCode  string    `json:"material_code" gorm:"size:64"`      // 物料编码
	MaterialName  string    `json:"material_name" gorm:"size:200"`     // 物料名称
	Unit          string    `json:"unit" gorm:"size:20"`               // 单位
	MinStock      float64   `json:"min_stock"`                         // 最小库存
	MaxStock      float64   `json:"max_stock"`                         // 最大库存
	CurrentStock  float64   `json:"current_stock"`                     // 当前库存
	AvailableStock float64  `json:"available_stock"`                   // 可用库存
	ConsumeQty    float64   `json:"consume_qty"`                       // 累计消耗
	LastConsumeDate *time.Time `json:"last_consume_date"`              // 最后消耗日期
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (VmiMaterial) TableName() string {
	return "vmi_material"
}

// VmiTransaction VMI事务记录
type VmiTransaction struct {
	ID            int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID      int64     `json:"tenant_id" gorm:"index;not null;default:1"`
	VendorID      int64     `json:"vendor_id" gorm:"index"`           // 供应商ID
	MaterialID    int64     `json:"material_id" gorm:"index"`         // 物料ID
	TransactionType int     `json:"transaction_type" gorm:"index"`    // 类型：1入库 2消耗 3盘点 4调整
	Quantity      float64   `json:"quantity"`                          // 数量
	BeforeQty     float64   `json:"before_qty"`                       // 变动前库存
	AfterQty      float64   `json:"after_qty"`                        // 变动后库存
	ReferenceNo   string    `json:"reference_no" gorm:"size:64"`      // 参考单号
	Remarks       string    `json:"remarks" gorm:"type:text"`          // 备注
	OperatorID    int64     `json:"operator_id"`                     // 操作人
	OperatorName  string    `json:"operator_name" gorm:"size:64"`     // 操作人名称
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (VmiTransaction) TableName() string {
	return "vmi_transaction"
}

// VmiVendorQuery VMI供应商查询
type VmiVendorQuery struct {
	VendorCode   string `form:"vendor_code"`
	VendorName   string `form:"vendor_name"`
	WarehouseID  int64  `form:"warehouse_id"`
	IsActive     int    `form:"is_active"`
	Status       int    `form:"status"`
	Page         int    `form:"page"`
	PageSize     int    `form:"page_size"`
}

// VmiMaterialQuery VMI物料查询
type VmiMaterialQuery struct {
	VendorID     int64  `form:"vendor_id"`
	MaterialCode string `form:"material_code"`
	MaterialName string `form:"material_name"`
	Page         int    `form:"page"`
	PageSize     int    `form:"page_size"`
}

// VmiTransactionQuery VMI事务查询
type VmiTransactionQuery struct {
	VendorID     int64  `form:"vendor_id"`
	MaterialID   int64  `form:"material_id"`
	TransactionType int `form:"transaction_type"`
	StartDate    string `form:"start_date"`
	EndDate      string `form:"end_date"`
	Page         int    `form:"page"`
	PageSize     int    `form:"page_size"`
}

// VmiVendorCreateReq 创建VMI供应商请求
type VmiVendorCreateReq struct {
	VendorID     int64   `json:"vendor_id" binding:"required"`
	VendorCode   string  `json:"vendor_code"`
	VendorName   string  `json:"vendor_name"`
	WarehouseID  int64   `json:"warehouse_id" binding:"required"`
	WarehouseName string `json:"warehouse_name"`
	Contact      string  `json:"contact"`
	Phone        string  `json:"phone"`
	MinStock     float64 `json:"min_stock"`
	MaxStock     float64 `json:"max_stock"`
	ReplenishCycle int   `json:"replenish_cycle"`
	Remarks      string  `json:"remarks"`
}

// VmiConsumeReq 消耗确认请求
type VmiConsumeReq struct {
	VendorID    int64    `json:"vendor_id" binding:"required"`
	MaterialID  int64    `json:"material_id" binding:"required"`
	Quantity    float64  `json:"quantity" binding:"required"`
	ReferenceNo string   `json:"reference_no"`
	Remarks     string   `json:"remarks"`
}

// VmiReplenishReq VMI补货请求
type VmiReplenishReq struct {
	VendorID    int64    `json:"vendor_id" binding:"required"`
	MaterialID int64     `json:"material_id" binding:"required"`
	Quantity   float64   `json:"quantity" binding:"required"`
	Remarks    string    `json:"remarks"`
}