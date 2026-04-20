package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// FinService 结算服务
type FinService struct {
	purchaseSettlementRepo *repository.PurchaseSettlementRepository
	salesSettlementRepo     *repository.SalesSettlementRepository
	paymentRequestRepo     *repository.PaymentRequestRepository
	purchaseAdvanceRepo     *repository.PurchaseAdvanceRepository
	salesReceiptRepo        *repository.SalesReceiptRepository
	supplierStatementRepo   *repository.SupplierStatementRepository
}

func NewFinService(
	purchaseSettlementRepo *repository.PurchaseSettlementRepository,
	salesSettlementRepo *repository.SalesSettlementRepository,
	paymentRequestRepo *repository.PaymentRequestRepository,
	purchaseAdvanceRepo *repository.PurchaseAdvanceRepository,
	salesReceiptRepo *repository.SalesReceiptRepository,
	supplierStatementRepo *repository.SupplierStatementRepository,
) *FinService {
	return &FinService{
		purchaseSettlementRepo: purchaseSettlementRepo,
		salesSettlementRepo:     salesSettlementRepo,
		paymentRequestRepo:     paymentRequestRepo,
		purchaseAdvanceRepo:     purchaseAdvanceRepo,
		salesReceiptRepo:       salesReceiptRepo,
		supplierStatementRepo:   supplierStatementRepo,
	}
}

// ==================== 采购结算 ====================

func (s *FinService) ListPurchaseSettlements(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.PurchaseSettlement, int64, error) {
	return s.purchaseSettlementRepo.List(ctx, tenantID, query)
}

func (s *FinService) GetPurchaseSettlement(ctx context.Context, id uint) (*model.PurchaseSettlement, error) {
	return s.purchaseSettlementRepo.GetByID(ctx, id)
}

func (s *FinService) CreatePurchaseSettlement(ctx context.Context, tenantID int64, req *model.PurchaseSettlementCreate, createdBy string) (*model.PurchaseSettlement, error) {
	// 生成结算单号
	now := time.Now()
	seq := now.UnixNano() % 10000
	settlementNo := fmt.Sprintf("PS%s%04d", now.Format("20060102"), seq)

	// 计算金额
	var goodsAmount, taxAmount, totalAmount float64
	for _, item := range req.Items {
		item.GoodsAmount = item.UnitPrice * item.ThisSettleQty
		item.TaxAmount = item.GoodsAmount * item.TaxRate / 100
		item.LineAmount = item.GoodsAmount + item.TaxAmount
		goodsAmount += item.GoodsAmount
		taxAmount += item.TaxAmount
	}
	totalAmount = goodsAmount + taxAmount

	settlement := &model.PurchaseSettlement{
		SettlementNo:   settlementNo,
		SettlementType: "NORMAL",
		RelatedType:    req.RelatedType,
		RelatedID:      req.RelatedID,
		SupplierID:     req.SupplierID,
		InvoiceNo:     req.InvoiceNo,
		InvoiceDate:   req.InvoiceDate,
		GoodsAmount:   goodsAmount,
		TaxAmount:     taxAmount,
		TotalAmount:   totalAmount,
		PaymentTerms:  req.PaymentTerms,
		PaymentDueDate: req.PaymentDueDate,
		PaymentMethod: req.PaymentMethod,
		Status:        "PENDING",
		Remark:        req.Remark,
		TenantID:      tenantID,
		CreatedBy:      &createdBy,
	}

	if err := s.purchaseSettlementRepo.Create(ctx, settlement); err != nil {
		return nil, err
	}

	// 创建明细
	for i, item := range req.Items {
		item.SettlementID = settlement.ID
		item.LineNo = i + 1
		item.TenantID = tenantID
		item.GoodsAmount = item.UnitPrice * item.ThisSettleQty
		item.TaxAmount = item.GoodsAmount * item.TaxRate / 100
		item.LineAmount = item.GoodsAmount + item.TaxAmount
		if err := s.purchaseSettlementRepo.CreateItem(ctx, &item); err != nil {
			return nil, err
		}
	}

	return settlement, nil
}

func (s *FinService) SubmitPurchaseSettlement(ctx context.Context, id uint) error {
	return s.purchaseSettlementRepo.Update(ctx, id, map[string]interface{}{
		"status": "SUBMITTED",
	})
}

func (s *FinService) ApprovePurchaseSettlement(ctx context.Context, id uint, approvedBy int64) error {
	now := time.Now()
	return s.purchaseSettlementRepo.Update(ctx, id, map[string]interface{}{
		"status":       "APPROVED",
		"approved_by":   approvedBy,
		"approved_time": now,
	})
}

