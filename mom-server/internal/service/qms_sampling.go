package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// QMSSamplingService 抽样方案服务
type QMSSamplingService struct {
	planRepo   *repository.QMSSamplingPlanRepository
	ruleRepo   *repository.QMSSamplingRuleRepository
	recordRepo *repository.QMSSamplingRecordRepository
}

func NewQMSSamplingService(
	planRepo *repository.QMSSamplingPlanRepository,
	ruleRepo *repository.QMSSamplingRuleRepository,
	recordRepo *repository.QMSSamplingRecordRepository,
) *QMSSamplingService {
	return &QMSSamplingService{
		planRepo:   planRepo,
		ruleRepo:   ruleRepo,
		recordRepo: recordRepo,
	}
}

// ListPlan 查询抽样方案列表
func (s *QMSSamplingService) ListPlan(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.QMSSamplingPlan, int64, error) {
	return s.planRepo.List(ctx, tenantID, query)
}

// GetPlan 获取抽样方案详情
func (s *QMSSamplingService) GetPlan(ctx context.Context, id int64) (*model.QMSSamplingPlanDetailRespVO, error) {
	plan, err := s.planRepo.GetByID(ctx, uint(id))
	if err != nil {
		return nil, err
	}
	rules, _ := s.ruleRepo.ListByPlanID(ctx, id)
	return &model.QMSSamplingPlanDetailRespVO{
		QMSSamplingPlan: *plan,
		Rules:          rules,
	}, nil
}

// CreatePlan 创建抽样方案
func (s *QMSSamplingService) CreatePlan(ctx context.Context, tenantID int64, req *model.QMSSamplingPlanCreateReqVO, username string) (*model.QMSSamplingPlan, error) {
	// 检查编码唯一性
	existing, _ := s.planRepo.GetByPlanCode(ctx, tenantID, req.PlanCode)
	if existing != nil && existing.ID > 0 {
		return nil, fmt.Errorf("方案编码已存在")
	}

	plan := &model.QMSSamplingPlan{
		TenantID:        tenantID,
		PlanCode:       req.PlanCode,
		PlanName:       req.PlanName,
		InspectionLevel: req.InspectionLevel,
		SampleType:     req.SampleType,
		AQL:            req.AQL,
		Status:         "ACTIVE",
		Remark:         req.Remark,
	}

	if err := s.planRepo.Create(ctx, plan); err != nil {
		return nil, fmt.Errorf("创建抽样方案失败: %w", err)
	}

	// 创建规则
	if len(req.Rules) > 0 {
		var rules []model.QMSSamplingRule
		for _, r := range req.Rules {
			rules = append(rules, model.QMSSamplingRule{
				PlanID:       int64(plan.ID),
				BatchQtyFrom: r.BatchQtyFrom,
				BatchQtyTo:   r.BatchQtyTo,
				SampleSize:   r.SampleSize,
				AcAccept:     r.AcAccept,
				ReReject:     r.ReReject,
			})
		}
		s.ruleRepo.CreateBatch(ctx, rules)
	}

	return plan, nil
}

// UpdatePlan 更新抽样方案
func (s *QMSSamplingService) UpdatePlan(ctx context.Context, id int64, req *model.QMSSamplingPlanUpdateReqVO) error {
	updates := map[string]interface{}{}
	if req.PlanName != "" {
		updates["plan_name"] = req.PlanName
	}
	if req.InspectionLevel != "" {
		updates["inspection_level"] = req.InspectionLevel
	}
	if req.SampleType != "" {
		updates["sample_type"] = req.SampleType
	}
	if req.AQL > 0 {
		updates["aql"] = req.AQL
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	return s.planRepo.Update(ctx, uint(id), updates)
}

// DeletePlan 删除抽样方案
func (s *QMSSamplingService) DeletePlan(ctx context.Context, id int64) error {
	// 先删除规则
	s.ruleRepo.DeleteByPlanID(ctx, id)
	return s.planRepo.Delete(ctx, uint(id))
}

// UpdateRules 更新抽样规则
func (s *QMSSamplingService) UpdateRules(ctx context.Context, planID int64, req *model.QMSSamplingRulesUpdateReqVO) error {
	// 删除旧规则
	s.ruleRepo.DeleteByPlanID(ctx, planID)
	// 创建新规则
	if len(req.Rules) > 0 {
		var rules []model.QMSSamplingRule
		for _, r := range req.Rules {
			rules = append(rules, model.QMSSamplingRule{
				PlanID:       planID,
				BatchQtyFrom: r.BatchQtyFrom,
				BatchQtyTo:   r.BatchQtyTo,
				SampleSize:   r.SampleSize,
				AcAccept:     r.AcAccept,
				ReReject:     r.ReReject,
			})
		}
		return s.ruleRepo.CreateBatch(ctx, rules)
	}
	return nil
}

// Calculate 计算样本数
func (s *QMSSamplingService) Calculate(ctx context.Context, req *model.QMSSamplingCalculateReqVO) (*model.QMSSamplingCalculateRespVO, error) {
	rule, err := s.ruleRepo.GetRuleForBatchQty(ctx, req.PlanID, req.BatchQty)
	if err != nil {
		return nil, fmt.Errorf("未找到对应的抽样规则: %w", err)
	}
	return &model.QMSSamplingCalculateRespVO{
		SampleSize: rule.SampleSize,
		AcAccept:   rule.AcAccept,
		ReReject:   rule.ReReject,
	}, nil
}

// CreateRecord 创建抽样记录
func (s *QMSSamplingService) CreateRecord(ctx context.Context, tenantID int64, req *model.QMSSamplingRecordCreateReqVO) error {
	now := time.Now()
	record := &model.QMSSamplingRecord{
		TenantID:     tenantID,
		PlanID:      req.PlanID,
		PlanCode:    req.PlanCode,
		InspectionID: req.InspectionID,
		BatchQty:    req.BatchQty,
		SampleSize:  req.SampleSize,
		DefectCount: req.DefectCount,
		AcResult:    req.AcResult,
		Inspector:   req.Inspector,
		InspectTime: &now,
	}
	return s.recordRepo.Create(ctx, record)
}

// ListRecord 查询抽样记录
func (s *QMSSamplingService) ListRecord(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.QMSSamplingRecord, int64, error) {
	return s.recordRepo.List(ctx, tenantID, query)
}
