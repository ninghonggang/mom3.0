package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type DeptService struct {
	repo *repository.DeptRepository
}

func NewDeptService(repo *repository.DeptRepository) *DeptService {
	return &DeptService{repo: repo}
}

func (s *DeptService) List(ctx context.Context) ([]model.Dept, error) {
	return s.repo.List(ctx, 0)
}

func (s *DeptService) Tree(ctx context.Context) ([]model.Dept, error) {
	return s.repo.Tree(ctx, 0)
}

func (s *DeptService) GetByID(ctx context.Context, id string) (*model.Dept, error) {
	var deptID uint
	_, err := fmt.Sscanf(id, "%d", &deptID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, deptID)
}

func (s *DeptService) Create(ctx context.Context, dept *model.Dept) error {
	return s.repo.Create(ctx, dept)
}

func (s *DeptService) Update(ctx context.Context, id string, dept *model.Dept) error {
	var deptID uint
	_, err := fmt.Sscanf(id, "%d", &deptID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, deptID, map[string]interface{}{
		"parent_id": dept.ParentID,
		"dept_name": dept.DeptName,
		"dept_sort": dept.DeptSort,
		"leader":    dept.Leader,
		"phone":     dept.Phone,
		"email":     dept.Email,
		"status":    dept.Status,
	})
}

func (s *DeptService) Delete(ctx context.Context, id string) error {
	var deptID uint
	_, err := fmt.Sscanf(id, "%d", &deptID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, deptID)
}
