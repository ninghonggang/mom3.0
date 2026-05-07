package mes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
)

// TempBOMHandler MES日计划临时替代BOM处理器
type TempBOMHandler struct {
	tempBOMSvc *service.TempBOMService
}

func NewTempBOMHandler(tempBOMSvc *service.TempBOMService) *TempBOMHandler {
	return &TempBOMHandler{tempBOMSvc: tempBOMSvc}
}

// Create POST /mes/orderday/temp-bom/create
func (h *TempBOMHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	username := middleware.GetUsername(c)

	var req model.TempBOMCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tempBOM, err := h.tempBOMSvc.Create(c.Request.Context(), tenantID, &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, tempBOM)
}

// Update PUT /mes/orderday/temp-bom/update
func (h *TempBOMHandler) Update(c *gin.Context) {
	username := middleware.GetUsername(c)

	var req model.TempBOMUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	tempBOM, err := h.tempBOMSvc.Update(c.Request.Context(), id, &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, tempBOM)
}

// Delete DELETE /mes/orderday/temp-bom/:id
func (h *TempBOMHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	if err := h.tempBOMSvc.Delete(c.Request.Context(), int64(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.SuccessWithMsg(c, "删除成功", nil)
}

// ListByOrderDayItem GET /mes/orderday/temp-bom/listByOrderDayItem?orderDayItemId=
func (h *TempBOMHandler) ListByOrderDayItem(c *gin.Context) {
	orderDayItemID, err := strconv.ParseInt(c.Query("orderDayItemId"), 10, 64)
	if err != nil || orderDayItemID <= 0 {
		response.BadRequest(c, "无效的日计划明细项ID")
		return
	}

	list, err := h.tempBOMSvc.ListByOrderDayItemID(c.Request.Context(), orderDayItemID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, gin.H{"list": list, "total": len(list)})
}

// Get GET /mes/orderday/temp-bom/:id
func (h *TempBOMHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	tempBOM, err := h.tempBOMSvc.GetByID(c.Request.Context(), int64(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, tempBOM)
}

// Approve PUT /mes/orderday/temp-bom/approve
func (h *TempBOMHandler) Approve(c *gin.Context) {
	username := middleware.GetUsername(c)

	var req model.TempBOMApprove
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.tempBOMSvc.Approve(c.Request.Context(), req.ID, req.Status, username); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.SuccessWithMsg(c, "审核成功", nil)
}
