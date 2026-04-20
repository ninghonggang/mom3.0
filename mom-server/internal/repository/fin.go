package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"mom-server/internal/model"
	"time"

	"gorm.io/gorm"
)

// PurchaseSettlementRepository 采购结算仓储
type PurchaseSettlementRepository struct {
	db *gorm.DB
}

func NewPurchaseSettlementRepository(db *gorm.DB) *PurchaseSettlementRepository {
	return &PurchaseSettlementRepository{db: db}
}

func (r *PurchaseSettlementRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.PurchaseSettlement, int64, error) {
	var list []model.PurchaseSettlement
	var total int64

	q := r.db.WithContext(ctx).Model(&model.PurchaseSettlement{}).Where("tenant_id = ?", tenantID)

	if supplierID, ok := query["supplier_id"]; ok && supplierID.(int64) > 0 {
		q = q.Where("supplier_id = ?", supplierID)
	}
	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}
	if startDate, ok := query["start_date"]; ok && startDate != "" {
		q = q.Where("settlement_date >= ?", startDate)
	}
	if endDate, ok := query["end_date"]; ok && endDate != "" {
		q = q.Where("settlement_date <= ?", endDate)
	}
	if settlementNo, ok := query["settlement_no"]; ok && settlementNo != "" {
		q = q.Where("settlement_no LIKE ?", "%"+settlementNo.(string)+"%")
	}

	q.Count(&total)

	page := 1
	limit := 20
	if p, ok := query["page"].(int); ok && p > 0 {
		page = p
	}
	if l, ok := query["limit"].(int); ok && l > 0 {
		limit = l
	}
	q = q.Offset((page - 1) * limit).Limit(limit).Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *PurchaseSettlementRepository) GetByID(ctx context.Context, id uint) (*model.PurchaseSettlement, error) {
	var settlement model.PurchaseSettlement
	err := r.db.WithContext(ctx).First(&settlement, id).Error
	return &settlement, err
}

func (r *PurchaseSettlementRepository) GetBySettlementNo(ctx context.Context, tenantID int64, no string) (*model.PurchaseSettlement, error) {
	var settlement model.PurchaseSettlement
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND settlement_no = ?", tenantID, no).First(&settlement).Error
	return &settlement, err
}

func (r *PurchaseSettlementRepository) Create(ctx context.Context, s *model.PurchaseSettlement) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *PurchaseSettlementRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.PurchaseSettlement{}).Where("id = ?", id).Updates(updates).Error
}

func (r *PurchaseSettlementRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.PurchaseSettlement{}, id).Error
}

// Items

func (r *PurchaseSettlementRepository) GetItems(ctx context.Context, settlementID int64) ([]model.PurchaseSettlementItem, error) {
	var items []model.PurchaseSettlementItem
	err := r.db.WithContext(ctx).Where("settlement_id = ?", settlementID).Order("line_no").Find(&items).Error
	return items, err
}

func (r *PurchaseSettlementRepository) CreateItem(ctx context.Context, item *model.PurchaseSettlementItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *PurchaseSettlementRepository) DeleteItems(ctx context.Context, settlementID int64) error {
	return r.db.WithContext(ctx).Where("settlement_id = ?", settlementID).Delete(&model.PurchaseSettlementItem{}).Error
}

// SalesSettlementRepository 销售结算仓储
type SalesSettlementRepository struct {
	db *gorm.DB
}

func NewSalesSettlementRepository(db *gorm.DB) *SalesSettlementRepository {
	return &SalesSettlementRepository{db: db}
}

func (r *SalesSettlementRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.SalesSettlement, int64, error) {
	var list []model.SalesSettlement
	var total int64

	q := r.db.WithContext(ctx).Model(&model.SalesSettlement{}).Where("tenant_id = ?", tenantID)

	if customerID, ok := query["customer_id"]; ok && customerID.(int64) > 0 {
		q = q.Where("customer_id = ?", customerID)
	}
	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}
	if startDate, ok := query["start_date"]; ok && startDate != "" {
		q = q.Where("settlement_date >= ?", startDate)
	}
	if endDate, ok := query["end_date"]; ok && endDate != "" {
		q = q.Where("settlement_date <= ?", endDate)
	}
	if settlementNo, ok := query["settlement_no"]; ok && settlementNo != "" {
		q = q.Where("settlement_no LIKE ?", "%"+settlementNo.(string)+"%")
	}

	q.Count(&total)

	page := 1
	limit := 20
	if p, ok := query["page"].(int); ok && p > 0 {
		page = p
	}
	if l, ok := query["limit"].(int); ok && l > 0 {
		limit = l
	}
	q = q.Offset((page - 1) * limit).Limit(limit).Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *SalesSettlementRepository) GetByID(ctx context.Context, id uint) (*model.SalesSettlement, error) {
	var settlement model.SalesSettlement
	err := r.db.WithContext(ctx).First(&settlement, id).Error
	return &settlement, err
}

