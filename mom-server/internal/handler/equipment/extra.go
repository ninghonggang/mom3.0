package equipment

import (
	"strconv"

	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type EquipmentCheckHandler struct {
	svc *service.EquipmentCheckService
}

func NewEquipmentCheckHandler(svc *service.EquipmentCheckService) *EquipmentCheckHandler {
	return &EquipmentCheckHandler{svc: svc}
}

func (h *EquipmentCheckHandler) List(c *gin.Context) {
	list, total, err := h.svc.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *EquipmentCheckHandler) Create(c *gin.Context) {
	var req model.EquipmentCheck
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *EquipmentCheckHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	check, err := h.svc.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, check)
}

func (h *EquipmentCheckHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req model.EquipmentCheck
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.Update(c.Request.Context(), uint(id), map[string]interface{}{
		"equipment_id":  req.EquipmentID,
		"equipment_code": req.EquipmentCode,
		"equipment_name": req.EquipmentName,
		"check_plan_id":  req.CheckPlanID,
		"check_user_id":  req.CheckUserID,
		"check_user_name": req.CheckUserName,
		"check_date":     req.CheckDate,
		"check_result":   req.CheckResult,
		"status":         req.Status,
		"remark":         req.Remark,
	}); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *EquipmentCheckHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.svc.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

type EquipmentMaintenanceHandler struct {
	svc *service.EquipmentMaintenanceService
}

func NewEquipmentMaintenanceHandler(svc *service.EquipmentMaintenanceService) *EquipmentMaintenanceHandler {
	return &EquipmentMaintenanceHandler{svc: svc}
}

func (h *EquipmentMaintenanceHandler) List(c *gin.Context) {
	list, total, err := h.svc.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *EquipmentMaintenanceHandler) Create(c *gin.Context) {
	var req model.EquipmentMaintenance
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *EquipmentMaintenanceHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	maint, err := h.svc.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, maint)
}

func (h *EquipmentMaintenanceHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req model.EquipmentMaintenance
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.Update(c.Request.Context(), uint(id), map[string]interface{}{
		"equipment_id":    req.EquipmentID,
		"equipment_code":  req.EquipmentCode,
		"equipment_name":  req.EquipmentName,
		"maint_type":      req.MaintType,
		"maint_plan_id":   req.MaintPlanID,
		"maint_user_id":   req.MaintUserID,
		"maint_user_name": req.MaintUserName,
		"maint_date":      req.MaintDate,
		"start_time":      req.StartTime,
		"end_time":        req.EndTime,
		"duration":        req.Duration,
		"content":         req.Content,
		"cost":            req.Cost,
		"status":          req.Status,
	}); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *EquipmentMaintenanceHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.svc.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

type EquipmentRepairHandler struct {
	svc *service.EquipmentRepairService
}

func NewEquipmentRepairHandler(svc *service.EquipmentRepairService) *EquipmentRepairHandler {
	return &EquipmentRepairHandler{svc: svc}
}

func (h *EquipmentRepairHandler) List(c *gin.Context) {
	list, total, err := h.svc.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *EquipmentRepairHandler) Create(c *gin.Context) {
	var req model.EquipmentRepair
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *EquipmentRepairHandler) Start(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req struct{ RepairUserID int64 }
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.Start(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *EquipmentRepairHandler) Complete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.svc.Complete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *EquipmentRepairHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	repair, err := h.svc.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, repair)
}

func (h *EquipmentRepairHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req model.EquipmentRepair
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.Update(c.Request.Context(), uint(id), map[string]interface{}{
		"equipment_id":    req.EquipmentID,
		"equipment_code": req.EquipmentCode,
		"equipment_name": req.EquipmentName,
		"fault_desc":     req.FaultDesc,
		"fault_time":     req.FaultTime,
		"report_user_id":  req.ReportUserID,
		"repair_user_id":  req.RepairUserID,
		"start_time":      req.StartTime,
		"end_time":        req.EndTime,
		"duration":        req.Duration,
		"repair_content":  req.RepairContent,
		"cost":            req.Cost,
		"status":          req.Status,
	}); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *EquipmentRepairHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.svc.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

type SparePartHandler struct {
	svc *service.SparePartService
}

func NewSparePartHandler(svc *service.SparePartService) *SparePartHandler {
	return &SparePartHandler{svc: svc}
}

func (h *SparePartHandler) List(c *gin.Context) {
	list, total, err := h.svc.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *SparePartHandler) Create(c *gin.Context) {
	var req model.SparePart
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *SparePartHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	sp, err := h.svc.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, sp)
}

func (h *SparePartHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req model.SparePart
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.Update(c.Request.Context(), uint(id), map[string]interface{}{
		"spare_part_code": req.SparePartCode,
		"spare_part_name": req.SparePartName,
		"spec":            req.Spec,
		"unit":            req.Unit,
		"quantity":        req.Quantity,
		"min_quantity":    req.MinQuantity,
		"price":           req.Price,
		"supplier":        req.Supplier,
		"status":          req.Status,
	}); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *SparePartHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.svc.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
