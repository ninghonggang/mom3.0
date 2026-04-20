package model

import (
	"time"
)

// QualityWeeklyReport 质量周报
type QualityWeeklyReport struct {
	BaseModel
	TenantID int64 `json:"tenant_id" gorm:"index;not null"`
	ReportYear   int       `json:"report_year" gorm:"index"`
	ReportWeek   int       `json:"report_week" gorm:"index"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	WorkshopID   int64     `json:"workshop_id" gorm:"index"`
	WorkshopName string    `json:"workshop_name" gorm:"size:100"`

	TotalInspectionQty int     `json:"total_inspection_qty"` // 总检验数
	QualifiedQty       int     `json:"qualified_qty"`        // 合格数
	DefectQty          int     `json:"defect_qty"`           // 不良数
	PassRate           float64 `json:"pass_rate"`            // 合格率%

	// IQC
	IQCInspQty      int `json:"iqc_insp_qty"`
	IQCQualifiedQty int `json:"iqc_qualified_qty"`
	IQCDefectQty    int `json:"iqc_defect_qty"`

	// IPQC
	IPQCInspQty      int `json:"ipqc_insp_qty"`
	IPQCQualifiedQty int `json:"ipqc_qualified_qty"`
	IPQCDefectQty    int `json:"ipqc_defect_qty"`

	// FQC
	FQCInspQty      int `json:"fqc_insp_qty"`
	FQCQualifiedQty int `json:"fqc_qualified_qty"`
	FQCDefectQty    int `json:"fqc_defect_qty"`

	// OQC
	OQCInspQty      int `json:"oqc_insp_qty"`
	OQCQualifiedQty int `json:"oqc_qualified_qty"`
	OQCDefectQty    int `json:"oqc_defect_qty"`

	NCRCount             int    `json:"ncr_count"`              // NCR数量
	CustomerComplaintCount int    `json:"customer_complaint_count"` // 客诉数量
	Remark               string `json:"remark" gorm:"size:500"`
}

func (QualityWeeklyReport) TableName() string {
	return "quality_weekly_report"
}