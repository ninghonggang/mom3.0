package eam

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
)

// EquipmentDowntimeHandler 设备停机Handler
type EquipmentDowntimeHandler struct {
	svc *service.EquipmentDowntimeService
}

// NewEquipmentDowntimeHandler 创建设备停机Handler
func NewEquipmentDowntimeHandler(svc *service.EquipmentDowntimeService) *EquipmentDowntimeHandler {
	return &EquipmentDowntimeHandler{svc: svc}
}

// List 获取设备停机列表
// @Summary 获取设备停机列表
// @Tags EAM-设备停机
// @Param equipment_id query uint64 false "设备ID"
// @Param equipment_code query string false "设备编号"
// @Param downtime_type query string false "停机类型"
// @Param status query string false "状态"
// @Param start_time query string false "开始时间"
// @Param end_time query string false "结束时间"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.Response{data=[]model.EquipmentDowntime}
// @Router /eam/downtime [get]
func (h *EquipmentDowntimeHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	query := make(map[string]interface{})
	if equipID := c.Query("equipment_id"); equipID != "" {
		if id, err := strconv.ParseInt(equipID, 10, 64); err == nil {
			query["equipment_id"] = id
		}
	}
	if equipCode := c.Query("equipment_code"); equipCode != "" {
		query["equipment_code"] = equipCode
	}
	if dtType := c.Query("downtime_type"); dtType != "" {
		query["downtime_type"] = dtType
	}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}
	if startTime := c.Query("start_time"); startTime != "" {
		query["start_time"] = startTime
	}
	if endTime := c.Query("end_time"); endTime != "" {
		query["end_time"] = endTime
	}
	if page, err := strconv.Atoi(c.DefaultQuery("page", "1")); err == nil && page > 0 {
		query["page"] = page
	}
	if pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "20")); err == nil && pageSize > 0 {
		query["page_size"] = pageSize
	}

	list, total, err := h.svc.List(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// Get 获取设备停机详情
// @Summary 获取设备停机详情
// @Tags EAM-设备停机
// @Param id path int true "ID"
// @Success 200 {object} response.Response{data=model.EquipmentDowntime}
// @Router /eam/downtime/{id} [get]
func (h *EquipmentDowntimeHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	downtime, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, downtime)
}

// Create 创建设备停机记录
// @Summary 创建设备停机记录
// @Tags EAM-设备停机
// @Param body body model.EquipmentDowntimeCreateRequest true "停机信息"
// @Success 200 {object} response.Response
// @Router /eam/downtime [post]
func (h *EquipmentDowntimeHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}
	username := middleware.GetUsername(c)

	var req model.EquipmentDowntimeCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	downtime, err := h.svc.Create(c.Request.Context(), tenantID, &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, downtime)
}

// Update 更新设备停机记录
// @Summary 更新设备停机记录
// @Tags EAM-设备停机
// @Param id path int true "ID"
// @Param body body model.EquipmentDowntimeUpdateRequest true "更新内容"
// @Success 200 {object} response.Response
// @Router /eam/downtime/{id} [put]
func (h *EquipmentDowntimeHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req model.EquipmentDowntimeUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	if err := h.svc.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Delete 删除设备停机记录
// @Summary 删除设备停机记录
// @Tags EAM-设备停机
// @Param id path int true "ID"
// @Success 200 {object} response.Response
// @Router /eam/downtime/{id} [delete]
func (h *EquipmentDowntimeHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// StartDowntime 开始停机
// @Summary 开始停机
// @Tags EAM-设备停机
// @Param id path int true "ID"
// @Success 200 {object} response.Response
// @Router /eam/downtime/{id}/start [post]
func (h *EquipmentDowntimeHandler) StartDowntime(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.svc.StartDowntime(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// EndDowntime 结束停机
// @Summary 结束停机
// @Tags EAM-设备停机
// @Param id path int true "ID"
// @Success 200 {object} response.Response
// @Router /eam/downtime/{id}/end [post]
func (h *EquipmentDowntimeHandler) EndDowntime(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.svc.EndDowntime(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
