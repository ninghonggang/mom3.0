package model

import (
	"encoding/json"
	"time"
)

// AlertBaseModel 告警公共字段
type AlertBaseModel struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// AlertEscalationRule 告警升级规则
type AlertEscalationRule struct {
	AlertBaseModel
	RuleCode        string          `json:"rule_code" gorm:"size:50;uniqueIndex;not null"` // 规则编码
	RuleName        string          `json:"rule_name" gorm:"size:200;not null"`           // 规则名称
	AlertType       *string         `json:"alert_type" gorm:"size:30"`                    // 告警类型（为空表示通用）
	SeverityLevel   *string         `json:"severity_level" gorm:"size:10"`               // 严重程度
	EscalationLevels json.RawMessage `json:"escalation_levels" gorm:"type:jsonb"`           // 升级层级配置
	IsEnabled       int             `json:"is_enabled" gorm:"default:1"`                   // 是否启用
	TenantID        int64           `json:"tenant_id" gorm:"index;not null"`
	UpdatedAt       *time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}

func (AlertEscalationRule) TableName() string {
	return "alert_escalation_rule"
}

// AlertRule 告警规则配置
type AlertRule struct {
	AlertBaseModel
	RuleCode            string          `json:"rule_code" gorm:"size:50;uniqueIndex;not null"`  // 规则编码
	RuleName            string          `json:"rule_name" gorm:"size:200;not null"`              // 规则名称
	AlertType           string          `json:"alert_type" gorm:"size:30;not null"`             // 告警类型
	BizModule           *string         `json:"biz_module" gorm:"size:30"`                      // 业务模块
	ConditionExpression string          `json:"condition_expression" gorm:"size:500;not null"`   // 条件表达式
	ConditionParams     json.RawMessage `json:"condition_params" gorm:"type:jsonb"`              // 条件参数
	SeverityLevel       string          `json:"severity_level" gorm:"size:10;default:MEDIUM"`    // 严重程度
	NotificationChannels json.RawMessage `json:"notification_channels" gorm:"type:jsonb"`        // 通知渠道列表
	NotifyTemplates     json.RawMessage `json:"notify_templates" gorm:"type:jsonb"`              // 通知模板配置
	EscalationRuleID    *int64          `json:"escalation_rule_id"`                              // 升级规则ID
	IsEnabled           int             `json:"is_enabled" gorm:"default:1"`                     // 是否启用
	CheckInterval       int             `json:"check_interval" gorm:"default:60"`               // 检查间隔（秒）
	MaxTriggerCount     int             `json:"max_trigger_count" gorm:"default:0"`              // 最大触发次数
	TriggerCount        int             `json:"trigger_count" gorm:"default:0"`                 // 当前触发次数
	LastTriggerTime     *time.Time      `json:"last_trigger_time"`                              // 最近触发时间
	TenantID            int64           `json:"tenant_id" gorm:"index;not null"`
	CreatedBy           *string         `json:"created_by" gorm:"size:50"`                      // 创建人
	UpdatedBy           *string         `json:"updated_by" gorm:"size:50"`                      // 更新人
	UpdatedAt           *time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}

func (AlertRule) TableName() string {
	return "alert_rule"
}

// AlertRecord 告警记录
type AlertRecord struct {
	AlertBaseModel
	AlertNo             string          `json:"alert_no" gorm:"size:50;uniqueIndex;not null"`  // 告警编号
	RuleID              *int64          `json:"rule_id"`                                        // 关联规则ID
	RuleCode            *string         `json:"rule_code" gorm:"size:50"`                      // 规则编码
	RuleName            *string         `json:"rule_name" gorm:"size:200"`                      // 规则名称
	AlertType           string          `json:"alert_type" gorm:"size:30;not null"`             // 告警类型
	SeverityLevel       string          `json:"severity_level" gorm:"size:10;not null"`         // 严重程度
	Title               string          `json:"title" gorm:"size:200;not null"`                 // 告警标题
	Content             string          `json:"content" gorm:"type:text;not null"`              // 告警内容
	TriggerTime         time.Time       `json:"trigger_time" gorm:"not null"`                    // 触发时间
	SourceModule        *string         `json:"source_module" gorm:"size:30"`                    // 来源模块
	SourceID            *int64          `json:"source_id"`                                      // 来源ID
	SourceNo            *string         `json:"source_no" gorm:"size:50"`                        // 来源单号
	SourceData          json.RawMessage `json:"source_data" gorm:"type:jsonb"`                   // 来源详细数据
	Status              string          `json:"status" gorm:"size:20;default:TRIGGERED"`        // 状态
	UrgencyLevel        string          `json:"urgency_level" gorm:"size:10;default:NORMAL"`   // 紧急程度
	AcknowledgedBy      *int64          `json:"acknowledged_by"`                               // 确认人ID
	AcknowledgedByName  *string         `json:"acknowledged_by_name" gorm:"size:50"`            // 确认人姓名
	AcknowledgedTime    *time.Time      `json:"acknowledged_time"`                              // 确认时间
	AcknowledgedRemark  *string         `json:"acknowledged_remark" gorm:"size:200"`            // 确认备注
	ResolvedBy          *int64          `json:"resolved_by"`                                    // 解决人ID
	ResolvedByName      *string         `json:"resolved_by_name" gorm:"size:50"`                // 解决人姓名
	ResolvedTime        *time.Time      `json:"resolved_time"`                                  // 解决时间
	ResolutionRemark    *string         `json:"resolution_remark" gorm:"size:500"`              // 解决备注
	ClosedTime          *time.Time      `json:"closed_time"`                                    // 关闭时间
	EscalationCount     int             `json:"escalation_count" gorm:"default:0"`              // 升级次数
	LastEscalationTime  *time.Time      `json:"last_escalation_time"`                          // 最近升级时间
	TenantID            int64           `json:"tenant_id" gorm:"index;not null"`
	UpdatedAt           *time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}

