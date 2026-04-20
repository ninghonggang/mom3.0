package model

import "time"

// Contact 联系人
type Contact struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID  uint64    `gorm:"index;not null;default:1" json:"tenant_id"`
	OwnerType string    `gorm:"size:20;index" json:"owner_type"` // CUSTOMER/SUPPLIER
	OwnerID   uint64    `gorm:"index" json:"owner_id"`
	Name      string    `gorm:"size:50" json:"name"`
	Gender    string    `gorm:"size:10" json:"gender"` // MALE/FEMALE
	Phone     string    `gorm:"size:20" json:"phone"`
	Mobile    string    `gorm:"size:20" json:"mobile"`
	Email     string    `gorm:"size:100" json:"email"`
	Department string   `gorm:"size:50" json:"department"`
	Position  string    `gorm:"size:50" json:"position"`
	IsPrimary bool      `gorm:"default:false" json:"is_primary"`
	Remark    string    `gorm:"size:500" json:"remark"`
	CreatedBy string    `gorm:"size:50" json:"created_by"`
	UpdatedBy string    `gorm:"size:50" json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Contact) TableName() string {
	return "mdm_contact"
}

// ContactCreateRequest 创建联系人请求
type ContactCreateRequest struct {
	OwnerType  string `json:"owner_type" binding:"required"`
	OwnerID    uint64 `json:"owner_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
	Department string `json:"department"`
	Position   string `json:"position"`
	IsPrimary  bool   `json:"is_primary"`
	Remark     string `json:"remark"`
}

// ContactUpdateRequest 更新联系人请求
type ContactUpdateRequest struct {
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
	Department string `json:"department"`
	Position   string `json:"position"`
	IsPrimary  bool   `json:"is_primary"`
	Remark     string `json:"remark"`
}

// ContactQuery 查询联系人
type ContactQuery struct {
	OwnerType string `json:"owner_type"`
	OwnerID   uint64 `json:"owner_id"`
}

// BankAccount 银行账户
type BankAccount struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID    uint64    `gorm:"index;not null;default:1" json:"tenant_id"`
	OwnerType   string    `gorm:"size:20;index" json:"owner_type"` // CUSTOMER/SUPPLIER
	OwnerID     uint64    `gorm:"index" json:"owner_id"`
	BankName    string    `gorm:"size:100" json:"bank_name"`      // 开户银行
	BankAccount string    `gorm:"size:50" json:"bank_account"`    // 银行账号
	AccountName string    `gorm:"size:100" json:"account_name"`  // 账户名称
	BranchName  string    `gorm:"size:100" json:"branch_name"`   // 支行名称
	Currency    string    `gorm:"size:10" json:"currency"`       // 币种默认CNY
	IsPrimary   bool      `gorm:"default:false" json:"is_primary"`
	Status      string    `gorm:"size:20" json:"status"`         // ACTIVE/INACTIVE
	Remark      string    `gorm:"size:500" json:"remark"`
	CreatedBy   string    `gorm:"size:50" json:"created_by"`
	UpdatedBy   string    `gorm:"size:50" json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (BankAccount) TableName() string {
	return "mdm_bank_account"
}

// BankAccountCreateRequest 创建银行账户请求
type BankAccountCreateRequest struct {
	OwnerType  string `json:"owner_type" binding:"required"`
	OwnerID    uint64 `json:"owner_id" binding:"required"`
	BankName   string `json:"bank_name" binding:"required"`
	BankAccount string `json:"bank_account" binding:"required"`
	AccountName string `json:"account_name" binding:"required"`
	BranchName string `json:"branch_name"`
	Currency   string `json:"currency"`
	IsPrimary  bool   `json:"is_primary"`
	Status     string `json:"status"`
	Remark     string `json:"remark"`
}

// BankAccountUpdateRequest 更新银行账户请求
type BankAccountUpdateRequest struct {
	BankName    string `json:"bank_name"`
	BankAccount string `json:"bank_account"`
	AccountName string `json:"account_name"`
	BranchName  string `json:"branch_name"`
	Currency    string `json:"currency"`
	IsPrimary   bool   `json:"is_primary"`
	Status      string `json:"status"`
	Remark      string `json:"remark"`
}

// BankAccountQuery 查询银行账户
type BankAccountQuery struct {
	OwnerType string `json:"owner_type"`
	OwnerID   uint64 `json:"owner_id"`
}

// Attachment 附件
type Attachment struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID    uint64    `gorm:"index;not null;default:1" json:"tenant_id"`
	OwnerType   string    `gorm:"size:20;index" json:"owner_type"` // CUSTOMER/SUPPLIER/PRODUCT/ORDER等
	OwnerID     uint64    `gorm:"index" json:"owner_id"`
	FileName    string    `gorm:"size:200" json:"file_name"`     // 原始文件名
	FilePath    string    `gorm:"size:500" json:"file_path"`     // 存储路径
	FileSize    int64     `json:"file_size"`                    // 文件大小
	FileType    string    `gorm:"size:50" json:"file_type"`     // 文件类型MIME
	Category    string    `gorm:"size:50" json:"category"`      // 附件分类LICENSE/CERTIFICATE/CONTRACT/OTHER
	Description string    `gorm:"size:500" json:"description"`
	CreatedBy   string    `gorm:"size:50" json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

func (Attachment) TableName() string {
	return "mdm_attachment"
}

// AttachmentCreateRequest 创建附件请求
type AttachmentCreateRequest struct {
	OwnerType   string `json:"owner_type" binding:"required"`
	OwnerID     uint64 `json:"owner_id" binding:"required"`
	FileName    string `json:"file_name" binding:"required"`
	FilePath    string `json:"file_path" binding:"required"`
	FileSize    int64  `json:"file_size"`
	FileType    string `json:"file_type"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

// AttachmentUpdateRequest 更新附件请求
type AttachmentUpdateRequest struct {
	FileName    string `json:"file_name"`
	FilePath    string `json:"file_path"`
	FileSize    int64  `json:"file_size"`
	FileType    string `json:"file_type"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

// AttachmentQuery 查询附件
type AttachmentQuery struct {
	OwnerType string `json:"owner_type"`
	OwnerID   uint64 `json:"owner_id"`
	Category  string `json:"category"`
}
