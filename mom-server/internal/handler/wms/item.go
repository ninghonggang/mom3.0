package wms

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

// WMSItemHandler 货品管理处理器
type WMSItemHandler struct {
	service *service.WMSItemService
}

// NewWMSItemHandler 创建货品处理器实例
func NewWMSItemHandler(s *service.WMSItemService) *WMSItemHandler {
	return &WMSItemHandler{service: s}
}

// List 获取货品列表
func (h *WMSItemHandler) List(c *gin.Context) {
	var query model.WMSItemQueryVO
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

// Get 获取货品详情
func (h *WMSItemHandler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, item)
}

// Search 搜索货品
func (h *WMSItemHandler) Search(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		response.BadRequest(c, "keyword is required")
		return
	}

	list, err := h.service.Search(c.Request.Context(), keyword)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list": list,
	})
}

// Create 创建货品
func (h *WMSItemHandler) Create(c *gin.Context) {
	var req model.WMSItemCreateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// 设置租户ID
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.CategoryID = nil // 确保为nil如果未提供

	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Update 更新货品
func (h *WMSItemHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.WMSItemUpdateReqVO
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

// Delete 删除货品
func (h *WMSItemHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}