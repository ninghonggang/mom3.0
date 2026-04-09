package ai

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"mom-server/internal/service"
	"mom-server/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type AIConfigHandler struct {
	configRepo *repository.AIConfigRepository
	aiService  *service.AIService
}

func NewAIConfigHandler(configRepo *repository.AIConfigRepository, aiService *service.AIService) *AIConfigHandler {
	return &AIConfigHandler{
		configRepo: configRepo,
		aiService:  aiService,
	}
}

// GetConfig 获取当前租户的AI配置
// GET /api/v1/ai/config
func (h *AIConfigHandler) GetConfig(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	cfg, err := h.configRepo.GetByTenant(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, "AI配置未找到，请先配置")
		return
	}
	// Mask API key for security
	if cfg.APIKey != "" {
		cfg.APIKey = "******"
	}
	response.Success(c, cfg)
}

// UpdateConfig 更新AI配置
// PUT /api/v1/ai/config
func (h *AIConfigHandler) UpdateConfig(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)

	var req model.AIConfig
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req.TenantID = tenantID

	// 如果APIKey被设置为******则不更新
	if req.APIKey == "******" {
		existing, _ := h.configRepo.GetByTenant(c.Request.Context(), tenantID)
		if existing != nil {
			req.APIKey = existing.APIKey
		}
	}

	// 检查是否已存在
	existing, _ := h.configRepo.GetByTenant(c.Request.Context(), tenantID)
	if existing != nil {
		req.ID = existing.ID
		if err := h.configRepo.Update(c.Request.Context(), &req); err != nil {
			response.ErrorMsg(c, err.Error())
			return
		}
	} else {
		if err := h.configRepo.Create(c.Request.Context(), &req); err != nil {
			response.ErrorMsg(c, err.Error())
			return
		}
	}

	// 返回时mask API key
	req.APIKey = "******"
	response.Success(c, req)
}

// TestConfig 测试AI配置
// POST /api/v1/ai/config/test
func (h *AIConfigHandler) TestConfig(c *gin.Context) {
	var req model.AIConfig
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.aiService.TestConnection(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, "连接失败: "+err.Error())
		return
	}

	response.Success(c, "连接成功")
}

// GetSchema 获取系统Schema
// GET /api/v1/ai/schema
func (h *AIConfigHandler) GetSchema(c *gin.Context) {
	schema := h.aiService.BuildSystemPrompt()
	response.Success(c, gin.H{"schema": schema})
}
