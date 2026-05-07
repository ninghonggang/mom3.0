-- MOM3.0 菜单种子数据 v2.0
-- 按职能域重新分类，符合MOM行业最佳实践
-- Format: id, parent_id, name, type, path, component, perms, icon, sort, visible, status

-- ============================================
-- 1. 首页
-- ============================================
(1, 0, '首页', 'C', '/dashboard', 'None', '', 'HomeFilled', 1, 1, '1')

-- ============================================
-- 2. 生产执行 (Production Operations)
-- ============================================
(100, 0, '生产执行', 'M', '/production', 'None', '', 'List', 2, 1, '1')
(101, 100, '生产工单', 'C', '/production/order', 'production/ProductionOrderList.vue', 'production:order:list', 'List', 1, 1, '1')
(102, 100, '生产报工', 'C', '/production/report', 'production/ReportList.vue', 'production:report:list', 'DocumentCheck', 2, 1, '1')
(103, 100, '派工管理', 'C', '/production/dispatch', 'production/DispatchList.vue', 'production:dispatch:list', 'Tickets', 3, 1, '1')
(104, 100, '生产发料', 'C', '/production/issue', 'production/ProductionIssueList.vue', 'None', 'Top', 4, 1, '1')
(105, 100, '生产退料', 'C', '/production/return', 'production/ProductionReturnList.vue', 'None', 'Bottom', 5, 1, '1')
(107, 100, '流转卡', 'C', '/production/flow-card', 'production/FlowCardList.vue', 'None', 'List', 7, 1, '1')
(108, 100, '首末件检验', 'C', '/production/first-last-inspect', 'production/FirstLastInspectList.vue', 'production:firstlastinspect:list', 'Stamp', 8, 1, '1')
(109, 100, '电子SOP', 'C', '/production/electronic-sop', 'production/ElectronicSOPList.vue', 'None', 'Document', 9, 1, '1')
(110, 100, '工单变更', 'C', '/production/order-change', 'production/OrderChangeList.vue', 'production:orderchange:list', 'Edit', 10, 1, '1')
(111, 100, '生产看板', 'C', '/production/kanban', 'production/KanbanBoard.vue', 'production:kanban:list', 'Board', 11, 1, '1')
(112, 100, '包装条码', 'C', '/production/package', 'production/PackageList.vue', 'production:package:list', 'Printer', 12, 1, '1')
(113, 100, '编码规则', 'C', '/production/code-rule', 'production/CodeRuleList.vue', 'None', 'Key', 13, 1, '1')

-- ============================================
-- 3. 质量管理 (Quality Management)
-- ============================================
(200, 0, '质量管理', 'M', '/quality', 'None', '', 'CircleCheck', 3, 1, '1')
(201, 200, 'IQC检验', 'C', '/quality/iqc', 'quality/IQCList.vue', 'quality:iqc:list', 'CircleCheck', 1, 1, '1')
(202, 200, 'IPQC检验', 'C', '/quality/ipqc', 'quality/IPQCList.vue', 'quality:ipqc:list', 'Check', 2, 1, '1')
(203, 200, 'FQC检验', 'C', '/quality/fqc', 'quality/FQCList.vue', 'quality:fqc:list', 'CircleCheck', 3, 1, '1')
(204, 200, 'OQC检验', 'C', '/quality/oqc', 'quality/OQCList.vue', 'quality:oqc:list', 'Box', 4, 1, '1')
(205, 200, '检验计划', 'C', '/quality/inspection-plans', 'quality/InspectionPlanList.vue', 'None', 'Schedule', 5, 1, '1')
(206, 200, 'NCR处理', 'C', '/quality/ncr', 'quality/NCRList.vue', 'quality:ncr:list', 'Close', 6, 1, '1')
(207, 200, '缺陷代码', 'C', '/quality/defect-code', 'quality/DefectCodeList.vue', 'quality:defectcode:list', 'Warning', 7, 1, '1')
(208, 200, '缺陷记录', 'C', '/quality/defect-record', 'quality/DefectRecordList.vue', 'quality:defectrecord:list', 'Document', 8, 1, '1')
(209, 200, 'SPC数据', 'C', '/quality/spc', 'quality/SPCDataList.vue', 'quality:spc:list', 'DataLine', 9, 1, '1')
(210, 200, 'QRCI质量闭环', 'C', '/quality/qrci', 'quality/QRCIList.vue', 'None', 'CircleCheck', 10, 1, '1')
(211, 200, 'LPA分层审核', 'C', '/quality/lpa', 'quality/LPAStandardList.vue', 'None', 'Checked', 11, 1, '1')
(212, 200, '动态规则', 'C', '/quality/dynamic-rule', 'quality/DynamicRuleList.vue', 'None', 'Setting', 12, 1, '1')
(213, 200, 'AQL抽样标准', 'C', '/quality/aql', 'quality/AQLList.vue', 'None', 'DataLine', 13, 1, '1')
-- 实验室
(214, 200, '检测样品', 'C', '/quality/lab/samples', 'quality/lab/LabSampleList.vue', 'None', 'Box', 14, 1, '1')
(215, 200, '检测报告', 'C', '/quality/lab/reports', 'quality/lab/LabReportList.vue', 'None', 'Document', 15, 1, '1')
(216, 200, '检测仪器', 'C', '/quality/lab-instrument', 'quality/lab/LabInstrumentList.vue', 'None', 'Scale', 16, 1, '1')

