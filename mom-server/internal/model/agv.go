package model

import (
	"time"
)

// AGVTaskType AGV任务类型
type AGVTaskType string

const (
	AGVTaskTypeDelivery AGVTaskType = "DELIVERY" // 配送任务
	AGVTaskTypePickup   AGVTaskType = "PICKUP"   // 取货任务
	AGVTaskTypeTransfer AGVTaskType = "TRANSFER"  // 转移任务
)

// AGVTaskStatus AGV任务状态
type AGVTaskStatus string

const (
	AGVTaskStatusPending    AGVTaskStatus = "PENDING"    // 待分配
	AGVTaskStatusAssigned   AGVTaskStatus = "ASSIGNED"  // 已分配
	AGVTaskStatusInProgress AGVTaskStatus = "IN_PROGRESS" // 执行中
	AGVTaskStatusCompleted  AGVTaskStatus = "COMPLETED"  // 已完成
	AGVTaskStatusCancelled   AGVTaskStatus = "CANCELLED"  // 已取消
	AGVTaskStatusException   AGVTaskStatus = "EXCEPTION"  // 异常
)

// AGVDeviceStatus AGV设备状态
type AGVDeviceStatus string

const (
	AGVDeviceStatusOnline   AGVDeviceStatus = "ONLINE"   // 在线
	AGVDeviceStatusOffline  AGVDeviceStatus = "OFFLINE"  // 离线
	AGVDeviceStatusCharging  AGVDeviceStatus = "CHARGING" // 充电中
	AGVDeviceStatusFault     AGVDeviceStatus = "FAULT"   // 故障
	AGVDeviceStatusBusy      AGVDeviceStatus = "BUSY"     // 忙碌
)

// AGVLocationType 库位类型
type AGVLocationType string

const (
	AGVLocationTypeStation  AGVLocationType = "STATION"  // 工位
	AGVLocationTypeCharge   AGVLocationType = "CHARGE"  // 充电位
	AGVLocationTypeParking AGVLocationType = "PARKING" // 停车位
	AGVLocationTypeWarehouse AGVLocationType = "WAREHOUSE" // 仓库
)

