package model

import (
	"time"
)

// BpmTaskMessageRule BPM任务消息规则
type BpmTaskMessageRule struct {
	ID            uint      `json:"id" gorm:"primarykey"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	TenantID      int64     `json:"tenant_id" gorm:"index;not null"`
	RuleCode      string    `json:"rule_code" gorm:"size:50;not null"`      // 规则编码
	RuleName      string    `json:"rule_name" gorm:"size:100;not null"`     // 规则名称
	ProcessDefKey string    `json:"process_def_key" gorm:"size:50"`        // 流程定义Key
	TaskDefKey    string    `json:"task_def_key" gorm:"size:50"`           // 任务定义Key
	MessageType   string    `json:"message_type" gorm:"size:20"`           // EMAIL/SMS/WEBSOCKET
	TemplateCode  string    `json:"template_code" gorm:"size:50"`          // 消息模板编码
	IsEnabled     bool      `json:"is_enabled" gorm:"default:true"`
	Priority      int       `json:"priority" gorm:"default:5"`
	Remark        string    `json:"remark" gorm:"type:text"`
}

func (BpmTaskMessageRule) TableName() string {
	return "bpm_task_message_rule"
}

// BpmTaskMessageRuleCreateReqVO 创建请求
type BpmTaskMessageRuleCreateReqVO struct {
	RuleCode      string `json:"rule_code"`
	RuleName      string `json:"rule_name"`
	ProcessDefKey string `json:"process_def_key"`
	TaskDefKey    string `json:"task_def_key"`
	MessageType   string `json:"message_type"`
	TemplateCode  string `json:"template_code"`
	IsEnabled     bool   `json:"is_enabled"`
	Priority      int    `json:"priority"`
	Remark        string `json:"remark"`
}

// BpmTaskMessageRuleUpdateReqVO 更新请求
type BpmTaskMessageRuleUpdateReqVO struct {
	Id            int64  `json:"id"`
	RuleName      string `json:"rule_name"`
	ProcessDefKey string `json:"process_def_key"`
	TaskDefKey    string `json:"task_def_key"`
	MessageType   string `json:"message_type"`
	TemplateCode  string `json:"template_code"`
	IsEnabled     bool   `json:"is_enabled"`
	Priority      int    `json:"priority"`
	Remark        string `json:"remark"`
}
