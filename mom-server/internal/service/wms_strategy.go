package service

import (
	"context"
	"fmt"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// WmsStrategyService 策略配置服务层
type WmsStrategyService struct {
	repo *repository.WmsStrategyRepository
}

// NewWmsStrategyService 创建策略配置服务实例
func NewWmsStrategyService(repo *repository.WmsStrategyRepository) *WmsStrategyService {
	return &WmsStrategyService{repo: repo}
}

// List 获取策略配置列表
func (s *WmsStrategyService) List(ctx context.Context, query *model.WmsStrategyQueryVO) ([]model.WmsStrategy, int64, error) {
	// 设置默认分页
	if query != nil {
		if query.Page <= 0 {
			query.Page = 1
		}
		if query.PageSize <= 0 {
			query.PageSize = 20
		}
	}
	return s.repo.List(ctx, query)
}

// Get 获取策略配置详情
func (s *WmsStrategyService) Get(ctx context.Context, id string) (*model.WmsStrategy, error) {
	var strategyID uint64
	_, err := fmt.Sscanf(id, "%d", &strategyID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, strategyID)
}

// Create 创建策略配置
func (s *WmsStrategyService) Create(ctx context.Context, req *model.WmsStrategyCreateReqVO) error {
	strategy := &model.WmsStrategy{
		StrategyCode: req.StrategyCode,
		StrategyName: req.StrategyName,
		StrategyType: req.StrategyType,
		RuleContent:  req.RuleContent,
		Priority:     req.Priority,
		Status:       "ACTIVE",
	}
	return s.repo.Create(ctx, strategy)
}

// Update 更新策略配置
func (s *WmsStrategyService) Update(ctx context.Context, id string, req *model.WmsStrategyUpdateReqVO) error {
	var strategyID uint64
	_, err := fmt.Sscanf(id, "%d", &strategyID)
	if err != nil {
		return err
	}
	updates := map[string]interface{}{
		"strategy_name": req.StrategyName,
		"strategy_type": req.StrategyType,
		"rule_content":  req.RuleContent,
		"priority":      req.Priority,
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}
	return s.repo.Update(ctx, strategyID, updates)
}

// Delete 删除策略配置
func (s *WmsStrategyService) Delete(ctx context.Context, id string) error {
	var strategyID uint64
	_, err := fmt.Sscanf(id, "%d", &strategyID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, strategyID)
}
