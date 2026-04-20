package model

// WMSItem 仓库货品基础信息
type WMSItem struct {
	BaseModel
	TenantID      int64   `json:"tenant_id" gorm:"index;not null"`
	ItemCode      string  `json:"item_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_item_code"`
	ItemName      string  `json:"item_name" gorm:"size:100;not null"`
	Specification string  `json:"specification" gorm:"size:200"` // 规格型号
	Unit          string  `json:"unit" gorm:"size:20"`            // 单位
	ItemType      string  `json:"item_type" gorm:"size:20"`      // 类型: RAW/PACKAGE/FINISHED
	CategoryID    *int64  `json:"category_id" gorm:""`           // 货品分类ID
	Barcode       string  `json:"barcode" gorm:"size:50"`        // 条码
	SafetyStock   float64 `json:"safety_stock" gorm:"type:decimal(18,3);default:0"` // 安全库存
	Status        string  `json:"status" gorm:"size:20;default:'ACTIVE'"`
}

func (WMSItem) TableName() string {
	return "wms_itembasic"
}

// WMSItemCreateReqVO 货品创建请求
type WMSItemCreateReqVO struct {
	ItemCode      string  `json:"item_code"`
	ItemName      string  `json:"item_name"`
	Specification string  `json:"specification"`
	Unit          string  `json:"unit"`
	ItemType      string  `json:"item_type"`
	CategoryID    *int64  `json:"category_id"`
	Barcode       string  `json:"barcode"`
	SafetyStock   float64 `json:"safety_stock"`
}

// WMSItemUpdateReqVO 货品更新请求
type WMSItemUpdateReqVO struct {
	Id           int64   `json:"id"`
	ItemName     string  `json:"item_name"`
	Specification string `json:"specification"`
	Unit         string  `json:"unit"`
	CategoryID   *int64  `json:"category_id"`
	Barcode      string  `json:"barcode"`
	SafetyStock  float64 `json:"safety_stock"`
	Status       string  `json:"status"`
}

// WMSItemQueryVO 货品查询请求
type WMSItemQueryVO struct {
	Keyword  string `form:"keyword"`  // 搜索关键字
	ItemType string `form:"item_type"` // 货品类型
	Status   string `form:"status"`    // 状态
	Page     int    `form:"page"`      // 页码
	PageSize int    `form:"page_size"` // 每页数量
}

// WMSItemRespVO 货品响应
type WMSItemRespVO struct {
	ID           int64   `json:"id"`
	ItemCode     string  `json:"item_code"`
	ItemName     string  `json:"item_name"`
	Specification string `json:"specification"`
	Unit         string  `json:"unit"`
	ItemType     string  `json:"item_type"`
	CategoryID   *int64  `json:"category_id"`
	Barcode      string  `json:"barcode"`
	SafetyStock  float64 `json:"safety_stock"`
	Status       string  `json:"status"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}