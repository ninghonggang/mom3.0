package repository

import (
	"context"
	"fmt"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// PurchaseOrderRepository 采购订单仓库
type PurchaseOrderRepository struct {
	db *gorm.DB
}

func NewPurchaseOrderRepository(db *gorm.DB) *PurchaseOrderRepository {
	return &PurchaseOrderRepository{db: db}
}

func (r *PurchaseOrderRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.PurchaseOrder, int64, error) {
	var list []model.PurchaseOrder
	var total int64

	q := r.db.WithContext(ctx).Model(&model.PurchaseOrder{}).Where("tenant_id = ?", tenantID)

	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}
	if supplierID, ok := query["supplier_id"]; ok && supplierID != nil {
		q = q.Where("supplier_id = ?", supplierID)
	}
	if startDate, ok := query["start_date"]; ok && startDate != "" {
		q = q.Where("order_date >= ?", startDate)
	}
	if endDate, ok := query["end_date"]; ok && endDate != "" {
		q = q.Where("order_date <= ?", endDate)
	}

	q.Count(&total)
	q = q.Preload("Items").Order("id DESC")
	if page, ok := query["page"].(int); ok && page > 0 {
		limit := 20
		if limitVal, ok := query["limit"].(int); ok && limitVal > 0 {
			limit = limitVal
		}
		q = q.Offset((page - 1) * limit).Limit(limit)
	}

	err := q.Find(&list).Error
	return list, total, err
}

func (r *PurchaseOrderRepository) GetByID(ctx context.Context, id uint) (*model.PurchaseOrder, error) {
	var order model.PurchaseOrder
	err := r.db.WithContext(ctx).Preload("Items").First(&order, id).Error
	return &order, err
}

func (r *PurchaseOrderRepository) GetByNo(ctx context.Context, tenantID int64, poNo string) (*model.PurchaseOrder, error) {
	var order model.PurchaseOrder
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND po_no = ?", tenantID, poNo).First(&order).Error
	return &order, err
}

func (r *PurchaseOrderRepository) Create(ctx context.Context, order *model.PurchaseOrder) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *PurchaseOrderRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.PurchaseOrder{}).Where("id = ?", id).Updates(updates).Error
}

func (r *PurchaseOrderRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		tx.Where("po_id = ?", id).Delete(&model.PurchaseOrderItem{})
		return tx.Delete(&model.PurchaseOrder{}, id).Error
	})
}

func (r *PurchaseOrderRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).Model(&model.PurchaseOrder{}).Where("id = ?", id).Update("status", status).Error
}

// GetItem 获取行项目
func (r *PurchaseOrderRepository) GetItem(ctx context.Context, itemID uint) (*model.PurchaseOrderItem, error) {
	var item model.PurchaseOrderItem
	err := r.db.WithContext(ctx).First(&item, itemID).Error
	return &item, err
}

// UpdateItem 更新行项目
func (r *PurchaseOrderRepository) UpdateItem(ctx context.Context, itemID uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.PurchaseOrderItem{}).Where("id = ?", itemID).Updates(updates).Error
}

// UpdateItemStatus 更新行项目状态
func (r *PurchaseOrderRepository) UpdateItemStatus(ctx context.Context, itemID uint, status string) error {
	return r.db.WithContext(ctx).Model(&model.PurchaseOrderItem{}).Where("id = ?", itemID).Update("status", status).Error
}

// UpdateItemWithPORecalc 在事务中更新行项目并重新计算PO汇总
func (r *PurchaseOrderRepository) UpdateItemWithPORecalc(ctx context.Context, poID int64, itemID uint, itemUpdates map[string]interface{}) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 更新行项目
		if err := tx.Model(&model.PurchaseOrderItem{}).Where("id = ?", itemID).Updates(itemUpdates).Error; err != nil {
			return err
		}
		// 重新计算PO汇总
		return r.updatePOWithItemsRecalcTx(ctx, tx, poID)
	})
}

