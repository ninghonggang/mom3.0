package quality

import (
	"strconv"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

// QRCIHandler QRCI质量闭环处理器
type QRCIHandler struct {
	svc           *service.QRCIService
	whySvc        *service.QRCI5WhyService
	actionSvc     *service.QRCIActionService
	verifySvc     *service.QRCIVerificationService
}

func NewQRCIHandler(svc *service.QRCIService, whySvc *service.QRCI5WhyService, actionSvc *service.QRCIActionService, verifySvc *service.QRCIVerificationService) *QRCIHandler {
	return &QRCIHandler{svc: svc, whySvc: whySvc, actionSvc: actionSvc, verifySvc: verifySvc}
}

// List 查询QRCI列表
func (h *QRCIHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	offset := (page - 1) * pageSize

	filters := map[string]interface{}{}
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if tenantID := c.Query("tenant_id"); tenantID != "" {
		filters["tenant_id"] = tenantID
	}

	list, total, err := h.svc.List(c.Request.Context(), offset, pageSize, filters)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// Get 获取QRCI详情
func (h *QRCIHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	qrci, err := h.svc.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, qrci)
}

// Create 创建QRCI
func (h *QRCIHandler) Create(c *gin.Context) {
	var req model.QRCI
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// Update 更新QRCI
func (h *QRCIHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.QRCI
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	updates := map[string]interface{}{
		"defect_description": req.DefectDescription,
		"severity_level":      req.SeverityLevel,
		"discovery_location":  req.DiscoveryLocation,
		"responsible_dept_id":  req.ResponsibleDeptID,
		"responsible_dept_name": req.ResponsibleDeptName,
		"owner_id":            req.OwnerID,
		"owner_name":          req.OwnerName,
		"target_close_date":   req.TargetCloseDate,
		"status":              req.Status,
		"remark":              req.Remark,
	}

	if err := h.svc.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Close 关闭QRCI
func (h *QRCIHandler) Close(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	now := time.Now()
	updates := map[string]interface{}{
		"status":           "CLOSED",
		"actual_close_date": &now,
	}

	if err := h.svc.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Delete 删除QRCI
func (h *QRCIHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.svc.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// List5Why 获取5Why分析列表
func (h *QRCIHandler) List5Why(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	list, err := h.whySvc.ListByQRCI(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list})
}

// Add5Why 添加5Why分析
func (h *QRCIHandler) Add5Why(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.QRCI5Why
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.QRCIID = int64(id)

	if err := h.whySvc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// ListActions 获取纠正措施列表
func (h *QRCIHandler) ListActions(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	list, err := h.actionSvc.ListByQRCI(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list})
}

// AddAction 添加纠正措施
func (h *QRCIHandler) AddAction(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.QRCIAction
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.QRCIID = int64(id)

	if err := h.actionSvc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// UpdateAction 更新纠正措施
func (h *QRCIHandler) UpdateAction(c *gin.Context) {
	actionID, err := strconv.ParseUint(c.Param("actionId"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid action id")
		return
	}

	var req model.QRCIAction
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	updates := map[string]interface{}{
		"action_description": req.ActionDescription,
		"responsible_id":      req.ResponsibleID,
		"responsible_name":   req.ResponsibleName,
		"due_date":           req.DueDate,
		"completed_date":      req.CompletedDate,
		"evidence_urls":       req.EvidenceURLs,
		"status":             req.Status,
		"remark":             req.Remark,
	}

	if err := h.actionSvc.Update(c.Request.Context(), uint(actionID), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// AddVerification 添加效果确认
func (h *QRCIHandler) AddVerification(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.QRCIVerification
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.QRCIID = int64(id)

	if err := h.verifySvc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}
