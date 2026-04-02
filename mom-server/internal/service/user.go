package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"mom-server/internal/dto"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

var (
	ErrUserNotFound     = errors.New("用户不存在")
	ErrUserExists      = errors.New("用户名已存在")
	ErrInvalidPassword = errors.New("用户名或密码错误")
	ErrUserDisabled    = errors.New("用户已被禁用")
)

// UserService 用户服务
type UserService struct {
	userRepo    *repository.UserRepository
	roleRepo    *repository.RoleRepository
	menuRepo    *repository.MenuRepository
	roleMenuDB  *repository.RoleMenuRepository
}

func NewUserService(userRepo *repository.UserRepository, roleRepo *repository.RoleRepository, menuRepo *repository.MenuRepository, roleMenuDB *repository.RoleMenuRepository) *UserService {
	return &UserService{
		userRepo:   userRepo,
		roleRepo:   roleRepo,
		menuRepo:   menuRepo,
		roleMenuDB: roleMenuDB,
	}
}

// ValidateLogin 验证登录（不生成Token）
func (s *UserService) ValidateLogin(ctx context.Context, req dto.LoginRequest) (*model.User, error) {
	// 登录时查询所有租户，找到匹配的用户
	user, err := s.userRepo.FindByUsernameAllTenants(ctx, req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidPassword
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, ErrInvalidPassword
	}

	if user.Status == 0 {
		return nil, ErrUserDisabled
	}

	return user, nil
}

// GetUserRoles 获取用户角色ID列表
func (s *UserService) GetUserRoles(ctx context.Context, userID int64) ([]int64, error) {
	return s.userRepo.GetUserRoles(ctx, userID)
}

// GetAllRoles 获取所有角色
func (s *UserService) GetAllRoles(ctx context.Context, tenantID int64) ([]model.Role, error) {
	return s.roleRepo.FindAll(ctx, tenantID)
}

// AssignRoles 分配角色
func (s *UserService) AssignRoles(ctx context.Context, userID int64, roleIDs []int64) error {
	return s.userRepo.AssignRoles(ctx, userID, roleIDs)
}

// Create 创建用户
func (s *UserService) Create(ctx context.Context, tenantID int64, req dto.CreateUserRequest) error {
	// 检查用户名是否存在
	existing, _ := s.userRepo.FindByUsername(ctx, tenantID, req.Username)
	if existing != nil {
		return ErrUserExists
	}

	// 加密密码（使用更高的成本因子，生产环境建议12-14）
	bcryptCost := 12
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcryptCost)
	if err != nil {
		return err
	}

	user := &model.User{
		TenantModel: model.TenantModel{TenantID: tenantID},
		Username:    req.Username,
		Nickname:    req.Nickname,
		Password:    string(hashedPassword),
		Email:       req.Email,
		Phone:       req.Phone,
		DeptID:      req.DeptID,
		Status:      req.Status,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return err
	}

	// 分配角色
	if len(req.RoleIDs) > 0 {
		s.userRepo.AssignRoles(ctx, user.ID, req.RoleIDs)
	}

	return nil
}

// Update 更新用户
func (s *UserService) Update(ctx context.Context, id int64, req dto.UpdateUserRequest) error {
	updates := map[string]interface{}{
		"nickname": req.Nickname,
		"email":    req.Email,
		"phone":    req.Phone,
		"dept_id":  req.DeptID,
		"status":   req.Status,
	}

	if err := s.userRepo.Update(ctx, id, updates); err != nil {
		return err
	}

	// 更新角色
	if req.RoleIDs != nil {
		s.userRepo.AssignRoles(ctx, id, req.RoleIDs)
	}

	return nil
}

// Delete 删除用户
func (s *UserService) Delete(ctx context.Context, id int64) error {
	return s.userRepo.Delete(ctx, id)
}

