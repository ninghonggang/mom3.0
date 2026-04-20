package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// ProcessModelRepository 流程模型仓库
type ProcessModelRepository struct {
	db *gorm.DB
}

func NewProcessModelRepository(db *gorm.DB) *ProcessModelRepository {
	return &ProcessModelRepository{db: db}
}

func (r *ProcessModelRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.ProcessModel, int64, error) {
	var list []model.ProcessModel
	var total int64

	q := r.db.WithContext(ctx).Model(&model.ProcessModel{}).Where("tenant_id = ?", tenantID)

	if modelType, ok := query["model_type"]; ok && modelType != "" {
		q = q.Where("model_type = ?", modelType)
	}
	if category, ok := query["category"]; ok && category != "" {
		q = q.Where("category = ?", category)
	}

	q.Count(&total)
	q = q.Preload("Nodes").Preload("Flows").Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *ProcessModelRepository) GetByID(ctx context.Context, id uint) (*model.ProcessModel, error) {
	var model model.ProcessModel
	err := r.db.WithContext(ctx).Preload("Nodes").Preload("Flows").First(&model, id).Error
	return &model, err
}

func (r *ProcessModelRepository) GetByCode(ctx context.Context, tenantID int64, code string) (*model.ProcessModel, error) {
	var m model.ProcessModel
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND model_code = ?", tenantID, code).First(&m).Error
	return &m, err
}

func (r *ProcessModelRepository) Create(ctx context.Context, m *model.ProcessModel) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *ProcessModelRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ProcessModel{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ProcessModelRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		tx.Where("model_id = ?", id).Delete(&model.NodeDefinition{})
		tx.Where("model_id = ?", id).Delete(&model.SequenceFlow{})
		return tx.Delete(&model.ProcessModel{}, id).Error
	})
}

// NodeDefinitionRepository 流程节点仓库
type NodeDefinitionRepository struct {
	db *gorm.DB
}

func NewNodeDefinitionRepository(db *gorm.DB) *NodeDefinitionRepository {
	return &NodeDefinitionRepository{db: db}
}

func (r *NodeDefinitionRepository) ListByModelID(ctx context.Context, modelID uint) ([]model.NodeDefinition, error) {
	var list []model.NodeDefinition
	err := r.db.WithContext(ctx).Where("model_id = ?", modelID).Order("sort_order ASC").Find(&list).Error
	return list, err
}

func (r *NodeDefinitionRepository) Create(ctx context.Context, node *model.NodeDefinition) error {
	return r.db.WithContext(ctx).Create(node).Error
}

func (r *NodeDefinitionRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.NodeDefinition{}).Where("id = ?", id).Updates(updates).Error
}

func (r *NodeDefinitionRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.NodeDefinition{}, id).Error
}

// SequenceFlowRepository 流程连线仓库
type SequenceFlowRepository struct {
	db *gorm.DB
}

func NewSequenceFlowRepository(db *gorm.DB) *SequenceFlowRepository {
	return &SequenceFlowRepository{db: db}
}

func (r *SequenceFlowRepository) ListByModelID(ctx context.Context, modelID uint) ([]model.SequenceFlow, error) {
	var list []model.SequenceFlow
	err := r.db.WithContext(ctx).Where("model_id = ?", modelID).Order("sort_order ASC").Find(&list).Error
	return list, err
}

func (r *SequenceFlowRepository) Create(ctx context.Context, flow *model.SequenceFlow) error {
	return r.db.WithContext(ctx).Create(flow).Error
}

func (r *SequenceFlowRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.SequenceFlow{}).Where("id = ?", id).Updates(updates).Error
}

func (r *SequenceFlowRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.SequenceFlow{}, id).Error
}

// FormDefinitionRepository 表单定义仓库
type FormDefinitionRepository struct {
	db *gorm.DB
}

func NewFormDefinitionRepository(db *gorm.DB) *FormDefinitionRepository {
	return &FormDefinitionRepository{db: db}
}

func (r *FormDefinitionRepository) List(ctx context.Context, tenantID int64) ([]model.FormDefinition, int64, error) {
	var list []model.FormDefinition
	var total int64
	err := r.db.WithContext(ctx).Model(&model.FormDefinition{}).Where("tenant_id = ?", tenantID).Count(&total).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *FormDefinitionRepository) GetByID(ctx context.Context, id uint) (*model.FormDefinition, error) {
	var form model.FormDefinition
	err := r.db.WithContext(ctx).Preload("Fields").First(&form, id).Error
	return &form, err
}

func (r *FormDefinitionRepository) Create(ctx context.Context, form *model.FormDefinition) error {
	return r.db.WithContext(ctx).Create(form).Error
}

func (r *FormDefinitionRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.FormDefinition{}).Where("id = ?", id).Updates(updates).Error
}

func (r *FormDefinitionRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		tx.Where("form_id = ?", id).Delete(&model.FormField{})
		return tx.Delete(&model.FormDefinition{}, id).Error
	})
}

// FormFieldRepository 表单字段仓库
type FormFieldRepository struct {
	db *gorm.DB
}

