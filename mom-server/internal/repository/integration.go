package repository

import (
	"context"
	"time"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

// InterfaceConfigRepository
type InterfaceConfigRepository struct {
	db *gorm.DB
}

func NewInterfaceConfigRepository(db *gorm.DB) *InterfaceConfigRepository {
	return &InterfaceConfigRepository{db: db}
}

func (r *InterfaceConfigRepository) List(ctx context.Context, q *model.InterfaceConfigQuery) ([]model.InterfaceConfig, int64, error) {
	var list []model.InterfaceConfig
	var total int64

	query := r.db.WithContext(ctx).Model(&model.InterfaceConfig{})
	if q.TenantID > 0 {
		query = query.Where("tenant_id = ?", q.TenantID)
	}
	if q.Category != "" {
		query = query.Where("category = ?", q.Category)
	}
	if q.Status != "" {
		query = query.Where("status = ?", q.Status)
	}
	if q.Keyword != "" {
		kw := "%" + q.Keyword + "%"
		query = query.Where("name LIKE ? OR code LIKE ?", kw, kw)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if q.Page > 0 && q.PageSize > 0 {
		offset := (q.Page - 1) * q.PageSize
		query = query.Offset(offset).Limit(q.PageSize)
	}
	if err := query.Order("id DESC").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *InterfaceConfigRepository) GetByID(ctx context.Context, id int64) (*model.InterfaceConfig, error) {
	var cfg model.InterfaceConfig
	err := r.db.WithContext(ctx).First(&cfg, id).Error
	return &cfg, err
}

func (r *InterfaceConfigRepository) GetByCode(ctx context.Context, code string) (*model.InterfaceConfig, error) {
	var cfg model.InterfaceConfig
	err := r.db.WithContext(ctx).Where("code = ?", code).First(&cfg).Error
	return &cfg, err
}

func (r *InterfaceConfigRepository) Create(ctx context.Context, cfg *model.InterfaceConfig) error {
	return r.db.WithContext(ctx).Create(cfg).Error
}

func (r *InterfaceConfigRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.InterfaceConfig{}).Where("id = ?", id).Updates(updates).Error
}

func (r *InterfaceConfigRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("interface_config_id = ?", id).Delete(&model.InterfaceFieldMap{}).Error; err != nil {
			return err
		}
		if err := tx.Where("interface_config_id = ?", id).Delete(&model.InterfaceTrigger{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.InterfaceConfig{}, id).Error
	})
}

// FieldMap operations
func (r *InterfaceConfigRepository) ListFieldMaps(ctx context.Context, configID int64) ([]model.InterfaceFieldMap, error) {
	var maps []model.InterfaceFieldMap
	err := r.db.WithContext(ctx).Where("interface_config_id = ?", configID).Order("sort_order ASC, id ASC").Find(&maps).Error
	return maps, err
}

func (r *InterfaceConfigRepository) CreateFieldMaps(ctx context.Context, configID int64, maps []model.InterfaceFieldMap) error {
	if len(maps) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).CreateInBatches(maps, 100).Error
}

func (r *InterfaceConfigRepository) DeleteFieldMaps(ctx context.Context, configID int64) error {
	return r.db.WithContext(ctx).Where("interface_config_id = ?", configID).Delete(&model.InterfaceFieldMap{}).Error
}

func (r *InterfaceConfigRepository) UpdateFieldMap(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.InterfaceFieldMap{}).Where("id = ?", id).Updates(updates).Error
}

func (r *InterfaceConfigRepository) DeleteFieldMap(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.InterfaceFieldMap{}, id).Error
}

// Trigger operations
func (r *InterfaceConfigRepository) ListTriggers(ctx context.Context, configID int64) ([]model.InterfaceTrigger, error) {
	var triggers []model.InterfaceTrigger
	err := r.db.WithContext(ctx).Where("interface_config_id = ?", configID).Order("id ASC").Find(&triggers).Error
	return triggers, err
}

