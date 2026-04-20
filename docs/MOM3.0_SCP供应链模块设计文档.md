# MOM3.0 SCP供应链模块设计文档

**版本**: V2.0 | **所属模块**: M16 SCP供应链 | **基于**: [MOM3.0_主设计文档](./MOM3.0_主设计文档.md)

---

## 1. 模块概述

### 1.1 功能定位

SCP供应链模块覆盖从客户需求到供应商交付的完整供应链链路，包括采购管理、销售管理、询价管理、供应商绩效等核心功能。

### 1.2 核心功能

| 功能 | 说明 |
|------|------|
| 采购订单 | 采购申请、审批、执行、收货 |
| 询价RFQ | 供应商询价、报价、比价 |
| 供应商报价 | 供应商历史报价管理 |
| 销售订单 | 销售接单、确认、发货 |
| 客户询价 | 客户询价处理流程 |
| 供应商绩效 | KPI评分管理 |

---

## 2. 页面清单

| 页面 | 路由路径 | 核心功能 |
|------|----------|----------|
| 采购订单 | `/scp/purchase` | 采购订单CRUD、审批、发布 |
| 询价单 | `/scp/rfq` | RFQ发布、报价、比价 |
| 供应商报价 | `/scp/supplier-quote` | 报价记录查询 |
| 销售订单 | `/scp/sales-order` | 销售订单CRUD、确认 |
| 客户询价 | `/scp/customer-inquiry` | 客户询价处理 |
| 供应商绩效 | `/scp/supplier-kpi` | KPI评分、排名 |

---

## 3. UI设计规范

### 3.1 页面基本结构

同MES模块标准布局：搜索+工具栏+表格+详情弹窗。

### 3.2 采购订单详情页按钮

```vue
<template #footer>
  <el-button @click="detailVisible = false">关闭</el-button>
  <el-button type="success" v-if="detailData.approval_status === 'PENDING'" @click="handleApprove">
    审批通过
  </el-button>
  <el-button type="danger" v-if="detailData.approval_status === 'PENDING'" @click="handleReject">
    拒绝
  </el-button>
  <el-button type="warning" v-if="['DRAFT', 'PENDING', 'APPROVED'].includes(detailData.status)" @click="handleCancel">
    取消
  </el-button>
  <el-button type="primary" v-if="detailData.status === 'APPROVED'" @click="handleIssue">
    发布执行
  </el-button>
  <el-button type="danger" v-if="['ISSUED', 'PARTIAL', 'RECEIVED'].includes(detailData.status)" @click="handleClose">
    关闭订单
  </el-button>
</template>
```

### 3.3 状态映射

**采购订单状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| DRAFT | info | 草稿 |
| PENDING | warning | 待审批 |
| APPROVED | primary | 审批通过 |
| ISSUED | primary | 已发布 |
| PARTIAL | warning | 部分收货 |
| RECEIVED | success | 已收货 |
| CLOSED | info | 已关闭 |
| CANCELLED | info | 已取消 |

**客户询价状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| DRAFT | info | 草稿 |
| SENT | primary | 已发送 |
| QUOTED | warning | 已报价 |
| WON | success | 已赢单 |
| LOST | danger | 已丢单 |
| CANCELLED | info | 已取消 |

---

## 4. 业务流程

### 4.1 采购订单流程

```
创建 → 提交审批 → 审批通过 → 发布执行 → 供应商发货 → ASN到货 → 收货入库
```

### 4.2 客户询价流程

```
创建询价 → 发送供应商 → 收集报价 → 选择最优 → 赢单/丢单
```

---

## 5. 数据模型

### 5.1 采购订单

| 字段 | 类型 | 说明 |
|------|------|------|
| order_no | VARCHAR(50) | 采购单号 |
| supplier_id | BIGINT | 供应商ID |
| order_date | DATE | 订货日期 |
| expected_date | DATE | 期望交付日期 |
| total_amount | DECIMAL(18,2) | 订单总额 |
| status | VARCHAR(20) | 状态 |
| approval_status | VARCHAR(20) | 审批状态 |

### 5.2 客户询价

| 字段 | 类型 | 说明 |
|------|------|------|
| inquiry_no | VARCHAR(50) | 询价单号 |
| customer_id | BIGINT | 客户ID |
| inquiry_date | DATE | 询价日期 |
| quoted_amount | DECIMAL(18,2) | 报价金额 |
| winner_supplier_id | BIGINT | 中标供应商 |
| status | VARCHAR(20) | 状态 |

---

## 6. API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /scp/purchase-orders/list | 采购订单列表 |
| POST | /scp/purchase-orders | 创建采购订单 |
| POST | /scp/purchase-orders/:id/submit | 提交审批 |
| POST | /scp/purchase-orders/:id/approve | 审批通过 |
| POST | /scp/purchase-orders/:id/issue | 发布执行 |
| GET | /scp/rfq/list | 询价单列表 |
| POST | /scp/rfq/:id/award | RFQ授标 |
| GET | /scp/customer-inquiry/list | 客户询价列表 |
| POST | /scp/customer-inquiry/:id/send | 发送询价 |
| POST | /scp/customer-inquiry/:id/quote | 报价 |
| POST | /scp/customer-inquiry/:id/win | 标记赢单 |

