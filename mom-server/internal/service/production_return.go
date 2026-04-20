package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// ProductionReturnService 生产退料服务
type ProductionReturnService struct {
	returnRepo *repository.ProductionReturnRepository
	itemRepo   *repository.ProductionReturnItemRepository
}

func NewProductionReturnService(
	returnRepo *repository.ProductionReturnRepository,
	itemRepo *repository.ProductionReturnItemRepository,
) *ProductionReturnService {
	return &ProductionReturnService{
		returnRepo: returnRepo,
		itemRepo:   itemRepo,
	}
}

// List 查询退料单列表
func (s *ProductionReturnService) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.ProductionReturn, int64, error) {
	return s.returnRepo.List(ctx, tenantID, query)
}

// GetByID 获取退料单详情
func (s *ProductionReturnService) GetByID(ctx context.Context, id uint) (*model.ProductionReturn, error) {
	return s.returnRepo.GetByID(ctx, id)
}

// Create 创建退料单
func (s *ProductionReturnService) Create(ctx context.Context, tenantID int64, req *model.ProductionReturnCreate, username string) (*model.ProductionReturn, error) {
	now := time.Now()
	seq := now.UnixNano() % 10000
	returnNo := fmt.Sprintf("PR%s%04d", now.Format("20060102"), seq)

	ret := &model.ProductionReturn{
		ReturnNo:           returnNo,
		ProductionOrderID:  req.ProductionOrderID,
		ReturnType:         req.ReturnType,
		WorkstationID:      req.WorkstationID,
		WorkshopID:         req.WorkshopID,
		Status:             "PENDING",
		RequestBy:          nil,
		RequestTime:        &now,
		TenantID:           tenantID,
		CreatedBy:          &username,
	}

	if err := s.returnRepo.Create(ctx, ret); err != nil {
		return nil, err
	}

	// 创建明细
	for i, item := range req.Items {
		ri := model.ProductionReturnItem{
			ReturnID:     uint(ret.ID),
			LineNo:       i + 1,
			MaterialID:   item.MaterialID,
			MaterialCode: item.MaterialCode,
			MaterialName: item.MaterialName,
			Unit:         item.Unit,
			IssuedQty:    item.IssuedQty,
			ReturnQty:    item.ReturnQty,
			WarehouseID:  item.WarehouseID,
			LocationID:   item.LocationID,
			BatchNo:      item.BatchNo,
			Reason:       item.Reason,
			Remark:       item.Remark,
			TenantID:     tenantID,
		}
		if err := s.itemRepo.CreateBatch(ctx, []model.ProductionReturnItem{ri}); err != nil {
			return nil, err
		}
	}

	return s.returnRepo.GetByID(ctx, ret.ID)
}

// Update 更新退料单（仅PENDING状态）
func (s *ProductionReturnService) Update(ctx context.Context, id uint, req *model.ProductionReturnUpdate) error {
	ret, err := s.returnRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if ret.Status != "PENDING" {
		return fmt.Errorf("只有待提交状态可以编辑")
	}

	updates := map[string]interface{}{}
	if req.WorkstationID != nil {
		updates["workstation_id"] = req.WorkstationID
	}
	if req.WorkshopID != nil {
		updates["workshop_id"] = req.WorkshopID
	}
	if req.Remark != nil {
		updates["remark"] = req.Remark
	}

	return s.returnRepo.Update(ctx, id, updates)
}

// Delete 删除退料单（仅PENDING状态）
func (s *ProductionReturnService) Delete(ctx context.Context, id uint) error {
	ret, err := s.returnRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if ret.Status != "PENDING" {
		return fmt.Errorf("只有待提交状态可以删除")
	}
	return s.returnRepo.Delete(ctx, id)
}

// Submit 提交退料单（PENDING → APPROVED）
func (s *ProductionReturnService) Submit(ctx context.Context, id uint, req *model.ProductionReturnSubmit) error {
	ret, err := s.returnRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if ret.Status != "PENDING" {
		return fmt.Errorf("当前状态不允许提交")
	}

	// 更新明细的退料数量
	for _, item := range req.Items {
		items, _ := s.itemRepo.ListByReturnID(ctx, int64(id))
		for _, existing := range items {
			if existing.MaterialID == item.MaterialID {
				s.itemRepo.UpdateReturnQty(ctx, existing.ID, item.ReturnQty)
				break
			}
		}
	}

	return s.returnRepo.Update(ctx, id, map[string]interface{}{
		"status": "APPROVED",
	})
}

// Approve 审批退料单（APPROVED 确认）
func (s *ProductionReturnService) Approve(ctx context.Context, id uint, approvedBy int64) error {
	ret, err := s.returnRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if ret.Status != "APPROVED" {
		return fmt.Errorf("当前状态不允许审批")
	}

	now := time.Now()
	return s.returnRepo.Update(ctx, id, map[string]interface{}{
		"status":       "RETURNING",
		"approved_by":  approvedBy,
		"approved_time": now,
	})
}

// StartReturn 开始退料（APPROVED → RETURNING）
func (s *ProductionReturnService) StartReturn(ctx context.Context, id uint) error {
	ret, err := s.returnRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if ret.Status != "APPROVED" {
		return fmt.Errorf("当前状态不允许开始退料")
	}
	return s.returnRepo.Update(ctx, id, map[string]interface{}{
		"status": "RETURNING",
	})
}

// ConfirmReturn 确认退料（RETURNING → RETURNED）
func (s *ProductionReturnService) ConfirmReturn(ctx context.Context, id uint, items []model.ProductionReturnItemConfirm) error {
	ret, err := s.returnRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if ret.Status != "RETURNING" {
		return fmt.Errorf("当前状态不允许确认退料")
	}

	// 更新每个物料的退料数量
	for _, item := range items {
		itemList, _ := s.itemRepo.ListByReturnID(ctx, int64(id))
		for _, existing := range itemList {
			if existing.MaterialID == item.MaterialID {
				s.itemRepo.UpdateReturnQty(ctx, existing.ID, item.ReturnQty)
				break
			}
		}
	}

	now := time.Now()
	return s.returnRepo.Update(ctx, id, map[string]interface{}{
		"status":        "RETURNED",
		"returned_by":   nil, // TODO: get from context
		"returned_time": now,
	})
}

// Cancel 取消退料单
func (s *ProductionReturnService) Cancel(ctx context.Context, id uint) error {
	ret, err := s.returnRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if ret.Status == "RETURNED" {
		return fmt.Errorf("已退料状态不允许取消")
	}
	return s.returnRepo.Update(ctx, id, map[string]interface{}{
		"status": "CANCELLED",
	})
}
