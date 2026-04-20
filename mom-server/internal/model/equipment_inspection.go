package model

import (
	"time"
)

// ========== 设备点检模块 ==========

// InspectionTemplate 点检标准模板
type InspectionTemplate struct {
	BaseModel
	TenantID            int64      `json:"tenant_id" gorm:"index;not null"`
	TemplateCode        string     `json:"template_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_template"`
	TemplateName        string     `json:"template_name" gorm:"size:200;not null"`
	TemplateType        string     `json:"template_type" gorm:"size:20"` // DAILY/WEEKLY/MONTHLY/SPECIAL
	EquipmentTypeID     *int64     `json:"equipment_type_id"`
	Version             string     `json:"version" gorm:"size:20;default:1.0"`
	FrequencyType       string     `json:"frequency_type" gorm:"size:20"` // DAILY/WEEKLY/MONTHLY/ONCE
	FrequencyValue      int        `json:"frequency_value" gorm:"default:1"`
	ExecutionTime       string     `json:"execution_time" gorm:"size:50"`
	EstimatedMinutes    int        `json:"estimated_minutes" gorm:"default:30"`
	RequiredSignature   int        `json:"required_signature" gorm:"default:1"`
	RequirePhoto        int        `json:"require_photo" gorm:"default:0"`
	RequireDefectReport int        `json:"require_defect_report" gorm:"default:1"`
	IsActive            int        `json:"is_active" gorm:"default:1"`
	EffectiveDate       *time.Time `json:"effective_date"`
	ExpiryDate          *time.Time `json:"expiry_date"`
	Remark              *string    `json:"remark" gorm:"type:text"`
}

func (InspectionTemplate) TableName() string {
	return "eam_inspection_template"
}

// InspectionItem 点检项目定义
type InspectionItem struct {
	BaseModel
	TemplateID      int64   `json:"template_id" gorm:"index;not null"`
	ItemCode        string  `json:"item_code" gorm:"size:50;not null"`
	ItemName        string  `json:"item_name" gorm:"size:200;not null"`
	ItemCategory    string  `json:"item_category" gorm:"size:50"` // 机械/电气/安全/精度
	CheckMethod     string  `json:"check_method" gorm:"size:100"`
	CheckStandard   string  `json:"check_standard" gorm:"size:200"`
	Unit            string  `json:"unit" gorm:"size:20"`
	UpperLimit      float64 `json:"upper_limit" gorm:"type:decimal(18,6)"`
	LowerLimit      float64 `json:"lower_limit" gorm:"type:decimal(18,6)"`
	StandardValue   float64 `json:"standard_value" gorm:"type:decimal(18,6)"`
	CheckType       string  `json:"check_type" gorm:"size:20"` // VISUAL/MEASURE/TAP/AUDIO/OTHER
	IsCriticalPoint int     `json:"is_critical_point" gorm:"default:0"`
	IsMandatory     int     `json:"is_mandatory" gorm:"default:1"`
	ResultType      string  `json:"result_type" gorm:"size:20"` // OK_NG/NUMERIC/TEXT
	SortOrder       int     `json:"sort_order" gorm:"default:0"`
	Remark          *string `json:"remark" gorm:"type:text"`
}

func (InspectionItem) TableName() string {
	return "eam_inspection_item"
}

// InspectionPlan 点检计划
type InspectionPlan struct {
	BaseModel
	TenantID           int64      `json:"tenant_id" gorm:"index;not null"`
	PlanNo             string     `json:"plan_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_plan"`
	PlanName           string     `json:"plan_name" gorm:"size:200;not null"`
	TemplateID         int64      `json:"template_id"`
	TemplateName       string     `json:"template_name" gorm:"size:200"`
	EquipmentID        int64      `json:"equipment_id"`
	EquipmentCode      string     `json:"equipment_code" gorm:"size:50"`
	EquipmentName      string     `json:"equipment_name" gorm:"size:100"`
	WorkstationID      *int64     `json:"workstation_id"`
	WorkstationName    *string    `json:"workstation_name" gorm:"size:100"`
	PlanDate           time.Time  `json:"plan_date"`
	PlanShift          string     `json:"plan_shift" gorm:"size:20"`
	PlanStartTime      string     `json:"plan_start_time" gorm:"size:10"`
	PlanEndTime        string     `json:"plan_end_time" gorm:"size:10"`
	AssignedTo         *int64     `json:"assigned_to"`
	AssignedName       *string    `json:"assigned_name" gorm:"size:50"`
	Status             string     `json:"status" gorm:"size:20;default:PENDING"` // PENDING/ASSIGNED/EXECUTING/COMPLETED/CANCELLED
	IsGenerated        int        `json:"is_generated" gorm:"default:0"`
	WorkshopID         *int64     `json:"workshop_id"`
}

func (InspectionPlan) TableName() string {
	return "eam_inspection_plan"
}

