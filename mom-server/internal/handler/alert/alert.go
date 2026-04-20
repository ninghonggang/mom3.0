package alert

import (
	"fmt"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type AlertHandler struct {
	alertService *service.AlertService
}

func NewAlertHandler(s *service.AlertService) *AlertHandler {
	return &AlertHandler{alertService: s}
}

// ==================== 告警规则 ====================

func (h *AlertHandler) ListRules(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if alertType := c.Query("alert_type"); alertType != "" {
		query["alert_type"] = alertType
	}
	if bizModule := c.Query("biz_module"); bizModule != "" {
		query["biz_module"] = bizModule
	}
	if severity := c.Query("severity_level"); severity != "" {
		query["severity_level"] = severity
	}
	if isEnabled := c.Query("is_enabled"); isEnabled != "" {
		query["is_enabled"] = isEnabled
	}

	list, total, err := h.alertService.ListAlertRules(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *AlertHandler) GetRule(c *gin.Context) {
	id := c.Param("id")
	rule, err := h.alertService.GetAlertRule(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, rule)
}

func (h *AlertHandler) CreateRule(c *gin.Context) {
	var req model.AlertRule
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.alertService.CreateAlertRule(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *AlertHandler) UpdateRule(c *gin.Context) {
	id := c.Param("id")
	var req model.AlertRule
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.alertService.UpdateAlertRule(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *AlertHandler) DeleteRule(c *gin.Context) {
	id := c.Param("id")
	if err := h.alertService.DeleteAlertRule(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *AlertHandler) EnableRule(c *gin.Context) {
	id := c.Param("id")
	if err := h.alertService.EnableAlertRule(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *AlertHandler) DisableRule(c *gin.Context) {
	id := c.Param("id")
	if err := h.alertService.DisableAlertRule(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ==================== 告警记录 ====================

func (h *AlertHandler) ListRecords(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if alertType := c.Query("alert_type"); alertType != "" {
		query["alert_type"] = alertType
	}
	if severity := c.Query("severity_level"); severity != "" {
		query["severity_level"] = severity
	}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}
	if sourceModule := c.Query("source_module"); sourceModule != "" {
		query["source_module"] = sourceModule
	}
	if startDate := c.Query("start_date"); startDate != "" {
		query["start_date"] = startDate
	}
	if endDate := c.Query("end_date"); endDate != "" {
		query["end_date"] = endDate
	}
	if keyword := c.Query("keyword"); keyword != "" {
		query["keyword"] = keyword
	}

	list, total, err := h.alertService.ListAlertRecords(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *AlertHandler) GetRecord(c *gin.Context) {
	id := c.Param("id")
	record, err := h.alertService.GetAlertRecord(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, record)
}

func (h *AlertHandler) AcknowledgeRecord(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		UserID   int64  `json:"user_id"`
		UserName string `json:"user_name"`
		Remark  string `json:"remark"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.alertService.AcknowledgeAlert(c.Request.Context(), id, req.UserID, req.UserName, req.Remark); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *AlertHandler) ResolveRecord(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		UserID   int64  `json:"user_id"`
		UserName string `json:"user_name"`
		Remark  string `json:"remark"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.alertService.ResolveAlert(c.Request.Context(), id, req.UserID, req.UserName, req.Remark); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *AlertHandler) CloseRecord(c *gin.Context) {
	id := c.Param("id")
	if err := h.alertService.CloseAlert(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ==================== 统计 ====================

func (h *AlertHandler) GetStatistics(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	stats, err := h.alertService.GetAlertStatistics(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, stats)
}

// ==================== 通知日志 ====================

func (h *AlertHandler) ListNotificationLogs(c *gin.Context) {
	id := c.Param("id")
	logs, err := h.alertService.ListNotificationLogs(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list": logs,
	})
}

// ==================== 升级规则 ====================

func (h *AlertHandler) ListEscalationRules(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	list, err := h.alertService.ListEscalationRules(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list": list,
	})
}

func (h *AlertHandler) CreateEscalationRule(c *gin.Context) {
	var req model.AlertEscalationRule
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.alertService.CreateEscalationRule(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// ==================== 发送通知 ====================

func (h *AlertHandler) SendNotification(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var req model.SendNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.alertService.SendNotification(c.Request.Context(), tenantID, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ==================== 通知渠道 ====================

func (h *AlertHandler) ListChannels(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if channelType := c.Query("channel_type"); channelType != "" {
		query["channel_type"] = channelType
	}
	if isEnabled := c.Query("is_enabled"); isEnabled != "" {
		query["is_enabled"] = isEnabled
	}

	list, total, err := h.alertService.ListChannels(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *AlertHandler) GetChannel(c *gin.Context) {
	id := c.Param("id")
	var channelID uint
	_, err := fmt.Sscanf(id, "%d", &channelID)
	if err != nil {
		response.ErrorMsg(c, "invalid channel id")
		return
	}
	ch, err := h.alertService.GetChannel(c.Request.Context(), channelID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, ch)
}

func (h *AlertHandler) CreateChannel(c *gin.Context) {
	var req model.NotificationChannel
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.alertService.CreateChannel(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *AlertHandler) UpdateChannel(c *gin.Context) {
	id := c.Param("id")
	var channelID uint
	_, err := fmt.Sscanf(id, "%d", &channelID)
	if err != nil {
		response.ErrorMsg(c, "invalid channel id")
		return
	}
	var req model.NotificationChannel
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	updates := map[string]interface{}{
		"channel_name": req.ChannelName,
		"channel_type": req.ChannelType,
		"config":       req.Config,
		"is_enabled":   req.IsEnabled,
		"priority":     req.Priority,
	}
	if err := h.alertService.UpdateChannel(c.Request.Context(), channelID, updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *AlertHandler) DeleteChannel(c *gin.Context) {
	id := c.Param("id")
	var channelID uint
	_, err := fmt.Sscanf(id, "%d", &channelID)
	if err != nil {
		response.ErrorMsg(c, "invalid channel id")
		return
	}
	if err := h.alertService.DeleteChannel(c.Request.Context(), channelID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *AlertHandler) EnableChannel(c *gin.Context) {
	id := c.Param("id")
	var channelID uint
	_, err := fmt.Sscanf(id, "%d", &channelID)
	if err != nil {
		response.ErrorMsg(c, "invalid channel id")
		return
	}
	if err := h.alertService.EnableChannel(c.Request.Context(), channelID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *AlertHandler) DisableChannel(c *gin.Context) {
	id := c.Param("id")
	var channelID uint
	_, err := fmt.Sscanf(id, "%d", &channelID)
	if err != nil {
		response.ErrorMsg(c, "invalid channel id")
		return
	}
	if err := h.alertService.DisableChannel(c.Request.Context(), channelID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
