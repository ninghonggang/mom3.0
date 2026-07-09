package model

import (
	"testing"

	"mom-server/internal/pkg/status"
)

// 测试 MesWorkSchedulingDetail StatusCode / SetStatus(BATCH 3-5 / 2026-07-09)
func TestMesWorkSchedulingDetailStatusHelpers(t *testing.T) {
	d1 := &MesWorkSchedulingDetail{Status: "IN_PROGRESS"}
	if got := d1.StatusCode(); got != string(status.MesWorkSchedulingDetailStatusInProgress) {
		t.Errorf("legacy IN_PROGRESS → %s, want IN_PROGRESS", got)
	}

	d2 := &MesWorkSchedulingDetail{StatusV2: "PAUSED"}
	if got := d2.StatusCode(); got != "PAUSED" {
		t.Errorf("StatusV2 priority = %s, want PAUSED", got)
	}

	d3 := &MesWorkSchedulingDetail{Status: "PENDING"}
	d3.SetStatus(string(status.MesWorkSchedulingDetailStatusCompleted))
	if d3.Status != "COMPLETED" || d3.StatusV2 != "COMPLETED" {
		t.Errorf("SetStatus 双轨不一致: Status=%s StatusV2=%s", d3.Status, d3.StatusV2)
	}

	d4 := &MesWorkSchedulingDetail{Status: "PENDING"}
	d4.SetStatus("INVALID")
	if d4.StatusCode() != string(status.MesWorkSchedulingDetailStatusPending) {
		t.Errorf("非法码降级失败, got %s", d4.StatusCode())
	}
}