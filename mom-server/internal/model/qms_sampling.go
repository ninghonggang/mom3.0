package model

import (
	"time"
)

// QMSSamplingPlan - 抽样方案
type QMSSamplingPlan struct {
	BaseModel
	TenantID         int64   `json:"tenant_id" gorm:"index;not null"`
	PlanCode         string  `json:"plan_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_plan_code"`
	PlanName         string  `json:"plan_name" gorm:"size:100;not null"`
	InspectionLevel  string  `json:"inspection_level" gorm:"size:20"` // General/Special/Skip
	SampleType       string  `json:"sample_type" gorm:"size:20"`   // NORMAL/DOUBLE/SKIP
	AQL              float64 `json:"aql" gorm:"type:decimal(5,2)"`  // 接收质量限
	Status           string  `json:"status" gorm:"size:20;default:'ACTIVE'"`
	Remark           string  `json:"remark" gorm:"type:text"`
}

func (QMSSamplingPlan) TableName() string {
	return "qms_sampling_plan"
}

// QMSSamplingRule - 抽样规则
type QMSSamplingRule struct {
	BaseModel
	PlanID       int64   `json:"plan_id" gorm:"index;not null"`
	BatchQtyFrom float64 `json:"batch_qty_from" gorm:"type:decimal(18,3)"` // 批量范围起
	BatchQtyTo   float64 `json:"batch_qty_to" gorm:"type:decimal(18,3)"`     // 批量范围止
	SampleSize   int     `json:"sample_size"`                                 // 样本数
	AcAccept     int     `json:"ac_accept"`                                  // 接收数
	ReReject     int     `json:"re_reject"`                                  // 拒收数
}

func (QMSSamplingRule) TableName() string {
	return "qms_sampling_rule"
}

// QMSSamplingRecord - 抽样记录
type QMSSamplingRecord struct {
	BaseModel
	TenantID     int64      `json:"tenant_id" gorm:"index;not null"`
	PlanID      int64      `json:"plan_id" gorm:"index;not null"`
	PlanCode    string     `json:"plan_code" gorm:"size:50"`
	InspectionID int64    `json:"inspection_id" gorm:"index;not null"` // 关联检验单
	BatchQty    float64    `json:"batch_qty" gorm:"type:decimal(18,3)"`
	SampleSize  int        `json:"sample_size"`
	DefectCount int        `json:"defect_count"`
	AcResult    bool       `json:"ac_result" gorm:"default:false"` // true=Accept, false=Reject
	Inspector   string     `json:"inspector" gorm:"size:50"`
	InspectTime *time.Time `json:"inspect_time"`
}

func (QMSSamplingRecord) TableName() string {
	return "qms_sampling_record"
}

// ========== Request/Response VO ==========

// QMSSamplingPlanCreateReqVO - 创建抽样方案请求
type QMSSamplingPlanCreateReqVO struct {
	PlanCode         string                   `json:"planCode" binding:"required"`
	PlanName         string                   `json:"planName" binding:"required"`
	InspectionLevel  string                   `json:"inspectionLevel"`
	SampleType       string                   `json:"sampleType"`
	AQL              float64                  `json:"aql"`
	Remark           string                   `json:"remark"`
	Rules            []QMSSamplingRuleReqVO   `json:"rules"`
}

// QMSSamplingPlanUpdateReqVO - 更新抽样方案请求
type QMSSamplingPlanUpdateReqVO struct {
	PlanName        string                 `json:"planName"`
	InspectionLevel string                 `json:"inspectionLevel"`
	SampleType      string                 `json:"sampleType"`
	AQL             float64                `json:"aql"`
	Status         string                 `json:"status"`
	Remark         string                 `json:"remark"`
}

// QMSSamplingRuleReqVO - 抽样规则请求
type QMSSamplingRuleReqVO struct {
	BatchQtyFrom float64 `json:"batchQtyFrom"`
	BatchQtyTo   float64 `json:"batchQtyTo"`
	SampleSize   int     `json:"sampleSize"`
	AcAccept     int     `json:"acAccept"`
	ReReject     int     `json:"reReject"`
}

// QMSSamplingRulesUpdateReqVO - 更新抽样规则请求
type QMSSamplingRulesUpdateReqVO struct {
	Rules []QMSSamplingRuleReqVO `json:"rules" binding:"required"`
}

// QMSSamplingCalculateReqVO - 计算样本数请求
type QMSSamplingCalculateReqVO struct {
	PlanID   int64   `json:"planId" binding:"required"`
	BatchQty float64 `json:"batchQty" binding:"required"`
}

// QMSSamplingCalculateRespVO - 计算样本数响应
type QMSSamplingCalculateRespVO struct {
	SampleSize int  `json:"sampleSize"`
	AcAccept   int  `json:"acAccept"`
	ReReject   int  `json:"reject"`
}

// QMSSamplingRecordCreateReqVO - 创建抽样记录请求
type QMSSamplingRecordCreateReqVO struct {
	PlanID       int64  `json:"planId" binding:"required"`
	PlanCode     string `json:"planCode"`
	InspectionID int64  `json:"inspectionId" binding:"required"`
	BatchQty     float64 `json:"batchQty"`
	SampleSize   int    `json:"sampleSize"`
	DefectCount  int    `json:"defectCount"`
	AcResult     bool   `json:"acResult"`
	Inspector    string `json:"inspector"`
}

// QMSSamplingPlanDetailRespVO - 抽样方案详情响应
type QMSSamplingPlanDetailRespVO struct {
	QMSSamplingPlan
	Rules []QMSSamplingRule `json:"rules"`
}
