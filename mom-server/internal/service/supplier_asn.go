package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// SupplierASNService ASN服务
type SupplierASNService struct {
	asnRepo      *repository.SupplierASNRepository
	supplierRepo *repository.SupplierRepository
	materialRepo *repository.MaterialRepository
}

// NewSupplierASNService 创建ASN服务
func NewSupplierASNService(
	asnRepo *repository.SupplierASNRepository,
	supplierRepo *repository.SupplierRepository,
	materialRepo *repository.MaterialRepository,
) *SupplierASNService {
	return &SupplierASNService{
		asnRepo:      asnRepo,
		supplierRepo: supplierRepo,
		materialRepo: materialRepo,
	}
}

// Create 创建ASN
func (s *SupplierASNService) Create(ctx context.Context, req *model.CreateSupplierASNRequest) (*model.SupplierASN, error) {
	// 生成ASN编号
	asnNo, err := s.asnRepo.GenerateASNNo(ctx)
	if err != nil {
		return nil, fmt.Errorf("生成ASN编号失败: %w", err)
	}

	asn := &model.SupplierASN{
		ASNNo:         asnNo,
		SupplierID:    req.SupplierID,
		SupplierCode:  req.SupplierCode,
		SupplierName:  req.SupplierName,
		DeliveryType:  req.DeliveryType,
		WarehouseCode: req.WarehouseCode,
		ContactPerson: req.ContactPerson,
		ContactPhone:  req.ContactPhone,
		Status:        model.ASNStatus.Draft,
		Remark:        req.Remark,
		TenantID:      req.TenantID,
	}

	if req.DeliveryDate != "" {
		t, _ := time.Parse("2006-01-02", req.DeliveryDate)
		asn.DeliveryDate = &t
	}

	// 计算总数量和金额
	for i, item := range req.Items {
		asn.TotalQty += item.Qty
		asn.TotalAmount += item.Amount
		if item.Price > 0 && item.Qty > 0 && item.Amount == 0 {
			asn.TotalAmount += item.Price * item.Qty
		}
		_ = i // avoid unused
	}

	if err := s.asnRepo.Create(ctx, asn); err != nil {
		return nil, fmt.Errorf("创建ASN失败: %w", err)
	}

	// 创建明细
	for i, item := range req.Items {
		asnItem := &model.SupplierASNItem{
			ASNID:        asn.ID,
			LineNo:       i + 1,
			MaterialCode: item.MaterialCode,
			MaterialName: item.MaterialName,
			Spec:         item.Spec,
			Unit:         item.Unit,
			BatchNo:      item.BatchNo,
			Qty:          item.Qty,
			Price:        item.Price,
			Amount:       item.Amount,
			PackingQty:   item.PackingQty,
			PackingUnit:  item.PackingUnit,
			TenantID:     req.TenantID,
		}
		if item.Price > 0 && item.Qty > 0 && item.Amount == 0 {
			asnItem.Amount = item.Price * item.Qty
		}
		s.asnRepo.CreateItem(ctx, asnItem)
	}

	return asn, nil
}

// GetByID 根据ID获取ASN
func (s *SupplierASNService) GetByID(ctx context.Context, id int64) (*model.SupplierASN, error) {
	asn, err := s.asnRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("ASN不存在: %w", err)
	}

	items, _ := s.asnRepo.GetItemsByASNID(ctx, id)
	asn.Items = items
	return asn, nil
}

// GetByASNNo 根据编号获取ASN
func (s *SupplierASNService) GetByASNNo(ctx context.Context, asnNo string) (*model.SupplierASN, error) {
	asn, err := s.asnRepo.GetByASNNo(ctx, asnNo)
	if err != nil {
		return nil, fmt.Errorf("ASN不存在: %w", err)
	}

	items, _ := s.asnRepo.GetItemsByASNID(ctx, asn.ID)
	asn.Items = items
	return asn, nil
}

// List 查询ASN列表
func (s *SupplierASNService) List(ctx context.Context, q *model.SupplierASNQuery) ([]model.SupplierASN, int64, error) {
	return s.asnRepo.List(ctx, q)
}

// ListBySupplier 按供应商查询
func (s *SupplierASNService) ListBySupplier(ctx context.Context, tenantID, supplierID int64, status string) ([]model.SupplierASN, error) {
	return s.asnRepo.ListBySupplier(ctx, tenantID, supplierID, status)
}

