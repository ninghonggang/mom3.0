package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// AlertRuleRepository 告警规则仓库
type AlertRuleRepository struct {
	db *gorm.DB
}

func NewAlertRuleRepository(db *gorm.DB) *AlertRuleRepository {
	return &AlertRuleRepository{db: db}
}

func (r *AlertRuleRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.AlertRule, int64, error) {
	var list []model.AlertRule
	var total int64

	q := r.db.WithContext(ctx).Model(&model.AlertRule{}).Where("tenant_id = ?", tenantID)

	if alertType, ok := query["alert_type"]; ok && alertType != "" {
		q = q.Where("alert_type = ?", alertType)
	}
	if bizModule, ok := query["biz_module"]; ok && bizModule != "" {
		q = q.Where("biz_module = ?", bizModule)
	}
	if severity, ok := query["severity_level"]; ok && severity != "" {
		q = q.Where("severity_level = ?", severity)
	}
	if isEnabled, ok := query["is_enabled"]; ok && isEnabled != "" {
		q = q.Where("is_enabled = ?", isEnabled)
	}

	q.Count(&total)
	q = q.Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *AlertRuleRepository) GetByID(ctx context.Context, id uint) (*model.AlertRule, error) {
	var rule model.AlertRule
	err := r.db.WithContext(ctx).First(&rule, id).Error
	return &rule, err
}

func (r *AlertRuleRepository) GetByCode(ctx context.Context, tenantID int64, code string) (*model.AlertRule, error) {
	var rule model.AlertRule
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND rule_code = ?", tenantID, code).First(&rule).Error
	return &rule, err
}

func (r *AlertRuleRepository) Create(ctx context.Context, rule *model.AlertRule) error {
	return r.db.WithContext(ctx).Create(rule).Error
}

func (r *AlertRuleRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.AlertRule{}).Where("id = ?", id).Updates(updates).Error
}

func (r *AlertRuleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.AlertRule{}, id).Error
}

// AlertRecordRepository 告警记录仓库
type AlertRecordRepository struct {
	db *gorm.DB
}

func NewAlertRecordRepository(db *gorm.DB) *AlertRecordRepository {
	return &AlertRecordRepository{db: db}
}

func (r *AlertRecordRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.AlertRecord, int64, error) {
	var list []model.AlertRecord
	var total int64

	q := r.db.WithContext(ctx).Model(&model.AlertRecord{}).Where("tenant_id = ?", tenantID)

	if alertType, ok := query["alert_type"]; ok && alertType != "" {
		q = q.Where("alert_type = ?", alertType)
	}
	if severity, ok := query["severity_level"]; ok && severity != "" {
		q = q.Where("severity_level = ?", severity)
	}
	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}
	if sourceModule, ok := query["source_module"]; ok && sourceModule != "" {
		q = q.Where("source_module = ?", sourceModule)
	}
	if startDate, ok := query["start_date"]; ok && startDate != "" {
		q = q.Where("trigger_time >= ?", startDate)
	}
	if endDate, ok := query["end_date"]; ok && endDate != "" {
		q = q.Where("trigger_time <= ?", endDate)
	}
	if keyword, ok := query["keyword"]; ok && keyword != "" {
		q = q.Where("title LIKE ? OR content LIKE ?", "%"+keyword.(string)+"%", "%"+keyword.(string)+"%")
	}

	q.Count(&total)
	q = q.Order("id DESC")

	page := 1
	limit := 20
	if p, ok := query["page"].(int); ok && p > 0 {
		page = p
	}
	if l, ok := query["limit"].(int); ok && l > 0 {
		limit = l
	}
	q = q.Offset((page - 1) * limit).Limit(limit)

	err := q.Find(&list).Error
	return list, total, err
}

func (r *AlertRecordRepository) GetByID(ctx context.Context, id uint) (*model.AlertRecord, error) {
	var record model.AlertRecord
	err := r.db.WithContext(ctx).First(&record, id).Error
	return &record, err
}

func (r *AlertRecordRepository) Create(ctx context.Context, record *model.AlertRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *AlertRecordRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.AlertRecord{}).Where("id = ?", id).Updates(updates).Error
}

func (r *AlertRecordRepository) GetStatistics(ctx context.Context, tenantID int64) (map[string]interface{}, error) {
	var todayCount, unresolvedCount, criticalCount int64
	var avgResponseTime, avgResolutionTime float64

	r.db.WithContext(ctx).Model(&model.AlertRecord{}).Where("tenant_id = ? AND DATE(trigger_time) = CURRENT_DATE", tenantID).Count(&todayCount)
	r.db.WithContext(ctx).Model(&model.AlertRecord{}).Where("tenant_id = ? AND status IN ('TRIGGERED', 'ACKNOWLEDGED')", tenantID).Count(&unresolvedCount)
	r.db.WithContext(ctx).Model(&model.AlertRecord{}).Where("tenant_id = ? AND severity_level = 'CRITICAL' AND status IN ('TRIGGERED', 'ACKNOWLEDGED')", tenantID).Count(&criticalCount)

	return map[string]interface{}{
		"today_count":        todayCount,
		"unresolved_count":   unresolvedCount,
		"critical_count":     criticalCount,
		"avg_response_time":  avgResponseTime,
		"avg_resolution_time": avgResolutionTime,
	}, nil
}

