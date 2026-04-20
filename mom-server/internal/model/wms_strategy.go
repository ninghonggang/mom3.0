package model

import (
	"time"
)

// WmsStrategy 策略配置
type WmsStrategy struct {
	ID           uint64    `json:"id" gorm:"primaryKey"`
	TenantID     int64     `json:"tenant_id" gorm:"index"`
	StrategyCode string    `json:"strategy_code" gorm:"size:50;uniqueIndex:idx_tenant_code"`
	StrategyName string    `json:"strategy_name" gorm:"size:100"`
	StrategyType string    `json:"strategy_type" gorm:"size:20"`  // PICK/PUTAWAY/TRANSFER
	RuleContent  string    `json:"rule_content" gorm:"type:text"`  // JSON格式规则
	Priority     int       `json:"priority" gorm:"default:0"`
	Status       string    `json:"status" gorm:"size:20"`
	CreatedBy    uint64    `json:"created_by"`
	UpdatedBy    uint64    `json:"updated_by"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (WmsStrategy) TableName() string {
	return "wms_strategy"
}

// WmsStrategyCreateReqVO 策略配置创建请求
type WmsStrategyCreateReqVO struct {
	StrategyCode string `json:"strategy_code"`
	StrategyName string `json:"strategy_name"`
	StrategyType string `json:"strategy_type"`
	RuleContent  string `json:"rule_content"`
	Priority     int    `json:"priority"`
	Status       string `json:"status"`
}

// WmsStrategyUpdateReqVO 策略配置更新请求
type WmsStrategyUpdateReqVO struct {
	Id           uint64 `json:"id"`
	StrategyName string `json:"strategy_name"`
	StrategyType string `json:"strategy_type"`
	RuleContent  string `json:"rule_content"`
	Priority     int    `json:"priority"`
	Status       string `json:"status"`
}

// WmsStrategyQueryVO 策略配置查询请求
type WmsStrategyQueryVO struct {
	Keyword      string `form:"keyword"`
	StrategyType string `form:"strategy_type"`
	Status       string `form:"status"`
	Page         int    `form:"page"`
	PageSize     int    `form:"page_size"`
}
