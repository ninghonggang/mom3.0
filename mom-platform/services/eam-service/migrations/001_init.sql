-- ===========================================================================
-- EAM Service - Initial Database Migration
-- Description: Enterprise Asset Management core tables
-- Version: 1.0.0
-- ===========================================================================

-- ---------------------------------------------------------------------------
-- 1. equipment (设备台账)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS equipment (
    id                  BIGSERIAL       PRIMARY KEY,
    tenant_id           BIGINT          NOT NULL,
    equipment_no        VARCHAR(64)     NOT NULL,
    equipment_name      VARCHAR(256)    NOT NULL,
    equipment_type      VARCHAR(128),
    model               VARCHAR(128),
    manufacturer        VARCHAR(256),
    serial_no           VARCHAR(128),
    workshop_id         BIGINT,
    line_id             BIGINT,
    workstation_id      BIGINT,
    status              VARCHAR(32)     NOT NULL DEFAULT 'IDLE',
    purchase_date       DATE,
    warranty_end        DATE,
    created_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);

COMMENT ON TABLE equipment IS '设备台账';
COMMENT ON COLUMN equipment.status IS 'IDLE/RUNNING/MAINTENANCE/BREAKDOWN/SCRAPPED';

CREATE INDEX idx_eq_tenant_id      ON equipment (tenant_id);
CREATE INDEX idx_eq_equipment_no   ON equipment (equipment_no);
CREATE INDEX idx_eq_status         ON equipment (status);
CREATE INDEX idx_eq_workshop_id    ON equipment (workshop_id);
CREATE INDEX idx_eq_workstation_id ON equipment (workstation_id);

-- ---------------------------------------------------------------------------
-- 2. maintenance_plans (维保计划)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS maintenance_plans (
    id                  BIGSERIAL       PRIMARY KEY,
    tenant_id           BIGINT          NOT NULL,
    plan_no             VARCHAR(64)     NOT NULL,
    plan_name           VARCHAR(256)    NOT NULL,
    equipment_id        BIGINT          NOT NULL,
    plan_type           VARCHAR(64)     NOT NULL DEFAULT 'PREVENTIVE',
    frequency_type      VARCHAR(64)     NOT NULL DEFAULT 'DATE',
    frequency_value     INTEGER         NOT NULL DEFAULT 1,
    frequency_unit      VARCHAR(32)     NOT NULL DEFAULT 'MONTH',
    estimated_duration  INTEGER,
    status              VARCHAR(32)     NOT NULL DEFAULT 'ACTIVE',
    next_scheduled      TIMESTAMPTZ,
    created_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ,

    CONSTRAINT fk_mp_equipment FOREIGN KEY (equipment_id)
        REFERENCES equipment (id) ON DELETE CASCADE
);

COMMENT ON TABLE maintenance_plans IS '维保计划';
COMMENT ON COLUMN maintenance_plans.plan_type IS 'PREVENTIVE/PREDICTIVE/CORRECTIVE';
COMMENT ON COLUMN maintenance_plans.status IS 'ACTIVE/INACTIVE/COMPLETED';

CREATE INDEX idx_mp_tenant_id     ON maintenance_plans (tenant_id);
CREATE INDEX idx_mp_equipment_id  ON maintenance_plans (equipment_id);
CREATE INDEX idx_mp_status        ON maintenance_plans (status);
CREATE INDEX idx_mp_next_sched    ON maintenance_plans (next_scheduled);

-- ---------------------------------------------------------------------------
-- 3. repair_orders (维修工单)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS repair_orders (
    id                  BIGSERIAL       PRIMARY KEY,
    tenant_id           BIGINT          NOT NULL,
    repair_no           VARCHAR(64)     NOT NULL,
    equipment_id        BIGINT          NOT NULL,
    plan_id             BIGINT,
    fault_description   TEXT,
    repair_type         VARCHAR(64)     NOT NULL DEFAULT 'BREAKDOWN',
    severity            VARCHAR(32)     NOT NULL DEFAULT 'MEDIUM',
    status              VARCHAR(32)     NOT NULL DEFAULT 'REPORTED',
    reported_by         BIGINT,
    reported_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    assigned_to         BIGINT,
    started_at          TIMESTAMPTZ,
    completed_at        TIMESTAMPTZ,
    downtime_minutes    INTEGER         NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ,

    CONSTRAINT fk_ro_equipment FOREIGN KEY (equipment_id)
        REFERENCES equipment (id) ON DELETE CASCADE,
    CONSTRAINT fk_ro_plan FOREIGN KEY (plan_id)
        REFERENCES maintenance_plans (id) ON DELETE SET NULL
);

COMMENT ON TABLE repair_orders IS '维修工单';
COMMENT ON COLUMN repair_orders.status IS 'REPORTED/ASSIGNED/IN_PROGRESS/COMPLETED/CLOSED';
COMMENT ON COLUMN repair_orders.severity IS 'LOW/MEDIUM/HIGH/CRITICAL';

