# SCP供应链模块设计文档

## 1. 模块概述

SCP（Supply Chain Planning）供应链模块是SFMS3.0的重要组成部分，负责企业采购、销售、供应商等供应链核心业务的执行与管理，连接WMS仓库管理和MES生产执行系统。

**路径**: `win-module-scp`

**子模块结构**:
- `win-module-scp-api` - API接口定义
- `win-module-scp-biz` - 业务实现

---

## 2. 模块职责

SCP模块主要包含以下业务功能域：

### 2.1 采购管理
- 采购订单管理
- 采购计划（MPS/MRP）
- 采购收货与入库
- 采购退换货
- 供应商结算

### 2.2 销售管理
- 销售订单管理
- 销售发货
- 客户结算
- 客户退货

### 2.3 供应商管理
- 供应商档案
- 供应商评价
- 供应商交货记录

### 2.4 客户管理
- 客户档案
- 客户信用管理
- 客户交货预测

### 2.5 数据统计
- MRS统计（物料需求计划统计）
- 采购统计
- 销售统计

---

## 3. 核心类/接口

### 3.1 目录结构

SCP模块采用标准的三层架构（Controller-Service-DAL）：

```
win-module-scp-biz/src/main/java/com/win/module/scp/
    |-- controller/         # REST控制器
    |-- service/            # 业务服务接口与实现
    |-- dal/                # 数据访问层
    |     |-- dataobject/   # DO实体类
    |     |-- mysql/        # Mapper接口
    |-- convert/            # 对象转换
    |-- framework/          # 框架配置
    |-- util/               # 工具类
```

### 3.2 API模块

**API模块**定义对外暴露的接口和枚举：

- `DictTypeConstants` - 字典类型常量（预留）
- `ErrorCodeConstants` - 错误码定义（预留）
- `package-info.java` - 包说明

### 3.3 控制器

控制器层处理HTTP请求，目前SCP模块包含：

- 采购相关控制器
- 销售相关控制器
- 供应商相关控制器
- 客户相关控制器

### 3.4 服务层

服务层实现核心业务逻辑，采用接口+实现类的设计模式：

- `XxxService` - 业务接口
- `XxxServiceImpl` - 业务实现

### 3.5 数据访问层

DAL层负责数据库操作：

- `XxxDO` - 数据对象（继承BaseDO）
- `XxxMapper` - MyBatis Mapper接口

---

## 4. 数据结构

### 4.1 通用设计模式

SCP模块遵循统一的DO设计规范：

```java
@TableName("scp_xxx")  // 表名
@KeySequence("scp_xxx_seq")  // 主键序列
@Data
@EqualsAndHashCode(callSuper = true)
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class XxxDO extends BaseDO {
    // 继承BaseDO获取createTime, updateTime, creator, updater等字段
}
```

### 4.2 典型数据结构

| 业务域 | 表前缀 | 说明 |
|--------|--------|------|
| 采购 | scp_purchase | 采购订单、计划、收货等 |
| 销售 | scp_sale | 销售订单、发货等 |
| 供应商 | scp_supplier | 供应商档案、交货记录 |
| 客户 | scp_customer | 客户档案、信用等 |

---

## 5. 数据流向

```
[ERP/外部系统]
        |
        v
[SCP模块 - Controller]
        |
        v
[SCP模块 - Service] <--> [WMS模块] (物料库存)
        |                      |
        |                      v
        |              [MES模块] (生产需求)
        |
        v
[MySQL数据库]
```

---

## 6. 关键技术实现

### 6.1 技术栈

- **框架**: Spring Boot + MyBatis-Plus
- **数据库**: MySQL
- **安全**: Spring Security + 权限注解
- **API文档**: Swagger/OpenAPI 3

### 6.2 业务特性

- **导入导出**: 支持Excel批量导入导出
- **权限控制**: 基于`@PreAuthorize`的细粒度权限管理
- **操作日志**: 使用`@OperateLog`记录关键操作
- **分页查询**: 统一使用`PageParam`和`PageResult`

### 6.3 模块集成

SCP模块与系统其他模块的集成关系：

```
SCP模块
    |
    +-- 依赖 Infra模块（安全、认证）
    |
    +-- 依赖 WMS模块（库存数据）
    |
    +-- 依赖 System模块（用户、组织数据）
    |
    +-- 被 MES模块调用（物料需求触发采购）
```

