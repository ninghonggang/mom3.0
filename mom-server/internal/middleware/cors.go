package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// CORS CORS中间件 - 仅允许指定的来源
func CORS() gin.HandlerFunc {
	// 从环境变量获取允许的origins，多个用逗号分隔
	allowedOrigins := []string{
		"http://localhost:5173",
		"http://localhost:5174",
		"http://localhost:5175",
		"http://localhost:5176",
		"http://localhost:5177",
		"http://localhost:5178",
		"http://localhost:5179",
		"http://localhost:9080",
		"http://localhost:9081",
	}
	
	// 如果设置了环境变量，使用环境变量中的配置
	if envOrigins := os.Getenv("CORS_ALLOWED_ORIGINS"); envOrigins != "" {
		allowedOrigins = strings.Split(envOrigins, ",")
		// 清理空格
		for i := range allowedOrigins {
			allowedOrigins[i] = strings.TrimSpace(allowedOrigins[i])
		}
	}

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// 检查Origin是否在允许列表中
		allowed := false
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				c.Header("Access-Control-Allow-Origin", origin)
				allowed = true
				break
			}
		}

		// 如果是预检请求，直接返回
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Tenant-ID")
			c.Header("Access-Control-Max-Age", "86400")
			c.AbortWithStatus(204)
			return
		}

		// 如果不是允许的Origin，拒绝请求
		if !allowed && origin != "" {
			c.AbortWithStatus(403)
			return
		}

		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Tenant-ID")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		c.Next()
	}
}

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		_ = time.Since(start) // latency tracked but not logged
	}
}

// Recovery 恢复中间件
func Recovery() gin.HandlerFunc {
	return gin.Recovery()
}