---

## 7. 关联文档

- [MOM3.0_主设计文档](./MOM3.0_主设计文档.md) - 系统总览
- [MOM3.0_UI设计规范](./MOM3.0_UI设计规范.md) - UI规范详情

---

## 8. 数据表DDL（待补充）

### 8.1 采购计划表 (scp_purchase_plan)

```sql
-- 采购计划表：存储MPS/MRP驱动的采购计划
CREATE TABLE scp_purchase_plan (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    plan_no VARCHAR(50) UNIQUE NOT NULL,           -- 计划单号
    plan_type VARCHAR(20) NOT NULL,                -- 计划类型：MPS/MRP
    source_type VARCHAR(20),                        -- 来源类型：工单/预测/安全库存
    source_no VARCHAR(50),                         -- 来源单号
    supplier_id BIGINT,                             -- 供应商ID
    supplier_code VARCHAR(50),                      -- 供应商编码
    supplier_name VARCHAR(200),                     -- 供应商名称
    plan_date DATE NOT NULL,                        -- 计划日期
    required_date DATE,                             -- 需求日期
    total_amount DECIMAL(18,2),                     -- 计划总额
    total_qty DECIMAL(18,4),                        -- 计划总量
    confirmed_qty DECIMAL(18,4),                     -- 已确认数量
    status VARCHAR(20) DEFAULT 'DRAFT',             -- 状态：DRAFT/CONFIRMED/PUBLISHED/CLOSED
    approval_status VARCHAR(20) DEFAULT 'PENDING',  -- 审批状态
    remarks TEXT,
    created_by VARCHAR(100),
    updated_by VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 采购计划明细表
CREATE TABLE scp_purchase_plan_item (
    id BIGSERIAL PRIMARY KEY,
    plan_id BIGINT NOT NULL REFERENCES scp_purchase_plan(id),
    line_no INTEGER NOT NULL,                       -- 行号
    material_id BIGINT NOT NULL,                    -- 物料ID
    material_code VARCHAR(50) NOT NULL,             -- 物料编码
    material_name VARCHAR(200),                     -- 物料名称
    specification VARCHAR(200),                      -- 规格型号
    unit VARCHAR(20),                               -- 单位
    plan_qty DECIMAL(18,4) NOT NULL,                -- 计划数量
    confirmed_qty DECIMAL(18,4),                     -- 已确认数量
    delivered_qty DECIMAL(18,4),                     -- 已发货数量
    unit_price DECIMAL(18,4),                       -- 单价
    line_amount DECIMAL(18,4),                       -- 行金额
    required_date DATE,                              -- 需求日期
    promised_date DATE,                              -- 承诺日期
    status VARCHAR(20) DEFAULT 'PENDING',           -- 状态
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### 8.2 供应商主表 (scp_supplier)

```sql
-- 供应商主表
CREATE TABLE scp_supplier (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    code VARCHAR(50) UNIQUE NOT NULL,               -- 供应商编码
    name VARCHAR(200) NOT NULL,                     -- 供应商名称
    type VARCHAR(20),                               -- 供应商类型：原材料/辅料/设备/服务
    category VARCHAR(50),                           -- 物料类别
    level VARCHAR(10) DEFAULT 'B',                  -- 供应商等级：A/B/C
    contact VARCHAR(100),                            -- 联系人
    phone VARCHAR(50),                              -- 联系电话
    email VARCHAR(100),                             -- 邮箱
    address VARCHAR(500),                           -- 地址
    status INTEGER DEFAULT 1,                       -- 状态：1=启用/2=禁用
    credit_limit DECIMAL(18,2),                     -- 信用额度
    payment_terms VARCHAR(100),                     -- 付款条款
    tax_rate DECIMAL(5,2),                          -- 税率
    bank_name VARCHAR(200),                          -- 开户银行
    bank_account VARCHAR(100),                       -- 银行账号
    tax_no VARCHAR(50),                             -- 税务登记号
    business_license VARCHAR(200),                   -- 营业执照
    cert_expire_date DATE,                          -- 证书过期日期
    evaluation_score DECIMAL(5,2),                  -- 评价得分
    last_eval_date DATE,                           -- 最近评价日期
    remarks TEXT,
    created_by VARCHAR(100),
    updated_by VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### 8.3 供应商联系人表 (scp_supplier_contact)

```sql
-- 供应商联系人表
CREATE TABLE scp_supplier_contact (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    supplier_id BIGINT NOT NULL,                    -- 供应商ID
    contact_name VARCHAR(100) NOT NULL,             -- 联系人姓名
    gender VARCHAR(10),                              -- 性别
    position VARCHAR(100),                          -- 职位
    department VARCHAR(100),                        -- 部门
    phone VARCHAR(50),                              -- 办公电话
    mobile VARCHAR(50),                             -- 手机
    email VARCHAR(100),                             -- 邮箱
    fax VARCHAR(50),                                -- 传真
    is_primary INTEGER DEFAULT 0,                   -- 是否主要联系人
    remarks TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### 8.4 供应商银行账户表 (scp_supplier_bank)

```sql
-- 供应商银行账户表
CREATE TABLE scp_supplier_bank (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    supplier_id BIGINT NOT NULL,                    -- 供应商ID
    bank_name VARCHAR(200) NOT NULL,                -- 开户银行
    bank_branch VARCHAR(200),                       -- 支行名称
    bank_account VARCHAR(100) NOT NULL,              -- 银行账号
    account_name VARCHAR(200),                      -- 账户名称
    account_type VARCHAR(20),                       -- 账户类型：对公/对私
    is_default INTEGER DEFAULT 0,                   -- 是否默认账户
    currency VARCHAR(10) DEFAULT 'CNY',             -- 币种
    swift_code VARCHAR(20),                         -- SWIFT代码
    remarks TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### 8.5 客户主表 (scp_customer)

```sql
-- 客户主表
CREATE TABLE scp_customer (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    code VARCHAR(50) UNIQUE NOT NULL,               -- 客户编码
    name VARCHAR(200) NOT NULL,                     -- 客户名称
    type VARCHAR(20),                               -- 客户类型：企业/个人
    category VARCHAR(50),                           -- 客户类别
    level VARCHAR(10) DEFAULT 'B',                  -- 客户等级：A/B/C
    contact VARCHAR(100),                            -- 联系人
    phone VARCHAR(50),                              -- 联系电话
    email VARCHAR(100),                             -- 邮箱
    address VARCHAR(500),                           -- 地址
    status INTEGER DEFAULT 1,                       -- 状态：1=启用/2=禁用
    credit_limit DECIMAL(18,2),                     -- 信用额度
    used_credit DECIMAL(18,2) DEFAULT 0,            -- 已用信用
    payment_terms VARCHAR(100),                     -- 付款条款
    tax_rate DECIMAL(5,2),                          -- 税率
    tax_no VARCHAR(50),                             -- 税务登记号
    bank_name VARCHAR(200),                        -- 开户银行
    bank_account VARCHAR(100),                     -- 银行账号
    delivery_address VARCHAR(500),                 -- 交货地址
    sales_person VARCHAR(100),                      -- 业务员
    discount_rate DECIMAL(5,2),                     -- 折扣率
    remarks TEXT,
    created_by VARCHAR(100),
    updated_by VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### 8.6 客户联系人表 (scp_customer_contact)

```sql
-- 客户联系人表
CREATE TABLE scp_customer_contact (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    customer_id BIGINT NOT NULL,                    -- 客户ID
    contact_name VARCHAR(100) NOT NULL,            -- 联系人姓名
    gender VARCHAR(10),                              -- 性别
    position VARCHAR(100),                          -- 职位
    department VARCHAR(100),                        -- 部门
    phone VARCHAR(50),                              -- 办公电话
    mobile VARCHAR(50),                             -- 手机
    email VARCHAR(100),                             -- 邮箱
    fax VARCHAR(50),                                -- 传真
    is_primary INTEGER DEFAULT 0,                  -- 是否主要联系人
    is_delivery INTEGER DEFAULT 0,                  -- 是否收货联系人
    remarks TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### 8.7 客户银行账户表 (scp_customer_bank)

```sql
-- 客户银行账户表
CREATE TABLE scp_customer_bank (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    customer_id BIGINT NOT NULL,                    -- 客户ID
    bank_name VARCHAR(200) NOT NULL,                -- 开户银行
    bank_branch VARCHAR(200),                       -- 支行名称
    bank_account VARCHAR(100) NOT NULL,              -- 银行账号
    account_name VARCHAR(200),                      -- 账户名称
    account_type VARCHAR(20),                       -- 账户类型
    is_default INTEGER DEFAULT 0,                   -- 是否默认账户
    currency VARCHAR(10) DEFAULT 'CNY',             -- 币种
    remarks TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### 8.8 客户信用表 (scp_customer_credit)

```sql
-- 客户信用表
CREATE TABLE scp_customer_credit (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    customer_id BIGINT NOT NULL,                    -- 客户ID
    credit_limit DECIMAL(18,2) NOT NULL,           -- 信用额度
    used_credit DECIMAL(18,2) DEFAULT 0,            -- 已用信用
    available_credit DECIMAL(18,2),                 -- 可用信用
    frozen_credit DECIMAL(18,2) DEFAULT 0,          -- 冻结信用
    credit_level VARCHAR(10),                       -- 信用等级
    overdue_amount DECIMAL(18,2) DEFAULT 0,         -- 逾期金额
    overdue_times INTEGER DEFAULT 0,                 -- 逾期次数
    last_overdue_date DATE,                         -- 最近逾期日期
    eval_date DATE,                                 -- 评估日期
    eval_by VARCHAR(100),                           -- 评估人
    effective_date DATE,                            -- 生效日期
    expire_date DATE,                               -- 失效日期
    status INTEGER DEFAULT 1,                       -- 状态
    remarks TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 客户信用记录表
CREATE TABLE scp_customer_credit_log (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    credit_id BIGINT NOT NULL,                      -- 信用ID
    customer_id BIGINT NOT NULL,                    -- 客户ID
    change_type VARCHAR(20) NOT NULL,               -- 变更类型：USE/RELEASE/FREEZE/UNFREEZE/ADJUST
    change_amount DECIMAL(18,2),                    -- 变更金额
    before_credit DECIMAL(18,2),                     -- 变更前信用
    after_credit DECIMAL(18,2),                      -- 变更后信用
    order_no VARCHAR(50),                           -- 相关单号
    reason VARCHAR(500),                            -- 变更原因
    operator VARCHAR(100),                         -- 操作人
    created_at TIMESTAMP DEFAULT NOW()
);
```

---

## 9. 供应商确认流程API

### 9.1 供应商待确认列表

```
POST /wms/purchase-plan-main/supplier-to-confirm
```

**功能说明**：供应商查看待确认的要货计划

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| supplier_id | BIGINT | 是 | 供应商ID |
| status | STRING | 否 | 状态筛选 |

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "plan_no": "PP-2026-04001",
        "plan_date": "2026-04-01",
        "required_date": "2026-04-15",
        "total_amount": 125500.00,
        "status": "READY",
        "item_count": 5
      }
    ],
    "total": 10
  }
}
```

### 9.2 供应商确认明细

```
GET /wms/purchase-plan-main/supplier-confirm-detail
```

**功能说明**：供应商确认弹窗明细，查看计划详细信息

**查询参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| plan_id | BIGINT | 是 | 计划ID |

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "plan": {
      "id": 1,
      "plan_no": "PP-2026-04001",
      "supplier_name": "华东精密机械有限公司",
      "plan_date": "2026-04-01",
      "required_date": "2026-04-15"
    },
    "items": [
      {
        "line_no": 1,
        "material_code": "MAT-RAW-001",
        "material_name": "钢板A3",
        "specification": "1200*2400mm",
        "plan_qty": 100,
        "unit": "PCS",
        "confirmed_qty": 80,
        "remarks": "分批交货"
      }
    ]
  }
}
```

### 9.3 供应商确认保存

```
POST /wms/purchase-plan-main/supplier-confirm
```

**功能说明**：供应商确认要货计划，可部分接受

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| plan_id | BIGINT | 是 | 计划ID |
| confirm_items | ARRAY | 是 | 确认明细 |

**confirm_items结构**：

| 参数名 | 类型 | 说明 |
|--------|------|------|
| item_id | BIGINT | 计划明细ID |
| confirmed_qty | DECIMAL | 确认数量 |
| promised_date | DATE | 承诺日期 |
| remarks | STRING | 备注 |

### 9.4 计划员确认

```
POST /wms/purchase-plan-main/planer-confirm
```

**功能说明**：计划员审核供应商确认结果

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| plan_id | BIGINT | 是 | 计划ID |
| action | STRING | 是 | 动作：ACCEPT/REJECT |
| reason | STRING | 否 | 驳回原因 |

---

## 10. 采购统计API

### 10.1 采购汇总统计

```
GET /scp/purchase-statistics/summary
```

**功能说明**：采购金额/数量汇总统计

**查询参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| start_date | DATE | 否 | 开始日期 |
| end_date | DATE | 否 | 结束日期 |
| supplier_id | BIGINT | 否 | 供应商ID |
| group_by | STRING | 否 | 分组维度：supplier/material/month |

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "total_amount": 1250000.00,
    "total_qty": 5000,
    "order_count": 150,
    "supplier_count": 25,
    "trends": [
      {"month": "2026-01", "amount": 350000.00, "qty": 1200},
      {"month": "2026-02", "amount": 420000.00, "qty": 1800}
    ]
  }
}
```

### 10.2 供应商采购排名

```
GET /scp/purchase-statistics/supplier-ranking
```

**功能说明**：按供应商统计采购金额排名

**查询参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| start_date | DATE | 否 | 开始日期 |
| end_date | DATE | 否 | 结束日期 |
| limit | INTEGER | 否 | 返回条数，默认100 |

### 10.3 物料采购分析

```
GET /scp/purchase-statistics/material-analysis
```

**功能说明**：按物料分析采购趋势

**查询参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| material_code | STRING | 否 | 物料编码 |
| start_date | DATE | 否 | 开始日期 |
| end_date | DATE | 否 | 结束日期 |

### 10.4 采购计划执行率

```
GET /scp/purchase-statistics/plan-execution
```

**功能说明**：采购计划执行情况统计

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "total_plans": 100,
    "completed": 75,
    "in_progress": 15,
    "pending": 10,
    "execution_rate": 85.5,
    "on_time_rate": 92.3
  }
}
```

---

## 11. 客户管理API

### 11.1 客户档案CRUD

#### 创建客户

```
POST /mdm/customer/create
POST /scp/customer/create
```

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| code | STRING | 是 | 客户编码 |
| name | STRING | 是 | 客户名称 |
| type | STRING | 否 | 客户类型 |
| contact | STRING | 否 | 联系人 |
| phone | STRING | 否 | 联系电话 |
| email | STRING | 否 | 邮箱 |
| address | STRING | 否 | 地址 |
| credit_limit | DECIMAL | 否 | 信用额度 |
| payment_terms | STRING | 否 | 付款条款 |

#### 查询客户列表

```
GET /mdm/customer/list
GET /scp/customer/list
```

**查询参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| keyword | STRING | 否 | 搜索关键词 |
| status | INTEGER | 否 | 状态 |
| level | STRING | 否 | 客户等级 |

#### 获取客户详情

```
GET /mdm/customer/get?id={id}
GET /scp/customer/get?id={id}
```

#### 更新客户

```
PUT /mdm/customer/update
PUT /scp/customer/update
```

#### 删除客户

```
DELETE /mdm/customer/delete?id={id}
DELETE /scp/customer/delete?id={id}
```

### 11.2 客户联系人管理

#### 联系人列表

```
GET /mdm/customer-contact/list?customer_id={customer_id}
```

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "contact_name": "张经理",
        "position": "采购总监",
        "mobile": "13800138001",
        "email": "zhang@example.com",
        "is_primary": 1
      }
    ]
  }
}
```

