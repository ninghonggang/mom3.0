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
