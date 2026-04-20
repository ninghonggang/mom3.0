package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type SupplierMaterialService struct {
	repo *repository.SupplierMaterialRepository
}

func NewSupplierMaterialService(repo *repository.SupplierMaterialRepository) *SupplierMaterialService {
	return &SupplierMaterialService{repo: repo}
}

func (s *SupplierMaterialService) List(ctx context.Context, tenantID int64, query string) ([]model.SupplierMaterial, int64, error) {
	return s.repo.List(ctx, tenantID, query)
}

func (s *SupplierMaterialService) GetByID(ctx context.Context, id string) (*model.SupplierMaterial, error) {
	var itemID uint
	_, err := fmt.Sscanf(id, "%d", &itemID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, itemID)
}

func (s *SupplierMaterialService) ListBySupplier(ctx context.Context, supplierID string) ([]model.SupplierMaterial, error) {
	var sid int64
	_, err := fmt.Sscanf(supplierID, "%d", &sid)
	if err != nil {
		return nil, err
	}
	return s.repo.ListBySupplier(ctx, sid)
}

func (s *SupplierMaterialService) ListByMaterial(ctx context.Context, materialID string) ([]model.SupplierMaterial, error) {
	var mid int64
	_, err := fmt.Sscanf(materialID, "%d", &mid)
	if err != nil {
		return nil, err
	}
	return s.repo.ListByMaterial(ctx, mid)
}

func (s *SupplierMaterialService) Create(ctx context.Context, item *model.SupplierMaterial) error {
	return s.repo.Create(ctx, item)
}

func (s *SupplierMaterialService) Update(ctx context.Context, id string, item *model.SupplierMaterial) error {
	var itemID uint
	_, err := fmt.Sscanf(id, "%d", &itemID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, itemID, map[string]any{
		"supplier_id":     item.SupplierID,
		"material_id":     item.MaterialID,
		"material_code":   item.MaterialCode,
		"material_name":   item.MaterialName,
		"supplier_part_no": item.SupplierPartNo,
		"price":          item.Price,
		"currency":       item.Currency,
		"min_order_qty":  item.MinOrderQty,
		"lead_time":      item.LeadTime,
		"is_preferred":   item.IsPreferred,
		"status":         item.Status,
	})
}

func (s *SupplierMaterialService) Delete(ctx context.Context, id string) error {
	var itemID uint
	_, err := fmt.Sscanf(id, "%d", &itemID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, itemID)
}

func (s *SupplierMaterialService) SetPreferred(ctx context.Context, id string) error {
	var itemID uint
	_, err := fmt.Sscanf(id, "%d", &itemID)
	if err != nil {
		return err
	}
	// 先获取记录
	item, err := s.repo.GetByID(ctx, itemID)
	if err != nil {
		return err
	}
	// 清除同供应商下的首选
	if err := s.repo.ClearPreferredBySupplier(ctx, item.SupplierID); err != nil {
		return err
	}
	// 设置新的首选
	return s.repo.Update(ctx, itemID, map[string]any{
		"is_preferred": 1,
	})
}
