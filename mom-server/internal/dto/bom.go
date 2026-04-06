package dto

// BOMCreateReq BOM创建/更新请求
type BOMCreateReq struct {
	TenantID     int64         `json:"tenant_id"`
	BOMCode      string        `json:"bom_code" binding:"required"`
	BOMName      string        `json:"bom_name" binding:"required"`
	MaterialID   int64         `json:"material_id" binding:"required"`
	MaterialCode string        `json:"material_code"`
	MaterialName string        `json:"material_name"`
	Version      string        `json:"version"`
	Status       string        `json:"status"`
	EffDate      string        `json:"eff_date"`
	ExpDate      string        `json:"exp_date"`
	Remark       string        `json:"remark"`
	ErpBomCode   string        `json:"erp_bom_code"`
	ErpSyncStatus string        `json:"erp_sync_status"`
	IsCurrent    int           `json:"is_current"`
	Items        []BOMItemDTO  `json:"items"`
}

// BOMItemDTO BOM明细
type BOMItemDTO struct {
	LineNo          int     `json:"line_no"`
	MaterialID      int64   `json:"material_id" binding:"required"`
	MaterialCode    string  `json:"material_code"`
	MaterialName    string  `json:"material_name"`
	Quantity        float64 `json:"quantity" binding:"required"`
	Unit            string  `json:"unit"`
	ScrapRate       float64 `json:"scrap_rate"`
	SubstituteGroup string  `json:"substitute_group"`
	IsAlternative   int     `json:"is_alternative"`
}

// BOMStatusReq BOM状态更新请求
type BOMStatusReq struct {
	Status string `json:"status" binding:"required"`
}
