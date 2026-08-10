package repository

import (
	"context"

	"mom-platform/services/andon-service/internal/model"
)

type Repository interface {
	// AndonCall
	CreateAndonCall(ctx context.Context, call *model.AndonCall) error
	GetAndonCall(ctx context.Context, id uint) (*model.AndonCall, error)
	GetAndonCallByNo(ctx context.Context, andonNo string) (*model.AndonCall, error)
	UpdateAndonCall(ctx context.Context, call *model.AndonCall) error
	ListActiveAndonCalls(ctx context.Context, tenantID string) ([]*model.AndonCall, error)
	ListAndonCalls(ctx context.Context, offset, limit int, andonType, status, workstationID string) ([]*model.AndonCall, int64, error)

	// AndonAction
	CreateAndonAction(ctx context.Context, action *model.AndonAction) error
	GetAndonActions(ctx context.Context, andonID uint) ([]*model.AndonAction, error)

	// AlertConfig
	CreateAlertConfig(ctx context.Context, cfg *model.AlertConfig) error
	GetAlertConfig(ctx context.Context, id uint) (*model.AlertConfig, error)
	GetAlertConfigByCode(ctx context.Context, code string) (*model.AlertConfig, error)

	// Alert
	CreateAlert(ctx context.Context, alert *model.Alert) error
	GetAlert(ctx context.Context, id uint) (*model.Alert, error)
	UpdateAlert(ctx context.Context, alert *model.Alert) error
	ListActiveAlerts(ctx context.Context) ([]*model.Alert, error)
	ListAlerts(ctx context.Context, offset, limit int, status, severity string) ([]*model.Alert, int64, error)

	// AlertEscalation
	CreateAlertEscalation(ctx context.Context, escalation *model.AlertEscalation) error
	GetAlertEscalations(ctx context.Context, alertID uint) ([]*model.AlertEscalation, error)
}
