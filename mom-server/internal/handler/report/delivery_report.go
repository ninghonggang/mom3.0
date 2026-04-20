package report

import (
	"strconv"
	"time"

	"mom-server/internal/middleware"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type DeliveryReportHandler struct {
	svc *service.DeliveryReportService
}

func NewDeliveryReportHandler(svc *service.DeliveryReportService) *DeliveryReportHandler {
	return &DeliveryReportHandler{svc: svc}
}

// ListDeliveryReports GET /list
func (h *DeliveryReportHandler) ListDeliveryReports(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	customerID, _ := strconv.ParseInt(c.Query("customer_id"), 10, 64)
	startMonthStr := c.Query("start_month")
	endMonthStr := c.Query("end_month")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	var startMonth, endMonth *time.Time
	if startMonthStr != "" {
		t, err := time.Parse("2006-01", startMonthStr)
		if err == nil {
			startMonth = &t
		}
	}
	if endMonthStr != "" {
		t, err := time.Parse("2006-01", endMonthStr)
		if err == nil {
			endMonth = &t
		}
	}

	query := &service.DeliveryReportQuery{
		TenantID:   tenantID,
		CustomerID: customerID,
		StartMonth: startMonth,
		EndMonth:   endMonth,
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

// GetDeliveryReport GET /:id
func (h *DeliveryReportHandler) GetDeliveryReport(c *gin.Context) {
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

// GenerateDeliveryReport POST /generate
func (h *DeliveryReportHandler) GenerateDeliveryReport(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var req struct {
		ReportMonth  string `json:"report_month" binding:"required"`
		CustomerID   int64  `json:"customer_id" binding:"required"`
		CustomerName string `json:"customer_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	reportMonth, err := time.Parse("2006-01", req.ReportMonth)
	if err != nil {
		response.BadRequest(c, "invalid report_month format, use YYYY-MM")
		return
	}

	report, err := h.svc.GenerateDeliveryReport(c.Request.Context(), tenantID, reportMonth, req.CustomerID, req.CustomerName)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, report)
}