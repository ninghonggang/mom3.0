package report

import (
	"strconv"
	"time"

	"mom-server/internal/middleware"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductionDailyReportHandler struct {
	svc *service.ProductionDailyReportService
}

func NewProductionDailyReportHandler(svc *service.ProductionDailyReportService) *ProductionDailyReportHandler {
	return &ProductionDailyReportHandler{svc: svc}
}

// ListProductionDailyReports GET /list
func (h *ProductionDailyReportHandler) ListProductionDailyReports(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	workshopID, _ := strconv.ParseInt(c.Query("workshop_id"), 10, 64)
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
			endDate = &t
		}
	}

	query := &service.ProductionDailyReportQuery{
		TenantID:   tenantID,
		WorkshopID: workshopID,
		StartDate:  startDate,
		EndDate:    endDate,
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

// GetProductionDailyReport GET /:id
func (h *ProductionDailyReportHandler) GetProductionDailyReport(c *gin.Context) {
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

// GenerateDailyReport POST /generate
func (h *ProductionDailyReportHandler) GenerateDailyReport(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var req struct {
		ReportDate   string `json:"report_date" binding:"required"`
		WorkshopID   int64  `json:"workshop_id" binding:"required"`
		WorkshopName string `json:"workshop_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	reportDate, err := time.Parse("2006-01-02", req.ReportDate)
	if err != nil {
		response.BadRequest(c, "invalid report_date format, use YYYY-MM-DD")
		return
	}

	report, err := h.svc.GenerateDailyReport(c.Request.Context(), tenantID, reportDate, req.WorkshopID, req.WorkshopName)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, report)
}

// GetDailyReportSummary GET /summary
func (h *ProductionDailyReportHandler) GetDailyReportSummary(c *gin.Context) {
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
			endDate = &t
		}
	}

	summary, err := h.svc.GetSummary(c.Request.Context(), tenantID, startDate, endDate, workshopID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, summary)
}