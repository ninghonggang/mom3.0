package repository

import (
	"context"

	"github.com/ninghonggang/mom-platform/services/mes-service/internal/model"
	"gorm.io/gorm"
)

// --- Order Repository (GORM) ---

type orderRepo struct{ db *gorm.DB }

func NewOrderRepository(db *gorm.DB) OrderRepository { return &orderRepo{db: db} }

func (r *orderRepo) GetByID(ctx context.Context, id int64) (*model.ProductionOrder, error) {
	var o model.ProductionOrder
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&o).Error
	if err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *orderRepo) List(ctx context.Context, f OrderFilter) ([]model.ProductionOrder, int64, error) {
	var orders []model.ProductionOrder
	var total int64

	q := r.db.WithContext(ctx).Model(&model.ProductionOrder{})
	if f.TenantID != 0 {
		q = q.Where("tenant_id = ?", f.TenantID)
	}

	if f.Keyword != "" {
		q = q.Where("order_no ILIKE ? OR material_code ILIKE ? OR material_name ILIKE ?",
			"%"+f.Keyword+"%", "%"+f.Keyword+"%", "%"+f.Keyword+"%")
	}
	if f.WorkshopID > 0 {
		q = q.Where("workshop_id = ?", f.WorkshopID)
	}
	if f.LineID > 0 {
		q = q.Where("line_id = ?", f.LineID)
	}
	if f.Status != nil {
		q = q.Where("status = ?", *f.Status)
	}
	if f.DateFrom != "" {
		q = q.Where("created_at >= ?", f.DateFrom)
	}
	if f.DateTo != "" {
		q = q.Where("created_at <= ?", f.DateTo)
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if f.Page > 0 && f.PageSize > 0 {
		q = q.Offset((f.Page - 1) * f.PageSize).Limit(f.PageSize)
	}
	q = q.Order("created_at DESC")

	if err := q.Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	return orders, total, nil
}

func (r *orderRepo) Create(ctx context.Context, o *model.ProductionOrder) error {
	return r.db.WithContext(ctx).Create(o).Error
}

func (r *orderRepo) UpdateStatus(ctx context.Context, id int64, status model.OrderStatus) error {
	return r.db.WithContext(ctx).Model(&model.ProductionOrder{}).
		Where("id = ?", id).Update("status", status).Error
}

func (r *orderRepo) Update(ctx context.Context, o *model.ProductionOrder) error {
	return r.db.WithContext(ctx).Save(o).Error
}

// --- Report Repository (GORM) ---

type reportRepo struct{ db *gorm.DB }

func NewReportRepository(db *gorm.DB) JobReportRepository { return &reportRepo{db: db} }

func (r *reportRepo) GetByID(ctx context.Context, id int64) (*model.MobileJobReport, error) {
	var m model.MobileJobReport
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *reportRepo) List(ctx context.Context, f ReportFilter) ([]model.MobileJobReport, int64, error) {
	var reports []model.MobileJobReport
	var total int64

	q := r.db.WithContext(ctx).Model(&model.MobileJobReport{})
	if f.TenantID != 0 {
		q = q.Where("tenant_id = ?", f.TenantID)
	}

	if f.OrderID > 0 {
		q = q.Where("order_id = ?", f.OrderID)
	}
	if f.EmployeeID > 0 {
		q = q.Where("employee_id = ?", f.EmployeeID)
	}
	if f.DateFrom != "" {
		q = q.Where("created_at >= ?", f.DateFrom)
	}
	if f.DateTo != "" {
		q = q.Where("created_at <= ?", f.DateTo)
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if f.Page > 0 && f.PageSize > 0 {
		q = q.Offset((f.Page - 1) * f.PageSize).Limit(f.PageSize)
	}
	q = q.Order("created_at DESC")

	if err := q.Find(&reports).Error; err != nil {
		return nil, 0, err
	}
	return reports, total, nil
}

func (r *reportRepo) Create(ctx context.Context, m *model.MobileJobReport) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *reportRepo) UpdateStatus(ctx context.Context, id int64, status model.ReportStatus) error {
	return r.db.WithContext(ctx).Model(&model.MobileJobReport{}).
		Where("id = ?", id).Update("status", status).Error
}

// --- Dispatch Repository (GORM) ---

type dispatchRepo struct{ db *gorm.DB }

func NewDispatchRepository(db *gorm.DB) DispatchRepository { return &dispatchRepo{db: db} }

func (r *dispatchRepo) List(ctx context.Context, f DispatchFilter) ([]model.Dispatch, error) {
	var dispatches []model.Dispatch
	q := r.db.WithContext(ctx).Model(&model.Dispatch{})
	if f.TenantID != 0 {
		q = q.Where("tenant_id = ?", f.TenantID)
	}
	if f.OrderID > 0 {
		q = q.Where("order_id = ?", f.OrderID)
	}
	if f.LineID > 0 {
		q = q.Where("line_id = ?", f.LineID)
	}
	if f.WorkstationID > 0 {
		q = q.Where("workstation_id = ?", f.WorkstationID)
	}
	if err := q.Find(&dispatches).Error; err != nil {
		return nil, err
	}
	return dispatches, nil
}

func (r *dispatchRepo) CreateBatch(ctx context.Context, dispatches []model.Dispatch) error {
	return r.db.WithContext(ctx).Create(&dispatches).Error
}

// --- ProductionComplete Repository (GORM) ---

type completeRepo struct{ db *gorm.DB }

// NewCompleteRepository returns a GORM-backed CompleteRepository.
func NewCompleteRepository(db *gorm.DB) CompleteRepository { return &completeRepo{db: db} }

func (r *completeRepo) Create(ctx context.Context, c *model.ProductionComplete) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *completeRepo) GetByID(ctx context.Context, id int64) (*model.ProductionComplete, error) {
	var c model.ProductionComplete
	if err := r.db.WithContext(ctx).First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *completeRepo) ListByOrder(ctx context.Context, orderID int64) ([]model.ProductionComplete, error) {
	var list []model.ProductionComplete
	if err := r.db.WithContext(ctx).
		Where("order_id = ?", orderID).
		Order("id DESC").
		Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// SumQtyByOrder returns the total already-completed quantity for an order.
func (r *completeRepo) SumQtyByOrder(ctx context.Context, orderID int64) (float64, error) {
	var total *float64
	if err := r.db.WithContext(ctx).
		Model(&model.ProductionComplete{}).
		Where("order_id = ?", orderID).
		Select("COALESCE(SUM(quantity), 0)").
		Scan(&total).Error; err != nil {
		return 0, err
	}
	if total == nil {
		return 0, nil
	}
	return *total, nil
}
