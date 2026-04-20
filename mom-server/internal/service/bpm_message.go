package service

import (
	"context"
	"fmt"
	"log"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// BpmMessageService 流程消息服务
type BpmMessageService struct {
	instanceRepo *repository.ProcessInstanceRepository
	taskRepo    *repository.TaskInstanceRepository
	eventBus    *EventBus
}

// NewBpmMessageService creates a new BPM message service
func NewBpmMessageService(
	instanceRepo *repository.ProcessInstanceRepository,
	taskRepo *repository.TaskInstanceRepository,
) *BpmMessageService {
	return &BpmMessageService{
		instanceRepo: instanceRepo,
		taskRepo:    taskRepo,
		eventBus:    GetEventBus(),
	}
}

// SendTaskCreatedMessage 发送任务创建消息
func (s *BpmMessageService) SendTaskCreatedMessage(ctx context.Context, taskID int64, taskName string, assignee string) error {
	event := NewDomainEvent("bpm:task:created", 0, map[string]interface{}{
		"task_id":    taskID,
		"task_name":  taskName,
		"assignee":   assignee,
	})

	if s.eventBus != nil {
		s.eventBus.Publish(ctx, event)
	}

	log.Printf("[BPM Message] Task created: taskID=%d, taskName=%s, assignee=%s", taskID, taskName, assignee)
	return nil
}

// SendTaskCompletedMessage 发送任务完成消息
func (s *BpmMessageService) SendTaskCompletedMessage(ctx context.Context, taskID int64, processInstanceID int64, comment string) error {
	event := NewDomainEvent("bpm:task:completed", 0, map[string]interface{}{
		"task_id":            taskID,
		"process_instance_id": processInstanceID,
		"comment":            comment,
	})

	if s.eventBus != nil {
		s.eventBus.Publish(ctx, event)
	}

	log.Printf("[BPM Message] Task completed: taskID=%d, processInstanceID=%d, comment=%s", taskID, processInstanceID, comment)
	return nil
}

// SendProcessCompletedMessage 发送流程完成消息
func (s *BpmMessageService) SendProcessCompletedMessage(ctx context.Context, processInstanceID int64, businessKey string) error {
	event := NewDomainEvent("bpm:process:completed", 0, map[string]interface{}{
		"process_instance_id": processInstanceID,
		"business_key":       businessKey,
	})

	if s.eventBus != nil {
		s.eventBus.Publish(ctx, event)
	}

	log.Printf("[BPM Message] Process completed: processInstanceID=%d, businessKey=%s", processInstanceID, businessKey)
	return nil
}

// SendTaskAssignedMessage 发送任务分配消息
func (s *BpmMessageService) SendTaskAssignedMessage(ctx context.Context, taskID int64, taskName string, assignee string, assigneeName string) error {
	event := NewDomainEvent("bpm:task:assigned", 0, map[string]interface{}{
		"task_id":        taskID,
		"task_name":      taskName,
		"assignee":       assignee,
		"assignee_name":  assigneeName,
	})

	if s.eventBus != nil {
		s.eventBus.Publish(ctx, event)
	}

	log.Printf("[BPM Message] Task assigned: taskID=%d, taskName=%s, assignee=%s, assigneeName=%s",
		taskID, taskName, assignee, assigneeName)
	return nil
}

// SendProcessCanceledMessage 发送流程取消消息
func (s *BpmMessageService) SendProcessCanceledMessage(ctx context.Context, processInstanceID int64, businessKey string, cancelReason string) error {
	event := NewDomainEvent("bpm:process:canceled", 0, map[string]interface{}{
		"process_instance_id": processInstanceID,
		"business_key":        businessKey,
		"cancel_reason":       cancelReason,
	})

	if s.eventBus != nil {
		s.eventBus.Publish(ctx, event)
	}

	log.Printf("[BPM Message] Process canceled: processInstanceID=%d, businessKey=%s, reason=%s",
		processInstanceID, businessKey, cancelReason)
	return nil
}

// SendProcessTerminatedMessage 发送流程终止消息
func (s *BpmMessageService) SendProcessTerminatedMessage(ctx context.Context, processInstanceID int64, businessKey string, terminateReason string) error {
	event := NewDomainEvent("bpm:process:terminated", 0, map[string]interface{}{
		"process_instance_id": processInstanceID,
		"business_key":        businessKey,
		"terminate_reason":    terminateReason,
	})

	if s.eventBus != nil {
		s.eventBus.Publish(ctx, event)
	}

	log.Printf("[BPM Message] Process terminated: processInstanceID=%d, businessKey=%s, reason=%s",
		processInstanceID, businessKey, terminateReason)
	return nil
}

// SendTaskRejectedMessage 发送任务驳回消息
func (s *BpmMessageService) SendTaskRejectedMessage(ctx context.Context, taskID int64, processInstanceID int64, rejectedBy string, comment string) error {
	event := NewDomainEvent("bpm:task:rejected", 0, map[string]interface{}{
		"task_id":             taskID,
		"process_instance_id": processInstanceID,
		"rejected_by":        rejectedBy,
		"comment":             comment,
	})

	if s.eventBus != nil {
		s.eventBus.Publish(ctx, event)
	}

	log.Printf("[BPM Message] Task rejected: taskID=%d, processInstanceID=%d, rejectedBy=%s, comment=%s",
		taskID, processInstanceID, rejectedBy, comment)
	return nil
}

// ==================== 跨模块API ====================

// StartProcessReq 启动流程请求
type StartProcessReq struct {
	ProcessDefKey string                 `json:"processDefKey" binding:"required"`
	BusinessKey  string                 `json:"businessKey"`
	Variables    map[string]interface{} `json:"variables"`
}

// CompleteTaskReq 完成任务请求
type CompleteTaskReq struct {
	TaskID    int64                  `json:"taskId" binding:"required"`
	Variables map[string]interface{} `json:"variables"`
}

// ProcessInstanceRespVO 流程实例响应
type ProcessInstanceRespVO struct {
	Id           int64  `json:"id"`
	InstanceNo   string `json:"instanceNo"`
	ModelCode    string `json:"modelCode"`
	ModelName    string `json:"modelName"`
	Status       string `json:"status"`
	StartTime    string `json:"startTime"`
	EndTime      string `json:"endTime"`
}

// BpmProcessInstanceApi 跨模块流程实例API
type BpmProcessInstanceApi struct {
	bpmSvc *BPMService
}

// NewBpmProcessInstanceApi creates a new BPM process instance API
func NewBpmProcessInstanceApi(bpmSvc *BPMService) *BpmProcessInstanceApi {
	return &BpmProcessInstanceApi{
		bpmSvc: bpmSvc,
	}
}

// StartProcessInstance 启动流程实例
func (s *BpmProcessInstanceApi) StartProcessInstance(ctx context.Context, req *StartProcessReq) (*ProcessInstanceRespVO, error) {
	if req.ProcessDefKey == "" {
		return nil, fmt.Errorf("process definition key is required")
	}

	// Create process instance using BPM service
	instance := &model.ProcessInstance{
		ModelCode: &req.ProcessDefKey,
		BizNo:     &req.BusinessKey,
		Title:     fmt.Sprintf("Process-%s-%s", req.ProcessDefKey, req.BusinessKey),
		Status:    "RUNNING",
	}
	if err := s.bpmSvc.CreateProcessInstance(ctx, instance); err != nil {
		return nil, fmt.Errorf("failed to start process instance: %w", err)
	}

	resp := &ProcessInstanceRespVO{
		Id:         int64(instance.ID),
		InstanceNo: instance.InstanceNo,
		Status:     instance.Status,
	}
	if instance.CompletedAt != nil {
		resp.EndTime = instance.CompletedAt.Format("2006-01-02 15:04:05")
	}

	return resp, nil
}

// CompleteTask 完成任务
func (s *BpmProcessInstanceApi) CompleteTask(ctx context.Context, taskID int64, variables map[string]interface{}) error {
	if taskID <= 0 {
		return fmt.Errorf("task ID is required")
	}

	// Get task to find the process instance
	task, err := s.bpmSvc.GetTask(ctx, fmt.Sprintf("%d", taskID))
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}

	// Use ApproveTask to complete the task
	if err := s.bpmSvc.ApproveTask(ctx, fmt.Sprintf("%d", taskID), 0, "system", ""); err != nil {
		return fmt.Errorf("failed to complete task: %w", err)
	}

	// Send completion message
	s.sendTaskCompletedMessage(ctx, taskID, task.InstanceID, "")

	return nil
}

