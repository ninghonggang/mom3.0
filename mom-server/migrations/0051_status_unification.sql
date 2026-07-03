-- Migration 0051: 状态字段统一方案
-- 设计日期: 2026-07-03
-- 设计人: 架构组 / 小二
-- 关联文档: docs/MOM3.0_状态字段统一方案.md
-- 目标: 全 MOM 3.0 状态字段统一为 varchar(30) + 字典表 mdm_status_dict
-- 策略: 双轨制 — 加 status_v2 字段,不删旧 status,逐步切读

BEGIN;

-- ============================================================
-- 1. 建 mdm_status_dict 字典表
-- ============================================================
CREATE TABLE IF NOT EXISTS mdm_status_dict (
    id              BIGSERIAL PRIMARY KEY,
    tenant_id       BIGINT       NOT NULL DEFAULT 1,
    domain          VARCHAR(50)  NOT NULL,            -- 模块域: production/mes/aps/quality/wms/scp/eam
    entity          VARCHAR(50)  NOT NULL,            -- 实体: production_order/mes_process/mps/schedule_result
    code            VARCHAR(30)  NOT NULL,            -- 状态码: DRAFT/RELEASED/IN_PROGRESS/...
    label           VARCHAR(50)  NOT NULL,            -- 中文显示
    element_type    VARCHAR(20)  NOT NULL DEFAULT 'info', -- Element Plus 标签类型
    is_terminal     BOOLEAN      NOT NULL DEFAULT FALSE, -- 是否终态
    sort_order      INT          NOT NULL DEFAULT 0,  -- 排序
    description     TEXT,                              -- 描述
    legacy_int      INT,                               -- 旧 int 值(迁移期对应,NULL = 新增无对应)
    created_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,                       -- 软删除
    UNIQUE (tenant_id, entity, code)
);
CREATE INDEX IF NOT EXISTS idx_status_dict_domain_entity ON mdm_status_dict(domain, entity);
CREATE INDEX IF NOT EXISTS idx_status_dict_entity_code ON mdm_status_dict(entity, code);

COMMENT ON TABLE mdm_status_dict IS 'MOM 3.0 状态字段字典表 - 跨 module 统一状态码';
COMMENT ON COLUMN mdm_status_dict.code IS '业务状态码: DRAFT/RELEASED/IN_PROGRESS/HOLD/COMPLETED/CLOSED/CANCELLED';
COMMENT ON COLUMN mdm_status_dict.element_type IS 'Element Plus 标签类型: success/warning/danger/info/primary';
COMMENT ON COLUMN mdm_status_dict.is_terminal IS 'TRUE = 终态,不可再转移';
COMMENT ON COLUMN mdm_status_dict.legacy_int IS '旧 int 字段值,迁移期用于回查对应关系;NULL = 新增状态';

-- ============================================================
-- 2. seed 字典数据(7 个 domain × 5-7 个状态)
-- ============================================================
-- production.production_order
INSERT INTO mdm_status_dict (domain, entity, code, label, element_type, is_terminal, sort_order, legacy_int, description) VALUES
('production', 'production_order', 'DRAFT',       '草稿',     'info',    FALSE, 1, 1, '工单创建后未下达'),
('production', 'production_order', 'RELEASED',    '已下达',   'primary', FALSE, 2, 2, '工单已下达,等待开工'),
('production', 'production_order', 'IN_PROGRESS', '生产中',   'warning', FALSE, 3, 3, '工单正在生产'),
('production', 'production_order', 'HOLD',        '挂起',     'warning', FALSE, 4, NULL, '物料/设备/质量异常'),
('production', 'production_order', 'COMPLETED',   '已完成',   'success', FALSE, 5, 4, '所有工序已报工完成'),
('production', 'production_order', 'CLOSED',      '已关闭',   'info',    TRUE,  6, 5, '工单关闭(超过3天无活动)'),
('production', 'production_order', 'CANCELLED',   '已取消',   'info',    TRUE,  7, NULL, '工单取消');

