package model

import (
	"encoding/json"
	"time"
)

// ElectronicSOP 电子SOP
type ElectronicSOP struct {
	BaseModel
	TenantID      int64           `json:"tenant_id" gorm:"index;not null"`
	SopNo        string          `json:"sop_no" gorm:"size:50;not null"` // SOP编号
	SopName      string          `json:"sop_name" gorm:"size:200;not null"` // SOP名称
	MaterialID   int64           `json:"material_id"` // 关联物料ID
	MaterialCode string          `json:"material_code" gorm:"size:50"` // 物料编码
	MaterialName string          `json:"material_name" gorm:"size:100"` // 物料名称
	Version      string          `json:"version" gorm:"size:20"` // 版本号
	ProcessID    int64           `json:"process_id"` // 工序ID
	ProcessName  string          `json:"process_name" gorm:"size:100"` // 工序名称
	ContentType  string          `json:"content_type" gorm:"size:20"` // PDF/VIDEO/IMAGE/HTML
	ContentURL   string          `json:"content_url" gorm:"size:500"` // 内容URL
	ThumbnailURL string          `json:"thumbnail_url" gorm:"size:500"` // 缩略图URL
	Steps        json.RawMessage `json:"steps" gorm:"type:jsonb"` // SOP步骤JSON
	WorkstationID int64         `json:"workstation_id"` // 适用工位ID
	WorkstationName string       `json:"workstation_name" gorm:"size:100"` // 工位名称
	WorkshopID   int64           `json:"workshop_id"` // 适用车间ID
	WorkshopName string          `json:"workshop_name" gorm:"size:100"` // 车间名称
	Status       int             `json:"status" gorm:"default:1"` // 1草稿/2已发布/3已作废
	EffDate      *time.Time      `json:"eff_date"` // 生效日期
	ExpDate      *time.Time      `json:"exp_date"` // 失效日期
	ApprovedBy   string          `json:"approved_by" gorm:"size:50"` // 审批人
	ApprovedAt   *time.Time      `json:"approved_at"` // 审批时间
	Remark       string          `json:"remark" gorm:"size:500"` // 备注
}

func (ElectronicSOP) TableName() string {
	return "mes_electronic_sop"
}

// SOPContentStep SOP步骤
type SOPContentStep struct {
	StepNo   int    `json:"step_no"`   // 步骤号
	Title    string `json:"title"`     // 步骤标题
	Content  string `json:"content"`    // 步骤内容
	ImageURL string `json:"image_url"` // 图片URL
	VideoURL string `json:"video_url"` // 视频URL
	Duration int    `json:"duration"`  // 预计时长(秒)
}
