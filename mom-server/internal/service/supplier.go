package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type SupplierService struct {
	repo *repository.SupplierRepository
}

func NewSupplierService(repo *repository.SupplierRepository) *SupplierService {
	return &SupplierService{repo: repo}
}

func (s *SupplierService) List(ctx context.Context, tenantID int64) ([]model.Supplier, int64, error) {
	return s.repo.List(ctx, tenantID)
}

func (s *SupplierService) GetByID(ctx context.Context, id string) (*model.Supplier, error) {
	var supplierID uint
	_, err := fmt.Sscanf(id, "%d", &supplierID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, supplierID)
}

func (s *SupplierService) Create(ctx context.Context, supplier *model.Supplier) error {
	return s.repo.Create(ctx, supplier)
}

func (s *SupplierService) Update(ctx context.Context, id string, supplier *model.Supplier) error {
	var supplierID uint
	_, err := fmt.Sscanf(id, "%d", &supplierID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, supplierID, map[string]interface{}{
		"code":     supplier.Code,
		"name":     supplier.Name,
		"type":     supplier.Type,
		"contact":  supplier.Contact,
		"phone":    supplier.Phone,
		"email":    supplier.Email,
		"address":  supplier.Address,
		"category": supplier.Category,
		"level":    supplier.Level,
		"status":   supplier.Status,
		"remark":   supplier.Remark,
	})
}

func (s *SupplierService) Delete(ctx context.Context, id string) error {
	var supplierID uint
	_, err := fmt.Sscanf(id, "%d", &supplierID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, supplierID)
}
