-- System extensions
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, 'Notice', 'C', '/system/notice', 'Bell', 20, (SELECT id FROM sys_menu WHERE path='/system' AND tenant_id=1), 1),
(1, 'PrintTemplate', 'C', '/system/print-template', 'Printer', 21, (SELECT id FROM sys_menu WHERE path='/system' AND tenant_id=1), 1);

-- Production extensions
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, 'ElectronicSOP', 'C', '/production/electronic-sop', 'Document', 12, (SELECT id FROM sys_menu WHERE path='/production' AND tenant_id=1), 1),
(1, 'FlowCard', 'C', '/production/flow-card', 'List', 13, (SELECT id FROM sys_menu WHERE path='/production' AND tenant_id=1), 1),
(1, 'CodeRule', 'C', '/production/code-rule', 'Key', 14, (SELECT id FROM sys_menu WHERE path='/production' AND tenant_id=1), 1);

-- Quality extension
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, 'DynamicRule', 'C', '/quality/dynamic-rule', 'Setting', 17, (SELECT id FROM sys_menu WHERE path='/quality' AND tenant_id=1), 1);

-- MES extension
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, 'MaterialTrace', 'C', '/mes/material-trace', 'Search', 24, (SELECT id FROM sys_menu WHERE path='/mes' AND tenant_id=1), 1);

-- Alert extension
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, 'AlertNotification', 'C', '/alert/notification', 'Message', 5, (SELECT id FROM sys_menu WHERE path='/alert' AND tenant_id=1), 1);
