package model

import (
	"time"
)

// AndonReport 安灯报表
type AndonReport struct {
	BaseModel
	TenantID int64 `json:"tenant_id" gorm:"index;not null"`
	ReportDate   time.Time `json:"report_date" gorm:"index;not null"`
	WorkshopID   int64    `json:"workshop_id" gorm:"index"`
	WorkshopName string   `json:"workshop_name" gorm:"size:100"`
	LineID       int64    `json:"line_id" gorm:"index"`
	LineName     string   `json:"line_name" gorm:"size:100"`
	StationID    int64    `json:"station_id" gorm:"index"`
	StationName  string   `json:"station_name" gorm:"size:100"`

	TotalCallCount     int     `json:"total_call_count"`      // 呼叫总数
	MaterialCallCount  int     `json:"material_call_count"`   // 物料呼叫
	QualityCallCount   int     `json:"quality_call_count"`    // 质量呼叫
	EquipmentCallCount int     `json:"equipment_call_count"`  // 设备呼叫
	OtherCallCount     int     `json:"other_call_count"`      // 其他呼叫
	AvgResponseTime    float64 `json:"avg_response_time"`     // 平均响应时间(分钟)
	AvgResolveTime     float64 `json:"avg_resolve_time"`     // 平均解决时间(分钟)
	UnresolvedCount    int     `json:"unresolved_count"`     // 未解决数
	Remark             string  `json:"remark" gorm:"size:500"`
}

func (AndonReport) TableName() string {
	return "andon_report"
}