package mdm

import (
	"github.com/gin-gonic/gin"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
)

// ProductUnitHandler MDM计量单位处理器
type ProductUnitHandler struct {
	svc *service.MaterialService
}

// NewProductUnitHandler 创建MDM计量单位处理器
func NewProductUnitHandler(s *service.MaterialService) *ProductUnitHandler {
	return &ProductUnitHandler{svc: s}
}

// List 获取计量单位列表
// GET /mdm/product-unit/list
func (h *ProductUnitHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	list, total, err := h.svc.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// Get 获取计量单位详情
// GET /mdm/product-unit/:id
func (h *ProductUnitHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	material, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, material)
}

// Create 创建计量单位
// POST /mdm/product-unit
func (h *ProductUnitHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var req model.Material
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req.TenantID = tenantID

	if err := h.svc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// Update 更新计量单位
// PUT /mdm/product-unit/:id
func (h *ProductUnitHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	var req model.Material
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Delete 删除计量单位
// DELETE /mdm/product-unit/:id
func (h *ProductUnitHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ProductUnitRequest 计量单位请求结构
type ProductUnitRequest struct {
	MaterialCode string `json:"material_code"`
	MaterialName string `json:"material_name"`
	MaterialType string `json:"material_type"`
	Spec         string `json:"spec"`
	Unit         string `json:"unit"`
	Status       int    `json:"status"`
}

// GetByCode 根据编码获取计量单位
// GET /mdm/product-unit/code/:code
func (h *ProductUnitHandler) GetByCode(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		response.BadRequest(c, "code is required")
		return
	}

	// 暂未实现根据code查询，先返回空数据
	response.Success(c, nil)
}

// ListByType 按类型获取计量单位列表
// GET /mdm/product-unit/list-by-type
func (h *ProductUnitHandler) ListByType(c *gin.Context) {
	materialType := c.DefaultQuery("type", "")

	list, total, err := h.svc.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	// 如果提供了类型参数，进行过滤
	if materialType != "" {
		filtered := make([]model.Material, 0)
		for _, m := range list {
			if m.MaterialType == materialType {
				filtered = append(filtered, m)
			}
		}
		list = filtered
		total = int64(len(filtered))
	}

	response.Success(c, gin.H{"list": list, "total": total})
}

// BatchDelete 批量删除计量单位
// DELETE /mdm/product-unit/batch
func (h *ProductUnitHandler) BatchDelete(c *gin.Context) {
	var req struct {
		IDs []string `json:"ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	for _, id := range req.IDs {
		if err := h.svc.Delete(c.Request.Context(), id); err != nil {
			response.ErrorMsg(c, err.Error())
			return
		}
	}
	response.Success(c, nil)
}

// Enable 启用计量单位
// POST /mdm/product-unit/:id/enable
func (h *ProductUnitHandler) Enable(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	var req model.Material
	req.Status = 1 // 启用状态

	if err := h.svc.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Disable 禁用计量单位
// POST /mdm/product-unit/:id/disable
func (h *ProductUnitHandler) Disable(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	var req model.Material
	req.Status = 0 // 禁用状态

	if err := h.svc.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetUnit 获取计量单位信息（简化版）
// GET /mdm/product-unit/unit/:id
func (h *ProductUnitHandler) GetUnit(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	material, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	// 返回简化的单位信息
	response.Success(c, gin.H{
		"id":   material.ID,
		"unit": material.Unit,
		"name": material.MaterialName,
	})
}
