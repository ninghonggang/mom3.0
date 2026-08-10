package service

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"mom-platform/services/andon-service/internal/model"
	"mom-platform/services/andon-service/internal/repository"
)

type AndonService struct {
	repo   repository.Repository
	redis  *redis.Client
	logger *zap.Logger
}

func NewAndonService(repo repository.Repository, redis *redis.Client, logger *zap.Logger) *AndonService {
	return &AndonService{repo: repo, redis: redis, logger: logger}
}

// lastN 安全地返回字符串末尾 n 个字符；长度不足时返回原串，空串返回 "0"。
// 避免对空 tenantID 做负索引切片导致 panic。
func lastN(s string, n int) string {
	if s == "" {
		return "0"
	}
	if len(s) <= n {
		return s
	}
	return s[len(s)-n:]
}

// TriggerAndon 触发安灯呼叫
func (s *AndonService) TriggerAndon(ctx context.Context, tenantID, workstationID, reporterID, andonType, description string) (*model.AndonCall, error) {
	now := time.Now()
	andonNo := fmt.Sprintf("AN-%s-%d", lastN(tenantID, 8), now.UnixNano())

	call := &model.AndonCall{
		TenantID:      tenantID,
		AndonNo:       andonNo,
		WorkstationID: workstationID,
		ReporterID:    reporterID,
		AndonType:     andonType,
		Description:   description,
		Status:        "TRIGGERED",
		TriggeredAt:   now,
	}

	if err := s.repo.CreateAndonCall(ctx, call); err != nil {
		return nil, fmt.Errorf("create andon call: %w", err)
	}

	// 记录创建动作
	action := &model.AndonAction{
		AndonID:    call.ID,
		ActionType: "TRIGGERED",
		ActionDesc: fmt.Sprintf("Andon triggered: %s", description),
		OperatorID: reporterID,
		ActionTime: now,
	}
	_ = s.repo.CreateAndonAction(ctx, action)

	s.logger.Info("andon call triggered",
		zap.String("andon_no", andonNo),
		zap.String("type", andonType),
	)
	return call, nil
}

// AcknowledgeAndon 确认安灯呼叫
func (s *AndonService) AcknowledgeAndon(ctx context.Context, andonID uint, operatorID string) (*model.AndonCall, error) {
	call, err := s.repo.GetAndonCall(ctx, andonID)
	if err != nil {
		return nil, fmt.Errorf("get andon call: %w", err)
	}

	if call.Status != "TRIGGERED" && call.Status != "ESCALATED" {
		return nil, fmt.Errorf("cannot acknowledge andon in status: %s", call.Status)
	}

	now := time.Now()
	responseSecs := int64(now.Sub(call.TriggeredAt).Seconds())

	call.Status = "ACKNOWLEDGED"
	call.AcknowledgedAt = &now
	call.ResponseSeconds = &responseSecs

	if err := s.repo.UpdateAndonCall(ctx, call); err != nil {
		return nil, fmt.Errorf("update andon call: %w", err)
	}

	// 记录确认动作
	action := &model.AndonAction{
		AndonID:    call.ID,
		ActionType: "ACKNOWLEDGED",
		ActionDesc: fmt.Sprintf("Andon acknowledged, response time: %ds", responseSecs),
		OperatorID: operatorID,
		ActionTime: now,
	}
	_ = s.repo.CreateAndonAction(ctx, action)

	s.logger.Info("andon acknowledged",
		zap.String("andon_no", call.AndonNo),
		zap.Int64("response_seconds", responseSecs),
	)
	return call, nil
}

// ResolveAndon 解决安灯呼叫
func (s *AndonService) ResolveAndon(ctx context.Context, andonID uint, operatorID, resolution string) (*model.AndonCall, error) {
	call, err := s.repo.GetAndonCall(ctx, andonID)
	if err != nil {
		return nil, fmt.Errorf("get andon call: %w", err)
	}

	now := time.Now()
	call.Status = "RESOLVED"
	call.ResolvedAt = &now

	if err := s.repo.UpdateAndonCall(ctx, call); err != nil {
		return nil, fmt.Errorf("update andon call: %w", err)
	}

	action := &model.AndonAction{
		AndonID:    call.ID,
		ActionType: "RESOLVED",
		ActionDesc: resolution,
		OperatorID: operatorID,
		ActionTime: now,
	}
	_ = s.repo.CreateAndonAction(ctx, action)

	s.logger.Info("andon resolved",
		zap.String("andon_no", call.AndonNo),
	)
	return call, nil
}

