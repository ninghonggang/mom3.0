package service

import (
	"context"
	"sync"
	"time"

	"mom-server/internal/model"
)

// EventHandler is a function that handles domain events
type EventHandler func(ctx context.Context, event *DomainEvent)

// DomainEvent represents a domain event
type DomainEvent struct {
	EventType string
	Timestamp time.Time
	TenantID  int64
	Data      map[string]interface{}
}

// EventBus manages event subscriptions and publishing
type EventBus struct {
	handlers map[string][]EventHandler
	mu       sync.RWMutex
}

var globalEventBus *EventBus

// InitEventBus initializes the global event bus
func InitEventBus() {
	globalEventBus = &EventBus{
		handlers: make(map[string][]EventHandler),
	}
}

// GetEventBus returns the global event bus instance
func GetEventBus() *EventBus {
	return globalEventBus
}

// Subscribe registers a handler for a specific event type
func (eb *EventBus) Subscribe(eventType string, handler EventHandler) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.handlers[eventType] = append(eb.handlers[eventType], handler)
}

// Publish publishes an event to all registered handlers
func (eb *EventBus) Publish(ctx context.Context, event *DomainEvent) {
	eb.mu.RLock()
	handlers := eb.handlers[event.EventType]
	eb.mu.RUnlock()

	for _, h := range handlers {
		go func(handler EventHandler) {
			handler(ctx, event)
		}(h)
	}
}

// PublishSync publishes an event synchronously
func (eb *EventBus) PublishSync(ctx context.Context, event *DomainEvent) {
	eb.mu.RLock()
	handlers := eb.handlers[event.EventType]
	eb.mu.RUnlock()

	for _, h := range handlers {
		h(ctx, event)
	}
}

// NewDomainEvent creates a new domain event
func NewDomainEvent(eventType string, tenantID int64, data map[string]interface{}) *DomainEvent {
	return &DomainEvent{
		EventType: eventType,
		Timestamp: time.Now(),
		TenantID:  tenantID,
		Data:      data,
	}
}

// Predefined event creation helpers

// NewProductionCompleteEvent creates a production complete event
func NewProductionCompleteEvent(tenantID int64, orderID int64, orderNo string, qty float64, workshopID int64) *DomainEvent {
	return NewDomainEvent(model.EventTypeProductionComplete, tenantID, map[string]interface{}{
		"order_id":    orderID,
		"order_no":   orderNo,
		"qty":        qty,
		"workshop_id": workshopID,
	})
}

// NewQualityInspectEvent creates a quality inspection complete event
func NewQualityInspectEvent(tenantID int64, inspectID int64, inspectType string, result string, orderID int64) *DomainEvent {
	return NewDomainEvent(model.EventTypeQualityInspect, tenantID, map[string]interface{}{
		"inspect_id":   inspectID,
		"inspect_type": inspectType,
		"result":       result,
		"order_id":     orderID,
	})
}

// NewStockInEvent creates a stock in event
func NewStockInEvent(tenantID int64, stockInID int64, materialID int64, qty float64, warehouseID int64) *DomainEvent {
	return NewDomainEvent(model.EventTypeStockIn, tenantID, map[string]interface{}{
		"stock_in_id":  stockInID,
		"material_id": materialID,
		"qty":         qty,
		"warehouse_id": warehouseID,
	})
}

// NewStockOutEvent creates a stock out event
func NewStockOutEvent(tenantID int64, stockOutID int64, materialID int64, qty float64, warehouseID int64) *DomainEvent {
	return NewDomainEvent(model.EventTypeStockOut, tenantID, map[string]interface{}{
		"stock_out_id": stockOutID,
		"material_id": materialID,
		"qty":         qty,
		"warehouse_id": warehouseID,
	})
}

// NewPurchaseAwardEvent creates a purchase award event
func NewPurchaseAwardEvent(tenantID int64, awardID int64, supplierID int64, orderNo string, amount float64) *DomainEvent {
	return NewDomainEvent(model.EventTypePurchaseAward, tenantID, map[string]interface{}{
		"award_id":    awardID,
		"supplier_id": supplierID,
		"order_no":   orderNo,
		"amount":      amount,
	})
}

// NewPurchaseReceiveEvent creates a purchase receive event
func NewPurchaseReceiveEvent(tenantID int64, poID int64, poNo string, supplierID int64, supplierName string, itemID uint, materialID int64, materialCode string, materialName string, qty float64, batchNo string) *DomainEvent {
	return NewDomainEvent(model.EventTypePurchaseReceive, tenantID, map[string]interface{}{
		"po_id":          poID,
		"po_no":          poNo,
		"supplier_id":    supplierID,
		"supplier_name":  supplierName,
		"item_id":        itemID,
		"material_id":    materialID,
		"material_code":  materialCode,
		"material_name":  materialName,
		"qty":           qty,
		"batch_no":      batchNo,
	})
}

// NewSalesShipEvent creates a sales ship event
func NewSalesShipEvent(tenantID int64, shipID int64, customerID int64, orderNo string, qty float64) *DomainEvent {
	return NewDomainEvent(model.EventTypeSalesShip, tenantID, map[string]interface{}{
		"ship_id":     shipID,
		"customer_id": customerID,
		"order_no":   orderNo,
		"qty":         qty,
	})
}
