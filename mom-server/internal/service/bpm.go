package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type BPMService struct {
	processModelRepo *repository.ProcessModelRepository
	nodeRepo         *repository.NodeDefinitionRepository
	flowRepo         *repository.SequenceFlowRepository
	formRepo         *repository.FormDefinitionRepository
	fieldRepo        *repository.FormFieldRepository
	instanceRepo     *repository.ProcessInstanceRepository
	taskRepo         *repository.TaskInstanceRepository
	delegateRepo     *repository.DelegateRecordRepository
	approvalRepo     *repository.ApprovalRecordRepository
	userRepo         *repository.UserRepository
	roleRepo         *repository.RoleRepository
}

func NewBPMService(
	processModelRepo *repository.ProcessModelRepository,
	nodeRepo *repository.NodeDefinitionRepository,
	flowRepo *repository.SequenceFlowRepository,
	formRepo *repository.FormDefinitionRepository,
	fieldRepo *repository.FormFieldRepository,
	instanceRepo *repository.ProcessInstanceRepository,
	taskRepo *repository.TaskInstanceRepository,
	delegateRepo *repository.DelegateRecordRepository,
	approvalRepo *repository.ApprovalRecordRepository,
	userRepo *repository.UserRepository,
	roleRepo *repository.RoleRepository,
) *BPMService {
	return &BPMService{
		processModelRepo: processModelRepo,
		nodeRepo:        nodeRepo,
		flowRepo:        flowRepo,
		formRepo:        formRepo,
		fieldRepo:       fieldRepo,
		instanceRepo:    instanceRepo,
		taskRepo:        taskRepo,
		delegateRepo:    delegateRepo,
		approvalRepo:    approvalRepo,
		userRepo:        userRepo,
		roleRepo:        roleRepo,
	}
}

// ==================== 流程模型 ====================

func (s *BPMService) ListProcessModels(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.ProcessModel, int64, error) {
	return s.processModelRepo.List(ctx, tenantID, query)
}

func (s *BPMService) GetProcessModel(ctx context.Context, id string) (*model.ProcessModel, error) {
	var modelID uint
	_, err := fmt.Sscanf(id, "%d", &modelID)
	if err != nil {
		return nil, err
	}
	return s.processModelRepo.GetByID(ctx, modelID)
}

func (s *BPMService) CreateProcessModel(ctx context.Context, m *model.ProcessModel) error {
	return s.processModelRepo.Create(ctx, m)
}

func (s *BPMService) UpdateProcessModel(ctx context.Context, id string, m *model.ProcessModel) error {
	var modelID uint
	_, err := fmt.Sscanf(id, "%d", &modelID)
	if err != nil {
		return err
	}
	return s.processModelRepo.Update(ctx, modelID, map[string]interface{}{
		"model_name":  m.ModelName,
		"category":    m.Category,
		"description": m.Description,
		"form_type":   m.FormType,
		"form_url":    m.FormURL,
	})
}

func (s *BPMService) DeleteProcessModel(ctx context.Context, id string) error {
	var modelID uint
	_, err := fmt.Sscanf(id, "%d", &modelID)
	if err != nil {
		return err
	}
	return s.processModelRepo.Delete(ctx, modelID)
}

func (s *BPMService) PublishProcessModel(ctx context.Context, id string, publishedBy int64) error {
	var modelID uint
	_, err := fmt.Sscanf(id, "%d", &modelID)
	if err != nil {
		return err
	}
	now := time.Now()
	return s.processModelRepo.Update(ctx, modelID, map[string]interface{}{
		"is_published": 1,
		"published_at": now,
		"published_by": publishedBy,
	})
}

// ==================== 流程节点 ====================

func (s *BPMService) ListNodes(ctx context.Context, modelID string) ([]model.NodeDefinition, error) {
	var id uint
	_, err := fmt.Sscanf(modelID, "%d", &id)
	if err != nil {
		return nil, err
	}
	return s.nodeRepo.ListByModelID(ctx, id)
}

func (s *BPMService) CreateNode(ctx context.Context, node *model.NodeDefinition) error {
	return s.nodeRepo.Create(ctx, node)
}

func (s *BPMService) UpdateNode(ctx context.Context, id string, node *model.NodeDefinition) error {
	var nodeID uint
	_, err := fmt.Sscanf(id, "%d", &nodeID)
	if err != nil {
		return err
	}
	return s.nodeRepo.Update(ctx, nodeID, map[string]interface{}{
		"node_name":   node.NodeName,
		"node_type":   node.NodeType,
		"position_x":  node.PositionX,
		"position_y":  node.PositionY,
		"node_config": node.NodeConfig,
		"sort_order":  node.SortOrder,
	})
}

