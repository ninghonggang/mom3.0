package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type WarehouseService struct {
	db                 *gorm.DB
	warehouseRepo      *repository.WarehouseRepository
	locationRepo       *repository.LocationRepository
	inventoryRepo      *repository.InventoryRepository
	receiveOrderRepo   *repository.ReceiveOrderRepository
	receiveOrderItemRepo *repository.ReceiveOrderItemRepository
	deliveryOrderRepo  *repository.DeliveryOrderRepository
	deliveryOrderItemRepo *repository.DeliveryOrderItemRepository
}

func NewWarehouseService(db *gorm.DB, wh *repository.WarehouseRepository, loc *repository.LocationRepository, inv *repository.InventoryRepository, recvRepo *repository.ReceiveOrderRepository, recvItemRepo *repository.ReceiveOrderItemRepository, delRepo *repository.DeliveryOrderRepository, delItemRepo *repository.DeliveryOrderItemRepository) *WarehouseService {
	return &WarehouseService{
		db:                   db,
		warehouseRepo:        wh,
		locationRepo:         loc,
		inventoryRepo:        inv,
		receiveOrderRepo:     recvRepo,
		receiveOrderItemRepo: recvItemRepo,
		deliveryOrderRepo:    delRepo,
		deliveryOrderItemRepo: delItemRepo,
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

// ========== 收货单服务 ==========

func (s *WarehouseService) ListReceiveOrder(ctx context.Context, tenantID int64) ([]model.ReceiveOrder, int64, error) {
	return s.receiveOrderRepo.List(ctx, tenantID)
}

func (s *WarehouseService) GetReceiveOrderByID(ctx context.Context, id string) (*model.ReceiveOrder, error) {
	var receiveID uint
	_, err := fmt.Sscanf(id, "%d", &receiveID)
	if err != nil {
		return nil, err
	}
	return s.receiveOrderRepo.GetByID(ctx, receiveID)
}

func (s *WarehouseService) GetReceiveOrderWithItems(ctx context.Context, id string) (*ReceiveOrderWithItems, error) {
	receiveID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	order, err := s.receiveOrderRepo.GetByID(ctx, uint(receiveID))
	if err != nil {
		return nil, err
	}
	items, err := s.receiveOrderItemRepo.ListByReceiveID(ctx, uint(receiveID))
	if err != nil {
		return nil, err
	}
	return &ReceiveOrderWithItems{Order: order, Items: items}, nil
}

func (s *WarehouseService) CreateReceiveOrder(ctx context.Context, req *ReceiveOrderCreateReq) error {
	order := &model.ReceiveOrder{
		TenantID:     req.TenantID,
		ReceiveNo:    req.ReceiveNo,
		SupplierID:   req.SupplierID,
		SupplierName: req.SupplierName,
		WarehouseID:  req.WarehouseID,
		ReceiveDate:  req.ReceiveDate,
		Status:       1, // 待收货
		Remark:       req.Remark,
	}

	if err := s.receiveOrderRepo.Create(ctx, order); err != nil {
		return err
	}

	// 创建明细
	for _, item := range req.Items {
		orderItem := &model.ReceiveOrderItem{
			ReceiveID:   order.ID,
			MaterialID:  item.MaterialID,
			MaterialCode: item.MaterialCode,
			MaterialName: item.MaterialName,
			Quantity:    item.Quantity,
			ReceivedQty: 0,
			Unit:       item.Unit,
			BatchNo:    item.BatchNo,
		}
		if err := s.receiveOrderItemRepo.Create(ctx, orderItem); err != nil {
			return err
		}
	}
	return nil
}

func (s *WarehouseService) UpdateReceiveOrder(ctx context.Context, id string, req *ReceiveOrderCreateReq) error {
	receiveID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	updates := map[string]interface{}{
		"supplier_id":    req.SupplierID,
		"supplier_name":  req.SupplierName,
		"warehouse_id":   req.WarehouseID,
		"receive_date":   req.ReceiveDate,
	}
	if req.Remark != nil {
		updates["remark"] = req.Remark
	}

	if err := s.receiveOrderRepo.Update(ctx, uint(receiveID), updates); err != nil {
		return err
	}

	// 更新明细：先删后插
	if err := s.receiveOrderItemRepo.DeleteByReceiveID(ctx, uint(receiveID)); err != nil {
		return err
	}
	for _, item := range req.Items {
		orderItem := &model.ReceiveOrderItem{
			ReceiveID:   int64(receiveID),
			MaterialID:  item.MaterialID,
			MaterialCode: item.MaterialCode,
			MaterialName: item.MaterialName,
			Quantity:    item.Quantity,
			ReceivedQty: 0,
			Unit:       item.Unit,
			BatchNo:    item.BatchNo,
		}
		if err := s.receiveOrderItemRepo.Create(ctx, orderItem); err != nil {
			return err
		}
	}
	return nil
}

func (s *WarehouseService) DeleteReceiveOrder(ctx context.Context, id string) error {
	receiveID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	// 先删明细
	if err := s.receiveOrderItemRepo.DeleteByReceiveID(ctx, uint(receiveID)); err != nil {
		return err
	}
	return s.receiveOrderRepo.Delete(ctx, uint(receiveID))
}

func (s *WarehouseService) ConfirmReceiveOrder(ctx context.Context, id string) error {
	receiveID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	order, err := s.receiveOrderRepo.GetByID(ctx, uint(receiveID))
	if err != nil {
		return err
	}
	if order.Status != 1 {
		return fmt.Errorf("收货单状态不是待收货，无法确认")
	}
	return s.receiveOrderRepo.Update(ctx, uint(receiveID), map[string]interface{}{"status": 3}) // 3-已完成
}

func (s *WarehouseService) GenerateReceiveNo(ctx context.Context, tenantID int64) (string, error) {
	dateStr := time.Now().Format("20060102")
	prefix := "RK" + dateStr

	// 查找今天最大的序号
	var maxNo string
	s.db.WithContext(ctx).Model(&model.ReceiveOrder{}).
		Where("tenant_id = ? AND receive_no LIKE ?", tenantID, prefix+"%").
		Order("receive_no DESC").Limit(1).Pluck("receive_no", &maxNo)

	seq := 1
	if maxNo != "" {
		var num int
		fmt.Sscanf(maxNo[len(prefix):], "%d", &num)
		seq = num + 1
	}
	return fmt.Sprintf("%s%04d", prefix, seq), nil
}

// ========== 发货单服务 ==========

func (s *WarehouseService) ListDeliveryOrder(ctx context.Context, tenantID int64) ([]model.DeliveryOrder, int64, error) {
	return s.deliveryOrderRepo.List(ctx, tenantID)
}

func (s *WarehouseService) GetDeliveryOrderByID(ctx context.Context, id string) (*model.DeliveryOrder, error) {
	var deliveryID uint
	_, err := fmt.Sscanf(id, "%d", &deliveryID)
	if err != nil {
		return nil, err
	}
	return s.deliveryOrderRepo.GetByID(ctx, deliveryID)
}

func (s *WarehouseService) GetDeliveryOrderWithItems(ctx context.Context, id string) (*DeliveryOrderWithItems, error) {
	deliveryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	order, err := s.deliveryOrderRepo.GetByID(ctx, uint(deliveryID))
	if err != nil {
		return nil, err
	}
	items, err := s.deliveryOrderItemRepo.ListByDeliveryID(ctx, uint(deliveryID))
	if err != nil {
		return nil, err
	}
	return &DeliveryOrderWithItems{Order: order, Items: items}, nil
}

func (s *WarehouseService) CreateDeliveryOrder(ctx context.Context, req *DeliveryOrderCreateReq) error {
	order := &model.DeliveryOrder{
		TenantID:     req.TenantID,
		DeliveryNo:   req.DeliveryNo,
		CustomerID:   req.CustomerID,
		CustomerName: req.CustomerName,
		WarehouseID:  req.WarehouseID,
		DeliveryDate: req.DeliveryDate,
		Status:       1, // 待发货
		Remark:       req.Remark,
	}

	if err := s.deliveryOrderRepo.Create(ctx, order); err != nil {
		return err
	}

	// 创建明细
	for _, item := range req.Items {
		orderItem := &model.DeliveryOrderItem{
			DeliveryID:  order.ID,
			MaterialID:  item.MaterialID,
			MaterialCode: item.MaterialCode,
			MaterialName: item.MaterialName,
			Quantity:   item.Quantity,
			ShippedQty: 0,
			Unit:      item.Unit,
			BatchNo:   item.BatchNo,
		}
		if err := s.deliveryOrderItemRepo.Create(ctx, orderItem); err != nil {
			return err
		}
	}
	return nil
}

func (s *WarehouseService) UpdateDeliveryOrder(ctx context.Context, id string, req *DeliveryOrderCreateReq) error {
	deliveryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	updates := map[string]interface{}{
		"customer_id":    req.CustomerID,
		"customer_name":  req.CustomerName,
		"warehouse_id":  req.WarehouseID,
		"delivery_date": req.DeliveryDate,
	}
	if req.Remark != nil {
		updates["remark"] = req.Remark
	}

	if err := s.deliveryOrderRepo.Update(ctx, uint(deliveryID), updates); err != nil {
		return err
	}

	// 更新明细：先删后插
	if err := s.deliveryOrderItemRepo.DeleteByDeliveryID(ctx, uint(deliveryID)); err != nil {
		return err
	}
	for _, item := range req.Items {
		orderItem := &model.DeliveryOrderItem{
			DeliveryID:  int64(deliveryID),
			MaterialID:  item.MaterialID,
			MaterialCode: item.MaterialCode,
			MaterialName: item.MaterialName,
			Quantity:   item.Quantity,
			ShippedQty: 0,
			Unit:      item.Unit,
			BatchNo:   item.BatchNo,
		}
		if err := s.deliveryOrderItemRepo.Create(ctx, orderItem); err != nil {
			return err
		}
	}
	return nil
}

func (s *WarehouseService) DeleteDeliveryOrder(ctx context.Context, id string) error {
	deliveryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	// 先删明细
	if err := s.deliveryOrderItemRepo.DeleteByDeliveryID(ctx, uint(deliveryID)); err != nil {
		return err
	}
	return s.deliveryOrderRepo.Delete(ctx, uint(deliveryID))
}

func (s *WarehouseService) ConfirmDeliveryOrder(ctx context.Context, id string) error {
	deliveryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	order, err := s.deliveryOrderRepo.GetByID(ctx, uint(deliveryID))
	if err != nil {
		return err
	}
	if order.Status != 1 {
		return fmt.Errorf("发货单状态不是待发货，无法确认")
	}
	return s.deliveryOrderRepo.Update(ctx, uint(deliveryID), map[string]interface{}{"status": 3}) // 3-已完成
}

func (s *WarehouseService) GenerateDeliveryNo(ctx context.Context, tenantID int64) (string, error) {
	dateStr := time.Now().Format("20060102")
	prefix := "CK" + dateStr

	var maxNo string
	s.db.WithContext(ctx).Model(&model.DeliveryOrder{}).
		Where("tenant_id = ? AND delivery_no LIKE ?", tenantID, prefix+"%").
		Order("delivery_no DESC").Limit(1).Pluck("delivery_no", &maxNo)

	seq := 1
	if maxNo != "" {
		var num int
		fmt.Sscanf(maxNo[len(prefix):], "%d", &num)
		seq = num + 1
	}
	return fmt.Sprintf("%s%04d", prefix, seq), nil
}

// ========== 辅助结构 ==========

// ReceiveOrderWithItems 收货单及明细
type ReceiveOrderWithItems struct {
	Order *model.ReceiveOrder    `json:"order"`
	Items []model.ReceiveOrderItem `json:"items"`
}

// DeliveryOrderWithItems 发货单及明细
type DeliveryOrderWithItems struct {
	Order *model.DeliveryOrder     `json:"order"`
	Items []model.DeliveryOrderItem `json:"items"`
}

// ReceiveOrderCreateReq 收货单创建请求
type ReceiveOrderCreateReq struct {
	TenantID     int64         `json:"tenant_id"`
	ReceiveNo    string        `json:"receive_no"`
	SupplierID   int64         `json:"supplier_id"`
	SupplierName *string       `json:"supplier_name"`
	WarehouseID  int64         `json:"warehouse_id"`
	ReceiveDate  *time.Time    `json:"receive_date"`
	Remark       *string       `json:"remark"`
	Items        []OrderItemDTO `json:"items"`
}

// DeliveryOrderCreateReq 发货单创建请求
type DeliveryOrderCreateReq struct {
	TenantID     int64         `json:"tenant_id"`
	DeliveryNo   string        `json:"delivery_no"`
	CustomerID   int64         `json:"customer_id"`
	CustomerName *string       `json:"customer_name"`
	WarehouseID  int64         `json:"warehouse_id"`
	DeliveryDate *time.Time    `json:"delivery_date"`
	Remark       *string       `json:"remark"`
	Items        []OrderItemDTO `json:"items"`
}

// OrderItemDTO 订单明细DTO
type OrderItemDTO struct {
	MaterialID   int64   `json:"material_id"`
	MaterialCode string  `json:"material_code"`
	MaterialName string  `json:"material_name"`
	Quantity     float64 `json:"quantity"`
	Unit         string  `json:"unit"`
	BatchNo      *string `json:"batch_no"`
}
