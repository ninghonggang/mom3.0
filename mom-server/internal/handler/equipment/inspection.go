package equipment

import (
	"strconv"

	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

// InspectionHandler 点检管理处理器
type InspectionHandler struct {
	templateSvc *service.InspectionTemplateService
	itemSvc     *service.InspectionItemService
	planSvc     *service.InspectionPlanService
	recordSvc   *service.InspectionRecordService
	resultSvc   *service.InspectionResultService
	defectSvc   *service.InspectionDefectService
}

func NewInspectionHandler(
	templateSvc *service.InspectionTemplateService,
	itemSvc *service.InspectionItemService,
	planSvc *service.InspectionPlanService,
	recordSvc *service.InspectionRecordService,
	resultSvc *service.InspectionResultService,
	defectSvc *service.InspectionDefectService,
) *InspectionHandler {
	return &InspectionHandler{
		templateSvc: templateSvc,
		itemSvc:     itemSvc,
		planSvc:     planSvc,
		recordSvc:   recordSvc,
		resultSvc:   resultSvc,
		defectSvc:   defectSvc,
	}
}

// ========== 点检标准 ==========

// ListTemplates 点检标准列表
func (h *InspectionHandler) ListTemplates(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	offset := (page - 1) * pageSize

	filters := map[string]interface{}{}
	if templateType := c.Query("template_type"); templateType != "" {
		filters["template_type"] = templateType
	}

	list, total, err := h.templateSvc.List(c.Request.Context(), offset, pageSize, filters)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// GetTemplate 获取点检标准详情
func (h *InspectionHandler) GetTemplate(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	tpl, err := h.templateSvc.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	items, _ := h.itemSvc.ListByTemplate(c.Request.Context(), uint(id))

	response.Success(c, gin.H{
		"template": tpl,
		"items":    items,
	})
}

// CreateTemplate 创建点检标准
func (h *InspectionHandler) CreateTemplate(c *gin.Context) {
	var req struct {
		Template model.InspectionTemplate `json:"template"`
		Items    []model.InspectionItem   `json:"items"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.templateSvc.Create(c.Request.Context(), &req.Template); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	if len(req.Items) > 0 {
		for i := range req.Items {
			req.Items[i].TemplateID = req.Template.ID
		}
		h.itemSvc.CreateBatch(c.Request.Context(), req.Items)
	}

	response.Success(c, req.Template)
}

// UpdateTemplate 更新点检标准
func (h *InspectionHandler) UpdateTemplate(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req struct {
		Template model.InspectionTemplate `json:"template"`
		Items    []model.InspectionItem   `json:"items"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	updates := map[string]interface{}{
		"template_name":        req.Template.TemplateName,
		"template_type":        req.Template.TemplateType,
		"equipment_type_id":   req.Template.EquipmentTypeID,
		"version":             req.Template.Version,
		"frequency_type":      req.Template.FrequencyType,
		"frequency_value":     req.Template.FrequencyValue,
		"execution_time":      req.Template.ExecutionTime,
		"estimated_minutes":  req.Template.EstimatedMinutes,
		"is_active":           req.Template.IsActive,
		"effective_date":      req.Template.EffectiveDate,
		"expiry_date":         req.Template.ExpiryDate,
		"remark":             req.Template.Remark,
	}

	if err := h.templateSvc.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// DeleteTemplate 删除点检标准
func (h *InspectionHandler) DeleteTemplate(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.templateSvc.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ========== 点检计划 ==========

// ListPlans 点检计划列表
func (h *InspectionHandler) ListPlans(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	offset := (page - 1) * pageSize

	filters := map[string]interface{}{}
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if planDate := c.Query("plan_date"); planDate != "" {
		filters["plan_date"] = planDate
	}

	list, total, err := h.planSvc.List(c.Request.Context(), offset, pageSize, filters)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// GetPlan 获取点检计划详情
func (h *InspectionHandler) GetPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	plan, err := h.planSvc.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, plan)
}

// CreatePlan 创建点检计划
func (h *InspectionHandler) CreatePlan(c *gin.Context) {
	var req model.InspectionPlan
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.planSvc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// UpdatePlan 更新点检计划
func (h *InspectionHandler) UpdatePlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.InspectionPlan
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	updates := map[string]interface{}{
		"plan_name":     req.PlanName,
		"equipment_id":  req.EquipmentID,
		"plan_date":     req.PlanDate,
		"plan_shift":    req.PlanShift,
		"assigned_to":   req.AssignedTo,
		"status":        req.Status,
	}

	if err := h.planSvc.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// AssignPlan 指派点检计划
func (h *InspectionHandler) AssignPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req struct {
		AssignedTo   int64  `json:"assigned_to"`
		AssignedName string `json:"assigned_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	updates := map[string]interface{}{
		"assigned_to":   req.AssignedTo,
		"assigned_name": req.AssignedName,
		"status":        "ASSIGNED",
	}

	if err := h.planSvc.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// CancelPlan 取消点检计划
func (h *InspectionHandler) CancelPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	updates := map[string]interface{}{
		"status": "CANCELLED",
	}

	if err := h.planSvc.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ========== 点检执行 ==========

// ListRecords 点检记录列表
func (h *InspectionHandler) ListRecords(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	offset := (page - 1) * pageSize

	filters := map[string]interface{}{}
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}

	list, total, err := h.recordSvc.List(c.Request.Context(), offset, pageSize, filters)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// GetRecord 获取点检记录详情
func (h *InspectionHandler) GetRecord(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	record, err := h.recordSvc.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	results, _ := h.resultSvc.ListByRecord(c.Request.Context(), uint(id))

	response.Success(c, gin.H{
		"record":  record,
		"results": results,
	})
}

// StartInspection 开始点检
func (h *InspectionHandler) StartInspection(c *gin.Context) {
	var req model.InspectionRecord
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.recordSvc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// CompleteInspection 完成点检
func (h *InspectionHandler) CompleteInspection(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req struct {
		Results []model.InspectionResult `json:"results"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if len(req.Results) > 0 {
		for i := range req.Results {
			req.Results[i].RecordID = int64(id)
		}
		h.resultSvc.CreateBatch(c.Request.Context(), req.Results)
	}

	// 计算结果统计
	var okCount, ngCount int
	for _, r := range req.Results {
		if r.ResultStatus == "OK" {
			okCount++
		} else if r.ResultStatus == "NG" {
			ngCount++
		}
	}

	overallResult := "OK"
	if ngCount > 0 {
		overallResult = "NG"
	} else if okCount == 0 && len(req.Results) > 0 {
		overallResult = "PARTIAL"
	}

	updates := map[string]interface{}{
		"status":         "COMPLETED",
		"overall_result": overallResult,
		"ok_count":       okCount,
		"ng_count":       ngCount,
	}

	if err := h.recordSvc.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ========== 异常处理 ==========

// ListDefects 异常列表
func (h *InspectionHandler) ListDefects(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	offset := (page - 1) * pageSize

	filters := map[string]interface{}{}
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}

	list, total, err := h.defectSvc.List(c.Request.Context(), offset, pageSize, filters)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// GetDefect 获取异常详情
func (h *InspectionHandler) GetDefect(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	defect, err := h.defectSvc.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, defect)
}

// CreateDefect 创建异常
func (h *InspectionHandler) CreateDefect(c *gin.Context) {
	var req model.InspectionDefect
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.defectSvc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// AssignDefect 指派异常处理人
func (h *InspectionHandler) AssignDefect(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req struct {
		AssignedTo   int64  `json:"assigned_to"`
		AssignedName string `json:"assigned_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	updates := map[string]interface{}{
		"assigned_to":    req.AssignedTo,
		"assigned_name":  req.AssignedName,
		"status":         "ACKNOWLEDGED",
	}

	if err := h.defectSvc.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ResolveDefect 处理完成
func (h *InspectionHandler) ResolveDefect(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req struct {
		Resolution string `json:"resolution"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	updates := map[string]interface{}{
		"resolution":    req.Resolution,
		"status":        "RESOLVED",
	}

	if err := h.defectSvc.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