-- ============================================
-- 4. 计划排程 (Planning & Scheduling)
-- ============================================
(300, 0, '计划排程', 'M', '/aps', 'None', '', 'Calendar', 4, 1, '1')
(301, 300, 'MPS计划', 'C', '/aps/mps', 'aps/MPSList.vue', 'aps:mps:list', 'Calendar', 1, 1, '1')
(302, 300, 'MRP计划', 'C', '/aps/mrp', 'aps/MRPList.vue', 'aps:mrp:list', 'Grid', 2, 1, '1')
(303, 300, 'APS排程', 'C', '/aps/schedule', 'aps/ScheduleList.vue', 'aps:schedule:list', 'List', 3, 1, '1')
(304, 300, '工作中心', 'C', '/aps/work-center', 'aps/WorkCenterList.vue', 'aps:workcenter:list', 'OfficeBuilding', 4, 1, '1')
(305, 300, '滚动排程', 'C', '/aps/rolling-config', 'aps/RollingConfigList.vue', 'None', 'Refresh', 5, 1, '1')
(306, 300, '交付分析', 'C', '/aps/delivery-analysis', 'aps/DeliveryAnalysisList.vue', 'None', 'TrendCharts', 6, 1, '1')
(307, 300, '缺料分析', 'C', '/aps/material-shortage', 'aps/MaterialShortageList.vue', 'None', 'Warning', 7, 1, '1')
(308, 300, '缺料规则', 'C', '/aps/shortage-rule', 'aps/ShortageRuleList.vue', 'None', 'SetUp', 8, 1, '1')
(309, 300, '换型矩阵', 'C', '/aps/changeover-matrix', 'aps/ChangeoverMatrixList.vue', 'None', 'Grid', 9, 1, '1')
(310, 300, '产品族', 'C', '/aps/product-family', 'aps/ProductFamilyList.vue', 'None', 'Collection', 10, 1, '1')

-- ============================================
-- 5. 设备管理 (Equipment Management)
-- ============================================
(400, 0, '设备管理', 'M', '/equipment', 'None', '', 'Monitor', 5, 1, '1')
(401, 400, '设备台账', 'C', '/equipment', 'equipment/EquipmentList.vue', 'equipment:list:list', 'Monitor', 1, 1, '1')
(402, 400, '设备点检', 'C', '/equipment/check', 'equipment/CheckList.vue', 'equipment:check:list', 'Check', 2, 1, '1')
(403, 400, '设备保养', 'C', '/equipment/maintenance', 'equipment/MaintenanceList.vue', 'equipment:maintenance:list', 'Tools', 3, 1, '1')
(404, 400, '设备维修', 'C', '/equipment/repair', 'equipment/RepairList.vue', 'equipment:repair:list', 'Tool', 4, 1, '1')
(405, 400, 'OEE分析', 'C', '/equipment/oee', 'equipment/OEELIst.vue', 'equipment:oee:list', 'DataLine', 5, 1, '1')
(406, 400, '备件管理', 'C', '/equipment/spare', 'equipment/SparePartList.vue', 'equipment:spare:list', 'Box', 6, 1, '1')
(407, 400, '点检模板', 'C', '/equipment/inspection/templates', 'equipment/InspectionTemplateList.vue', 'None', 'Tickets', 7, 1, '1')
(408, 400, '点检计划', 'C', '/equipment/inspection/plans', 'equipment/InspectionPlanList.vue', 'None', 'Schedule', 8, 1, '1')
(409, 400, '点检记录', 'C', '/equipment/inspection/records', 'equipment/InspectionRecordList.vue', 'None', 'Document', 9, 1, '1')
(410, 400, '点检缺陷', 'C', '/equipment/inspection/defects', 'equipment/InspectionDefectList.vue', 'None', 'Warning', 10, 1, '1')
(411, 400, '量检具管理', 'C', '/equipment/gauge', 'equipment/GaugeList.vue', 'None', 'Scale', 11, 1, '1')

