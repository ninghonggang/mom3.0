package service

import (
	"context"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type ProductionDailyReportService struct {
	repo *repository.ProductionDailyReportRepository
}

func NewProductionDailyReportService(repo *repository.ProductionDailyReportRepository) *ProductionDailyReportService {
	return &ProductionDailyReportService{repo: repo}
}

type ProductionDailyReportQuery struct {
	TenantID   int64
	WorkshopID int64
	StartDate  *time.Time
	EndDate    *time.Time
	Page       int
	PageSize   int
}

func (s *ProductionDailyReportService) List(ctx context.Context, query *ProductionDailyReportQuery) ([]model.ProductionDailyReport, int64, error) {
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 20
	}
	return s.repo.List(ctx, query.TenantID, query.StartDate, query.EndDate, query.WorkshopID, query.Page, query.PageSize)
}

func (s *ProductionDailyReportService) GetByID(ctx context.Context, id int64) (*model.ProductionDailyReport, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ProductionDailyReportService) Create(ctx context.Context, report *model.ProductionDailyReport) error {
	return s.repo.Create(ctx, report)
}

func (s *ProductionDailyReportService) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *ProductionDailyReportService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *ProductionDailyReportService) GenerateDailyReport(ctx context.Context, tenantID int64, reportDate time.Time, workshopID int64, workshopName string) (*model.ProductionDailyReport, error) {
	// Check if report already exists
	existing, _ := s.repo.GetByDateAndWorkshop(ctx, reportDate, workshopID, tenantID)

	// Get actual production stats
	stats, err := s.repo.GetProductionStatsByDate(ctx, tenantID, reportDate, workshopID)
	if err != nil {
		stats = map[string]interface{}{}
	}

	orderCount := 0
	completedCount := 0
	totalOutput := 0.0
	qualifiedQty := 0.0
	defectQty := 0.0

	if v, ok := stats["order_count"].(int); ok {
		orderCount = v
	}
	if v, ok := stats["completed_count"].(int); ok {
		completedCount = v
	}
	if v, ok := stats["total_output"].(float64); ok {
		totalOutput = v
	}
	if v, ok := stats["qualified_qty"].(float64); ok {
		qualifiedQty = v
	}
	if v, ok := stats["defect_qty"].(float64); ok {
		defectQty = v
	}

	// Calculate rates
	passRate := 0.0
	if totalOutput > 0 {
		passRate = qualifiedQty / totalOutput * 100
	}
	firstPassRate := passRate // simplified: use passRate as first pass rate
	oee := passRate * 0.85   // simplified OEE estimate
	outputPerHour := 0.0
	workerCount := 0
	workingHours := 8.0

	if completedCount > 0 {
		outputPerHour = totalOutput / float64(completedCount) / workingHours
	}

	if existing != nil {
		// Update existing report
		updates := map[string]interface{}{
			"production_order_count": orderCount,
			"completed_order_count":  completedCount,
			"total_output_qty":       totalOutput,
			"qualified_qty":         qualifiedQty,
			"defect_qty":            defectQty,
			"pass_rate":             passRate,
			"first_pass_rate":       firstPassRate,
			"oee":                   oee,
			"output_per_hour":       outputPerHour,
			"worker_count":          workerCount,
			"working_hours":         workingHours,
			"workshop_name":         workshopName,
		}
		s.repo.Update(ctx, existing.ID, updates)
		existing.ProductionOrderCount = orderCount
		existing.CompletedOrderCount = completedCount
		existing.TotalOutputQty = totalOutput
		existing.QualifiedQty = qualifiedQty
		existing.DefectQty = defectQty
		existing.PassRate = passRate
		existing.FirstPassRate = firstPassRate
		existing.OEE = oee
		existing.OutputPerHour = outputPerHour
		existing.WorkerCount = workerCount
		existing.WorkingHours = workingHours
		return existing, nil
	}

	// Create new report
	report := &model.ProductionDailyReport{
		TenantID:             tenantID,
		ReportDate:           reportDate,
		WorkshopID:           workshopID,
		WorkshopName:         workshopName,
		ProductionOrderCount: orderCount,
		CompletedOrderCount:  completedCount,
		TotalOutputQty:       totalOutput,
		QualifiedQty:         qualifiedQty,
		DefectQty:            defectQty,
		PassRate:             passRate,
		FirstPassRate:        firstPassRate,
		OEE:                  oee,
		OutputPerHour:        outputPerHour,
		WorkerCount:          workerCount,
		WorkingHours:         workingHours,
	}

	if err := s.repo.Create(ctx, report); err != nil {
		return nil, err
	}
	return report, nil
}

func (s *ProductionDailyReportService) GetSummary(ctx context.Context, tenantID int64, startDate, endDate *time.Time, workshopID int64) (map[string]interface{}, error) {
	return s.repo.GetSummary(ctx, tenantID, startDate, endDate, workshopID)
}