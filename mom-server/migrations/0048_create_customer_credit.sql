-- Customer Credit Table Migration

CREATE TABLE IF NOT EXISTS customer_credit (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    customer_id BIGINT NOT NULL,
    customer_code VARCHAR(64),
    customer_name VARCHAR(200),
    credit_limit DECIMAL(15, 2) DEFAULT 0,
    used_credit DECIMAL(15, 2) DEFAULT 0,
    available_credit DECIMAL(15, 2) DEFAULT 0,
    credit_level VARCHAR(20),
    payment_days INTEGER DEFAULT 0,
    risk_level INTEGER DEFAULT 1,
    alert_threshold DECIMAL(5, 2) DEFAULT 0.8,
    total_orders INTEGER DEFAULT 0,
    total_amount DECIMAL(15, 2) DEFAULT 0,
    overdue_amount DECIMAL(15, 2) DEFAULT 0,
    blacklist INTEGER DEFAULT 0,
    status INTEGER DEFAULT 1,
    remarks TEXT,
    last_trade_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_customer_credit_tenant ON customer_credit(tenant_id);
CREATE INDEX IF NOT EXISTS idx_customer_credit_customer ON customer_credit(customer_id);