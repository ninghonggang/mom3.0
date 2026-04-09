package mdm

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerService *service.CustomerService
}

func NewCustomerHandler(s *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: s}
}

func (h *CustomerHandler) List(c *gin.Context) {
	list, total, err := h.customerService.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *CustomerHandler) Get(c *gin.Context) {
	id := c.Param("id")
	customer, err := h.customerService.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, customer)
}

func (h *CustomerHandler) Create(c *gin.Context) {
	var req model.Customer
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	// 设置默认租户ID
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.customerService.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *CustomerHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.Customer
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.customerService.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *CustomerHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.customerService.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
