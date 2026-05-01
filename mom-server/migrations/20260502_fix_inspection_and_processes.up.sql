-- MOM3.0 Schema Fix: 新增设备点检三表
-- 执行时间: 2026-05-02
-- 修复: eam_inspection_plan / eam_inspection_record / eam_inspection_defect 三表缺失
-- 数据库: mom3_db

-- 1. eam_inspection_plan 点检计划表
CREATE TABLE IF NOT EXISTS eam_inspection_plan (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    plan_no VARCHAR(50) NOT NULL,
    plan_name VARCHAR(200) NOT NULL,
    template_id BIGINT,
    template_name VARCHAR(200),
    equipment_id BIGINT,
    equipment_code VARCHAR(50),
    equipment_name VARCHAR(100),
    workstation_id BIGINT,
    workstation_name VARCHAR(100),
    plan_date TIMESTAMP,
    plan_shift VARCHAR(20),
    plan_start_time VARCHAR(10),
    plan_end_time VARCHAR(10),
    assigned_to BIGINT,
    assigned_name VARCHAR(50),
    status VARCHAR(20) DEFAULT 'PENDING',
    is_generated INT DEFAULT 0,
    workshop_id BIGINT,
    CONSTRAINT idx_tenant_plan_pn UNIQUE (tenant_id, plan_no)
);

-- 2. eam_inspection_record 点检记录表
CREATE TABLE IF NOT EXISTS eam_inspection_record (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    record_no VARCHAR(50) NOT NULL,
    plan_id BIGINT,
    template_id BIGINT,
    template_name VARCHAR(200),
    equipment_id BIGINT,
    equipment_code VARCHAR(50),
    equipment_name VARCHAR(100),
    inspector_id BIGINT,
    inspector_name VARCHAR(50),
    inspection_start_time TIMESTAMP,
    inspection_end_time TIMESTAMP,
    actual_duration INT DEFAULT 0,
    overall_result VARCHAR(20),
    ok_count INT DEFAULT 0,
    ng_count INT DEFAULT 0,
    shift VARCHAR(20),
    shift_name VARCHAR(50),
    location_lat VARCHAR(20),
    location_lng VARCHAR(20),
    signature_data TEXT,
    signature_time TIMESTAMP,
    status VARCHAR(20) DEFAULT 'EXECUTING',
    remark TEXT,
    workshop_id BIGINT,
    CONSTRAINT idx_tenant_record_rn UNIQUE (tenant_id, record_no)
);

-- 3. eam_inspection_defect 点检异常表
CREATE TABLE IF NOT EXISTS eam_inspection_defect (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    defect_no VARCHAR(50) NOT NULL,
    record_id BIGINT,
    result_id BIGINT,
    defect_type VARCHAR(20),
    defect_code VARCHAR(50),
    defect_name VARCHAR(200),
    defect_level VARCHAR(10),
    description TEXT,
    photos TEXT,
    status VARCHAR(20) DEFAULT 'REPORTED',
    reported_by BIGINT,
    reported_time TIMESTAMP,
    assigned_to BIGINT,
    assigned_name VARCHAR(50),
    assignment_time TIMESTAMP,
    resolution VARCHAR(200),
    resolved_by BIGINT,
    resolved_time TIMESTAMP,
    resolution_photos TEXT,
    create_repair_order INT DEFAULT 0,
    repair_order_id BIGINT,
    remark TEXT,
    CONSTRAINT idx_tenant_defect_dn UNIQUE (tenant_id, defect_no)
);

-- 4. processes 表补充缺失列
--    Go model MesProcess 引用了 operation_code 和 is_key_process，
--    但 init.sql 建表时未包含，导致 Create/Update 时 GORM 写入被忽略
ALTER TABLE processes ADD COLUMN IF NOT EXISTS operation_code VARCHAR(50);
ALTER TABLE processes ADD COLUMN IF NOT EXISTS is_key_process BIGINT DEFAULT 0;

-- 验证
-- SELECT 'eam_inspection_plan' as tbl, count(*) as cnt FROM eam_inspection_plan
-- UNION ALL SELECT 'eam_inspection_record', count(*) FROM eam_inspection_record
-- UNION ALL SELECT 'eam_inspection_defect', count(*) FROM eam_inspection_defect
-- UNION ALL SELECT 'processes operation_code', count(*) FROM processes WHERE operation_code IS NOT NULL;
