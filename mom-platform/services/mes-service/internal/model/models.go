package model

import (
	"time"

	"gorm.io/gorm"
)

// --- Domain models (neutral, no proto dependency) ---

type OrderStatus int

const (
	OrderPending    OrderStatus = 1
	OrderInProgress OrderStatus = 2
	OrderCompleted  OrderStatus = 3
	OrderClosed     OrderStatus = 4
	OrderSuspended  OrderStatus = 5
)

type Priority int

const (
	PriorityNormal   Priority = 1
	PriorityUrgent   Priority = 2
	PriorityCritical Priority = 3
)

type ReportStatus int

const (
	ReportSubmitted ReportStatus = 1
	ReportConfirmed ReportStatus = 2
	ReportAudited   ReportStatus = 3
)

type ReportType int

const (
	ReportNormal        ReportType = 1
	ReportSupplementary ReportType = 2
	ReportExceptional   ReportType = 3
)

type DispatchStatus int

const (
	DispatchPending    DispatchStatus = 1
	DispatchDispatched DispatchStatus = 2
	DispatchInProgress DispatchStatus = 3
	DispatchCompleted  DispatchStatus = 4
	DispatchWithdrawn  DispatchStatus = 5
)

// ProductionOrder — 生产工单
type ProductionOrder struct {
	ID             int64          `gorm:"primaryKey" json:"id"`
	TenantID       int64          `gorm:"index;not null" json:"tenant_id"`
	OrderNo        string         `gorm:"uniqueIndex;size:50;not null" json:"order_no"`
	SalesOrderNo   string         `gorm:"size:50" json:"sales_order_no"`
	MaterialID     int64          `gorm:"index;not null" json:"material_id"`
	MaterialCode   string         `gorm:"size:50;not null" json:"material_code"`
	MaterialName   string         `gorm:"size:100;not null" json:"material_name"`
	MaterialSpec   string         `gorm:"size:200" json:"material_spec"`
	Quantity       float64        `gorm:"not null" json:"quantity"`
	CompletedQty   float64        `gorm:"default:0" json:"completed_qty"`
	RejectedQty    float64        `gorm:"default:0" json:"rejected_qty"`
	WorkshopID     int64          `gorm:"index" json:"workshop_id"`
	WorkshopName   string         `gorm:"size:100" json:"workshop_name"`
	LineID         int64          `gorm:"index" json:"line_id"`
	LineName       string         `gorm:"size:100" json:"line_name"`
	Status         OrderStatus    `gorm:"default:1;not null" json:"status"`
	Priority       Priority       `gorm:"default:1;not null" json:"priority"`
	PlanStartDate  *time.Time     `json:"plan_start_date"`
	PlanEndDate    *time.Time     `json:"plan_end_date"`
	ActualStartDate *time.Time    `json:"actual_start_date"`
	ActualEndDate  *time.Time     `json:"actual_end_date"`
	Remark         string         `gorm:"size:500" json:"remark"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

func (ProductionOrder) TableName() string { return "production_orders" }

// MobileJobReport — 报工记录
type MobileJobReport struct {
	ID              int64          `gorm:"primaryKey" json:"id"`
	TenantID        int64          `gorm:"index;not null" json:"tenant_id"`
	OrderID         int64          `gorm:"index;not null" json:"order_id"`
	OrderNo         string         `gorm:"size:50;not null" json:"order_no"`
	ProcessID       int64          `gorm:"index" json:"process_id"`
	OperationID     int64          `gorm:"index" json:"operation_id"`
	OperationName   string         `gorm:"size:100" json:"operation_name"`
	EmployeeID      int64          `gorm:"index;not null" json:"employee_id"`
	EmployeeName    string         `gorm:"size:50;not null" json:"employee_name"`
	WorkstationID   int64          `gorm:"index" json:"workstation_id"`
	ReportedQty     float64        `gorm:"not null" json:"reported_qty"`
	QualifiedQty    float64        `gorm:"default:0" json:"qualified_qty"`
	DefectiveQty    float64        `gorm:"default:0" json:"defective_qty"`
	WorkMinutes     int            `gorm:"default:0" json:"work_minutes"`
	ReportType      ReportType     `gorm:"default:1" json:"report_type"`
	Status          ReportStatus   `gorm:"default:1" json:"status"`
	StartTime       *time.Time     `json:"start_time"`
	EndTime         *time.Time     `json:"end_time"`
	DefectCodes     string         `gorm:"size:500" json:"defect_codes"`
	Remark          string         `gorm:"size:500" json:"remark"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

func (MobileJobReport) TableName() string { return "mobile_job_report" }

// Dispatch — 派工单
type Dispatch struct {
	ID            int64          `gorm:"primaryKey" json:"id"`
	TenantID      int64          `gorm:"index;not null" json:"tenant_id"`
	OrderID       int64          `gorm:"index;not null" json:"order_id"`
	OrderNo       string         `gorm:"size:50" json:"order_no"`
	LineID        int64          `gorm:"index" json:"line_id"`
	WorkstationID int64          `gorm:"index" json:"workstation_id"`
	ProcessID     int64          `gorm:"index" json:"process_id"`
	OperationID   int64          `gorm:"index" json:"operation_id"`
	OperationName string         `gorm:"size:100" json:"operation_name"`
	EmployeeID    int64          `json:"employee_id"`
	EmployeeName  string         `gorm:"size:50" json:"employee_name"`
	PlannedQty    float64        `gorm:"default:0" json:"planned_qty"`
	CompletedQty  float64        `gorm:"default:0" json:"completed_qty"`
	PlanStartTime *time.Time     `json:"plan_start_time"`
	PlanEndTime   *time.Time     `json:"plan_end_time"`
	Status        DispatchStatus `gorm:"default:1" json:"status"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Dispatch) TableName() string { return "dispatch" }

// ProductionComplete — 生产完工入库记录
// 记录一次工单完工的入库动作：数量、批次、目标仓库/库位与完工时间。
type ProductionComplete struct {
	ID           int64          `gorm:"primaryKey" json:"id"`
	TenantID     int64          `gorm:"index;not null" json:"tenant_id"`
	OrderID      int64          `gorm:"index;not null" json:"order_id"`
	OrderNo      string         `gorm:"size:50" json:"order_no"`
	WarehouseID  int64          `gorm:"index" json:"warehouse_id"`
	LocationID   int64          `gorm:"index" json:"location_id"`
	Quantity     float64        `gorm:"not null" json:"quantity"`
	BatchNo      string         `gorm:"size:50" json:"batch_no"`
	CompleteTime *time.Time     `json:"complete_time"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

func (ProductionComplete) TableName() string { return "production_complete" }
