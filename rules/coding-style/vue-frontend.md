# Vue3 前端编码规范

> 本文件继承 [common/coding-style.md](../../common/coding-style.md) 并针对 MOM3.0 Vue3 前端项目进行扩展。

## 1. 项目结构

```
mom-web/src/
├── main.ts                    # 入口文件
├── App.vue                    # 根组件
├── router/                    # Vue Router路由配置
├── stores/                    # Pinia状态管理
│   ├── auth.ts              # 认证状态
│   ├── app.ts               # 应用全局状态
│   └── websocket.ts         # WebSocket状态
├── api/                      # API请求（按模块）
│   ├── system.js           # 系统管理
│   ├── mdm.js              # 主数据
│   ├── production.js       # 生产执行
│   ├── aps.js              # 计划排程
│   ├── quality.js          # 质量管理
│   ├── equipment.js        # 设备管理
│   ├── wms.js              # 仓储管理
│   ├── trace.js            # 追溯管理
│   ├── andon.js            # 安东系统
│   ├── datacollect.js      # 数据采集
│   ├── report.js           # 报表分析
│   ├── digital-twin.js     # 数字孪生
│   ├── lowcode.js          # 低代码平台
│   ├── ai.js               # AI助手
│   └── energy.js           # 能源管理
├── views/                   # 页面视图（按模块）
│   ├── system/            # 系统管理
│   ├── mdm/               # 主数据
│   ├── production/         # 生产执行
│   ├── aps/               # 计划排程
│   ├── quality/           # 质量管理
│   ├── equipment/         # 设备管理
│   ├── wms/               # 仓储管理
│   ├── trace/             # 追溯管理
│   ├── andon/             # 安东系统
│   ├── datacollect/       # 数据采集
│   ├── report/            # 报表分析
│   ├── digital-twin/      # 数字孪生
│   ├── lowcode/           # 低代码平台
│   ├── ai/                # AI助手
│   ├── energy/            # 能源管理
│   └── dashboard/         # 综合看板
├── components/             # 公共组件
│   ├── layout/            # 布局组件
│   ├── charts/            # ECharts封装
│   ├── table/             # 增强表格
│   ├── form/              # 低代码表单渲染
│   ├── upload/            # 文件上传
│   └── chat/              # AI聊天组件
├── composables/            # 组合式函数
│   ├── useWebSocket.ts
│   ├── usePermission.ts
│   ├── useTable.ts
│   └── useChart.ts
├── utils/                 # 工具函数
└── styles/                 # 全局样式
```

## 2. 命名规范

### 文件命名 (CRITICAL)
- 组件文件: `PascalCase` - `UserList.vue`, `ProductionOrder.vue`
- 工具函数: `camelCase` - `formatDate.js`, `validateForm.js`
- 配置文件: `camelCase` - `router.js`, `store.js`
- 常量文件: `UPPER_SNAKE_CASE` - `apiConstants.js`

### 组件命名 (CRITICAL)
- 基础组件: `BaseXxx` - `BaseTable`, `BaseForm`, `BaseModal`
- 业务组件: `XxxTable`, `XxxForm` - `UserTable`, `OrderForm`
- 页面组件: `XxxList`, `XxxDetail` - `UserList`, `UserDetail`

### 变量命名 (HIGH)
- 组件 Props: `camelCase` - `userList`, `searchParams`
- 组件 Emits: `camelCase` - `onSearch`, `onSubmit`
- 事件处理: `handleXxx` - `handleSearch`, `handleDelete`
- 异步请求: `fetchXxx`, `loadXxx` - `fetchUserList`, `loadOrderDetail`

## 3. Vue 3 Composition API 规范

### Setup 语法 (CRITICAL)

```vue
<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/store/user'
import { ElMessage } from 'element-plus'

// 1. 组合式函数
const router = useRouter()
const userStore = useUserStore()

// 2. 响应式状态
const loading = ref(false)
const tableData = ref([])
const searchForm = ref({
  name: '',
  status: ''
})

// 3. 计算属性
const hasPermission = computed(() => userStore.hasPermission('user:delete'))

// 4. 方法
const handleSearch = async () => {
  loading.value = true
  try {
    const res = await fetchUserList(searchForm.value)
    tableData.value = res.data
  } catch (error) {
    ElMessage.error('加载失败')
  } finally {
    loading.value = false
  }
}

// 5. 生命周期
onMounted(() => {
  handleSearch()
})
</script>
```

