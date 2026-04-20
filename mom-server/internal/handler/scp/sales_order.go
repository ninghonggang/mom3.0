package scp

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type SalesOrderHandler struct {
	scpService *service.SCPService
}

func NewSalesOrderHandler(s *service.SCPService) *SalesOrderHandler {
	return &SalesOrderHandler{scpService: s}
}

func (h *SalesOrderHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}
	if customerID := c.Query("customer_id"); customerID != "" {
		query["customer_id"] = customerID
	}

	list, total, err := h.scpService.ListSalesOrders(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *SalesOrderHandler) Get(c *gin.Context) {
	id := c.Param("id")
	order, err := h.scpService.GetSalesOrder(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, order)
}

func (h *SalesOrderHandler) Create(c *gin.Context) {
	var req model.SCPSalesOrder
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.scpService.CreateSalesOrder(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *SalesOrderHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.SCPSalesOrder
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.scpService.UpdateSalesOrder(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *SalesOrderHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.DeleteSalesOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *SalesOrderHandler) Submit(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.SubmitSalesOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *SalesOrderHandler) Approve(c *gin.Context) {
	id := c.Param("id")
	userID := middleware.GetUserID(c)
	if err := h.scpService.ApproveSalesOrder(c.Request.Context(), id, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *SalesOrderHandler) Reject(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.RejectSalesOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *SalesOrderHandler) Confirm(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.ConfirmSalesOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *SalesOrderHandler) Close(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.CloseSalesOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *SalesOrderHandler) Cancel(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.CancelSalesOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
