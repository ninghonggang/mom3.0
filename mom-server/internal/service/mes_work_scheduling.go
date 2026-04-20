package service

import (
	"context"
	"errors"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// MesWorkSchedulingService 工单排程服务
type MesWorkSchedulingService struct {
	repo       *repository.MesWorkSchedulingRepository
	detailRepo *repository.MesWorkSchedulingDetailRepository
}

func NewMesWorkSchedulingService(
	repo *repository.MesWorkSchedulingRepository,
	detailRepo *repository.MesWorkSchedulingDetailRepository,
) *MesWorkSchedulingService {
	return &MesWorkSchedulingService{repo: repo, detailRepo: detailRepo}
}

// Create 创建工单排程
func (s *MesWorkSchedulingService) Create(ctx context.Context, tenantID int64, userID int64, req *model.MesWorkSchedulingCreateReqVO) (*model.MesWorkScheduling, error) {
	schedulingCode := req.SchedulingCode
	if schedulingCode == "" {
		var err error
		schedulingCode, err = s.repo.GenerateSchedulingCode(ctx, tenantID)
		if err != nil {
			return nil, err
		}
	}

	status := req.Status
	if status == "" {
		status = "PENDING"
	}

	var planDate time.Time
	if req.PlanDate != "" {
		var err error
		planDate, err = time.Parse("2006-01-02", req.PlanDate)
		if err != nil {
			return nil, errors.New("计划日期格式错误，请使用 YYYY-MM-DD")
		}
	}

	m := &model.MesWorkScheduling{
		TenantID:       tenantID,
		PlanNoDay:      req.PlanNoDay,
		SchedulingCode: schedulingCode,
		ProductCode:    req.ProductCode,
		ProductName:    req.ProductName,
		Status:         status,
		Quantity:       req.Quantity,
		WorkMode:       req.WorkMode,
		TaskMode:       req.TaskMode,
		PlanDate:       planDate,
		WorkshopCode:   req.WorkshopCode,
		LineCode:       req.LineCode,
		CreatedBy:      userID,
		UpdatedBy:      userID,
	}

	if err := s.repo.Create(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

// Update 更新工单排程
func (s *MesWorkSchedulingService) Update(ctx context.Context, userID int64, req *model.WorkScheduleUpdateVO) (*model.MesWorkScheduling, error) {
	existing, err := s.repo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, errors.New("工单排程不存在")
	}

	updates := map[string]interface{}{
		"updated_by": userID,
	}
	if req.PlanNoDay != "" {
		updates["plan_no_day"] = req.PlanNoDay
	}
	if req.ProductCode != "" {
		updates["product_code"] = req.ProductCode
	}
	if req.ProductName != "" {
		updates["product_name"] = req.ProductName
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}
	if req.Quantity > 0 {
		updates["quantity"] = req.Quantity
	}
	if req.FinishedQty > 0 {
		updates["finished_qty"] = req.FinishedQty
	}
	if req.WorkMode != "" {
		updates["work_mode"] = req.WorkMode
	}
	if req.TaskMode != "" {
		updates["task_mode"] = req.TaskMode
	}
	if req.PlanDate != "" {
		planDate, err := time.Parse("2006-01-02", req.PlanDate)
		if err != nil {
			return nil, errors.New("计划日期格式错误，请使用 YYYY-MM-DD")
		}
		updates["plan_date"] = planDate
	}
	if req.WorkshopCode != "" {
		updates["workshop_code"] = req.WorkshopCode
	}
	if req.LineCode != "" {
		updates["line_code"] = req.LineCode
	}

	if err := s.repo.Update(ctx, existing.ID, updates); err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, existing.ID)
}

// Delete 删除工单排程（同时删除明细）
func (s *MesWorkSchedulingService) Delete(ctx context.Context, id int64) error {
	if _, err := s.repo.GetByID(ctx, id); err != nil {
		return errors.New("工单排程不存在")
	}
	if err := s.detailRepo.DeleteBySchedulingID(ctx, id); err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}

// Get 获取工单排程（含明细）
func (s *MesWorkSchedulingService) Get(ctx context.Context, id int64) (*model.MesWorkScheduling, error) {
	m, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("工单排程不存在")
	}
	details, err := s.detailRepo.ListBySchedulingID(ctx, id)
	if err != nil {
		return nil, err
	}
	m.Details = details
	return m, nil
}

