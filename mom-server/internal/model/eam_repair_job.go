package model

import (
	"time"
)

// EamRepairJob 维修工单
type EamRepairJob struct {
	BaseModel
	TenantID        int64      `gorm:"index;not null" json:"tenant_id"`
	JobCode         string     `gorm:"size:50;uniqueIndex:idx_eam_repair_job_code" json:"job_code"`
	EquipmentID     int64      `gorm:"index" json:"equipment_id"`
	EquipmentCode   string     `gorm:"size:50" json:"equipment_code"`
	FaultType       string     `gorm:"size:30" json:"fault_type"`
	FaultReason     string     `gorm:"size:30" json:"fault_reason"`
	FaultDesc       string     `gorm:"size:500" json:"fault_desc"`
	Level           string     `gorm:"size:20" json:"level"`           // URGENT/HIGH/NORMAL/LOW
	Status          string     `gorm:"size:20;default:'PENDING'" json:"status"` // PENDING/ASSIGNED/IN_PROGRESS/COMPLETED/CANCELLED
	ReporterID      int64      `json:"reporter_id"`
	ReporterName    string     `gorm:"size:100" json:"reporter_name"`
	AssigneeID      int64      `json:"assignee_id"`
	AssigneeName    string     `gorm:"size:100" json:"assignee_name"`
	PlanStartTime   *time.Time `json:"plan_start_time"`
	PlanEndTime     *time.Time `json:"plan_end_time"`
	ActualStartTime *time.Time `json:"actual_start_time"`
	ActualEndTime   *time.Time `json:"actual_end_time"`
	Result          string     `gorm:"size:500" json:"result"`
	Evaluation      string     `gorm:"size:20" json:"evaluation"`
	CreatedBy       int64      `json:"created_by"`
	UpdatedBy       int64      `json:"updated_by"`
}

func (EamRepairJob) TableName() string {
	return "eam_repair_job"
}

// EamRepairJobCreateReq 创建维修工单请求
type EamRepairJobCreateReq struct {
	JobCode       string `json:"job_code" binding:"required"`
	EquipmentID   int64  `json:"equipment_id"`
	EquipmentCode string `json:"equipment_code"`
	FaultType     string `json:"fault_type"`
	FaultReason   string `json:"fault_reason"`
	FaultDesc     string `json:"fault_desc"`
	Level         string `json:"level"`
	ReporterID    int64  `json:"reporter_id"`
	ReporterName  string `json:"reporter_name"`
	PlanStartTime string `json:"plan_start_time"`
	PlanEndTime   string `json:"plan_end_time"`
}

// EamRepairJobUpdateReq 更新维修工单请求
type EamRepairJobUpdateReq struct {
	ID            int64  `json:"id" binding:"required"`
	FaultType     string `json:"fault_type"`
	FaultReason   string `json:"fault_reason"`
	FaultDesc     string `json:"fault_desc"`
	Level         string `json:"level"`
	PlanStartTime string `json:"plan_start_time"`
	PlanEndTime   string `json:"plan_end_time"`
}

// EamRepairJobAssignReq 派工请求
type EamRepairJobAssignReq struct {
	ID           int64  `json:"id" binding:"required"`
	AssigneeID   int64  `json:"assignee_id" binding:"required"`
	AssigneeName string `json:"assignee_name"`
}

// EamRepairJobCompleteReq 完工请求
type EamRepairJobCompleteReq struct {
	ID     int64  `json:"id" binding:"required"`
	Result string `json:"result"`
}

// EamRepairJobEvaluateReq 评价请求
type EamRepairJobEvaluateReq struct {
	ID         int64  `json:"id" binding:"required"`
	Evaluation string `json:"evaluation" binding:"required"`
}

// EamRepairJobPageReq 分页查询请求
type EamRepairJobPageReq struct {
	JobCode     string `form:"job_code"`
	Status      string `form:"status"`
	EquipmentID int64  `form:"equipment_id"`
	Page        int    `form:"page"`
	PageSize    int    `form:"page_size"`
}
