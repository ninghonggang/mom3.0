package model

import (
	"time"
)

// ProductionDailyReport 生产日报
type ProductionDailyReport struct {
	BaseModel
	TenantID int64 `json:"tenant_id" gorm:"index;not null"`
	ReportDate            time.Time `json:"report_date" gorm:"index;not null"` // 报表日期
	WorkshopID            int64    `json:"workshop_id" gorm:"index"`
	WorkshopName          string   `json:"workshop_name" gorm:"size:100"`
	ProductionOrderCount  int      `json:"production_order_count"` // 生产工单数
	CompletedOrderCount   int      `json:"completed_order_count"`  // 完成工单数
	TotalOutputQty       float64  `json:"total_output_qty"`      // 总产出数量
	QualifiedQty          float64  `json:"qualified_qty"`         // 合格数量
	DefectQty             float64  `json:"defect_qty"`           // 不良数量
	PassRate              float64  `json:"pass_rate"`             // 合格率%
	FirstPassRate         float64  `json:"first_pass_rate"`      // 一次合格率%
	OEE                   float64  `json:"oee"`                  // 设备效率%
	OutputPerHour         float64  `json:"output_per_hour"`      // 每小时产量
	WorkerCount           int      `json:"worker_count"`         // 作业人数
	WorkingHours          float64  `json:"working_hours"`        // 作业工时
	Remark                string   `json:"remark" gorm:"size:500"`
}

func (ProductionDailyReport) TableName() string {
	return "production_daily_report"
}