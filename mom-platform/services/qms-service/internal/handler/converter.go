package handler

import (
	"strconv"
	"time"

	qms "github.com/ninghonggang/mom-platform/gen/qms"
	common "github.com/ninghonggang/mom-platform/gen/common"

	"mom-platform/services/qms-service/internal/model"
)

// =============================================================================
// InspectionType
// =============================================================================

var modelToProtoInspectionType = map[string]qms.InspectionType{
	"INCOMING":     qms.InspectionType_INSPECTION_TYPE_IQC,
	"IN_PROCESS":   qms.InspectionType_INSPECTION_TYPE_IPQC,
	"FINAL":        qms.InspectionType_INSPECTION_TYPE_FQC,
	"FIRST_ARTICLE": qms.InspectionType_INSPECTION_TYPE_OQC,
}

var protoToModelInspectionType = map[qms.InspectionType]string{
	qms.InspectionType_INSPECTION_TYPE_IQC:  "INCOMING",
	qms.InspectionType_INSPECTION_TYPE_IPQC: "IN_PROCESS",
	qms.InspectionType_INSPECTION_TYPE_FQC:  "FINAL",
	qms.InspectionType_INSPECTION_TYPE_OQC:  "FIRST_ARTICLE",
}

func modelToProtoIT(s string) qms.InspectionType {
	if t, ok := modelToProtoInspectionType[s]; ok {
		return t
	}
	return qms.InspectionType_INSPECTION_TYPE_UNSPECIFIED
}

func protoToModelIT(t qms.InspectionType) string {
	if s, ok := protoToModelInspectionType[t]; ok {
		return s
	}
	return "INCOMING"
}

// =============================================================================
// InspectionSheetStatus
// =============================================================================

var modelToProtoSheetStatus = map[string]qms.InspectionSheetStatus{
	model.SheetStatusPending:    qms.InspectionSheetStatus_INSPECTION_SHEET_STATUS_PENDING,
	model.SheetStatusInProgress: qms.InspectionSheetStatus_INSPECTION_SHEET_STATUS_IN_PROGRESS,
	model.SheetStatusPassed:     qms.InspectionSheetStatus_INSPECTION_SHEET_STATUS_PASSED,
	model.SheetStatusFailed:     qms.InspectionSheetStatus_INSPECTION_SHEET_STATUS_FAILED,
	model.SheetStatusWaived:     qms.InspectionSheetStatus_INSPECTION_SHEET_STATUS_WAIVED,
	model.SheetStatusCancelled:  qms.InspectionSheetStatus_INSPECTION_SHEET_STATUS_CANCELLED,
}

var protoToModelSheetStatus = map[qms.InspectionSheetStatus]string{
	qms.InspectionSheetStatus_INSPECTION_SHEET_STATUS_PENDING:     model.SheetStatusPending,
	qms.InspectionSheetStatus_INSPECTION_SHEET_STATUS_IN_PROGRESS: model.SheetStatusInProgress,
	qms.InspectionSheetStatus_INSPECTION_SHEET_STATUS_PASSED:      model.SheetStatusPassed,
	qms.InspectionSheetStatus_INSPECTION_SHEET_STATUS_FAILED:      model.SheetStatusFailed,
	qms.InspectionSheetStatus_INSPECTION_SHEET_STATUS_WAIVED:      model.SheetStatusWaived,
	qms.InspectionSheetStatus_INSPECTION_SHEET_STATUS_CANCELLED:   model.SheetStatusCancelled,
}

func modelToProtoSS(s string) qms.InspectionSheetStatus {
	if t, ok := modelToProtoSheetStatus[s]; ok {
		return t
	}
	return qms.InspectionSheetStatus_INSPECTION_SHEET_STATUS_UNSPECIFIED
}

func protoToModelSS(t qms.InspectionSheetStatus) string {
	if s, ok := protoToModelSheetStatus[t]; ok {
		return s
	}
	return model.SheetStatusPending
}

// =============================================================================
// CharDataType
// =============================================================================

var modelToProtoDataType = map[string]qms.CharDataType{
	"NUMERIC": qms.CharDataType_CHAR_DATA_TYPE_NUMBER,
	"TEXT":    qms.CharDataType_CHAR_DATA_TYPE_TEXT,
	"BOOLEAN": qms.CharDataType_CHAR_DATA_TYPE_BOOLEAN,
}

