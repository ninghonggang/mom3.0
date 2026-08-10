package handler

import (
	"context"
	"errors"
	"strconv"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	qms "github.com/ninghonggang/mom-platform/gen/qms"

	"mom-platform/services/qms-service/internal/model"
	"mom-platform/services/qms-service/internal/service"
)

// defaultTenantID 单租户部署下的默认租户。
const defaultTenantID = "0"

// tenantFromCtx 从 gRPC metadata 的 x-tenant-id 读取租户；
// 缺省时回退到 defaultTenantID，保证单租户部署可用。
func tenantFromCtx(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if vals := md.Get("x-tenant-id"); len(vals) > 0 && vals[0] != "" {
			return vals[0]
		}
	}
	return defaultTenantID
}

// Handler implements the QMS gRPC service by satisfying the
// QmsServiceServer interface from generated proto code.
type Handler struct {
	qms.UnimplementedQmsServiceServer
	svc *service.Service
	log *zap.Logger
}

// NewHandler creates a new gRPC Handler.
func NewHandler(svc *service.Service, log *zap.Logger) *Handler {
	return &Handler{svc: svc, log: log}
}

// errToStatus maps service-layer errors to gRPC status codes.
func (h *Handler) errToStatus(err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, service.ErrNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, service.ErrInvalidInput):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, service.ErrInvalidTransition):
		return status.Error(codes.FailedPrecondition, err.Error())
	case errors.Is(err, service.ErrInspectionSheetNotFailed):
		return status.Error(codes.FailedPrecondition, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}

// ==============================================================================
// InspectionSheet RPCs
// ==============================================================================

