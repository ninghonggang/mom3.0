package quality

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InspectionFeatureHandler struct {
	service *service.InspectionFeatureService
}

func NewInspectionFeatureHandler(svc *service.InspectionFeatureService) *InspectionFeatureHandler {
	return &InspectionFeatureHandler{service: svc}
}

// ListInspectionFeatures 查询检验特性列表
func (h *InspectionFeatureHandler) ListInspectionFeatures(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := make(map[string]interface{})
	if productIDStr := c.Query("product_id"); productIDStr != "" {
		if productID, err := strconv.ParseUint(productIDStr, 10, 64); err == nil && productID > 0 {
			query["product_id"] = productID
		}
	}
	if inspectionType := c.Query("inspection_type"); inspectionType != "" {
		query["inspection_type"] = inspectionType
	}
	if featureType := c.Query("feature_type"); featureType != "" {
		query["feature_type"] = featureType
	}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}

	list, total, err := h.service.List(c.Request.Context(), uint64(tenantID), query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

// GetInspectionFeature 根据ID获取检验特性
func (h *InspectionFeatureHandler) GetInspectionFeature(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	feature, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, feature)
}

// CreateInspectionFeature 创建检验特性
func (h *InspectionFeatureHandler) CreateInspectionFeature(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var req model.InspectionFeatureCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	username := middleware.GetUsername(c)

	feature, err := h.service.Create(c.Request.Context(), uint64(tenantID), &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, feature)
}

// UpdateInspectionFeature 更新检验特性
func (h *InspectionFeatureHandler) UpdateInspectionFeature(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.InspectionFeatureUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	username := middleware.GetUsername(c)

	err = h.service.Update(c.Request.Context(), id, &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeleteInspectionFeature 删除检验特性
func (h *InspectionFeatureHandler) DeleteInspectionFeature(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	err = h.service.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// BatchCreateInspectionFeature 批量创建检验特性
func (h *InspectionFeatureHandler) BatchCreateInspectionFeature(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var req model.InspectionFeatureBatchCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	username := middleware.GetUsername(c)

	features, err := h.service.BatchCreate(c.Request.Context(), uint64(tenantID), &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, features)
}

// GetFeaturesByProduct 获取产品的所有检验特性
func (h *InspectionFeatureHandler) GetFeaturesByProduct(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	productIDStr := c.Param("productId")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid product id")
		return
	}

	list, err := h.service.GetByProductID(c.Request.Context(), uint64(tenantID), productID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list": list,
	})
}