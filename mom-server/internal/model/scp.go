package model

import (
	"encoding/json"
	"time"
)

// SCPBaseModel 供应链公共字段
type SCPBaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	TenantID  int64          `json:"tenant_id" gorm:"index;not null"`
}

// PurchaseOrder 采购订单头表
type PurchaseOrder struct {
	SCPBaseModel
	PONo             string          `json:"po_no" gorm:"size:50;uniqueIndex;not null"`     // 采购订单号
	POType           string          `json:"po_type" gorm:"size:20;default:STANDARD"`        // STANDARD/URGENT/LONG_TERM
	SupplierID       int64           `json:"supplier_id" gorm:"not null"`                     // 供应商ID
	SupplierCode     string          `json:"supplier_code" gorm:"size:50"`                    // 供应商编码
	SupplierName     string          `json:"supplier_name" gorm:"size:100"`                   // 供应商名称
	ContactPerson    *string         `json:"contact_person" gorm:"size:50"`                   // 联系人
	ContactPhone     *string         `json:"contact_phone" gorm:"size:20"`                    // 联系电话
	ContactEmail     *string         `json:"contact_email" gorm:"size:100"`                   // 邮箱
	OrderDate        time.Time       `json:"order_date" gorm:"type:date;not null"`            // 订单日期
	PromisedDate     *time.Time      `json:"promised_date" gorm:"type:date"`                  // 承诺交货日期
	Currency         string          `json:"currency" gorm:"size:10;default:CNY"`             // 币种
	PaymentTerms     *string         `json:"payment_terms" gorm:"size:50"`                    // 付款条款
	TaxRate          float64         `json:"tax_rate" gorm:"type:decimal(5,2);default:13.00"` // 税率
	TotalAmount      float64         `json:"total_amount" gorm:"type:decimal(18,2);default:0"` // 总金额
	TotalQty         float64         `json:"total_qty" gorm:"type:decimal(18,3);default:0"`  // 总数量
	ApprovedBy       *int64          `json:"approved_by"`                                     // 审批人
	ApprovedTime     *time.Time      `json:"approved_time"`                                  // 审批时间
	ApprovalStatus   string          `json:"approval_status" gorm:"size:20;default:PENDING"` // PENDING/APPROVED/REJECTED
	Status           string          `json:"status" gorm:"size:20;default:DRAFT"`             // DRAFT/PENDING/APPROVED/ISSUED/PARTIAL/RECEIVED/CLOSED/CANCELLED
	CloseReason      *string         `json:"close_reason" gorm:"size:500"`                   // 关闭原因
	SourceType       *string         `json:"source_type" gorm:"size:20"`                     // MANUAL/MPS/PR
	SourceNo         *string         `json:"source_no" gorm:"size:50"`                        // 来源单据号
	Remark           *string         `json:"remark" gorm:"type:text"`                        // 备注
	Items            []PurchaseOrderItem `json:"items" gorm:"foreignKey:POID"`                 // 订单明细
}

func (PurchaseOrder) TableName() string {
	return "scp_purchase_order"
}

// PurchaseOrderItem 采购订单明细
type PurchaseOrderItem struct {
	ID                  uint           `json:"id" gorm:"primarykey"`
	POID                int64          `json:"po_id" gorm:"not null;index"`                   // 订单ID
	LineNo              int            `json:"line_no" gorm:"not null"`                       // 行号
	MaterialID          *int64         `json:"material_id"`                                  // 物料ID
	MaterialCode        string         `json:"material_code" gorm:"size:50;not null"`         // 物料编码
	MaterialName        *string        `json:"material_name" gorm:"size:100"`                 // 物料名称
	Specification       *string        `json:"specification" gorm:"size:200"`                 // 规格型号
	Unit                string         `json:"unit" gorm:"size:20;default:PCS"`               // 单位
	UnitPrice           float64        `json:"unit_price" gorm:"type:decimal(18,4);default:0"` // 单价
	OrderQty            float64        `json:"order_qty" gorm:"type:decimal(18,3);not null"` // 订单数量
	DeliveredQty        float64        `json:"delivered_qty" gorm:"type:decimal(18,3);default:0"` // 已交货数量
	ReceivedQty         float64        `json:"received_qty" gorm:"type:decimal(18,3);default:0"` // 已收货数量
	TaxAmount           float64        `json:"tax_amount" gorm:"type:decimal(18,2);default:0"` // 税额
	LineAmount          float64        `json:"line_amount" gorm:"type:decimal(18,2);default:0"` // 行金额
	PromisedDate        *time.Time     `json:"promised_date" gorm:"type:date"`                // 承诺交期
	ActualDeliveryDate  *time.Time     `json:"actual_delivery_date" gorm:"type:date"`         // 实际交货日期
	BatchNo             *string        `json:"batch_no" gorm:"size:50"`                      // 批号
	QualityRequire      *string        `json:"quality_require" gorm:"type:text"`              // 质量要求
	PackageRequire      *string        `json:"package_require" gorm:"size:200"`                // 包装要求
	IsGifted            int            `json:"is_gifted" gorm:"default:0"`                    // 是否赠品
	Status              string         `json:"status" gorm:"size:20;default:PENDING"`         // PENDING/PARTIAL/COMPLETED/CANCELLED
	Remark              *string        `json:"remark" gorm:"type:text"`                       // 备注
	CreatedAt           time.Time      `json:"created_at" gorm:"autoCreateTime"`
}

