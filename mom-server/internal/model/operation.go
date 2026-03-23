package model

// MDM 工序表
type MdmOperation struct {
	BaseModel
	TenantID          int64   `json:"tenant_id" gorm:"index;not null"`
	OperationCode     string  `json:"operation_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_op_code"`
	OperationName     string  `json:"operation_name" gorm:"size:100;not null"`
	WorkcenterID      *int64  `json:"workcenter_id"`
	WorkcenterName    *string `json:"workcenter_name" gorm:"size:100"`
	StandardWorktime  int     `json:"standard_worktime" gorm:"default:0"` // 标准工时(分钟)
	QualityStd       *string `json:"quality_std" gorm:"size:500"` // 质量标准
	IsKeyProcess      int     `json:"is_key_process" gorm:"default:0"` // 是否关键工序 0否 1是
	IsQCPoint        int     `json:"is_qc_point" gorm:"default:0"` // 是否质检点 0否 1是
	Sequence         int     `json:"sequence" gorm:"default:0"` // 顺序号
	Remark           *string `json:"remark" gorm:"size:500"`
}

func (MdmOperation) TableName() string {
	return "mdm_operation"
}
