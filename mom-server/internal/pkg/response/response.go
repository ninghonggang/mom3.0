package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 响应码定义
const (
	CodeSuccess       = 200
	CodeParamError    = 40001
	CodeUnauthorized  = 40101
	CodeForbidden     = 40301
	CodeNotFound      = 40401
	CodeInternalError = 50001
)

// Response 通用响应
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: "success",
		Data:    data,
	})
}

// SuccessWithMsg 自定义消息成功响应
func SuccessWithMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: msg,
		Data:    data,
	})
}

// PageSuccess 分页成功响应
func PageSuccess(c *gin.Context, list interface{}, total int64, page, pageSize int) {
	Success(c, gin.H{
		"list":      list,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// Error 错误响应
func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// ErrorMsg 简化错误响应（使用默认错误码）
func ErrorMsg(c *gin.Context, message string) {
	Error(c, CodeInternalError, message)
}

// BadRequest 请求参数错误
func BadRequest(c *gin.Context, message string) {
	ParamError(c, message)
}

// ParamError 参数错误
func ParamError(c *gin.Context, message string) {
	Error(c, CodeParamError, message)
}

// Unauthorized 未授权
func Unauthorized(c *gin.Context, message string) {
	if message == "" {
		message = "未授权，请登录"
	}
	c.JSON(http.StatusUnauthorized, Response{
		Code:    CodeUnauthorized,
		Message: message,
	})
}

// Forbidden 禁止访问
func Forbidden(c *gin.Context, message string) {
	if message == "" {
		message = "没有权限"
	}
	Error(c, CodeForbidden, message)
}

// NotFound 资源不存在
func NotFound(c *gin.Context, message string) {
	if message == "" {
		message = "资源不存在"
	}
	Error(c, CodeNotFound, message)
}

// InternalError 服务器内部错误
func InternalError(c *gin.Context, message string) {
	if message == "" {
		message = "服务器内部错误"
	}
	Error(c, CodeInternalError, message)
}
