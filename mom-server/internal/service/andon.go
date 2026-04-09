package service

import (
	"context"
	"fmt"
	"log"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"time"

	"go.uber.org/zap"
)

type AndonService struct {
	andonRepo        *repository.AndonRepository
	ruleRepo         *repository.EscalationRuleRepository
	escalationRepo   *repository.EscalationLogRepository
	notifyRepo       *repository.NotificationLogRepository
	logger           *zap.Logger
}

func NewAndonService(
	andonRepo *repository.AndonRepository,
	ruleRepo *repository.EscalationRuleRepository,
	escalationRepo *repository.EscalationLogRepository,
	notifyRepo *repository.NotificationLogRepository,
	logger *zap.Logger,
) *AndonService {
	if logger == nil {
		logger, _ = zap.NewProduction()
	}
	return &AndonService{
		andonRepo:      andonRepo,
		ruleRepo:       ruleRepo,
		escalationRepo: escalationRepo,
		notifyRepo:     notifyRepo,
		logger:         logger,
	}
}

// AndonQuery 查询参数
type AndonQuery struct {
	TenantID    int64
	WorkshopID  int64
	Status      string
	AndonType   string
	CallNo      string
	StartDate   *time.Time
	EndDate     *time.Time
	Page        int
	PageSize    int
}

// List 查询呼叫列表
func (s *AndonService) List(ctx context.Context, query *AndonQuery) ([]model.AndonCall, int64, error) {
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 20
	}
	return s.andonRepo.List(ctx, query.TenantID, query.WorkshopID, query.Status, query.AndonType, query.CallNo, query.StartDate, query.EndDate, query.Page, query.PageSize)
}

// GetByID 根据ID获取
func (s *AndonService) GetByID(ctx context.Context, id int64) (*model.AndonCall, error) {
	return s.andonRepo.GetByID(ctx, id)
}

// Create 创建呼叫
func (s *AndonService) Create(ctx context.Context, call *model.AndonCall) error {
	// 生成呼叫编号
	callNo := fmt.Sprintf("ANDON-%s-%04d", time.Now().Format("20060102"), time.Now().Unix()%10000)
	call.CallNo = callNo
	call.CallLevel = 1
	if call.Status == "" {
		call.Status = "CALLING"
	}
	if call.CallTime.IsZero() {
		call.CallTime = time.Now()
	}
	return s.andonRepo.Create(ctx, call)
}

// Respond 响应呼叫
func (s *AndonService) Respond(ctx context.Context, id int64, responseBy string, responseRemark string) error {
	call, err := s.andonRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("呼叫不存在: %w", err)
	}

	updates := map[string]interface{}{
		"status":        "RESPONDED",
		"response_by":   responseBy,
		"response_time": time.Now(),
	}

	if call.ResponseTime != nil {
		// 计算响应时长
		responseDuration := int(time.Since(*call.ResponseTime).Seconds())
		updates["response_duration"] = responseDuration
	}

	if responseRemark != "" {
		updates["remark"] = responseRemark
	}

	return s.andonRepo.Update(ctx, id, updates)
}

// Resolve 处理完成
func (s *AndonService) Resolve(ctx context.Context, id int64, handleResult, handleRemark string, relatedRepairID, relatedNCRID *int64) error {
	call, err := s.andonRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("呼叫不存在: %w", err)
	}

	now := time.Now()
	updates := map[string]interface{}{
		"status":        "RESOLVED",
		"handle_by":    call.ResponseBy,
		"handle_time":   now,
		"handle_result": handleResult,
	}

	if handleRemark != "" {
		updates["handle_remark"] = handleRemark
	}

	if relatedRepairID != nil {
		updates["related_repair_id"] = *relatedRepairID
	}
	if relatedNCRID != nil {
		updates["related_ncr_id"] = *relatedNCRID
	}

	// 计算处理时长
	if call.ResponseTime != nil {
		handleDuration := int(now.Sub(*call.ResponseTime).Seconds())
		updates["handle_duration"] = handleDuration
	}

	return s.andonRepo.Update(ctx, id, updates)
}

