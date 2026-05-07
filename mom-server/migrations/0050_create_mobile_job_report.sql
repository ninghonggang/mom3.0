-- Mobile Job Report Table Migration

CREATE TABLE IF NOT EXISTS mobile_job_report (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    workshop_id BIGINT,
    workshop_name VARCHAR(100),
    production_line_id BIGINT,
    production_line_name VARCHAR(100),
    workstation_id BIGINT,
    workstation_name VARCHAR(100),
    order_id BIGINT,
    order_code VARCHAR(64),
    process_id BIGINT,
    process_name VARCHAR(100),
    employee_id BIGINT,
    employee_name VARCHAR(64),
    reported_quantity DECIMAL(15, 2) DEFAULT 0,
    qualified_quantity DECIMAL(15, 2) DEFAULT 0,
    defective_quantity DECIMAL(15, 2) DEFAULT 0,
    work_minutes INTEGER DEFAULT 0,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    report_type INTEGER DEFAULT 1,
    status INTEGER DEFAULT 1,
    remarks TEXT,
    confirm_by BIGINT,
    confirm_at TIMESTAMP,
    audit_by BIGINT,
    audit_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_mobile_job_report_tenant ON mobile_job_report(tenant_id);
CREATE INDEX IF NOT EXISTS idx_mobile_job_report_workshop ON mobile_job_report(workshop_id);
CREATE INDEX IF NOT EXISTS idx_mobile_job_report_order ON mobile_job_report(order_id);
CREATE INDEX IF NOT EXISTS idx_mobile_job_report_employee ON mobile_job_report(employee_id);