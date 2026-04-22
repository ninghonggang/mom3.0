-- =====================================================
-- MOM3.0 新增功能菜单 - 修复版 (处理已存在父菜单)
-- =====================================================

-- 1. SCP供应链管理 - 父菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, '供应链管理', 'M', '/scp', NULL, 'Connection', 12, 0, 1
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = '/scp' AND tenant_id = 1);

-- 2. Alert统一告警中心 - 父菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, '统一告警', 'M', '/alert', NULL, 'Bell', 13, 0, 1
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = '/alert' AND tenant_id = 1);

-- 3. BPM流程引擎 - 父菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, '流程管理', 'M', '/bpm', NULL, 'Guide', 14, 0, 1
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = '/bpm' AND tenant_id = 1);

-- 插入SCP子菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, name.val, 'C', name.path, name.comp, name.icon, name.sort,
       (SELECT id FROM sys_menu WHERE path = '/scp' AND tenant_id = 1), 1
FROM (VALUES
  ('采购订单', '/scp/purchase', 'scp/PurchaseOrderList.vue', 'ShoppingCart', 1),
  ('询价单', '/scp/rfq', 'scp/RFQList.vue', 'PriceTag', 2),
  ('供应商报价', '/scp/supplier-quote', 'scp/SupplierQuoteList.vue', 'Document', 3),
  ('销售订单', '/scp/sales-order', 'scp/SCPSalesOrderList.vue', 'DocumentCopy', 4),
  ('供应商绩效', '/scp/supplier-kpi', 'scp/SupplierKPIList.vue', 'DataLine', 5),
  ('客户询价', '/scp/customer-inquiry', 'scp/CustomerInquiryList.vue', 'Message', 6)
) AS name(val, path, comp, icon, sort)
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = name.path AND tenant_id = 1);

-- 插入Alert子菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, name.val, 'C', name.path, name.comp, name.icon, name.sort,
       (SELECT id FROM sys_menu WHERE path = '/alert' AND tenant_id = 1), 1
FROM (VALUES
  ('告警规则', '/alert/rules', 'alert/AlertRulesList.vue', 'SetUp', 1),
  ('告警记录', '/alert/records', 'alert/AlertRecordsList.vue', 'List', 2),
  ('升级规则', '/alert/escalation', 'alert/AlertEscalationList.vue', 'Top', 3),
  ('告警统计', '/alert/statistics', 'alert/AlertStatistics.vue', 'DataAnalysis', 4),
  ('通知日志', '/alert/notification', 'alert/AlertNotification.vue', 'Log', 5)
) AS name(val, path, comp, icon, sort)
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = name.path AND tenant_id = 1);

-- 插入BPM子菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, name.val, 'C', name.path, name.comp, name.icon, name.sort,
       (SELECT id FROM sys_menu WHERE path = '/bpm' AND tenant_id = 1), 1
FROM (VALUES
  ('流程模型', '/bpm/process', 'bpm/ProcessList.vue', 'FlowChart', 1),
  ('流程实例', '/bpm/instance', 'bpm/InstanceList.vue', 'Connection', 2),
  ('任务实例', '/bpm/task', 'bpm/TaskList.vue', 'Tickets', 3)
) AS name(val, path, comp, icon, sort)
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = name.path AND tenant_id = 1);

-- 插入APS子菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, name.val, 'C', name.path, name.comp, name.icon, name.sort,
       (SELECT id FROM sys_menu WHERE path = '/aps' AND tenant_id = 1), 1
FROM (VALUES
  ('滚动排程', '/aps/rolling-config', 'aps/RollingConfigList.vue', 'Refresh', 4),
  ('交付分析', '/aps/delivery-analysis', 'aps/DeliveryAnalysisList.vue', 'TrendCharts', 5),
  ('缺料分析', '/aps/material-shortage', 'aps/MaterialShortageList.vue', 'Warning', 6),
  ('缺料规则', '/aps/shortage-rule', 'aps/ShortageRuleList.vue', 'SetUp', 7),
  ('换型矩阵', '/aps/changeover-matrix', 'aps/ChangeoverMatrixList.vue', 'Grid', 8),
  ('产品族', '/aps/product-family', 'aps/ProductFamilyList.vue', 'Collection', 9)
) AS name(val, path, comp, icon, sort)
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = name.path AND tenant_id = 1);

-- 插入WMS子菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, name.val, 'C', name.path, name.comp, name.icon, name.sort,
       (SELECT id FROM sys_menu WHERE path = '/wms' AND tenant_id = 1), 1