// Escalate 手动升级
func (s *AndonService) Escalate(ctx context.Context, id int64, toLevel int, escalationReason, triggerUser string) error {
	call, err := s.andonRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("呼叫不存在: %w", err)
	}

	if call.Status == "RESOLVED" || call.Status == "CLOSED" {
		return fmt.Errorf("已关闭的呼叫不能升级")
	}

	// 记录升级日志
	escalationLog := &model.AndonEscalationLog{
		CallID:         int64(id),
		FromLevel:      call.CallLevel,
		ToLevel:        toLevel,
		EscalationType: "MANUAL",
		TriggerUser:    triggerUser,
		TriggerReason:  escalationReason,
		CreatedAt:      time.Now(),
	}
	if err := s.escalationRepo.Create(ctx, escalationLog); err != nil {
		s.logger.Error("创建升级日志失败", zap.Error(err))
	}

	// 获取升级规则
	rule, err := s.ruleRepo.GetApplicableRule(ctx, call.TenantID, call.WorkshopID, call.AndonType, call.Priority)
	if err != nil {
		s.logger.Warn("获取升级规则失败", zap.Error(err))
	}

	// 发送通知
	if rule != nil {
		go s.sendEscalationNotifications(call, toLevel, rule)
	}

	// 更新呼叫
	updates := map[string]interface{}{
		"call_level":        toLevel,
		"is_escalated":      1,
		"escalated_at":      time.Now(),
		"escalation_count":  call.EscalationCount + 1,
		"status":            "HANDLING",
	}

	return s.andonRepo.Update(ctx, id, updates)
}

// CheckAndEscalate 检查并升级(定时任务调用)
func (s *AndonService) CheckAndEscalate(ctx context.Context, callID int64) error {
	call, err := s.andonRepo.GetByID(ctx, callID)
	if err != nil {
		return fmt.Errorf("获取呼叫失败: %w", err)
	}

	// 已关闭的呼叫不检查
	if call.Status == "RESOLVED" || call.Status == "CLOSED" {
		return nil
	}

	// 获取适用的升级规则
	rule, err := s.ruleRepo.GetApplicableRule(ctx, call.TenantID, call.WorkshopID, call.AndonType, call.Priority)
	if err != nil {
		s.logger.Warn("获取升级规则失败", zap.Int64("call_id", callID), zap.Error(err))
		return nil // 不因为规则获取失败而中断
	}

	// 计算当前应该的等级
	elapsed := time.Since(call.CallTime).Minutes()
	targetLevel := s.calculateTargetLevel(rule, elapsed, call.CallLevel)

	// 如果目标等级高于当前等级，执行升级
	if targetLevel > call.CallLevel && targetLevel <= rule.MaxEscalationLevel {
		return s.doEscalate(ctx, call, rule, targetLevel, "TIMEOUT")
	}

	return nil
}

// CheckAllPendingEscalations 检查所有待升级的呼叫(定时任务调用)
func (s *AndonService) CheckAllPendingEscalations(ctx context.Context) {
	calls, err := s.andonRepo.ListPendingEscalations(ctx)
	if err != nil {
		s.logger.Error("查询待升级呼叫失败", zap.Error(err))
		return
	}

	for _, call := range calls {
		if err := s.CheckAndEscalate(ctx, call.ID); err != nil {
			s.logger.Error("升级检查失败",
				zap.Int64("call_id", call.ID),
				zap.Error(err))
		}
	}
}

// calculateTargetLevel 计算目标等级
func (s *AndonService) calculateTargetLevel(rule *model.AndonEscalationRule, elapsed float64, currentLevel int) int {
	if rule == nil {
		// 默认规则: L1=5分钟, L2=10分钟, L3=15分钟, L4=20分钟
		switch {
		case elapsed >= 20:
			return 4
		case elapsed >= 15:
			return 3
		case elapsed >= 10:
			return 2
		case elapsed >= 5:
			return 1
		default:
			return currentLevel
		}
	}

	switch {
	case elapsed >= float64(rule.Level4Timeout) && rule.Level4Timeout > 0:
		return 4
	case elapsed >= float64(rule.Level3Timeout) && rule.Level3Timeout > 0:
		return 3
	case elapsed >= float64(rule.Level2Timeout) && rule.Level2Timeout > 0:
		return 2
	case elapsed >= float64(rule.Level1Timeout):
		return 1
	default:
		return currentLevel
	}
}

