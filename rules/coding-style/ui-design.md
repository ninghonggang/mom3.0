# UI 设计规范

> 本规范定义 MOM3.0 系统的界面设计标准，确保用户体验一致性。

## 1. 设计原则

### 1.1 核心原则

| 原则 | 说明 | 优先级 |
|------|------|--------|
| **一致性** | 整个系统保持统一的视觉风格和交互模式 | CRITICAL |
| **效率** | 用户能快速完成任务，减少操作步骤 | CRITICAL |
| **清晰** | 信息层次分明，重点突出 | HIGH |
| **反馈** | 所有操作都有明确的系统反馈 | HIGH |

### 1.2 设计目标

- **专业感**: 工业级系统的稳重与可信
- **高效感**: 减少操作路径，快速完成任务
- **现代感**: 符合当下主流 UI 趋势
- **MES特色**: 制造业场景的沉浸式体验

## 2. 色彩系统

### 2.1 主色调

```
主色 (Primary):     #409EFF  (Element Plus 蓝 - 工业感)
主色悬停:           #66B1FF
主色点击:           #3A8EE6
```

**使用场景**: 主要按钮、链接、选中状态、图标

### 2.2 功能色

```
成功 (Success):      #67C23A  (绿色 - 表示正常、完成)
警告 (Warning):      #E6A23C  (橙色 - 表示提醒、待处理)
危险 (Danger):      #F56C6C  (红色 - 表示错误、删除)
信息 (Info):        #909399  (灰色 - 表示中性信息)
```

### 2.3 中性色

```
主要文字:           #303133  (标题、重要内容)
常规文字:          #606266  (正文、描述)
次要文字:          #909399  (辅助信息、占位符)
占位符:            #C0C4CC  (输入框占位符)
边框色:            #DCDFE6  (分割线、边框)
分割线:            #E4E7ED  (表格分割、列表分割)
背景色:            #F5F7FA  (页面背景、卡片背景)
深色背景:          #FAFAFA  (表格斑马纹)
```

### 2.4 状态色 (制造业场景)

```
设备运行:           #67C23A  (绿色)
设备待机:           #E6A23C  (橙色)
设备故障:           #F56C6C  (红色)
设备维修:           #909399  (灰色)

工单进行中:         #409EFF  (蓝色)
工单已完成:         #67C23A  (绿色)
工单已延误:         #F56C6C  (红色)
工单待开始:         #E6A23C  (橙色)
```

## 3. 字体系统

### 3.1 字体家族

```
主字体:             'Helvetica Neue', 'Helvetica', 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', sans-serif
等宽字体:          'Menlo', 'Monaco', 'Consolas', 'monospace'  (用于代码、数字)
```

### 3.2 字号规范

```
页面标题:          24px / 1.5    (h1)
卡片标题:          18px / 1.5    (h2)
区块标题:          16px / 1.5    (h3)
正文:              14px / 1.5    (body)
辅助文字:          12px / 1.5    (small)
数字显示:          16px / 1.2    (强调数字)
```

### 3.3 字重规范

```
标题:              500 或 600    (Bold)
正文:              400          (Regular)
辅助文字:          400          (Regular)
按钮文字:          500          (Medium)
```

## 4. 间距系统

### 4.1 基础间距

```
xs: 4px    (组件内部紧密元素)
sm: 8px    (组件内部松散元素)
md: 16px   (组件之间、卡片内边距)
lg: 24px   (区块之间)
xl: 32px   (页面主要区块)
xxl: 48px  (页面大幅分割)
```

### 4.2 组件间距

```
表单项目间距:      16px
表格列间距:        默认
按钮组间距:        8px
卡片内边距:        16px 或 20px
```

## 5. 组件规范

### 5.1 按钮

```
主要按钮:          bg: #409EFF, color: #FFF
次要按钮:          bg: #FFF, border: #409EFF, color: #409EFF
文字按钮:          color: #409EFF, 无边框
危险按钮:          bg: #F56C6C, color: #FFF

高度:              32px (默认), 28px (小), 40px (大)
圆角:              4px (工业感，不宜过圆)
间距:              按钮之间 8px
```

### 5.2 输入框

```
高度:              32px
边框:              1px solid #DCDFE6
聚焦边框:          1px solid #409EFF
圆角:              4px
内边距:            0 12px
```

### 5.3 表格

```
表头背景:          #F5F7FA
斑马纹:            #FAFAFA
行高:              44px (数据行)
表头高:            44px
边框:              1px solid #E4E7ED
```

### 5.4 卡片

```
背景:              #FFF
圆角:              4px
阴影:              0 2px 12px 0 rgba(0, 0, 0, 0.1)
内边距:            16px 或 20px
```

### 5.5 对话框 (Dialog)

```
宽度:              500px (小), 700px (中), 900px (大)
圆角:              8px
遮罩:              rgba(0, 0, 0, 0.5)
```

## 6. 布局规范

