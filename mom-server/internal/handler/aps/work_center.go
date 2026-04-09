package aps

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type WorkCenterHandler struct {
	service *service.WorkCenterService
}

func NewWorkCenterHandler(s *service.WorkCenterService) *WorkCenterHandler {
	return &WorkCenterHandler{service: s}
}

func (h *WorkCenterHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	list, total, err := h.service.List(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *WorkCenterHandler) Get(c *gin.Context) {
	id := c.Param("id")
	wc, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, wc)
}

func (h *WorkCenterHandler) Create(c *gin.Context) {
	var req model.WorkCenter
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *WorkCenterHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.WorkCenter
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := h.service.Update(c.Request.Context(), id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *WorkCenterHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *WorkCenterHandler) ListByWorkshop(c *gin.Context) {
	workshopID := c.Query("workshop_id")
	if workshopID == "" {
		response.BadRequest(c, "workshop_id is required")
		return
	}
	list, err := h.service.ListByWorkshop(c.Request.Context(), 0)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, list)
}
