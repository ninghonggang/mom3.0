package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type SCPService struct {
	poRepo          *repository.PurchaseOrderRepository
	rfqRepo        *repository.RFQRepository
	quoteRepo      *repository.SupplierQuoteRepository
	salesOrderRepo *repository.SCPSalesOrderRepository
	inquiryRepo    *repository.CustomerInquiryRepository
	kpiRepo        *repository.SupplierKPIRepository
	purchaseInfoRepo *repository.SupplierPurchaseInfoRepository
	poChangeRepo   *repository.POChangeLogRepository
	quoteCompRepo  *repository.QuoteComparisonRepository
}

func NewSCPService(
	poRepo *repository.PurchaseOrderRepository,
	rfqRepo *repository.RFQRepository,
	quoteRepo *repository.SupplierQuoteRepository,
	salesOrderRepo *repository.SCPSalesOrderRepository,
	inquiryRepo *repository.CustomerInquiryRepository,
	kpiRepo *repository.SupplierKPIRepository,
	purchaseInfoRepo *repository.SupplierPurchaseInfoRepository,
	poChangeRepo *repository.POChangeLogRepository,
	quoteCompRepo *repository.QuoteComparisonRepository,
) *SCPService {
	return &SCPService{
		poRepo:          poRepo,
		rfqRepo:        rfqRepo,
		quoteRepo:      quoteRepo,
		salesOrderRepo: salesOrderRepo,
		inquiryRepo:    inquiryRepo,
		kpiRepo:        kpiRepo,
		purchaseInfoRepo: purchaseInfoRepo,
		poChangeRepo:   poChangeRepo,
		quoteCompRepo:  quoteCompRepo,
	}
}

// ==================== 采购订单 ====================

func (s *SCPService) ListPurchaseOrders(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.PurchaseOrder, int64, error) {
	return s.poRepo.List(ctx, tenantID, query)
}

func (s *SCPService) GetPurchaseOrder(ctx context.Context, id string) (*model.PurchaseOrder, error) {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return nil, err
	}
	return s.poRepo.GetByID(ctx, orderID)
}

func (s *SCPService) CreatePurchaseOrder(ctx context.Context, order *model.PurchaseOrder) error {
	if order.PONo == "" {
		order.PONo = generateCode("PO", order.ID)
	}
	// 计算总金额和总数量
	for _, item := range order.Items {
		item.LineAmount = item.UnitPrice * item.OrderQty
		item.TaxAmount = item.LineAmount * order.TaxRate / 100
		order.TotalAmount += item.LineAmount + item.TaxAmount
		order.TotalQty += item.OrderQty
	}
	return s.poRepo.Create(ctx, order)
}

func (s *SCPService) UpdatePurchaseOrder(ctx context.Context, id string, order *model.PurchaseOrder) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.poRepo.Update(ctx, orderID, map[string]interface{}{
		"supplier_id":     order.SupplierID,
		"supplier_code":   order.SupplierCode,
		"supplier_name":   order.SupplierName,
		"contact_person":  order.ContactPerson,
		"contact_phone":   order.ContactPhone,
		"contact_email":   order.ContactEmail,
		"promised_date":   order.PromisedDate,
		"currency":        order.Currency,
		"payment_terms":   order.PaymentTerms,
		"tax_rate":        order.TaxRate,
		"remark":          order.Remark,
	})
}

func (s *SCPService) DeletePurchaseOrder(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.poRepo.Delete(ctx, orderID)
}

func (s *SCPService) SubmitPurchaseOrder(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.poRepo.UpdateStatus(ctx, orderID, "PENDING")
}

func (s *SCPService) ApprovePurchaseOrder(ctx context.Context, id string, approvedBy int64) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	now := time.Now()
	return s.poRepo.Update(ctx, orderID, map[string]interface{}{
		"approval_status": "APPROVED",
		"approved_by":     approvedBy,
		"approved_time":   now,
		"status":         "APPROVED",
	})
}

func (s *SCPService) RejectPurchaseOrder(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.poRepo.UpdateStatus(ctx, orderID, "REJECTED")
}

