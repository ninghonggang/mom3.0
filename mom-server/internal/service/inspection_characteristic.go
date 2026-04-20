package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type InspectionCharacteristicService struct {
	repo *repository.InspectionCharacteristicRepository
}

func NewInspectionCharacteristicService(repo *repository.InspectionCharacteristicRepository) *InspectionCharacteristicService {
	return &InspectionCharacteristicService{repo: repo}
}

func (s *InspectionCharacteristicService) List(ctx context.Context, tenantID int64, query string) ([]model.InspectionCharacteristic, int64, error) {
	return s.repo.List(ctx, tenantID, query)
}

func (s *InspectionCharacteristicService) GetByID(ctx context.Context, id string) (*model.InspectionCharacteristic, error) {
	var uid uint
	fmt.Sscanf(id, "%d", &uid)
	return s.repo.GetByID(ctx, uid)
}

func (s *InspectionCharacteristicService) Create(ctx context.Context, char *model.InspectionCharacteristic) error {
	return s.repo.Create(ctx, char)
}

func (s *InspectionCharacteristicService) Update(ctx context.Context, id string, updates map[string]any) error {
	var uid uint
	fmt.Sscanf(id, "%d", &uid)
	return s.repo.Update(ctx, uid, updates)
}

func (s *InspectionCharacteristicService) Delete(ctx context.Context, id string) error {
	var uid uint
	fmt.Sscanf(id, "%d", &uid)
	return s.repo.Delete(ctx, uid)
}
