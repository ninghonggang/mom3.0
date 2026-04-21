package mes

import (
	"os"
	"path/filepath"
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type SopHandler struct {
	sopService *service.MesSopService
}

func NewSopHandler(sopSvc *service.MesSopService) *SopHandler {
	return &SopHandler{sopService: sopSvc}
}

// Upload 上传SOP PDF文档
// POST /mes/sop/upload
func (h *SopHandler) Upload(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	// 获取上传文件
	file, err := c.FormFile("file")
	if err != nil {
		response.BadRequest(c, "请选择要上传的文件")
		return
	}

	// 验证文件类型
	if file.Header.Get("Content-Type") != "application/pdf" {
		ext := filepath.Ext(file.Filename)
		if ext != ".pdf" && ext != ".PDF" {
			response.BadRequest(c, "只支持PDF格式文件")
			return
		}
	}

	// 获取其他参数
	sopName := c.PostForm("sopName")
	if sopName == "" {
		response.BadRequest(c, "请填写SOP文档名称")
		return
	}

	processRouteIdStr := c.PostForm("processRouteId")
	processRouteId, _ := strconv.ParseInt(processRouteIdStr, 10, 64)

	workOrderIdStr := c.PostForm("workOrderId")
	workOrderId, _ := strconv.ParseInt(workOrderIdStr, 10, 64)

	version := c.PostForm("version")
	uploader := c.PostForm("uploader")

	// 构建上传请求
	req := &service.UploadReq{
		File:           file,
		SopName:        sopName,
		ProcessRouteId: processRouteId,
		WorkOrderId:    workOrderId,
		Version:        version,
		Uploader:       uploader,
	}

	sop, err := h.sopService.Upload(c.Request.Context(), tenantID, req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, gin.H{
		"id":       sop.ID,
		"sopName":  sop.SopName,
		"fileName": filepath.Base(sop.ContentURL),
		"fileUrl":  sop.ContentURL,
		"version":  sop.Version,
	})
}

// GetByWorkOrder 获取工单关联的SOP文档
// GET /mes/sop/getPDF?workOrderId=xxx
func (h *SopHandler) GetByWorkOrder(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	workOrderIdStr := c.Query("workOrderId")
	workOrderId, err := strconv.ParseInt(workOrderIdStr, 10, 64)
	if err != nil || workOrderId <= 0 {
		response.BadRequest(c, "无效的工单ID")
		return
	}

	list, err := h.sopService.GetByWorkOrder(c.Request.Context(), tenantID, workOrderId)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	// 转换为响应VO
	var result []service.SopRespVO
	for i := range list {
		result = append(result, h.sopService.ToRespVO(&list[i]))
	}

	response.Success(c, gin.H{"list": result, "total": len(result)})
}

// ListByProcessRoute 获取工艺路线关联的SOP列表
// GET /mes/sop/listByProcessRoute?processRouteId=xxx
func (h *SopHandler) ListByProcessRoute(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	processRouteIdStr := c.Query("processRouteId")
	processRouteId, err := strconv.ParseInt(processRouteIdStr, 10, 64)
	if err != nil || processRouteId <= 0 {
		response.BadRequest(c, "无效的工艺路线ID")
		return
	}

	list, err := h.sopService.GetByProcessRoute(c.Request.Context(), tenantID, processRouteId)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	// 转换为响应VO
	var result []service.SopRespVO
	for i := range list {
		result = append(result, h.sopService.ToRespVO(&list[i]))
	}

	response.Success(c, gin.H{"list": result, "total": len(result)})
}

// Get 获取SOP详情
// GET /mes/sop/:id
func (h *SopHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	sop, err := h.sopService.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, h.sopService.ToRespVO(sop))
}

// Delete 删除SOP文档
// DELETE /mes/sop/:id
func (h *SopHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	if err := h.sopService.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// Download 下载SOP文档
// GET /mes/sop/download/:id
func (h *SopHandler) Download(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	filePath, fileName, err := h.sopService.Download(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		response.ErrorMsg(c, "文件不存在")
		return
	}

	// 设置响应头
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/pdf")
	c.File(filePath)
}

// List 获取所有SOP列表
// GET /mes/sop/list
func (h *SopHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := c.Query("query")
	list, total, err := h.sopService.ListAll(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	// 转换为响应VO
	var result []service.SopRespVO
	for i := range list {
		result = append(result, h.sopService.ToRespVO(&list[i]))
	}

	response.Success(c, gin.H{"list": result, "total": total})
}
