package model

// EamRepairFlow 维修流程配置
type EamRepairFlow struct {
	BaseModel
	TenantID  int64  `gorm:"index;not null" json:"tenant_id"`
	FlowCode  string `gorm:"size:50;uniqueIndex:idx_eam_repair_flow_code" json:"flow_code"`
	FlowName  string `gorm:"size:100" json:"flow_name"`
	FlowSteps string `gorm:"type:text" json:"flow_steps"` // JSON格式步骤
	Status    string `gorm:"size:20;default:'ACTIVE'" json:"status"`
	CreatedBy int64  `json:"created_by"`
	UpdatedBy int64  `json:"updated_by"`
}

func (EamRepairFlow) TableName() string {
	return "eam_repair_flow"
}

// EamRepairFlowCreateReq 创建维修流程请求
type EamRepairFlowCreateReq struct {
	FlowCode  string `json:"flow_code" binding:"required"`
	FlowName  string `json:"flow_name" binding:"required"`
	FlowSteps string `json:"flow_steps"`
	Status    string `json:"status"`
}

// EamRepairFlowUpdateReq 更新维修流程请求
type EamRepairFlowUpdateReq struct {
	ID        int64  `json:"id" binding:"required"`
	FlowName  string `json:"flow_name"`
	FlowSteps string `json:"flow_steps"`
	Status    string `json:"status"`
}

// EamRepairFlowPageReq 分页查询请求
type EamRepairFlowPageReq struct {
	FlowCode string `form:"flow_code"`
	FlowName string `form:"flow_name"`
	Status   string `form:"status"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}