---

## 7. 错误码

错误码定义遵循系统规范（1-xxx-xxx-xxx）：

- SCP模块使用独立错误码段
- 错误码定义在`ErrorCodeConstants`接口中

---

## 8. 与其他模块的集成

### 8.1 与WMS集成

- 采购入库：SCP生成采购单 → WMS执行入库
- 销售出库：SCP生成销售单 → WMS执行出库

### 8.2 与MES集成

- 物料需求：MES生产工单 → SCP物料需求计划
- 供应计划：SCP采购计划 → 供应商交货

### 8.3 与System集成

- 用户认证：复用System的OAuth2TokenApi
- 权限校验：使用System的权限框架

---

## 9. 完整API清单

> 说明：以下API均位于 `win-module-scp` 模块中。采购、销售等业务的核心Controller部署在 `win-module-wms` 模块（路径前缀 `/wms/xxx`），与SCP模块共享供应链业务实体和数据库表结构。

### 9.1 采购管理

#### 9.1.1 采购订单主 (PurchaseMainController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| POST | `/wms/purchase-main/create` | 创建采购订单主 |
| PUT | `/wms/purchase-main/update` | 更新采购订单主 |
| DELETE | `/wms/purchase-main/delete` | 删除采购订单主 |
| GET | `/wms/purchase-main/get` | 获得采购订单主 |
| GET | `/wms/purchase-main/list` | 获得采购订单主列表 |
| GET | `/wms/purchase-main/page` | 获得采购订单主分页 |
| POST | `/wms/purchase-main/senior` | 高级搜索获得采购订单主分页 |
| GET | `/wms/purchase-main/export-excel` | 导出采购订单主 Excel |
| GET | `/wms/purchase-main/export-excelWMS` | 从WMS系统导出采购订单主 Excel |
| POST | `/wms/purchase-main/export-excel-senior` | 导出采购订单主 Excel（xml查询） |
| POST | `/wms/purchase-main/export-excel-senior1` | 导出采购订单主 Excel（高级搜索） |
| POST | `/wms/purchase-main/export-excel-seniorWMS` | 导出采购订单主 Excel（WMS xml查询） |
| GET | `/wms/purchase-main/get-import-template` | 获得导入采购订单模板 |
| POST | `/wms/purchase-main/import` | 导入采购订单 |
| POST | `/wms/purchase-main/close` | 关闭采购订单主 |
| POST | `/wms/purchase-main/open` | 打开采购订单主 |
| POST | `/wms/purchase-main/publish` | 发布采购订单主 |
| POST | `/wms/purchase-main/wit` | 下架采购订单主 |

#### 9.1.2 采购订单子 (PurchaseDetailController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| POST | `/wms/purchase-detail/create` | 创建采购订单子 |
| PUT | `/wms/purchase-detail/update` | 更新采购订单子 |
| DELETE | `/wms/purchase-detail/delete` | 删除采购订单子 |
| GET | `/wms/purchase-detail/get` | 获得采购订单子 |
| GET | `/wms/purchase-detail/list` | 获得采购订单子列表 |
| GET | `/wms/purchase-detail/page` | 获得采购订单子分页 |
| POST | `/wms/purchase-detail/senior` | 高级搜索获得采购订单子信息分页 |
| GET | `/wms/purchase-detail/pageWMS` | WMS获得采购订单子分页 |
| POST | `/wms/purchase-detail/seniorWMS` | WMS高级搜索获得采购订单子信息分页 |
| GET | `/wms/purchase-detail/purchasereceiptRequestPageWMS` | WMS获得采购订单子分页（收货申请） |
| POST | `/wms/purchase-detail/purchasereceiptRequestSeniorWMS` | WMS高级搜索采购订单子（收货申请） |
| GET | `/wms/purchase-detail/pageWMS-Spare` | WMS获得采购订单子维修备件收货分页 |
| POST | `/wms/purchase-detail/seniorWMS-Spare` | WMS高级搜索采购订单子维修备件收货分页 |
| GET | `/wms/purchase-detail/pageWMS-MOrderType` | WMS获得采购订单子分页（M型采购订单） |
| POST | `/wms/purchase-detail/seniorWMS-MOrderType` | WMS高级搜索获得采购订单子分页（M型采购订单） |
| GET | `/wms/purchase-detail/export-excel` | 导出采购订单子 Excel |
| GET | `/wms/purchase-detail/selectAll` | 获得采购订单子分页（按masterId查询） |
| GET | `/wms/purchase-detail/pageCheckData` | 获得采购订单子分页（核对数据） |
| GET | `/wms/purchase-detail/pageM` | WMS根据订单号获取M型采购收货订单明细 |
| POST | `/wms/purchase-detail/seniorM` | 高级搜索M型采购收货订单明细 |

