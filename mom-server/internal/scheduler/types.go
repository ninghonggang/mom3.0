package scheduler

import (
	"time"
)

// Extended SchedulingRule enum - new rules added to existing FIFO/EDD/SPT/LPT in scheduler.go
const (
	RuleJITFirst   SchedulingRule = "JIT_FIRST"  // JIT需求优先
	RuleCR         SchedulingRule = "CR"          // 紧迫系数 (交期-当前)/剩余工时
	RuleFamily     SchedulingRule = "FAMILY"    // 产品族聚类
	RuleBottleneck  SchedulingRule = "BOTTLENECK" // 瓶颈工序优先
)

// ScheduleConstraints 排程约束
type ScheduleConstraints struct {
	RespectJIT       bool    `json:"respect_jit"`        // 优先满足JIT时间窗
	MaxChangeoverPct float64 `json:"max_changeover_pct"` // 换型时间占比上限(默认30%)
	MinUtilization   float64 `json:"min_utilization"`   // 最低利用率
	AllowOvertime    bool    `json:"allow_overtime"`     // 允许加班
	FamilyGrouping   bool    `json:"family_grouping"`   // 产品族聚类
}

// ScheduleRequest 排程请求
type ScheduleRequest struct {
	PlanID          int64                `json:"plan_id"`
	WorkshopID      int64                `json:"workshop_id"`
	AlgorithmType   string               `json:"algorithm_type"` // FIFO/EDD/SPT/LPT/JIT_FIRST/CR/FAMILY/BOTTLENECK
	Direction       SchedulingDirection  `json:"direction"`      // FORWARD/BACKWARD
	StartDate       time.Time            `json:"start_date"`
	EndDate         time.Time            `json:"end_date"`
	Orders          []ScheduleOrder      `json:"orders"`
	WorkCenters     []WorkCenterInfo     `json:"work_centers"`
	Constraints     ScheduleConstraints  `json:"constraints"`
	ChangeoverMatrix map[string]float64  `json:"-"` // key: "from-to" -> minutes
	FamilyGroups    map[string][]string  `json:"-"` // family -> product_codes
}

// ScheduleOrder 工单信息
type ScheduleOrder struct {
	OrderID       int64     `json:"order_id"`
	OrderNo       string    `json:"order_no"`
	ProductID     int64     `json:"product_id"`
	ProductCode   string    `json:"product_code"`
	ProductName   string    `json:"product_name"`
	FamilyCode    string    `json:"family_code"`    // 产品族编码
	Quantity      float64   `json:"quantity"`
	StandardHours float64   `json:"standard_hours"` // 标准工时
	DueDate       time.Time `json:"due_date"`
	JITTime       *time.Time `json:"jit_time"`      // JIT需求时间
	Priority      int       `json:"priority"`
	ProcessCount  int       `json:"process_count"` // 工序数量
	Bottleneck    bool      `json:"bottleneck"`     // 是否瓶颈工序
}

// WorkCenterInfo 工作中心信息
type WorkCenterInfo struct {
	WorkCenterID   int64   `json:"work_center_id"`
	WorkCenterName string  `json:"work_center_name"`
	LineID         int64   `json:"line_id"`
	LineName       string  `json:"line_name"`
	Capacity       float64 `json:"capacity"`        // 产能(每小时)
	Efficiency     float64 `json:"efficiency"`      // 效率%
	MaxOvertime    float64 `json:"max_overtime"`   // 最大加班产能
}

// OptimizationSuggestion 优化建议
type OptimizationSuggestion struct {
	Type    string `json:"type"`    // CAPACITY/JIT/CHANGEVER/DELIVERY
	Level   string `json:"level"`   // INFO/WARNING/CRITICAL
	Message string `json:"message"`
	Impact  string `json:"impact"`  // 影响描述
}
