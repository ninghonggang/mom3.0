package andon

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type RuleHandler struct {
	ruleSvc *service.EscalationRuleService
}

func NewRuleHandler(ruleSvc *service.EscalationRuleService) *RuleHandler {
	return &RuleHandler{ruleSvc: ruleSvc}
}

// List 查询规则列表
// GET /api/v1/andon/escalation-rules
func (h *RuleHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	andonType := c.Query("andon_type")
	workshopID, _ := strconv.ParseInt(c.Query("workshop_id"), 10, 64)

	list, total, err := h.ruleSvc.List(c.Request.Context(), tenantID, andonType, workshopID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

// Get 获取单个规则
// GET /api/v1/andon/escalation-rules/:id
func (h *RuleHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	rule, err := h.ruleSvc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, rule)
}

// Create 创建规则
// POST /api/v1/andon/escalation-rules
func (h *RuleHandler) Create(c *gin.Context) {
	var req model.AndonEscalationRule
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID

	// 获取创建人
	req.CreatedBy = middleware.GetUsername(c)

	// 设置默认升级模式
	if req.EscalationMode == "" {
		req.EscalationMode = "TIMEOUT"
	}
	// 设置最大升级等级
	if req.MaxEscalationLevel == 0 {
		req.MaxEscalationLevel = 4
	}

	if err := h.ruleSvc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// Update 更新规则
// PUT /api/v1/andon/escalation-rules/:id
func (h *RuleHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.AndonEscalationRule
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.ruleSvc.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// Delete 删除规则
// DELETE /api/v1/andon/escalation-rules/:id
func (h *RuleHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.ruleSvc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