// EscalateAndon 升级安灯呼叫
func (s *AndonService) EscalateAndon(ctx context.Context, andonID uint, operatorID, reason string) (*model.AndonCall, error) {
	call, err := s.repo.GetAndonCall(ctx, andonID)
	if err != nil {
		return nil, fmt.Errorf("get andon call: %w", err)
	}

	now := time.Now()
	call.Status = "ESCALATED"

	if err := s.repo.UpdateAndonCall(ctx, call); err != nil {
		return nil, fmt.Errorf("update andon call: %w", err)
	}

	action := &model.AndonAction{
		AndonID:    call.ID,
		ActionType: "ESCALATED",
		ActionDesc: reason,
		OperatorID: operatorID,
		ActionTime: now,
	}
	_ = s.repo.CreateAndonAction(ctx, action)

	s.logger.Info("andon escalated",
		zap.String("andon_no", call.AndonNo),
		zap.String("reason", reason),
	)
	return call, nil
}

// GetAndonCall 获取安灯呼叫详情
func (s *AndonService) GetAndonCall(ctx context.Context, andonID uint) (*model.AndonCall, error) {
	return s.repo.GetAndonCall(ctx, andonID)
}

// ListAndonCalls 分页列表安灯呼叫
func (s *AndonService) ListAndonCalls(ctx context.Context, page, pageSize int, andonType, status, workstationID string) ([]*model.AndonCall, int64, error) {
	offset := (page - 1) * pageSize
	return s.repo.ListAndonCalls(ctx, offset, pageSize, andonType, status, workstationID)
}

// CreateAlertConfig 创建告警配置
func (s *AndonService) CreateAlertConfig(ctx context.Context, configCode, configName, triggerType, severity, triggerCondition, notifyChannels string) (*model.AlertConfig, error) {
	cfg := &model.AlertConfig{
		ConfigCode:       configCode,
		ConfigName:       configName,
		TriggerType:      triggerType,
		Severity:         severity,
		TriggerCondition: triggerCondition,
		NotifyChannels:   notifyChannels,
		Status:           "ENABLED",
	}

	if err := s.repo.CreateAlertConfig(ctx, cfg); err != nil {
		return nil, fmt.Errorf("create alert config: %w", err)
	}

	s.logger.Info("alert config created", zap.String("config_code", configCode))
	return cfg, nil
}

// TriggerAlert 根据配置编码触发告警
func (s *AndonService) TriggerAlert(ctx context.Context, configCode, targetID, targetType string) (*model.Alert, error) {
	cfg, err := s.repo.GetAlertConfigByCode(ctx, configCode)
	if err != nil {
		return nil, fmt.Errorf("get alert config by code %q: %w", configCode, err)
	}
	return s.triggerWithConfig(ctx, cfg, targetID, targetType)
}

// TriggerAlertByConfigID 根据配置主键触发告警
func (s *AndonService) TriggerAlertByConfigID(ctx context.Context, configID uint, targetID, targetType string) (*model.Alert, error) {
	if configID == 0 {
		return nil, fmt.Errorf("config_id is required")
	}
	cfg, err := s.repo.GetAlertConfig(ctx, configID)
	if err != nil {
		return nil, fmt.Errorf("get alert config by id %d: %w", configID, err)
	}
	return s.triggerWithConfig(ctx, cfg, targetID, targetType)
}

// triggerWithConfig 使用已解析的配置创建告警记录
func (s *AndonService) triggerWithConfig(ctx context.Context, cfg *model.AlertConfig, targetID, targetType string) (*model.Alert, error) {
	configCode := cfg.ConfigCode
	if targetType == "" {
		targetType = "UNKNOWN"
	}

	now := time.Now()
	alert := &model.Alert{
		ConfigID:    cfg.ID,
		TargetID:    targetID,
		TargetType:  targetType,
		Status:      "ACTIVE",
		TriggeredAt: now,
	}

	if err := s.repo.CreateAlert(ctx, alert); err != nil {
		return nil, fmt.Errorf("create alert: %w", err)
	}

	// P0 alerts: full channel notification (stub)
	if cfg.Severity == "P0" {
		s.logger.Warn("P0 alert triggered — full channel notification",
			zap.String("config_code", configCode),
			zap.String("target_id", targetID),
			zap.String("notify_channels", cfg.NotifyChannels),
		)
	}

	s.logger.Info("alert triggered",
		zap.String("config_code", configCode),
		zap.String("severity", cfg.Severity),
		zap.Uint("alert_id", alert.ID),
	)
	return alert, nil
}

