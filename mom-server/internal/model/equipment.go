package model

import (
	"time"
)

// ========== 设备管理模块 ==========

// Equipment 设备台账
type Equipment struct {
	BaseModel
	TenantID     int64      `json:"tenant_id" gorm:"index;not null"`
	EquipmentCode string    `json:"equipment_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_equip"`
	EquipmentName string    `json:"equipment_name" gorm:"size:100;not null"`
	EquipmentType string    `json:"equipment_type" gorm:"size:20"` // 设备类型
	Brand         *string   `json:"brand" gorm:"size:50"`         // 品牌
	Model         *string   `json:"model" gorm:"size:50"`         // 型号
	SerialNumber  *string   `json:"serial_number" gorm:"size:100"` // 序列号
	WorkshopID    int64     `json:"workshop_id"`                 // 车间ID
	WorkshopName   *string   `json:"workshop_name" gorm:"size:100"`
	LineID        *int64    `json:"line_id"`                    // 生产线ID
	LineName      *string   `json:"line_name" gorm:"size:100"`
	StationID     *int64    `json:"station_id"`                // 工位ID
	StationName   *string   `json:"station_name" gorm:"size:100"`
	Supplier      *string   `json:"supplier" gorm:"size:100"` // 供应商
	PurchaseDate  *time.Time `json:"purchase_date"`            // 采购日期
	PurchasePrice *float64  `json:"purchase_price"`            // 采购价格
	 WarranryEndDate *time.Time `json:"warranty_end_date"` // 保修期结束
	Status        int        `json:"status" gorm:"default:1"` // 1运行/2待机/3故障/4维修/5报废
}

func (Equipment) TableName() string {
	return "equ_equipment"
}

// EquipmentCheck 设备点检
type EquipmentCheck struct {
	BaseModel
	TenantID       int64      `json:"tenant_id" gorm:"index;not null"`
	EquipmentID     int64      `json:"equipment_id"`
	EquipmentCode   string     `json:"equipment_code" gorm:"size:50"`
	EquipmentName   string     `json:"equipment_name" gorm:"size:100"`
	CheckPlanID    int64      `json:"check_plan_id"` // 点检计划ID
	CheckUserID    int64      `json:"check_user_id"`
	CheckUserName  *string    `json:"check_user_name" gorm:"size:50"`
	CheckDate      *time.Time `json:"check_date"`
	CheckResult    int        `json:"check_result"` // 1正常/2异常
	Status         int        `json:"status" gorm:"default:1"` // 1待点检/2已完成/3异常待处理
	Remark         *string    `json:"remark" gorm:"size:500"`
}

func (EquipmentCheck) TableName() string {
	return "equ_equipment_check"
}

// EquipmentMaintenance 设备保养
type EquipmentMaintenance struct {
	BaseModel
	TenantID      int64      `json:"tenant_id" gorm:"index;not null"`
	EquipmentID    int64      `json:"equipment_id"`
	EquipmentCode  string     `json:"equipment_code" gorm:"size:50"`
	EquipmentName  string     `json:"equipment_name" gorm:"size:100"`
	MaintType      string     `json:"maint_type" gorm:"size:20"` // 保养类型
	MaintPlanID   int64      `json:"maint_plan_id"` // 保养计划ID
	MaintUserID   int64      `json:"maint_user_id"`
	MaintUserName *string    `json:"maint_user_name" gorm:"size:50"`
	MaintDate     *time.Time `json:"maint_date"`
	StartTime     *time.Time `json:"start_time"`
	EndTime       *time.Time `json:"end_time"`
	Duration      int        `json:"duration"` // 保养时长(分钟)
	Content       *string    `json:"content" gorm:"type:text"` // 保养内容
	Cost          float64    `json:"cost" gorm:"type:decimal(18,2)"` // 费用
	Status        int        `json:"status" gorm:"default:1"` // 1待执行/2进行中/3已完成
}

func (EquipmentMaintenance) TableName() string {
	return "equ_equipment_maintenance"
}

// EquipmentRepair 设备维修
type EquipmentRepair struct {
	BaseModel
	TenantID      int64      `json:"tenant_id" gorm:"index;not null"`
	EquipmentID    int64      `json:"equipment_id"`
	EquipmentCode  string     `json:"equipment_code" gorm:"size:50"`
	EquipmentName  string     `json:"equipment_name" gorm:"size:100"`
	FaultDesc     string     `json:"fault_desc" gorm:"type:text"` // 故障描述
	FaultTime     *time.Time `json:"fault_time"` // 故障时间
	ReportUserID  int64      `json:"report_user_id"`
	ReportUserName *string   `json:"report_user_name" gorm:"size:50"`
	RepairUserID  *int64     `json:"repair_user_id"` // 维修人
	RepairUserName *string   `json:"repair_user_name" gorm:"size:50"`
	StartTime     *time.Time `json:"start_time"` // 开始时间
	EndTime       *time.Time `json:"end_time"` // 结束时间
	Duration      int        `json:"duration"` // 维修时长(分钟)
	RepairContent *string    `json:"repair_content" gorm:"type:text"` // 维修内容
	Cost          float64    `json:"cost" gorm:"type:decimal(18,2)"` // 维修费用
	Status        int        `json:"status" gorm:"default:1"` // 1待维修/2维修中/3已完成
}

func (EquipmentRepair) TableName() string {
	return "equ_equipment_repair"
}

// SparePart 备件
type SparePart struct {
	BaseModel
	TenantID       int64   `json:"tenant_id" gorm:"index;not null"`
	SparePartCode  string  `json:"spare_part_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_spare"`
	SparePartName  string  `json:"spare_part_name" gorm:"size:100;not null"`
	Spec           *string `json:"spec" gorm:"size:100"` // 规格
	Unit           string  `json:"unit" gorm:"size:20"`  // 单位
	Quantity       float64 `json:"quantity" gorm:"type:decimal(18,2);default:0"` // 库存数量
	MinQuantity    float64 `json:"min_quantity" gorm:"type:decimal(18,2)"` // 最小库存
	Price          float64 `json:"price" gorm:"type:decimal(18,2)"` // 单价
	Supplier       *string `json:"supplier" gorm:"size:100"`
	Status         int     `json:"status" gorm:"default:1"`
}

func (SparePart) TableName() string {
	return "equ_spare_part"
}
