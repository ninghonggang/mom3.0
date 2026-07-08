package service

import (
	"context"
	"errors"
	"mom-server/internal/model"
	"mom-server/internal/pkg/status"
	"mom-server/internal/repository"
	"time"
)

// EquipmentDowntimeService 设备停机服务
type EquipmentDowntimeService struct {
	repo *repository.EquipmentDowntimeRepository
}

// NewEquipmentDowntimeService 创建设备停机服务
func NewEquipmentDowntimeService(repo *repository.EquipmentDowntimeRepository) *EquipmentDowntimeService {
	return &EquipmentDowntimeService{repo: repo}
}

// List 获取设备停机列表
func (s *EquipmentDowntimeService) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.EquipmentDowntime, int64, error) {
	return s.repo.List(ctx, tenantID, query)
}

// GetByID 根据ID获取设备停机记录
func (s *EquipmentDowntimeService) GetByID(ctx context.Context, id int64) (*model.EquipmentDowntime, error) {
	if id == 0 {
		return nil, errors.New("invalid id")
	}
	return s.repo.GetByID(ctx, id)
}

// Create 创建设备停机记录
func (s *EquipmentDowntimeService) Create(ctx context.Context, tenantID int64, req *model.EquipmentDowntimeCreateRequest, username string) (*model.EquipmentDowntime, error) {
	startTime, err := time.Parse(time.RFC3339, req.StartTime)
	if err != nil {
		startTime = time.Now()
	}

	downtime := &model.EquipmentDowntime{
		TenantID:        tenantID,
		EquipmentID:     req.EquipmentID,
		EquipmentCode:   req.EquipmentCode,
		EquipmentName:   req.EquipmentName,
		DowntimeType:    req.DowntimeType,
		DowntimeReason:  req.DowntimeReason,
		StartTime:       startTime,
		LostProduction:  req.LostProduction,
		WorkOrderID:     req.WorkOrderID,
		WorkOrderCode:   req.WorkOrderCode,
		ShiftID:        req.ShiftID,
		OperatorID:     req.OperatorID,
		OperatorName:   req.OperatorName,
		MaintainerID:   req.MaintainerID,
		MaintainerName: req.MaintainerName,
		Status:         "OPEN",
		Remark:         req.Remark,
		CreatedBy:      username,
		UpdatedBy:      username,
	}

	if req.EndTime != nil {
		endTime, err := time.Parse(time.RFC3339, *req.EndTime)
		if err == nil {
			downtime.EndTime = &endTime
			downtime.Duration = int(endTime.Sub(startTime).Minutes())
			downtime.SetStatus("CLOSED") // V2.1 双轨: 同时写 status + status_v2
		}
	}

	if err := s.repo.Create(ctx, downtime); err != nil {
		return nil, err
	}
	return downtime, nil
}

// Update 更新设备停机记录
func (s *EquipmentDowntimeService) Update(ctx context.Context, id int64, req *model.EquipmentDowntimeUpdateRequest) error {
	downtime, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if req.DowntimeType != "" {
		downtime.DowntimeType = req.DowntimeType
	}
	if req.DowntimeReason != "" {
		downtime.DowntimeReason = req.DowntimeReason
	}
	if req.StartTime != "" {
		if t, err := time.Parse(time.RFC3339, req.StartTime); err == nil {
			downtime.StartTime = t
		}
	}
	if req.EndTime != nil {
		if t, err := time.Parse(time.RFC3339, *req.EndTime); err == nil {
			downtime.EndTime = &t
			downtime.Duration = int(t.Sub(downtime.StartTime).Minutes())
		}
	}
	if req.Duration > 0 {
		downtime.Duration = req.Duration
	}
	if req.LostProduction > 0 {
		downtime.LostProduction = req.LostProduction
	}
	downtime.WorkOrderID = req.WorkOrderID
	downtime.WorkOrderCode = req.WorkOrderCode
	downtime.ShiftID = req.ShiftID
	downtime.OperatorID = req.OperatorID
	downtime.OperatorName = req.OperatorName
	downtime.MaintainerID = req.MaintainerID
	downtime.MaintainerName = req.MaintainerName
	if req.Status != "" {
		downtime.SetStatus(req.Status) // V2.1 双轨 + 字典码校验
	}
	if req.Remark != "" {
		downtime.Remark = req.Remark
	}

	return s.repo.Update(ctx, downtime)
}

// Delete 删除设备停机记录
func (s *EquipmentDowntimeService) Delete(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("invalid id")
	}
	return s.repo.Delete(ctx, id)
}

// StartDowntime 开始停机
func (s *EquipmentDowntimeService) StartDowntime(ctx context.Context, id int64) error {
	downtime, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if downtime.StatusCode() != string(status.DowntimeOpen) {
		return errors.New("only OPEN status can be started")
	}
	now := time.Now()
	downtime.StartTime = now
	downtime.SetStatus("INPROGRESS") // V2.1 双轨
	return s.repo.Update(ctx, downtime)
}

// EndDowntime 结束停机
func (s *EquipmentDowntimeService) EndDowntime(ctx context.Context, id int64) error {
	downtime, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if downtime.StatusCode() == string(status.DowntimeClosed) {
		return errors.New("already closed")
	}
	now := time.Now()
	downtime.EndTime = &now
	downtime.Duration = int(now.Sub(downtime.StartTime).Minutes())
	downtime.SetStatus("CLOSED") // V2.1 双轨
	return s.repo.Update(ctx, downtime)
}