// Page 分页查询工单排程
func (s *MesWorkSchedulingService) Page(ctx context.Context, tenantID int64, req *model.WorkSchedulePageVO) ([]model.MesWorkScheduling, int64, error) {
	return s.repo.Page(ctx, tenantID, req)
}

// ===================== 明细相关 =====================

// CreateDetail 创建工序排程明细
func (s *MesWorkSchedulingService) CreateDetail(ctx context.Context, tenantID int64, userID int64, req *model.MesWorkSchedulingDetailCreateReqVO) (*model.MesWorkSchedulingDetail, error) {
	if _, err := s.repo.GetByID(ctx, req.SchedulingID); err != nil {
		return nil, errors.New("工单排程不存在")
	}

	status := req.Status
	if status == "" {
		status = "PENDING"
	}

	d := &model.MesWorkSchedulingDetail{
		TenantID:        tenantID,
		SchedulingID:    req.SchedulingID,
		WorkingNode:     req.WorkingNode,
		WorkingName:     req.WorkingName,
		Status:          status,
		EquipmentID:     req.EquipmentID,
		EquipmentCode:   req.EquipmentCode,
		WorkstationID:   req.WorkstationID,
		WorkstationName: req.WorkstationName,
		WorkerID:        req.WorkerID,
		WorkerName:      req.WorkerName,
		PlanQty:         req.PlanQty,
		WorkMinutes:     req.WorkMinutes,
		CreatedBy:       userID,
		UpdatedBy:       userID,
	}

	if err := s.detailRepo.Create(ctx, d); err != nil {
		return nil, err
	}
	return d, nil
}

// UpdateDetail 更新工序排程明细
func (s *MesWorkSchedulingService) UpdateDetail(ctx context.Context, userID int64, req *model.MesWorkSchedulingDetailUpdateReqVO) (*model.MesWorkSchedulingDetail, error) {
	existing, err := s.detailRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, errors.New("工序排程明细不存在")
	}

	updates := map[string]interface{}{"updated_by": userID}
	if req.WorkingNode != "" {
		updates["working_node"] = req.WorkingNode
	}
	if req.WorkingName != "" {
		updates["working_name"] = req.WorkingName
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}
	if req.EquipmentID > 0 {
		updates["equipment_id"] = req.EquipmentID
	}
	if req.EquipmentCode != "" {
		updates["equipment_code"] = req.EquipmentCode
	}
	if req.WorkstationID > 0 {
		updates["workstation_id"] = req.WorkstationID
	}
	if req.WorkstationName != "" {
		updates["workstation_name"] = req.WorkstationName
	}
	if req.WorkerID > 0 {
		updates["worker_id"] = req.WorkerID
	}
	if req.WorkerName != "" {
		updates["worker_name"] = req.WorkerName
	}
	if req.PlanQty > 0 {
		updates["plan_qty"] = req.PlanQty
	}
	if req.FinishedQty > 0 {
		updates["finished_qty"] = req.FinishedQty
	}
	if req.WorkMinutes > 0 {
		updates["work_minutes"] = req.WorkMinutes
	}

	if err := s.detailRepo.Update(ctx, existing.ID, updates); err != nil {
		return nil, err
	}
	return s.detailRepo.GetByID(ctx, existing.ID)
}

// DeleteDetail 删除工序排程明细
func (s *MesWorkSchedulingService) DeleteDetail(ctx context.Context, id int64) error {
	if _, err := s.detailRepo.GetByID(ctx, id); err != nil {
		return errors.New("工序排程明细不存在")
	}
	return s.detailRepo.Delete(ctx, id)
}

// GetDetail 获取工序排程明细
func (s *MesWorkSchedulingService) GetDetail(ctx context.Context, id int64) (*model.MesWorkSchedulingDetail, error) {
	d, err := s.detailRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("工序排程明细不存在")
	}
	return d, nil
}