---

### 9.2 要货计划管理

#### 9.2.1 要货计划主 (PurchasePlanMainController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| POST | `/wms/purchase-plan-main/create` | 创建要货计划主 |
| PUT | `/wms/purchase-plan-main/update` | 更新要货计划主 |
| PUT | `/wms/purchase-plan-main/updateALL` | 更新要货计划主（除关闭状态） |
| DELETE | `/wms/purchase-plan-main/delete` | 删除要货计划主 |
| GET | `/wms/purchase-plan-main/get` | 获得要货计划主 |
| GET | `/wms/purchase-plan-main/list` | 获得要货计划主列表 |
| GET | `/wms/purchase-plan-main/page` | 获得要货计划主分页 |
| POST | `/wms/purchase-plan-main/senior` | 高级搜索获得要货计划主分页 |
| GET | `/wms/purchase-plan-main/export-excel` | 导出要货计划二维表 |
| GET | `/wms/purchase-plan-main/export-excel-detail` | 导出要货计划明细 |
| POST | `/wms/purchase-plan-main/export-excel-senior` | 导出高级搜索明细 Excel |
| GET | `/wms/purchase-plan-main/get-import-template` | 获得导入要货计划信息模板 |
| POST | `/wms/purchase-plan-main/import` | 导入要货计划 |
| POST | `/wms/purchase-plan-main/close` | 关闭要货计划主 |
| GET | `/wms/purchase-plan-main/close-attachment-files` | 关闭原因附件列表 |
| POST | `/wms/purchase-plan-main/open` | 打开要货计划主 |
| POST | `/wms/purchase-plan-main/supplier-to-confirm` | 供应商待确认（仅准备状态） |
| GET | `/wms/purchase-plan-main/supplier-confirm-detail` | 供应商确认弹窗明细 |
| POST | `/wms/purchase-plan-main/supplier-confirm` | 供应商确认保存 |
| POST | `/wms/purchase-plan-main/planer-confirm` | 计划员确认（仅供应商已确认状态） |
| POST | `/wms/purchase-plan-main/publish` | 发布要货计划主 |
| POST | `/wms/purchase-plan-main/wit` | 下架要货计划主 |
| POST | `/wms/purchase-plan-main/acc` | 接受要货计划主 |
| POST | `/wms/purchase-plan-main/rej` | 驳回要货计划主 |
| GET | `/wms/purchase-plan-main/queryPurchasePlan` | 获取采购计划策略S001 |

#### 9.2.2 要货计划子 (PurchasePlanDetailController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| POST | `/wms/purchase-plan-detail/create` | 创建要货计划子 |
| PUT | `/wms/purchase-plan-detail/update` | 更新要货计划子 |
| DELETE | `/wms/purchase-plan-detail/delete` | 删除要货计划子 |
| GET | `/wms/purchase-plan-detail/get` | 获得要货计划子 |
| GET | `/wms/purchase-plan-detail/list` | 获得要货计划子列表 |
| GET | `/wms/purchase-plan-detail/page` | 获得要货计划子分页 |
| POST | `/wms/purchase-plan-detail/senior` | 高级搜索获得要货计划子信息分页 |
| GET | `/wms/purchase-plan-detail/export-excel` | 导出要货计划子 Excel |
| GET | `/wms/purchase-plan-detail/allList` | 高级搜索获得要货计划子信息 |
| GET | `/wms/purchase-plan-detail/pageWMS` | WMS获得要货计划子分页 |
| POST | `/wms/purchase-plan-detail/seniorWMS` | WMS高级搜索获得要货计划子信息分页 |
| GET | `/wms/purchase-plan-detail/clickDetailsPage` | 供应商发货申请选择要货计划点击明细 |
| POST | `/wms/purchase-plan-detail/clickDetailsSenior` | 供应商发货申请点击明细高级搜索 |

