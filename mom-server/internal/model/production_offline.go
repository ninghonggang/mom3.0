package model

import (
	"time"
)

// ProductionOffline 产品离线记录
type ProductionOffline struct {
	ID             uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID       uint64    `gorm:"index;not null;default:1" json:"tenant_id"`
	OfflineCode    string    `gorm:"size:50;uniqueIndex" json:"offline_code"` // 离线单号
	WorkOrderID    uint64    `gorm:"index" json:"work_order_id"`              // 工单ID
	WorkOrderCode  string    `gorm:"size:50" json:"work_order_code"`          // 工单编号
	ProductID      uint64    `gorm:"index" json:"product_id"`                 // 产品ID
	ProductCode    string    `gorm:"size:50" json:"product_code"`             // 产品编号
	ProductName    string    `gorm:"size:100" json:"product_name"`            // 产品名称
	OfflineType    string    `gorm:"size:20;index" json:"offline_type"`       // QUALITY/EQUIPMENT/MATERIAL/OTHER
	OfflineReason  string    `gorm:"size:200" json:"offline_reason"`          // 离线原因
	OfflineQty     float64   `json:"offline_qty"`                             // 离线数量
	ProcessRouteID uint64    `json:"process_route_id"`                        // 工艺路线ID
	CurrentOpID    uint64    `json:"current_op_id"`                           // 当前工序ID
	CurrentOpName  string    `gorm:"size:50" json:"current_op_name"`          // 当前工序名称
	HandleMethod   string    `gorm:"size:20;index" json:"handle_method"`      // REWORK/DOWNGRADE/SCRAP
	HandleQty      float64   `json:"handle_qty"`                              // 处理数量
	HandleResult   string    `gorm:"size:20" json:"handle_result"`            // PENDING/COMPLETED
	ReworkOrderID  uint64    `json:"rework_order_id"`                         // 返工工单ID
	ScrapQty       float64   `json:"scrap_qty"`                               // 报废数量
	DowngradeQty   float64   `json:"downgrade_qty"`                           // 降级数量
	Status         string    `gorm:"size:20;index" json:"status"`             // OPEN/HANDLED/CLOSED
	OperatorID     uint64    `json:"operator_id"`                             // 操作员ID
	OperatorName   string    `gorm:"size:50" json:"operator_name"`            // 操作员姓名
	Remark         string    `gorm:"size:500" json:"remark"`
	CreatedBy      string    `gorm:"size:50" json:"created_by"`
	UpdatedBy      string    `gorm:"size:50" json:"updated_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (ProductionOffline) TableName() string {
	return "production_offline"
}

// ProductionOfflineItem 离线产品明细
type ProductionOfflineItem struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID     uint64    `gorm:"index;not null;default:1" json:"tenant_id"`
	OfflineID    uint64    `gorm:"index;not null" json:"offline_id"` // 离线主表ID
	SerialNo     string    `gorm:"size:50" json:"serial_no"`         // 产品序列号
	BatchNo      string    `gorm:"size:50" json:"batch_no"`          // 批次号
	OfflineQty   float64   `json:"offline_qty"`                      // 离线数量
	HandleMethod string    `gorm:"size:20" json:"handle_method"`     // REWORK/DOWNGRADE/SCRAP
	HandleQty    float64   `json:"handle_qty"`                       // 处理数量
	HandleResult string    `gorm:"size:20" json:"handle_result"`     // PENDING/COMPLETED
	Remark       string    `gorm:"size:200" json:"remark"`
	CreatedAt    time.Time `json:"created_at"`
}

func (ProductionOfflineItem) TableName() string {
	return "production_offline_item"
}

// ProductionOfflineCreateRequest 创建请求
type ProductionOfflineCreateRequest struct {
	WorkOrderID    uint64                        `json:"work_order_id" binding:"required"`
	WorkOrderCode  string                        `json:"work_order_code" binding:"required"`
	ProductID      uint64                        `json:"product_id" binding:"required"`
	ProductCode    string                        `json:"product_code" binding:"required"`
	ProductName    string                        `json:"product_name" binding:"required"`
	OfflineType    string                        `json:"offline_type" binding:"required"` // QUALITY/EQUIPMENT/MATERIAL/OTHER
	OfflineReason  string                        `json:"offline_reason"`
	OfflineQty     float64                       `json:"offline_qty" binding:"required"`
	ProcessRouteID uint64                        `json:"process_route_id"`
	CurrentOpID    uint64                        `json:"current_op_id"`
	CurrentOpName  string                        `json:"current_op_name"`
	OperatorID     uint64                        `json:"operator_id"`
	OperatorName   string                        `json:"operator_name"`
	Remark         string                        `json:"remark"`
	Items          []ProductionOfflineItemCreate `json:"items"`
}

// ProductionOfflineItemCreate 离线产品明细创建
type ProductionOfflineItemCreate struct {
	SerialNo   string  `json:"serial_no"`
	BatchNo    string  `json:"batch_no"`
	OfflineQty float64 `json:"offline_qty"`
	Remark     string  `json:"remark"`
}

// ProductionOfflineUpdateRequest 更新请求
type ProductionOfflineUpdateRequest struct {
	OfflineType   string  `json:"offline_type"`
	OfflineReason string  `json:"offline_reason"`
	OfflineQty    float64 `json:"offline_qty"`
	CurrentOpID   uint64  `json:"current_op_id"`
	CurrentOpName string  `json:"current_op_name"`
	OperatorID    uint64  `json:"operator_id"`
	OperatorName  string  `json:"operator_name"`
	Remark        string  `json:"remark"`
}

// ProductionOfflineHandleRequest 处理请求
type ProductionOfflineHandleRequest struct {
	HandleMethod  string                        `json:"handle_method" binding:"required"` // REWORK/DOWNGRADE/SCRAP
	HandleQty     float64                       `json:"handle_qty" binding:"required"`
	ReworkOrderID uint64                        `json:"rework_order_id"`
	Items         []ProductionOfflineItemHandle `json:"items"`
}

// ProductionOfflineItemHandle 离线产品明细处理
type ProductionOfflineItemHandle struct {
	ID           uint64  `json:"id"`
	HandleMethod string  `json:"handle_method"`
	HandleQty    float64 `json:"handle_qty"`
	HandleResult string  `json:"handle_result"`
	Remark       string  `json:"remark"`
}

// ProductionOfflineQuery 查询参数
type ProductionOfflineQuery struct {
	WorkOrderCode string `form:"work_order_code"`
	ProductCode   string `form:"product_code"`
	ProductName   string `form:"product_name"`
	OfflineType   string `form:"offline_type"`
	HandleMethod  string `form:"handle_method"`
	Status        string `form:"status"`
	StartDate     string `form:"start_date"`
	EndDate       string `form:"end_date"`
	Page          int    `form:"page"`
	PageSize      int    `form:"page_size"`
}
