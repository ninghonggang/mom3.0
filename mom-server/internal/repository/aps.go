package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type MPSRepository struct {
	db *gorm.DB
}

func NewMPSRepository(db *gorm.DB) *MPSRepository {
	return &MPSRepository{db: db}
}

func (r *MPSRepository) List(ctx context.Context, tenantID int64) ([]model.MPS, int64, error) {
	var list []model.MPS
	var total int64
	err := r.db.WithContext(ctx).Model(&model.MPS{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *MPSRepository) GetByID(ctx context.Context, id uint) (*model.MPS, error) {
	var mps model.MPS
	err := r.db.WithContext(ctx).First(&mps, id).Error
	return &mps, err
}

func (r *MPSRepository) Create(ctx context.Context, mps *model.MPS) error {
	return r.db.WithContext(ctx).Create(mps).Error
}

func (r *MPSRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.MPS{}).Where("id = ?", id).Updates(updates).Error
}

func (r *MPSRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.MPS{}, id).Error
}

type MRPRepository struct {
	db *gorm.DB
}

func NewMRPRepository(db *gorm.DB) *MRPRepository {
	return &MRPRepository{db: db}
}

func (r *MRPRepository) List(ctx context.Context, tenantID int64) ([]model.MRP, int64, error) {
	var list []model.MRP
	var total int64
	err := r.db.WithContext(ctx).Model(&model.MRP{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *MRPRepository) GetByID(ctx context.Context, id uint) (*model.MRP, error) {
	var mrp model.MRP
	err := r.db.WithContext(ctx).First(&mrp, id).Error
	return &mrp, err
}

func (r *MRPRepository) Create(ctx context.Context, mrp *model.MRP) error {
	return r.db.WithContext(ctx).Create(mrp).Error
}

func (r *MRPRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.MRP{}).Where("id = ?", id).Updates(updates).Error
}

func (r *MRPRepository) CreateItem(ctx context.Context, item *model.MRPItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *MRPRepository) GetItemsByMRPID(ctx context.Context, mrpID int64) ([]model.MRPItem, error) {
	var items []model.MRPItem
	err := r.db.WithContext(ctx).Where("mrp_id = ?", mrpID).Find(&items).Error
	return items, err
}

type ScheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) *ScheduleRepository {
	return &ScheduleRepository{db: db}
}

func (r *ScheduleRepository) List(ctx context.Context, tenantID int64) ([]model.SchedulePlan, int64, error) {
	var list []model.SchedulePlan
	var total int64
	err := r.db.WithContext(ctx).Model(&model.SchedulePlan{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *ScheduleRepository) GetByID(ctx context.Context, id uint) (*model.SchedulePlan, error) {
	var plan model.SchedulePlan
	err := r.db.WithContext(ctx).First(&plan, id).Error
	return &plan, err
}

func (r *ScheduleRepository) Create(ctx context.Context, plan *model.SchedulePlan) error {
	return r.db.WithContext(ctx).Create(plan).Error
}

func (r *ScheduleRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.SchedulePlan{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ScheduleRepository) CreateResult(ctx context.Context, result *model.ScheduleResult) error {
	return r.db.WithContext(ctx).Create(result).Error
}

func (r *ScheduleRepository) GetResultsByPlanID(ctx context.Context, planID int64) ([]model.ScheduleResult, error) {
	var results []model.ScheduleResult
	err := r.db.WithContext(ctx).Where("plan_id = ?", planID).Order("sequence ASC").Find(&results).Error
	return results, err
}

func (r *ScheduleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.SchedulePlan{}, id).Error
}

func (r *ScheduleRepository) UpdateResult(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ScheduleResult{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ScheduleRepository) GetResultByID(ctx context.Context, id uint) (*model.ScheduleResult, error) {
	var result model.ScheduleResult
	err := r.db.WithContext(ctx).First(&result, id).Error
	return &result, err
}
