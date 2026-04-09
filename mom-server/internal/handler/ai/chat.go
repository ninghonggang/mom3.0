package ai

import (
	"encoding/json"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"mom-server/internal/service"
	"mom-server/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type AIChatHandler struct {
	aiService  *service.AIService
	aiExecutor *service.AIExecutor
	convRepo   *repository.ConversationRepository
	msgRepo    *repository.MessageRepository
}

func NewAIChatHandler(
	aiService *service.AIService,
	aiExecutor *service.AIExecutor,
	convRepo *repository.ConversationRepository,
	msgRepo *repository.MessageRepository,
) *AIChatHandler {
	return &AIChatHandler{
		aiService:  aiService,
		aiExecutor: aiExecutor,
		convRepo:   convRepo,
		msgRepo:    msgRepo,
	}
}

// SendMessageRequest 发送消息请求
type SendMessageRequest struct {
	SessionID string `json:"session_id"` // 空表示新会话
	Message   string `json:"message" binding:"required"`
}

// SendMessage 发送消息
// POST /api/v1/ai/chat/send
func (h *AIChatHandler) SendMessage(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	userID := middleware.GetUserID(c)

	var req SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// 获取对话历史用于上下文
	var history []*model.ChatMessage
	if req.SessionID != "" {
		msgs, _ := h.aiService.GetConversationMessages(c.Request.Context(), tenantID, req.SessionID, 20)
		history = make([]*model.ChatMessage, len(msgs))
		for i := range msgs {
			history[i] = &msgs[i]
		}
	}

	resp, err := h.aiService.SendMessage(c.Request.Context(), tenantID, userID, req.SessionID, req.Message, history)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, resp)
}

// ExecuteOperationRequest 执行操作请求
type ExecuteOperationRequest struct {
	MessageID int64 `json:"message_id" binding:"required"`
	Confirmed bool  `json:"confirmed"`
}

// ExecuteOperation 执行确认的写操作
// POST /api/v1/ai/chat/execute
func (h *AIChatHandler) ExecuteOperation(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	userID := middleware.GetUserID(c)

	var req ExecuteOperationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if !req.Confirmed {
		response.ErrorMsg(c, "操作已取消")
		return
	}

	// 获取消息
	msg, err := h.msgRepo.GetByID(c.Request.Context(), req.MessageID)
	if err != nil {
		response.ErrorMsg(c, "消息不存在")
		return
	}

	if msg.IntentJSON == nil {
		response.ErrorMsg(c, "无操作意图")
		return
	}

	// 解析意图
	var intent service.AIIntent
	if err := json.Unmarshal([]byte(*msg.IntentJSON), &intent); err != nil {
		response.ErrorMsg(c, "解析意图失败")
		return
	}

	result, err := h.aiExecutor.ExecuteConfirmedOperation(c.Request.Context(), &intent, tenantID, userID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	// 更新消息状态
	status := "executed"
	if !result.Success {
		status = "rejected"
	}
	h.msgRepo.UpdateStatus(c.Request.Context(), req.MessageID, status)

	// 如果有结果数据，更新工具结果
	if result.Data != nil {
		dataJSON, _ := json.Marshal(result.Data)
		h.msgRepo.UpdateToolResult(c.Request.Context(), req.MessageID, string(dataJSON))
	}

	response.Success(c, result)
}

// ListConversations 获取会话列表
// GET /api/v1/ai/chat/conversations
func (h *AIChatHandler) ListConversations(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	userID := middleware.GetUserID(c)

	convs, err := h.aiService.GetConversations(c.Request.Context(), tenantID, userID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, gin.H{"list": convs})
}

// GetConversation 获取会话详情
// GET /api/v1/ai/chat/conversations/:session_id
func (h *AIChatHandler) GetConversation(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	sessionID := c.Param("session_id")

	msgs, err := h.aiService.GetConversationMessages(c.Request.Context(), tenantID, sessionID, 100)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, gin.H{"list": msgs})
}

// DeleteConversation 删除会话
// DELETE /api/v1/ai/chat/conversations/:session_id
func (h *AIChatHandler) DeleteConversation(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	sessionID := c.Param("session_id")

	if err := h.aiService.DeleteConversation(c.Request.Context(), tenantID, sessionID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}
