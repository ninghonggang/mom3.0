package repository

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"time"

	"gorm.io/gorm"
)

// ========== OrderMonthRepository 月计划仓储 ==========

type OrderMonthRepository struct {
	db *gorm.DB
}

func NewOrderMonthRepository(db *gorm.DB) *OrderMonthRepository {
	return &OrderMonthRepository{db: db}
}

func (r *OrderMonthRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.OrderMonth, int64, error) {
	var list []model.OrderMonth
	var total int64

	q := r.db.WithContext(ctx).Model(&model.OrderMonth{}).Where("tenant_id = ?", tenantID)

	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("approval_status = ?", status)
	}
	if planMonth, ok := query["plan_month"]; ok && planMonth != "" {
		q = q.Where("plan_month = ?", planMonth)
	}
	if sourceType, ok := query["source_type"]; ok && sourceType != "" {
		q = q.Where("source_type = ?", sourceType)
	}
	if workshopID, ok := query["workshop_id"]; ok && workshopID.(int64) > 0 {
		q = q.Where("workshop_id = ?", workshopID)
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

func (r *OrderMonthRepository) GetByID(ctx context.Context, id uint) (*model.OrderMonth, error) {
	var order model.OrderMonth
	err := r.db.WithContext(ctx).First(&order, id).Error
	return &order, err
}

func (r *OrderMonthRepository) GetByMonthPlanNo(ctx context.Context, tenantID int64, no string) (*model.OrderMonth, error) {
	var order model.OrderMonth
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND month_plan_no = ?", tenantID, no).First(&order).Error
	return &order, err
}

func (r *OrderMonthRepository) Create(ctx context.Context, order *model.OrderMonth) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *OrderMonthRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.OrderMonth{}).Where("id = ?", id).Updates(updates).Error
}

func (r *OrderMonthRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.OrderMonth{}, id).Error
}

// GenerateMonthPlanNo 生成月计划单号 MP-YYYYMM
func (r *OrderMonthRepository) GenerateMonthPlanNo(ctx context.Context, tenantID int64, planMonth string) (string, error) {
	var count int64
	r.db.WithContext(ctx).Model(&model.OrderMonth{}).Where("tenant_id = ? AND plan_month = ?", tenantID, planMonth).Count(&count)
	return fmt.Sprintf("MP-%s-%03d", planMonth, count+1), nil
}

// ========== OrderMonthItemRepository 月计划明细仓储 ==========

type OrderMonthItemRepository struct {
	db *gorm.DB
}

func NewOrderMonthItemRepository(db *gorm.DB) *OrderMonthItemRepository {
	return &OrderMonthItemRepository{db: db}
}

func (r *OrderMonthItemRepository) ListByMonthPlanID(ctx context.Context, monthPlanID int64) ([]model.OrderMonthItem, error) {
	var items []model.OrderMonthItem
	err := r.db.WithContext(ctx).Where("month_plan_id = ?", monthPlanID).Order("line_no").Find(&items).Error
	return items, err
}

func (r *OrderMonthItemRepository) Create(ctx context.Context, item *model.OrderMonthItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *OrderMonthItemRepository) CreateBatch(ctx context.Context, items []model.OrderMonthItem) error {
	return r.db.WithContext(ctx).Create(&items).Error
}

func (r *OrderMonthItemRepository) DeleteByMonthPlanID(ctx context.Context, monthPlanID int64) error {
	return r.db.WithContext(ctx).Where("month_plan_id = ?", monthPlanID).Delete(&model.OrderMonthItem{}).Error
}

func (r *OrderMonthItemRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.OrderMonthItem{}).Where("id = ?", id).Updates(updates).Error
}

// ========== OrderMonthAuditRepository 月计划审核记录仓储 ==========

type OrderMonthAuditRepository struct {
	db *gorm.DB
}

func NewOrderMonthAuditRepository(db *gorm.DB) *OrderMonthAuditRepository {
	return &OrderMonthAuditRepository{db: db}
}

func (r *OrderMonthAuditRepository) Create(ctx context.Context, audit *model.OrderMonthAudit) error {
	return r.db.WithContext(ctx).Create(audit).Error
}

func (r *OrderMonthAuditRepository) ListByMonthPlanID(ctx context.Context, monthPlanID int64) ([]model.OrderMonthAudit, error) {
	var audits []model.OrderMonthAudit
	err := r.db.WithContext(ctx).Where("month_plan_id = ?", monthPlanID).Order("created_at DESC").Find(&audits).Error
	return audits, err
}

// ========== OrderDayRepository 日计划仓储 ==========

type OrderDayRepository struct {
	db *gorm.DB
}

func NewOrderDayRepository(db *gorm.DB) *OrderDayRepository {
	return &OrderDayRepository{db: db}
}

