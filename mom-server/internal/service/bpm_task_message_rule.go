package service

import (
	"context"
	"fmt"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type BpmTaskMessageRuleService struct {
	repo *repository.BpmTaskMessageRuleRepository
}

func NewBpmTaskMessageRuleService(repo *repository.BpmTaskMessageRuleRepository) *BpmTaskMessageRuleService {
	return &BpmTaskMessageRuleService{repo: repo}
}

func (s *BpmTaskMessageRuleService) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.BpmTaskMessageRule, int64, error) {
	return s.repo.List(ctx, tenantID, query)
}

func (s *BpmTaskMessageRuleService) Get(ctx context.Context, id string) (*model.BpmTaskMessageRule, error) {
	var ruleID uint
	_, err := fmt.Sscanf(id, "%d", &ruleID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, ruleID)
}

func (s *BpmTaskMessageRuleService) Create(ctx context.Context, rule *model.BpmTaskMessageRule) error {
	return s.repo.Create(ctx, rule)
}

func (s *BpmTaskMessageRuleService) Update(ctx context.Context, id string, req *model.BpmTaskMessageRuleUpdateReqVO) error {
	var ruleID uint
	_, err := fmt.Sscanf(id, "%d", &ruleID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, ruleID, map[string]interface{}{
		"rule_name":       req.RuleName,
		"process_def_key": req.ProcessDefKey,
		"task_def_key":    req.TaskDefKey,
		"message_type":    req.MessageType,
		"template_code":   req.TemplateCode,
		"is_enabled":      req.IsEnabled,
		"priority":        req.Priority,
		"remark":          req.Remark,
	})
}

func (s *BpmTaskMessageRuleService) Delete(ctx context.Context, id string) error {
	var ruleID uint
	_, err := fmt.Sscanf(id, "%d", &ruleID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, ruleID)
}

func (s *BpmTaskMessageRuleService) Enable(ctx context.Context, id string) error {
	var ruleID uint
	_, err := fmt.Sscanf(id, "%d", &ruleID)
	if err != nil {
		return err
	}
	return s.repo.Enable(ctx, ruleID)
}

func (s *BpmTaskMessageRuleService) Disable(ctx context.Context, id string) error {
	var ruleID uint
	_, err := fmt.Sscanf(id, "%d", &ruleID)
	if err != nil {
		return err
	}
	return s.repo.Disable(ctx, ruleID)
}
