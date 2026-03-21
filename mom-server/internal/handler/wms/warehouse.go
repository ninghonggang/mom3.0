package wms

import (
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
