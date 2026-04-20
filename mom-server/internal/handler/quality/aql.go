package quality

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
)

type AQLHandler struct {
	svc *service.AQLService
}

func NewAQLHandler(svc *service.AQLService) *AQLHandler {
	return &AQLHandler{svc: svc}
}

func (h *AQLHandler) ListAQLLevels(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	list, total, err := h.svc.ListAQLLevels(c.Request.Context(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *AQLHandler) GetAQLLevel(c *gin.Context) {
	id := c.Param("id")
	item, err := h.svc.GetAQLLevel(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.Success(c, item)
}

func (h *AQLHandler) CreateAQLLevel(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	var req model.AQLLevel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.TenantID = tenantID
	if err := h.svc.CreateAQLLevel(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.Success(c, req)
}

func (h *AQLHandler) UpdateAQLLevel(c *gin.Context) {
	id := c.Param("id")
	var req model.AQLLevel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.UpdateAQLLevel(c.Request.Context(), id, map[string]any{
		"level": req.Level, "name": req.Name, "type": req.Type,
		"order": req.Order, "status": req.Status, "remark": req.Remark,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.Success(c, nil)
}

func (h *AQLHandler) DeleteAQLLevel(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.DeleteAQLLevel(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.Success(c, nil)
}

func (h *AQLHandler) ListAQLTableRows(c *gin.Context) {
	levelID := c.Query("level_id")
	if levelID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "level_id is required"})
		return
	}
	list, err := h.svc.ListAQLTableRows(c.Request.Context(), 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.Success(c, gin.H{"list": list})
}

func (h *AQLHandler) CreateAQLTableRow(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	var req model.AQLTableRow
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.TenantID = tenantID
	if err := h.svc.CreateAQLTableRow(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.Success(c, req)
}

func (h *AQLHandler) CalculateSampleSize(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	batchSize := 0
	fmt.Sscanf(c.Query("batch_size"), "%d", &batchSize)
	aqlValue := c.Query("aql")
	if batchSize <= 0 || aqlValue == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "batch_size and aql are required"})
		return
	}
	result, err := h.svc.CalculateSampleSize(c.Request.Context(), tenantID, batchSize, aqlValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.Success(c, result)
}

func (h *AQLHandler) ListSamplingPlans(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	query := c.Query("query")
	list, total, err := h.svc.ListSamplingPlans(c.Request.Context(), tenantID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *AQLHandler) GetSamplingPlan(c *gin.Context) {
	id := c.Param("id")
	item, err := h.svc.GetSamplingPlan(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.Success(c, item)
}

func (h *AQLHandler) CreateSamplingPlan(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	var req model.SamplingPlan
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.TenantID = tenantID
	if err := h.svc.CreateSamplingPlan(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.Success(c, req)
}

func (h *AQLHandler) UpdateSamplingPlan(c *gin.Context) {
	id := c.Param("id")
	var req model.SamplingPlan
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.UpdateSamplingPlan(c.Request.Context(), id, map[string]any{
		"name": req.Name, "inspection_type": req.InspectionType,
		"aql_level_id": req.AQLLevelID, "default_aql": req.DefaultAQL,
		"min_batch_size": req.MinBatchSize, "max_batch_size": req.MaxBatchSize,
		"status": req.Status, "remark": req.Remark,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.Success(c, nil)
}

func (h *AQLHandler) DeleteSamplingPlan(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.DeleteSamplingPlan(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.Success(c, nil)
}