### 6.1 页面布局

```
顶部导航:          60px 高度
侧边栏:            200px (可折叠至 64px)
内容区:            自适应，右侧可留 16px 边距
```

### 6.2 列表页布局

```
┌─────────────────────────────────────────┐
│  搜索区域 (Search Area)     16px padding │
├─────────────────────────────────────────┤
│  工具栏 (Toolbar)           16px padding │
├─────────────────────────────────────────┤
│                                         │
│  表格区域 (Table)           自适应       │
│                                         │
├─────────────────────────────────────────┤
│  分页 (Pagination)         16px padding │
└─────────────────────────────────────────┘
```

### 6.3 详情页布局

```
┌─────────────────────────────────────────┐
│  标题区域                     24px margin│
├─────────────────────────────────────────┤
│  ┌─────────────┐  ┌─────────────────┐  │
│  │  基本信息    │  │   扩展信息      │  │
│  │  卡片        │  │   卡片          │  │
│  └─────────────┘  └─────────────────┘  │
├─────────────────────────────────────────┤
│  附件/日志/其他                     ...  │
└─────────────────────────────────────────┘
```

## 7. 制造业特色 UI

### 7.1 状态指示

```vue
<!-- 设备状态指示 -->
<template>
  <div class="status-indicator">
    <span class="status-dot" :class="statusClass"></span>
    <span class="status-text">{{ statusText }}</span>
  </div>
</template>

<style scoped>
.status-dot {
  display: inline-block;
  width: 8px;
  height: 8px;
  border-radius: 50%;
  margin-right: 6px;
}

.status-dot.running { background: #67C23A; }
.status-dot.idle { background: #E6A23C; }
.status-dot.fault { background: #F56C6C; animation: pulse 2s infinite; }
.status-dot.maintenance { background: #909399; }

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.5; }
}
</style>
```

### 7.2 进度条

```vue
<!-- 工单进度 -->
<el-progress
  :percentage="progress"
  :status="progressStatus"
  :stroke-width="20"
  :text-inside="true"
/>
```

### 7.3 看板卡片

```vue
<!-- Andon 看板卡片 -->
<template>
  <div class="kanban-card" :class="['level-' + level]">
    <div class="card-header">
      <span class="line-name">{{ lineName }}</span>
      <span class="call-time">{{ callTime }}</span>
    </div>
    <div class="card-body">
      <div class="call-type">{{ callType }}</div>
      <div class="call-desc">{{ description }}</div>
    </div>
  </div>
</template>

<style scoped>
.kanban-card {
  background: #FFF;
  border-left: 4px solid;
  border-radius: 4px;
  padding: 16px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.kanban-card.level-critical { border-left-color: #F56C6C; }
.kanban-card.level-major { border-left-color: #E6A23C; }
.kanban-card.level-minor { border-left-color: #409EFF; }
</style>
```

### 7.4 甘特图 (APS)

```vue
<!-- 使用 Gantt 组件 -->
<gantt-chart
  :tasks="tasks"
  :view-mode="viewMode"
  @task-click="handleTaskClick"
/>
```

## 8. 动画规范

### 8.1 过渡动画

```css
/* 淡入淡出 */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* 滑入滑出 */
.slide-enter-active,
.slide-leave-active {
  transition: transform 0.3s ease;
}

.slide-enter-from {
  transform: translateX(100%);
}

.slide-leave-to {
  transform: translateX(-100%);
}
```

### 8.2 交互动画

```css
/* 按钮悬停 */
.el-button {
  transition: all 0.2s ease;
}

.el-button:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 8px rgba(64, 158, 255, 0.3);
}

/* 数字跳动 */
.number-change {
  animation: numberPop 0.3s ease;
}

@keyframes numberPop {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.1); color: #409EFF; }
}
```

## 9. 响应式断点

```
大屏:              ≥1920px  (展示更多列)
桌面:              ≥1440px  (标准布局)
笔记本:            ≥1200px  (可接受压缩)
平板:              ≥992px   (单列布局)
```

## 10. 无障碍设计

### 10.1 颜色对比度

- 文字与背景对比度 ≥ 4.5:1
- 大文字对比度 ≥ 3:1
- 禁止仅用颜色传达信息

### 10.2 键盘导航

- 所有交互元素可键盘操作
- 焦点样式清晰可见
- Tab 顺序合理

### 10.3 屏幕阅读器

- 所有图片有 alt 属性
- 表单有 label 关联
- 动态内容有 aria-live 通知

## 11. 主题定制

### 11.1 深色主题

```css
/* 可选深色主题支持 */
.dark-theme {
  --bg-color: #1a1a1a;
  --text-primary: #E5E5E5;
  --border-color: #3A3A3A;
}
```

### 11.2 色弱模式

```css
/* 色弱友好调色板 */
.color-blind-safe {
  --primary: #0077BB;
  --success: #009E73;
  --warning: #EE7733;
  --danger: #CC3311;
}
```
