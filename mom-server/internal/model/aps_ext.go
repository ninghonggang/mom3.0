package model

import "time"

// CapacityAnalysis 产能分析
type CapacityAnalysis struct {
	BaseModel
	TenantID      int64     `json:"tenant_id" gorm:"index;not null"`
	WorkshopID   int64     `json:"workshop_id"` // 车间ID
	WorkshopName string    `json:"workshop_name" gorm:"size:100"` // 车间名称
	LineID       int64     `json:"line_id"` // 产线ID
	LineName     string    `json:"line_name" gorm:"size:100"` // 产线名称
	WorkDate     string    `json:"work_date" gorm:"size:10"` // 工作日期
	ShiftID      int64     `json:"shift_id"` // 班次ID
	ShiftName    string    `json:"shift_name" gorm:"size:50"` // 班次名称
	PlanCapacity float64   `json:"plan_capacity" gorm:"type:decimal(18,2)"` // 计划产能
	ActualCapacity float64 `json:"actual_capacity" gorm:"type:decimal(18,2)"` // 实际产能
	Utilization  float64   `json:"utilization" gorm:"type:decimal(5,2)"` // 利用率%
	OutputQty    float64   `json:"output_qty" gorm:"type:decimal(18,2)"` // 产出数量
	TargetQty    float64   `json:"target_qty" gorm:"type:decimal(18,2)"` // 目标数量
	RejectQty    float64   `json:"reject_qty" gorm:"type:decimal(18,2)"` // 不良数量
	RejectRate   float64   `json:"reject_rate" gorm:"type:decimal(5,2)"` // 不良率%
	Uptime       float64   `json:"uptime" gorm:"type:decimal(5,2)"` // 运行时间(小时)
	Downtime     float64   `json:"downtime" gorm:"type:decimal(5,2)"` // 停机时间(小时)
	Status       int       `json:"status" gorm:"default:1"` // 1正常/2异常
}

func (CapacityAnalysis) TableName() string {
	return "aps_capacity_analysis"
}

// DeliveryRate 交付率
type DeliveryRate struct {
	BaseModel
	TenantID        int64     `json:"tenant_id" gorm:"index;not null"`
	OrderNo         string    `json:"order_no" gorm:"size:50"` // 订单编号
	CustomerID      int64     `json:"customer_id"` // 客户ID
	CustomerName     string    `json:"customer_name" gorm:"size:100"` // 客户名称
	MaterialID      int64     `json:"material_id"` // 物料ID
	MaterialCode    string    `json:"material_code" gorm:"size:50"` // 物料编码
	MaterialName    string    `json:"material_name" gorm:"size:100"` // 物料名称
	PlanDeliveryDate string    `json:"plan_delivery_date" gorm:"size:10"` // 计划交付日期
	ActualDeliveryDate *time.Time `json:"actual_delivery_date"` // 实际交付日期
	PlanQty         float64   `json:"plan_qty" gorm:"type:decimal(18,4)"` // 计划数量
	DeliveryQty     float64   `json:"delivery_qty" gorm:"type:decimal(18,4)"` // 已交付数量
	OnTimeQty       float64   `json:"on_time_qty" gorm:"type:decimal(18,4)"` // 准时交付数量
	OnTimeRate      float64   `json:"on_time_rate" gorm:"type:decimal(5,2)"` // 准时交付率%
	FulfillmentRate float64   `json:"fulfillment_rate" gorm:"type:decimal(5,2)"` // 完成率%
	Status          int       `json:"status" gorm:"default:1"` // 1待交付/2部分交付/3已完成/4逾期
}

func (DeliveryRate) TableName() string {
	return "aps_delivery_rate"
}

