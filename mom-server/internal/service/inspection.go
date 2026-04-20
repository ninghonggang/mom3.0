package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type InspectionService struct {
	repo *repository.InspectionRepository
}

func NewInspectionService(repo *repository.InspectionRepository) *InspectionService {
	return &InspectionService{repo: repo}
}

// ListPlans 查询检验计划列表
func (s *InspectionService) ListPlans(ctx context.Context, tenantID uint64, query map[string]interface{}) ([]model.QualityInspectionPlan, int64, error) {
	return s.repo.ListPlans(ctx, tenantID, query)
}

// GetPlanByID 根据ID获取检验计划
func (s *InspectionService) GetPlanByID(ctx context.Context, id uint64) (*model.QualityInspectionPlan, error) {
	return s.repo.GetPlanByID(ctx, id)
}

// CreatePlan 创建检验计划
func (s *InspectionService) CreatePlan(ctx context.Context, tenantID uint64, req *model.QualityInspectionPlanCreateRequest, username string) (*model.QualityInspectionPlan, error) {
	plan := &model.QualityInspectionPlan{
		TenantID:        tenantID,
		PlanCode:       req.PlanCode,
		PlanName:       req.PlanName,
		InspectionType: req.InspectionType,
		AQLLevel:       req.AQLLevel,
		SampleSize:     req.SampleSize,
		BatchMin:       req.BatchMin,
		BatchMax:       req.BatchMax,
		AcCount:        req.AcCount,
		ReCount:        req.ReCount,
		CheckItems:     req.CheckItems,
		Status:         req.Status,
		Remark:         req.Remark,
		CreatedBy:      username,
		UpdatedBy:      username,
	}

	// 默认状态
	if plan.Status == "" {
		plan.Status = "ACTIVE"
	}

	err := s.repo.CreatePlan(ctx, plan)
	if err != nil {
		return nil, fmt.Errorf("failed to create inspection plan: %w", err)
	}
	return plan, nil
}

// UpdatePlan 更新检验计划
func (s *InspectionService) UpdatePlan(ctx context.Context, id uint64, req *model.QualityInspectionPlanUpdateRequest, username string) error {
	plan, err := s.repo.GetPlanByID(ctx, id)
	if err != nil {
		return fmt.Errorf("plan not found: %w", err)
	}

	// 更新字段
	if req.PlanName != "" {
		plan.PlanName = req.PlanName
	}
	if req.AQLLevel != "" {
		plan.AQLLevel = req.AQLLevel
	}
	if req.SampleSize > 0 {
		plan.SampleSize = req.SampleSize
	}
	if req.BatchMin > 0 {
		plan.BatchMin = req.BatchMin
	}
	if req.BatchMax > 0 {
		plan.BatchMax = req.BatchMax
	}
	if req.AcCount > 0 {
		plan.AcCount = req.AcCount
	}
	if req.ReCount > 0 {
		plan.ReCount = req.ReCount
	}
	if req.CheckItems != "" {
		plan.CheckItems = req.CheckItems
	}
	if req.Status != "" {
		plan.Status = req.Status
	}
	if req.Remark != "" {
		plan.Remark = req.Remark
	}
	plan.UpdatedBy = username

	return s.repo.UpdatePlan(ctx, plan)
}

// DeletePlan 删除检验计划
func (s *InspectionService) DeletePlan(ctx context.Context, id uint64) error {
	return s.repo.DeletePlan(ctx, id)
}

// CalculateSampleSize 根据检验类型、批量大小计算抽样方案
func (s *InspectionService) CalculateSampleSize(ctx context.Context, tenantID uint64, inspectionType string, batchSize int) (*model.AQLSampleSize, error) {
	// 根据检验类型确定默认检验水平
	level := "II" // 默认使用一般检验水平II
	switch inspectionType {
	case "IQC", "IQC_INCOMING":
		level = "I"
	case "FQC", "OQC":
		level = "II"
	case "IPQC", "PQC":
		level = "S-2"
	default:
		level = "II"
	}

	// 获取样本量字码
	sampleSizeCode := repository.GenerateSampleSizeCode(batchSize)

	// 根据批量大小确定AQL范围
	var aqlValue float64 = 1.0 // 默认AQL 1.0

	// 查询AQL抽样表
	result, err := s.repo.GetAQLSampleSizeByCode(ctx, tenantID, sampleSizeCode, aqlValue, level)
	if err != nil {
		// 如果没有找到精确匹配，尝试根据批量范围查找
		result, err = s.repo.GetAQLSampleSizeByBatchSize(ctx, tenantID, batchSize, level)
		if err != nil {
			return nil, fmt.Errorf("no AQL sample size found for batch size %d and level %s: %w", batchSize, level, err)
		}
	}

	return result, nil
}

