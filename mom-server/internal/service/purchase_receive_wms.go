package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"

	"gorm.io/gorm"
)

// PurchaseReceiveWMSHandler 采购收货WMS联动处理器
type PurchaseReceiveWMSHandler struct {
	db                   *gorm.DB
	receiveOrderRepo     *repository.ReceiveOrderRepository
	receiveOrderItemRepo *repository.ReceiveOrderItemRepository
	warehouseRepo        *repository.WarehouseRepository
}

// NewPurchaseReceiveWMSHandler 创建采购收货WMS联动处理器
func NewPurchaseReceiveWMSHandler(
	db *gorm.DB,
	receiveOrderRepo *repository.ReceiveOrderRepository,
	receiveOrderItemRepo *repository.ReceiveOrderItemRepository,
	warehouseRepo *repository.WarehouseRepository,
) *PurchaseReceiveWMSHandler {
	return &PurchaseReceiveWMSHandler{
		db:                   db,
		receiveOrderRepo:     receiveOrderRepo,
		receiveOrderItemRepo: receiveOrderItemRepo,
		warehouseRepo:        warehouseRepo,
	}
}

// HandlePurchaseReceive 处理采购收货事件，创建WMS收货单
func (h *PurchaseReceiveWMSHandler) HandlePurchaseReceive(ctx context.Context, event *DomainEvent) {
	data := event.Data

	poIDStr, ok := data["po_id"].(float64)
	if !ok {
		fmt.Printf("[WMS] Invalid po_id type in event\n")
		return
	}
	poID := int64(poIDStr)
	_ = poID // suppress unused warning

	poNo, ok := data["po_no"].(string)
	if !ok {
		poNo = ""
	}

	supplierID, ok := data["supplier_id"].(float64)
	if !ok {
		supplierID = 0
	}

	var supplierName *string
	if sn, ok := data["supplier_name"].(string); ok {
		supplierName = &sn
	}

	itemID, ok := data["item_id"].(float64)
	if !ok {
		itemID = 0
	}

	materialID, ok := data["material_id"].(float64)
	if !ok {
		materialID = 0
	}

	materialCode, _ := data["material_code"].(string)
	materialName, _ := data["material_name"].(string)

	qty, ok := data["qty"].(float64)
	if !ok {
		qty = 0
	}

	batchNo, _ := data["batch_no"].(string)

	// 生成收货单号
	receiveNo, err := h.generateReceiveNo(ctx, event.TenantID)
	if err != nil {
		fmt.Printf("[WMS] Failed to generate receive no: %v\n", err)
		return
	}

	// 获取默认仓库
	warehouseID := h.getDefaultWarehouse(ctx)
	if warehouseID == 0 {
		fmt.Printf("[WMS] No default warehouse found\n")
		return
	}

	// 创建收货单
	remark := fmt.Sprintf("来源采购订单: %s", poNo)
	now := time.Now()

	receiveOrder := &model.ReceiveOrder{
		TenantID:     event.TenantID,
		ReceiveNo:    receiveNo,
		SupplierID:   int64(supplierID),
		SupplierName: supplierName,
		WarehouseID:  warehouseID,
		ReceiveDate:  &now,
		Status:       1, // 待收货
		Remark:       &remark,
	}

	if err := h.receiveOrderRepo.Create(ctx, receiveOrder); err != nil {
		fmt.Printf("[WMS] Failed to create receive order: %v\n", err)
		return
	}

	// 创建收货单明细
	receiveOrderItem := &model.ReceiveOrderItem{
		ReceiveID:    receiveOrder.ID,
		MaterialID:   int64(materialID),
		MaterialCode: materialCode,
		MaterialName: materialName,
		Quantity:     qty,
		ReceivedQty:  0,
		Unit:         "PCS",
	}
	if batchNo != "" {
		receiveOrderItem.BatchNo = &batchNo
	}

	if err := h.receiveOrderItemRepo.Create(ctx, receiveOrderItem); err != nil {
		fmt.Printf("[WMS] Failed to create receive order item: %v\n", err)
		return
	}

	fmt.Printf("[WMS] Created receive order %s for PO %s, item %d, qty %.3f\n",
		receiveNo, poNo, uint(itemID), qty)
}

func (h *PurchaseReceiveWMSHandler) generateReceiveNo(ctx context.Context, tenantID int64) (string, error) {
	dateStr := time.Now().Format("20060102")
	prefix := fmt.Sprintf("RK%s", dateStr[2:])

	var count int64
	h.db.WithContext(ctx).Model(&model.ReceiveOrder{}).
		Where("tenant_id = ? AND receive_no LIKE ?", tenantID, prefix+"%").
		Count(&count)

	return fmt.Sprintf("%s%04d", prefix, count+1), nil
}

func (h *PurchaseReceiveWMSHandler) getDefaultWarehouse(ctx context.Context) int64 {
	var warehouse model.Warehouse
	err := h.db.WithContext(ctx).Where("status = 'ACTIVE'").First(&warehouse).Error
	if err != nil {
		return 0
	}
	return warehouse.ID
}

// Subscribe 订阅采购收货事件
func (h *PurchaseReceiveWMSHandler) Subscribe() {
	eventBus := GetEventBus()
	eventBus.Subscribe(model.EventTypePurchaseReceive, func(ctx context.Context, event *DomainEvent) {
		h.HandlePurchaseReceive(ctx, event)
	})
}
