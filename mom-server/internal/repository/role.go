package repository

import (
	"context"

	"gorm.io/gorm"
	"mom-server/internal/dto"
	"mom-server/internal/model"
)

// RoleRepository 角色仓储
type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

// Create 创建角色
func (r *RoleRepository) Create(ctx context.Context, role *model.Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

// Update 更新角色
func (r *RoleRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Role{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除角色
func (r *RoleRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Role{}, id).Error
}

// FindByID 根据ID查询
func (r *RoleRepository) FindByID(ctx context.Context, id int64) (*model.Role, error) {
	var role model.Role
	err := r.db.WithContext(ctx).First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// FindByPage 分页查询
func (r *RoleRepository) FindByPage(ctx context.Context, tenantID int64, req dto.PageRequest, roleName, status string) ([]model.Role, int64, error) {
	var roles []model.Role
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Role{}).Where("tenant_id = ?", tenantID)

	if roleName != "" {
		query = query.Where("role_name LIKE ?", "%"+roleName+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(req.GetOffset()).Limit(req.GetPageSize()).Order("role_sort ASC").Find(&roles).Error
	return roles, total, err
}

// FindAll 获取所有角色
func (r *RoleRepository) FindAll(ctx context.Context, tenantID int64) ([]model.Role, error) {
	var roles []model.Role
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND status = ?", tenantID, 1).Find(&roles).Error
	return roles, err
}

// AssignMenus 分配菜单权限
func (r *RoleRepository) AssignMenus(ctx context.Context, roleID int64, menuIDs []int64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 删除原有菜单
		if err := tx.Where("role_id = ?", roleID).Delete(&model.RoleMenu{}).Error; err != nil {
			return err
		}
		// 添加新菜单
		for _, menuID := range menuIDs {
			if err := tx.Create(&model.RoleMenu{RoleID: roleID, MenuID: menuID}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// GetRoleMenus 获取角色菜单
func (r *RoleRepository) GetRoleMenus(ctx context.Context, roleID int64) ([]int64, error) {
	var menuIDs []int64
	err := r.db.WithContext(ctx).Model(&model.RoleMenu{}).
		Where("role_id = ?", roleID).
		Pluck("menu_id", &menuIDs).Error
	return menuIDs, err
}

// RoleMenu 角色菜单关联
type RoleMenu struct {
	RoleID int64 `gorm:"primaryKey"`
	MenuID int64 `gorm:"primaryKey"`
}

func (RoleMenu) TableName() string {
	return "sys_role_menu"
}