func (PurchaseOrderItem) TableName() string {
	return "scp_purchase_order_item"
}

// POChangeLog 采购订单变更记录
type POChangeLog struct {
	ID           uint      `json:"id" gorm:"primarykey"`
	POID         int64     `json:"po_id" gorm:"not null;index"`
	ChangeType   string    `json:"change_type" gorm:"size:30;not null"` // QTY_CHANGE/PRICE_CHANGE/DATE_CHANGE/SUPPLIER_CHANGE/ITEM_ADD/ITEM_REMOVE
	ChangeField  *string   `json:"change_field" gorm:"size:50"`
	OldValue     *string   `json:"old_value" gorm:"size:200"`
	NewValue     *string   `json:"new_value" gorm:"size:200"`
	ChangedBy    *int64    `json:"changed_by"`
	ChangedByName *string  `json:"changed_by_name" gorm:"size:50"`
	ChangeTime   time.Time `json:"change_time" gorm:"autoCreateTime"`
	Reason       *string   `json:"reason" gorm:"size:200"`
	TenantID     int64     `json:"tenant_id" gorm:"index;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (POChangeLog) TableName() string {
	return "scp_po_change_log"
}

// RFQ 询价单头表
type RFQ struct {
	SCPBaseModel
	RFQNo               string    `json:"rfq_no" gorm:"size:50;uniqueIndex;not null"` // 询价单号
	RFQName              *string   `json:"rfq_name" gorm:"size:200"`                   // 询价单名称
	RFQType              string    `json:"rfq_type" gorm:"size:20;default:STANDARD"`    // STANDARD/QUICK/ANNUAL
	InquiryDate          time.Time `json:"inquiry_date" gorm:"type:date;not null"`       // 询价日期
	DeadlineDate         time.Time `json:"deadline_date" gorm:"type:date;not null"`      // 报价截止日期
	Currency             string    `json:"currency" gorm:"size:10;default:CNY"`         // 币种
	PaymentTerms         *string   `json:"payment_terms" gorm:"size:50"`                  // 付款条款
	DeliveryTerms        *string   `json:"delivery_terms" gorm:"size:100"`               // 交货条款
	QualityStandard      *string   `json:"quality_standard" gorm:"size:100"`             // 质量标准
	Status               string    `json:"status" gorm:"size:20;default:DRAFT"`          // DRAFT/PUBLISHED/CLOSED/AWARDED/CANCELLED
	AwardedSupplierID   *int64    `json:"awarded_supplier_id"`                          // 中标供应商
	AwardedTotalAmount  *float64  `json:"awarded_total_amount" gorm:"type:decimal(18,2)"` // 中标金额
	TotalBids            int       `json:"total_bids" gorm:"default:0"`                 // 收到报价数
	IsEvaluated          int       `json:"is_evaluated" gorm:"default:0"`                // 是否已评估
	EvaluationBy         *int64    `json:"evaluation_by"`                                // 评估人
	EvaluationTime       *time.Time `json:"evaluation_time"`                             // 评估时间
	Remark               *string   `json:"remark" gorm:"type:text"`                      // 备注
	CreatedBy            *string   `json:"created_by" gorm:"size:50"`                     // 创建人
	Items                []RFQItem  `json:"items" gorm:"foreignKey:RFQID"`                // 询价明细
	Invites              []RFQInvite `json:"invites" gorm:"foreignKey:RFQID"`            // 邀请供应商
}

func (RFQ) TableName() string {
	return "scp_rfq"
}

// RFQItem 询价单明细
type RFQItem struct {
	ID               uint       `json:"id" gorm:"primarykey"`
	RFQID            int64      `json:"rfq_id" gorm:"not null;index"`
	LineNo           int        `json:"line_no" gorm:"not null"`
	MaterialID       *int64     `json:"material_id"`
	MaterialCode     string     `json:"material_code" gorm:"size:50;not null"`
	MaterialName     *string    `json:"material_name" gorm:"size:100"`
	Specification    *string    `json:"specification" gorm:"size:200"`
	Unit             string     `json:"unit" gorm:"size:20;default:PCS"`
	RequiredQty      *float64   `json:"required_qty" gorm:"type:decimal(18,3)"`
	TargetPrice      *float64   `json:"target_price" gorm:"type:decimal(18,4)"`      // 目标单价
	MarketPrice      *float64   `json:"market_price" gorm:"type:decimal(18,4)"`     // 市场参考价
	RequestedDate    *time.Time `json:"requested_date" gorm:"type:date"`             // 要求交货日期
	QualityRequire   *string    `json:"quality_require" gorm:"type:text"`
	Remark           *string    `json:"remark" gorm:"type:text"`
	CreatedAt        time.Time  `json:"created_at" gorm:"autoCreateTime"`
}

func (RFQItem) TableName() string {
	return "scp_rfq_item"
}

// RFQInvite 询价邀请供应商
type RFQInvite struct {
	ID              uint       `json:"id" gorm:"primarykey"`
	RFQID           int64      `json:"rfq_id" gorm:"not null;index"`
	SupplierID      int64      `json:"supplier_id" gorm:"not null"`
	SupplierCode    string     `json:"supplier_code" gorm:"size:50"`
	SupplierName    string     `json:"supplier_name" gorm:"size:100"`
	ContactPerson   *string    `json:"contact_person" gorm:"size:50"`
	ContactEmail    *string    `json:"contact_email" gorm:"size:100"`
	InviteDate      *time.Time `json:"invite_date" gorm:"type:date"`
	ResponseStatus  string     `json:"response_status" gorm:"size:20;default:PENDING"` // PENDING/QUOTED/DECLINED
	QuotedDate      *time.Time `json:"quoted_date" gorm:"type:date"`
	TenantID        int64      `json:"tenant_id" gorm:"index;not null"`
	CreatedAt       time.Time  `json:"created_at" gorm:"autoCreateTime"`
}

func (RFQInvite) TableName() string {
	return "scp_rfq_invite"
}

// SupplierQuote 供应商报价
type SupplierQuote struct {
	SCPBaseModel
	RFQID               int64          `json:"rfq_id" gorm:"not null;index"`
	RFQNo               string         `json:"rfq_no" gorm:"size:50"`
	SupplierID          int64          `json:"supplier_id" gorm:"not null"`
	SupplierCode        string         `json:"supplier_code" gorm:"size:50"`
	SupplierName        string         `json:"supplier_name" gorm:"size:100"`
	QuoteNo             string         `json:"quote_no" gorm:"size:50;uniqueIndex;not null"` // 报价单号
	QuoteDate           time.Time      `json:"quote_date" gorm:"type:date;not null"`          // 报价日期
	ValidUntil          *time.Time     `json:"valid_until" gorm:"type:date"`                 // 报价有效期
	Currency            string         `json:"currency" gorm:"size:10;default:CNY"`
	PaymentTerms        *string        `json:"payment_terms" gorm:"size:50"`
	DeliveryDays        *int           `json:"delivery_days"`                                // 交货周期(天)
	TotalAmount         float64        `json:"total_amount" gorm:"type:decimal(18,2);default:0"`
	IsAccepted          int            `json:"is_accepted" gorm:"default:0"`                 // 是否中标
	IsLowest            int            `json:"is_lowest" gorm:"default:0"`                   // 是否最低价
	RankPosition        *int           `json:"rank_position"`                               // 排名
	QuoteStatus         string         `json:"quote_status" gorm:"size:20;default:SUBMITTED"` // SUBMITTED/REVISED/WITHDRAWN
	EvaluationScore     *float64       `json:"evaluation_score" gorm:"type:decimal(5,2)"`
	EvaluationResult    *string        `json:"evaluation_result" gorm:"size:20"`            // WIN/LOSE/PENDING
	EvaluatorID         *int64         `json:"evaluator_id"`
	EvaluationTime      *time.Time     `json:"evaluation_time"`
	EvaluationRemark    *string        `json:"evaluation_remark" gorm:"type:text"`
	Remark              *string        `json:"remark" gorm:"type:text"`
	Items               []QuoteItem    `json:"items" gorm:"foreignKey:QuoteID"`
}

func (SupplierQuote) TableName() string {
	return "scp_supplier_quote"
}

// QuoteItem 供应商报价明细
type QuoteItem struct {
	ID             uint       `json:"id" gorm:"primarykey"`
	QuoteID        int64      `json:"quote_id" gorm:"not null;index"`
	RFQLineID      *int64     `json:"rfq_line_id"`
	MaterialID     *int64     `json:"material_id"`
	MaterialCode   *string    `json:"material_code" gorm:"size:50"`
	MaterialName   *string    `json:"material_name" gorm:"size:100"`
	Unit           *string    `json:"unit" gorm:"size:20"`
	QuotedQty      *float64   `json:"quoted_qty" gorm:"type:decimal(18,3)"`
	UnitPrice      float64    `json:"unit_price" gorm:"type:decimal(18,4);not null"`
	LineAmount     float64    `json:"line_amount" gorm:"type:decimal(18,2);default:0"`
	DeliveryDate   *time.Time `json:"delivery_date" gorm:"type:date"`
	LeadTimeDays   *int       `json:"lead_time_days"`
	Remark         *string    `json:"remark" gorm:"type:text"`
	CreatedAt      time.Time  `json:"created_at" gorm:"autoCreateTime"`
}

func (QuoteItem) TableName() string {
	return "scp_quote_item"
}

// QuoteComparison 报价对比分析
type QuoteComparison struct {
	SCPBaseModel
	RFQID            int64           `json:"rfq_id" gorm:"not null;index"`
	ComparisonNo     string          `json:"comparison_no" gorm:"size:50;uniqueIndex;not null"`
	ComparedBy      *int64          `json:"compared_by"`
	ComparedAt      time.Time       `json:"compared_at" gorm:"autoCreateTime"`
	SummaryData     json.RawMessage `json:"summary_data" gorm:"type:jsonb"`           // 对比汇总数据
	Recommendation  *string         `json:"recommendation" gorm:"size:100"`          // 推荐供应商
	DecisionRemark  *string         `json:"decision_remark" gorm:"type:text"`
}

func (QuoteComparison) TableName() string {
	return "scp_quote_comparison"
}

// SCPSalesOrder 销售订单头表
type SCPSalesOrder struct {
	SCPBaseModel
	SONo              string           `json:"so_no" gorm:"size:50;uniqueIndex;not null"` // 销售订单号
	SOType            string           `json:"so_type" gorm:"size:20;default:STANDARD"`    // STANDARD/URGENT/DISTRIBUTION
	CustomerID        int64            `json:"customer_id" gorm:"not null"`
	CustomerCode      string           `json:"customer_code" gorm:"size:50"`
	CustomerName      string           `json:"customer_name" gorm:"size:100"`
	ContactPerson     *string          `json:"contact_person" gorm:"size:50"`
	ContactPhone      *string          `json:"contact_phone" gorm:"size:20"`
	ContactEmail      *string          `json:"contact_email" gorm:"size:100"`
	SalesPersonID     *int64           `json:"sales_person_id"`
	SalesPersonName   *string          `json:"sales_person_name" gorm:"size:50"`
	OrderDate         time.Time        `json:"order_date" gorm:"type:date;not null"`
	PromisedDate      *time.Time       `json:"promised_date" gorm:"type:date"`
	Currency          string           `json:"currency" gorm:"size:10;default:CNY"`
	PaymentTerms      *string          `json:"payment_terms" gorm:"size:50"`
	TaxRate           float64          `json:"tax_rate" gorm:"type:decimal(5,2);default:13.00"`
	TotalAmount       float64          `json:"total_amount" gorm:"type:decimal(18,2);default:0"`
	TotalQty          float64          `json:"total_qty" gorm:"type:decimal(18,3);default:0"`
	DeliveredAmount   float64          `json:"delivered_amount" gorm:"type:decimal(18,2);default:0"`
	DeliveredQty      float64          `json:"delivered_qty" gorm:"type:decimal(18,3);default:0"`
	ApprovedBy        *int64           `json:"approved_by"`
	ApprovedTime      *time.Time       `json:"approved_time"`
	ApprovalStatus    string           `json:"approval_status" gorm:"size:20;default:PENDING"`
	Status            string           `json:"status" gorm:"size:20;default:DRAFT"`     // DRAFT/PENDING/APPROVED/CONFIRMED/PARTIAL/SHIPPED/CLOSED/CANCELLED
	SourceType        *string          `json:"source_type" gorm:"size:20"`              // MANUAL/CUSTOMER_PO/CRM
	SourceNo          *string          `json:"source_no" gorm:"size:50"`
	DeliveryAddress   *string          `json:"delivery_address" gorm:"type:text"`
	DeliveryWarehouseID *int64         `json:"delivery_warehouse_id"`
	Remark            *string          `json:"remark" gorm:"type:text"`
	Items             []SCPSalesOrderItem `json:"items" gorm:"foreignKey:SOID"`
}

func (SCPSalesOrder) TableName() string {
	return "scp_sales_order"
}

// SCPSalesOrderItem 销售订单明细
type SCPSalesOrderItem struct {
	ID                  uint           `json:"id" gorm:"primarykey"`
	SOID                int64          `json:"so_id" gorm:"not null;index"`
	LineNo              int            `json:"line_no" gorm:"not null"`
	MaterialID          *int64         `json:"material_id"`
	MaterialCode        string         `json:"material_code" gorm:"size:50;not null"`
	MaterialName        *string        `json:"material_name" gorm:"size:100"`
	Specification       *string        `json:"specification" gorm:"size:200"`
	Unit                string         `json:"unit" gorm:"size:20;default:PCS"`
	UnitPrice           float64        `json:"unit_price" gorm:"type:decimal(18,4);default:0"`
	OrderQty            float64        `json:"order_qty" gorm:"type:decimal(18,3);not null"`
	DeliveredQty        float64        `json:"delivered_qty" gorm:"type:decimal(18,3);default:0"`
	ShippedQty          float64        `json:"shipped_qty" gorm:"type:decimal(18,3);default:0"`
	TaxAmount           float64        `json:"tax_amount" gorm:"type:decimal(18,2);default:0"`
	LineAmount          float64        `json:"line_amount" gorm:"type:decimal(18,2);default:0"`
	PromisedDate        *time.Time     `json:"promised_date" gorm:"type:date"`
	ActualDeliveryDate  *time.Time     `json:"actual_delivery_date" gorm:"type:date"`
	ProductionOrderID   *int64         `json:"production_order_id"`                        // 关联生产工单
	Status              string         `json:"status" gorm:"size:20;default:PENDING"`
	Remark              *string        `json:"remark" gorm:"type:text"`
	CreatedAt           time.Time      `json:"created_at" gorm:"autoCreateTime"`
}

func (SCPSalesOrderItem) TableName() string {
	return "scp_sales_order_item"
}

// SOChangeLog 销售订单变更记录
type SOChangeLog struct {
	ID           uint      `json:"id" gorm:"primarykey"`
	SOID         int64     `json:"so_id" gorm:"not null;index"`
	ChangeType   string    `json:"change_type" gorm:"size:30;not null"`
	ChangeField  *string   `json:"change_field" gorm:"size:50"`
	OldValue     *string   `json:"old_value" gorm:"size:200"`
	NewValue     *string   `json:"new_value" gorm:"size:200"`
	ChangedBy    *int64    `json:"changed_by"`
	ChangedByName *string  `json:"changed_by_name" gorm:"size:50"`
	ChangeTime   time.Time `json:"change_time" gorm:"autoCreateTime"`
	Reason       *string   `json:"reason" gorm:"size:200"`
	TenantID     int64     `json:"tenant_id" gorm:"index;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (SOChangeLog) TableName() string {
	return "scp_so_change_log"
}

