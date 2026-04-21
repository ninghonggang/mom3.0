package model

import "time"

// MesJobReportLog 报工记录
type MesJobReportLog struct {
	ID            uint64    `gorm:"primaryKey" json:"id"`
	TenantID      int64     `gorm:"index" json:"tenantId"`
	WorkOrderId   uint64    `gorm:"index" json:"workOrderId"`
	WorkOrderCode string    `gorm:"size:50" json:"workOrderCode"`
	ProcessCode   string    `gorm:"size:50" json:"processCode"`
	ProcessName   string    `gorm:"size:100" json:"processName"`
	ReportType    string    `gorm:"size:20" json:"reportType"`   // NORMAL/PRODUCTION/SAMPLE
	Quantity      float64   `gorm:"type:decimal(18,2)" json:"quantity"`
	ReportTime    time.Time `gorm:"index" json:"reportTime"`
	ReporterId    uint64    `gorm:"index" json:"reporterId"`
	ReporterName  string    `gorm:"size:100" json:"reporterName"`
	Remark        string    `gorm:"size:500" json:"remark"`
	Status        string    `gorm:"size:20" json:"status"`        // PENDING/APPROVED/REJECTED
	CreatedAt     time.Time `json:"createdAt"`
}

func (MesJobReportLog) TableName() string {
	return "mes_job_report_log"
}

// MesJobReportLogCreateReqVO 创建报工记录请求
type MesJobReportLogCreateReqVO struct {
	WorkOrderId   uint64  `json:"workOrderId" binding:"required"`
	WorkOrderCode string  `json:"workOrderCode"`
	ProcessCode   string  `json:"processCode"`
	ProcessName   string  `json:"processName"`
	ReportType    string  `json:"reportType"`    // NORMAL/PRODUCTION/SAMPLE
	Quantity      float64 `json:"quantity" binding:"required"`
	ReportTime    string  `json:"reportTime"`    // YYYY-MM-DD HH:mm:ss
	ReporterId    uint64  `json:"reporterId"`
	ReporterName  string  `json:"reporterName"`
	Remark        string  `json:"remark"`
	Status        string  `json:"status"`
}

// MesJobReportLogQueryVO 查询报工记录请求
type MesJobReportLogQueryVO struct {
	Page        int    `json:"page" form:"page"`
	PageSize    int    `json:"page_size" form:"page_size"`
	Keyword     string `json:"keyword" form:"keyword"`       // search work_order_code, process_name
	WorkOrderId int64  `json:"workOrderId" form:"workOrderId"`
	ProcessCode string `json:"processCode" form:"processCode"`
	ReporterId  int64  `json:"reporterId" form:"reporterId"`
	StartDate   string `json:"startDate" form:"startDate"`
	EndDate     string `json:"endDate" form:"endDate"`
}

// MesJobReportLogSeniorReqVO 高级搜索请求
type MesJobReportLogSeniorReqVO struct {
	Conditions []map[string]interface{} `json:"conditions"`
	Page       int                       `json:"page"`
	PageSize   int                       `json:"page_size"`
}
