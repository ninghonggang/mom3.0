package model

import (
	"time"
)

// QualityCertificate 质量证书
type QualityCertificate struct {
	ID           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID     int64     `json:"tenant_id" gorm:"index;not null;default:1"`
	CertCode     string    `json:"cert_code" gorm:"size:64;uniqueIndex"` // 证书编号
	CertType     string    `json:"cert_type" gorm:"size:32"`             // 证书类型：COC,质检报告
	OrderID      int64     `json:"order_id" gorm:"index"`               // 关联订单ID
	OrderCode    string    `json:"order_code" gorm:"size:64"`            // 订单编号
	ProductID    int64     `json:"product_id" gorm:"index"`             // 产品ID
	ProductCode  string    `json:"product_code" gorm:"size:64"`          // 产品编码
	ProductName  string    `json:"product_name" gorm:"size:200"`        // 产品名称
	BatchNo      string    `json:"batch_no" gorm:"size:64"`             // 批次号
	Quantity     float64   `json:"quantity"`                            // 数量
	Unit         string    `json:"unit" gorm:"size:20"`                 // 单位
	Inspector    string    `json:"inspector" gorm:"size:64"`            // 检验员
	InspectDate  time.Time `json:"inspect_date"`                        // 检验日期
	Result       int       `json:"result" gorm:"default:1"`            // 检验结果：1合格 0不合格
	Status       int       `json:"status" gorm:"default:1"`             // 状态：1有效 0无效
	IssueDate    time.Time `json:"issue_date"`                          // 发证日期
	ExpiryDate   time.Time `json:"expiry_date"`                         // 有效期
	Remarks      string    `json:"remarks" gorm:"type:text"`            // 备注
	Attachments  string    `json:"attachments" gorm:"type:text"`         // 附件JSON
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (QualityCertificate) TableName() string {
	return "quality_certificate"
}

// QualityCertificateQuery 质量证书查询
type QualityCertificateQuery struct {
	OrderCode   string `form:"order_code"`
	ProductCode string `form:"product_code"`
	ProductName string `form:"product_name"`
	BatchNo     string `form:"batch_no"`
	CertType    string `form:"cert_type"`
	Result      int    `form:"result"`
	Status      int    `form:"status"`
	StartDate   string `form:"start_date"`
	EndDate     string `form:"end_date"`
	Page        int    `form:"page"`
	PageSize    int    `form:"page_size"`
}

// QualityCertificateCreateReq 创建质量证书请求
type QualityCertificateCreateReq struct {
	CertType    string  `json:"cert_type" binding:"required"`
	OrderID     int64   `json:"order_id"`
	OrderCode   string  `json:"order_code"`
	ProductID   int64   `json:"product_id" binding:"required"`
	ProductCode string  `json:"product_code"`
	ProductName string  `json:"product_name"`
	BatchNo     string  `json:"batch_no"`
	Quantity    float64 `json:"quantity"`
	Unit        string  `json:"unit"`
	Inspector   string  `json:"inspector"`
	InspectDate string  `json:"inspect_date"`
	Result      int     `json:"result"`
	IssueDate   string  `json:"issue_date"`
	ExpiryDate  string  `json:"expiry_date"`
	Remarks     string  `json:"remarks"`
	Attachments string  `json:"attachments"`
}