// CustomerInquiry 客户询价头表
type CustomerInquiry struct {
	SCPBaseModel
	InquiryNo       string          `json:"inquiry_no" gorm:"size:50;uniqueIndex;not null"`
	InquiryDate     time.Time       `json:"inquiry_date" gorm:"type:date;not null"`
	CustomerID     *int64          `json:"customer_id"`
	CustomerName   *string         `json:"customer_name" gorm:"size:100"`
	ContactPerson  *string         `json:"contact_person" gorm:"size:50"`
	ContactPhone   *string         `json:"contact_phone" gorm:"size:20"`
	ContactEmail   *string         `json:"contact_email" gorm:"size:100"`
	ExpectedDate   *time.Time      `json:"expected_date" gorm:"type:date"`
	ValidUntil     *time.Time      `json:"valid_until" gorm:"type:date"`
	Currency       string          `json:"currency" gorm:"size:10;default:CNY"`
	Status         string          `json:"status" gorm:"size:20;default:DRAFT"` // DRAFT/SENT/QUOTED/WON/LOST/CANCELLED
	QuotedAmount   *float64        `json:"quoted_amount" gorm:"type:decimal(18,2)"`
	WinnerSupplierID *int64        `json:"winner_supplier_id"`                  // 中标供应商
	Remark         *string         `json:"remark" gorm:"type:text"`
	CreatedBy      *string         `json:"created_by" gorm:"size:50"`
	Items          []InquiryItem   `json:"items" gorm:"foreignKey:InquiryID"`
}

