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

func (SalesOrder) TableName() string {
	return "pro_sales_order"
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

func (SalesOrderItem) TableName() string {
	return "pro_sales_order_item"
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
	Status          int        `json:"status" gorm:"default:1"` // legacy bigint,与 status_v2 双轨 - 1草稿/2已下达/3生产中/4已完成/5已关闭/6已取消
	StatusV2        string     `json:"status_v2" gorm:"size:30;index"` // 双轨:status_v2 varchar(30) - MOM 3.0 V2.1, 与 mdm_status_dict 对齐
	Remark          *string    `json:"remark" gorm:"size:500"`
}

func (ProductionOrder) TableName() string {
	return "production_orders"
}

// StatusCode 返回有效状态码(优先 status_v2, fallback status int → 字典码)
func (o *ProductionOrder) StatusCode() string {
	if o.StatusV2 != "" {
		return o.StatusV2
	}
	return map[int]string{1: "DRAFT", 2: "RELEASED", 3: "IN_PROGRESS", 4: "COMPLETED", 5: "CLOSED", 6: "CANCELLED"}[o.Status]
}

// SetStatus 同时写 status(int) + status_v2(varchar) - 双轨制
func (o *ProductionOrder) SetStatus(code string) {
	o.StatusV2 = code
	switch code {
	case "DRAFT":
		o.Status = 1
	case "RELEASED":
		o.Status = 2
	case "IN_PROGRESS":
		o.Status = 3
	case "COMPLETED":
		o.Status = 4
	case "CLOSED":
		o.Status = 5
	case "CANCELLED":
		o.Status = 6
	}
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
	Status        int        `json:"status" gorm:"default:1"` // 1已提交/2已确认/3已审核
	StatusV2      string     `json:"status_v2" gorm:"size:30;index"` // 双轨:MOM 3.0 V2.1
	Remark        *string    `json:"remark" gorm:"size:500"`
}

func (ProductionReport) TableName() string {
	return "pro_production_report"
}

// StatusCode 返回有效状态码(优先 status_v2, fallback status int → 字典码)
func (o *ProductionReport) StatusCode() string {
	if o.StatusV2 != "" {
		return o.StatusV2
	}
	return map[int]string{1: "SUBMITTED", 2: "CONFIRMED", 3: "AUDITED", 4: "REJECTED"}[o.Status]
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
	StatusV2      string     `json:"status_v2" gorm:"size:30;index"` // 双轨:MOM 3.0 V2.1
}

func (Dispatch) TableName() string {
	return "pro_dispatch"
}

// StatusCode 返回有效状态码(优先 status_v2, fallback status int → 字典码)
func (o *Dispatch) StatusCode() string {
	if o.StatusV2 != "" {
		return o.StatusV2
	}
	return map[int]string{1: "PENDING", 2: "IN_PROGRESS", 3: "COMPLETED"}[o.Status]
}

// ProductionOrderChangeLog 生产工单变更日志
type ProductionOrderChangeLog struct {
	BaseModel
	TenantID      int64  `json:"tenant_id" gorm:"index;not null"`
	OrderID       int64  `json:"order_id" gorm:"index;not null"`
	OrderNo       string `json:"order_no" gorm:"size:50"`
	ChangeType    string `json:"change_type" gorm:"size:32;not null"` // QUANTITY_CHANGE, DATE_CHANGE, PRIORITY_CHANGE, LINE_CHANGE, STATUS_CHANGE
	OldValue      string `json:"old_value" gorm:"type:text"`
	NewValue      string `json:"new_value" gorm:"type:text"`
	ChangeReason  string `json:"change_reason" gorm:"size:256"`
	ChangedBy     string `json:"changed_by" gorm:"size:50"`
	ChangedAt     string `json:"changed_at" gorm:"size:50"`
}

func (ProductionOrderChangeLog) TableName() string {
	return "mes_production_order_change_log"
}
