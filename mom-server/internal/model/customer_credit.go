package model

import (
	"time"
)

// CustomerCredit 客户信用
type CustomerCredit struct {
	ID              int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID        int64     `json:"tenant_id" gorm:"index;not null;default:1"`
	CustomerID      int64     `json:"customer_id" gorm:"index;not null"`   // 客户ID
	CustomerCode    string    `json:"customer_code" gorm:"size:64"`         // 客户编码
	CustomerName    string    `json:"customer_name" gorm:"size:200"`        // 客户名称
	CreditLimit     float64   `json:"credit_limit"`                         // 信用额度
	UsedCredit      float64   `json:"used_credit"`                          // 已用额度
	AvailableCredit float64   `json:"available_credit"`                     // 可用额度
	CreditLevel     string    `json:"credit_level" gorm:"size:20"`          // 信用等级：A/B/C/D
	PaymentDays     int       `json:"payment_days"`                          // 账期天数
	RiskLevel       int       `json:"risk_level" gorm:"default:1"`           // 风险等级：1低/2中/3高
	AlertThreshold  float64   `json:"alert_threshold"`                      // 预警阈值比例
	TotalOrders     int       `json:"total_orders"`                         // 累计订单数
	TotalAmount     float64   `json:"total_amount"`                         // 累计交易额
	OverdueAmount   float64   `json:"overdue_amount"`                        // 逾期金额
	Blacklist       int       `json:"blacklist" gorm:"default:0"`           // 黑名单：0否 1是
	Status          int       `json:"status" gorm:"default:1"`              // 状态：1正常 0冻结
	Remarks         string    `json:"remarks" gorm:"type:text"`             // 备注
	LastTradeDate   time.Time `json:"last_trade_date"`                      // 最后交易日期
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (CustomerCredit) TableName() string {
	return "customer_credit"
}

// CustomerCreditQuery 客户信用查询
type CustomerCreditQuery struct {
	CustomerCode string `form:"customer_code"`
	CustomerName string `form:"customer_name"`
	CreditLevel  string `form:"credit_level"`
	RiskLevel    int    `form:"risk_level"`
	Blacklist    int    `form:"blacklist"`
	Status       int    `form:"status"`
	Page         int    `form:"page"`
	PageSize     int    `form:"page_size"`
}

// CustomerCreditCreateReq 创建客户信用请求
type CustomerCreditCreateReq struct {
	CustomerID     int64   `json:"customer_id" binding:"required"`
	CustomerCode   string  `json:"customer_code"`
	CustomerName   string  `json:"customer_name"`
	CreditLimit    float64 `json:"credit_limit" binding:"required"`
	CreditLevel    string  `json:"credit_level"`
	PaymentDays    int     `json:"payment_days"`
	RiskLevel      int     `json:"risk_level"`
	AlertThreshold float64 `json:"alert_threshold"`
	Remarks        string  `json:"remarks"`
}