func (s *SCPService) IssuePurchaseOrder(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.poRepo.UpdateStatus(ctx, orderID, "ISSUED")
}

func (s *SCPService) ClosePurchaseOrder(ctx context.Context, id string, closeReason string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.poRepo.Update(ctx, orderID, map[string]interface{}{
		"status":       "CLOSED",
		"close_reason": closeReason,
	})
}

func (s *SCPService) CancelPurchaseOrder(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.poRepo.UpdateStatus(ctx, orderID, "CANCELLED")
}

// ReceivePurchaseOrderItem 按行项目收货
// itemID: 订单行项目ID
// receivedQty: 本次收货数量
// batchNo: 批号（可选）
func (s *SCPService) ReceivePurchaseOrderItem(ctx context.Context, itemID uint, receivedQty float64, batchNo string) error {
	// 获取行项目
	item, err := s.poRepo.GetItem(ctx, itemID)
	if err != nil {
		return fmt.Errorf("行项目不存在: %w", err)
	}

	if item.Status == "COMPLETED" {
		return fmt.Errorf("该行项目已收货完成，无法重复操作")
	}

	if item.Status == "CANCELLED" {
		return fmt.Errorf("该行项目已取消，无法收货")
	}

	// 获取PO信息用于事件
	po, err := s.poRepo.GetByID(ctx, uint(item.POID))
	if err != nil {
		return fmt.Errorf("采购订单不存在: %w", err)
	}

	// 计算新的已收货数量
	newReceivedQty := item.ReceivedQty + receivedQty
	if newReceivedQty > item.OrderQty {
		return fmt.Errorf("收货数量超过订单数量，当前订单数量: %.3f，已收货: %.3f，本次: %.3f",
			item.OrderQty, item.ReceivedQty, receivedQty)
	}

	// 确定行项目状态
	newStatus := "PARTIAL"
	if newReceivedQty >= item.OrderQty {
		newStatus = "COMPLETED"
	}

	// 更新行项目
	updates := map[string]interface{}{
		"received_qty":          newReceivedQty,
		"status":                 newStatus,
		"actual_delivery_date":   time.Now(),
	}
	if batchNo != "" {
		updates["batch_no"] = batchNo
	}

	// 在事务中更新行项目并重新计算订单汇总
	if err := s.poRepo.UpdateItemWithPORecalc(ctx, item.POID, itemID, updates); err != nil {
		return err
	}

	// 发布采购收货事件，触发WMS入库
	var materialID int64
	if item.MaterialID != nil {
		materialID = *item.MaterialID
	}
	var materialName string
	if item.MaterialName != nil {
		materialName = *item.MaterialName
	}

	event := NewPurchaseReceiveEvent(
		po.TenantID,
		int64(po.ID),
		po.PONo,
		po.SupplierID,
		po.SupplierName,
		itemID,
		materialID,
		item.MaterialCode,
		materialName,
		receivedQty,
		batchNo,
	)
	GetEventBus().Publish(ctx, event)

	return nil
}

// ReceivePurchaseOrder 整单收货（兼容旧接口）
func (s *SCPService) ReceivePurchaseOrder(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.poRepo.UpdateStatus(ctx, orderID, "RECEIVED")
}

// ==================== 询价单 RFQ ====================

func (s *SCPService) ListRFQs(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.RFQ, int64, error) {
	return s.rfqRepo.List(ctx, tenantID, query)
}

func (s *SCPService) GetRFQ(ctx context.Context, id string) (*model.RFQ, error) {
	var rfqID uint
	_, err := fmt.Sscanf(id, "%d", &rfqID)
	if err != nil {
		return nil, err
	}
	return s.rfqRepo.GetByID(ctx, rfqID)
}

func (s *SCPService) CreateRFQ(ctx context.Context, rfq *model.RFQ) error {
	if rfq.RFQNo == "" {
		rfq.RFQNo = generateCode("RFQ", rfq.ID)
	}
	return s.rfqRepo.Create(ctx, rfq)
}