// doEscalate 执行升级
func (s *AndonService) doEscalate(ctx context.Context, call *model.AndonCall, rule *model.AndonEscalationRule, toLevel int, escalationType string) error {
	fromLevel := call.CallLevel

	// 记录升级日志
	triggerReason := fmt.Sprintf("L%d超时%d分钟未处理", fromLevel, s.getTimeoutMinutes(rule, fromLevel))
	escalationLog := &model.AndonEscalationLog{
		CallID:         int64(call.ID),
		FromLevel:      fromLevel,
		ToLevel:        toLevel,
		EscalationType: escalationType,
		TriggerReason:  triggerReason,
		CreatedAt:      time.Now(),
	}
	if err := s.escalationRepo.Create(ctx, escalationLog); err != nil {
		s.logger.Error("创建升级日志失败", zap.Error(err))
	}

	// 发送通知
	go s.sendEscalationNotifications(call, toLevel, rule)

	// 更新呼叫
	updates := map[string]interface{}{
		"call_level":        toLevel,
		"is_escalated":      1,
		"escalated_at":      time.Now(),
		"escalation_count":  call.EscalationCount + 1,
	}

	// L3及以上自动变为HANDLING状态
	if toLevel >= 3 {
		updates["status"] = "HANDLING"
	}

	return s.andonRepo.Update(ctx, call.ID, updates)
}

// getTimeoutMinutes 获取等级超时分钟数
func (s *AndonService) getTimeoutMinutes(rule *model.AndonEscalationRule, level int) int {
	if rule == nil {
		switch level {
		case 1:
			return 5
		case 2:
			return 10
		case 3:
			return 15
		default:
			return 20
		}
	}

	switch level {
	case 1:
		return rule.Level1Timeout
	case 2:
		return rule.Level2Timeout
	case 3:
		return rule.Level3Timeout
	default:
		return rule.Level4Timeout
	}
}

// sendEscalationNotifications 发送升级通知
func (s *AndonService) sendEscalationNotifications(call *model.AndonCall, level int, rule *model.AndonEscalationRule) {
	var notifiers []model.NotifyInfo
	switch level {
	case 1:
		notifiers = rule.GetLevel1Notifiers()
	case 2:
		notifiers = rule.GetLevel2Notifiers()
	case 3:
		notifiers = rule.GetLevel3Notifiers()
	case 4:
		notifiers = rule.GetLevel4Notifiers()
	}

	// 构建通知内容
	title := fmt.Sprintf("安灯呼叫升级 - L%d", level)
	content := fmt.Sprintf("车间: %s\n工位: %s\n类型: %s\n描述: %s\n呼叫人: %s\n呼叫时间: %s",
		call.WorkshopName, getString(call.WorkstationName), call.AndonTypeName,
		call.Description, call.CallBy, call.CallTime.Format("2006-01-02 15:04:05"))

	// 记录通知日志
	for _, notify := range notifiers {
		notifyLog := &model.AndonNotificationLog{
			CallID:       int64(call.ID),
			Channel:      rule.Level1NotifyType,
			ReceiverType: "USER",
			ReceiverID:   notify.UserID,
			ReceiverName: notify.UserName,
			Title:        title,
			Content:      content,
			Priority:     level,
			SendTime:     time.Now(),
			SendResult:   "PENDING",
			TenantID:     call.TenantID,
		}
		if err := s.notifyRepo.Create(context.Background(), notifyLog); err != nil {
			s.logger.Error("创建通知日志失败", zap.Error(err))
		}
	}

	s.logger.Info("发送升级通知",
		zap.Int64("call_id", call.ID),
		zap.Int("level", level),
		zap.Int("notifier_count", len(notifiers)))
}

// GetStatistics 获取统计数据
func (s *AndonService) GetStatistics(ctx context.Context, tenantID, workshopID int64, startDate, endDate *time.Time) (map[string]interface{}, error) {
	return s.andonRepo.GetStatistics(ctx, tenantID, workshopID, startDate, endDate)
}

// ========== EscalationRule Service ==========

