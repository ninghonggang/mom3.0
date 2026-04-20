package repository

import (
	"context"
	"time"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

type QualityWeeklyReportRepository struct {
	db *gorm.DB
}

func NewQualityWeeklyReportRepository(db *gorm.DB) *QualityWeeklyReportRepository {
	return &QualityWeeklyReportRepository{db: db}
}

func (r *QualityWeeklyReportRepository) List(ctx context.Context, tenantID int64, year, week int, workshopID int64, page, pageSize int) ([]model.QualityWeeklyReport, int64, error) {
	var list []model.QualityWeeklyReport
	var total int64

	query := r.db.WithContext(ctx).Model(&model.QualityWeeklyReport{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if workshopID > 0 {
		query = query.Where("workshop_id = ?", workshopID)
	}
	if year > 0 {
		query = query.Where("report_year = ?", year)
	}
	if week > 0 {
		query = query.Where("report_week = ?", week)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		query = query.Offset(offset).Limit(pageSize)
	}

	err = query.Order("report_year DESC, report_week DESC").Find(&list).Error
	return list, total, err
}

func (r *QualityWeeklyReportRepository) GetByID(ctx context.Context, id int64) (*model.QualityWeeklyReport, error) {
	var report model.QualityWeeklyReport
	err := r.db.WithContext(ctx).First(&report, id).Error
	return &report, err
}

func (r *QualityWeeklyReportRepository) GetByYearWeekAndWorkshop(ctx context.Context, year, week int, workshopID, tenantID int64) (*model.QualityWeeklyReport, error) {
	var report model.QualityWeeklyReport
	err := r.db.WithContext(ctx).
		Where("report_year = ? AND report_week = ? AND workshop_id = ? AND tenant_id = ?", year, week, workshopID, tenantID).
		First(&report).Error
	return &report, err
}

func (r *QualityWeeklyReportRepository) Create(ctx context.Context, report *model.QualityWeeklyReport) error {
	return r.db.WithContext(ctx).Create(report).Error
}

func (r *QualityWeeklyReportRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.QualityWeeklyReport{}).Where("id = ?", id).Updates(updates).Error
}

func (r *QualityWeeklyReportRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.QualityWeeklyReport{}, id).Error
}

// GetWeeklyQualityStats 获取指定年/周的质检统计数据
func (r *QualityWeeklyReportRepository) GetWeeklyQualityStats(ctx context.Context, tenantID int64, year, week int, workshopID int64) (map[string]interface{}, error) {
	type IQCStats struct {
		TotalCount     int
		QualifiedCount int
		DefectCount   int
	}
	type IPQCStats struct {
		TotalCount     int
		QualifiedCount int
		DefectCount   int
	}
	type FQCOQCStats struct {
		TotalCount     int
		QualifiedCount int
		DefectCount   int
	}
	type NCRStats struct {
		Count int
	}

	// ISO week: find Monday of week 1, then add (week-1)*7 days
	jan4 := time.Date(year, 1, 4, 0, 0, 0, 0, time.UTC)
	week1Monday := jan4.AddDate(0, 0, -int(jan4.Weekday())+1)
	if jan4.Weekday() == time.Sunday {
		week1Monday = week1Monday.AddDate(0, 0, -7)
	}
	startDate := week1Monday.AddDate(0, 0, (week-1)*7)
	endDate := startDate.AddDate(0, 0, 6)
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 0, time.UTC)

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
		"iqc_insp_qty":              iqc.TotalCount,
		"iqc_qualified_qty":         iqc.QualifiedCount,
		"iqc_defect_qty":           iqc.DefectCount,
		"ipqc_insp_qty":            ipqc.TotalCount,
		"ipqc_qualified_qty":       ipqc.QualifiedCount,
		"ipqc_defect_qty":          ipqc.DefectCount,
		"fqc_insp_qty":             fqc.TotalCount,
		"fqc_qualified_qty":        fqc.QualifiedCount,
		"fqc_defect_qty":           fqc.DefectCount,
		"oqc_insp_qty":             oqc.TotalCount,
		"oqc_qualified_qty":         oqc.QualifiedCount,
		"oqc_defect_qty":           oqc.DefectCount,
		"total_inspection_qty":     totalInsp,
		"qualified_qty":            totalQualified,
		"defect_qty":               totalDefect,
		"ncr_count":                 ncr.Count,
		"customer_complaint_count":  0,
		"start_date":               startDate,
		"end_date":                 endDate,
	}, nil
}