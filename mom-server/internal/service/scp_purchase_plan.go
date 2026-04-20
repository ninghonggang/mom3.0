package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type ScpPurchasePlanService struct {
	planRepo *repository.ScpPurchasePlanRepository
}

func NewScpPurchasePlanService(planRepo *repository.ScpPurchasePlanRepository) *ScpPurchasePlanService {
	return &ScpPurchasePlanService{planRepo: planRepo}
}

// ListPurchasePlans 查询采购计划列表
func (s *ScpPurchasePlanService) ListPurchasePlans(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.ScpPurchasePlan, int64, error) {
	return s.planRepo.List(ctx, tenantID, query)
}

// GetPurchasePlan 获取采购计划详情
func (s *ScpPurchasePlanService) GetPurchasePlan(ctx context.Context, id string) (*model.ScpPurchasePlan, error) {
	var planID uint
	_, err := fmt.Sscanf(id, "%d", &planID)
	if err != nil {
		return nil, err
	}
	return s.planRepo.GetByID(ctx, planID)
}

// CreatePurchasePlan 创建采购计划
func (s *ScpPurchasePlanService) CreatePurchasePlan(ctx context.Context, tenantID int64, req *model.ScpPurchasePlanCreateReqVO) (*model.ScpPurchasePlan, error) {
	// 生成计划编号
	planNo := generatePurchasePlanNo(tenantID)

	plan := &model.ScpPurchasePlan{
		TenantID:   tenantID,
		PlanNo:     planNo,
		Title:      req.Title,
		PlanType:   req.PlanType,
		PlanYear:   req.PlanYear,
		PlanMonth:  req.PlanMonth,
		Quarter:    req.Quarter,
		Currency:   req.Currency,
		Department: req.Department,
		Remark:     req.Remark,
		Status:     "DRAFT",
	}

	if plan.Currency == "" {
		plan.Currency = "CNY"
	}

	// 处理明细
	var totalAmount float64
	for i, itemReq := range req.Items {
		item := model.ScpPurchasePlanItem{
			PlanNo:         planNo,
			LineNo:         i + 1,
			MaterialID:      int64ToPtr(itemReq.MaterialID),
			MaterialCode:   itemReq.MaterialCode,
			MaterialName:   itemReq.MaterialName,
			Spec:           itemReq.Spec,
			Unit:           itemReq.Unit,
			ReqQty:         itemReq.ReqQty,
			UnitPrice:      itemReq.UnitPrice,
			LineAmount:     itemReq.ReqQty * itemReq.UnitPrice,
			SupplierID:      int64ToPtr(itemReq.SupplierID),
			SupplierCode:   itemReq.SupplierCode,
			SupplierName:   itemReq.SupplierName,
			MrsNo:          itemReq.MrsNo,
			MrsLineNo:      itemReq.MrsLineNo,
			Status:         "PENDING",
		}
		if itemReq.ReqDeliveryDate != "" {
			t, _ := time.Parse("2006-01-02", itemReq.ReqDeliveryDate)
			item.ReqDeliveryDate = &t
		}
		if itemReq.PromiseDate != "" {
			t, _ := time.Parse("2006-01-02", itemReq.PromiseDate)
			item.PromiseDate = &t
		}
		plan.Items = append(plan.Items, item)
		plan.TotalItems = i + 1
		totalAmount += item.LineAmount
	}
	plan.TotalAmount = totalAmount

	if err := s.planRepo.CreateWithItems(ctx, plan); err != nil {
		return nil, err
	}

	return plan, nil
}

