package model

import "time"

// ========== APS滚动排程与交付分析模块 ==========

// RollingConfig 滚动排程配置
type RollingConfig struct {
	BaseModel
	TenantID            int64     `json:"tenant_id" gorm:"index;not null"`
	ConfigCode         string    `json:"config_code" gorm:"size:50;uniqueIndex"`
	ConfigName         string    `json:"config_name" gorm:"size:100"`
	ConfigType         string    `json:"config_type" gorm:"size:20"`    // DAILY/HOURLY/EVENT
	TriggerType        string    `json:"trigger_type" gorm:"size:20"`   // CRON/MANUAL/EVENT
	TriggerCron        string    `json:"trigger_cron" gorm:"size:50"`   // Cron表达式
	HorizonDays        int       `json:"horizon_days" gorm:"default:7"`  // 排程视野(天数)
	LeadTimeDays       int       `json:"lead_time_days" gorm:"default:3"` // 提前下单期(天)
	SchedulingAlgorithm string    `json:"scheduling_algorithm" gorm:"size:20;default:FIFO"` // FIFO/EDD/SPT/LPT/JIT_FIRST/CR/FAMILY/BOTTLENECK
	Direction          string    `json:"direction" gorm:"size:20;default:FORWARD"` // FORWARD/BACKWARD
	OptimizeTarget     string    `json:"optimize_target" gorm:"size:20;default:DELIVERY"` // DELIVERY/UTILITY/EQUILIBRIUM
	RespectDueDate    int       `json:"respect_due_date" gorm:"default:1"`    // 是否严格遵守交付期
	MaxChangeoverPct  float64   `json:"max_changeover_pct" gorm:"type:decimal(5,2);default:30"` // 换型时间占比上限
	MinResourceUtilization float64 `json:"min_resource_utilization" gorm:"type:decimal(5,2);default:70"` // 最低资源利用率
	AllowOvertime     int       `json:"allow_overtime" gorm:"default:0"`     // 是否允许加班
	FamilyGrouping    int       `json:"family_grouping" gorm:"default:1"`     // 是否启用产品族聚类
	LockedTasksHandling string   `json:"locked_tasks_handling" gorm:"size:20;default:RETAIN"` // IGNORE/ADJUST/RETAIN
	AutoExecute       int       `json:"auto_execute" gorm:"default:0"`        // 是否自动执行
	IsEnabled         int       `json:"is_enabled" gorm:"default:1"`          // 是否启用
	LastExecuteTime   *time.Time `json:"last_execute_time"`
	ExecuteLog        string    `json:"execute_log" gorm:"type:text"`         // 执行日志
	Remark            string    `json:"remark" gorm:"size:500"`
}

func (RollingConfig) TableName() string {
	return "aps_rolling_config"
}

// RollingScheduleResult 滚动排程结果记录
type RollingScheduleResult struct {
	BaseModel
	ScheduleNo        string    `json:"schedule_no" gorm:"size:50;uniqueIndex"`
	ConfigID          *int64    `json:"config_id"`
	ScheduleDate      string    `json:"schedule_date" gorm:"size:10"`
	WorkshopID        *int64    `json:"workshop_id"`
	WorkshopName      string    `json:"workshop_name" gorm:"size:100"`
	Algorithm         string    `json:"algorithm" gorm:"size:20"`
	Direction         string    `json:"direction" gorm:"size:20"`
	HorizonStart      string    `json:"horizon_start" gorm:"size:10"`
	HorizonEnd        string    `json:"horizon_end" gorm:"size:10"`
	TotalOrders       int       `json:"total_orders" gorm:"default:0"`
	ScheduledOrders   int       `json:"scheduled_orders" gorm:"default:0"`
	FailedOrders      int       `json:"failed_orders" gorm:"default:0"`
	OnTimeOrders      int       `json:"on_time_orders" gorm:"default:0"`
	AvgUtilization    float64   `json:"avg_utilization" gorm:"type:decimal(5,2)"`
	MaxUtilization    float64   `json:"max_utilization" gorm:"type:decimal(5,2)"`
	ChangeoverTimeMinutes int   `json:"changeover_time_minutes" gorm:"default:0"`
	TotalProductionHours float64 `json:"total_production_hours" gorm:"type:decimal(10,2)"`
	Status            string    `json:"status" gorm:"size:20;default:COMPLETED"` // RUNNING/COMPLETED/FAILED
	ErrorMessage      string    `json:"error_message" gorm:"type:text"`
	ExecuteDurationMs int       `json:"execute_duration_ms"`
	TenantID          int64     `json:"tenant_id" gorm:"index;not null"`
}

