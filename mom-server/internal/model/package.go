package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type SerialNos []string

func (s SerialNos) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *SerialNos) Scan(value interface{}) error {
	if value == nil {
		*s = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, s)
}

// MesPackage 包装箱
type MesPackage struct {
	ID                uint        `json:"id" gorm:"primaryKey"`
	TenantID          int64       `json:"tenant_id" gorm:"index;not null"`
	WorkshopID        *int64      `json:"workshop_id" gorm:"index"`
	PackageNo         string      `json:"package_no" gorm:"size:100;uniqueIndex;not null"` // 箱条码
	PackageType       string      `json:"package_type" gorm:"size:20"`                    // SMALL_BOX/BIG_BOX/PALLET
	ProductionOrderID *int64      `json:"production_order_id" gorm:"index"`
	ProductID         int64       `json:"product_id" gorm:"not null"`
	ProductCode       string      `json:"product_code" gorm:"size:50"`
	Qty               int         `json:"qty" gorm:"not null"`
	SerialNos         SerialNos   `json:"serial_nos" gorm:"type:jsonb"` // 箱内序列号列表
	Status            string      `json:"status" gorm:"size:20;default:'OPEN'"` // OPEN/SEALED/SHIPPED
	SealTime          *time.Time  `json:"seal_time"`
	SealBy            string      `json:"seal_by" gorm:"size:50"`
	ShipTime          *time.Time  `json:"ship_time"`
	CustomerID        *int64      `json:"customer_id" gorm:"index"`
	ContainerID       *int64      `json:"container_id" gorm:"index"` // 器具ID
	CreatedAt         time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

func (MesPackage) TableName() string {
	return "mes_package"
}
