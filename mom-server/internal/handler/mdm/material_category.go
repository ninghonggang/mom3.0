package mdm

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type MaterialCategoryHandler struct {
	materialCategoryService *service.MaterialCategoryService
}

func NewMaterialCategoryHandler(mcs *service.MaterialCategoryService) *MaterialCategoryHandler {
	return &MaterialCategoryHandler{materialCategoryService: mcs}
}

func (h *MaterialCategoryHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	categories, err := h.materialCategoryService.List(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, categories)
}

func (h *MaterialCategoryHandler) Tree(c *gin.Context) {
	var tenantID int64
	if middleware.IsSuperAdmin(c) {
		tenantID = 0
	} else {
		tenantID = middleware.GetTenantID(c)
	}
	categories, err := h.materialCategoryService.Tree(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, categories)
}

func (h *MaterialCategoryHandler) Get(c *gin.Context) {
	id := c.Param("id")
	category, err := h.materialCategoryService.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, category)
}

func (h *MaterialCategoryHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}
	var req model.MaterialCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req.TenantID = tenantID
	err := h.materialCategoryService.Create(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, req)
}

func (h *MaterialCategoryHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.MaterialCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.materialCategoryService.Update(c.Request.Context(), id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, req)
}

func (h *MaterialCategoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.materialCategoryService.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
