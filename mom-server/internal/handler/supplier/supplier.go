package supplier

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type SupplierHandler struct {
	supplierService *service.SupplierService
}

func NewSupplierHandler(s *service.SupplierService) *SupplierHandler {
	return &SupplierHandler{supplierService: s}
}

func (h *SupplierHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	list, total, err := h.supplierService.List(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *SupplierHandler) Get(c *gin.Context) {
	id := c.Param("id")
	supplier, err := h.supplierService.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, supplier)
}

func (h *SupplierHandler) Create(c *gin.Context) {
	var req model.Supplier
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
	if err := h.supplierService.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *SupplierHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.Supplier
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.supplierService.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *SupplierHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.supplierService.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
