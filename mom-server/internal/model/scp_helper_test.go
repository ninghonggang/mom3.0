package model

import (
	"testing"

	"mom-server/internal/pkg/status"
)

// 测试 StatusCode() / SetStatus() helper(BATCH 3-1 / 2026-07-09)
func TestPurchaseOrderItemStatusHelpers(t *testing.T) {
	// 1. legacy Status 优先回退到字典码
	p1 := &PurchaseOrderItem{Status: "COMPLETED"}
	if got := p1.StatusCode(); got != string(status.PurchaseOrderItemCompleted) {
		t.Errorf("legacy StatusCode = %s, want COMPLETED", got)
	}

	// 2. StatusV2 优先 Status
	p2 := &PurchaseOrderItem{Status: "PENDING", StatusV2: "PARTIAL"}
	if got := p2.StatusCode(); got != "PARTIAL" {
		t.Errorf("StatusV2 priority = %s, want PARTIAL", got)
	}

	// 3. SetStatus 双轨写回
	p3 := &PurchaseOrderItem{Status: "PENDING"}
	p3.SetStatus(string(status.PurchaseOrderItemCompleted))
	if p3.Status != "COMPLETED" || p3.StatusV2 != "COMPLETED" {
		t.Errorf("SetStatus 双轨不一致: Status=%s, StatusV2=%s", p3.Status, p3.StatusV2)
	}

	// 4. SetStatus 非法码降级
	p4 := &PurchaseOrderItem{Status: "PENDING"}
	p4.SetStatus("INVALID_CODE")
	if p4.StatusCode() != string(status.PurchaseOrderItemPending) {
		t.Errorf("非法码降级失败, got %s", p4.StatusCode())
	}
}
