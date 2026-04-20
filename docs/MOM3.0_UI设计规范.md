# MOM3.0 UI设计规范

**版本**: V1.0 | **项目**: 闻荫科技MOM3.0 | **参考**: Element Plus + 自定义主题

> 本文档定义MOM3.0前端UI设计规范，所有模块页面开发均需遵循此规范。

---

## 目录

1. [设计原则](#1-设计原则)
2. [色彩系统](#2-色彩系统)
3. [字体系统](#3-字体系统)
4. [间距系统](#4-间距系统)
5. [组件规范](#5-组件规范)
6. [页面布局规范](#6-页面布局规范)
7. [状态设计](#7-状态设计)
8. [模块UI规范索引](#8-模块ui规范索引)

---

## 1. 设计原则

### 1.1 核心原则

- **清晰优先**: 信息层次分明，关键数据一眼可见
- **高效操作**: 常用操作3步内完成
- **一致性**: 同类场景统一交互模式
- **容错性**: 操作可撤销，错误有提示

### 1.2 设计类型

本系统属于 **工业MES应用** — 数据密集、操作导向、效率优先

**设计特征**:
- 冷静克制的视觉层次，避免过度装饰
- 深色顶栏 + 浅色内容区
- 状态色彩语义化（红=告警，绿=正常，黄=警告）
- 大面积数据表格 + 紧凑工具栏
- 操作按钮突出，次要操作收敛

---

## 2. 色彩系统

### 2.1 主色（Theme Colors）

使用 Element Plus 主题变量覆盖方式：

```css
/* src/styles/variables.scss */
:root {
  /* 主色 - 深蓝（工业稳重感） */
  --el-color-primary: #2563eb;
  --el-color-primary-light-3: #60a5fa;
  --el-color-primary-light-5: #93c5fd;
  --el-color-primary-light-7: #bfdbfe;
  --el-color-primary-light-8: #dbeafe;
  --el-color-primary-light-9: #eff6ff;
  --el-color-primary-dark-2: #1d4ed8;

  /* 功能色 */
  --el-color-success: #10b981;   /* 绿色 - 正常/完成 */
  --el-color-warning: #f59e0b;   /* 橙色 - 警告/待处理 */
  --el-color-danger: #ef4444;    /* 红色 - 异常/紧急 */
  --el-color-info: #6b7280;      /* 灰色 - 中性信息 */
}
```

### 2.2 语义色（Status Colors）

| 用途 | 颜色 | Hex | Element Plus变量 |
|------|------|-----|----------------|
| 正常/合格 | 绿色 | `#10b981` | `--el-color-success` |
| 警告/待处理 | 橙色 | `#f59e0b` | `--el-color-warning` |
| 异常/不合格 | 红色 | `#ef4444` | `--el-color-danger` |
| 进行中/启用 | 蓝色 | `#2563eb` | `--el-color-primary` |
| 已完成/禁用 | 灰色 | `#6b7280` | `--el-color-info` |

### 2.3 Andon灯色

| 状态 | 颜色 | Hex | 用途 |
|------|------|-----|------|
| 红色 | `#ef4444` | 紧急呼叫 | 设备故障/质量异常 |
| 橙色 | `#f59e0b` | 等待处理 | 物料短缺/技术指导 |
| 黄色 | `#eab308` | 提醒 | 保养提示/换型 |
| 绿色 | `#22c55e` | 正常运行 | 无异常 |
| 蓝色 | `#3b82f6` | 信息 | 进度提示 |

### 2.4 背景色

| 区域 | 颜色 | Hex |
|------|------|-----|
| 顶栏背景 | 深蓝灰 | `#1e293b` |
| 侧边菜单背景 | 深灰 | `#1f2937` |
| 页面背景 | 浅灰白 | `#f1f5f9` |
| 卡片背景 | 白色 | `#ffffff` |
| 表格斑马纹 | 极浅灰 | `#f8fafc` |
| 边框色 | 中灰 | `#e2e8f0` |

---

## 3. 字体系统

### 3.1 字体栈

```css
/* 中文优先 */
--el-font-family: "PingFang SC", "Microsoft YaHei", "Helvetica Neue", Helvetica, Arial, sans-serif;

/* 代码/数字 */
--el-font-family-code: "JetBrains Mono", "Fira Code", "Consolas", monospace;
```

### 3.2 字号规范

| 用途 | 字号 | 行高 | 字重 | 示例 |
|------|------|------|------|------|
| 页面大标题 | 18px | 28px | 600 | 页面标题 |
| 卡片标题 | 16px | 24px | 600 | 卡片标题/表格列名 |
| 正文 | 14px | 22px | 400 | 表单标签/表格内容 |
| 辅助文字 | 12px | 18px | 400 | 提示/时间戳 |
| 数字强调 | 14px | 22px | 600 | 统计数据/OEE数值 |
| 代码/编码 | 13px | 20px | 400 | 物料编码/批次号 |

### 3.3 文字色

```css
--el-text-color-primary: #1f2937;    /* 主要文字 - 深灰黑 */
--el-text-color-regular: #4b5563;    /* 常规文字 - 中灰 */
--el-text-color-secondary: #9ca3af; /* 次要文字 - 浅灰 */
--el-text-color-placeholder: #d1d5db; /* 占位文字 - 更浅 */
```

---

## 4. 间距系统

### 4.1 间距刻度

基于 4px 网格系统：

| Token | 值 | 用途 |
|-------|---|------|
| `space-xs` | 4px | 标签与输入框间距 |
| `space-sm` | 8px | 紧凑元素间距 |
| `space-md` | 16px | 标准元素间距/卡片内边距 |
| `space-lg` | 24px | 区块间距 |
| `space-xl` | 32px | 页面边距 |
| `space-2xl` | 48px | 大区块分隔 |

### 4.2 卡片间距

- 卡片间间距: `16px`
- 卡片内边距: `16px`（紧凑）/ `20px`（标准）
- 搜索区域与表格间距: `16px`
- 表格与分页间距: `16px`

---

## 5. 组件规范

### 5.1 页面基本结构

```vue
<template>
  <div class="page-container">
    <!-- 1. 搜索区域 - 折叠式 -->
    <el-card class="search-card" v-show="!searchCollapsed">
      <el-form :model="searchForm" inline>
        <!-- 搜索字段... -->
      </el-form>
    </el-card>

    <!-- 2. 工具栏 - 批量操作 + 主操作 -->
    <el-card class="toolbar-card">
      <div class="toolbar-left">
        <el-button type="primary">新增</el-button>
        <el-button type="danger" :disabled="!selectedRows.length">批量删除</el-button>
      </div>
      <div class="toolbar-right">
        <el-button text @click="searchCollapsed = !searchCollapsed">
          {{ searchCollapsed ? '展开搜索' : '收起搜索' }}
          <el-icon><ArrowUp v-if="!searchCollapsed" /><ArrowDown v-else /></el-icon>
        </el-button>
      </div>
    </el-card>

    <!-- 3. 表格区域 -->
    <el-card class="table-card">
      <el-table :data="tableData" v-loading="loading">
        <!-- 列定义... -->
      </el-table>
      <div class="pagination">
        <el-pagination v-model:current-page="pagination.page" ... />
      </div>
    </el-card>

    <!-- 4. 弹窗编辑 -->
    <el-dialog v-model="dialogVisible" title="编辑" width="600px">
      <!-- 表单... -->
    </el-dialog>
  </div>
</template>
```

### 5.2 按钮规范

#### 主操作按钮
- 类型: `type="primary"`
- 颜色: `#2563eb`（蓝色）
- 位置: 工具栏最左侧
- 图标: 左侧图标（`Plus`, `Delete`, `Search`）

#### 次操作按钮
- 类型: `type="default"` 或 无类型
- 颜色: 灰色边框
- 位置: 工具栏右侧/表单底部
- 示例: "取消", "重置", "导出"

#### 危险操作按钮
- 类型: `type="danger"`
- 颜色: `#ef4444`（红色）
- 用途: 删除、禁用、取消等不可逆操作
- 需二次确认: `ElMessageBox.confirm`

#### 文字按钮
- 类型: `link` 或 `text`
- 用途: 表格内操作列
- 颜色: 主色/红色

```vue
<!-- 标准工具栏按钮 -->
<el-button type="primary" @click="handleAdd">
  <el-icon><Plus /></el-icon>新增
</el-button>
<el-button type="danger" :disabled="!selectedRows.length" @click="handleBatchDelete">
  <el-icon><Delete /></el-icon>批量删除
</el-button>
<el-button @click="handleExport">
  <el-icon><Download /></el-icon>导出
</el-button>

<!-- 表格操作列 -->
<el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
<el-button link type="danger" @click="handleDelete(row)">删除</el-button>
```

### 5.3 表格规范

#### 标准表格结构
```vue
<el-table :data="tableData" v-loading="loading" stripe border>
  <el-table-column type="selection" width="55" />
  <el-table-column prop="code" label="编码" min-width="120" />
  <el-table-column prop="name" label="名称" min-width="150" show-overflow-tooltip />
  <el-table-column prop="status" label="状态" width="100">
    <template #default="{ row }">
      <el-tag :type="getStatusType(row.status)">{{ row.statusText }}</el-tag>
    </template>
  </el-table-column>
  <el-table-column prop="created_at" label="创建时间" width="180" />
  <el-table-column label="操作" width="180" fixed="right">
    <template #default="{ row }">
      <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
      <el-button link type="danger" @click="handleDelete(row)">删除</el-button>
    </template>
  </el-table-column>
</el-table>
```

#### 状态标签映射
```js
const statusMap = {
  'ACTIVE': { type: 'success', text: '启用' },
  'INACTIVE': { type: 'info', text: '禁用' },
  'PENDING': { type: 'warning', text: '待处理' },
  'RUNNING': { type: 'primary', text: '运行中' },
  'COMPLETED': { type: 'success', text: '已完成' },
  'FAILED': { type: 'danger', text: '失败' },
  'QUALIFIED': { type: 'success', text: '合格' },
  'UNQUALIFIED': { type: 'danger', text: '不合格' },
}
```

### 5.4 表单规范

#### 弹窗表单
```vue
<el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
  <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
    <el-row :gutter="20">
      <el-col :span="12">
        <el-form-item label="编码" prop="code">
          <el-input v-model="formData.code" :disabled="isEdit" />
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item label="名称" prop="name">
          <el-input v-model="formData.name" />
        </el-form-item>
      </el-col>
    </el-row>
    <el-form-item label="描述" prop="description">
      <el-input v-model="formData.description" type="textarea" :rows="3" />
    </el-form-item>
  </el-form>
  <template #footer>
    <el-button @click="dialogVisible = false">取消</el-button>
    <el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button>
  </template>
</el-dialog>
```

#### 表单校验规则
```js
const rules: FormRules = {
  code: [{ required: true, message: '请输入编码', trigger: 'blur' }],
  name: [{ required: true, message: '请输入名称', trigger: 'blur' }],
  email: [{ type: 'email', message: '请输入正确的邮箱', trigger: 'blur' }],
  phone: [{ pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号', trigger: 'blur' }],
  quantity: [{ type: 'number', min: 1, message: '数量必须大于0', trigger: 'blur' }],
}
```

### 5.5 搜索表单

```vue
<el-card class="search-card">
  <el-form :model="searchForm" inline>
    <el-form-item label="编码">
      <el-input v-model="searchForm.code" placeholder="请输入编码" clearable />
    </el-form-item>
    <el-form-item label="状态">
      <el-select v-model="searchForm.status" placeholder="请选择" clearable>
        <el-option label="启用" :value="1" />
        <el-option label="禁用" :value="0" />
      </el-select>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="handleSearch">
        <el-icon><Search /></el-icon>查询
      </el-button>
      <el-button @click="handleReset">
        <el-icon><Refresh /></el-icon>重置
      </el-button>
    </el-form-item>
  </el-form>
</el-card>
```

### 5.6 分页规范

```vue
<div class="pagination">
  <el-pagination
    v-model:current-page="pagination.page"
    v-model:page-size="pagination.pageSize"
    :total="pagination.total"
    :page-sizes="[10, 20, 50, 100]"
    layout="total, sizes, prev, pager, next, jumper"
    @size-change="loadData"
    @current-change="loadData"
  />
</div>
```

### 5.7 卡片规范

```vue
<!-- 搜索卡片 -->
<el-card class="search-card">
  <!-- 默认无边框，内部padding 16px -->
</el-card>

<!-- 工具栏卡片 -->
<el-card class="toolbar-card">
  <div class="toolbar-content">
    <!-- 左对齐: 主要操作按钮 -->
    <!-- 右对齐: 辅助操作按钮 -->
  </div>
</el-card>

<!-- 表格卡片 -->
<el-card class="table-card">
  <!-- 无额外样式 -->
</el-card>
```

```scss
// 卡片样式
.search-card,
.toolbar-card {
  margin-bottom: 16px;
}

.toolbar-card :deep(.el-card__body) {
  padding: 12px 16px;
}

.pagination {
  margin-top: 16px;
  display: flex;
  justify-content: flex-end;
}
```

---

## 6. 页面布局规范

### 6.1 页面容器

```vue
<template>
  <div class="page-container">
    <!-- 页面标题 -->
    <div class="page-header">
      <h2 class="page-title">{{ pageTitle }}</h2>
    </div>

    <!-- 搜索区域 -->
    <el-card class="search-card">...</el-card>

    <!-- 工具栏 -->
    <el-card class="toolbar-card">...</el-card>

    <!-- 表格 -->
    <el-card class="table-card">...</el-card>
  </div>
</template>

<style scoped>
.page-container {
  padding: 16px;
}

.page-header {
  margin-bottom: 16px;
}

.page-title {
  font-size: 18px;
  font-weight: 600;
  color: var(--el-text-color-primary);
  margin: 0;
}
</style>
```

### 6.2 栅格布局

| 场景 | 栅格 | 说明 |
|------|------|------|
| 表单（2列） | `el-col :span="12"` | 2列布局 |
| 表单（3列） | `el-col :span="8"` | 3列布局 |
| 统计卡片 | `el-col :span="6"` | 4列布局 |
| 大屏看板 | `el-col :span="8"` | 3列布局 |

### 6.3 Andon大屏布局

```vue
<!-- Andon车间大屏 -->
<div class="andon-screen">
  <!-- 顶部标题栏 -->
  <div class="screen-header">
    <h1>安灯监控系统 - {{ workshopName }}</h1>
    <div class="current-time">{{ currentTime }}</div>
  </div>

  <!-- 中间状态区 - 全屏色块 -->
  <div class="status-zone">
    <div
      v-for="zone in zones"
      :key="zone.id"
      class="zone-card"
      :class="zone.status"
    >
      <div class="zone-name">{{ zone.name }}</div>
      <div class="zone-status-text">{{ zone.statusText }}</div>
      <div class="zone-count">{{ zone.callCount }}</div>
    </div>
  </div>

  <!-- 底部统计区 -->
  <div class="screen-footer">
    <div class="stat-card">今日呼叫: {{ todayCalls }}</div>
    <div class="stat-card">平均响应: {{ avgResponse }}分钟</div>
    <div class="stat-card">解决率: {{ resolutionRate }}%</div>
  </div>
</div>
```

### 6.4 甘特图布局

```vue
<!-- APS甘特图页面 -->
<div class="gantt-page">
  <!-- 左侧: 工作中心列表 -->
  <div class="gantt-sidebar">
    <div class="workcenter-list">
      <div v-for="wc in workCenters" :key="wc.id" class="wc-item">
        <span class="wc-name">{{ wc.name }}</span>
        <span class="wc-utilization">{{ wc.utilization }}%</span>
      </div>
    </div>
  </div>

  <!-- 右侧: 甘特图时间轴 -->
  <div class="gantt-timeline">
    <!-- 时间刻度 -->
    <div class="timeline-header">
      <div
        v-for="hour in timeRange"
        :key="hour"
        class="hour-cell"
        :style="{ width: cellWidth + 'px' }"
      >
        {{ hour }}:00
      </div>
    </div>

    <!-- 任务条 -->
    <div class="timeline-body">
      <div
        v-for="task in tasks"
        :key="task.id"
        class="task-bar"
        :class="task.productFamily"
        :style="{
          top: task.rowIndex * rowHeight + 'px',
          left: task.startOffset + 'px',
          width: task.duration + 'px'
        }"
        @click="handleTaskClick(task)"
      >
        {{ task.productName }}
      </div>
    </div>
  </div>
</div>
```

---

## 7. 状态设计

### 7.1 加载状态

```vue
<!-- 全页加载 -->
<el-table v-loading="loading" ...>
<!-- 或全页 -->
<div v-loading="fullLoading" class="full-loading">
  <el-skeleton :rows="10" animated />
</div>
```

### 7.2 空状态

```vue
<el-empty
  v-if="tableData.length === 0 && !loading"
  :description="emptyDescription"
>
  <el-button type="primary" @click="handleAdd">新增数据</el-button>
</el-empty>
```

**空状态文案规范**:
- 表格空: "暂无数据"
- 搜索无结果: "未找到匹配的结果，请调整搜索条件"
- 收藏为空: "暂无收藏记录"

### 7.3 错误状态

```vue
<el-result
  v-if="error"
  icon="error"
  title="加载失败"
  :sub-title="errorMessage"
>
  <template #extra>
    <el-button type="primary" @click="loadData">重试</el-button>
  </template>
</el-result>
```

### 7.4 操作反馈

```js
// 成功
ElMessage.success('操作成功')

// 警告
ElMessage.warning('数据已保存，但存在以下问题...')

// 错误
ElMessage.error('操作失败，请重试')

// 确认
await ElMessageBox.confirm(
  '此操作将永久删除该数据，是否继续？',
  '删除确认',
  { type: 'warning' }
)
```

---

## 8. 模块UI规范索引

### 8.1 Andon系统页面

| 页面 | 路径 | 布局类型 | 特殊组件 |
|------|------|---------|---------|
| Andon呼叫 | `/andon/call` | 大屏+工位端 | 全屏色块/灯色/倒计时 |
| 响应处理 | `/andon/response` | 列表+详情 | 计时器/处理记录 |
| 升级规则 | `/andon/rule` | 标准列表 | 规则配置表单 |
| 统计分析 | `/andon/stats` | 看板+图表 | 柱状图/饼图/趋势图 |

### 8.2 APS排程页面

| 页面 | 路径 | 布局类型 | 特殊组件 |
|------|------|---------|---------|
| 甘特图 | `/aps/gantt` | 双栏 | dhtmlx-gantt/拖拽 |
| 排程计划 | `/aps/plan` | 标准列表 | 状态标签 |
| 工作中心 | `/aps/work-center` | 标准列表+树 | 产能利用率条 |
| 换型矩阵 | `/aps/changeover` | 矩阵表格 | 颜色热力图 |

### 8.3 质量管理页面

| 页面 | 路径 | 布局类型 | 特殊组件 |
|------|------|---------|---------|
| IQC来料检验 | `/quality/iqc` | 标准列表+详情 | 蓝牙量具控件 |
| IPQC过程检验 | `/quality/ipqc` | 标准列表+详情 | 检验项目勾选 |
| SPC控制图 | `/quality/spc-chart` | 图表 | ECharts控制图 |
| NCR处理 | `/quality/ncr` | 标准列表+流程 | 流程步骤条 |

### 8.4 设备管理页面

| 页面 | 路径 | 布局类型 | 特殊组件 |
|------|------|---------|---------|
| OEE分析 | `/equipment/oee` | 仪表盘 | 环形图/趋势图 |
| TEEP分析 | `/equipment/teep` | 多指标 | 堆积面积图 |
| 模具管理 | `/equipment/mold` | 标准列表 | 寿命进度条 |
| 设备点检 | `/equipment/check` | 卡片+日历 | 日历控件 |

---

## 附录A: 组件开发检查清单

每个Vue页面开发前需确认：

- [ ] 搜索表单使用 `inline` 模式
- [ ] 工具栏有收起/展开搜索功能
- [ ] 表格有 `v-loading` 和 `stripe` 属性
- [ ] 状态使用 `el-tag` + 语义化颜色
- [ ] 操作列使用 `link` 类型按钮
- [ ] 分页在右下角对齐
- [ ] 弹窗标题区分新增/编辑
- [ ] 表单有 `label-width`
- [ ] 删除操作有二次确认
- [ ] API调用有 `try-catch` 和错误提示
- [ ] 按钮有权限控制 `v-if="hasPermission(...)"`

## 附录B: 命名规范

| 类型 | 规范 | 示例 |
|------|------|------|
| Vue文件 | PascalCase | `UserList.vue`, `OEEAnalysis.vue` |
| API函数 | camelCase | `getUserList`, `createOrder` |
| 路由路径 | kebab-case | `/system/user`, `/aps/gantt` |
| CSS类 | kebab-case | `.page-container`, `.search-card` |
| 状态值 | SCREAMING_SNAKE | `ACTIVE`, `PENDING`, `COMPLETED` |
