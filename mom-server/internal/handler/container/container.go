package container

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type ContainerHandler struct {
	service *service.ContainerService
}

func NewContainerHandler(s *service.ContainerService) *ContainerHandler {
	return &ContainerHandler{service: s}
}

// List 器具列表
func (h *ContainerHandler) List(c *gin.Context) {
	var params model.ContainerQueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		params.Page = 1
		params.PageSize = 20
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	list, total, err := h.service.List(c.Request.Context(), tenantID, params)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

// Get 器具详情
func (h *ContainerHandler) Get(c *gin.Context) {
	id := c.Param("id")
	container, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, container)
}

// Create 创建器具
func (h *ContainerHandler) Create(c *gin.Context) {
	var req model.ContainerMaster
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

// Update 更新器具
func (h *ContainerHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.ContainerMaster
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

// Delete 删除器具
func (h *ContainerHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// In 入库
func (h *ContainerHandler) In(c *gin.Context) {
	id := c.Param("id")
	var req model.ContainerMovement
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID

	err := h.service.In(c.Request.Context(), id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Out 出库
func (h *ContainerHandler) Out(c *gin.Context) {
	id := c.Param("id")
	var req model.ContainerMovement
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID

	err := h.service.Out(c.Request.Context(), id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Return 退回
func (h *ContainerHandler) Return(c *gin.Context) {
	id := c.Param("id")
	var req model.ContainerMovement
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID

	err := h.service.Return(c.Request.Context(), id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Transfer 调拨
func (h *ContainerHandler) Transfer(c *gin.Context) {
	id := c.Param("id")
	var req model.ContainerMovement
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID

	err := h.service.Transfer(c.Request.Context(), id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Clean 清洁
func (h *ContainerHandler) Clean(c *gin.Context) {
	id := c.Param("id")
	var req model.ContainerMovement
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID

	err := h.service.Clean(c.Request.Context(), id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Movements 流转记录
func (h *ContainerHandler) Movements(c *gin.Context) {
	id := c.Param("id")
	page := 1
	pageSize := 20

	list, total, err := h.service.GetMovements(c.Request.Context(), id, page, pageSize)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}