---

### 9.3 供应商发货管理

#### 9.3.1 供应商发货申请主 (SupplierdeliverRequestMainController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| POST | `/wms/supplierdeliver-request-main/create` | 创建供应商发货申请主 |
| PUT | `/wms/supplierdeliver-request-main/update` | 更新供应商发货申请主 |
| DELETE | `/wms/supplierdeliver-request-main/delete` | 删除供应商发货申请主 |
| GET | `/wms/supplierdeliver-request-main/get` | 获得供应商发货申请主 |
| GET | `/wms/supplierdeliver-request-main/list` | 获得供应商发货申请主列表 |
| GET | `/wms/supplierdeliver-request-main/page` | 获得供应商发货申请主分页 |
| POST | `/wms/supplierdeliver-request-main/senior` | 高级搜索获得供应商发货申请主分页 |
| GET | `/wms/supplierdeliver-request-main/export-excel` | 导出供应商发货申请主 Excel |
| POST | `/wms/supplierdeliver-request-main/export-excel-senior` | 导出供应商发货申请主 Excel（高级搜索） |
| GET | `/wms/supplierdeliver-request-main/get-import-template` | 获得导入供应商发货申请信息模板 |
| POST | `/wms/supplierdeliver-request-main/import` | 导入供应商发货申请 |
| POST | `/wms/supplierdeliver-request-main/close` | 关闭供应商发货申请主 |
| POST | `/wms/supplierdeliver-request-main/open` | 打开供应商发货申请主 |
| POST | `/wms/supplierdeliver-request-main/sub` | 提交供应商发货申请主 |
| POST | `/wms/supplierdeliver-request-main/app` | 审批通过供应商发货申请主 |
| POST | `/wms/supplierdeliver-request-main/rej` | 驳回供应商发货申请主 |
| POST | `/wms/supplierdeliver-request-main/selfCheckReport` | 自检报告 |
| POST | `/wms/supplierdeliver-request-main/genLabel` | 生成标签 |
| POST | `/wms/supplierdeliver-request-main/checkPackQty` | 检查包装数量 |
| POST | `/wms/supplierdeliver-request-main/genRecords` | 生成记录 |
| GET | `/wms/supplierdeliver-request-main/queryQualityInspection` | 查询质检明细 |
| GET | `/wms/supplierdeliver-request-main/querySupplierResume` | 查询供应商履历表 |
| POST | `/wms/supplierdeliver-request-main/deleteOldLabels` | 删除之前打印的标签 |

#### 9.3.2 供应商发货申请子 (SupplierdeliverRequestDetailController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| POST | `/wms/supplierdeliver-request-detail/create` | 创建供应商发货申请子 |
| PUT | `/wms/supplierdeliver-request-detail/update` | 更新供应商发货申请子 |
| DELETE | `/wms/supplierdeliver-request-detail/delete` | 删除供应商发货申请子 |
| GET | `/wms/supplierdeliver-request-detail/get` | 获得供应商发货申请子 |
| GET | `/wms/supplierdeliver-request-detail/list` | 获得供应商发货申请子列表 |
| GET | `/wms/supplierdeliver-request-detail/page` | 获得供应商发货申请子分页 |
| POST | `/wms/supplierdeliver-request-detail/senior` | 高级搜索获得供应商发货申请子分页 |
| GET | `/wms/supplierdeliver-request-detail/generateLabelList` | 获得供应商发货申请打印标签 |
| GET | `/wms/supplierdeliver-request-detail/generateLabelParentList` | 获得供应商发货申请打印标签（父级） |
| GET | `/wms/supplierdeliver-request-detail/export-excel` | 导出供应商发货申请子 Excel |

