package production

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductionOrderHandler struct {
	service *service.ProductionOrderService
}

func NewProductionOrderHandler(s *service.ProductionOrderService) *ProductionOrderHandler {
	return &ProductionOrderHandler{service: s}
}

func (h *ProductionOrderHandler) List(c *gin.Context) {
	list, total, err := h.service.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *ProductionOrderHandler) Get(c *gin.Context) {
	id := c.Param("id")
	order, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, order)
}

func (h *ProductionOrderHandler) Create(c *gin.Context) {
	var req model.ProductionOrder
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

	err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, req)
}

func (h *ProductionOrderHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.ProductionOrder
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// 获取当前用户名
	username := middleware.GetUsername(c)

	err := h.service.Update(c.Request.Context(), id, &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, req)
}

func (h *ProductionOrderHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *ProductionOrderHandler) Start(c *gin.Context) {
	id := c.Param("id")
	username := middleware.GetUsername(c)
	err := h.service.Start(c.Request.Context(), id, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *ProductionOrderHandler) Complete(c *gin.Context) {
	id := c.Param("id")
	username := middleware.GetUsername(c)
	err := h.service.Complete(c.Request.Context(), id, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
