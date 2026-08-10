package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"mom-platform/services/eam-service/internal/model"
)

// gormRepository implements the Repository interface using GORM.
type gormRepository struct {
	db *gorm.DB
}

// NewRepository creates a new Repository backed by GORM.
func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

func (r *gormRepository) Equipment() EquipmentRepository       { return &equipmentRepo{db: r.db} }
func (r *gormRepository) RepairOrder() RepairOrderRepository   { return &repairOrderRepo{db: r.db} }
func (r *gormRepository) Oee() OeeRepository                   { return &oeeRepo{db: r.db} }
func (r *gormRepository) MaintenancePlan() MaintenancePlanRepository {
	return &maintenancePlanRepo{db: r.db}
}
func (r *gormRepository) Check() CheckRepository               { return &checkRepo{db: r.db} }
func (r *gormRepository) Downtime() DowntimeRepository         { return &downtimeRepo{db: r.db} }

// ============ Equipment ============

type equipmentRepo struct {
	db *gorm.DB
}

func (r *equipmentRepo) Create(ctx context.Context, e *model.Equipment) error {
	return r.db.WithContext(ctx).Create(e).Error
}

func (r *equipmentRepo) GetByID(ctx context.Context, id int64) (*model.Equipment, error) {
	var e model.Equipment
	if err := r.db.WithContext(ctx).First(&e, id).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *equipmentRepo) GetByCode(ctx context.Context, code string) (*model.Equipment, error) {
	var e model.Equipment
	if err := r.db.WithContext(ctx).Where("equipment_code = ?", code).First(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *equipmentRepo) Update(ctx context.Context, e *model.Equipment) error {
	return r.db.WithContext(ctx).Save(e).Error
}

func (r *equipmentRepo) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Equipment{}, id).Error
}

func (r *equipmentRepo) List(ctx context.Context, filter EquipmentFilter, page model.PageQuery) ([]model.Equipment, int64, error) {
	q := r.db.WithContext(ctx).Model(&model.Equipment{})
	if filter.Type != "" {
		q = q.Where("equipment_type = ?", filter.Type)
	}
	if filter.Status != "" {
		q = q.Where("status = ?", filter.Status)
	}
	if filter.WorkshopID != nil {
		q = q.Where("workshop_id = ?", *filter.WorkshopID)
	}
	if filter.LineID != nil {
		q = q.Where("line_id = ?", *filter.LineID)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var items []model.Equipment
	if err := q.Order("id DESC").Offset(page.Offset()).Limit(int(page.PageSize)).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// ============ RepairOrder ============

type repairOrderRepo struct {
	db *gorm.DB
}

func (r *repairOrderRepo) Create(ctx context.Context, ro *model.RepairOrder) error {
	return r.db.WithContext(ctx).Create(ro).Error
}

func (r *repairOrderRepo) GetByID(ctx context.Context, id int64) (*model.RepairOrder, error) {
	var ro model.RepairOrder
	if err := r.db.WithContext(ctx).First(&ro, id).Error; err != nil {
		return nil, err
	}
	return &ro, nil
}

func (r *repairOrderRepo) Update(ctx context.Context, ro *model.RepairOrder) error {
	return r.db.WithContext(ctx).Save(ro).Error
}

func (r *repairOrderRepo) List(ctx context.Context, filter RepairOrderFilter, page model.PageQuery) ([]model.RepairOrder, int64, error) {
	q := r.db.WithContext(ctx).Model(&model.RepairOrder{})
	if filter.EquipmentID != nil {
		q = q.Where("equipment_id = ?", *filter.EquipmentID)
	}
	if filter.Status != "" {
		q = q.Where("status = ?", filter.Status)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var items []model.RepairOrder
	if err := q.Order("id DESC").Offset(page.Offset()).Limit(int(page.PageSize)).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (r *repairOrderRepo) CountByStatus(ctx context.Context, status model.RepairOrderStatus) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&model.RepairOrder{}).Where("status = ?", status).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// ============ EquipmentOee ============

type oeeRepo struct {
	db *gorm.DB
}

func (r *oeeRepo) Create(ctx context.Context, o *model.EquipmentOee) error {
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "equipment_id"}, {Name: "calc_date"}},
			DoUpdates: clause.AssignmentColumns([]string{"availability", "performance", "quality", "oee", "updated_at"}),
		}).
		Create(o).Error
}

