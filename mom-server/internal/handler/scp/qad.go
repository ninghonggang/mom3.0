package scp

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

// QadHandler QAD对接接口处理
type QadHandler struct {
	svc *service.ScpQadService
}

// NewQadHandler 创建QAD处理器
func NewQadHandler(svc *service.ScpQadService) *QadHandler {
	return &QadHandler{svc: svc}
}

// Sync POST /scp/qad/sync - 同步数据到QAD
func (h *QadHandler) Sync(c *gin.Context) {
	var req model.QadSyncRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	result, err := h.svc.SyncToQad(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, result)
}

// GetSyncStatus GET /scp/qad/sync/status/:syncId - 同步状态查询
func (h *QadHandler) GetSyncStatus(c *gin.Context) {
	idStr := c.Param("syncId")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid syncId")
		return
	}

	log, err := h.svc.GetSyncStatus(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, log)
}

// GetSyncLog GET /scp/qad/sync/log/:docNo - 同步日志查询
func (h *QadHandler) GetSyncLog(c *gin.Context) {
	docNo := c.Param("docNo")
	if docNo == "" {
		response.BadRequest(c, "docNo is required")
		return
	}

	logs, err := h.svc.GetSyncLogByDocNo(c.Request.Context(), docNo)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, logs)
}

// Confirm POST /scp/qad/confirm - QAD订单确认回调
func (h *QadHandler) Confirm(c *gin.Context) {
	var req model.QadConfirmRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	if err := h.svc.HandleQadConfirm(c.Request.Context(), tenantID, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Delivery POST /scp/qad/delivery - QAD发货通知回调
func (h *QadHandler) Delivery(c *gin.Context) {
	var req model.QadDeliveryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	if err := h.svc.HandleQadDelivery(c.Request.Context(), tenantID, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
