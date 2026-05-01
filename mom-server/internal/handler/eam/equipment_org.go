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

// EAMEquipmentOrgHandler 设备组织Handler
type EAMEquipmentOrgHandler struct {
	db *gorm.DB
}

// NewEAMEquipmentOrgHandler 创建设备组织Handler
func NewEAMEquipmentOrgHandler(db *gorm.DB) *EAMEquipmentOrgHandler {
	return &EAMEquipmentOrgHandler{db: db}
}

// listRequest 查询参数
type eamEquipmentOrgListRequest struct {
	FactoryID  int64 `form:"factory_id"`
	WorkshopID int64 `form:"workshop_id"`
	Status     int   `form:"status"`
}

// List 获取设备组织列表
func (h *EAMEquipmentOrgHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var req eamEquipmentOrgListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	query := h.db.WithContext(c.Request.Context()).Model(&model.EquipmentOrg{}).Where("tenant_id = ?", tenantID)

	if req.FactoryID > 0 {
		query = query.Where("factory_id = ?", req.FactoryID)
	}
	if req.WorkshopID > 0 {
		query = query.Where("workshop_id = ?", req.WorkshopID)
	}
	if req.Status > 0 {
		query = query.Where("status = ?", req.Status)
	}

	var list []model.EquipmentOrg
	var total int64
	query.Count(&total)
	query.Order("factory_id, workshop_id, line_id").Find(&list)

	response.Success(c, gin.H{"list": list, "total": total})
}

// Get 获取设备组织详情
func (h *EAMEquipmentOrgHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var org model.EquipmentOrg
	if err := h.db.WithContext(c.Request.Context()).First(&org, id).Error; err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, org)
}

// Create 创建设备组织
func (h *EAMEquipmentOrgHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var req model.EquipmentOrg
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

// Update 更新设备组织
func (h *EAMEquipmentOrgHandler) Update(c *gin.Context) {
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

	if err := h.db.WithContext(c.Request.Context()).Model(&model.EquipmentOrg{}).Where("id = ?", id).Updates(req).Error; err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// Delete 删除设备组织
func (h *EAMEquipmentOrgHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.db.WithContext(c.Request.Context()).Delete(&model.EquipmentOrg{}, id).Error; err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}
