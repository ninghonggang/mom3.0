package service

import (
	"context"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"time"
)

type EquipmentCheckService struct {
	repo *repository.EquipmentCheckRepository
}

func NewEquipmentCheckService(repo *repository.EquipmentCheckRepository) *EquipmentCheckService {
	return &EquipmentCheckService{repo: repo}
}

func (s *EquipmentCheckService) List(ctx context.Context) ([]model.EquipmentCheck, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *EquipmentCheckService) Create(ctx context.Context, check *model.EquipmentCheck) error {
	return s.repo.Create(ctx, check)
}

type EquipmentMaintenanceService struct {
	repo *repository.EquipmentMaintenanceRepository
}

func NewEquipmentMaintenanceService(repo *repository.EquipmentMaintenanceRepository) *EquipmentMaintenanceService {
	return &EquipmentMaintenanceService{repo: repo}
}

func (s *EquipmentMaintenanceService) List(ctx context.Context) ([]model.EquipmentMaintenance, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *EquipmentMaintenanceService) Create(ctx context.Context, m *model.EquipmentMaintenance) error {
	return s.repo.Create(ctx, m)
}

type EquipmentRepairService struct {
	repo *repository.EquipmentRepairRepository
}

func NewEquipmentRepairService(repo *repository.EquipmentRepairRepository) *EquipmentRepairService {
	return &EquipmentRepairService{repo: repo}
}

func (s *EquipmentRepairService) List(ctx context.Context) ([]model.EquipmentRepair, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *EquipmentRepairService) Create(ctx context.Context, repair *model.EquipmentRepair) error {
	return s.repo.Create(ctx, repair)
}

func (s *EquipmentRepairService) Start(ctx context.Context, id uint) error {
	now := time.Now()
	return s.repo.Update(ctx, id, map[string]interface{}{
		"status":     2,
		"start_time": now,
	})
}

func (s *EquipmentRepairService) Complete(ctx context.Context, id uint) error {
	now := time.Now()
	repair, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	var duration int
	if repair.StartTime != nil {
		duration = int(now.Sub(*repair.StartTime).Minutes())
	}
	return s.repo.Update(ctx, id, map[string]interface{}{
		"status":     3,
		"end_time":   now,
		"duration":   duration,
	})
}

type SparePartService struct {
	repo *repository.SparePartRepository
}

func NewSparePartService(repo *repository.SparePartRepository) *SparePartService {
	return &SparePartService{repo: repo}
}

func (s *SparePartService) List(ctx context.Context) ([]model.SparePart, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *SparePartService) Create(ctx context.Context, sp *model.SparePart) error {
	return s.repo.Create(ctx, sp)
}
