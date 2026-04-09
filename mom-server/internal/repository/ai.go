package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// AIConfigRepository AI模型配置仓储
type AIConfigRepository struct {
	db *gorm.DB
}

// NewAIConfigRepository 创建AI配置仓储
func NewAIConfigRepository(db *gorm.DB) *AIConfigRepository {
	return &AIConfigRepository{db: db}
}

// GetByTenant 根据租户获取启用的配置
func (r *AIConfigRepository) GetByTenant(ctx context.Context, tenantID int64) (*model.AIConfig, error) {
	var cfg model.AIConfig
	err := r.db.WithContext(ctx).
		Where("tenant_id = ? AND enable = ?", tenantID, true).
		First(&cfg).Error
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// GetByID 根据ID查询
func (r *AIConfigRepository) GetByID(ctx context.Context, id int64) (*model.AIConfig, error) {
	var cfg model.AIConfig
	err := r.db.WithContext(ctx).First(&cfg, id).Error
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Create 创建配置
func (r *AIConfigRepository) Create(ctx context.Context, cfg *model.AIConfig) error {
	return r.db.WithContext(ctx).Create(cfg).Error
}

// Update 更新配置
func (r *AIConfigRepository) Update(ctx context.Context, cfg *model.AIConfig) error {
	return r.db.WithContext(ctx).Save(cfg).Error
}

// Delete 删除配置
func (r *AIConfigRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.AIConfig{}, id).Error
}

// ListByTenant 查询租户下所有配置
func (r *AIConfigRepository) ListByTenant(ctx context.Context, tenantID int64) ([]model.AIConfig, error) {
	var cfgs []model.AIConfig
	err := r.db.WithContext(ctx).
		Where("tenant_id = ?", tenantID).
		Order("created_at DESC").
		Find(&cfgs).Error
	return cfgs, err
}

// ConversationRepository AI聊天会话仓储
type ConversationRepository struct {
	db *gorm.DB
}

// NewConversationRepository 创建会话仓储
func NewConversationRepository(db *gorm.DB) *ConversationRepository {
	return &ConversationRepository{db: db}
}

// ListByUser 查询用户会话列表
func (r *ConversationRepository) ListByUser(ctx context.Context, tenantID, userID int64, limit int) ([]model.ChatConversation, error) {
	var convs []model.ChatConversation
	query := r.db.WithContext(ctx).
		Where("tenant_id = ? AND user_id = ? AND deleted_at IS NULL", tenantID, userID).
		Order("created_at DESC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	err := query.Find(&convs).Error
	return convs, err
}

// GetBySessionID 根据SessionID查询
func (r *ConversationRepository) GetBySessionID(ctx context.Context, tenantID int64, sessionID string) (*model.ChatConversation, error) {
	var conv model.ChatConversation
	err := r.db.WithContext(ctx).
		Where("tenant_id = ? AND session_id = ? AND deleted_at IS NULL", tenantID, sessionID).
		First(&conv).Error
	if err != nil {
		return nil, err
	}
	return &conv, nil
}

// Create 创建会话
func (r *ConversationRepository) Create(ctx context.Context, conv *model.ChatConversation) error {
	return r.db.WithContext(ctx).Create(conv).Error
}

// Update 更新会话
func (r *ConversationRepository) Update(ctx context.Context, conv *model.ChatConversation) error {
	return r.db.WithContext(ctx).Save(conv).Error
}

// Delete 删除会话(软删除)
func (r *ConversationRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.ChatConversation{}, id).Error
}

// MessageRepository AI聊天消息仓储
type MessageRepository struct {
	db *gorm.DB
}

// NewMessageRepository 创建消息仓储
func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

// ListByConversation 查询会话消息列表
func (r *MessageRepository) ListByConversation(ctx context.Context, conversationID int64, limit int) ([]model.ChatMessage, error) {
	var msgs []model.ChatMessage
	query := r.db.WithContext(ctx).
		Where("conversation_id = ?", conversationID).
		Order("created_at ASC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	err := query.Find(&msgs).Error
	return msgs, err
}

// Create 创建消息
func (r *MessageRepository) Create(ctx context.Context, msg *model.ChatMessage) error {
	return r.db.WithContext(ctx).Create(msg).Error
}

// UpdateStatus 更新消息状态
func (r *MessageRepository) UpdateStatus(ctx context.Context, id int64, status string) error {
	return r.db.WithContext(ctx).
		Model(&model.ChatMessage{}).
		Where("id = ?", id).
		Update("status", status).Error
}

// UpdateToolResult 更新工具执行结果
func (r *MessageRepository) UpdateToolResult(ctx context.Context, id int64, result string) error {
	return r.db.WithContext(ctx).
		Model(&model.ChatMessage{}).
		Where("id = ?", id).
		Update("tool_result", result).Error
}

// GetByID 根据ID查询消息
func (r *MessageRepository) GetByID(ctx context.Context, id int64) (*model.ChatMessage, error) {
	var msg model.ChatMessage
	err := r.db.WithContext(ctx).First(&msg, id).Error
	if err != nil {
		return nil, err
	}
	return &msg, nil
}
