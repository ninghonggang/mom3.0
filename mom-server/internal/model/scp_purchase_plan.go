package model

import (
	"time"
)

// ScpPurchasePlan 采购计划主表
type ScpPurchasePlan struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	TenantID     int64         `json:"tenant_id" gorm:"index;not null"`
	PlanNo       string        `json:"plan_no" gorm:"size:50;not null"`     // 计划单号
	Title        string        `json:"title" gorm:"size:200"`                // 计划标题
	PlanType     string        `json:"plan_type" gorm:"size:20"`            // 计划类型: MONTHLY/QUARTERLY/YEARLY
	PlanYear     int           `json:"plan_year" gorm:""`                   // 计划年份
	PlanMonth    int           `json:"plan_month" gorm:""`                  // 计划月份(1-12)
	Quarter      int           `json:"quarter" gorm:""`                     // 季度(1-4)
	Status       string        `json:"status" gorm:"size:20;default:'DRAFT'"` // DRAFT/CONFIRMED/PUBLISHED/CLOSED
	TotalItems   int           `json:"total_items" gorm:"default:0"`         // 明细行数
	TotalAmount  float64       `json:"total_amount" gorm:"type:decimal(18,2)"` // 计划总金额
	Currency     string        `json:"currency" gorm:"size:10;default:'CNY'"` // 币种
	Department   string        `json:"department" gorm:"size:100"`           // 需求部门
	SubmitterID  *int64        `json:"submitter_id"`                         // 提交人
	SubmittedAt  *time.Time    `json:"submitted_at"`                         // 提交时间
	ConfirmedBy  *int64        `json:"confirmed_by"`                         // 确认人
	ConfirmedAt  *time.Time    `json:"confirmed_at"`                         // 确认时间
	PublishedBy  *int64        `json:"published_by"`                         // 发布人
	PublishedAt  *time.Time    `json:"published_at"`                         // 发布时间
	ClosedBy     *int64        `json:"closed_by"`                             // 关闭人
	ClosedAt     *time.Time    `json:"closed_at"`                             // 关闭时间
	CloseReason  string        `json:"close_reason" gorm:"size:500"`          // 关闭原因
	Remark       string        `json:"remark" gorm:"type:text"`              // 备注
	Items        []ScpPurchasePlanItem `json:"items" gorm:"foreignKey:PlanID"`
}

func (ScpPurchasePlan) TableName() string {
	return "scp_purchase_plan"
}

// ScpPurchasePlanItem 采购计划明细表
type ScpPurchasePlanItem struct {
	ID              uint       `json:"id" gorm:"primarykey"`
	CreatedAt       time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	PlanID          int64      `json:"plan_id" gorm:"index;not null"`      // 计划主表ID
	PlanNo          string     `json:"plan_no" gorm:"size:50;not null"`     // 计划单号
	LineNo          int        `json:"line_no" gorm:""`                     // 行号
	MaterialID      *int64     `json:"material_id"`                         // 物料ID
	MaterialCode    string     `json:"material_code" gorm:"size:50"`       // 物料编码
	MaterialName    string     `json:"material_name" gorm:"size:100"`       // 物料名称
	Spec            string     `json:"spec" gorm:"size:200"`                // 规格型号
	Unit            string     `json:"unit" gorm:"size:20"`                  // 单位
	ReqQty          float64    `json:"req_qty" gorm:"type:decimal(18,3)"`   // 需求数量
	UnitPrice       float64    `json:"unit_price" gorm:"type:decimal(18,6)"` // 预计单价
	LineAmount      float64    `json:"line_amount" gorm:"type:decimal(18,2)"` // 行金额
	SupplierID      *int64     `json:"supplier_id"`                         // 供应商ID
	SupplierCode    string     `json:"supplier_code" gorm:"size:50"`         // 供应商编码
	SupplierName    string     `json:"supplier_name" gorm:"size:100"`        // 供应商名称
	ReqDeliveryDate *time.Time `json:"req_delivery_date" gorm:"type:date"`   // 需求交期
	PromiseDate     *time.Time `json:"promise_date" gorm:"type:date"`        // 承诺交期
	MrsNo           string     `json:"mrs_no" gorm:"size:50"`                // 来源MRS单号
	MrsLineNo       int        `json:"mrs_line_no"`                         // 来源MRS行号
	Status          string     `json:"status" gorm:"size:20;default:'PENDING'"` // PENDING/ORDERED/RECEIVED/CLOSED
	Remark          string     `json:"remark" gorm:"type:text"`              // 备注
}

