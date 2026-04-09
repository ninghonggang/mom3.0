package service

import (
	"context"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// WorkshopConfigService 车间配置服务
type WorkshopConfigService struct {
	repo *repository.WorkshopConfigRepository
}

// NewWorkshopConfigService 创建车间配置服务
func NewWorkshopConfigService(repo *repository.WorkshopConfigRepository) *WorkshopConfigService {
	return &WorkshopConfigService{repo: repo}
}

// GetByWorkshopID 获取车间配置
func (s *WorkshopConfigService) GetByWorkshopID(ctx context.Context, workshopID int64) (*model.WorkshopConfig, error) {
	return s.repo.GetByWorkshopID(ctx, workshopID)
}

// Create 创建车间配置
func (s *WorkshopConfigService) Create(ctx context.Context, config *model.WorkshopConfig) error {
	return s.repo.Create(ctx, config)
}

// Update 更新车间配置
func (s *WorkshopConfigService) Update(ctx context.Context, workshopID int64, updates map[string]interface{}) error {
	return s.repo.Update(ctx, workshopID, updates)
}

// WorkingCalendarService 工厂日历服务
type WorkingCalendarService struct {
	repo *repository.WorkingCalendarRepository
}

// NewWorkingCalendarService 创建工厂日历服务
func NewWorkingCalendarService(repo *repository.WorkingCalendarRepository) *WorkingCalendarService {
	return &WorkingCalendarService{repo: repo}
}

// GetByWorkshopID 获取车间日历
func (s *WorkingCalendarService) GetByWorkshopID(ctx context.Context, workshopID int64) ([]model.WorkingCalendar, error) {
	return s.repo.GetByWorkshopID(ctx, workshopID)
}

// GetEffective 获取生效日历
func (s *WorkingCalendarService) GetEffective(ctx context.Context, workshopID int64) (*model.WorkingCalendar, error) {
	return s.repo.GetEffective(ctx, workshopID)
}

// Create 创建日历
func (s *WorkingCalendarService) Create(ctx context.Context, calendar *model.WorkingCalendar) error {
	return s.repo.Create(ctx, calendar)
}

// Update 更新日历
func (s *WorkingCalendarService) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

// Delete 删除日历
func (s *WorkingCalendarService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

// GetWorkHours 获取工作时长(小时)
func (s *WorkingCalendarService) GetWorkHours(ctx context.Context, workshopID int64) (float64, error) {
	cal, err := s.repo.GetEffective(ctx, workshopID)
	if err != nil {
		return 8.0, nil // 默认8小时
	}
	return cal.GetWorkHoursPerDay(), nil
}
