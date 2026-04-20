package service

import (
	"context"
	"fmt"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type AQLService struct {
	levelRepo *repository.AQLLevelRepository
	rowRepo   *repository.AQLTableRowRepository
	planRepo  *repository.SamplingPlanRepository
}

func NewAQLService(levelRepo *repository.AQLLevelRepository, rowRepo *repository.AQLTableRowRepository, planRepo *repository.SamplingPlanRepository) *AQLService {
	return &AQLService{levelRepo: levelRepo, rowRepo: rowRepo, planRepo: planRepo}
}

func (s *AQLService) ListAQLLevels(ctx context.Context, tenantID int64) ([]model.AQLLevel, int64, error) {
	return s.levelRepo.List(ctx, tenantID)
}

func (s *AQLService) GetAQLLevel(ctx context.Context, id string) (*model.AQLLevel, error) {
	var uid uint
	fmt.Sscanf(id, "%d", &uid)
	return s.levelRepo.GetByID(ctx, uid)
}

func (s *AQLService) CreateAQLLevel(ctx context.Context, level *model.AQLLevel) error {
	return s.levelRepo.Create(ctx, level)
}

func (s *AQLService) UpdateAQLLevel(ctx context.Context, id string, updates map[string]any) error {
	var uid uint
	fmt.Sscanf(id, "%d", &uid)
	return s.levelRepo.Update(ctx, uid, updates)
}

func (s *AQLService) DeleteAQLLevel(ctx context.Context, id string) error {
	var uid uint
	fmt.Sscanf(id, "%d", &uid)
	return s.levelRepo.Delete(ctx, uid)
}

func (s *AQLService) ListAQLTableRows(ctx context.Context, aqlLevelID int64) ([]model.AQLTableRow, error) {
	return s.rowRepo.ListByLevelID(ctx, aqlLevelID)
}

func (s *AQLService) CreateAQLTableRow(ctx context.Context, row *model.AQLTableRow) error {
	return s.rowRepo.Create(ctx, row)
}

func (s *AQLService) CalculateSampleSize(ctx context.Context, tenantID int64, batchSize int, aqlValue string) (*model.AQLTableRow, error) {
	return s.rowRepo.GetByBatchAndAQL(ctx, tenantID, batchSize, aqlValue)
}

func (s *AQLService) ListSamplingPlans(ctx context.Context, tenantID int64, query string) ([]model.SamplingPlan, int64, error) {
	return s.planRepo.List(ctx, tenantID, query)
}

func (s *AQLService) GetSamplingPlan(ctx context.Context, id string) (*model.SamplingPlan, error) {
	var uid uint
	fmt.Sscanf(id, "%d", &uid)
	return s.planRepo.GetByID(ctx, uid)
}

func (s *AQLService) CreateSamplingPlan(ctx context.Context, plan *model.SamplingPlan) error {
	return s.planRepo.Create(ctx, plan)
}

func (s *AQLService) UpdateSamplingPlan(ctx context.Context, id string, updates map[string]any) error {
	var uid uint
	fmt.Sscanf(id, "%d", &uid)
	return s.planRepo.Update(ctx, uid, updates)
}

func (s *AQLService) DeleteSamplingPlan(ctx context.Context, id string) error {
	var uid uint
	fmt.Sscanf(id, "%d", &uid)
	return s.planRepo.Delete(ctx, uid)
}