// InspectionRecord 点检记录
type InspectionRecord struct {
	BaseModel
	RecordNo            string     `json:"record_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_record"`
	PlanID              *int64     `json:"plan_id"`
	TemplateID          int64      `json:"template_id"`
	TemplateName        string     `json:"template_name" gorm:"size:200"`
	EquipmentID         int64      `json:"equipment_id"`
	EquipmentCode       string     `json:"equipment_code" gorm:"size:50"`
	EquipmentName       string     `json:"equipment_name" gorm:"size:100"`
	InspectorID         int64      `json:"inspector_id"`
	InspectorName       string     `json:"inspector_name" gorm:"size:50"`
	InspectionStartTime time.Time  `json:"inspection_start_time"`
	InspectionEndTime   *time.Time `json:"inspection_end_time"`
	ActualDuration      int        `json:"actual_duration"` // 实际耗时(分钟)
	OverallResult       string     `json:"overall_result" gorm:"size:20"` // OK/NG/PARTIAL
	OKCount             int        `json:"ok_count" gorm:"default:0"`
	NGCount             int        `json:"ng_count" gorm:"default:0"`
	Shift               string     `json:"shift" gorm:"size:20"`
	ShiftName           *string    `json:"shift_name" gorm:"size:50"`
	LocationLat         *string    `json:"location_lat" gorm:"size:20"`
	LocationLng         *string    `json:"location_lng" gorm:"size:20"`
	SignatureData       *string    `json:"signature_data" gorm:"type:text"`
	SignatureTime       *time.Time `json:"signature_time"`
	Status              string     `json:"status" gorm:"size:20;default:EXECUTING"` // EXECUTING/COMPLETED/ABNORMAL
	Remark              *string    `json:"remark" gorm:"type:text"`
	WorkshopID          *int64     `json:"workshop_id"`
}

func (InspectionRecord) TableName() string {
	return "eam_inspection_record"
}

// InspectionResult 点检结果明细
type InspectionResult struct {
	BaseModel
	RecordID             int64     `json:"record_id" gorm:"index;not null"`
	ItemID               int64     `json:"item_id"`
	ItemCode             string    `json:"item_code" gorm:"size:50"`
	ItemName             string    `json:"item_name" gorm:"size:200"`
	ItemCategory         string    `json:"item_category" gorm:"size:50"`
	ResultType           string    `json:"result_type" gorm:"size:20"` // OK/NG/NUMERIC/TEXT
	ResultValue          string    `json:"result_value" gorm:"size:200"`
	ResultStatus         string    `json:"result_status" gorm:"size:20;not null"` // OK/NG/N/A
	IsAbnormal           int       `json:"is_abnormal" gorm:"default:0"`
	AbnormalDescription   *string   `json:"abnormal_description" gorm:"type:text"`
	AbnormalPhotos        *string   `json:"abnormal_photos" gorm:"type:text"` // JSON array
	DefectCode           *string   `json:"defect_code" gorm:"size:50"`
	DefectName           *string   `json:"defect_name" gorm:"size:100"`
	HandlingAction       *string   `json:"handling_action" gorm:"size:100"`
	HandledBy            *int64    `json:"handled_by"`
	HandledByName        *string   `json:"handled_by_name" gorm:"size:50"`
	HandledTime          *time.Time `json:"handled_time"`
	Remark               *string   `json:"remark" gorm:"type:text"`
}

func (InspectionResult) TableName() string {
	return "eam_inspection_result"
}

// InspectionDefect 点检异常处理
type InspectionDefect struct {
	BaseModel
	DefectNo            string     `json:"defect_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_defect_no"`
	RecordID            int64      `json:"record_id" gorm:"index"`
	ResultID            int64      `json:"result_id"`
	DefectType          string     `json:"defect_type" gorm:"size:20"` // ABNORMAL/FAIL/POTENTIAL
	DefectCode          *string    `json:"defect_code" gorm:"size:50"`
	DefectName          *string    `json:"defect_name" gorm:"size:200"`
	DefectLevel         string     `json:"defect_level" gorm:"size:10"` // A/B/C
	Description         string     `json:"description" gorm:"type:text"`
	Photos              *string    `json:"photos" gorm:"type:text"` // JSON array
	Status              string     `json:"status" gorm:"size:20;default:REPORTED"` // REPORTED/ACKNOWLEDGED/PROCESSING/RESOLVED/CLOSED
	ReportedBy          *int64     `json:"reported_by"`
	ReportedTime        *time.Time `json:"reported_time"`
	AssignedTo          *int64     `json:"assigned_to"`
	AssignedName        *string    `json:"assigned_name" gorm:"size:50"`
	AssignmentTime      *time.Time `json:"assignment_time"`
	Resolution          *string    `json:"resolution" gorm:"size:200"`
	ResolvedBy          *int64     `json:"resolved_by"`
	ResolvedTime        *time.Time `json:"resolved_time"`
	ResolutionPhotos    *string    `json:"resolution_photos" gorm:"type:text"` // JSON array
	CreateRepairOrder   int        `json:"create_repair_order" gorm:"default:0"`
	RepairOrderID      *int64     `json:"repair_order_id"`
	Remark              *string    `json:"remark" gorm:"type:text"`
}

func (InspectionDefect) TableName() string {
	return "eam_inspection_defect"
}

// InspectionStatistics 点检完成率统计
type InspectionStatistics struct {
	BaseModel
	StatDate         time.Time `json:"stat_date"`
	WorkshopID       *int64    `json:"workshop_id"`
	TemplateType     *string   `json:"template_type" gorm:"size:20"`
	TotalPlans       int       `json:"total_plans" gorm:"default:0"`
	CompletedPlans   int       `json:"completed_plans" gorm:"default:0"`
	CancelledPlans   int       `json:"cancelled_plans" gorm:"default:0"`
	OverduePlans     int       `json:"overdue_plans" gorm:"default:0"`
	CompletionRate   float64   `json:"completion_rate" gorm:"type:decimal(5,2)"`
	OnTimeRate       float64   `json:"on_time_rate" gorm:"type:decimal(5,2)"`
	AbnormalityRate  float64   `json:"abnormality_rate" gorm:"type:decimal(5,2)"`
	TotalItems       int       `json:"total_items" gorm:"default:0"`
	OKItems          int       `json:"ok_items" gorm:"default:0"`
	NGItems          int       `json:"ng_items" gorm:"default:0"`
	NGRate           float64   `json:"ng_rate" gorm:"type:decimal(5,2)"`
}

func (InspectionStatistics) TableName() string {
	return "eam_inspection_statistics"
}
