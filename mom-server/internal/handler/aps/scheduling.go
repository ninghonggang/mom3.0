package aps

import (
	"strconv"
	"time"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

// SchedulingHandler 排程处理器
type SchedulingHandler struct {
	scheduleService *service.ScheduleService
}

// NewSchedulingHandler 创建排程处理器
func NewSchedulingHandler(scheduleService *service.ScheduleService) *SchedulingHandler {
	return &SchedulingHandler{scheduleService: scheduleService}
}

// ExecuteRequest 执行排程请求
type ExecuteRequest struct {
	PlanID    int64  `json:"plan_id" binding:"required"`
	Rule      string `json:"rule"`       // FIFO/EDD/SPT/LPT
	Direction string `json:"direction"`  // FORWARD/BACKWARD
}

// ExecuteScheduling 执行排程
// POST /api/v1/aps/schedule/execute
func (h *SchedulingHandler) ExecuteScheduling(c *gin.Context) {
	var req ExecuteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if req.Rule == "" {
		req.Rule = "FIFO"
	}
	if req.Direction == "" {
		req.Direction = "FORWARD"
	}

	err := h.scheduleService.ExecuteSchedulingWithRule(c.Request.Context(), service.SchedulingRequest{
		PlanID:    req.PlanID,
		Rule:      req.Rule,
		Direction: req.Direction,
	})
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// GetResults 获取排程结果
// GET /api/v1/aps/schedule/results?plan_id=xxx
func (h *SchedulingHandler) GetResults(c *gin.Context) {
	planIDStr := c.Query("plan_id")
	if planIDStr == "" {
		response.BadRequest(c, "plan_id is required")
		return
	}

	planID, err := strconv.ParseInt(planIDStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid plan_id")
		return
	}

	results, err := h.scheduleService.GetGanttData(c.Request.Context(), planID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, results)
}

// GanttDataRequest 获取甘特图数据请求
type GanttDataRequest struct {
	PlanID int64 `form:"plan_id" binding:"required"`
}

// GetGanttData 获取甘特图数据
// GET /api/v1/aps/gantt/data?plan_id=xxx
func (h *SchedulingHandler) GetGanttData(c *gin.Context) {
	planIDStr := c.Query("plan_id")
	if planIDStr == "" {
		response.BadRequest(c, "plan_id is required")
		return
	}

	planID, err := strconv.ParseInt(planIDStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid plan_id")
		return
	}

	ganttData, err := h.scheduleService.GetGanttData(c.Request.Context(), planID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, ganttData)
}

// UpdateTaskTimeRequest 更新任务时间请求
type UpdateTaskTimeRequest struct {
	StartTime int64 `json:"start_time" binding:"required"` // Unix timestamp (秒)
	EndTime   int64 `json:"end_time" binding:"required"`   // Unix timestamp (秒)
}

// UpdateTaskTime 更新任务时间
// PUT /api/v1/aps/gantt/tasks/:id
func (h *SchedulingHandler) UpdateTaskTime(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid task id")
		return
	}

	var req UpdateTaskTimeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	startTime := time.Unix(req.StartTime, 0)
	endTime := time.Unix(req.EndTime, 0)

	err = h.scheduleService.UpdateTaskTime(c.Request.Context(), uint(id), startTime, endTime)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// GetLoadData 获取产能负载数据
// GET /api/v1/aps/load/data?plan_id=xxx
func (h *SchedulingHandler) GetLoadData(c *gin.Context) {
	planIDStr := c.Query("plan_id")
	if planIDStr == "" {
		response.BadRequest(c, "plan_id is required")
		return
	}

	planID, err := strconv.ParseInt(planIDStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid plan_id")
		return
	}

	loads, err := h.scheduleService.CalculateLoad(c.Request.Context(), planID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, gin.H{"loads": loads})
}

// OptimizeRequest 优化请求
type OptimizeRequest struct {
	PlanID int64 `json:"plan_id" binding:"required"`
}

// Optimize 优化排程
// POST /api/v1/aps/schedule/optimize
func (h *SchedulingHandler) Optimize(c *gin.Context) {
	var req OptimizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.scheduleService.OptimizeSchedule(c.Request.Context(), req.PlanID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// List 获取排程计划列表
// GET /api/v1/aps/schedule/list
func (h *SchedulingHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	list, total, err := h.scheduleService.List(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// Create 创建排程计划
// POST /api/v1/aps/schedule
func (h *SchedulingHandler) Create(c *gin.Context) {
	var req model.SchedulePlan
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.scheduleService.Create(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, req)
}

// Delete 删除排程计划
// DELETE /api/v1/aps/schedule/:id
func (h *SchedulingHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.scheduleService.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Get 获取单个排程结果详情
// GET /api/v1/aps/schedule/:id
func (h *SchedulingHandler) Get(c *gin.Context) {
	id := c.Param("id")
	results, err := h.scheduleService.GetResults(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, results)
}

