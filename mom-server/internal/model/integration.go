package model

import (
	"time"
)

// InterfaceCategory 接口分类
const (
	InterfaceCategoryERP   = "ERP"
	InterfaceCategoryAGV    = "AGV"
	InterfaceCategoryMES   = "MES"
	InterfaceCategoryWMS   = "WMS"
	InterfaceCategoryOther = "OTHER"
)

// InterfaceDirection 接口方向
const (
	InterfaceDirectionOutbound = "OUTBOUND" // 推（主动推送外部系统）
	InterfaceDirectionInbound = "INBOUND"  // 拉（从外部系统拉取）
)

// HTTPMethod HTTP方法
const (
	HTTPMethodGET    = "GET"
	HTTPMethodPOST   = "POST"
	HTTPMethodPUT    = "PUT"
	HTTPMethodDELETE = "DELETE"
)

// AuthType 认证类型
const (
	AuthTypeNone      = "NONE"
	AuthTypeBasic     = "BASIC"
	AuthTypeAPIKey    = "API_KEY"
	AuthTypeOAuth2    = "OAUTH2"
	AuthTypeBearer    = "BEARER_TOKEN"
)

// ContentType 内容类型
const (
	ContentTypeJSON  = "JSON"
	ContentTypeXML   = "XML"
	ContentTypeFORM  = "FORM"
	ContentTypeTEXT  = "TEXT"
)

// TriggerType 触发类型
const (
	TriggerTypeManual   = "MANUAL"
	TriggerTypeSchedule = "SCHEDULE"
	TriggerTypeEvent   = "EVENT"
)

// ExecutionStatus 执行状态
const (
	ExecutionStatusSuccess = "SUCCESS"
	ExecutionStatusFailed  = "FAILED"
	ExecutionStatusPartial = "PARTIAL"
)

// SourceType 数据源类型
const (
	SourceTypeTableQuery = "TABLE_QUERY"
	SourceTypeAPICall    = "API_CALL"
	SourceTypeEventPayload = "EVENT_PAYLOAD"
)

// MapType 映射类型
const (
	MapTypeConst     = "CONST"
	MapTypeField     = "FIELD"
	MapTypeExpr      = "EXPRESSION"
	MapTypeJSONPath  = "JSONPATH"
)

// Predefined event types
const (
	EventTypeProductionComplete = "PRODUCTION_COMPLETE"
	EventTypeQualityInspect    = "QUALITY_INSPECT"
	EventTypeStockIn           = "STOCK_IN"
	EventTypeStockOut          = "STOCK_OUT"
	EventTypePurchaseAward    = "PURCHASE_AWARD"
	EventTypePurchaseReceive   = "PURCHASE_RECEIVE"
	EventTypeSalesShip         = "SALES_SHIP"
)

// InterfaceConfig 接口配置
type InterfaceConfig struct {
	BaseModel
	TenantID    int64  `json:"tenant_id" gorm:"index;not null"`
	Name        string `json:"name" gorm:"size:100;not null"`         // 接口名称
	Code        string `json:"code" gorm:"size:50;not null;uniqueIndex:idx_tenant_code"` // 唯一编码
	Category    string `json:"category" gorm:"size:20"`              // ERP/AGV/MES/WMS/OTHER
	Description string `json:"description" gorm:"size:500"`           // 描述

	// 调用配置
	Direction            string `json:"direction" gorm:"size:20"`         // OUTBOUND/INBOUND
	Method               string `json:"method" gorm:"size:10"`            // GET/POST/PUT/DELETE
	BaseURL              string `json:"base_url" gorm:"size:500"`         // 基础URL
	Path                 string `json:"path" gorm:"size:500"`             // 请求路径
	AuthType             string `json:"auth_type" gorm:"size:20"`        // 认证类型
	AuthConfig           string `json:"auth_config" gorm:"type:text"`    // 认证配置JSON
	RequestContentType   string `json:"request_content_type" gorm:"size:20"` // 请求内容类型
	RequestBodyTemplate  string `json:"request_body_template" gorm:"type:text"` // 请求体模板
	ResponseFormat       string `json:"response_format" gorm:"size:20"`  // 响应格式

	// 超时与重试
	Timeout       int `json:"timeout" gorm:"default:30"`      // 超时秒数
	RetryCount    int `json:"retry_count" gorm:"default:0"`  // 重试次数
	RetryInterval int `json:"retry_interval" gorm:"default:3"` // 重试间隔秒数

	// 数据源配置
	SourceType          string `json:"source_type" gorm:"size:20"`           // TABLE_QUERY/API_CALL/EVENT_PAYLOAD
	SourceTable         string `json:"source_table" gorm:"size:100"`         // 源表名
	SourceAPI           string `json:"source_api" gorm:"size:200"`          // 源API路径
	SourceFilter        string `json:"source_filter" gorm:"type:text"`       // 查询条件/参数
	SourceFields        string `json:"source_fields" gorm:"size:500"`        // 查询字段
	PrimaryKey          string `json:"primary_key" gorm:"size:50"`           // 主键字段
	IncrementalField    string `json:"incremental_field" gorm:"size:50"`     // 增量时间字段
	IncrementalWindow   int    `json:"incremental_window" gorm:"default:0"`   // 增量窗口（分钟）

	Status    string `json:"status" gorm:"size:20;default:ENABLE"` // ENABLE/DISABLE
	Remark    string `json:"remark" gorm:"size:500"`
}

