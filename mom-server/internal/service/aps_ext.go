package service

import (
	"context"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type CapacityAnalysisService struct {
	repo *repository.CapacityAnalysisRepository
}

func NewCapacityAnalysisService(repo *repository.CapacityAnalysisRepository) *CapacityAnalysisService {
	return &CapacityAnalysisService{repo: repo}
}

func (s *CapacityAnalysisService) List(ctx context.Context, query string) ([]model.CapacityAnalysis, int64, error) {
	return s.repo.List(ctx, 1, query)
}

func (s *CapacityAnalysisService) GetByID(ctx context.Context, id uint) (*model.CapacityAnalysis, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CapacityAnalysisService) Create(ctx context.Context, a *model.CapacityAnalysis) error {
	if a.TenantID == 0 {
		a.TenantID = 1
	}
	return s.repo.Create(ctx, a)
}

func (s *CapacityAnalysisService) Update(ctx context.Context, id uint, updates map[string]any) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *CapacityAnalysisService) GetStatsByDate(ctx context.Context, startDate, endDate string) ([]model.CapacityAnalysis, error) {
	return s.repo.GetStatsByDate(ctx, 1, startDate, endDate)
}

type DeliveryRateService struct {
	repo *repository.DeliveryRateRepository
}

func NewDeliveryRateService(repo *repository.DeliveryRateRepository) *DeliveryRateService {
	return &DeliveryRateService{repo: repo}
}

func (s *DeliveryRateService) List(ctx context.Context, query string) ([]model.DeliveryRate, int64, error) {
	return s.repo.List(ctx, 1, query)
}

func (s *DeliveryRateService) GetByID(ctx context.Context, id uint) (*model.DeliveryRate, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *DeliveryRateService) Create(ctx context.Context, d *model.DeliveryRate) error {
	if d.TenantID == 0 {
		d.TenantID = 1
	}
	return s.repo.Create(ctx, d)
}

func (s *DeliveryRateService) Update(ctx context.Context, id uint, updates map[string]any) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *DeliveryRateService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

type ChangeoverMatrixService struct {
	repo *repository.ChangeoverMatrixRepository
}

func NewChangeoverMatrixService(repo *repository.ChangeoverMatrixRepository) *ChangeoverMatrixService {
	return &ChangeoverMatrixService{repo: repo}
}

func (s *ChangeoverMatrixService) List(ctx context.Context) ([]model.ChangeoverMatrix, error) {
	return s.repo.List(ctx, 1)
}

func (s *ChangeoverMatrixService) GetByID(ctx context.Context, id uint) (*model.ChangeoverMatrix, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ChangeoverMatrixService) Create(ctx context.Context, c *model.ChangeoverMatrix) error {
	if c.TenantID == 0 {
		c.TenantID = 1
	}
	return s.repo.Create(ctx, c)
}

func (s *ChangeoverMatrixService) Update(ctx context.Context, id uint, updates map[string]any) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *ChangeoverMatrixService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *ChangeoverMatrixService) GetByProducts(ctx context.Context, fromProductID, toProductID int64) (*model.ChangeoverMatrix, error) {
	return s.repo.GetByProducts(ctx, fromProductID, toProductID)
}

type RollingScheduleService struct {
	repo *repository.RollingScheduleRepository
}

func NewRollingScheduleService(repo *repository.RollingScheduleRepository) *RollingScheduleService {
	return &RollingScheduleService{repo: repo}
}

func (s *RollingScheduleService) List(ctx context.Context, query string) ([]model.RollingSchedule, int64, error) {
	return s.repo.List(ctx, 1, query)
}

func (s *RollingScheduleService) GetByID(ctx context.Context, id uint) (*model.RollingSchedule, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *RollingScheduleService) Create(ctx context.Context, r *model.RollingSchedule) error {
	if r.TenantID == 0 {
		r.TenantID = 1
	}
	return s.repo.Create(ctx, r)
}

func (s *RollingScheduleService) Update(ctx context.Context, id uint, updates map[string]any) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *RollingScheduleService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

type JITDemandService struct {
	repo *repository.JITDemandRepository
}

func NewJITDemandService(repo *repository.JITDemandRepository) *JITDemandService {
	return &JITDemandService{repo: repo}
}

func (s *JITDemandService) List(ctx context.Context, query string) ([]model.JITDemand, int64, error) {
	return s.repo.List(ctx, 1, query)
}

func (s *JITDemandService) GetByID(ctx context.Context, id uint) (*model.JITDemand, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *JITDemandService) Create(ctx context.Context, j *model.JITDemand) error {
	if j.TenantID == 0 {
		j.TenantID = 1
	}
	return s.repo.Create(ctx, j)
}

func (s *JITDemandService) Update(ctx context.Context, id uint, updates map[string]any) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *JITDemandService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
