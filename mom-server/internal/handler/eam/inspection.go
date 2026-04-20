package eam

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
)

// EAMInspectionHandler 巡检Handler
type EAMInspectionHandler struct {
	svc *service.EAMInspectionService
}

// NewEAMInspectionHandler 创建巡检Handler
func NewEAMInspectionHandler(svc *service.EAMInspectionService) *EAMInspectionHandler {
	return &EAMInspectionHandler{svc: svc}
}

// ========== 巡检计划 ==========

// CreatePlan 创建巡检计划
// @Summary 创建巡检计划
// @Tags EAM-巡检管理
// @Param body body model.EAMInspectionPlanCreateReqVO true "巡检计划信息"
// @Success 200 {object} response.Response
// @Router /eam/inspection/plan/create [post]
func (h *EAMInspectionHandler) CreatePlan(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var req model.EAMInspectionPlanCreateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	plan, err := h.svc.CreatePlanWithItems(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, plan)
}

// UpdatePlan 更新巡检计划
// @Summary 更新巡检计划
// @Tags EAM-巡检管理
// @Param id path int true "ID"
// @Param body body model.EAMInspectionPlanCreateReqVO true "巡检计划信息"
// @Success 200 {object} response.Response
// @Router /eam/inspection/plan/update [put]
func (h *EAMInspectionHandler) UpdatePlan(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req model.EAMInspectionPlanCreateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	if err := h.svc.UpdatePlan(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeletePlan 删除巡检计划
// @Summary 删除巡检计划
// @Tags EAM-巡检管理
// @Param id path int true "ID"
// @Success 200 {object} response.Response
// @Router /eam/inspection/plan/:id [delete]
func (h *EAMInspectionHandler) DeletePlan(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.svc.DeletePlan(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ListPlan 巡检计划列表
// @Summary 巡检计划列表
// @Tags EAM-巡检管理
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Param plan_no query string false "计划编号"
// @Param plan_name query string false "计划名称"
// @Param status query string false "状态"
// @Success 200 {object} response.Response
// @Router /eam/inspection/plan/list [get]
func (h *EAMInspectionHandler) ListPlan(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	offset := (page - 1) * pageSize

	filters := map[string]interface{}{}
	if planNo := c.Query("plan_no"); planNo != "" {
		filters["plan_no"] = planNo
	}
	if planName := c.Query("plan_name"); planName != "" {
		filters["plan_name"] = planName
	}
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}

	list, total, err := h.svc.ListPlan(c.Request.Context(), offset, pageSize, filters)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// GetPlan 获取巡检计划详情
// @Summary 获取巡检计划详情
// @Tags EAM-巡检管理
// @Param id path int true "ID"
// @Success 200 {object} response.Response
// @Router /eam/inspection/plan/:id [get]
func (h *EAMInspectionHandler) GetPlan(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	plan, items, err := h.svc.GetPlanWithItems(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"plan":  plan,
		"items": items,
	})
}

// UpdateItems 更新巡检项目
// @Summary 更新巡检项目
// @Tags EAM-巡检管理
// @Param id path int true "计划ID"
// @Param body body []model.EAMInspectionItemReqVO true "巡检项目列表"
// @Success 200 {object} response.Response
// @Router /eam/inspection/plan/:id/items [put]
func (h *EAMInspectionHandler) UpdateItems(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var req []model.EAMInspectionItemReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	items := make([]model.EAMInspectionItem, len(req))
	for i, item := range req {
		items[i] = model.EAMInspectionItem{
			TenantID:      tenantID,
			PlanID:        id,
			ItemCode:      item.ItemCode,
			ItemName:      item.ItemName,
			CheckMethod:   item.CheckMethod,
			CheckStandard: item.CheckStandard,
			IsRequired:    item.IsRequired,
			SortOrder:     item.SortOrder,
		}
	}

	if err := h.svc.UpdateItems(c.Request.Context(), id, items); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ========== 巡检方案(执行) ==========

// CreateScheme 创建巡检方案
// @Summary 创建巡检方案
// @Tags EAM-巡检管理
// @Param body body model.EAMInspectionSchemeCreateReqVO true "巡检方案信息"
// @Success 200 {object} response.Response
// @Router /eam/inspection/scheme/create [post]
func (h *EAMInspectionHandler) CreateScheme(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var req model.EAMInspectionSchemeCreateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	scheme, err := h.svc.CreateScheme(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, scheme)
}

// ListScheme 巡检方案列表
// @Summary 巡检方案列表
// @Tags EAM-巡检管理
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Param plan_no query string false "计划编号"
// @Param status query string false "状态"
// @Success 200 {object} response.Response
// @Router /eam/inspection/scheme/list [get]
func (h *EAMInspectionHandler) ListScheme(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	offset := (page - 1) * pageSize

	filters := map[string]interface{}{}
	if planNo := c.Query("plan_no"); planNo != "" {
		filters["plan_no"] = planNo
	}
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}

	list, total, err := h.svc.ListScheme(c.Request.Context(), offset, pageSize, filters)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// GetScheme 获取巡检方案详情
// @Summary 获取巡检方案详情
// @Tags EAM-巡检管理
// @Param id path int true "ID"
// @Success 200 {object} response.Response
// @Router /eam/inspection/scheme/:id [get]
func (h *EAMInspectionHandler) GetScheme(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	scheme, results, err := h.svc.GetSchemeWithResults(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"scheme":  scheme,
		"results": results,
	})
}

// StartScheme 开始巡检
// @Summary 开始巡检
// @Tags EAM-巡检管理
// @Param id path int true "ID"
// @Success 200 {object} response.Response
// @Router /eam/inspection/scheme/:id/start [put]
func (h *EAMInspectionHandler) StartScheme(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.svc.StartScheme(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// CompleteScheme 完成巡检
// @Summary 完成巡检
// @Tags EAM-巡检管理
// @Param id path int true "ID"
// @Success 200 {object} response.Response
// @Router /eam/inspection/scheme/:id/complete [put]
func (h *EAMInspectionHandler) CompleteScheme(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req struct {
		Results []model.EAMInspectionResultItemReqVO `json:"results"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	if err := h.svc.CompleteScheme(c.Request.Context(), id, req.Results); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// SubmitResult 提交巡检结果
// @Summary 提交巡检结果
// @Tags EAM-巡检管理
// @Param body body model.EAMInspectionResultSubmitReqVO true "巡检结果"
// @Success 200 {object} response.Response
// @Router /eam/inspection/scheme/:id/result [put]
func (h *EAMInspectionHandler) SubmitResult(c *gin.Context) {
	var req model.EAMInspectionResultSubmitReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	if err := h.svc.SubmitResult(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}