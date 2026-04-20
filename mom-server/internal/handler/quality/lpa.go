package quality

import (
	"strconv"

	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

// LPAHandler LPA分层审核处理器
type LPAHandler struct {
	standardSvc  *service.LPAStandardService
	questionSvc  *service.LPAQuestionService
	recordSvc    *service.LPARecordService
	itemSvc      *service.LPARecordItemService
}

func NewLPAHandler(standardSvc *service.LPAStandardService, questionSvc *service.LPAQuestionService, recordSvc *service.LPARecordService, itemSvc *service.LPARecordItemService) *LPAHandler {
	return &LPAHandler{standardSvc: standardSvc, questionSvc: questionSvc, recordSvc: recordSvc, itemSvc: itemSvc}
}

// ListStandards 查询审核标准列表
func (h *LPAHandler) ListStandards(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	offset := (page - 1) * pageSize

	filters := map[string]interface{}{}
	if isActive := c.Query("is_active"); isActive != "" {
		filters["is_active"] = isActive
	}

	list, total, err := h.standardSvc.List(c.Request.Context(), offset, pageSize, filters)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// GetStandard 获取审核标准详情
func (h *LPAHandler) GetStandard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	std, err := h.standardSvc.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, std)
}

// CreateStandard 创建审核标准
func (h *LPAHandler) CreateStandard(c *gin.Context) {
	var req model.LPAStandard
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.standardSvc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// UpdateStandard 更新审核标准
func (h *LPAHandler) UpdateStandard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.LPAStandard
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	updates := map[string]interface{}{
		"standard_name":   req.StandardName,
		"version":         req.Version,
		"dept_id":         req.DeptID,
		"dept_name":       req.DeptName,
		"audit_frequency": req.AuditFrequency,
		"auditor_levels":  req.AuditorLevels,
		"passing_score":   req.PassingScore,
		"is_active":       req.IsActive,
		"effective_date":  req.EffectiveDate,
		"remark":          req.Remark,
	}

	if err := h.standardSvc.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeleteStandard 删除审核标准
func (h *LPAHandler) DeleteStandard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.standardSvc.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ListQuestions 获取标准问题项
func (h *LPAHandler) ListQuestions(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	list, err := h.questionSvc.ListByStandard(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list})
}

// AddQuestion 添加问题项
func (h *LPAHandler) AddQuestion(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.LPAQuestion
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.StandardID = int64(id)

	if err := h.questionSvc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// ListRecords 查询审核记录列表
func (h *LPAHandler) ListRecords(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	offset := (page - 1) * pageSize

	filters := map[string]interface{}{}
	if standardID := c.Query("standard_id"); standardID != "" {
		filters["standard_id"] = standardID
	}
	if auditorID := c.Query("auditor_id"); auditorID != "" {
		filters["auditor_id"] = auditorID
	}

	list, total, err := h.recordSvc.List(c.Request.Context(), offset, pageSize, filters)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// GetRecord 获取审核记录详情
func (h *LPAHandler) GetRecord(c *gin.Context) {
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

	items, _ := h.itemSvc.ListByRecord(c.Request.Context(), uint(id))

	response.Success(c, gin.H{
		"record": record,
		"items":  items,
	})
}

// CreateRecord 创建审核记录
func (h *LPAHandler) CreateRecord(c *gin.Context) {
	var req struct {
		Record   model.LPARecord       `json:"record"`
		Items    []model.LPARecordItem `json:"items"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.recordSvc.Create(c.Request.Context(), &req.Record); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	if len(req.Items) > 0 {
		for i := range req.Items {
			req.Items[i].RecordID = req.Record.ID
		}
		h.itemSvc.CreateBatch(c.Request.Context(), req.Items)
	}

	response.Success(c, req.Record)
}

// VerifyRecord 确认审核记录
func (h *LPAHandler) VerifyRecord(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	updates := map[string]interface{}{
		"status": "VERIFIED",
	}

	if err := h.recordSvc.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
