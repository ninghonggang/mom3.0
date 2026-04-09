package production

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ElectronicSOPHandler struct {
	service *service.ElectronicSOPService
}

func NewElectronicSOPHandler(s *service.ElectronicSOPService) *ElectronicSOPHandler {
	return &ElectronicSOPHandler{service: s}
}

func (h *ElectronicSOPHandler) List(c *gin.Context) {
	query := c.Query("query")
	list, total, err := h.service.List(c.Request.Context(), query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *ElectronicSOPHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	sop, err := h.service.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, sop)
}

func (h *ElectronicSOPHandler) Create(c *gin.Context) {
	var req model.ElectronicSOP
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *ElectronicSOPHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var updates map[string]any
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *ElectronicSOPHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

type CodeRuleHandler struct {
	service *service.CodeRuleService
}

func NewCodeRuleHandler(s *service.CodeRuleService) *CodeRuleHandler {
	return &CodeRuleHandler{service: s}
}

func (h *CodeRuleHandler) List(c *gin.Context) {
	list, err := h.service.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list})
}

func (h *CodeRuleHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	rule, err := h.service.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, rule)
}

func (h *CodeRuleHandler) Create(c *gin.Context) {
	var req model.CodeRule
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *CodeRuleHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var updates map[string]any
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *CodeRuleHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *CodeRuleHandler) GenerateCode(c *gin.Context) {
	ruleCode := c.Query("rule_code")
	if ruleCode == "" {
		response.BadRequest(c, "rule_code is required")
		return
	}
	code, err := h.service.GenerateCode(c.Request.Context(), ruleCode)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"code": code})
}

type FlowCardHandler struct {
	service *service.FlowCardService
}

func NewFlowCardHandler(s *service.FlowCardService) *FlowCardHandler {
	return &FlowCardHandler{service: s}
}

func (h *FlowCardHandler) List(c *gin.Context) {
	query := c.Query("query")
	list, total, err := h.service.List(c.Request.Context(), query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *FlowCardHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	card, err := h.service.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	details, _ := h.service.GetDetails(c.Request.Context(), int64(card.ID))
	response.Success(c, gin.H{"card": card, "details": details})
}

func (h *FlowCardHandler) Create(c *gin.Context) {
	var req struct {
		Card    model.FlowCard           `json:"card"`
		Details []model.FlowCardDetail   `json:"details"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.Card.TenantID = tenantID
	if err := h.service.CreateWithDetails(c.Request.Context(), &req.Card, req.Details); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req.Card)
}

func (h *FlowCardHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var updates map[string]any
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *FlowCardHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
