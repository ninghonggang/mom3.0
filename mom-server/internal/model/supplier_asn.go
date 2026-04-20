package model

import "time"

// SupplierASN 供应商ASN到货通知
type SupplierASN struct {
	ID             int64      `json:"id" gorm:"primaryKey"`
	ASNNo          string     `json:"asnNo" gorm:"column:asn_no;uniqueIndex;size:50;not null"`
	SupplierID     int64      `json:"supplierId" gorm:"column:supplier_id;not null"`
	SupplierCode   string     `json:"supplierCode" gorm:"column:supplier_code;size:50;not null"`
	SupplierName   string     `json:"supplierName" gorm:"column:supplier_name;size:100"`
	DeliveryType   string     `json:"deliveryType" gorm:"column:delivery_type;size:20;default:NORMAL"` // NORMAL/URGENT
	DeliveryDate   *time.Time `json:"deliveryDate" gorm:"column:delivery_date"`
	DeliveryStart  string     `json:"deliveryTimeStart" gorm:"column:delivery_time_start;size:20"`
	DeliveryEnd    string     `json:"deliveryTimeEnd" gorm:"column:delivery_time_end;size:20"`
	WarehouseCode  string     `json:"warehouseCode" gorm:"column:warehouse_code;size:50"`
	ContactPerson  string     `json:"contactPerson" gorm:"column:contact_person;size:50"`
	ContactPhone   string     `json:"contactPhone" gorm:"column:contact_phone;size:20"`
	Status         string     `json:"status" gorm:"column:status;size:20;default:DRAFT"` // DRAFT/SUBMITTED/CONFIRMED/RECEIVING/COMPLETED/CANCELLED
	TotalQty       float64    `json:"totalQty" gorm:"column:total_qty;type:decimal(18,3)"`
	TotalAmount    float64    `json:"totalAmount" gorm:"column:total_amount;type:decimal(18,2)"`
	Remark         string     `json:"remark" gorm:"column:remark;type:text"`
	TenantID       int64      `json:"tenantId" gorm:"column:tenant_id;not null"`
	CreatedBy      int64      `json:"createdBy" gorm:"column:created_by"`
	CreatedAt      time.Time  `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt      time.Time  `json:"updatedAt" gorm:"column:updated_at"`
	Items          []SupplierASNItem `json:"items" gorm:"foreignKey:ASNID"`
}

// SupplierASNItem ASN明细
type SupplierASNItem struct {
	ID           int64   `json:"id" gorm:"primaryKey"`
	ASNID        int64   `json:"asnId" gorm:"column:asn_id;not null"`
	LineNo       int     `json:"lineNo" gorm:"column:line_no"`
	MaterialCode string  `json:"materialCode" gorm:"column:material_code;size:50;not null"`
	MaterialName string  `json:"materialName" gorm:"column:material_name;size:100"`
	Spec         string  `json:"spec" gorm:"column:spec;size:100"`
	Unit         string  `json:"unit" gorm:"column:unit;size:20"`
	BatchNo      string  `json:"batchNo" gorm:"column:batch_no;size:50"`
	Qty          float64 `json:"qty" gorm:"column:qty;type:decimal(18,3)"`
	QualifiedQty float64 `json:"qualifiedQty" gorm:"column:qualified_qty;type:decimal(18,3)"`
	Price        float64 `json:"price" gorm:"column:price;type:decimal(18,4)"`
	Amount       float64 `json:"amount" gorm:"column:amount;type:decimal(18,2)"`
	PackingQty   float64 `json:"packingQty" gorm:"column:packing_qty;type:decimal(18,3)"`
	PackingUnit  string  `json:"packingUnit" gorm:"column:packing_unit;size:20"`
	ReceivedQty  float64 `json:"receivedQty" gorm:"column:received_qty;type:decimal(18,3)"`
	TenantID     int64   `json:"tenantId" gorm:"column:tenant_id;not null"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at"`
}

// SupplierASNQuery ASN查询
type SupplierASNQuery struct {
	TenantID     int64  `json:"tenantId"`
	SupplierCode string `json:"supplierCode"`
	Status       string `json:"status"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	Page         int    `json:"page"`
	PageSize     int    `json:"pageSize"`
}

// CreateSupplierASNRequest 创建ASN请求
type CreateSupplierASNRequest struct {
	TenantID       int64                    `json:"tenantId"`
	SupplierID     int64                    `json:"supplierId" binding:"required"`
	SupplierCode   string                   `json:"supplierCode" binding:"required"`
	SupplierName   string                   `json:"supplierName"`
	DeliveryType   string                   `json:"deliveryType"`
	DeliveryDate   string                   `json:"deliveryDate"`
	DeliveryStart  string                   `json:"deliveryTimeStart"`
	DeliveryEnd    string                   `json:"deliveryTimeEnd"`
	WarehouseCode  string                   `json:"warehouseCode"`
	ContactPerson  string                   `json:"contactPerson"`
	ContactPhone   string                   `json:"contactPhone"`
	Remark         string                   `json:"remark"`
	Items          []CreateASNItemRequest   `json:"items"`
}

// CreateASNItemRequest 创建ASN明细请求
type CreateASNItemRequest struct {
	MaterialCode string  `json:"materialCode" binding:"required"`
	MaterialName string  `json:"materialName"`
	Spec         string  `json:"spec"`
	Unit         string  `json:"unit"`
	BatchNo      string  `json:"batchNo"`
	Qty          float64 `json:"qty" binding:"required"`
	Price        float64 `json:"price"`
	Amount       float64 `json:"amount"`
	PackingQty   float64 `json:"packingQty"`
	PackingUnit  string  `json:"packingUnit"`
}

// UpdateSupplierASNRequest 更新ASN请求
type UpdateSupplierASNRequest struct {
	DeliveryType   string  `json:"deliveryType"`
	DeliveryDate   string  `json:"deliveryDate"`
	DeliveryStart  string  `json:"deliveryTimeStart"`
	DeliveryEnd    string  `json:"deliveryTimeEnd"`
	WarehouseCode  string  `json:"warehouseCode"`
	ContactPerson  string  `json:"contactPerson"`
	ContactPhone   string  `json:"contactPhone"`
	Status         string  `json:"status"`
	Remark         string  `json:"remark"`
}

// ConfirmASNRequest 确认ASN请求
type ConfirmASNRequest struct {
	ConfirmStatus       string `json:"confirmStatus"` // CONFIRMED/REJECTED
	ConfirmedDeliveryDate string `json:"confirmedDeliveryDate"`
	ConfirmedQty        float64 `json:"confirmedQty"`
	RejectReason        string  `json:"rejectReason"`
	ModifiedItems       []ModifiedItem `json:"modifiedItems"`
}

// ModifiedItem 修改的明细项
type ModifiedItem struct {
	ItemID   int64   `json:"itemId"`
	Qty      float64 `json:"qty"`
	Price    float64 `json:"price"`
}

// ASNStatus ASN状态常量
var ASNStatus = struct {
	Draft      string
	Submitted  string
	Confirmed  string
	Receiving  string
	Completed  string
	Cancelled  string
}{
	Draft:      "DRAFT",
	Submitted:  "SUBMITTED",
	Confirmed:  "CONFIRMED",
	Receiving:  "RECEIVING",
	Completed:  "COMPLETED",
	Cancelled:  "CANCELLED",
}
