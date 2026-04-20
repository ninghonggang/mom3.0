package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// LabSampleService 检测样品服务
type LabSampleService struct {
	repo     *repository.LabSampleRepository
	codeRule *CodeRuleService
}

func NewLabSampleService(repo *repository.LabSampleRepository, codeRule *CodeRuleService) *LabSampleService {
	return &LabSampleService{repo: repo, codeRule: codeRule}
}

func (s *LabSampleService) List(ctx context.Context, query *model.LabSampleQuery) ([]model.LabSample, int64, error) {
	if query.Page == 0 {
		query.Page = 1
	}
	if query.PageSize == 0 {
		query.PageSize = 20
	}
	return s.repo.List(ctx, query)
}

func (s *LabSampleService) GetByID(ctx context.Context, id uint64) (*model.LabSample, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *LabSampleService) Create(ctx context.Context, sample *model.LabSample) error {
	if sample.SampleCode == "" && s.codeRule != nil {
		no, err := s.codeRule.GenerateCode(ctx, "LAB_SAMPLE")
		if err != nil {
			return fmt.Errorf("failed to generate sample code: %w", err)
		}
		sample.SampleCode = no
	}
	return s.repo.Create(ctx, sample)
}

func (s *LabSampleService) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *LabSampleService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

func (s *LabSampleService) SubmitForInspection(ctx context.Context, id uint64) error {
	return s.repo.Update(ctx, id, map[string]interface{}{
		"status": "INSPECTING",
	})
}

// LabTestItemService 检测项目服务
type LabTestItemService struct {
	repo *repository.LabTestItemRepository
}

func NewLabTestItemService(repo *repository.LabTestItemRepository) *LabTestItemService {
	return &LabTestItemService{repo: repo}
}

func (s *LabTestItemService) ListBySampleID(ctx context.Context, sampleID uint64) ([]model.LabTestItem, error) {
	return s.repo.ListBySampleID(ctx, sampleID)
}

func (s *LabTestItemService) GetByID(ctx context.Context, id uint64) (*model.LabTestItem, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *LabTestItemService) Create(ctx context.Context, item *model.LabTestItem) error {
	return s.repo.Create(ctx, item)
}

func (s *LabTestItemService) BatchCreate(ctx context.Context, sampleID uint64, tenantID uint64, items []model.LabTestItem) error {
	if len(items) == 0 {
		return nil
	}
	for i := range items {
		items[i].SampleID = sampleID
		items[i].TenantID = tenantID
	}
	return s.repo.BatchCreate(ctx, items)
}

func (s *LabTestItemService) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *LabTestItemService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

// LabReportService 检测报告服务
type LabReportService struct {
	repo     *repository.LabReportRepository
	codeRule *CodeRuleService
}

func NewLabReportService(repo *repository.LabReportRepository, codeRule *CodeRuleService) *LabReportService {
	return &LabReportService{repo: repo, codeRule: codeRule}
}

func (s *LabReportService) List(ctx context.Context, query *model.LabReportQuery) ([]model.LabReport, int64, error) {
	if query.Page == 0 {
		query.Page = 1
	}
	if query.PageSize == 0 {
		query.PageSize = 20
	}
	return s.repo.List(ctx, query)
}

func (s *LabReportService) GetByID(ctx context.Context, id uint64) (*model.LabReport, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *LabReportService) GetBySampleID(ctx context.Context, sampleID uint64) (*model.LabReport, error) {
	return s.repo.GetBySampleID(ctx, sampleID)
}

func (s *LabReportService) Create(ctx context.Context, report *model.LabReport) error {
	if report.ReportNo == "" && s.codeRule != nil {
		no, err := s.codeRule.GenerateCode(ctx, "LAB_REPORT")
		if err != nil {
			return fmt.Errorf("failed to generate report no: %w", err)
		}
		report.ReportNo = no
	}
	return s.repo.Create(ctx, report)
}

func (s *LabReportService) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *LabReportService) Approve(ctx context.Context, id uint64, approvedBy string) error {
	return s.repo.Update(ctx, id, map[string]interface{}{
		"approved_by": approvedBy,
	})
}

func (s *LabReportService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}
