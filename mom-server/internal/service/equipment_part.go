package service

import (
	"context"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type EquipmentPartService struct {
	repo *repository.EquipmentPartRepository
}

func NewEquipmentPartService(repo *repository.EquipmentPartRepository) *EquipmentPartService {
	return &EquipmentPartService{repo: repo}
}

func (s *EquipmentPartService) List(ctx context.Context, tenantID int64, query map[string]any) ([]model.EquipmentPart, int64, error) {
	return s.repo.List(ctx, tenantID, query)
}

func (s *EquipmentPartService) GetByID(ctx context.Context, id uint) (*model.EquipmentPart, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *EquipmentPartService) Create(ctx context.Context, tenantID int64, req *model.EquipmentPartCreate) (*model.EquipmentPart, error) {
	part := &model.EquipmentPart{
		TenantID:       tenantID,
		EquipmentID:   req.EquipmentID,
		PartCode:      req.PartCode,
		PartName:      req.PartName,
		Spec:          req.Spec,
		Unit:          req.Unit,
		Qty:           req.Qty,
		Supplier:      req.Supplier,
		UnitPrice:     req.UnitPrice,
		TotalPrice:    req.TotalPrice,
		ReplacementFreq: req.ReplacementFreq,
		MaxStock:     req.MaxStock,
		MinStock:     req.MinStock,
		CurrentStock: req.CurrentStock,
		Status:       1,
	}
	if err := s.repo.Create(ctx, part); err != nil {
		return nil, err
	}
	return part, nil
}

func (s *EquipmentPartService) Update(ctx context.Context, id uint, req *model.EquipmentPartUpdate) error {
	updates := make(map[string]any)
	if req.PartCode != nil {
		updates["part_code"] = *req.PartCode
	}
	if req.PartName != nil {
		updates["part_name"] = *req.PartName
	}
	if req.Spec != nil {
		updates["spec"] = *req.Spec
	}
	if req.Unit != nil {
		updates["unit"] = *req.Unit
	}
	if req.Qty != nil {
		updates["qty"] = *req.Qty
	}
	if req.Supplier != nil {
		updates["supplier"] = *req.Supplier
	}
	if req.UnitPrice != nil {
		updates["unit_price"] = *req.UnitPrice
	}
	if req.TotalPrice != nil {
		updates["total_price"] = *req.TotalPrice
	}
	if req.ReplacementFreq != nil {
		updates["replacement_freq"] = *req.ReplacementFreq
	}
	if req.MaxStock != nil {
		updates["max_stock"] = *req.MaxStock
	}
	if req.MinStock != nil {
		updates["min_stock"] = *req.MinStock
	}
	if req.CurrentStock != nil {
		updates["current_stock"] = *req.CurrentStock
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	return s.repo.Update(ctx, id, updates)
}

func (s *EquipmentPartService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *EquipmentPartService) ListByEquipmentID(ctx context.Context, equipmentID int64) ([]model.EquipmentPart, error) {
	return s.repo.ListByEquipmentID(ctx, equipmentID)
}
