package model

import (
	"time"
)

// ScpSupplierBank 供应商银行账户表
type ScpSupplierBank struct {
	ID           uint      `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	TenantID     int64     `json:"tenant_id" gorm:"index;not null"`
	SupplierID   int64     `json:"supplier_id" gorm:"index;not null"`   // 供应商ID
	SupplierCode string    `json:"supplier_code" gorm:"size:50"`        // 供应商编码
	SupplierName string    `json:"supplier_name" gorm:"size:100"`       // 供应商名称
	BankName     string    `json:"bank_name" gorm:"size:100;not null"`   // 开户银行
	BankCode     string    `json:"bank_code" gorm:"size:50"`            // 银行编码
	BranchName   string    `json:"branch_name" gorm:"size:100"`         // 支行名称
	BranchCode   string    `json:"branch_code" gorm:"size:50"`         // 支行编码
	AccountType  string    `json:"account_type" gorm:"size:20"`         // 账户类型: BASIC/GENERAL/VIRTUAL
	AccountNo    string    `json:"account_no" gorm:"size:50;not null"`  // 银行账号
	AccountName  string    `json:"account_name" gorm:"size:100;not null"`// 账户名称
	Currency     string    `json:"currency" gorm:"size:10;default:'CNY'"` // 币种
	IsPrimary    bool      `json:"is_primary" gorm:"default:false"`      // 是否主账户
	IsActive     bool      `json:"is_active" gorm:"default:true"`        // 是否启用
	Status       string    `json:"status" gorm:"size:20;default:'ACTIVE'"` // ACTIVE/INACTIVE
	Remark       string    `json:"remark" gorm:"type:text"`              // 备注
}

func (ScpSupplierBank) TableName() string {
	return "scp_supplier_bank"
}

// ScpSupplierBankCreateReqVO 创建供应商银行账户请求
type ScpSupplierBankCreateReqVO struct {
	SupplierID   int64  `json:"supplierId" binding:"required"`
	SupplierCode  string `json:"supplierCode"`
	SupplierName  string `json:"supplierName"`
	BankName      string `json:"bankName" binding:"required"`
	BankCode      string `json:"bankCode"`
	BranchName    string `json:"branchName"`
	BranchCode    string `json:"branchCode"`
	AccountType   string `json:"accountType"`
	AccountNo     string `json:"accountNo" binding:"required"`
	AccountName   string `json:"accountName" binding:"required"`
	Currency      string `json:"currency"`
	IsPrimary     bool   `json:"isPrimary"`
	Remark        string `json:"remark"`
}

// ScpSupplierBankUpdateReqVO 更新供应商银行账户请求
type ScpSupplierBankUpdateReqVO struct {
	BankName     string `json:"bankName"`
	BankCode     string `json:"bankCode"`
	BranchName   string `json:"branchName"`
	BranchCode   string `json:"branchCode"`
	AccountType  string `json:"accountType"`
	AccountNo    string `json:"accountNo"`
	AccountName  string `json:"accountName"`
	Currency     string `json:"currency"`
	IsPrimary    bool   `json:"isPrimary"`
	IsActive     bool   `json:"isActive"`
	Status       string `json:"status"`
	Remark       string `json:"remark"`
}
