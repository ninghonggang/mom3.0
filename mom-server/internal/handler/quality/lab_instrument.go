package quality

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// LabInstrumentHandler 实验室仪器处理器
type LabInstrumentHandler struct {
	service *service.LabInstrumentService
}

func NewLabInstrumentHandler(s *service.LabInstrumentService) *LabInstrumentHandler {
	return &LabInstrumentHandler{service: s}
}

// ListLabInstruments 获取仪器列表
func (h *LabInstrumentHandler) ListLabInstruments(c *gin.Context) {
	var query model.LabInstrumentQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	list, total, err := h.service.List(c.Request.Context(), &query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// GetLabInstrument 获取仪器详情
func (h *LabInstrumentHandler) GetLabInstrument(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	instrument, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, instrument)
}

// CreateLabInstrument 创建仪器
func (h *LabInstrumentHandler) CreateLabInstrument(c *gin.Context) {
	var req model.LabInstrument
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = uint64(tenantID)

	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// UpdateLabInstrument 更新仪器
func (h *LabInstrumentHandler) UpdateLabInstrument(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.service.Update(c.Request.Context(), id, updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeleteLabInstrument 删除仪器
func (h *LabInstrumentHandler) DeleteLabInstrument(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetLabInstrumentCalibrations 获取仪器的校准记录
func (h *LabInstrumentHandler) GetLabInstrumentCalibrations(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var query model.LabCalibrationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	list, total, err := h.service.GetCalibrations(c.Request.Context(), id, &query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// RecordCalibration 记录校准
func (h *LabInstrumentHandler) RecordCalibration(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.RecordCalibrationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	if err := h.service.RecordCalibration(c.Request.Context(), id, &req, uint64(tenantID)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}