func (s *FinService) CancelPurchaseSettlement(ctx context.Context, id uint) error {
	return s.purchaseSettlementRepo.Update(ctx, id, map[string]interface{}{
		"status": "CANCELLED",
	})
}

func (s *FinService) DeletePurchaseSettlement(ctx context.Context, id uint) error {
	// 先删除明细
	if err := s.purchaseSettlementRepo.DeleteItems(ctx, int64(id)); err != nil {
		return err
	}
	return s.purchaseSettlementRepo.Delete(ctx, id)
}

// ==================== 销售结算 ====================

func (s *FinService) ListSalesSettlements(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.SalesSettlement, int64, error) {
	return s.salesSettlementRepo.List(ctx, tenantID, query)
}

func (s *FinService) GetSalesSettlement(ctx context.Context, id uint) (*model.SalesSettlement, error) {
	return s.salesSettlementRepo.GetByID(ctx, id)
}

func (s *FinService) CreateSalesSettlement(ctx context.Context, tenantID int64, req *model.SalesSettlementCreate, createdBy string) (*model.SalesSettlement, error) {
	now := time.Now()
	seq := now.UnixNano() % 10000
	settlementNo := fmt.Sprintf("SS%s%04d", now.Format("20060102"), seq)

	var goodsAmount, taxAmount, totalAmount float64
	for _, item := range req.Items {
		item.GoodsAmount = item.UnitPrice * item.ThisSettleQty
		item.TaxAmount = item.GoodsAmount * item.TaxRate / 100
		item.LineAmount = item.GoodsAmount + item.TaxAmount
		goodsAmount += item.GoodsAmount
		taxAmount += item.TaxAmount
	}
	totalAmount = goodsAmount + taxAmount

	settlement := &model.SalesSettlement{
		SettlementNo:   settlementNo,
		SettlementType: "NORMAL",
		RelatedType:    req.RelatedType,
		RelatedID:      req.RelatedID,
		CustomerID:     req.CustomerID,
		InvoiceNo:     req.InvoiceNo,
		InvoiceDate:   req.InvoiceDate,
		GoodsAmount:   goodsAmount,
		TaxAmount:     taxAmount,
		TotalAmount:   totalAmount,
		PaymentTerms:  req.PaymentTerms,
		PaymentDueDate: req.PaymentDueDate,
		PaymentMethod: req.PaymentMethod,
		Status:        "PENDING",
		Remark:        req.Remark,
		TenantID:      tenantID,
		CreatedBy:      &createdBy,
	}

	if err := s.salesSettlementRepo.Create(ctx, settlement); err != nil {
		return nil, err
	}

	for i, item := range req.Items {
		item.SettlementID = settlement.ID
		item.LineNo = i + 1
		item.TenantID = tenantID
		item.GoodsAmount = item.UnitPrice * item.ThisSettleQty
		item.TaxAmount = item.GoodsAmount * item.TaxRate / 100
		item.LineAmount = item.GoodsAmount + item.TaxAmount
		if err := s.salesSettlementRepo.CreateItem(ctx, &item); err != nil {
			return nil, err
		}
	}

	return settlement, nil
}

func (s *FinService) SubmitSalesSettlement(ctx context.Context, id uint) error {
	return s.salesSettlementRepo.Update(ctx, id, map[string]interface{}{
		"status": "SUBMITTED",
	})
}

func (s *FinService) ApproveSalesSettlement(ctx context.Context, id uint, approvedBy int64) error {
	now := time.Now()
	return s.salesSettlementRepo.Update(ctx, id, map[string]interface{}{
		"status":       "APPROVED",
		"approved_by":   approvedBy,
		"approved_time": now,
	})
}

func (s *FinService) CancelSalesSettlement(ctx context.Context, id uint) error {
	return s.salesSettlementRepo.Update(ctx, id, map[string]interface{}{
		"status": "CANCELLED",
	})
}

func (s *FinService) DeleteSalesSettlement(ctx context.Context, id uint) error {
	if err := s.salesSettlementRepo.DeleteItems(ctx, int64(id)); err != nil {
		return err
	}
	return s.salesSettlementRepo.Delete(ctx, id)
}

// ==================== 付款申请 ====================

func (s *FinService) ListPaymentRequests(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.PaymentRequest, int64, error) {
	return s.paymentRequestRepo.List(ctx, tenantID, query)
}

