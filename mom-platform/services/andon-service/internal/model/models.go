package model

import (
	"time"

	"gorm.io/gorm"
)

// AndonCall 安灯呼叫记录
type AndonCall struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	TenantID        string         `gorm:"column:tenant_id;size:64;index" json:"tenant_id"`
	AndonNo         string         `gorm:"column:andon_no;size:128;uniqueIndex" json:"andon_no"`
	WorkstationID   string         `gorm:"column:workstation_id;size:64;index" json:"workstation_id"`
	ReporterID      string         `gorm:"column:reporter_id;size:64" json:"reporter_id"`
	AndonType       string         `gorm:"column:andon_type;size:32" json:"andon_type"` // MATERIAL/EQUIPMENT/QUALITY/SAFETY
	Description     string         `gorm:"column:description;type:text" json:"description"`
	Status          string         `gorm:"column:status;size:32;default:TRIGGERED" json:"status"` // TRIGGERED/ACKNOWLEDGED/IN_PROGRESS/RESOLVED/CLOSED/CANCELLED/ESCALATED
	TriggeredAt     time.Time      `gorm:"column:triggered_at" json:"triggered_at"`
	AcknowledgedAt  *time.Time     `gorm:"column:acknowledged_at" json:"acknowledged_at"`
	ResolvedAt      *time.Time     `gorm:"column:resolved_at" json:"resolved_at"`
	ResponseSeconds *int64         `gorm:"column:response_seconds" json:"response_seconds"`
	CreatedAt       time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	Actions []AndonAction `gorm:"foreignKey:AndonID" json:"actions,omitempty"`
}

func (AndonCall) TableName() string { return "andon_calls" }

// AndonAction 安灯操作记录
type AndonAction struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	AndonID      uint      `gorm:"column:andon_id;index" json:"andon_id"`
	ActionType   string    `gorm:"column:action_type;size:32" json:"action_type"`
	ActionDesc   string    `gorm:"column:action_desc;size:512" json:"action_desc"`
	OperatorID   string    `gorm:"column:operator_id;size:64" json:"operator_id"`
	ActionTime   time.Time `gorm:"column:action_time" json:"action_time"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
}

func (AndonAction) TableName() string { return "andon_actions" }

// AlertConfig 告警配置
type AlertConfig struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	ConfigCode       string         `gorm:"column:config_code;size:128;uniqueIndex" json:"config_code"`
	ConfigName       string         `gorm:"column:config_name;size:256" json:"config_name"`
	TriggerType      string         `gorm:"column:trigger_type;size:32" json:"trigger_type"` // THRESHOLD/EVENT/SCHEDULE
	Severity         string         `gorm:"column:severity;size:8" json:"severity"`          // P0/P1/P2/P3
	TriggerCondition string         `gorm:"column:trigger_condition;type:text" json:"trigger_condition"`
	NotifyChannels   string         `gorm:"column:notify_channels;size:256" json:"notify_channels"` // comma-separated
	Status           string         `gorm:"column:status;size:32;default:ENABLED" json:"status"`
	CreatedAt        time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}

func (AlertConfig) TableName() string { return "alert_configs" }

// Alert 告警记录
type Alert struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	ConfigID       uint           `gorm:"column:config_id;index" json:"config_id"`
	TargetID       string         `gorm:"column:target_id;size:64;index" json:"target_id"`
	TargetType     string         `gorm:"column:target_type;size:32" json:"target_type"`
	Status         string         `gorm:"column:status;size:32;default:ACTIVE" json:"status"` // ACTIVE/ACKNOWLEDGED/RESOLVED/ESCALATED/SUPPRESSED/CLOSED
	TriggeredAt    time.Time      `gorm:"column:triggered_at" json:"triggered_at"`
	AcknowledgedAt *time.Time     `gorm:"column:acknowledged_at" json:"acknowledged_at"`
	ResolvedAt     *time.Time     `gorm:"column:resolved_at" json:"resolved_at"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	Config      *AlertConfig      `gorm:"foreignKey:ConfigID" json:"config,omitempty"`
	Escalations []AlertEscalation `gorm:"foreignKey:AlertID" json:"escalations,omitempty"`
}

func (Alert) TableName() string { return "alerts" }

// AlertEscalation 告警升级记录
type AlertEscalation struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	AlertID          uint      `gorm:"column:alert_id;index" json:"alert_id"`
	Level            int       `gorm:"column:level" json:"level"` // 1/2/3
	EscalateToUserID string    `gorm:"column:escalate_to_user_id;size:64" json:"escalate_to_user_id"`
	TimeoutSeconds   int       `gorm:"column:timeout_seconds" json:"timeout_seconds"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
}

func (AlertEscalation) TableName() string { return "alert_escalations" }
