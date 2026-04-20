package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"

	"gorm.io/gorm"
)

// SupplierKPICalculator 供应商绩效计算器
type SupplierKPICalculator struct {
	db              *gorm.DB
	deliveryRepo    *repository.SupplierDeliveryRecordRepository
	kpiRepo         *repository.SupplierKPIRepository
	supplierRepo    *repository.SupplierRepository
}

// NewSupplierKPICalculator 创建供应商绩效计算器
func NewSupplierKPICalculator(
	db *gorm.DB,
	deliveryRepo *repository.SupplierDeliveryRecordRepository,
	kpiRepo *repository.SupplierKPIRepository,
	supplierRepo *repository.SupplierRepository,
) *SupplierKPICalculator {
	return &SupplierKPICalculator{
		db:           db,
		deliveryRepo: deliveryRepo,
		kpiRepo:      kpiRepo,
		supplierRepo: supplierRepo,
	}
}

// HandlePurchaseReceive 处理采购收货事件，记录交货记录并更新供应商绩效
func (h *SupplierKPICalculator) HandlePurchaseReceive(ctx context.Context, event *DomainEvent) {
	data := event.Data

	supplierID, ok := data["supplier_id"].(float64)
	if !ok || supplierID == 0 {
		return
	}

	poNo, _ := data["po_no"].(string)
	itemID, _ := data["item_id"].(float64)

	// 创建交货记录
	poNoPtr := &poNo
	var itemIDPtr *int64
	if itemID > 0 {
		id := int64(itemID)
		itemIDPtr = &id
	}
	status := "ON_TIME"
	tenantID := event.TenantID
	now := time.Now()

	record := &model.SupplierDeliveryRecord{
		SupplierID:         int64(supplierID),
		PONo:              poNoPtr,
		POLineID:          itemIDPtr,
		ActualDeliveryDate: &now,
		DeliveryStatus:     &status,
		TenantID:          tenantID,
	}

	if err := h.deliveryRepo.Create(ctx, record); err != nil {
		fmt.Printf("[KPI] Failed to create delivery record: %v\n", err)
		return
	}

	// 更新供应商绩效评分
	h.calculateSupplierKPI(ctx, int64(supplierID), tenantID)

	fmt.Printf("[KPI] Recorded delivery for supplier %d, PO %s\n", int64(supplierID), poNo)
}

func (h *SupplierKPICalculator) calculateSupplierKPI(ctx context.Context, supplierID int64, tenantID int64) {
	// 获取当月
	now := time.Now()
	currentMonth := now.Format("2006-01")

	// 查询当月交货记录
	records, err := h.deliveryRepo.ListBySupplier(ctx, supplierID, 100)
	if err != nil {
		fmt.Printf("[KPI] Failed to list delivery records: %v\n", err)
		return
	}

	// 统计
	totalOrders := len(records)
	onTimeCount := 0
	delayedCount := 0
	totalDelayDays := 0

	for _, record := range records {
		if record.DeliveryStatus != nil && *record.DeliveryStatus == "ON_TIME" {
			onTimeCount++
		} else if record.DeliveryStatus != nil && *record.DeliveryStatus == "DELAYED" {
			delayedCount++
			totalDelayDays += record.DelayDays
		}
	}

	// 计算准时交货率
	var onTimeRate float64
	if totalOrders > 0 {
		onTimeRate = float64(onTimeCount) / float64(totalOrders) * 100
	}

	// 查询现有KPI
	existingKPI, err := h.kpiRepo.GetBySupplierMonthly(ctx, supplierID, currentMonth)
	if err != nil {
		// KPI不存在，创建新的
		newKPI := &model.SupplierKPI{
			SupplierID:           supplierID,
			EvaluationMonth:      currentMonth,
			EvaluationDate:       now,
			TotalDeliveryOrders: totalOrders,
			OnTimeDeliveryRate:  &onTimeRate,
		}
		h.kpiRepo.Create(ctx, newKPI)
		fmt.Printf("[KPI] Created new KPI for supplier %d, month %s\n", supplierID, currentMonth)
	} else {
		// 更新现有KPI
		h.kpiRepo.Update(ctx, existingKPI.ID, map[string]interface{}{
			"total_delivery_orders": totalOrders,
			"on_time_delivery_rate":  onTimeRate,
		})
		fmt.Printf("[KPI] Updated KPI for supplier %d, month %s, on-time rate: %.2f%%\n",
			supplierID, currentMonth, onTimeRate)
	}
}

// Subscribe 订阅采购收货事件
func (h *SupplierKPICalculator) Subscribe() {
	eventBus := GetEventBus()
	eventBus.Subscribe(model.EventTypePurchaseReceive, func(ctx context.Context, event *DomainEvent) {
		h.HandlePurchaseReceive(ctx, event)
	})
}
