package model

import (
	"time"
)

// TEEPData TEEP分析数据 (Total Effective Equipment Performance)
type TEEPData struct {
	BaseModel
	TenantID      int64     `json:"tenant_id" gorm:"index;not null"`
	EquipmentID   int64     `json:"equipment_id"`
	EquipmentCode string    `json:"equipment_code" gorm:"size:50"`
	EquipmentName string    `json:"equipment_name" gorm:"size:100"`
	ReportDate    time.Time `json:"report_date"` // 统计日期
	// 时间稼动率 (Availability)
	PlanTime      float64   `json:"plan_time"`       // 计划开机时间(分钟)
	DownTime      float64   `json:"down_time"`       // 停机时间(分钟)
	ActualTime    float64   `json:"actual_time"`     // 实际开机时间(分钟)
	Availability  float64   `json:"availability"`    // 时间稼动率 = ActualTime / PlanTime
	// 性能稼动率 (Performance)
	IdealCycleTime float64  `json:"ideal_cycle_time"` // 标准周期时间(秒)
	ActualOutput   int64    `json:"actual_output"`    // 实际产量
	IdealOutput    float64  `json:"ideal_output"`     // 理论产量
	Performance    float64  `json:"performance"`      // 性能稼动率 = IdealOutput / ActualOutput
	// 良品率 (Quality)
	PassOutput    int64     `json:"pass_output"`    // 良品数量
	FailOutput    int64     `json:"fail_output"`    // 不良品数量
	Quality       float64   `json:"quality"`        // 良品率 = PassOutput / ActualOutput
	// TEEP综合指标
	TEEP          float64   `json:"teep"`           // TEEP = Availability * Performance * Quality
	OEE           float64   `json:"oee"`            // OEE = Performance * Quality * Availability
	Remark        *string   `json:"remark" gorm:"size:500"`
}

func (TEEPData) TableName() string {
	return "equ_teep_data"
}

// Mold 模具管理
type Mold struct {
	BaseModel
	TenantID     int64      `json:"tenant_id" gorm:"index;not null"`
	MoldCode     string     `json:"mold_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_mold"`
	MoldName     string     `json:"mold_name" gorm:"size:100;not null"`
	MoldType     string     `json:"mold_type" gorm:"size:50"` // 模具类型
	ProductCode  *string    `json:"product_code" gorm:"size:50"` // 适用产品编码
	ProductName  *string    `json:"product_name" gorm:"size:100"` // 适用产品名称
	CavityCount  int        `json:"cavity_count"` // 型腔数
	Lifecycle    int        `json:"lifecycle"` // 设计寿命(次数)
	UsedCount    int        `json:"used_count"` // 已使用次数
	WorkshopID   int64      `json:"workshop_id"`
	WorkshopName *string    `json:"workshop_name" gorm:"size:100"`
	LocationID   *int64     `json:"location_id"`
	LocationName *string    `json:"location_name" gorm:"size:100"`
	Status       int        `json:"status" gorm:"default:1"` // 1正常使用/2维修中/3闲置/4报废
	Remark       *string    `json:"remark" gorm:"size:500"`
}

func (Mold) TableName() string {
	return "equ_mold"
}

// MoldMaintenance 模具保养记录
type MoldMaintenance struct {
	BaseModel
	TenantID    int64      `json:"tenant_id" gorm:"index;not null"`
	MoldID      int64      `json:"mold_id"`
	MoldCode    string     `json:"mold_code" gorm:"size:50"`
	MoldName    string     `json:"mold_name" gorm:"size:100"`
	MaintType   string     `json:"maint_type" gorm:"size:20"` // 保养类型
	MaintDate   time.Time  `json:"maint_date"`
	MaintUserID int64      `json:"maint_user_id"`
	UserName    *string    `json:"user_name" gorm:"size:50"`
	Content     *string    `json:"content" gorm:"type:text"`
	Duration    int        `json:"duration"` // 保养时长(分钟)
	Cost        float64    `json:"cost" gorm:"type:decimal(18,2)"`
	Status      int        `json:"status" gorm:"default:1"` // 1待执行/2进行中/3已完成
}