func (r *oeeRepo) GetByEquipmentAndDate(ctx context.Context, equipmentID int64, date string) (*model.EquipmentOee, error) {
	var o model.EquipmentOee
	if err := r.db.WithContext(ctx).
		Where("equipment_id = ? AND calc_date = ?", equipmentID, date).
		First(&o).Error; err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *oeeRepo) List(ctx context.Context, filter OeeFilter) ([]model.EquipmentOee, error) {
	q := r.db.WithContext(ctx).Model(&model.EquipmentOee{})
	if filter.EquipmentID > 0 {
		q = q.Where("equipment_id = ?", filter.EquipmentID)
	}
	if filter.BeginDate != "" {
		q = q.Where("calc_date >= ?", filter.BeginDate)
	}
	if filter.EndDate != "" {
		q = q.Where("calc_date <= ?", filter.EndDate)
	}
	var items []model.EquipmentOee
	if err := q.Order("calc_date DESC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

// ============ MaintenancePlan ============

type maintenancePlanRepo struct {
	db *gorm.DB
}

func (r *maintenancePlanRepo) Create(ctx context.Context, p *model.MaintenancePlan) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *maintenancePlanRepo) GetByID(ctx context.Context, id int64) (*model.MaintenancePlan, error) {
	var p model.MaintenancePlan
	if err := r.db.WithContext(ctx).First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *maintenancePlanRepo) Update(ctx context.Context, p *model.MaintenancePlan) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r *maintenancePlanRepo) List(ctx context.Context, filter MaintenancePlanFilter, page model.PageQuery) ([]model.MaintenancePlan, int64, error) {
	q := r.db.WithContext(ctx).Model(&model.MaintenancePlan{})
	if filter.EquipmentID != nil {
		q = q.Where("equipment_id = ?", *filter.EquipmentID)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var items []model.MaintenancePlan
	if err := q.Order("id DESC").Offset(page.Offset()).Limit(int(page.PageSize)).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// ============ EquipmentCheck ============

type checkRepo struct {
	db *gorm.DB
}

func (r *checkRepo) Create(ctx context.Context, c *model.EquipmentCheck) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *checkRepo) List(ctx context.Context, filter CheckFilter, page model.PageQuery) ([]model.EquipmentCheck, int64, error) {
	q := r.db.WithContext(ctx).Model(&model.EquipmentCheck{})
	if filter.EquipmentID != nil {
		q = q.Where("equipment_id = ?", *filter.EquipmentID)
	}
	if filter.BeginTime != nil {
		q = q.Where("check_time >= ?", fmt.Sprintf("to_timestamp(%d)", *filter.BeginTime))
	}
	if filter.EndTime != nil {
		q = q.Where("check_time <= ?", fmt.Sprintf("to_timestamp(%d)", *filter.EndTime))
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var items []model.EquipmentCheck
	if err := q.Order("id DESC").Offset(page.Offset()).Limit(int(page.PageSize)).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// ============ EquipmentDowntime ============

type downtimeRepo struct {
	db *gorm.DB
}

func (r *downtimeRepo) Create(ctx context.Context, d *model.EquipmentDowntime) error {
	return r.db.WithContext(ctx).Create(d).Error
}

func (r *downtimeRepo) GetByID(ctx context.Context, id int64) (*model.EquipmentDowntime, error) {
	var d model.EquipmentDowntime
	if err := r.db.WithContext(ctx).First(&d, id).Error; err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *downtimeRepo) Update(ctx context.Context, d *model.EquipmentDowntime) error {
	return r.db.WithContext(ctx).Save(d).Error
}

func (r *downtimeRepo) List(ctx context.Context, filter DowntimeFilter, page model.PageQuery) ([]model.EquipmentDowntime, int64, error) {
	q := r.db.WithContext(ctx).Model(&model.EquipmentDowntime{})
	if filter.EquipmentID != nil {
		q = q.Where("equipment_id = ?", *filter.EquipmentID)
	}
	if filter.Status != "" {
		q = q.Where("status = ?", filter.Status)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var items []model.EquipmentDowntime
	if err := q.Order("id DESC").Offset(page.Offset()).Limit(int(page.PageSize)).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// AutoMigrate runs database migrations for all EAM models.
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Equipment{},
		&model.RepairOrder{},
		&model.EquipmentOee{},
		&model.MaintenancePlan{},
		&model.EquipmentCheck{},
		&model.EquipmentDowntime{},
	)
}