func (s *SCPService) UpdateRFQ(ctx context.Context, id string, rfq *model.RFQ) error {
	var rfqID uint
	_, err := fmt.Sscanf(id, "%d", &rfqID)
	if err != nil {
		return err
	}
	return s.rfqRepo.Update(ctx, rfqID, map[string]interface{}{
		"rfq_name":       rfq.RFQName,
		"deadline_date":  rfq.DeadlineDate,
		"currency":       rfq.Currency,
		"payment_terms":  rfq.PaymentTerms,
		"delivery_terms": rfq.DeliveryTerms,
		"quality_standard": rfq.QualityStandard,
		"remark":         rfq.Remark,
	})
}

func (s *SCPService) DeleteRFQ(ctx context.Context, id string) error {
	var rfqID uint
	_, err := fmt.Sscanf(id, "%d", &rfqID)
	if err != nil {
		return err
	}
	return s.rfqRepo.Delete(ctx, rfqID)
}

func (s *SCPService) PublishRFQ(ctx context.Context, id string) error {
	var rfqID uint
	_, err := fmt.Sscanf(id, "%d", &rfqID)
	if err != nil {
		return err
	}
	return s.rfqRepo.Update(ctx, rfqID, map[string]interface{}{
		"status": "PUBLISHED",
	})
}

func (s *SCPService) CloseRFQ(ctx context.Context, id string) error {
	var rfqID uint
	_, err := fmt.Sscanf(id, "%d", &rfqID)
	if err != nil {
		return err
	}
	return s.rfqRepo.Update(ctx, rfqID, map[string]interface{}{
		"status": "CLOSED",
	})
}

// ==================== 供应商报价 ====================

func (s *SCPService) ListQuotesByRFQ(ctx context.Context, rfqID string) ([]model.SupplierQuote, error) {
	var id uint
	_, err := fmt.Sscanf(rfqID, "%d", &id)
	if err != nil {
		return nil, err
	}
	return s.quoteRepo.ListByRFQ(ctx, id)
}

func (s *SCPService) CreateQuote(ctx context.Context, quote *model.SupplierQuote) error {
	if quote.QuoteNo == "" {
		quote.QuoteNo = generateCode("QUOTE", quote.ID)
	}
	return s.quoteRepo.Create(ctx, quote)
}

func (s *SCPService) ListSupplierQuotes(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.SupplierQuote, int64, error) {
	return s.quoteRepo.List(ctx, tenantID, query)
}

func (s *SCPService) GetSupplierQuote(ctx context.Context, id string) (*model.SupplierQuote, error) {
	var quoteID uint
	_, err := fmt.Sscanf(id, "%d", &quoteID)
	if err != nil {
		return nil, err
	}
	return s.quoteRepo.GetByID(ctx, quoteID)
}

func (s *SCPService) AwardRFQ(ctx context.Context, rfqID string, supplierID int64, totalAmount float64) error {
	var id uint
	_, err := fmt.Sscanf(rfqID, "%d", &id)
	if err != nil {
		return err
	}
	return s.rfqRepo.Update(ctx, id, map[string]interface{}{
		"status":               "AWARDED",
		"awarded_supplier_id":  supplierID,
		"awarded_total_amount": totalAmount,
	})
}

// ==================== 销售订单 ====================

func (s *SCPService) ListSalesOrders(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.SCPSalesOrder, int64, error) {
	return s.salesOrderRepo.List(ctx, tenantID, query)
}

func (s *SCPService) GetSalesOrder(ctx context.Context, id string) (*model.SCPSalesOrder, error) {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return nil, err
	}
	return s.salesOrderRepo.GetByID(ctx, orderID)
}

func (s *SCPService) CreateSalesOrder(ctx context.Context, order *model.SCPSalesOrder) error {
	if order.SONo == "" {
		order.SONo = generateCode("SO", order.ID)
	}
	for _, item := range order.Items {
		item.LineAmount = item.UnitPrice * item.OrderQty
		item.TaxAmount = item.LineAmount * order.TaxRate / 100
		order.TotalAmount += item.LineAmount + item.TaxAmount
		order.TotalQty += item.OrderQty
	}
	return s.salesOrderRepo.Create(ctx, order)
}

