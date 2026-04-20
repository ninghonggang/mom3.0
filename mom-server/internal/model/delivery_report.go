package model

import (
	"time"
)

// DeliveryReport 交付率报表
type DeliveryReport struct {
	BaseModel
	TenantID int64 `json:"tenant_id" gorm:"index;not null"`
	ReportMonth   time.Time `json:"report_month" gorm:"index;not null"`
	CustomerID    int64    `json:"customer_id" gorm:"index"`
	CustomerName  string   `json:"customer_name" gorm:"size:100"`

	OrderCount      int     `json:"order_count"`       // 订单数
	TotalOrderQty   float64 `json:"total_order_qty"`   // 订单总量
	DeliveredQty    float64 `json:"delivered_qty"`     // 已交付量
	OnTimeDeliverQty float64 `json:"on_time_deliver_qty"` // 准时交付量
	DeliveryRate    float64 `json:"delivery_rate"`     // 交付率%
	OnTimeRate      float64 `json:"on_time_rate"`      // 准时率%
	LateDeliverQty  float64 `json:"late_deliver_qty"`  // 延期交付量
	Remark          string  `json:"remark" gorm:"size:500"`
}

func (DeliveryReport) TableName() string {
	return "delivery_report"
}