func (CustomerInquiry) TableName() string {
	return "scp_customer_inquiry"
}

// InquiryItem 客户询价明细
type InquiryItem struct {
	ID              uint       `json:"id" gorm:"primarykey"`
	InquiryID       int64      `json:"inquiry_id" gorm:"not null;index"`
	LineNo          int        `json:"line_no" gorm:"not null"`
	MaterialID      *int64     `json:"material_id"`
	MaterialCode    string     `json:"material_code" gorm:"size:50;not null"`
	MaterialName    *string    `json:"material_name" gorm:"size:100"`
	Specification   *string    `json:"specification" gorm:"size:200"`
	Unit            string     `json:"unit" gorm:"size:20;default:PCS"`
	RequiredQty     *float64   `json:"required_qty" gorm:"type:decimal(18,3)"`
	TargetPrice     *float64   `json:"target_price" gorm:"type:decimal(18,4)"`
	QuotedPrice     *float64   `json:"quoted_price" gorm:"type:decimal(18,4)"`
	QuotedSupplierID *int64    `json:"quoted_supplier_id"`                   // 询价供应商
	LeadTimeDays    *int       `json:"lead_time_days"`
	Remark          *string    `json:"remark" gorm:"type:text"`
	CreatedAt       time.Time  `json:"created_at" gorm:"autoCreateTime"`
}

