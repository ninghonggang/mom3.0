# MOM3.0 APS计划模块设计文档

**版本**: V2.0 | **所属模块**: M04 APS计划 | **基于**: [MOM3.0_主设计文档](./MOM3.0_主设计文档.md)

---

## 1. 模块概述

### 1.1 功能定位

APS高级计划排程模块实现MPS主生产计划、MRP物料需求运算、甘特图排程、交付分析、缺料分析等核心功能，支持汽车零部件行业的多品种小批量生产模式。

### 1.2 核心功能

| 功能 | 说明 |
|------|------|
| MPS主生产计划 | 月度生产计划编制 |
| MRP运算 | 物料需求计算 |
| 工作中心 | 产能管理 |
| 甘特图排程 | 可视化排程编辑 |
| 滚动排程 | 自动调度 |
| 交付分析 | 交付率分析 |
| 缺料分析 | 物料短缺预警 |
| 换型矩阵 | 产品换型时间 |
| 产品族 | 产品族管理 |

---

## 2. 页面清单

| 页面 | 路由路径 | 核心功能 |
|------|----------|----------|
| MPS计划 | `/aps/mps` | 主生产计划列表 |
| 甘特图 | `/aps/gantt` | 排程甘特图 |
| 排程配置 | `/aps/rolling-config-list` | 滚动排程配置 |
| 交付分析 | `/aps/delivery-analysis` | 交付性能分析 |
| 交付预警 | `/aps/delivery-warning` | 交付预警列表 |
| 缺料分析 | `/aps/material-shortage` | 缺料分析 |
| 工作中心 | `/aps/work-center` | 工作中心管理 |
| 换型矩阵 | `/aps/changeover-matrix` | 产品换型时间 |
| 产品族 | `/aps/product-family` | 产品族管理 |
| 工厂日历 | `/aps/calendar` | 班次日历配置 |

---

## 3. UI设计规范

### 3.1 页面基本结构

同MES模块标准布局：搜索+工具栏+表格+详情弹窗。

### 3.2 状态映射

**计划状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| DRAFT | info | 草稿 |
| RELEASED | primary | 已下达 |
| IN_PROGRESS | warning | 执行中 |
| COMPLETED | success | 已完成 |
| CANCELLED | info | 已取消 |

**排程结果状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| RUNNING | primary | 排程中 |
| COMPLETED | success | 排程完成 |
| FAILED | danger | 排程失败 |

**交付预警等级**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| CRITICAL | danger | 严重延期 |
| HIGH | warning | 高风险 |
| MEDIUM | warning | 中风险 |
| LOW | info | 低风险 |

### 3.3 甘特图交互

```vue
<!-- 甘特图拖拽调整 -->
<template>
  <div class="gantt-container">
    <div class="gantt-timeline">
      <div class="time-header">
        <div v-for="day in days" :key="day" class="day-cell">
          {{ day }}
        </div>
      </div>
      <div class="gantt-rows">
        <div v-for="order in orders" :key="order.id" class="gantt-row">
          <div class="order-bar"
               :style="{ left: getBarLeft(order), width: getBarWidth(order) }"
               :class="order.status"
               draggable="true"
               @dragstart="handleDragStart($event, order)"
               @dragend="handleDragEnd($event, order)">
            {{ order.order_no }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
```

---

## 4. 业务流程

### 4.1 MPS计划流程

```
创建MPS计划 → 编排月计划 → 分解周计划 → 下达日计划 → 工单生成
```

### 4.2 滚动排程流程

```
每日18:00自动触发 → 读取未排工单 → 应用排程算法 → 生成排程结果 → 发布甘特图
```

### 4.3 交付预警流程

```
每日分析 → 计算剩余天数 → 风险评估 → 预警推送 → 确认处理 → 升级机制
```

---

## 5. 数据模型

### 5.1 MPS计划

| 字段 | 类型 | 说明 |
|------|------|------|
| plan_no | VARCHAR(50) | 计划编号 |
| plan_month | VARCHAR(7) | 计划月份 |
| workshop_id | BIGINT | 车间ID |
| total_orders | INTEGER | 订单数量 |
| status | VARCHAR(20) | 状态 |
| released_at | TIMESTAMP | 下达时间 |

### 5.2 滚动排程配置

| 字段 | 类型 | 说明 |
|------|------|------|
| config_code | VARCHAR(50) | 配置编码 |
| config_type | VARCHAR(20) | DAILY/HOURLY |
| trigger_cron | VARCHAR(50) | Cron表达式 |
| horizon_days | INTEGER | 排程视野(天) |
| scheduling_algorithm | VARCHAR(20) | FIFO/EDD/SPT/LPT |
| is_enabled | SMALLINT | 是否启用 |

### 5.3 排程结果

| 字段 | 类型 | 说明 |
|------|------|------|
| schedule_no | VARCHAR(50) | 排程编号 |
| schedule_date | DATE | 排程日期 |
| total_orders | INTEGER | 总订单数 |
| scheduled_orders | INTEGER | 已排订单 |
| avg_utilization | DECIMAL(5,2) | 平均利用率 |
| status | VARCHAR(20) | 状态 |

### 5.4 交付分析

| 字段 | 类型 | 说明 |
|------|------|------|
| analysis_no | VARCHAR(50) | 分析编号 |
| analysis_date | DATE | 分析日期 |
| total_orders | INTEGER | 总订单数 |
| on_time_orders | INTEGER | 准时交付数 |
| on_time_rate | DECIMAL(5,2) | 准时交付率 |
| avg_delay_days | DECIMAL(8,2) | 平均延期天数 |

### 5.5 缺料分析

| 字段 | 类型 | 说明 |
|------|------|------|
| analysis_no | VARCHAR(50) | 分析编号 |
| material_id | BIGINT | 物料ID |
| required_qty | DECIMAL(18,3) | 需求数量 |
| available_qty | DECIMAL(18,3) | 可用量 |
| shortage_qty | DECIMAL(18,3) | 缺口数量 |
| shortage_level | VARCHAR(10) | CRITICAL/HIGH/MEDIUM/LOW |

---

## 6. API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /aps/mps/list | MPS计划列表 |
| POST | /aps/mps | 创建MPS计划 |
| POST | /aps/mps/:id/release | 下达计划 |
| GET | /aps/schedule/results | 排程结果列表 |
| POST | /aps/schedule/execute | 执行排程 |
| GET | /aps/schedule/results/:id/gantt | 甘特图数据 |
| GET | /aps/delivery/analysis | 交付分析列表 |
| GET | /aps/delivery/warnings | 交付预警列表 |
| PUT | /aps/delivery/warnings/:id/acknowledge | 确认预警 |
| GET | /aps/material-shortage | 缺料分析列表 |
| GET | /aps/work-centers | 工作中心列表 |
| GET | /aps/changeover-matrix | 换型矩阵 |
| POST | /aps/changeover-matrix | 添加换型时间 |
| GET | /aps/product-families | 产品族列表 |
| GET | /aps/calendar | 工厂日历 |

---

## 7. 关联文档

- [MOM3.0_主设计文档](./MOM3.0_主设计文档.md) - 系统总览
- [MOM3.0_UI设计规范](./MOM3.0_UI