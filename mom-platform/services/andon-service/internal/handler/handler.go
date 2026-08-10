package handler

import (
	"context"
	"errors"
	"strconv"

	"github.com/ninghonggang/mom-platform/pkg/eventbus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	andon "github.com/ninghonggang/mom-platform/gen/andon"
	common "github.com/ninghonggang/mom-platform/gen/common"

	"mom-platform/services/andon-service/internal/service"
)

// AndonHandler implements andon.AndonServiceServer gRPC interface.
type AndonHandler struct {
	andon.UnimplementedAndonServiceServer
	svc    *service.AndonService
	pub    *eventbus.EventPublisher
	logger *zap.Logger
}

func NewAndonHandler(svc *service.AndonService, pub *eventbus.EventPublisher, logger *zap.Logger) *AndonHandler {
	return &AndonHandler{svc: svc, pub: pub, logger: logger}
}

// --- 安灯呼叫 ---

func (h *AndonHandler) TriggerAndon(ctx context.Context, req *andon.TriggerAndonRequest) (*andon.AndonCall, error) {
	workstationID := strconv.FormatInt(req.WorkstationId, 10)
	reporterID := strconv.FormatInt(req.ReporterId, 10)
	andonTypeStr := protoAndonTypeToString(req.AndonType)

	call, err := h.svc.TriggerAndon(ctx,
		"", workstationID, reporterID,
		andonTypeStr, req.Description,
	)
	if err != nil {
		h.logger.Error("TriggerAndon failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Publish andon.triggered event
	if h.pub != nil {
		_ = h.pub.Publish(ctx, eventbus.SubjectAndonTriggered, call)
	}

	return modelToProtoAndonCall(call), nil
}

func (h *AndonHandler) AcknowledgeAndon(ctx context.Context, req *andon.AcknowledgeAndonRequest) (*andon.AndonCall, error) {
	operatorID := strconv.FormatInt(req.OperatorId, 10)

	call, err := h.svc.AcknowledgeAndon(ctx, uint(req.Id), operatorID)
	if err != nil {
		h.logger.Error("AcknowledgeAndon failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return modelToProtoAndonCall(call), nil
}

func (h *AndonHandler) ResolveAndon(ctx context.Context, req *andon.ResolveAndonRequest) (*andon.AndonCall, error) {
	operatorID := strconv.FormatInt(req.OperatorId, 10)

	call, err := h.svc.ResolveAndon(ctx, uint(req.Id), operatorID, req.Resolution)
	if err != nil {
		h.logger.Error("ResolveAndon failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return modelToProtoAndonCall(call), nil
}

func (h *AndonHandler) EscalateAndon(ctx context.Context, req *andon.EscalateAndonRequest) (*andon.AndonCall, error) {
	reason := "escalated to level " + strconv.FormatInt(int64(req.Level), 10)

	call, err := h.svc.EscalateAndon(ctx, uint(req.Id), "system", reason)
	if err != nil {
		h.logger.Error("EscalateAndon failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Publish andon.escalated event
	if h.pub != nil {
		_ = h.pub.Publish(ctx, eventbus.SubjectAndonEscalated, call)
	}

	return modelToProtoAndonCall(call), nil
}

func (h *AndonHandler) GetAndonCall(ctx context.Context, req *common.IdRequest) (*andon.AndonCall, error) {
	call, err := h.svc.GetAndonCall(ctx, uint(req.Id))
	if err != nil {
		h.logger.Error("GetAndonCall failed", zap.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return modelToProtoAndonCall(call), nil
}

func (h *AndonHandler) ListAndonCalls(ctx context.Context, req *andon.ListAndonCallsRequest) (*andon.ListAndonCallsResponse, error) {
	page := int32(1)
	pageSize := int32(20)
	if req.Page != nil {
		page = req.Page.Page
		pageSize = req.Page.PageSize
	}
	andonTypeStr := protoAndonTypeToString(req.AndonType)
	statusStr := protoAndonCallStatusToString(req.Status)
	workstationID := strconv.FormatInt(req.WorkstationId, 10)

	calls, total, err := h.svc.ListAndonCalls(ctx, int(page), int(pageSize), andonTypeStr, statusStr, workstationID)
	if err != nil {
		h.logger.Error("ListAndonCalls failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	items := make([]*andon.AndonCall, len(calls))
	for i, c := range calls {
		items[i] = modelToProtoAndonCall(c)
	}

	totalPages := int32(0)
	if pageSize > 0 {
		totalPages = (int32(total) + pageSize - 1) / pageSize
	}

	return &andon.ListAndonCallsResponse{
		Items: items,
		Page:  protoPageResponse(page, pageSize, int32(total), totalPages),
	}, nil
}

// --- 告警配置 ---

func (h *AndonHandler) CreateAlertConfig(ctx context.Context, req *andon.CreateAlertConfigRequest) (*andon.AlertConfig, error) {
	triggerTypeStr := protoTriggerTypeToString(req.TriggerType)
	severityStr := protoAlertSeverityToString(req.Severity)

	cfg, err := h.svc.CreateAlertConfig(ctx, req.ConfigCode, req.ConfigName,
		triggerTypeStr, severityStr,
		req.TriggerCondition, req.NotifyChannels)
	if err != nil {
		h.logger.Error("CreateAlertConfig failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return modelToProtoAlertConfig(cfg), nil
}

// --- 告警 ---

func (h *AndonHandler) TriggerAlert(ctx context.Context, req *andon.TriggerAlertRequest) (*andon.Alert, error) {
	if req.GetConfigId() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "config_id is required")
	}
	targetID := strconv.FormatInt(req.TargetId, 10)

	// 按配置主键解析告警配置
	alertEntity, err := h.svc.TriggerAlertByConfigID(ctx, uint(req.GetConfigId()), targetID, req.TargetType)
	if err != nil {
		h.logger.Error("TriggerAlert failed", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "alert config %d not found", req.GetConfigId())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Publish alert.triggered event
	if h.pub != nil {
		_ = h.pub.Publish(ctx, eventbus.SubjectAlertTriggered, alertEntity)
	}

	return modelToProtoAlert(alertEntity), nil
}

func (h *AndonHandler) AcknowledgeAlert(ctx context.Context, req *common.IdRequest) (*andon.Alert, error) {
	alertEntity, err := h.svc.AcknowledgeAlert(ctx, uint(req.Id))
	if err != nil {
		h.logger.Error("AcknowledgeAlert failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return modelToProtoAlert(alertEntity), nil
}

func (h *AndonHandler) ResolveAlert(ctx context.Context, req *common.IdRequest) (*andon.Alert, error) {
	alertEntity, err := h.svc.ResolveAlert(ctx, uint(req.Id))
	if err != nil {
		h.logger.Error("ResolveAlert failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return modelToProtoAlert(alertEntity), nil
}

func (h *AndonHandler) ListAlerts(ctx context.Context, req *andon.ListAlertsRequest) (*andon.ListAlertsResponse, error) {
	page := int32(1)
	pageSize := int32(20)
	if req.Page != nil {
		page = req.Page.Page
		pageSize = req.Page.PageSize
	}
	statusStr := protoAlertStatusToString(req.Status)
	severityStr := protoAlertSeverityToString(req.Severity)

	alerts, total, err := h.svc.ListAlerts(ctx, int(page), int(pageSize), statusStr, severityStr)
	if err != nil {
		h.logger.Error("ListAlerts failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	items := make([]*andon.Alert, len(alerts))
	for i, a := range alerts {
		items[i] = modelToProtoAlert(a)
	}

	totalPages := int32(0)
	if pageSize > 0 {
		totalPages = (int32(total) + pageSize - 1) / pageSize
	}

	return &andon.ListAlertsResponse{
		Items: items,
		Page:  protoPageResponse(page, pageSize, int32(total), totalPages),
	}, nil
}
