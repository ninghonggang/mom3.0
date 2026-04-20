package agv

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

// AGVHandler AGV接口处理
type AGVHandler struct {
	svc *service.AGVService
}

// NewAGVHandler 创建AGV处理器
func NewAGVHandler(svc *service.AGVService) *AGVHandler {
	return &AGVHandler{svc: svc}
}

// ListTasks GET /agv/task/list - 查询任务列表
func (h *AGVHandler) ListTasks(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	q := &model.AGVTaskQuery{
		TenantID:  tenantID,
		TaskNo:    c.Query("taskNo"),
		TaskType:  model.AGVTaskType(c.Query("taskType")),
		Status:    model.AGVTaskStatus(c.Query("status")),
		AGVCode:   c.Query("agvCode"),
		StartDate: c.Query("startDate"),
		EndDate:   c.Query("endDate"),
		Page:      page,
		PageSize:  pageSize,
	}

	list, total, err := h.svc.ListTasks(c.Request.Context(), q)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// GetTask GET /agv/task/:id - 获取任务详情
func (h *AGVHandler) GetTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	task, err := h.svc.GetTaskStatus(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, task)
}

// CreateTask POST /agv/task - 创建任务
func (h *AGVHandler) CreateTask(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var req model.CreateAGVTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.TenantID = tenantID

	task, err := h.svc.CreateDeliveryTask(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, task)
}

// CancelTask PUT /agv/task/:id/cancel - 取消任务
func (h *AGVHandler) CancelTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.svc.CancelTask(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "任务已取消"})
}

// AssignTask PUT /agv/task/:id/assign - 分配任务给AGV
func (h *AGVHandler) AssignTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req struct {
		AGVCode string `json:"agvCode" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.AssignTaskToAGV(c.Request.Context(), id, req.AGVCode); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "任务已分配"})
}

// CompleteTask PUT /agv/task/:id/complete - 完成任务
func (h *AGVHandler) CompleteTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.svc.CompleteTask(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "任务已完成"})
}

// StartTask PUT /agv/task/:id/start - 开始任务
func (h *AGVHandler) StartTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.svc.StartTask(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "任务已开始"})
}

// ListDevices GET /agv/device/list - 查询设备列表
func (h *AGVHandler) ListDevices(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	q := &model.AGVDeviceQuery{
		TenantID: tenantID,
		AGVCode: c.Query("agvCode"),
		Status:  model.AGVDeviceStatus(c.Query("status")),
		Page:    page,
		PageSize: pageSize,
	}

	list, total, err := h.svc.ListDevices(c.Request.Context(), q)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// GetDevice GET /agv/device/:id - 获取设备详情
func (h *AGVHandler) GetDevice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	device, err := h.svc.GetDevice(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, device)
}

// CreateDevice POST /agv/device - 创建设备
func (h *AGVHandler) CreateDevice(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var req model.CreateAGVDeviceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.TenantID = tenantID

	device, err := h.svc.CreateDevice(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, device)
}

// UpdateDeviceStatus PUT /agv/device/:id/status - 更新设备状态
func (h *AGVHandler) UpdateDeviceStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.UpdateAGVDeviceStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	device, err := h.svc.GetDevice(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	if err := h.svc.UpdateDeviceStatus(c.Request.Context(), device.AGVCode, req.Status, req.BatteryLevel, req.CurrentLocation); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "状态已更新"})
}

// ListLocations GET /agv/location/list - 查询库位映射列表
func (h *AGVHandler) ListLocations(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	enabled := c.Query("enabled")
	var enabledPtr *bool
	if enabled != "" {
		b := enabled == "true"
		enabledPtr = &b
	}

	q := &model.AGVLocationQuery{
		TenantID:     tenantID,
		LocationCode: c.Query("locationCode"),
		LocationType: model.AGVLocationType(c.Query("locationType")),
		Enabled:      enabledPtr,
		Page:         page,
		PageSize:     pageSize,
	}

	list, total, err := h.svc.ListLocations(c.Request.Context(), q)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// GetLocation GET /agv/location/:id - 获取库位映射详情
func (h *AGVHandler) GetLocation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	location, err := h.svc.GetLocation(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, location)
}

// CreateLocation POST /agv/location - 创建库位映射
func (h *AGVHandler) CreateLocation(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var req model.CreateAGVLocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.TenantID = tenantID

	location, err := h.svc.CreateLocation(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, location)
}

// UpdateLocation PUT /agv/location/:id - 更新库位映射
func (h *AGVHandler) UpdateLocation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.UpdateLocation(c.Request.Context(), id, updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "更新成功"})
}

// DeleteLocation DELETE /agv/location/:id - 删除库位映射
func (h *AGVHandler) DeleteLocation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.svc.DeleteLocation(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

// Heartbeat POST /agv/callback/heartbeat - AGV心跳回调
func (h *AGVHandler) Heartbeat(c *gin.Context) {
	var req struct {
		AGVCode      string  `json:"agv_code" binding:"required"`
		BatteryLevel float64 `json:"battery_level"`
		Position     string  `json:"position"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.RegisterAGVHeartbeat(c.Request.Context(), req.AGVCode, req.BatteryLevel, req.Position); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "心跳已记录"})
}

// TaskCallback POST /agv/callback/task - AGV任务状态回调
func (h *AGVHandler) TaskCallback(c *gin.Context) {
	var req service.AGVCallbackResult
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.HandleAGVCallback(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "回调已处理"})
}

// GetAvailableAGVs GET /agv/device/available - 获取可用AGV
func (h *AGVHandler) GetAvailableAGVs(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	devices, err := h.svc.GetAvailableAGVs(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": devices})
}
