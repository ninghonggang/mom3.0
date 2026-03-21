# 表格组件规范

> 本文件适用于 MOM3.0 Vue3 前端项目。

## 1. 标准表格结构

### 1.1 基础表格

```vue
<template>
  <div class="table-container">
    <!-- 搜索区域 -->
    <div class="search-area">
      <SearchForm @search="handleSearch" @reset="handleReset" />
    </div>

    <!-- 工具栏 -->
    <div class="toolbar">
      <el-button type="primary" @click="handleAdd">新增</el-button>
      <el-button type="danger" :disabled="!selectedRows.length" @click="handleBatchDelete">
        批量删除
      </el-button>
      <el-button @click="handleExport">导出</el-button>
    </div>

    <!-- 表格 -->
    <el-table
      ref="tableRef"
      v-loading="loading"
      :data="tableData"
      :height="tableHeight"
      stripe
      border
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column type="index" label="序号" width="60" align="center" />
      <el-table-column prop="name" label="名称" min-width="150" show-overflow-tooltip />
      <el-table-column prop="code" label="编码" width="120" />
      <el-table-column prop="status" label="状态" width="100" align="center">
        <template #default="{ row }">
          <el-tag :type="row.status === 1 ? 'success' : 'info'">
            {{ row.status === 1 ? '启用' : '禁用' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180" fixed="right" align="center">
        <template #default="{ row }">
          <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
          <el-button link type="danger" @click="handleDelete(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页 -->
    <div class="pagination-wrapper">
      <el-pagination
        v-model:current-page="pagination.page"
        v-model:page-size="pagination.pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :total="pagination.total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handlePageChange"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'

// 状态
const loading = ref(false)
const tableData = ref([])
const selectedRows = ref([])
const tableRef = ref()

const searchForm = reactive({
  name: '',
  status: ''
})

const pagination = reactive({
  page: 1,
  pageSize: 20,
  total: 0
})

// 计算表格高度
const tableHeight = computed(() => {
  return 'calc(100vh - 280px)'
})

// 加载数据
const loadData = async () => {
  loading.value = true
  try {
    const res = await fetchList({
      ...searchForm,
      page: pagination.page,
      pageSize: pagination.pageSize
    })
    tableData.value = res.data.list
    pagination.total = res.data.total
  } finally {
    loading.value = false
  }
}

// 事件处理
const handleSearch = () => {
  pagination.page = 1
  loadData()
}

const handleReset = () => {
  Object.assign(searchForm, { name: '', status: '' })
  handleSearch()
}

const handleSelectionChange = (rows) => {
  selectedRows.value = rows
}

const handleSizeChange = () => {
  pagination.page = 1
  loadData()
}

const handlePageChange = () => {
  loadData()
}

const handleAdd = () => {
  // 跳转新增页或打开弹窗
}

const handleEdit = (row) => {
  // 跳转编辑页或打开弹窗
}

const handleDelete = async (row) => {
  // 删除确认
  await deleteById(row.id)
  loadData()
}

const handleBatchDelete = async () => {
  // 批量删除
}

const handleExport = () => {
  // 导出
}

// 生命周期
onMounted(() => {
  loadData()
})
</script>

<style scoped>
.table-container {
  background: #fff;
  padding: 16px;
  border-radius: 4px;
}

.search-area {
  margin-bottom: 16px;
}

.toolbar {
  margin-bottom: 16px;
  display: flex;
  gap: 8px;
}

.pagination-wrapper {
  margin-top: 16px;
  display: flex;
  justify-content: flex-end;
}
</style>
```

## 2. 列配置

### 2.1 列类型

| 类型 | 说明 | 使用场景 |
|------|------|----------|
| `selection` | 多选框 | 批量操作 |
| `index` | 序号 | 列表展示 |
| `expand` | 展开行 | 详情展示 |
| `prop` | 普通列 | 常规数据 |

### 2.2 列属性

