package mes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
)

// JobReportHandler 报工管理处理器
type JobReportHandler struct {
	svc *service.MesJobReportLogService
}

func NewJobReportHandler(svc *service.MesJobReportLogService) *JobReportHandler {
	return &JobReportHandler{svc: svc}
}

// Create POST /mes/mes-job-report-log/create
func (h *JobReportHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	userID := middleware.GetUserID(c)

	var req model.MesJobReportLogCreateReqVO
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

// Get GET /mes/mes-job-report-log/get?id=xxx
func (h *JobReportHandler) Get(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
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

// Page GET /mes/mes-job-report-log/page
func (h *JobReportHandler) Page(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)

	var req model.MesJobReportLogQueryVO
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

// Senior POST /mes/mes-job-report-log/senior
func (h *JobReportHandler) Senior(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)

	var req model.MesJobReportLogSeniorReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	list, total, err := h.svc.Senior(c.Request.Context(), tenantID, req.Conditions)
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
