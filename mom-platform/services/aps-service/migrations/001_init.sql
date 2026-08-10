-- ===========================================================================
-- APS Service - Initial Database Migration
-- Description: Advanced Planning & Scheduling core tables
-- Version: 1.0.0
-- ===========================================================================

-- ---------------------------------------------------------------------------
-- 1. work_centers (工作中心)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS work_centers (
    id                  BIGSERIAL       PRIMARY KEY,
    tenant_id           BIGINT          NOT NULL,
    work_center_no      VARCHAR(64)     NOT NULL,
    work_center_name    VARCHAR(256)    NOT NULL,
    work_center_type    VARCHAR(64)     NOT NULL DEFAULT 'PRODUCTION',
    capacity_per_hour   DECIMAL(18,2)   NOT NULL DEFAULT 0,
    efficiency          DECIMAL(5,2)    NOT NULL DEFAULT 100,
    available_hours_day DECIMAL(5,2)    NOT NULL DEFAULT 24,
    max_simultaneous    INTEGER         NOT NULL DEFAULT 1,
    is_active           BOOLEAN         NOT NULL DEFAULT TRUE,
    created_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);

COMMENT ON TABLE work_centers IS '工作中心';
COMMENT ON COLUMN work_centers.work_center_type IS 'PRODUCTION/ASSEMBLY/TESTING/PACKAGING';

CREATE INDEX idx_wc_tenant_id       ON work_centers (tenant_id);
CREATE INDEX idx_wc_work_center_no  ON work_centers (work_center_no);
CREATE INDEX idx_wc_is_active       ON work_centers (is_active);