func (InquiryItem) TableName() string {
	return "scp_inquiry_item"
}

// SupplierKPI 供应商绩效评分
type SupplierKPI struct {
	SCPBaseModel
	SupplierID              int64    `json:"supplier_id" gorm:"not null;index"`
	SupplierCode           string   `json:"supplier_code" gorm:"size:50"`
	SupplierName           string   `json:"supplier_name" gorm:"size:100"`
	EvaluationMonth        string   `json:"evaluation_month" gorm:"size:7;not null"` // 格式: 2026-01
	EvaluationDate        time.Time `json:"evaluation_date" gorm:"type:date;not null"`
	EvaluatedBy            *int64   `json:"evaluated_by"`
	EvaluatedByName        *string  `json:"evaluated_by_name" gorm:"size:50"`

	// 交货绩效
	OnTimeDeliveryRate     *float64 `json:"on_time_delivery_rate" gorm:"type:decimal(5,2)"` // 准时交货率(%)
	TotalDeliveryOrders    int      `json:"total_delivery_orders" gorm:"default:0"`
	OnTimeDeliveryCount    int      `json:"on_time_delivery_count" gorm:"default:0"`
	AvgDelayDays           *float64 `json:"avg_delay_days" gorm:"type:decimal(5,2)"`       // 平均延迟天数

	// 质量绩效
	QualityPassRate        *float64 `json:"quality_pass_rate" gorm:"type:decimal(5,2)"`    // 来料合格率(%)
	TotalIQCOrders         int      `json:"total_iqc_orders" gorm:"default:0"`
	PassedIQCOrders        int      `json:"passed_iqc_orders" gorm:"default:0"`
	DefectPartsCount       int      `json:"defect_parts_count" gorm:"default:0"`
	DefectRate             *float64 `json:"defect_rate" gorm:"type:decimal(5,4)"`

	// 价格绩效
	PriceCompetitiveness   *float64 `json:"price_competitiveness" gorm:"type:decimal(5,2)"` // 价格竞争力(%)
	LastPurchasePrice      *float64 `json:"last_purchase_price" gorm:"type:decimal(18,4)"`
	MarketAvgPrice         *float64 `json:"market_avg_price" gorm:"type:decimal(18,4)"`

	// 综合评分
	TotalScore             float64  `json:"total_score" gorm:"type:decimal(5,2);default:0"` // 综合得分
	Grade                  *string  `json:"grade" gorm:"size:10"`                           // A/B/C/D
	RankPosition           *int     `json:"rank_position"`                                   // 排名
	EvaluationRemark       *string  `json:"evaluation_remark" gorm:"type:text"`
}

