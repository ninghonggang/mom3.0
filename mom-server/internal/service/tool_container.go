package service

import (
	"context"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"time"
)

// ToolService 器具服务
type ToolService struct {
	repo *repository.ToolRepository
}

func NewToolService(repo *repository.ToolRepository) *ToolService {
	return &ToolService{repo: repo}
}

func (s *ToolService) List(ctx context.Context, tenantID int64, query *model.ToolQuery) ([]model.Tool, int64, error) {
	return s.repo.List(ctx, tenantID, query)
}

func (s *ToolService) GetByID(ctx context.Context, id int64) (*model.Tool, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ToolService) Create(ctx context.Context, tool *model.Tool) error {
	if tool.TenantID == 0 {
		tool.TenantID = 1
	}
	return s.repo.Create(ctx, tool)
}

func (s *ToolService) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *ToolService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

// ToolContainerService 容器服务
type ToolContainerService struct {
	repo *repository.ToolContainerRepository
}

func NewToolContainerService(repo *repository.ToolContainerRepository) *ToolContainerService {
	return &ToolContainerService{repo: repo}
}

func (s *ToolContainerService) List(ctx context.Context, tenantID int64, query *model.ToolContainerQuery) ([]model.ToolContainer, int64, error) {
	return s.repo.List(ctx, tenantID, query)
}

func (s *ToolContainerService) GetByID(ctx context.Context, id int64) (*model.ToolContainer, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ToolContainerService) Create(ctx context.Context, container *model.ToolContainer) error {
	if container.TenantID == 0 {
		container.TenantID = 1
	}
	return s.repo.Create(ctx, container)
}

func (s *ToolContainerService) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *ToolContainerService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

// ToolContainerBindingService 器具容器绑定服务
type ToolContainerBindingService struct {
	bindingRepo *repository.ToolContainerBindingRepository
	toolRepo    *repository.ToolRepository
	containerRepo *repository.ToolContainerRepository
}

func NewToolContainerBindingService(bindingRepo *repository.ToolContainerBindingRepository, toolRepo *repository.ToolRepository, containerRepo *repository.ToolContainerRepository) *ToolContainerBindingService {
	return &ToolContainerBindingService{
		bindingRepo:  bindingRepo,
		toolRepo:    toolRepo,
		containerRepo: containerRepo,
	}
}

func (s *ToolContainerBindingService) List(ctx context.Context, tenantID int64, query *model.ToolContainerBindingQuery) ([]model.ToolContainerBinding, int64, error) {
	return s.bindingRepo.List(ctx, tenantID, query)
}

func (s *ToolContainerBindingService) GetByID(ctx context.Context, id int64) (*model.ToolContainerBinding, error) {
	return s.bindingRepo.GetByID(ctx, id)
}

func (s *ToolContainerBindingService) Bind(ctx context.Context, tenantID int64, req *model.ToolContainerBindingCreateRequest) error {
	// 检查器具是否存在
	tool, err := s.toolRepo.GetByID(ctx, req.ToolID)
	if err != nil {
		return err
	}

	// 检查容器是否存在
	container, err := s.containerRepo.GetByID(ctx, req.ContainerID)
	if err != nil {
		return err
	}

	// 检查器具是否已有活跃绑定
	existing, _ := s.bindingRepo.GetActiveBinding(ctx, tenantID, req.ToolID)
	if existing != nil {
		return err
	}

	binding := &model.ToolContainerBinding{
		TenantID:      tenantID,
		ToolID:        req.ToolID,
		ToolCode:      tool.ToolCode,
		ContainerID:   req.ContainerID,
		ContainerCode: container.ContainerCode,
		BindTime:      time.Now(),
		Status:        "BOUND",
		OperatorID:    req.OperatorID,
		OperatorName:  req.OperatorName,
		Remark:        req.Remark,
	}

	return s.bindingRepo.Create(ctx, binding)
}

func (s *ToolContainerBindingService) Unbind(ctx context.Context, id int64) error {
	now := time.Now()
	return s.bindingRepo.Update(ctx, id, map[string]interface{}{
		"status":      "UNBOUND",
		"unbind_time": now,
	})
}

func (s *ToolContainerBindingService) GetByToolID(ctx context.Context, tenantID int64, toolID int64) ([]model.ToolContainerBinding, error) {
	return s.bindingRepo.GetByToolID(ctx, tenantID, toolID)
}

func (s *ToolContainerBindingService) GetByContainerID(ctx context.Context, tenantID int64, containerID int64) ([]model.ToolContainerBinding, error) {
	return s.bindingRepo.GetByContainerID(ctx, tenantID, containerID)
}

func (s *ToolContainerBindingService) GetActiveBinding(ctx context.Context, tenantID int64, toolID int64) (*model.ToolContainerBinding, error) {
	return s.bindingRepo.GetActiveBinding(ctx, tenantID, toolID)
}
