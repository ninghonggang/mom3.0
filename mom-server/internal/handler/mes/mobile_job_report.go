package mes

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type MobileJobReportHandler struct {
	svc *service.MobileJobReportService
}

func NewMobileJobReportHandler(svc *service.MobileJobReportService) *MobileJobReportHandler {
	return &MobileJobReportHandler{svc: svc}
}

// List 获取移动端报工列表(分页)
// GET /mes/mobile-job-report/page
func (h *MobileJobReportHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	workshopID, _ := strconv.ParseInt(c.Query("workshop_id"), 10, 64)
	productionLineID, _ := strconv.ParseInt(c.Query("production_line_id"), 10, 64)
	workstationID, _ := strconv.ParseInt(c.Query("workstation_id"), 10, 64)
	orderID, _ := strconv.ParseInt(c.Query("order_id"), 10, 64)
	employeeID, _ := strconv.ParseInt(c.Query("employee_id"), 10, 64)
	reportType, _ := strconv.Atoi(c.Query("report_type"))
	status, _ := strconv.Atoi(c.Query("status"))
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	req := &model.MobileJobReportQuery{
		WorkshopID:      workshopID,
		ProductionLineID: productionLineID,
		WorkstationID:   workstationID,
		OrderID:         orderID,
		EmployeeID:      employeeID,
		ReportType:      reportType,
		Status:          status,
		StartDate:       startDate,
		EndDate:         endDate,
		Page:            page,
		PageSize:        pageSize,
	}

	list, total, err := h.svc.List(c.Request.Context(), tenantID, req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// Get 获取单个报工
// GET /mes/mobile-job-report/:id
func (h *MobileJobReportHandler) Get(c *gin.Context) {
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

// Create 创建报工
// POST /mes/mobile-job-report
func (h *MobileJobReportHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	var req model.MobileJobReportCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	report, err := h.svc.Create(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, report)
}

// Confirm 确认报工
// PUT /mes/mobile-job-report/:id/confirm
func (h *MobileJobReportHandler) Confirm(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	userID := middleware.GetUserID(c)
	if err := h.svc.Confirm(c.Request.Context(), id, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Audit 审核报工
// PUT /mes/mobile-job-report/:id/audit
func (h *MobileJobReportHandler) Audit(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	userID := middleware.GetUserID(c)
	if err := h.svc.Audit(c.Request.Context(), id, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetPendingOrders 获取待报工工单
// GET /mes/mobile-job-report/pending-orders
func (h *MobileJobReportHandler) GetPendingOrders(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	workshopID, _ := strconv.ParseInt(c.Query("workshop_id"), 10, 64)
	employeeID, _ := strconv.ParseInt(c.Query("employee_id"), 10, 64)

	list, err := h.svc.GetPendingOrders(c.Request.Context(), tenantID, workshopID, employeeID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list})
}

// Delete 删除报工
// DELETE /mes/mobile-job-report/:id
func (h *MobileJobReportHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}