package model

import (
	"time"
)

// IdocRecord IDOC记录
type IdocRecord struct {
	ID            int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID      int64     `json:"tenant_id" gorm:"index;not null;default:1"`
	IdocNumber    string    `json:"idoc_number" gorm:"size:64;uniqueIndex"` // IDOC编号
	IdocType      string    `json:"idoc_type" gorm:"size:30;index"`        // IDOC类型：MATMAS/ORDERS/DESADV等
	Direction     int       `json:"direction" gorm:"default:1"`             // 方向：1接收 2发送
	Status        int       `json:"status" gorm:"default:1"`               // 状态：1新建 2处理中 3成功 4失败
	PartnerType   string    `json:"partner_type" gorm:"size:20"`            // 伙伴类型：LI供应商 KU客户 WE工厂
	PartnerNo     string    `json:"partner_no" gorm:"size:20"`             // 伙伴编号
	MessageType   string    `json:"message_type" gorm:"size:30"`           // 消息类型
	ReferenceNo   string    `json:"reference_no" gorm:"size:64"`           // 参考编号
	RawContent    string    `json:"raw_content" gorm:"type:text"`           // 原始IDOC内容
	ParsedData    string    `json:"parsed_data" gorm:"type:text"`           // 解析后的数据(JSON)
	ErrorMessage  string    `json:"error_message" gorm:"type:text"`        // 错误信息
	RetryCount    int       `json:"retry_count" gorm:"default:0"`           // 重试次数
	ProcessedAt   *time.Time `json:"processed_at"`                         // 处理时间
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (IdocRecord) TableName() string {
	return "idoc_record"
}

// IdocTypeConfig IDOC类型配置
type IdocTypeConfig struct {
	ID           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID     int64     `json:"tenant_id" gorm:"index;not null;default:1"`
	IdocType     string    `json:"idoc_type" gorm:"size:30;uniqueIndex"` // IDOC类型
	MessageType  string    `json:"message_type" gorm:"size:30"`          // 消息类型
	Description  string    `json:"description" gorm:"size:200"`        // 描述
	Endpoint     string    `json:"endpoint" gorm:"size:200"`             // 接收/发送地址
	IsActive     int       `json:"is_active" gorm:"default:1"`           // 是否启用
	MappingRule  string    `json:"mapping_rule" gorm:"type:text"`       // 映射规则(JSON)
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (IdocTypeConfig) TableName() string {
	return "idoc_type_config"
}

// IdocQuery IDOC查询
type IdocQuery struct {
	IdocType    string `form:"idoc_type"`
	Direction   int    `form:"direction"`
	Status      int    `form:"status"`
	PartnerNo   string `form:"partner_no"`
	StartDate   string `form:"start_date"`
	EndDate     string `form:"end_date"`
	Page        int    `form:"page"`
	PageSize    int    `form:"page_size"`
}

// IdocReceiveReq IDOC接收请求
type IdocReceiveReq struct {
	IdocType    string `json:"idoc_type" binding:"required"`
	PartnerType string `json:"partner_type"`
	PartnerNo   string `json:"partner_no"`
	MessageType string `json:"message_type"`
	ReferenceNo string `json:"reference_no"`
	RawContent  string `json:"raw_content" binding:"required"`
}

// IdocSendReq IDOC发送请求
type IdocSendReq struct {
	IdocType    string `json:"idoc_type" binding:"required"`
	TargetType  string `json:"target_type"`
	TargetNo    string `json:"target_no"`
	MessageType string `json:"message_type"`
	Data        map[string]interface{} `json:"data" binding:"required"`
}