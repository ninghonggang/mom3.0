package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// BpmTaskMessageRuleRepository BPM任务消息规则仓库
type BpmTaskMessageRuleRepository struct {
	db *gorm.DB
}

func NewBpmTaskMessageRuleRepository(db *gorm.DB) *BpmTaskMessageRuleRepository {
	return &BpmTaskMessageRuleRepository{db: db}
}

func (r *BpmTaskMessageRuleRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.BpmTaskMessageRule, int64, error) {
	var list []model.BpmTaskMessageRule
	var total int64

	q := r.db.WithContext(ctx).Model(&model.BpmTaskMessageRule{}).Where("tenant_id = ?", tenantID)

	if ruleCode, ok := query["rule_code"]; ok && ruleCode != "" {
		q = q.Where("rule_code LIKE ?", "%"+ruleCode.(string)+"%")
	}
	if ruleName, ok := query["rule_name"]; ok && ruleName != "" {
		q = q.Where("rule_name LIKE ?", "%"+ruleName.(string)+"%")
	}
	if processDefKey, ok := query["process_def_key"]; ok && processDefKey != "" {
		q = q.Where("process_def_key = ?", processDefKey)
	}
	if messageType, ok := query["message_type"]; ok && messageType != "" {
		q = q.Where("message_type = ?", messageType)
	}
	if isEnabled, ok := query["is_enabled"]; ok {
		q = q.Where("is_enabled = ?", isEnabled)
	}

	q.Count(&total)
	q = q.Order("priority DESC, id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *BpmTaskMessageRuleRepository) GetByID(ctx context.Context, id uint) (*model.BpmTaskMessageRule, error) {
	var rule model.BpmTaskMessageRule
	err := r.db.WithContext(ctx).First(&rule, id).Error
	return &rule, err
}

func (r *BpmTaskMessageRuleRepository) GetByCode(ctx context.Context, tenantID int64, code string) (*model.BpmTaskMessageRule, error) {
	var rule model.BpmTaskMessageRule
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND rule_code = ?", tenantID, code).First(&rule).Error
	return &rule, err
}

func (r *BpmTaskMessageRuleRepository) Create(ctx context.Context, rule *model.BpmTaskMessageRule) error {
	return r.db.WithContext(ctx).Create(rule).Error
}

func (r *BpmTaskMessageRuleRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.BpmTaskMessageRule{}).Where("id = ?", id).Updates(updates).Error
}

func (r *BpmTaskMessageRuleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.BpmTaskMessageRule{}, id).Error
}

func (r *BpmTaskMessageRuleRepository) Enable(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&model.BpmTaskMessageRule{}).Where("id = ?", id).Update("is_enabled", true).Error
}

func (r *BpmTaskMessageRuleRepository) Disable(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&model.BpmTaskMessageRule{}).Where("id = ?", id).Update("is_enabled", false).Error
}
