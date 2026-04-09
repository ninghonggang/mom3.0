package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type MaterialCategoryRepository struct {
	db *gorm.DB
}

func NewMaterialCategoryRepository(db *gorm.DB) *MaterialCategoryRepository {
	return &MaterialCategoryRepository{db: db}
}

func (r *MaterialCategoryRepository) List(ctx context.Context, tenantID int64) ([]model.MaterialCategory, error) {
	var list []model.MaterialCategory
	query := r.db.WithContext(ctx).Model(&model.MaterialCategory{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	err := query.Order("sort ASC, id ASC").Find(&list).Error
	return list, err
}

func (r *MaterialCategoryRepository) Tree(ctx context.Context, tenantID int64) ([]model.MaterialCategory, error) {
	categories, err := r.List(ctx, tenantID)
	if err != nil {
		return nil, err
	}
	return r.BuildTree(categories), nil
}

func (r *MaterialCategoryRepository) BuildTree(categories []model.MaterialCategory) []model.MaterialCategory {
	var result []model.MaterialCategory
	categoryMap := make(map[int64]*model.MaterialCategory)

	for i := range categories {
		categoryMap[categories[i].ID] = &categories[i]
	}

	for i := range categories {
		if categories[i].ParentID == 0 {
			result = append(result, categories[i])
		} else {
			if parent, ok := categoryMap[categories[i].ParentID]; ok {
				parent.Children = append(parent.Children, categories[i])
			}
		}
	}

	return result
}

func (r *MaterialCategoryRepository) GetByID(ctx context.Context, id uint) (*model.MaterialCategory, error) {
	var category model.MaterialCategory
	err := r.db.WithContext(ctx).First(&category, id).Error
	return &category, err
}

func (r *MaterialCategoryRepository) Create(ctx context.Context, category *model.MaterialCategory) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *MaterialCategoryRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.MaterialCategory{}).Where("id = ?", id).Updates(updates).Error
}

func (r *MaterialCategoryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.MaterialCategory{}, id).Error
}
