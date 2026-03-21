package system

import (
	"github.com/gin-gonic/gin"
	"mom-server/internal/dto"
	"mom-server/internal/middleware"
	"mom-server/internal/pkg/jwt"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
)

// AuthHandler 认证处理器
type AuthHandler struct {
	userSvc *service.UserService
	jwtSvc  *jwt.JWT
}

func NewAuthHandler(userSvc *service.UserService, jwtSvc *jwt.JWT) *AuthHandler {
	return &AuthHandler{userSvc: userSvc, jwtSvc: jwtSvc}
}

// Login 登录
// @Summary 用户登录
// @Tags 认证
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "登录请求"
// @Success 200 {object} response.Response
// @Router /api/v1/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, "参数错误: "+err.Error())
		return
	}

	resp, err := h.userSvc.Login(c.Request.Context(), req)
	if err != nil {
		response.Error(c, 40001, err.Error())
		return
	}

	response.Success(c, resp)
}

// Logout 退出登录
// @Summary 退出登录
// @Tags 认证
// @Success 200 {object} response.Response
// @Router /api/v1/auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	// TODO: 实现退出登录逻辑
	response.Success(c, nil)
}

// GetUserInfo 获取用户信息
// @Summary 获取当前用户信息
// @Tags 认证
// @Success 200 {object} response.Response
// @Router /api/v1/auth/info [get]
func (h *AuthHandler) GetUserInfo(c *gin.Context) {
	userID := middleware.GetUserID(c)

	user, err := h.userSvc.GetByID(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, 40401, err.Error())
		return
	}

	response.Success(c, user)
}

// ChangePassword 修改密码
// @Summary 修改密码
// @Tags 认证
// @Accept json
// @Produce json
// @Param request body dto.ChangePasswordRequest true "密码请求"
// @Success 200 {object} response.Response
// @Router /api/v1/auth/password [put]
func (h *AuthHandler) ChangePassword(c *gin.Context) {
	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, "参数错误: "+err.Error())
		return
	}

	userID := middleware.GetUserID(c)
	err := h.userSvc.ChangePassword(c.Request.Context(), userID, req.OldPassword, req.NewPassword)
	if err != nil {
		response.Error(c, 40001, err.Error())
		return
	}

	response.Success(c, nil)
}

// RefreshToken 刷新Token
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, "参数错误: " + err.Error())
		return
	}

	accessToken, refreshToken, err := h.jwtSvc.RefreshToken(req.RefreshToken)
	if err != nil {
		response.Error(c, 40101, "Token刷新失败")
		return
	}

	response.Success(c, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
