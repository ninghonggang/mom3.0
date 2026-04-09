package system

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type ImportHandler struct {
	importService *service.ImportService
}

func NewImportHandler(importService *service.ImportService) *ImportHandler {
	return &ImportHandler{importService: importService}
}

// ImportMaterials 导入物料
func (h *ImportHandler) ImportMaterials(c *gin.Context) {
	// 获取租户ID
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	// 获取当前用户
	username := ""
	if user, exists := c.Get("username"); exists {
		username = user.(string)
	}

	// 解析表单
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.BadRequest(c, "请选择要导入的文件")
		return
	}
	defer file.Close()

	// 验证文件类型
	if !strings.HasSuffix(strings.ToLower(header.Filename), ".xlsx") && !strings.HasSuffix(strings.ToLower(header.Filename), ".xls") {
		response.BadRequest(c, "只支持 .xlsx 或 .xls 格式的Excel文件")
		return
	}

	// 保存上传的文件
	uploadDir := "./uploads/import"
	filePath, err := h.importService.SaveUploadedFile(file, header, uploadDir)
	if err != nil {
		response.ErrorMsg(c, fmt.Sprintf("保存文件失败: %v", err))
		return
	}

	// 创建导入任务
	task, err := h.importService.CreateImportTask(c.Request.Context(), tenantID, model.ImportTypeMaterial, header.Filename, filePath, username)
	if err != nil {
		response.ErrorMsg(c, fmt.Sprintf("创建导入任务失败: %v", err))
		return
	}

	// 重新打开文件用于解析
	file2, _, err := c.Request.FormFile("file")
	if err != nil {
		response.ErrorMsg(c, "重新读取文件失败")
		return
	}
	defer file2.Close()

	// 解析Excel
	rows, err := h.importService.ParseExcel(file2)
	if err != nil {
		response.ErrorMsg(c, fmt.Sprintf("解析Excel失败: %v", err))
		return
	}

	if len(rows) == 0 {
		response.BadRequest(c, "Excel文件中没有数据行")
		return
	}

	// 启动导入（异步）
	go func() {
		_ = h.importService.ImportMaterials(c.Request.Context(), uint(task.ID), rows, tenantID)
	}()

	response.Success(c, gin.H{
		"task_id":  task.ID,
		"task_no":  task.TaskNo,
		"total":    len(rows),
		"message":  "导入任务已创建，正在处理中...",
	})
}

// GetImportTask 获取导入任务状态
func (h *ImportHandler) GetImportTask(c *gin.Context) {
	id := c.Param("id")
	task, err := h.importService.GetImportTask(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, task)
}

// DownloadTemplate 下载导入模板
func (h *ImportHandler) DownloadTemplate(c *gin.Context) {
	f, err := h.importService.GenerateMaterialTemplate()
	if err != nil {
		response.ErrorMsg(c, fmt.Sprintf("生成模板失败: %v", err))
		return
	}

	// 设置响应头
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "物料导入模板.xlsx"))

	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write template"})
	}
}

// GetImportTaskResult 获取导入结果详情
func (h *ImportHandler) GetImportTaskResult(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid task id")
		return
	}

	task, err := h.importService.GetImportTask(c.Request.Context(), fmt.Sprintf("%d", taskID))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, gin.H{
		"task_id":     task.ID,
		"task_no":     task.TaskNo,
		"status":      task.Status,
		"total_rows":  task.TotalRows,
		"success_rows": task.SuccessRows,
		"fail_rows":   task.FailRows,
		"fail_data":   task.FailDataJSON,
		"created_at":  task.CreatedAt,
		"completed_at": task.CompletedAt,
	})
}

