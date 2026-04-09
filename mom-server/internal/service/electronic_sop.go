package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"time"
)

type ElectronicSOPService struct {
	repo *repository.ElectronicSOPRepository
}

func NewElectronicSOPService(repo *repository.ElectronicSOPRepository) *ElectronicSOPService {
	return &ElectronicSOPService{repo: repo}
}

func (s *ElectronicSOPService) List(ctx context.Context, query string) ([]model.ElectronicSOP, int64, error) {
	return s.repo.List(ctx, 1, query)
}

func (s *ElectronicSOPService) GetByID(ctx context.Context, id uint) (*model.ElectronicSOP, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ElectronicSOPService) Create(ctx context.Context, sop *model.ElectronicSOP) error {
	if sop.TenantID == 0 {
		sop.TenantID = 1
	}
	return s.repo.Create(ctx, sop)
}

func (s *ElectronicSOPService) Update(ctx context.Context, id uint, updates map[string]any) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *ElectronicSOPService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *ElectronicSOPService) GetByMaterial(ctx context.Context, materialID int64) ([]model.ElectronicSOP, error) {
	return s.repo.GetByMaterial(ctx, materialID)
}

type CodeRuleService struct {
	repo *repository.CodeRuleRepository
}

func NewCodeRuleService(repo *repository.CodeRuleRepository) *CodeRuleService {
	return &CodeRuleService{repo: repo}
}

func (s *CodeRuleService) List(ctx context.Context) ([]model.CodeRule, error) {
	return s.repo.List(ctx, 1)
}

func (s *CodeRuleService) GetByID(ctx context.Context, id uint) (*model.CodeRule, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CodeRuleService) Create(ctx context.Context, rule *model.CodeRule) error {
	if rule.TenantID == 0 {
		rule.TenantID = 1
	}
	return s.repo.Create(ctx, rule)
}

func (s *CodeRuleService) Update(ctx context.Context, id uint, updates map[string]any) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *CodeRuleService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

// GenerateCode 生成编码
func (s *CodeRuleService) GenerateCode(ctx context.Context, ruleCode string) (string, error) {
	rule, err := s.repo.GetByCode(ctx, 1, ruleCode)
	if err != nil {
		return "", fmt.Errorf("编码规则不存在: %s", ruleCode)
	}

	if rule.Status != 1 {
		return "", fmt.Errorf("编码规则已禁用")
	}

	// 检查是否需要重置序号
	now := time.Now()
	genDate := now.Format("20060102")
	needReset := false

	switch rule.ResetType {
	case "DAILY":
		if rule.SeqCurrent == 0 || rule.LastGenDate != genDate {
			needReset = true
		}
	case "MONTHLY":
		if rule.SeqCurrent == 0 || rule.LastGenDate[:6] != genDate[:6] {
			needReset = true
		}
	case "YEARLY":
		if rule.SeqCurrent == 0 || rule.LastGenDate[:4] != genDate[:4] {
			needReset = true
		}
	}

	newSeq := rule.SeqCurrent + 1
	if needReset {
		newSeq = rule.SeqStart
	}

	// 构建编码
	code := rule.Prefix
	if rule.DateFormat != "" {
		dateStr := ""
		switch rule.DateFormat {
		case "YYYYMMDD":
			dateStr = now.Format("20060102")
		case "YYYYMM":
			dateStr = now.Format("200601")
		case "YYYY":
			dateStr = now.Format("2006")
		}
		code += dateStr
	}
	if rule.MidFix != "" {
		code += rule.MidFix
	}
	// 格式化序号
	seqStr := fmt.Sprintf("%0"+fmt.Sprintf("%d", rule.SeqLength)+"d", newSeq)
	code += seqStr
	if rule.Suffix != "" {
		code += rule.Suffix
	}

	// 更新序号
	s.repo.UpdateSeq(ctx, rule.ID, newSeq)

	// 记录
	record := &model.CodeRuleRecord{
		TenantID:    1,
		RuleID:     rule.ID,
		RuleCode:   ruleCode,
		EntityType: rule.EntityType,
		GenDate:    genDate,
		SeqValue:   newSeq,
		GenCode:    code,
	}
	s.repo.CreateRecord(ctx, record)

	return code, nil
}

type FlowCardService struct {
	repo *repository.FlowCardRepository
}

func NewFlowCardService(repo *repository.FlowCardRepository) *FlowCardService {
	return &FlowCardService{repo: repo}
}

func (s *FlowCardService) List(ctx context.Context, query string) ([]model.FlowCard, int64, error) {
	return s.repo.List(ctx, 1, query)
}

func (s *FlowCardService) GetByID(ctx context.Context, id uint) (*model.FlowCard, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *FlowCardService) Create(ctx context.Context, card *model.FlowCard) error {
	if card.TenantID == 0 {
		card.TenantID = 1
	}
	return s.repo.Create(ctx, card)
}

func (s *FlowCardService) Update(ctx context.Context, id uint, updates map[string]any) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *FlowCardService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *FlowCardService) GetDetails(ctx context.Context, cardID int64) ([]model.FlowCardDetail, error) {
	return s.repo.GetDetails(ctx, cardID)
}

func (s *FlowCardService) CreateWithDetails(ctx context.Context, card *model.FlowCard, details []model.FlowCardDetail) error {
	err := s.repo.Create(ctx, card)
	if err != nil {
		return err
	}
	if len(details) > 0 {
		for i := range details {
			details[i].CardID = card.ID
		}
		return s.repo.CreateDetails(ctx, details)
	}
	return nil
}
