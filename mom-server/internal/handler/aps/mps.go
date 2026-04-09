package aps

import (
	"fmt"
	"strconv"
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
	// 设置默认租户ID
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

func (h *MRPHandler) Get(c *gin.Context) {
	id := c.Param("id")
	mrp, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, mrp)
}

func (h *MRPHandler) Create(c *gin.Context) {
	var req model.MRP
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

func (h *MRPHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.MRP
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	updates := map[string]interface{}{}
	if req.MRPNo != "" {
		updates["mrp_no"] = req.MRPNo
	}
	if req.MRPType != "" {
		updates["mrp_type"] = req.MRPType
	}
	if req.PlanDate != nil {
		updates["plan_date"] = req.PlanDate
	}
	if req.Remark != nil {
		updates["remark"] = req.Remark
	}
	if req.Status != 0 {
		updates["status"] = req.Status
	}
	if err := h.service.Update(c.Request.Context(), id, updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *MRPHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *MRPHandler) Calculate(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Calculate(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Run 执行完整的MRP计算（基于MPS和BOM展开）
func (h *MRPHandler) Run(c *gin.Context) {
	var req struct {
		ID        int64  `json:"id" binding:"required"`
		PlanMonth string `json:"plan_month"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.RunMRP(c.Request.Context(), req.ID, req.PlanMonth); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetResults 获取MRP计算结果
func (h *MRPHandler) GetResults(c *gin.Context) {
	id := c.Param("id")
	var mrpID int64
	_, err := fmt.Sscanf(id, "%d", &mrpID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	items, err := h.service.GetMrpResults(c.Request.Context(), mrpID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": items, "total": len(items)})
}

// GetShortage 缺料分析
func (h *MRPHandler) GetShortage(c *gin.Context) {
	id := c.Param("id")
	var mrpID int64
	_, err := fmt.Sscanf(id, "%d", &mrpID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	items, err := h.service.AnalyzeShortage(c.Request.Context(), mrpID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": items, "total": len(items)})
}

// GetPurchaseSuggestion 获取采购建议
func (h *MRPHandler) GetPurchaseSuggestion(c *gin.Context) {
	id := c.Param("id")
	var mrpID int64
	_, err := fmt.Sscanf(id, "%d", &mrpID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	suggestions, err := h.service.GeneratePurchaseSuggestion(c.Request.Context(), mrpID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": suggestions, "total": len(suggestions)})
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

// ExecuteConstrainedRequest 带约束排程请求
type ExecuteConstrainedRequest struct {
	PlanID        int64  `json:"plan_id" binding:"required"`
	AlgorithmType string `json:"algorithm_type"` // FIFO/EDD/SPT/LPT/JIT_FIRST/CR/FAMILY/BOTTLENECK
	Direction    string `json:"direction"`        // FORWARD/BACKWARD
	WorkshopID   int64  `json:"workshop_id"`
	RespectJIT       bool    `json:"respect_jit"`
	MaxChangeoverPct float64 `json:"max_changeover_pct"`
	MinUtilization   float64 `json:"min_utilization"`
	AllowOvertime    bool    `json:"allow_overtime"`
	FamilyGrouping   bool    `json:"family_grouping"`
}

// ExecuteConstrained 执行带约束的排程
// POST /api/v1/aps/schedule/execute-constrained
func (h *ScheduleHandler) ExecuteConstrained(c *gin.Context) {
	var req ExecuteConstrainedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if req.AlgorithmType == "" {
		req.AlgorithmType = "FIFO"
	}
	if req.Direction == "" {
		req.Direction = "FORWARD"
	}
	if req.MinUtilization == 0 {
		req.MinUtilization = 70.0
	}
	if req.MaxChangeoverPct == 0 {
		req.MaxChangeoverPct = 30.0
	}

	err := h.service.ExecuteConstrainedScheduling(c.Request.Context(), service.ConstrainedScheduleRequest{
		PlanID:        req.PlanID,
		AlgorithmType: req.AlgorithmType,
		Direction:     req.Direction,
		WorkshopID:   req.WorkshopID,
		Constraints: service.ScheduleConstraints{
			RespectJIT:       req.RespectJIT,
			MaxChangeoverPct: req.MaxChangeoverPct,
			MinUtilization:   req.MinUtilization,
			AllowOvertime:    req.AllowOvertime,
			FamilyGrouping:   req.FamilyGrouping,
		},
	})
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// GetSuggestions 获取优化建议
// GET /api/v1/aps/schedule/suggestions/:plan_id
func (h *ScheduleHandler) GetSuggestions(c *gin.Context) {
	planIDStr := c.Param("plan_id")
	if planIDStr == "" {
		response.BadRequest(c, "plan_id is required")
		return
	}

	planID, err := strconv.ParseInt(planIDStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid plan_id")
		return
	}

	suggestions, err := h.service.GetOptimizationSuggestions(c.Request.Context(), planID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, suggestions)
}
