package scp

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type PurchaseOrderHandler struct {
	scpService *service.SCPService
}

func NewPurchaseOrderHandler(s *service.SCPService) *PurchaseOrderHandler {
	return &PurchaseOrderHandler{scpService: s}
}

func (h *PurchaseOrderHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}
	if supplierID := c.Query("supplier_id"); supplierID != "" {
		query["supplier_id"] = supplierID
	}
	if startDate := c.Query("start_date"); startDate != "" {
		query["start_date"] = startDate
	}
	if endDate := c.Query("end_date"); endDate != "" {
		query["end_date"] = endDate
	}
	if page := c.DefaultQuery("page", "1"); page != "" {
		query["page"] = 1
	}

	list, total, err := h.scpService.ListPurchaseOrders(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *PurchaseOrderHandler) Get(c *gin.Context) {
	id := c.Param("id")
	order, err := h.scpService.GetPurchaseOrder(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, order)
}

func (h *PurchaseOrderHandler) Create(c *gin.Context) {
	var req model.PurchaseOrder
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.scpService.CreatePurchaseOrder(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *PurchaseOrderHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.PurchaseOrder
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.scpService.UpdatePurchaseOrder(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *PurchaseOrderHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.DeletePurchaseOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *PurchaseOrderHandler) Submit(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.SubmitPurchaseOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *PurchaseOrderHandler) Approve(c *gin.Context) {
	id := c.Param("id")
	userID := middleware.GetUserID(c)
	if err := h.scpService.ApprovePurchaseOrder(c.Request.Context(), id, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *PurchaseOrderHandler) Reject(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.RejectPurchaseOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *PurchaseOrderHandler) Issue(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.IssuePurchaseOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *PurchaseOrderHandler) Close(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		CloseReason string `json:"close_reason" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "请提供关闭原因")
		return
	}
	if err := h.scpService.ClosePurchaseOrder(c.Request.Context(), id, req.CloseReason); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *PurchaseOrderHandler) Cancel(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.CancelPurchaseOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *PurchaseOrderHandler) Receive(c *gin.Context) {
	var req struct {
		ItemID      uint    `json:"item_id" binding:"required"`
		ReceivedQty float64 `json:"received_qty" binding:"required"`
		BatchNo     string  `json:"batch_no"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "请提供行项目ID和收货数量")
		return
	}
	if err := h.scpService.ReceivePurchaseOrderItem(c.Request.Context(), req.ItemID, req.ReceivedQty, req.BatchNo); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
