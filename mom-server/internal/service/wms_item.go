package service

import (
	"context"
	"fmt"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// WMSItemService 货品管理服务层
type WMSItemService struct {
	repo *repository.WMSItemRepository
}

// NewWMSItemService 创建货品服务实例
func NewWMSItemService(repo *repository.WMSItemRepository) *WMSItemService {
	return &WMSItemService{repo: repo}
}

// List 获取货品列表
func (s *WMSItemService) List(ctx context.Context, query *model.WMSItemQueryVO) ([]model.WMSItem, int64, error) {
	// 设置默认分页
	if query != nil {
		if query.Page <= 0 {
			query.Page = 1
		}
		if query.PageSize <= 0 {
			query.PageSize = 20
		}
	}
	return s.repo.List(ctx, query)
}

// Get 获取货品详情
func (s *WMSItemService) Get(ctx context.Context, id string) (*model.WMSItem, error) {
	var itemID uint
	_, err := fmt.Sscanf(id, "%d", &itemID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, itemID)
}

// Search 搜索货品
func (s *WMSItemService) Search(ctx context.Context, keyword string) ([]model.WMSItem, error) {
	return s.repo.Search(ctx, keyword)
}

// Create 创建货品
func (s *WMSItemService) Create(ctx context.Context, req *model.WMSItemCreateReqVO) error {
	item := &model.WMSItem{
		ItemCode:      req.ItemCode,
		ItemName:      req.ItemName,
		Specification: req.Specification,
		Unit:          req.Unit,
		ItemType:      req.ItemType,
		CategoryID:    req.CategoryID,
		Barcode:       req.Barcode,
		SafetyStock:   req.SafetyStock,
		Status:        "ACTIVE",
	}
	return s.repo.Create(ctx, item)
}

// Update 更新货品
func (s *WMSItemService) Update(ctx context.Context, id string, req *model.WMSItemUpdateReqVO) error {
	var itemID uint
	_, err := fmt.Sscanf(id, "%d", &itemID)
	if err != nil {
		return err
	}
	updates := map[string]interface{}{
		"item_name":     req.ItemName,
		"specification": req.Specification,
		"unit":          req.Unit,
		"category_id":   req.CategoryID,
		"barcode":       req.Barcode,
		"safety_stock":  req.SafetyStock,
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}
	return s.repo.Update(ctx, itemID, updates)
}

// Delete 删除货品
func (s *WMSItemService) Delete(ctx context.Context, id string) error {
	var itemID uint
	_, err := fmt.Sscanf(id, "%d", &itemID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, itemID)
}
