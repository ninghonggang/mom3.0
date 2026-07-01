# MOM 3.0 UI 设计规范

> 版本：V1.0 | 最后更新：2026-07-01 | 维护人：前端架构
> 适用范围：MOM 3.0 全部前端页面（`mom-web/src/views/`）
> 基础规范：Element Plus 2.5 设计语言、Element Plus Icons Vue 2.3

---

## 1. 整体布局

### 1.1 三栏式布局

```
┌─────────────────────────────────────────────────────────────┐
│  顶栏（高 56px）：Logo  │ 系统名  │ 租户切换  │ 全屏  │ 通知 │ 用户│
├──────────┬──────────────────────────────────────────────────┤
│          │  面包屑导航                                       │
│  侧边菜单  ├──────────────────────────────────────────────────┤
│  (240px)  │                                                  │
│          │            主体内容区                             │
│  - 首页   │            (自适应高度)                            │
│  - 生产   │                                                  │
│  - 质量   │                                                  │
│  ...      │                                                  │
│          │                                                  │
├──────────┴──────────────────────────────────────────────────┤
│  底栏（高 32px）：版权  │ 版本号  │ 技术支持                  │
└─────────────────────────────────────────────────────────────┘
```

- **侧边菜单**：可折叠（240px ↔ 64px），菜单数据来自 `sys_menu` 表
- **顶栏**：租户切换器仅多租户版本显示
- **主体内容**：自适应高度，内部再分「查询区 + 工具栏 + 表格 + 分页」

### 1.2 屏幕适配

| 设备 | 宽度 | 处理 |
|------|------|------|
| 桌面（默认）| ≥ 1280px | 完整三栏布局 |
| 笔记本 | 1024-1279px | 侧边菜单默认折叠 |
| 平板（横向）| 768-1023px | 侧边菜单默认折叠，表格列允许横向滚动 |
| 平板（竖向）| 600-767px | 侧边菜单隐藏（顶栏汉堡菜单弹出），表格降密度 |
| 手机 | < 600px | 移动端 / PDA 页面（`mom-mobile` 项目，不在本系统）|

> 工业现场平板常见（768-1280px），必须保证此区间可用

---

## 2. 列表页模板

### 2.1 标准结构

```vue
<template>
  <div class="page-container">
    <!-- ① 查询区 -->
    <div class="search-area">
      <el-form :inline="true" :model="query" class="search-form">
        <el-form-item label="工单号">
          <el-input v-model="query.orderNo" placeholder="请输入" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="query.status" placeholder="全部" clearable>
            <el-option v-for="o in STATUS_OPTIONS" :key="o.value" :label="o.label" :value="o.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="日期">
          <el-date-picker v-model="query.dateRange" type="daterange" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :icon="Search" @click="handleSearch">查询</el-button>
          <el-button :icon="Refresh" @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- ② 工具栏 -->
    <div class="toolbar">
      <el-button type="primary" :icon="Plus" v-permission="['production:order:create']" @click="handleCreate">新增</el-button>
      <el-button type="success" :icon="Edit" :disabled="!selection.length" @click="handleBatchEdit">批量编辑</el-button>
      <el-button type="danger" :icon="Delete" :disabled="!selection.length" v-permission="['production:order:delete']" @click="handleBatchDelete">批量删除</el-button>
      <el-button :icon="Download" @click="handleExport">导出</el-button>
      <el-button :icon="Upload" @click="handleImport">导入</el-button>
      <div class="toolbar-right">
        <el-tooltip content="刷新"><el-button :icon="Refresh" circle @click="loadData" /></el-tooltip>
        <el-tooltip content="列设置"><el-button :icon="Setting" circle @click="columnSettingVisible = true" /></el-tooltip>
      </div>
    </div>

    <!-- ③ 表格 -->
    <el-table
      v-loading="loading"
      :data="list"
      border
      stripe
      @selection-change="val => selection = val"
    >
      <el-table-column type="selection" width="44" />
      <el-table-column prop="orderNo" label="工单号" min-width="160" fixed="left" />
      <el-table-column prop="productName" label="产品" min-width="200" show-overflow-tooltip />
      <el-table-column prop="quantity" label="数量" width="100" align="right" />
      <el-table-column prop="status" label="状态" width="100">
        <template #default="{ row }">
          <el-tag :type="STATUS_TAG_MAP[row.status]?.type" disable-transitions>
            {{ STATUS_TAG_MAP[row.status]?.text }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="startTime" label="开始时间" width="160" />
      <el-table-column prop="endTime" label="结束时间" width="160" />
      <el-table-column label="操作" width="220" fixed="right">
        <template #default="{ row }">
          <el-button type="primary" link @click="handleView(row)">查看</el-button>
          <el-button type="primary" link v-permission="['production:order:update']" @click="handleEdit(row)">编辑</el-button>
          <el-button type="primary" link v-permission="['production:order:release']" @click="handleRelease(row)" v-if="row.status === 'DRAFT'">下达</el-button>
          <el-button type="danger" link v-permission="['production:order:delete']" @click="handleDelete(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- ④ 分页 -->
    <div class="pagination">
      <el-pagination
        v-model:current-page="query.page"
        v-model:page-size="query.pageSize"
        :total="total"
        :page-sizes="[20, 50, 100, 200]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="loadData"
        @current-change="loadData"
      />
    </div>
  </div>
</template>
```