func (s *FinService) GetPaymentRequest(ctx context.Context, id uint) (*model.PaymentRequest, error) {
	return s.paymentRequestRepo.GetByID(ctx, id)
}

func (s *FinService) CreatePaymentRequest(ctx context.Context, tenantID int64, req *model.PaymentRequestCreate, createdBy string) (*model.PaymentRequest, error) {
	now := time.Now()
	seq := now.UnixNano() % 10000
	requestNo := fmt.Sprintf("PR%s%04d", now.Format("20060102"), seq)

	paymentReq := &model.PaymentRequest{
		RequestNo:    requestNo,
		RequestType:  req.RequestType,
		RequestAmount: req.RequestAmount,
		Purpose:      req.Purpose,
		BankName:     req.BankName,
		BankAccount:  req.BankAccount,
		Status:       "PENDING",
		ApprovalStatus: "PENDING",
		PaymentStatus: "UNPAID",
		TenantID:     tenantID,
		CreatedBy:   &createdBy,
	}

	if err := s.paymentRequestRepo.Create(ctx, paymentReq); err != nil {
		return nil, err
	}

	return paymentReq, nil
}

func (s *FinService) SubmitPaymentRequest(ctx context.Context, id uint) error {
	return s.paymentRequestRepo.Update(ctx, id, map[string]interface{}{
		"status": "SUBMITTED",
	})
}

func (s *FinService) ApprovePaymentRequest(ctx context.Context, id uint, approvedBy int64, comment string) error {
	now := time.Now()
	updates := map[string]interface{}{
		"status":          "APPROVED",
		"approval_status": "APPROVED",
		"approved_by":      approvedBy,
		"approved_time":    now,
	}
	if comment != "" {
		updates["approver_comment"] = comment
	}
	return s.paymentRequestRepo.Update(ctx, id, updates)
}

func (s *FinService) RejectPaymentRequest(ctx context.Context, id uint, rejectedBy int64, comment string) error {
	now := time.Now()
	return s.paymentRequestRepo.Update(ctx, id, map[string]interface{}{
		"status":            "REJECTED",
		"approval_status":    "REJECTED",
		"approved_by":        rejectedBy,
		"approved_time":      now,
		"approver_comment":   comment,
	})
}

func (s *FinService) PayPaymentRequest(ctx context.Context, id uint, paidBy int64) error {
	now := time.Now()
	return s.paymentRequestRepo.Update(ctx, id, map[string]interface{}{
		"payment_status": "PAID",
		"paid_by":         paidBy,
		"paid_time":       now,
	})
}

func (s *FinService) DeletePaymentRequest(ctx context.Context, id uint) error {
	return s.paymentRequestRepo.Delete(ctx, id)
}

// ==================== 采购预付款 ====================

func (s *FinService) ListPurchaseAdvances(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.PurchaseAdvance, int64, error) {
	return s.purchaseAdvanceRepo.List(ctx, tenantID, query)
}

func (s *FinService) CreatePurchaseAdvance(ctx context.Context, tenantID int64, advance *model.PurchaseAdvance) error {
	now := time.Now()
	seq := now.UnixNano() % 10000
	advanceNo := fmt.Sprintf("PA%s%04d", now.Format("20060102"), seq)
	advance.AdvanceNo = advanceNo
	advance.Status = "PENDING"
	advance.TenantID = tenantID
	return s.purchaseAdvanceRepo.Create(ctx, advance)
}

// ==================== 销售收款 ====================

func (s *FinService) ListSalesReceipts(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.SalesReceipt, int64, error) {
	return s.salesReceiptRepo.List(ctx, tenantID, query)
}

func (s *FinService) CreateSalesReceipt(ctx context.Context, tenantID int64, receipt *model.SalesReceipt) error {
	now := time.Now()
	seq := now.UnixNano() % 10000
	receiptNo := fmt.Sprintf("SR%s%04d", now.Format("20060102"), seq)
	receipt.ReceiptNo = receiptNo
	receipt.Status = "PENDING"
	receipt.TenantID = tenantID
	return s.salesReceiptRepo.Create(ctx, receipt)
}

// ==================== 供应商对账 ====================

func (s *FinService) ListSupplierStatements(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.SupplierStatement, int64, error) {
	return s.supplierStatementRepo.List(ctx, tenantID, query)
}

func (s *FinService) GetSupplierStatement(ctx context.Context, id uint) (*model.SupplierStatement, error) {
	return s.supplierStatementRepo.GetByID(ctx, id)
}
