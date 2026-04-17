package scheduler

import (
	"context"
	"fmt"
	"sort"
	"time"

	"mom-server/internal/model"
)

// SchedulingRule 排程规则
type SchedulingRule string

const (
	RuleFIFO SchedulingRule = "FIFO" // 先进先出：按创建时间
	RuleEDD  SchedulingRule = "EDD"  // 最早交付日期
	RuleSPT  SchedulingRule = "SPT"  // 最短加工时间
	RuleLPT  SchedulingRule = "LPT"  // 最长加工时间
)

// SchedulingDirection 排程方向
type SchedulingDirection string

const (
	DirectionForward  SchedulingDirection = "FORWARD"  // 前向排程：从可用时间开始向后
	DirectionBackward SchedulingDirection = "BACKWARD" // 后向排程：从交付日期开始向前
)

// TaskInfo 任务信息
type TaskInfo struct {
	OrderID       int64
	OrderNo      string
	Priority     int
	Quantity     float64
	StandardHours float64 // 标准工时(小时)
	DueDate      time.Time
	CreateTime   time.Time
	ProcessCount int // 工序数量
}

// ResourceInfo 资源信息
type ResourceInfo struct {
	ResourceID   int64
	ResourceName string
	Capacity     float64 // 产能(每小时)
	Efficiency   float64 // 效率%
}

// ScheduledTask 排程后的任务
type ScheduledTask struct {
	TaskID         int64
	OrderID        int64
	OrderNo       string
	ResourceID     int64
	ResourceName   string
	Sequence       int
	PlanStartTime  time.Time
	PlanEndTime    time.Time
	StandardHours  float64
	LoadHours      float64 // 加载工时
}

// SchedulingResult 排程结果
type SchedulingResult struct {
	Tasks        []ScheduledTask
	TotalHours   float64
	StartTime    time.Time
	EndTime      time.Time
	Utilization  float64           // 资源利用率
	Suggestions  []OptimizationSuggestion
}

// Scheduler 排程器接口
type Scheduler interface {
	Schedule(tasks []TaskInfo, resources []ResourceInfo, direction SchedulingDirection) (*SchedulingResult, error)
}

// PriorityRuleScheduler 基于优先级规则的排程器
type PriorityRuleScheduler struct {
	Rule SchedulingRule
}

// NewPriorityRuleScheduler 创建优先级规则排程器
func NewPriorityRuleScheduler(rule SchedulingRule) *PriorityRuleScheduler {
	return &PriorityRuleScheduler{Rule: rule}
}

// Schedule 执行排程
func (s *PriorityRuleScheduler) Schedule(tasks []TaskInfo, resources []ResourceInfo, direction SchedulingDirection) (*SchedulingResult, error) {
	if len(tasks) == 0 {
		return &SchedulingResult{}, nil
	}

	if len(resources) == 0 {
		return nil, fmt.Errorf("资源列表为空")
	}

	// 根据规则排序任务
	sortedTasks := s.sortTasks(tasks)

	result := &SchedulingResult{
		Tasks: make([]ScheduledTask, 0, len(sortedTasks)),
	}

	// 初始化资源可用时间
	resourceAvailability := make(map[int64]time.Time)
	for _, r := range resources {
		resourceAvailability[r.ResourceID] = time.Now()
	}

	// 按顺序分配任务
	for i, task := range sortedTasks {
		// 选择最佳资源(负载最低)
		bestResource := s.selectBestResource(task, resources, resourceAvailability)

		var startTime, endTime time.Time

		if direction == DirectionForward {
			// 前向排程：从资源可用时间开始
			startTime = resourceAvailability[bestResource.ResourceID]
			loadHours := task.StandardHours / (bestResource.Efficiency / 100)
			endTime = startTime.Add(time.Duration(loadHours * float64(time.Hour)))
		} else {
			// 后向排程：从任务截止日期向前
			endTime = task.DueDate
			loadHours := task.StandardHours / (bestResource.Efficiency / 100)
			startTime = endTime.Add(-time.Duration(loadHours * float64(time.Hour)))
			// 确保不早于资源可用时间
			if startTime.Before(resourceAvailability[bestResource.ResourceID]) {
				startTime = resourceAvailability[bestResource.ResourceID]
				endTime = startTime.Add(time.Duration(loadHours * float64(time.Hour)))
			}
		}

		// 更新资源可用时间
		resourceAvailability[bestResource.ResourceID] = endTime

		scheduledTask := ScheduledTask{
			TaskID:        int64(i + 1),
			OrderID:       task.OrderID,
			OrderNo:      task.OrderNo,
			ResourceID:    bestResource.ResourceID,
			ResourceName:  bestResource.ResourceName,
			Sequence:      i + 1,
			PlanStartTime: startTime,
			PlanEndTime:   endTime,
			StandardHours: task.StandardHours,
			LoadHours:     task.StandardHours / (bestResource.Efficiency / 100),
		}
		result.Tasks = append(result.Tasks, scheduledTask)
	}

	// 计算统计信息
	s.calculateStats(result)

	return result, nil
}