// ChangeoverMatrix 换型矩阵
type ChangeoverMatrix struct {
	BaseModel
	TenantID       int64   `json:"tenant_id" gorm:"index;not null"`
	FromProductID  int64   `json:"from_product_id"` // 换出产品ID
	FromProductCode string  `json:"from_product_code" gorm:"size:50"` // 换出产品编码
	FromProductName string  `json:"from_product_name" gorm:"size:100"` // 换出产品名称
	ToProductID    int64   `json:"to_product_id"` // 换入产品ID
	ToProductCode  string  `json:"to_product_code" gorm:"size:50"` // 换入产品编码
	ToProductName  string  `json:"to_product_name" gorm:"size:100"` // 换入产品名称
	ChangeoverTime float64 `json:"changeover_time" gorm:"type:decimal(10,2)"` // 换型时间(分钟)
	SetupTime      float64 `json:"setup_time" gorm:"type:decimal(10,2)"` // 准备时间
	CleanTime      float64 `json:"clean_time" gorm:"type:decimal(10,2)"` // 清洁时间
	IsOptimized    int     `json:"is_optimized" gorm:"default:0"` // 是否已优化
	Priority       int     `json:"priority" gorm:"default:1"` // 优先级
	Status         int     `json:"status" gorm:"default:1"` // 1启用/0禁用
	Remark         string  `json:"remark" gorm:"size:500"` // 备注
}

func (ChangeoverMatrix) TableName() string {
	return "aps_changeover_matrix"
}

// RollingSchedule 滚动排程
type RollingSchedule struct {
	BaseModel
	TenantID     int64      `json:"tenant_id" gorm:"index;not null"`
	ScheduleNo   string     `json:"schedule_no" gorm:"size:50"` // 排程单号
	ScheduleType string     `json:"schedule_type" gorm:"size:20"` // DAILY/WEEKLY/MONTHLY
	StartDate    string     `json:"start_date" gorm:"size:10"` // 开始日期
	EndDate      string     `json:"end_date" gorm:"size:10"` // 结束日期
	WorkshopID   int64      `json:"workshop_id"` // 车间ID
	WorkshopName string     `json:"workshop_name" gorm:"size:100"` // 车间名称
	LineID       int64      `json:"line_id"` // 产线ID
	LineName     string     `json:"line_name" gorm:"size:100"` // 产线名称
	PlanQty     float64    `json:"plan_qty" gorm:"type:decimal(18,4)"` // 计划数量
	CompletedQty float64    `json:"completed_qty" gorm:"type:decimal(18,4)"` // 已完成数量
	StartTime    *time.Time `json:"start_time"` // 开始时间
	EndTime      *time.Time `json:"end_time"` // 结束时间
	HorizonDays  int       `json:"horizon_days"` // 展望期(天)
	Status       int        `json:"status" gorm:"default:1"` // 1待执行/2执行中/3已完成
	ExecuteTime  *time.Time `json:"execute_time"` // 执行时间
}

func (RollingSchedule) TableName() string {
	return "aps_rolling_schedule"
}

// JITDemand JIT需求
type JITDemand struct {
	BaseModel
	TenantID      int64      `json:"tenant_id" gorm:"index;not null"`
	DemandNo     string     `json:"demand_no" gorm:"size:50"` // 需求单号
	DemandType   string     `json:"demand_type" gorm:"size:20"` // JIT/JIS
	MaterialID   int64      `json:"material_id"` // 物料ID
	MaterialCode string     `json:"material_code" gorm:"size:50"` // 物料编码
	MaterialName string     `json:"material_name" gorm:"size:100"` // 物料名称
	CustomerID   int64      `json:"customer_id"` // 客户ID
	CustomerName string     `json:"customer_name" gorm:"size:100"` // 客户名称
	DemandQty   float64     `json:"demand_qty" gorm:"type:decimal(18,4)"` // 需求数量
	DemandTime  *time.Time  `json:"demand_time"` // 需求时间
	Priority    int         `json:"priority" gorm:"default:1"` // 优先级
	Frequency   string      `json:"frequency" gorm:"size:20"` // 供货频率
	LeadTime    int         `json:"lead_time"` // 前置时间(分钟)
	KanbanQty   float64     `json:"kanban_qty" gorm:"type:decimal(18,4)"` // 看板数量
	Status      int          `json:"status" gorm:"default:1"` // 1待配送/2配送中/3已配送
}

func (JITDemand) TableName() string {
	return "aps_jit_demand"
}
