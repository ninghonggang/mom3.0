package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type InspectionFeatureService struct {
	repo *repository.InspectionFeatureRepository
}

func NewInspectionFeatureService(repo *repository.InspectionFeatureRepository) *InspectionFeatureService {
	return &InspectionFeatureService{repo: repo}
}

// List 查询检验特性列表
func (s *InspectionFeatureService) List(ctx context.Context, tenantID uint64, query map[string]interface{}) ([]model.InspectionFeature, int64, error) {
	return s.repo.List(ctx, tenantID, query)
}

// GetByID 根据ID获取检验特性
func (s *InspectionFeatureService) GetByID(ctx context.Context, id uint64) (*model.InspectionFeature, error) {
	return s.repo.GetByID(ctx, id)
}

// Create 创建检验特性
func (s *InspectionFeatureService) Create(ctx context.Context, tenantID uint64, req *model.InspectionFeatureCreateRequest, username string) (*model.InspectionFeature, error) {
	feature := &model.InspectionFeature{
		TenantID:       tenantID,
		FeatureCode:    req.FeatureCode,
		FeatureName:    req.FeatureName,
		ProductID:      req.ProductID,
		ProductCode:    req.ProductCode,
		ProductName:    req.ProductName,
		InspectionType: req.InspectionType,
		FeatureType:    req.FeatureType,
		TechnicalSpec:  req.TechnicalSpec,
		LowerLimit:     req.LowerLimit,
		UpperLimit:     req.UpperLimit,
		Unit:           req.Unit,
		SampleSize:     req.SampleSize,
		GaugesMethod:   req.GaugesMethod,
		AQLLevel:       req.AQLLevel,
		Status:         req.Status,
		Remark:         req.Remark,
		CreatedBy:      username,
		UpdatedBy:      username,
	}

	// 默认状态
	if feature.Status == "" {
		feature.Status = "ACTIVE"
	}

	err := s.repo.Create(ctx, feature)
	if err != nil {
		return nil, fmt.Errorf("failed to create inspection feature: %w", err)
	}
	return feature, nil
}

// Update 更新检验特性
func (s *InspectionFeatureService) Update(ctx context.Context, id uint64, req *model.InspectionFeatureUpdateRequest, username string) error {
	feature, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("inspection feature not found: %w", err)
	}

	// 更新字段
	if req.FeatureName != "" {
		feature.FeatureName = req.FeatureName
	}
	if req.ProductID > 0 {
		feature.ProductID = req.ProductID
	}
	if req.ProductCode != "" {
		feature.ProductCode = req.ProductCode
	}
	if req.ProductName != "" {
		feature.ProductName = req.ProductName
	}
	if req.InspectionType != "" {
		feature.InspectionType = req.InspectionType
	}
	if req.FeatureType != "" {
		feature.FeatureType = req.FeatureType
	}
	if req.TechnicalSpec != "" {
		feature.TechnicalSpec = req.TechnicalSpec
	}
	if req.LowerLimit != nil {
		feature.LowerLimit = req.LowerLimit
	}
	if req.UpperLimit != nil {
		feature.UpperLimit = req.UpperLimit
	}
	if req.Unit != nil {
		feature.Unit = req.Unit
	}
	if req.SampleSize > 0 {
		feature.SampleSize = req.SampleSize
	}
	if req.GaugesMethod != "" {
		feature.GaugesMethod = req.GaugesMethod
	}
	if req.AQLLevel != "" {
		feature.AQLLevel = req.AQLLevel
	}
	if req.Status != "" {
		feature.Status = req.Status
	}
	if req.Remark != "" {
		feature.Remark = req.Remark
	}
	feature.UpdatedBy = username

	return s.repo.Update(ctx, feature)
}

// Delete 删除检验特性
func (s *InspectionFeatureService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

// GetByProductID 获取产品的所有检验特性
func (s *InspectionFeatureService) GetByProductID(ctx context.Context, tenantID uint64, productID uint64) ([]model.InspectionFeature, error) {
	return s.repo.GetByProductID(ctx, tenantID, productID)
}

// BatchCreate 批量创建检验特性
func (s *InspectionFeatureService) BatchCreate(ctx context.Context, tenantID uint64, req *model.InspectionFeatureBatchCreateRequest, username string) ([]model.InspectionFeature, error) {
	if len(req.Features) == 0 {
		return nil, fmt.Errorf("features list is empty")
	}

	features := make([]model.InspectionFeature, 0, len(req.Features))
	for _, f := range req.Features {
		feature := model.InspectionFeature{
			TenantID:       tenantID,
			FeatureCode:    f.FeatureCode,
			FeatureName:    f.FeatureName,
			ProductID:      f.ProductID,
			ProductCode:    f.ProductCode,
			ProductName:    f.ProductName,
			InspectionType: f.InspectionType,
			FeatureType:    f.FeatureType,
			TechnicalSpec:  f.TechnicalSpec,
			LowerLimit:     f.LowerLimit,
			UpperLimit:     f.UpperLimit,
			Unit:           f.Unit,
			SampleSize:     f.SampleSize,
			GaugesMethod:   f.GaugesMethod,
			AQLLevel:       f.AQLLevel,
			Status:         f.Status,
			Remark:         f.Remark,
			CreatedBy:      username,
			UpdatedBy:      username,
		}
		if feature.Status == "" {
			feature.Status = "ACTIVE"
		}
		features = append(features, feature)
	}

	err := s.repo.BatchCreate(ctx, features)
	if err != nil {
		return nil, fmt.Errorf("failed to batch create inspection features: %w", err)
	}
	return features, nil
}