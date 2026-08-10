package model

import (
	"time"

	"gorm.io/gorm"
)

// ============ 设备台账 ============

// EquipmentType 设备类型
type EquipmentType string

const (
	EquipmentTypeMachine    EquipmentType = "MACHINE"
	EquipmentTypeMold       EquipmentType = "MOLD"
	EquipmentTypeInstrument EquipmentType = "INSTRUMENT"
)

// EquipmentStatus 设备状态
type EquipmentStatus string

const (
	EquipmentStatusRunning     EquipmentStatus = "RUNNING"
	EquipmentStatusIdle        EquipmentStatus = "IDLE"
	EquipmentStatusMaintenance EquipmentStatus = "MAINTENANCE"
	EquipmentStatusRepair      EquipmentStatus = "REPAIR"
	EquipmentStatusScrapped    EquipmentStatus = "SCRAPPED"
)

// Equipment 设备台账
type Equipment struct {
	ID              int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID        int64          `gorm:"index;not null" json:"tenant_id"`
	EquipmentCode   string         `gorm:"uniqueIndex;size:64;not null" json:"equipment_code"`
	EquipmentName   string         `gorm:"size:128;not null" json:"equipment_name"`
	EquipmentType   EquipmentType  `gorm:"size:32;not null" json:"equipment_type"`
	EquipmentClassID *int64        `gorm:"index" json:"equipment_class_id,omitempty"`
	Model           string         `gorm:"size:128" json:"model"`
	Specification   string         `gorm:"size:256" json:"specification"`
	ManufacturerID  *int64         `gorm:"index" json:"manufacturer_id,omitempty"`
	SupplierID      *int64         `gorm:"index" json:"supplier_id,omitempty"`
	SerialNo        string         `gorm:"size:128" json:"serial_no"`
	WorkshopID      *int64         `gorm:"index" json:"workshop_id,omitempty"`
	LineID          *int64         `gorm:"index" json:"line_id,omitempty"`
	LocationID      *int64         `gorm:"index" json:"location_id,omitempty"`
	Status          EquipmentStatus `gorm:"size:32;index;default:IDLE" json:"status"`
	TargetOee       string         `gorm:"size:16" json:"target_oee"`
	PurchaseDate    *time.Time     `json:"purchase_date,omitempty"`
	PurchaseAmount  string         `gorm:"size:16" json:"purchase_amount"`
	ServiceLife     int32          `json:"service_life"`
	InstallDate     *time.Time     `json:"install_date,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Equipment) TableName() string { return "eam_equipment" }

// ============ 维修单 ============

// Urgency 紧急程度
type Urgency string

const (
	UrgencyNormal    Urgency = "NORMAL"
	UrgencyUrgent    Urgency = "URGENT"
	UrgencyEmergency Urgency = "EMERGENCY"
)

// RepairOrderStatus 维修单状态
type RepairOrderStatus string

const (
	RepairStatusReported    RepairOrderStatus = "REPORTED"
	RepairStatusAssigned    RepairOrderStatus = "ASSIGNED"
	RepairStatusInProgress  RepairOrderStatus = "IN_PROGRESS"
	RepairStatusPendingParts RepairOrderStatus = "PENDING_PARTS"
	RepairStatusCompleted   RepairOrderStatus = "COMPLETED"
	RepairStatusVerified    RepairOrderStatus = "VERIFIED"
	RepairStatusCancelled   RepairOrderStatus = "CANCELLED"
)

// RepairOrder 维修单
type RepairOrder struct {
	ID          int64             `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID    int64             `gorm:"index;not null" json:"tenant_id"`
	RepairNo    string            `gorm:"uniqueIndex;size:64;not null" json:"repair_no"`
	EquipmentID int64             `gorm:"index;not null" json:"equipment_id"`
	FaultType   string            `gorm:"size:64" json:"fault_type"`
	FaultDesc   string            `gorm:"size:512" json:"fault_desc"`
	Urgency     Urgency           `gorm:"size:32;default:NORMAL" json:"urgency"`
	ReporterID  *int64            `gorm:"index" json:"reporter_id,omitempty"`
	RepairmanID *int64            `gorm:"index" json:"repairman_id,omitempty"`
	Status      RepairOrderStatus `gorm:"size:32;index;default:REPORTED" json:"status"`
	ReportedAt  time.Time         `json:"reported_at"`
	CompletedAt *time.Time        `json:"completed_at,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	DeletedAt   gorm.DeletedAt    `gorm:"index" json:"deleted_at,omitempty"`
}

func (RepairOrder) TableName() string { return "eam_repair_order" }

// ============ 设备OEE ============

// EquipmentOee 设备OEE记录
type EquipmentOee struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	// (EquipmentID, CalcDate) 组合唯一：同一设备同一天只保留一条 OEE，重复上报走 upsert 覆盖。
	// 该唯一约束是 repository.oeeRepo.Create 中 ON CONFLICT 子句能生效的前提。
	EquipmentID  int64     `gorm:"uniqueIndex:uk_eam_oee_equipment_date;not null" json:"equipment_id"`
	CalcDate     string    `gorm:"size:10;uniqueIndex:uk_eam_oee_equipment_date" json:"calc_date"` // YYYY-MM-DD
	Availability float64   `gorm:"type:decimal(5,4)" json:"availability"`
	Performance  float64   `gorm:"type:decimal(5,4)" json:"performance"`
	Quality      float64   `gorm:"type:decimal(5,4)" json:"quality"`
	OEE          float64   `gorm:"type:decimal(5,4)" json:"oee"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (EquipmentOee) TableName() string { return "eam_equipment_oee" }

// ============ 保养计划 ============

// MaintenancePlanStatus 保养计划状态
type MaintenancePlanStatus string

const (
	MaintenanceStatusScheduled  MaintenancePlanStatus = "SCHEDULED"
	MaintenanceStatusInProgress MaintenancePlanStatus = "IN_PROGRESS"
	MaintenanceStatusCompleted  MaintenancePlanStatus = "COMPLETED"
	MaintenanceStatusSkipped    MaintenancePlanStatus = "SKIPPED"
)

// MaintenancePlan 保养计划
type MaintenancePlan struct {
	ID                 int64                 `gorm:"primaryKey;autoIncrement" json:"id"`
	EquipmentID        int64                 `gorm:"index;not null" json:"equipment_id"`
	PlanNo             string                `gorm:"uniqueIndex;size:64;not null" json:"plan_no"`
	MaintenanceType    string                `gorm:"size:64" json:"maintenance_type"`
	CycleDays          int32                 `json:"cycle_days"`
	NextMaintenanceDate string               `gorm:"size:10" json:"next_maintenance_date"` // YYYY-MM-DD
	Status             MaintenancePlanStatus `gorm:"size:32;index;default:SCHEDULED" json:"status"`
	CreatedAt          time.Time             `json:"created_at"`
	UpdatedAt          time.Time             `json:"updated_at"`
	DeletedAt          gorm.DeletedAt        `gorm:"index" json:"deleted_at,omitempty"`
}

func (MaintenancePlan) TableName() string { return "eam_maintenance_plan" }

// ============ 点检记录 ============

// CheckResult 点检结果
type CheckResult string

const (
	CheckResultOK CheckResult = "OK"
	CheckResultNG CheckResult = "NG"
	CheckResultNA CheckResult = "NA"
)

// EquipmentCheck 点检记录
type EquipmentCheck struct {
	ID         int64       `gorm:"primaryKey;autoIncrement" json:"id"`
	EquipmentID int64      `gorm:"index;not null" json:"equipment_id"`
	CheckNo    string      `gorm:"uniqueIndex;size:64;not null" json:"check_no"`
	CheckStdID *int64      `gorm:"index" json:"check_std_id,omitempty"`
	CheckerID  *int64      `gorm:"index" json:"checker_id,omitempty"`
	CheckTime  time.Time   `json:"check_time"`
	Result     CheckResult `gorm:"size:8;index" json:"result"`
	Remark     string      `gorm:"size:512" json:"remark"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (EquipmentCheck) TableName() string { return "eam_equipment_check" }

// ============ 设备停机 ============

// DowntimeType 停机类型
type DowntimeType string

const (
	DowntimeTypeUnplanned DowntimeType = "UNPLANNED"
	DowntimeTypePlanned   DowntimeType = "PLANNED"
)

// DowntimeStatus 停机状态
type DowntimeStatus string

const (
	DowntimeStatusActive   DowntimeStatus = "ACTIVE"
	DowntimeStatusResolved DowntimeStatus = "RESOLVED"
	DowntimeStatusPlanned  DowntimeStatus = "PLANNED"
)

// EquipmentDowntime 设备停机记录
type EquipmentDowntime struct {
	ID               int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	EquipmentID      int64          `gorm:"index;not null" json:"equipment_id"`
	DowntimeType     DowntimeType   `gorm:"size:16;not null" json:"downtime_type"`
	StartTime        time.Time      `json:"start_time"`
	EndTime          *time.Time     `json:"end_time,omitempty"`
	DurationSeconds  int32          `json:"duration_seconds"`
	Reason           string         `gorm:"size:512" json:"reason"`
	Status           DowntimeStatus `gorm:"size:16;index;default:ACTIVE" json:"status"`
	// 恢复时留痕：没有这两个字段，"谁修的、怎么修好的" 会随请求一起丢掉，
	// 停机记录也就失去了复盘价值。
	Resolution string `gorm:"size:512" json:"resolution"`
	ResolverID string `gorm:"size:64" json:"resolver_id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (EquipmentDowntime) TableName() string { return "eam_equipment_downtime" }

// ============ 通用分页 ============

// PageQuery 分页查询参数
type PageQuery struct {
	Page     int32
	PageSize int32
}

// PageResult 分页结果
type PageResult struct {
	Page       int32
	PageSize   int32
	Total      int64
	TotalPages int32
}

// NewPageQuery creates a PageQuery with defaults.
func NewPageQuery(page, pageSize int32) PageQuery {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	// 超限时截断到上限，而不是回落为默认值——后者会让 pageSize=500 静默变成 20。
	if pageSize > 200 {
		pageSize = 200
	}
	return PageQuery{Page: page, PageSize: pageSize}
}

// Offset returns the SQL offset for pagination.
func (p PageQuery) Offset() int {
	return int((p.Page - 1) * p.PageSize)
}

// CalcTotalPages computes total pages from total count.
func CalcTotalPages(total int64, pageSize int32) int32 {
	if pageSize <= 0 {
		return 0
	}
	pages := int32(total / int64(pageSize))
	if total%int64(pageSize) > 0 {
		pages++
	}
	return pages
}
