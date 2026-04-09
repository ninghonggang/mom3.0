package repository

import (
	"context"
	"database/sql"
	"mom-server/internal/model"
	"time"

	"gorm.io/gorm"
)

type AndonRepository struct {
	db *gorm.DB
}

func NewAndonRepository(db *gorm.DB) *AndonRepository {
	return &AndonRepository{db: db}
}

// List 查询呼叫列表
func (r *AndonRepository) List(ctx context.Context, tenantID int64, workshopID int64, status string, andonType string, callNo string, startDate, endDate *time.Time, page, pageSize int) ([]model.AndonCall, int64, error) {
	var list []model.AndonCall
	var total int64

	query := r.db.WithContext(ctx).Model(&model.AndonCall{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if workshopID > 0 {
		query = query.Where("workshop_id = ?", workshopID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if andonType != "" {
		query = query.Where("andon_type = ?", andonType)
	}
	if callNo != "" {
		query = query.Where("call_no LIKE ?", "%"+callNo+"%")
	}
	if startDate != nil {
		query = query.Where("call_time >= ?", startDate)
	}
	if endDate != nil {
		query = query.Where("call_time <= ?", endDate)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		query = query.Offset(offset).Limit(pageSize)
	}

	err = query.Order("call_time DESC").Find(&list).Error
	return list, total, err
}

// ListPendingEscalations 查询所有未关闭的呼叫(用于定时升级检查)
func (r *AndonRepository) ListPendingEscalations(ctx context.Context) ([]model.AndonCall, error) {
	var list []model.AndonCall
	err := r.db.WithContext(ctx).
		Where("status IN ?", []string{"CALLING", "RESPONDED", "HANDLING"}).
		Order("call_time ASC").
		Find(&list).Error
	return list, err
}

// GetByID 根据ID获取
func (r *AndonRepository) GetByID(ctx context.Context, id int64) (*model.AndonCall, error) {
	var call model.AndonCall
	err := r.db.WithContext(ctx).First(&call, id).Error
	return &call, err
}

// GetByIDWithTenant 根据ID和租户ID获取
func (r *AndonRepository) GetByIDWithTenant(ctx context.Context, id, tenantID int64) (*model.AndonCall, error) {
	var call model.AndonCall
	err := r.db.WithContext(ctx).Where("id = ? AND tenant_id = ?", id, tenantID).First(&call).Error
	return &call, err
}

// Update 更新呼叫记录
func (r *AndonRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.AndonCall{}).Where("id = ?", id).Updates(updates).Error
}

// Create 创建呼叫记录
func (r *AndonRepository) Create(ctx context.Context, call *model.AndonCall) error {
	return r.db.WithContext(ctx).Create(call).Error
}

// GetStatistics 获取统计数据
func (r *AndonRepository) GetStatistics(ctx context.Context, tenantID, workshopID int64, startDate, endDate *time.Time) (map[string]interface{}, error) {
	var totalCalls int64
	var resolvedCalls int64
	var escalatedCalls int64

	// 基础查询条件
	baseWhere := r.db.WithContext(ctx).Model(&model.AndonCall{})
	if tenantID > 0 {
		baseWhere = baseWhere.Where("tenant_id = ?", tenantID)
	}
	if workshopID > 0 {
		baseWhere = baseWhere.Where("workshop_id = ?", workshopID)
	}
	if startDate != nil {
		baseWhere = baseWhere.Where("call_time >= ?", startDate)
	}
	if endDate != nil {
		baseWhere = baseWhere.Where("call_time <= ?", endDate)
	}

	// 总呼叫数
	baseWhere.Count(&totalCalls)

	// 已解决数量
	r.db.WithContext(ctx).Model(&model.AndonCall{}).
		Where("status IN ?", []string{"RESOLVED", "CLOSED"}).
		Count(&resolvedCalls)

	// 升级数量
	r.db.WithContext(ctx).Model(&model.AndonCall{}).
		Where("is_escalated = ?", 1).
		Count(&escalatedCalls)

	// 平均响应时间
	var avgResponse sql.NullFloat64
	r.db.WithContext(ctx).Model(&model.AndonCall{}).
		Select("AVG(response_duration)").
		Where("response_duration > 0").
		Scan(&avgResponse)

	// 平均处理时间
	var avgHandle sql.NullFloat64
	r.db.WithContext(ctx).Model(&model.AndonCall{}).
		Select("AVG(handle_duration)").
		Where("handle_duration > 0").
		Scan(&avgHandle)

	// 按类型统计
	type TypeStats struct {
		AndonType string `json:"andon_type"`
		Count     int64  `json:"count"`
	}
	var byType []TypeStats
	r.db.WithContext(ctx).Model(&model.AndonCall{}).
		Select("andon_type, COUNT(*) as count").
		Group("andon_type").
		Scan(&byType)

	// 按小时分布
	type HourStats struct {
		Hour  int   `json:"hour"`
		Count int64 `json:"count"`
	}
	var byHour []HourStats
	r.db.WithContext(ctx).Model(&model.AndonCall{}).
		Select("EXTRACT(HOUR FROM call_time) as hour, COUNT(*) as count").
		Group("EXTRACT(HOUR FROM call_time)").
		Order("hour").
		Scan(&byHour)

	// 未结束呼叫
	var unresolvedCalls []model.AndonCall
	r.db.WithContext(ctx).
		Where("status IN ?", []string{"CALLING", "RESPONDED", "HANDLING"}).
		Order("call_time DESC").
		Limit(10).
		Find(&unresolvedCalls)

	result := map[string]interface{}{
		"total_calls":       totalCalls,
		"resolved_calls":    resolvedCalls,
		"escalated_calls":   escalatedCalls,
		"avg_response_time": avgResponse.Float64,
		"avg_handle_time":   avgHandle.Float64,
		"escalation_rate":   0.0,
		"call_by_type":      byType,
		"call_by_hour":      byHour,
		"unresolved_calls": unresolvedCalls,
	}

	if totalCalls > 0 {
		result["escalation_rate"] = float64(escalatedCalls) / float64(totalCalls) * 100
	}

	return result, nil
}

// ========== EscalationRule Repository ==========

type EscalationRuleRepository struct {
	db *gorm.DB
}

func NewEscalationRuleRepository(db *gorm.DB) *EscalationRuleRepository {
	return &EscalationRuleRepository{db: db}
}

// List 查询规则列表
func (r *EscalationRuleRepository) List(ctx context.Context, tenantID int64, andonType string, workshopID int64) ([]model.AndonEscalationRule, int64, error) {
	var list []model.AndonEscalationRule
	var total int64

	query := r.db.WithContext(ctx).Model(&model.AndonEscalationRule{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if andonType != "" {
		query = query.Where("andon_type = ? OR andon_type = '' OR andon_type IS NULL", andonType)
	}
	if workshopID > 0 {
		query = query.Where("workshop_id = ? OR workshop_id IS NULL OR workshop_id = 0", workshopID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("sort_order ASC, id ASC").Find(&list).Error
	return list, total, err
}

// GetByID 根据ID获取
func (r *EscalationRuleRepository) GetByID(ctx context.Context, id int64) (*model.AndonEscalationRule, error) {
	var rule model.AndonEscalationRule
	err := r.db.WithContext(ctx).First(&rule, id).Error
	return &rule, err
}

// GetDefaultRule 获取默认规则
func (r *EscalationRuleRepository) GetDefaultRule(ctx context.Context, tenantID int64) (*model.AndonEscalationRule, error) {
	var rule model.AndonEscalationRule
	err := r.db.WithContext(ctx).
		Where("tenant_id = ? AND is_enabled = 1 AND is_default = 1", tenantID).
		First(&rule).Error
	if err == gorm.ErrRecordNotFound {
		// 如果没有默认规则，返回任意一个启用规则
		err = r.db.WithContext(ctx).
			Where("tenant_id = ? AND is_enabled = 1", tenantID).
			First(&rule).Error
	}
	return &rule, err
}

// GetApplicableRule 获取适用的升级规则
func (r *EscalationRuleRepository) GetApplicableRule(ctx context.Context, tenantID, workshopID int64, andonType string, priority int) (*model.AndonEscalationRule, error) {
	var rule model.AndonEscalationRule

	// 优先查找完全匹配的规则
	err := r.db.WithContext(ctx).
		Where("tenant_id = ? AND is_enabled = 1", tenantID).
		Where("andon_type = ? AND workshop_id = ?", andonType, workshopID).
		Where("is_default = 0").
		First(&rule).Error

	if err == gorm.ErrRecordNotFound {
		// 查找类型匹配但车间不限制的规则
		err = r.db.WithContext(ctx).
			Where("tenant_id = ? AND is_enabled = 1", tenantID).
			Where("andon_type = ? AND (workshop_id IS NULL OR workshop_id = 0)", andonType).
			First(&rule).Error
	}

	if err == gorm.ErrRecordNotFound {
		// 查找默认规则(不限类型和车间)
		err = r.db.WithContext(ctx).
			Where("tenant_id = ? AND is_enabled = 1 AND is_default = 1", tenantID).
			First(&rule).Error
	}

	if err == gorm.ErrRecordNotFound {
		// 查找任意启用规则
		err = r.db.WithContext(ctx).
			Where("tenant_id = ? AND is_enabled = 1", tenantID).
			First(&rule).Error
	}

	if err != nil {
		return nil, err
	}

	return &rule, nil
}

// Create 创建规则
func (r *EscalationRuleRepository) Create(ctx context.Context, rule *model.AndonEscalationRule) error {
	return r.db.WithContext(ctx).Create(rule).Error
}

// Update 更新规则
func (r *EscalationRuleRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.AndonEscalationRule{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除规则
func (r *EscalationRuleRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.AndonEscalationRule{}, id).Error
}

// ========== EscalationLog Repository ==========

type EscalationLogRepository struct {
	db *gorm.DB
}

func NewEscalationLogRepository(db *gorm.DB) *EscalationLogRepository {
	return &EscalationLogRepository{db: db}
}

// Create 创建升级日志
func (r *EscalationLogRepository) Create(ctx context.Context, log *model.AndonEscalationLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

// ListByCallID 根据呼叫ID获取升级历史
func (r *EscalationLogRepository) ListByCallID(ctx context.Context, callID int64) ([]model.AndonEscalationLog, error) {
	var list []model.AndonEscalationLog
	err := r.db.WithContext(ctx).
		Where("call_id = ?", callID).
		Order("created_at ASC").
		Find(&list).Error
	return list, err
}

// ========== NotificationLog Repository ==========

type NotificationLogRepository struct {
	db *gorm.DB
}

func NewNotificationLogRepository(db *gorm.DB) *NotificationLogRepository {
	return &NotificationLogRepository{db: db}
}

// Create 创建通知日志
func (r *NotificationLogRepository) Create(ctx context.Context, log *model.AndonNotificationLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

// ListByCallID 根据呼叫ID获取通知历史
func (r *NotificationLogRepository) ListByCallID(ctx context.Context, callID int64) ([]model.AndonNotificationLog, error) {
	var list []model.AndonNotificationLog
	err := r.db.WithContext(ctx).
		Where("call_id = ?", callID).
		Order("created_at DESC").
		Find(&list).Error
	return list, err
}
