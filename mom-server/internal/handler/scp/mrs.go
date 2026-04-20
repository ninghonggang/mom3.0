package scp

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type MRSHandler struct {
	mrsService *service.ScpMRSService
}

func NewMRSHandler(s *service.ScpMRSService) *MRSHandler {
	return &MRSHandler{mrsService: s}
}

// List 查询MRS列表
func (h *MRSHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}
	if planMonth := c.Query("planMonth"); planMonth != "" {
		query["plan_month"] = planMonth
	}
	if sourceType := c.Query("sourceType"); sourceType != "" {
		query["source_type"] = sourceType
	}

	list, total, err := h.mrsService.ListMRS(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

// Get 获取MRS详情
func (h *MRSHandler) Get(c *gin.Context) {
	id := c.Param("id")
	mrs, err := h.mrsService.GetMRS(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, mrs)
}

// Create 创建MRS
func (h *MRSHandler) Create(c *gin.Context) {
	var req model.ScpMRSCreateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	mrs, err := h.mrsService.CreateMRS(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, mrs)
}

// Update 更新MRS
func (h *MRSHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.ScpMRSUpdateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	mrs, err := h.mrsService.UpdateMRS(c.Request.Context(), tenantID, id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, mrs)
}

// Delete 删除MRS
func (h *MRSHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.mrsService.DeleteMRS(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Publish 发布MRS
func (h *MRSHandler) Publish(c *gin.Context) {
	id := c.Param("id")
	userID := middleware.GetUserID(c)
	if err := h.mrsService.PublishMRS(c.Request.Context(), id, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Close 关闭MRS
func (h *MRSHandler) Close(c *gin.Context) {
	id := c.Param("id")
	if err := h.mrsService.CloseMRS(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetItems 获取MRS明细
func (h *MRSHandler) GetItems(c *gin.Context) {
	id := c.Param("id")
	items, err := h.mrsService.GetItems(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list": items,
	})
}

// Sync 从外部系统同步MRS
func (h *MRSHandler) Sync(c *gin.Context) {
	var req model.ScpMRSSyncReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	mrs, err := h.mrsService.SyncMRS(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, mrs)
}
