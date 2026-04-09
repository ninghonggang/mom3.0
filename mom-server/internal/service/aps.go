package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
	"mom-server/internal/scheduler"
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
	bomRepo    *repository.BOMRepository
	bomItemRepo *repository.BOMItemRepository
	mpsRepo    *repository.MPSRepository
}

func NewMRPService(mrp *repository.MRPRepository, inv *repository.InventoryRepository, bomRepo *repository.BOMRepository, bomItemRepo *repository.BOMItemRepository, mpsRepo *repository.MPSRepository) *MRPService {
	return &MRPService{
		mrpRepo:    mrp,
		invRepo:    inv,
		bomRepo:    bomRepo,
		bomItemRepo: bomItemRepo,
		mpsRepo:    mpsRepo,
	}
}

func (s *MRPService) List(ctx context.Context, tenantID int64) ([]model.MRP, int64, error) {
	return s.mrpRepo.List(ctx, tenantID)
}

func (s *MRPService) GetByID(ctx context.Context, id string) (*model.MRP, error) {
	var mrpID uint
	_, err := fmt.Sscanf(id, "%d", &mrpID)
	if err != nil {
		return nil, err
	}
	return s.mrpRepo.GetByID(ctx, mrpID)
}

func (s *MRPService) Create(ctx context.Context, mrp *model.MRP) error {
	return s.mrpRepo.Create(ctx, mrp)
}

func (s *MRPService) Update(ctx context.Context, id string, updates map[string]interface{}) error {
	var mrpID uint
	_, err := fmt.Sscanf(id, "%d", &mrpID)
	if err != nil {
		return err
	}
	return s.mrpRepo.Update(ctx, mrpID, updates)
}

func (s *MRPService) Delete(ctx context.Context, id string) error {
	var mrpID uint
	_, err := fmt.Sscanf(id, "%d", &mrpID)
	if err != nil {
		return err
	}
	// 删除关联的MRP明细
	if err := s.mrpRepo.DeleteItemsByMRPID(ctx, int64(mrpID)); err != nil {
		return err
	}
	return s.mrpRepo.Delete(ctx, mrpID)
}

// CalculateNetDemand 计算净需求
// 净需求 = 毛需求 - 可用库存 - 已分配数量
func (s *MRPService) CalculateNetDemand(ctx context.Context, materialID int64, grossDemand float64) (float64, float64, error) {
	// 获取物料库存信息
	inventories, _, err := s.invRepo.ListByMaterialID(ctx, materialID)
	if err != nil {
		return 0, 0, fmt.Errorf("获取库存失败: %w", err)
	}

	var availableQty float64
	for _, inv := range inventories {
		availableQty += inv.AvailableQty
	}

	netDemand := grossDemand - availableQty
	if netDemand < 0 {
		netDemand = 0
	}

	return netDemand, availableQty, nil
}

// AnalyzeShortage 分析缺料情况
// 返回缺料物料列表
func (s *MRPService) AnalyzeShortage(ctx context.Context, mrpID int64) ([]model.MRPItem, error) {
	items, err := s.mrpRepo.GetItemsByMRPID(ctx, mrpID)
	if err != nil {
		return nil, err
	}

	var shortageItems []model.MRPItem
	for _, item := range items {
		// 计算净需求
		netQty := item.Quantity - item.StockQty - item.AllocatedQty
		if netQty > 0 {
			shortageItems = append(shortageItems, model.MRPItem{
				MaterialID:   item.MaterialID,
				MaterialCode: item.MaterialCode,
				MaterialName: item.MaterialName,
				Quantity:     item.Quantity,
				StockQty:    item.StockQty,
				AllocatedQty: item.AllocatedQty,
				NetQty:      netQty,
				SourceType:  "SHORTAGE",
			})
		}
	}

	return shortageItems, nil
}

