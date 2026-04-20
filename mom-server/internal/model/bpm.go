package model

import (
	"encoding/json"
	"time"
)

// BPMBaseModel BPM公共字段
type BPMBaseModel struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// ProcessModel 流程模型定义
type ProcessModel struct {
	BPMBaseModel
	ModelCode        string          `json:"model_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_model_version"` // 流程标识
	ModelName        string          `json:"model_name" gorm:"size:200;not null"`                                    // 流程名称
	ModelType        string          `json:"model_type" gorm:"size:20;not null"`                                     // 流程类型
	Version          string          `json:"version" gorm:"size:20;default:1.0"`                                     // 版本号
	Category         *string         `json:"category" gorm:"size:50"`                                                // 流程分类
	Description      *string         `json:"description" gorm:"type:text"`                                            // 描述
	FormType         string          `json:"form_type" gorm:"size:20;not null"`                                      // 表单类型
	FormDefinitionID *int64          `json:"form_definition_id"`                                                    // 关联表单定义ID
	FormURL          *string         `json:"form_url" gorm:"size:500"`                                              // 外部表单URL
	IsPublished      int             `json:"is_published" gorm:"default:0"`                                          // 是否已发布
	IsActive         int             `json:"is_active" gorm:"default:1"`                                            // 是否启用
	PublishedAt      *time.Time      `json:"published_at"`                                                          // 发布时间
	PublishedBy     *int64          `json:"published_by"`                                                          // 发布人
	Config           json.RawMessage `json:"config" gorm:"type:jsonb"`                                               // 流程配置
	TenantID         int64           `json:"tenant_id" gorm:"index;not null"`
	CreatedBy        *string         `json:"created_by" gorm:"size:50"`
	UpdatedAt        *time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	Nodes            []NodeDefinition `json:"nodes" gorm:"foreignKey:ModelID"`
	Flows            []SequenceFlow  `json:"flows" gorm:"foreignKey:ModelID"`
}

func (ProcessModel) TableName() string {
	return "bpm_process_model"
}

// NodeDefinition 流程节点定义
type NodeDefinition struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	ModelID     int64          `json:"model_id" gorm:"not null;index"`
	NodeID      string         `json:"node_id" gorm:"size:50;not null;uniqueIndex:idx_tenant_model_node"`
	NodeName    string         `json:"node_name" gorm:"size:200;not null"`
	NodeType    string         `json:"node_type" gorm:"size:30;not null"` // START/END/APPROVAL/COPY/CONDITION/ACTION
	PositionX   int            `json:"position_x" gorm:"default:0"`      // 图形化X坐标
	PositionY   int            `json:"position_y" gorm:"default:0"`      // 图形化Y坐标
	Width       int            `json:"width" gorm:"default:120"`
	Height      int            `json:"height" gorm:"default:80"`
	NodeConfig  json.RawMessage `json:"node_config" gorm:"type:jsonb;not null"` // 节点配置
	SortOrder   int            `json:"sort_order" gorm:"default:0"`
	TenantID    int64          `json:"tenant_id" gorm:"index;not null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
}

func (NodeDefinition) TableName() string {
	return "bpm_node_definition"
}

// SequenceFlow 流程连线定义
type SequenceFlow struct {
	ID                 uint           `json:"id" gorm:"primarykey"`
	ModelID            int64          `json:"model_id" gorm:"not null;index"`
	FlowID             string         `json:"flow_id" gorm:"size:50;not null;uniqueIndex:idx_tenant_model_flow"`
	SourceNodeID       string         `json:"source_node_id" gorm:"size:50;not null"` // 起始节点
	TargetNodeID       string         `json:"target_node_id" gorm:"size:50;not null"` // 目标节点
	FlowName           *string        `json:"flow_name" gorm:"size:200"`
	ConditionType      string         `json:"condition_type" gorm:"size:20;default:NONE"` // NONE/CONDITION/DEFAULT
	ConditionExpression *string        `json:"condition_expression" gorm:"type:text"`     // 条件表达式
	IsDefault          int            `json:"is_default" gorm:"default:0"`               // 是否默认分支
	FlowConfig         json.RawMessage `json:"flow_config" gorm:"type:jsonb"`
	SortOrder          int            `json:"sort_order" gorm:"default:0"`
	TenantID           int64          `json:"tenant_id" gorm:"index;not null"`
	CreatedAt          time.Time      `json:"created_at" gorm:"autoCreateTime"`
}

func (SequenceFlow) TableName() string {
	return "bpm_sequence_flow"
}

// ModelVersion 流程模型版本历史
type ModelVersion struct {
	BPMBaseModel
	ModelID           int64     `json:"model_id" gorm:"not null;index"`
	Version           string    `json:"version" gorm:"size:20;not null"`
	ChangeDescription *string   `json:"change_description" gorm:"type:text"`
	PublishedBy       *int64    `json:"published_by"`
	PublishedAt       *time.Time `json:"published_at"`
	IsCurrent         int       `json:"is_current" gorm:"default:0"`
	TenantID          int64     `json:"tenant_id" gorm:"index;not null"`
}

func (ModelVersion) TableName() string {
	return "bpm_model_version"
}

// FormDefinition 表单定义
type FormDefinition struct {
	BPMBaseModel
	FormCode   string    `json:"form_code" gorm:"size:50;uniqueIndex;not null"` // 表单编码
	FormName   string    `json:"form_name" gorm:"size:200;not null"`             // 表单名称
	FormType   string    `json:"form_type" gorm:"size:20;not null"`             // NORMAL/TEMPLATE
	Version    string    `json:"version" gorm:"size:20;default:1.0"`
	Category   *string   `json:"category" gorm:"size:50"`
	IsPublished int      `json:"is_published" gorm:"default:0"`
	PublishedAt *time.Time `json:"published_at"`
	TenantID   int64     `json:"tenant_id" gorm:"index;not null"`
	CreatedBy  *string   `json:"created_by" gorm:"size:50"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Fields     []FormField `json:"fields" gorm:"foreignKey:FormID"`
}

