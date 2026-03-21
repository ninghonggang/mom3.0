package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TenantMiddleware 租户中间件
type TenantMiddleware struct {
	db *gorm.DB
}

// NewTenantMiddleware 创建租户中间件
func NewTenantMiddleware(db *gorm.DB) *TenantMiddleware {
	return &TenantMiddleware{db: db}
}

// TenantID 租户ID
const TenantID = "tenant_id"

// TenantCheck 租户检查中间件
func (m *TenantMiddleware) TenantCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Token中获取租户ID（已经在JWT中间件中设置）
		tenantID := GetTenantID(c)
		if tenantID == 0 {
			tenantID = 1 // 默认租户
		}
		c.Set(TenantID, tenantID)
		c.Next()
	}
}

// GetTenantIDFromContext 从上下文获取租户ID
func GetTenantIDFromContext(c *gin.Context) int64 {
	if v, exists := c.Get(TenantID); exists {
		return v.(int64)
	}
	return 1
}