-- ============================================
-- 6. 仓储物流 (Warehouse & Logistics)
-- ============================================
(500, 0, '仓储物流', 'M', '/wms', 'None', '', 'House', 6, 1, '1')
(501, 500, '仓库管理', 'C', '/wms/warehouse', 'wms/WarehouseList.vue', 'wms:warehouse:list', 'House', 1, 1, '1')
(502, 500, '库位管理', 'C', '/wms/location', 'wms/LocationList.vue', 'wms:location:list', 'Location', 2, 1, '1')
(503, 500, '库存管理', 'C', '/wms/inventory', 'wms/InventoryList.vue', 'wms:inventory:list', 'Box', 3, 1, '1')
(504, 500, '收货单', 'C', '/wms/receive', 'wms/ReceiveOrderList.vue', 'wms:receive:list', 'Download', 4, 1, '1')
(505, 500, '发货单', 'C', '/wms/delivery', 'wms/DeliveryOrderList.vue', 'wms:delivery:list', 'Upload', 5, 1, '1')
(506, 500, '调拨管理', 'C', '/wms/transfer', 'wms/TransferOrderList.vue', 'None', 'Switch', 6, 1, '1')
(507, 500, '盘点管理', 'C', '/wms/stock-check', 'wms/StockCheckList.vue', 'None', 'DocumentChecked', 7, 1, '1')
(508, 500, '数据采集点', 'C', '/wms/data-point', 'wms/DataPointList.vue', 'wms:datapoint:list', 'DataAnalysis', 8, 1, '1')
(509, 500, '扫描记录', 'C', '/wms/scan-log', 'wms/ScanLogList.vue', 'wms:scanlog:list', 'Scanner', 9, 1, '1')

-- ============================================
-- 7. 供应链 (Supply Chain)
-- ============================================
(600, 0, '供应链', 'M', '/scp', 'None', '', 'Connection', 7, 1, '1')
(601, 600, '采购订单', 'C', '/scp/purchase', 'scp/PurchaseOrderList.vue', 'None', 'ShoppingCart', 1, 1, '1')
(602, 600, '销售订单', 'C', '/scp/sales-order', 'scp/SCPSalesOrderList.vue', 'None', 'DocumentCopy', 2, 1, '1')
(603, 600, '询价单', 'C', '/scp/rfq', 'scp/RFQList.vue', 'None', 'PriceTag', 3, 1, '1')
(604, 600, '供应商报价', 'C', '/scp/supplier-quote', 'scp/SupplierQuoteList.vue', 'None', 'Document', 4, 1, '1')
(605, 600, '供应商绩效', 'C', '/scp/supplier-kpi', 'scp/SupplierKPIList.vue', 'None', 'DataLine', 5, 1, '1')
(606, 600, '客户询价', 'C', '/scp/customer-inquiry', 'scp/CustomerInquiryList.vue', 'None', 'Message', 6, 1, '1')

-- ============================================
-- 8. 主数据 (Master Data)
-- ============================================
(700, 0, '主数据', 'M', '/mdm', 'None', '', 'Box', 8, 1, '1')
(701, 700, '物料管理', 'C', '/mdm/material', 'mdm/MaterialList.vue', 'mdm:material:list', 'Box', 1, 1, '1')
(702, 700, '物料分类', 'C', '/mdm/material-category', 'mdm/MaterialCategoryList.vue', 'mdm:materialcategory:list', 'Folder', 2, 1, '1')
(703, 700, 'BOM管理', 'C', '/mdm/bom', 'mdm/BomList.vue', 'mdm:bom:list', 'Files', 3, 1, '1')
(704, 700, '工序管理', 'C', '/mdm/operation', 'mdm/OperationList.vue', 'mdm:operation:list', 'Operation', 4, 1, '1')
(705, 700, '供应商管理', 'C', '/mdm/supplier', 'mdm/SupplierList.vue', 'mdm:supplier:list', 'OfficeBuilding', 5, 1, '1')
(706, 700, '客户管理', 'C', '/mdm/customer', 'mdm/CustomerList.vue', 'mdm:customer:list', 'OfficeBuilding', 6, 1, '1')
(707, 700, '联系人管理', 'C', '/mdm/contact', 'mdm/ContactList.vue', 'mdm:contact:list', 'User', 7, 1, '1')
(708, 700, '银行账户', 'C', '/mdm/bank-account', 'mdm/BankAccountList.vue', 'mdm:bankaccount:list', 'CreditCard', 8, 1, '1')
(709, 700, '收货地址', 'C', '/mdm/delivery-address', 'mdm/DeliveryAddressList.vue', 'mdm:deliveryaddress:list', 'Location', 9, 1, '1')
(710, 700, '附件管理', 'C', '/mdm/attachment', 'mdm/AttachmentList.vue', 'mdm:attachment:list', 'Folder', 10, 1, '1')

