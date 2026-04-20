package mes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/service"
)

// MesHandler MES月计划/日计划处理器
type MesHandler struct {
	orderMonthSvc *service.OrderMonthService
	orderDaySvc   *service.OrderDayService
}

func NewMesHandler(orderMonthSvc *service.OrderMonthService, orderDaySvc *service.OrderDayService) *MesHandler {
	return &MesHandler{
		orderMonthSvc: orderMonthSvc,
		orderDaySvc:   orderDaySvc,
	}
}

// ========== 月计划 CRUD ==========

// ListMonthPlans GET /month-plans
func (h *MesHandler) ListMonthPlans(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	query := map[string]interface{}{
		"page":       1,
		"limit":      20,
		"plan_month": c.Query("plan_month"),
		"status":     c.Query("status"),
		"source_type": c.Query("source_type"),
	}

	if page, err := strconv.Atoi(c.Query("page")); err == nil && page > 0 {
		query["page"] = page
	}
	if limit, err := strconv.Atoi(c.Query("limit")); err == nil && limit > 0 {
		query["limit"] = limit
	}

	list, total, err := h.orderMonthSvc.List(c.Request.Context(), tenantID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": list,
		"meta": gin.H{"total": total, "page": query["page"], "limit": query["limit"]},
	})
}

// GetMonthPlan GET /month-plans/:id
func (h *MesHandler) GetMonthPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	order, err := h.orderMonthSvc.Get(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": order})
}

// CreateMonthPlan POST /month-plans
func (h *MesHandler) CreateMonthPlan(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	username := middleware.GetUsername(c)

	var req model.OrderMonthCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orderMonthSvc.Create(c.Request.Context(), tenantID, &req, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": order})
}

// UpdateMonthPlan PUT /month-plans/:id
func (h *MesHandler) UpdateMonthPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	username := middleware.GetUsername(c)

	var req model.OrderMonthUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orderMonthSvc.Update(c.Request.Context(), int64(id), &req, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": order})
}

// DeleteMonthPlan DELETE /month-plans/:id
func (h *MesHandler) DeleteMonthPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	if err := h.orderMonthSvc.Delete(c.Request.Context(), int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除成功"})
}

// ========== 月计划状态操作 ==========

// SubmitMonthPlan POST /month-plans/:id/submit
func (h *MesHandler) SubmitMonthPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	userID := middleware.GetUserID(c)
	username := middleware.GetUsername(c)

	if err := h.orderMonthSvc.Submit(c.Request.Context(), int64(id), userID, username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "提交成功"})
}

// ApproveMonthPlan POST /month-plans/:id/approve
func (h *MesHandler) ApproveMonthPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	userID := middleware.GetUserID(c)
	username := middleware.GetUsername(c)

	var req struct {
		Comment string `json:"comment"`
	}
	c.ShouldBindJSON(&req)

	if err := h.orderMonthSvc.Approve(c.Request.Context(), int64(id), userID, username, req.Comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "审核成功"})
}

// ReleaseMonthPlan POST /month-plans/:id/release
func (h *MesHandler) ReleaseMonthPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	userID := middleware.GetUserID(c)
	username := middleware.GetUsername(c)

	if err := h.orderMonthSvc.Release(c.Request.Context(), int64(id), userID, username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "下达成功"})
}

// CloseMonthPlan POST /month-plans/:id/close
func (h *MesHandler) CloseMonthPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	userID := middleware.GetUserID(c)
	username := middleware.GetUsername(c)

	if err := h.orderMonthSvc.Close(c.Request.Context(), int64(id), userID, username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "关闭成功"})
}

// CancelMonthPlan POST /month-plans/:id/cancel
func (h *MesHandler) CancelMonthPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	userID := middleware.GetUserID(c)
	username := middleware.GetUsername(c)

	var req struct {
		Comment string `json:"comment"`
	}
	c.ShouldBindJSON(&req)

	if err := h.orderMonthSvc.Cancel(c.Request.Context(), int64(id), userID, username, req.Comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "取消成功"})
}

