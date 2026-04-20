package service

import (
	"context"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type IntegrationService struct {
	configRepo *repository.InterfaceConfigRepository
	logRepo   *repository.InterfaceExecutionLogRepository
}

func NewIntegrationService(
	configRepo *repository.InterfaceConfigRepository,
	logRepo *repository.InterfaceExecutionLogRepository,
) *IntegrationService {
	return &IntegrationService{
		configRepo: configRepo,
		logRepo:    logRepo,
	}
}

// ListConfigs 分页查询接口配置
func (s *IntegrationService) ListConfigs(ctx context.Context, q *model.InterfaceConfigQuery) ([]model.InterfaceConfig, int64, error) {
	return s.configRepo.List(ctx, q)
}

// GetConfig 获取单个配置（含字段映射和触发器）
func (s *IntegrationService) GetConfig(ctx context.Context, id int64) (*model.InterfaceConfig, []model.InterfaceFieldMap, []model.InterfaceTrigger, error) {
	return s.configRepo.GetWithRelations(ctx, id)
}

// CreateConfig 创建接口配置
func (s *IntegrationService) CreateConfig(ctx context.Context, tenantID int64, req *model.InterfaceConfigCreate) (*model.InterfaceConfig, error) {
	// Check code uniqueness
	existing, _ := s.configRepo.GetByCode(ctx, req.Code)
	if existing != nil {
		return nil, nil
	}

	cfg := &model.InterfaceConfig{
		TenantID:    tenantID,
		Name:        req.Name,
		Code:        req.Code,
		Category:    req.Category,
		Description: req.Description,
		Direction:   req.Direction,
		Method:      req.Method,
		BaseURL:     req.BaseURL,
		Path:        req.Path,
		AuthType:    req.AuthType,
		AuthConfig:  req.AuthConfig,
		RequestContentType:  req.RequestContentType,
		RequestBodyTemplate: req.RequestBodyTemplate,
		ResponseFormat:      req.ResponseFormat,
		Timeout:        req.Timeout,
		RetryCount:     req.RetryCount,
		RetryInterval:  req.RetryInterval,
		SourceType:    req.SourceType,
		SourceTable:   req.SourceTable,
		SourceAPI:     req.SourceAPI,
		SourceFilter:  req.SourceFilter,
		SourceFields:  req.SourceFields,
		PrimaryKey:    req.PrimaryKey,
		IncrementalField: req.IncrementalField,
		IncrementalWindow: req.IncrementalWindow,
		Status: "ENABLE",
		Remark: req.Remark,
	}

	if err := s.configRepo.Create(ctx, cfg); err != nil {
		return nil, err
	}

	// Create field maps
	if len(req.FieldMaps) > 0 {
		maps := make([]model.InterfaceFieldMap, len(req.FieldMaps))
		for i, m := range req.FieldMaps {
			maps[i] = model.InterfaceFieldMap{
				InterfaceConfigID: cfg.ID,
				FieldName:     m.FieldName,
				FieldType:    m.FieldType,
				MapType:      m.MapType,
				MapValue:     m.MapValue,
				Required:     m.Required,
				DefaultValue: m.DefaultValue,
				TransformFunc: m.TransformFunc,
				SortOrder:   m.SortOrder,
			}
		}
		s.configRepo.UpsertFieldMaps(ctx, cfg.ID, maps)
	}

	// Create triggers
	if len(req.Triggers) > 0 {
		triggers := make([]model.InterfaceTrigger, len(req.Triggers))
		for i, t := range req.Triggers {
			triggers[i] = model.InterfaceTrigger{
				InterfaceConfigID: cfg.ID,
				TriggerType:     t.TriggerType,
				CronExpr:        t.CronExpr,
				EventSource:     t.EventSource,
				PayloadTemplate: t.PayloadTemplate,
				Condition:       t.Condition,
				FallbackMinutes: t.FallbackMinutes,
				Status:          "ENABLE",
			}
		}
		s.configRepo.UpsertTriggers(ctx, cfg.ID, triggers)
	}

	return cfg, nil
}

// UpdateConfig 更新接口配置
func (s *IntegrationService) UpdateConfig(ctx context.Context, id int64, req *model.InterfaceConfigUpdate) error {
	updates := map[string]interface{}{}
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Direction != "" {
		updates["direction"] = req.Direction
	}
	if req.Method != "" {
		updates["method"] = req.Method
	}
	if req.BaseURL != "" {
		updates["base_url"] = req.BaseURL
	}
	if req.Path != "" {
		updates["path"] = req.Path
	}
	if req.AuthType != "" {
		updates["auth_type"] = req.AuthType
	}
	if req.AuthConfig != "" {
		updates["auth_config"] = req.AuthConfig
	}
	if req.RequestContentType != "" {
		updates["request_content_type"] = req.RequestContentType
	}
	if req.RequestBodyTemplate != "" {
		updates["request_body_template"] = req.RequestBodyTemplate
	}
	if req.ResponseFormat != "" {
		updates["response_format"] = req.ResponseFormat
	}
	if req.Timeout > 0 {
		updates["timeout"] = req.Timeout
	}
	if req.RetryCount >= 0 {
		updates["retry_count"] = req.RetryCount
	}
	if req.RetryInterval >= 0 {
		updates["retry_interval"] = req.RetryInterval
	}
	if req.SourceType != "" {
		updates["source_type"] = req.SourceType
	}
	if req.SourceTable != "" {
		updates["source_table"] = req.SourceTable
	}
	if req.SourceAPI != "" {
		updates["source_api"] = req.SourceAPI
	}
	if req.SourceFilter != "" {
		updates["source_filter"] = req.SourceFilter
	}
	if req.SourceFields != "" {
		updates["source_fields"] = req.SourceFields
	}
	if req.PrimaryKey != "" {
		updates["primary_key"] = req.PrimaryKey
	}
	if req.IncrementalField != "" {
		updates["incremental_field"] = req.IncrementalField
	}
	if req.IncrementalWindow >= 0 {
		updates["incremental_window"] = req.IncrementalWindow
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	return s.configRepo.Update(ctx, id, updates)
}

// DeleteConfig 删除接口配置（级联删除字段映射和触发器）
func (s *IntegrationService) DeleteConfig(ctx context.Context, id int64) error {
	return s.configRepo.Delete(ctx, id)
}

// FieldMap CRUD
func (s *IntegrationService) CreateFieldMap(ctx context.Context, configID int64, req *model.InterfaceFieldMapCreate) error {
	m := model.InterfaceFieldMap{
		InterfaceConfigID: configID,
		FieldName:     req.FieldName,
		FieldType:    req.FieldType,
		MapType:      req.MapType,
		MapValue:     req.MapValue,
		Required:     req.Required,
		DefaultValue: req.DefaultValue,
		TransformFunc: req.TransformFunc,
		SortOrder:   req.SortOrder,
	}
	return s.configRepo.CreateFieldMaps(ctx, configID, []model.InterfaceFieldMap{m})
}

func (s *IntegrationService) UpdateFieldMap(ctx context.Context, id int64, req *model.InterfaceFieldMapCreate) error {
	updates := map[string]interface{}{
		"field_name":     req.FieldName,
		"field_type":    req.FieldType,
		"map_type":      req.MapType,
		"map_value":     req.MapValue,
		"required":      req.Required,
		"default_value": req.DefaultValue,
		"transform_func": req.TransformFunc,
		"sort_order":    req.SortOrder,
	}
	return s.configRepo.UpdateFieldMap(ctx, id, updates)
}

func (s *IntegrationService) DeleteFieldMap(ctx context.Context, id int64) error {
	return s.configRepo.DeleteFieldMap(ctx, id)
}

// Trigger CRUD
func (s *IntegrationService) CreateTrigger(ctx context.Context, configID int64, req *model.InterfaceTriggerCreate) error {
	t := model.InterfaceTrigger{
		InterfaceConfigID: configID,
		TriggerType:     req.TriggerType,
		CronExpr:        req.CronExpr,
		EventSource:     req.EventSource,
		PayloadTemplate: req.PayloadTemplate,
		Condition:       req.Condition,
		FallbackMinutes: req.FallbackMinutes,
		Status:          "ENABLE",
	}
	return s.configRepo.CreateTriggers(ctx, configID, []model.InterfaceTrigger{t})
}

func (s *IntegrationService) UpdateTrigger(ctx context.Context, id int64, req *model.InterfaceTriggerCreate) error {
	updates := map[string]interface{}{
		"trigger_type":      req.TriggerType,
		"cron_expr":        req.CronExpr,
		"event_source":     req.EventSource,
		"payload_template": req.PayloadTemplate,
		"condition":        req.Condition,
		"fallback_minutes": req.FallbackMinutes,
	}
	return s.configRepo.UpdateTrigger(ctx, id, updates)
}

func (s *IntegrationService) DeleteTrigger(ctx context.Context, id int64) error {
	return s.configRepo.DeleteTrigger(ctx, id)
}

// Execution log
func (s *IntegrationService) ListExecutionLogs(ctx context.Context, q *repository.ExecutionLogQuery) ([]model.InterfaceExecutionLog, int64, error) {
	return s.logRepo.List(ctx, q)
}

func (s *IntegrationService) GetExecutionLog(ctx context.Context, id int64) (*model.InterfaceExecutionLog, error) {
	return s.logRepo.GetByID(ctx, id)
}

// FindByEventSource finds configs matching the given event
func (s *IntegrationService) FindByEventSource(ctx context.Context, eventSource string) ([]model.InterfaceConfig, error) {
	return s.configRepo.FindByEventSource(ctx, eventSource)
}

// GetScheduleTriggers gets all schedule triggers for the scheduler
func (s *IntegrationService) GetScheduleTriggers(ctx context.Context) ([]model.InterfaceTrigger, error) {
	return s.configRepo.FindScheduleTriggers(ctx)
}

// GetConfigByID gets config by ID only
func (s *IntegrationService) GetConfigByID(ctx context.Context, id int64) (*model.InterfaceConfig, error) {
	return s.configRepo.GetByID(ctx, id)
}

// GetFieldMaps gets all field maps for a config
func (s *IntegrationService) GetFieldMaps(ctx context.Context, configID int64) ([]model.InterfaceFieldMap, error) {
	return s.configRepo.ListFieldMaps(ctx, configID)
}

// GetLastExecutionTime gets the last successful execution time for a config
func (s *IntegrationService) GetLastExecutionTime(ctx context.Context, configID int64) (*time.Time, error) {
	return s.logRepo.GetLastExecutionTime(ctx, configID)
}
