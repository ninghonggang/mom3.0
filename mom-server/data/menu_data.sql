-- =====================================================
-- 新增功能菜单 (MOM3.0 V2.0)
-- 数据库: mom3_db
-- =====================================================

-- 1. SCP供应链管理 - 父菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '供应链管理', 'M', '/scp', 'Connection', 12, 0, 1);

-- SCP子菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '采购订单', 'C', '/scp/purchase', 'ShoppingCart', 1, (SELECT id FROM sys_menu WHERE path='/scp' AND tenant_id=1), 1),
(1, '询价单', 'C', '/scp/rfq', 'PriceTag', 2, (SELECT id FROM sys_menu WHERE path='/scp' AND tenant_id=1), 1),
(1, '供应商报价', 'C', '/scp/supplier-quote', 'Document', 3, (SELECT id FROM sys_menu WHERE path='/scp' AND tenant_id=1), 1),
(1, '销售订单', 'C', '/scp/sales-order', 'DocumentCopy', 4, (SELECT id FROM sys_menu WHERE path='/scp' AND tenant_id=1), 1),
(1, '供应商绩效', 'C', '/scp/supplier-kpi', 'DataLine', 5, (SELECT id FROM sys_menu WHERE path='/scp' AND tenant_id=1), 1),
(1, '客户询价', 'C', '/scp/customer-inquiry', 'Message', 6, (SELECT id FROM sys_menu WHERE path='/scp' AND tenant_id=1), 1);

-- 2. Alert统一告警中心 - 父菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '统一告警', 'M', '/alert', 'Bell', 13, 0, 1);

INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '告警规则', 'C', '/alert/rules', 'SetUp', 1, (SELECT id FROM sys_menu WHERE path='/alert' AND tenant_id=1), 1),
(1, '告警记录', 'C', '/alert/records', 'List', 2, (SELECT id FROM sys_menu WHERE path='/alert' AND tenant_id=1), 1),
(1, '升级规则', 'C', '/alert/escalation-rules', 'Top', 3, (SELECT id FROM sys_menu WHERE path='/alert' AND tenant_id=1), 1),
(1, '告警统计', 'C', '/alert/statistics', 'DataAnalysis', 4, (SELECT id FROM sys_menu WHERE path='/alert' AND tenant_id=1), 1),
(1, '通知日志', 'C', '/alert/notification-logs', 'Log', 5, (SELECT id FROM sys_menu WHERE path='/alert' AND tenant_id=1), 1),
(1, '通知渠道', 'C', '/alert/channels', 'Connection', 6, (SELECT id FROM sys_menu WHERE path='/alert' AND tenant_id=1), 1);

-- 3. BPM流程引擎 - 父菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '流程管理', 'M', '/bpm', 'Guide', 14, 0, 1);

INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '流程模型', 'C', '/bpm/process', 'FlowChart', 1, (SELECT id FROM sys_menu WHERE path='/bpm' AND tenant_id=1), 1),
(1, '流程实例', 'C', '/bpm/instance', 'Connection', 2, (SELECT id FROM sys_menu WHERE path='/bpm' AND tenant_id=1), 1),
(1, '任务实例', 'C', '/bpm/task', 'Tickets', 3, (SELECT id FROM sys_menu WHERE path='/bpm' AND tenant_id=1), 1);

-- 4. APS扩展菜单 - 添加到现有APS父菜单下
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '滚动排程', 'C', '/aps/rolling-config', 'Refresh', 4, (SELECT id FROM sys_menu WHERE path='/aps' AND tenant_id=1), 1),
(1, '交付分析', 'C', '/aps/delivery-analysis', 'TrendCharts', 5, (SELECT id FROM sys_menu WHERE path='/aps' AND tenant_id=1), 1),
(1, '缺料分析', 'C', '/aps/material-shortage', 'Warning', 6, (SELECT id FROM sys_menu WHERE path='/aps' AND tenant_id=1), 1),
(1, '缺料规则', 'C', '/aps/shortage-rule', 'SetUp', 7, (SELECT id FROM sys_menu WHERE path='/aps' AND tenant_id=1), 1),
(1, '换型矩阵', 'C', '/aps/changeover-matrix', 'Grid', 8, (SELECT id FROM sys_menu WHERE path='/aps' AND tenant_id=1), 1),
(1, '产品族', 'C', '/aps/product-family', 'Collection', 9, (SELECT id FROM sys_menu WHERE path='/aps' AND tenant_id=1), 1);

