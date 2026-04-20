package service

import (
	"context"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type DeliveryReportService struct {
	repo *repository.DeliveryReportRepository
}

func NewDeliveryReportService(repo *repository.DeliveryReportRepository) *DeliveryReportService {
	return &DeliveryReportService{repo: repo}
}

type DeliveryReportQuery struct {
	TenantID    int64
	CustomerID  int64
	StartMonth  *time.Time
	EndMonth    *time.Time
	Page        int
	PageSize    int
}

func (s *DeliveryReportService) List(ctx context.Context, query *DeliveryReportQuery) ([]model.DeliveryReport, int64, error) {
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 20
	}
	return s.repo.List(ctx, query.TenantID, query.StartMonth, query.EndMonth, query.CustomerID, query.Page, query.PageSize)
}

func (s *DeliveryReportService) GetByID(ctx context.Context, id int64) (*model.DeliveryReport, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *DeliveryReportService) Create(ctx context.Context, report *model.DeliveryReport) error {
	return s.repo.Create(ctx, report)
}

func (s *DeliveryReportService) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *DeliveryReportService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *DeliveryReportService) GenerateDeliveryReport(ctx context.Context, tenantID int64, reportMonth time.Time, customerID int64, customerName string) (*model.DeliveryReport, error) {
	// Check if report already exists
	existing, _ := s.repo.GetByMonthAndCustomer(ctx, reportMonth, customerID, tenantID)
	if existing != nil {
		return existing, nil
	}

	// Create new report (placeholder for actual business logic)
	report := &model.DeliveryReport{
		TenantID:    tenantID,
		ReportMonth: reportMonth,
		CustomerID:  customerID,
		CustomerName: customerName,
	}

	if err := s.repo.Create(ctx, report); err != nil {
		return nil, err
	}
	return report, nil
}