### 2.2 关键约定

| 项 | 规范 |
|----|------|
| **查询区** | inline 表单，每行 3-4 个字段；过多字段折叠到「高级搜索」 |
| **工具栏** | 主操作按钮在左、刷新/列设置在右；批量操作 :disabled="!selection.length" |
| **表格** | border + stripe；首列 selection，最后列操作（fixed="right"）|
| **状态列** | 必须用 `el-tag`，禁用 transitions（避免动画卡顿）|
| **数字列** | `align="right"` |
| **长文本** | `show-overflow-tooltip` |
| **操作列** | 文字链接样式（`type="primary" link`），最多 4 个；多于 4 个折叠到「更多」|
| **分页** | 默认 20 条/页，可选 50/100/200 |

---

## 3. 表单弹窗模板

### 3.1 新增 / 编辑弹窗

```vue
<el-dialog
  v-model="dialogVisible"
  :title="dialogMode === 'create' ? '新增工单' : '编辑工单'"
  width="720px"
  :close-on-click-modal="false"
  @closed="handleDialogClosed"
>
  <el-form ref="formRef" :model="form" :rules="rules" label-width="100px">
    <el-row :gutter="20">
      <el-col :span="12">
        <el-form-item label="工单号" prop="orderNo">
          <el-input v-model="form.orderNo" placeholder="自动生成或手动输入" />
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item label="产品" prop="productId">
          <el-select v-model="form.productId" filterable>
            <el-option v-for="m in materials" :key="m.id" :label="`${m.code} / ${m.name}`" :value="m.id" />
          </el-select>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="12">
        <el-form-item label="数量" prop="quantity">
          <el-input-number v-model="form.quantity" :min="1" :precision="0" />
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item label="计划开始" prop="startTime">
          <el-date-picker v-model="form.startTime" type="datetime" />
        </el-form-item>
      </el-col>
    </el-row>
    <el-form-item label="备注" prop="remark">
      <el-input v-model="form.remark" type="textarea" :rows="3" maxlength="500" show-word-limit />
    </el-form-item>
  </el-form>

  <template #footer>
    <el-button @click="dialogVisible = false">取消</el-button>
    <el-button type="primary" :loading="submitting" @click="handleSubmit">保存</el-button>
  </template>
</el-dialog>
```

### 3.2 关键约定