func (InterfaceConfig) TableName() string {
	return "sys_interface_config"
}

// InterfaceFieldMap 字段映射
type InterfaceFieldMap struct {
	BaseModel
	InterfaceConfigID int64  `json:"interface_config_id" gorm:"index;not null"`
	FieldName        string `json:"field_name" gorm:"size:100;not null"` // 目标字段名
	FieldType        string `json:"field_type" gorm:"size:20"`            // STRING/NUMBER/DATE/BOOL
	MapType          string `json:"map_type" gorm:"size:20"`             // CONST/FIELD/EXPRESSION/JSONPATH
	MapValue         string `json:"map_value" gorm:"size:500"`           // 映射值/路径/表达式
	Required         bool   `json:"required" gorm:"default:false"`
	DefaultValue     string `json:"default_value" gorm:"size:200"`       // 默认值
	TransformFunc    string `json:"transform_func" gorm:"size:100"`      // 转换函数
	SortOrder       int    `json:"sort_order" gorm:"default:0"`         // 排序
}

func (InterfaceFieldMap) TableName() string {
	return "sys_interface_field_map"
}

// InterfaceTrigger 触发配置
type InterfaceTrigger struct {
	BaseModel
	InterfaceConfigID int64  `json:"interface_config_id" gorm:"index;not null"`
	TriggerType      string `json:"trigger_type" gorm:"size:20;not null"` // MANUAL/SCHEDULE/EVENT
	CronExpr        string `json:"cron_expr" gorm:"size:50"`            // Cron表达式
	EventSource     string `json:"event_source" gorm:"size:50"`         // 事件类型
	PayloadTemplate string `json:"payload_template" gorm:"type:text"`  // 事件数据模板
	Condition       string `json:"condition" gorm:"size:500"`           // 触发条件表达式
	FallbackMinutes  int    `json:"fallback_minutes" gorm:"default:0"`   // 兜底触发间隔（分钟）
	Status          string `json:"status" gorm:"size:20;default:ENABLE"`
}

func (InterfaceTrigger) TableName() string {
	return "sys_interface_trigger"
}

// InterfaceExecutionLog 执行日志
type InterfaceExecutionLog struct {
	BaseModel
	InterfaceConfigID int64     `json:"interface_config_id" gorm:"index"`
	ConfigName       string    `json:"config_name" gorm:"size:100"`     // 配置名称（冗余）
	TriggerType      string    `json:"trigger_type" gorm:"size:20"`     // 触发类型
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	Duration         int64     `json:"duration"`           // 耗时ms
	RequestURL       string    `json:"request_url" gorm:"size:500"`
	RequestMethod    string    `json:"request_method" gorm:"size:10"`
	RequestHeaders   string    `json:"request_headers" gorm:"type:text"`
	RequestBody      string    `json:"request_body" gorm:"type:text"`
	ResponseBody     string    `json:"response_body" gorm:"type:text"`
	ResponseCode     int       `json:"response_code"`
	Status           string    `json:"status" gorm:"size:20"`   // SUCCESS/FAILED/PARTIAL
	ErrorMessage     string    `json:"error_message" gorm:"size:1000"`
	RetryCount       int       `json:"retry_count" gorm:"default:0"`
	RecordsProcessed int       `json:"records_processed"` // 处理记录数
}