func (FormDefinition) TableName() string {
	return "bpm_form_definition"
}

// FormField 表单字段定义
type FormField struct {
	ID              uint           `json:"id" gorm:"primarykey"`
	FormID         int64          `json:"form_id" gorm:"not null;index"`
	FieldCode      string         `json:"field_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_form_field"`
	FieldName      string         `json:"field_name" gorm:"size:100;not null"`
	FieldType      string         `json:"field_type" gorm:"size:30;not null"` // TEXT/NUMBER/DATE/SELECT/MULTI_SELECT/ATTACHMENT/SIGNATURE
	FieldConfig    json.RawMessage `json:"field_config" gorm:"type:jsonb;not null"` // 字段配置
	DefaultValue   *string        `json:"default_value" gorm:"type:text"`
	ValidationRules json.RawMessage `json:"validation_rules" gorm:"type:jsonb"`  // 自定义校验规则
	IsRequired     int            `json:"is_required" gorm:"default:0"`
	IsReadonly     int            `json:"is_readonly" gorm:"default:0"`
	IsHidden       int            `json:"is_hidden" gorm:"default:0"`
	SortOrder      int            `json:"sort_order" gorm:"default:0"`
	TenantID       int64          `json:"tenant_id" gorm:"index;not null"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime"`
}

func (FormField) TableName() string {
	return "bpm_form_field"
}

// FormVersion 表单版本历史
type FormVersion struct {
	BPMBaseModel
	FormID          int64           `json:"form_id" gorm:"not null;index"`
	Version         string          `json:"version" gorm:"size:20;not null"`
	ChangeDescription *string       `json:"change_description" gorm:"type:text"`
	FieldsSnapshot  json.RawMessage `json:"fields_snapshot" gorm:"type:jsonb"` // 字段快照
	PublishedBy    *int64          `json:"published_by"`
	PublishedAt    *time.Time       `json:"published_at"`
	IsCurrent      int             `json:"is_current" gorm:"default:0"`
	TenantID       int64           `json:"tenant_id" gorm:"index;not null"`
}

func (FormVersion) TableName() string {
	return "bpm_form_version"
}

