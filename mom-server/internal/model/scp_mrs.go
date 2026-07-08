package model

import (
	"time"

	"mom-server/internal/pkg/status"
)

// ScpMRS MRS汇总表
type ScpMRS struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	TenantID     int64          `json:"tenant_id" gorm:"index;not null"`
	MrsNo        string         `json:"mrs_no" gorm:"size:50;not null"`        // MRS单号
	PlanMonth    string         `json:"plan_month" gorm:"size:7"`              // 计划月份 YYYY-MM
	Status       string         `json:"status" gorm:"size:20;default:'DRAFT'"` // DRAFT/PUBLISHED/CLOSED(legacy,保留双轨)
	StatusV2     string         `json:"status_v2" gorm:"size:30;index"`        // 双轨:varchar - MOM 3.0 V2.1(P2 扩展 batch 2)
	SourceType   string         `json:"source_type" gorm:"size:20"`           // MANUAL/APS
	SourceNo     string         `json:"source_no" gorm:"size:50"`              // 来源单号
	TotalItems   int            `json:"total_items" gorm:"default:0"`
	TotalQty     float64        `json:"total_qty" gorm:"type:decimal(18,3)"`
	PublishedBy  *int64         `json:"published_by"`
	PublishedAt  *time.Time     `json:"published_at"`
	Remark       string         `json:"remark" gorm:"type:text"`
	Items        []ScpMRSItem   `json:"items" gorm:"foreignKey:MrsID"`
}

func (ScpMRS) TableName() string {
	return "scp_mrs"
}

// StatusCode 返回有效状态码(优先 status_v2, fallback status → 字典码)
// P2 扩展 batch 2,与 MdmBOM.StatusCode() 同模式
func (s *ScpMRS) StatusCode() string {
	if s.StatusV2 != "" {
		return s.StatusV2
	}
	return string(status.ScpMrsFromLegacyVarchar(s.Status))
}

// SetStatus 同时写 status + status_v2(双轨制)
func (s *ScpMRS) SetStatus(code string) {
	c := status.Code(code)
	if !c.IsValid(status.ScpMrsAll) {
		code = string(status.ScpMrsDraft)
	}
	s.Status = code
	s.StatusV2 = code
}

// ScpMRSItem MRS明细
type ScpMRSItem struct {
	ID           uint       `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	MrsID        int64      `json:"mrs_id" gorm:"index;not null"`
	MrsNo        string     `json:"mrs_no" gorm:"size:50;not null"`
	MaterialID   int64      `json:"material_id" gorm:""`
	MaterialCode string     `json:"material_code" gorm:"size:50"`
	MaterialName string     `json:"material_name" gorm:"size:100"`
	Spec         string     `json:"spec" gorm:"size:200"`
	Unit         string     `json:"unit" gorm:"size:20"`
	ReqQty       float64    `json:"req_qty" gorm:"type:decimal(18,3)"`  // 需求数量
	OnHandQty    float64    `json:"on_hand_qty" gorm:"type:decimal(18,3)"` // 库存数量
	ShortQty     float64    `json:"short_qty" gorm:"type:decimal(18,3)"`  // 短缺数量
	SupplierID   *int64     `json:"supplier_id"`
	SupplierName string     `json:"supplier_name" gorm:"size:100"`
	PromiseDate  *time.Time `json:"promise_date" gorm:"type:date"`
	Status       string     `json:"status" gorm:"size:20;default:'PENDING'"` // PENDING/PRINTED/SENT
}

func (ScpMRSItem) TableName() string {
	return "scp_mrs_item"
}

// ScpMRSCreateReqVO 创建MRS请求
type ScpMRSCreateReqVO struct {
	PlanMonth  string                   `json:"planMonth"`
	SourceType string                   `json:"sourceType"`
	SourceNo   string                   `json:"sourceNo"`
	Remark     string                   `json:"remark"`
	Items      []ScpMRSItemCreateReqVO  `json:"items"`
}

// ScpMRSItemCreateReqVO 创建MRS明细请求
type ScpMRSItemCreateReqVO struct {
	MaterialID   int64    `json:"materialId"`
	MaterialCode string   `json:"materialCode"`
	MaterialName string   `json:"materialName"`
	Spec         string   `json:"spec"`
	Unit         string   `json:"unit"`
	ReqQty       float64  `json:"reqQty"`
	OnHandQty    float64  `json:"onHandQty"`
	ShortQty     float64  `json:"shortQty"`
	SupplierID   *int64   `json:"supplierId"`
	SupplierName string   `json:"supplierName"`
	PromiseDate  string   `json:"promiseDate"`
}

// ScpMRSUpdateReqVO 更新MRS请求
type ScpMRSUpdateReqVO struct {
	PlanMonth  string                   `json:"planMonth"`
	SourceType string                   `json:"sourceType"`
	SourceNo   string                   `json:"sourceNo"`
	Remark     string                   `json:"remark"`
	Items      []ScpMRSItemCreateReqVO `json:"items"`
}

// ScpMRSRespVO MRS响应
type ScpMRSRespVO struct {
	Id         int64                 `json:"id"`
	MrsNo      string                `json:"mrsNo"`
	PlanMonth  string                `json:"planMonth"`
	Status     string                `json:"status"`
	TotalItems int                   `json:"totalItems"`
	TotalQty   float64               `json:"totalQty"`
	Items      []ScpMRSItemRespVO    `json:"items"`
}

// ScpMRSItemRespVO MRS明细响应
type ScpMRSItemRespVO struct {
	Id           int64      `json:"id"`
	MaterialID   int64      `json:"materialId"`
	MaterialCode string     `json:"materialCode"`
	MaterialName string     `json:"materialName"`
	Spec         string     `json:"spec"`
	Unit         string     `json:"unit"`
	ReqQty       float64    `json:"reqQty"`
	OnHandQty    float64    `json:"onHandQty"`
	ShortQty     float64    `json:"shortQty"`
	SupplierID   *int64     `json:"supplierId"`
	SupplierName string     `json:"supplierName"`
	PromiseDate  *time.Time `json:"promiseDate"`
	Status       string     `json:"status"`
}

// ScpMRSSyncReqVO 同步MRS请求
type ScpMRSSyncReqVO struct {
	SourceNo string `json:"sourceNo"`
	Data     []struct {
		MaterialCode string  `json:"materialCode"`
		MaterialName string  `json:"materialName"`
		Spec         string  `json:"spec"`
		Unit         string  `json:"unit"`
		ReqQty       float64 `json:"reqQty"`
		OnHandQty    float64 `json:"onHandQty"`
		ShortQty     float64 `json:"shortQty"`
		SupplierName string  `json:"supplierName"`
		PromiseDate  string  `json:"promiseDate"`
	} `json:"data"`
}