```javascript
const columns = [
  // 多选列
  { type: 'selection', width: 55, fixed: 'left' },

  // 序号列
  { type: 'index', label: '序号', width: 60, fixed: 'left' },

  // 普通文本列
  { prop: 'name', label: '名称', min-width: 150, show-overflow-tooltip: true },

  // 带格式化的列
  { prop: 'amount', label: '金额', width: 120, align: 'right',
    formatter: (row) => `¥${row.amount.toFixed(2)}`
  },

  // 自定义列
  { prop: 'status', label: '状态', width: 100,
    slots: { default: 'status' }
  },

  // 操作列
  { label: '操作', width: 180, fixed: 'right',
    slots: { default: 'action' }
  }
]
```

## 3. 单元格渲染

### 3.1 状态标签

```vue
<template #status="{ row }">
  <el-tag :type="getStatusType(row.status)">
    {{ getStatusText(row.status) }}
  </el-tag>
</template>

<script setup>
const STATUS_MAP = {
  1: { text: '启用', type: 'success' },
  0: { text: '禁用', type: 'info' },
  pending: { text: '待处理', type: 'warning' },
  processing: { text: '处理中', type: 'primary' },
  completed: { text: '已完成', type: 'success' },
  failed: { text: '失败', type: 'danger' }
}

const getStatusType = (status) => STATUS_MAP[status]?.type || 'info'
const getStatusText = (status) => STATUS_MAP[status]?.text || status
</script>
```

### 3.2 颜色指示

```vue
<template #level="{ row }">
  <div class="level-indicator">
    <span class="level-dot" :class="'level-' + row.level"></span>
    {{ getLevelText(row.level) }}
  </div>
</template>

<style scoped>
.level-dot {
  display: inline-block;
  width: 8px;
  height: 8px;
  border-radius: 50%;
  margin-right: 4px;
}
.level-dot.level-high { background: #F56C6C; }
.level-dot.level-medium { background: #E6A23C; }
.level-dot.level-low { background: #67C23A; }
</style>
```

### 3.3 进度条

```vue
<template #progress="{ row }">
  <el-progress
    :percentage="row.progress"
    :stroke-width="12"
    :text-inside="true"
    :status="row.progress >= 100 ? 'success' : undefined"
  />
</template>
```

### 3.4 图片展示

```vue
<template #image="{ row }">
  <el-image
    :src="row.image"
    :preview-src-list="[row.image]"
    fit="cover"
    style="width: 40px; height: 40px"
  />
</template>
```

### 3.5 链接/按钮

```vue
<template #action="{ row }">
  <el-button link type="primary" @click="handleView(row)">查看</el-button>
  <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
  <el-button link type="danger" @click="handleDelete(row)">删除</el-button>
</template>
```

## 4. 高级特性

### 4.1 行展开

```vue
<el-table :data="tableData" @expand-change="handleExpand">
  <el-table-column type="expand">
    <template #default="{ row }">
      <div class="expand-content">
        <el-descriptions :column="2" border>
          <el-descriptions-item label="创建时间">{{ row.createdAt }}</el-descriptions-item>
          <el-descriptions-item label="更新时间">{{ row.updatedAt }}</el-descriptions-item>
          <el-descriptions-item label="备注" :span="2">{{ row.remark || '-' }}</el-descriptions-item>
        </el-descriptions>
      </div>
    </template>
  </el-table-column>
  <el-table-column prop="name" label="名称" />
</el-table>
```

### 4.2 列动态显示

```vue
<el-table :data="tableData">
  <el-table-column
    v-for="col in visibleColumns"
    :key="col.prop"
    :prop="col.prop"
    :label="col.label"
    :width="col.width"
  />
</el-table>

<script setup>
const columns = [
  { prop: 'name', label: '名称', visible: true },
  { prop: 'code', label: '编码', visible: true },
  { prop: 'remark', label: '备注', visible: false }
]

const visibleColumns = computed(() => columns.filter(c => c.visible))
</script>
```

### 4.3 拖拽排序

