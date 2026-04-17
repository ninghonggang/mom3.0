package scheduler

import (
	"context"
	"fmt"
	"sort"
	"time"
)

// ConstrainedScheduler 带约束的排程器
type ConstrainedScheduler struct {
	calendarGetter func(workshopID int64) (interface{}, error) // returns *WorkingCalendar
	changeover    *ChangeoverAwareScheduler
}

// NewConstrainedScheduler 创建约束排程器
func NewConstrainedScheduler(
	calendarGetter func(workshopID int64) (interface{}, error),
	changeover *ChangeoverAwareScheduler,
) *ConstrainedScheduler {
	return &ConstrainedScheduler{
		calendarGetter: calendarGetter,
		changeover:    changeover,
	}
}

// Schedule 执行带约束的排程
func (s *ConstrainedScheduler) Schedule(ctx context.Context, req *ScheduleRequest) (*SchedulingResult, error) {
	if len(req.Orders) == 0 {
		return &SchedulingResult{}, nil
	}

	if len(req.WorkCenters) == 0 {
		return nil, fmt.Errorf("工作中心列表为空")
	}

	// 1. 获取日历，计算每个资源可用时段
	workHoursPerDay := 8.0
	if s.calendarGetter != nil && req.WorkshopID > 0 {
		if cal, err := s.calendarGetter(req.WorkshopID); err == nil && cal != nil {
			// 使用接口方式获取工时
			if wc, ok := cal.(interface{ GetWorkHoursPerDay() float64 }); ok {
				workHoursPerDay = wc.GetWorkHoursPerDay()
			}
		}
	}

	// 2. 按规则排序任务
	sortedOrders := s.sortOrders(req)

	// 3. 考虑换型时间分配任务
	result := s.allocateTasksWithChangeover(sortedOrders, req.WorkCenters, req.Direction, workHoursPerDay)

	// 4. 计算统计信息
	s.calculateConstrainedStats(result, req)

	// 5. 生成优化建议
	result.Suggestions = s.generateSuggestions(req, result)

	return result, nil
}

// sortOrders 按算法类型排序工单（统一使用 CalculateOrder）
func (s *ConstrainedScheduler) sortOrders(req *ScheduleRequest) []ScheduleOrder {
	orders := make([]ScheduleOrder, len(req.Orders))
	copy(orders, req.Orders)

	// 如果启用产品族聚类，使用换型优化
	if req.Constraints.FamilyGrouping && s.changeover != nil {
		return s.changeover.OptimizeByChangeover(orders)
	}

	algorithm := SchedulingRule(req.AlgorithmType)

	// 统一使用 CalculateOrder 计算优先级得分进行排序
	sort.Slice(orders, func(i, j int) bool {
		scoreI := CalculateOrder(orders[i], string(algorithm))
		scoreJ := CalculateOrder(orders[j], string(algorithm))
		return scoreI < scoreJ
	})

	return orders
}

// allocateTasksWithChangeover 分配任务（考虑换型）
func (s *ConstrainedScheduler) allocateTasksWithChangeover(orders []ScheduleOrder, resources []WorkCenterInfo, direction SchedulingDirection, workHoursPerDay float64) *SchedulingResult {
	result := &SchedulingResult{
		Suggestions: make([]OptimizationSuggestion, 0),
	}

	// 初始化资源可用时间
	resourceAvailability := make(map[int64]time.Time)
	for _, r := range resources {
		resourceAvailability[r.LineID] = time.Now()
	}

	var lastProductCode string
	for i, order := range orders {
		// 选择最佳资源
		bestResource := s.selectBestResource(order, resources, resourceAvailability)

		// 计算换型时间
		changeoverTime := 0.0
		if lastProductCode != "" && lastProductCode != order.ProductCode && s.changeover != nil {
			changeoverTime = s.changeover.CalculateChangeoverTime(lastProductCode, order.ProductCode)
		}
		lastProductCode = order.ProductCode

		var startTime, endTime time.Time
		loadHours := order.StandardHours / (bestResource.Efficiency / 100)

		if direction == DirectionForward {
			startTime = resourceAvailability[bestResource.LineID]
			if changeoverTime > 0 {
				// 换型时间追加
				startTime = startTime.Add(time.Duration(changeoverTime * float64(time.Minute)))
			}
			endTime = startTime.Add(time.Duration(loadHours * float64(time.Hour)))
		} else {
			endTime = order.DueDate
			if changeoverTime > 0 {
				endTime = endTime.Add(-time.Duration(changeoverTime * float64(time.Minute)))
			}
			loadHoursCalc := loadHours
			endTime = endTime.Add(-time.Duration(loadHoursCalc * float64(time.Hour)))
			startTime = endTime
			if startTime.Before(resourceAvailability[bestResource.LineID]) {
				startTime = resourceAvailability[bestResource.LineID]
				endTime = startTime.Add(time.Duration(loadHours * float64(time.Hour)))
			}
		}

		resourceAvailability[bestResource.LineID] = endTime

		result.Tasks = append(result.Tasks, ScheduledTask{
			TaskID:        int64(i + 1),
			OrderID:       order.OrderID,
			OrderNo:      order.OrderNo,
			ResourceID:    bestResource.LineID,
			ResourceName:  bestResource.LineName,
			Sequence:      i + 1,
			PlanStartTime: startTime,
			PlanEndTime:   endTime,
			StandardHours: order.StandardHours,
			LoadHours:     loadHours,
		})
	}

	s.calculateStats(result)
	return result
}

