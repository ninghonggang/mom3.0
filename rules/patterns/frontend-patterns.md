# 前端设计模式

> 本文件适用于 MOM3.0 Vue3 前端项目。

## 1. 组合式 API 模式

### 1.1 典型页面结构

```vue
<!-- 标准列表页组件 -->
<template>
  <div class="page-container">
    <!-- 搜索区域 -->
    <SearchArea @search="handleSearch" @reset="handleReset" />

    <!-- 工具栏 -->
    <Toolbar>
      <el-button type="primary" @click="handleAdd">新增</el-button>
    </Toolbar>

    <!-- 数据表格 -->
    <DataTable
      :data="tableData"
      :loading="loading"
      @edit="handleEdit"
      @delete="handleDelete"
    />

    <!-- 分页 -->
    <Pagination
      v-model:page="pagination.page"
      v-model:page-size="pagination.pageSize"
      :total="pagination.total"
      @change="loadData"
    />
  </div>
</template>

<script setup>
// 导入
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessageBox, ElMessage } from 'element-plus'

// 组合式函数
const router = useRouter()

// 状态
const loading = ref(false)
const tableData = ref([])
const searchForm = reactive({ name: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

// 方法
const loadData = async () => {
  loading.value = true
  try {
    const res = await fetchList({ ...searchForm, ...pagination })
    tableData.value = res.data.list
    pagination.total = res.data.total
  } catch (error) {
    ElMessage.error('加载失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.page = 1
  loadData()
}

const handleReset = () => {
  Object.assign(searchForm, { name: '', status: '' })
  handleSearch()
}

const handleAdd = () => {
  router.push('/path/to/add')
}

const handleEdit = (row) => {
  router.push(`/path/to/edit/${row.id}`)
}

const handleDelete = async (row) => {
  await ElMessageBox.confirm('确认删除?', '提示', { type: 'warning' })
  await deleteById(row.id)
  ElMessage.success('删除成功')
  loadData()
}

// 生命周期
onMounted(() => {
  loadData()
})
</script>
```

### 1.2 自定义 Hook

```javascript
// hooks/usePagination.js
import { reactive, computed } from 'vue'

export function usePagination(options = {}) {
  const defaultPageSize = options.defaultPageSize || 20

  const pagination = reactive({
    page: 1,
    pageSize: defaultPageSize,
    total: 0
  })

  const handlePageChange = (page) => {
    pagination.page = page
  }

  const handleSizeChange = (pageSize) => {
    pagination.pageSize = pageSize
    pagination.page = 1
  }

  const resetPagination = () => {
    pagination.page = 1
  }

  const setTotal = (total) => {
    pagination.total = total
  }

  return {
    pagination,
    handlePageChange,
    handleSizeChange,
    resetPagination,
    setTotal
  }
}

// hooks/useSearch.js
import { reactive, watch } from 'vue'

export function useSearch(loadFn, options = {}) {
  const { immediate = true, debounce = 300 } = options

  const searchForm = reactive({})
  const loading = ref(false)

  let debounceTimer = null

  const handleSearch = () => {
    if (debounceTimer) clearTimeout(debounceTimer)
    debounceTimer = setTimeout(() => {
      loadFn(searchForm)
    }, debounce)
  }

  const handleReset = (defaultValues = {}) => {
    Object.assign(searchForm, defaultValues)
    handleSearch()
  }

  if (immediate) {
    onMounted(() => {
      handleSearch()
    })
  }

  return {
    searchForm,
    loading,
    handleSearch,
    handleReset
  }
}
```

## 2. 状态管理模式

### 2.1 Pinia Store 规范

