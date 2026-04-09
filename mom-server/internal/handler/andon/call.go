package andon

import (
	"strconv"
	"time"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CallHandler struct {
	andonSvc *service.AndonService
	logger   *zap.Logger
}

func NewCallHandler(andonSvc *service.AndonService, logger *zap.Logger) *CallHandler {
	if logger == nil {
		logger, _ = zap.NewProduction()
	}
	return &CallHandler{andonSvc: andonSvc, logger: logger}
}

// List 查询呼叫列表
// GET /api/v1/andon/calls
func (h *CallHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	workshopID, _ := strconv.ParseInt(c.Query("workshop_id"), 10, 64)
	status := c.Query("status")
	andonType := c.Query("andon_type")
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	var startDate, endDate *time.Time
	if startDateStr != "" {
		t, err := time.Parse("2006-01-02", startDateStr)
		if err == nil {
			startDate = &t
		}
	}
	if endDateStr != "" {
		t, err := time.Parse("2006-01-02", endDateStr)
		if err == nil {
			// 设置为当天结束
			end := t.Add(24*time.Hour - time.Second)
			endDate = &end
		}
	}

	query := &service.AndonQuery{
		TenantID:   tenantID,
		WorkshopID: workshopID,
		Status:     status,
		AndonType:  andonType,
		StartDate:  startDate,
		EndDate:    endDate,
		Page:       page,
		PageSize:   pageSize,
	}

	list, total, err := h.andonSvc.List(c.Request.Context(), query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
		"page":  page,
		"page_size": pageSize,
	})
}

// Get 获取单个呼叫
// GET /api/v1/andon/calls/:id
func (h *CallHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	call, err := h.andonSvc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, call)
}

// Create 发起呼叫
// POST /api/v1/andon/calls
func (h *CallHandler) Create(c *gin.Context) {
	var req model.AndonCall
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	req.CallTime = time.Now()
	req.Status = "CALLING"
	req.CallLevel = 1

	if err := h.andonSvc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// Respond 响应呼叫
// PUT /api/v1/andon/calls/:id/respond
func (h *CallHandler) Respond(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req struct {
		ResponseBy      string `json:"response_by"`
		ResponseRemark  string `json:"response_remark"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.andonSvc.Respond(c.Request.Context(), id, req.ResponseBy, req.ResponseRemark); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Resolve 处理完成
// PUT /api/v1/andon/calls/:id/resolve
func (h *CallHandler) Resolve(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req struct {
		HandleResult    string `json:"handle_result" binding:"required"` // RESOLVED/CARRY_OVER
		HandleRemark    string `json:"handle_remark"`
		RelatedRepairID *int64 `json:"related_repair_id"`
		RelatedNCRID   *int64 `json:"related_ncr_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.andonSvc.Resolve(c.Request.Context(), id, req.HandleResult, req.HandleRemark, req.RelatedRepairID, req.RelatedNCRID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Escalate 手动升级
// PUT /api/v1/andon/calls/:id/escalate
func (h *CallHandler) Escalate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req struct {
		ToLevel          int    `json:"escalate_to_level"` // 升级到哪个等级
		EscalationReason string `json:"escalation_reason"`
		TriggerUser      string `json:"trigger_user"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if req.ToLevel <= 0 {
		req.ToLevel = 0 // 默认为下一级
	}

	if err := h.andonSvc.Escalate(c.Request.Context(), id, req.ToLevel, req.EscalationReason, req.TriggerUser); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetStatistics 获取统计数据
// GET /api/v1/andon/statistics
func (h *CallHandler) GetStatistics(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	workshopID, _ := strconv.ParseInt(c.Query("workshop_id"), 10, 64)
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate *time.Time
	if startDateStr != "" {
		t, err := time.Parse("2006-01-02", startDateStr)
		if err == nil {
			startDate = &t
		}
	}
	if endDateStr != "" {
		t, err := time.Parse("2006-01-02", endDateStr)
		if err == nil {
			end := t.Add(24*time.Hour - time.Second)
			endDate = &end
		}
	}

	stats, err := h.andonSvc.GetStatistics(c.Request.Context(), tenantID, workshopID, startDate, endDate)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, stats)
}