#### 创建联系人

```
POST /mdm/customer-contact/create
```

#### 更新联系人

```
PUT /mdm/customer-contact/update
```

#### 删除联系人

```
DELETE /mdm/customer-contact/delete?id={id}
```

### 11.3 客户银行账户管理

#### 银行账户列表

```
GET /mdm/customer-bank/list?customer_id={customer_id}
```

#### 创建银行账户

```
POST /mdm/customer-bank/create
```

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| customer_id | BIGINT | 是 | 客户ID |
| bank_name | STRING | 是 | 开户银行 |
| bank_account | STRING | 是 | 银行账号 |
| account_name | STRING | 是 | 账户名称 |
| is_default | INTEGER | 否 | 是否默认 |

### 11.4 客户信用管理

#### 获取客户信用

```
GET /scp/customer-credit/get?customer_id={customer_id}
```

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "credit_limit": 1000000.00,
    "used_credit": 350000.00,
    "available_credit": 650000.00,
    "credit_level": "A",
    "overdue_times": 0,
    "eval_date": "2026-04-01"
  }
}
```

#### 更新客户信用

```
PUT /scp/customer-credit/update
```

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| customer_id | BIGINT | 是 | 客户ID |
| credit_limit | DECIMAL | 是 | 信用额度 |
| remarks | STRING | 否 | 备注 |

#### 信用记录查询

```
GET /scp/customer-credit/log?customer_id={customer_id}
```

### 11.5 客户发货预测

```
POST /wms/customer-delivery-forecast/create
GET /wms/customer-delivery-forecast/page
PUT /wms/customer-delivery-forecast/update
DELETE /wms/customer-delivery-forecast/delete
```

---

## 12. 供应商管理API

### 12.1 供应商档案CRUD

#### 创建供应商

```
POST /mdm/supplier/create
POST /scp/supplier/create
```

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| code | STRING | 是 | 供应商编码 |
| name | STRING | 是 | 供应商名称 |
| type | STRING | 否 | 供应商类型 |
| category | STRING | 否 | 物料类别 |
| contact | STRING | 否 | 联系人 |
| phone | STRING | 否 | 联系电话 |
| email | STRING | 否 | 邮箱 |
| address | STRING | 否 | 地址 |

#### 查询供应商列表

```
GET /mdm/supplier/list
GET /scp/supplier/list
```

#### 获取供应商详情

```
GET /mdm/supplier/get?id={id}
GET /scp/supplier/get?id={id}
```

### 12.2 供应商联系人管理

```
GET /mdm/supplier-contact/list?supplier_id={supplier_id}
POST /mdm/supplier-contact/create
PUT /mdm/supplier-contact/update
DELETE /mdm/supplier-contact/delete?id={id}
```

### 12.3 供应商银行账户管理

```
GET /mdm/supplier-bank/list?supplier_id={supplier_id}
POST /mdm/supplier-bank/create
PUT /mdm/supplier-bank/update
DELETE /mdm/supplier-bank/delete?id={id}
```

### 12.4 供应商KPI管理

```
GET /scp/supplier-kpi/list
GET /scp/supplier-kpi/monthly?supplier_id={id}&month={month}
POST /scp/supplier-kpi/create
GET /scp/supplier-kpi/ranking?month={month}
```

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": [
    {
      "rank": 1,
      "supplier_id": 1,
      "supplier_name": "华东精密机械有限公司",
      "total_score": 95.5,
      "grade": "A",
      "on_time_rate": 98.5,
      "quality_pass_rate": 99.2
    }
  ]
}
```

