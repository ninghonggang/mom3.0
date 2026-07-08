// Package status 提供 MOM 3.0 状态字段统一字典常量与 helper。
//
// 设计原则:
//   - 单一事实源 (SSOT): 与 mdm_status_dict 表的 code 字段对齐,DB 与代码同步
//   - 多 domain: production / mes / aps / quality / wms / scp / eam
//   - 双轨制: 不破坏现有 bigint legacy 字段, 新代码用 V2 helper
//
// 配套迁移: migrations/0051_status_unification.sql
// 设计文档: docs/MOM3.0_状态字段统一方案.md
package status

// Code 状态码(VARCHAR(30)),与 mdm_status_dict.code 对齐
type Code string

// IsValid 判断状态码是否在该 entity 允许的字典内
func (c Code) IsValid(allowed []Code) bool {
	for _, a := range allowed {
		if c == a {
			return true
		}
	}
	return false
}

// String 实现 fmt.Stringer
func (c Code) String() string { return string(c) }

// ========== production.production_order ==========
// 字典来源: mdm_status_dict WHERE entity='production_order'
// legacy_int: 1=DRAFT, 2=RELEASED, 3=IN_PROGRESS, 5=CLOSED, 4=COMPLETED
const (
	ProductionOrderDraft       Code = "DRAFT"
	ProductionOrderReleased    Code = "RELEASED"
	ProductionOrderInProgress  Code = "IN_PROGRESS"
	ProductionOrderHold        Code = "HOLD"
	ProductionOrderCompleted   Code = "COMPLETED"
	ProductionOrderClosed      Code = "CLOSED"
	ProductionOrderCancelled   Code = "CANCELLED"
)

// ProductionOrderAll production_order 完整允许状态集(7 个)
var ProductionOrderAll = []Code{
	ProductionOrderDraft, ProductionOrderReleased, ProductionOrderInProgress,
	ProductionOrderHold, ProductionOrderCompleted, ProductionOrderClosed,
	ProductionOrderCancelled,
}

// ProductionOrderFromLegacyInt bigint legacy → V2 字典码
// ⚠️ legacy int 是 short form(5 状态),与 V2 6 状态字典不是一一对应:
//   - 0051 迁移映射: 1=DRAFT, 2=IN_PROGRESS, 3=COMPLETED, 4=CANCELLED, 5=CLOSED
//   - RELEASED / HOLD 在 legacy 中没有专属位(2 已被 IN_PROGRESS 占用)
//   - 反向解析遇到 status=2 一律返回 IN_PROGRESS(已部署数据兼容)
func ProductionOrderFromLegacyInt(i int) Code {
	switch i {
	case 1:
		return ProductionOrderDraft
	case 2:
		return ProductionOrderInProgress
	case 3:
		return ProductionOrderCompleted
	case 4:
		return ProductionOrderCancelled
	case 5:
		return ProductionOrderClosed
	case 6:
		return ProductionOrderCancelled
	default:
		return ProductionOrderDraft
	}
}

// ProductionOrderToLegacyInt V2 字典码 → bigint legacy(双轨回写)
// ⚠️ 多对一映射: RELEASED/HOLD/COMPLETED 都回写为 3,存在信息丢失
// 仅用于老接口兼容,新代码直接读 status_v2
func ProductionOrderToLegacyInt(c Code) int {
	switch c {
	case ProductionOrderDraft:
		return 1
	case ProductionOrderReleased:
		return 3 // ⚠️ legacy 无专属位,借 IN_PROGRESS 位
	case ProductionOrderInProgress:
		return 2 // 跟 0051 迁移反向一致
	case ProductionOrderHold:
		return 3 // ⚠️ legacy 无专属位
	case ProductionOrderCompleted:
		return 3 // ⚠️ 跟 0051 一致(legacy short form 不区分 IN_PROGRESS/COMPLETED)
	case ProductionOrderClosed:
		return 5
	case ProductionOrderCancelled:
		return 4 // 跟 0051 一致
	default:
		return 1
	}
}

