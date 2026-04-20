package model

import (
	"time"
)

// WMSPutawayJob 上架作业单
type WMSPutawayJob struct {
	BaseModel
	TenantID     int64      `json:"tenant_id" gorm:"index;not null"`
	PutawayNo    string     `json:"putaway_no" gorm:"size:50;not null"`   // 上架单号
	SourceType   string     `json:"source_type" gorm:"size:20"`            // RECEIVE/RETURN/TRANSFER
	SourceNo     string     `json:"source_no" gorm:"size:50"`             // 来源单号
	WarehouseID  int64      `json:"warehouse_id" gorm:"index"`
	Status       string     `json:"status" gorm:"size:20;default:'PENDING'"` // PENDING/ASSIGNED/PUTAWAYING/COMPLETED/CANCELLED
	AssignTime   *time.Time `json:"assign_time"`
	OperatorID   *int64     `json:"operator_id"`
	OperatorName string     `json:"operator_name" gorm:"size:50"`
	PutawayTime  *time.Time `json:"putaway_time"`
	Remark       string     `json:"remark" gorm:"type:text"`
}

func (WMSPutawayJob) TableName() string {
	return "wms_putaway_job"
}

// WMSPutawayRecord 上架明细
type WMSPutawayRecord struct {
	BaseModel
	TenantID       int64   `json:"tenant_id" gorm:"index;not null"`
	PutawayJobID   int64   `json:"putaway_job_id" gorm:"index;not null"`
	PutawayNo     string  `json:"putaway_no" gorm:"size:50;not null"`
	ItemID        int64   `json:"item_id"`
	ItemCode      string  `json:"item_code" gorm:"size:50"`
	ItemName      string  `json:"item_name" gorm:"size:100"`
	FromLocationID *int64 `json:"from_location_id"` // 来源库位
	ToLocationID  *int64  `json:"to_location_id"`   // 目标库位
	PutawayQty    float64 `json:"putaway_qty" gorm:"type:decimal(18,3)"`
	PutawardedQty float64 `json:"putawayed_qty" gorm:"type:decimal(18,3);default:0"`
	Status        string  `json:"status" gorm:"size:20;default:'PENDING'"`
}

func (WMSPutawayRecord) TableName() string {
	return "wms_putaway_record"
}

// WMSPutawayJobCreateReqVO 创建上架作业单请求
type WMSPutawayJobCreateReqVO struct {
	SourceType  string `json:"sourceType"`
	SourceNo    string `json:"sourceNo"`
	WarehouseID int64  `json:"warehouseId"`
}

// WMSPutawayJobAssignReqVO 分配操作员请求
type WMSPutawayJobAssignReqVO struct {
	ID           int64  `json:"id"`
	OperatorID   int64  `json:"operatorId"`
	OperatorName string `json:"operatorName"`
}

// WMSPutawayJobRespVO 上架作业单响应
type WMSPutawayJobRespVO struct {
	ID           int64                      `json:"id"`
	PutawayNo    string                     `json:"putawayNo"`
	SourceType   string                     `json:"sourceType"`
	SourceNo     string                     `json:"sourceNo"`
	Status       string                     `json:"status"`
	OperatorID   *int64                     `json:"operatorId"`
	OperatorName string                     `json:"operatorName"`
	Records      []WMSPutawayRecordRespVO   `json:"records"`
}

// WMSPutawayRecordRespVO 上架明细响应
type WMSPutawayRecordRespVO struct {
	ID             int64   `json:"id"`
	ItemID         int64   `json:"itemId"`
	ItemCode       string  `json:"itemCode"`
	ItemName       string  `json:"itemName"`
	FromLocationID *int64  `json:"fromLocationId"`
	ToLocationID   *int64  `json:"toLocationId"`
	PutawayQty     float64 `json:"putawayQty"`
	PutawardedQty  float64 `json:"putawayedQty"`
	Status         string  `json:"status"`
}

// WMSPutawayCancelReqVO 取消上架请求
type WMSPutawayCancelReqVO struct {
	ID     int64  `json:"id"`
	Reason string `json:"reason"`
}