#### 9.3.3 供应商发货记录主 (SupplierdeliverRecordMainController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| POST | `/wms/supplierdeliver-record-main/createMq` | 创建供应商发货记录主（MQ专用） |
| POST | `/wms/supplierdeliver-record-main/create` | 创建供应商发货记录主 |
| PUT | `/wms/supplierdeliver-record-main/update` | 更新供应商发货记录主 |
| DELETE | `/wms/supplierdeliver-record-main/delete` | 删除供应商发货记录主 |
| GET | `/wms/supplierdeliver-record-main/get` | 获得供应商发货记录主 |
| GET | `/wms/supplierdeliver-record-main/list` | 获得供应商发货记录主列表 |
| GET | `/wms/supplierdeliver-record-main/page` | 获得供应商发货记录主分页 |
| POST | `/wms/supplierdeliver-record-main/senior` | 高级搜索获得供应商发货记录主分页 |
| POST | `/wms/supplierdeliver-record-main/abolish` | 作废供应商发货记录 |
| GET | `/wms/supplierdeliver-record-main/export-excel` | 导出供应商发货记录主 Excel |
| POST | `/wms/supplierdeliver-record-main/export-excel-senior` | 导出供应商发货记录主 Excel（高级搜索） |
| GET | `/wms/supplierdeliver-record-main/getSupplierdeliverRecordById` | APP获得供应商发货记录主子表明细列表 |
| POST | `/wms/supplierdeliver-record-main/createPurchasereceiptRequest` | 创建采购收货申请 |
| POST | `/wms/supplierdeliver-record-main/cancelShipmentByAsnNumber` | 撤销发货通过ASN单号 |

---

### 9.4 采购收货/退货管理

#### 9.4.1 采购收货任务主 (PurchasereceiptJobMainController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| GET | `/wms/purchasereceipt-job-main/page` | 获得采购收货任务主分页 |
| POST | `/wms/purchasereceipt-job-main/senior` | 高级搜索获得采购收货任务主分页 |
| GET | `/wms/purchasereceipt-job-main/export-excel` | 导出采购收货任务主 Excel |
| POST | `/wms/purchasereceipt-job-main/export-excel-senior` | 导出采购收货任务主 Excel（高级搜索） |
| GET | `/wms/purchasereceipt-job-main/export-excel-spare` | 导出备件收货任务主 Excel |
| POST | `/wms/purchasereceipt-job-main/export-excel-spare-senior` | 导出备件收货任务主 Excel（高级搜索） |
| GET | `/wms/purchasereceipt-job-main/getPurchasereceiptJobyId` | APP获得采购收货任务主子表明细列表 |
| POST | `/wms/purchasereceipt-job-main/getCountByStatus` | APP获得采购收货任务数量根据任务状态 |
| PUT | `/wms/purchasereceipt-job-main/accept` | 承接任务 |
| PUT | `/wms/purchasereceipt-job-main/abandon` | 放弃任务 |
| PUT | `/wms/purchasereceipt-job-main/close` | 关闭任务 |
| PUT | `/wms/purchasereceipt-job-main/execute` | 执行采购收货任务主 |
| PUT | `/wms/purchasereceipt-job-main/executeSpare` | 执行备件收货任务主 |
| POST | `/wms/purchasereceipt-job-main/refusal` | 拒收任务 |
| GET | `/wms/purchasereceipt-job-main/queryInspectionFreeFlag` | APP获得采购收货任务主子表明细列表（质检免检） |
| PUT | `/wms/purchasereceipt-job-main/updatePurchasereceiptJobConfig` | 更新任务单据设置 |
| PUT | `/wms/purchasereceipt-job-main/updatePurchasereceiptJobConfigSpare` | 维修备件更新任务配置 |

#### 9.4.2 采购退货任务主 (PurchasereturnJobMainController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| GET | `/wms/purchasereturn-job-main/page` | 获得采购退货任务主分页 |
| POST | `/wms/purchasereturn-job-main/senior` | 高级搜索获得采购退货任务分页 |
| GET | `/wms/purchasereturn-job-main/export-excel` | 导出采购退货任务 Excel |
| POST | `/wms/purchasereturn-job-main/export-excel-senior` | 导出采购退货任务 Excel（高级搜索） |
| GET | `/wms/purchasereturn-job-main/getReturnJobById` | APP获得采购退货任务主子表明细列表 |
| POST | `/wms/purchasereturn-job-main/getCountByStatus` | APP获得采购退货任务数量根据任务状态 |
| PUT | `/wms/purchasereturn-job-main/accept` | 承接任务 |
| PUT | `/wms/purchasereturn-job-main/abandon` | 放弃任务 |
| PUT | `/wms/purchasereturn-job-main/close` | 关闭任务 |
| PUT | `/wms/purchasereturn-job-main/execute` | 执行采购退货任务主 |
| PUT | `/wms/purchasereturn-job-main/updatePurchasereturnJobConfig` | 更新任务单据设置 |

