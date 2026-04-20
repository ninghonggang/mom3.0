package container

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ContainerLifecycleHandler 容器生命周期处理器
type ContainerLifecycleHandler struct {
	service *service.ContainerLifecycleService
}

func NewContainerLifecycleHandler(s *service.ContainerLifecycleService) *ContainerLifecycleHandler {
	return &ContainerLifecycleHandler{service: s}
}

// ListContainerLifecycles 获取生命周期记录列表
func (h *ContainerLifecycleHandler) ListContainerLifecycles(c *gin.Context) {
	var query model.ContainerLifecycleQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		query.Page = 1
		query.PageSize = 20
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	list, total, err := h.service.ListContainerLifecycles(c.Request.Context(), tenantID, &query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

// GetContainerLifecycle 获取单条生命周期记录
func (h *ContainerLifecycleHandler) GetContainerLifecycle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	lifecycle, err := h.service.GetContainerLifecycle(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, lifecycle)
}

// InitializeContainer 初始化容器
func (h *ContainerLifecycleHandler) InitializeContainer(c *gin.Context) {
	containerIdStr := c.Param("containerId")
	containerId, err := strconv.ParseUint(containerIdStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的容器ID")
		return
	}

	var req model.InitializeContainerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.ContainerID = int64(containerId)

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	if err := h.service.InitializeContainer(c.Request.Context(), tenantID, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// RecordMaintenance 记录维修
func (h *ContainerLifecycleHandler) RecordMaintenance(c *gin.Context) {
	containerIdStr := c.Param("containerId")
	containerId, err := strconv.ParseUint(containerIdStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的容器ID")
		return
	}

	var req model.RecordMaintenanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.ContainerID = int64(containerId)

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	if err := h.service.RecordMaintenance(c.Request.Context(), tenantID, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// CompleteMaintenance 完成维修
func (h *ContainerLifecycleHandler) CompleteMaintenance(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	var req model.CompleteMaintenanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	if err := h.service.CompleteMaintenance(c.Request.Context(), tenantID, uint(id), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// RetireContainer 报废容器
func (h *ContainerLifecycleHandler) RetireContainer(c *gin.Context) {
	containerIdStr := c.Param("containerId")
	containerId, err := strconv.ParseUint(containerIdStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的容器ID")
		return
	}

	var req model.RetireContainerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.ContainerID = int64(containerId)

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	if err := h.service.RetireContainer(c.Request.Context(), tenantID, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetContainerTimeline 获取容器时间线
func (h *ContainerLifecycleHandler) GetContainerTimeline(c *gin.Context) {
	containerIdStr := c.Param("containerId")
	containerId, err := strconv.ParseUint(containerIdStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的容器ID")
		return
	}

	timeline, err := h.service.GetContainerTimeline(c.Request.Context(), int64(containerId))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list": timeline,
	})
}

// ListContainerMaintenances 获取维修记录列表
func (h *ContainerLifecycleHandler) ListContainerMaintenances(c *gin.Context) {
	var query model.ContainerMaintenanceQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		query.Page = 1
		query.PageSize = 20
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	list, total, err := h.service.ListContainerMaintenances(c.Request.Context(), tenantID, &query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}