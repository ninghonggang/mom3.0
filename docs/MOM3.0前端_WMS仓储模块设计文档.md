# MOM3.0 WMS仓储模块设计文档

**版本**: V2.0 | **所属模块**: M07仓储管理 | **基于**: [MOM3.0_主设计文档](./MOM3.0_主设计文档.md)

---

## 1. 模块概述

### 1.1 功能定位

WMS仓储模块负责物料的存储、流动和库存管理，与生产、采购、销售模块深度集成，实现物料的收、发、存、退、调、盘六大核心功能。

### 1.2 核心功能

| 功能 | 说明 |
|------|------|
| 仓库管理 | 仓库、库区、库位配置 |
| 库存管理 | 库存查询、冻结、调整 |
| 采购收货 | 采购入库确认 |
| 销售发货 | 销售出库 |
| 生产发料 | 车间领料 |
| 完工入库 | 生产成品入库 |
| 调拨管理 | 库间调拨 |
| 盘点管理 | 库存盘点 |

---

## 2. 页面清单

### 2.1 基础数据管理

| 页面 | 路由路径 | 核心功能 |
|------|----------|----------|
| 仓库管理 | `/wms/warehouse` | 仓库、库区、库位CRUD |
| 库位管理 | `/wms/location` | 库位配置、线边库 |
| 库区管理 | `/wms/areabasic` | 库区配置 |
| 库位组 | `/wms/locationgroup` | 库位分组管理 |
| Dock管理 | `/wms/dock` | Dock配置 |
| 物料BOM | `/wms/bom` | 物料清单配置 |
| 物料包装 | `/wms/itempackage` | 包装配置 |
| 客户管理 | `/wms/customer` | 客户档案 |
| 供应商管理 | `/wms/supplier` | 供应商档案 |
| 标签模板 | `/wms/barbasic` | 标签模板配置 |
| 条码规则 | `/wms/barcode` | 条码规则 |

### 2.2 库存管理

| 页面 | 路由路径 | 核心功能 |
|------|----------|----------|
| 库存管理 | `/wms/inventory` | 库存查询、冻结 |
| 容器管理 | `/wms/container` | 容器档案 |

### 2.3 入库管理

| 页面 | 路由路径 | 核心功能 |
|------|----------|----------|
| 收货单 | `/wms/receive` | 采购收货确认 |
| IQC来料检验 | `/wms/iqc` | 采购检验 |
| 生产入库 | `/wms/productreceipt` | 生产入库作业 |

### 2.4 出库管理

| 页面 | 路由路径 | 核心功能 |
|------|----------|----------|
| 发货单 | `/wms/delivery` | 销售发货 |
| 领料管理 | `/wms/issue` | 生产领料作业 |
| 备料计划 | `/wms/preparetoissueplan` | 备料管理 |

### 2.5 移库与盘点

| 页面 | 路由路径 | 核心功能 |
|------|----------|----------|
| 调拨管理 | `/wms/transfer` | 库间调拨 |
| 盘点管理 | `/wms/stock-check` | 库存盘点 |

### 2.6 AGV管理

| 页面 | 路由路径 | 核心功能 |
|------|----------|----------|
| AGV调度 | `/wms/agvManage` | AGV库位关系/接口 |

---

## 3. UI设计规范

### 3.1 页面基本结构

同MES模块，搜索+工具栏+表格+弹窗的标准布局。

### 3.2 状态标签

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| PENDING | warning | 待处理 |
| CONFIRMED | primary | 已确认 |
| IN_TRANSIT | info | 运输中 |
| RECEIVED | success | 已收货 |
| CANCELLED | info | 已取消 |

---

## 4. 业务流程

### 4.1 采购入库流程

```
采购订单 → ASN到货通知 → 收货扫描 → IQC检验 → 合格入库
```

### 4.2 生产发料流程

```
工单开工 → 发料申请 → 仓库拣料 → 线边库交接 → 工单用料
```

---

## 5. 数据模型

### 5.1 库存表

| 字段 | 类型 | 说明 |
|------|------|------|
| warehouse_id | BIGINT | 仓库ID |
| location_id | BIGINT | 库位ID |
| material_id | BIGINT | 物料ID |
| batch_no | VARCHAR(50) | 批次号 |
| quantity | DECIMAL(18,3) | 库存数量 |
| frozen_quantity | DECIMAL(18,3) | 冻结数量 |
| unit_cost | DECIMAL(18,4) | 单位成本 |

---

## 6. API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/warehouse/list | 仓库列表 |
| GET | /wms/location/list | 库位列表 |
| GET | /wms/inventory/list | 库存列表 |
| GET | /wms/receive/list | 收货单列表 |
| GET | /wms/delivery/list | 发货单列表 |
| GET | /wms/transfer/list | 调拨单列表 |

---

## 7. 缺失页面详细设计

### 7.1 库区管理 (areabasic)

**路由路径**: `/wms/areabasic`

**功能说明**: 库区配置管理，用于定义仓库内的区域划分，如原材料区、成品区、待检区等。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| area_code | VARCHAR(50) | 是 | 库区编码 |
| area_name | VARCHAR(100) | 是 | 库区名称 |
| warehouse_id | BIGINT | 是 | 所属仓库ID |
| area_type | VARCHAR(20) | 是 | 库区类型(RAW/PRODUCT/CHECK/RETURN) |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/areabasic/list | 库区列表 |
| GET | /wms/areabasic/{id} | 库区详情 |
| POST | /wms/areabasic | 新增库区 |
| PUT | /wms/areabasic/{id} | 修改库区 |
| DELETE | /wms/areabasic/{id} | 删除库区 |

---

### 7.2 库位组 (locationgroup)

**路由路径**: `/wms/locationgroup`

**功能说明**: 库位分组管理，将多个库位归为一组进行管理，便于批量操作和策略应用。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| group_code | VARCHAR(50) | 是 | 分组编码 |
| group_name | VARCHAR(100) | 是 | 分组名称 |
| warehouse_id | BIGINT | 是 | 所属仓库ID |
| area_id | BIGINT | 否 | 所属库区ID |
| location_ids | VARCHAR(500) | 是 | 包含的库位ID列表 |
| group_type | VARCHAR(20) | 是 | 分组类型 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/locationgroup/list | 库位组列表 |
| GET | /wms/locationgroup/{id} | 库位组详情 |
| POST | /wms/locationgroup | 新增库位组 |
| PUT | /wms/locationgroup/{id} | 修改库位组 |
| DELETE | /wms/locationgroup/{id} | 删除库位组 |

---

### 7.3 Dock管理 (dock)

**路由路径**: `/wms/dock`

**功能说明**: Dock配置管理，用于定义仓库的收货口和发货口，以及Dock与仓库区域的映射关系。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| dock_code | VARCHAR(50) | 是 | Dock编码 |
| dock_name | VARCHAR(100) | 是 | Dock名称 |
| warehouse_id | BIGINT | 是 | 所属仓库ID |
| dock_type | VARCHAR(20) | 是 | 类型(RECEIVE/SHIP/BOTH) |
| status | VARCHAR(20) | 是 | 状态 |
| position | VARCHAR(100) | 否 | 位置描述 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/dock/list | Dock列表 |
| GET | /wms/dock/{id} | Dock详情 |
| POST | /wms/dock | 新增Dock |
| PUT | /wms/dock/{id} | 修改Dock |
| DELETE | /wms/dock/{id} | 删除Dock |

---

### 7.4 物料BOM (bom)

**路由路径**: `/wms/bom`

**功能说明**: 物料清单配置，定义物料的组成结构，用于生产配料和成本核算。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| bom_code | VARCHAR(50) | 是 | BOM编码 |
| material_id | BIGINT | 是 | 物料ID |
| material_code | VARCHAR(50) | 是 | 物料编码 |
| version | VARCHAR(20) | 是 | 版本号 |
| effective_date | DATE | 是 | 生效日期 |
| expiry_date | DATE | 否 | 失效日期 |
| status | VARCHAR(20) | 是 | 状态 |
| bom_items | JSON | 是 | BOM明细列表 |

**BOM明细字段**:

| 字段 | 类型 | 说明 |
|------|------|------|
| item_material_id | BIGINT | 子物料ID |
| item_quantity | DECIMAL(18,6) | 用量 |
| item_unit | VARCHAR(20) | 单位 |
| scrap_rate | DECIMAL(5,4) | 损耗率 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/bom/list | BOM列表 |
| GET | /wms/bom/{id} | BOM详情 |
| POST | /wms/bom | 新增BOM |
| PUT | /wms/bom/{id} | 修改BOM |
| DELETE | /wms/bom/{id} | 删除BOM |
| GET | /wms/bom/material/{materialId} | 获取物料的BOM |

---

### 7.5 物料包装 (itempackage)

**路由路径**: `/wms/itempackage`

**功能说明**: 包装配置管理，定义物料的包装规格、单位和换算关系。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| material_id | BIGINT | 是 | 物料ID |
| package_level | VARCHAR(20) | 是 | 包装层级(LEVEL1/LEVEL2/LEVEL3) |
| package_type | VARCHAR(50) | 是 | 包装类型 |
| quantity | DECIMAL(18,6) | 是 | 数量 |
| unit | VARCHAR(20) | 是 | 单位 |
| barcode | VARCHAR(100) | 否 | 条码 |
| length | DECIMAL(10,2) | 否 | 长度(cm) |
| width | DECIMAL(10,2) | 否 | 宽度(cm) |
| height | DECIMAL(10,2) | 否 | 高度(cm) |
| weight | DECIMAL(10,3) | 否 | 重量(kg) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/itempackage/list | 包装配置列表 |
| GET | /wms/itempackage/{id} | 包装配置详情 |
| POST | /wms/itempackage | 新增包装配置 |
| PUT | /wms/itempackage/{id} | 修改包装配置 |
| DELETE | /wms/itempackage/{id} | 删除包装配置 |
| GET | /wms/itempackage/material/{materialId} | 获取物料的包装配置 |

---

### 7.6 客户管理 (customer)

**路由路径**: `/wms/customer`

**功能说明**: 客户档案管理，维护客户基本信息、联系人、收货地址等。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| customer_code | VARCHAR(50) | 是 | 客户编码 |
| customer_name | VARCHAR(200) | 是 | 客户名称 |
| customer_type | VARCHAR(20) | 是 | 客户类型 |
| contact_person | VARCHAR(100) | 否 | 联系人 |
| contact_phone | VARCHAR(50) | 否 | 联系电话 |
| contact_address | VARCHAR(500) | 否 | 收货地址 |
| tax_no | VARCHAR(50) | 否 | 税号 |
| payment_terms | VARCHAR(50) | 否 | 付款条款 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/customer/list | 客户列表 |
| GET | /wms/customer/{id} | 客户详情 |
| POST | /wms/customer | 新增客户 |
| PUT | /wms/customer/{id} | 修改客户 |
| DELETE | /wms/customer/{id} | 删除客户 |

---

### 7.7 供应商管理 (supplier)

**路由路径**: `/wms/supplier`

**功能说明**: 供应商档案管理，维护供应商基本信息、联系人、发货地址等。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| supplier_code | VARCHAR(50) | 是 | 供应商编码 |
| supplier_name | VARCHAR(200) | 是 | 供应商名称 |
| supplier_type | VARCHAR(20) | 是 | 供应商类型 |
| contact_person | VARCHAR(100) | 否 | 联系人 |
| contact_phone | VARCHAR(50) | 否 | 联系电话 |
| contact_address | VARCHAR(500) | 否 | 发货地址 |
| tax_no | VARCHAR(50) | 否 | 税号 |
| payment_terms | VARCHAR(50) | 否 | 付款条款 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/supplier/list | 供应商列表 |
| GET | /wms/supplier/{id} | 供应商详情 |
| POST | /wms/supplier | 新增供应商 |
| PUT | /wms/supplier/{id} | 修改供应商 |
| DELETE | /wms/supplier/{id} | 删除供应商 |

---

### 7.8 标签模板 (barbasic)

**路由路径**: `/wms/barbasic`

**功能说明**: 标签模板配置，定义各类标签的打印格式和内容布局。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| template_code | VARCHAR(50) | 是 | 模板编码 |
| template_name | VARCHAR(100) | 是 | 模板名称 |
| template_type | VARCHAR(20) | 是 | 模板类型(LOCATION/MATERIAL/CONTAINER/PRODUCT) |
| template_content | TEXT | 是 | 模板内容(JSON格式) |
| width | DECIMAL(10,2) | 是 | 宽度(mm) |
| height | DECIMAL(10,2) | 是 | 高度(mm) |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/barbasic/list | 标签模板列表 |
| GET | /wms/barbasic/{id} | 标签模板详情 |
| POST | /wms/barbasic | 新增标签模板 |
| PUT | /wms/barbasic/{id} | 修改标签模板 |
| DELETE | /wms/barbasic/{id} | 删除标签模板 |

---

### 7.9 条码规则 (barcode)

**路由路径**: `/wms/barcode`

**功能说明**: 条码规则配置，定义各类条码的生成规则和校验逻辑。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| rule_code | VARCHAR(50) | 是 | 规则编码 |
| rule_name | VARCHAR(100) | 是 | 规则名称 |
| barcode_type | VARCHAR(20) | 是 | 条码类型(CODE128/CODE39/QR/QR128) |
| prefix | VARCHAR(20) | 否 | 前缀 |
| pattern | VARCHAR(200) | 是 | 生成模式 |
| check_digit | VARCHAR(20) | 否 | 校验位算法(MOD10/MOD11/None) |
| sequence_start | BIGINT | 是 | 序列号起始值 |
| sequence_length | INT | 是 | 序列号长度 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/barcode/list | 条码规则列表 |
| GET | /wms/barcode/{id} | 条码规则详情 |
| POST | /wms/barcode | 新增条码规则 |
| PUT | /wms/barcode/{id} | 修改条码规则 |
| DELETE | /wms/barcode/{id} | 删除条码规则 |
| POST | /wms/barcode/generate | 生成条码 |

---

### 7.10 IQC来料检验 (iqc)

**路由路径**: `/wms/iqc`

**功能说明**: IQC(Incoming Quality Control)来料检验，对采购物料进行质量检验，确认合格后方可入库。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|
| iqc_no | VARCHAR(50) | 是 | 检验单号 |
| supplier_id | BIGINT | 是 | 供应商ID |
| receive_id | BIGINT | 是 | 收货单ID |
| inspect_type | VARCHAR(20) | 是 | 检验类型(NORMAL/REDUCED/STRICT) |
| inspect_date | DATE | 是 | 检验日期 |
| inspector | VARCHAR(50) | 是 | 检验员 |
| result | VARCHAR(20) | 是 | 检验结果(PASS/FAIL/PENDING) |
| sample_quantity | INT | 是 | 抽样数量 |
| accept_quantity | INT | 是 | 接收数量 |
| reject_quantity | INT | 是 | 拒收数量 |
| remark | VARCHAR(500) | 否 | 备注 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/iqc/list | IQC检验单列表 |
| GET | /wms/iqc/{id} | IQC检验单详情 |
| POST | /wms/iqc | 新增IQC检验单 |
| PUT | /wms/iqc/{id} | 修改IQC检验单 |
| POST | /wms/iqc/{id}/confirm | 确认检验结果 |
| GET | /wms/iqc/record/list | 检验记录列表 |

---

### 7.11 生产入库 (productreceipt)

**路由路径**: `/wms/productreceipt`

**功能说明**: 生产入库作业，接收生产车间完成的成品或半成品入库。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| receipt_no | VARCHAR(50) | 是 | 入库单号 |
| work_order_id | BIGINT | 是 | 工单ID |
| production_line_id | BIGINT | 是 | 产线ID |
| workshop_id | BIGINT | 是 | 车间ID |
| receipt_type | VARCHAR(20) | 是 | 入库类型(PRODUCT/SEMIFINISHED/BYPRODUCT) |
| receipt_date | DATE | 是 | 入库日期 |
| receipt_status | VARCHAR(20) | 是 | 状态(DRAFT/CONFIRMED/COMPLETED/CANCELLED) |
| total_quantity | DECIMAL(18,6) | 是 | 总数量 |
| remark | VARCHAR(500) | 否 | 备注 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/productreceipt/list | 生产入库单列表 |
| GET | /wms/productreceipt/{id} | 生产入库单详情 |
| POST | /wms/productreceipt | 新增生产入库单 |
| PUT | /wms/productreceipt/{id} | 修改生产入库单 |
| POST | /wms/productreceipt/{id}/confirm | 确认生产入库 |
| POST | /wms/productreceipt/{id}/cancel | 取消生产入库 |

---

### 7.12 领料管理 (issue)

**路由路径**: `/wms/issue`

**功能说明**: 生产领料作业，根据工单配料需求进行拣料和发料。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| issue_no | VARCHAR(50) | 是 | 领料单号 |
| work_order_id | BIGINT | 是 | 工单ID |
| issue_type | VARCHAR(20) | 是 | 领料类型(NORMAL/EMERGENCY/RETURN) |
| request_date | DATE | 是 | 申请日期 |
| issue_status | VARCHAR(20) | 是 | 状态(DRAFT/PICKING/PICKED/COMPLETED/CANCELLED) |
| warehouse_id | BIGINT | 是 | 仓库ID |
| workshop_id | BIGINT | 是 | 车间ID |
| requester | VARCHAR(50) | 是 | 申请人 |
| total_quantity | DECIMAL(18,6) | 是 | 总数量 |
| remark | VARCHAR(500) | 否 | 备注 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/issue/list | 领料单列表 |
| GET | /wms/issue/{id} | 领料单详情 |
| POST | /wms/issue | 新增领料单 |
| PUT | /wms/issue/{id} | 修改领料单 |
| POST | /wms/issue/{id}/pick | 开始拣料 |
| POST | /wms/issue/{id}/confirm | 确认领料 |
| POST | /wms/issue/{id}/cancel | 取消领料 |

