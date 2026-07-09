package model

import (
	"testing"

	"mom-server/internal/pkg/status"
)

// 测试 InterfaceConfig StatusCode / SetStatus(BATCH 3-3 / 2026-07-09)
func TestInterfaceConfigStatusHelpers(t *testing.T) {
	c1 := &InterfaceConfig{Status: "ENABLE"}
	if got := c1.StatusCode(); got != string(status.IntegrationConfigStatusEnable) {
		t.Errorf("legacy ENABLE → %s, want ENABLE", got)
	}

	c2 := &InterfaceConfig{StatusV2: "DISABLE"}
	if got := c2.StatusCode(); got != "DISABLE" {
		t.Errorf("StatusV2 priority = %s, want DISABLE", got)
	}

	c3 := &InterfaceConfig{Status: "ENABLE"}
	c3.SetStatus(string(status.IntegrationConfigStatusDisable))
	if c3.Status != "DISABLE" || c3.StatusV2 != "DISABLE" {
		t.Errorf("SetStatus 双轨不一致: Status=%s StatusV2=%s", c3.Status, c3.StatusV2)
	}

	c4 := &InterfaceConfig{Status: "ENABLE"}
	c4.SetStatus("INVALID")
	if c4.StatusCode() != string(status.IntegrationConfigStatusEnable) {
		t.Errorf("非法码降级失败, got %s", c4.StatusCode())
	}
}