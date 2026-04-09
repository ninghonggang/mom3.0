package model

import (
	"time"
)

// ========== 主数据模块 ==========

// Material 物料
type Material struct {
	BaseModel
	TenantID     int64   `json:"tenant_id" gorm:"index;not null"`
	MaterialCode string   `json:"material_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_material"`
	MaterialName string   `json:"material_name" gorm:"size:100;not null"`
	MaterialType string   `json:"material_type" gorm:"size:20"` // 原材料/半成品/成品
	Spec         *string `json:"spec" gorm:"size:100"`          // 规格
	Unit         string   `json:"unit" gorm:"size:20"`           // 单位
	UnitName     *string `json:"unit_name" gorm:"size:20"`       // 单位名称
	Weight       *float64 `json:"weight"`                        // 重量
	Length       *float64 `json:"length"`                        // 长度
	Width        *float64 `json:"width"`                        // 宽度
	Height       *float64 `json:"height"`                        // 高度
	CategoryID   *int64   `json:"category_id"`                 // 分类ID
	Status       int      `json:"status" gorm:"default:1"`
}

func (Material) TableName() string {
	return "mdm_material"
}

// MaterialCategory 物料分类
type MaterialCategory struct {
	BaseModel
	TenantID   int64  `json:"tenant_id" gorm:"index;not null"`
	ParentID   int64  `json:"parent_id" gorm:"default:0;index"`
	CategoryName string `json:"category_name" gorm:"size:100;not null"`
	CategoryCode string `json:"category_code" gorm:"size:50"`
	Sort       int     `json:"sort" gorm:"default:0"`
	Status     int     `json:"status" gorm:"default:1"`
	Children   []MaterialCategory `json:"children" gorm:"-"`
}

func (MaterialCategory) TableName() string {
	return "mdm_material_category"
}

// BOM 物料清单
type BOM struct {
	BaseModel
	TenantID    int64   `json:"tenant_id" gorm:"index;not null"`
	BOMCode     string  `json:"bom_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_bom"`
	ProductID   int64   `json:"product_id" gorm:"not null"` // 产品ID
	ProductCode string  `json:"product_code" gorm:"size:50"`
	ProductName string  `json:"product_name" gorm:"size:100"`
	Version     string  `json:"version" gorm:"size:20"` // 版本
	BOMType     string  `json:"bom_type" gorm:"size:20"` // 生产BOM/维修BOM
	Status      int     `json:"status" gorm:"default:1"` // 1草稿/2生效/3失效
	EffDate     *time.Time `json:"eff_date"`             // 生效日期
	ExpDate     *time.Time `json:"exp_date"`             // 失效日期
}

func (BOM) TableName() string {
	return "mdm_bom"
}

// BOMItem BOM明细
type BOMItem struct {
	BaseModel
	BOMID       int64   `json:"bom_id" gorm:"index;not null"`
	MaterialID  int64   `json:"material_id" gorm:"not null"`
	MaterialCode string  `json:"material_code" gorm:"size:50"`
	MaterialName string  `json:"material_name" gorm:"size:100"`
	Quantity    float64 `json:"quantity" gorm:"type:decimal(18,4)"` // 用量
	Unit        string  `json:"unit" gorm:"size:20"`
	ScrapRate   float64 `json:"scrap_rate" gorm:"default:0"` // 损耗率
	Level       int     `json:"level" gorm:"default:1"`        // 层级
	ParentID    *int64  `json:"parent_id"`                     // 父级ID
	Sort        int     `json:"sort" gorm:"default:0"`
}

func (BOMItem) TableName() string {
	return "mdm_bom_item"
}

// Process 工艺路线
type Process struct {
	BaseModel
	TenantID    int64   `json:"tenant_id" gorm:"index;not null"`
	ProcessCode string  `json:"process_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_process"`
	ProcessName string  `json:"process_name" gorm:"size:100;not null"`
	ProcessType string  `json:"process_type" gorm:"size:20"` // 加工/检验/装配
	Sequence    int     `json:"sequence" gorm:"default:0"`   // 序号
	StationID   *int64  `json:"station_id"`                   // 工位ID
	Status      int     `json:"status" gorm:"default:1"`
}

func (Process) TableName() string {
	return "mdm_operation"
}

// Route 工序
type Route struct {
	BaseModel
	TenantID   int64   `json:"tenant_id" gorm:"index;not null"`
	RouteCode  string  `json:"route_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_route"`
	RouteName  string  `json:"route_name" gorm:"size:100;not null"`
	MaterialID int64   `json:"material_id" gorm:"not null"` // 物料ID
	Version    string  `json:"version" gorm:"size:20"`      // 版本
	Status     int     `json:"status" gorm:"default:1"`
}

