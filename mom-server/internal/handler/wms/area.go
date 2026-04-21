package wms

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

// WmsAreaHandler 库区处理器
type WmsAreaHandler struct {
	svc *service.WmsAreaService
}

// NewWmsAreaHandler 创建库区处理器
func NewWmsAreaHandler(svc *service.WmsAreaService) *WmsAreaHandler {
	return &WmsAreaHandler{svc: svc}
}

// Create 创建库区
// POST /wms/area/create
func (h *WmsAreaHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var req model.WmsAreaCreateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.Create(c.Request.Context(), tenantID, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// Update 更新库区
// PUT /wms/area/update
func (h *WmsAreaHandler) Update(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	var req model.WmsAreaUpdateReqVO
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

// Delete 删除库区
// DELETE /wms/area/delete
func (h *WmsAreaHandler) Delete(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// Get 获取库区
// GET /wms/area/get
func (h *WmsAreaHandler) Get(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	code := c.Query("areaCode")
	if code == "" {
		response.BadRequest(c, "库区编码不能为空")
		return
	}

	area, err := h.svc.GetByCode(c.Request.Context(), tenantID, code)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, area)
}

// Page 分页查询库区
// GET /wms/area/page
func (h *WmsAreaHandler) Page(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var query model.WmsAreaQueryVO
	if err := c.ShouldBindQuery(&query); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	list, total, err := h.svc.Page(c.Request.Context(), tenantID, &query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, gin.H{"list": list, "total": total})
}

// Tree 获取库区树形结构
// GET /wms/area/tree
func (h *WmsAreaHandler) Tree(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	tree, err := h.svc.GetTree(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, gin.H{"list": tree, "total": len(tree)})
}

// ListByWarehouse 按仓库获取库区列表
// GET /wms/area/listByWarehouse
func (h *WmsAreaHandler) ListByWarehouse(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	warehouseCode := c.Query("warehouseCode")
	if warehouseCode == "" {
		response.BadRequest(c, "仓库编码不能为空")
		return
	}

	list, err := h.svc.ListByWarehouse(c.Request.Context(), tenantID, warehouseCode)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, gin.H{"list": list, "total": len(list)})
}
