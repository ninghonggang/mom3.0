package model

import (
	"time"

	"gorm.io/gorm"
)

// AIConfig AI模型配置
type AIConfig struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID    int64     `json:"tenant_id" gorm:"index;not null"`
	ConfigName  string    `json:"config_name" gorm:"size:100"`
	Provider    string    `json:"provider" gorm:"size:50"` // openai/azure/ollama/custom
	Endpoint    string    `json:"endpoint" gorm:"size:500"`
	APIVersion  string    `json:"api_version" gorm:"size:100"` // For Azure OpenAI
	ModelName   string    `json:"model_name" gorm:"size:100"`
	APIKey      string    `json:"api_key" gorm:"size:500"` // Encrypted, frontend sends api_key in requests
	Temperature float64   `json:"temperature" gorm:"type:decimal(3,2)"`
	MaxTokens   int       `json:"max_tokens"`
	Timeout     int       `json:"timeout"` // Seconds
	Enable      bool      `json:"enable" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (AIConfig) TableName() string {
	return "ai_config"
}

// ChatConversation AI聊天会话
type ChatConversation struct {
	ID        int64          `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID  int64          `json:"tenant_id" gorm:"index;not null"`
	UserID    int64          `json:"user_id" gorm:"index;not null"`
	SessionID string         `json:"session_id" gorm:"type:uuid;index:idx_tenant_session,unique"`
	Title     *string        `json:"title" gorm:"size:200"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (ChatConversation) TableName() string {
	return "ai_chat_conversation"
}

// ChatMessage AI聊天消息
type ChatMessage struct {
	ID             int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID       int64     `json:"tenant_id" gorm:"index;not null"`
	UserID         int64     `json:"user_id" gorm:"index;not null"`
	ConversationID int64     `json:"conversation_id" gorm:"index;not null"`
	Role           string    `json:"role" gorm:"size:20"` // user/assistant/system
	Content        string    `json:"content" gorm:"type:text"`
	IntentJSON     *string   `json:"intent_json" gorm:"type:text"`    // Structured intent
	OperationType  *string   `json:"operation_type" gorm:"size:20"` // query/write/analysis/null
	Status         string    `json:"status" gorm:"size:20;default:pending"` // pending/executed/confirmed/rejected
	ToolResult     *string   `json:"tool_result" gorm:"type:text"`
	CreatedAt      time.Time `json:"created_at"`
}

func (ChatMessage) TableName() string {
	return "ai_chat_message"
}
