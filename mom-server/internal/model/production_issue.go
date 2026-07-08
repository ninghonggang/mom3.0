package model

import (
	"time"

	"mom-server/internal/pkg/status"
)

// ProductionIssue 生产发料单
type ProductionIssue struct {
	ID            uint      `json:"id" gorm:"primarykey"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	IssueNo       string    `json:"issue_no" gorm:"size:50;uniqueIndex;not null"`
	IssueType     string    `json:"issue_type" gorm:"size:20;not null"` // NORMAL/SUPPLEMENT/CALL
	ProductionOrderID int64 `json:"production_order_id" gorm:"not null;index"`
	OrderNo       *string   `json:"order_no" gorm:"size:50"`
	WorkstationID *int64    `json:"workstation_id" gorm:"index"`
	WorkshopID    *int64    `json:"workshop_id" gorm:"index"`
	Status        string    `json:"status" gorm:"size:20;not null;default:PENDING"`           // PENDING/APPROVED/PICKING/PICKED/ISSUED/CANCELLED(legacy,保留双轨)
	StatusV2      string    `json:"status_v2" gorm:"size:30;index"`                          // 双轨:varchar - MOM 3.0 V2.1(P2 扩展)
	PickStatus    string    `json:"pick_status" gorm:"size:20;default:PENDING"`              // PENDING/PICKING/PICKED
	RequestBy     *int64    `json:"request_by"`
	RequestTime   *time.Time `json:"request_time"`
	IssuedBy      *int64    `json:"issued_by"`
	IssuedTime    *time.Time `json:"issued_time"`
	Remark        *string   `json:"remark" gorm:"type:text"`
	TenantID      int64     `json:"tenant_id" gorm:"index;not null"`
	CreatedBy     *string   `json:"created_by" gorm:"size:50"`
	Items         []ProductionIssueItem `json:"items" gorm:"foreignKey:IssueID"`
}

// StatusCode 返回有效状态码(优先 status_v2, fallback status → 字典码)
// 与 MdmBOM.StatusCode() 同模式(P2 扩展,统一 MOM 3.0 切读路径)
func (i *ProductionIssue) StatusCode() string {
	if i.StatusV2 != "" {
		return i.StatusV2
	}
	return string(status.ProductionIssueFromLegacyVarchar(i.Status))
}

// SetStatus 同时写 status + status_v2(双轨制)
func (i *ProductionIssue) SetStatus(code string) {
	c := status.Code(code)
	if !c.IsValid(status.ProductionIssueAll) {
		code = string(status.ProductionIssuePending)
	}
	i.Status = code
	i.StatusV2 = code
}

func (ProductionIssue) TableName() string {
	return "wms_production_issue"
}

// ProductionIssueItem 生产发料明细
type ProductionIssueItem struct {
	ID            uint    `json:"id" gorm:"primarykey"`
	IssueID      uint   `json:"issue_id" gorm:"not null;index"`
	LineNo       int     `json:"line_no" gorm:"not null"`
	MaterialID   int64   `json:"material_id" gorm:"not null;index"`
	MaterialCode *string `json:"material_code" gorm:"size:50"`
	MaterialName *string `json:"material_name" gorm:"size:100"`
	Unit         *string `json:"unit" gorm:"size:20"`
	RequiredQty  float64 `json:"required_qty" gorm:"type:decimal(18,3);not null"`
	PickedQty    float64 `json:"picked_qty" gorm:"type:decimal(18,3);default:0"`
	IssuedQty    float64 `json:"issued_qty" gorm:"type:decimal(18,3);default:0"`
	WarehouseID  *int64  `json:"warehouse_id"`
	LocationID   *int64  `json:"location_id"`
	BatchNo      *string `json:"batch_no" gorm:"size:50"`
	Remark       *string `json:"remark" gorm:"type:text"`
	TenantID     int64   `json:"tenant_id" gorm:"index;not null"`
}

func (ProductionIssueItem) TableName() string {
	return "wms_production_issue_item"
}

// ProductionIssueCreate 创建请求
type ProductionIssueCreate struct {
	IssueType        string `json:"issue_type" binding:"required"`
	ProductionOrderID int64 `json:"production_order_id" binding:"required"`
	WorkstationID   *int64 `json:"workstation_id"`
	WorkshopID      *int64 `json:"workshop_id"`
	Items           []ProductionIssueItemCreate `json:"items" binding:"required,min=1"`
}

// ProductionIssueItemCreate 明细创建请求
type ProductionIssueItemCreate struct {
	MaterialID   int64   `json:"material_id" binding:"required"`
	MaterialCode *string `json:"material_code"`
	MaterialName *string `json:"material_name"`
	Unit         *string `json:"unit"`
	RequiredQty  float64 `json:"required_qty" binding:"required"`
	WarehouseID  *int64  `json:"warehouse_id"`
	LocationID   *int64  `json:"location_id"`
	BatchNo      *string `json:"batch_no"`
	Remark       *string `json:"remark"`
}

// ProductionIssueUpdate 更新请求
type ProductionIssueUpdate struct {
	WorkstationID *int64 `json:"workstation_id"`
	WorkshopID    *int64 `json:"workshop_id"`
	Remark        *string `json:"remark"`
}

// ProductionIssueSubmit 提交请求
type ProductionIssueSubmit struct {
	Items []ProductionIssueItemSubmit `json:"items"`
}

// ProductionIssueItemSubmit 明细提交
type ProductionIssueItemSubmit struct {
	MaterialID  int64   `json:"material_id" binding:"required"`
	PickedQty  float64 `json:"picked_qty" binding:"required"`
	WarehouseID *int64 `json:"warehouse_id"`
	LocationID *int64  `json:"location_id"`
	BatchNo    *string `json:"batch_no"`
}
