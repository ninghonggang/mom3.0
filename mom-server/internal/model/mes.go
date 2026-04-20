package model

import (
	"time"
)

// ========== MES生产执行模块 - 月计划/日计划 ==========

// OrderMonth 月度计划主表
type OrderMonth struct {
	BaseModel
	TenantID        int64      `json:"tenant_id" gorm:"index;not null"`
	MonthPlanNo     string     `json:"month_plan_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_month_plan"` // 计划单号，例: MP-2026-04
	PlanMonth       string     `json:"plan_month" gorm:"size:7;not null"`                                        // 计划月份，例: 2026-04
	Title           string     `json:"title" gorm:"size:200"`                                                    // 计划标题
	SourceType      string     `json:"source_type" gorm:"size:20"`                                               // ERP/APS/FIS/MANUAL
	SourceNo        *string    `json:"source_no" gorm:"size:50"`                                                 // 来源单号

	WorkshopID      int64      `json:"workshop_id"`
	WorkshopName    *string    `json:"workshop_name" gorm:"size:100"`

	TotalProductCount int       `json:"total_product_count" gorm:"default:0"`   // 产品种类数
	TotalPlanQty      float64   `json:"total_plan_qty" gorm:"type:decimal(18,3);default:0"`   // 计划总数量
	TotalCompletedQty float64   `json:"total_completed_qty" gorm:"type:decimal(18,3);default:0"` // 已完成数量

	// 审核状态: DRAFT/SUBMITTED/APPROVED/RELEASED/CLOSED/CANCELLED
	ApprovalStatus   string     `json:"approval_status" gorm:"size:20;default:'DRAFT'"`
	SubmittedBy      *int64     `json:"submitted_by"`
	SubmittedAt      *time.Time `json:"submitted_at"`
	ApprovedBy       *int64     `json:"approved_by"`
	ApprovedAt       *time.Time `json:"approved_at"`
	ReleasedBy       *int64     `json:"released_by"`
	ReleasedAt       *time.Time `json:"released_at"`

	Remark          *string    `json:"remark" gorm:"type:text"`
	CreatedBy       *string    `json:"created_by" gorm:"size:50"`

	Items           []OrderMonthItem `json:"items" gorm:"foreignKey:MonthPlanID"`
}

func (OrderMonth) TableName() string {
	return "mes_order_month"
}

// OrderMonthItem 月计划明细行
type OrderMonthItem struct {
	BaseModel
	MonthPlanID     int64      `json:"month_plan_id" gorm:"index;not null"`
	LineNo         int        `json:"line_no" gorm:"not null"`

	ProductID      int64      `json:"product_id" gorm:"not null"`
	ProductCode    *string    `json:"product_code" gorm:"size:50"`
	ProductName    *string    `json:"product_name" gorm:"size:100"`
	Specification  *string    `json:"specification" gorm:"size:200"`
	Unit           *string    `json:"unit" gorm:"size:20"`

	PlanQty        float64    `json:"plan_qty" gorm:"type:decimal(18,3);not null"`     // 计划数量
	CompletedQty   float64    `json:"completed_qty" gorm:"type:decimal(18,3);default:0"` // 已完成数量
	ReleasedQty    float64    `json:"released_qty" gorm:"type:decimal(18,3);default:0"`  // 已下达数量

	DeliveryDate   *time.Time `json:"delivery_date"` // 交付要求
	Priority       int        `json:"priority" gorm:"default:5"` // 优先级 1-5

	Remark         *string    `json:"remark" gorm:"type:text"`
	TenantID       int64      `json:"tenant_id" gorm:"index;not null"`
}

func (OrderMonthItem) TableName() string {
	return "mes_order_month_item"
}

// OrderMonthAudit 月计划审核记录
type OrderMonthAudit struct {
	BaseModel
	MonthPlanID     int64      `json:"month_plan_id" gorm:"index;not null"`
	ApprovalStatus  string     `json:"approval_status" gorm:"size:20;not null"` // SUBMIT/APPROVE/REJECT/RELEASE/CLOSE/CANCEL
	ApproverID      *int64     `json:"approver_id"`
	ApproverName    *string    `json:"approver_name" gorm:"size:50"`
	ApprovalTime    *time.Time `json:"approval_time"`
	Comment         *string    `json:"comment" gorm:"type:text"`
	TenantID        int64      `json:"tenant_id" gorm:"index;not null"`
}

