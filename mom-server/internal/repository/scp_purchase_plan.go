package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// ScpPurchasePlanRepository 采购计划仓库
type ScpPurchasePlanRepository struct {
	db *gorm.DB
}

func NewScpPurchasePlanRepository(db *gorm.DB) *ScpPurchasePlanRepository {
	return &ScpPurchasePlanRepository{db: db}
}

func (r *ScpPurchasePlanRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.ScpPurchasePlan, int64, error) {
	var list []model.ScpPurchasePlan
	var total int64

	q := r.db.WithContext(ctx).Model(&model.ScpPurchasePlan{}).Where("tenant_id = ?", tenantID)

	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}
	if planType, ok := query["plan_type"]; ok && planType != "" {
		q = q.Where("plan_type = ?", planType)
	}
	if planYear, ok := query["plan_year"]; ok && planYear != nil {
		q = q.Where("plan_year = ?", planYear)
	}
	if planMonth, ok := query["plan_month"]; ok && planMonth != nil {
		q = q.Where("plan_month = ?", planMonth)
	}

	q.Count(&total)
	q = q.Preload("Items").Order("id DESC")

	if page, ok := query["page"].(int); ok && page > 0 {
		limit := 20
		if limitVal, ok := query["limit"].(int); ok && limitVal > 0 {
			limit = limitVal
		}
		q = q.Offset((page - 1) * limit).Limit(limit)
	}

	err := q.Find(&list).Error
	return list, total, err
}

func (r *ScpPurchasePlanRepository) GetByID(ctx context.Context, id uint) (*model.ScpPurchasePlan, error) {
	var plan model.ScpPurchasePlan
	err := r.db.WithContext(ctx).Preload("Items").First(&plan, id).Error
	return &plan, err
}

func (r *ScpPurchasePlanRepository) GetByNo(ctx context.Context, tenantID int64, planNo string) (*model.ScpPurchasePlan, error) {
	var plan model.ScpPurchasePlan
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND plan_no = ?", tenantID, planNo).First(&plan).Error
	return &plan, err
}

func (r *ScpPurchasePlanRepository) Create(ctx context.Context, plan *model.ScpPurchasePlan) error {
	return r.db.WithContext(ctx).Create(plan).Error
}

func (r *ScpPurchasePlanRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ScpPurchasePlan{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ScpPurchasePlanRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		tx.Where("plan_id = ?", id).Delete(&model.ScpPurchasePlanItem{})
		return tx.Delete(&model.ScpPurchasePlan{}, id).Error
	})
}

func (r *ScpPurchasePlanRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).Model(&model.ScpPurchasePlan{}).Where("id = ?", id).Update("status", status).Error
}

// CreateWithItems 创建采购计划并包含明细
func (r *ScpPurchasePlanRepository) CreateWithItems(ctx context.Context, plan *model.ScpPurchasePlan) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(plan).Error; err != nil {
			return err
		}
		return nil
	})
}

// UpdateWithItems 更新采购计划并更新明细
func (r *ScpPurchasePlanRepository) UpdateWithItems(ctx context.Context, id uint, plan *model.ScpPurchasePlan, itemUpdates []map[string]interface{}) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 更新主表
		updates := map[string]interface{}{
			"title":         plan.Title,
			"plan_type":     plan.PlanType,
			"plan_year":     plan.PlanYear,
			"plan_month":    plan.PlanMonth,
			"quarter":       plan.Quarter,
			"currency":      plan.Currency,
			"department":    plan.Department,
			"total_items":   plan.TotalItems,
			"total_amount":  plan.TotalAmount,
			"remark":        plan.Remark,
		}
		if err := tx.Model(&model.ScpPurchasePlan{}).Where("id = ?", id).Updates(updates).Error; err != nil {
			return err
		}
		// 删除旧明细
		if err := tx.Where("plan_id = ?", id).Delete(&model.ScpPurchasePlanItem{}).Error; err != nil {
			return err
		}
		// 创建新明细
		for i := range plan.Items {
			plan.Items[i].PlanID = int64(id)
			if err := tx.Create(&plan.Items[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// GetItems 获取采购计划明细
func (r *ScpPurchasePlanRepository) GetItems(ctx context.Context, planID uint) ([]model.ScpPurchasePlanItem, error) {
	var items []model.ScpPurchasePlanItem
	err := r.db.WithContext(ctx).Where("plan_id = ?", planID).Order("line_no ASC").Find(&items).Error
	return items, err
}

// UpdateItemStatus 更新明细行状态
func (r *ScpPurchasePlanRepository) UpdateItemStatus(ctx context.Context, itemID uint, status string) error {
	return r.db.WithContext(ctx).Model(&model.ScpPurchasePlanItem{}).Where("id = ?", itemID).Update("status", status).Error
}