-- production.mobile_job_report
INSERT INTO mdm_status_dict (domain, entity, code, label, element_type, is_terminal, sort_order, legacy_int, description) VALUES
('production', 'mobile_job_report', 'SUBMITTED', '已提交',  'primary', FALSE, 1, 1, '工人扫码报工已提交'),
('production', 'mobile_job_report', 'CONFIRMED', '已确认',  'success', FALSE, 2, 2, '班组长已确认'),
('production', 'mobile_job_report', 'AUDITED',   '已审核',  'success', TRUE,  3, 3, '质量/车间已审核'),
('production', 'mobile_job_report', 'REJECTED',  '已驳回',  'danger',  FALSE, 4, NULL, '报工被驳回');

-- mes.mes_process
INSERT INTO mdm_status_dict (domain, entity, code, label, element_type, is_terminal, sort_order, legacy_int, description) VALUES
('mes', 'mes_process', 'DRAFT',     '草稿',     'info',    FALSE, 1, NULL, '工艺路线草稿'),
('mes', 'mes_process', 'ACTIVE',    '生效',     'success', FALSE, 2, NULL, '工艺路线当前生效版本'),
('mes', 'mes_process', 'OBSOLETE',  '失效',     'info',    TRUE,  3, NULL, '工艺路线被新版本取代');

-- aps.mps
INSERT INTO mdm_status_dict (domain, entity, code, label, element_type, is_terminal, sort_order, legacy_int, description) VALUES
('aps', 'mps', 'DRAFT',       '草稿',     'info',    FALSE, 1, 1, 'MPS 计划草稿'),
('aps', 'mps', 'RELEASED',    '已下达',   'primary', FALSE, 2, 2, 'MPS 已下达'),
('aps', 'mps', 'IN_PROGRESS', '执行中',   'warning', FALSE, 3, 3, 'MPS 执行中'),
('aps', 'mps', 'COMPLETED',   '已完成',   'success', TRUE,  4, NULL, 'MPS 全部完成'),
('aps', 'mps', 'CANCELLED',   '已取消',   'info',    TRUE,  5, NULL, 'MPS 取消');

-- aps.mrp
INSERT INTO mdm_status_dict (domain, entity, code, label, element_type, is_terminal, sort_order, legacy_int, description) VALUES
('aps', 'mrp', 'PENDING',   '待计算',   'info',    FALSE, 1, 1, 'MRP 待计算'),
('aps', 'mrp', 'RUNNING',   '计算中',   'primary', FALSE, 2, 2, 'MRP 计算中'),
('aps', 'mrp', 'COMPLETED', '已完成',   'success', TRUE,  3, 3, 'MRP 计算完成'),
('aps', 'mrp', 'FAILED',    '失败',     'danger',  TRUE,  4, NULL, 'MRP 计算失败');

-- aps.schedule_plan
INSERT INTO mdm_status_dict (domain, entity, code, label, element_type, is_terminal, sort_order, legacy_int, description) VALUES
('aps', 'schedule_plan', 'PENDING',   '待排程',   'info',    FALSE, 1, 1, '排程任务待执行'),
('aps', 'schedule_plan', 'RUNNING',   '排程中',   'primary', FALSE, 2, 2, '排程算法执行中'),
('aps', 'schedule_plan', 'COMPLETED', '已完成',   'success', TRUE,  3, 3, '排程完成'),
('aps', 'schedule_plan', 'FAILED',    '失败',     'danger',  TRUE,  4, NULL, '排程失败');

