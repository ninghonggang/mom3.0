# MOM3.0 结算模块设计文档

**版本**: V2.0 | **所属模块**: M07结算管理 | **基于**: [MOM3.0_主设计文档](./MOM3.0_主设计文档.md)

---

## 1. 模块概述

### 1.1 功能定位

结算模块处理采购和销售业务的财务结算，支持在线实时结算和离线手工结算两种模式，实现货、票、款一致的财务闭环管理。

### 1.2 核心功能

| 功能 | 说明 |
|------|------|
| 采购结算 | 采购到货结算、退货结算 |
| 销售结算 | 销售发货结算、退货结算 |
| 预付款管理 | 采购预付/销售预收 |
| 付款申请 | 付款审批流程 |
| 收款管理 | 销售收款认领 |
| 供应商对账 | 供应商对账单生成与确认 |

---

## 2. 页面清单

| 页面 | 路由路径 | 核心功能 |
|------|----------|----------|
| 采购结算 | `/fin/purchase-settlement` | 采购结算CRUD、审批 |
| 销售结算 | `/fin/sales-settlement` | 销售结算CRUD、审批 |
| 预付款管理 | `/fin/advance` | 预付款登记、使用 |
| 付款申请 | `/fin/payment-request` | 付款申请、审批 |
| 收款管理 | `/fin/receipt` | 收款认领、冲账 |
| 供应商对账 | `/fin/supplier-statement` | 对账单生成、确认 |

---

## 3. UI设计规范

### 3.1 页面基本结构

同MES模块标准布局：搜索+工具栏+表格+详情弹窗。

### 3.2 状态映射

**结算单状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| PENDING | warning | 待审批 |
| APPROVED | primary | 审批通过 |
| PAID | success | 已付款 |
| CANCELLED | info | 已取消 |

**付款申请状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| PENDING | warning | 待审批 |
| APPROVED | primary | 审批通过 |
| UNPAID | info | 未付款 |
| PARTIAL_PAID | warning | 部分付款 |
| PAID | success | 已付款 |

**对账单状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| PENDING | warning | 待确认 |
| CONFIRMED | success | 已确认 |
| DISPUTED | danger | 有异议 |

### 3.3 结算单详情页按钮

```vue
<template #footer>
  <el-button @click="detailVisible = false">关闭</el-button>
  <el-button type="primary" v-if="detailData.status === 'PENDING'" @click="handleApprove">
    审批通过
  </el-button>
  <el-button type="danger" v-if="detailData.status === 'PENDING'" @click="handleReject">
    拒绝
  </el-button>
  <el-button type="warning" v-if="['PENDING', 'APPROVED'].includes(detailData.status)" @click="handleCancel">
    取消结算
  </el-button>
  <el-button type="success" v-if="detailData.status === 'APPROVED'" @click="handlePay">
    确认付款
  </el-button>
</template>
```

---

## 4. 业务流程

### 4.1 采购结算流程

```
采购收货 → 生成结算单 → 提交审批 → 审批通过 → 确认付款 → 完成
```

### 4.2 销售结算流程

```
销售发货 → 生成结算单 → 提交审批 → 审批通过 → 收款认领 → 完成
```

### 4.3 付款申请流程

```
创建申请 → 提交审批 → 各级审批 → 审批通过 → 出纳付款 → 完成
```

---

## 5. 数据模型

### 5.1 采购结算

| 字段 | 类型 | 说明 |
|------|------|------|
| settlement_no | VARCHAR(50) | 结算单号 |
| settlement_type | VARCHAR(20) | NORMAL/RETURN/ADVANCE |
| related_type | VARCHAR(20) | 来源类型 |
| related_id | BIGINT | 来源ID |
| supplier_id | BIGINT | 供应商ID |
| invoice_no | VARCHAR(50) | 发票号 |
| goods_amount | DECIMAL(18,2) | 货款金额 |
| tax_amount | DECIMAL(18,2) | 税额 |
| total_amount | DECIMAL(18,2) | 结算总额 |
| paid_amount | DECIMAL(18,2) | 已付款金额 |
| payment_due_date | DATE | 应付日期 |
| status | VARCHAR(20) | 状态 |

### 5.2 销售结算

| 字段 | 类型 | 说明 |
|------|------|------|
| settlement_no | VARCHAR(50) | 结算单号 |
| customer_id | BIGINT | 客户ID |
| invoice_no | VARCHAR(50) | 发票号 |
| goods_amount | DECIMAL(18,2) | 货款金额 |
| tax_amount | DECIMAL(18,2) | 税额 |
| total_amount | DECIMAL(18,2) | 结算总额 |
| received_amount | DECIMAL(18,2) | 已收款金额 |
| payment_due_date | DATE | 应收日期 |
| status | VARCHAR(20) | 状态 |

### 5.3 付款申请

| 字段 | 类型 | 说明 |
|------|------|------|
| request_no | VARCHAR(50) | 申请单号 |
| request_type | VARCHAR(20) | PURCHASE/SALES/EXPENSE |
| supplier_customer_id | BIGINT | 供应商/客户ID |
| request_amount | DECIMAL(18,2) | 申请金额 |
| purpose | VARCHAR(200) | 付款用途 |
| bank_name | VARCHAR(100) | 银行名称 |
| bank_account | VARCHAR(100) | 银行账号 |
| status | VARCHAR(20) | 状态 |
| approval_status | VARCHAR(20) | 审批状态 |
| payment_status | VARCHAR(20) | 付款状态 |

---

## 6. API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /fin/purchase-settlements/list | 采购结算列表 |
| POST | /fin/purchase-settlements | 创建采购结算 |
| POST | /fin/purchase-settlements/:id/submit | 提交审批 |
| POST | /fin/purchase-settlements/:id/approve | 审批通过 |
| POST | /fin/purchase-settlements/:id/pay | 确认付款 |
| GET | /fin/sales-settlements/list | 销售结算列表 |
| POST | /fin/sales-settlements | 创建销售结算 |
| POST | /fin/sales-settlements/:id/submit | 提交审批 |
| GET | /fin/payment-requests/list | 付款申请列表 |
| POST | /fin/payment-requests | 创建申请 |
| POST | /fin/payment-requests/:id/approve | 审批通过 |
| POST | /fin/payment-requests/:id/pay | 确认付款 |
| GET | /fin/supplier-statements/list | 对账单列表 |
| POST | /fin/supplier-statements | 生成对账单 |

---

## 7. 关联文档

- [MOM3.0_主设计文档](./MOM3.0_主设计文档.md) - 系统总览
- [MOM3.0_UI设计规范](./MOM3.0_UI设计规范.md) - UI规范详情