### Props 定义 (CRITICAL)

```vue
<script setup>
// ✅ 正确: 使用 defineProps 带类型定义
const props = defineProps({
  title: {
    type: String,
    required: true
  },
  data: {
    type: Array,
    default: () => []
  },
  loading: {
    type: Boolean,
    default: false
  }
})

// 使用 props.title 访问
</script>
```

### Emits 定义 (CRITICAL)

```vue
<script setup>
// ✅ 正确: 使用 defineEmits
const emit = defineEmits(['search', 'row-click', 'update:visible'])

const handleConfirm = () => {
  emit('search', searchForm.value)
}

const handleRowClick = (row) => {
  emit('row-click', row)
}
</script>
```

## 4. 组件设计原则

### Single Responsibility (CRITICAL)
- 每个组件只负责一个功能
- 超过 300 行必须拆分

### Props 传递 (HIGH)

```vue
<!-- ✅ 正确: 简洁的 Props 传递 -->
<BaseTable :data="tableData" :loading="loading" @row-click="handleRowClick" />

<!-- ❌ 错误: 传递不必要的 props -->
<BaseTable
  :data="tableData"
  :loading="loading"
  :pagination="pagination"
  :configs="configs"  <!-- 过于复杂 -->
  @search="handleSearch"
/>
```

### 插槽使用 (HIGH)

```vue
<!-- 命名插槽 -->
<template #status="{ row }">
  <el-tag :type="getStatusType(row.status)">
    {{ getStatusText(row.status) }}
  </el-tag>
</template>

<!-- 默认插槽 -->
<template #default="{ row }">
  <span>{{ row.name }}</span>
</template>
```

## 5. 状态管理 (Pinia)

### Store 结构 (HIGH)

```javascript
// stores/user.js
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { getUserList, createUser, updateUser, deleteUser } from '@/api/user'

export const useUserStore = defineStore('user', () => {
  // State
  const userList = ref([])
  const currentUser = ref(null)
  const loading = ref(false)

  // Getters
  const userCount = computed(() => userList.value.length)
  const adminList = computed(() => userList.value.filter(u => u.role === 'admin'))

  // Actions
  const fetchUserList = async (params) => {
    loading.value = true
    try {
      const res = await getUserList(params)
      userList.value = res.data
    } finally {
      loading.value = false
    }
  }

  const addUser = async (data) => {
    await createUser(data)
    await fetchUserList()
  }

  return {
    userList,
    currentUser,
    loading,
    userCount,
    adminList,
    fetchUserList,
    addUser
  }
})
```

## 6. API 请求规范

### API 模块划分 (CRITICAL)

```javascript
// api/system.js - 系统管理模块
import request from '@/utils/request'

export const getUserList = (params) => request.get('/users', { params })
export const getUserById = (id) => request.get(`/users/${id}`)
export const createUser = (data) => request.post('/users', data)
export const updateUser = (id, data) => request.put(`/users/${id}`, data)
export const deleteUser = (id) => request.delete(`/users/${id}`)
```

### 请求封装 (HIGH)

```javascript
// utils/request.js
import axios from 'axios'
import { ElMessage } from 'element-plus'

const service = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: 30000
})

// 请求拦截器
service.interceptors.request.use(
  config => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  error => Promise.reject(error)
)

// 响应拦截器
service.interceptors.response.use(
  response => {
    const { code, message, data } = response.data
    if (code === 0) {
      return response.data
    }
    ElMessage.error(message || '请求失败')
    return Promise.reject(new Error(message))
  },
  error => {
    ElMessage.error(error.message || '网络错误')
    return Promise.reject(error)
  }
)

export default service
```

## 7. 表格组件规范

### 标准表格模式 (CRITICAL)