func (SupplierKPI) TableName() string {
	return "scp_supplier_kpi"
}

// SupplierDeliveryRecord 供应商绩效明细-交货记录
type SupplierDeliveryRecord struct {
	ID               uint       `json:"id" gorm:"primarykey"`
	SupplierID       int64      `json:"supplier_id" gorm:"not null;index"`
	PONo             *string    `json:"po_no" gorm:"size:50"`
	POLineID         *int64     `json:"po_line_id"`
	PromisedDate     *time.Time `json:"promised_date" gorm:"type:date"`
	ActualDeliveryDate *time.Time `json:"actual_delivery_date" gorm:"type:date"`
	DeliveryStatus   *string    `json:"delivery_status" gorm:"size:20"`   // ON_TIME/DELAYED/EARLY/PARTIAL
	DelayDays        int        `json:"delay_days" gorm:"default:0"`
	IsPenaltyApplied int        `json:"is_penalty_applied" gorm:"default:0"` // 是否扣款
	PenaltyAmount    *float64   `json:"penalty_amount" gorm:"type:decimal(18,2)"`
	Remark           *string    `json:"remark" gorm:"size:200"`
	TenantID         int64      `json:"tenant_id" gorm:"index;not null"`
	CreatedAt        time.Time  `json:"created_at" gorm:"autoCreateTime"`
}

