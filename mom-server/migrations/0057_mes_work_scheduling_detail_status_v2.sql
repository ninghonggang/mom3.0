-- ============================================================
-- 0057 / 2026-07-09 / batch 3-5
-- 给 mes_work_scheduling_detail 加 status_v2 字段 + seed 字典
-- 配套:
--   internal/pkg/status/status.go 加 MesWorkSchedulingDetailStatus 字典
--   internal/model/mes_work_scheduling_detail.go 加 StatusV2/StatusCode/SetStatus
--   internal/service/mes_work_scheduling.go StartDetail/PauseDetail/ResumeDetail/CompleteDetail 切字典码(4 处)
-- ============================================================

-- 1. 加 status_v2 字段(双轨)
ALTER TABLE mes_work_scheduling_detail ADD COLUMN IF NOT EXISTS status_v2 VARCHAR(30);
CREATE INDEX IF NOT EXISTS idx_mes_work_scheduling_detail_status_v2 ON mes_work_scheduling_detail (status_v2);

-- 2. seed mdm_status_dict
INSERT INTO mdm_status_dict (domain, entity, code, label, element_type, is_terminal, sort_order, description)
SELECT v.domain, v.entity, v.code, v.label, v.element_type, v.is_terminal, v.sort_order, v.description
FROM (VALUES
    ('mes', 'mes_work_scheduling_detail', 'PENDING',     '待开工',   'info',    FALSE, 1, '工序待开工'),
    ('mes', 'mes_work_scheduling_detail', 'IN_PROGRESS', '进行中',   'warning', FALSE, 2, '工序进行中'),
    ('mes', 'mes_work_scheduling_detail', 'PAUSED',      '已暂停',   'info',    FALSE, 3, '工序暂停'),
    ('mes', 'mes_work_scheduling_detail', 'COMPLETED',   '已完工',   'success', TRUE,  4, '工序完工')
) AS v(domain, entity, code, label, element_type, is_terminal, sort_order, description)
WHERE NOT EXISTS (
    SELECT 1 FROM mdm_status_dict d
    WHERE d.entity = v.entity AND d.code = v.code
);

-- 3. 数据回填
DO $$
BEGIN
    UPDATE mes_work_scheduling_detail
       SET status_v2 = CASE status
           WHEN 'PENDING'     THEN 'PENDING'
           WHEN 'IN_PROGRESS' THEN 'IN_PROGRESS'
           WHEN 'PAUSED'      THEN 'PAUSED'
           WHEN 'COMPLETED'   THEN 'COMPLETED'
           ELSE 'PENDING'
       END
     WHERE status_v2 IS NULL OR status_v2 = '';
    RAISE NOTICE 'mes_work_scheduling_detail status_v2 回填完成';
EXCEPTION WHEN OTHERS THEN
    RAISE NOTICE 'mes_work_scheduling_detail status_v2 回填跳过: %', SQLERRM;
END
$$;

-- 4. 回滚
/*
ALTER TABLE mes_work_scheduling_detail DROP COLUMN IF EXISTS status_v2;
DROP INDEX IF EXISTS idx_mes_work_scheduling_detail_status_v2;
DELETE FROM mdm_status_dict WHERE entity = 'mes_work_scheduling_detail';
*/