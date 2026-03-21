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

func (s *WarehouseService) GetLocationByID(ctx context.Context, id string) (*model.Location, error) {
	var locationID uint
	_, err := fmt.Sscanf(id, "%d", &locationID)
	if err != nil {
		return nil, err
	}
	return s.locationRepo.GetByID(ctx, locationID)
}

func (s *WarehouseService) CreateLocation(ctx context.Context, location *model.Location) error {
	return s.locationRepo.Create(ctx, location)
}

func (s *WarehouseService) UpdateLocation(ctx context.Context, id string, location *model.Location) error {
	var locationID uint
	_, err := fmt.Sscanf(id, "%d", &locationID)
	if err != nil {
		return err
	}
	return s.locationRepo.Update(ctx, locationID, map[string]interface{}{
		"location_code":  location.LocationCode,
		"location_name":  location.LocationName,
		"warehouse_id":   location.WarehouseID,
		"zone_code":      location.ZoneCode,
		"row":            location.Row,
		"col":            location.Col,
		"layer":          location.Layer,
		"location_type":  location.LocationType,
		"capacity":       location.Capacity,
		"status":         location.Status,
	})
}

func (s *WarehouseService) DeleteLocation(ctx context.Context, id string) error {
	var locationID uint
	_, err := fmt.Sscanf(id, "%d", &locationID)
	if err != nil {
		return err
	}
	return s.locationRepo.Delete(ctx, locationID)
}

func (s *WarehouseService) ListInventory(ctx context.Context) ([]model.Inventory, int64, error) {
	return s.inventoryRepo.List(ctx, 0)
}

func (s *WarehouseService) GetInventoryByID(ctx context.Context, id string) (*model.Inventory, error) {
	var inventoryID uint
	_, err := fmt.Sscanf(id, "%d", &inventoryID)
	if err != nil {
		return nil, err
	}
	return s.inventoryRepo.GetByID(ctx, inventoryID)
}

func (s *WarehouseService) CreateInventory(ctx context.Context, inventory *model.Inventory) error {
	return s.inventoryRepo.Create(ctx, inventory)
}

func (s *WarehouseService) UpdateInventory(ctx context.Context, id string, inventory *model.Inventory) error {
	var inventoryID uint
	_, err := fmt.Sscanf(id, "%d", &inventoryID)
	if err != nil {
		return err
	}
	return s.inventoryRepo.Update(ctx, inventoryID, map[string]interface{}{
		"material_id":    inventory.MaterialID,
		"material_code":  inventory.MaterialCode,
		"material_name":  inventory.MaterialName,
		"warehouse_id":   inventory.WarehouseID,
		"location_id":    inventory.LocationID,
		"quantity":       inventory.Quantity,
		"available_qty":  inventory.AvailableQty,
		"allocated_qty":  inventory.AllocatedQty,
		"locked_qty":     inventory.LockedQty,
		"batch_no":       inventory.BatchNo,
	})
}

func (s *WarehouseService) DeleteInventory(ctx context.Context, id string) error {
	var inventoryID uint
	_, err := fmt.Sscanf(id, "%d", &inventoryID)
	if err != nil {
		return err
	}
	return s.inventoryRepo.Delete(ctx, inventoryID)
}