func (AlertRecord) TableName() string {
	return "alert_record"
}

// AlertNotificationLog 告警通知日志
type AlertNotificationLog struct {
	AlertBaseModel
	AlertID            int64     `json:"alert_id" gorm:"index;not null"`                           // 关联告警ID
	AlertNo            string    `json:"alert_no" gorm:"size:50"`                                   // 告警编号
	Channel            string    `json:"channel" gorm:"size:20;not null"`                          // 通知渠道
	ReceiverType       *string   `json:"receiver_type" gorm:"size:20"`                             // 接收者类型
	ReceiverID         *int64    `json:"receiver_id"`                                             // 接收者ID
	ReceiverName       *string   `json:"receiver_name" gorm:"size:50"`                             // 接收者姓名
	ReceiverValue      *string   `json:"receiver_value" gorm:"size:200"`                            // 接收者值
	NotificationStatus *string   `json:"notification_status" gorm:"size:20"`                       // 通知状态
	SentTime           *time.Time `json:"sent_time"`                                             // 发送时间
	ReadTime           *time.Time `json:"read_time"`                                             // 阅读时间
	ErrorCode          *string   `json:"error_code" gorm:"size:50"`                               // 错误码
	ErrorMsg           *string   `json:"error_msg" gorm:"size:500"`                               // 错误信息
	RetryCount         int       `json:"retry_count" gorm:"default:0"`                            // 重试次数
	TenantID           int64     `json:"tenant_id" gorm:"index;not null"`
}

func (AlertNotificationLog) TableName() string {
	return "alert_notification_log"
}

// NotificationChannel 通知渠道配置
type NotificationChannel struct {
	AlertBaseModel
	ChannelCode    string         `json:"channel_code" gorm:"size:50;uniqueIndex;not null"` // 渠道编码
	ChannelName    string         `json:"channel_name" gorm:"size:100;not null"`           // 渠道名称
	ChannelType    string         `json:"channel_type" gorm:"size:20;not null"`             // 渠道类型: IN_SITE/FEISHU/WECOM/EMAIL
	Config         json.RawMessage `json:"config" gorm:"type:jsonb"`                       // 渠道配置JSON
	IsEnabled      int            `json:"is_enabled" gorm:"default:1"`                    // 是否启用
	Priority       int            `json:"priority" gorm:"default:0"`                       // 优先级
	TenantID       int64          `json:"tenant_id" gorm:"index;not null"`
	CreatedBy      *string        `json:"created_by" gorm:"size:50"`
	UpdatedBy      *string        `json:"updated_by" gorm:"size:50"`
	UpdatedAt      *time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
}

func (NotificationChannel) TableName() string {
	return "notification_channel"
}

// ChannelConfig 渠道配置结构
type ChannelConfig struct {
	WebhookURL  string           `json:"webhook_url"`            // Webhook地址（飞书/企微）
	SMTPHost   string           `json:"smtp_host"`             // SMTP服务器
	SMTPPort   int              `json:"smtp_port"`             // SMTP端口
	SMTPUser   string           `json:"smtp_user"`             // SMTP用户名
	SMTPPassword string         `json:"smtp_password"`         // SMTP密码
	SMTPFrom   string           `json:"smtp_from"`             // 发件人
	SMTPTLS    bool             `json:"smtp_tls"`              // 是否使用TLS
	Receivers  []ReceiverConfig `json:"receivers"`            // 默认接收人列表
}

// ReceiverConfig 接收者配置
type ReceiverConfig struct {
	ReceiverType  string `json:"receiver_type"`  // USER/ROLE/DEPT
	ReceiverID    int64  `json:"receiver_id"`
	ReceiverName  string `json:"receiver_name"`
	ReceiverValue string `json:"receiver_value"` // 邮箱/手机/用户ID
}

// SendNotificationRequest 发送通知请求
type SendNotificationRequest struct {
	AlertID       *int64  `json:"alert_id"`
	AlertNo       *string `json:"alert_no"`
	ChannelType   string  `json:"channel_type" binding:"required"` // IN_SITE/FEISHU/WECOM/EMAIL
	Title         string  `json:"title" binding:"required"`
	Content       string  `json:"content" binding:"required"`
	ReceiverType  *string `json:"receiver_type"`
	ReceiverID    *int64  `json:"receiver_id"`
	ReceiverName  *string `json:"receiver_name"`
	ReceiverValue *string `json:"receiver_value"`
}
