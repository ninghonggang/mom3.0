package service

import (
	"context"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type EquipmentBomService struct {
	repo *repository.EquipmentBomRepository
}

func NewEquipmentBomService(repo *repository.EquipmentBomRepository) *EquipmentBomService {
	return &EquipmentBomService{repo: repo}
}

func (s *EquipmentBomService) Create(ctx context.Context, tenantID int64, req *model.EquipmentBomCreateReq) (*model.EquipmentBom, error) {
	bom := &model.EquipmentBom{
		TenantID:      tenantID,
		EquipmentID:   req.EquipmentID,
		EquipmentCode: req.EquipmentCode,
		EquipmentName: req.EquipmentName,
		MaterialID:    req.MaterialID,
		MaterialCode:   req.MaterialCode,
		MaterialName:   req.MaterialName,
		Quantity:       req.Quantity,
		Unit:           req.Unit,
		Position:       req.Position,
		ReplaceCycle:   req.ReplaceCycle,
		IsCritical:     req.IsCritical,
		Remark:         req.Remark,
		Status:         req.Status,
	}
	if bom.Status == 0 {
		bom.Status = 1
	}
	if err := s.repo.Create(ctx, bom); err != nil {
		return nil, err
	}
	return bom, nil
}

func (s *EquipmentBomService) GetByID(ctx context.Context, id int64) (*model.EquipmentBom, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *EquipmentBomService) ListByEquipmentID(ctx context.Context, equipmentID int64) ([]model.EquipmentBom, error) {
	return s.repo.ListByEquipmentID(ctx, equipmentID)
}

func (s *EquipmentBomService) List(ctx context.Context, tenantID int64, req *model.EquipmentBomQuery) ([]model.EquipmentBom, int64, error) {
	return s.repo.Page(ctx, tenantID, req)
}

func (s *EquipmentBomService) Update(ctx context.Context, id int64, req *model.EquipmentBomCreateReq) error {
	updates := map[string]interface{}{
		"equipment_id":   req.EquipmentID,
		"equipment_code": req.EquipmentCode,
		"equipment_name": req.EquipmentName,
		"material_id":    req.MaterialID,
		"material_code":   req.MaterialCode,
		"material_name":   req.MaterialName,
		"quantity":       req.Quantity,
		"unit":           req.Unit,
		"position":       req.Position,
		"replace_cycle":   req.ReplaceCycle,
		"is_critical":     req.IsCritical,
		"remark":         req.Remark,
		"status":         req.Status,
	}
	return s.repo.Update(ctx, id, updates)
}

func (s *EquipmentBomService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}