// Update 更新ASN
func (s *SupplierASNService) Update(ctx context.Context, id int64, req *model.UpdateSupplierASNRequest) error {
	updates := map[string]interface{}{}

	if req.DeliveryType != "" {
		updates["delivery_type"] = req.DeliveryType
	}
	if req.DeliveryDate != "" {
		t, _ := time.Parse("2006-01-02", req.DeliveryDate)
		updates["delivery_date"] = &t
	}
	if req.DeliveryStart != "" {
		updates["delivery_time_start"] = req.DeliveryStart
	}
	if req.DeliveryEnd != "" {
		updates["delivery_time_end"] = req.DeliveryEnd
	}
	if req.WarehouseCode != "" {
		updates["warehouse_code"] = req.WarehouseCode
	}
	if req.ContactPerson != "" {
		updates["contact_person"] = req.ContactPerson
	}
	if req.ContactPhone != "" {
		updates["contact_phone"] = req.ContactPhone
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	return s.asnRepo.Update(ctx, id, updates)
}

// UpdateStatus 更新状态
func (s *SupplierASNService) UpdateStatus(ctx context.Context, id int64, status string) error {
	return s.asnRepo.UpdateStatus(ctx, id, status)
}

// Submit 提交ASN
func (s *SupplierASNService) Submit(ctx context.Context, id int64) error {
	asn, err := s.asnRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("ASN不存在: %w", err)
	}

	if asn.Status != model.ASNStatus.Draft {
		return fmt.Errorf("当前状态 %s 不允许提交", asn.Status)
	}

	return s.asnRepo.UpdateStatus(ctx, id, model.ASNStatus.Submitted)
}

// Confirm 确认ASN
func (s *SupplierASNService) Confirm(ctx context.Context, id int64, req *model.ConfirmASNRequest) error {
	asn, err := s.asnRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("ASN不存在: %w", err)
	}

	if asn.Status != model.ASNStatus.Submitted {
		return fmt.Errorf("当前状态 %s 不允许确认", asn.Status)
	}

	if req.ConfirmStatus == "REJECTED" {
		return s.asnRepo.UpdateStatus(ctx, id, model.ASNStatus.Draft)
	}

	// 应用修改的明细
	for _, item := range req.ModifiedItems {
		s.asnRepo.UpdateItem(ctx, item.ItemID, map[string]interface{}{
			"qty":   item.Qty,
			"price": item.Price,
		})
	}

	return s.asnRepo.UpdateStatus(ctx, id, model.ASNStatus.Confirmed)
}

// StartReceiving 开始收货
func (s *SupplierASNService) StartReceiving(ctx context.Context, id int64) error {
	asn, err := s.asnRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("ASN不存在: %w", err)
	}

	if asn.Status != model.ASNStatus.Confirmed {
		return fmt.Errorf("当前状态 %s 不允许开始收货", asn.Status)
	}

	return s.asnRepo.UpdateStatus(ctx, id, model.ASNStatus.Receiving)
}

// CompleteReceiving 完成收货
func (s *SupplierASNService) CompleteReceiving(ctx context.Context, id int64) error {
	asn, err := s.asnRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("ASN不存在: %w", err)
	}

	if asn.Status != model.ASNStatus.Receiving {
		return fmt.Errorf("当前状态 %s 不允许完成收货", asn.Status)
	}

	return s.asnRepo.UpdateStatus(ctx, id, model.ASNStatus.Completed)
}

// Cancel 取消ASN
func (s *SupplierASNService) Cancel(ctx context.Context, id int64) error {
	asn, err := s.asnRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("ASN不存在: %w", err)
	}

	if asn.Status == model.ASNStatus.Completed || asn.Status == model.ASNStatus.Cancelled {
		return fmt.Errorf("当前状态 %s 不允许取消", asn.Status)
	}

	return s.asnRepo.UpdateStatus(ctx, id, model.ASNStatus.Cancelled)
}

// Delete 删除ASN
func (s *SupplierASNService) Delete(ctx context.Context, id int64) error {
	asn, err := s.asnRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("ASN不存在: %w", err)
	}

	if asn.Status != model.ASNStatus.Draft {
		return fmt.Errorf("只有草稿状态可以删除")
	}

	return s.asnRepo.Delete(ctx, id)
}

// AddItem 添加ASN明细
func (s *SupplierASNService) AddItem(ctx context.Context, asnID int64, item *model.CreateASNItemRequest) (*model.SupplierASNItem, error) {
	asn, err := s.asnRepo.GetByID(ctx, asnID)
	if err != nil {
		return nil, fmt.Errorf("ASN不存在: %w", err)
	}

	if asn.Status != model.ASNStatus.Draft {
		return nil, fmt.Errorf("只有草稿状态可以添加明细")
	}

	items, _ := s.asnRepo.GetItemsByASNID(ctx, asnID)
	lineNo := len(items) + 1

	asnItem := &model.SupplierASNItem{
		ASNID:        asnID,
		LineNo:       lineNo,
		MaterialCode: item.MaterialCode,
		MaterialName: item.MaterialName,
		Spec:         item.Spec,
		Unit:         item.Unit,
		BatchNo:      item.BatchNo,
		Qty:          item.Qty,
		Price:        item.Price,
		Amount:       item.Amount,
		PackingQty:   item.PackingQty,
		PackingUnit:  item.PackingUnit,
		TenantID:     asn.TenantID,
	}
	if item.Price > 0 && item.Qty > 0 && item.Amount == 0 {
		asnItem.Amount = item.Price * item.Qty
	}

	if err := s.asnRepo.CreateItem(ctx, asnItem); err != nil {
		return nil, fmt.Errorf("添加明细失败: %w", err)
	}

	return asnItem, nil
}
