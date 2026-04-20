package model

// ========== EAM设备组织模块 ==========

// Factory 厂区
type Factory struct {
	BaseModel
	TenantID     int64   `json:"tenant_id" gorm:"index;not null"`
	FactoryCode  string  `json:"factory_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_factory_code"`
	FactoryName  string  `json:"factory_name" gorm:"size:100;not null"`
	Province     *string `json:"province" gorm:"size:50"`    // 省份
	City         *string `json:"city" gorm:"size:50"`        // 城市
	District     *string `json:"district" gorm:"size:50"`    // 区县
	Address      *string `json:"address" gorm:"size:200"`    // 详细地址
	Manager      *string `json:"manager" gorm:"size:50"`     // 负责人
	Phone        *string `json:"phone" gorm:"size:20"`       // 联系电话
	AreaSize     *float64 `json:"area_size" gorm:"type:decimal(18,2)"` // 占地面积(平方米)
	Status       int     `json:"status" gorm:"default:1"`    // 1启用 0禁用
}

func (Factory) TableName() string {
	return "eam_factory"
}

// EquipmentOrg 设备组织结构（厂区/车间/产线层级关系视图）
// 用于展示完整的层级关系：Factory -> Workshop -> Line
type EquipmentOrg struct {
	BaseModel
	TenantID     int64   `json:"tenant_id" gorm:"index;not null"`
	FactoryID    int64   `json:"factory_id" gorm:"not null"`    // 厂区ID
	FactoryCode  string  `json:"factory_code" gorm:"size:50"`  // 厂区编码
	FactoryName  string  `json:"factory_name" gorm:"size:100"` // 厂区名称
	WorkshopID   int64   `json:"workshop_id" gorm:"not null"`  // 车间ID
	WorkshopCode string  `json:"workshop_code" gorm:"size:50"` // 车间编码
	WorkshopName string  `json:"workshop_name" gorm:"size:100"` // 车间名称
	LineID       int64   `json:"line_id" gorm:"not null"`      // 产线ID
	LineCode     string  `json:"line_code" gorm:"size:50"`     // 产线编码
	LineName     string  `json:"line_name" gorm:"size:100"`    // 产线名称
	Status       int     `json:"status" gorm:"default:1"`       // 状态
}

func (EquipmentOrg) TableName() string {
	return "eam_equipment_org"
}

// EquipmentOrgNode 设备组织节点（用于树形结构）
type EquipmentOrgNode struct {
	ID           int64               `json:"id"`
	FactoryID    int64               `json:"factory_id"`
	FactoryCode  string              `json:"factory_code"`
	FactoryName  string              `json:"factory_name"`
	Status       int                 `json:"status"`
	Children     []WorkshopNode      `json:"children,omitempty"`
}

// WorkshopNode 车间节点
type WorkshopNode struct {
	ID         int64          `json:"id"`
	WorkshopID int64          `json:"workshop_id"`
	Code       string         `json:"code"`
	Name       string         `json:"name"`
	Status     int            `json:"status"`
	Children   []LineNode     `json:"children,omitempty"`
}

// LineNode 产线节点
type LineNode struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Status int  `json:"status"`
}

// EquipmentOrgQuery 设备组织查询条件
type EquipmentOrgQuery struct {
	FactoryID  int64 `json:"factory_id"`  // 厂区ID
	WorkshopID int64 `json:"workshop_id"` // 车间ID
	Status     int   `json:"status"`      // 状态
}
