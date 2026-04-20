package wms

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

// WmsLabelTemplateHandler 标签模板处理器
type WmsLabelTemplateHandler struct {
	service *service.WmsLabelTemplateService
}

// NewWmsLabelTemplateHandler 创建标签模板处理器实例
func NewWmsLabelTemplateHandler(s *service.WmsLabelTemplateService) *WmsLabelTemplateHandler {
	return &WmsLabelTemplateHandler{service: s}
}

// List 获取标签模板列表
func (h *WmsLabelTemplateHandler) List(c *gin.Context) {
	var query model.WmsLabelTemplateQueryVO
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

// Get 获取标签模板详情
func (h *WmsLabelTemplateHandler) Get(c *gin.Context) {
	id := c.Param("id")
	template, err := h.service.Get(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, template)
}

// Create 创建标签模板
func (h *WmsLabelTemplateHandler) Create(c *gin.Context) {
	var req model.WmsLabelTemplateCreateReqVO
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

// Update 更新标签模板
func (h *WmsLabelTemplateHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.WmsLabelTemplateUpdateReqVO
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

// Delete 删除标签模板
func (h *WmsLabelTemplateHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
