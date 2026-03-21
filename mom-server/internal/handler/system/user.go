package system

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"mom-server/internal/dto"
	"mom-server/internal/middleware"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
)

// UserHandler 用户管理处理器
type UserHandler struct {
	userSvc *service.UserService
}

func NewUserHandler(userSvc *service.UserService) *UserHandler {
	return &UserHandler{userSvc: userSvc}
}

// GetList 获取用户列表
// @Summary 获取用户列表
// @Tags 系统管理-用户管理
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Param username query string false "用户名"
// @Param status query string false "状态"
// @Success 200 {object} response.Response
// @Router /api/v1/system/user/list [get]
func (h *UserHandler) GetList(c *gin.Context) {
	var req dto.PageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ParamError(c, "参数错误")
		return
	}

	username := c.Query("username")
	status := c.Query("status")

	data, err := h.userSvc.GetList(c.Request.Context(), req, username, status)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.PageSuccess(c, data.List, data.Total, data.Page, data.PageSize)
}

// GetByID 获取用户详情
// @Summary 获取用户详情
// @Tags 系统管理-用户管理
// @Success 200 {object} response.Response
// @Router /api/v1/system/user/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	var userID int64
	_, err := fmt.Sscanf(id, "%d", &userID)
	if err != nil {
		response.ParamError(c, "ID格式错误")
		return
	}

	user, err := h.userSvc.GetByID(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, 40401, err.Error())
		return
	}

	response.Success(c, user)
}

// Create 创建用户
// @Summary 创建用户
// @Tags 系统管理-用户管理
// @Accept json
// @Produce json
// @Param request body dto.CreateUserRequest true "用户信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/user [post]
func (h *UserHandler) Create(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, "参数错误: "+err.Error())
		return
	}

	// 设置默认密码
	if req.Password == "" {
		req.Password = "123456"
	}

	err := h.userSvc.Create(c.Request.Context(), req)
	if err != nil {
		response.Error(c, 40001, err.Error())
		return
	}

	response.Success(c, nil)
}

// Update 更新用户
// @Summary 更新用户
// @Tags 系统管理-用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param request body dto.UpdateUserRequest true "用户信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/user/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var userID int64
	_, err := fmt.Sscanf(id, "%d", &userID)
	if err != nil {
		response.ParamError(c, "ID格式错误")
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, "参数错误: "+err.Error())
		return
	}

	err = h.userSvc.Update(c.Request.Context(), userID, req)
	if err != nil {
		response.Error(c, 40001, err.Error())
		return
	}

	response.Success(c, nil)
}

// Delete 删除用户
// @Summary 删除用户
// @Tags 系统管理-用户管理
// @Success 200 {object} response.Response
// @Router /api/v1/system/user/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	var userID int64
	_, err := fmt.Sscanf(id, "%d", &userID)
	if err != nil {
		response.ParamError(c, "ID格式错误")
		return
	}

	err = h.userSvc.Delete(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, 40001, err.Error())
		return
	}

	response.Success(c, nil)
}

// ResetPassword 重置密码
// @Summary 重置密码
// @Tags 系统管理-用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param request body dto.ResetPasswordRequest true "密码"
// @Success 200 {object} response.Response
// @Router /api/v1/system/user/{id}/password [put]
func (h *UserHandler) ResetPassword(c *gin.Context) {
	id := c.Param("id")
	var userID int64
	_, err := fmt.Sscanf(id, "%d", &userID)
	if err != nil {
		response.ParamError(c, "ID格式错误")
		return
	}

	var req dto.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, "参数错误")
		return
	}

	err = h.userSvc.ResetPassword(c.Request.Context(), userID, req.Password)
	if err != nil {
		response.Error(c, 40001, err.Error())
		return
	}

	response.Success(c, nil)
}

// GetCurrentUser 获取当前用户
func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	userID := middleware.GetUserID(c)
	user, err := h.userSvc.GetByID(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, 40401, err.Error())
		return
	}
	response.Success(c, user)
}