func (RollingScheduleResult) TableName() string {
	return "aps_schedule_result"
}

// ScheduleWarning 排程冲突/警告
type ScheduleWarning struct {
	BaseModel
	ScheduleResultID  int64     `json:"schedule_result_id" gorm:"index"`
	WarningType       string    `json:"warning_type" gorm:"size:30"` // CAPACITY_OVERLOAD/DELAY_RISK/CHANGEVER_LONG/MATERIAL_SHORTAGE
	Severity          string    `json:"severity" gorm:"size:10"`    // HIGH/MEDIUM/LOW
	WorkCenterID      *int64    `json:"work_center_id"`
	WorkCenterName    string    `json:"work_center_name" gorm:"size:100"`
	OrderID           *int64    `json:"order_id"`
	OrderNo           string    `json:"order_no" gorm:"size:50"`
	Description       string    `json:"description" gorm:"size:500"`
	SuggestedAction   string    `json:"suggested_action" gorm:"size:200"`
	IsResolved        int       `json:"is_resolved" gorm:"default:0"`
	ResolvedBy        *int64    `json:"resolved_by"`
	ResolvedTime      *time.Time `json:"resolved_time"`
	TenantID          int64     `json:"tenant_id" gorm:"index;not null"`
}

func (ScheduleWarning) TableName() string {
	return "aps_schedule_warning"
}

// DeliveryAnalysis 交付性能分析
type DeliveryAnalysis struct {
	BaseModel
	AnalysisNo       string    `json:"analysis_no" gorm:"size:50;uniqueIndex"`
	AnalysisDate     string    `json:"analysis_date" gorm:"size:10"`
	AnalysisType     string    `json:"analysis_type" gorm:"size:20;default:DAILY"` // DAILY/WEEKLY/MONTHLY
	WorkshopID       *int64    `json:"workshop_id"`
	WorkshopName     string    `json:"workshop_name" gorm:"size:100"`
	TotalOrders     int       `json:"total_orders" gorm:"default:0"`
	OnTimeOrders    int       `json:"on_time_orders" gorm:"default:0"`      // 准时交付
	EarlyOrders     int       `json:"early_orders" gorm:"default:0"`         // 提前交付
	LateOrders      int       `json:"late_orders" gorm:"default:0"`          // 延期交付
	CriticalLateOrders int    `json:"critical_late_orders" gorm:"default:0"` // 严重延期(>3天)
	OnTimeRate      float64   `json:"on_time_rate" gorm:"type:decimal(5,2)"` // 准时交付率
	EarlyRate       float64   `json:"early_rate" gorm:"type:decimal(5,2)"`  // 提前交付率
	LateRate        float64   `json:"late_rate" gorm:"type:decimal(5,2)"`  // 延期率
	AvgDelayDays    float64   `json:"avg_delay_days" gorm:"type:decimal(8,2)"` // 平均延期天数
	MaxDelayDays    int       `json:"max_delay_days"`
	DelayDistribution string   `json:"delay_distribution" gorm:"type:jsonb"`   // 延期分布
	OtdTarget       float64   `json:"otd_target" gorm:"type:decimal(5,2)"` // 目标准时交付率
	OtdGap          float64   `json:"otd_gap" gorm:"type:decimal(5,2)"`    // 与目标差距
	ImprovementDirection string `json:"improvement_direction" gorm:"size:10"` // IMPROVE/DETERIORATE/STABLE
	DelayReasons    string     `json:"delay_reasons" gorm:"type:jsonb"`     // 延期原因统计
	BottleneckWorkcenters string `json:"bottleneck_workcenters" gorm:"type:jsonb"` // 瓶颈工作中心
	AnalysisSummary string     `json:"analysis_summary" gorm:"type:text"`
	TenantID        int64      `json:"tenant_id" gorm:"index;not null"`
	CreatedBy       string     `json:"created_by" gorm:"size:50"`
}

func (DeliveryAnalysis) TableName() string {
	return "aps_delivery_analysis"
}

