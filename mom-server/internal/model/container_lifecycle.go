package model

import (
	"encoding/json"
	"time"
)

// ContainerEventType 容器事件类型
type ContainerEventType string

const (
	ContainerEventInitialize   ContainerEventType = "INITIALIZE"   // 初始化
	ContainerEventInService    ContainerEventType = "IN_SERVICE"   // 在役
	ContainerEventMaintenance  ContainerEventType = "MAINTENANCE"  // 维修
	ContainerEventRepair      ContainerEventType = "REPAIR"       // 修理
	ContainerEventRetire      ContainerEventType = "RETIRE"      // 退役
	ContainerEventScrap       ContainerEventType = "SCRAP"       // 报废
)

// ContainerLifecycleStatusChange 容器状态变更
type ContainerLifecycleStatusChange string

const (
	LifecycleStatusBefore ContainerLifecycleStatusChange = "BEFORE_STATUS" // 变更前状态
	LifecycleStatusAfter  ContainerLifecycleStatusChange = "AFTER_STATUS"  // 变更后状态
)

// ContainerMaintenanceType 容器维修类型
type ContainerMaintenanceType string

const (
	MaintenanceTypeCleaning     ContainerMaintenanceType = "CLEANING"     // 清洁
	MaintenanceTypeRepair       ContainerMaintenanceType = "REPAIR"       // 修理
	MaintenanceTypeReplacement  ContainerMaintenanceType = "REPLACEMENT"  // 更换
)

// ContainerMaintenanceStatus 容器维修状态
type ContainerMaintenanceStatus string

const (
	MaintenanceStatusPending    ContainerMaintenanceStatus = "PENDING"     // 待处理
	MaintenanceStatusInProgress ContainerMaintenanceStatus = "IN_PROGRESS" // 处理中
	MaintenanceStatusCompleted  ContainerMaintenanceStatus = "COMPLETED"   // 已完成
)

// ContainerLifecycle 容器生命周期记录表
type ContainerLifecycle struct {
	ID            uint                 `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time            `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time            `json:"updated_at" gorm:"autoUpdateTime"`
	TenantID     int64                `json:"tenant_id" gorm:"index;not null;default:1"`
	ContainerID   int64                `json:"container_id" gorm:"index;not null"` // 容器ID
	ContainerCode string               `json:"container_code" gorm:"size:50"`      // 容器编码
	EventType     ContainerEventType   `json:"event_type" gorm:"size:20;not null"` // 事件类型
	EventDate     time.Time            `json:"event_date" gorm:"not null"`        // 事件日期
	OperatorID    int64                `json:"operator_id" gorm:"index"`          // 操作人ID
	OperatorName  string               `json:"operator_name" gorm:"size:50"`      // 操作人姓名
	LocationID    int64                `json:"location_id" gorm:"index"`          // 位置ID
	LocationName  string               `json:"location_name" gorm:"size:100"`     // 位置名称
	Status        ContainerLifecycleStatusChange `json:"status" gorm:"size:20"` // 状态
	Remark        string               `json:"remark" gorm:"size:500"`          // 备注
}

func (ContainerLifecycle) TableName() string {
	return "mes_container_lifecycle"
}

// ContainerLifecycleQuery 容器生命周期查询
type ContainerLifecycleQuery struct {
	Page         int                 `form:"page" json:"page"`
	PageSize     int                 `form:"pageSize" json:"pageSize"`
	ContainerID  int64               `form:"container_id" json:"container_id"`
	EventType    ContainerEventType  `form:"event_type" json:"event_type"`
	StartDate    *time.Time          `form:"start_date" json:"start_date"`
	EndDate      *time.Time          `form:"end_date" json:"end_date"`
}