// GetMonthPlanAudits GET /month-plans/:id/audits
func (h *MesHandler) GetMonthPlanAudits(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	audits, err := h.orderMonthSvc.GetAudits(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": audits})
}

// ========== 日计划 CRUD ==========

// ListDayPlans GET /day-plans
func (h *MesHandler) ListDayPlans(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	query := map[string]interface{}{
		"page":        1,
		"limit":       20,
		"status":      c.Query("status"),
		"plan_date":   c.Query("plan_date"),
		"start_date":  c.Query("start_date"),
		"end_date":    c.Query("end_date"),
		"line_id":     c.Query("line_id"),
		"month_plan_id": c.Query("month_plan_id"),
	}

	if page, err := strconv.Atoi(c.Query("page")); err == nil && page > 0 {
		query["page"] = page
	}
	if limit, err := strconv.Atoi(c.Query("limit")); err == nil && limit > 0 {
		query["limit"] = limit
	}
	if lineID, err := strconv.ParseInt(c.Query("line_id"), 10, 64); err == nil && lineID > 0 {
		query["line_id"] = lineID
	}
	if monthPlanID, err := strconv.ParseInt(c.Query("month_plan_id"), 10, 64); err == nil && monthPlanID > 0 {
		query["month_plan_id"] = monthPlanID
	}

	list, total, err := h.orderDaySvc.List(c.Request.Context(), tenantID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": list,
		"meta": gin.H{"total": total, "page": query["page"], "limit": query["limit"]},
	})
}

// GetDayPlan GET /day-plans/:id
func (h *MesHandler) GetDayPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	order, err := h.orderDaySvc.Get(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": order})
}

// CreateDayPlan POST /day-plans
func (h *MesHandler) CreateDayPlan(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	username := middleware.GetUsername(c)

	var req model.OrderDayCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orderDaySvc.Create(c.Request.Context(), tenantID, &req, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": order})
}

// UpdateDayPlan PUT /day-plans/:id
func (h *MesHandler) UpdateDayPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	username := middleware.GetUsername(c)

	var req model.OrderDayUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orderDaySvc.Update(c.Request.Context(), int64(id), &req, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": order})
}

// DeleteDayPlan DELETE /day-plans/:id
func (h *MesHandler) DeleteDayPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	if err := h.orderDaySvc.Delete(c.Request.Context(), int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除成功"})
}

// ========== 日计划状态操作 ==========

// PublishDayPlan POST /day-plans/:id/publish
func (h *MesHandler) PublishDayPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	userID := middleware.GetUserID(c)
	username := middleware.GetUsername(c)

	if err := h.orderDaySvc.Publish(c.Request.Context(), int64(id), userID, username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "发布成功"})
}

// CompleteDayPlan POST /day-plans/:id/complete
func (h *MesHandler) CompleteDayPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	if err := h.orderDaySvc.Complete(c.Request.Context(), int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "完成成功"})
}

// TerminateDayPlan POST /day-plans/:id/terminate
func (h *MesHandler) TerminateDayPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	userID := middleware.GetUserID(c)
	username := middleware.GetUsername(c)

	if err := h.orderDaySvc.Terminate(c.Request.Context(), int64(id), userID, username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "终止成功"})
}

// KitCheckDayPlan POST /day-plans/:id/kit-check
func (h *MesHandler) KitCheckDayPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	userID := middleware.GetUserID(c)

	if err := h.orderDaySvc.KitCheck(c.Request.Context(), int64(id), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "齐套检查完成"})
}

// DecomposeMonthPlan POST /month-plans/:id/decompose
func (h *MesHandler) DecomposeMonthPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	username := middleware.GetUsername(c)

	days, err := h.orderMonthSvc.Decompose(c.Request.Context(), int64(id), username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": days, "message": fmt.Sprintf("成功分解到 %d 个日计划", len(days))})
}

// GetDayPlansByMonth GET /month-plans/:id/day-plans
func (h *MesHandler) GetDayPlansByMonth(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	days, err := h.orderDaySvc.GetByMonthPlanID(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": days, "meta": gin.H{"total": len(days)}})
}
