package mdm

import (
	"strconv"

	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type ShiftHandler struct {
	shiftService *service.MdmShiftService
}

func NewShiftHandler(shiftSvc *service.MdmShiftService) *ShiftHandler {
	return &ShiftHandler{shiftService: shiftSvc}
}

// List 获取班次列表
func (h *ShiftHandler) List(c *gin.Context) {
	list, total, err := h.shiftService.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// Get 获取单个班次
func (h *ShiftHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	shift, err := h.shiftService.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, shift)
}

// Create 创建班次
func (h *ShiftHandler) Create(c *gin.Context) {
	var req model.MdmShift
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.shiftService.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// Update 更新班次
func (h *ShiftHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req model.MdmShift
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.shiftService.Update(c.Request.Context(), uint(id), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Delete 删除班次
func (h *ShiftHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.shiftService.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