---

## 13. API响应格式规范

### 13.1 统一响应结构

```json
{
  "code": 0,
  "msg": "success",
  "data": {}
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| code | INTEGER | 状态码，0=成功，其他=失败 |
| msg | STRING | 消息 |
| data | OBJECT | 数据对象 |

### 13.2 分页响应结构

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "list": [],
    "total": 100,
    "page": 1,
    "page_size": 20
  }
}
```

### 13.3 错误响应示例

```json
{
  "code": 400,
  "msg": "参数错误：code不能为空",
  "data": null
}
```

---

## 14. 状态码定义

### 14.1 采购计划状态

| 状态码 | 说明 |
|--------|------|
| DRAFT | 草稿 |
| READY | 已准备 |
| CONFIRMED | 已确认 |
| PUBLISHED | 已发布 |
| PARTIAL | 部分执行 |
| CLOSED | 已关闭 |
| CANCELLED | 已取消 |

### 14.2 供应商确认状态

| 状态码 | 说明 |
|--------|------|
| PENDING | 待确认 |
| CONFIRMED | 已确认 |
| REJECTED | 已驳回 |
| PARTIAL | 部分接受 |

### 14.3 客户信用状态

| 状态码 | 说明 |
|--------|------|
| NORMAL | 正常 |
| WARNING | 预警 |
| OVERDUE | 逾期 |
| FROZEN | 冻结 |

