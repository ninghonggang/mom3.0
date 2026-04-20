package model

import (
	"time"
)

// ========== 齐套检查模块 ==========

// MesCompleteInspect 齐套检查配置表
type MesCompleteInspect struct {
	BaseModel
	TenantID     int64  `json:"tenant_id" gorm:"index;not null"`
	OrderDayID   int64  `json:"order_day_id" gorm:"index;not null"`
	OrderDayNo   string `json:"order_day_no" gorm:"size:50"`          // 日计划单号
	EnableKitCheck int   `json:"enable_kit_check" gorm:"default:1"`  // 是否启用齐套检查: 0=否, 1=是
	KitCheckStatus string `json:"kit_check_status" gorm:"size:20"`    // 齐套检查状态: PENDING/CHECKING/READY/NOT_READY
	KitCheckTime *time.Time `json:"kit_check_time"`
	KitCheckBy   *int64     `json:"kit_check_by"`
	Remark       string     `json:"remark" gorm:"size:500"`
}

func (MesCompleteInspect) TableName() string {
	return "plan_mes_complete_inspect"
}

// MesConfigInfo 配置信息(用于GET /mes/complete-inspect/get)
type MesConfigInfo struct {
	ID        int64   `json:"id"`
	ConfigCode string  `json:"config_code"`
	ConfigName string  `json:"config_name"`
	ConfigValue string `json:"config_value"`
	Status    int      `json:"status"`
}

// MesWorkSchedulingBaseVO 齐套检查基础请求(MesWorkSchedulingBaseVO)
type MesWorkSchedulingBaseVO struct {
	OrderDayID   int64  `json:"order_day_id" binding:"required"`
	OrderDayNo   string `json:"order_day_no"`
	PlanDate     string `json:"plan_date"`
}

// MesWorkSchedulingPageReqVO 齐套检查分页请求(MesWorkSchedulingPageReqVO)
type MesWorkSchedulingPageReqVO struct {
	OrderDayID   int64  `json:"order_day_id"`
	OrderDayNo   string `json:"order_day_no"`
	PlanDate     string `json:"plan_date"`
	Page         int    `json:"page" gorm:"default:1"`
	PageSize     int    `json:"page_size" gorm:"default:20"`
}

// MesOrderDayBomRespVO 日计划BOM响应(MesOrderDayBomRespVO)
type MesOrderDayBomRespVO struct {
	ID              int64   `json:"id"`
	OrderDayID      int64   `json:"order_day_id"`
	OrderDayNo      string  `json:"order_day_no"`
	ProductID       int64   `json:"product_id"`
	ProductCode     string  `json:"product_code"`
	ProductName     string  `json:"product_name"`
	MaterialID      int64   `json:"material_id"`
	MaterialCode    string  `json:"material_code"`
	MaterialName    string  `json:"material_name"`
	Specification   string  `json:"specification"`
	Unit            string  `json:"unit"`
	RequiredQty     float64 `json:"required_qty"`    // 需求数量
	AvailableQty    float64 `json:"available_qty"`   // 可用数量
	ShortageQty     float64 `json:"shortage_qty"`    // 缺料数量
	KitStatus       string  `json:"kit_status"`      // 齐套状态: READY/NOT_READY/PENDING
	WarehouseID     int64   `json:"warehouse_id"`
	WarehouseName   string  `json:"warehouse_name"`
}

// MesOrderDayWorkerRespVO 日计划人员响应(MesOrderDayWorkerRespVO)
type MesOrderDayWorkerRespVO struct {
	ID              int64   `json:"id"`
	OrderDayID      int64   `json:"order_day_id"`
	OrderDayNo      string  `json:"order_day_no"`
	ProductID       int64   `json:"product_id"`
	ProductCode     string  `json:"product_code"`
	ProductName     string  `json:"product_name"`
	ProcessRouteID  int64   `json:"process_route_id"`
	ProcessRouteName string `json:"process_route_name"`
	WorkerID        int64   `json:"worker_id"`
	WorkerCode      string  `json:"worker_code"`
	WorkerName      string  `json:"worker_name"`
	TeamID          int64   `json:"team_id"`
	TeamName        string  `json:"team_name"`
	ShiftType       string  `json:"shift_type"`
	KitStatus       string  `json:"kit_status"`       // 齐套状态: READY/NOT_READY/PENDING
}

// MesOrderDayEquipmentRespVO 日计划设备响应(MesOrderDayEquipmentRespVO)
type MesOrderDayEquipmentRespVO struct {
	ID                int64   `json:"id"`
	OrderDayID        int64   `json:"order_day_id"`
	OrderDayNo        string  `json:"order_day_no"`
	ProductID         int64   `json:"product_id"`
	ProductCode       string  `json:"product_code"`
	ProductName       string  `json:"product_name"`
	ProcessRouteID    int64   `json:"process_route_id"`
	ProcessRouteName  string  `json:"process_route_name"`
	EquipmentID       int64   `json:"equipment_id"`
	EquipmentCode     string  `json:"equipment_code"`
	EquipmentName     string  `json:"equipment_name"`
	WorkstationID     int64   `json:"workstation_id"`
	WorkstationName   string  `json:"workstation_name"`
	Status            string  `json:"status"`          // 设备状态
	KitStatus         string  `json:"kit_status"`      // 齐套状态: READY/NOT_READY/PENDING
}

// MesWorkSchedulingUpdateReqVO 更新生产日工单请求(MesWorkSchedulingUpdateReqVO)
type MesWorkSchedulingUpdateReqVO struct {
	OrderDayID       int64  `json:"order_day_id" binding:"required"`
	KitCheckStatus   string `json:"kit_check_status"`  // PENDING/CHECKING/READY/NOT_READY
	KitCheckRemark   string `json:"kit_check_remark"`
}

// CompleteInspectUpdate 更新齐套检查状态请求
type CompleteInspectUpdate struct {
	OrderDayID     int64  `json:"order_day_id" binding:"required"`
	KitCheckStatus string `json:"kit_check_status"`
	Remark         string `json:"remark"`
}
