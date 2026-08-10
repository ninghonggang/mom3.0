package handler

import (
	"context"
	"strconv"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	eam "github.com/ninghonggang/mom-platform/gen/eam"
	common "github.com/ninghonggang/mom-platform/gen/common"
	"mom-platform/services/eam-service/internal/model"
	"mom-platform/services/eam-service/internal/repository"
	"mom-platform/services/eam-service/internal/service"
	"github.com/ninghonggang/mom-platform/pkg/eventbus"
)

// Handler implements the gRPC EamServiceServer interface.
type Handler struct {
	eam.UnimplementedEamServiceServer
	svc       *service.Service
	logger    *zap.Logger
	eventPub  *eventbus.EventPublisher
}

// New creates a new Handler.
func New(svc *service.Service, logger *zap.Logger, eventPub *eventbus.EventPublisher) *Handler {
	return &Handler{svc: svc, logger: logger, eventPub: eventPub}
}

// ============ Equipment ============

func (h *Handler) CreateEquipment(ctx context.Context, req *eam.CreateEquipmentRequest) (*eam.Equipment, error) {
	h.logger.Info("CreateEquipment", zap.String("code", req.GetEquipmentCode()))

	if req.GetEquipmentCode() == "" {
		return nil, status.Error(codes.InvalidArgument, "equipment_code is required")
	}

	equip := &model.Equipment{
		EquipmentCode: req.GetEquipmentCode(),
		EquipmentName: req.GetEquipmentName(),
		EquipmentType: protoToModelEquipmentType(req.GetEquipmentType()),
		Model:         req.GetModel(),
		Specification: req.GetSpecification(),
		TargetOee:     req.GetTargetOee(),
	}
	if req.GetWorkshopId() > 0 {
		wid := req.GetWorkshopId()
		equip.WorkshopID = &wid
	}
	if req.GetLineId() > 0 {
		lid := req.GetLineId()
		equip.LineID = &lid
	}

	if err := h.svc.CreateEquipment(ctx, equip); err != nil {
		h.logger.Error("create equipment failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "create equipment: %v", err)
	}
	return modelToProtoEquipment(equip), nil
}

func (h *Handler) GetEquipment(ctx context.Context, req *eam.GetEquipmentRequest) (*eam.Equipment, error) {
	h.logger.Info("GetEquipment", zap.Int64("id", req.GetId()))

	equip, err := h.svc.GetEquipment(ctx, req.GetId())
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "equipment %d not found", req.GetId())
		}
		h.logger.Error("get equipment failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "get equipment: %v", err)
	}
	return modelToProtoEquipment(equip), nil
}

func (h *Handler) ListEquipment(ctx context.Context, req *eam.ListEquipmentRequest) (*eam.ListEquipmentResponse, error) {
	h.logger.Info("ListEquipment")

	page := protoToModelPage(req.GetPage())
	filter := repository.EquipmentFilter{
		Type: protoToModelEquipmentType(req.GetEquipmentType()),
		Status: protoToModelEquipmentStatus(req.GetStatus()),
	}
	if req.GetWorkshopId() > 0 {
		wid := req.GetWorkshopId()
		filter.WorkshopID = &wid
	}
	if req.GetLineId() > 0 {
		lid := req.GetLineId()
		filter.LineID = &lid
	}

	items, total, err := h.svc.ListEquipment(ctx, filter, page)
	if err != nil {
		h.logger.Error("list equipment failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "list equipment: %v", err)
	}

	resp := &eam.ListEquipmentResponse{
		Items: make([]*eam.Equipment, 0, len(items)),
		Page:  modelToProtoPage(page, total),
	}
	for i := range items {
		resp.Items = append(resp.Items, modelToProtoEquipment(&items[i]))
	}
	return resp, nil
}

// ============ RepairOrder ============

func (h *Handler) CreateRepairOrder(ctx context.Context, req *eam.CreateRepairOrderRequest) (*eam.RepairOrder, error) {
	h.logger.Info("CreateRepairOrder", zap.Int64("equipment_id", req.GetEquipmentId()))

	if req.GetEquipmentId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "equipment_id is required")
	}

	ro := &model.RepairOrder{
		EquipmentID: req.GetEquipmentId(),
		FaultType:   req.GetFaultType(),
		FaultDesc:   req.GetFaultDesc(),
		Urgency:     protoToModelUrgency(req.GetUrgency()),
	}
	if req.GetReporterId() > 0 {
		rid := req.GetReporterId()
		ro.ReporterID = &rid
	}

	created, err := h.svc.CreateRepairOrder(ctx, ro)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "equipment %d not found", req.GetEquipmentId())
		}
		h.logger.Error("create repair order failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "create repair order: %v", err)
	}
	return modelToProtoRepairOrder(created), nil
}

