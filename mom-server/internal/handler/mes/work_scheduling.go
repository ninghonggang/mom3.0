package mes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
)

// WorkSchedulingHandler 工单排程处理器
type WorkSchedulingHandler struct {
	svc *service.MesWorkSchedulingService
}

func NewWorkSchedulingHandler(svc *service.MesWorkSchedulingService) *WorkSchedulingHandler {
	return &WorkSchedulingHandler{svc: svc}
}

// ==================== 工单排程主表 ====================

// Create POST /mes/work-scheduling/create
func (h *WorkSchedulingHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	userID := middleware.GetUserID(c)

	var req model.MesWorkSchedulingCreateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	m, err := h.svc.Create(c.Request.Context(), tenantID, userID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, m)
}

// Update PUT /mes/work-scheduling/update
func (h *WorkSchedulingHandler) Update(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req model.WorkScheduleUpdateVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	m, err := h.svc.Update(c.Request.Context(), userID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, m)
}

// Delete DELETE /mes/work-scheduling/delete?id=xxx
func (h *WorkSchedulingHandler) Delete(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		response.BadRequest(c, "无效的ID")
		return
	}

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Get GET /mes/work-scheduling/get?id=xxx
func (h *WorkSchedulingHandler) Get(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		response.BadRequest(c, "无效的ID")
		return
	}

	m, err := h.svc.Get(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, m)
}

// Page GET /mes/work-scheduling/page
func (h *WorkSchedulingHandler) Page(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)

	var req model.WorkSchedulePageVO
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	list, total, err := h.svc.Page(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}
	response.PageSuccess(c, list, total, page, pageSize)
}

// ==================== 工序排程明细 ====================

// CreateDetail POST /mes/work-scheduling-detail/create
func (h *WorkSchedulingHandler) CreateDetail(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	userID := middleware.GetUserID(c)

	var req model.MesWorkSchedulingDetailCreateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	d, err := h.svc.CreateDetail(c.Request.Context(), tenantID, userID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, d)
}

// UpdateDetail PUT /mes/work-scheduling-detail/update
func (h *WorkSchedulingHandler) UpdateDetail(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req model.MesWorkSchedulingDetailUpdateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	d, err := h.svc.UpdateDetail(c.Request.Context(), userID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, d)
}

// DeleteDetail DELETE /mes/work-scheduling-detail/delete?id=xxx
func (h *WorkSchedulingHandler) DeleteDetail(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		response.BadRequest(c, "无效的ID")
		return
	}

	if err := h.svc.DeleteDetail(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetDetail GET /mes/work-scheduling-detail/get?id=xxx
func (h *WorkSchedulingHandler) GetDetail(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		response.BadRequest(c, "无效的ID")
		return
	}

	d, err := h.svc.GetDetail(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, d)
}

// PageDetail GET /mes/work-scheduling-detail/page
func (h *WorkSchedulingHandler) PageDetail(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)

	var req model.MesWorkSchedulingDetailPageReqVO
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	list, total, err := h.svc.PageDetail(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}
	response.PageSuccess(c, list, total, page, pageSize)
}

// ListDetail GET /mes/work-scheduling-detail/list?scheduling_id=xxx
func (h *WorkSchedulingHandler) ListDetail(c *gin.Context) {
	idStr := c.Query("scheduling_id")
	schedulingID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || schedulingID <= 0 {
		response.BadRequest(c, "无效的scheduling_id")
		return
	}

	list, err := h.svc.ListDetailBySchedulingID(c.Request.Context(), schedulingID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": len(list)})
}

// ==================== 工序操作 ====================

// StartDetail PUT /mes/work-scheduling-detail/start
func (h *WorkSchedulingHandler) StartDetail(c *gin.Context) {
	userID := middleware.GetUserID(c)
	var req struct {
		ID int64 `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.StartDetail(c.Request.Context(), req.ID, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// PauseDetail PUT /mes/work-scheduling-detail/pause
func (h *WorkSchedulingHandler) PauseDetail(c *gin.Context) {
	userID := middleware.GetUserID(c)
	var req struct {
		ID int64 `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.PauseDetail(c.Request.Context(), req.ID, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ResumeDetail PUT /mes/work-scheduling-detail/resume
func (h *WorkSchedulingHandler) ResumeDetail(c *gin.Context) {
	userID := middleware.GetUserID(c)
	var req struct {
		ID int64 `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.ResumeDetail(c.Request.Context(), req.ID, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// CompleteDetail PUT /mes/work-scheduling-detail/complete
func (h *WorkSchedulingHandler) CompleteDetail(c *gin.Context) {
	userID := middleware.GetUserID(c)
	var req struct {
		ID int64 `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.CompleteDetail(c.Request.Context(), req.ID, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ReportDetail POST /mes/work-scheduling-detail/report
func (h *WorkSchedulingHandler) ReportDetail(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req model.MesWorkSchedulingDetailReportReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.ReportDetail(c.Request.Context(), req.ID, userID, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// BindEquipment PUT /mes/work-scheduling-detail/bindEquipment
func (h *WorkSchedulingHandler) BindEquipment(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req model.MesWorkSchedulingDetailBindEquipmentReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.BindEquipment(c.Request.Context(), userID, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// BindWorker PUT /mes/work-scheduling-detail/bindWorker
func (h *WorkSchedulingHandler) BindWorker(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req model.MesWorkSchedulingDetailBindWorkerReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.BindWorker(c.Request.Context(), userID, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