// AlertNotificationLogRepository 告警通知日志仓库
type AlertNotificationLogRepository struct {
	db *gorm.DB
}

func NewAlertNotificationLogRepository(db *gorm.DB) *AlertNotificationLogRepository {
	return &AlertNotificationLogRepository{db: db}
}

func (r *AlertNotificationLogRepository) ListByAlertID(ctx context.Context, alertID uint) ([]model.AlertNotificationLog, error) {
	var list []model.AlertNotificationLog
	err := r.db.WithContext(ctx).Where("alert_id = ?", alertID).Order("created_at DESC").Find(&list).Error
	return list, err
}

func (r *AlertNotificationLogRepository) Create(ctx context.Context, log *model.AlertNotificationLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

// AlertEscalationRuleRepository 告警升级规则仓库
type AlertEscalationRuleRepository struct {
	db *gorm.DB
}

func NewAlertEscalationRuleRepository(db *gorm.DB) *AlertEscalationRuleRepository {
	return &AlertEscalationRuleRepository{db: db}
}

func (r *AlertEscalationRuleRepository) List(ctx context.Context, tenantID int64) ([]model.AlertEscalationRule, error) {
	var list []model.AlertEscalationRule
	err := r.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Order("id DESC").Find(&list).Error
	return list, err
}

func (r *AlertEscalationRuleRepository) GetByID(ctx context.Context, id uint) (*model.AlertEscalationRule, error) {
	var rule model.AlertEscalationRule
	err := r.db.WithContext(ctx).First(&rule, id).Error
	return &rule, err
}

func (r *AlertEscalationRuleRepository) Create(ctx context.Context, rule *model.AlertEscalationRule) error {
	return r.db.WithContext(ctx).Create(rule).Error
}

func (r *AlertEscalationRuleRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.AlertEscalationRule{}).Where("id = ?", id).Updates(updates).Error
}

// NotificationChannelRepository 通知渠道配置仓库
type NotificationChannelRepository struct {
	db *gorm.DB
}

func NewNotificationChannelRepository(db *gorm.DB) *NotificationChannelRepository {
	return &NotificationChannelRepository{db: db}
}

func (r *NotificationChannelRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.NotificationChannel, int64, error) {
	var list []model.NotificationChannel
	var total int64

	q := r.db.WithContext(ctx).Model(&model.NotificationChannel{}).Where("tenant_id = ?", tenantID)

	if channelType, ok := query["channel_type"]; ok && channelType != "" {
		q = q.Where("channel_type = ?", channelType)
	}
	if isEnabled, ok := query["is_enabled"]; ok && isEnabled != "" {
		q = q.Where("is_enabled = ?", isEnabled)
	}

	q.Count(&total)
	q = q.Order("priority DESC, id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *NotificationChannelRepository) GetByID(ctx context.Context, id uint) (*model.NotificationChannel, error) {
	var ch model.NotificationChannel
	err := r.db.WithContext(ctx).First(&ch, id).Error
	return &ch, err
}

func (r *NotificationChannelRepository) GetByType(ctx context.Context, tenantID int64, channelType string) ([]model.NotificationChannel, error) {
	var list []model.NotificationChannel
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND channel_type = ? AND is_enabled = 1", tenantID, channelType).Order("priority DESC").Find(&list).Error
	return list, err
}

func (r *NotificationChannelRepository) GetByCode(ctx context.Context, tenantID int64, code string) (*model.NotificationChannel, error) {
	var ch model.NotificationChannel
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND channel_code = ?", tenantID, code).First(&ch).Error
	return &ch, err
}

func (r *NotificationChannelRepository) Create(ctx context.Context, ch *model.NotificationChannel) error {
	return r.db.WithContext(ctx).Create(ch).Error
}

func (r *NotificationChannelRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.NotificationChannel{}).Where("id = ?", id).Updates(updates).Error
}

func (r *NotificationChannelRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.NotificationChannel{}, id).Error
}

func (r *NotificationChannelRepository) Enable(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&model.NotificationChannel{}).Where("id = ?", id).Updates(map[string]interface{}{"is_enabled": 1}).Error
}

func (r *NotificationChannelRepository) Disable(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&model.NotificationChannel{}).Where("id = ?", id).Updates(map[string]interface{}{"is_enabled": 0}).Error
}