// DeliveryWarning 交付预警
type DeliveryWarning struct {
	BaseModel
	WarningNo         string    `json:"warning_no" gorm:"size:50;uniqueIndex"`
	OrderID           int64     `json:"order_id" gorm:"index"`
	OrderNo           string    `json:"order_no" gorm:"size:50"`
	CustomerName      string    `json:"customer_name" gorm:"size:100"`
	DueDate           string    `json:"due_date" gorm:"size:10"`
	RemainingDays     int       `json:"remaining_days"`
	RiskLevel         string    `json:"risk_level" gorm:"size:10"`  // CRITICAL/HIGH/MEDIUM/LOW
	CurrentStatus     string    `json:"current_status" gorm:"size:30"` // PENDING/SCHEDULED/IN_PRODUCTION
	ScheduledCompletionDate *string `json:"scheduled_completion_date" gorm:"size:10"`
	DelayRiskDays    int        `json:"delay_risk_days"`           // 预计延误天数
	AffectedQuantity float64     `json:"affected_quantity" gorm:"type:decimal(18,3)"`
	CustomerPriority string      `json:"customer_priority" gorm:"size:10"` // VIP/HIGH/NORMAL
	PenaltyRisk      float64     `json:"penalty_risk" gorm:"type:decimal(18,2)"` // 违约风险金额
	SuggestedActions string      `json:"suggested_actions" gorm:"type:jsonb"`     // 建议措施
	EscalationRequired int       `json:"escalation_required" gorm:"default:0"`
	Status           string      `json:"status" gorm:"size:20;default:OPEN"` // OPEN/ACKNOWLEDGED/MITIGATED/CLOSED
	AcknowledgedBy   *int64      `json:"acknowledged_by"`
	AcknowledgedTime *time.Time  `json:"acknowledged_time"`
	TenantID         int64       `json:"tenant_id" gorm:"index;not null"`
}

func (DeliveryWarning) TableName() string {
	return "aps_delivery_warning"
}

// DeliveryMonthly 交付统计汇总(月度)
type DeliveryMonthly struct {
	BaseModel
	StatMonth       string    `json:"stat_month" gorm:"size:7;uniqueIndex:idx_month_workshop"`
	WorkshopID     *int64    `json:"workshop_id" gorm:"index;uniqueIndex:idx_month_workshop"`
	TotalOrders   int        `json:"total_orders" gorm:"default:0"`
	OnTimeOrders  int        `json:"on_time_orders" gorm:"default:0"`
	OnTimeRate   float64     `json:"on_time_rate" gorm:"type:decimal(5,2)"`
	AvgFulfillmentRate float64 `json:"avg_fulfillment_rate" gorm:"type:decimal(5,2)"`
	TopDeliveryReason string  `json:"top_delivery_reason" gorm:"size:100"`
	TenantID     int64        `json:"tenant_id" gorm:"index;not null"`
}

func (DeliveryMonthly) TableName() string {
	return "aps_delivery_monthly"
}

