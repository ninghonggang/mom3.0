package system

import (
	"mom-server/internal/dto"
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	roleService *service.RoleService
}

func NewRoleHandler(rs *service.RoleService) *RoleHandler {
	return &RoleHandler{roleService: rs}
}

func (h *RoleHandler) List(c *gin.Context) {
	var req dto.RoleListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	list, total, err := h.roleService.List(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *RoleHandler) Get(c *gin.Context) {
	id := c.Param("id")
	role, err := h.roleService.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, role)
}

func (h *RoleHandler) Create(c *gin.Context) {
	var req model.Role
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

	err := h.roleService.Create(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, req)
}

func (h *RoleHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.Role
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.roleService.Update(c.Request.Context(), id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, req)
}

func (h *RoleHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.roleService.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *RoleHandler) GetMenus(c *gin.Context) {
	id := c.Param("id")
	menus, err := h.roleService.GetMenus(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, menus)
}

func (h *RoleHandler) AssignMenus(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		MenuIDs []uint `json:"menu_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.roleService.AssignMenus(c.Request.Context(), id, req.MenuIDs)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *RoleHandler) GetPerms(c *gin.Context) {
	id := c.Param("id")
	perms, err := h.roleService.GetRolePerms(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, perms)
}

func (h *RoleHandler) AssignPerms(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Perms []string `json:"perms"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.roleService.AssignPerms(c.Request.Context(), id, req.Perms)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
