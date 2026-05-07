package mes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
)

// ProcessHandler 工艺路线处理器
type ProcessHandler struct {
	processSvc *service.MesProcessService
}

func NewProcessHandler(processSvc *service.MesProcessService) *ProcessHandler {
	return &ProcessHandler{processSvc: processSvc}
}

// List 获取工艺路线列表
func (h *ProcessHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	query := map[string]interface{}{
		"page":          1,
		"limit":         20,
		"status":        c.Query("status"),
		"process_code":  c.Query("process_code"),
		"process_name":  c.Query("process_name"),
	}

	if page, err := strconv.Atoi(c.Query("page")); err == nil && page > 0 {
		query["page"] = page
	}
	if limit, err := strconv.Atoi(c.Query("limit")); err == nil && limit > 0 {
		query["limit"] = limit
	}
	if materialID, err := strconv.ParseInt(c.Query("material_id"), 10, 64); err == nil && materialID > 0 {
		query["material_id"] = materialID
	}

	list, total, err := h.processSvc.List(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, gin.H{"list": list, "total": total, "page": query["page"], "page_size": query["limit"]})
}

// Get 获取工艺路线详情
func (h *ProcessHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	process, err := h.processSvc.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, process)
}

// GetByMaterial 获取产品的有效工艺路线
func (h *ProcessHandler) GetByMaterial(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	materialID, err := strconv.ParseInt(c.Query("material_id"), 10, 64)
	if err != nil || materialID <= 0 {
		response.BadRequest(c, "无效的物料ID")
		return
	}

	list, err := h.processSvc.GetByMaterialID(c.Request.Context(), tenantID, materialID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, gin.H{"list": list})
}

// Create 创建工艺路线
func (h *ProcessHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)

	var req model.MesProcessCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	process, err := h.processSvc.Create(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, process)
}

// Update 更新工艺路线
func (h *ProcessHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	var req model.MesProcessUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	process, err := h.processSvc.Update(c.Request.Context(), uint(id), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, process)
}

// Delete 删除工艺路线
func (h *ProcessHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	if err := h.processSvc.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "删除成功"})
}

// UpdateStatus 更新状态
func (h *ProcessHandler) UpdateStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.processSvc.UpdateStatus(c.Request.Context(), uint(id), req.Status); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "状态更新成功"})
}

// Copy 复制工艺路线
func (h *ProcessHandler) Copy(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	process, err := h.processSvc.Copy(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, process)
}

// Validate 验证工艺路线
func (h *ProcessHandler) Validate(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	if err := h.processSvc.ValidateProcess(c.Request.Context(), uint(id)); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "验证通过"})
}
