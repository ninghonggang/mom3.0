package model

import (
	"time"
)

// ========== 采购结算模块 ==========

// PurchaseSettlement 采购结算单
type PurchaseSettlement struct {
	BaseModel
	SettlementNo   string     `json:"settlement_no" gorm:"size:50;uniqueIndex;not null"` // 结算单号
	SettlementType string     `json:"settlement_type" gorm:"size:20;not null"`             // NORMAL/RETURN/ADVANCE
	RelatedType    string     `json:"related_type" gorm:"size:20;not null"`               // PURCHASE_ORDER/PURCHASE_RCV/RETURN
	RelatedID      int64      `json:"related_id"`                                         // 关联单据ID
	RelatedNo      string     `json:"related_no" gorm:"size:50"`                         // 采购单号/收货单号

	SupplierID   int64   `json:"supplier_id"`
	SupplierCode string  `json:"supplier_code" gorm:"size:50"`
	SupplierName string  `json:"supplier_name" gorm:"size:100"`

	InvoiceNo   *string `json:"invoice_no" gorm:"size:50"`    // 发票号
	InvoiceDate *string `json:"invoice_date" gorm:"size:10"`  // 发票日期

	GoodsAmount   float64 `json:"goods_amount" gorm:"type:decimal(18,2);default:0"`    // 货款金额
	TaxAmount     float64 `json:"tax_amount" gorm:"type:decimal(18,2);default:0"`      // 税额
	TotalAmount   float64 `json:"total_amount" gorm:"type:decimal(18,2);default:0"`     // 结算总额
	PaidAmount    float64 `json:"paid_amount" gorm:"type:decimal(18,2);default:0"`       // 已付款金额
	DiscountAmount float64 `json:"discount_amount" gorm:"type:decimal(18,2);default:0"`  // 折扣金额

	Currency    string  `json:"currency" gorm:"size:10;default:CNY"`
	ExchangeRate float64 `json:"exchange_rate" gorm:"type:decimal(10,4);default:1"`

	PaymentTerms    *string `json:"payment_terms" gorm:"size:50"`     // 付款条款
	PaymentDueDate  *string `json:"payment_due_date" gorm:"size:10"` // 应付日期
	PaymentMethod   *string `json:"payment_method" gorm:"size:20"`    // 付款方式

	Status      string     `json:"status" gorm:"size:20;default:PENDING"` // PENDING/APPROVED/PAID/CANCELLED
	ApprovedBy  *int64     `json:"approved_by"`
	ApprovedTime *time.Time `json:"approved_time"`

	SettlementDate *string `json:"settlement_date" gorm:"size:10"` // 结算日期

	Remark   *string `json:"remark" gorm:"type:text"`
	TenantID int64   `json:"tenant_id" gorm:"index;not null"`
	CreatedBy *string `json:"created_by" gorm:"size:50"`
}

func (PurchaseSettlement) TableName() string {
	return "fin_purchase_settlement"
}

// PurchaseSettlementItem 采购结算明细
type PurchaseSettlementItem struct {
	BaseModel
	SettlementID int64   `json:"settlement_id" gorm:"index;not null"`
	LineNo      int     `json:"line_no"`

	MaterialID   int64   `json:"material_id"`
	MaterialCode string  `json:"material_code" gorm:"size:50"`
	MaterialName string  `json:"material_name" gorm:"size:100"`
	Specification *string `json:"specification" gorm:"size:200"`
	Unit        *string `json:"unit" gorm:"size:20"`

	InvoiceQty      float64 `json:"invoice_qty" gorm:"type:decimal(18,3)"`       // 发票数量
	ReceivedQty    float64 `json:"received_qty" gorm:"type:decimal(18,3)"`     // 收货数量
	SettledQty     float64 `json:"settled_qty" gorm:"type:decimal(18,3)"`      // 已结算数量
	ThisSettleQty  float64 `json:"this_settle_qty" gorm:"type:decimal(18,3)"`  // 本次结算数量

	UnitPrice   float64 `json:"unit_price" gorm:"type:decimal(18,4)"`  // 单价
	TaxRate     float64 `json:"tax_rate" gorm:"type:decimal(5,2)"`     // 税率
	GoodsAmount float64 `json:"goods_amount" gorm:"type:decimal(18,2)"` // 货款
	TaxAmount   float64 `json:"tax_amount" gorm:"type:decimal(18,2)"`  // 税额
	LineAmount  float64 `json:"line_amount" gorm:"type:decimal(18,2)"`  // 行金额

	BatchNo     *string `json:"batch_no" gorm:"size:50"`
	WarehouseID *int64  `json:"warehouse_id"`
	WarehouseName *string `json:"warehouse_name" gorm:"size:100"`

	Remark   *string `json:"remark" gorm:"type:text"`
	TenantID int64   `json:"tenant_id" gorm:"index;not null"`
}

