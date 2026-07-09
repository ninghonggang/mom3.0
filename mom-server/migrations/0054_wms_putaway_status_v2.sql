-- ============================================================
-- 0054 / 2026-07-09 / batch 3-2
-- 给 wms_putaway_job + wms_putaway_record 加 status_v2 字段 + seed 字典
-- 配套:
--   internal/pkg/status/status.go 加 WMSPutawayJob + WMSPutawayRecord 字典
--   internal/model/wms_putaway.go 加 StatusV2 + StatusCode()/SetStatus()
--   internal/service/wms_putaway.go 切字典码(4 处 == + 1 处赋值)
-- ============================================================

-- 1. 先建表(如果不存在 — model 声明了但 DB 未建,batch 3-2 踩过这个坑)
CREATE TABLE IF NOT EXISTS wms_putaway_job (
    id              BIGSERIAL PRIMARY KEY,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP DEFAULT NULL,
    tenant_id       BIGINT NOT NULL,
    putaway_no      VARCHAR(50) NOT NULL,
    source_type     VARCHAR(20),
    source_no       VARCHAR(50),
    warehouse_id    BIGINT,
    status          VARCHAR(20) DEFAULT 'PENDING',
    status_v2       VARCHAR(30),
    assign_time     TIMESTAMP,
    operator_id     BIGINT,
    operator_name   VARCHAR(50),
    putaway_time    TIMESTAMP,
    remark          TEXT
);
CREATE INDEX IF NOT EXISTS idx_wms_putaway_job_tenant_id   ON wms_putaway_job (tenant_id);
CREATE INDEX IF NOT EXISTS idx_wms_putaway_job_warehouse_id ON wms_putaway_job (warehouse_id);

CREATE TABLE IF NOT EXISTS wms_putaway_record (
    id              BIGSERIAL PRIMARY KEY,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP DEFAULT NULL,
    tenant_id       BIGINT NOT NULL,
    putaway_job_id  BIGINT NOT NULL,
    putaway_no      VARCHAR(50) NOT NULL,
    item_id         BIGINT,
    item_code       VARCHAR(50),
    item_name       VARCHAR(100),
    from_location_id BIGINT,
    to_location_id  BIGINT,
    putaway_qty     DECIMAL(18,3),
    putawarded_qty  DECIMAL(18,3) DEFAULT 0,
    status          VARCHAR(20) DEFAULT 'PENDING',
    status_v2       VARCHAR(30)
);
CREATE INDEX IF NOT EXISTS idx_wms_putaway_record_tenant_id    ON wms_putaway_record (tenant_id);
CREATE INDEX IF NOT EXISTS idx_wms_putaway_record_putaway_job_id ON wms_putaway_record (putaway_job_id);

-- 2. 加 status_v2 字段(双轨,IF EXISTS 防重复)
ALTER TABLE wms_putaway_job     ADD COLUMN IF NOT EXISTS status_v2 VARCHAR(30);
ALTER TABLE wms_putaway_record  ADD COLUMN IF NOT EXISTS status_v2 VARCHAR(30);

CREATE INDEX IF NOT EXISTS idx_wms_putaway_job_status_v2     ON wms_putaway_job (status_v2);
CREATE INDEX IF NOT EXISTS idx_wms_putaway_record_status_v2  ON wms_putaway_record (status_v2);

-- 2. seed mdm_status_dict(wms.wms_putaway_job)
INSERT INTO mdm_status_dict (domain, entity, code, label, element_type, is_terminal, sort_order, description)
SELECT v.domain, v.entity, v.code, v.label, v.element_type, v.is_terminal, v.sort_order, v.description
FROM (VALUES
    ('wms', 'wms_putaway_job', 'PENDING',    '待分配',     'info',    FALSE, 1, '上架作业待分配操作员'),
    ('wms', 'wms_putaway_job', 'ASSIGNED',   '已分配',     'warning', FALSE, 2, '上架作业已分配'),
    ('wms', 'wms_putaway_job', 'PUTAWAYING', '上架中',     'warning', FALSE, 3, '上架进行中'),
    ('wms', 'wms_putaway_job', 'COMPLETED',  '已完成',     'success', TRUE,  4, '上架完成'),
    ('wms', 'wms_putaway_job', 'CANCELLED',  '已取消',     'info',    TRUE,  5, '上架取消'),
    ('wms', 'wms_putaway_record', 'PENDING',    '待上架',     'info',    FALSE, 1, '上架明细待处理'),
    ('wms', 'wms_putaway_record', 'PUTAWAYING', '上架中',     'warning', FALSE, 2, '上架明细处理中'),
    ('wms', 'wms_putaway_record', 'COMPLETED',  '已完成',     'success', TRUE,  3, '上架明细完成')
) AS v(domain, entity, code, label, element_type, is_terminal, sort_order, description)
WHERE NOT EXISTS (
    SELECT 1 FROM mdm_status_dict d
    WHERE d.entity = v.entity AND d.code = v.code
);

-- 3. 数据回填(空表或 0 行时 NOTICE 跳过)
DO $$
BEGIN
    UPDATE wms_putaway_job
       SET status_v2 = CASE status
           WHEN 'PENDING'    THEN 'PENDING'
           WHEN 'ASSIGNED'   THEN 'ASSIGNED'
           WHEN 'PUTAWAYING' THEN 'PUTAWAYING'
           WHEN 'COMPLETED'  THEN 'COMPLETED'
           WHEN 'CANCELLED'  THEN 'CANCELLED'
           ELSE 'PENDING'
       END
     WHERE status_v2 IS NULL OR status_v2 = '';
    RAISE NOTICE 'wms_putaway_job status_v2 回填完成';

    UPDATE wms_putaway_record
       SET status_v2 = CASE status
           WHEN 'PENDING'    THEN 'PENDING'
           WHEN 'PUTAWAYING' THEN 'PUTAWAYING'
           WHEN 'COMPLETED'  THEN 'COMPLETED'
           ELSE 'PENDING'
       END
     WHERE status_v2 IS NULL OR status_v2 = '';
    RAISE NOTICE 'wms_putaway_record status_v2 回填完成';
EXCEPTION WHEN OTHERS THEN
    RAISE NOTICE 'wms_putaway status_v2 回填跳过: %', SQLERRM;
END
$$;

-- 4. 回滚(用 DO 块逐表容错)
/*
ALTER TABLE wms_putaway_job     DROP COLUMN IF EXISTS status_v2;
ALTER TABLE wms_putaway_record  DROP COLUMN IF EXISTS status_v2;
DROP INDEX IF EXISTS idx_wms_putaway_job_status_v2;
DROP INDEX IF EXISTS idx_wms_putaway_record_status_v2;
DELETE FROM mdm_status_dict WHERE entity IN ('wms_putaway_job', 'wms_putaway_record');
*/