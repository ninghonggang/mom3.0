package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type WarehouseService struct {
	warehouseRepo *repository.WarehouseRepository
	locationRepo  *repository.LocationRepository
	inventoryRepo *repository.InventoryRepository
}

func NewWarehouseService(wh *repository.WarehouseRepository, loc *repository.LocationRepository, inv *repository.InventoryRepository) *WarehouseService {
	return &WarehouseService{
		warehouseRepo: wh,
		locationRepo:  loc,
		inventoryRepo: inv,
	}
}

func (s *WarehouseService) ListWarehouse(ctx context.Context) ([]model.Warehouse, int64, error) {
	return s.warehouseRepo.List(ctx, 0)
}

func (s *WarehouseService) GetWarehouseByID(ctx context.Context, id string) (*model.Warehouse, error) {
	var warehouseID uint
	_, err := fmt.Sscanf(id, "%d", &warehouseID)
	if err != nil {
		return nil, err
	}
	return s.warehouseRepo.GetByID(ctx, warehouseID)
}

func (s *WarehouseService) CreateWarehouse(ctx context.Context, warehouse *model.Warehouse) error {
	return s.warehouseRepo.Create(ctx, warehouse)
}

func (s *WarehouseService) UpdateWarehouse(ctx context.Context, id string, warehouse *model.Warehouse) error {
	var warehouseID uint
	_, err := fmt.Sscanf(id, "%d", &warehouseID)
	if err != nil {
		return err
	}
	return s.warehouseRepo.Update(ctx, warehouseID, map[string]interface{}{
		"warehouse_name": warehouse.WarehouseName,
		"warehouse_code": warehouse.WarehouseCode,
		"warehouse_type": warehouse.WarehouseType,
		"address":        warehouse.Address,
		"status":         warehouse.Status,
	})
}

func (s *WarehouseService) DeleteWarehouse(ctx context.Context, id string) error {
	var warehouseID uint
	_, err := fmt.Sscanf(id, "%d", &warehouseID)
	if err != nil {
		return err
	}
	return s.warehouseRepo.Delete(ctx, warehouseID)
}

func (s *WarehouseService) ListLocation(ctx context.Context) ([]model.Location, int64, error) {
	return s.locationRepo.List(ctx, 0)
}

func (s *WarehouseService) ListInventory(ctx context.Context) ([]model.Inventory, int64, error) {
	return s.inventoryRepo.List(ctx, 0)
}
