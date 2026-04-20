package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"time"
)

// ContainerLifecycleService 容器生命周期服务
type ContainerLifecycleService struct {
	lifecycleRepo *repository.ContainerLifecycleRepository
	maintenanceRepo *repository.ContainerMaintenanceRepository
	containerRepo *repository.ContainerRepository
}

func NewContainerLifecycleService(
	lifecycleRepo *repository.ContainerLifecycleRepository,
	maintenanceRepo *repository.ContainerMaintenanceRepository,
	containerRepo *repository.ContainerRepository,
) *ContainerLifecycleService {
	return &ContainerLifecycleService{
		lifecycleRepo: lifecycleRepo,
		maintenanceRepo: maintenanceRepo,
		containerRepo: containerRepo,
	}
}

// ListContainerLifecycles 分页查询容器生命周期记录
func (s *ContainerLifecycleService) ListContainerLifecycles(ctx context.Context, tenantID int64, query *model.ContainerLifecycleQuery) ([]model.ContainerLifecycle, int64, error) {
	return s.lifecycleRepo.List(ctx, tenantID, query)
}

// GetContainerLifecycle 获取单条生命周期记录
func (s *ContainerLifecycleService) GetContainerLifecycle(ctx context.Context, id uint) (*model.ContainerLifecycle, error) {
	return s.lifecycleRepo.GetByID(ctx, id)
}

// InitializeContainer 初始化容器
func (s *ContainerLifecycleService) InitializeContainer(ctx context.Context, tenantID int64, req *model.InitializeContainerRequest) error {
	// 检查容器是否存在
	container, err := s.containerRepo.GetByID(ctx, uint(req.ContainerID))
	if err != nil {
		return fmt.Errorf("容器不存在: %w", err)
	}

	// 检查是否已有初始化记录
	latest, _ := s.lifecycleRepo.GetLatestByContainerID(ctx, req.ContainerID)
	if latest != nil && latest.EventType == model.ContainerEventInitialize {
		return fmt.Errorf("容器已初始化，不能重复初始化")
	}

	// 创建生命周期记录
	now := time.Now()
	lifecycle := &model.ContainerLifecycle{
		TenantID:      tenantID,
		ContainerID:    req.ContainerID,
		ContainerCode: container.ContainerCode,
		EventType:     model.ContainerEventInitialize,
		EventDate:     now,
		OperatorID:    req.OperatorID,
		OperatorName:  req.OperatorName,
		LocationID:    req.LocationID,
		LocationName:  req.LocationName,
		Status:        model.LifecycleStatusBefore,
		Remark:        req.Remark,
	}

	if err := s.lifecycleRepo.Create(ctx, lifecycle); err != nil {
		return fmt.Errorf("创建生命周期记录失败: %w", err)
	}

	// 更新容器状态为在役
	updates := map[string]interface{}{
		"status": model.ContainerStatusInUse,
	}
	if req.LocationName != "" {
		updates["current_location"] = req.LocationName
	}

	if err := s.containerRepo.Update(ctx, uint(req.ContainerID), updates); err != nil {
		return fmt.Errorf("更新容器状态失败: %w", err)
	}

	return nil
}

// RecordMaintenance 记录维修
func (s *ContainerLifecycleService) RecordMaintenance(ctx context.Context, tenantID int64, req *model.RecordMaintenanceRequest) error {
	// 检查容器是否存在
	container, err := s.containerRepo.GetByID(ctx, uint(req.ContainerID))
	if err != nil {
		return fmt.Errorf("容器不存在: %w", err)
	}

	// 创建维修记录
	maintenance := &model.ContainerMaintenance{
		TenantID:           tenantID,
		ContainerID:         req.ContainerID,
		ContainerCode:      container.ContainerCode,
		MaintenanceType:    req.MaintenanceType,
		MaintenanceDate:    req.MaintenanceDate,
		TechnicianID:       req.TechnicianID,
		TechnicianName:     req.TechnicianName,
		FaultDescription:   req.FaultDescription,
		MaintenanceContent:  req.MaintenanceContent,
		SparePartsUsed:     req.SparePartsUsed,
		Cost:               req.Cost,
		Status:             model.MaintenanceStatusPending,
		Remark:             req.Remark,
	}

	if err := s.maintenanceRepo.Create(ctx, maintenance); err != nil {
		return fmt.Errorf("创建维修记录失败: %w", err)
	}

	// 创建生命周期记录（维修事件）
	lifecycle := &model.ContainerLifecycle{
		TenantID:      tenantID,
		ContainerID:    req.ContainerID,
		ContainerCode: container.ContainerCode,
		EventType:     model.ContainerEventMaintenance,
		EventDate:     req.MaintenanceDate,
		OperatorID:    req.TechnicianID,
		OperatorName:  req.TechnicianName,
		LocationID:    0,
		LocationName:  "",
		Status:        model.LifecycleStatusBefore,
		Remark:        req.FaultDescription,
	}

	if err := s.lifecycleRepo.Create(ctx, lifecycle); err != nil {
		return fmt.Errorf("创建生命周期记录失败: %w", err)
	}

	return nil
}

