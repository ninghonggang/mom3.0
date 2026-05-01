package eam

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
)

// EAMFactoryHandler 工厂Handler
type EAMFactoryHandler struct {
	db *gorm.DB
}

// NewEAMFactoryHandler 创建工厂Handler
func NewEAMFactoryHandler(db *gorm.DB) *EAMFactoryHandler {
	return &EAMFactoryHandler{db: db}
}

// List 获取工厂列表
func (h *EAMFactoryHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var list []model.Factory
	var total int64

	query := h.db.WithContext(c.Request.Context()).Model(&model.Factory{}).Where("tenant_id = ?", tenantID)

	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)
	query.Order("created_at DESC").Find(&list)

	response.Success(c, gin.H{"list": list, "total": total})
}

// Get 获取工厂详情
func (h *EAMFactoryHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var factory model.Factory
	if err := h.db.WithContext(c.Request.Context()).First(&factory, id).Error; err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, factory)
}

// Create 创建工厂
func (h *EAMFactoryHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var req model.Factory
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	req.TenantID = tenantID
	if err := h.db.WithContext(c.Request.Context()).Create(&req).Error; err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, req)
}

// Update 更新工厂
func (h *EAMFactoryHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	if err := h.db.WithContext(c.Request.Context()).Model(&model.Factory{}).Where("id = ?", id).Updates(req).Error; err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// Delete 删除工厂
func (h *EAMFactoryHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.db.WithContext(c.Request.Context()).Delete(&model.Factory{}, id).Error; err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}
