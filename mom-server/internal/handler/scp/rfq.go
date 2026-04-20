package scp

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type RFQHandler struct {
	scpService *service.SCPService
}

func NewRFQHandler(s *service.SCPService) *RFQHandler {
	return &RFQHandler{scpService: s}
}

func (h *RFQHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}

	list, total, err := h.scpService.ListRFQs(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *RFQHandler) Get(c *gin.Context) {
	id := c.Param("id")
	rfq, err := h.scpService.GetRFQ(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, rfq)
}

func (h *RFQHandler) Create(c *gin.Context) {
	var req model.RFQ
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.scpService.CreateRFQ(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *RFQHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.RFQ
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.scpService.UpdateRFQ(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *RFQHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.DeleteRFQ(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *RFQHandler) Publish(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.PublishRFQ(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *RFQHandler) Close(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.CloseRFQ(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *RFQHandler) GetQuotes(c *gin.Context) {
	id := c.Param("id")
	quotes, err := h.scpService.ListQuotesByRFQ(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list": quotes,
	})
}

func (h *RFQHandler) Award(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		SupplierID   int64   `json:"supplier_id"`
		TotalAmount  float64 `json:"total_amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.scpService.AwardRFQ(c.Request.Context(), id, req.SupplierID, req.TotalAmount); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
