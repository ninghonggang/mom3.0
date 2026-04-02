package service

import (
	"context"
	"errors"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type TenantService struct {
	tenantRepo *repository.TenantRepository
}

func NewTenantService(tenantRepo *repository.TenantRepository) *TenantService {
	return &TenantService{tenantRepo: tenantRepo}
}

func (s *TenantService) Create(ctx context.Context, req *model.Tenant) error {
	// 检查TenantKey唯一性
	existing, _ := s.tenantRepo.FindByKey(ctx, req.TenantKey)
	if existing != nil && existing.ID != 0 {
		return errors.New("租户标识已存在")
	}
	return s.tenantRepo.Create(ctx, req)
}

func (s *TenantService) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return s.tenantRepo.Update(ctx, id, updates)
}

func (s *TenantService) Delete(ctx context.Context, id int64) error {
	return s.tenantRepo.Delete(ctx, id)
}

func (s *TenantService) GetByID(ctx context.Context, id int64) (*model.Tenant, error) {
	return s.tenantRepo.FindByID(ctx, id)
}

func (s *TenantService) GetList(ctx context.Context, tenantName, tenantKey, status string, page, pageSize int) ([]model.Tenant, int64, error) {
	return s.tenantRepo.FindByPage(ctx, tenantName, tenantKey, status, page, pageSize)
}

func (s *TenantService) GetAll(ctx context.Context) ([]model.Tenant, error) {
	tenants, _, err := s.tenantRepo.FindByPage(ctx, "", "", "", 1, 1000)
	if err != nil {
		return nil, err
	}
	return tenants, nil
}
