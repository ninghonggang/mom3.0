package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type ProductionReportService struct {
	repo *repository.ProductionReportRepository
}

func NewProductionReportService(repo *repository.ProductionReportRepository) *ProductionReportService {
	return &ProductionReportService{repo: repo}
}

func (s *ProductionReportService) List(ctx context.Context) ([]model.ProductionReport, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *ProductionReportService) Create(ctx context.Context, report *model.ProductionReport) error {
	return s.repo.Create(ctx, report)
}

type DispatchService struct {
	repo *repository.DispatchRepository
}

func NewDispatchService(repo *repository.DispatchRepository) *DispatchService {
	return &DispatchService{repo: repo}
}

func (s *DispatchService) List(ctx context.Context) ([]model.Dispatch, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *DispatchService) Create(ctx context.Context, dispatch *model.Dispatch) error {
	return s.repo.Create(ctx, dispatch)
}

func (s *DispatchService) Update(ctx context.Context, id string, dispatch *model.Dispatch) error {
	var dispatchID uint
	_, err := fmt.Sscanf(id, "%d", &dispatchID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, dispatchID, map[string]interface{}{
		"order_id":      dispatch.OrderID,
		"order_no":      dispatch.OrderNo,
		"process_id":    dispatch.ProcessID,
		"process_name":  dispatch.ProcessName,
		"station_id":    dispatch.StationID,
		"station_name":   dispatch.StationName,
		"assign_user_id":  dispatch.AssignUserID,
		"assign_user_name": dispatch.AssignUserName,
		"quantity":      dispatch.Quantity,
		"status":        dispatch.Status,
	})
}

func (s *DispatchService) Start(ctx context.Context, id string) error {
	var dispatchID uint
	_, err := fmt.Sscanf(id, "%d", &dispatchID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, dispatchID, map[string]interface{}{
		"status": 2,
	})
}

func (s *DispatchService) Complete(ctx context.Context, id string) error {
	var dispatchID uint
	_, err := fmt.Sscanf(id, "%d", &dispatchID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, dispatchID, map[string]interface{}{
		"status": 3,
	})
}
