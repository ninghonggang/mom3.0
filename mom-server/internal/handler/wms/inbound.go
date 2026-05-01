package wms

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
)

// WMSInboundHandler WMS入库管理处理器
type WMSInboundHandler struct {
	svc *service.WarehouseService
}

// NewWMSInboundHandler 创建WMS入库管理处理器
func NewWMSInboundHandler(s *service.WarehouseService) *WMSInboundHandler {
	return &WMSInboundHandler{svc: s}
}

// List 获取入库单列表
// GET /wms/inbound/list
func (h *WMSInboundHandler) List(c *gin.Context) {
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

	list, total, err := h.svc.ListReceiveOrder(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// Get 获取入库单详情
// GET /wms/inbound/:id
func (h *WMSInboundHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	result, err := h.svc.GetReceiveOrderWithItems(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, result)
}

// Create 创建入库单
// POST /wms/inbound
func (h *WMSInboundHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var req service.ReceiveOrderCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req.TenantID = tenantID

	// 自动生成单号
	receiveNo, err := h.svc.GenerateReceiveNo(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	req.ReceiveNo = receiveNo

	if err := h.svc.CreateReceiveOrder(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Update 更新入库单
// PUT /wms/inbound/:id
func (h *WMSInboundHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	var req service.ReceiveOrderCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.UpdateReceiveOrder(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Delete 删除入库单
// DELETE /wms/inbound/:id
func (h *WMSInboundHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	if err := h.svc.DeleteReceiveOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Confirm 确认入库单
// POST /wms/inbound/:id/confirm
func (h *WMSInboundHandler) Confirm(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	if err := h.svc.ConfirmReceiveOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// InboundRequest 入库请求结构
type InboundRequest struct {
	ReceiveNo    string `json:"receive_no"`
	SupplierID   int64  `json:"supplier_id"`
	SupplierName string `json:"supplier_name"`
	WarehouseID  int64  `json:"warehouse_id"`
	Remark       string `json:"remark"`
}

// GetByID 根据ID获取入库单
// GET /wms/inbound/get/:id
func (h *WMSInboundHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	receiveOrder, err := h.svc.GetReceiveOrderByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, receiveOrder)
}

// ListByStatus 按状态获取入库单列表
// GET /wms/inbound/list-by-status
func (h *WMSInboundHandler) ListByStatus(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	status := c.DefaultQuery("status", "")
	list, total, err := h.svc.ListReceiveOrder(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	// 简单过滤状态（如果提供了status参数）
	if status != "" {
		statusInt, _ := strconv.Atoi(status)
		filtered := make([]model.ReceiveOrder, 0)
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

// BatchDelete 批量删除入库单
// DELETE /wms/inbound/batch
func (h *WMSInboundHandler) BatchDelete(c *gin.Context) {
	var req struct {
		IDs []string `json:"ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	for _, id := range req.IDs {
		if err := h.svc.DeleteReceiveOrder(c.Request.Context(), id); err != nil {
			response.ErrorMsg(c, err.Error())
			return
		}
	}
	response.Success(c, nil)
}

// GetReceiveOrder 获取收货单详情（别名）
// GET /wms/inbound/receive-order/:id
func (h *WMSInboundHandler) GetReceiveOrder(c *gin.Context) {
	h.Get(c)
}
