package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"time"

	"gorm.io/gorm"
)

type ProductionCompleteService struct {
	db             *gorm.DB
	completeRepo  *repository.ProductionCompleteRepository
	stockInRepo   *repository.ProductionStockInRepository
	orderRepo     *repository.ProductionOrderRepository
	inventoryRepo *repository.InventoryRepository
}

func NewProductionCompleteService(
	db *gorm.DB,
	completeRepo *repository.ProductionCompleteRepository,
	stockInRepo *repository.ProductionStockInRepository,
	orderRepo *repository.ProductionOrderRepository,
	inventoryRepo *repository.InventoryRepository,
) *ProductionCompleteService {
	return &ProductionCompleteService{
		db:             db,
		completeRepo:  completeRepo,
		stockInRepo:   stockInRepo,
		orderRepo:     orderRepo,
		inventoryRepo: inventoryRepo,
	}
}

func (s *ProductionCompleteService) GenerateCompleteNo(ctx context.Context, tenantID int64) (string, error) {
	var count int64
	today := time.Now().Format("2006-01-02")
	s.db.WithContext(ctx).Model(&model.ProductionComplete{}).Where("DATE(created_at) = ?", today).Count(&count)
	return fmt.Sprintf("PC%s%04d", today[2:10], count+1), nil
}

