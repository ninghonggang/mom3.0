package report

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type QualityWeeklyReportHandler struct {
	svc *service.QualityWeeklyReportService
}

func NewQualityWeeklyReportHandler(svc *service.QualityWeeklyReportService) *QualityWeeklyReportHandler {
	return &QualityWeeklyReportHandler{svc: svc}
}

// ListQualityWeeklyReports GET /list
func (h *QualityWeeklyReportHandler) ListQualityWeeklyReports(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	workshopID, _ := strconv.ParseInt(c.Query("workshop_id"), 10, 64)
	year, _ := strconv.Atoi(c.Query("year"))
	week, _ := strconv.Atoi(c.Query("week"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	query := &service.QualityWeeklyReportQuery{
		TenantID:   tenantID,
		WorkshopID: workshopID,
		Year:       year,
		Week:       week,
		Page:       page,
		PageSize:   pageSize,
	}

	list, total, err := h.svc.List(c.Request.Context(), query)
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

// GetQualityWeeklyReport GET /:id
func (h *QualityWeeklyReportHandler) GetQualityWeeklyReport(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	report, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, report)
}

// GenerateWeeklyReport POST /generate
func (h *QualityWeeklyReportHandler) GenerateWeeklyReport(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var req struct {
		Year        int    `json:"year" binding:"required"`
		Week        int    `json:"week" binding:"required"`
		WorkshopID  int64  `json:"workshop_id" binding:"required"`
		WorkshopName string `json:"workshop_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	report, err := h.svc.GenerateWeeklyReport(c.Request.Context(), tenantID, req.Year, req.Week, req.WorkshopID, req.WorkshopName)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, report)
}