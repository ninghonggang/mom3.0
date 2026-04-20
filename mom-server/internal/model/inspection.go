package model

import (
	"time"
)

// QualityInspectionPlan 检验计划（主表）
type QualityInspectionPlan struct {
	ID              uint64  `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID        uint64  `gorm:"index;not null;default:1" json:"tenant_id"`
	PlanCode        string  `gorm:"size:50;uniqueIndex" json:"plan_code"`       // 检验计划编号
	PlanName        string  `gorm:"size:100;not null" json:"plan_name"`         // 检验计划名称
	InspectionType  string  `gorm:"size:20;index" json:"inspection_type"`       // IQC/PQC/OQC/FQC
	AQLLevel        string  `gorm:"size:10" json:"aql_level"`                    // AQL等级如 I/II/III
	SampleSize      int     `json:"sample_size"`                                // 抽样数量
	BatchMin        int     `json:"batch_min"`                                  // 批次最小数量
	BatchMax        int     `json:"batch_max"`                                  // 批次最大数量
	AcCount         int     `json:"ac_count"`                                   // 接受数(Accept)
	ReCount         int     `json:"re_count"`                                   // 拒收数(Reject)
	CheckItems      string  `gorm:"type:text" json:"check_items"`               // 检验项目JSON
	Status          string  `gorm:"size:20;index" json:"status"`                // ACTIVE/INACTIVE
	Remark          string  `gorm:"size:500" json:"remark"`
	CreatedBy       string  `gorm:"size:50" json:"created_by"`
	UpdatedBy       string  `gorm:"size:50" json:"updated_by"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (QualityInspectionPlan) TableName() string {
	return "qc_inspection_plan"
}

// QualityInspectionPlanCreateRequest 创建检验计划请求
type QualityInspectionPlanCreateRequest struct {
	PlanCode       string  `json:"plan_code" binding:"required"`
	PlanName       string  `json:"plan_name" binding:"required"`
	InspectionType string  `json:"inspection_type" binding:"required"`
	AQLLevel       string  `json:"aql_level"`
	SampleSize     int     `json:"sample_size"`
	BatchMin       int     `json:"batch_min"`
	BatchMax       int     `json:"batch_max"`
	AcCount        int     `json:"ac_count"`
	ReCount        int     `json:"re_count"`
	CheckItems     string  `json:"check_items"`
	Status         string  `json:"status"`
	Remark         string  `json:"remark"`
}

// QualityInspectionPlanUpdateRequest 更新检验计划请求
type QualityInspectionPlanUpdateRequest struct {
	PlanName       string `json:"plan_name"`
	AQLLevel       string `json:"aql_level"`
	SampleSize     int    `json:"sample_size"`
	BatchMin       int    `json:"batch_min"`
	BatchMax       int    `json:"batch_max"`
	AcCount        int    `json:"ac_count"`
	ReCount        int    `json:"re_count"`
	CheckItems     string `json:"check_items"`
	Status         string `json:"status"`
	Remark         string `json:"remark"`
}

// QualityInspectionPlanQuery 查询参数
type QualityInspectionPlanQuery struct {
	PlanCode       string `form:"plan_code"`
	PlanName       string `form:"plan_name"`
	InspectionType string `form:"inspection_type"`
	AQLLevel       string `form:"aql_level"`
	Status         string `form:"status"`
}

// AQLSampleSize AQL标准抽样表（国标GB/T 2828.1）
type AQLSampleSize struct {
	ID               uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID         uint64    `gorm:"index;not null;default:1" json:"tenant_id"`
	SampleSizeCode   string    `gorm:"size:20" json:"sample_size_code"`   // 样本量字码
	BatchSizeMin     int       `json:"batch_size_min"`                   // 批量范围起点
	BatchSizeMax     int       `json:"batch_size_max"`                   // 批量范围终点
	InspectionLevel  string    `gorm:"size:10" json:"inspection_level"`  // 检验水平I/II/III/S-1/S-2/S-3
	AQLValue        float64   `json:"aql_value"`                        // AQL值
	SampleSize      int       `json:"sample_size"`                      // 样本量
	Ac1             int       `json:"ac_1"`                             // 接收数Ac
	Re1             int       `json:"re_1"`                             // 拒收数Re
	CreatedAt       time.Time `json:"created_at"`
}

func (AQLSampleSize) TableName() string {
	return "qc_aql_sample_size"
}
