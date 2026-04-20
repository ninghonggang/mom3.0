package service

import (
	"context"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"time"
)

// ========== 滚动排程配置 ==========

type RollingConfigService struct {
	repo *repository.RollingConfigRepository
}

func NewRollingConfigService(repo *repository.RollingConfigRepository) *RollingConfigService {
	return &RollingConfigService{repo: repo}
}

func (s *RollingConfigService) List(ctx context.Context) ([]model.RollingConfig, int64, error) {
	return s.repo.List(ctx, 1)
}

func (s *RollingConfigService) GetByID(ctx context.Context, id uint) (*model.RollingConfig, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *RollingConfigService) Create(ctx context.Context, r *model.RollingConfig) error {
	if r.TenantID == 0 {
		r.TenantID = 1
	}
	return s.repo.Create(ctx, r)
}

func (s *RollingConfigService) Update(ctx context.Context, id uint, updates map[string]any) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *RollingConfigService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *RollingConfigService) Execute(ctx context.Context, id uint) error {
	// 更新最后执行时间
	now := time.Now()
	if err := s.repo.Update(ctx, id, map[string]any{"last_execute_time": now}); err != nil {
		return err
	}
	// 模拟执行 - 实际应调用排程引擎
	return nil
}

// ========== 排程结果 ==========

type ScheduleResultService struct {
	repo      *repository.ScheduleResultRepository
	warningRepo *repository.ScheduleWarningRepository
}

func NewScheduleResultService(repo *repository.ScheduleResultRepository, warningRepo *repository.ScheduleWarningRepository) *ScheduleResultService {
	return &ScheduleResultService{repo: repo, warningRepo: warningRepo}
}

func (s *ScheduleResultService) List(ctx context.Context) ([]model.RollingScheduleResult, int64, error) {
	return s.repo.List(ctx, 1)
}

func (s *ScheduleResultService) GetByID(ctx context.Context, id uint) (*model.RollingScheduleResult, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ScheduleResultService) Create(ctx context.Context, r *model.RollingScheduleResult) error {
	if r.TenantID == 0 {
		r.TenantID = 1
	}
	return s.repo.Create(ctx, r)
}

func (s *ScheduleResultService) Execute(ctx context.Context, configID uint) error {
	// 模拟执行排程 - 实际需要调用排程引擎
	now := time.Now()
	result := &model.RollingScheduleResult{
		ScheduleNo:    "SCH" + now.Format("20060102150405"),
		ConfigID:      int64Ptr(int64(configID)),
		ScheduleDate:  now.Format("2006-01-02"),
		Status:        "COMPLETED",
		TotalOrders:   0,
		ScheduledOrders: 0,
		TenantID:      1,
	}
	return s.repo.Create(ctx, result)
}

func (s *ScheduleResultService) GetWarnings(ctx context.Context, resultID uint) ([]model.ScheduleWarning, error) {
	if s.warningRepo == nil {
		return []model.ScheduleWarning{}, nil
	}
	return s.warningRepo.ListByResultID(ctx, resultID)
}

// ========== 交付分析 ==========

type DeliveryAnalysisService struct {
	repo *repository.DeliveryAnalysisRepository
}

func NewDeliveryAnalysisService(repo *repository.DeliveryAnalysisRepository) *DeliveryAnalysisService {
	return &DeliveryAnalysisService{repo: repo}
}

func (s *DeliveryAnalysisService) List(ctx context.Context) ([]model.DeliveryAnalysis, int64, error) {
	return s.repo.List(ctx, 1)
}

func (s *DeliveryAnalysisService) GetByID(ctx context.Context, id uint) (*model.DeliveryAnalysis, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *DeliveryAnalysisService) Create(ctx context.Context, r *model.DeliveryAnalysis) error {
	if r.TenantID == 0 {
		r.TenantID = 1
	}
	return s.repo.Create(ctx, r)
}

func (s *DeliveryAnalysisService) GetDailyStats(ctx context.Context) ([]model.DeliveryAnalysis, error) {
	return s.repo.GetStatsByDate(ctx, 1, time.Now().Format("2006-01-02"), time.Now().Format("2006-01-02"))
}

func (s *DeliveryAnalysisService) GetWeeklyStats(ctx context.Context) ([]model.DeliveryAnalysis, error) {
	now := time.Now()
	start := now.AddDate(0, 0, -7).Format("2006-01-02")
	return s.repo.GetStatsByDate(ctx, 1, start, now.Format("2006-01-02"))
}

func (s *DeliveryAnalysisService) GetMonthlyStats(ctx context.Context) ([]model.DeliveryAnalysis, error) {
	now := time.Now()
	start := now.AddDate(0, -1, 0).Format("2006-01-02")
	return s.repo.GetStatsByDate(ctx, 1, start, now.Format("2006-01-02"))
}

// ========== 交付预警 ==========

type DeliveryWarningService struct {
	repo *repository.DeliveryWarningRepository
}

func NewDeliveryWarningService(repo *repository.DeliveryWarningRepository) *DeliveryWarningService {
	return &DeliveryWarningService{repo: repo}
}

func (s *DeliveryWarningService) List(ctx context.Context) ([]model.DeliveryWarning, int64, error) {
	return s.repo.List(ctx, 1)
}

