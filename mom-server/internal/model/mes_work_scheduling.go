package model

import "time"

// MesWorkScheduling 工单排程主表
type MesWorkScheduling struct {
	BaseModel
	TenantID       int64     `json:"tenant_id" gorm:"index;not null;uniqueIndex:idx_tenant_scheduling_code"`
	PlanNoDay      string    `json:"plan_no_day" gorm:"size:50"`
	SchedulingCode string    `json:"scheduling_code" gorm:"size:50;uniqueIndex:idx_tenant_scheduling_code"`
	ProductCode    string    `json:"product_code" gorm:"size:50"`
	ProductName    string    `json:"product_name" gorm:"size:200"`
	Status         string    `json:"status" gorm:"size:20"`
	Quantity       float64   `json:"quantity" gorm:"type:decimal(18,2)"`
	FinishedQty    float64   `json:"finished_qty" gorm:"type:decimal(18,2)"`
	WorkMode       string    `json:"work_mode" gorm:"size:20"`
	TaskMode       string    `json:"task_mode" gorm:"size:20"`
	PlanDate       time.Time `json:"plan_date" gorm:"type:date"`
	WorkshopCode   string    `json:"workshop_code" gorm:"size:50"`
	LineCode       string    `json:"line_code" gorm:"size:50"`
	CreatedBy      int64     `json:"created_by"`
	UpdatedBy      int64     `json:"updated_by"`

	Details []MesWorkSchedulingDetail `json:"details,omitempty" gorm:"foreignKey:SchedulingID"`
}

func (MesWorkScheduling) TableName() string {
	return "mes_work_scheduling"
}

// MesWorkSchedulingCreateReqVO 创建工单排程请求
type MesWorkSchedulingCreateReqVO struct {
	PlanNoDay      string  `json:"plan_no_day"`
	SchedulingCode string  `json:"scheduling_code"`
	ProductCode    string  `json:"product_code"`
	ProductName    string  `json:"product_name"`
	Status         string  `json:"status"`
	Quantity       float64 `json:"quantity"`
	WorkMode       string  `json:"work_mode"`
	TaskMode       string  `json:"task_mode"`
	PlanDate       string  `json:"plan_date"` // YYYY-MM-DD
	WorkshopCode   string  `json:"workshop_code"`
	LineCode       string  `json:"line_code"`
}

// WorkScheduleUpdateVO 更新工单排程请求
type WorkScheduleUpdateVO struct {
	ID             int64   `json:"id" binding:"required"`
	PlanNoDay      string  `json:"plan_no_day"`
	ProductCode    string  `json:"product_code"`
	ProductName    string  `json:"product_name"`
	Status         string  `json:"status"`
	Quantity       float64 `json:"quantity"`
	FinishedQty    float64 `json:"finished_qty"`
	WorkMode       string  `json:"work_mode"`
	TaskMode       string  `json:"task_mode"`
	PlanDate       string  `json:"plan_date"` // YYYY-MM-DD
	WorkshopCode   string  `json:"workshop_code"`
	LineCode       string  `json:"line_code"`
}

// WorkSchedulePageVO 分页查询请求
type WorkSchedulePageVO struct {
	Page           int    `json:"page" form:"page"`
	PageSize       int    `json:"page_size" form:"page_size"`
	SchedulingCode string `json:"scheduling_code" form:"scheduling_code"`
	ProductCode    string `json:"product_code" form:"product_code"`
	Status         string `json:"status" form:"status"`
	PlanDate       string `json:"plan_date" form:"plan_date"`
	WorkshopCode   string `json:"workshop_code" form:"workshop_code"`
	LineCode       string `json:"line_code" form:"line_code"`
}
