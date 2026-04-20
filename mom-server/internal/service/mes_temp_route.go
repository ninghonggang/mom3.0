package service

import (
	"context"
	"errors"
	"fmt"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// TempRouteService 临时工艺路线服务
type TempRouteService struct {
	repo *repository.TempRouteRepository
}

func NewTempRouteService(repo *repository.TempRouteRepository) *TempRouteService {
	return &TempRouteService{repo: repo}
}

// Create 创建临时工艺路线
func (s *TempRouteService) Create(ctx context.Context, tenantID int64, req *model.TempRouteCreate, username string) (*model.TempRoute, error) {
	tempRoute := &model.TempRoute{
		TenantID:      tenantID,
		OrderDayID:    req.OrderDayID,
		TempRouteName: req.TempRouteName,
		RouteContent:  req.RouteContent,
		Reason:        req.Reason,
		Status:        0, // 待审核
		Creator:       username,
	}

	if err := s.repo.Create(ctx, tempRoute); err != nil {
		return nil, fmt.Errorf("创建临时工艺路线失败: %w", err)
	}

	return s.repo.GetByID(ctx, uint(tempRoute.ID))
}

// GetByID 获取临时工艺路线
func (s *TempRouteService) GetByID(ctx context.Context, id int64) (*model.TempRoute, error) {
	return s.repo.GetByID(ctx, uint(id))
}

// Update 更新临时工艺路线
func (s *TempRouteService) Update(ctx context.Context, id int64, req *model.TempRouteUpdate, username string) (*model.TempRoute, error) {
	tempRoute, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	// 仅待审核状态可更新
	if tempRoute.Status != 0 {
		return nil, errors.New("仅待审核状态可更新")
	}

	updates := map[string]interface{}{}
	if req.TempRouteName != "" {
		updates["temp_route_name"] = req.TempRouteName
	}
	if req.RouteContent != "" {
		updates["route_content"] = req.RouteContent
	}
	if req.Reason != "" {
		updates["reason"] = req.Reason
	}

	if err := s.repo.Update(ctx, uint(id), updates); err != nil {
		return nil, fmt.Errorf("更新临时工艺路线失败: %w", err)
	}

	return s.repo.GetByID(ctx, uint(id))
}

// Delete 删除临时工艺路线
func (s *TempRouteService) Delete(ctx context.Context, id int64) error {
	tempRoute, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return err
	}

	// 仅待审核状态可删除
	if tempRoute.Status != 0 {
		return errors.New("仅待审核状态可删除")
	}

	return s.repo.Delete(ctx, uint(id))
}

// ListByOrderDayID 根据日计划ID查询临时工艺路线列表
func (s *TempRouteService) ListByOrderDayID(ctx context.Context, orderDayID int64) ([]model.TempRoute, error) {
	return s.repo.ListByOrderDayID(ctx, orderDayID)
}

// Approve 审核临时工艺路线
func (s *TempRouteService) Approve(ctx context.Context, id int64, status int, username string) error {
	tempRoute, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return err
	}

	// 仅待审核状态可审核
	if tempRoute.Status != 0 {
		return errors.New("仅待审核状态可审核")
	}

	if status != 1 && status != 2 {
		return errors.New("审核状态无效，1=批准，2=拒绝")
	}

	return s.repo.Update(ctx, uint(id), map[string]interface{}{
		"status": status,
	})
}