#### 9.4.3 采购退货记录主 (PurchasereturnRecordMainController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| POST | `/wms/purchasereturn-record-main/create` | 创建采购退货记录 |
| GET | `/wms/purchasereturn-record-main/page` | 获得采购退货记录主分页 |
| POST | `/wms/purchasereturn-record-main/senior` | 高级搜索获得采购退货记录分页 |
| GET | `/wms/purchasereturn-record-main/export-excel` | 导出采购退货记录 Excel |
| POST | `/wms/purchasereturn-record-main/export-excel-senior` | 导出采购退货记录 Excel（高级搜索） |
| GET | `/wms/purchasereturn-record-main/export-excel-spare` | 导出维修备件退货记录 Excel |
| POST | `/wms/purchasereturn-record-main/export-excel-senior-spare` | 导出维修备件退货记录 Excel（高级搜索） |
| GET | `/wms/purchasereturn-record-main/export-excel-mordertype` | 导出采购退货记录 Excel（M型订单类型） |
| POST | `/wms/purchasereturn-record-main/export-excel-senior-mordertype` | 导出采购退货记录 Excel（M型订单，高级搜索） |
| GET | `/wms/purchasereturn-record-main/export-excel-SCP` | 导出采购退货记录主 Excel（SCP供应商视角） |
| POST | `/wms/purchasereturn-record-main/export-excel-senior-SCP` | 导出采购退货记录主 Excel（SCP供应商视角，高级搜索） |

---

### 9.5 供应商管理

#### 9.5.1 供应商物料 (SupplieritemController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| POST | `/wms/supplieritem/create` | 创建供应商物料 |
| POST | `/wms/supplieritem/createSCP` | 创建供应商物料（SCP） |
| PUT | `/wms/supplieritem/update` | 更新供应商物料 |
| PUT | `/wms/supplieritem/updateSCP` | 更新供应商物料（SCP） |
| DELETE | `/wms/supplieritem/delete` | 删除供应商物料 |
| DELETE | `/wms/supplieritem/deleteSCP` | 删除供应商物料（SCP） |
| GET | `/wms/supplieritem/get` | 获得供应商物料 |
| GET | `/wms/supplieritem/page` | 获得供应商物料分页 |
| POST | `/wms/supplieritem/senior` | 高级搜索获得供应商物料分页 |
| GET | `/wms/supplieritem/pageSCP` | 获得供应商物料分页（SCP供应商视角） |
| POST | `/wms/supplieritem/seniorSCP` | 高级搜索获得供应商物料分页（SCP） |
| GET | `/wms/supplieritem/pageItembasicTypeToSupplieritem` | 获取物料属性筛选出的供应商物料 |
| POST | `/wms/supplieritem/pageItembasicTypeToSupplieritemSenior` | 高级搜索物料属性筛选出的供应商物料 |
| GET | `/wms/supplieritem/export-excel` | 导出供应商物料 Excel |
| POST | `/wms/supplieritem/export-excel-senior` | 导出供应商物料 Excel（高级搜索） |
| GET | `/wms/supplieritem/export-excel-SCP` | 导出供应商物料 Excel（SCP供应商视角） |
| POST | `/wms/supplieritem/export-excel-senior-SCP` | 导出供应商物料 Excel（SCP高级搜索） |
| GET | `/wms/supplieritem/get-import-template` | 获得导入供应商物料信息模板 |
| POST | `/wms/supplieritem/import` | 导入供应商物料信息 |
| POST | `/wms/supplieritem/getDefaultLocationCode` | 获取供应商物料默认库位 |
| GET | `/wms/supplieritem/listByCodes` | 根据代码查询供应商物料基本信息 |
| GET | `/wms/supplieritem/querySupplierByCode` | 根据物料代码查询供应商信息 |
| GET | `/wms/supplieritem/querySupplierByCodeAndType` | 根据物料代码和类型查询供应商信息 |

