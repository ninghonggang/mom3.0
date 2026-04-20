package model

import (
	"time"
)

// InspectionFeature 检验特性表
type InspectionFeature struct {
	ID              uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	TenantID        uint64     `gorm:"index;not null;default:1" json:"tenant_id"`
	FeatureCode     string     `gorm:"size:50;uniqueIndex" json:"feature_code"`     // 特性编码
	FeatureName     string     `gorm:"size:100;not null" json:"feature_name"`       // 特性名称
	ProductID       uint64     `gorm:"index" json:"product_id"`                     // 产品ID
	ProductCode     string     `gorm:"size:50;index" json:"product_code"`           // 产品编码
	ProductName     string     `gorm:"size:100" json:"product_name"`                // 产品名称
	InspectionType  string     `gorm:"size:20;index" json:"inspection_type"`        // IQC/IPQC/FQC/OQC
	FeatureType     string     `gorm:"size:20" json:"feature_type"`                 // QUALITATIVE/QUANTITATIVE 定性/定量
	TechnicalSpec   string     `gorm:"size:500" json:"technical_spec"`              // 技术规格
	LowerLimit      *float64   `json:"lower_limit"`                                 // 下限
	UpperLimit      *float64   `json:"upper_limit"`                                 // 上限
	Unit            *string    `gorm:"size:20" json:"unit"`                          // 单位
	SampleSize      int        `json:"sample_size"`                                 // 样本大小
	GaugesMethod    string     `gorm:"size:100" json:"gauges_method"`                // 检测方法/量具
	AQLLevel        string     `gorm:"size:20" json:"aql_level"`                    // CRITICAL/MAJOR/MINOR
	Status          string     `gorm:"size:20;index" json:"status"`                 // ACTIVE/INACTIVE
	Remark          string     `gorm:"size:500" json:"remark"`
	CreatedBy       string     `gorm:"size:50" json:"created_by"`
	UpdatedBy       string     `gorm:"size:50" json:"updated_by"`
}

func (InspectionFeature) TableName() string {
	return "qc_inspection_feature"
}

// InspectionFeatureCreateRequest 创建检验特性请求
type InspectionFeatureCreateRequest struct {
	FeatureCode    string   `json:"feature_code" binding:"required"`
	FeatureName    string   `json:"feature_name" binding:"required"`
	ProductID      uint64   `json:"product_id"`
	ProductCode    string   `json:"product_code"`
	ProductName    string   `json:"product_name"`
	InspectionType string   `json:"inspection_type" binding:"required"`
	FeatureType    string   `json:"feature_type" binding:"required"`
	TechnicalSpec  string   `json:"technical_spec"`
	LowerLimit     *float64 `json:"lower_limit"`
	UpperLimit     *float64 `json:"upper_limit"`
	Unit           *string  `json:"unit"`
	SampleSize     int      `json:"sample_size"`
	GaugesMethod   string   `json:"gauges_method"`
	AQLLevel       string   `json:"aql_level"`
	Status         string   `json:"status"`
	Remark         string   `json:"remark"`
}

// InspectionFeatureUpdateRequest 更新检验特性请求
type InspectionFeatureUpdateRequest struct {
	FeatureName    string   `json:"feature_name"`
	ProductID      uint64   `json:"product_id"`
	ProductCode    string   `json:"product_code"`
	ProductName    string   `json:"product_name"`
	InspectionType string   `json:"inspection_type"`
	FeatureType    string   `json:"feature_type"`
	TechnicalSpec  string   `json:"technical_spec"`
	LowerLimit     *float64 `json:"lower_limit"`
	UpperLimit     *float64 `json:"upper_limit"`
	Unit           *string  `json:"unit"`
	SampleSize     int      `json:"sample_size"`
	GaugesMethod   string   `json:"gauges_method"`
	AQLLevel       string   `json:"aql_level"`
	Status         string   `json:"status"`
	Remark         string   `json:"remark"`
}

// InspectionFeatureQuery 查询参数
type InspectionFeatureQuery struct {
	ProductID      uint64 `form:"product_id"`
	InspectionType string `form:"inspection_type"`
	Status         string `form:"status"`
	FeatureType    string `form:"feature_type"`
}

// InspectionFeatureBatchCreateRequest 批量创建请求
type InspectionFeatureBatchCreateRequest struct {
	Features []InspectionFeatureCreateRequest `json:"features" binding:"required"`
}