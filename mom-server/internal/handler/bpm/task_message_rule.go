package bpm

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type BpmTaskMessageRuleHandler struct {
	svc *service.BpmTaskMessageRuleService
}

func NewBpmTaskMessageRuleHandler(svc *service.BpmTaskMessageRuleService) *BpmTaskMessageRuleHandler {
	return &BpmTaskMessageRuleHandler{svc: svc}
}

func (h *BpmTaskMessageRuleHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if ruleCode := c.Query("rule_code"); ruleCode != "" {
		query["rule_code"] = ruleCode
	}
	if ruleName := c.Query("rule_name"); ruleName != "" {
		query["rule_name"] = ruleName
	}
	if processDefKey := c.Query("process_def_key"); processDefKey != "" {
		query["process_def_key"] = processDefKey
	}
	if messageType := c.Query("message_type"); messageType != "" {
		query["message_type"] = messageType
	}

	list, total, err := h.svc.List(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *BpmTaskMessageRuleHandler) Get(c *gin.Context) {
	id := c.Param("id")
	rule, err := h.svc.Get(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, rule)
}

func (h *BpmTaskMessageRuleHandler) Create(c *gin.Context) {
	var req model.BpmTaskMessageRuleCreateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	rule := &model.BpmTaskMessageRule{
		TenantID:      tenantID,
		RuleCode:      req.RuleCode,
		RuleName:      req.RuleName,
		ProcessDefKey: req.ProcessDefKey,
		TaskDefKey:    req.TaskDefKey,
		MessageType:   req.MessageType,
		TemplateCode:  req.TemplateCode,
		IsEnabled:     req.IsEnabled,
		Priority:      req.Priority,
		Remark:        req.Remark,
	}

	if err := h.svc.Create(c.Request.Context(), rule); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, rule)
}

func (h *BpmTaskMessageRuleHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.BpmTaskMessageRuleUpdateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *BpmTaskMessageRuleHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *BpmTaskMessageRuleHandler) Enable(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.Enable(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *BpmTaskMessageRuleHandler) Disable(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.Disable(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