---

## 15. 备注

本文档为MOM3.0 SCP供应链模块设计文档的补充内容，涵盖：

1. **DDL表结构**：采购计划表、供应商主表、供应商联系人表、供应商银行账户表、客户主表、客户联系人表、客户银行账户表、客户信用表

2. **供应商确认API**：供应商待确认列表、确认明细、确认保存、计划员确认

3. **采购统计API**：汇总统计、供应商排名、物料分析、计划执行率

4. **客户管理API**：客户档案CRUD、联系人管理、银行账户管理、信用管理

5. **供应商管理API**：供应商档案CRUD、联系人管理、银行账户管理、KPI管理

---

## 16. MRS外部接口

### 16.1 功能说明

MRS外部接口用于与外部MRS（物料需求计划系统）对接，实现要货计划的创建、查询和需求数据同步。

### 16.2 业务流程

```
MRS系统创建要货计划 → 同步需求数据 → 查询汇总统计 → 确认要货计划
```

### 16.3 API接口

#### 16.3.1 从MRS创建要货计划

```
POST /scp/mrs/createPlan
```

**功能说明**：接收MRS系统发送的要货计划数据，创建内部采购计划

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| mrs_no | STRING | 是 | MRS单号 |
| plan_type | STRING | 是 | 计划类型：MRS |
| supplier_id | BIGINT | 是 | 供应商ID |
| plan_date | DATE | 是 | 计划日期 |
| required_date | DATE | 是 | 需求日期 |
| items | ARRAY | 是 | 计划明细 |

