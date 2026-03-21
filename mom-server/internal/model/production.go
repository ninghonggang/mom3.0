package model

import (
	"time"
)

// ========== 生产执行模块 ==========

// SalesOrder 销售订单
type SalesOrder struct {
	BaseModel
	TenantID      int64      `json:"tenant_id" gorm:"index;not null"`
	OrderNo       string     `json:"order_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_order"`
	CustomerID    int64      `json:"customer_id"`
	CustomerName  *string    `json:"customer_name" gorm:"size:100"`
	OrderDate     *time.Time `json:"order_date"`
	DeliveryDate  *time.Time `json:"delivery_date"`
	OrderType     string     `json:"order_type" gorm:"size:20"` // 标准/定制
	Priority      int        `json:"priority" gorm:"default:1"`    // 1普通/2紧急/3加急
	Status        int        `json:"status" gorm:"default:1"`      // 1待确认/2已确认/3生产中/4已完成/5已关闭
	Remark        *string    `json:"remark" gorm:"size:500"`
}

// SalesOrderItem 销售订单明细
type SalesOrderItem struct {
	BaseModel
	OrderID      int64   `json:"order_id" gorm:"index;not null"`
	MaterialID   int64   `json:"material_id"`
	MaterialCode string   `json:"material_code" gorm:"size:50"`
	MaterialName string   `json:"material_name" gorm:"size:100"`
	Quantity     float64 `json:"quantity" gorm:"type:decimal(18,4)"`
	Unit         string   `json:"unit" gorm:"size:20"`
	Price        float64 `json:"price" gorm:"type:decimal(18,2)"` // 单价
	Amount       float64 `json:"amount" gorm:"type:decimal(18,2)"` // 金额
	ShippedQty   float64 `json:"shipped_qty" gorm:"type:decimal(18,4);default:0"`
}

// ProductionOrder 生产工单
type ProductionOrder struct {
	BaseModel
	TenantID       int64      `json:"tenant_id" gorm:"index;not null"`
	OrderNo         string     `json:"order_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_porder"`
	SalesOrderNo    *string    `json:"sales_order_no" gorm:"size:50"` // 销售订单号
	MaterialID      int64      `json:"material_id"`
	MaterialCode    string     `json:"material_code" gorm:"size:50"`
	MaterialName    string     `json:"material_name" gorm:"size:100"`
	MaterialSpec    *string    `json:"material_spec" gorm:"size:100"`
	Unit            string     `json:"unit" gorm:"size:20"`
	Quantity        float64    `json:"quantity" gorm:"type:decimal(18,4)"` // 计划数量
	CompletedQty    float64    `json:"completed_qty" gorm:"type:decimal(18,4);default:0"` // 已完成数量
	RejectedQty     float64    `json:"rejected_qty" gorm:"type:decimal(18,4);default:0"` // 不良品数量
	WorkshopID     int64      `json:"workshop_id"` // 车间ID
	WorkshopName    *string    `json:"workshop_name" gorm:"size:100"`
	LineID          int64      `json:"line_id"` // 生产线ID
	LineName        *string    `json:"line_name" gorm:"size:100"`
	RouteID         int64      `json:"route_id"` // 工艺路线ID
	BOMID           int64      `json:"bom_id"` // BOM ID
	PlanStartDate   *time.Time `json:"plan_start_date"` // 计划开始
	PlanEndDate     *time.Time `json:"plan_end_date"` // 计划结束
	ActualStartDate *time.Time `json:"actual_start_date"` // 实际开始
	ActualEndDate   *time.Time `json:"actual_end_date"` // 实际结束
	Priority        int        `json:"priority" gorm:"default:1"`
	Status          int        `json:"status" gorm:"default:1"` // 1待生产/2生产中/3已完成/4已取消
	Remark          *string    `json:"remark" gorm:"size:500"`
}

// ProductionReport 生产报工
type ProductionReport struct {
	BaseModel
	TenantID      int64      `json:"tenant_id" gorm:"index;not null"`
	OrderID       int64      `json:"order_id" gorm:"index;not null"`
	OrderNo       string     `json:"order_no" gorm:"size:50"`
	ProcessID     int64      `json:"process_id"` // 工序ID
	ProcessName   *string    `json:"process_name" gorm:"size:100"`
	StationID     int64      `json:"station_id"` // 工位ID
	StationName   *string    `json:"station_name" gorm:"size:100"`
	ReportUserID  int64      `json:"report_user_id"`
	ReportUserName *string   `json:"report_user_name" gorm:"size:50"`
	ReportDate    *time.Time `json:"report_date"`
	Quantity      float64    `json:"quantity" gorm:"type:decimal(18,4)"` // 报工数量
	QualifiedQty  float64    `json:"qualified_qty" gorm:"type:decimal(18,4)"` // 合格数量
	RejectedQty   float64    `json:"rejected_qty" gorm:"type:decimal(18,4);default:0"` // 不良数量
	WorkTime      int        `json:"work_time"` // 工时(分钟)
	Remark        *string    `json:"remark" gorm:"size:500"`
}

// Dispatch 派工单
type Dispatch struct {
	BaseModel
	TenantID      int64      `json:"tenant_id" gorm:"index;not null"`
	OrderID       int64      `json:"order_id" gorm:"index;not null"`
	OrderNo       string     `json:"order_no" gorm:"size:50"`
	ProcessID     int64      `json:"process_id"`
	ProcessName   *string    `json:"process_name" gorm:"size:100"`
	StationID     int64      `json:"station_id"`
	StationName   *string    `json:"station_name" gorm:"size:100"`
	AssignUserID  int64      `json:"assign_user_id"` // 分配给
	AssignUserName *string   `json:"assign_user_name" gorm:"size:50"`
	Quantity      float64    `json:"quantity" gorm:"type:decimal(18,4)"` // 派工数量
	Status        int        `json:"status" gorm:"default:1"` // 1待开始/2进行中/3已完成
}
