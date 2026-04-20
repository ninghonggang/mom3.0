-- Assign role 1 permissions for new menus
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/system/notice';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/system/print-template';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/production/electronic-sop';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/production/flow-card';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/production/code-rule';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/quality/dynamic-rule';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/mes/material-trace';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/alert/notification';
