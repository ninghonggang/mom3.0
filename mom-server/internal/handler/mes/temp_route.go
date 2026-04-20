package mes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/service"
)

// TempRouteHandler MES日计划临时工艺路线处理器
type TempRouteHandler struct {
	tempRouteSvc *service.TempRouteService
}

func NewTempRouteHandler(tempRouteSvc *service.TempRouteService) *TempRouteHandler {
	return &TempRouteHandler{tempRouteSvc: tempRouteSvc}
}

// Create POST /mes/orderday/temp-route/create
func (h *TempRouteHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	username := middleware.GetUsername(c)

	var req model.TempRouteCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tempRoute, err := h.tempRouteSvc.Create(c.Request.Context(), tenantID, &req, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": tempRoute})
}

// Update PUT /mes/orderday/temp-route/update
func (h *TempRouteHandler) Update(c *gin.Context) {
	username := middleware.GetUsername(c)

	var req model.TempRouteUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	tempRoute, err := h.tempRouteSvc.Update(c.Request.Context(), id, &req, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": tempRoute})
}

// Delete DELETE /mes/orderday/temp-route/:id
func (h *TempRouteHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	if err := h.tempRouteSvc.Delete(c.Request.Context(), int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除成功"})
}

// ListByOrderDay GET /mes/orderday/temp-route/listByOrderDay?orderDayId=
func (h *TempRouteHandler) ListByOrderDay(c *gin.Context) {
	orderDayID, err := strconv.ParseInt(c.Query("orderDayId"), 10, 64)
	if err != nil || orderDayID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的日计划ID"})
		return
	}

	list, err := h.tempRouteSvc.ListByOrderDayID(c.Request.Context(), orderDayID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": list, "meta": gin.H{"total": len(list)}})
}

// Get GET /mes/orderday/temp-route/:id
func (h *TempRouteHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	tempRoute, err := h.tempRouteSvc.GetByID(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": tempRoute})
}

// Approve PUT /mes/orderday/temp-route/approve
func (h *TempRouteHandler) Approve(c *gin.Context) {
	username := middleware.GetUsername(c)

	var req model.TempRouteApprove
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.tempRouteSvc.Approve(c.Request.Context(), req.ID, req.Status, username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "审核成功"})
}
