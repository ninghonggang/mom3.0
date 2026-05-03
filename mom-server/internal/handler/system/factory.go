package system

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type FactoryHandler struct {
	svc *service.FactoryService
}

func NewFactoryHandler(svc *service.FactoryService) *FactoryHandler {
	return &FactoryHandler{svc: svc}
}

// List 获取工厂列表
// GET /system/factory/list
func (h *FactoryHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	factoryCode := c.Query("factory_code")
	factoryName := c.Query("factory_name")
	status, _ := strconv.Atoi(c.Query("status"))

	req := &model.FactoryQuery{
		FactoryCode: factoryCode,
		FactoryName: factoryName,
		Status:      status,
		Page:        page,
		PageSize:    pageSize,
	}

	list, total, err := h.svc.List(c.Request.Context(), tenantID, req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// Get 获取单个工厂
// GET /system/factory/:id
func (h *FactoryHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	factory, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, factory)
}

// Create 创建工厂
// POST /system/factory
func (h *FactoryHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	var req model.FactoryCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	factory, err := h.svc.Create(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, factory)
}

// Update 更新工厂
// PUT /system/factory/:id
func (h *FactoryHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req model.FactoryCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Delete 删除工厂
// DELETE /system/factory/:id
func (h *FactoryHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// SetDefault 设置默认工厂
// PUT /system/factory/default/:id
func (h *FactoryHandler) SetDefault(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.svc.SetDefault(c.Request.Context(), tenantID, id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetCurrentFactory 获取当前工厂
// GET /system/factory/current
func (h *FactoryHandler) GetCurrentFactory(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	userID := middleware.GetUserID(c)
	factory, err := h.svc.GetCurrentFactory(c.Request.Context(), tenantID, userID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, factory)
}

// SetCurrentFactory 设置当前工厂
// PUT /system/factory/current
func (h *FactoryHandler) SetCurrentFactory(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	userID := middleware.GetUserID(c)
	var req struct {
		FactoryID int64 `json:"factory_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.SetCurrentFactory(c.Request.Context(), tenantID, userID, req.FactoryID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}