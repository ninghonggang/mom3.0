package model

import (
	"time"
)

// ========== QRCI质量闭环模块 ==========

// QRCI QRCI头表
type QRCI struct {
	BaseModel
	TenantID             int64      `json:"tenant_id" gorm:"index;not null"`
	QRICNo               string     `json:"qrci_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_qrci"`
	SourceType           string     `json:"source_type" gorm:"size:20"` // NCR/COMPLAINT/AUDIT/INTERNAL
	SourceID             *int64     `json:"source_id"`
	DefectDescription    string     `json:"defect_description" gorm:"type:text"`
	SeverityLevel       string     `json:"severity_level" gorm:"size:10"` // CRITICAL/HIGH/MEDIUM/LOW
	DiscoveryDate        time.Time  `json:"discovery_date"`
	DiscoveryLocation   *string    `json:"discovery_location" gorm:"size:100"`
	ResponsibleDeptID   *int64     `json:"responsible_dept_id"`
	ResponsibleDeptName *string    `json:"responsible_dept_name" gorm:"size:100"`
	OwnerID              *int64     `json:"owner_id"`
	OwnerName            *string    `json:"owner_name" gorm:"size:50"`
	TargetCloseDate     *time.Time `json:"target_close_date"`
	ActualCloseDate     *time.Time `json:"actual_close_date"`
	Status               string     `json:"status" gorm:"size:20;default:OPEN"` // OPEN/IN_PROGRESS/VERIFICATION/CLOSED/CANCELLED
	VerificationResult  *string    `json:"verification_result" gorm:"size:20"` // EFFECTIVE/INEFFECTIVE/FOLLOW_UP
	VerificationBy      *int64     `json:"verification_by"`
	VerificationTime    *time.Time `json:"verification_time"`
	Remark              *string    `json:"remark" gorm:"size:500"`
}

func (QRCI) TableName() string {
	return "qms_qrci"
}

// QRCI5Why QRCI 5Why分析
type QRCI5Why struct {
	BaseModel
	QRCIID       int64  `json:"qrci_id" gorm:"index;not null"`
	WhyLevel     int    `json:"why_level"`
	Question     string `json:"question" gorm:"size:500"`
	Answer       string `json:"answer" gorm:"size:500"`
	IsRootCause  int    `json:"is_root_cause" gorm:"default:0"` // 0-否, 1-是
}

func (QRCI5Why) TableName() string {
	return "qms_qrci_5why"
}

// QRCIAction QRCI纠正措施
type QRCIAction struct {
	BaseModel
	QRCIID             int64      `json:"qrci_id" gorm:"index;not null"`
	ActionType        string     `json:"action_type" gorm:"size:20"` // CORRECTIVE/PREVENTIVE
	ActionDescription string     `json:"action_description" gorm:"size:500"`
	ResponsibleID      int64      `json:"responsible_id"`
	ResponsibleName   *string    `json:"responsible_name" gorm:"size:50"`
	DueDate           time.Time  `json:"due_date"`
	CompletedDate     *time.Time `json:"completed_date"`
	EvidenceURLs      *string    `json:"evidence_urls" gorm:"type:text"` // JSON array
	Status            string     `json:"status" gorm:"size:20;default:PENDING"` // PENDING/IN_PROGRESS/COMPLETED/VERIFIED
	VerificationResult *string   `json:"verification_result" gorm:"size:20"`
	VerificationBy    *int64     `json:"verification_by"`
	VerificationTime  *time.Time `json:"verification_time"`
	Remark            *string    `json:"remark" gorm:"size:500"`
}

func (QRCIAction) TableName() string {
	return "qms_qrci_action"
}

// QRCIVerification QRCI效果确认
type QRCIVerification struct {
	BaseModel
	QRCIID             int64      `json:"qrci_id" gorm:"index;not null"`
	VerificationDate   time.Time  `json:"verification_date"`
	VerifierID         int64      `json:"verifier_id"`
	VerifierName      *string    `json:"verifier_name" gorm:"size:50"`
	Effectiveness     string     `json:"effectiveness" gorm:"size:20"` // EFFECTIVE/INEFFECTIVE
	EvidenceDescription *string   `json:"evidence_description" gorm:"size:500"`
	EvidenceURLs       *string    `json:"evidence_urls" gorm:"type:text"` // JSON array
	FollowUpRequired  int        `json:"follow_up_required" gorm:"default:0"` // 0-否, 1-是
	FollowUpRemark    *string    `json:"follow_up_remark" gorm:"size:500"`
}

func (QRCIVerification) TableName() string {
	return "qms_qrci_verification"
}

// ========== LPA分层审核模块 ==========