```javascript
// stores/user.js
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import * as api from '@/api/user'

export const useUserStore = defineStore('user', () => {
  // ========== State ==========
  const userList = ref([])
  const currentUser = ref(null)
  const loading = ref(false)
  const pagination = ref({ page: 1, pageSize: 20, total: 0 })

  // ========== Getters ==========
  const userCount = computed(() => userList.value.length)

  const activeUsers = computed(() =>
    userList.value.filter(u => u.status === 1)
  )

  const usersByRole = computed(() => (role) =>
    userList.value.filter(u => u.role === role)
  )

  // ========== Actions ==========
  const fetchUserList = async (params = {}) => {
    loading.value = true
    try {
      const res = await api.getUserList(params)
      userList.value = res.data.list
      pagination.value.total = res.data.total
    } finally {
      loading.value = false
    }
  }

  const getUserById = async (id) => {
    const res = await api.getUserById(id)
    currentUser.value = res.data
    return res.data
  }

  const createUser = async (data) => {
    await api.createUser(data)
    await fetchUserList()
  }

  const updateUser = async (id, data) => {
    await api.updateUser(id, data)
    await fetchUserList()
  }

  const deleteUser = async (id) => {
    await api.deleteUser(id)
    await fetchUserList()
  }

  // ========== Reset ==========
  const $reset = () => {
    userList.value = []
    currentUser.value = null
    pagination.value = { page: 1, pageSize: 20, total: 0 }
  }

  return {
    // State
    userList,
    currentUser,
    loading,
    pagination,
    // Getters
    userCount,
    activeUsers,
    usersByRole,
    // Actions
    fetchUserList,
    getUserById,
    createUser,
    updateUser,
    deleteUser,
    $reset
  }
})
```

### 2.2 Store 模块划分

```
store/
├── index.js           # 组合所有 store
├── user.js            # 用户相关
├── permission.js       # 权限相关
├── settings.js        # 系统设置
└── slices/            # 大型 store 拆分
    ├── user/
    │   ├── index.js
    │   ├── list.js
    │   └── detail.js
    └── order/
        ├── index.js
        ├── list.js
        └── detail.js
```

## 3. API 分层模式

### 3.1 API 模块划分

```javascript
// api/index.js - API 入口
import request from './request'
import system from './system'
import production from './production'
import quality from './quality'
import equipment from './equipment'

export default {
  system,
  production,
  quality,
  equipment
}

// api/production.js - 生产模块
import request from './request'

export const getWorkOrderList = (params) =>
  request.get('/production/orders', { params })

export const getWorkOrderById = (id) =>
  request.get(`/production/orders/${id}`)

export const createWorkOrder = (data) =>
  request.post('/production/orders', data)

export const updateWorkOrder = (id, data) =>
  request.put(`/production/orders/${id}`, data)

export const deleteWorkOrder = (id) =>
  request.delete(`/production/orders/${id}`)

export const reportProduction = (id, data) =>
  request.post(`/production/orders/${id}/report`, data)

// api/production.js - 批量操作
export const batchDeleteOrders = (ids) =>
  request.delete('/production/orders/batch', { data: { ids } })

export const exportOrders = (params) =>
  request.get('/production/orders/export', { params, responseType: 'blob' })

export const importOrders = (file) => {
  const formData = new FormData()
  formData.append('file', file)
  return request.post('/production/orders/import', formData, {
    headers: { 'Content-Type': 'multipart/form-data' }
  })
}
```

### 3.2 API 封装模式