var protoToModelDataType = map[qms.CharDataType]string{
	qms.CharDataType_CHAR_DATA_TYPE_NUMBER:  "NUMERIC",
	qms.CharDataType_CHAR_DATA_TYPE_TEXT:    "TEXT",
	qms.CharDataType_CHAR_DATA_TYPE_BOOLEAN: "BOOLEAN",
}

func modelToProtoDT(s string) qms.CharDataType {
	if t, ok := modelToProtoDataType[s]; ok {
		return t
	}
	return qms.CharDataType_CHAR_DATA_TYPE_UNSPECIFIED
}

func protoToModelDT(t qms.CharDataType) string {
	if s, ok := protoToModelDataType[t]; ok {
		return s
	}
	return "NUMERIC"
}

// =============================================================================
// NcrSeverity
// =============================================================================

var modelToProtoSeverity = map[string]qms.NcrSeverity{
	"CRITICAL": qms.NcrSeverity_NCR_SEVERITY_CRITICAL,
	"MAJOR":    qms.NcrSeverity_NCR_SEVERITY_MAJOR,
	"MINOR":    qms.NcrSeverity_NCR_SEVERITY_MINOR,
}

var protoToModelSeverity = map[qms.NcrSeverity]string{
	qms.NcrSeverity_NCR_SEVERITY_CRITICAL: "CRITICAL",
	qms.NcrSeverity_NCR_SEVERITY_MAJOR:    "MAJOR",
	qms.NcrSeverity_NCR_SEVERITY_MINOR:    "MINOR",
}

func modelToProtoSev(s string) qms.NcrSeverity {
	if t, ok := modelToProtoSeverity[s]; ok {
		return t
	}
	return qms.NcrSeverity_NCR_SEVERITY_UNSPECIFIED
}

func protoToModelSev(t qms.NcrSeverity) string {
	if s, ok := protoToModelSeverity[t]; ok {
		return s
	}
	return "MINOR"
}

// =============================================================================
// NcrStatus
// =============================================================================

var modelToProtoNcrStatus = map[string]qms.NcrStatus{
	model.NcrStatusOpen:          qms.NcrStatus_NCR_STATUS_OPEN,
	model.NcrStatusInvestigating: qms.NcrStatus_NCR_STATUS_INVESTIGATING,
	model.NcrStatusDispositioned: qms.NcrStatus_NCR_STATUS_DISPOSITIONED,
	model.NcrStatusVerified:      qms.NcrStatus_NCR_STATUS_VERIFIED,
	model.NcrStatusClosed:        qms.NcrStatus_NCR_STATUS_CLOSED,
	model.NcrStatusCancelled:     qms.NcrStatus_NCR_STATUS_CANCELLED,
	model.NcrStatusReopened:      qms.NcrStatus_NCR_STATUS_REOPENED,
}

var protoToModelNcrStatus = map[qms.NcrStatus]string{
	qms.NcrStatus_NCR_STATUS_OPEN:          model.NcrStatusOpen,
	qms.NcrStatus_NCR_STATUS_INVESTIGATING: model.NcrStatusInvestigating,
	qms.NcrStatus_NCR_STATUS_DISPOSITIONED: model.NcrStatusDispositioned,
	qms.NcrStatus_NCR_STATUS_VERIFIED:      model.NcrStatusVerified,
	qms.NcrStatus_NCR_STATUS_CLOSED:        model.NcrStatusClosed,
	qms.NcrStatus_NCR_STATUS_CANCELLED:     model.NcrStatusCancelled,
	qms.NcrStatus_NCR_STATUS_REOPENED:      model.NcrStatusReopened,
}

func modelToProtoNS(s string) qms.NcrStatus {
	if t, ok := modelToProtoNcrStatus[s]; ok {
		return t
	}
	return qms.NcrStatus_NCR_STATUS_UNSPECIFIED
}

func protoToModelNS(t qms.NcrStatus) string {
	if s, ok := protoToModelNcrStatus[t]; ok {
		return s
	}
	return model.NcrStatusOpen
}

// =============================================================================
// InspectionSheet <-> proto
// =============================================================================

