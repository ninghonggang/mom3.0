package model

import (
	"time"
)

// ========== 追溯模块 ==========

// SerialNumber 序列号
type SerialNumber struct {
	BaseModel
	TenantID      int64      `json:"tenant_id" gorm:"index;not null"`
	SerialNumber   string     `json:"serial_number" gorm:"size:50;not null;uniqueIndex:idx_tenant_serial"`
	MaterialID    int64      `json:"material_id"`
	MaterialCode  string     `json:"material_code" gorm:"size:50"`
	MaterialName  string     `json:"material_name" gorm:"size:100"`
	BatchNo       *string    `json:"batch_no" gorm:"size:50"` // 批次号
	LineID        *int64     `json:"line_id"` // 生产线ID
	LineName      *string    `json:"line_name" gorm:"size:100"`
	OrderID        *int64     `json:"order_id"` // 生产工单ID
	OrderNo       *string    `json:"order_no" gorm:"size:50"`
	ProductionDate *time.Time `json:"production_date"`
	Status        int        `json:"status" gorm:"default:1"` // 1在制/2入库/3出库/4已使用
}

func (SerialNumber) TableName() string {
	return "tra_serial_number"
}

// TraceRecord 追溯记录
type TraceRecord struct {
	BaseModel
	TenantID      int64      `json:"tenant_id" gorm:"index;not null"`
	TraceNo       string     `json:"trace_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_trace"`
	MaterialID    int64      `json:"material_id"`
	MaterialCode  string     `json:"material_code" gorm:"size:50"`
	MaterialName  string     `json:"material_name" gorm:"size:100"`
	SerialNumber  *string    `json:"serial_number" gorm:"size:50"`
	BatchNo       *string    `json:"batch_no" gorm:"size:50"`
	ProcessID     int64      `json:"process_id"`
	ProcessName   string     `json:"process_name" gorm:"size:100"`
	StationID     int64      `json:"station_id"`
	StationName   *string    `json:"station_name" gorm:"size:100"`
	OperatorID    *int64     `json:"operator_id"`
	OperatorName  *string    `json:"operator_name" gorm:"size:50"`
	OperateTime   time.Time  `json:"operate_time"`
	OperateType   string     `json:"operate_type" gorm:"size:20"` // 投料/加工/检验/包装
	InputQty      float64    `json:"input_qty" gorm:"type:decimal(18,4)"`
	OutputQty     float64    `json:"output_qty" gorm:"type:decimal(18,4)"`
	RejectQty     float64    `json:"reject_qty" gorm:"type:decimal(18,4);default:0"`
}

func (TraceRecord) TableName() string {
	return "tra_trace_record"
}

// ========== 数据采集 ==========

// DataCollection 数据采集
type DataCollection struct {
	BaseModel
	TenantID    int64      `json:"tenant_id" gorm:"index;not null"`
	EquipmentID int64      `json:"equipment_id"`
	EquipmentCode string    `json:"equipment_code" gorm:"size:50"`
	StationID   int64      `json:"station_id"`
	DataType   string     `json:"data_type" gorm:"size:20"` // 产量/质量/能耗/工艺参数
	DataKey    string     `json:"data_key" gorm:"size:50"` // 数据Key
	DataValue  string     `json:"data_value" gorm:"size:200"` // 数据值
	Unit       *string    `json:"unit" gorm:"size:20"`
	CollectTime time.Time `json:"collect_time"`
}

func (DataCollection) TableName() string {
	return "tra_data_collection"
}

// ========== 能源管理 ==========

// EnergyRecord 能耗记录
type EnergyRecord struct {
	BaseModel
	TenantID     int64      `json:"tenant_id" gorm:"index;not null"`
	EnergyType   string     `json:"energy_type" gorm:"size:20"` // 电/水/气/蒸汽
	WorkshopID   *int64     `json:"workshop_id"`
	WorkshopName *string    `json:"workshop_name" gorm:"size:100"`
	EquipmentID  *int64     `json:"equipment_id"`
	EquipmentName *string   `json:"equipment_name" gorm:"size:100"`
	MeterNo     *string    `json:"meter_no" gorm:"size:50"` // 表号
	Quantity    float64    `json:"quantity" gorm:"type:decimal(18,2)"` // 用量
	Unit        string     `json:"unit" gorm:"size:20"` // kWh/m³/m³/kg
	UnitPrice   *float64   `json:"unit_price"` // 单价
	Amount      *float64   `json:"amount"` // 金额
	RecordDate  time.Time  `json:"record_date"`
	Remark      *string    `json:"remark" gorm:"size:500"`
}

func (EnergyRecord) TableName() string {
	return "ene_energy_record"
}