func (OrderMonthAudit) TableName() string {
	return "mes_order_month_audit"
}

// OrderDay 日计划主表
type OrderDay struct {
	BaseModel
	TenantID        int64      `json:"tenant_id" gorm:"index;not null"`
	DayPlanNo       string     `json:"day_plan_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_day_plan"` // 日计划单号，例: DP-20260415
	PlanDate        time.Time  `json:"plan_date" gorm:"type:date;not null"`                                    // 计划日期

	MonthPlanID     *int64     `json:"month_plan_id"`                                  // 关联月计划
	MonthPlanNo     *string    `json:"month_plan_no" gorm:"size:50"`                  // 月计划单号

	WorkshopID      int64      `json:"workshop_id"`
	WorkshopName    *string    `json:"workshop_name" gorm:"size:100"`
	ProductionLineID *int64    `json:"production_line_id"`
	LineName        *string    `json:"line_name" gorm:"size:100"`

	TotalProductCount int       `json:"total_product_count" gorm:"default:0"`
	TotalPlanQty      float64   `json:"total_plan_qty" gorm:"type:decimal(18,3);default:0"`
	TotalCompletedQty float64   `json:"total_completed_qty" gorm:"type:decimal(18,3);default:0"`

	// 齐套检查: PENDING/CHECKING/READY/NOT_READY
	KitStatus       string     `json:"kit_status" gorm:"size:20;default:'PENDING'"`
	KitCheckTime    *time.Time `json:"kit_check_time"`
	KitCheckBy      *int64     `json:"kit_check_by"`

	// 状态: DRAFT/PUBLISHED/IN_PRODUCTION/COMPLETED/TERMINATED
	Status          string     `json:"status" gorm:"size:20;default:'DRAFT'"`
	ShiftType       string     `json:"shift_type" gorm:"size:20"` // DAY/AFTERNOON/NIGHT/ALL

	PublishedAt     *time.Time `json:"published_at"`
	PublishedBy     *int64     `json:"published_by"`

	Remark          *string    `json:"remark" gorm:"type:text"`
	CreatedBy       *string    `json:"created_by" gorm:"size:50"`

	Items           []OrderDayItem `json:"items" gorm:"foreignKey:DayPlanID"`
}

func (OrderDay) TableName() string {
	return "mes_order_day"
}

// OrderDayItem 日计划明细行
type OrderDayItem struct {
	BaseModel
	DayPlanID       int64      `json:"day_plan_id" gorm:"index;not null"`
	LineNo         int        `json:"line_no" gorm:"not null"`

	ProductID      int64      `json:"product_id" gorm:"not null"`
	ProductCode    *string    `json:"product_code" gorm:"size:50"`
	ProductName    *string    `json:"product_name" gorm:"size:100"`
	Specification  *string    `json:"specification" gorm:"size:200"`
	Unit           *string    `json:"unit" gorm:"size:20"`

	PlanQty        float64    `json:"plan_qty" gorm:"type:decimal(18,3);not null"`
	CompletedQty   float64    `json:"completed_qty" gorm:"type:decimal(18,3);default:0"`

	BOMID          *int64     `json:"bom_id"`
	BOMVersion     *string    `json:"bom_version" gorm:"size:20"`
	ProcessRouteID *int64     `json:"process_route_id"`                     // 工艺路线ID
	RouteVersion   *string    `json:"route_version" gorm:"size:20"`

	ProductionMode string     `json:"production_mode" gorm:"size:20;default:'BATCH'"` // BATCH/SINGLE

	// 齐套明细
	KitStatus      string     `json:"kit_status" gorm:"size:20;default:'PENDING'"` // PENDING/READY/NOT_READY
	KitCheckRemark *string    `json:"kit_check_remark" gorm:"type:text"`

	// 来源月计划明细
	MonthPlanItemID *int64    `json:"month_plan_item_id"`

	// PENDING/IN_PROGRESS/COMPLETED
	ItemStatus     string     `json:"item_status" gorm:"size:20;default:'PENDING'"`

	Priority       int        `json:"priority" gorm:"default:5"`
	Remark         *string    `json:"remark" gorm:"type:text"`
	TenantID       int64      `json:"tenant_id" gorm:"index;not null"`
}

