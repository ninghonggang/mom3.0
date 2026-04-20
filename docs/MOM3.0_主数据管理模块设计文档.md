# MOM3.0 主数据管理模块设计文档

**版本**: V2.0 | **所属模块**: M02主数据管理 | **基于**: [MOM3.0_主设计文档](./MOM3.0_主设计文档.md)

---

## 1. 模块概述

### 1.1 功能定位

主数据管理模块是MOM3.0的基础数据层，覆盖物料、BOM、工艺、客户、供应商等核心主数据的管理和维护。

### 1.2 核心功能

| 功能 | 说明 |
|------|------|
| 物料管理 | 物料主数据、分类、单位 |
| BOM管理 | 物料清单、批量导入 |
| 工艺路线 | 工序定义、工时设置 |
| 客户管理 | 客户主数据、联系人 |
| 供应商管理 | 供应商主数据、联系人 |
| 车间产线 | 车间、产线、工位 |
| 班次管理 | 工作班次定义 |

---

## 2. 页面清单

| 页面 | 路由路径 | 核心功能 |
|------|----------|----------|
| 物料管理 | `/mdm/material` | 物料列表、新增/编辑 |
| 物料分类 | `/mdm/material-category` | 物料分类管理 |
| BOM管理 | `/mdm/bom` | BOM列表、结构查看 |
| BOM编辑 | `/mdm/bom/:id` | BOM编辑 |
| 工艺路线 | `/mdm/process-route` | 工艺路线管理 |
| 客户列表 | `/mdm/customer-list` | 客户主数据 |
| 客户详情 | `/mdm/customer-view/:id` | 客户详情 |
| 供应商列表 | `/mdm/supplier-list` | 供应商主数据 |
| 供应商详情 | `/mdm/supplier-view/:id` | 供应商详情 |
| 车间产线 | `/mdm/workshop` | 车间产线配置 |
| 班次管理 | `/mdm/shift` | 班次定义 |

---

## 3. UI设计规范

### 3.1 页面基本结构

同MES模块标准布局。

### 3.2 状态映射

**客户/供应商状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| PENDING | warning | 待审核 |
| ACTIVE | success | 正常 |
| SUSPENDED | info | 已停用 |
| BLACKLIST | danger | 黑名单 |

**供应商等级**

| 等级 | 说明 |
|------|------|
| A级 | 90-100分，首选供应商 |
| B级 | 75-89分，正常采购 |
| C级 | 60-74分，观察供应商 |
| D级 | <60分，整改或淘汰 |

---

## 4. 业务流程

### 4.1 物料新增流程

```
创建物料 → 完善属性 → 关联分类 → 设置单位 → 审核发布
```

### 4.2 BOM维护流程

```
选择产品 → 维护层级 → 添加物料 → 设置用量 → 保存BOM → 发布生效
```

---

## 5. 数据模型

### 5.1 物料

| 字段 | 类型 | 说明 |
|------|------|------|
| material_code | VARCHAR(50) | 物料编码 |
| material_name | VARCHAR(100) | 物料名称 |
| material_type | VARCHAR(30) | 物料类型 |
| category_id | BIGINT | 分类ID |
| unit | VARCHAR(20) | 单位 |
| spec | VARCHAR(200) | 规格 |
| safety_stock | DECIMAL(18,3) | 安全库存 |
| is_active | SMALLINT | 是否启用 |

### 5.2 BOM

| 字段 | 类型 | 说明 |
|------|------|------|
| bom_code | VARCHAR(50) | BOM编码 |
| product_id | BIGINT | 产品ID |
| version | VARCHAR(20) | 版本 |
| status | VARCHAR(20) | 状态 |

### 5.3 客户

| 字段 | 类型 | 说明 |
|------|------|------|
| customer_code | VARCHAR(50) | 客户编码 |
| customer_name | VARCHAR(200) | 客户名称 |
| customer_type | VARCHAR(20) | 类型 |
| credit_limit | DECIMAL(18,2) | 信用额度 |
| payment_terms | VARCHAR(50) | 付款条款 |

### 5.4 供应商

| 字段 | 类型 | 说明 |
|------|------|------|
| supplier_code | VARCHAR(50) | 供应商编码 |
| supplier_name | VARCHAR(200) | 供应商名称 |
| supplier_type | VARCHAR(20) | 类型 |
| contact_phone | VARCHAR(50) | 联系电话 |
| grade | VARCHAR(10) | 等级 |

---

## 6. API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /mdm/materials | 物料列表 |
| POST | /mdm/materials | 创建物料 |
| GET | /mdm/materials/:id | 物料详情 |
| PUT | /mdm/materials/:id | 更新物料 |
| GET | /mdm/bom | BOM列表 |
| GET | /mdm/bom/:id | BOM详情 |
| GET | /mdm/process-routes | 工艺路线列表 |
| GET | /mdm/customers | 客户列表 |
| POST | /mdm/customers | 创建客户 |
| GET | /mdm/customers/:id | 客户详情 |
| GET | /mdm/suppliers | 供应商列表 |
| POST | /mdm/suppliers | 创建供应商 |
| GET | /mdm/suppliers/:id | 供应商详情 |
| POST | /mdm/suppliers/:id/approve | 审批通过 |
| POST | /mdm/suppliers/:id/blacklist | 加入黑名单 |

---

## 7. 关联文档

- [MOM3.0_主设计文档](./MOM3.0_主设计文档.md) - 系统总览
- [