package model

// EquipmentBom 设备BOM表
type EquipmentBom struct {
	BaseModel
	TenantID      int64   `json:"tenant_id" gorm:"index;not null"`
	EquipmentID   int64   `json:"equipment_id" gorm:"not null"`     // 设备ID
	EquipmentCode string  `json:"equipment_code" gorm:"size:50"`     // 设备编码
	EquipmentName string  `json:"equipment_name" gorm:"size:100"` // 设备名称
	MaterialID    int64   `json:"material_id" gorm:"not null"`      // 物料ID
	MaterialCode  string  `json:"material_code" gorm:"size:50"`      // 物料编码
	MaterialName  string  `json:"material_name" gorm:"size:100"`    // 物料名称
	Quantity       float64 `json:"quantity" gorm:"type:decimal(18,4);default:1"` // 标准用量
	Unit           string  `json:"unit" gorm:"size:20"`              // 单位
	Position       *string `json:"position" gorm:"size:100"`         // 安装位置
	ReplaceCycle   int     `json:"replace_cycle" gorm:"default:0"`   // 更换周期(天)
	IsCritical     int     `json:"is_critical" gorm:"default:0"`      // 是否关键备件 0否 1是
	Remark         *string `json:"remark" gorm:"size:500"`           // 备注
	Status         int     `json:"status" gorm:"default:1"`         // 1启用 0禁用
}

func (EquipmentBom) TableName() string {
	return "eqp_bom"
}

// EquipmentBomCreateReq 设备BOM创建请求
type EquipmentBomCreateReq struct {
	EquipmentID   int64   `json:"equipment_id" binding:"required"`
	EquipmentCode string  `json:"equipment_code"`
	EquipmentName string  `json:"equipment_name"`
	MaterialID    int64   `json:"material_id" binding:"required"`
	MaterialCode   string  `json:"material_code"`
	MaterialName   string  `json:"material_name"`
	Quantity       float64 `json:"quantity"`
	Unit           string  `json:"unit"`
	Position       *string `json:"position"`
	ReplaceCycle   int     `json:"replace_cycle"`
	IsCritical     int     `json:"is_critical"`
	Remark         *string `json:"remark"`
	Status         int     `json:"status"`
}

// EquipmentBomQuery 设备BOM查询请求
type EquipmentBomQuery struct {
	EquipmentID int64  `json:"equipment_id" form:"equipment_id"`
	MaterialCode string `json:"material_code" form:"material_code"`
	MaterialName string `json:"material_name" form:"material_name"`
	IsCritical  int    `json:"is_critical" form:"is_critical"`
	Page        int    `json:"page" form:"page"`
	PageSize    int    `json:"page_size" form:"page_size"`
}