func (r *OrderDayRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.OrderDay, int64, error) {
	var list []model.OrderDay
	var total int64

	q := r.db.WithContext(ctx).Model(&model.OrderDay{}).Where("tenant_id = ?", tenantID)

	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}
	if planDate, ok := query["plan_date"]; ok && planDate != "" {
		q = q.Where("plan_date = ?", planDate)
	}
	if startDate, ok := query["start_date"]; ok && startDate != "" {
		q = q.Where("plan_date >= ?", startDate)
	}
	if endDate, ok := query["end_date"]; ok && endDate != "" {
		q = q.Where("plan_date <= ?", endDate)
	}
	if lineID, ok := query["line_id"]; ok && lineID.(int64) > 0 {
		q = q.Where("production_line_id = ?", lineID)
	}
	if monthPlanID, ok := query["month_plan_id"]; ok && monthPlanID.(int64) > 0 {
		q = q.Where("month_plan_id = ?", monthPlanID)
	}
	if workshopID, ok := query["workshop_id"]; ok && workshopID.(int64) > 0 {
		q = q.Where("workshop_id = ?", workshopID)
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

func (r *OrderDayRepository) GetByID(ctx context.Context, id uint) (*model.OrderDay, error) {
	var order model.OrderDay
	err := r.db.WithContext(ctx).First(&order, id).Error
	return &order, err
}

func (r *OrderDayRepository) GetByDayPlanNo(ctx context.Context, tenantID int64, no string) (*model.OrderDay, error) {
	var order model.OrderDay
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND day_plan_no = ?", tenantID, no).First(&order).Error
	return &order, err
}

func (r *OrderDayRepository) Create(ctx context.Context, order *model.OrderDay) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *OrderDayRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.OrderDay{}).Where("id = ?", id).Updates(updates).Error
}

func (r *OrderDayRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.OrderDay{}, id).Error
}

// GenerateDayPlanNo 生成日计划单号 DP-YYYYMMDD
func (r *OrderDayRepository) GenerateDayPlanNo(ctx context.Context, tenantID int64, planDate time.Time) (string, error) {
	dateStr := planDate.Format("20060102")
	var count int64
	r.db.WithContext(ctx).Model(&model.OrderDay{}).Where("tenant_id = ? AND plan_date = ?", tenantID, planDate.Format("2006-01-02")).Count(&count)
	return fmt.Sprintf("DP-%s-%03d", dateStr, count+1), nil
}

// GetOrderDaysByMonthID 获取指定月计划关联的所有日计划
func (r *OrderDayRepository) GetOrderDaysByMonthID(ctx context.Context, monthPlanID int64) ([]model.OrderDay, error) {
	var list []model.OrderDay
	err := r.db.WithContext(ctx).Where("month_plan_id = ?", monthPlanID).Order("plan_date ASC").Find(&list).Error
	return list, err
}

// ========== OrderDayItemRepository 日计划明细仓储 ==========

type OrderDayItemRepository struct {
	db *gorm.DB
}

func NewOrderDayItemRepository(db *gorm.DB) *OrderDayItemRepository {
	return &OrderDayItemRepository{db: db}
}

func (r *OrderDayItemRepository) ListByDayPlanID(ctx context.Context, dayPlanID int64) ([]model.OrderDayItem, error) {
	var items []model.OrderDayItem
	err := r.db.WithContext(ctx).Where("day_plan_id = ?", dayPlanID).Order("line_no").Find(&items).Error
	return items, err
}

func (r *OrderDayItemRepository) Create(ctx context.Context, item *model.OrderDayItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *OrderDayItemRepository) CreateBatch(ctx context.Context, items []model.OrderDayItem) error {
	return r.db.WithContext(ctx).Create(&items).Error
}

func (r *OrderDayItemRepository) DeleteByDayPlanID(ctx context.Context, dayPlanID int64) error {
	return r.db.WithContext(ctx).Where("day_plan_id = ?", dayPlanID).Delete(&model.OrderDayItem{}).Error
}

func (r *OrderDayItemRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.OrderDayItem{}).Where("id = ?", id).Updates(updates).Error
}

func (r *OrderDayItemRepository) GetByID(ctx context.Context, id uint) (*model.OrderDayItem, error) {
	var item model.OrderDayItem
	err := r.db.WithContext(ctx).First(&item, id).Error
	return &item, err
}

// ========== OrderDayWorkOrderMapRepository 工单生成记录仓储 ==========

type OrderDayWorkOrderMapRepository struct {
	db *gorm.DB
}

func NewOrderDayWorkOrderMapRepository(db *gorm.DB) *OrderDayWorkOrderMapRepository {
	return &OrderDayWorkOrderMapRepository{db: db}
}

func (r *OrderDayWorkOrderMapRepository) Create(ctx context.Context, m *model.OrderDayWorkOrderMap) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *OrderDayWorkOrderMapRepository) CreateBatch(ctx context.Context, maps []model.OrderDayWorkOrderMap) error {
	return r.db.WithContext(ctx).Create(&maps).Error
}

func (r *OrderDayWorkOrderMapRepository) ListByDayPlanID(ctx context.Context, dayPlanID int64) ([]model.OrderDayWorkOrderMap, error) {
	var maps []model.OrderDayWorkOrderMap
	err := r.db.WithContext(ctx).Where("day_plan_id = ?", dayPlanID).Find(&maps).Error
	return maps, err
}

func (r *OrderDayWorkOrderMapRepository) ListByDayPlanItemID(ctx context.Context, dayPlanItemID int64) ([]model.OrderDayWorkOrderMap, error) {
	var maps []model.OrderDayWorkOrderMap
	err := r.db.WithContext(ctx).Where("day_plan_item_id = ?", dayPlanItemID).Find(&maps).Error
	return maps, err
}