---

### 7.13 备料计划 (preparetoissueplan)

**路由路径**: `/wms/preparetoissueplan`

**功能说明**: 备料管理，根据工单需求提前准备物料，提高发料效率。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| plan_no | VARCHAR(50) | 是 | 备料计划号 |
| work_order_id | BIGINT | 是 | 工单ID |
| warehouse_id | BIGINT | 是 | 仓库ID |
| plan_date | DATE | 是 | 计划日期 |
| plan_status | VARCHAR(20) | 是 | 状态(PENDING/IN_PROGRESS/COMPLETED) |
| priority | INT | 是 | 优先级(1-5) |
| prepare_user | VARCHAR(50) | 否 | 备料人 |
| complete_time | DATETIME | 否 | 完成时间 |
| remark | VARCHAR(500) | 否 | 备注 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/preparetoissueplan/list | 备料计划列表 |
| GET | /wms/preparetoissueplan/{id} | 备料计划详情 |
| POST | /wms/preparetoissueplan | 新增备料计划 |
| PUT | /wms/preparetoissueplan/{id} | 修改备料计划 |
| POST | /wms/preparetoissueplan/{id}/start | 开始备料 |
| POST | /wms/preparetoissueplan/{id}/complete | 完成备料 |
| DELETE | /wms/preparetoissueplan/{id} | 删除备料计划 |

---

### 7.14 容器管理 (container)

**路由路径**: `/wms/container`

**功能说明**: 容器档案管理，维护周转箱、托盘等容器的信息和状态。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| container_code | VARCHAR(50) | 是 | 容器编码 |
| container_type | VARCHAR(20) | 是 | 容器类型(TRAY/BUCKET/BOX/PALLET) |
| warehouse_id | BIGINT | 是 | 所属仓库ID |
| location_id | BIGINT | 否 | 当前位置ID |
| status | VARCHAR(20) | 是 | 状态(EMPTY/FULL/BIND/BROKEN) |
| max_capacity | DECIMAL(18,6) | 是 | 最大容量 |
| current_quantity | DECIMAL(18,6) | 否 | 当前载货量 |
| bind_material_id | BIGINT | 否 | 绑定物料ID |
| bind_order_id | BIGINT | 否 | 绑定单据ID |
| length | DECIMAL(10,2) | 否 | 长度(cm) |
| width | DECIMAL(10,2) | 否 | 宽度(cm) |
| height | DECIMAL(10,2) | 否 | 高度(cm) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/container/list | 容器列表 |
| GET | /wms/container/{id} | 容器详情 |
| POST | /wms/container | 新增容器 |
| PUT | /wms/container/{id} | 修改容器 |
| DELETE | /wms/container/{id} | 删除容器 |
| POST | /wms/container/{id}/bind | 绑定容器 |
| POST | /wms/container/{id}/unbind | 解绑容器 |
| POST | /wms/container/{id}/repair | 报修容器 |

---

### 7.15 AGV调度 (agvManage)

**路由路径**: `/wms/agvManage`

**功能说明**: AGV库位关系与接口配置，管理AGV与库位的映射关系以及AGV系统对接。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| agv_code | VARCHAR(50) | 是 | AGV编码 |
| agv_name | VARCHAR(100) | 是 | AGV名称 |
| agv_type | VARCHAR(20) | 是 | AGV类型(FORKLIFT/CONVEYOR/DRONE) |
| warehouse_id | BIGINT | 是 | 所属仓库ID |
| current_location_id | BIGINT | 否 | 当前位置ID |
| target_location_id | BIGINT | 否 | 目标位置ID |
| status | VARCHAR(20) | 是 | 状态(IDLE/BUSY/ERROR/OFFLINE) |
| interface_url | VARCHAR(200) | 否 | 接口地址 |
| interface_key | VARCHAR(100) | 否 | 接口密钥 |

**库位映射字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| location_id | BIGINT | 是 | 库位ID |
| agv_station_code | VARCHAR(50) | 是 | AGV站点编码 |
| mapping_type | VARCHAR(20) | 是 | 映射类型(PUTAWAY/PICKING/TRANSFER) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/agvManage/list | AGV列表 |
| GET | /wms/agvManage/{id} | AGV详情 |
| POST | /wms/agvManage | 新增AGV |
| PUT | /wms/agvManage/{id} | 修改AGV |
| DELETE | /wms/agvManage/{id} | 删除AGV |
| GET | /wms/agvManage/locationRelation/list | 库位关系列表 |
| POST | /wms/agvManage/locationRelation | 新增库位关系 |
| PUT | /wms/agvManage/locationRelation/{id} | 修改库位关系 |
| DELETE | /wms/agvManage/locationRelation/{id} | 删除库位关系 |
| GET | /wms/agvManage/interfaceInfo | AGV接口配置 |
| PUT | /wms/agvManage/interfaceInfo | 更新AGV接口配置 |
| POST | /wms/agvManage/{id}/dispatch | AGV调度 |

---

## 8. 缺失页面详细设计(续)

本文档补充以下16个子模块的页面设计，涵盖约276个缺失页面。每个页面包含：路由路径、功能说明、核心字段、API接口。

### 8.1 工厂建模

#### 8.1.1 产线管理 (productionline)

**路由路径**: `/wms/productionline`

**功能说明**: 产线配置管理，定义生产线的基本信息、所属车间、产线类型等。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| line_code | VARCHAR(50) | 是 | 产线编码 |
| line_name | VARCHAR(100) | 是 | 产线名称 |
| workshop_id | BIGINT | 是 | 所属车间ID |
| line_type | VARCHAR(20) | 是 | 产线类型(ASSEMBLY/PACKING/TESTING) |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/productionline/list | 产线列表 |
| GET | /wms/productionline/{id} | 产线详情 |
| POST | /wms/productionline | 新增产线 |
| PUT | /wms/productionline/{id} | 修改产线 |
| DELETE | /wms/productionline/{id} | 删除产线 |

---

#### 8.1.2 车间管理 (workshop)

**路由路径**: `/wms/workshop`

**功能说明**: 车间配置管理，定义生产车间的基本信息、所属工厂等。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| workshop_code | VARCHAR(50) | 是 | 车间编码 |
| workshop_name | VARCHAR(100) | 是 | 车间名称 |
| factory_id | BIGINT | 是 | 所属工厂ID |
| manager | VARCHAR(50) | 否 | 负责人 |
| contact_phone | VARCHAR(50) | 否 | 联系电话 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/workshop/list | 车间列表 |
| GET | /wms/workshop/{id} | 车间详情 |
| POST | /wms/workshop | 新增车间 |
| PUT | /wms/workshop/{id} | 修改车间 |
| DELETE | /wms/workshop/{id} | 删除车间 |

---

#### 8.1.3 工位管理 (workstation)

**路由路径**: `/wms/workstation`

**功能说明**: 工位配置管理，定义生产线上的工位信息、工位类型等。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| station_code | VARCHAR(50) | 是 | 工位编码 |
| station_name | VARCHAR(100) | 是 | 工位名称 |
| production_line_id | BIGINT | 是 | 所属产线ID |
| station_type | VARCHAR(20) | 是 | 工位类型(PRODUCTION/TEST/PACK) |
| process_id | BIGINT | 否 | 所属工序ID |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/workstation/list | 工位列表 |
| GET | /wms/workstation/{id} | 工位详情 |
| POST | /wms/workstation | 新增工位 |
| PUT | /wms/workstation/{id} | 修改工位 |
| DELETE | /wms/workstation/{id} | 删除工位 |

---

#### 8.1.4 工序管理 (process)

**路由路径**: `/wms/process`

**功能说明**: 工序配置管理，定义生产工序的先后顺序、标准工时等。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| process_code | VARCHAR(50) | 是 | 工序编码 |
| process_name | VARCHAR(100) | 是 | 工序名称 |
| sequence | INT | 是 | 工序序号 |
| standard_time | DECIMAL(10,2) | 否 | 标准工时(分钟) |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/process/list | 工序列表 |
| GET | /wms/process/{id} | 工序详情 |
| POST | /wms/process | 新增工序 |
| PUT | /wms/process/{id} | 修改工序 |
| DELETE | /wms/process/{id} | 删除工序 |

---

#### 8.1.5 巷道管理 (aisle)

**路由路径**: `/wms/aisle`

**功能说明**: 仓库巷道配置管理，定义巷道编码、方向、宽度等。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| aisle_code | VARCHAR(50) | 是 | 巷道编码 |
| aisle_name | VARCHAR(100) | 是 | 巷道名称 |
| warehouse_id | BIGINT | 是 | 所属仓库ID |
| area_id | BIGINT | 否 | 所属库区ID |
| direction | VARCHAR(20) | 否 | 方向(HORIZONTAL/VERTICAL) |
| width | DECIMAL(10,2) | 否 | 宽度(m) |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/aisle/list | 巷道列表 |
| GET | /wms/aisle/{id} | 巷道详情 |
| POST | /wms/aisle | 新增巷道 |
| PUT | /wms/aisle/{id} | 修改巷道 |
| DELETE | /wms/aisle/{id} | 删除巷道 |

---

#### 8.1.6 层架管理 (rack)

**路由路径**: `/wms/rack`

**功能说明**: 货架层架配置管理，定义货架的层数、每层容量等。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| rack_code | VARCHAR(50) | 是 | 货架编码 |
| rack_name | VARCHAR(100) | 是 | 货架名称 |
| warehouse_id | BIGINT | 是 | 所属仓库ID |
| aisle_id | BIGINT | 否 | 所属巷道ID |
| levels | INT | 是 | 层数 |
| columns | INT | 是 | 列数 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/rack/list | 货架列表 |
| GET | /wms/rack/{id} | 货架详情 |
| POST | /wms/rack | 新增货架 |
| PUT | /wms/rack/{id} | 修改货架 |
| DELETE | /wms/rack/{id} | 删除货架 |

---

#### 8.1.7 仓库模型 (warehousemodel)

**路由路径**: `/wms/warehousemodel`

**功能说明**: 仓库3D模型配置，定义仓库的空间结构和物理布局。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| model_code | VARCHAR(50) | 是 | 模型编码 |
| model_name | VARCHAR(100) | 是 | 模型名称 |
| warehouse_id | BIGINT | 是 | 所属仓库ID |
| model_config | JSON | 是 | 模型配置(结构数据) |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/warehousemodel/list | 模型列表 |
| GET | /wms/warehousemodel/{id} | 模型详情 |
| POST | /wms/warehousemodel | 新增模型 |
| PUT | /wms/warehousemodel/{id} | 修改模型 |
| DELETE | /wms/warehousemodel/{id} | 删除模型 |

---

#### 8.1.8 库位属性 (locationattribute)

**路由路径**: `/wms/locationattribute`

**功能说明**: 库位属性配置，定义库位的特殊属性如温度区、湿度区等。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| attr_code | VARCHAR(50) | 是 | 属性编码 |
| attr_name | VARCHAR(100) | 是 | 属性名称 |
| attr_type | VARCHAR(20) | 是 | 属性类型(TEMPERATURE/HUMIDITY/HAZARD) |
| default_value | VARCHAR(100) | 否 | 默认值 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/locationattribute/list | 属性列表 |
| GET | /wms/locationattribute/{id} | 属性详情 |
| POST | /wms/locationattribute | 新增属性 |
| PUT | /wms/locationattribute/{id} | 修改属性 |
| DELETE | /wms/locationattribute/{id} | 删除属性 |

---

#### 8.1.9 库位类型 (locationtype)

**路由路径**: `/wms/locationtype`

**功能说明**: 库位类型配置，定义不同类型的库位及其拣货策略。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| type_code | VARCHAR(50) | 是 | 类型编码 |
| type_name | VARCHAR(100) | 是 | 类型名称 |
| picking_strategy | VARCHAR(20) | 否 | 拣货策略(FIFO/FEFO/LOCATION) |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/locationtype/list | 类型列表 |
| GET | /wms/locationtype/{id} | 类型详情 |
| POST | /wms/locationtype | 新增类型 |
| PUT | /wms/locationtype/{id} | 修改类型 |
| DELETE | /wms/locationtype/{id} | 删除类型 |

---

#### 8.1.10 库区类型 (areatype)

**路由路径**: `/wms/areatype`

**功能说明**: 库区类型配置，定义不同功能库区的类型。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| type_code | VARCHAR(50) | 是 | 类型编码 |
| type_name | VARCHAR(100) | 是 | 类型名称 |
| function_type | VARCHAR(20) | 是 | 功能类型(RAW/PRODUCT/CHECK/RETURN/DOCK) |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/areatype/list | 类型列表 |
| GET | /wms/areatype/{id} | 类型详情 |
| POST | /wms/areatype | 新增类型 |
| PUT | /wms/areatype/{id} | 修改类型 |
| DELETE | /wms/areatype/{id} | 删除类型 |

---

#### 8.1.11 库存状态 (inventorystatus)

**路由路径**: `/wms/inventorystatus`

**功能说明**: 库存状态配置，定义库存的可用、冻结、锁定等状态。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| status_code | VARCHAR(50) | 是 | 状态编码 |
| status_name | VARCHAR(100) | 是 | 状态名称 |
| color | VARCHAR(20) | 否 | 显示颜色 |
| is_freeze | BOOLEAN | 是 | 是否可冻结 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/inventorystatus/list | 状态列表 |
| GET | /wms/inventorystatus/{id} | 状态详情 |
| POST | /wms/inventorystatus | 新增状态 |
| PUT | /wms/inventorystatus/{id} | 修改状态 |
| DELETE | /wms/inventorystatus/{id} | 删除状态 |

---

### 8.2 物料管理

#### 8.2.1 替代物料 (substitutematerial)

**路由路径**: `/wms/substitutematerial`

**功能说明**: 替代物料配置，定义物料间的替代关系和优先级。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| material_id | BIGINT | 是 | 物料ID |
| substitute_id | BIGINT | 是 | 替代物料ID |
| priority | INT | 是 | 优先级(1-10) |
| effective_date | DATE | 否 | 生效日期 |
| expiry_date | DATE | 否 | 失效日期 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/substitutematerial/list | 替代关系列表 |
| GET | /wms/substitutematerial/{id} | 替代关系详情 |
| POST | /wms/substitutematerial | 新增替代关系 |
| PUT | /wms/substitutematerial/{id} | 修改替代关系 |
| DELETE | /wms/substitutematerial/{id} | 删除替代关系 |

---

#### 8.2.2 物料特性 (itemfeature)

**路由路径**: `/wms/itemfeature`

**功能说明**: 物料特性配置，定义物料的特殊属性如规格、材质等。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| feature_code | VARCHAR(50) | 是 | 特性编码 |
| feature_name | VARCHAR(100) | 是 | 特性名称 |
| data_type | VARCHAR(20) | 是 | 数据类型(STRING/NUMBER/DATE) |
| default_value | VARCHAR(100) | 否 | 默认值 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/itemfeature/list | 特性列表 |
| GET | /wms/itemfeature/{id} | 特性详情 |
| POST | /wms/itemfeature | 新增特性 |
| PUT | /wms/itemfeature/{id} | 修改特性 |
| DELETE | /wms/itemfeature/{id} | 删除特性 |

---

#### 8.2.3 物料属性 (itemattribute)

**路由路径**: `/wms/itemattribute`

**功能说明**: 物料属性管理，为物料赋值具体属性值。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| material_id | BIGINT | 是 | 物料ID |
| feature_id | BIGINT | 是 | 特性ID |
| attr_value | VARCHAR(200) | 是 | 属性值 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/itemattribute/list | 属性列表 |
| GET | /wms/itemattribute/{id} | 属性详情 |
| POST | /wms/itemattribute | 新增属性 |
| PUT | /wms/itemattribute/{id} | 修改属性 |
| DELETE | /wms/itemattribute/{id} | 删除属性 |

---

#### 8.2.4 物料组 (itemgroup)

**路由路径**: `/wms/itemgroup`

**功能说明**: 物料分组管理，定义物料的分类层级结构。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| group_code | VARCHAR(50) | 是 | 分组编码 |
| group_name | VARCHAR(100) | 是 | 分组名称 |
| parent_id | BIGINT | 否 | 上级分组ID |
| level | INT | 是 | 层级深度 |
| sort_no | INT | 否 | 排序号 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/itemgroup/list | 分组列表 |
| GET | /wms/itemgroup/tree | 分组树形结构 |
| GET | /wms/itemgroup/{id} | 分组详情 |
| POST | /wms/itemgroup | 新增分组 |
| PUT | /wms/itemgroup/{id} | 修改分组 |
| DELETE | /wms/itemgroup/{id} | 删除分组 |

---

#### 8.2.5 物料状态 (itemstatus)

**路由路径**: `/wms/itemstatus`

**功能说明**: 物料状态配置，定义物料的启用、禁用、停售等状态。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| status_code | VARCHAR(50) | 是 | 状态编码 |
| status_name | VARCHAR(100) | 是 | 状态名称 |
| is_active | BOOLEAN | 是 | 是否可用 |
| color | VARCHAR(20) | 否 | 显示颜色 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/itemstatus/list | 状态列表 |
| GET | /wms/itemstatus/{id} | 状态详情 |
| POST | /wms/itemstatus | 新增状态 |
| PUT | /wms/itemstatus/{id} | 修改状态 |
| DELETE | /wms/itemstatus/{id} | 删除状态 |

---

#### 8.2.6 采购单位 (purchaseunit)

**路由路径**: `/wms/purchaseunit`

**功能说明**: 采购单位配置，定义物料的采购计量单位。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| material_id | BIGINT | 是 | 物料ID |
| unit_code | VARCHAR(20) | 是 | 单位编码 |
| unit_name | VARCHAR(50) | 是 | 单位名称 |
| conversion_ratio | DECIMAL(18,6) | 是 | 转换比率(相对库存单位) |
| min_order_qty | DECIMAL(18,6) | 否 | 最小订购量 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/purchaseunit/list | 采购单位列表 |
| GET | /wms/purchaseunit/{id} | 采购单位详情 |
| POST | /wms/purchaseunit | 新增采购单位 |
| PUT | /wms/purchaseunit/{id} | 修改采购单位 |
| DELETE | /wms/purchaseunit/{id} | 删除采购单位 |