// MaterialShortage 缺料分析
type MaterialShortage struct {
	BaseModel
	AnalysisNo      string    `json:"analysis_no" gorm:"size:50;uniqueIndex"`
	AnalysisDate    string    `json:"analysis_date" gorm:"size:10"`
	PlanID          *int64    `json:"plan_id"`
	PlanNo          string    `json:"plan_no" gorm:"size:50"`
	MaterialID      int64     `json:"material_id" gorm:"index"`
	MaterialCode    string    `json:"material_code" gorm:"size:50"`
	MaterialName    string    `json:"material_name" gorm:"size:100"`
	MaterialType    string    `json:"material_type" gorm:"size:30"`  // 原材料/辅料/成品/半成品
	RequiredQty     float64   `json:"required_qty" gorm:"type:decimal(18,3)"`
	RequiredDate    string    `json:"required_date" gorm:"size:10"`
	AvailableQty    float64   `json:"available_qty" gorm:"type:decimal(18,3)"` // 可用量(含在途)
	OnHandQty      float64   `json:"on_hand_qty" gorm:"type:decimal(18,3)"`   // 现有量
	AllocatedQty    float64   `json:"allocated_qty" gorm:"type:decimal(18,3)"` // 已分配量
	InTransitQty    float64   `json:"in_transit_qty" gorm:"type:decimal(18,3)"` // 在途量
	ShortageQty     float64   `json:"shortage_qty" gorm:"type:decimal(18,3)"`  // 缺口数量
	ShortagePct     float64   `json:"shortage_pct" gorm:"type:decimal(5,2)"`   // 缺口比例
	ShortageLevel   string    `json:"shortage_level" gorm:"size:10"`  // CRITICAL/HIGH/MEDIUM/LOW
	SupplierID      *int64    `json:"supplier_id"`
	SupplierName    string    `json:"supplier_name" gorm:"size:100"`
	LeadTimeDays    int       `json:"lead_time_days"`
	ExpectedSupplyDate *string `json:"expected_supply_date" gorm:"size:10"`
	AlternativeSources string `json:"alternative_sources" gorm:"type:jsonb"` // 替代方案
	HasAlternative  int       `json:"has_alternative" gorm:"default:0"`
	SuggestedAlternativeID *int64 `json:"suggested_alternative_id"`
	AffectedOrders  string    `json:"affected_orders" gorm:"type:jsonb"`     // 受影响订单
	AffectedOrderCount int    `json:"affected_order_count" gorm:"default:0"`
	SuggestedAction string    `json:"suggested_action" gorm:"size:200"`
	IsUrgentPurchaseRequired int `json:"is_urgent_purchase_required" gorm:"default:0"`
	PurchaseLeadTimeDays int   `json:"purchase_lead_time_days"`
	Status          string    `json:"status" gorm:"size:20;default:ANALYZED"` // ANALYZED/ORDERED/PARTIAL/PROJECTED/SUPPLIED
	ResolvedBy      *int64    `json:"resolved_by"`
	ResolvedTime    *time.Time `json:"resolved_time"`
	ResolutionRemark string   `json:"resolution_remark" gorm:"size:500"`
	TenantID        int64     `json:"tenant_id" gorm:"index;not null"`
}

func (MaterialShortage) TableName() string {
	return "aps_material_shortage"
}

// ShortageWarningRule 缺料预警规则
type ShortageWarningRule struct {
	BaseModel
	RuleCode          string    `json:"rule_code" gorm:"size:50;uniqueIndex"`
	RuleName          string    `json:"rule_name" gorm:"size:100"`
	MaterialType      string    `json:"material_type" gorm:"size:30"` // 适用物料类型，为空则全部
	ShortageThresholdPct float64 `json:"shortage_threshold_pct" gorm:"type:decimal(5,2);default:20"` // 缺料预警阈值
	CheckHorizonDays  int       `json:"check_horizon_days" gorm:"default:7"` // 检查视野(天)
	NotifyChannels    string    `json:"notify_channels" gorm:"type:jsonb"`  // 通知配置
	AutoCreatePurchaseRequest int `json:"auto_create_purchase_request" gorm:"default:0"` // 是否自动生成采购申请
	IsEnabled         int       `json:"is_enabled" gorm:"default:1"`
	TenantID          int64     `json:"tenant_id" gorm:"index;not null"`
}

func (ShortageWarningRule) TableName() string {
	return "aps_shortage_warning_rule"
}

// ShortageStatistics 缺料统计汇总
type ShortageStatistics struct {
	BaseModel
	StatDate          string    `json:"stat_date" gorm:"size:10;uniqueIndex:idx_stat_workshop"`
	WorkshopID       *int64    `json:"workshop_id" gorm:"index;uniqueIndex:idx_stat_workshop"`
	TotalShortageCases int      `json:"total_shortage_cases" gorm:"default:0"`
	CriticalCases    int        `json:"critical_cases" gorm:"default:0"`
	HighCases        int        `json:"high_cases" gorm:"default:0"`
	MediumCases      int        `json:"medium_cases" gorm:"default:0"`
	LowCases         int        `json:"low_cases" gorm:"default:0"`
	TotalShortageAmount float64 `json:"total_shortage_amount" gorm:"type:decimal(18,2)"`
	AffectedOrdersCount int     `json:"affected_orders_count" gorm:"default:0"`
	TopShortageMaterials string `json:"top_shortage_materials" gorm:"type:jsonb"` // TOP缺料物料
	TenantID         int64     `json:"tenant_id" gorm:"index;not null"`
}

func (ShortageStatistics) TableName() string {
	return "aps_shortage_statistics"
}

