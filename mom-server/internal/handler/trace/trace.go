package trace

import (
	"strconv"
	"time"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type TraceHandler struct {
	traceSvc *service.TraceService
}

func NewTraceHandler(traceSvc *service.TraceService) *TraceHandler {
	return &TraceHandler{traceSvc: traceSvc}
}

func (h *TraceHandler) TraceBySerial(c *gin.Context) {
	sn := c.Query("serial_number")
	if sn == "" {
		response.BadRequest(c, "serial_number is required")
		return
	}
	snModel, records, err := h.traceSvc.TraceBySerial(c.Request.Context(), sn)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"serial":  snModel,
		"records": records,
	})
}

func (h *TraceHandler) TraceByBatch(c *gin.Context) {
	batchNo := c.Query("batch_no")
	if batchNo == "" {
		response.BadRequest(c, "batch_no is required")
		return
	}
	list, err := h.traceSvc.TraceByBatch(c.Request.Context(), batchNo)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list})
}

func (h *TraceHandler) TraceByOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid order id")
		return
	}
	list, err := h.traceSvc.TraceByOrder(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list})
}

func (h *TraceHandler) ForwardTrace(c *gin.Context) {
	sn := c.Query("serial_number")
	if sn == "" {
		response.BadRequest(c, "serial_number is required")
		return
	}
	records, err := h.traceSvc.ForwardTrace(c.Request.Context(), sn)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": records})
}

func (h *TraceHandler) BackwardTrace(c *gin.Context) {
	sn := c.Query("serial_number")
	if sn == "" {
		response.BadRequest(c, "serial_number is required")
		return
	}
	records, err := h.traceSvc.BackwardTrace(c.Request.Context(), sn)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": records})
}

type AndonHandler struct {
	andonSvc *service.AndonService
}

func NewAndonHandler(andonSvc *service.AndonService) *AndonHandler {
	return &AndonHandler{andonSvc: andonSvc}
}

func (h *AndonHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	status := c.Query("status")
	callNo := c.Query("call_no")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	query := &service.AndonQuery{
		TenantID: tenantID,
		Status:   status,
		Page:     page,
		PageSize: pageSize,
		CallNo:   callNo,
	}
	list, total, err := h.andonSvc.List(c.Request.Context(), query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *AndonHandler) Create(c *gin.Context) {
	var req model.AndonCall
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
	req.CallTime = time.Now()
	req.Status = "CALLING"
	if err := h.andonSvc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *AndonHandler) Response(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	responseBy := c.Query("response_by")
	if err := h.andonSvc.Respond(c.Request.Context(), id, responseBy, ""); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *AndonHandler) Resolve(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.andonSvc.Resolve(c.Request.Context(), id, "RESOLVED", "", nil, nil); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

type EnergyHandler struct {
	energySvc *service.EnergyService
}

func NewEnergyHandler(energySvc *service.EnergyService) *EnergyHandler {
	return &EnergyHandler{energySvc: energySvc}
}

func (h *EnergyHandler) List(c *gin.Context) {
	energyType := c.Query("energy_type")
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate time.Time
	if startDateStr != "" {
		startDate, _ = time.Parse("2006-01-02", startDateStr)
	}
	if endDateStr != "" {
		endDate, _ = time.Parse("2006-01-02", endDateStr)
	}

	list, total, err := h.energySvc.List(c.Request.Context(), energyType, startDate, endDate)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *EnergyHandler) GetStats(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	startDate, _ := time.Parse("2006-01-02", startDateStr)
	endDate, _ := time.Parse("2006-01-02", endDateStr)

	stats, err := h.energySvc.GetStats(c.Request.Context(), startDate, endDate)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, stats)
}

func (h *EnergyHandler) GetTrend(c *gin.Context) {
	energyType := c.Query("energy_type")
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	startDate, _ := time.Parse("2006-01-02", startDateStr)
	endDate, _ := time.Parse("2006-01-02", endDateStr)

	trend, err := h.energySvc.GetTrend(c.Request.Context(), energyType, startDate, endDate)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": trend})
}

func (h *EnergyHandler) Create(c *gin.Context) {
	var req model.EnergyRecord
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
	if err := h.energySvc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}
