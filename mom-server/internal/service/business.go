package service

import (
	"context"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type ProductionLineService struct {
	repo *repository.ProductionLineRepository
}

func NewProductionLineService(repo *repository.ProductionLineRepository) *ProductionLineService {
	return &ProductionLineService{repo: repo}
}

func (s *ProductionLineService) List(ctx context.Context) ([]model.ProductionLine, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *ProductionLineService) Create(ctx context.Context, line *model.ProductionLine) error {
	return s.repo.Create(ctx, line)
}

func (s *ProductionLineService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *ProductionLineService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

type WorkstationService struct {
	repo *repository.WorkstationRepository
}

func NewWorkstationService(repo *repository.WorkstationRepository) *WorkstationService {
	return &WorkstationService{repo: repo}
}

func (s *WorkstationService) List(ctx context.Context) ([]model.Workstation, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *WorkstationService) Create(ctx context.Context, ws *model.Workstation) error {
	return s.repo.Create(ctx, ws)
}

func (s *WorkstationService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *WorkstationService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

type ShiftService struct {
	repo *repository.ShiftRepository
}

func NewShiftService(repo *repository.ShiftRepository) *ShiftService {
	return &ShiftService{repo: repo}
}

func (s *ShiftService) List(ctx context.Context) ([]model.Shift, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *ShiftService) Create(ctx context.Context, shift *model.Shift) error {
	return s.repo.Create(ctx, shift)
}

func (s *ShiftService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *ShiftService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
