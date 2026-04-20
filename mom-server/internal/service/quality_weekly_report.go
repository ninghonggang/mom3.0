package service

import (
	"context"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type QualityWeeklyReportService struct {
	repo *repository.QualityWeeklyReportRepository
}

func NewQualityWeeklyReportService(repo *repository.QualityWeeklyReportRepository) *QualityWeeklyReportService {
	return &QualityWeeklyReportService{repo: repo}
}

type QualityWeeklyReportQuery struct {
	TenantID   int64
	WorkshopID int64
	Year       int
	Week       int
	Page       int
	PageSize   int
}

func (s *QualityWeeklyReportService) List(ctx context.Context, query *QualityWeeklyReportQuery) ([]model.QualityWeeklyReport, int64, error) {
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 20
	}
	return s.repo.List(ctx, query.TenantID, query.Year, query.Week, query.WorkshopID, query.Page, query.PageSize)
}

func (s *QualityWeeklyReportService) GetByID(ctx context.Context, id int64) (*model.QualityWeeklyReport, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *QualityWeeklyReportService) Create(ctx context.Context, report *model.QualityWeeklyReport) error {
	return s.repo.Create(ctx, report)
}

func (s *QualityWeeklyReportService) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *QualityWeeklyReportService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *QualityWeeklyReportService) GenerateWeeklyReport(ctx context.Context, tenantID int64, year, week int, workshopID int64, workshopName string) (*model.QualityWeeklyReport, error) {
	// Check if report already exists
	existing, _ := s.repo.GetByYearWeekAndWorkshop(ctx, year, week, workshopID, tenantID)

	// Get actual quality stats
	stats, err := s.repo.GetWeeklyQualityStats(ctx, tenantID, year, week, workshopID)
	if err != nil {
		stats = map[string]interface{}{}
	}

	// Extract stats with safe type conversion
	getInt := func(m map[string]interface{}, key string) int {
		if v, ok := m[key].(int); ok {
			return v
		}
		return 0
	}
	getTime := func(m map[string]interface{}, key string) (t time.Time, ok bool) {
		if v, ok := m[key].(time.Time); ok {
			return v, true
		}
		return
	}

	totalInsp := getInt(stats, "total_inspection_qty")
	qualifiedQty := getInt(stats, "qualified_qty")
	defectQty := getInt(stats, "defect_qty")

	passRate := 0.0
	if totalInsp > 0 {
		passRate = float64(qualifiedQty) / float64(totalInsp) * 100
	}

	startDate, _ := getTime(stats, "start_date")
	endDate, _ := getTime(stats, "end_date")

	if existing != nil {
		updates := map[string]interface{}{
			"total_inspection_qty":       totalInsp,
			"qualified_qty":              qualifiedQty,
			"defect_qty":                defectQty,
			"pass_rate":                 passRate,
			"iqc_insp_qty":              getInt(stats, "iqc_insp_qty"),
			"iqc_qualified_qty":         getInt(stats, "iqc_qualified_qty"),
			"iqc_defect_qty":           getInt(stats, "iqc_defect_qty"),
			"ipqc_insp_qty":            getInt(stats, "ipqc_insp_qty"),
			"ipqc_qualified_qty":        getInt(stats, "ipqc_qualified_qty"),
			"ipqc_defect_qty":          getInt(stats, "ipqc_defect_qty"),
			"fqc_insp_qty":             getInt(stats, "fqc_insp_qty"),
			"fqc_qualified_qty":         getInt(stats, "fqc_qualified_qty"),
			"fqc_defect_qty":           getInt(stats, "fqc_defect_qty"),
			"oqc_insp_qty":             getInt(stats, "oqc_insp_qty"),
			"oqc_qualified_qty":         getInt(stats, "oqc_qualified_qty"),
			"oqc_defect_qty":           getInt(stats, "oqc_defect_qty"),
			"ncr_count":                 getInt(stats, "ncr_count"),
			"customer_complaint_count":  getInt(stats, "customer_complaint_count"),
			"start_date":                startDate,
			"end_date":                  endDate,
			"workshop_name":             workshopName,
		}
		s.repo.Update(ctx, existing.ID, updates)
		existing.TotalInspectionQty = totalInsp
		existing.QualifiedQty = qualifiedQty
		existing.DefectQty = defectQty
		existing.PassRate = passRate
		existing.StartDate = startDate
		existing.EndDate = endDate
		return existing, nil
	}

	report := &model.QualityWeeklyReport{
		TenantID:                tenantID,
		ReportYear:              year,
		ReportWeek:              week,
		WorkshopID:              workshopID,
		WorkshopName:            workshopName,
		StartDate:               startDate,
		EndDate:                 endDate,
		TotalInspectionQty:     totalInsp,
		QualifiedQty:           qualifiedQty,
		DefectQty:             defectQty,
		PassRate:              passRate,
		IQCInspQty:            getInt(stats, "iqc_insp_qty"),
		IQCQualifiedQty:       getInt(stats, "iqc_qualified_qty"),
		IQCDefectQty:          getInt(stats, "iqc_defect_qty"),
		IPQCInspQty:           getInt(stats, "ipqc_insp_qty"),
		IPQCQualifiedQty:      getInt(stats, "ipqc_qualified_qty"),
		IPQCDefectQty:         getInt(stats, "ipqc_defect_qty"),
		FQCInspQty:            getInt(stats, "fqc_insp_qty"),
		FQCQualifiedQty:       getInt(stats, "fqc_qualified_qty"),
		FQCDefectQty:         getInt(stats, "fqc_defect_qty"),
		OQCInspQty:            getInt(stats, "oqc_insp_qty"),
		OQCQualifiedQty:      getInt(stats, "oqc_qualified_qty"),
		OQCDefectQty:         getInt(stats, "oqc_defect_qty"),
		NCRCount:              getInt(stats, "ncr_count"),
		CustomerComplaintCount: getInt(stats, "customer_complaint_count"),
	}

	if err := s.repo.Create(ctx, report); err != nil {
		return nil, err
	}
	return report, nil
}