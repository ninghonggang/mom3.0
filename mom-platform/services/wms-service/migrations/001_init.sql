-- ===========================================================================
-- WMS Service - Initial Database Migration
-- Description: Warehouse Management System core tables
-- Version: 1.0.0
-- ===========================================================================

-- ---------------------------------------------------------------------------
-- 1. warehouses (仓库)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS warehouses (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    warehouse_no    VARCHAR(64)     NOT NULL,
    warehouse_name  VARCHAR(256)    NOT NULL,
    warehouse_type  VARCHAR(64)     NOT NULL DEFAULT 'RAW_MATERIAL',
    address         VARCHAR(512),
    is_active       BOOLEAN         NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

COMMENT ON TABLE warehouses IS '仓库';
COMMENT ON COLUMN warehouses.warehouse_type IS 'RAW_MATERIAL/FINISHED_GOODS/WIP/RETURNS/SPARE_PARTS';

CREATE INDEX idx_wh_tenant_id     ON warehouses (tenant_id);
CREATE INDEX idx_wh_warehouse_no  ON warehouses (warehouse_no);
CREATE INDEX idx_wh_is_active     ON warehouses (is_active);

-- ---------------------------------------------------------------------------
-- 2. locations (库位)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS locations (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    warehouse_id    BIGINT          NOT NULL,
    location_no     VARCHAR(64)     NOT NULL,
    location_type   VARCHAR(64)     NOT NULL DEFAULT 'FLOOR',
    zone            VARCHAR(64),
    aisle           VARCHAR(32),
    rack            VARCHAR(32),
    level           VARCHAR(32),
    max_weight      DECIMAL(18,2),
    max_volume      DECIMAL(18,4),
    is_active       BOOLEAN         NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_loc_warehouse FOREIGN KEY (warehouse_id)
        REFERENCES warehouses (id) ON DELETE CASCADE
);

COMMENT ON TABLE locations IS '库位';
COMMENT ON COLUMN locations.location_type IS 'FLOOR/RACK/SHELF/BIN';

CREATE INDEX idx_loc_tenant_id    ON locations (tenant_id);
CREATE INDEX idx_loc_warehouse_id ON locations (warehouse_id);
CREATE INDEX idx_loc_location_no  ON locations (location_no);
CREATE INDEX idx_loc_is_active    ON locations (is_active);

-- ---------------------------------------------------------------------------
-- 3. inventory_balances (库存余额)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS inventory_balances (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    warehouse_id    BIGINT          NOT NULL,
    location_id     BIGINT,
    material_id     BIGINT          NOT NULL,
    batch_no        VARCHAR(128),
    lot_no          VARCHAR(128),
    on_hand_qty     DECIMAL(18,4)   NOT NULL DEFAULT 0,
    reserved_qty    DECIMAL(18,4)   NOT NULL DEFAULT 0,
    available_qty   DECIMAL(18,4)   NOT NULL DEFAULT 0,
    uom             VARCHAR(32)     NOT NULL DEFAULT 'PCS',
    last_counted_at TIMESTAMPTZ,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_ib_warehouse FOREIGN KEY (warehouse_id)
        REFERENCES warehouses (id) ON DELETE CASCADE
);

COMMENT ON TABLE inventory_balances IS '库存余额';

CREATE INDEX idx_ib_tenant_id    ON inventory_balances (tenant_id);
CREATE INDEX idx_ib_warehouse_id ON inventory_balances (warehouse_id);
CREATE INDEX idx_ib_location_id  ON inventory_balances (location_id);
CREATE INDEX idx_ib_material_id  ON inventory_balances (material_id);
CREATE INDEX idx_ib_batch_no     ON inventory_balances (batch_no);

-- ---------------------------------------------------------------------------
-- 4. receive_orders (收货单)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS receive_orders (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    receive_no      VARCHAR(64)     NOT NULL,
    warehouse_id    BIGINT          NOT NULL,
    supplier_id     BIGINT,
    reference_no    VARCHAR(128),
    receive_type    VARCHAR(64)     NOT NULL DEFAULT 'PURCHASE',
    status          VARCHAR(32)     NOT NULL DEFAULT 'PENDING',
    expected_at     TIMESTAMPTZ,
    received_at     TIMESTAMPTZ,
    received_by     BIGINT,
    remark          TEXT,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_ro_warehouse FOREIGN KEY (warehouse_id)
        REFERENCES warehouses (id) ON DELETE CASCADE
);

COMMENT ON TABLE receive_orders IS '收货单';
COMMENT ON COLUMN receive_orders.receive_type IS 'PURCHASE/RETURN/TRANSFER';
COMMENT ON COLUMN receive_orders.status IS 'PENDING/PARTIAL/COMPLETED/CLOSED';

CREATE INDEX idx_ro_wh_tenant_id    ON receive_orders (tenant_id);
CREATE INDEX idx_ro_wh_warehouse_id ON receive_orders (warehouse_id);
CREATE INDEX idx_ro_wh_status       ON receive_orders (status);
CREATE INDEX idx_ro_wh_supplier_id  ON receive_orders (supplier_id);
CREATE INDEX idx_ro_wh_received_at  ON receive_orders (received_at);

-- ---------------------------------------------------------------------------
-- 5. receive_order_lines (收货单明细)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS receive_order_lines (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    receive_id      BIGINT          NOT NULL,
    line_no         INTEGER         NOT NULL DEFAULT 1,
    material_id     BIGINT          NOT NULL,
    expected_qty    DECIMAL(18,4)   NOT NULL DEFAULT 0,
    received_qty    DECIMAL(18,4)   NOT NULL DEFAULT 0,
    batch_no        VARCHAR(128),
    uom             VARCHAR(32)     NOT NULL DEFAULT 'PCS',
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_rol_receive FOREIGN KEY (receive_id)
        REFERENCES receive_orders (id) ON DELETE CASCADE
);

COMMENT ON TABLE receive_order_lines IS '收货单明细';

CREATE INDEX idx_rol_tenant_id   ON receive_order_lines (tenant_id);
CREATE INDEX idx_rol_receive_id  ON receive_order_lines (receive_id);
CREATE INDEX idx_rol_material_id ON receive_order_lines (material_id);

-- ---------------------------------------------------------------------------
-- 6. delivery_orders (发货单)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS delivery_orders (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    delivery_no     VARCHAR(64)     NOT NULL,
    warehouse_id    BIGINT          NOT NULL,
    customer_id     BIGINT,
    order_id        BIGINT,
    delivery_type   VARCHAR(64)     NOT NULL DEFAULT 'SALES',
    status          VARCHAR(32)     NOT NULL DEFAULT 'PENDING',
    expected_at     TIMESTAMPTZ,
    delivered_at    TIMESTAMPTZ,
    delivered_by    BIGINT,
    remark          TEXT,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_do_warehouse FOREIGN KEY (warehouse_id)
        REFERENCES warehouses (id) ON DELETE CASCADE
);

COMMENT ON TABLE delivery_orders IS '发货单';
COMMENT ON COLUMN delivery_orders.delivery_type IS 'SALES/TRANSFER/SAMPLE';
COMMENT ON COLUMN delivery_orders.status IS 'PENDING/PARTIAL/COMPLETED/CLOSED';

CREATE INDEX idx_do_tenant_id     ON delivery_orders (tenant_id);
CREATE INDEX idx_do_warehouse_id  ON delivery_orders (warehouse_id);
CREATE INDEX idx_do_status        ON delivery_orders (status);
CREATE INDEX idx_do_customer_id   ON delivery_orders (customer_id);
CREATE INDEX idx_do_delivered_at  ON delivery_orders (delivered_at);

-- ---------------------------------------------------------------------------
-- 7. delivery_order_lines (发货单明细)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS delivery_order_lines (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    delivery_id     BIGINT          NOT NULL,
    line_no         INTEGER         NOT NULL DEFAULT 1,
    material_id     BIGINT          NOT NULL,
    ordered_qty     DECIMAL(18,4)   NOT NULL DEFAULT 0,
    delivered_qty   DECIMAL(18,4)   NOT NULL DEFAULT 0,
    batch_no        VARCHAR(128),
    uom             VARCHAR(32)     NOT NULL DEFAULT 'PCS',
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_dol_delivery FOREIGN KEY (delivery_id)
        REFERENCES delivery_orders (id) ON DELETE CASCADE
);

COMMENT ON TABLE delivery_order_lines IS '发货单明细';

CREATE INDEX idx_dol_tenant_id   ON delivery_order_lines (tenant_id);
CREATE INDEX idx_dol_delivery_id ON delivery_order_lines (delivery_id);
CREATE INDEX idx_dol_material_id ON delivery_order_lines (material_id);

-- ---------------------------------------------------------------------------
-- 8. putaway_records (上架记录)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS putaway_records (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    receive_line_id BIGINT          NOT NULL,
    location_id     BIGINT          NOT NULL,
    putaway_qty     DECIMAL(18,4)   NOT NULL DEFAULT 0,
    putaway_by      BIGINT,
    putaway_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_pa_receive_line FOREIGN KEY (receive_line_id)
        REFERENCES receive_order_lines (id) ON DELETE CASCADE,
    CONSTRAINT fk_pa_location FOREIGN KEY (location_id)
        REFERENCES locations (id) ON DELETE CASCADE
);

COMMENT ON TABLE putaway_records IS '上架记录';

CREATE INDEX idx_pa_tenant_id       ON putaway_records (tenant_id);
CREATE INDEX idx_pa_receive_line_id ON putaway_records (receive_line_id);
CREATE INDEX idx_pa_location_id     ON putaway_records (location_id);
CREATE INDEX idx_pa_putaway_at      ON putaway_records (putaway_at);

-- ---------------------------------------------------------------------------
-- 9. pick_records (拣货记录)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS pick_records (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    delivery_line_id BIGINT         NOT NULL,
    location_id     BIGINT          NOT NULL,
    pick_qty        DECIMAL(18,4)   NOT NULL DEFAULT 0,
    picked_by       BIGINT,
    picked_at       TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_pk_delivery_line FOREIGN KEY (delivery_line_id)
        REFERENCES delivery_order_lines (id) ON DELETE CASCADE,
    CONSTRAINT fk_pk_location FOREIGN KEY (location_id)
        REFERENCES locations (id) ON DELETE CASCADE
);

COMMENT ON TABLE pick_records IS '拣货记录';

CREATE INDEX idx_pk_tenant_id        ON pick_records (tenant_id);
CREATE INDEX idx_pk_delivery_line_id ON pick_records (delivery_line_id);
CREATE INDEX idx_pk_location_id      ON pick_records (location_id);
CREATE INDEX idx_pk_picked_at        ON pick_records (picked_at);

-- ---------------------------------------------------------------------------
-- 10. count_plans (盘点计划)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS count_plans (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    count_no        VARCHAR(64)     NOT NULL,
    warehouse_id    BIGINT          NOT NULL,
    count_type      VARCHAR(64)     NOT NULL DEFAULT 'FULL',
    status          VARCHAR(32)     NOT NULL DEFAULT 'PENDING',
    planned_at      TIMESTAMPTZ,
    counted_at      TIMESTAMPTZ,
    created_by      BIGINT,
    remark          TEXT,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_cp_warehouse FOREIGN KEY (warehouse_id)
        REFERENCES warehouses (id) ON DELETE CASCADE
);

COMMENT ON TABLE count_plans IS '盘点计划';
COMMENT ON COLUMN count_plans.count_type IS 'FULL/CYCLE/SPOT';
COMMENT ON COLUMN count_plans.status IS 'PENDING/IN_PROGRESS/COMPLETED/CLOSED';

CREATE INDEX idx_cp_tenant_id    ON count_plans (tenant_id);
CREATE INDEX idx_cp_warehouse_id ON count_plans (warehouse_id);
CREATE INDEX idx_cp_status       ON count_plans (status);
CREATE INDEX idx_cp_planned_at   ON count_plans (planned_at);

-- ---------------------------------------------------------------------------
-- 11. count_records (盘点记录)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS count_records (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    count_plan_id   BIGINT          NOT NULL,
    location_id     BIGINT          NOT NULL,
    material_id     BIGINT          NOT NULL,
    batch_no        VARCHAR(128),
    system_qty      DECIMAL(18,4)   NOT NULL DEFAULT 0,
    counted_qty     DECIMAL(18,4),
    variance_qty    DECIMAL(18,4),
    counted_by      BIGINT,
    counted_at      TIMESTAMPTZ,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_cr_plan FOREIGN KEY (count_plan_id)
        REFERENCES count_plans (id) ON DELETE CASCADE,
    CONSTRAINT fk_cr_location FOREIGN KEY (location_id)
        REFERENCES locations (id) ON DELETE CASCADE
);

COMMENT ON TABLE count_records IS '盘点记录';

CREATE INDEX idx_cr_tenant_id    ON count_records (tenant_id);
CREATE INDEX idx_cr_count_plan_id ON count_records (count_plan_id);
CREATE INDEX idx_cr_location_id  ON count_records (location_id);
CREATE INDEX idx_cr_material_id  ON count_records (material_id);
CREATE INDEX idx_cr_counted_at   ON count_records (counted_at);