-- 5. WMS扩展菜单 - 添加到现有WMS父菜单下
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '调拨管理', 'C', '/wms/transfer', 'Switch', 4, (SELECT id FROM sys_menu WHERE path='/wms' AND tenant_id=1), 1),
(1, '盘点管理', 'C', '/wms/stock-check', 'DocumentChecked', 5, (SELECT id FROM sys_menu WHERE path='/wms' AND tenant_id=1), 1),
(1, '采购退货', 'C', '/wms/purchase-return', 'Return', 6, (SELECT id FROM sys_menu WHERE path='/wms' AND tenant_id=1), 1),
(1, '销售退货', 'C', '/wms/sales-return', 'Return', 7, (SELECT id FROM sys_menu WHERE path='/wms' AND tenant_id=1), 1);

-- 5.1 生产发料管理 - 添加到现有生产执行父菜单下
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '生产发料', 'C', '/production/issue', 'Box', 10, (SELECT id FROM sys_menu WHERE path='/production' AND tenant_id=1), 1);

-- 5.2 生产退料管理 - 添加到现有生产执行父菜单下
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '生产退料', 'C', '/production/return', 'Return', 11, (SELECT id FROM sys_menu WHERE path='/production' AND tenant_id=1), 1);

-- 6. Quality扩展菜单 - 添加到现有Quality父菜单下
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, 'QRCI质量闭环', 'C', '/quality/qrci', 'CircleCheck', 10, (SELECT id FROM sys_menu WHERE path='/quality' AND tenant_id=1), 1),
(1, 'LPA分层审核', 'C', '/quality/lpa', 'Checked', 11, (SELECT id FROM sys_menu WHERE path='/quality' AND tenant_id=1), 1),
(1, '检验计划', 'C', '/quality/inspection-plans', 'Document', 12, (SELECT id FROM sys_menu WHERE path='/quality' AND tenant_id=1), 1),
(1, 'AQL抽样方案', 'C', '/quality/aql', 'DataAnalysis', 13, (SELECT id FROM sys_menu WHERE path='/quality' AND tenant_id=1), 1),
(1, '检测样品', 'C', '/quality/lab/samples', 'TestTube', 14, (SELECT id FROM sys_menu WHERE path='/quality' AND tenant_id=1), 1),
(1, '检验特性', 'C', '/quality/inspection-feature', 'Grid', 16, (SELECT id FROM sys_menu WHERE path='/quality' AND tenant_id=1), 1),
(1, '检测报告', 'C', '/quality/lab/reports', 'Document', 15, (SELECT id FROM sys_menu WHERE path='/quality' AND tenant_id=1), 1),
(1, '仪器管理', 'C', '/quality/lab-instrument', 'Tools', 16, (SELECT id FROM sys_menu WHERE path='/quality' AND tenant_id=1), 1);

-- 7. MES班组管理 - 添加到现有MES父菜单下
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '班组管理', 'C', '/mes/team', 'Team', 20, (SELECT id FROM sys_menu WHERE path='/mes' AND tenant_id=1), 1);

-- 7.1 MES工艺路线管理 - 添加到现有MES父菜单下
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '工艺路线', 'C', '/mes/process-routes', 'Guide', 21, (SELECT id FROM sys_menu WHERE path='/mes' AND tenant_id=1), 1);

-- 7.2 MES产品离线管理 - 添加到现有MES父菜单下
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '产品离线', 'C', '/mes/offline', 'Warning', 22, (SELECT id FROM sys_menu WHERE path='/mes' AND tenant_id=1), 1);

-- 7.3 MES人员能力矩阵 - 添加到现有MES父菜单下
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '人员能力', 'C', '/mes/person-skill', 'User', 23, (SELECT id FROM sys_menu WHERE path='/mes' AND tenant_id=1), 1);

-- 7.4 EAM设备组织层级 - 添加到现有MDM父菜单下
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '设备组织', 'M', '/eam', 'OfficeBuilding', 8, 0, 1);

INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '厂区管理', 'C', '/eam/factory', 'HomeFilled', 1, (SELECT id FROM sys_menu WHERE path='/eam' AND tenant_id=1), 1),
(1, '设备组织层级', 'C', '/eam/equipment-org', 'Connection', 2, (SELECT id FROM sys_menu WHERE path='/eam' AND tenant_id=1), 1);

INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '停机记录', 'C', '/eam/downtime', 'Clock', 13, (SELECT id FROM sys_menu WHERE path='/eam' AND tenant_id=1), 1);