-- aps.schedule_result
INSERT INTO mdm_status_dict (domain, entity, code, label, element_type, is_terminal, sort_order, legacy_int, description) VALUES
('aps', 'schedule_result', 'PENDING',     '待执行',   'info',    FALSE, 1, 1, '排程结果待执行'),
('aps', 'schedule_result', 'IN_PROGRESS', '执行中',   'warning', FALSE, 2, 2, '排程结果执行中'),
('aps', 'schedule_result', 'COMPLETED',   '已完成',   'success', TRUE,  3, 3, '排程结果完成'),
('aps', 'schedule_result', 'CANCELLED',   '已取消',   'info',    TRUE,  4, NULL, '排程结果取消');

-- aps.work_center
INSERT INTO mdm_status_dict (domain, entity, code, label, element_type, is_terminal, sort_order, legacy_int, description) VALUES
('aps', 'work_center', 'ACTIVE',    '生效',   'success', FALSE, 1, NULL, '工作中心启用'),
('aps', 'work_center', 'INACTIVE',  '停用',   'info',    TRUE,  2, NULL, '工作中心停用'),
('aps', 'work_center', 'MAINTENANCE','维护中', 'warning', FALSE, 3, NULL, '工作中心维护中');

-- ============================================================
-- 3. 加 status_v2 字段(双轨,不删旧 status)
-- ============================================================
ALTER TABLE production_orders     ADD COLUMN IF NOT EXISTS status_v2 VARCHAR(30);
ALTER TABLE production_dispatch   ADD COLUMN IF NOT EXISTS status_v2 VARCHAR(30);
ALTER TABLE production_report     ADD COLUMN IF NOT EXISTS status_v2 VARCHAR(30);
ALTER TABLE mobile_job_report     ADD COLUMN IF NOT EXISTS status_v2 VARCHAR(30);
ALTER TABLE mes_process           ADD COLUMN IF NOT EXISTS status_v2 VARCHAR(30);
ALTER TABLE aps_mps               ADD COLUMN IF NOT EXISTS status_v2 VARCHAR(30);
ALTER TABLE aps_mrp               ADD COLUMN IF NOT EXISTS status_v2 VARCHAR(30);
ALTER TABLE aps_schedule_plan     ADD COLUMN IF NOT EXISTS status_v2 VARCHAR(30);
ALTER TABLE aps_schedule_result   ADD COLUMN IF NOT EXISTS status_v2 VARCHAR(30);
ALTER TABLE aps_work_center       ADD COLUMN IF NOT EXISTS status_v2 VARCHAR(30);

-- ============================================================
-- 4. 数据迁移:status(int/varchar) → status_v2(varchar)
-- ============================================================
-- production_orders
UPDATE production_orders SET status_v2 = 'DRAFT'       WHERE status = 1 AND status_v2 IS NULL;
UPDATE production_orders SET status_v2 = 'IN_PROGRESS' WHERE status = 2 AND status_v2 IS NULL;
UPDATE production_orders SET status_v2 = 'COMPLETED'   WHERE status = 3 AND status_v2 IS NULL;
UPDATE production_orders SET status_v2 = 'CANCELLED'   WHERE status = 4 AND status_v2 IS NULL;
UPDATE production_orders SET status_v2 = 'CLOSED'      WHERE status = 5 AND status_v2 IS NULL;

-- mobile_job_report
UPDATE mobile_job_report SET status_v2 = 'SUBMITTED' WHERE status = 1 AND status_v2 IS NULL;
UPDATE mobile_job_report SET status_v2 = 'CONFIRMED' WHERE status = 2 AND status_v2 IS NULL;
UPDATE mobile_job_report SET status_v2 = 'AUDITED'   WHERE status = 3 AND status_v2 IS NULL;

-- aps_mps
UPDATE aps_mps SET status_v2 = 'DRAFT'       WHERE status = 1 AND status_v2 IS NULL;
UPDATE aps_mps SET status_v2 = 'IN_PROGRESS' WHERE status = 2 AND status_v2 IS NULL;
UPDATE aps_mps SET status_v2 = 'COMPLETED'   WHERE status = 3 AND status_v2 IS NULL;