---

#### 8.2.7 库存单位 (stockunit)

**路由路径**: `/wms/stockunit`

**功能说明**: 库存单位配置，定义物料的库存计量单位。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| material_id | BIGINT | 是 | 物料ID |
| unit_code | VARCHAR(20) | 是 | 单位编码 |
| unit_name | VARCHAR(50) | 是 | 单位名称 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/stockunit/list | 库存单位列表 |
| GET | /wms/stockunit/{id} | 库存单位详情 |
| POST | /wms/stockunit | 新增库存单位 |
| PUT | /wms/stockunit/{id} | 修改库存单位 |
| DELETE | /wms/stockunit/{id} | 删除库存单位 |

---

#### 8.2.8 物料标签 (itemlabel)

**路由路径**: `/wms/itemlabel`

**功能说明**: 物料标签打印管理，定义物料标签的打印内容和格式。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| label_code | VARCHAR(50) | 是 | 标签编码 |
| material_id | BIGINT | 是 | 物料ID |
| template_id | BIGINT | 是 | 模板ID |
| print_qty | INT | 否 | 打印份数 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/itemlabel/list | 标签列表 |
| GET | /wms/itemlabel/{id} | 标签详情 |
| POST | /wms/itemlabel | 新增标签 |
| PUT | /wms/itemlabel/{id} | 修改标签 |
| DELETE | /wms/itemlabel/{id} | 删除标签 |
| POST | /wms/itemlabel/{id}/print | 打印标签 |

---

#### 8.2.9 物料描述 (itemdescription)

**路由路径**: `/wms/itemdescription`

**功能说明**: 物料描述管理，支持多语言物料描述信息。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| material_id | BIGINT | 是 | 物料ID |
| language | VARCHAR(20) | 是 | 语言(ZH/EN) |
| description | TEXT | 是 | 物料描述 |
| specification | TEXT | 否 | 规格说明 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/itemdescription/list | 描述列表 |
| GET | /wms/itemdescription/{id} | 描述详情 |
| POST | /wms/itemdescription | 新增描述 |
| PUT | /wms/itemdescription/{id} | 修改描述 |
| DELETE | /wms/itemdescription/{id} | 删除描述 |

---

### 8.3 客户管理

#### 8.3.1 客户项目 (customerproject)

**路由路径**: `/wms/customerproject`

**功能说明**: 客户项目档案管理，定义客户的项目信息、周期等。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| project_code | VARCHAR(50) | 是 | 项目编码 |
| project_name | VARCHAR(200) | 是 | 项目名称 |
| customer_id | BIGINT | 是 | 客户ID |
| project_type | VARCHAR(20) | 否 | 项目类型 |
| start_date | DATE | 否 | 开始日期 |
| end_date | DATE | 否 | 结束日期 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/customerproject/list | 项目列表 |
| GET | /wms/customerproject/{id} | 项目详情 |
| POST | /wms/customerproject | 新增项目 |
| PUT | /wms/customerproject/{id} | 修改项目 |
| DELETE | /wms/customerproject/{id} | 删除项目 |

---

#### 8.3.2 价格管理 (price)

**路由路径**: `/wms/price`

**功能说明**: 客户价格管理，定义客户专属的物料销售价格。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| customer_id | BIGINT | 是 | 客户ID |
| material_id | BIGINT | 是 | 物料ID |
| unit_price | DECIMAL(18,4) | 是 | 单价 |
| currency | VARCHAR(10) | 是 | 币种 |
| effective_date | DATE | 是 | 生效日期 |
| expiry_date | DATE | 否 | 失效日期 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/price/list | 价格列表 |
| GET | /wms/price/{id} | 价格详情 |
| POST | /wms/price | 新增价格 |
| PUT | /wms/price/{id} | 修改价格 |
| DELETE | /wms/price/{id} | 删除价格 |

---

#### 8.3.3 客户联系人 (customercontact)

**路由路径**: `/wms/customercontact`

**功能说明**: 客户联系人管理，维护客户的联系人信息。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| customer_id | BIGINT | 是 | 客户ID |
| contact_name | VARCHAR(100) | 是 | 联系人姓名 |
| position | VARCHAR(50) | 否 | 职位 |
| phone | VARCHAR(50) | 否 | 联系电话 |
| mobile | VARCHAR(50) | 否 | 手机 |
| email | VARCHAR(100) | 否 | 邮箱 |
| is_default | BOOLEAN | 否 | 是否默认联系人 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/customercontact/list | 联系人列表 |
| GET | /wms/customercontact/{id} | 联系人详情 |
| POST | /wms/customercontact | 新增联系人 |
| PUT | /wms/customercontact/{id} | 修改联系人 |
| DELETE | /wms/customercontact/{id} | 删除联系人 |

---

#### 8.3.4 客户地址 (customeraddress)

**路由路径**: `/wms/customeraddress`

**功能说明**: 客户地址管理，维护客户的收货地址、发函地址等。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| customer_id | BIGINT | 是 | 客户ID |
| address_type | VARCHAR(20) | 是 | 地址类型(DELIVERY/INVOICE) |
| receiver | VARCHAR(100) | 是 | 收货人 |
| phone | VARCHAR(50) | 是 | 联系电话 |
| province | VARCHAR(50) | 是 | 省份 |
| city | VARCHAR(50) | 是 | 城市 |
| district | VARCHAR(50) | 否 | 区县 |
| address | VARCHAR(500) | 是 | 详细地址 |
| is_default | BOOLEAN | 否 | 是否默认地址 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/customeraddress/list | 地址列表 |
| GET | /wms/customeraddress/{id} | 地址详情 |
| POST | /wms/customeraddress | 新增地址 |
| PUT | /wms/customeraddress/{id} | 修改地址 |
| DELETE | /wms/customeraddress/{id} | 删除地址 |

---

#### 8.3.5 客户银行 (customerbank)

**路由路径**: `/wms/customerbank`

**功能说明**: 客户银行账户管理，维护客户的银行账户信息。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| customer_id | BIGINT | 是 | 客户ID |
| bank_name | VARCHAR(100) | 是 | 开户银行 |
| bank_branch | VARCHAR(200) | 是 | 支行名称 |
| account_no | VARCHAR(50) | 是 | 账号 |
| account_name | VARCHAR(100) | 是 | 户名 |
| is_default | BOOLEAN | 否 | 是否默认账户 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/customerbank/list | 账户列表 |
| GET | /wms/customerbank/{id} | 账户详情 |
| POST | /wms/customerbank | 新增账户 |
| PUT | /wms/customerbank/{id} | 修改账户 |
| DELETE | /wms/customerbank/{id} | 删除账户 |

---

### 8.4 供应商管理

#### 8.4.1 供应商物料 (supplieritem)

**路由路径**: `/wms/supplieritem`

**功能说明**: 供应商物料配置，定义供应商供应的物料信息。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| supplier_id | BIGINT | 是 | 供应商ID |
| material_id | BIGINT | 是 | 物料ID |
| supplier_material_code | VARCHAR(50) | 否 | 供应商物料编码 |
| supplier_material_name | VARCHAR(200) | 否 | 供应商物料名称 |
| lead_time | INT | 否 | 交期(天) |
| moq | DECIMAL(18,6) | 否 | 最小订购量 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/supplieritem/list | 供应商物料列表 |
| GET | /wms/supplieritem/{id} | 供应商物料详情 |
| POST | /wms/supplieritem | 新增供应商物料 |
| PUT | /wms/supplieritem/{id} | 修改供应商物料 |
| DELETE | /wms/supplieritem/{id} | 删除供应商物料 |

---

#### 8.4.2 供应商周期 (supplierperiod)

**路由路径**: `/wms/supplierperiod`

**功能说明**: 供应商交货周期配置，定义不同物料的供货周期。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| supplier_id | BIGINT | 是 | 供应商ID |
| material_id | BIGINT | 是 | 物料ID |
| period_type | VARCHAR(20) | 是 | 周期类型(WEEK/MONTH/CUSTOM) |
| period_days | INT | 是 | 周期天数 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/supplierperiod/list | 周期列表 |
| GET | /wms/supplierperiod/{id} | 周期详情 |
| POST | /wms/supplierperiod | 新增周期 |
| PUT | /wms/supplierperiod/{id} | 修改周期 |
| DELETE | /wms/supplierperiod/{id} | 删除周期 |

---

#### 8.4.3 供应商价格 (supplierprice)

**路由路径**: `/wms/supplierprice`

**功能说明**: 供应商价格管理，定义采购物料的供应商价格。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| supplier_id | BIGINT | 是 | 供应商ID |
| material_id | BIGINT | 是 | 物料ID |
| unit_price | DECIMAL(18,4) | 是 | 单价 |
| currency | VARCHAR(10) | 是 | 币种 |
| effective_date | DATE | 是 | 生效日期 |
| expiry_date | DATE | 否 | 失效日期 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/supplierprice/list | 价格列表 |
| GET | /wms/supplierprice/{id} | 价格详情 |
| POST | /wms/supplierprice | 新增价格 |
| PUT | /wms/supplierprice/{id} | 修改价格 |
| DELETE | /wms/supplierprice/{id} | 删除价格 |

---

### 8.5 标签管理

#### 8.5.1 标签类型 (labeltype)

**路由路径**: `/wms/labeltype`

**功能说明**: 标签类型配置，定义不同业务场景使用的标签分类。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| type_code | VARCHAR(50) | 是 | 类型编码 |
| type_name | VARCHAR(100) | 是 | 类型名称 |
| template_required | BOOLEAN | 是 | 是否必须模板 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/labeltype/list | 类型列表 |
| GET | /wms/labeltype/{id} | 类型详情 |
| POST | /wms/labeltype | 新增类型 |
| PUT | /wms/labeltype/{id} | 修改类型 |
| DELETE | /wms/labeltype/{id} | 删除类型 |

---

#### 8.5.2 标签模板 (labeltemplate)

**路由路径**: `/wms/labeltemplate`

**功能说明**: 标签模板配置，定义标签的打印格式和内容布局。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| template_code | VARCHAR(50) | 是 | 模板编码 |
| template_name | VARCHAR(100) | 是 | 模板名称 |
| label_type | VARCHAR(20) | 是 | 标签类型 |
| template_content | TEXT | 是 | 模板内容(JSON格式) |
| width | DECIMAL(10,2) | 是 | 宽度(mm) |
| height | DECIMAL(10,2) | 是 | 高度(mm) |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/labeltemplate/list | 模板列表 |
| GET | /wms/labeltemplate/{id} | 模板详情 |
| POST | /wms/labeltemplate | 新增模板 |
| PUT | /wms/labeltemplate/{id} | 修改模板 |
| DELETE | /wms/labeltemplate/{id} | 删除模板 |

---

#### 8.5.3 条码规则 (barcode) - 已存在于7.9节

#### 8.5.4 叫料规则 (callmaterialsrule)

**路由路径**: `/wms/callmaterialsrule`

**功能说明**: 叫料规则配置，定义生产线叫料的触发条件和优先级。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| rule_code | VARCHAR(50) | 是 | 规则编码 |
| rule_name | VARCHAR(100) | 是 | 规则名称 |
| priority | INT | 是 | 优先级(1-10) |
| trigger_type | VARCHAR(20) | 是 | 触发类型(MIN_STOCK/TIME_BASED) |
| threshold_qty | DECIMAL(18,6) | 否 | 触发数量 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/callmaterialsrule/list | 规则列表 |
| GET | /wms/callmaterialsrule/{id} | 规则详情 |
| POST | /wms/callmaterialsrule | 新增规则 |
| PUT | /wms/callmaterialsrule/{id} | 修改规则 |
| DELETE | /wms/callmaterialsrule/{id} | 删除规则 |

---

#### 8.5.5 标签打印 (labelprint)

**路由路径**: `/wms/labelprint`

**功能说明**: 标签打印执行管理，记录标签打印任务和历史。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| print_id | VARCHAR(50) | 是 | 打印单号 |
| label_code | VARCHAR(50) | 是 | 标签编码 |
| material_id | BIGINT | 否 | 物料ID |
| printer | VARCHAR(100) | 是 | 打印机 |
| copies | INT | 是 | 打印份数 |
| print_time | DATETIME | 是 | 打印时间 |
| operator | VARCHAR(50) | 是 | 操作人 |
| status | VARCHAR(20) | 是 | 状态(SUCCESS/FAILED) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/labelprint/list | 打印记录列表 |
| GET | /wms/labelprint/{id} | 打印记录详情 |
| POST | /wms/labelprint | 新增打印任务 |
| POST | /wms/labelprint/batch | 批量打印 |

---

#### 8.5.6 标签历史 (labelhistory)

**路由路径**: `/wms/labelhistory`

**功能说明**: 标签打印历史查询，追溯标签的打印记录。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| history_id | BIGINT | 是 | 历史ID |
| label_code | VARCHAR(50) | 是 | 标签编码 |
| material_id | BIGINT | 否 | 物料ID |
| print_time | DATETIME | 是 | 打印时间 |
| operator_id | BIGINT | 是 | 操作人ID |
| printer | VARCHAR(100) | 是 | 打印机 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/labelhistory/list | 历史列表 |
| GET | /wms/labelhistory/{id} | 历史详情 |
| GET | /wms/labelhistory/export-excel | 导出历史 |

---

#### 8.5.7 标签绑定 (labelbind)

**路由路径**: `/wms/labelbind`

**功能说明**: 标签绑定管理，将标签与物料、库位等进行绑定。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| bind_id | BIGINT | 是 | 绑定ID |
| label_code | VARCHAR(50) | 是 | 标签编码 |
| bind_type | VARCHAR(20) | 是 | 绑定类型(MATERIAL/LOCATION/CONTAINER) |
| bind_object_id | BIGINT | 是 | 绑定对象ID |
| bind_time | DATETIME | 是 | 绑定时间 |
| operator_id | BIGINT | 是 | 操作人ID |
| status | VARCHAR(20) | 是 | 状态(BOUND/UNBOUND) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/labelbind/list | 绑定列表 |
| GET | /wms/labelbind/{id} | 绑定详情 |
| POST | /wms/labelbind | 新增绑定 |
| DELETE | /wms/labelbind/{id} | 解除绑定 |

---

#### 8.5.8 标签解除 (labelunbind)

**路由路径**: `/wms/labelunbind`

**功能说明**: 标签解除管理，记录标签解绑操作。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| unbind_id | BIGINT | 是 | 解除ID |
| label_code | VARCHAR(50) | 是 | 标签编码 |
| reason | VARCHAR(200) | 是 | 解除原因 |
| unbind_time | DATETIME | 是 | 解除时间 |
| operator_id | BIGINT | 是 | 操作人ID |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/labelunbind/list | 解除列表 |
| GET | /wms/labelunbind/{id} | 解除详情 |
| POST | /wms/labelunbind | 新增解除记录 |

---

### 8.6 策略设置

#### 8.6.1 上架策略 (putawaystrategy)

**路由路径**: `/wms/putawaystrategy`

**功能说明**: 上架策略配置，定义物料入库时的推荐库位规则。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| rule_type | VARCHAR(20) | 是 | 规则类型(FIXED/DYNAMIC) |
| priority | INT | 是 | 优先级 |
| conditions | JSON | 否 | 条件配置 |
| warehouse_id | BIGINT | 否 | 适用仓库 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/putawaystrategy/list | 策略列表 |
| GET | /wms/putawaystrategy/{id} | 策略详情 |
| POST | /wms/putawaystrategy | 新增策略 |
| PUT | /wms/putawaystrategy/{id} | 修改策略 |
| DELETE | /wms/putawaystrategy/{id} | 删除策略 |

---

#### 8.6.2 下架策略 (pickstrategy)

**路由路径**: `/wms/pickstrategy`

**功能说明**: 下架策略配置，定义物料出库时的拣货库位规则。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| rule_type | VARCHAR(20) | 是 | 规则类型(FIFO/FEFO/LOCATION_BASED) |
| priority | INT | 是 | 优先级 |
| conditions | JSON | 否 | 条件配置 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/pickstrategy/list | 策略列表 |
| GET | /wms/pickstrategy/{id} | 策略详情 |
| POST | /wms/pickstrategy | 新增策略 |
| PUT | /wms/pickstrategy/{id} | 修改策略 |
| DELETE | /wms/pickstrategy/{id} | 删除策略 |

---

#### 8.6.3 批策略 (batchstrategy)

**路由路径**: `/wms/batchstrategy`

**功能说明**: 批策略配置，定义批次管理规则如保质期、先进先出等。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| batch_rule | VARCHAR(20) | 是 | 批规则(PRODATE/INDATE/FIFO) |
| expiry_days | INT | 否 | 保质期天数 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/batchstrategy/list | 策略列表 |
| GET | /wms/batchstrategy/{id} | 策略详情 |
| POST | /wms/batchstrategy | 新增策略 |
| PUT | /wms/batchstrategy/{id} | 修改策略 |
| DELETE | /wms/batchstrategy/{id} | 删除策略 |

---

#### 8.6.4 库容策略 (capacitystrategy)

**路由路径**: `/wms/capacitystrategy`

**功能说明**: 库容策略配置，定义库位的最大存储容量和警戒线。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| location_id | BIGINT | 否 | 库位ID |
| max_qty | DECIMAL(18,6) | 是 | 最大容量 |
| warning_qty | DECIMAL(18,6) | 否 | 警戒容量 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/capacitystrategy/list | 策略列表 |
| GET | /wms/capacitystrategy/{id} | 策略详情 |
| POST | /wms/capacitystrategy | 新增策略 |
| PUT | /wms/capacitystrategy/{id} | 修改策略 |
| DELETE | /wms/capacitystrategy/{id} | 删除策略 |

