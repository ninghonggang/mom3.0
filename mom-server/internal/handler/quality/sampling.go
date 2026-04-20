package quality

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/service"
)

// QMSSamplingHandler 抽样方案处理器
type QMSSamplingHandler struct {
	svc *service.QMSSamplingService
}

func NewQMSSamplingHandler(svc *service.QMSSamplingService) *QMSSamplingHandler {
	return &QMSSamplingHandler{svc: svc}
}

// CreatePlan POST /qms/sampling/plan/create
func (h *QMSSamplingHandler) CreatePlan(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	username := middleware.GetUsername(c)

	var req model.QMSSamplingPlanCreateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plan, err := h.svc.CreatePlan(c.Request.Context(), tenantID, &req, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": plan})
}

// UpdatePlan PUT /qms/sampling/plan/update
func (h *QMSSamplingHandler) UpdatePlan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	var req model.QMSSamplingPlanUpdateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.UpdatePlan(c.Request.Context(), id, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "更新成功"})
}

// DeletePlan DELETE /qms/sampling/plan/:id
func (h *QMSSamplingHandler) DeletePlan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	if err := h.svc.DeletePlan(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除成功"})
}

// ListPlan GET /qms/sampling/plan/list
func (h *QMSSamplingHandler) ListPlan(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	query := map[string]interface{}{
		"page":  1,
		"limit": 20,
		"status": c.Query("status"),
	}

	if page, err := strconv.Atoi(c.Query("page")); err == nil && page > 0 {
		query["page"] = page
	}
	if limit, err := strconv.Atoi(c.Query("limit")); err == nil && limit > 0 {
		query["limit"] = limit
	}

	list, total, err := h.svc.ListPlan(c.Request.Context(), tenantID, query)
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

// GetPlan GET /qms/sampling/plan/:id
func (h *QMSSamplingHandler) GetPlan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	plan, err := h.svc.GetPlan(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": plan})
}

// UpdateRules PUT /qms/sampling/plan/:id/rules
func (h *QMSSamplingHandler) UpdateRules(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	var req model.QMSSamplingRulesUpdateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.UpdateRules(c.Request.Context(), id, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "规则更新成功"})
}

// Calculate GET /qms/sampling/calculate
func (h *QMSSamplingHandler) Calculate(c *gin.Context) {
	planID, err := strconv.ParseInt(c.Query("planId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的planId"})
		return
	}

	batchQty, err := strconv.ParseFloat(c.Query("batchQty"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的batchQty"})
		return
	}

	result, err := h.svc.Calculate(c.Request.Context(), &model.QMSSamplingCalculateReqVO{
		PlanID:   planID,
		BatchQty: batchQty,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": result})
}

// CreateRecord POST /qms/sampling/record
func (h *QMSSamplingHandler) CreateRecord(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)

	var req model.QMSSamplingRecordCreateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.CreateRecord(c.Request.Context(), tenantID, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "抽样记录创建成功"})
}

// ListRecord GET /qms/sampling/record/list
func (h *QMSSamplingHandler) ListRecord(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	query := map[string]interface{}{
		"page": 1,
		"limit": 20,
	}

	if page, err := strconv.Atoi(c.Query("page")); err == nil && page > 0 {
		query["page"] = page
	}
	if limit, err := strconv.Atoi(c.Query("limit")); err == nil && limit > 0 {
		query["limit"] = limit
	}
	if planID, err := strconv.ParseInt(c.Query("planId"), 10, 64); err == nil && planID > 0 {
		query["plan_id"] = planID
	}
	if inspectionID, err := strconv.ParseInt(c.Query("inspectionId"), 10, 64); err == nil && inspectionID > 0 {
		query["inspection_id"] = inspectionID
	}

	list, total, err := h.svc.ListRecord(c.Request.Context(), tenantID, query)
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