FROM (VALUES
  ('调拨管理', '/wms/transfer', 'wms/TransferOrderList.vue', 'Switch', 4),
  ('盘点管理', '/wms/stock-check', 'wms/StockCheckList.vue', 'Document', 5)
) AS name(val, path, comp, icon, sort)
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = name.path AND tenant_id = 1);

-- 插入Production子菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, name.val, 'C', name.path, name.comp, name.icon, name.sort,
       (SELECT id FROM sys_menu WHERE path = '/production' AND tenant_id = 1), 1
FROM (VALUES
  ('生产发料', '/production/issue', 'production/ProductionIssueList.vue', 'Top', 1),
  ('生产退料', '/production/return', 'production/ProductionReturnList.vue', 'Bottom', 2)
) AS name(val, path, comp, icon, sort)
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = name.path AND tenant_id = 1);

-- 插入Quality子菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, name.val, 'C', name.path, name.comp, name.icon, name.sort,
       (SELECT id FROM sys_menu WHERE path = '/quality' AND tenant_id = 1), 1
FROM (VALUES
  ('QRCI', '/quality/qrci', 'quality/QRCIList.vue', 'Edit', 1),
  ('LPA', '/quality/lpa', 'quality/LPAList.vue', 'Document', 2),
  ('检验计划', '/quality/inspection-plans', 'quality/InspectionPlanList.vue', 'SetUp', 3),
  ('AQL配置', '/quality/aql', 'quality/AQLList.vue', 'DataLine', 4),
  ('实验室样品', '/quality/lab/samples', 'quality/lab/LabSampleList.vue', 'Sample', 5),
  ('实验室报告', '/quality/lab/reports', 'quality/lab/LabReportList.vue', 'Document', 6),
  ('检测设备', '/quality/lab/instrument', 'quality/lab/LabInstrumentList.vue', 'Tools', 7)
) AS name(val, path, comp, icon, sort)
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = name.path AND tenant_id = 1);

-- 插入MES子菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, name.val, 'C', name.path, name.comp, name.icon, name.sort,
       (SELECT id FROM sys_menu WHERE path = '/mes' AND tenant_id = 1), 1
FROM (VALUES
  ('班组管理', '/mes/team', 'mes/TeamList.vue', 'User', 1),
  ('工艺路线', '/mes/process-routes', 'mes/ProcessRouteList.vue', 'Flow', 2),
  ('离线作业', '/mes/offline', 'mes/OfflineOperationList.vue', 'VideoPause', 3),
  ('人员技能', '/mes/person-skill', 'mes/PersonSkillList.vue', 'Star', 4)
) AS name(val, path, comp, icon, sort)
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = name.path AND tenant_id = 1);

-- 插入EAM子菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, name.val, 'C', name.path, name.comp, name.icon, name.sort,
       (SELECT id FROM sys_menu WHERE path = '/eam' AND tenant_id = 1), 1
FROM (VALUES
  ('工厂日历', '/eam/factory', 'eam/FactoryList.vue', 'Calendar', 1),
  ('设备组织', '/eam/equipment-org', 'eam/EquipmentOrgList.vue', 'Office', 2),
  ('停机记录', '/eam/downtime', 'eam/DowntimeList.vue', 'Switch', 3)
) AS name(val, path, comp, icon, sort)
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = name.path AND tenant_id = 1);

-- 插入Equipment子菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, name.val, 'C', name.path, name.comp, name.icon, name.sort,
       (SELECT id FROM sys_menu WHERE path = '/equipment' AND tenant_id = 1), 1
FROM (VALUES
  ('量具管理', '/equipment/gauge', 'equipment/GaugeList.vue', 'Scale', 1),
  ('检验计划', '/equipment/inspection/plans', 'equipment/inspection/InspectionPlanList.vue', 'SetUp', 2),
  ('检验记录', '/equipment/inspection/records', 'equipment/inspection/InspectionRecordList.vue', 'List', 3),
  ('检验模板', '/equipment/inspection/templates', 'equipment/inspection/InspectionTemplateList.vue', 'Document', 4),
  ('缺陷管理', '/equipment/inspection/defects', 'equipment/inspection/DefectList.vue', 'Warning', 5)
) AS name(val, path, comp, icon, sort)
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = name.path AND tenant_id = 1);

-- 插入Fin子菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, name.val, 'C', name.path, name.comp, name.icon, name.sort,
       (SELECT id FROM sys_menu WHERE path = '/fin' AND tenant_id = 1), 1
FROM (VALUES
  ('付款申请', '/fin/payment-request', 'fin/PaymentRequestList.vue', 'Money', 1),
  ('采购结算', '/fin/purchase-settlement', 'fin/PurchaseSettlementList.vue', 'ShoppingCart', 2),
  ('销售结算', '/fin/sales-settlement', 'fin/SalesSettlementList.vue', 'Sell', 3)
) AS name(val, path, comp, icon, sort)
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = name.path AND tenant_id = 1);

