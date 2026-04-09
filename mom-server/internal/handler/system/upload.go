package system

import (
	"fmt"
	"mom-server/internal/pkg/response"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	uploadDir string
}

func NewUploadHandler(uploadDir string) *UploadHandler {
	// 确保上传目录存在
	os.MkdirAll(uploadDir, 0755)
	return &UploadHandler{uploadDir: uploadDir}
}

// Upload 上传文件
func (h *UploadHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.ErrorMsg(c, "请选择要上传的文件")
		return
	}

	// 检查文件大小 (最大 5MB)
	if file.Size > 5*1024*1024 {
		response.ErrorMsg(c, "文件大小不能超过5MB")
		return
	}

	// 生成文件名
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), randomString(8), ext)
	filepath := filepath.Join(h.uploadDir, filename)

	// 保存文件
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		response.ErrorMsg(c, "文件保存失败")
		return
	}

	// 返回访问URL
	url := fmt.Sprintf("/uploads/%s", filename)
	response.Success(c, gin.H{
		"url":  url,
		"name": file.Filename,
	})
}

// randomString 生成随机字符串
func randomString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[time.Now().UnixNano()%int64(len(chars))]
	}
	return string(result)
}

// ServeStatic 提供静态文件服务
func ServeStatic(engine *gin.Engine, uploadDir string) {
	engine.Static("/uploads", uploadDir)
}

// Health 健康检查
func (h *UploadHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
