package model

import "time"

// DCDataPoint 数据采集点配置
type DCDataPoint struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	TenantID     int64     `json:"tenant_id" gorm:"index;not null"`
	PointCode    string    `json:"point_code" gorm:"size:50;uniqueIndex:idx_tenant_point;not null"`
	PointName    string    `json:"point_name" gorm:"size:100;not null"`
	DeviceID     int64     `json:"device_id" gorm:"not null"`
	DataType     string    `json:"data_type" gorm:"size:20"`                     // ANALOG/DIGITAL/TEXT/ENUM
	Protocol     string    `json:"protocol" gorm:"size:20;not null"`           // OPCUA/MQTT/MODBUS/MANUAL
	Address      string    `json:"address" gorm:"size:200"`                      // OPC节点ID / MQTT Topic / Modbus地址
	Unit         string    `json:"unit" gorm:"size:20"`
	ScanRate     int       `json:"scan_rate" gorm:"default:1000"`               // 采集周期(ms)
	Deadband     float64   `json:"deadband" gorm:"type:decimal(10,4);default:0"` // 死区值
	StorePolicy  string    `json:"store_policy" gorm:"size:20;default:'ALL'"`    // ALL/CHANGE/PERIOD
	AlarmEnabled int       `json:"alarm_enabled" gorm:"default:0"`
	AlarmHigh    *float64  `json:"alarm_high" gorm:"type:decimal(18,4)"`
	AlarmLow     *float64  `json:"alarm_low" gorm:"type:decimal(18,4)"`
	MapToField   string    `json:"map_to_field" gorm:"size:100"` // 映射到OEE/报工字段
	Status       string    `json:"status" gorm:"size:20;default:'ACTIVE'"` // ACTIVE/INACTIVE
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (DCDataPoint) TableName() string {
	return "dc_data_point"
}

// DCCollectRecord 采集记录（只写不查，定期清理）
type DCCollectRecord struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	TenantID      int64     `json:"tenant_id" gorm:"index;not null"`
	PointID      int64     `json:"point_id" gorm:"index;not null"`
	DeviceID     int64     `json:"device_id" gorm:"not null"`
	ValueRaw     string    `json:"value_raw" gorm:"size:200"`
	ValueNumeric *float64  `json:"value_numeric" gorm:"type:decimal(18,6)"`
	ValueText    string    `json:"value_text" gorm:"size:500"`
	Quality      string    `json:"quality" gorm:"size:20;default:'GOOD'` // GOOD/BAD/UNCERTAIN
	CollectTime  time.Time `json:"collect_time" gorm:"not null;index"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (DCCollectRecord) TableName() string {
	return "dc_collect_record"
}

// DCScanLog 扫描记录
type DCScanLog struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	TenantID      int64     `json:"tenant_id" gorm:"index;not null"`
	WorkshopID    *int64    `json:"workshop_id" gorm:"index"`
	ScanType      string    `json:"scan_type" gorm:"size:20;not null"`        // MATERIAL/PRODUCT/TOOL/CONTAINER
	ScanCode      string    `json:"scan_code" gorm:"size:200;not null"`
	ScanDevice    string    `json:"scan_device" gorm:"size:50"` // 扫描枪设备ID
	WorkstationID *int64    `json:"workstation_id" gorm:"index"`
	ScanUserID    *int64    `json:"scan_user_id" gorm:"index"`
	ScanTime      time.Time `json:"scan_time" gorm:"index"`
	ParseResult   string    `json:"parse_result" gorm:"type:jsonb"`      // 解析结果
	BusinessType  string    `json:"business_type" gorm:"size:50"`       // 触发的业务类型
	RelatedID     *int64    `json:"related_id" gorm:"index"`             // 关联业务ID
	Status        string    `json:"status" gorm:"size:20;default:'SUCCESS'"` // SUCCESS/FAILED/DUPLICATE
	FailReason    string    `json:"fail_reason" gorm:"size:200"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (DCScanLog) TableName() string {
	return "dc_scan_log"
}
