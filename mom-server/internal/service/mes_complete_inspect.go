package service

import (
	"context"
	"fmt"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// CompleteInspectService 齐套检查服务
type CompleteInspectService struct {
	repo *repository.CompleteInspectRepository
}

func NewCompleteInspectService(repo *repository.CompleteInspectRepository) *CompleteInspectService {
	return &CompleteInspectService{repo: repo}
}

// GetConfig 获取齐套检查标识配置
func (s *CompleteInspectService) GetConfig(ctx context.Context, paramCode string) (*model.MesConfigInfo, error) {
	// paramCode 应该是配置码，如 "PRE_START_CHECK_ON"
	config, err := s.repo.GetConfigByCode(ctx, paramCode)
	if err != nil {
		// 如果没找到配置，返回默认值
		return &model.MesConfigInfo{
			ConfigCode:  paramCode,
			ConfigValue: "true",
			Status:      1,
		}, nil
	}
	return config, nil
}

// GetOrderDayBom 获取日计划Bom信息
func (s *CompleteInspectService) GetOrderDayBom(ctx context.Context, tenantID int64, req *model.MesWorkSchedulingBaseVO) ([]model.MesOrderDayBomRespVO, error) {
	if req.OrderDayID <= 0 {
		return nil, fmt.Errorf("order_day_id is required")
	}
	return s.repo.ListOrderDayBom(ctx, req.OrderDayID)
}

// GetOrderDayBomPage 获取日计划Bom信息(分页)
func (s *CompleteInspectService) GetOrderDayBomPage(ctx context.Context, tenantID int64, req *model.MesWorkSchedulingPageReqVO) ([]model.MesOrderDayBomRespVO, int64, error) {
	query := map[string]interface{}{
		"order_day_id": req.OrderDayID,
		"order_day_no": req.OrderDayNo,
		"page":         req.Page,
		"page_size":    req.PageSize,
	}
	return s.repo.ListOrderDayBomPage(ctx, tenantID, query)
}

// GetOrderDayWorker 获取日计划人员信息
func (s *CompleteInspectService) GetOrderDayWorker(ctx context.Context, tenantID int64, req *model.MesWorkSchedulingBaseVO) ([]model.MesOrderDayWorkerRespVO, error) {
	if req.OrderDayID <= 0 {
		return nil, fmt.Errorf("order_day_id is required")
	}
	return s.repo.ListOrderDayWorker(ctx, req.OrderDayID)
}

// GetOrderDayWorkerPage 获取日计划Worker信息(分页)
func (s *CompleteInspectService) GetOrderDayWorkerPage(ctx context.Context, tenantID int64, req *model.MesWorkSchedulingPageReqVO) ([]model.MesOrderDayWorkerRespVO, int64, error) {
	query := map[string]interface{}{
		"order_day_id": req.OrderDayID,
		"order_day_no": req.OrderDayNo,
		"page":         req.Page,
		"page_size":    req.PageSize,
	}
	return s.repo.ListOrderDayWorkerPage(ctx, tenantID, query)
}

// GetOrderDayEquipment 获取日计划设备信息
func (s *CompleteInspectService) GetOrderDayEquipment(ctx context.Context, tenantID int64, req *model.MesWorkSchedulingBaseVO) ([]model.MesOrderDayEquipmentRespVO, error) {
	if req.OrderDayID <= 0 {
		return nil, fmt.Errorf("order_day_id is required")
	}
	return s.repo.ListOrderDayEquipment(ctx, req.OrderDayID)
}

// GetOrderDayEquipmentPage 获取日计划Equipment信息(分页)
func (s *CompleteInspectService) GetOrderDayEquipmentPage(ctx context.Context, tenantID int64, req *model.MesWorkSchedulingPageReqVO) ([]model.MesOrderDayEquipmentRespVO, int64, error) {
	query := map[string]interface{}{
		"order_day_id": req.OrderDayID,
		"order_day_no": req.OrderDayNo,
		"page":         req.Page,
		"page_size":    req.PageSize,
	}
	return s.repo.ListOrderDayEquipmentPage(ctx, tenantID, query)
}

// Update 更新齐套检查状态
func (s *CompleteInspectService) Update(ctx context.Context, tenantID int64, req *model.CompleteInspectUpdate, username string) error {
	if req.OrderDayID <= 0 {
		return fmt.Errorf("order_day_id is required")
	}

	updates := map[string]interface{}{}
	if req.KitCheckStatus != "" {
		updates["kit_check_status"] = req.KitCheckStatus
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	// 先尝试获取现有记录
	_, err := s.repo.GetByOrderDayID(ctx, req.OrderDayID)
	if err != nil {
		// 记录不存在，创建新记录
		inspect := &model.MesCompleteInspect{
			TenantID:       tenantID,
			OrderDayID:     req.OrderDayID,
			KitCheckStatus: req.KitCheckStatus,
			Remark:         req.Remark,
		}
		return s.repo.Create(ctx, inspect)
	}

	// 记录存在，更新
	return s.repo.UpdateByOrderDayID(ctx, req.OrderDayID, updates)
}