-- 8. 设备点检子菜单 - 添加到现有设备父菜单下
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '点检模板', 'C', '/equipment/inspection/templates', 'Tickets', 6, (SELECT id FROM sys_menu WHERE path='/equipment' AND tenant_id=1), 1),
(1, '点检计划', 'C', '/equipment/inspection/plans', 'Schedule', 7, (SELECT id FROM sys_menu WHERE path='/equipment' AND tenant_id=1), 1),
(1, '点检记录', 'C', '/equipment/inspection/records', 'Document', 8, (SELECT id FROM sys_menu WHERE path='/equipment' AND tenant_id=1), 1),
(1, '点检缺陷', 'C', '/equipment/inspection/defects', 'Warning', 9, (SELECT id FROM sys_menu WHERE path='/equipment' AND tenant_id=1), 1),
(1, '量检具管理', 'C', '/equipment/gauge', 'Scale', 10, (SELECT id FROM sys_menu WHERE path='/equipment' AND tenant_id=1), 1),
(1, '器具管理', 'C', '/equipment/tools', 'Tool', 11, (SELECT id FROM sys_menu WHERE path='/equipment' AND tenant_id=1), 1),
(1, '容器管理', 'C', '/equipment/containers', 'Box', 12, (SELECT id FROM sys_menu WHERE path='/equipment' AND tenant_id=1), 1),
(1, '器具容器绑定', 'C', '/equipment/tool-bindings', 'Connection', 13, (SELECT id FROM sys_menu WHERE path='/equipment' AND tenant_id=1), 1),
(1, '容器生命周期', 'C', '/equipment/container-lifecycle', 'Refresh', 14, (SELECT id FROM sys_menu WHERE path='/equipment' AND tenant_id=1), 1);

-- 9. MDM客户供应商扩展菜单 - 添加到现有MDM父菜单下
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '联系人管理', 'C', '/mdm/contacts', 'User', 30, (SELECT id FROM sys_menu WHERE path='/mdm' AND tenant_id=1), 1),
(1, '银行账户', 'C', '/mdm/bank-accounts', 'CreditCard', 31, (SELECT id FROM sys_menu WHERE path='/mdm' AND tenant_id=1), 1),
(1, '附件管理', 'C', '/mdm/attachments', 'Folder', 32, (SELECT id FROM sys_menu WHERE path='/mdm' AND tenant_id=1), 1);

-- AI视觉检测菜单 - 添加到现有AI父菜单下（假设AI菜单path为/ai）
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '视觉检测', 'C', '/ai/visual-inspection', 'Camera', 10, (SELECT id FROM sys_menu WHERE path='/ai' AND tenant_id=1), 1);

-- 给管理员角色分配新菜单
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path LIKE '/scp%';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path LIKE '/alert%';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path LIKE '/bpm%';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path IN ('/aps/rolling-config','/aps/delivery-analysis','/aps/material-shortage','/aps/shortage-rule','/aps/changeover-matrix','/aps/product-family');
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path IN ('/wms/transfer','/wms/stock-check');
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/production/issue';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/production/return';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path IN ('/quality/qrci','/quality/lpa','/quality/inspection-plans','/quality/aql','/quality/lab/samples','/quality/lab/reports','/quality/lab-instrument');
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path LIKE '/equipment/inspection%';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/equipment/gauge';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/mes/team';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/mes/process-routes';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/mes/offline';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/mes/person-skill';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path LIKE '/eam%';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path IN ('/quality/inspection-plans','/quality/aql');
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path IN ('/equipment/tools','/equipment/containers','/equipment/tool-bindings','/equipment/container-lifecycle');
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path IN ('/mdm/contacts','/mdm/bank-accounts','/mdm/attachments');
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/ai/visual-inspection';

-- 8. 系统集成 - 接口配置管理（作为一级菜单）
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '系统集成', 'M', '/integration', 'Connection', 12, 0, 1);

-- 8.1 接口配置管理
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '接口配置', 'C', '/integration/interface-config', 'Interface', 1, (SELECT id FROM sys_menu WHERE path='/integration' AND tenant_id=1), 1);

-- 8.2 执行日志查询
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '执行日志', 'C', '/integration/execution-log', 'List', 2, (SELECT id FROM sys_menu WHERE path='/integration' AND tenant_id=1), 1);

-- 分配系统集成菜单权限
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/integration';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/integration/interface-config';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/integration/execution-log';

-- 9. AGV调度
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, 'AGV调度', 'M', '/agv', 'Apple', 13, 0, 1);

-- 9.1 AGV任务管理
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, 'AGV任务', 'C', '/agv/task', 'Document', 1, (SELECT id FROM sys_menu WHERE path='/agv' AND tenant_id=1), 1);

-- 9.2 AGV设备管理
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, 'AGV设备', 'C', '/agv/device', 'Processor', 2, (SELECT id FROM sys_menu WHERE path='/agv' AND tenant_id=1), 1);

-- 9.3 库位映射管理
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '库位映射', 'C', '/agv/location', 'Location', 3, (SELECT id FROM sys_menu WHERE path='/agv' AND tenant_id=1), 1);

