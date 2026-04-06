package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type InspectItem struct {
	ItemName    string  `json:"item_name"`
	StdValue    string  `json:"std_value"`
	ActualValue string  `json:"actual_value"`
	Result      string  `json:"result"` // OK/NG
}

type InspectItems []InspectItem

func (i InspectItems) Value() (driver.Value, error) {
	return json.Marshal(i)
}

func (i *InspectItems) Scan(value interface{}) error {
	if value == nil {
		*i = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, i)
}

// MesFirstLastInspect 首末件检验单
type MesFirstLastInspect struct {
	ID                 uint         `json:"id" gorm:"primaryKey"`
	TenantID           int64        `json:"tenant_id" gorm:"index;not null"`
	WorkshopID         *int64       `json:"workshop_id" gorm:"index"`
	InspectNo          string       `json:"inspect_no" gorm:"size:50;uniqueIndex;not null"`
	InspectType        string       `json:"inspect_type" gorm:"size:10;not null"` // FIRST/LAST
	ProductionOrderID  int64        `json:"production_order_id" gorm:"index;not null"`
	ProcessID          int64        `json:"process_id" gorm:"not null"`
	WorkstationID      int64        `json:"workstation_id" gorm:"not null"`
	ShiftID            int64        `json:"shift_id" gorm:"not null"`
	ProductID          int64        `json:"product_id" gorm:"not null"`
	SerialNo           string       `json:"serial_no" gorm:"size:100"`
	InspectItems       InspectItems `json:"inspect_items" gorm:"type:jsonb"`
	OverallResult      string       `json:"overall_result" gorm:"size:10"` // OK/NG/PENDING
	InspectorID        *int64       `json:"inspector_id" gorm:"index"`
	InspectorName      string       `json:"inspector_name" gorm:"size:50"`
	InspectTime        *time.Time   `json:"inspect_time"`
	BluetoothDeviceID  string       `json:"bluetooth_device_id" gorm:"size:100"`
	Remark             string       `json:"remark" gorm:"size:500"`
	Status             string       `json:"status" gorm:"size:20;default:'PENDING'"` // PENDING/COMPLETED
	CreatedAt          time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}

func (MesFirstLastInspect) TableName() string {
	return "mes_first_last_inspect"
}