```vue
<template>
  <div class="table-container">
    <!-- 搜索区域 -->
    <div class="search-area">
      <el-form :model="searchForm" inline>
        <el-form-item label="名称">
          <el-input v-model="searchForm.name" placeholder="请输入名称" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="启用" value="1" />
            <el-option label="禁用" value="0" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 操作栏 -->
    <div class="toolbar">
      <el-button type="primary" @click="handleAdd">新增</el-button>
      <el-button type="danger" :disabled="!selectedRows.length" @click="handleBatchDelete">
        批量删除
      </el-button>
    </div>

    <!-- 表格 -->
    <el-table
      v-loading="loading"
      :data="tableData"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column prop="name" label="名称" min-width="120" />
      <el-table-column prop="status" label="状态" width="100">
        <template #default="{ row }">
          <el-tag :type="row.status === 1 ? 'success' : 'info'">
            {{ row.status === 1 ? '启用' : '禁用' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180" fixed="right">
        <template #default="{ row }">
          <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
          <el-button link type="danger" @click="handleDelete(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页 -->
    <div class="pagination">
      <el-pagination
        v-model:current-page="pagination.page"
        v-model:page-size="pagination.pageSize"
        :total="pagination.total"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'

// State
const loading = ref(false)
const tableData = ref([])
const selectedRows = ref([])
const searchForm = reactive({ name: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

// Methods
const handleSearch = async () => {
  pagination.page = 1
  await loadData()
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getList({ ...searchForm, ...pagination })
    tableData.value = res.data.list
    pagination.total = res.data.total
  } finally {
    loading.value = false
  }
}

const handleSelectionChange = (rows) => {
  selectedRows.value = rows
}
</script>
```

## 8. 表单组件规范

### 表单验证规则 (CRITICAL)

```javascript
// 表单验证规则
const rules = {
  name: [
    { required: true, message: '请输入名称', trigger: 'blur' },
    { min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur' }
  ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ],
  status: [
    { required: true, message: '请选择状态', trigger: 'change' }
  ]
}
```

### Dialog 表单模式 (HIGH)

```vue
<template>
  <el-dialog
    v-model="dialogVisible"
    :title="dialogTitle"
    width="600px"
    @close="handleClose"
  >
    <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
      <el-form-item label="名称" prop="name">
        <el-input v-model="formData.name" placeholder="请输入名称" />
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-switch v-model="formData.status" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="handleClose">取消</el-button>
      <el-button type="primary" :loading="submitLoading" @click="handleSubmit">
        确定
      </el-button>
    </template>
  </el-dialog>
</template>
```

## 9. 样式规范

### CSS 作用域 (CRITICAL)

```vue
<style scoped>
/* ✅ 正确: 使用 scoped */
.table-container {
  background: #fff;
  padding: 16px;
  border-radius: 4px;
}

.table-container :deep(.el-table) {
  font-size: 14px;
}
</style>
```

### CSS 变量 (HIGH)

```css
/* index.css 或 variables.css */
:root {
  /* 主色调 */
  --primary-color: #409eff;
  --success-color: #67c23a;
  --warning-color: #e6a23c;
  --danger-color: #f56c6c;
  --info-color: #909399;

  /* 文字 */
  --text-primary: #303133;
  --text-regular: #606266;
  --text-secondary: #909399;

  /* 边框 */
  --border-color: #dcdfe6;
  --border-light: #e4e7ed;
}
```

## 10. 性能优化 (HIGH)

### 虚拟滚动 (大列表)

```vue
<template>
  <!-- 大量数据使用虚拟滚动 -->
  <el-table-v2
    :columns="columns"
    :data="largeDataList"
    :height="400"
    :estimated-row-height="50"
  />
</template>
```

### 图片懒加载

```vue
<img v-lazy="imageUrl" alt="image" />
```

### KeepAlive 缓存

```vue
<router-view v-slot="{ Component }">
  <keep-alive :include="cachedViews">
    <component :is="Component" />
  </keep-alive>
</router-view>
```

## 11. 权限控制

### 按钮级权限 (CRITICAL)

```vue
<template>
  <!-- 使用 v-hasPermission 指令 -->
  <el-button v-has-permission="'user:create'" type="primary" @click="handleAdd">
    新增
  </el-button>
  <el-button v-has-permission="'user:delete'" type="danger" @click="handleDelete">
    删除
  </el-button>
</template>

<script setup>
// 或使用 composable
const { hasPermission } = usePermission()

const canDelete = hasPermission('user:delete')
</script>
```

## 12. 错误处理

### 全局错误边界

```javascript
// main.js
import { globalErrorHandler } from '@/utils/error'

// Vue 错误处理
app.config.errorHandler = (err, instance, info) => {
  globalErrorHandler(err, { info, route: router.currentRoute.value })
}

// Promise 未捕获
window.addEventListener('unhandledrejection', (event) => {
  globalErrorHandler(event.reason, { type: 'unhandledrejection' })
})
```
