-- ============================================================
-- 0056 / 2026-07-09 / batch 3-3
-- 给 sys_interface_config 加 status_v2 字段 + seed 字典
-- 配套:
--   internal/pkg/status/status.go 加 IntegrationConfigStatus 字典
--   internal/model/integration.go 加 InterfaceConfig.StatusV2/StatusCode/SetStatus
--   internal/service/integration_executor.go 切字典码(1 处 == ENABLE)
-- ============================================================

-- 1. 加 status_v2 字段(双轨)
ALTER TABLE sys_interface_config ADD COLUMN IF NOT EXISTS status_v2 VARCHAR(30);
CREATE INDEX IF NOT EXISTS idx_sys_interface_config_status_v2 ON sys_interface_config (status_v2);

-- 2. seed mdm_status_dict
INSERT INTO mdm_status_dict (domain, entity, code, label, element_type, is_terminal, sort_order, description)
SELECT v.domain, v.entity, v.code, v.label, v.element_type, v.is_terminal, v.sort_order, v.description
FROM (VALUES
    ('integration', 'sys_interface_config', 'ENABLE',  '启用',  'success', TRUE, 1, '接口配置启用'),
    ('integration', 'sys_interface_config', 'DISABLE', '禁用',  'info',   TRUE, 2, '接口配置禁用')
) AS v(domain, entity, code, label, element_type, is_terminal, sort_order, description)
WHERE NOT EXISTS (
    SELECT 1 FROM mdm_status_dict d
    WHERE d.entity = v.entity AND d.code = v.code
);

-- 3. 数据回填
DO $$
BEGIN
    UPDATE sys_interface_config
       SET status_v2 = CASE status
           WHEN 'ENABLE'  THEN 'ENABLE'
           WHEN 'DISABLE' THEN 'DISABLE'
           ELSE 'ENABLE'
       END
     WHERE status_v2 IS NULL OR status_v2 = '';
    RAISE NOTICE 'sys_interface_config status_v2 回填完成';
EXCEPTION WHEN OTHERS THEN
    RAISE NOTICE 'sys_interface_config status_v2 回填跳过: %', SQLERRM;
END
$$;

-- 4. 回滚
/*
ALTER TABLE sys_interface_config DROP COLUMN IF EXISTS status_v2;
DROP INDEX IF EXISTS idx_sys_interface_config_status_v2;
DELETE FROM mdm_status_dict WHERE entity = 'sys_interface_config';
*/