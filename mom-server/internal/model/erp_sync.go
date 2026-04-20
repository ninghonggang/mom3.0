package model

import "time"

// IntegrationERPSyncLog ERP同步日志
type IntegrationERPSyncLog struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	SyncType     string    `json:"syncType" gorm:"column:sync_type;size:50;not null"` // BOM/PRODUCTION_ORDER/REPORT/STOCK_IN/QUALITY/STOCK
	Direction    string    `json:"direction" gorm:"column:direction;size:10;not null"` // OUTBOUND/INBOUND
	ERPBillNo    string    `json:"erpBillNo" gorm:"column:erp_bill_no;size:100"`
	MESBillNo    string    `json:"mesBillNo" gorm:"column:mes_bill_no;size:100"`
	RequestBody  string    `json:"requestBody" gorm:"column:request_body;type:text"`
	ResponseBody string    `json:"responseBody" gorm:"column:response_body;type:text"`
	Status       string    `json:"status" gorm:"column:status;size:20"` // PENDING/SUCCESS/FAILED
	ErrorMsg     string    `json:"errorMsg" gorm:"column:error_msg;type:text"`
	RetryCount   int       `json:"retryCount" gorm:"column:retry_count;default:0"`
	TenantID     int64     `json:"tenantId" gorm:"column:tenant_id;not null"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

// IntegrationERPMapping ERP字段映射
type IntegrationERPMapping struct {
	ID            int64     `json:"id" gorm:"primaryKey"`
	ERPTableName  string    `json:"erpTableName" gorm:"column:erp_table_name;size:50;not null"`
	ERPFieldName  string    `json:"erpFieldName" gorm:"column:erp_field_name;size:50;not null"`
	MESTableName  string    `json:"mesTableName" gorm:"column:mes_table_name;size:50;not null"`
	MESFieldName  string    `json:"mesFieldName" gorm:"column:mes_field_name;size:50;not null"`
	TransformRule string    `json:"transformRule" gorm:"column:transform_rule;size:200"`
	TenantID      int64     `json:"tenantId" gorm:"column:tenant_id;not null"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

// ERPSyncLogQuery 同步日志查询
type ERPSyncLogQuery struct {
	TenantID  int64  `json:"tenantId"`
	SyncType  string `json:"syncType"`
	Direction string `json:"direction"`
	Status    string `json:"status"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Page      int    `json:"page"`
	PageSize  int    `json:"pageSize"`
}

// CreateERPSyncLogRequest 创建同步日志请求
type CreateERPSyncLogRequest struct {
	SyncType    string `json:"syncType" binding:"required"`
	Direction   string `json:"direction" binding:"required"`
	ERPBillNo   string `json:"erpBillNo"`
	MESBillNo   string `json:"mesBillNo"`
	RequestBody string `json:"requestBody"`
	TenantID    int64  `json:"tenantId"`
}

// UpdateERPSyncLogRequest 更新同步日志请求
type UpdateERPSyncLogRequest struct {
	ResponseBody string `json:"responseBody"`
	Status       string `json:"status"`
	ErrorMsg     string `json:"errorMsg"`
	RetryCount   int    `json:"retryCount"`
}

// ERPSyncStatus 常量
var ERPSyncStatus = struct {
	Pending string
	Success string
	Failed  string
}{
	Pending: "PENDING",
	Success: "SUCCESS",
	Failed:  "FAILED",
}

// ERPSyncDirection 常量
var ERPSyncDirection = struct {
	Outbound string
	Inbound  string
}{
	Outbound: "OUTBOUND",
	Inbound:  "INBOUND",
}

// ERPSyncType 常量
var ERPSyncType = struct {
	BOM              string
	ProductionOrder   string
	Report           string
	StockIn          string
	Quality          string
	Stock            string
}{
	BOM:              "BOM",
	ProductionOrder:  "PRODUCTION_ORDER",
	Report:           "REPORT",
	StockIn:          "STOCK_IN",
	Quality:          "QUALITY",
	Stock:            "STOCK",
}

// ERPSyncBOMRequest BOM同步请求
type ERPSyncBOMRequest struct {
	ERPBillNo    string         `json:"erpBillNo"`
	Items        []BOMSyncItem  `json:"items"`
}

// BOMSyncItem BOM同步项
type BOMSyncItem struct {
	MaterialCode string  `json:"materialCode"`
	MaterialName string  `json:"materialName"`
	ParentCode   string  `json:"parentCode"`
	ChildCode    string  `json:"childCode"`
	Qty          float64 `json:"qty"`
	Unit         string  `json:"unit"`
}

// ERPSyncProductionOrderRequest 生产订单同步请求
type ERPSyncProductionOrderRequest struct {
	ERPBillNo     string                  `json:"erpBillNo"`
	OrderNo       string                  `json:"orderNo"`
	MaterialCode  string                  `json:"materialCode"`
	Qty           float64                 `json:"qty"`
	StartDate     string                  `json:"startDate"`
	EndDate       string                  `json:"endDate"`
	WorkshopCode  string                  `json:"workshopCode"`
}

// ERPSyncStockRequest 库存同步请求
type ERPSyncStockRequest struct {
	MaterialCode string  `json:"materialCode"`
	WarehouseCode string `json:"warehouseCode"`
	Qty          float64 `json:"qty"`
	FrozenQty    float64 `json:"frozenQty"`
	AvailableQty float64 `json:"availableQty"`
}

// ERPPushReportRequest 报工回传请求
type ERPPushReportRequest struct {
	MESBillNo    string  `json:"mesBillNo"`
	OrderNo      string  `json:"orderNo"`
	ReportQty    float64 `json:"reportQty"`
	ReportDate   string  `json:"reportDate"`
	ReportUser   string  `json:"reportUser"`
	WorkstationCode string `json:"workstationCode"`
}

// ERPPushStockInRequest 入库通知回传请求
type ERPPushStockInRequest struct {
	MESBillNo    string  `json:"mesBillNo"`
	OrderNo      string  `json:"orderNo"`
	MaterialCode string  `json:"materialCode"`
	InQty        float64 `json:"inQty"`
	InDate       string  `json:"inDate"`
	WarehouseCode string `json:"warehouseCode"`
}

// ERPPushQualityRequest 质检数据回传请求
type ERPPushQualityRequest struct {
	InspectNo    string  `json:"inspectNo"`
	MaterialCode string  `json:"materialCode"`
	BatchNo      string  `json:"batchNo"`
	QualifiedQty float64 `json:"qualifiedQty"`
	RejectQty    float64 `json:"rejectQty"`
	InspectDate  string  `json:"inspectDate"`
}

// ERPSyncResult 同步结果
type ERPSyncResult struct {
	Success   bool   `json:"success"`
	ERPBillNo string `json:"erpBillNo"`
	MESBillNo string `json:"mesBillNo"`
	Message   string `json:"message"`
}