-- 分配AGV菜单权限
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/agv';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/agv/task';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/agv/device';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/agv/location';

-- 10. 供应商ASN
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '供应商ASN', 'M', '/supplier-asn', 'Truck', 14, 0, 1);

-- 10.1 ASN到货通知
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, 'ASN到货', 'C', '/supplier-asn/asn', 'Document', 1, (SELECT id FROM sys_menu WHERE path='/supplier-asn' AND tenant_id=1), 1);

-- 分配供应商ASN菜单权限
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/supplier-asn';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/supplier-asn/asn';

-- 11. 报表模块
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '生产报表', 'M', '/report', 'DataLine', 15, 0, 1);

INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '生产日报', 'C', '/report/production-daily', 'DataLine', 1, (SELECT id FROM sys_menu WHERE path='/report' AND tenant_id=1), 1),
(1, '质量周报', 'C', '/report/quality-weekly', 'DataAnalysis', 2, (SELECT id FROM sys_menu WHERE path='/report' AND tenant_id=1), 1),
(1, 'OEE报表', 'C', '/report/oee', 'TrendCharts', 3, (SELECT id FROM sys_menu WHERE path='/report' AND tenant_id=1), 1),
(1, '交付报表', 'C', '/report/delivery', 'Truck', 4, (SELECT id FROM sys_menu WHERE path='/report' AND tenant_id=1), 1),
(1, '安东报表', 'C', '/report/andon', 'Bell', 5, (SELECT id FROM sys_menu WHERE path='/report' AND tenant_id=1), 1);

-- 分配报表菜单权限
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/report';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/report/production-daily';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/report/quality-weekly';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/report/oee';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/report/delivery';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/report/andon';

-- 12. 财务模块
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '财务管理', 'M', '/fin', 'Money', 16, 0, 1);

INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '采购结算', 'C', '/fin/purchase-settlement', 'Document', 1, (SELECT id FROM sys_menu WHERE path='/fin' AND tenant_id=1), 1),
(1, '销售结算', 'C', '/fin/sales-settlement', 'DocumentCopy', 2, (SELECT id FROM sys_menu WHERE path='/fin' AND tenant_id=1), 1),
(1, '付款申请', 'C', '/fin/payment-request', 'Money', 3, (SELECT id FROM sys_menu WHERE path='/fin' AND tenant_id=1), 1);

-- 分配财务菜单权限
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/fin';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/fin/purchase-settlement';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/fin/sales-settlement';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/fin/payment-request';

-- 13. System扩展菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '通知公告', 'C', '/system/notice', 'Bell', 20, (SELECT id FROM sys_menu WHERE path='/system' AND tenant_id=1), 1),
(1, '打印模板', 'C', '/system/print-template', 'Printer', 21, (SELECT id FROM sys_menu WHERE path='/system' AND tenant_id=1), 1);

-- 分配System扩展菜单权限
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/system/notice';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/system/print-template';

-- 14. Production扩展菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '电子SOP', 'C', '/production/electronic-sop', 'Document', 12, (SELECT id FROM sys_menu WHERE path='/production' AND tenant_id=1), 1),
(1, '流转卡', 'C', '/production/flow-card', 'List', 13, (SELECT id FROM sys_menu WHERE path='/production' AND tenant_id=1), 1),
(1, '编码规则', 'C', '/production/code-rule', 'Key', 14, (SELECT id FROM sys_menu WHERE path='/production' AND tenant_id=1), 1);

-- 分配Production扩展菜单权限
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/production/electronic-sop';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/production/flow-card';
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/production/code-rule';

-- 15. Quality扩展菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '动态规则', 'C', '/quality/dynamic-rule', 'Setting', 17, (SELECT id FROM sys_menu WHERE path='/quality' AND tenant_id=1), 1);

-- 分配Quality扩展菜单权限
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/quality/dynamic-rule';

-- 16. MES扩展菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '物料追溯', 'C', '/mes/material-trace', 'Search', 24, (SELECT id FROM sys_menu WHERE path='/mes' AND tenant_id=1), 1);

-- 分配MES扩展菜单权限
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/mes/material-trace';

-- 17. Alert扩展菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, parent_id, status) VALUES
(1, '告警通知', 'C', '/alert/notification', 'Message', 5, (SELECT id FROM sys_menu WHERE path='/alert' AND tenant_id=1), 1);

-- 分配Alert扩展菜单权限
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu WHERE tenant_id = 1 AND path = '/alert/notification';

-- 验证
SELECT COUNT(*) as total_menus FROM sys_menu WHERE tenant_id = 1;