// SeedAQLData 初始化AQL标准抽样数据
func (s *InspectionService) SeedAQLData(ctx context.Context) error {
	tenantID := uint64(1)

	// 检查是否已有数据
	count, err := s.repo.CountAQLData(ctx, tenantID)
	if err != nil {
		return fmt.Errorf("failed to count AQL data: %w", err)
	}
	if count > 0 {
		return nil // 已有数据，无需初始化
	}

	// 生成GB/T 2828.1标准抽样数据
	data := generateAQLSampleSizeData(tenantID)
	return s.repo.SeedData(ctx, data)
}

// generateAQLSampleSizeData 生成AQL标准抽样数据（GB/T 2828.1简化版）
func generateAQLSampleSizeData(tenantID uint64) []model.AQLSampleSize {
	var data []model.AQLSampleSize

	// 检验水平: I, II, III, S-1, S-2, S-3
	// 样本量字码: D, E, F, G, H, J, K, L, M, N, P, Q, R, S, T
	// AQL值: 0.065, 0.10, 0.15, 0.25, 0.40, 0.65, 1.0, 1.5, 2.5, 4.0

	batchRanges := []struct {
		min      int
		max      int
		code     string
	}{
		{2, 8, "D"},
		{9, 15, "E"},
		{16, 25, "F"},
		{26, 50, "G"},
		{51, 90, "H"},
		{91, 150, "J"},
		{151, 280, "K"},
		{281, 500, "L"},
		{501, 1200, "M"},
		{1201, 3200, "N"},
		{3201, 10000, "P"},
		{10001, 35000, "Q"},
		{35001, 150000, "R"},
		{150001, 500000, "S"},
		{500001, 999999999, "T"},
	}

	levels := []string{"S-1", "S-2", "S-3", "I", "II", "III"}

	// 样本量表
	sampleSizes := map[string]int{
		"D": 2, "E": 3, "F": 5, "G": 8, "H": 13, "J": 20, "K": 32,
		"L": 50, "M": 80, "N": 125, "P": 200, "Q": 315, "R": 500, "S": 800, "T": 1250,
	}

	aqlValues := []float64{0.065, 0.10, 0.15, 0.25, 0.40, 0.65, 1.0, 1.5, 2.5, 4.0}

	// Ac/Re 查表简化版
	acReTable := map[string]map[float64]struct{ ac, re int }{
		"D": {0.065: {0, 1}, 0.10: {0, 1}, 0.15: {0, 1}, 0.25: {0, 1}, 0.40: {0, 1}, 0.65: {0, 1}, 1.0: {0, 1}, 1.5: {0, 1}, 2.5: {0, 1}, 4.0: {1, 2}},
		"E": {0.065: {0, 1}, 0.10: {0, 1}, 0.15: {0, 1}, 0.25: {0, 1}, 0.40: {0, 1}, 0.65: {0, 1}, 1.0: {0, 1}, 1.5: {1, 2}, 2.5: {1, 2}, 4.0: {1, 2}},
		"F": {0.065: {0, 1}, 0.10: {0, 1}, 0.15: {0, 1}, 0.25: {0, 1}, 0.40: {0, 1}, 0.65: {1, 2}, 1.0: {1, 2}, 1.5: {1, 2}, 2.5: {2, 3}, 4.0: {2, 3}},
		"G": {0.065: {0, 1}, 0.10: {0, 1}, 0.15: {0, 1}, 0.25: {0, 1}, 0.40: {1, 2}, 0.65: {1, 2}, 1.0: {1, 2}, 1.5: {2, 3}, 2.5: {3, 4}, 4.0: {3, 4}},
		"H": {0.065: {0, 1}, 0.10: {0, 1}, 0.15: {0, 1}, 0.25: {1, 2}, 0.40: {1, 2}, 0.65: {1, 2}, 1.0: {2, 3}, 1.5: {3, 4}, 2.5: {4, 5}, 4.0: {5, 6}},
		"J": {0.065: {0, 1}, 0.10: {0, 1}, 0.15: {1, 2}, 0.25: {1, 2}, 0.40: {1, 2}, 0.65: {2, 3}, 1.0: {3, 4}, 1.5: {4, 5}, 2.5: {6, 7}, 4.0: {7, 8}},
		"K": {0.065: {0, 1}, 0.10: {1, 2}, 0.15: {1, 2}, 0.25: {1, 2}, 0.40: {2, 3}, 0.65: {3, 4}, 1.0: {4, 5}, 1.5: {6, 7}, 2.5: {7, 8}, 4.0: {9, 10}},
		"L": {0.065: {1, 2}, 0.10: {1, 2}, 0.15: {1, 2}, 0.25: {2, 3}, 0.40: {3, 4}, 0.65: {4, 5}, 1.0: {5, 6}, 1.5: {7, 8}, 2.5: {10, 11}, 4.0: {12, 13}},
		"M": {0.065: {1, 2}, 0.10: {1, 2}, 0.15: {2, 3}, 0.25: {3, 4}, 0.40: {4, 5}, 0.65: {5, 6}, 1.0: {7, 8}, 1.5: {10, 11}, 2.5: {14, 15}, 4.0: {18, 19}},
		"N": {0.065: {1, 2}, 0.10: {2, 3}, 0.15: {3, 4}, 0.25: {4, 5}, 0.40: {6, 7}, 0.65: {7, 8}, 1.0: {10, 11}, 1.5: {14, 15}, 2.5: {18, 19}, 4.0: {23, 24}},
		"P": {0.065: {2, 3}, 0.10: {3, 4}, 0.15: {4, 5}, 0.25: {6, 7}, 0.40: {8, 9}, 0.65: {10, 11}, 1.0: {14, 15}, 1.5: {18, 19}, 2.5: {23, 24}, 4.0: {30, 31}},
		"Q": {0.065: {3, 4}, 0.10: {4, 5}, 0.15: {6, 7}, 0.25: {8, 9}, 0.40: {10, 11}, 0.65: {14, 15}, 1.0: {18, 19}, 1.5: {23, 24}, 2.5: {30, 31}, 4.0: {39, 40}},
		"R": {0.065: {4, 5}, 0.10: {5, 6}, 0.15: {7, 8}, 0.25: {10, 11}, 0.40: {13, 14}, 0.65: {18, 19}, 1.0: {23, 24}, 1.5: {30, 31}, 2.5: {39, 40}, 4.0: {50, 51}},
		"S": {0.065: {5, 6}, 0.10: {7, 8}, 0.15: {9, 10}, 0.25: {13, 14}, 0.40: {18, 19}, 0.65: {23, 24}, 1.0: {30, 31}, 1.5: {39, 40}, 2.5: {50, 51}, 4.0: {63, 64}},
		"T": {0.065: {7, 8}, 0.10: {9, 10}, 0.15: {13, 14}, 0.25: {18, 19}, 0.40: {23, 24}, 0.65: {30, 31}, 1.0: {39, 40}, 1.5: {50, 51}, 2.5: {63, 64}, 4.0: {80, 81}},
	}

	for _, br := range batchRanges {
		for _, level := range levels {
			for _, aql := range aqlValues {
				acRe := acReTable[br.code][aql]
				ss := sampleSizes[br.code]
				data = append(data, model.AQLSampleSize{
					TenantID:        tenantID,
					SampleSizeCode:  br.code,
					BatchSizeMin:    br.min,
					BatchSizeMax:    br.max,
					InspectionLevel: level,
					AQLValue:        aql,
					SampleSize:      ss,
					Ac1:             acRe.ac,
					Re1:             acRe.re,
				})
			}
		}
	}

	return data
}
