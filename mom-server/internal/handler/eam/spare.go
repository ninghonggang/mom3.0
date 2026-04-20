package eam

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/service"
)

// SpareHandler 备件Handler
type SpareHandler struct {
	svc *service.EquipmentSpareService
}

// NewSpareHandler 创建备件Handler
func NewSpareHandler(svc *service.EquipmentSpareService) *SpareHandler {
	return &SpareHandler{svc: svc}
}

// List 获取备件列表
// @Summary 获取备件列表
// @Tags EAM-备件管理
// @Param tenant_id header int false "租户ID"
// @Param keyword query string false "关键词"
// @Param category query string false "类别"
// @Param status query string false "状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.Response
// @Router /eam/spare/list [get]
func (h *SpareHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	query := make(map[string]interface{})
	if keyword := c.Query("keyword"); keyword != "" {
		query["keyword"] = keyword
	}
	if category := c.Query("category"); category != "" {
		query["category"] = category
	}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}

	spares, total, err := h.svc.ListPage(tenantID, query, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取备件列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "success",
		"data": gin.H{
			"list":  spares,
			"total": total,
			"page":  page,
		},
	})
}

// Get 获取备件详情
// @Summary 获取备件详情
// @Tags EAM-备件管理
// @Param id path int true "备件ID"
// @Success 200 {object} response.Response
// @Router /eam/spare/{id} [get]
func (h *SpareHandler) Get(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	spare, err := h.svc.GetByID(tenantID, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取备件失败"})
		return
	}
	if spare == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "备件不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "success",
		"data": spare,
	})
}

// Create 创建备件
// @Summary 创建备件
// @Tags EAM-备件管理
// @Param body body model.SpareCreateReq true "备件信息"
// @Success 200 {object} response.Response
// @Router /eam/spare [post]
func (h *SpareHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var req model.SpareCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	if err := h.svc.Create(tenantID, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建备件失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "success",
	})
}

// Update 更新备件
// @Summary 更新备件
// @Tags EAM-备件管理
// @Param body body model.SpareUpdateReq true "备件信息"
// @Success 200 {object} response.Response
// @Router /eam/spare [put]
func (h *SpareHandler) Update(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var req model.SpareUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	if err := h.svc.Update(tenantID, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新备件失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "success",
	})
}

// Delete 删除备件
// @Summary 删除备件
// @Tags EAM-备件管理
// @Param id path int true "备件ID"
// @Success 200 {object} response.Response
// @Router /eam/spare/{id} [delete]
func (h *SpareHandler) Delete(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	if err := h.svc.Delete(tenantID, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除备件失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "success",
	})
}

// Input 备件入库
// @Summary 备件入库
// @Tags EAM-备件管理
// @Param body body model.SpareTransactionReq true "入库信息"
// @Success 200 {object} response.Response
// @Router /eam/spare/input [post]
func (h *SpareHandler) Input(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var req model.SpareTransactionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	handlerID := uint(0)
	handlerName := ""

	if err := h.svc.In(tenantID, handlerID, handlerName, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "入库失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "success",
	})
}

// Output 备件出库
// @Summary 备件出库
// @Tags EAM-备件管理
// @Param body body model.SpareTransactionReq true "出库信息"
// @Success 200 {object} response.Response
// @Router /eam/spare/output [post]
func (h *SpareHandler) Output(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var req model.SpareTransactionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	handlerID := uint(0)
	handlerName := ""

	if err := h.svc.Out(tenantID, handlerID, handlerName, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "出库失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "success",
	})
}

// Transactions 获取事务记录
// @Summary 获取事务记录
// @Tags EAM-备件管理
// @Param spare_id query int true "备件ID"
// @Success 200 {object} response.Response
// @Router /eam/spare/transactions [get]
func (h *SpareHandler) Transactions(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	spareID, err := strconv.ParseUint(c.Query("spare_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的备件ID"})
		return
	}

	txs, err := h.svc.GetTransactions(tenantID, uint(spareID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取事务记录失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "success",
		"data": txs,
	})
}
