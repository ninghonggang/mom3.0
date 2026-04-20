package model

import "time"

// ScpQadSyncLog QAD同步记录
type ScpQadSyncLog struct {
	ID              uint64    `json:"id" gorm:"primaryKey"`
	TenantID        int64     `json:"tenant_id" gorm:"index"`
	SyncType        string    `json:"sync_type" gorm:"size:20"`    // ORDER/DELIVERY/INVENTORY
	SyncDirection   string    `json:"sync_direction" gorm:"size:10"` // UPLOAD/DOWNLOAD
	QadDocNo        string    `json:"qad_doc_no" gorm:"size:50;index"`
	MomDocNo        string    `json:"mom_doc_no" gorm:"size:50;index"`
	Status          string    `json:"status" gorm:"size:20"`       // PENDING/SUCCESS/FAILED
	RequestContent  string    `json:"request_content" gorm:"type:text"`
	ResponseContent string    `json:"response_content" gorm:"type:text"`
	ErrorMsg        string    `json:"error_msg" gorm:"size:500"`
	SyncTime        time.Time `json:"sync_time"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (ScpQadSyncLog) TableName() string {
	return "scp_qad_sync_log"
}

// QadSyncType QAD同步类型常量
var QadSyncType = struct {
	Order     string
	Delivery  string
	Inventory string
}{
	Order:     "ORDER",
	Delivery:  "DELIVERY",
	Inventory: "INVENTORY",
}

// QadSyncDirection QAD同步方向常量
var QadSyncDirection = struct {
	Upload   string
	Download string
}{
	Upload:   "UPLOAD",
	Download: "DOWNLOAD",
}

// QadSyncStatus QAD同步状态常量
var QadSyncStatus = struct {
	Pending string
	Success string
	Failed  string
}{
	Pending: "PENDING",
	Success: "SUCCESS",
	Failed:  "FAILED",
}

// QadSyncRequest 同步数据到QAD请求
type QadSyncRequest struct {
	SyncType string                 `json:"syncType" binding:"required"`
	DocNo    string                 `json:"docNo" binding:"required"`
	Data     map[string]interface{} `json:"data"`
}

// QadSyncResponse 同步响应
type QadSyncResponse struct {
	SyncID  uint64 `json:"syncId"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// QadConfirmRequest QAD订单确认回调请求
type QadConfirmRequest struct {
	QadDocNo    string `json:"qadDocNo" binding:"required"`
	MomDocNo    string `json:"momDocNo" binding:"required"`
	Status      string `json:"status" binding:"required"` // CONFIRMED/REJECTED
	ConfirmTime string `json:"confirmTime"`
	Remark      string `json:"remark"`
}

// QadDeliveryRequest QAD发货通知回调请求
type QadDeliveryRequest struct {
	QadDocNo     string  `json:"qadDocNo" binding:"required"`
	MomDocNo     string  `json:"momDocNo" binding:"required"`
	DeliveryDate string  `json:"deliveryDate"`
	MaterialCode string  `json:"materialCode"`
	Qty          float64 `json:"qty"`
	BatchNo      string  `json:"batchNo"`
	Remark       string  `json:"remark"`
}

// QadSyncLogQuery 同步日志查询
type QadSyncLogQuery struct {
	TenantID int64  `json:"tenantId"`
	SyncType string `json:"syncType"`
	Status   string `json:"status"`
	DocNo    string `json:"docNo"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}
