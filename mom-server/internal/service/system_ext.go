package service

import (
	"context"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type PrintTemplateService struct {
	repo *repository.PrintTemplateRepository
}

func NewPrintTemplateService(repo *repository.PrintTemplateRepository) *PrintTemplateService {
	return &PrintTemplateService{repo: repo}
}

func (s *PrintTemplateService) List(ctx context.Context) ([]model.PrintTemplate, error) {
	return s.repo.List(ctx, 1)
}

func (s *PrintTemplateService) GetByID(ctx context.Context, id uint) (*model.PrintTemplate, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *PrintTemplateService) Create(ctx context.Context, t *model.PrintTemplate) error {
	if t.TenantID == 0 {
		t.TenantID = 1
	}
	return s.repo.Create(ctx, t)
}

func (s *PrintTemplateService) Update(ctx context.Context, id uint, updates map[string]any) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *PrintTemplateService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

type NoticeService struct {
	repo *repository.NoticeRepository
}

func NewNoticeService(repo *repository.NoticeRepository) *NoticeService {
	return &NoticeService{repo: repo}
}

func (s *NoticeService) List(ctx context.Context, query string) ([]model.Notice, int64, error) {
	return s.repo.List(ctx, 1, query)
}

func (s *NoticeService) GetByID(ctx context.Context, id uint) (*model.Notice, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *NoticeService) Create(ctx context.Context, n *model.Notice) error {
	if n.TenantID == 0 {
		n.TenantID = 1
	}
	return s.repo.Create(ctx, n)
}

func (s *NoticeService) Update(ctx context.Context, id uint, updates map[string]any) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *NoticeService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *NoticeService) Publish(ctx context.Context, id uint) error {
	return s.repo.Update(ctx, id, map[string]any{"status": 2})
}

func (s *NoticeService) View(ctx context.Context, id uint) error {
	s.repo.IncrementViewCount(ctx, id)
	return nil
}

func (s *NoticeService) GetMyNotices(ctx context.Context, userID int64) ([]model.Notice, error) {
	return s.repo.GetMyNotices(ctx, 1, userID)
}
