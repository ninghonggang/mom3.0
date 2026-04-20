package model

import "time"

// ========== 盘点管理扩展 ==========

// StockCheck 盘点单
type StockCheck struct {
	BaseModel
	TenantID            int64      `json:"tenant_id" gorm:"index;not null"`
	CheckNo             string     `json:"check_no" gorm:"size:50;uniqueIndex:idx_tenant_check"`
	CheckType           string     `json:"check_type" gorm:"size:20"` // FULL/CYCLE/BLIND/ADJUSTMENT
	WarehouseID         int64      `json:"warehouse_id"`
	WarehouseName       string     `json:"warehouse_name" gorm:"size:100"`
	AreaIDs             string     `json:"area_ids" gorm:"type:jsonb"`      // 盘点库区列表
	LocationIDs         string     `json:"location_ids" gorm:"type:jsonb"`  // 盘点库位列表
	Status              string     `json:"status" gorm:"size:20;default:DRAFT"` // DRAFT/IN_PROGRESS/COUNTING/COMPLETED/CANCELLED
	PlanStartDate       string     `json:"plan_start_date" gorm:"size:10"`
	PlanEndDate         string     `json:"plan_end_date" gorm:"size:10"`
	ActualStartDate     *string    `json:"actual_start_date" gorm:"size:10"`
	ActualEndDate       *string    `json:"actual_end_date" gorm:"size:10"`
	IncludeZeroStock    int        `json:"include_zero_stock" gorm:"default:1"`    // 是否包含零库存库位
	IncludeExpiredStock int        `json:"include_expired_stock" gorm:"default:0"` // 是否包含过期物料
	IsBlindMode        int        `json:"is_blind_mode" gorm:"default:0"`       // 是否盲盘
	CheckerID          int64      `json:"checker_id"`
	CheckerName        string     `json:"checker_name" gorm:"size:50"`
	AuditID            *int64     `json:"audit_id"`
	AuditName          *string    `json:"audit_name" gorm:"size:50"`
	TotalLocations     int        `json:"total_locations" gorm:"default:0"`   // 盘点库位数
	CountedLocations   int        `json:"counted_locations" gorm:"default:0"`  // 已盘点库位数
	TotalMaterials     int        `json:"total_materials" gorm:"default:0"`    // 盘点物料种数
	VarianceCount      int        `json:"variance_count" gorm:"default:0"`     // 差异条数
	VarianceRate       float64    `json:"variance_rate" gorm:"type:decimal(5,2)"` // 差异率
	ApprovalStatus     string     `json:"approval_status" gorm:"size:20;default:PENDING"` // PENDING/APPROVED/REJECTED
	ApprovedBy         *int64     `json:"approved_by"`
	ApprovedTime       *time.Time `json:"approved_time"`
	ApprovalComment    *string    `json:"approval_comment" gorm:"type:text"`
	Remark             *string    `json:"remark" gorm:"type:text"`
	WorkshopID         *int64     `json:"workshop_id"`
}

func (StockCheck) TableName() string {
	return "wms_stock_check"
}

// StockCheckItem 盘点明细
type StockCheckItem struct {
	BaseModel
	TenantID        int64      `json:"tenant_id" gorm:"index;not null"`
	CheckID         int64      `json:"check_id"`
	LineNo          int        `json:"line_no"`
	MaterialID      int64      `json:"material_id"`
	MaterialCode    string     `json:"material_code" gorm:"size:50"`
	MaterialName    string     `json:"material_name" gorm:"size:100"`
	Specification   *string    `json:"specification" gorm:"size:200"`
	Unit            *string    `json:"unit" gorm:"size:20"`
	LocationID      int64      `json:"location_id"`
	LocationName    string     `json:"location_name" gorm:"size:100"`
	AreaName        *string    `json:"area_name" gorm:"size:100"`
	BatchNo         *string    `json:"batch_no" gorm:"size:50"`
	SystemQty       float64    `json:"system_qty" gorm:"type:decimal(18,3)"`   // 系统库存(盲盘时不显示)
	CountedQty      float64    `json:"counted_qty" gorm:"type:decimal(18,3)"`   // 盘点数量
	VarianceQty     float64    `json:"variance_qty" gorm:"type:decimal(18,3)"` // 差异数量
	VarianceAmount  float64    `json:"variance_amount" gorm:"type:decimal(18,2)"` // 差异金额
	VarianceReason  *string    `json:"variance_reason" gorm:"size:200"`          // 差异原因
	HandleStatus    string     `json:"handle_status" gorm:"size:20;default:PENDING"` // PENDING/APPROVED/PROCESSED
	HandleMethod    *string    `json:"handle_method" gorm:"size:20"`             // ADJUST/WRITE_OFF/WRITE_IN
	HandleQty       float64    `json:"handle_qty" gorm:"type:decimal(18,3)"`
	HandleTime      *time.Time `json:"handle_time"`
	HandlerID       *int64     `json:"handler_id"`
	HandlerName     *string    `json:"handler_name" gorm:"size:50"`
	CountStatus     string     `json:"count_status" gorm:"size:20;default:PENDING"` // PENDING/COUNTED/CONFIRMED
	CountTime       *time.Time `json:"count_time"`
	CounterID       *int64     `json:"counter_id"`
	CounterName     *string    `json:"counter_name" gorm:"size:50"`
	ConfirmTime     *time.Time `json:"confirm_time"`
	ConfirmerID     *int64     `json:"confirmer_id"`
	ConfirmerName   *string    `json:"confirmer_name" gorm:"size:50"`
	RecountQty      float64    `json:"recount_qty" gorm:"type:decimal(18,3)"`
	RecountTime     *time.Time `json:"recount_time"`
	RecountBy       *int64     `json:"recount_by"`
	RecountByName   *string    `json:"recount_by_name" gorm:"size:50"`
	IsFinalConfirmed int       `json:"is_final_confirmed" gorm:"default:0"`
	Photos          string     `json:"photos" gorm:"type:jsonb"` // 盘点照片
	Remark          *string    `json:"remark" gorm:"size:500"`
}