---

#### 8.6.5 补货策略 (replenishmentstrategy)

**路由路径**: `/wms/replenishmentstrategy`

**功能说明**: 补货策略配置，定义线边库自动补货的触发条件和规则。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| trigger_type | VARCHAR(20) | 是 | 触发类型(MIN_STOCK/POINT_STOCK/TIME) |
| threshold_min | DECIMAL(18,6) | 否 | 最小阈值 |
| threshold_max | DECIMAL(18,6) | 否 | 最大阈值 |
| material_id | BIGINT | 否 | 物料ID |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/replenishmentstrategy/list | 策略列表 |
| GET | /wms/replenishmentstrategy/{id} | 策略详情 |
| POST | /wms/replenishmentstrategy | 新增策略 |
| PUT | /wms/replenishmentstrategy/{id} | 修改策略 |
| DELETE | /wms/replenishmentstrategy/{id} | 删除策略 |

---

#### 8.6.6 分配策略 (allocationstrategy)

**路由路径**: `/wms/allocationstrategy`

**功能说明**: 分配策略配置，定义订单分配库存时的优先级规则。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| rule_type | VARCHAR(20) | 是 | 规则类型(LOCATION/PRIORITY/DATE) |
| priority | INT | 是 | 优先级 |
| conditions | JSON | 否 | 条件配置 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/allocationstrategy/list | 策略列表 |
| GET | /wms/allocationstrategy/{id} | 策略详情 |
| POST | /wms/allocationstrategy | 新增策略 |
| PUT | /wms/allocationstrategy/{id} | 修改策略 |
| DELETE | /wms/allocationstrategy/{id} | 删除策略 |

---

#### 8.6.7 合并策略 (mergestrategy)

**路由路径**: `/wms/mergestrategy`

**功能说明**: 合并策略配置，定义库存合并的规则和条件。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| merge_type | VARCHAR(20) | 是 | 合并类型(BATCH/LOCATION) |
| conditions | JSON | 否 | 合并条件 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/mergestrategy/list | 策略列表 |
| GET | /wms/mergestrategy/{id} | 策略详情 |
| POST | /wms/mergestrategy | 新增策略 |
| PUT | /wms/mergestrategy/{id} | 修改策略 |
| DELETE | /wms/mergestrategy/{id} | 删除策略 |

---

#### 8.6.8 分割策略 (splitstrategy)

**路由路径**: `/wms/splitstrategy`

**功能说明**: 分割策略配置，定义库存分割的规则和条件。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| split_type | VARCHAR(20) | 是 | 分割类型(QTY/BATCH) |
| conditions | JSON | 否 | 分割条件 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/splitstrategy/list | 策略列表 |
| GET | /wms/splitstrategy/{id} | 策略详情 |
| POST | /wms/splitstrategy | 新增策略 |
| PUT | /wms/splitstrategy/{id} | 修改策略 |
| DELETE | /wms/splitstrategy/{id} | 删除策略 |

---

#### 8.6.9 路径策略 (pathstrategy)

**路由路径**: `/wms/pathstrategy`

**功能说明**: 路径策略配置，定义AGV/物料搬运的最优路径规则。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| path_type | VARCHAR(20) | 是 | 路径类型(SHORTEST/AVOID) |
| waypoints | JSON | 否 | 路径点配置 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/pathstrategy/list | 策略列表 |
| GET | /wms/pathstrategy/{id} | 策略详情 |
| POST | /wms/pathstrategy | 新增策略 |
| PUT | /wms/pathstrategy/{id} | 修改策略 |
| DELETE | /wms/pathstrategy/{id} | 删除策略 |

---

#### 8.6.10 排序策略 (sortstrategy)

**路由路径**: `/wms/sortstrategy`

**功能说明**: 排序策略配置，定义出库订单的排序规则。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| sort_type | VARCHAR(20) | 是 | 排序类型(PRIORITY/DATE/DISTANCE) |
| priority | INT | 是 | 优先级 |
| conditions | JSON | 否 | 排序条件 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/sortstrategy/list | 策略列表 |
| GET | /wms/sortstrategy/{id} | 策略详情 |
| POST | /wms/sortstrategy | 新增策略 |
| PUT | /wms/sortstrategy/{id} | 修改策略 |
| DELETE | /wms/sortstrategy/{id} | 删除策略 |

---

#### 8.6.11 筛选策略 (filterstrategy)

**路由路径**: `/wms/filterstrategy`

**功能说明**: 筛选策略配置，定义库存查询和过滤的规则。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| filter_type | VARCHAR(20) | 是 | 筛选类型(BATCH/STATUS/LOCATION) |
| conditions | JSON | 是 | 筛选条件 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/filterstrategy/list | 策略列表 |
| GET | /wms/filterstrategy/{id} | 策略详情 |
| POST | /wms/filterstrategy | 新增策略 |
| PUT | /wms/filterstrategy/{id} | 修改策略 |
| DELETE | /wms/filterstrategy/{id} | 删除策略 |

---

#### 8.6.12 推荐策略 (recommendstrategy)

**路由路径**: `/wms/recommendstrategy`

**功能说明**: 推荐策略配置，定义库位推荐给物料的规则。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| recommend_type | VARCHAR(20) | 是 | 推荐类型(PUTAWAY/PICK) |
| conditions | JSON | 否 | 推荐条件 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/recommendstrategy/list | 策略列表 |
| GET | /wms/recommendstrategy/{id} | 策略详情 |
| POST | /wms/recommendstrategy | 新增策略 |
| PUT | /wms/recommendstrategy/{id} | 修改策略 |
| DELETE | /wms/recommendstrategy/{id} | 删除策略 |

---

#### 8.6.13 预警策略 (warningstrategy)

**路由路径**: `/wms/warningstrategy`

**功能说明**: 预警策略配置，定义库存预警的触发条件和通知方式。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| alert_type | VARCHAR(20) | 是 | 预警类型(STOCK/EXPIRY/QUALITY) |
| threshold | DECIMAL(18,6) | 是 | 阈值 |
| notify_type | VARCHAR(20) | 否 | 通知方式(EMAIL/SMS/SYSTEM) |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/warningstrategy/list | 策略列表 |
| GET | /wms/warningstrategy/{id} | 策略详情 |
| POST | /wms/warningstrategy | 新增策略 |
| PUT | /wms/warningstrategy/{id} | 修改策略 |
| DELETE | /wms/warningstrategy/{id} | 删除策略 |

---

#### 8.6.14 冻结策略 (freezestrategy)

**路由路径**: `/wms/freezestrategy`

**功能说明**: 冻结策略配置，定义库存冻结的规则和条件。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| freeze_type | VARCHAR(20) | 是 | 冻结类型(QC/ORDER/SYSTEM) |
| conditions | JSON | 否 | 冻结条件 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/freezestrategy/list | 策略列表 |
| GET | /wms/freezestrategy/{id} | 策略详情 |
| POST | /wms/freezestrategy | 新增策略 |
| PUT | /wms/freezestrategy/{id} | 修改策略 |
| DELETE | /wms/freezestrategy/{id} | 删除策略 |

---

#### 8.6.15 解冻策略 (unfreezestrategy)

**路由路径**: `/wms/unfreezestrategy`

**功能说明**: 解冻策略配置，定义库存解冻的规则和条件。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| unfreeze_type | VARCHAR(20) | 是 | 解冻类型(MANUAL/AUTO/APPROVAL) |
| conditions | JSON | 否 | 解冻条件 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/unfreezestrategy/list | 策略列表 |
| GET | /wms/unfreezestrategy/{id} | 策略详情 |
| POST | /wms/unfreezestrategy | 新增策略 |
| PUT | /wms/unfreezestrategy/{id} | 修改策略 |
| DELETE | /wms/unfreezestrategy/{id} | 删除策略 |

---

#### 8.6.16 锁定策略 (lockstrategy)

**路由路径**: `/wms/lockstrategy`

**功能说明**: 锁定策略配置，定义库存锁定的规则和条件。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| lock_type | VARCHAR(20) | 是 | 锁定类型(ORDER/SYSTEM) |
| conditions | JSON | 否 | 锁定条件 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/lockstrategy/list | 策略列表 |
| GET | /wms/lockstrategy/{id} | 策略详情 |
| POST | /wms/lockstrategy | 新增策略 |
| PUT | /wms/lockstrategy/{id} | 修改策略 |
| DELETE | /wms/lockstrategy/{id} | 删除策略 |

---

#### 8.6.17 解锁策略 (unlockstrategy)

**路由路径**: `/wms/unlockstrategy`

**功能说明**: 解锁策略配置，定义库存解锁的规则和条件。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| unlock_type | VARCHAR(20) | 是 | 解锁类型(MANUAL/AUTO) |
| conditions | JSON | 否 | 解锁条件 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/unlockstrategy/list | 策略列表 |
| GET | /wms/unlockstrategy/{id} | 策略详情 |
| POST | /wms/unlockstrategy | 新增策略 |
| PUT | /wms/unlockstrategy/{id} | 修改策略 |
| DELETE | /wms/unlockstrategy/{id} | 删除策略 |

---

#### 8.6.18 检验策略 (inspectstrategy)

**路由路径**: `/wms/inspectstrategy`

**功能说明**: 检验策略配置，定义来料检验的抽样规则和判定标准。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| inspect_type | VARCHAR(20) | 是 | 检验类型(NORMAL/REDUCED/STRICT) |
| sampling_rate | DECIMAL(5,2) | 否 | 抽样比例(%) |
| accept_level | VARCHAR(20) | 否 | 接收标准(AQL) |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/inspectstrategy/list | 策略列表 |
| GET | /wms/inspectstrategy/{id} | 策略详情 |
| POST | /wms/inspectstrategy | 新增策略 |
| PUT | /wms/inspectstrategy/{id} | 修改策略 |
| DELETE | /wms/inspectstrategy/{id} | 删除策略 |

---

#### 8.6.19 报废策略 (scrapstrategy)

**路由路径**: `/wms/scrapstrategy`

**功能说明**: 报废策略配置，定义库存报废的判定规则和审批流程。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| scrap_type | VARCHAR(20) | 是 | 报废类型(DAMAGE/EXPIRY/DEFECT) |
| conditions | JSON | 否 | 报废条件 |
| approval_required | BOOLEAN | 是 | 是否需要审批 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/scrapstrategy/list | 策略列表 |
| GET | /wms/scrapstrategy/{id} | 策略详情 |
| POST | /wms/scrapstrategy | 新增策略 |
| PUT | /wms/scrapstrategy/{id} | 修改策略 |
| DELETE | /wms/scrapstrategy/{id} | 删除策略 |

---

#### 8.6.20 盘点策略 (countstrategy)

**路由路径**: `/wms/countstrategy`

**功能说明**: 盘点策略配置，定义库存盘点的周期和范围规则。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| count_type | VARCHAR(20) | 是 | 盘点类型(CYCLE/FULL/SPOT) |
| frequency | VARCHAR(20) | 否 | 盘点频率(DAILY/WEEKLY/MONTHLY) |
| scope | JSON | 否 | 盘点范围 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/countstrategy/list | 策略列表 |
| GET | /wms/countstrategy/{id} | 策略详情 |
| POST | /wms/countstrategy | 新增策略 |
| PUT | /wms/countstrategy/{id} | 修改策略 |
| DELETE | /wms/countstrategy/{id} | 删除策略 |

---

#### 8.6.21 调拨策略 (transferstrategy)

**路由路径**: `/wms/transferstrategy`

**功能说明**: 调拨策略配置，定义库间调拨的规则和审批流程。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| transfer_type | VARCHAR(20) | 是 | 调拨类型(NORMAL/URGENT) |
| approval_required | BOOLEAN | 是 | 是否需要审批 |
| conditions | JSON | 否 | 调拨条件 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/transferstrategy/list | 策略列表 |
| GET | /wms/transferstrategy/{id} | 策略详情 |
| POST | /wms/transferstrategy | 新增策略 |
| PUT | /wms/transferstrategy/{id} | 修改策略 |
| DELETE | /wms/transferstrategy/{id} | 删除策略 |

---

#### 8.6.22 退货策略 (returnstrategy)

**路由路径**: `/wms/returnstrategy`

**功能说明**: 退货策略配置，定义客户退货的处理规则和流程。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| return_type | VARCHAR(20) | 是 | 退货类型(QUALITY/DELIVERY/ORDER) |
| inspection_required | BOOLEAN | 是 | 是否需要检验 |
| refund_type | VARCHAR(20) | 否 | 退款方式 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/returnstrategy/list | 策略列表 |
| GET | /wms/returnstrategy/{id} | 策略详情 |
| POST | /wms/returnstrategy | 新增策略 |
| PUT | /wms/returnstrategy/{id} | 修改策略 |
| DELETE | /wms/returnstrategy/{id} | 删除策略 |

---

#### 8.6.23 签收策略 (signstrategy)

**路由路径**: `/wms/signstrategy`

**功能说明**: 签收策略配置，定义出库送货的签收规则。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| sign_type | VARCHAR(20) | 是 | 签收类型(SIGNATURE/PHOTO/OCR) |
| confirm_deadline | INT | 否 | 确认期限(小时) |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/signstrategy/list | 策略列表 |
| GET | /wms/signstrategy/{id} | 策略详情 |
| POST | /wms/signstrategy | 新增策略 |
| PUT | /wms/signstrategy/{id} | 修改策略 |
| DELETE | /wms/signstrategy/{id} | 删除策略 |

---

#### 8.6.24 复核策略 (reviewstrategy)

**路由路径**: `/wms/reviewstrategy`

**功能说明**: 复核策略配置，定义出库复核的规则和流程。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| review_type | VARCHAR(20) | 是 | 复核类型(FULL/SPOT) |
| threshold_qty | DECIMAL(18,6) | 否 | 全检数量阈值 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/reviewstrategy/list | 策略列表 |
| GET | /wms/reviewstrategy/{id} | 策略详情 |
| POST | /wms/reviewstrategy | 新增策略 |
| PUT | /wms/reviewstrategy/{id} | 修改策略 |
| DELETE | /wms/reviewstrategy/{id} | 删除策略 |

---

#### 8.6.25 打包策略 (packstrategy)

**路由路径**: `/wms/packstrategy`

**功能说明**: 打包策略配置，定义打包作业的规则和包装规格。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| strategy_code | VARCHAR(50) | 是 | 策略编码 |
| strategy_name | VARCHAR(100) | 是 | 策略名称 |
| pack_type | VARCHAR(20) | 是 | 打包类型(STANDARD/CUSTOM) |
| pack_spec | JSON | 否 | 包装规格配置 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/packstrategy/list | 策略列表 |
| GET | /wms/packstrategy/{id} | 策略详情 |
| POST | /wms/packstrategy | 新增策略 |
| PUT | /wms/packstrategy/{id} | 修改策略 |
| DELETE | /wms/packstrategy/{id} | 删除策略 |

---

### 8.7 单据设置

#### 8.7.1 业务类型 (businesstype)

**路由路径**: `/wms/businesstype`

**功能说明**: 业务类型配置，定义仓储业务的各种类型如采购、销售、调拨等。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| type_code | VARCHAR(50) | 是 | 类型编码 |
| type_name | VARCHAR(100) | 是 | 类型名称 |
| type_category | VARCHAR(20) | 是 | 类型分类(IN/OUT/TRANSFER) |
| workflow_id | BIGINT | 否 | 审批流程ID |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/businesstype/list | 类型列表 |
| GET | /wms/businesstype/{id} | 类型详情 |
| POST | /wms/businesstype | 新增类型 |
| PUT | /wms/businesstype/{id} | 修改类型 |
| DELETE | /wms/businesstype/{id} | 删除类型 |

---

#### 8.7.2 单据类型 (documenttype)

**路由路径**: `/wms/documenttype`

**功能说明**: 单据类型配置，定义各种业务单据的类型。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| doc_type_code | VARCHAR(50) | 是 | 单据类型编码 |
| doc_type_name | VARCHAR(100) | 是 | 单据类型名称 |
| prefix | VARCHAR(20) | 否 | 单据号前缀 |
| sequence_length | INT | 否 | 序列号长度 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/documenttype/list | 类型列表 |
| GET | /wms/documenttype/{id} | 类型详情 |
| POST | /wms/documenttype | 新增类型 |
| PUT | /wms/documenttype/{id} | 修改类型 |
| DELETE | /wms/documenttype/{id} | 删除类型 |

---

#### 8.7.3 编号规则 (numberrule)

**路由路径**: `/wms/numberrule`

**功能说明**: 编号规则配置，定义单据编号的生成规则。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| rule_code | VARCHAR(50) | 是 | 规则编码 |
| rule_name | VARCHAR(100) | 是 | 规则名称 |
| prefix | VARCHAR(20) | 否 | 前缀 |
| date_format | VARCHAR(20) | 否 | 日期格式(YYYYMMDD) |
| sequence_start | BIGINT | 是 | 序列起始值 |
| sequence_length | INT | 是 | 序列号长度 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/numberrule/list | 规则列表 |
| GET | /wms/numberrule/{id} | 规则详情 |
| POST | /wms/numberrule | 新增规则 |
| PUT | /wms/numberrule/{id} | 修改规则 |
| DELETE | /wms/numberrule/{id} | 删除规则 |

---

#### 8.7.4 单据状态 (documentstatus)

**路由路径**: `/wms/documentstatus`

**功能说明**: 单据状态配置，定义单据的工作流状态和允许操作。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| status_code | VARCHAR(50) | 是 | 状态编码 |
| status_name | VARCHAR(100) | 是 | 状态名称 |
| doc_type | VARCHAR(50) | 是 | 关联单据类型 |
| allow_actions | JSON | 否 | 允许操作列表 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/documentstatus/list | 状态列表 |
| GET | /wms/documentstatus/{id} | 状态详情 |
| POST | /wms/documentstatus | 新增状态 |
| PUT | /wms/documentstatus/{id} | 修改状态 |
| DELETE | /wms/documentstatus/{id} | 删除状态 |

