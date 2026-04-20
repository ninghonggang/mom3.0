package model

import "time"

// MesWorkSchedulingDetail 工单排程明细表
type MesWorkSchedulingDetail struct {
	BaseModel
	TenantID        int64      `json:"tenant_id" gorm:"index"`
	SchedulingID    int64      `json:"scheduling_id" gorm:"index;not null"`
	WorkingNode     string     `json:"working_node" gorm:"size:50"`
	WorkingName     string     `json:"working_name" gorm:"size:100"`
	Status          string     `json:"status" gorm:"size:20"`
	EquipmentID     int64      `json:"equipment_id"`
	EquipmentCode   string     `json:"equipment_code" gorm:"size:50"`
	WorkstationID   int64      `json:"workstation_id"`
	WorkstationName string     `json:"workstation_name" gorm:"size:100"`
	WorkerID        int64      `json:"worker_id"`
	WorkerName      string     `json:"worker_name" gorm:"size:100"`
	StartTime       *time.Time `json:"start_time"`
	EndTime         *time.Time `json:"end_time"`
	PlanQty         float64    `json:"plan_qty" gorm:"type:decimal(18,2)"`
	FinishedQty     float64    `json:"finished_qty" gorm:"type:decimal(18,2)"`
	WorkMinutes     int        `json:"work_minutes"`
	CreatedBy       int64      `json:"created_by"`
	UpdatedBy       int64      `json:"updated_by"`
}

func (MesWorkSchedulingDetail) TableName() string {
	return "mes_work_scheduling_detail"
}

// MesWorkSchedulingDetailCreateReqVO 创建工序排程请求
type MesWorkSchedulingDetailCreateReqVO struct {
	SchedulingID    int64   `json:"scheduling_id" binding:"required"`
	WorkingNode     string  `json:"working_node"`
	WorkingName     string  `json:"working_name"`
	Status          string  `json:"status"`
	EquipmentID     int64   `json:"equipment_id"`
	EquipmentCode   string  `json:"equipment_code"`
	WorkstationID   int64   `json:"workstation_id"`
	WorkstationName string  `json:"workstation_name"`
	WorkerID        int64   `json:"worker_id"`
	WorkerName      string  `json:"worker_name"`
	PlanQty         float64 `json:"plan_qty"`
	WorkMinutes     int     `json:"work_minutes"`
}

// MesWorkSchedulingDetailUpdateReqVO 更新工序排程请求
type MesWorkSchedulingDetailUpdateReqVO struct {
	ID              int64   `json:"id" binding:"required"`
	WorkingNode     string  `json:"working_node"`
	WorkingName     string  `json:"working_name"`
	Status          string  `json:"status"`
	EquipmentID     int64   `json:"equipment_id"`
	EquipmentCode   string  `json:"equipment_code"`
	WorkstationID   int64   `json:"workstation_id"`
	WorkstationName string  `json:"workstation_name"`
	WorkerID        int64   `json:"worker_id"`
	WorkerName      string  `json:"worker_name"`
	PlanQty         float64 `json:"plan_qty"`
	FinishedQty     float64 `json:"finished_qty"`
	WorkMinutes     int     `json:"work_minutes"`
}

// MesWorkSchedulingDetailPageReqVO 分页查询请求
type MesWorkSchedulingDetailPageReqVO struct {
	Page         int    `json:"page" form:"page"`
	PageSize     int    `json:"page_size" form:"page_size"`
	SchedulingID int64  `json:"scheduling_id" form:"scheduling_id"`
	Status       string `json:"status" form:"status"`
	WorkingNode  string `json:"working_node" form:"working_node"`
}

// MesWorkSchedulingDetailReportReqVO 工序报工请求
type MesWorkSchedulingDetailReportReqVO struct {
	ID          int64   `json:"id" binding:"required"`
	FinishedQty float64 `json:"finished_qty" binding:"required"`
	Remark      string  `json:"remark"`
}

// MesWorkSchedulingDetailBindEquipmentReqVO 绑定设备请求
type MesWorkSchedulingDetailBindEquipmentReqVO struct {
	ID            int64  `json:"id" binding:"required"`
	EquipmentID   int64  `json:"equipment_id" binding:"required"`
	EquipmentCode string `json:"equipment_code"`
}

// MesWorkSchedulingDetailBindWorkerReqVO 绑定人员请求
type MesWorkSchedulingDetailBindWorkerReqVO struct {
	ID         int64  `json:"id" binding:"required"`
	WorkerID   int64  `json:"worker_id" binding:"required"`
	WorkerName string `json:"worker_name"`
}
