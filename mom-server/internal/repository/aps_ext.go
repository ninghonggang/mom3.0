package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type CapacityAnalysisRepository struct {
	db *gorm.DB
}

func NewCapacityAnalysisRepository(db *gorm.DB) *CapacityAnalysisRepository {
	return &CapacityAnalysisRepository{db: db}
}

func (r *CapacityAnalysisRepository) List(ctx context.Context, tenantID int64, query string) ([]model.CapacityAnalysis, int64, error) {
	var list []model.CapacityAnalysis
	var total int64

	db := r.db.WithContext(ctx).Model(&model.CapacityAnalysis{}).Where("tenant_id = ?", tenantID)

	if query != "" {
		db = db.Where("workshop_name LIKE ? OR line_name LIKE ? OR work_date LIKE ?",
			"%"+query+"%", "%"+query+"%", "%"+query+"%")
	}

	db.Count(&total)
	err := db.Order("work_date DESC").Find(&list).Error
	return list, total, err
}

func (r *CapacityAnalysisRepository) GetByID(ctx context.Context, id uint) (*model.CapacityAnalysis, error) {
	var a model.CapacityAnalysis
	err := r.db.WithContext(ctx).First(&a, id).Error
	return &a, err
}

func (r *CapacityAnalysisRepository) Create(ctx context.Context, a *model.CapacityAnalysis) error {
	return r.db.WithContext(ctx).Create(a).Error
}

func (r *CapacityAnalysisRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.CapacityAnalysis{}).Where("id = ?", id).Updates(updates).Error
}

func (r *CapacityAnalysisRepository) GetStatsByDate(ctx context.Context, tenantID int64, startDate, endDate string) ([]model.CapacityAnalysis, error) {
	var list []model.CapacityAnalysis
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND work_date >= ? AND work_date <= ?",
		tenantID, startDate, endDate).Order("work_date DESC").Find(&list).Error
	return list, err
}

type DeliveryRateRepository struct {
	db *gorm.DB
}

func NewDeliveryRateRepository(db *gorm.DB) *DeliveryRateRepository {
	return &DeliveryRateRepository{db: db}
}

func (r *DeliveryRateRepository) List(ctx context.Context, tenantID int64, query string) ([]model.DeliveryRate, int64, error) {
	var list []model.DeliveryRate
	var total int64

	db := r.db.WithContext(ctx).Model(&model.DeliveryRate{}).Where("tenant_id = ?", tenantID)

	if query != "" {
		db = db.Where("order_no LIKE ? OR customer_name LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	db.Count(&total)
	err := db.Order("created_at DESC").Find(&list).Error
	return list, total, err
}

func (r *DeliveryRateRepository) GetByID(ctx context.Context, id uint) (*model.DeliveryRate, error) {
	var d model.DeliveryRate
	err := r.db.WithContext(ctx).First(&d, id).Error
	return &d, err
}

func (r *DeliveryRateRepository) Create(ctx context.Context, d *model.DeliveryRate) error {
	return r.db.WithContext(ctx).Create(d).Error
}

func (r *DeliveryRateRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.DeliveryRate{}).Where("id = ?", id).Updates(updates).Error
}

func (r *DeliveryRateRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.DeliveryRate{}, id).Error
}

type ChangeoverMatrixRepository struct {
	db *gorm.DB
}

func NewChangeoverMatrixRepository(db *gorm.DB) *ChangeoverMatrixRepository {
	return &ChangeoverMatrixRepository{db: db}
}

func (r *ChangeoverMatrixRepository) List(ctx context.Context, tenantID int64) ([]model.ChangeoverMatrix, error) {
	var list []model.ChangeoverMatrix
	err := r.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Order("priority ASC").Find(&list).Error
	return list, err
}

func (r *ChangeoverMatrixRepository) GetByID(ctx context.Context, id uint) (*model.ChangeoverMatrix, error) {
	var c model.ChangeoverMatrix
	err := r.db.WithContext(ctx).First(&c, id).Error
	return &c, err
}

func (r *ChangeoverMatrixRepository) Create(ctx context.Context, c *model.ChangeoverMatrix) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *ChangeoverMatrixRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.ChangeoverMatrix{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ChangeoverMatrixRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.ChangeoverMatrix{}, id).Error
}

func (r *ChangeoverMatrixRepository) GetByProducts(ctx context.Context, fromProductID, toProductID int64) (*model.ChangeoverMatrix, error) {
	var c model.ChangeoverMatrix
	err := r.db.WithContext(ctx).Where("from_product_id = ? AND to_product_id = ?", fromProductID, toProductID).First(&c).Error
	return &c, err
}

type RollingScheduleRepository struct {
	db *gorm.DB
}

func NewRollingScheduleRepository(db *gorm.DB) *RollingScheduleRepository {
	return &RollingScheduleRepository{db: db}
}

func (r *RollingScheduleRepository) List(ctx context.Context, tenantID int64, query string) ([]model.RollingSchedule, int64, error) {
	var list []model.RollingSchedule
	var total int64

	db := r.db.WithContext(ctx).Model(&model.RollingSchedule{}).Where("tenant_id = ?", tenantID)

	if query != "" {
		db = db.Where("schedule_no LIKE ? OR workshop_name LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	db.Count(&total)
	err := db.Order("created_at DESC").Find(&list).Error
	return list, total, err
}

func (r *RollingScheduleRepository) GetByID(ctx context.Context, id uint) (*model.RollingSchedule, error) {
	var r2 model.RollingSchedule
	err := r.db.WithContext(ctx).First(&r2, id).Error
	return &r2, err
}

func (r *RollingScheduleRepository) Create(ctx context.Context, r2 *model.RollingSchedule) error {
	return r.db.WithContext(ctx).Create(r2).Error
}

func (r *RollingScheduleRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.RollingSchedule{}).Where("id = ?", id).Updates(updates).Error
}

func (r *RollingScheduleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.RollingSchedule{}, id).Error
}

type JITDemandRepository struct {
	db *gorm.DB
}

func NewJITDemandRepository(db *gorm.DB) *JITDemandRepository {
	return &JITDemandRepository{db: db}
}

func (r *JITDemandRepository) List(ctx context.Context, tenantID int64, query string) ([]model.JITDemand, int64, error) {
	var list []model.JITDemand
	var total int64

	db := r.db.WithContext(ctx).Model(&model.JITDemand{}).Where("tenant_id = ?", tenantID)

	if query != "" {
		db = db.Where("demand_no LIKE ? OR material_code LIKE ? OR customer_name LIKE ?",
			"%"+query+"%", "%"+query+"%", "%"+query+"%")
	}

	db.Count(&total)
	err := db.Order("priority DESC, demand_time ASC").Find(&list).Error
	return list, total, err
}

func (r *JITDemandRepository) GetByID(ctx context.Context, id uint) (*model.JITDemand, error) {
	var j model.JITDemand
	err := r.db.WithContext(ctx).First(&j, id).Error
	return &j, err
}

func (r *JITDemandRepository) Create(ctx context.Context, j *model.JITDemand) error {
	return r.db.WithContext(ctx).Create(j).Error
}

func (r *JITDemandRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.JITDemand{}).Where("id = ?", id).Updates(updates).Error
}

func (r *JITDemandRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.JITDemand{}, id).Error
}
