package bpm

import (
	"fmt"
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

func toInt64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

func toInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

type BPMHandler struct {
	bpmSvc *service.BPMService
}

func NewBPMHandler(bpmSvc *service.BPMService) *BPMHandler {
	return &BPMHandler{bpmSvc: bpmSvc}
}

// ==================== 流程模型 ====================

func (h *BPMHandler) ListProcessModels(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if modelType := c.Query("model_type"); modelType != "" {
		query["model_type"] = modelType
	}
	if category := c.Query("category"); category != "" {
		query["category"] = category
	}
	if page := c.Query("page"); page != "" {
		query["page"] = toInt(page)
	}
	if limit := c.Query("limit"); limit != "" {
		query["limit"] = toInt(limit)
	}

	list, total, err := h.bpmSvc.ListProcessModels(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *BPMHandler) GetProcessModel(c *gin.Context) {
	id := c.Param("id")
	model, err := h.bpmSvc.GetProcessModel(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, model)
}

func (h *BPMHandler) CreateProcessModel(c *gin.Context) {
	var req model.ProcessModel
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID

	if err := h.bpmSvc.CreateProcessModel(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *BPMHandler) UpdateProcessModel(c *gin.Context) {
	id := c.Param("id")
	var req model.ProcessModel
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.bpmSvc.UpdateProcessModel(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *BPMHandler) DeleteProcessModel(c *gin.Context) {
	id := c.Param("id")
	if err := h.bpmSvc.DeleteProcessModel(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *BPMHandler) PublishProcessModel(c *gin.Context) {
	id := c.Param("id")
	userID := middleware.GetUserID(c)

	if err := h.bpmSvc.PublishProcessModel(c.Request.Context(), id, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ==================== 流程节点 ====================

func (h *BPMHandler) ListNodes(c *gin.Context) {
	modelID := c.Param("id")
	nodes, err := h.bpmSvc.ListNodes(c.Request.Context(), modelID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": nodes})
}

func (h *BPMHandler) CreateNode(c *gin.Context) {
	modelID := c.Param("id")
	var req model.NodeDefinition
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var modelIDInt int64
	fmt.Sscanf(modelID, "%d", &modelIDInt)
	req.ModelID = modelIDInt
	req.TenantID = tenantID

	if err := h.bpmSvc.CreateNode(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *BPMHandler) UpdateNode(c *gin.Context) {
	id := c.Param("id")
	var req model.NodeDefinition
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.bpmSvc.UpdateNode(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *BPMHandler) DeleteNode(c *gin.Context) {
	id := c.Param("id")
	if err := h.bpmSvc.DeleteNode(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ==================== 流程连线 ====================

func (h *BPMHandler) ListFlows(c *gin.Context) {
	modelID := c.Param("id")
	flows, err := h.bpmSvc.ListFlows(c.Request.Context(), modelID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": flows})
}

func (h *BPMHandler) CreateFlow(c *gin.Context) {
	modelID := c.Param("id")
	var req model.SequenceFlow
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var modelIDInt int64
	fmt.Sscanf(modelID, "%d", &modelIDInt)
	req.ModelID = modelIDInt
	req.TenantID = tenantID

	if err := h.bpmSvc.CreateFlow(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *BPMHandler) UpdateFlow(c *gin.Context) {
	id := c.Param("id")
	var req model.SequenceFlow
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.bpmSvc.UpdateFlow(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *BPMHandler) DeleteFlow(c *gin.Context) {
	id := c.Param("id")
	if err := h.bpmSvc.DeleteFlow(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ==================== 表单定义 ====================

func (h *BPMHandler) ListFormDefinitions(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	list, total, err := h.bpmSvc.ListFormDefinitions(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *BPMHandler) GetFormDefinition(c *gin.Context) {
	id := c.Param("id")
	form, err := h.bpmSvc.GetFormDefinition(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, form)
}

func (h *BPMHandler) CreateFormDefinition(c *gin.Context) {
	var req model.FormDefinition
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID

	if err := h.bpmSvc.CreateFormDefinition(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *BPMHandler) UpdateFormDefinition(c *gin.Context) {
	id := c.Param("id")
	var req model.FormDefinition
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.bpmSvc.UpdateFormDefinition(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *BPMHandler) DeleteFormDefinition(c *gin.Context) {
	id := c.Param("id")
	if err := h.bpmSvc.DeleteFormDefinition(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ==================== 表单字段 ====================

func (h *BPMHandler) ListFormFields(c *gin.Context) {
	formID := c.Param("id")
	fields, err := h.bpmSvc.ListFormFields(c.Request.Context(), formID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": fields})
}

func (h *BPMHandler) CreateFormField(c *gin.Context) {
	formID := c.Param("id")
	var req model.FormField
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var formIDInt int64
	fmt.Sscanf(formID, "%d", &formIDInt)
	req.FormID = formIDInt
	req.TenantID = tenantID

	if err := h.bpmSvc.CreateFormField(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *BPMHandler) UpdateFormField(c *gin.Context) {
	id := c.Param("id")
	var req model.FormField
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.bpmSvc.UpdateFormField(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *BPMHandler) DeleteFormField(c *gin.Context) {
	id := c.Param("id")
	if err := h.bpmSvc.DeleteFormField(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ==================== 流程实例 ====================

func (h *BPMHandler) ListProcessInstances(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}
	if initiatorID := c.Query("initiator_id"); initiatorID != "" {
		query["initiator_id"] = toInt64(initiatorID)
	}
	if page := c.Query("page"); page != "" {
		query["page"] = toInt(page)
	}
	if limit := c.Query("limit"); limit != "" {
		query["limit"] = toInt(limit)
	}

	list, total, err := h.bpmSvc.ListProcessInstances(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *BPMHandler) GetProcessInstance(c *gin.Context) {
	id := c.Param("id")
	instance, err := h.bpmSvc.GetProcessInstance(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, instance)
}

func (h *BPMHandler) CreateProcessInstance(c *gin.Context) {
	var req model.ProcessInstance
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID

	if err := h.bpmSvc.CreateProcessInstance(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *BPMHandler) CancelProcessInstance(c *gin.Context) {
	id := c.Param("id")
	if err := h.bpmSvc.CancelProcessInstance(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *BPMHandler) TerminateProcessInstance(c *gin.Context) {
	id := c.Param("id")
	if err := h.bpmSvc.TerminateProcessInstance(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ==================== 任务实例 ====================

func (h *BPMHandler) ListTasksByAssignee(c *gin.Context) {
	userID := middleware.GetUserID(c)

	query := map[string]interface{}{}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}
	if page := c.Query("page"); page != "" {
		query["page"] = toInt(page)
	}
	if limit := c.Query("limit"); limit != "" {
		query["limit"] = toInt(limit)
	}

	list, total, err := h.bpmSvc.ListTasksByAssignee(c.Request.Context(), userID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *BPMHandler) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := h.bpmSvc.GetTask(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, task)
}

func (h *BPMHandler) ApproveTask(c *gin.Context) {
	id := c.Param("id")
	userID := middleware.GetUserID(c)
	username := middleware.GetUsername(c)

	var req struct {
		Comment string `json:"comment"`
	}
	c.ShouldBindJSON(&req)

	if err := h.bpmSvc.ApproveTask(c.Request.Context(), id, userID, username, req.Comment); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *BPMHandler) RejectTask(c *gin.Context) {
	id := c.Param("id")
	userID := middleware.GetUserID(c)
	username := middleware.GetUsername(c)

	var req struct {
		Comment string `json:"comment"`
	}
	c.ShouldBindJSON(&req)

	if err := h.bpmSvc.RejectTask(c.Request.Context(), id, userID, username, req.Comment); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ==================== 委托记录 ====================

func (h *BPMHandler) ListDelegates(c *gin.Context) {
	userID := middleware.GetUserID(c)

	list, err := h.bpmSvc.ListDelegates(c.Request.Context(), userID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list})
}

func (h *BPMHandler) CreateDelegate(c *gin.Context) {
	var req model.DelegateRecord
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID

	if err := h.bpmSvc.CreateDelegate(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *BPMHandler) UpdateDelegate(c *gin.Context) {
	id := c.Param("id")
	var req model.DelegateRecord
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.bpmSvc.UpdateDelegate(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *BPMHandler) DeleteDelegate(c *gin.Context) {
	id := c.Param("id")
	if err := h.bpmSvc.DeleteDelegate(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ==================== 审批记录 ====================

func (h *BPMHandler) ListApprovalRecords(c *gin.Context) {
	taskID := c.Param("id")
	records, err := h.bpmSvc.ListApprovalRecords(c.Request.Context(), taskID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": records})
}