-- ============================================
-- 9. 生产基础 (Production Foundation) -- MERGED INTO 生产执行 + 主数据
-- ============================================
-- (7100, 0, '生产基础', 'M', '/mes', 'mes/MesHome.vue', 'None', 'Grid', 9, 1, '1') -- MERGED INTO 生产执行
-- 车间/生产线/工位/班次/班组 → 生产执行 (100)
(7101, 100, '车间管理', 'C', '/mdm/workshop', 'mdm/WorkshopList.vue', 'mdm:workshop:list', 'OfficeBuilding', 14, 1, '1')
(7102, 100, '生产线管理', 'C', '/mdm/line', 'mdm/LineList.vue', 'mdm:line:list', 'Connection', 15, 1, '1')
(7103, 100, '工位管理', 'C', '/mdm/workstation', 'mdm/WorkstationList.vue', 'mdm:workstation:list', 'Grid', 16, 1, '1')
(7104, 100, '班次管理', 'C', '/mdm/mdm-shift', 'mdm/ShiftList.vue', 'mdm:shift:list', 'Clock', 17, 1, '1')
(7105, 100, '班组管理', 'C', '/mes/team', 'mes/TeamList.vue', 'None', 'User', 18, 1, '1')
-- 工艺路线 → 主数据 (700) - part of BOM/工艺管理
(7106, 700, '工艺路线', 'C', '/mes/process-routes', 'mes/ProcessRouteList.vue', 'None', 'Guide', 11, 1, '1')
-- 人员能力 → 生产执行
(7107, 100, '人员能力', 'C', '/mes/person-skill', 'mes/PersonSkillList.vue', 'None', 'Star', 19, 1, '1')
-- 产品离线 → 生产执行
(7108, 100, '产品离线', 'C', '/mes/offline', 'mes/OfflineList.vue', 'None', 'Warning', 20, 1, '1')
-- 物料追溯 → 追溯中心 (already exists)
-- (7109, 100, '物料追溯', 'C', '/mes/material-trace', 'mes/MaterialTrace.vue', 'None', 'Search', 9, 1, '1')

-- ============================================
-- 10. 追溯中心 (Traceability Center)
-- ============================================
(800, 0, '追溯中心', 'M', '/trace', 'None', '', 'Search', 10, 1, '1')
(801, 800, '追溯查询', 'C', '/trace/query', 'trace/TraceQuery.vue', 'trace:query:list', 'Search', 1, 1, '1')
(802, 800, '安东呼叫', 'C', '/trace/andon', 'trace/AndonCall.vue', 'trace:andon:list', 'Bell', 2, 1, '1')
(803, 800, '安东统计', 'C', '/trace/andon-stats', 'trace/AndonStats.vue', 'trace:andon:stats', 'DataLine', 3, 1, '1')
(804, 800, '升级规则', 'C', '/trace/escalation-rules', 'trace/EscalationRuleList.vue', 'trace:escalation:list', 'Guide', 4, 1, '1')

-- ============================================
-- 11. 运营分析 (Analytics & Reports)
-- ============================================
(900, 0, '运营分析', 'M', '/report', 'None', 'None', 'DataLine', 11, 1, '1')
(901, 900, '生产日报', 'C', '/report/production-daily', 'report/ProductionDailyReport.vue', 'None', 'DataLine', 1, 1, '1')
(902, 900, '质量周报', 'C', '/report/quality-weekly', 'report/QualityWeeklyReport.vue', 'None', 'DataAnalysis', 2, 1, '1')
(903, 900, 'OEE报表', 'C', '/report/oee', 'report/OEEReport.vue', 'None', 'TrendCharts', 3, 1, '1')
(904, 900, '交付报表', 'C', '/report/delivery', 'report/DeliveryReport.vue', 'None', 'Truck', 4, 1, '1')
(905, 900, '安东报表', 'C', '/report/andon', 'report/AndonReport.vue', 'None', 'Bell', 5, 1, '1')

