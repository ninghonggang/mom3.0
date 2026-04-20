package model

import (
	"time"
)

// ScpSupplierContact 供应商联系人表
type ScpSupplierContact struct {
	ID           uint      `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	TenantID     int64     `json:"tenant_id" gorm:"index;not null"`
	SupplierID   int64     `json:"supplier_id" gorm:"index;not null"`   // 供应商ID
	SupplierCode string    `json:"supplier_code" gorm:"size:50"`        // 供应商编码
	SupplierName string    `json:"supplier_name" gorm:"size:100"`       // 供应商名称
	Name         string    `json:"name" gorm:"size:50;not null"`        // 联系人姓名
	Gender       string    `json:"gender" gorm:"size:10"`               // 性别: MALE/FEMALE
	Department   string    `json:"department" gorm:"size:50"`           // 部门
	Position     string    `json:"position" gorm:"size:50"`             // 职位
	Phone        string    `json:"phone" gorm:"size:20"`               // 办公电话
	Mobile       string    `json:"mobile" gorm:"size:20"`              // 手机
	Email        string    `json:"email" gorm:"size:100"`              // 邮箱
	Wechat       string    `json:"wechat" gorm:"size:50"`              // 微信
	QQ           string    `json:"qq" gorm:"size:20"`                 // QQ
	IsPrimary    bool      `json:"is_primary" gorm:"default:false"`    // 是否主要联系人
	IsActive     bool      `json:"is_active" gorm:"default:true"`      // 是否启用
	Remark       string    `json:"remark" gorm:"type:text"`            // 备注
}

func (ScpSupplierContact) TableName() string {
	return "scp_supplier_contact"
}

// ScpSupplierContactCreateReqVO 创建供应商联系人请求
type ScpSupplierContactCreateReqVO struct {
	SupplierID   int64  `json:"supplierId" binding:"required"`
	SupplierCode string `json:"supplierCode"`
	SupplierName string `json:"supplierName"`
	Name         string `json:"name" binding:"required"`
	Gender       string `json:"gender"`
	Department   string `json:"department"`
	Position     string `json:"position"`
	Phone        string `json:"phone"`
	Mobile       string `json:"mobile"`
	Email        string `json:"email"`
	Wechat       string `json:"wechat"`
	QQ           string `json:"qq"`
	IsPrimary    bool   `json:"isPrimary"`
	Remark       string `json:"remark"`
}

// ScpSupplierContactUpdateReqVO 更新供应商联系人请求
type ScpSupplierContactUpdateReqVO struct {
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Department string `json:"department"`
	Position   string `json:"position"`
	Phone      string `json:"phone"`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
	Wechat     string `json:"wechat"`
	QQ         string `json:"qq"`
	IsPrimary  bool   `json:"isPrimary"`
	IsActive   bool   `json:"isActive"`
	Remark     string `json:"remark"`
}
