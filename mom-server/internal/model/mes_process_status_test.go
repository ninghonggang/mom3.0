package model

import "testing"

func TestMesProcessStatusCode(t *testing.T) {
	// 优先返回 status_v2
	p := &MesProcess{Status: "DRAFT", StatusV2: "ACTIVE"}
	if got := p.StatusCode(); got != "ACTIVE" {
		t.Errorf("StatusCode() = %s, want ACTIVE", got)
	}

	// fallback 到 legacy varchar
	p2 := &MesProcess{Status: "DRAFT", StatusV2: ""}
	if got := p2.StatusCode(); got != "DRAFT" {
		t.Errorf("StatusCode() = %s, want DRAFT", got)
	}

	// legacy EXPIRED → OBSOLETE 兼容
	p3 := &MesProcess{Status: "EXPIRED", StatusV2: ""}
	if got := p3.StatusCode(); got != "OBSOLETE" {
		t.Errorf("StatusCode() = %s, want OBSOLETE (EXPIRED 别名兼容)", got)
	}
}

func TestMesProcessSetStatus(t *testing.T) {
	// 合法码双轨写
	p := &MesProcess{}
	p.SetStatus("ACTIVE")
	if p.Status != "ACTIVE" || p.StatusV2 != "ACTIVE" {
		t.Errorf("SetStatus(ACTIVE) = status=%s, status_v2=%s, want both ACTIVE", p.Status, p.StatusV2)
	}

	// 非法码 fallback DRAFT
	p2 := &MesProcess{}
	p2.SetStatus("GARBAGE")
	if p2.Status != "DRAFT" || p2.StatusV2 != "DRAFT" {
		t.Errorf("SetStatus(GARBAGE) = status=%s, status_v2=%s, want both DRAFT (fallback)", p2.Status, p2.StatusV2)
	}
}