// ========== mes.mes_process ==========
// 字典来源: mdm_status_dict WHERE entity='mes_process'
// 原表 status 字段是 varchar,枚举值 DRAFT/ACTIVE/EXPIRED
// V2 规范改为 DRAFT/ACTIVE/OBSOLETE(语义更清晰),旧数据 EXPIRED → OBSOLETE 兼容
const (
	MesProcessDraft    Code = "DRAFT"
	MesProcessActive   Code = "ACTIVE"
	MesProcessObsolete Code = "OBSOLETE"
)

// MesProcessAll mes_process 完整允许状态集
var MesProcessAll = []Code{
	MesProcessDraft, MesProcessActive, MesProcessObsolete,
}

// MesProcessFromLegacyVarchar varchar legacy → V2(兼容 EXPIRED 别名)
func MesProcessFromLegacyVarchar(s string) Code {
	switch s {
	case "DRAFT":
		return MesProcessDraft
	case "ACTIVE":
		return MesProcessActive
	case "EXPIRED", "OBSOLETE":
		return MesProcessObsolete
	default:
		return MesProcessDraft
	}
}

// ========== mdm.bom ==========
// 字典来源: mdm_status_dict WHERE entity='bom'
// 与 mes_process 同模式:DRAFT 草稿 / ACTIVE 生效 / OBSOLETE 失效
// 原表 status varchar,枚举值 DRAFT/ACTIVE/EXPIRED,V2 兼容 EXPIRED→OBSOLETE
const (
	MDMBomDraft    Code = "DRAFT"
	MDMBomActive   Code = "ACTIVE"
	MDMBomObsolete Code = "OBSOLETE"
)

// MDMBomAll bom 完整允许状态集
var MDMBomAll = []Code{
	MDMBomDraft, MDMBomActive, MDMBomObsolete,
}

// MDMBomFromLegacyVarchar varchar legacy → V2(兼容 EXPIRED 别名,与 mes_process 同模式)
func MDMBomFromLegacyVarchar(s string) Code {
	switch s {
	case "DRAFT":
		return MDMBomDraft
	case "ACTIVE":
		return MDMBomActive
	case "EXPIRED", "OBSOLETE":
		return MDMBomObsolete
	default:
		return MDMBomDraft
	}
}

// ========== wms.production_issue(BATCH 2026-07-08,扩展 P2) ==========
// 字典来源:wms_production_issue.status 现有枚举
// 状态机:PENDING → APPROVED → PICKING → PICKED → ISSUED (+ CANCELLED 终态)
const (
	ProductionIssuePending   Code = "PENDING"
	ProductionIssueApproved  Code = "APPROVED"
	ProductionIssuePicking   Code = "PICKING"
	ProductionIssuePicked    Code = "PICKED"
	ProductionIssueIssued    Code = "ISSUED"
	ProductionIssueCancelled Code = "CANCELLED"
)

// ProductionIssueAll production_issue 完整允许状态集
var ProductionIssueAll = []Code{
	ProductionIssuePending, ProductionIssueApproved, ProductionIssuePicking,
	ProductionIssuePicked, ProductionIssueIssued, ProductionIssueCancelled,
}

// ProductionIssueFromLegacyVarchar varchar legacy → V2(原表本身即 varchar,直接透传)
func ProductionIssueFromLegacyVarchar(s string) Code {
	switch s {
	case "PENDING":
		return ProductionIssuePending
	case "APPROVED":
		return ProductionIssueApproved
	case "PICKING":
		return ProductionIssuePicking
	case "PICKED":
		return ProductionIssuePicked
	case "ISSUED":
		return ProductionIssueIssued
	case "CANCELLED":
		return ProductionIssueCancelled
	default:
		return ProductionIssuePending
	}
}

