package production

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductionReturnHandler struct {
	returnSvc *service.ProductionReturnService
}

func NewProductionReturnHandler(returnSvc *service.ProductionReturnService) *ProductionReturnHandler {
	return &ProductionReturnHandler{returnSvc: returnSvc}
}

func toUint64(s string) uint {
	i, _ := strconv.ParseUint(s, 10, 64)
	return uint(i)
}

// ListProductionReturns GET /production-return
func (h *ProductionReturnHandler) ListProductionReturns(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if productionOrderID := c.Query("production_order_id"); productionOrderID != "" {
		query["production_order_id"], _ = strconv.ParseInt(productionOrderID, 10, 64)
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

// GetProductionReturn GET /production-return/:id
func (h *ProductionReturnHandler) GetProductionReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	ret, err := h.returnSvc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, ret)
}

// CreateProductionReturn POST /production-return
func (h *ProductionReturnHandler) CreateProductionReturn(c *gin.Context) {
	var req model.ProductionReturnCreate
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

// UpdateProductionReturn PUT /production-return/:id
func (h *ProductionReturnHandler) UpdateProductionReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	var req model.ProductionReturnUpdate
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

// DeleteProductionReturn DELETE /production-return/:id
func (h *ProductionReturnHandler) DeleteProductionReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	if err := h.returnSvc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// SubmitProductionReturn POST /production-return/:id/submit
func (h *ProductionReturnHandler) SubmitProductionReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	var req model.ProductionReturnSubmit
	c.ShouldBindJSON(&req)

	if err := h.returnSvc.Submit(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ApproveProductionReturn POST /production-return/:id/approve
func (h *ProductionReturnHandler) ApproveProductionReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	userID := middleware.GetUserID(c)

	if err := h.returnSvc.Approve(c.Request.Context(), id, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// StartReturnProductionReturn POST /production-return/:id/start-return
func (h *ProductionReturnHandler) StartReturnProductionReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	if err := h.returnSvc.StartReturn(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ConfirmReturnProductionReturn POST /production-return/:id/confirm-return
func (h *ProductionReturnHandler) ConfirmReturnProductionReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	var req model.ProductionReturnConfirm
	c.ShouldBindJSON(&req)

	if err := h.returnSvc.ConfirmReturn(c.Request.Context(), id, req.Items); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// CancelProductionReturn POST /production-return/:id/cancel
func (h *ProductionReturnHandler) CancelProductionReturn(c *gin.Context) {
	id := toUint64(c.Param("id"))
	if err := h.returnSvc.Cancel(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
