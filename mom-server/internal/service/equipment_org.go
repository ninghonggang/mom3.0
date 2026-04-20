package service

import (
	"context"
	"errors"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"strconv"
)

// EquipmentOrgService 设备组织服务
type EquipmentOrgService struct {
	repo          *repository.EquipmentOrgRepository
	workshopRepo  *repository.WorkshopRepository
	lineRepo      *repository.ProductionLineRepository
}

// NewEquipmentOrgService 创建设备组织服务
func NewEquipmentOrgService(
	repo *repository.EquipmentOrgRepository,
	workshopRepo *repository.WorkshopRepository,
	lineRepo *repository.ProductionLineRepository,
) *EquipmentOrgService {
	return &EquipmentOrgService{
		repo:         repo,
		workshopRepo: workshopRepo,
		lineRepo:     lineRepo,
	}
}

// List 获取设备组织列表
func (s *EquipmentOrgService) List(ctx context.Context, query *model.EquipmentOrgQuery) ([]model.EquipmentOrg, int64, error) {
	return s.repo.List(ctx, 0, query)
}

// GetByID 根据ID获取设备组织
func (s *EquipmentOrgService) GetByID(ctx context.Context, id string) (*model.EquipmentOrg, error) {
	orgID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, errors.New("invalid equipment org id")
	}
	return s.repo.GetByID(ctx, uint(orgID))
}

// SyncFromMasterData 从主数据（车间、产线）同步设备组织关系
func (s *EquipmentOrgService) SyncFromMasterData(ctx context.Context, tenantID int64) error {
	// 获取所有车间
	workshops, _, err := s.workshopRepo.List(ctx, tenantID)
	if err != nil {
		return err
	}

	// 获取所有产线
	lines, _, err := s.lineRepo.List(ctx, tenantID)
	if err != nil {
		return err
	}

	// 按车间分组产线
	lineMap := make(map[int64][]model.ProductionLine)
	for _, line := range lines {
		if line.WorkshopID > 0 {
			lineMap[line.WorkshopID] = append(lineMap[line.WorkshopID], line)
		}
	}

	// 厂区信息（从租户或配置获取，这里简化处理）
	for _, workshop := range workshops {
		factoryCode := "F001"
		factoryName := "默认厂区"
		if workshop.WorkshopCode != "" {
			// 实际应该关联厂区表
			factoryCode = "F" + workshop.WorkshopCode[:1]
			factoryName = workshop.WorkshopName + "-厂区"
		}

		workshopLines := lineMap[workshop.ID]
		if len(workshopLines) == 0 {
			// 只有车间没有产线时，也创建一条记录
			org := &model.EquipmentOrg{
				TenantID:     tenantID,
				FactoryID:    1,
				FactoryCode:  factoryCode,
				FactoryName:  factoryName,
				WorkshopID:   workshop.ID,
				WorkshopCode: workshop.WorkshopCode,
				WorkshopName: workshop.WorkshopName,
				LineID:       0,
				LineCode:     "",
				LineName:     "",
				Status:       workshop.Status,
			}
			s.repo.CreateSync(ctx, org)
		}

		for _, line := range workshopLines {
			org := &model.EquipmentOrg{
				TenantID:     tenantID,
				FactoryID:    1,
				FactoryCode:  factoryCode,
				FactoryName:  factoryName,
				WorkshopID:   workshop.ID,
				WorkshopCode: workshop.WorkshopCode,
				WorkshopName: workshop.WorkshopName,
				LineID:       line.ID,
				LineCode:     line.LineCode,
				LineName:     line.LineName,
				Status:       line.Status,
			}
			if err := s.repo.CreateSync(ctx, org); err != nil {
				return err
			}
		}
	}

	return nil
}

// Update 更新设备组织
func (s *EquipmentOrgService) Update(ctx context.Context, id string, updates map[string]interface{}) error {
	orgID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("invalid equipment org id")
	}
	return s.repo.Update(ctx, uint(orgID), updates)
}

// Delete 删除设备组织
func (s *EquipmentOrgService) Delete(ctx context.Context, id string) error {
	orgID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("invalid equipment org id")
	}
	return s.repo.Delete(ctx, uint(orgID))
}

// GetFactoryList 获取厂区列表
func (s *EquipmentOrgService) GetFactoryList(ctx context.Context) ([]model.Factory, error) {
	return s.repo.GetFactoryList(ctx, 0)
}

// CreateFactory 创建厂区
func (s *EquipmentOrgService) CreateFactory(ctx context.Context, factory *model.Factory) error {
	return s.repo.CreateFactory(ctx, factory)
}

// GetFactoryByID 根据ID获取厂区
func (s *EquipmentOrgService) GetFactoryByID(ctx context.Context, id string) (*model.Factory, error) {
	factoryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, errors.New("invalid factory id")
	}
	return s.repo.GetFactoryByID(ctx, uint(factoryID))
}

// GetFactoryByCode 根据编码获取厂区
func (s *EquipmentOrgService) GetFactoryByCode(ctx context.Context, tenantID int64, code string) (*model.Factory, error) {
	return s.repo.GetFactoryByCode(ctx, tenantID, code)
}

// UpdateFactory 更新厂区
func (s *EquipmentOrgService) UpdateFactory(ctx context.Context, id string, updates map[string]interface{}) error {
	factoryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("invalid factory id")
	}
	return s.repo.UpdateFactory(ctx, uint(factoryID), updates)
}

// DeleteFactory 删除厂区
func (s *EquipmentOrgService) DeleteFactory(ctx context.Context, id string) error {
	factoryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("invalid factory id")
	}
	return s.repo.DeleteFactory(ctx, uint(factoryID))
}