| 项 | 规范 |
|----|------|
| **宽度** | 简单表单 600px；中等 720px；复杂 960px；含明细 80vw |
| **栅格** | `:gutter="20"`，一行两列（`:span="12"`），3 列（`:span="8"`）|
| **label-width** | 固定 100px 或 120px，统一 |
| **校验** | 必填项必须有 `rules` 规则；失焦触发 |
| **提交按钮** | loading 态；防止重复提交 |
| **关闭行为** | `:close-on-click-modal="false"`，强制走「取消」|
| **明细/子表** | 用嵌套 `el-table`，底部有「添加行」按钮 |

---

## 4. 详情弹窗模板

### 4.1 标准结构

```vue
<el-dialog v-model="detailVisible" :title="`工单详情 - ${detail.orderNo}`" width="960px">
  <el-descriptions :column="3" border>
    <el-descriptions-item label="工单号">{{ detail.orderNo }}</el-descriptions-item>
    <el-descriptions-item label="状态">
      <el-tag :type="STATUS_TAG_MAP[detail.status]?.type">{{ STATUS_TAG_MAP[detail.status]?.text }}</el-tag>
    </el-descriptions-item>
    <el-descriptions-item label="创建时间">{{ formatDate(detail.createdAt) }}</el-descriptions-item>
    <el-descriptions-item label="产品" :span="3">{{ detail.productName }}</el-descriptions-item>
    <el-descriptions-item label="数量">{{ detail.quantity }}</el-descriptions-item>
    <el-descriptions-item label="已报工">{{ detail.reportedQty }}</el-descriptions-item>
    <el-descriptions-item label="不良数">{{ detail.defectQty }}</el-descriptions-item>
    <el-descriptions-item label="计划开始">{{ formatDate(detail.startTime) }}</el-descriptions-item>
    <el-descriptions-item label="计划结束">{{ formatDate(detail.endTime) }}</el-descriptions-item>
    <el-descriptions-item label="实际开始">{{ formatDate(detail.actualStart) }}</el-descriptions-item>
    <el-descriptions-item label="备注" :span="3">{{ detail.remark || '-' }}</el-descriptions-item>
  </el-descriptions>

  <el-tabs class="detail-tabs">
    <el-tab-pane label="工序明细">
      <el-table :data="detail.processes" border size="small">
        ...
      </el-table>
    </el-tab-pane>
    <el-tab-pane label="报工记录">
      <el-table :data="detail.reports" border size="small">
        ...
      </el-table>
    </el-tab-pane>
    <el-tab-pane label="变更记录">
      <el-timeline>
        <el-timeline-item v-for="log in detail.changeLogs" :key="log.id" :timestamp="log.time">
          {{ log.content }}
        </el-timeline-item>
      </el-timeline>
    </el-tab-pane>
  </el-tabs>

  <template #footer>
    <el-button @click="detailVisible = false">关闭</el-button>
    <el-button type="primary" v-if="detail.status === 'DRAFT'" @click="handleRelease(detail)">下达</el-button>
    <el-button type="danger" v-if="['DRAFT', 'PENDING'].includes(detail.status)" @click="handleCancel(detail)">取消</el-button>
  </template>
</el-dialog>
```

### 4.2 关键约定

- **基本信息**：`el-descriptions` 3 列展示
- **明细/历史**：用 `el-tabs` 切换
- **变更记录**：用 `el-timeline`
- **操作按钮**：根据 `status` 动态显示（见下方状态色规范）

---

## 5. 状态色规范

### 5.1 通用语义

| 业务语义 | Element Plus type | 使用场景 |
|---------|------------------|---------|
| **成功 / 已完成 / 正常** | `success`（绿） | 完成、通过、有效、合格、正常 |
| **警告 / 待处理 / 即将到期** | `warning`（黄） | 待审核、即将到期、中风险 |
| **危险 / 失败 / 不良** | `danger`（红） | 不合格、已停用、超期、严重延期 |
| **信息 / 草稿 / 已取消** | `info`（灰） | 草稿、已取消、归档 |
| **进行中 / 主要操作** | `primary`（蓝） | 审批中、已派工、运行中 |

