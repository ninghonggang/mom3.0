package scp

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type CustomerInquiryHandler struct {
	scpService *service.SCPService
}

func NewCustomerInquiryHandler(s *service.SCPService) *CustomerInquiryHandler {
	return &CustomerInquiryHandler{scpService: s}
}

func (h *CustomerInquiryHandler) List(c *gin.Context) {
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
	if inquiryNo := c.Query("inquiry_no"); inquiryNo != "" {
		query["inquiry_no"] = inquiryNo
	}

	list, total, err := h.scpService.ListCustomerInquiries(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *CustomerInquiryHandler) Get(c *gin.Context) {
	id := c.Param("id")
	inquiry, err := h.scpService.GetCustomerInquiry(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, inquiry)
}

func (h *CustomerInquiryHandler) Create(c *gin.Context) {
	var req model.CustomerInquiry
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.scpService.CreateCustomerInquiry(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *CustomerInquiryHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.CustomerInquiry
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.scpService.UpdateCustomerInquiry(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *CustomerInquiryHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.DeleteCustomerInquiry(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *CustomerInquiryHandler) Send(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.SendCustomerInquiry(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *CustomerInquiryHandler) Quote(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		QuotedAmount float64 `json:"quoted_amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.scpService.QuoteCustomerInquiry(c.Request.Context(), id, req.QuotedAmount); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *CustomerInquiryHandler) Win(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		SupplierID int64 `json:"supplier_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.scpService.WinCustomerInquiry(c.Request.Context(), id, req.SupplierID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *CustomerInquiryHandler) Lose(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.LoseCustomerInquiry(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *CustomerInquiryHandler) Cancel(c *gin.Context) {
	id := c.Param("id")
	if err := h.scpService.CancelCustomerInquiry(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
