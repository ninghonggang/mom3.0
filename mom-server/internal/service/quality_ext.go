package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// ========== QRCI Service ==========

type QRCIService struct {
	repo     *repository.QRCIRepository
	codeRule *CodeRuleService
}

func NewQRCIService(repo *repository.QRCIRepository, codeRule *CodeRuleService) *QRCIService {
	return &QRCIService{repo: repo, codeRule: codeRule}
}

func (s *QRCIService) Create(ctx context.Context, m *model.QRCI) error {
	if m.QRICNo == "" {
		no, err := s.codeRule.GenerateCode(ctx, "QRCI")
		if err != nil {
			return fmt.Errorf("failed to generate QRCI no: %w", err)
		}
		m.QRICNo = no
	}
	return s.repo.Create(ctx, m)
}

func (s *QRCIService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *QRCIService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *QRCIService) GetByID(ctx context.Context, id uint) (*model.QRCI, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *QRCIService) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.QRCI, int64, error) {
	return s.repo.List(ctx, offset, limit, filters)
}

type QRCI5WhyService struct {
	repo *repository.QRCI5WhyRepository
}

func NewQRCI5WhyService(repo *repository.QRCI5WhyRepository) *QRCI5WhyService {
	return &QRCI5WhyService{repo: repo}
}

func (s *QRCI5WhyService) Create(ctx context.Context, m *model.QRCI5Why) error {
	return s.repo.Create(ctx, m)
}

func (s *QRCI5WhyService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *QRCI5WhyService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *QRCI5WhyService) ListByQRCI(ctx context.Context, qrciID uint) ([]model.QRCI5Why, error) {
	return s.repo.ListByQRCI(ctx, qrciID)
}

type QRCIActionService struct {
	repo *repository.QRCIActionRepository
}

func NewQRCIActionService(repo *repository.QRCIActionRepository) *QRCIActionService {
	return &QRCIActionService{repo: repo}
}

func (s *QRCIActionService) Create(ctx context.Context, m *model.QRCIAction) error {
	return s.repo.Create(ctx, m)
}

func (s *QRCIActionService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *QRCIActionService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *QRCIActionService) GetByID(ctx context.Context, id uint) (*model.QRCIAction, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *QRCIActionService) ListByQRCI(ctx context.Context, qrciID uint) ([]model.QRCIAction, error) {
	return s.repo.ListByQRCI(ctx, qrciID)
}

type QRCIVerificationService struct {
	repo *repository.QRCIVerificationRepository
}

func NewQRCIVerificationService(repo *repository.QRCIVerificationRepository) *QRCIVerificationService {
	return &QRCIVerificationService{repo: repo}
}

func (s *QRCIVerificationService) Create(ctx context.Context, m *model.QRCIVerification) error {
	return s.repo.Create(ctx, m)
}

func (s *QRCIVerificationService) ListByQRCI(ctx context.Context, qrciID uint) ([]model.QRCIVerification, error) {
	return s.repo.ListByQRCI(ctx, qrciID)
}

// ========== LPA Service ==========

type LPAStandardService struct {
	repo *repository.LPAStandardRepository
}

func NewLPAStandardService(repo *repository.LPAStandardRepository) *LPAStandardService {
	return &LPAStandardService{repo: repo}
}

func (s *LPAStandardService) Create(ctx context.Context, m *model.LPAStandard) error {
	return s.repo.Create(ctx, m)
}

func (s *LPAStandardService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *LPAStandardService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *LPAStandardService) GetByID(ctx context.Context, id uint) (*model.LPAStandard, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *LPAStandardService) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.LPAStandard, int64, error) {
	return s.repo.List(ctx, offset, limit, filters)
}

type LPAQuestionService struct {
	repo *repository.LPAQuestionRepository
}

func NewLPAQuestionService(repo *repository.LPAQuestionRepository) *LPAQuestionService {
	return &LPAQuestionService{repo: repo}
}

func (s *LPAQuestionService) Create(ctx context.Context, m *model.LPAQuestion) error {
	return s.repo.Create(ctx, m)
}

func (s *LPAQuestionService) CreateBatch(ctx context.Context, items []model.LPAQuestion) error {
	return s.repo.CreateBatch(ctx, items)
}

func (s *LPAQuestionService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *LPAQuestionService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *LPAQuestionService) ListByStandard(ctx context.Context, standardID uint) ([]model.LPAQuestion, error) {
	return s.repo.ListByStandard(ctx, standardID)
}

type LPARecordService struct {
	repo     *repository.LPARecordRepository
	codeRule *CodeRuleService
}

func NewLPARecordService(repo *repository.LPARecordRepository, codeRule *CodeRuleService) *LPARecordService {
	return &LPARecordService{repo: repo, codeRule: codeRule}
}

func (s *LPARecordService) Create(ctx context.Context, m *model.LPARecord) error {
	if m.RecordNo == "" {
		no, err := s.codeRule.GenerateCode(ctx, "LPA_RECORD")
		if err != nil {
			return fmt.Errorf("failed to generate record no: %w", err)
		}
		m.RecordNo = no
	}
	return s.repo.Create(ctx, m)
}

func (s *LPARecordService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *LPARecordService) GetByID(ctx context.Context, id uint) (*model.LPARecord, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *LPARecordService) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.LPARecord, int64, error) {
	return s.repo.List(ctx, offset, limit, filters)
}

type LPARecordItemService struct {
	repo *repository.LPARecordItemRepository
}

func NewLPARecordItemService(repo *repository.LPARecordItemRepository) *LPARecordItemService {
	return &LPARecordItemService{repo: repo}
}

func (s *LPARecordItemService) CreateBatch(ctx context.Context, items []model.LPARecordItem) error {
	return s.repo.CreateBatch(ctx, items)
}

func (s *LPARecordItemService) ListByRecord(ctx context.Context, recordID uint) ([]model.LPARecordItem, error) {
	return s.repo.ListByRecord(ctx, recordID)
}