// ========== purchase.purchase_return(BATCH 2026-07-08 P2 扩展) ==========
// 字典来源:purchase_return.status 现有枚举
// 状态机:PENDING → APPROVED → RETURNING → RETURNED
const (
	PurchaseReturnPending   Code = "PENDING"
	PurchaseReturnApproved  Code = "APPROVED"
	PurchaseReturnReturning Code = "RETURNING"
	PurchaseReturnReturned  Code = "RETURNED"
)

var PurchaseReturnAll = []Code{
	PurchaseReturnPending, PurchaseReturnApproved, PurchaseReturnReturning, PurchaseReturnReturned,
}

func PurchaseReturnFromLegacyVarchar(s string) Code {
	switch s {
	case "PENDING":
		return PurchaseReturnPending
	case "APPROVED":
		return PurchaseReturnApproved
	case "RETURNING":
		return PurchaseReturnReturning
	case "RETURNED":
		return PurchaseReturnReturned
	default:
		return PurchaseReturnPending
	}
}

// ========== production.production_return(BATCH 2026-07-08 P2 扩展) ==========
// 与 purchase_return 同字典(返回生产物料到仓库 vs 采购返回)
const (
	ProductionReturnPending   Code = "PENDING"
	ProductionReturnApproved  Code = "APPROVED"
	ProductionReturnReturning Code = "RETURNING"
	ProductionReturnReturned  Code = "RETURNED"
)

var ProductionReturnAll = []Code{
	ProductionReturnPending, ProductionReturnApproved, ProductionReturnReturning, ProductionReturnReturned,
}

func ProductionReturnFromLegacyVarchar(s string) Code {
	switch s {
	case "PENDING":
		return ProductionReturnPending
	case "APPROVED":
		return ProductionReturnApproved
	case "RETURNING":
		return ProductionReturnReturning
	case "RETURNED":
		return ProductionReturnReturned
	default:
		return ProductionReturnPending
	}
}

// ========== eam.andon_call(BATCH 2026-07-08 P2 扩展) ==========
// 字典来源:andon.status 现有枚举
// 状态机:CALLING → RESPONDED → HANDLING → RESOLVED → CLOSED
const (
	AndonCallCalling   Code = "CALLING"
	AndonCallResponded Code = "RESPONDED"
	AndonCallHandling  Code = "HANDLING"
	AndonCallResolved  Code = "RESOLVED"
	AndonCallClosed    Code = "CLOSED"
)

var AndonCallAll = []Code{
	AndonCallCalling, AndonCallResponded, AndonCallHandling, AndonCallResolved, AndonCallClosed,
}

func AndonCallFromLegacyVarchar(s string) Code {
	switch s {
	case "CALLING":
		return AndonCallCalling
	case "RESPONDED":
		return AndonCallResponded
	case "HANDLING":
		return AndonCallHandling
	case "RESOLVED":
		return AndonCallResolved
	case "CLOSED":
		return AndonCallClosed
	default:
		return AndonCallCalling
	}
}

// ========== eam.downtime(BATCH 2026-07-08 P2 扩展) ==========
// 字典来源:downtime.status 现有枚举
// 状态机:OPEN → INPROGRESS → CLOSED
const (
	DowntimeOpen       Code = "OPEN"
	DowntimeInProgress Code = "INPROGRESS"
	DowntimeClosed     Code = "CLOSED"
)

var DowntimeAll = []Code{
	DowntimeOpen, DowntimeInProgress, DowntimeClosed,
}

func DowntimeFromLegacyVarchar(s string) Code {
	switch s {
	case "OPEN":
		return DowntimeOpen
	case "INPROGRESS":
		return DowntimeInProgress
	case "CLOSED":
		return DowntimeClosed
	default:
		return DowntimeOpen
	}
}

// ========== production.mobile_job_report / production_report / dispatch ==========
const (
	ReportSubmitted Code = "SUBMITTED"
	ReportConfirmed Code = "CONFIRMED"
	ReportAudited   Code = "AUDITED"
	ReportRejected  Code = "REJECTED"

	DispatchPending     Code = "PENDING"
	DispatchInProgress  Code = "IN_PROGRESS"
	DispatchCompleted   Code = "COMPLETED"
)