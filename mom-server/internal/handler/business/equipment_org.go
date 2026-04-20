package business

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
)

// EquipmentOrgHandler 设备组织层级Handler
type EquipmentOrgHandler struct {
	svc *service.EquipmentOrgService
}

// NewEquipmentOrgHandler 创建设备组织Handler
func NewEquipmentOrgHandler(svc *service.EquipmentOrgService) *EquipmentOrgHandler {
	return &EquipmentOrgHandler{svc: svc}
}

// listRequest 查询参数
type listRequest struct {
	FactoryID  int64 `form:"factory_id"`
	WorkshopID int64 `form:"workshop_id"`
	Status     int   `form:"status"`
}

// List 获取设备组织列表
// @Summary 获取设备组织列表
// @Tags EAM-设备组织
// @Param factory_id query int false "厂区ID"
// @Param workshop_id query int false "车间ID"
// @Param status query int false "状态"
// @Success 200 {object} response.Response{data=[]model.EquipmentOrg}
// @Router /eam/equipment-org/list [get]
func (h *EquipmentOrgHandler) List(c *gin.Context) {
	var req listRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	query := &model.EquipmentOrgQuery{
		FactoryID:  req.FactoryID,
		WorkshopID: req.WorkshopID,
		Status:     req.Status,
	}

	list, total, err := h.svc.List(c.Request.Context(), query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// Get 获取设备组织详情
// @Summary 获取设备组织详情
// @Tags EAM-设备组织
// @Param id path int true "ID"
// @Success 200 {object} response.Response{data=model.EquipmentOrg}
// @Router /eam/equipment-org/{id} [get]
func (h *EquipmentOrgHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	org, err := h.svc.GetByID(c.Request.Context(), idStr)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, org)
}

// SyncFromMasterData 从主数据同步设备组织关系
// @Summary 从主数据同步设备组织关系
// @Tags EAM-设备组织
// @Success 200 {object} response.Response
// @Router /eam/equipment-org/sync [post]
func (h *EquipmentOrgHandler) SyncFromMasterData(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	if err := h.svc.SyncFromMasterData(c.Request.Context(), tenantID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Update 更新设备组织
// @Summary 更新设备组织
// @Tags EAM-设备组织
// @Param id path int true "ID"
// @Param body body map[string]interface{} true "更新内容"
// @Success 200 {object} response.Response
// @Router /eam/equipment-org/{id} [put]
func (h *EquipmentOrgHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	if err := h.svc.Update(c.Request.Context(), idStr, req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Delete 删除设备组织
// @Summary 删除设备组织
// @Tags EAM-设备组织
// @Param id path int true "ID"
// @Success 200 {object} response.Response
// @Router /eam/equipment-org/{id} [delete]
func (h *EquipmentOrgHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	if err := h.svc.Delete(c.Request.Context(), idStr); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetFactoryList 获取厂区列表
// @Summary 获取厂区列表
// @Tags EAM-设备组织
// @Success 200 {object} response.Response{data=[]model.Factory}
// @Router /eam/factory/list [get]
func (h *EquipmentOrgHandler) GetFactoryList(c *gin.Context) {
	list, err := h.svc.GetFactoryList(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": len(list)})
}

// GetFactory 获取厂区详情
// @Summary 获取厂区详情
// @Tags EAM-设备组织
// @Param id path int true "厂区ID"
// @Success 200 {object} response.Response{data=model.Factory}
// @Router /eam/factory/{id} [get]
func (h *EquipmentOrgHandler) GetFactory(c *gin.Context) {
	idStr := c.Param("id")
	factory, err := h.svc.GetFactoryByID(c.Request.Context(), idStr)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, factory)
}

// CreateFactory 创建厂区
// @Summary 创建厂区
// @Tags EAM-设备组织
// @Param body body model.Factory true "厂区信息"
// @Success 200 {object} response.Response
// @Router /eam/factory [post]
func (h *EquipmentOrgHandler) CreateFactory(c *gin.Context) {
	var req model.Factory
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID

	if err := h.svc.CreateFactory(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// UpdateFactory 更新厂区
// @Summary 更新厂区
// @Tags EAM-设备组织
// @Param id path int true "厂区ID"
// @Param body body map[string]interface{} true "更新内容"
// @Success 200 {object} response.Response
// @Router /eam/factory/{id} [put]
func (h *EquipmentOrgHandler) UpdateFactory(c *gin.Context) {
	idStr := c.Param("id")
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	if err := h.svc.UpdateFactory(c.Request.Context(), idStr, req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeleteFactory 删除厂区
// @Summary 删除厂区
// @Tags EAM-设备组织
// @Param id path int true "厂区ID"
// @Success 200 {object} response.Response
// @Router /eam/factory/{id} [delete]
func (h *EquipmentOrgHandler) DeleteFactory(c *gin.Context) {
	idStr := c.Param("id")
	if err := h.svc.DeleteFactory(c.Request.Context(), idStr); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