CREATE INDEX idx_ro_tenant_id    ON repair_orders (tenant_id);
CREATE INDEX idx_ro_equipment_id ON repair_orders (equipment_id);
CREATE INDEX idx_ro_status       ON repair_orders (status);
CREATE INDEX idx_ro_reported_at  ON repair_orders (reported_at);
CREATE INDEX idx_ro_assigned_to  ON repair_orders (assigned_to);

-- ---------------------------------------------------------------------------
-- 4. equipment_oee (OEE记录)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS equipment_oee (
    id                  BIGSERIAL       PRIMARY KEY,
    tenant_id           BIGINT          NOT NULL,
    equipment_id        BIGINT          NOT NULL,
    record_date         DATE            NOT NULL,
    shift               VARCHAR(32),
    planned_runtime     DECIMAL(18,2)   NOT NULL DEFAULT 0,
    actual_runtime      DECIMAL(18,2)   NOT NULL DEFAULT 0,
    downtime            DECIMAL(18,2)   NOT NULL DEFAULT 0,
    ideal_cycle_time    DECIMAL(18,4),
    actual_output       INTEGER         NOT NULL DEFAULT 0,
    good_output         INTEGER         NOT NULL DEFAULT 0,
    availability        DECIMAL(5,2),
    performance         DECIMAL(5,2),
    quality             DECIMAL(5,2),
    oee                 DECIMAL(5,2),
    created_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ,

    CONSTRAINT fk_oee_equipment FOREIGN KEY (equipment_id)
        REFERENCES equipment (id) ON DELETE CASCADE
);

COMMENT ON TABLE equipment_oee IS '设备综合效率 OEE 记录';

CREATE INDEX idx_oee_tenant_id    ON equipment_oee (tenant_id);
CREATE INDEX idx_oee_equipment_id ON equipment_oee (equipment_id);
CREATE INDEX idx_oee_record_date  ON equipment_oee (record_date);

-- ---------------------------------------------------------------------------
-- 5. equipment_checks (设备点检)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS equipment_checks (
    id                  BIGSERIAL       PRIMARY KEY,
    tenant_id           BIGINT          NOT NULL,
    equipment_id        BIGINT          NOT NULL,
    check_no            VARCHAR(64)     NOT NULL,
    check_type          VARCHAR(64)     NOT NULL DEFAULT 'DAILY',
    check_status        VARCHAR(32)     NOT NULL DEFAULT 'PENDING',
    check_result        VARCHAR(32),
    checked_by          BIGINT,
    checked_at          TIMESTAMPTZ,
    remark              TEXT,
    created_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ,

    CONSTRAINT fk_ec_equipment FOREIGN KEY (equipment_id)
        REFERENCES equipment (id) ON DELETE CASCADE
);

COMMENT ON TABLE equipment_checks IS '设备点检记录';
COMMENT ON COLUMN equipment_checks.check_type IS 'DAILY/WEEKLY/MONTHLY/SPECIAL';
COMMENT ON COLUMN equipment_checks.check_result IS 'PASS/FAIL/NOT_APPLICABLE';

CREATE INDEX idx_ec_tenant_id    ON equipment_checks (tenant_id);
CREATE INDEX idx_ec_equipment_id ON equipment_checks (equipment_id);
CREATE INDEX idx_ec_check_status ON equipment_checks (check_status);
CREATE INDEX idx_ec_checked_at   ON equipment_checks (checked_at);

-- ---------------------------------------------------------------------------
-- 6. equipment_downtimes (设备停机记录)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS equipment_downtimes (
    id                  BIGSERIAL       PRIMARY KEY,
    tenant_id           BIGINT          NOT NULL,
    equipment_id        BIGINT          NOT NULL,
    repair_order_id     BIGINT,
    downtime_type       VARCHAR(64)     NOT NULL DEFAULT 'UNPLANNED',
    reason              VARCHAR(512),
    started_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    ended_at            TIMESTAMPTZ,
    duration_minutes    INTEGER         NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ,

    CONSTRAINT fk_ed_equipment FOREIGN KEY (equipment_id)
        REFERENCES equipment (id) ON DELETE CASCADE,
    CONSTRAINT fk_ed_repair_order FOREIGN KEY (repair_order_id)
        REFERENCES repair_orders (id) ON DELETE SET NULL
);

COMMENT ON TABLE equipment_downtimes IS '设备停机记录';
COMMENT ON COLUMN equipment_downtimes.downtime_type IS 'PLANNED/UNPLANNED/SETUP/CHANGEOVER';

CREATE INDEX idx_ed_tenant_id     ON equipment_downtimes (tenant_id);
CREATE INDEX idx_ed_equipment_id  ON equipment_downtimes (equipment_id);
CREATE INDEX idx_ed_started_at    ON equipment_downtimes (started_at);
CREATE INDEX idx_ed_ended_at      ON equipment_downtimes (ended_at);
