-- ===========================================================================
-- QMS Service - Initial Database Migration
-- Description: Quality Management System core tables
-- Version: 1.0.0
-- ===========================================================================

-- ---------------------------------------------------------------------------
-- 1. inspection_plans (检验计划)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS inspection_plans (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    plan_no         VARCHAR(64)     NOT NULL,
    plan_name       VARCHAR(256)    NOT NULL,
    material_id     BIGINT,
    plan_type       VARCHAR(64)     NOT NULL DEFAULT 'INCOMING',
    status          VARCHAR(32)     NOT NULL DEFAULT 'ACTIVE',
    valid_from      TIMESTAMPTZ,
    valid_to        TIMESTAMPTZ,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

COMMENT ON TABLE inspection_plans IS '检验计划';
COMMENT ON COLUMN inspection_plans.plan_type IS 'INCOMING/PROCESS/FINAL';
COMMENT ON COLUMN inspection_plans.status IS 'ACTIVE/INACTIVE/ARCHIVED';

CREATE INDEX idx_ip_tenant_id   ON inspection_plans (tenant_id);
CREATE INDEX idx_ip_plan_no     ON inspection_plans (plan_no);
CREATE INDEX idx_ip_material_id ON inspection_plans (material_id);
CREATE INDEX idx_ip_status      ON inspection_plans (status);

-- ---------------------------------------------------------------------------
-- 2. inspection_characteristics (检验特性)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS inspection_characteristics (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    char_no         VARCHAR(64)     NOT NULL,
    char_name       VARCHAR(256)    NOT NULL,
    char_type       VARCHAR(64)     NOT NULL DEFAULT 'VARIABLE',
    unit            VARCHAR(32),
    usl             DECIMAL(18,4),
    lsl             DECIMAL(18,4),
    target_value    DECIMAL(18,4),
    testing_method  VARCHAR(256),
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

COMMENT ON TABLE inspection_characteristics IS '检验特性';
COMMENT ON COLUMN inspection_characteristics.char_type IS 'VARIABLE/ATTRIBUTE/VISUAL';

CREATE INDEX idx_ic_tenant_id ON inspection_characteristics (tenant_id);
CREATE INDEX idx_ic_char_no   ON inspection_characteristics (char_no);

-- ---------------------------------------------------------------------------
-- 3. inspection_plan_characteristics (检验计划-特性关联)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS inspection_plan_characteristics (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    plan_id         BIGINT          NOT NULL,
    characteristic_id BIGINT        NOT NULL,
    sample_size     INTEGER         NOT NULL DEFAULT 1,
    aql             DECIMAL(18,4),
    sort_order      INTEGER         NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_ipc_plan FOREIGN KEY (plan_id)
        REFERENCES inspection_plans (id) ON DELETE CASCADE,
    CONSTRAINT fk_ipc_char FOREIGN KEY (characteristic_id)
        REFERENCES inspection_characteristics (id) ON DELETE CASCADE,
    CONSTRAINT uk_ipc_plan_char UNIQUE (plan_id, characteristic_id)
);

COMMENT ON TABLE inspection_plan_characteristics IS '检验计划特性关联';

CREATE INDEX idx_ipc_tenant_id         ON inspection_plan_characteristics (tenant_id);
CREATE INDEX idx_ipc_plan_id           ON inspection_plan_characteristics (plan_id);
CREATE INDEX idx_ipc_characteristic_id ON inspection_plan_characteristics (characteristic_id);

-- ---------------------------------------------------------------------------
-- 4. inspection_sheets (检验单)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS inspection_sheets (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    sheet_no        VARCHAR(64)     NOT NULL,
    plan_id         BIGINT,
    order_id        BIGINT,
    material_id     BIGINT          NOT NULL,
    batch_no        VARCHAR(128),
    lot_qty         DECIMAL(18,4)   NOT NULL DEFAULT 1,
    sample_qty      INTEGER         NOT NULL DEFAULT 0,
    sheet_status    VARCHAR(32)     NOT NULL DEFAULT 'PENDING',
    result          VARCHAR(32),
    inspected_by    BIGINT,
    inspected_at    TIMESTAMPTZ,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_isheet_plan FOREIGN KEY (plan_id)
        REFERENCES inspection_plans (id) ON DELETE SET NULL
);

COMMENT ON TABLE inspection_sheets IS '检验单';
COMMENT ON COLUMN inspection_sheets.sheet_status IS 'PENDING/IN_PROGRESS/COMPLETED/CLOSED';
COMMENT ON COLUMN inspection_sheets.result IS 'PASS/FAIL/CONCESSION';

CREATE INDEX idx_is_tenant_id    ON inspection_sheets (tenant_id);
CREATE INDEX idx_is_sheet_no     ON inspection_sheets (sheet_no);
CREATE INDEX idx_is_plan_id      ON inspection_sheets (plan_id);
CREATE INDEX idx_is_material_id  ON inspection_sheets (material_id);
CREATE INDEX idx_is_sheet_status ON inspection_sheets (sheet_status);
CREATE INDEX idx_is_inspected_at ON inspection_sheets (inspected_at);

-- ---------------------------------------------------------------------------
-- 5. inspection_results (检验结果明细)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS inspection_results (
    id                  BIGSERIAL       PRIMARY KEY,
    tenant_id           BIGINT          NOT NULL,
    sheet_id            BIGINT          NOT NULL,
    characteristic_id   BIGINT          NOT NULL,
    measured_value      DECIMAL(18,4),
    is_pass             BOOLEAN,
    inspector_id        BIGINT,
    inspected_at        TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    remark              VARCHAR(512),
    created_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ,

    CONSTRAINT fk_ir_sheet FOREIGN KEY (sheet_id)
        REFERENCES inspection_sheets (id) ON DELETE CASCADE,
    CONSTRAINT fk_ir_char FOREIGN KEY (characteristic_id)
        REFERENCES inspection_characteristics (id) ON DELETE CASCADE
);

COMMENT ON TABLE inspection_results IS '检验结果明细';

CREATE INDEX idx_ir_tenant_id         ON inspection_results (tenant_id);
CREATE INDEX idx_ir_sheet_id          ON inspection_results (sheet_id);
CREATE INDEX idx_ir_characteristic_id ON inspection_results (characteristic_id);
CREATE INDEX idx_ir_inspected_at      ON inspection_results (inspected_at);

-- ---------------------------------------------------------------------------
-- 6. ncrs (不合格品报告)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS ncrs (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    ncr_no          VARCHAR(64)     NOT NULL,
    sheet_id        BIGINT,
    order_id        BIGINT,
    defect_code_id  BIGINT,
    material_id     BIGINT          NOT NULL,
    defect_qty      DECIMAL(18,4)   NOT NULL DEFAULT 0,
    severity        VARCHAR(32)     NOT NULL DEFAULT 'MINOR',
    ncr_status      VARCHAR(32)     NOT NULL DEFAULT 'OPEN',
    reported_by     BIGINT,
    reported_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    description     TEXT,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_ncr_sheet FOREIGN KEY (sheet_id)
        REFERENCES inspection_sheets (id) ON DELETE SET NULL
);

COMMENT ON TABLE ncrs IS '不合格品报告 (NCR)';
COMMENT ON COLUMN ncrs.severity IS 'MINOR/MAJOR/CRITICAL';
COMMENT ON COLUMN ncrs.ncr_status IS 'OPEN/UNDER_REVIEW/DISPOSITIONED/CLOSED';

CREATE INDEX idx_ncr_tenant_id    ON ncrs (tenant_id);
CREATE INDEX idx_ncr_ncr_no       ON ncrs (ncr_no);
CREATE INDEX idx_ncr_material_id  ON ncrs (material_id);
CREATE INDEX idx_ncr_ncr_status   ON ncrs (ncr_status);
CREATE INDEX idx_ncr_reported_at  ON ncrs (reported_at);

-- ---------------------------------------------------------------------------
-- 7. ncr_actions (不合格品处置措施)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS ncr_actions (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    ncr_id          BIGINT          NOT NULL,
    action_type     VARCHAR(64)     NOT NULL,
    action_status   VARCHAR(32)     NOT NULL DEFAULT 'PENDING',
    assigned_to     BIGINT,
    description     TEXT,
    due_date        TIMESTAMPTZ,
    completed_at    TIMESTAMPTZ,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_ncra_ncr FOREIGN KEY (ncr_id)
        REFERENCES ncrs (id) ON DELETE CASCADE
);

COMMENT ON TABLE ncr_actions IS 'NCR处置措施';
COMMENT ON COLUMN ncr_actions.action_type IS 'REWORK/SCRAP/REGRADE/CONCESSION/RETURN_TO_SUPPLIER';
COMMENT ON COLUMN ncr_actions.action_status IS 'PENDING/IN_PROGRESS/COMPLETED';

CREATE INDEX idx_ncra_tenant_id    ON ncr_actions (tenant_id);
CREATE INDEX idx_ncra_ncr_id       ON ncr_actions (ncr_id);
CREATE INDEX idx_ncra_action_status ON ncr_actions (action_status);

-- ---------------------------------------------------------------------------
-- 8. defect_codes (缺陷代码)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS defect_codes (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    defect_code     VARCHAR(64)     NOT NULL,
    defect_name     VARCHAR(256)    NOT NULL,
    defect_category VARCHAR(128),
    severity        VARCHAR(32)     NOT NULL DEFAULT 'MINOR',
    is_active       BOOLEAN         NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

COMMENT ON TABLE defect_codes IS '缺陷代码';

CREATE INDEX idx_dc_tenant_id    ON defect_codes (tenant_id);
CREATE INDEX idx_dc_defect_code  ON defect_codes (defect_code);
CREATE INDEX idx_dc_is_active    ON defect_codes (is_active);

-- ---------------------------------------------------------------------------
-- 9. spc_data (SPC数据采集)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS spc_data (
    id                  BIGSERIAL       PRIMARY KEY,
    tenant_id           BIGINT          NOT NULL,
    characteristic_id   BIGINT          NOT NULL,
    sheet_id            BIGINT,
    sample_no           INTEGER         NOT NULL DEFAULT 1,
    measured_value      DECIMAL(18,4)   NOT NULL,
    measured_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    machine_id          BIGINT,
    batch_no            VARCHAR(128),
    created_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ,

    CONSTRAINT fk_spc_char FOREIGN KEY (characteristic_id)
        REFERENCES inspection_characteristics (id) ON DELETE CASCADE
);

COMMENT ON TABLE spc_data IS 'SPC统计过程控制数据';

CREATE INDEX idx_spc_tenant_id         ON spc_data (tenant_id);
CREATE INDEX idx_spc_characteristic_id ON spc_data (characteristic_id);
CREATE INDEX idx_spc_measured_at       ON spc_data (measured_at);
CREATE INDEX idx_spc_batch_no          ON spc_data (batch_no);