func (MoldMaintenance) TableName() string {
	return "equ_mold_maintenance"
}

// MoldRepair 模具维修记录
type MoldRepair struct {
	BaseModel
	TenantID    int64      `json:"tenant_id" gorm:"index;not null"`
	MoldID      int64      `json:"mold_id"`
	MoldCode    string     `json:"mold_code" gorm:"size:50"`
	MoldName    string     `json:"mold_name" gorm:"size:100"`
	RepairType  string     `json:"repair_type" gorm:"size:20"` // 维修类型
	RepairDate  time.Time  `json:"repair_date"`
	RepairUserID int64     `json:"repair_user_id"`
	UserName    *string    `json:"user_name" gorm:"size:50"`
	Reason      *string    `json:"reason" gorm:"size:500"` // 损坏原因
	Content     *string    `json:"content" gorm:"type:text"` // 维修内容
	PartsUsed   *string    `json:"parts_used" gorm:"size:500"` // 使用配件
	Duration    int        `json:"duration"` // 维修时长(小时)
	Cost        float64    `json:"cost" gorm:"type:decimal(18,2)"`
	Status      int        `json:"status" gorm:"default:1"` // 1待维修/2维修中/3已完成
}

func (MoldRepair) TableName() string {
	return "equ_mold_repair"
}

// Gauge 量检具管理
type Gauge struct {
	BaseModel
	TenantID      int64      `json:"tenant_id" gorm:"index;not null"`
	GaugeCode     string     `json:"gauge_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_gauge"`
	GaugeName     string     `json:"gallery_name" gorm:"size:100;not null"`
	GaugeType     string     `json:"gauge_type" gorm:"size:20"` // 量检具类型
	Spec          *string    `json:"spec" gorm:"size:100"` // 规格型号
	Precision     *string    `json:"precision" gorm:"size:50"` // 精度等级
	MeasureRange  *string    `json:"measure_range" gorm:"size:100"` // 测量范围
	CalCycle      int        `json:"cal_cycle"` // 校准周期(天)
	LastCalDate   *time.Time `json:"last_cal_date"` // 上次校准日期
	NextCalDate   *time.Time `json:"next_cal_date"` // 下次校准日期
	WorkshopID    int64      `json:"workshop_id"`
	WorkshopName  *string    `json:"workshop_name" gorm:"size:100"`
	LocationID    *int64     `json:"location_id"`
	LocationName  *string    `json:"location_name" gorm:"size:100"`
	Status        int        `json:"status" gorm:"default:1"` // 1正常使用/2校准中/3维修中/4报废
	Remark        *string    `json:"remark" gorm:"size:500"`
}

func (Gauge) TableName() string {
	return "equ_gauge"
}

// GaugeCalibration 量检具校准记录
type GaugeCalibration struct {
	BaseModel
	TenantID    int64      `json:"tenant_id" gorm:"index;not null"`
	GaugeID     int64      `json:"gauge_id"`
	GaugeCode   string     `json:"gauge_code" gorm:"size:50"`
	GaugeName   string     `json:"gauge_name" gorm:"size:100"`
	CalDate     time.Time  `json:"cal_date"`
	CalType     string     `json:"cal_type" gorm:"size:20"` // 校准类型
	CalResult   int        `json:"cal_result"` // 1合格/2不合格
	Standard    *string    `json:"standard" gorm:"size:100"` // 标准值
	ActualValue *string    `json:"actual_value" gorm:"size:100"` // 实际值
	ErrorValue  *string    `json:"error_value" gorm:"size:100"` // 误差值
	CalUserID   int64      `json:"cal_user_id"`
	UserName    *string    `json:"user_name" gorm:"size:50"`
	Agency      *string    `json:"agency" gorm:"size:100"` // 校准机构
	Certificate *string    `json:"certificate" gorm:"size:100"` // 证书编号
	Cost        float64    `json:"cost" gorm:"type:decimal(18,2)"`
	Remark      *string    `json:"remark" gorm:"size:500"`
}

func (GaugeCalibration) TableName() string {
	return "equ_gauge_calibration"
}
