package model

import (
	"encoding/json"
	"time"
)

// ContainerStatus 器具状态
type ContainerStatus string

const (
	ContainerStatusInStock    ContainerStatus = "IN_STOCK"    // 在库
	ContainerStatusInUse       ContainerStatus = "IN_USE"      // 使用中
	ContainerStatusSentOut     ContainerStatus = "SENT_OUT"    // 已发出
	ContainerStatusMaintenance ContainerStatus = "MAINTENANCE" // 维修中
	ContainerStatusScrapped    ContainerStatus = "SCRAPPED"   // 已报废
)

// ContainerType 器具类型
type ContainerType string

const (
	ContainerTypeRack   ContainerType = "RACK"   // 货架
	ContainerTypeTray   ContainerType = "TRAY"   // 托盘
	ContainerTypeBox    ContainerType = "BOX"    // 箱子
	ContainerTypePallet ContainerType = "PALLET" // 托盘
)

// MovementType 流转类型
type MovementType string

const (
	MovementTypeIn       MovementType = "IN"       // 入库
	MovementTypeOut      MovementType = "OUT"      // 出库
	MovementTypeReturn   MovementType = "RETURN"   // 退回
	MovementTypeTransfer MovementType = "TRANSFER"  // 调拨
	MovementTypeClean    MovementType = "CLEAN"    // 清洁
)

// ContainerMaster 器具主表
type ContainerMaster struct {
	ID                 uint             `json:"id" gorm:"primaryKey"`
	TenantID           int64            `json:"tenant_id" gorm:"index;not null"`
	ContainerCode      string           `json:"container_code" gorm:"size:50;uniqueIndex;not null"` // 器具编号
	ContainerName      string           `json:"container_name" gorm:"size:100;not null"`            // 器具名称
	ContainerType      ContainerType    `json:"container_type" gorm:"size:30;not null"`            // RACK/TRAY/BOX/PALLET
	StandardQty        int              `json:"standard_qty" gorm:"default:0"`                     // 标准承载数量
	ApplicableProducts json.RawMessage  `json:"applicable_products" gorm:"type:jsonb"`              // 适用产品JSONB
	Status             ContainerStatus  `json:"status" gorm:"size:20;default:'IN_STOCK'"`          // IN_STOCK/IN_USE/SENT_OUT/MAINTENANCE/SCRAPPED
	LocationType       string           `json:"location_type" gorm:"size:20"`                      // INCOMING/FACTORY/OUTGOING
	CurrentLocation    string           `json:"current_location" gorm:"size:100"`                  // 当前位置
	CustomerID         int64            `json:"customer_id" gorm:"index"`                          // 客户ID(客户器具)
	Barcode            string           `json:"barcode" gorm:"size:100;uniqueIndex"`               // 条码
	TotalTrips         int              `json:"total_trips" gorm:"default:0"`                      // 总使用次数
	LastCleanDate      *time.Time       `json:"last_clean_date"`                                    // 上次清洁日期
	CreatedAt          time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
}

func (ContainerMaster) TableName() string {
	return "container_master"
}

// ContainerMovement 器具流转记录
type ContainerMovement struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	TenantID       int64        `json:"tenant_id" gorm:"index;not null"`
	ContainerID    int64        `json:"container_id" gorm:"index;not null"`    // 器具ID
	ContainerCode  string       `json:"container_code" gorm:"size:50"`         // 器具编号
	MovementType   MovementType `json:"movement_type" gorm:"size:20;not null"` // IN/OUT/RETURN/TRANSFER/CLEAN
	FromLocation   string       `json:"from_location" gorm:"size:100"`         // 来源位置
	ToLocation     string       `json:"to_location" gorm:"size:100"`           // 目标位置
	Qty            int          `json:"qty" gorm:"default:1"`                  // 数量
	RelatedOrderNo string       `json:"related_order_no" gorm:"size:50"`       // 关联单号
	OperatorID    int64        `json:"operator_id" gorm:"index"`              // 操作人ID
	OperatorName   string       `json:"operator_name" gorm:"size:50"`          // 操作人姓名
	MovementTime   time.Time    `json:"movement_time" gorm:"not null"`         // 流转时间
	Remark         string       `json:"remark" gorm:"size:200"`                // 备注
	CreatedAt      time.Time    `json:"created_at" gorm:"autoCreateTime"`
}

func (ContainerMovement) TableName() string {
	return "container_movement"
}

// ContainerQueryParams 器具查询参数
type ContainerQueryParams struct {
	Page         int              `form:"page" json:"page"`
	PageSize     int              `form:"pageSize" json:"pageSize"`
	Keyword      string           `form:"keyword" json:"keyword"`
	Status       ContainerStatus  `form:"status" json:"status"`
	ContainerType ContainerType   `form:"container_type" json:"container_type"`
	CustomerID   int64            `form:"customer_id" json:"customer_id"`
}

// ContainerMovementDTO 流转记录响应
type ContainerMovementDTO struct {
	ID             uint         `json:"id"`
	ContainerID    int64        `json:"container_id"`
	ContainerCode  string       `json:"container_code"`
	ContainerName  string       `json:"container_name"`
	MovementType   MovementType `json:"movement_type"`
	FromLocation   string       `json:"from_location"`
	ToLocation     string       `json:"to_location"`
	Qty            int          `json:"qty"`
	RelatedOrderNo string       `json:"related_order_no"`
	OperatorID     int64        `json:"operator_id"`
	OperatorName   string       `json:"operator_name"`
	MovementTime   time.Time    `json:"movement_time"`
	Remark         string       `json:"remark"`
	CreatedAt      time.Time    `json:"created_at"`
}