-- ---------------------------------------------------------------------------
-- 2. mps_plans (主生产计划 MPS)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS mps_plans (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    plan_no         VARCHAR(64)     NOT NULL,
    plan_name       VARCHAR(256)    NOT NULL,
    plan_period     VARCHAR(64)     NOT NULL DEFAULT 'MONTHLY',
    material_id     BIGINT          NOT NULL,
    planned_qty     DECIMAL(18,4)   NOT NULL DEFAULT 0,
    plan_start      TIMESTAMPTZ,
    plan_end        TIMESTAMPTZ,
    status          VARCHAR(32)     NOT NULL DEFAULT 'DRAFT',
    frozen          BOOLEAN         NOT NULL DEFAULT FALSE,
    created_by      BIGINT,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

COMMENT ON TABLE mps_plans IS '主生产计划';
COMMENT ON COLUMN mps_plans.plan_period IS 'DAILY/WEEKLY/MONTHLY';
COMMENT ON COLUMN mps_plans.status IS 'DRAFT/CONFIRMED/RELEASED/FROZEN/CLOSED';

CREATE INDEX idx_mps_tenant_id   ON mps_plans (tenant_id);
CREATE INDEX idx_mps_material_id ON mps_plans (material_id);
CREATE INDEX idx_mps_status      ON mps_plans (status);
CREATE INDEX idx_mps_plan_start  ON mps_plans (plan_start);

-- ---------------------------------------------------------------------------
-- 3. mrp_plans (物料需求计划 MRP)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS mrp_plans (
    id                  BIGSERIAL       PRIMARY KEY,
    tenant_id           BIGINT          NOT NULL,
    mps_plan_id         BIGINT,
    material_id         BIGINT          NOT NULL,
    gross_requirement   DECIMAL(18,4)   NOT NULL DEFAULT 0,
    scheduled_receipt   DECIMAL(18,4)   NOT NULL DEFAULT 0,
    projected_on_hand   DECIMAL(18,4)   NOT NULL DEFAULT 0,
    net_requirement     DECIMAL(18,4)   NOT NULL DEFAULT 0,
    planned_order_qty   DECIMAL(18,4)   NOT NULL DEFAULT 0,
    planned_start       TIMESTAMPTZ,
    planned_end         TIMESTAMPTZ,
    lead_time_days      INTEGER         NOT NULL DEFAULT 0,
    lot_size            DECIMAL(18,4),
    status              VARCHAR(32)     NOT NULL DEFAULT 'DRAFT',
    period_start        DATE            NOT NULL,
    period_end          DATE            NOT NULL,
    created_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ,

    CONSTRAINT fk_mrp_mps FOREIGN KEY (mps_plan_id)
        REFERENCES mps_plans (id) ON DELETE SET NULL
);

COMMENT ON TABLE mrp_plans IS '物料需求计划';

CREATE INDEX idx_mrp_tenant_id    ON mrp_plans (tenant_id);
CREATE INDEX idx_mrp_mps_plan_id  ON mrp_plans (mps_plan_id);
CREATE INDEX idx_mrp_material_id  ON mrp_plans (material_id);
CREATE INDEX idx_mrp_status       ON mrp_plans (status);
CREATE INDEX idx_mrp_period_start ON mrp_plans (period_start);

-- ---------------------------------------------------------------------------
-- 4. schedule_jobs (排程作业)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS schedule_jobs (
    id                  BIGSERIAL       PRIMARY KEY,
    tenant_id           BIGINT          NOT NULL,
    order_id            BIGINT          NOT NULL,
    work_center_id      BIGINT          NOT NULL,
    job_status          VARCHAR(32)     NOT NULL DEFAULT 'PLANNED',
    priority            INTEGER         NOT NULL DEFAULT 0,
    scheduled_start     TIMESTAMPTZ,
    scheduled_end       TIMESTAMPTZ,
    actual_start        TIMESTAMPTZ,
    actual_end          TIMESTAMPTZ,
    setup_time_min      INTEGER         NOT NULL DEFAULT 0,
    process_time_min    INTEGER         NOT NULL DEFAULT 0,
    resource_group      VARCHAR(128),
    predecessor_job_id  BIGINT,
    created_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ,

    CONSTRAINT fk_sj_work_center FOREIGN KEY (work_center_id)
        REFERENCES work_centers (id) ON DELETE CASCADE
);

COMMENT ON TABLE schedule_jobs IS '排程作业';
COMMENT ON COLUMN schedule_jobs.job_status IS 'PLANNED/DISPATCHED/IN_PROGRESS/COMPLETED/DELAYED';

CREATE INDEX idx_sj_tenant_id       ON schedule_jobs (tenant_id);
CREATE INDEX idx_sj_order_id        ON schedule_jobs (order_id);
CREATE INDEX idx_sj_work_center_id  ON schedule_jobs (work_center_id);
CREATE INDEX idx_sj_job_status      ON schedule_jobs (job_status);
CREATE INDEX idx_sj_scheduled_start ON schedule_jobs (scheduled_start);
CREATE INDEX idx_sj_scheduled_end   ON schedule_jobs (scheduled_end);

-- ---------------------------------------------------------------------------
-- 5. schedule_constraints (排程约束)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS schedule_constraints (
    id                  BIGSERIAL       PRIMARY KEY,
    tenant_id           BIGINT          NOT NULL,
    constraint_name     VARCHAR(256)    NOT NULL,
    constraint_type     VARCHAR(64)     NOT NULL,
    work_center_id      BIGINT,
    material_id         BIGINT,
    start_time          TIME,
    end_time            TIME,
    constraint_value    VARCHAR(256),
    valid_from          DATE,
    valid_to            DATE,
    is_active           BOOLEAN         NOT NULL DEFAULT TRUE,
    created_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ,

    CONSTRAINT fk_sc_work_center FOREIGN KEY (work_center_id)
        REFERENCES work_centers (id) ON DELETE SET NULL
);

COMMENT ON TABLE schedule_constraints IS '排程约束';
COMMENT ON COLUMN schedule_constraints.constraint_type IS 'SHIFT/TIME_WINDOW/CAPACITY/MATERIAL_AVAILABLE';

CREATE INDEX idx_sc_tenant_id      ON schedule_constraints (tenant_id);
CREATE INDEX idx_sc_work_center_id ON schedule_constraints (work_center_id);
CREATE INDEX idx_sc_is_active      ON schedule_constraints (is_active);
CREATE INDEX idx_sc_valid_range    ON schedule_constraints (valid_from, valid_to);

-- ---------------------------------------------------------------------------
-- 6. changeovers (换型/切换记录)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS changeovers (
    id                  BIGSERIAL       PRIMARY KEY,
    tenant_id           BIGINT          NOT NULL,
    work_center_id      BIGINT          NOT NULL,
    from_material_id    BIGINT,
    to_material_id      BIGINT,
    changeover_time_min INTEGER         NOT NULL DEFAULT 0,
    changeover_type     VARCHAR(64)     NOT NULL DEFAULT 'SETUP',
    is_active           BOOLEAN         NOT NULL DEFAULT TRUE,
    created_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ,

    CONSTRAINT fk_co_work_center FOREIGN KEY (work_center_id)
        REFERENCES work_centers (id) ON DELETE CASCADE
);

COMMENT ON TABLE changeovers IS '换型/切换矩阵';
COMMENT ON COLUMN changeovers.changeover_type IS 'SETUP/CLEANING/QUALIFICATION';

CREATE INDEX idx_co_tenant_id      ON changeovers (tenant_id);
CREATE INDEX idx_co_work_center_id ON changeovers (work_center_id);
CREATE INDEX idx_co_from_material  ON changeovers (from_material_id);
CREATE INDEX idx_co_to_material    ON changeovers (to_material_id);
