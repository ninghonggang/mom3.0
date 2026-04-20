INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES (1, 'ShengChan RiBao', 'C', '/report/production-daily', 'DataLine', 1, 46, 1);
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES (1, 'ZhiLiang ZhouBao', 'C', '/report/quality-weekly', 'DataAnalysis', 2, 46, 1);
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES (1, 'OEE BaoBiao', 'C', '/report/oee', 'TrendCharts', 3, 46, 1);
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES (1, 'JiaoFu BaoBiao', 'C', '/report/delivery', 'Truck', 4, 46, 1);
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES (1, 'AnDong BaoBiao', 'C', '/report/andon', 'Bell', 5, 46, 1);
INSERT INTO sys_role_menu (role_id, menu_id) SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path IN ('/report', '/report/production-daily', '/report/quality-weekly', '/report/oee', '/report/delivery', '/report/andon');
