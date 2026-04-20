# MOM3.0 系统集成模块设计文档

**版本**: V2.0 | **所属模块**: M14系统集成 | **基于**: [MOM3.0_主设计文档](./MOM3.0_主设计文档.md)

---

## 1. 模块概述

### 1.1 功能定位

系统集成模块负责MOM3.0与外部系统（ERP、AGV、视觉检测等）的对接，实现数据同步和业务协同。

### 1.2 核心功能

| 功能 | 说明 |
|------|------|
| 金蝶ERP集成 | BOM/工单/报工数据同步 |
| AGV调度集成 | 物料配送任务下发 |
| 视觉检测对接 | AI检测结果接收 |
| 消息推送 | 飞书/企微通知 |
| 供应商Portal | 供应商ASN管理 |

---

## 2. 页面清单

| 页面 | 路由路径 | 核心功能 |
|------|----------|----------|
| 接口配置 | `/integration/interface-config` | 接口定义管理 |
| 字段映射 | `/integration/field-mapping` | 字段转换规则 |
| 执行日志 | `/integration/execution-log` | 接口调用日志 |
| ERP同步 | `/integration/erp-sync` | ERP同步状态 |
| AGV任务 | `/integration/agv-task` | AGV任务管理 |

---

## 3. UI设计规范

### 3.1 页面基本结构

同MES模块标准布局。

### 3.2 同步状态映射

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| PENDING | warning | 待同步 |
| SYNCING | primary | 同步中 |
| SUCCESS | success | 同步成功 |
| FAILED | danger | 同步失败 |

---

## 4. 业务流程

### 4.1 ERP同步流程

```
ERP BOM变更 → MES接收 → 解析转换 → 更新BOM → 触发MRP重算
```

### 4.2 AGV调度流程

```
工单发料 → 生成配送任务 → 下发AGV → AGV取货 → AGV送货 → 工位接收
```

---

## 5. 数据模型

### 5.1 接口配置

| 字段 | 类型 | 说明 |
|------|------|------|
| interface_code | VARCHAR(50) | 接口编码 |
| interface_name | VARCHAR(100) | 接口名称 |
| system_type | VARCHAR(20) | 目标系统 |
| endpoint_url | VARCHAR(500) | 接口地址 |
| auth_type | VARCHAR(20) | 认证方式 |
| trigger_type | VARCHAR(20) | SCHEDULE/EVENT/MANUAL |
| is_enabled | SMALLINT | 是否启用 |

### 5.2 同步日志

| 字段 | 类型 | 说明 |
|------|------|------|
| sync_type | VARCHAR(50) | 同步类型 |
| direction | VARCHAR(10) | IN/OUT |
| erp_bill_no | VARCHAR(100) | ERP单据号 |
| mes_bill_no | VARCHAR(100) | MES单据号 |
| request_body | TEXT | 请求内容 |
| response_body | TEXT | 响应内容 |
| status | VARCHAR(20) | 状态 |
| error_msg | TEXT | 错误信息 |
| retry_count | INTEGER | 重试次数 |

### 5.3 AGV任务

| 字段 | 类型 | 说明 |
|------|------|------|
| task_id | VARCHAR(50) | 任务ID |
| pickup_point | VARCHAR(100) | 取货点 |
| delivery_point | VARCHAR(100) | 送货点 |
| material_code | VARCHAR(50) | 物料编码 |
| qty | INTEGER | 数量 |
| priority | INTEGER | 优先级 |
| status | VARCHAR(20) | PENDING/EXECUTING/COMPLETED/FAILED |

---

## 6. API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /integration/interfaces | 接口配置列表 |
| POST | /integration/interfaces | 创建接口配置 |
| GET | /integration/logs | 同步日志 |
| GET | /integration/logs/:id | 日志详情 |
| POST | /integration/erp/sync-bom | 手动同步BOM |
| POST | /integration/erp/sync-order | 手动同步工单 |
| GET | /agv/tasks | AGV任务列表 |
| POST | /agv/tasks | 创建AGV任务 |
| GET | /agv/tasks/:id/status | 查询任务状态 |
| POST | /vision/results | 接收视觉检测结果 |
| GET | /vision/results | 视觉检测结果列表 |

---

## 7. 集成规范

### 7.1 ERP同步接口

| 方向 | 内容 | 触发时机 |
|------|------|---------|
| ERP→MES | BOM数据 | 每日+变更触发 |
| ERP→MES | 生产订单 | 实时 |
| MES→ERP | 报工信息 | 工序完工 |
| MES→ERP | 入库通知 | FQC合格 |

### 7.2 消息推送格式(飞书)

```json
{
  "msg_type": "interactive",
  "card": {
    "header": {"title": {"content": "🔔 异常告警", "tag": "plain_text"}},
    "elements": [
      {"tag": "div", "text": {"content": "**车间**: 平衡轴\n**工位**: WS-001", "tag": "lark_md"}}
    ]
  }
}
```

---

## 8. 关联文档

- [MOM3.0_主设计文档](./MOM3.0_主设计文档.md) - 系统总览
- [MOM3.0_UI设计规范](./MOM3.0_UI设计规范.md) - UI规范详情