func (InterfaceExecutionLog) TableName() string {
	return "sys_interface_exec_log"
}

// InterfaceFieldMapCreate 字段映射创建请求
type InterfaceFieldMapCreate struct {
	FieldName    string `json:"field_name" binding:"required"`
	FieldType    string `json:"field_type" binding:"required"`
	MapType      string `json:"map_type" binding:"required"`
	MapValue     string `json:"map_value"`
	Required     bool   `json:"required"`
	DefaultValue string `json:"default_value"`
	TransformFunc string `json:"transform_func"`
	SortOrder    int    `json:"sort_order"`
}

// InterfaceTriggerCreate 触发配置创建请求
type InterfaceTriggerCreate struct {
	TriggerType      string `json:"trigger_type" binding:"required"`
	CronExpr        string `json:"cron_expr"`
	EventSource     string `json:"event_source"`
	PayloadTemplate string `json:"payload_template"`
	Condition       string `json:"condition"`
	FallbackMinutes int    `json:"fallback_minutes"`
}

// InterfaceConfigCreate 接口配置创建请求
type InterfaceConfigCreate struct {
	Name               string                    `json:"name" binding:"required"`
	Code              string                    `json:"code" binding:"required"`
	Category          string                    `json:"category"`
	Description       string                    `json:"description"`
	Direction         string                    `json:"direction"`
	Method            string                    `json:"method"`
	BaseURL           string                    `json:"base_url"`
	Path              string                    `json:"path"`
	AuthType          string                    `json:"auth_type"`
	AuthConfig        string                    `json:"auth_config"`
	RequestContentType string                  `json:"request_content_type"`
	RequestBodyTemplate string                  `json:"request_body_template"`
	ResponseFormat    string                    `json:"response_format"`
	Timeout           int                       `json:"timeout"`
	RetryCount        int                       `json:"retry_count"`
	RetryInterval     int                       `json:"retry_interval"`
	SourceType        string                    `json:"source_type"`
	SourceTable       string                    `json:"source_table"`
	SourceAPI         string                    `json:"source_api"`
	SourceFilter      string                    `json:"source_filter"`
	SourceFields      string                    `json:"source_fields"`
	PrimaryKey        string                    `json:"primary_key"`
	IncrementalField  string                    `json:"incremental_field"`
	IncrementalWindow int                       `json:"incremental_window"`
	FieldMaps         []InterfaceFieldMapCreate  `json:"field_maps"`
	Triggers          []InterfaceTriggerCreate  `json:"triggers"`
	Remark            string                    `json:"remark"`
}

// InterfaceConfigUpdate 接口配置更新请求
type InterfaceConfigUpdate struct {
	Name               string                   `json:"name"`
	Description        string                   `json:"description"`
	Direction          string                   `json:"direction"`
	Method             string                   `json:"method"`
	BaseURL            string                   `json:"base_url"`
	Path               string                   `json:"path"`
	AuthType           string                   `json:"auth_type"`
	AuthConfig         string                   `json:"auth_config"`
	RequestContentType string                   `json:"request_content_type"`
	RequestBodyTemplate string                  `json:"request_body_template"`
	ResponseFormat     string                   `json:"response_format"`
	Timeout            int                      `json:"timeout"`
	RetryCount         int                      `json:"retry_count"`
	RetryInterval      int                      `json:"retry_interval"`
	SourceType         string                   `json:"source_type"`
	SourceTable        string                   `json:"source_table"`
	SourceAPI          string                   `json:"source_api"`
	SourceFilter       string                   `json:"source_filter"`
	SourceFields       string                   `json:"source_fields"`
	PrimaryKey         string                   `json:"primary_key"`
	IncrementalField   string                   `json:"incremental_field"`
	IncrementalWindow  int                      `json:"incremental_window"`
	Status             string                   `json:"status"`
	Remark             string                   `json:"remark"`
}

// InterfaceConfigQuery 查询条件
type InterfaceConfigQuery struct {
	TenantID  int64
	Category  string
	Status    string
	Keyword   string // 模糊匹配 name/code
	Page      int
	PageSize  int
}
