package mes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/service"
)

// CompleteInspectHandler 齐套检查处理器
type CompleteInspectHandler struct {
	completeInspectSvc *service.CompleteInspectService
}

func NewCompleteInspectHandler(completeInspectSvc *service.CompleteInspectService) *CompleteInspectHandler {
	return &CompleteInspectHandler{completeInspectSvc: completeInspectSvc}
}

// GetConfig GET /mes/complete-inspect/get - 获得齐套检查标识
func (h *CompleteInspectHandler) GetConfig(c *gin.Context) {
	paramCode := c.Query("paramCode")
	if paramCode == "" {
		paramCode = "PRE_START_CHECK_ON" // 默认配置码
	}

	config, err := h.completeInspectSvc.GetConfig(c.Request.Context(), paramCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": config})
}

// GetOrderDayBom POST /mes/complete-inspect/get-orderDay-bom - 获取日计划Bom信息
func (h *CompleteInspectHandler) GetOrderDayBom(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)

	var req model.MesWorkSchedulingBaseVO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list, err := h.completeInspectSvc.GetOrderDayBom(c.Request.Context(), tenantID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": list})
}

// GetOrderDayBomPage POST /mes/complete-inspect/get-orderDay-bom-page - 获取日计划Bom信息(分页)
func (h *CompleteInspectHandler) GetOrderDayBomPage(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)

	var req model.MesWorkSchedulingPageReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	list, total, err := h.completeInspectSvc.GetOrderDayBomPage(c.Request.Context(), tenantID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": list,
		"meta": gin.H{"total": total, "page": req.Page, "page_size": req.PageSize},
	})
}

// GetOrderDayWorkerPage POST /mes/complete-inspect/get-orderDay-worker-page - 获取日计划Worker信息(分页)
func (h *CompleteInspectHandler) GetOrderDayWorkerPage(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)

	var req model.MesWorkSchedulingPageReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	list, total, err := h.completeInspectSvc.GetOrderDayWorkerPage(c.Request.Context(), tenantID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": list,
		"meta": gin.H{"total": total, "page": req.Page, "page_size": req.PageSize},
	})
}

// GetOrderDayEquipmentPage POST /mes/complete-inspect/get-orderDay-equipment-page - 获取日计划Equipment信息(分页)
func (h *CompleteInspectHandler) GetOrderDayEquipmentPage(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)

	var req model.MesWorkSchedulingPageReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	list, total, err := h.completeInspectSvc.GetOrderDayEquipmentPage(c.Request.Context(), tenantID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": list,
		"meta": gin.H{"total": total, "page": req.Page, "page_size": req.PageSize},
	})
}

// GetOrderDayEquipment POST /mes/complete-inspect/get-orderDay-equipment - 获取日计划设备信息
func (h *CompleteInspectHandler) GetOrderDayEquipment(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)

	var req model.MesWorkSchedulingBaseVO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list, err := h.completeInspectSvc.GetOrderDayEquipment(c.Request.Context(), tenantID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": list})
}

// GetOrderDayWorker POST /mes/complete-inspect/get-orderDay-worker - 获取日计划人员信息
func (h *CompleteInspectHandler) GetOrderDayWorker(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)

	var req model.MesWorkSchedulingBaseVO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list, err := h.completeInspectSvc.GetOrderDayWorker(c.Request.Context(), tenantID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": list})
}

// Update POST /mes/complete-inspect/update - 更新生产日工单
func (h *CompleteInspectHandler) Update(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	username := middleware.GetUsername(c)

	var req model.CompleteInspectUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderDayID, err := strconv.ParseInt(c.Query("orderDayId"), 10, 64)
	if err != nil || orderDayID <= 0 {
		// 如果没有通过query传递，尝试从body获取
		orderDayID = req.OrderDayID
	}

	if orderDayID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order_day_id is required"})
		return
	}

	req.OrderDayID = orderDayID

	if err := h.completeInspectSvc.Update(c.Request.Context(), tenantID, &req, username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "更新成功"})
}
