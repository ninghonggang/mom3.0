package system

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	menuService *service.MenuService
}

func NewMenuHandler(ms *service.MenuService) *MenuHandler {
	return &MenuHandler{menuService: ms}
}

func (h *MenuHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	menus, err := h.menuService.List(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, menus)
}

func (h *MenuHandler) Tree(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	menus, err := h.menuService.Tree(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, menus)
}

func (h *MenuHandler) Get(c *gin.Context) {
	id := c.Param("id")
	menu, err := h.menuService.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, menu)
}

func (h *MenuHandler) Create(c *gin.Context) {
	var req model.Menu
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

	err := h.menuService.Create(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, req)
}

func (h *MenuHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.Menu
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.menuService.Update(c.Request.Context(), id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, req)
}

func (h *MenuHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.menuService.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
