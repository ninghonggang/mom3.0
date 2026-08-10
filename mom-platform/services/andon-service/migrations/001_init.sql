-- ===========================================================================
-- Andon Service - Initial Database Migration
-- Description: Andon System core tables
-- Version: 1.0.0
-- ===========================================================================

-- ---------------------------------------------------------------------------
-- 1. andon_calls (Andon呼叫)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS andon_calls (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    call_no         VARCHAR(64)     NOT NULL,
    workstation_id  BIGINT          NOT NULL,
    line_id         BIGINT,
    equipment_id    BIGINT,
    call_type       VARCHAR(64)     NOT NULL DEFAULT 'QUALITY',
    call_status     VARCHAR(32)     NOT NULL DEFAULT 'OPEN',
    severity        VARCHAR(32)     NOT NULL DEFAULT 'MEDIUM',
    description     TEXT,
    called_by       BIGINT,
    called_at       TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    assigned_to     BIGINT,
    acknowledged_at TIMESTAMPTZ,
    resolved_at     TIMESTAMPTZ,
    resolved_by     BIGINT,
    resolution      TEXT,
    close_reason    VARCHAR(256),
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

COMMENT ON TABLE andon_calls IS 'Andon呼叫';
COMMENT ON COLUMN andon_calls.call_type IS 'QUALITY/MATERIAL/MAINTENANCE/SAFETY/SUPERVISOR';
COMMENT ON COLUMN andon_calls.call_status IS 'OPEN/ACKNOWLEDGED/IN_PROGRESS/RESOLVED/CLOSED';
COMMENT ON COLUMN andon_calls.severity IS 'LOW/MEDIUM/HIGH/CRITICAL';

CREATE INDEX idx_ac_tenant_id      ON andon_calls (tenant_id);
CREATE INDEX idx_ac_workstation_id ON andon_calls (workstation_id);
CREATE INDEX idx_ac_call_status    ON andon_calls (call_status);
CREATE INDEX idx_ac_called_at      ON andon_calls (called_at);
CREATE INDEX idx_ac_assigned_to    ON andon_calls (assigned_to);
CREATE INDEX idx_ac_equipment_id   ON andon_calls (equipment_id);

-- ---------------------------------------------------------------------------
-- 2. andon_actions (Andon处理动作)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS andon_actions (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    call_id         BIGINT          NOT NULL,
    action_type     VARCHAR(64)     NOT NULL,
    action_desc     TEXT,
    performed_by    BIGINT,
    performed_at    TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    previous_status VARCHAR(32),
    new_status      VARCHAR(32),
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_aa_call FOREIGN KEY (call_id)
        REFERENCES andon_calls (id) ON DELETE CASCADE
);

COMMENT ON TABLE andon_actions IS 'Andon处理动作/历史';
COMMENT ON COLUMN andon_actions.action_type IS 'CREATED/ACKNOWLEDGED/ASSIGNED/ESCALATED/RESOLVED/CLOSED';

CREATE INDEX idx_aa_tenant_id   ON andon_actions (tenant_id);
CREATE INDEX idx_aa_call_id     ON andon_actions (call_id);
CREATE INDEX idx_aa_performed_at ON andon_actions (performed_at);

-- ---------------------------------------------------------------------------
-- 3. alert_configs (告警配置)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS alert_configs (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    config_name     VARCHAR(256)    NOT NULL,
    alert_type      VARCHAR(64)     NOT NULL,
    target_type     VARCHAR(64)     NOT NULL DEFAULT 'WORKSTATION',
    target_id       BIGINT,
    condition_field VARCHAR(128)    NOT NULL,
    condition_op    VARCHAR(32)     NOT NULL DEFAULT 'GT',
    condition_value VARCHAR(128)    NOT NULL,
    notify_method   VARCHAR(64)     NOT NULL DEFAULT 'ANDON_BOARD',
    notify_targets  JSONB,
    trigger_cooldown_sec INTEGER    NOT NULL DEFAULT 300,
    is_active       BOOLEAN         NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

COMMENT ON TABLE alert_configs IS '告警配置';
COMMENT ON COLUMN alert_configs.alert_type IS 'THRESHOLD/EVENT/DURATION/TREND';
COMMENT ON COLUMN alert_configs.target_type IS 'WORKSTATION/LINE/EQUIPMENT/ANDON_BOARD';
COMMENT ON COLUMN alert_configs.condition_op IS 'GT/LT/EQ/NE/GTE/LTE';

CREATE INDEX idx_alc_tenant_id  ON alert_configs (tenant_id);
CREATE INDEX idx_alc_is_active  ON alert_configs (is_active);
CREATE INDEX idx_alc_target_id  ON alert_configs (target_id);

-- ---------------------------------------------------------------------------
-- 4. alerts (告警记录)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS alerts (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    config_id       BIGINT,
    call_id         BIGINT,
    alert_no        VARCHAR(64)     NOT NULL,
    alert_type      VARCHAR(64)     NOT NULL,
    alert_status    VARCHAR(32)     NOT NULL DEFAULT 'FIRED',
    severity        VARCHAR(32)     NOT NULL DEFAULT 'MEDIUM',
    title           VARCHAR(512)    NOT NULL,
    message         TEXT,
    triggered_at    TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    acknowledged_at TIMESTAMPTZ,
    acknowledged_by BIGINT,
    resolved_at     TIMESTAMPTZ,
    resolved_by     BIGINT,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_alert_config FOREIGN KEY (config_id)
        REFERENCES alert_configs (id) ON DELETE SET NULL,
    CONSTRAINT fk_alert_call FOREIGN KEY (call_id)
        REFERENCES andon_calls (id) ON DELETE SET NULL
);

COMMENT ON TABLE alerts IS '告警记录';
COMMENT ON COLUMN alerts.alert_status IS 'FIRED/ACKNOWLEDGED/RESOLVED/SUPPRESSED';

CREATE INDEX idx_al_tenant_id   ON alerts (tenant_id);
CREATE INDEX idx_al_config_id   ON alerts (config_id);
CREATE INDEX idx_al_call_id     ON alerts (call_id);
CREATE INDEX idx_al_alert_status ON alerts (alert_status);
CREATE INDEX idx_al_triggered_at ON alerts (triggered_at);

-- ---------------------------------------------------------------------------
-- 5. alert_escalations (告警升级)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS alert_escalations (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    alert_id        BIGINT          NOT NULL,
    escalation_level INTEGER        NOT NULL DEFAULT 1,
    escalate_to     BIGINT,
    escalate_method VARCHAR(64)     NOT NULL DEFAULT 'NOTIFICATION',
    escalate_reason VARCHAR(256),
    escalate_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    acknowledged    BOOLEAN         NOT NULL DEFAULT FALSE,
    acknowledged_at TIMESTAMPTZ,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_ae_alert FOREIGN KEY (alert_id)
        REFERENCES alerts (id) ON DELETE CASCADE
);

COMMENT ON TABLE alert_escalations IS '告警升级记录';
COMMENT ON COLUMN alert_escalations.escalate_method IS 'NOTIFICATION/EMAIL/SMS/CALL';

CREATE INDEX idx_ae_tenant_id   ON alert_escalations (tenant_id);
CREATE INDEX idx_ae_alert_id    ON alert_escalations (alert_id);
CREATE INDEX idx_ae_escalate_at ON alert_escalations (escalate_at);
CREATE INDEX idx_ae_escalate_to ON alert_escalations (escalate_to);