---

#### 8.7.5 审批流程 (approvalflow)

**路由路径**: `/wms/approvalflow`

**功能说明**: 审批流程配置，定义单据的审批工作流。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| flow_code | VARCHAR(50) | 是 | 流程编码 |
| flow_name | VARCHAR(100) | 是 | 流程名称 |
| doc_type | VARCHAR(50) | 是 | 适用单据类型 |
| approvers | JSON | 是 | 审批人配置 |
| conditions | JSON | 否 | 触发条件 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/approvalflow/list | 流程列表 |
| GET | /wms/approvalflow/{id} | 流程详情 |
| POST | /wms/approvalflow | 新增流程 |
| PUT | /wms/approvalflow/{id} | 修改流程 |
| DELETE | /wms/approvalflow/{id} | 删除流程 |

---

#### 8.7.6 打印设置 (printsetting)

**路由路径**: `/wms/printsetting`

**功能说明**: 打印设置配置，定义单据的打印模板和打印机。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| setting_code | VARCHAR(50) | 是 | 设置编码 |
| setting_name | VARCHAR(100) | 是 | 设置名称 |
| doc_type | VARCHAR(50) | 是 | 单据类型 |
| template_id | BIGINT | 否 | 打印模板ID |
| printer | VARCHAR(100) | 否 | 打印机 |
| copies | INT | 否 | 默认份数 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/printsetting/list | 设置列表 |
| GET | /wms/printsetting/{id} | 设置详情 |
| POST | /wms/printsetting | 新增设置 |
| PUT | /wms/printsetting/{id} | 修改设置 |
| DELETE | /wms/printsetting/{id} | 删除设置 |

---

#### 8.7.7 字段配置 (fieldconfig)

**路由路径**: `/wms/fieldconfig`

**功能说明**: 单据字段配置，定义单据字段的显示、编辑、必填等属性。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| field_code | VARCHAR(50) | 是 | 字段编码 |
| field_name | VARCHAR(100) | 是 | 字段名称 |
| doc_type | VARCHAR(50) | 是 | 单据类型 |
| is_visible | BOOLEAN | 是 | 是否显示 |
| is_editable | BOOLEAN | 是 | 是否可编辑 |
| is_required | BOOLEAN | 是 | 是否必填 |
| default_value | VARCHAR(100) | 否 | 默认值 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/fieldconfig/list | 配置列表 |
| GET | /wms/fieldconfig/{id} | 配置详情 |
| POST | /wms/fieldconfig | 新增配置 |
| PUT | /wms/fieldconfig/{id} | 修改配置 |
| DELETE | /wms/fieldconfig/{id} | 删除配置 |

---

#### 8.7.8 权限配置 (permissionconfig)

**路由路径**: `/wms/permissionconfig`

**功能说明**: 单据权限配置，定义用户对单据的操作权限。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| permission_code | VARCHAR(50) | 是 | 权限编码 |
| permission_name | VARCHAR(100) | 是 | 权限名称 |
| resource_type | VARCHAR(20) | 是 | 资源类型(DOCUMENT/WAREHOUSE) |
| constraints | JSON | 否 | 权限约束条件 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/permissionconfig/list | 权限列表 |
| GET | /wms/permissionconfig/{id} | 权限详情 |
| POST | /wms/permissionconfig | 新增权限 |
| PUT | /wms/permissionconfig/{id} | 修改权限 |
| DELETE | /wms/permissionconfig/{id} | 删除权限 |

---

### 8.8 库存管理

#### 8.8.1 库存变化 (inventorychange)

**路由路径**: `/wms/inventorychange`

**功能说明**: 库存变化记录查询，追溯库存的每一次变动。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| change_id | BIGINT | 是 | 变化ID |
| balance_id | BIGINT | 是 | 库存台账ID |
| change_type | VARCHAR(20) | 是 | 变化类型(IN/OUT/TRANSFER/ADJUST) |
| before_qty | DECIMAL(18,6) | 是 | 变化前数量 |
| after_qty | DECIMAL(18,6) | 是 | 变化后数量 |
| change_qty | DECIMAL(18,6) | 是 | 变化数量 |
| change_time | DATETIME | 是 | 变化时间 |
| operator_id | BIGINT | 是 | 操作人ID |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/inventorychange/list | 变化记录列表 |
| GET | /wms/inventorychange/{id} | 变化记录详情 |
| GET | /wms/inventorychange/export-excel | 导出变化记录 |

---

#### 8.8.2 库存快照 (inventorysnapshot)

**路由路径**: `/wms/inventorysnapshot`

**功能说明**: 库存快照查询，查看特定时间点的库存状况。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| snapshot_id | BIGINT | 是 | 快照ID |
| snapshot_time | DATETIME | 是 | 快照时间 |
| warehouse_id | BIGINT | 是 | 仓库ID |
| material_id | BIGINT | 是 | 物料ID |
| location_id | BIGINT | 否 | 库位ID |
| quantity | DECIMAL(18,6) | 是 | 库存数量 |
| frozen_quantity | DECIMAL(18,6) | 是 | 冻结数量 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/inventorysnapshot/list | 快照列表 |
| GET | /wms/inventorysnapshot/{id} | 快照详情 |
| GET | /wms/inventorysnapshot/time | 按时间查询快照 |

---

#### 8.8.3 容器管理 (container) - 已存在于7.14节

#### 8.8.4 库存冻结 (inventoryfreeze)

**路由路径**: `/wms/inventoryfreeze`

**功能说明**: 库存冻结管理，对库存进行冻结和解冻操作。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| freeze_id | BIGINT | 是 | 冻结ID |
| balance_id | BIGINT | 是 | 库存台账ID |
| freeze_type | VARCHAR(20) | 是 | 冻结类型(QC/PENDING/ORDER) |
| freeze_qty | DECIMAL(18,6) | 是 | 冻结数量 |
| reason | VARCHAR(200) | 否 | 冻结原因 |
| freeze_time | DATETIME | 是 | 冻结时间 |
| unfreeze_time | DATETIME | 否 | 解冻时间 |
| status | VARCHAR(20) | 是 | 状态(FROZEN/UNFROZEN) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/inventoryfreeze/list | 冻结记录列表 |
| GET | /wms/inventoryfreeze/{id} | 冻结记录详情 |
| POST | /wms/inventoryfreeze | 新增冻结 |
| PUT | /wms/inventoryfreeze/{id}/unfreeze | 解冻 |

---

#### 8.8.5 库存锁定 (inventorylock)

**路由路径**: `/wms/inventorylock`

**功能说明**: 库存锁定管理，对库存进行锁定操作。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| lock_id | BIGINT | 是 | 锁定ID |
| balance_id | BIGINT | 是 | 库存台账ID |
| lock_type | VARCHAR(20) | 是 | 锁定类型(ORDER/SYSTEM) |
| lock_qty | DECIMAL(18,6) | 是 | 锁定数量 |
| reason | VARCHAR(200) | 否 | 锁定原因 |
| lock_time | DATETIME | 是 | 锁定时间 |
| unlock_time | DATETIME | 否 | 解锁时间 |
| status | VARCHAR(20) | 是 | 状态(LOCKED/UNLOCKED) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/inventorylock/list | 锁定记录列表 |
| GET | /wms/inventorylock/{id} | 锁定记录详情 |
| POST | /wms/inventorylock | 新增锁定 |
| PUT | /wms/inventorylock/{id}/unlock | 解锁 |

---

#### 8.8.6 库存分配 (inventoryallocation)

**路由路径**: `/wms/inventoryallocation`

**功能说明**: 库存分配记录，查看订单的库存分配情况。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| allocation_id | BIGINT | 是 | 分配ID |
| order_id | BIGINT | 是 | 订单ID |
| order_type | VARCHAR(20) | 是 | 订单类型 |
| balance_id | BIGINT | 是 | 库存台账ID |
| allocated_qty | DECIMAL(18,6) | 是 | 分配数量 |
| create_time | DATETIME | 是 | 创建时间 |
| status | VARCHAR(20) | 是 | 状态(ALLOCATED/RELEASED) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/inventoryallocation/list | 分配记录列表 |
| GET | /wms/inventoryallocation/{id} | 分配记录详情 |
| PUT | /wms/inventoryallocation/{id}/release | 释放分配 |

---

#### 8.8.7 库存预留 (inventoryreservation)

**路由路径**: `/wms/inventoryreservation`

**功能说明**: 库存预留管理，为未来订单预留库存。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| reservation_id | BIGINT | 是 | 预留ID |
| order_id | BIGINT | 否 | 订单ID |
| balance_id | BIGINT | 是 | 库存台账ID |
| reserved_qty | DECIMAL(18,6) | 是 | 预留数量 |
| expiry_time | DATETIME | 否 | 预留失效时间 |
| create_time | DATETIME | 是 | 创建时间 |
| status | VARCHAR(20) | 是 | 状态(RESERVED/RELEASED/EXPIRED) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/inventoryreservation/list | 预留记录列表 |
| GET | /wms/inventoryreservation/{id} | 预留记录详情 |
| POST | /wms/inventoryreservation | 新增预留 |
| PUT | /wms/inventoryreservation/{id}/release | 释放预留 |

---

#### 8.8.8 库存查询 (inventoryquery)

**路由路径**: `/wms/inventoryquery`

**功能说明**: 综合库存查询，支持多条件筛选和批量导出。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| warehouse_id | BIGINT | 否 | 仓库ID |
| material_id | BIGINT | 否 | 物料ID |
| location_id | BIGINT | 否 | 库位ID |
| batch_no | VARCHAR(50) | 否 | 批次号 |
| status | VARCHAR(20) | 否 | 库存状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/inventoryquery/list | 库存列表 |
| GET | /wms/inventoryquery/summary | 库存汇总 |
| GET | /wms/inventoryquery/export-excel | 导出库存 |
| POST | /wms/inventoryquery/senior | 高级搜索 |

---

### 8.9 入库管理

#### 8.9.1 采购订单 (purchaseorder)

**路由路径**: `/wms/purchaseorder`

**功能说明**: 采购订单查询，查看来自SCP的采购订单信息。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| po_no | VARCHAR(50) | 是 | 采购订单号 |
| supplier_id | BIGINT | 是 | 供应商ID |
| order_date | DATE | 是 | 订单日期 |
| expected_date | DATE | 是 | 预计到货日期 |
| total_qty | DECIMAL(18,6) | 是 | 总数量 |
| total_amount | DECIMAL(18,2) | 否 | 总金额 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/purchaseorder/list | 采购订单列表 |
| GET | /wms/purchaseorder/{id} | 采购订单详情 |
| GET | /wms/purchaseorder/export-excel | 导出采购订单 |

---

#### 8.9.2 采购到货 (purchasedelivery)

**路由路径**: `/wms/purchasedelivery`

**功能说明**: 采购到货记录，登记供应商的实际到货信息。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| delivery_id | BIGINT | 是 | 到货记录ID |
| po_no | VARCHAR(50) | 是 | 采购订单号 |
| supplier_id | BIGINT | 是 | 供应商ID |
| arrive_date | DATE | 是 | 到货日期 |
| arrive_qty | DECIMAL(18,6) | 是 | 到货数量 |
| quality_status | VARCHAR(20) | 是 | 质量状态(PENDING/PASS/FAIL) |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/purchasedelivery/list | 到货记录列表 |
| GET | /wms/purchasedelivery/{id} | 到货记录详情 |
| POST | /wms/purchasedelivery | 新增到货记录 |
| PUT | /wms/purchasedelivery/{id} | 修改到货记录 |

---

#### 8.9.3 IQC来料检验 (iqc) - 已存在于7.10节

#### 8.9.4 生产入库 (productreceipt) - 已存在于7.11节

#### 8.9.5 退货入库 (returnreceipt)

**路由路径**: `/wms/returnreceipt`

**功能说明**: 退货入库管理，处理客户退货或供应商退货的入库业务。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| return_no | VARCHAR(50) | 是 | 退货单号 |
| source_type | VARCHAR(20) | 是 | 来源类型(CUSTOMER/SUPPLIER) |
| source_no | VARCHAR(50) | 是 | 来源单号 |
| return_date | DATE | 是 | 退货日期 |
| total_qty | DECIMAL(18,6) | 是 | 总数量 |
| warehouse_id | BIGINT | 是 | 收货仓库ID |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/returnreceipt/list | 退货单列表 |
| GET | /wms/returnreceipt/{id} | 退货单详情 |
| POST | /wms/returnreceipt | 新增退货单 |
| PUT | /wms/returnreceipt/{id} | 修改退货单 |
| POST | /wms/returnreceipt/{id}/confirm | 确认入库 |

---

#### 8.9.6 调拨入库 (transferreceipt)

**路由路径**: `/wms/transferreceipt`

**功能说明**: 调拨入库管理，处理库间调拨的入库确认。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| transfer_no | VARCHAR(50) | 是 | 调拨单号 |
| source_warehouse_id | BIGINT | 是 | 源仓库ID |
| dest_warehouse_id | BIGINT | 是 | 目标仓库ID |
| transfer_date | DATE | 是 | 调拨日期 |
| total_qty | DECIMAL(18,6) | 是 | 总数量 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/transferreceipt/list | 调拨入库列表 |
| GET | /wms/transferreceipt/{id} | 调拨入库详情 |
| POST | /wms/transferreceipt/{id}/confirm | 确认入库 |

---

#### 8.9.7 其他入库 (otherreceipt)

**路由路径**: `/wms/otherreceipt`

**功能说明**: 其他入库单管理，处理盘点盈亏、调整等特殊入库业务。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| receipt_no | VARCHAR(50) | 是 | 入库单号 |
| receipt_type | VARCHAR(20) | 是 | 入库类型(ADJUST/PROFIT/LENDING) |
| receipt_date | DATE | 是 | 入库日期 |
| total_qty | DECIMAL(18,6) | 是 | 总数量 |
| warehouse_id | BIGINT | 是 | 仓库ID |
| remark | VARCHAR(500) | 否 | 备注 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/otherreceipt/list | 其他入库列表 |
| GET | /wms/otherreceipt/{id} | 其他入库详情 |
| POST | /wms/otherreceipt | 新增其他入库 |
| PUT | /wms/otherreceipt/{id} | 修改其他入库 |
| POST | /wms/otherreceipt/{id}/confirm | 确认入库 |

---

#### 8.9.8 入库确认 (receiptconfirm)

**路由路径**: `/wms/receiptconfirm`

**功能说明**: 入库确认执行，对入库单进行确认操作。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| confirm_id | BIGINT | 是 | 确认ID |
| receipt_id | BIGINT | 是 | 入库单ID |
| receipt_type | VARCHAR(20) | 是 | 入库类型 |
| confirm_qty | DECIMAL(18,6) | 是 | 确认数量 |
| confirm_time | DATETIME | 是 | 确认时间 |
| operator_id | BIGINT | 是 | 操作人ID |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/receiptconfirm/list | 确认记录列表 |
| POST | /wms/receiptconfirm | 执行确认 |

---

#### 8.9.9 入库上架 (receiptputaway)

**路由路径**: `/wms/receiptputaway`

**功能说明**: 入库上架执行，将到货物料上架到指定库位。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| putaway_id | BIGINT | 是 | 上架ID |
| receipt_id | BIGINT | 是 | 入库单ID |
| location_id | BIGINT | 是 | 目标库位ID |
| putaway_qty | DECIMAL(18,6) | 是 | 上架数量 |
| putaway_time | DATETIME | 是 | 上架时间 |
| operator_id | BIGINT | 是 | 操作人ID |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/receiptputaway/list | 上架记录列表 |
| GET | /wms/receiptputaway/{id} | 上架记录详情 |
| POST | /wms/receiptputaway | 执行上架 |
| POST | /wms/receiptputaway/batch | 批量上架 |

---

#### 8.9.10 ASN管理 (asn)

**路由路径**: `/wms/asn`

**功能说明**: ASN(Advanced Shipping Notification)管理，供应商发货预报。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| asn_no | VARCHAR(50) | 是 | ASN单号 |
| supplier_id | BIGINT | 是 | 供应商ID |
| expected_date | DATE | 是 | 预计到货日期 |
| total_qty | DECIMAL(18,6) | 是 | 总数量 |
| actual_date | DATE | 否 | 实际到货日期 |
| actual_qty | DECIMAL(18,6) | 否 | 实际数量 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/asn/list | ASN列表 |
| GET | /wms/asn/{id} | ASN详情 |
| POST | /wms/asn | 新增ASN |
| PUT | /wms/asn/{id} | 修改ASN |
| POST | /wms/asn/{id}/receive | 到货接收 |

---

#### 8.9.11 到货登记 (deliveryregister)

**路由路径**: `/wms/deliveryregister`

**功能说明**: 到货登记管理，记录供应商送货的到达信息。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| register_id | BIGINT | 是 | 登记ID |
| asn_no | VARCHAR(50) | 是 | ASN单号 |
| arrive_date | DATE | 是 | 到达日期 |
| arrive_time | TIME | 是 | 到达时间 |
| vehicle_no | VARCHAR(50) | 否 | 车牌号 |
| driver_name | VARCHAR(50) | 否 | 司机姓名 |
| driver_phone | VARCHAR(50) | 否 | 司机电话 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/deliveryregister/list | 登记列表 |
| GET | /wms/deliveryregister/{id} | 登记详情 |
| POST | /wms/deliveryregister | 新增登记 |
| PUT | /wms/deliveryregister/{id} | 修改登记 |

---

### 8.10 出库管理

#### 8.10.1 销售订单 (saleorder)

**路由路径**: `/wms/saleorder`