func (h *Handler) UpdateRepairOrder(ctx context.Context, req *eam.UpdateRepairOrderRequest) (*eam.RepairOrder, error) {
	h.logger.Info("UpdateRepairOrder", zap.Int64("id", req.GetId()), zap.String("status", req.GetStatus().String()))

	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	var repairmanID *int64
	if req.GetRepairmanId() > 0 {
		rid := req.GetRepairmanId()
		repairmanID = &rid
	}

	updated, err := h.svc.UpdateRepairOrder(ctx, req.GetId(), protoToModelRepairStatus(req.GetStatus()), repairmanID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "repair order %d not found", req.GetId())
		}
		h.logger.Error("update repair order failed", zap.Error(err))
		return nil, status.Errorf(codes.FailedPrecondition, "update repair order: %v", err)
	}
	return modelToProtoRepairOrder(updated), nil
}

func (h *Handler) GetRepairOrder(ctx context.Context, req *eam.GetEquipmentRequest) (*eam.RepairOrder, error) {
	h.logger.Info("GetRepairOrder", zap.Int64("id", req.GetId()))

	ro, err := h.svc.GetRepairOrder(ctx, req.GetId())
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "repair order %d not found", req.GetId())
		}
		h.logger.Error("get repair order failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "get repair order: %v", err)
	}
	return modelToProtoRepairOrder(ro), nil
}

func (h *Handler) ListRepairOrders(ctx context.Context, req *eam.ListRepairOrdersRequest) (*eam.ListRepairOrdersResponse, error) {
	h.logger.Info("ListRepairOrders")

	page := protoToModelPage(req.GetPage())
	filter := repository.RepairOrderFilter{
		Status: protoToModelRepairStatus(req.GetStatus()),
	}
	if req.GetEquipmentId() > 0 {
		eid := req.GetEquipmentId()
		filter.EquipmentID = &eid
	}

	items, total, err := h.svc.ListRepairOrders(ctx, filter, page)
	if err != nil {
		h.logger.Error("list repair orders failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "list repair orders: %v", err)
	}

	resp := &eam.ListRepairOrdersResponse{
		Items: make([]*eam.RepairOrder, 0, len(items)),
		Page:  modelToProtoPage(page, total),
	}
	for i := range items {
		resp.Items = append(resp.Items, modelToProtoRepairOrder(&items[i]))
	}
	return resp, nil
}

// ============ OEE ============

// SaveOee 接收某设备某日的 OEE 三要素并落库，OEE 由 service 层按 A×P×Q 计算。
// 同一 (设备, 日期) 重复上报为幂等覆盖。
func (h *Handler) SaveOee(ctx context.Context, req *eam.SaveOeeRequest) (*eam.EquipmentOee, error) {
	h.logger.Info("SaveOee",
		zap.Int64("equipment_id", req.GetEquipmentId()),
		zap.String("calc_date", req.GetCalcDate()),
	)

	if req.GetEquipmentId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "equipment_id is required")
	}

	availability, err := parseRatio(req.GetAvailability(), "availability")
	if err != nil {
		return nil, err
	}
	performance, err := parseRatio(req.GetPerformance(), "performance")
	if err != nil {
		return nil, err
	}
	quality, err := parseRatio(req.GetQuality(), "quality")
	if err != nil {
		return nil, err
	}

	calcDate := req.GetCalcDate()
	if calcDate == "" {
		calcDate = time.Now().Format("2006-01-02")
	}
	if _, err := time.Parse("2006-01-02", calcDate); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "calc_date must be YYYY-MM-DD, got %q", calcDate)
	}

	saved, err := h.svc.SaveOee(ctx, &model.EquipmentOee{
		EquipmentID:  req.GetEquipmentId(),
		CalcDate:     calcDate,
		Availability: availability,
		Performance:  performance,
		Quality:      quality,
	})
	if err != nil {
		h.logger.Error("save oee failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "save oee: %v", err)
	}
	return modelToProtoOee(saved), nil
}

