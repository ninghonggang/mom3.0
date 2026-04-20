package production

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductionIssueHandler struct {
	issueSvc *service.ProductionIssueService
}

func NewProductionIssueHandler(issueSvc *service.ProductionIssueService) *ProductionIssueHandler {
	return &ProductionIssueHandler{issueSvc: issueSvc}
}

func toUint(s string) uint {
	v, _ := strconv.ParseUint(s, 10, 64)
	return uint(v)
}

// ListProductionIssues GET /production-issues
func (h *ProductionIssueHandler) ListProductionIssues(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if productionOrderID := c.Query("production_order_id"); productionOrderID != "" {
		query["production_order_id"] = toUint(productionOrderID)
	}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}
	if issueType := c.Query("issue_type"); issueType != "" {
		query["issue_type"] = issueType
	}
	if page := c.Query("page"); page != "" {
		query["page"], _ = strconv.Atoi(page)
	}
	if limit := c.Query("limit"); limit != "" {
		query["limit"], _ = strconv.Atoi(limit)
	}

	list, total, err := h.issueSvc.List(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// GetProductionIssue GET /production-issues/:id
func (h *ProductionIssueHandler) GetProductionIssue(c *gin.Context) {
	id := toUint(c.Param("id"))
	issue, err := h.issueSvc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, issue)
}

// CreateProductionIssue POST /production-issues
func (h *ProductionIssueHandler) CreateProductionIssue(c *gin.Context) {
	var req model.ProductionIssueCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	username := middleware.GetUsername(c)

	issue, err := h.issueSvc.Create(c.Request.Context(), tenantID, &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, issue)
}

// UpdateProductionIssue PUT /production-issues/:id
func (h *ProductionIssueHandler) UpdateProductionIssue(c *gin.Context) {
	id := toUint(c.Param("id"))
	var req model.ProductionIssueUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.issueSvc.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeleteProductionIssue DELETE /production-issues/:id
func (h *ProductionIssueHandler) DeleteProductionIssue(c *gin.Context) {
	id := toUint(c.Param("id"))
	if err := h.issueSvc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// SubmitProductionIssue POST /production-issues/:id/submit
func (h *ProductionIssueHandler) SubmitProductionIssue(c *gin.Context) {
	id := toUint(c.Param("id"))
	var req model.ProductionIssueSubmit
	c.ShouldBindJSON(&req)

	if err := h.issueSvc.Submit(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// StartPickProductionIssue POST /production-issues/:id/start-pick
func (h *ProductionIssueHandler) StartPickProductionIssue(c *gin.Context) {
	id := toUint(c.Param("id"))
	if err := h.issueSvc.StartPick(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ConfirmPickProductionIssue POST /production-issues/:id/confirm-pick
func (h *ProductionIssueHandler) ConfirmPickProductionIssue(c *gin.Context) {
	id := toUint(c.Param("id"))
	var req struct {
		Items []model.ProductionIssueItemSubmit `json:"items"`
	}
	c.ShouldBindJSON(&req)

	if err := h.issueSvc.ConfirmPick(c.Request.Context(), id, req.Items); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// IssueProductionIssue POST /production-issues/:id/issue
func (h *ProductionIssueHandler) IssueProductionIssue(c *gin.Context) {
	id := toUint(c.Param("id"))
	userID := middleware.GetUserID(c)

	if err := h.issueSvc.Issue(c.Request.Context(), id, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// CancelProductionIssue POST /production-issues/:id/cancel
func (h *ProductionIssueHandler) CancelProductionIssue(c *gin.Context) {
	id := toUint(c.Param("id"))
	if err := h.issueSvc.Cancel(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
