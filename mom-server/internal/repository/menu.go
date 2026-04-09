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
	var result []*model.Menu
	menuMap := make(map[int64]*model.Menu)

	// 过滤掉按钮类型的菜单(F)和链接类型的菜单(L)
	filtered := make([]model.Menu, 0)
	for _, m := range menus {
		if m.MenuType == "F" || m.MenuType == "L" {
			continue // 跳过按钮和链接
		}
		filtered = append(filtered, m)
	}

	for i := range filtered {
		menuMap[filtered[i].ID] = &filtered[i]
	}

	for i := range filtered {
		if filtered[i].ParentID == 0 {
			result = append(result, &filtered[i])
		} else {
			if parent, ok := menuMap[filtered[i].ParentID]; ok {
				parent.Children = append(parent.Children, filtered[i])
			}
		}
	}

	// 转换为普通 slice
	final := make([]model.Menu, len(result))
	for i, m := range result {
		final[i] = *m
	}
	return final
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

func (r *RoleMenuRepository) GetMenuIDsByRoleID(ctx context.Context, roleID int64) ([]int64, error) {
	var menuIDs []int64
	err := r.db.WithContext(ctx).Model(&model.RoleMenu{}).
		Where("role_id = ?", roleID).
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

// GetRolePerms 获取角色权限列表
func (r *RoleMenuRepository) GetRolePerms(ctx context.Context, roleID int64) ([]string, error) {
	var perms []string
	err := r.db.WithContext(ctx).Model(&model.RolePerm{}).
		Where("role_id = ?", roleID).
		Pluck("perm", &perms).Error
	return perms, err
}

// AssignPerms 分配角色权限
func (r *RoleMenuRepository) AssignPerms(ctx context.Context, roleID int64, perms []string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id = ?", roleID).Delete(&model.RolePerm{}).Error; err != nil {
			return err
		}
		for _, perm := range perms {
			if err := tx.Create(&model.RolePerm{RoleID: roleID, Perm: perm}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
