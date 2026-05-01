package eam

import (
	"github.com/gin-gonic/gin"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
)

// AssetHandler EAM设备资产处理器
type AssetHandler struct {
	svc *service.EquipmentService
}

// NewAssetHandler 创建EAM设备资产处理器
func NewAssetHandler(s *service.EquipmentService) *AssetHandler {
	return &AssetHandler{svc: s}
}

// List 获取设备资产列表
// GET /eam/asset/list
func (h *AssetHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	list, total, err := h.svc.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// Get 获取设备资产详情
// GET /eam/asset/:id
func (h *AssetHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	equipment, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, equipment)
}

// Create 创建设备资产
// POST /eam/asset
func (h *AssetHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var req model.Equipment
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req.TenantID = tenantID

	if err := h.svc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// Update 更新设备资产
// PUT /eam/asset/:id
func (h *AssetHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	var req model.Equipment
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Delete 删除设备资产
// DELETE /eam/asset/:id
func (h *AssetHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// AssetRequest 设备资产请求结构
type AssetRequest struct {
	EquipmentCode string  `json:"equipment_code"`
	EquipmentName string  `json:"equipment_name"`
	EquipmentType string  `json:"equipment_type"`
	Brand         string  `json:"brand"`
	Model         string  `json:"model"`
	Status        int     `json:"status"`
	WorkshopID    int64   `json:"workshop_id"`
	Supplier      string  `json:"supplier"`
	PurchasePrice float64 `json:"purchase_price"`
}

// GetByCode 根据编码获取设备资产
// GET /eam/asset/code/:code
func (h *AssetHandler) GetByCode(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		response.BadRequest(c, "code is required")
		return
	}

	// 暂未实现根据code查询，先返回空数据
	response.Success(c, nil)
}

// ListByStatus 按状态获取设备资产列表
// GET /eam/asset/list-by-status
func (h *AssetHandler) ListByStatus(c *gin.Context) {
	statusStr := c.DefaultQuery("status", "")

	list, total, err := h.svc.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	// 如果提供了状态参数，进行过滤
	if statusStr != "" {
		// 暂未实现过滤
		_ = statusStr
	}

	response.Success(c, gin.H{"list": list, "total": total})
}

// ListByType 按类型获取设备资产列表
// GET /eam/asset/list-by-type
func (h *AssetHandler) ListByType(c *gin.Context) {
	equipmentType := c.DefaultQuery("type", "")

	list, total, err := h.svc.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	// 如果提供了类型参数，进行过滤
	if equipmentType != "" {
		filtered := make([]model.Equipment, 0)
		for _, e := range list {
			if e.EquipmentType == equipmentType {
				filtered = append(filtered, e)
			}
		}
		list = filtered
		total = int64(len(filtered))
	}

	response.Success(c, gin.H{"list": list, "total": total})
}

// ListByWorkshop 按车间获取设备资产列表
// GET /eam/asset/list-by-workshop
func (h *AssetHandler) ListByWorkshop(c *gin.Context) {
	workshopIDStr := c.DefaultQuery("workshop_id", "")

	list, total, err := h.svc.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	// 如果提供了车间ID参数，进行过滤
	if workshopIDStr != "" {
		// 暂未实现精确过滤
		_ = workshopIDStr
	}

	response.Success(c, gin.H{"list": list, "total": total})
}

// BatchDelete 批量删除设备资产
// DELETE /eam/asset/batch
func (h *AssetHandler) BatchDelete(c *gin.Context) {
	var req struct {
		IDs []string `json:"ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	for _, id := range req.IDs {
		if err := h.svc.Delete(c.Request.Context(), id); err != nil {
			response.ErrorMsg(c, err.Error())
			return
		}
	}
	response.Success(c, nil)
}

// GetStatus 获取设备资产状态
// GET /eam/asset/:id/status
func (h *AssetHandler) GetStatus(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	status, err := h.svc.GetStatus(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"status": status})
}

// Start 使用设备资产
// POST /eam/asset/:id/start
func (h *AssetHandler) Start(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	// 设置状态为运行中(1)
	var req model.Equipment
	req.Status = 1

	if err := h.svc.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Stop 停用设备资产
// POST /eam/asset/:id/stop
func (h *AssetHandler) Stop(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	// 设置状态为停机(2)
	var req model.Equipment
	req.Status = 2

	if err := h.svc.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Scrap 报废设备资产
// POST /eam/asset/:id/scrap
func (h *AssetHandler) Scrap(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	// 设置状态为报废(5)
	var req model.Equipment
	req.Status = 5

	if err := h.svc.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Transfer 转移设备资产
// POST /eam/asset/:id/transfer
func (h *AssetHandler) Transfer(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	var req struct {
		WorkshopID   int64  `json:"workshop_id"`
		WorkshopName string `json:"workshop_name"`
		LineID       *int64 `json:"line_id"`
		LineName     string `json:"line_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// 暂未实现完整的转移逻辑
	response.Success(c, nil)
}
