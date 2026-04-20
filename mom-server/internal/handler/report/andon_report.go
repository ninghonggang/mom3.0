package report

import (
	"strconv"
	"time"

	"mom-server/internal/middleware"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type AndonReportHandler struct {
	svc *service.AndonReportService
}

func NewAndonReportHandler(svc *service.AndonReportService) *AndonReportHandler {
	return &AndonReportHandler{svc: svc}
}

// ListAndonReports GET /list
func (h *AndonReportHandler) ListAndonReports(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	workshopID, _ := strconv.ParseInt(c.Query("workshop_id"), 10, 64)
	lineID, _ := strconv.ParseInt(c.Query("line_id"), 10, 64)
	stationID, _ := strconv.ParseInt(c.Query("station_id"), 10, 64)
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

	query := &service.AndonReportQuery{
		TenantID:   tenantID,
		WorkshopID: workshopID,
		LineID:     lineID,
		StationID:  stationID,
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

// GetAndonReport GET /:id
func (h *AndonReportHandler) GetAndonReport(c *gin.Context) {
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

// GenerateAndonReport POST /generate
func (h *AndonReportHandler) GenerateAndonReport(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var req struct {
		ReportDate   string `json:"report_date" binding:"required"`
		WorkshopID   int64  `json:"workshop_id" binding:"required"`
		WorkshopName string `json:"workshop_name"`
		LineID       int64  `json:"line_id" binding:"required"`
		LineName     string `json:"line_name"`
		StationID    int64  `json:"station_id" binding:"required"`
		StationName  string `json:"station_name"`
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

	report, err := h.svc.GenerateAndonReport(c.Request.Context(), tenantID, reportDate, req.WorkshopID, req.WorkshopName, req.LineID, req.LineName, req.StationID, req.StationName)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, report)
}