// AGVTask AGV任务
type AGVTask struct {
	ID               int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID         int64          `gorm:"column:tenant_id;default:1" json:"tenantId"`
	TaskNo           string         `gorm:"column:task_no;type:varchar(50);uniqueIndex;not null" json:"taskNo"` // 任务编号
	TaskType         AGVTaskType    `gorm:"column:task_type;type:varchar(20);not null" json:"taskType"`        // 任务类型
	Status           AGVTaskStatus  `gorm:"column:status;type:varchar(20);default:PENDING" json:"status"`      // 任务状态
	Priority         int            `gorm:"column:priority;default:0" json:"priority"`                          // 优先级
	SourceLocationID int64          `gorm:"column:source_location_id" json:"sourceLocationId"`                 // 起始位置ID
	SourceLocation   string         `gorm:"column:source_location;type:varchar(100)" json:"sourceLocation"`     // 起始位置名称
	TargetLocationID int64          `gorm:"column:target_location_id" json:"targetLocationId"`                 // 目标位置ID
	TargetLocation   string         `gorm:"column:target_location;type:varchar(100)" json:"targetLocation"`   // 目标位置名称
	MaterialID       int64          `gorm:"column:material_id" json:"materialId"`                               // 物料ID
	MaterialCode     string         `gorm:"column:material_code;type:varchar(50)" json:"materialCode"`         // 物料编码
	MaterialName     string         `gorm:"column:material_name;type:varchar(100)" json:"materialName"`        // 物料名称
	Quantity         float64        `gorm:"column:quantity;type:decimal(18,6)" json:"quantity"`                // 数量
	Unit             string         `gorm:"column:unit;type:varchar(20)" json:"unit"`                         // 单位
	AssignedAGVCode  string         `gorm:"column:assigned_agv_code;type:varchar(50)" json:"assignedAgvCode"` // 分配的AGV编号
	AssignedAGVName  string         `gorm:"column:assigned_agv_name;type:varchar(100)" json:"assignedAgvName"` // 分配的AGV名称
	StartedAt        *time.Time     `gorm:"column:started_at" json:"startedAt"`                                // 开始时间
	CompletedAt      *time.Time     `gorm:"column:completed_at" json:"completedAt"`                             // 完成时间
	ErrorMessage     string         `gorm:"column:error_message;type:text" json:"errorMessage"`               // 错误信息
	RelatedOrderNo   string         `gorm:"column:related_order_no;type:varchar(50)" json:"relatedOrderNo"`   // 关联单据编号
	RelatedOrderType string         `gorm:"column:related_order_type;type:varchar(50)" json:"relatedOrderType"` // 关联单据类型
	ExtData          string         `gorm:"column:ext_data;type:jsonb" json:"extData"`                        // 扩展数据
	CreatedBy        string         `gorm:"column:created_by;type:varchar(50)" json:"createdBy"`              // 创建人
	CreatedAt        time.Time      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (AGVTask) TableName() string {
	return "wms_agv_task"
}

// AGVDevice AGV设备
type AGVDevice struct {
	ID              int64           `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID        int64           `gorm:"column:tenant_id;default:1" json:"tenantId"`
	AGVCode         string          `gorm:"column:agv_code;type:varchar(50);uniqueIndex;not null" json:"agvCode"` // AGV编号
	AGVName         string          `gorm:"column:agv_name;type:varchar(100)" json:"agvName"`                   // AGV名称
	AGVType         string          `gorm:"column:agv_type;type:varchar(30)" json:"agvType"`                     // AGV类型
	Status          AGVDeviceStatus `gorm:"column:status;type:varchar(20);default:OFFLINE" json:"status"`         // 设备状态
	CurrentLocation string          `gorm:"column:current_location;type:varchar(100)" json:"currentLocation"`     // 当前位置
	BatteryLevel    float64         `gorm:"column:battery_level;type:decimal(5,2)" json:"batteryLevel"`           // 电池电量
	MaxLoad         float64         `gorm:"column:max_load;type:decimal(18,6)" json:"maxLoad"`                  // 最大载重
	ExtConfig       string          `gorm:"column:ext_config;type:jsonb" json:"extConfig"`                       // 扩展配置
	LastHeartbeat   *time.Time      `gorm:"column:last_heartbeat" json:"lastHeartbeat"`                         // 最后心跳时间
	Remark          string          `gorm:"column:remark;type:text" json:"remark"`                               // 备注
	CreatedAt       time.Time       `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt       time.Time       `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (AGVDevice) TableName() string {
	return "wms_agv_device"
}

// AGVLocationMapping AGV库位映射
type AGVLocationMapping struct {
	ID               int64           `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID         int64           `gorm:"column:tenant_id;default:1" json:"tenantId"`
	LocationCode     string          `gorm:"column:location_code;type:varchar(50);not null" json:"locationCode"` // 库位编码
	LocationName     string          `gorm:"column:location_name;type:varchar(100)" json:"locationName"`          // 库位名称
	LocationType     AGVLocationType `gorm:"column:location_type;type:varchar(30)" json:"locationType"`          // 库位类型
	AGVLocationCode  string          `gorm:"column:agv_location_code;type:varchar(50)" json:"agvLocationCode"`   // AGV定位编码
	XCoord           float64         `gorm:"column:x_coord;type:decimal(10,2)" json:"xCoord"`                    // X坐标
	YCoord           float64         `gorm:"column:y_coord;type:decimal(10,2)" json:"yCoord"`                    // Y坐标
	Priority         int             `gorm:"column:priority;default:0" json:"priority"`                          // 优先级
	Enabled          bool            `gorm:"column:enabled;default:true" json:"enabled"`                         // 是否启用
	ExtData          string          `gorm:"column:ext_data;type:jsonb" json:"extData"`                         // 扩展数据
	CreatedAt        time.Time       `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt        time.Time       `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (AGVLocationMapping) TableName() string {
	return "wms_agv_location_mapping"
}

// ============ DTOs ============

// AGVTaskQuery AGV任务查询
type AGVTaskQuery struct {
	TenantID    int64          `form:"tenantId" json:"tenantId"`
	TaskNo      string         `form:"taskNo" json:"taskNo"`
	TaskType    AGVTaskType    `form:"taskType" json:"taskType"`
	Status      AGVTaskStatus   `form:"status" json:"status"`
	Priority    int            `form:"priority" json:"priority"`
	AGVCode     string         `form:"agvCode" json:"agvCode"`
	StartDate   string         `form:"startDate" json:"startDate"`
	EndDate     string         `form:"endDate" json:"endDate"`
	Page        int            `form:"page" json:"page"`
	PageSize    int            `form:"pageSize" json:"pageSize"`
}

// AGVDeviceQuery AGV设备查询
type AGVDeviceQuery struct {
	TenantID int64           `form:"tenantId" json:"tenantId"`
	AGVCode  string         `form:"agvCode" json:"agvCode"`
	Status   AGVDeviceStatus `form:"status" json:"status"`
	Page     int            `form:"page" json:"page"`
	PageSize int            `form:"pageSize" json:"pageSize"`
}

// AGVLocationQuery AGV库位查询
type AGVLocationQuery struct {
	TenantID    int64           `form:"tenantId" json:"tenantId"`
	LocationCode string         `form:"locationCode" json:"locationCode"`
	LocationType AGVLocationType `form:"locationType" json:"locationType"`
	Enabled     *bool           `form:"enabled" json:"enabled"`
	Page        int            `form:"page" json:"page"`
	PageSize    int            `form:"pageSize" json:"pageSize"`
}

// CreateAGVTaskRequest 创建AGV任务请求
type CreateAGVTaskRequest struct {
	TenantID         int64   `json:"tenantId" binding:"required"`
	TaskType         string  `json:"taskType" binding:"required"` // DELIVERY/PICKUP/TRANSFER
	Priority         int     `json:"priority"`
	SourceLocationID int64   `json:"sourceLocationId"`
	SourceLocation   string  `json:"sourceLocation"`
	TargetLocationID int64   `json:"targetLocationId"`
	TargetLocation   string  `json:"targetLocation"`
	MaterialID       int64   `json:"materialId"`
	MaterialCode     string  `json:"materialCode"`
	MaterialName     string  `json:"materialName"`
	Quantity         float64 `json:"quantity"`
	Unit             string  `json:"unit"`
	RelatedOrderNo   string  `json:"relatedOrderNo"`
	RelatedOrderType string  `json:"relatedOrderType"`
}

// UpdateAGVTaskStatusRequest 更新AGV任务状态请求
type UpdateAGVTaskStatusRequest struct {
	Status     string `json:"status"`
	ErrorMsg   string `json:"errorMsg"`
	AGVCode    string `json:"agvCode"`
	AGVName    string `json:"agvName"`
}

// CreateAGVDeviceRequest 创建AGV设备请求
type CreateAGVDeviceRequest struct {
	TenantID   int64   `json:"tenantId" binding:"required"`
	AGVCode    string  `json:"agvCode" binding:"required"`
	AGVName    string  `json:"agvName"`
	AGVType    string  `json:"agvType"`
	MaxLoad    float64 `json:"maxLoad"`
	Remark     string  `json:"remark"`
}

// UpdateAGVDeviceStatusRequest 更新AGV设备状态请求
type UpdateAGVDeviceStatusRequest struct {
	Status          string  `json:"status"`
	CurrentLocation string  `json:"currentLocation"`
	BatteryLevel   float64 `json:"batteryLevel"`
}

// CreateAGVLocationRequest 创建库位映射请求
type CreateAGVLocationRequest struct {
	TenantID     int64   `json:"tenantId" binding:"required"`
	LocationCode string  `json:"locationCode" binding:"required"`
	LocationName string  `json:"locationName"`
	LocationType string  `json:"locationType"`
	AGVLocationCode string `json:"agvLocationCode"`
	XCoord       float64 `json:"xCoord"`
	YCoord       float64 `json:"yCoord"`
	Priority     int     `json:"priority"`
	Enabled      bool    `json:"enabled"`
}
