package model

// TempBOM MES日计划临时替代BOM
type TempBOM struct {
	BaseModel
	TenantID       int64  `json:"tenant_id" gorm:"index;not null"`
	OrderDayID     int64  `json:"order_day_id" gorm:"index;not null"`
	OrderDayItemID int64  `json:"order_day_item_id" gorm:"index;not null"`
	OriginalBOMID  int64  `json:"original_bom_id" gorm:"index"`
	TempBOMName    string `json:"temp_bom_name" gorm:"size:200;not null"`
	BOMContent     string `json:"bom_content" gorm:"type:text"`
	Reason         string `json:"reason" gorm:"size:500"`
	Status         int    `json:"status" gorm:"default:0"`
	Creator       string `json:"creator" gorm:"size:64"`
}

func (TempBOM) TableName() string {
	return "plan_mes_order_day_temp_bom"
}

// TempBOMCreate 创建临时替代BOM请求
type TempBOMCreate struct {
	OrderDayID     int64  `json:"order_day_id" binding:"required"`
	OrderDayItemID int64  `json:"order_day_item_id" binding:"required"`
	OriginalBOMID  int64  `json:"original_bom_id"`
	TempBOMName    string `json:"temp_bom_name" binding:"required"`
	BOMContent     string `json:"bom_content"`
	Reason        string `json:"reason"`
}

// TempBOMUpdate 更新临时替代BOM请求
type TempBOMUpdate struct {
	TempBOMName string `json:"temp_bom_name"`
	BOMContent string `json:"bom_content"`
	Reason     string `json:"reason"`
}

// TempBOMApprove 审核临时替代BOM请求
type TempBOMApprove struct {
	ID      int64  `json:"id" binding:"required"`
	Status  int    `json:"status" binding:"required"`
	Comment string `json:"comment"`
}
