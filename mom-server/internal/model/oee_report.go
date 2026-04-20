package model

import (
	"time"
)

// OEEReport OEE报表
type OEEReport struct {
	BaseModel
	TenantID int64 `json:"tenant_id" gorm:"index;not null"`
	ReportDate   time.Time `json:"report_date" gorm:"index;not null"`
	WorkshopID   int64    `json:"workshop_id" gorm:"index"`
	WorkshopName string   `json:"workshop_name" gorm:"size:100"`
	LineID       int64    `json:"line_id" gorm:"index"`
	LineName     string   `json:"line_name" gorm:"size:100"`

	Availability float64 `json:"availability"` // 可用率%
	Performance  float64 `json:"performance"`  // 性能率%
	Quality      float64 `json:"quality"`       // 质量率%
	OEE          float64 `json:"oee"`           // 综合效率%

	PlannedProductionTime int `json:"planned_production_time"` // 计划生产时间(分钟)
	ActualProductionTime  int `json:"actual_production_time"`  // 实际生产时间(分钟)
	DownTime             int `json:"down_time"`               // 停机时间(分钟)
	SpeedLoss            float64 `json:"speed_loss"`           // 速度损失%
	DefectLoss           float64 `json:"defect_loss"`          // 不良损失%

	Remark string `json:"remark" gorm:"size:500"`
}

func (OEEReport) TableName() string {
	return "oee_report"
}