-- aps_mrp
UPDATE aps_mrp SET status_v2 = 'PENDING'   WHERE status = 1 AND status_v2 IS NULL;
UPDATE aps_mrp SET status_v2 = 'RUNNING'   WHERE status = 2 AND status_v2 IS NULL;
UPDATE aps_mrp SET status_v2 = 'COMPLETED' WHERE status = 3 AND status_v2 IS NULL;

-- aps_schedule_plan
UPDATE aps_schedule_plan SET status_v2 = 'PENDING'   WHERE status = 1 AND status_v2 IS NULL;
UPDATE aps_schedule_plan SET status_v2 = 'RUNNING'   WHERE status = 2 AND status_v2 IS NULL;
UPDATE aps_schedule_plan SET status_v2 = 'COMPLETED' WHERE status = 3 AND status_v2 IS NULL;

-- aps_schedule_result
UPDATE aps_schedule_result SET status_v2 = 'PENDING'      WHERE status = 1 AND status_v2 IS NULL;
UPDATE aps_schedule_result SET status_v2 = 'IN_PROGRESS'  WHERE status = 2 AND status_v2 IS NULL;
UPDATE aps_schedule_result SET status_v2 = 'COMPLETED'    WHERE status = 3 AND status_v2 IS NULL;

-- mes_process(原本就是 varchar,直接 copy)
UPDATE mes_process SET status_v2 = status WHERE status_v2 IS NULL;

-- aps_work_center(原本就是 varchar,直接 copy)
UPDATE aps_work_center SET status_v2 = status WHERE status_v2 IS NULL;

-- ============================================================
-- 5. 校验(若 status_v2 还有 NULL 报警)
-- ============================================================
DO $$
DECLARE
    rec RECORD;
    null_count INT;
BEGIN
    FOR rec IN
        SELECT 'production_orders'   AS tbl UNION ALL
        SELECT 'production_dispatch' UNION ALL
        SELECT 'production_report'   UNION ALL
        SELECT 'mobile_job_report'   UNION ALL
        SELECT 'mes_process'         UNION ALL
        SELECT 'aps_mps'             UNION ALL
        SELECT 'aps_mrp'             UNION ALL
        SELECT 'aps_schedule_plan'   UNION ALL
        SELECT 'aps_schedule_result' UNION ALL
        SELECT 'aps_work_center'
    LOOP
        EXECUTE format('SELECT COUNT(*) FROM %I WHERE status_v2 IS NULL', rec.tbl) INTO null_count;
        IF null_count > 0 THEN
            RAISE WARNING 'table %: % rows with NULL status_v2 (review needed)', rec.tbl, null_count;
        ELSE
            RAISE NOTICE 'table %: status_v2 fully migrated ✓', rec.tbl;
        END IF;
    END LOOP;
END $$;

COMMIT;

-- ============================================================
-- 回滚脚本(0051_status_unification.down.sql)
-- ============================================================
-- BEGIN;
-- ALTER TABLE production_orders     DROP COLUMN IF EXISTS status_v2;
-- ALTER TABLE production_dispatch   DROP COLUMN IF EXISTS status_v2;
-- ALTER TABLE production_report     DROP COLUMN IF EXISTS status_v2;
-- ALTER TABLE mobile_job_report     DROP COLUMN IF EXISTS status_v2;
-- ALTER TABLE mes_process           DROP COLUMN IF EXISTS status_v2;
-- ALTER TABLE aps_mps               DROP COLUMN IF EXISTS status_v2;
-- ALTER TABLE aps_mrp               DROP COLUMN IF EXISTS status_v2;
-- ALTER TABLE aps_schedule_plan     DROP COLUMN IF EXISTS status_v2;
-- ALTER TABLE aps_schedule_result   DROP COLUMN IF EXISTS status_v2;
-- ALTER TABLE aps_work_center       DROP COLUMN IF EXISTS status_v2;
-- -- mdm_status_dict 字典表本身保留(无副作用)
-- COMMIT;
