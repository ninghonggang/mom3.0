-- Quality Certificate Table Migration

CREATE TABLE IF NOT EXISTS quality_certificate (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    cert_code VARCHAR(64),
    cert_type VARCHAR(32),
    order_id BIGINT,
    order_code VARCHAR(64),
    product_id BIGINT,
    product_code VARCHAR(64),
    product_name VARCHAR(200),
    batch_no VARCHAR(64),
    quantity DECIMAL(15, 2) DEFAULT 0,
    unit VARCHAR(20),
    inspector VARCHAR(64),
    inspect_date TIMESTAMP,
    result INTEGER DEFAULT 1,
    status INTEGER DEFAULT 1,
    issue_date TIMESTAMP,
    expiry_date TIMESTAMP,
    remarks TEXT,
    attachments TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_quality_certificate_tenant ON quality_certificate(tenant_id);
CREATE INDEX IF NOT EXISTS idx_quality_certificate_order ON quality_certificate(order_id);
CREATE INDEX IF NOT EXISTS idx_quality_certificate_product ON quality_certificate(product_id);
CREATE INDEX IF NOT EXISTS idx_quality_certificate_cert_code ON quality_certificate(cert_code);