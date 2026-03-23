package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type BOMRepository struct {
	db *gorm.DB
}

func NewBOMRepository(db *gorm.DB) *BOMRepository {
	return &BOMRepository{db: db}
}

func (r *BOMRepository) List(ctx context.Context, tenantID int64) ([]model.MdmBOM, int64, error) {
	var list []model.MdmBOM
	var total int64

	query := r.db.WithContext(ctx).Model(&model.MdmBOM{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *BOMRepository) GetByID(ctx context.Context, id uint) (*model.MdmBOM, error) {
	var bom model.MdmBOM
	err := r.db.WithContext(ctx).First(&bom, id).Error
	return &bom, err
}

func (r *BOMRepository) Create(ctx context.Context, bom *model.MdmBOM) error {
	return r.db.WithContext(ctx).Create(bom).Error
}

func (r *BOMRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.MdmBOM{}).Where("id = ?", id).Updates(updates).Error
}

func (r *BOMRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.MdmBOM{}, id).Error
}

type BOMItemRepository struct {
	db *gorm.DB
}

func NewBOMItemRepository(db *gorm.DB) *BOMItemRepository {
	return &BOMItemRepository{db: db}
}

func (r *BOMItemRepository) ListByBOMID(ctx context.Context, bomID uint) ([]model.MdmBOMItem, error) {
	var items []model.MdmBOMItem
	err := r.db.WithContext(ctx).Where("bom_id = ?", bomID).Order("line_no ASC").Find(&items).Error
	return items, err
}

func (r *BOMItemRepository) Create(ctx context.Context, item *model.MdmBOMItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *BOMItemRepository) DeleteByBOMID(ctx context.Context, bomID uint) error {
	return r.db.WithContext(ctx).Where("bom_id = ?", bomID).Delete(&model.MdmBOMItem{}).Error
}