// APSShift 班次
type APSShift struct {
	BaseModel
	ShiftCode   string    `json:"shift_code" gorm:"size:50;uniqueIndex"`
	ShiftName   string    `json:"shift_name" gorm:"size:100"`
	ShiftType   string    `json:"shift_type" gorm:"size:20"` // REGULAR/OVERTIME/HOLIDAY
	StartTime   string    `json:"start_time" gorm:"size:10"` // HH:MM
	EndTime     string    `json:"end_time" gorm:"size:10"`   // HH:MM
	BreakStart  *string   `json:"break_start" gorm:"size:10"`
	BreakEnd    *string   `json:"break_end" gorm:"size:10"`
	WorkHours   float64   `json:"work_hours" gorm:"type:decimal(5,2)"`
	IsNightShift int      `json:"is_night_shift" gorm:"default:0"`
	Color       string    `json:"color" gorm:"size:20"`  // 甘特图颜色
	SortOrder   int       `json:"sort_order" gorm:"default:0"`
	Remark      string    `json:"remark" gorm:"size:200"`
	TenantID    int64     `json:"tenant_id" gorm:"index;not null"`
}

func (APSShift) TableName() string {
	return "aps_shift"
}

// WorkshopCalendar 车间日日历
type WorkshopCalendar struct {
	BaseModel
	WorkshopID   int64     `json:"workshop_id" gorm:"index"`
	WorkDate    string     `json:"work_date" gorm:"size:10;uniqueIndex:idx_workshop_date"`
	ShiftID     *int64     `json:"shift_id"`
	ShiftName   string     `json:"shift_name" gorm:"size:100"`
	IsWorkingDay int      `json:"is_working_day" gorm:"default:1"`  // 1=工作日, 0=休息日
	IsHoliday   int        `json:"is_holiday" gorm:"default:0"`
	HolidayName string     `json:"holiday_name" gorm:"size:100"`
	OvertimeHours float64  `json:"overtime_hours" gorm:"type:decimal(5,2);default:0"`
	Remark      string     `json:"remark" gorm:"size:200"`
	TenantID    int64      `json:"tenant_id" gorm:"index;not null"`
}

func (WorkshopCalendar) TableName() string {
	return "aps_workshop_calendar"
}

// ChangeoverMatrix 换型矩阵
type APSChangeoverMatrix struct {
	BaseModel
	TenantID          int64    `json:"tenant_id" gorm:"index;not null"`
	FromProductID     int64    `json:"from_product_id"`
	FromProductCode   string   `json:"from_product_code" gorm:"size:50"`
	FromProductName   string   `json:"from_product_name" gorm:"size:100"`
	ToProductID       int64    `json:"to_product_id"`
	ToProductCode     string   `json:"to_product_code" gorm:"size:50"`
	ToProductName     string   `json:"to_product_name" gorm:"size:100"`
	ChangeoverTime    float64  `json:"changeover_time" gorm:"type:decimal(10,2);default:0"` // 换型时间(分钟)
	SetupTime         float64  `json:"setup_time" gorm:"type:decimal(10,2)"`
	CleanTime         float64  `json:"clean_time" gorm:"type:decimal(10,2)"`
	IsColorChange     int      `json:"is_color_change" gorm:"default:0"`
	IsMaterialChange  int      `json:"is_material_change" gorm:"default:0"`
	ChangeoverType    string   `json:"changeover_type" gorm:"size:20"` // CLEAN/RETOOL/ADJUST/OTHER
	Remark            string   `json:"remark" gorm:"size:200"`
}

func (APSChangeoverMatrix) TableName() string {
	return "aps_changeover_matrix_detail"
}

// ProductFamily 产品族定义
type ProductFamily struct {
	BaseModel
	FamilyCode         string   `json:"family_code" gorm:"size:50;uniqueIndex"`
	FamilyName         string   `json:"family_name" gorm:"size:100"`
	FamilyColor        string   `json:"family_color" gorm:"size:20"`  // 甘特图颜色
	ProductIDs         string   `json:"product_ids" gorm:"type:jsonb"` // 该族包含的产品ID列表
	AvgChangeoverTime  float64  `json:"avg_changeover_time" gorm:"type:decimal(8,2)"`
	Description        string   `json:"description" gorm:"size:500"`
	TenantID           int64    `json:"tenant_id" gorm:"index;not null"`
}

func (ProductFamily) TableName() string {
	return "aps_product_family"
}