func NewFormFieldRepository(db *gorm.DB) *FormFieldRepository {
	return &FormFieldRepository{db: db}
}

func (r *FormFieldRepository) ListByFormID(ctx context.Context, formID uint) ([]model.FormField, error) {
	var list []model.FormField
	err := r.db.WithContext(ctx).Where("form_id = ?", formID).Order("sort_order ASC").Find(&list).Error
	return list, err
}

func (r *FormFieldRepository) Create(ctx context.Context, field *model.FormField) error {
	return r.db.WithContext(ctx).Create(field).Error
}

func (r *FormFieldRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.FormField{}).Where("id = ?", id).Updates(updates).Error
}

func (r *FormFieldRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.FormField{}, id).Error
}

// ProcessInstanceRepository 流程实例仓库
type ProcessInstanceRepository struct {
	db *gorm.DB
}

func NewProcessInstanceRepository(db *gorm.DB) *ProcessInstanceRepository {
	return &ProcessInstanceRepository{db: db}
}

func (r *ProcessInstanceRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.ProcessInstance, int64, error) {
	var list []model.ProcessInstance
	var total int64

	q := r.db.WithContext(ctx).Model(&model.ProcessInstance{}).Where("tenant_id = ?", tenantID)

	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}
	if initiatorID, ok := query["initiator_id"]; ok {
		q = q.Where("initiator_id = ?", initiatorID)
	}

	q.Count(&total)
	q = q.Preload("Tasks").Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *ProcessInstanceRepository) GetByID(ctx context.Context, id uint) (*model.ProcessInstance, error) {
	var instance model.ProcessInstance
	err := r.db.WithContext(ctx).Preload("Tasks").First(&instance, id).Error
	return &instance, err
}

func (r *ProcessInstanceRepository) Create(ctx context.Context, instance *model.ProcessInstance) error {
	return r.db.WithContext(ctx).Create(instance).Error
}

func (r *ProcessInstanceRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ProcessInstance{}).Where("id = ?", id).Updates(updates).Error
}

// TaskInstanceRepository 任务实例仓库
type TaskInstanceRepository struct {
	db *gorm.DB
}

func NewTaskInstanceRepository(db *gorm.DB) *TaskInstanceRepository {
	return &TaskInstanceRepository{db: db}
}

func (r *TaskInstanceRepository) ListByAssignee(ctx context.Context, assigneeID int64, query map[string]interface{}) ([]model.TaskInstance, int64, error) {
	var list []model.TaskInstance
	var total int64

	q := r.db.WithContext(ctx).Model(&model.TaskInstance{}).Where("assignee_id = ? AND status IN ('PENDING', 'IN_PROGRESS')", assigneeID)

	q.Count(&total)
	q = q.Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *TaskInstanceRepository) ListByInstance(ctx context.Context, instanceID uint) ([]model.TaskInstance, error) {
	var list []model.TaskInstance
	err := r.db.WithContext(ctx).Where("instance_id = ?", instanceID).Order("created_at ASC").Find(&list).Error
	return list, err
}

func (r *TaskInstanceRepository) GetByID(ctx context.Context, id uint) (*model.TaskInstance, error) {
	var task model.TaskInstance
	err := r.db.WithContext(ctx).First(&task, id).Error
	return &task, err
}

func (r *TaskInstanceRepository) Create(ctx context.Context, task *model.TaskInstance) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *TaskInstanceRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.TaskInstance{}).Where("id = ?", id).Updates(updates).Error
}

// GetDB 获取数据库连接
func (r *TaskInstanceRepository) GetDB() *gorm.DB {
	return r.db
}

// DelegateRecordRepository 流程委托记录仓库
type DelegateRecordRepository struct {
	db *gorm.DB
}

func NewDelegateRecordRepository(db *gorm.DB) *DelegateRecordRepository {
	return &DelegateRecordRepository{db: db}
}

func (r *DelegateRecordRepository) ListByDelegate(ctx context.Context, delegateID int64) ([]model.DelegateRecord, error) {
	var list []model.DelegateRecord
	err := r.db.WithContext(ctx).Where("delegate_id = ? AND is_active = 1", delegateID).Order("created_at DESC").Find(&list).Error
	return list, err
}

func (r *DelegateRecordRepository) Create(ctx context.Context, record *model.DelegateRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *DelegateRecordRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.DelegateRecord{}).Where("id = ?", id).Updates(updates).Error
}

func (r *DelegateRecordRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.DelegateRecord{}, id).Error
}

// ApprovalRecordRepository 审批记录仓库
type ApprovalRecordRepository struct {
	db *gorm.DB
}

func NewApprovalRecordRepository(db *gorm.DB) *ApprovalRecordRepository {
	return &ApprovalRecordRepository{db: db}
}

func (r *ApprovalRecordRepository) ListByTaskID(ctx context.Context, taskID uint) ([]model.ApprovalRecord, error) {
	var list []model.ApprovalRecord
	err := r.db.WithContext(ctx).Where("task_id = ?", taskID).Order("action_time ASC").Find(&list).Error
	return list, err
}

func (r *ApprovalRecordRepository) Create(ctx context.Context, record *model.ApprovalRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}
