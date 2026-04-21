package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// WmsAreaService 库区服务层
type WmsAreaService struct {
	repo *repository.WmsAreaRepository
}

// NewWmsAreaService 创建库区服务
func NewWmsAreaService(repo *repository.WmsAreaRepository) *WmsAreaService {
	return &WmsAreaService{repo: repo}
}

// Create 创建库区
func (s *WmsAreaService) Create(ctx context.Context, tenantID int64, req *model.WmsAreaCreateReqVO) error {
	if tenantID <= 0 {
		tenantID = 1
	}

	// 检查编码是否已存在
	existing, _ := s.repo.GetByCode(ctx, tenantID, req.AreaCode)
	if existing != nil && existing.ID > 0 {
		return fmt.Errorf("库区编码已存在: %s", req.AreaCode)
	}

	area := &model.WmsArea{
		TenantID:      tenantID,
		WarehouseCode: req.WarehouseCode,
		WarehouseName: req.WarehouseName,
		AreaCode:      req.AreaCode,
		AreaName:      req.AreaName,
		AreaType:      req.AreaType,
		ParentCode:    req.ParentCode,
		Level:         req.Level,
		Status:        req.Status,
		Remark:        req.Remark,
	}

	if area.Status == "" {
		area.Status = "ACTIVE"
	}

	return s.repo.Create(ctx, area)
}

// Update 更新库区
func (s *WmsAreaService) Update(ctx context.Context, id uint64, req *model.WmsAreaUpdateReqVO) error {
	updates := make(map[string]interface{})
	if req.WarehouseName != "" {
		updates["warehouse_name"] = req.WarehouseName
	}
	if req.AreaName != "" {
		updates["area_name"] = req.AreaName
	}
	if req.AreaType != "" {
		updates["area_type"] = req.AreaType
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	return s.repo.Update(ctx, id, updates)
}

// Delete 删除库区
func (s *WmsAreaService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

// GetByID 根据ID获取库区
func (s *WmsAreaService) GetByID(ctx context.Context, id uint64) (*model.WmsArea, error) {
	return s.repo.GetByID(ctx, id)
}

// GetByCode 根据编码获取库区
func (s *WmsAreaService) GetByCode(ctx context.Context, tenantID int64, areaCode string) (*model.WmsArea, error) {
	return s.repo.GetByCode(ctx, tenantID, areaCode)
}

// Page 分页查询库区
func (s *WmsAreaService) Page(ctx context.Context, tenantID int64, query *model.WmsAreaQueryVO) ([]model.WmsArea, int64, error) {
	return s.repo.Page(ctx, tenantID, query)
}

// ListByWarehouse 按仓库获取库区列表
func (s *WmsAreaService) ListByWarehouse(ctx context.Context, tenantID int64, warehouseCode string) ([]model.WmsArea, error) {
	return s.repo.ListByWarehouse(ctx, tenantID, warehouseCode)
}

// GetTree 获取库区树形结构
func (s *WmsAreaService) GetTree(ctx context.Context, tenantID int64) ([]model.WmsAreaTreeVO, error) {
	allAreas, err := s.repo.ListAll(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	// 构建树形结构
	return buildAreaTree(allAreas), nil
}

// buildAreaTree 构建库区树
func buildAreaTree(areas []model.WmsArea) []model.WmsAreaTreeVO {
	// 建立映射
	areaMap := make(map[string]*model.WmsAreaTreeVO)
	for i := range areas {
		areaMap[areas[i].AreaCode] = &model.WmsAreaTreeVO{
			WmsArea: areas[i],
			Children: []model.WmsAreaTreeVO{},
		}
	}

	// 构建父子关系
	var roots []model.WmsAreaTreeVO
	for _, area := range areas {
		node := areaMap[area.AreaCode]
		if area.ParentCode == "" {
			roots = append(roots, *node)
		} else {
			if parent, ok := areaMap[area.ParentCode]; ok {
				parent.Children = append(parent.Children, *node)
			} else {
				// 如果父节点不存在，作为根节点
				roots = append(roots, *node)
			}
		}
	}

	return roots
}