func (PurchaseSettlementItem) TableName() string {
	return "fin_purchase_settlement_item"
}

// PurchaseAdvance 采购预付款
type PurchaseAdvance struct {
	BaseModel
	AdvanceNo     string  `json:"advance_no" gorm:"size:50;uniqueIndex;not null"`
	SupplierID    int64   `json:"supplier_id"`
	SupplierName  string  `json:"supplier_name" gorm:"size:100"`
	AdvanceAmount float64 `json:"advance_amount" gorm:"type:decimal(18,2);not null"`
	PaidDate     string  `json:"paid_date" gorm:"size:10;not null"`
	PaymentMethod *string `json:"payment_method" gorm:"size:20"`
	PaymentAccount *string `json:"payment_account" gorm:"size:50"`
	BankFlowNo   *string `json:"bank_flow_no" gorm:"size:50"`

	Status      string  `json:"status" gorm:"size:20;default:PENDING"` // PENDING/USED/FINISHED/CANCELLED
	UsedAmount  float64 `json:"used_amount" gorm:"type:decimal(18,2);default:0"`
	SettlementIDs string `json:"settlement_ids" gorm:"type:jsonb"` // 关联结算单ID列表

	Remark   *string `json:"remark" gorm:"type:text"`
	TenantID int64   `json:"tenant_id" gorm:"index;not null"`
	CreatedBy *string `json:"created_by" gorm:"size:50"`
}

func (PurchaseAdvance) TableName() string {
	return "fin_purchase_advance"
}

// ========== 销售结算模块 ==========

// SalesSettlement 销售结算单
type SalesSettlement struct {
	BaseModel
	SettlementNo   string  `json:"settlement_no" gorm:"size:50;uniqueIndex;not null"`
	SettlementType string `json:"settlement_type" gorm:"size:20;not null"`  // NORMAL/RETURN/ADVANCE
	RelatedType    string `json:"related_type" gorm:"size:20;not null"`      // SALES_ORDER/DELIVERY/RETURN
	RelatedID      int64  `json:"related_id"`
	RelatedNo      string `json:"related_no" gorm:"size:50"`

	CustomerID   int64  `json:"customer_id"`
	CustomerCode string `json:"customer_code" gorm:"size:50"`
	CustomerName string `json:"customer_name" gorm:"size:100"`

	InvoiceNo   *string `json:"invoice_no" gorm:"size:50"`
	InvoiceDate *string `json:"invoice_date" gorm:"size:10"`

	GoodsAmount    float64 `json:"goods_amount" gorm:"type:decimal(18,2);default:0"`
	TaxAmount      float64 `json:"tax_amount" gorm:"type:decimal(18,2);default:0"`
	TotalAmount    float64 `json:"total_amount" gorm:"type:decimal(18,2);default:0"`
	ReceivedAmount float64 `json:"received_amount" gorm:"type:decimal(18,2);default:0"`
	DiscountAmount float64 `json:"discount_amount" gorm:"type:decimal(18,2);default:0"`

	Currency string `json:"currency" gorm:"size:10;default:CNY"`

	PaymentTerms   *string `json:"payment_terms" gorm:"size:50"`
	PaymentDueDate *string `json:"payment_due_date" gorm:"size:10"`
	PaymentMethod  *string `json:"payment_method" gorm:"size:20"`

	Status      string     `json:"status" gorm:"size:20;default:PENDING"`
	ApprovedBy  *int64     `json:"approved_by"`
	ApprovedTime *time.Time `json:"approved_time"`

	SettlementDate *string `json:"settlement_date" gorm:"size:10"`

	Remark   *string `json:"remark" gorm:"type:text"`
	TenantID int64   `json:"tenant_id" gorm:"index;not null"`
	CreatedBy *string `json:"created_by" gorm:"size:50"`
}