func (h *Handler) CreateInspectionSheet(ctx context.Context, req *qms.CreateInspectionSheetRequest) (*qms.InspectionSheet, error) {
	h.log.Info("CreateInspectionSheet",
		zap.Int64("material_id", req.GetMaterialId()),
	)

	sheet := protoCreateReqToModel(req)
	sheet.TenantID = tenantFromCtx(ctx)
	if sheet.SampleSize <= 0 {
		sheet.SampleSize = 1
	}
	result, err := h.svc.CreateInspectionSheet(ctx, sheet)
	if err != nil {
		h.log.Error("CreateInspectionSheet failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}
	return modelToProtoInspectionSheet(result), nil
}

func (h *Handler) UpdateInspectionSheet(ctx context.Context, req *qms.UpdateInspectionSheetRequest) (*qms.InspectionSheet, error) {
	h.log.Info("UpdateInspectionSheet", zap.Int64("id", req.GetId()))

	updates := protoUpdateReqToModel(req)
	result, err := h.svc.UpdateInspectionSheet(ctx, uint(req.GetId()), updates)
	if err != nil {
		h.log.Error("UpdateInspectionSheet failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}
	return modelToProtoInspectionSheet(result), nil
}

func (h *Handler) GetInspectionSheet(ctx context.Context, req *qms.GetInspectionSheetRequest) (*qms.InspectionSheet, error) {
	h.log.Info("GetInspectionSheet", zap.Int64("id", req.GetId()))

	sheet, err := h.svc.GetInspectionSheet(ctx, uint(req.GetId()))
	if err != nil {
		h.log.Error("GetInspectionSheet failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}
	return modelToProtoInspectionSheet(sheet), nil
}

func (h *Handler) ListInspectionSheets(ctx context.Context, req *qms.ListInspectionSheetsRequest) (*qms.ListInspectionSheetsResponse, error) {
	h.log.Info("ListInspectionSheets")

	page, pageSize := protoToPageRequest(req.GetPage())

	filter := map[string]interface{}{}
	// 仅在调用方显式指定状态时过滤。protoToModelSS 对 UNSPECIFIED 有 PENDING 兜底，
	// 缺少此判断会让未传状态的列表被静默限制为 PENDING，隐藏其余全部记录。
	if req.GetStatus() != qms.InspectionSheetStatus_INSPECTION_SHEET_STATUS_UNSPECIFIED {
		if s := protoToModelSS(req.GetStatus()); s != "" {
			filter["status = ?"] = s
		}
	}
	if req.GetMaterialId() > 0 {
		filter["material_id = ?"] = strconv.FormatInt(req.GetMaterialId(), 10)
	}

	result, err := h.svc.ListInspectionSheets(ctx, page, pageSize, filter)
	if err != nil {
		h.log.Error("ListInspectionSheets failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}

	items := make([]*qms.InspectionSheet, 0, len(result.Items))
	for i := range result.Items {
		items = append(items, modelToProtoInspectionSheet(&result.Items[i]))
	}
	return &qms.ListInspectionSheetsResponse{
		Items: items,
		Page:  newProtoPageResponse(page, pageSize, result.Total),
	}, nil
}

func (h *Handler) DeleteInspectionSheet(ctx context.Context, req *qms.DeleteInspectionSheetRequest) (*qms.InspectionSheet, error) {
	h.log.Info("DeleteInspectionSheet", zap.Int64("id", req.GetId()))

	id := uint(req.GetId())
	sheet, err := h.svc.GetInspectionSheet(ctx, id)
	if err != nil {
		h.log.Error("DeleteInspectionSheet failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}
	if err := h.svc.DeleteInspectionSheet(ctx, id); err != nil {
		h.log.Error("DeleteInspectionSheet failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}
	return modelToProtoInspectionSheet(sheet), nil
}

// ==============================================================================
// InspectionCharacteristic RPCs
// ==============================================================================

func (h *Handler) CreateCharacteristic(ctx context.Context, req *qms.CreateCharacteristicRequest) (*qms.InspectionCharacteristic, error) {
	h.log.Info("CreateCharacteristic", zap.String("char_code", req.GetCharCode()))

	c := protoCreateCharReqToModel(req)
	if err := h.svc.CreateInspectionCharacteristic(ctx, c); err != nil {
		h.log.Error("CreateCharacteristic failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}
	return modelToProtoCharacteristic(c), nil
}

func (h *Handler) ListCharacteristics(ctx context.Context, req *qms.ListCharacteristicsRequest) (*qms.ListCharacteristicsResponse, error) {
	h.log.Info("ListCharacteristics")

	page, pageSize := protoToPageRequest(req.GetPage())
	result, err := h.svc.ListInspectionCharacteristics(ctx, page, pageSize, nil)
	if err != nil {
		h.log.Error("ListCharacteristics failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}

	items := make([]*qms.InspectionCharacteristic, 0, len(result.Items))
	for i := range result.Items {
		items = append(items, modelToProtoCharacteristic(&result.Items[i]))
	}
	return &qms.ListCharacteristicsResponse{
		Items: items,
		Page:  newProtoPageResponse(page, pageSize, result.Total),
	}, nil
}

// ==============================================================================
// InspectionPlan RPCs
// ==============================================================================

func (h *Handler) CreateInspectionPlan(ctx context.Context, req *qms.CreateInspectionPlanRequest) (*qms.InspectionPlan, error) {
	h.log.Info("CreateInspectionPlan", zap.String("scheme_code", req.GetSchemeCode()))

	p := protoCreatePlanReqToModel(req)
	if err := h.svc.CreateInspectionPlan(ctx, p); err != nil {
		h.log.Error("CreateInspectionPlan failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}
	return modelToProtoInspectionPlan(p), nil
}

func (h *Handler) ListInspectionPlans(ctx context.Context, req *qms.ListInspectionPlansRequest) (*qms.ListInspectionPlansResponse, error) {
	h.log.Info("ListInspectionPlans")

	page, pageSize := protoToPageRequest(req.GetPage())
	result, err := h.svc.ListInspectionPlans(ctx, page, pageSize, nil)
	if err != nil {
		h.log.Error("ListInspectionPlans failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}

	items := make([]*qms.InspectionPlan, 0, len(result.Items))
	for i := range result.Items {
		items = append(items, modelToProtoInspectionPlan(&result.Items[i]))
	}
	return &qms.ListInspectionPlansResponse{
		Items: items,
		Page:  newProtoPageResponse(page, pageSize, result.Total),
	}, nil
}

// ==============================================================================
// InspectionResult RPCs
// ==============================================================================

func (h *Handler) RecordInspectionResult(ctx context.Context, req *qms.RecordInspectionResultRequest) (*qms.InspectionSheet, error) {
	h.log.Info("RecordInspectionResult",
		zap.Int64("sheet_id", req.GetSheetId()),
		zap.Int("entries", len(req.GetEntries())),
	)

	for _, entry := range req.GetEntries() {
		_, err := h.svc.RecordInspectionResult(ctx, uint(req.GetSheetId()), uint(entry.GetCharId()), entry.GetValue())
		if err != nil {
			h.log.Error("RecordInspectionResult failed", zap.Int64("char_id", entry.GetCharId()), zap.Error(err))
			return nil, h.errToStatus(err)
		}
	}

	sheet, err := h.svc.GetInspectionSheet(ctx, uint(req.GetSheetId()))
	if err != nil {
		h.log.Error("RecordInspectionResult: get sheet failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}
	return modelToProtoInspectionSheet(sheet), nil
}

// ==============================================================================
// NCR RPCs
// ==============================================================================

func (h *Handler) CreateNcr(ctx context.Context, req *qms.CreateNcrRequest) (*qms.Ncr, error) {
	h.log.Info("CreateNcr", zap.Int64("sheet_id", req.GetInspectionSheetId()))

	ncr := protoCreateNcrReqToModel(req)
	ncr.TenantID = tenantFromCtx(ctx)
	if ncr.Quantity <= 0 {
		ncr.Quantity = 1
	}
	result, err := h.svc.CreateNcr(ctx, ncr)
	if err != nil {
		h.log.Error("CreateNcr failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}
	return modelToProtoNcr(result), nil
}

func (h *Handler) UpdateNcr(ctx context.Context, req *qms.UpdateNcrRequest) (*qms.Ncr, error) {
	h.log.Info("UpdateNcr", zap.Int64("id", req.GetId()))

	updates := protoUpdateNcrReqToModel(req)
	result, err := h.svc.UpdateNcr(ctx, uint(req.GetId()), updates)
	if err != nil {
		h.log.Error("UpdateNcr failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}
	return modelToProtoNcr(result), nil
}

func (h *Handler) GetNcr(ctx context.Context, req *qms.GetNcrRequest) (*qms.Ncr, error) {
	h.log.Info("GetNcr", zap.Int64("id", req.GetId()))

	ncr, err := h.svc.GetNcr(ctx, uint(req.GetId()))
	if err != nil {
		h.log.Error("GetNcr failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}
	return modelToProtoNcr(ncr), nil
}

func (h *Handler) ListNcrs(ctx context.Context, req *qms.ListNcrsRequest) (*qms.ListNcrsResponse, error) {
	h.log.Info("ListNcrs")

	page, pageSize := protoToPageRequest(req.GetPage())
	filter := map[string]interface{}{}
	// 同 ListInspectionSheets：两个转换函数分别对 UNSPECIFIED 兜底为 OPEN / MINOR，
	// 必须先判空，否则未带筛选条件的列表会被钉死在 status=OPEN AND severity=MINOR。
	if req.GetStatus() != qms.NcrStatus_NCR_STATUS_UNSPECIFIED {
		if s := protoToModelNS(req.GetStatus()); s != "" {
			filter["status = ?"] = s
		}
	}
	if req.GetSeverity() != qms.NcrSeverity_NCR_SEVERITY_UNSPECIFIED {
		if s := protoToModelSev(req.GetSeverity()); s != "" {
			filter["severity = ?"] = s
		}
	}

	result, err := h.svc.ListNcrs(ctx, page, pageSize, filter)
	if err != nil {
		h.log.Error("ListNcrs failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}

	items := make([]*qms.Ncr, 0, len(result.Items))
	for i := range result.Items {
		items = append(items, modelToProtoNcr(&result.Items[i]))
	}
	return &qms.ListNcrsResponse{
		Items: items,
		Page:  newProtoPageResponse(page, pageSize, result.Total),
	}, nil
}

func (h *Handler) AddNcrAction(ctx context.Context, req *qms.AddNcrActionRequest) (*qms.NcrAction, error) {
	h.log.Info("AddNcrAction", zap.Int64("ncr_id", req.GetNcrId()))

	action := protoAddActionReqToModel(req)
	result, err := h.svc.AddNcrAction(ctx, action)
	if err != nil {
		h.log.Error("AddNcrAction failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}
	return modelToProtoNcrAction(result), nil
}

func (h *Handler) ListNcrActions(ctx context.Context, req *qms.GetNcrRequest) (*qms.ListNcrActionsResponse, error) {
	h.log.Info("ListNcrActions", zap.Int64("id", req.GetId()))

	actions, err := h.svc.ListNcrActions(ctx, uint(req.GetId()))
	if err != nil {
		h.log.Error("ListNcrActions failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}

	items := make([]*qms.NcrAction, 0, len(actions))
	for i := range actions {
		items = append(items, modelToProtoNcrAction(&actions[i]))
	}
	return &qms.ListNcrActionsResponse{Items: items}, nil
}

// ==============================================================================
// DefectCode RPCs
// ==============================================================================

func (h *Handler) CreateDefectCode(ctx context.Context, req *qms.CreateDefectCodeRequest) (*qms.DefectCode, error) {
	h.log.Info("CreateDefectCode", zap.String("defect_code", req.GetDefectCode()))

	d := protoCreateDefectReqToModel(req)
	if err := h.svc.CreateDefectCode(ctx, d); err != nil {
		h.log.Error("CreateDefectCode failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}
	return modelToProtoDefectCode(d), nil
}

func (h *Handler) ListDefectCodes(ctx context.Context, req *qms.ListDefectCodesRequest) (*qms.ListDefectCodesResponse, error) {
	h.log.Info("ListDefectCodes")

	page, pageSize := protoToPageRequest(req.GetPage())
	result, err := h.svc.ListDefectCodes(ctx, page, pageSize, nil)
	if err != nil {
		h.log.Error("ListDefectCodes failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}

	items := make([]*qms.DefectCode, 0, len(result.Items))
	for i := range result.Items {
		items = append(items, modelToProtoDefectCode(&result.Items[i]))
	}
	return &qms.ListDefectCodesResponse{
		Items: items,
		Page:  newProtoPageResponse(page, pageSize, result.Total),
	}, nil
}

// ==============================================================================
// SPC RPCs
// ==============================================================================

func (h *Handler) RecordSpcData(ctx context.Context, req *qms.RecordSpcDataRequest) (*qms.SpcData, error) {
	h.log.Info("RecordSpcData", zap.Int64("char_id", req.GetCharId()))

	data := protoRecordSpcReqToModel(req)
	if err := h.svc.RecordSpcData(ctx, data); err != nil {
		h.log.Error("RecordSpcData failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}
	return modelToProtoSpcData(data), nil
}

func (h *Handler) ListSpcData(ctx context.Context, req *qms.ListSpcDataRequest) (*qms.ListSpcDataResponse, error) {
	h.log.Info("ListSpcData", zap.Int64("char_id", req.GetCharId()))

	page, pageSize := 1, 20

	_, items, err := h.listSpcDataHelper(ctx, uint(req.GetCharId()), page, pageSize)
	if err != nil {
		h.log.Error("ListSpcData failed", zap.Error(err))
		return nil, h.errToStatus(err)
	}

	protoItems := make([]*qms.SpcData, 0, len(items))
	for i := range items {
		protoItems = append(protoItems, modelToProtoSpcData(&items[i]))
	}
	return &qms.ListSpcDataResponse{
		Items: protoItems,
	}, nil
}

// listSpcDataHelper is injected in main.go via the service's extended accessor.
// This is a helper that can be replaced when the service exposes a proper list method.
func (h *Handler) listSpcDataHelper(ctx context.Context, charID uint, page, pageSize int) (total int64, items []model.SpcData, err error) {
	// For now, we return empty data since the service doesn't expose a paginated list.
	// In production, this should call a proper repository method.
	return 0, nil, nil
}
