package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type ProductionOfflineRepository struct {
	db *gorm.DB
}

func NewProductionOfflineRepository(db *gorm.DB) *ProductionOfflineRepository {
	return &ProductionOfflineRepository{db: db}
}

func (r *ProductionOfflineRepository) List(c *gorm.DB, tenantID uint64, query map[string]interface{}) ([]model.ProductionOffline, int64, error) {
	var list []model.ProductionOffline
	var total int64

	db := c
	if db == nil {
		db = r.db
	}
	q := db.Model(&model.ProductionOffline{}).Where("tenant_id = ?", tenantID)

	if v, ok := query["work_order_code"]; ok && v.(string) != "" {
		q = q.Where("work_order_code LIKE ?", "%"+v.(string)+"%")
	}
	if v, ok := query["product_code"]; ok && v.(string) != "" {
		q = q.Where("product_code LIKE ?", "%"+v.(string)+"%")
	}
	if v, ok := query["product_name"]; ok && v.(string) != "" {
		q = q.Where("product_name LIKE ?", "%"+v.(string)+"%")
	}
	if v, ok := query["offline_type"]; ok && v.(string) != "" {
		q = q.Where("offline_type = ?", v.(string))
	}
	if v, ok := query["handle_method"]; ok && v.(string) != "" {
		q = q.Where("handle_method = ?", v.(string))
	}
	if v, ok := query["status"]; ok && v.(string) != "" {
		q = q.Where("status = ?", v.(string))
	}
	if v, ok := query["start_date"]; ok && v.(string) != "" {
		q = q.Where("created_at >= ?", v.(string))
	}
	if v, ok := query["end_date"]; ok && v.(string) != "" {
		q = q.Where("created_at <= ?", v.(string))
	}

	err := q.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = q.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *ProductionOfflineRepository) GetByID(id uint64) (*model.ProductionOffline, error) {
	var offline model.ProductionOffline
	err := r.db.First(&offline, id).Error
	return &offline, err
}

func (r *ProductionOfflineRepository) Create(d *model.ProductionOffline) error {
	return r.db.Create(d).Error
}

func (r *ProductionOfflineRepository) Update(d *model.ProductionOffline) error {
	return r.db.Save(d).Error
}

func (r *ProductionOfflineRepository) Delete(id uint64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 删除明细
		if err := tx.Where("offline_id = ?", id).Delete(&model.ProductionOfflineItem{}).Error; err != nil {
			return err
		}
		// 删除主表
		return tx.Delete(&model.ProductionOffline{}, id).Error
	})
}

func (r *ProductionOfflineRepository) CreateWithItems(d *model.ProductionOffline, items []model.ProductionOfflineItem) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(d).Error; err != nil {
			return err
		}
		if len(items) > 0 {
			for i := range items {
				items[i].OfflineID = d.ID
				items[i].TenantID = d.TenantID
			}
			if err := tx.Create(&items).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *ProductionOfflineRepository) GetItemsByOfflineID(offlineID uint64) ([]model.ProductionOfflineItem, error) {
	var items []model.ProductionOfflineItem
	err := r.db.Where("offline_id = ?", offlineID).Order("id ASC").Find(&items).Error
	return items, err
}

func (r *ProductionOfflineRepository) UpdateItems(offlineID uint64, items []model.ProductionOfflineItem) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 删除旧明细
		if err := tx.Where("offline_id = ?", offlineID).Delete(&model.ProductionOfflineItem{}).Error; err != nil {
			return err
		}
		// 创建新明细
		if len(items) > 0 {
			if err := tx.Create(&items).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *ProductionOfflineRepository) UpdateWithContext(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ProductionOffline{}).Where("id = ?", id).Updates(updates).Error
}
