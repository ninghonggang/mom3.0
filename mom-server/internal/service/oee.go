package service

import (
	"context"
	"errors"
	"fmt"
	"math"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type OEEService struct {
	repo *repository.OEERepository
	eventRepo *repository.OEEEventRepository
}

func NewOEEService(repo *repository.OEERepository, eventRepo *repository.OEEEventRepository) *OEEService {
	return &OEEService{repo: repo, eventRepo: eventRepo}
}

// OEE计算请求
type OECalculateReq struct {
	EquipmentID   int64  `json:"equipment_id"`
	EquipmentCode string `json:"equipment_code"`
	EquipmentName string `json:"equipment_name"`
	WorkshopID    int64  `json:"workshop_id"`
	RecordDate    string `json:"record_date"`
	PlanTime      int    `json:"plan_time"`       // 计划时间(分钟)
	RunTime       int    `json:"run_time"`        // 运行时间(分钟)
	DownTime      int    `json:"down_time"`       // 停机时间(分钟)
	IdleTime      int    `json:"idle_time"`       // 空闲时间(分钟)
	PlanStopTime  int    `json:"plan_stop_time"`  // 计划停机时间(分钟)
	OutputQty     int    `json:"output_qty"`      // 产出数量
	QualifiedQty  int    `json:"qualified_qty"`   // 合格数量
	TheoreticalOutput int `json:"theoretical_output"` // 理论产量
}

func (s *OEEService) List(ctx context.Context, params map[string]interface{}) ([]model.OEE, int64, error) {
	return s.repo.List(ctx, 1, params)
}

func (s *OEEService) GetByID(ctx context.Context, id int64) (*model.OEE, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *OEEService) Calculate(ctx context.Context, req *OECalculateReq) (*model.OEE, error) {
	// 参数校验
	if req.EquipmentID == 0 {
		return nil, errors.New("设备ID不能为空")
	}
	if req.RecordDate == "" {
		return nil, errors.New("记录日期不能为空")
	}
	if req.PlanTime <= 0 {
		return nil, errors.New("计划时间必须大于0")
	}

	// 计算OEE
	oee := &model.OEE{
		TenantID:       1,
		EquipmentID:    req.EquipmentID,
		EquipmentCode:  req.EquipmentCode,
		EquipmentName:  req.EquipmentName,
		WorkshopID:     req.WorkshopID,
		RecordDate:     req.RecordDate,
		PlanTime:       req.PlanTime,
		RunTime:        req.RunTime,
		DownTime:       req.DownTime,
		IdleTime:       req.IdleTime,
		PlanStopTime:   req.PlanStopTime,
		OutputQty:      req.OutputQty,
		QualifiedQty:   req.QualifiedQty,
	}

	// 计算可用率 = 运行时间 / (运行时间 + 停机时间) * 100
	runAndDown := oee.RunTime + oee.DownTime
	if runAndDown > 0 {
		oee.Availability = math.Round(float64(oee.RunTime)/float64(runAndDown)*10000) / 100
	}

	// 计算性能率 = 实际产量 / 理论产量 * 100
	if req.TheoreticalOutput > 0 {
		oee.Performance = math.Round(float64(req.OutputQty)/float64(req.TheoreticalOutput)*10000) / 100
	} else if oee.RunTime > 0 {
		// 如果没有理论产量，使用计划时间作为参考
		oee.Performance = math.Round(float64(oee.RunTime)/float64(oee.PlanTime)*10000) / 100
		if oee.Performance > 100 {
			oee.Performance = 100
		}
	}

	// 计算质量率 = 合格数量 / 总产量 * 100
	if oee.OutputQty > 0 {
		oee.Quality = math.Round(float64(oee.QualifiedQty)/float64(oee.OutputQty)*10000) / 100
	}

	// 计算OEE = 可用率 × 性能率 × 质量率
	oee.OEE = math.Round(oee.Availability*oee.Performance*oee.Quality) / 100

	// 检查是否已存在记录
	existing, err := s.repo.GetByEquipmentAndDate(ctx, req.EquipmentID, req.RecordDate)
	if err == nil && existing != nil {
		// 更新已有记录
		oee.ID = existing.ID
		err = s.repo.Update(ctx, existing.ID, map[string]interface{}{
			"plan_time":      oee.PlanTime,
			"run_time":       oee.RunTime,
			"down_time":      oee.DownTime,
			"idle_time":      oee.IdleTime,
			"plan_stop_time": oee.PlanStopTime,
			"output_qty":     oee.OutputQty,
			"qualified_qty":  oee.QualifiedQty,
			"availability":   oee.Availability,
			"performance":    oee.Performance,
			"quality":        oee.Quality,
			"oee":           oee.OEE,
		})
		if err != nil {
			return nil, fmt.Errorf("更新OEE记录失败: %w", err)
		}
	} else {
		// 创建新记录
		err = s.repo.Create(ctx, oee)
		if err != nil {
			return nil, fmt.Errorf("创建OEE记录失败: %w", err)
		}
	}

	return oee, nil
}

func (s *OEEService) GetChartData(ctx context.Context, params map[string]interface{}) ([]model.OEE, error) {
	return s.repo.GetChartData(ctx, 1, params)
}

func (s *OEEService) Delete(ctx context.Context, id int64) error {
	// 删除关联事件记录
	_ = s.eventRepo.DeleteByOEEID(ctx, id)
	return s.repo.Delete(ctx, id)
}
