package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"mom-server/internal/pkg/jwt"
	"mom-server/internal/pkg/response"
)

const (
	ContextKeyUserID   = "user_id"
	ContextKeyTenantID = "tenant_id"
	ContextKeyUsername = "username"
	ContextKeyRoles    = "roles"
)

// JWTAuth JWT认证中间件
func JWTAuth(jwtUtil *jwt.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Header获取Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, "请求头中Authorization为空")
			c.Abort()
			return
		}

		// Bearer Token格式
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Unauthorized(c, "Token格式错误")
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 解析Token
		claims, err := jwtUtil.ParseToken(tokenString)
		if err != nil {
			response.Unauthorized(c, "Token解析失败: " + err.Error())
			c.Abort()
			return
		}

		// 设置上下文
		c.Set(ContextKeyUserID, claims.UserID)
		c.Set(ContextKeyTenantID, claims.TenantID)
		c.Set(ContextKeyUsername, claims.Username)
		c.Set(ContextKeyRoles, claims.Roles)

		c.Next()
	}
}

// GetUserID 获取用户ID
func GetUserID(c *gin.Context) int64 {
	if v, exists := c.Get(ContextKeyUserID); exists {
		if userID, ok := v.(int64); ok {
			return userID
		}
	}
	return 0
}

// GetTenantID 获取租户ID
func GetTenantID(c *gin.Context) int64 {
	if v, exists := c.Get(ContextKeyTenantID); exists {
		if tenantID, ok := v.(int64); ok {
			return tenantID
		}
	}
	return 1 // 默认返回租户1，避免0值问题
}

// IsSuperAdmin 判断是否为超级管理员
func IsSuperAdmin(c *gin.Context) bool {
	roles := GetRoles(c)
	for _, role := range roles {
		if role == "super_admin" || role == "admin" {
			return true
		}
	}
	return false
}

// GetUsername 获取用户名
func GetUsername(c *gin.Context) string {
	if v, exists := c.Get(ContextKeyUsername); exists {
		if username, ok := v.(string); ok {
			return username
		}
	}
	return ""
}

// GetRoles 获取角色列表
func GetRoles(c *gin.Context) []string {
	if v, exists := c.Get(ContextKeyRoles); exists {
		if roles, ok := v.([]string); ok {
			return roles
		}
	}
	return nil
}
