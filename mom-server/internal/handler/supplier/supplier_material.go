package supplier

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type SupplierMaterialHandler struct {
	svc *service.SupplierMaterialService
}

func NewSupplierMaterialHandler(svc *service.SupplierMaterialService) *SupplierMaterialHandler {
	return &SupplierMaterialHandler{svc: svc}
}

func (h *SupplierMaterialHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	query := c.Query("query")
	list, total, err := h.svc.List(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *SupplierMaterialHandler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, item)
}

func (h *SupplierMaterialHandler) Create(c *gin.Context) {
	var req model.SupplierMaterial
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.svc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *SupplierMaterialHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.SupplierMaterial
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *SupplierMaterialHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *SupplierMaterialHandler) ListBySupplier(c *gin.Context) {
	supplierID := c.Param("supplier_id")
	list, err := h.svc.ListBySupplier(c.Request.Context(), supplierID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": len(list)})
}

func (h *SupplierMaterialHandler) ListByMaterial(c *gin.Context) {
	materialID := c.Param("material_id")
	list, err := h.svc.ListByMaterial(c.Request.Context(), materialID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": len(list)})
}

func (h *SupplierMaterialHandler) SetPreferred(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.SetPreferred(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