**items结构**：

| 参数名 | 类型 | 说明 |
|--------|------|------|
| material_code | STRING | 物料编码 |
| material_name | STRING | 物料名称 |
| plan_qty | DECIMAL | 计划数量 |
| unit_price | DECIMAL | 单价 |
| required_date | DATE | 需求日期 |

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "plan_id": 1001,
    "plan_no": "PP-MRS-2026-04001"
  }
}
```

#### 16.3.2 获取MRS汇总统计

```
GET /scp/mrs/getStatistics
```

**功能说明**：查询MRS汇总统计数据

**查询参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| start_date | DATE | 否 | 开始日期 |
| end_date | DATE | 否 | 结束日期 |
| supplier_id | BIGINT | 否 | 供应商ID |
| status | STRING | 否 | 状态筛选 |

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "mrs_no": "MRS-2026-001",
        "statistics_date": "2026-04-15",
        "total_demand_qty": 5000,
        "total_supply_qty": 4500,
        "gap_qty": 500,
        "status": "CONFIRMED"
      }
    ],
    "total": 10
  }
}
```

#### 16.3.3 同步MRS需求数据

```
POST /scp/mrs/syncDemand
```

**功能说明**：从MRS系统同步需求数据到内部系统

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| mrs_no | STRING | 是 | MRS单号 |
| sync_type | STRING | 是 | 同步类型：FULL/INCREMENTAL |
| demand_data | ARRAY | 是 | 需求数据 |

**demand_data结构**：

| 参数名 | 类型 | 说明 |
|--------|------|------|
| material_code | STRING | 物料编码 |
| demand_qty | DECIMAL | 需求数量 |
| supply_qty | DECIMAL | 供给数量 |
| gap_qty | DECIMAL | 缺口数量 |

#### 16.3.4 获取统计明细

```
GET /scp/mrs/getStatisticsDetail
```

**功能说明**：获取MRS汇总统计的明细数据

**查询参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| main_id | BIGINT | 是 | 汇总统计ID |

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "main": {
      "id": 1,
      "mrs_no": "MRS-2026-001",
      "statistics_date": "2026-04-15"
    },
    "details": [
      {
        "material_code": "MAT-001",
        "material_name": "钢板A3",
        "demand_qty": 1000,
        "supply_qty": 900,
        "gap_qty": 100,
        "version": "V1"
      }
    ]
  }
}
```

---

## 17. QAD对接接口

### 17.1 功能说明

QAD对接接口用于与QAD ERP系统同步要货预测数据，实现跨系统数据互通。

### 17.2 业务流程

```
QAD系统 → 同步预测数据 → 预测确认 → 版本管理
```

### 17.3 API接口

#### 17.3.1 同步QAD要货预测

```
POST /scp/qad/syncForecast
```

**功能说明**：从QAD系统同步要货预测数据

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| qad_doc_no | STRING | 是 | QAD单据号 |
| sync_time | DATETIME | 是 | 同步时间 |
| forecast_data | ARRAY | 是 | 预测数据 |

**forecast_data结构**：

| 参数名 | 类型 | 说明 |
|--------|------|------|
| material_code | STRING | 物料编码 |
| material_name | STRING | 物料名称 |
| demand_qty | DECIMAL | 需求数量 |
| demand_date | DATE | 需求日期 |
| priority | INTEGER | 优先级 |

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "sync_count": 50,
    "sync_time": "2026-04-15 10:30:00"
  }
}
```

