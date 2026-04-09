package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type MaterialCategoryService struct {
	repo *repository.MaterialCategoryRepository
}

func NewMaterialCategoryService(repo *repository.MaterialCategoryRepository) *MaterialCategoryService {
	return &MaterialCategoryService{repo: repo}
}

func (s *MaterialCategoryService) List(ctx context.Context, tenantID int64) ([]model.MaterialCategory, error) {
	return s.repo.List(ctx, tenantID)
}

func (s *MaterialCategoryService) Tree(ctx context.Context, tenantID int64) ([]model.MaterialCategory, error) {
	return s.repo.Tree(ctx, tenantID)
}

func (s *MaterialCategoryService) GetByID(ctx context.Context, id string) (*model.MaterialCategory, error) {
	var categoryID uint
	_, err := fmt.Sscanf(id, "%d", &categoryID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, categoryID)
}

func (s *MaterialCategoryService) Create(ctx context.Context, category *model.MaterialCategory) error {
	return s.repo.Create(ctx, category)
}

func (s *MaterialCategoryService) Update(ctx context.Context, id string, category *model.MaterialCategory) error {
	var categoryID uint
	_, err := fmt.Sscanf(id, "%d", &categoryID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, categoryID, map[string]interface{}{
		"parent_id":     category.ParentID,
		"category_name": category.CategoryName,
		"category_code": category.CategoryCode,
		"sort":          category.Sort,
		"status":        category.Status,
	})
}

func (s *MaterialCategoryService) Delete(ctx context.Context, id string) error {
	var categoryID uint
	_, err := fmt.Sscanf(id, "%d", &categoryID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, categoryID)
}
