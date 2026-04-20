package integration

import (
	"strconv"
	"time"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/repository"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type IntegrationHandler struct {
	svc      *service.IntegrationService
	executor *service.IntegrationExecutor
}

func NewIntegrationHandler(svc *service.IntegrationService, executor *service.IntegrationExecutor) *IntegrationHandler {
	return &IntegrationHandler{svc: svc, executor: executor}
}

// ListConfigs GET /integration/interface-config/list
func (h *IntegrationHandler) ListConfigs(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	q := &model.InterfaceConfigQuery{
		TenantID:  tenantID,
		Category: c.Query("category"),
		Status:   c.Query("status"),
		Keyword:  c.Query("keyword"),
		Page:     page,
		PageSize: pageSize,
	}

	list, total, err := h.svc.ListConfigs(c.Request.Context(), q)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// GetConfig GET /integration/interface-config/:id
func (h *IntegrationHandler) GetConfig(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	cfg, fieldMaps, triggers, err := h.svc.GetConfig(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"config":      cfg,
		"field_maps": fieldMaps,
		"triggers":   triggers,
	})
}

// CreateConfig POST /integration/interface-config
func (h *IntegrationHandler) CreateConfig(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var req model.InterfaceConfigCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	cfg, err := h.svc.CreateConfig(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	if cfg == nil {
		response.BadRequest(c, "配置编码已存在")
		return
	}
	response.Success(c, cfg)
}

// UpdateConfig PUT /integration/interface-config/:id
func (h *IntegrationHandler) UpdateConfig(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.InterfaceConfigUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.UpdateConfig(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "更新成功"})
}

// DeleteConfig DELETE /integration/interface-config/:id
func (h *IntegrationHandler) DeleteConfig(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.svc.DeleteConfig(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

// ExecuteConfig POST /integration/interface-config/:id/execute
func (h *IntegrationHandler) ExecuteConfig(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	log, err := h.executor.ExecuteConfig(c.Request.Context(), id, "MANUAL")
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, log)
}

// TestConfig POST /integration/interface-config/:id/test
func (h *IntegrationHandler) TestConfig(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	cfg, err := h.svc.GetConfigByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	fieldMaps, _ := h.svc.GetFieldMaps(c.Request.Context(), id)

	// Optional test data from request body
	var testData map[string]interface{}
	c.ShouldBindJSON(&testData)

	statusCode, body, err := h.executor.TestConfig(c.Request.Context(), cfg, fieldMaps, testData)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"status_code": statusCode,
		"body":        body,
	})
}

// ListFieldMaps GET /integration/interface-config/:id/field-maps
func (h *IntegrationHandler) ListFieldMaps(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	maps, err := h.svc.GetFieldMaps(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": maps})
}

// CreateFieldMap POST /integration/interface-config/:id/field-maps
func (h *IntegrationHandler) CreateFieldMap(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.InterfaceFieldMapCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.CreateFieldMap(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "创建成功"})
}

// UpdateFieldMap PUT /integration/field-map/:id
func (h *IntegrationHandler) UpdateFieldMap(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.InterfaceFieldMapCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.UpdateFieldMap(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "更新成功"})
}

// DeleteFieldMap DELETE /integration/field-map/:id
func (h *IntegrationHandler) DeleteFieldMap(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.svc.DeleteFieldMap(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

// ListTriggers GET /integration/interface-config/:id/triggers
func (h *IntegrationHandler) ListTriggers(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	_, _, triggers, err := h.svc.GetConfig(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": triggers})
}

// CreateTrigger POST /integration/interface-config/:id/triggers
func (h *IntegrationHandler) CreateTrigger(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.InterfaceTriggerCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.CreateTrigger(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "创建成功"})
}

// UpdateTrigger PUT /integration/trigger/:id
func (h *IntegrationHandler) UpdateTrigger(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.InterfaceTriggerCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.UpdateTrigger(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "更新成功"})
}

// DeleteTrigger DELETE /integration/trigger/:id
func (h *IntegrationHandler) DeleteTrigger(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.svc.DeleteTrigger(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

// ListExecutionLogs GET /integration/execution-log/list
func (h *IntegrationHandler) ListExecutionLogs(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	configID, _ := strconv.ParseInt(c.Query("interface_config_id"), 10, 64)

	var startDate, endDate *time.Time
	if sd := c.Query("start_date"); sd != "" {
		if t, err := time.Parse("2006-01-02", sd); err == nil {
			startDate = &t
		}
	}
	if ed := c.Query("end_date"); ed != "" {
		if t, err := time.Parse("2006-01-02", ed); err == nil {
			endDate = &t
		}
	}

	q := &repository.ExecutionLogQuery{
		TenantID:          tenantID,
		InterfaceConfigID: configID,
		Status:    c.Query("status"),
		StartDate: startDate,
		EndDate:   endDate,
		Page:      page,
		PageSize:  pageSize,
	}

	list, total, err := h.svc.ListExecutionLogs(c.Request.Context(), q)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// GetExecutionLog GET /integration/execution-log/:id
func (h *IntegrationHandler) GetExecutionLog(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	log, err := h.svc.GetExecutionLog(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, log)
}

// GetConstantOptions GET /integration/options
// Returns available options for dropdowns
func (h *IntegrationHandler) GetConstantOptions(c *gin.Context) {
	response.Success(c, gin.H{
		"categories": []string{"ERP", "AGV", "MES", "WMS", "OTHER"},
		"directions": []string{"OUTBOUND", "INBOUND"},
		"methods":    []string{"GET", "POST", "PUT", "DELETE"},
		"auth_types": []string{"NONE", "BASIC", "API_KEY", "OAUTH2", "BEARER_TOKEN"},
		"content_types": []string{"JSON", "XML", "FORM", "TEXT"},
		"trigger_types": []string{"MANUAL", "SCHEDULE", "EVENT"},
		"source_types":  []string{"TABLE_QUERY", "API_CALL", "EVENT_PAYLOAD"},
		"map_types":     []string{"CONST", "FIELD", "EXPRESSION", "JSONPATH"},
		"transform_funcs": []string{"upper", "lower", "trim", "date_format", "datetime_format", "md5", "uuid"},
		"event_sources": []string{
			"PRODUCTION_COMPLETE", "QUALITY_INSPECT", "STOCK_IN",
			"STOCK_OUT", "PURCHASE_AWARD", "SALES_SHIP",
		},
	})
}
