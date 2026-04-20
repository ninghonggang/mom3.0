package service

import (
	"context"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"time"
)

// EAMInspectionService 巡检服务
type EAMInspectionService struct {
	planRepo    *repository.EAMInspectionPlanRepository
	itemRepo    *repository.EAMInspectionItemRepository
	schemeRepo  *repository.EAMInspectionSchemeRepository
	resultRepo  *repository.EAMInspectionResultRepository
}

// NewEAMInspectionService 创建巡检服务
func NewEAMInspectionService(
	planRepo *repository.EAMInspectionPlanRepository,
	itemRepo *repository.EAMInspectionItemRepository,
	schemeRepo *repository.EAMInspectionSchemeRepository,
	resultRepo *repository.EAMInspectionResultRepository,
) *EAMInspectionService {
	return &EAMInspectionService{
		planRepo:   planRepo,
		itemRepo:   itemRepo,
		schemeRepo: schemeRepo,
		resultRepo: resultRepo,
	}
}

// ========== 巡检计划 ==========

// CreatePlan 创建巡检计划
func (s *EAMInspectionService) CreatePlan(ctx context.Context, tenantID int64, req *model.EAMInspectionPlanCreateReqVO) (*model.EAMInspectionPlan, error) {
	plan := &model.EAMInspectionPlan{
		TenantID:          tenantID,
		PlanNo:            req.PlanNo,
		PlanName:          req.PlanName,
		EquipmentCategory: req.EquipmentCategory,
		CycleType:        req.CycleType,
		CycleDays:        req.CycleDays,
		Status:           "ACTIVE",
		Remark:           req.Remark,
	}
	if req.StartDate != "" {
		plan.StartDate = parseTime(req.StartDate)
	}
	if err := s.planRepo.Create(ctx, plan); err != nil {
		return nil, err
	}
	return plan, nil
}

// CreatePlanWithItems 创建巡检计划及项目
func (s *EAMInspectionService) CreatePlanWithItems(ctx context.Context, tenantID int64, req *model.EAMInspectionPlanCreateReqVO) (*model.EAMInspectionPlan, error) {
	plan := &model.EAMInspectionPlan{
		TenantID:          tenantID,
		PlanNo:            req.PlanNo,
		PlanName:          req.PlanName,
		EquipmentCategory: req.EquipmentCategory,
		CycleType:        req.CycleType,
		CycleDays:        req.CycleDays,
		Status:           "ACTIVE",
		Remark:           req.Remark,
	}
	if req.StartDate != "" {
		plan.StartDate = parseTime(req.StartDate)
	}
	if err := s.planRepo.Create(ctx, plan); err != nil {
		return nil, err
	}

	// 创建巡检项目
	if len(req.Items) > 0 {
		items := make([]model.EAMInspectionItem, len(req.Items))
		for i, item := range req.Items {
			items[i] = model.EAMInspectionItem{
				TenantID:      tenantID,
				PlanID:        plan.ID,
				ItemCode:      item.ItemCode,
				ItemName:      item.ItemName,
				CheckMethod:   item.CheckMethod,
				CheckStandard: item.CheckStandard,
				IsRequired:    item.IsRequired,
				SortOrder:     item.SortOrder,
			}
		}
		if err := s.itemRepo.CreateBatch(ctx, items); err != nil {
			return nil, err
		}
	}
	return plan, nil
}

// UpdatePlan 更新巡检计划
func (s *EAMInspectionService) UpdatePlan(ctx context.Context, id int64, req *model.EAMInspectionPlanCreateReqVO) error {
	updates := map[string]interface{}{
		"plan_name":           req.PlanName,
		"equipment_category":  req.EquipmentCategory,
		"cycle_type":          req.CycleType,
		"cycle_days":          req.CycleDays,
		"remark":              req.Remark,
	}
	if req.StartDate != "" {
		updates["start_date"] = parseTime(req.StartDate)
	}
	return s.planRepo.Update(ctx, id, updates)
}

// DeletePlan 删除巡检计划
func (s *EAMInspectionService) DeletePlan(ctx context.Context, id int64) error {
	return s.planRepo.Delete(ctx, id)
}

// GetPlan 获取巡检计划详情
func (s *EAMInspectionService) GetPlan(ctx context.Context, id int64) (*model.EAMInspectionPlan, error) {
	return s.planRepo.GetByID(ctx, id)
}