#### 9.5.2 供应商用户关联 (SupplierUserController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| POST | `/wms/supplier-user/create` | 创建供应商用户关联信息 |
| PUT | `/wms/supplier-user/update` | 更新供应商用户关联信息 |
| DELETE | `/wms/supplier-user/delete` | 删除供应商用户关联信息 |
| GET | `/wms/supplier-user/get` | 获得供应商用户关联信息 |
| GET | `/wms/supplier-user/list` | 获得供应商用户关联信息列表 |
| GET | `/wms/supplier-user/page` | 获得供应商用户关联信息分页 |
| POST | `/wms/supplier-user/senior` | 高级搜索获得供应商用户关联信息分页 |
| GET | `/wms/supplier-user/export-excel` | 导出供应商用户关联信息 Excel |
| POST | `/wms/supplier-user/export-excel-senior` | 导出供应商用户关联信息 Excel（高级搜索） |
| GET | `/wms/supplier-user/get-import-template` | 获得导入供应商用户关联信息模板 |
| POST | `/wms/supplier-user/import` | 导入供应商用户关联信息基本信息 |
| GET | `/wms/supplier-user/getSupplierUserList` | 获取用户信息父子列表（用户树状图展示） |

#### 9.5.3 供应商应付款余额 (SupplierApbalanceMainController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| POST | `/wms/supplier-apbalance-main/create` | 创建供应商余额明细主 |
| PUT | `/wms/supplier-apbalance-main/update` | 更新供应商余额明细主 |
| DELETE | `/wms/supplier-apbalance-main/delete` | 删除供应商余额明细主 |
| GET | `/wms/supplier-apbalance-main/get` | 获得供应商余额明细主 |
| GET | `/wms/supplier-apbalance-main/getPrintInfo` | 获得供应商余额明细主打印信息 |
| GET | `/wms/supplier-apbalance-main/page` | 获得供应商余额明细主分页 |
| POST | `/wms/supplier-apbalance-main/senior` | 高级搜索获得供应商余额明细主分页 |
| GET | `/wms/supplier-apbalance-main/confirmationPage` | 询证查询 |
| GET | `/wms/supplier-apbalance-main/export-excel` | 导出供应商余额明细主 Excel |
| POST | `/wms/supplier-apbalance-main/export-excel-senior` | 导出供应商余额明细主 Excel（高级搜索） |
| GET | `/wms/supplier-apbalance-main/get-import-template` | 获得导入供应商余额明细主模板 |

---

### 9.6 MRS统计/接口

#### 9.6.1 要货计划汇总统计 (PurchaseMrsStatisticsController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| POST | `/wms/purchase-mrs-statistics/create` | 创建要货计划汇总统计给MRS |
| PUT | `/wms/purchase-mrs-statistics/update` | 更新要货计划汇总统计给MRS |
| DELETE | `/wms/purchase-mrs-statistics/delete` | 删除要货计划汇总统计给MRS |
| GET | `/wms/purchase-mrs-statistics/get` | 获得要货计划汇总统计给MRS |
| GET | `/wms/purchase-mrs-statistics/page` | 获得要货计划汇总统计给MRS分页 |
| POST | `/wms/purchase-mrs-statistics/senior` | 高级搜索获得要货计划汇总统计分页 |
| GET | `/wms/purchase-mrs-statistics/export-excel` | 导出要货计划汇总统计给MRS Excel |
| GET | `/wms/purchase-mrs-statistics/get-import-template` | 获得导入要货计划汇总统计给MRS模板 |
| POST | `/wms/purchase-mrs-statistics/import` | 导入要货计划汇总统计给MRS基本信息 |

#### 9.6.2 MRS外部接口 (MrsPurchasePlanMainController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| POST | `/wms/mrs/purchase-plan-main/create` | MRS创建要货计划（独立入参，服务端补全单号/策略仓库等） |

#### 9.6.3 MRS外部接口 - 要货计划汇总统计 (MrsPurchaseMrsStatisticsController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| GET | `/wms/mrs/purchase-mrs-statistics/list` | MRS查询要货计划汇总统计（条件可空，默认查全部可用数据） |

---

### 9.7 客户管理

