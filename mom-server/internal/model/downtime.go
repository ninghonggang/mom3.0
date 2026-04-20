package model

import (
	"time"
)

// EquipmentDowntime 设备停机记录
type EquipmentDowntime struct {
	ID              uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID        int64      `gorm:"index;not null;default:1" json:"tenant_id"`
	EquipmentID     int64      `gorm:"index;not null" json:"equipment_id"`      // 设备ID
	EquipmentCode   string     `gorm:"size:50" json:"equipment_code"`          // 设备编号
	EquipmentName   string     `gorm:"size:100" json:"equipment_name"`          // 设备名称
	DowntimeType    string     `gorm:"size:20;index" json:"downtime_type"`      // DOWNFAULT/DOWNPLAN/DOWNMAINT
	DowntimeReason  string     `gorm:"size:200" json:"downtime_reason"`         // 停机原因
	StartTime       time.Time  `gorm:"index" json:"start_time"`                // 停机开始时间
	EndTime         *time.Time `json:"end_time"`                               // 停机结束时间
	Duration        int        `json:"duration"`                               // 停机时长(分钟)
	LostProduction  float64    `json:"lost_production"`                        // 损失产量
	WorkOrderID     int64      `gorm:"index" json:"work_order_id"`             // 工单ID
	WorkOrderCode   string     `gorm:"size:50" json:"work_order_code"`          // 工单编号
	ShiftID         int64      `json:"shift_id"`                               // 班组ID
	OperatorID      int64      `json:"operator_id"`                            // 操作员ID
	OperatorName    string     `gorm:"size:50" json:"operator_name"`           // 操作员姓名
	MaintainerID    int64      `json:"maintainer_id"`                          // 维修人员ID
	MaintainerName  string     `gorm:"size:50" json:"maintainer_name"`         // 维修人员姓名
	Status          string     `gorm:"size:20;index" json:"status"`            // OPEN/INPROGRESS/CLOSED
	Remark          string     `gorm:"size:500" json:"remark"`                  // 备注
	CreatedBy       string     `gorm:"size:50" json:"created_by"`
	UpdatedBy       string     `gorm:"size:50" json:"updated_by"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

func (EquipmentDowntime) TableName() string {
	return "eam_equipment_downtime"
}

// EquipmentDowntimeCreateRequest 设备停机创建请求
type EquipmentDowntimeCreateRequest struct {
	EquipmentID    int64   `json:"equipment_id" binding:"required"`
	EquipmentCode  string  `json:"equipment_code"`
	EquipmentName  string  `json:"equipment_name"`
	DowntimeType   string  `json:"downtime_type" binding:"required"`
	DowntimeReason string  `json:"downtime_reason"`
	StartTime      string  `json:"start_time" binding:"required"`
	EndTime        *string `json:"end_time"`
	Duration       int     `json:"duration"`
	LostProduction float64 `json:"lost_production"`
	WorkOrderID    int64   `json:"work_order_id"`
	WorkOrderCode  string  `json:"work_order_code"`
	ShiftID        int64   `json:"shift_id"`
	OperatorID     int64   `json:"operator_id"`
	OperatorName   string  `json:"operator_name"`
	MaintainerID   int64   `json:"maintainer_id"`
	MaintainerName string  `json:"maintainer_name"`
	Remark         string  `json:"remark"`
}

// EquipmentDowntimeUpdateRequest 设备停机更新请求
type EquipmentDowntimeUpdateRequest struct {
	DowntimeType   string  `json:"downtime_type"`
	DowntimeReason string  `json:"downtime_reason"`
	StartTime      string  `json:"start_time"`
	EndTime        *string `json:"end_time"`
	Duration       int     `json:"duration"`
	LostProduction float64 `json:"lost_production"`
	WorkOrderID    int64   `json:"work_order_id"`
	WorkOrderCode  string  `json:"work_order_code"`
	ShiftID        int64   `json:"shift_id"`
	OperatorID     int64   `json:"operator_id"`
	OperatorName   string  `json:"operator_name"`
	MaintainerID   int64   `json:"maintainer_id"`
	MaintainerName string  `json:"maintainer_name"`
	Status         string  `json:"status"`
	Remark         string  `json:"remark"`
}

// EquipmentDowntimeQuery 设备停机查询参数
type EquipmentDowntimeQuery struct {
	EquipmentID    int64  `json:"equipment_id"`
	EquipmentCode string `json:"equipment_code"`
	DowntimeType string `json:"downtime_type"`
	Status       string `json:"status"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Page         int    `json:"page"`
	PageSize     int    `json:"page_size"`
}
