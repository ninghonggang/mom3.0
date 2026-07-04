package model

import (
	"time"
)

// MobileJobReport 移动端报工
type MobileJobReport struct {
	ID              int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID        int64     `json:"tenant_id" gorm:"index;not null;default:1"`
	WorkshopID      int64     `json:"workshop_id" gorm:"index"`              // 车间ID
	WorkshopName    string    `json:"workshop_name" gorm:"size:100"`         // 车间名称
	ProductionLineID int64    `json:"production_line_id" gorm:"index"`      // 产线ID
	ProductionLineName string  `json:"production_line_name" gorm:"size:100"` // 产线名称
	WorkstationID   int64     `json:"workstation_id" gorm:"index"`          // 工位ID
	WorkstationName string    `json:"workstation_name" gorm:"size:100"`       // 工位名称
	OrderID         int64     `json:"order_id" gorm:"index"`                // 工单ID
	OrderCode       string    `json:"order_code" gorm:"size:64"`             // 工单编号
	ProcessID       int64     `json:"process_id" gorm:"index"`              // 工序ID
	ProcessName     string    `json:"process_name" gorm:"size:100"`         // 工序名称
	EmployeeID      int64     `json:"employee_id" gorm:"index"`             // 员工ID
	EmployeeName    string    `json:"employee_name" gorm:"size:64"`          // 员工名称
	ReportedQuantity float64   `json:"reported_quantity"`                    // 报工数量
	QualifiedQuantity float64  `json:"qualified_quantity"`                    // 合格数量
	DefectiveQuantity float64 `json:"defective_quantity"`                    // 不良数量
	WorkMinutes    int        `json:"work_minutes"`                          // 工作时长(分钟)
	StartTime      time.Time  `json:"start_time"`                           // 开始时间
	EndTime        time.Time  `json:"end_time"`                             // 结束时间
	ReportType     int        `json:"report_type" gorm:"default:1"`         // 报工类型：1正常 2补报 3异常
	Status         int        `json:"status" gorm:"default:1"`              // 状态：1已提交 2已确认 3已审核
	StatusV2       string     `json:"status_v2" gorm:"size:30;index"`       // 双轨:varchar - MOM 3.0 V2.1
	Remarks        string     `json:"remarks" gorm:"type:text"`             // 备注
	ConfirmBy      int64      `json:"confirm_by"`                           // 确认人
	ConfirmAt      *time.Time `json:"confirm_at"`                           // 确认时间
	AuditBy        int64      `json:"audit_by"`                             // 审核人
	AuditAt        *time.Time `json:"audit_at"`                             // 审核时间
	CreatedAt      time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

// StatusCode 返回有效状态码(优先 status_v2)
func (o *MobileJobReport) StatusCode() string {
	if o.StatusV2 != "" {
		return o.StatusV2
	}
	return map[int]string{1: "SUBMITTED", 2: "CONFIRMED", 3: "AUDITED", 4: "REJECTED"}[o.Status]
}

func (MobileJobReport) TableName() string {
	return "mobile_job_report"
}

// MobileJobReportQuery 移动端报工查询
type MobileJobReportQuery struct {
	WorkshopID      int64  `form:"workshop_id"`
	ProductionLineID int64 `form:"production_line_id"`
	WorkstationID   int64  `form:"workstation_id"`
	OrderID         int64  `form:"order_id"`
	EmployeeID      int64  `form:"employee_id"`
	ReportType      int    `form:"report_type"`
	Status          int    `form:"status"`
	StartDate       string `form:"start_date"`
	EndDate         string `form:"end_date"`
	Page            int    `form:"page"`
	PageSize        int    `form:"page_size"`
}

// MobileJobReportCreateReq 创建报工请求
type MobileJobReportCreateReq struct {
	WorkshopID      int64   `json:"workshop_id" binding:"required"`
	ProductionLineID int64  `json:"production_line_id" binding:"required"`
	WorkstationID   int64   `json:"workstation_id" binding:"required"`
	OrderID         int64   `json:"order_id" binding:"required"`
	ProcessID       int64   `json:"process_id"`
	EmployeeID      int64   `json:"employee_id" binding:"required"`
	ReportedQuantity float64 `json:"reported_quantity" binding:"required"`
	QualifiedQuantity float64 `json:"qualified_quantity"`
	DefectiveQuantity float64 `json:"defective_quantity"`
	WorkMinutes    int      `json:"work_minutes"`
	StartTime      string   `json:"start_time"`
	EndTime        string   `json:"end_time"`
	ReportType     int      `json:"report_type"`
	Remarks        string   `json:"remarks"`
}

// MobilePendingOrder 待报工工单
type MobilePendingOrder struct {
	OrderID         int64   `json:"order_id"`
	OrderCode       string  `json:"order_code"`
	ProductCode     string  `json:"product_code"`
	ProductName     string  `json:"product_name"`
	ProcessName     string  `json:"process_name"`
	ProcessID       int64   `json:"process_id"`
	WorkstationID   int64   `json:"workstation_id"`
	WorkstationName string  `json:"workstation_name"`
	PlannedQuantity float64 `json:"planned_quantity"`
	ReportedQuantity float64 `json:"reported_quantity"`
	RemainingQuantity float64 `json:"remaining_quantity"`
	Priority        int     `json:"priority"`
	DueDate         string  `json:"due_date"`
}