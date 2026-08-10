-- ===========================================================================
-- MDM Service - Initial Database Migration
-- Description: Master Data Management core tables
-- Version: 1.0.0
-- ===========================================================================

-- ---------------------------------------------------------------------------
-- 1. materials (物料主数据)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS materials (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    material_no     VARCHAR(64)     NOT NULL,
    material_name   VARCHAR(256)    NOT NULL,
    material_type   VARCHAR(64)     NOT NULL DEFAULT 'RAW',
    material_group  VARCHAR(128),
    spec            VARCHAR(256),
    model           VARCHAR(128),
    base_uom        VARCHAR(32)     NOT NULL DEFAULT 'PCS',
    unit_weight     DECIMAL(18,4),
    unit_volume     DECIMAL(18,4),
    shelf_life_days INTEGER,
    safety_stock    DECIMAL(18,4)   NOT NULL DEFAULT 0,
    min_order_qty   DECIMAL(18,4)   NOT NULL DEFAULT 0,
    status          VARCHAR(32)     NOT NULL DEFAULT 'ACTIVE',
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

COMMENT ON TABLE materials IS '物料主数据';
COMMENT ON COLUMN materials.material_type IS 'RAW/WIP/FINISHED/PACKAGING/SERVICE';
COMMENT ON COLUMN materials.status IS 'ACTIVE/INACTIVE/OBSOLETE';

CREATE INDEX idx_mat_tenant_id   ON materials (tenant_id);
CREATE INDEX idx_mat_material_no ON materials (material_no);
CREATE INDEX idx_mat_status      ON materials (status);
CREATE INDEX idx_mat_type        ON materials (material_type);

-- ---------------------------------------------------------------------------
-- 2. boms (BOM主表)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS boms (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    bom_no          VARCHAR(64)     NOT NULL,
    bom_name        VARCHAR(256)    NOT NULL,
    material_id     BIGINT          NOT NULL,
    bom_type        VARCHAR(64)     NOT NULL DEFAULT 'MANUFACTURING',
    version         VARCHAR(32)     NOT NULL DEFAULT '1.0',
    quantity        DECIMAL(18,4)   NOT NULL DEFAULT 1,
    status          VARCHAR(32)     NOT NULL DEFAULT 'ACTIVE',
    valid_from      TIMESTAMPTZ,
    valid_to        TIMESTAMPTZ,
    is_default      BOOLEAN         NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

COMMENT ON TABLE boms IS 'BOM主表 - 物料清单';
COMMENT ON COLUMN boms.bom_type IS 'MANUFACTURING/ENGINEERING/SERVICE';
COMMENT ON COLUMN boms.status IS 'ACTIVE/INACTIVE/ARCHIVED';

CREATE INDEX idx_bom_tenant_id   ON boms (tenant_id);
CREATE INDEX idx_bom_material_id ON boms (material_id);
CREATE INDEX idx_bom_status      ON boms (status);
CREATE INDEX idx_bom_is_default  ON boms (is_default);

-- ---------------------------------------------------------------------------
-- 3. bom_items (BOM明细)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS bom_items (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    bom_id          BIGINT          NOT NULL,
    item_no         INTEGER         NOT NULL DEFAULT 1,
    component_id    BIGINT          NOT NULL,
    quantity        DECIMAL(18,4)   NOT NULL DEFAULT 1,
    uom             VARCHAR(32)     NOT NULL DEFAULT 'PCS',
    scrap_rate      DECIMAL(5,4)    NOT NULL DEFAULT 0,
    is_critical     BOOLEAN         NOT NULL DEFAULT FALSE,
    remark          VARCHAR(512),
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_bi_bom FOREIGN KEY (bom_id)
        REFERENCES boms (id) ON DELETE CASCADE
);

COMMENT ON TABLE bom_items IS 'BOM明细';

CREATE INDEX idx_bi_tenant_id    ON bom_items (tenant_id);
CREATE INDEX idx_bi_bom_id       ON bom_items (bom_id);
CREATE INDEX idx_bi_component_id ON bom_items (component_id);

-- ---------------------------------------------------------------------------
-- 4. workshops (车间)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS workshops (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    workshop_no     VARCHAR(64)     NOT NULL,
    workshop_name   VARCHAR(256)    NOT NULL,
    workshop_type   VARCHAR(64),
    manager_id      BIGINT,
    is_active       BOOLEAN         NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

COMMENT ON TABLE workshops IS '车间';

CREATE INDEX idx_ws_tenant_id    ON workshops (tenant_id);
CREATE INDEX idx_ws_workshop_no  ON workshops (workshop_no);
CREATE INDEX idx_ws_is_active    ON workshops (is_active);

-- ---------------------------------------------------------------------------
-- 5. production_lines (产线)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS production_lines (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    workshop_id     BIGINT          NOT NULL,
    line_no         VARCHAR(64)     NOT NULL,
    line_name       VARCHAR(256)    NOT NULL,
    line_type       VARCHAR(64),
    capacity_per_hour DECIMAL(18,2),
    is_active       BOOLEAN         NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_pl_workshop FOREIGN KEY (workshop_id)
        REFERENCES workshops (id) ON DELETE CASCADE
);

COMMENT ON TABLE production_lines IS '产线';

CREATE INDEX idx_pl_tenant_id   ON production_lines (tenant_id);
CREATE INDEX idx_pl_workshop_id ON production_lines (workshop_id);
CREATE INDEX idx_pl_is_active   ON production_lines (is_active);

-- ---------------------------------------------------------------------------
-- 6. workstations (工位)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS workstations (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    line_id         BIGINT          NOT NULL,
    workstation_no  VARCHAR(64)     NOT NULL,
    workstation_name VARCHAR(256)   NOT NULL,
    sequence        INTEGER         NOT NULL DEFAULT 1,
    is_active       BOOLEAN         NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_wst_line FOREIGN KEY (line_id)
        REFERENCES production_lines (id) ON DELETE CASCADE
);

COMMENT ON TABLE workstations IS '工位';

CREATE INDEX idx_wst_tenant_id      ON workstations (tenant_id);
CREATE INDEX idx_wst_line_id        ON workstations (line_id);
CREATE INDEX idx_wst_workstation_no ON workstations (workstation_no);
CREATE INDEX idx_wst_is_active      ON workstations (is_active);

-- ---------------------------------------------------------------------------
-- 7. customers (客户)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS customers (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    customer_no     VARCHAR(64)     NOT NULL,
    customer_name   VARCHAR(256)    NOT NULL,
    short_name      VARCHAR(128),
    contact_person  VARCHAR(128),
    contact_phone   VARCHAR(32),
    contact_email   VARCHAR(256),
    address         VARCHAR(512),
    status          VARCHAR(32)     NOT NULL DEFAULT 'ACTIVE',
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

COMMENT ON TABLE customers IS '客户主数据';
COMMENT ON COLUMN customers.status IS 'ACTIVE/INACTIVE/BLOCKED';

CREATE INDEX idx_cust_tenant_id   ON customers (tenant_id);
CREATE INDEX idx_cust_customer_no ON customers (customer_no);
CREATE INDEX idx_cust_status      ON customers (status);

-- ---------------------------------------------------------------------------
-- 8. suppliers (供应商)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS suppliers (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    supplier_no     VARCHAR(64)     NOT NULL,
    supplier_name   VARCHAR(256)    NOT NULL,
    short_name      VARCHAR(128),
    contact_person  VARCHAR(128),
    contact_phone   VARCHAR(32),
    contact_email   VARCHAR(256),
    address         VARCHAR(512),
    supplier_type   VARCHAR(64),
    status          VARCHAR(32)     NOT NULL DEFAULT 'ACTIVE',
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

COMMENT ON TABLE suppliers IS '供应商主数据';
COMMENT ON COLUMN suppliers.status IS 'ACTIVE/INACTIVE/BLOCKED';

CREATE INDEX idx_sup_tenant_id   ON suppliers (tenant_id);
CREATE INDEX idx_sup_supplier_no ON suppliers (supplier_no);
CREATE INDEX idx_sup_status      ON suppliers (status);
