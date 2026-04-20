package model

import (
	"time"
)

// WmsLabelTemplate 标签模板
type WmsLabelTemplate struct {
	ID           uint64    `json:"id" gorm:"primaryKey"`
	TenantID     int64     `json:"tenant_id" gorm:"index"`
	TemplateCode string    `json:"template_code" gorm:"size:50;uniqueIndex:idx_tenant_code"`
	TemplateName string    `json:"template_name" gorm:"size:100"`
	TemplateType string    `json:"template_type" gorm:"size:20"`  // ITEM/INVENTORY/LOCATION/PALLET
	Width        float64   `json:"width" gorm:"type:decimal(10,2)"`
	Height       float64   `json:"height" gorm:"type:decimal(10,2)"`
	Content      string    `json:"content" gorm:"type:text"`  // JSON格式标签内容
	Status       string    `json:"status" gorm:"size:20"`
	CreatedBy    uint64    `json:"created_by"`
	UpdatedBy    uint64    `json:"updated_by"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (WmsLabelTemplate) TableName() string {
	return "wms_label_template"
}

// WmsLabelTemplateCreateReqVO 标签模板创建请求
type WmsLabelTemplateCreateReqVO struct {
	TemplateCode string  `json:"template_code"`
	TemplateName string  `json:"template_name"`
	TemplateType string  `json:"template_type"`
	Width        float64 `json:"width"`
	Height       float64 `json:"height"`
	Content      string  `json:"content"`
	Status       string  `json:"status"`
}

// WmsLabelTemplateUpdateReqVO 标签模板更新请求
type WmsLabelTemplateUpdateReqVO struct {
	Id           uint64  `json:"id"`
	TemplateName string  `json:"template_name"`
	TemplateType string  `json:"template_type"`
	Width        float64 `json:"width"`
	Height       float64 `json:"height"`
	Content      string  `json:"content"`
	Status       string  `json:"status"`
}

// WmsLabelTemplateQueryVO 标签模板查询请求
type WmsLabelTemplateQueryVO struct {
	Keyword      string `form:"keyword"`
	TemplateType string `form:"template_type"`
	Status       string `form:"status"`
	Page         int    `form:"page"`
	PageSize     int    `form:"page_size"`
}
