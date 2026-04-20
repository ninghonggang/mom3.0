package model

import (
	"time"
)

// WMSPickJob 拣货作业单
type WMSPickJob struct {
	BaseModel
	TenantID    int64      `json:"tenant_id" gorm:"index;not null"`
	PickNo      string     `json:"pick_no" gorm:"size:50;not null"`
	PickType    string     `json:"pick_type" gorm:"size:20"`           // PICK/SHORT
	SourceType  string     `json:"source_type" gorm:"size:20"`        // SO/PO/TRANSFER
	SourceNo    string     `json:"source_no" gorm:"size:50"`          // 来源单号
	WarehouseID int64      `json:"warehouse_id" gorm:"index"`
	WarehouseName string   `json:"warehouse_name" gorm:"size:100"`
	Status      string     `json:"status" gorm:"size:20;default:'PENDING'"` // PENDING/ASSIGNED/PICKING/COMPLETED/CANCELLED
	AssignTime  *time.Time `json:"assign_time"`
	PickerID    *int64     `json:"picker_id" gorm:"index"`
	PickerName  string     `json:"picker_name" gorm:"size:50"`
	PickedTime  *time.Time `json:"picked_time"`
	Remark      string     `json:"remark" gorm:"type:text"`
}

func (WMSPickJob) TableName() string {
	return "wms_pick_job"
}

// WMSPickRecord 拣货明细
type WMSPickRecord struct {
	BaseModel
	TenantID    int64   `json:"tenant_id" gorm:"index;not null"`
	PickJobID   int64   `json:"pick_job_id" gorm:"index;not null"`
	PickNo      string  `json:"pick_no" gorm:"size:50;not null"`
	ItemID      int64   `json:"item_id" gorm:"index"`
	ItemCode    string  `json:"item_code" gorm:"size:50"`
	ItemName    string  `json:"item_name" gorm:"size:100"`
	LocationID  int64   `json:"location_id" gorm:"index"`
	LocationCode string `json:"location_code" gorm:"size:50"`
	PickQty     float64 `json:"pick_qty" gorm:"type:decimal(18,3)"`              // 拣货数量
	PickedQty   float64 `json:"picked_qty" gorm:"type:decimal(18,3);default:0"` // 已拣数量
	Status      string  `json:"status" gorm:"size:20;default:'PENDING'"`
}

func (WMSPickRecord) TableName() string {
	return "wms_pick_record"
}

// ========== Request/Response VO ==========

// WMSPickJobCreateReqVO 创建拣货作业请求
type WMSPickJobCreateReqVO struct {
	SourceType   string `json:"sourceType"`
	SourceNo     string `json:"sourceNo"`
	WarehouseID  int64  `json:"warehouseId"`
}

// WMSPickJobAssignReqVO 分配拣货人请求
type WMSPickJobAssignReqVO struct {
	Id         int64  `json:"id"`
	PickerID   int64  `json:"pickerId"`
	PickerName string `json:"pickerName"`
}

// WMSPickJobCancelReqVO 取消拣货作业请求
type WMSPickJobCancelReqVO struct {
	Id     int64  `json:"id"`
	Reason string `json:"reason"`
}

// WMSPickRecordRespVO 拣货明细响应
type WMSPickRecordRespVO struct {
	Id           int64   `json:"id"`
	ItemID       int64   `json:"itemId"`
	ItemCode     string  `json:"itemCode"`
	ItemName     string  `json:"itemName"`
	LocationID   int64   `json:"locationId"`
	LocationCode string  `json:"locationCode"`
	PickQty      float64 `json:"pickQty"`
	PickedQty    float64 `json:"pickedQty"`
	Status       string  `json:"status"`
}

// WMSPickJobRespVO 拣货作业响应
type WMSPickJobRespVO struct {
	Id           int64                  `json:"id"`
	PickNo       string                 `json:"pickNo"`
	PickType     string                 `json:"pickType"`
	SourceType   string                 `json:"sourceType"`
	SourceNo     string                 `json:"sourceNo"`
	WarehouseID  int64                  `json:"warehouseId"`
	WarehouseName string                `json:"warehouseName"`
	Status       string                 `json:"status"`
	PickerID     *int64                 `json:"pickerId"`
	PickerName   string                 `json:"pickerName"`
	AssignTime   *time.Time             `json:"assignTime"`
	PickedTime   *time.Time             `json:"pickedTime"`
	Remark       string                 `json:"remark"`
	Records      []WMSPickRecordRespVO  `json:"records"`
}