func (Route) TableName() string {
	return "mdm_route"
}

// RouteOperation 工序明细
type RouteOperation struct {
	BaseModel
	RouteID     int64   `json:"route_id" gorm:"index;not null"`
	ProcessID   int64   `json:"process_id" gorm:"not null"`
	ProcessName string  `json:"process_name" gorm:"size:100"`
	Sequence    int     `json:"sequence" gorm:"default:0"`
	StationID   *int64  `json:"station_id"`
	StationName *string `json:"station_name" gorm:"size:100"`
	StandardTime int    `json:"standard_time"` // 标准工时(秒)
	QueueTime   int     `json:"queue_time"`    // 排队时间(秒)
	TransportTime int   `json:"transport_time"` // 搬运时间(秒)
}

func (RouteOperation) TableName() string {
	return "mdm_route_operation"
}

// Workshop 车间
type Workshop struct {
	BaseModel
	TenantID     int64   `json:"tenant_id" gorm:"index;not null"`
	WorkshopCode string  `json:"workshop_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_workshop"`
	WorkshopName string  `json:"workshop_name" gorm:"size:100;not null"`
	WorkshopType string  `json:"workshop_type" gorm:"size:20"` // 加工/装配/检验
	Manager      *string `json:"manager" gorm:"size:50"`
	Phone        *string `json:"phone" gorm:"size:20"`
	Address      *string `json:"address" gorm:"size:200"`
	Status       int     `json:"status" gorm:"default:1"`
}

func (Workshop) TableName() string {
	return "mdm_workshop"
}

// ProductionLine 生产线
type ProductionLine struct {
	BaseModel
	TenantID     int64   `json:"tenant_id" gorm:"index;not null"`
	LineCode     string  `json:"line_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_line"`
	LineName     string  `json:"line_name" gorm:"size:100;not null"`
	WorkshopID   int64   `json:"workshop_id" gorm:"not null"`
	LineType     string  `json:"line_type" gorm:"size:20"` // 自动化/半自动化/手工
	Status       int     `json:"status" gorm:"default:1"`
}

func (ProductionLine) TableName() string {
	return "mdm_production_line"
}

// Workstation 工位
type Workstation struct {
	BaseModel
	TenantID   int64   `json:"tenant_id" gorm:"index;not null"`
	StationCode string  `json:"station_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_station"`
	StationName string  `json:"station_name" gorm:"size:100;not null"`
	LineID     int64   `json:"line_id" gorm:"not null"`
	StationType string  `json:"station_type" gorm:"size:20"` // 加工/装配/检验/物料
	Status     int     `json:"status" gorm:"default:1"`
}

func (Workstation) TableName() string {
	return "mdm_workstation"
}

// Shift 班次 (旧版本，兼容)
type Shift struct {
	BaseModel
	TenantID   int64   `json:"tenant_id" gorm:"index;not null"`
	ShiftCode  string  `json:"shift_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_shift"`
	ShiftName  string  `json:"shift_name" gorm:"size:100;not null"`
	StartTime  string  `json:"start_time" gorm:"size:10"` // HH:mm
	EndTime    string  `json:"end_time" gorm:"size:10"`   // HH:mm
	BreakStart *string `json:"break_start" gorm:"size:10"`
	BreakEnd   *string `json:"break_end" gorm:"size:10"`
	Status     int     `json:"status" gorm:"default:1"`
}

func (Shift) TableName() string {
	return "mdm_shift"
}

// MdmShift MDM班次
type MdmShift struct {
	BaseModel
	TenantID   int64   `json:"tenant_id" gorm:"index;not null"`
	ShiftCode  string  `json:"shift_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_mdm_shift_code"`
	ShiftName  string  `json:"shift_name" gorm:"size:100;not null"`
	StartTime  string  `json:"start_time" gorm:"size:10"` // HH:mm
	EndTime    string  `json:"end_time" gorm:"size:10"`   // HH:mm
	WorkHours  float64 `json:"work_hours" gorm:"type:decimal(10,2);default:8"` // 工作时长
	IsNight    int     `json:"is_night" gorm:"default:0"` // 是否夜班 0否 1是
	Remark     *string `json:"remark" gorm:"size:500"`
}

func (MdmShift) TableName() string {
	return "mdm_shift"
}
