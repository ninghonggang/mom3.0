package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type MenuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

func (r *MenuRepository) List(ctx context.Context, tenantID int64) ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Order("sort ASC").Find(&menus).Error
	return menus, err
}

func (r *MenuRepository) Tree(ctx context.Context, tenantID int64) ([]model.Menu, error) {
	menus, err := r.List(ctx, tenantID)
	if err != nil {
		return nil, err
	}
	return r.BuildTree(menus), nil
}

func (r *MenuRepository) BuildTree(menus []model.Menu) []model.Menu {
	var result []model.Menu
	menuMap := make(map[int64]*model.Menu)

	for i := range menus {
		menuMap[menus[i].ID] = &menus[i]
	}

	for i := range menus {
		if menus[i].ParentID == 0 {
			result = append(result, menus[i])
		} else {
			if parent, ok := menuMap[menus[i].ParentID]; ok {
				parent.Children = append(parent.Children, menus[i])
			}
		}
	}

	return result
}

func (r *MenuRepository) GetByID(ctx context.Context, id uint) (*model.Menu, error) {
	var menu model.Menu
	err := r.db.WithContext(ctx).First(&menu, id).Error
	return &menu, err
}

func (r *MenuRepository) Create(ctx context.Context, menu *model.Menu) error {
	return r.db.WithContext(ctx).Create(menu).Error
}

func (r *MenuRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Menu{}).Where("id = ?", id).Updates(updates).Error
}

func (r *MenuRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Menu{}, id).Error
}

func (r *MenuRepository) GetByIDs(ctx context.Context, ids []uint) ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&menus).Error
	return menus, err
}

type RoleMenuRepository struct {
	db *gorm.DB
}

func NewRoleMenuRepository(db *gorm.DB) *RoleMenuRepository {
	return &RoleMenuRepository{db: db}
}

func (r *RoleMenuRepository) GetMenuIDsByRoleID(ctx context.Context, roleID string) ([]uint, error) {
	var menuIDs []uint
	id := roleID
	err := r.db.WithContext(ctx).Model(&model.RoleMenu{}).
		Where("role_id = ?", id).
		Pluck("menu_id", &menuIDs).Error
	return menuIDs, err
}

func (r *RoleMenuRepository) AssignMenus(ctx context.Context, roleID int64, menuIDs []int64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id = ?", roleID).Delete(&model.RoleMenu{}).Error; err != nil {
			return err
		}
		for _, menuID := range menuIDs {
			if err := tx.Create(&model.RoleMenu{RoleID: roleID, MenuID: menuID}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
