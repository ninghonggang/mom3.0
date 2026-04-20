package wms

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type PurchaseReturnHandler struct {
	returnSvc *service.PurchaseReturnService
}

func NewPurchaseReturnHandler(returnSvc *service.PurchaseReturnService) *PurchaseReturnHandler {
	return &PurchaseReturnHandler{returnSvc: returnSvc}
}

func toUint64(s string) uint {
	i, _ := strconv.ParseUint(s, 10, 64)
	return uint(i)
}

// ListPurchaseReturns GET /purchase-return
func (h *PurchaseReturnHandler) ListPurchaseReturns(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if purchaseOrderID := c.Query("purchase_order_id"); purchaseOrderID != "" {
		query["purchase_order_id"], _ = strconv.ParseInt(purchaseOrderID, 10, 64)
	}
	if supplierID := c.Query("supplier_id"); supplierID != "" {
		query["supplier_id"], _ = strconv.ParseInt(supplierID, 10, 64)
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

// GetPurchaseReturn GET /purchase-return/:id
func (h *PurchaseReturnHandler) GetPurchaseReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	ret, err := h.returnSvc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, ret)
}

// CreatePurchaseReturn POST /purchase-return
func (h *PurchaseReturnHandler) CreatePurchaseReturn(c *gin.Context) {
	var req model.PurchaseReturnCreate
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

// UpdatePurchaseReturn PUT /purchase-return/:id
func (h *PurchaseReturnHandler) UpdatePurchaseReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	var req model.PurchaseReturnUpdate
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

// DeletePurchaseReturn DELETE /purchase-return/:id
func (h *PurchaseReturnHandler) DeletePurchaseReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	if err := h.returnSvc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// SubmitPurchaseReturn POST /purchase-return/:id/submit
func (h *PurchaseReturnHandler) SubmitPurchaseReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	var req model.PurchaseReturnSubmit
	c.ShouldBindJSON(&req)

	if err := h.returnSvc.Submit(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ApprovePurchaseReturn POST /purchase-return/:id/approve
func (h *PurchaseReturnHandler) ApprovePurchaseReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	userID := middleware.GetUserID(c)

	if err := h.returnSvc.Approve(c.Request.Context(), id, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// StartReturnPurchaseReturn POST /purchase-return/:id/start-return
func (h *PurchaseReturnHandler) StartReturnPurchaseReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	if err := h.returnSvc.StartReturn(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ConfirmPurchaseReturn POST /purchase-return/:id/confirm
func (h *PurchaseReturnHandler) ConfirmPurchaseReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	var req model.PurchaseReturnConfirm
	c.ShouldBindJSON(&req)

	if err := h.returnSvc.ConfirmReturn(c.Request.Context(), id, req.Items); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// CancelPurchaseReturn POST /purchase-return/:id/cancel
func (h *PurchaseReturnHandler) CancelPurchaseReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	if err := h.returnSvc.Cancel(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}