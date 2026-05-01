package wms

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
)

// WMSOutboundHandler WMS出库管理处理器
type WMSOutboundHandler struct {
	svc *service.WarehouseService
}

// NewWMSOutboundHandler 创建WMS出库管理处理器
func NewWMSOutboundHandler(s *service.WarehouseService) *WMSOutboundHandler {
	return &WMSOutboundHandler{svc: s}
}

// List 获取出库单列表
// GET /wms/outbound/list
func (h *WMSOutboundHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "20")
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}

	list, total, err := h.svc.ListDeliveryOrder(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// Get 获取出库单详情
// GET /wms/outbound/:id
func (h *WMSOutboundHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	result, err := h.svc.GetDeliveryOrderWithItems(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, result)
}

// Create 创建出库单
// POST /wms/outbound
func (h *WMSOutboundHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var req service.DeliveryOrderCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req.TenantID = tenantID

	// 自动生成单号
	deliveryNo, err := h.svc.GenerateDeliveryNo(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	req.DeliveryNo = deliveryNo

	if err := h.svc.CreateDeliveryOrder(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Update 更新出库单
// PUT /wms/outbound/:id
func (h *WMSOutboundHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	var req service.DeliveryOrderCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.UpdateDeliveryOrder(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Delete 删除出库单
// DELETE /wms/outbound/:id
func (h *WMSOutboundHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	if err := h.svc.DeleteDeliveryOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Confirm 确认出库单
// POST /wms/outbound/:id/confirm
func (h *WMSOutboundHandler) Confirm(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	if err := h.svc.ConfirmDeliveryOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// OutboundRequest 出库请求结构
type OutboundRequest struct {
	DeliveryNo   string `json:"delivery_no"`
	CustomerID   int64  `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	WarehouseID  int64  `json:"warehouse_id"`
	Remark       string `json:"remark"`
}

// GetByID 根据ID获取出库单
// GET /wms/outbound/get/:id
func (h *WMSOutboundHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	deliveryOrder, err := h.svc.GetDeliveryOrderByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, deliveryOrder)
}

// ListByStatus 按状态获取出库单列表
// GET /wms/outbound/list-by-status
func (h *WMSOutboundHandler) ListByStatus(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	status := c.DefaultQuery("status", "")
	list, total, err := h.svc.ListDeliveryOrder(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	// 简单过滤状态（如果提供了status参数）
	if status != "" {
		statusInt, _ := strconv.Atoi(status)
		filtered := make([]model.DeliveryOrder, 0)
		for _, item := range list {
			if item.Status == statusInt {
				filtered = append(filtered, item)
			}
		}
		list = filtered
		total = int64(len(filtered))
	}

	response.Success(c, gin.H{"list": list, "total": total})
}

// BatchDelete 批量删除出库单
// DELETE /wms/outbound/batch
func (h *WMSOutboundHandler) BatchDelete(c *gin.Context) {
	var req struct {
		IDs []string `json:"ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	for _, id := range req.IDs {
		if err := h.svc.DeleteDeliveryOrder(c.Request.Context(), id); err != nil {
			response.ErrorMsg(c, err.Error())
			return
		}
	}
	response.Success(c, nil)
}

// GetDeliveryOrder 获取发货单详情（别名）
// GET /wms/outbound/delivery-order/:id
func (h *WMSOutboundHandler) GetDeliveryOrder(c *gin.Context) {
	h.Get(c)
}
