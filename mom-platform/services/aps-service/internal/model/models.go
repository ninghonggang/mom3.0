package model

import "time"

type MpsPlan struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID    string    `gorm:"column:tenant_id;type:varchar(64);index;not null" json:"tenant_id"`
	PlanNo      string    `gorm:"column:plan_no;type:varchar(64);uniqueIndex;not null" json:"plan_no"`
	PlanMonth   string    `gorm:"column:plan_month;type:varchar(7);not null" json:"plan_month"` // YYYY-MM
	MaterialID  uint64    `gorm:"column:material_id;not null" json:"material_id"`
	PlannedQty  float64   `gorm:"column:planned_qty;type:decimal(18,4);not null" json:"planned_qty"`
	Status      string    `gorm:"column:status;type:varchar(64);default:DRAFT" json:"status"` // DRAFT/CONFIRMED/RELEASED/CLOSED
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (MpsPlan) TableName() string { return "mps_plan" }

type MrpPlan struct {
	ID                  uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID            string    `gorm:"column:tenant_id;type:varchar(64);index;not null" json:"tenant_id"`
	PlanNo              string    `gorm:"column:plan_no;type:varchar(64);uniqueIndex;not null" json:"plan_no"`
	MpsID               uint64    `gorm:"column:mps_id;index" json:"mps_id"`
	MaterialID          uint64    `gorm:"column:material_id;not null" json:"material_id"`
	GrossReq            float64   `gorm:"column:gross_req;type:decimal(18,4)" json:"gross_req"`
	ScheduledReceipt    float64   `gorm:"column:scheduled_receipt;type:decimal(18,4)" json:"scheduled_receipt"`
	ProjectedOnHand     float64   `gorm:"column:projected_on_hand;type:decimal(18,4)" json:"projected_on_hand"`
	NetReq              float64   `gorm:"column:net_req;type:decimal(18,4)" json:"net_req"`
	PlannedOrderRelease float64   `gorm:"column:planned_order_release;type:decimal(18,4)" json:"planned_order_release"`
	PlannedOrderReceipt float64   `gorm:"column:planned_order_receipt;type:decimal(18,4)" json:"planned_order_receipt"`
	CreatedAt           time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (MrpPlan) TableName() string { return "mrp_plan" }

type WorkCenter struct {
	ID             uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID       string    `gorm:"column:tenant_id;type:varchar(64);index;not null" json:"tenant_id"`
	CenterCode     string    `gorm:"column:center_code;type:varchar(64);uniqueIndex;not null" json:"center_code"`
	CenterName     string    `gorm:"column:center_name;type:varchar(128);not null" json:"center_name"`
	WorkshopID     uint64    `gorm:"column:workshop_id;not null" json:"workshop_id"`
	CapacityPerDay float64   `gorm:"column:capacity_per_day;type:decimal(12,2)" json:"capacity_per_day"`
	Efficiency     float64   `gorm:"column:efficiency;type:decimal(5,2);default:100" json:"efficiency"`
	Status         string    `gorm:"column:status;type:varchar(64);default:ACTIVE" json:"status"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (WorkCenter) TableName() string { return "work_center" }

type ScheduleJob struct {
	ID                uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID          string     `gorm:"column:tenant_id;type:varchar(64);index;not null" json:"tenant_id"`
	JobNo             string     `gorm:"column:job_no;type:varchar(64);uniqueIndex;not null" json:"job_no"`
	MpsID             uint64     `gorm:"column:mps_id;index" json:"mps_id"`
	PlanType          string     `gorm:"column:plan_type;type:varchar(32)" json:"plan_type"` // ROUGH/FINE
	Algorithm         string     `gorm:"column:algorithm;type:varchar(32)" json:"algorithm"` // FIFO/EDD/SPT/LPT
	ProductionOrderID uint64     `gorm:"column:production_order_id;index" json:"production_order_id"`
	WorkCenterID      uint64     `gorm:"column:work_center_id;not null" json:"work_center_id"`
	PlannedStart      time.Time  `gorm:"column:planned_start" json:"planned_start"`
	PlannedEnd        time.Time  `gorm:"column:planned_end" json:"planned_end"`
	ActualStart       *time.Time `gorm:"column:actual_start" json:"actual_start,omitempty"`
	ActualEnd         *time.Time `gorm:"column:actual_end" json:"actual_end,omitempty"`
	Status            string     `gorm:"column:status;type:varchar(64);default:PENDING" json:"status"` // PENDING/SCHEDULED/RUNNING/COMPLETED/DELAYED
	CreatedAt         time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (ScheduleJob) TableName() string { return "schedule_job" }

type ScheduleConstraint struct {
	ID             uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	JobID          uint64    `gorm:"column:job_id;index;not null" json:"job_id"`
	ConstraintType string    `gorm:"column:constraint_type;type:varchar(32);not null" json:"constraint_type"`
	Description    string    `gorm:"column:description;type:varchar(512)" json:"description"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (ScheduleConstraint) TableName() string { return "schedule_constraint" }

type Changeover struct {
	ID                   uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	FromMaterialID       uint64    `gorm:"column:from_material_id;index;not null" json:"from_material_id"`
	ToMaterialID         uint64    `gorm:"column:to_material_id;index;not null" json:"to_material_id"`
	ChangeoverTimeMinutes float64  `gorm:"column:changeover_time_minutes;type:decimal(10,2);not null" json:"changeover_time_minutes"`
	CreatedAt            time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Changeover) TableName() string { return "changeover" }