// parseRatio 解析 0~1 之间的比率字符串，空串按 0 处理，越界或非法直接拒绝，
// 避免脏数据把 OEE 算成大于 1 的荒谬值。
func parseRatio(s, field string) (float64, error) {
	if s == "" {
		return 0, nil
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, status.Errorf(codes.InvalidArgument, "%s must be a number, got %q", field, s)
	}
	if v < 0 || v > 1 {
		return 0, status.Errorf(codes.InvalidArgument, "%s must be within [0,1], got %v", field, v)
	}
	return v, nil
}

func (h *Handler) ListOee(ctx context.Context, req *eam.ListOeeRequest) (*eam.ListOeeResponse, error) {
	h.logger.Info("ListOee", zap.Int64("equipment_id", req.GetEquipmentId()))

	filter := repository.OeeFilter{
		EquipmentID: req.GetEquipmentId(),
		BeginDate:   req.GetBeginDate(),
		EndDate:     req.GetEndDate(),
	}

	items, err := h.svc.ListOee(ctx, filter)
	if err != nil {
		h.logger.Error("list oee failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "list oee: %v", err)
	}

	resp := &eam.ListOeeResponse{
		Items: make([]*eam.EquipmentOee, 0, len(items)),
	}
	for i := range items {
		resp.Items = append(resp.Items, modelToProtoOee(&items[i]))
	}
	return resp, nil
}

// ============ MaintenancePlan ============

func (h *Handler) CreateMaintenancePlan(ctx context.Context, req *eam.CreateMaintenancePlanRequest) (*eam.MaintenancePlan, error) {
	h.logger.Info("CreateMaintenancePlan", zap.Int64("equipment_id", req.GetEquipmentId()))

	if req.GetEquipmentId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "equipment_id is required")
	}

	plan := &model.MaintenancePlan{
		EquipmentID:         req.GetEquipmentId(),
		MaintenanceType:     req.GetMaintenanceType(),
		CycleDays:           req.GetCycleDays(),
		NextMaintenanceDate: req.GetNextMaintenanceDate(),
	}

	created, err := h.svc.CreateMaintenancePlan(ctx, plan)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "equipment %d not found", req.GetEquipmentId())
		}
		h.logger.Error("create maintenance plan failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "create maintenance plan: %v", err)
	}
	return modelToProtoMaintenancePlan(created), nil
}

