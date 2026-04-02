package service

import (
	"context"
	"fmt"
	"mom-server/internal/dto"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type RoleService struct {
	repo       *repository.RoleRepository
	menuRepo   *repository.MenuRepository
	roleMenuDB *repository.RoleMenuRepository
}

func NewRoleService(repo *repository.RoleRepository, menuRepo *repository.MenuRepository, roleMenuDB *repository.RoleMenuRepository) *RoleService {
	return &RoleService{
		repo:       repo,
		menuRepo:   menuRepo,
		roleMenuDB: roleMenuDB,
	}
}

func (s *RoleService) List(ctx context.Context, tenantID int64, req *dto.RoleListReq) ([]model.Role, int64, error) {
	pageReq := dto.PageRequest{Page: req.Page, PageSize: req.PageSize}
	return s.repo.FindByPage(ctx, tenantID, pageReq, req.RoleName, "")
}

func (s *RoleService) GetByID(ctx context.Context, id string) (*model.Role, error) {
	var roleID int64
	_, err := fmt.Sscanf(id, "%d", &roleID)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByID(ctx, roleID)
}

func (s *RoleService) Create(ctx context.Context, role *model.Role) error {
	return s.repo.Create(ctx, role)
}

func (s *RoleService) Update(ctx context.Context, id string, role *model.Role) error {
	var roleID int64
	_, err := fmt.Sscanf(id, "%d", &roleID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, roleID, map[string]interface{}{
		"role_name": role.RoleName,
		"role_key":  role.RoleKey,
		"role_sort": role.RoleSort,
		"status":    role.Status,
		"remark":    role.Remark,
	})
}

func (s *RoleService) Delete(ctx context.Context, id string) error {
	var roleID int64
	_, err := fmt.Sscanf(id, "%d", &roleID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, roleID)
}

func (s *RoleService) GetMenus(ctx context.Context, roleID string) ([]model.Menu, error) {
	var rID int64
	_, err := fmt.Sscanf(roleID, "%d", &rID)
	if err != nil {
		return nil, err
	}
	menuIDs, err := s.roleMenuDB.GetMenuIDsByRoleID(ctx, rID)
	if err != nil {
		return nil, err
	}
	if len(menuIDs) == 0 {
		return []model.Menu{}, nil
	}
	uids := make([]uint, len(menuIDs))
	for i, m := range menuIDs {
		uids[i] = uint(m)
	}
	return s.menuRepo.GetByIDs(ctx, uids)
}

func (s *RoleService) AssignMenus(ctx context.Context, roleID string, menuIDs []uint) error {
	var rID int64
	_, err := fmt.Sscanf(roleID, "%d", &rID)
	if err != nil {
		return err
	}
	mIDs := make([]int64, len(menuIDs))
	for i, m := range menuIDs {
		mIDs[i] = int64(m)
	}
	return s.roleMenuDB.AssignMenus(ctx, rID, mIDs)
}

func (s *RoleService) GetRoleMenusTree(ctx context.Context, roleID string) ([]model.Menu, error) {
	menus, err := s.GetMenus(ctx, roleID)
	if err != nil {
		return nil, err
	}
	return s.menuRepo.BuildTree(menus), nil
}

// GetRolePerms 获取角色权限列表
func (s *RoleService) GetRolePerms(ctx context.Context, roleID string) ([]string, error) {
	var rID int64
	_, err := fmt.Sscanf(roleID, "%d", &rID)
	if err != nil {
		return nil, err
	}
	return s.roleMenuDB.GetRolePerms(ctx, rID)
}

// AssignPerms 分配角色权限
func (s *RoleService) AssignPerms(ctx context.Context, roleID string, perms []string) error {
	var rID int64
	_, err := fmt.Sscanf(roleID, "%d", &rID)
	if err != nil {
		return err
	}
	return s.roleMenuDB.AssignPerms(ctx, rID, perms)
}
