package model

// MdmFactory 工厂表（多工厂支持）
type MdmFactory struct {
	BaseModel
	TenantID     int64   `json:"tenant_id" gorm:"index;not null"`
	FactoryCode  string  `json:"factory_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_factory_code"`
	FactoryName  string  `json:"factory_name" gorm:"size:100;not null"`
	Province     *string `json:"province" gorm:"size:50"`
	City         *string `json:"city" gorm:"size:50"`
	District     *string `json:"district" gorm:"size:50"`
	Address      *string `json:"address" gorm:"size:200"`
	Manager      *string `json:"manager" gorm:"size:50"`
	Phone        *string `json:"phone" gorm:"size:20"`
	AreaSize     *float64 `json:"area_size" gorm:"type:decimal(18,2)"`
	Status       int     `json:"status" gorm:"default:1"` // 1启用 0禁用
	IsDefault    int     `json:"is_default" gorm:"default:0"` // 是否默认工厂
}

func (MdmFactory) TableName() string {
	return "mdm_factory"
}

// FactoryCreateReq 工厂创建请求
type FactoryCreateReq struct {
	FactoryCode string  `json:"factory_code" binding:"required"`
	FactoryName string  `json:"factory_name" binding:"required"`
	Province    *string `json:"province"`
	City        *string `json:"city"`
	District    *string `json:"district"`
	Address     *string `json:"address"`
	Manager     *string `json:"manager"`
	Phone       *string `json:"phone"`
	AreaSize    *float64 `json:"area_size"`
	Status      int     `json:"status"`
	IsDefault   int     `json:"is_default"`
}

// FactoryQuery 工厂查询请求
type FactoryQuery struct {
	FactoryCode string `json:"factory_code" form:"factory_code"`
	FactoryName string `json:"factory_name" form:"factory_name"`
	Status     int    `json:"status" form:"status"`
	Page       int    `json:"page" form:"page"`
	PageSize   int    `json:"page_size" form:"page_size"`
}

// TenantFactory 租户工厂关联表（用于设置当前工厂）
type TenantFactory struct {
	BaseModel
	TenantID   int64 `json:"tenant_id" gorm:"index;not null"`
	FactoryID  int64 `json:"factory_id" gorm:"index;not null"`
	UserID     int64 `json:"user_id" gorm:"index;not null"` // 0表示租户级设置
	IsCurrent  int   `json:"is_current" gorm:"default:0"`    // 1当前工厂 0非当前
}

func (TenantFactory) TableName() string {
	return "sys_tenant_factory"
}