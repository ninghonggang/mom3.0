package model

import (
	"time"
)

// MDM BOM 物料清单头表
type MdmBOM struct {
	BaseModel
	TenantID      int64      `json:"tenant_id" gorm:"index;not null"`
	BOMCode      string     `json:"bom_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_bom_code"`
	BOMName      string     `json:"bom_name" gorm:"size:200;not null"`
	MaterialID   int64      `json:"material_id" gorm:"not null"`
	MaterialCode string     `json:"material_code" gorm:"size:50"`
	MaterialName string     `json:"material_name" gorm:"size:100"`
	Version      string     `json:"version" gorm:"size:20"` // 版本
	Status       string     `json:"status" gorm:"size:20;default:'DRAFT'"` // DRAFT草稿/ACTIVE生效/EXPIRED失效
	EffDate      *time.Time `json:"eff_date" gorm:"type:date"` // 生效日期
	ExpDate      *time.Time `json:"exp_date" gorm:"type:date"` // 失效日期
	Remark       *string    `json:"remark" gorm:"size:500"`
}

func (MdmBOM) TableName() string {
	return "mdm_bom"
}

// MDM BOM 物料清单行表
type MdmBOMItem struct {
	BaseModel
	TenantID       int64   `json:"tenant_id" gorm:"index;not null"`
	BOMID         int64   `json:"bom_id" gorm:"index;not null"`
	LineNo        int     `json:"line_no" gorm:"default:0"` // 行号
	MaterialID    int64   `json:"material_id" gorm:"not null"`
	MaterialCode  string  `json:"material_code" gorm:"size:50"`
	MaterialName  string  `json:"material_name" gorm:"size:100"`
	Quantity      float64 `json:"quantity" gorm:"type:decimal(18,4);default:0"` // 用量
	Unit          string  `json:"unit" gorm:"size:20"`
	ScrapRate     float64 `json:"scrap_rate" gorm:"type:decimal(10,4);default:0"` // 损耗率%
	SubstituteGroup *string `json:"substitute_group" gorm:"size:50"` // 替代组
	IsAlternative int     `json:"is_alternative" gorm:"default:0"` // 是否替代料 0否 1是
}

func (MdmBOMItem) TableName() string {
	return "mdm_bom_item"
}
