package model

import (
	"time"
)

// InstrumentType 仪器类型
type InstrumentType string

const (
	InstrumentTypeMeasuring  InstrumentType = "MEASURING"  // 测量仪器
	InstrumentTypeTesting   InstrumentType = "TESTING"    // 测试仪器
	InstrumentTypeAnalytical InstrumentType = "ANALYTICAL" // 分析仪器
)

// CalibrationStatus 校准状态
type CalibrationStatus string

const (
	CalibrationStatusCalibrated CalibrationStatus = "CALIBRATED" // 已校准
	CalibrationStatusDue       CalibrationStatus = "DUE"        // 即将到期
	CalibrationStatusOverdue   CalibrationStatus = "OVERDUE"   // 已过期
)

// InstrumentStatus 仪器状态
type InstrumentStatus string

const (
	InstrumentStatusActive      InstrumentStatus = "ACTIVE"      // 正常
	InstrumentStatusMaintenance InstrumentStatus = "MAINTENANCE" // 维护中
	InstrumentStatusRetired    InstrumentStatus = "RETIRED"     // 已报废
)

// LabInstrument 实验室仪器表
type LabInstrument struct {
	ID                  uint64            `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID            uint64            `gorm:"index;not null;default:1" json:"tenant_id"`
	InstrumentCode      string            `gorm:"size:50;uniqueIndex" json:"instrument_code"`        // 仪器编码
	InstrumentName      string            `gorm:"size:100" json:"instrument_name"`                 // 仪器名称
	InstrumentType      InstrumentType    `gorm:"size:20;index" json:"instrument_type"`            // 仪器类型
	Manufacturer        string            `gorm:"size:100" json:"manufacturer"`                     // 制造商
	Model               string            `gorm:"size:100" json:"model"`                           // 型号
	SerialNumber        string            `gorm:"size:100" json:"serial_number"`                    // 序列号
	CalibrationCycle    int               `json:"calibration_cycle"`                                // 校准周期(天数)
	LastCalibrationDate time.Time         `json:"last_calibration_date"`                           // 上次校准日期
	NextCalibrationDate time.Time         `json:"next_calibration_date"`                           // 下次校准日期
	CalibrationStatus   CalibrationStatus `gorm:"size:20;index" json:"calibration_status"`         // 校准状态
	LabLocation         string            `gorm:"size:100" json:"lab_location"`                    // 实验室位置
	Status              InstrumentStatus  `gorm:"size:20;index" json:"status"`                     // 状态
	Remark              string            `gorm:"size:500" json:"remark"`                          // 备注
	CreatedBy           string            `gorm:"size:50" json:"created_by"`                       // 创建人
	UpdatedBy           string            `gorm:"size:50" json:"updated_by"`                       // 更新人
	CreatedAt           time.Time         `json:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at"`
}

func (LabInstrument) TableName() string {
	return "lab_instrument"
}

// LabInstrumentQuery 查询实验室仪器
type LabInstrumentQuery struct {
	InstrumentCode      string `form:"instrument_code"`
	InstrumentName      string `form:"instrument_name"`
	InstrumentType      string `form:"instrument_type"`
	CalibrationStatus   string `form:"calibration_status"`
	Status              string `form:"status"`
	Page                int    `form:"page"`
	PageSize            int    `form:"page_size"`
}

// CalibrationResult 校准结果
type CalibrationResult string

const (
	CalibrationResultPass CalibrationResult = "PASS" // 合格
	CalibrationResultFail CalibrationResult = "FAIL" // 不合格
)

// LabCalibration 校准记录表
type LabCalibration struct {
	ID                 uint64            `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID           uint64            `gorm:"index;not null;default:1" json:"tenant_id"`
	InstrumentID       uint64            `gorm:"index;not null" json:"instrument_id"`              // 仪器ID
	CalibrationDate    time.Time         `json:"calibration_date"`                                 // 校准日期
	CalibrationResult  CalibrationResult `gorm:"size:10" json:"calibration_result"`               // 校准结果
	CalibratedBy       string            `gorm:"size:50" json:"calibrated_by"`                     // 校准人
	CertificateNo      string            `gorm:"size:100" json:"certificate_no"`                   // 证书编号
	NextCalibrationDate time.Time        `json:"next_calibration_date"`                            // 下次校准日期
	CalibrationItems   string            `gorm:"type:text" json:"calibration_items"`               // 校准项目JSON
	AttachmentURL      string            `gorm:"size:500" json:"attachment_url"`                   // 证书附件URL
	Remark             string            `gorm:"size:500" json:"remark"`                           // 备注
	CreatedBy          string            `gorm:"size:50" json:"created_by"`                        // 创建人
	UpdatedBy          string            `gorm:"size:50" json:"updated_by"`                        // 更新人
	CreatedAt          time.Time         `json:"created_at"`
	UpdatedAt          time.Time         `json:"updated_at"`
}

func (LabCalibration) TableName() string {
	return "lab_calibration"
}

// LabCalibrationQuery 查询校准记录
type LabCalibrationQuery struct {
	InstrumentID uint64 `form:"instrument_id"`
	Page         int    `form:"page"`
	PageSize     int    `form:"page_size"`
}

// LabInstrumentCreateRequest 创建仪器请求
type LabInstrumentCreateRequest struct {
	InstrumentCode      string         `json:"instrument_code" binding:"required"`
	InstrumentName      string         `json:"instrument_name" binding:"required"`
	InstrumentType      InstrumentType `json:"instrument_type" binding:"required"`
	Manufacturer        string         `json:"manufacturer"`
	Model               string         `json:"model"`
	SerialNumber        string         `json:"serial_number"`
	CalibrationCycle    int            `json:"calibration_cycle"`
	LastCalibrationDate time.Time      `json:"last_calibration_date"`
	NextCalibrationDate time.Time      `json:"next_calibration_date"`
	CalibrationStatus   CalibrationStatus `json:"calibration_status"`
	LabLocation         string         `json:"lab_location"`
	Status              InstrumentStatus `json:"status"`
	Remark              string         `json:"remark"`
}

// LabInstrumentUpdateRequest 更新仪器请求
type LabInstrumentUpdateRequest struct {
	InstrumentName      string            `json:"instrument_name"`
	InstrumentType      InstrumentType    `json:"instrument_type"`
	Manufacturer        string            `json:"manufacturer"`
	Model               string            `json:"model"`
	SerialNumber        string            `json:"serial_number"`
	CalibrationCycle    int               `json:"calibration_cycle"`
	LastCalibrationDate time.Time         `json:"last_calibration_date"`
	NextCalibrationDate time.Time         `json:"next_calibration_date"`
	CalibrationStatus   CalibrationStatus `json:"calibration_status"`
	LabLocation         string            `json:"lab_location"`
	Status              InstrumentStatus  `json:"status"`
	Remark              string            `json:"remark"`
}

// RecordCalibrationRequest 记录校准请求
type RecordCalibrationRequest struct {
	CalibrationDate     time.Time        `json:"calibration_date" binding:"required"`
	CalibrationResult   CalibrationResult `json:"calibration_result" binding:"required"`
	CalibratedBy        string           `json:"calibrated_by" binding:"required"`
	CertificateNo       string           `json:"certificate_no"`
	NextCalibrationDate time.Time        `json:"next_calibration_date"`
	CalibrationItems    string           `json:"calibration_items"`
	AttachmentURL       string           `json:"attachment_url"`
	Remark              string           `json:"remark"`
}