// updatePOWithItemsRecalcTx 事务中重新计算PO汇总
func (r *PurchaseOrderRepository) updatePOWithItemsRecalcTx(ctx context.Context, tx *gorm.DB, poID int64) error {
	var items []model.PurchaseOrderItem
	if err := tx.Where("po_id = ?", poID).Find(&items).Error; err != nil {
		return err
	}

	var totalQty, totalAmount, receivedQty float64
	for _, item := range items {
		totalQty += item.OrderQty
		totalAmount += item.LineAmount
		receivedQty += item.ReceivedQty
	}

	// 判断订单状态
	status := "ISSUED"
	allCompleted := true
	anyPartial := false
	for _, item := range items {
		if item.Status == "COMPLETED" {
			continue
		} else if item.Status == "PARTIAL" {
			anyPartial = true
			allCompleted = false
		} else {
			allCompleted = false
		}
	}
	if allCompleted && len(items) > 0 {
		status = "RECEIVED"
	} else if anyPartial {
		status = "PARTIAL"
	}

	return tx.Model(&model.PurchaseOrder{}).Where("id = ?", poID).Updates(map[string]interface{}{
		"total_qty":    totalQty,
		"total_amount": totalAmount,
		"received_qty":  receivedQty,
		"status":       status,
	}).Error
}

// RFQRepository 询价单仓库
type RFQRepository struct {
	db *gorm.DB
}

func NewRFQRepository(db *gorm.DB) *RFQRepository {
	return &RFQRepository{db: db}
}

func (r *RFQRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.RFQ, int64, error) {
	var list []model.RFQ
	var total int64

	q := r.db.WithContext(ctx).Model(&model.RFQ{}).Where("tenant_id = ?", tenantID)

	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}

	q.Count(&total)
	q = q.Preload("Items").Preload("Invites").Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *RFQRepository) GetByID(ctx context.Context, id uint) (*model.RFQ, error) {
	var rfq model.RFQ
	err := r.db.WithContext(ctx).Preload("Items").Preload("Invites").First(&rfq, id).Error
	return &rfq, err
}

func (r *RFQRepository) GetByNo(ctx context.Context, tenantID int64, rfqNo string) (*model.RFQ, error) {
	var rfq model.RFQ
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND rfq_no = ?", tenantID, rfqNo).First(&rfq).Error
	return &rfq, err
}

func (r *RFQRepository) Create(ctx context.Context, rfq *model.RFQ) error {
	return r.db.WithContext(ctx).Create(rfq).Error
}

func (r *RFQRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.RFQ{}).Where("id = ?", id).Updates(updates).Error
}

func (r *RFQRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		tx.Where("rfq_id = ?", id).Delete(&model.RFQItem{})
		tx.Where("rfq_id = ?", id).Delete(&model.RFQInvite{})
		return tx.Delete(&model.RFQ{}, id).Error
	})
}

// SupplierQuoteRepository 供应商报价仓库
type SupplierQuoteRepository struct {
	db *gorm.DB
}

func NewSupplierQuoteRepository(db *gorm.DB) *SupplierQuoteRepository {
	return &SupplierQuoteRepository{db: db}
}

func (r *SupplierQuoteRepository) ListByRFQ(ctx context.Context, rfqID uint) ([]model.SupplierQuote, error) {
	var list []model.SupplierQuote
	err := r.db.WithContext(ctx).Preload("Items").Where("rfq_id = ?", rfqID).Order("total_amount ASC").Find(&list).Error
	return list, err
}

func (r *SupplierQuoteRepository) GetByID(ctx context.Context, id uint) (*model.SupplierQuote, error) {
	var quote model.SupplierQuote
	err := r.db.WithContext(ctx).Preload("Items").First(&quote, id).Error
	return &quote, err
}

func (r *SupplierQuoteRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.SupplierQuote, int64, error) {
	var list []model.SupplierQuote
	var total int64

	q := r.db.WithContext(ctx).Model(&model.SupplierQuote{}).Where("tenant_id = ?", tenantID)

	if rfqNo := query["rfq_no"]; rfqNo != nil && rfqNo != "" {
		q = q.Where("rfq_no LIKE ?", "%"+rfqNo.(string)+"%")
	}
	if supplierName := query["supplier_name"]; supplierName != nil && supplierName != "" {
		q = q.Where("supplier_name LIKE ?", "%"+supplierName.(string)+"%")
	}
	if status := query["status"]; status != nil && status != "" {
		q = q.Where("quote_status = ?", status)
	}

	q.Count(&total)
	q = q.Preload("Items").Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *SupplierQuoteRepository) Create(ctx context.Context, quote *model.SupplierQuote) error {
	return r.db.WithContext(ctx).Create(quote).Error
}

func (r *SupplierQuoteRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.SupplierQuote{}).Where("id = ?", id).Updates(updates).Error
}

// SCPSalesOrderRepository 销售订单仓库
type SCPSalesOrderRepository struct {
	db *gorm.DB
}