#### 9.7.1 客户发货预测 (CustomerDeliveryForecastController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| POST | `/wms/customer-delivery-forecast/create` | 创建客户发货预测 |
| PUT | `/wms/customer-delivery-forecast/update` | 更新客户发货预测 |
| DELETE | `/wms/customer-delivery-forecast/delete` | 删除客户发货预测 |
| GET | `/wms/customer-delivery-forecast/get` | 获得客户发货预测 |
| GET | `/wms/customer-delivery-forecast/page` | 获得客户发货预测分页 |
| POST | `/wms/customer-delivery-forecast/senior` | 高级搜索获得客户发货预测分页 |
| GET | `/wms/customer-delivery-forecast/export-excel` | 导出客户发货预测 Excel |
| GET | `/wms/customer-delivery-forecast/get-import-template` | 获得导入客户发货预测模板 |
| POST | `/wms/customer-delivery-forecast/import` | 导入客户发货预测基本信息 |

#### 9.7.2 要货预测子 (DemandforecastingDetailController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| POST | `/wms/demandforecasting-detail/create` | 创建要货预测子 |
| PUT | `/wms/demandforecasting-detail/update` | 更新要货预测子 |
| DELETE | `/wms/demandforecasting-detail/delete` | 删除要货预测子 |
| GET | `/wms/demandforecasting-detail/get` | 获得要货预测子 |
| GET | `/wms/demandforecasting-detail/list` | 获得要货预测子列表 |
| GET | `/wms/demandforecasting-detail/page` | 获得要货预测子分页 |
| GET | `/wms/demandforecasting-detail/export-excel` | 导出要货预测子 Excel |
| GET | `/wms/demandforecasting-detail/queryPageTableHead` | 获取表格头部数据 |
| GET | `/wms/demandforecasting-detail/queryVersion` | 获取版本 |
| POST | `/wms/demandforecasting-detail/queryQADDemandforecasting` | 查询QAD的要货预测数据 |

#### 9.7.3 客户退货任务主 (CustomerreturnJobMainController)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|----------|
| POST | `/wms/customerreturn-job-main/create` | 创建客户退货任务主 |
| PUT | `/wms/customerreturn-job-main/update` | 更新客户退货任务主 |
| DELETE | `/wms/customerreturn-job-main/delete` | 删除客户退货任务主 |
| GET | `/wms/customerreturn-job-main/get` | 获得客户退货任务主 |
| GET | `/wms/customerreturn-job-main/list` | 获得客户退货任务主列表 |
| GET | `/wms/customerreturn-job-main/page` | 获得客户退货任务主分页 |
| POST | `/wms/customerreturn-job-main/senior` | 高级搜索获得客户退货任务主分页 |
| GET | `/wms/customerreturn-job-main/export-excel` | 导出客户退货任务主 Excel |
| POST | `/wms/customerreturn-job-main/export-excel-senior` | 导出客户退货任务主 Excel（高级搜索） |
| GET | `/wms/customerreturn-job-main/getCustomerreturnJobById` | APP获得客户退货任务主子表明细列表 |
| POST | `/wms/customerreturn-job-main/getCountByStatus` | APP获得客户退货任务数量根据任务状态 |
| PUT | `/wms/customerreturn-job-main/accept` | 承接客户退货任务 |
| PUT | `/wms/customerreturn-job-main/abandon` | 取消承接客户退货任务 |
| PUT | `/wms/customerreturn-job-main/close` | 关闭客户退货任务主 |
| PUT | `/wms/customerreturn-job-main/execute` | 执行客户退货任务主 |

---

### 9.8 API统计汇总

| 子模块 | Controller数量 | 主要功能 |
|--------|---------------|----------|
| 采购管理 | 2 | 采购订单主/子 CRUD、导入导出、状态管理 |
| 要货计划管理 | 2 | 要货计划主/子 CRUD、确认流程、发布下架 |
| 供应商发货管理 | 3 | 发货申请主/子、发货记录主、标签生成 |
| 采购收货/退货 | 3 | 收货任务主、退货任务主、退货记录主 |
| 供应商管理 | 3 | 供应商物料、用户关联、应付款余额 |
| MRS统计/接口 | 3 | MRS汇总统计、外部接口（创建/查询） |
| 客户管理 | 3 | 客户发货预测、要货预测子、客户退货任务 |
| **合计** | **19** | **约140+ API端点** |

> 注：SCP模块的API实际承载在 `win-module-wms` 模块的Controller中，通过统一的 `/wms/` 路径前缀对外暴露。`win-module-scp` 模块作为API定义和数据结构层，为上层业务提供公共的枚举、常量和数据结构支撑。