func (OrderDayItem) TableName() string {
	return "mes_order_day_item"
}

// OrderDayWorkOrderMap 工单生成记录（日计划发布时自动生成）
type OrderDayWorkOrderMap struct {
	BaseModel
	DayPlanID       int64      `json:"day_plan_id" gorm:"index;not null"`
	DayPlanItemID   int64      `json:"day_plan_item_id" gorm:"index;not null"`
	WorkOrderID     int64      `json:"work_order_id" gorm:"index;not null"`
	WorkOrderNo     string     `json:"work_order_no" gorm:"size:50;not null"`
	TenantID        int64      `json:"tenant_id" gorm:"index;not null"`
}

func (OrderDayWorkOrderMap) TableName() string {
	return "bpm_order_day_workorder_map"
}

// ========== 请求/响应结构 ==========

// OrderMonthCreate 月计划创建请求
type OrderMonthCreate struct {
	Title      string                 `json:"title"`
	PlanMonth  string                 `json:"plan_month"` // YYYYMM
	SourceType string                 `json:"source_type"`
	SourceNo   string                 `json:"source_no"`
	WorkshopID int64                  `json:"workshop_id"`
	Remark     string                 `json:"remark"`
	Items      []OrderMonthItemCreate `json:"items"`
}

// OrderMonthItemCreate 月计划明细创建请求
type OrderMonthItemCreate struct {
	ProductID    int64     `json:"product_id"`
	ProductCode  string    `json:"product_code"`
	ProductName  string    `json:"product_name"`
	Specification string   `json:"specification"`
	Unit         string    `json:"unit"`
	PlanQty      float64   `json:"plan_qty"`
	DeliveryDate string    `json:"delivery_date"` // YYYY-MM-DD
	Priority     int       `json:"priority"`
	Remark       string    `json:"remark"`
}

// OrderMonthUpdate 月计划更新请求
type OrderMonthUpdate struct {
	Title   string                 `json:"title"`
	Remark  string                 `json:"remark"`
	Items   []OrderMonthItemCreate `json:"items"`
}

// OrderDayCreate 日计划创建请求
type OrderDayCreate struct {
	PlanDate        string                `json:"plan_date"` // YYYY-MM-DD
	MonthPlanID     *int64                `json:"month_plan_id"`
	WorkshopID      int64                 `json:"workshop_id"`
	ProductionLineID *int64               `json:"production_line_id"`
	ShiftType       string                `json:"shift_type"`
	Remark          string                `json:"remark"`
	Items           []OrderDayItemCreate  `json:"items"`
}

// OrderDayItemCreate 日计划明细创建请求
type OrderDayItemCreate struct {
	ProductID      int64   `json:"product_id"`
	ProductCode    string  `json:"product_code"`
	ProductName    string  `json:"product_name"`
	Specification  string  `json:"specification"`
	Unit           string  `json:"unit"`
	PlanQty        float64 `json:"plan_qty"`
	BOMID          *int64  `json:"bom_id"`
	ProcessRouteID *int64  `json:"process_route_id"`
	ProductionMode string  `json:"production_mode"`
	MonthPlanItemID *int64 `json:"month_plan_item_id"`
	Priority       int     `json:"priority"`
	Remark         string  `json:"remark"`
}

// OrderDayUpdate 日计划更新请求
type OrderDayUpdate struct {
	ProductionLineID *int64 `json:"production_line_id"`
	ShiftType        string `json:"shift_type"`
	Remark           string `json:"remark"`
	Items            []OrderDayItemCreate `json:"items"`
}
