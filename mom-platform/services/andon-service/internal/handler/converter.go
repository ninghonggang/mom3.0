package handler

import (
	common "github.com/ninghonggang/mom-platform/gen/common"
	andon "github.com/ninghonggang/mom-platform/gen/andon"

	"mom-platform/services/andon-service/internal/model"
)

// --- model -> proto ---

func modelToProtoAndonCall(m *model.AndonCall) *andon.AndonCall {
	var acknowledgedAt, resolvedAt int64
	var responseSeconds int32
	if m.AcknowledgedAt != nil {
		acknowledgedAt = m.AcknowledgedAt.Unix()
	}
	if m.ResolvedAt != nil {
		resolvedAt = m.ResolvedAt.Unix()
	}
	if m.ResponseSeconds != nil {
		responseSeconds = int32(*m.ResponseSeconds)
	}

	return &andon.AndonCall{
		Id:              int64(m.ID),
		AndonNo:         m.AndonNo,
		WorkstationId:   parseInt64FromString(m.WorkstationID),
		ReporterId:      parseInt64FromString(m.ReporterID),
		AndonType:       stringToAndonType(m.AndonType),
		Description:     m.Description,
		Status:          stringToAndonCallStatus(m.Status),
		TriggeredAt:     m.TriggeredAt.Unix(),
		AcknowledgedAt:  acknowledgedAt,
		ResolvedAt:      resolvedAt,
		ResponseSeconds: responseSeconds,
	}
}

func modelToProtoAlertConfig(m *model.AlertConfig) *andon.AlertConfig {
	return &andon.AlertConfig{
		Id:               int64(m.ID),
		ConfigCode:       m.ConfigCode,
		ConfigName:       m.ConfigName,
		TriggerType:      stringToTriggerType(m.TriggerType),
		Severity:         stringToAlertSeverity(m.Severity),
		TriggerCondition: m.TriggerCondition,
		NotifyChannels:   m.NotifyChannels,
		Status:           m.Status,
	}
}

func modelToProtoAlert(m *model.Alert) *andon.Alert {
	var acknowledgedAt, resolvedAt int64
	if m.AcknowledgedAt != nil {
		acknowledgedAt = m.AcknowledgedAt.Unix()
	}
	if m.ResolvedAt != nil {
		resolvedAt = m.ResolvedAt.Unix()
	}
	var configID int64
	if m.Config != nil {
		configID = int64(m.Config.ID)
	}

	return &andon.Alert{
		Id:             int64(m.ID),
		ConfigId:       configID,
		TargetId:       parseInt64FromString(m.TargetID),
		TargetType:     m.TargetType,
		Status:         stringToAlertStatus(m.Status),
		TriggeredAt:    m.TriggeredAt.Unix(),
		AcknowledgedAt: acknowledgedAt,
		ResolvedAt:     resolvedAt,
	}
}

func protoPageResponse(page, pageSize, total, totalPages int32) *common.PageResponse {
	return &common.PageResponse{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
	}
}

// --- string <-> proto enum helpers ---

func stringToAndonType(s string) andon.AndonType {
	switch s {
	case "MATERIAL":
		return andon.AndonType_ANDON_TYPE_MATERIAL
	case "EQUIPMENT":
		return andon.AndonType_ANDON_TYPE_EQUIPMENT
	case "QUALITY":
		return andon.AndonType_ANDON_TYPE_QUALITY
	case "SAFETY":
		return andon.AndonType_ANDON_TYPE_SAFETY
	default:
		return andon.AndonType_ANDON_TYPE_UNSPECIFIED
	}
}

func protoAndonTypeToString(t andon.AndonType) string {
	switch t {
	case andon.AndonType_ANDON_TYPE_MATERIAL:
		return "MATERIAL"
	case andon.AndonType_ANDON_TYPE_EQUIPMENT:
		return "EQUIPMENT"
	case andon.AndonType_ANDON_TYPE_QUALITY:
		return "QUALITY"
	case andon.AndonType_ANDON_TYPE_SAFETY:
		return "SAFETY"
	default:
		return ""
	}
}

func stringToAndonCallStatus(s string) andon.AndonCallStatus {
	switch s {
	case "TRIGGERED":
		return andon.AndonCallStatus_ANDON_CALL_STATUS_TRIGGERED
	case "ACKNOWLEDGED":
		return andon.AndonCallStatus_ANDON_CALL_STATUS_ACKNOWLEDGED
	case "IN_PROGRESS":
		return andon.AndonCallStatus_ANDON_CALL_STATUS_IN_PROGRESS
	case "RESOLVED":
		return andon.AndonCallStatus_ANDON_CALL_STATUS_RESOLVED
	case "CLOSED":
		return andon.AndonCallStatus_ANDON_CALL_STATUS_CLOSED
	case "CANCELLED":
		return andon.AndonCallStatus_ANDON_CALL_STATUS_CANCELLED
	case "ESCALATED":
		return andon.AndonCallStatus_ANDON_CALL_STATUS_ESCALATED
	default:
		return andon.AndonCallStatus_ANDON_CALL_STATUS_UNSPECIFIED
	}
}

