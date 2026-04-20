package model

import "time"

// DeliveryAddress 客户收货地址
type DeliveryAddress struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID     uint64    `gorm:"index;not null;default:1" json:"tenant_id"`
	CustomerID    uint64    `gorm:"index;not null" json:"customer_id"`
	AddressName   string    `gorm:"size:100" json:"address_name"`     // 地址别名
	ContactPerson string    `gorm:"size:50" json:"contact_person"`   // 联系人
	ContactPhone  string    `gorm:"size:50" json:"contact_phone"`    // 联系电话
	Province      string    `gorm:"size:50" json:"province"`         // 省
	City          string    `gorm:"size:50" json:"city"`             // 市
	District      string    `gorm:"size:50" json:"district"`         // 区
	AddressDetail string    `gorm:"type:text" json:"address_detail"` // 详细地址
	IsDefault     bool      `gorm:"default:false" json:"is_default"` // 是否默认地址
	IsActive      bool      `gorm:"default:true" json:"is_active"`   // 是否启用
	CreatedAt     time.Time `json:"created_at"`
}

func (DeliveryAddress) TableName() string {
	return "mdm_customer_delivery_address"
}

// DeliveryAddressCreateRequest 创建收货地址请求
type DeliveryAddressCreateRequest struct {
	CustomerID    uint64 `json:"customer_id" binding:"required"`
	AddressName   string `json:"address_name"`
	ContactPerson string `json:"contact_person"`
	ContactPhone  string `json:"contact_phone"`
	Province      string `json:"province"`
	City          string `json:"city"`
	District      string `json:"district"`
	AddressDetail string `json:"address_detail"`
	IsDefault     bool   `json:"is_default"`
	IsActive      bool   `json:"is_active"`
}

// DeliveryAddressUpdateRequest 更新收货地址请求
type DeliveryAddressUpdateRequest struct {
	AddressName   string `json:"address_name"`
	ContactPerson string `json:"contact_person"`
	ContactPhone  string `json:"contact_phone"`
	Province      string `json:"province"`
	City          string `json:"city"`
	District      string `json:"district"`
	AddressDetail string `json:"address_detail"`
	IsDefault     bool   `json:"is_default"`
	IsActive      bool   `json:"is_active"`
}

// DeliveryAddressQuery 查询收货地址
type DeliveryAddressQuery struct {
	CustomerID uint64 `json:"customer_id"`
}