func (r *InterfaceConfigRepository) CreateTriggers(ctx context.Context, configID int64, triggers []model.InterfaceTrigger) error {
	if len(triggers) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).CreateInBatches(triggers, 100).Error
}

func (r *InterfaceConfigRepository) DeleteTriggers(ctx context.Context, configID int64) error {
	return r.db.WithContext(ctx).Where("interface_config_id = ?", configID).Delete(&model.InterfaceTrigger{}).Error
}

func (r *InterfaceConfigRepository) UpdateTrigger(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.InterfaceTrigger{}).Where("id = ?", id).Updates(updates).Error
}

func (r *InterfaceConfigRepository) DeleteTrigger(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.InterfaceTrigger{}, id).Error
}

// FindByEventSource finds all enabled configs that have a trigger for the given event source
func (r *InterfaceConfigRepository) FindByEventSource(ctx context.Context, eventSource string) ([]model.InterfaceConfig, error) {
	var configs []model.InterfaceConfig
	err := r.db.WithContext(ctx).
		Joins("JOIN sys_interface_trigger ON sys_interface_trigger.interface_config_id = sys_interface_config.id").
		Where("sys_interface_config.status = ?", "ENABLE").
		Where("sys_interface_trigger.trigger_type = ?", "EVENT").
		Where("sys_interface_trigger.event_source = ?", eventSource).
		Where("sys_interface_trigger.status = ?", "ENABLE").
		Find(&configs).Error
	return configs, err
}

// FindScheduleTriggers finds all enabled schedule triggers
func (r *InterfaceConfigRepository) FindScheduleTriggers(ctx context.Context) ([]model.InterfaceTrigger, error) {
	var triggers []model.InterfaceTrigger
	err := r.db.WithContext(ctx).
		Where("status = ?", "ENABLE").
		Where("trigger_type = ?", "SCHEDULE").
		Where("cron_expr != ''").
		Find(&triggers).Error
	return triggers, err
}

// QuerySourceData queries data from a source table with optional incremental filtering
func (r *InterfaceConfigRepository) QuerySourceData(ctx context.Context, tenantID int64, cfg *model.InterfaceConfig, lastExecTime *time.Time) ([]map[string]interface{}, error) {
	if cfg.SourceTable == "" {
		return nil, nil
	}

	sql := "SELECT "
	if cfg.SourceFields != "" {
		sql += cfg.SourceFields
	} else {
		sql += "*"
	}
	sql += " FROM " + cfg.SourceTable + " WHERE tenant_id = ?"
	args := []interface{}{tenantID}

	// Incremental filter
	if cfg.IncrementalField != "" && lastExecTime != nil && cfg.IncrementalWindow > 0 {
		since := lastExecTime.Add(-time.Duration(cfg.IncrementalWindow) * time.Minute)
		sql += " AND " + cfg.IncrementalField + " >= ?"
		args = append(args, since)
	}

	sql += " ORDER BY id ASC LIMIT 1000"

	rows, err := r.db.WithContext(ctx).Raw(sql, args...).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}
		row := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]
			if b, ok := val.([]byte); ok {
				row[col] = string(b)
			} else {
				row[col] = val
			}
		}
		results = append(results, row)
	}
	return results, nil
}

// InterfaceExecutionLogRepository
type InterfaceExecutionLogRepository struct {
	db *gorm.DB
}

func NewInterfaceExecutionLogRepository(db *gorm.DB) *InterfaceExecutionLogRepository {
	return &InterfaceExecutionLogRepository{db: db}
}

type ExecutionLogQuery struct {
	TenantID         int64
	InterfaceConfigID int64
	Status           string
	TriggerType      string
	StartDate        *time.Time
	EndDate          *time.Time
	Page             int
	PageSize         int
}

