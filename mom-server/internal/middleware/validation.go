package middleware

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"mom-server/internal/pkg/errno"
	"mom-server/internal/pkg/response"
)

// ValidateIDParam ID参数验证中间件
func ValidateIDParam() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		if idStr == "" {
			response.Error(c, errno.ErrParamInvalid, "ID参数不能为空")
			c.Abort()
			return
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil || id <= 0 {
			response.Error(c, errno.ErrParamInvalid, "ID参数必须是有效的正整数")
			c.Abort()
			return
		}

		// 将解析后的ID存入上下文
		c.Set("validated_id", id)
		c.Next()
	}
}

// ValidatePageParams 分页参数验证中间件
func ValidatePageParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证 page
		pageStr := c.DefaultQuery("page", "1")
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			response.Error(c, errno.ErrParamInvalid, "page参数必须是大于等于1的整数")
			c.Abort()
			return
		}

		// 验证 page_size
		pageSizeStr := c.DefaultQuery("page_size", "10")
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil || pageSize < 1 || pageSize > 100 {
			response.Error(c, errno.ErrParamInvalid, "page_size参数必须是1-100之间的整数")
			c.Abort()
			return
		}

		c.Next()
	}
}

// GetValidatedID 从上下文获取验证后的ID
func GetValidatedID(c *gin.Context) int64 {
	if v, exists := c.Get("validated_id"); exists {
		if id, ok := v.(int64); ok {
			return id
		}
	}
	return 0
}

// RequiredParams 必填参数验证
type RequiredParams struct {
	Query  []string // 查询参数
	Body   []string // 请求体参数
	Headers []string // 请求头
}

// ValidateRequired 验证必填参数
func ValidateRequired(params RequiredParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证查询参数
		for _, key := range params.Query {
			if c.Query(key) == "" {
				response.Error(c, errno.ErrParamInvalid, fmt.Sprintf("查询参数 %s 不能为空", key))
				c.Abort()
				return
			}
		}

		// 验证请求头
		for _, key := range params.Headers {
			if c.GetHeader(key) == "" {
				response.Error(c, errno.ErrParamInvalid, fmt.Sprintf("请求头 %s 不能为空", key))
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
