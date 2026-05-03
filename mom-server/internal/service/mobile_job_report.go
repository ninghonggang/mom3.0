package service

import (
	"context"
	"errors"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type MobileJobReportService struct {
	repo *repository.MobileJobReportRepository
}

func NewMobileJobReportService(repo *repository.MobileJobReportRepository) *MobileJobReportService {
	return &MobileJobReportService{repo: repo}
}

func (s *MobileJobReportService) List(ctx context.Context, tenantID int64, query *model.MobileJobReportQuery) ([]model.MobileJobReport, int64, error) {
	return s.repo.List(ctx, tenantID, query)
}

func (s *MobileJobReportService) GetByID(ctx context.Context, id int64) (*model.MobileJobReport, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *MobileJobReportService) Create(ctx context.Context, tenantID int64, req *model.MobileJobReportCreateReq) (*model.MobileJobReport, error) {
	report := &model.MobileJobReport{
		TenantID:         tenantID,
		WorkshopID:       req.WorkshopID,
		ProductionLineID:  req.ProductionLineID,
		WorkstationID:    req.WorkstationID,
		OrderID:          req.OrderID,
		ProcessID:        req.ProcessID,
		EmployeeID:       req.EmployeeID,
		ReportedQuantity: req.ReportedQuantity,
		QualifiedQuantity: req.QualifiedQuantity,
		DefectiveQuantity: req.DefectiveQuantity,
		WorkMinutes:      req.WorkMinutes,
		ReportType:       req.ReportType,
		Status:           1,
		Remarks:          req.Remarks,
	}

	if req.StartTime != "" {
		if t, err := time.Parse("2006-01-02T15:04:05", req.StartTime); err == nil {
			report.StartTime = t
		}
	}
	if req.EndTime != "" {
		if t, err := time.Parse("2006-01-02T15:04:05", req.EndTime); err == nil {
			report.EndTime = t
		}
	}

	// Default qualified to reported if not specified
	if report.QualifiedQuantity == 0 {
		report.QualifiedQuantity = report.ReportedQuantity
	}
	// Calculate defective
	if report.DefectiveQuantity == 0 {
		report.DefectiveQuantity = report.ReportedQuantity - report.QualifiedQuantity
	}

	if err := s.repo.Create(ctx, report); err != nil {
		return nil, err
	}
	return report, nil
}

func (s *MobileJobReportService) Confirm(ctx context.Context, id int64, userID int64) error {
	report, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if report == nil {
		return errors.New("report not found")
	}

	report.Status = 2
	report.ConfirmBy = userID
	now := time.Now()
	report.ConfirmAt = &now

	return s.repo.Update(ctx, report)
}

func (s *MobileJobReportService) Audit(ctx context.Context, id int64, userID int64) error {
	report, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if report == nil {
		return errors.New("report not found")
	}

	report.Status = 3
	report.AuditBy = userID
	now := time.Now()
	report.AuditAt = &now

	return s.repo.Update(ctx, report)
}

func (s *MobileJobReportService) GetPendingOrders(ctx context.Context, tenantID int64, workshopID int64, employeeID int64) ([]model.MobilePendingOrder, error) {
	return s.repo.GetPendingOrders(ctx, tenantID, workshopID, employeeID)
}

func (s *MobileJobReportService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}