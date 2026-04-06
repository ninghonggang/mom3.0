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