// UploadFile 上传文件（兼容前端）
func (h *ImportHandler) UploadFile(c *gin.Context) {
	// 获取租户ID
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	username := ""
	if user, exists := c.Get("username"); exists {
		username = user.(string)
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.BadRequest(c, "请选择要上传的文件")
		return
	}
	defer file.Close()

	// 验证文件类型
	if !strings.HasSuffix(strings.ToLower(header.Filename), ".xlsx") && !strings.HasSuffix(strings.ToLower(header.Filename), ".xls") {
		response.BadRequest(c, "只支持 .xlsx 或 .xls 格式的Excel文件")
		return
	}

	// 保存文件
	uploadDir := "./uploads/import"
	filePath, err := h.importService.SaveUploadedFile(file, header, uploadDir)
	if err != nil {
		response.ErrorMsg(c, fmt.Sprintf("保存文件失败: %v", err))
		return
	}

	// 创建导入任务
	task, err := h.importService.CreateImportTask(c.Request.Context(), tenantID, model.ImportTypeMaterial, header.Filename, filePath, username)
	if err != nil {
		response.ErrorMsg(c, fmt.Sprintf("创建导入任务失败: %v", err))
		return
	}

	response.Success(c, gin.H{
		"task_id":  task.ID,
		"task_no":  task.TaskNo,
		"file_name": header.Filename,
		"file_path": filePath,
	})
}

// DoImport 执行导入
func (h *ImportHandler) DoImport(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var req struct {
		TaskID uint `json:"task_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if req.TaskID == 0 {
		response.BadRequest(c, "task_id is required")
		return
	}

	// 重新打开文件进行解析
	task, err := h.importService.GetImportTask(c.Request.Context(), fmt.Sprintf("%d", req.TaskID))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	// 解析Excel
	// 这里需要重新读取文件，但文件已在UploadFile中读取过
	// 简化处理：在ImportMaterials中重新实现文件读取
	rows, err := h.importService.ParseExcelFromPath(task.FilePath)
	if err != nil {
		response.ErrorMsg(c, fmt.Sprintf("解析Excel失败: %v", err))
		return
	}

	// 执行导入
	go func() {
		_ = h.importService.ImportMaterials(c.Request.Context(), req.TaskID, rows, tenantID)
	}()

	response.Success(c, gin.H{
		"message": "导入任务已启动，正在处理中...",
		"task_id": req.TaskID,
	})
}

// ImportBOM 上传并解析BOM导入
func (h *ImportHandler) ImportBOM(c *gin.Context) {
	// 获取租户ID
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	// 获取当前用户
	username := ""
	if user, exists := c.Get("username"); exists {
		username = user.(string)
	}

	// 解析表单
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.BadRequest(c, "请选择要导入的文件")
		return
	}
	defer file.Close()

	// 验证文件类型
	if !strings.HasSuffix(strings.ToLower(header.Filename), ".xlsx") && !strings.HasSuffix(strings.ToLower(header.Filename), ".xls") {
		response.BadRequest(c, "只支持 .xlsx 或 .xls 格式的Excel文件")
		return
	}

	// 保存上传的文件
	uploadDir := "./uploads/import"
	filePath, err := h.importService.SaveUploadedFile(file, header, uploadDir)
	if err != nil {
		response.ErrorMsg(c, fmt.Sprintf("保存文件失败: %v", err))
		return
	}

	// 创建导入任务
	task, err := h.importService.CreateImportTask(c.Request.Context(), tenantID, model.ImportTypeBom, header.Filename, filePath, username)
	if err != nil {
		response.ErrorMsg(c, fmt.Sprintf("创建导入任务失败: %v", err))
		return
	}

	// 解析BOM Excel
	rows, err := h.importService.ParseBOMExcel(filePath)
	if err != nil {
		response.ErrorMsg(c, fmt.Sprintf("解析Excel失败: %v", err))
		return
	}

	if len(rows) == 0 {
		response.BadRequest(c, "Excel文件中没有数据行")
		return
	}

	// 启动导入（异步）
	go func() {
		_ = h.importService.ImportBOMs(c.Request.Context(), uint(task.ID), rows, tenantID)
	}()

	response.Success(c, gin.H{
		"task_id": task.ID,
		"task_no": task.TaskNo,
		"total":   len(rows),
		"message": "导入任务已创建，正在处理中...",
	})
}

// DownloadBOMTemplate 下载BOM导入模板
func (h *ImportHandler) DownloadBOMTemplate(c *gin.Context) {
	f, err := h.importService.GenerateBOMTemplate()
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment;filename=bom_template.xlsx")
	if err := f.Write(c.Writer); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
}
