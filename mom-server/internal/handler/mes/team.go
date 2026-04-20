package mes

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	teamService *service.MesTeamService
}

func NewTeamHandler(teamSvc *service.MesTeamService) *TeamHandler {
	return &TeamHandler{teamService: teamSvc}
}

// List 获取班组列表
func (h *TeamHandler) List(c *gin.Context) {
	list, total, err := h.teamService.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// Get 获取单个班组
func (h *TeamHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	team, err := h.teamService.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, team)
}

// Create 创建班组
func (h *TeamHandler) Create(c *gin.Context) {
	var req model.MesTeam
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.teamService.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// Update 更新班组
func (h *TeamHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req model.MesTeam
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.teamService.Update(c.Request.Context(), uint(id), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Delete 删除班组
func (h *TeamHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.teamService.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ListMembers 获取班组成员列表
func (h *TeamHandler) ListMembers(c *gin.Context) {
	teamIDStr := c.Query("team_id")
	teamID, err := strconv.ParseUint(teamIDStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid team_id")
		return
	}
	members, err := h.teamService.ListMembers(c.Request.Context(), int64(teamID))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": members, "total": len(members)})
}

// AddMember 添加班组成员
func (h *TeamHandler) AddMember(c *gin.Context) {
	var req model.MesTeamMember
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.teamService.AddMember(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// UpdateMember 更新班组成员
func (h *TeamHandler) UpdateMember(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req model.MesTeamMember
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.teamService.UpdateMember(c.Request.Context(), uint(id), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// RemoveMember 移除班组成员
func (h *TeamHandler) RemoveMember(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.teamService.RemoveMember(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ListShifts 获取班组排班列表
func (h *TeamHandler) ListShifts(c *gin.Context) {
	teamIDStr := c.Query("team_id")
	teamID, err := strconv.ParseUint(teamIDStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid team_id")
		return
	}
	shifts, err := h.teamService.ListShifts(c.Request.Context(), int64(teamID))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": shifts, "total": len(shifts)})
}

// CreateShift 创建班组排班
func (h *TeamHandler) CreateShift(c *gin.Context) {
	var req model.MesTeamShift
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.teamService.CreateShift(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// UpdateShift 更新班组排班
func (h *TeamHandler) UpdateShift(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req model.MesTeamShift
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.teamService.UpdateShift(c.Request.Context(), uint(id), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeleteShift 删除班组排班
func (h *TeamHandler) DeleteShift(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.teamService.DeleteShift(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
