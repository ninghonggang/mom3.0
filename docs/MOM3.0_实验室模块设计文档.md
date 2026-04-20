# MOM3.0 实验室模块设计文档

**版本**: V2.0 | **所属模块**: M11实验室、M12量检具 | **基于**: [MOM3.0_主设计文档](./MOM3.0_主设计文档.md)

---

## 1. 模块概述

### 1.1 功能定位

实验室模块包含M11实验室管理和M12量检具管理两部分，负责样品检测、仪器校准、量具借用归还等核心功能。

### 1.2 核心功能

| 功能 | 说明 |
|------|------|
| 检测申请 | 样品送检申请管理 |
| 检测结果 | 检验报告录入管理 |
| 仪器台账 | 实验室仪器设备管理 |
| 校准记录 | 仪器校准历史记录 |
| 量检具台账 | 卡尺/千分表等量具管理 |
| 量检具校准 | 量检具定期校准 |
| 借用管理 | 量检具借用归还 |

---

## 2. 页面清单

| 页面 | 路由路径 | 核心功能 |
|------|----------|----------|
| 检测申请 | `/lab/request` | 检测申请CRUD、状态管理 |
| 检测结果 | `/lab/result` | 检验结果录入、报告 |
| 实验室仪器 | `/lab/instrument` | 仪器台账管理 |
| 校准记录 | `/lab/calibration` | 校准历史查询 |
| 量检具台账 | `/lab/gauge` | 量检具管理 |
| 量检具校准 | `/lab/gauge-calibration` | 校准记录 |
| 借用管理 | `/lab/borrow` | 借用/归还操作 |

---

## 3. UI设计规范

### 3.1 页面基本结构

同MES模块标准布局：搜索+工具栏+表格+详情弹窗。

### 3.2 状态映射

**检测申请状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| SUBMITTED | warning | 已提交 |
| ACCEPTED | primary | 已接受 |
| INSPECTING | info | 检验中 |
| COMPLETED | success | 已完成 |
| CANCELLED | info | 已取消 |

**校准状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| VALID | success | 有效 |
| EXPIRED | danger | 已过期 |
| CALIBRATING | primary | 校准中 |
| LOCKED | warning | 已锁定 |

**借用状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| BORROWED | warning | 已借出 |
| RETURNED | success | 已归还 |
| OVERDUE | danger | 超期 |

### 3.3 校准到期提醒

```javascript
// 校准状态判断
const getCalibrationStatus = (nextDate) => {
  const today = new Date()
  const expiry = new Date(nextDate)
  const daysUntilExpiry = Math.ceil((expiry - today) / (1000 * 60 * 60 * 24))

  if (daysUntilExpiry < 0) return { type: 'danger', text: '已过期' }
  if (daysUntilExpiry <= 7) return { type: 'warning', text: `${daysUntilExpiry}天后到期` }
  return { type: 'success', text: '有效' }
}
```

---

## 4. 业务流程

### 4.1 检测申请流程

```
创建申请 → 提交 → 接受指派 → 检验 → 录入结果 → 完成
```

### 4.2 校准管理流程

```
仪器/量检具 → 到期提醒 → 送检/自检 → 校准记录 → 更新下次到期日
```

### 4.3 借用管理流程

```
查询可用量检具 → 借用登记 → 领用 → 使用 → 归还 → 确认归还
```

---

## 5. 数据模型

### 5.1 检测申请

| 字段 | 类型 | 说明 |
|------|------|------|
| request_no | VARCHAR(50) | 申请单号 |
| request_type | VARCHAR(20) | INTERNAL/EXTERNAL |
| applicant_id | BIGINT | 申请人ID |
| sample_name | VARCHAR(200) | 样品名称 |
| sample_code | VARCHAR(100) | 样品编码 |
| sample_batch | VARCHAR(100) | 样品批次 |
| inspect_items | JSONB | 检测项目列表 |
| priority | VARCHAR(10) | URGENT/HIGH/NORMAL |
| status | VARCHAR(20) | 状态 |

### 5.2 实验室仪器

| 字段 | 类型 | 说明 |
|------|------|------|
| instrument_code | VARCHAR(50) | 仪器编码 |
| instrument_name | VARCHAR(100) | 仪器名称 |
| instrument_type | VARCHAR(30) | 仪器类型 |
| manufacturer | VARCHAR(100) | 制造商 |
| calibration_cycle | INTEGER | 校准周期(天) |
| last_calibration_date | DATE | 上次校准 |
| next_calibration_date | DATE | 下次校准 |
| calibration_status | VARCHAR(20) | 校准状态 |
| location | VARCHAR(100) | 存放位置 |

### 5.3 量检具

| 字段 | 类型 | 说明 |
|------|------|------|
| gauge_code | VARCHAR(50) | 量检具编码 |
| gauge_name | VARCHAR(100) | 名称 |
| gauge_type | VARCHAR(30) | 类型(卡尺/千分尺等) |
| specification | VARCHAR(100) | 规格型号 |
| calibration_cycle | INTEGER | 校准周期(天) |
| calibration_status | VARCHAR(20) | 校准状态 |
| current_location | VARCHAR(100) | 当前位置 |
| current_holder | BIGINT | 持有人ID |

### 5.4 借用记录

| 字段 | 类型 | 说明 |
|------|------|------|
| gauge_id | BIGINT | 量检具ID |
| borrower_id | BIGINT | 借用人ID |
| borrow_time | TIMESTAMP | 借用时间 |
| expected_return_time | TIMESTAMP | 预计归还 |
| actual_return_time | TIMESTAMP | 实际归还 |
| status | VARCHAR(20) | 状态 |

---

## 6. API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /lab/requests/list | 检测申请列表 |
| POST | /lab/requests | 创建申请 |
| PUT | /lab/requests/:id/accept | 接受申请 |
| PUT | /lab/requests/:id/complete | 完成检测 |
| GET | /lab/instruments/list | 仪器列表 |
| GET | /lab/instruments/expiring | 即将到期仪器 |
| GET | /lab/gauges/list | 量检具列表 |
| GET | /lab/gauges/expiring | 即将到期量检具 |
| POST | /lab/gauges/:id/borrow | 借用量检具 |
| PUT | /lab/gauges/:id/return | 归还量检具 |

---

## 7. 关联文档

- [MOM3.0_主设计文档](./MOM3.0_主设计文档.md) - 系统总览
- [MOM3.0_UI设计规范](./MOM3.0_UI设计规范.md) - UI规范详情