func (s *SCPService) UpdateSalesOrder(ctx context.Context, id string, order *model.SCPSalesOrder) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.salesOrderRepo.Update(ctx, orderID, map[string]interface{}{
		"customer_id":       order.CustomerID,
		"customer_code":     order.CustomerCode,
		"customer_name":     order.CustomerName,
		"contact_person":    order.ContactPerson,
		"contact_phone":     order.ContactPhone,
		"contact_email":     order.ContactEmail,
		"promised_date":     order.PromisedDate,
		"currency":          order.Currency,
		"payment_terms":    order.PaymentTerms,
		"tax_rate":         order.TaxRate,
		"delivery_address": order.DeliveryAddress,
		"remark":           order.Remark,
	})
}

func (s *SCPService) DeleteSalesOrder(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.salesOrderRepo.Delete(ctx, orderID)
}

func (s *SCPService) SubmitSalesOrder(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.salesOrderRepo.UpdateStatus(ctx, orderID, "PENDING")
}

func (s *SCPService) ApproveSalesOrder(ctx context.Context, id string, approvedBy int64) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	now := time.Now()
	return s.salesOrderRepo.Update(ctx, orderID, map[string]interface{}{
		"approval_status": "APPROVED",
		"approved_by":     approvedBy,
		"approved_time":   now,
		"status":         "APPROVED",
	})
}

func (s *SCPService) RejectSalesOrder(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.salesOrderRepo.UpdateStatus(ctx, orderID, "REJECTED")
}

func (s *SCPService) ConfirmSalesOrder(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.salesOrderRepo.UpdateStatus(ctx, orderID, "CONFIRMED")
}

func (s *SCPService) CloseSalesOrder(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.salesOrderRepo.UpdateStatus(ctx, orderID, "CLOSED")
}

func (s *SCPService) CancelSalesOrder(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.salesOrderRepo.UpdateStatus(ctx, orderID, "CANCELLED")
}

// ==================== 客户询价 ====================

func (s *SCPService) ListCustomerInquiries(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.CustomerInquiry, int64, error) {
	return s.inquiryRepo.List(ctx, tenantID, query)
}

func (s *SCPService) GetCustomerInquiry(ctx context.Context, id string) (*model.CustomerInquiry, error) {
	var inquiryID uint
	_, err := fmt.Sscanf(id, "%d", &inquiryID)
	if err != nil {
		return nil, err
	}
	return s.inquiryRepo.GetByID(ctx, inquiryID)
}

func (s *SCPService) CreateCustomerInquiry(ctx context.Context, inquiry *model.CustomerInquiry) error {
	if inquiry.InquiryNo == "" {
		inquiry.InquiryNo = generateCode("INQ", inquiry.ID)
	}
	return s.inquiryRepo.Create(ctx, inquiry)
}

func (s *SCPService) UpdateCustomerInquiry(ctx context.Context, id string, inquiry *model.CustomerInquiry) error {
	var inquiryID uint
	_, err := fmt.Sscanf(id, "%d", &inquiryID)
	if err != nil {
		return err
	}
	return s.inquiryRepo.Update(ctx, inquiryID, map[string]interface{}{
		"customer_id":     inquiry.CustomerID,
		"customer_name":   inquiry.CustomerName,
		"contact_person":  inquiry.ContactPerson,
		"contact_phone":   inquiry.ContactPhone,
		"contact_email":   inquiry.ContactEmail,
		"expected_date":   inquiry.ExpectedDate,
		"valid_until":    inquiry.ValidUntil,
		"remark":         inquiry.Remark,
	})
}

func (s *SCPService) DeleteCustomerInquiry(ctx context.Context, id string) error {
	var inquiryID uint
	_, err := fmt.Sscanf(id, "%d", &inquiryID)
	if err != nil {
		return err
	}
	// 先删除明细
	_ = s.inquiryRepo.DeleteItems(ctx, inquiryID)
	return s.inquiryRepo.Delete(ctx, inquiryID)
}

func (s *SCPService) SendCustomerInquiry(ctx context.Context, id string) error {
	var inquiryID uint
	_, err := fmt.Sscanf(id, "%d", &inquiryID)
	if err != nil {
		return err
	}
	return s.inquiryRepo.Update(ctx, inquiryID, map[string]interface{}{
		"status": "SENT",
	})
}