func NewSCPSalesOrderRepository(db *gorm.DB) *SCPSalesOrderRepository {
	return &SCPSalesOrderRepository{db: db}
}

func (r *SCPSalesOrderRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.SCPSalesOrder, int64, error) {
	var list []model.SCPSalesOrder
	var total int64

	q := r.db.WithContext(ctx).Model(&model.SCPSalesOrder{}).Where("tenant_id = ?", tenantID)

	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}
	if customerID, ok := query["customer_id"]; ok && customerID != nil {
		q = q.Where("customer_id = ?", customerID)
	}

	q.Count(&total)
	q = q.Preload("Items").Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *SCPSalesOrderRepository) GetByID(ctx context.Context, id uint) (*model.SCPSalesOrder, error) {
	var order model.SCPSalesOrder
	err := r.db.WithContext(ctx).Preload("Items").First(&order, id).Error
	return &order, err
}

func (r *SCPSalesOrderRepository) GetByNo(ctx context.Context, tenantID int64, soNo string) (*model.SCPSalesOrder, error) {
	var order model.SCPSalesOrder
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND so_no = ?", tenantID, soNo).First(&order).Error
	return &order, err
}

func (r *SCPSalesOrderRepository) Create(ctx context.Context, order *model.SCPSalesOrder) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *SCPSalesOrderRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.SCPSalesOrder{}).Where("id = ?", id).Updates(updates).Error
}

func (r *SCPSalesOrderRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		tx.Where("so_id = ?", id).Delete(&model.SCPSalesOrderItem{})
		return tx.Delete(&model.SCPSalesOrder{}, id).Error
	})
}

func (r *SCPSalesOrderRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).Model(&model.SCPSalesOrder{}).Where("id = ?", id).Update("status", status).Error
}

// CustomerInquiryRepository 客户询价仓库
type CustomerInquiryRepository struct {
	db *gorm.DB
}

func NewCustomerInquiryRepository(db *gorm.DB) *CustomerInquiryRepository {
	return &CustomerInquiryRepository{db: db}
}

func (r *CustomerInquiryRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.CustomerInquiry, int64, error) {
	var list []model.CustomerInquiry
	var total int64

	q := r.db.WithContext(ctx).Model(&model.CustomerInquiry{}).Where("tenant_id = ?", tenantID)

	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}

	q.Count(&total)
	q = q.Preload("Items").Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *CustomerInquiryRepository) GetByID(ctx context.Context, id uint) (*model.CustomerInquiry, error) {
	var inquiry model.CustomerInquiry
	err := r.db.WithContext(ctx).Preload("Items").First(&inquiry, id).Error
	return &inquiry, err
}

func (r *CustomerInquiryRepository) Create(ctx context.Context, inquiry *model.CustomerInquiry) error {
	return r.db.WithContext(ctx).Create(inquiry).Error
}

func (r *CustomerInquiryRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.CustomerInquiry{}).Where("id = ?", id).Updates(updates).Error
}

func (r *CustomerInquiryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.CustomerInquiry{}).Error
}

func (r *CustomerInquiryRepository) DeleteItems(ctx context.Context, inquiryID uint) error {
	return r.db.WithContext(ctx).Where("inquiry_id = ?", inquiryID).Delete(&model.InquiryItem{}).Error
}

// SupplierKPIRepository 供应商绩效仓库
type SupplierKPIRepository struct {
	db *gorm.DB
}

func NewSupplierKPIRepository(db *gorm.DB) *SupplierKPIRepository {
	return &SupplierKPIRepository{db: db}
}

func (r *SupplierKPIRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.SupplierKPI, int64, error) {
	var list []model.SupplierKPI
	var total int64

	q := r.db.WithContext(ctx).Model(&model.SupplierKPI{}).Where("tenant_id = ?", tenantID)

	if supplierID, ok := query["supplier_id"]; ok && supplierID != nil {
		q = q.Where("supplier_id = ?", supplierID)
	}
	if month, ok := query["evaluation_month"]; ok && month != "" {
		q = q.Where("evaluation_month = ?", month)
	}

	q.Count(&total)
	q = q.Order("evaluation_month DESC, total_score DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *SupplierKPIRepository) GetBySupplierMonthly(ctx context.Context, supplierID int64, month string) (*model.SupplierKPI, error) {
	var kpi model.SupplierKPI
	err := r.db.WithContext(ctx).Where("supplier_id = ? AND evaluation_month = ?", supplierID, month).First(&kpi).Error
	return &kpi, err
}

func (r *SupplierKPIRepository) Create(ctx context.Context, kpi *model.SupplierKPI) error {
	return r.db.WithContext(ctx).Create(kpi).Error
}

func (r *SupplierKPIRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.SupplierKPI{}).Where("id = ?", id).Updates(updates).Error
}