func (r *SalesSettlementRepository) Create(ctx context.Context, s *model.SalesSettlement) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *SalesSettlementRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.SalesSettlement{}).Where("id = ?", id).Updates(updates).Error
}

func (r *SalesSettlementRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.SalesSettlement{}, id).Error
}

func (r *SalesSettlementRepository) GetItems(ctx context.Context, settlementID int64) ([]model.SalesSettlementItem, error) {
	var items []model.SalesSettlementItem
	err := r.db.WithContext(ctx).Where("settlement_id = ?", settlementID).Order("line_no").Find(&items).Error
	return items, err
}

func (r *SalesSettlementRepository) CreateItem(ctx context.Context, item *model.SalesSettlementItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *SalesSettlementRepository) DeleteItems(ctx context.Context, settlementID int64) error {
	return r.db.WithContext(ctx).Where("settlement_id = ?", settlementID).Delete(&model.SalesSettlementItem{}).Error
}

// PaymentRequestRepository 付款申请仓储
type PaymentRequestRepository struct {
	db *gorm.DB
}

func NewPaymentRequestRepository(db *gorm.DB) *PaymentRequestRepository {
	return &PaymentRequestRepository{db: db}
}

func (r *PaymentRequestRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.PaymentRequest, int64, error) {
	var list []model.PaymentRequest
	var total int64

	q := r.db.WithContext(ctx).Model(&model.PaymentRequest{}).Where("tenant_id = ?", tenantID)

	if requestType, ok := query["request_type"]; ok && requestType != "" {
		q = q.Where("request_type = ?", requestType)
	}
	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}
	if startDate, ok := query["start_date"]; ok && startDate != "" {
		q = q.Where("created_at >= ?", startDate)
	}
	if endDate, ok := query["end_date"]; ok && endDate != "" {
		q = q.Where("created_at <= ?", endDate)
	}

	q.Count(&total)

	page := 1
	limit := 20
	if p, ok := query["page"].(int); ok && p > 0 {
		page = p
	}
	if l, ok := query["limit"].(int); ok && l > 0 {
		limit = l
	}
	q = q.Offset((page - 1) * limit).Limit(limit).Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *PaymentRequestRepository) GetByID(ctx context.Context, id uint) (*model.PaymentRequest, error) {
	var req model.PaymentRequest
	err := r.db.WithContext(ctx).First(&req, id).Error
	return &req, err
}

func (r *PaymentRequestRepository) Create(ctx context.Context, req *model.PaymentRequest) error {
	return r.db.WithContext(ctx).Create(req).Error
}

func (r *PaymentRequestRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.PaymentRequest{}).Where("id = ?", id).Updates(updates).Error
}

func (r *PaymentRequestRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.PaymentRequest{}, id).Error
}

// PurchaseAdvanceRepository 采购预付款仓储
type PurchaseAdvanceRepository struct {
	db *gorm.DB
}

func NewPurchaseAdvanceRepository(db *gorm.DB) *PurchaseAdvanceRepository {
	return &PurchaseAdvanceRepository{db: db}
}

func (r *PurchaseAdvanceRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.PurchaseAdvance, int64, error) {
	var list []model.PurchaseAdvance
	var total int64

	q := r.db.WithContext(ctx).Model(&model.PurchaseAdvance{}).Where("tenant_id = ?", tenantID)

	if supplierID, ok := query["supplier_id"]; ok && supplierID.(int64) > 0 {
		q = q.Where("supplier_id = ?", supplierID)
	}
	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}

	q.Count(&total)

	page := 1
	limit := 20
	if p, ok := query["page"].(int); ok && p > 0 {
		page = p
	}
	if l, ok := query["limit"].(int); ok && l > 0 {
		limit = l
	}
	q = q.Offset((page - 1) * limit).Limit(limit).Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *PurchaseAdvanceRepository) GetByID(ctx context.Context, id uint) (*model.PurchaseAdvance, error) {
	var advance model.PurchaseAdvance
	err := r.db.WithContext(ctx).First(&advance, id).Error
	return &advance, err
}

