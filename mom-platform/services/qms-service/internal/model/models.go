package model

import (
	"time"

	"gorm.io/gorm"
)

// InspectionSheet status constants.
const (
	SheetStatusPending    = "PENDING"
	SheetStatusInProgress = "IN_PROGRESS"
	SheetStatusPassed     = "PASSED"
	SheetStatusFailed     = "FAILED"
	SheetStatusWaived     = "WAIVED"
	SheetStatusCancelled  = "CANCELLED"
)

// NCR status constants.
const (
	NcrStatusOpen          = "OPEN"
	NcrStatusInvestigating = "INVESTIGATING"
	NcrStatusDispositioned = "DISPOSITIONED"
	NcrStatusVerified      = "VERIFIED"
	NcrStatusClosed        = "CLOSED"
	NcrStatusCancelled     = "CANCELLED"
	NcrStatusReopened      = "REOPENED"
)

// InspectionSheet represents a quality inspection sheet (检验单).
type InspectionSheet struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	TenantID        string         `gorm:"size:64;index;not null" json:"tenant_id"`
	SheetNo         string         `gorm:"size:64;uniqueIndex;not null" json:"sheet_no"`
	InspectionType  string         `gorm:"size:32;not null" json:"inspection_type"`
	MaterialID      string         `gorm:"size:64;index" json:"material_id"`
	BatchID         string         `gorm:"size:64;index" json:"batch_id"`
	SampleSize      int            `gorm:"not null" json:"sample_size"`
	DefectCount     int            `gorm:"default:0" json:"defect_count"`
	Status          string         `gorm:"size:32;index;default:PENDING" json:"status"`
	InspectorID     string         `gorm:"size:64;index" json:"inspector_id"`
	InspectedAt     *time.Time     `json:"inspected_at"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName overrides the table name.
func (InspectionSheet) TableName() string { return "qms_inspection_sheets" }

// InspectionCharacteristic represents a measurable inspection characteristic (检验特性).
type InspectionCharacteristic struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	CharCode  string  `gorm:"size:64;uniqueIndex;not null" json:"char_code"`
	CharName  string  `gorm:"size:256;not null" json:"char_name"`
	DataType  string  `gorm:"size:32;not null" json:"data_type"` // NUMERIC, BOOLEAN, TEXT
	USL       float64 `gorm:"comment:upper_spec_limit" json:"usl"`
	LSL       float64 `gorm:"comment:lower_spec_limit" json:"lsl"`
	Target    float64 `json:"target"`
	Unit      string  `gorm:"size:32" json:"unit"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (InspectionCharacteristic) TableName() string { return "qms_inspection_characteristics" }

// InspectionPlan represents an inspection plan/scheme (检验计划).
type InspectionPlan struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	SchemeCode   string `gorm:"size:64;uniqueIndex;not null" json:"scheme_code"`
	SchemeName   string `gorm:"size:256;not null" json:"scheme_name"`
	SchemeType   string `gorm:"size:32;not null" json:"scheme_type"`
	TemplateID   string `gorm:"size:64;index" json:"template_id"`
	Status       string `gorm:"size:32;default:ACTIVE" json:"status"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (InspectionPlan) TableName() string { return "qms_inspection_plans" }

// InspectionResult represents a single measured inspection result (检验结果).
type InspectionResult struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	SheetID   uint    `gorm:"index;not null" json:"sheet_id"`
	CharID    uint    `gorm:"index;not null" json:"char_id"`
	Value     string  `gorm:"type:text" json:"value"`
	Pass      bool    `json:"pass"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (InspectionResult) TableName() string { return "qms_inspection_results" }

// Ncr represents a non-conformance report / defect disposition (不良品处置单).
type Ncr struct {
	ID                uint       `gorm:"primaryKey" json:"id"`
	TenantID          string     `gorm:"size:64;index;not null" json:"tenant_id"`
	NcrNo             string     `gorm:"size:64;uniqueIndex;not null" json:"ncr_no"`
	InspectionSheetID uint       `gorm:"index;not null" json:"inspection_sheet_id"`
	MaterialID        string     `gorm:"size:64;index" json:"material_id"`
	BatchID           string     `gorm:"size:64;index" json:"batch_id"`
	Quantity          float64    `gorm:"not null" json:"quantity"`
	Severity          string     `gorm:"size:32" json:"severity"`
	Status            string     `gorm:"size:32;index;default:OPEN" json:"status"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

func (Ncr) TableName() string { return "qms_ncrs" }

// NcrAction represents a corrective/disposition action (处置措施).
type NcrAction struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	NcrID         uint   `gorm:"index;not null" json:"ncr_id"`
	ActionType    string `gorm:"size:32;not null" json:"action_type"`
	ActionDesc    string `gorm:"type:text" json:"action_desc"`
	ResponsibleID string `gorm:"size:64" json:"responsible_id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (NcrAction) TableName() string { return "qms_ncr_actions" }

// DefectCode represents a defect classification code (缺陷代码).
type DefectCode struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	DefectCode  string `gorm:"size:64;uniqueIndex;not null" json:"defect_code"`
	DefectName  string `gorm:"size:256;not null" json:"defect_name"`
	DefectClass string `gorm:"size:32" json:"defect_class"`
	Severity    string `gorm:"size:32" json:"severity"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (DefectCode) TableName() string { return "qms_defect_codes" }

// SpcData represents a single SPC sample data point (SPC数据).
type SpcData struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CharID      uint      `gorm:"index;not null" json:"char_id"`
	SampleValue float64   `gorm:"not null" json:"sample_value"`
	SampleTime  time.Time `gorm:"index;not null" json:"sample_time"`
	Xbar        float64   `json:"xbar"`
	RValue      float64   `json:"r_value"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (SpcData) TableName() string { return "qms_spc_data" }