### 5.2 业务状态映射（标准字典）

**生产工单状态**：

| 状态值 | 显示文本 | type |
|--------|---------|------|
| `DRAFT` | 草稿 | info |
| `PENDING` | 待审批 | warning |
| `APPROVED` | 已审批 | primary |
| `RELEASED` | 已下达 | primary |
| `IN_PROGRESS` | 执行中 | primary |
| `HOLD` | 挂起 | warning |
| `COMPLETED` | 已完成 | success |
| `CLOSED` | 已关闭 | info |
| `CANCELLED` | 已取消 | info |

**质量检验单状态**：

| 状态值 | 显示文本 | type |
|--------|---------|------|
| `PENDING` | 待检验 | warning |
| `INSPECTING` | 检验中 | primary |
| `PASSED` | 合格 | success |
| `FAILED` | 不合格 | danger |
| `ACCEPTED` | 已接收 | success |
| `REJECTED` | 已拒收 | danger |
| `WAIVED` | 让步接收 | warning |

**NCR 状态**：

| 状态值 | 显示文本 | type |
|--------|---------|------|
| `OPEN` | 待处理 | warning |
| `INVESTIGATING` | 处理中 | primary |
| `RESOLVED` | 已解决 | success |
| `CLOSED` | 已关闭 | info |
| `CANCELLED` | 已取消 | info |

**设备状态**：

| 状态值 | 显示文本 | type |
|--------|---------|------|
| `IDLE` | 空闲 | info |
| `RUNNING` | 运行中 | primary |
| `DOWN` | 停机 | danger |
| `MAINTENANCE` | 保养中 | warning |
| `REPAIR` | 维修中 | warning |
| `RETIRED` | 已报废 | info |

完整状态映射规范见 [`MOM3.0_附录.md`](./MOM3.0_附录.md)（阶段 2 补充）

---

## 6. 按钮规范

### 6.1 类型与位置

| 类型 | type | 位置 | 场景 |
|------|------|------|------|
| **主操作** | `primary`（蓝） | 工具栏首位 | 新增、保存、下达 |
| **次操作** | `default`（白） | 主操作之后 | 编辑、详情、刷新 |
| **危险** | `danger`（红） | 删除/取消 | 删除、强制结束 |
| **成功** | `success`（绿） | 流转类 | 审批通过、提交 |
| **链接** | `primary` `link` | 操作列 | 查看、编辑、删除 |
| **文字按钮** | `text` | 行内次要 | 备注、提示 |

### 6.2 禁用与确认

```vue
<!-- 危险操作必须二次确认 -->
<el-popconfirm title="确认删除该工单？" @confirm="handleDelete(row)">
  <template #reference>
    <el-button type="danger" link>删除</el-button>
  </template>
</el-popconfirm>

<!-- 重要操作必须带 loading -->
<el-button type="primary" :loading="submitting" @click="handleSubmit">保存</el-button>
```

---

## 7. 空状态 / 加载 / 错误展示

### 7.1 空状态

```vue
<el-empty
  v-if="!loading && list.length === 0"
  description="暂无数据"
  :image-size="120"
>
  <el-button type="primary" @click="handleCreate">新增第一条</el-button>
</el-empty>
```

- **空状态**：必须有图标 + 文案 + 操作建议（不要空白）
- **第一次使用**：推荐一个「新增」CTA
- **已筛过无结果**：提示「没有符合条件的记录，[清空筛选]」

### 7.2 加载状态

| 场景 | 处理 |
|------|------|
| **首屏加载** | 表格 `v-loading="loading"`，骨架屏可选 |
| **按钮提交** | 按钮 `:loading="submitting"`，禁用防重 |
| **页面级** | `<el-skeleton :rows="5" animated />` |
| **下拉/分页** | 表格 inline loading，不阻塞 |

### 7.3 错误展示

