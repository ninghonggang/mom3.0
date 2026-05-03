package service

import (
	"context"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type FactoryService struct {
	repo *repository.FactoryRepository
}

func NewFactoryService(repo *repository.FactoryRepository) *FactoryService {
	return &FactoryService{repo: repo}
}

func (s *FactoryService) Create(ctx context.Context, tenantID int64, req *model.FactoryCreateReq) (*model.MdmFactory, error) {
	factory := &model.MdmFactory{
		TenantID:    tenantID,
		FactoryCode: req.FactoryCode,
		FactoryName: req.FactoryName,
		Province:    req.Province,
		City:        req.City,
		District:    req.District,
		Address:     req.Address,
		Manager:    req.Manager,
		Phone:      req.Phone,
		AreaSize:   req.AreaSize,
		Status:     req.Status,
		IsDefault:  req.IsDefault,
	}
	if factory.Status == 0 {
		factory.Status = 1
	}
	if err := s.repo.Create(ctx, factory); err != nil {
		return nil, err
	}
	return factory, nil
}

func (s *FactoryService) GetByID(ctx context.Context, id int64) (*model.MdmFactory, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *FactoryService) List(ctx context.Context, tenantID int64, req *model.FactoryQuery) ([]model.MdmFactory, int64, error) {
	return s.repo.Page(ctx, tenantID, req)
}

func (s *FactoryService) ListAll(ctx context.Context, tenantID int64) ([]model.MdmFactory, error) {
	return s.repo.ListByTenant(ctx, tenantID)
}

func (s *FactoryService) Update(ctx context.Context, id int64, req *model.FactoryCreateReq) error {
	updates := map[string]interface{}{
		"factory_code": req.FactoryCode,
		"factory_name": req.FactoryName,
		"province":     req.Province,
		"city":         req.City,
		"district":     req.District,
		"address":      req.Address,
		"manager":      req.Manager,
		"phone":        req.Phone,
		"area_size":    req.AreaSize,
		"status":       req.Status,
		"is_default":   req.IsDefault,
	}
	return s.repo.Update(ctx, id, updates)
}

func (s *FactoryService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *FactoryService) SetDefault(ctx context.Context, tenantID, factoryID int64) error {
	return s.repo.SetDefault(ctx, tenantID, factoryID)
}

func (s *FactoryService) SetCurrentFactory(ctx context.Context, tenantID, userID, factoryID int64) error {
	return s.repo.SetCurrentFactory(ctx, tenantID, userID, factoryID)
}

func (s *FactoryService) GetCurrentFactory(ctx context.Context, tenantID, userID int64) (*model.MdmFactory, error) {
	return s.repo.GetCurrentFactory(ctx, tenantID, userID)
}