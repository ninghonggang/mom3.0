-- MDM BOM 物料清单头表
CREATE TABLE IF NOT EXISTS mdm_bom (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    bom_code VARCHAR(50) NOT NULL,
    bom_name VARCHAR(200) NOT NULL,
    material_id BIGINT NOT NULL,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    version VARCHAR(20) DEFAULT 'V1',
    status VARCHAR(20) DEFAULT 'DRAFT',
    eff_date DATE,
    exp_date DATE,
    remark VARCHAR(500),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_bom_code ON mdm_bom(tenant_id, bom_code) WHERE deleted_at IS NULL;

-- MDM BOM 物料清单行表
CREATE TABLE IF NOT EXISTS mdm_bom_item (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    bom_id BIGINT NOT NULL,
    line_no INT DEFAULT 0,
    material_id BIGINT NOT NULL,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    quantity DECIMAL(18,4) DEFAULT 0,
    unit VARCHAR(20),
    scrap_rate DECIMAL(10,4) DEFAULT 0,
    substitute_group VARCHAR(50),
    is_alternative INT DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_mdm_bom_item_bom_id ON mdm_bom_item(bom_id) WHERE deleted_at IS NULL;

-- MDM 工序表
CREATE TABLE IF NOT EXISTS mdm_operation (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    operation_code VARCHAR(50) NOT NULL,
    operation_name VARCHAR(100) NOT NULL,
    workcenter_id BIGINT,
    workcenter_name VARCHAR(100),
    standard_worktime INT DEFAULT 0,
    quality_std VARCHAR(500),
    is_key_process INT DEFAULT 0,
    is_qc_point INT DEFAULT 0,
    sequence INT DEFAULT 0,
    remark VARCHAR(500),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_op_code ON mdm_operation(tenant_id, operation_code) WHERE deleted_at IS NULL;

-- MDM 班次表
CREATE TABLE IF NOT EXISTS mdm_shift (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    shift_code VARCHAR(50) NOT NULL,
    shift_name VARCHAR(100) NOT NULL,
    start_time VARCHAR(10),
    end_time VARCHAR(10),
    work_hours DECIMAL(10,2) DEFAULT 8,
    is_night INT DEFAULT 0,
    remark VARCHAR(500),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_mdm_shift_code ON mdm_shift(tenant_id, shift_code) WHERE deleted_at IS NULL;