func (SupplierDeliveryRecord) TableName() string {
	return "scp_supplier_delivery_record"
}

// SupplierQualityRecord 供应商绩效明细-质量记录
type SupplierQualityRecord struct {
	ID             uint           `json:"id" gorm:"primarykey"`
	SupplierID     int64          `json:"supplier_id" gorm:"not null;index"`
	MaterialCode   *string        `json:"material_code" gorm:"size:50"`
	IQCRecordID    *int64         `json:"iqc_record_id"`
	IQCNo          *string        `json:"iqc_no" gorm:"size:50"`
	InspectDate    *time.Time     `json:"inspect_date" gorm:"type:date"`
	InspectQty     *float64       `json:"inspect_qty" gorm:"type:decimal(18,3)"`
	QualifiedQty   *float64       `json:"qualified_qty" gorm:"type:decimal(18,3)"`
	DefectQty      *float64       `json:"defect_qty" gorm:"type:decimal(18,3)"`
	DefectRate     *float64       `json:"defect_rate" gorm:"type:decimal(5,4)"`
	DefectTypes    json.RawMessage `json:"defect_types" gorm:"type:jsonb"` // [{code, name, qty}]
	IsChargeback   int            `json:"is_chargeback" gorm:"default:0"`  // 是否索赔
	ChargebackAmount *float64     `json:"chargeback_amount" gorm:"type:decimal(18,2)"`
	Remark         *string        `json:"remark" gorm:"size:200"`
	TenantID       int64          `json:"tenant_id" gorm:"index;not null"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime"`
}

