import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue'),
    meta: { title: '登录', public: true }
  },
  {
    path: '/',
    component: () => import('@/components/layout/MainLayout.vue'),
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/Dashboard.vue'),
        meta: { title: '首页' }
      },
      // 系统管理
      {
        path: 'system/user',
        name: 'UserList',
        component: () => import('@/views/system/UserList.vue'),
        meta: { title: '用户管理', icon: 'User' }
      },
      {
        path: 'system/role',
        name: 'RoleList',
        component: () => import('@/views/system/RoleList.vue'),
        meta: { title: '角色管理', icon: 'Key' }
      },
      {
        path: 'system/menu',
        name: 'MenuList',
        component: () => import('@/views/system/MenuList.vue'),
        meta: { title: '菜单管理', icon: 'Menu' }
      },
      {
        path: 'system/dept',
        name: 'DeptList',
        component: () => import('@/views/system/DeptList.vue'),
        meta: { title: '部门管理', icon: 'OfficeBuilding' }
      },
      {
        path: 'system/dict',
        name: 'DictList',
        component: () => import('@/views/system/DictList.vue'),
        meta: { title: '字典管理', icon: 'Document' }
      },
      {
        path: 'system/post',
        name: 'PostList',
        component: () => import('@/views/system/PostList.vue'),
        meta: { title: '岗位管理', icon: 'Postcard' }
      },
      {
        path: 'system/tenant',
        name: 'TenantList',
        component: () => import('@/views/system/TenantList.vue'),
        meta: { title: '租户管理', icon: 'Building' }
      },
      {
        path: 'system/login-log',
        name: 'LoginLogList',
        component: () => import('@/views/system/LoginLogList.vue'),
        meta: { title: '登录日志', icon: 'Key' }
      },
      {
        path: 'system/oper-log',
        name: 'OperLogList',
        component: () => import('@/views/system/OperLogList.vue'),
        meta: { title: '操作日志', icon: 'Document' }
      },
      {
        path: 'system/config',
        name: 'SystemConfig',
        component: () => import('@/views/system/SystemConfig.vue'),
        meta: { title: '系统配置', icon: 'Setting' }
      },
      {
        path: 'system/ai-config',
        name: 'AiConfig',
        component: () => import('@/views/system/AiConfigView.vue'),
        meta: { title: 'AI助手配置', icon: 'ChatDotRound' }
      },
      {
        path: 'system/notice',
        name: 'NoticeList',
        component: () => import('@/views/system/NoticeList.vue'),
        meta: { title: '通知公告', icon: 'Bell' }
      },
      {
        path: 'system/print-template',
        name: 'PrintTemplateList',
        component: () => import('@/views/system/PrintTemplateList.vue'),
        meta: { title: '打印模板', icon: 'Printer' }
      },
      // 主数据管理
      {
        path: 'mdm/material',
        name: 'MaterialList',
        component: () => import('@/views/mdm/MaterialList.vue'),
        meta: { title: '物料管理', icon: 'Box' }
      },
      {
        path: 'mdm/material-category',
        name: 'MaterialCategoryList',
        component: () => import('@/views/mdm/MaterialCategoryList.vue'),
        meta: { title: '物料分类', icon: 'Folder' }
      },
      {
        path: 'mdm/workshop',
        name: 'WorkshopList',
        component: () => import('@/views/mdm/WorkshopList.vue'),
        meta: { title: '车间管理', icon: 'OfficeBuilding' }
      },
      {
        path: 'mdm/line',
        name: 'ProductionLineList',
        component: () => import('@/views/mdm/LineList.vue'),
        meta: { title: '生产线管理', icon: 'Connection' }
      },
      {
        path: 'mdm/workstation',
        name: 'WorkstationList',
        component: () => import('@/views/mdm/WorkstationList.vue'),
        meta: { title: '工位管理', icon: 'Grid' }
      },
      {
        path: 'mdm/mdm-shift',
        name: 'MdmShiftList',
        component: () => import('@/views/mdm/ShiftList.vue'),
        meta: { title: '班次管理', icon: 'Clock' }
      },
      {
        path: 'mdm/bom',
        name: 'BomList',
        component: () => import('@/views/mdm/BomList.vue'),
        meta: { title: 'BOM管理', icon: 'Files' }
      },
      {
        path: 'mdm/operation',
        name: 'OperationList',
        component: () => import('@/views/mdm/OperationList.vue'),
        meta: { title: '工序管理', icon: 'Operation' }
      },
      {
        path: 'mdm/customer',
        name: 'CustomerList',
        component: () => import('@/views/mdm/CustomerList.vue'),
        meta: { title: '客户管理', icon: 'User' }
      },
      {
        path: 'mdm/supplier',
        name: 'SupplierList',
        component: () => import('@/views/mdm/SupplierList.vue'),
        meta: { title: '供应商管理', icon: 'OfficeBuilding' }
      },
      {
        path: 'mdm/contact',
        name: 'ContactList',
        component: () => import('@/views/mdm/ContactList.vue'),
        meta: { title: '联系人管理', icon: 'UserFilled' }
      },
      {
        path: 'mdm/bank-account',
        name: 'BankAccountList',
        component: () => import('@/views/mdm/BankAccountList.vue'),
        meta: { title: '银行账户', icon: 'CreditCard' }
      },
      {
        path: 'mdm/attachment',
        name: 'AttachmentList',
        component: () => import('@/views/mdm/AttachmentList.vue'),
        meta: { title: '附件管理', icon: 'Folder' }
      },
      {
        path: 'mdm/delivery-address',
        name: 'DeliveryAddressList',
        component: () => import('@/views/mdm/DeliveryAddressList.vue'),
        meta: { title: '收货地址', icon: 'Location' }
      },
      // 生产执行
      {
        path: 'production/sales-order',
        name: 'SalesOrderList',
        component: () => import('@/views/production/SalesOrderList.vue'),
        meta: { title: '销售订单', icon: 'Document' }
      },
      {
        path: 'production/report',
        name: 'ReportList',
        component: () => import('@/views/production/ReportList.vue'),
        meta: { title: '生产报工', icon: 'DocumentCheck' }
      },
      {
        path: 'production/dispatch',
        name: 'DispatchList',
        component: () => import('@/views/production/DispatchList.vue'),
        meta: { title: '派工管理', icon: 'Tickets' }
      },
      {
        path: 'production/order',
        name: 'ProductionOrderList',
        component: () => import('@/views/production/ProductionOrderList.vue'),
        meta: { title: '生产工单', icon: 'List' }
      },
      {
        path: 'production/kanban',
        name: 'KanbanBoard',
        component: () => import('@/views/production/KanbanBoard.vue'),
        meta: { title: '生产看板', icon: 'DataBoard' }
      },
      {
        path: 'production/order-change',
        name: 'OrderChangeList',
        component: () => import('@/views/production/OrderChangeList.vue'),
        meta: { title: '工单变更', icon: 'Edit' }
      },
      {
        path: 'production/package',
        name: 'PackageList',
        component: () => import('@/views/production/PackageList.vue'),
        meta: { title: '包装条码', icon: 'Box' }
      },
      {
        path: 'production/first-last-inspect',
        name: 'FirstLastInspectList',
        component: () => import('@/views/production/FirstLastInspectList.vue'),
        meta: { title: '首末件检验', icon: 'Check' }
      },
      {
        path: 'production/electronic-sop',
        name: 'ElectronicSOPList',
        component: () => import('@/views/production/ElectronicSOPList.vue'),
        meta: { title: '电子SOP', icon: 'Document' }
      },
      {
        path: 'production/flow-card',
        name: 'FlowCardList',
        component: () => import('@/views/production/FlowCardList.vue'),
        meta: { title: '流转卡', icon: 'List' }
      },
      {
        path: 'production/code-rule',
        name: 'CodeRuleList',
        component: () => import('@/views/production/CodeRuleList.vue'),
        meta: { title: '编码规则', icon: 'Key' }
      },
      // 设备管理
      {
        path: 'equipment',
        name: 'EquipmentList',
        component: () => import('@/views/equipment/EquipmentList.vue'),
        meta: { title: '设备台账', icon: 'Monitor' }
      },
      {
        path: 'equipment/check',
        name: 'EquipmentCheck',
        component: () => import('@/views/equipment/CheckList.vue'),
        meta: { title: '设备点检', icon: 'Check' }
      },
      {
        path: 'equipment/maintenance',
        name: 'EquipmentMaintenance',
        component: () => import('@/views/equipment/MaintenanceList.vue'),
        meta: { title: '设备保养', icon: 'Tools' }
      },
      {
        path: 'equipment/repair',
        name: 'EquipmentRepair',
        component: () => import('@/views/equipment/RepairList.vue'),
        meta: { title: '设备维修', icon: 'Tool' }
      },
      {
        path: 'equipment/spare',
        name: 'SparePartList',
        component: () => import('@/views/equipment/SparePartList.vue'),
        meta: { title: '备件管理', icon: 'Box' }
      },
      {
        path: 'equipment/oee',
        name: 'OEELIst',
        component: () => import('@/views/equipment/OEELIst.vue'),
        meta: { title: 'OEE分析', icon: 'DataLine' }
      },
      // 仓储管理
      {
        path: 'wms/warehouse',
        name: 'WarehouseList',
        component: () => import('@/views/wms/WarehouseList.vue'),
        meta: { title: '仓库管理', icon: 'House' }
      },
      {
        path: 'wms/data-point',
        name: 'DataPointList',
        component: () => import('@/views/wms/DataPointList.vue'),
        meta: { title: '数据采集点', icon: 'DataLine' }
      },
      {
        path: 'wms/scan-log',
        name: 'ScanLogList',
        component: () => import('@/views/wms/ScanLogList.vue'),
        meta: { title: '扫描记录', icon: 'Scanner' }
      },
      {
        path: 'wms/location',
        name: 'LocationList',
        component: () => import('@/views/wms/LocationList.vue'),
        meta: { title: '库位管理', icon: 'Location' }
      },
      {
        path: 'wms/inventory',
        name: 'InventoryList',
        component: () => import('@/views/wms/InventoryList.vue'),
        meta: { title: '库存管理', icon: 'Box' }
      },
      {
        path: 'wms/receive',
        name: 'ReceiveOrderList',
        component: () => import('@/views/wms/ReceiveOrderList.vue'),
        meta: { title: '收货单', icon: 'Download' }
      },
      {
        path: 'wms/delivery',
        name: 'DeliveryOrderList',
        component: () => import('@/views/wms/DeliveryOrderList.vue'),
        meta: { title: '发货单', icon: 'Upload' }
      },
      // 质量管理
      {
        path: 'quality/iqc',
        name: 'IQCList',
        component: () => import('@/views/quality/IQCList.vue'),
        meta: { title: 'IQC检验', icon: 'CircleCheck' }
      },
      {
        path: 'quality/ipqc',
        name: 'IPQCList',
        component: () => import('@/views/quality/IPQCList.vue'),
        meta: { title: 'IPQC检验', icon: 'Check' }
      },
      {
        path: 'quality/fqc',
        name: 'FQCList',
        component: () => import('@/views/quality/FQCList.vue'),
        meta: { title: 'FQC检验', icon: 'CircleCheck' }
      },
      {
        path: 'quality/oqc',
        name: 'OQCList',
        component: () => import('@/views/quality/OQCList.vue'),
        meta: { title: 'OQC检验', icon: 'Box' }
      },
      {
        path: 'quality/defect-code',
        name: 'DefectCodeList',
        component: () => import('@/views/quality/DefectCodeList.vue'),
        meta: { title: '不良代码', icon: 'Warning' }
      },
      {
        path: 'quality/defect-record',
        name: 'DefectRecordList',
        component: () => import('@/views/quality/DefectRecordList.vue'),
        meta: { title: '不良记录', icon: 'Document' }
      },
      {
        path: 'quality/ncr',
        name: 'NCRList',
        component: () => import('@/views/quality/NCRList.vue'),
        meta: { title: 'NCR', icon: 'FolderOpened' }
      },
      {
        path: 'quality/spc',
        name: 'SPCDataList',
        component: () => import('@/views/quality/SPCDataList.vue'),
        meta: { title: 'SPC数据', icon: 'DataLine' }
      },
      // APS计划
      {
        path: 'aps/mps',
        name: 'MPSList',
        component: () => import('@/views/aps/MPSList.vue'),
        meta: { title: 'MPS计划', icon: 'Calendar' }
      },
      {
        path: 'aps/mrp',
        name: 'MRPList',
        component: () => import('@/views/aps/MRPList.vue'),
        meta: { title: 'MRP计划', icon: 'Grid' }
      },
      {
        path: 'aps/schedule',
        name: 'ScheduleList',
        component: () => import('@/views/aps/ScheduleList.vue'),
        meta: { title: '排程计划', icon: 'List' }
      },
      {
        path: 'aps/work-center',
        name: 'WorkCenterList',
        component: () => import('@/views/aps/WorkCenterList.vue'),
        meta: { title: '工作中心', icon: 'OfficeBuilding' }
      },
      // 追溯管理
      {
        path: 'trace/query',
        name: 'TraceQuery',
        component: () => import('@/views/trace/TraceQuery.vue'),
        meta: { title: '追溯查询', icon: 'Search' }
      },
      {
        path: 'trace/andon',
        name: 'AndonCall',
        component: () => import('@/views/trace/AndonCall.vue'),
        meta: { title: '安东呼叫', icon: 'Bell' }
      },
      // 能源管理
      {
        path: 'energy/monitor',
        name: 'EnergyMonitor',
        component: () => import('@/views/energy/EnergyMonitor.vue'),
        meta: { title: '能源监控', icon: 'Lightning' }
      },
      // 供应链管理
      {
        path: 'scp/purchase',
        name: 'PurchaseOrderList',
        component: () => import('@/views/scp/PurchaseOrderList.vue'),
        meta: { title: '采购订单', icon: 'ShoppingCart' }
      },
      {
        path: 'scp/rfq',
        name: 'RFQList',
        component: () => import('@/views/scp/RFQList.vue'),
        meta: { title: '询价单', icon: 'PriceTag' }
      },
      {
        path: 'scp/supplier-quote',
        name: 'SupplierQuoteList',
        component: () => import('@/views/scp/SupplierQuoteList.vue'),
        meta: { title: '供应商报价', icon: 'Document' }
      },
      {
        path: 'scp/sales-order',
        name: 'SCPSalesOrderList',
        component: () => import('@/views/scp/SCPSalesOrderList.vue'),
        meta: { title: '销售订单', icon: 'DocumentCopy' }
      },
      {
        path: 'scp/supplier-kpi',
        name: 'SupplierKPIList',
        component: () => import('@/views/scp/SupplierKPIList.vue'),
        meta: { title: '供应商绩效', icon: 'DataLine' }
      },
      {
        path: 'scp/customer-inquiry',
        name: 'CustomerInquiryList',
        component: () => import('@/views/scp/CustomerInquiryList.vue'),
        meta: { title: '客户询价', icon: 'Message' }
      },
      // 统一告警
      {
        path: 'alert/rules',
        name: 'AlertRulesList',
        component: () => import('@/views/alert/AlertRulesList.vue'),
        meta: { title: '告警规则', icon: 'SetUp' }
      },
      {
        path: 'alert/records',
        name: 'AlertRecordsList',
        component: () => import('@/views/alert/AlertRecordsList.vue'),
        meta: { title: '告警记录', icon: 'List' }
      },
      {
        path: 'alert/escalation',
        name: 'AlertEscalationList',
        component: () => import('@/views/alert/AlertEscalationList.vue'),
        meta: { title: '升级规则', icon: 'Top' }
      },
      {
        path: 'alert/statistics',
        name: 'AlertStatistics',
        component: () => import('@/views/alert/AlertStatistics.vue'),
        meta: { title: '告警统计', icon: 'DataAnalysis' }
      },
      {
        path: 'alert/notification',
        name: 'AlertNotification',
        component: () => import('@/views/alert/AlertNotification.vue'),
        meta: { title: '告警通知', icon: 'Message' }
      },
      // 流程管理
      {
        path: 'bpm/process',
        name: 'BPMProcessList',
        component: () => import('@/views/bpm/ProcessList.vue'),
        meta: { title: '流程模型', icon: 'FlowChart' }
      },
      {
        path: 'bpm/instance',
        name: 'BPMInstanceList',
        component: () => import('@/views/bpm/InstanceList.vue'),
        meta: { title: '流程实例', icon: 'Connection' }
      },
      {
        path: 'bpm/task',
        name: 'BPMTaskList',
        component: () => import('@/views/bpm/TaskList.vue'),
        meta: { title: '任务实例', icon: 'Tickets' }
      },
      // APS扩展
      {
        path: 'aps/rolling-config',
        name: 'RollingConfigList',
        component: () => import('@/views/aps/RollingConfigList.vue'),
        meta: { title: '滚动排程', icon: 'Refresh' }
      },
      {
        path: 'aps/delivery-analysis',
        name: 'DeliveryAnalysisList',
        component: () => import('@/views/aps/DeliveryAnalysisList.vue'),
        meta: { title: '交付分析', icon: 'TrendCharts' }
      },
      {
        path: 'aps/material-shortage',
        name: 'MaterialShortageList',
        component: () => import('@/views/aps/MaterialShortageList.vue'),
        meta: { title: '缺料分析', icon: 'Warning' }
      },
      {
        path: 'aps/shortage-rule',
        name: 'ShortageRuleList',
        component: () => import('@/views/aps/ShortageRuleList.vue'),
        meta: { title: '缺料规则', icon: 'SetUp' }
      },
      {
        path: 'aps/changeover-matrix',
        name: 'ChangeoverMatrixList',
        component: () => import('@/views/aps/ChangeoverMatrixList.vue'),
        meta: { title: '换型矩阵', icon: 'Grid' }
      },
      {
        path: 'aps/product-family',
        name: 'ProductFamilyList',
        component: () => import('@/views/aps/ProductFamilyList.vue'),
        meta: { title: '产品族', icon: 'Collection' }
      },
      // WMS扩展
      {
        path: 'wms/transfer',
        name: 'TransferOrderList',
        component: () => import('@/views/wms/TransferOrderList.vue'),
        meta: { title: '调拨管理', icon: 'Switch' }
      },
      {
        path: 'wms/stock-check',
        name: 'StockCheckList',
        component: () => import('@/views/wms/StockCheckList.vue'),
        meta: { title: '盘点管理', icon: 'DocumentChecked' }
      },
      // 质量扩展
      {
        path: 'quality/qrci',
        name: 'QRCIList',
        component: () => import('@/views/quality/QRCIList.vue'),
        meta: { title: 'QRCI质量闭环', icon: 'CircleCheck' }
      },
      {
        path: 'quality/lpa',
        name: 'LPAStandardList',
        component: () => import('@/views/quality/LPAStandardList.vue'),
        meta: { title: 'LPA分层审核', icon: 'Checked' }
      },
      {
        path: 'quality/dynamic-rule',
        name: 'DynamicRuleList',
        component: () => import('@/views/quality/DynamicRuleList.vue'),
        meta: { title: '动态规则', icon: 'Setting' }
      },
      {
        path: 'quality/aql',
        name: 'AQLList',
        component: () => import('@/views/quality/AQLList.vue'),
        meta: { title: 'AQL抽样标准', icon: 'DataLine' }
      },
      {
        path: 'quality/inspection-plans',
        name: 'QualityInspectionPlanList',
        component: () => import('@/views/quality/InspectionPlanList.vue'),
        meta: { title: '检验计划', icon: 'Schedule' }
      },
      {
        path: 'quality/lab/samples',
        name: 'LabSampleList',
        component: () => import('@/views/quality/lab/LabSampleList.vue'),
        meta: { title: '检测样品', icon: 'Box' }
      },
      {
        path: 'quality/lab/reports',
        name: 'LabReportList',
        component: () => import('@/views/quality/lab/LabReportList.vue'),
        meta: { title: '检测报告', icon: 'Document' }
      },
      {
        path: 'quality/lab-instrument',
        name: 'LabInstrumentList',
        component: () => import('@/views/quality/lab/LabInstrumentList.vue'),
        meta: { title: '检测仪器', icon: 'Scale' }
      },
      // EAM模块
      {
        path: 'eam/factory',
        name: 'FactoryList',
        component: () => import('@/views/eam/FactoryList.vue'),
        meta: { title: '厂区管理', icon: 'OfficeBuilding' }
      },
      {
        path: 'eam/equipment-org',
        name: 'EquipmentOrgList',
        component: () => import('@/views/eam/EquipmentOrgList.vue'),
        meta: { title: '设备层级', icon: 'Grid' }
      },
      {
        path: 'eam/downtime',
        name: 'DowntimeList',
        component: () => import('@/views/eam/DowntimeList.vue'),
        meta: { title: '停机记录', icon: 'Warning' }
      },
      // 供应商ASN
      {
        path: 'supplier/asn',
        name: 'ASNList',
        component: () => import('@/views/supplier/ASNList.vue'),
        meta: { title: 'ASN到货通知', icon: 'Bell' }
      },
      // 财务模块
      {
        path: 'fin/purchase-settlement',
        name: 'PurchaseSettlementList',
        component: () => import('@/views/fin/PurchaseSettlementList.vue'),
        meta: { title: '采购结算', icon: 'Document' }
      },
      {
        path: 'fin/sales-settlement',
        name: 'SalesSettlementList',
        component: () => import('@/views/fin/SalesSettlementList.vue'),
        meta: { title: '销售结算', icon: 'DocumentCopy' }
      },
      {
        path: 'fin/payment-request',
        name: 'PaymentRequestList',
        component: () => import('@/views/fin/PaymentRequestList.vue'),
        meta: { title: '付款申请', icon: 'Money' }
      },
      // 生产扩展
      {
        path: 'production/issue',
        name: 'ProductionIssueList',
        component: () => import('@/views/production/ProductionIssueList.vue'),
        meta: { title: '生产发料', icon: 'Upload' }
      },
      {
        path: 'production/return',
        name: 'ProductionReturnList',
        component: () => import('@/views/production/ProductionReturnList.vue'),
        meta: { title: '生产退料', icon: 'Download' }
      },
      // 设备扩展
      {
        path: 'equipment/inspection/templates',
        name: 'InspectionTemplateList',
        component: () => import('@/views/equipment/InspectionTemplateList.vue'),
        meta: { title: '点检模板', icon: 'Tickets' }
      },
      {
        path: 'equipment/inspection/plans',
        name: 'InspectionPlanList',
        component: () => import('@/views/equipment/InspectionPlanList.vue'),
        meta: { title: '点检计划', icon: 'Schedule' }
      },
      {
        path: 'equipment/inspection/records',
        name: 'InspectionRecordList',
        component: () => import('@/views/equipment/InspectionRecordList.vue'),
        meta: { title: '点检记录', icon: 'Document' }
      },
      {
        path: 'equipment/inspection/defects',
        name: 'InspectionDefectList',
        component: () => import('@/views/equipment/InspectionDefectList.vue'),
        meta: { title: '点检缺陷', icon: 'Warning' }
      },
      {
        path: 'equipment/gauge',
        name: 'GaugeList',
        component: () => import('@/views/equipment/GaugeList.vue'),
        meta: { title: '量检具管理', icon: 'Scale' }
      },
      // 报表模块
      {
        path: 'report/production-daily',
        name: 'ProductionDailyReport',
        component: () => import('@/views/report/ProductionDailyReport.vue'),
        meta: { title: '生产日报', icon: 'DataLine' }
      },
      {
        path: 'report/quality-weekly',
        name: 'QualityWeeklyReport',
        component: () => import('@/views/report/QualityWeeklyReport.vue'),
        meta: { title: '质量周报', icon: 'DataAnalysis' }
      },
      {
        path: 'report/oee',
        name: 'OEEReport',
        component: () => import('@/views/report/OEEReport.vue'),
        meta: { title: 'OEE报表', icon: 'TrendCharts' }
      },
      {
        path: 'report/delivery',
        name: 'DeliveryReport',
        component: () => import('@/views/report/DeliveryReport.vue'),
        meta: { title: '交付报表', icon: ' Truck' }
      },
      {
        path: 'report/andon',
        name: 'AndonReport',
        component: () => import('@/views/report/AndonReport.vue'),
        meta: { title: '安东报表', icon: 'Bell' }
      },
      // MES模块
      {
        path: 'mes/team',
        name: 'TeamList',
        component: () => import('@/views/mes/TeamList.vue'),
        meta: { title: '班组管理', icon: 'User' }
      },
      {
        path: 'mes/process-routes',
        name: 'ProcessRouteList',
        component: () => import('@/views/mes/ProcessRouteList.vue'),
        meta: { title: '工艺路线', icon: 'Guide' }
      },
      {
        path: 'mes/person-skill',
        name: 'PersonSkillList',
        component: () => import('@/views/mes/PersonSkillList.vue'),
        meta: { title: '人员能力', icon: 'Star' }
      },
      {
        path: 'mes/offline',
        name: 'OfflineList',
        component: () => import('@/views/mes/OfflineList.vue'),
        meta: { title: '产品离线', icon: 'Warning' }
      },
      {
        path: 'mes/material-trace',
        name: 'MaterialTrace',
        component: () => import('@/views/mes/MaterialTrace.vue'),
        meta: { title: '物料追溯', icon: 'Search' }
      },
      // AGV模块
      {
        path: 'agv/task',
        name: 'AGVTaskList',
        component: () => import('@/views/agv/TaskList.vue'),
        meta: { title: 'AGV任务', icon: 'Document' }
      },
      {
        path: 'agv/device',
        name: 'AGVDeviceList',
        component: () => import('@/views/agv/DeviceList.vue'),
        meta: { title: 'AGV设备', icon: 'Processor' }
      },
      {
        path: 'agv/location',
        name: 'AGVLocationList',
        component: () => import('@/views/agv/LocationList.vue'),
        meta: { title: '库位映射', icon: 'Location' }
      },
      // 系统集成
      {
        path: 'integration/interface-config',
        name: 'InterfaceConfigList',
        component: () => import('@/views/integration/InterfaceConfigList.vue'),
        meta: { title: '接口配置', icon: 'Interface' }
      },
      {
        path: 'integration/execution-log',
        name: 'ExecutionLogList',
        component: () => import('@/views/integration/ExecutionLogList.vue'),
        meta: { title: '执行日志', icon: 'List' }
      }
    ]
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/views/Error404.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫
router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore()

  // 设置页面标题
  document.title = (to.meta.title as string) || 'MOM3.0'

  // 公开路由
  if (to.meta.public) {
    next()
    return
  }

  // 检查登录状态
  if (!authStore.isLoggedIn) {
    next({ name: 'Login', query: { redirect: to.fullPath } })
    return
  }

  // 如果没有用户信息，获取一下
  if (!authStore.userInfo) {
    await authStore.getUserInfoAction()
  }

  next()
})

export default router
