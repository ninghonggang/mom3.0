package integration

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type IdocHandler struct {
	svc *service.IdocService
}

func NewIdocHandler(svc *service.IdocService) *IdocHandler {
	return &IdocHandler{svc: svc}
}

// List 获取IDOC列表(分页)
// GET /integration/idoc/page
func (h *IdocHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	idocType := c.Query("idoc_type")
	direction, _ := strconv.Atoi(c.Query("direction"))
	status, _ := strconv.Atoi(c.Query("status"))
	partnerNo := c.Query("partner_no")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	req := &model.IdocQuery{
		IdocType:  idocType,
		Direction: direction,
		Status:    status,
		PartnerNo: partnerNo,
		StartDate: startDate,
		EndDate:   endDate,
		Page:      page,
		PageSize:  pageSize,
	}

	list, total, err := h.svc.List(c.Request.Context(), tenantID, req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// Get 获取单个IDOC
// GET /integration/idoc/:id
func (h *IdocHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	record, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, record)
}

// Receive 接收IDOC
// POST /integration/idoc/receive
func (h *IdocHandler) Receive(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	var req model.IdocReceiveReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	record, err := h.svc.Receive(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, record)
}

// Send 发送IDOC
// POST /integration/idoc/send
func (h *IdocHandler) Send(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	var req model.IdocSendReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	record, err := h.svc.Send(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, record)
}

// Retry 重试IDOC
// POST /integration/idoc/:id/retry
func (h *IdocHandler) Retry(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.svc.Retry(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ListConfigs 获取IDOC类型配置列表
// GET /integration/idoc/configs
func (h *IdocHandler) ListConfigs(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	list, err := h.svc.ListConfigs(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list})
}