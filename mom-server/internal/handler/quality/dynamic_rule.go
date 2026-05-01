package quality

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
)

// DynamicRuleHandler 动态规则Handler
type DynamicRuleHandler struct {
	db *gorm.DB
}

// NewDynamicRuleHandler 创建动态规则Handler
func NewDynamicRuleHandler(db *gorm.DB) *DynamicRuleHandler {
	return &DynamicRuleHandler{db: db}
}

// List 获取动态规则列表
func (h *DynamicRuleHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var list []model.DynamicRule
	var total int64

	query := h.db.WithContext(c.Request.Context()).Model(&model.DynamicRule{}).Where("tenant_id = ?", tenantID)

	if ruleType := c.Query("rule_type"); ruleType != "" {
		query = query.Where("rule_type = ?", ruleType)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)
	query.Order("created_at DESC").Find(&list)

	response.Success(c, gin.H{"list": list, "total": total})
}

// Get 获取动态规则详情
func (h *DynamicRuleHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var rule model.DynamicRule
	if err := h.db.WithContext(c.Request.Context()).First(&rule, id).Error; err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, rule)
}

// Create 创建动态规则
func (h *DynamicRuleHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID == 0 {
		tenantID = 1
	}

	var req model.DynamicRule
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	req.TenantID = tenantID
	if err := h.db.WithContext(c.Request.Context()).Create(&req).Error; err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, req)
}

// Update 更新动态规则
func (h *DynamicRuleHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	if err := h.db.WithContext(c.Request.Context()).Model(&model.DynamicRule{}).Where("id = ?", id).Updates(req).Error; err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// Delete 删除动态规则
func (h *DynamicRuleHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.db.WithContext(c.Request.Context()).Delete(&model.DynamicRule{}, id).Error; err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// Activate 激活动态规则
func (h *DynamicRuleHandler) Activate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.db.WithContext(c.Request.Context()).Model(&model.DynamicRule{}).Where("id = ?", id).Update("status", "ACTIVE").Error; err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, nil)
}