func protoAndonCallStatusToString(s andon.AndonCallStatus) string {
	switch s {
	case andon.AndonCallStatus_ANDON_CALL_STATUS_TRIGGERED:
		return "TRIGGERED"
	case andon.AndonCallStatus_ANDON_CALL_STATUS_ACKNOWLEDGED:
		return "ACKNOWLEDGED"
	case andon.AndonCallStatus_ANDON_CALL_STATUS_IN_PROGRESS:
		return "IN_PROGRESS"
	case andon.AndonCallStatus_ANDON_CALL_STATUS_RESOLVED:
		return "RESOLVED"
	case andon.AndonCallStatus_ANDON_CALL_STATUS_CLOSED:
		return "CLOSED"
	case andon.AndonCallStatus_ANDON_CALL_STATUS_CANCELLED:
		return "CANCELLED"
	case andon.AndonCallStatus_ANDON_CALL_STATUS_ESCALATED:
		return "ESCALATED"
	default:
		return ""
	}
}

func stringToTriggerType(s string) andon.TriggerType {
	switch s {
	case "THRESHOLD":
		return andon.TriggerType_TRIGGER_TYPE_THRESHOLD
	case "EVENT":
		return andon.TriggerType_TRIGGER_TYPE_EVENT
	case "SCHEDULE":
		return andon.TriggerType_TRIGGER_TYPE_SCHEDULE
	default:
		return andon.TriggerType_TRIGGER_TYPE_UNSPECIFIED
	}
}

func protoTriggerTypeToString(t andon.TriggerType) string {
	switch t {
	case andon.TriggerType_TRIGGER_TYPE_THRESHOLD:
		return "THRESHOLD"
	case andon.TriggerType_TRIGGER_TYPE_EVENT:
		return "EVENT"
	case andon.TriggerType_TRIGGER_TYPE_SCHEDULE:
		return "SCHEDULE"
	default:
		return ""
	}
}

func stringToAlertSeverity(s string) andon.AlertSeverity {
	switch s {
	case "P0":
		return andon.AlertSeverity_ALERT_SEVERITY_P0
	case "P1":
		return andon.AlertSeverity_ALERT_SEVERITY_P1
	case "P2":
		return andon.AlertSeverity_ALERT_SEVERITY_P2
	case "P3":
		return andon.AlertSeverity_ALERT_SEVERITY_P3
	default:
		return andon.AlertSeverity_ALERT_SEVERITY_UNSPECIFIED
	}
}

func protoAlertSeverityToString(s andon.AlertSeverity) string {
	switch s {
	case andon.AlertSeverity_ALERT_SEVERITY_P0:
		return "P0"
	case andon.AlertSeverity_ALERT_SEVERITY_P1:
		return "P1"
	case andon.AlertSeverity_ALERT_SEVERITY_P2:
		return "P2"
	case andon.AlertSeverity_ALERT_SEVERITY_P3:
		return "P3"
	default:
		return ""
	}
}

func stringToAlertStatus(s string) andon.AlertStatus {
	switch s {
	case "ACTIVE":
		return andon.AlertStatus_ALERT_STATUS_ACTIVE
	case "ACKNOWLEDGED":
		return andon.AlertStatus_ALERT_STATUS_ACKNOWLEDGED
	case "RESOLVED":
		return andon.AlertStatus_ALERT_STATUS_RESOLVED
	case "ESCALATED":
		return andon.AlertStatus_ALERT_STATUS_ESCALATED
	case "SUPPRESSED":
		return andon.AlertStatus_ALERT_STATUS_SUPPRESSED
	case "CLOSED":
		return andon.AlertStatus_ALERT_STATUS_CLOSED
	default:
		return andon.AlertStatus_ALERT_STATUS_UNSPECIFIED
	}
}

func protoAlertStatusToString(s andon.AlertStatus) string {
	switch s {
	case andon.AlertStatus_ALERT_STATUS_ACTIVE:
		return "ACTIVE"
	case andon.AlertStatus_ALERT_STATUS_ACKNOWLEDGED:
		return "ACKNOWLEDGED"
	case andon.AlertStatus_ALERT_STATUS_RESOLVED:
		return "RESOLVED"
	case andon.AlertStatus_ALERT_STATUS_ESCALATED:
		return "ESCALATED"
	case andon.AlertStatus_ALERT_STATUS_SUPPRESSED:
		return "SUPPRESSED"
	case andon.AlertStatus_ALERT_STATUS_CLOSED:
		return "CLOSED"
	default:
		return ""
	}
}

// --- helpers ---

func parseInt64FromString(s string) int64 {
	if s == "" {
		return 0
	}
	var n int64
	for _, c := range s {
		if c >= '0' && c <= '9' {
			n = n*10 + int64(c-'0')
		} else {
			return 0
		}
	}
	return n
}
