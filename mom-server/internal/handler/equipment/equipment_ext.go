package equipment

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ========== TEEP分析 ==========

type TEEPDataHandler struct {
	service *service.TEEPDataService
}

func NewTEEPDataHandler(s *service.TEEPDataService) *TEEPDataHandler {
	return &TEEPDataHandler{service: s}
}

func (h *TEEPDataHandler) List(c *gin.Context) {
	equipmentID, _ := strconv.ParseInt(c.Query("equipment_id"), 10, 64)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	list, total, err := h.service.List(c.Request.Context(), equipmentID, startDate, endDate)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *TEEPDataHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	data, err := h.service.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, data)
}

func (h *TEEPDataHandler) Create(c *gin.Context) {
	var req model.TEEPData
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *TEEPDataHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *TEEPDataHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ========== 模具管理 ==========

type MoldHandler struct {
	service *service.MoldService
}

func NewMoldHandler(s *service.MoldService) *MoldHandler {
	return &MoldHandler{service: s}
}

func (h *MoldHandler) List(c *gin.Context) {
	query := c.Query("query")
	list, total, err := h.service.List(c.Request.Context(), query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *MoldHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	mold, err := h.service.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, mold)
}

func (h *MoldHandler) Create(c *gin.Context) {
	var req model.Mold
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *MoldHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *MoldHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

type MoldMaintenanceHandler struct {
	service *service.MoldMaintenanceService
}

func NewMoldMaintenanceHandler(s *service.MoldMaintenanceService) *MoldMaintenanceHandler {
	return &MoldMaintenanceHandler{service: s}
}

func (h *MoldMaintenanceHandler) List(c *gin.Context) {
	moldID, _ := strconv.ParseInt(c.Query("mold_id"), 10, 64)
	list, total, err := h.service.List(c.Request.Context(), moldID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *MoldMaintenanceHandler) Create(c *gin.Context) {
	var req model.MoldMaintenance
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

type MoldRepairHandler struct {
	service *service.MoldRepairService
}

func NewMoldRepairHandler(s *service.MoldRepairService) *MoldRepairHandler {
	return &MoldRepairHandler{service: s}
}

func (h *MoldRepairHandler) List(c *gin.Context) {
	moldID, _ := strconv.ParseInt(c.Query("mold_id"), 10, 64)
	list, total, err := h.service.List(c.Request.Context(), moldID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *MoldRepairHandler) Create(c *gin.Context) {
	var req model.MoldRepair
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// ========== 量检具管理 ==========

type GaugeHandler struct {
	service *service.GaugeService
}

func NewGaugeHandler(s *service.GaugeService) *GaugeHandler {
	return &GaugeHandler{service: s}
}

func (h *GaugeHandler) List(c *gin.Context) {
	query := c.Query("query")
	list, total, err := h.service.List(c.Request.Context(), query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *GaugeHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	gauge, err := h.service.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gauge)
}

func (h *GaugeHandler) Create(c *gin.Context) {
	var req model.Gauge
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *GaugeHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *GaugeHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

type GaugeCalibrationHandler struct {
	service *service.GaugeCalibrationService
}

func NewGaugeCalibrationHandler(s *service.GaugeCalibrationService) *GaugeCalibrationHandler {
	return &GaugeCalibrationHandler{service: s}
}

func (h *GaugeCalibrationHandler) List(c *gin.Context) {
	gaugeID, _ := strconv.ParseInt(c.Query("gauge_id"), 10, 64)
	list, total, err := h.service.List(c.Request.Context(), gaugeID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *GaugeCalibrationHandler) Create(c *gin.Context) {
	var req model.GaugeCalibration
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}