// ProcessInstance 流程实例
type ProcessInstance struct {
	BPMBaseModel
	InstanceNo    string          `json:"instance_no" gorm:"size:50;uniqueIndex;not null"` // 实例编号
	ModelID      int64           `json:"model_id" gorm:"not null;index"`
	ModelCode    *string         `json:"model_code" gorm:"size:50"`
	ModelName    *string         `json:"model_name" gorm:"size:200"`
	Version      *string         `json:"version" gorm:"size:20"`
	BizType      *string         `json:"biz_type" gorm:"size:50"`               // 业务类型
	BizID        *int64          `json:"biz_id"`                                // 业务单据ID
	BizNo        *string         `json:"biz_no" gorm:"size:100"`                // 业务单据号
	Title        string          `json:"title" gorm:"size:500;not null"`        // 流程标题
	InitiatorID  int64           `json:"initiator_id" gorm:"not null"`          // 发起人
	InitiatorName *string        `json:"initiator_name" gorm:"size:50"`
	CurrentNodeID *string        `json:"current_node_id" gorm:"size:50"`        // 当前节点ID
	CurrentNodeName *string      `json:"current_node_name" gorm:"size:200"`     // 当前节点名称
	Status       string          `json:"status" gorm:"size:20;not null"`       // DRAFT/RUNNING/SUSPENDED/CANCELLED/COMPLETED/TERMINATED
	Priority     int             `json:"priority" gorm:"default:5"`             // 优先级 1-5
	DueDate      *time.Time      `json:"due_date"`                             // 期望完成时间
	CompletedAt  *time.Time      `json:"completed_at"`
	CancelReason *string         `json:"cancel_reason" gorm:"type:text"`
	FormData    json.RawMessage `json:"form_data" gorm:"type:jsonb"`           // 表单数据快照
	BusinessData json.RawMessage `json:"business_data" gorm:"type:jsonb"`      // 业务数据快照
	TenantID    int64           `json:"tenant_id" gorm:"index;not null"`
	UpdatedAt   *time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	Tasks       []TaskInstance  `json:"tasks" gorm:"foreignKey:InstanceID"`
}

func (ProcessInstance) TableName() string {
	return "bpm_process_instance"
}

// TaskInstance 任务节点实例
type TaskInstance struct {
	ID             uint           `json:"id" gorm:"primarykey"`
	TaskNo        string         `json:"task_no" gorm:"size:50;uniqueIndex;not null"` // 任务编号
	InstanceID    int64          `json:"instance_id" gorm:"not null;index"`
	NodeID        string         `json:"node_id" gorm:"size:50;not null"`             // 节点定义ID
	NodeName      string         `json:"node_name" gorm:"size:200;not null"`
	NodeType      string         `json:"node_type" gorm:"size:30;not null"`
	TaskType      string         `json:"task_type" gorm:"size:20;default:APPROVAL"`   // APPROVAL/COPY/ACTION
	AssignType    string         `json:"assign_type" gorm:"size:20;not null"`        // ASSIGN/ROLE/INITIATOR/SCRIPT
	AssignValue   *string        `json:"assign_value" gorm:"size:200"`                // 分配值
	AssigneeID    *int64         `json:"assignee_id"`                                // 当前处理人
	AssigneeName  *string        `json:"assignee_name" gorm:"size:50"`
	AssigneeList  json.RawMessage `json:"assignee_list" gorm:"type:jsonb"`           // 所有处理人列表
	ActionResult  *string        `json:"action_result" gorm:"size:20"`                // 当前动作
	ActionComment *string        `json:"action_comment" gorm:"type:text"`            // 审批意见
	ActionTime    *time.Time     `json:"action_time"`
	SignType      *string        `json:"sign_type" gorm:"size:20"`                  // 会签类型
	RequiredApprovers int        `json:"required_approvers" gorm:"default:1"`       // 需审批人数
	CurrentApprovers  int        `json:"current_approvers" gorm:"default:0"`        // 当前审批人数
	Status        string         `json:"status" gorm:"size:20;not null"`            // PENDING/IN_PROGRESS/COMPLETED/CANCELLED/SKIPPED
	IsCurrent     int            `json:"is_current" gorm:"default:0"`               // 是否当前任务
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	CompletedAt   *time.Time     `json:"completed_at"`
}

func (TaskInstance) TableName() string {
	return "bpm_task_instance"
}

// ApprovalRecord 审批记录
type ApprovalRecord struct {
	BPMBaseModel
	TaskID        int64     `json:"task_id" gorm:"not null;index"`
	InstanceID    int64     `json:"instance_id" gorm:"not null;index"`
	ApproverID   int64     `json:"approver_id" gorm:"not null"`
	ApproverName string    `json:"approver_name" gorm:"size:50;not null"`
	ApproverDept *string   `json:"approver_dept" gorm:"size:100"`
	Action       string    `json:"action" gorm:"size:20;not null"` // AGREE/REJECT/TRANSFER/WITHDRAW/ADD_SIGN/REMOVE_SIGN
	Comment      *string   `json:"comment" gorm:"type:text"`
	ActionTime   time.Time `json:"action_time" gorm:"not null"`
	AssigneeID   *int64    `json:"assignee_id"`                  // 转交给谁
	AssigneeName *string   `json:"assignee_name" gorm:"size:50"`
	IsRollback   int       `json:"is_rollback" gorm:"default:0"` // 是否回退
	RollbackTarget *string `json:"rollback_target" gorm:"size:50"` // 回退目标节点
	TenantID     int64     `json:"tenant_id" gorm:"index;not null"`
}

