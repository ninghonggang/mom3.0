package mes

import (
	"strconv"
	"time"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type PersonSkillHandler struct {
	skillService *service.PersonSkillService
}

func NewPersonSkillHandler(skillSvc *service.PersonSkillService) *PersonSkillHandler {
	return &PersonSkillHandler{skillService: skillSvc}
}

// ListPersonSkills 获取人员技能列表
func (h *PersonSkillHandler) ListPersonSkills(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "20")
	personIDStr := c.Query("person_id")
	workshopIDStr := c.Query("workshop_id")
	skillLevel := c.Query("skill_level")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	personID, _ := strconv.ParseInt(personIDStr, 10, 64)
	workshopID, _ := strconv.ParseInt(workshopIDStr, 10, 64)

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	list, total, err := h.skillService.List(c.Request.Context(), tenantID, personID, workshopID, skillLevel, page, pageSize)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// GetPersonSkill 获取单个人员技能
func (h *PersonSkillHandler) GetPersonSkill(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	skill, err := h.skillService.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, skill)
}

// CreatePersonSkill 创建人员技能
func (h *PersonSkillHandler) CreatePersonSkill(c *gin.Context) {
	var req model.PersonSkill
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.skillService.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// UpdatePersonSkill 更新人员技能
func (h *PersonSkillHandler) UpdatePersonSkill(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req model.PersonSkill
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.skillService.Update(c.Request.Context(), uint(id), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeletePersonSkill 删除人员技能
func (h *PersonSkillHandler) DeletePersonSkill(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.skillService.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetPersonSkillDetail 获取人员技能详情（包含评分列表）
func (h *PersonSkillHandler) GetPersonSkillDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	detail, err := h.skillService.GetDetail(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, detail)
}

// EvaluateSkill 添加技能评分
func (h *PersonSkillHandler) EvaluateSkill(c *gin.Context) {
	idStr := c.Param("id")
	_, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req model.PersonSkillScore
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	req.EvaluatedAt = time.Now()
	if err := h.skillService.EvaluateSkill(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// GetPersonCapability 获取人员能力报告
func (h *PersonSkillHandler) GetPersonCapability(c *gin.Context) {
	personIdStr := c.Param("personId")
	personID, err := strconv.ParseInt(personIdStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid personId")
		return
	}
	capability, err := h.skillService.GetPersonCapability(c.Request.Context(), personID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, capability)
}