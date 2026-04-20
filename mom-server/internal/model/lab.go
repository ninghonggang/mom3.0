package model

import (
	"time"
)

// LabSample 检测样品
type LabSample struct {
	ID             uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID       uint64    `gorm:"index;not null;default:1" json:"tenant_id"`
	SampleCode     string    `gorm:"size:50;uniqueIndex" json:"sample_code"`      // 样品编号
	SampleName     string    `gorm:"size:100" json:"sample_name"`                // 样品名称
	InspectionType string    `gorm:"size:20;index" json:"inspection_type"`       // IQC/PQC/FQC
	SourceType     string    `gorm:"size:20" json:"source_type"`                 // PURCHASE/PRODUCT/ORDER
	SourceID       uint64    `gorm:"index" json:"source_id"`                     // 来源ID（采购单/工单ID）
	SourceNo       string    `gorm:"size:50" json:"source_no"`                   // 来源单号
	Quantity       float64   `json:"quantity"`                                   // 送检数量
	SampleQty      float64   `json:"sample_qty"`                                 // 抽样数量
	ReceivedBy     string    `gorm:"size:50" json:"received_by"`                // 接收人
	ReceivedAt     time.Time `json:"received_at"`                               // 接收时间
	DueDate        time.Time `json:"due_date"`                                   // 要求完成日期
	Status         string    `gorm:"size:20;index" json:"status"`               // PENDING/INSPECTING/COMPLETED/CANCELLED
	Remark         string    `gorm:"size:500" json:"remark"`
	CreatedBy      string    `gorm:"size:50" json:"created_by"`
	UpdatedBy      string    `gorm:"size:50" json:"updated_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (LabSample) TableName() string {
	return "lab_sample"
}

// LabSampleCreateRequest 创建检测样品请求
type LabSampleCreateRequest struct {
	SampleName     string    `json:"sample_name" binding:"required"`
	InspectionType string    `json:"inspection_type" binding:"required"`
	SourceType     string    `json:"source_type"`
	SourceID       uint64    `json:"source_id"`
	SourceNo       string    `json:"source_no"`
	Quantity       float64   `json:"quantity"`
	SampleQty      float64   `json:"sample_qty"`
	ReceivedBy     string    `json:"received_by"`
	ReceivedAt     time.Time `json:"received_at"`
	DueDate        time.Time `json:"due_date"`
	Remark         string    `json:"remark"`
}

// LabSampleUpdateRequest 更新检测样品请求
type LabSampleUpdateRequest struct {
	SampleName     string    `json:"sample_name"`
	InspectionType string    `json:"inspection_type"`
	SourceType     string    `json:"source_type"`
	SourceID       uint64    `json:"source_id"`
	SourceNo       string    `json:"source_no"`
	Quantity       float64   `json:"quantity"`
	SampleQty      float64   `json:"sample_qty"`
	ReceivedBy     string    `json:"received_by"`
	ReceivedAt     time.Time `json:"received_at"`
	DueDate        time.Time `json:"due_date"`
	Status         string    `json:"status"`
	Remark         string    `json:"remark"`
}

// LabSampleQuery 查询检测样品
type LabSampleQuery struct {
	SampleCode     string `form:"sample_code"`
	SampleName     string `form:"sample_name"`
	InspectionType string `form:"inspection_type"`
	Status         string `form:"status"`
	Page           int    `form:"page"`
	PageSize       int    `form:"page_size"`
}

// LabTestItem 检测项目
type LabTestItem struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID    uint64    `gorm:"index;not null;default:1" json:"tenant_id"`
	SampleID    uint64    `gorm:"index;not null" json:"sample_id"`
	ItemName    string    `gorm:"size:100" json:"item_name"`             // 检测项目名称
	TestMethod  string    `gorm:"size:100" json:"test_method"`          // 检测方法
	Standard    string    `gorm:"size:100" json:"standard"`             // 判定标准
	UpperLimit  float64  `json:"upper_limit"`                          // 上限
	LowerLimit  float64  `json:"lower_limit"`                          // 下限
	Unit        string    `gorm:"size:20" json:"unit"`                  // 单位
	Result      string    `gorm:"size:20" json:"result"`                // PASS/FAIL/NA
	ResultValue float64  `json:"result_value"`                         // 检测值
	TesterID    uint64    `json:"tester_id"`
	TesterName  string    `gorm:"size:50" json:"tester_name"`
	TestedAt    time.Time `json:"tested_at"`
	Remark      string    `gorm:"size:200" json:"remark"`
	CreatedAt   time.Time `json:"created_at"`
}

func (LabTestItem) TableName() string {
	return "lab_test_item"
}

// LabTestItemCreateRequest 创建检测项目请求
type LabTestItemCreateRequest struct {
	ItemName    string  `json:"item_name" binding:"required"`
	TestMethod  string  `json:"test_method"`
	Standard    string  `json:"standard"`
	UpperLimit  float64 `json:"upper_limit"`
	LowerLimit  float64 `json:"lower_limit"`
	Unit        string  `json:"unit"`
	ResultValue float64 `json:"result_value"`
	Result      string  `json:"result"`
	TesterID    uint64  `json:"tester_id"`
	TesterName  string  `json:"tester_name"`
	TestedAt    time.Time `json:"tested_at"`
	Remark      string  `json:"remark"`
}

// LabReport 检测报告
type LabReport struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID     uint64    `gorm:"index;not null;default:1" json:"tenant_id"`
	ReportNo     string    `gorm:"size:50;uniqueIndex" json:"report_no"` // 报告编号
	SampleID     uint64    `gorm:"index;not null" json:"sample_id"`
	Conclusion   string    `gorm:"size:20" json:"conclusion"`            // QUALIFIED/UNQUALIFIED
	Remarks      string    `gorm:"type:text" json:"remarks"`             // 备注
	Attachments  string    `gorm:"type:text" json:"attachments"`        // 附件JSON
	InspectorID  uint64    `json:"inspector_id"`
	InspectorName string   `gorm:"size:50" json:"inspector_name"`
	ApprovedBy   string    `gorm:"size:50" json:"approved_by"` // 批准人
	ReportDate   time.Time `json:"report_date"`
	CreatedBy    string    `gorm:"size:50" json:"created_by"`
	UpdatedBy    string    `gorm:"size:50" json:"updated_by"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (LabReport) TableName() string {
	return "lab_report"
}

// LabReportQuery 查询检测报告
type LabReportQuery struct {
	ReportNo     string `form:"report_no"`
	Conclusion   string `form:"conclusion"`
	SampleID     uint64 `form:"sample_id"`
	InspectorID  uint64 `form:"inspector_id"`
	Page         int    `form:"page"`
	PageSize     int    `form:"page_size"`
}

// LabReportCreateRequest 创建检测报告请求
type LabReportCreateRequest struct {
	SampleID      uint64    `json:"sample_id" binding:"required"`
	Conclusion    string    `json:"conclusion" binding:"required"`
	Remarks       string    `json:"remarks"`
	Attachments   string    `json:"attachments"`
	InspectorID   uint64    `json:"inspector_id"`
	InspectorName string    `json:"inspector_name"`
	ReportDate    time.Time `json:"report_date"`
}
