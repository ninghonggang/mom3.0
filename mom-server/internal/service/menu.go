package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type MenuService struct {
	repo *repository.MenuRepository
}

func NewMenuService(repo *repository.MenuRepository) *MenuService {
	return &MenuService{repo: repo}
}

func (s *MenuService) List(ctx context.Context, tenantID int64) ([]model.Menu, error) {
	return s.repo.List(ctx, tenantID)
}

func (s *MenuService) Tree(ctx context.Context, tenantID int64) ([]model.Menu, error) {
	return s.repo.Tree(ctx, tenantID)
}

func (s *MenuService) GetByID(ctx context.Context, id string) (*model.Menu, error) {
	var menuID uint
	_, err := fmt.Sscanf(id, "%d", &menuID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, menuID)
}

func (s *MenuService) Create(ctx context.Context, menu *model.Menu) error {
	return s.repo.Create(ctx, menu)
}

func (s *MenuService) Update(ctx context.Context, id string, menu *model.Menu) error {
	var menuID uint
	_, err := fmt.Sscanf(id, "%d", &menuID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, menuID, map[string]interface{}{
		"parent_id":  menu.ParentID,
		"menu_name":  menu.MenuName,
		"menu_type":  menu.MenuType,
		"path":       menu.Path,
		"component":  menu.Component,
		"perms":      menu.Perms,
		"icon":       menu.Icon,
		"sort":       menu.Sort,
		"visible":    menu.Visible,
		"status":     menu.Status,
	})
}

func (s *MenuService) Delete(ctx context.Context, id string) error {
	var menuID uint
	_, err := fmt.Sscanf(id, "%d", &menuID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, menuID)
}