-- 插入Report子菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, name.val, 'C', name.path, name.comp, name.icon, name.sort,
       (SELECT id FROM sys_menu WHERE path = '/report' AND tenant_id = 1), 1
FROM (VALUES
  ('生产日报', '/report/production-daily', 'report/ProductionDailyReport.vue', 'DataLine', 1),
  ('质量周报', '/report/quality-weekly', 'report/QualityWeeklyReport.vue', 'DataAnalysis', 2),
  ('OEE报表', '/report/oee', 'report/OEEReport.vue', 'TrendCharts', 3),
  ('交付报表', '/report/delivery', 'report/DeliveryReport.vue', 'Vehicle', 4),
  ('安灯报表', '/report/andon', 'report/AndonReport.vue', 'Bell', 5)
) AS name(val, path, comp, icon, sort)
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = name.path AND tenant_id = 1);

-- 插入Integration子菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, name.val, 'C', name.path, name.comp, name.icon, name.sort,
       (SELECT id FROM sys_menu WHERE path = '/integration' AND tenant_id = 1), 1
FROM (VALUES
  ('接口配置', '/integration/interface-config', 'integration/InterfaceConfigList.vue', 'Setting', 1),
  ('执行日志', '/integration/execution-log', 'integration/ExecutionLogList.vue', 'List', 2)
) AS name(val, path, comp, icon, sort)
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = name.path AND tenant_id = 1);

-- 插入AGV子菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, name.val, 'C', name.path, name.comp, name.icon, name.sort,
       (SELECT id FROM sys_menu WHERE path = '/agv' AND tenant_id = 1), 1
FROM (VALUES
  ('AGV任务', '/agv/task', 'agv/AGVTaskList.vue', 'VideoPlay', 1),
  ('AGV设备', '/agv/device', 'agv/AGVDeviceList.vue', 'Van', 2),
  ('AGV站点', '/agv/location', 'agv/AGVLocationList.vue', 'Location', 3)
) AS name(val, path, comp, icon, sort)
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = name.path AND tenant_id = 1);

-- 插入Supplier子菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, icon, sort, parent_id, status)
SELECT 1, name.val, 'C', name.path, name.comp, name.icon, name.sort,
       (SELECT id FROM sys_menu WHERE path = '/supplier' AND tenant_id = 1), 1
FROM (VALUES
  ('ASN', '/supplier/asn', 'supplier/ASNList.vue', 'Truck', 1)
) AS name(val, path, comp, icon, sort)
WHERE NOT EXISTS (SELECT 1 FROM sys_menu WHERE path = name.path AND tenant_id = 1);

-- =====================================================
-- 角色菜单权限 (role_id=1 管理员)
-- =====================================================
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, m.id FROM sys_menu m WHERE m.tenant_id = 1
AND m.path IN (
  '/scp/purchase', '/scp/rfq', '/scp/supplier-quote', '/scp/sales-order', '/scp/supplier-kpi', '/scp/customer-inquiry',
  '/alert/rules', '/alert/records', '/alert/escalation', '/alert/statistics', '/alert/notification',
  '/bpm/process', '/bpm/instance', '/bpm/task',
  '/aps/rolling-config', '/aps/delivery-analysis', '/aps/material-shortage', '/aps/shortage-rule', '/aps/changeover-matrix', '/aps/product-family',
  '/wms/transfer', '/wms/stock-check',
  '/production/issue', '/production/return',
  '/quality/qrci', '/quality/lpa', '/quality/inspection-plans', '/quality/aql',
  '/quality/lab/samples', '/quality/lab/reports', '/quality/lab/instrument',
  '/mes/team', '/mes/process-routes', '/mes/offline', '/mes/person-skill',
  '/eam/factory', '/eam/equipment-org', '/eam/downtime',
  '/equipment/gauge',
  '/equipment/inspection/plans', '/equipment/inspection/records', '/equipment/inspection/templates', '/equipment/inspection/defects',
  '/fin/payment-request', '/fin/purchase-settlement', '/fin/sales-settlement',
  '/report/production-daily', '/report/quality-weekly', '/report/oee', '/report/delivery', '/report/andon',
  '/integration/interface-config', '/integration/execution-log',
  '/agv/task', '/agv/device', '/agv/location',
  '/supplier/asn'
)
AND NOT EXISTS (SELECT 1 FROM sys_role_menu rm WHERE rm.role_id = 1 AND rm.menu_id = m.id);