func (s *BPMService) DeleteNode(ctx context.Context, id string) error {
	var nodeID uint
	_, err := fmt.Sscanf(id, "%d", &nodeID)
	if err != nil {
		return err
	}
	return s.nodeRepo.Delete(ctx, nodeID)
}

// ==================== 流程连线 ====================

func (s *BPMService) ListFlows(ctx context.Context, modelID string) ([]model.SequenceFlow, error) {
	var id uint
	_, err := fmt.Sscanf(modelID, "%d", &id)
	if err != nil {
		return nil, err
	}
	return s.flowRepo.ListByModelID(ctx, id)
}

func (s *BPMService) CreateFlow(ctx context.Context, flow *model.SequenceFlow) error {
	return s.flowRepo.Create(ctx, flow)
}

func (s *BPMService) UpdateFlow(ctx context.Context, id string, flow *model.SequenceFlow) error {
	var flowID uint
	_, err := fmt.Sscanf(id, "%d", &flowID)
	if err != nil {
		return err
	}
	return s.flowRepo.Update(ctx, flowID, map[string]interface{}{
		"flow_name":            flow.FlowName,
		"source_node_id":       flow.SourceNodeID,
		"target_node_id":       flow.TargetNodeID,
		"condition_type":       flow.ConditionType,
		"condition_expression": flow.ConditionExpression,
		"is_default":           flow.IsDefault,
		"flow_config":          flow.FlowConfig,
		"sort_order":           flow.SortOrder,
	})
}

func (s *BPMService) DeleteFlow(ctx context.Context, id string) error {
	var flowID uint
	_, err := fmt.Sscanf(id, "%d", &flowID)
	if err != nil {
		return err
	}
	return s.flowRepo.Delete(ctx, flowID)
}

// ==================== 表单定义 ====================

func (s *BPMService) ListFormDefinitions(ctx context.Context, tenantID int64) ([]model.FormDefinition, int64, error) {
	return s.formRepo.List(ctx, tenantID)
}

func (s *BPMService) GetFormDefinition(ctx context.Context, id string) (*model.FormDefinition, error) {
	var formID uint
	_, err := fmt.Sscanf(id, "%d", &formID)
	if err != nil {
		return nil, err
	}
	return s.formRepo.GetByID(ctx, formID)
}

func (s *BPMService) CreateFormDefinition(ctx context.Context, form *model.FormDefinition) error {
	return s.formRepo.Create(ctx, form)
}

func (s *BPMService) UpdateFormDefinition(ctx context.Context, id string, form *model.FormDefinition) error {
	var formID uint
	_, err := fmt.Sscanf(id, "%d", &formID)
	if err != nil {
		return err
	}
	return s.formRepo.Update(ctx, formID, map[string]interface{}{
		"form_name": form.FormName,
		"category":  form.Category,
	})
}

func (s *BPMService) DeleteFormDefinition(ctx context.Context, id string) error {
	var formID uint
	_, err := fmt.Sscanf(id, "%d", &formID)
	if err != nil {
		return err
	}
	return s.formRepo.Delete(ctx, formID)
}

// ==================== 表单字段 ====================

func (s *BPMService) ListFormFields(ctx context.Context, formID string) ([]model.FormField, error) {
	var id uint
	_, err := fmt.Sscanf(formID, "%d", &id)
	if err != nil {
		return nil, err
	}
	return s.fieldRepo.ListByFormID(ctx, id)
}

func (s *BPMService) CreateFormField(ctx context.Context, field *model.FormField) error {
	return s.fieldRepo.Create(ctx, field)
}

func (s *BPMService) UpdateFormField(ctx context.Context, id string, field *model.FormField) error {
	var fieldID uint
	_, err := fmt.Sscanf(id, "%d", &fieldID)
	if err != nil {
		return err
	}
	return s.fieldRepo.Update(ctx, fieldID, map[string]interface{}{
		"field_name":    field.FieldName,
		"field_type":    field.FieldType,
		"field_config":  field.FieldConfig,
		"default_value": field.DefaultValue,
		"is_required":   field.IsRequired,
		"is_readonly":   field.IsReadonly,
		"is_hidden":     field.IsHidden,
		"sort_order":    field.SortOrder,
	})
}

func (s *BPMService) DeleteFormField(ctx context.Context, id string) error {
	var fieldID uint
	_, err := fmt.Sscanf(id, "%d", &fieldID)
	if err != nil {
		return err
	}
	return s.fieldRepo.Delete(ctx, fieldID)
}