// sortTasks 根据规则排序任务
func (s *PriorityRuleScheduler) sortTasks(tasks []TaskInfo) []TaskInfo {
	sorted := make([]TaskInfo, len(tasks))
	copy(sorted, tasks)

	switch s.Rule {
	case RuleFIFO:
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].CreateTime.Before(sorted[j].CreateTime)
		})
	case RuleEDD:
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].DueDate.Before(sorted[j].DueDate)
		})
	case RuleSPT:
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].StandardHours < sorted[j].StandardHours
		})
	case RuleLPT:
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].StandardHours > sorted[j].StandardHours
		})
	default:
		// 默认FIFO
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].CreateTime.Before(sorted[j].CreateTime)
		})
	}

	return sorted
}

// selectBestResource 选择最佳资源(负载最低且满足能力要求)
func (s *PriorityRuleScheduler) selectBestResource(task TaskInfo, resources []ResourceInfo, availability map[int64]time.Time) ResourceInfo {
	best := resources[0]
	bestLoad := float64(0)

	for _, r := range resources {
		var load float64
		if r.Capacity > 0 {
			load = float64(availability[r.ResourceID].Unix()) / r.Capacity
		}
		if bestLoad == 0 || load < bestLoad {
			best = r
			bestLoad = load
		}
	}

	return best
}

// calculateStats 计算统计信息
func (s *PriorityRuleScheduler) calculateStats(result *SchedulingResult) {
	if len(result.Tasks) == 0 {
		return
	}

	var totalLoad float64
	result.StartTime = result.Tasks[0].PlanStartTime
	result.EndTime = result.Tasks[0].PlanEndTime

	for _, t := range result.Tasks {
		totalLoad += t.LoadHours
		if t.PlanStartTime.Before(result.StartTime) {
			result.StartTime = t.PlanStartTime
		}
		if t.PlanEndTime.After(result.EndTime) {
			result.EndTime = t.PlanEndTime
		}
	}

	result.TotalHours = totalLoad

	// 计算时间跨度内的资源利用率(简化计算)
	timeSpan := result.EndTime.Sub(result.StartTime).Hours()
	if timeSpan > 0 {
		result.Utilization = (totalLoad / timeSpan) * 100
		if result.Utilization > 100 {
			result.Utilization = 100
		}
	}
}

// ForwardScheduler 前向排程器
type ForwardScheduler struct {
	*PriorityRuleScheduler
}

// NewForwardScheduler 创建前向排程器
func NewForwardScheduler(rule SchedulingRule) *ForwardScheduler {
	return &ForwardScheduler{
		PriorityRuleScheduler: NewPriorityRuleScheduler(rule),
	}
}

// Schedule 执行前向排程(从可用时间开始向后)
func (s *ForwardScheduler) Schedule(tasks []TaskInfo, resources []ResourceInfo, _ SchedulingDirection) (*SchedulingResult, error) {
	return s.PriorityRuleScheduler.Schedule(tasks, resources, DirectionForward)
}

// BackwardScheduler 后向排程器
type BackwardScheduler struct {
	*PriorityRuleScheduler
}

// NewBackwardScheduler 创建后向排程器
func NewBackwardScheduler(rule SchedulingRule) *BackwardScheduler {
	return &BackwardScheduler{
		PriorityRuleScheduler: NewPriorityRuleScheduler(rule),
	}
}

// Schedule 执行后向排程(从截止日期开始向前)
func (s *BackwardScheduler) Schedule(tasks []TaskInfo, resources []ResourceInfo, _ SchedulingDirection) (*SchedulingResult, error) {
	return s.PriorityRuleScheduler.Schedule(tasks, resources, DirectionBackward)
}

// SchedulingService 排程服务
type SchedulingService struct {
	scheduler Scheduler
}

// NewSchedulingService 创建排程服务
func NewSchedulingService(scheduler Scheduler) *SchedulingService {
	return &SchedulingService{scheduler: scheduler}
}

// ExecuteScheduling 执行排程
func (s *SchedulingService) ExecuteScheduling(ctx context.Context, tasks []TaskInfo, resources []ResourceInfo, direction SchedulingDirection) (*SchedulingResult, error) {
	return s.scheduler.Schedule(tasks, resources, direction)
}

