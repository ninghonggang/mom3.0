package repository

import (
	"context"
	"time"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

type ProductionDailyReportRepository struct {
	db *gorm.DB
}

func NewProductionDailyReportRepository(db *gorm.DB) *ProductionDailyReportRepository {
	return &ProductionDailyReportRepository{db: db}
}

func (r *ProductionDailyReportRepository) List(ctx context.Context, tenantID int64, startDate, endDate *time.Time, workshopID int64, page, pageSize int) ([]model.ProductionDailyReport, int64, error) {
	var list []model.ProductionDailyReport
	var total int64

	query := r.db.WithContext(ctx).Model(&model.ProductionDailyReport{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if workshopID > 0 {
		query = query.Where("workshop_id = ?", workshopID)
	}
	if startDate != nil {
		query = query.Where("report_date >= ?", startDate)
	}
	if endDate != nil {
		query = query.Where("report_date <= ?", endDate)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		query = query.Offset(offset).Limit(pageSize)
	}

	err = query.Order("report_date DESC").Find(&list).Error
	return list, total, err
}

func (r *ProductionDailyReportRepository) GetByID(ctx context.Context, id int64) (*model.ProductionDailyReport, error) {
	var report model.ProductionDailyReport
	err := r.db.WithContext(ctx).First(&report, id).Error
	return &report, err
}

func (r *ProductionDailyReportRepository) GetByDateAndWorkshop(ctx context.Context, reportDate time.Time, workshopID, tenantID int64) (*model.ProductionDailyReport, error) {
	var report model.ProductionDailyReport
	err := r.db.WithContext(ctx).
		Where("report_date = ? AND workshop_id = ? AND tenant_id = ?", reportDate, workshopID, tenantID).
		First(&report).Error
	return &report, err
}

func (r *ProductionDailyReportRepository) Create(ctx context.Context, report *model.ProductionDailyReport) error {
	return r.db.WithContext(ctx).Create(report).Error
}

func (r *ProductionDailyReportRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ProductionDailyReport{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ProductionDailyReportRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.ProductionDailyReport{}, id).Error
}

// GetWeeklyQualityStats 获取指定年/周的质检统计数据（供生成周报使用）
func (r *ProductionDailyReportRepository) GetWeeklyQualityStats(ctx context.Context, tenantID int64, year, week int, workshopID int64) (map[string]interface{}, error) {
	type IQCStats struct {
		TotalCount    int
		QualifiedCount int
		DefectCount   int
	}
	type IPQCStats struct {
		TotalCount    int
		QualifiedCount int
		DefectCount   int
	}
	type FQCOQCStats struct {
		TotalCount    int
		QualifiedCount int
		DefectCount   int
	}
	type NCRStats struct {
		Count int
	}
	type ComplaintStats struct {
		Count int
	}

	// Calculate week start and end dates (ISO week: Monday as first day)
	jan4 := time.Date(year, 1, 4, 0, 0, 0, 0, time.UTC)
	week1Monday := jan4.AddDate(0, 0, -int(jan4.Weekday())+1)
	if jan4.Weekday() == time.Sunday {
		week1Monday = week1Monday.AddDate(0, 0, -7)
	}
	startDate := week1Monday.AddDate(0, 0, (week-1)*7)
	endDate := startDate.AddDate(0, 0, 6)

	var iqc IQCStats
	r.db.WithContext(ctx).Model(&model.IQC{}).
		Where("tenant_id = ?", tenantID).
		Where("check_date >= ? AND check_date <= ?", startDate, endDate).
		Select("COUNT(*) as total_count, "+
			"SUM(CASE WHEN result = 2 THEN 1 ELSE 0 END) as qualified_count, "+
			"SUM(CASE WHEN result = 3 THEN 1 ELSE 0 END) as defect_count").
		Scan(&iqc)

	var ipqc IPQCStats
	r.db.WithContext(ctx).Model(&model.IPQC{}).
		Where("tenant_id = ?", tenantID).
		Where("check_date >= ? AND check_date <= ?", startDate, endDate).
		Select("COUNT(*) as total_count, "+
			"SUM(CASE WHEN result = 2 THEN 1 ELSE 0 END) as qualified_count, "+
			"SUM(CASE WHEN result = 3 THEN 1 ELSE 0 END) as defect_count").
		Scan(&ipqc)

	var fqc FQCOQCStats
	r.db.WithContext(ctx).Model(&model.FQC{}).
		Where("tenant_id = ?", tenantID).
		Where("check_date >= ? AND check_date <= ?", startDate, endDate).
		Select("COUNT(*) as total_count, "+
			"SUM(CASE WHEN result = 2 THEN 1 ELSE 0 END) as qualified_count, "+
			"SUM(CASE WHEN result = 3 THEN 1 ELSE 0 END) as defect_count").
		Scan(&fqc)

	var oqc FQCOQCStats
	r.db.WithContext(ctx).Model(&model.OQC{}).
		Where("tenant_id = ?", tenantID).
		Where("check_date >= ? AND check_date <= ?", startDate, endDate).
		Select("COUNT(*) as total_count, "+
			"SUM(CASE WHEN result = 2 THEN 1 ELSE 0 END) as qualified_count, "+
			"SUM(CASE WHEN result = 3 THEN 1 ELSE 0 END) as defect_count").
		Scan(&oqc)

	var ncr NCRStats
	r.db.WithContext(ctx).Model(&model.NCR{}).
		Where("tenant_id = ?", tenantID).
		Where("check_date >= ? AND check_date <= ?", startDate, endDate).
		Select("COUNT(*) as count").
		Scan(&ncr)

	totalInsp := iqc.TotalCount + ipqc.TotalCount + fqc.TotalCount + oqc.TotalCount
	totalQualified := iqc.QualifiedCount + ipqc.QualifiedCount + fqc.QualifiedCount + oqc.QualifiedCount
	totalDefect := iqc.DefectCount + ipqc.DefectCount + fqc.DefectCount + oqc.DefectCount

	return map[string]interface{}{
		"iqc_insp_qty":       iqc.TotalCount,
		"iqc_qualified_qty":  iqc.QualifiedCount,
		"iqc_defect_qty":    iqc.DefectCount,
		"ipqc_insp_qty":     ipqc.TotalCount,
		"ipqc_qualified_qty": ipqc.QualifiedCount,
		"ipqc_defect_qty":   ipqc.DefectCount,
		"fqc_insp_qty":      fqc.TotalCount,
		"fqc_qualified_qty":  fqc.QualifiedCount,
		"fqc_defect_qty":    fqc.DefectCount,
		"oqc_insp_qty":      oqc.TotalCount,
		"oqc_qualified_qty":  oqc.QualifiedCount,
		"oqc_defect_qty":    oqc.DefectCount,
		"total_inspection_qty": totalInsp,
		"qualified_qty":       totalQualified,
		"defect_qty":          totalDefect,
		"ncr_count":           ncr.Count,
		"customer_complaint_count": 0, // 客诉需单独表，暂置0
		"start_date":          startDate,
		"end_date":            endDate,
	}, nil
}

// GetProductionStatsByDate 获取指定日期的生产统计数据（供生成日报使用）
func (r *ProductionDailyReportRepository) GetProductionStatsByDate(ctx context.Context, tenantID int64, reportDate time.Time, workshopID int64) (map[string]interface{}, error) {
	type Stats struct {
		OrderCount      int
		CompletedCount  int
		TotalOutput     float64
		QualifiedQty    float64
		DefectQty       float64
	}
	var stats Stats

	// 按完成日期统计
	query := r.db.WithContext(ctx).Model(&model.ProductionOrder{}).
		Where("tenant_id = ?", tenantID).
		Where("DATE(actual_end_date) = ?", reportDate.Format("2006-01-02")).
		Where("status = 3") // 已完成
	if workshopID > 0 {
		query = query.Where("workshop_id = ?", workshopID)
	}

	err := query.Select("COUNT(*) as order_count, "+
		"SUM(CASE WHEN status = 3 THEN 1 ELSE 0 END) as completed_count, "+
		"SUM(completed_qty) as total_output, "+
		"SUM(completed_qty - rejected_qty) as qualified_qty, "+
		"SUM(rejected_qty) as defect_qty").
		Scan(&stats).Error
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"order_count":     stats.OrderCount,
		"completed_count": stats.CompletedCount,
		"total_output":    stats.TotalOutput,
		"qualified_qty":   stats.QualifiedQty,
		"defect_qty":      stats.DefectQty,
	}, nil
}

func (r *ProductionDailyReportRepository) GetSummary(ctx context.Context, tenantID int64, startDate, endDate *time.Time, workshopID int64) (map[string]interface{}, error) {
	type Summary struct {
		TotalOrders        int
		CompletedOrders    int
		TotalOutput        float64
		TotalQualified     float64
		TotalDefect        float64
		AvgPassRate        float64
		AvgFirstPassRate   float64
		AvgOEE             float64
	}
	var summary Summary

	query := r.db.WithContext(ctx).Model(&model.ProductionDailyReport{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if workshopID > 0 {
		query = query.Where("workshop_id = ?", workshopID)
	}
	if startDate != nil {
		query = query.Where("report_date >= ?", startDate)
	}
	if endDate != nil {
		query = query.Where("report_date <= ?", endDate)
	}

	err := query.Select("COUNT(*) as total_orders, SUM(completed_order_count) as completed_orders, "+
		"SUM(total_output_qty) as total_output, SUM(qualified_qty) as total_qualified, "+
		"SUM(defect_qty) as total_defect, AVG(pass_rate) as avg_pass_rate, "+
		"AVG(first_pass_rate) as avg_first_pass_rate, AVG(oee) as avg_oee").Scan(&summary).Error
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"total_orders":      summary.TotalOrders,
		"completed_orders":  summary.CompletedOrders,
		"total_output":      summary.TotalOutput,
		"total_qualified":   summary.TotalQualified,
		"total_defect":     summary.TotalDefect,
		"avg_pass_rate":    summary.AvgPassRate,
		"avg_first_pass_rate": summary.AvgFirstPassRate,
		"avg_oee":          summary.AvgOEE,
	}, nil
}