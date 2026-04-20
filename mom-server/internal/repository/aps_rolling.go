package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// ========== 滚动排程配置 ==========

type RollingConfigRepository struct {
	db *gorm.DB
}

func NewRollingConfigRepository(db *gorm.DB) *RollingConfigRepository {
	return &RollingConfigRepository{db: db}
}

func (r *RollingConfigRepository) List(ctx context.Context, tenantID int64) ([]model.RollingConfig, int64, error) {
	var list []model.RollingConfig
	var total int64
	db := r.db.WithContext(ctx).Model(&model.RollingConfig{}).Where("tenant_id = ?", tenantID)
	db.Count(&total)
	err := db.Order("created_at DESC").Find(&list).Error
	return list, total, err
}

func (r *RollingConfigRepository) GetByID(ctx context.Context, id uint) (*model.RollingConfig, error) {
	var c model.RollingConfig
	err := r.db.WithContext(ctx).First(&c, id).Error
	return &c, err
}

func (r *RollingConfigRepository) Create(ctx context.Context, c *model.RollingConfig) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *RollingConfigRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.RollingConfig{}).Where("id = ?", id).Updates(updates).Error
}

func (r *RollingConfigRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.RollingConfig{}, id).Error
}

// ========== 排程结果 ==========

type ScheduleResultRepository struct {
	db *gorm.DB
}

func NewScheduleResultRepository(db *gorm.DB) *ScheduleResultRepository {
	return &ScheduleResultRepository{db: db}
}

func (r *ScheduleResultRepository) List(ctx context.Context, tenantID int64) ([]model.RollingScheduleResult, int64, error) {
	var list []model.RollingScheduleResult
	var total int64
	db := r.db.WithContext(ctx).Model(&model.RollingScheduleResult{}).Where("tenant_id = ?", tenantID)
	db.Count(&total)
	err := db.Order("created_at DESC").Find(&list).Error
	return list, total, err
}

func (r *ScheduleResultRepository) GetByID(ctx context.Context, id uint) (*model.RollingScheduleResult, error) {
	var s model.RollingScheduleResult
	err := r.db.WithContext(ctx).First(&s, id).Error
	return &s, err
}

func (r *ScheduleResultRepository) Create(ctx context.Context, s *model.RollingScheduleResult) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *ScheduleResultRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.RollingScheduleResult{}).Where("id = ?", id).Updates(updates).Error
}

// ========== 排程警告 ==========

type ScheduleWarningRepository struct {
	db *gorm.DB
}

func NewScheduleWarningRepository(db *gorm.DB) *ScheduleWarningRepository {
	return &ScheduleWarningRepository{db: db}
}

func (r *ScheduleWarningRepository) ListByResultID(ctx context.Context, resultID uint) ([]model.ScheduleWarning, error) {
	var list []model.ScheduleWarning
	err := r.db.WithContext(ctx).Where("schedule_result_id = ?", resultID).Find(&list).Error
	return list, err
}

// ========== 交付分析 ==========

type DeliveryAnalysisRepository struct {
	db *gorm.DB
}

func NewDeliveryAnalysisRepository(db *gorm.DB) *DeliveryAnalysisRepository {
	return &DeliveryAnalysisRepository{db: db}
}

func (r *DeliveryAnalysisRepository) List(ctx context.Context, tenantID int64) ([]model.DeliveryAnalysis, int64, error) {
	var list []model.DeliveryAnalysis
	var total int64
	db := r.db.WithContext(ctx).Model(&model.DeliveryAnalysis{}).Where("tenant_id = ?", tenantID)
	db.Count(&total)
	err := db.Order("analysis_date DESC").Find(&list).Error
	return list, total, err
}

func (r *DeliveryAnalysisRepository) GetByID(ctx context.Context, id uint) (*model.DeliveryAnalysis, error) {
	var d model.DeliveryAnalysis
	err := r.db.WithContext(ctx).First(&d, id).Error
	return &d, err
}

