package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type AlertService struct {
	ruleRepo       *repository.AlertRuleRepository
	recordRepo     *repository.AlertRecordRepository
	notifyLogRepo  *repository.AlertNotificationLogRepository
	escalationRepo *repository.AlertEscalationRuleRepository
	channelRepo    *repository.NotificationChannelRepository
}

func NewAlertService(
	ruleRepo *repository.AlertRuleRepository,
	recordRepo *repository.AlertRecordRepository,
	notifyLogRepo *repository.AlertNotificationLogRepository,
	escalationRepo *repository.AlertEscalationRuleRepository,
	channelRepo *repository.NotificationChannelRepository,
) *AlertService {
	return &AlertService{
		ruleRepo:       ruleRepo,
		recordRepo:     recordRepo,
		notifyLogRepo:  notifyLogRepo,
		escalationRepo: escalationRepo,
		channelRepo:    channelRepo,
	}
}

// ==================== 告警规则 ====================

func (s *AlertService) ListAlertRules(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.AlertRule, int64, error) {
	return s.ruleRepo.List(ctx, tenantID, query)
}

func (s *AlertService) GetAlertRule(ctx context.Context, id string) (*model.AlertRule, error) {
	var ruleID uint
	_, err := fmt.Sscanf(id, "%d", &ruleID)
	if err != nil {
		return nil, err
	}
	return s.ruleRepo.GetByID(ctx, ruleID)
}

func (s *AlertService) CreateAlertRule(ctx context.Context, rule *model.AlertRule) error {
	return s.ruleRepo.Create(ctx, rule)
}

func (s *AlertService) UpdateAlertRule(ctx context.Context, id string, rule *model.AlertRule) error {
	var ruleID uint
	_, err := fmt.Sscanf(id, "%d", &ruleID)
	if err != nil {
		return err
	}
	return s.ruleRepo.Update(ctx, ruleID, map[string]interface{}{
		"rule_name":             rule.RuleName,
		"alert_type":           rule.AlertType,
		"biz_module":           rule.BizModule,
		"condition_expression":  rule.ConditionExpression,
		"condition_params":     rule.ConditionParams,
		"severity_level":       rule.SeverityLevel,
		"notification_channels": rule.NotificationChannels,
		"notify_templates":     rule.NotifyTemplates,
		"escalation_rule_id":   rule.EscalationRuleID,
		"check_interval":       rule.CheckInterval,
		"is_enabled":          rule.IsEnabled,
	})
}

func (s *AlertService) DeleteAlertRule(ctx context.Context, id string) error {
	var ruleID uint
	_, err := fmt.Sscanf(id, "%d", &ruleID)
	if err != nil {
		return err
	}
	return s.ruleRepo.Delete(ctx, ruleID)
}

func (s *AlertService) EnableAlertRule(ctx context.Context, id string) error {
	var ruleID uint
	_, err := fmt.Sscanf(id, "%d", &ruleID)
	if err != nil {
		return err
	}
	return s.ruleRepo.Update(ctx, ruleID, map[string]interface{}{"is_enabled": 1})
}

func (s *AlertService) DisableAlertRule(ctx context.Context, id string) error {
	var ruleID uint
	_, err := fmt.Sscanf(id, "%d", &ruleID)
	if err != nil {
		return err
	}
	return s.ruleRepo.Update(ctx, ruleID, map[string]interface{}{"is_enabled": 0})
}

// ==================== 告警记录 ====================

func (s *AlertService) ListAlertRecords(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.AlertRecord, int64, error) {
	return s.recordRepo.List(ctx, tenantID, query)
}

func (s *AlertService) GetAlertRecord(ctx context.Context, id string) (*model.AlertRecord, error) {
	var recordID uint
	_, err := fmt.Sscanf(id, "%d", &recordID)
	if err != nil {
		return nil, err
	}
	return s.recordRepo.GetByID(ctx, recordID)
}

func (s *AlertService) CreateAlertRecord(ctx context.Context, record *model.AlertRecord) error {
	return s.recordRepo.Create(ctx, record)
}

func (s *AlertService) AcknowledgeAlert(ctx context.Context, id string, userID int64, userName, remark string) error {
	var recordID uint
	_, err := fmt.Sscanf(id, "%d", &recordID)
	if err != nil {
		return err
	}
	now := time.Now()
	return s.recordRepo.Update(ctx, recordID, map[string]interface{}{
		"status":               "ACKNOWLEDGED",
		"acknowledged_by":      userID,
		"acknowledged_by_name": userName,
		"acknowledged_time":    now,
		"acknowledged_remark": remark,
	})
}

func (s *AlertService) ResolveAlert(ctx context.Context, id string, userID int64, userName, remark string) error {
	var recordID uint
	_, err := fmt.Sscanf(id, "%d", &recordID)
	if err != nil {
		return err
	}
	now := time.Now()
	return s.recordRepo.Update(ctx, recordID, map[string]interface{}{
		"status":            "RESOLVED",
		"resolved_by":       userID,
		"resolved_by_name":  userName,
		"resolved_time":     now,
		"resolution_remark": remark,
	})
}

func (s *AlertService) CloseAlert(ctx context.Context, id string) error {
	var recordID uint
	_, err := fmt.Sscanf(id, "%d", &recordID)
	if err != nil {
		return err
	}
	now := time.Now()
	return s.recordRepo.Update(ctx, recordID, map[string]interface{}{
		"status":      "CLOSED",
		"closed_time": now,
	})
}

