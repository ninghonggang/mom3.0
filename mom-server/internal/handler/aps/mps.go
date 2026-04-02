package aps

import (
	"fmt"
	"time"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type MPSHandler struct {
	service *service.MPSService
}

func NewMPSHandler(s *service.MPSService) *MPSHandler {
	return &MPSHandler{service: s}
}

func (h *MPSHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	list, total, err := h.service.List(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *MPSHandler) Get(c *gin.Context) {
	id := c.Param("id")
	mps, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, mps)
}

func (h *MPSHandler) Create(c *gin.Context) {
	var req model.MPS
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *MPSHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.MPS
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

func (h *MPSHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *MPSHandler) Submit(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Submit(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

type MRPHandler struct {
	service *service.MRPService
}

func NewMRPHandler(s *service.MRPService) *MRPHandler {
	return &MRPHandler{service: s}
}

func (h *MRPHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	list, total, err := h.service.List(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *MRPHandler) Calculate(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Calculate(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

type ScheduleHandler struct {
	service *service.ScheduleService
}

func NewScheduleHandler(s *service.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{service: s}
}

func (h *ScheduleHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	list, total, err := h.service.List(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *ScheduleHandler) Create(c *gin.Context) {
	var req model.SchedulePlan
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *ScheduleHandler) Execute(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Execute(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *ScheduleHandler) GetResults(c *gin.Context) {
	id := c.Param("id")
	results, err := h.service.GetResults(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, results)
}

func (h *ScheduleHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DragUpdateRequest 拖拽更新请求
type DragUpdateRequest struct {
	ResultID      uint  `json:"result_id" binding:"required"`
	LineID        int64 `json:"line_id"`
	StationID     int64 `json:"station_id"`
	PlanStartTime int64 `json:"plan_start_time"` // 时间戳(秒)
	PlanEndTime   int64 `json:"plan_end_time"`   // 时间戳(秒)
}

func (h *ScheduleHandler) DragUpdate(c *gin.Context) {
	var req DragUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	planStartTime := time.Unix(req.PlanStartTime, 0)
	planEndTime := time.Unix(req.PlanEndTime, 0)
	err := h.service.DragUpdate(c.Request.Context(), req.ResultID, req.LineID, req.StationID, planStartTime, planEndTime)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	result, err := h.service.GetResultByID(c.Request.Context(), fmt.Sprintf("%d", req.ResultID))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, result)
}
