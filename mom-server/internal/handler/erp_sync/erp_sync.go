package erp_sync

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

// ERPSyncHandler ERP同步接口处理
type ERPSyncHandler struct {
	svc *service.ERPSyncService
}

// NewERPSyncHandler 创建ERP同步处理器
func NewERPSyncHandler(svc *service.ERPSyncService) *ERPSyncHandler {
	return &ERPSyncHandler{svc: svc}
}

// ListSyncLogs GET /integration/erp/sync-log/list - 查询同步日志列表
func (h *ERPSyncHandler) ListSyncLogs(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	q := &model.ERPSyncLogQuery{
		TenantID:  tenantID,
		SyncType: c.Query("syncType"),
		Direction: c.Query("direction"),
		Status:   c.Query("status"),
		StartDate: c.Query("startDate"),
		EndDate:   c.Query("endDate"),
		Page:      page,
		PageSize:  pageSize,
	}

	list, total, err := h.svc.ListSyncLogs(c.Request.Context(), q)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// GetSyncLog GET /integration/erp/sync-log/:id - 获取同步日志详情
func (h *ERPSyncHandler) GetSyncLog(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	log, err := h.svc.GetSyncLog(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, log)
}

// SyncBOM POST /integration/erp/bom/sync - BOM数据同步(ERP→MES)
func (h *ERPSyncHandler) SyncBOM(c *gin.Context) {
	var req model.ERPSyncBOMRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	result, err := h.svc.SyncBOM(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, result)
}

// SyncProductionOrder POST /integration/erp/order/sync - 生产订单同步(ERP→MES)
func (h *ERPSyncHandler) SyncProductionOrder(c *gin.Context) {
	var req model.ERPSyncProductionOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	result, err := h.svc.SyncProductionOrder(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, result)
}

// SyncStock POST /integration/erp/stock/sync - 库存同步(ERP→MES)
func (h *ERPSyncHandler) SyncStock(c *gin.Context) {
	var req model.ERPSyncStockRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	result, err := h.svc.SyncStock(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, result)
}

// PushReport POST /integration/erp/report/push - 报工回传(MES→ERP)
func (h *ERPSyncHandler) PushReport(c *gin.Context) {
	var req model.ERPPushReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	result, err := h.svc.PushReport(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, result)
}

// PushStockIn POST /integration/erp/stockin/push - 入库通知回传(MES→ERP)
func (h *ERPSyncHandler) PushStockIn(c *gin.Context) {
	var req model.ERPPushStockInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	result, err := h.svc.PushStockIn(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, result)
}

// PushQualityData POST /integration/erp/quality/push - 质检数据回传(MES→ERP)
func (h *ERPSyncHandler) PushQualityData(c *gin.Context) {
	var req model.ERPPushQualityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	result, err := h.svc.PushQualityData(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, result)
}

// GetSyncStatus GET /integration/erp/status/:syncId - 查询同步状态
func (h *ERPSyncHandler) GetSyncStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("syncId"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid syncId")
		return
	}

	status, err := h.svc.GetSyncStatus(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, status)
}