// AcknowledgeAlert 确认告警
func (s *AndonService) AcknowledgeAlert(ctx context.Context, alertID uint) (*model.Alert, error) {
	alert, err := s.repo.GetAlert(ctx, alertID)
	if err != nil {
		return nil, fmt.Errorf("get alert: %w", err)
	}

	if alert.Status != "ACTIVE" && alert.Status != "ESCALATED" {
		return nil, fmt.Errorf("cannot acknowledge alert in status: %s", alert.Status)
	}

	now := time.Now()
	alert.Status = "ACKNOWLEDGED"
	alert.AcknowledgedAt = &now

	if err := s.repo.UpdateAlert(ctx, alert); err != nil {
		return nil, fmt.Errorf("update alert: %w", err)
	}

	s.logger.Info("alert acknowledged", zap.Uint("alert_id", alert.ID))
	return alert, nil
}

// ResolveAlert 解决告警
func (s *AndonService) ResolveAlert(ctx context.Context, alertID uint) (*model.Alert, error) {
	alert, err := s.repo.GetAlert(ctx, alertID)
	if err != nil {
		return nil, fmt.Errorf("get alert: %w", err)
	}

	now := time.Now()
	alert.Status = "RESOLVED"
	alert.ResolvedAt = &now

	if err := s.repo.UpdateAlert(ctx, alert); err != nil {
		return nil, fmt.Errorf("update alert: %w", err)
	}

	s.logger.Info("alert resolved", zap.Uint("alert_id", alert.ID))
	return alert, nil
}

// ListAlerts 分页列表告警
func (s *AndonService) ListAlerts(ctx context.Context, page, pageSize int, status, severity string) ([]*model.Alert, int64, error) {
	offset := (page - 1) * pageSize
	return s.repo.ListAlerts(ctx, offset, pageSize, status, severity)
}

// EscalateAlert 告警升级：5min -> L1, 15min -> L2, 30min -> L3
func (s *AndonService) EscalateAlert(ctx context.Context, alertID uint) ([]*model.AlertEscalation, error) {
	alert, err := s.repo.GetAlert(ctx, alertID)
	if err != nil {
		return nil, fmt.Errorf("get alert: %w", err)
	}

	if alert.Status != "ACTIVE" && alert.Status != "ESCALATED" {
		return nil, fmt.Errorf("cannot escalate alert in status: %s", alert.Status)
	}

	elapsed := time.Since(alert.TriggeredAt)
	existingEscalations, err := s.repo.GetAlertEscalations(ctx, alertID)
	if err != nil {
		return nil, fmt.Errorf("get existing escalations: %w", err)
	}

	maxLevel := 0
	for _, e := range existingEscalations {
		if e.Level > maxLevel {
			maxLevel = e.Level
		}
	}

	// Escalation rules: 5min -> L1, 15min -> L2, 30min -> L3
	var level int
	var timeout int
	switch {
	case elapsed >= 30*time.Minute && maxLevel < 3:
		level = 3
		timeout = 1800
	case elapsed >= 15*time.Minute && maxLevel < 2:
		level = 2
		timeout = 900
	case elapsed >= 5*time.Minute && maxLevel < 1:
		level = 1
		timeout = 300
	default:
		return existingEscalations, nil // no escalation needed yet
	}

	// Get escalate target from config notify_channels (simplified — use supervisor role)
	escalation := &model.AlertEscalation{
		AlertID:          alertID,
		Level:            level,
		EscalateToUserID: "supervisor", // stub — would be config-based
		TimeoutSeconds:   timeout,
	}

	if err := s.repo.CreateAlertEscalation(ctx, escalation); err != nil {
		return nil, fmt.Errorf("create escalation: %w", err)
	}

	alert.Status = "ESCALATED"
	_ = s.repo.UpdateAlert(ctx, alert)

	s.logger.Warn("alert escalated",
		zap.Uint("alert_id", alertID),
		zap.Int("level", level),
		zap.Duration("elapsed", elapsed),
	)

	escalations := append(existingEscalations, escalation)
	return escalations, nil
}
