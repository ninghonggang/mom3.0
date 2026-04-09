package model

// Customer 客户
type Customer struct {
	BaseModel
	TenantID int64   `json:"tenant_id" gorm:"index;not null"`
	Code     string  `json:"code" gorm:"size:50;not null;uniqueIndex:idx_tenant_customer_code"` // 客户编码
	Name     string  `json:"name" gorm:"size:200;not null"`                                   // 客户名称
	Type     *string `json:"type" gorm:"size:50"`                                            // 客户类型
	Contact  *string `json:"contact" gorm:"size:100"`                                        // 联系人
	Phone    *string `json:"phone" gorm:"size:50"`                                           // 联系电话
	Email    *string `json:"email" gorm:"size:100"`                                          // 邮箱
	Address  *string `json:"address" gorm:"size:500"`                                        // 地址
	Status   int     `json:"status" gorm:"default:1"`                                       // 状态: 1=启用/2=禁用
}

func (Customer) TableName() string {
	return "mdm_customer"
}
