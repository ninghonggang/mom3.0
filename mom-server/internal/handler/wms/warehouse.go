package wms

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type WarehouseHandler struct {
	service *service.WarehouseService
}

func NewWarehouseHandler(s *service.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{service: s}
}

func (h *WarehouseHandler) ListWarehouse(c *gin.Context) {
	list, total, err := h.service.ListWarehouse(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *WarehouseHandler) CreateWarehouse(c *gin.Context) {
	var req model.Warehouse
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	// 设置默认租户ID
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	err := h.service.CreateWarehouse(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *WarehouseHandler) UpdateWarehouse(c *gin.Context) {
	id := c.Param("id")
	var req model.Warehouse
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := h.service.UpdateWarehouse(c.Request.Context(), id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *WarehouseHandler) DeleteWarehouse(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteWarehouse(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *WarehouseHandler) ListLocation(c *gin.Context) {
	list, total, err := h.service.ListLocation(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *WarehouseHandler) GetLocation(c *gin.Context) {
	id := c.Param("id")
	location, err := h.service.GetLocationByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, location)
}

func (h *WarehouseHandler) CreateLocation(c *gin.Context) {
	var req model.Location
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	// 设置默认租户ID
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	err := h.service.CreateLocation(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *WarehouseHandler) UpdateLocation(c *gin.Context) {
	id := c.Param("id")
	var req model.Location
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := h.service.UpdateLocation(c.Request.Context(), id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *WarehouseHandler) DeleteLocation(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteLocation(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *WarehouseHandler) ListInventory(c *gin.Context) {
	list, total, err := h.service.ListInventory(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *WarehouseHandler) GetInventory(c *gin.Context) {
	id := c.Param("id")
	inventory, err := h.service.GetInventoryByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, inventory)
}

func (h *WarehouseHandler) CreateInventory(c *gin.Context) {
	var req model.Inventory
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	// 设置默认租户ID
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	err := h.service.CreateInventory(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *WarehouseHandler) UpdateInventory(c *gin.Context) {
	id := c.Param("id")
	var req model.Inventory
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := h.service.UpdateInventory(c.Request.Context(), id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *WarehouseHandler) DeleteInventory(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteInventory(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ========== 收货单 ==========

// ListReceiveOrder 获取收货单列表
func (h *WarehouseHandler) ListReceiveOrder(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	list, total, err := h.service.ListReceiveOrder(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

// GetReceiveOrder 获取收货单详情
func (h *WarehouseHandler) GetReceiveOrder(c *gin.Context) {
	id := c.Param("id")
	result, err := h.service.GetReceiveOrderWithItems(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, result)
}

// CreateReceiveOrder 创建收货单
func (h *WarehouseHandler) CreateReceiveOrder(c *gin.Context) {
	var req service.ReceiveOrderCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID

	// 自动生成单号
	receiveNo, err := h.service.GenerateReceiveNo(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	req.ReceiveNo = receiveNo

	if err := h.service.CreateReceiveOrder(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// UpdateReceiveOrder 更新收货单
func (h *WarehouseHandler) UpdateReceiveOrder(c *gin.Context) {
	id := c.Param("id")
	var req service.ReceiveOrderCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.UpdateReceiveOrder(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeleteReceiveOrder 删除收货单
func (h *WarehouseHandler) DeleteReceiveOrder(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteReceiveOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ConfirmReceiveOrder 确认收货单
func (h *WarehouseHandler) ConfirmReceiveOrder(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.ConfirmReceiveOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ========== 发货单 ==========

// ListDeliveryOrder 获取发货单列表
func (h *WarehouseHandler) ListDeliveryOrder(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	list, total, err := h.service.ListDeliveryOrder(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

// GetDeliveryOrder 获取发货单详情
func (h *WarehouseHandler) GetDeliveryOrder(c *gin.Context) {
	id := c.Param("id")
	result, err := h.service.GetDeliveryOrderWithItems(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, result)
}

// CreateDeliveryOrder 创建发货单
func (h *WarehouseHandler) CreateDeliveryOrder(c *gin.Context) {
	var req service.DeliveryOrderCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID

	// 自动生成单号
	deliveryNo, err := h.service.GenerateDeliveryNo(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	req.DeliveryNo = deliveryNo

	if err := h.service.CreateDeliveryOrder(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// UpdateDeliveryOrder 更新发货单
func (h *WarehouseHandler) UpdateDeliveryOrder(c *gin.Context) {
	id := c.Param("id")
	var req service.DeliveryOrderCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.UpdateDeliveryOrder(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeleteDeliveryOrder 删除发货单
func (h *WarehouseHandler) DeleteDeliveryOrder(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteDeliveryOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ConfirmDeliveryOrder 确认发货单
func (h *WarehouseHandler) ConfirmDeliveryOrder(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.ConfirmDeliveryOrder(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