func (r *SupplierKPIRepository) GetRanking(ctx context.Context, tenantID int64, month string) ([]model.SupplierKPI, error) {
	var list []model.SupplierKPI
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND evaluation_month = ?", tenantID, month).Order("rank_position ASC").Find(&list).Error
	return list, err
}

// SupplierPurchaseInfoRepository 供应商采购信息仓库
type SupplierPurchaseInfoRepository struct {
	db *gorm.DB
}

func NewSupplierPurchaseInfoRepository(db *gorm.DB) *SupplierPurchaseInfoRepository {
	return &SupplierPurchaseInfoRepository{db: db}
}

func (r *SupplierPurchaseInfoRepository) GetBySupplierID(ctx context.Context, supplierID uint) (*model.SupplierPurchaseInfo, error) {
	var info model.SupplierPurchaseInfo
	err := r.db.WithContext(ctx).Where("supplier_id = ?", supplierID).First(&info).Error
	return &info, err
}

func (r *SupplierPurchaseInfoRepository) Create(ctx context.Context, info *model.SupplierPurchaseInfo) error {
	return r.db.WithContext(ctx).Create(info).Error
}

func (r *SupplierPurchaseInfoRepository) Update(ctx context.Context, supplierID uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.SupplierPurchaseInfo{}).Where("supplier_id = ?", supplierID).Updates(updates).Error
}

// POChangeLogRepository 采购订单变更记录仓库
type POChangeLogRepository struct {
	db *gorm.DB
}

func NewPOChangeLogRepository(db *gorm.DB) *POChangeLogRepository {
	return &POChangeLogRepository{db: db}
}

func (r *POChangeLogRepository) ListByPOID(ctx context.Context, poID uint) ([]model.POChangeLog, error) {
	var list []model.POChangeLog
	err := r.db.WithContext(ctx).Where("po_id = ?", poID).Order("change_time DESC").Find(&list).Error
	return list, err
}

func (r *POChangeLogRepository) Create(ctx context.Context, log *model.POChangeLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

// QuoteComparisonRepository 报价对比仓库
type QuoteComparisonRepository struct {
	db *gorm.DB
}

func NewQuoteComparisonRepository(db *gorm.DB) *QuoteComparisonRepository {
	return &QuoteComparisonRepository{db: db}
}

func (r *QuoteComparisonRepository) List(ctx context.Context, tenantID int64) ([]model.QuoteComparison, int64, error) {
	var list []model.QuoteComparison
	var total int64
	err := r.db.WithContext(ctx).Model(&model.QuoteComparison{}).Where("tenant_id = ?", tenantID).Count(&total).Order("compared_at DESC").Find(&list).Error
	return list, total, err
}

func (r *QuoteComparisonRepository) GetByRFQID(ctx context.Context, rfqID uint) (*model.QuoteComparison, error) {
	var comp model.QuoteComparison
	err := r.db.WithContext(ctx).Where("rfq_id = ?", rfqID).First(&comp).Error
	return &comp, err
}

func (r *QuoteComparisonRepository) Create(ctx context.Context, comp *model.QuoteComparison) error {
	return r.db.WithContext(ctx).Create(comp).Error
}

// SupplierDeliveryRecordRepository 供应商交货记录仓库
type SupplierDeliveryRecordRepository struct {
	db *gorm.DB
}

func NewSupplierDeliveryRecordRepository(db *gorm.DB) *SupplierDeliveryRecordRepository {
	return &SupplierDeliveryRecordRepository{db: db}
}

func (r *SupplierDeliveryRecordRepository) Create(ctx context.Context, record *model.SupplierDeliveryRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *SupplierDeliveryRecordRepository) ListBySupplier(ctx context.Context, supplierID int64, limit int) ([]model.SupplierDeliveryRecord, error) {
	var list []model.SupplierDeliveryRecord
	err := r.db.WithContext(ctx).Where("supplier_id = ?", supplierID).Order("created_at DESC").Limit(limit).Find(&list).Error
	return list, err
}

// Helper function to generate code
func GenerateCode(prefix string, seq uint) string {
	return fmt.Sprintf("%s-%06d", prefix, seq)
}