// CalculateOrder 计算工单优先级得分（供约束排程使用）
// 返回值越小优先级越高
func CalculateOrder(order ScheduleOrder, algorithmType string) float64 {
	switch SchedulingRule(algorithmType) {
	case RuleFamily:
		// FAMILY: 按产品族分组，族内按优先级
		if order.FamilyCode == "" {
			return 999 // 无产品族信息优先级最低
		}
		return float64(order.Priority) * 1000 // 族内按优先级排序

	case RuleBottleneck:
		// BOTTLENECK: 瓶颈工序优先，然后按优先级
		if order.Bottleneck {
			return float64(order.Priority)
		}
		return float64(order.Priority) + 1000 // 非瓶颈排后面

	case RuleCR:
		// CR (Critical Ratio): (交期-当前)/剩余工时
		// 值越小越紧急
		if order.StandardHours <= 0 {
			return 999
		}
		now := time.Now()
		hoursUntilDue := order.DueDate.Sub(now).Hours()
		if hoursUntilDue < 0 {
			hoursUntilDue = 0 // 已逾期，标记为最紧急
		}
		return hoursUntilDue / order.StandardHours

	case RuleJITFirst:
		// JIT_FIRST: JIT需求时间优先
		if order.JITTime == nil {
			return float64(order.Priority) * 10000 // 无JIT时间按优先级
		}
		now := time.Now()
		// JIT时间距离现在越近，优先级越高
		hoursUntilJIT := order.JITTime.Sub(now).Hours()
		if hoursUntilJIT < 0 {
			hoursUntilJIT = 0 // JIT时间已过，标记为最紧急
		}
		return hoursUntilJIT

	case RuleEDD:
		// EDD: 最早交付日期优先
		return float64(order.DueDate.Unix())

	case RuleSPT:
		// SPT: 最短加工时间优先
		return order.StandardHours

	case RuleLPT:
		// LPT: 最长加工时间优先
		return -order.StandardHours

	case RuleFIFO:
		// FIFO: 按工单ID顺序
		return float64(order.OrderID)

	default:
		return float64(order.Priority) * 10000
	}
}

// CompareOrders 比较两个工单的优先级
// 返回true表示o1比o2优先级高(应该排在前面的)
func CompareOrders(o1, o2 ScheduleOrder, algorithmType string) bool {
	score1 := CalculateOrder(o1, algorithmType)
	score2 := CalculateOrder(o2, algorithmType)
	return score1 < score2
}

// ConvertToScheduleResults 转换为数据库模型
func ConvertToScheduleResults(planID int64, result *SchedulingResult) []*model.ScheduleResult {
	results := make([]*model.ScheduleResult, len(result.Tasks))
	for i, t := range result.Tasks {
		r := &model.ScheduleResult{
			PlanID:         planID,
			OrderID:        t.OrderID,
			OrderNo:        t.OrderNo,
			Sequence:       t.Sequence,
			LineID:         t.ResourceID,
			PlanStartTime:  &t.PlanStartTime,
			PlanEndTime:    &t.PlanEndTime,
		}
		lineName := t.ResourceName
		r.LineName = &lineName
		results[i] = r
	}
	return results
}

// CalculateLoad 计算产能负载
func CalculateLoad(tasks []ScheduledTask, resourceID int64, startTime, endTime time.Time) float64 {
	var load float64
	for _, t := range tasks {
		if t.ResourceID == resourceID {
			// 计算在指定时间范围内的负载
			taskStart := t.PlanStartTime
			taskEnd := t.PlanEndTime

			// 如果任务在时间范围内，则计算负载
			if taskEnd.After(startTime) && taskStart.Before(endTime) {
				load += t.LoadHours
			}
		}
	}
	return load
}

// OptimizeSchedule 优化排程结果
func OptimizeSchedule(result *SchedulingResult, resources []ResourceInfo) *SchedulingResult {
	if result == nil || len(result.Tasks) == 0 {
		return result
	}

	// 简单优化：尝试平衡资源负载
	optimized := &SchedulingResult{
		Tasks:      make([]ScheduledTask, len(result.Tasks)),
		StartTime:  result.StartTime,
		EndTime:    result.EndTime,
	}

	copy(optimized.Tasks, result.Tasks)

	// 计算每个资源的当前负载
	resourceLoads := make(map[int64]float64)
	for _, t := range result.Tasks {
		resourceLoads[t.ResourceID] += t.LoadHours
	}

	// 找出最负载和最空闲的资源
	var maxLoadResource, minLoadResource int64
	var maxLoad, minLoad float64

	for rid, load := range resourceLoads {
		if maxLoad == 0 || load > maxLoad {
			maxLoad = load
			maxLoadResource = rid
		}
		if minLoad == 0 || load < minLoad {
			minLoad = load
			minLoadResource = rid
		}
	}

	// 如果负载差异超过20%，尝试迁移任务
	if maxLoad > 0 && minLoad > 0 && (maxLoad-minLoad)/maxLoad > 0.2 {
		// 从高负载资源迁移任务到低负载资源
		for i := range optimized.Tasks {
			if optimized.Tasks[i].ResourceID == maxLoadResource && optimized.Tasks[i].Sequence > 1 {
				// 找到最晚在该任务之前结束的低负载资源任务
				for j := range optimized.Tasks {
					if optimized.Tasks[j].ResourceID == minLoadResource {
						// 交换资源分配(简化处理)
						optimized.Tasks[i].ResourceID, optimized.Tasks[j].ResourceID = minLoadResource, maxLoadResource
						optimized.Tasks[i].ResourceName, optimized.Tasks[j].ResourceName = resources[j].ResourceName, resources[i].ResourceName
						break
					}
				}
				break
			}
		}
	}

	return optimized
}