// LPAStandard LPA审核标准
type LPAStandard struct {
	BaseModel
	StandardCode    string     `json:"standard_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_lpa_std"`
	StandardName    string     `json:"standard_name" gorm:"size:200;not null"`
	Version          string     `json:"version" gorm:"size:20;default:1.0"`
	DeptID          *int64     `json:"dept_id"`
	DeptName        *string    `json:"dept_name" gorm:"size:100"`
	AuditFrequency  string     `json:"audit_frequency" gorm:"size:20"` // DAILY/WEEKLY/MONTHLY/QUARTERLY
	AuditorLevels   *string    `json:"auditor_levels" gorm:"type:text"` // JSON array
	QuestionCount   int        `json:"question_count" gorm:"default:0"`
	PassingScore    float64    `json:"passing_score" gorm:"type:decimal(5,2)"`
	IsActive        int        `json:"is_active" gorm:"default:1"`
	EffectiveDate   *time.Time `json:"effective_date"`
	Remark          *string    `json:"remark" gorm:"size:500"`
}

func (LPAStandard) TableName() string {
	return "qms_lpa_standard"
}

// LPAQuestion LPA审核问题项
type LPAQuestion struct {
	BaseModel
	StandardID        int64   `json:"standard_id" gorm:"index;not null"`
	QuestionNo        string  `json:"question_no" gorm:"size:10;not null"`
	QuestionText      string  `json:"question_text" gorm:"size:500;not null"`
	AuditPoint        *string `json:"audit_point" gorm:"size:200"`
	Severity          string  `json:"severity" gorm:"size:10"` // CRITICAL/MAJOR/MINOR
	IsCriticalPoint   int     `json:"is_critical_point" gorm:"default:0"` // 0-否, 1-是
	SortOrder         int     `json:"sort_order" gorm:"default:0"`
}

func (LPAQuestion) TableName() string {
	return "qms_lpa_question"
}

// LPARecord LPA审核记录
type LPARecord struct {
	BaseModel
	RecordNo         string     `json:"record_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_lpa_record"`
	StandardID       int64      `json:"standard_id"`
	StandardName     string     `json:"standard_name" gorm:"size:200"`
	AuditorID        int64      `json:"auditor_id"`
	AuditorName      string     `json:"auditor_name" gorm:"size:50"`
	AuditorLevel     string     `json:"auditor_level" gorm:"size:20"` // L1/L2/L3/L4
	AuditDeptID      *int64     `json:"audit_dept_id"`
	AuditDeptName    *string    `json:"audit_dept_name" gorm:"size:100"`
	AuditLocation   *string    `json:"audit_location" gorm:"size:100"`
	AuditDate        time.Time  `json:"audit_date"`
	AuditTime        string     `json:"audit_time" gorm:"size:10"`
	Shift            string     `json:"shift" gorm:"size:20"` // DAY/EVENING/NIGHT
	TotalQuestions   int        `json:"total_questions" gorm:"default:0"`
	YesCount         int        `json:"yes_count" gorm:"default:0"`
	NoCount          int        `json:"no_count" gorm:"default:0"`
	NACount          int        `json:"na_count" gorm:"default:0"`
	Score            float64    `json:"score" gorm:"type:decimal(5,2)"`
	Result           string     `json:"result" gorm:"size:10"` // PASS/FAIL
	Findings         *string    `json:"findings" gorm:"type:text"` // JSON array
	NextAction       *string    `json:"next_action" gorm:"size:500"`
	Status           string     `json:"status" gorm:"size:20;default:SUBMITTED"` // SUBMITTED/VERIFIED
	VerifiedBy       *int64     `json:"verified_by"`
	VerifiedTime     *time.Time `json:"verified_time"`
}

func (LPARecord) TableName() string {
	return "qms_lpa_record"
}

// LPARecordItem LPA审核记录明细
type LPARecordItem struct {
	BaseModel
	RecordID      int64  `json:"record_id" gorm:"index;not null"`
	QuestionID    int64  `json:"question_id"`
	QuestionNo   string `json:"question_no" gorm:"size:10"`
	QuestionText string `json:"question_text" gorm:"size:500"`
	Result       string `json:"result" gorm:"size:10"` // YES/NO/N/A
	Finding      *string `json:"finding" gorm:"type:text"`
	Remark       *string `json:"remark" gorm:"size:500"`
}

func (LPARecordItem) TableName() string {
	return "qms_lpa_record_item"
}

// LPAStatistics LPA统计分析
type LPAStatistics struct {
	BaseModel
	StatDate       time.Time `json:"stat_date"`
	DeptID        *int64    `json:"dept_id"`
	DeptName      *string   `json:"dept_name" gorm:"size:100"`
	StandardID    *int64    `json:"standard_id"`
	TotalRecords  int       `json:"total_records" gorm:"default:0"`
	PassRecords   int       `json:"pass_records" gorm:"default:0"`
	FailRecords   int       `json:"fail_records" gorm:"default:0"`
	PassRate      float64   `json:"pass_rate" gorm:"type:decimal(5,2)"`
	AvgScore      float64   `json:"avg_score" gorm:"type:decimal(5,2)"`
	TotalFindings int       `json:"total_findings" gorm:"default:0"`
}

func (LPAStatistics) TableName() string {
	return "qms_lpa_statistics"
}
