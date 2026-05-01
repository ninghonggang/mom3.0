-- =====================================================
-- MOM3.0 完整菜单权限分配 (V2.0)
-- 为管理员角色(role_id=1)分配所有新菜单权限
-- =====================================================

-- System 扩展
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/system/notice';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/system/print-template';

-- Production 扩展
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/production/electronic-sop';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/production/flow-card';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/production/code-rule';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/production/issue';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/production/return';

-- Quality 扩展
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/quality/dynamic-rule';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/quality/qrci';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/quality/lpa';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/quality/inspection-plans';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/quality/aql';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/quality/lab/samples';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/quality/lab/reports';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/quality/lab-instrument';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/quality/inspection-feature';

-- MES 扩展
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/mes/material-trace';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/mes/team';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/mes/process-routes';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/mes/offline';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/mes/person-skill';

-- Alert 扩展
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/alert/notification';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/alert/rules';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/alert/records';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/alert/escalation-rules';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/alert/statistics';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/alert/notification-logs';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/alert/channels';

-- BPM 流程
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/bpm/process';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/bpm/instance';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/bpm/task';

-- APS 扩展
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/aps/rolling-config';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/aps/delivery-analysis';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/aps/material-shortage';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/aps/shortage-rule';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/aps/changeover-matrix';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/aps/product-family';

-- WMS 扩展
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/wms/transfer';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/wms/stock-check';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/wms/purchase-return';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/wms/sales-return';

-- SCP 供应链
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/scp/purchase';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/scp/rfq';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/scp/supplier-quote';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/scp/sales-order';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/scp/supplier-kpi';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/scp/customer-inquiry';

-- EAM 设备资产
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/eam/factory';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/eam/equipment-org';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/eam/downtime';

-- 父级菜单也需要分配 (如果尚未分配)
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/scp';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/alert';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/bpm';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/eam';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/mes';