#### 17.3.2 获取预测列表

```
GET /scp/qad/getForecastList
```

**功能说明**：查询已同步的QAD预测数据列表

**查询参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| qad_doc_no | STRING | 否 | QAD单据号 |
| material_code | STRING | 否 | 物料编码 |
| start_date | DATE | 否 | 开始日期 |
| end_date | DATE | 否 | 结束日期 |
| status | STRING | 否 | 状态 |
| page | INTEGER | 否 | 页码 |
| page_size | INTEGER | 否 | 每页条数 |

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "qad_doc_no": "QAD-FC-2026-001",
        "material_code": "MAT-001",
        "material_name": "钢板A3",
        "demand_qty": 500,
        "demand_date": "2026-04-20",
        "status": "CONFIRMED",
        "version": "V2"
      }
    ],
    "total": 100,
    "page": 1,
    "page_size": 20
  }
}
```

#### 17.3.3 确认预测数据

```
POST /scp/qad/confirmForecast
```

**功能说明**：确认QAD预测数据，支持批量确认

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| forecast_ids | ARRAY | 是 | 预测ID列表 |
| confirm_items | ARRAY | 否 | 确认明细（可选） |
| remarks | STRING | 否 | 备注 |

**confirm_items结构**：

| 参数名 | 类型 | 说明 |
|--------|------|------|
| forecast_id | BIGINT | 预测ID |
| adjusted_qty | DECIMAL | 调整后数量 |
| adjusted_date | DATE | 调整后日期 |
| reason | STRING | 调整原因 |

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "confirmed_count": 10,
    "version": "V3"
  }
}
```

---

## 18. 要货预测子表接口

### 18.1 功能说明

要货预测子表接口用于管理预测明细数据，支持按主表查询和版本管理。

### 18.2 API接口

#### 18.2.1 按主表查询明细

```
GET /scp/demandforecasting-detail/listByMain
```

**功能说明**：根据主表ID查询预测明细列表

**查询参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| main_id | BIGINT | 是 | 主表ID |
| material_code | STRING | 否 | 物料编码 |
| version | STRING | 否 | 版本号 |

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "main_id": 100,
        "material_code": "MAT-001",
        "material_name": "钢板A3",
        "demand_qty": 1000,
        "supply_qty": 900,
        "gap_qty": 100,
        "version": "V1"
      }
    ],
    "total": 5
  }
}
```

#### 18.2.2 获取版本列表

```
GET /scp/demandforecasting-detail/getVersions
```

**功能说明**：获取指定物料的预测版本历史

**查询参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| material_code | STRING | 是 | 物料编码 |

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "versions": [
      {"version": "V1", "create_time": "2026-04-01 10:00:00"},
      {"version": "V2", "create_time": "2026-04-10 10:00:00"},
      {"version": "V3", "create_time": "2026-04-15 10:00:00"}
    ]
  }
}
```

#### 18.2.3 导入预测数据

```
POST /scp/demandforecasting-detail/import
```

**功能说明**：批量导入预测数据

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| main_id | BIGINT | 是 | 主表ID |
| import_data | ARRAY | 是 | 导入数据 |

**import_data结构**：

| 参数名 | 类型 | 说明 |
|--------|------|------|
| material_code | STRING | 物料编码 |
| material_name | STRING | 物料名称 |
| demand_qty | DECIMAL | 需求数量 |
| supply_qty | DECIMAL | 供给数量 |
| gap_qty | DECIMAL | 缺口数量 |

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "import_count": 50,
    "fail_count": 0,
    "errors": []
  }
}
```

---

## 19. 采购计划策略S001

### 19.1 功能说明

采购计划策略S001是驱动采购计划自动生成的核心策略引擎，根据预设规则自动计算采购需求并生成采购计划。

### 19.2 策略规则配置

#### 19.2.1 策略参数

| 参数名 | 类型 | 说明 |
|--------|------|------|
| strategy_code | STRING | 策略编码：S001 |
| strategy_name | STRING | 策略名称 |
| enabled | INTEGER | 是否启用：1=启用/0=禁用 |
| priority | INTEGER | 优先级 |
| trigger_type | STRING | 触发类型：MANUAL/AUTO/SCHEDULE |
| schedule_cron | STRING | 调度表达式（trigger_type为SCHEDULE时） |

#### 19.2.2 策略规则

| 规则编码 | 规则名称 | 说明 |
|----------|----------|------|
| R001 | 最小采购量规则 | 当采购量小于最小采购量时，按最小采购量采购 |
| R002 | 批量采购规则 | 按指定批量倍数进行采购 |
| R003 | 安全库存规则 | 采购量 = 需求数量 + 安全库存 - 当前库存 |
| R004 | 供应商交期规则 | 根据供应商交期调整采购计划 |
| R005 | 优先级规则 | 按物料优先级排序，高优先级优先采购 |

#### 19.2.3 策略执行流程

```
1. 触发条件检查 → 2. 数据采集 → 3. 规则匹配 → 4. 计划计算 → 5. 计划生成 → 6. 结果输出
```

#### 19.2.4 API接口

```
GET /scp/strategy/list?code=S001
POST /scp/strategy/execute
POST /scp/strategy/config/update
```

**execute请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| strategy_code | STRING | 是 | 策略编码 |
| start_date | DATE | 是 | 开始日期 |
| end_date | DATE | 是 | 结束日期 |
| supplier_id | BIGINT | 否 | 供应商ID（可选） |

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "execute_id": "EX-2026-0415001",
    "plans_generated": 15,
    "total_amount": 125000.00,
    "execution_time": "2026-04-15 10:30:00"
  }
}
```

