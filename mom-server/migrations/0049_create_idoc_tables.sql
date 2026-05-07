-- IDOC Tables Migration

-- IDOC记录表
CREATE TABLE IF NOT EXISTS idoc_record (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    idoc_number VARCHAR(64) NOT NULL,
    idoc_type VARCHAR(30),
    direction INTEGER DEFAULT 1,
    status INTEGER DEFAULT 1,
    partner_type VARCHAR(20),
    partner_no VARCHAR(20),
    message_type VARCHAR(30),
    reference_no VARCHAR(64),
    raw_content TEXT,
    parsed_data TEXT,
    error_message TEXT,
    retry_count INTEGER DEFAULT 0,
    processed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_idoc_record_tenant ON idoc_record(tenant_id);
CREATE INDEX IF NOT EXISTS idx_idoc_record_idoc_type ON idoc_record(idoc_type);
CREATE INDEX IF NOT EXISTS idx_idoc_record_status ON idoc_record(status);

-- IDOC类型配置表
CREATE TABLE IF NOT EXISTS idoc_type_config (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    idoc_type VARCHAR(30) NOT NULL,
    message_type VARCHAR(30),
    description VARCHAR(200),
    endpoint VARCHAR(200),
    is_active INTEGER DEFAULT 1,
    mapping_rule TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_idoc_type_config_tenant ON idoc_type_config(tenant_id);
CREATE UNIQUE INDEX IF NOT EXISTS idx_idoc_type_config_type ON idoc_type_config(idoc_type);