// GetPlanWithItems 获取巡检计划及项目
func (s *EAMInspectionService) GetPlanWithItems(ctx context.Context, id int64) (*model.EAMInspectionPlan, []model.EAMInspectionItem, error) {
	plan, err := s.planRepo.GetByID(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	items, err := s.itemRepo.ListByPlanID(ctx, id)
	if err != nil {
		return plan, nil, err
	}
	return plan, items, nil
}

// ListPlan 巡检计划列表
func (s *EAMInspectionService) ListPlan(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.EAMInspectionPlan, int64, error) {
	return s.planRepo.List(ctx, offset, limit, filters)
}

// UpdateItems 更新巡检项目
func (s *EAMInspectionService) UpdateItems(ctx context.Context, planID int64, items []model.EAMInspectionItem) error {
	return s.itemRepo.UpdateByPlanID(ctx, planID, items)
}

// ========== 巡检方案(执行) ==========

// CreateScheme 创建巡检方案
func (s *EAMInspectionService) CreateScheme(ctx context.Context, tenantID int64, req *model.EAMInspectionSchemeCreateReqVO) (*model.EAMInspectionScheme, error) {
	// 获取计划信息
	plan, err := s.planRepo.GetByID(ctx, req.PlanID)
	if err != nil {
		return nil, err
	}
	scheme := &model.EAMInspectionScheme{
		TenantID:       tenantID,
		PlanID:         req.PlanID,
		PlanNo:         plan.PlanNo,
		EquipmentID:    req.EquipmentID,
		EquipmentCode:  req.EquipmentCode,
		EquipmentName:  req.EquipmentName,
		Status:         "PENDING",
	}
	// 生成执行单号
	scheme.SchemeNo = generateSchemeNo()
	if req.InspectorID > 0 {
		scheme.InspectorID = &req.InspectorID
	}
	scheme.InspectorName = req.InspectorName
	if err := s.schemeRepo.Create(ctx, scheme); err != nil {
		return nil, err
	}
	return scheme, nil
}

// GetScheme 获取巡检方案
func (s *EAMInspectionService) GetScheme(ctx context.Context, id int64) (*model.EAMInspectionScheme, error) {
	return s.schemeRepo.GetByID(ctx, id)
}

// GetSchemeWithResults 获取巡检方案及结果
func (s *EAMInspectionService) GetSchemeWithResults(ctx context.Context, id int64) (*model.EAMInspectionScheme, []model.EAMInspectionResult, error) {
	scheme, err := s.schemeRepo.GetByID(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	results, err := s.resultRepo.ListBySchemeID(ctx, id)
	if err != nil {
		return scheme, nil, err
	}
	return scheme, results, nil
}

// ListScheme 巡检方案列表
func (s *EAMInspectionService) ListScheme(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.EAMInspectionScheme, int64, error) {
	return s.schemeRepo.List(ctx, offset, limit, filters)
}

// StartScheme 开始巡检
func (s *EAMInspectionService) StartScheme(ctx context.Context, id int64) error {
	updates := map[string]interface{}{
		"status": "IN_PROGRESS",
	}
	return s.schemeRepo.Update(ctx, id, updates)
}

// CompleteScheme 完成巡检
func (s *EAMInspectionService) CompleteScheme(ctx context.Context, id int64, results []model.EAMInspectionResultItemReqVO) error {
	// 计算整体结果
	hasAbnormal := false
	for _, r := range results {
		if !r.IsNormal {
			hasAbnormal = true
			break
		}
	}
	result := "OK"
	if hasAbnormal {
		result = "NG"
	}
	updates := map[string]interface{}{
		"status":    "COMPLETED",
		"result":    result,
	}
	if err := s.schemeRepo.Update(ctx, id, updates); err != nil {
		return err
	}
	// 保存巡检结果明细
	if len(results) > 0 {
		items := make([]model.EAMInspectionResult, len(results))
		for i, r := range results {
			items[i] = model.EAMInspectionResult{
				SchemeID:   id,
				ItemID:     r.ItemID,
				ItemName:   r.ItemName,
				CheckValue: r.CheckValue,
				IsNormal:   r.IsNormal,
				Remark:     r.Remark,
			}
		}
		return s.resultRepo.CreateBatch(ctx, items)
	}
	return nil
}

// SubmitResult 提交巡检结果
func (s *EAMInspectionService) SubmitResult(ctx context.Context, req *model.EAMInspectionResultSubmitReqVO) error {
	return s.CompleteScheme(ctx, req.SchemeID, req.Results)
}

// ========== 辅助函数 ==========

func generateSchemeNo() string {
	return "XJ" + time.Now().Format("20060102150405")
}

func parseTime(s string) (t time.Time) {
	// 简单解析 YYYY-MM-DD
	if len(s) >= 10 {
		parsed, _ := time.Parse("2006-01-02", s[:10])
		return parsed
	}
	return
}