package wms

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// WMSPickHandler 拣货作业处理器
type WMSPickHandler struct {
	service *service.WMSPickService
}

func NewWMSPickHandler(s *service.WMSPickService) *WMSPickHandler {
	return &WMSPickHandler{service: s}
}

// Create 创建拣货作业
func (h *WMSPickHandler) Create(c *gin.Context) {
	var req model.WMSPickJobCreateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	job, err := h.service.Create(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, job)
}

// Assign 分配拣货人
func (h *WMSPickHandler) Assign(c *gin.Context) {
	var req model.WMSPickJobAssignReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.service.Assign(c.Request.Context(), uint(req.Id), req.PickerID, req.PickerName); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// List 查询拣货作业列表
func (h *WMSPickHandler) List(c *gin.Context) {
	query := c.Query("query")
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	list, total, err := h.service.List(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// Get 获取拣货作业详情
func (h *WMSPickHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	resp, err := h.service.GetWithRecords(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, resp)
}

// Start 开始拣货
func (h *WMSPickHandler) Start(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Start(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Complete 完成拣货
func (h *WMSPickHandler) Complete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Complete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Cancel 取消拣货作业
func (h *WMSPickHandler) Cancel(c *gin.Context) {
	var req model.WMSPickJobCancelReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.service.Cancel(c.Request.Context(), uint(req.Id), req.Reason); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
