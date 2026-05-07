-- VMI Tables Migration
-- Run this to create VMI tables

-- VMI Vendor table
CREATE TABLE IF NOT EXISTS vmi_vendor (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    vendor_id BIGINT NOT NULL,
    vendor_code VARCHAR(64),
    vendor_name VARCHAR(200),
    warehouse_id BIGINT NOT NULL,
    warehouse_name VARCHAR(100),
    contact VARCHAR(100),
    phone VARCHAR(50),
    min_stock DECIMAL(15, 2) DEFAULT 0,
    max_stock DECIMAL(15, 2) DEFAULT 0,
    replenish_cycle INTEGER DEFAULT 7,
    is_active INTEGER DEFAULT 1,
    status INTEGER DEFAULT 1,
    remarks TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_vmi_vendor_tenant ON vmi_vendor(tenant_id);

-- VMI Material table
CREATE TABLE IF NOT EXISTS vmi_material (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    vendor_id BIGINT NOT NULL,
    material_id BIGINT NOT NULL,
    material_code VARCHAR(64),
    material_name VARCHAR(200),
    unit VARCHAR(20),
    min_stock DECIMAL(15, 2) DEFAULT 0,
    max_stock DECIMAL(15, 2) DEFAULT 0,
    current_stock DECIMAL(15, 2) DEFAULT 0,
    available_stock DECIMAL(15, 2) DEFAULT 0,
    consume_qty DECIMAL(15, 2) DEFAULT 0,
    last_consume_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_vmi_material_tenant ON vmi_material(tenant_id);

-- VMI Transaction table
CREATE TABLE IF NOT EXISTS vmi_transaction (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    vendor_id BIGINT NOT NULL,
    material_id BIGINT NOT NULL,
    transaction_type INTEGER NOT NULL,
    quantity DECIMAL(15, 2) NOT NULL,
    before_qty DECIMAL(15, 2) NOT NULL,
    after_qty DECIMAL(15, 2) NOT NULL,
    reference_no VARCHAR(100),
    operator_id BIGINT,
    operator_name VARCHAR(100),
    remarks TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_vmi_transaction_tenant ON vmi_transaction(tenant_id);
CREATE INDEX IF NOT EXISTS idx_vmi_transaction_created ON vmi_transaction(created_at);
