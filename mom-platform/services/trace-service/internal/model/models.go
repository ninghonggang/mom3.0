package model

import (
	"time"

	"gorm.io/gorm"
)

// TraceRecord 追溯记录
type TraceRecord struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	TenantID          string         `gorm:"column:tenant_id;size:64;index" json:"tenant_id"`
	TraceNo           string         `gorm:"column:trace_no;size:128;uniqueIndex" json:"trace_no"`
	TraceType         string         `gorm:"column:trace_type;size:32" json:"trace_type"` // SERIAL/BATCH/ORDER/MATERIAL
	SerialNo          string         `gorm:"column:serial_no;size:128;index" json:"serial_no"`
	BatchNo           string         `gorm:"column:batch_no;size:128;index" json:"batch_no"`
	MaterialID        string         `gorm:"column:material_id;size:64;index" json:"material_id"`
	ProductionOrderID string         `gorm:"column:production_order_id;size:64;index" json:"production_order_id"`
	Status            string         `gorm:"column:status;size:32;default:PENDING" json:"status"` // PENDING/ACTIVE/BROKEN/ARCHIVED
	TraceAt           time.Time      `gorm:"column:trace_at" json:"trace_at"`
	CreatedAt         time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}

func (TraceRecord) TableName() string { return "trace_records" }

// TraceLink 追溯链
type TraceLink struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	TraceID        uint           `gorm:"column:trace_id;index" json:"trace_id"`
	ParentTraceID  *uint          `gorm:"column:parent_trace_id;index" json:"parent_trace_id"`
	LinkType       string         `gorm:"column:link_type;size:32" json:"link_type"` // MATERIAL/PROCESS/ORDER
	Level          int            `gorm:"column:level" json:"level"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	Trace       TraceRecord  `gorm:"foreignKey:TraceID" json:"trace,omitempty"`
	ParentTrace *TraceRecord `gorm:"foreignKey:ParentTraceID" json:"parent_trace,omitempty"`
}

func (TraceLink) TableName() string { return "trace_links" }

// SerialNumber 序列号
type SerialNumber struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	SerialNo          string         `gorm:"column:serial_no;size:128;uniqueIndex" json:"serial_no"`
	MaterialID        string         `gorm:"column:material_id;size:64;index" json:"material_id"`
	ProductionOrderID string         `gorm:"column:production_order_id;size:64;index" json:"production_order_id"`
	BatchNo           string         `gorm:"column:batch_no;size:128;index" json:"batch_no"`
	Status            string         `gorm:"column:status;size:32;default:UNUSED" json:"status"`
	CreatedAt         time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}

func (SerialNumber) TableName() string { return "serial_numbers" }

// DataPoint 数据采集点
type DataPoint struct {
	ID                     uint           `gorm:"primaryKey" json:"id"`
	TenantID               string         `gorm:"column:tenant_id;size:64;index" json:"tenant_id"`
	PointCode              string         `gorm:"column:point_code;size:128;uniqueIndex" json:"point_code"`
	PointName              string         `gorm:"column:point_name;size:256" json:"point_name"`
	EquipmentID            string         `gorm:"column:equipment_id;size:64;index" json:"equipment_id"`
	DataType               string         `gorm:"column:data_type;size:32" json:"data_type"` // NUMBER/STRING/BOOLEAN
	UpperLimit             *float64       `gorm:"column:upper_limit" json:"upper_limit"`
	LowerLimit             *float64       `gorm:"column:lower_limit" json:"lower_limit"`
	CollectIntervalSeconds int            `gorm:"column:collect_interval_seconds;default:60" json:"collect_interval_seconds"`
	Status                 string         `gorm:"column:status;size:32;default:ACTIVE" json:"status"` // ACTIVE/PAUSED/ERROR
	CreatedAt              time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt              time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"index" json:"-"`
}

func (DataPoint) TableName() string { return "data_points" }

// CollectRecord 采集记录
type CollectRecord struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	TenantID    string         `gorm:"column:tenant_id;size:64;index" json:"tenant_id"`
	DataPointID uint           `gorm:"column:data_point_id;index" json:"data_point_id"`
	Value       string         `gorm:"column:value;size:256" json:"value"`
	Quality     string         `gorm:"column:quality;size:32" json:"quality"` // GOOD/BAD/UNCERTAIN
	CollectedAt time.Time      `gorm:"column:collected_at;index" json:"collected_at"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	DataPoint *DataPoint `gorm:"foreignKey:DataPointID" json:"data_point,omitempty"`
}

func (CollectRecord) TableName() string { return "collect_records" }

// ScanLog 扫码日志
type ScanLog struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	TenantID      string         `gorm:"column:tenant_id;size:64;index" json:"tenant_id"`
	ScanCode      string         `gorm:"column:scan_code;size:256" json:"scan_code"`
	ScanType      string         `gorm:"column:scan_type;size:32" json:"scan_type"`
	OperatorID    string         `gorm:"column:operator_id;size:64" json:"operator_id"`
	EquipmentID   string         `gorm:"column:equipment_id;size:64" json:"equipment_id"`
	WorkstationID string         `gorm:"column:workstation_id;size:64" json:"workstation_id"`
	TraceID       *uint          `gorm:"column:trace_id" json:"trace_id"`
	ScanTime      time.Time      `gorm:"column:scan_time" json:"scan_time"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	Trace *TraceRecord `gorm:"foreignKey:TraceID" json:"trace,omitempty"`
}

func (ScanLog) TableName() string { return "scan_logs" }
