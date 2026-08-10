-- ===========================================================================
-- Trace Service - Initial Database Migration
-- Description: Product Traceability core tables
-- Version: 1.0.0
-- ===========================================================================

-- ---------------------------------------------------------------------------
-- 1. serial_numbers (序列号管理)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS serial_numbers (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    material_id     BIGINT          NOT NULL,
    serial_no       VARCHAR(128)    NOT NULL,
    batch_no        VARCHAR(128),
    lot_no          VARCHAR(128),
    production_date DATE,
    expiry_date     DATE,
    status          VARCHAR(32)     NOT NULL DEFAULT 'ACTIVE',
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

COMMENT ON TABLE serial_numbers IS '序列号';
COMMENT ON COLUMN serial_numbers.status IS 'ACTIVE/QUARANTINED/SCRAPPED/SHIPPED';

CREATE INDEX idx_sn_tenant_id   ON serial_numbers (tenant_id);
CREATE INDEX idx_sn_material_id ON serial_numbers (material_id);
CREATE INDEX idx_sn_serial_no   ON serial_numbers (serial_no);
CREATE INDEX idx_sn_batch_no    ON serial_numbers (batch_no);
CREATE INDEX idx_sn_lot_no      ON serial_numbers (lot_no);
CREATE INDEX idx_sn_status      ON serial_numbers (status);

-- ---------------------------------------------------------------------------
-- 2. trace_records (追溯记录)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS trace_records (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    serial_id       BIGINT          NOT NULL,
    trace_type      VARCHAR(64)     NOT NULL,
    event           VARCHAR(128)    NOT NULL,
    event_time      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    operator_id     BIGINT,
    workstation_id  BIGINT,
    order_id        BIGINT,
    detail          JSONB,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_tr_serial FOREIGN KEY (serial_id)
        REFERENCES serial_numbers (id) ON DELETE CASCADE
);

COMMENT ON TABLE trace_records IS '追溯记录';
COMMENT ON COLUMN trace_records.trace_type IS 'RECEIVE/ISSUE/PRODUCE/INSPECT/SHIP/RETURN';

CREATE INDEX idx_tr_tenant_id   ON trace_records (tenant_id);
CREATE INDEX idx_tr_serial_id   ON trace_records (serial_id);
CREATE INDEX idx_tr_event_time  ON trace_records (event_time);
CREATE INDEX idx_tr_order_id    ON trace_records (order_id);
CREATE INDEX idx_tr_detail      ON trace_records USING GIN (detail);

-- ---------------------------------------------------------------------------
-- 3. trace_links (追溯链路)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS trace_links (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    parent_serial_id BIGINT         NOT NULL,
    child_serial_id BIGINT          NOT NULL,
    link_type       VARCHAR(64)     NOT NULL DEFAULT 'ASSEMBLY',
    quantity        DECIMAL(18,4)   NOT NULL DEFAULT 1,
    order_id        BIGINT,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_tl_parent FOREIGN KEY (parent_serial_id)
        REFERENCES serial_numbers (id) ON DELETE CASCADE,
    CONSTRAINT fk_tl_child FOREIGN KEY (child_serial_id)
        REFERENCES serial_numbers (id) ON DELETE CASCADE,
    CONSTRAINT uk_tl_link UNIQUE (parent_serial_id, child_serial_id, link_type)
);

COMMENT ON TABLE trace_links IS '追溯链路 - 物料树形结构';
COMMENT ON COLUMN trace_links.link_type IS 'ASSEMBLY/DISASSEMBLY/REWORK/SPLIT/MERGE';

CREATE INDEX idx_tl_tenant_id        ON trace_links (tenant_id);
CREATE INDEX idx_tl_parent_serial_id ON trace_links (parent_serial_id);
CREATE INDEX idx_tl_child_serial_id  ON trace_links (child_serial_id);
CREATE INDEX idx_tl_order_id         ON trace_links (order_id);

-- ---------------------------------------------------------------------------
-- 4. data_points (数据采集点)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS data_points (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    serial_id       BIGINT,
    data_point_code VARCHAR(64)     NOT NULL,
    data_point_name VARCHAR(256)    NOT NULL,
    data_type       VARCHAR(64)     NOT NULL DEFAULT 'STRING',
    unit            VARCHAR(32),
    workstation_id  BIGINT,
    process_step    VARCHAR(128),
    is_mandatory    BOOLEAN         NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

COMMENT ON TABLE data_points IS '数据采集点定义';
COMMENT ON COLUMN data_points.data_type IS 'STRING/NUMERIC/DATE/BOOLEAN/JSON';

CREATE INDEX idx_dp_tenant_id       ON data_points (tenant_id);
CREATE INDEX idx_dp_data_point_code ON data_points (data_point_code);
CREATE INDEX idx_dp_workstation_id  ON data_points (workstation_id);

-- ---------------------------------------------------------------------------
-- 5. collect_records (采集记录)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS collect_records (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    serial_id       BIGINT          NOT NULL,
    data_point_id   BIGINT          NOT NULL,
    value           TEXT            NOT NULL,
    collected_at    TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    collected_by    BIGINT,
    equipment_id    BIGINT,
    trace_record_id BIGINT,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT fk_cr_serial FOREIGN KEY (serial_id)
        REFERENCES serial_numbers (id) ON DELETE CASCADE,
    CONSTRAINT fk_cr_data_point FOREIGN KEY (data_point_id)
        REFERENCES data_points (id) ON DELETE CASCADE
);

COMMENT ON TABLE collect_records IS '数据采集记录';

CREATE INDEX idx_col_tenant_id      ON collect_records (tenant_id);
CREATE INDEX idx_col_serial_id      ON collect_records (serial_id);
CREATE INDEX idx_col_data_point_id  ON collect_records (data_point_id);
CREATE INDEX idx_col_collected_at   ON collect_records (collected_at);
CREATE INDEX idx_col_equipment_id   ON collect_records (equipment_id);

-- ---------------------------------------------------------------------------
-- 6. scan_logs (扫描日志)
-- ---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS scan_logs (
    id              BIGSERIAL       PRIMARY KEY,
    tenant_id       BIGINT          NOT NULL,
    serial_id       BIGINT,
    scan_code       VARCHAR(256)    NOT NULL,
    scan_type       VARCHAR(64)     NOT NULL DEFAULT 'BARCODE',
    scan_action     VARCHAR(64)     NOT NULL,
    workstation_id  BIGINT,
    operator_id     BIGINT,
    scanned_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    is_valid        BOOLEAN,
    error_message   VARCHAR(512),
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

COMMENT ON TABLE scan_logs IS '扫描日志';
COMMENT ON COLUMN scan_logs.scan_type IS 'BARCODE/QR/RFID/MANUAL';

CREATE INDEX idx_sl_tenant_id       ON scan_logs (tenant_id);
CREATE INDEX idx_sl_serial_id       ON scan_logs (serial_id);
CREATE INDEX idx_sl_scan_code       ON scan_logs (scan_code);
CREATE INDEX idx_sl_scanned_at      ON scan_logs (scanned_at);
CREATE INDEX idx_sl_workstation_id  ON scan_logs (workstation_id);
