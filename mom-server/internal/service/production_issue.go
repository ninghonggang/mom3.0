package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// ProductionIssueService 生产发料服务
type ProductionIssueService struct {
	issueRepo *repository.ProductionIssueRepository
	itemRepo  *repository.ProductionIssueItemRepository
	invRepo   *repository.InventoryRepository
}

func NewProductionIssueService(
	issueRepo *repository.ProductionIssueRepository,
	itemRepo *repository.ProductionIssueItemRepository,
	invRepo *repository.InventoryRepository,
) *ProductionIssueService {
	return &ProductionIssueService{
		issueRepo: issueRepo,
		itemRepo:  itemRepo,
		invRepo:   invRepo,
	}
}

// List 查询发料单列表
func (s *ProductionIssueService) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.ProductionIssue, int64, error) {
	return s.issueRepo.List(ctx, tenantID, query)
}

// GetByID 获取发料单详情
func (s *ProductionIssueService) GetByID(ctx context.Context, id uint) (*model.ProductionIssue, error) {
	return s.issueRepo.GetByID(ctx, id)
}

// Create 创建发料单
func (s *ProductionIssueService) Create(ctx context.Context, tenantID int64, req *model.ProductionIssueCreate, username string) (*model.ProductionIssue, error) {
	now := time.Now()
	seq := now.UnixNano() % 10000
	issueNo := fmt.Sprintf("PI%s%04d", now.Format("20060102"), seq)

	issue := &model.ProductionIssue{
		IssueNo:           issueNo,
		IssueType:         req.IssueType,
		ProductionOrderID: req.ProductionOrderID,
		WorkstationID:     req.WorkstationID,
		WorkshopID:         req.WorkshopID,
		Status:             "PENDING",
		PickStatus:         "PENDING",
		RequestBy:          nil,
		RequestTime:        &now,
		TenantID:           tenantID,
		CreatedBy:         &username,
	}

	if err := s.issueRepo.Create(ctx, issue); err != nil {
		return nil, err
	}

	// 创建明细
	for i, item := range req.Items {
		pi := model.ProductionIssueItem{
			IssueID:      uint(issue.ID),
			LineNo:       i + 1,
			MaterialID:   item.MaterialID,
			MaterialCode: item.MaterialCode,
			MaterialName: item.MaterialName,
			Unit:         item.Unit,
			RequiredQty:  item.RequiredQty,
			PickedQty:    0,
			IssuedQty:    0,
			WarehouseID:  item.WarehouseID,
			LocationID:   item.LocationID,
			BatchNo:      item.BatchNo,
			Remark:       item.Remark,
			TenantID:     tenantID,
		}
		if err := s.itemRepo.CreateBatch(ctx, []model.ProductionIssueItem{pi}); err != nil {
			return nil, err
		}
	}

	return s.issueRepo.GetByID(ctx, issue.ID)
}

// Update 更新发料单（仅PENDING状态）
func (s *ProductionIssueService) Update(ctx context.Context, id uint, req *model.ProductionIssueUpdate) error {
	issue, err := s.issueRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if issue.Status != "PENDING" {
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

	return s.issueRepo.Update(ctx, id, updates)
}

// Delete 删除发料单（仅PENDING状态）
func (s *ProductionIssueService) Delete(ctx context.Context, id uint) error {
	issue, err := s.issueRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if issue.Status != "PENDING" {
		return fmt.Errorf("只有待提交状态可以删除")
	}
	return s.issueRepo.Delete(ctx, id)
}

// Submit 提交发料单（PENDING → APPROVED）
func (s *ProductionIssueService) Submit(ctx context.Context, id uint, req *model.ProductionIssueSubmit) error {
	issue, err := s.issueRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if issue.Status != "PENDING" {
		return fmt.Errorf("当前状态不允许提交")
	}

	// 更新明细的拣配数量
	for _, item := range req.Items {
		items, _ := s.itemRepo.ListByIssueID(ctx, int64(id))
		for _, existing := range items {
			if existing.MaterialID == item.MaterialID {
				s.itemRepo.UpdatePickedQty(ctx, existing.ID, item.PickedQty)
				break
			}
		}
	}

	return s.issueRepo.Update(ctx, id, map[string]interface{}{
		"status": "APPROVED",
	})
}

// StartPick 开始拣配（APPROVED → PICKING）
func (s *ProductionIssueService) StartPick(ctx context.Context, id uint) error {
	issue, err := s.issueRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if issue.Status != "APPROVED" {
		return fmt.Errorf("当前状态不允许开始拣配")
	}
	return s.issueRepo.Update(ctx, id, map[string]interface{}{
		"status":      "PICKING",
		"pick_status": "PICKING",
	})
}

// ConfirmPick 确认拣配（PICKING → PICKED）
func (s *ProductionIssueService) ConfirmPick(ctx context.Context, id uint, items []model.ProductionIssueItemSubmit) error {
	issue, err := s.issueRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if issue.Status != "PICKING" {
		return fmt.Errorf("当前状态不允许确认拣配")
	}

	// 更新每个物料的已拣配数量并预占库存
	warehouseID := issue.WorkshopID
	if warehouseID == nil {
		return fmt.Errorf("工单车间ID为空")
	}
	for _, item := range items {
		itemList, _ := s.itemRepo.ListByIssueID(ctx, int64(id))
		for _, existing := range itemList {
			if existing.MaterialID == item.MaterialID {
				s.itemRepo.UpdatePickedQty(ctx, existing.ID, item.PickedQty)
				s.invRepo.Allocate(ctx, existing.MaterialID, *warehouseID, item.PickedQty)
				break
			}
		}
	}

	return s.issueRepo.Update(ctx, id, map[string]interface{}{
		"status":      "PICKED",
		"pick_status": "PICKED",
	})
}

// Issue 发料（PICKED → ISSUED，扣减库存）
func (s *ProductionIssueService) Issue(ctx context.Context, id uint, issuedBy int64) error {
	issue, err := s.issueRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if issue.Status != "PICKED" {
		return fmt.Errorf("当前状态不允许发料")
	}

	warehouseID := issue.WorkshopID
	if warehouseID == nil {
		return fmt.Errorf("工单车间ID为空")
	}
	itemList, _ := s.itemRepo.ListByIssueID(ctx, int64(id))
	for _, item := range itemList {
		s.invRepo.DeductAllocated(ctx, item.MaterialID, *warehouseID, item.IssuedQty)
	}

	now := time.Now()
	return s.issueRepo.Update(ctx, id, map[string]interface{}{
		"status":      "ISSUED",
		"issued_by":   issuedBy,
		"issued_time": now,
	})
}

// Cancel 取消发料单（释放预占库存）
func (s *ProductionIssueService) Cancel(ctx context.Context, id uint) error {
	issue, err := s.issueRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if issue.Status == "ISSUED" {
		return fmt.Errorf("已发料状态不允许取消")
	}

	if issue.Status == "PICKED" {
		warehouseID := issue.WorkshopID
		if warehouseID == nil {
			return fmt.Errorf("工单车间ID为空")
		}
		itemList, _ := s.itemRepo.ListByIssueID(ctx, int64(id))
		for _, item := range itemList {
			s.invRepo.ReleaseAllocation(ctx, item.MaterialID, *warehouseID, item.IssuedQty)
		}
	}

	return s.issueRepo.Update(ctx, id, map[string]interface{}{
		"status": "CANCELLED",
	})
}
