package mes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/service"
)

// TempBOMHandler MES日计划临时替代BOM处理器
type TempBOMHandler struct {
	tempBOMSvc *service.TempBOMService
}

func NewTempBOMHandler(tempBOMSvc *service.TempBOMService) *TempBOMHandler {
	return &TempBOMHandler{tempBOMSvc: tempBOMSvc}
}

// Create POST /mes/orderday/temp-bom/create
func (h *TempBOMHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	username := middleware.GetUsername(c)

	var req model.TempBOMCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tempBOM, err := h.tempBOMSvc.Create(c.Request.Context(), tenantID, &req, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": tempBOM})
}

// Update PUT /mes/orderday/temp-bom/update
func (h *TempBOMHandler) Update(c *gin.Context) {
	username := middleware.GetUsername(c)

	var req model.TempBOMUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	tempBOM, err := h.tempBOMSvc.Update(c.Request.Context(), id, &req, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": tempBOM})
}

// Delete DELETE /mes/orderday/temp-bom/:id
func (h *TempBOMHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	if err := h.tempBOMSvc.Delete(c.Request.Context(), int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除成功"})
}

// ListByOrderDayItem GET /mes/orderday/temp-bom/listByOrderDayItem?orderDayItemId=
func (h *TempBOMHandler) ListByOrderDayItem(c *gin.Context) {
	orderDayItemID, err := strconv.ParseInt(c.Query("orderDayItemId"), 10, 64)
	if err != nil || orderDayItemID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的日计划明细项ID"})
		return
	}

	list, err := h.tempBOMSvc.ListByOrderDayItemID(c.Request.Context(), orderDayItemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": list, "meta": gin.H{"total": len(list)}})
}

// Get GET /mes/orderday/temp-bom/:id
func (h *TempBOMHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	tempBOM, err := h.tempBOMSvc.GetByID(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": tempBOM})
}

// Approve PUT /mes/orderday/temp-bom/approve
func (h *TempBOMHandler) Approve(c *gin.Context) {
	username := middleware.GetUsername(c)

	var req model.TempBOMApprove
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.tempBOMSvc.Approve(c.Request.Context(), req.ID, req.Status, username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "审核成功"})
}