// CompleteMaintenance 完成维修
func (s *ContainerLifecycleService) CompleteMaintenance(ctx context.Context, tenantID int64, id uint, req *model.CompleteMaintenanceRequest) error {
	// 检查维修记录是否存在
	maintenance, err := s.maintenanceRepo.GetByIDOnly(ctx, id)
	if err != nil {
		return fmt.Errorf("维修记录不存在: %w", err)
	}

	if maintenance.Status == model.MaintenanceStatusCompleted {
		return fmt.Errorf("维修已完成，不能重复完成")
	}

	// 更新维修记录
	updates := map[string]interface{}{
		"status":             model.MaintenanceStatusCompleted,
		"completed_date":      req.CompletedDate,
		"maintenance_content": req.MaintenanceContent,
		"spare_parts_used":   req.SparePartsUsed,
		"cost":               req.Cost,
		"remark":              req.Remark,
	}

	if err := s.maintenanceRepo.Update(ctx, id, updates); err != nil {
		return fmt.Errorf("更新维修记录失败: %w", err)
	}

	// 创建生命周期记录（维修完成事件）
	lifecycle := &model.ContainerLifecycle{
		TenantID:      tenantID,
		ContainerID:    maintenance.ContainerID,
		ContainerCode: maintenance.ContainerCode,
		EventType:     model.ContainerEventInService,
		EventDate:     req.CompletedDate,
		OperatorID:    maintenance.TechnicianID,
		OperatorName:  maintenance.TechnicianName,
		LocationID:    0,
		LocationName:  "",
		Status:        model.LifecycleStatusAfter,
		Remark:        "维修完成",
	}

	if err := s.lifecycleRepo.Create(ctx, lifecycle); err != nil {
		return fmt.Errorf("创建生命周期记录失败: %w", err)
	}

	// 更新容器状态为使用中
	containerUpdates := map[string]interface{}{
		"status": model.ContainerStatusInUse,
	}
	if err := s.containerRepo.Update(ctx, uint(maintenance.ContainerID), containerUpdates); err != nil {
		return fmt.Errorf("更新容器状态失败: %w", err)
	}

	return nil
}

// RetireContainer 报废容器
func (s *ContainerLifecycleService) RetireContainer(ctx context.Context, tenantID int64, req *model.RetireContainerRequest) error {
	// 检查容器是否存在
	container, err := s.containerRepo.GetByID(ctx, uint(req.ContainerID))
	if err != nil {
		return fmt.Errorf("容器不存在: %w", err)
	}

	now := time.Now()

	// 创建生命周期记录
	lifecycle := &model.ContainerLifecycle{
		TenantID:      tenantID,
		ContainerID:    req.ContainerID,
		ContainerCode: container.ContainerCode,
		EventType:     model.ContainerEventScrap,
		EventDate:     now,
		OperatorID:    req.OperatorID,
		OperatorName:  req.OperatorName,
		LocationID:    req.LocationID,
		LocationName:  req.LocationName,
		Status:        model.LifecycleStatusAfter,
		Remark:        req.Remark,
	}

	if err := s.lifecycleRepo.Create(ctx, lifecycle); err != nil {
		return fmt.Errorf("创建生命周期记录失败: %w", err)
	}

	// 更新容器状态为已报废
	updates := map[string]interface{}{
		"status": model.ContainerStatusScrapped,
	}
	if err := s.containerRepo.Update(ctx, uint(req.ContainerID), updates); err != nil {
		return fmt.Errorf("更新容器状态失败: %w", err)
	}

	return nil
}

// GetContainerTimeline 获取容器时间线
func (s *ContainerLifecycleService) GetContainerTimeline(ctx context.Context, containerID int64) ([]model.ContainerTimelineItem, error) {
	lifecycles, err := s.lifecycleRepo.GetByContainerID(ctx, containerID)
	if err != nil {
		return nil, err
	}

	timeline := make([]model.ContainerTimelineItem, 0, len(lifecycles))
	for _, lc := range lifecycles {
		timeline = append(timeline, model.ContainerTimelineItem{
			EventType:    string(lc.EventType),
			EventDate:    lc.EventDate,
			OperatorName: lc.OperatorName,
			LocationName: lc.LocationName,
			Status:       string(lc.Status),
			Remark:       lc.Remark,
		})
	}

	return timeline, nil
}

// ListContainerMaintenances 分页查询维修记录
func (s *ContainerLifecycleService) ListContainerMaintenances(ctx context.Context, tenantID int64, query *model.ContainerMaintenanceQuery) ([]model.ContainerMaintenance, int64, error) {
	return s.maintenanceRepo.List(ctx, tenantID, query)
}

// GetContainerMaintenance 获取单条维修记录
func (s *ContainerLifecycleService) GetContainerMaintenance(ctx context.Context, id uint) (*model.ContainerMaintenance, error) {
	return s.maintenanceRepo.GetByID(ctx, id)
}