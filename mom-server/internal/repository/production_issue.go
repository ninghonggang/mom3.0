package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// ProductionIssueRepository 生产发料仓储
type ProductionIssueRepository struct {
	db *gorm.DB
}

func NewProductionIssueRepository(db *gorm.DB) *ProductionIssueRepository {
	return &ProductionIssueRepository{db: db}
}

func (r *ProductionIssueRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.ProductionIssue, int64, error) {
	var list []model.ProductionIssue
	var total int64

	q := r.db.WithContext(ctx).Model(&model.ProductionIssue{}).Where("tenant_id = ?", tenantID)

	if orderID, ok := query["production_order_id"]; ok && orderID.(int64) > 0 {
		q = q.Where("production_order_id = ?", orderID)
	}
	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}
	if issueType, ok := query["issue_type"]; ok && issueType != "" {
		q = q.Where("issue_type = ?", issueType)
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
	q = q.Preload("Items").Offset((page-1)*limit).Limit(limit).Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *ProductionIssueRepository) GetByID(ctx context.Context, id uint) (*model.ProductionIssue, error) {
	var issue model.ProductionIssue
	err := r.db.WithContext(ctx).Preload("Items").First(&issue, id).Error
	return &issue, err
}

func (r *ProductionIssueRepository) GetByNo(ctx context.Context, tenantID int64, no string) (*model.ProductionIssue, error) {
	var issue model.ProductionIssue
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND issue_no = ?", tenantID, no).First(&issue).Error
	return &issue, err
}

func (r *ProductionIssueRepository) Create(ctx context.Context, issue *model.ProductionIssue) error {
	return r.db.WithContext(ctx).Create(issue).Error
}

func (r *ProductionIssueRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ProductionIssue{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ProductionIssueRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		tx.Where("issue_id = ?", id).Delete(&model.ProductionIssueItem{})
		return tx.Delete(&model.ProductionIssue{}, id).Error
	})
}

// ProductionIssueItemRepository 生产发料明细仓储
type ProductionIssueItemRepository struct {
	db *gorm.DB
}

func NewProductionIssueItemRepository(db *gorm.DB) *ProductionIssueItemRepository {
	return &ProductionIssueItemRepository{db: db}
}

func (r *ProductionIssueItemRepository) ListByIssueID(ctx context.Context, issueID int64) ([]model.ProductionIssueItem, error) {
	var items []model.ProductionIssueItem
	err := r.db.WithContext(ctx).Where("issue_id = ?", issueID).Order("line_no").Find(&items).Error
	return items, err
}

func (r *ProductionIssueItemRepository) CreateBatch(ctx context.Context, items []model.ProductionIssueItem) error {
	return r.db.WithContext(ctx).Create(&items).Error
}

func (r *ProductionIssueItemRepository) DeleteByIssueID(ctx context.Context, issueID int64) error {
	return r.db.WithContext(ctx).Where("issue_id = ?", issueID).Delete(&model.ProductionIssueItem{}).Error
}

func (r *ProductionIssueItemRepository) UpdatePickedQty(ctx context.Context, id uint, pickedQty float64) error {
	return r.db.WithContext(ctx).Model(&model.ProductionIssueItem{}).Where("id = ?", id).Updates(map[string]interface{}{
		"picked_qty": pickedQty,
	}).Error
}

func (r *ProductionIssueItemRepository) UpdateIssuedQty(ctx context.Context, id uint, issuedQty float64) error {
	return r.db.WithContext(ctx).Model(&model.ProductionIssueItem{}).Where("id = ?", id).Updates(map[string]interface{}{
		"issued_qty": issuedQty,
	}).Error
}
