package equipment

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type EquipmentBomHandler struct {
	svc *service.EquipmentBomService
}

func NewEquipmentBomHandler(svc *service.EquipmentBomService) *EquipmentBomHandler {
	return &EquipmentBomHandler{svc: svc}
}

// ListByEquipment 获取设备BOM列表
// GET /equipment/bom/list?equipment_id=xxx
func (h *EquipmentBomHandler) ListByEquipment(c *gin.Context) {
	equipmentIDStr := c.Query("equipment_id")
	if equipmentIDStr == "" {
		response.BadRequest(c, "equipment_id is required")
		return
	}
	equipmentID, err := strconv.ParseInt(equipmentIDStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid equipment_id")
		return
	}
	list, err := h.svc.ListByEquipmentID(c.Request.Context(), equipmentID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": len(list)})
}

// List 获取BOM列表(分页)
// GET /equipment/bom/page
func (h *EquipmentBomHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	equipmentID, _ := strconv.ParseInt(c.Query("equipment_id"), 10, 64)
	materialCode := c.Query("material_code")
	materialName := c.Query("material_name")
	isCritical, _ := strconv.Atoi(c.Query("is_critical"))

	req := &model.EquipmentBomQuery{
		EquipmentID: equipmentID,
		MaterialCode: materialCode,
		MaterialName: materialName,
		IsCritical: isCritical,
		Page: page,
		PageSize: pageSize,
	}

	list, total, err := h.svc.List(c.Request.Context(), tenantID, req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// Get 获取单个BOM
// GET /equipment/bom/:id
func (h *EquipmentBomHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	bom, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, bom)
}

// Create 创建BOM
// POST /equipment/bom
func (h *EquipmentBomHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	var req model.EquipmentBomCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	bom, err := h.svc.Create(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, bom)
}

// Update 更新BOM
// PUT /equipment/bom/:id
func (h *EquipmentBomHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req model.EquipmentBomCreateReq
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

// Delete 删除BOM
// DELETE /equipment/bom/:id
func (h *EquipmentBomHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}