func (h *Handler) ListMaintenancePlans(ctx context.Context, req *eam.ListMaintenancePlansRequest) (*eam.ListMaintenancePlansResponse, error) {
	h.logger.Info("ListMaintenancePlans")

	page := protoToModelPage(req.GetPage())
	filter := repository.MaintenancePlanFilter{}
	if req.GetEquipmentId() > 0 {
		eid := req.GetEquipmentId()
		filter.EquipmentID = &eid
	}

	items, total, err := h.svc.ListMaintenancePlans(ctx, filter, page)
	if err != nil {
		h.logger.Error("list maintenance plans failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "list maintenance plans: %v", err)
	}

	resp := &eam.ListMaintenancePlansResponse{
		Items: make([]*eam.MaintenancePlan, 0, len(items)),
		Page:  modelToProtoPage(page, total),
	}
	for i := range items {
		resp.Items = append(resp.Items, modelToProtoMaintenancePlan(&items[i]))
	}
	return resp, nil
}

// ============ EquipmentCheck ============

func (h *Handler) CreateCheck(ctx context.Context, req *eam.CreateCheckRequest) (*eam.EquipmentCheck, error) {
	h.logger.Info("CreateCheck", zap.Int64("equipment_id", req.GetEquipmentId()))

	if req.GetEquipmentId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "equipment_id is required")
	}

	check := &model.EquipmentCheck{
		EquipmentID: req.GetEquipmentId(),
		Result:      protoToModelCheckResult(req.GetResult()),
		Remark:      req.GetRemark(),
	}
	if req.GetCheckStdId() > 0 {
		csid := req.GetCheckStdId()
		check.CheckStdID = &csid
	}
	if req.GetCheckerId() > 0 {
		cid := req.GetCheckerId()
		check.CheckerID = &cid
	}

	created, err := h.svc.CreateCheck(ctx, check)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "equipment %d not found", req.GetEquipmentId())
		}
		h.logger.Error("create check failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "create check: %v", err)
	}
	return modelToProtoCheck(created), nil
}

func (h *Handler) ListChecks(ctx context.Context, req *eam.ListChecksRequest) (*eam.ListChecksResponse, error) {
	h.logger.Info("ListChecks")

	page := protoToModelPage(req.GetPage())
	filter := repository.CheckFilter{}
	if req.GetEquipmentId() > 0 {
		eid := req.GetEquipmentId()
		filter.EquipmentID = &eid
	}
	if req.GetBeginTime() > 0 {
		bt := req.GetBeginTime()
		filter.BeginTime = &bt
	}
	if req.GetEndTime() > 0 {
		et := req.GetEndTime()
		filter.EndTime = &et
	}

	items, total, err := h.svc.ListChecks(ctx, filter, page)
	if err != nil {
		h.logger.Error("list checks failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "list checks: %v", err)
	}

	resp := &eam.ListChecksResponse{
		Items: make([]*eam.EquipmentCheck, 0, len(items)),
		Page:  modelToProtoPage(page, total),
	}
	for i := range items {
		resp.Items = append(resp.Items, modelToProtoCheck(&items[i]))
	}
	return resp, nil
}

// ============ EquipmentDowntime ============

func (h *Handler) StartDowntime(ctx context.Context, req *eam.StartDowntimeRequest) (*eam.EquipmentDowntime, error) {
	h.logger.Info("StartDowntime", zap.Int64("equipment_id", req.GetEquipmentId()))

	if req.GetEquipmentId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "equipment_id is required")
	}

	dt, err := h.svc.StartDowntime(ctx, req.GetEquipmentId(), protoToModelDowntimeType(req.GetDowntimeType()), req.GetReason())
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "equipment %d not found", req.GetEquipmentId())
		}
		h.logger.Error("start downtime failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "start downtime: %v", err)
	}

	// Publish downtime start event.
	if h.eventPub != nil {
		if pubErr := h.eventPub.Publish(ctx, eventbus.SubjectEAMDowntimeStart, dt); pubErr != nil {
			h.logger.Error("failed to publish downtime start event", zap.Error(pubErr))
		}
	}

	return modelToProtoDowntime(dt), nil
}

func (h *Handler) ResolveDowntime(ctx context.Context, req *eam.ResolveDowntimeRequest) (*eam.EquipmentDowntime, error) {
	h.logger.Info("ResolveDowntime",
		zap.Int64("id", req.GetId()),
		zap.String("resolver_id", req.GetResolverId()),
	)

	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	dt, err := h.svc.ResolveDowntime(ctx, req.GetId(), req.GetResolution(), req.GetResolverId())
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "downtime %d not found", req.GetId())
		}
		h.logger.Error("resolve downtime failed", zap.Error(err))
		return nil, status.Errorf(codes.FailedPrecondition, "resolve downtime: %v", err)
	}

	// Publish downtime resolve event.
	if h.eventPub != nil {
		if pubErr := h.eventPub.Publish(ctx, eventbus.SubjectEAMDowntimeResolve, dt); pubErr != nil {
			h.logger.Error("failed to publish downtime resolve event", zap.Error(pubErr))
		}
	}

	return modelToProtoDowntime(dt), nil
}