func (s *AlertService) GetAlertStatistics(ctx context.Context, tenantID int64) (map[string]interface{}, error) {
	return s.recordRepo.GetStatistics(ctx, tenantID)
}

// ==================== 通知日志 ====================

func (s *AlertService) ListNotificationLogs(ctx context.Context, alertID string) ([]model.AlertNotificationLog, error) {
	var id uint
	_, err := fmt.Sscanf(alertID, "%d", &id)
	if err != nil {
		return nil, err
	}
	return s.notifyLogRepo.ListByAlertID(ctx, id)
}

// ==================== 升级规则 ====================

func (s *AlertService) ListEscalationRules(ctx context.Context, tenantID int64) ([]model.AlertEscalationRule, error) {
	return s.escalationRepo.List(ctx, tenantID)
}

func (s *AlertService) CreateEscalationRule(ctx context.Context, rule *model.AlertEscalationRule) error {
	return s.escalationRepo.Create(ctx, rule)
}

// ==================== 通知渠道管理 ====================

func (s *AlertService) ListChannels(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.NotificationChannel, int64, error) {
	return s.channelRepo.List(ctx, tenantID, query)
}

func (s *AlertService) GetChannel(ctx context.Context, id uint) (*model.NotificationChannel, error) {
	return s.channelRepo.GetByID(ctx, id)
}

func (s *AlertService) CreateChannel(ctx context.Context, ch *model.NotificationChannel) error {
	return s.channelRepo.Create(ctx, ch)
}

func (s *AlertService) UpdateChannel(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.channelRepo.Update(ctx, id, updates)
}

func (s *AlertService) DeleteChannel(ctx context.Context, id uint) error {
	return s.channelRepo.Delete(ctx, id)
}

func (s *AlertService) EnableChannel(ctx context.Context, id uint) error {
	return s.channelRepo.Enable(ctx, id)
}

func (s *AlertService) DisableChannel(ctx context.Context, id uint) error {
	return s.channelRepo.Disable(ctx, id)
}

// ==================== 发送通知 ====================

// SendNotification 发送通知
func (s *AlertService) SendNotification(ctx context.Context, tenantID int64, req *model.SendNotificationRequest) error {
	// 创建通知日志
	log := &model.AlertNotificationLog{
		AlertID:       0,
		AlertNo:       "",
		Channel:       req.ChannelType,
		ReceiverType:  req.ReceiverType,
		ReceiverID:    req.ReceiverID,
		ReceiverName:  req.ReceiverName,
		ReceiverValue: req.ReceiverValue,
		NotificationStatus: stringPtr("PENDING"),
		TenantID:      tenantID,
	}
	if req.AlertID != nil {
		log.AlertID = *req.AlertID
	}
	if req.AlertNo != nil {
		log.AlertNo = *req.AlertNo
	}

	// 根据渠道类型发送
	var err error
	switch req.ChannelType {
	case "IN_SITE":
		err = s.sendInSite(ctx, log, req)
	case "FEISHU":
		err = s.sendFeishu(ctx, log, req)
	case "WECOM":
		err = s.sendWecom(ctx, log, req)
	case "EMAIL":
		err = s.sendEmail(ctx, log, req)
	default:
		err = fmt.Errorf("不支持的渠道类型: %s", req.ChannelType)
	}

	// 更新日志状态
	if err != nil {
		errCode := "SEND_FAILED"
		log.NotificationStatus = &errCode
		log.ErrorMsg = stringPtr(err.Error())
	} else {
		status := "SENT"
		log.NotificationStatus = &status
		now := time.Now()
		log.SentTime = &now
	}

	return s.notifyLogRepo.Create(ctx, log)
}

// sendInSite 发送站内信
func (s *AlertService) sendInSite(ctx context.Context, log *model.AlertNotificationLog, req *model.SendNotificationRequest) error {
	return nil
}

// sendFeishu 发送飞书通知
func (s *AlertService) sendFeishu(ctx context.Context, log *model.AlertNotificationLog, req *model.SendNotificationRequest) error {
	channels, err := s.channelRepo.GetByType(ctx, log.TenantID, "FEISHU")
	if err != nil || len(channels) == 0 {
		return fmt.Errorf("飞书渠道未配置")
	}
	_ = channels[0] // TODO: 使用配置发送
	return nil
}

// sendWecom 发送企业微信通知
func (s *AlertService) sendWecom(ctx context.Context, log *model.AlertNotificationLog, req *model.SendNotificationRequest) error {
	channels, err := s.channelRepo.GetByType(ctx, log.TenantID, "WECOM")
	if err != nil || len(channels) == 0 {
		return fmt.Errorf("企业微信渠道未配置")
	}
	_ = channels[0] // TODO: 使用配置发送
	return nil
}

// sendEmail 发送邮件通知
func (s *AlertService) sendEmail(ctx context.Context, log *model.AlertNotificationLog, req *model.SendNotificationRequest) error {
	channels, err := s.channelRepo.GetByType(ctx, log.TenantID, "EMAIL")
	if err != nil || len(channels) == 0 {
		return fmt.Errorf("邮件渠道未配置")
	}

	// TODO: 实现SMTP邮件发送
	return nil
}
