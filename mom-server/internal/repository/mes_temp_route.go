package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// TempRouteRepository 临时工艺路线仓库
type TempRouteRepository struct {
	db *gorm.DB
}

func NewTempRouteRepository(db *gorm.DB) *TempRouteRepository {
	return &TempRouteRepository{db: db}
}

// Create 创建临时工艺路线
func (r *TempRouteRepository) Create(ctx context.Context, tempRoute *model.TempRoute) error {
	return r.db.WithContext(ctx).Create(tempRoute).Error
}

// GetByID 根据ID获取临时工艺路线
func (r *TempRouteRepository) GetByID(ctx context.Context, id uint) (*model.TempRoute, error) {
	var tempRoute model.TempRoute
	if err := r.db.WithContext(ctx).First(&tempRoute, id).Error; err != nil {
		return nil, err
	}
	return &tempRoute, nil
}

// Update 更新临时工艺路线
func (r *TempRouteRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.TempRoute{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除临时工艺路线
func (r *TempRouteRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.TempRoute{}, id).Error
}

// ListByOrderDayID 根据日计划ID查询临时工艺路线列表
func (r *TempRouteRepository) ListByOrderDayID(ctx context.Context, orderDayID int64) ([]model.TempRoute, error) {
	var list []model.TempRoute
	if err := r.db.WithContext(ctx).Where("order_day_id = ?", orderDayID).Order("created_at desc").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// ListByTenantID 查询租户下所有临时工艺路线（分页）
func (r *TempRouteRepository) ListByTenantID(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.TempRoute, int64, error) {
	var list []model.TempRoute
	var total int64

	db := r.db.WithContext(ctx).Model(&model.TempRoute{}).Where("tenant_id = ?", tenantID)

	if orderDayID, ok := query["order_day_id"].(int64); ok && orderDayID > 0 {
		db = db.Where("order_day_id = ?", orderDayID)
	}
	if status, ok := query["status"].(int); ok && status >= 0 {
		db = db.Where("status = ?", status)
	}

	db.Count(&total)

	limit := 20
	page := 1
	if l, ok := query["limit"].(int); ok && l > 0 {
		limit = l
	}
	if p, ok := query["page"].(int); ok && p > 0 {
		page = p
	}

	offset := (page - 1) * limit
	if err := db.Offset(offset).Limit(limit).Order("created_at desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
