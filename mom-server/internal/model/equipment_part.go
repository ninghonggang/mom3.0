package model

// EquipmentPart 设备部件
type EquipmentPart struct {
	BaseModel
	TenantID       int64   `json:"tenant_id" gorm:"index;not null"`
	EquipmentID    int64   `json:"equipment_id" gorm:"index;not null"`
	EquipmentCode  string  `json:"equipment_code" gorm:"size:50"`
	EquipmentName  string  `json:"equipment_name" gorm:"size:100"`
	PartCode       string  `json:"part_code" gorm:"size:50;not null"`     // 部件编号
	PartName       string  `json:"part_name" gorm:"size:100;not null"`    // 部件名称
	Spec           *string `json:"spec" gorm:"size:100"`                 // 规格型号
	Unit           string  `json:"unit" gorm:"size:20"`                   // 单位
	Qty            float64 `json:"qty" gorm:"type:decimal(18,4)"`          // 数量
	Supplier       *string `json:"supplier" gorm:"size:100"`              // 供应商
	UnitPrice      float64 `json:"unit_price" gorm:"type:decimal(18,2)"`  // 单价
	TotalPrice     float64 `json:"total_price" gorm:"type:decimal(18,2)"` // 总价
	ReplacementFreq int     `json:"replacement_freq"`                      // 更换频率(小时)
	MaxStock       float64 `json:"max_stock" gorm:"type:decimal(18,4)"`   // 最大库存
	MinStock       float64 `json:"min_stock" gorm:"type:decimal(18,4)"`   // 最小库存
	CurrentStock   float64 `json:"current_stock" gorm:"type:decimal(18,4)"` // 当前库存
	Status         int     `json:"status" gorm:"default:1"`                // 1正常/2停用
}

func (EquipmentPart) TableName() string {
	return "equ_equipment_part"
}

// EquipmentPartCreate 设备部件创建
type EquipmentPartCreate struct {
	EquipmentID    int64   `json:"equipment_id" binding:"required"`
	PartCode       string  `json:"part_code" binding:"required"`
	PartName       string  `json:"part_name" binding:"required"`
	Spec           *string `json:"spec"`
	Unit           string  `json:"unit"`
	Qty            float64 `json:"qty"`
	Supplier       *string `json:"supplier"`
	UnitPrice      float64 `json:"unit_price"`
	TotalPrice     float64 `json:"total_price"`
	ReplacementFreq int    `json:"replacement_freq"`
	MaxStock       float64 `json:"max_stock"`
	MinStock       float64 `json:"min_stock"`
	CurrentStock   float64 `json:"current_stock"`
}

type EquipmentPartUpdate struct {
	PartCode       *string `json:"part_code"`
	PartName       *string `json:"part_name"`
	Spec           *string `json:"spec"`
	Unit           *string `json:"unit"`
	Qty            *float64 `json:"qty"`
	Supplier       *string `json:"supplier"`
	UnitPrice      *float64 `json:"unit_price"`
	TotalPrice     *float64 `json:"total_price"`
	ReplacementFreq *int    `json:"replacement_freq"`
	MaxStock       *float64 `json:"max_stock"`
	MinStock       *float64 `json:"min_stock"`
	CurrentStock   *float64 `json:"current_stock"`
	Status         *int     `json:"status"`
}