// PageDetail 分页查询工序明细
func (s *MesWorkSchedulingService) PageDetail(ctx context.Context, tenantID int64, req *model.MesWorkSchedulingDetailPageReqVO) ([]model.MesWorkSchedulingDetail, int64, error) {
	return s.detailRepo.Page(ctx, tenantID, req)
}

// ListDetailBySchedulingID 根据排程ID列出明细
func (s *MesWorkSchedulingService) ListDetailBySchedulingID(ctx context.Context, schedulingID int64) ([]model.MesWorkSchedulingDetail, error) {
	return s.detailRepo.ListBySchedulingID(ctx, schedulingID)
}

// StartDetail 工序开工
func (s *MesWorkSchedulingService) StartDetail(ctx context.Context, id int64, userID int64) error {
	d, err := s.detailRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("工序排程明细不存在")
	}
	if d.Status != "PENDING" && d.Status != "PAUSED" {
		return errors.New("只有待开工或暂停状态的工序可以开工")
	}
	now := time.Now()
	updates := map[string]interface{}{
		"start_time": &now,
		"updated_by": userID,
	}
	return s.detailRepo.UpdateStatus(ctx, id, "IN_PROGRESS", updates)
}

// PauseDetail 工序暂停
func (s *MesWorkSchedulingService) PauseDetail(ctx context.Context, id int64, userID int64) error {
	d, err := s.detailRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("工序排程明细不存在")
	}
	if d.Status != "IN_PROGRESS" {
		return errors.New("只有进行中的工序可以暂停")
	}
	return s.detailRepo.UpdateStatus(ctx, id, "PAUSED", map[string]interface{}{"updated_by": userID})
}

// ResumeDetail 工序恢复
func (s *MesWorkSchedulingService) ResumeDetail(ctx context.Context, id int64, userID int64) error {
	d, err := s.detailRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("工序排程明细不存在")
	}
	if d.Status != "PAUSED" {
		return errors.New("只有暂停状态的工序可以恢复")
	}
	return s.detailRepo.UpdateStatus(ctx, id, "IN_PROGRESS", map[string]interface{}{"updated_by": userID})
}

// CompleteDetail 工序完工
func (s *MesWorkSchedulingService) CompleteDetail(ctx context.Context, id int64, userID int64) error {
	d, err := s.detailRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("工序排程明细不存在")
	}
	if d.Status != "IN_PROGRESS" {
		return errors.New("只有进行中的工序可以完工")
	}
	now := time.Now()
	updates := map[string]interface{}{
		"end_time":   &now,
		"updated_by": userID,
	}
	return s.detailRepo.UpdateStatus(ctx, id, "COMPLETED", updates)
}

// ReportDetail 工序报工
func (s *MesWorkSchedulingService) ReportDetail(ctx context.Context, id int64, userID int64, req *model.MesWorkSchedulingDetailReportReqVO) error {
	if _, err := s.detailRepo.GetByID(ctx, id); err != nil {
		return errors.New("工序排程明细不存在")
	}
	updates := map[string]interface{}{
		"finished_qty": req.FinishedQty,
		"updated_by":   userID,
	}
	return s.detailRepo.Update(ctx, id, updates)
}

// BindEquipment 绑定设备
func (s *MesWorkSchedulingService) BindEquipment(ctx context.Context, userID int64, req *model.MesWorkSchedulingDetailBindEquipmentReqVO) error {
	if _, err := s.detailRepo.GetByID(ctx, req.ID); err != nil {
		return errors.New("工序排程明细不存在")
	}
	updates := map[string]interface{}{
		"equipment_id":   req.EquipmentID,
		"equipment_code": req.EquipmentCode,
		"updated_by":     userID,
	}
	return s.detailRepo.Update(ctx, req.ID, updates)
}

// BindWorker 绑定人员
func (s *MesWorkSchedulingService) BindWorker(ctx context.Context, userID int64, req *model.MesWorkSchedulingDetailBindWorkerReqVO) error {
	if _, err := s.detailRepo.GetByID(ctx, req.ID); err != nil {
		return errors.New("工序排程明细不存在")
	}
	updates := map[string]interface{}{
		"worker_id":   req.WorkerID,
		"worker_name": req.WorkerName,
		"updated_by":  userID,
	}
	return s.detailRepo.Update(ctx, req.ID, updates)
}