**功能说明**: 销售订单查询，查看来自SCP的销售订单信息。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| so_no | VARCHAR(50) | 是 | 销售订单号 |
| customer_id | BIGINT | 是 | 客户ID |
| order_date | DATE | 是 | 订单日期 |
| expected_date | DATE | 是 | 预计发货日期 |
| total_qty | DECIMAL(18,6) | 是 | 总数量 |
| total_amount | DECIMAL(18,2) | 否 | 总金额 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/saleorder/list | 销售订单列表 |
| GET | /wms/saleorder/{id} | 销售订单详情 |
| GET | /wms/saleorder/export-excel | 导出销售订单 |

---

#### 8.10.2 销售发货 (saledelivery)

**路由路径**: `/wms/saledelivery`

**功能说明**: 销售发货管理，执行销售订单的发货操作。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| delivery_id | BIGINT | 是 | 发货ID |
| so_no | VARCHAR(50) | 是 | 销售订单号 |
| customer_id | BIGINT | 是 | 客户ID |
| ship_date | DATE | 是 | 发货日期 |
| ship_qty | DECIMAL(18,6) | 是 | 发货数量 |
| vehicle_no | VARCHAR(50) | 否 | 车牌号 |
| driver_name | VARCHAR(50) | 否 | 司机姓名 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/saledelivery/list | 发货记录列表 |
| GET | /wms/saledelivery/{id} | 发货记录详情 |
| POST | /wms/saledelivery | 新增发货记录 |
| PUT | /wms/saledelivery/{id} | 修改发货记录 |
| POST | /wms/saledelivery/{id}/confirm | 确认发货 |

---

#### 8.10.3 领料管理 (issue) - 已存在于7.12节

#### 8.10.4 备料计划 (preparetoissueplan) - 已存在于7.13节

#### 8.10.5 生产领料 (productionissue)

**路由路径**: `/wms/productionissue`

**功能说明**: 生产领料执行，记录生产车间的实际领料情况。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| issue_id | BIGINT | 是 | 领料ID |
| work_order_id | BIGINT | 是 | 工单ID |
| material_id | BIGINT | 是 | 物料ID |
| issue_qty | DECIMAL(18,6) | 是 | 领料数量 |
| issue_time | DATETIME | 是 | 领料时间 |
| operator_id | BIGINT | 是 | 操作人ID |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/productionissue/list | 领料记录列表 |
| GET | /wms/productionissue/{id} | 领料记录详情 |
| POST | /wms/productionissue | 新增领料记录 |
| POST | /wms/productionissue/batch | 批量领料 |

---

#### 8.10.6 拣货作业 (pickingjob)

**路由路径**: `/wms/pickingjob`

**功能说明**: 拣货作业管理，执行销售/领料单的拣货下架操作。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| picking_id | BIGINT | 是 | 拣货ID |
| order_id | BIGINT | 是 | 订单ID |
| order_type | VARCHAR(20) | 是 | 订单类型 |
| location_id | BIGINT | 是 | 拣货库位ID |
| material_id | BIGINT | 是 | 物料ID |
| picked_qty | DECIMAL(18,6) | 是 | 拣货数量 |
| picker_id | BIGINT | 是 | 拣货人ID |
| picking_time | DATETIME | 是 | 拣货时间 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/pickingjob/list | 拣货记录列表 |
| GET | /wms/pickingjob/{id} | 拣货记录详情 |
| POST | /wms/pickingjob | 执行拣货 |
| POST | /wms/pickingjob/batch | 批量拣货 |

---

#### 8.10.7 补料作业 (replenishjob)

**路由路径**: `/wms/replenishjob`

**功能说明**: 补料作业管理，执行线边库的补料操作。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| replenish_id | BIGINT | 是 | 补料ID |
| location_id | BIGINT | 是 | 补料库位ID |
| material_id | BIGINT | 是 | 物料ID |
| replenish_qty | DECIMAL(18,6) | 是 | 补料数量 |
| replenish_time | DATETIME | 是 | 补料时间 |
| operator_id | BIGINT | 是 | 操作人ID |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/replenishjob/list | 补料记录列表 |
| GET | /wms/replenishjob/{id} | 补料记录详情 |
| POST | /wms/replenishjob | 执行补料 |
| POST | /wms/replenishjob/batch | 批量补料 |

---

#### 8.10.8 出库确认 (deliveryconfirm)

**路由路径**: `/wms/deliveryconfirm`

**功能说明**: 出库确认执行，对出库单进行确认操作。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| confirm_id | BIGINT | 是 | 确认ID |
| delivery_id | BIGINT | 是 | 出库单ID |
| confirm_qty | DECIMAL(18,6) | 是 | 确认数量 |
| confirm_time | DATETIME | 是 | 确认时间 |
| operator_id | BIGINT | 是 | 操作人ID |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/deliveryconfirm/list | 确认记录列表 |
| POST | /wms/deliveryconfirm | 执行确认 |

---

#### 8.10.9 出库复核 (deliveryreview)

**路由路径**: `/wms/deliveryreview`

**功能说明**: 出库复核管理，对出库物料进行复核确认。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| review_id | BIGINT | 是 | 复核ID |
| delivery_id | BIGINT | 是 | 出库单ID |
| review_qty | DECIMAL(18,6) | 是 | 复核数量 |
| reviewer_id | BIGINT | 是 | 复核人ID |
| review_time | DATETIME | 是 | 复核时间 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/deliveryreview/list | 复核记录列表 |
| GET | /wms/deliveryreview/{id} | 复核记录详情 |
| POST | /wms/deliveryreview | 执行复核 |

---

#### 8.10.10 装车管理 (loadingmanagement)

**路由路径**: `/wms/loadingmanagement`

**功能说明**: 装车管理，管理发货商品的装车信息。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| loading_id | BIGINT | 是 | 装车ID |
| vehicle_no | VARCHAR(50) | 是 | 车牌号 |
| delivery_ids | VARCHAR(500) | 是 | 发货单ID列表 |
| loading_time | DATETIME | 是 | 装车时间 |
| operator_id | BIGINT | 是 | 操作人ID |
| total_qty | DECIMAL(18,6) | 是 | 装车总数量 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/loadingmanagement/list | 装车记录列表 |
| GET | /wms/loadingmanagement/{id} | 装车记录详情 |
| POST | /wms/loadingmanagement | 新增装车 |
| PUT | /wms/loadingmanagement/{id} | 修改装车 |

---

#### 8.10.11 送货管理 (deliverymanagement)

**路由路径**: `/wms/deliverymanagement`

**功能说明**: 送货管理，管理货物的运输和送达信息。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| delivery_id | BIGINT | 是 | 送货ID |
| loading_id | BIGINT | 是 | 装车ID |
| vehicle_no | VARCHAR(50) | 是 | 车牌号 |
| driver_name | VARCHAR(50) | 是 | 司机姓名 |
| driver_phone | VARCHAR(50) | 是 | 司机电话 |
| delivery_address | VARCHAR(500) | 是 | 送货地址 |
| estimated_arrival | DATETIME | 否 | 预计到达时间 |
| actual_arrival | DATETIME | 否 | 实际到达时间 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/deliverymanagement/list | 送货记录列表 |
| GET | /wms/deliverymanagement/{id} | 送货记录详情 |
| POST | /wms/deliverymanagement | 新增送货 |
| PUT | /wms/deliverymanagement/{id} | 修改送货 |
| PUT | /wms/deliverymanagement/{id}/depart | 发车 |
| PUT | /wms/deliverymanagement/{id}/arrive | 到达 |

---

#### 8.10.12 签收管理 (signmanagement)

**路由路径**: `/wms/signmanagement`

**功能说明**: 签收管理，记录客户对货物的签收确认。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| sign_id | BIGINT | 是 | 签收ID |
| delivery_id | BIGINT | 是 | 送货ID |
| sign_time | DATETIME | 是 | 签收时间 |
| sign_by | VARCHAR(100) | 是 | 签收人 |
| sign_qty | DECIMAL(18,6) | 是 | 签收数量 |
| sign_type | VARCHAR(20) | 是 | 签收类型(SIGNATURE/PHOTO) |
| sign_image | VARCHAR(500) | 否 | 签收照片 |
| remark | VARCHAR(500) | 否 | 备注 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/signmanagement/list | 签收记录列表 |
| GET | /wms/signmanagement/{id} | 签收记录详情 |
| POST | /wms/signmanagement | 新增签收 |
| PUT | /wms/signmanagement/{id} | 修改签收 |

---

#### 8.10.13 运费结算 (freightsettlement)

**路由路径**: `/wms/freightsettlement`

**功能说明**: 运费结算管理，记录和结算运输费用。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| settlement_id | BIGINT | 是 | 结算ID |
| carrier_id | BIGINT | 是 | 承运商ID |
| settlement_date | DATE | 是 | 结算日期 |
| total_amount | DECIMAL(18,2) | 是 | 结算总金额 |
| delivery_count | INT | 是 | 发货票数 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/freightsettlement/list | 结算记录列表 |
| GET | /wms/freightsettlement/{id} | 结算记录详情 |
| POST | /wms/freightsettlement | 新增结算 |
| PUT | /wms/freightsettlement/{id}/confirm | 确认结算 |

---

### 8.11 移库管理

#### 8.11.1 移库申请 (inventorymoverequest)

**路由路径**: `/wms/inventorymoverequest`

**功能说明**: 移库申请管理，申请库存的库位移动。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| move_no | VARCHAR(50) | 是 | 移库单号 |
| source_location_id | BIGINT | 是 | 源库位ID |
| dest_location_id | BIGINT | 是 | 目标库位ID |
| material_id | BIGINT | 是 | 物料ID |
| move_qty | DECIMAL(18,6) | 是 | 移动数量 |
| reason | VARCHAR(200) | 否 | 移库原因 |
| request_time | DATETIME | 是 | 申请时间 |
| applicant_id | BIGINT | 是 | 申请人ID |
| status | VARCHAR(20) | 是 | 状态(PENDING/APPROVED/REJECTED/EXECUTED) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/inventorymoverequest/list | 移库申请列表 |
| GET | /wms/inventorymoverequest/{id} | 移库申请详情 |
| POST | /wms/inventorymoverequest | 新增移库申请 |
| PUT | /wms/inventorymoverequest/{id} | 修改移库申请 |
| PUT | /wms/inventorymoverequest/{id}/submit | 提交申请 |
| PUT | /wms/inventorymoverequest/{id}/approve | 审批通过 |
| PUT | /wms/inventorymoverequest/{id}/reject | 审批拒绝 |

---

#### 8.11.2 移库作业 (inventorymovejob)

**路由路径**: `/wms/inventorymovejob`

**功能说明**: 移库作业执行，执行库位间的库存移动任务。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| job_id | BIGINT | 是 | 作业ID |
| move_no | VARCHAR(50) | 是 | 移库单号 |
| operator_id | BIGINT | 是 | 操作人ID |
| execute_time | DATETIME | 是 | 执行时间 |
| status | VARCHAR(20) | 是 | 状态(PENDING/EXECUTING/COMPLETED) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/inventorymovejob/list | 移库作业列表 |
| GET | /wms/inventorymovejob/{id} | 移库作业详情 |
| POST | /wms/inventorymovejob | 执行移库作业 |
| POST | /wms/inventorymovejob/batch | 批量执行 |

---

#### 8.11.3 移库记录 (inventorymoverecord)

**路由路径**: `/wms/inventorymoverecord`

**功能说明**: 移库记录查询，追溯历史移库操作。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| record_id | BIGINT | 是 | 记录ID |
| move_no | VARCHAR(50) | 是 | 移库单号 |
| source_location_id | BIGINT | 是 | 源库位ID |
| dest_location_id | BIGINT | 是 | 目标库位ID |
| material_id | BIGINT | 是 | 物料ID |
| move_qty | DECIMAL(18,6) | 是 | 移动数量 |
| move_time | DATETIME | 是 | 移动时间 |
| operator_id | BIGINT | 是 | 操作人ID |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/inventorymoverecord/list | 移库记录列表 |
| GET | /wms/inventorymoverecord/{id} | 移库记录详情 |
| GET | /wms/inventorymoverecord/export-excel | 导出移库记录 |

---

#### 8.11.4 库间调拨 (warehousetransfer)

**路由路径**: `/wms/warehousetransfer`

**功能说明**: 库间调拨管理，处理不同仓库间的库存调拨业务。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| transfer_no | VARCHAR(50) | 是 | 调拨单号 |
| source_warehouse_id | BIGINT | 是 | 源仓库ID |
| dest_warehouse_id | BIGINT | 是 | 目标仓库ID |
| transfer_date | DATE | 是 | 调拨日期 |
| total_qty | DECIMAL(18,6) | 是 | 总数量 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/warehousetransfer/list | 调拨单列表 |
| GET | /wms/warehousetransfer/{id} | 调拨单详情 |
| POST | /wms/warehousetransfer | 新增调拨单 |
| PUT | /wms/warehousetransfer/{id} | 修改调拨单 |
| POST | /wms/warehousetransfer/{id}/confirm | 确认调拨 |

---

#### 8.11.5 库内移动 (locationmove)

**路由路径**: `/wms/locationmove`

**功能说明**: 库内移动管理，处理同一仓库内库位间的移动。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| move_no | VARCHAR(50) | 是 | 移动单号 |
| source_location_id | BIGINT | 是 | 源库位ID |
| dest_location_id | BIGINT | 是 | 目标库位ID |
| material_id | BIGINT | 是 | 物料ID |
| move_qty | DECIMAL(18,6) | 是 | 移动数量 |
| reason | VARCHAR(200) | 否 | 移动原因 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/locationmove/list | 库内移动列表 |
| GET | /wms/locationmove/{id} | 库内移动详情 |
| POST | /wms/locationmove | 执行库内移动 |
| POST | /wms/locationmove/batch | 批量移动 |

---

#### 8.11.6 盘点调整 (countadjust)

**路由路径**: `/wms/countadjust`

**功能说明**: 盘点调整申请，记录盘点差异并申请调整。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| adjust_id | BIGINT | 是 | 调整ID |
| count_id | BIGINT | 是 | 盘点任务ID |
| material_id | BIGINT | 是 | 物料ID |
| location_id | BIGINT | 是 | 库位ID |
| system_qty | DECIMAL(18,6) | 是 | 系统数量 |
| actual_qty | DECIMAL(18,6) | 是 | 实际数量 |
| adjust_qty | DECIMAL(18,6) | 是 | 调整数量 |
| reason | VARCHAR(200) | 否 | 调整原因 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/countadjust/list | 调整记录列表 |
| GET | /wms/countadjust/{id} | 调整记录详情 |
| POST | /wms/countadjust | 新增调整 |
| PUT | /wms/countadjust/{id}/approve | 审批调整 |
| PUT | /wms/countadjust/{id}/reject | 拒绝调整 |

---

#### 8.11.7 库存损益 (inventoryprofitloss)

**路由路径**: `/wms/inventoryprofitloss`

**功能说明**: 库存损益记录，管理盘点发现的盈亏情况。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| pl_id | BIGINT | 是 | 损益ID |
| material_id | BIGINT | 是 | 物料ID |
| location_id | BIGINT | 是 | 库位ID |
| profit_loss_type | VARCHAR(20) | 是 | 损益类型(PROFIT/LOSS) |
| qty | DECIMAL(18,6) | 是 | 损益数量 |
| amount | DECIMAL(18,2) | 否 | 损益金额 |
| reason | VARCHAR(200) | 否 | 原因 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/inventoryprofitloss/list | 损益记录列表 |
| GET | /wms/inventoryprofitloss/{id} | 损益记录详情 |
| POST | /wms/inventoryprofitloss | 新增损益记录 |
| PUT | /wms/inventoryprofitloss/{id}/approve | 审批损益 |

---

#### 8.11.8 库存盘点 (stockcheck) - 盘点任务执行

**路由路径**: `/wms/stockcheck`

**功能说明**: 库存盘点执行，执行盘点的具体任务。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| check_id | BIGINT | 是 | 盘点ID |
| check_plan_id | BIGINT | 是 | 盘点计划ID |
| material_id | BIGINT | 是 | 物料ID |
| location_id | BIGINT | 是 | 库位ID |
| system_qty | DECIMAL(18,6) | 是 | 系统数量 |
| count_qty | DECIMAL(18,6) | 是 | 盘点数量 |
| diff_qty | DECIMAL(18,6) | 是 | 差异数量 |
| count_time | DATETIME | 是 | 盘点时间 |
| counter_id | BIGINT | 是 | 盘点人ID |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/stockcheck/list | 盘点记录列表 |
| GET | /wms/stockcheck/{id} | 盘点记录详情 |
| POST | /wms/stockcheck | 执行盘点 |
| POST | /wms/stockcheck/batch | 批量盘点 |

---

#### 8.11.9 库存复核 (countreview)

**路由路径**: `/wms/countreview`

**功能说明**: 盘点复核管理，对盘点结果进行复核确认。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| review_id | BIGINT | 是 | 复核ID |
| check_id | BIGINT | 是 | 盘点ID |
| review_qty | DECIMAL(18,6) | 是 | 复核数量 |
| reviewer_id | BIGINT | 是 | 复核人ID |
| review_time | DATETIME | 是 | 复核时间 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/countreview/list | 复核记录列表 |
| GET | /wms/countreview/{id} | 复核记录详情 |
| POST | /wms/countreview | 执行复核 |

---

#### 8.11.10 差异处理 (differencehandling)

**路由路径**: `/wms/differencehandling`

**功能说明**: 盘点差异处理，记录和审批盘点差异。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| diff_id | BIGINT | 是 | 差异ID |
| check_id | BIGINT | 是 | 盘点ID |
| material_id | BIGINT | 是 | 物料ID |
| location_id | BIGINT | 是 | 库位ID |
| diff_qty | DECIMAL(18,6) | 是 | 差异数量 |
| handle_type | VARCHAR(20) | 是 | 处理方式(ADJUST/SCRAP/IGNORE) |
| handle_remark | VARCHAR(200) | 否 | 处理备注 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/differencehandling/list | 差异列表 |
| GET | /wms/differencehandling/{id} | 差异详情 |
| POST | /wms/differencehandling | 新增差异处理 |
| PUT | /wms/differencehandling/{id}/approve | 审批处理 |

