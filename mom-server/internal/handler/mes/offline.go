package mes

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type OfflineHandler struct {
	offlineService *service.ProductionOfflineService
}

func NewOfflineHandler(offlineSvc *service.ProductionOfflineService) *OfflineHandler {
	return &OfflineHandler{offlineService: offlineSvc}
}

// List 获取离线记录列表
func (h *OfflineHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{
		"work_order_code": c.Query("work_order_code"),
		"product_code":    c.Query("product_code"),
		"product_name":    c.Query("product_name"),
		"offline_type":    c.Query("offline_type"),
		"handle_method":   c.Query("handle_method"),
		"status":          c.Query("status"),
		"start_date":      c.Query("start_date"),
		"end_date":        c.Query("end_date"),
	}

	list, total, err := h.offlineService.List(c.Request.Context(), uint64(tenantID), query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// Get 获取单个离线记录
func (h *OfflineHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	offline, err := h.offlineService.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, offline)
}

// Create 创建离线记录
func (h *OfflineHandler) Create(c *gin.Context) {
	var req model.ProductionOfflineCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	username := middleware.GetUsername(c)

	offline, err := h.offlineService.Create(c.Request.Context(), uint64(tenantID), &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, offline)
}

// Update 更新离线记录
func (h *OfflineHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.ProductionOfflineUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	username := middleware.GetUsername(c)

	if err := h.offlineService.Update(c.Request.Context(), id, &req, username); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Delete 删除离线记录
func (h *OfflineHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.offlineService.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Handle 处理离线记录（返工/降级/报废）
func (h *OfflineHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.ProductionOfflineHandleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	username := middleware.GetUsername(c)

	if err := h.offlineService.HandleOffline(c.Request.Context(), id, &req, username); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetItems 获取离线产品明细
func (h *OfflineHandler) GetItems(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	items, err := h.offlineService.GetItems(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": items, "total": len(items)})
}