func (ApprovalRecord) TableName() string {
	return "bpm_approval_record"
}

// DelegateRecord 流程委托记录
type DelegateRecord struct {
	BPMBaseModel
	DelegateID   int64           `json:"delegate_id" gorm:"not null"`         // 委托人
	DelegateName *string         `json:"delegate_name" gorm:"size:50"`
	DelegateeID  int64           `json:"delegatee_id" gorm:"not null"`        // 受托人
	DelegateeName *string        `json:"delegatee_name" gorm:"size:50"`
	StartDate    *time.Time      `json:"start_date" gorm:"type:date"`
	EndDate      *time.Time      `json:"end_date" gorm:"type:date"`
	BizTypes     json.RawMessage `json:"biz_types" gorm:"type:jsonb"`        // 委托的业务类型
	IsActive     int             `json:"is_active" gorm:"default:1"`
	TenantID     int64           `json:"tenant_id" gorm:"index;not null"`
}

func (DelegateRecord) TableName() string {
	return "bpm_delegate_record"
}

// ReminderRecord 流程催办记录
type ReminderRecord struct {
	BPMBaseModel
	InstanceID   int64     `json:"instance_id" gorm:"not null;index"`
	TaskID      *int64    `json:"task_id" gorm:"index"`
	RemindType  string    `json:"remind_type" gorm:"size:20;not null"`  // AUTO/MANUAL
	RemindFrom  *string   `json:"remind_from" gorm:"size:50"`
	RemindTo    *string   `json:"remind_to" gorm:"size:50"`
	RemindTime  time.Time `json:"remind_time" gorm:"not null"`
	RemindCount int       `json:"remind_count" gorm:"default:1"`
	TenantID    int64     `json:"tenant_id" gorm:"index;not null"`
}

func (ReminderRecord) TableName() string {
	return "bpm_reminder_record"
}

// TaskAssignment 任务分配规则
type TaskAssignment struct {
	ID            uint      `json:"id" gorm:"primarykey"`
	ModelID      int64     `json:"model_id" gorm:"not null;index"`
	NodeID      string    `json:"node_id" gorm:"size:50;not null"`
	RuleType    string    `json:"rule_type" gorm:"size:20;not null"`   // FIXED/ROLE/DEPT_HEAD/INITIATOR/SCRIPT/DYNAMIC
	RuleValue   string    `json:"rule_value" gorm:"type:text;not null"` // 规则值
	Priority    int       `json:"priority" gorm:"default:5"`
	AllowAssign int       `json:"allow_assign" gorm:"default:1"`     // 允许转交
	AllowTransfer int     `json:"allow_transfer" gorm:"default:1"`    // 允许转派
	TenantID    int64     `json:"tenant_id" gorm:"index;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (TaskAssignment) TableName() string {
	return "bpm_task_assignment"
}

// NodeTimeout 节点超时配置
type NodeTimeout struct {
	ID               uint      `json:"id" gorm:"primarykey"`
	ModelID         int64     `json:"model_id" gorm:"not null;index"`
	NodeID         string    `json:"node_id" gorm:"size:50;not null"`
	TimeoutMinutes  int       `json:"timeout_minutes" gorm:"default:0"`   // 超时时间(分钟)
	TimeoutAction   *string   `json:"timeout_action" gorm:"size:20"`     // NOTIFY/ESCALATE/AUTO_APPROVE/TERMINATE
	TimeoutNoticeRoles json.RawMessage `json:"timeout_notice_roles" gorm:"type:jsonb"` // 超时通知角色
	EscalationNodeID *string  `json:"escalation_node_id" gorm:"size:50"` // 升级目标节点
	TenantID       int64     `json:"tenant_id" gorm:"index;not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (NodeTimeout) TableName() string {
	return "bpm_node_timeout"
}

// BizMapping 业务模块与流程绑定
type BizMapping struct {
	BPMBaseModel
	BizType      string    `json:"biz_type" gorm:"size:50;not null;uniqueIndex"` // NCR/PURCHASE_ORDER/EQUIPMENT_REPAIR
	ModelID      int64     `json:"model_id" gorm:"not null"`
	TriggerEvent string    `json:"trigger_event" gorm:"size:50;not null"`       // CREATE/SUBMIT/APPROVED
	IsActive     int       `json:"is_active" gorm:"default:1"`
	TenantID     int64     `json:"tenant_id" gorm:"index;not null"`
}

func (BizMapping) TableName() string {
	return "bpm_biz_mapping"
}
