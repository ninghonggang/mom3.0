package scheduler

import (
	"fmt"
	"sort"
)

// ChangeoverAwareScheduler 换型感知排程器
type ChangeoverAwareScheduler struct {
	changeoverMatrix map[string]float64 // key: "from_product-to_product"
	familyGroups     map[string][]string // 产品族成员
	defaultChangeover float64             // 默认换型时间(分钟)
}

// NewChangeoverAwareScheduler 创建换型感知排程器
func NewChangeoverAwareScheduler(matrix map[string]float64, families map[string][]string) *ChangeoverAwareScheduler {
	return &ChangeoverAwareScheduler{
		changeoverMatrix: matrix,
		familyGroups:     families,
		defaultChangeover: 10.0, // 默认10分钟
	}
}

// CalculateChangeoverTime 计算换型时间
func (s *ChangeoverAwareScheduler) CalculateChangeoverTime(fromProduct, toProduct string) float64 {
	if fromProduct == "" || toProduct == "" || fromProduct == toProduct {
		return 0
	}

	key := fmt.Sprintf("%s-%s", fromProduct, toProduct)
	if time, ok := s.changeoverMatrix[key]; ok {
		return time
	}

	// 查反向
	reverseKey := fmt.Sprintf("%s-%s", toProduct, fromProduct)
	if time, ok := s.changeoverMatrix[reverseKey]; ok {
		return time
	}

	// 同产品族减少换型时间
	if s.areSameFamily(fromProduct, toProduct) {
		return s.defaultChangeover * 0.5 // 同族减半
	}

	return s.defaultChangeover
}

// areSameFamily 判断是否同产品族
func (s *ChangeoverAwareScheduler) areSameFamily(product1, product2 string) bool {
	if s.familyGroups == nil {
		return false
	}
	for _, products := range s.familyGroups {
		for _, p1 := range products {
			if p1 == product1 {
				for _, p2 := range products {
					if p2 == product2 {
						return true
					}
				}
			}
		}
	}
	return false
}

// GroupByFamily 按产品族聚类任务
func (s *ChangeoverAwareScheduler) GroupByFamily(tasks []ScheduleOrder) [][]ScheduleOrder {
	if len(tasks) == 0 {
		return nil
	}

	// 按产品族分组
	familyMap := make(map[string][]ScheduleOrder)
	var noFamily []ScheduleOrder

	for _, task := range tasks {
		if task.FamilyCode == "" {
			noFamily = append(noFamily, task)
		} else {
			familyMap[task.FamilyCode] = append(familyMap[task.FamilyCode], task)
		}
	}

	// 排序保持稳定性
	var families []string
	for family := range familyMap {
		families = append(families, family)
	}
	sort.Strings(families)

	result := make([][]ScheduleOrder, 0, len(familyMap)+1)
	for _, family := range families {
		result = append(result, familyMap[family])
	}
	if len(noFamily) > 0 {
		result = append(result, noFamily)
	}

	return result
}

// OptimizeByChangeover 换型优化排序
func (s *ChangeoverAwareScheduler) OptimizeByChangeover(tasks []ScheduleOrder) []ScheduleOrder {
	if len(tasks) <= 1 {
		return tasks
	}

	// 先按产品族聚类
	groups := s.GroupByFamily(tasks)

	result := make([]ScheduleOrder, 0, len(tasks))
	for _, group := range groups {
		// 每组内按CR值排序
		sort.Slice(group, func(i, j int) bool {
			return s.calculateCR(group[i]) < s.calculateCR(group[j])
		})
		result = append(result, group...)
	}

	return result
}

// calculateCR 计算紧迫系数 CR = (交期-当前)/剩余工时
func (s *ChangeoverAwareScheduler) calculateCR(task ScheduleOrder) float64 {
	// 简化实现
	if task.StandardHours <= 0 {
		return 999
	}
	// 按优先级和标准工时估算
	return float64(task.Priority) * 100 / task.StandardHours
}