func (SalesSettlement) TableName() string {
	return "fin_sales_settlement"
}

// SalesSettlementItem 销售结算明细
type SalesSettlementItem struct {
	BaseModel
	SettlementID int64   `json:"settlement_id" gorm:"index;not null"`
	LineNo      int     `json:"line_no"`

	MaterialID   int64   `json:"material_id"`
	MaterialCode  string  `json:"material_code" gorm:"size:50"`
	MaterialName  string  `json:"material_name" gorm:"size:100"`
	Specification *string `json:"specification" gorm:"size:200"`
	Unit         *string `json:"unit" gorm:"size:20"`

	InvoiceQty     float64 `json:"invoice_qty" gorm:"type:decimal(18,3)"`
	ShippedQty    float64 `json:"shipped_qty" gorm:"type:decimal(18,3)"`
	SettledQty    float64 `json:"settled_qty" gorm:"type:decimal(18,3)"`
	ThisSettleQty float64 `json:"this_settle_qty" gorm:"type:decimal(18,3)"`

	UnitPrice   float64 `json:"unit_price" gorm:"type:decimal(18,4)"`
	TaxRate     float64 `json:"tax_rate" gorm:"type:decimal(5,2)"`
	GoodsAmount float64 `json:"goods_amount" gorm:"type:decimal(18,2)"`
	TaxAmount   float64 `json:"tax_amount" gorm:"type:decimal(18,2)"`
	LineAmount  float64 `json:"line_amount" gorm:"type:decimal(18,2)"`

	BatchNo    *string `json:"batch_no" gorm:"size:50"`
	Remark     *string `json:"remark" gorm:"type:text"`
	TenantID   int64   `json:"tenant_id" gorm:"index;not null"`
}

func (SalesSettlementItem) TableName() string {
	return "fin_sales_settlement_item"
}

// SalesReceipt 销售收款单
type SalesReceipt struct {
	BaseModel
	ReceiptNo     string  `json:"receipt_no" gorm:"size:50;uniqueIndex;not null"`
	CustomerID    int64   `json:"customer_id"`
	CustomerName  string  `json:"customer_name" gorm:"size:100"`
	ReceiptAmount float64 `json:"receipt_amount" gorm:"type:decimal(18,2);not null"`
	ReceiptDate   string  `json:"receipt_date" gorm:"size:10;not null"`
	ReceiptMethod *string `json:"receipt_method" gorm:"size:20"`   // CASH/TRANSFER/ENDORSEMENT
	ReceiptAccount *string `json:"receipt_account" gorm:"size:50"`
	BankFlowNo    *string `json:"bank_flow_no" gorm:"size:50"`
	SettlementIDs  string  `json:"settlement_ids" gorm:"type:jsonb"` // 关联结算单
	UsedAmount    float64 `json:"used_amount" gorm:"type:decimal(18,2);default:0"`
	Status        string  `json:"status" gorm:"size:20;default:PENDING"` // PENDING/USED/FINISHED
	Remark        *string `json:"remark" gorm:"type:text"`
	TenantID      int64   `json:"tenant_id" gorm:"index;not null"`
	CreatedBy     *string `json:"created_by" gorm:"size:50"`
}

func (SalesReceipt) TableName() string {
	return "fin_sales_receipt"
}

// ========== 付款申请模块 ==========