-- ============================================
-- 12. 统一告警 (Alert Management) -- MERGED INTO 系统管理
-- ============================================
-- (1000, 0, '统一告警', 'M', '/alert', 'None', 'None', 'Bell', 12, 1, '1') -- MERGED INTO 系统管理
(1001, 2000, '告警规则', 'C', '/alert/rules', 'alert/AlertRulesList.vue', 'None', 'SetUp', 14, 1, '1')
(1002, 2000, '告警记录', 'C', '/alert/records', 'alert/AlertRecordsList.vue', 'None', 'List', 15, 1, '1')
(1003, 2000, '升级规则', 'C', '/alert/escalation', 'alert/AlertEscalationList.vue', 'None', 'Top', 16, 1, '1')
(1004, 2000, '告警统计', 'C', '/alert/statistics', 'alert/AlertStatistics.vue', 'None', 'DataAnalysis', 17, 1, '1')
(1005, 2000, '告警通知', 'C', '/alert/notification', 'alert/AlertNotification.vue', 'None', 'Message', 18, 1, '1')

-- ============================================
-- 13. 流程管理 (BPM) -- MERGED INTO 系统管理
-- ============================================
-- (1100, 0, '流程管理', 'M', '/bpm', 'None', 'None', 'Guide', 13, 1, '1') -- MERGED INTO 系统管理
(1101, 2000, '流程模型', 'C', '/bpm/process', 'bpm/ProcessList.vue', 'None', 'FlowChart', 19, 1, '1')
(1102, 2000, '流程实例', 'C', '/bpm/instance', 'bpm/InstanceList.vue', 'None', 'Connection', 20, 1, '1')
(1103, 2000, '任务实例', 'C', '/bpm/task', 'bpm/TaskList.vue', 'None', 'Tickets', 21, 1, '1')

-- ============================================
-- 14. AGV管理 -- MERGED INTO 仓储物流
-- ============================================
-- (1200, 0, 'AGV管理', 'M', '/agv', 'agv/AGVHome.vue', 'None', 'Van', 14, 1, '1') -- MERGED INTO 仓储物流
(1201, 500, 'AGV任务', 'C', '/agv/task', 'agv/TaskList.vue', 'None', 'Document', 10, 1, '1')
(1202, 500, 'AGV设备', 'C', '/agv/device', 'agv/DeviceList.vue', 'None', 'Processor', 11, 1, '1')
(1203, 500, '库位映射', 'C', '/agv/location', 'agv/LocationList.vue', 'None', 'Location', 12, 1, '1')

-- ============================================
-- 15. 能源管理 (Energy Management) -- MERGED INTO 设备管理
-- ============================================
-- (1300, 0, '能源管理', 'M', '/energy', 'None', '', 'Lightning', 15, 1, '1') -- MERGED INTO 设备管理
(1301, 400, '能源监控', 'C', '/energy/monitor', 'energy/EnergyMonitor.vue', 'energy:monitor:list', 'Lightning', 15, 1, '1')

-- ============================================
-- 16. EAM资产管理
-- ============================================
-- (1400, 0, 'EAM资产', 'M', '/eam', 'eam/EAMHome.vue', 'None', 'Office', 16, 1, '1') -- MERGED INTO 设备管理
(1401, 400, '厂区管理', 'C', '/eam/factory', 'eam/FactoryList.vue', 'None', 'OfficeBuilding', 12, 1, '1')
(1402, 400, '设备层级', 'C', '/eam/equipment-org', 'eam/EquipmentOrgList.vue', 'None', 'Grid', 13, 1, '1')
(1403, 400, '停机记录', 'C', '/eam/downtime', 'eam/DowntimeList.vue', 'None', 'Switch', 14, 1, '1')

-- ============================================
-- 17. 集成管理 (Integration) -- MERGED INTO 系统管理
-- ============================================
-- (1500, 0, '集成管理', 'M', '/integration', 'integration/IntegrationHome.vue', 'None', 'Connection', 17, 1, '1') -- MERGED INTO 系统管理
(1501, 2000, '接口配置', 'C', '/integration/interface-config', 'integration/InterfaceConfigList.vue', 'None', 'Setting', 22, 1, '1')
(1502, 2000, '执行日志', 'C', '/integration/execution-log', 'integration/ExecutionLogList.vue', 'None', 'List', 23, 1, '1')

-- ============================================
-- 18. 供应商门户 (Supplier Portal) -- MERGED INTO 供应链
-- ============================================
-- (1600, 0, '供应商门户', 'M', '/supplier', 'supplier/SupplierHome.vue', 'None', 'Truck', 18, 1, '1') -- MERGED INTO 供应链
(1601, 600, 'ASN到货通知', 'C', '/supplier/asn', 'supplier/ASNList.vue', 'None', 'Bell', 7, 1, '1')

