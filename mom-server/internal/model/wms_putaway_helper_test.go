package model

import (
	"testing"

	"mom-server/internal/pkg/status"
)

// 测试 WMSPutawayJob / WMSPutawayRecord StatusCode / SetStatus(BATCH 3-2 / 2026-07-09)
func TestWMSPutawayStatusHelpers(t *testing.T) {
	// WMSPutawayJob
	j1 := &WMSPutawayJob{Status: "PENDING"}
	if got := j1.StatusCode(); got != string(status.WMSPutawayJobPending) {
		t.Errorf("Job legacy PENDING → %s, want PENDING", got)
	}

	j2 := &WMSPutawayJob{StatusV2: "PUTAWAYING"}
	if got := j2.StatusCode(); got != "PUTAWAYING" {
		t.Errorf("Job StatusV2 priority = %s, want PUTAWAYING", got)
	}

	j3 := &WMSPutawayJob{Status: "PENDING"}
	j3.SetStatus(string(status.WMSPutawayJobCompleted))
	if j3.Status != "COMPLETED" || j3.StatusV2 != "COMPLETED" {
		t.Errorf("Job SetStatus 双轨不一致: Status=%s StatusV2=%s", j3.Status, j3.StatusV2)
	}

	j4 := &WMSPutawayJob{Status: "PENDING"}
	j4.SetStatus("INVALID")
	if j4.StatusCode() != string(status.WMSPutawayJobPending) {
		t.Errorf("Job 非法码降级失败, got %s", j4.StatusCode())
	}

	// WMSPutawayRecord
	r1 := &WMSPutawayRecord{Status: "PUTAWAYING"}
	if got := r1.StatusCode(); got != string(status.WMSPutawayRecordPutawaying) {
		t.Errorf("Record legacy PUTAWAYING → %s, want PUTAWAYING", got)
	}

	r2 := &WMSPutawayRecord{Status: "PENDING"}
	r2.SetStatus(string(status.WMSPutawayRecordCompleted))
	if r2.Status != "COMPLETED" || r2.StatusV2 != "COMPLETED" {
		t.Errorf("Record SetStatus 双轨不一致: Status=%s StatusV2=%s", r2.Status, r2.StatusV2)
	}
}