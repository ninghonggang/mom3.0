package model

import "time"

// WmsArea 库区表
type WmsArea struct {
	ID            uint64    `gorm:"primaryKey" json:"id"`
	TenantID      int64     `gorm:"index;uniqueIndex:idx_tenant_code" json:"tenantId"`
	WarehouseCode string    `gorm:"size:50" json:"warehouseCode"`
	WarehouseName string    `gorm:"size:100" json:"warehouseName"`
	AreaCode      string    `gorm:"size:50;uniqueIndex:idx_tenant_code" json:"areaCode"`
	AreaName      string    `gorm:"size:100" json:"areaName"`
	AreaType      string    `gorm:"size:20" json:"areaType"` // STORAGE/PICKING/RECEIVING/SHIPPING/RETURN/QC
	ParentCode    string    `gorm:"size:50" json:"parentCode"`
	Level         int       `gorm:"default:1" json:"level"`
	Status        string    `gorm:"size:20;default:'ACTIVE'" json:"status"`
	Remark        string    `gorm:"size:500" json:"remark"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func (WmsArea) TableName() string {
	return "wms_area"
}

// WmsAreaQueryVO 查询参数
type WmsAreaQueryVO struct {
	Page          int    `form:"page"`
	PageSize      int    `form:"pageSize"`
	Keyword       string `form:"keyword"`
	WarehouseCode string `form:"warehouseCode"`
	AreaType      string `form:"areaType"`
	Level         int    `form:"level"`
	Status        string `form:"status"`
}

// WmsAreaTreeVO 树形结构
type WmsAreaTreeVO struct {
	WmsArea
	Children []WmsAreaTreeVO `json:"children"`
}

// WmsAreaCreateReqVO 创建请求
type WmsAreaCreateReqVO struct {
	WarehouseCode string `json:"warehouseCode" binding:"required"`
	WarehouseName string `json:"warehouseName"`
	AreaCode      string `json:"areaCode" binding:"required"`
	AreaName      string `json:"areaName" binding:"required"`
	AreaType      string `json:"areaType"`
	ParentCode    string `json:"parentCode"`
	Level         int    `json:"level"`
	Status        string `json:"status"`
	Remark        string `json:"remark"`
}

// WmsAreaUpdateReqVO 更新请求
type WmsAreaUpdateReqVO struct {
	WarehouseName string `json:"warehouseName"`
	AreaName      string `json:"areaName"`
	AreaType      string `json:"areaType"`
	Status        string `json:"status"`
	Remark        string `json:"remark"`
}