-- ============================================
-- 19. 财务管理 (Finance) -- MERGED INTO 供应链
-- ============================================
-- (1700, 0, '财务管理', 'M', '/fin', 'fin/FinHome.vue', 'None', 'Money', 19, 1, '1') -- MERGED INTO 供应链
(1701, 600, '付款申请', 'C', '/fin/payment-request', 'fin/PaymentRequestList.vue', 'None', 'Money', 8, 1, '1')
(1702, 600, '采购结算', 'C', '/fin/purchase-settlement', 'fin/PurchaseSettlementList.vue', 'None', 'ShoppingCart', 9, 1, '1')
(1703, 600, '销售结算', 'C', '/fin/sales-settlement', 'fin/SalesSettlementList.vue', 'None', 'Sell', 10, 1, '1')

-- ============================================
-- 20. 系统管理 (System Administration)
-- ============================================
(2000, 0, '系统管理', 'M', '/system', 'None', '', 'Setting', 20, 1, '1')
(2001, 2000, '用户管理', 'C', '/system/user', 'system/UserList.vue', 'system:user:list', 'User', 1, 1, '1')
(2002, 2000, '角色管理', 'C', '/system/role', 'system/RoleList.vue', 'system:role:list', 'Key', 2, 1, '1')
(2003, 2000, '菜单管理', 'C', '/system/menu', 'system/MenuList.vue', 'system:menu:list', 'Menu', 3, 1, '1')
(2004, 2000, '部门管理', 'C', '/system/dept', 'system/DeptList.vue', 'system:dept:list', 'OfficeBuilding', 4, 1, '1')
(2005, 2000, '岗位管理', 'C', '/system/post', 'system/PostList.vue', 'system:post:list', 'Postcard', 5, 1, '1')
(2006, 2000, '租户管理', 'C', '/system/tenant', 'system/TenantList.vue', 'system:tenant:list', 'Building', 6, 1, '1')
(2007, 2000, '字典管理', 'C', '/system/dict', 'system/DictList.vue', 'system:dict:list', 'Document', 7, 1, '1')
(2008, 2000, '登录日志', 'C', '/system/login-log', 'system/LoginLogList.vue', 'system:loginlog:list', 'Key', 8, 1, '1')
(2009, 2000, '操作日志', 'C', '/system/oper-log', 'system/OperLogList.vue', 'system:operlog:list', 'Document', 9, 1, '1')
(2010, 2000, '系统配置', 'C', '/system/config', 'system/SystemConfig.vue', 'system:config:list', 'Setting', 10, 1, '1')
(2011, 2000, 'AI助手配置', 'C', '/system/ai-config', 'system/AiConfigView.vue', 'system:aiconfig:list', 'ChatDotRound', 11, 1, '1')
(2012, 2000, '通知公告', 'C', '/system/notice', 'system/NoticeList.vue', 'None', 'Bell', 12, 1, '1')
(2013, 2000, '打印模板', 'C', '/system/print-template', 'system/PrintTemplateList.vue', 'None', 'Printer', 13, 1, '1')

-- ============================================
-- 功能按钮 (Buttons) - 按模块分组
-- ============================================
-- 生产执行功能按钮
(30001, 101, '新增工单', 'F', '', 'None', 'production:order:add', '#', 1, 0, '1')
(30002, 101, '修改工单', 'F', '', 'None', 'production:order:edit', '#', 2, 0, '1')
(30003, 101, '删除工单', 'F', '', 'None', 'production:order:delete', '#', 3, 0, '1')
(30004, 101, '审核工单', 'F', '', 'None', 'production:order:approve', '#', 4, 0, '1')
(30005, 101, '下达工单', 'F', '', 'None', 'production:order:release', '#', 5, 0, '1')

-- 质量检验功能按钮
(30010, 201, '新增检验', 'F', '', 'None', 'quality:iqc:add', '#', 1, 0, '1')
(30011, 201, '修改检验', 'F', '', 'None', 'quality:iqc:edit', '#', 2, 0, '1')
(30012, 201, '删除检验', 'F', '', 'None', 'quality:iqc:delete', '#', 3, 0, '1')
(30013, 201, '审核检验', 'F', '', 'None', 'quality:iqc:approve', '#', 4, 0, '1')

