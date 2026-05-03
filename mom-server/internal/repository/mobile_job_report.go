package repository

import (
	"context"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

type MobileJobReportRepository struct {
	db *gorm.DB
}

func NewMobileJobReportRepository(db *gorm.DB) *MobileJobReportRepository {
	return &MobileJobReportRepository{db: db}
}

func (r *MobileJobReportRepository) List(ctx context.Context, tenantID int64, query *model.MobileJobReportQuery) ([]model.MobileJobReport, int64, error) {
	var list []model.MobileJobReport
	var total int64

	db := r.db.WithContext(ctx).Model(&model.MobileJobReport{}).Where("tenant_id = ?", tenantID)

	if query.WorkshopID > 0 {
		db = db.Where("workshop_id = ?", query.WorkshopID)
	}
	if query.ProductionLineID > 0 {
		db = db.Where("production_line_id = ?", query.ProductionLineID)
	}
	if query.WorkstationID > 0 {
		db = db.Where("workstation_id = ?", query.WorkstationID)
	}
	if query.OrderID > 0 {
		db = db.Where("order_id = ?", query.OrderID)
	}
	if query.EmployeeID > 0 {
		db = db.Where("employee_id = ?", query.EmployeeID)
	}
	if query.ReportType > 0 {
		db = db.Where("report_type = ?", query.ReportType)
	}
	if query.Status > 0 {
		db = db.Where("status = ?", query.Status)
	}
	if query.StartDate != "" {
		db = db.Where("created_at >= ?", query.StartDate)
	}
	if query.EndDate != "" {
		db = db.Where("created_at <= ?", query.EndDate)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 20
	}
	offset := (query.Page - 1) * query.PageSize

	if err := db.Order("id DESC").Offset(offset).Limit(query.PageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *MobileJobReportRepository) GetByID(ctx context.Context, id int64) (*model.MobileJobReport, error) {
	var report model.MobileJobReport
	if err := r.db.WithContext(ctx).First(&report, id).Error; err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *MobileJobReportRepository) Create(ctx context.Context, report *model.MobileJobReport) error {
	return r.db.WithContext(ctx).Create(report).Error
}

func (r *MobileJobReportRepository) Update(ctx context.Context, report *model.MobileJobReport) error {
	return r.db.WithContext(ctx).Where("id = ?", report.ID).Updates(report).Error
}

func (r *MobileJobReportRepository) GetPendingOrders(ctx context.Context, tenantID int64, workshopID int64, employeeID int64) ([]model.MobilePendingOrder, error) {
	// This queries production orders that are in progress and have pending quantities
	var list []model.MobilePendingOrder
	err := r.db.WithContext(ctx).Raw(`
		SELECT
			po.id as order_id,
			po.order_code,
			p.product_code,
			p.product_name,
			pp.process_name,
			pp.id as process_id,
			ws.id as workstation_id,
			ws.workstation_name,
			po.planned_quantity,
			COALESCE(SUM(mjr.reported_quantity), 0) as reported_quantity,
			(po.planned_quantity - COALESCE(SUM(mjr.reported_quantity), 0)) as remaining_quantity,
			po.priority,
			po.due_date
		FROM production_orders po
		LEFT JOIN mdm_product p ON po.product_id = p.id
		LEFT JOIN production_process pp ON po.process_id = pp.id
		LEFT JOIN business_workstation ws ON po.workstation_id = ws.id
		LEFT JOIN mobile_job_report mjr ON po.id = mjr.order_id AND mjr.status >= 1
		WHERE po.tenant_id = ?
			AND po.status = 2
			AND (? = 0 OR po.workshop_id = ?)
		GROUP BY po.id, p.product_code, p.product_name, pp.process_name, pp.id, ws.id, ws.workstation_name, po.planned_quantity, po.priority, po.due_date
		HAVING (po.planned_quantity - COALESCE(SUM(mjr.reported_quantity), 0)) > 0
		ORDER BY po.priority DESC, po.due_date ASC
	`, tenantID, workshopID, workshopID).Scan(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *MobileJobReportRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.MobileJobReport{}, id).Error
}