func (r *PurchaseAdvanceRepository) Create(ctx context.Context, a *model.PurchaseAdvance) error {
	return r.db.WithContext(ctx).Create(a).Error
}

func (r *PurchaseAdvanceRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.PurchaseAdvance{}).Where("id = ?", id).Updates(updates).Error
}

// SalesReceiptRepository 销售收款单仓储
type SalesReceiptRepository struct {
	db *gorm.DB
}

func NewSalesReceiptRepository(db *gorm.DB) *SalesReceiptRepository {
	return &SalesReceiptRepository{db: db}
}

func (r *SalesReceiptRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.SalesReceipt, int64, error) {
	var list []model.SalesReceipt
	var total int64

	q := r.db.WithContext(ctx).Model(&model.SalesReceipt{}).Where("tenant_id = ?", tenantID)

	if customerID, ok := query["customer_id"]; ok && customerID.(int64) > 0 {
		q = q.Where("customer_id = ?", customerID)
	}
	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}

	q.Count(&total)

	page := 1
	limit := 20
	if p, ok := query["page"].(int); ok && p > 0 {
		page = p
	}
	if l, ok := query["limit"].(int); ok && l > 0 {
		limit = l
	}
	q = q.Offset((page - 1) * limit).Limit(limit).Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *SalesReceiptRepository) GetByID(ctx context.Context, id uint) (*model.SalesReceipt, error) {
	var receipt model.SalesReceipt
	err := r.db.WithContext(ctx).First(&receipt, id).Error
	return &receipt, err
}

func (r *SalesReceiptRepository) Create(ctx context.Context, s *model.SalesReceipt) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *SalesReceiptRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.SalesReceipt{}).Where("id = ?", id).Updates(updates).Error
}

// SupplierStatementRepository 供应商对账单仓储
type SupplierStatementRepository struct {
	db *gorm.DB
}

func NewSupplierStatementRepository(db *gorm.DB) *SupplierStatementRepository {
	return &SupplierStatementRepository{db: db}
}

func (r *SupplierStatementRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.SupplierStatement, int64, error) {
	var list []model.SupplierStatement
	var total int64

	q := r.db.WithContext(ctx).Model(&model.SupplierStatement{}).Where("tenant_id = ?", tenantID)

	if supplierID, ok := query["supplier_id"]; ok && supplierID.(int64) > 0 {
		q = q.Where("supplier_id = ?", supplierID)
	}
	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}

	q.Count(&total)

	page := 1
	limit := 20
	if p, ok := query["page"].(int); ok && p > 0 {
		page = p
	}
	if l, ok := query["limit"].(int); ok && l > 0 {
		limit = l
	}
	q = q.Offset((page - 1) * limit).Limit(limit).Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *SupplierStatementRepository) GetByID(ctx context.Context, id uint) (*model.SupplierStatement, error) {
	var s model.SupplierStatement
	err := r.db.WithContext(ctx).First(&s, id).Error
	return &s, err
}

func (r *SupplierStatementRepository) Create(ctx context.Context, s *model.SupplierStatement) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *SupplierStatementRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.SupplierStatement{}).Where("id = ?", id).Updates(updates).Error
}

func (r *SupplierStatementRepository) GetDetails(ctx context.Context, statementID int64) ([]model.StatementDetail, error) {
	var details []model.StatementDetail
	err := r.db.WithContext(ctx).Where("statement_id = ?", statementID).Order("biz_date").Find(&details).Error
	return details, err
}

func (r *SupplierStatementRepository) CreateDetail(ctx context.Context, d *model.StatementDetail) error {
	return r.db.WithContext(ctx).Create(d).Error
}

// GenerateStatementNo 生成对账单号
func (r *SupplierStatementRepository) GenerateStatementNo(ctx context.Context, tenantID int64) (string, error) {
	var count int64
	r.db.WithContext(ctx).Model(&model.SupplierStatement{}).Where("tenant_id = ?", tenantID).Count(&count)
	return fmt.Sprintf("SS%s%c", time.Now().Format("200601"), 'A'+count%26), nil
}

// Helper for JSON marshal/unmarshal
func parseJSONArray[T any](data string) ([]T, error) {
	if data == "" {
		return []T{}, nil
	}
	var result []T
	err := json.Unmarshal([]byte(data), &result)
	return result, err
}
