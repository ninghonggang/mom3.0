package repository

import (
	"context"

	"mom-platform/services/eam-service/internal/model"
)

// ============ Equipment ============

// EquipmentFilter holds filter criteria for listing equipment.
type EquipmentFilter struct {
	Type       model.EquipmentType
	Status     model.EquipmentStatus
	WorkshopID *int64
	LineID     *int64
}

// EquipmentRepository defines CRUD operations for Equipment.
type EquipmentRepository interface {
	Create(ctx context.Context, e *model.Equipment) error
	GetByID(ctx context.Context, id int64) (*model.Equipment, error)
	GetByCode(ctx context.Context, code string) (*model.Equipment, error)
	Update(ctx context.Context, e *model.Equipment) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, filter EquipmentFilter, page model.PageQuery) ([]model.Equipment, int64, error)
}

// ============ RepairOrder ============

// RepairOrderFilter holds filter criteria for listing repair orders.
type RepairOrderFilter struct {
	EquipmentID *int64
	Status      model.RepairOrderStatus
}

// RepairOrderRepository defines CRUD operations for RepairOrder.
type RepairOrderRepository interface {
	Create(ctx context.Context, r *model.RepairOrder) error
	GetByID(ctx context.Context, id int64) (*model.RepairOrder, error)
	Update(ctx context.Context, r *model.RepairOrder) error
	List(ctx context.Context, filter RepairOrderFilter, page model.PageQuery) ([]model.RepairOrder, int64, error)
	CountByStatus(ctx context.Context, status model.RepairOrderStatus) (int64, error)
}

// ============ EquipmentOee ============

// OeeFilter holds filter criteria for listing OEE records.
type OeeFilter struct {
	EquipmentID int64
	BeginDate   string
	EndDate     string
}

// OeeRepository defines CRUD operations for EquipmentOee.
type OeeRepository interface {
	Create(ctx context.Context, o *model.EquipmentOee) error
	GetByEquipmentAndDate(ctx context.Context, equipmentID int64, date string) (*model.EquipmentOee, error)
	List(ctx context.Context, filter OeeFilter) ([]model.EquipmentOee, error)
}

// ============ MaintenancePlan ============

// MaintenancePlanFilter holds filter criteria for listing maintenance plans.
type MaintenancePlanFilter struct {
	EquipmentID *int64
}

// MaintenancePlanRepository defines CRUD operations for MaintenancePlan.
type MaintenancePlanRepository interface {
	Create(ctx context.Context, p *model.MaintenancePlan) error
	GetByID(ctx context.Context, id int64) (*model.MaintenancePlan, error)
	Update(ctx context.Context, p *model.MaintenancePlan) error
	List(ctx context.Context, filter MaintenancePlanFilter, page model.PageQuery) ([]model.MaintenancePlan, int64, error)
}

// ============ EquipmentCheck ============

// CheckFilter holds filter criteria for listing check records.
type CheckFilter struct {
	EquipmentID *int64
	BeginTime   *int64 // unix timestamp
	EndTime     *int64 // unix timestamp
}

// CheckRepository defines CRUD operations for EquipmentCheck.
type CheckRepository interface {
	Create(ctx context.Context, c *model.EquipmentCheck) error
	List(ctx context.Context, filter CheckFilter, page model.PageQuery) ([]model.EquipmentCheck, int64, error)
}

// ============ EquipmentDowntime ============

// DowntimeFilter holds filter criteria for listing downtime records.
type DowntimeFilter struct {
	EquipmentID *int64
	Status      model.DowntimeStatus
}

// DowntimeRepository defines CRUD operations for EquipmentDowntime.
type DowntimeRepository interface {
	Create(ctx context.Context, d *model.EquipmentDowntime) error
	GetByID(ctx context.Context, id int64) (*model.EquipmentDowntime, error)
	Update(ctx context.Context, d *model.EquipmentDowntime) error
	List(ctx context.Context, filter DowntimeFilter, page model.PageQuery) ([]model.EquipmentDowntime, int64, error)
}

// Repository aggregates all individual repositories.
type Repository interface {
	Equipment() EquipmentRepository
	RepairOrder() RepairOrderRepository
	Oee() OeeRepository
	MaintenancePlan() MaintenancePlanRepository
	Check() CheckRepository
	Downtime() DowntimeRepository
}
