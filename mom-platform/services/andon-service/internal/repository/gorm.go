package repository

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"mom-platform/services/andon-service/internal/model"
)

type gormRepo struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewGormRepository(db *gorm.DB, logger *zap.Logger) Repository {
	return &gormRepo{db: db, logger: logger}
}

// AndonCall
func (r *gormRepo) CreateAndonCall(ctx context.Context, call *model.AndonCall) error {
	return r.db.WithContext(ctx).Create(call).Error
}

func (r *gormRepo) GetAndonCall(ctx context.Context, id uint) (*model.AndonCall, error) {
	var call model.AndonCall
	err := r.db.WithContext(ctx).Preload("Actions").First(&call, id).Error
	if err != nil {
		return nil, err
	}
	return &call, nil
}

func (r *gormRepo) GetAndonCallByNo(ctx context.Context, andonNo string) (*model.AndonCall, error) {
	var call model.AndonCall
	err := r.db.WithContext(ctx).Where("andon_no = ?", andonNo).First(&call).Error
	if err != nil {
		return nil, err
	}
	return &call, nil
}

func (r *gormRepo) UpdateAndonCall(ctx context.Context, call *model.AndonCall) error {
	return r.db.WithContext(ctx).Save(call).Error
}

func (r *gormRepo) ListActiveAndonCalls(ctx context.Context, tenantID string) ([]*model.AndonCall, error) {
	var calls []*model.AndonCall
	q := r.db.WithContext(ctx).Where("status NOT IN ?", []string{"RESOLVED", "CLOSED", "CANCELLED"})
	if tenantID != "" && tenantID != "0" {
		q = q.Where("tenant_id = ?", tenantID)
	}
	err := q.Order("triggered_at DESC").Find(&calls).Error
	return calls, err
}

func (r *gormRepo) ListAndonCalls(ctx context.Context, offset, limit int, andonType, status, workstationID string) ([]*model.AndonCall, int64, error) {
	q := r.db.WithContext(ctx).Model(&model.AndonCall{})

	if andonType != "" {
		q = q.Where("andon_type = ?", andonType)
	}
	if status != "" {
		q = q.Where("status = ?", status)
	}
	if workstationID != "" {
		q = q.Where("workstation_id = ?", workstationID)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var calls []*model.AndonCall
	err := q.Offset(offset).Limit(limit).Order("triggered_at DESC").Find(&calls).Error
	return calls, total, err
}

// AndonAction
func (r *gormRepo) CreateAndonAction(ctx context.Context, action *model.AndonAction) error {
	return r.db.WithContext(ctx).Create(action).Error
}

func (r *gormRepo) GetAndonActions(ctx context.Context, andonID uint) ([]*model.AndonAction, error) {
	var actions []*model.AndonAction
	err := r.db.WithContext(ctx).Where("andon_id = ?", andonID).Order("action_time ASC").Find(&actions).Error
	return actions, err
}

// AlertConfig
func (r *gormRepo) CreateAlertConfig(ctx context.Context, cfg *model.AlertConfig) error {
	return r.db.WithContext(ctx).Create(cfg).Error
}

func (r *gormRepo) GetAlertConfig(ctx context.Context, id uint) (*model.AlertConfig, error) {
	var cfg model.AlertConfig
	err := r.db.WithContext(ctx).First(&cfg, id).Error
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (r *gormRepo) GetAlertConfigByCode(ctx context.Context, code string) (*model.AlertConfig, error) {
	var cfg model.AlertConfig
	err := r.db.WithContext(ctx).Where("config_code = ?", code).First(&cfg).Error
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Alert
func (r *gormRepo) CreateAlert(ctx context.Context, alert *model.Alert) error {
	return r.db.WithContext(ctx).Create(alert).Error
}

func (r *gormRepo) GetAlert(ctx context.Context, id uint) (*model.Alert, error) {
	var alert model.Alert
	err := r.db.WithContext(ctx).Preload("Config").First(&alert, id).Error
	if err != nil {
		return nil, err
	}
	return &alert, nil
}

func (r *gormRepo) UpdateAlert(ctx context.Context, alert *model.Alert) error {
	return r.db.WithContext(ctx).Save(alert).Error
}

func (r *gormRepo) ListActiveAlerts(ctx context.Context) ([]*model.Alert, error) {
	var alerts []*model.Alert
	err := r.db.WithContext(ctx).Where("status IN ?", []string{"ACTIVE"}).
		Order("triggered_at ASC").Find(&alerts).Error
	return alerts, err
}

func (r *gormRepo) ListAlerts(ctx context.Context, offset, limit int, status, severity string) ([]*model.Alert, int64, error) {
	q := r.db.WithContext(ctx).Model(&model.Alert{})

	if status != "" {
		q = q.Where("status = ?", status)
	}
	if severity != "" {
		q = q.Where("severity = ?", severity)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var alerts []*model.Alert
	err := q.Offset(offset).Limit(limit).Order("triggered_at DESC").Find(&alerts).Error
	return alerts, total, err
}

// AlertEscalation
func (r *gormRepo) CreateAlertEscalation(ctx context.Context, escalation *model.AlertEscalation) error {
	return r.db.WithContext(ctx).Create(escalation).Error
}

func (r *gormRepo) GetAlertEscalations(ctx context.Context, alertID uint) ([]*model.AlertEscalation, error) {
	var escalations []*model.AlertEscalation
	err := r.db.WithContext(ctx).Where("alert_id = ?", alertID).Order("level ASC").Find(&escalations).Error
	return escalations, err
}