// ContainerMaintenance 容器维修记录表
type ContainerMaintenance struct {
	ID                uint                      `json:"id" gorm:"primaryKey"`
	CreatedAt         time.Time                 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time                 `json:"updated_at" gorm:"autoUpdateTime"`
	TenantID         int64                     `json:"tenant_id" gorm:"index;not null;default:1"`
	ContainerID       int64                     `json:"container_id" gorm:"index;not null"`           // 容器ID
	ContainerCode     string                    `json:"container_code" gorm:"size:50"`              // 容器编码
	MaintenanceType   ContainerMaintenanceType   `json:"maintenance_type" gorm:"size:20;not null"`  // 维修类型
	MaintenanceDate   time.Time                 `json:"maintenance_date" gorm:"not null"`          // 维修日期
	CompletedDate     *time.Time                `json:"completed_date"`                           // 完成日期
	TechnicianID     int64                     `json:"technician_id" gorm:"index"`               // 维修技师ID
	TechnicianName   string                    `json:"technician_name" gorm:"size:50"`            // 维修技师姓名
	FaultDescription string                    `json:"fault_description" gorm:"size:500"`          // 故障描述
	MaintenanceContent string                  `json:"maintenance_content" gorm:"size:500"`       // 维修内容
	SparePartsUsed   json.RawMessage           `json:"spare_parts_used" gorm:"type:jsonb"`         // 使用的备件JSON
	Cost             float64                   `json:"cost" gorm:"type:decimal(18,2)"`            // 维修费用
	Status           ContainerMaintenanceStatus `json:"status" gorm:"size:20;default:'PENDING'"` // 状态
	Remark           string                    `json:"remark" gorm:"size:500"`                   // 备注
}

func (ContainerMaintenance) TableName() string {
	return "mes_container_maintenance"
}

// ContainerMaintenanceQuery 容器维修记录查询
type ContainerMaintenanceQuery struct {
	Page          int                        `form:"page" json:"page"`
	PageSize      int                        `form:"pageSize" json:"pageSize"`
	ContainerID   int64                      `form:"container_id" json:"container_id"`
	ContainerCode string                      `form:"container_code" json:"container_code"`
	MaintenanceType ContainerMaintenanceType  `form:"maintenance_type" json:"maintenance_type"`
	Status        ContainerMaintenanceStatus  `form:"status" json:"status"`
	StartDate     *time.Time                 `form:"start_date" json:"start_date"`
	EndDate       *time.Time                 `form:"end_date" json:"end_date"`
}

// InitializeContainerRequest 初始化容器请求
type InitializeContainerRequest struct {
	ContainerID  int64   `json:"container_id" binding:"required"`
	OperatorID   int64   `json:"operator_id"`
	OperatorName string  `json:"operator_name" binding:"required"`
	LocationID   int64   `json:"location_id"`
	LocationName string  `json:"location_name"`
	Remark       string  `json:"remark"`
}

// RecordMaintenanceRequest 记录维修请求
type RecordMaintenanceRequest struct {
	ContainerID       int64                    `json:"container_id" binding:"required"`
	MaintenanceType   ContainerMaintenanceType  `json:"maintenance_type" binding:"required"`
	MaintenanceDate   time.Time                `json:"maintenance_date" binding:"required"`
	TechnicianID      int64                    `json:"technician_id"`
	TechnicianName    string                   `json:"technician_name"`
	FaultDescription  string                   `json:"fault_description"`
	MaintenanceContent string                  `json:"maintenance_content"`
	SparePartsUsed    json.RawMessage          `json:"spare_parts_used"`
	Cost              float64                  `json:"cost"`
	Remark            string                   `json:"remark"`
}

// CompleteMaintenanceRequest 完成维修请求
type CompleteMaintenanceRequest struct {
	CompletedDate    time.Time       `json:"completed_date" binding:"required"`
	MaintenanceContent string        `json:"maintenance_content"`
	SparePartsUsed   json.RawMessage `json:"spare_parts_used"`
	Cost             float64         `json:"cost"`
	Remark           string          `json:"remark"`
}

// RetireContainerRequest 报废容器请求
type RetireContainerRequest struct {
	ContainerID  int64   `json:"container_id" binding:"required"`
	OperatorID   int64   `json:"operator_id"`
	OperatorName string  `json:"operator_name" binding:"required"`
	LocationID   int64   `json:"location_id"`
	LocationName string  `json:"location_name"`
	Remark       string  `json:"remark"`
}

// ContainerTimelineItem 容器时间线项
type ContainerTimelineItem struct {
	EventType   string    `json:"event_type"`
	EventDate   time.Time `json:"event_date"`
	OperatorName string    `json:"operator_name"`
	LocationName string   `json:"location_name"`
	Status      string    `json:"status"`
	Remark      string    `json:"remark"`
}