// selectBestResource 选择最佳资源
func (s *ConstrainedScheduler) selectBestResource(order ScheduleOrder, resources []WorkCenterInfo, availability map[int64]time.Time) WorkCenterInfo {
	best := resources[0]
	minLoad := float64(0)

	for _, r := range resources {
		load := float64(availability[r.LineID].Unix())
		if minLoad == 0 || load < minLoad {
			best = r
			minLoad = load
		}
	}
	return best
}

// calculateConstrainedStats 计算约束相关的统计
func (s *ConstrainedScheduler) calculateConstrainedStats(result *SchedulingResult, req *ScheduleRequest) {
	if len(result.Tasks) == 0 {
		return
	}

	result.StartTime = result.Tasks[0].PlanStartTime
	result.EndTime = result.Tasks[0].PlanEndTime

	for _, t := range result.Tasks {
		result.TotalHours += t.LoadHours
		if t.PlanStartTime.Before(result.StartTime) {
			result.StartTime = t.PlanStartTime
		}
		if t.PlanEndTime.After(result.EndTime) {
			result.EndTime = t.PlanEndTime
		}
	}

	// 计算利用率
	timeSpan := result.EndTime.Sub(result.StartTime).Hours()
	if timeSpan > 0 {
		result.Utilization = (result.TotalHours / timeSpan) * 100
		if result.Utilization > 100 {
			result.Utilization = 100
		}
	}
}

// calculateStats 计算统计信息
func (s *ConstrainedScheduler) calculateStats(result *SchedulingResult) {
	if len(result.Tasks) == 0 {
		return
	}

	result.StartTime = result.Tasks[0].PlanStartTime
	result.EndTime = result.Tasks[0].PlanEndTime

	for _, t := range result.Tasks {
		result.TotalHours += t.LoadHours
		if t.PlanStartTime.Before(result.StartTime) {
			result.StartTime = t.PlanStartTime
		}
		if t.PlanEndTime.After(result.EndTime) {
			result.EndTime = t.PlanEndTime
		}
	}
}

// generateSuggestions 生成优化建议
func (s *ConstrainedScheduler) generateSuggestions(req *ScheduleRequest, result *SchedulingResult) []OptimizationSuggestion {
	suggestions := make([]OptimizationSuggestion, 0)

	// 检查利用率
	if result.Utilization < float64(req.Constraints.MinUtilization) {
		suggestions = append(suggestions, OptimizationSuggestion{
			Type:    "CAPACITY",
			Level:   "WARNING",
			Message: fmt.Sprintf("资源利用率 %.1f%% 低于目标 %.1f%%", result.Utilization, req.Constraints.MinUtilization),
			Impact:  "产能过剩，建议减少资源投入或增加订单",
		})
	}

	// 检查交付率
	for _, task := range result.Tasks {
		for _, order := range req.Orders {
			if task.OrderID == order.OrderID && task.PlanEndTime.After(order.DueDate) {
				suggestions = append(suggestions, OptimizationSuggestion{
					Type:    "DELIVERY",
					Level:   "CRITICAL",
					Message: fmt.Sprintf("订单 %s 计划结束时间晚于交期", order.OrderNo),
					Impact:  "将导致逾期交付",
				})
			}
		}
	}

	return suggestions
}