---

### 8.12 库存作业管理

#### 8.12.1 容器绑定 (containerbind)

**路由路径**: `/wms/containerbind`

**功能说明**: 容器绑定管理，将物料与容器进行绑定。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| bind_id | BIGINT | 是 | 绑定ID |
| container_id | BIGINT | 是 | 容器ID |
| material_id | BIGINT | 是 | 物料ID |
| location_id | BIGINT | 是 | 库位ID |
| bind_qty | DECIMAL(18,6) | 是 | 绑定数量 |
| bind_time | DATETIME | 是 | 绑定时间 |
| operator_id | BIGINT | 是 | 操作人ID |
| status | VARCHAR(20) | 是 | 状态(BOUND/UNBOUND) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/containerbind/list | 绑定记录列表 |
| GET | /wms/containerbind/{id} | 绑定记录详情 |
| POST | /wms/containerbind | 执行绑定 |
| DELETE | /wms/containerbind/{id} | 解除绑定 |

---

#### 8.12.2 容器解绑 (containerunbind)

**路由路径**: `/wms/containerunbind`

**功能说明**: 容器解绑管理，解除物料与容器的绑定关系。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| unbind_id | BIGINT | 是 | 解绑ID |
| container_id | BIGINT | 是 | 容器ID |
| reason | VARCHAR(200) | 是 | 解绑原因 |
| unbind_time | DATETIME | 是 | 解绑时间 |
| operator_id | BIGINT | 是 | 操作人ID |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/containerunbind/list | 解绑记录列表 |
| GET | /wms/containerunbind/{id} | 解绑记录详情 |
| POST | /wms/containerunbind | 执行解绑 |

---

#### 8.12.3 包装管理 (packagingmanagement)

**路由路径**: `/wms/packagingmanagement`

**功能说明**: 包装管理，定义包装规格和包装层级。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| pack_id | BIGINT | 是 | 包装ID |
| material_id | BIGINT | 是 | 物料ID |
| pack_type | VARCHAR(20) | 是 | 包装类型(BOX/TRAY/PALLET) |
| pack_qty | DECIMAL(18,6) | 是 | 包装数量 |
| pack_unit | VARCHAR(20) | 是 | 包装单位 |
| length | DECIMAL(10,2) | 否 | 长度(cm) |
| width | DECIMAL(10,2) | 否 | 宽度(cm) |
| height | DECIMAL(10,2) | 否 | 高度(cm) |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/packagingmanagement/list | 包装列表 |
| GET | /wms/packagingmanagement/{id} | 包装详情 |
| POST | /wms/packagingmanagement | 新增包装 |
| PUT | /wms/packagingmanagement/{id} | 修改包装 |
| DELETE | /wms/packagingmanagement/{id} | 删除包装 |

---

#### 8.12.4 打包作业 (packjob)

**路由路径**: `/wms/packjob`

**功能说明**: 打包作业执行，对物料进行包装作业。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| job_id | BIGINT | 是 | 作业ID |
| order_id | BIGINT | 是 | 订单ID |
| pack_type | VARCHAR(20) | 是 | 包装类型 |
| material_id | BIGINT | 是 | 物料ID |
| pack_qty | DECIMAL(18,6) | 是 | 包装数量 |
| operator_id | BIGINT | 是 | 操作人ID |
| pack_time | DATETIME | 是 | 打包时间 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/packjob/list | 打包记录列表 |
| GET | /wms/packjob/{id} | 打包记录详情 |
| POST | /wms/packjob | 执行打包 |
| POST | /wms/packjob/batch | 批量打包 |

---

#### 8.12.5 拆包作业 (unpackjob)

**路由路径**: `/wms/unpackjob`

**功能说明**: 拆包作业执行，对包装物料进行拆包。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| job_id | BIGINT | 是 | 作业ID |
| container_id | BIGINT | 是 | 容器ID |
| material_id | BIGINT | 是 | 物料ID |
| unpack_qty | DECIMAL(18,6) | 是 | 拆包数量 |
| operator_id | BIGINT | 是 | 操作人ID |
| unpack_time | DATETIME | 是 | 拆包时间 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/unpackjob/list | 拆包记录列表 |
| GET | /wms/unpackjob/{id} | 拆包记录详情 |
| POST | /wms/unpackjob | 执行拆包 |
| POST | /wms/unpackjob/batch | 批量拆包 |

---

#### 8.12.6 换箱作业 (switchboxjob)

**路由路径**: `/wms/switchboxjob`

**功能说明**: 换箱作业执行，更换物料的包装容器。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| job_id | BIGINT | 是 | 作业ID |
| old_container_id | BIGINT | 是 | 原容器ID |
| new_container_id | BIGINT | 是 | 新容器ID |
| material_id | BIGINT | 是 | 物料ID |
| switch_qty | DECIMAL(18,6) | 是 | 换箱数量 |
| operator_id | BIGINT | 是 | 操作人ID |
| switch_time | DATETIME | 是 | 换箱时间 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/switchboxjob/list | 换箱记录列表 |
| GET | /wms/switchboxjob/{id} | 换箱记录详情 |
| POST | /wms/switchboxjob | 执行换箱 |
| POST | /wms/switchboxjob/batch | 批量换箱 |

---

#### 8.12.7 合并作业 (mergejob)

**路由路径**: `/wms/mergejob`

**功能说明**: 合并作业执行，将多个容器的物料合并。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| job_id | BIGINT | 是 | 作业ID |
| source_container_ids | VARCHAR(500) | 是 | 源容器ID列表 |
| target_container_id | BIGINT | 是 | 目标容器ID |
| material_id | BIGINT | 是 | 物料ID |
| merge_qty | DECIMAL(18,6) | 是 | 合并数量 |
| operator_id | BIGINT | 是 | 操作人ID |
| merge_time | DATETIME | 是 | 合并时间 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/mergejob/list | 合并记录列表 |
| GET | /wms/mergejob/{id} | 合并记录详情 |
| POST | /wms/mergejob | 执行合并 |

---

#### 8.12.8 分割作业 (splitjob)

**路由路径**: `/wms/splitjob`

**功能说明**: 分割作业执行，将一个容器的物料分割到多个容器。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| job_id | BIGINT | 是 | 作业ID |
| source_container_id | BIGINT | 是 | 源容器ID |
| target_container_ids | VARCHAR(500) | 是 | 目标容器ID列表 |
| material_id | BIGINT | 是 | 物料ID |
| split_qty | DECIMAL(18,6) | 是 | 分割数量 |
| operator_id | BIGINT | 是 | 操作人ID |
| split_time | DATETIME | 是 | 分割时间 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/splitjob/list | 分割记录列表 |
| GET | /wms/splitjob/{id} | 分割记录详情 |
| POST | /wms/splitjob | 执行分割 |

---

#### 8.12.9 报废管理 (scrapmanagement)

**路由路径**: `/wms/scrapmanagement`

**功能说明**: 报废综合管理，记录和审批库存报废。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| scrap_id | BIGINT | 是 | 报废ID |
| material_id | BIGINT | 是 | 物料ID |
| location_id | BIGINT | 是 | 库位ID |
| scrap_qty | DECIMAL(18,6) | 是 | 报废数量 |
| scrap_reason | VARCHAR(200) | 是 | 报废原因 |
| scrap_time | DATETIME | 是 | 报废时间 |
| applicant_id | BIGINT | 是 | 申请人ID |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/scrapmanagement/list | 报废记录列表 |
| GET | /wms/scrapmanagement/{id} | 报废记录详情 |
| POST | /wms/scrapmanagement | 新增报废 |
| PUT | /wms/scrapmanagement/{id}/approve | 审批报废 |
| PUT | /wms/scrapmanagement/{id}/reject | 拒绝报废 |

---

#### 8.12.10 报废申请 (scrapa pplication)

**路由路径**: `/wms/scrapa pplication`

**功能说明**: 报废申请管理，提交库存报废申请。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| apply_no | VARCHAR(50) | 是 | 申请单号 |
| material_id | BIGINT | 是 | 物料ID |
| scrap_qty | DECIMAL(18,6) | 是 | 报废数量 |
| reason | VARCHAR(200) | 是 | 报废原因 |
| apply_time | DATETIME | 是 | 申请时间 |
| applicant_id | BIGINT | 是 | 申请人ID |
| status | VARCHAR(20) | 是 | 状态(PENDING/APPROVED/REJECTED) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/scrapa pplication/list | 报废申请列表 |
| GET | /wms/scrapa pplication/{id} | 报废申请详情 |
| POST | /wms/scrapa pplication | 新增报废申请 |
| PUT | /wms/scrapa pplication/{id}/submit | 提交申请 |
| PUT | /wms/scrapa pplication/{id}/approve | 审批通过 |
| PUT | /wms/scrapa pplication/{id}/reject | 审批拒绝 |

---

#### 8.12.11 报废审批 (scrapapproval)

**路由路径**: `/wms/scrapapproval`

**功能说明**: 报废审批管理，对报废申请进行审批。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| approval_id | BIGINT | 是 | 审批ID |
| apply_no | VARCHAR(50) | 是 | 申请单号 |
| approver_id | BIGINT | 是 | 审批人ID |
| approval_result | VARCHAR(20) | 是 | 审批结果(APPROVE/REJECT) |
| approval_time | DATETIME | 是 | 审批时间 |
| approval_remark | VARCHAR(200) | 否 | 审批备注 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/scrapapproval/list | 审批记录列表 |
| GET | /wms/scrapapproval/{id} | 审批记录详情 |
| POST | /wms/scrapapproval | 执行审批 |

---

#### 8.12.12 备件管理 (sparemanagement)

**路由路径**: `/wms/sparemanagement`

**功能说明**: 备件档案管理，维护备件的基本信息。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| spare_id | BIGINT | 是 | 备件ID |
| spare_code | VARCHAR(50) | 是 | 备件编码 |
| spare_name | VARCHAR(200) | 是 | 备件名称 |
| spec | VARCHAR(100) | 否 | 规格型号 |
| unit | VARCHAR(20) | 是 | 单位 |
| stock_qty | DECIMAL(18,6) | 是 | 库存数量 |
| min_qty | DECIMAL(18,6) | 否 | 最小库存 |
| max_qty | DECIMAL(18,6) | 否 | 最大库存 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/sparemanagement/list | 备件列表 |
| GET | /wms/sparemanagement/{id} | 备件详情 |
| POST | /wms/sparemanagement | 新增备件 |
| PUT | /wms/sparemanagement/{id} | 修改备件 |
| DELETE | /wms/sparemanagement/{id} | 删除备件 |

---

#### 8.12.13 备件申请 (spareapplication)

**路由路径**: `/wms/spareapplication`

**功能说明**: 备件申请管理，申请领用备件。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| apply_no | VARCHAR(50) | 是 | 申请单号 |
| spare_id | BIGINT | 是 | 备件ID |
| apply_qty | DECIMAL(18,6) | 是 | 申请数量 |
| reason | VARCHAR(200) | 是 | 申请原因 |
| apply_time | DATETIME | 是 | 申请时间 |
| applicant_id | BIGINT | 是 | 申请人ID |
| status | VARCHAR(20) | 是 | 状态(PENDING/APPROVED/REJECTED) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/spareapplication/list | 备件申请列表 |
| GET | /wms/spareapplication/{id} | 备件申请详情 |
| POST | /wms/spareapplication | 新增备件申请 |
| PUT | /wms/spareapplication/{id}/submit | 提交申请 |
| PUT | /wms/spareapplication/{id}/approve | 审批通过 |
| PUT | /wms/spareapplication/{id}/reject | 审批拒绝 |

---

#### 8.12.14 备件入库 (sparereceipt)

**路由路径**: `/wms/sparereceipt`

**功能说明**: 备件入库管理，处理备件的入库业务。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| receipt_id | BIGINT | 是 | 入库ID |
| apply_no | VARCHAR(50) | 是 | 申请单号 |
| spare_id | BIGINT | 是 | 备件ID |
| receipt_qty | DECIMAL(18,6) | 是 | 入库数量 |
| receipt_time | DATETIME | 是 | 入库时间 |
| operator_id | BIGINT | 是 | 操作人ID |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/sparereceipt/list | 备件入库列表 |
| GET | /wms/sparereceipt/{id} | 备件入库详情 |
| POST | /wms/sparereceipt | 执行备件入库 |

---

#### 8.12.15 备件出库 (spareissue)

**路由路径**: `/wms/spareissue`

**功能说明**: 备件出库管理，处理备件的出库业务。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| issue_id | BIGINT | 是 | 出库ID |
| apply_no | VARCHAR(50) | 是 | 申请单号 |
| spare_id | BIGINT | 是 | 备件ID |
| issue_qty | DECIMAL(18,6) | 是 | 出库数量 |
| issue_time | DATETIME | 是 | 出库时间 |
| operator_id | BIGINT | 是 | 操作人ID |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/spareissue/list | 备件出库列表 |
| GET | /wms/spareissue/{id} | 备件出库详情 |
| POST | /wms/spareissue | 执行备件出库 |

---

#### 8.12.16 备件库存 (spareinventory)

**路由路径**: `/wms/spareinventory`

**功能说明**: 备件库存查询，查看备件的实时库存。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| spare_id | BIGINT | 是 | 备件ID |
| warehouse_id | BIGINT | 是 | 仓库ID |
| location_id | BIGINT | 否 | 库位ID |
| stock_qty | DECIMAL(18,6) | 是 | 库存数量 |
| frozen_qty | DECIMAL(18,6) | 是 | 冻结数量 |
| available_qty | DECIMAL(18,6) | 是 | 可用数量 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/spareinventory/list | 备件库存列表 |
| GET | /wms/spareinventory/{id} | 备件库存详情 |

---

#### 8.12.17 备件位置 (sparelocation)

**路由路径**: `/wms/sparelocation`

**功能说明**: 备件库位管理，定义备件的存放位置。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| location_id | BIGINT | 是 | 位置ID |
| spare_id | BIGINT | 是 | 备件ID |
| warehouse_id | BIGINT | 是 | 仓库ID |
| location_code | VARCHAR(50) | 是 | 库位编码 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/sparelocation/list | 备件位置列表 |
| GET | /wms/sparelocation/{id} | 备件位置详情 |
| POST | /wms/sparelocation | 新增备件位置 |
| PUT | /wms/sparelocation/{id} | 修改备件位置 |
| DELETE | /wms/sparelocation/{id} | 删除备件位置 |

---

#### 8.12.18 库存预警 (inventorywarning)

**路由路径**: `/wms/inventorywarning`

**功能说明**: 库存预警查询，显示库存不足或过剩的预警信息。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| warning_id | BIGINT | 是 | 预警ID |
| material_id | BIGINT | 是 | 物料ID |
| warehouse_id | BIGINT | 是 | 仓库ID |
| warning_type | VARCHAR(20) | 是 | 预警类型(LOW_STOCK/HIGH_STOCK) |
| threshold | DECIMAL(18,6) | 是 | 阈值 |
| current_qty | DECIMAL(18,6) | 是 | 当前库存 |
| warning_level | VARCHAR(20) | 是 | 预警级别(WARNING/CRITICAL) |
| create_time | DATETIME | 是 | 生成时间 |
| status | VARCHAR(20) | 是 | 状态(UNREAD/READ/HANDLED) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/inventorywarning/list | 预警列表 |
| GET | /wms/inventorywarning/{id} | 预警详情 |
| PUT | /wms/inventorywarning/{id}/handle | 处理预警 |
| PUT | /wms/inventorywarning/batch/handle | 批量处理预警 |

---

#### 8.12.19 效期预警 (expirywarning)

**路由路径**: `/wms/expirywarning`

**功能说明**: 效期预警查询，显示临近保质期到期的物料预警。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| warning_id | BIGINT | 是 | 预警ID |
| material_id | BIGINT | 是 | 物料ID |
| batch_no | VARCHAR(50) | 是 | 批次号 |
| production_date | DATE | 是 | 生产日期 |
| expiry_date | DATE | 是 | 过期日期 |
| remaining_days | INT | 是 | 剩余天数 |
| stock_qty | DECIMAL(18,6) | 是 | 库存数量 |
| warning_level | VARCHAR(20) | 是 | 预警级别(INFO/WARNING/CRITICAL) |
| status | VARCHAR(20) | 是 | 状态(UNREAD/READ/HANDLED) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/expirywarning/list | 效期预警列表 |
| GET | /wms/expirywarning/{id} | 效期预警详情 |
| PUT | /wms/expirywarning/{id}/handle | 处理预警 |
| PUT | /wms/expirywarning/batch/handle | 批量处理预警 |

---

### 8.13 结算管理

#### 8.13.1 客户收款 (customerreceipt)

**路由路径**: `/wms/customerreceipt`

**功能说明**: 客户收款记录，管理客户的收款信息。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| receipt_id | BIGINT | 是 | 收款ID |
| customer_id | BIGINT | 是 | 客户ID |
| receipt_amount | DECIMAL(18,2) | 是 | 收款金额 |
| receipt_date | DATE | 是 | 收款日期 |
| payment_method | VARCHAR(20) | 是 | 付款方式(CASH/TRANSFER/NOTE) |
| receipt_no | VARCHAR(50) | 是 | 收款单号 |
| remark | VARCHAR(500) | 否 | 备注 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/customerreceipt/list | 收款记录列表 |
| GET | /wms/customerreceipt/{id} | 收款记录详情 |
| POST | /wms/customerreceipt | 新增收款 |
| PUT | /wms/customerreceipt/{id} | 修改收款 |

---

#### 8.13.2 客户结算 (customersettlement)