// GeneratePurchaseSuggestion 生成采购建议
// 基于净需求生成采购建议
func (s *MRPService) GeneratePurchaseSuggestion(ctx context.Context, mrpID int64) ([]PurchaseSuggestion, error) {
	items, err := s.mrpRepo.GetItemsByMRPID(ctx, mrpID)
	if err != nil {
		return nil, err
	}

	var suggestions []PurchaseSuggestion
	for _, item := range items {
		netQty := item.NetQty
		if netQty > 0 {
			suggestion := PurchaseSuggestion{
				MaterialID:   item.MaterialID,
				MaterialCode: item.MaterialCode,
				MaterialName: item.MaterialName,
				DemandQty:    netQty,
				UrgentLevel:  "NORMAL",
			}
			// 根据缺料程度判断紧急级别
			if item.StockQty <= 0 {
				suggestion.UrgentLevel = "URGENT"
			} else if item.StockQty < item.Quantity*0.2 {
				suggestion.UrgentLevel = "HIGH"
			}
			suggestions = append(suggestions, suggestion)
		}
	}

	return suggestions, nil
}

// PurchaseSuggestion 采购建议结构
type PurchaseSuggestion struct {
	MaterialID    int64   `json:"material_id"`
	MaterialCode  string  `json:"material_code"`
	MaterialName  string  `json:"material_name"`
	DemandQty     float64 `json:"demand_qty"`
	UrgentLevel   string  `json:"urgent_level"` // URGENT/HIGH/NORMAL
	SuggestedQty  float64 `json:"suggested_qty"`
}

// RunMRP 执行完整的MRP计算（基于MPS和BOM展开）
func (s *MRPService) RunMRP(ctx context.Context, mrpID int64, planMonth string) error {
	// 1. 更新MRP状态为计算中
	if err := s.mrpRepo.Update(ctx, uint(mrpID), map[string]interface{}{"status": 2}); err != nil {
		return fmt.Errorf("更新MRP状态失败: %w", err)
	}

	// 2. 获取已提交的MPS计划
	mpsList, _, err := s.mpsRepo.ListByStatus(ctx, 2) // 2=已提交
	if err != nil {
		return fmt.Errorf("获取MPS计划失败: %w", err)
	}

	// 3. 过滤指定月份的计划
	var relevantMPS []model.MPS
	for _, mps := range mpsList {
		if planMonth != "" && mps.PlanMonth == planMonth {
			relevantMPS = append(relevantMPS, mps)
		} else if planMonth == "" {
			relevantMPS = append(relevantMPS, mps)
		}
	}

	// 4. 删除旧的MRP明细
	if err := s.mrpRepo.DeleteItemsByMRPID(ctx, mrpID); err != nil {
		return fmt.Errorf("删除旧MRP明细失败: %w", err)
	}

	// 5. BOM展开计算
	for _, mps := range relevantMPS {
		// 获取物料的BOM
		bom, err := s.bomRepo.GetByMaterialID(ctx, mps.MaterialID)
		if err != nil {
			// 没有BOM，直接添加为需求
			netQty, availableQty, err := s.CalculateNetDemand(ctx, mps.MaterialID, mps.Quantity)
			if err != nil {
				return fmt.Errorf("计算净需求失败: %w", err)
			}
			item := &model.MRPItem{
				MRPID:       mrpID,
				MaterialID:  mps.MaterialID,
				MaterialCode: mps.MaterialCode,
				MaterialName: mps.MaterialName,
				Quantity:    mps.Quantity,
				StockQty:    availableQty,
				NetQty:      netQty,
				SourceType:  "MPS",
				SourceNo:    &mps.MPSNo,
			}
			if err := s.mrpRepo.CreateItem(ctx, item); err != nil {
				return fmt.Errorf("创建MRP明细失败: %w", err)
			}
			continue
		}

		// BOM展开：计算下层物料需求
		if err := s.bomExplosion(ctx, mrpID, bom.ID, mps.Quantity, "MPS", mps.MPSNo); err != nil {
			return fmt.Errorf("BOM展开失败: %w", err)
		}
	}

	// 6. 更新MRP状态为已完成
	if err := s.mrpRepo.Update(ctx, uint(mrpID), map[string]interface{}{"status": 3}); err != nil {
		return fmt.Errorf("更新MRP状态失败: %w", err)
	}

	return nil
}