func (r *InterfaceExecutionLogRepository) List(ctx context.Context, q *ExecutionLogQuery) ([]model.InterfaceExecutionLog, int64, error) {
	var list []model.InterfaceExecutionLog
	var total int64

	query := r.db.WithContext(ctx).Model(&model.InterfaceExecutionLog{})
	if q.TenantID > 0 {
		query = query.Where("tenant_id = ?", q.TenantID)
	}
	if q.InterfaceConfigID > 0 {
		query = query.Where("interface_config_id = ?", q.InterfaceConfigID)
	}
	if q.Status != "" {
		query = query.Where("status = ?", q.Status)
	}
	if q.TriggerType != "" {
		query = query.Where("trigger_type = ?", q.TriggerType)
	}
	if q.StartDate != nil {
		query = query.Where("start_time >= ?", q.StartDate)
	}
	if q.EndDate != nil {
		query = query.Where("start_time <= ?", q.EndDate)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if q.Page > 0 && q.PageSize > 0 {
		offset := (q.Page - 1) * q.PageSize
		query = query.Offset(offset).Limit(q.PageSize)
	}
	if err := query.Order("start_time DESC").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *InterfaceExecutionLogRepository) GetByID(ctx context.Context, id int64) (*model.InterfaceExecutionLog, error) {
	var log model.InterfaceExecutionLog
	err := r.db.WithContext(ctx).First(&log, id).Error
	return &log, err
}

func (r *InterfaceExecutionLogRepository) Create(ctx context.Context, log *model.InterfaceExecutionLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *InterfaceExecutionLogRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.InterfaceExecutionLog{}).Where("id = ?", id).Updates(updates).Error
}

func (r *InterfaceExecutionLogRepository) GetLastExecutionTime(ctx context.Context, configID int64) (*time.Time, error) {
	var log model.InterfaceExecutionLog
	err := r.db.WithContext(ctx).
		Where("interface_config_id = ?", configID).
		Where("status = ?", "SUCCESS").
		Order("end_time DESC").
		First(&log).Error
	if err != nil {
		return nil, err
	}
	return &log.EndTime, nil
}

// DeleteOldLogs deletes logs older than days
func (r *InterfaceExecutionLogRepository) DeleteOldLogs(ctx context.Context, days int) error {
	if days <= 0 {
		return nil
	}
	cutoff := time.Now().AddDate(0, 0, -days)
	return r.db.WithContext(ctx).Where("start_time < ?", cutoff).Delete(&model.InterfaceExecutionLog{}).Error
}

// UpsertFieldMaps replaces all field maps for a config
func (r *InterfaceConfigRepository) UpsertFieldMaps(ctx context.Context, configID int64, maps []model.InterfaceFieldMap) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("interface_config_id = ?", configID).Delete(&model.InterfaceFieldMap{}).Error; err != nil {
			return err
		}
		for i := range maps {
			maps[i].InterfaceConfigID = configID
		}
		if len(maps) > 0 {
			return tx.CreateInBatches(maps, 100).Error
		}
		return nil
	})
}

// UpsertTriggers replaces all triggers for a config
func (r *InterfaceConfigRepository) UpsertTriggers(ctx context.Context, configID int64, triggers []model.InterfaceTrigger) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("interface_config_id = ?", configID).Delete(&model.InterfaceTrigger{}).Error; err != nil {
			return err
		}
		for i := range triggers {
			triggers[i].InterfaceConfigID = configID
		}
		if len(triggers) > 0 {
			return tx.CreateInBatches(triggers, 100).Error
		}
		return nil
	})
}

// GetWithRelations loads config with field maps and triggers
func (r *InterfaceConfigRepository) GetWithRelations(ctx context.Context, id int64) (*model.InterfaceConfig, []model.InterfaceFieldMap, []model.InterfaceTrigger, error) {
	var cfg model.InterfaceConfig
	if err := r.db.WithContext(ctx).First(&cfg, id).Error; err != nil {
		return nil, nil, nil, err
	}
	maps, _ := r.ListFieldMaps(ctx, id)
	triggers, _ := r.ListTriggers(ctx, id)
	return &cfg, maps, triggers, nil
}
