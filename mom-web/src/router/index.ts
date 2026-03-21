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
      // 主数据管理
      {
        path: 'mdm/material',
        name: 'MaterialList',
        component: () => import('@/views/mdm/MaterialList.vue'),
        meta: { title: '物料管理', icon: 'Box' }
      },
      {
        path: 'mdm/workshop',
        name: 'WorkshopList',
        component: () => import('@/views/mdm/WorkshopList.vue'),
        meta: { title: '车间管理', icon: 'OfficeBuilding' }
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
      // 设备管理
      {
        path: 'equipment',
        name: 'EquipmentList',
        component: () => import('@/views/equipment/EquipmentList.vue'),
        meta: { title: '设备台账', icon: 'Monitor' }
      },
      // 仓储管理
      {
        path: 'wms/warehouse',
        name: 'WarehouseList',
        component: () => import('@/views/wms/WarehouseList.vue'),
        meta: { title: '仓库管理', icon: 'House' }
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
      // 质量管理
      {
        path: 'quality/iqc',
        name: 'IQCList',
        component: () => import('@/views/quality/IQCList.vue'),
        meta: { title: 'IQC检验', icon: 'CircleCheck' }
      },
      // APS计划
      {
        path: 'aps/mps',
        name: 'MPSList',
        component: () => import('@/views/aps/MPSList.vue'),
        meta: { title: 'MPS计划', icon: 'Calendar' }
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

  next()
})

export default router
