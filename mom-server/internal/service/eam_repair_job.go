package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// EamRepairJobService 维修工单/流程/标准服务
type EamRepairJobService struct {
	jobRepo  *repository.EamRepairJobRepository
	flowRepo *repository.EamRepairFlowRepository
	stdRepo  *repository.EamRepairStdRepository
}

// NewEamRepairJobService 创建维修服务
func NewEamRepairJobService(
	jobRepo *repository.EamRepairJobRepository,
	flowRepo *repository.EamRepairFlowRepository,
	stdRepo *repository.EamRepairStdRepository,
) *EamRepairJobService {
	return &EamRepairJobService{
		jobRepo:  jobRepo,
		flowRepo: flowRepo,
		stdRepo:  stdRepo,
	}
}

// ========== 维修工单 ==========

// CreateJob 创建维修工单
func (s *EamRepairJobService) CreateJob(ctx context.Context, tenantID int64, req *model.EamRepairJobCreateReq) (*model.EamRepairJob, error) {
	job := &model.EamRepairJob{
		TenantID:      tenantID,
		JobCode:       req.JobCode,
		EquipmentID:   req.EquipmentID,
		EquipmentCode: req.EquipmentCode,
		FaultType:     req.FaultType,
		FaultReason:   req.FaultReason,
		FaultDesc:     req.FaultDesc,
		Level:         req.Level,
		Status:        "PENDING",
		ReporterID:    req.ReporterID,
		ReporterName:  req.ReporterName,
	}
	if req.PlanStartTime != "" {
		t := parseRepairTime(req.PlanStartTime)
		job.PlanStartTime = &t
	}
	if req.PlanEndTime != "" {
		t := parseRepairTime(req.PlanEndTime)
		job.PlanEndTime = &t
	}
	if err := s.jobRepo.Create(ctx, job); err != nil {
		return nil, fmt.Errorf("创建维修工单失败: %w", err)
	}
	return job, nil
}

// UpdateJob 更新维修工单
func (s *EamRepairJobService) UpdateJob(ctx context.Context, req *model.EamRepairJobUpdateReq) error {
	updates := map[string]interface{}{
		"fault_type":   req.FaultType,
		"fault_reason": req.FaultReason,
		"fault_desc":   req.FaultDesc,
		"level":        req.Level,
	}
	if req.PlanStartTime != "" {
		updates["plan_start_time"] = parseRepairTime(req.PlanStartTime)
	}
	if req.PlanEndTime != "" {
		updates["plan_end_time"] = parseRepairTime(req.PlanEndTime)
	}
	if err := s.jobRepo.Update(ctx, req.ID, updates); err != nil {
		return fmt.Errorf("更新维修工单失败: %w", err)
	}
	return nil
}

// DeleteJob 删除维修工单
func (s *EamRepairJobService) DeleteJob(ctx context.Context, id int64) error {
	if err := s.jobRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("删除维修工单失败: %w", err)
	}
	return nil
}

// GetJob 获取维修工单详情
func (s *EamRepairJobService) GetJob(ctx context.Context, id int64) (*model.EamRepairJob, error) {
	return s.jobRepo.GetByID(ctx, id)
}

// PageJob 分页查询维修工单
func (s *EamRepairJobService) PageJob(ctx context.Context, tenantID int64, req *model.EamRepairJobPageReq) ([]model.EamRepairJob, int64, error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	filters := map[string]interface{}{"tenant_id": tenantID}
	if req.JobCode != "" {
		filters["job_code"] = req.JobCode
	}
	if req.Status != "" {
		filters["status"] = req.Status
	}
	if req.EquipmentID > 0 {
		filters["equipment_id"] = req.EquipmentID
	}
	return s.jobRepo.List(ctx, offset, pageSize, filters)
}

// AssignJob 派工
func (s *EamRepairJobService) AssignJob(ctx context.Context, req *model.EamRepairJobAssignReq) error {
	updates := map[string]interface{}{
		"assignee_id":   req.AssigneeID,
		"assignee_name": req.AssigneeName,
		"status":        "ASSIGNED",
	}
	if err := s.jobRepo.Update(ctx, req.ID, updates); err != nil {
		return fmt.Errorf("派工失败: %w", err)
	}
	return nil
}

// AcceptJob 接单
func (s *EamRepairJobService) AcceptJob(ctx context.Context, id int64) error {
	now := time.Now()
	updates := map[string]interface{}{
		"status":            "IN_PROGRESS",
		"actual_start_time": now,
	}
	if err := s.jobRepo.Update(ctx, id, updates); err != nil {
		return fmt.Errorf("接单失败: %w", err)
	}
	return nil
}

// CompleteJob 完工
func (s *EamRepairJobService) CompleteJob(ctx context.Context, req *model.EamRepairJobCompleteReq) error {
	now := time.Now()
	updates := map[string]interface{}{
		"status":          "COMPLETED",
		"actual_end_time": now,
		"result":          req.Result,
	}
	if err := s.jobRepo.Update(ctx, req.ID, updates); err != nil {
		return fmt.Errorf("完工失败: %w", err)
	}
	return nil
}

// EvaluateJob 评价
func (s *EamRepairJobService) EvaluateJob(ctx context.Context, req *model.EamRepairJobEvaluateReq) error {
	updates := map[string]interface{}{
		"evaluation": req.Evaluation,
	}
	if err := s.jobRepo.Update(ctx, req.ID, updates); err != nil {
		return fmt.Errorf("评价失败: %w", err)
	}
	return nil
}

