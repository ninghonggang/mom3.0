-- Add report menus
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status)
VALUES (1, ' ShengChan BaoBiao ', 'M', '/report', 'DataLine', 15, 0, 1)
ON CONFLICT (tenant_id, path) DO NOTHING;

INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status)
VALUES
(1, ' ShengChan RiBao ', 'C', '/report/production-daily', 'DataLine', 1, (SELECT id FROM sys_menu WHERE path='/report' AND tenant_id=1), 1),
(1, ' ZhiLiang ZhouBao ', 'C', '/report/quality-weekly', 'DataAnalysis', 2, (SELECT id FROM sys_menu WHERE path='/report' AND tenant_id=1), 1),
(1, ' OEE BaoBiao ', 'C', '/report/oee', 'TrendCharts', 3, (SELECT id FROM sys_menu WHERE path='/report' AND tenant_id=1), 1),
(1, ' JiaoFu BaoBiao ', 'C', '/report/delivery', 'Truck', 4, (SELECT id FROM sys_menu WHERE path='/report' AND tenant_id=1), 1),
(1, ' AnDong BaoBiao ', 'C', '/report/andon', 'Bell', 5, (SELECT id FROM sys_menu WHERE path='/report' AND tenant_id=1), 1)
ON CONFLICT (tenant_id, path) DO NOTHING;

INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path IN (
  '/report', '/report/production-daily', '/report/quality-weekly',
  '/report/oee', '/report/delivery', '/report/andon'
) ON CONFLICT DO NOTHING;
