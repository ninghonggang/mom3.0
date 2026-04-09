package model

import (
	"encoding/json"
	"time"
)

// ========== Andon升级机制 ==========

// AndonCall 安东呼叫记录
type AndonCall struct {
	BaseModel
	TenantID           int64           `json:"tenant_id" gorm:"index;not null"`
	CallNo             string          `json:"call_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_andon"`
	WorkshopID         int64          `json:"workshop_id" gorm:"index"`
	WorkshopName       string          `json:"workshop_name" gorm:"size:100"`
	ProductionLineID   *int64         `json:"production_line_id" gorm:"index"`
	ProductionLineName *string        `json:"production_line_name" gorm:"size:100"`
	WorkstationID      int64          `json:"workstation_id" gorm:"index"`
	WorkstationName    *string        `json:"workstation_name" gorm:"size:100"`
	AndonType          string         `json:"andon_type" gorm:"size:30;index"` // EQUIPMENT/MATERIAL/QUALITY/TECHNICAL/TOOLING/SAFETY/OTHER
	AndonTypeName      string         `json:"andon_type_name" gorm:"size:50"`
	CallLevel          int            `json:"call_level" gorm:"default:1"` // 1/2/3/4
	Priority           int            `json:"priority" gorm:"default:5"`   // 1-10
	Description        string         `json:"description" gorm:"size:500"`
	Photos             json.RawMessage `json:"photos" gorm:"type:jsonb"`   // [{"url": "", "uploaded_at": ""}]
	AudioURL           string         `json:"audio_url" gorm:"size:500"`
	CallBy             string         `json:"call_by" gorm:"size:50"`
	CallTime           time.Time      `json:"call_time" gorm:"index"`
	ResponseBy         string         `json:"response_by" gorm:"size:50"`
	ResponseTime       *time.Time     `json:"response_time"`
	HandleBy           string         `json:"handle_by" gorm:"size:50"`
	HandleTime         *time.Time     `json:"handle_time"`
	HandleResult       string         `json:"handle_result" gorm:"size:20"` // RESOLVED/CARRY_OVER/ESCALATED
	HandleRemark       string         `json:"handle_remark" gorm:"size:500"`
	RelatedOrderID     *int64         `json:"related_order_id" gorm:"index"`
	RelatedNCRID       *int64         `json:"related_ncr_id" gorm:"index"`
	RelatedRepairID    *int64         `json:"related_repair_id" gorm:"index"`
	Status             string         `json:"status" gorm:"size:20;default:CALLING;index"` // CALLING/RESPONDED/HANDLING/RESOLVED/CLOSED
	IsEscalated        int            `json:"is_escalated" gorm:"default:0"`                // 0未升级/1已升级
	EscalatedAt        *time.Time     `json:"escalated_at"`
	EscalationCount    int            `json:"escalation_count" gorm:"default:0"`
	ResponseDuration   *int           `json:"response_duration"` // 响应时长(秒)
	HandleDuration     *int           `json:"handle_duration"`   // 处理时长(秒)
}

func (AndonCall) TableName() string {
	return "andon_call"
}

// AndonEscalationRule 升级规则
type AndonEscalationRule struct {
	BaseModel
	RuleCode       string           `json:"rule_code" gorm:"size:50;not null;uniqueIndex"`
	RuleName       string           `json:"rule_name" gorm:"size:100;not null"`
	AndonType      string           `json:"andon_type" gorm:"size:30"` // 适用类型，为空=全部适用
	WorkshopID     *int64           `json:"workshop_id" gorm:"index"`  // 适用车间，为空=全部
	PriorityRange  string           `json:"priority_range" gorm:"size:20"` // 适用优先级范围，如 "1-5"
	IsDefault      int              `json:"is_default" gorm:"default:0"`  // 1是默认规则

	// L1 第一级
	Level1Timeout    int              `json:"level1_timeout" gorm:"not null"` // 超时时间(分钟)
	Level1NotifyType string           `json:"level1_notify_type" gorm:"size:20"` // WORKSTATION/SCREEN/PUSH/ALL
	Level1NotifyJSON json.RawMessage  `json:"level1_notify_json" gorm:"type:jsonb"` // [{"user_id": "", "user_name": "", "role": ""}]

	// L2 第二级
	Level2Timeout    int              `json:"level2_timeout"`
	Level2NotifyType string           `json:"level2_notify_type" gorm:"size:20"`
	Level2NotifyJSON json.RawMessage  `json:"level2_notify_json" gorm:"type:jsonb"`

	// L3 第三级
	Level3Timeout    int              `json:"level3_timeout"`
	Level3NotifyType string          `json:"level3_notify_type" gorm:"size:20"`
	Level3NotifyJSON json.RawMessage `json:"level3_notify_json" gorm:"type:jsonb"`

	// L4 第四级(最高)
	Level4Timeout    int              `json:"level4_timeout"`
	Level4NotifyType string          `json:"level4_notify_type" gorm:"size:20"`
	Level4NotifyJSON json.RawMessage `json:"level4_notify_json" gorm:"type:jsonb"`

	// 升级规则
	EscalationMode     string `json:"escalation_mode" gorm:"size:20;default:TIMEOUT"` // TIMEOUT/MANUAL/AUTO
	MaxEscalationLevel int    `json:"max_escalation_level" gorm:"default:4"`

	// 广播配置
	AudioEnabled          int    `json:"audio_enabled" gorm:"default:0"`
	AudioMessageTemplate  string `json:"audio_message_template" gorm:"size:200"`
	AudioRepeatTimes      int    `json:"audio_repeat_times" gorm:"default:3"`

	// 状态
	IsEnabled  int    `json:"is_enabled" gorm:"default:1"`
	SortOrder  int    `json:"sort_order" gorm:"default:0"`
	Remark     string `json:"remark" gorm:"size:500"`
	TenantID   int64  `json:"tenant_id" gorm:"index;not null"`
	CreatedBy  string `json:"created_by" gorm:"size:50"`
}

