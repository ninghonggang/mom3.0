package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type WorkshopRepository struct {
	db *gorm.DB
}

func NewWorkshopRepository(db *gorm.DB) *WorkshopRepository {
	return &WorkshopRepository{db: db}
}

func (r *WorkshopRepository) List(ctx context.Context, tenantID int64) ([]model.Workshop, int64, error) {
	var list []model.Workshop
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Workshop{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *WorkshopRepository) GetByID(ctx context.Context, id uint) (*model.Workshop, error) {
	var workshop model.Workshop
	err := r.db.WithContext(ctx).First(&workshop, id).Error
	return &workshop, err
}

func (r *WorkshopRepository) Create(ctx context.Context, workshop *model.Workshop) error {
	return r.db.WithContext(ctx).Create(workshop).Error
}

func (r *WorkshopRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Workshop{}).Where("id = ?", id).Updates(updates).Error
}

func (r *WorkshopRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Workshop{}, id).Error
}

type ProductionLineRepository struct {
	db *gorm.DB
}

func NewProductionLineRepository(db *gorm.DB) *ProductionLineRepository {
	return &ProductionLineRepository{db: db}
}

func (r *ProductionLineRepository) List(ctx context.Context, tenantID int64) ([]model.ProductionLine, int64, error) {
	var list []model.ProductionLine
	var total int64

	err := r.db.WithContext(ctx).Model(&model.ProductionLine{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *ProductionLineRepository) GetByID(ctx context.Context, id uint) (*model.ProductionLine, error) {
	var line model.ProductionLine
	err := r.db.WithContext(ctx).First(&line, id).Error
	return &line, err
}

func (r *ProductionLineRepository) Create(ctx context.Context, line *model.ProductionLine) error {
	return r.db.WithContext(ctx).Create(line).Error
}

func (r *ProductionLineRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ProductionLine{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ProductionLineRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.ProductionLine{}, id).Error
}

type ShiftRepository struct {
	db *gorm.DB
}

func NewShiftRepository(db *gorm.DB) *ShiftRepository {
	return &ShiftRepository{db: db}
}

func (r *ShiftRepository) List(ctx context.Context, tenantID int64) ([]model.Shift, int64, error) {
	var list []model.Shift
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Shift{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *ShiftRepository) GetByID(ctx context.Context, id uint) (*model.Shift, error) {
	var shift model.Shift
	err := r.db.WithContext(ctx).First(&shift, id).Error
	return &shift, err
}

func (r *ShiftRepository) Create(ctx context.Context, shift *model.Shift) error {
	return r.db.WithContext(ctx).Create(shift).Error
}

func (r *ShiftRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Shift{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ShiftRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Shift{}, id).Error
}
