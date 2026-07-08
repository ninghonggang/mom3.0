-- Migration 0052: mdm.bom(boms + bom_items)建表
-- 设计日期: 2026-07-08
-- 设计人: 架构组 / 小二
-- 关联文档:
--   - docs/MOM3.0_状态字段统一方案.md(状态字段双轨)
--   - migrations/0051_status_unification.sql(已加 mdm.bom seed 字典)
-- 触发原因:
--   - model/internal/model/bom.go 写 TableName()="boms" / "bom_items"
--   - 但 DB 实际没有这两张表(P0 bug,model 层先写,migration 后建)
--   - 0051 跑时 DO 块跳过 boms 段,留下 blocker
-- 策略:
--   - 用 IF NOT EXISTS 兼容重复跑
--   - 双轨制:status varchar(20) + status_v2 varchar(30) 同模式 0051
--   - 与 0051 mdm.bom 字典 seed 配套:DRAFT(草稿)/ACTIVE(生效)/OBSOLETE(失效)

-- ============================================================
-- 1. 建 boms 表(MOM 3.0 mdm.bom 实体,BOM 头)
-- ============================================================
CREATE TABLE IF NOT EXISTS boms (
    id              BIGSERIAL    PRIMARY KEY,
    tenant_id       BIGINT       NOT NULL                DEFAULT 1,
    bom_code        VARCHAR(50)  NOT NULL,
    bom_name        VARCHAR(200) NOT NULL,
    material_id     BIGINT       NOT NULL                DEFAULT 0,
    material_code   VARCHAR(50),
    material_name   VARCHAR(100),
    version         VARCHAR(20),
    status          VARCHAR(20)  NOT NULL                DEFAULT 'DRAFT',  -- legacy varchar,与 status_v2 双轨
    status_v2       VARCHAR(30),                                         -- 双轨:V2.1 与 mdm_status_dict 对齐
    eff_date        DATE,
    exp_date        DATE,
    remark          VARCHAR(500),
    erp_bom_code    VARCHAR(50),
    erp_sync_time   TIMESTAMPTZ,
    erp_sync_status VARCHAR(20),
    is_current      INT          NOT NULL                DEFAULT 1,
    created_at      TIMESTAMPTZ  NOT NULL                DEFAULT NOW(),
    updated_at      TIMESTAMPTZ  NOT NULL                DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,
    CONSTRAINT uniq_boms_tenant_code UNIQUE (tenant_id, bom_code)
);

CREATE INDEX IF NOT EXISTS idx_boms_tenant           ON boms(tenant_id);
CREATE INDEX IF NOT EXISTS idx_boms_material_id      ON boms(material_id);
CREATE INDEX IF NOT EXISTS idx_boms_status_v2        ON boms(status_v2);
CREATE INDEX IF NOT EXISTS idx_boms_deleted_at       ON boms(deleted_at);

COMMENT ON TABLE  boms              IS 'MOM 3.0 MDM BOM 物料清单头表';
COMMENT ON COLUMN boms.status       IS 'legacy varchar 状态:DRAFT/ACTIVE/EXPIRED(双轨保留)';
COMMENT ON COLUMN boms.status_v2    IS 'V2.1 双轨:与 mdm_status_dict DRAFT/ACTIVE/OBSOLETE 对齐';
COMMENT ON COLUMN boms.is_current   IS '是否当前版本:0 否 / 1 是';
COMMENT ON COLUMN boms.erp_sync_status IS '金蝶同步状态:SYNCED/PENDING/FAILED';

-- ============================================================
-- 2. 建 bom_items 表(MOM 3.0 mdm.bom 实体,BOM 行)
-- ============================================================
CREATE TABLE IF NOT EXISTS bom_items (
    id                 BIGSERIAL    PRIMARY KEY,
    tenant_id          BIGINT       NOT NULL                DEFAULT 1,
    bom_id             BIGINT       NOT NULL,
    line_no            INT          NOT NULL                DEFAULT 0,
    material_id        BIGINT       NOT NULL,
    material_code      VARCHAR(50),
    material_name      VARCHAR(100),
    quantity           DECIMAL(18,4) NOT NULL               DEFAULT 0,
    unit               VARCHAR(20),
    scrap_rate         DECIMAL(10,4) NOT NULL               DEFAULT 0,
    substitute_group   VARCHAR(50),
    is_alternative     INT          NOT NULL                DEFAULT 0,
    created_at         TIMESTAMPTZ  NOT NULL                DEFAULT NOW(),
    updated_at         TIMESTAMPTZ  NOT NULL                DEFAULT NOW(),
    deleted_at         TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_bom_items_tenant      ON bom_items(tenant_id);
CREATE INDEX IF NOT EXISTS idx_bom_items_bom_id      ON bom_items(bom_id);
CREATE INDEX IF NOT EXISTS idx_bom_items_material_id ON bom_items(material_id);
CREATE INDEX IF NOT EXISTS idx_bom_items_deleted_at  ON bom_items(deleted_at);

COMMENT ON TABLE  bom_items                  IS 'MOM 3.0 MDM BOM 物料清单行表';
COMMENT ON COLUMN bom_items.scrap_rate       IS '损耗率%,DECIMAL(10,4)';
COMMENT ON COLUMN bom_items.is_alternative   IS '是否替代料:0 否 / 1 是';

-- ============================================================
-- 3. 数据迁移(boms 表刚建,没历史数据,seed 占位即可)
-- ============================================================
UPDATE boms SET status_v2 = status WHERE status_v2 IS NULL;

-- ============================================================
-- 4. 校验
-- ============================================================
DO $$
DECLARE
    null_count INT;
    bom_count  INT;
    item_count INT;
BEGIN
    SELECT COUNT(*) INTO bom_count  FROM boms;
    SELECT COUNT(*) INTO item_count FROM bom_items;
    RAISE NOTICE 'boms 行数: %, bom_items 行数: %', bom_count, item_count;

    SELECT COUNT(*) INTO null_count FROM boms WHERE status_v2 IS NULL;
    IF null_count > 0 THEN
        RAISE WARNING 'boms: % rows with NULL status_v2', null_count;
    ELSE
        RAISE NOTICE 'boms: status_v2 fully migrated ✓';
    END IF;
END $$;

-- ============================================================
-- 回滚脚本(0052_mdm_bom_table.down.sql)
-- ============================================================
-- DROP TABLE IF EXISTS bom_items;
-- DROP TABLE IF EXISTS boms;
-- -- mdm_status_dict 字典表里 mdm.bom 的 3 条 seed 保留(无副作用)
