package mdm

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"mom-server/internal/model"
	"mom-server/internal/service"
	"mom-server/internal/pkg/response"
)

// WorkshopConfigHandler 车间配置Handler
type WorkshopConfigHandler struct {
	svc *service.WorkshopConfigService
}

// NewWorkshopConfigHandler 创建车间配置Handler
func NewWorkshopConfigHandler(svc *service.WorkshopConfigService) *WorkshopConfigHandler {
	return &WorkshopConfigHandler{svc: svc}
}

// GetConfig 获取车间配置
// @Summary 获取车间配置
// @Tags MDC-车间配置
// @Param workshop_id path int true "车间ID"
// @Success 200 {object} response.Response{data=model.WorkshopConfig}
// @Router /mdm/workshop-config/{workshop_id} [get]
func (h *WorkshopConfigHandler) GetConfig(c *gin.Context) {
	workshopID, err := strconv.ParseInt(c.Param("workshop_id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的车间ID")
		return
	}

	config, err := h.svc.GetByWorkshopID(c.Request.Context(), workshopID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "获取车间配置失败")
		return
	}

	response.Success(c, config)
}

// UpdateConfig 更新车间配置
// @Summary 更新车间配置
// @Tags MDC-车间配置
// @Param workshop_id path int true "车间ID"
// @Param body body model.WorkshopConfig true "车间配置"
// @Success 200 {object} response.Response
// @Router /mdm/workshop-config/{workshop_id} [put]
func (h *WorkshopConfigHandler) UpdateConfig(c *gin.Context) {
	workshopID, err := strconv.ParseInt(c.Param("workshop_id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的车间ID")
		return
	}

	var req model.WorkshopConfig
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	updates := map[string]interface{}{
		"erp_plant_code":      req.ErpPlantCode,
		"max_devices":         req.MaxDevices,
		"max_workers":         req.MaxWorkers,
		"max_capacity_per_day": req.MaxCapacityPerDay,
		"time_zone":           req.TimeZone,
		"is_default":          req.IsDefault,
	}

	if err := h.svc.Update(c.Request.Context(), workshopID, updates); err != nil {
		response.Error(c, http.StatusInternalServerError, "更新车间配置失败")
		return
	}

	response.Success(c, nil)
}

// WorkingCalendarHandler 工厂日历Handler
type WorkingCalendarHandler struct {
	svc *service.WorkingCalendarService
}

// NewWorkingCalendarHandler 创建工厂日历Handler
func NewWorkingCalendarHandler(svc *service.WorkingCalendarService) *WorkingCalendarHandler {
	return &WorkingCalendarHandler{svc: svc}
}

// GetCalendars 获取日历列表
// @Summary 获取日历列表
// @Tags APS-工厂日历
// @Param workshop_id query int false "车间ID"
// @Success 200 {object} response.Response{data=[]model.WorkingCalendar}
// @Router /aps/calendar [get]
func (h *WorkingCalendarHandler) GetCalendars(c *gin.Context) {
	workshopIDStr := c.Query("workshop_id")
	if workshopIDStr == "" {
		response.Error(c, http.StatusBadRequest, "请提供车间ID")
		return
	}

	workshopID, err := strconv.ParseInt(workshopIDStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的车间ID")
		return
	}

	calendars, err := h.svc.GetByWorkshopID(c.Request.Context(), workshopID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "获取日历失败")
		return
	}

	response.Success(c, calendars)
}

// CreateCalendar 创建日历
// @Summary 创建日历
// @Tags APS-工厂日历
// @Param body body model.WorkingCalendar true "日历"
// @Success 200 {object} response.Response
// @Router /aps/calendar [post]
func (h *WorkingCalendarHandler) CreateCalendar(c *gin.Context) {
	var req model.WorkingCalendar
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	if err := h.svc.Create(c.Request.Context(), &req); err != nil {
		response.Error(c, http.StatusInternalServerError, "创建日历失败")
		return
	}

	response.Success(c, nil)
}

// UpdateCalendar 更新日历
// @Summary 更新日历
// @Tags APS-工厂日历
// @Param id path int true "日历ID"
// @Param body body model.WorkingCalendar true "日历"
// @Success 200 {object} response.Response
// @Router /aps/calendar/{id} [put]
func (h *WorkingCalendarHandler) UpdateCalendar(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的日历ID")
		return
	}

	var req model.WorkingCalendar
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	updates := map[string]interface{}{
		"calendar_name":     req.CalendarName,
		"work_days":         req.WorkDays,
		"shifts":            req.Shifts,
		"holiday_dates":     req.HolidayDates,
		"special_work_dates": req.SpecialWorkDates,
		"effective_from":    req.EffectiveFrom,
		"effective_to":      req.EffectiveTo,
		"status":            req.Status,
	}

	if err := h.svc.Update(c.Request.Context(), id, updates); err != nil {
		response.Error(c, http.StatusInternalServerError, "更新日历失败")
		return
	}

	response.Success(c, nil)
}

// DeleteCalendar 删除日历
// @Summary 删除日历
// @Tags APS-工厂日历
// @Param id path int true "日历ID"
// @Success 200 {object} response.Response
// @Router /aps/calendar/{id} [delete]
func (h *WorkingCalendarHandler) DeleteCalendar(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的日历ID")
		return
	}

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, "删除日历失败")
		return
	}

	response.Success(c, nil)
}
