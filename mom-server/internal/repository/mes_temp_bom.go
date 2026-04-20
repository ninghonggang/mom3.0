package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// TempBOMRepository 临时替代BOM仓库
type TempBOMRepository struct {
	db *gorm.DB
}

func NewTempBOMRepository(db *gorm.DB) *TempBOMRepository {
	return &TempBOMRepository{db: db}
}

// Create 创建临时替代BOM
func (r *TempBOMRepository) Create(ctx context.Context, tempBOM *model.TempBOM) error {
	return r.db.WithContext(ctx).Create(tempBOM).Error
}

// GetByID 根据ID获取临时替代BOM
func (r *TempBOMRepository) GetByID(ctx context.Context, id uint) (*model.TempBOM, error) {
	var tempBOM model.TempBOM
	if err := r.db.WithContext(ctx).First(&tempBOM, id).Error; err != nil {
		return nil, err
	}
	return &tempBOM, nil
}

// Update 更新临时替代BOM
func (r *TempBOMRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.TempBOM{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除临时替代BOM
func (r *TempBOMRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.TempBOM{}, id).Error
}

// ListByOrderDayItemID 根据日计划明细项ID查询临时替代BOM列表
func (r *TempBOMRepository) ListByOrderDayItemID(ctx context.Context, orderDayItemID int64) ([]model.TempBOM, error) {
	var list []model.TempBOM
	if err := r.db.WithContext(ctx).Where("order_day_item_id = ?", orderDayItemID).Order("created_at desc").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
