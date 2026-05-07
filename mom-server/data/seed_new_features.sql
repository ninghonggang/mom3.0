-- 新增功能菜单配置
-- 执行前请确保数据库已初始化

-- ============================================
-- 1. VMI管理 (仓储物流下)
-- ============================================
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, perms, icon, sort, parent_id, status)
SELECT 1, 'VMI管理', 'C', '/wms/vmi', 'wms/VmiList.vue', 'wms:vmi:list', 'Box', 20, id, 1
FROM sys_menu WHERE path = '/wms' AND tenant_id = 1
ON CONFLICT DO NOTHING;

-- 获取VMI菜单ID
DO $$
DECLARE
    vmi_menu_id BIGINT;
BEGIN
    SELECT id INTO vmi_menu_id FROM sys_menu WHERE path = '/wms/vmi' AND tenant_id = 1;
    IF vmi_menu_id IS NOT NULL THEN
        -- VMI供应商
        INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, perms, icon, sort, parent_id, status)
        VALUES (1, 'VMI供应商', 'C', '/wms/vmi/vendor', 'None', 'wms:vmi:vendor:list', 'OfficeBuilding', 1, vmi_menu_id, 1)
        ON CONFLICT DO NOTHING;

        -- VMI物料
        INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, perms, icon, sort, parent_id, status)
        VALUES (1, 'VMI物料', 'C', '/wms/vmi/material', 'None', 'wms:vmi:material:list', 'Box', 2, vmi_menu_id, 1)
        ON CONFLICT DO NOTHING;

        -- VMI事务
        INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, perms, icon, sort, parent_id, status)
        VALUES (1, 'VMI事务', 'C', '/wms/vmi/transaction', 'None', 'wms:vmi:transaction:list', 'List', 3, vmi_menu_id, 1)
        ON CONFLICT DO NOTHING;
    END IF;
END $$;

-- ============================================
-- 2. 质量证书 (质量管理下)
-- ============================================
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, perms, icon, sort, parent_id, status)
SELECT 1, '质量证书', 'C', '/quality/certificate', 'quality/QualityCertificateList.vue', 'quality:certificate:list', 'Document', 20, id, 1
FROM sys_menu WHERE path = '/quality' AND tenant_id = 1
ON CONFLICT DO NOTHING;

-- ============================================
-- 3. 客户信用 (供应链下)
-- ============================================
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, perms, icon, sort, parent_id, status)
SELECT 1, '客户信用', 'C', '/scp/customer-credit', 'scp/CustomerCreditList.vue', 'scp:customer-credit:list', 'CreditCard', 20, id, 1
FROM sys_menu WHERE path = '/scp' AND tenant_id = 1
ON CONFLICT DO NOTHING;

-- ============================================
-- 4. IDOC管理 (集成管理下)
-- ============================================
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, perms, icon, sort, parent_id, status)
SELECT 1, 'IDOC管理', 'C', '/integration/idoc', 'integration/IdocList.vue', 'integration:idoc:list', 'Connection', 30, id, 1
FROM sys_menu WHERE path = '/integration' AND tenant_id = 1
ON CONFLICT DO NOTHING;

-- ============================================
-- 5. 移动报工 (MES作业下)
-- ============================================
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, perms, icon, sort, parent_id, status)
SELECT 1, '移动报工', 'C', '/mes/mobile-job-report', 'mes/MobileJobReportList.vue', 'mes:mobile-job-report:list', 'Cellphone', 30, id, 1
FROM sys_menu WHERE path = '/mes' AND tenant_id = 1
ON CONFLICT DO NOTHING;

-- ============================================
-- 分配菜单权限给管理员角色
-- ============================================
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path IN (
    '/wms/vmi', '/wms/vmi/vendor', '/wms/vmi/material', '/wms/vmi/transaction',
    '/quality/certificate',
    '/scp/customer-credit',
    '/integration/idoc',
    '/mes/mobile-job-report'
) ON CONFLICT DO NOTHING;
