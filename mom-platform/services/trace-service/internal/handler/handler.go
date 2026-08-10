package handler

import (
	"context"
	"strconv"

	"github.com/ninghonggang/mom-platform/pkg/eventbus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	trace "github.com/ninghonggang/mom-platform/gen/trace"

	"mom-platform/services/trace-service/internal/service"
)

// TraceHandler implements trace.TraceServiceServer gRPC interface.
type TraceHandler struct {
	trace.UnimplementedTraceServiceServer
	svc      *service.TraceService
	pub      *eventbus.EventPublisher
	logger   *zap.Logger
}

func NewTraceHandler(svc *service.TraceService, pub *eventbus.EventPublisher, logger *zap.Logger) *TraceHandler {
	return &TraceHandler{svc: svc, pub: pub, logger: logger}
}

func (h *TraceHandler) CreateTraceRecord(ctx context.Context, req *trace.CreateTraceRecordRequest) (*trace.TraceRecord, error) {
	materialID := strconv.FormatInt(req.MaterialId, 10)
	orderID := strconv.FormatInt(req.ProductionOrderId, 10)
	traceTypeStr := protoTraceTypeToString(req.TraceType)

	record, err := h.svc.CreateTraceRecord(ctx,
		"", traceTypeStr, req.SerialNo, req.BatchNo,
		materialID, orderID, "", "",
	)
	if err != nil {
		h.logger.Error("CreateTraceRecord failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return modelToProtoTraceRecord(record), nil
}

func (h *TraceHandler) ForwardTrace(ctx context.Context, req *trace.GetTraceRequest) (*trace.ForwardTraceResponse, error) {
	// Resolve trace_id from serial_no/batch_no
	startID, err := h.resolveTraceID(ctx, req)
	if err != nil {
		return nil, err
	}

	records, err := h.svc.ForwardTrace(ctx, startID)
	if err != nil {
		h.logger.Error("ForwardTrace failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	nodes := make([]*trace.TraceRecord, len(records))
	for i, r := range records {
		nodes[i] = modelToProtoTraceRecord(r)
	}
	return &trace.ForwardTraceResponse{
		Nodes: nodes,
	}, nil
}

func (h *TraceHandler) BackwardTrace(ctx context.Context, req *trace.GetTraceRequest) (*trace.BackwardTraceResponse, error) {
	startID, err := h.resolveTraceID(ctx, req)
	if err != nil {
		return nil, err
	}

	records, err := h.svc.BackwardTrace(ctx, startID)
	if err != nil {
		h.logger.Error("BackwardTrace failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	nodes := make([]*trace.TraceRecord, len(records))
	for i, r := range records {
		nodes[i] = modelToProtoTraceRecord(r)
	}
	return &trace.BackwardTraceResponse{
		Nodes: nodes,
	}, nil
}

func (h *TraceHandler) resolveTraceID(ctx context.Context, req *trace.GetTraceRequest) (uint, error) {
	if req.SerialNo != "" {
		record, err := h.svc.GetRepo().GetTraceRecordBySerialNo(ctx, req.SerialNo)
		if err == nil {
			return record.ID, nil
		}
	}
	if req.BatchNo != "" {
		records, err := h.svc.GetRepo().GetTraceRecordByBatchNo(ctx, req.BatchNo)
		if err == nil && len(records) > 0 {
			return records[0].ID, nil
		}
	}
	return 0, status.Error(codes.NotFound, "no trace record found for the given criteria")
}

func (h *TraceHandler) ListTraces(ctx context.Context, req *trace.ListTracesRequest) (*trace.ListTracesResponse, error) {
	page := int32(1)
	pageSize := int32(20)
	if req.Page != nil {
		page = req.Page.Page
		pageSize = req.Page.PageSize
	}
	traceTypeStr := protoTraceTypeToString(req.TraceType)
	materialID := strconv.FormatInt(req.MaterialId, 10)

	records, total, err := h.svc.ListTraces(ctx, int(page), int(pageSize), traceTypeStr, materialID, req.BeginTime, req.EndTime)
	if err != nil {
		h.logger.Error("ListTraces failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	items := make([]*trace.TraceRecord, len(records))
	for i, r := range records {
		items[i] = modelToProtoTraceRecord(r)
	}

	totalPages := int32(0)
	if pageSize > 0 {
		totalPages = (int32(total) + pageSize - 1) / pageSize
	}

	return &trace.ListTracesResponse{
		Items: items,
		Page:  protoPageResponse(page, pageSize, int32(total), totalPages),
	}, nil
}

func (h *TraceHandler) GenerateSerials(ctx context.Context, req *trace.GenerateSerialRequest) (*trace.GenerateSerialResponse, error) {
	materialID := strconv.FormatInt(req.MaterialId, 10)
	orderID := strconv.FormatInt(req.ProductionOrderId, 10)
	prefix := "SN"

	sns, err := h.svc.GenerateSerials(ctx, prefix, int(req.Count), materialID, orderID, req.BatchNo)
	if err != nil {
		h.logger.Error("GenerateSerials failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	items := make([]*trace.SerialNumber, len(sns))
	for i, sn := range sns {
		items[i] = modelToProtoSerialNumber(sn)
	}
	return &trace.GenerateSerialResponse{Serials: items}, nil
}

func (h *TraceHandler) CreateDataPoint(ctx context.Context, req *trace.CreateDataPointRequest) (*trace.DataPoint, error) {
	equipmentID := strconv.FormatInt(req.EquipmentId, 10)
	dataTypeStr := protoDataTypeToString(req.DataType)

	var upper, lower *float64
	if req.UpperLimit != "" {
		if v, err := strconv.ParseFloat(req.UpperLimit, 64); err == nil {
			upper = &v
		}
	}
	if req.LowerLimit != "" {
		if v, err := strconv.ParseFloat(req.LowerLimit, 64); err == nil {
			lower = &v
		}
	}

	dp, err := h.svc.CreateDataPoint(ctx, "", req.PointCode, req.PointName, equipmentID,
		dataTypeStr, upper, lower, int(req.CollectIntervalSeconds))
	if err != nil {
		h.logger.Error("CreateDataPoint failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return modelToProtoDataPoint(dp), nil
}

func (h *TraceHandler) ListDataPoints(ctx context.Context, req *trace.ListDataPointsRequest) (*trace.ListDataPointsResponse, error) {
	page := int32(1)
	pageSize := int32(20)
	if req.Page != nil {
		page = req.Page.Page
		pageSize = req.Page.PageSize
	}

	points, total, err := h.svc.ListDataPoints(ctx, req.EquipmentId, dataPointStatusToString(req.Status), int(page), int(pageSize))
	if err != nil {
		h.logger.Error("ListDataPoints failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	items := make([]*trace.DataPoint, len(points))
	for i, p := range points {
		items[i] = modelToProtoDataPoint(p)
	}

	totalPages := int32(0)
	if pageSize > 0 {
		totalPages = (int32(total) + pageSize - 1) / pageSize
	}

	return &trace.ListDataPointsResponse{
		Items: items,
		Page:  protoPageResponse(page, pageSize, int32(total), totalPages),
	}, nil
}

func (h *TraceHandler) CollectData(ctx context.Context, req *trace.CollectDataRequest) (*trace.CollectRecord, error) {
	cr, err := h.svc.CollectData(ctx, "", uint(req.DataPointId), req.Value)
	if err != nil {
		h.logger.Error("CollectData failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return modelToProtoCollectRecord(cr), nil
}

func (h *TraceHandler) ListCollectRecords(ctx context.Context, req *trace.ListCollectRecordsRequest) (*trace.ListCollectRecordsResponse, error) {
	page := int32(1)
	pageSize := int32(20)
	if req.Page != nil {
		page = req.Page.Page
		pageSize = req.Page.PageSize
	}

	records, total, err := h.svc.ListCollectRecords(ctx, req.DataPointId, req.BeginTime, req.EndTime, int(page), int(pageSize))
	if err != nil {
		h.logger.Error("ListCollectRecords failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	items := make([]*trace.CollectRecord, len(records))
	for i, r := range records {
		items[i] = modelToProtoCollectRecord(r)
	}

	totalPages := int32(0)
	if pageSize > 0 {
		totalPages = (int32(total) + pageSize - 1) / pageSize
	}

	return &trace.ListCollectRecordsResponse{
		Items: items,
		Page:  protoPageResponse(page, pageSize, int32(total), totalPages),
	}, nil
}

func (h *TraceHandler) CreateScanLog(ctx context.Context, req *trace.CreateScanLogRequest) (*trace.ScanLog, error) {
	operatorID := strconv.FormatInt(req.OperatorId, 10)
	equipmentID := strconv.FormatInt(req.EquipmentId, 10)
	workstationID := strconv.FormatInt(req.WorkstationId, 10)

	sl, _, err := h.svc.CreateScanLog(ctx, "", req.ScanCode, req.ScanType,
		operatorID, equipmentID, workstationID)
	if err != nil {
		h.logger.Error("CreateScanLog failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return modelToProtoScanLog(sl), nil
}