```vue
<el-result icon="warning" title="加载失败" sub-title="网络异常，请稍后重试">
  <template #extra>
    <el-button type="primary" @click="loadData">重试</el-button>
  </template>
</el-result>
```

- **统一错误码提示**：响应拦截器处理，弹 `ElMessage.error`
- **可重试错误**：页面级显示 `<el-result>` + 重试按钮
- **业务错误**：从响应 `data.message` 提取

---

## 8. 移动端 / PDA 适配

### 8.1 适配范围

- **车间现场**：扫描 PDA（Android，720×1280 或 1080×1920）
- **移动审批**：手机端（iOS/Android，375-414px 宽）

### 8.2 适配要点

| 项 | 规范 |
|----|------|
| **触控目标** | 最小 44×44px |
| **按钮高度** | 48px（比桌面 32px 更大）|
| **字号** | 基础 16px（避免 iOS 自动放大）|
| **间距** | 8/16/24px 三档（不放 4px）|
| **手势** | 左滑操作（删除/确认），下拉刷新 |
| **扫码** | 调用 `@zxing/browser` 或原生接口 |

### 8.3 移动端模板（MOM 3.0 中 PDA 复用 Web 页面，单独移动 App 在 P3 阶段）

---

## 9. 国际化（i18n）

### 9.1 规范

- **当前状态**：MOM 3.0 **尚未启用多语言**（仅简体中文），但代码结构和命名应预留 i18n 扩展空间
- **未来语言**：英文（en）、繁体中文（zh-TW）、日语（ja）

### 9.2 文案规范

| 项 | 规范 |
|----|------|
| **避免缩写** | 「工单号」不要写成「单号」 |
| **统一术语** | 用「生产工单」不要混用「生产单」「工单」 |
| **动宾结构** | 按钮用「新增工单」「下达计划」 |
| **标点** | 中文文案用全角（，。：），按钮不要尾标点 |
| **占位符** | 提示文案用 `请输入`「下拉」用 `请选择`「日期」用 `请选择日期范围` |
| **复数** | `${count} 条记录已删除`，避免「记录删除成功 1」 |

### 9.3 i18n key 命名（未来扩展时）

```
<模块>.<页面>.<元素>
production.orderList.search.orderNo.label
production.orderList.search.orderNo.placeholder
production.orderList.button.create
production.orderList.status.RELEASED
```

未来启用时使用 `vue-i18n`：

```ts
// i18n/index.ts
import { createI18n } from 'vue-i18n'
import zhCN from './locales/zh-CN.json'
import enUS from './locales/en-US.json'

export default createI18n({
  locale: 'zh-CN',
  fallbackLocale: 'en-US',
  messages: { 'zh-CN': zhCN, 'en-US': enUS }
})
```

---

## 10. 图标规范

### 10.1 来源

- **Element Plus Icons Vue**（`@element-plus/icons-vue` 2.3.1）— 80% 场景
- **自定义 SVG**（`/src/assets/icons/`）— 业务专用图标

### 10.2 一级菜单图标（菜单 seed 已配置）

| 菜单 | 图标 |
|------|------|
| 首页 | `HomeFilled` |
| 生产执行 | `List` |
| 质量管理 | `CircleCheck` |
| 计划排程 | `Calendar` |
| 设备管理 | `Monitor` |
| 仓储物流 | `House` |
| 供应链 | `Connection` |
| 主数据 | `Box` |
| 追溯中心 | `Search` |
| 运营分析 | `DataLine` |
| 系统管理 | `Setting` |

### 10.3 子菜单常用图标

| 功能 | 图标 |
|------|------|
| 新增 | `Plus` |
| 编辑 | `Edit` |
| 删除 | `Delete` |
| 查询 | `Search` |
| 重置 | `Refresh` |
| 导出 | `Download` |
| 导入 | `Upload` |
| 详情 | `View` |
| 审批 | `Check` |
| 下达/发布 | `Promotion` |
| 完成 | `Finished` |
| 取消 | `CircleClose` |

