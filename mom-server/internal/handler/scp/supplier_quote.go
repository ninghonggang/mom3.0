package scp

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type SupplierQuoteHandler struct {
	scpService *service.SCPService
}

func NewSupplierQuoteHandler(s *service.SCPService) *SupplierQuoteHandler {
	return &SupplierQuoteHandler{scpService: s}
}

func (h *SupplierQuoteHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if rfqNo := c.Query("rfq_no"); rfqNo != "" {
		query["rfq_no"] = rfqNo
	}
	if supplierName := c.Query("supplier_name"); supplierName != "" {
		query["supplier_name"] = supplierName
	}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}

	list, total, err := h.scpService.ListSupplierQuotes(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *SupplierQuoteHandler) Get(c *gin.Context) {
	id := c.Param("id")
	quote, err := h.scpService.GetSupplierQuote(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, quote)
}

func (h *SupplierQuoteHandler) Create(c *gin.Context) {
	var req model.SupplierQuote
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.scpService.CreateQuote(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *SupplierQuoteHandler) GetQuotes(c *gin.Context) {
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

func (h *SupplierQuoteHandler) Award(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		SupplierID  int64   `json:"supplier_id"`
		TotalAmount float64 `json:"total_amount"`
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
