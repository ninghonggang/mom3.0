package repository

import (
	"context"

	"gorm.io/gorm"

	"mom-server/internal/model"
)

// WorkshopConfigRepository 车间配置Repository
type WorkshopConfigRepository struct {
	db *gorm.DB
}

// NewWorkshopConfigRepository 创建车间配置Repository
func NewWorkshopConfigRepository(db *gorm.DB) *WorkshopConfigRepository {
	return &WorkshopConfigRepository{db: db}
}

// GetByWorkshopID 根据车间ID获取配置
func (r *WorkshopConfigRepository) GetByWorkshopID(ctx context.Context, workshopID int64) (*model.WorkshopConfig, error) {
	var config model.WorkshopConfig
	err := r.db.WithContext(ctx).Where("workshop_id = ?", workshopID).First(&config).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// Create 创建车间配置
func (r *WorkshopConfigRepository) Create(ctx context.Context, config *model.WorkshopConfig) error {
	return r.db.WithContext(ctx).Create(config).Error
}

// Update 更新车间配置
func (r *WorkshopConfigRepository) Update(ctx context.Context, workshopID int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.WorkshopConfig{}).Where("workshop_id = ?", workshopID).Updates(updates).Error
}

// WorkingCalendarRepository 工厂日历Repository
type WorkingCalendarRepository struct {
	db *gorm.DB
}

// NewWorkingCalendarRepository 创建工厂日历Repository
func NewWorkingCalendarRepository(db *gorm.DB) *WorkingCalendarRepository {
	return &WorkingCalendarRepository{db: db}
}

// GetByWorkshopID 获取车间日历列表
func (r *WorkingCalendarRepository) GetByWorkshopID(ctx context.Context, workshopID int64) ([]model.WorkingCalendar, error) {
	var calendars []model.WorkingCalendar
	err := r.db.WithContext(ctx).Where("workshop_id = ?", workshopID).Order("created_at desc").Find(&calendars).Error
	return calendars, err
}

// GetEffective 获取生效日历
func (r *WorkingCalendarRepository) GetEffective(ctx context.Context, workshopID int64) (*model.WorkingCalendar, error) {
	var calendar model.WorkingCalendar
	err := r.db.WithContext(ctx).Where("workshop_id = ? AND status = 1", workshopID).Order("effective_from desc").First(&calendar).Error
	if err != nil {
		return nil, err
	}
	return &calendar, nil
}

// Create 创建日历
func (r *WorkingCalendarRepository) Create(ctx context.Context, calendar *model.WorkingCalendar) error {
	return r.db.WithContext(ctx).Create(calendar).Error
}

// Update 更新日历
func (r *WorkingCalendarRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.WorkingCalendar{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除日历
func (r *WorkingCalendarRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.WorkingCalendar{}, id).Error
}
