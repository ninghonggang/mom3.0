package listener

import (
	"context"
	"log"
	"sync"
)

// BpmProcessInstanceResultEventListener 流程实例审批结果监听器
type BpmProcessInstanceResultEventListener struct {
	// Handlers for approved/rejected events
	handlers []BpmResultHandler
	mu       sync.RWMutex
}

// BpmResultHandler 审批结果处理接口
type BpmResultHandler interface {
	OnApproved(instanceID int64, businessKey string) error
	OnRejected(instanceID int64, businessKey string, comment string) error
}

// BpmResultEvent 审批结果事件
type BpmResultEvent struct {
	InstanceID  int64
	BusinessKey string
	Result      string // APPROVED or REJECTED
	Comment     string
}

// NewBpmProcessInstanceResultEventListener creates a new BPM result listener
func NewBpmProcessInstanceResultEventListener() *BpmProcessInstanceResultEventListener {
	return &BpmProcessInstanceResultEventListener{
		handlers: make([]BpmResultHandler, 0),
	}
}

// RegisterHandler 注册审批结果处理器
func (l *BpmProcessInstanceResultEventListener) RegisterHandler(handler BpmResultHandler) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.handlers = append(l.handlers, handler)
}

// OnApproved 审批通过处理
func (l *BpmProcessInstanceResultEventListener) OnApproved(instanceID int64, businessKey string) error {
	l.mu.RLock()
	handlers := make([]BpmResultHandler, len(l.handlers))
	copy(handlers, l.handlers)
	l.mu.RUnlock()

	var lastErr error
	for _, h := range handlers {
		if err := h.OnApproved(instanceID, businessKey); err != nil {
			log.Printf("[BPM Result Listener] OnApproved handler error: %v", err)
			lastErr = err
		}
	}

	if lastErr != nil {
		return lastErr
	}

	log.Printf("[BPM Result Listener] Instance approved: instanceID=%d, businessKey=%s, handlers=%d",
		instanceID, businessKey, len(handlers))
	return nil
}

// OnRejected 审批驳回处理
func (l *BpmProcessInstanceResultEventListener) OnRejected(instanceID int64, businessKey string, comment string) error {
	l.mu.RLock()
	handlers := make([]BpmResultHandler, len(l.handlers))
	copy(handlers, l.handlers)
	l.mu.RUnlock()

	var lastErr error
	for _, h := range handlers {
		if err := h.OnRejected(instanceID, businessKey, comment); err != nil {
			log.Printf("[BPM Result Listener] OnRejected handler error: %v", err)
			lastErr = err
		}
	}

	if lastErr != nil {
		return lastErr
	}

	log.Printf("[BPM Result Listener] Instance rejected: instanceID=%d, businessKey=%s, comment=%s, handlers=%d",
		instanceID, businessKey, comment, len(handlers))
	return nil
}

// HandleEvent 处理审批结果事件
func (l *BpmProcessInstanceResultEventListener) HandleEvent(event *BpmResultEvent) error {
	if event == nil {
		return nil
	}

	switch event.Result {
	case "APPROVED":
		return l.OnApproved(event.InstanceID, event.BusinessKey)
	case "REJECTED":
		return l.OnRejected(event.InstanceID, event.BusinessKey, event.Comment)
	default:
		log.Printf("[BPM Result Listener] Unknown result: %s", event.Result)
		return nil
	}
}

// ==================== 默认实现 ====================

// DefaultBpmResultHandler 默认审批结果处理器（空实现，可被扩展）
type DefaultBpmResultHandler struct{}

// OnApproved 默认通过处理
func (h *DefaultBpmResultHandler) OnApproved(instanceID int64, businessKey string) error {
	log.Printf("[DefaultBpmResultHandler] Instance approved: instanceID=%d, businessKey=%s", instanceID, businessKey)
	return nil
}

// OnRejected 默认驳回处理
func (h *DefaultBpmResultHandler) OnRejected(instanceID int64, businessKey string, comment string) error {
	log.Printf("[DefaultBpmResultHandler] Instance rejected: instanceID=%d, businessKey=%s, comment=%s",
		instanceID, businessKey, comment)
	return nil
}

// ==================== 事件订阅集成 ====================

// BpmEventSubscriber BPM事件订阅器
type BpmEventSubscriber struct {
	listener *BpmProcessInstanceResultEventListener
	eventBus interface {
		Subscribe(eventType string, handler func(ctx context.Context, event interface{}))
	}
}

// NewBpmEventSubscriber 创建BPM事件订阅器
func NewBpmEventSubscriber(listener *BpmProcessInstanceResultEventListener) *BpmEventSubscriber {
	return &BpmEventSubscriber{
		listener: listener,
	}
}

// SubscribeToEvents 订阅BPM事件
func (s *BpmEventSubscriber) SubscribeToEvents(eventBus interface {
	Subscribe(eventType string, handler func(ctx context.Context, event interface{}))
}) {
	s.eventBus = eventBus

	// 订阅任务完成事件
	if s.eventBus != nil {
		s.eventBus.Subscribe("bpm:task:completed", s.handleTaskCompleted)
		s.eventBus.Subscribe("bpm:task:rejected", s.handleTaskRejected)
		s.eventBus.Subscribe("bpm:process:completed", s.handleProcessCompleted)
	}
}

// handleTaskCompleted 处理任务完成事件
func (s *BpmEventSubscriber) handleTaskCompleted(ctx context.Context, event interface{}) {
	data, ok := event.(map[string]interface{})
	if !ok {
		log.Printf("[BpmEventSubscriber] Invalid task completed event data")
		return
	}

	taskID, _ := data["task_id"].(float64)
	processInstanceID, _ := data["process_instance_id"].(float64)

	log.Printf("[BpmEventSubscriber] Task completed: taskID=%d, processInstanceID=%d",
		int64(taskID), int64(processInstanceID))
}

// handleTaskRejected 处理任务驳回事件
func (s *BpmEventSubscriber) handleTaskRejected(ctx context.Context, event interface{}) {
	data, ok := event.(map[string]interface{})
	if !ok {
		log.Printf("[BpmEventSubscriber] Invalid task rejected event data")
		return
	}

	taskID, _ := data["task_id"].(float64)
	processInstanceID, _ := data["process_instance_id"].(float64)
	comment, _ := data["comment"].(string)

	log.Printf("[BpmEventSubscriber] Task rejected: taskID=%d, processInstanceID=%d, comment=%s",
		int64(taskID), int64(processInstanceID), comment)
}

// handleProcessCompleted 处理流程完成事件
func (s *BpmEventSubscriber) handleProcessCompleted(ctx context.Context, event interface{}) {
	data, ok := event.(map[string]interface{})
	if !ok {
		log.Printf("[BpmEventSubscriber] Invalid process completed event data")
		return
	}

	processInstanceID, _ := data["process_instance_id"].(float64)
	businessKey, _ := data["business_key"].(string)

	// 触发审批通过事件
	if err := s.listener.OnApproved(int64(processInstanceID), businessKey); err != nil {
		log.Printf("[BpmEventSubscriber] Failed to handle approval: %v", err)
	}
}