```vue
<el-table :data="tableData" row-key="id">
  <el-table-column type="index" width="50" />
  <el-table-column prop="name" label="名称" />
  <el-table-column prop="sort" label="排序" width="100">
    <template #default>
      <el-icon class="drag-handle"><Rank /></el-icon>
    </template>
  </el-table-column>
</el-table>

<script setup>
import { VueDraggable } from 'vuedraggable'

// 需要安装 vuedraggable
// npm install vuedraggable@next
</script>
```

### 4.4 虚拟滚动

```vue
<!-- 大数据量使用 el-table-v2 -->
<el-table-v2
  :columns="columns"
  :data="largeDataList"
  :height="600"
  :estimated-row-height="44"
  :width="800"
/>
```

## 5. 表格事件

### 5.1 常用事件

| 事件名 | 说明 | 参数 |
|--------|------|------|
| `selection-change` | 选择变化 | selection |
| `row-click` | 行点击 | row, column, event |
| `row-dblclick` | 行双击 | row, column, event |
| `cell-click` | 单元格点击 | row, column, event |
| `sort-change` | 排序变化 | { prop, order } |
| `filter-change` | 筛选变化 | filters |

### 5.2 事件处理

```javascript
const handleRowClick = (row) => {
  // 跳转详情
  router.push(`/detail/${row.id}`)
}

const handleSelectionChange = (selection) => {
  selectedRows.value = selection
  // 批量操作按钮状态
}

const handleSortChange = ({ prop, order }) => {
  searchForm.sortField = prop
  searchForm.sortOrder = order
  loadData()
}
```

## 6. 分页配置

### 6.1 分页参数

```javascript
const pagination = reactive({
  page: 1,
  pageSize: 20,
  total: 0,
  // 可选：自定义布局
  // layout: 'total, sizes, prev, pager, next, jumper'
})
```

### 6.2 分页事件

```javascript
const handleSizeChange = (size) => {
  pagination.pageSize = size
  pagination.page = 1
  loadData()
}

const handlePageChange = (page) => {
  pagination.page = page
  loadData()
}
```

## 7. 表格操作

### 7.1 行选择

```javascript
// 获取选中行
const getSelectedRows = () => {
  return tableRef.value?.getSelectionRows()
}

// 选中某行
const toggleRowSelection = (row) => {
  tableRef.value?.toggleRowSelection(row)
}

// 全选/取消全选
const toggleAllSelection = () => {
  tableRef.value?.toggleAllSelection()
}

// 清除选择
const clearSelection = () => {
  tableRef.value?.clearSelection()
}
```

### 7.2 滚动操作

```javascript
// 滚动到顶部
const scrollToTop = () => {
  tableRef.value?.setScrollTop(0)
}

// 滚动到指定行
const scrollToRow = (row) => {
  tableRef.value?.scrollToRow(row)
}
```

### 7.3 排序操作

```javascript
// 设置排序
const setSort = (prop, order) => {
  tableRef.value?.sort(prop, order)
}

// 清除排序
const clearSort = () => {
  tableRef.value?.clearSort()
}
```

## 8. 性能优化

### 8.1 大数据处理

```vue
<script setup>
// 使用 Web Worker 处理大数据
import { ref, onMounted } from 'vue'

const largeDataList = ref([])

onMounted(() => {
  // 后台线程处理
  Worker.postMessage({ type: 'load', params })

  Worker.onmessage = (e) => {
    largeDataList.value = e.data
  }
})
</script>

<!-- 使用虚拟滚动 -->
<el-table-v2 :data="largeDataList" ... />
```

### 8.2 列冻结

```vue
<!-- 固定左侧列 -->
<el-table-column fixed prop="name" label="名称" />

<!-- 固定右侧列 -->
<el-table-column prop="action" label="操作" fixed="right" />
```

## 9. 无障碍设计

### 9.1 键盘导航

```vue
<!-- 使用键盘操作 -->
<el-table
  @keydown.enter="handleEnter"
  @keydown.escape="handleEscape"
>
```

### 9.2 屏幕阅读器

```vue
<!-- 为表格添加描述 -->
<el-table
  aria-label="用户列表"
  :aria-describedby="'table-caption'"
>
  <template #empty>
    <span id="table-caption">暂无数据</span>
  </template>
</el-table>
```