func (s *DeliveryWarningService) GetByID(ctx context.Context, id uint) (*model.DeliveryWarning, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *DeliveryWarningService) Acknowledge(ctx context.Context, id uint) error {
	now := time.Now()
	return s.repo.Update(ctx, id, map[string]any{
		"status":            "ACKNOWLEDGED",
		"acknowledged_time": now,
	})
}

func (s *DeliveryWarningService) Mitigate(ctx context.Context, id uint, action string) error {
	return s.repo.Update(ctx, id, map[string]any{
		"status": "MITIGATED",
	})
}

// ========== 缺料分析 ==========

type MaterialShortageService struct {
	repo *repository.MaterialShortageRepository
}

func NewMaterialShortageService(repo *repository.MaterialShortageRepository) *MaterialShortageService {
	return &MaterialShortageService{repo: repo}
}

func (s *MaterialShortageService) List(ctx context.Context) ([]model.MaterialShortage, int64, error) {
	return s.repo.List(ctx, 1)
}

func (s *MaterialShortageService) GetByID(ctx context.Context, id uint) (*model.MaterialShortage, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *MaterialShortageService) Create(ctx context.Context, r *model.MaterialShortage) error {
	if r.TenantID == 0 {
		r.TenantID = 1
	}
	return s.repo.Create(ctx, r)
}

func (s *MaterialShortageService) GetDailyStats(ctx context.Context) ([]model.MaterialShortage, error) {
	return s.repo.GetByDate(ctx, 1, time.Now().Format("2006-01-02"))
}

// ========== 缺料预警规则 ==========

type ShortageWarningRuleService struct {
	repo *repository.ShortageWarningRuleRepository
}

func NewShortageWarningRuleService(repo *repository.ShortageWarningRuleRepository) *ShortageWarningRuleService {
	return &ShortageWarningRuleService{repo: repo}
}

func (s *ShortageWarningRuleService) List(ctx context.Context) ([]model.ShortageWarningRule, int64, error) {
	return s.repo.List(ctx, 1)
}

func (s *ShortageWarningRuleService) GetByID(ctx context.Context, id uint) (*model.ShortageWarningRule, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ShortageWarningRuleService) Create(ctx context.Context, r *model.ShortageWarningRule) error {
	if r.TenantID == 0 {
		r.TenantID = 1
	}
	return s.repo.Create(ctx, r)
}

func (s *ShortageWarningRuleService) Update(ctx context.Context, id uint, updates map[string]any) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *ShortageWarningRuleService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

// ========== APS班次 ==========

type APSShiftService struct {
	repo *repository.APSShiftRepository
}

func NewAPSShiftService(repo *repository.APSShiftRepository) *APSShiftService {
	return &APSShiftService{repo: repo}
}

func (s *APSShiftService) List(ctx context.Context) ([]model.APSShift, int64, error) {
	return s.repo.List(ctx, 1)
}

func (s *APSShiftService) GetByID(ctx context.Context, id uint) (*model.APSShift, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *APSShiftService) Create(ctx context.Context, r *model.APSShift) error {
	if r.TenantID == 0 {
		r.TenantID = 1
	}
	return s.repo.Create(ctx, r)
}

func (s *APSShiftService) Update(ctx context.Context, id uint, updates map[string]any) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *APSShiftService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

// ========== APS换型矩阵 ==========

type APSChangeoverService struct {
	repo *repository.APSChangeoverRepository
}

func NewAPSChangeoverService(repo *repository.APSChangeoverRepository) *APSChangeoverService {
	return &APSChangeoverService{repo: repo}
}

func (s *APSChangeoverService) List(ctx context.Context) ([]model.APSChangeoverMatrix, error) {
	return s.repo.List(ctx, 1)
}

func (s *APSChangeoverService) GetByID(ctx context.Context, id uint) (*model.APSChangeoverMatrix, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *APSChangeoverService) Create(ctx context.Context, r *model.APSChangeoverMatrix) error {
	if r.TenantID == 0 {
		r.TenantID = 1
	}
	return s.repo.Create(ctx, r)
}

func (s *APSChangeoverService) Update(ctx context.Context, id uint, updates map[string]any) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *APSChangeoverService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *APSChangeoverService) GetByProducts(ctx context.Context, fromProductID, toProductID int64) (*model.APSChangeoverMatrix, error) {
	return s.repo.GetByProducts(ctx, fromProductID, toProductID)
}

// ========== APS产品族 ==========

type ProductFamilyService struct {
	repo *repository.ProductFamilyRepository
}

func NewProductFamilyService(repo *repository.ProductFamilyRepository) *ProductFamilyService {
	return &ProductFamilyService{repo: repo}
}

func (s *ProductFamilyService) List(ctx context.Context) ([]model.ProductFamily, int64, error) {
	return s.repo.List(ctx, 1)
}

func (s *ProductFamilyService) GetByID(ctx context.Context, id uint) (*model.ProductFamily, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ProductFamilyService) Create(ctx context.Context, r *model.ProductFamily) error {
	if r.TenantID == 0 {
		r.TenantID = 1
	}
	return s.repo.Create(ctx, r)
}

func (s *ProductFamilyService) Update(ctx context.Context, id uint, updates map[string]any) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *ProductFamilyService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

// ========== 辅助函数 ==========

func int64Ptr(i int64) *int64 {
	return &i
}