```javascript
// utils/request.js
import axios from 'axios'
import { ElMessage, ElLoading } from 'element-plus'

const service = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: 30000
})

let loadingInstance = null
let requestCount = 0

const showLoading = () => {
  if (requestCount === 0) {
    loadingInstance = ElLoading.service({
      lock: true,
      text: '加载中...',
      background: 'rgba(255, 255, 255, 0.7)'
    })
  }
  requestCount++
}

const hideLoading = () => {
  requestCount--
  if (requestCount <= 0) {
    requestCount = 0
    loadingInstance?.close()
  }
}

// 请求拦截
service.interceptors.request.use(
  config => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }

    const tenantId = localStorage.getItem('tenantId')
    if (tenantId) {
      config.headers['X-Tenant-ID'] = tenantId
    }

    if (config.showLoading !== false) {
      showLoading()
    }

    return config
  },
  error => {
    hideLoading()
    return Promise.reject(error)
  }
)

// 响应拦截
service.interceptors.response.use(
  response => {
    hideLoading()

    const { code, message, data } = response.data

    if (code === 0 || code === 200) {
      return response.data
    }

    if (code === 401) {
      ElMessage.error('登录已过期，请重新登录')
      router.push('/login')
      return Promise.reject(new Error(message))
    }

    if (code === 403) {
      ElMessage.error('没有权限')
      return Promise.reject(new Error(message))
    }

    ElMessage.error(message || '请求失败')
    return Promise.reject(new Error(message))
  },
  error => {
    hideLoading()
    ElMessage.error(error.message || '网络错误')
    return Promise.reject(error)
  }
)

export default service
```

## 4. 表格驱动模式

### 4.1 配置化表格

```javascript
// 配置定义
const columns = [
  { type: 'selection', width: 55 },
  { prop: 'orderNo', label: '工单号', minWidth: 120 },
  { prop: 'materialName', label: '物料名称', minWidth: 150 },
  { prop: 'quantity', label: '数量', width: 100, align: 'right' },
  { prop: 'status', label: '状态', width: 100,
    formatter: (row) => STATUS_OPTIONS.find(s => s.value === row.status)?.label
  },
  { prop: 'createdAt', label: '创建时间', width: 160,
    formatter: (row) => formatDate(row.createdAt)
  },
  { prop: 'action', label: '操作', width: 150, fixed: 'right',
    slots: 'action'
  }
]

// 通用表格组件
<template>
  <el-table :data="data" v-bind="$attrs">
    <el-table-column
      v-for="col in columns"
      :key="col.prop"
      v-bind="col"
    >
      <template v-if="col.slots" #default="scope">
        <slot :name="col.slots" :row="scope.row" />
      </template>
    </el-table-column>
  </el-table>
</template>

<script setup>
defineProps({
  data: { type: Array, default: () => [] },
  columns: { type: Array, default: () => [] }
})
</script>
```

### 4.2 配置化搜索表单

```javascript
// 搜索配置
const searchConfig = [
  { type: 'input', prop: 'orderNo', label: '工单号', placeholder: '请输入工单号' },
  { type: 'input', prop: 'materialName', label: '物料名称', placeholder: '请输入物料名称' },
  { type: 'select', prop: 'status', label: '状态',
    options: STATUS_OPTIONS, placeholder: '请选择状态' },
  { type: 'date', prop: 'dateRange', label: '日期范围', type: 'daterange' },
  { type: 'tree-select', prop: 'deptId', label: '部门',
    data: deptTree, placeholder: '请选择部门' }
]

// 通用搜索组件
<template>
  <el-form :model="form" inline>
    <el-form-item
      v-for="item in config"
      :key="item.prop"
      :label="item.label"
    >
      <component
        :is="getComponent(item.type)"
        v-model="form[item.prop]"
        v-bind="item"
      />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="handleSearch">查询</el-button>
      <el-button @click="handleReset">重置</el-button>
    </el-form-item>
  </el-form>
</template>
```

## 5. 权限控制模式

### 5.1 权限指令

```javascript
// directives/hasPermission.js
import { useUserStore } from '@/store/user'

export default {
  mounted(el, binding) {
    const userStore = useUserStore()
    const permission = binding.value

    if (!hasPermission(permission)) {
      el.parentNode?.removeChild(el)
    }
  }
}

function hasPermission(permission) {
  const userStore = useUserStore()
  if (!permission) return true
  if (typeof permission === 'string') {
    return userStore.permissions.includes(permission)
  }
  if (Array.isArray(permission)) {
    return permission.some(p => userStore.permissions.includes(p))
  }
  return false
}
```

### 5.2 权限组件