func modelToProtoInspectionSheet(s *model.InspectionSheet) *qms.InspectionSheet {
	p := &qms.InspectionSheet{
		Id:             int64(s.ID),
		TenantId:       parseInt64(s.TenantID),
		SheetNo:        s.SheetNo,
		InspectionType: modelToProtoIT(s.InspectionType),
		MaterialId:     parseInt64(s.MaterialID),
		BatchId:        parseInt64(s.BatchID),
		SampleSize:     strconv.Itoa(s.SampleSize),
		DefectCount:    strconv.Itoa(s.DefectCount),
		Status:         modelToProtoSS(s.Status),
		InspectorId:    parseInt64(s.InspectorID),
		CreatedAt:      s.CreatedAt.Unix(),
		UpdatedAt:      s.UpdatedAt.Unix(),
	}
	if s.InspectedAt != nil {
		p.InspectedAt = s.InspectedAt.Unix()
	}
	return p
}

func protoCreateReqToModel(req *qms.CreateInspectionSheetRequest) *model.InspectionSheet {
	return &model.InspectionSheet{
		InspectionType: protoToModelIT(req.GetInspectionType()),
		MaterialID:     strconv.FormatInt(req.GetMaterialId(), 10),
		BatchID:        strconv.FormatInt(req.GetBatchId(), 10),
		SampleSize:     int(protoParseInt(req.GetSampleSize())),
		InspectorID:    strconv.FormatInt(req.GetInspectorId(), 10),
	}
}

func protoUpdateReqToModel(req *qms.UpdateInspectionSheetRequest) *model.InspectionSheet {
	return &model.InspectionSheet{
		Status:      protoToModelSS(req.GetStatus()),
		DefectCount: int(protoParseInt(req.GetDefectCount())),
		InspectorID: strconv.FormatInt(req.GetInspectorId(), 10),
	}
}

// =============================================================================
// InspectionCharacteristic <-> proto
// =============================================================================

func modelToProtoCharacteristic(c *model.InspectionCharacteristic) *qms.InspectionCharacteristic {
	return &qms.InspectionCharacteristic{
		Id:       int64(c.ID),
		CharCode: c.CharCode,
		CharName: c.CharName,
		DataType: modelToProtoDT(c.DataType),
		Usl:      strconv.FormatFloat(c.USL, 'f', -1, 64),
		Lsl:      strconv.FormatFloat(c.LSL, 'f', -1, 64),
		Target:   strconv.FormatFloat(c.Target, 'f', -1, 64),
		Unit:     c.Unit,
	}
}

func protoCreateCharReqToModel(req *qms.CreateCharacteristicRequest) *model.InspectionCharacteristic {
	return &model.InspectionCharacteristic{
		CharCode: req.GetCharCode(),
		CharName: req.GetCharName(),
		DataType: protoToModelDT(req.GetDataType()),
		USL:      parseFloat64(req.GetUsl()),
		LSL:      parseFloat64(req.GetLsl()),
		Target:   parseFloat64(req.GetTarget()),
		Unit:     req.GetUnit(),
	}
}

// =============================================================================
// InspectionPlan <-> proto
// =============================================================================

func modelToProtoInspectionPlan(p *model.InspectionPlan) *qms.InspectionPlan {
	return &qms.InspectionPlan{
		Id:          int64(p.ID),
		SchemeCode:  p.SchemeCode,
		SchemeName:  p.SchemeName,
		SchemeType:  p.SchemeType,
		TemplateId:  parseInt64(p.TemplateID),
		Status:      p.Status,
	}
}

func protoCreatePlanReqToModel(req *qms.CreateInspectionPlanRequest) *model.InspectionPlan {
	return &model.InspectionPlan{
		SchemeCode: req.GetSchemeCode(),
		SchemeName: req.GetSchemeName(),
		SchemeType: req.GetSchemeType(),
		TemplateID: strconv.FormatInt(req.GetTemplateId(), 10),
	}
}

// =============================================================================
// Ncr <-> proto
// =============================================================================

func modelToProtoNcr(n *model.Ncr) *qms.Ncr {
	return &qms.Ncr{
		Id:                int64(n.ID),
		TenantId:          parseInt64(n.TenantID),
		NcrNo:             n.NcrNo,
		InspectionSheetId: int64(n.InspectionSheetID),
		MaterialId:        parseInt64(n.MaterialID),
		BatchId:           parseInt64(n.BatchID),
		Quantity:          strconv.FormatFloat(n.Quantity, 'f', -1, 64),
		Severity:          modelToProtoSev(n.Severity),
		Status:            modelToProtoNS(n.Status),
		CreatedAt:         n.CreatedAt.Unix(),
		UpdatedAt:         n.UpdatedAt.Unix(),
	}
}

