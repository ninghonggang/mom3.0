package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// SalesReturnService 销售退货服务
type SalesReturnService struct {
	returnRepo *repository.SalesReturnRepository
	itemRepo   *repository.SalesReturnItemRepository
}

func NewSalesReturnService(
	returnRepo *repository.SalesReturnRepository,
	itemRepo *repository.SalesReturnItemRepository,
) *SalesReturnService {
	return &SalesReturnService{
		returnRepo: returnRepo,
		itemRepo:   itemRepo,
	}
}

// List 查询退货单列表
func (s *SalesReturnService) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.SalesReturn, int64, error) {
	return s.returnRepo.List(ctx, tenantID, query)
}

// GetByID 获取退货单详情
func (s *SalesReturnService) GetByID(ctx context.Context, id uint) (*model.SalesReturn, error) {
	return s.returnRepo.GetByID(ctx, id)
}

// Create 创建退货单
func (s *SalesReturnService) Create(ctx context.Context, tenantID int64, req *model.SalesReturnCreate, username string) (*model.SalesReturn, error) {
	now := time.Now()
	seq := now.UnixNano() % 10000
	returnNo := fmt.Sprintf("SR%s%04d", now.Format("20060102"), seq)

	ret := &model.SalesReturn{
		ReturnNo:       returnNo,
		SalesOrderID:  req.SalesOrderID,
		CustomerID:    req.CustomerID,
		CustomerName:  req.CustomerName,
		WarehouseID:   req.WarehouseID,
		ReturnDate:   req.ReturnDate,
		ReturnType:   req.ReturnType,
		Status:        "PENDING",
		RequestBy:     nil,
		RequestTime:   &now,
		TenantID:      tenantID,
		CreatedBy:     &username,
		Remark:        req.Remark,
	}

	if err := s.returnRepo.Create(ctx, ret); err != nil {
		return nil, err
	}

	// 创建明细
	for i, item := range req.Items {
		ri := model.SalesReturnItem{
			ReturnID:     uint(ret.ID),
			LineNo:       i + 1,
			MaterialID:   item.MaterialID,
			MaterialCode: item.MaterialCode,
			MaterialName: item.MaterialName,
			Unit:         item.Unit,
			ReturnQty:    item.ReturnQty,
			Reason:       item.Reason,
			TenantID:     tenantID,
		}
		if err := s.itemRepo.CreateBatch(ctx, []model.SalesReturnItem{ri}); err != nil {
			return nil, err
		}
	}

	return s.returnRepo.GetByID(ctx, ret.ID)
}

// Update 更新退货单（仅PENDING状态）
func (s *SalesReturnService) Update(ctx context.Context, id uint, req *model.SalesReturnUpdate) error {
	ret, err := s.returnRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if ret.Status != "PENDING" {
		return fmt.Errorf("只有待提交状态可以编辑")
	}

	updates := map[string]interface{}{}
	if req.Remark != nil {
		updates["remark"] = req.Remark
	}

	return s.returnRepo.Update(ctx, id, updates)
}

// Delete 删除退货单（仅PENDING状态）
func (s *SalesReturnService) Delete(ctx context.Context, id uint) error {
	ret, err := s.returnRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if ret.Status != "PENDING" {
		return fmt.Errorf("只有待提交状态可以删除")
	}
	return s.returnRepo.Delete(ctx, id)
}

// Submit 提交退货单（PENDING → APPROVED）
func (s *SalesReturnService) Submit(ctx context.Context, id uint, req *model.SalesReturnSubmit) error {
	ret, err := s.returnRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if ret.Status != "PENDING" {
		return fmt.Errorf("当前状态不允许提交")
	}

	// 更新明细的退货数量
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

// Approve 审批退货单（APPROVED 确认）
func (s *SalesReturnService) Approve(ctx context.Context, id uint, approvedBy int64) error {
	ret, err := s.returnRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if ret.Status != "APPROVED" {
		return fmt.Errorf("当前状态不允许审批")
	}

	now := time.Now()
	return s.returnRepo.Update(ctx, id, map[string]interface{}{
		"status":        "RETURNING",
		"approved_by":   approvedBy,
		"approved_time": now,
	})
}

// StartReturn 开始退货（APPROVED → RETURNING）
func (s *SalesReturnService) StartReturn(ctx context.Context, id uint) error {
	ret, err := s.returnRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if ret.Status != "APPROVED" {
		return fmt.Errorf("当前状态不允许开始退货")
	}
	return s.returnRepo.Update(ctx, id, map[string]interface{}{
		"status": "RETURNING",
	})
}

// ConfirmReturn 确认退货（RETURNING → RETURNED）
func (s *SalesReturnService) ConfirmReturn(ctx context.Context, id uint, items []model.SalesReturnItemConfirm) error {
	ret, err := s.returnRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if ret.Status != "RETURNING" {
		return fmt.Errorf("当前状态不允许确认退货")
	}

	// 更新每个物料的退货数量
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

// Cancel 取消退货单
func (s *SalesReturnService) Cancel(ctx context.Context, id uint) error {
	ret, err := s.returnRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if ret.Status == "RETURNED" {
		return fmt.Errorf("已退货状态不允许取消")
	}
	return s.returnRepo.Update(ctx, id, map[string]interface{}{
		"status": "CANCELLED",
	})
}