func (s *SCPService) QuoteCustomerInquiry(ctx context.Context, id string, quotedAmount float64) error {
	var inquiryID uint
	_, err := fmt.Sscanf(id, "%d", &inquiryID)
	if err != nil {
		return err
	}
	return s.inquiryRepo.Update(ctx, inquiryID, map[string]interface{}{
		"status":         "QUOTED",
		"quoted_amount":   quotedAmount,
	})
}

func (s *SCPService) WinCustomerInquiry(ctx context.Context, id string, supplierID int64) error {
	var inquiryID uint
	_, err := fmt.Sscanf(id, "%d", &inquiryID)
	if err != nil {
		return err
	}
	return s.inquiryRepo.Update(ctx, inquiryID, map[string]interface{}{
		"status":              "WON",
		"winner_supplier_id":  supplierID,
	})
}

func (s *SCPService) LoseCustomerInquiry(ctx context.Context, id string) error {
	var inquiryID uint
	_, err := fmt.Sscanf(id, "%d", &inquiryID)
	if err != nil {
		return err
	}
	return s.inquiryRepo.Update(ctx, inquiryID, map[string]interface{}{
		"status": "LOST",
	})
}

func (s *SCPService) CancelCustomerInquiry(ctx context.Context, id string) error {
	var inquiryID uint
	_, err := fmt.Sscanf(id, "%d", &inquiryID)
	if err != nil {
		return err
	}
	return s.inquiryRepo.Update(ctx, inquiryID, map[string]interface{}{
		"status": "CANCELLED",
	})
}

// ==================== 供应商绩效 KPI ====================

func (s *SCPService) ListSupplierKPI(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.SupplierKPI, int64, error) {
	return s.kpiRepo.List(ctx, tenantID, query)
}

func (s *SCPService) GetSupplierKPIByMonthly(ctx context.Context, supplierID int64, month string) (*model.SupplierKPI, error) {
	return s.kpiRepo.GetBySupplierMonthly(ctx, supplierID, month)
}

func (s *SCPService) CreateSupplierKPI(ctx context.Context, kpi *model.SupplierKPI) error {
	return s.kpiRepo.Create(ctx, kpi)
}

func (s *SCPService) GetSupplierRanking(ctx context.Context, tenantID int64, month string) ([]model.SupplierKPI, error) {
	return s.kpiRepo.GetRanking(ctx, tenantID, month)
}

// ==================== 供应商采购信息 ====================

func (s *SCPService) GetSupplierPurchaseInfo(ctx context.Context, supplierID string) (*model.SupplierPurchaseInfo, error) {
	var id uint
	_, err := fmt.Sscanf(supplierID, "%d", &id)
	if err != nil {
		return nil, err
	}
	return s.purchaseInfoRepo.GetBySupplierID(ctx, id)
}

func (s *SCPService) CreateSupplierPurchaseInfo(ctx context.Context, info *model.SupplierPurchaseInfo) error {
	return s.purchaseInfoRepo.Create(ctx, info)
}

func (s *SCPService) UpdateSupplierPurchaseInfo(ctx context.Context, supplierID string, info *model.SupplierPurchaseInfo) error {
	var id uint
	_, err := fmt.Sscanf(supplierID, "%d", &id)
	if err != nil {
		return err
	}
	return s.purchaseInfoRepo.Update(ctx, id, map[string]interface{}{
		"payment_terms":           info.PaymentTerms,
		"credit_limit":            info.CreditLimit,
		"tax_rate":                info.TaxRate,
		"min_order_amount":        info.MinOrderAmount,
		"lead_time_days":          info.LeadTimeDays,
		"supplier_grade":         info.SupplierGrade,
		"is_preferred":            info.IsPreferred,
		"is_blacklist":            info.IsBlacklist,
		"blacklist_reason":        info.BlacklistReason,
		"cooperation_start_date":  info.CooperationStartDate,
		"cooperation_end_date":   info.CooperationEndDate,
	})
}

// Helper function to generate code
func generateCode(prefix string, id uint) string {
	return fmt.Sprintf("%s-%06d", prefix, id)
}
