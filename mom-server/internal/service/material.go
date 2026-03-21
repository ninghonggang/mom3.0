package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type MaterialService struct {
	repo *repository.MaterialRepository
}

func NewMaterialService(repo *repository.MaterialRepository) *MaterialService {
	return &MaterialService{repo: repo}
}

func (s *MaterialService) List(ctx context.Context) ([]model.Material, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *MaterialService) GetByID(ctx context.Context, id string) (*model.Material, error) {
	var materialID uint
	_, err := fmt.Sscanf(id, "%d", &materialID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, materialID)
}

func (s *MaterialService) Create(ctx context.Context, material *model.Material) error {
	return s.repo.Create(ctx, material)
}

func (s *MaterialService) Update(ctx context.Context, id string, material *model.Material) error {
	var materialID uint
	_, err := fmt.Sscanf(id, "%d", &materialID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, materialID, map[string]interface{}{
		"material_code": material.MaterialCode,
		"material_name": material.MaterialName,
		"material_type": material.MaterialType,
		"spec":         material.Spec,
		"unit":         material.Unit,
		"status":       material.Status,
	})
}

func (s *MaterialService) Delete(ctx context.Context, id string) error {
	var materialID uint
	_, err := fmt.Sscanf(id, "%d", &materialID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, materialID)
}
