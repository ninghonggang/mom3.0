package model

// DynamicRule 动态规则
type DynamicRule struct {
	BaseModel
	TenantID   int64  `json:"tenant_id" gorm:"index;not null"`
	RuleCode   string `json:"rule_code" gorm:"size:50;not null"`    // 规则编码
	RuleName   string `json:"rule_name" gorm:"size:100;not null"`   // 规则名称
	RuleType   string `json:"rule_type" gorm:"size:50"`              // 规则类型
	RuleConfig string `json:"rule_config" gorm:"type:text"`         // 规则配置(JSON)
	Priority   int    `json:"priority" gorm:"default:0"`            // 优先级
	Status     string `json:"status" gorm:"size:20;default:ACTIVE"` // 状态 ACTIVE INACTIVE
	Remark     string `json:"remark" gorm:"size:500"`               // 备注
}

func (DynamicRule) TableName() string {
	return "qc_dynamic_rule"
}
