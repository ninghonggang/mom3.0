package ai

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/service"
	"mom-server/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

// VisualInspectionHandler 视觉检测处理器
type VisualInspectionHandler struct {
	svc *service.VisualInspectionService
}

// NewVisualInspectionHandler 创建视觉检测处理器
func NewVisualInspectionHandler(svc *service.VisualInspectionService) *VisualInspectionHandler {
	return &VisualInspectionHandler{svc: svc}
}

// ListVisualInspectionTasks 获取任务列表
// GET /api/v1/ai/visual-inspection/list
func (h *VisualInspectionHandler) ListVisualInspectionTasks(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)

	// 解析分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	// 解析过滤参数
	taskType := c.Query("task_type")
	status := c.Query("status")
	productID, _ := strconv.ParseInt(c.Query("product_id"), 10, 64)

	tasks, total, err := h.svc.ListTasks(c.Request.Context(), tenantID, taskType, status, productID, page, pageSize)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, gin.H{
		"list":  tasks,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// GetVisualInspectionTask 获取任务详情
// GET /api/v1/ai/visual-inspection/:id
func (h *VisualInspectionHandler) GetVisualInspectionTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的任务ID")
		return
	}

	task, err := h.svc.GetTask(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, "任务不存在")
		return
	}

	response.Success(c, task)
}

// CreateVisualInspectionTask 创建检测任务
// POST /api/v1/ai/visual-inspection
func (h *VisualInspectionHandler) CreateVisualInspectionTask(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	userID := middleware.GetUserID(c)

	var req model.CreateVisualInspectionTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	task, err := h.svc.CreateTask(c.Request.Context(), &req, tenantID, userID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, task)
}

// DeleteVisualInspectionTask 删除任务
// DELETE /api/v1/ai/visual-inspection/:id
func (h *VisualInspectionHandler) DeleteVisualInspectionTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的任务ID")
		return
	}

	if err := h.svc.DeleteTask(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// GetVisualInspectionResult 获取检测结果
// GET /api/v1/ai/visual-inspection/:id/result
func (h *VisualInspectionHandler) GetVisualInspectionResult(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的任务ID")
		return
	}

	result, err := h.svc.GetResult(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, "检测结果不存在")
		return
	}

	response.Success(c, result)
}

// ManualReview 人工复核
// POST /api/v1/ai/visual-inspection/:id/manual-review
func (h *VisualInspectionHandler) ManualReview(c *gin.Context) {
	userID := middleware.GetUserID(c)

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的任务ID")
		return
	}

	var req model.ManualReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.ManualReview(c.Request.Context(), uint(id), userID, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, "复核成功")
}

// GetVisualInspectionStats 获取统计数据
// GET /api/v1/ai/visual-inspection/stats
func (h *VisualInspectionHandler) GetVisualInspectionStats(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)

	stats, err := h.svc.GetStats(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, stats)
}
