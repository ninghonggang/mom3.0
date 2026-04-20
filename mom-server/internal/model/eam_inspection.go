package model

import (
	"time"
)

// ========== EAM巡检模块 ==========

// EAMInspectionPlan 巡检计划
type EAMInspectionPlan struct {
	BaseModel
	TenantID           int64      `json:"tenant_id" gorm:"index;not null"`
	PlanNo             string     `json:"plan_no" gorm:"size:50;not null"`
	PlanName           string     `json:"plan_name" gorm:"size:100;not null"`
	EquipmentCategory   string     `json:"equipment_category" gorm:"size:50"` // 设备分类
	CycleType          string     `json:"cycle_type" gorm:"size:20"`         // DAILY/WEEKLY/MONTHLY/QUARTERLY/YEARLY
	CycleDays          int        `json:"cycle_days" gorm:"default:1"`       // 周期天数
	StartDate          time.Time  `json:"start_date"`
	EndDate            *time.Time `json:"end_date"`
	Status             string     `json:"status" gorm:"size:20;default:'ACTIVE'"`
	Remark             string     `json:"remark" gorm:"type:text"`
}

func (EAMInspectionPlan) TableName() string {
	return "eam_inspection_plan"
}

// EAMInspectionItem 巡检项目
type EAMInspectionItem struct {
	BaseModel
	TenantID       int64  `json:"tenant_id" gorm:"index;not null"`
	PlanID         int64  `json:"plan_id" gorm:"index;not null"`
	ItemCode       string `json:"item_code" gorm:"size:50;not null"`
	ItemName       string `json:"item_name" gorm:"size:100;not null"`
	CheckMethod    string `json:"check_method" gorm:"size:200"`    // 检查方法
	CheckStandard  string `json:"check_standard" gorm:"size:200"`  // 检查标准
	IsRequired     bool   `json:"is_required" gorm:"default:true"`
	SortOrder      int    `json:"sort_order" gorm:"default:0"`
}

func (EAMInspectionItem) TableName() string {
	return "eam_inspection_item"
}

// EAMInspectionScheme 巡检方案(每次执行的记录)
type EAMInspectionScheme struct {
	BaseModel
	TenantID       int64      `json:"tenant_id" gorm:"index;not null"`
	PlanID         int64      `json:"plan_id" gorm:"index;not null"`
	PlanNo         string     `json:"plan_no" gorm:"size:50;not null"`
	SchemeNo       string     `json:"scheme_no" gorm:"size:50;not null"`   // 执行单号
	EquipmentID    int64      `json:"equipment_id" gorm:"index;not null"`
	EquipmentCode  string     `json:"equipment_code" gorm:"size:50"`
	EquipmentName  string     `json:"equipment_name" gorm:"size:100"`
	InspectorID    *int64     `json:"inspector_id"`
	InspectorName  string     `json:"inspector_name" gorm:"size:50"`
	InspectTime    *time.Time `json:"inspect_time"`
	Status         string     `json:"status" gorm:"size:20;default:'PENDING'"` // PENDING/IN_PROGRESS/COMPLETED/ABNORMAL
	Result         string     `json:"result" gorm:"size:20"`                   // OK/NG
	Remark         string     `json:"remark" gorm:"type:text"`
}

func (EAMInspectionScheme) TableName() string {
	return "eam_inspection_scheme"
}

// EAMInspectionResult 巡检结果明细
type EAMInspectionResult struct {
	BaseModel
	TenantID   int64  `json:"tenant_id" gorm:"index;not null"`
	SchemeID   int64  `json:"scheme_id" gorm:"index;not null"`
	ItemID     int64  `json:"item_id" gorm:"index;not null"`
	ItemName   string `json:"item_name" gorm:"size:100"`
	CheckValue string `json:"check_value" gorm:"size:200"` // 检查值
	IsNormal   bool   `json:"is_normal" gorm:"default:true"` // 是否正常
	Remark     string `json:"remark" gorm:"type:text"`
}

func (EAMInspectionResult) TableName() string {
	return "eam_inspection_result"
}

// ========== Request/Response VO ==========

// EAMInspectionPlanCreateReqVO 创建巡检计划请求
type EAMInspectionPlanCreateReqVO struct {
	PlanNo           string                     `json:"planNo"`
	PlanName         string                     `json:"planName"`
	EquipmentCategory string                     `json:"equipmentCategory"`
	CycleType        string                     `json:"cycleType"`
	CycleDays        int                        `json:"cycleDays"`
	StartDate        string                     `json:"startDate"`
	Remark           string                     `json:"remark"`
	Items            []EAMInspectionItemReqVO   `json:"items"`
}

// EAMInspectionItemReqVO 巡检项目请求
type EAMInspectionItemReqVO struct {
	ItemCode      string `json:"itemCode"`
	ItemName      string `json:"itemName"`
	CheckMethod   string `json:"checkMethod"`
	CheckStandard string `json:"checkStandard"`
	IsRequired    bool   `json:"isRequired"`
	SortOrder     int    `json:"sortOrder"`
}

// EAMInspectionSchemeCreateReqVO 创建巡检方案请求
type EAMInspectionSchemeCreateReqVO struct {
	PlanID       int64  `json:"planId"`
	EquipmentID  int64  `json:"equipmentId"`
	EquipmentCode string `json:"equipmentCode"`
	EquipmentName string `json:"equipmentName"`
	InspectorID  int64  `json:"inspectorId"`
	InspectorName string `json:"inspectorName"`
}

// EAMInspectionResultSubmitReqVO 提交巡检结果请求
type EAMInspectionResultSubmitReqVO struct {
	SchemeID int64                             `json:"schemeId"`
	Results  []EAMInspectionResultItemReqVO    `json:"results"`
}

// EAMInspectionResultItemReqVO 巡检结果明细项
type EAMInspectionResultItemReqVO struct {
	ItemID     int64  `json:"itemId"`
	ItemName   string `json:"itemName"`
	CheckValue string `json:"checkValue"`
	IsNormal   bool   `json:"isNormal"`
	Remark     string `json:"remark"`
}

// EAMInspectionPlanListReqVO 巡检计划列表查询
type EAMInspectionPlanListReqVO struct {
	PlanNo    string `json:"planNo"`
	PlanName  string `json:"planName"`
	Status    string `json:"status"`
	Page      int    `json:"page"`
	PageSize  int    `json:"pageSize"`
}

// EAMInspectionSchemeListReqVO 巡检执行列表查询
type EAMInspectionSchemeListReqVO struct {
	PlanNo    string `json:"planNo"`
	Status    string `json:"status"`
	Page      int    `json:"page"`
	PageSize  int    `json:"pageSize"`
}