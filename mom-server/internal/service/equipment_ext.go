package service

import (
	"context"
	"math"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// ========== TEEP分析 ==========

type TEEPDataService struct {
	repo *repository.TEEPDataRepository
}

func NewTEEPDataService(repo *repository.TEEPDataRepository) *TEEPDataService {
	return &TEEPDataService{repo: repo}
}

func (s *TEEPDataService) List(ctx context.Context, equipmentID int64, startDate, endDate string) ([]model.TEEPData, int64, error) {
	return s.repo.List(ctx, 1, equipmentID, startDate, endDate)
}

func (s *TEEPDataService) GetByID(ctx context.Context, id uint) (*model.TEEPData, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TEEPDataService) Create(ctx context.Context, data *model.TEEPData) error {
	if data.TenantID == 0 {
		data.TenantID = 1
	}
	// 自动计算OEE和TEEP
	s.calculateOEE(data)
	return s.repo.Create(ctx, data)
}

func (s *TEEPDataService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *TEEPDataService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

// calculateOEE 计算OEE和TEEP
func (s *TEEPDataService) calculateOEE(data *model.TEEPData) {
	// 时间稼动率
	if data.PlanTime > 0 {
		data.Availability = data.ActualTime / data.PlanTime
	} else {
		data.Availability = 0
	}

	// 性能稼动率
	if data.ActualOutput > 0 && data.IdealOutput > 0 {
		data.Performance = data.IdealOutput / float64(data.ActualOutput)
	} else {
		data.Performance = 0
	}
	if data.Performance > 1 {
		data.Performance = 1
	}

	// 良品率
	if data.ActualOutput > 0 {
		data.Quality = float64(data.PassOutput) / float64(data.ActualOutput)
	} else {
		data.Quality = 0
	}

	// OEE = 性能 * 质量 * 时间
	data.OEE = data.Performance * data.Quality * data.Availability * 100

	// TEEP = OEE * 时间稼动率 (简化计算)
	data.TEEP = data.OEE * data.Availability / 100
}

// CalculateOEEFromData 从原始数据计算OEE
func CalculateOEE(planTime, downTime float64, actualOutput, passOutput int64, idealCycleTime float64) (oee, availability, performance, quality float64) {
	// 时间稼动率
	actualTime := planTime - downTime
	if planTime > 0 {
		availability = actualTime / planTime
	}

	// 理论产量
	idealOutput := (actualTime * 60 / idealCycleTime) * 1000 // 转换为每小时产量

	// 性能稼动率
	if actualOutput > 0 && idealOutput > 0 {
		performance = math.Min(idealOutput/float64(actualOutput), 1.0)
	}

	// 良品率
	if actualOutput > 0 {
		quality = float64(passOutput) / float64(actualOutput)
	}

	// OEE
	oee = availability * performance * quality * 100
	return
}

// ========== 模具管理 ==========

type MoldService struct {
	repo *repository.MoldRepository
}

func NewMoldService(repo *repository.MoldRepository) *MoldService {
	return &MoldService{repo: repo}
}

func (s *MoldService) List(ctx context.Context, query string) ([]model.Mold, int64, error) {
	return s.repo.List(ctx, 1, query)
}

func (s *MoldService) GetByID(ctx context.Context, id uint) (*model.Mold, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *MoldService) Create(ctx context.Context, mold *model.Mold) error {
	if mold.TenantID == 0 {
		mold.TenantID = 1
	}
	return s.repo.Create(ctx, mold)
}

func (s *MoldService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *MoldService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

type MoldMaintenanceService struct {
	repo *repository.MoldMaintenanceRepository
}

func NewMoldMaintenanceService(repo *repository.MoldMaintenanceRepository) *MoldMaintenanceService {
	return &MoldMaintenanceService{repo: repo}
}

func (s *MoldMaintenanceService) List(ctx context.Context, moldID int64) ([]model.MoldMaintenance, int64, error) {
	return s.repo.List(ctx, 1, moldID)
}

func (s *MoldMaintenanceService) Create(ctx context.Context, m *model.MoldMaintenance) error {
	if m.TenantID == 0 {
		m.TenantID = 1
	}
	return s.repo.Create(ctx, m)
}

type MoldRepairService struct {
	repo *repository.MoldRepairRepository
}

func NewMoldRepairService(repo *repository.MoldRepairRepository) *MoldRepairService {
	return &MoldRepairService{repo: repo}
}

func (s *MoldRepairService) List(ctx context.Context, moldID int64) ([]model.MoldRepair, int64, error) {
	return s.repo.List(ctx, 1, moldID)
}

func (s *MoldRepairService) Create(ctx context.Context, m *model.MoldRepair) error {
	if m.TenantID == 0 {
		m.TenantID = 1
	}
	return s.repo.Create(ctx, m)
}

// ========== 量检具管理 ==========

type GaugeService struct {
	repo *repository.GaugeRepository
}

func NewGaugeService(repo *repository.GaugeRepository) *GaugeService {
	return &GaugeService{repo: repo}
}

func (s *GaugeService) List(ctx context.Context, query string) ([]model.Gauge, int64, error) {
	return s.repo.List(ctx, 1, query)
}

func (s *GaugeService) GetByID(ctx context.Context, id uint) (*model.Gauge, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *GaugeService) Create(ctx context.Context, gauge *model.Gauge) error {
	if gauge.TenantID == 0 {
		gauge.TenantID = 1
	}
	return s.repo.Create(ctx, gauge)
}

func (s *GaugeService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *GaugeService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

type GaugeCalibrationService struct {
	repo *repository.GaugeCalibrationRepository
}

func NewGaugeCalibrationService(repo *repository.GaugeCalibrationRepository) *GaugeCalibrationService {
	return &GaugeCalibrationService{repo: repo}
}

func (s *GaugeCalibrationService) List(ctx context.Context, gaugeID int64) ([]model.GaugeCalibration, int64, error) {
	return s.repo.List(ctx, 1, gaugeID)
}

func (s *GaugeCalibrationService) Create(ctx context.Context, c *model.GaugeCalibration) error {
	if c.TenantID == 0 {
		c.TenantID = 1
	}
	return s.repo.Create(ctx, c)
}
