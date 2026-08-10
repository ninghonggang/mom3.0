-- ===========================================================================
-- MES Service - Initial Database Migration
-- Description: Manufacturing Execution System core tables
-- Version: 1.0.0
-- ===========================================================================

-- ---------------------------------------------------------------------------
-- 1. production_orders (工单/生产工单)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS production_orders (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    order_no        VARCHAR(64)     NOT NULL,
    bom_id          BIGINT,
    material_id     BIGINT          NOT NULL,
    planned_qty     DECIMAL(18,4)   NOT NULL DEFAULT 0,
    completed_qty   DECIMAL(18,4)   NOT NULL DEFAULT 0,
    scrapped_qty    DECIMAL(18,4)   NOT NULL DEFAULT 0,
    planned_start   TIMESTAMPTZ,
    planned_end     TIMESTAMPTZ,
    actual_start    TIMESTAMPTZ,
    actual_end      TIMESTAMPTZ,
    status          VARCHAR(32)     NOT NULL DEFAULT 'DRAFT',
    priority        SMALLINT        NOT NULL DEFAULT 0,
    workshop_id     BIGINT,
    line_id         BIGINT,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

COMMENT ON TABLE production_orders IS '生产工单';
COMMENT ON COLUMN production_orders.status IS 'DRAFT/RELEASED/IN_PROGRESS/COMPLETED/CLOSED/CANCELLED';

CREATE INDEX idx_po_tenant_id     ON production_orders (tenant_id);
CREATE INDEX idx_po_order_no      ON production_orders (order_no);
CREATE INDEX idx_po_status        ON production_orders (status);
CREATE INDEX idx_po_material_id   ON production_orders (material_id);
CREATE INDEX idx_po_planned_start ON production_orders (planned_start);
CREATE INDEX idx_po_planned_end   ON production_orders (planned_end);

-- ---------------------------------------------------------------------------
-- 2. dispatch_records (派工记录)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS dispatch_records (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    order_id        BIGINT          NOT NULL,
    workstation_id  BIGINT          NOT NULL,
    operator_id     BIGINT,
    dispatched_qty  DECIMAL(18,4)   NOT NULL DEFAULT 0,
    dispatched_at   TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_dispatch_order FOREIGN KEY (order_id)
        REFERENCES production_orders (id) ON DELETE CASCADE
);

COMMENT ON TABLE dispatch_records IS '派工记录';

CREATE INDEX idx_dr_tenant_id      ON dispatch_records (tenant_id);
CREATE INDEX idx_dr_order_id       ON dispatch_records (order_id);
CREATE INDEX idx_dr_workstation_id ON dispatch_records (workstation_id);
CREATE INDEX idx_dr_operator_id    ON dispatch_records (operator_id);

-- ---------------------------------------------------------------------------
-- 3. job_reports (报工记录)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS job_reports (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    order_id        BIGINT          NOT NULL,
    dispatch_id     BIGINT,
    workstation_id  BIGINT          NOT NULL,
    operator_id     BIGINT,
    reported_qty    DECIMAL(18,4)   NOT NULL DEFAULT 0,
    pass_qty        DECIMAL(18,4)   NOT NULL DEFAULT 0,
    ng_qty          DECIMAL(18,4)   NOT NULL DEFAULT 0,
    report_type     VARCHAR(32)     NOT NULL DEFAULT 'NORMAL',
    reported_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_report_order FOREIGN KEY (order_id)
        REFERENCES production_orders (id) ON DELETE CASCADE,
    CONSTRAINT fk_report_dispatch FOREIGN KEY (dispatch_id)
        REFERENCES dispatch_records (id) ON DELETE SET NULL
);

COMMENT ON TABLE job_reports IS '报工记录';
COMMENT ON COLUMN job_reports.report_type IS 'NORMAL/REWORK/SUPPLEMENT';

CREATE INDEX idx_jr_tenant_id      ON job_reports (tenant_id);
CREATE INDEX idx_jr_order_id       ON job_reports (order_id);
CREATE INDEX idx_jr_workstation_id ON job_reports (workstation_id);
CREATE INDEX idx_jr_reported_at    ON job_reports (reported_at);

-- ---------------------------------------------------------------------------
-- 4. completion_records (完工记录)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS completion_records (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    order_id        BIGINT          NOT NULL,
    completed_qty   DECIMAL(18,4)   NOT NULL DEFAULT 0,
    completed_at    TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_completion_order FOREIGN KEY (order_id)
        REFERENCES production_orders (id) ON DELETE CASCADE
);

COMMENT ON TABLE completion_records IS '完工记录';

CREATE INDEX idx_cr_tenant_id    ON completion_records (tenant_id);
CREATE INDEX idx_cr_order_id     ON completion_records (order_id);
CREATE INDEX idx_cr_completed_at ON completion_records (completed_at);

-- ---------------------------------------------------------------------------
-- 5. production_tasks (生产任务)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS production_tasks (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    order_id        BIGINT          NOT NULL,
    task_type       VARCHAR(64)     NOT NULL,
    task_status     VARCHAR(32)     NOT NULL DEFAULT 'PENDING',
    assigned_to     BIGINT,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_task_order FOREIGN KEY (order_id)
        REFERENCES production_orders (id) ON DELETE CASCADE
);

COMMENT ON TABLE production_tasks IS '生产任务';
COMMENT ON COLUMN production_tasks.task_status IS 'PENDING/IN_PROGRESS/COMPLETED/CANCELLED';

CREATE INDEX idx_pt_tenant_id    ON production_tasks (tenant_id);
CREATE INDEX idx_pt_order_id     ON production_tasks (order_id);
CREATE INDEX idx_pt_task_status  ON production_tasks (task_status);
CREATE INDEX idx_pt_assigned_to  ON production_tasks (assigned_to);

-- ---------------------------------------------------------------------------
-- 6. work_cards (流转卡)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS work_cards (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    order_id        BIGINT          NOT NULL,
    card_no         VARCHAR(64)     NOT NULL,
    current_process VARCHAR(128),
    next_process    VARCHAR(128),
    status          VARCHAR(32)     NOT NULL DEFAULT 'ACTIVE',
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_workcard_order FOREIGN KEY (order_id)
        REFERENCES production_orders (id) ON DELETE CASCADE
);

COMMENT ON TABLE work_cards IS '流转卡';
COMMENT ON COLUMN work_cards.status IS 'ACTIVE/COMPLETED/CANCELLED';

CREATE INDEX idx_wc_tenant_id ON work_cards (tenant_id);
CREATE INDEX idx_wc_order_id  ON work_cards (order_id);
CREATE INDEX idx_wc_card_no   ON work_cards (card_no);
CREATE INDEX idx_wc_status    ON work_cards (status);