// GetProcessInstance 获取流程实例
func (s *BpmProcessInstanceApi) GetProcessInstance(ctx context.Context, instanceID int64) (*ProcessInstanceRespVO, error) {
	if instanceID <= 0 {
		return nil, fmt.Errorf("instance ID is required")
	}

	instance, err := s.bpmSvc.GetProcessInstance(ctx, fmt.Sprintf("%d", instanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get process instance: %w", err)
	}

	resp := &ProcessInstanceRespVO{
		Id:         int64(instance.ID),
		InstanceNo: instance.InstanceNo,
		Status:    instance.Status,
	}
	if instance.ModelCode != nil {
		resp.ModelCode = *instance.ModelCode
	}
	if instance.ModelName != nil {
		resp.ModelName = *instance.ModelName
	}
	if instance.CompletedAt != nil {
		resp.EndTime = instance.CompletedAt.Format("2006-01-02 15:04:05")
	}

	return resp, nil
}

// GetTask 获取任务实例
func (s *BpmProcessInstanceApi) GetTask(ctx context.Context, taskID int64) (*model.TaskInstance, error) {
	return s.bpmSvc.GetTask(ctx, fmt.Sprintf("%d", taskID))
}

// sendTaskCompletedMessage sends task completed message (internal helper)
func (s *BpmProcessInstanceApi) sendTaskCompletedMessage(ctx context.Context, taskID, instanceID int64, comment string) {
	// This would use the BpmMessageService in production
	log.Printf("[BPM Instance API] Task completed internally: taskID=%d, instanceID=%d", taskID, instanceID)
}
