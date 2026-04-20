package quality

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InspectionHandler struct {
	service *service.InspectionService
}

func NewInspectionHandler(svc *service.InspectionService) *InspectionHandler {
	return &InspectionHandler{service: svc}
}

// ListPlans 查询检验计划列表
func (h *InspectionHandler) ListPlans(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
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

	list, total, err := h.service.ListPlans(c.Request.Context(), uint64(tenantID), query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

// GetPlanByID 根据ID获取检验计划
func (h *InspectionHandler) GetPlanByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	plan, err := h.service.GetPlanByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, plan)
}

// CreatePlan 创建检验计划
func (h *InspectionHandler) CreatePlan(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var req model.QualityInspectionPlanCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	username := middleware.GetUsername(c)

	plan, err := h.service.CreatePlan(c.Request.Context(), uint64(tenantID), &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, plan)
}

// UpdatePlan 更新检验计划
func (h *InspectionHandler) UpdatePlan(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.QualityInspectionPlanUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	username := middleware.GetUsername(c)

	err = h.service.UpdatePlan(c.Request.Context(), id, &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeletePlan 删除检验计划
func (h *InspectionHandler) DeletePlan(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	err = h.service.DeletePlan(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// CalculateSampleSize 计算抽样方案
func (h *InspectionHandler) CalculateSampleSize(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	inspectionType := c.Query("inspection_type")
	batchSizeStr := c.Query("batch_size")

	if inspectionType == "" || batchSizeStr == "" {
		response.BadRequest(c, "inspection_type and batch_size are required")
		return
	}

	batchSize, err := strconv.Atoi(batchSizeStr)
	if err != nil || batchSize <= 0 {
		response.BadRequest(c, "invalid batch_size")
		return
	}

	result, err := h.service.CalculateSampleSize(c.Request.Context(), uint64(tenantID), inspectionType, batchSize)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, result)
}

// SeedAQLData 初始化AQL标准抽样数据
func (h *InspectionHandler) SeedAQLData(c *gin.Context) {
	err := h.service.SeedAQLData(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "AQL data seeded successfully"})
}
