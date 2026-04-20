package service

import (
	"context"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type AndonReportService struct {
	repo *repository.AndonReportRepository
}

func NewAndonReportService(repo *repository.AndonReportRepository) *AndonReportService {
	return &AndonReportService{repo: repo}
}

type AndonReportQuery struct {
	TenantID   int64
	WorkshopID int64
	LineID     int64
	StationID  int64
	StartDate  *time.Time
	EndDate    *time.Time
	Page       int
	PageSize   int
}

func (s *AndonReportService) List(ctx context.Context, query *AndonReportQuery) ([]model.AndonReport, int64, error) {
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 20
	}
	return s.repo.List(ctx, query.TenantID, query.StartDate, query.EndDate, query.WorkshopID, query.LineID, query.StationID, query.Page, query.PageSize)
}

func (s *AndonReportService) GetByID(ctx context.Context, id int64) (*model.AndonReport, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *AndonReportService) Create(ctx context.Context, report *model.AndonReport) error {
	return s.repo.Create(ctx, report)
}

func (s *AndonReportService) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *AndonReportService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *AndonReportService) GenerateAndonReport(ctx context.Context, tenantID int64, reportDate time.Time, workshopID int64, workshopName string, lineID int64, lineName string, stationID int64, stationName string) (*model.AndonReport, error) {
	// Check if report already exists
	existing, _ := s.repo.GetByDateAndStation(ctx, reportDate, stationID, tenantID)
	if existing != nil {
		return existing, nil
	}

	// Create new report (placeholder for actual business logic)
	report := &model.AndonReport{
		TenantID:   tenantID,
		ReportDate: reportDate,
		WorkshopID: workshopID,
		WorkshopName: workshopName,
		LineID:     lineID,
		LineName:   lineName,
		StationID:  stationID,
		StationName: stationName,
	}

	if err := s.repo.Create(ctx, report); err != nil {
		return nil, err
	}
	return report, nil
}