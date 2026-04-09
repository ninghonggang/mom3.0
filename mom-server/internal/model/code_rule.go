package model

import "time"

// CodeRule 编码规则
type CodeRule struct {
	BaseModel
	TenantID    int64  `json:"tenant_id" gorm:"index;not null"`
	RuleCode   string `json:"rule_code" gorm:"size:50;not null"` // 规则编码
	RuleName   string `json:"rule_name" gorm:"size:100;not null"` // 规则名称
	EntityType string `json:"entity_type" gorm:"size:50;not null"` // 实体类型: PRODUCTION_ORDER/SALES_ORDER/PACKAGE etc.
	Prefix     string `json:"prefix" gorm:"size:20"` // 前缀
	DateFormat string `json:"date_format" gorm:"size:20"` // 日期格式: YYYYMMDD
	SeqLength  int    `json:"seq_length" gorm:"default:4"` // 序号长度
	SeqStart   int    `json:"seq_start" gorm:"default:1"` // 起始序号
	SeqCurrent int    `json:"seq_current" gorm:"default:0"` // 当前序号
	MidFix    string `json:"mid_fix" gorm:"size:20"` // 中间固定字符
	Suffix    string `json:"suffix" gorm:"size:20"` // 后缀
	ResetType string `json:"reset_type" gorm:"size:20"` // 重置类型: DAILY/MONTHLY/YEARLY/NONE
	LastGenDate string `json:"last_gen_date" gorm:"size:20"` // 最后生成日期
	Example   string `json:"example" gorm:"size:100"` // 示例
	Status    int    `json:"status" gorm:"default:1"` // 1启用/0禁用
	Remark    string `json:"remark" gorm:"size:500"` // 备注
}

func (CodeRule) TableName() string {
	return "mes_code_rule"
}

// CodeRuleRecord 编码记录
type CodeRuleRecord struct {
	BaseModel
	TenantID    int64  `json:"tenant_id" gorm:"index;not null"`
	RuleID     int64  `json:"rule_id" gorm:"index;not null"` // 规则ID
	RuleCode   string `json:"rule_code" gorm:"size:50"` // 规则编码
	EntityType string `json:"entity_type" gorm:"size:50"` // 实体类型
	GenDate    string `json:"gen_date" gorm:"size:20"` // 生成日期
	SeqValue   int    `json:"seq_value"` // 序号值
	GenCode    string `json:"gen_code" gorm:"size:100;uniqueIndex"` // 生成的编码
}

func (CodeRuleRecord) TableName() string {
	return "mes_code_rule_record"
}

// FlowCard 生产指示单
type FlowCard struct {
	BaseModel
	TenantID      int64      `json:"tenant_id" gorm:"index;not null"`
	CardNo       string     `json:"card_no" gorm:"size:50;not null"` // 指示单号
	OrderID      int64      `json:"order_id"` // 生产工单ID
	OrderNo      string     `json:"order_no" gorm:"size:50"` // 工单编号
	MaterialID   int64      `json:"material_id"` // 物料ID
	MaterialCode string     `json:"material_code" gorm:"size:50"` // 物料编码
	MaterialName string     `json:"material_name" gorm:"size:100"` // 物料名称
	WorkshopID   int64      `json:"workshop_id"` // 车间ID
	WorkshopName string     `json:"workshop_name" gorm:"size:100"` // 车间名称
	LineID       int64      `json:"line_id"` // 产线ID
	LineName     string     `json:"line_name" gorm:"size:100"` // 产线名称
	ProcessID    int64      `json:"process_id"` // 当前工序ID
	ProcessName  string     `json:"process_name" gorm:"size:100"` // 当前工序名称
	StationID    int64      `json:"station_id"` // 工位ID
	StationName  string     `json:"station_name" gorm:"size:100"` // 工位名称
	PlanQty      float64    `json:"plan_qty" gorm:"type:decimal(18,4)"` // 计划数量
	CompletedQty float64    `json:"completed_qty" gorm:"type:decimal(18,4)"` // 已完成数量
	Status       int        `json:"status" gorm:"default:1"` // 1待生产/2生产中/3已完成/4已取消
	Priority     int        `json:"priority" gorm:"default:1"` // 1普通/2紧急/3加急
	PlanStartTime *time.Time `json:"plan_start_time"` // 计划开始时间
	PlanEndTime   *time.Time `json:"plan_end_time"` // 计划结束时间
	ActualStartTime *time.Time `json:"actual_start_time"` // 实际开始时间
	ActualEndTime   *time.Time `json:"actual_end_time"` // 实际结束时间
	Remark       string     `json:"remark" gorm:"size:500"` // 备注
}

func (FlowCard) TableName() string {
	return "mes_flow_card"
}

// FlowCardDetail 指示单明细
type FlowCardDetail struct {
	BaseModel
	CardID     int64   `json:"card_id" gorm:"index;not null"` // 指示单ID
	StepNo     int     `json:"step_no"` // 步骤序号
	ProcessID  int64   `json:"process_id"` // 工序ID
	ProcessName string  `json:"process_name" gorm:"size:100"` // 工序名称
	StationID  int64   `json:"station_id"` // 工位ID
	StationName string `json:"station_name" gorm:"size:100"` // 工位名称
	WorkContent string  `json:"work_content" gorm:"size:500"` // 工作内容
	StdCycleTime int    `json:"std_cycle_time"` // 标准周期时间(秒)
	SeqQty     float64 `json:"seq_qty" gorm:"type:decimal(18,4)"` // 工序数量
}

func (FlowCardDetail) TableName() string {
	return "mes_flow_card_detail"
}
