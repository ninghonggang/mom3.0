package wms

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

// WmsStrategyHandler 策略配置处理器
type WmsStrategyHandler struct {
	service *service.WmsStrategyService
}

// NewWmsStrategyHandler 创建策略配置处理器实例
func NewWmsStrategyHandler(s *service.WmsStrategyService) *WmsStrategyHandler {
	return &WmsStrategyHandler{service: s}
}

// List 获取策略配置列表
func (h *WmsStrategyHandler) List(c *gin.Context) {
	var query model.WmsStrategyQueryVO
	if err := c.ShouldBindQuery(&query); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	list, total, err := h.service.List(c.Request.Context(), &query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

// Get 获取策略配置详情
func (h *WmsStrategyHandler) Get(c *gin.Context) {
	id := c.Param("id")
	strategy, err := h.service.Get(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, strategy)
}

// Create 创建策略配置
func (h *WmsStrategyHandler) Create(c *gin.Context) {
	var req model.WmsStrategyCreateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// 设置租户ID
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Update 更新策略配置
func (h *WmsStrategyHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.WmsStrategyUpdateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req.Id = 0 // 防止通过body传入id
	if err := h.service.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Delete 删除策略配置
func (h *WmsStrategyHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