-- 设备管理功能按钮
(30020, 401, '新增设备', 'F', '', 'None', 'equipment:list:add', '#', 1, 0, '1')
(30021, 401, '修改设备', 'F', '', 'None', 'equipment:list:edit', '#', 2, 0, '1')
(30022, 401, '删除设备', 'F', '', 'None', 'equipment:list:delete', '#', 3, 0, '1')
(30023, 402, '新增点检', 'F', '', 'None', 'equipment:check:add', '#', 1, 0, '1')
(30024, 402, '修改点检', 'F', '', 'None', 'equipment:check:edit', '#', 2, 0, '1')
(30025, 402, '删除点检', 'F', '', 'None', 'equipment:check:delete', '#', 3, 0, '1')
(30026, 403, '新增保养', 'F', '', 'None', 'equipment:maintenance:add', '#', 1, 0, '1')
(30027, 403, '修改保养', 'F', '', 'None', 'equipment:maintenance:edit', '#', 2, 0, '1')
(30028, 403, '删除保养', 'F', '', 'None', 'equipment:maintenance:delete', '#', 3, 0, '1')
(30029, 404, '新增维修', 'F', '', 'None', 'equipment:repair:add', '#', 1, 0, '1')
(30030, 404, '修改维修', 'F', '', 'None', 'equipment:repair:edit', '#', 2, 0, '1')
(30031, 404, '删除维修', 'F', '', 'None', 'equipment:repair:delete', '#', 3, 0, '1')
(30032, 404, '开始维修', 'F', '', 'None', 'equipment:repair:start', '#', 4, 0, '1')
(30033, 404, '完成维修', 'F', '', 'None', 'equipment:repair:complete', '#', 5, 0, '1')
(30034, 405, '导出OEE', 'F', '', 'None', 'equipment:oee:export', '#', 1, 0, '1')
(30035, 405, '新增OEE', 'F', '', 'None', 'equipment:oee:add', '#', 2, 0, '1')
(30036, 405, '查看OEE', 'F', '', 'None', 'equipment:oee:view', '#', 3, 0, '1')
(30037, 405, '删除OEE', 'F', '', 'None', 'equipment:oee:delete', '#', 4, 0, '1')

-- 仓储管理功能按钮
(30040, 501, '新增仓库', 'F', '', 'None', 'wms:warehouse:add', '#', 1, 0, '1')
(30041, 501, '修改仓库', 'F', '', 'None', 'wms:warehouse:edit', '#', 2, 0, '1')
(30042, 501, '删除仓库', 'F', '', 'None', 'wms:warehouse:delete', '#', 3, 0, '1')
(30043, 502, '新增库位', 'F', '', 'None', 'wms:location:add', '#', 1, 0, '1')
(30044, 502, '修改库位', 'F', '', 'None', 'wms:location:edit', '#', 2, 0, '1')
(30045, 502, '删除库位', 'F', '', 'None', 'wms:location:delete', '#', 3, 0, '1')
(30046, 503, '导出库存', 'F', '', 'None', 'wms:inventory:export', '#', 1, 0, '1')
(30047, 503, '盘点', 'F', '', 'None', 'wms:inventory:check', '#', 2, 0, '1')
(30048, 503, '新增库存', 'F', '', 'None', 'wms:inventory:add', '#', 3, 0, '1')
(30049, 503, '修改库存', 'F', '', 'None', 'wms:inventory:edit', '#', 4, 0, '1')
(30050, 503, '删除库存', 'F', '', 'None', 'wms:inventory:delete', '#', 5, 0, '1')

-- APS功能按钮
(30060, 301, '新增MPS', 'F', '', 'None', 'aps:mps:add', '#', 1, 0, '1')
(30061, 301, '修改MPS', 'F', '', 'None', 'aps:mps:edit', '#', 2, 0, '1')
(30062, 301, '删除MPS', 'F', '', 'None', 'aps:mps:delete', '#', 3, 0, '1')
(30063, 301, '运行MPS', 'F', '', 'None', 'aps:mps:run', '#', 4, 0, '1')
(30064, 301, '下达MPS', 'F', '', 'None', 'aps:mps:release', '#', 5, 0, '1')
(30065, 301, 'MPS明细', 'F', '', 'None', 'aps:mps:detail', '#', 6, 0, '1')
(30066, 302, '新增MRP', 'F', '', 'None', 'aps:mrp:add', '#', 1, 0, '1')
(30067, 302, '修改MRP', 'F', '', 'None', 'aps:mrp:edit', '#', 2, 0, '1')
(30068, 302, '删除MRP', 'F', '', 'None', 'aps:mrp:delete', '#', 3, 0, '1')
(30069, 302, '执行MRP', 'F', '', 'None', 'aps:mrp:execute', '#', 4, 0, '1')
(30070, 302, '缺料分析', 'F', '', 'None', 'aps:mrp:shortage', '#', 5, 0, '1')
(30071, 302, '计算结果', 'F', '', 'None', 'aps:mrp:results', '#', 6, 0, '1')
(30072, 303, '新增排程', 'F', '', 'None', 'aps:schedule:add', '#', 1, 0, '1')
(30073, 303, '修改排程', 'F', '', 'None', 'aps:schedule:edit', '#', 2, 0, '1')
(30074, 303, '删除排程', 'F', '', 'None', 'aps:schedule:delete', '#', 3, 0, '1')
(30075, 303, '执行排程', 'F', '', 'None', 'aps:schedule:execute', '#', 4, 0, '1')
(30076, 303, '甘特图', 'F', '', 'None', 'aps:schedule:gantt', '#', 5, 0, '1')
(30077, 303, '排程结果', 'F', '', 'None', 'aps:schedule:results', '#', 6, 0, '1')
(30078, 304, '新增工作中心', 'F', '', 'None', 'aps:workcenter:add', '#', 1, 0, '1')
(30079, 304, '修改工作中心', 'F', '', 'None', 'aps:workcenter:edit', '#', 2, 0, '1')
(30080, 304, '删除工作中心', 'F', '', 'None', 'aps:workcenter:delete', '#', 3, 0, '1')

