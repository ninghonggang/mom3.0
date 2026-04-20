package repository

import (
	"context"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

// WmsStrategyRepository 策略配置仓储层
type WmsStrategyRepository struct {
	db *gorm.DB
}

// NewWmsStrategyRepository 创建策略配置仓储实例
func NewWmsStrategyRepository(db *gorm.DB) *WmsStrategyRepository {
	return &WmsStrategyRepository{db: db}
}

// List 获取策略配置列表（分页）
func (r *WmsStrategyRepository) List(ctx context.Context, query *model.WmsStrategyQueryVO) ([]model.WmsStrategy, int64, error) {
	var list []model.WmsStrategy
	var total int64

	queryDB := r.db.WithContext(ctx).Model(&model.WmsStrategy{})

	if query != nil && query.Keyword != "" {
		queryDB = queryDB.Where("strategy_code LIKE ? OR strategy_name LIKE ?",
			"%"+query.Keyword+"%", "%"+query.Keyword+"%")
	}
	if query != nil && query.StrategyType != "" {
		queryDB = queryDB.Where("strategy_type = ?", query.StrategyType)
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

	err = queryDB.Order("priority DESC, id DESC").Find(&list).Error
	return list, total, err
}

// GetByID 根据ID获取策略配置
func (r *WmsStrategyRepository) GetByID(ctx context.Context, id uint64) (*model.WmsStrategy, error) {
	var strategy model.WmsStrategy
	err := r.db.WithContext(ctx).First(&strategy, id).Error
	if err != nil {
		return nil, err
	}
	return &strategy, nil
}

// GetByCode 根据编码获取策略配置
func (r *WmsStrategyRepository) GetByCode(ctx context.Context, strategyCode string) (*model.WmsStrategy, error) {
	var strategy model.WmsStrategy
	err := r.db.WithContext(ctx).Where("strategy_code = ?", strategyCode).First(&strategy).Error
	if err != nil {
		return nil, err
	}
	return &strategy, nil
}

// Create 创建策略配置
func (r *WmsStrategyRepository) Create(ctx context.Context, strategy *model.WmsStrategy) error {
	return r.db.WithContext(ctx).Create(strategy).Error
}

// Update 更新策略配置
func (r *WmsStrategyRepository) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.WmsStrategy{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除策略配置
func (r *WmsStrategyRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.WmsStrategy{}, id).Error
}
