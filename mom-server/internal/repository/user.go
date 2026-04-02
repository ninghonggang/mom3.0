package repository

import (
	"context"

	"gorm.io/gorm"
	"mom-server/internal/dto"
	"mom-server/internal/model"
)

// UserRepository 用户仓储
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create 创建用户
func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// Update 更新用户
func (r *UserRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除用户
func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.User{}, id).Error
}

// FindByID 根据ID查询
func (r *UserRepository) FindByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByUsername 根据用户名查询（指定租户）
func (r *UserRepository) FindByUsername(ctx context.Context, tenantID int64, username string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).
		Where("tenant_id = ? AND username = ?", tenantID, username).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByUsernameAllTenants 根据用户名查询所有租户（用于登录）
func (r *UserRepository) FindByUsernameAllTenants(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).
		Where("username = ?", username).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByPage 分页查询
func (r *UserRepository) FindByPage(ctx context.Context, tenantID int64, req dto.PageRequest, username, status string) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	query := r.db.WithContext(ctx).Model(&model.User{}).Where("tenant_id = ?", tenantID)

	if username != "" {
		query = query.Where("username LIKE ? OR nickname LIKE ?", "%"+username+"%", "%"+username+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(req.GetOffset()).Limit(req.GetPageSize()).Order("id DESC").Find(&users).Error
	return users, total, err
}

// AssignRoles 分配角色
func (r *UserRepository) AssignRoles(ctx context.Context, userID int64, roleIDs []int64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 删除原有角色
		if err := tx.Where("user_id = ?", userID).Delete(&model.UserRole{}).Error; err != nil {
			return err
		}
		// 添加新角色
		for _, roleID := range roleIDs {
			if err := tx.Create(&model.UserRole{UserID: userID, RoleID: roleID}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// GetUserRoles 获取用户角色
func (r *UserRepository) GetUserRoles(ctx context.Context, userID int64) ([]int64, error) {
	var roleIDs []int64
	err := r.db.WithContext(ctx).Model(&model.UserRole{}).
		Where("user_id = ?", userID).
		Pluck("role_id", &roleIDs).Error
	return roleIDs, err
}

// GetUserByIDs 批量获取用户
func (r *UserRepository) GetUserByIDs(ctx context.Context, ids []int64) ([]model.User, error) {
	var users []model.User
	err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&users).Error
	return users, err
}

// UserRole 用户角色关联
type UserRole struct {
	UserID int64 `gorm:"primaryKey"`
	RoleID int64 `gorm:"primaryKey"`
}

func (UserRole) TableName() string {
	return "sys_user_role"
}