-- 主数据功能按钮
(30100, 701, '新增物料', 'F', '', 'None', 'mdm:material:add', '#', 1, 0, '1')
(30101, 701, '修改物料', 'F', '', 'None', 'mdm:material:edit', '#', 2, 0, '1')
(30102, 701, '删除物料', 'F', '', 'None', 'mdm:material:delete', '#', 3, 0, '1')
(30103, 701, '导入物料', 'F', '', 'None', 'mdm:material:import', '#', 4, 0, '1')
(30104, 701, '导出物料', 'F', '', 'None', 'mdm:material:export', '#', 5, 0, '1')

-- 系统管理功能按钮
(30200, 2001, '新增用户', 'F', '', 'None', 'system:user:add', '#', 1, 0, '1')
(30201, 2001, '修改用户', 'F', '', 'None', 'system:user:edit', '#', 2, 0, '1')
(30202, 2001, '删除用户', 'F', '', 'None', 'system:user:delete', '#', 3, 0, '1')
(30203, 2001, '重置密码', 'F', '', 'None', 'system:user:resetPwd', '#', 4, 0, '1')
(30204, 2001, '导出用户', 'F', '', 'None', 'system:user:export', '#', 5, 0, '1')
(30205, 2002, '新增角色', 'F', '', 'None', 'system:role:add', '#', 1, 0, '1')
(30206, 2002, '修改角色', 'F', '', 'None', 'system:role:edit', '#', 2, 0, '1')
(30207, 2002, '删除角色', 'F', '', 'None', 'system:role:delete', '#', 3, 0, '1')
(30208, 2002, '分配权限', 'F', '', 'None', 'system:role:assign', '#', 4, 0, '1')
(30209, 2003, '新增菜单', 'F', '', 'None', 'system:menu:add', '#', 1, 0, '1')
(30210, 2003, '修改菜单', 'F', '', 'None', 'system:menu:edit', '#', 2, 0, '1')
(30211, 2003, '删除菜单', 'F', '', 'None', 'system:menu:delete', '#', 3, 0, '1')
(30212, 2004, '新增部门', 'F', '', 'None', 'system:dept:add', '#', 1, 0, '1')
(30213, 2004, '修改部门', 'F', '', 'None', 'system:dept:edit', '#', 2, 0, '1')
(30214, 2004, '删除部门', 'F', '', 'None', 'system:dept:delete', '#', 3, 0, '1')
(30215, 2005, '新增岗位', 'F', '', 'None', 'system:post:add', '#', 1, 0, '1')
(30216, 2005, '修改岗位', 'F', '', 'None', 'system:post:edit', '#', 2, 0, '1')
(30217, 2005, '删除岗位', 'F', '', 'None', 'system:post:delete', '#', 3, 0, '1')
(30218, 2006, '新增租户', 'F', '', 'None', 'system:tenant:add', '#', 1, 0, '1')
(30219, 2006, '修改租户', 'F', '', 'None', 'system:tenant:edit', '#', 2, 0, '1')
(30220, 2006, '删除租户', 'F', '', 'None', 'system:tenant:delete', '#', 3, 0, '1')
(30221, 2007, '新增字典', 'F', '', 'None', 'system:dict:add', '#', 1, 0, '1')
(30222, 2007, '修改字典', 'F', '', 'None', 'system:dict:edit', '#', 2, 0, '1')
(30223, 2007, '删除字典', 'F', '', 'None', 'system:dict:delete', '#', 3, 0, '1')
(30224, 2008, '导出日志', 'F', '', 'None', 'system:loginlog:export', '#', 1, 0, '1')
(30225, 2008, '清空日志', 'F', '', 'None', 'system:loginlog:clean', '#', 2, 0, '1')
(30226, 2009, '导出日志', 'F', '', 'None', 'system:operlog:export', '#', 1, 0, '1')
