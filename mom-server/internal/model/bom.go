package model

import (
	"time"

	"mom-server/internal/pkg/status"
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
	Status       string     `json:"status" gorm:"size:20;default:'DRAFT'"` // DRAFT草稿/ACTIVE生效/EXPIRED失效(legacy,保留双轨)
	StatusV2     string     `json:"status_v2" gorm:"size:30;index"`       // 双轨:varchar - MOM 3.0 V2.1,与 mdm_status_dict 对齐
	EffDate      *time.Time `json:"eff_date" gorm:"type:date"` // 生效日期
	ExpDate      *time.Time `json:"exp_date" gorm:"type:date"` // 失效日期
	Remark       *string    `json:"remark" gorm:"size:500"`
	ErpBomCode   string     `json:"erp_bom_code" gorm:"size:50"`           // 金蝶BOM编码
	ErpSyncTime  *time.Time `json:"erp_sync_time"`                         // 同步时间
	ErpSyncStatus string     `json:"erp_sync_status" gorm:"size:20"`       // SYNCED/PENDING/FAILED
	IsCurrent    int        `json:"is_current" gorm:"default:1"`         // 是否当前版本 0否 1是
}

// StatusCode 返回有效状态码(优先 status_v2, fallback status → 字典码)
// 与 mes_process.StatusCode() 同模式,统一 MOM 3.0 切读路径
func (b *MdmBOM) StatusCode() string {
	if b.StatusV2 != "" {
		return b.StatusV2
	}
	return string(status.MDMBomFromLegacyVarchar(b.Status))
}

// SetStatus 同时写 status + status_v2(双轨制,与 mdm_status_dict 对齐)
// 入参:字典码(DRAFT / ACTIVE / OBSOLETE),非法码自动 fallback DRAFT
func (b *MdmBOM) SetStatus(code string) {
	c := status.Code(code)
	if !c.IsValid(status.MDMBomAll) {
		code = string(status.MDMBomDraft)
	}
	b.Status = code
	b.StatusV2 = code
}

func (MdmBOM) TableName() string {
	return "boms"
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
	return "bom_items"
}
