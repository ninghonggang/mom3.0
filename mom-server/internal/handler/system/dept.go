package system

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type DeptHandler struct {
	deptService *service.DeptService
}

func NewDeptHandler(ds *service.DeptService) *DeptHandler {
	return &DeptHandler{deptService: ds}
}

func (h *DeptHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	depts, err := h.deptService.List(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, depts)
}

func (h *DeptHandler) Tree(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	depts, err := h.deptService.Tree(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, depts)
}

func (h *DeptHandler) Get(c *gin.Context) {
	id := c.Param("id")
	dept, err := h.deptService.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, dept)
}

func (h *DeptHandler) Create(c *gin.Context) {
	var req model.Dept
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.deptService.Create(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, req)
}

func (h *DeptHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.Dept
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.deptService.Update(c.Request.Context(), id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, req)
}

func (h *DeptHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deptService.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
