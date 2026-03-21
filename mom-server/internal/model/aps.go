package model

import (
	"time"
)

// ========== APS模块 ==========

// MRP物料需求计划
type MRP struct {
	BaseModel
	TenantID     int64      `json:"tenant_id" gorm:"index;not null"`
	MRPNo        string     `json:"mrp_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_mrp"`
	MRPType      string     `json:"mrp_type" gorm:"size:20"` // MPS/MRP
	PlanDate     *time.Time `json:"plan_date"` // 计划日期
	Status       int        `json:"status" gorm:"default:1"` // 1待计算/2计算中/3已完成
	Remark       *string    `json:"remark" gorm:"size:500"`
}

// MRPItem MRP明细
type MRPItem struct {
	BaseModel
	MRPID       int64   `json:"mrp_id" gorm:"index"`
	MaterialID  int64   `json:"material_id"`
	MaterialCode string  `json:"material_code" gorm:"size:50"`
	MaterialName string  `json:"material_name" gorm:"size:100"`
	Quantity    float64 `json:"quantity" gorm:"type:decimal(18,4)"` // 需求数量
	StockQty    float64 `json:"stock_qty" gorm:"type:decimal(18,4)"` // 库存数量
	AllocatedQty float64 `json:"allocated_qty" gorm:"type:decimal(18,4)"` // 已分配数量
	NetQty      float64 `json:"net_qty" gorm:"type:decimal(18,4)"` // 净需求
	SourceType  string  `json:"source_type" gorm:"size:20"` // 订单/库存/采购
	SourceNo    *string `json:"source_no" gorm:"size:50"`
}

// MPS主生产计划
type MPS struct {
	BaseModel
	TenantID     int64      `json:"tenant_id" gorm:"index;not null"`
	MPSNo        string     `json:"mps_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_mps"`
	PlanMonth    string     `json:"plan_month" gorm:"size:10"` // 计划月份 YYYY-MM
	MaterialID    int64      `json:"material_id"`
	MaterialCode  string     `json:"material_code" gorm:"size:50"`
	MaterialName  string     `json:"material_name" gorm:"size:100"`
	Quantity     float64    `json:"quantity" gorm:"type:decimal(18,4)"`
	Status       int        `json:"status" gorm:"default:1"`
}

// SchedulePlan 排程计划
type SchedulePlan struct {
	BaseModel
	TenantID     int64      `json:"tenant_id" gorm:"index;not null"`
	PlanNo       string     `json:"plan_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_schedule"`
	PlanType     string     `json:"plan_type" gorm:"size:20"` // 粗排/细排
	StartDate    *time.Time `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
	Algorithm    string     `json:"algorithm" gorm:"size:20"` // 遗传/粒子群/启发式
	Status       int        `json:"status" gorm:"default:1"` // 1待排程/2排程中/3已完成
	Remark       *string    `json:"remark" gorm:"size:500"`
}

// ScheduleResult 排程结果
type ScheduleResult struct {
	BaseModel
	PlanID       int64      `json:"plan_id" gorm:"index"`
	OrderID      int64      `json:"order_id" gorm:"index"`
	OrderNo      string     `json:"order_no" gorm:"size:50"`
	Sequence     int        `json:"sequence"` // 排程顺序
	LineID       int64      `json:"line_id"`
	LineName     *string    `json:"line_name" gorm:"size:100"`
	StationID    *int64     `json:"station_id"`
	StationName  *string    `json:"station_name" gorm:"size:100"`
	PlanStartTime *time.Time `json:"plan_start_time"`
	PlanEndTime  *time.Time `json:"plan_end_time"`
	ActualStartTime *time.Time `json:"actual_start_time"`
	ActualEndTime *time.Time `json:"actual_end_time"`
	Status       int        `json:"status" gorm:"default:1"` // 1待执行/2执行中/3已完成
}

// Resource 资源
type Resource struct {
	BaseModel
	TenantID     int64   `json:"tenant_id" gorm:"index;not null"`
	ResourceCode string  `json:"resource_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_resource"`
	ResourceName string  `json:"resource_name" gorm:"size:100;not null"`
	ResourceType string  `json:"resource_type" gorm:"size:20"` // 设备/人员/模具
	WorkshopID   int64   `json:"workshop_id"`
	Capacity     float64 `json:"capacity" gorm:"type:decimal(18,2)"` // 产能
	Unit         string  `json:"unit" gorm:"size:20"` // 单位时间产能
	Efficiency   float64 `json:"efficiency" gorm:"default:100"` // 效率%
	Status       int     `json:"status" gorm:"default:1"`
}

// WorkCenter 工作中心
type WorkCenter struct {
	BaseModel
	TenantID      int64   `json:"tenant_id" gorm:"index;not null"`
	WorkCenterCode string `json:"work_center_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_wc"`
	WorkCenterName string `json:"work_center_name" gorm:"size:100;not null"`
	WorkshopID    int64   `json:"workshop_id"`
	Capacity      float64 `json:"capacity" gorm:"type:decimal(18,2)"`
	Status        int     `json:"status" gorm:"default:1"`
}
