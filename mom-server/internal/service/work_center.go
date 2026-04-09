package service

import (
	"context"
	"fmt"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type WorkCenterService struct {
	repo *repository.WorkCenterRepository
}

func NewWorkCenterService(repo *repository.WorkCenterRepository) *WorkCenterService {
	return &WorkCenterService{repo: repo}
}

func (s *WorkCenterService) List(ctx context.Context, tenantID int64) ([]model.WorkCenter, int64, error) {
	return s.repo.List(ctx, tenantID)
}

func (s *WorkCenterService) GetByID(ctx context.Context, id string) (*model.WorkCenter, error) {
	var wcID uint
	_, err := fmt.Sscanf(id, "%d", &wcID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, wcID)
}

func (s *WorkCenterService) Create(ctx context.Context, wc *model.WorkCenter) error {
	// 检查编码唯一性
	existing, err := s.repo.GetByCode(ctx, wc.WorkCenterCode)
	if err == nil && existing.ID > 0 {
		return fmt.Errorf("工作中心编码 %s 已存在", wc.WorkCenterCode)
	}
	return s.repo.Create(ctx, wc)
}

func (s *WorkCenterService) Update(ctx context.Context, id string, wc *model.WorkCenter) error {
	var wcID uint
	_, err := fmt.Sscanf(id, "%d", &wcID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, wcID, map[string]interface{}{
		"work_center_code":   wc.WorkCenterCode,
		"work_center_name":   wc.WorkCenterName,
		"work_center_type":   wc.WorkCenterType,
		"workshop_id":        wc.WorkshopID,
		"capacity_unit":      wc.CapacityUnit,
		"standard_capacity":  wc.StandardCapacity,
		"max_capacity":       wc.MaxCapacity,
		"efficiency_factor":  wc.EfficiencyFactor,
		"utilization_target": wc.UtilizationTarget,
		"setup_time":         wc.SetupTime,
		"status":             wc.Status,
		"description":        wc.Description,
	})
}

func (s *WorkCenterService) Delete(ctx context.Context, id string) error {
	var wcID uint
	_, err := fmt.Sscanf(id, "%d", &wcID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, wcID)
}

func (s *WorkCenterService) ListByWorkshop(ctx context.Context, workshopID int64) ([]model.WorkCenter, error) {
	return s.repo.ListByWorkshop(ctx, workshopID)
}
