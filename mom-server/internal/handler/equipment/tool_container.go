package equipment

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ToolHandler 器具处理器
type ToolHandler struct {
	service *service.ToolService
}

func NewToolHandler(s *service.ToolService) *ToolHandler {
	return &ToolHandler{service: s}
}

func (h *ToolHandler) List(c *gin.Context) {
	var params model.ToolQuery
	if err := c.ShouldBindQuery(&params); err != nil {
		params.Page = 1
		params.PageSize = 20
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	list, total, err := h.service.List(c.Request.Context(), tenantID, &params)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *ToolHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	tool, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, tool)
}

func (h *ToolHandler) Create(c *gin.Context) {
	var req model.ToolCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	tool := &model.Tool{
		TenantID:      tenantID,
		ToolCode:      req.ToolCode,
		ToolName:      req.ToolName,
		ToolType:      req.ToolType,
		Specification: req.Specification,
		Unit:          req.Unit,
		Status:        req.Status,
		Location:      req.Location,
		Remark:        req.Remark,
	}

	if err := h.service.Create(c.Request.Context(), tool); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, tool)
}

func (h *ToolHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var req model.ToolUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	updates := map[string]interface{}{}
	if req.ToolName != "" {
		updates["tool_name"] = req.ToolName
	}
	if req.ToolType != "" {
		updates["tool_type"] = req.ToolType
	}
	if req.Specification != "" {
		updates["specification"] = req.Specification
	}
	if req.Unit != "" {
		updates["unit"] = req.Unit
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}
	if req.Location != "" {
		updates["location"] = req.Location
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	if err := h.service.Update(c.Request.Context(), id, updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *ToolHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ToolContainerHandler 容器处理器
type ToolContainerHandler struct {
	service *service.ToolContainerService
}

func NewToolContainerHandler(s *service.ToolContainerService) *ToolContainerHandler {
	return &ToolContainerHandler{service: s}
}

func (h *ToolContainerHandler) List(c *gin.Context) {
	var params model.ToolContainerQuery
	if err := c.ShouldBindQuery(&params); err != nil {
		params.Page = 1
		params.PageSize = 20
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	list, total, err := h.service.List(c.Request.Context(), tenantID, &params)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *ToolContainerHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	container, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, container)
}

func (h *ToolContainerHandler) Create(c *gin.Context) {
	var req model.ToolContainerCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	container := &model.ToolContainer{
		TenantID:      tenantID,
		ContainerCode: req.ContainerCode,
		ContainerName: req.ContainerName,
		ContainerType: req.ContainerType,
		Capacity:      req.Capacity,
		Status:        req.Status,
		Location:      req.Location,
		Remark:        req.Remark,
	}

	if err := h.service.Create(c.Request.Context(), container); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, container)
}

func (h *ToolContainerHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var req model.ToolContainerUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	updates := map[string]interface{}{}
	if req.ContainerName != "" {
		updates["container_name"] = req.ContainerName
	}
	if req.ContainerType != "" {
		updates["container_type"] = req.ContainerType
	}
	if req.Capacity > 0 {
		updates["capacity"] = req.Capacity
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}
	if req.Location != "" {
		updates["location"] = req.Location
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	if err := h.service.Update(c.Request.Context(), id, updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *ToolContainerHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ToolContainerBindingHandler 器具容器绑定处理器
type ToolContainerBindingHandler struct {
	service *service.ToolContainerBindingService
}

func NewToolContainerBindingHandler(s *service.ToolContainerBindingService) *ToolContainerBindingHandler {
	return &ToolContainerBindingHandler{service: s}
}

func (h *ToolContainerBindingHandler) List(c *gin.Context) {
	var params model.ToolContainerBindingQuery
	if err := c.ShouldBindQuery(&params); err != nil {
		params.Page = 1
		params.PageSize = 20
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	list, total, err := h.service.List(c.Request.Context(), tenantID, &params)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *ToolContainerBindingHandler) Bind(c *gin.Context) {
	var req model.ToolContainerBindingCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	if err := h.service.Bind(c.Request.Context(), tenantID, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *ToolContainerBindingHandler) Unbind(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.service.Unbind(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *ToolContainerBindingHandler) GetToolBinding(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	binding, err := h.service.GetActiveBinding(c.Request.Context(), tenantID, id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, binding)
}

func (h *ToolContainerBindingHandler) GetContainerBindings(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	list, err := h.service.GetByContainerID(c.Request.Context(), tenantID, id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list})
}
