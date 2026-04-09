package service

import (
	"context"
	"errors"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"time"
)

type ContainerService struct {
	repo         *repository.ContainerRepository
	movementRepo *repository.ContainerMovementRepository
}

func NewContainerService(repo *repository.ContainerRepository, movementRepo *repository.ContainerMovementRepository) *ContainerService {
	return &ContainerService{
		repo:         repo,
		movementRepo: movementRepo,
	}
}

// List 分页查询器具
func (s *ContainerService) List(ctx context.Context, tenantID int64, params model.ContainerQueryParams) ([]model.ContainerMaster, int64, error) {
	return s.repo.List(ctx, tenantID, params)
}

// GetByID 根据ID查询
func (s *ContainerService) GetByID(ctx context.Context, id string) (*model.ContainerMaster, error) {
	var containerID uint
	_, err := parseID(id, &containerID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, containerID)
}

// Create 创建器具
func (s *ContainerService) Create(ctx context.Context, container *model.ContainerMaster) error {
	// 检查编号是否重复
	existing, err := s.repo.GetByCode(ctx, container.ContainerCode)
	if err == nil && existing != nil && existing.ID > 0 {
		return errors.New("器具编号已存在")
	}
	// 检查条码是否重复
	if container.Barcode != "" {
		existing, err := s.repo.GetByBarcode(ctx, container.Barcode)
		if err == nil && existing != nil && existing.ID > 0 {
			return errors.New("器具条码已存在")
		}
	}
	container.Status = model.ContainerStatusInStock
	return s.repo.Create(ctx, container)
}

// Update 更新器具
func (s *ContainerService) Update(ctx context.Context, id string, container *model.ContainerMaster) error {
	var containerID uint
	_, err := parseID(id, &containerID)
	if err != nil {
		return err
	}

	updates := map[string]interface{}{
		"container_name":      container.ContainerName,
		"container_type":      container.ContainerType,
		"standard_qty":        container.StandardQty,
		"applicable_products": container.ApplicableProducts,
		"location_type":      container.LocationType,
		"current_location":   container.CurrentLocation,
		"customer_id":        container.CustomerID,
	}

	return s.repo.Update(ctx, containerID, updates)
}

// Delete 删除器具
func (s *ContainerService) Delete(ctx context.Context, id string) error {
	var containerID uint
	_, err := parseID(id, &containerID)
	if err != nil {
		return err
	}

	// 检查是否有未完结的流转记录
	hasOpen, err := s.repo.HasOpenMovements(ctx, containerID)
	if err != nil {
		return err
	}
	if hasOpen {
		return errors.New("器具存在未完结的流转记录，无法删除")
	}

	return s.repo.Delete(ctx, containerID)
}

// In 入库操作
func (s *ContainerService) In(ctx context.Context, id string, req *model.ContainerMovement) error {
	var containerID uint
	_, err := parseID(id, &containerID)
	if err != nil {
		return err
	}

	// 获取器具信息
	container, err := s.repo.GetByID(ctx, containerID)
	if err != nil {
		return errors.New("器具不存在")
	}

	// 更新状态为在库
	err = s.repo.Update(ctx, containerID, map[string]interface{}{
		"status":           model.ContainerStatusInStock,
		"current_location": req.ToLocation,
	})
	if err != nil {
		return err
	}

	// 记录流转
	movement := &model.ContainerMovement{
		TenantID:       req.TenantID,
		ContainerID:    int64(containerID),
		ContainerCode:  container.ContainerCode,
		MovementType:   model.MovementTypeIn,
		FromLocation:   req.FromLocation,
		ToLocation:     req.ToLocation,
		Qty:            1,
		RelatedOrderNo: req.RelatedOrderNo,
		OperatorID:     req.OperatorID,
		OperatorName:   req.OperatorName,
		MovementTime:   time.Now(),
		Remark:         req.Remark,
	}

	return s.movementRepo.Create(ctx, movement)
}

// Out 出库操作
func (s *ContainerService) Out(ctx context.Context, id string, req *model.ContainerMovement) error {
	var containerID uint
	_, err := parseID(id, &containerID)
	if err != nil {
		return err
	}

	// 获取器具信息
	container, err := s.repo.GetByID(ctx, containerID)
	if err != nil {
		return errors.New("器具不存在")
	}

	if container.Status != model.ContainerStatusInStock {
		return errors.New("器具当前状态不允许出库")
	}

	// 更新状态为使用中
	err = s.repo.Update(ctx, containerID, map[string]interface{}{
		"status": model.ContainerStatusInUse,
	})
	if err != nil {
		return err
	}

	// 使用次数+1
	s.repo.IncrementTrips(ctx, containerID)

	// 记录流转
	movement := &model.ContainerMovement{
		TenantID:       req.TenantID,
		ContainerID:    int64(containerID),
		ContainerCode:  container.ContainerCode,
		MovementType:   model.MovementTypeOut,
		FromLocation:   req.FromLocation,
		ToLocation:     req.ToLocation,
		Qty:            1,
		RelatedOrderNo: req.RelatedOrderNo,
		OperatorID:     req.OperatorID,
		OperatorName:   req.OperatorName,
		MovementTime:   time.Now(),
		Remark:         req.Remark,
	}

	return s.movementRepo.Create(ctx, movement)
}

