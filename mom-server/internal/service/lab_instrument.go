package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// LabInstrumentService 实验室仪器服务
type LabInstrumentService struct {
	repo           *repository.LabInstrumentRepository
	calibrationRepo *repository.LabCalibrationRepository
}

func NewLabInstrumentService(repo *repository.LabInstrumentRepository, calibrationRepo *repository.LabCalibrationRepository) *LabInstrumentService {
	return &LabInstrumentService{repo: repo, calibrationRepo: calibrationRepo}
}

func (s *LabInstrumentService) List(ctx context.Context, query *model.LabInstrumentQuery) ([]model.LabInstrument, int64, error) {
	if query.Page == 0 {
		query.Page = 1
	}
	if query.PageSize == 0 {
		query.PageSize = 20
	}
	return s.repo.List(ctx, query)
}

func (s *LabInstrumentService) GetByID(ctx context.Context, id uint64) (*model.LabInstrument, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *LabInstrumentService) Create(ctx context.Context, instrument *model.LabInstrument) error {
	// 检查仪器编码是否已存在
	existing, err := s.repo.GetByCode(ctx, instrument.InstrumentCode)
	if err == nil && existing != nil {
		return fmt.Errorf("仪器编码 %s 已存在", instrument.InstrumentCode)
	}
	return s.repo.Create(ctx, instrument)
}

func (s *LabInstrumentService) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *LabInstrumentService) Delete(ctx context.Context, id uint64) error {
	// 删除关联的校准记录
	if err := s.calibrationRepo.DeleteByInstrumentID(ctx, id); err != nil {
		return fmt.Errorf("删除关联校准记录失败: %w", err)
	}
	return s.repo.Delete(ctx, id)
}

func (s *LabInstrumentService) GetCalibrations(ctx context.Context, instrumentID uint64, query *model.LabCalibrationQuery) ([]model.LabCalibration, int64, error) {
	if query.Page == 0 {
		query.Page = 1
	}
	if query.PageSize == 0 {
		query.PageSize = 20
	}
	return s.calibrationRepo.ListByInstrumentID(ctx, instrumentID, query)
}

func (s *LabInstrumentService) RecordCalibration(ctx context.Context, instrumentID uint64, req *model.RecordCalibrationRequest, tenantID uint64) error {
	// 获取仪器信息
	instrument, err := s.repo.GetByID(ctx, instrumentID)
	if err != nil {
		return fmt.Errorf("仪器不存在: %w", err)
	}

	// 创建校准记录
	calibration := &model.LabCalibration{
		TenantID:            tenantID,
		InstrumentID:        instrumentID,
		CalibrationDate:     req.CalibrationDate,
		CalibrationResult:   req.CalibrationResult,
		CalibratedBy:       req.CalibratedBy,
		CertificateNo:      req.CertificateNo,
		NextCalibrationDate: req.NextCalibrationDate,
		CalibrationItems:   req.CalibrationItems,
		AttachmentURL:      req.AttachmentURL,
		Remark:             req.Remark,
	}
	if err := s.calibrationRepo.Create(ctx, calibration); err != nil {
		return fmt.Errorf("创建校准记录失败: %w", err)
	}

	// 更新仪器的校准信息
	updates := map[string]interface{}{
		"last_calibration_date":  req.CalibrationDate,
		"next_calibration_date": req.NextCalibrationDate,
		"calibration_status":    model.CalibrationStatusCalibrated,
	}

	// 计算下次校准状态
	if !req.NextCalibrationDate.IsZero() {
		daysUntilDue := req.NextCalibrationDate.Sub(time.Now()).Hours() / 24
		if daysUntilDue < 0 {
			updates["calibration_status"] = model.CalibrationStatusOverdue
		} else if daysUntilDue <= 30 {
			updates["calibration_status"] = model.CalibrationStatusDue
		}
	}

	// 如果仪器状态不是维护中，保持ACTIVE状态
	if instrument.Status != model.InstrumentStatusMaintenance {
		updates["status"] = model.InstrumentStatusActive
	}

	return s.repo.Update(ctx, instrumentID, updates)
}