func (r *DeliveryAnalysisRepository) Create(ctx context.Context, d *model.DeliveryAnalysis) error {
	return r.db.WithContext(ctx).Create(d).Error
}

func (r *DeliveryAnalysisRepository) GetStatsByDate(ctx context.Context, tenantID int64, startDate, endDate string) ([]model.DeliveryAnalysis, error) {
	var list []model.DeliveryAnalysis
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND analysis_date >= ? AND analysis_date <= ?",
		tenantID, startDate, endDate).Order("analysis_date DESC").Find(&list).Error
	return list, err
}

// ========== 交付预警 ==========

type DeliveryWarningRepository struct {
	db *gorm.DB
}

func NewDeliveryWarningRepository(db *gorm.DB) *DeliveryWarningRepository {
	return &DeliveryWarningRepository{db: db}
}

func (r *DeliveryWarningRepository) List(ctx context.Context, tenantID int64) ([]model.DeliveryWarning, int64, error) {
	var list []model.DeliveryWarning
	var total int64
	db := r.db.WithContext(ctx).Model(&model.DeliveryWarning{}).Where("tenant_id = ?", tenantID)
	db.Count(&total)
	err := db.Order("created_at DESC").Find(&list).Error
	return list, total, err
}

func (r *DeliveryWarningRepository) GetByID(ctx context.Context, id uint) (*model.DeliveryWarning, error) {
	var d model.DeliveryWarning
	err := r.db.WithContext(ctx).First(&d, id).Error
	return &d, err
}

func (r *DeliveryWarningRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.DeliveryWarning{}).Where("id = ?", id).Updates(updates).Error
}

// ========== 缺料分析 ==========

type MaterialShortageRepository struct {
	db *gorm.DB
}

func NewMaterialShortageRepository(db *gorm.DB) *MaterialShortageRepository {
	return &MaterialShortageRepository{db: db}
}

func (r *MaterialShortageRepository) List(ctx context.Context, tenantID int64) ([]model.MaterialShortage, int64, error) {
	var list []model.MaterialShortage
	var total int64
	db := r.db.WithContext(ctx).Model(&model.MaterialShortage{}).Where("tenant_id = ?", tenantID)
	db.Count(&total)
	err := db.Order("created_at DESC").Find(&list).Error
	return list, total, err
}

func (r *MaterialShortageRepository) GetByID(ctx context.Context, id uint) (*model.MaterialShortage, error) {
	var m model.MaterialShortage
	err := r.db.WithContext(ctx).First(&m, id).Error
	return &m, err
}

func (r *MaterialShortageRepository) Create(ctx context.Context, m *model.MaterialShortage) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *MaterialShortageRepository) GetByDate(ctx context.Context, tenantID int64, date string) ([]model.MaterialShortage, error) {
	var list []model.MaterialShortage
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND analysis_date = ?", tenantID, date).Find(&list).Error
	return list, err
}

// ========== 缺料预警规则 ==========

type ShortageWarningRuleRepository struct {
	db *gorm.DB
}

func NewShortageWarningRuleRepository(db *gorm.DB) *ShortageWarningRuleRepository {
	return &ShortageWarningRuleRepository{db: db}
}

func (r *ShortageWarningRuleRepository) List(ctx context.Context, tenantID int64) ([]model.ShortageWarningRule, int64, error) {
	var list []model.ShortageWarningRule
	var total int64
	db := r.db.WithContext(ctx).Model(&model.ShortageWarningRule{}).Where("tenant_id = ?", tenantID)
	db.Count(&total)
	err := db.Order("created_at DESC").Find(&list).Error
	return list, total, err
}

func (r *ShortageWarningRuleRepository) GetByID(ctx context.Context, id uint) (*model.ShortageWarningRule, error) {
	var s model.ShortageWarningRule
	err := r.db.WithContext(ctx).First(&s, id).Error
	return &s, err
}