```vue
<!-- components/PermissionButton.vue -->
<template>
  <el-tooltip v-if="!visible" :content="disabledText">
    <span>
      <el-button v-bind="$attrs" :disabled="true" />
    </span>
  </el-tooltip>
  <el-button v-else v-bind="$attrs" />
</template>

<script setup>
import { computed } from 'vue'
import { usePermission } from '@/hooks/usePermission'

const props = defineProps({
  permission: { type: [String, Array], default: '' },
  disabledText: { type: String, default: '没有权限' }
})

const { hasPermission } = usePermission()

const visible = computed(() => {
  if (!props.permission) return true
  return hasPermission(props.permission)
})
</script>
```

## 6. 路由守卫模式

### 6.1 路由权限

```javascript
// router/permission.js
import router from './index'
import { useUserStore } from '@/store/user'
import { getToken } from '@/utils/auth'

const whiteList = ['/login', '/404']

router.beforeEach(async (to, from, next) => {
  const hasToken = getToken()

  if (hasToken) {
    if (to.path === '/login') {
      next({ path: '/' })
    } else {
      const userStore = useUserStore()
      const hasRoles = userStore.roles.length > 0

      if (hasRoles) {
        next()
      } else {
        try {
          await userStore.getUserInfo()
          const accessRoutes = await userStore.generateRoutes()
          accessRoutes.forEach(route => {
            router.addRoute(route)
          })
          next({ ...to, replace: true })
        } catch (error) {
          await userStore.logout()
          next(`/login?redirect=${to.path}`)
        }
      }
    }
  } else {
    if (whiteList.includes(to.path)) {
      next()
    } else {
      next(`/login?redirect=${to.path}`)
    }
  }
})
```

## 7. 依赖注入模式

### 7.1 Provide/Inject

```javascript
// contexts/ThemeContext.vue
<script setup>
import { provide, ref } from 'vue'

const theme = ref('light')

const setTheme = (newTheme) => {
  theme.value = newTheme
  document.documentElement.setAttribute('data-theme', newTheme)
}

provide('theme', {
  theme,
  setTheme
})
</script>

// 子组件使用
<script setup>
import { inject } from 'vue'

const { theme, setTheme } = inject('theme')
</script>
```

## 8. 错误处理模式

### 8.1 统一错误处理

```javascript
// hooks/useErrorHandler.js
import { ElMessage } from 'element-plus'

export function useErrorHandler() {
  const handleError = (error, options = {}) => {
    const { showMessage = true, log = true } = options

    if (log) {
      console.error('[Error]', error)
    }

    if (showMessage) {
      const message = error?.response?.data?.message
        || error?.message
        || '操作失败'
      ElMessage.error(message)
    }
  }

  const handleSuccess = (message = '操作成功') => {
    ElMessage.success(message)
  }

  return {
    handleError,
    handleSuccess
  }
}

// 使用
const { handleError, handleSuccess } = useErrorHandler()

const handleSave = async () => {
  try {
    await saveData(data)
    handleSuccess()
    loadData()
  } catch (error) {
    handleError(error)
  }
}
```

## 9. 异步数据加载模式

### 9.1 组合式数据加载

```javascript
// hooks/useAsyncData.js
import { ref, watch } from 'vue'

export function useAsyncData(fetchFn, options = {}) {
  const { immediate = true, debounce = 0, onSuccess, onError } = options

  const data = ref(null)
  const loading = ref(false)
  const error = ref(null)

  const execute = async (...args) => {
    loading.value = true
    error.value = null

    try {
      const result = await fetchFn(...args)
      data.value = result
      onSuccess?.(result)
      return result
    } catch (e) {
      error.value = e
      onError?.(e)
      throw e
    } finally {
      loading.value = false
    }
  }

  if (immediate) {
    execute()
  }

  return {
    data,
    loading,
    error,
    execute
  }
}

// 使用
const { data, loading, execute } = useAsyncData(
  (params) => fetchUserList(params),
  { immediate: true }
)

// 重新加载
const handleRefresh = () => execute(searchParams)
```