func protoCreateNcrReqToModel(req *qms.CreateNcrRequest) *model.Ncr {
	return &model.Ncr{
		InspectionSheetID: uint(req.GetInspectionSheetId()),
		MaterialID:        strconv.FormatInt(req.GetMaterialId(), 10),
		BatchID:           strconv.FormatInt(req.GetBatchId(), 10),
		Quantity:          parseFloat64(req.GetQuantity()),
		Severity:          protoToModelSev(req.GetSeverity()),
	}
}

func protoUpdateNcrReqToModel(req *qms.UpdateNcrRequest) *model.Ncr {
	return &model.Ncr{
		ID:     uint(req.GetId()),
		Status: protoToModelNS(req.GetStatus()),
	}
}

// =============================================================================
// NcrAction <-> proto
// =============================================================================

func modelToProtoNcrAction(a *model.NcrAction) *qms.NcrAction {
	return &qms.NcrAction{
		Id:            int64(a.ID),
		NcrId:         int64(a.NcrID),
		ActionType:    a.ActionType,
		ActionDesc:    a.ActionDesc,
		ResponsibleId: parseInt64(a.ResponsibleID),
	}
}

func protoAddActionReqToModel(req *qms.AddNcrActionRequest) *model.NcrAction {
	return &model.NcrAction{
		NcrID:         uint(req.GetNcrId()),
		ActionType:    req.GetActionType(),
		ActionDesc:    req.GetActionDesc(),
		ResponsibleID: strconv.FormatInt(req.GetResponsibleId(), 10),
	}
}

// =============================================================================
// DefectCode <-> proto
// =============================================================================

func modelToProtoDefectCode(d *model.DefectCode) *qms.DefectCode {
	return &qms.DefectCode{
		Id:          int64(d.ID),
		DefectCode:  d.DefectCode,
		DefectName:  d.DefectName,
		DefectClass: d.DefectClass,
		Severity:    modelToProtoSev(d.Severity),
	}
}

func protoCreateDefectReqToModel(req *qms.CreateDefectCodeRequest) *model.DefectCode {
	return &model.DefectCode{
		DefectCode:  req.GetDefectCode(),
		DefectName:  req.GetDefectName(),
		DefectClass: req.GetDefectClass(),
		Severity:    protoToModelSev(req.GetSeverity()),
	}
}

// =============================================================================
// SpcData <-> proto
// =============================================================================

func modelToProtoSpcData(s *model.SpcData) *qms.SpcData {
	return &qms.SpcData{
		Id:          int64(s.ID),
		CharId:      int64(s.CharID),
		SampleValue: strconv.FormatFloat(s.SampleValue, 'f', -1, 64),
		SampleTime:  s.SampleTime.Unix(),
		Xbar:        strconv.FormatFloat(s.Xbar, 'f', -1, 64),
		RValue:      strconv.FormatFloat(s.RValue, 'f', -1, 64),
	}
}

func protoRecordSpcReqToModel(req *qms.RecordSpcDataRequest) *model.SpcData {
	return &model.SpcData{
		CharID:      uint(req.GetCharId()),
		SampleValue: parseFloat64(req.GetSampleValue()),
		SampleTime:  time.Now(),
	}
}

// =============================================================================
// Page helpers
// =============================================================================

func protoToPageRequest(p *common.PageRequest) (page, pageSize int) {
	if p == nil {
		return 1, 20
	}
	page = int(p.GetPage())
	pageSize = int(p.GetPageSize())
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	// 超限时截断到上限，而不是回落为默认值——后者会让 pageSize=500 静默变成 20。
	if pageSize > 200 {
		pageSize = 200
	}
	return
}

func newProtoPageResponse(page, pageSize int, total int64) *common.PageResponse {
	totalPages := int32(0)
	if pageSize > 0 {
		totalPages = int32((total + int64(pageSize) - 1) / int64(pageSize))
	}
	return &common.PageResponse{
		Page:       int32(page),
		PageSize:   int32(pageSize),
		Total:      int32(total),
		TotalPages: totalPages,
	}
}

// =============================================================================
// Numeric helpers
// =============================================================================

func parseInt64(s string) int64 {
	if s == "" {
		return 0
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return v
}

func parseFloat64(s string) float64 {
	if s == "" {
		return 0
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return v
}

func protoParseInt(s string) int64 {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return v
}
