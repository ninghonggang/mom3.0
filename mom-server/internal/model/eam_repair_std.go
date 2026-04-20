package model

// EamRepairStd 维修标准
type EamRepairStd struct {
	BaseModel
	TenantID          int64   `gorm:"index;not null" json:"tenant_id"`
	StdCode           string  `gorm:"size:50;uniqueIndex:idx_eam_repair_std_code" json:"std_code"`
	StdName           string  `gorm:"size:100" json:"std_name"`
	FaultType         string  `gorm:"size:30" json:"fault_type"`
	RepairSteps       string  `gorm:"type:text" json:"repair_steps"`          // JSON格式维修步骤
	ToolsRequired     string  `gorm:"size:500" json:"tools_required"`
	MaterialsRequired string  `gorm:"size:500" json:"materials_required"`
	StandardHours     float64 `json:"standard_hours"`
	Status            string  `gorm:"size:20;default:'ACTIVE'" json:"status"`
	CreatedBy         int64   `json:"created_by"`
	UpdatedBy         int64   `json:"updated_by"`
}

func (EamRepairStd) TableName() string {
	return "eam_repair_std"
}

// EamRepairStdCreateReq 创建维修标准请求
type EamRepairStdCreateReq struct {
	StdCode           string  `json:"std_code" binding:"required"`
	StdName           string  `json:"std_name" binding:"required"`
	FaultType         string  `json:"fault_type"`
	RepairSteps       string  `json:"repair_steps"`
	ToolsRequired     string  `json:"tools_required"`
	MaterialsRequired string  `json:"materials_required"`
	StandardHours     float64 `json:"standard_hours"`
	Status            string  `json:"status"`
}

// EamRepairStdUpdateReq 更新维修标准请求
type EamRepairStdUpdateReq struct {
	ID                int64   `json:"id" binding:"required"`
	StdName           string  `json:"std_name"`
	FaultType         string  `json:"fault_type"`
	RepairSteps       string  `json:"repair_steps"`
	ToolsRequired     string  `json:"tools_required"`
	MaterialsRequired string  `json:"materials_required"`
	StandardHours     float64 `json:"standard_hours"`
	Status            string  `json:"status"`
}

// EamRepairStdPageReq 分页查询请求
type EamRepairStdPageReq struct {
	StdCode   string `form:"std_code"`
	StdName   string `form:"std_name"`
	FaultType string `form:"fault_type"`
	Status    string `form:"status"`
	Page      int    `form:"page"`
	PageSize  int    `form:"page_size"`
}
