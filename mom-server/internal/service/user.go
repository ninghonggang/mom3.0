package service

import (
	"context"
	"errors"
	"fmt"
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
	userRepo *repository.UserRepository
	roleRepo *repository.RoleRepository
}

func NewUserService(userRepo *repository.UserRepository, roleRepo *repository.RoleRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

// Login 登录
func (s *UserService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	// 查询用户
	user, err := s.userRepo.FindByUsername(ctx, 1, req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidPassword
		}
		return nil, err
	}

	// Debug: 打印密码和哈希
	fmt.Printf("DEBUG: username=%s, password=%s, hash=%s\n", req.Username, req.Password, user.Password)

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		fmt.Printf("DEBUG: bcrypt error: %v\n", err)
		return nil, ErrInvalidPassword
	}

	// 检查状态
	if user.Status == 0 {
		return nil, ErrUserDisabled
	}

	// 获取用户角色
	roleIDs, _ := s.userRepo.GetUserRoles(ctx, user.ID)
	var roles []string
	if len(roleIDs) > 0 {
		roleList, _ := s.roleRepo.FindAll(ctx, user.TenantID)
		roleMap := make(map[int64]string)
		for _, r := range roleList {
			roleMap[r.ID] = r.RoleKey
		}
		for _, rid := range roleIDs {
			if rk, ok := roleMap[rid]; ok {
				roles = append(roles, rk)
			}
		}
	}

	// 生成Token（简化版，实际应该使用JWT服务）
	token := "mock-token-" + user.Username
	refreshToken := "mock-refresh-" + user.Username

	return &dto.LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		ExpiresIn:    7200,
		User: &dto.UserDTO{
			ID:       user.ID,
			Username: user.Username,
			Nickname: user.Nickname,
			Email:    user.Email,
			Phone:    user.Phone,
			Avatar:   user.Avatar,
			DeptID:   user.DeptID,
			Status:   user.Status,
			Roles:    roles,
		},
	}, nil
}

// ValidateLogin 验证登录（不生成Token）
func (s *UserService) ValidateLogin(ctx context.Context, req dto.LoginRequest) (*model.User, error) {
	user, err := s.userRepo.FindByUsername(ctx, 1, req.Username)
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

// Create 创建用户
func (s *UserService) Create(ctx context.Context, req dto.CreateUserRequest) error {
	// 检查用户名是否存在
	existing, _ := s.userRepo.FindByUsername(ctx, 1, req.Username)
	if existing != nil {
		return ErrUserExists
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		TenantModel: model.TenantModel{TenantID: 1},
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

// GetByID 获取用户详情
func (s *UserService) GetByID(ctx context.Context, id int64) (*dto.UserDTO, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// 获取角色
	roleIDs, _ := s.userRepo.GetUserRoles(ctx, user.ID)
	var roles []string
	if len(roleIDs) > 0 {
		roleList, _ := s.roleRepo.FindAll(ctx, user.TenantID)
		roleMap := make(map[int64]string)
		for _, r := range roleList {
			roleMap[r.ID] = r.RoleKey
		}
		for _, rid := range roleIDs {
			if rk, ok := roleMap[rid]; ok {
				roles = append(roles, rk)
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
	}, nil
}

// GetList 获取用户列表
func (s *UserService) GetList(ctx context.Context, req dto.PageRequest, username, status string) (*dto.PageData, error) {
	users, total, err := s.userRepo.FindByPage(ctx, 1, req, username, status)
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
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