func (AndonEscalationRule) TableName() string {
	return "andon_escalation_rule"
}

// AndonEscalationLog 升级历史
type AndonEscalationLog struct {
	BaseModel
	CallID         int64            `json:"call_id" gorm:"index;not null"`
	FromLevel      int              `json:"from_level" gorm:"not null"`
	ToLevel        int              `json:"to_level" gorm:"not null"`
	EscalationType string           `json:"escalation_type" gorm:"size:20"` // TIMEOUT/MANUAL/AUTO
	TriggerUser    string           `json:"trigger_user" gorm:"size:50"`   // 触发人(手动升级)
	TriggerReason  string           `json:"trigger_reason" gorm:"size:200"`
	NotifyResult   json.RawMessage  `json:"notify_result" gorm:"type:jsonb"` // [{"channel": "", "status": "", "message": ""}]
	CreatedAt      time.Time        `json:"created_at" gorm:"default:now()"`
}

func (AndonEscalationLog) TableName() string {
	return "andon_escalation_log"
}

// AndonNotificationLog 消息推送记录
type AndonNotificationLog struct {
	BaseModel
	CallID       int64     `json:"call_id" gorm:"index;not null"`
	Channel      string    `json:"channel" gorm:"size:20;not null"` // FEISHU/WECHAT/SMS/AUDIO/WEBSOCKET
	ReceiverType string    `json:"receiver_type" gorm:"size:20"`   // USER/ROLE/DEPARTMENT
	ReceiverID   string    `json:"receiver_id" gorm:"size:50"`
	ReceiverName string    `json:"receiver_name" gorm:"size:100"`
	Title        string    `json:"title" gorm:"size:200"`
	Content      string    `json:"content" gorm:"type:text"`
	Priority     int       `json:"priority"`
	SendTime     time.Time `json:"send_time"`
	SendResult   string    `json:"send_result" gorm:"size:20"` // SUCCESS/FAILED/PENDING
	ErrorMsg     string    `json:"error_msg" gorm:"size:500"`
	TenantID     int64     `json:"tenant_id" gorm:"index;not null"`
}

func (AndonNotificationLog) TableName() string {
	return "andon_notification_log"
}

// NotifyInfo 通知对象
type NotifyInfo struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
}

// GetLevel1Notifiers 获取L1通知对象
func (r *AndonEscalationRule) GetLevel1Notifiers() []NotifyInfo {
	return getNotifiersFromJSON(r.Level1NotifyJSON)
}

// GetLevel2Notifiers 获取L2通知对象
func (r *AndonEscalationRule) GetLevel2Notifiers() []NotifyInfo {
	return getNotifiersFromJSON(r.Level2NotifyJSON)
}

// GetLevel3Notifiers 获取L3通知对象
func (r *AndonEscalationRule) GetLevel3Notifiers() []NotifyInfo {
	return getNotifiersFromJSON(r.Level3NotifyJSON)
}

// GetLevel4Notifiers 获取L4通知对象
func (r *AndonEscalationRule) GetLevel4Notifiers() []NotifyInfo {
	return getNotifiersFromJSON(r.Level4NotifyJSON)
}

func getNotifiersFromJSON(data json.RawMessage) []NotifyInfo {
	if data == nil || string(data) == "" || string(data) == "null" {
		return []NotifyInfo{}
	}
	var notifiers []NotifyInfo
	if err := json.Unmarshal(data, &notifiers); err != nil {
		return []NotifyInfo{}
	}
	return notifiers
}
