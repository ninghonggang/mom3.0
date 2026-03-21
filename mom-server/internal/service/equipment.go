package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type EquipmentService struct {
	repo *repository.EquipmentRepository
}

func NewEquipmentService(repo *repository.EquipmentRepository) *EquipmentService {
	return &EquipmentService{repo: repo}
}

func (s *EquipmentService) List(ctx context.Context) ([]model.Equipment, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *EquipmentService) GetByID(ctx context.Context, id string) (*model.Equipment, error) {
	var equipmentID uint
	_, err := fmt.Sscanf(id, "%d", &equipmentID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, equipmentID)
}

func (s *EquipmentService) Create(ctx context.Context, equipment *model.Equipment) error {
	return s.repo.Create(ctx, equipment)
}

func (s *EquipmentService) Update(ctx context.Context, id string, equipment *model.Equipment) error {
	var equipmentID uint
	_, err := fmt.Sscanf(id, "%d", &equipmentID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, equipmentID, map[string]interface{}{
		"equipment_name": equipment.EquipmentName,
		"equipment_type": equipment.EquipmentType,
		"brand":          equipment.Brand,
		"model":          equipment.Model,
		"status":         equipment.Status,
	})
}

func (s *EquipmentService) Delete(ctx context.Context, id string) error {
	var equipmentID uint
	_, err := fmt.Sscanf(id, "%d", &equipmentID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, equipmentID)
}

func (s *EquipmentService) GetStatus(ctx context.Context, id string) (int, error) {
	equipment, err := s.GetByID(ctx, id)
	if err != nil {
		return 0, err
	}
	return equipment.Status, nil
}
