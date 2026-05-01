package eam

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
)

// EamRepairJobHandler 维修工单/流程/标准Handler
type EamRepairJobHandler struct {
	svc *service.EamRepairJobService
}

// NewEamRepairJobHandler 创建维修Handler
func NewEamRepairJobHandler(svc *service.EamRepairJobService) *EamRepairJobHandler {
	return &EamRepairJobHandler{svc: svc}
}

// ========== 维修工单 ==========

// CreateJob 创建维修工单
// @Summary 创建维修工单
// @Tags EAM-维修管理
// @Param body body model.EamRepairJobCreateReq true "维修工单信息"
// @Success 200 {object} response.Response
// @Router /eam/repair-job/create [post]
func (h *EamRepairJobHandler) CreateJob(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}
	var req model.EamRepairJobCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	job, err := h.svc.CreateJob(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, job)
}

// UpdateJob 更新维修工单
// @Summary 更新维修工单
// @Tags EAM-维修管理
// @Param body body model.EamRepairJobUpdateReq true "维修工单信息"
// @Success 200 {object} response.Response
// @Router /eam/repair-job/update [put]
func (h *EamRepairJobHandler) UpdateJob(c *gin.Context) {
	var req model.EamRepairJobUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	if err := h.svc.UpdateJob(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeleteJob 删除维修工单
// @Summary 删除维修工单
// @Tags EAM-维修管理
// @Param id query int true "ID"
// @Success 200 {object} response.Response
// @Router /eam/repair-job/delete [delete]
func (h *EamRepairJobHandler) DeleteJob(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil || id <= 0 {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	if err := h.svc.DeleteJob(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetJob 获取维修工单详情
// @Summary 获取维修工单详情
// @Tags EAM-维修管理
// @Param id query int true "ID"
// @Success 200 {object} response.Response
// @Router /eam/repair-job/get [get]
func (h *EamRepairJobHandler) GetJob(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil || id <= 0 {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	job, err := h.svc.GetJob(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, job)
}

// PageJob 分页查询维修工单
// @Summary 分页查询维修工单
// @Tags EAM-维修管理
// @Param job_code query string false "工单编号"
// @Param status query string false "状态"
// @Param equipment_id query int false "设备ID"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.Response
// @Router /eam/repair-job/page [get]
func (h *EamRepairJobHandler) PageJob(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}
	log.Printf("[DEBUG] PageJob: tenantID=%d, h.svc=%v", tenantID, h.svc)
	var req model.EamRepairJobPageReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	log.Printf("[DEBUG] PageJob calling svc.PageJob: tenantID=%d, req=%+v", tenantID, req)
	list, total, err := h.svc.PageJob(c.Request.Context(), tenantID, &req)
	log.Printf("[DEBUG] PageJob returned: listLen=%d, total=%d, err=%v", len(list), total, err)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}
	response.PageSuccess(c, list, total, page, pageSize)
}

// AssignJob 派工
// @Summary 派工
// @Tags EAM-维修管理
// @Param body body model.EamRepairJobAssignReq true "派工信息"
// @Success 200 {object} response.Response
// @Router /eam/repair-job/assign [post]
func (h *EamRepairJobHandler) AssignJob(c *gin.Context) {
	var req model.EamRepairJobAssignReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	if err := h.svc.AssignJob(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// AcceptJob 接单
// @Summary 接单
// @Tags EAM-维修管理
// @Param body body object true "{"id": 1}"
// @Success 200 {object} response.Response
// @Router /eam/repair-job/accept [post]
func (h *EamRepairJobHandler) AcceptJob(c *gin.Context) {
	var req struct {
		ID int64 `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	if err := h.svc.AcceptJob(c.Request.Context(), req.ID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// CompleteJob 完工
// @Summary 完工
// @Tags EAM-维修管理
// @Param body body model.EamRepairJobCompleteReq true "完工信息"
// @Success 200 {object} response.Response
// @Router /eam/repair-job/complete [post]
func (h *EamRepairJobHandler) CompleteJob(c *gin.Context) {
	var req model.EamRepairJobCompleteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	if err := h.svc.CompleteJob(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// EvaluateJob 评价
// @Summary 评价维修工单
// @Tags EAM-维修管理
// @Param body body model.EamRepairJobEvaluateReq true "评价信息"
// @Success 200 {object} response.Response
// @Router /eam/repair-job/evaluate [post]
func (h *EamRepairJobHandler) EvaluateJob(c *gin.Context) {
	var req model.EamRepairJobEvaluateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	if err := h.svc.EvaluateJob(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ========== 维修流程 ==========

// CreateFlow 创建维修流程
// @Summary 创建维修流程
// @Tags EAM-维修管理
// @Param body body model.EamRepairFlowCreateReq true "维修流程信息"
// @Success 200 {object} response.Response
// @Router /eam/repair-flow/create [post]
func (h *EamRepairJobHandler) CreateFlow(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}
	var req model.EamRepairFlowCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	flow, err := h.svc.CreateFlow(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, flow)
}

// UpdateFlow 更新维修流程
// @Summary 更新维修流程
// @Tags EAM-维修管理
// @Param body body model.EamRepairFlowUpdateReq true "维修流程信息"
// @Success 200 {object} response.Response
// @Router /eam/repair-flow/update [put]
func (h *EamRepairJobHandler) UpdateFlow(c *gin.Context) {
	var req model.EamRepairFlowUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	if err := h.svc.UpdateFlow(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeleteFlow 删除维修流程
// @Summary 删除维修流程
// @Tags EAM-维修管理
// @Param id query int true "ID"
// @Success 200 {object} response.Response
// @Router /eam/repair-flow/delete [delete]
func (h *EamRepairJobHandler) DeleteFlow(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil || id <= 0 {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	if err := h.svc.DeleteFlow(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetFlow 获取维修流程详情
// @Summary 获取维修流程详情
// @Tags EAM-维修管理
// @Param id query int true "ID"
// @Success 200 {object} response.Response
// @Router /eam/repair-flow/get [get]
func (h *EamRepairJobHandler) GetFlow(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil || id <= 0 {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	flow, err := h.svc.GetFlow(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, flow)
}

// PageFlow 分页查询维修流程
// @Summary 分页查询维修流程
// @Tags EAM-维修管理
// @Param flow_code query string false "流程编号"
// @Param flow_name query string false "流程名称"
// @Param status query string false "状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.Response
// @Router /eam/repair-flow/page [get]
func (h *EamRepairJobHandler) PageFlow(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}
	var req model.EamRepairFlowPageReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	list, total, err := h.svc.PageFlow(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}
	response.PageSuccess(c, list, total, page, pageSize)
}

// ========== 维修标准 ==========

// CreateStd 创建维修标准
// @Summary 创建维修标准
// @Tags EAM-维修管理
// @Param body body model.EamRepairStdCreateReq true "维修标准信息"
// @Success 200 {object} response.Response
// @Router /eam/repair-std/create [post]
func (h *EamRepairJobHandler) CreateStd(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}
	var req model.EamRepairStdCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	std, err := h.svc.CreateStd(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, std)
}

// UpdateStd 更新维修标准
// @Summary 更新维修标准
// @Tags EAM-维修管理
// @Param body body model.EamRepairStdUpdateReq true "维修标准信息"
// @Success 200 {object} response.Response
// @Router /eam/repair-std/update [put]
func (h *EamRepairJobHandler) UpdateStd(c *gin.Context) {
	var req model.EamRepairStdUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	if err := h.svc.UpdateStd(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeleteStd 删除维修标准
// @Summary 删除维修标准
// @Tags EAM-维修管理
// @Param id query int true "ID"
// @Success 200 {object} response.Response
// @Router /eam/repair-std/delete [delete]
func (h *EamRepairJobHandler) DeleteStd(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil || id <= 0 {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	if err := h.svc.DeleteStd(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetStd 获取维修标准详情
// @Summary 获取维修标准详情
// @Tags EAM-维修管理
// @Param id query int true "ID"
// @Success 200 {object} response.Response
// @Router /eam/repair-std/get [get]
func (h *EamRepairJobHandler) GetStd(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil || id <= 0 {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	std, err := h.svc.GetStd(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, std)
}

// PageStd 分页查询维修标准
// @Summary 分页查询维修标准
// @Tags EAM-维修管理
// @Param std_code query string false "标准编号"
// @Param std_name query string false "标准名称"
// @Param fault_type query string false "故障类型"
// @Param status query string false "状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.Response
// @Router /eam/repair-std/page [get]
func (h *EamRepairJobHandler) PageStd(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}
	var req model.EamRepairStdPageReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	list, total, err := h.svc.PageStd(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}
	response.PageSuccess(c, list, total, page, pageSize)
}