// UpdatePurchasePlan 更新采购计划
func (s *ScpPurchasePlanService) UpdatePurchasePlan(ctx context.Context, tenantID int64, id string, req *model.ScpPurchasePlanUpdateReqVO) (*model.ScpPurchasePlan, error) {
	var planID uint
	_, err := fmt.Sscanf(id, "%d", &planID)
	if err != nil {
		return nil, err
	}

	// 获取现有计划
	existing, err := s.planRepo.GetByID(ctx, planID)
	if err != nil {
		return nil, err
	}

	if existing.Status != "DRAFT" {
		return nil, fmt.Errorf("只有草稿状态的采购计划可以更新")
	}

	plan := &model.ScpPurchasePlan{
		Title:      req.Title,
		PlanType:   req.PlanType,
		PlanYear:   req.PlanYear,
		PlanMonth:  req.PlanMonth,
		Quarter:    req.Quarter,
		Currency:   req.Currency,
		Department: req.Department,
		Remark:     req.Remark,
	}

	// 处理明细
	var totalAmount float64
	for i, itemReq := range req.Items {
		item := model.ScpPurchasePlanItem{
			PlanID:         int64(planID),
			PlanNo:         existing.PlanNo,
			LineNo:         i + 1,
			MaterialID:      int64ToPtr(itemReq.MaterialID),
			MaterialCode:   itemReq.MaterialCode,
			MaterialName:   itemReq.MaterialName,
			Spec:           itemReq.Spec,
			Unit:           itemReq.Unit,
			ReqQty:         itemReq.ReqQty,
			UnitPrice:      itemReq.UnitPrice,
			LineAmount:     itemReq.ReqQty * itemReq.UnitPrice,
			SupplierID:      int64ToPtr(itemReq.SupplierID),
			SupplierCode:   itemReq.SupplierCode,
			SupplierName:   itemReq.SupplierName,
			MrsNo:          itemReq.MrsNo,
			MrsLineNo:      itemReq.MrsLineNo,
			Status:         "PENDING",
		}
		if itemReq.ReqDeliveryDate != "" {
			t, _ := time.Parse("2006-01-02", itemReq.ReqDeliveryDate)
			item.ReqDeliveryDate = &t
		}
		if itemReq.PromiseDate != "" {
			t, _ := time.Parse("2006-01-02", itemReq.PromiseDate)
			item.PromiseDate = &t
		}
		plan.Items = append(plan.Items, item)
		plan.TotalItems = i + 1
		totalAmount += item.LineAmount
	}
	plan.TotalAmount = totalAmount

	if err := s.planRepo.UpdateWithItems(ctx, planID, plan, nil); err != nil {
		return nil, err
	}

	return s.planRepo.GetByID(ctx, planID)
}

// DeletePurchasePlan 删除采购计划
func (s *ScpPurchasePlanService) DeletePurchasePlan(ctx context.Context, id string) error {
	var planID uint
	_, err := fmt.Sscanf(id, "%d", &planID)
	if err != nil {
		return err
	}

	// 获取现有计划
	existing, err := s.planRepo.GetByID(ctx, planID)
	if err != nil {
		return err
	}

	if existing.Status != "DRAFT" {
		return fmt.Errorf("只有草稿状态的采购计划可以删除")
	}

	return s.planRepo.Delete(ctx, planID)
}

// ConfirmPurchasePlan 确认采购计划
func (s *ScpPurchasePlanService) ConfirmPurchasePlan(ctx context.Context, id string, userID int64) error {
	var planID uint
	_, err := fmt.Sscanf(id, "%d", &planID)
	if err != nil {
		return err
	}

	existing, err := s.planRepo.GetByID(ctx, planID)
	if err != nil {
		return err
	}

	if existing.Status != "DRAFT" {
		return fmt.Errorf("只有草稿状态的采购计划可以确认")
	}

	now := time.Now()
	return s.planRepo.Update(ctx, planID, map[string]interface{}{
		"status":       "CONFIRMED",
		"confirmed_by": userID,
		"confirmed_at": now,
	})
}

// PublishPurchasePlan 发布采购计划
func (s *ScpPurchasePlanService) PublishPurchasePlan(ctx context.Context, id string, userID int64) error {
	var planID uint
	_, err := fmt.Sscanf(id, "%d", &planID)
	if err != nil {
		return err
	}

	existing, err := s.planRepo.GetByID(ctx, planID)
	if err != nil {
		return err
	}

	if existing.Status != "CONFIRMED" {
		return fmt.Errorf("只有已确认状态的采购计划可以发布")
	}

	now := time.Now()
	return s.planRepo.Update(ctx, planID, map[string]interface{}{
		"status":       "PUBLISHED",
		"published_by": userID,
		"published_at": now,
	})
}

// ClosePurchasePlan 关闭采购计划
func (s *ScpPurchasePlanService) ClosePurchasePlan(ctx context.Context, id string, userID int64, closeReason string) error {
	var planID uint
	_, err := fmt.Sscanf(id, "%d", &planID)
	if err != nil {
		return err
	}

	existing, err := s.planRepo.GetByID(ctx, planID)
	if err != nil {
		return err
	}

	if existing.Status != "PUBLISHED" {
		return fmt.Errorf("只有已发布状态的采购计划可以关闭")
	}

	now := time.Now()
	return s.planRepo.Update(ctx, planID, map[string]interface{}{
		"status":       "CLOSED",
		"closed_by":    userID,
		"closed_at":    now,
		"close_reason": closeReason,
	})
}

// GetPurchasePlanItems 获取采购计划明细
func (s *ScpPurchasePlanService) GetPurchasePlanItems(ctx context.Context, planID string) ([]model.ScpPurchasePlanItem, error) {
	var id uint
	_, err := fmt.Sscanf(planID, "%d", &id)
	if err != nil {
		return nil, err
	}
	return s.planRepo.GetItems(ctx, id)
}

// Helper function
func int64ToPtr(i int64) *int64 {
	return &i
}

// generatePurchasePlanNo 生成采购计划编号
func generatePurchasePlanNo(tenantID int64) string {
	return fmt.Sprintf("PP-%s-%d", time.Now().Format("20060102"), tenantID)
}
