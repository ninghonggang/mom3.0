package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type InspectionTemplateService struct {
	repo *repository.InspectionTemplateRepository
}

func NewInspectionTemplateService(repo *repository.InspectionTemplateRepository) *InspectionTemplateService {
	return &InspectionTemplateService{repo: repo}
}

func (s *InspectionTemplateService) Create(ctx context.Context, m *model.InspectionTemplate) error {
	return s.repo.Create(ctx, m)
}

func (s *InspectionTemplateService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *InspectionTemplateService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *InspectionTemplateService) GetByID(ctx context.Context, id uint) (*model.InspectionTemplate, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *InspectionTemplateService) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.InspectionTemplate, int64, error) {
	return s.repo.List(ctx, offset, limit, filters)
}

type InspectionItemService struct {
	repo *repository.InspectionItemRepository
}

func NewInspectionItemService(repo *repository.InspectionItemRepository) *InspectionItemService {
	return &InspectionItemService{repo: repo}
}

func (s *InspectionItemService) Create(ctx context.Context, m *model.InspectionItem) error {
	return s.repo.Create(ctx, m)
}

func (s *InspectionItemService) CreateBatch(ctx context.Context, items []model.InspectionItem) error {
	return s.repo.CreateBatch(ctx, items)
}

func (s *InspectionItemService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *InspectionItemService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *InspectionItemService) ListByTemplate(ctx context.Context, templateID uint) ([]model.InspectionItem, error) {
	return s.repo.ListByTemplate(ctx, templateID)
}

type InspectionPlanService struct {
	repo     *repository.InspectionPlanRepository
	codeRule *CodeRuleService
}

func NewInspectionPlanService(repo *repository.InspectionPlanRepository, codeRule *CodeRuleService) *InspectionPlanService {
	return &InspectionPlanService{repo: repo, codeRule: codeRule}
}

func (s *InspectionPlanService) Create(ctx context.Context, m *model.InspectionPlan) error {
	if m.PlanNo == "" {
		no, err := s.codeRule.GenerateCode(ctx, "INSPECTION_PLAN")
		if err != nil {
			return fmt.Errorf("failed to generate plan no: %w", err)
		}
		m.PlanNo = no
	}
	return s.repo.Create(ctx, m)
}

func (s *InspectionPlanService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *InspectionPlanService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *InspectionPlanService) GetByID(ctx context.Context, id uint) (*model.InspectionPlan, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *InspectionPlanService) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.InspectionPlan, int64, error) {
	return s.repo.List(ctx, offset, limit, filters)
}

type InspectionRecordService struct {
	repo     *repository.InspectionRecordRepository
	codeRule *CodeRuleService
}

func NewInspectionRecordService(repo *repository.InspectionRecordRepository, codeRule *CodeRuleService) *InspectionRecordService {
	return &InspectionRecordService{repo: repo, codeRule: codeRule}
}

func (s *InspectionRecordService) Create(ctx context.Context, m *model.InspectionRecord) error {
	if m.RecordNo == "" {
		no, err := s.codeRule.GenerateCode(ctx, "INSPECTION_RECORD")
		if err != nil {
			return fmt.Errorf("failed to generate record no: %w", err)
		}
		m.RecordNo = no
	}
	return s.repo.Create(ctx, m)
}

func (s *InspectionRecordService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *InspectionRecordService) GetByID(ctx context.Context, id uint) (*model.InspectionRecord, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *InspectionRecordService) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.InspectionRecord, int64, error) {
	return s.repo.List(ctx, offset, limit, filters)
}

type InspectionResultService struct {
	repo *repository.InspectionResultRepository
}

func NewInspectionResultService(repo *repository.InspectionResultRepository) *InspectionResultService {
	return &InspectionResultService{repo: repo}
}

func (s *InspectionResultService) Create(ctx context.Context, m *model.InspectionResult) error {
	return s.repo.Create(ctx, m)
}

func (s *InspectionResultService) CreateBatch(ctx context.Context, items []model.InspectionResult) error {
	return s.repo.CreateBatch(ctx, items)
}

func (s *InspectionResultService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *InspectionResultService) ListByRecord(ctx context.Context, recordID uint) ([]model.InspectionResult, error) {
	return s.repo.ListByRecord(ctx, recordID)
}

type InspectionDefectService struct {
	repo     *repository.InspectionDefectRepository
	codeRule *CodeRuleService
}

func NewInspectionDefectService(repo *repository.InspectionDefectRepository, codeRule *CodeRuleService) *InspectionDefectService {
	return &InspectionDefectService{repo: repo, codeRule: codeRule}
}

func (s *InspectionDefectService) Create(ctx context.Context, m *model.InspectionDefect) error {
	if m.DefectNo == "" {
		no, err := s.codeRule.GenerateCode(ctx, "INSPECTION_DEFECT")
		if err != nil {
			return fmt.Errorf("failed to generate defect no: %w", err)
		}
		m.DefectNo = no
	}
	return s.repo.Create(ctx, m)
}

func (s *InspectionDefectService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *InspectionDefectService) GetByID(ctx context.Context, id uint) (*model.InspectionDefect, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *InspectionDefectService) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.InspectionDefect, int64, error) {
	return s.repo.List(ctx, offset, limit, filters)
}
