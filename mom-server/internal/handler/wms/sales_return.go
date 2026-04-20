package wms

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type SalesReturnHandler struct {
	returnSvc *service.SalesReturnService
}

func NewSalesReturnHandler(returnSvc *service.SalesReturnService) *SalesReturnHandler {
	return &SalesReturnHandler{returnSvc: returnSvc}
}

// ListSalesReturns GET /sales-return
func (h *SalesReturnHandler) ListSalesReturns(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if salesOrderID := c.Query("sales_order_id"); salesOrderID != "" {
		query["sales_order_id"], _ = strconv.ParseInt(salesOrderID, 10, 64)
	}
	if customerID := c.Query("customer_id"); customerID != "" {
		query["customer_id"], _ = strconv.ParseInt(customerID, 10, 64)
	}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}
	if returnType := c.Query("return_type"); returnType != "" {
		query["return_type"] = returnType
	}
	if page := c.Query("page"); page != "" {
		query["page"], _ = strconv.Atoi(page)
	}
	if limit := c.Query("limit"); limit != "" {
		query["limit"], _ = strconv.Atoi(limit)
	}

	list, total, err := h.returnSvc.List(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// GetSalesReturn GET /sales-return/:id
func (h *SalesReturnHandler) GetSalesReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	ret, err := h.returnSvc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, ret)
}

// CreateSalesReturn POST /sales-return
func (h *SalesReturnHandler) CreateSalesReturn(c *gin.Context) {
	var req model.SalesReturnCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	username := middleware.GetUsername(c)

	ret, err := h.returnSvc.Create(c.Request.Context(), tenantID, &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, ret)
}

// UpdateSalesReturn PUT /sales-return/:id
func (h *SalesReturnHandler) UpdateSalesReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	var req model.SalesReturnUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.returnSvc.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeleteSalesReturn DELETE /sales-return/:id
func (h *SalesReturnHandler) DeleteSalesReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	if err := h.returnSvc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// SubmitSalesReturn POST /sales-return/:id/submit
func (h *SalesReturnHandler) SubmitSalesReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	var req model.SalesReturnSubmit
	c.ShouldBindJSON(&req)

	if err := h.returnSvc.Submit(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ApproveSalesReturn POST /sales-return/:id/approve
func (h *SalesReturnHandler) ApproveSalesReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	userID := middleware.GetUserID(c)

	if err := h.returnSvc.Approve(c.Request.Context(), id, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// StartReturnSalesReturn POST /sales-return/:id/start-return
func (h *SalesReturnHandler) StartReturnSalesReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	if err := h.returnSvc.StartReturn(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ConfirmSalesReturn POST /sales-return/:id/confirm
func (h *SalesReturnHandler) ConfirmSalesReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	var req model.SalesReturnConfirm
	c.ShouldBindJSON(&req)

	if err := h.returnSvc.ConfirmReturn(c.Request.Context(), id, req.Items); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// CancelSalesReturn POST /sales-return/:id/cancel
func (h *SalesReturnHandler) CancelSalesReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	if err := h.returnSvc.Cancel(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}