func (h *Handler) ListDowntimes(ctx context.Context, req *eam.ListDowntimesRequest) (*eam.ListDowntimesResponse, error) {
	h.logger.Info("ListDowntimes")

	page := protoToModelPage(req.GetPage())
	filter := repository.DowntimeFilter{
		Status: protoToModelDowntimeStatus(req.GetStatus()),
	}
	if req.GetEquipmentId() > 0 {
		eid := req.GetEquipmentId()
		filter.EquipmentID = &eid
	}

	items, total, err := h.svc.ListDowntimes(ctx, filter, page)
	if err != nil {
		h.logger.Error("list downtimes failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "list downtimes: %v", err)
	}

	resp := &eam.ListDowntimesResponse{
		Items: make([]*eam.EquipmentDowntime, 0, len(items)),
		Page:  modelToProtoPage(page, total),
	}
	for i := range items {
		resp.Items = append(resp.Items, modelToProtoDowntime(&items[i]))
	}
	return resp, nil
}

// ============ Conversion helpers ============

func protoToModelPage(p *common.PageRequest) model.PageQuery {
	if p == nil {
		return model.NewPageQuery(1, 20)
	}
	return model.NewPageQuery(p.GetPage(), p.GetPageSize())
}

func modelToProtoPage(q model.PageQuery, total int64) *common.PageResponse {
	return &common.PageResponse{
		Page:       q.Page,
		PageSize:   q.PageSize,
		Total:      int32(total),
		TotalPages: model.CalcTotalPages(total, q.PageSize),
	}
}

func protoToModelEquipmentType(t eam.EquipmentType) model.EquipmentType {
	switch t {
	case eam.EquipmentType_EQUIPMENT_TYPE_MACHINE:
		return model.EquipmentTypeMachine
	case eam.EquipmentType_EQUIPMENT_TYPE_MOLD:
		return model.EquipmentTypeMold
	case eam.EquipmentType_EQUIPMENT_TYPE_INSTRUMENT:
		return model.EquipmentTypeInstrument
	default:
		return ""
	}
}

func protoToModelEquipmentStatus(s eam.EquipmentStatus) model.EquipmentStatus {
	switch s {
	case eam.EquipmentStatus_EQUIPMENT_STATUS_RUNNING:
		return model.EquipmentStatusRunning
	case eam.EquipmentStatus_EQUIPMENT_STATUS_IDLE:
		return model.EquipmentStatusIdle
	case eam.EquipmentStatus_EQUIPMENT_STATUS_MAINTENANCE:
		return model.EquipmentStatusMaintenance
	case eam.EquipmentStatus_EQUIPMENT_STATUS_REPAIR:
		return model.EquipmentStatusRepair
	case eam.EquipmentStatus_EQUIPMENT_STATUS_SCRAPPED:
		return model.EquipmentStatusScrapped
	default:
		return ""
	}
}

func modelToProtoEquipment(e *model.Equipment) *eam.Equipment {
	return &eam.Equipment{
		Id:              e.ID,
		TenantId:        e.TenantID,
		EquipmentCode:   e.EquipmentCode,
		EquipmentName:   e.EquipmentName,
		EquipmentType:   modelToProtoEquipmentType(e.EquipmentType),
		EquipmentClassId: derefInt64(e.EquipmentClassID),
		Model:           e.Model,
		Specification:   e.Specification,
		ManufacturerId:  derefInt64(e.ManufacturerID),
		SupplierId:      derefInt64(e.SupplierID),
		SerialNo:        e.SerialNo,
		WorkshopId:      derefInt64(e.WorkshopID),
		LineId:          derefInt64(e.LineID),
		LocationId:      derefInt64(e.LocationID),
		Status:          modelToProtoEquipmentStatus(e.Status),
		TargetOee:       e.TargetOee,
		PurchaseAmount:  e.PurchaseAmount,
		ServiceLife:     e.ServiceLife,
		InstallDate:     formatTimeStr(e.InstallDate),
	}
}

func modelToProtoEquipmentType(t model.EquipmentType) eam.EquipmentType {
	switch t {
	case model.EquipmentTypeMachine:
		return eam.EquipmentType_EQUIPMENT_TYPE_MACHINE
	case model.EquipmentTypeMold:
		return eam.EquipmentType_EQUIPMENT_TYPE_MOLD
	case model.EquipmentTypeInstrument:
		return eam.EquipmentType_EQUIPMENT_TYPE_INSTRUMENT
	default:
		return eam.EquipmentType_EQUIPMENT_TYPE_UNSPECIFIED
	}
}

func modelToProtoEquipmentStatus(s model.EquipmentStatus) eam.EquipmentStatus {
	switch s {
	case model.EquipmentStatusRunning:
		return eam.EquipmentStatus_EQUIPMENT_STATUS_RUNNING
	case model.EquipmentStatusIdle:
		return eam.EquipmentStatus_EQUIPMENT_STATUS_IDLE
	case model.EquipmentStatusMaintenance:
		return eam.EquipmentStatus_EQUIPMENT_STATUS_MAINTENANCE
	case model.EquipmentStatusRepair:
		return eam.EquipmentStatus_EQUIPMENT_STATUS_REPAIR
	case model.EquipmentStatusScrapped:
		return eam.EquipmentStatus_EQUIPMENT_STATUS_SCRAPPED
	default:
		return eam.EquipmentStatus_EQUIPMENT_STATUS_UNSPECIFIED
	}
}

func protoToModelUrgency(u eam.Urgency) model.Urgency {
	switch u {
	case eam.Urgency_URGENCY_NORMAL:
		return model.UrgencyNormal
	case eam.Urgency_URGENCY_URGENT:
		return model.UrgencyUrgent
	case eam.Urgency_URGENCY_EMERGENCY:
		return model.UrgencyEmergency
	default:
		return model.UrgencyNormal
	}
}

func protoToModelRepairStatus(s eam.RepairOrderStatus) model.RepairOrderStatus {
	switch s {
	case eam.RepairOrderStatus_REPAIR_STATUS_REPORTED:
		return model.RepairStatusReported
	case eam.RepairOrderStatus_REPAIR_STATUS_ASSIGNED:
		return model.RepairStatusAssigned
	case eam.RepairOrderStatus_REPAIR_STATUS_IN_PROGRESS:
		return model.RepairStatusInProgress
	case eam.RepairOrderStatus_REPAIR_STATUS_PENDING_PARTS:
		return model.RepairStatusPendingParts
	case eam.RepairOrderStatus_REPAIR_STATUS_COMPLETED:
		return model.RepairStatusCompleted
	case eam.RepairOrderStatus_REPAIR_STATUS_VERIFIED:
		return model.RepairStatusVerified
	case eam.RepairOrderStatus_REPAIR_STATUS_CANCELLED:
		return model.RepairStatusCancelled
	default:
		return ""
	}
}

func modelToProtoRepairOrder(ro *model.RepairOrder) *eam.RepairOrder {
	resp := &eam.RepairOrder{
		Id:          ro.ID,
		TenantId:    ro.TenantID,
		RepairNo:    ro.RepairNo,
		EquipmentId: ro.EquipmentID,
		FaultType:   ro.FaultType,
		FaultDesc:   ro.FaultDesc,
		Urgency:     modelToProtoUrgency(ro.Urgency),
		ReporterId:  derefInt64(ro.ReporterID),
		RepairmanId: derefInt64(ro.RepairmanID),
		Status:      modelToProtoRepairStatus(ro.Status),
		ReportedAt:  ro.ReportedAt.Unix(),
	}
	if ro.CompletedAt != nil {
		resp.CompletedAt = ro.CompletedAt.Unix()
	}
	return resp
}

func modelToProtoUrgency(u model.Urgency) eam.Urgency {
	switch u {
	case model.UrgencyNormal:
		return eam.Urgency_URGENCY_NORMAL
	case model.UrgencyUrgent:
		return eam.Urgency_URGENCY_URGENT
	case model.UrgencyEmergency:
		return eam.Urgency_URGENCY_EMERGENCY
	default:
		return eam.Urgency_URGENCY_UNSPECIFIED
	}
}

func modelToProtoRepairStatus(s model.RepairOrderStatus) eam.RepairOrderStatus {
	switch s {
	case model.RepairStatusReported:
		return eam.RepairOrderStatus_REPAIR_STATUS_REPORTED
	case model.RepairStatusAssigned:
		return eam.RepairOrderStatus_REPAIR_STATUS_ASSIGNED
	case model.RepairStatusInProgress:
		return eam.RepairOrderStatus_REPAIR_STATUS_IN_PROGRESS
	case model.RepairStatusPendingParts:
		return eam.RepairOrderStatus_REPAIR_STATUS_PENDING_PARTS
	case model.RepairStatusCompleted:
		return eam.RepairOrderStatus_REPAIR_STATUS_COMPLETED
	case model.RepairStatusVerified:
		return eam.RepairOrderStatus_REPAIR_STATUS_VERIFIED
	case model.RepairStatusCancelled:
		return eam.RepairOrderStatus_REPAIR_STATUS_CANCELLED
	default:
		return eam.RepairOrderStatus_REPAIR_STATUS_UNSPECIFIED
	}
}

func modelToProtoOee(o *model.EquipmentOee) *eam.EquipmentOee {
	return &eam.EquipmentOee{
		Id:          o.ID,
		EquipmentId: o.EquipmentID,
		CalcDate:    o.CalcDate,
		Availability: formatFloat(o.Availability),
		Performance:  formatFloat(o.Performance),
		Quality:      formatFloat(o.Quality),
		Oee:          formatFloat(o.OEE),
	}
}

func modelToProtoMaintenancePlan(p *model.MaintenancePlan) *eam.MaintenancePlan {
	return &eam.MaintenancePlan{
		Id:                  p.ID,
		EquipmentId:         p.EquipmentID,
		PlanNo:              p.PlanNo,
		MaintenanceType:     p.MaintenanceType,
		CycleDays:           p.CycleDays,
		NextMaintenanceDate: p.NextMaintenanceDate,
		Status:              modelToProtoMaintenanceStatus(p.Status),
	}
}

func modelToProtoMaintenanceStatus(s model.MaintenancePlanStatus) eam.MaintenancePlanStatus {
	switch s {
	case model.MaintenanceStatusScheduled:
		return eam.MaintenancePlanStatus_MAINTENANCE_STATUS_SCHEDULED
	case model.MaintenanceStatusInProgress:
		return eam.MaintenancePlanStatus_MAINTENANCE_STATUS_IN_PROGRESS
	case model.MaintenanceStatusCompleted:
		return eam.MaintenancePlanStatus_MAINTENANCE_STATUS_COMPLETED
	case model.MaintenanceStatusSkipped:
		return eam.MaintenancePlanStatus_MAINTENANCE_STATUS_SKIPPED
	default:
		return eam.MaintenancePlanStatus_MAINTENANCE_STATUS_UNSPECIFIED
	}
}

func protoToModelCheckResult(r eam.CheckResult) model.CheckResult {
	switch r {
	case eam.CheckResult_CHECK_RESULT_OK:
		return model.CheckResultOK
	case eam.CheckResult_CHECK_RESULT_NG:
		return model.CheckResultNG
	case eam.CheckResult_CHECK_RESULT_NA:
		return model.CheckResultNA
	default:
		return ""
	}
}

func modelToProtoCheck(c *model.EquipmentCheck) *eam.EquipmentCheck {
	return &eam.EquipmentCheck{
		Id:          c.ID,
		EquipmentId: c.EquipmentID,
		CheckNo:     c.CheckNo,
		CheckStdId:  derefInt64(c.CheckStdID),
		CheckerId:   derefInt64(c.CheckerID),
		CheckTime:   c.CheckTime.Unix(),
		Result:      modelToProtoCheckResult(c.Result),
		Remark:      c.Remark,
	}
}

func modelToProtoCheckResult(r model.CheckResult) eam.CheckResult {
	switch r {
	case model.CheckResultOK:
		return eam.CheckResult_CHECK_RESULT_OK
	case model.CheckResultNG:
		return eam.CheckResult_CHECK_RESULT_NG
	case model.CheckResultNA:
		return eam.CheckResult_CHECK_RESULT_NA
	default:
		return eam.CheckResult_CHECK_RESULT_UNSPECIFIED
	}
}

func protoToModelDowntimeType(t eam.DowntimeType) model.DowntimeType {
	switch t {
	case eam.DowntimeType_DOWNTIME_TYPE_UNPLANNED:
		return model.DowntimeTypeUnplanned
	case eam.DowntimeType_DOWNTIME_TYPE_PLANNED:
		return model.DowntimeTypePlanned
	default:
		return model.DowntimeTypeUnplanned
	}
}

func protoToModelDowntimeStatus(s eam.DowntimeStatus) model.DowntimeStatus {
	switch s {
	case eam.DowntimeStatus_DOWNTIME_STATUS_ACTIVE:
		return model.DowntimeStatusActive
	case eam.DowntimeStatus_DOWNTIME_STATUS_RESOLVED:
		return model.DowntimeStatusResolved
	case eam.DowntimeStatus_DOWNTIME_STATUS_PLANNED:
		return model.DowntimeStatusPlanned
	default:
		return ""
	}
}

func modelToProtoDowntime(d *model.EquipmentDowntime) *eam.EquipmentDowntime {
	resp := &eam.EquipmentDowntime{
		Id:              d.ID,
		EquipmentId:     d.EquipmentID,
		DowntimeType:    modelToProtoDowntimeType(d.DowntimeType),
		StartTime:       d.StartTime.Unix(),
		DurationSeconds: d.DurationSeconds,
		Reason:          d.Reason,
		Status:          modelToProtoDowntimeStatus(d.Status),
		Resolution:      d.Resolution,
		ResolverId:      d.ResolverID,
	}
	if d.EndTime != nil {
		resp.EndTime = d.EndTime.Unix()
	}
	return resp
}

func modelToProtoDowntimeType(t model.DowntimeType) eam.DowntimeType {
	switch t {
	case model.DowntimeTypeUnplanned:
		return eam.DowntimeType_DOWNTIME_TYPE_UNPLANNED
	case model.DowntimeTypePlanned:
		return eam.DowntimeType_DOWNTIME_TYPE_PLANNED
	default:
		return eam.DowntimeType_DOWNTIME_TYPE_UNSPECIFIED
	}
}

func modelToProtoDowntimeStatus(s model.DowntimeStatus) eam.DowntimeStatus {
	switch s {
	case model.DowntimeStatusActive:
		return eam.DowntimeStatus_DOWNTIME_STATUS_ACTIVE
	case model.DowntimeStatusResolved:
		return eam.DowntimeStatus_DOWNTIME_STATUS_RESOLVED
	case model.DowntimeStatusPlanned:
		return eam.DowntimeStatus_DOWNTIME_STATUS_PLANNED
	default:
		return eam.DowntimeStatus_DOWNTIME_STATUS_UNSPECIFIED
	}
}

// ============ Utility ============

func derefInt64(p *int64) int64 {
	if p == nil {
		return 0
	}
	return *p
}

func formatTimeStr(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02")
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', 4, 64)
}
