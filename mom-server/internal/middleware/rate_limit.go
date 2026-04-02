package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"mom-server/internal/pkg/errno"
	"mom-server/internal/pkg/response"
)

// RateLimiter 限流器
type RateLimiter struct {
	requests map[string][]time.Time
	maxRequests int
	window time.Duration
	mu sync.RWMutex
}

// NewRateLimiter 创建限流器
// maxRequests: 时间窗口内允许的最大请求数
// window: 时间窗口
func NewRateLimiter(maxRequests int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		maxRequests: maxRequests,
		window: window,
	}
}

// Allow 检查是否允许请求
func (rl *RateLimiter) Allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-rl.window)

	// 清理过期的请求记录
	if records, exists := rl.requests[key]; exists {
		var validRecords []time.Time
		for _, record := range records {
			if record.After(cutoff) {
				validRecords = append(validRecords, record)
			}
		}
		rl.requests[key] = validRecords
	}

	// 检查请求数是否超过限制
	if len(rl.requests[key]) >= rl.maxRequests {
		return false
	}

	// 添加当前请求
	rl.requests[key] = append(rl.requests[key], now)
	return true
}

// RateLimit 限流中间件（按IP）
func RateLimit(maxRequests int, window time.Duration) gin.HandlerFunc {
	limiter := NewRateLimiter(maxRequests, window)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		
		if !limiter.Allow(ip) {
			response.Error(c, errno.ErrTooManyRequest, "请求过于频繁，请稍后再试")
			c.Abort()
			return
		}

		c.Next()
	}
}

// RateLimitByUser 限流中间件（按用户ID）
func RateLimitByUser(maxRequests int, window time.Duration) gin.HandlerFunc {
	limiter := NewRateLimiter(maxRequests, window)

	return func(c *gin.Context) {
		userID := GetUserID(c)
		if userID == 0 {
			// 未登录用户按IP限流
			ip := c.ClientIP()
			if !limiter.Allow(ip) {
				response.Error(c, errno.ErrTooManyRequest, "请求过于频繁，请稍后再试")
				c.Abort()
				return
			}
		} else {
			// 已登录用户按用户ID限流
			key := string(rune(userID))
			if !limiter.Allow(key) {
				response.Error(c, errno.ErrTooManyRequest, "请求过于频繁，请稍后再试")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

// LoginRateLimit 登录专用限流（更严格）
func LoginRateLimit() gin.HandlerFunc {
	// 每IP每分钟最多5次登录尝试
	return RateLimit(5, time.Minute)
}