func (SupplierQualityRecord) TableName() string {
	return "scp_supplier_quality_record"
}

// SupplierGradeStandard 供应商评级标准
type SupplierGradeStandard struct {
	ID                    uint       `json:"id" gorm:"primarykey"`
	Grade                 string     `json:"grade" gorm:"size:10;not null"` // A/B/C/D
	GradeName             *string    `json:"grade_name" gorm:"size:50"`
	MinScore              float64    `json:"min_score" gorm:"type:decimal(5,2);not null"`
	MaxScore              float64    `json:"max_score" gorm:"type:decimal(5,2);not null"`
	OnTimeRateThreshold   *float64   `json:"on_time_rate_threshold" gorm:"type:decimal(5,2)"`  // 准时率门槛
	QualityRateThreshold *float64   `json:"quality_rate_threshold" gorm:"type:decimal(5,2)"` // 合格率门槛
	IsActive              int        `json:"is_active" gorm:"default:1"`
	TenantID              int64      `json:"tenant_id" gorm:"index;not null"`
	CreatedAt             time.Time  `json:"created_at" gorm:"autoCreateTime"`
}

func (SupplierGradeStandard) TableName() string {
	return "scp_supplier_grade_standard"
}

// SupplierPurchaseInfo 供应商采购信息扩展
type SupplierPurchaseInfo struct {
	ID                    uint       `json:"id" gorm:"primarykey"`
	SupplierID            int64      `json:"supplier_id" gorm:"uniqueIndex;not null"`
	PaymentTerms          *string    `json:"payment_terms" gorm:"size:50"`
	CreditLimit           *float64   `json:"credit_limit" gorm:"type:decimal(18,2)"`
	TaxRate               float64    `json:"tax_rate" gorm:"type:decimal(5,2);default:13.00"`
	MinOrderAmount        *float64   `json:"min_order_amount" gorm:"type:decimal(18,2)"`
	LeadTimeDays          *int       `json:"lead_time_days"`                    // 标准交货周期
	SupplierGrade         *string    `json:"supplier_grade" gorm:"size:10"`     // 供应商等级
	IsPreferred           int        `json:"is_preferred" gorm:"default:0"`    // 是否首选供应商
	IsBlacklist           int        `json:"is_blacklist" gorm:"default:0"`   // 是否黑名单
	BlacklistReason       *string    `json:"blacklist_reason" gorm:"size:200"`
	CooperationStartDate  *time.Time `json:"cooperation_start_date" gorm:"type:date"`
	CooperationEndDate    *time.Time `json:"cooperation_end_date" gorm:"type:date"`
	TotalCooperationAmount *float64  `json:"total_cooperation_amount" gorm:"type:decimal(18,2)"` // 累计合作金额
	TenantID              int64      `json:"tenant_id" gorm:"index;not null"`
	CreatedAt             time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt             time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

func (SupplierPurchaseInfo) TableName() string {
	return "scp_supplier_purchase_info"
}

// SupplierMaterial 供应商物料关联
type SupplierMaterial struct {
	ID             uint      `json:"id" gorm:"primarykey"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	SupplierID     int64     `json:"supplier_id" gorm:"index;not null"`
	MaterialID     int64     `json:"material_id" gorm:"index;not null"`
	MaterialCode   *string   `json:"material_code" gorm:"size:50"`
	MaterialName   *string   `json:"material_name" gorm:"size:100"`
	SupplierPartNo *string   `json:"supplier_part_no" gorm:"size:50"` // 供应商料号
	Price          float64   `json:"price" gorm:"type:decimal(18,6)"`
	Currency       *string   `json:"currency" gorm:"size:10"`         // 币种
	MinOrderQty    float64   `json:"min_order_qty" gorm:"default:0"`  // 最小订购量
	LeadTime       int       `json:"lead_time" gorm:"default:0"`      // 交期(天)
	IsPreferred    int       `json:"is_preferred" gorm:"default:0"`   // 是否首选 0-否 1-是
	Status         int       `json:"status" gorm:"default:1"`       // 1-有效 0-无效
	TenantID       int64     `json:"tenant_id" gorm:"index;not null"`
}

func (SupplierMaterial) TableName() string {
	return "mdm_supplier_material"
}
