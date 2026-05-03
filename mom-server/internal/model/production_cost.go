package model

import (
	"time"
)

// ProductionCost 生产成本归集表
type ProductionCost struct {
	BaseModel
	TenantID      int64      `json:"tenant_id" gorm:"index;not null"`
	OrderID       int64      `json:"order_id" gorm:"index;not null"`         // 工单ID
	OrderNo       string     `json:"order_no" gorm:"size:50"`                // 工单号
	CostType      string     `json:"cost_type" gorm:"size:20"`               // 成本类型: material/labor/overhead
	CostItem      string     `json:"cost_item" gorm:"size:100"`              // 成本项目
	Quantity      float64    `json:"quantity" gorm:"type:decimal(18,4)"`      // 数量
	UnitPrice     float64    `json:"unit_price" gorm:"type:decimal(18,4)"`   // 单价
	Amount        float64    `json:"amount" gorm:"type:decimal(18,2)"`       // 金额
	ProcessID     *int64     `json:"process_id"`                            // 工序ID
	ProcessName   *string    `json:"process_name" gorm:"size:100"`           // 工序名称
	DepartmentID  *int64     `json:"department_id"`                          // 部门ID
	DepartmentName *string   `json:"department_name" gorm:"size:100"`        // 部门名称
	WorkerID      *int64     `json:"worker_id"`                              // 人员ID
	WorkerName    *string    `json:"worker_name" gorm:"size:50"`             // 人员姓名
	EquipmentID   *int64     `json:"equipment_id"`                           // 设备ID
	EquipmentName *string    `json:"equipment_name" gorm:"size:100"`         // 设备名称
	CostDate      *time.Time `json:"cost_date"`                             // 成本日期
	Remark        *string    `json:"remark" gorm:"size:500"`                // 备注
}

func (ProductionCost) TableName() string {
	return "pro_production_cost"
}

// ProductionCostCreateReq 成本创建请求
type ProductionCostCreateReq struct {
	OrderID        int64   `json:"order_id" binding:"required"`
	OrderNo        string  `json:"order_no"`
	CostType       string  `json:"cost_type" binding:"required"` // material/labor/overhead
	CostItem       string  `json:"cost_item" binding:"required"`
	Quantity       float64 `json:"quantity"`
	UnitPrice      float64 `json:"unit_price"`
	Amount         float64 `json:"amount" binding:"required"`
	ProcessID      *int64  `json:"process_id"`
	ProcessName    *string `json:"process_name"`
	DepartmentID   *int64  `json:"department_id"`
	DepartmentName *string `json:"department_name"`
	WorkerID       *int64  `json:"worker_id"`
	WorkerName     *string `json:"worker_name"`
	EquipmentID    *int64  `json:"equipment_id"`
	EquipmentName  *string `json:"equipment_name"`
	CostDate       *string `json:"cost_date"`
	Remark         *string `json:"remark"`
}

// ProductionCostQuery 成本查询请求
type ProductionCostQuery struct {
	OrderID   int64  `json:"order_id" form:"order_id"`
	OrderNo   string `json:"order_no" form:"order_no"`
	CostType  string `json:"cost_type" form:"cost_type"`
	StartDate string `json:"start_date" form:"start_date"`
	EndDate   string `json:"end_date" form:"end_date"`
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"page_size" form:"page_size"`
}

// ProductionCostSummary 成本汇总
type ProductionCostSummary struct {
	OrderID        int64   `json:"order_id"`
	OrderNo        string  `json:"order_no"`
	MaterialCost   float64 `json:"material_cost"`   // 材料成本
	LaborCost      float64 `json:"labor_cost"`      // 人工成本
	OverheadCost   float64 `json:"overhead_cost"`   // 制造费用
	TotalCost      float64 `json:"total_cost"`      // 总成本
	UnitCost       float64 `json:"unit_cost"`       // 单位成本
	CompletedQty   float64 `json:"completed_qty"`   // 已完成数量
}