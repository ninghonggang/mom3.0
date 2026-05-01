package quality

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
)

// QualityInspectionPlanHandler 质量检验计划Handler
type QualityInspectionPlanHandler struct {
	svc *service.InspectionService
}

// NewQualityInspectionPlanHandler 创建质量检验计划Handler
func NewQualityInspectionPlanHandler(svc *service.InspectionService) *QualityInspectionPlanHandler {
	return &QualityInspectionPlanHandler{svc: svc}
}

// List 获取检验计划列表
func (h *QualityInspectionPlanHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	query := make(map[string]interface{})
	if planCode := c.Query("plan_code"); planCode != "" {
		query["plan_code"] = planCode
	}
	if planName := c.Query("plan_name"); planName != "" {
		query["plan_name"] = planName
	}
	if inspectionType := c.Query("inspection_type"); inspectionType != "" {
		query["inspection_type"] = inspectionType
	}
	if aqlLevel := c.Query("aql_level"); aqlLevel != "" {
		query["aql_level"] = aqlLevel
	}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}

	list, total, err := h.svc.ListPlans(c.Request.Context(), uint64(tenantID), query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// Get 获取检验计划详情
func (h *QualityInspectionPlanHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	plan, err := h.svc.GetPlanByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, plan)
}

// Create 创建检验计划
func (h *QualityInspectionPlanHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var req model.QualityInspectionPlanCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	username := middleware.GetUsername(c)

	plan, err := h.svc.CreatePlan(c.Request.Context(), uint64(tenantID), &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, plan)
}

// Update 更新检验计划
func (h *QualityInspectionPlanHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req model.QualityInspectionPlanUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	username := middleware.GetUsername(c)

	err = h.svc.UpdatePlan(c.Request.Context(), id, &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// Delete 删除检验计划
func (h *QualityInspectionPlanHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.svc.DeletePlan(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}