**路由路径**: `/wms/customersettlement`

**功能说明**: 客户结算管理，处理与客户的货款结算。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| settlement_id | BIGINT | 是 | 结算ID |
| customer_id | BIGINT | 是 | 客户ID |
| settlement_date | DATE | 是 | 结算日期 |
| total_amount | DECIMAL(18,2) | 是 | 结算总金额 |
| paid_amount | DECIMAL(18,2) | 是 | 已付金额 |
| pending_amount | DECIMAL(18,2) | 是 | 待付金额 |
| status | VARCHAR(20) | 是 | 状态(PENDING/COMPLETED) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/customersettlement/list | 结算记录列表 |
| GET | /wms/customersettlement/{id} | 结算记录详情 |
| POST | /wms/customersettlement | 新增结算 |
| PUT | /wms/customersettlement/{id}/confirm | 确认结算 |

---

#### 8.13.3 备货结算 (stockingsettlement)

**路由路径**: `/wms/stockingsettlement`

**功能说明**: 备货服务结算，管理仓储备货服务的费用结算。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| settlement_id | BIGINT | 是 | 结算ID |
| customer_id | BIGINT | 是 | 客户ID |
| service_date | DATE | 是 | 服务日期 |
| service_type | VARCHAR(20) | 是 | 服务类型(STOCKING/PICKING/PACKING) |
| service_fee | DECIMAL(18,2) | 是 | 服务费用 |
| quantity | DECIMAL(18,6) | 否 | 操作数量 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/stockingsettlement/list | 结算记录列表 |
| GET | /wms/stockingsettlement/{id} | 结算记录详情 |
| POST | /wms/stockingsettlement | 新增结算 |
| PUT | /wms/stockingsettlement/{id}/confirm | 确认结算 |

---

#### 8.13.4 应付管理 (accountspayable)

**路由路径**: `/wms/accountspayable`

**功能说明**: 应付账款管理，处理应支付给供应商的款项。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| ap_id | BIGINT | 是 | 应付ID |
| supplier_id | BIGINT | 是 | 供应商ID |
| invoice_no | VARCHAR(50) | 是 | 发票号 |
| invoice_amount | DECIMAL(18,2) | 是 | 发票金额 |
| paid_amount | DECIMAL(18,2) | 是 | 已付金额 |
| due_date | DATE | 是 | 到期日期 |
| status | VARCHAR(20) | 是 | 状态(PENDING/PARTIAL/COMPLETED) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/accountspayable/list | 应付记录列表 |
| GET | /wms/accountspayable/{id} | 应付记录详情 |
| POST | /wms/accountspayable | 新增应付记录 |
| PUT | /wms/accountspayable/{id}/pay | 支付 |

---

#### 8.13.5 收款核销 (receiptverification)

**路由路径**: `/wms/receiptverification`

**功能说明**: 收款核销管理，将收款与应收款进行核销。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| verify_id | BIGINT | 是 | 核销ID |
| receipt_id | BIGINT | 是 | 收款ID |
| invoice_id | BIGINT | 是 | 发票ID |
| verify_amount | DECIMAL(18,2) | 是 | 核销金额 |
| verify_time | DATETIME | 是 | 核销时间 |
| operator_id | BIGINT | 是 | 操作人ID |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/receiptverification/list | 核销记录列表 |
| GET | /wms/receiptverification/{id} | 核销记录详情 |
| POST | /wms/receiptverification | 执行核销 |

---

#### 8.13.6 结算单据 (settlementdocument)

**路由路径**: `/wms/settlementdocument`

**功能说明**: 结算单据管理，管理结算相关的单据凭证。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| doc_id | BIGINT | 是 | 单据ID |
| settlement_id | BIGINT | 是 | 结算ID |
| doc_type | VARCHAR(20) | 是 | 单据类型(INVOICE/RECEIPT/VOUCHER) |
| doc_no | VARCHAR(50) | 是 | 单据号 |
| amount | DECIMAL(18,2) | 是 | 金额 |
| doc_date | DATE | 是 | 单据日期 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/settlementdocument/list | 单据列表 |
| GET | /wms/settlementdocument/{id} | 单据详情 |
| POST | /wms/settlementdocument | 新增单据 |
| PUT | /wms/settlementdocument/{id} | 修改单据 |
| DELETE | /wms/settlementdocument/{id} | 删除单据 |

---

### 8.14 供应商管理

#### 8.14.1 供应商账期 (supplierperiod) - 已在8.4.2节部分覆盖

#### 8.14.2 供应商发票 (supplierinvoice)

**路由路径**: `/wms/supplierinvoice`

**功能说明**: 供应商发票管理，管理供应商提供的发票。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| invoice_id | BIGINT | 是 | 发票ID |
| supplier_id | BIGINT | 是 | 供应商ID |
| invoice_no | VARCHAR(50) | 是 | 发票号 |
| invoice_amount | DECIMAL(18,2) | 是 | 发票金额 |
| invoice_date | DATE | 是 | 发票日期 |
| tax_amount | DECIMAL(18,2) | 否 | 税额 |
| status | VARCHAR(20) | 是 | 状态(PENDING/CONFIRMED/CANCELLED) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/supplierinvoice/list | 发票列表 |
| GET | /wms/supplierinvoice/{id} | 发票详情 |
| POST | /wms/supplierinvoice | 新增发票 |
| PUT | /wms/supplierinvoice/{id} | 修改发票 |
| DELETE | /wms/supplierinvoice/{id} | 删除发票 |

---

#### 8.14.3 供应商索赔 (supplierclaim)

**路由路径**: `/wms/supplierclaim`

**功能说明**: 供应商索赔管理，处理供应商提出的索赔请求。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| claim_id | BIGINT | 是 | 索赔ID |
| supplier_id | BIGINT | 是 | 供应商ID |
| claim_type | VARCHAR(20) | 是 | 索赔类型(QUALITY/DELIVERY/PRICE) |
| claim_amount | DECIMAL(18,2) | 是 | 索赔金额 |
| reason | VARCHAR(500) | 是 | 索赔原因 |
| claim_date | DATE | 是 | 索赔日期 |
| status | VARCHAR(20) | 是 | 状态(PENDING/HANDLED/REJECTED) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/supplierclaim/list | 索赔列表 |
| GET | /wms/supplierclaim/{id} | 索赔详情 |
| POST | /wms/supplierclaim | 新增索赔 |
| PUT | /wms/supplierclaim/{id} | 修改索赔 |
| PUT | /wms/supplierclaim/{id}/handle | 处理索赔 |
| PUT | /wms/supplierclaim/{id}/reject | 拒绝索赔 |

---

#### 8.14.4 供应商对账 (supplierreconciliation)

**路由路径**: `/wms/supplierreconciliation`

**功能说明**: 供应商对账管理，进行供应商往来账核对。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| recon_id | BIGINT | 是 | 对账ID |
| supplier_id | BIGINT | 是 | 供应商ID |
| recon_date | DATE | 是 | 对账日期 |
| statement_amount | DECIMAL(18,2) | 是 | 账单金额 |
| actual_amount | DECIMAL(18,2) | 是 | 实际金额 |
| diff_amount | DECIMAL(18,2) | 是 | 差异金额 |
| remark | VARCHAR(500) | 否 | 备注 |
| status | VARCHAR(20) | 是 | 状态(PENDING/CONFIRMED) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/supplierreconciliation/list | 对账列表 |
| GET | /wms/supplierreconciliation/{id} | 对账详情 |
| POST | /wms/supplierreconciliation | 新增对账 |
| PUT | /wms/supplierreconciliation/{id}/confirm | 确认对账 |

---

#### 8.14.5 供应商评级 (supplierrating)

**路由路径**: `/wms/supplierrating`

**功能说明**: 供应商评级管理，对供应商进行等级评定。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| rating_id | BIGINT | 是 | 评级ID |
| supplier_id | BIGINT | 是 | 供应商ID |
| rating_level | VARCHAR(10) | 是 | 评级等级(A/B/C/D) |
| rating_date | DATE | 是 | 评级日期 |
| evaluator_id | BIGINT | 是 | 评价人ID |
| quality_score | DECIMAL(5,2) | 否 | 质量评分 |
| delivery_score | DECIMAL(5,2) | 否 | 交期评分 |
| price_score | DECIMAL(5,2) | 否 | 价格评分 |
| overall_score | DECIMAL(5,2) | 是 | 综合评分 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/supplierrating/list | 评级列表 |
| GET | /wms/supplierrating/{id} | 评级详情 |
| POST | /wms/supplierrating | 新增评级 |
| PUT | /wms/supplierrating/{id} | 修改评级 |

---

#### 8.14.6 供应商资质 (supplierqualification)

**路由路径**: `/wms/supplierqualification`

**功能说明**: 供应商资质管理，管理供应商的资质证书。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| qual_id | BIGINT | 是 | 资质ID |
| supplier_id | BIGINT | 是 | 供应商ID |
| qual_type | VARCHAR(20) | 是 | 资质类型(CERT/LICENSE/PERMISSION) |
| qual_name | VARCHAR(200) | 是 | 资质名称 |
| cert_no | VARCHAR(50) | 是 | 证书编号 |
| valid_from | DATE | 是 | 生效日期 |
| valid_to | DATE | 是 | 失效日期 |
| attach_path | VARCHAR(500) | 否 | 附件路径 |
| status | VARCHAR(20) | 是 | 状态(VALID/EXPIRED/INVALID) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/supplierqualification/list | 资质列表 |
| GET | /wms/supplierqualification/{id} | 资质详情 |
| POST | /wms/supplierqualification | 新增资质 |
| PUT | /wms/supplierqualification/{id} | 修改资质 |
| DELETE | /wms/supplierqualification/{id} | 删除资质 |

---

#### 8.14.7 供应商合同 (suppliercontract)

**路由路径**: `/wms/suppliercontract`

**功能说明**: 供应商合同管理，管理与供应商签订的合同。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| contract_id | BIGINT | 是 | 合同ID |
| supplier_id | BIGINT | 是 | 供应商ID |
| contract_no | VARCHAR(50) | 是 | 合同编号 |
| contract_name | VARCHAR(200) | 是 | 合同名称 |
| contract_amount | DECIMAL(18,2) | 是 | 合同金额 |
| sign_date | DATE | 是 | 签订日期 |
| start_date | DATE | 是 | 开始日期 |
| end_date | DATE | 是 | 结束日期 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/EXPIRED/TERMINATED) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/suppliercontract/list | 合同列表 |
| GET | /wms/suppliercontract/{id} | 合同详情 |
| POST | /wms/suppliercontract | 新增合同 |
| PUT | /wms/suppliercontract/{id} | 修改合同 |
| DELETE | /wms/suppliercontract/{id} | 删除合同 |

---

#### 8.14.8 供应商考核 (supplierevaluation)

**路由路径**: `/wms/supplierevaluation`

**功能说明**: 供应商考核管理，对供应商进行定期考核评价。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| eval_id | BIGINT | 是 | 考核ID |
| supplier_id | BIGINT | 是 | 供应商ID |
| eval_date | DATE | 是 | 考核日期 |
| quality_score | DECIMAL(5,2) | 是 | 质量得分 |
| delivery_score | DECIMAL(5,2) | 是 | 交期得分 |
| service_score | DECIMAL(5,2) | 是 | 服务得分 |
| price_score | DECIMAL(5,2) | 是 | 价格得分 |
| overall_score | DECIMAL(5,2) | 是 | 综合得分 |
| eval_result | VARCHAR(20) | 是 | 考核结果(PASS/FAIL) |
| remark | VARCHAR(500) | 否 | 备注 |
| status | VARCHAR(20) | 是 | 状态 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/supplierevaluation/list | 考核列表 |
| GET | /wms/supplierevaluation/{id} | 考核详情 |
| POST | /wms/supplierevaluation | 新增考核 |
| PUT | /wms/supplierevaluation/{id} | 修改考核 |

---

### 8.15 AGV管理

#### 8.15.1 AGV设备 (agvequipment)

**路由路径**: `/wms/agvequipment`

**功能说明**: AGV设备管理，定义AGV的基本信息和工作状态。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| agv_id | BIGINT | 是 | AGV ID |
| agv_code | VARCHAR(50) | 是 | AGV编码 |
| agv_name | VARCHAR(100) | 是 | AGV名称 |
| agv_type | VARCHAR(20) | 是 | AGV类型(FORKLIFT/CONVEYOR/DRONE) |
| warehouse_id | BIGINT | 是 | 所属仓库ID |
| current_location_id | BIGINT | 否 | 当前位置ID |
| status | VARCHAR(20) | 是 | 状态(IDLE/BUSY/ERROR/OFFLINE) |
| battery_level | DECIMAL(5,2) | 否 | 电池电量(%) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/agvequipment/list | AGV列表 |
| GET | /wms/agvequipment/{id} | AGV详情 |
| POST | /wms/agvequipment | 新增AGV |
| PUT | /wms/agvequipment/{id} | 修改AGV |
| DELETE | /wms/agvequipment/{id} | 删除AGV |

---

#### 8.15.2 AGV库位 (agvlocation)

**路由路径**: `/wms/agvlocation`

**功能说明**: AGV库位管理，定义AGV可识别的库位信息。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| location_id | BIGINT | 是 | 库位ID |
| agv_station_code | VARCHAR(50) | 是 | AGV站点编码 |
| location_type | VARCHAR(20) | 是 | 位置类型(STORAGE/CHARGE/LOADING) |
| warehouse_id | BIGINT | 是 | 所属仓库ID |
| coordinate_x | DECIMAL(10,2) | 否 | X坐标 |
| coordinate_y | DECIMAL(10,2) | 否 | Y坐标 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE/OCCUPIED) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/agvlocation/list | AGV库位列表 |
| GET | /wms/agvlocation/{id} | AGV库位详情 |
| POST | /wms/agvlocation | 新增AGV库位 |
| PUT | /wms/agvlocation/{id} | 修改AGV库位 |
| DELETE | /wms/agvlocation/{id} | 删除AGV库位 |

---

#### 8.15.3 AGV任务 (agvtask)

**路由路径**: `/wms/agvtask`

**功能说明**: AGV任务管理，查看和管理AGV的运输任务。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| task_id | BIGINT | 是 | 任务ID |
| agv_id | BIGINT | 是 | AGV ID |
| task_type | VARCHAR(20) | 是 | 任务类型(PUTAWAY/PICK/TRANSFER) |
| source_location_id | BIGINT | 是 | 源位置ID |
| dest_location_id | BIGINT | 是 | 目标位置ID |
| material_id | BIGINT | 否 | 物料ID |
| priority | INT | 是 | 优先级(1-10) |
| status | VARCHAR(20) | 是 | 状态(PENDING/EXECUTING/COMPLETED/FAILED) |
| create_time | DATETIME | 是 | 创建时间 |
| execute_time | DATETIME | 否 | 执行时间 |
| complete_time | DATETIME | 否 | 完成时间 |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/agvtask/list | AGV任务列表 |
| GET | /wms/agvtask/{id} | AGV任务详情 |
| POST | /wms/agvtask | 创建AGV任务 |
| PUT | /wms/agvtask/{id}/cancel | 取消任务 |
| PUT | /wms/agvtask/{id}/retry | 重试任务 |

---

#### 8.15.4 AGV库位关系 (agvlocationrelation) - 已在7.15节

#### 8.15.5 AGV接口 (agvinterface)

**路由路径**: `/wms/agvinterface`

**功能说明**: AGV接口配置，管理与AGV系统的对接配置。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| interface_id | BIGINT | 是 | 接口ID |
| agv_id | BIGINT | 否 | AGV ID |
| interface_url | VARCHAR(200) | 是 | 接口地址 |
| interface_key | VARCHAR(100) | 否 | 接口密钥 |
| protocol_type | VARCHAR(20) | 是 | 协议类型(HTTP/MQTT/WebSocket) |
| timeout | INT | 否 | 超时时间(秒) |
| retry_times | INT | 否 | 重试次数 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/agvinterface/list | 接口列表 |
| GET | /wms/agvinterface/{id} | 接口详情 |
| POST | /wms/agvinterface | 新增接口 |
| PUT | /wms/agvinterface/{id} | 修改接口 |
| DELETE | /wms/agvinterface/{id} | 删除接口 |
| POST | /wms/agvinterface/{id}/test | 测试接口连接 |

---

### 8.16 MES条码对接

#### 8.16.1 MES集成配置 (mesintegration)

**路由路径**: `/wms/mesintegration`

**功能说明**: MES系统集成配置，管理与MES系统的条码对接参数。

**核心字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| config_id | BIGINT | 是 | 配置ID |
| config_type | VARCHAR(20) | 是 | 配置类型(BARCODE/WORK_ORDER/PRODUCTION) |
| mes_url | VARCHAR(200) | 是 | MES系统地址 |
| api_key | VARCHAR(100) | 否 | API密钥 |
| sync_frequency | INT | 否 | 同步频率(分钟) |
| last_sync_time | DATETIME | 否 | 最后同步时间 |
| status | VARCHAR(20) | 是 | 状态(ACTIVE/INACTIVE) |

**API接口**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /wms/mesintegration/list | 配置列表 |
| GET | /wms/mesintegration/{id} | 配置详情 |
| POST | /wms/mesintegration | 新增配置 |
| PUT | /wms/mesintegration/{id} | 修改配置 |
| DELETE | /wms/mesintegration/{id} | 删除配置 |
| POST | /wms/mesintegration/{id}/sync | 手动触发同步 |
| GET | /wms/mesintegration/{id}/sync-log | 同步日志 |

---

## 9. 关联文档

- [MOM3.0_主设计文档](./MOM3.0_主设计文档.md) - 系统总览
- [MOM3.0_UI设计规范](./MOM3.0_UI设计规范.md) - UI规范详情
- [sfms3.0/WMS模块设计文档](./sfms3.0/WMS模块设计文档.md) - SFMS3.0 WMS参考文档