func (StockCheckItem) TableName() string {
	return "wms_stock_check_item"
}

// StocktakeTask 盘点任务
type StocktakeTask struct {
	BaseModel
	StocktakeID    int64      `json:"stocktake_id" gorm:"index"`
	TaskNo         string     `json:"task_no" gorm:"size:50;uniqueIndex"`
	AssigneeID     int64      `json:"assignee_id"`
	AssigneeName   string     `json:"assignee_name" gorm:"size:50"`
	LocationIDs    string     `json:"location_ids" gorm:"type:jsonb"`  // 分配的库位列表
	LocationCount  int        `json:"location_count" gorm:"default:0"`  // 库位数量
	Status         string     `json:"status" gorm:"size:20;default:PENDING"` // PENDING/IN_PROGRESS/COMPLETED
	StartedAt      *time.Time `json:"started_at"`
	CompletedAt    *time.Time `json:"completed_at"`
	CountedCount   int        `json:"counted_count" gorm:"default:0"`  // 已盘点库位
	VarianceCount  int        `json:"variance_count" gorm:"default:0"` // 差异数量
	Remark         *string    `json:"remark" gorm:"size:500"`
	TenantID       int64      `json:"tenant_id" gorm:"index;not null"`
}

func (StocktakeTask) TableName() string {
	return "wms_stocktake_task"
}

// CycleCountConfig 循环盘点配置
type CycleCountConfig struct {
	BaseModel
	ConfigCode      string    `json:"config_code" gorm:"size:50;uniqueIndex"`
	ConfigName      string    `json:"config_name" gorm:"size:100"`
	WarehouseID     int64     `json:"warehouse_id"`
	WarehouseName   string    `json:"warehouse_name" gorm:"size:100"`
	CountClassAEvery int      `json:"count_class_a_every" gorm:"default:30"`  // A类物料盘点周期(天)
	CountClassBEvery int      `json:"count_class_b_every" gorm:"default:60"`  // B类物料盘点周期(天)
	CountClassCEvery int      `json:"count_class_c_every" gorm:"default:180"` // C类物料盘点周期(天)
	TriggerType     string    `json:"trigger_type" gorm:"size:20;default:SCHEDULE"` // SCHEDULE/AUTO/MANUAL
	TriggerCron     string    `json:"trigger_cron" gorm:"size:50"`
	PriorityRules   string    `json:"priority_rules" gorm:"type:jsonb"`  // 优先规则
	AutoAssign     int       `json:"auto_assign" gorm:"default:1"`     // 是否自动分配任务
	AssignRules    string    `json:"assign_rules" gorm:"type:jsonb"`   // 分配规则
	IsEnabled      int       `json:"is_enabled" gorm:"default:1"`
	Remark         string    `json:"remark" gorm:"size:500"`
	TenantID       int64     `json:"tenant_id" gorm:"index;not null"`
}

func (CycleCountConfig) TableName() string {
	return "wms_cycle_count_config"
}

// StocktakeRecord 盘点记录(历史存档)
type StocktakeRecord struct {
	BaseModel
	StocktakeNo    string    `json:"stocktake_no" gorm:"size:50"`
	StocktakeType string    `json:"stocktake_type" gorm:"size:20"`
	WarehouseID   *int64    `json:"warehouse_id"`
	WarehouseName string    `json:"warehouse_name" gorm:"size:100"`
	StocktakeDate string    `json:"stocktake_date" gorm:"size:10"`
	CheckerID     *int64    `json:"checker_id"`
	CheckerName   string    `json:"checker_name" gorm:"size:50"`
	TotalLocations int      `json:"total_locations"`
	TotalMaterials int      `json:"total_materials"`
	VarianceCount int       `json:"variance_count"`
	VarianceRate  float64   `json:"variance_rate" gorm:"type:decimal(5,2)"`
	TotalVarianceAmount float64 `json:"total_variance_amount" gorm:"type:decimal(18,2)"`
	Details       string    `json:"details" gorm:"type:jsonb"` // 盘点明细JSON
	TenantID      int64     `json:"tenant_id" gorm:"index;not null"`
}

func (StocktakeRecord) TableName() string {
	return "wms_stocktake_record"
}