---

## 20. DDL表结构（补充）

### 20.1 MRS汇总统计表 (scp_purchase_mrs_statistics)

```sql
-- MRS汇总统计表：存储MRS系统的要货计划汇总数据
CREATE TABLE scp_purchase_mrs_statistics (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    mrs_no VARCHAR(50) NOT NULL,                    -- MRS单号
    statistics_date DATE NOT NULL,                   -- 统计日期
    total_demand_qty DECIMAL(18,4) DEFAULT 0,         -- 总需求数量
    total_supply_qty DECIMAL(18,4) DEFAULT 0,         -- 总供给数量
    gap_qty DECIMAL(18,4) DEFAULT 0,                 -- 缺口数量
    status VARCHAR(20) DEFAULT 'DRAFT',              -- 状态：DRAFT/CONFIRMED/CLOSED
    creator VARCHAR(100),
    create_time TIMESTAMP DEFAULT NOW(),
    updater VARCHAR(100),
    update_time TIMESTAMP DEFAULT NOW(),
    deleted INTEGER DEFAULT 0                         -- 软删除标记
);

COMMENT ON TABLE scp_purchase_mrs_statistics IS 'MRS汇总统计表';
COMMENT ON COLUMN scp_purchase_mrs_statistics.mrs_no IS 'MRS单号，外部系统唯一标识';
COMMENT ON COLUMN scp_purchase_mrs_statistics.gap_qty IS '缺口数量，total_demand_qty - total_supply_qty';
```

### 20.2 要货预测子表 (scp_demand_forecasting_detail)

```sql
-- 要货预测子表：存储要货预测的明细数据
CREATE TABLE scp_demand_forecasting_detail (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    main_id BIGINT NOT NULL,                         -- 主表ID（关联scp_purchase_mrs_statistics）
    material_code VARCHAR(50) NOT NULL,               -- 物料编码
    material_name VARCHAR(200),                      -- 物料名称
    demand_qty DECIMAL(18,4) DEFAULT 0,              -- 需求数量
    supply_qty DECIMAL(18,4) DEFAULT 0,               -- 供给数量
    gap_qty DECIMAL(18,4) DEFAULT 0,                 -- 缺口数量
    version VARCHAR(20) DEFAULT 'V1',                -- 版本号
    creator VARCHAR(100),
    create_time TIMESTAMP DEFAULT NOW(),
    updater VARCHAR(100),
    update_time TIMESTAMP DEFAULT NOW(),
    tenant_id BIGINT NOT NULL DEFAULT 1,
    deleted INTEGER DEFAULT 0                         -- 软删除标记
);

CREATE INDEX idx_demand_forecasting_detail_main_id ON scp_demand_forecasting_detail(main_id);
CREATE INDEX idx_demand_forecasting_detail_material_code ON scp_demand_forecasting_detail(material_code);
CREATE INDEX idx_demand_forecasting_detail_version ON scp_demand_forecasting_detail(version);

COMMENT ON TABLE scp_demand_forecasting_detail IS '要货预测子表';
COMMENT ON COLUMN scp_demand_forecasting_detail.main_id IS '关联主表ID';
COMMENT ON COLUMN scp_demand_forecasting_detail.version IS '版本号，用于版本管理';
```

---

## 21. 附录：新增接口汇总

### 21.1 MRS外部接口 /scp/mrs/*

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /scp/mrs/createPlan | 从MRS创建要货计划 |
| GET | /scp/mrs/getStatistics | 获取MRS汇总统计 |
| POST | /scp/mrs/syncDemand | 同步MRS需求数据 |
| GET | /scp/mrs/getStatisticsDetail | 获取统计明细 |

### 21.2 QAD对接接口 /scp/qad/*

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /scp/qad/syncForecast | 同步QAD要货预测 |
| GET | /scp/qad/getForecastList | 获取预测列表 |
| POST | /scp/qad/confirmForecast | 确认预测数据 |

### 21.3 要货预测子表接口 /scp/demandforecasting-detail/*

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /scp/demandforecasting-detail/listByMain | 按主表查询明细 |
| GET | /scp/demandforecasting-detail/getVersions | 获取版本列表 |
| POST | /scp/demandforecasting-detail/import | 导入预测数据 |

### 21.4 采购计划策略接口 /scp/strategy/*

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /scp/strategy/list | 策略列表查询 |
| POST | /scp/strategy/execute | 执行策略 |
| POST | /scp/strategy/config/update | 更新策略配置 |
