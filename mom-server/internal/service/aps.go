package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type MPSService struct {
	repo *repository.MPSRepository
}

func NewMPSService(repo *repository.MPSRepository) *MPSService {
	return &MPSService{repo: repo}
}

func (s *MPSService) List(ctx context.Context, tenantID int64) ([]model.MPS, int64, error) {
	return s.repo.List(ctx, tenantID)
}

func (s *MPSService) GetByID(ctx context.Context, id string) (*model.MPS, error) {
	var mpsID uint
	_, err := fmt.Sscanf(id, "%d", &mpsID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, mpsID)
}

func (s *MPSService) Create(ctx context.Context, mps *model.MPS) error {
	return s.repo.Create(ctx, mps)
}

func (s *MPSService) Update(ctx context.Context, id string, mps *model.MPS) error {
	var mpsID uint
	_, err := fmt.Sscanf(id, "%d", &mpsID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, mpsID, map[string]interface{}{
		"plan_month":    mps.PlanMonth,
		"material_id":   mps.MaterialID,
		"material_code": mps.MaterialCode,
		"material_name": mps.MaterialName,
		"quantity":     mps.Quantity,
		"status":       mps.Status,
	})
}

func (s *MPSService) Delete(ctx context.Context, id string) error {
	var mpsID uint
	_, err := fmt.Sscanf(id, "%d", &mpsID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, mpsID)
}

func (s *MPSService) Submit(ctx context.Context, id string) error {
	var mpsID uint
	_, err := fmt.Sscanf(id, "%d", &mpsID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, mpsID, map[string]interface{}{
		"status": 2,
	})
}

type MRPService struct {
	mrpRepo    *repository.MRPRepository
	invRepo    *repository.InventoryRepository
}

func NewMRPService(mrp *repository.MRPRepository, inv *repository.InventoryRepository) *MRPService {
	return &MRPService{mrpRepo: mrp, invRepo: inv}
}

func (s *MRPService) List(ctx context.Context, tenantID int64) ([]model.MRP, int64, error) {
	return s.mrpRepo.List(ctx, tenantID)
}

func (s *MRPService) Calculate(ctx context.Context, id string) error {
	var mrpID uint
	_, err := fmt.Sscanf(id, "%d", &mrpID)
	if err != nil {
		return err
	}
	// 简单MRP计算逻辑：净需求 = 需求数量 - 库存数量 - 已分配数量
	// 这里简化处理，实际需要BOM展开等复杂计算
	items, err := s.mrpRepo.GetItemsByMRPID(ctx, int64(mrpID))
	if err != nil {
		return err
	}
	for i := range items {
		netQty := items[i].Quantity - items[i].StockQty - items[i].AllocatedQty
		if netQty < 0 {
			netQty = 0
		}
		items[i].NetQty = netQty
		// 更新净需求到数据库
		s.mrpRepo.Update(ctx, uint(items[i].ID), map[string]interface{}{
			"net_qty": netQty,
		})
	}
	return s.mrpRepo.Update(ctx, mrpID, map[string]interface{}{
		"status": 3, // 已完成
	})
}

type ScheduleService struct {
	scheduleRepo *repository.ScheduleRepository
	orderRepo   *repository.ProductionOrderRepository
}

func NewScheduleService(schedule *repository.ScheduleRepository, order *repository.ProductionOrderRepository) *ScheduleService {
	return &ScheduleService{scheduleRepo: schedule, orderRepo: order}
}

func (s *ScheduleService) List(ctx context.Context, tenantID int64) ([]model.SchedulePlan, int64, error) {
	return s.scheduleRepo.List(ctx, tenantID)
}

func (s *ScheduleService) Create(ctx context.Context, plan *model.SchedulePlan) error {
	return s.scheduleRepo.Create(ctx, plan)
}

func (s *ScheduleService) Execute(ctx context.Context, id string) error {
	var planID uint
	_, err := fmt.Sscanf(id, "%d", &planID)
	if err != nil {
		return err
	}
	plan, err := s.scheduleRepo.GetByID(ctx, planID)
	if err != nil {
		return err
	}
	// 删除旧的排程结果
	results, _ := s.scheduleRepo.GetResultsByPlanID(ctx, int64(planID))
	for _, r := range results {
		s.scheduleRepo.Delete(ctx, uint(r.ID))
	}
	// 简单排程：按工序顺序排，标准工时估算
	// 实际需要遗传算法/粒子群等优化
	orders, _, err := s.orderRepo.List(ctx, 0)
	if err != nil {
		return err
	}
	seq := 1
	currentTime := time.Now()
	if plan.StartDate != nil {
		currentTime = *plan.StartDate
	}
	for _, order := range orders {
		// 简单分配：按顺序分配产线和工位
		// 实际需要根据资源能力、瓶颈等优化
		lineID := int64(seq % 3) // 模拟3条产线轮换
		lineName := fmt.Sprintf("产线%d", lineID+1)
		stationID := int64(seq % 5)
		stationName := fmt.Sprintf("工位%d", stationID+1)
		// 估算工时：默认8小时
		duration := 8 * time.Hour
		planStartTime := currentTime.Add(time.Duration((seq-1)*8) * time.Hour)
		planEndTime := planStartTime.Add(duration)
		result := &model.ScheduleResult{
			PlanID:         int64(planID),
			OrderID:        order.ID,
			OrderNo:        order.OrderNo,
			Sequence:       seq,
			LineID:         lineID,
			LineName:       &lineName,
			StationID:      &stationID,
			StationName:    &stationName,
			PlanStartTime:  &planStartTime,
			PlanEndTime:    &planEndTime,
		}
		s.scheduleRepo.CreateResult(ctx, result)
		seq++
	}
	return s.scheduleRepo.Update(ctx, planID, map[string]interface{}{
		"status": 3, // 已完成
	})
}

func (s *ScheduleService) GetResults(ctx context.Context, id string) ([]model.ScheduleResult, error) {
	var planID uint
	_, err := fmt.Sscanf(id, "%d", &planID)
	if err != nil {
		return nil, err
	}
	return s.scheduleRepo.GetResultsByPlanID(ctx, int64(planID))
}

func (s *ScheduleService) Delete(ctx context.Context, id string) error {
	var planID uint
	_, err := fmt.Sscanf(id, "%d", &planID)
	if err != nil {
		return err
	}
	return s.scheduleRepo.Delete(ctx, planID)
}

// DragUpdate 更新排程结果（拖拽后保存）
func (s *ScheduleService) DragUpdate(ctx context.Context, resultID uint, lineID int64, stationID int64, planStartTime, planEndTime time.Time) error {
	updates := map[string]interface{}{
		"line_id":         lineID,
		"station_id":      stationID,
		"plan_start_time": planStartTime,
		"plan_end_time":   planEndTime,
	}
	return s.scheduleRepo.UpdateResult(ctx, resultID, updates)
}

// GetResultByID 获取单个排程结果
func (s *ScheduleService) GetResultByID(ctx context.Context, id string) (*model.ScheduleResult, error) {
	var resultID uint
	_, err := fmt.Sscanf(id, "%d", &resultID)
	if err != nil {
		return nil, err
	}
	return s.scheduleRepo.GetResultByID(ctx, resultID)
}