type EscalationRuleService struct {
	ruleRepo       *repository.EscalationRuleRepository
	escalationRepo *repository.EscalationLogRepository
	notifyRepo     *repository.NotificationLogRepository
}

func NewEscalationRuleService(ruleRepo *repository.EscalationRuleRepository, escalationRepo *repository.EscalationLogRepository, notifyRepo *repository.NotificationLogRepository) *EscalationRuleService {
	return &EscalationRuleService{
		ruleRepo:       ruleRepo,
		escalationRepo: escalationRepo,
		notifyRepo:     notifyRepo,
	}
}

// List 查询规则列表
func (s *EscalationRuleService) List(ctx context.Context, tenantID int64, andonType string, workshopID int64) ([]model.AndonEscalationRule, int64, error) {
	return s.ruleRepo.List(ctx, tenantID, andonType, workshopID)
}

// GetByID 根据ID获取
func (s *EscalationRuleService) GetByID(ctx context.Context, id int64) (*model.AndonEscalationRule, error) {
	return s.ruleRepo.GetByID(ctx, id)
}

// Create 创建规则
func (s *EscalationRuleService) Create(ctx context.Context, rule *model.AndonEscalationRule) error {
	// 如果设置为默认规则，先取消其他默认规则
	if rule.IsDefault == 1 {
		if err := s.ruleRepo.Update(ctx, 0, map[string]interface{}{"is_default": 0}); err != nil {
			log.Printf("取消其他默认规则失败: %v", err)
		}
	}
	return s.ruleRepo.Create(ctx, rule)
}

// Update 更新规则
func (s *EscalationRuleService) Update(ctx context.Context, id int64, rule *model.AndonEscalationRule) error {
	updates := map[string]interface{}{
		"rule_name":         rule.RuleName,
		"andon_type":        rule.AndonType,
		"workshop_id":       rule.WorkshopID,
		"priority_range":    rule.PriorityRange,
		"is_default":        rule.IsDefault,
		"level1_timeout":    rule.Level1Timeout,
		"level1_notify_type": rule.Level1NotifyType,
		"level1_notify_json": rule.Level1NotifyJSON,
		"level2_timeout":    rule.Level2Timeout,
		"level2_notify_type": rule.Level2NotifyType,
		"level2_notify_json": rule.Level2NotifyJSON,
		"level3_timeout":    rule.Level3Timeout,
		"level3_notify_type": rule.Level3NotifyType,
		"level3_notify_json": rule.Level3NotifyJSON,
		"level4_timeout":    rule.Level4Timeout,
		"level4_notify_type": rule.Level4NotifyType,
		"level4_notify_json": rule.Level4NotifyJSON,
		"escalation_mode":   rule.EscalationMode,
		"max_escalation_level": rule.MaxEscalationLevel,
		"audio_enabled":      rule.AudioEnabled,
		"audio_message_template": rule.AudioMessageTemplate,
		"audio_repeat_times": rule.AudioRepeatTimes,
		"is_enabled":        rule.IsEnabled,
		"sort_order":        rule.SortOrder,
		"remark":            rule.Remark,
	}

	// 如果设置为默认规则，先取消其他默认规则
	if rule.IsDefault == 1 {
		if err := s.ruleRepo.Update(ctx, 0, map[string]interface{}{"is_default": 0}); err != nil {
			log.Printf("取消其他默认规则失败: %v", err)
		}
	}

	return s.ruleRepo.Update(ctx, id, updates)
}

// Delete 删除规则
func (s *EscalationRuleService) Delete(ctx context.Context, id int64) error {
	return s.ruleRepo.Delete(ctx, id)
}

// GetEscalationLogs 获取升级历史
func (s *EscalationRuleService) GetEscalationLogs(ctx context.Context, callID int64) ([]model.AndonEscalationLog, error) {
	return s.escalationRepo.ListByCallID(ctx, callID)
}

// GetNotificationLogs 获取通知历史
func (s *EscalationRuleService) GetNotificationLogs(ctx context.Context, callID int64) ([]model.AndonNotificationLog, error) {
	return s.notifyRepo.ListByCallID(ctx, callID)
}

// Helper function
func getString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