---

## 11. 性能与体验

### 11.1 性能基线

| 指标 | 目标 |
|------|------|
| **首屏渲染** | ≤ 1.5s |
| **列表查询** | P95 ≤ 1s |
| **详情查询** | P95 ≤ 0.5s |
| **表单保存** | P95 ≤ 1.5s |
| **批量导入 1000 行** | ≤ 5s |

### 11.2 体验规范

- **防抖**：搜索框 input 节流 300ms
- **取消竞态**：请求竞态保护（同一接口未返回不发新请求）
- **预加载**：菜单 hover 时预加载常用页面
- **缓存**：字典数据用 Pinia 缓存（不重复请求）
- **乐观更新**：状态切换类操作可考虑乐观更新（如暂停/启用）
- **键盘可达**：Tab 顺序合理，表单 Enter 提交
- **无障碍**：颜色对比 ≥ 4.5:1，提供 icon + text（不要纯 icon）

---

## 12. 代码规范（前端）

### 12.1 文件结构

```
mom-web/src/views/<module>/<Page>.vue
├── <Page>.vue                 # 主组件
├── components/                # 子组件（如有）
│   └── <SubComponent>.vue
├── composables/               # 组合式函数（如有）
│   └── use<Xxx>.ts
├── types.ts                   # 类型定义（如有）
└── index.ts                   # 导出（如需）
```

### 12.2 命名

| 项 | 命名 |
|----|------|
| 组件文件 | PascalCase：`ProductionOrderList.vue` |
| 组合函数 | camelCase `useXxx`：`useOrderForm.ts` |
| 类型/接口 | PascalCase：`ProductionOrder` |
| 常量 | UPPER_SNAKE：`STATUS_OPTIONS` |
| API 函数 | camelCase：`getProductionOrders` |
| 状态字段 | UPPER_SNAKE：`DRAFT` / `IN_PROGRESS` |

### 12.3 Pinia Store 规范

```ts
// stores/order.ts
import { defineStore } from 'pinia'

export const useOrderStore = defineStore('order', {
  state: () => ({
    orders: [] as ProductionOrder[],
    current: null as ProductionOrder | null
  }),
  getters: {
    orderMap: (state) => new Map(state.orders.map(o => [o.id, o]))
  },
  actions: {
    async fetchOrders(query: OrderQuery) { ... }
  }
})
```

---

## 13. 附录

### 13.1 表格列宽规范

| 列宽 | 适用 |
|------|------|
| `width="80"` | 序号、单选 |
| `width="100"` | 数量、状态 |
| `width="120"` | 短编码、日期 |
| `width="160"` | 长编码、名称 |
| `width="200"` | 长名称 |
| `min-width="160"` | 内容可能变长的列（推荐）|

### 13.2 必查的全局组件

| 组件 | 路径 |
|------|------|
| `MainLayout` | `@/components/layout/MainLayout.vue` |
| `SearchForm` | `@/components/form/SearchForm.vue` |
| `PageContainer` | `@/components/layout/PageContainer.vue` |

### 13.3 设计资源

- Element Plus 设计资源：<https://element-plus.org/zh-CN/guide/design.html>
- Ant Design 设计语言（参考）：<https://ant.design/docs/spec/introduce-cn>
- Arco Design（参考）：<https://arco.design/docs/spec/introduce>

---

## 14. 修订记录

| 版本 | 日期 | 修订人 | 说明 |
|------|------|--------|------|
| V1.0 | 2026-07-01 | 前端架构 | 初版；建立 13 节规范；解决各模块文档「同 X 模块标准布局」= 0 内容的循环引用问题 |

---

**下一步**：阶段 2 将建立 13 章节模块设计模板 + 5 个组件级规范（表格/表单/详情/树/上传）；详见 [`MOM3.0-design-doc-improvement-2026-07-01.md`](./research/MOM3.0-design-doc-improvement-2026-07-01.md)。