// PaymentRequest 付款申请单
type PaymentRequest struct {
	BaseModel
	RequestNo    string  `json:"request_no" gorm:"size:50;uniqueIndex;not null"`
	RequestType  string  `json:"request_type" gorm:"size:20;not null"` // PURCHASE/SALES/EXPENSE/OTHER
	SupplierCustomerID *int64 `json:"supplier_customer_id"`
	SupplierCustomerName *string `json:"supplier_customer_name" gorm:"size:100"`

	RequestAmount float64 `json:"request_amount" gorm:"type:decimal(18,2);not null"`
	AmountInWords *string `json:"amount_in_words" gorm:"size:200"` // 大写金额

	Purpose      *string `json:"purpose" gorm:"size:200"` // 付款用途
	BankName     *string `json:"bank_name" gorm:"size:100"`
	BankAccount  *string `json:"bank_account" gorm:"size:100"`

	SettlementIDs string `json:"settlement_ids" gorm:"type:jsonb"`    // 关联结算单
	AttachmentURLs string `json:"attachment_urls" gorm:"type:jsonb"` // 附件URL

	Status        string `json:"status" gorm:"size:20;default:PENDING"`
	ApprovalStatus string `json:"approval_status" gorm:"size:20;default:PENDING"`
	ApprovedBy    *int64  `json:"approved_by"`
	ApprovedTime *time.Time `json:"approved_time"`
	ApproverComment *string `json:"approver_comment" gorm:"type:text"`

	PaidBy   *int64     `json:"paid_by"`
	PaidTime *time.Time `json:"paid_time"`
	PaymentStatus string `json:"payment_status" gorm:"size:20;default:UNPAID"` // UNPAID/PARTIAL_PAID/PAID

	TenantID int64   `json:"tenant_id" gorm:"index;not null"`
	CreatedBy *string `json:"created_by" gorm:"size:50"`
}

func (PaymentRequest) TableName() string {
	return "fin_payment_request"
}

// ========== 对账模块 ==========

// SupplierStatement 供应商对账单
type SupplierStatement struct {
	BaseModel
	StatementNo    string  `json:"statement_no" gorm:"size:50;uniqueIndex;not null"`
	SupplierID     int64   `json:"supplier_id"`
	SupplierName   string  `json:"supplier_name" gorm:"size:100"`
	StatementPeriod string `json:"statement_period" gorm:"size:20;not null"` // 2026-01
	StartDate      string  `json:"start_date" gorm:"size:10;not null"`
	EndDate        string  `json:"end_date" gorm:"size:10;not null"`

	BeginningAmount   float64 `json:"beginning_amount" gorm:"type:decimal(18,2);default:0"`    // 期初
	PurchaseAmount   float64 `json:"purchase_amount" gorm:"type:decimal(18,2);default:0"`    // 本期采购
	OtherAmount      float64 `json:"other_amount" gorm:"type:decimal(18,2);default:0"`       // 其他应付
	PaymentAmount    float64 `json:"payment_amount" gorm:"type:decimal(18,2);default:0"`      // 本期付款
	DiscountAmount   float64 `json:"discount_amount" gorm:"type:decimal(18,2);default:0"`   // 折扣
	AdjustmentAmount float64 `json:"adjustment_amount" gorm:"type:decimal(18,2);default:0"` // 调整

	EndingAmount float64 `json:"ending_amount" gorm:"type:decimal(18,2);default:0"` // 期末

	Status        string     `json:"status" gorm:"size:20;default:PENDING"` // PENDING/CONFIRMED/DISPUTED
	ConfirmedBy   *int64      `json:"confirmed_by"`
	ConfirmedTime *time.Time  `json:"confirmed_time"`
	DisputeReason *string     `json:"dispute_reason" gorm:"type:text"`

	TenantID   int64   `json:"tenant_id" gorm:"index;not null"`
	CreatedBy *string  `json:"created_by" gorm:"size:50"`
}

func (SupplierStatement) TableName() string {
	return "fin_supplier_statement"
}