func (s *ProductionCompleteService) Create(ctx context.Context, tenantID int64, req *model.ProductionCompleteCreate) (*model.ProductionComplete, error) {
	completeNo, err := s.GenerateCompleteNo(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	complete := &model.ProductionComplete{
		CompleteNo:        completeNo,
		ProductionOrderID: req.ProductionOrderID,
		WorkshopID:       req.WorkshopID,
		WorkstationID:    req.WorkstationID,
		CompleteQty:      req.CompleteQty,
		QualifiedQty:     req.QualifiedQty,
		Status:          "PENDING",
		CompleteTime:    nil,
		TenantID:        tenantID,
	}

	for i, item := range req.Items {
		complete.Items = append(complete.Items, model.ProductionCompleteItem{
			CompleteID:   0, // will be set after create
			LineNo:       i + 1,
			MaterialID:   item.MaterialID,
			MaterialCode: item.MaterialCode,
			MaterialName: item.MaterialName,
			Unit:         item.Unit,
			CompleteQty:  item.CompleteQty,
			QualifiedQty: item.QualifiedQty,
			WarehouseID: item.WarehouseID,
			LocationID:  item.LocationID,
			BatchNo:     item.BatchNo,
			TenantID:    tenantID,
		})
	}

	if err := s.completeRepo.Create(ctx, complete); err != nil {
		return nil, err
	}
	return complete, nil
}

func (s *ProductionCompleteService) SubmitForInspect(ctx context.Context, id uint, req *model.ProductionCompleteSubmitForInspect) error {
	return s.completeRepo.Update(ctx, id, map[string]any{
		"status": "INSPECTING",
	})
}

func (s *ProductionCompleteService) Qualify(ctx context.Context, id uint, req *model.ProductionCompleteSubmitForInspect) error {
	var totalQualified float64
	for _, item := range req.Items {
		totalQualified += item.QualifiedQty
	}
	return s.completeRepo.Update(ctx, id, map[string]any{
		"status":       "QUALIFIED",
		"qualified_qty": totalQualified,
	})
}

func (s *ProductionCompleteService) StockIn(ctx context.Context, tenantID int64, req *model.ProductionStockInCreate) (*model.ProductionStockIn, error) {
	// 获取完工单
	complete, err := s.completeRepo.GetByID(ctx, uint(req.CompleteID))
	if err != nil {
		return nil, fmt.Errorf("完工单不存在: %w", err)
	}

	if complete.Status != "QUALIFIED" {
		return nil, fmt.Errorf("完工单状态必须为已质检，当前状态: %s", complete.Status)
	}

	// 获取关联工单，获取产品信息用于库存更新
	var productID int64
	var productCode, productName string
	if complete.ProductionOrderID > 0 {
		order, err := s.orderRepo.GetByID(ctx, uint(complete.ProductionOrderID))
		if err == nil && order != nil {
			productID = order.MaterialID
			productCode = order.MaterialCode
			productName = order.MaterialName
		}
	}

	// 生成入库单号
	stockInNo, _ := s.stockInRepo.GenerateStockInNo(ctx, tenantID)

	now := time.Now()
	stockIn := &model.ProductionStockIn{
		StockInNo:   stockInNo,
		CompleteID:  req.CompleteID,
		CompleteNo:  &complete.CompleteNo,
		WarehouseID: req.WarehouseID,
		LocationID:  req.LocationID,
		Status:      "PENDING",
		TenantID:    tenantID,
	}

	// 优先使用前端传入的入库明细，否则用完工明细
	if len(req.Items) > 0 {
		for i, item := range req.Items {
			stockIn.Items = append(stockIn.Items, model.ProductionStockInItem{
				StockInID:    0,
				LineNo:       i + 1,
				MaterialID:   item.MaterialID,
				MaterialCode: item.MaterialCode,
				MaterialName: item.MaterialName,
				Unit:         item.Unit,
				StockInQty:   item.StockInQty,
				WarehouseID:  item.WarehouseID,
				LocationID:   item.LocationID,
				BatchNo:      item.BatchNo,
				TenantID:     tenantID,
			})
		}
	} else {
		// 使用完工明细中的合格数量作为入库数量
		for i, item := range complete.Items {
			stockIn.Items = append(stockIn.Items, model.ProductionStockInItem{
				StockInID:    0,
				LineNo:       i + 1,
				MaterialID:   item.MaterialID,
				MaterialCode: item.MaterialCode,
				MaterialName: item.MaterialName,
				Unit:         item.Unit,
				StockInQty:   item.QualifiedQty,
				WarehouseID:  item.WarehouseID,
				LocationID:   item.LocationID,
				BatchNo:      item.BatchNo,
				TenantID:     tenantID,
			})
		}
	}

	// 入库单创建（事务保证）
	if err := s.stockInRepo.Create(ctx, stockIn); err != nil {
		return nil, err
	}

	// 更新入库单状态为已入库
	s.stockInRepo.Update(ctx, stockIn.ID, map[string]any{
		"status":       "STORED",
		"stock_in_time": now,
	})

	// 更新完工单状态
	s.completeRepo.Update(ctx, uint(req.CompleteID), map[string]any{
		"status":        "STORED",
		"complete_time": now,
	})

	// 更新工单状态为已完成
	if complete.ProductionOrderID > 0 {
		s.orderRepo.Update(ctx, uint(complete.ProductionOrderID), map[string]any{
			"status":          3, // 已完成
			"actual_end_date": now,
			"completed_qty":   complete.QualifiedQty,
		})
	}

	// 更新入库成品库存（使用工单的产品信息）
	if productID > 0 && complete.QualifiedQty > 0 {
		// 查找是否有现成库存记录
		inv, err := s.inventoryRepo.GetByMaterialAndLocation(ctx, productID, req.WarehouseID)
		if err == nil && inv.ID > 0 {
			// 存在则累加
			s.inventoryRepo.Update(ctx, uint(inv.ID), map[string]any{
				"quantity":      inv.Quantity + complete.QualifiedQty,
				"available_qty": inv.AvailableQty + complete.QualifiedQty,
			})
		} else {
			// 不存在则新建
			newInv := &model.Inventory{
				TenantID:     tenantID,
				MaterialID:   productID,
				MaterialCode: productCode,
				MaterialName: productName,
				WarehouseID:  req.WarehouseID,
				LocationID:   0,
				Quantity:     complete.QualifiedQty,
				AvailableQty: complete.QualifiedQty,
			}
			if req.LocationID != nil {
				newInv.LocationID = *req.LocationID
			}
			if len(stockIn.Items) > 0 && stockIn.Items[0].BatchNo != nil {
				newInv.BatchNo = stockIn.Items[0].BatchNo
			}
			s.inventoryRepo.Create(ctx, newInv)
		}
	}

	return stockIn, nil
}

func (s *ProductionCompleteService) List(ctx context.Context, tenantID int64, query string, page, pageSize int) ([]model.ProductionComplete, int64, error) {
	return s.completeRepo.List(ctx, tenantID, query, page, pageSize)
}

func (s *ProductionCompleteService) GetByID(ctx context.Context, id string) (*model.ProductionComplete, error) {
	var itemID uint
	fmt.Sscanf(id, "%d", &itemID)
	return s.completeRepo.GetByID(ctx, itemID)
}

func (s *ProductionCompleteService) ListStockIn(ctx context.Context, tenantID int64, query string, page, pageSize int) ([]model.ProductionStockIn, int64, error) {
	return s.stockInRepo.List(ctx, tenantID, query, page, pageSize)
}
