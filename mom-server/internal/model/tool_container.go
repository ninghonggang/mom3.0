package model

import (
	"time"
)

// Tool 器具
type Tool struct {
	ID            int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID      int64     `json:"tenant_id" gorm:"index;not null;default:1"`
	ToolCode      string    `json:"tool_code" gorm:"size:50;uniqueIndex:idx_tenant_tool"`
	ToolName      string    `json:"tool_name" gorm:"size:100"`
	ToolType      string    `json:"tool_type" gorm:"size:50"` // MOLD/Fixture/Gauge/OTHER
	Specification string    `json:"specification" gorm:"size:100"`
	Unit          string    `json:"unit" gorm:"size:20"`
	Status        string    `json:"status" gorm:"size:20;index"` // IDLE/INUSE/MAINTENANCE/SCRAPPED
	Location      string    `json:"location" gorm:"size:100"`
	Remark        string    `json:"remark" gorm:"size:500"`
	CreatedBy     string    `json:"created_by" gorm:"size:50"`
	UpdatedBy     string    `json:"updated_by" gorm:"size:50"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (Tool) TableName() string {
	return "mes_tool"
}

// ToolCreateRequest 器具创建请求
type ToolCreateRequest struct {
	ToolCode      string `json:"tool_code" binding:"required"`
	ToolName      string `json:"tool_name" binding:"required"`
	ToolType      string `json:"tool_type"`
	Specification string `json:"specification"`
	Unit          string `json:"unit"`
	Status        string `json:"status"`
	Location      string `json:"location"`
	Remark        string `json:"remark"`
}

// ToolUpdateRequest 器具更新请求
type ToolUpdateRequest struct {
	ToolName      string `json:"tool_name"`
	ToolType      string `json:"tool_type"`
	Specification string `json:"specification"`
	Unit          string `json:"unit"`
	Status        string `json:"status"`
	Location      string `json:"location"`
	Remark        string `json:"remark"`
}

// ToolQuery 器具查询
type ToolQuery struct {
	Page     int    `form:"page" json:"page"`
	PageSize int    `form:"pageSize" json:"pageSize"`
	Query    string `form:"query" json:"query"`
	ToolType string `form:"tool_type" json:"tool_type"`
	Status   string `form:"status" json:"status"`
}

// ToolContainer 容器
type ToolContainer struct {
	ID            int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID      int64     `json:"tenant_id" gorm:"index;not null;default:1"`
	ContainerCode string    `json:"container_code" gorm:"size:50;uniqueIndex:idx_tenant_tool_container"`
	ContainerName string    `json:"container_name" gorm:"size:100"`
	ContainerType string    `json:"container_type" gorm:"size:50"` // PALLET/BOX/TRAY/OTHER
	Capacity      float64   `json:"capacity"`
	Status        string    `json:"status" gorm:"size:20;index"` // EMPTY/FULL/INTRANSIT
	Location      string    `json:"location" gorm:"size:100"`
	Remark        string    `json:"remark" gorm:"size:500"`
	CreatedBy     string    `json:"created_by" gorm:"size:50"`
	UpdatedBy     string    `json:"updated_by" gorm:"size:50"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (ToolContainer) TableName() string {
	return "mes_tool_container"
}

// ToolContainerCreateRequest 容器创建请求
type ToolContainerCreateRequest struct {
	ContainerCode string  `json:"container_code" binding:"required"`
	ContainerName string  `json:"container_name" binding:"required"`
	ContainerType string  `json:"container_type"`
	Capacity      float64 `json:"capacity"`
	Status        string  `json:"status"`
	Location      string  `json:"location"`
	Remark        string  `json:"remark"`
}

// ToolContainerUpdateRequest 容器更新请求
type ToolContainerUpdateRequest struct {
	ContainerName string  `json:"container_name"`
	ContainerType string  `json:"container_type"`
	Capacity      float64 `json:"capacity"`
	Status        string  `json:"status"`
	Location      string  `json:"location"`
	Remark        string  `json:"remark"`
}

// ToolContainerQuery 容器查询
type ToolContainerQuery struct {
	Page          int    `form:"page" json:"page"`
	PageSize      int    `form:"pageSize" json:"pageSize"`
	Query         string `form:"query" json:"query"`
	ContainerType string `form:"container_type" json:"container_type"`
	Status        string `form:"status" json:"status"`
}

// ToolContainerBinding 器具容器绑定
type ToolContainerBinding struct {
	ID            int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID      int64      `json:"tenant_id" gorm:"index;not null;default:1"`
	ToolID        int64      `json:"tool_id" gorm:"index;not null"`
	ToolCode      string     `json:"tool_code" gorm:"size:50"`
	ContainerID   int64      `json:"container_id" gorm:"index;not null"`
	ContainerCode string     `json:"container_code" gorm:"size:50"`
	BindTime      time.Time  `json:"bind_time"`
	UnbindTime    *time.Time `json:"unbind_time"`
	Status        string     `json:"status" gorm:"size:20;index"` // BOUND/UNBOUND
	OperatorID    int64      `json:"operator_id"`
	OperatorName  string     `json:"operator_name" gorm:"size:50"`
	Remark        string     `json:"remark" gorm:"size:500"`
	CreatedBy     string     `json:"created_by" gorm:"size:50"`
	CreatedAt     time.Time  `json:"created_at"`
}

func (ToolContainerBinding) TableName() string {
	return "mes_tool_container_binding"
}

// ToolContainerBindingCreateRequest 绑定请求
type ToolContainerBindingCreateRequest struct {
	ToolID      int64  `json:"tool_id" binding:"required"`
	ContainerID int64  `json:"container_id" binding:"required"`
	OperatorID  int64  `json:"operator_id"`
	OperatorName string `json:"operator_name"`
	Remark      string `json:"remark"`
}

// ToolContainerBindingQuery 绑定查询
type ToolContainerBindingQuery struct {
	Page         int    `form:"page" json:"page"`
	PageSize     int    `form:"pageSize" json:"pageSize"`
	Query        string `form:"query" json:"query"`
	ToolID       int64  `form:"tool_id" json:"tool_id"`
	ContainerID  int64  `form:"container_id" json:"container_id"`
	Status       string `form:"status" json:"status"`
}
