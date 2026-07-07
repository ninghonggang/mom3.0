package status

import "testing"

func TestCodeIsValid(t *testing.T) {
	allowed := []Code{MesProcessDraft, MesProcessActive, MesProcessObsolete}
	cases := []struct {
		in   Code
		want bool
	}{
		{MesProcessDraft, true},
		{MesProcessActive, true},
		{MesProcessObsolete, true},
		{Code("UNKNOWN"), false},
		{Code(""), false},
		{Code("draft"), false}, // 大小写敏感
	}
	for _, c := range cases {
		if got := c.in.IsValid(allowed); got != c.want {
			t.Errorf("IsValid(%s) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestMesProcessFromLegacyVarchar(t *testing.T) {
	cases := []struct {
		in   string
		want Code
	}{
		{"DRAFT", MesProcessDraft},
		{"ACTIVE", MesProcessActive},
		{"EXPIRED", MesProcessObsolete}, // 别名兼容
		{"OBSOLETE", MesProcessObsolete},
		{"unknown", MesProcessDraft}, // 未知值 fallback
		{"", MesProcessDraft},
	}
	for _, c := range cases {
		if got := MesProcessFromLegacyVarchar(c.in); got != c.want {
			t.Errorf("MesProcessFromLegacyVarchar(%q) = %s, want %s", c.in, got, c.want)
		}
	}
}

func TestProductionOrderRoundTrip(t *testing.T) {
	// legacy short form 完整往返 4 个无歧义状态
	for _, c := range []Code{
		ProductionOrderDraft,
		ProductionOrderInProgress,
		ProductionOrderClosed,
		ProductionOrderCancelled,
	} {
		legacy := ProductionOrderToLegacyInt(c)
		got := ProductionOrderFromLegacyInt(legacy)
		if got != c {
			t.Errorf("round-trip %s -> %d -> %s broken", c, legacy, got)
		}
	}

	// 多对一映射验证: legacy=3 被 COMPLETED 占用(0051 定),RELEASED/HOLD/COMPLETED 都映射到 3
	// round-trip 会折叠到 COMPLETED(语义损失,但 round-trip 闭环)
	for _, c := range []Code{
		ProductionOrderReleased,
		ProductionOrderHold,
		ProductionOrderCompleted,
	} {
		legacy := ProductionOrderToLegacyInt(c)
		if legacy != 3 {
			t.Errorf("%s → legacy 应为 3,实际 %d", c, legacy)
		}
		got := ProductionOrderFromLegacyInt(legacy)
		if got != ProductionOrderCompleted {
			t.Errorf("legacy=3 → 应折叠到 COMPLETED,实际 %s", got)
		}
	}
}

func TestProductionOrderAllCount(t *testing.T) {
	// 防御: 字典数量变化时提醒
	if len(ProductionOrderAll) != 7 {
		t.Errorf("ProductionOrderAll 期望 7 个状态, 实际 %d", len(ProductionOrderAll))
	}
	if len(MesProcessAll) != 3 {
		t.Errorf("MesProcessAll 期望 3 个状态, 实际 %d", len(MesProcessAll))
	}
}