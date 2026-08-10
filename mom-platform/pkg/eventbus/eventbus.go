package eventbus

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
)

// EventPublisher — NATS JetStream 事件发布器
type EventPublisher struct {
	nc    *nats.Conn
	js    jetstream.JetStream
	log   *zap.Logger
}

// NewEventPublisher — 创建事件发布器
func NewEventPublisher(natsURL string, log *zap.Logger) (*EventPublisher, error) {
	nc, err := nats.Connect(natsURL,
		nats.MaxReconnects(-1),
		nats.ReconnectWait(2*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("connect nats: %w", err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		return nil, fmt.Errorf("create jetstream: %w", err)
	}

	log.Info("NATS JetStream connected", zap.String("url", natsURL))

	return &EventPublisher{nc: nc, js: js, log: log}, nil
}

// Publish — 发布领域事件
// subject 格式: "mom.{domain}.{entity}.{action}"，例如 "mom.production.order.created"
func (p *EventPublisher) Publish(ctx context.Context, subject string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal event payload: %w", err)
	}

	_, err = p.js.Publish(ctx, subject, data)
	if err != nil {
		p.log.Error("failed to publish event",
			zap.String("subject", subject),
			zap.Error(err),
		)
		return fmt.Errorf("publish event: %w", err)
	}

	p.log.Debug("event published",
		zap.String("subject", subject),
		zap.Int("bytes", len(data)),
	)
	return nil
}

// Close — 关闭连接
func (p *EventPublisher) Close() {
	if p.nc != nil {
		p.nc.Drain()
	}
}

// EnsureStream — 确保 JetStream stream 存在
func (p *EventPublisher) EnsureStream(ctx context.Context, streamName string, subjects []string) error {
	_, err := p.js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:      streamName,
		Subjects:  subjects,
		Retention: jetstream.LimitsPolicy,
		MaxAge:    72 * time.Hour,
	})
	if err != nil {
		return fmt.Errorf("create stream %s: %w", streamName, err)
	}
	p.log.Info("stream ensured", zap.String("stream", streamName), zap.Strings("subjects", subjects))
	return nil
}

// NATS subjects — 各业务域事件路由
const (
	SubjectMDMMaterialCreated    = "mom.mdm.material.created"
	SubjectMDMMaterialObsolete   = "mom.mdm.material.obsolete"
	SubjectMDMBomActivated       = "mom.mdm.bom.activated"
	SubjectMDMSupplierActivated  = "mom.mdm.supplier.activated"
	SubjectMDMSupplierBlacklist  = "mom.mdm.supplier.blacklisted"

	SubjectMESOrderCreated       = "mom.mes.order.created"
	SubjectMESOrderCompleted     = "mom.mes.order.completed"
	SubjectMESOrderHold          = "mom.mes.order.hold"
	SubjectMESReportAudited      = "mom.mes.report.audited"
	SubjectMESException          = "mom.mes.exception.created"

	SubjectAPSMPSReleased        = "mom.aps.mps.released"
	SubjectAPSMRPCompleted       = "mom.aps.mrp.completed"
	SubjectAPSSchedulePublished  = "mom.aps.schedule.published"
	SubjectAPSMaterialShortage   = "mom.aps.material.shortage"

	SubjectWMSReceiveCompleted   = "mom.wms.receive.completed"
	SubjectWMSShipped            = "mom.wms.delivery.shipped"

	SubjectAndonTriggered        = "mom.andon.call.triggered"
	SubjectAndonEscalated        = "mom.andon.call.escalated"
	SubjectAlertTriggered        = "mom.andon.alert.triggered"

	SubjectEAMDowntimeStart      = "mom.eam.downtime.start"
	SubjectEAMDowntimeResolve    = "mom.eam.downtime.resolve"
)