// ==================== 流程实例 ====================

func (s *BPMService) ListProcessInstances(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.ProcessInstance, int64, error) {
	return s.instanceRepo.List(ctx, tenantID, query)
}

func (s *BPMService) GetProcessInstance(ctx context.Context, id string) (*model.ProcessInstance, error) {
	var instanceID uint
	_, err := fmt.Sscanf(id, "%d", &instanceID)
	if err != nil {
		return nil, err
	}
	return s.instanceRepo.GetByID(ctx, instanceID)
}

func (s *BPMService) CreateProcessInstance(ctx context.Context, instance *model.ProcessInstance) error {
	return s.instanceRepo.Create(ctx, instance)
}

func (s *BPMService) CancelProcessInstance(ctx context.Context, id string) error {
	var instanceID uint
	_, err := fmt.Sscanf(id, "%d", &instanceID)
	if err != nil {
		return err
	}
	return s.instanceRepo.Update(ctx, instanceID, map[string]interface{}{
		"status": "CANCELLED",
	})
}

func (s *BPMService) TerminateProcessInstance(ctx context.Context, id string) error {
	var instanceID uint
	_, err := fmt.Sscanf(id, "%d", &instanceID)
	if err != nil {
		return err
	}
	return s.instanceRepo.Update(ctx, instanceID, map[string]interface{}{
		"status": "TERMINATED",
	})
}

// ==================== 任务实例 ====================

func (s *BPMService) ListTasksByAssignee(ctx context.Context, assigneeID int64, query map[string]interface{}) ([]model.TaskInstance, int64, error) {
	return s.taskRepo.ListByAssignee(ctx, assigneeID, query)
}

func (s *BPMService) GetTask(ctx context.Context, id string) (*model.TaskInstance, error) {
	var taskID uint
	_, err := fmt.Sscanf(id, "%d", &taskID)
	if err != nil {
		return nil, err
	}
	return s.taskRepo.GetByID(ctx, taskID)
}

func (s *BPMService) ApproveTask(ctx context.Context, id string, approverID int64, approverName, comment string) error {
	var taskID uint
	_, err := fmt.Sscanf(id, "%d", &taskID)
	if err != nil {
		return err
	}
	now := time.Now()
	return s.taskRepo.Update(ctx, taskID, map[string]interface{}{
		"status":         "COMPLETED",
		"action_result":  "AGREE",
		"action_comment": comment,
		"action_time":    now,
		"completed_at":   now,
	})
}

func (s *BPMService) RejectTask(ctx context.Context, id string, approverID int64, approverName, comment string) error {
	var taskID uint
	_, err := fmt.Sscanf(id, "%d", &taskID)
	if err != nil {
		return err
	}
	now := time.Now()
	return s.taskRepo.Update(ctx, taskID, map[string]interface{}{
		"status":         "COMPLETED",
		"action_result":  "REJECT",
		"action_comment": comment,
		"action_time":    now,
		"completed_at":   now,
	})
}

// ==================== 委托记录 ====================

func (s *BPMService) ListDelegates(ctx context.Context, delegateID int64) ([]model.DelegateRecord, error) {
	return s.delegateRepo.ListByDelegate(ctx, delegateID)
}

func (s *BPMService) CreateDelegate(ctx context.Context, record *model.DelegateRecord) error {
	return s.delegateRepo.Create(ctx, record)
}

func (s *BPMService) UpdateDelegate(ctx context.Context, id string, record *model.DelegateRecord) error {
	var delegateID uint
	_, err := fmt.Sscanf(id, "%d", &delegateID)
	if err != nil {
		return err
	}
	return s.delegateRepo.Update(ctx, delegateID, map[string]interface{}{
		"delegatee_id": record.DelegateeID,
		"start_date":   record.StartDate,
		"end_date":     record.EndDate,
		"biz_types":    record.BizTypes,
		"is_active":    record.IsActive,
	})
}

func (s *BPMService) DeleteDelegate(ctx context.Context, id string) error {
	var delegateID uint
	_, err := fmt.Sscanf(id, "%d", &delegateID)
	if err != nil {
		return err
	}
	return s.delegateRepo.Delete(ctx, delegateID)
}

// ==================== 审批记录 ====================

func (s *BPMService) ListApprovalRecords(ctx context.Context, taskID string) ([]model.ApprovalRecord, error) {
	var id uint
	_, err := fmt.Sscanf(taskID, "%d", &id)
	if err != nil {
		return nil, err
	}
	return s.approvalRepo.ListByTaskID(ctx, id)
}