// bomExplosion BOM展开
func (s *MRPService) bomExplosion(ctx context.Context, mrpID int64, bomID int64, parentQty float64, sourceType string, sourceNo string) error {
	bomItems, err := s.bomItemRepo.ListByBOMID(ctx, uint(bomID))
	if err != nil {
		return err
	}

	for _, bomItem := range bomItems {
		// 计算物料需求：父项数量 × BOM用量 × (1 + 损耗率)
		grossDemand := parentQty * bomItem.Quantity * (1 + bomItem.ScrapRate)

		// 获取库存信息
		netQty, availableQty, err := s.CalculateNetDemand(ctx, bomItem.MaterialID, grossDemand)
		if err != nil {
			return fmt.Errorf("计算净需求失败: %w", err)
		}

		item := &model.MRPItem{
			MRPID:       mrpID,
			MaterialID:  bomItem.MaterialID,
			MaterialCode: bomItem.MaterialCode,
			MaterialName: bomItem.MaterialName,
			Quantity:    grossDemand,
			StockQty:    availableQty,
			AllocatedQty: 0,
			NetQty:      netQty,
			SourceType:  sourceType,
			SourceNo:    &sourceNo,
		}
		if err := s.mrpRepo.CreateItem(ctx, item); err != nil {
			return err
		}

		// 检查是否有子BOM，递归展开
		subBOM, err := s.bomRepo.GetByMaterialID(ctx, bomItem.MaterialID)
		if err == nil && subBOM.ID > 0 {
			if err := s.bomExplosion(ctx, mrpID, subBOM.ID, netQty, "BOM", sourceNo); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetMrpResults 获取MRP计算结果
func (s *MRPService) GetMrpResults(ctx context.Context, mrpID int64) ([]model.MRPItem, error) {
	return s.mrpRepo.GetItemsByMRPID(ctx, mrpID)
}

// Calculate 执行MRP计算（简单版本，基于已有明细）
func (s *MRPService) Calculate(ctx context.Context, id string) error {
	var mrpID uint
	_, err := fmt.Sscanf(id, "%d", &mrpID)
	if err != nil {
		return err
	}

	// 更新状态为计算中
	if err := s.mrpRepo.Update(ctx, mrpID, map[string]interface{}{"status": 2}); err != nil {
		return err
	}

	// 获取MRP明细
	items, err := s.mrpRepo.GetItemsByMRPID(ctx, int64(mrpID))
	if err != nil {
		return err
	}

	// 逐项计算净需求
	for i := range items {
		netQty := items[i].Quantity - items[i].StockQty - items[i].AllocatedQty
		if netQty < 0 {
			netQty = 0
		}
		items[i].NetQty = netQty

		// 更新净需求到数据库
		if err := s.mrpRepo.UpdateItem(ctx, uint(items[i].ID), map[string]interface{}{
			"net_qty": netQty,
		}); err != nil {
			return err
		}
	}

	// 更新MRP状态为已完成
	return s.mrpRepo.Update(ctx, mrpID, map[string]interface{}{
		"status": 3,
	})
}

type ScheduleService struct {
	scheduleRepo *repository.ScheduleRepository
	orderRepo    *repository.ProductionOrderRepository
	lineRepo     *repository.ProductionLineRepository
}

func NewScheduleService(schedule *repository.ScheduleRepository, order *repository.ProductionOrderRepository, line *repository.ProductionLineRepository) *ScheduleService {
	return &ScheduleService{scheduleRepo: schedule, orderRepo: order, lineRepo: line}
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

// SchedulingRequest 排程请求
type SchedulingRequest struct {
	PlanID    int64  `json:"plan_id"`
	Rule      string `json:"rule"`       // FIFO/EDD/SPT/LPT
	Direction string `json:"direction"`  // FORWARD/BACKWARD
}

// ConstrainedScheduleRequest 带约束排程请求
type ConstrainedScheduleRequest struct {
	PlanID        int64              `json:"plan_id"`
	AlgorithmType string             `json:"algorithm_type"`
	Direction     string             `json:"direction"`
	WorkshopID    int64              `json:"workshop_id"`
	Constraints   ScheduleConstraints `json:"constraints"`
}

// ScheduleConstraints 排程约束
type ScheduleConstraints struct {
	RespectJIT       bool    `json:"respect_jit"`
	MaxChangeoverPct float64 `json:"max_changeover_pct"`
	MinUtilization   float64 `json:"min_utilization"`
	AllowOvertime    bool    `json:"allow_overtime"`
	FamilyGrouping   bool    `json:"family_grouping"`
}

// GanttData 甘特图数据
type GanttData struct {
	Tasks     []GanttTask      `json:"tasks"`
	Resources []GanttResource   `json:"resources"`
}

// GanttTask 甘特图任务
type GanttTask struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Start     string   `json:"start"`
	End       string   `json:"end"`
	Progress  int      `json:"progress"`
	Status    string   `json:"status"`
	Resources []string `json:"resources"`
	LineID    int64    `json:"line_id"`
	LineName  string   `json:"line_name"`
	OrderID   int64    `json:"order_id"`
}

// GanttResource 甘特图资源
type GanttResource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ExecuteSchedulingWithRule 使用指定规则执行排程
func (s *ScheduleService) ExecuteSchedulingWithRule(ctx context.Context, req SchedulingRequest) error {
	plan, err := s.scheduleRepo.GetByID(ctx, uint(req.PlanID))
	if err != nil {
		return fmt.Errorf("获取排程计划失败: %w", err)
	}

	// 删除旧的排程结果
	results, _ := s.scheduleRepo.GetResultsByPlanID(ctx, req.PlanID)
	for _, r := range results {
		s.scheduleRepo.Delete(ctx, uint(r.ID))
	}

	// 获取待排程的生产工单
	orders, _, err := s.orderRepo.List(ctx, 0)
	if err != nil {
		return fmt.Errorf("获取工单失败: %w", err)
	}

	// 转换为排程任务
	tasks := s.convertToTasks(orders)

	// 获取产线资源
	lines, _, err := s.lineRepo.List(ctx, 0)
	if err != nil {
		return fmt.Errorf("获取产线失败: %w", err)
	}
	resources := s.convertToResources(lines)

	// 根据规则创建排程器
	var rule scheduler.SchedulingRule
	var direction scheduler.SchedulingDirection

	switch req.Rule {
	case "EDD":
		rule = scheduler.RuleEDD
	case "SPT":
		rule = scheduler.RuleSPT
	case "LPT":
		rule = scheduler.RuleLPT
	default:
		rule = scheduler.RuleFIFO
	}

	if req.Direction == "BACKWARD" {
		direction = scheduler.DirectionBackward
	} else {
		direction = scheduler.DirectionForward
	}

	var schedulingService *scheduler.SchedulingService
	if direction == scheduler.DirectionForward {
		schedulingService = scheduler.NewSchedulingService(scheduler.NewForwardScheduler(rule))
	} else {
		schedulingService = scheduler.NewSchedulingService(scheduler.NewBackwardScheduler(rule))
	}

	result, err := schedulingService.ExecuteScheduling(ctx, tasks, resources, direction)
	if err != nil {
		return fmt.Errorf("排程失败: %w", err)
	}

	// 保存排程结果
	scheduleResults := scheduler.ConvertToScheduleResults(req.PlanID, result)
	for _, sr := range scheduleResults {
		if err := s.scheduleRepo.CreateResult(ctx, sr); err != nil {
			return fmt.Errorf("保存排程结果失败: %w", err)
		}
	}

	// 更新计划状态
	if err := s.scheduleRepo.Update(ctx, uint(req.PlanID), map[string]interface{}{
		"status": 3,
	}); err != nil {
		return fmt.Errorf("更新计划状态失败: %w", err)
	}

	_ = plan // avoid unused variable
	return nil
}

// OptimizeSchedule 优化排程结果
func (s *ScheduleService) OptimizeSchedule(ctx context.Context, planID int64) error {
	results, err := s.scheduleRepo.GetResultsByPlanID(ctx, planID)
	if err != nil {
		return err
	}

	// 获取产线资源
	lines, _, err := s.lineRepo.List(ctx, 0)
	if err != nil {
		return err
	}
	resources := s.convertToResources(lines)

	// 转换为排程任务格式
	scheduledTasks := make([]scheduler.ScheduledTask, len(results))
	for i, r := range results {
		loadHours := float64(8) // 默认8小时
		if r.PlanStartTime != nil && r.PlanEndTime != nil {
			loadHours = r.PlanEndTime.Sub(*r.PlanStartTime).Hours()
		}
		scheduledTasks[i] = scheduler.ScheduledTask{
			TaskID:       int64(r.ID),
			OrderID:      r.OrderID,
			OrderNo:      r.OrderNo,
			ResourceID:   r.LineID,
			ResourceName: *r.LineName,
			Sequence:     r.Sequence,
			LoadHours:    loadHours,
		}
		if r.PlanStartTime != nil {
			scheduledTasks[i].PlanStartTime = *r.PlanStartTime
		}
		if r.PlanEndTime != nil {
			scheduledTasks[i].PlanEndTime = *r.PlanEndTime
		}
	}

	// 优化
	optimized := scheduler.OptimizeSchedule(&scheduler.SchedulingResult{
		Tasks: scheduledTasks,
	}, resources)

	// 更新结果
	for _, t := range optimized.Tasks {
		if err := s.scheduleRepo.UpdateResult(ctx, uint(t.TaskID), map[string]interface{}{
			"line_id": t.ResourceID,
		}); err != nil {
			return err
		}
	}

	return nil
}

// CalculateLoad 计算产能负载
func (s *ScheduleService) CalculateLoad(ctx context.Context, planID int64) (map[int64]float64, error) {
	results, err := s.scheduleRepo.GetResultsByPlanID(ctx, planID)
	if err != nil {
		return nil, err
	}

	loads := make(map[int64]float64)
	for _, r := range results {
		if r.PlanStartTime != nil && r.PlanEndTime != nil {
			hours := r.PlanEndTime.Sub(*r.PlanStartTime).Hours()
			loads[r.LineID] += hours
		}
	}

	return loads, nil
}

// GetGanttData 获取甘特图数据
func (s *ScheduleService) GetGanttData(ctx context.Context, planID int64) (*GanttData, error) {
	results, err := s.scheduleRepo.GetResultsByPlanID(ctx, planID)
	if err != nil {
		return nil, err
	}

	// 获取产线信息用于去重
	lines, _, err := s.lineRepo.List(ctx, 0)
	if err != nil {
		return nil, err
	}

	ganttData := &GanttData{
		Tasks:     make([]GanttTask, 0, len(results)),
		Resources: make([]GanttResource, 0, len(lines)),
	}

	// 添加资源
	lineMap := make(map[int64]bool)
	for _, line := range lines {
		if !lineMap[line.ID] {
			lineMap[line.ID] = true
			ganttData.Resources = append(ganttData.Resources, GanttResource{
				ID:   fmt.Sprintf("line-%d", line.ID),
				Name: line.LineName,
			})
		}
	}

	// 添加任务
	for _, r := range results {
		status := "WAITING"
		progress := 0
		if r.Status == 2 {
			status = "RUNNING"
			progress = 50
		} else if r.Status == 3 {
			status = "COMPLETED"
			progress = 100
		}

		startStr := ""
		endStr := ""
		if r.PlanStartTime != nil {
			startStr = r.PlanStartTime.Format("2006-01-02 15:04")
		}
		if r.PlanEndTime != nil {
			endStr = r.PlanEndTime.Format("2006-01-02 15:04")
		}

		lineName := ""
		if r.LineName != nil {
			lineName = *r.LineName
		}

		task := GanttTask{
			ID:        fmt.Sprintf("task-%d", r.ID),
			Name:      r.OrderNo,
			Start:     startStr,
			End:       endStr,
			Progress:  progress,
			Status:    status,
			Resources: []string{lineName},
			LineID:    r.LineID,
			LineName:  lineName,
			OrderID:   r.OrderID,
		}
		ganttData.Tasks = append(ganttData.Tasks, task)
	}

	return ganttData, nil
}

// UpdateTaskTime 更新任务时间
func (s *ScheduleService) UpdateTaskTime(ctx context.Context, id uint, startTime, endTime time.Time) error {
	return s.scheduleRepo.UpdateResult(ctx, id, map[string]interface{}{
		"plan_start_time": startTime,
		"plan_end_time":   endTime,
	})
}

// convertToTasks 转换工单为排程任务
func (s *ScheduleService) convertToTasks(orders []model.ProductionOrder) []scheduler.TaskInfo {
	tasks := make([]scheduler.TaskInfo, 0, len(orders))
	for _, order := range orders {
		if order.Status == 1 || order.Status == 2 { // 只排待生产和生产中的工单
			// 默认标准工时：数量 * 0.5小时
			stdHours := order.Quantity * 0.5
			dueDate := time.Now().Add(72 * time.Hour) // 默认3天后
			if order.PlanEndDate != nil {
				dueDate = *order.PlanEndDate
			}
			createTime := order.CreatedAt

			task := scheduler.TaskInfo{
				OrderID:       order.ID,
				OrderNo:       order.OrderNo,
				Priority:      order.Priority,
				Quantity:      order.Quantity,
				StandardHours: stdHours,
				DueDate:      dueDate,
				CreateTime:   createTime,
			}
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// convertToResources 转换产线为排程资源
func (s *ScheduleService) convertToResources(lines []model.ProductionLine) []scheduler.ResourceInfo {
	resources := make([]scheduler.ResourceInfo, 0, len(lines))
	for _, line := range lines {
		resource := scheduler.ResourceInfo{
			ResourceID:   line.ID,
			ResourceName: line.LineName,
			Capacity:     1.0, // 默认产能
			Efficiency:   100, // 默认效率
		}
		resources = append(resources, resource)
	}
	return resources
}

// ExecuteConstrainedScheduling 执行带约束的排程
func (s *ScheduleService) ExecuteConstrainedScheduling(ctx context.Context, req ConstrainedScheduleRequest) error {
	plan, err := s.scheduleRepo.GetByID(ctx, uint(req.PlanID))
	if err != nil {
		return fmt.Errorf("获取排程计划失败: %w", err)
	}

	// 删除旧的排程结果
	results, _ := s.scheduleRepo.GetResultsByPlanID(ctx, req.PlanID)
	for _, r := range results {
		s.scheduleRepo.Delete(ctx, uint(r.ID))
	}

	// 获取待排程的生产工单
	orders, _, err := s.orderRepo.List(ctx, 0)
	if err != nil {
		return fmt.Errorf("获取工单失败: %w", err)
	}

	// 转换为带约束的排程任务
	scheduleOrders := s.convertToScheduleOrders(orders)

	// 获取产线资源
	lines, _, err := s.lineRepo.List(ctx, 0)
	if err != nil {
		return fmt.Errorf("获取产线失败: %w", err)
	}
	workCenters := s.convertToWorkCenters(lines)

	// 创建换型感知排程器
	changeoverScheduler := scheduler.NewChangeoverAwareScheduler(nil, nil)

	// 创建约束排程器
	constrainedScheduler := scheduler.NewConstrainedScheduler(nil, changeoverScheduler)

	// 转换为排程请求
	scheduleReq := &scheduler.ScheduleRequest{
		PlanID:        req.PlanID,
		WorkshopID:   req.WorkshopID,
		AlgorithmType: req.AlgorithmType,
		Direction:     scheduler.SchedulingDirection(req.Direction),
		Orders:       scheduleOrders,
		WorkCenters:  workCenters,
		Constraints: scheduler.ScheduleConstraints{
			RespectJIT:       req.Constraints.RespectJIT,
			MaxChangeoverPct: req.Constraints.MaxChangeoverPct,
			MinUtilization:   req.Constraints.MinUtilization,
			AllowOvertime:    req.Constraints.AllowOvertime,
			FamilyGrouping:   req.Constraints.FamilyGrouping,
		},
	}

	result, err := constrainedScheduler.Schedule(ctx, scheduleReq)
	if err != nil {
		return fmt.Errorf("排程失败: %w", err)
	}

	// 保存排程结果
	for _, sr := range result.Tasks {
		lineName := ""
		for _, wc := range workCenters {
			if wc.LineID == sr.ResourceID {
				lineName = wc.LineName
				break
			}
		}
		scheduleResult := &model.ScheduleResult{
			PlanID:        req.PlanID,
			OrderID:       sr.OrderID,
			OrderNo:       sr.OrderNo,
			Sequence:      sr.Sequence,
			LineID:        sr.ResourceID,
			LineName:      &lineName,
			PlanStartTime: &sr.PlanStartTime,
			PlanEndTime:   &sr.PlanEndTime,
		}
		if err := s.scheduleRepo.CreateResult(ctx, scheduleResult); err != nil {
			return fmt.Errorf("保存排程结果失败: %w", err)
		}
	}

	// 更新计划状态
	if err := s.scheduleRepo.Update(ctx, uint(req.PlanID), map[string]interface{}{
		"status": 3,
	}); err != nil {
		return fmt.Errorf("更新计划状态失败: %w", err)
	}

	_ = plan // avoid unused
	return nil
}

// GetOptimizationSuggestions 获取优化建议
func (s *ScheduleService) GetOptimizationSuggestions(ctx context.Context, planID int64) ([]scheduler.OptimizationSuggestion, error) {
	results, err := s.scheduleRepo.GetResultsByPlanID(ctx, planID)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return []scheduler.OptimizationSuggestion{}, nil
	}

	suggestions := make([]scheduler.OptimizationSuggestion, 0)

	// 计算利用率
	var totalHours float64
	var maxEndTime, minStartTime time.Time

	for _, r := range results {
		if r.PlanStartTime != nil && r.PlanEndTime != nil {
			hours := r.PlanEndTime.Sub(*r.PlanStartTime).Hours()
			totalHours += hours
			if maxEndTime.IsZero() || r.PlanEndTime.After(maxEndTime) {
				maxEndTime = *r.PlanEndTime
			}
			if minStartTime.IsZero() || r.PlanStartTime.Before(minStartTime) {
				minStartTime = *r.PlanStartTime
			}
		}
	}

	timeSpan := maxEndTime.Sub(minStartTime).Hours()
	utilization := 0.0
	if timeSpan > 0 && len(results) > 0 {
		// 简化计算：假设有3条产线
		utilization = (totalHours / timeSpan) / 3 * 100
	}

	if utilization < 70 {
		suggestions = append(suggestions, scheduler.OptimizationSuggestion{
			Type:    "CAPACITY",
			Level:   "WARNING",
			Message: fmt.Sprintf("资源利用率 %.1f%% 偏低，建议优化排程顺序", utilization),
			Impact:  "产能未充分利用",
		})
	}

	// 检查交付逾期
	for _, r := range results {
		order, _ := s.orderRepo.GetByID(ctx, uint(r.OrderID))
		if order != nil && order.PlanEndDate != nil && r.PlanEndTime != nil {
			if r.PlanEndTime.After(*order.PlanEndDate) {
				suggestions = append(suggestions, scheduler.OptimizationSuggestion{
					Type:    "DELIVERY",
					Level:   "CRITICAL",
					Message: fmt.Sprintf("订单 %s 计划结束时间晚于交期", r.OrderNo),
					Impact:  "将导致逾期交付",
				})
			}
		}
	}

	return suggestions, nil
}

// convertToScheduleOrders 转换工单为排程订单
func (s *ScheduleService) convertToScheduleOrders(orders []model.ProductionOrder) []scheduler.ScheduleOrder {
	scheduleOrders := make([]scheduler.ScheduleOrder, 0, len(orders))
	for _, order := range orders {
		if order.Status == 1 || order.Status == 2 {
			stdHours := order.Quantity * 0.5
			dueDate := time.Now().Add(72 * time.Hour)
			if order.PlanEndDate != nil {
				dueDate = *order.PlanEndDate
			}
			scheduleOrders = append(scheduleOrders, scheduler.ScheduleOrder{
				OrderID:       order.ID,
				OrderNo:       order.OrderNo,
				ProductID:     order.MaterialID,
				ProductCode:   order.MaterialCode,
				ProductName:   order.MaterialName,
				Quantity:      order.Quantity,
				StandardHours: stdHours,
				DueDate:      dueDate,
				Priority:     order.Priority,
			})
		}
	}
	return scheduleOrders
}

// convertToWorkCenters 转换产线为工作中心
func (s *ScheduleService) convertToWorkCenters(lines []model.ProductionLine) []scheduler.WorkCenterInfo {
	workCenters := make([]scheduler.WorkCenterInfo, 0, len(lines))
	for _, line := range lines {
		workCenters = append(workCenters, scheduler.WorkCenterInfo{
			WorkCenterID:   line.ID,
			WorkCenterName: line.LineName,
			LineID:        line.ID,
			LineName:       line.LineName,
			Capacity:      1.0,
			Efficiency:    100,
		})
	}
	return workCenters
}
