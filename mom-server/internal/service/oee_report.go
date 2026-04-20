package service

import (
	"context"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type OEEReportService struct {
	repo *repository.OEEReportRepository
}

func NewOEEReportService(repo *repository.OEEReportRepository) *OEEReportService {
	return &OEEReportService{repo: repo}
}

type OEEReportQuery struct {
	TenantID   int64
	WorkshopID int64
	LineID     int64
	StartDate  *time.Time
	EndDate    *time.Time
	Page       int
	PageSize   int
}

func (s *OEEReportService) List(ctx context.Context, query *OEEReportQuery) ([]model.OEEReport, int64, error) {
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 20
	}
	return s.repo.List(ctx, query.TenantID, query.StartDate, query.EndDate, query.WorkshopID, query.LineID, query.Page, query.PageSize)
}

func (s *OEEReportService) GetByID(ctx context.Context, id int64) (*model.OEEReport, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *OEEReportService) Create(ctx context.Context, report *model.OEEReport) error {
	return s.repo.Create(ctx, report)
}

func (s *OEEReportService) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *OEEReportService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *OEEReportService) GenerateOEEReport(ctx context.Context, tenantID int64, reportDate time.Time, workshopID int64, workshopName string, lineID int64, lineName string) (*model.OEEReport, error) {
	// Check if report already exists
	existing, _ := s.repo.GetByDateAndLine(ctx, reportDate, lineID, tenantID)
	if existing != nil {
		return existing, nil
	}

	// Create new report (placeholder for actual business logic)
	report := &model.OEEReport{
		TenantID:   tenantID,
		ReportDate: reportDate,
		WorkshopID: workshopID,
		WorkshopName: workshopName,
		LineID:     lineID,
		LineName:   lineName,
	}

	if err := s.repo.Create(ctx, report); err != nil {
		return nil, err
	}
	return report, nil
}