// StatementDetail 对账单明细
type StatementDetail struct {
	BaseModel
	StatementID int64   `json:"statement_id" gorm:"index;not null"`
	DetailType  string  `json:"detail_type" gorm:"size:20;not null"` // PURCHASE/PAYMENT/ADJUSTMENT
	RelatedNo   string  `json:"related_no" gorm:"size:50"`
	BizDate     *string `json:"biz_date" gorm:"size:10"`
	Amount      float64 `json:"amount" gorm:"type:decimal(18,2)"`
	Remark      *string `json:"remark" gorm:"type:text"`
	TenantID    int64   `json:"tenant_id" gorm:"index;not null"`
}

func (StatementDetail) TableName() string {
	return "fin_statement_detail"
}

// ========== 查询请求结构 ==========

// PurchaseSettlementQuery 采购结算查询
type PurchaseSettlementQuery struct {
	SupplierID   *int64  `json:"supplier_id"`
	Status      *string `json:"status"`
	StartDate   *string `json:"start_date"`
	EndDate     *string `json:"end_date"`
	SettlementNo *string `json:"settlement_no"`
	Page        int     `json:"page"`
	Limit       int     `json:"limit"`
}

// SalesSettlementQuery 销售结算查询
type SalesSettlementQuery struct {
	CustomerID   *int64  `json:"customer_id"`
	Status      *string `json:"status"`
	StartDate   *string `json:"start_date"`
	EndDate     *string `json:"end_date"`
	SettlementNo *string `json:"settlement_no"`
	Page        int     `json:"page"`
	Limit       int     `json:"limit"`
}

// PaymentRequestQuery 付款申请查询
type PaymentRequestQuery struct {
	RequestType *string `json:"request_type"`
	Status      *string `json:"status"`
	StartDate   *string `json:"start_date"`
	EndDate     *string `json:"end_date"`
	Page        int     `json:"page"`
	Limit       int     `json:"limit"`
}

// PurchaseSettlementCreate 创建采购结算请求
type PurchaseSettlementCreate struct {
	SupplierID    int64                   `json:"supplier_id" binding:"required"`
	RelatedType   string                  `json:"related_type" binding:"required"` // PURCHASE_ORDER/PURCHASE_RCV
	RelatedID     int64                   `json:"related_id" binding:"required"`
	InvoiceNo     *string                 `json:"invoice_no"`
	InvoiceDate   *string                 `json:"invoice_date"`
	Items        []PurchaseSettlementItem `json:"items"`
	PaymentTerms  *string                 `json:"payment_terms"`
	PaymentDueDate *string                `json:"payment_due_date"`
	PaymentMethod *string                 `json:"payment_method"`
	Remark       *string                  `json:"remark"`
}

// SalesSettlementCreate 创建销售结算请求
type SalesSettlementCreate struct {
	CustomerID    int64                    `json:"customer_id" binding:"required"`
	RelatedType  string                   `json:"related_type" binding:"required"` // SALES_ORDER/DELIVERY
	RelatedID    int64                    `json:"related_id" binding:"required"`
	InvoiceNo    *string                  `json:"invoice_no"`
	InvoiceDate  *string                  `json:"invoice_date"`
	Items       []SalesSettlementItem      `json:"items"`
	PaymentTerms *string                  `json:"payment_terms"`
	PaymentDueDate *string                `json:"payment_due_date"`
	PaymentMethod *string                 `json:"payment_method"`
	Remark       *string                  `json:"remark"`
}

// PaymentRequestCreate 创建付款申请请求
type PaymentRequestCreate struct {
	RequestType  string   `json:"request_type" binding:"required"`
	SupplierCustomerID *int64 `json:"supplier_customer_id"`
	RequestAmount float64 `json:"request_amount" binding:"required"`
	Purpose     *string  `json:"purpose"`
	BankName    *string  `json:"bank_name"`
	BankAccount *string  `json:"bank_account"`
	SettlementIDs []int64 `json:"settlement_ids"`
	Remark      *string  `json:"remark"`
}
