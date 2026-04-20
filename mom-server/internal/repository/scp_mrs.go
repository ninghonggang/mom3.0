package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// ScpMRSRepository MRS仓库
type ScpMRSRepository struct {
	db *gorm.DB
}

func NewScpMRSRepository(db *gorm.DB) *ScpMRSRepository {
	return &ScpMRSRepository{db: db}
}

func (r *ScpMRSRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.ScpMRS, int64, error) {
	var list []model.ScpMRS
	var total int64

	q := r.db.WithContext(ctx).Model(&model.ScpMRS{}).Where("tenant_id = ?", tenantID)

	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}
	if planMonth, ok := query["plan_month"]; ok && planMonth != "" {
		q = q.Where("plan_month = ?", planMonth)
	}
	if sourceType, ok := query["source_type"]; ok && sourceType != "" {
		q = q.Where("source_type = ?", sourceType)
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

func (r *ScpMRSRepository) GetByID(ctx context.Context, id uint) (*model.ScpMRS, error) {
	var mrs model.ScpMRS
	err := r.db.WithContext(ctx).Preload("Items").First(&mrs, id).Error
	return &mrs, err
}

func (r *ScpMRSRepository) GetByNo(ctx context.Context, tenantID int64, mrsNo string) (*model.ScpMRS, error) {
	var mrs model.ScpMRS
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND mrs_no = ?", tenantID, mrsNo).First(&mrs).Error
	return &mrs, err
}

func (r *ScpMRSRepository) Create(ctx context.Context, mrs *model.ScpMRS) error {
	return r.db.WithContext(ctx).Create(mrs).Error
}

func (r *ScpMRSRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ScpMRS{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ScpMRSRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		tx.Where("mrs_id = ?", id).Delete(&model.ScpMRSItem{})
		return tx.Delete(&model.ScpMRS{}, id).Error
	})
}

func (r *ScpMRSRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).Model(&model.ScpMRS{}).Where("id = ?", id).Update("status", status).Error
}

// CreateWithItems 创建MRS并包含明细
func (r *ScpMRSRepository) CreateWithItems(ctx context.Context, mrs *model.ScpMRS) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(mrs).Error; err != nil {
			return err
		}
		return nil
	})
}

// UpdateWithItems 更新MRS并更新明细
func (r *ScpMRSRepository) UpdateWithItems(ctx context.Context, id uint, mrs *model.ScpMRS, itemUpdates []map[string]interface{}) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 更新主表
		updates := map[string]interface{}{
			"plan_month":  mrs.PlanMonth,
			"source_type": mrs.SourceType,
			"source_no":   mrs.SourceNo,
			"remark":      mrs.Remark,
			"total_items": mrs.TotalItems,
			"total_qty":   mrs.TotalQty,
		}
		if err := tx.Model(&model.ScpMRS{}).Where("id = ?", id).Updates(updates).Error; err != nil {
			return err
		}
		// 删除旧明细
		if err := tx.Where("mrs_id = ?", id).Delete(&model.ScpMRSItem{}).Error; err != nil {
			return err
		}
		// 创建新明细
		for _, item := range mrs.Items {
			item.MrsID = int64(id)
			if err := tx.Create(&item).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// GetItems 获取MRS明细
func (r *ScpMRSRepository) GetItems(ctx context.Context, mrsID uint) ([]model.ScpMRSItem, error) {
	var items []model.ScpMRSItem
	err := r.db.WithContext(ctx).Where("mrs_id = ?", mrsID).Order("id ASC").Find(&items).Error
	return items, err
}