// GetByID 获取用户详情（优化N+1查询）
func (s *UserService) GetByID(ctx context.Context, id int64) (*dto.UserDTO, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// 获取角色IDs
	roleIDs, _ := s.userRepo.GetUserRoles(ctx, user.ID)
	if len(roleIDs) == 0 {
		return &dto.UserDTO{
			ID:       user.ID,
			Username: user.Username,
			Nickname: user.Nickname,
			Email:    user.Email,
			Phone:    user.Phone,
			Avatar:   user.Avatar,
			DeptID:   user.DeptID,
			Status:   user.Status,
			Roles:    []string{},
			RoleIDs:  []int64{},
			Perms:    []string{},
			Menus:    []model.Menu{},
		}, nil
	}

	// 批量查询所有角色（单次查询替代N次查询）
	roleList, _ := s.roleRepo.FindAll(ctx, user.TenantID)
	roleMap := make(map[int64]string)
	var roles []string
	for _, r := range roleList {
		roleMap[r.ID] = r.RoleKey
	}
	for _, rid := range roleIDs {
		if rk, ok := roleMap[rid]; ok {
			roles = append(roles, rk)
		}
	}

	// 获取用户菜单（根据角色权限）
	var menus []model.Menu
	isAdmin := false
	for _, r := range roles {
		if r == "admin" {
			isAdmin = true
			break
		}
	}

	var perms []string
	if isAdmin {
		// 超管返回所有菜单（单次查询）
		menus, _ = s.menuRepo.Tree(ctx, user.TenantID)
		// 超管拥有所有权限
		perms = []string{"*"}
	} else if len(roleIDs) > 0 {
		// 普通用户根据角色权限获取菜单
		var allMenuIDs []int64
		for _, roleID := range roleIDs {
			menuIDs, _ := s.roleMenuDB.GetMenuIDsByRoleID(ctx, roleID)
			allMenuIDs = append(allMenuIDs, menuIDs...)
		}
		if len(allMenuIDs) > 0 {
			// 去重menuIDs
			seen := make(map[int64]bool)
			uniqueMenuIDs := []int64{}
			for _, id := range allMenuIDs {
				if !seen[id] {
					seen[id] = true
					uniqueMenuIDs = append(uniqueMenuIDs, id)
				}
			}

			uids := make([]uint, len(uniqueMenuIDs))
			for i, m := range uniqueMenuIDs {
				uids[i] = uint(m)
			}
			menus, _ = s.menuRepo.GetByIDs(ctx, uids)
			menus = s.menuRepo.BuildTree(menus)
			// 收集权限标识（解析逗号分隔的多个权限码）
			for _, m := range menus {
				if m.Perms != "" {
					// 拆分逗号分隔的权限码
					for _, p := range strings.Split(m.Perms, ",") {
						p = strings.TrimSpace(p)
						if p != "" {
							perms = append(perms, p)
						}
					}
				}
			}
		}
	}

	return &dto.UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Phone:    user.Phone,
		Avatar:   user.Avatar,
		DeptID:   user.DeptID,
		Status:   user.Status,
		Roles:    roles,
		RoleIDs:  roleIDs,
		Perms:    perms,
		Menus:    menus,
	}, nil
}

// GetList 获取用户列表
func (s *UserService) GetList(ctx context.Context, tenantID int64, req dto.PageRequest, username, status string) (*dto.PageData, error) {
	users, total, err := s.userRepo.FindByPage(ctx, tenantID, req, username, status)
	if err != nil {
		return nil, err
	}

	// 转换DTO
	list := make([]dto.UserDTO, len(users))
	for i, user := range users {
		list[i] = dto.UserDTO{
			ID:       user.ID,
			Username: user.Username,
			Nickname: user.Nickname,
			Email:    user.Email,
			Phone:    user.Phone,
			Status:   user.Status,
			DeptID:   user.DeptID,
		}
	}

	return &dto.PageData{
		List:     list,
		Total:    total,
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
	}, nil
}

// ResetPassword 重置密码
func (s *UserService) ResetPassword(ctx context.Context, id int64, password string) error {
	// 加密密码（使用更高的成本因子，生产环境建议12-14）
	bcryptCost := 12
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return err
	}

	return s.userRepo.Update(ctx, id, map[string]interface{}{
		"password": string(hashedPassword),
	})
}

// ChangePassword 修改密码
func (s *UserService) ChangePassword(ctx context.Context, id int64, oldPassword, newPassword string) error {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// 验证旧密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return ErrInvalidPassword
	}

	// 设置新密码
	return s.ResetPassword(ctx, id, newPassword)
}

// UpdateLoginInfo 更新登录信息
func (s *UserService) UpdateLoginInfo(ctx context.Context, id int64, ip string) error {
	now := time.Now()
	return s.userRepo.Update(ctx, id, map[string]interface{}{
		"login_ip":   ip,
		"login_date": now,
	})
}
