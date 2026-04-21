package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// WmsAreaRepository 库区仓储层
type WmsAreaRepository struct {
	db *gorm.DB
}

// NewWmsAreaRepository 创建库区仓储实例
func NewWmsAreaRepository(db *gorm.DB) *WmsAreaRepository {
	return &WmsAreaRepository{db: db}
}

// Create 创建库区
func (r *WmsAreaRepository) Create(ctx context.Context, area *model.WmsArea) error {
	return r.db.WithContext(ctx).Create(area).Error
}

// Update 更新库区
func (r *WmsAreaRepository) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.WmsArea{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除库区
func (r *WmsAreaRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.WmsArea{}, id).Error
}

// GetByID 根据ID获取库区
func (r *WmsAreaRepository) GetByID(ctx context.Context, id uint64) (*model.WmsArea, error) {
	var area model.WmsArea
	err := r.db.WithContext(ctx).First(&area, id).Error
	if err != nil {
		return nil, err
	}
	return &area, nil
}

// GetByCode 根据编码获取库区
func (r *WmsAreaRepository) GetByCode(ctx context.Context, tenantID int64, areaCode string) (*model.WmsArea, error) {
	var area model.WmsArea
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND area_code = ?", tenantID, areaCode).First(&area).Error
	if err != nil {
		return nil, err
	}
	return &area, nil
}

// Page 分页查询库区
func (r *WmsAreaRepository) Page(ctx context.Context, tenantID int64, query *model.WmsAreaQueryVO) ([]model.WmsArea, int64, error) {
	var list []model.WmsArea
	var total int64

	queryDB := r.db.WithContext(ctx).Model(&model.WmsArea{}).Where("tenant_id = ?", tenantID)

	if query != nil && query.Keyword != "" {
		queryDB = queryDB.Where("area_code LIKE ? OR area_name LIKE ? OR warehouse_code LIKE ?",
			"%"+query.Keyword+"%", "%"+query.Keyword+"%", "%"+query.Keyword+"%")
	}
	if query != nil && query.WarehouseCode != "" {
		queryDB = queryDB.Where("warehouse_code = ?", query.WarehouseCode)
	}
	if query != nil && query.AreaType != "" {
		queryDB = queryDB.Where("area_type = ?", query.AreaType)
	}
	if query != nil && query.Level > 0 {
		queryDB = queryDB.Where("level = ?", query.Level)
	}
	if query != nil && query.Status != "" {
		queryDB = queryDB.Where("status = ?", query.Status)
	}

	err := queryDB.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if query != nil && query.Page > 0 && query.PageSize > 0 {
		offset := (query.Page - 1) * query.PageSize
		queryDB = queryDB.Offset(offset).Limit(query.PageSize)
	}

	err = queryDB.Order("level ASC, area_code ASC").Find(&list).Error
	return list, total, err
}

// ListByWarehouse 按仓库获取库区列表
func (r *WmsAreaRepository) ListByWarehouse(ctx context.Context, tenantID int64, warehouseCode string) ([]model.WmsArea, error) {
	var list []model.WmsArea
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND warehouse_code = ?", tenantID, warehouseCode).
		Order("level ASC, area_code ASC").Find(&list).Error
	return list, err
}

// ListAll 获取所有库区
func (r *WmsAreaRepository) ListAll(ctx context.Context, tenantID int64) ([]model.WmsArea, error) {
	var list []model.WmsArea
	err := r.db.WithContext(ctx).Where("tenant_id = ?", tenantID).
		Order("level ASC, area_code ASC").Find(&list).Error
	return list, err
}