func (ScpPurchasePlanItem) TableName() string {
	return "scp_purchase_plan_item"
}

// ScpPurchasePlanCreateReqVO 创建采购计划请求
type ScpPurchasePlanCreateReqVO struct {
	Title    string                      `json:"title"`
	PlanType string                      `json:"planType" binding:"required"`
	PlanYear int                         `json:"planYear" binding:"required"`
	PlanMonth int                        `json:"planMonth"`
	Quarter  int                         `json:"quarter"`
	Currency string                      `json:"currency"`
	Department string                    `json:"department"`
	Remark   string                      `json:"remark"`
	Items    []ScpPurchasePlanItemCreateReqVO `json:"items"`
}

// ScpPurchasePlanItemCreateReqVO 创建采购计划明细请求
type ScpPurchasePlanItemCreateReqVO struct {
	MaterialID      int64   `json:"materialId"`
	MaterialCode    string  `json:"materialCode" binding:"required"`
	MaterialName    string  `json:"materialName" binding:"required"`
	Spec            string  `json:"spec"`
	Unit            string  `json:"unit"`
	ReqQty          float64 `json:"reqQty" binding:"required"`
	UnitPrice       float64 `json:"unitPrice"`
	SupplierID      int64   `json:"supplierId"`
	SupplierCode    string  `json:"supplierCode"`
	SupplierName    string  `json:"supplierName"`
	ReqDeliveryDate string  `json:"reqDeliveryDate"`
	PromiseDate     string  `json:"promiseDate"`
	MrsNo           string  `json:"mrsNo"`
	MrsLineNo       int     `json:"mrsLineNo"`
	Remark          string  `json:"remark"`
}

// ScpPurchasePlanUpdateReqVO 更新采购计划请求
type ScpPurchasePlanUpdateReqVO struct {
	Title    string                      `json:"title"`
	PlanType string                      `json:"planType"`
	PlanYear int                         `json:"planYear"`
	PlanMonth int                        `json:"planMonth"`
	Quarter  int                         `json:"quarter"`
	Currency string                      `json:"currency"`
	Department string                    `json:"department"`
	Remark   string                      `json:"remark"`
	Items    []ScpPurchasePlanItemCreateReqVO `json:"items"`
}

// ScpPurchasePlanRespVO 采购计划响应
type ScpPurchasePlanRespVO struct {
	ID          int64                            `json:"id"`
	PlanNo      string                           `json:"planNo"`
	Title       string                           `json:"title"`
	PlanType    string                           `json:"planType"`
	PlanYear    int                              `json:"planYear"`
	PlanMonth   int                              `json:"planMonth"`
	Quarter     int                              `json:"quarter"`
	Status      string                           `json:"status"`
	TotalItems  int                              `json:"totalItems"`
	TotalAmount float64                          `json:"totalAmount"`
	Currency    string                           `json:"currency"`
	Department  string                           `json:"department"`
	Items       []ScpPurchasePlanItemRespVO      `json:"items"`
}

// ScpPurchasePlanItemRespVO 采购计划明细响应
type ScpPurchasePlanItemRespVO struct {
	ID              int64      `json:"id"`
	LineNo          int        `json:"lineNo"`
	MaterialID      int64      `json:"materialId"`
	MaterialCode    string     `json:"materialCode"`
	MaterialName    string     `json:"materialName"`
	Spec            string     `json:"spec"`
	Unit            string     `json:"unit"`
	ReqQty          float64    `json:"reqQty"`
	UnitPrice       float64    `json:"unitPrice"`
	LineAmount      float64    `json:"lineAmount"`
	SupplierID      int64      `json:"supplierId"`
	SupplierCode    string     `json:"supplierCode"`
	SupplierName    string     `json:"supplierName"`
	ReqDeliveryDate *time.Time `json:"reqDeliveryDate"`
	PromiseDate     *time.Time `json:"promiseDate"`
	MrsNo           string     `json:"mrsNo"`
	MrsLineNo       int        `json:"mrsLineNo"`
	Status          string     `json:"status"`
}