// ========== 维修流程 ==========

// CreateFlow 创建维修流程
func (s *EamRepairJobService) CreateFlow(ctx context.Context, tenantID int64, req *model.EamRepairFlowCreateReq) (*model.EamRepairFlow, error) {
	status := req.Status
	if status == "" {
		status = "ACTIVE"
	}
	flow := &model.EamRepairFlow{
		TenantID:  tenantID,
		FlowCode:  req.FlowCode,
		FlowName:  req.FlowName,
		FlowSteps: req.FlowSteps,
		Status:    status,
	}
	if err := s.flowRepo.Create(ctx, flow); err != nil {
		return nil, fmt.Errorf("创建维修流程失败: %w", err)
	}
	return flow, nil
}

// UpdateFlow 更新维修流程
func (s *EamRepairJobService) UpdateFlow(ctx context.Context, req *model.EamRepairFlowUpdateReq) error {
	updates := map[string]interface{}{
		"flow_name":  req.FlowName,
		"flow_steps": req.FlowSteps,
		"status":     req.Status,
	}
	if err := s.flowRepo.Update(ctx, req.ID, updates); err != nil {
		return fmt.Errorf("更新维修流程失败: %w", err)
	}
	return nil
}

// DeleteFlow 删除维修流程
func (s *EamRepairJobService) DeleteFlow(ctx context.Context, id int64) error {
	if err := s.flowRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("删除维修流程失败: %w", err)
	}
	return nil
}

// GetFlow 获取维修流程详情
func (s *EamRepairJobService) GetFlow(ctx context.Context, id int64) (*model.EamRepairFlow, error) {
	return s.flowRepo.GetByID(ctx, id)
}

// PageFlow 分页查询维修流程
func (s *EamRepairJobService) PageFlow(ctx context.Context, tenantID int64, req *model.EamRepairFlowPageReq) ([]model.EamRepairFlow, int64, error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	filters := map[string]interface{}{"tenant_id": tenantID}
	if req.FlowCode != "" {
		filters["flow_code"] = req.FlowCode
	}
	if req.FlowName != "" {
		filters["flow_name"] = req.FlowName
	}
	if req.Status != "" {
		filters["status"] = req.Status
	}
	return s.flowRepo.List(ctx, offset, pageSize, filters)
}

// ========== 维修标准 ==========

// CreateStd 创建维修标准
func (s *EamRepairJobService) CreateStd(ctx context.Context, tenantID int64, req *model.EamRepairStdCreateReq) (*model.EamRepairStd, error) {
	status := req.Status
	if status == "" {
		status = "ACTIVE"
	}
	std := &model.EamRepairStd{
		TenantID:          tenantID,
		StdCode:           req.StdCode,
		StdName:           req.StdName,
		FaultType:         req.FaultType,
		RepairSteps:       req.RepairSteps,
		ToolsRequired:     req.ToolsRequired,
		MaterialsRequired: req.MaterialsRequired,
		StandardHours:     req.StandardHours,
		Status:            status,
	}
	if err := s.stdRepo.Create(ctx, std); err != nil {
		return nil, fmt.Errorf("创建维修标准失败: %w", err)
	}
	return std, nil
}

// UpdateStd 更新维修标准
func (s *EamRepairJobService) UpdateStd(ctx context.Context, req *model.EamRepairStdUpdateReq) error {
	updates := map[string]interface{}{
		"std_name":           req.StdName,
		"fault_type":         req.FaultType,
		"repair_steps":       req.RepairSteps,
		"tools_required":     req.ToolsRequired,
		"materials_required": req.MaterialsRequired,
		"standard_hours":     req.StandardHours,
		"status":             req.Status,
	}
	if err := s.stdRepo.Update(ctx, req.ID, updates); err != nil {
		return fmt.Errorf("更新维修标准失败: %w", err)
	}
	return nil
}

// DeleteStd 删除维修标准
func (s *EamRepairJobService) DeleteStd(ctx context.Context, id int64) error {
	if err := s.stdRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("删除维修标准失败: %w", err)
	}
	return nil
}

// GetStd 获取维修标准详情
func (s *EamRepairJobService) GetStd(ctx context.Context, id int64) (*model.EamRepairStd, error) {
	return s.stdRepo.GetByID(ctx, id)
}

// PageStd 分页查询维修标准
func (s *EamRepairJobService) PageStd(ctx context.Context, tenantID int64, req *model.EamRepairStdPageReq) ([]model.EamRepairStd, int64, error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	filters := map[string]interface{}{"tenant_id": tenantID}
	if req.StdCode != "" {
		filters["std_code"] = req.StdCode
	}
	if req.StdName != "" {
		filters["std_name"] = req.StdName
	}
	if req.FaultType != "" {
		filters["fault_type"] = req.FaultType
	}
	if req.Status != "" {
		filters["status"] = req.Status
	}
	return s.stdRepo.List(ctx, offset, pageSize, filters)
}

// ========== 辅助函数 ==========

func parseRepairTime(s string) time.Time {
	if len(s) >= 10 {
		t, _ := time.Parse("2006-01-02", s[:10])
		return t
	}
	return time.Time{}
}
