package mdm

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type MaterialHandler struct {
	materialService *service.MaterialService
}

func NewMaterialHandler(ms *service.MaterialService) *MaterialHandler {
	return &MaterialHandler{materialService: ms}
}

func (h *MaterialHandler) List(c *gin.Context) {
	list, total, err := h.materialService.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *MaterialHandler) Get(c *gin.Context) {
	id := c.Param("id")
	material, err := h.materialService.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, material)
}

func (h *MaterialHandler) Create(c *gin.Context) {
	var req model.Material
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	// 设置租户ID
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID

	if err := h.materialService.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *MaterialHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.Material
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.materialService.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *MaterialHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.materialService.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

type WorkshopHandler struct {
	workshopService *service.WorkshopService
}

func NewWorkshopHandler(ws *service.WorkshopService) *WorkshopHandler {
	return &WorkshopHandler{workshopService: ws}
}

func (h *WorkshopHandler) List(c *gin.Context) {
	list, total, err := h.workshopService.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *WorkshopHandler) Get(c *gin.Context) {
	id := c.Param("id")
	workshop, err := h.workshopService.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, workshop)
}

func (h *WorkshopHandler) Create(c *gin.Context) {
	var req model.Workshop
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	// 设置租户ID
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID

	if err := h.workshopService.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *WorkshopHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.Workshop
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.workshopService.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *WorkshopHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.workshopService.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