func (r *ShortageWarningRuleRepository) Create(ctx context.Context, s *model.ShortageWarningRule) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *ShortageWarningRuleRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.ShortageWarningRule{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ShortageWarningRuleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.ShortageWarningRule{}, id).Error
}

// ========== APS班次 ==========

type APSShiftRepository struct {
	db *gorm.DB
}

func NewAPSShiftRepository(db *gorm.DB) *APSShiftRepository {
	return &APSShiftRepository{db: db}
}

func (r *APSShiftRepository) List(ctx context.Context, tenantID int64) ([]model.APSShift, int64, error) {
	var list []model.APSShift
	var total int64
	db := r.db.WithContext(ctx).Model(&model.APSShift{}).Where("tenant_id = ?", tenantID)
	db.Count(&total)
	err := db.Order("sort_order ASC").Find(&list).Error
	return list, total, err
}

func (r *APSShiftRepository) GetByID(ctx context.Context, id uint) (*model.APSShift, error) {
	var s model.APSShift
	err := r.db.WithContext(ctx).First(&s, id).Error
	return &s, err
}

func (r *APSShiftRepository) Create(ctx context.Context, s *model.APSShift) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *APSShiftRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.APSShift{}).Where("id = ?", id).Updates(updates).Error
}

func (r *APSShiftRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.APSShift{}, id).Error
}

// ========== APS换型矩阵 ==========

type APSChangeoverRepository struct {
	db *gorm.DB
}

func NewAPSChangeoverRepository(db *gorm.DB) *APSChangeoverRepository {
	return &APSChangeoverRepository{db: db}
}

func (r *APSChangeoverRepository) List(ctx context.Context, tenantID int64) ([]model.APSChangeoverMatrix, error) {
	var list []model.APSChangeoverMatrix
	err := r.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Find(&list).Error
	return list, err
}

func (r *APSChangeoverRepository) GetByID(ctx context.Context, id uint) (*model.APSChangeoverMatrix, error) {
	var c model.APSChangeoverMatrix
	err := r.db.WithContext(ctx).First(&c, id).Error
	return &c, err
}

func (r *APSChangeoverRepository) Create(ctx context.Context, c *model.APSChangeoverMatrix) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *APSChangeoverRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.APSChangeoverMatrix{}).Where("id = ?", id).Updates(updates).Error
}

func (r *APSChangeoverRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.APSChangeoverMatrix{}, id).Error
}

func (r *APSChangeoverRepository) GetByProducts(ctx context.Context, fromProductID, toProductID int64) (*model.APSChangeoverMatrix, error) {
	var c model.APSChangeoverMatrix
	err := r.db.WithContext(ctx).Where("from_product_id = ? AND to_product_id = ?", fromProductID, toProductID).First(&c).Error
	return &c, err
}

// ========== APS产品族 ==========

type ProductFamilyRepository struct {
	db *gorm.DB
}

func NewProductFamilyRepository(db *gorm.DB) *ProductFamilyRepository {
	return &ProductFamilyRepository{db: db}
}

func (r *ProductFamilyRepository) List(ctx context.Context, tenantID int64) ([]model.ProductFamily, int64, error) {
	var list []model.ProductFamily
	var total int64
	db := r.db.WithContext(ctx).Model(&model.ProductFamily{}).Where("tenant_id = ?", tenantID)
	db.Count(&total)
	err := db.Order("created_at DESC").Find(&list).Error
	return list, total, err
}

func (r *ProductFamilyRepository) GetByID(ctx context.Context, id uint) (*model.ProductFamily, error) {
	var p model.ProductFamily
	err := r.db.WithContext(ctx).First(&p, id).Error
	return &p, err
}

func (r *ProductFamilyRepository) Create(ctx context.Context, p *model.ProductFamily) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *ProductFamilyRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.ProductFamily{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ProductFamilyRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.ProductFamily{}, id).Error
}
