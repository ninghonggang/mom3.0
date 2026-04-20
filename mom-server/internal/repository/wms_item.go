package repository

import (
	"context"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

// WMSItemRepository 货品管理仓储层
type WMSItemRepository struct {
	db *gorm.DB
}

// NewWMSItemRepository 创建货品仓储实例
func NewWMSItemRepository(db *gorm.DB) *WMSItemRepository {
	return &WMSItemRepository{db: db}
}

// List 获取货品列表（分页）
func (r *WMSItemRepository) List(ctx context.Context, query *model.WMSItemQueryVO) ([]model.WMSItem, int64, error) {
	var list []model.WMSItem
	var total int64

	queryDB := r.db.WithContext(ctx).Model(&model.WMSItem{})

	// 按租户筛选
	if query != nil && query.Keyword != "" {
		queryDB = queryDB.Where("item_code LIKE ? OR item_name LIKE ? OR barcode LIKE ?",
			"%"+query.Keyword+"%", "%"+query.Keyword+"%", "%"+query.Keyword+"%")
	}
	if query != nil && query.ItemType != "" {
		queryDB = queryDB.Where("item_type = ?", query.ItemType)
	}
	if query != nil && query.Status != "" {
		queryDB = queryDB.Where("status = ?", query.Status)
	}

	err := queryDB.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页
	if query != nil && query.Page > 0 && query.PageSize > 0 {
		offset := (query.Page - 1) * query.PageSize
		queryDB = queryDB.Offset(offset).Limit(query.PageSize)
	}

	err = queryDB.Order("id DESC").Find(&list).Error
	return list, total, err
}

// GetByID 根据ID获取货品
func (r *WMSItemRepository) GetByID(ctx context.Context, id uint) (*model.WMSItem, error) {
	var item model.WMSItem
	err := r.db.WithContext(ctx).First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// GetByCode 根据货品编码获取货品
func (r *WMSItemRepository) GetByCode(ctx context.Context, itemCode string) (*model.WMSItem, error) {
	var item model.WMSItem
	err := r.db.WithContext(ctx).Where("item_code = ?", itemCode).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// Search 搜索货品（关键字匹配）
func (r *WMSItemRepository) Search(ctx context.Context, keyword string) ([]model.WMSItem, error) {
	var list []model.WMSItem
	query := r.db.WithContext(ctx).
		Where("item_code LIKE ? OR item_name LIKE ? OR barcode LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").
		Where("status = ?", "ACTIVE").
		Order("id DESC").
		Limit(50) // 限制返回数量

	err := query.Find(&list).Error
	return list, err
}

// Create 创建货品
func (r *WMSItemRepository) Create(ctx context.Context, item *model.WMSItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

// Update 更新货品
func (r *WMSItemRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.WMSItem{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除货品
func (r *WMSItemRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.WMSItem{}, id).Error
}