// Return 退回操作
func (s *ContainerService) Return(ctx context.Context, id string, req *model.ContainerMovement) error {
	var containerID uint
	_, err := parseID(id, &containerID)
	if err != nil {
		return err
	}

	// 获取器具信息
	container, err := s.repo.GetByID(ctx, containerID)
	if err != nil {
		return errors.New("器具不存在")
	}

	// 更新状态为在库，更新位置
	err = s.repo.Update(ctx, containerID, map[string]interface{}{
		"status":           model.ContainerStatusInStock,
		"current_location": req.ToLocation,
	})
	if err != nil {
		return err
	}

	// 记录流转
	movement := &model.ContainerMovement{
		TenantID:       req.TenantID,
		ContainerID:    int64(containerID),
		ContainerCode:  container.ContainerCode,
		MovementType:   model.MovementTypeReturn,
		FromLocation:   req.FromLocation,
		ToLocation:     req.ToLocation,
		Qty:            1,
		RelatedOrderNo: req.RelatedOrderNo,
		OperatorID:     req.OperatorID,
		OperatorName:   req.OperatorName,
		MovementTime:   time.Now(),
		Remark:         req.Remark,
	}

	return s.movementRepo.Create(ctx, movement)
}

// Transfer 调拨操作
func (s *ContainerService) Transfer(ctx context.Context, id string, req *model.ContainerMovement) error {
	var containerID uint
	_, err := parseID(id, &containerID)
	if err != nil {
		return err
	}

	// 获取器具信息
	container, err := s.repo.GetByID(ctx, containerID)
	if err != nil {
		return errors.New("器具不存在")
	}

	// 更新位置
	err = s.repo.Update(ctx, containerID, map[string]interface{}{
		"current_location": req.ToLocation,
	})
	if err != nil {
		return err
	}

	// 记录流转
	movement := &model.ContainerMovement{
		TenantID:       req.TenantID,
		ContainerID:    int64(containerID),
		ContainerCode:  container.ContainerCode,
		MovementType:   model.MovementTypeTransfer,
		FromLocation:   req.FromLocation,
		ToLocation:     req.ToLocation,
		Qty:            1,
		RelatedOrderNo: req.RelatedOrderNo,
		OperatorID:     req.OperatorID,
		OperatorName:   req.OperatorName,
		MovementTime:   time.Now(),
		Remark:         req.Remark,
	}

	return s.movementRepo.Create(ctx, movement)
}

// Clean 清洁操作
func (s *ContainerService) Clean(ctx context.Context, id string, req *model.ContainerMovement) error {
	var containerID uint
	_, err := parseID(id, &containerID)
	if err != nil {
		return err
	}

	// 获取器具信息
	container, err := s.repo.GetByID(ctx, containerID)
	if err != nil {
		return errors.New("器具不存在")
	}

	// 更新清洁日期
	s.repo.UpdateLastCleanDate(ctx, containerID)

	// 记录流转
	movement := &model.ContainerMovement{
		TenantID:       req.TenantID,
		ContainerID:    int64(containerID),
		ContainerCode:  container.ContainerCode,
		MovementType:   model.MovementTypeClean,
		FromLocation:   req.FromLocation,
		ToLocation:     req.ToLocation,
		Qty:            1,
		OperatorID:     req.OperatorID,
		OperatorName:   req.OperatorName,
		MovementTime:   time.Now(),
		Remark:         req.Remark,
	}

	return s.movementRepo.Create(ctx, movement)
}

// GetMovements 获取流转记录
func (s *ContainerService) GetMovements(ctx context.Context, id string, page, pageSize int) ([]model.ContainerMovement, int64, error) {
	var containerID uint
	_, err := parseID(id, &containerID)
	if err != nil {
		return nil, 0, err
	}
	return s.movementRepo.ListByContainerID(ctx, containerID, page, pageSize)
}

// parseID 解析ID字符串
func parseID(s string, id *uint) (bool, error) {
	var v uint
	_, err := parseUint(s, &v)
	*id = v
	return err == nil, err
}

func parseUint(s string, v *uint) (bool, error) {
	n, err := parseInt(s)
	*v = uint(n)
	return err == nil, err
}

func parseInt(s string) (int, error) {
	var n int
	_, err := parseStringToInt(s, &n)
	return n, err
}

func parseStringToInt(s string, v *int) (bool, error) {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false, errors.New("invalid id")
		}
	}
	n := 0
	for _, c := range s {
		n = n*10 + int(c-'0')
	}
	*v = n
	return true, nil
}
