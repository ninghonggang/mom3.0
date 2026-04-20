# WMS仓库管理模块设计文档

## 1. 模块概述

WMS（Warehouse Management System）仓库管理模块是SFMS3.0的核心业务模块，负责企业仓储全流程管理，涵盖基础档案、库内作业、入出库业务、库存盘点、AGV调度等核心仓储功能。

**路径**: `win-module-wms`

**子模块结构**:
- `win-module-wms-api` - API接口定义
- `win-module-wms-biz` - 业务实现

---

## 2. 模块职责

WMS模块包含约214个业务控制器，216个服务类，是系统中最庞大的模块。主要功能域如下：

### 2.1 基础档案管理

| 子模块 | 职责 |
|--------|------|
| 仓库 (warehouse) | 仓库档案管理 |
| 库区 (areabasic) | 库区配置 |
| 库位 (location) | 库位定义与容量 |
| 货品 (itembasic) | 物料/产品档案 |
| 容器 (container) | 容器规格管理 |
| 包装 (itempackage) | 包装单位 |
| 业务类型 (businesstype) | 业务类型配置 |
| 策略 (strategy) | 业务策略规则 |

### 2.2 库内作业管理

| 子模块 | 职责 |
|--------|------|
| 盘点计划 (countPlan) | 盘点计划制定 |
| 盘点任务 (countJob) | 盘点执行任务 |
| 盘点记录 (countRecord) | 盘点结果记录 |
| 移库作业 (inventorymoveJob) | 库位间移库 |
| 调账作业 (countadjustRequest) | 库存调整 |

### 2.3 入库管理

| 子模块 | 职责 |
|--------|------|
| 采购入库 (purchasereceiptJob) | 采购收货入库 |
| 生产入库 (productreceiptJob) | 生产成品入库 |
| 销售退货 (customerreturnJob) | 客户退货入库 |
| 供应商退货 (purchasereturnJob) | 采购退货 |
| 库内入库 (putawayJob) | 容器入库上架 |
| 退库 (productionreturnJob) | 生产退料 |

### 2.4 出库管理

| 子模块 | 职责 |
|--------|------|
| 销售发货 (saleShipmentRecord) | 销售订单发货 |
| 领料 (issueJob) | 生产领料出库 |
| 采购退料 (purchasereturnRecord) | 采购退料出库 |
| 库内出库 (pickJob) | 拣货下架 |
| 客户退货出库 (customerreturnRecord) | 退货出库 |

### 2.5 库存管理

| 子模块 | 职责 |
|--------|------|
| 库存台账 (balance) | 实时库存 |
| 库存变化 (inventorychangeRecord) | 库存异动记录 |
| 库存初始化 (inventoryinitRecord) | 期初库存 |

### 2.6 容器管理

| 子模块 | 职责 |
|--------|------|
| 容器绑定 (containerBind) | 货品与容器绑定 |
| 容器解绑 (containerUnbind) | 容器释放 |
| 容器初始化 (containerinit) | 容器初始化 |
| 容器维修 (containerRepair) | 容器维修记录 |

### 2.7 AGV调度

| 子模块 | 职责 |
|--------|------|
| AGV任务 (agvService) | AGV作业调度 |
| AGV库位关系 (agvlocationrelation) | AGV与库位映射 |

### 2.8 质检管理

| 子模块 | 职责 |
|--------|------|
| 检验请求 (inspectRequest) | 质检申请 |
| 检验记录 (inspectRecord) | 检验结果 |
| 供应商来料检验 (supplierdeliverinspectiondetail) | IQC |

### 2.9 标签打印

| 子模块 | 职责 |
|--------|------|
| 标签模板 (labelBarbasic) | 标签格式定义 |
| 打印配置 (print) | 打印设置 |
| 条码 (barcode) | 条码规则 |

---

## 3. 核心类/接口

### 3.1 目录结构

```
win-module-wms-biz/src/main/java/com/win/module/wms/
    |-- controller/          # REST控制器 (214个)
    |-- service/              # 业务服务接口与实现 (216个)
    |-- dal/                  # 数据访问层
    |     |-- dataobject/    # DO实体类
    |     |-- mysql/         # Mapper接口
    |-- convert/             # 对象转换 (Dozer/MyBatis-Plus)
    |-- framework/            # 框架配置
    |-- job/                  # 定时任务
    |-- runner/               # 启动执行器
    |-- enums/                # 枚举定义
    |-- util/                 # 工具类
```

### 3.2 典型控制器

**Controller**: `WarehouseController`
- 路径: `controller/warehouse/WarehouseController.java`
- 职责: 仓库档案的CRUD、导入导出、高级搜索

```java
@Tag(name = "管理后台 - 仓库")
@RestController
@RequestMapping("/wms/warehouse")
public class WarehouseController {
    // POST   /wms/warehouse/create           - 创建仓库
    // PUT    /wms/warehouse/update           - 更新仓库
    // DELETE /wms/warehouse/delete          - 删除仓库
    // GET    /wms/warehouse/get              - 获取仓库
    // GET    /wms/warehouse/list            - 获取仓库列表
    // GET    /wms/warehouse/page            - 分页查询
    // POST   /wms/warehouse/senior          - 高级搜索
    // GET    /wms/warehouse/export-excel    - 导出Excel
    // POST   /wms/warehouse/import          - 导入Excel
}
```

### 3.3 核心服务模式

**Service接口**: `WarehouseService`
```java
public interface WarehouseService {
    Long createWarehouse(WarehouseCreateReqVO createReqVO);
    void updateWarehouse(WarehouseUpdateReqVO updateReqVO);
    void deleteWarehouse(Long id);
    WarehouseDO getWarehouse(Long id);
    List<WarehouseDO> getWarehouseList(WarehouseExportReqVO exportReqVO);
    PageResult<WarehouseDO> getWarehousePage(WarehousePageReqVO pageVO);
    // 高级搜索
    PageResult<WarehouseDO> getWarehouseSenior(CustomConditions conditions);
    // Excel导入
    List<WarehouseImportErrorVO> importWarehouseList(List<WarehouseImportExcelVo> list, Integer mode, Boolean updatePart);
}
```

### 3.4 数据对象

**DO**: `WarehouseDO`
```java
@TableName("wms_warehouse")
@KeySequence("wms_warehouse_seq")
@Data
@EqualsAndHashCode(callSuper = true)
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class WarehouseDO extends BaseDO {
    private Long id;
    private String code;          // 仓库编码
    private String name;          // 仓库名称
    private String address;       // 地址
    private Integer status;       // 状态
    private String businessType;  // 业务类型
    // ... 更多字段
}
```

---

## 4. 数据结构

### 4.1 通用设计模式

WMS模块遵循统一的DO设计规范：

```java
@TableName("wms_xxx")      // 表名前缀 wms_
@KeySequence("wms_xxx_seq") // 主键序列
@Data
@EqualsAndHashCode(callSuper = true)
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class XxxDO extends BaseDO {
    // 继承BaseDO获取: id, createTime, updateTime, creator, updater, deleted, tenantId
}
```

### 4.2 核心表结构

| 业务域 | 表前缀 | 说明 |
|--------|--------|------|
| 仓库 | wms_warehouse | 仓库档案 |
| 库区 | wms_areabasic | 库区配置 |
| 库位 | wms_location | 库位定义 |
| 物料 | wms_itembasic | 物料档案 |
| 容器 | wms_container_main / detail | 容器主从 |
| 库存 | wms_balance | 库存台账 |
| 入库 | wms_putaway_job / record | 上架作业/记录 |
| 出库 | wms_pick_job / record | 拣货作业/记录 |

### 4.3 VO对象

| VO类型 | 说明 |
|--------|------|
| XxxCreateReqVO | 创建请求 |
| XxxUpdateReqVO | 更新请求 |
| XxxRespVO | 响应对象 |
| XxxPageReqVO | 分页请求 |
| XxxExportReqVO | 导出请求 |
| XxxImportExcelVo | 导入模板 |
| XxxImportErrorVO | 导入错误 |

---

## 5. 数据流向

### 5.1 入库流程

```
[采购收货] --> purchasereceiptJob --> [验货] --> [上架] --> putawayJob --> [库存台账] balance
                                          |
                                          v
                                   inventorychangeRecord
```

### 5.2 出库流程

```
[销售订单] --> [拣货下架] pickJob --> [出库] --> [库存扣减] --> inventorychangeRecord
                                                    |
                                                    v
                                              balance
```

### 5.3 盘点流程

```
[盘点计划] countPlan --> [盘点任务] countJob --> [盘点执行] countRecord --> [库存调整] countadjustRequest
```

### 5.4 整体数据流

```
[SCP采购模块]  --> 入库 --> [WMS模块] --> [MES生产模块]
                                    |
                                    v
                            [库存台账 balance]
                                    |
                                    v
[SCP销售模块]  --> 出库 --> [Report报表模块]
```

---

## 6. 关键技术实现

### 6.1 技术栈

- **框架**: Spring Boot + MyBatis-Plus
- **数据库**: MySQL
- **安全**: Spring Security + 权限注解
- **导入导出**: 自定义Excel工具（ExcelUtils）
- **API文档**: Swagger/OpenAPI 3
- **AGV集成**: RESTful API

### 6.2 高级搜索

使用`CustomConditions`实现动态条件组合：

```java
public PageResult<WarehouseDO> getWarehouseSenior(CustomConditions conditions) {
    // 动态拼接SQL条件
}
```

### 6.3 幂等控制

使用`@Idempotent`注解防止重复提交：

```java
@PostMapping("/create")
@Idempotent(timeout = 60, message = "正在执行中，请勿重复提交")
public CommonResult<Long> createWarehouse(...) { }
```

### 6.4 导入校验

```java
public List<WarehouseImportErrorVO> importWarehouseList(
    List<WarehouseImportExcelVo> list, Integer mode, Boolean updatePart) {
    // 1. 数据校验
    // 2. 错误收集
    // 3. 返回错误列表
}
```

### 6.5 权限控制

使用`@PreAuthorize("@ss.hasPermission('wms:warehouse:create')")`注解实现细粒度权限控制。

### 6.6 业务类型驱动

仓库和业务类型关联，实现不同业务类型的仓库筛选：

```java
@GetMapping("/pageBusinessTypeToWarehouse")
public PageResult<WarehouseRespVO> selectBusinessTypeToWarehouse(
    @RequestParam("businessType") String businessType,
    @RequestParam("isIn") String isIn) {
    // 根据业务类型筛选可用仓库
}
```

---

## 7. 集成关系

### 7.1 与SCP模块

- **采购入库**: SCP采购单 → WMS执行收货上架
- **销售出库**: SCP销售单 → WMS执行拣货出库
- **采购退货**: WMS退货 → SCP供应商退货单

### 7.2 与MES模块

- **生产领料**: MES工单 → WMS领料出库
- **生产入库**: MES报工 → WMS成品入库

### 7.3 与Report模块

- WMS提供库存数据供Report模块查询展示

### 7.4 与System模块

- 用户认证: 复用System的OAuth2TokenApi
- 权限校验: 使用System的权限框架

---

## 8. 错误码

错误码定义遵循系统规范，WMS模块使用独立错误码段。

---

## 9. 模块特点

### 9.1 规模最大

WMS是SFMS3.0中规模最大的模块：
- 214个控制器
- 216个服务类
- 涵盖仓储全流程

### 9.2 业务复杂

- 多种业务类型（采购、销售、生产、调拨等）
- 主从表结构（Job/Record分离）
- 库存实时记账

### 9.3 高度可配置

- 业务类型驱动
- 策略规则引擎
- 打印标签可自定义

---

## 10. 补充设计

### 10.1 补充DDL表结构

#### 10.1.1 wms_areabasic - 库区基础表（仓库-库区-库位三级结构）

> 管理仓库、库区、库位的层级结构关系，支持按仓库→库区→库位逐级细化。

```sql
CREATE TABLE `wms_areabasic` (
  `id`            bigint       NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `code`          varchar(50)  NOT NULL                   COMMENT '库区编码',
  `name`          varchar(100) NOT NULL                   COMMENT '库区名称',
  `warehouse_code` varchar(50)  NOT NULL                   COMMENT '所属仓库编码',
  `area_type`     varchar(20)  DEFAULT NULL               COMMENT '库区类型：STORAGE存储区/PICKING拣货区/RECEIVING收货区/SHIPPING发货区',
  `parent_code`   varchar(50)  DEFAULT NULL               COMMENT '上级库区编码（支持多级）',
  `level`         tinyint      DEFAULT 1                   COMMENT '层级：1仓库/2库区/3库位',
  `status`        tinyint      DEFAULT 1                   COMMENT '状态：0禁用/1启用',
  `remark`        varchar(500) DEFAULT NULL               COMMENT '备注',
  `deleted`       bit(1)       DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`     bigint       DEFAULT NULL                COMMENT '租户编号',
  `create_time`   datetime     DEFAULT NULL                COMMENT '创建时间',
  `update_time`   datetime     DEFAULT NULL                COMMENT '更新时间',
  `creator`       varchar(64)   DEFAULT ''                   COMMENT '创建者',
  `updater`       varchar(64)   DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_code` (`code`, `tenant_id`, `deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库区基础表';
```

#### 10.1.2 wms_itembasic - 货品基础表（物料与货品对应关系）

> 建立MDM物料与WMS货品的映射关系，维护物料在仓库中的货品化档案。

```sql
CREATE TABLE `wms_itembasic` (
  `id`              bigint       NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `item_code`       varchar(50)  NOT NULL                   COMMENT '货品编码',
  `item_name`       varchar(200) NOT NULL                   COMMENT '货品名称',
  `material_code`   varchar(50)  DEFAULT NULL               COMMENT '对应物料编码（MDM）',
  `material_name`   varchar(200) DEFAULT NULL               COMMENT '对应物料名称',
  `warehouse_code`  varchar(50)  DEFAULT NULL               COMMENT '默认仓库编码',
  `default_area`    varchar(50)  DEFAULT NULL               COMMENT '默认库区编码',
  `default_location` varchar(50) DEFAULT NULL               COMMENT '默认库位编码',
  `item_type`       varchar(20)  DEFAULT NULL               COMMENT '货品类型：RAW原材料/SEMI半成品/FG成品/MATERIAL辅料',
  `unit`            varchar(20)  DEFAULT NULL               COMMENT '计量单位',
  `spec`            varchar(100) DEFAULT NULL               COMMENT '规格型号',
  `batch_enabled`   bit(1)      DEFAULT b'0'                COMMENT '是否启用批次管理',
  `shelf_life_days` int          DEFAULT NULL               COMMENT '保质期天数',
  `min_stock`       decimal(18,6) DEFAULT NULL             COMMENT '最小库存量',
  `max_stock`       decimal(18,6) DEFAULT NULL               COMMENT '最大库存量',
  `status`          tinyint      DEFAULT 1                   COMMENT '状态：0禁用/1启用',
  `remark`          varchar(500) DEFAULT NULL               COMMENT '备注',
  `deleted`         bit(1)       DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`       bigint       DEFAULT NULL                COMMENT '租户编号',
  `create_time`     datetime     DEFAULT NULL                COMMENT '创建时间',
  `update_time`     datetime     DEFAULT NULL                COMMENT '更新时间',
  `creator`         varchar(64)   DEFAULT ''                   COMMENT '创建者',
  `updater`         varchar(64)   DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_item_code` (`item_code`, `tenant_id`, `deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='货品基础表';
```

#### 10.1.3 wms_container_main - 容器主表

> 管理容器（托盘、周转箱等）的整体档案及当前状态。

```sql
CREATE TABLE `wms_container_main` (
  `id`              bigint        NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `container_code`  varchar(100)  NOT NULL                   COMMENT '容器编码（如托盘号）',
  `container_type`  varchar(50)   DEFAULT NULL               COMMENT '容器类型：TRAY托盘/BOX周转箱/Pallet托盘',
  `spec`            varchar(100)  DEFAULT NULL               COMMENT '规格（长*宽*高）',
  `max_capacity`    decimal(18,4) DEFAULT NULL               COMMENT '最大容量（件数）',
  `max_weight`      decimal(18,4) DEFAULT NULL               COMMENT '最大承重（kg）',
  `warehouse_code`  varchar(50)   DEFAULT NULL               COMMENT '所属仓库编码',
  `area_code`       varchar(50)   DEFAULT NULL               COMMENT '所属库区编码',
  `location_code`   varchar(50)   DEFAULT NULL               COMMENT '当前库位编码',
  `status`          varchar(20)   DEFAULT 'IDLE'             COMMENT '状态：IDLE空闲/IN_USE使用中/MAINTENANCE维修/BLOCKED冻结',
  `bind_item_code`  varchar(50)   DEFAULT NULL               COMMENT '绑定货品编码',
  `bind_batch`      varchar(50)   DEFAULT NULL               COMMENT '绑定批次号',
  `remark`          varchar(500)  DEFAULT NULL               COMMENT '备注',
  `deleted`         bit(1)        DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`       bigint        DEFAULT NULL                COMMENT '租户编号',
  `create_time`     datetime      DEFAULT NULL                COMMENT '创建时间',
  `update_time`     datetime      DEFAULT NULL                COMMENT '更新时间',
  `creator`         varchar(64)   DEFAULT ''                   COMMENT '创建者',
  `updater`         varchar(64)   DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_container_code` (`container_code`, `tenant_id`, `deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='容器主表';
```

#### 10.1.4 wms_container_detail - 容器明细表

> 记录容器内各货品的明细数量，支持一容器多货品。

```sql
CREATE TABLE `wms_container_detail` (
  `id`              bigint        NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `container_id`    bigint        NOT NULL                   COMMENT '容器主表ID',
  `container_code`  varchar(100)  NOT NULL                   COMMENT '容器编码',
  `item_code`       varchar(50)   NOT NULL                   COMMENT '货品编码',
  `item_name`       varchar(200)  DEFAULT NULL               COMMENT '货品名称',
  `batch_no`        varchar(50)   DEFAULT NULL               COMMENT '批次号',
  `quantity`        decimal(18,6) NOT NULL  DEFAULT 0         COMMENT '数量',
  `unit`            varchar(20)   DEFAULT NULL               COMMENT '单位',
  `production_date` datetime      DEFAULT NULL               COMMENT '生产日期',
  `expire_date`     datetime      DEFAULT NULL               COMMENT '有效期至',
  `location_code`   varchar(50)   DEFAULT NULL               COMMENT '库位编码',
  `status`          tinyint       DEFAULT 1                   COMMENT '状态：0无效/1有效',
  `deleted`         bit(1)        DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`       bigint        DEFAULT NULL                COMMENT '租户编号',
  `create_time`     datetime      DEFAULT NULL                COMMENT '创建时间',
  `update_time`     datetime      DEFAULT NULL                COMMENT '更新时间',
  `creator`         varchar(64)   DEFAULT ''                   COMMENT '创建者',
  `updater`         varchar(64)   DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_container_id` (`container_id`),
  KEY `idx_item_code` (`item_code`),
  KEY `idx_batch_no` (`batch_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='容器明细表';
```

#### 10.1.5 wms_pick_job - 拣货作业表

> 记录销售发货、生产领料等出库场景的拣货下架任务。

```sql
CREATE TABLE `wms_pick_job` (
  `id`                bigint        NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `job_no`            varchar(50)   NOT NULL                   COMMENT '拣货单号',
  `source_type`       varchar(30)  NOT NULL                   COMMENT '来源类型：SALE销售/ISSUE领料/RETURN退货',
  `source_no`         varchar(50)  DEFAULT NULL               COMMENT '来源单号（销售单/工单等）',
  `warehouse_code`    varchar(50)  NOT NULL                   COMMENT '仓库编码',
  `area_code`         varchar(50)  DEFAULT NULL               COMMENT '库区编码',
  `location_code`     varchar(50)  DEFAULT NULL               COMMENT '目标库位',
  `container_code`    varchar(100) DEFAULT NULL               COMMENT '容器编码',
  `priority`          tinyint      DEFAULT 5                   COMMENT '优先级（1最高）',
  `status`            varchar(20)  DEFAULT 'PENDING'          COMMENT '状态：PENDING待执行/IN_PROGRESS执行中/COMPLETED已完成/CANCELLED已取消',
  `total_qty`         decimal(18,6) DEFAULT 0                  COMMENT '总数量',
  `picked_qty`        decimal(18,6) DEFAULT 0                  COMMENT '已拣数量',
  `picker`            varchar(64)  DEFAULT NULL               COMMENT '拣货员',
  `pick_time`        datetime      DEFAULT NULL               COMMENT '拣货时间',
  `complete_time`    datetime      DEFAULT NULL               COMMENT '完成时间',
  `remark`            varchar(500) DEFAULT NULL               COMMENT '备注',
  `deleted`           bit(1)       DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`         bigint       DEFAULT NULL                COMMENT '租户编号',
  `create_time`       datetime     DEFAULT NULL                COMMENT '创建时间',
  `update_time`       datetime     DEFAULT NULL                COMMENT '更新时间',
  `creator`           varchar(64)  DEFAULT ''                   COMMENT '创建者',
  `updater`           varchar(64)  DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_job_no` (`job_no`, `tenant_id`),
  KEY `idx_source_no` (`source_no`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='拣货作业表';
```

#### 10.1.6 wms_pick_record - 拣货记录表

> 记录拣货作业的执行明细，作为出库的正式凭证。

```sql
CREATE TABLE `wms_pick_record` (
  `id`              bigint        NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `record_no`       varchar(50)   NOT NULL                   COMMENT '拣货记录单号',
  `job_id`          bigint        NOT NULL                   COMMENT '作业表ID',
  `job_no`          varchar(50)   NOT NULL                   COMMENT '作业单号',
  `item_code`       varchar(50)   NOT NULL                   COMMENT '货品编码',
  `item_name`       varchar(200)  DEFAULT NULL               COMMENT '货品名称',
  `batch_no`        varchar(50)   DEFAULT NULL               COMMENT '批次号',
  `from_location`   varchar(50)    NOT NULL                   COMMENT '拣货库位',
  `container_code`  varchar(100)  DEFAULT NULL               COMMENT '容器编码',
  `pick_qty`        decimal(18,6) NOT NULL  DEFAULT 0         COMMENT '拣货数量',
  `unit`            varchar(20)   DEFAULT NULL               COMMENT '单位',
  `pick_time`       datetime      NOT NULL                   COMMENT '拣货时间',
  `picker`          varchar(64)   NOT NULL                   COMMENT '拣货员',
  `status`          tinyint       DEFAULT 1                   COMMENT '状态：0无效/1有效',
  `deleted`         bit(1)        DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`       bigint        DEFAULT NULL                COMMENT '租户编号',
  `create_time`     datetime      DEFAULT NULL                COMMENT '创建时间',
  `update_time`     datetime      DEFAULT NULL                COMMENT '更新时间',
  `creator`         varchar(64)   DEFAULT ''                   COMMENT '创建者',
  `updater`         varchar(64)   DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_job_id` (`job_id`),
  KEY `idx_pick_time` (`pick_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='拣货记录表';
```

#### 10.1.7 wms_putaway_job - 上架作业表

> 记录验货后货品入库上架的执行任务。

```sql
CREATE TABLE `wms_putaway_job` (
  `id`                bigint        NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `job_no`            varchar(50)   NOT NULL                   COMMENT '上架单号',
  `source_type`       varchar(30)  NOT NULL                   COMMENT '来源类型：PURCHASE采购/PRODUCT生产/RETURN退货',
  `source_no`         varchar(50)  DEFAULT NULL               COMMENT '来源单号',
  `warehouse_code`    varchar(50)  NOT NULL                   COMMENT '仓库编码',
  `target_area`       varchar(50)  DEFAULT NULL               COMMENT '目标库区',
  `target_location`   varchar(50)  DEFAULT NULL               COMMENT '目标库位',
  `container_code`    varchar(100) DEFAULT NULL               COMMENT '容器编码',
  `priority`          tinyint      DEFAULT 5                   COMMENT '优先级',
  `status`            varchar(20)  DEFAULT 'PENDING'            COMMENT '状态：PENDING待上架/IN_PROGRESS上架中/COMPLETED已完成/CANCELLED已取消',
  `total_qty`         decimal(18,6) DEFAULT 0                  COMMENT '总数量',
  `putaway_qty`       decimal(18,6) DEFAULT 0                  COMMENT '已上架数量',
  `putaway_by`        varchar(64)  DEFAULT NULL               COMMENT '上架人',
  `putaway_time`      datetime     DEFAULT NULL               COMMENT '上架时间',
  `complete_time`     datetime     DEFAULT NULL               COMMENT '完成时间',
  `remark`            varchar(500) DEFAULT NULL               COMMENT '备注',
  `deleted`           bit(1)       DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`         bigint       DEFAULT NULL                COMMENT '租户编号',
  `create_time`       datetime     DEFAULT NULL                COMMENT '创建时间',
  `update_time`       datetime     DEFAULT NULL                COMMENT '更新时间',
  `creator`           varchar(64)  DEFAULT ''                   COMMENT '创建者',
  `updater`           varchar(64)  DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_job_no` (`job_no`, `tenant_id`),
  KEY `idx_source_no` (`source_no`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='上架作业表';
```

#### 10.1.8 wms_putaway_record - 上架记录表

> 记录上架执行明细，作为入库的正式凭证。

```sql
CREATE TABLE `wms_putaway_record` (
  `id`              bigint        NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `record_no`       varchar(50)   NOT NULL                   COMMENT '上架记录单号',
  `job_id`          bigint        NOT NULL                   COMMENT '作业表ID',
  `job_no`          varchar(50)   NOT NULL                   COMMENT '作业单号',
  `item_code`       varchar(50)   NOT NULL                   COMMENT '货品编码',
  `item_name`       varchar(200)  DEFAULT NULL               COMMENT '货品名称',
  `batch_no`        varchar(50)   DEFAULT NULL               COMMENT '批次号',
  `to_location`     varchar(50)   NOT NULL                   COMMENT '上架库位',
  `container_code`  varchar(100)  DEFAULT NULL               COMMENT '容器编码',
  `putaway_qty`     decimal(18,6) NOT NULL  DEFAULT 0         COMMENT '上架数量',
  `unit`            varchar(20)   DEFAULT NULL               COMMENT '单位',
  `putaway_time`    datetime      NOT NULL                   COMMENT '上架时间',
  `putaway_by`      varchar(64)   NOT NULL                   COMMENT '上架人',
  `status`          tinyint       DEFAULT 1                   COMMENT '状态：0无效/1有效',
  `deleted`         bit(1)        DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`       bigint        DEFAULT NULL                COMMENT '租户编号',
  `create_time`     datetime      DEFAULT NULL                COMMENT '创建时间',
  `update_time`     datetime      DEFAULT NULL                COMMENT '更新时间',
  `creator`         varchar(64)   DEFAULT ''                   COMMENT '创建者',
  `updater`         varchar(64)   DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_job_id` (`job_id`),
  KEY `idx_putaway_time` (`putaway_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='上架记录表';
```

#### 10.1.9 wms_count_plan - 盘点计划表

> 制定盘点范围、周期和规则，生成盘点任务。

```sql
CREATE TABLE `wms_count_plan` (
  `id`              bigint        NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `plan_no`         varchar(50)   NOT NULL                   COMMENT '盘点计划编号',
  `plan_name`       varchar(200)  NOT NULL                   COMMENT '盘点计划名称',
  `warehouse_code`  varchar(50)   NOT NULL                   COMMENT '仓库编码',
  `area_codes`      varchar(500)  DEFAULT NULL               COMMENT '库区编码列表（JSON）',
  `count_type`      varchar(20)   NOT NULL                   COMMENT '盘点类型：FULL全盘/PARTIAL抽盘/CYCLE循环盘点',
  `count_cycle`     varchar(20)   DEFAULT NULL               COMMENT '盘点周期：DAILY日/MONTHLY月/QUARTERLY季度',
  `plan_start_date` datetime      DEFAULT NULL               COMMENT '计划开始日期',
  `plan_end_date`   datetime      DEFAULT NULL               COMMENT '计划结束日期',
  `status`          varchar(20)   DEFAULT 'DRAFT'             COMMENT '状态：DRAFT草稿/SUBMITTED已提交/APPROVED已审批/PUBLISHED已发布/IN_PROGRESS盘点中/COMPLETED已完成/CLOSED已关闭',
  `job_count`       int           DEFAULT 0                  COMMENT '生成任务数',
  `complated_count` int           DEFAULT 0                  COMMENT '已完成任务数',
  `remark`          varchar(500)  DEFAULT NULL               COMMENT '备注',
  `deleted`         bit(1)        DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`       bigint        DEFAULT NULL                COMMENT '租户编号',
  `create_time`     datetime      DEFAULT NULL                COMMENT '创建时间',
  `update_time`     datetime      DEFAULT NULL                COMMENT '更新时间',
  `creator`         varchar(64)   DEFAULT ''                   COMMENT '创建者',
  `updater`         varchar(64)   DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_plan_no` (`plan_no`, `tenant_id`),
  KEY `idx_status` (`status`),
  KEY `idx_plan_date` (`plan_start_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='盘点计划表';
```

#### 10.1.10 wms_count_job - 盘点任务表

> 盘点计划下发的具体盘点任务，分配给执行人。

```sql
CREATE TABLE `wms_count_job` (
  `id`              bigint        NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `job_no`          varchar(50)   NOT NULL                   COMMENT '盘点任务编号',
  `plan_id`         bigint        NOT NULL                   COMMENT '盘点计划ID',
  `plan_no`         varchar(50)   NOT NULL                   COMMENT '盘点计划编号',
  `warehouse_code`  varchar(50)   NOT NULL                   COMMENT '仓库编码',
  `area_code`       varchar(50)   NOT NULL                   COMMENT '库区编码',
  `location_code`   varchar(50)   DEFAULT NULL               COMMENT '库位编码（精确定位）',
  `item_code`       varchar(50)   DEFAULT NULL               COMMENT '货品编码（抽盘时指定）',
  `count_by`        varchar(64)   DEFAULT NULL               COMMENT '盘点负责人',
  `status`          varchar(20)   DEFAULT 'PENDING'           COMMENT '状态：PENDING待盘点/IN_PROGRESS盘点中/SUBMITTED已提交/APPROVED已审批/COMPLETED已完成',
  `expected_qty`    decimal(18,6) DEFAULT NULL               COMMENT '账面数量',
  `actual_qty`      decimal(18,6) DEFAULT NULL               COMMENT '实盘数量',
  `diff_qty`        decimal(18,6) DEFAULT NULL               COMMENT '差异数量',
  `start_time`      datetime      DEFAULT NULL               COMMENT '开始时间',
  `complete_time`   datetime      DEFAULT NULL               COMMENT '完成时间',
  `remark`          varchar(500)  DEFAULT NULL               COMMENT '备注',
  `deleted`         bit(1)        DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`       bigint        DEFAULT NULL                COMMENT '租户编号',
  `create_time`     datetime      DEFAULT NULL                COMMENT '创建时间',
  `update_time`     datetime      DEFAULT NULL                COMMENT '更新时间',
  `creator`         varchar(64)   DEFAULT ''                   COMMENT '创建者',
  `updater`         varchar(64)   DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_job_no` (`job_no`, `tenant_id`),
  KEY `idx_plan_id` (`plan_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='盘点任务表';
```

#### 10.1.11 wms_count_record - 盘点记录表

> 记录盘点执行的明细数据（逐行记录每次盘点结果）。

```sql
CREATE TABLE `wms_count_record` (
  `id`              bigint        NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `record_no`       varchar(50)   NOT NULL                   COMMENT '盘点记录编号',
  `job_id`          bigint        NOT NULL                   COMMENT '盘点任务ID',
  `job_no`          varchar(50)  NOT NULL                   COMMENT '盘点任务编号',
  `item_code`       varchar(50)   NOT NULL                   COMMENT '货品编码',
  `item_name`       varchar(200)  DEFAULT NULL               COMMENT '货品名称',
  `batch_no`       varchar(50)   DEFAULT NULL               COMMENT '批次号',
  `location_code`   varchar(50)   NOT NULL                   COMMENT '库位编码',
  `expected_qty`   decimal(18,6) DEFAULT NULL               COMMENT '账面数量',
  `actual_qty`     decimal(18,6) NOT NULL                   COMMENT '实盘数量',
  `diff_qty`       decimal(18,6) DEFAULT NULL               COMMENT '差异数量（actual-expected）',
  `diff_rate`       decimal(10,4) DEFAULT NULL               COMMENT '差异率',
  `unit`            varchar(20)   DEFAULT NULL               COMMENT '单位',
  `count_time`      datetime      NOT NULL                   COMMENT '盘点时间',
  `count_by`        varchar(64)   NOT NULL                   COMMENT '盘点人',
  `remark`          varchar(500)  DEFAULT NULL               COMMENT '备注',
  `status`          tinyint       DEFAULT 1                   COMMENT '状态：0无效/1有效',
  `deleted`         bit(1)        DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`       bigint        DEFAULT NULL                COMMENT '租户编号',
  `create_time`     datetime      DEFAULT NULL                COMMENT '创建时间',
  `update_time`     datetime      DEFAULT NULL                COMMENT '更新时间',
  `creator`         varchar(64)   DEFAULT ''                   COMMENT '创建者',
  `updater`         varchar(64)   DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_job_id` (`job_id`),
  KEY `idx_count_time` (`count_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='盘点记录表';
```

#### 10.1.12 wms_inventorychange_record - 库存异动记录表

> 记录所有库存变动的流水，是库存追溯和账务核对的核心。

```sql
CREATE TABLE `wms_inventorychange_record` (
  `id`                bigint         NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `record_no`         varchar(50)    NOT NULL                   COMMENT '异动单号',
  `change_type`       varchar(30)   NOT NULL                   COMMENT '异动类型：IN入库/OUT出库/TRANSFER移库/ADJUST调整/LOCK锁定/UNLOCK解锁',
  `business_type`     varchar(30)   NOT NULL                   COMMENT '业务类型：PURCHASE采购/PRODUCT生产/SALE销售/ISSUE领料/TRANSFER调拨/COUNT盘点/ADJUST调账',
  `source_no`         varchar(50)   DEFAULT NULL               COMMENT '来源单号',
  `warehouse_code`     varchar(50)   NOT NULL                   COMMENT '仓库编码',
  `area_code`          varchar(50)   DEFAULT NULL               COMMENT '库区编码',
  `location_code`     varchar(50)   DEFAULT NULL               COMMENT '库位编码',
  `item_code`          varchar(50)   NOT NULL                   COMMENT '货品编码',
  `item_name`          varchar(200)  DEFAULT NULL               COMMENT '货品名称',
  `batch_no`           varchar(50)   DEFAULT NULL               COMMENT '批次号',
  `container_code`     varchar(100)  DEFAULT NULL               COMMENT '容器编码',
  `change_qty`         decimal(18,6) NOT NULL                   COMMENT '异动数量（正数入库/负数出库）',
  `balance_before`     decimal(18,6) DEFAULT NULL               COMMENT '变动前余额',
  `balance_after`      decimal(18,6) DEFAULT NULL               COMMENT '变动后余额',
  `unit`               varchar(20)   DEFAULT NULL               COMMENT '单位',
  `change_reason`      varchar(200)  DEFAULT NULL               COMMENT '异动原因',
  `operator`           varchar(64)   DEFAULT NULL               COMMENT '操作人',
  `operate_time`       datetime      NOT NULL                   COMMENT '操作时间',
  `remark`             varchar(500)  DEFAULT NULL               COMMENT '备注',
  `deleted`            bit(1)        DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`          bigint        DEFAULT NULL                COMMENT '租户编号',
  `create_time`        datetime      DEFAULT NULL                COMMENT '创建时间',
  `update_time`        datetime      DEFAULT NULL                COMMENT '更新时间',
  `creator`            varchar(64)   DEFAULT ''                   COMMENT '创建者',
  `updater`            varchar(64)   DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_record_no` (`record_no`, `tenant_id`),
  KEY `idx_item_code` (`item_code`),
  KEY `idx_source_no` (`source_no`),
  KEY `idx_change_type` (`change_type`),
  KEY `idx_operate_time` (`operate_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存异动记录表';
```

#### 10.1.13 wms_agv_task - AGV任务表

> 管理AGV搬运任务的完整生命周期（任务下发→执行→完成）。

```sql
CREATE TABLE `wms_agv_task` (
  `id`              bigint        NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `task_no`         varchar(50)   NOT NULL                   COMMENT 'AGV任务编号',
  `task_type`       varchar(30)  NOT NULL                   COMMENT '任务类型：PUTAWAY上架/PICK出库/TRANSFER移库/CHARGE充电/RETURN归还',
  `priority`        tinyint      DEFAULT 5                   COMMENT '优先级（1最高）',
  `warehouse_code`  varchar(50)  NOT NULL                   COMMENT '仓库编码',
  `from_location`   varchar(50)  DEFAULT NULL               COMMENT '起始库位',
  `to_location`     varchar(50)  DEFAULT NULL               COMMENT '目标库位',
  `container_code`  varchar(100) DEFAULT NULL               COMMENT '容器编码',
  `agv_code`        varchar(50)  DEFAULT NULL               COMMENT 'AGV编号',
  `agv_status`      varchar(20)  DEFAULT 'WAITING'            COMMENT 'AGV状态：WAITING等待/ASSIGNED已分配/RUNNING运行中/COMPLETED完成/FAILED失败',
  `dispatch_time`   datetime     DEFAULT NULL               COMMENT '下发时间',
  `start_time`      datetime     DEFAULT NULL               COMMENT '开始执行时间',
  `complete_time`   datetime     DEFAULT NULL               COMMENT '完成时间',
  `error_code`      varchar(50)  DEFAULT NULL               COMMENT '错误代码',
  `error_msg`       varchar(500) DEFAULT NULL               COMMENT '错误信息',
  `retry_count`     tinyint      DEFAULT 0                   COMMENT '重试次数',
  `remark`          varchar(500) DEFAULT NULL               COMMENT '备注',
  `deleted`         bit(1)       DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`       bigint       DEFAULT NULL                COMMENT '租户编号',
  `create_time`     datetime     DEFAULT NULL                COMMENT '创建时间',
  `update_time`     datetime     DEFAULT NULL                COMMENT '更新时间',
  `creator`         varchar(64)  DEFAULT ''                   COMMENT '创建者',
  `updater`         varchar(64)  DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_task_no` (`task_no`, `tenant_id`),
  KEY `idx_agv_code` (`agv_code`),
  KEY `idx_status` (`agv_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='AGV任务表';
```

#### 10.1.14 wms_agv_job - AGV作业表

> 记录AGV的具体搬运作业明细，作为AGV执行的流水台账。

```sql
CREATE TABLE `wms_agv_job` (
  `id`              bigint        NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `job_no`          varchar(50)   NOT NULL                   COMMENT 'AGV作业编号',
  `task_id`         bigint        NOT NULL                   COMMENT 'AGV任务ID',
  `task_no`         varchar(50)   NOT NULL                   COMMENT 'AGV任务编号',
  `agv_code`        varchar(50)  NOT NULL                   COMMENT 'AGV编号',
  `agv_type`        varchar(30)  DEFAULT NULL               COMMENT 'AGV车型',
  `from_location`   varchar(50)   NOT NULL                   COMMENT '起始点',
  `to_location`     varchar(50)   NOT NULL                   COMMENT '目标点',
  `container_code`  varchar(100) DEFAULT NULL               COMMENT '容器编码',
  `status`          varchar(20)  DEFAULT 'WAITING'            COMMENT '状态：WAITING等待/DISPATCHED已下发/RUNNING执行中/COMPLETED完成/EXCEPTION异常',
  `step_no`         tinyint      DEFAULT 1                   COMMENT '步骤序号',
  `start_time`      datetime     DEFAULT NULL               COMMENT '开始时间',
  `arrive_time`     datetime     DEFAULT NULL               COMMENT '到达时间',
  `complete_time`   datetime     DEFAULT NULL               COMMENT '完成时间',
  `operator`        varchar(64)  DEFAULT NULL               COMMENT '操作员',
  `remark`          varchar(500) DEFAULT NULL               COMMENT '备注',
  `deleted`         bit(1)       DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`       bigint       DEFAULT NULL                COMMENT '租户编号',
  `create_time`     datetime     DEFAULT NULL                COMMENT '创建时间',
  `update_time`     datetime     DEFAULT NULL                COMMENT '更新时间',
  `creator`         varchar(64)  DEFAULT ''                   COMMENT '创建者',
  `updater`         varchar(64)  DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_task_id` (`task_id`),
  KEY `idx_agv_code` (`agv_code`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='AGV作业表';
```

#### 10.1.15 wms_label_template - 标签模板表

> 定义各类标签的打印模板（格式、字段布局、条码规则）。

```sql
CREATE TABLE `wms_label_template` (
  `id`              bigint        NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `template_code`  varchar(50)   NOT NULL                   COMMENT '模板编码',
  `template_name`  varchar(200)  NOT NULL                   COMMENT '模板名称',
  `label_type`     varchar(30)   NOT NULL                   COMMENT '标签类型：ITEM货品标签/CONTAINER容器标签/LOCATION库位标签/PALLET托盘标签/SHIPPING发货标签',
  `paper_width`     decimal(10,2) DEFAULT NULL               COMMENT '纸张宽度（mm）',
  `paper_height`    decimal(10,2) DEFAULT NULL               COMMENT '纸张高度（mm）',
  `template_content` text        DEFAULT NULL               COMMENT '模板内容（JSON格式，定义字段位置和样式）',
  `barcode_type`   varchar(30)  DEFAULT 'CODE128'          COMMENT '条码类型：CODE128/QR_CODE/EAN13',
  `barcode_field`   varchar(50)  DEFAULT NULL               COMMENT '条码绑定字段',
  `printer`         varchar(100) DEFAULT NULL               COMMENT '默认打印机',
  `status`          tinyint      DEFAULT 1                   COMMENT '状态：0禁用/1启用',
  `remark`          varchar(500) DEFAULT NULL               COMMENT '备注',
  `deleted`         bit(1)       DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`       bigint       DEFAULT NULL                COMMENT '租户编号',
  `create_time`     datetime     DEFAULT NULL                COMMENT '创建时间',
  `update_time`     datetime     DEFAULT NULL                COMMENT '更新时间',
  `creator`         varchar(64)  DEFAULT ''                   COMMENT '创建者',
  `updater`         varchar(64)  DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_template_code` (`template_code`, `tenant_id`, `deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签模板表';
```

#### 10.1.16 wms_label_print_record - 标签打印记录表

> 记录标签打印的历史流水，支持追溯和补打。

```sql
CREATE TABLE `wms_label_print_record` (
  `id`              bigint        NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `record_no`       varchar(50)   NOT NULL                   COMMENT '打印记录编号',
  `template_id`      bigint        NOT NULL                   COMMENT '标签模板ID',
  `template_code`    varchar(50)   NOT NULL                   COMMENT '模板编码',
  `label_type`       varchar(30)  NOT NULL                   COMMENT '标签类型',
  `target_code`     varchar(100)  NOT NULL                   COMMENT '打印对象编码（如货品编码/容器编码）',
  `target_name`     varchar(200)  DEFAULT NULL               COMMENT '打印对象名称',
  `barcode_value`    varchar(200)  DEFAULT NULL               COMMENT '条码值',
  `print_count`      int          DEFAULT 1                   COMMENT '打印张数',
  `printer`          varchar(100) DEFAULT NULL               COMMENT '打印机名称',
  `print_by`         varchar(64)  NOT NULL                   COMMENT '打印人',
  `print_time`      datetime     NOT NULL                   COMMENT '打印时间',
  `source_type`      varchar(30)  DEFAULT NULL               COMMENT '来源业务类型',
  `source_no`        varchar(50)  DEFAULT NULL               COMMENT '来源单号',
  `remark`          varchar(500) DEFAULT NULL               COMMENT '备注',
  `deleted`         bit(1)       DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`       bigint       DEFAULT NULL                COMMENT '租户编号',
  `create_time`     datetime     DEFAULT NULL                COMMENT '创建时间',
  `update_time`     datetime     DEFAULT NULL                COMMENT '更新时间',
  `creator`         varchar(64)  DEFAULT ''                   COMMENT '创建者',
  `updater`         varchar(64)  DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_target_code` (`target_code`),
  KEY `idx_print_time` (`print_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签打印记录表';
```

#### 10.1.17 wms_businesstype - 业务类型表

> 维护WMS支持的各类业务类型（采购入库、销售出库等），驱动不同仓库的业务隔离。

```sql
CREATE TABLE `wms_businesstype` (
  `id`              bigint        NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `business_type`   varchar(30)  NOT NULL                   COMMENT '业务类型编码',
  `business_name`   varchar(200) NOT NULL                   COMMENT '业务类型名称',
  `direction`       varchar(10)  NOT NULL                   COMMENT '方向：IN入库/OUT出库/TRANSFER移库',
  `category`        varchar(30)  DEFAULT NULL               COMMENT '类别：PROCUREMENT采购/PRODUCTION生产/SALE销售/MAINTENANCE维保/OTHER其他',
  `in_out_flag`     tinyint      DEFAULT NULL               COMMENT '出入库标识：0出库/1入库',
  `workflow_enabled` bit(1)      DEFAULT b'0'                COMMENT '是否启用审批流',
  `auto_create_job` bit(1)       DEFAULT b'1'                COMMENT '是否自动生成作业',
  `strategy_code`   varchar(50)  DEFAULT NULL               COMMENT '默认策略编码',
  `status`          tinyint      DEFAULT 1                   COMMENT '状态：0禁用/1启用',
  `sort`            int          DEFAULT 0                   COMMENT '排序',
  `remark`          varchar(500) DEFAULT NULL               COMMENT '备注',
  `deleted`         bit(1)       DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`       bigint       DEFAULT NULL                COMMENT '租户编号',
  `create_time`     datetime     DEFAULT NULL                COMMENT '创建时间',
  `update_time`     datetime     DEFAULT NULL                COMMENT '更新时间',
  `creator`         varchar(64)  DEFAULT ''                   COMMENT '创建者',
  `updater`         varchar(64)  DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_business_type` (`business_type`, `tenant_id`, `deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='业务类型表';
```

#### 10.1.18 wms_strategy - 策略配置表

> 配置上架策略、拣货策略、库位分配规则等业务规则引擎。

```sql
CREATE TABLE `wms_strategy` (
  `id`              bigint        NOT NULL  AUTO_INCREMENT  COMMENT '主键',
  `strategy_code`  varchar(50)   NOT NULL                   COMMENT '策略编码',
  `strategy_name`  varchar(200)  NOT NULL                   COMMENT '策略名称',
  `strategy_type`  varchar(30)  NOT NULL                   COMMENT '策略类型：PUTAWAY上架策略/PICK拣货策略/LOCATION_ASSIGN库位分配策略/CONTAINER_BIND容器绑定策略',
  `business_type`  varchar(30)  DEFAULT NULL               COMMENT '适用业务类型',
  `warehouse_code`  varchar(50)  DEFAULT NULL               COMMENT '适用仓库（为空则全局）',
  `item_category`  varchar(50)  DEFAULT NULL               COMMENT '适用货品分类',
  `rule_config`     text         DEFAULT NULL               COMMENT '规则配置（JSON格式）',
  `priority`        int          DEFAULT 100                COMMENT '优先级（数字越小优先级越高）',
  `enabled`         bit(1)       DEFAULT b'1'                COMMENT '是否启用',
  `effective_start` datetime     DEFAULT NULL               COMMENT '生效开始时间',
  `effective_end`   datetime     DEFAULT NULL               COMMENT '生效结束时间',
  `status`          tinyint      DEFAULT 1                   COMMENT '状态：0禁用/1启用',
  `remark`          varchar(500) DEFAULT NULL               COMMENT '备注',
  `deleted`         bit(1)       DEFAULT b'0'                COMMENT '是否删除',
  `tenant_id`       bigint       DEFAULT NULL                COMMENT '租户编号',
  `create_time`     datetime     DEFAULT NULL                COMMENT '创建时间',
  `update_time`     datetime     DEFAULT NULL                COMMENT '更新时间',
  `creator`         varchar(64)  DEFAULT ''                   COMMENT '创建者',
  `updater`         varchar(64)  DEFAULT ''                   COMMENT '更新者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_strategy_code` (`strategy_code`, `tenant_id`, `deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='策略配置表';
```

---

### 10.2 补充API设计

#### 10.2.1 /wms/area/* - 库区管理API

> 实现仓库→库区→库位三级结构的CRUD管理，支持树形层级查询。

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/area/create` | createArea | **库区** |
| PUT | `/wms/area/update` | updateArea | **库区** |
| DELETE | `/wms/area/delete` | deleteArea | **库区** |
| GET | `/wms/area/get` | getArea | **编号** |
| GET | `/wms/area/page` | getAreaPage | **库区** |
| GET | `/wms/area/tree` | getAreaTree | **仓库** |
| GET | `/wms/area/listByWarehouse` | listByWarehouse | **仓库** |
| POST | `/wms/area/senior` | getAreaSenior | **库区** |
| GET | `/wms/area/export-excel` | exportAreaExcel | **库区** |
| POST | `/wms/area/import` | importAreaExcel | **库区** |

#### 10.2.2 /wms/item/* - 货品管理API

> 实现物料到货品的映射管理，配置货品的仓库默认参数。

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/item/create` | createItem | **货品** |
| PUT | `/wms/item/update` | updateItem | **货品** |
| DELETE | `/wms/item/delete` | deleteItem | **货品** |
| GET | `/wms/item/get` | getItem | **编号** |
| GET | `/wms/item/page` | getItemPage | **货品** |
| GET | `/wms/item/list` | getItemList | **货品** |
| GET | `/wms/item/listByMaterial` | listByMaterial | **物料** |
| POST | `/wms/item/senior` | getItemSenior | **货品** |
| GET | `/wms/item/export-excel` | exportItemExcel | **货品** |
| POST | `/wms/item/import` | importItemExcel | **货品** |

#### 10.2.3 /wms/container/* - 容器管理API

> 实现容器的CRUD及绑定/解绑/维修操作。

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/container/create` | createContainer | **容器** |
| PUT | `/wms/container/update` | updateContainer | **容器** |
| DELETE | `/wms/container/delete` | deleteContainer | **容器** |
| GET | `/wms/container/get` | getContainer | **编号** |
| GET | `/wms/container/page` | getContainerPage | **容器** |
| GET | `/wms/container/list` | getContainerList | **容器** |
| PUT | `/wms/container/bind` | bindContainer | **容器** |
| PUT | `/wms/container/unbind` | unbindContainer | **容器** |
| PUT | `/wms/container/maintain` | maintainContainer | **容器** |
| PUT | `/wms/container/block` | blockContainer | **容器** |

#### 10.2.4 /wms/pick/* - 拣货作业API

> 实现销售发货拣货、生产领料拣货的作业管理。

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/pick/create` | createPickJob | **拣货作业** |
| PUT | `/wms/pick/assign` | assignPickJob | **拣货作业** |
| PUT | `/wms/pick/execute` | executePickJob | **拣货作业** |
| PUT | `/wms/pick/complete` | completePickJob | **拣货作业** |
| PUT | `/wms/pick/cancel` | cancelPickJob | **拣货作业** |
| GET | `/wms/pick/get` | getPickJob | **编号** |
| GET | `/wms/pick/page` | getPickJobPage | **拣货作业** |
| GET | `/wms/pick/pageBySource` | getPickJobPageBySource | **来源单号** |
| POST | `/wms/pick/senior` | getPickJobSenior | **拣货作业** |
| GET | `/wms/pick/export-excel` | exportPickJobExcel | **拣货作业** |

#### 10.2.5 /wms/putaway/* - 上架作业API

> 实现验货后货品入库上架的执行管理。

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/putaway/create` | createPutawayJob | **上架作业** |
| PUT | `/wms/putaway/assign` | assignPutawayJob | **上架作业** |
| PUT | `/wms/putaway/execute` | executePutawayJob | **上架作业** |
| PUT | `/wms/putaway/complete` | completePutawayJob | **上架作业** |
| PUT | `/wms/putaway/cancel` | cancelPutawayJob | **上架作业** |
| GET | `/wms/putaway/get` | getPutawayJob | **编号** |
| GET | `/wms/putaway/page` | getPutawayJobPage | **上架作业** |
| GET | `/wms/putaway/pageBySource` | getPutawayJobPageBySource | **来源单号** |
| POST | `/wms/putaway/senior` | getPutawayJobSenior | **上架作业** |
| GET | `/wms/putaway/export-excel` | exportPutawayJobExcel | **上架作业** |

#### 10.2.6 /wms/count/* - 盘点管理API

> 实现盘点计划制定、任务下发、执行记录全流程管理。

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/count/plan/create` | createCountPlan | **盘点计划** |
| PUT | `/wms/count/plan/submit` | submitCountPlan | **盘点计划** |
| PUT | `/wms/count/plan/publish` | publishCountPlan | **盘点计划** |
| PUT | `/wms/count/plan/resetting` | resettingCountPlan | **盘点计划** |
| GET | `/wms/count/plan/get` | getCountPlan | **编号** |
| GET | `/wms/count/plan/page` | getCountPlanPage | **盘点计划** |
| PUT | `/wms/count/job/execute` | executeCountJob | **盘点任务** |
| PUT | `/wms/count/job/submit` | submitCountJob | **盘点任务** |
| GET | `/wms/count/job/page` | getCountJobPage | **盘点任务** |
| GET | `/wms/count/record/page` | getCountRecordPage | **盘点记录** |

#### 10.2.7 /wms/agv/* - AGV调度API

> 实现AGV任务下发、状态跟踪和作业记录。

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/agv/task/dispatch` | dispatchAgvTask | **AGV任务** |
| PUT | `/wms/agv/task/cancel` | cancelAgvTask | **AGV任务** |
| GET | `/wms/agv/task/get` | getAgvTask | **编号** |
| GET | `/wms/agv/task/page` | getAgvTaskPage | **AGV任务** |
| GET | `/wms/agv/task/status` | getAgvTaskStatus | **任务编号** |
| POST | `/wms/agv/task/callback` | agvTaskCallback | **AGV回调** |
| GET | `/wms/agv/job/page` | getAgvJobPage | **AGV作业** |
| GET | `/wms/agv/job/export` | exportAgvJobExcel | **AGV作业** |
| GET | `/wms/agv/device/list` | listAgvDevice | **AGV设备** |
| GET | `/wms/agv/device/status` | getAgvDeviceStatus | **AGV设备** |

#### 10.2.8 /wms/label/* - 标签打印API

> 实现标签模板管理和打印执行。

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/label/template/create` | createLabelTemplate | **标签模板** |
| PUT | `/wms/label/template/update` | updateLabelTemplate | **标签模板** |
| DELETE | `/wms/label/template/delete` | deleteLabelTemplate | **标签模板** |
| GET | `/wms/label/template/get` | getLabelTemplate | **编号** |
| GET | `/wms/label/template/list` | getLabelTemplateList | **标签类型** |
| GET | `/wms/label/template/page` | getLabelTemplatePage | **标签模板** |
| POST | `/wms/label/print` | printLabel | **打印** |
| POST | `/wms/label/batchPrint` | batchPrintLabel | **批量打印** |
| GET | `/wms/label/record/page` | getLabelPrintRecordPage | **打印记录** |
| GET | `/wms/label/record/reprint` | reprintLabel | **打印记录** |

#### 10.2.9 /wms/strategy/* - 策略配置API

> 实现上架策略、拣货策略、库位分配规则的配置管理。

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/strategy/create` | createStrategy | **策略配置** |
| PUT | `/wms/strategy/update` | updateStrategy | **策略配置** |
| DELETE | `/wms/strategy/delete` | deleteStrategy | **策略配置** |
| GET | `/wms/strategy/get` | getStrategy | **编号** |
| GET | `/wms/strategy/page` | getStrategyPage | **策略配置** |
| GET | `/wms/strategy/listByType` | listByType | **策略类型** |
| POST | `/wms/strategy/senior` | getStrategySenior | **策略配置** |
| PUT | `/wms/strategy/enable` | enableStrategy | **策略配置** |
| PUT | `/wms/strategy/disable` | disableStrategy | **策略配置** |
| GET | `/wms/strategy/export-excel` | exportStrategyExcel | **策略配置** |

---

### 10.3 补充业务功能说明

#### 10.3.1 库区管理（仓库-库区-库位三级结构）

库区管理是WMS的基础档案之一，采用仓库→库区→库位三级层级结构：

- **仓库层**：顶级实体，代表一个物理仓库或虚拟仓库，配置业务类型和负责人。
- **库区层**：仓库下的功能分区，如收货区、存储区、拣货区、发货区、质检区、退货区等，支持多级子库区。
- **库位层**：库区下的具体存放位置，支持定义库位类型（立体库位、地面库位）、容量限制（最大货位数、最长边限制）。

**核心功能**：
- 树形层级维护，支持拖拽调整层级关系
- 库位容量管理（最大存放数量、最长边/宽/高/承重）
- 库区类型过滤（收货区只允许收货作业，拣货区只允许出库）
- 层级级联状态（仓库停用则下属库区自动停用）

#### 10.3.2 货品管理（物料与货品对应关系）

货品档案是物料在WMS中的业务化映射，建立MDM主数据与仓储业务的关联：

- **映射关系**：一个物料编码（MDM）对应一条货品档案（WMS），可配置默认仓库/库区/库位。
- **库存策略**：配置批次管理、保质期预警、最小/最大库存警戒线。
- **包装配置**：维护货品的销售单位、库存单位、入库单位的换算关系。
- **仓库适配**：配置货品在不同仓库的存储参数（如冷链要求、危险品级别）。

**核心功能**：
- 物料编码批量映射导入
- 货品-仓库默认配置（不同仓库可设置不同参数）
- 库存预警通知触发配置
- 货品分类ABC分析辅助库位优化

#### 10.3.3 容器管理（容器绑定/解绑/维修）

容器（托盘、周转箱）是WMS中货品承载和搬运的基本单元：

- **容器主档**：维护容器类型（托盘/周转箱）、规格（长*宽*高）、最大承重/容量。
- **容器绑定**：将容器与货品、批次绑定，确保货品与容器的一一对应关系，用于追溯。
- **容器解绑**：货品出库后，释放容器使其恢复空闲状态。
- **容器维修**：记录容器的维修历史和当前状态（空闲/使用中/维修中/冻结）。
- **容器明细**：支持一容器多货品，记录容器内各货品的明细数量和批次信息。

**核心功能**：
- 容器与AGV搬运任务的联动（AGV自动搬运绑定货品的容器）
- 容器状态全程跟踪
- 容器使用率统计分析
- 容器维修记录和到期提醒

#### 10.3.4 AGV调度（AGV任务下发与跟踪）

AGV调度模块负责将仓库搬运任务自动分配给AGV执行：

- **任务类型**：上架搬运（收货暂存→库位）、出库搬运（库位→发货暂存）、库间移库、容器归位、AGV充电。
- **任务下发**：根据策略规则（最近原则、负载均衡、路径最优）自动分配AGV设备。
- **状态跟踪**：实时接收AGV设备的状态回调，更新任务进度（等待→已分配→执行中→完成/异常）。
- **异常处理**：AGV执行异常时自动重试，超限后人工介入处理。
- **AGV作业**：一条AGV任务可拆分为多个作业步骤（如：从A点到B点→取货→从B点到C点→放货）。

**核心功能**：
- 多AGV统一调度和路径规划
- 任务优先级管理
- AGV设备状态实时监控
- 搬运效率统计（平均搬运时长、AGV利用率）

#### 10.3.5 标签打印（标签模板设计、条码规则）

标签打印模块提供灵活的标签设计和批量打印能力：

- **标签类型**：货品标签（物料名称/规格/批次）、容器标签（容器编码/绑定货品）、库位标签（库位编码/容量）、托盘标签、发货标签（客户信息/订单信息）。
- **模板设计**：通过JSON格式定义标签布局（文字/条码/二维码/图片的位置、字体、尺寸），支持预览。
- **条码规则**：支持CODE128、QR_CODE、EAN13等条码格式，绑定对应字段自动生成条码值。
- **打印追溯**：记录每次打印的时间、人、打印机，支持按来源单号追溯打印记录。
- **批量打印**：支持从入库单/出库单等业务单据批量触发标签打印。

**核心功能**：
- 标签模板的增删改查和版本管理
- 打印参数配置（纸张规格、打印机选择）
- 按批次/单据批量打印
- 打印记录追溯和补打

#### 10.3.6 拣货下架（销售发货/领料拣货）

拣货下架是出库的核心环节，将库存中的货品按订单要求从库位移至发货暂存区：

- **销售发货拣货**：根据销售订单生成拣货任务，拣货员按最优路径从库位取货。
- **生产领料拣货**：根据MES工单生成领料拣货任务，从原材料库区拣货至产线工位。
- **拣货策略**：支持单人拣货、接力拣货（先拣后分）、分区拣货等多种模式。
- **作业流程**：任务创建→任务分配→执行拣货→完成复核→货品移交发货暂存区。
- **差异处理**：拣货数量与计划不符时，记录差异原因（货品损坏/库位错误/库存不足）。

**核心功能**：
- 拣货路径优化（按库位顺序规划最短路径）
- PDA扫码复核（扫描货品条码和库位条码双重校验）
- 拣货进度实时跟踪
- 拣货效率统计（单人拣货件数/时长）

#### 10.3.7 上架策略（入库上架规则配置）

上架策略决定验货后的货品应放置到哪个库位：

- **策略类型**：同品集中（相同货品/批次集中存放）、随机就近（快速上架）、先入先出（按批次日期）、温湿度分区（特殊货品专用库区）。
- **策略规则**：按业务类型/货品分类/仓库分别配置策略优先级，命中规则后自动推荐库位。
- **库位推荐**：根据策略规则和库位当前容量，从可用库位中自动推荐最优上架位置。
- **上架执行**：上架员按推荐库位执行上架，支持扫码确认和人工调整库位。

**核心功能**：
- 策略规则可视化配置（JSON编辑器）
- 策略生效时间区间配置（支持临时策略）
- 上架推荐库位的二次确认/调整
- 上架效率统计（平均上架时长/库位利用率）

#### 10.3.8 盘点管理（盘点计划/任务/执行）

盘点管理通过计划→任务→执行三层结构实现定期或不定期的库存核查：

- **盘点计划**：制定盘点范围（按仓库/库区/货品）、盘点类型（全盘/抽盘/循环盘点）、周期和时间安排。
- **审批发布**：盘点计划提交后经审批通过方可发布，发布后自动生成盘点任务。
- **任务分配**：盘点任务精确到库位级（精盘）或货品级（抽盘），分配给指定盘点人员。
- **执行记录**：盘点员通过PDA扫描库位和货品，输入实盘数量，系统自动计算差异。
- **差异审批**：盘点差异经审批后自动生成调账单，更新库存台账。
- **循环盘点**：按ABC分类自动生成抽盘任务，实现库存的持续性核查。

**核心功能**：
- 计划制定→审批→发布→执行→审批→调账完整流程管控
- PDA扫码盘点（库位+货品双重扫码确认）
- 差异率阈值配置（超阈值需复盘或审批）
- 盘点报表（盘点准确率、差异汇总、损益金额）

---

## 11. 文档状态

### 11.1 文档版本

| 版本 | 日期 | 作者 | 说明 |
|------|------|------|------|
| V1.0 | 2026-04-17 | CI | 初始版本，涵盖基础档案、入出库、库存、质检、标签、AGV核心功能 |
| V1.1 | 2026-04-17 | CI | 补充DDL表结构（18张）、API设计（9大模块）、业务功能说明（8项） |

### 11.2 章节完成状态

| 章节 | 内容 | 状态 |
|------|------|------|
| 1. 模块概述 | 路径、子模块结构 | 已完成 |
| 2. 模块职责 | 9大功能域职责 | 已完成 |
| 3. 核心类/接口 | 目录结构、控制器、服务、DO设计 | 已完成 |
| 4. 数据结构 | 通用设计模式、核心表结构、VO对象 | 已完成 |
| 5. 数据流向 | 入库/出库/盘点流程、数据流图 | 已完成 |
| 6. 关键技术实现 | 高级搜索、幂等、导入导出、权限、业务类型驱动 | 已完成 |
| 7. 集成关系 | SCP、MES、Report、System模块集成 | 已完成 |
| 8. 错误码 | WMS模块错误码定义 | 已完成 |
| 9. 模块特点 | 规模、业务复杂度、可配置性 | 已完成 |
| **10. 补充设计** | **DDL表结构（18张）、API设计（9模块）、业务功能（8项）** | **已完成** |
| **11. 文档状态** | **版本记录、章节完成状态** | **已完成** |

### 11.3 待补充内容

> 以下内容暂未纳入本文档，后续版本迭代补充。

| 序号 | 待补充内容 | 优先级 | 备注 |
|------|-----------|--------|------|
| 1 | 移库作业完整API与DDD设计 | P1 | 库内移库流程与AGV联动 |
| 2 | 调账申请/记录完整API | P1 | 库存调整审批流 |
| 3 | 序列号（SN）管理设计 | P2 | 唯一序列号追溯 |
| 4 | 波次分派（WAVE）设计 | P2 | 订单波次合并优化 |
| 5 | 报表与数据看板 | P2 | 库存周转率、库容分析 |

## 完整API清单

> 本清单由脚本自动从 359 个Controller文件中提取，共 3950 个API端点。

### 账期日历 (accountcalendar)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/accountcalendar/create` | createAccountcalendar | **账期日历** |
| PUT | `/wms/accountcalendar/update` | updateAccountcalendar | **账期日历** |
| DELETE | `/wms/accountcalendar/delete` | deleteAccountcalendar | **账期日历** |
| GET | `/wms/accountcalendar/get` | getAccountcalendar | **编号** |
| GET | `/wms/accountcalendar/page` | getAccountcalendarPage | **编号** |
| POST | `/wms/accountcalendar/senior` | getAccountcalendarSenior | **账期日历** |
| GET | `/wms/accountcalendar/export-excel` | exportAccountcalendarExcel | **账期日历** |
| POST | `/wms/accountcalendar/export-excel-senior` | exportAccountcalendarExcel | **账期日历** |
| GET | `/wms/accountcalendar/get-import-template` | importTemplate | **账期日历** |
| POST | `/wms/accountcalendar/import` | importExcel | **账期日历** |

### AGV库位关系 (agvlocationrelation)  (11个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/agv-locationrelation/create` | createAgvLocationrelation | **AGV库位转换** |
| PUT | `/wms/agv-locationrelation/update` | updateAgvLocationrelation | **AGV库位转换** |
| DELETE | `/wms/agv-locationrelation/delete` | deleteAgvLocationrelation | **AGV库位转换** |
| GET | `/wms/agv-locationrelation/get` | getAgvLocationrelation | **编号** |
| GET | `/wms/agv-locationrelation/list` | getAgvLocationrelationList | **编号** |
| GET | `/wms/agv-locationrelation/page` | getAgvLocationrelationPage | **编号列表** |
| POST | `/wms/agv-locationrelation/senior` | getAgvLocationrelationSenior | **AGV库位转换** |
| GET | `/wms/agv-locationrelation/export-excel` | exportAgvLocationrelationExcel | **AGV库位转换** |
| POST | `/wms/agv-locationrelation/export-excel-senior` | exportAgvLocationrelationExcel | **AGV库位转换** |
| GET | `/wms/agv-locationrelation/get-import-template` | importTemplate | **AGV库位转换** |
| POST | `/wms/agv-locationrelation/import` | importExcel | **AGV库位转换** |

### 综合作业 (allJob)  (1个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/all-job/list` | getAllJobCountList | **PDA根据用户权限操作所有任务模块** |

### 库区 (areabasic)  (14个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/areabasic/create` | createAreabasic | **库区** |
| PUT | `/wms/areabasic/update` | updateAreabasic | **库区** |
| DELETE | `/wms/areabasic/delete` | deleteAreabasic | **库区** |
| GET | `/wms/areabasic/get` | getAreabasic | **编号** |
| GET | `/wms/areabasic/getArea` | selectAreabasicDOByCode | **编号** |
| GET | `/wms/areabasic/list` | getAreabasicList | **库位代码** |
| GET | `/wms/areabasic/page` | getAreabasicPage | **库区** |
| GET | `/wms/areabasic/export-excel` | exportAreabasicExcel | **库区** |
| POST | `/wms/areabasic/export-excel-senior` | exportAreabasicExcel | **库区** |
| POST | `/wms/areabasic/senior` | getAreabasicSenior | **库区** |
| GET | `/wms/areabasic/get-import-template` | importTemplate | **库区** |
| POST | `/wms/areabasic/import` | importExcel | **库区** |
| GET | `/wms/areabasic/listByCodes` | getAreabasicByCodes | **Excel 文件** |
| GET | `/wms/areabasic/listAreabasicByCode` | listAreabasicByCode | **库区** |

### QAD反冲明细 (backflushdetailbqad)  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/backflush-detailb-qad/create` | createBackflushDetailbQad | **制品收货记录子表QAD返回数据** |
| PUT | `/wms/backflush-detailb-qad/update` | updateBackflushDetailbQad | **制品收货记录子表QAD返回数据** |
| DELETE | `/wms/backflush-detailb-qad/delete` | deleteBackflushDetailbQad | **制品收货记录子表QAD返回数据** |
| GET | `/wms/backflush-detailb-qad/get` | getBackflushDetailbQad | **编号** |
| GET | `/wms/backflush-detailb-qad/page` | getBackflushDetailbQadPage | **编号** |
| POST | `/wms/backflush-detailb-qad/senior` | getBackflushDetailbQadSenior | **制品收货记录子表QAD返回数据** |
| GET | `/wms/backflush-detailb-qad/export-excel` | exportBackflushDetailbQadExcel | **制品收货记录子表QAD返回数据** |
| GET | `/wms/backflush-detailb-qad/get-import-template` | importTemplate | **制品收货记录子表QAD返回数据** |
| POST | `/wms/backflush-detailb-qad/import` | importExcel | **制品收货记录子表QAD返回数据** |

### 库存台账 (balance)  (44个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| PUT | `/wms/balance/update` | updateBalance | **库存余额** |
| GET | `/wms/balance/page` | getBalancePage | **库存余额** |
| GET | `/wms/balance/pageAll` | getBalancePageAll | **库存余额** |
| POST | `/wms/balance/seniorAll` | getBalanceSeniorAll | **库存余额** |
| GET | `/wms/balance/pageContainer` | getBalancePageContainer | **库存余额** |
| POST | `/wms/balance/seniorContainer` | getBalanceSeniorContainer | **库存余额** |
| GET | `/wms/balance/pageContainerToScrap` | getBalancePageContainerToScrap | **库存余额** |
| POST | `/wms/balance/seniorContainerToScrap` | getBalanceSeniorContainerToScrap | **库存余额** |
| GET | `/wms/balance/pageReturn` | getBalancePageReturn | **库存余额** |
| POST | `/wms/balance/seniorReturn` | getBalanceSeniorReturn | **库存余额** |
| GET | `/wms/balance/pagePutaway` | getBalancePagePutaway | **库存余额** |
| POST | `/wms/balance/senior` | getBalanceSenior | **库存余额** |
| GET | `/wms/balance/export-excel` | exportBalanceExcel | **库存余额** |
| POST | `/wms/balance/export-excel-senior` | exportBalanceExcel | **库存余额** |
| GET | `/wms/balance/pageItems` | getBalanceAndItemPage | **库存余额** |
| POST | `/wms/balance/seniorItems` | getBalanceItemSenior | **库存余额** |
| GET | `/wms/balance/pageSpareItem` | getBalanceAndItemPageSpare | **库存余额** |
| POST | `/wms/balance/seniorSpareItems` | getBalanceItemSeniorSpare | **库存余额** |
| GET | `/wms/balance/pageBusinessType` | getBalancePageBusinessType | **库存余额** |
| POST | `/wms/balance/seniorBusinessType` | getBalanceSeniorBusinessType | **库存余额** |
| GET | `/wms/balance/pageBusinessTypeByItemType` | getBalancePageBusinessTypeByItemType | **库存余额** |
| POST | `/wms/balance/seniorBusinessTypeByItemType` | getBalanceSeniorBusinessTypeItemType | **库存余额** |
| GET | `/wms/balance/pageLocationCodeToBalance` | selectLocationTypeToBalance | **库存余额** |
| POST | `/wms/balance/pageLocationCodeToBalanceSenior` | getLocationTypeToBalanceSenior | **业务类型** |
| GET | `/wms/balance/pageConfigToBalance` | selectConfigToBalance | **库存余额** |
| POST | `/wms/balance/pageConfigToBalanceSenior` | selectConfigToBalanceSenior | **配置** |
| GET | `/wms/balance/pageBusinessTypeToBalance` | balancePageBusinessType | **库存余额** |
| POST | `/wms/balance/pageBusinessTypeToBalanceSenior` | balanceSeniorBusinessType | **库存余额** |
| GET | `/wms/balance/pageBusinessCategoryToBalance` | balancePageBusinessCategory | **库存余额** |
| POST | `/wms/balance/pageBusinessCategoryToBalanceSenior` | balanceSeniorBusinessCategory | **库存余额** |
| GET | `/wms/balance/summary` | queryGroupBySummary | **库存余额** |
| POST | `/wms/balance/batchPrintLabel` | batchPrintLabel | **库存余额** |
| GET | `/wms/balance/getSumByConditions` | getSumByConditions | **库存余额** |
| GET | `/wms/balance/listByCodes` | getBalanceListByCodes | **库存余额** |
| POST | `/wms/balance/getBalanceListByPackage` | getBalanceListByPackage | **物料编码** |
| GET | `/wms/balance/pageBOM` | getBalancePageBOM | **库存余额** |
| POST | `/wms/balance/seniorBOM` | getBalanceSeniorBOM | **库存余额** |
| GET | `/wms/balance/queryBalancePurchaseReceiptReturn` | queryBalancePurchaseReceiptReturn | **库存余额** |
| GET | `/wms/balance/summaryByBusinessType` | summaryByBusinessType | **库存余额** |
| GET | `/wms/balance/totalBalanceTree` | getTransactionBalances | **库存余额** |
| GET | `/wms/balance/exportTotalBalanceTree` | exportTotalBalanceExcel | **库存余额** |
| POST | `/wms/balance/balanceSeniorByLocation` | getBalanceSeniorByLocation | **库存余额** |
| GET | `/wms/balance/balancePageByLocation` | getBalancePageByLocation | **库存余额** |
| GET | `/wms/balance/exportTotalBalanceTree` | exportTotalBalanceTree | **库存余额** |

### 库存变化历史 (balancechangehistory)  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/balance-change-history/create` | createBalanceChangeHistory | **库存余额变更记录** |
| PUT | `/wms/balance-change-history/update` | updateBalanceChangeHistory | **库存余额变更记录** |
| DELETE | `/wms/balance-change-history/delete` | deleteBalanceChangeHistory | **库存余额变更记录** |
| GET | `/wms/balance-change-history/get` | getBalanceChangeHistory | **编号** |
| GET | `/wms/balance-change-history/page` | getBalanceChangeHistoryPage | **编号** |
| POST | `/wms/balance-change-history/senior` | getBalanceChangeHistorySenior | **库存余额变更记录** |
| GET | `/wms/balance-change-history/export-excel` | exportBalanceChangeHistoryExcel | **库存余额变更记录** |
| GET | `/wms/balance-change-history/get-import-template` | importTemplate | **库存余额变更记录** |
| POST | `/wms/balance-change-history/import` | importExcel | **库存余额变更记录** |

### 条码 (barcode)  (11个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/label/barcode/create` | createBarcode | **标签规则设置** |
| PUT | `/label/barcode/update` | updateBarcode | **标签规则设置** |
| DELETE | `/label/barcode/delete` | deleteBarcode | **标签规则设置** |
| GET | `/label/barcode/get` | getBarcode | **编号** |
| GET | `/label/barcode/list` | getBarcodeList | **编号** |
| GET | `/label/barcode/page` | getBarcodePage | **编号列表** |
| POST | `/label/barcode/senior` | getBarcodeSenior | **标签规则设置** |
| GET | `/label/barcode/export-excel` | exportBarcodeExcel | **标签规则设置** |
| POST | `/label/barcode/export-excel-senior` | exportBarcodeExcel | **标签规则设置** |
| GET | `/label/barcode/get-import-template` | importTemplate | **标签规则设置** |
| POST | `/label/barcode/import` | importExcel | **标签规则设置** |

### 看板 (board)  (4个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/board/purchasereceipt` | purchasereceipt | **看板** |
| GET | `/wms/board/issue` | issue | **看板** |
| GET | `/wms/board/production` | production | **看板** |
| GET | `/wms/board/deliver` | deliver | **看板** |

### BOM (bom)  (11个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/bom/create` | createBom | **物料清单** |
| PUT | `/wms/bom/update` | updateBom | **物料清单** |
| DELETE | `/wms/bom/delete` | deleteBom | **物料清单** |
| GET | `/wms/bom/get` | getBom | **编号** |
| GET | `/wms/bom/page` | getBomPage | **编号** |
| POST | `/wms/bom/senior` | getBomSenior | **物料清单** |
| GET | `/wms/bom/export-excel` | exportBomExcel | **物料清单** |
| POST | `/wms/bom/export-excel-senior` | exportBomExcel | **物料清单** |
| GET | `/wms/bom/get-import-template` | importTemplate | **物料清单** |
| POST | `/wms/bom/import` | importExcel | **物料清单** |
| GET | `/wms/bom/queryBom` | getPurchasereceiptJobyId | **Excel 文件** |

### 业务类型 (businesstype)  (11个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/businesstype/create` | createBusinesstype | **业务类型** |
| PUT | `/wms/businesstype/update` | updateBusinesstype | **业务类型** |
| DELETE | `/wms/businesstype/delete` | deleteBusinesstype | **业务类型** |
| GET | `/wms/businesstype/get` | getBusinesstype | **编号** |
| GET | `/wms/businesstype/list` | getBusinesstypeList | **编号** |
| GET | `/wms/businesstype/page` | getBusinesstypePage | **编号列表** |
| POST | `/wms/businesstype/senior` | getBusinesstypeSenior | **业务类型** |
| GET | `/wms/businesstype/export-excel` | exportBusinesstypeExcel | **业务类型** |
| GET | `/wms/businesstype/get-import-template` | importTemplate | **业务类型** |
| POST | `/wms/businesstype/import` | importExcel | **业务类型** |
| GET | `/wms/businesstype/getBusinesstypeCode` | getBusinesstypeCode | **Excel 文件** |

### 叫料 (callmaterials)  (11个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/callmaterials/create` | createCallmaterials | **叫料标签** |
| PUT | `/wms/callmaterials/update` | updateCallmaterials | **叫料标签** |
| DELETE | `/wms/callmaterials/delete` | deleteCallmaterials | **叫料标签** |
| GET | `/wms/callmaterials/get` | getCallmaterials | **编号** |
| GET | `/wms/callmaterials/list` | getCallmaterialsList | **编号** |
| GET | `/wms/callmaterials/page` | getCallmaterialsPage | **编号列表** |
| POST | `/wms/callmaterials/senior` | getCallmaterialsSenior | **叫料标签** |
| GET | `/wms/callmaterials/export-excel` | exportCallmaterialsExcel | **叫料标签** |
| POST | `/wms/callmaterials/export-excel-senior` | exportCallmaterialsSeniorExcel | **叫料标签** |
| GET | `/wms/callmaterials/get-import-template` | importTemplate | **叫料标签** |
| POST | `/wms/callmaterials/import` | importExcel | **叫料标签** |

### 承运商 (carrier)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/carrier/create` | createCarrier | **承运商** |
| PUT | `/wms/carrier/update` | updateCarrier | **承运商** |
| DELETE | `/wms/carrier/delete` | deleteCarrier | **承运商** |
| GET | `/wms/carrier/get` | getCarrier | **编号** |
| GET | `/wms/carrier/page` | getCarrierPage | **编号** |
| POST | `/wms/carrier/senior` | getCarrierSenior | **承运商** |
| GET | `/wms/carrier/export-excel` | exportCarrierExcel | **承运商** |
| POST | `/wms/carrier/export-excel-senior` | exportCarrierExcel | **承运商** |
| GET | `/wms/carrier/get-import-template` | importTemplate | **承运商** |
| POST | `/wms/carrier/import` | importExcel | **承运商** |

### 条件 (condition)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/condition/create` | createCondition | **条件** |
| PUT | `/wms/condition/update` | updateCondition | **条件** |
| DELETE | `/wms/condition/delete` | deleteCondition | **条件** |
| GET | `/wms/condition/get` | getCondition | **编号** |
| GET | `/wms/condition/list` | getConditionList | **编号** |
| GET | `/wms/condition/page` | getConditionPage | **编号列表** |
| POST | `/wms/condition/senior` | getConditionSenior | **条件** |
| GET | `/wms/condition/export-excel` | exportConditionExcel | **条件** |
| GET | `/wms/condition/get-import-template` | importTemplate | **条件** |
| POST | `/wms/condition/import` | importExcel | **条件** |

### 配置 (configuration)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/configuration/create` | createConfiguration | **配置** |
| PUT | `/wms/configuration/update` | updateConfiguration | **配置** |
| DELETE | `/wms/configuration/delete` | deleteConfiguration | **配置** |
| GET | `/wms/configuration/get` | getConfiguration | **编号** |
| GET | `/wms/configuration/list` | getConfigurationList | **编号** |
| GET | `/wms/configuration/page` | getConfigurationPage | **编号列表** |
| POST | `/wms/configuration/senior` | getConfigurationSenior | **配置** |
| GET | `/wms/configuration/export-excel` | exportConfigurationExcel | **配置** |
| GET | `/wms/configuration/get-import-template` | importTemplate | **配置** |
| POST | `/wms/configuration/import` | importExcel | **配置** |

### 配置设置 (configurationsetting)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/configurationsetting/create` | createConfigurationsetting | **配置设置** |
| PUT | `/wms/configurationsetting/update` | updateConfigurationsetting | **配置设置** |
| DELETE | `/wms/configurationsetting/delete` | deleteConfigurationsetting | **配置设置** |
| GET | `/wms/configurationsetting/get` | getConfigurationsetting | **编号** |
| GET | `/wms/configurationsetting/page` | getConfigurationsettingPage | **编号** |
| POST | `/wms/configurationsetting/senior` | getConfigurationsettingSenior | **配置设置** |
| GET | `/wms/configurationsetting/export-excel` | exportConfigurationsettingExcel | **配置设置** |
| POST | `/wms/configurationsetting/export-excel-senior` | exportConfigurationsettingExcel | **配置设置** |
| GET | `/wms/configurationsetting/get-import-template` | importTemplate | **配置设置** |
| POST | `/wms/configurationsetting/import` | importExcel | **配置设置** |

### 容器 (container)  (21个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/container-detail/create` | createContainerDetail | **器具子** |
| PUT | `/wms/container-detail/update` | updateContainerDetail | **器具子** |
| DELETE | `/wms/container-detail/delete` | deleteContainerDetail | **器具子** |
| GET | `/wms/container-detail/get` | getContainerDetail | **编号** |
| GET | `/wms/container-detail/list` | getContainerDetailList | **编号** |
| GET | `/wms/container-detail/page` | getContainerDetailPage | **编号列表** |
| POST | `/wms/container-detail/senior` | getContainerDetailSenior | **器具子** |
| GET | `/wms/container-detail/export-excel` | exportContainerDetailExcel | **器具子** |
| POST | `/wms/container-main/create` | createContainerMain | **器具主** |
| PUT | `/wms/container-main/update` | updateContainerMain | **器具主** |
| DELETE | `/wms/container-main/delete` | deleteContainerMain | **器具主** |
| GET | `/wms/container-main/get` | getContainerMain | **编号** |
| GET | `/wms/container-main/list` | getContainerMainList | **编号** |
| GET | `/wms/container-main/page` | getContainerMainPage | **编号列表** |
| POST | `/wms/container-main/senior` | getContainerMainSenior | **器具主** |
| GET | `/wms/container-main/export-excel` | exportContainerMainExcel | **器具主** |
| GET | `/wms/container-main/getContainerByNumber` | getContainerByNumber | **器具主** |
| POST | `/wms/container-main/containerBind` | pdaBind | **编号** |
| POST | `/wms/container-main/containerUnBind` | pdaUnBind | **器具主** |
| PUT | `/wms/container-main/repair` | repairSubmitContainerRequestMain | **器具主** |
| PUT | `/wms/container-main/scrap` | scrapSubmitContainerRequestMain | **器具主** |

### 容器绑定 (containerBind)  (21个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/container-bind-record-detail/create` | createContainerBindRecordDetail | **器具绑定记录子** |
| PUT | `/wms/container-bind-record-detail/update` | updateContainerBindRecordDetail | **器具绑定记录子** |
| DELETE | `/wms/container-bind-record-detail/delete` | deleteContainerBindRecordDetail | **器具绑定记录子** |
| GET | `/wms/container-bind-record-detail/get` | getContainerBindRecordDetail | **编号** |
| GET | `/wms/container-bind-record-detail/list` | getContainerBindRecordDetailList | **编号** |
| GET | `/wms/container-bind-record-detail/page` | getContainerBindRecordDetailPage | **编号列表** |
| POST | `/wms/container-bind-record-detail/senior` | getInventorychangeRecordDetailSenior | **器具绑定记录子** |
| GET | `/wms/container-bind-record-detail/export-excel` | exportContainerBindRecordDetailExcel | **器具绑定记录子** |
| GET | `/wms/container-bind-record-detail/get-import-template` | importTemplate | **器具绑定记录子** |
| POST | `/wms/container-bind-record-detail/import` | importExcel | **器具绑定记录子** |
| POST | `/wms/container-bind-record-main/create` | createContainerBindRecordMain | **器具绑定记录主** |
| PUT | `/wms/container-bind-record-main/update` | updateContainerBindRecordMain | **器具绑定记录主** |
| DELETE | `/wms/container-bind-record-main/delete` | deleteContainerBindRecordMain | **器具绑定记录主** |
| GET | `/wms/container-bind-record-main/get` | getContainerBindRecordMain | **编号** |
| GET | `/wms/container-bind-record-main/list` | getContainerBindRecordMainList | **编号** |
| GET | `/wms/container-bind-record-main/page` | getContainerBindRecordMainPage | **编号列表** |
| POST | `/wms/container-bind-record-main/senior` | getContainerBindRecordMainSenior | **器具绑定记录主** |
| GET | `/wms/container-bind-record-main/export-excel` | exportContainerBindRecordMainExcel | **器具绑定记录主** |
| POST | `/wms/container-bind-record-main/export-excel-senior` | exportContainerBindRecordMainExcel | **器具绑定记录主** |
| GET | `/wms/container-bind-record-main/get-import-template` | importTemplate | **器具绑定记录主** |
| POST | `/wms/container-bind-record-main/import` | importExcel | **器具绑定记录主** |

### 容器记录 (containerRecord)  (21个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/container-record-detail/create` | createContainerRecordDetail | **器具管理记录子** |
| PUT | `/wms/container-record-detail/update` | updateContainerRecordDetail | **器具管理记录子** |
| DELETE | `/wms/container-record-detail/delete` | deleteContainerRecordDetail | **器具管理记录子** |
| GET | `/wms/container-record-detail/get` | getContainerRecordDetail | **编号** |
| GET | `/wms/container-record-detail/page` | getContainerRecordDetailPage | **编号** |
| POST | `/wms/container-record-detail/senior` | getContainerRecordDetailSenior | **器具管理记录子** |
| GET | `/wms/container-record-detail/export-excel` | exportContainerRecordDetailExcel | **器具管理记录子** |
| POST | `/wms/container-record-main/create` | createContainerRecordMain | **器具管理记录主** |
| PUT | `/wms/container-record-main/update` | updateContainerRecordMain | **器具管理记录主** |
| DELETE | `/wms/container-record-main/delete` | deleteContainerRecordMain | **器具管理记录主** |
| GET | `/wms/container-record-main/get` | getContainerRecordMain | **编号** |
| GET | `/wms/container-record-main/page` | getContainerRecordMainPage | **编号** |
| POST | `/wms/container-record-main/senior` | getContainerRecordMainSenior | **器具管理记录主** |
| GET | `/wms/container-record-main/export-excel` | exportContainerRecordMainExcel | **器具管理记录主** |
| POST | `/wms/container-record-main/export-excel-senior` | exportContainerRecordMainExcelSenior | **器具管理记录主** |
| GET | `/wms/container-record-main/export-excel-init` | exportContainerRecordInitMainExcel | **器具管理记录主** |
| POST | `/wms/container-record-main/export-excel-init-senior` | exportContainerRecordMainInitExcelSenior | **器具管理记录主** |
| GET | `/wms/container-record-main/export-excel-scrap` | exportContainerRecordScrapMainExcel | **器具管理记录主** |
| POST | `/wms/container-record-main/export-excel-scrap-senior` | exportContainerRecordMainScrapExcelSenior | **器具管理记录主** |
| GET | `/wms/container-record-main/get-import-template` | importTemplate | **器具管理记录主** |
| POST | `/wms/container-record-main/import` | importExcel | **器具管理记录主** |

### 容器维修 (containerRepair)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/container-repair-record-detail/create` | createContainerRepairRecordDetail | **器具维修记录子** |
| PUT | `/wms/container-repair-record-detail/update` | updateContainerRepairRecordDetail | **器具维修记录子** |
| DELETE | `/wms/container-repair-record-detail/delete` | deleteContainerRepairRecordDetail | **器具维修记录子** |
| GET | `/wms/container-repair-record-detail/get` | getContainerRepairRecordDetail | **编号** |
| GET | `/wms/container-repair-record-detail/list` | getContainerRepairRecordDetailList | **编号** |
| GET | `/wms/container-repair-record-detail/page` | getContainerRepairRecordDetailPage | **编号列表** |
| POST | `/wms/container-repair-record-detail/senior` | getContainerRepairRecordDetailSenior | **器具维修记录子** |
| GET | `/wms/container-repair-record-detail/export-excel` | exportContainerRepairRecordDetailExcel | **器具维修记录子** |
| GET | `/wms/container-repair-record-detail/get-import-template` | importTemplate | **器具维修记录子** |
| POST | `/wms/container-repair-record-detail/import` | importExcel | **器具维修记录子** |
| POST | `/wms/container-repair-record-main/create` | createContainerRepairRecordMain | **器具维修记录主** |
| PUT | `/wms/container-repair-record-main/update` | updateContainerRepairRecordMain | **器具维修记录主** |
| DELETE | `/wms/container-repair-record-main/delete` | deleteContainerRepairRecordMain | **器具维修记录主** |
| GET | `/wms/container-repair-record-main/get` | getContainerRepairRecordMain | **编号** |
| GET | `/wms/container-repair-record-main/list` | getContainerRepairRecordMainList | **编号** |
| GET | `/wms/container-repair-record-main/page` | getContainerRepairRecordMainPage | **编号列表** |
| GET | `/wms/container-repair-record-main/export-excel` | exportContainerRepairRecordMainExcel | **器具维修记录主** |
| GET | `/wms/container-repair-record-main/get-import-template` | importTemplate | **器具维修记录主** |
| POST | `/wms/container-repair-record-main/import` | importExcel | **器具维修记录主** |

### 容器申请 (containerRequest)  (27个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/container-detail-request/create` | createContainerDetailRequest | **器具管理申请子** |
| PUT | `/wms/container-detail-request/update` | updateContainerDetailRequest | **器具管理申请子** |
| DELETE | `/wms/container-detail-request/delete` | deleteContainerDetailRequest | **器具管理申请子** |
| GET | `/wms/container-detail-request/get` | getContainerDetailRequest | **编号** |
| GET | `/wms/container-detail-request/page` | getContainerDetailRequestPage | **编号** |
| POST | `/wms/container-detail-request/senior` | getContainerDetailRequestSenior | **器具管理申请子** |
| GET | `/wms/container-detail-request/export-excel` | exportContainerDetailRequestExcel | **器具管理申请子** |
| POST | `/wms/container-main-request/create` | createContainerMainRequest | **器具管理申请主** |
| PUT | `/wms/container-main-request/update` | updateContainerMainRequest | **器具管理申请主** |
| DELETE | `/wms/container-main-request/delete` | deleteContainerMainRequest | **器具管理申请主** |
| GET | `/wms/container-main-request/get` | getContainerMainRequest | **编号** |
| GET | `/wms/container-main-request/page` | getContainerMainRequestPage | **编号** |
| POST | `/wms/container-main-request/senior` | getContainerMainRequestSenior | **器具管理申请主** |
| GET | `/wms/container-main-request/export-excel` | exportContainerMainRequestExcel | **器具管理申请主** |
| POST | `/wms/container-main-request/export-excel-senior` | exportContainerMainRequestSeniorExcel | **器具管理申请主** |
| GET | `/wms/container-main-request/get-import-template-initial` | importTemplateInitial | **器具管理申请主** |
| GET | `/wms/container-main-request/get-import-template-create` | importTemplateCreate | **器具管理申请主** |
| GET | `/wms/container-main-request/get-import-template-scrap` | importTemplateScrap | **器具管理申请主** |
| GET | `/wms/container-main-request/get-import-template-return` | importTemplateReturn | **器具管理申请主** |
| GET | `/wms/container-main-request/get-import-template-move` | importTemplateMove | **器具管理申请主** |
| POST | `/wms/container-main-request/import` | importExcel | **器具管理申请主** |
| PUT | `/wms/container-main-request/close` | closeContainerRequestMain | **Excel 文件** |
| PUT | `/wms/container-main-request/reAdd` | reAddContainerRequestMain | **编号** |
| PUT | `/wms/container-main-request/submit` | submitContainerRequestMain | **编号** |
| PUT | `/wms/container-main-request/agree` | agreeContainerRequestMain | **编号** |
| PUT | `/wms/container-main-request/handle` | handleContainerRequestMain | **编号** |
| PUT | `/wms/container-main-request/refused` | refusedContainerRequestMain | **编号** |

### 容器解绑 (containerUnbind)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/container-unbind-record-detail/create` | createContainerUnbindRecordDetail | **器具解绑记录子** |
| PUT | `/wms/container-unbind-record-detail/update` | updateContainerUnbindRecordDetail | **器具解绑记录子** |
| DELETE | `/wms/container-unbind-record-detail/delete` | deleteContainerUnbindRecordDetail | **器具解绑记录子** |
| GET | `/wms/container-unbind-record-detail/get` | getContainerUnbindRecordDetail | **编号** |
| GET | `/wms/container-unbind-record-detail/list` | getContainerUnbindRecordDetailList | **编号** |
| GET | `/wms/container-unbind-record-detail/page` | getContainerUnbindRecordDetailPage | **编号列表** |
| POST | `/wms/container-unbind-record-detail/senior` | getContainerUnbindRecordDetailSenior | **器具解绑记录子** |
| GET | `/wms/container-unbind-record-detail/export-excel` | exportContainerUnbindRecordDetailExcel | **器具解绑记录子** |
| GET | `/wms/container-unbind-record-detail/get-import-template` | importTemplate | **器具解绑记录子** |
| POST | `/wms/container-unbind-record-detail/import` | importExcel | **器具解绑记录子** |
| POST | `/wms/container-unbind-record-main/create` | createContainerUnbindRecordMain | **器具解绑记录主** |
| PUT | `/wms/container-unbind-record-main/update` | updateContainerUnbindRecordMain | **器具解绑记录主** |
| DELETE | `/wms/container-unbind-record-main/delete` | deleteContainerUnbindRecordMain | **器具解绑记录主** |
| GET | `/wms/container-unbind-record-main/get` | getContainerUnbindRecordMain | **编号** |
| GET | `/wms/container-unbind-record-main/list` | getContainerUnbindRecordMainList | **编号** |
| GET | `/wms/container-unbind-record-main/page` | getContainerUnbindRecordMainPage | **编号列表** |
| GET | `/wms/container-unbind-record-main/export-excel` | exportContainerUnbindRecordMainExcel | **器具解绑记录主** |
| GET | `/wms/container-unbind-record-main/get-import-template` | importTemplate | **器具解绑记录主** |
| POST | `/wms/container-unbind-record-main/import` | importExcel | **器具解绑记录主** |

### 容器初始化 (containerinit)  (20个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/container-init-record-detail/create` | createContainerInitRecordDetail | **器具初始化记录子** |
| PUT | `/wms/container-init-record-detail/update` | updateContainerInitRecordDetail | **器具初始化记录子** |
| DELETE | `/wms/container-init-record-detail/delete` | deleteContainerInitRecordDetail | **器具初始化记录子** |
| GET | `/wms/container-init-record-detail/get` | getContainerInitRecordDetail | **编号** |
| GET | `/wms/container-init-record-detail/list` | getContainerInitRecordDetailList | **编号** |
| GET | `/wms/container-init-record-detail/page` | getContainerInitRecordDetailPage | **编号列表** |
| POST | `/wms/container-init-record-detail/senior` | getInventorychangeRecordDetailSenior | **器具初始化记录子** |
| GET | `/wms/container-init-record-detail/export-excel` | exportContainerInitRecordDetailExcel | **器具初始化记录子** |
| GET | `/wms/container-init-record-detail/get-import-template` | importTemplate | **器具初始化记录子** |
| POST | `/wms/container-init-record-detail/import` | importExcel | **器具初始化记录子** |
| POST | `/wms/container-init-record-main/create` | createContainerInitRecordMain | **器具初始化记录主** |
| PUT | `/wms/container-init-record-main/update` | updateContainerInitRecordMain | **器具初始化记录主** |
| DELETE | `/wms/container-init-record-main/delete` | deleteContainerInitRecordMain | **器具初始化记录主** |
| GET | `/wms/container-init-record-main/get` | getContainerInitRecordMain | **编号** |
| GET | `/wms/container-init-record-main/list` | getContainerInitRecordMainList | **编号** |
| GET | `/wms/container-init-record-main/page` | getContainerInitRecordMainPage | **编号列表** |
| GET | `/wms/container-init-record-main/export-excel` | exportContainerInitRecordMainExcel | **器具初始化记录主** |
| GET | `/wms/container-init-record-main/export-excel-senior` | exportContainerInitRecordMainExcel | **器具初始化记录主** |
| GET | `/wms/container-init-record-main/get-import-template` | importTemplate | **器具初始化记录主** |
| POST | `/wms/container-init-record-main/import` | importExcel | **器具初始化记录主** |

### 盘点任务 (countJob)  (21个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/count-job-detail/page` | getCountJobDetailPage | **盘点任务子** |
| POST | `/wms/count-job-detail/senior` | getCountJobDetailSenior | **盘点任务子** |
| GET | `/wms/count-job-main/page` | getCountJobMainPage | **盘点任务主** |
| POST | `/wms/count-job-main/senior` | getCountJobMainSenior | **盘点任务主** |
| POST | `/wms/count-job-main/senior-pda` | getCountJobMainSeniorPda | **盘点任务主** |
| GET | `/wms/count-job-main/export-excel` | exportCountJobMainExcel | **盘点任务主** |
| POST | `/wms/count-job-main/export-excel-senior` | exportCountJobMainSeniorExcel | **盘点任务主** |
| GET | `/wms/count-job-main/getCountJobById` | getCountJobById | **盘点任务主** |
| GET | `/wms/count-job-main/getCountJobById-pda` | getCountJobByIdPda | **编号** |
| POST | `/wms/count-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/count-job-main/accept` | acceptCountJobMain | **盘点任务主** |
| PUT | `/wms/count-job-main/abandon` | abandonInspectJobMain | **盘点任务主** |
| PUT | `/wms/count-job-main/close` | closeInspectJobMain | **盘点任务主** |
| PUT | `/wms/count-job-main/execute` | executeJobMain | **盘点任务主** |
| PUT | `/wms/count-job-main/validateItem` | validateItem | **编号** |
| PUT | `/wms/count-job-main/finish` | finishJobMain | **编号** |
| GET | `/wms/count-job-main/export-excel-single` | exportCountJobMainExcelSingle | **盘点任务主** |
| GET | `/wms/count-job-main/get-import-template` | importTemplate | **盘点任务主** |
| POST | `/wms/count-job-main/import` | importSingleExcel | **盘点任务主** |
| POST | `/wms/count-job-main/lineTypeImport` | lineTypeImport | **Excel 文件** |
| GET | `/wms/count-job-main/get-lineTypeImport-template` | lineTypeImportTemplate | **Excel 文件** |

### 盘点计划 (countPlan)  (21个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/count-plan-detail/create` | createCountPlanDetail | **盘点计划子** |
| PUT | `/wms/count-plan-detail/update` | updateCountPlanDetail | **盘点计划子** |
| DELETE | `/wms/count-plan-detail/delete` | deleteCountPlanDetail | **盘点计划子** |
| GET | `/wms/count-plan-detail/get` | getCountPlanDetail | **编号** |
| GET | `/wms/count-plan-detail/page` | getCountPlanDetailPage | **编号** |
| POST | `/wms/count-plan-detail/senior` | getCountPlanDetailSenior | **盘点计划子** |
| POST | `/wms/count-plan-main/create` | createCountPlanMain | **盘点计划主** |
| PUT | `/wms/count-plan-main/update` | updateCountPlanMain | **盘点计划主** |
| DELETE | `/wms/count-plan-main/delete` | deleteCountPlanMain | **盘点计划主** |
| GET | `/wms/count-plan-main/get` | getCountPlanMain | **编号** |
| GET | `/wms/count-plan-main/page` | getCountPlanMainPage | **编号** |
| POST | `/wms/count-plan-main/senior` | getCountPlanMainSenior | **盘点计划主** |
| GET | `/wms/count-plan-main/export-excel` | exportCountPlanMainExcel | **盘点计划主** |
| POST | `/wms/count-plan-main/export-excel-senior` | exportCountPlanMainSeniorExcel | **盘点计划主** |
| PUT | `/wms/count-plan-main/close` | closeCountPlanMain | **盘点计划主** |
| PUT | `/wms/count-plan-main/open` | openCountPlanMain | **盘点计划主** |
| PUT | `/wms/count-plan-main/submit` | submitCountPlanMain | **盘点计划主** |
| PUT | `/wms/count-plan-main/agree` | agreeCountPlanMain | **盘点计划主** |
| PUT | `/wms/count-plan-main/reject` | rejectCountPlanMain | **盘点计划主** |
| PUT | `/wms/count-plan-main/publish` | publishCountPlanMain | **盘点计划主** |
| PUT | `/wms/count-plan-main/resetting` | resettingCountPlanMain | **盘点计划主** |

### 盘点记录 (countRecord)  (6个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/count-record-detail/page` | getCountRecordDetailPage | **盘点记录子** |
| POST | `/wms/count-record-detail/senior` | getCountRecordDetailSenior | **盘点记录子** |
| GET | `/wms/count-record-main/page` | getCountRecordMainPage | **盘点记录主** |
| POST | `/wms/count-record-main/senior` | getCountRecordMainSenior | **盘点记录主** |
| GET | `/wms/count-record-main/export-excel` | exportCountRecordMainExcel | **盘点记录主** |
| POST | `/wms/count-record-main/export-excel-senior` | exportCountRecordMainSeniorExcel | **盘点记录主** |

### 盘点申请 (countRequest)  (25个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/count-request-detail/create` | createCountRequestDetail | **盘点申请子** |
| PUT | `/wms/count-request-detail/update` | updateCountRequestDetail | **盘点申请子** |
| DELETE | `/wms/count-request-detail/delete` | deleteCountRequestDetail | **盘点申请子** |
| GET | `/wms/count-request-detail/get` | getCountRequestDetail | **编号** |
| GET | `/wms/count-request-detail/page` | getCountRequestDetailPage | **编号** |
| POST | `/wms/count-request-detail/senior` | getCountRequestDetailSenior | **盘点申请子** |
| POST | `/wms/count-request-main/create` | createCountRequestMain | **盘点申请主** |
| PUT | `/wms/count-request-main/update` | updateCountRequestMain | **盘点申请主** |
| GET | `/wms/count-request-main/get` | getCountRequestMain | **盘点申请主** |
| GET | `/wms/count-request-main/page` | getCountRequestMainPage | **编号** |
| POST | `/wms/count-request-main/senior` | getCountRequestMainSenior | **盘点申请主** |
| GET | `/wms/count-request-main/export-excel` | exportCountRequestMainExcel | **盘点申请主** |
| POST | `/wms/count-request-main/export-excel-senior` | exportCountRequestMainSeniorExcel | **盘点申请主** |
| GET | `/wms/count-request-main/get-import-template` | importTemplate | **盘点申请主** |
| POST | `/wms/count-request-main/import` | importExcel | **盘点申请主** |
| PUT | `/wms/count-request-main/close` | closeCountRequestMain | **Excel 文件** |
| PUT | `/wms/count-request-main/reAdd` | openCountRequestMain | **编号** |
| PUT | `/wms/count-request-main/submit` | submitCountRequestMain | **编号** |
| PUT | `/wms/count-request-main/agree` | agreeCountRequestMain | **编号** |
| PUT | `/wms/count-request-main/handle` | handleCountRequestMain | **编号** |
| PUT | `/wms/count-request-main/refused` | refusedCountRequestMain | **编号** |
| PUT | `/wms/count-request-main/reCount` | createReCountJob | **编号** |
| PUT | `/wms/count-request-main/supervise` | createSuperviseCountJob | **盘点申请主** |
| PUT | `/wms/count-request-main/generateCountadjustRequest` | generateCountadjustRequest | **盘点申请主** |
| PUT | `/wms/count-request-main/thaw` | thaw | **编号** |

### 盘点调整记录 (countadjustRecord)  (6个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/countadjust-record-detail/page` | getCountadjustRecordDetailPage | **盘点调整记录子** |
| POST | `/wms/countadjust-record-detail/senior` | getCountadjustRecordDetailSenior | **盘点调整记录子** |
| GET | `/wms/countadjust-record-main/page` | getCountadjustRecordMainPage | **盘点调整记录主** |
| POST | `/wms/countadjust-record-main/senior` | getItembasicSenior | **盘点调整记录主** |
| GET | `/wms/countadjust-record-main/export-excel` | exportCountadjustRecordMainExcel | **盘点调整记录主** |
| POST | `/wms/countadjust-record-main/export-excel-senior` | exportCountadjustRecordMainSeniorExcel | **盘点调整记录主** |

### 盘点调整申请 (countadjustRequest)  (15个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/countadjust-request-detail/page` | getCountadjustRequestDetailPage | **盘点调整申请子** |
| POST | `/wms/countadjust-request-detail/senior` | getCountadjustRequestDetailSenior | **盘点调整申请子** |
| DELETE | `/wms/countadjust-request-detail/delete` | delete | **盘点调整申请子** |
| GET | `/wms/countadjust-request-main/page` | getCountadjustRequestMainPage | **盘点调整申请主** |
| POST | `/wms/countadjust-request-main/senior` | getCountadjustRequestMainSenior | **盘点调整申请主** |
| GET | `/wms/countadjust-request-main/export-excel` | exportCountadjustRequestMainExcel | **盘点调整申请主** |
| POST | `/wms/countadjust-request-main/export-excel-senior` | exportCountadjustRequestMainSeniorExcel | **盘点调整申请主** |
| PUT | `/wms/countadjust-request-main/close` | closeCountadjustRequestMain | **盘点调整申请主** |
| PUT | `/wms/countadjust-request-main/reAdd` | openCountadjustRequestMain | **编号** |
| PUT | `/wms/countadjust-request-main/submit` | submitCountadjustRequestMain | **编号** |
| PUT | `/wms/countadjust-request-main/agree` | agreeCountadjustRequestMain | **编号** |
| PUT | `/wms/countadjust-request-main/handle` | handleCountadjustRequestMain | **编号** |
| PUT | `/wms/countadjust-request-main/refused` | abortCountadjustRequestMain | **编号** |
| GET | `/wms/countadjust-request-main/generateQadData` | generateQadData | **编号** |
| GET | `/wms/countadjust-request-main/iceItem` | iceItem | **编号** |

### 汇率 (currencyexchange)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/currencyexchange/create` | createCurrencyexchange | **货币转换** |
| PUT | `/wms/currencyexchange/update` | updateCurrencyexchange | **货币转换** |
| DELETE | `/wms/currencyexchange/delete` | deleteCurrencyexchange | **货币转换** |
| GET | `/wms/currencyexchange/get` | getCurrencyexchange | **编号** |
| GET | `/wms/currencyexchange/page` | getCurrencyexchangePage | **编号** |
| POST | `/wms/currencyexchange/senior` | getCurrencyexchangeSenior | **货币转换** |
| GET | `/wms/currencyexchange/export-excel` | exportCurrencyexchangeExcel | **货币转换** |
| POST | `/wms/currencyexchange/export-excel-senior` | exportCurrencyexchangeExcel | **货币转换** |
| GET | `/wms/currencyexchange/get-import-template` | importTemplate | **货币转换** |
| POST | `/wms/currencyexchange/import` | importExcel | **货币转换** |

### 客户 (customer)  (12个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/customer/create` | createCustomer | **客户** |
| PUT | `/wms/customer/update` | updateCustomer | **客户** |
| DELETE | `/wms/customer/delete` | deleteCustomer | **客户** |
| GET | `/wms/customer/get` | getCustomer | **编号** |
| GET | `/wms/customer/list` | getCustomerList | **编号** |
| GET | `/wms/customer/page` | getCustomerPage | **客户** |
| POST | `/wms/customer/senior` | getCustomerSenior | **客户** |
| GET | `/wms/customer/export-excel` | exportCustomerExcel | **客户** |
| POST | `/wms/customer/export-excel-senior` | exportCustomerExcel | **客户** |
| GET | `/wms/customer/get-import-template` | importTemplate | **客户** |
| POST | `/wms/customer/import` | importExcel | **客户** |
| GET | `/wms/customer/listByCodes` | getCustomerByCodes | **Excel 文件** |

### 客户交期预测 (customerdeliveryforecast)  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/customer-delivery-forecast/create` | createCustomerDeliveryForecast | **客户发货预测** |
| PUT | `/wms/customer-delivery-forecast/update` | updateCustomerDeliveryForecast | **客户发货预测** |
| DELETE | `/wms/customer-delivery-forecast/delete` | deleteCustomerDeliveryForecast | **客户发货预测** |
| GET | `/wms/customer-delivery-forecast/get` | getCustomerDeliveryForecast | **编号** |
| GET | `/wms/customer-delivery-forecast/page` | getCustomerDeliveryForecastPage | **编号** |
| POST | `/wms/customer-delivery-forecast/senior` | getCustomerDeliveryForecastSenior | **客户发货预测** |
| GET | `/wms/customer-delivery-forecast/export-excel` | exportCustomerDeliveryForecastExcel | **客户发货预测** |
| GET | `/wms/customer-delivery-forecast/get-import-template` | importTemplate | **客户发货预测** |
| POST | `/wms/customer-delivery-forecast/import` | importExcel | **客户发货预测** |

### 客户月台 (customerdock)  (15个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/customerdock/create` | createCustomerdock | **客户月台** |
| PUT | `/wms/customerdock/update` | updateCustomerdock | **客户月台** |
| DELETE | `/wms/customerdock/delete` | deleteCustomerdock | **客户月台** |
| GET | `/wms/customerdock/get` | getCustomerdock | **编号** |
| GET | `/wms/customerdock/page` | getCustomerdockPage | **编号** |
| GET | `/wms/customerdock/deliverPages` | getCustomerdockDeliverPage | **客户月台** |
| POST | `/wms/customerdock/senior` | getCustomerdockSenior | **客户月台** |
| GET | `/wms/customerdock/export-excel` | exportCustomerdockExcel | **客户月台** |
| POST | `/wms/customerdock/export-excel-senior` | exportCustomerdockExcel | **客户月台** |
| GET | `/wms/customerdock/get-import-template` | importTemplate | **客户月台** |
| POST | `/wms/customerdock/import` | importExcel | **客户月台** |
| GET | `/wms/customerdock/pageCustomerCodeToCustomerDock` | selectPageCustomerCodeToCustomerDock | **Excel 文件** |
| POST | `/wms/customerdock/pageCustomerCodeToCustomerDockSenior` | selectPageCustomerCodeToCustomerDockSenior | **客户代码** |
| GET | `/wms/customerdock/pageCustomerCodeToCustomerDockReceiving` | selectPageCustomerCodeToCustomerDockReceiving | **客户月台** |
| POST | `/wms/customerdock/pageCustomerCodeToCustomerDockReceivingSenior` | selectPageCustomerCodeToCustomerDockReceivingSenior | **客户代码** |

### 客户货品 (customeritem)  (13个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/customeritem/create` | createCustomeritem | **客户物料** |
| PUT | `/wms/customeritem/update` | updateCustomeritem | **客户物料** |
| DELETE | `/wms/customeritem/delete` | deleteCustomeritem | **客户物料** |
| GET | `/wms/customeritem/get` | getCustomeritem | **编号** |
| GET | `/wms/customeritem/page` | getCustomeritemPage | **编号** |
| POST | `/wms/customeritem/senior` | getCustomeritemSenior | **客户物料** |
| GET | `/wms/customeritem/pageBusinessTypeToItemCode` | selectBusinessTypeToItemCode | **客户物料** |
| POST | `/wms/customeritem/pageBusinessTypeToLocationSenior` | selectBusinessTypeToItemCodeSenior | **业务类型** |
| GET | `/wms/customeritem/export-excel` | exportCustomeritemExcel | **客户物料** |
| POST | `/wms/customeritem/export-excel-senior` | exportCustomeritemExcel | **客户物料** |
| GET | `/wms/customeritem/get-import-template` | importTemplate | **客户物料** |
| POST | `/wms/customeritem/import` | importExcel | **客户物料** |
| GET | `/wms/customeritem/listByCodes` | getCustomeritemListByCodes | **Excel 文件** |

### 客收记录 (customerreceiptRecord)  (17个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/customerreceipt-record-detail/create` | createCustomerreceiptRecordDetail | **客户收货记录子** |
| PUT | `/wms/customerreceipt-record-detail/update` | updateCustomerreceiptRecordDetail | **客户收货记录子** |
| DELETE | `/wms/customerreceipt-record-detail/delete` | deleteCustomerreceiptRecordDetail | **客户收货记录子** |
| GET | `/wms/customerreceipt-record-detail/get` | getCustomerreceiptRecordDetail | **编号** |
| GET | `/wms/customerreceipt-record-detail/list` | getCustomerreceiptRecordDetailList | **编号** |
| GET | `/wms/customerreceipt-record-detail/page` | getCustomerreceiptRecordDetailPage | **编号列表** |
| POST | `/wms/customerreceipt-record-detail/senior` | getCustomerreceiptRecordDetailSenior | **客户收货记录子** |
| GET | `/wms/customerreceipt-record-detail/export-excel` | exportCustomerreceiptRecordDetailExcel | **客户收货记录子** |
| POST | `/wms/customerreceipt-record-main/create` | createCustomerreceiptRecordMain | **客户收货记录主** |
| PUT | `/wms/customerreceipt-record-main/update` | updateCustomerreceiptRecordMain | **客户收货记录主** |
| DELETE | `/wms/customerreceipt-record-main/delete` | deleteCustomerreceiptRecordMain | **客户收货记录主** |
| GET | `/wms/customerreceipt-record-main/get` | getCustomerreceiptRecordMain | **编号** |
| GET | `/wms/customerreceipt-record-main/list` | getCustomerreceiptRecordMainList | **编号** |
| GET | `/wms/customerreceipt-record-main/page` | getCustomerreceiptRecordMainPage | **编号列表** |
| POST | `/wms/customerreceipt-record-main/senior` | getCustomerreceiptRecordMainSenior | **客户收货记录主** |
| GET | `/wms/customerreceipt-record-main/export-excel` | exportCustomerreceiptRecordMainExcel | **客户收货记录主** |
| POST | `/wms/customerreceipt-record-main/export-excel-senior` | exportCustomerreceiptRecordMainSeniorExcel | **客户收货记录主** |

### 客收申请 (customerreceiptRequest)  (23个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/customerreceipt-request-detail/create` | createCustomerreceiptRequestDetail | **客户收货申请子** |
| PUT | `/wms/customerreceipt-request-detail/update` | updateCustomerreceiptRequestDetail | **客户收货申请子** |
| DELETE | `/wms/customerreceipt-request-detail/delete` | deleteCustomerreceiptRequestDetail | **客户收货申请子** |
| GET | `/wms/customerreceipt-request-detail/get` | getCustomerreceiptRequestDetail | **编号** |
| GET | `/wms/customerreceipt-request-detail/list` | getCustomerreceiptRequestDetailList | **编号** |
| GET | `/wms/customerreceipt-request-detail/page` | getCustomerreceiptRequestDetailPage | **编号列表** |
| POST | `/wms/customerreceipt-request-detail/senior` | getCustomerreceiptRequestDetailSenior | **客户收货申请子** |
| GET | `/wms/customerreceipt-request-detail/export-excel` | exportCustomerreceiptRequestDetailExcel | **客户收货申请子** |
| POST | `/wms/customerreceipt-request-main/create` | createCustomerreceiptRequestMain | **客户收货申请主** |
| PUT | `/wms/customerreceipt-request-main/update` | updateCustomerreceiptRequestMain | **客户收货申请主** |
| DELETE | `/wms/customerreceipt-request-main/delete` | deleteCustomerreceiptRequestMain | **客户收货申请主** |
| GET | `/wms/customerreceipt-request-main/get` | getCustomerreceiptRequestMain | **编号** |
| GET | `/wms/customerreceipt-request-main/list` | getCustomerreceiptRequestMainList | **编号** |
| GET | `/wms/customerreceipt-request-main/page` | getCustomerreceiptRequestMainPage | **编号列表** |
| POST | `/wms/customerreceipt-request-main/senior` | getCustomerreceiptRequestMainSenior | **客户收货申请主** |
| GET | `/wms/customerreceipt-request-main/export-excel` | exportCustomerreceiptRequestMainExcel | **客户收货申请主** |
| POST | `/wms/customerreceipt-request-main/export-excel-senior` | exportCustomerreceiptRequestMainSeniorExcel | **客户收货申请主** |
| PUT | `/wms/customerreceipt-request-main/close` | closeCustomerreceiptRequestMain | **客户收货申请主** |
| PUT | `/wms/customerreceipt-request-main/reAdd` | openCustomerreceiptRequestMain | **编号** |
| PUT | `/wms/customerreceipt-request-main/submit` | submitCustomerreceiptRequestMain | **编号** |
| PUT | `/wms/customerreceipt-request-main/agree` | agreeCustomerreceiptRequestMain | **编号** |
| PUT | `/wms/customerreceipt-request-main/handle` | handleCustomerreceiptRequestMain | **编号** |
| PUT | `/wms/customerreceipt-request-main/refused` | refusedCustomerreceiptRequestMain | **编号** |

### 客退任务 (customerreturnJob)  (23个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/customerreturn-job-detail/create` | createCustomerreturnJobDetail | **客户退货任务子** |
| PUT | `/wms/customerreturn-job-detail/update` | updateCustomerreturnJobDetail | **客户退货任务子** |
| DELETE | `/wms/customerreturn-job-detail/delete` | deleteCustomerreturnJobDetail | **客户退货任务子** |
| GET | `/wms/customerreturn-job-detail/get` | getCustomerreturnJobDetail | **编号** |
| GET | `/wms/customerreturn-job-detail/list` | getCustomerreturnJobDetailList | **编号** |
| GET | `/wms/customerreturn-job-detail/page` | getCustomerreturnJobDetailPage | **编号列表** |
| POST | `/wms/customerreturn-job-detail/senior` | getCustomerreturnJobDetailSenior | **客户退货任务子** |
| GET | `/wms/customerreturn-job-detail/export-excel` | exportCustomerreturnJobDetailExcel | **客户退货任务子** |
| POST | `/wms/customerreturn-job-main/create` | createCustomerreturnJobMain | **客户退货任务主** |
| PUT | `/wms/customerreturn-job-main/update` | updateCustomerreturnJobMain | **客户退货任务主** |
| DELETE | `/wms/customerreturn-job-main/delete` | deleteCustomerreturnJobMain | **客户退货任务主** |
| GET | `/wms/customerreturn-job-main/get` | getCustomerreturnJobMain | **编号** |
| GET | `/wms/customerreturn-job-main/list` | getCustomerreturnJobMainList | **编号** |
| GET | `/wms/customerreturn-job-main/page` | getCustomerreturnJobMainPage | **编号列表** |
| POST | `/wms/customerreturn-job-main/senior` | getCustomerreturnJobMainSenior | **客户退货任务主** |
| GET | `/wms/customerreturn-job-main/export-excel` | exportCustomerreturnJobMainExcel | **客户退货任务主** |
| POST | `/wms/customerreturn-job-main/export-excel-senior` | exportCustomerreturnJobMainSeniorExcel | **客户退货任务主** |
| GET | `/wms/customerreturn-job-main/getCustomerreturnJobById` | getCustomerreturnJobById | **客户退货任务主** |
| POST | `/wms/customerreturn-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/customerreturn-job-main/accept` | acceptCustomerreturnJobMain | **类型数组** |
| PUT | `/wms/customerreturn-job-main/abandon` | abandonCustomerreturnRequestMain | **编号** |
| PUT | `/wms/customerreturn-job-main/close` | closeCustomerreturnRequestMain | **编号** |
| PUT | `/wms/customerreturn-job-main/execute` | closeCustomerreturnRequestMain | **编号** |

### 客退记录 (customerreturnRecord)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/customerreturn-record-detail/create` | createCustomerreturnRecordDetail | **客户退货记录子** |
| PUT | `/wms/customerreturn-record-detail/update` | updateCustomerreturnRecordDetail | **客户退货记录子** |
| DELETE | `/wms/customerreturn-record-detail/delete` | deleteCustomerreturnRecordDetail | **客户退货记录子** |
| GET | `/wms/customerreturn-record-detail/get` | getCustomerreturnRecordDetail | **编号** |
| GET | `/wms/customerreturn-record-detail/list` | getCustomerreturnRecordDetailList | **编号** |
| GET | `/wms/customerreturn-record-detail/page` | getCustomerreturnRecordDetailPage | **编号列表** |
| POST | `/wms/customerreturn-record-detail/senior` | getCustomerreturnRecordDetailSenior | **客户退货记录子** |
| GET | `/wms/customerreturn-record-detail/export-excel` | exportCustomerreturnRecordDetailExcel | **客户退货记录子** |
| POST | `/wms/customerreturn-record-main/create` | createCustomerreturnRecordMain | **客户退货记录主** |
| PUT | `/wms/customerreturn-record-main/update` | updateCustomerreturnRecordMain | **客户退货记录主** |
| DELETE | `/wms/customerreturn-record-main/delete` | deleteCustomerreturnRecordMain | **客户退货记录主** |
| GET | `/wms/customerreturn-record-main/get` | getCustomerreturnRecordMain | **编号** |
| GET | `/wms/customerreturn-record-main/list` | getCustomerreturnRecordMainList | **编号** |
| GET | `/wms/customerreturn-record-main/page` | getCustomerreturnRecordMainPage | **编号列表** |
| POST | `/wms/customerreturn-record-main/senior` | getCustomerreturnRecordMainSenior | **客户退货记录主** |
| GET | `/wms/customerreturn-record-main/export-excel` | exportCustomerreturnRecordMainExcel | **客户退货记录主** |
| POST | `/wms/customerreturn-record-main/export-excel-senior` | exportCustomerreturnRecordMainSeniorExcel | **客户退货记录主** |
| PUT | `/wms/customerreturn-record-main/receive` | receive | **客户退货记录主** |
| PUT | `/wms/customerreturn-record-main/refuse` | refuse | **编号** |

### 客退申请 (customerreturnRequest)  (29个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/customerreturn-request-detail/create` | createCustomerreturnRequestDetail | **客户退货申请子** |
| PUT | `/wms/customerreturn-request-detail/update` | updateCustomerreturnRequestDetail | **客户退货申请子** |
| DELETE | `/wms/customerreturn-request-detail/delete` | deleteCustomerreturnRequestDetail | **客户退货申请子** |
| GET | `/wms/customerreturn-request-detail/get` | getCustomerreturnRequestDetail | **编号** |
| GET | `/wms/customerreturn-request-detail/list` | getCustomerreturnRequestDetailList | **编号** |
| GET | `/wms/customerreturn-request-detail/page` | getCustomerreturnRequestDetailPage | **编号列表** |
| POST | `/wms/customerreturn-request-detail/senior` | getCustomerreturnRequestDetailSenior | **客户退货申请子** |
| GET | `/wms/customerreturn-request-detail/export-excel` | exportCustomerreturnRequestDetailExcel | **客户退货申请子** |
| POST | `/wms/customerreturn-request-main/create` | createCustomerreturnRequestMain | **客户退货申请主** |
| PUT | `/wms/customerreturn-request-main/update` | updateCustomerreturnRequestMain | **客户退货申请主** |
| DELETE | `/wms/customerreturn-request-main/delete` | deleteCustomerreturnRequestMain | **客户退货申请主** |
| GET | `/wms/customerreturn-request-main/get` | getCustomerreturnRequestMain | **编号** |
| GET | `/wms/customerreturn-request-main/list` | getCustomerreturnRequestMainList | **编号** |
| GET | `/wms/customerreturn-request-main/page` | getCustomerreturnRequestMainPage | **编号列表** |
| POST | `/wms/customerreturn-request-main/senior` | getCustomerreturnRequestMainSenior | **客户退货申请主** |
| GET | `/wms/customerreturn-request-main/export-excel` | exportCustomerreturnRequestMainExcel | **客户退货申请主** |
| POST | `/wms/customerreturn-request-main/export-excel-senior` | exportCustomerreturnRequestMainSeniorExcel | **客户退货申请主** |
| GET | `/wms/customerreturn-request-main/get-import-template` | importTemplate | **客户退货申请主** |
| POST | `/wms/customerreturn-request-main/import` | importExcel | **客户退货申请主** |
| GET | `/wms/customerreturn-request-main/getCustomerreturnRequestById` | getCustomerreturnRequestById | **Excel 文件** |
| PUT | `/wms/customerreturn-request-main/close` | closeCustomerreturnRequestMain | **编号** |
| PUT | `/wms/customerreturn-request-main/reAdd` | openCustomerreturnRequestMain | **编号** |
| PUT | `/wms/customerreturn-request-main/submit` | submitCustomerreturnRequestMain | **编号** |
| PUT | `/wms/customerreturn-request-main/agree` | agreeCustomerreturnRequestMain | **编号** |
| PUT | `/wms/customerreturn-request-main/handle` | handleCustomerreturnRequestMain | **编号** |
| PUT | `/wms/customerreturn-request-main/refused` | abortCustomerreturnRequestMain | **编号** |
| GET | `/wms/customerreturn-request-main/pageItemCodeToBalance` | selectPageItemCodeToBalance | **编号** |
| POST | `/wms/customerreturn-request-main/pageItemCodeToBalanceSenior` | selectpageItemCodeToBalanceSenior | **客户代码** |
| POST | `/wms/customerreturn-request-main/genLabel` | genLabel | **客户代码** |

### 客户结算记录 (customersettleRecord)  (17个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/customersettle-record-detail/create` | createCustomersettleRecordDetail | **客户结算记录子** |
| PUT | `/wms/customersettle-record-detail/update` | updateCustomersettleRecordDetail | **客户结算记录子** |
| DELETE | `/wms/customersettle-record-detail/delete` | deleteCustomersettleRecordDetail | **客户结算记录子** |
| GET | `/wms/customersettle-record-detail/get` | getCustomersettleRecordDetail | **编号** |
| GET | `/wms/customersettle-record-detail/list` | getCustomersettleRecordDetailList | **编号** |
| GET | `/wms/customersettle-record-detail/page` | getCustomersettleRecordDetailPage | **编号列表** |
| POST | `/wms/customersettle-record-detail/senior` | getCustomersettleRecordDetailSenior | **客户结算记录子** |
| GET | `/wms/customersettle-record-detail/export-excel` | exportCustomersettleRecordDetailExcel | **客户结算记录子** |
| POST | `/wms/customersettle-record-main/create` | createCustomersettleRecordMain | **客户结算记录主** |
| PUT | `/wms/customersettle-record-main/update` | updateCustomersettleRecordMain | **客户结算记录主** |
| DELETE | `/wms/customersettle-record-main/delete` | deleteCustomersettleRecordMain | **客户结算记录主** |
| GET | `/wms/customersettle-record-main/get` | getCustomersettleRecordMain | **编号** |
| GET | `/wms/customersettle-record-main/list` | getCustomersettleRecordMainList | **编号** |
| GET | `/wms/customersettle-record-main/page` | getCustomersettleRecordMainPage | **编号列表** |
| POST | `/wms/customersettle-record-main/senior` | getCustomersettleRecordMainSenior | **客户结算记录主** |
| GET | `/wms/customersettle-record-main/export-excel` | exportCustomersettleRecordMainExcel | **客户结算记录主** |
| POST | `/wms/customersettle-record-main/export-excel-senior` | exportCustomersettleRecordMainSeniorExcel | **客户结算记录主** |

### 客户结算申请 (customersettleRequest)  (25个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/customersettle-request-detail/create` | createCustomersettleRequestDetail | **客户结算申请子** |
| PUT | `/wms/customersettle-request-detail/update` | updateCustomersettleRequestDetail | **客户结算申请子** |
| DELETE | `/wms/customersettle-request-detail/delete` | deleteCustomersettleRequestDetail | **客户结算申请子** |
| GET | `/wms/customersettle-request-detail/get` | getCustomersettleRequestDetail | **编号** |
| GET | `/wms/customersettle-request-detail/list` | getCustomersettleRequestDetailList | **编号** |
| GET | `/wms/customersettle-request-detail/page` | getCustomersettleRequestDetailPage | **编号列表** |
| POST | `/wms/customersettle-request-detail/senior` | getCustomersettleRequestDetailSenior | **客户结算申请子** |
| GET | `/wms/customersettle-request-detail/export-excel` | exportCustomersettleRequestDetailExcel | **客户结算申请子** |
| POST | `/wms/customersettle-request-main/create` | createCustomersettleRequestMain | **客户结算申请主** |
| PUT | `/wms/customersettle-request-main/update` | updateCustomersettleRequestMain | **客户结算申请主** |
| DELETE | `/wms/customersettle-request-main/delete` | deleteCustomersettleRequestMain | **客户结算申请主** |
| GET | `/wms/customersettle-request-main/get` | getCustomersettleRequestMain | **编号** |
| GET | `/wms/customersettle-request-main/list` | getCustomersettleRequestMainList | **编号** |
| GET | `/wms/customersettle-request-main/page` | getCustomersettleRequestMainPage | **编号列表** |
| POST | `/wms/customersettle-request-main/senior` | getCustomersettleRequestMainSenior | **客户结算申请主** |
| GET | `/wms/customersettle-request-main/export-excel` | exportCustomersettleRequestMainExcel | **客户结算申请主** |
| POST | `/wms/customersettle-request-main/export-excel-senior` | exportCustomersettleRequestMainSeniorExcel | **客户结算申请主** |
| GET | `/wms/customersettle-request-main/get-import-template` | importTemplate | **客户结算申请主** |
| POST | `/wms/customersettle-request-main/import` | importExcel | **客户结算申请主** |
| PUT | `/wms/customersettle-request-main/close` | closeCustomersettleRequestMain | **Excel 文件** |
| PUT | `/wms/customersettle-request-main/reAdd` | reAddCustomersettleRequestMain | **编号** |
| PUT | `/wms/customersettle-request-main/submit` | submitCustomersettleRequestMain | **编号** |
| PUT | `/wms/customersettle-request-main/agree` | agreeCustomersettleRequestMain | **编号** |
| PUT | `/wms/customersettle-request-main/handle` | handleCustomersettleRequestMain | **编号** |
| PUT | `/wms/customersettle-request-main/refused` | abortCustomersettleRequestMain | **编号** |

### 发货任务 (deliverJob)  (23个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/deliver-job-detail/create` | createDeliverJobDetail | **发货任务子** |
| PUT | `/wms/deliver-job-detail/update` | updateDeliverJobDetail | **发货任务子** |
| DELETE | `/wms/deliver-job-detail/delete` | deleteDeliverJobDetail | **发货任务子** |
| GET | `/wms/deliver-job-detail/get` | getDeliverJobDetail | **编号** |
| GET | `/wms/deliver-job-detail/list` | getDeliverJobDetailList | **编号** |
| GET | `/wms/deliver-job-detail/page` | getDeliverJobDetailPage | **编号列表** |
| POST | `/wms/deliver-job-detail/senior` | getDeliverJobDetailSenior | **发货任务子** |
| GET | `/wms/deliver-job-detail/export-excel` | exportDeliverJobDetailExcel | **发货任务子** |
| POST | `/wms/deliver-job-main/create` | createDeliverJobMain | **发货任务主** |
| PUT | `/wms/deliver-job-main/update` | updateDeliverJobMain | **发货任务主** |
| DELETE | `/wms/deliver-job-main/delete` | deleteDeliverJobMain | **发货任务主** |
| GET | `/wms/deliver-job-main/get` | getDeliverJobMain | **编号** |
| GET | `/wms/deliver-job-main/list` | getDeliverJobMainList | **编号** |
| GET | `/wms/deliver-job-main/page` | getDeliverJobMainPage | **编号列表** |
| POST | `/wms/deliver-job-main/senior` | getDeliverJobMainSenior | **发货任务主** |
| GET | `/wms/deliver-job-main/export-excel` | exportDeliverJobMainExcel | **发货任务主** |
| POST | `/wms/deliver-job-main/export-excel-senior` | exportDeliverJobMainSeniorExcel | **发货任务主** |
| GET | `/wms/deliver-job-main/getDeliverJobById` | getDeliverJobById | **发货任务主** |
| POST | `/wms/deliver-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/deliver-job-main/accept` | acceptDeliverJobMain | **类型数组** |
| PUT | `/wms/deliver-job-main/abandon` | abandonDeliverJobMain | **编号** |
| PUT | `/wms/deliver-job-main/close` | closeDeliverJobMain | **编号** |
| PUT | `/wms/deliver-job-main/execute` | executeDeliverJobMain | **编号** |

### 发货计划 (deliverPlan)  (27个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/deliver-plan-detail/create` | createDeliverPlanDetail | **发货计划子** |
| PUT | `/wms/deliver-plan-detail/update` | updateDeliverPlanDetail | **发货计划子** |
| DELETE | `/wms/deliver-plan-detail/delete` | deleteDeliverPlanDetail | **发货计划子** |
| GET | `/wms/deliver-plan-detail/get` | getDeliverPlanDetail | **编号** |
| GET | `/wms/deliver-plan-detail/list` | getDeliverPlanDetailList | **编号** |
| GET | `/wms/deliver-plan-detail/page` | getDeliverPlanDetailPage | **编号列表** |
| POST | `/wms/deliver-plan-detail/senior` | getDeliverPlanDetailSenior | **发货计划子** |
| GET | `/wms/deliver-plan-detail/export-excel` | exportDeliverPlanDetailExcel | **发货计划子** |
| GET | `/wms/deliver-plan-detail/detailList` | selectDetailByMasterID | **发货计划子** |
| POST | `/wms/deliver-plan-main/create` | createDeliverPlanMain | **发货计划主** |
| PUT | `/wms/deliver-plan-main/update` | updateDeliverPlanMain | **发货计划主** |
| PUT | `/wms/deliver-plan-main/close` | closeDeliverPlanMain | **发货计划主** |
| PUT | `/wms/deliver-plan-main/open` | openDeliverPlanMain | **发货计划主** |
| PUT | `/wms/deliver-plan-main/submit` | submitDeliverPlanMain | **发货计划主** |
| PUT | `/wms/deliver-plan-main/agree` | agreeDeliverPlanMain | **发货计划主** |
| PUT | `/wms/deliver-plan-main/reject` | rejectDeliverPlanMain | **发货计划主** |
| PUT | `/wms/deliver-plan-main/publish` | publishDeliverPlanMain | **发货计划主** |
| PUT | `/wms/deliver-plan-main/resetting` | resettingDeliverPlanMain | **发货计划主** |
| DELETE | `/wms/deliver-plan-main/delete` | deleteDeliverPlanMain | **发货计划主** |
| GET | `/wms/deliver-plan-main/get` | getDeliverPlanMain | **编号** |
| GET | `/wms/deliver-plan-main/list` | getDeliverPlanMainList | **编号** |
| GET | `/wms/deliver-plan-main/page` | getDeliverPlanMainPage | **编号列表** |
| POST | `/wms/deliver-plan-main/senior` | getDeliverPlanMainSenior | **发货计划主** |
| GET | `/wms/deliver-plan-main/export-excel` | exportDeliverPlanMainExcel | **发货计划主** |
| POST | `/wms/deliver-plan-main/export-excel-senior` | exportPurchasereceiptRequestMainSeniorExcel | **发货计划主** |
| GET | `/wms/deliver-plan-main/get-import-template` | importTemplate | **发货计划主** |
| POST | `/wms/deliver-plan-main/import` | importExcel | **发货计划主** |

### 发货记录 (deliverRecord)  (21个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/deliver-record-detail/create` | createDeliverRecordDetail | **发货记录子** |
| PUT | `/wms/deliver-record-detail/update` | updateDeliverRecordDetail | **发货记录子** |
| DELETE | `/wms/deliver-record-detail/delete` | deleteDeliverRecordDetail | **发货记录子** |
| GET | `/wms/deliver-record-detail/get` | getDeliverRecordDetail | **编号** |
| GET | `/wms/deliver-record-detail/list` | getDeliverRecordDetailList | **编号** |
| GET | `/wms/deliver-record-detail/listToRepeatDeliverReceipt` | getDeliverRecordDetailListToRepeat | **编号列表** |
| GET | `/wms/deliver-record-detail/page` | getDeliverRecordDetailPage | **发货记录子** |
| POST | `/wms/deliver-record-detail/senior` | getDeliverRecordDetailSenior | **发货记录子** |
| GET | `/wms/deliver-record-detail/pageCustomerreturn` | getDeliverRecordDetailPageCustomerreturn | **发货记录子** |
| POST | `/wms/deliver-record-detail/seniorCustomerreturn` | getDeliverRecordDetailSeniorCustomerreturn | **发货记录子** |
| GET | `/wms/deliver-record-detail/export-excel` | exportDeliverRecordDetailExcel | **发货记录子** |
| POST | `/wms/deliver-record-main/create` | createDeliverRecordMain | **发货记录主** |
| PUT | `/wms/deliver-record-main/update` | updateDeliverRecordMain | **发货记录主** |
| DELETE | `/wms/deliver-record-main/delete` | deleteDeliverRecordMain | **发货记录主** |
| GET | `/wms/deliver-record-main/get` | getDeliverRecordMain | **编号** |
| GET | `/wms/deliver-record-main/list` | getDeliverRecordMainList | **编号** |
| GET | `/wms/deliver-record-main/page` | getDeliverRecordMainPage | **编号列表** |
| POST | `/wms/deliver-record-main/senior` | getDeliverRecordMainSenior | **发货记录主** |
| POST | `/wms/deliver-record-main/toDeliverTypeSenior` | getDeliverRecordMainToDeliverTypeSenior | **发货记录主** |
| GET | `/wms/deliver-record-main/export-excel` | exportDeliverRecordMainExcel | **发货记录主** |
| POST | `/wms/deliver-record-main/export-excel-senior` | exportDeliverRecordMainSeniorExcel | **发货记录主** |

### 发货申请 (deliverRequest)  (26个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/deliver-request-detail/create` | createDeliverRequestDetail | **发货申请子** |
| PUT | `/wms/deliver-request-detail/update` | updateDeliverRequestDetail | **发货申请子** |
| DELETE | `/wms/deliver-request-detail/delete` | deleteDeliverRequestDetail | **发货申请子** |
| GET | `/wms/deliver-request-detail/get` | getDeliverRequestDetail | **编号** |
| GET | `/wms/deliver-request-detail/list` | getDeliverRequestDetailList | **编号** |
| GET | `/wms/deliver-request-detail/page` | getDeliverRequestDetailPage | **编号列表** |
| POST | `/wms/deliver-request-detail/senior` | getDeliverRequestDetailSenior | **发货申请子** |
| GET | `/wms/deliver-request-detail/export-excel` | exportDeliverRequestDetailExcel | **发货申请子** |
| POST | `/wms/deliver-request-main/create` | createDeliverRequestMain | **发货申请主** |
| PUT | `/wms/deliver-request-main/update` | updateDeliverRequestMain | **发货申请主** |
| DELETE | `/wms/deliver-request-main/delete` | deleteDeliverRequestMain | **发货申请主** |
| GET | `/wms/deliver-request-main/get` | getDeliverRequestMain | **编号** |
| GET | `/wms/deliver-request-main/list` | getDeliverRequestMainList | **编号** |
| GET | `/wms/deliver-request-main/page` | getDeliverRequestMainPage | **编号列表** |
| POST | `/wms/deliver-request-main/senior` | getDeliverRequestMainSenior | **发货申请主** |
| GET | `/wms/deliver-request-main/export-excel` | exportDeliverRequestMainExcel | **发货申请主** |
| POST | `/wms/deliver-request-main/export-excel-senior` | exportPurchasereceiptRequestMainSeniorExcel | **发货申请主** |
| GET | `/wms/deliver-request-main/getDeliverRequestById` | getDeliverRequestById | **发货申请主** |
| GET | `/wms/deliver-request-main/get-import-template` | importTemplate | **编号** |
| POST | `/wms/deliver-request-main/import` | importExcel | **发货申请主** |
| PUT | `/wms/deliver-request-main/close` | closeDeliverRequestMain | **Excel 文件** |
| PUT | `/wms/deliver-request-main/reAdd` | reAddDeliverRequestMain | **编号** |
| PUT | `/wms/deliver-request-main/submit` | submitDeliverRequestMain | **编号** |
| PUT | `/wms/deliver-request-main/agree` | agreeDeliverRequestMain | **编号** |
| PUT | `/wms/deliver-request-main/handle` | handleDeliverRequestMain | **编号** |
| PUT | `/wms/deliver-request-main/refused` | abortDeliverRequestMain | **编号** |

### 需求预测 (demandforecasting)  (31个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/demandforecasting-detail/create` | createDemandforecastingDetail | **要货预测子** |
| PUT | `/wms/demandforecasting-detail/update` | updateDemandforecastingDetail | **要货预测子** |
| DELETE | `/wms/demandforecasting-detail/delete` | deleteDemandforecastingDetail | **要货预测子** |
| GET | `/wms/demandforecasting-detail/get` | getDemandforecastingDetail | **编号** |
| GET | `/wms/demandforecasting-detail/list` | getDemandforecastingDetailList | **编号** |
| GET | `/wms/demandforecasting-detail/page` | getDemandforecastingDetailPage | **编号列表** |
| GET | `/wms/demandforecasting-detail/queryPageTableHead` | queryPageTableHead | **要货预测子** |
| GET | `/wms/demandforecasting-detail/export-excel` | exportDemandforecastingDetailExcel | **要货预测子** |
| GET | `/wms/demandforecasting-detail/queryVersion` | queryVersion | **要货预测子** |
| POST | `/wms/demandforecasting-detail/queryQADDemandforecasting` | queryQADDemandforecasting | **要货预测子** |
| POST | `/wms/demandforecasting-main/create` | createDemandforecastingMain | **要货预测主** |
| PUT | `/wms/demandforecasting-main/update` | updateDemandforecastingMain | **要货预测主** |
| DELETE | `/wms/demandforecasting-main/delete` | deleteDemandforecastingMain | **要货预测主** |
| GET | `/wms/demandforecasting-main/get` | getDemandforecastingMain | **编号** |
| GET | `/wms/demandforecasting-main/list` | getDemandforecastingMainList | **编号** |
| GET | `/wms/demandforecasting-main/page` | getDemandforecastingMainPage | **编号列表** |
| POST | `/wms/demandforecasting-main/senior` | getDemandforecastingMainSenior | **要货预测主** |
| GET | `/wms/demandforecasting-main/export-excel` | exportDemandforecastingMainExcel | **要货预测主** |
| GET | `/wms/demandforecasting-main/export-excel` | exportDemandforecastingMainExcel | **要货预测主** |
| POST | `/wms/demandforecasting-main/export-excel-senior` | exportDemandforecastingMainSeniorExcel | **要货预测主** |
| GET | `/wms/demandforecasting-main/get-import-template` | importTemplate | **要货预测主** |
| POST | `/wms/demandforecasting-main/import` | importExcel | **要货预测主** |
| POST | `/wms/demandforecasting-main/close` | closeDemandforecastingMain | **Excel 文件** |
| POST | `/wms/demandforecasting-main/open` | openDemandforecastingMain | **编号** |
| POST | `/wms/demandforecasting-main/publish` | publishDemandforecastingMain | **编号** |
| POST | `/wms/demandforecasting-main/wit` | witDemandforecastingMain | **编号** |
| POST | `/wms/demandforecasting-main/queryUserPlanerList` | queryUserPlanerList | **编号** |
| POST | `/wms/demandforecasting-main/querySupplierList` | querySupplierList | **要货预测主** |
| POST | `/wms/demandforecasting-main/queryQadSupplierList` | queryQadSupplierList | **要货预测主** |
| POST | `/wms/demandforecasting-main/queryQadItemCodeList` | queryQadItemCodeList | **要货预测主** |
| POST | `/wms/demandforecasting-main/updateIsRead` | updateIsRead | **要货预测主** |

### 月台 (dock)  (11个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/dock/create` | createDock | **月台** |
| PUT | `/wms/dock/update` | updateDock | **月台** |
| DELETE | `/wms/dock/delete` | deleteDock | **月台** |
| GET | `/wms/dock/get` | getDock | **编号** |
| GET | `/wms/dock/list` | getDockList | **编号** |
| GET | `/wms/dock/page` | getDockPage | **月台** |
| POST | `/wms/dock/senior` | getDockSenior | **月台** |
| GET | `/wms/dock/export-excel` | exportDockExcel | **月台** |
| POST | `/wms/dock/export-excel-senior` | exportDockExcel | **月台** |
| GET | `/wms/dock/get-import-template` | importTemplate | **月台** |
| POST | `/wms/dock/import` | importExcel | **月台** |

### 单据切换 (documentSwitch)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/switch/create` | createSwitch | **单据开关** |
| PUT | `/wms/switch/update` | updateSwitch | **单据开关** |
| DELETE | `/wms/switch/delete` | deleteSwitch | **单据开关** |
| GET | `/wms/switch/get` | getSwitch | **编号** |
| GET | `/wms/switch/getByCode` | getByCode | **编号** |
| POST | `/wms/switch/getSwitchList` | getSwitch | **编号** |
| GET | `/wms/switch/list` | getSwitchList | **编号** |
| POST | `/wms/switch/senior` | getSwitchSenior | **编号列表** |
| GET | `/wms/switch/page` | getSwitchPage | **单据开关** |
| GET | `/wms/switch/export-excel` | exportSwitchExcel | **单据开关** |

### 单据设置 (documentsetting)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/documentsetting/create` | createDocumentsetting | **单据设置** |
| PUT | `/wms/documentsetting/update` | updateDocumentsetting | **单据设置** |
| DELETE | `/wms/documentsetting/delete` | deleteDocumentsetting | **单据设置** |
| GET | `/wms/documentsetting/get` | getDocumentsetting | **编号** |
| GET | `/wms/documentsetting/list` | getDocumentsettingList | **编号** |
| GET | `/wms/documentsetting/page` | getDocumentsettingPage | **编号列表** |
| POST | `/wms/documentsetting/senior` | getDocumentsettingSenior | **单据设置** |
| GET | `/wms/documentsetting/export-excel` | exportDocumentsettingExcel | **单据设置** |
| GET | `/wms/documentsetting/get-import-template` | importTemplate | **单据设置** |
| POST | `/wms/documentsetting/import` | importExcel | **单据设置** |

### 企业 (enterprise)  (8个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/enterprise/create` | createEnterprise | **企业** |
| PUT | `/wms/enterprise/update` | updateEnterprise | **企业** |
| DELETE | `/wms/enterprise/delete` | deleteEnterprise | **企业** |
| GET | `/wms/enterprise/get` | getEnterprise | **编号** |
| GET | `/wms/enterprise/list` | getEnterpriseList | **编号** |
| GET | `/wms/enterprise/page` | getEnterprisePage | **编号列表** |
| POST | `/wms/enterprise/senior` | getEnterpriseSenior | **企业** |
| GET | `/wms/enterprise/export-excel` | exportEnterpriseExcel | **企业** |

### 预期入库 (expectin)  (8个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| DELETE | `/wms/expectin/delete` | deleteExpectin | **预计入库存** |
| DELETE | `/wms/expectin/deleteByJobNumber` | deleteExpectin | **任务号** |
| DELETE | `/wms/expectin/deleteIds` | deleteExpectinIds | **编号** |
| GET | `/wms/expectin/get` | getExpectin | **编号** |
| GET | `/wms/expectin/page` | getExpectinPage | **编号** |
| POST | `/wms/expectin/senior` | getExpectinSenior | **预计入库存** |
| GET | `/wms/expectin/export-excel` | exportExpectinExcel | **预计入库存** |
| POST | `/wms/expectin/export-excel-senior` | exportExpectinExcel | **预计入库存** |

### 预期出库 (expectout)  (8个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| DELETE | `/wms/expectout/delete` | deleteExpectout | **预计出库存** |
| DELETE | `/wms/expectout/deleteByJobNumber` | deleteExpectout | **编号** |
| DELETE | `/wms/expectout/deleteIds` | deleteExpectoutIds | **任务号** |
| GET | `/wms/expectout/get` | getExpectout | **编号** |
| GET | `/wms/expectout/page` | getExpectoutPage | **编号** |
| POST | `/wms/expectout/senior` | getExpectoutSenior | **预计出库存** |
| GET | `/wms/expectout/export-excel` | exportExpectoutExcel | **预计出库存** |
| POST | `/wms/expectout/export-excel-senior` | exportExpectoutExcel | **预计出库存** |

### 首页 (index)  (13个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/index/indexSupplier` | indexSupplier | **首页** |
| GET | `/wms/index/pagePurchasereturnRecordMonth` | pagePurchasereturnRecordMonth | **首页** |
| GET | `/wms/index/pagePurchaseclaimRecordMonth` | pagePurchaseclaimRecordMonth | **首页** |
| GET | `/wms/index/indexMaterial` | indexMaterial | **首页** |
| GET | `/wms/index/pageStagnantBalance` | pageStagnantBalance | **首页** |
| GET | `/wms/index/pageOverdueBalance` | pageOverdueBalance | **首页** |
| GET | `/wms/index/pageWarningBalance` | pageWarningBalance | **首页** |
| GET | `/wms/index/indexProduct` | indexProduct | **首页** |
| GET | `/wms/index/indexProduce` | indexProduce | **首页** |
| GET | `/wms/index/pageProductionToday` | pageProductionToday | **首页** |
| GET | `/wms/index/pageSafeLocation` | pageSafeLocation | **首页** |
| GET | `/wms/index/pageProductputawayJobDetail` | pageProductputawayJobDetail | **首页** |
| GET | `/wms/index/indexPda` | indexPda | **首页** |

### 质检任务 (inspectJob)  (16个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/inspect-job-detail/page` | getInspectJobDetailPage | **检验任务子** |
| POST | `/wms/inspect-job-detail/senior` | getInspectJobDetailSenior | **检验任务子** |
| GET | `/wms/inspect-job-main/page` | getInspectJobMainPage | **检验任务主** |
| POST | `/wms/inspect-job-main/senior` | getInspectJobMainSenior | **检验任务主** |
| GET | `/wms/inspect-job-main/export-excel` | exportInspectJobMainExcel | **检验任务主** |
| POST | `/wms/inspect-job-main/export-excel-senior` | exportInspectJobMainSeniorExcel | **检验任务主** |
| GET | `/wms/inspect-job-main/getInspectJobById` | getInspectJobById | **检验任务主** |
| POST | `/wms/inspect-job-main/getInspectJobPageByStatusAndTime` | getInspectJobPageByStatusAndTime | **编号** |
| POST | `/wms/inspect-job-main/getInspectJobMainSenior` | getInspectJobMainSenior | **今日开始结束时间** |
| POST | `/wms/inspect-job-main/getCountByStatus` | getCountByStatus | **发货单号** |
| PUT | `/wms/inspect-job-main/accept` | acceptInspectJobMain | **检验任务主** |
| PUT | `/wms/inspect-job-main/abandon` | abandonInspectJobMain | **检验任务主** |
| PUT | `/wms/inspect-job-main/close` | closeInspectJobMain | **检验任务主** |
| PUT | `/wms/inspect-job-main/execute` | executeInspectJobMain | **检验任务主** |
| POST | `/wms/inspect-job-main/purchaseReceiptInspectResult` | purchaseReceiptInspectResult | **编号** |
| POST | `/wms/inspect-job-main/purchaseReceiptInspectResultMoni` | purchaseReceiptInspectResultMoni | **检验任务主** |

### 质检记录 (inspectRecord)  (8个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/inspect-record-detail/page` | getInspectRecordDetailPage | **检验记录子** |
| POST | `/wms/inspect-record-detail/senior` | getInspectRecordDetailSenior | **检验记录子** |
| GET | `/wms/inspect-record-main/page` | getInspectRecordMainPage | **检验记录主** |
| POST | `/wms/inspect-record-main/senior` | getInspectRecordMainSenior | **检验记录主** |
| GET | `/wms/inspect-record-main/export-excel` | exportInspectRecordMainExcel | **检验记录主** |
| POST | `/wms/inspect-record-main/export-excel-senior` | exportInspectRecordMainSeniorExcel | **检验记录主** |
| POST | `/wms/inspect-record-main/createPutAwayRequest` | createPutAwayRequest | **检验记录主** |
| POST | `/wms/inspect-record-main/createPutAwayRequestPC` | createPutAwayRequestPC | **检验记录主** |

### 质检申请 (inspectRequest)  (23个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/inspect-request-detail/create` | createInspectRequestDetail | **检验申请子** |
| PUT | `/wms/inspect-request-detail/update` | updateInspectRequestDetail | **检验申请子** |
| DELETE | `/wms/inspect-request-detail/delete` | deleteInspectRequestDetail | **检验申请子** |
| GET | `/wms/inspect-request-detail/get` | getInspectRequestDetail | **编号** |
| GET | `/wms/inspect-request-detail/page` | getInspectRequestDetailPage | **编号** |
| POST | `/wms/inspect-request-detail/senior` | getInspectRequestDetailSenior | **检验申请子** |
| POST | `/wms/inspect-request-main/create` | createInspectRequestMain | **检验申请主** |
| PUT | `/wms/inspect-request-main/update` | updateInspectRequestMain | **检验申请主** |
| DELETE | `/wms/inspect-request-main/delete` | deleteInspectRequestMain | **检验申请主** |
| GET | `/wms/inspect-request-main/get` | getInspectRequestMain | **编号** |
| GET | `/wms/inspect-request-main/page` | getInspectRequestMainPage | **编号** |
| POST | `/wms/inspect-request-main/senior` | getInspectRequestMainSenior | **检验申请主** |
| GET | `/wms/inspect-request-main/export-excel` | exportInspectRequestMainExcel | **检验申请主** |
| POST | `/wms/inspect-request-main/export-excel-senior` | exportInspectRequestMainSeniorExcel | **检验申请主** |
| GET | `/wms/inspect-request-main/get-import-template` | importTemplate | **检验申请主** |
| POST | `/wms/inspect-request-main/import` | importExcel | **检验申请主** |
| PUT | `/wms/inspect-request-main/close` | closeInspectRequestMain | **Excel 文件** |
| PUT | `/wms/inspect-request-main/submit` | submitInspectRequestMain | **编号** |
| PUT | `/wms/inspect-request-main/reAdd` | openInspectRequestMain | **编号** |
| PUT | `/wms/inspect-request-main/agree` | agreeInspectRequestMain | **编号** |
| PUT | `/wms/inspect-request-main/handle` | handleInspectRequestMain | **编号** |
| PUT | `/wms/inspect-request-main/refused` | refusedInspectRequestMain | **编号** |
| GET | `/wms/inspect-request-main/getInspectRequestById` | getInspectRequestById | **编号** |

### 接口信息 (interfaceinfo)  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/interface-info/create` | createInterfaceInfo | **接口调用信息** |
| PUT | `/wms/interface-info/update` | updateInterfaceInfo | **接口调用信息** |
| DELETE | `/wms/interface-info/delete` | deleteInterfaceInfo | **接口调用信息** |
| GET | `/wms/interface-info/get` | getInterfaceInfo | **编号** |
| GET | `/wms/interface-info/page` | getInterfaceInfoPage | **编号** |
| POST | `/wms/interface-info/senior` | getInterfaceInfoSenior | **接口调用信息** |
| GET | `/wms/interface-info/export-excel` | exportInterfaceInfoExcel | **接口调用信息** |
| GET | `/wms/interface-info/get-import-template` | importTemplate | **接口调用信息** |
| POST | `/wms/interface-info/import` | importExcel | **接口调用信息** |

### 库存异动记录 (inventorychangeRecord)  (17个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/inventorychange-record-detail/create` | createInventorychangeRecordDetail | **库存修改记录子** |
| PUT | `/wms/inventorychange-record-detail/update` | updateInventorychangeRecordDetail | **库存修改记录子** |
| DELETE | `/wms/inventorychange-record-detail/delete` | deleteInventorychangeRecordDetail | **库存修改记录子** |
| GET | `/wms/inventorychange-record-detail/get` | getInventorychangeRecordDetail | **编号** |
| GET | `/wms/inventorychange-record-detail/list` | getInventorychangeRecordDetailList | **编号** |
| GET | `/wms/inventorychange-record-detail/page` | getInventorychangeRecordDetailPage | **编号列表** |
| POST | `/wms/inventorychange-record-detail/senior` | getInventorychangeRecordDetailSenior | **库存修改记录子** |
| GET | `/wms/inventorychange-record-detail/export-excel` | exportInventorychangeRecordDetailExcel | **库存修改记录子** |
| POST | `/wms/inventorychange-record-main/create` | createInventorychangeRecordMain | **库存修改记录主** |
| PUT | `/wms/inventorychange-record-main/update` | updateInventorychangeRecordMain | **库存修改记录主** |
| DELETE | `/wms/inventorychange-record-main/delete` | deleteInventorychangeRecordMain | **库存修改记录主** |
| GET | `/wms/inventorychange-record-main/get` | getInventorychangeRecordMain | **编号** |
| GET | `/wms/inventorychange-record-main/list` | getInventorychangeRecordMainList | **编号** |
| GET | `/wms/inventorychange-record-main/page` | getInventorychangeRecordMainPage | **编号列表** |
| POST | `/wms/inventorychange-record-main/senior` | getInventorychangeRecordMainSenior | **库存修改记录主** |
| GET | `/wms/inventorychange-record-main/export-excel` | exportInventorychangeRecordMainExcel | **库存修改记录主** |
| POST | `/wms/inventorychange-record-main/export-excel-senior` | exportInventorychangeRecordMainSeniorExcel | **库存修改记录主** |

### 库存异动申请 (inventorychangeRequest)  (25个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/inventorychange-request-detail/create` | createInventorychangeRequestDetail | **库存修改申请子** |
| PUT | `/wms/inventorychange-request-detail/update` | updateInventorychangeRequestDetail | **库存修改申请子** |
| DELETE | `/wms/inventorychange-request-detail/delete` | deleteInventorychangeRequestDetail | **库存修改申请子** |
| GET | `/wms/inventorychange-request-detail/get` | getInventorychangeRequestDetail | **编号** |
| GET | `/wms/inventorychange-request-detail/list` | getInventorychangeRequestDetailList | **编号** |
| GET | `/wms/inventorychange-request-detail/page` | getInventorychangeRequestDetailPage | **编号列表** |
| POST | `/wms/inventorychange-request-detail/senior` | getInventorychangeRequestDetailSenior | **库存修改申请子** |
| GET | `/wms/inventorychange-request-detail/export-excel` | exportInventorychangeRequestDetailExcel | **库存修改申请子** |
| POST | `/wms/inventorychange-request-main/create` | createInventorychangeRequestMain | **库存修改申请主** |
| PUT | `/wms/inventorychange-request-main/update` | updateInventorychangeRequestMain | **库存修改申请主** |
| DELETE | `/wms/inventorychange-request-main/delete` | deleteInventorychangeRequestMain | **库存修改申请主** |
| GET | `/wms/inventorychange-request-main/get` | getInventorychangeRequestMain | **编号** |
| GET | `/wms/inventorychange-request-main/list` | getInventorychangeRequestMainList | **编号** |
| GET | `/wms/inventorychange-request-main/page` | getInventorychangeRequestMainPage | **编号列表** |
| POST | `/wms/inventorychange-request-main/senior` | getInventorychangeRequestMainSenior | **库存修改申请主** |
| GET | `/wms/inventorychange-request-main/export-excel` | exportInventorychangeRequestMainExcel | **库存修改申请主** |
| POST | `/wms/inventorychange-request-main/export-excel-senior` | exportInventorychangeRequestMainSeniorExcel | **库存修改申请主** |
| GET | `/wms/inventorychange-request-main/get-import-template` | importTemplate | **库存修改申请主** |
| POST | `/wms/inventorychange-request-main/import` | importExcel | **库存修改申请主** |
| PUT | `/wms/inventorychange-request-main/close` | closeInventorychangeRequestMain | **Excel 文件** |
| PUT | `/wms/inventorychange-request-main/reAdd` | reAddInventorychangeRequestMain | **编号** |
| PUT | `/wms/inventorychange-request-main/submit` | submitInventorychangeRequestMain | **编号** |
| PUT | `/wms/inventorychange-request-main/agree` | agreeInventorychangeRequestMain | **编号** |
| PUT | `/wms/inventorychange-request-main/handle` | handleInventorychangeRequestMain | **编号** |
| PUT | `/wms/inventorychange-request-main/refused` | abortInventorychangeRequestMain | **编号** |

### 库存初始化记录 (inventoryinitRecord)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/inventoryinit-record-detail/create` | createInventoryinitRecordDetail | **库存初始化记录子** |
| PUT | `/wms/inventoryinit-record-detail/update` | updateInventoryinitRecordDetail | **库存初始化记录子** |
| DELETE | `/wms/inventoryinit-record-detail/delete` | deleteInventoryinitRecordDetail | **库存初始化记录子** |
| GET | `/wms/inventoryinit-record-detail/get` | getInventoryinitRecordDetail | **编号** |
| GET | `/wms/inventoryinit-record-detail/list` | getInventoryinitRecordDetailList | **编号** |
| GET | `/wms/inventoryinit-record-detail/page` | getInventoryinitRecordDetailPage | **编号列表** |
| POST | `/wms/inventoryinit-record-detail/senior` | getInventoryinitRecordDetailSenior | **库存初始化记录子** |
| GET | `/wms/inventoryinit-record-detail/export-excel` | exportInventoryinitRecordDetailExcel | **库存初始化记录子** |
| POST | `/wms/inventoryinit-record-main/create` | createInventoryinitRecordMain | **库存初始化记录主** |
| PUT | `/wms/inventoryinit-record-main/update` | updateInventoryinitRecordMain | **库存初始化记录主** |
| DELETE | `/wms/inventoryinit-record-main/delete` | deleteInventoryinitRecordMain | **库存初始化记录主** |
| GET | `/wms/inventoryinit-record-main/get` | getInventoryinitRecordMain | **编号** |
| GET | `/wms/inventoryinit-record-main/list` | getInventoryinitRecordMainList | **编号** |
| GET | `/wms/inventoryinit-record-main/page` | getInventoryinitRecordMainPage | **编号列表** |
| POST | `/wms/inventoryinit-record-main/senior` | getInventoryinitRecordMainSenior | **库存初始化记录主** |
| GET | `/wms/inventoryinit-record-main/export-excel` | exportInventoryinitRecordMainExcel | **库存初始化记录主** |
| POST | `/wms/inventoryinit-record-main/export-excel-senior` | exportInventoryinitRequestMainSeniorExcel | **库存初始化记录主** |
| POST | `/wms/inventoryinit-record-main/printLabelBatchById` | printLabelBatchById | **库存初始化记录主** |
| POST | `/wms/inventoryinit-record-main/printLabelBatchByMasterId` | printLabelBatchByMasterId | **库存初始化记录主** |

### 库存初始化申请 (inventoryinitRequest)  (26个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/inventoryinit-request-detail/create` | createInventoryinitRequestDetail | **库存初始化申请子** |
| PUT | `/wms/inventoryinit-request-detail/update` | updateInventoryinitRequestDetail | **库存初始化申请子** |
| DELETE | `/wms/inventoryinit-request-detail/delete` | deleteInventoryinitRequestDetail | **库存初始化申请子** |
| GET | `/wms/inventoryinit-request-detail/get` | getInventoryinitRequestDetail | **编号** |
| GET | `/wms/inventoryinit-request-detail/list` | getInventoryinitRequestDetailList | **编号** |
| GET | `/wms/inventoryinit-request-detail/page` | getInventoryinitRequestDetailPage | **编号列表** |
| POST | `/wms/inventoryinit-request-detail/senior` | getInventoryinitRequestDetailSenior | **库存初始化申请子** |
| GET | `/wms/inventoryinit-request-detail/export-excel` | exportInventoryinitRequestDetailExcel | **库存初始化申请子** |
| POST | `/wms/inventoryinit-request-main/create` | createInventoryinitRequestMain | **库存初始化申请主** |
| PUT | `/wms/inventoryinit-request-main/update` | updateInventoryinitRequestMain | **库存初始化申请主** |
| DELETE | `/wms/inventoryinit-request-main/delete` | deleteInventoryinitRequestMain | **库存初始化申请主** |
| GET | `/wms/inventoryinit-request-main/get` | getInventoryinitRequestMain | **编号** |
| GET | `/wms/inventoryinit-request-main/list` | getInventoryinitRequestMainList | **编号** |
| GET | `/wms/inventoryinit-request-main/page` | getInventoryinitRequestMainPage | **编号列表** |
| POST | `/wms/inventoryinit-request-main/senior` | getInventoryinitRequestMainSenior | **库存初始化申请主** |
| GET | `/wms/inventoryinit-request-main/export-excel` | exportInventoryinitRequestMainExcel | **库存初始化申请主** |
| POST | `/wms/inventoryinit-request-main/export-excel-senior` | exportInventoryinitRequestMainSeniorExcel | **库存初始化申请主** |
| GET | `/wms/inventoryinit-request-main/get-import-template` | importTemplate | **库存初始化申请主** |
| POST | `/wms/inventoryinit-request-main/import` | importExcel | **库存初始化申请主** |
| POST | `/wms/inventoryinit-request-main/importLine` | importExcelLine | **Excel 文件** |
| PUT | `/wms/inventoryinit-request-main/close` | closeInventoryinitRequestMain | **Excel 文件** |
| PUT | `/wms/inventoryinit-request-main/reAdd` | openInventoryinitRequestMain | **编号** |
| PUT | `/wms/inventoryinit-request-main/submit` | submitInventoryinitRequestMain | **编号** |
| PUT | `/wms/inventoryinit-request-main/agree` | agreeInventoryinitRequestMain | **编号** |
| PUT | `/wms/inventoryinit-request-main/handle` | handleInventoryinitRequestMain | **编号** |
| PUT | `/wms/inventoryinit-request-main/refused` | abortInventoryinitRequestMain | **编号** |

### 移库任务 (inventorymoveJob)  (23个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/inventorymove-job-detail/create` | createInventorymoveJobDetail | **库存转移任务子** |
| PUT | `/wms/inventorymove-job-detail/update` | updateInventorymoveJobDetail | **库存转移任务子** |
| DELETE | `/wms/inventorymove-job-detail/delete` | deleteInventorymoveJobDetail | **库存转移任务子** |
| GET | `/wms/inventorymove-job-detail/get` | getInventorymoveJobDetail | **编号** |
| GET | `/wms/inventorymove-job-detail/list` | getInventorymoveJobDetailList | **编号** |
| GET | `/wms/inventorymove-job-detail/page` | getInventorymoveJobDetailPage | **编号列表** |
| POST | `/wms/inventorymove-job-detail/senior` | getInventorymoveJobDetailSenior | **库存转移任务子** |
| GET | `/wms/inventorymove-job-detail/export-excel` | exportInventorymoveJobDetailExcel | **库存转移任务子** |
| POST | `/wms/inventorymove-job-main/create` | createInventorymoveJobMain | **库存转移任务主** |
| PUT | `/wms/inventorymove-job-main/update` | updateInventorymoveJobMain | **库存转移任务主** |
| DELETE | `/wms/inventorymove-job-main/delete` | deleteInventorymoveJobMain | **库存转移任务主** |
| GET | `/wms/inventorymove-job-main/get` | getInventorymoveJobMain | **编号** |
| GET | `/wms/inventorymove-job-main/list` | getInventorymoveJobMainList | **编号** |
| GET | `/wms/inventorymove-job-main/page` | getInventorymoveJobMainPage | **编号列表** |
| POST | `/wms/inventorymove-job-main/senior` | getInventorymoveJobMainSenior | **库存转移任务主** |
| GET | `/wms/inventorymove-job-main/export-excel` | exportInventorymoveJobMainExcel | **库存转移任务主** |
| POST | `/wms/inventorymove-job-main/export-excel-senior` | exportInventorymoveJobMainSeniorExcel | **库存转移任务主** |
| GET | `/wms/inventorymove-job-main/getInventorymoveJobById` | getInventorymoveJobById | **库存转移任务主** |
| POST | `/wms/inventorymove-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/inventorymove-job-main/accept` | acceptInventorymoveJobMain | **类型数组** |
| PUT | `/wms/inventorymove-job-main/abandon` | abandonInventorymoveJobMain | **编号** |
| PUT | `/wms/inventorymove-job-main/close` | closeInventorymoveJobMain | **编号** |
| PUT | `/wms/inventorymove-job-main/execute` | executeInventorymoveJobMain | **编号** |

### 移库记录 (inventorymoveRecord)  (25个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/inventorymove-record-detail/create` | createInventorymoveRecordDetail | **库存转移记录子** |
| PUT | `/wms/inventorymove-record-detail/update` | updateInventorymoveRecordDetail | **库存转移记录子** |
| DELETE | `/wms/inventorymove-record-detail/delete` | deleteInventorymoveRecordDetail | **库存转移记录子** |
| GET | `/wms/inventorymove-record-detail/get` | getInventorymoveRecordDetail | **编号** |
| GET | `/wms/inventorymove-record-detail/list` | getInventorymoveRecordDetailList | **编号** |
| GET | `/wms/inventorymove-record-detail/page` | getInventorymoveRecordDetailPage | **编号列表** |
| POST | `/wms/inventorymove-record-detail/senior` | getInventorymoveRecordDetailSenior | **库存转移记录子** |
| GET | `/wms/inventorymove-record-detail/export-excel` | exportInventorymoveRecordDetailExcel | **库存转移记录子** |
| POST | `/wms/inventorymove-record-main/create` | createInventorymoveRecordMain | **库存转移记录主** |
| POST | `/wms/inventorymove-record-main/createMove` | createMoveRecordMain | **库存转移记录主** |
| POST | `/wms/inventorymove-record-main/createInventorymoveAGV` | createInventorymoveAGV | **库存转移记录主** |
| PUT | `/wms/inventorymove-record-main/update` | updateInventorymoveRecordMain | **库存转移记录主** |
| DELETE | `/wms/inventorymove-record-main/delete` | deleteInventorymoveRecordMain | **库存转移记录主** |
| GET | `/wms/inventorymove-record-main/get` | getInventorymoveRecordMain | **编号** |
| GET | `/wms/inventorymove-record-main/list` | getInventorymoveRecordMainList | **编号** |
| GET | `/wms/inventorymove-record-main/page` | getInventorymoveRecordMainPage | **编号列表** |
| POST | `/wms/inventorymove-record-main/senior` | getInventorymoveRecordMainSenior | **库存转移记录主** |
| GET | `/wms/inventorymove-record-main/export-excel` | exportInventorymoveRecordMainExcel | **库存转移记录主** |
| POST | `/wms/inventorymove-record-main/export-excel-senior` | exportInventorymoveRecordMainSeniorExcel | **库存转移记录主** |
| GET | `/wms/inventorymove-record-main/get-import-template` | importTemplate | **库存转移记录主** |
| GET | `/wms/inventorymove-record-main/get-import-template-exceptMove` | importTemplateExceptMove | **库存转移记录主** |
| POST | `/wms/inventorymove-record-main/import` | importExcel | **库存转移记录主** |
| POST | `/wms/inventorymove-record-main/importMove` | importExcelMove | **Excel 文件** |
| PUT | `/wms/inventorymove-record-main/receive` | receive | **Excel 文件** |
| PUT | `/wms/inventorymove-record-main/refuse` | refuse | **编号** |

### 移库申请 (inventorymoveRequest)  (28个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/inventorymove-request-detail/create` | createInventorymoveRequestDetail | **库存转移申请子** |
| PUT | `/wms/inventorymove-request-detail/update` | updateInventorymoveRequestDetail | **库存转移申请子** |
| DELETE | `/wms/inventorymove-request-detail/delete` | deleteInventorymoveRequestDetail | **库存转移申请子** |
| GET | `/wms/inventorymove-request-detail/get` | getInventorymoveRequestDetail | **编号** |
| GET | `/wms/inventorymove-request-detail/list` | getInventorymoveRequestDetailList | **编号** |
| GET | `/wms/inventorymove-request-detail/page` | getInventorymoveRequestDetailPage | **编号列表** |
| POST | `/wms/inventorymove-request-detail/senior` | getInventorymoveRequestDetailSenior | **库存转移申请子** |
| GET | `/wms/inventorymove-request-detail/export-excel` | exportInventorymoveRequestDetailExcel | **库存转移申请子** |
| POST | `/wms/inventorymove-request-main/create` | createInventorymoveRequestMain | **库存转移申请主** |
| PUT | `/wms/inventorymove-request-main/update` | updateInventorymoveRequestMain | **库存转移申请主** |
| DELETE | `/wms/inventorymove-request-main/delete` | deleteInventorymoveRequestMain | **库存转移申请主** |
| GET | `/wms/inventorymove-request-main/get` | getInventorymoveRequestMain | **编号** |
| GET | `/wms/inventorymove-request-main/list` | getInventorymoveRequestMainList | **编号** |
| GET | `/wms/inventorymove-request-main/page` | getInventorymoveRequestMainPage | **编号列表** |
| POST | `/wms/inventorymove-request-main/senior` | getInventorymoveRequestMainSenior | **库存转移申请主** |
| GET | `/wms/inventorymove-request-main/getInventorymoveRequestById` | getInventorymoveRequestById | **库存转移申请主** |
| GET | `/wms/inventorymove-request-main/export-excel` | exportInventorymoveRequestMainExcel | **编号** |
| POST | `/wms/inventorymove-request-main/export-excel-senior` | exportInventorymoveRequestMainSeniorExcel | **库存转移申请主** |
| GET | `/wms/inventorymove-request-main/get-import-template` | importTemplate | **库存转移申请主** |
| GET | `/wms/inventorymove-request-main/get-import-template-hold-ok` | importTemplateHoldok | **库存转移申请主** |
| GET | `/wms/inventorymove-request-main/get-import-template-exceptMove` | importTemplateExceptMove | **库存转移申请主** |
| POST | `/wms/inventorymove-request-main/import` | importExcel | **库存转移申请主** |
| PUT | `/wms/inventorymove-request-main/close` | closeInventorymoveRequestMain | **Excel 文件** |
| PUT | `/wms/inventorymove-request-main/reAdd` | reAddInventorymoveRequestMain | **编号** |
| PUT | `/wms/inventorymove-request-main/submit` | submitInventorymoveRequestMain | **编号** |
| PUT | `/wms/inventorymove-request-main/agree` | agreeInventorymoveRequestMain | **编号** |
| PUT | `/wms/inventorymove-request-main/handle` | handleInventorymoveRequestMain | **编号** |
| PUT | `/wms/inventorymove-request-main/refused` | abortInventorymoveRequestMain | **编号** |

### 开票日历 (invoicingcalendar)  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/invoicingcalendar/create` | createInvoicingcalendar | **开票日历** |
| PUT | `/wms/invoicingcalendar/update` | updateInvoicingcalendar | **开票日历** |
| DELETE | `/wms/invoicingcalendar/delete` | deleteInvoicingcalendar | **开票日历** |
| GET | `/wms/invoicingcalendar/get` | getInvoicingcalendar | **编号** |
| GET | `/wms/invoicingcalendar/page` | getInvoicingcalendarPage | **编号** |
| POST | `/wms/invoicingcalendar/senior` | getInvoicingcalendarSenior | **开票日历** |
| GET | `/wms/invoicingcalendar/export-excel` | exportInvoicingcalendarExcel | **开票日历** |
| GET | `/wms/invoicingcalendar/get-import-template` | importTemplate | **开票日历** |
| POST | `/wms/invoicingcalendar/import` | importExcel | **开票日历** |

### 领料任务 (issueJob)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/issue-job-detail/get` | getIssueJobDetail | **发料任务子** |
| GET | `/wms/issue-job-detail/list` | getIssueJobDetailList | **编号** |
| GET | `/wms/issue-job-detail/page` | getIssueJobDetailPage | **编号列表** |
| POST | `/wms/issue-job-detail/senior` | getIssueJobDetailSenior | **发料任务子** |
| GET | `/wms/issue-job-detail/export-excel` | exportIssueJobDetailExcel | **发料任务子** |
| GET | `/wms/issue-job-main/get` | getIssueJobMain | **发料任务主** |
| GET | `/wms/issue-job-main/list` | getIssueJobMainList | **编号** |
| GET | `/wms/issue-job-main/page` | getIssueJobMainPage | **编号列表** |
| POST | `/wms/issue-job-main/senior` | getIssueJobMainSenior | **发料任务主** |
| GET | `/wms/issue-job-main/export-excel` | exportIssueJobMainExcel | **发料任务主** |
| POST | `/wms/issue-job-main/export-excel-senior` | exportIssueJobMainSeniorExcel | **发料任务主** |
| GET | `/wms/issue-job-main/getIssueJobById` | getIssueJobById | **发料任务主** |
| POST | `/wms/issue-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/issue-job-main/accept` | acceptIssueJobMain | **类型数组** |
| PUT | `/wms/issue-job-main/abandon` | abandonIssueJobMain | **发料任务主** |
| PUT | `/wms/issue-job-main/close` | closeIssueJobMain | **发料任务主** |
| PUT | `/wms/issue-job-main/execute` | executeIssueJobMain | **发料任务主** |
| PUT | `/wms/issue-job-main/updateIssueJobConfig` | updateIssueJobConfig | **发料任务主** |
| GET | `/wms/issue-job-main/getIssueJobByProductionline` | getIssueJobByProductionline | **发料任务主** |

### 领料记录 (issueRecord)  (12个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/issue-record-detail/get` | getIssueRecordDetail | **发料记录子** |
| GET | `/wms/issue-record-detail/list` | getIssueRecordDetailList | **编号** |
| GET | `/wms/issue-record-detail/page` | getIssueRecordDetailPage | **编号列表** |
| POST | `/wms/issue-record-detail/senior` | getIssueRecordDetailSenior | **发料记录子** |
| GET | `/wms/issue-record-detail/export-excel` | exportIssueRecordDetailExcel | **发料记录子** |
| POST | `/wms/issue-record-main/create` | createIssueRecordMain | **发料记录主** |
| GET | `/wms/issue-record-main/get` | getIssueRecordMain | **发料记录主** |
| GET | `/wms/issue-record-main/list` | getIssueRecordMainList | **编号** |
| GET | `/wms/issue-record-main/page` | getIssueRecordMainPage | **编号列表** |
| POST | `/wms/issue-record-main/senior` | getIssueRecordMainSenior | **发料记录主** |
| GET | `/wms/issue-record-main/export-excel` | exportIssueRecordMainExcel | **发料记录主** |
| POST | `/wms/issue-record-main/export-excel-senior` | exportIssueRecordMainSeniorExcel | **发料记录主** |

### 领料申请 (issueRequest)  (30个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/issue-request-detail/create` | createIssueRequestDetail | **发料申请子** |
| PUT | `/wms/issue-request-detail/update` | updateIssueRequestDetail | **发料申请子** |
| DELETE | `/wms/issue-request-detail/delete` | deleteIssueRequestDetail | **发料申请子** |
| GET | `/wms/issue-request-detail/get` | getIssueRequestDetail | **编号** |
| GET | `/wms/issue-request-detail/list` | getIssueRequestDetailList | **编号** |
| GET | `/wms/issue-request-detail/page` | getIssueRequestDetailPage | **编号列表** |
| POST | `/wms/issue-request-detail/senior` | getIssueRequestDetailSenior | **发料申请子** |
| POST | `/wms/issue-request-main/create` | createIssueRequestMain | **发料申请主** |
| POST | `/wms/issue-request-main/createPDA` | createIssueRequestMainSetInterval | **发料申请主** |
| PUT | `/wms/issue-request-main/update` | updateIssueRequestMain | **发料申请主** |
| DELETE | `/wms/issue-request-main/delete` | deleteIssueRequestMain | **发料申请主** |
| GET | `/wms/issue-request-main/get` | getIssueRequestMain | **编号** |
| GET | `/wms/issue-request-main/list` | getIssueRequestMainList | **编号** |
| GET | `/wms/issue-request-main/page` | getIssueRequestMainPage | **编号列表** |
| POST | `/wms/issue-request-main/senior` | getIssueRequestMainSenior | **发料申请主** |
| POST | `/wms/issue-request-main/export-excel-senior` | exportIssueRequestMainSeniorExcel | **发料申请主** |
| GET | `/wms/issue-request-main/export-excel` | exportIssueRequestMainExcel | **发料申请主** |
| GET | `/wms/issue-request-main/get-import-template` | importTemplate | **发料申请主** |
| POST | `/wms/issue-request-main/import` | importExcel | **发料申请主** |
| PUT | `/wms/issue-request-main/close` | closeIssueRequestMain | **Excel 文件** |
| PUT | `/wms/issue-request-main/reAdd` | reAddIssueRequestMain | **编号** |
| PUT | `/wms/issue-request-main/submit` | submitIssueRequestMain | **编号** |
| PUT | `/wms/issue-request-main/refused` | refusedIssueRequestMain | **编号** |
| PUT | `/wms/issue-request-main/agree` | agreeIssueRequestMain | **编号** |
| PUT | `/wms/issue-request-main/handle` | handleIssueRequestMain | **编号** |
| GET | `/wms/issue-request-main/get-workshop-productionline-workstation` | getWorkshopProductionlineWorkstation | **编号** |
| GET | `/wms/issue-request-main/getIssueRequestById` | getIssueRequestById | **发料申请主** |
| GET | `/wms/issue-request-main/getBalanceByBatchOffShelf` | getBalanceByBatchOffShelf | **编号** |
| POST | `/wms/issue-request-main/getPackUnit` | getPackUnit | **编号** |
| POST | `/wms/issue-request-main/getCallmaterials` | getCallmaterials | **发料申请主** |

### 货品库区 (itemarea)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/itemarea/create` | createItemarea | **物料库区配置** |
| PUT | `/wms/itemarea/update` | updateItemarea | **物料库区配置** |
| DELETE | `/wms/itemarea/delete` | deleteItemarea | **物料库区配置** |
| GET | `/wms/itemarea/get` | getItemarea | **编号** |
| GET | `/wms/itemarea/list` | getItemareaList | **编号** |
| GET | `/wms/itemarea/page` | getItemareaPage | **编号列表** |
| POST | `/wms/itemarea/senior` | getItemareaPageSenior | **物料库区配置** |
| GET | `/wms/itemarea/export-excel` | exportItemareaExcel | **物料库区配置** |
| POST | `/wms/itemarea/export-excel-senior` | exportItemareaSeniorExcel | **物料库区配置** |
| GET | `/wms/itemarea/get-import-template` | importTemplate | **物料库区配置** |
| POST | `/wms/itemarea/import` | importExcel | **物料库区配置** |
| POST | `/wms/itemarea-detail/create` | createItemareaDetail | **物料库区配置表子** |
| PUT | `/wms/itemarea-detail/update` | updateItemareaDetail | **物料库区配置表子** |
| DELETE | `/wms/itemarea-detail/delete` | deleteItemareaDetail | **物料库区配置表子** |
| GET | `/wms/itemarea-detail/get` | getItemareaDetail | **编号** |
| GET | `/wms/itemarea-detail/list` | getItemareaDetailList | **编号** |
| GET | `/wms/itemarea-detail/page` | getItemareaDetailPage | **编号列表** |
| POST | `/wms/itemarea-detail/senior` | getItemareaDetailPageSenior | **物料库区配置表子** |
| GET | `/wms/itemarea-detail/export-excel` | exportItemareaDetailExcel | **物料库区配置表子** |

### 货品档案 (itembasic)  (22个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/itembasic/create` | createItembasic | **物料基本信息** |
| PUT | `/wms/itembasic/update` | updateItembasic | **物料基本信息** |
| PUT | `/wms/itembasic/updateStatus` | updateItembasicStatus | **物料基本信息** |
| DELETE | `/wms/itembasic/delete` | deleteItembasic | **物料基本信息** |
| GET | `/wms/itembasic/get` | getItembasic | **编号** |
| GET | `/wms/itembasic/getProduct` | getItembasic | **编号** |
| GET | `/wms/itembasic/list` | getItembasicList | **编码** |
| GET | `/wms/itembasic/page` | getItembasicPage | **物料基本信息** |
| POST | `/wms/itembasic/senior` | getItembasicSenior | **物料基本信息** |
| GET | `/wms/itembasic/itembasicPageToFgAndSemi` | selecttembasicPageToFgAndSemi | **物料基本信息** |
| POST | `/wms/itembasic/itembasicPageToFgAndSemiSenior` | selectItembasicPageToFgAndSemiSenior | **物料基本信息** |
| GET | `/wms/itembasic/export-excel` | exportItembasicExcel | **物料基本信息** |
| POST | `/wms/itembasic/export-excel-senior` | exportItembasicExcel | **物料基本信息** |
| GET | `/wms/itembasic/get-import-template` | importTemplate | **物料基本信息** |
| POST | `/wms/itembasic/import` | importExcel | **物料基本信息** |
| GET | `/wms/itembasic/pageTypeToItembasic` | selectTypeToItembasic | **Excel 文件** |
| POST | `/wms/itembasic/pageTypeToItembasicSenior` | selectTypeToItembasicSenior | **物料基本信息** |
| GET | `/wms/itembasic/pageConfigToItembasic` | selectConfigToItembasic | **物料基本信息** |
| POST | `/wms/itembasic/pageConfigToItembasicSenior` | selectConfigToItembasicSenior | **物料基本信息** |
| POST | `/wms/itembasic/queryItemCodeInfo` | queryItemCodeInfo | **物料基本信息** |
| GET | `/wms/itembasic/listByCodes` | getItembasicListByCodes | **物料基本信息** |
| GET | `/wms/itembasic/selectContainermanageItemCode` | selectContainermanageItemCode | **编号列表** |

### 货品包装 (itempackage)  (21个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/itempackage/create` | createItempackaging | **物料包装信息 ** |
| PUT | `/wms/itempackage/update` | updateItempackaging | **物料包装信息 ** |
| DELETE | `/wms/itempackage/delete` | deleteItempackaging | **物料包装信息 ** |
| GET | `/wms/itempackage/get` | getItempackaging | **编号** |
| GET | `/wms/itempackage/page` | getItempackagingPage | **编号** |
| POST | `/wms/itempackage/senior` | getItempackagingSenior | **物料包装信息 ** |
| GET | `/wms/itempackage/pageTree` | getItempackagingPageTree | **物料包装信息 ** |
| POST | `/wms/itempackage/seniorTree` | getItempackagingSeniorTree | **物料包装信息 ** |
| GET | `/wms/itempackage/pageTreeSCP` | getItempackagingPageTreeSCP | **物料包装信息 ** |
| POST | `/wms/itempackage/seniorTreeSCP` | getItempackagingSeniorTreeSCP | **物料包装信息 ** |
| GET | `/wms/itempackage/pageBySupplierdeliver` | getItempackagingPageBySupplierdeliver | **物料包装信息 ** |
| GET | `/wms/itempackage/pageByCustomerreturn` | getItempackagingPageByCustomerreturn | **物料包装信息 ** |
| POST | `/wms/itempackage/seniorByCustomerreturn` | getItempackagingSeniorByCustomerreturn | **物料包装信息 ** |
| POST | `/wms/itempackage/seniorByProductreceipt` | getItempackagingSeniorByProductreceipt | **物料包装信息 ** |
| GET | `/wms/itempackage/pageByProductreceipt` | getItempackagingPageByProductreceipt | **物料包装信息 ** |
| POST | `/wms/itempackage/seniorBySupplierdeliver` | getItempackagingSeniorBySupplierdeliver | **物料包装信息 ** |
| GET | `/wms/itempackage/export-excel` | exportItempackagingExcel | **物料包装信息 ** |
| POST | `/wms/itempackage/export-excel-senior` | exportItempackagingExcel | **物料包装信息 ** |
| GET | `/wms/itempackage/get-import-template` | importTemplate | **物料包装信息 ** |
| POST | `/wms/itempackage/import` | importExcel | **物料包装信息 ** |
| GET | `/wms/itempackage/listByCodes` | getItempackageListByCodes | **Excel 文件** |

### 货品仓库 (itemwarehouse)  (11个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/itemwarehouse/create` | createItemwarehouse | **物料仓库默认配置** |
| PUT | `/wms/itemwarehouse/update` | updateItemwarehouse | **物料仓库默认配置** |
| DELETE | `/wms/itemwarehouse/delete` | deleteItemwarehouse | **物料仓库默认配置** |
| GET | `/wms/itemwarehouse/get` | getItemwarehouse | **编号** |
| GET | `/wms/itemwarehouse/list` | getItemwarehouseList | **编号** |
| GET | `/wms/itemwarehouse/page` | getItemwarehousePage | **编号列表** |
| POST | `/wms/itemwarehouse/senior` | getItemwarehouseSenior | **物料仓库默认配置** |
| GET | `/wms/itemwarehouse/export-excel` | exportItemwarehouseExcel | **物料仓库默认配置** |
| POST | `/wms/itemwarehouse/export-excel-senior` | exportItemwarehouseExcel | **物料仓库默认配置** |
| GET | `/wms/itemwarehouse/get-import-template` | importTemplate | **物料仓库默认配置** |
| POST | `/wms/itemwarehouse/import` | importExcel | **物料仓库默认配置** |

### 作业设置 (jobsetting)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/jobsetting/create` | createJobsetting | **任务设置** |
| PUT | `/wms/jobsetting/update` | updateJobsetting | **任务设置** |
| DELETE | `/wms/jobsetting/delete` | deleteJobsetting | **任务设置** |
| GET | `/wms/jobsetting/get` | getJobsetting | **编号** |
| GET | `/wms/jobsetting/list` | getJobsettingList | **编号** |
| GET | `/wms/jobsetting/page` | getJobsettingPage | **编号列表** |
| GET | `/wms/jobsetting/export-excel` | exportJobsettingExcel | **任务设置** |
| POST | `/wms/jobsetting/senior` | getJobsettingSenior | **任务设置** |
| GET | `/wms/jobsetting/get-import-template` | importTemplate | **任务设置** |
| POST | `/wms/jobsetting/import` | importExcel | **任务设置** |

### 标签模板 (labelBarbasic)  (11个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/barbasic/create` | createBarbasic | **采购件标签/制造件标签/器具标签/库位标签/叫料标签** |
| PUT | `/wms/barbasic/update` | updateBarbasic | **采购件标签/制造件标签/器具标签/库位标签/叫料标签** |
| DELETE | `/wms/barbasic/delete` | deleteBarbasic | **采购件标签/制造件标签/器具标签/库位标签/叫料标签** |
| GET | `/wms/barbasic/get` | getBarbasic | **编号** |
| GET | `/wms/barbasic/list` | getBarbasicList | **编号** |
| GET | `/wms/barbasic/page` | getBarbasicPage | **编号列表** |
| POST | `/wms/barbasic/senior` | getBarbasicSenior | **采购件标签/制造件标签/器具标签/库位标签/叫料标签** |
| GET | `/wms/barbasic/export-excel` | exportBarbasicExcel | **采购件标签/制造件标签/器具标签/库位标签/叫料标签** |
| POST | `/wms/barbasic/export-excel-senior` | exportItembasicExcel | **采购件标签/制造件标签/器具标签/库位标签/叫料标签** |
| GET | `/wms/barbasic/get-import-template` | importTemplate | **采购件标签/制造件标签/器具标签/库位标签/叫料标签** |
| GET | `/wms/barbasic/getBarbasicByPackingNumber` | getBarbasicByPackingNumber | **采购件标签/制造件标签/器具标签/库位标签/叫料标签** |

### 标签类型 (labeltype)  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/labeltype/create` | createlabeltype | **标签定义** |
| PUT | `/wms/labeltype/update` | updatetLabeltype | **标签定义** |
| DELETE | `/wms/labeltype/delete` | deleteLabeltype | **标签定义** |
| GET | `/wms/labeltype/get` | gettype | **编号** |
| GET | `/wms/labeltype/list` | getLabeltypeList | **编号** |
| GET | `/wms/labeltype/page` | gettypePage | **编号列表** |
| POST | `/wms/labeltype/senior` | getLabeltypeSenior | **标签定义** |
| GET | `/wms/labeltype/export-excel` | exporttypeExcel | **标签定义** |
| GET | `/wms/labeltype/getDetailsByHeader` | getDetailsByHeader | **标签定义** |

### 库位 (location)  (38个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/location/create` | createLocation | **库位** |
| PUT | `/wms/location/update` | updateLocation | **库位** |
| DELETE | `/wms/location/delete` | deleteLocation | **库位** |
| GET | `/wms/location/get` | getLocation | **编号** |
| GET | `/wms/location/list` | getLocationList | **编号** |
| GET | `/wms/location/page` | getLocationPage | **库位** |
| GET | `/wms/location/Mpage` | getMLocationPage | **库位** |
| GET | `/wms/location/pageForRepleinsh` | pageForRepleinsh | **库位** |
| POST | `/wms/location/senior` | getLocationSenior | **库位** |
| POST | `/wms/location/Msenior` | getLocationMSenior | **库位** |
| POST | `/wms/location/getForRepleinshSenior` | getForRepleinshSenior | **库位** |
| GET | `/wms/location/export-excel` | exportLocationExcel | **库位** |
| POST | `/wms/location/export-excel-senior` | exportLocationExcel | **库位** |
| GET | `/wms/location/get-import-template` | importTemplate | **库位** |
| POST | `/wms/location/import` | importExcel | **库位** |
| POST | `/wms/location/recommendLocation` | inspectLocation | **Excel 文件** |
| POST | `/wms/location/recommendLocationNew` | inspectLocationNew | **库位** |
| POST | `/wms/location/recommendLocationEmpty` | recommendLocationEmpty | **库位** |
| POST | `/wms/location/recommendLocationNew` | inspectLocationNew | **库位** |
| POST | `/wms/location/recommendLocationExpectin` | recommendLocationExpectin | **库位** |
| POST | `/wms/location/recommendLocationRemoveExpectin` | recommendLocationRemoveExpectin | **库位** |
| POST | `/wms/location/checkRecommendLocation` | checkRecommendLocation | **库位** |
| POST | `/wms/location/validate` | validate | **库位** |
| GET | `/wms/location/pageBusinessTypeToLocation` | selectBusinessTypeToLocation | **库位** |
| POST | `/wms/location/pageBusinessTypeToLocationSenior` | selectBusinessTypeToLocationSenior | **业务类型** |
| GET | `/wms/location/pageBusinessTypeOutLocation` | selectBusinessTypeOutLocation | **库位** |
| POST | `/wms/location/pageBusinessTypeOutLocationSenior` | selectBusinessTypeOutLocationSenior | **业务类型** |
| GET | `/wms/location/pageBusinessTypeOutLocationAll` | selectBusinessTypeOutLocationAll | **库位** |
| POST | `/wms/location/pageBusinessTypeOutLocationSeniorAll` | selectBusinessTypeOutLocationSeniorAll | **业务类型** |
| GET | `/wms/location/pageItemAreaToLocation` | selectItemAreaToLocation | **库位** |
| POST | `/wms/location/pageItemAreaToLocationSenior` | selectItemAreaToLocationSenior | **配置** |
| GET | `/wms/location/pageConfigToLocation` | selectConfigToLocation | **库位** |
| POST | `/wms/location/pageConfigToLocationSenior` | selectConfigToLocationSenior | **配置** |
| GET | `/wms/location/pageBusinessTypeToLocation1` | pageBusinessTypeToLocation1 | **库位** |
| POST | `/wms/location/pageBusinessTypeToLocationSenior1` | pageBusinessTypeToLocationSenior1 | **业务类型** |
| GET | `/wms/location/listByCodes` | getLocationByCodes | **库位** |
| GET | `/wms/location/listLocationByCode` | listLocationByCode | **库位** |
| GET | `/wms/location/queryLocationByOverflowAreaTypeByConfig` | queryLocationByOverflowAreaTypeByConfig | **库位** |

### 库位容量 (locationcapacity)  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/locationcapacity/create` | createLocationcapacity | **库位容量** |
| PUT | `/wms/locationcapacity/update` | updateLocationcapacity | **库位容量** |
| DELETE | `/wms/locationcapacity/delete` | deleteLocationcapacity | **库位容量** |
| GET | `/wms/locationcapacity/get` | getLocationcapacity | **编号** |
| GET | `/wms/locationcapacity/list` | getLocationcapacityList | **编号** |
| GET | `/wms/locationcapacity/page` | getLocationcapacityPage | **编号列表** |
| POST | `/wms/locationcapacity/senior` | getLocationcapacitySenior | **库位容量** |
| GET | `/wms/locationcapacity/export-excel` | exportLocationcapacityExcel | **库位容量** |
| POST | `/wms/locationcapacity/export-excel-senior` | exportLocationcapacitySeniorExcel | **库位容量** |

### 库位组 (locationgroup)  (12个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/locationgroup/create` | createLocationgroup | **库位组** |
| PUT | `/wms/locationgroup/update` | updateLocationgroup | **库位组** |
| DELETE | `/wms/locationgroup/delete` | deleteLocationgroup | **库位组** |
| GET | `/wms/locationgroup/get` | getLocationgroup | **编号** |
| GET | `/wms/locationgroup/list` | getLocationgroupList | **编号** |
| GET | `/wms/locationgroup/page` | getLocationgroupPage | **库位组** |
| POST | `/wms/locationgroup/senior` | getLocationgroupSenior | **库位组** |
| GET | `/wms/locationgroup/export-excel` | exportLocationgroupExcel | **库位组** |
| POST | `/wms/locationgroup/export-excel-senior` | exportLocationgroupExcel | **库位组** |
| GET | `/wms/locationgroup/get-import-template` | importTemplate | **库位组** |
| POST | `/wms/locationgroup/import` | importExcel | **库位组** |
| GET | `/wms/locationgroup/ListByCode` | ListByCode | **Excel 文件** |

### MES条码 (mesbarcode)  (11个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/mes-bar-code/create` | createMesBarCode | **生产条码清单** |
| PUT | `/wms/mes-bar-code/update` | updateMesBarCode | **生产条码清单** |
| DELETE | `/wms/mes-bar-code/delete` | deleteMesBarCode | **生产条码清单** |
| GET | `/wms/mes-bar-code/get` | getMesBarCode | **编号** |
| GET | `/wms/mes-bar-code/list` | getMesBarCodeList | **编号** |
| GET | `/wms/mes-bar-code/page` | getMesBarCodePage | **编号列表** |
| POST | `/wms/mes-bar-code/senior` | getMesBarCodeSenior | **生产条码清单** |
| GET | `/wms/mes-bar-code/export-excel` | exportMesBarCodeExcel | **生产条码清单** |
| POST | `/wms/mes-bar-code/export-excel-senior` | exportMesBarCodeExcel | **生产条码清单** |
| GET | `/wms/mes-bar-code/get-import-template` | importTemplate | **生产条码清单** |
| POST | `/wms/mes-bar-code/import` | importExcel | **生产条码清单** |

### MRS (mrs)  (4个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/mrs/demandforecasting-main/create-day` | createDayFromMrs | **外部接口 - MRS 要货预测** |
| POST | `/wms/mrs/demandforecasting-main/create-month` | createMonthFromMrs | **外部接口 - MRS 要货预测** |
| GET | `/wms/mrs/purchase-mrs-statistics/list` | list | **外部接口 - MRS 要货计划汇总统计** |
| POST | `/wms/mrs/purchase-plan-main/create` | createPurchasePlanMainFromMrs | **外部接口 - MRS 要货计划** |

### MSTR (mstr)  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/mstr/create` | createMstr | **产品类信息** |
| PUT | `/wms/mstr/update` | updateMstr | **产品类信息** |
| DELETE | `/wms/mstr/delete` | deleteMstr | **产品类信息** |
| GET | `/wms/mstr/get` | getMstr | **编号** |
| GET | `/wms/mstr/page` | getMstrPage | **编号** |
| POST | `/wms/mstr/senior` | getMstrSenior | **产品类信息** |
| GET | `/wms/mstr/export-excel` | exportMstrExcel | **产品类信息** |
| GET | `/wms/mstr/get-import-template` | importTemplate | **产品类信息** |
| POST | `/wms/mstr/import` | importExcel | **产品类信息** |

### 线下结算记录 (offlinesettlementRecord)  (17个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/offlinesettlement-record-detail/create` | createOfflinesettlementRecordDetail | **下线结算记录子** |
| PUT | `/wms/offlinesettlement-record-detail/update` | updateOfflinesettlementRecordDetail | **下线结算记录子** |
| DELETE | `/wms/offlinesettlement-record-detail/delete` | deleteOfflinesettlementRecordDetail | **下线结算记录子** |
| GET | `/wms/offlinesettlement-record-detail/get` | getOfflinesettlementRecordDetail | **编号** |
| GET | `/wms/offlinesettlement-record-detail/list` | getOfflinesettlementRecordDetailList | **编号** |
| GET | `/wms/offlinesettlement-record-detail/page` | getOfflinesettlementRecordDetailPage | **编号列表** |
| POST | `/wms/offlinesettlement-record-detail/senior` | getOfflinesettlementRecordDetailSenior | **下线结算记录子** |
| GET | `/wms/offlinesettlement-record-detail/export-excel` | exportOfflinesettlementRecordDetailExcel | **下线结算记录子** |
| POST | `/wms/offlinesettlement-record-main/create` | createOfflinesettlementRecordMain | **下线结算记录主** |
| PUT | `/wms/offlinesettlement-record-main/update` | updateOfflinesettlementRecordMain | **下线结算记录主** |
| DELETE | `/wms/offlinesettlement-record-main/delete` | deleteOfflinesettlementRecordMain | **下线结算记录主** |
| GET | `/wms/offlinesettlement-record-main/get` | getOfflinesettlementRecordMain | **编号** |
| GET | `/wms/offlinesettlement-record-main/list` | getOfflinesettlementRecordMainList | **编号** |
| GET | `/wms/offlinesettlement-record-main/page` | getOfflinesettlementRecordMainPage | **编号列表** |
| POST | `/wms/offlinesettlement-record-main/senior` | getOfflinesettlementRecordMainSenior | **下线结算记录主** |
| GET | `/wms/offlinesettlement-record-main/export-excel` | exportOfflinesettlementRecordMainExcel | **下线结算记录主** |
| POST | `/wms/offlinesettlement-record-main/export-excel-senior` | exportOfflinesettlementRecordMainSeniorExcel | **下线结算记录主** |

### 线下结算申请 (offlinesettlementRequest)  (17个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/offlinesettlement-request-detail/create` | createOfflinesettlementRequestDetail | **下线结算申请子** |
| PUT | `/wms/offlinesettlement-request-detail/update` | updateOfflinesettlementRequestDetail | **下线结算申请子** |
| DELETE | `/wms/offlinesettlement-request-detail/delete` | deleteOfflinesettlementRequestDetail | **下线结算申请子** |
| GET | `/wms/offlinesettlement-request-detail/get` | getOfflinesettlementRequestDetail | **编号** |
| GET | `/wms/offlinesettlement-request-detail/list` | getOfflinesettlementRequestDetailList | **编号** |
| GET | `/wms/offlinesettlement-request-detail/page` | getOfflinesettlementRequestDetailPage | **编号列表** |
| POST | `/wms/offlinesettlement-request-detail/senior` | getOfflinesettlementRequestDetailSenior | **下线结算申请子** |
| GET | `/wms/offlinesettlement-request-detail/export-excel` | exportOfflinesettlementRequestDetailExcel | **下线结算申请子** |
| POST | `/wms/offlinesettlement-request-main/create` | createOfflinesettlementRequestMain | **下线结算申请主** |
| PUT | `/wms/offlinesettlement-request-main/update` | updateOfflinesettlementRequestMain | **下线结算申请主** |
| DELETE | `/wms/offlinesettlement-request-main/delete` | deleteOfflinesettlementRequestMain | **下线结算申请主** |
| GET | `/wms/offlinesettlement-request-main/get` | getOfflinesettlementRequestMain | **编号** |
| GET | `/wms/offlinesettlement-request-main/list` | getOfflinesettlementRequestMainList | **编号** |
| GET | `/wms/offlinesettlement-request-main/page` | getOfflinesettlementRequestMainPage | **编号列表** |
| POST | `/wms/offlinesettlement-request-main/senior` | getOfflinesettlementRequestMainSenior | **下线结算申请主** |
| GET | `/wms/offlinesettlement-request-main/export-excel` | exportOfflinesettlementRequestMainExcel | **下线结算申请主** |
| POST | `/wms/offlinesettlement-request-main/export-excel-senior` | exportOfflinesettlementRequestMainSeniorExcel | **下线结算申请主** |

### 线上结算记录 (onlinesettlementRecord)  (14个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/onlinesettlement-record-detail/create` | createOnlinesettlementRecordDetail | **上线结算记录子** |
| PUT | `/wms/onlinesettlement-record-detail/update` | updateOnlinesettlementRecordDetail | **上线结算记录子** |
| DELETE | `/wms/onlinesettlement-record-detail/delete` | deleteOnlinesettlementRecordDetail | **上线结算记录子** |
| GET | `/wms/onlinesettlement-record-detail/get` | getOnlinesettlementRecordDetail | **编号** |
| GET | `/wms/onlinesettlement-record-detail/list` | getOnlinesettlementRecordDetailList | **编号** |
| GET | `/wms/onlinesettlement-record-detail/page` | getOnlinesettlementRecordDetailPage | **编号列表** |
| POST | `/wms/onlinesettlement-record-detail/senior` | getOnlinesettlementRecordDetailSenior | **上线结算记录子** |
| GET | `/wms/onlinesettlement-record-detail/export-excel` | exportOnlinesettlementRecordDetailExcel | **上线结算记录子** |
| GET | `/wms/onlinesettlement-record-main/get` | getOnlinesettlementRecordMain | **上线结算记录主** |
| GET | `/wms/onlinesettlement-record-main/list` | getOnlinesettlementRecordMainList | **编号** |
| GET | `/wms/onlinesettlement-record-main/page` | getOnlinesettlementRecordMainPage | **编号列表** |
| POST | `/wms/onlinesettlement-record-main/senior` | getOnlinesettlementRecordMainSenior | **上线结算记录主** |
| GET | `/wms/onlinesettlement-record-main/export-excel` | exportOnlinesettlementRecordMainExcel | **上线结算记录主** |
| POST | `/wms/onlinesettlement-record-main/export-excel-senior` | exportOnlinesettlementRecordMainSeniorExcel | **上线结算记录主** |

### 线上结算申请 (onlinesettlementRequest)  (17个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/onlinesettlement-request-detail/create` | createOnlinesettlementRequestDetail | **上线结算申请子** |
| PUT | `/wms/onlinesettlement-request-detail/update` | updateOnlinesettlementRequestDetail | **上线结算申请子** |
| DELETE | `/wms/onlinesettlement-request-detail/delete` | deleteOnlinesettlementRequestDetail | **上线结算申请子** |
| GET | `/wms/onlinesettlement-request-detail/get` | getOnlinesettlementRequestDetail | **编号** |
| GET | `/wms/onlinesettlement-request-detail/list` | getOnlinesettlementRequestDetailList | **编号** |
| GET | `/wms/onlinesettlement-request-detail/page` | getOnlinesettlementRequestDetailPage | **编号列表** |
| POST | `/wms/onlinesettlement-request-detail/senior` | getOnlinesettlementRequestDetailSenior | **上线结算申请子** |
| GET | `/wms/onlinesettlement-request-detail/export-excel` | exportOnlinesettlementRequestDetailExcel | **上线结算申请子** |
| POST | `/wms/onlinesettlement-request-main/create` | createOnlinesettlementRequestMain | **上线结算申请主** |
| PUT | `/wms/onlinesettlement-request-main/update` | updateOnlinesettlementRequestMain | **上线结算申请主** |
| DELETE | `/wms/onlinesettlement-request-main/delete` | deleteOnlinesettlementRequestMain | **上线结算申请主** |
| GET | `/wms/onlinesettlement-request-main/get` | getOnlinesettlementRequestMain | **编号** |
| GET | `/wms/onlinesettlement-request-main/list` | getOnlinesettlementRequestMainList | **编号** |
| GET | `/wms/onlinesettlement-request-main/page` | getOnlinesettlementRequestMainPage | **编号列表** |
| POST | `/wms/onlinesettlement-request-main/senior` | getOnlinesettlementRequestMainSenior | **上线结算申请主** |
| GET | `/wms/onlinesettlement-request-main/export-excel` | exportOnlinesettlementRequestMainExcel | **上线结算申请主** |
| POST | `/wms/onlinesettlement-request-main/export-excel-senior` | exportOnlinesettlementRequestMainSeniorExcel | **上线结算申请主** |

### 货主 (owner)  (12个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/owner/create` | createOwner | **货主** |
| PUT | `/wms/owner/update` | updateOwner | **货主** |
| DELETE | `/wms/owner/delete` | deleteOwner | **货主** |
| GET | `/wms/owner/get` | getOwner | **编号** |
| GET | `/wms/owner/page` | getOwnerPage | **编号** |
| GET | `/wms/owner/list` | getOwnerList | **货主** |
| POST | `/wms/owner/senior` | getOwnerSenior | **货主** |
| GET | `/wms/owner/export-excel` | exportOwnerExcel | **货主** |
| POST | `/wms/owner/export-excel-senior` | exportOwnerExcel | **货主** |
| GET | `/wms/owner/get-import-template` | importTemplate | **货主** |
| POST | `/wms/owner/import` | importExcel | **货主** |
| GET | `/wms/owner/listByCodes` | getOwnerByCodes | **Excel 文件** |

### 包裹信息 (packageMassage)  (23个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/package/create` | createPackage | **包装** |
| POST | `/wms/package/createLabel` | createPackageLabel | **包装** |
| PUT | `/wms/package/update` | updatePackage | **包装** |
| DELETE | `/wms/package/delete` | deletePackage | **包装** |
| GET | `/wms/package/get` | getPackage | **编号** |
| GET | `/wms/package/list` | getPackageList | **编号** |
| GET | `/wms/package/page` | getPackagePage | **编号列表** |
| GET | `/wms/package/getBalanceToPackage` | getBalanceToPackage | **包装** |
| GET | `/wms/package/export-excel` | exportPackageExcel | **编号** |
| POST | `/wms/package/export-excel-senior` | exportPackageSeniorExcel | **包装** |
| POST | `/wms/package/senior` | getPackageSenior | **包装** |
| GET | `/wms/package/get-import-template` | importTemplate | **包装** |
| POST | `/wms/package/import` | importExcel | **包装** |
| GET | `/wms/package/getMakeLabel` | getMakeLabel | **Excel 文件** |
| GET | `/wms/package/getLabel` | getPurchaseLabel | **包装** |
| GET | `/wms/package/getLabelDetailPage` | getLabelDetailPage | **包装** |
| GET | `/wms/package/getProductreceiptLabelDetailPage` | getProductreceiptLabelDetailPage | **包装** |
| GET | `/wms/package/getCallMaterialsLabel` | getCallMaterialsLabel | **包装** |
| GET | `/wms/package/queryPackageInfo` | queryPackageTree | **包装** |
| POST | `/wms/package/batchPrintingLable` | batchPrintingLable | **包装** |
| POST | `/wms/package/batchPrintingLables` | batchPrintingLable | **包装** |
| POST | `/wms/package/batchPrintingLablesForDL` | batchPrintingLablesForDL | **包装** |
| GET | `/wms/package/getLabelDetailPageByRecordId` | getLabelDetailPageByRecordId | **包装** |

### 合包 (packagemergemain)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/packagemerge-detail/create` | createPackagemergeDetail | **合包记录子** |
| PUT | `/wms/packagemerge-detail/update` | updatePackagemergeDetail | **合包记录子** |
| DELETE | `/wms/packagemerge-detail/delete` | deletePackagemergeDetail | **合包记录子** |
| GET | `/wms/packagemerge-detail/get` | getPackagemergeDetail | **编号** |
| GET | `/wms/packagemerge-detail/list` | getPackagemergeDetailList | **编号** |
| GET | `/wms/packagemerge-detail/page` | getPackagemergeDetailPage | **编号列表** |
| POST | `/wms/packagemerge-detail/senior` | getPackagemergeDetailSenior | **合包记录子** |
| GET | `/wms/packagemerge-detail/export-excel` | exportPackagemergeDetailExcel | **合包记录子** |
| GET | `/wms/packagemerge-detail/get-import-template` | importTemplate | **合包记录子** |
| POST | `/wms/packagemerge-main/create` | createPackagemergeMain | **合包记录主** |
| PUT | `/wms/packagemerge-main/update` | updatePackagemergeMain | **合包记录主** |
| DELETE | `/wms/packagemerge-main/delete` | deletePackagemergeMain | **合包记录主** |
| GET | `/wms/packagemerge-main/get` | getPackagemergeMain | **编号** |
| GET | `/wms/packagemerge-main/list` | getPackagemergeMainList | **编号** |
| GET | `/wms/packagemerge-main/page` | getPackagemergeMainPage | **编号列表** |
| POST | `/wms/packagemerge-main/senior` | getPackagemergeRecordMainSenior | **合包记录主** |
| GET | `/wms/packagemerge-main/export-excel` | exportPackagemergeMainExcel | **合包记录主** |
| POST | `/wms/packagemerge-main/export-excel-senior` | exportPackagemergeMainSeniorExcel | **合包记录主** |
| GET | `/wms/packagemerge-main/get-import-template` | importTemplate | **合包记录主** |

### 包装完工任务 (packageoverJob)  (24个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/packageover-job-detail/create` | createPackageoverJobDetail | **翻包任务子** |
| PUT | `/wms/packageover-job-detail/update` | updatePackageoverJobDetail | **翻包任务子** |
| DELETE | `/wms/packageover-job-detail/delete` | deletePackageoverJobDetail | **翻包任务子** |
| GET | `/wms/packageover-job-detail/get` | getPackageoverJobDetail | **编号** |
| GET | `/wms/packageover-job-detail/list` | getPackageoverJobDetailList | **编号** |
| GET | `/wms/packageover-job-detail/page` | getPackageoverJobDetailPage | **编号列表** |
| POST | `/wms/packageover-job-detail/senior` | getPackageoverJobDetailSenior | **翻包任务子** |
| GET | `/wms/packageover-job-detail/export-excel` | exportPackageoverJobDetailExcel | **翻包任务子** |
| GET | `/wms/packageover-job-detail/get-import-template` | importTemplate | **翻包任务子** |
| POST | `/wms/packageover-job-main/create` | createPackageoverJobMain | **翻包任务主** |
| PUT | `/wms/packageover-job-main/update` | updatePackageoverJobMain | **翻包任务主** |
| DELETE | `/wms/packageover-job-main/delete` | deletePackageoverJobMain | **翻包任务主** |
| GET | `/wms/packageover-job-main/get` | getPackageoverJobMain | **编号** |
| GET | `/wms/packageover-job-main/list` | getPackageoverJobMainList | **编号** |
| GET | `/wms/packageover-job-main/page` | getPackageoverJobMainPage | **编号列表** |
| POST | `/wms/packageover-job-main/senior` | getPackageoverJobMainSenior | **翻包任务主** |
| GET | `/wms/packageover-job-main/export-excel` | exportPackageoverJobMainExcel | **翻包任务主** |
| GET | `/wms/packageover-job-main/export-excel-senior` | exportPackageoverJobMainSeniorExcel | **翻包任务主** |
| GET | `/wms/packageover-job-main/getPackageoverJobById` | getPackageoverJobById | **翻包任务主** |
| POST | `/wms/packageover-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/packageover-job-main/accept` | acceptPackageoverJobMain | **类型数组** |
| PUT | `/wms/packageover-job-main/abandon` | abandonPackageoverJobMain | **翻包任务主** |
| PUT | `/wms/packageover-job-main/close` | closePackageoverJobMain | **翻包任务主** |
| PUT | `/wms/packageover-job-main/execute` | executePackageoverJobMain | **翻包任务主** |

### 包装完工记录 (packageoverRecord)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/packageover-record-detail/create` | createPackageoverDetail | **翻包记录子** |
| PUT | `/wms/packageover-record-detail/update` | updatePackageoverDetail | **翻包记录子** |
| DELETE | `/wms/packageover-record-detail/delete` | deletePackageoverDetail | **翻包记录子** |
| GET | `/wms/packageover-record-detail/get` | getPackageoverDetail | **编号** |
| GET | `/wms/packageover-record-detail/list` | getPackageoverDetailList | **编号** |
| GET | `/wms/packageover-record-detail/page` | getPackageoverDetailPage | **编号列表** |
| POST | `/wms/packageover-record-detail/senior` | getPackageoverDetailSenior | **翻包记录子** |
| GET | `/wms/packageover-record-detail/export-excel` | exportPackageoverDetailExcel | **翻包记录子** |
| GET | `/wms/packageover-record-detail/get-import-template` | importTemplate | **翻包记录子** |
| POST | `/wms/packageover-record-main/create` | createPackageoverMain | **翻包记录主** |
| PUT | `/wms/packageover-record-main/update` | updatePackageoverMain | **翻包记录主** |
| DELETE | `/wms/packageover-record-main/delete` | deletePackageoverMain | **翻包记录主** |
| GET | `/wms/packageover-record-main/get` | getPackageoverMain | **编号** |
| GET | `/wms/packageover-record-main/list` | getPackageoverMainList | **编号** |
| GET | `/wms/packageover-record-main/page` | getPackageoverMainPage | **编号列表** |
| POST | `/wms/packageover-record-main/senior` | getPackageoverMainSenior | **翻包记录主** |
| GET | `/wms/packageover-record-main/export-excel` | exportPackageoverMainExcel | **翻包记录主** |
| POST | `/wms/packageover-record-main/export-excel-senior` | exportPackageoverMainSeniorExcel | **翻包记录主** |
| GET | `/wms/packageover-record-main/get-import-template` | importTemplate | **翻包记录主** |

### 包装完工申请 (packageoverRequest)  (26个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/packageover-request-detail/create` | createPackageoverRequestDetail | **翻包申请子** |
| PUT | `/wms/packageover-request-detail/update` | updatePackageoverRequestDetail | **翻包申请子** |
| DELETE | `/wms/packageover-request-detail/delete` | deletePackageoverRequestDetail | **翻包申请子** |
| GET | `/wms/packageover-request-detail/get` | getPackageoverRequestDetail | **编号** |
| GET | `/wms/packageover-request-detail/list` | getPackageoverRequestDetailList | **编号** |
| GET | `/wms/packageover-request-detail/page` | getPackageoverRequestDetailPage | **编号列表** |
| POST | `/wms/packageover-request-detail/senior` | getPackageoverRequestDetailSenior | **翻包申请子** |
| GET | `/wms/packageover-request-detail/export-excel` | exportPackageoverRequestDetailExcel | **翻包申请子** |
| GET | `/wms/packageover-request-detail/get-import-template` | importTemplate | **翻包申请子** |
| POST | `/wms/packageover-request-main/create` | createPackageoverRequestMain | **翻包申请主** |
| PUT | `/wms/packageover-request-main/update` | updatePackageoverRequestMain | **翻包申请主** |
| DELETE | `/wms/packageover-request-main/delete` | deletePackageoverRequestMain | **翻包申请主** |
| GET | `/wms/packageover-request-main/get` | getPackageoverRequestMain | **编号** |
| GET | `/wms/packageover-request-main/list` | getPackageoverRequestMainList | **编号** |
| GET | `/wms/packageover-request-main/page` | getPackageoverRequestMainPage | **编号列表** |
| POST | `/wms/packageover-request-main/senior` | getPackageoverRequestMainSenior | **翻包申请主** |
| GET | `/wms/packageover-request-main/export-excel-senior` | exportPackageoverRequestMainSeniorExcel | **翻包申请主** |
| GET | `/wms/packageover-request-main/export-excel` | exportPackageoverRequestMainExcel | **翻包申请主** |
| GET | `/wms/packageover-request-main/get-import-template` | importTemplate | **翻包申请主** |
| POST | `/wms/packageover-request-main/import` | importExcel | **翻包申请主** |
| PUT | `/wms/packageover-request-main/close` | closePackageoverRequestMain | **Excel 文件** |
| PUT | `/wms/packageover-request-main/reAdd` | reAddPackageoverRequestMain | **编号** |
| PUT | `/wms/packageover-request-main/submit` | submitPackageoverRequestMain | **编号** |
| PUT | `/wms/packageover-request-main/refused` | refusedPackageoverRequestMain | **编号** |
| PUT | `/wms/packageover-request-main/agree` | agreePackageoverRequestMain | **编号** |
| PUT | `/wms/packageover-request-main/handle` | handlePackageoverRequestMain | **编号** |

### 包装追溯 (packageoverretrospect)  (6个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/packageover-retrospect/get` | getPackageoverRetrospect | **翻包追溯记录主** |
| GET | `/wms/packageover-retrospect/list` | getPackageoverRetrospectList | **编号** |
| GET | `/wms/packageover-retrospect/page` | getPackageoverRetrospectPage | **编号列表** |
| POST | `/wms/packageover-retrospect/senior` | getPackageoverRetrospectSenior | **翻包追溯记录主** |
| GET | `/wms/packageover-retrospect/export-excel` | exportPackageoverRetrospectExcel | **翻包追溯记录主** |
| GET | `/wms/packageover-retrospect/export-excel-senior` | exportPackageoverRetrospectSeniorExcel | **翻包追溯记录主** |

### 拆包 (packagesplitmain)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/packagesplit-detail/create` | createPackagesplitDetail | **拆包记录子** |
| PUT | `/wms/packagesplit-detail/update` | updatePackagesplitDetail | **拆包记录子** |
| DELETE | `/wms/packagesplit-detail/delete` | deletePackagesplitDetail | **拆包记录子** |
| GET | `/wms/packagesplit-detail/get` | getPackagesplitDetail | **编号** |
| GET | `/wms/packagesplit-detail/list` | getPackagesplitDetailList | **编号** |
| GET | `/wms/packagesplit-detail/page` | getPackagesplitDetailPage | **编号列表** |
| POST | `/wms/packagesplit-detail/senior` | getPackagesplitDetailSenior | **拆包记录子** |
| GET | `/wms/packagesplit-detail/export-excel` | exportPackagesplitDetailExcel | **拆包记录子** |
| GET | `/wms/packagesplit-detail/get-import-template` | importTemplate | **拆包记录子** |
| POST | `/wms/packagesplit-main/create` | createPackagesplitMain | **拆包记录主** |
| PUT | `/wms/packagesplit-main/update` | updatePackagesplitMain | **拆包记录主** |
| DELETE | `/wms/packagesplit-main/delete` | deletePackagesplitMain | **拆包记录主** |
| GET | `/wms/packagesplit-main/get` | getPackagesplitMain | **编号** |
| GET | `/wms/packagesplit-main/list` | getPackagesplitMainList | **编号** |
| GET | `/wms/packagesplit-main/page` | getPackagesplitMainPage | **编号列表** |
| POST | `/wms/packagesplit-main/senior` | getPackagesplitMainSenior | **拆包记录主** |
| GET | `/wms/packagesplit-main/export-excel` | exportPackagesplitMainExcel | **拆包记录主** |
| POST | `/wms/packagesplit-main/export-excel-senior` | exportPackagesplitMainSeniorExcel | **拆包记录主** |
| GET | `/wms/packagesplit-main/get-import-template` | importTemplate | **拆包记录主** |

### 包装单位 (packageunit)  (16个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/packageunit/create` | createPackageunit | **物品包装规格信息 ** |
| PUT | `/wms/packageunit/update` | updatePackageunit | **物品包装规格信息 ** |
| DELETE | `/wms/packageunit/delete` | deletePackageunit | **物品包装规格信息 ** |
| GET | `/wms/packageunit/get` | getPackageunit | **编号** |
| GET | `/wms/packageunit/list` | getPackageunitList | **编号** |
| GET | `/wms/packageunit/page` | getPackageunitPage | **编号列表** |
| POST | `/wms/packageunit/senior` | getPackageunitSenior | **物品包装规格信息 ** |
| GET | `/wms/packageunit/pageTree` | getPackageunitPageTree | **物品包装规格信息 ** |
| POST | `/wms/packageunit/seniorTree` | getPackageunitSeniorTree | **物品包装规格信息 ** |
| GET | `/wms/packageunit/pageParent` | getPackageunitPageParent | **物品包装规格信息 ** |
| POST | `/wms/packageunit/seniorParent` | getPackageunitSeniorParent | **物品包装规格信息 ** |
| GET | `/wms/packageunit/export-excel` | exportPackageunitExcel | **物品包装规格信息 ** |
| POST | `/wms/packageunit/export-excel-senior` | exportPackageunitExcel | **物品包装规格信息 ** |
| GET | `/wms/packageunit/get-import-template` | importTemplate | **物品包装规格信息 ** |
| POST | `/wms/packageunit/import` | importExcel | **物品包装规格信息 ** |
| GET | `/wms/packageunit/listByCodes` | getPackageunitListByCodes | **Excel 文件** |

### 参数设置 (paramsetting)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/paramsetting/create` | createParamsetting | **参数设置** |
| PUT | `/wms/paramsetting/update` | updateParamsetting | **参数设置** |
| DELETE | `/wms/paramsetting/delete` | deleteParamsetting | **参数设置** |
| GET | `/wms/paramsetting/get` | getParamsetting | **编号** |
| GET | `/wms/paramsetting/page` | getParamsettingPage | **编号** |
| GET | `/wms/paramsetting/export-excel` | exportParamsettingExcel | **参数设置** |
| POST | `/wms/paramsetting/export-excel-senior` | exportParamsettingExcel | **参数设置** |
| GET | `/wms/paramsetting/get-import-template` | importTemplate | **参数设置** |
| POST | `/wms/paramsetting/import` | importExcel | **参数设置** |
| POST | `/wms/paramsetting/senior` | getParamsettingSenior | **Excel 文件** |

### 拣货任务 (pickJob)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/pick-job-detail/create` | createPickJobDetail | **备货任务子** |
| PUT | `/wms/pick-job-detail/update` | updatePickJobDetail | **备货任务子** |
| DELETE | `/wms/pick-job-detail/delete` | deletePickJobDetail | **备货任务子** |
| GET | `/wms/pick-job-detail/get` | getPickJobDetail | **编号** |
| GET | `/wms/pick-job-detail/list` | getPickJobDetailList | **编号** |
| GET | `/wms/pick-job-detail/page` | getPickJobDetailPage | **编号列表** |
| GET | `/wms/pick-job-detail/export-excel` | exportPickJobDetailExcel | **备货任务子** |
| POST | `/wms/pick-job-detail/senior` | getPickJobDetailSenior | **备货任务子** |
| POST | `/wms/pick-job-main/create` | createPickJobMain | **备货任务主** |
| PUT | `/wms/pick-job-main/update` | updatePickJobMain | **备货任务主** |
| DELETE | `/wms/pick-job-main/delete` | deletePickJobMain | **备货任务主** |
| GET | `/wms/pick-job-main/get` | getPickJobMain | **编号** |
| GET | `/wms/pick-job-main/list` | getPickJobMainList | **编号** |
| GET | `/wms/pick-job-main/page` | getPickJobMainPage | **编号列表** |
| GET | `/wms/pick-job-main/export-excel` | exportPickJobMainExcel | **备货任务主** |
| GET | `/wms/pick-job-main/export-excel-senior` | exportPickJobMainSeniorExcel | **备货任务主** |
| POST | `/wms/pick-job-main/senior` | getPickJobMainSenior | **备货任务主** |
| GET | `/wms/pick-job-main/getPickJobById` | getPickJobById | **备货任务主** |
| POST | `/wms/pick-job-main/getCountByStatus` | getCountByStatus | **编号** |

### 拣货记录 (pickRecord)  (17个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/pick-record-detail/create` | createPickRecordDetail | **备货记录子** |
| PUT | `/wms/pick-record-detail/update` | updatePickRecordDetail | **备货记录子** |
| DELETE | `/wms/pick-record-detail/delete` | deletePickRecordDetail | **备货记录子** |
| GET | `/wms/pick-record-detail/get` | getPickRecordDetail | **编号** |
| GET | `/wms/pick-record-detail/list` | getPickRecordDetailList | **编号** |
| GET | `/wms/pick-record-detail/page` | getPickRecordDetailPage | **编号列表** |
| GET | `/wms/pick-record-detail/export-excel` | exportPickRecordDetailExcel | **备货记录子** |
| POST | `/wms/pick-record-detail/senior` | getPickRecordDetailSenior | **备货记录子** |
| POST | `/wms/pick-record-main/create` | createPickRecordMain | **备货记录主** |
| PUT | `/wms/pick-record-main/update` | updatePickRecordMain | **备货记录主** |
| DELETE | `/wms/pick-record-main/delete` | deletePickRecordMain | **备货记录主** |
| GET | `/wms/pick-record-main/get` | getPickRecordMain | **编号** |
| GET | `/wms/pick-record-main/list` | getPickRecordMainList | **编号** |
| GET | `/wms/pick-record-main/page` | getPickRecordMainPage | **编号列表** |
| GET | `/wms/pick-record-main/export-excel` | exportPickRecordMainExcel | **备货记录主** |
| GET | `/wms/pick-record-main/export-excel-senior` | exportPickRecordMainSeniorExcel | **备货记录主** |
| POST | `/wms/pick-record-main/senior` | getPickRecordMainSenior | **备货记录主** |

### 拣货申请 (pickRequest)  (18个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/pick-request-detail/create` | createPickRequestDetail | **备货申请子** |
| PUT | `/wms/pick-request-detail/update` | updatePickRequestDetail | **备货申请子** |
| DELETE | `/wms/pick-request-detail/delete` | deletePickRequestDetail | **备货申请子** |
| GET | `/wms/pick-request-detail/get` | getPickRequestDetail | **编号** |
| GET | `/wms/pick-request-detail/list` | getPickRequestDetailList | **编号** |
| GET | `/wms/pick-request-detail/page` | getPickRequestDetailPage | **编号列表** |
| GET | `/wms/pick-request-detail/export-excel` | exportPickRequestDetailExcel | **备货申请子** |
| POST | `/wms/pick-request-detail/senior` | getPickRequestDetailSenior | **备货申请子** |
| POST | `/wms/pick-request-main/create` | createPickRequestMain | **备货申请主** |
| PUT | `/wms/pick-request-main/update` | updatePickRequestMain | **备货申请主** |
| DELETE | `/wms/pick-request-main/delete` | deletePickRequestMain | **备货申请主** |
| GET | `/wms/pick-request-main/get` | getPickRequestMain | **编号** |
| GET | `/wms/pick-request-main/list` | getPickRequestMainList | **编号** |
| GET | `/wms/pick-request-main/page` | getPickRequestMainPage | **编号列表** |
| GET | `/wms/pick-request-main/export-excel` | exportPickRequestMainExcel | **备货申请主** |
| GET | `/wms/pick-request-main/export-excel-senior` | exportPickRequestMainSeniorExcel | **备货申请主** |
| POST | `/wms/pick-request-main/senior` | getPickRequestMainSenior | **备货申请主** |
| GET | `/wms/pick-request-main/getPickRequestById` | getPickRequestById | **备货申请主** |

### 计划设置 (plansetting)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/plansetting/create` | createPlansetting | **计划设置** |
| PUT | `/wms/plansetting/update` | updatePlansetting | **计划设置** |
| DELETE | `/wms/plansetting/delete` | deletePlansetting | **计划设置** |
| GET | `/wms/plansetting/get` | getPlansetting | **编号** |
| GET | `/wms/plansetting/list` | getPlansettingList | **编号** |
| GET | `/wms/plansetting/page` | getPlansettingPage | **编号列表** |
| GET | `/wms/plansetting/export-excel` | exportPlansettingExcel | **计划设置** |
| GET | `/wms/plansetting/export-excel-senior` | exportPickRequestMainSeniorExcel | **计划设置** |
| GET | `/wms/plansetting/get-import-template` | importTemplate | **计划设置** |
| POST | `/wms/plansetting/import` | importExcel | **计划设置** |

### 备料发料 (preparetoissue)  (27个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/preparetoissue-detail/create` | createPreparetoissueDetail | **备料计划子** |
| PUT | `/wms/preparetoissue-detail/update` | updatePreparetoissueDetail | **备料计划子** |
| DELETE | `/wms/preparetoissue-detail/delete` | deletePreparetoissueDetail | **备料计划子** |
| GET | `/wms/preparetoissue-detail/get` | getPreparetoissueDetail | **编号** |
| GET | `/wms/preparetoissue-detail/list` | getPreparetoissueDetailList | **编号** |
| GET | `/wms/preparetoissue-detail/page` | getPreparetoissueDetailPage | **编号列表** |
| POST | `/wms/preparetoissue-detail/senior` | getPreparetoissueDetailSenior | **备料计划子** |
| POST | `/wms/preparetoissue-main/create` | createPreparetoissueMain | **备料计划主** |
| PUT | `/wms/preparetoissue-main/update` | updatePreparetoissueMain | **备料计划主** |
| DELETE | `/wms/preparetoissue-main/delete` | deletePreparetoissueMain | **备料计划主** |
| GET | `/wms/preparetoissue-main/get` | getPreparetoissueMain | **编号** |
| GET | `/wms/preparetoissue-main/list` | getPreparetoissueMainList | **编号** |
| GET | `/wms/preparetoissue-main/page` | getPreparetoissueMainPage | **编号列表** |
| POST | `/wms/preparetoissue-main/senior` | getPreparetoissueMainSenior | **备料计划主** |
| POST | `/wms/preparetoissue-main/import` | importExcel | **备料计划主** |
| GET | `/wms/preparetoissue-main/get-import-template` | importTemplate | **Excel 文件** |
| GET | `/wms/preparetoissue-main/export-excel` | exportPreparetoissueMainExcel | **备料计划主** |
| POST | `/wms/preparetoissue-main/export-excel-senior` | exportProductionMainSeniorExcel | **备料计划主** |
| GET | `/wms/preparetoissue-main/getBomDisassemble` | getBomDisassemble | **备料计划主** |
| PUT | `/wms/preparetoissue-main/close` | closePreparetoissueMain | **编号** |
| PUT | `/wms/preparetoissue-main/submit` | submitPreparetoissueMain | **编号** |
| PUT | `/wms/preparetoissue-main/open` | openPreparetoissueMain | **编号** |
| PUT | `/wms/preparetoissue-main/reject` | rejectPreparetoissueMain | **编号** |
| PUT | `/wms/preparetoissue-main/agree` | agreePreparetoissueMain | **编号** |
| PUT | `/wms/preparetoissue-main/publish` | publishPreparetoissueMain | **编号** |
| PUT | `/wms/preparetoissue-main/resetting` | resettingPreparetoissueMain | **编号** |
| POST | `/wms/preparetoissue-main/generateIssueRequest` | generateIssueRequest | **编号** |

### 打印 (print)  (3个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/print/printerList` | getPrinterList | **PDA根据用户权限操作所有任务模块** |
| GET | `/wms/print/modelList` | getModelList | **PDA根据用户权限操作所有任务模块** |
| POST | `/wms/print/printModel` | postPrintModel | **PDA根据用户权限操作所有任务模块** |

### 打印模板 (printbusinesstypetemplate)  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/print-businesstype-template/create` | createPrintBusinesstypeTemplate | **打印服务业务类型和打印模板关系** |
| PUT | `/wms/print-businesstype-template/update` | updatePrintBusinesstypeTemplate | **打印服务业务类型和打印模板关系** |
| DELETE | `/wms/print-businesstype-template/delete` | deletePrintBusinesstypeTemplate | **打印服务业务类型和打印模板关系** |
| GET | `/wms/print-businesstype-template/get` | getPrintBusinesstypeTemplate | **编号** |
| GET | `/wms/print-businesstype-template/page` | getPrintBusinesstypeTemplatePage | **编号** |
| POST | `/wms/print-businesstype-template/senior` | getPrintBusinesstypeTemplateSenior | **打印服务业务类型和打印模板关系** |
| GET | `/wms/print-businesstype-template/export-excel` | exportPrintBusinesstypeTemplateExcel | **打印服务业务类型和打印模板关系** |
| GET | `/wms/print-businesstype-template/get-import-template` | importTemplate | **打印服务业务类型和打印模板关系** |
| POST | `/wms/print-businesstype-template/import` | importExcel | **打印服务业务类型和打印模板关系** |

### 客户端打印机 (printclientprinter)  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/print-client-printer/create` | createPrintClientPrinter | **打印服务客户端与打印机名称关系** |
| PUT | `/wms/print-client-printer/update` | updatePrintClientPrinter | **打印服务客户端与打印机名称关系** |
| DELETE | `/wms/print-client-printer/delete` | deletePrintClientPrinter | **打印服务客户端与打印机名称关系** |
| GET | `/wms/print-client-printer/get` | getPrintClientPrinter | **编号** |
| GET | `/wms/print-client-printer/page` | getPrintClientPrinterPage | **编号** |
| POST | `/wms/print-client-printer/senior` | getPrintClientPrinterSenior | **打印服务客户端与打印机名称关系** |
| GET | `/wms/print-client-printer/export-excel` | exportPrintClientPrinterExcel | **打印服务客户端与打印机名称关系** |
| GET | `/wms/print-client-printer/get-import-template` | importTemplate | **打印服务客户端与打印机名称关系** |
| POST | `/wms/print-client-printer/import` | importExcel | **打印服务客户端与打印机名称关系** |

### 打印关联 (printcorrelation)  (6个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/print-correlation/create` | createPrintCorrelation | **打印标签关联** |
| PUT | `/wms/print-correlation/update` | updatePrintCorrelation | **打印标签关联** |
| DELETE | `/wms/print-correlation/delete` | deletePrintCorrelation | **打印标签关联** |
| GET | `/wms/print-correlation/get` | getPrintCorrelation | **编号** |
| GET | `/wms/print-correlation/page` | getPrintCorrelationPage | **编号** |
| POST | `/wms/print-correlation/senior` | getPrintCorrelationSenior | **打印标签关联** |

### 工艺 (process)  (13个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/process/create` | createProcess | **工序** |
| PUT | `/wms/process/update` | updateProcess | **工序** |
| DELETE | `/wms/process/delete` | deleteProcess | **工序** |
| GET | `/wms/process/get` | getProcess | **编号** |
| GET | `/wms/process/getByCode` | getProcessByCode | **编号** |
| GET | `/wms/process/page` | getProcessPage | **编码** |
| GET | `/wms/process/export-excel` | exportProcessExcel | **工序** |
| POST | `/wms/process/export-excel-senior` | exportProcessExcel | **工序** |
| GET | `/wms/process/get-import-template` | importTemplate | **工序** |
| POST | `/wms/process/import` | importExcel | **工序** |
| POST | `/wms/process/senior` | getProcessSenior | **Excel 文件** |
| GET | `/wms/process/noPage` | getProcessNoPage | **工序** |
| GET | `/wms/process/listByCodes` | getProcessByCodes | **工序** |

### 工序生产记录 (processproductionRecord)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/processproduction-record-detail/create` | createProcessproductionRecordDetail | **工序报产记录子** |
| PUT | `/wms/processproduction-record-detail/update` | updateProcessproductionRecordDetail | **工序报产记录子** |
| DELETE | `/wms/processproduction-record-detail/delete` | deleteProcessproductionRecordDetail | **工序报产记录子** |
| GET | `/wms/processproduction-record-detail/get` | getProcessproductionRecordDetail | **编号** |
| GET | `/wms/processproduction-record-detail/page` | getProcessproductionRecordDetailPage | **编号** |
| POST | `/wms/processproduction-record-detail/senior` | getProcessproductionRecordDetailSenior | **工序报产记录子** |
| GET | `/wms/processproduction-record-detail/export-excel` | exportProcessproductionRecordDetailExcel | **工序报产记录子** |
| GET | `/wms/processproduction-record-detail/get-import-template` | importTemplate | **工序报产记录子** |
| POST | `/wms/processproduction-record-detail/import` | importExcel | **工序报产记录子** |
| POST | `/wms/processproduction-record-main/create` | createProcessproductionRecordMain | **工序报产记录主** |
| PUT | `/wms/processproduction-record-main/update` | updateProcessproductionRecordMain | **工序报产记录主** |
| DELETE | `/wms/processproduction-record-main/delete` | deleteProcessproductionRecordMain | **工序报产记录主** |
| GET | `/wms/processproduction-record-main/get` | getProcessproductionRecordMain | **编号** |
| GET | `/wms/processproduction-record-main/page` | getProcessproductionRecordMainPage | **编号** |
| POST | `/wms/processproduction-record-main/senior` | getProcessproductionRecordMainSenior | **工序报产记录主** |
| GET | `/wms/processproduction-record-main/export-excel` | exportProcessproductionRecordMainExcel | **工序报产记录主** |
| POST | `/wms/processproduction-record-main/export-excel-senior` | exportProcessproductionRecordMainSeniorExcel | **工序报产记录主** |
| GET | `/wms/processproduction-record-main/get-import-template` | importTemplate | **工序报产记录主** |
| POST | `/wms/processproduction-record-main/import` | importExcel | **工序报产记录主** |

### 工序生产申请 (processproductionRequest)  (27个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/processproduction-request-detail/create` | createProcessproductionRequestDetail | **工序报产申请子** |
| PUT | `/wms/processproduction-request-detail/update` | updateProcessproductionRequestDetail | **工序报产申请子** |
| DELETE | `/wms/processproduction-request-detail/delete` | deleteProcessproductionRequestDetail | **工序报产申请子** |
| GET | `/wms/processproduction-request-detail/get` | getProcessproductionRequestDetail | **编号** |
| GET | `/wms/processproduction-request-detail/page` | getProcessproductionRequestDetailPage | **编号** |
| POST | `/wms/processproduction-request-detail/senior` | getProcessproductionRequestDetailSenior | **工序报产申请子** |
| GET | `/wms/processproduction-request-detail/export-excel` | exportProcessproductionRequestDetailExcel | **工序报产申请子** |
| GET | `/wms/processproduction-request-detail/get-import-template` | importTemplate | **工序报产申请子** |
| POST | `/wms/processproduction-request-detail/import` | importExcel | **工序报产申请子** |
| POST | `/wms/processproduction-request-main/create` | createProcessproductionRequestMain | **工序报产申请主** |
| PUT | `/wms/processproduction-request-main/update` | updateProcessproductionRequestMain | **工序报产申请主** |
| DELETE | `/wms/processproduction-request-main/delete` | deleteProcessproductionRequestMain | **工序报产申请主** |
| GET | `/wms/processproduction-request-main/get` | getProcessproductionRequestMain | **编号** |
| GET | `/wms/processproduction-request-main/page` | getProcessproductionRequestMainPage | **编号** |
| POST | `/wms/processproduction-request-main/senior` | getProcessproductionRequestMainSenior | **工序报产申请主** |
| GET | `/wms/processproduction-request-main/export-excel` | exportProcessproductionRequestMainExcel | **工序报产申请主** |
| POST | `/wms/processproduction-request-main/export-excel-senior` | exportProcessproductionRequestMainSeniorExcel | **工序报产申请主** |
| GET | `/wms/processproduction-request-main/get-import-template` | importTemplate | **工序报产申请主** |
| POST | `/wms/processproduction-request-main/import` | importExcel | **工序报产申请主** |
| PUT | `/wms/processproduction-request-main/close` | closeProcessproductionRequestMain | **Excel 文件** |
| PUT | `/wms/processproduction-request-main/reAdd` | reAddProcessproductionRequestMain | **编号** |
| PUT | `/wms/processproduction-request-main/submit` | submitProcessproductionRequestMain | **编号** |
| PUT | `/wms/processproduction-request-main/refused` | abortProcessproductionRequestMain | **编号** |
| PUT | `/wms/processproduction-request-main/agree` | agreeProcessproductionRequestMain | **编号** |
| PUT | `/wms/processproduction-request-main/handle` | handleProcessproductionRequestMain | **编号** |
| GET | `/wms/processproduction-request-main/queryChildItemByParentCodePage` | queryChildItemByParentCodePage | **编号** |
| POST | `/wms/processproduction-request-main/queryChildItemByParentCodeSenior` | queryChildItemByParentCodeSenior | **工序报产申请主** |

### 产品拆解任务 (productdismantleJob)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productdismantle-job-detail/create` | createProductdismantleJobDetail | **制品拆解任务子** |
| PUT | `/wms/productdismantle-job-detail/update` | updateProductdismantleJobDetail | **制品拆解任务子** |
| DELETE | `/wms/productdismantle-job-detail/delete` | deleteProductdismantleJobDetail | **制品拆解任务子** |
| GET | `/wms/productdismantle-job-detail/get` | getProductdismantleJobDetail | **编号** |
| GET | `/wms/productdismantle-job-detail/list` | getProductdismantleJobDetailList | **编号** |
| GET | `/wms/productdismantle-job-detail/page` | getProductdismantleJobDetailPage | **编号列表** |
| GET | `/wms/productdismantle-job-detail/export-excel` | exportProductdismantleJobDetailExcel | **制品拆解任务子** |
| POST | `/wms/productdismantle-job-detail/senior` | getProductdismantleJobDetailSenior | **制品拆解任务子** |
| POST | `/wms/productdismantle-job-main/create` | createProductdismantleJobMain | **制品拆解任务主** |
| PUT | `/wms/productdismantle-job-main/update` | updateProductdismantleJobMain | **制品拆解任务主** |
| DELETE | `/wms/productdismantle-job-main/delete` | deleteProductdismantleJobMain | **制品拆解任务主** |
| GET | `/wms/productdismantle-job-main/get` | getProductdismantleJobMain | **编号** |
| GET | `/wms/productdismantle-job-main/list` | getProductdismantleJobMainList | **编号** |
| GET | `/wms/productdismantle-job-main/page` | getProductdismantleJobMainPage | **编号列表** |
| POST | `/wms/productdismantle-job-main/senior` | getProductdismantleJobMainSenior | **制品拆解任务主** |
| GET | `/wms/productdismantle-job-main/export-excel` | exportProductdismantleJobMainExcel | **制品拆解任务主** |
| POST | `/wms/productdismantle-job-main/export-excel-senior` | exportProductdismantleJobMainExcel | **制品拆解任务主** |
| GET | `/wms/productdismantle-job-main/getProductdismantleJobById` | getProductdismantleJobById | **制品拆解任务主** |
| POST | `/wms/productdismantle-job-main/getCountByStatus` | getCountByStatus | **编号** |

### 产品拆解记录 (productdismantleRecord)  (25个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/dismantle-record-detailb/create` | createDismantleRecordDetailb | **制品拆解记录子** |
| PUT | `/wms/dismantle-record-detailb/update` | updateDismantleRecordDetailb | **制品拆解记录子** |
| DELETE | `/wms/dismantle-record-detailb/delete` | deleteDismantleRecordDetailb | **制品拆解记录子** |
| GET | `/wms/dismantle-record-detailb/get` | getDismantleRecordDetailb | **编号** |
| GET | `/wms/dismantle-record-detailb/list` | getDismantleRecordDetailbList | **编号** |
| GET | `/wms/dismantle-record-detailb/page` | getDismantleRecordDetailbPage | **编号列表** |
| POST | `/wms/dismantle-record-detailb/senior` | getDismantleRecordDetailbSenior | **制品拆解记录子** |
| GET | `/wms/dismantle-record-detailb/export-excel` | exportDismantleRecordDetailbExcel | **制品拆解记录子** |
| POST | `/wms/productdismantle-record-detaila/create` | createProductdismantleRecordDetaila | **制品拆解记录子** |
| PUT | `/wms/productdismantle-record-detaila/update` | updateProductdismantleRecordDetaila | **制品拆解记录子** |
| DELETE | `/wms/productdismantle-record-detaila/delete` | deleteProductdismantleRecordDetaila | **制品拆解记录子** |
| GET | `/wms/productdismantle-record-detaila/get` | getProductdismantleRecordDetaila | **编号** |
| GET | `/wms/productdismantle-record-detaila/list` | getProductdismantleRecordDetailaList | **编号** |
| GET | `/wms/productdismantle-record-detaila/page` | getProductdismantleRecordDetailaPage | **编号列表** |
| GET | `/wms/productdismantle-record-detaila/export-excel` | exportProductdismantleRecordDetailaExcel | **制品拆解记录子** |
| POST | `/wms/productdismantle-record-detaila/senior` | getProductdismantleRecordDetailaSenior | **制品拆解记录子** |
| POST | `/wms/productdismantle-record-main/create` | createProductdismantleRecordMain | **制品拆解记录主** |
| PUT | `/wms/productdismantle-record-main/update` | updateProductdismantleRecordMain | **制品拆解记录主** |
| DELETE | `/wms/productdismantle-record-main/delete` | deleteProductdismantleRecordMain | **制品拆解记录主** |
| GET | `/wms/productdismantle-record-main/get` | getProductdismantleRecordMain | **编号** |
| GET | `/wms/productdismantle-record-main/list` | getProductdismantleRecordMainList | **编号** |
| GET | `/wms/productdismantle-record-main/page` | getProductdismantleRecordMainPage | **编号列表** |
| POST | `/wms/productdismantle-record-main/senior` | getProductdismantleRecordMainSenior | **制品拆解记录主** |
| GET | `/wms/productdismantle-record-main/export-excel` | exportProductdismantleRecordMainExcel | **制品拆解记录主** |
| POST | `/wms/productdismantle-record-main/export-excel-senior` | exportProductdismantleRecordMainSeniorExcel | **制品拆解记录主** |

### 产品拆解申请 (productdismantleRequest)  (32个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productdismantle-request-detaila/create` | createProductdismantleRequestDetaila | **制品拆解申请子** |
| PUT | `/wms/productdismantle-request-detaila/update` | updateProductdismantleRequestDetaila | **制品拆解申请子** |
| DELETE | `/wms/productdismantle-request-detaila/delete` | deleteProductdismantleRequestDetaila | **制品拆解申请子** |
| GET | `/wms/productdismantle-request-detaila/get` | getProductdismantleRequestDetaila | **编号** |
| GET | `/wms/productdismantle-request-detaila/list` | getProductdismantleRequestDetailaList | **编号** |
| GET | `/wms/productdismantle-request-detaila/page` | getProductdismantleRequestDetailaPage | **编号列表** |
| GET | `/wms/productdismantle-request-detaila/export-excel` | exportProductdismantleRequestDetailaExcel | **制品拆解申请子** |
| POST | `/wms/productdismantle-request-detaila/senior` | getProductdismantleRequestDetailaSenior | **制品拆解申请子** |
| POST | `/wms/dismantle-request-detailb/create` | createDismantleRequestDetailb | **制品拆解申请子** |
| POST | `/wms/dismantle-request-detailb/update` | updateDismantleRequestDetailb | **制品拆解申请子** |
| DELETE | `/wms/dismantle-request-detailb/delete` | deleteDismantleRequestDetailb | **制品拆解申请子** |
| GET | `/wms/dismantle-request-detailb/get` | getDismantleRequestDetailb | **编号** |
| GET | `/wms/dismantle-request-detailb/list` | getDismantleRequestDetailbList | **编号** |
| GET | `/wms/dismantle-request-detailb/page` | getDismantleRequestDetailbPage | **编号列表** |
| GET | `/wms/dismantle-request-detailb/export-excel` | exportDismantleRequestDetailbExcel | **制品拆解申请子** |
| POST | `/wms/dismantle-request-detailb/senior` | getProductdismantleRequestDetailbSenior | **制品拆解申请子** |
| GET | `/wms/dismantle-request-detailb/bomPage` | getBomInfoPage | **制品拆解申请子** |
| POST | `/wms/productdismantle-request-main/create` | createProductdismantleRequestMain | **制品拆解申请主** |
| PUT | `/wms/productdismantle-request-main/update` | updateProductdismantleRequestMain | **制品拆解申请主** |
| DELETE | `/wms/productdismantle-request-main/delete` | deleteProductdismantleRequestMain | **制品拆解申请主** |
| GET | `/wms/productdismantle-request-main/get` | getProductdismantleRequestMain | **编号** |
| GET | `/wms/productdismantle-request-main/list` | getProductdismantleRequestMainList | **编号** |
| GET | `/wms/productdismantle-request-main/page` | getProductdismantleRequestMainPage | **编号列表** |
| POST | `/wms/productdismantle-request-main/senior` | getProductdismantleRequestMainSenior | **制品拆解申请主** |
| GET | `/wms/productdismantle-request-main/export-excel` | exportProductdismantleRequestMainExcel | **制品拆解申请主** |
| POST | `/wms/productdismantle-request-main/export-excel-senior` | exportProductdismantleRequestMainSeniorExcel | **制品拆解申请主** |
| PUT | `/wms/productdismantle-request-main/close` | closeProductdismantleRequestMain | **制品拆解申请主** |
| PUT | `/wms/productdismantle-request-main/reAdd` | reAddProductdismantleRequestMain | **编号** |
| PUT | `/wms/productdismantle-request-main/submit` | submitProductdismantleRequestMain | **编号** |
| PUT | `/wms/productdismantle-request-main/refused` | refusedProductdismantleRequestMain | **编号** |
| PUT | `/wms/productdismantle-request-main/agree` | agreeProductdismantleRequestMain | **编号** |
| PUT | `/wms/productdismantle-request-main/handle` | handleProductdismantleRequestMain | **编号** |

### 生产 (production)  (31个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/production-detail/create` | createProductionDetail | **生产计划子** |
| PUT | `/wms/production-detail/update` | updateProductionDetail | **生产计划子** |
| DELETE | `/wms/production-detail/delete` | deleteProductionDetail | **生产计划子** |
| GET | `/wms/production-detail/get` | getProductionDetail | **编号** |
| GET | `/wms/production-detail/list` | getProductionDetailList | **编号** |
| GET | `/wms/production-detail/page` | getProductionDetailPage | **编号列表** |
| POST | `/wms/production-detail/senior` | getProductionDetailSenior | **生产计划子** |
| POST | `/wms/production-main/create` | createProductionMain | **生产计划主** |
| PUT | `/wms/production-main/update` | updateProductionMain | **生产计划主** |
| DELETE | `/wms/production-main/delete` | deleteProductionMain | **生产计划主** |
| GET | `/wms/production-main/get` | getProductionMain | **编号** |
| GET | `/wms/production-main/list` | getProductionMainList | **编号** |
| GET | `/wms/production-main/getDetailByAvailable` | getDetailByAvailable | **编号列表** |
| GET | `/wms/production-main/page` | getProductionMainPage | **编号** |
| POST | `/wms/production-main/senior` | getProductionMainSenior | **生产计划主** |
| GET | `/wms/production-main/export-excel` | exportProductionMainExcel | **生产计划主** |
| POST | `/wms/production-main/export-excel-senior` | exportProductionMainSeniorExcel | **生产计划主** |
| GET | `/wms/production-main/get-import-template` | importTemplate | **生产计划主** |
| POST | `/wms/production-main/import` | importExcel | **生产计划主** |
| PUT | `/wms/production-main/close` | closeProductionMain | **Excel 文件** |
| PUT | `/wms/production-main/submit` | submitProductionMain | **编号** |
| PUT | `/wms/production-main/open` | openProductionMain | **编号** |
| PUT | `/wms/production-main/reject` | rejectProductionMain | **编号** |
| PUT | `/wms/production-main/agree` | agreeProductionMain | **编号** |
| PUT | `/wms/production-main/publish` | publishProductionMain | **编号** |
| PUT | `/wms/production-main/resetting` | resettingProductionMain | **编号** |
| POST | `/wms/production-main/generateProductreceiptRequest` | generateProductreceiptRequest | **编号** |
| POST | `/wms/production-main/generatePreparetoissue` | generatePreparetoissue | **生产计划主** |
| GET | `/wms/production-main/getPlanProductionByProductionLineAndPlanDate` | getPlanProductionByProductionLineAndPlanDate | **生产计划主** |
| GET | `/wms/production-main/getProductionlineAndWorkStation` | getProductionlineAndWorkStation | **生产计划主** |
| GET | `/wms/production-main/getPackageByItemCode` | getPackageByItemCode | **生产计划主** |

### 生产物料编码 (productionitemcodespareitemcode)  (12个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productionitemcode-spareitemcode/create` | createProductionitemcodeSpareitemcode | **量产件备件关系** |
| PUT | `/wms/productionitemcode-spareitemcode/update` | updateProductionitemcodeSpareitemcode | **量产件备件关系** |
| DELETE | `/wms/productionitemcode-spareitemcode/delete` | deleteProductionitemcodeSpareitemcode | **量产件备件关系** |
| GET | `/wms/productionitemcode-spareitemcode/get` | getProductionitemcodeSpareitemcode | **编号** |
| GET | `/wms/productionitemcode-spareitemcode/getRelation` | getProductionitemcodeSpareitemcodeRelation | **编号** |
| POST | `/wms/productionitemcode-spareitemcode/getRelationSenior` | getProductionitemcodeSpareitemcodeRelationSenior | **变更前物料代码** |
| GET | `/wms/productionitemcode-spareitemcode/page` | getProductionitemcodeSpareitemcodePage | **量产件备件关系** |
| POST | `/wms/productionitemcode-spareitemcode/senior` | getProductionitemcodeSpareitemcodeSenior | **量产件备件关系** |
| GET | `/wms/productionitemcode-spareitemcode/export-excel` | exportProductionitemcodeSpareitemcodeExcel | **量产件备件关系** |
| POST | `/wms/productionitemcode-spareitemcode/export-excel-senior` | exportItembasicExcel | **量产件备件关系** |
| GET | `/wms/productionitemcode-spareitemcode/get-import-template` | importTemplate | **量产件备件关系** |
| POST | `/wms/productionitemcode-spareitemcode/import` | importExcel | **量产件备件关系** |

### 产线 (productionline)  (12个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productionline/create` | createProductionline | **生产线** |
| PUT | `/wms/productionline/update` | updateProductionline | **生产线** |
| DELETE | `/wms/productionline/delete` | deleteProductionline | **生产线** |
| GET | `/wms/productionline/get` | getProductionline | **编号** |
| GET | `/wms/productionline/page` | getProductionlinePage | **编号** |
| GET | `/wms/productionline/export-excel` | exportProductionlineExcel | **生产线** |
| POST | `/wms/productionline/export-excel-senior` | exportProductionlineExcel | **生产线** |
| GET | `/wms/productionline/get-import-template` | importTemplate | **生产线** |
| POST | `/wms/productionline/import` | importExcel | **生产线** |
| POST | `/wms/productionline/senior` | getProductionlineSenior | **Excel 文件** |
| GET | `/wms/productionline/noPage` | getProductionlineNoPage | **生产线** |
| GET | `/wms/productionline/listByCodes` | getProductionlineByCodes | **生产线** |

### 产线项目 (productionlineitem)  (17个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productionlineitem/create` | createProductionlineitem | **生产线物料关系** |
| PUT | `/wms/productionlineitem/update` | updateProductionlineitem | **生产线物料关系** |
| DELETE | `/wms/productionlineitem/delete` | deleteProductionlineitem | **生产线物料关系** |
| GET | `/wms/productionlineitem/get` | getProductionlineitem | **编号** |
| GET | `/wms/productionlineitem/selectItemCodeToProductionLineCode` | selectItemCodeToProductionLineCode | **编号** |
| GET | `/wms/productionlineitem/page` | getProductionlineitemPage | **编号** |
| POST | `/wms/productionlineitem/seniorBom` | getProductionlineitemSeniorBomList | **生产线物料关系** |
| GET | `/wms/productionlineitem/pageBom` | getProductionlineitemPageBomList | **生产线物料关系** |
| GET | `/wms/productionlineitem/export-excel` | exportProductionlineitemExcel | **生产线物料关系** |
| POST | `/wms/productionlineitem/export-excel-senior` | exportProductionlineitemExcel | **生产线物料关系** |
| GET | `/wms/productionlineitem/get-import-template` | importTemplate | **生产线物料关系** |
| POST | `/wms/productionlineitem/import` | importExcel | **生产线物料关系** |
| POST | `/wms/productionlineitem/senior` | getProductionlineitemSenior | **Excel 文件** |
| GET | `/wms/productionlineitem/pageByItemtype` | getProductionlineitemPageByItemtype | **生产线物料关系** |
| POST | `/wms/productionlineitem/pageByItemtypeSenior` | getProductionlineitemPageByItemtypeSenior | **生产线物料关系** |
| GET | `/wms/productionlineitem/listByCodes` | getProductionlineitemListByCodes | **生产线物料关系** |
| GET | `/wms/productionlineitem/getBomVersionByProductionline` | getBomVersionByProductionline | **生产线物料关系** |

### 生产入库任务 (productionreceiptJob)  (25个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productionreceipt-job-detail/create` | createProductionreceiptJobDetail | **生产收料任务子** |
| PUT | `/wms/productionreceipt-job-detail/update` | updateProductionreceiptJobDetail | **生产收料任务子** |
| DELETE | `/wms/productionreceipt-job-detail/delete` | deleteProductionreceiptJobDetail | **生产收料任务子** |
| GET | `/wms/productionreceipt-job-detail/get` | getProductionreceiptJobDetail | **编号** |
| GET | `/wms/productionreceipt-job-detail/list` | getProductionreceiptJobDetailList | **编号** |
| GET | `/wms/productionreceipt-job-detail/page` | getProductionreceiptJobDetailPage | **编号列表** |
| GET | `/wms/productionreceipt-job-detail/export-excel` | exportProductionreceiptJobDetailExcel | **生产收料任务子** |
| POST | `/wms/productionreceipt-job-detail/senior` | getProductionreceiptJobDetailSenior | **生产收料任务子** |
| POST | `/wms/productionreceipt-job-main/create` | createProductionreceiptJobMain | **生产收料任务主** |
| PUT | `/wms/productionreceipt-job-main/update` | updateProductionreceiptJobMain | **生产收料任务主** |
| DELETE | `/wms/productionreceipt-job-main/delete` | deleteProductionreceiptJobMain | **生产收料任务主** |
| GET | `/wms/productionreceipt-job-main/get` | getProductionreceiptJobMain | **编号** |
| GET | `/wms/productionreceipt-job-main/list` | getProductionreceiptJobMainList | **编号** |
| GET | `/wms/productionreceipt-job-main/page` | getProductionreceiptJobMainPage | **编号列表** |
| POST | `/wms/productionreceipt-job-main/senior` | getProductionreceiptJobMainSenior | **生产收料任务主** |
| GET | `/wms/productionreceipt-job-main/export-excel` | exportProductionreceiptJobMainExcel | **生产收料任务主** |
| POST | `/wms/productionreceipt-job-main/export-excel-senior` | exportProductionreceiptJobMainSeniorExcel | **生产收料任务主** |
| GET | `/wms/productionreceipt-job-main/getProductionreceiptJobById` | getProductionreceiptJobById | **生产收料任务主** |
| POST | `/wms/productionreceipt-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/productionreceipt-job-main/accept` | acceptProductionreceiptJobMain | **类型数组** |
| PUT | `/wms/productionreceipt-job-main/abandon` | abandonProductionreceiptJobMain | **生产收料任务主** |
| PUT | `/wms/productionreceipt-job-main/close` | closeProductionreceiptJobMain | **生产收料任务主** |
| PUT | `/wms/productionreceipt-job-main/execute` | executeProductionreceiptJobMain | **生产收料任务主** |
| GET | `/wms/productionreceipt-job-main/getProductionreceiptJobByProductionline` | getProductionreceiptJobByProductionline | **生产收料任务主** |
| PUT | `/wms/productionreceipt-job-main/updateProductionreceiptJobConfig` | updateProductionreceiptJobConfig | **生产收料任务主** |

### 生产入库记录 (productionreceiptRecord)  (14个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productionreceipt-record-detail/create` | createProductionreceiptRecordDetail | **生产收料记录子** |
| PUT | `/wms/productionreceipt-record-detail/update` | updateProductionreceiptRecordDetail | **生产收料记录子** |
| DELETE | `/wms/productionreceipt-record-detail/delete` | deleteProductionreceiptRecordDetail | **生产收料记录子** |
| GET | `/wms/productionreceipt-record-detail/get` | getProductionreceiptRecordDetail | **编号** |
| GET | `/wms/productionreceipt-record-detail/list` | getProductionreceiptRecordDetailList | **编号** |
| GET | `/wms/productionreceipt-record-detail/page` | getProductionreceiptRecordDetailPage | **编号列表** |
| GET | `/wms/productionreceipt-record-detail/export-excel` | exportProductionreceiptRecordDetailExcel | **生产收料记录子** |
| POST | `/wms/productionreceipt-record-detail/senior` | getProductionreceiptRecordDetailSenior | **生产收料记录子** |
| GET | `/wms/productionreceipt-record-main/get` | getProductionreceiptRecordMain | **生产收料记录主** |
| GET | `/wms/productionreceipt-record-main/list` | getProductionreceiptRecordMainList | **编号** |
| GET | `/wms/productionreceipt-record-main/page` | getProductionreceiptRecordMainPage | **编号列表** |
| POST | `/wms/productionreceipt-record-main/senior` | getProductionreceiptRecordMainSenior | **生产收料记录主** |
| GET | `/wms/productionreceipt-record-main/export-excel` | exportProductionreceiptRecordMainExcel | **生产收料记录主** |
| POST | `/wms/productionreceipt-record-main/export-excel-senior` | exportProductionreceiptRecordMainSeniorExcel | **生产收料记录主** |

### 生产退料任务 (productionreturnJob)  (31个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productionreturn-job-detail/create` | createProductionreturnJobDetail | **生产退料任务子** |
| PUT | `/wms/productionreturn-job-detail/update` | updateProductionreturnJobDetail | **生产退料任务子** |
| DELETE | `/wms/productionreturn-job-detail/delete` | deleteProductionreturnJobDetail | **生产退料任务子** |
| GET | `/wms/productionreturn-job-detail/get` | getProductionreturnJobDetail | **编号** |
| GET | `/wms/productionreturn-job-detail/list` | getProductionreturnJobDetailList | **编号** |
| GET | `/wms/productionreturn-job-detail/page` | getProductionreturnJobDetailPage | **编号列表** |
| GET | `/wms/productionreturn-job-detail/export-excel` | exportProductionreturnJobDetailExcel | **生产退料任务子** |
| POST | `/wms/productionreturn-job-detail/senior` | getProductionreturnJobDetailSenior | **生产退料任务子** |
| GET | `/wms/productionreturn-job-detail-hold/page` | getProductionreturnJobDetailPage | **生产退料任务子** |
| GET | `/wms/productionreturn-job-detail-hold/export-excel` | exportProductionreturnJobDetailExcel | **生产退料任务子** |
| POST | `/wms/productionreturn-job-detail-hold/senior` | getProductionreturnJobDetailSenior | **生产退料任务子** |
| GET | `/wms/productionreturn-job-detail-store/page` | getProductionreturnJobDetailPage | **生产退料任务子** |
| GET | `/wms/productionreturn-job-detail-store/export-excel` | exportProductionreturnJobDetailExcel | **生产退料任务子** |
| POST | `/wms/productionreturn-job-detail-store/senior` | getProductionreturnJobDetailSenior | **生产退料任务子** |
| POST | `/wms/productionreturn-job-main/create` | createProductionreturnJobMain | **生产退料任务主** |
| PUT | `/wms/productionreturn-job-main/update` | updateProductionreturnJobMain | **生产退料任务主** |
| DELETE | `/wms/productionreturn-job-main/delete` | deleteProductionreturnJobMain | **生产退料任务主** |
| GET | `/wms/productionreturn-job-main/get` | getProductionreturnJobMain | **编号** |
| GET | `/wms/productionreturn-job-main/list` | getProductionreturnJobMainList | **编号** |
| GET | `/wms/productionreturn-job-main/page` | getProductionreturnJobMainPage | **编号列表** |
| POST | `/wms/productionreturn-job-main/senior` | getProductionreturnJobMainSenior | **生产退料任务主** |
| GET | `/wms/productionreturn-job-main/export-excel` | exportProductionreturnJobMainExcel | **生产退料任务主** |
| POST | `/wms/productionreturn-job-main/export-excel-senior` | exportProductionreturnJobMainSeniorExcel | **生产退料任务主** |
| GET | `/wms/productionreturn-job-main/getProductionreturnJobById` | getProductionreturnJobById | **生产退料任务主** |
| POST | `/wms/productionreturn-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/productionreturn-job-main/accept` | acceptProductionreturnJobMain | **类型数组** |
| PUT | `/wms/productionreturn-job-main/abandon` | abandonProductionreturnJobMain | **生产退料任务主** |
| PUT | `/wms/productionreturn-job-main/close` | closeProductionreturnJobMain | **生产退料任务主** |
| PUT | `/wms/productionreturn-job-main/execute` | executeProductionreturnJobMain | **生产退料任务主** |
| PUT | `/wms/productionreturn-job-main/updateProductionreturnJobConfig` | updateProductionreturnJobConfig | **生产退料任务主** |
| PUT | `/wms/productionreturn-job-main/updateProductionreturnJobHoldConfig` | updateProductionreturnJobHoldConfig | **生产退料任务主** |

### 生产退料记录 (productionreturnRecord)  (27个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productionreturn-record-detail/create` | createProductionreturnRecordDetail | **生产退料记录子** |
| PUT | `/wms/productionreturn-record-detail/update` | updateProductionreturnRecordDetail | **生产退料记录子** |
| DELETE | `/wms/productionreturn-record-detail/delete` | deleteProductionreturnRecordDetail | **生产退料记录子** |
| GET | `/wms/productionreturn-record-detail/get` | getProductionreturnRecordDetail | **编号** |
| GET | `/wms/productionreturn-record-detail/list` | getProductionreturnRecordDetailList | **编号** |
| GET | `/wms/productionreturn-record-detail/page` | getProductionreturnRecordDetailPage | **编号列表** |
| GET | `/wms/productionreturn-record-detail/export-excel` | exportProductionreturnRecordDetailExcel | **生产退料记录子** |
| POST | `/wms/productionreturn-record-detail/senior` | getProductionreturnRecordDetailSenior | **生产退料记录子** |
| GET | `/wms/productionreturn-record-detail-hold/page` | getProductionreturnRecordDetailPage | **隔离退料记录子** |
| POST | `/wms/productionreturn-record-detail-hold/senior` | getProductionreturnRecordDetailSenior | **隔离退料记录子** |
| GET | `/wms/productionreturn-record-detail-hold/export-excel` | exportProductionreturnRecordDetailExcel | **隔离退料记录子** |
| POST | `/wms/productionreturn-record-detail-hold/export-excel-senior` | exportProductionreturnRecordMainSeniorExcel | **隔离退料记录子** |
| GET | `/wms/productionreturn-record-detail-store/page` | getProductionreturnRecordDetailPage | **生产合格退料记录子** |
| POST | `/wms/productionreturn-record-detail-store/senior` | getProductionreturnRecordDetailSenior | **生产合格退料记录子** |
| GET | `/wms/productionreturn-record-detail-store/export-excel` | exportProductionreturnRecordDetailExcel | **生产合格退料记录子** |
| POST | `/wms/productionreturn-record-detail-store/export-excel-senior` | exportProductionreturnRecordMainSeniorExcel | **生产合格退料记录子** |
| POST | `/wms/productionreturn-record-main/create` | createProductionreturnRecordMain | **生产退料记录主** |
| PUT | `/wms/productionreturn-record-main/update` | updateProductionreturnRecordMain | **生产退料记录主** |
| DELETE | `/wms/productionreturn-record-main/delete` | deleteProductionreturnRecordMain | **生产退料记录主** |
| GET | `/wms/productionreturn-record-main/get` | getProductionreturnRecordMain | **编号** |
| GET | `/wms/productionreturn-record-main/list` | getProductionreturnRecordMainList | **编号** |
| GET | `/wms/productionreturn-record-main/page` | getProductionreturnRecordMainPage | **编号列表** |
| POST | `/wms/productionreturn-record-main/senior` | getProductionreturnRecordMainSenior | **生产退料记录主** |
| GET | `/wms/productionreturn-record-main/export-excel` | exportProductionreturnRecordMainExcel | **生产退料记录主** |
| POST | `/wms/productionreturn-record-main/export-excel-senior` | exportProductionreturnRecordMainSeniorExcel | **生产退料记录主** |
| PUT | `/wms/productionreturn-record-main/receive` | receive | **生产退料记录主** |
| PUT | `/wms/productionreturn-record-main/refuse` | refuse | **编号** |

### 生产退料申请 (productionreturnRequest)  (29个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productionreturn-request-detail/create` | createProductionreturnRequestDetail | **生产退料申请子** |
| PUT | `/wms/productionreturn-request-detail/update` | updateProductionreturnRequestDetail | **生产退料申请子** |
| DELETE | `/wms/productionreturn-request-detail/delete` | deleteProductionreturnRequestDetail | **生产退料申请子** |
| GET | `/wms/productionreturn-request-detail/get` | getProductionreturnRequestDetail | **编号** |
| GET | `/wms/productionreturn-request-detail/list` | getProductionreturnRequestDetailList | **编号** |
| GET | `/wms/productionreturn-request-detail/page` | getProductionreturnRequestDetailPage | **编号列表** |
| GET | `/wms/productionreturn-request-detail/export-excel` | exportProductionreturnRequestDetailExcel | **生产退料申请子** |
| POST | `/wms/productionreturn-request-detail/senior` | getProductionreturnRequestDetailSenior | **生产退料申请子** |
| PUT | `/wms/productionreturn-request-detail/updateDetailPackingNumber` | updateProductreceiptPackingNumber | **生产退料申请子** |
| POST | `/wms/productionreturn-request-main/create` | createProductionreturnRequestMain | **生产退料申请主** |
| PUT | `/wms/productionreturn-request-main/update` | updateProductionreturnRequestMain | **生产退料申请主** |
| DELETE | `/wms/productionreturn-request-main/delete` | deleteProductionreturnRequestMain | **生产退料申请主** |
| GET | `/wms/productionreturn-request-main/get` | getProductionreturnRequestMain | **编号** |
| GET | `/wms/productionreturn-request-main/list` | getProductionreturnRequestMainList | **编号** |
| GET | `/wms/productionreturn-request-main/page` | getProductionreturnRequestMainPage | **编号列表** |
| POST | `/wms/productionreturn-request-main/export-excel-senior` | exportProductionreturnRequestMainSeniorExcel | **生产退料申请主** |
| GET | `/wms/productionreturn-request-main/export-excel` | exportProductionreturnRequestMainExcel | **生产退料申请主** |
| POST | `/wms/productionreturn-request-main/senior` | getProductionreturnRequestMainSenior | **生产退料申请主** |
| GET | `/wms/productionreturn-request-main/getProductionreturnRequestById` | getProductionreturnRequestById | **生产退料申请主** |
| GET | `/wms/productionreturn-request-main/get-import-template` | importTemplate | **编号** |
| GET | `/wms/productionreturn-request-main/get-import-template-hold` | importTemplateHold | **生产退料申请主** |
| POST | `/wms/productionreturn-request-main/import` | importExcel | **生产退料申请主** |
| POST | `/wms/productionreturn-request-main/hold-import` | importExcelNO | **Excel 文件** |
| PUT | `/wms/productionreturn-request-main/close` | closeProductionreturnRequestMain | **Excel 文件** |
| PUT | `/wms/productionreturn-request-main/reAdd` | reAddProductionreturnRequestMain | **编号** |
| PUT | `/wms/productionreturn-request-main/submit` | submitProductionreturnRequestMain | **编号** |
| PUT | `/wms/productionreturn-request-main/refused` | refusedProductionreturnRequestMain | **编号** |
| PUT | `/wms/productionreturn-request-main/agree` | agreeProductionreturnRequestMain | **编号** |
| PUT | `/wms/productionreturn-request-main/handle` | handleProductionreturnRequestMain | **编号** |

### 生产报废记录 (productionscrapRecord)  (20个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productionscrap-record-detail/create` | createProductionscrapRecordDetail | **线边原料报废记录子** |
| PUT | `/wms/productionscrap-record-detail/update` | updateProductionscrapRecordDetail | **线边原料报废记录子** |
| DELETE | `/wms/productionscrap-record-detail/delete` | deleteProductionscrapRecordDetail | **线边原料报废记录子** |
| GET | `/wms/productionscrap-record-detail/get` | getProductionscrapRecordDetail | **编号** |
| GET | `/wms/productionscrap-record-detail/page` | getProductionscrapRecordDetailPage | **编号** |
| POST | `/wms/productionscrap-record-detail/senior` | getProductionscrapRecordDetailSenior | **线边原料报废记录子** |
| GET | `/wms/productionscrap-record-detail/export-excel` | exportProductionscrapRecordDetailExcel | **线边原料报废记录子** |
| GET | `/wms/productionscrap-record-detail/get-import-template` | importTemplate | **线边原料报废记录子** |
| POST | `/wms/productionscrap-record-detail/import` | importExcel | **线边原料报废记录子** |
| POST | `/wms/productionscrap-record-main/create` | createProductionscrapRecordMain | **线边原料报废记录主** |
| PUT | `/wms/productionscrap-record-main/update` | updateProductionscrapRecordMain | **线边原料报废记录主** |
| DELETE | `/wms/productionscrap-record-main/delete` | deleteProductionscrapRecordMain | **线边原料报废记录主** |
| GET | `/wms/productionscrap-record-main/get` | getProductionscrapRecordMain | **编号** |
| GET | `/wms/productionscrap-record-main/page` | getProductionscrapRecordMainPage | **编号** |
| POST | `/wms/productionscrap-record-main/senior` | getProductionscrapRecordMainSenior | **线边原料报废记录主** |
| GET | `/wms/productionscrap-record-main/export-excel` | exportProductionscrapRecordMainExcel | **线边原料报废记录主** |
| POST | `/wms/productionscrap-record-main/export-excel-senior` | exportProductionscrapRecordMainSeniorExcel | **线边原料报废记录主** |
| GET | `/wms/productionscrap-record-main/get-import-template` | importTemplate | **线边原料报废记录主** |
| POST | `/wms/productionscrap-record-main/import` | importExcel | **线边原料报废记录主** |
| GET | `/wms/productionscrap-record-main/revoke` | revoke | **Excel 文件** |

### 生产报废申请 (productionscrapRequest)  (24个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productionscrap-request-detail/create` | createProductionscrapRequestDetail | **线边原料报废子** |
| PUT | `/wms/productionscrap-request-detail/update` | updateProductionscrapRequestDetail | **线边原料报废子** |
| DELETE | `/wms/productionscrap-request-detail/delete` | deleteProductionscrapRequestDetail | **线边原料报废子** |
| GET | `/wms/productionscrap-request-detail/get` | getProductionscrapRequestDetail | **编号** |
| GET | `/wms/productionscrap-request-detail/page` | getProductionscrapRequestDetailPage | **编号** |
| POST | `/wms/productionscrap-request-detail/senior` | getProductionscrapRequestDetailSenior | **线边原料报废子** |
| GET | `/wms/productionscrap-request-detail/export-excel` | exportProductionscrapRequestDetailExcel | **线边原料报废子** |
| GET | `/wms/productionscrap-request-detail/get-import-template` | importTemplate | **线边原料报废子** |
| POST | `/wms/productionscrap-request-detail/import` | importExcel | **线边原料报废子** |
| POST | `/wms/productionscrap-request-main/create` | createProductionscrapRequestMain | **线边原料报废主** |
| PUT | `/wms/productionscrap-request-main/update` | updateProductionscrapRequestMain | **线边原料报废主** |
| DELETE | `/wms/productionscrap-request-main/delete` | deleteProductionscrapRequestMain | **线边原料报废主** |
| GET | `/wms/productionscrap-request-main/get` | getProductionscrapRequestMain | **编号** |
| GET | `/wms/productionscrap-request-main/page` | getProductionscrapRequestMainPage | **编号** |
| POST | `/wms/productionscrap-request-main/senior` | getProductionscrapRequestMainSenior | **线边原料报废主** |
| GET | `/wms/productionscrap-request-main/export-excel` | exportProductionscrapRequestMainExcel | **线边原料报废主** |
| GET | `/wms/productionscrap-request-main/get-import-template` | importTemplate | **线边原料报废主** |
| POST | `/wms/productionscrap-request-main/import` | importExcel | **线边原料报废主** |
| PUT | `/wms/productionscrap-request-main/close` | closeProductionscrapRequestMain | **Excel 文件** |
| PUT | `/wms/productionscrap-request-main/reAdd` | reAddProductionscrapRequestMain | **编号** |
| PUT | `/wms/productionscrap-request-main/submit` | submitProductionscrapRequestMain | **编号** |
| PUT | `/wms/productionscrap-request-main/refused` | refusedProductionscrapRequestMain | **编号** |
| PUT | `/wms/productionscrap-request-main/agree` | agreeProductionscrapRequestMain | **编号** |
| PUT | `/wms/productionscrap-request-main/handle` | handleProductionscrapRequestMain | **编号** |

### 生产上架任务 (productputawayJob)  (25个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productputaway-job-detail/create` | createProductputawayJobDetail | **制品上架任务子** |
| PUT | `/wms/productputaway-job-detail/update` | updateProductputawayJobDetail | **制品上架任务子** |
| DELETE | `/wms/productputaway-job-detail/delete` | deleteProductputawayJobDetail | **制品上架任务子** |
| GET | `/wms/productputaway-job-detail/get` | getProductputawayJobDetail | **编号** |
| GET | `/wms/productputaway-job-detail/list` | getProductputawayJobDetailList | **编号** |
| GET | `/wms/productputaway-job-detail/page` | getProductputawayJobDetailPage | **编号列表** |
| GET | `/wms/productputaway-job-detail/export-excel` | exportProductputawayJobDetailExcel | **制品上架任务子** |
| POST | `/wms/productputaway-job-detail/senior` | getProductputawayJobDetailSenior | **制品上架任务子** |
| POST | `/wms/productputaway-job-main/create` | createProductputawayJobMain | **制品上架任务主** |
| PUT | `/wms/productputaway-job-main/update` | updateProductputawayJobMain | **制品上架任务主** |
| DELETE | `/wms/productputaway-job-main/delete` | deleteProductputawayJobMain | **制品上架任务主** |
| GET | `/wms/productputaway-job-main/get` | getProductputawayJobMain | **编号** |
| GET | `/wms/productputaway-job-main/list` | getProductputawayJobMainList | **编号** |
| GET | `/wms/productputaway-job-main/page` | getProductputawayJobMainPage | **编号列表** |
| POST | `/wms/productputaway-job-main/senior` | getProductputawayJobMainSenior | **制品上架任务主** |
| GET | `/wms/productputaway-job-main/export-excel` | exportProductputawayJobMainExcel | **制品上架任务主** |
| POST | `/wms/productputaway-job-main/export-excel-senior` | exportProductputawayJobMainSeniorExcel | **制品上架任务主** |
| GET | `/wms/productputaway-job-main/getProductputawayJobById` | getProductputawayJobById | **制品上架任务主** |
| POST | `/wms/productputaway-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/productputaway-job-main/accept` | acceptProductputawayJobMain | **类型数组** |
| PUT | `/wms/productputaway-job-main/abandon` | abandonProductputawayJobMain | **制品上架任务主** |
| PUT | `/wms/productputaway-job-main/close` | closeProductputawayJobMain | **制品上架任务主** |
| PUT | `/wms/productputaway-job-main/execute` | executeProductputawayJobMain | **制品上架任务主** |
| PUT | `/wms/productputaway-job-main/updateProductputawayJobConfig` | updateProductputawayJobConfig | **制品上架任务主** |
| PUT | `/wms/productputaway-job-main/updateProductputawayJobAssembleConfig` | updateProductputawayJobAssembleConfig | **制品上架任务主** |

### 生产上架记录 (productputawayRecord)  (17个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productputaway-record-detail/create` | createProductputawayRecordDetail | **制品上架记录子** |
| PUT | `/wms/productputaway-record-detail/update` | updateProductputawayRecordDetail | **制品上架记录子** |
| DELETE | `/wms/productputaway-record-detail/delete` | deleteProductputawayRecordDetail | **制品上架记录子** |
| GET | `/wms/productputaway-record-detail/get` | getProductputawayRecordDetail | **编号** |
| GET | `/wms/productputaway-record-detail/list` | getProductputawayRecordDetailList | **编号** |
| GET | `/wms/productputaway-record-detail/page` | getProductputawayRecordDetailPage | **编号列表** |
| GET | `/wms/productputaway-record-detail/export-excel` | exportProductputawayRecordDetailExcel | **制品上架记录子** |
| POST | `/wms/productputaway-record-detail/senior` | getProductputawayRecordDetailSenior | **制品上架记录子** |
| POST | `/wms/productputaway-record-main/create` | createProductputawayRecordMain | **制品上架记录主** |
| PUT | `/wms/productputaway-record-main/update` | updateProductputawayRecordMain | **制品上架记录主** |
| DELETE | `/wms/productputaway-record-main/delete` | deleteProductputawayRecordMain | **制品上架记录主** |
| GET | `/wms/productputaway-record-main/get` | getProductputawayRecordMain | **编号** |
| GET | `/wms/productputaway-record-main/list` | getProductputawayRecordMainList | **编号** |
| GET | `/wms/productputaway-record-main/page` | getProductputawayRecordMainPage | **编号列表** |
| POST | `/wms/productputaway-record-main/senior` | getProductputawayRecordMainSenior | **制品上架记录主** |
| GET | `/wms/productputaway-record-main/export-excel` | exportProductputawayRecordMainExcel | **制品上架记录主** |
| POST | `/wms/productputaway-record-main/export-excel-senior` | exportProductputawayRecordMainSeniorExcel | **制品上架记录主** |

### 生产上架申请 (productputawayRequest)  (26个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productputaway-request-detail/create` | createProductputawayRequestDetail | **制品上架申请子** |
| PUT | `/wms/productputaway-request-detail/update` | updateProductputawayRequestDetail | **制品上架申请子** |
| DELETE | `/wms/productputaway-request-detail/delete` | deleteProductputawayRequestDetail | **制品上架申请子** |
| GET | `/wms/productputaway-request-detail/get` | getProductputawayRequestDetail | **编号** |
| GET | `/wms/productputaway-request-detail/list` | getProductputawayRequestDetailList | **编号** |
| GET | `/wms/productputaway-request-detail/page` | getProductputawayRequestDetailPage | **编号列表** |
| GET | `/wms/productputaway-request-detail/export-excel` | exportProductputawayRequestDetailExcel | **制品上架申请子** |
| POST | `/wms/productputaway-request-detail/senior` | getProductputawayRequestDetailSenior | **制品上架申请子** |
| POST | `/wms/productputaway-request-main/create` | createProductputawayRequestMain | **制品上架申请主** |
| PUT | `/wms/productputaway-request-main/update` | updateProductputawayRequestMain | **制品上架申请主** |
| DELETE | `/wms/productputaway-request-main/delete` | deleteProductputawayRequestMain | **制品上架申请主** |
| GET | `/wms/productputaway-request-main/get` | getProductputawayRequestMain | **编号** |
| GET | `/wms/productputaway-request-main/list` | getProductputawayRequestMainList | **编号** |
| GET | `/wms/productputaway-request-main/page` | getProductputawayRequestMainPage | **编号列表** |
| POST | `/wms/productputaway-request-main/senior` | getProductputawayRequestMainSenior | **制品上架申请主** |
| GET | `/wms/productputaway-request-main/get-import-template` | importTemplate | **制品上架申请主** |
| GET | `/wms/productputaway-request-main/export-excel` | exportProductputawayRequestMainExcel | **制品上架申请主** |
| POST | `/wms/productputaway-request-main/export-excel-senior` | exportProductputawayRequestMainSeniorExcel | **制品上架申请主** |
| POST | `/wms/productputaway-request-main/import` | importExcel | **制品上架申请主** |
| GET | `/wms/productputaway-request-main/getProductputawayRequestById` | getProductputawayRequestById | **Excel 文件** |
| PUT | `/wms/productputaway-request-main/close` | closeProductputawayRequestMain | **编号** |
| PUT | `/wms/productputaway-request-main/reAdd` | reAddProductputawayRequestMain | **编号** |
| PUT | `/wms/productputaway-request-main/submit` | submitProductputawayRequestMain | **编号** |
| PUT | `/wms/productputaway-request-main/refused` | refusedProductputawayRequestMain | **编号** |
| PUT | `/wms/productputaway-request-main/agree` | agreeProductputawayRequestMain | **编号** |
| PUT | `/wms/productputaway-request-main/handle` | handleProductputawayRequestMain | **编号** |

### 生产收货明细 (productreceiptDetailb)  (14个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productreceipt-detailb/create` | createProductreceiptDetailb | **制品收货记录子** |
| PUT | `/wms/productreceipt-detailb/update` | updateProductreceiptDetailb | **制品收货记录子** |
| DELETE | `/wms/productreceipt-detailb/delete` | deleteProductreceiptDetailb | **制品收货记录子** |
| GET | `/wms/productreceipt-detailb/get` | getProductreceiptDetailb | **编号** |
| GET | `/wms/productreceipt-detailb/list` | getProductreceiptDetailbList | **编号** |
| GET | `/wms/productreceipt-detailb/page` | getProductreceiptDetailbPage | **编号列表** |
| GET | `/wms/productreceipt-detailb/export-excel` | exportProductreceiptDetailbExcel | **制品收货记录子** |
| GET | `/wms/productreceipt-detailb/pageChildPackingNumber` | getProductreceiptDetailbPageChildPackingNumber | **制品收货记录子** |
| GET | `/wms/productreceipt-detailb/getCheckWhetherItExists` | getCheckWhetherItExists | **制品收货记录子** |
| GET | `/wms/productreceipt-detailb/getProductreceiptDetailbByPackingNumber` | getProductreceiptDetailbByPackingNumber | **制品收货记录子** |
| GET | `/wms/productreceipt-detailb/getAssemblyMaterialUsageMes` | getAssemblyMaterialUsageMes | **制品收货记录子** |
| GET | `/wms/productreceipt-detailb/assemblyMaterialUsageMesExport` | exportAssemblyMaterialUsageExcel | **制品收货记录子** |
| POST | `/wms/productreceipt-detailb/getAssemblyMaterialUsageMesSenior` | getAssemblyMaterialUsageMesSenior | **制品收货记录子** |
| POST | `/wms/productreceipt-detailb/getAssemblyMaterialUsageMesSeniorExport` | getAssemblyMaterialUsageMesSeniorExport | **制品收货记录子** |

### 生产收货任务 (productreceiptJob)  (27个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productreceipt-job-detail/create` | createProductreceiptJobDetail | **制品收货任务子** |
| PUT | `/wms/productreceipt-job-detail/update` | updateProductreceiptJobDetail | **制品收货任务子** |
| DELETE | `/wms/productreceipt-job-detail/delete` | deleteProductreceiptJobDetail | **制品收货任务子** |
| GET | `/wms/productreceipt-job-detail/get` | getProductreceiptJobDetail | **编号** |
| GET | `/wms/productreceipt-job-detail/list` | getProductreceiptJobDetailList | **编号** |
| GET | `/wms/productreceipt-job-detail/page` | getProductreceiptJobDetailPage | **编号列表** |
| GET | `/wms/productreceipt-job-detail/export-excel` | exportProductreceiptJobDetailExcel | **制品收货任务子** |
| POST | `/wms/productreceipt-job-detail/senior` | getProductreceiptJobDetailSenior | **制品收货任务子** |
| POST | `/wms/productreceipt-job-main/create` | createProductreceiptJobMain | **制品收货任务主** |
| PUT | `/wms/productreceipt-job-main/update` | updateProductreceiptJobMain | **制品收货任务主** |
| DELETE | `/wms/productreceipt-job-main/delete` | deleteProductreceiptJobMain | **制品收货任务主** |
| GET | `/wms/productreceipt-job-main/get` | getProductreceiptJobMain | **编号** |
| GET | `/wms/productreceipt-job-main/list` | getProductreceiptJobMainList | **编号** |
| GET | `/wms/productreceipt-job-main/page` | getProductreceiptJobMainPage | **编号列表** |
| POST | `/wms/productreceipt-job-main/senior` | getProductreceiptJobMainSenior | **制品收货任务主** |
| GET | `/wms/productreceipt-job-main/export-excel` | exportProductreceiptJobMainExcel | **制品收货任务主** |
| POST | `/wms/productreceipt-job-main/export-excel-senior` | exportProductreceiptJobMainSeniorExcel | **制品收货任务主** |
| GET | `/wms/productreceipt-job-main/getProductreceiptJobById` | getProductreceiptJobById | **制品收货任务主** |
| POST | `/wms/productreceipt-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/productreceipt-job-main/accept` | acceptProductreceiptJobMain | **类型数组** |
| PUT | `/wms/productreceipt-job-main/abandon` | abandonProductreceiptJobMain | **制品收货任务主** |
| PUT | `/wms/productreceipt-job-main/close` | closeProductreceiptJobMain | **制品收货任务主** |
| PUT | `/wms/productreceipt-job-main/execute` | executeProductreceiptJobMain | **制品收货任务主** |
| PUT | `/wms/productreceipt-job-main/executeForCC` | executeProductreceiptJobMainByCC | **制品收货任务主** |
| PUT | `/wms/productreceipt-job-main/updateProductreceiptJobConfig` | updateProductreceiptJobConfig | **制品收货任务主** |
| PUT | `/wms/productreceipt-job-main/updateProductreceiptJobAssembleConfig` | updateProductreceiptJobAssembleConfig | **制品收货任务主** |
| PUT | `/wms/productreceipt-job-main/updateProductreceiptScrapJobConfig` | updateProductreceiptScrapJobConfig | **制品收货任务主** |

### 生产收货记录 (productreceiptRecord)  (42个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/backflush-record-detailb/create` | createBackflushRecordDetailb | **制品收货记录子** |
| PUT | `/wms/backflush-record-detailb/update` | updateBackflushRecordDetailb | **制品收货记录子** |
| DELETE | `/wms/backflush-record-detailb/delete` | deleteBackflushRecordDetailb | **制品收货记录子** |
| GET | `/wms/backflush-record-detailb/get` | getBackflushRecordDetailb | **编号** |
| GET | `/wms/backflush-record-detailb/list` | getBackflushRecordDetailbList | **编号** |
| GET | `/wms/backflush-record-detailb/page` | getBackflushRecordDetailbPage | **编号列表** |
| GET | `/wms/backflush-record-detailb/export-excel` | exportBackflushRecordDetailbExcel | **制品收货记录子** |
| POST | `/wms/backflush-record-detailb/senior` | getBackflushRecordDetailbSenior | **制品收货记录子** |
| GET | `/wms/backflush-record-detailb/getAssemblyMaterialUsage` | getAssemblyMaterialUsage | **制品收货记录子** |
| GET | `/wms/backflush-record-detailb/assemblyMaterialUsageExport` | exportAssemblyMaterialUsageExcel | **制品收货记录子** |
| POST | `/wms/backflush-record-detailb/getAssemblyMaterialUsageSenior` | getAssemblyMaterialUsageSenior | **制品收货记录子** |
| POST | `/wms/backflush-record-detailb/getAssemblyMaterialUsageSeniorExport` | getAssemblyMaterialUsageSeniorExport | **制品收货记录子** |
| GET | `/wms/backflush-record-detailb/get-import-template-error` | importTemplate | **制品收货记录子** |
| POST | `/wms/backflush-record-detailb/importError` | importExcel | **制品收货记录子** |
| POST | `/wms/productreceipt-record-detail/create` | createProductreceiptRecordDetail | **制品收货记录子** |
| PUT | `/wms/productreceipt-record-detail/update` | updateProductreceiptRecordDetail | **制品收货记录子** |
| DELETE | `/wms/productreceipt-record-detail/delete` | deleteProductreceiptRecordDetail | **制品收货记录子** |
| GET | `/wms/productreceipt-record-detail/get` | getProductreceiptRecordDetail | **编号** |
| GET | `/wms/productreceipt-record-detail/list` | getProductreceiptRecordDetailList | **编号** |
| GET | `/wms/productreceipt-record-detail/page` | getProductreceiptRecordDetailPage | **编号列表** |
| GET | `/wms/productreceipt-record-detail/export-excel` | exportProductreceiptRecordDetailExcel | **制品收货记录子** |
| POST | `/wms/productreceipt-record-detail/senior` | getProductreceiptRecordDetailSenior | **制品收货记录子** |
| POST | `/wms/productreceipt-record-main/create` | createProductreceiptRecordMain | **制品收货记录主** |
| POST | `/wms/productreceipt-record-main/createByPlan` | createProductreceiptRecordMainByPlan | **制品收货记录主** |
| POST | `/wms/productreceipt-record-main/createByPlanToSenior` | createProductreceiptRecordMainByPlanToSenior | **制品收货记录主** |
| POST | `/wms/productreceipt-record-main/createProductreceiptAGV` | createProductreceiptRecordMain | **制品收货记录主** |
| PUT | `/wms/productreceipt-record-main/update` | updateProductreceiptRecordMain | **制品收货记录主** |
| DELETE | `/wms/productreceipt-record-main/delete` | deleteProductreceiptRecordMain | **制品收货记录主** |
| GET | `/wms/productreceipt-record-main/get` | getProductreceiptRecordMain | **编号** |
| GET | `/wms/productreceipt-record-main/list` | getProductreceiptRecordMainList | **编号** |
| GET | `/wms/productreceipt-record-main/page` | getProductreceiptRecordMainPage | **编号列表** |
| POST | `/wms/productreceipt-record-main/senior` | getProductreceiptRecordMainSenior | **制品收货记录主** |
| GET | `/wms/productreceipt-record-main/export-excel` | exportProductreceiptRecordMainExcel | **制品收货记录主** |
| POST | `/wms/productreceipt-record-main/export-excel-senior` | exportProductreceiptRecordMainSeniorExcel | **制品收货记录主** |
| GET | `/wms/productreceipt-record-main/export-excel-scrap` | exportProductreceiptRecordMainExcelScrap | **制品收货记录主** |
| POST | `/wms/productreceipt-record-main/export-excel-senior-scrap` | exportProductreceiptRecordMainSeniorExcelScarp | **制品收货记录主** |
| POST | `/wms/productreceipt-record-main/createInspectRequest` | createInspectRequest | **制品收货记录主** |
| POST | `/wms/productreceipt-record-main/createPutawayRequest` | createPutawayRequest | **制品收货记录主** |
| PUT | `/wms/productreceipt-record-main/handleRecovery` | handleProductreceiptRecordMain | **制品收货记录主** |
| POST | `/wms/productreceipt-record-main/replenishmentProductreceiptTransaction` | replenishmentProductreceiptTransaction | **制品收货记录主** |
| PUT | `/wms/productreceipt-record-main/receive` | receive | **制品收货记录主** |
| PUT | `/wms/productreceipt-record-main/refuse` | refuse | **编号** |

### 生产收货申请 (productreceiptRequest)  (35个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productreceipt-request-detail/create` | createProductreceiptRequestDetail | **制品收货申请子** |
| PUT | `/wms/productreceipt-request-detail/update` | updateProductreceiptRequestDetail | **制品收货申请子** |
| DELETE | `/wms/productreceipt-request-detail/delete` | deleteProductreceiptRequestDetail | **制品收货申请子** |
| GET | `/wms/productreceipt-request-detail/get` | getProductreceiptRequestDetail | **编号** |
| GET | `/wms/productreceipt-request-detail/list` | getProductreceiptRequestDetailList | **编号** |
| GET | `/wms/productreceipt-request-detail/page` | getProductreceiptRequestDetailPage | **编号列表** |
| GET | `/wms/productreceipt-request-detail/export-excel` | exportProductreceiptRequestDetailExcel | **制品收货申请子** |
| POST | `/wms/productreceipt-request-detail/senior` | getProductreceiptRequestDetailSenior | **制品收货申请子** |
| PUT | `/wms/productreceipt-request-detail/updateDetailPackingNumber` | updateProductreceiptPackingNumber | **制品收货申请子** |
| POST | `/wms/productreceipt-request-detailb/create` | createProductreceiptRequestDetailb | **制品收货申请子** |
| PUT | `/wms/productreceipt-request-detailb/update` | updateProductreceiptRequestDetailb | **制品收货申请子** |
| DELETE | `/wms/productreceipt-request-detailb/delete` | deleteProductreceiptRequestDetailb | **制品收货申请子** |
| GET | `/wms/productreceipt-request-detailb/get` | getProductreceiptRequestDetailb | **编号** |
| GET | `/wms/productreceipt-request-detailb/list` | getProductreceiptRequestDetailbList | **编号** |
| GET | `/wms/productreceipt-request-detailb/page` | getProductreceiptRequestDetailbPage | **编号列表** |
| GET | `/wms/productreceipt-request-detailb/export-excel` | exportProductreceiptRequestDetailbExcel | **制品收货申请子** |
| GET | `/wms/productreceipt-request-detailb/get-import-template` | importTemplate | **制品收货申请子** |
| POST | `/wms/productreceipt-request-main/create` | createProductreceiptRequestMain | **制品收货申请主** |
| PUT | `/wms/productreceipt-request-main/update` | updateProductreceiptRequestMain | **制品收货申请主** |
| DELETE | `/wms/productreceipt-request-main/delete` | deleteProductreceiptRequestMain | **制品收货申请主** |
| GET | `/wms/productreceipt-request-main/get` | getProductreceiptRequestMain | **编号** |
| GET | `/wms/productreceipt-request-main/list` | getProductreceiptRequestMainList | **编号** |
| GET | `/wms/productreceipt-request-main/page` | getProductreceiptRequestMainPage | **编号列表** |
| POST | `/wms/productreceipt-request-main/senior` | getProductreceiptRequestMainSenior | **制品收货申请主** |
| POST | `/wms/productreceipt-request-main/export-excel-senior` | exportProductreceiptRequestMainSeniorExcel | **制品收货申请主** |
| GET | `/wms/productreceipt-request-main/export-excel` | exportProductreceiptRequestMainExcel | **制品收货申请主** |
| GET | `/wms/productreceipt-request-main/get-import-template` | importTemplate | **制品收货申请主** |
| POST | `/wms/productreceipt-request-main/import` | importExcel | **制品收货申请主** |
| PUT | `/wms/productreceipt-request-main/close` | closeProductreceiptRequestMain | **Excel 文件** |
| PUT | `/wms/productreceipt-request-main/reAdd` | reAddProductreceiptRequestMain | **编号** |
| PUT | `/wms/productreceipt-request-main/submit` | submitProductreceiptRequestMain | **编号** |
| PUT | `/wms/productreceipt-request-main/refused` | refusedProductreceiptRequestMain | **编号** |
| PUT | `/wms/productreceipt-request-main/agree` | agreeProductreceiptRequestMain | **编号** |
| PUT | `/wms/productreceipt-request-main/handle` | handleProductreceiptRequestMain | **编号** |
| POST | `/wms/productreceipt-request-main/productCreateLabel` | productCreateLabel | **编号** |

### 生产整改任务 (productredressJob)  (24个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productredress-job-detail/create` | createProductredressJobDetail | **制品回收任务子** |
| PUT | `/wms/productredress-job-detail/update` | updateProductredressJobDetail | **制品回收任务子** |
| DELETE | `/wms/productredress-job-detail/delete` | deleteProductredressJobDetail | **制品回收任务子** |
| GET | `/wms/productredress-job-detail/get` | getProductredressJobDetail | **编号** |
| GET | `/wms/productredress-job-detail/list` | getProductredressJobDetailList | **编号** |
| GET | `/wms/productredress-job-detail/page` | getProductredressJobDetailPage | **编号列表** |
| GET | `/wms/productredress-job-detail/export-excel` | exportProductredressJobDetailExcel | **制品回收任务子** |
| POST | `/wms/productredress-job-detail/senior` | getProductreceiptJobDetailSenior | **制品回收任务子** |
| POST | `/wms/productredress-job-main/create` | createProductredressJobMain | **制品回收任务主** |
| PUT | `/wms/productredress-job-main/update` | updateProductredressJobMain | **制品回收任务主** |
| DELETE | `/wms/productredress-job-main/delete` | deleteProductredressJobMain | **制品回收任务主** |
| GET | `/wms/productredress-job-main/get` | getProductredressJobMain | **编号** |
| GET | `/wms/productredress-job-main/list` | getProductredressJobMainList | **编号** |
| GET | `/wms/productredress-job-main/page` | getProductredressJobMainPage | **编号列表** |
| POST | `/wms/productredress-job-main/senior` | getProductreceiptJobMainSenior | **制品回收任务主** |
| GET | `/wms/productredress-job-main/export-excel` | exportProductredressJobMainExcel | **制品回收任务主** |
| POST | `/wms/productredress-job-main/export-excel-senior` | exportProductredressJobMainSeniorExcel | **制品回收任务主** |
| GET | `/wms/productredress-job-main/getProductredressJobById` | getProductredressJobById | **制品回收任务主** |
| POST | `/wms/productredress-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/productredress-job-main/accept` | acceptProductredressJobMain | **类型数组** |
| PUT | `/wms/productredress-job-main/abandon` | abandonProductredressJobMain | **制品回收任务主** |
| PUT | `/wms/productredress-job-main/close` | closeProductredressJobMain | **制品回收任务主** |
| PUT | `/wms/productredress-job-main/execute` | executeProductredressJobMain | **制品回收任务主** |
| PUT | `/wms/productredress-job-main/updateProductredressJobConfig` | updateProductredressJobConfig | **制品回收任务主** |

### 生产整改记录 (productredressRecord)  (16个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productredress-record-detail/create` | createProductredressRecordDetail | **制品回收记录子** |
| PUT | `/wms/productredress-record-detail/update` | updateProductredressRecordDetail | **制品回收记录子** |
| DELETE | `/wms/productredress-record-detail/delete` | deleteProductredressRecordDetail | **制品回收记录子** |
| GET | `/wms/productredress-record-detail/get` | getProductredressRecordDetail | **编号** |
| GET | `/wms/productredress-record-detail/list` | getProductredressRecordDetailList | **编号** |
| GET | `/wms/productredress-record-detail/page` | getProductredressRecordDetailPage | **编号列表** |
| POST | `/wms/productredress-record-detail/senior` | ggetProductredressRecordDetailSenior | **制品回收记录子** |
| POST | `/wms/productredress-record-main/create` | createProductredressRecordMain | **制品回收记录主** |
| PUT | `/wms/productredress-record-main/update` | updateProductredressRecordMain | **制品回收记录主** |
| DELETE | `/wms/productredress-record-main/delete` | deleteProductredressRecordMain | **制品回收记录主** |
| GET | `/wms/productredress-record-main/get` | getProductredressRecordMain | **编号** |
| GET | `/wms/productredress-record-main/list` | getProductredressRecordMainList | **编号** |
| GET | `/wms/productredress-record-main/page` | getProductredressRecordMainPage | **编号列表** |
| POST | `/wms/productredress-record-main/senior` | getProductreceiptRecordMainSenior | **制品回收记录主** |
| GET | `/wms/productredress-record-main/export-excel` | exportProductredressRecordMainExcel | **制品回收记录主** |
| POST | `/wms/productredress-record-main/export-excel-senior` | exportProductredressRecordMainSeniorExcel | **制品回收记录主** |

### 生产整改申请 (productredressRequest)  (25个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productredress-request-detail/create` | createProductredressRequestDetail | **制品回收申请子** |
| PUT | `/wms/productredress-request-detail/update` | updateProductredressRequestDetail | **制品回收申请子** |
| DELETE | `/wms/productredress-request-detail/delete` | deleteProductredressRequestDetail | **制品回收申请子** |
| GET | `/wms/productredress-request-detail/get` | getProductredressRequestDetail | **编号** |
| GET | `/wms/productredress-request-detail/list` | getProductredressRequestDetailList | **编号** |
| GET | `/wms/productredress-request-detail/page` | getProductredressRequestDetailPage | **编号列表** |
| POST | `/wms/productredress-request-detail/senior` | getProductredressRequestDetailSenior | **制品回收申请子** |
| POST | `/wms/productredress-request-main/create` | createProductredressRequestMain | **制品回收申请主** |
| PUT | `/wms/productredress-request-main/update` | updateProductredressRequestMain | **制品回收申请主** |
| DELETE | `/wms/productredress-request-main/delete` | deleteProductredressRequestMain | **制品回收申请主** |
| GET | `/wms/productredress-request-main/get` | getProductredressRequestMain | **编号** |
| GET | `/wms/productredress-request-main/list` | getProductredressRequestMainList | **编号** |
| GET | `/wms/productredress-request-main/page` | getProductredressRequestMainPage | **编号列表** |
| POST | `/wms/productredress-request-main/senior` | getProductredressRequestMainSenior | **制品回收申请主** |
| POST | `/wms/productredress-request-main/export-excel-senior` | exportProductredressRequestMainSeniorExcel | **制品回收申请主** |
| GET | `/wms/productredress-request-main/export-excel` | exportProductredressRequestMainExcel | **制品回收申请主** |
| GET | `/wms/productredress-request-main/get-import-template` | importTemplate | **制品回收申请主** |
| POST | `/wms/productredress-request-main/import` | importExcel | **制品回收申请主** |
| GET | `/wms/productredress-request-main/getProductredressRequestById` | getProductredressRequestById | **Excel 文件** |
| PUT | `/wms/productredress-request-main/close` | closeProductredressRequestMain | **编号** |
| PUT | `/wms/productredress-request-main/reAdd` | reAddProductredressRequestMain | **编号** |
| PUT | `/wms/productredress-request-main/submit` | submitProductredressRequestMain | **编号** |
| PUT | `/wms/productredress-request-main/refused` | refusedProductredressRequestMain | **编号** |
| PUT | `/wms/productredress-request-main/agree` | agreeProductredressRequestMain | **编号** |
| PUT | `/wms/productredress-request-main/handle` | handleProductredressRequestMain | **编号** |

### 生产维修记录 (productrepairRecord)  (26个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/consume-record-detailb/create` | createConsumeRecordDetailb | **制品返修记录子** |
| PUT | `/wms/consume-record-detailb/update` | updateConsumeRecordDetailb | **制品返修记录子** |
| DELETE | `/wms/consume-record-detailb/delete` | deleteConsumeRecordDetailb | **制品返修记录子** |
| GET | `/wms/consume-record-detailb/get` | getConsumeRecordDetailb | **编号** |
| GET | `/wms/consume-record-detailb/list` | getConsumeRecordDetailbList | **编号** |
| GET | `/wms/consume-record-detailb/page` | getConsumeRecordDetailbPage | **编号列表** |
| GET | `/wms/consume-record-detailb/export-excel` | exportConsumeRecordDetailbExcel | **制品返修记录子** |
| POST | `/wms/consume-record-detailb/senior` | getConsumeRecordDetailbSenior | **制品返修记录子** |
| POST | `/wms/productrepair-record-detail/create` | createProductrepairRecordDetail | **制品返修记录子** |
| PUT | `/wms/productrepair-record-detail/update` | updateProductrepairRecordDetail | **制品返修记录子** |
| DELETE | `/wms/productrepair-record-detail/delete` | deleteProductrepairRecordDetail | **制品返修记录子** |
| GET | `/wms/productrepair-record-detail/get` | getProductrepairRecordDetail | **编号** |
| GET | `/wms/productrepair-record-detail/list` | getProductrepairRecordDetailList | **编号** |
| GET | `/wms/productrepair-record-detail/page` | getProductrepairRecordDetailPage | **编号列表** |
| GET | `/wms/productrepair-record-detail/export-excel` | exportProductrepairRecordDetailExcel | **制品返修记录子** |
| POST | `/wms/productrepair-record-detail/senior` | getProductrepairRecordDetailSenior | **制品返修记录子** |
| POST | `/wms/productrepair-record-main/create` | createProductrepairRecordMain | **制品返修记录主** |
| PUT | `/wms/productrepair-record-main/update` | updateProductrepairRecordMain | **制品返修记录主** |
| DELETE | `/wms/productrepair-record-main/delete` | deleteProductrepairRecordMain | **制品返修记录主** |
| GET | `/wms/productrepair-record-main/get` | getProductrepairRecordMain | **编号** |
| GET | `/wms/productrepair-record-main/list` | getProductrepairRecordMainList | **编号** |
| GET | `/wms/productrepair-record-main/page` | getProductrepairRecordMainPage | **编号列表** |
| GET | `/wms/productrepair-record-main/export-excel` | exportProductrepairRecordMainExcel | **制品返修记录主** |
| POST | `/wms/productrepair-record-main/export-excel-senior` | exportProductrepairRecordMainSeniorExcel | **制品返修记录主** |
| POST | `/wms/productrepair-record-main/senior` | getProductrepairRecordMainSenior | **制品返修记录主** |
| GET | `/wms/productrepair-record-main/bomPage` | getBomInfoPage | **制品返修记录主** |

### 生产维修申请 (productrepairRequest)  (35个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productrepair-request-detaila/create` | createProductrepairRequestDetaila | **制品返修申请子** |
| PUT | `/wms/productrepair-request-detaila/update` | updateProductrepairRequestDetaila | **制品返修申请子** |
| DELETE | `/wms/productrepair-request-detaila/delete` | deleteProductrepairRequestDetaila | **制品返修申请子** |
| GET | `/wms/productrepair-request-detaila/get` | getProductrepairRequestDetaila | **编号** |
| GET | `/wms/productrepair-request-detaila/list` | getProductrepairRequestDetailaList | **编号** |
| GET | `/wms/productrepair-request-detaila/page` | getProductrepairRequestDetailaPage | **编号列表** |
| GET | `/wms/productrepair-request-detaila/export-excel` | exportProductrepairRequestDetailaExcel | **制品返修申请子** |
| POST | `/wms/productrepair-request-detaila/senior` | getProductrepairRequestDetailaSenior | **制品返修申请子** |
| POST | `/wms/consumere-request-detailb/create` | createConsumereRequestDetailb | **制品返修申请子** |
| PUT | `/wms/consumere-request-detailb/update` | updateConsumereRequestDetailb | **制品返修申请子** |
| DELETE | `/wms/consumere-request-detailb/delete` | deleteConsumereRequestDetailb | **制品返修申请子** |
| GET | `/wms/consumere-request-detailb/get` | getConsumereRequestDetailb | **编号** |
| GET | `/wms/consumere-request-detailb/list` | getConsumereRequestDetailbList | **编号** |
| GET | `/wms/consumere-request-detailb/page` | getConsumereRequestDetailbPage | **编号列表** |
| POST | `/wms/consumere-request-detailb/senior` | getConsumereSenior | **制品返修申请子** |
| GET | `/wms/consumere-request-detailb/export-excel` | exportConsumereRequestDetailbExcel | **制品返修申请子** |
| POST | `/wms/productrepair-request-main/create` | createProductrepairRequestMain | **制品返修申请主** |
| PUT | `/wms/productrepair-request-main/update` | updateProductrepairRequestMain | **制品返修申请主** |
| DELETE | `/wms/productrepair-request-main/delete` | deleteProductrepairRequestMain | **制品返修申请主** |
| GET | `/wms/productrepair-request-main/get` | getProductrepairRequestMain | **编号** |
| GET | `/wms/productrepair-request-main/list` | getProductrepairRequestMainList | **编号** |
| GET | `/wms/productrepair-request-main/page` | getProductrepairRequestMainPage | **编号列表** |
| POST | `/wms/productrepair-request-main/export-excel-senior` | exportProductrepairRequestMainSeniorExcel | **制品返修申请主** |
| GET | `/wms/productrepair-request-main/export-excel` | exportProductrepairRequestMainExcel | **制品返修申请主** |
| POST | `/wms/productrepair-request-main/senior` | getProductrepairRequestMainSenior | **制品返修申请主** |
| GET | `/wms/productrepair-request-main/get-import-template` | importTemplate | **制品返修申请主** |
| POST | `/wms/productrepair-request-main/import` | importExcel | **制品返修申请主** |
| PUT | `/wms/productrepair-request-main/close` | closeProductrepairRequestMain | **Excel 文件** |
| PUT | `/wms/productrepair-request-main/reAdd` | reAddeProductrepairRequestMain | **编号** |
| PUT | `/wms/productrepair-request-main/submit` | submitProductrepairRequestMain | **编号** |
| PUT | `/wms/productrepair-request-main/refused` | refusedProductrepairRequestMain | **编号** |
| PUT | `/wms/productrepair-request-main/agree` | agreeProductrepairRequestMain | **编号** |
| PUT | `/wms/productrepair-request-main/handle` | handleProductrepairRequestMain | **编号** |
| GET | `/wms/productrepair-request-main/bomPage` | getBomInfoPage | **编号** |
| POST | `/wms/productrepair-request-main/updateBom` | updateProductrepairDetailRequestBom | **制品返修申请主** |

### 生产报废任务 (productscrapJob)  (17个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productscrap-job-detail/create` | createProductscrapJobDetail | **制品报废任务子** |
| PUT | `/wms/productscrap-job-detail/update` | updateProductscrapJobDetail | **制品报废任务子** |
| DELETE | `/wms/productscrap-job-detail/delete` | deleteProductscrapJobDetail | **制品报废任务子** |
| GET | `/wms/productscrap-job-detail/get` | getProductscrapJobDetail | **编号** |
| GET | `/wms/productscrap-job-detail/list` | getProductscrapJobDetailList | **编号** |
| GET | `/wms/productscrap-job-detail/page` | getProductscrapJobDetailPage | **编号列表** |
| POST | `/wms/productscrap-job-detail/senior` | getProductscrapJobDetailSenior | **制品报废任务子** |
| GET | `/wms/productscrap-job-detail/export-excel` | exportProductscrapJobDetailExcel | **制品报废任务子** |
| POST | `/wms/productscrap-job-main/senior` | getProductscrapJobMainSenior | **制品报废任务主** |
| POST | `/wms/productscrap-job-main/create` | createProductscrapJobMain | **制品报废任务主** |
| PUT | `/wms/productscrap-job-main/update` | updateProductscrapJobMain | **制品报废任务主** |
| DELETE | `/wms/productscrap-job-main/delete` | deleteProductscrapJobMain | **制品报废任务主** |
| GET | `/wms/productscrap-job-main/get` | getProductscrapJobMain | **编号** |
| GET | `/wms/productscrap-job-main/list` | getProductscrapJobMainList | **编号** |
| GET | `/wms/productscrap-job-main/page` | getProductscrapJobMainPage | **编号列表** |
| GET | `/wms/productscrap-job-main/export-excel` | exportProductscrapJobMainExcel | **制品报废任务主** |
| GET | `/wms/productscrap-job-main/export-excel-senior` | exportProductscrapJobMainSeniorExcel | **制品报废任务主** |

### 生产报废记录 (productscrapRecord)  (29个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productscrap-record-detail/create` | createProductscrapRecordDetail | **制品报废记录子** |
| PUT | `/wms/productscrap-record-detail/update` | updateProductscrapRecordDetail | **制品报废记录子** |
| DELETE | `/wms/productscrap-record-detail/delete` | deleteProductscrapRecordDetail | **制品报废记录子** |
| GET | `/wms/productscrap-record-detail/get` | getProductscrapRecordDetail | **编号** |
| GET | `/wms/productscrap-record-detail/list` | getProductscrapRecordDetailList | **编号** |
| GET | `/wms/productscrap-record-detail/page` | getProductscrapRecordDetailPage | **编号列表** |
| POST | `/wms/productscrap-record-detail/senior` | getProductscrapRecordDetailSenior | **制品报废记录子** |
| GET | `/wms/productscrap-record-detail/export-excel` | exportProductscrapRecordDetailExcel | **制品报废记录子** |
| POST | `/wms/productscrap-record-main/create` | createProductscrapRecordMain | **制品报废记录主** |
| PUT | `/wms/productscrap-record-main/update` | updateProductscrapRecordMain | **制品报废记录主** |
| DELETE | `/wms/productscrap-record-main/delete` | deleteProductscrapRecordMain | **制品报废记录主** |
| GET | `/wms/productscrap-record-main/get` | getProductscrapRecordMain | **编号** |
| GET | `/wms/productscrap-record-main/list` | getProductscrapRecordMainList | **编号** |
| GET | `/wms/productscrap-record-main/page` | getProductscrapRecordMainPage | **编号列表** |
| POST | `/wms/productscrap-record-main/senior` | getProductscrapRecordMainSenior | **制品报废记录主** |
| GET | `/wms/productscrap-record-main/export-excel` | exportProductscrapRecordMainExcel | **制品报废记录主** |
| POST | `/wms/productscrap-record-main/export-excel-senior` | exportProductscrapRecordMainSeniorExcel | **制品报废记录主** |
| GET | `/wms/productscrap-record-main/bomPage` | getBomInfoPage | **制品报废记录主** |
| GET | `/wms/productscrap-record-main/revoke` | revoke | **制品报废记录主** |
| POST | `/wms/rawscrap-record-detail/create` | createRawscrapRecordDetail | **制品返修记录子** |
| PUT | `/wms/rawscrap-record-detail/update` | updateRawscrapRecordDetail | **制品返修记录子** |
| DELETE | `/wms/rawscrap-record-detail/delete` | deleteRawscrapRecordDetail | **制品返修记录子** |
| GET | `/wms/rawscrap-record-detail/get` | getRawscrapRecordDetail | **编号** |
| GET | `/wms/rawscrap-record-detail/list` | getRawscrapRecordDetailList | **编号** |
| POST | `/wms/rawscrap-record-detail/senior` | getRawscrapRecordDetailSenior | **编号列表** |
| GET | `/wms/rawscrap-record-detail/page` | getRawscrapRecordDetailPage | **制品返修记录子** |
| GET | `/wms/rawscrap-record-detail/export-excel` | exportRawscrapRecordDetailExcel | **制品返修记录子** |
| GET | `/wms/rawscrap-record-detail/get-import-template` | importTemplate | **制品返修记录子** |
| POST | `/wms/rawscrap-record-detail/import` | importExcel | **制品返修记录子** |

### 生产报废申请 (productscrapRequest)  (42个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/productscrap-request-detail/create` | createProductscrapRequestDetail | **制品报废申请子** |
| PUT | `/wms/productscrap-request-detail/update` | updateProductscrapRequestDetail | **制品报废申请子** |
| DELETE | `/wms/productscrap-request-detail/delete` | deleteProductscrapRequestDetail | **制品报废申请子** |
| GET | `/wms/productscrap-request-detail/get` | getProductscrapRequestDetail | **编号** |
| GET | `/wms/productscrap-request-detail/list` | getProductscrapRequestDetailList | **编号** |
| GET | `/wms/productscrap-request-detail/page` | getProductscrapRequestDetailPage | **编号列表** |
| POST | `/wms/productscrap-request-detail/senior` | getProductscrapRequestDetailSenior | **制品报废申请子** |
| GET | `/wms/productscrap-request-detail/export-excel` | exportProductscrapRequestDetailExcel | **制品报废申请子** |
| POST | `/wms/productscrap-request-main/create` | createProductscrapRequestMain | **制品报废申请主** |
| PUT | `/wms/productscrap-request-main/update` | updateProductscrapRequestMain | **制品报废申请主** |
| GET | `/wms/productscrap-request-main/editReturnNew` | editProductscrapRequestMainNew | **制品报废申请主** |
| DELETE | `/wms/productscrap-request-main/delete` | deleteProductscrapRequestMain | **制品报废申请主** |
| POST | `/wms/productscrap-request-main/senior` | getProductscrapRequestMainSenior | **编号** |
| GET | `/wms/productscrap-request-main/get` | getProductscrapRequestMain | **制品报废申请主** |
| GET | `/wms/productscrap-request-main/list` | getProductscrapRequestMainList | **编号** |
| GET | `/wms/productscrap-request-main/page` | getProductscrapRequestMainPage | **编号列表** |
| GET | `/wms/productscrap-request-main/pageBom` | getProductscrapRequestMainPageBom | **制品报废申请主** |
| POST | `/wms/productscrap-request-main/updateBomQty` | updateBomQty | **制品报废申请主** |
| GET | `/wms/productscrap-request-main/export-excel` | exportProductscrapRequestMainExcel | **制品报废申请主** |
| POST | `/wms/productscrap-request-main/export-excel-senior` | exportProductscrapRequestMainSeniorExcel | **制品报废申请主** |
| GET | `/wms/productscrap-request-main/get-import-template` | importTemplate | **制品报废申请主** |
| POST | `/wms/productscrap-request-main/import` | importExcel | **制品报废申请主** |
| PUT | `/wms/productscrap-request-main/close` | closeProductscrapRequestMain | **Excel 文件** |
| PUT | `/wms/productscrap-request-main/reAdd` | reAddeProductscrapRequestMain | **编号** |
| PUT | `/wms/productscrap-request-main/submit` | submitProductscrapRequestMain | **编号** |
| PUT | `/wms/productscrap-request-main/refused` | refusedProductscrapRequestMain | **编号** |
| PUT | `/wms/productscrap-request-main/agree` | agreeProductscrapRequestMain | **编号** |
| PUT | `/wms/productscrap-request-main/handle` | handleProductscrapRequestMain | **编号** |
| GET | `/wms/productscrap-request-main/bomPage` | getBomInfoPage | **编号** |
| GET | `/wms/productscrap-request-main/bomRecordPage` | getBomRecordInfoPage | **制品报废申请主** |
| GET | `/wms/productscrap-request-main/bomPageBatch` | getBomInfoPageBatch | **制品报废申请主** |
| GET | `/wms/productscrap-request-main/bomRecordPageBatch` | getBomRecordInfoPageBatch | **制品报废申请主** |
| POST | `/wms/productscrap-request-main/updateBom` | updateProductscrapDetailRequestBom | **制品报废申请主** |
| POST | `/wms/rawscrap-request-detail/create` | createRawscrapRequestDetail | **制品报废-原材料报废申请** |
| PUT | `/wms/rawscrap-request-detail/update` | updateRawscrapRequestDetail | **制品报废-原材料报废申请** |
| DELETE | `/wms/rawscrap-request-detail/delete` | deleteRawscrapRequestDetail | **制品报废-原材料报废申请** |
| GET | `/wms/rawscrap-request-detail/get` | getRawscrapRequestDetail | **编号** |
| GET | `/wms/rawscrap-request-detail/list` | getRawscrapRequestDetailList | **编号** |
| GET | `/wms/rawscrap-request-detail/page` | getRawscrapRequestDetailPage | **编号列表** |
| GET | `/wms/rawscrap-request-detail/export-excel` | exportRawscrapRequestDetailExcel | **制品报废-原材料报废申请** |
| GET | `/wms/rawscrap-request-detail/get-import-template` | importTemplate | **制品报废-原材料报废申请** |
| POST | `/wms/rawscrap-request-detail/import` | importExcel | **制品报废-原材料报废申请** |

### 项目 (project)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/project/create` | createProject | **项目** |
| PUT | `/wms/project/update` | updateProject | **项目** |
| DELETE | `/wms/project/delete` | deleteProject | **项目** |
| GET | `/wms/project/get` | getProject | **编号** |
| POST | `/wms/project/senior` | getProjectSenior | **编号** |
| GET | `/wms/project/page` | getProjectPage | **项目** |
| GET | `/wms/project/export-excel` | exportProjectExcel | **项目** |
| POST | `/wms/project/export-excel-senior` | exportProjectExcel | **项目** |
| GET | `/wms/project/get-import-template` | importTemplate | **项目** |
| POST | `/wms/project/import` | importExcel | **项目** |

### 采购 (purchase)  (38个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/purchase-detail/create` | createPurchaseDetail | **采购订单子** |
| PUT | `/wms/purchase-detail/update` | updatePurchaseDetail | **采购订单子** |
| DELETE | `/wms/purchase-detail/delete` | deletePurchaseDetail | **采购订单子** |
| GET | `/wms/purchase-detail/get` | getPurchaseDetail | **编号** |
| GET | `/wms/purchase-detail/list` | getPurchaseDetailList | **编号** |
| GET | `/wms/purchase-detail/page` | getPurchaseDetailPage | **编号列表** |
| POST | `/wms/purchase-detail/senior` | getPurchaseDetailSenior | **采购订单子** |
| GET | `/wms/purchase-detail/pageWMS` | getPurchaseDetailPageWMS | **采购订单子** |
| POST | `/wms/purchase-detail/seniorWMS` | getPurchaseDetailSeniorWMS | **采购订单子** |
| GET | `/wms/purchase-detail/purchasereceiptRequestPageWMS` | getPurchasereceiptRequestPageWMS | **采购订单子** |
| POST | `/wms/purchase-detail/purchasereceiptRequestSeniorWMS` | getpurchasereceiptRequestSeniorWMS | **采购订单子** |
| GET | `/wms/purchase-detail/pageWMS-Spare` | getPurchaseDetailPageWMSSpare | **采购订单子** |
| POST | `/wms/purchase-detail/seniorWMS-Spare` | getPurchaseDetailSeniorWMSSpare | **采购订单子** |
| GET | `/wms/purchase-detail/pageWMS-MOrderType` | getPurchaseDetailPageWMSMOrderType | **采购订单子** |
| POST | `/wms/purchase-detail/seniorWMS-MOrderType` | getPurchaseDetailSeniorWMSMOrderType | **采购订单子** |
| GET | `/wms/purchase-detail/export-excel` | exportPurchaseDetailExcel | **采购订单子** |
| GET | `/wms/purchase-detail/selectAll` | selectAll | **采购订单子** |
| GET | `/wms/purchase-detail/pageCheckData` | getPurchaseDetailPageCheckData | **采购订单子** |
| GET | `/wms/purchase-detail/pageM` | getPurchaseDetailMPage | **采购订单子** |
| POST | `/wms/purchase-detail/seniorM` | getPurchaseDetailseniorM | **采购订单子** |
| POST | `/wms/purchase-main/create` | createPurchaseMain | **采购订单主** |
| PUT | `/wms/purchase-main/update` | updatePurchaseMain | **采购订单主** |
| DELETE | `/wms/purchase-main/delete` | deletePurchaseMain | **采购订单主** |
| GET | `/wms/purchase-main/get` | getPurchaseMain | **编号** |
| GET | `/wms/purchase-main/list` | getPurchaseMainList | **编号** |
| GET | `/wms/purchase-main/page` | getPurchaseMainPage | **编号列表** |
| POST | `/wms/purchase-main/senior` | getPurchaseMainSenior | **采购订单主** |
| POST | `/wms/purchase-main/export-excel-senior1` | exportPurchaseMainSeniorExcel | **采购订单主** |
| GET | `/wms/purchase-main/get-import-template` | importTemplate | **采购订单主** |
| POST | `/wms/purchase-main/import` | importExcel | **采购订单主** |
| POST | `/wms/purchase-main/close` | closePurchaseMain | **Excel 文件** |
| POST | `/wms/purchase-main/open` | openPurchaseMain | **编号** |
| POST | `/wms/purchase-main/publish` | publishPurchaseMain | **编号** |
| POST | `/wms/purchase-main/wit` | witPurchaseMain | **编号** |
| GET | `/wms/purchase-main/export-excel` | exportPurchaseMainExcel1 | **编号** |
| GET | `/wms/purchase-main/export-excelWMS` | exportPurchaseMainExcelWMS | **采购订单主** |
| POST | `/wms/purchase-main/export-excel-senior` | exportPurchaseMainSeniorExcel1 | **采购订单主** |
| POST | `/wms/purchase-main/export-excel-seniorWMS` | exportPurchaseMainSeniorExcelWMS | **采购订单主** |

### 采购计划 (purchasePlan)  (40个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/purchase-plan-detail/create` | createPurchasePlanDetail | **要货计划子** |
| PUT | `/wms/purchase-plan-detail/update` | updatePurchasePlanDetail | **要货计划子** |
| GET | `/wms/purchase-plan-detail/page` | getPurchasePlanDetailPage | **要货计划子** |
| POST | `/wms/purchase-plan-detail/senior` | getPurchasePlanDetailSenior | **要货计划子** |
| DELETE | `/wms/purchase-plan-detail/delete` | deletePurchasePlanDetail | **要货计划子** |
| GET | `/wms/purchase-plan-detail/get` | getPurchasePlanDetail | **编号** |
| GET | `/wms/purchase-plan-detail/list` | getPurchasePlanDetailList | **编号** |
| GET | `/wms/purchase-plan-detail/export-excel` | exportPurchasePlanDetailExcel | **编号列表** |
| GET | `/wms/purchase-plan-detail/allList` | selectAllList | **要货计划子** |
| GET | `/wms/purchase-plan-detail/pageWMS` | getPurchasePlanDetailPagewms | **要货计划子** |
| POST | `/wms/purchase-plan-detail/seniorWMS` | getPurchasePlanDetailSeniorwms | **要货计划子** |
| GET | `/wms/purchase-plan-detail/clickDetailsPage` | clickDetailsPage | **要货计划子** |
| POST | `/wms/purchase-plan-detail/clickDetailsSenior` | clickDetailsSenior | **要货计划子** |
| POST | `/wms/purchase-plan-main/create` | createPurchasePlanMain | **要货计划主** |
| PUT | `/wms/purchase-plan-main/update` | updatePurchasePlanMain | **要货计划主** |
| DELETE | `/wms/purchase-plan-main/delete` | deletePurchasePlanMain | **要货计划主** |
| GET | `/wms/purchase-plan-main/get` | getPurchasePlanMain | **编号** |
| GET | `/wms/purchase-plan-main/list` | getPurchasePlanMainList | **编号** |
| GET | `/wms/purchase-plan-main/page` | getPurchasePlanMainPage | **编号列表** |
| POST | `/wms/purchase-plan-main/senior` | getPurchasePlanMainSenior | **要货计划主** |
| GET | `/wms/purchase-plan-main/export-excel-detail` | exportPurchasePlanMainExcelDetail | **要货计划主** |
| GET | `/wms/purchase-plan-main/export-excel` | exportPurchasePlanMainExcel | **要货计划主** |
| GET | `/wms/purchase-plan-main/export-excel-detail` | exportPurchasePlanMainExcelDetail | **要货计划主** |
| POST | `/wms/purchase-plan-main/export-excel-senior` | exportPurchasePlanMainSeniorExcel | **要货计划主** |
| GET | `/wms/purchase-plan-main/export-excel-senior` | exportPurchasePlanMainSeniorExcel | **要货计划主** |
| GET | `/wms/purchase-plan-main/get-import-template` | importTemplate | **要货计划主** |
| POST | `/wms/purchase-plan-main/import` | importExcel | **要货计划主** |
| POST | `/wms/purchase-plan-main/close` | closePurchasePlanMain | **Excel 文件** |
| GET | `/wms/purchase-plan-main/close-attachment-files` | listCloseAttachmentFiles | **要货计划主** |
| POST | `/wms/purchase-plan-main/open` | openPurchasePlanMain | **要货计划主编号** |
| POST | `/wms/purchase-plan-main/supplier-to-confirm` | supplierToConfirmPurchasePlanMain | **编号** |
| GET | `/wms/purchase-plan-main/supplier-confirm-detail` | getSupplierConfirmDetailList | **编号** |
| POST | `/wms/purchase-plan-main/supplier-confirm` | supplierConfirmPurchasePlanMain | **编号** |
| POST | `/wms/purchase-plan-main/planer-confirm` | planerConfirmPurchasePlanMain | **要货计划主** |
| POST | `/wms/purchase-plan-main/publish` | publishPurchasePlanMain | **编号** |
| POST | `/wms/purchase-plan-main/wit` | witPurchasePlanMain | **编号** |
| POST | `/wms/purchase-plan-main/acc` | accPurchasePlanMain | **编号** |
| POST | `/wms/purchase-plan-main/rej` | rejPurchasePlanMain | **编号** |
| GET | `/wms/purchase-plan-main/queryPurchasePlan` | queryPurchasePlan | **编号** |
| PUT | `/wms/purchase-plan-main/updateALL` | updateALLPurchasePlanMain | **编号** |

### 采购领料记录 (purchaseclaimRecord)  (17个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/purchaseclaim-record-detail/create` | createPurchaseclaimRecordDetail | **采购索赔记录子** |
| PUT | `/wms/purchaseclaim-record-detail/update` | updatePurchaseclaimRecordDetail | **采购索赔记录子** |
| DELETE | `/wms/purchaseclaim-record-detail/delete` | deletePurchaseclaimRecordDetail | **采购索赔记录子** |
| GET | `/wms/purchaseclaim-record-detail/get` | getPurchaseclaimRecordDetail | **编号** |
| GET | `/wms/purchaseclaim-record-detail/list` | getPurchaseclaimRecordDetailList | **编号** |
| GET | `/wms/purchaseclaim-record-detail/page` | getPurchaseclaimRecordDetailPage | **编号列表** |
| POST | `/wms/purchaseclaim-record-detail/senior` | getPurchaseclaimRecordDetailSenior | **采购索赔记录子** |
| GET | `/wms/purchaseclaim-record-detail/export-excel` | exportPurchaseclaimRecordDetailExcel | **采购索赔记录子** |
| POST | `/wms/purchaseclaim-record-main/create` | createPurchaseclaimRecordMain | **采购索赔记录主** |
| PUT | `/wms/purchaseclaim-record-main/update` | updatePurchaseclaimRecordMain | **采购索赔记录主** |
| DELETE | `/wms/purchaseclaim-record-main/delete` | deletePurchaseclaimRecordMain | **采购索赔记录主** |
| GET | `/wms/purchaseclaim-record-main/get` | getPurchaseclaimRecordMain | **编号** |
| GET | `/wms/purchaseclaim-record-main/list` | getPurchaseclaimRecordMainList | **编号** |
| GET | `/wms/purchaseclaim-record-main/page` | getPurchaseclaimRecordMainPage | **编号列表** |
| POST | `/wms/purchaseclaim-record-main/senior` | getPurchaseclaimRecordMainSenior | **采购索赔记录主** |
| GET | `/wms/purchaseclaim-record-main/export-excel` | exportPurchaseclaimRecordMainExcel | **采购索赔记录主** |
| POST | `/wms/purchaseclaim-record-main/export-excel-senior` | exportPurchaseclaimRecordMainSeniorExcel | **采购索赔记录主** |

### 采购领料申请 (purchaseclaimRequest)  (22个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/purchaseclaim-request-detail/create` | createPurchaseclaimRequestDetail | **采购索赔申请子** |
| PUT | `/wms/purchaseclaim-request-detail/update` | updatePurchaseclaimRequestDetail | **采购索赔申请子** |
| DELETE | `/wms/purchaseclaim-request-detail/delete` | deletePurchaseclaimRequestDetail | **采购索赔申请子** |
| GET | `/wms/purchaseclaim-request-detail/get` | getPurchaseclaimRequestDetail | **编号** |
| GET | `/wms/purchaseclaim-request-detail/page` | getPurchaseclaimRequestDetailPage | **编号** |
| POST | `/wms/purchaseclaim-request-detail/senior` | getPurchaseclaimRequestDetailSenior | **采购索赔申请子** |
| POST | `/wms/purchaseclaim-request-main/create` | createPurchaseclaimRequestMain | **采购索赔申请主** |
| PUT | `/wms/purchaseclaim-request-main/update` | updatePurchaseclaimRequestMain | **采购索赔申请主** |
| DELETE | `/wms/purchaseclaim-request-main/delete` | deletePurchaseclaimRequestMain | **采购索赔申请主** |
| GET | `/wms/purchaseclaim-request-main/get` | getPurchaseclaimRequestMain | **编号** |
| GET | `/wms/purchaseclaim-request-main/page` | getPurchaseclaimRequestMainPage | **编号** |
| POST | `/wms/purchaseclaim-request-main/senior` | getPurchaseclaimRequestMainSenior | **采购索赔申请主** |
| GET | `/wms/purchaseclaim-request-main/export-excel` | exportPurchaseclaimRequestMainExcel | **采购索赔申请主** |
| POST | `/wms/purchaseclaim-request-main/export-excel-senior` | exportPurchaseclaimRequestMainSeniorExcel | **采购索赔申请主** |
| GET | `/wms/purchaseclaim-request-main/get-import-template` | importTemplate | **采购索赔申请主** |
| POST | `/wms/purchaseclaim-request-main/import` | importExcel | **采购索赔申请主** |
| POST | `/wms/purchaseclaim-request-main/close` | closePurchaseclaimRequestMain | **Excel 文件** |
| POST | `/wms/purchaseclaim-request-main/open` | openPurchaseclaimRequestMain | **编号** |
| POST | `/wms/purchaseclaim-request-main/sub` | subPurchaseclaimRequestMain | **编号** |
| POST | `/wms/purchaseclaim-request-main/app` | witPurchaseclaimRequestMain | **编号** |
| POST | `/wms/purchaseclaim-request-main/rej` | rejPurchaseclaimRequestMain | **编号** |
| POST | `/wms/purchaseclaim-request-main/genRecords` | genRecords | **编号** |

### 采购MRS统计 (purchasemrsstatistics)  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/purchase-mrs-statistics/create` | createPurchaseMrsStatistics | **要货计划汇总统计给MRS** |
| PUT | `/wms/purchase-mrs-statistics/update` | updatePurchaseMrsStatistics | **要货计划汇总统计给MRS** |
| DELETE | `/wms/purchase-mrs-statistics/delete` | deletePurchaseMrsStatistics | **要货计划汇总统计给MRS** |
| GET | `/wms/purchase-mrs-statistics/get` | getPurchaseMrsStatistics | **编号** |
| GET | `/wms/purchase-mrs-statistics/page` | getPurchaseMrsStatisticsPage | **编号** |
| POST | `/wms/purchase-mrs-statistics/senior` | getPurchaseMrsStatisticsSenior | **要货计划汇总统计给MRS** |
| GET | `/wms/purchase-mrs-statistics/export-excel` | exportPurchaseMrsStatisticsExcel | **要货计划汇总统计给MRS** |
| GET | `/wms/purchase-mrs-statistics/get-import-template` | importTemplate | **要货计划汇总统计给MRS** |
| POST | `/wms/purchase-mrs-statistics/import` | importExcel | **要货计划汇总统计给MRS** |

### 采购价格 (purchaseprice)  (14个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/purchaseprice/create` | createPurchaseprice | **采购价格单** |
| PUT | `/wms/purchaseprice/update` | updatePurchaseprice | **采购价格单** |
| DELETE | `/wms/purchaseprice/delete` | deletePurchaseprice | **采购价格单** |
| GET | `/wms/purchaseprice/get` | getPurchaseprice | **编号** |
| GET | `/wms/purchaseprice/page` | getPurchasepricePage | **编号** |
| POST | `/wms/purchaseprice/senior` | getPurchasepriceSenior | **采购价格单** |
| GET | `/wms/purchaseprice/pageSCP` | getPurchasepricePageSCP | **采购价格单** |
| POST | `/wms/purchaseprice/seniorSCP` | getPurchasepriceSeniorSCP | **采购价格单** |
| GET | `/wms/purchaseprice/export-excel` | exportPurchasepriceExcel | **采购价格单** |
| POST | `/wms/purchaseprice/export-excel-senior` | exportItembasicExcel | **采购价格单** |
| GET | `/wms/purchaseprice/export-excel-SCP` | exportPurchasepriceExcelSCP | **采购价格单** |
| POST | `/wms/purchaseprice/export-excel-senior-SCP` | exportItembasicExcelSCP | **采购价格单** |
| GET | `/wms/purchaseprice/get-import-template` | importTemplate | **采购价格单** |
| POST | `/wms/purchaseprice/import` | importExcel | **采购价格单** |

### 采购入库任务 (purchasereceiptJob)  (25个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/purchasereceipt-job-detail/page` | getPurchasereceiptJobDetailPage | **采购收货任务子** |
| POST | `/wms/purchasereceipt-job-detail/senior` | getPurchasereceiptJobDetailSenior | **采购收货任务子** |
| GET | `/wms/purchasereceipt-job-detail/pageSpare` | ggetPurchasereceiptJobDetailPageSpare | **采购收货任务子** |
| POST | `/wms/purchasereceipt-job-detail/seniorSpare` | getPurchasereceiptJobDetailSeniorSpare | **采购收货任务子** |
| GET | `/wms/purchasereceipt-job-detail/pageSCP` | getPurchasereceiptJobDetailPageSCP | **采购收货任务子** |
| POST | `/wms/purchasereceipt-job-detail/seniorSCP` | getPurchasereceiptJobDetailSeniorSCP | **采购收货任务子** |
| GET | `/wms/purchasereceipt-job-detail/pagePackingNumber` | getPurchasereceiptJobDetailPagePackingNumber | **采购收货任务子** |
| POST | `/wms/purchasereceipt-job-detail/seniorPackingNumber` | getPurchasereceiptJobDetailSeniorPackingNumber | **采购收货任务子** |
| GET | `/wms/purchasereceipt-job-main/page` | getPurchasereceiptJobMainPage | **采购收货任务主** |
| POST | `/wms/purchasereceipt-job-main/senior` | getPurchasereceiptJobMainSenior | **采购收货任务主** |
| GET | `/wms/purchasereceipt-job-main/export-excel` | exportPurchasereceiptJobMainExcel | **采购收货任务主** |
| POST | `/wms/purchasereceipt-job-main/export-excel-senior` | exportPurchasereceiptJobMainSeniorExcel | **采购收货任务主** |
| GET | `/wms/purchasereceipt-job-main/export-excel-spare` | exportSpareJobMainExcel | **采购收货任务主** |
| POST | `/wms/purchasereceipt-job-main/export-excel-spare-senior` | exportSpareJobMainSeniorExcel | **采购收货任务主** |
| GET | `/wms/purchasereceipt-job-main/getPurchasereceiptJobyId` | getPurchasereceiptJobyId | **采购收货任务主** |
| POST | `/wms/purchasereceipt-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/purchasereceipt-job-main/accept` | acceptPurchasereceiptJobMain | **类型数组** |
| PUT | `/wms/purchasereceipt-job-main/abandon` | abandonPurchasereceiptJobMain | **采购收货任务主** |
| PUT | `/wms/purchasereceipt-job-main/close` | closePurchasereceiptJobMain | **采购收货任务主** |
| PUT | `/wms/purchasereceipt-job-main/execute` | executePurchasereceiptJobMain | **采购收货任务主** |
| PUT | `/wms/purchasereceipt-job-main/executeSpare` | executePurchasereceiptJobMainSpare | **编号** |
| POST | `/wms/purchasereceipt-job-main/refusal` | refusalPurchasereceiptJobMain | **编号** |
| GET | `/wms/purchasereceipt-job-main/queryInspectionFreeFlag` | getPurchasereceiptJobyId | **采购收货任务主** |
| PUT | `/wms/purchasereceipt-job-main/updatePurchasereceiptJobConfig` | updatePurchasereceiptJobConfig | **编号** |
| PUT | `/wms/purchasereceipt-job-main/updatePurchasereceiptJobConfigSpare` | updatePurchasereceiptJobConfigSpare | **采购收货任务主** |

### 采购入库记录 (purchasereceiptRecord)  (36个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/purchasereceipt-record-detail/page` | getPurchasereceiptRecordDetailPage | **采购收货记录子** |
| POST | `/wms/purchasereceipt-record-detail/senior` | getPurchasereceiptRecordDetailSenior | **采购收货记录子** |
| GET | `/wms/purchasereceipt-record-detail/pageRefuse` | getPurchasereceiptRecordDetailPageRefuse | **采购收货记录子** |
| POST | `/wms/purchasereceipt-record-detail/seniorRefuse` | getPurchasereceiptRecordDetailSeniorRefuse | **采购收货记录子** |
| GET | `/wms/purchasereceipt-record-detail/pageReturn` | getPurchasereceiptRecordDetailPageReturn | **采购收货记录子** |
| POST | `/wms/purchasereceipt-record-detail/seniorReturn` | getPurchasereceiptRecordDetailSeniorReturn | **采购收货记录子** |
| GET | `/wms/purchasereceipt-record-detail/pageSpare` | getPurchasereceiptRecordDetailPageSpare | **采购收货记录子** |
| POST | `/wms/purchasereceipt-record-detail/seniorSpare` | getPurchasereceiptRecordDetailSeniorSpare | **采购收货记录子** |
| GET | `/wms/purchasereceipt-record-detail/pageSCP` | getPurchasereceiptRecordDetailPageSCP | **采购收货记录子** |
| POST | `/wms/purchasereceipt-record-detail/seniorSCP` | getPurchasereceiptRecordDetailSeniorSCP | **采购收货记录子** |
| GET | `/wms/purchasereceipt-record-detail/pageBySupplierCode` | getPurchasereceiptRecordDetailPageBySupplierCode | **采购收货记录子** |
| POST | `/wms/purchasereceipt-record-detail/seniorBySupplierCode` | getPurchasereceiptRecordDetailSeniorBySupplierCode | **采购收货记录子** |
| GET | `/wms/purchasereceipt-record-detail/list` | getPurchasereceiptRecordDetailList | **采购收货记录子** |
| GET | `/wms/purchasereceipt-record-detail/listToRepeatPurchaseReceipt` | getPurchasereceiptRecordDetailListToRepeat | **采购收货记录子** |
| GET | `/wms/purchasereceipt-record-detail/pageChildPackingNumber` | getPurchasereceiptRecordDetailPageChildPackingNumber | **采购收货记录子** |
| GET | `/wms/purchasereceipt-record-detail/queryPurchasereceiptRecordByItemCode` | queryPurchasereceiptRecordByItemCode | **采购收货记录子** |
| GET | `/wms/purchasereceipt-record-detail/queryPurchasereceiptSpareRecordBySupplier` | queryPurchasereceiptSpareRecordBySupplier | **采购收货记录子** |
| GET | `/wms/purchasereceipt-record-detail/pageForQ2` | pageForQ2 | **采购收货记录子** |
| POST | `/wms/purchasereceipt-record-main/createMq` | createPurchasereceiptRecordMainMq | **采购收货记录主** |
| POST | `/wms/purchasereceipt-record-main/create` | createPurchasereceiptRecordMain | **采购收货记录主** |
| GET | `/wms/purchasereceipt-record-main/page` | getPurchasereceiptRecordMainPage | **采购收货记录主** |
| POST | `/wms/purchasereceipt-record-main/senior` | getPurchasereceiptRecordMainSenior | **采购收货记录主** |
| GET | `/wms/purchasereceipt-record-main/export-excel` | exportPurchasereceiptRecordMainExcel | **采购收货记录主** |
| POST | `/wms/purchasereceipt-record-main/export-excel-senior` | exportPurchasereceiptRecordMainSeniorExcel | **采购收货记录主** |
| GET | `/wms/purchasereceipt-record-main/export-excel-refuse` | exportPurchasereceiptRecordMainExcelRefuse | **采购收货记录主** |
| POST | `/wms/purchasereceipt-record-main/export-excel-senior-refuse` | exportPurchasereceiptRecordMainSeniorExcelRefuse | **采购收货记录主** |
| GET | `/wms/purchasereceipt-record-main/export-excel-spare` | exportSpareRecordMainExcel | **采购收货记录主** |
| POST | `/wms/purchasereceipt-record-main/export-excel-spare-senior` | exportSpareRecordMainSeniorExcel | **采购收货记录主** |
| GET | `/wms/purchasereceipt-record-main/export-excel-SCP` | exportPurchasereceiptRecordMainExcelSCP | **采购收货记录主** |
| POST | `/wms/purchasereceipt-record-main/export-excel-senior-SCP` | exportPurchasereceiptRecordMainSeniorExcelSCP | **采购收货记录主** |
| POST | `/wms/purchasereceipt-record-main/createInspectRequest` | createInspectRequest | **采购收货记录主** |
| POST | `/wms/purchasereceipt-record-main/createPutawayRequest` | createPutawayRequest | **采购收货记录主** |
| POST | `/wms/purchasereceipt-record-main/createPurchasereturnRecord` | createPurchasereturnRecord | **采购收货记录主** |
| GET | `/wms/purchaseshortage-detail/page` | getPurchaseshortageDetailPage | **采购收货缺货子** |
| POST | `/wms/purchaseshortage-detail/senior` | getPurchaseshortageDetailSenior | **采购收货缺货子** |
| GET | `/wms/purchaseshortage-detail/pageChildPackingNumber` | getPurchaseshortageDetailPagePackingNumber | **采购收货缺货子** |

### purchasereceiptRequest  (38个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/purchasereceipt-request-detail/create` | createPurchasereceiptRequestDetail | **采购收货申请子** |
| PUT | `/wms/purchasereceipt-request-detail/update` | updatePurchasereceiptRequestDetail | **采购收货申请子** |
| DELETE | `/wms/purchasereceipt-request-detail/delete` | deletePurchasereceiptRequestDetail | **采购收货申请子** |
| GET | `/wms/purchasereceipt-request-detail/get` | getPurchasereceiptRequestDetail | **编号** |
| GET | `/wms/purchasereceipt-request-detail/page` | getPurchasereceiptRequestDetailPage | **编号** |
| POST | `/wms/purchasereceipt-request-detail/senior` | getPurchasereceiptRequestDetailSenior | **采购收货申请子** |
| GET | `/wms/purchasereceipt-request-detail/pageLabel` | getPurchasereceiptRequestDetailPageLabel | **采购收货申请子** |
| GET | `/wms/purchasereceipt-request-detail/pageOrderMType` | getPurchasereceiptRequestDetailPageOrderMType | **采购收货申请子** |
| POST | `/wms/purchasereceipt-request-detail/seniorOrderMType` | getPurchasereceiptRequestDetailSeniorOrderMType | **采购收货申请子** |
| GET | `/wms/purchasereceipt-request-detail/pageSpare` | getPurchasereceiptRequestDetailPageSpare | **采购收货申请子** |
| POST | `/wms/purchasereceipt-request-detail/seniorSpare` | getPurchasereceiptRequestDetailSeniorSpare | **采购收货申请子** |
| GET | `/wms/purchasereceipt-request-detail/queryPurchaseceiptChildPackingNumberPage` | queryPurchaseceiptChildPackingNumberPage | **采购收货申请子** |
| POST | `/wms/purchasereceipt-request-detail/queryPurchaseceiptChildPackingNumberSenior` | queryPurchaseceiptChildPackingNumberSenior | **采购收货申请子** |
| POST | `/wms/purchasereceipt-request-main/create` | createPurchasereceiptRequestMain | **采购收货申请主** |
| POST | `/wms/purchasereceipt-request-main/createSpare` | createPurchasereceiptRequestMainSpare | **采购收货申请主** |
| PUT | `/wms/purchasereceipt-request-main/update` | updatePurchasereceiptRequestMain | **采购收货申请主** |
| GET | `/wms/purchasereceipt-request-main/get` | getPurchasereceiptRequestMain | **采购收货申请主** |
| GET | `/wms/purchasereceipt-request-main/list` | getPurchasereceiptRequestMainList | **编号** |
| GET | `/wms/purchasereceipt-request-main/page` | getPurchasereceiptRequestMainPage | **编号列表** |
| POST | `/wms/purchasereceipt-request-main/senior` | getPurchasereceiptRequestMainSenior | **采购收货申请主** |
| GET | `/wms/purchasereceipt-request-main/export-excel` | exportPurchasereceiptRequestMainExcel | **采购收货申请主** |
| POST | `/wms/purchasereceipt-request-main/export-excel-senior` | exportPurchasereceiptRequestMainSeniorExcel | **采购收货申请主** |
| GET | `/wms/purchasereceipt-request-main/export-excel-orderTypeM` | exportPurchasereceiptRequestMainExcelOrderTypeM | **采购收货申请主** |
| POST | `/wms/purchasereceipt-request-main/export-excel-senior-orderTypeM` | exportPurchasereceiptRequestMainSeniorExcelOrderTypeM | **采购收货申请主** |
| GET | `/wms/purchasereceipt-request-main/export-excel-spare` | exportSpareRequestMainExcel | **采购收货申请主** |
| POST | `/wms/purchasereceipt-request-main/export-excel-spare-senior` | exportSpareRequestMainSeniorExcel | **采购收货申请主** |
| GET | `/wms/purchasereceipt-request-main/get-import-template` | importTemplate | **采购收货申请主** |
| POST | `/wms/purchasereceipt-request-main/import` | importExcel | **采购收货申请主** |
| POST | `/wms/purchasereceipt-request-main/importSpare` | importExcelSpare | **Excel 文件** |
| PUT | `/wms/purchasereceipt-request-main/close` | closePurchasereceiptRequestMain | **Excel 文件** |
| PUT | `/wms/purchasereceipt-request-main/reAdd` | openPurchasereceiptRequestMain | **编号** |
| PUT | `/wms/purchasereceipt-request-main/submit` | submitPurchasereceiptRequestMain | **编号** |
| PUT | `/wms/purchasereceipt-request-main/agree` | agreePurchasereceiptRequestMain | **编号** |
| PUT | `/wms/purchasereceipt-request-main/handle` | handlePurchasereceiptRequestMain | **编号** |
| PUT | `/wms/purchasereceipt-request-main/refused` | abortPurchasereceiptRequestMain | **编号** |
| POST | `/wms/purchasereceipt-request-main/genLabel` | genLabel | **编号** |
| POST | `/wms/purchasereceipt-request-main/deleteOldLabels` | deleteOldLabels | **编号** |
| POST | `/wms/purchasereceipt-request-main/queryPurchasePlan` | queryPurchasePlan | **编号** |

### 采购退货任务 (purchasereturnJob)  (13个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/purchasereturn-job-detail/page` | getPurchasereturnJobDetailPage | **采购退货任务子** |
| POST | `/wms/purchasereturn-job-detail/senior` | getPurchasereturnJobDetailSenior | **采购退货任务子** |
| GET | `/wms/purchasereturn-job-main/page` | getPurchasereturnJobMainPage | **采购退货任务主** |
| POST | `/wms/purchasereturn-job-main/senior` | getPurchasereturnJobMainSenior | **采购退货任务主** |
| GET | `/wms/purchasereturn-job-main/export-excel` | exportPurchasereturnJobMainExcel | **采购退货任务主** |
| POST | `/wms/purchasereturn-job-main/export-excel-senior` | exportPurchasereturnJobMainSeniorExcel | **采购退货任务主** |
| GET | `/wms/purchasereturn-job-main/getReturnJobById` | getReturnJobById | **采购退货任务主** |
| GET | `/wms/purchasereturn-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/purchasereturn-job-main/accept` | acceptPurchasereceiptJobMain | **类型数组** |
| PUT | `/wms/purchasereturn-job-main/abandon` | abandonPurchasereceiptJobMain | **采购退货任务主** |
| PUT | `/wms/purchasereturn-job-main/close` | closePurchasereceiptJobMain | **采购退货任务主** |
| PUT | `/wms/purchasereturn-job-main/updatePurchasereturnJobConfig` | updatePurchasereturnJobConfig | **采购退货任务主** |
| PUT | `/wms/purchasereturn-job-main/execute` | executePurchasereceiptJobMain | **采购退货任务主** |

### 采购退货记录 (purchasereturnRecord)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/purchasereturn-record-detail/page` | getPurchasereturnRecordDetailPage | **采购退货记录子** |
| POST | `/wms/purchasereturn-record-detail/senior` | getPurchasereturnRecordDetailSenior | **采购退货记录子** |
| GET | `/wms/purchasereturn-record-detail/pageSpare` | getPurchasereturnRecordDetailPageSpare | **采购退货记录子** |
| POST | `/wms/purchasereturn-record-detail/seniorSpare` | getPurchasereturnRecordDetailSeniorSpare | **采购退货记录子** |
| GET | `/wms/purchasereturn-record-detail/pageMorderType` | getPurchasereturnRecordDetailPageMorderType | **采购退货记录子** |
| POST | `/wms/purchasereturn-record-detail/seniorMorderType` | getPurchasereturnRecordDetailSeniorMorderType | **采购退货记录子** |
| GET | `/wms/purchasereturn-record-detail/pageSCP` | getPurchasereturnRecordDetailPageSCP | **采购退货记录子** |
| POST | `/wms/purchasereturn-record-detail/seniorSCP` | getPurchasereturnRecordDetailSeniorSCP | **采购退货记录子** |
| POST | `/wms/purchasereturn-record-main/create` | createPurchasereturnRecordMain | **采购退货记录主** |
| GET | `/wms/purchasereturn-record-main/page` | getPurchasereturnRecordMainPage | **采购退货记录主** |
| POST | `/wms/purchasereturn-record-main/senior` | getPurchasereturnRecordMainSenior | **采购退货记录主** |
| GET | `/wms/purchasereturn-record-main/export-excel` | exportPurchasereturnRecordMainExcel | **采购退货记录主** |
| POST | `/wms/purchasereturn-record-main/export-excel-senior` | exportPurchasereturnRecordMainSeniorExcel | **采购退货记录主** |
| GET | `/wms/purchasereturn-record-main/export-excel-spare` | exportPurchasereturnRecordMainExcelSpare | **采购退货记录主** |
| POST | `/wms/purchasereturn-record-main/export-excel-senior-spare` | exportPurchasereturnRecordMainSeniorExcelSpare | **采购退货记录主** |
| GET | `/wms/purchasereturn-record-main/export-excel-mordertype` | exportPurchasereturnRecordMainExcelMorderType | **采购退货记录主** |
| POST | `/wms/purchasereturn-record-main/export-excel-senior-mordertype` | exportPurchasereturnRecordMainSeniorExcelMorderType | **采购退货记录主** |
| GET | `/wms/purchasereturn-record-main/export-excel-SCP` | exportPurchasereturnRecordMainExcelSCP | **采购退货记录主** |
| POST | `/wms/purchasereturn-record-main/export-excel-senior-SCP` | exportPurchasereturnRecordMainSeniorExcelSCP | **采购退货记录主** |

### 采购退货申请 (purchasereturnRequest)  (47个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/purchasereturn-request-detail/create` | createPurchasereturnRequestDetail | **采购退货申请子** |
| PUT | `/wms/purchasereturn-request-detail/update` | updatePurchasereturnRequestDetail | **采购退货申请子** |
| DELETE | `/wms/purchasereturn-request-detail/delete` | deletePurchasereturnRequestDetail | **采购退货申请子** |
| GET | `/wms/purchasereturn-request-detail/get` | getPurchasereturnRequestDetail | **编号** |
| GET | `/wms/purchasereturn-request-detail/page` | getPurchasereturnRequestDetailPage | **编号** |
| POST | `/wms/purchasereturn-request-detail/senior` | getPurchasereturnRequestDetailSenior | **采购退货申请子** |
| GET | `/wms/purchasereturn-request-detail/pageSpare` | getPurchasereturnRequestDetailPageSpare | **采购退货申请子** |
| POST | `/wms/purchasereturn-request-detail/seniorSpare` | getPurchasereturnRequestDetailSeniorSpare | **采购退货申请子** |
| GET | `/wms/purchasereturn-request-detail/pageMorderType` | getPurchasereturnRequestDetailPageMorderType | **采购退货申请子** |
| POST | `/wms/purchasereturn-request-detail/seniorMorderType` | getPurchasereturnRequestDetailSeniorMorderType | **采购退货申请子** |
| GET | `/wms/purchasereturn-request-detail/getReturnDetailList` | getReturnDetailList | **采购退货申请子** |
| POST | `/wms/purchasereturn-request-main/create` | createPurchasereturnRequestMain | **采购退货申请主** |
| PUT | `/wms/purchasereturn-request-main/update` | updatePurchasereturnRequestMain | **采购退货申请主** |
| POST | `/wms/purchasereturn-request-main/createReturnNew` | createPurchasereturnRequestMainNew | **采购退货申请主** |
| GET | `/wms/purchasereturn-request-main/editReturnNew` | editPurchasereturnRequestMainNew | **采购退货申请主** |
| POST | `/wms/purchasereturn-request-main/updateReturnNew` | updatePurchasereturnRequestMainNew | **采购退货申请主** |
| GET | `/wms/purchasereturn-request-main/editReturnSpareNew` | editPurchasereturnRequestSpareMainNew | **采购退货申请主** |
| POST | `/wms/purchasereturn-request-main/updateReturnSpareNew` | updatePurchasereturnRequestSpareMainNew | **采购退货申请主** |
| POST | `/wms/purchasereturn-request-main/createReturnNewMtype` | createPurchasereturnRequestMainNewMtype | **采购退货申请主** |
| GET | `/wms/purchasereturn-request-main/editReturnNewMtype` | editPurchasereturnRequestMainNewMtype | **采购退货申请主** |
| POST | `/wms/purchasereturn-request-main/updateReturnNewMtype` | updatePurchasereturnRequestMainNewMtype | **采购退货申请主** |
| GET | `/wms/purchasereturn-request-main/page` | getPurchasereturnRequestMainPage | **采购退货申请主** |
| POST | `/wms/purchasereturn-request-main/senior` | getPurchasereturnRequestMainSenior | **采购退货申请主** |
| GET | `/wms/purchasereturn-request-main/export-excel` | exportPurchasereturnRequestMainExcel | **采购退货申请主** |
| POST | `/wms/purchasereturn-request-main/export-excel-senior` | exportPurchasereturnRequestMainSeniorExcel | **采购退货申请主** |
| GET | `/wms/purchasereturn-request-main/export-excel-mordertype` | exportPurchasereturnRequestMainExcelMorderType | **采购退货申请主** |
| POST | `/wms/purchasereturn-request-main/export-excel-senior-mordertype` | exportPurchasereturnRequestMainSeniorExcelMorderType | **采购退货申请主** |
| GET | `/wms/purchasereturn-request-main/export-excel-spare` | exportPurchasereturnRequestMainExcelSpare | **采购退货申请主** |
| POST | `/wms/purchasereturn-request-main/export-excel-senior-spare` | exportPurchasereturnRequestMainSeniorExcelSpare | **采购退货申请主** |
| GET | `/wms/purchasereturn-request-main/get-import-template` | importTemplate | **采购退货申请主** |
| GET | `/wms/purchasereturn-request-main/get-import-template-mordertype` | importTemplateMordertype | **采购退货申请主** |
| GET | `/wms/purchasereturn-request-main/get-import-template-spare` | importTemplateSpare | **采购退货申请主** |
| GET | `/wms/purchasereturn-request-main/get-import-template-new` | importTemplateNew | **采购退货申请主** |
| POST | `/wms/purchasereturn-request-main/import` | importExcel | **采购退货申请主** |
| POST | `/wms/purchasereturn-request-main/importSpare` | importExcelSpare | **Excel 文件** |
| POST | `/wms/purchasereturn-request-main/importNew` | importExcelNew | **Excel 文件** |
| POST | `/wms/purchasereturn-request-main/importMorderType` | importExcelMorderType | **Excel 文件** |
| GET | `/wms/purchasereturn-request-main/getPurchasereturnRequestById` | getPurchasereturnRequestById | **Excel 文件** |
| GET | `/wms/purchasereturn-request-main/queryBalancePurchaseReceiptSpareReturn` | queryBalancePurchaseReceiptSpareReturn | **编号** |
| PUT | `/wms/purchasereturn-request-main/close` | closePurchasereturnRequestMain | **采购退货申请主** |
| PUT | `/wms/purchasereturn-request-main/reAdd` | openPurchasereturnRequestMain | **编号** |
| PUT | `/wms/purchasereturn-request-main/submit` | submitPurchasereturnRequestMain | **编号** |
| PUT | `/wms/purchasereturn-request-main/agree` | agreePurchasereturnRequestMain | **编号** |
| PUT | `/wms/purchasereturn-request-main/handle` | handlePurchasereturnRequestMain | **编号** |
| PUT | `/wms/purchasereturn-request-main/handleNew` | handlePurchasereturnRequestMainNew | **编号** |
| PUT | `/wms/purchasereturn-request-main/refused` | abortPurchasereturnRequestMain | **编号** |
| POST | `/wms/purchasereturn-request-main/genLabel` | genLabel | **编号** |

### 上架任务 (putawayJob)  (18个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/putaway-job-detail/page` | getPutawayJobDetailPage | **上架任务子** |
| POST | `/wms/putaway-job-detail/senior` | getPutawayJobDetailSenior | **上架任务子** |
| GET | `/wms/putaway-job-detail/pageChildPackingNumber` | getPutawayJobDetailPageChildPackingNumber | **上架任务子** |
| GET | `/wms/putaway-job-main/page` | getPutawayJobMainPage | **上架任务主** |
| POST | `/wms/putaway-job-main/senior` | getPutawayJobMainSenior | **上架任务主** |
| GET | `/wms/putaway-job-main/export-excel` | exportPutawayJobMainExcel | **上架任务主** |
| POST | `/wms/putaway-job-main/export-excel-senior` | exportPutawayJobMainSeniorExcel | **上架任务主** |
| GET | `/wms/putaway-job-main/getPutawayJobById` | getPutawayJobById | **上架任务主** |
| POST | `/wms/putaway-job-main/getPutawayJobPageByStatusAndTime` | getPutawayJobPageByStatusAndTime | **编号** |
| POST | `/wms/putaway-job-main/getCountByStatus` | getCountByStatus | **今日开始结束时间** |
| PUT | `/wms/putaway-job-main/accept` | acceptPutawayJobMain | **上架任务主** |
| PUT | `/wms/putaway-job-main/abandon` | cancelAcceptPutawayJobMain | **上架任务主** |
| PUT | `/wms/putaway-job-main/close` | closePutawayJobMain | **上架任务主** |
| PUT | `/wms/putaway-job-main/updatePutawayJobConfig` | updatePutawayJobConfig | **上架任务主** |
| PUT | `/wms/putaway-job-main/execute` | executePutawayJobMain | **上架任务主** |
| PUT | `/wms/putaway-job-main/acceptBatch` | acceptPutawayJobMainBatch | **编号** |
| PUT | `/wms/putaway-job-main/abandonBatch` | cancelAcceptPutawayJobMainBatch | **编号** |
| PUT | `/wms/putaway-job-main/executeBatch` | executePutawayJobMainBatch | **上架任务主** |

### 上架记录 (putawayRecord)  (8个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/putaway-record-detail/page` | getPutawayRecordDetailPage | **上架记录子** |
| POST | `/wms/putaway-record-detail/senior` | getPutawayRecordDetailSenior | **上架记录子** |
| GET | `/wms/putaway-record-detail/pageChildPackingNumber` | getPutawayRecordDetailPageChildPackingNumber | **上架记录子** |
| POST | `/wms/putaway-record-main/create` | createPutawayRecordMain | **上架记录主** |
| GET | `/wms/putaway-record-main/page` | getPutawayRecordMainPage | **上架记录主** |
| POST | `/wms/putaway-record-main/senior` | getPutawayRecordMainSenior | **上架记录主** |
| GET | `/wms/putaway-record-main/export-excel` | exportPutawayRecordMainExcel | **上架记录主** |
| POST | `/wms/putaway-record-main/export-excel-senior` | exportPutawayRecordMainSeniorExcel | **上架记录主** |

### 上架申请 (putawayRequest)  (24个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/putaway-request-detail/create` | createPutawayRequestDetail | **上架申请子** |
| PUT | `/wms/putaway-request-detail/update` | updatePutawayRequestDetail | **上架申请子** |
| DELETE | `/wms/putaway-request-detail/delete` | deletePutawayRequestDetail | **上架申请子** |
| GET | `/wms/putaway-request-detail/get` | getPutawayRequestDetail | **编号** |
| GET | `/wms/putaway-request-detail/page` | getPutawayRequestDetailPage | **编号** |
| POST | `/wms/putaway-request-detail/senior` | getPutawayRequestDetailSenior | **上架申请子** |
| GET | `/wms/putaway-request-detail/pageChildPackingNumber` | getPutawayRequestDetailPageChildPackingNumber | **上架申请子** |
| POST | `/wms/putaway-request-main/create` | createPutawayRequestMain | **上架申请主** |
| PUT | `/wms/putaway-request-main/update` | updatePutawayRequestMain | **上架申请主** |
| DELETE | `/wms/putaway-request-main/delete` | deletePutawayRequestMain | **上架申请主** |
| GET | `/wms/putaway-request-main/get` | getPutawayRequestMain | **编号** |
| POST | `/wms/putaway-request-main/senior` | getPutawayRequestMainSenior | **编号** |
| GET | `/wms/putaway-request-main/page` | getPutawayRequestMainPage | **上架申请主** |
| GET | `/wms/putaway-request-main/export-excel` | exportPutawayRequestMainExcel | **上架申请主** |
| POST | `/wms/putaway-request-main/export-excel-senior` | exportPutawayRequestMainSeniorExcel | **上架申请主** |
| GET | `/wms/putaway-request-main/get-import-template` | importTemplate | **上架申请主** |
| GET | `/wms/putaway-request-main/getPutawayRequestById` | getPutawayRequestById | **上架申请主** |
| POST | `/wms/putaway-request-main/import` | importExcel | **编号** |
| PUT | `/wms/putaway-request-main/close` | closePutawayRequestMain | **Excel 文件** |
| PUT | `/wms/putaway-request-main/reAdd` | openPutawayRequestMain | **编号** |
| PUT | `/wms/putaway-request-main/submit` | submitPutawayRequestMain | **编号** |
| PUT | `/wms/putaway-request-main/agree` | agreePutawayRequestMain | **编号** |
| PUT | `/wms/putaway-request-main/handle` | handlePutawayRequestMain | **编号** |
| PUT | `/wms/putaway-request-main/refused` | abortPutawayRequestMain | **编号** |

### QAD成本中心 (qadCostcentre)  (11个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/qad-costcentre/create` | createQadCostcentre | **QAD成品中心** |
| PUT | `/wms/qad-costcentre/update` | updateQadCostcentre | **QAD成品中心** |
| DELETE | `/wms/qad-costcentre/delete` | deleteQadCostcentre | **QAD成品中心** |
| GET | `/wms/qad-costcentre/get` | getQadCostcentre | **编号** |
| GET | `/wms/qad-costcentre/page` | getQadCostcentrePage | **编号** |
| POST | `/wms/qad-costcentre/senior` | getQadCostcentreSenior | **QAD成品中心** |
| GET | `/wms/qad-costcentre/export-excel` | exportQadCostcentreExcel | **QAD成品中心** |
| GET | `/wms/qad-costcentre/get-import-template` | importTemplate | **QAD成品中心** |
| POST | `/wms/qad-costcentre/import` | importExcel | **QAD成品中心** |
| GET | `/wms/qad-costcentre/queryCostcentreType` | queryCostcentreType | **Excel 文件** |
| GET | `/wms/qad-costcentre/listByCostcentreCode` | queryCostcentreCode | **QAD成品中心** |

### QAD项目 (qadProject)  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/qad-project/create` | createQadProject | **QAD项目信息** |
| PUT | `/wms/qad-project/update` | updateQadProject | **QAD项目信息** |
| DELETE | `/wms/qad-project/delete` | deleteQadProject | **QAD项目信息** |
| GET | `/wms/qad-project/get` | getQadProject | **编号** |
| GET | `/wms/qad-project/page` | getQadProjectPage | **编号** |
| POST | `/wms/qad-project/senior` | getQadProjectSenior | **QAD项目信息** |
| GET | `/wms/qad-project/export-excel` | exportQadProjectExcel | **QAD项目信息** |
| GET | `/wms/qad-project/get-import-template` | importTemplate | **QAD项目信息** |
| POST | `/wms/qad-project/import` | importExcel | **QAD项目信息** |

### QAD生产计划 (qadproductionplan)  (12个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/qad-production-plan-detail/page` | getQadProductionPlanMainPage | **QAD生产计划子** |
| POST | `/wms/qad-production-plan-detail/senior` | getQadProductionPlanMainSenior | **QAD生产计划子** |
| POST | `/wms/qad-production-plan-main/create` | createQadProductionPlanMain | **QAD生产计划主** |
| PUT | `/wms/qad-production-plan-main/update` | updateQadProductionPlanMain | **QAD生产计划主** |
| DELETE | `/wms/qad-production-plan-main/delete` | deleteQadProductionPlanMain | **QAD生产计划主** |
| GET | `/wms/qad-production-plan-main/get` | getQadProductionPlanMain | **编号** |
| GET | `/wms/qad-production-plan-main/page` | getQadProductionPlanMainPage | **编号** |
| POST | `/wms/qad-production-plan-main/senior` | getQadProductionPlanMainSenior | **QAD生产计划主** |
| POST | `/wms/qad-production-plan-main/export-excel-senior` | exportQadProductionPlanMainExcelSenior | **QAD生产计划主** |
| GET | `/wms/qad-production-plan-main/export-excel` | exportQadProductionPlanMainExcel | **QAD生产计划主** |
| GET | `/wms/qad-production-plan-main/get-import-template` | importTemplate | **QAD生产计划主** |
| POST | `/wms/qad-production-plan-main/import` | importExcel | **QAD生产计划主** |

### 推荐库位历史 (recommendlocationhistory)  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/recommend-location-history/create` | createRecommendLocationHistory | **推荐库位推荐记录** |
| PUT | `/wms/recommend-location-history/update` | updateRecommendLocationHistory | **推荐库位推荐记录** |
| DELETE | `/wms/recommend-location-history/delete` | deleteRecommendLocationHistory | **推荐库位推荐记录** |
| GET | `/wms/recommend-location-history/get` | getRecommendLocationHistory | **编号** |
| GET | `/wms/recommend-location-history/page` | getRecommendLocationHistoryPage | **编号** |
| POST | `/wms/recommend-location-history/senior` | getRecommendLocationHistorySenior | **推荐库位推荐记录** |
| GET | `/wms/recommend-location-history/export-excel` | exportRecommendLocationHistoryExcel | **推荐库位推荐记录** |
| GET | `/wms/recommend-location-history/get-import-template` | importTemplate | **推荐库位推荐记录** |
| POST | `/wms/recommend-location-history/import` | importExcel | **推荐库位推荐记录** |

### reconciliation/notpackagebalancetemp  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/not-package-balance-temp/create` | createNotPackageBalanceTemp | **对账_无包装库存余额临时** |
| PUT | `/wms/not-package-balance-temp/update` | updateNotPackageBalanceTemp | **对账_无包装库存余额临时** |
| DELETE | `/wms/not-package-balance-temp/delete` | deleteNotPackageBalanceTemp | **对账_无包装库存余额临时** |
| GET | `/wms/not-package-balance-temp/get` | getNotPackageBalanceTemp | **编号** |
| GET | `/wms/not-package-balance-temp/page` | getNotPackageBalanceTempPage | **编号** |
| POST | `/wms/not-package-balance-temp/senior` | getNotPackageBalanceTempSenior | **对账_无包装库存余额临时** |
| GET | `/wms/not-package-balance-temp/export-excel` | exportNotPackageBalanceTempExcel | **对账_无包装库存余额临时** |
| GET | `/wms/not-package-balance-temp/get-import-template` | importTemplate | **对账_无包装库存余额临时** |
| POST | `/wms/not-package-balance-temp/import` | importExcel | **对账_无包装库存余额临时** |

### reconciliation/notpackagetransactionbalance  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/not-package-transaction-balance/create` | createNotPackageTransactionBalance | **对账_无包装事务余额对数结果** |
| PUT | `/wms/not-package-transaction-balance/update` | updateNotPackageTransactionBalance | **对账_无包装事务余额对数结果** |
| DELETE | `/wms/not-package-transaction-balance/delete` | deleteNotPackageTransactionBalance | **对账_无包装事务余额对数结果** |
| GET | `/wms/not-package-transaction-balance/get` | getNotPackageTransactionBalance | **编号** |
| GET | `/wms/not-package-transaction-balance/page` | getNotPackageTransactionBalancePage | **编号** |
| POST | `/wms/not-package-transaction-balance/senior` | getNotPackageTransactionBalanceSenior | **对账_无包装事务余额对数结果** |
| GET | `/wms/not-package-transaction-balance/export-excel` | exportNotPackageTransactionBalanceExcel | **对账_无包装事务余额对数结果** |
| GET | `/wms/not-package-transaction-balance/get-import-template` | importTemplate | **对账_无包装事务余额对数结果** |
| POST | `/wms/not-package-transaction-balance/import` | importExcel | **对账_无包装事务余额对数结果** |

### reconciliation/notpackagetransactiontemp  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/not-package-transaction-temp/create` | createNotPackageTransactionTemp | **对账_无包装库存事务临时** |
| PUT | `/wms/not-package-transaction-temp/update` | updateNotPackageTransactionTemp | **对账_无包装库存事务临时** |
| DELETE | `/wms/not-package-transaction-temp/delete` | deleteNotPackageTransactionTemp | **对账_无包装库存事务临时** |
| GET | `/wms/not-package-transaction-temp/get` | getNotPackageTransactionTemp | **编号** |
| GET | `/wms/not-package-transaction-temp/page` | getNotPackageTransactionTempPage | **编号** |
| POST | `/wms/not-package-transaction-temp/senior` | getNotPackageTransactionTempSenior | **对账_无包装库存事务临时** |
| GET | `/wms/not-package-transaction-temp/export-excel` | exportNotPackageTransactionTempExcel | **对账_无包装库存事务临时** |
| GET | `/wms/not-package-transaction-temp/get-import-template` | importTemplate | **对账_无包装库存事务临时** |
| POST | `/wms/not-package-transaction-temp/import` | importExcel | **对账_无包装库存事务临时** |

### reconciliation/transactionbalancepackage  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/transaction-balance-package/create` | createTransactionBalancePackage | **对账_包装号的事务汇总与余额的差异** |
| PUT | `/wms/transaction-balance-package/update` | updateTransactionBalancePackage | **对账_包装号的事务汇总与余额的差异** |
| DELETE | `/wms/transaction-balance-package/delete` | deleteTransactionBalancePackage | **对账_包装号的事务汇总与余额的差异** |
| GET | `/wms/transaction-balance-package/get` | getTransactionBalancePackage | **编号** |
| GET | `/wms/transaction-balance-package/page` | getTransactionBalancePackagePage | **编号** |
| POST | `/wms/transaction-balance-package/senior` | getTransactionBalancePackageSenior | **对账_包装号的事务汇总与余额的差异** |
| GET | `/wms/transaction-balance-package/export-excel` | exportTransactionBalancePackageExcel | **对账_包装号的事务汇总与余额的差异** |
| GET | `/wms/transaction-balance-package/get-import-template` | importTemplate | **对账_包装号的事务汇总与余额的差异** |
| POST | `/wms/transaction-balance-package/import` | importExcel | **对账_包装号的事务汇总与余额的差异** |

### 记录设置 (recordsetting)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/recordsetting/create` | createRecordsetting | **记录设置** |
| PUT | `/wms/recordsetting/update` | updateRecordsetting | **记录设置** |
| POST | `/wms/recordsetting/senior` | getRecordsettingSenior | **记录设置** |
| DELETE | `/wms/recordsetting/delete` | deleteRecordsetting | **记录设置** |
| GET | `/wms/recordsetting/get` | getRecordsetting | **编号** |
| GET | `/wms/recordsetting/list` | getRecordsettingList | **编号** |
| GET | `/wms/recordsetting/page` | getRecordsettingPage | **编号列表** |
| GET | `/wms/recordsetting/export-excel` | exportRecordsettingExcel | **记录设置** |
| GET | `/wms/recordsetting/get-import-template` | importTemplate | **记录设置** |
| POST | `/wms/recordsetting/import` | importExcel | **记录设置** |

### 转寄记录 (relegateRecord)  (15个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/relegate-record-detail/create` | createRelegateRecordDetail | **物料降级记录子** |
| PUT | `/wms/relegate-record-detail/update` | updateRelegateRecordDetail | **物料降级记录子** |
| DELETE | `/wms/relegate-record-detail/delete` | deleteRelegateRecordDetail | **物料降级记录子** |
| GET | `/wms/relegate-record-detail/get` | getRelegateRecordDetail | **编号** |
| GET | `/wms/relegate-record-detail/page` | getRelegateRecordDetailPage | **编号** |
| POST | `/wms/relegate-record-detail/senior` | getRelegateRecordDetailSenior | **物料降级记录子** |
| GET | `/wms/relegate-record-detail/export-excel` | exportRelegateRecordDetailExcel | **物料降级记录子** |
| POST | `/wms/relegate-record-main/create` | createRelegateRecordMain | **物料降级记录主** |
| PUT | `/wms/relegate-record-main/update` | updateRelegateRecordMain | **物料降级记录主** |
| DELETE | `/wms/relegate-record-main/delete` | deleteRelegateRecordMain | **物料降级记录主** |
| GET | `/wms/relegate-record-main/get` | getRelegateRecordMain | **编号** |
| GET | `/wms/relegate-record-main/page` | getRelegateRecordMainPage | **编号** |
| POST | `/wms/relegate-record-main/senior` | getRelegateRecordMainSenior | **物料降级记录主** |
| GET | `/wms/relegate-record-main/export-excel` | exportRelegateRecordMainExcel | **物料降级记录主** |
| POST | `/wms/relegate-record-main/export-excel-senior` | exportRelegateRecordMainSeniorExcel | **物料降级记录主** |

### 转寄申请 (relegateRequest)  (27个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/relegate-request-detail/create` | createRelegateRequestDetail | **物料降级信息** |
| PUT | `/wms/relegate-request-detail/update` | updateRelegateRequestDetail | **物料降级信息** |
| DELETE | `/wms/relegate-request-detail/delete` | deleteRelegateRequestDetail | **物料降级信息** |
| GET | `/wms/relegate-request-detail/get` | getRelegateRequestDetail | **编号** |
| GET | `/wms/relegate-request-detail/page` | getRelegateRequestDetailPage | **编号** |
| POST | `/wms/relegate-request-detail/senior` | getRelegateRequestDetailSenior | **物料降级信息** |
| GET | `/wms/relegate-request-detail/export-excel` | exportRelegateRequestDetailExcel | **物料降级信息** |
| GET | `/wms/relegate-request-detail/get-import-template` | importTemplate | **物料降级信息** |
| POST | `/wms/relegate-request-detail/import` | importExcel | **物料降级信息** |
| POST | `/wms/relegate-request-main/create` | createRelegateRequestMain | **物料变更申请主** |
| POST | `/wms/relegate-request-main/bind` | bindProductreceiptDetailbDO | **物料变更申请主** |
| PUT | `/wms/relegate-request-main/update` | updateRelegateRequestMain | **物料变更申请主** |
| DELETE | `/wms/relegate-request-main/delete` | deleteRelegateRequestMain | **物料变更申请主** |
| GET | `/wms/relegate-request-main/get` | getRelegateRequestMain | **编号** |
| GET | `/wms/relegate-request-main/page` | getRelegateRequestMainPage | **编号** |
| POST | `/wms/relegate-request-main/senior` | getRelegateRequestMainSenior | **物料变更申请主** |
| GET | `/wms/relegate-request-main/export-excel` | exportRelegateRequestMainExcel | **物料变更申请主** |
| POST | `/wms/relegate-request-main/export-excel-senior` | exportRelegateRequestMainSeniorExcel | **物料变更申请主** |
| GET | `/wms/relegate-request-main/get-import-template` | importTemplate | **物料变更申请主** |
| POST | `/wms/relegate-request-main/import` | importExcel | **物料变更申请主** |
| PUT | `/wms/relegate-request-main/close` | closeRelegateRequestMain | **Excel 文件** |
| PUT | `/wms/relegate-request-main/reAdd` | reAddRelegateRequestMain | **编号** |
| PUT | `/wms/relegate-request-main/submit` | submitRelegateRequestMain | **编号** |
| PUT | `/wms/relegate-request-main/agree` | agreeRelegateRequestMain | **编号** |
| PUT | `/wms/relegate-request-main/handle` | handleRelegateRequestMain | **编号** |
| PUT | `/wms/relegate-request-main/refused` | abortRelegateRequestMain | **编号** |
| POST | `/wms/relegate-request-main/relegateCreateLabel` | relegateCreateLabel | **编号** |

### 补货任务 (repleinshJob)  (22个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/job/repleinsh-job-detail/create` | createRepleinshJobDetail | **补料任务子** |
| PUT | `/job/repleinsh-job-detail/update` | updateRepleinshJobDetail | **补料任务子** |
| DELETE | `/job/repleinsh-job-detail/delete` | deleteRepleinshJobDetail | **补料任务子** |
| GET | `/job/repleinsh-job-detail/get` | getRepleinshJobDetail | **编号** |
| POST | `/job/repleinsh-job-detail/senior` | getRepleinshJobDetailSenior | **编号** |
| GET | `/job/repleinsh-job-detail/list` | getRepleinshJobDetailList | **补料任务子** |
| GET | `/job/repleinsh-job-detail/page` | getRepleinshJobDetailPage | **编号列表** |
| GET | `/job/repleinsh-job-detail/export-excel` | exportRepleinshJobDetailExcel | **补料任务子** |
| POST | `/wms/repleinsh-job-main/senior` | getRepleinshJobMainSenior | **补料任务主** |
| GET | `/wms/repleinsh-job-main/get` | getRepleinshJobMain | **补料任务主** |
| GET | `/wms/repleinsh-job-main/list` | getRepleinshJobMainList | **编号** |
| GET | `/wms/repleinsh-job-main/page` | getRepleinshJobMainPage | **编号列表** |
| GET | `/wms/repleinsh-job-main/export-excel` | exportRepleinshJobMainExcel | **补料任务主** |
| POST | `/wms/repleinsh-job-main/export-excel-senior` | exportRepleinshJobMainSeniorExcel | **补料任务主** |
| PUT | `/wms/repleinsh-job-main/accept` | acceptRepleinshJobMain | **补料任务主** |
| PUT | `/wms/repleinsh-job-main/abandon` | abandonRepleinshJobMain | **补料任务主** |
| PUT | `/wms/repleinsh-job-main/close` | closeRepleinshJobMain | **补料任务主** |
| PUT | `/wms/repleinsh-job-main/execute` | executeRepleinshJobMain | **补料任务主** |
| GET | `/wms/repleinsh-job-main/getRepleinshJobById` | getRepleinshJobById | **补料任务主** |
| POST | `/wms/repleinsh-job-main/getRepleinshJobbPageByStatusAndTime` | getRepleinshJobbPageByStatusAndTime | **编号** |
| POST | `/wms/repleinsh-job-main/getCountByStatus` | getCountByStatus | **今日开始结束时间** |
| PUT | `/wms/repleinsh-job-main/updateRepleinshJobConfig` | updateRepleinshJobConfig | **类型数组** |

### 补货记录 (repleinshRecord)  (15个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/repleinsh-record-detail/create` | createRepleinshRecordDetail | **补料记录子** |
| PUT | `/wms/repleinsh-record-detail/update` | updateRepleinshRecordDetail | **补料记录子** |
| DELETE | `/wms/repleinsh-record-detail/delete` | deleteRepleinshRecordDetail | **补料记录子** |
| POST | `/wms/repleinsh-record-detail/senior` | getRepleinshRecordDetailSenior | **编号** |
| GET | `/wms/repleinsh-record-detail/get` | getRepleinshRecordDetail | **补料记录子** |
| GET | `/wms/repleinsh-record-detail/list` | getRepleinshRecordDetailList | **编号** |
| GET | `/wms/repleinsh-record-detail/page` | getRepleinshRecordDetailPage | **编号列表** |
| GET | `/wms/repleinsh-record-detail/export-excel` | exportRepleinshRecordDetailExcel | **补料记录子** |
| POST | `/wms/repleinsh-record-main/create` | createRepleinshRecordMain | **补料记录主** |
| POST | `/wms/repleinsh-record-main/senior` | getRepleinshRecordMainSenior | **补料记录主** |
| GET | `/wms/repleinsh-record-main/get` | getRepleinshRecordMain | **补料记录主** |
| GET | `/wms/repleinsh-record-main/list` | getRepleinshRecordMainList | **编号** |
| GET | `/wms/repleinsh-record-main/page` | getRepleinshRecordMainPage | **编号列表** |
| GET | `/wms/repleinsh-record-main/export-excel` | exportRepleinshRecordMainExcel | **补料记录主** |
| POST | `/wms/repleinsh-record-main/export-excel-senior` | exportRepleinshRecordMainSeniorExcel | **补料记录主** |

### 补货申请 (repleinshRequest)  (24个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/repleinsh-request-detail/create` | createRepleinshRequestDetail | **补料申请子** |
| PUT | `/wms/repleinsh-request-detail/update` | updateRepleinshRequestDetail | **补料申请子** |
| DELETE | `/wms/repleinsh-request-detail/delete` | deleteRepleinshRequestDetail | **补料申请子** |
| POST | `/wms/repleinsh-request-detail/senior` | getRepleinshRequestDetailSenior | **编号** |
| GET | `/wms/repleinsh-request-detail/get` | getRepleinshRequestDetail | **补料申请子** |
| GET | `/wms/repleinsh-request-detail/list` | getRepleinshRequestDetailList | **编号** |
| GET | `/wms/repleinsh-request-detail/page` | getRepleinshRequestDetailPage | **编号列表** |
| POST | `/wms/repleinsh-request-main/create` | createRepleinshRequestMain | **补料申请主** |
| PUT | `/wms/repleinsh-request-main/update` | updateRepleinshRequestMain | **补料申请主** |
| DELETE | `/wms/repleinsh-request-main/delete` | deleteRepleinshRequestMain | **补料申请主** |
| POST | `/wms/repleinsh-request-main/senior` | getRepleinshRequestMainSenior | **编号** |
| GET | `/wms/repleinsh-request-main/get` | getRepleinshRequestMain | **补料申请主** |
| GET | `/wms/repleinsh-request-main/list` | getRepleinshRequestMainList | **编号** |
| GET | `/wms/repleinsh-request-main/page` | getRepleinshRequestMainPage | **编号列表** |
| POST | `/wms/repleinsh-request-main/export-excel-senior` | exportRepleinshRequestMainSeniorExcel | **补料申请主** |
| GET | `/wms/repleinsh-request-main/export-excel` | exportRepleinshRequestMainExcel | **补料申请主** |
| GET | `/wms/repleinsh-request-main/get-import-template` | importTemplate | **补料申请主** |
| POST | `/wms/repleinsh-request-main/import` | importExcel | **补料申请主** |
| PUT | `/wms/repleinsh-request-main/close` | closeRepleinshRequestMain | **Excel 文件** |
| PUT | `/wms/repleinsh-request-main/reAdd` | reAddRepleinshRequestMain | **编号** |
| PUT | `/wms/repleinsh-request-main/submit` | submitRepleinshRequestMain | **编号** |
| PUT | `/wms/repleinsh-request-main/refused` | refusedRepleinshRequestMain | **编号** |
| PUT | `/wms/repleinsh-request-main/agree` | agreeRepleinshRequestMain | **编号** |
| PUT | `/wms/repleinsh-request-main/handle` | handleRepleinshRequestMain | **编号** |

### 报表 (report)  (23个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/jimu-report/queryRecordSupplierdeliverMain` | queryRecordSupplierdeliverMain | **积木报表接口** |
| GET | `/wms/jimu-report/queryRecordSupplierdeliverDetail` | queryRecordSupplierdeliverDetail | **积木报表接口** |
| GET | `/wms/jimu-report/queryRequestSupplierinvoiceMain` | queryReportRequestSupplierinvoiceMain | **积木报表接口** |
| GET | `/wms/jimu-report/queryRecordSupplierinvoiceMain` | queryRecordSupplierinvoiceMain | **积木报表接口** |
| GET | `/wms/jimu-report/queryRecordSupplierinvoiceDetail` | queryRecordSupplierinvoiceDetail | **积木报表接口** |
| GET | `/wms/jimu-report/listSupplierInvoiceMain` | listSupplierInvoiceMain | **积木报表接口** |
| GET | `/wms/jimu-report/queryRequestSupplierinvoiceDetail` | queryReportRequestSupplierinvoiceDetail | **积木报表接口** |
| GET | `/wms/jimu-report/getCallMaterialsLabelList` | getCallMaterialsLabelList | **积木报表接口** |
| GET | `/wms/jimu-report/getYCLLabelList` | getYCLLabelList | **积木报表接口** |
| GET | `/wms/jimu-report/getLocationLabelList` | getLocationLabelList | **积木报表接口** |
| GET | `/wms/jimu-report/getPutawayJobMain` | getPutawayJobMain | **积木报表接口** |
| GET | `/wms/jimu-report/getPutawayJobDetail` | getPutawayJobDetail | **积木报表接口** |
| GET | `/wms/jimu-report/getPurchasereturnRecordMain` | getPurchasereturnRecordMain | **积木报表接口** |
| GET | `/wms/jimu-report/getPurchasereturnRecordMainSCP` | getPurchasereturnRecordMainSCP | **积木报表接口** |
| GET | `/wms/jimu-report/getPurchasereturnRecordDetail` | getPurchasereturnRecordDetail | **积木报表接口** |
| GET | `/wms/jimu-report/getPurchasereturnRecordDetailSCP` | getPurchasereturnRecordDetailSCP | **积木报表接口** |
| GET | `/wms/jimu-report/getProductPutawayJobMain` | getProductPutawayJobMain | **积木报表接口** |
| GET | `/wms/jimu-report/getProductPutawayJobDetail` | getProductPutawayJobDetail | **积木报表接口** |
| POST | `/wms/jimu-report/getPutawayJobDetailForPDA` | getPutawayJobDetailForPDA | **积木报表接口** |
| GET | `/wms/jimu-report/listSupplierInvoiceRequestMain` | listSupplierInvoiceRequestMain | **积木报表接口** |
| GET | `/wms/jimu-report/getInspectionInfo` | getInspectionInfo | **积木报表接口** |
| GET | `/wms/jimu-report/getCustContainerList` | getCustContainerList | **积木报表接口** |
| GET | `/wms/jimu-report/getFactoryContainerList` | getFactoryContainerList | **积木报表接口** |

### 申请设置 (requestsetting)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/requestsetting/create` | createRequestsetting | **申请设置** |
| PUT | `/wms/requestsetting/update` | updateRequestsetting | **申请设置** |
| DELETE | `/wms/requestsetting/delete` | deleteRequestsetting | **申请设置** |
| GET | `/wms/requestsetting/get` | getRequestsetting | **编号** |
| POST | `/wms/requestsetting/senior` | getRequestsettingSenior | **编号** |
| GET | `/wms/requestsetting/list` | getRequestsettingList | **申请设置** |
| GET | `/wms/requestsetting/page` | getRequestsettingPage | **编号列表** |
| GET | `/wms/requestsetting/export-excel` | exportRequestsettingExcel | **申请设置** |
| GET | `/wms/requestsetting/get-import-template` | importTemplate | **申请设置** |
| POST | `/wms/requestsetting/import` | importExcel | **申请设置** |

### 规则 (rule)  (13个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/rule/create` | createRule | **规则** |
| PUT | `/wms/rule/update` | updateRule | **规则** |
| DELETE | `/wms/rule/delete` | deleteRule | **规则** |
| GET | `/wms/rule/get` | getRule | **编号** |
| POST | `/wms/rule/senior` | getRuleSenior | **编号** |
| GET | `/wms/rule/page` | getRulePage | **规则** |
| POST | `/wms/rule/getPrecisionStrategyByItemCodes` | getPrecisionStrategyByItemCodes | **规则** |
| POST | `/wms/rule/getPrecisionStrategy` | getPrecisionStrategy | **规则** |
| GET | `/wms/rule/export-excel` | exportRuleExcel | **规则** |
| POST | `/wms/rule/export-excel-senior` | exportItembasicExcel | **规则** |
| POST | `/wms/rule/import` | importExcel | **规则** |
| GET | `/wms/rule/get-import-template` | importTemplate | **Excel 文件** |
| GET | `/wms/rule/getMaxPriority` | getMaxPriority | **规则** |

### 销售 (sale)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/sale-detail/create` | createSaleDetail | **销售订单子** |
| PUT | `/wms/sale-detail/update` | updateSaleDetail | **销售订单子** |
| DELETE | `/wms/sale-detail/delete` | deleteSaleDetail | **销售订单子** |
| GET | `/wms/sale-detail/get` | getSaleDetail | **编号** |
| GET | `/wms/sale-detail/list` | getSaleDetailList | **编号** |
| GET | `/wms/sale-detail/page` | getSaleDetailPage | **编号列表** |
| POST | `/wms/sale-detail/senior` | getSaleDetailSenior | **销售订单子** |
| GET | `/wms/sale-detail/export-excel` | exportSaleDetailExcel | **销售订单子** |
| POST | `/wms/sale-main/create` | createSaleMain | **销售订单主** |
| PUT | `/wms/sale-main/update` | updateSaleMain | **销售订单主** |
| POST | `/wms/sale-main/senior` | getSaleMainSenior | **销售订单主** |
| DELETE | `/wms/sale-main/delete` | deleteSaleMain | **销售订单主** |
| GET | `/wms/sale-main/get` | getSaleMain | **编号** |
| GET | `/wms/sale-main/list` | getSaleMainList | **编号** |
| GET | `/wms/sale-main/page` | getSaleMainPage | **编号列表** |
| GET | `/wms/sale-main/export-excel` | exportSaleMainExcel | **销售订单主** |
| POST | `/wms/sale-main/export-excel-senior` | exportRepleinshRequestMainSeniorExcel | **销售订单主** |
| GET | `/wms/sale-main/get-import-template` | importTemplate | **销售订单主** |
| POST | `/wms/sale-main/import` | importExcel | **销售订单主** |

### 销售发货记录 (saleShipmentRecord)  (17个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/sale-shipment-detail-record/create` | createSaleShipmentDetailRecord | **结算出库记录子** |
| PUT | `/wms/sale-shipment-detail-record/update` | updateSaleShipmentDetailRecord | **结算出库记录子** |
| DELETE | `/wms/sale-shipment-detail-record/delete` | deleteSaleShipmentDetailRecord | **结算出库记录子** |
| GET | `/wms/sale-shipment-detail-record/get` | getSaleShipmentDetailRecord | **编号** |
| GET | `/wms/sale-shipment-detail-record/page` | getSaleShipmentDetailRecordPage | **编号** |
| POST | `/wms/sale-shipment-detail-record/senior` | getSaleShipmentDetailRecordSenior | **结算出库记录子** |
| GET | `/wms/sale-shipment-detail-record/export-excel` | exportSaleShipmentDetailRecordExcel | **结算出库记录子** |
| GET | `/wms/sale-shipment-detail-record/get-import-template` | importTemplate | **结算出库记录子** |
| POST | `/wms/sale-shipment-main-record/create` | createSaleShipmentMainRecord | **结算出库记录主** |
| PUT | `/wms/sale-shipment-main-record/update` | updateSaleShipmentMainRecord | **结算出库记录主** |
| DELETE | `/wms/sale-shipment-main-record/delete` | deleteSaleShipmentMainRecord | **结算出库记录主** |
| GET | `/wms/sale-shipment-main-record/get` | getSaleShipmentMainRecord | **编号** |
| GET | `/wms/sale-shipment-main-record/page` | getSaleShipmentMainRecordPage | **编号** |
| POST | `/wms/sale-shipment-main-record/senior` | getSaleShipmentMainRecordSenior | **结算出库记录主** |
| GET | `/wms/sale-shipment-main-record/export-excel` | exportSaleShipmentMainRecordExcel | **结算出库记录主** |
| POST | `/wms/sale-shipment-main-record/export-excel-senior` | exportSaleShipmentMainRecordSeniorExcel | **结算出库记录主** |
| PUT | `/wms/sale-shipment-main-record/abort` | abort | **结算出库记录主** |

### 销售发货申请 (saleShipmentRequest)  (23个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/sale-shipment-detail-request/create` | createSaleShipmentDetail | **结算出库申请子** |
| PUT | `/wms/sale-shipment-detail-request/update` | updateSaleShipmentDetail | **结算出库申请子** |
| DELETE | `/wms/sale-shipment-detail-request/delete` | deleteSaleShipmentDetail | **结算出库申请子** |
| GET | `/wms/sale-shipment-detail-request/get` | getSaleShipmentDetail | **编号** |
| GET | `/wms/sale-shipment-detail-request/page` | getSaleShipmentDetailPage | **编号** |
| POST | `/wms/sale-shipment-detail-request/senior` | getSaleShipmentDetailSenior | **结算出库申请子** |
| GET | `/wms/sale-shipment-detail-request/export-excel` | exportSaleShipmentDetailExcel | **结算出库申请子** |
| POST | `/wms/sale-shipment-main-request/create` | createSaleShipmentMain | **结算出库申请主** |
| PUT | `/wms/sale-shipment-main-request/update` | updateSaleShipmentMain | **结算出库申请主** |
| DELETE | `/wms/sale-shipment-main-request/delete` | deleteSaleShipmentMain | **结算出库申请主** |
| GET | `/wms/sale-shipment-main-request/get` | getSaleShipmentMain | **编号** |
| GET | `/wms/sale-shipment-main-request/page` | getSaleShipmentMainPage | **编号** |
| POST | `/wms/sale-shipment-main-request/senior` | getSaleShipmentMainSenior | **结算出库申请主** |
| GET | `/wms/sale-shipment-main-request/export-excel` | exportSaleShipmentMainExcel | **结算出库申请主** |
| POST | `/wms/sale-shipment-main-request/export-excel-senior` | exportSaleShipmentMainSeniorExcel | **结算出库申请主** |
| GET | `/wms/sale-shipment-main-request/get-import-template` | importTemplate | **结算出库申请主** |
| POST | `/wms/sale-shipment-main-request/import` | importExcel | **结算出库申请主** |
| PUT | `/wms/sale-shipment-main-request/close` | closeSaleShipmentMainRequest | **Excel 文件** |
| PUT | `/wms/sale-shipment-main-request/reAdd` | reAddSaleShipmentMainRequest | **编号** |
| PUT | `/wms/sale-shipment-main-request/submit` | submitSaleShipmentMainRequest | **编号** |
| PUT | `/wms/sale-shipment-main-request/agree` | agreeSaleShipmentMainRequest | **编号** |
| PUT | `/wms/sale-shipment-main-request/handle` | handleSaleShipmentMainRequest | **编号** |
| PUT | `/wms/sale-shipment-main-request/refused` | abortSaleShipmentMainRequest | **编号** |

### 销售价格 (saleprice)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/saleprice/create` | createSaleprice | **销售价格单** |
| PUT | `/wms/saleprice/update` | updateSaleprice | **销售价格单** |
| DELETE | `/wms/saleprice/delete` | deleteSaleprice | **销售价格单** |
| GET | `/wms/saleprice/get` | getSaleprice | **编号** |
| GET | `/wms/saleprice/page` | getSalepricePage | **编号** |
| POST | `/wms/saleprice/senior` | getSalepriceSenior | **销售价格单** |
| GET | `/wms/saleprice/export-excel` | exportSalepriceExcel | **销售价格单** |
| POST | `/wms/saleprice/export-excel-senior` | exportSalepriceExcel | **销售价格单** |
| GET | `/wms/saleprice/get-import-template` | importTemplate | **销售价格单** |
| POST | `/wms/saleprice/import` | importExcel | **销售价格单** |

### 报废任务 (scrapJob)  (23个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/scrap-job-detail/create` | createScrapJobDetail | **报废出库任务子** |
| PUT | `/wms/scrap-job-detail/update` | updateScrapJobDetail | **报废出库任务子** |
| DELETE | `/wms/scrap-job-detail/delete` | deleteScrapJobDetail | **报废出库任务子** |
| GET | `/wms/scrap-job-detail/get` | getScrapJobDetail | **编号** |
| GET | `/wms/scrap-job-detail/list` | getScrapJobDetailList | **编号** |
| GET | `/wms/scrap-job-detail/page` | getScrapJobDetailPage | **编号列表** |
| POST | `/wms/scrap-job-detail/senior` | getScrapJobDetailSenior | **报废出库任务子** |
| GET | `/wms/scrap-job-detail/export-excel` | exportScrapJobDetailExcel | **报废出库任务子** |
| POST | `/wms/scrap-job-main/create` | createScrapJobMain | **报废出库任务主** |
| POST | `/wms/scrap-job-main/senior` | getScrapJobMainSenior | **报废出库任务主** |
| PUT | `/wms/scrap-job-main/update` | updateScrapJobMain | **报废出库任务主** |
| DELETE | `/wms/scrap-job-main/delete` | deleteScrapJobMain | **报废出库任务主** |
| GET | `/wms/scrap-job-main/get` | getScrapJobMain | **编号** |
| GET | `/wms/scrap-job-main/list` | getScrapJobMainList | **编号** |
| GET | `/wms/scrap-job-main/page` | getScrapJobMainPage | **编号列表** |
| GET | `/wms/scrap-job-main/export-excel` | exportScrapJobMainExcel | **报废出库任务主** |
| POST | `/wms/scrap-job-main/export-excel-senior` | exportScrapJobMainExcel | **报废出库任务主** |
| GET | `/wms/scrap-job-main/getScrapJobById` | getScrapJobById | **报废出库任务主** |
| POST | `/wms/scrap-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/scrap-job-main/accept` | acceptScrapJobMain | **类型数组** |
| PUT | `/wms/scrap-job-main/abandon` | abandonScrapJobMain | **报废出库任务主** |
| PUT | `/wms/scrap-job-main/close` | closeScrapJobMain | **报废出库任务主** |
| PUT | `/wms/scrap-job-main/execute` | executeScrapJobMain | **报废出库任务主** |

### 报废记录 (scrapRecord)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/scrap-record-detail/create` | createScrapRecordDetail | **报废出库记录子** |
| PUT | `/wms/scrap-record-detail/update` | updateScrapRecordDetail | **报废出库记录子** |
| DELETE | `/wms/scrap-record-detail/delete` | deleteScrapRecordDetail | **报废出库记录子** |
| GET | `/wms/scrap-record-detail/get` | getScrapRecordDetail | **编号** |
| GET | `/wms/scrap-record-detail/list` | getScrapRecordDetailList | **编号** |
| GET | `/wms/scrap-record-detail/page` | getScrapRecordDetailPage | **编号列表** |
| POST | `/wms/scrap-record-detail/senior` | getScrapRecordDetailSenior | **报废出库记录子** |
| GET | `/wms/scrap-record-detail/export-excel` | exportScrapRecordDetailExcel | **报废出库记录子** |
| POST | `/wms/scrap-record-main/create` | createScrapRecordMain | **报废出库记录主** |
| PUT | `/wms/scrap-record-main/update` | updateScrapRecordMain | **报废出库记录主** |
| DELETE | `/wms/scrap-record-main/delete` | deleteScrapRecordMain | **报废出库记录主** |
| GET | `/wms/scrap-record-main/get` | getScrapRecordMain | **编号** |
| GET | `/wms/scrap-record-main/list` | getScrapRecordMainList | **编号** |
| GET | `/wms/scrap-record-main/page` | getScrapRecordMainPage | **编号列表** |
| POST | `/wms/scrap-record-main/senior` | getScrapRecordMainSenior | **报废出库记录主** |
| GET | `/wms/scrap-record-main/export-excel` | exportScrapRecordMainExcel | **报废出库记录主** |
| POST | `/wms/scrap-record-main/export-excel-senior` | exportScrapRecordMainExcel | **报废出库记录主** |
| GET | `/wms/scrap-record-main/getDetailInfoById` | getDetailInfoById | **报废出库记录主** |
| GET | `/wms/scrap-record-main/revoke` | revoke | **报废出库记录主** |

### 报废申请 (scrapRequest)  (26个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/scrap-request-detail/create` | createScrapRequestDetail | **报废出库申请子** |
| PUT | `/wms/scrap-request-detail/update` | updateScrapRequestDetail | **报废出库申请子** |
| DELETE | `/wms/scrap-request-detail/delete` | deleteScrapRequestDetail | **报废出库申请子** |
| GET | `/wms/scrap-request-detail/get` | getScrapRequestDetail | **编号** |
| GET | `/wms/scrap-request-detail/list` | getScrapRequestDetailList | **编号** |
| GET | `/wms/scrap-request-detail/page` | getScrapRequestDetailPage | **编号列表** |
| POST | `/wms/scrap-request-detail/senior` | getScrapRequestDetailSenior | **报废出库申请子** |
| GET | `/wms/scrap-request-detail/export-excel` | exportScrapRequestDetailExcel | **报废出库申请子** |
| POST | `/wms/scrap-request-main/create` | createScrapRequestMain | **报废出库申请主** |
| PUT | `/wms/scrap-request-main/update` | updateScrapRequestMain | **报废出库申请主** |
| DELETE | `/wms/scrap-request-main/delete` | deleteScrapRequestMain | **报废出库申请主** |
| GET | `/wms/scrap-request-main/get` | getScrapRequestMain | **编号** |
| GET | `/wms/scrap-request-main/list` | getScrapRequestMainList | **编号** |
| GET | `/wms/scrap-request-main/page` | getScrapRequestMainPage | **编号列表** |
| POST | `/wms/scrap-request-main/senior` | getScrapRequestMainSenior | **报废出库申请主** |
| GET | `/wms/scrap-request-main/get-import-template` | importTemplate | **报废出库申请主** |
| POST | `/wms/scrap-request-main/import` | importExcel | **报废出库申请主** |
| GET | `/wms/scrap-request-main/export-excel` | exportScrapRequestMainExcel | **Excel 文件** |
| POST | `/wms/scrap-request-main/export-excel-senior` | exportScrapRequestMainSeniorExcel | **报废出库申请主** |
| GET | `/wms/scrap-request-main/getScrapRequestById` | getScrapRequestById | **报废出库申请主** |
| PUT | `/wms/scrap-request-main/close` | closeScrapRequestMain | **编号** |
| PUT | `/wms/scrap-request-main/reAdd` | reAddScrapRequestMain | **编号** |
| PUT | `/wms/scrap-request-main/submit` | submitScrapRequestMain | **编号** |
| PUT | `/wms/scrap-request-main/refused` | abortScrapRequestMain | **编号** |
| PUT | `/wms/scrap-request-main/agree` | agreeScrapRequestMain | **编号** |
| PUT | `/wms/scrap-request-main/handle` | handleScrapRequestMain | **编号** |

### 班次 (shift)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/shift/create` | createShift | **班次** |
| PUT | `/wms/shift/update` | updateShift | **班次** |
| DELETE | `/wms/shift/delete` | deleteShift | **班次** |
| GET | `/wms/shift/get` | getShift | **编号** |
| GET | `/wms/shift/page` | getShiftPage | **编号** |
| POST | `/wms/shift/senior` | getShiftSenior | **班次** |
| GET | `/wms/shift/export-excel` | exportShiftExcel | **班次** |
| POST | `/wms/shift/export-excel-senior` | exportShiftExcel | **班次** |
| GET | `/wms/shift/get-import-template` | importTemplate | **班次** |
| POST | `/wms/shift/import` | importExcel | **班次** |

### 备件货位 (spareitemlocation)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/spareitem-location/create` | createSpareitemLocation | **维修备件默认库位** |
| PUT | `/wms/spareitem-location/update` | updateSpareitemLocation | **维修备件默认库位** |
| DELETE | `/wms/spareitem-location/delete` | deleteSpareitemLocation | **维修备件默认库位** |
| GET | `/wms/spareitem-location/get` | getSpareitemLocation | **编号** |
| GET | `/wms/spareitem-location/page` | getSpareitemLocationPage | **编号** |
| POST | `/wms/spareitem-location/senior` | getSpareitemLocationSenior | **维修备件默认库位** |
| GET | `/wms/spareitem-location/export-excel` | exportSpareitemLocationExcel | **维修备件默认库位** |
| GET | `/wms/spareitem-location/get-import-template` | importTemplate | **维修备件默认库位** |
| POST | `/wms/spareitem-location/import` | importExcel | **维修备件默认库位** |
| POST | `/wms/spareitem-location/queryItemLocation` | queryItemLocation | **Excel 文件** |

### 标准成本 (stdcostprice)  (11个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/stdcostprice/create` | createStdcostprice | **标准成本价格单** |
| PUT | `/wms/stdcostprice/update` | updateStdcostprice | **标准成本价格单** |
| DELETE | `/wms/stdcostprice/delete` | deleteStdcostprice | **标准成本价格单** |
| GET | `/wms/stdcostprice/get` | getStdcostprice | **编号** |
| GET | `/wms/stdcostprice/page` | getStdcostpricePage | **编号** |
| POST | `/wms/stdcostprice/senior` | getStdcostpriceSenior | **标准成本价格单** |
| GET | `/wms/stdcostprice/export-excel` | exportStdcostpriceExcel | **标准成本价格单** |
| POST | `/wms/stdcostprice/export-excel-senior` | exportStdcostpriceExcel | **标准成本价格单** |
| GET | `/wms/stdcostprice/get-import-template` | importTemplate | **标准成本价格单** |
| POST | `/wms/stdcostprice/import` | importExcel | **标准成本价格单** |
| POST | `/wms/stdcostprice/queryStdcostpriceByItemCode` | queryStdcostpriceByItemCode | **Excel 文件** |

### 备货任务 (stockupJob)  (22个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/stockup-detail-job/create` | createStockupDetailJob | **备货任务子** |
| PUT | `/wms/stockup-detail-job/update` | updateStockupDetailJob | **备货任务子** |
| DELETE | `/wms/stockup-detail-job/delete` | deleteStockupDetailJob | **备货任务子** |
| GET | `/wms/stockup-detail-job/get` | getStockupDetailJob | **编号** |
| GET | `/wms/stockup-detail-job/list` | getStockupDetailJobList | **编号** |
| GET | `/wms/stockup-detail-job/page` | getStockupDetailJobPage | **编号列表** |
| POST | `/wms/stockup-detail-job/senior` | getStockupDetailJobPageSenior | **备货任务子** |
| GET | `/wms/stockup-detail-job/export-excel` | exportStockupDetailJobExcel | **备货任务子** |
| POST | `/wms/stockup-main-job/create` | createStockupMainJob | **备货任务主** |
| PUT | `/wms/stockup-main-job/update` | updateStockupMainJob | **备货任务主** |
| DELETE | `/wms/stockup-main-job/delete` | deleteStockupMainJob | **备货任务主** |
| GET | `/wms/stockup-main-job/get` | getStockupMainJob | **编号** |
| GET | `/wms/stockup-main-job/list` | getStockupMainJobList | **编号** |
| GET | `/wms/stockup-main-job/page` | getStockupMainJobPage | **编号列表** |
| POST | `/wms/stockup-main-job/senior` | getStockupMainJobPageSenior | **备货任务主** |
| GET | `/wms/stockup-main-job/getStockupMainJobById` | getDeliverJobById | **备货任务主** |
| GET | `/wms/stockup-main-job/export-excel` | exportStockupMainJobExcel | **编号** |
| POST | `/wms/stockup-main-job/export-excel-senior` | exportStockupMainJobMainSeniorExcel | **备货任务主** |
| PUT | `/wms/stockup-main-job/accept` | acceptStockupMainJob | **备货任务主** |
| PUT | `/wms/stockup-main-job/abandon` | abandonStockupMainJob | **编号** |
| PUT | `/wms/stockup-main-job/close` | closeStockupMainJob | **编号** |
| PUT | `/wms/stockup-main-job/execute` | executeStockupMainJob | **编号** |

### 备货记录 (stockupRecord)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/stockup-detail-record/create` | createStockupDetailRecord | **备货记录子** |
| PUT | `/wms/stockup-detail-record/update` | updateStockupDetailRecord | **备货记录子** |
| DELETE | `/wms/stockup-detail-record/delete` | deleteStockupDetailRecord | **备货记录子** |
| GET | `/wms/stockup-detail-record/get` | getStockupDetailRecord | **编号** |
| GET | `/wms/stockup-detail-record/list` | getStockupDetailRecordList | **编号** |
| GET | `/wms/stockup-detail-record/page` | getStockupDetailRecordPage | **编号列表** |
| POST | `/wms/stockup-detail-record/senior` | getStockupDetailRecordPageSenior | **备货记录子** |
| GET | `/wms/stockup-detail-record/export-excel` | exportStockupDetailRecordExcel | **备货记录子** |
| GET | `/wms/stockup-detail-record/get-import-template` | importTemplate | **备货记录子** |
| POST | `/wms/stockup-main-record/create` | createStockupMainRecord | **备货记录主** |
| PUT | `/wms/stockup-main-record/update` | updateStockupMainRecord | **备货记录主** |
| DELETE | `/wms/stockup-main-record/delete` | deleteStockupMainRecord | **备货记录主** |
| GET | `/wms/stockup-main-record/get` | getStockupMainRecord | **编号** |
| GET | `/wms/stockup-main-record/list` | getStockupMainRecordList | **编号** |
| GET | `/wms/stockup-main-record/page` | getStockupMainRecordPage | **编号列表** |
| POST | `/wms/stockup-main-record/senior` | getStockupMainRecordPageSenior | **备货记录主** |
| GET | `/wms/stockup-main-record/export-excel` | exportStockupMainRecordExcel | **备货记录主** |
| POST | `/wms/stockup-main-record/export-excel-senior` | exportStockupMainRequestMainSeniorExcel | **备货记录主** |
| GET | `/wms/stockup-main-record/get-import-template` | importTemplate | **备货记录主** |

### 备货申请 (stockupRequest)  (27个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/stockup-detail-request/create` | createStockupDetailRequest | **备货申请子** |
| PUT | `/wms/stockup-detail-request/update` | updateStockupDetailRequest | **备货申请子** |
| DELETE | `/wms/stockup-detail-request/delete` | deleteStockupDetailRequest | **备货申请子** |
| GET | `/wms/stockup-detail-request/get` | getStockupDetailRequest | **编号** |
| GET | `/wms/stockup-detail-request/list` | getStockupDetailRequestList | **编号** |
| GET | `/wms/stockup-detail-request/page` | getStockupDetailRequestPage | **编号列表** |
| POST | `/wms/stockup-detail-request/senior` | getStockupDetailRequestSenior | **备货申请子** |
| GET | `/wms/stockup-detail-request/export-excel` | exportStockupDetailRequestExcel | **备货申请子** |
| GET | `/wms/stockup-detail-request/get-import-template` | importTemplate | **备货申请子** |
| POST | `/wms/stockup-main-request/create` | createStockupMainRequest | **备货申请主** |
| PUT | `/wms/stockup-main-request/update` | updateStockupMainRequest | **备货申请主** |
| DELETE | `/wms/stockup-main-request/delete` | deleteStockupMainRequest | **备货申请主** |
| GET | `/wms/stockup-main-request/get` | getStockupMainRequest | **编号** |
| GET | `/wms/stockup-main-request/list` | getStockupMainRequestList | **编号** |
| GET | `/wms/stockup-main-request/page` | getStockupMainRequestPage | **编号列表** |
| POST | `/wms/stockup-main-request/senior` | getStockupMainRequestSenior | **备货申请主** |
| GET | `/wms/stockup-main-request/export-excel` | exportStockupMainRequestExcel | **备货申请主** |
| POST | `/wms/stockup-main-request/export-excel-senior` | exportStockupMainRequestMainSeniorExcel | **备货申请主** |
| GET | `/wms/stockup-main-request/get-import-template` | importTemplate | **备货申请主** |
| POST | `/wms/stockup-main-request/import` | importExcel | **备货申请主** |
| GET | `/wms/stockup-main-request/getStockupRequestById` | getDeliverRequestById | **Excel 文件** |
| PUT | `/wms/stockup-main-request/close` | closeStockupMainRequest | **编号** |
| PUT | `/wms/stockup-main-request/reAdd` | reAddStockupMainRequest | **编号** |
| PUT | `/wms/stockup-main-request/submit` | submitStockupMainRequest | **编号** |
| PUT | `/wms/stockup-main-request/agree` | agreeStockupMainRequest | **编号** |
| PUT | `/wms/stockup-main-request/handle` | handleStockupMainRequest | **编号** |
| PUT | `/wms/stockup-main-request/refused` | abortStockupMainRequest | **编号** |

### 策略 (strategy)  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/strategy/create` | createStrategy | **策略** |
| PUT | `/wms/strategy/update` | updateStrategy | **策略** |
| POST | `/wms/strategy/senior` | getStrategySenior | **策略** |
| DELETE | `/wms/strategy/delete` | deleteStrategy | **策略** |
| GET | `/wms/strategy/get` | getStrategy | **编号** |
| GET | `/wms/strategy/list` | getStrategyList | **编号** |
| GET | `/wms/strategy/page` | getStrategyPage | **编号列表** |
| GET | `/wms/strategy/export-excel` | exportStrategyExcel | **策略** |
| POST | `/wms/strategy/export-excel-senior` | exportStrategyExcel | **策略** |

### 科目账户 (subjectaccount)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/subject-account/create` | createSubjectAccount | **科目账户配置** |
| PUT | `/wms/subject-account/update` | updateSubjectAccount | **科目账户配置** |
| DELETE | `/wms/subject-account/delete` | deleteSubjectAccount | **科目账户配置** |
| GET | `/wms/subject-account/get` | getSubjectAccount | **编号** |
| GET | `/wms/subject-account/page` | getSubjectAccountPage | **编号** |
| POST | `/wms/subject-account/senior` | getSubjectAccountSenior | **科目账户配置** |
| GET | `/wms/subject-account/export-excel` | exportSubjectAccountExcel | **科目账户配置** |
| POST | `/wms/subject-account/export-excel-senior` | exportSubjectAccountSeniorExcel | **科目账户配置** |
| GET | `/wms/subject-account/get-import-template` | importTemplate | **科目账户配置** |
| POST | `/wms/subject-account/import` | importExcel | **科目账户配置** |

### 供应商 (supplier)  (16个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/supplier/create` | createSupplier | **供应商** |
| PUT | `/wms/supplier/update` | updateSupplier | **供应商** |
| DELETE | `/wms/supplier/delete` | deleteSupplier | **供应商** |
| GET | `/wms/supplier/get` | getSupplier | **编号** |
| GET | `/wms/supplier/list` | getSupplierList | **编号** |
| GET | `/wms/supplier/page` | getSupplierPage | **供应商** |
| POST | `/wms/supplier/senior` | getSupplierSenior | **供应商** |
| GET | `/wms/supplier/pageSCP` | getSupplierPageSCP | **供应商** |
| POST | `/wms/supplier/seniorSCP` | getSupplierSeniorSCP | **供应商** |
| GET | `/wms/supplier/export-excel` | exportSupplierExcel | **供应商** |
| POST | `/wms/supplier/export-excel-senior` | exportSupplierExcel | **供应商** |
| POST | `/wms/supplier/export-excel-senior-SCP` | exportSupplierExcelSCP | **供应商** |
| GET | `/wms/supplier/export-excel-SCP` | exportSupplierExcelSCP | **供应商** |
| GET | `/wms/supplier/get-import-template` | importTemplate | **供应商** |
| POST | `/wms/supplier/import` | importExcel | **供应商** |
| GET | `/wms/supplier/listByCodes` | getSupplierListByCodes | **Excel 文件** |

### 供应商AP余额 (supplierApbalance)  (28个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/supplier-apbalance-adjust/create` | createSupplierApbalanceAdjust | **供应商余额调节** |
| PUT | `/wms/supplier-apbalance-adjust/update` | updateSupplierApbalanceAdjust | **供应商余额调节** |
| DELETE | `/wms/supplier-apbalance-adjust/delete` | deleteSupplierApbalanceAdjust | **供应商余额调节** |
| GET | `/wms/supplier-apbalance-adjust/get` | getSupplierApbalanceAdjust | **编号** |
| GET | `/wms/supplier-apbalance-adjust/page` | getSupplierApbalanceAdjustPage | **编号** |
| GET | `/wms/supplier-apbalance-adjust/getList` | getSupplierApbalanceAdjustList | **供应商余额调节** |
| POST | `/wms/supplier-apbalance-adjust/senior` | getSupplierApbalanceAdjustSenior | **供应商余额调节** |
| GET | `/wms/supplier-apbalance-adjust/export-excel` | exportSupplierApbalanceAdjustExcel | **供应商余额调节** |
| GET | `/wms/supplier-apbalance-adjust/get-import-template` | importTemplate | **供应商余额调节** |
| POST | `/wms/supplier-apbalance-detail/create` | createSupplierApbalanceDetail | **供应商余额明细子** |
| PUT | `/wms/supplier-apbalance-detail/update` | updateSupplierApbalanceDetail | **供应商余额明细子** |
| DELETE | `/wms/supplier-apbalance-detail/delete` | deleteSupplierApbalanceDetail | **供应商余额明细子** |
| GET | `/wms/supplier-apbalance-detail/get` | getSupplierApbalanceDetail | **编号** |
| GET | `/wms/supplier-apbalance-detail/page` | getSupplierApbalanceDetailPage | **编号** |
| POST | `/wms/supplier-apbalance-detail/senior` | getSupplierApbalanceDetailSenior | **供应商余额明细子** |
| GET | `/wms/supplier-apbalance-detail/export-excel` | exportSupplierApbalanceDetailExcel | **供应商余额明细子** |
| GET | `/wms/supplier-apbalance-detail/get-import-template` | importTemplate | **供应商余额明细子** |
| POST | `/wms/supplier-apbalance-main/create` | createSupplierApbalanceMain | **供应商余额明细主** |
| PUT | `/wms/supplier-apbalance-main/update` | updateSupplierApbalanceMain | **供应商余额明细主** |
| DELETE | `/wms/supplier-apbalance-main/delete` | deleteSupplierApbalanceMain | **供应商余额明细主** |
| GET | `/wms/supplier-apbalance-main/get` | getSupplierApbalanceMain | **编号** |
| GET | `/wms/supplier-apbalance-main/getPrintInfo` | queryRecordSupplierdeliverMain | **编号** |
| GET | `/wms/supplier-apbalance-main/page` | getSupplierApbalanceMainPage | **供应商余额明细主** |
| POST | `/wms/supplier-apbalance-main/senior` | getSupplierApbalanceMainSenior | **供应商余额明细主** |
| GET | `/wms/supplier-apbalance-main/confirmationPage` | getSupplierApbalanceMainConfirmationPage | **供应商余额明细主** |
| GET | `/wms/supplier-apbalance-main/export-excel` | exportSupplierApbalanceMainExcel | **供应商余额明细主** |
| POST | `/wms/supplier-apbalance-main/export-excel-senior` | exportSupplierApbalanceMainExcelSenior | **供应商余额明细主** |
| GET | `/wms/supplier-apbalance-main/get-import-template` | importTemplate | **供应商余额明细主** |

### 供应商账期 (supplierapbalancecalendar)  (12个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/supplier-apbalance-calendar/create` | createSupplierApbalanceCalendar | **询证函日历** |
| PUT | `/wms/supplier-apbalance-calendar/update` | updateSupplierApbalanceCalendar | **询证函日历** |
| DELETE | `/wms/supplier-apbalance-calendar/delete` | deleteSupplierApbalanceCalendar | **询证函日历** |
| GET | `/wms/supplier-apbalance-calendar/get` | getSupplierApbalanceCalendar | **编号** |
| GET | `/wms/supplier-apbalance-calendar/page` | getSupplierApbalanceCalendarPage | **编号** |
| POST | `/wms/supplier-apbalance-calendar/senior` | getSupplierApbalanceCalendarSenior | **询证函日历** |
| GET | `/wms/supplier-apbalance-calendar/export-excel` | exportSupplierApbalanceCalendarExcel | **询证函日历** |
| GET | `/wms/supplier-apbalance-calendar/get-import-template` | importTemplate | **询证函日历** |
| POST | `/wms/supplier-apbalance-calendar/import` | importExcel | **询证函日历** |
| GET | `/wms/supplier-apbalance-calendar/getMonth` | getSupplierApbalanceCalendarMonth | **Excel 文件** |
| GET | `/wms/supplier-apbalance-calendar/getDay` | getSupplierApbalanceCalendarDay | **询证函日历** |
| GET | `/wms/supplier-apbalance-calendar/getMonthDay` | getSupplierApbalanceCalendarMonthDay | **询证函日历** |

### 供应商周期 (suppliercycle)  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/supplier-cycle/create` | createSupplierCycle | **要货预测周期** |
| PUT | `/wms/supplier-cycle/update` | updateSupplierCycle | **要货预测周期** |
| DELETE | `/wms/supplier-cycle/delete` | deleteSupplierCycle | **要货预测周期** |
| GET | `/wms/supplier-cycle/get` | getSupplierCycle | **编号** |
| GET | `/wms/supplier-cycle/page` | getSupplierCyclePage | **编号** |
| POST | `/wms/supplier-cycle/senior` | getSupplierCycleSenior | **要货预测周期** |
| GET | `/wms/supplier-cycle/export-excel` | exportSupplierCycleExcel | **要货预测周期** |
| GET | `/wms/supplier-cycle/get-import-template` | importTemplate | **要货预测周期** |
| POST | `/wms/supplier-cycle/import` | importExcel | **要货预测周期** |

### 供应商交货记录 (supplierdeliverRecord)  (25个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/supplierdeliver-record-detail/create` | createSupplierdeliverRecordDetail | **供应商发货记录子** |
| PUT | `/wms/supplierdeliver-record-detail/update` | updateSupplierdeliverRecordDetail | **供应商发货记录子** |
| DELETE | `/wms/supplierdeliver-record-detail/delete` | deleteSupplierdeliverRecordDetail | **供应商发货记录子** |
| GET | `/wms/supplierdeliver-record-detail/get` | getSupplierdeliverRecordDetail | **编号** |
| GET | `/wms/supplierdeliver-record-detail/page` | getSupplierdeliverRecordDetailPage | **编号** |
| POST | `/wms/supplierdeliver-record-detail/senior` | getSupplierdeliverRecordDetailSenior | **供应商发货记录子** |
| GET | `/wms/supplierdeliver-record-detail/list` | getSupplierdeliverRecordDetailList | **供应商发货记录子** |
| GET | `/wms/supplierdeliver-record-detail/export-excel` | exportSupplierdeliverRecordDetailExcel | **编号列表** |
| GET | `/wms/supplierdeliver-record-detail/allList` | selectAllList | **供应商发货记录子** |
| GET | `/wms/supplierdeliver-record-detail/queryChildPickingNumberPage` | queryChildPickingNumberPage | **供应商发货记录子** |
| POST | `/wms/supplierdeliver-record-detail/queryChildPickingNumberSenior` | queryChildPickingNumberSenior | **供应商发货记录子** |
| POST | `/wms/supplierdeliver-record-main/createMq` | createSupplierdeliverRecordMainMq | **供应商发货记录主** |
| POST | `/wms/supplierdeliver-record-main/create` | createSupplierdeliverRecordMain | **供应商发货记录主** |
| PUT | `/wms/supplierdeliver-record-main/update` | updateSupplierdeliverRecordMain | **供应商发货记录主** |
| POST | `/wms/supplierdeliver-record-main/senior` | getSupplierdeliverRecordMainSenior | **供应商发货记录主** |
| DELETE | `/wms/supplierdeliver-record-main/delete` | deleteSupplierdeliverRecordMain | **供应商发货记录主** |
| GET | `/wms/supplierdeliver-record-main/get` | getSupplierdeliverRecordMain | **编号** |
| GET | `/wms/supplierdeliver-record-main/list` | getSupplierdeliverRecordMainList | **编号** |
| GET | `/wms/supplierdeliver-record-main/page` | getSupplierdeliverRecordMainPage | **编号列表** |
| POST | `/wms/supplierdeliver-record-main/abolish` | abolishSupplierdeliverRequestMain | **供应商发货记录主** |
| GET | `/wms/supplierdeliver-record-main/export-excel` | exportSupplierdeliverRecordMainExcel | **编号** |
| POST | `/wms/supplierdeliver-record-main/export-excel-senior` | exportPurchasereceiptRequestMainSeniorExcel | **供应商发货记录主** |
| GET | `/wms/supplierdeliver-record-main/getSupplierdeliverRecordById` | getSupplierdeliverRecordById | **供应商发货记录主** |
| POST | `/wms/supplierdeliver-record-main/createPurchasereceiptRequest` | createPurchasereceiptRequest | **编号** |
| POST | `/wms/supplierdeliver-record-main/cancelShipmentByAsnNumber` | cancelShipmentByAsnNumber | **供应商发货记录主** |

### 供应商交货申请 (supplierdeliverRequest)  (34个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/supplierdeliver-request-detail/create` | createSupplierdeliverRequestDetail | **供应商发货申请子** |
| PUT | `/wms/supplierdeliver-request-detail/update` | updateSupplierdeliverRequestDetail | **供应商发货申请子** |
| GET | `/wms/supplierdeliver-request-detail/page` | getSupplierdeliverRequestDetailPage | **供应商发货申请子** |
| POST | `/wms/supplierdeliver-request-detail/senior` | getSupplierdeliverRequestDetailSenior | **供应商发货申请子** |
| GET | `/wms/supplierdeliver-request-detail/generateLabelList` | generateLabelList | **供应商发货申请子** |
| GET | `/wms/supplierdeliver-request-detail/generateLabelParentList` | generateLabelParentList | **供应商发货申请子** |
| DELETE | `/wms/supplierdeliver-request-detail/delete` | deleteSupplierdeliverRequestDetail | **供应商发货申请子** |
| GET | `/wms/supplierdeliver-request-detail/get` | getSupplierdeliverRequestDetail | **编号** |
| GET | `/wms/supplierdeliver-request-detail/list` | getSupplierdeliverRequestDetailList | **编号** |
| GET | `/wms/supplierdeliver-request-detail/export-excel` | exportSupplierdeliverRequestDetailExcel | **编号列表** |
| POST | `/wms/supplierdeliver-request-main/create` | createSupplierdeliverRequestMain | **供应商发货申请主** |
| PUT | `/wms/supplierdeliver-request-main/update` | updateSupplierdeliverRequestMain | **供应商发货申请主** |
| DELETE | `/wms/supplierdeliver-request-main/delete` | deleteSupplierdeliverRequestMain | **供应商发货申请主** |
| GET | `/wms/supplierdeliver-request-main/get` | getSupplierdeliverRequestMain | **编号** |
| GET | `/wms/supplierdeliver-request-main/list` | getSupplierdeliverRequestMainList | **编号** |
| POST | `/wms/supplierdeliver-request-main/senior` | getSupplierdeliverRequestMainSenior | **编号列表** |
| GET | `/wms/supplierdeliver-request-main/page` | getSupplierdeliverRequestMainPage | **供应商发货申请主** |
| GET | `/wms/supplierdeliver-request-main/export-excel` | exportSupplierdeliverRequestMainExcel | **供应商发货申请主** |
| POST | `/wms/supplierdeliver-request-main/export-excel-senior` | exportSupplierdeliverRequestMainSeniorExcel | **供应商发货申请主** |
| GET | `/wms/supplierdeliver-request-main/get-import-template` | importTemplate | **供应商发货申请主** |
| POST | `/wms/supplierdeliver-request-main/import` | importExcel | **供应商发货申请主** |
| POST | `/wms/supplierdeliver-request-main/import` | importExcel | **Excel 文件** |
| POST | `/wms/supplierdeliver-request-main/close` | closeSupplierdeliverRequestMain | **Excel 文件** |
| POST | `/wms/supplierdeliver-request-main/open` | openSupplierdeliverRequestMain | **编号** |
| POST | `/wms/supplierdeliver-request-main/sub` | subSupplierdeliverRequestMain | **编号** |
| POST | `/wms/supplierdeliver-request-main/app` | witSupplierdeliverRequestMain | **编号** |
| POST | `/wms/supplierdeliver-request-main/rej` | rejSupplierdeliverRequestMain | **编号** |
| POST | `/wms/supplierdeliver-request-main/selfCheckReport` | selfCheckReport | **编号** |
| POST | `/wms/supplierdeliver-request-main/genLabel` | genLabel | **编号** |
| POST | `/wms/supplierdeliver-request-main/checkPackQty` | checkPackQty | **编号** |
| POST | `/wms/supplierdeliver-request-main/genRecords` | genRecords | **编号** |
| GET | `/wms/supplierdeliver-request-main/queryQualityInspection` | queryQualityInspection | **编号** |
| GET | `/wms/supplierdeliver-request-main/querySupplierResume` | querySupplierResume | **供应商发货申请主** |
| POST | `/wms/supplierdeliver-request-main/deleteOldLabels` | deleteOldLabels | **供应商发货申请主** |

### 来料质检 (supplierdeliverinspectiondetail)  (8个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/supplierdeliver-inspection-detail/create` | createSupplierdeliverInspectionDetail | **供应商发货申请质检信息子** |
| PUT | `/wms/supplierdeliver-inspection-detail/update` | updateSupplierdeliverInspectionDetail | **供应商发货申请质检信息子** |
| DELETE | `/wms/supplierdeliver-inspection-detail/delete` | deleteSupplierdeliverInspectionDetail | **供应商发货申请质检信息子** |
| GET | `/wms/supplierdeliver-inspection-detail/get` | getSupplierdeliverInspectionDetail | **编号** |
| GET | `/wms/supplierdeliver-inspection-detail/page` | getSupplierdeliverInspectionDetailPage | **编号** |
| POST | `/wms/supplierdeliver-inspection-detail/senior` | getSupplierdeliverInspectionDetailSenior | **供应商发货申请质检信息子** |
| GET | `/wms/supplierdeliver-inspection-detail/export-excel` | exportSupplierdeliverInspectionDetailExcel | **供应商发货申请质检信息子** |
| GET | `/wms/supplierdeliver-inspection-detail/queryByMasterId` | queryByMasterId | **供应商发货申请质检信息子** |

### 供应商发票记录 (supplierinvoiceRecord)  (23个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/supplierinvoice-record-deatil/create` | createSupplierinvoiceRecordDeatil | **供应商发票记录子** |
| PUT | `/wms/supplierinvoice-record-deatil/update` | updateSupplierinvoiceRecordDeatil | **供应商发票记录子** |
| DELETE | `/wms/supplierinvoice-record-deatil/delete` | deleteSupplierinvoiceRecordDeatil | **供应商发票记录子** |
| GET | `/wms/supplierinvoice-record-deatil/get` | getSupplierinvoiceRecordDeatil | **编号** |
| GET | `/wms/supplierinvoice-record-deatil/list` | getSupplierinvoiceRecordDeatilList | **编号** |
| GET | `/wms/supplierinvoice-record-deatil/page` | getSupplierinvoiceRecordDeatilPage | **编号列表** |
| POST | `/wms/supplierinvoice-record-deatil/senior` | getSupplierinvoiceRecordDeatilSenior | **供应商发票记录子** |
| GET | `/wms/supplierinvoice-record-deatil/export-excel` | exportSupplierinvoiceRecordDeatilExcel | **供应商发票记录子** |
| POST | `/wms/supplierinvoice-record-main/create` | createSupplierinvoiceRecordMain | **供应商发票记录主** |
| PUT | `/wms/supplierinvoice-record-main/update` | updateSupplierinvoiceRecordMain | **供应商发票记录主** |
| DELETE | `/wms/supplierinvoice-record-main/delete` | deleteSupplierinvoiceRecordMain | **供应商发票记录主** |
| GET | `/wms/supplierinvoice-record-main/get` | getSupplierinvoiceRecordMain | **编号** |
| GET | `/wms/supplierinvoice-record-main/list` | getSupplierinvoiceRecordMainList | **编号** |
| GET | `/wms/supplierinvoice-record-main/page` | getSupplierinvoiceRecordMainPage | **编号列表** |
| POST | `/wms/supplierinvoice-record-main/senior` | getSupplierinvoiceRecordMainSenior | **供应商发票记录主** |
| GET | `/wms/supplierinvoice-record-main/pageDiscrete` | getSupplierinvoiceRecordMainPageDiscrete | **供应商发票记录主** |
| POST | `/wms/supplierinvoice-record-main/seniorDiscrete` | getSupplierinvoiceRecordMainSeniorDiscrete | **供应商发票记录主** |
| GET | `/wms/supplierinvoice-record-main/export-excel` | exportSupplierinvoiceRecordMainExcel | **供应商发票记录主** |
| POST | `/wms/supplierinvoice-record-main/export-excel-senior` | exportSupplierinvoiceRecordMainSeniorExcel | **供应商发票记录主** |
| GET | `/wms/supplierinvoice-record-main/getInvoicePrintList` | getInvoicePrintList | **供应商发票记录主** |
| GET | `/wms/supplierinvoice-record-main/getFileByRequestNumber` | getFileByRequestNumber | **表数据id** |
| POST | `/wms/supplierinvoice-record-main/reverse` | reverse | **供应商发票记录主** |
| GET | `/wms/supplierinvoice-record-main/discreteIsRead` | discreteIsRead | **编号** |

### 供应商发票申请 (supplierinvoiceRequest)  (40个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/supplierinvoice-request-detail/create` | createSupplierinvoiceRequestDetail | **供应商发票申请子** |
| PUT | `/wms/supplierinvoice-request-detail/update` | updateSupplierinvoiceRequestDetail | **供应商发票申请子** |
| DELETE | `/wms/supplierinvoice-request-detail/delete` | deleteSupplierinvoiceRequestDetail | **供应商发票申请子** |
| GET | `/wms/supplierinvoice-request-detail/get` | getSupplierinvoiceRequestDetail | **编号** |
| GET | `/wms/supplierinvoice-request-detail/list` | getSupplierinvoiceRequestDetailList | **编号** |
| GET | `/wms/supplierinvoice-request-detail/page` | getSupplierinvoiceRequestDetailPage | **编号列表** |
| POST | `/wms/supplierinvoice-request-detail/senior` | getSupplierinvoiceRequestDetailSenior | **供应商发票申请子** |
| GET | `/wms/supplierinvoice-request-detail/export-excel` | exportSupplierinvoiceRequestDetailExcel | **供应商发票申请子** |
| GET | `/wms/supplierinvoice-request-detail/getPoNumberPoLineInfo` | getPoNumberPoLineInfo | **供应商发票申请子** |
| POST | `/wms/supplierinvoice-request-detail/getPoNumbersenior` | getPoNumbersenior | **供应商发票申请子** |
| GET | `/wms/supplierinvoice-request-detail/getInvoicedPage` | getInvoicedPage | **供应商发票申请子** |
| POST | `/wms/supplierinvoice-request-detail/getInvoicedSenior` | getInvoicedSenior | **供应商发票申请子** |
| POST | `/wms/supplierinvoice-request-main/create` | createSupplierinvoiceRequestMain | **供应商发票申请主** |
| PUT | `/wms/supplierinvoice-request-main/update` | updateSupplierinvoiceRequestMain | **供应商发票申请主** |
| DELETE | `/wms/supplierinvoice-request-main/delete` | deleteSupplierinvoiceRequestMain | **供应商发票申请主** |
| GET | `/wms/supplierinvoice-request-main/get` | getSupplierinvoiceRequestMain | **编号** |
| GET | `/wms/supplierinvoice-request-main/list` | getSupplierinvoiceRequestMainList | **编号** |
| GET | `/wms/supplierinvoice-request-main/getLoginUserRoleList` | getLoginUserRoleList | **编号列表** |
| GET | `/wms/supplierinvoice-request-main/page` | getSupplierinvoiceRequestMainPage | **供应商发票申请主** |
| POST | `/wms/supplierinvoice-request-main/senior` | getSupplierinvoiceRequestMainSenior | **供应商发票申请主** |
| POST | `/wms/supplierinvoice-request-main/export-excel` | exportSupplierinvoiceRequestMainExcel | **供应商发票申请主** |
| POST | `/wms/supplierinvoice-request-main/export-excel-senior` | exportSupplierinvoiceRequestMainSeniorExcel | **供应商发票申请主** |
| POST | `/wms/supplierinvoice-request-main/import` | importExcel | **供应商发票申请主** |
| GET | `/wms/supplierinvoice-request-main/get-import-template` | importTemplate | **Excel 文件** |
| POST | `/wms/supplierinvoice-request-main/close` | closeSupplierinvoiceRequestMain | **供应商发票申请主** |
| POST | `/wms/supplierinvoice-request-main/open` | openSupplierinvoiceRequestMain | **编号** |
| POST | `/wms/supplierinvoice-request-main/sub` | subSupplierinvoiceRequestMain | **编号** |
| POST | `/wms/supplierinvoice-request-main/app` | witSupplierinvoiceRequestMain | **编号** |
| POST | `/wms/supplierinvoice-request-main/rej` | rejSupplierinvoiceRequestMain | **编号** |
| POST | `/wms/supplierinvoice-request-main/invoiceSentOut` | invoiceSentOutSupplierinvoiceRequestMain | **供应商发票申请主** |
| POST | `/wms/supplierinvoice-request-main/financeApp` | financeAppSupplierinvoiceRequestMain | **编号** |
| POST | `/wms/supplierinvoice-request-main/financeRej` | financeRejSupplierinvoiceRequestMain | **编号** |
| POST | `/wms/supplierinvoice-request-main/genRecords` | genRecords | **编号** |
| POST | `/wms/supplierinvoice-request-main/querySupplierRecord` | querySupplierRecord | **编号** |
| POST | `/wms/supplierinvoice-request-main/querySupplierRecordByMasterId` | querySupplierRecordByMasterId | **编号** |
| GET | `/wms/supplierinvoice-request-main/export-excel-detail` | exportSupplierinvoiceRequestDetailExcel | **编号** |
| GET | `/wms/supplierinvoice-request-main/queryUserInfoByRoleCodePage` | queryUserInfoByRoleCodePage | **供应商发票申请主** |
| GET | `/wms/supplierinvoice-request-main/checkInvoicingCalendar` | checkInvoicingCalendar | **供应商发票申请主** |
| POST | `/wms/supplierinvoice-request-main/repeal` | repealSupplierinvoiceRequestMain | **供应商发票申请主** |
| GET | `/wms/supplierinvoice-request-main/computeById` | computeById | **编号** |

### 供应商已开票 (supplierinvoiceinvoiced)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/supplierinvoice-invoiced/create` | createSupplierinvoiceInvoiced | **待开票** |
| PUT | `/wms/supplierinvoice-invoiced/update` | updateSupplierinvoiceInvoiced | **待开票** |
| DELETE | `/wms/supplierinvoice-invoiced/delete` | deleteSupplierinvoiceInvoiced | **待开票** |
| GET | `/wms/supplierinvoice-invoiced/get` | getSupplierinvoiceInvoiced | **编号** |
| GET | `/wms/supplierinvoice-invoiced/list` | getSupplierinvoiceInvoicedList | **编号** |
| GET | `/wms/supplierinvoice-invoiced/page` | getSupplierinvoiceInvoicedPage | **编号列表** |
| GET | `/wms/supplierinvoice-invoiced/pageDiscrete` | getSupplierinvoiceInvoicedPageDiscrete | **待开票** |
| GET | `/wms/supplierinvoice-invoiced/deletedPage` | getSupplierinvoiceInvoicedDeletedPage | **待开票** |
| POST | `/wms/supplierinvoice-invoiced/senior` | getSupplierinvoiceInvoicedSenior | **待开票** |
| POST | `/wms/supplierinvoice-invoiced/seniorDiscrete` | getSupplierinvoiceInvoicedSeniorDiscrete | **待开票** |
| GET | `/wms/supplierinvoice-invoiced/export-excel-schedule` | exportSupplierinvoiceInvoicedExcelSchedule | **待开票** |
| POST | `/wms/supplierinvoice-invoiced/export-excel-schedule-senior` | exportSupplierinvoiceInvoicedExcelScheduleSenior | **待开票** |
| GET | `/wms/supplierinvoice-invoiced/export-excel-discrete` | exportSupplierinvoiceInvoicedExcelDiscrete | **待开票** |
| POST | `/wms/supplierinvoice-invoiced/export-excel-discrete-senior` | exportSupplierinvoiceInvoicedExcelDiscreteSenior | **待开票** |
| GET | `/wms/supplierinvoice-invoiced/get-import-template` | importTemplate | **待开票** |
| POST | `/wms/supplierinvoice-invoiced/import` | importExcel | **待开票** |
| POST | `/wms/supplierinvoice-invoiced/agree` | witSupplierinvoiceAgree | **Excel 文件** |
| POST | `/wms/supplierinvoice-invoiced/recevery` | receverySupplierinvoice | **编号** |
| POST | `/wms/supplierinvoice-invoiced/refuse` | witSupplierinvoiceRefuse | **编号** |

### 供应商货品 (supplieritem)  (23个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/supplieritem/create` | createSupplieritem | **供应商物料** |
| POST | `/wms/supplieritem/createSCP` | createSCP | **供应商物料** |
| PUT | `/wms/supplieritem/update` | update | **供应商物料** |
| PUT | `/wms/supplieritem/updateSCP` | updateSCP | **供应商物料** |
| DELETE | `/wms/supplieritem/delete` | deleteSupplieritem | **供应商物料** |
| DELETE | `/wms/supplieritem/deleteSCP` | deleteSCP | **编号** |
| GET | `/wms/supplieritem/get` | getSupplieritem | **编号** |
| GET | `/wms/supplieritem/pageSCP` | getSupplieritemPageSCP | **编号** |
| POST | `/wms/supplieritem/seniorSCP` | getSupplieritemSeniorSCP | **供应商物料** |
| GET | `/wms/supplieritem/page` | getSupplieritemPage | **供应商物料** |
| POST | `/wms/supplieritem/senior` | getSupplieritemSenior | **供应商物料** |
| GET | `/wms/supplieritem/pageItembasicTypeToSupplieritem` | selectItembasicTypeToSupplieritem | **供应商物料** |
| POST | `/wms/supplieritem/pageItembasicTypeToSupplieritemSenior` | selectItembasicTypeToSupplieritemSenior | **供应商物料** |
| GET | `/wms/supplieritem/export-excel` | exportSupplieritemExcel | **供应商物料** |
| POST | `/wms/supplieritem/export-excel-senior` | exportSupplieritemExcel | **供应商物料** |
| GET | `/wms/supplieritem/export-excel-SCP` | exportSupplieritemExcelSCP | **供应商物料** |
| POST | `/wms/supplieritem/export-excel-senior-SCP` | exportSupplieritemExcelSCP | **供应商物料** |
| GET | `/wms/supplieritem/get-import-template` | importTemplate | **供应商物料** |
| POST | `/wms/supplieritem/import` | importExcel | **供应商物料** |
| POST | `/wms/supplieritemgetDefaultLocationCode` | getDefaultLocationCode | **Excel 文件** |
| GET | `/wms/supplieritem/listByCodes` | getSupplieritemListByCodes | **供应商物料** |
| GET | `/wms/supplieritem/querySupplierByCode` | querySupplierByCode | **供应商代码编码** |
| GET | `/wms/supplieritem/querySupplierByCodeAndType` | querySupplierByCodeAndType | **物料编码** |

### 供应商用户 (supplieruser)  (12个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/supplier-user/create` | createSupplierUser | **供应商用户关联信息** |
| PUT | `/wms/supplier-user/update` | updateSupplierUser | **供应商用户关联信息** |
| DELETE | `/wms/supplier-user/delete` | deleteSupplierUser | **供应商用户关联信息** |
| GET | `/wms/supplier-user/get` | getSupplierUser | **编号** |
| GET | `/wms/supplier-user/list` | getSupplierUserList | **编号** |
| GET | `/wms/supplier-user/page` | getSupplierUserPage | **编号列表** |
| POST | `/wms/supplier-user/senior` | getSupplierUserSenior | **供应商用户关联信息** |
| GET | `/wms/supplier-user/export-excel` | exportSupplierUserExcel | **供应商用户关联信息** |
| POST | `/wms/supplier-user/export-excel-senior` | exportSupplierUserExcel | **供应商用户关联信息** |
| GET | `/wms/supplier-user/get-import-template` | importTemplate | **供应商用户关联信息** |
| POST | `/wms/supplier-user/import` | importExcel | **供应商用户关联信息** |
| GET | `/wms/supplier-user/getSupplierUserList` | getSupplierUserList | **Excel 文件** |

### 系统安装包 (systemInstallPackage)  (8个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/system-install-package/create` | createSystemInstallPackage | **安装包信息** |
| PUT | `/wms/system-install-package/update` | updateSystemInstallPackage | **安装包信息** |
| DELETE | `/wms/system-install-package/delete` | deleteSystemInstallPackage | **安装包信息** |
| GET | `/wms/system-install-package/get` | getSystemInstallPackage | **编号** |
| GET | `/wms/system-install-package/list` | getSystemInstallPackageList | **编号** |
| GET | `/wms/system-install-package/page` | getSystemInstallPackagePage | **编号列表** |
| GET | `/wms/system-install-package/returnNewFile` | returnNewFileSystemInstallPackage | **安装包信息** |
| GET | `/wms/system-install-package/downloadApk` | downloadApk | **安装包信息** |

### 系统日历 (systemcalendar)  (10个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/systemcalendar/create` | createSystemcalendar | **系统日历** |
| PUT | `/wms/systemcalendar/update` | updateSystemcalendar | **系统日历** |
| DELETE | `/wms/systemcalendar/delete` | deleteSystemcalendar | **系统日历** |
| GET | `/wms/systemcalendar/get` | getSystemcalendar | **编号** |
| GET | `/wms/systemcalendar/page` | getSystemcalendarPage | **编号** |
| POST | `/wms/systemcalendar/senior` | getSystemcalendarSenior | **系统日历** |
| GET | `/wms/systemcalendar/export-excel` | exportSystemcalendarExcel | **系统日历** |
| POST | `/wms/systemcalendar/export-excel-senior` | exportSystemcalendarExcel | **系统日历** |
| GET | `/wms/systemcalendar/get-import-template` | importTemplate | **系统日历** |
| POST | `/wms/systemcalendar/import` | importExcel | **系统日历** |

### 班组 (team)  (13个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/team/create` | createTeam | **班组** |
| PUT | `/wms/team/update` | updateTeam | **班组** |
| DELETE | `/wms/team/delete` | deleteTeam | **班组** |
| GET | `/wms/team/get` | getTeam | **编号** |
| GET | `/wms/team/page` | getTeamPage | **编号** |
| GET | `/wms/team/getPageChildren` | getPageChildren | **班组** |
| POST | `/wms/team/senior` | getTeamSenior | **班组** |
| POST | `/wms/team/queryTeamUserByCode` | seniorTeamUserByCode | **班组** |
| GET | `/wms/team/export-excel` | exportTeamExcel | **班组** |
| POST | `/wms/team/export-excel-senior` | exportTeamExcel | **班组** |
| GET | `/wms/team/get-import-template` | importTemplate | **班组** |
| POST | `/wms/team/import` | importExcel | **班组** |
| GET | `/wms/team/noPage` | getTeamNoPage | **Excel 文件** |

### 事务 (transaction)  (9个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/transaction/create` | createTransaction | **库存事务** |
| POST | `/wms/transaction/createTransactionAndTransferLog` | createTransactionAndTransferLog | **库存事务** |
| GET | `/wms/transaction/page` | getTransactionPage | **库存事务** |
| GET | `/wms/transaction/page_balance` | getTransactionBalancePage | **库存事务** |
| POST | `/wms/transaction/senior` | getTransactionSenior | **库存事务** |
| GET | `/wms/transaction/export-excel` | exportTransactionExcel | **库存事务** |
| GET | `/wms/transaction/export-excel-checkCnt` | exportTransactionExcelCheck | **库存事务** |
| POST | `/wms/transaction/export-excel-senior` | exportTransactionSeniorExcel | **库存事务** |
| POST | `/wms/transaction/export-excel-senior-checkCnt` | exportTransactionSeniorExcelCheck | **库存事务** |

### 事务类型 (transactiontype)  (12个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/transactiontype/create` | createTransactiontype | **事务类型** |
| PUT | `/wms/transactiontype/update` | updateTransactiontype | **事务类型** |
| DELETE | `/wms/transactiontype/delete` | deleteTransactiontype | **事务类型** |
| GET | `/wms/transactiontype/get` | getTransactiontype | **编号** |
| GET | `/wms/transactiontype/list` | getTransactiontypeList | **编号** |
| GET | `/wms/transactiontype/page` | getTransactiontypePage | **编号列表** |
| POST | `/wms/transactiontype/senior` | getTransactiontypeSenior | **事务类型** |
| GET | `/wms/transactiontype/export-excel` | exportTransactiontypeExcel | **事务类型** |
| POST | `/wms/transactiontype/export-excel-senior` | exportTransactiontypeSeniorExcel | **事务类型** |
| GET | `/wms/transactiontype/get-import-template` | importTemplate | **事务类型** |
| POST | `/wms/transactiontype/import` | importExcel | **事务类型** |
| GET | `/wms/transactiontype/ListByCode` | ListByCode | **Excel 文件** |

### 调拨出库任务 (transferissueJob)  (23个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/transferissue-job-detail/create` | createTransferissueJobDetail | **调拨出库任务子** |
| PUT | `/wms/transferissue-job-detail/update` | updateTransferissueJobDetail | **调拨出库任务子** |
| DELETE | `/wms/transferissue-job-detail/delete` | deleteTransferissueJobDetail | **调拨出库任务子** |
| GET | `/wms/transferissue-job-detail/get` | getTransferissueJobDetail | **编号** |
| GET | `/wms/transferissue-job-detail/list` | getTransferissueJobDetailList | **编号** |
| GET | `/wms/transferissue-job-detail/page` | getTransferissueJobDetailPage | **编号列表** |
| POST | `/wms/transferissue-job-detail/senior` | getTransferissueJobDetailSenior | **调拨出库任务子** |
| GET | `/wms/transferissue-job-detail/export-excel` | exportTransferissueJobDetailExcel | **调拨出库任务子** |
| POST | `/wms/transferissue-job-main/create` | createTransferissueJobMain | **调拨出库任务主** |
| PUT | `/wms/transferissue-job-main/update` | updateTransferissueJobMain | **调拨出库任务主** |
| DELETE | `/wms/transferissue-job-main/delete` | deleteTransferissueJobMain | **调拨出库任务主** |
| GET | `/wms/transferissue-job-main/get` | getTransferissueJobMain | **编号** |
| GET | `/wms/transferissue-job-main/list` | getTransferissueJobMainList | **编号** |
| GET | `/wms/transferissue-job-main/page` | getTransferissueJobMainPage | **编号列表** |
| POST | `/wms/transferissue-job-main/senior` | getTransferissueJobMainSenior | **调拨出库任务主** |
| GET | `/wms/transferissue-job-main/export-excel` | exportTransferissueJobMainExcel | **调拨出库任务主** |
| POST | `/wms/transferissue-job-main/export-excel-senior` | exportTransferissueJobMainSeniorExcel | **调拨出库任务主** |
| GET | `/wms/transferissue-job-main/getTransferissueJobById` | getTransferissueJobById | **调拨出库任务主** |
| POST | `/wms/transferissue-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/transferissue-job-main/accept` | acceptTransferissueJobMain | **类型数组** |
| PUT | `/wms/transferissue-job-main/abandon` | abandonTransferissueJobMain | **调拨出库任务主** |
| PUT | `/wms/transferissue-job-main/close` | closeTransferissueJobMain | **调拨出库任务主** |
| PUT | `/wms/transferissue-job-main/execute` | executeTransferissueJobMain | **调拨出库任务主** |

### 调拨出库记录 (transferissueRecord)  (18个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/transferissue-record-detail/create` | createTransferissueRecordDetail | **调拨出库记录子** |
| PUT | `/wms/transferissue-record-detail/update` | updateTransferissueRecordDetail | **调拨出库记录子** |
| DELETE | `/wms/transferissue-record-detail/delete` | deleteTransferissueRecordDetail | **调拨出库记录子** |
| GET | `/wms/transferissue-record-detail/get` | getTransferissueRecordDetail | **编号** |
| GET | `/wms/transferissue-record-detail/list` | getTransferissueRecordDetailList | **编号** |
| POST | `/wms/transferissue-record-detail/senior` | getTransferissueRecordDetailSenior | **编号列表** |
| GET | `/wms/transferissue-record-detail/page` | getTransferissueRecordDetailPage | **调拨出库记录子** |
| GET | `/wms/transferissue-record-detail/export-excel` | exportTransferissueRecordDetailExcel | **调拨出库记录子** |
| POST | `/wms/transferissue-record-main/create` | createTransferissueRecordMain | **调拨出库记录主** |
| PUT | `/wms/transferissue-record-main/update` | updateTransferissueRecordMain | **调拨出库记录主** |
| DELETE | `/wms/transferissue-record-main/delete` | deleteTransferissueRecordMain | **调拨出库记录主** |
| GET | `/wms/transferissue-record-main/get` | getTransferissueRecordMain | **编号** |
| GET | `/wms/transferissue-record-main/list` | getTransferissueRecordMainList | **编号** |
| GET | `/wms/transferissue-record-main/page` | getTransferissueRecordMainPage | **编号列表** |
| POST | `/wms/transferissue-record-main/senior` | getTransferissueRecordMainSenior | **调拨出库记录主** |
| GET | `/wms/transferissue-record-main/export-excel` | exportTransferissueRecordMainExcel | **调拨出库记录主** |
| POST | `/wms/transferissue-record-main/export-excel-senior` | exportTransferissueRequestMainSeniorExcel | **调拨出库记录主** |
| GET | `/wms/transferissue-record-main/getDetailInfoById` | getDetailInfoById | **调拨出库记录主** |

### 调拨出库申请 (transferissueRequest)  (25个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/transferissue-request-detail/create` | createTransferissueRequestDetail | **调拨出库申请子** |
| PUT | `/wms/transferissue-request-detail/update` | updateTransferissueRequestDetail | **调拨出库申请子** |
| DELETE | `/wms/transferissue-request-detail/delete` | deleteTransferissueRequestDetail | **调拨出库申请子** |
| GET | `/wms/transferissue-request-detail/get` | getTransferissueRequestDetail | **编号** |
| GET | `/wms/transferissue-request-detail/list` | getTransferissueRequestDetailList | **编号** |
| GET | `/wms/transferissue-request-detail/page` | getTransferissueRequestDetailPage | **编号列表** |
| POST | `/wms/transferissue-request-detail/senior` | getTransferissueRequestDetailSenior | **调拨出库申请子** |
| GET | `/wms/transferissue-request-detail/export-excel` | exportTransferissueRequestDetailExcel | **调拨出库申请子** |
| POST | `/wms/transferissue-request-main/create` | createTransferissueRequestMain | **调拨出库申请主** |
| PUT | `/wms/transferissue-request-main/update` | updateTransferissueRequestMain | **调拨出库申请主** |
| DELETE | `/wms/transferissue-request-main/delete` | deleteTransferissueRequestMain | **调拨出库申请主** |
| GET | `/wms/transferissue-request-main/get` | getTransferissueRequestMain | **编号** |
| GET | `/wms/transferissue-request-main/list` | getTransferissueRequestMainList | **编号** |
| GET | `/wms/transferissue-request-main/page` | getTransferissueRequestMainPage | **编号列表** |
| POST | `/wms/transferissue-request-main/senior` | getTransferissueRequestMainSenior | **调拨出库申请主** |
| GET | `/wms/transferissue-request-main/export-excel` | exportTransferissueRequestMainExcel | **调拨出库申请主** |
| POST | `/wms/transferissue-request-main/export-excel-senior` | exportTransferissueRequestMainSeniorExcel | **调拨出库申请主** |
| GET | `/wms/transferissue-request-main/get-import-template` | importTemplate | **调拨出库申请主** |
| POST | `/wms/transferissue-request-main/import` | importExcel | **调拨出库申请主** |
| PUT | `/wms/transferissue-request-main/close` | closeTransferissueRequestMain | **Excel 文件** |
| PUT | `/wms/transferissue-request-main/reAdd` | reAddTransferissueRequestMain | **编号** |
| PUT | `/wms/transferissue-request-main/submit` | submitTransferissueRequestMain | **编号** |
| PUT | `/wms/transferissue-request-main/refused` | abortTransferissueRequestMain | **编号** |
| PUT | `/wms/transferissue-request-main/agree` | agreeTransferissueRequestMain | **编号** |
| PUT | `/wms/transferissue-request-main/handle` | handleTransferissueRequestMain | **编号** |

### 调拨日志 (transferlog)  (6个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| GET | `/wms/transferlog/get` | getTransferlog | **库存转移日志** |
| GET | `/wms/transferlog/list` | getTransferlogList | **编号** |
| GET | `/wms/transferlog/page` | getTransferlogPage | **编号列表** |
| POST | `/wms/transferlog/senior` | getTransferlogSenior | **库存转移日志** |
| GET | `/wms/transferlog/export-excel` | exportTransferlogExcel | **库存转移日志** |
| POST | `/wms/transferlog/export-excel-senior` | exportTransferlogSeniorExcel | **库存转移日志** |

### 调拨入库任务 (transferreceiptJob)  (23个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/transferreceipt-job-detail/create` | createTransferreceiptJobDetail | **调拨入库任务子** |
| PUT | `/wms/transferreceipt-job-detail/update` | updateTransferreceiptJobDetail | **调拨入库任务子** |
| DELETE | `/wms/transferreceipt-job-detail/delete` | deleteTransferreceiptJobDetail | **调拨入库任务子** |
| GET | `/wms/transferreceipt-job-detail/get` | getTransferreceiptJobDetail | **编号** |
| GET | `/wms/transferreceipt-job-detail/list` | getTransferreceiptJobDetailList | **编号** |
| GET | `/wms/transferreceipt-job-detail/page` | getTransferreceiptJobDetailPage | **编号列表** |
| POST | `/wms/transferreceipt-job-detail/senior` | getTransferreceiptJobDetailSenior | **调拨入库任务子** |
| GET | `/wms/transferreceipt-job-detail/export-excel` | exportTransferreceiptJobDetailExcel | **调拨入库任务子** |
| POST | `/wms/transferreceipt-job-main/create` | createTransferreceiptJobMain | **调拨入库任务主** |
| PUT | `/wms/transferreceipt-job-main/update` | updateTransferreceiptJobMain | **调拨入库任务主** |
| DELETE | `/wms/transferreceipt-job-main/delete` | deleteTransferreceiptJobMain | **调拨入库任务主** |
| GET | `/wms/transferreceipt-job-main/get` | getTransferreceiptJobMain | **编号** |
| GET | `/wms/transferreceipt-job-main/list` | getTransferreceiptJobMainList | **编号** |
| GET | `/wms/transferreceipt-job-main/page` | getTransferreceiptJobMainPage | **编号列表** |
| POST | `/wms/transferreceipt-job-main/senior` | getTransferreceiptJobMainSenior | **调拨入库任务主** |
| GET | `/wms/transferreceipt-job-main/export-excel` | exportTransferreceiptJobMainExcel | **调拨入库任务主** |
| POST | `/wms/transferreceipt-job-main/export-excel-senior` | exportTransferreceiptJobSeniorExcel | **调拨入库任务主** |
| GET | `/wms/transferreceipt-job-main/getTransferreceiptJobById` | getTransferreceiptJobById | **调拨入库任务主** |
| POST | `/wms/transferreceipt-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/transferreceipt-job-main/accept` | acceptTransferreceiptJobMain | **类型数组** |
| PUT | `/wms/transferreceipt-job-main/abandon` | abandonTransferreceiptJobMain | **调拨入库任务主** |
| PUT | `/wms/transferreceipt-job-main/close` | closeTransferreceiptJobMain | **调拨入库任务主** |
| PUT | `/wms/transferreceipt-job-main/execute` | executeTransferreceiptJobMain | **调拨入库任务主** |

### 调拨入库记录 (transferreceiptRecord)  (18个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/transferreceipt-record-detail/create` | createTransferreceiptRecordDetail | **调拨入库记录子** |
| PUT | `/wms/transferreceipt-record-detail/update` | updateTransferreceiptRecordDetail | **调拨入库记录子** |
| DELETE | `/wms/transferreceipt-record-detail/delete` | deleteTransferreceiptRecordDetail | **调拨入库记录子** |
| GET | `/wms/transferreceipt-record-detail/get` | getTransferreceiptRecordDetail | **编号** |
| GET | `/wms/transferreceipt-record-detail/list` | getTransferreceiptRecordDetailList | **编号** |
| GET | `/wms/transferreceipt-record-detail/page` | getTransferreceiptRecordDetailPage | **编号列表** |
| POST | `/wms/transferreceipt-record-detail/senior` | getTransferreceiptRecordDetailSenior | **调拨入库记录子** |
| GET | `/wms/transferreceipt-record-detail/export-excel` | exportTransferreceiptRecordDetailExcel | **调拨入库记录子** |
| POST | `/wms/transferreceipt-record-main/create` | createTransferreceiptRecordMain | **调拨入库记录主** |
| PUT | `/wms/transferreceipt-record-main/update` | updateTransferreceiptRecordMain | **调拨入库记录主** |
| DELETE | `/wms/transferreceipt-record-main/delete` | deleteTransferreceiptRecordMain | **调拨入库记录主** |
| GET | `/wms/transferreceipt-record-main/get` | getTransferreceiptRecordMain | **编号** |
| GET | `/wms/transferreceipt-record-main/list` | getTransferreceiptRecordMainList | **编号** |
| GET | `/wms/transferreceipt-record-main/page` | getTransferreceiptRecordMainPage | **编号列表** |
| POST | `/wms/transferreceipt-record-main/senior` | getTransferreceiptRecordMainSenior | **调拨入库记录主** |
| GET | `/wms/transferreceipt-record-main/export-excel` | exportTransferreceiptRecordMainExcel | **调拨入库记录主** |
| POST | `/wms/transferreceipt-record-main/export-excel-senior` | exportTransferreceiptRecordMainSeniorExcel | **调拨入库记录主** |
| GET | `/wms/transferreceipt-record-main/getDetailInfoById` | getDetailInfoById | **调拨入库记录主** |

### 调拨入库申请 (transferreceiptRequest)  (23个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/transferreceipt-request-detail/create` | createTransferreceiptRequestDetail | **调拨入库申请子** |
| PUT | `/wms/transferreceipt-request-detail/update` | updateTransferreceiptRequestDetail | **调拨入库申请子** |
| DELETE | `/wms/transferreceipt-request-detail/delete` | deleteTransferreceiptRequestDetail | **调拨入库申请子** |
| GET | `/wms/transferreceipt-request-detail/get` | getTransferreceiptRequestDetail | **编号** |
| GET | `/wms/transferreceipt-request-detail/list` | getTransferreceiptRequestDetailList | **编号** |
| GET | `/wms/transferreceipt-request-detail/page` | getTransferreceiptRequestDetailPage | **编号列表** |
| POST | `/wms/transferreceipt-request-detail/senior` | getTransferreceiptRequestDetailSenior | **调拨入库申请子** |
| GET | `/wms/transferreceipt-request-detail/export-excel` | exportTransferreceiptRequestDetailExcel | **调拨入库申请子** |
| POST | `/wms/transferreceipt-request-main/create` | createTransferreceiptRequestMain | **调拨入库申请主** |
| PUT | `/wms/transferreceipt-request-main/update` | updateTransferreceiptRequestMain | **调拨入库申请主** |
| DELETE | `/wms/transferreceipt-request-main/delete` | deleteTransferreceiptRequestMain | **调拨入库申请主** |
| GET | `/wms/transferreceipt-request-main/get` | getTransferreceiptRequestMain | **编号** |
| GET | `/wms/transferreceipt-request-main/list` | getTransferreceiptRequestMainList | **编号** |
| GET | `/wms/transferreceipt-request-main/page` | getTransferreceiptRequestMainPage | **编号列表** |
| POST | `/wms/transferreceipt-request-main/senior` | getTransferreceiptRequestMainSenior | **调拨入库申请主** |
| GET | `/wms/transferreceipt-request-main/export-excel` | exportTransferreceiptRequestMainExcel | **调拨入库申请主** |
| POST | `/wms/transferreceipt-request-main/export-excel-senior` | exportTransferreceiptRequestMainSeniorExcel | **调拨入库申请主** |
| PUT | `/wms/transferreceipt-request-main/close` | closeTransferreceiptRequestMain | **调拨入库申请主** |
| PUT | `/wms/transferreceipt-request-main/reAdd` | reAddTransferreceiptRequestMain | **编号** |
| PUT | `/wms/transferreceipt-request-main/submit` | submitTransferreceiptRequestMain | **编号** |
| PUT | `/wms/transferreceipt-request-main/refused` | abortTransferreceiptRequestMain | **编号** |
| PUT | `/wms/transferreceipt-request-main/agree` | agreeTransferreceiptRequestMain | **编号** |
| PUT | `/wms/transferreceipt-request-main/handle` | handleTransferreceiptRequestMain | **编号** |

### 非计划领料任务 (unplannedissueJob)  (23个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/unplannedissue-job-detail/create` | createUnplannedissueJobDetail | **计划外出库任务子** |
| PUT | `/wms/unplannedissue-job-detail/update` | updateUnplannedissueJobDetail | **计划外出库任务子** |
| DELETE | `/wms/unplannedissue-job-detail/delete` | deleteUnplannedissueJobDetail | **计划外出库任务子** |
| GET | `/wms/unplannedissue-job-detail/get` | getUnplannedissueJobDetail | **编号** |
| GET | `/wms/unplannedissue-job-detail/list` | getUnplannedissueJobDetailList | **编号** |
| GET | `/wms/unplannedissue-job-detail/page` | getUnplannedissueJobDetailPage | **编号列表** |
| POST | `/wms/unplannedissue-job-detail/senior` | getUnplannedissueJobDetailSenior | **计划外出库任务子** |
| GET | `/wms/unplannedissue-job-detail/export-excel` | exportUnplannedissueJobDetailExcel | **计划外出库任务子** |
| POST | `/wms/unplannedissue-job-main/create` | createUnplannedissueJobMain | **计划外出库任务主** |
| PUT | `/wms/unplannedissue-job-main/update` | updateUnplannedissueJobMain | **计划外出库任务主** |
| DELETE | `/wms/unplannedissue-job-main/delete` | deleteUnplannedissueJobMain | **计划外出库任务主** |
| GET | `/wms/unplannedissue-job-main/get` | getUnplannedissueJobMain | **编号** |
| GET | `/wms/unplannedissue-job-main/list` | getUnplannedissueJobMainList | **编号** |
| GET | `/wms/unplannedissue-job-main/page` | getUnplannedissueJobMainPage | **编号列表** |
| POST | `/wms/unplannedissue-job-main/senior` | getUnplannedissueJobMainSenior | **计划外出库任务主** |
| GET | `/wms/unplannedissue-job-main/export-excel` | exportUnplannedissueJobMainExcel | **计划外出库任务主** |
| POST | `/wms/unplannedissue-job-main/export-excel-senior` | exportUnplannedreceiptRequestMainSeniorExcel | **计划外出库任务主** |
| GET | `/wms/unplannedissue-job-main/getUnplannedissueJobById` | getUnplannedissueJobById | **计划外出库任务主** |
| POST | `/wms/unplannedissue-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/unplannedissue-job-main/accept` | acceptUnplannedissueJobMain | **类型数组** |
| PUT | `/wms/unplannedissue-job-main/abandon` | abandonUnplannedissueJobMain | **计划外出库任务主** |
| PUT | `/wms/unplannedissue-job-main/close` | closeUnplannedissueJobMain | **计划外出库任务主** |
| PUT | `/wms/unplannedissue-job-main/execute` | executeUnplannedissueJobMain | **计划外出库任务主** |

### 非计划领料记录 (unplannedissueRecord)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/unplannedissue-record-detail/create` | createUnplannedissueRecordDetail | **计划外出库记录子** |
| PUT | `/wms/unplannedissue-record-detail/update` | updateUnplannedissueRecordDetail | **计划外出库记录子** |
| DELETE | `/wms/unplannedissue-record-detail/delete` | deleteUnplannedissueRecordDetail | **计划外出库记录子** |
| GET | `/wms/unplannedissue-record-detail/get` | getUnplannedissueRecordDetail | **编号** |
| GET | `/wms/unplannedissue-record-detail/list` | getUnplannedissueRecordDetailList | **编号** |
| GET | `/wms/unplannedissue-record-detail/page` | getUnplannedissueRecordDetailPage | **编号列表** |
| POST | `/wms/unplannedissue-record-detail/senior` | getUnplannedissueRecordDetailSenior | **计划外出库记录子** |
| GET | `/wms/unplannedissue-record-detail/export-excel` | exportUnplannedissueRecordDetailExcel | **计划外出库记录子** |
| POST | `/wms/unplannedissue-record-main/create` | createUnplannedissueRecordMain | **计划外出库记录主** |
| PUT | `/wms/unplannedissue-record-main/update` | updateUnplannedissueRecordMain | **计划外出库记录主** |
| DELETE | `/wms/unplannedissue-record-main/delete` | deleteUnplannedissueRecordMain | **计划外出库记录主** |
| GET | `/wms/unplannedissue-record-main/get` | getUnplannedissueRecordMain | **编号** |
| GET | `/wms/unplannedissue-record-main/list` | getUnplannedissueRecordMainList | **编号** |
| GET | `/wms/unplannedissue-record-main/page` | getUnplannedissueRecordMainPage | **编号列表** |
| POST | `/wms/unplannedissue-record-main/senior` | getUnplannedissueRecordMainSenior | **计划外出库记录主** |
| GET | `/wms/unplannedissue-record-main/export-excel` | exportUnplannedissueRecordMainExcel | **计划外出库记录主** |
| POST | `/wms/unplannedissue-record-main/export-excel-senior` | exportUnplannedreceiptRecordMainSeniorExcel | **计划外出库记录主** |
| GET | `/wms/unplannedissue-record-main/getDetailInfoById` | getDetailInfoById | **计划外出库记录主** |
| GET | `/wms/unplannedissue-record-main/revoke` | revoke | **计划外出库记录主** |

### 非计划领料申请 (unplannedissueRequest)  (28个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/unplannedissue-request-detail/create` | createUnplannedissueRequestDetail | **计划外出库申请子** |
| PUT | `/wms/unplannedissue-request-detail/update` | updateUnplannedissueRequestDetail | **计划外出库申请子** |
| DELETE | `/wms/unplannedissue-request-detail/delete` | deleteUnplannedissueRequestDetail | **计划外出库申请子** |
| GET | `/wms/unplannedissue-request-detail/get` | getUnplannedissueRequestDetail | **编号** |
| GET | `/wms/unplannedissue-request-detail/list` | getUnplannedissueRequestDetailList | **编号** |
| GET | `/wms/unplannedissue-request-detail/page` | getUnplannedissueRequestDetailPage | **编号列表** |
| POST | `/wms/unplannedissue-request-detail/senior` | getUnplannedissueRequestDetailSenior | **计划外出库申请子** |
| GET | `/wms/unplannedissue-request-detail/export-excel` | exportUnplannedissueRequestDetailExcel | **计划外出库申请子** |
| POST | `/wms/unplannedissue-request-main/create` | createUnplannedissueRequestMain | **计划外出库申请主** |
| PUT | `/wms/unplannedissue-request-main/update` | updateUnplannedissueRequestMain | **计划外出库申请主** |
| DELETE | `/wms/unplannedissue-request-main/delete` | deleteUnplannedissueRequestMain | **计划外出库申请主** |
| GET | `/wms/unplannedissue-request-main/get` | getUnplannedissueRequestMain | **编号** |
| GET | `/wms/unplannedissue-request-main/list` | getUnplannedissueRequestMainList | **编号** |
| GET | `/wms/unplannedissue-request-main/page` | getUnplannedissueRequestMainPage | **编号列表** |
| POST | `/wms/unplannedissue-request-main/senior` | getUnplannedissueRequestMainSenior | **计划外出库申请主** |
| GET | `/wms/unplannedissue-request-main/export-excel` | exportUnplannedissueRequestMainExcel | **计划外出库申请主** |
| POST | `/wms/unplannedissue-request-main/export-excel-senior` | exportUnplannedreceiptRequestMainSeniorExcel | **计划外出库申请主** |
| GET | `/wms/unplannedissue-request-main/get-import-template` | importTemplate | **计划外出库申请主** |
| POST | `/wms/unplannedissue-request-main/import` | importExcel | **计划外出库申请主** |
| GET | `/wms/unplannedissue-request-main/getUnplannedissueRequestById` | getUnplannedissueRequestById | **Excel 文件** |
| PUT | `/wms/unplannedissue-request-main/close` | closeUnplannedissueRequestMain | **编号** |
| PUT | `/wms/unplannedissue-request-main/reAdd` | reAddUnplannedissueRequestMain | **编号** |
| PUT | `/wms/unplannedissue-request-main/submit` | submitUnplannedissueRequestMain | **编号** |
| PUT | `/wms/unplannedissue-request-main/refused` | abortUnplannedissueRequestMain | **编号** |
| PUT | `/wms/unplannedissue-request-main/agree` | agreeUnplannedissueRequestMain | **编号** |
| PUT | `/wms/unplannedissue-request-main/handle` | handleUnplannedissueRequestMain | **编号** |
| GET | `/wms/unplannedissue-request-main/get-import-template-spare` | importTemplateSpare | **编号** |
| POST | `/wms/unplannedissue-request-main/importSpare` | importExcelSpare | **计划外出库申请主** |

### 非计划收货任务 (unplannedreceiptJob)  (23个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/unplannedreceipt-job-detail/create` | createUnplannedreceiptJobDetail | **计划外入库任务子** |
| PUT | `/wms/unplannedreceipt-job-detail/update` | updateUnplannedreceiptJobDetail | **计划外入库任务子** |
| DELETE | `/wms/unplannedreceipt-job-detail/delete` | deleteUnplannedreceiptJobDetail | **计划外入库任务子** |
| GET | `/wms/unplannedreceipt-job-detail/get` | getUnplannedreceiptJobDetail | **编号** |
| GET | `/wms/unplannedreceipt-job-detail/list` | getUnplannedreceiptJobDetailList | **编号** |
| GET | `/wms/unplannedreceipt-job-detail/page` | getUnplannedreceiptJobDetailPage | **编号列表** |
| POST | `/wms/unplannedreceipt-job-detail/senior` | getUnplannedreceiptJobDetailSenior | **计划外入库任务子** |
| GET | `/wms/unplannedreceipt-job-detail/export-excel` | exportUnplannedreceiptJobDetailExcel | **计划外入库任务子** |
| POST | `/wms/unplannedreceipt-job-main/create` | createUnplannedreceiptJobMain | **计划外入库任务主** |
| PUT | `/wms/unplannedreceipt-job-main/update` | updateUnplannedreceiptJobMain | **计划外入库任务主** |
| DELETE | `/wms/unplannedreceipt-job-main/delete` | deleteUnplannedreceiptJobMain | **计划外入库任务主** |
| GET | `/wms/unplannedreceipt-job-main/get` | getUnplannedreceiptJobMain | **编号** |
| GET | `/wms/unplannedreceipt-job-main/list` | getUnplannedreceiptJobMainList | **编号** |
| GET | `/wms/unplannedreceipt-job-main/page` | getUnplannedreceiptJobMainPage | **编号列表** |
| POST | `/wms/unplannedreceipt-job-main/senior` | getUnplannedreceiptJobMainSenior | **计划外入库任务主** |
| GET | `/wms/unplannedreceipt-job-main/export-excel` | exportUnplannedreceiptJobMainExcel | **计划外入库任务主** |
| POST | `/wms/unplannedreceipt-job-main/export-excel-senior` | exportUnplannedreceiptJobMainSeniorExcel | **计划外入库任务主** |
| GET | `/wms/unplannedreceipt-job-main/getUnplannedreceiptJobById` | getUnplannedreceiptJobById | **计划外入库任务主** |
| POST | `/wms/unplannedreceipt-job-main/getCountByStatus` | getCountByStatus | **编号** |
| PUT | `/wms/unplannedreceipt-job-main/accept` | acceptUnplannedreceiptJobMain | **类型数组** |
| PUT | `/wms/unplannedreceipt-job-main/abandon` | abandonUnplannedreceiptJobMain | **计划外入库任务主** |
| PUT | `/wms/unplannedreceipt-job-main/close` | closeUnplannedreceiptJobMain | **计划外入库任务主** |
| PUT | `/wms/unplannedreceipt-job-main/execute` | executeUnplannedreceiptJobMain | **计划外入库任务主** |

### 非计划收货记录 (unplannedreceiptRecord)  (19个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/unplannedreceipt-record-detail/create` | createUnplannedreceiptRecordDetail | **计划外入库记录子** |
| PUT | `/wms/unplannedreceipt-record-detail/update` | updateUnplannedreceiptRecordDetail | **计划外入库记录子** |
| DELETE | `/wms/unplannedreceipt-record-detail/delete` | deleteUnplannedreceiptRecordDetail | **计划外入库记录子** |
| GET | `/wms/unplannedreceipt-record-detail/get` | getUnplannedreceiptRecordDetail | **编号** |
| GET | `/wms/unplannedreceipt-record-detail/list` | getUnplannedreceiptRecordDetailList | **编号** |
| GET | `/wms/unplannedreceipt-record-detail/page` | getUnplannedreceiptRecordDetailPage | **编号列表** |
| POST | `/wms/unplannedreceipt-record-detail/senior` | getUnplannedreceiptRecordDetailSenior | **计划外入库记录子** |
| GET | `/wms/unplannedreceipt-record-detail/export-excel` | exportUnplannedreceiptRecordDetailExcel | **计划外入库记录子** |
| POST | `/wms/unplannedreceipt-record-main/create` | createUnplannedreceiptRecordMain | **计划外入库记录主** |
| PUT | `/wms/unplannedreceipt-record-main/update` | updateUnplannedreceiptRecordMain | **计划外入库记录主** |
| DELETE | `/wms/unplannedreceipt-record-main/delete` | deleteUnplannedreceiptRecordMain | **计划外入库记录主** |
| POST | `/wms/unplannedreceipt-record-main/senior` | getUnplannedreceiptRecordMainSenior | **编号** |
| GET | `/wms/unplannedreceipt-record-main/get` | getUnplannedreceiptRecordMain | **计划外入库记录主** |
| GET | `/wms/unplannedreceipt-record-main/list` | getUnplannedreceiptRecordMainList | **编号** |
| GET | `/wms/unplannedreceipt-record-main/page` | getUnplannedreceiptRecordMainPage | **编号列表** |
| GET | `/wms/unplannedreceipt-record-main/export-excel` | exportUnplannedreceiptRecordMainExcel | **计划外入库记录主** |
| POST | `/wms/unplannedreceipt-record-main/export-excel-senior` | exportUnplannedreceiptRecordMainSeniorExcel | **计划外入库记录主** |
| GET | `/wms/unplannedreceipt-record-main/getDetailInfoById` | getDetailInfoById | **计划外入库记录主** |
| GET | `/wms/unplannedreceipt-record-main/revoke` | revoke | **计划外入库记录主** |

### 非计划收货申请 (unplannedreceiptRequest)  (29个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/unplannedreceipt-request-detail/create` | createUnplannedreceiptRequestDetail | **计划外入库申请子** |
| PUT | `/wms/unplannedreceipt-request-detail/update` | updateUnplannedreceiptRequestDetail | **计划外入库申请子** |
| DELETE | `/wms/unplannedreceipt-request-detail/delete` | deleteUnplannedreceiptRequestDetail | **计划外入库申请子** |
| GET | `/wms/unplannedreceipt-request-detail/get` | getUnplannedreceiptRequestDetail | **编号** |
| GET | `/wms/unplannedreceipt-request-detail/list` | getUnplannedreceiptRequestDetailList | **计划外入库申请子** |
| GET | `/wms/unplannedreceipt-request-detail/page` | getUnplannedreceiptRequestDetailPage | **编号列表** |
| POST | `/wms/unplannedreceipt-request-detail/senior` | getUnplannedreceiptRequestDetailSenior | **计划外入库申请子** |
| GET | `/wms/unplannedreceipt-request-detail/export-excel` | exportUnplannedreceiptRequestDetailExcel | **计划外入库申请子** |
| PUT | `/wms/unplannedreceipt-request-detail/updateDetailPackingNumber` | updateUnplannedreceiptPackingNumber | **计划外入库申请子** |
| GET | `/wms/unplannedreceipt-request-detail/pageCreateLabel` | getUnplannedreceiptRequestDetailPageCreateLabel | **计划外入库申请子** |
| POST | `/wms/unplannedreceipt-request-main/create` | createUnplannedreceiptRequestMain | **计划外入库申请主** |
| PUT | `/wms/unplannedreceipt-request-main/update` | updateUnplannedreceiptRequestMain | **计划外入库申请主** |
| DELETE | `/wms/unplannedreceipt-request-main/delete` | deleteUnplannedreceiptRequestMain | **计划外入库申请主** |
| GET | `/wms/unplannedreceipt-request-main/get` | getUnplannedreceiptRequestMain | **编号** |
| GET | `/wms/unplannedreceipt-request-main/list` | getUnplannedreceiptRequestMainList | **编号** |
| GET | `/wms/unplannedreceipt-request-main/page` | getUnplannedreceiptRequestMainPage | **编号列表** |
| POST | `/wms/unplannedreceipt-request-main/senior` | getUnplannedreceiptRequestMainSenior | **计划外入库申请主** |
| GET | `/wms/unplannedreceipt-request-main/export-excel` | exportUnplannedreceiptRequestMainExcel | **计划外入库申请主** |
| POST | `/wms/unplannedreceipt-request-main/export-excel-senior` | exportUnplannedreceiptRequestMainSeniorExcel | **计划外入库申请主** |
| GET | `/wms/unplannedreceipt-request-main/get-import-template` | importTemplate | **计划外入库申请主** |
| POST | `/wms/unplannedreceipt-request-main/import` | importExcel | **计划外入库申请主** |
| GET | `/wms/unplannedreceipt-request-main/getUnplannedreceiptRequestById` | getUnplannedreceiptRequestById | **Excel 文件** |
| PUT | `/wms/unplannedreceipt-request-main/close` | closeUnplannedreceiptRequestMain | **编号** |
| PUT | `/wms/unplannedreceipt-request-main/reAdd` | reAddUnplannedreceiptRequestMain | **编号** |
| PUT | `/wms/unplannedreceipt-request-main/submit` | submitUnplannedreceiptRequestMain | **编号** |
| PUT | `/wms/unplannedreceipt-request-main/refused` | abortUnplannedreceiptRequestMain | **编号** |
| PUT | `/wms/unplannedreceipt-request-main/agree` | agreeUnplannedreceiptRequestMain | **编号** |
| PUT | `/wms/unplannedreceipt-request-main/handle` | handleUnplannedreceiptRequestMain | **编号** |
| PUT | `/wms/unplannedreceipt-request-main/handleBack` | handleUnplannedreceiptBackRequestMain | **编号** |

### 仓库 (warehouse)  (15个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/warehouse/create` | createWarehouse | **仓库** |
| PUT | `/wms/warehouse/update` | updateWarehouse | **仓库** |
| DELETE | `/wms/warehouse/delete` | deleteWarehouse | **仓库** |
| GET | `/wms/warehouse/get` | getWarehouse | **编号** |
| GET | `/wms/warehouse/list` | getWarehouseList | **编号** |
| GET | `/wms/warehouse/page` | getWarehousePage | **仓库** |
| POST | `/wms/warehouse/senior` | getWarehouseSenior | **仓库** |
| GET | `/wms/warehouse/export-excel` | exportWarehouseExcel | **仓库** |
| POST | `/wms/warehouse/export-excel-senior` | exportWarehouseExcel | **仓库** |
| GET | `/wms/warehouse/get-import-template` | importTemplate | **仓库** |
| POST | `/wms/warehouse/import` | importExcel | **仓库** |
| GET | `/wms/warehouse/pageBusinessTypeToWarehouse` | selectBusinessTypeToWarehouse | **Excel 文件** |
| POST | `/wms/warehouse/pageBusinessTypeToWarehouseSenior` | selectBusinessTypeToWarehouseSenior | **业务类型** |
| GET | `/wms/warehouse/listByCodes` | getProcessByCodes | **仓库** |
| GET | `/wms/warehouse/ListByCode` | ListByCode | **仓库** |

### 作业 (work)  (18个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/work-detail/create` | createWorkDetail | **生产订单子** |
| PUT | `/wms/work-detail/update` | updateWorkDetail | **生产订单子** |
| DELETE | `/wms/work-detail/delete` | deleteWorkDetail | **生产订单子** |
| GET | `/wms/work-detail/get` | getWorkDetail | **编号** |
| GET | `/wms/work-detail/list` | getWorkDetailList | **编号** |
| GET | `/wms/work-detail/page` | getWorkDetailPage | **编号列表** |
| POST | `/wms/work-detail/senior` | getWorkdetailSenior | **生产订单子** |
| POST | `/wms/work-main/create` | createWorkMain | **生产订单主** |
| PUT | `/wms/work-main/update` | updateWorkMain | **生产订单主** |
| DELETE | `/wms/work-main/delete` | deleteWorkMain | **生产订单主** |
| GET | `/wms/work-main/get` | getWorkMain | **编号** |
| GET | `/wms/work-main/list` | getWorkMainList | **编号** |
| GET | `/wms/work-main/page` | getWorkMainPage | **编号列表** |
| POST | `/wms/work-main/senior` | getWorkMainSenior | **生产订单主** |
| GET | `/wms/work-main/export-excel` | exportWorkMainExcel | **生产订单主** |
| POST | `/wms/work-main/export-excel-senior` | exportWorkMainSeniorExcel | **生产订单主** |
| GET | `/wms/work-main/get-import-template` | importTemplate | **生产订单主** |
| POST | `/wms/work-main/import` | importExcel | **生产订单主** |

### 车间 (workshop)  (11个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/workshop/create` | createWorkshop | **车间** |
| PUT | `/wms/workshop/update` | updateWorkshop | **车间** |
| DELETE | `/wms/workshop/delete` | deleteWorkshop | **车间** |
| GET | `/wms/workshop/get` | getWorkshop | **编号** |
| GET | `/wms/workshop/page` | getWorkshopPage | **编号** |
| POST | `/wms/workshop/senior` | getWorkshopSenior | **车间** |
| GET | `/wms/workshop/export-excel` | exportWorkshopExcel | **车间** |
| POST | `/wms/workshop/export-excel-senior` | exportWorkshopExcel | **车间** |
| GET | `/wms/workshop/get-import-template` | importTemplate | **车间** |
| POST | `/wms/workshop/import` | importExcel | **车间** |
| GET | `/wms/workshop/noPage` | getWorkshopNoPage | **Excel 文件** |

### 工位 (workstation)  (14个API)

| HTTP方法 | 端点路径 | 方法名 | 说明 |
|---------|---------|-------|------|
| POST | `/wms/workstation/create` | createWorkstation | **工位** |
| PUT | `/wms/workstation/update` | updateWorkstation | **工位** |
| DELETE | `/wms/workstation/delete` | deleteWorkstation | **工位** |
| GET | `/wms/workstation/get` | getWorkstation | **编号** |
| GET | `/wms/workstation/page` | getWorkstationPage | **编号** |
| POST | `/wms/workstation/senior` | getWorkstationSenior | **工位** |
| GET | `/wms/workstation/export-excel` | exportWorkstationExcel | **工位** |
| POST | `/wms/workstation/export-excel-senior` | exportWorkstationExcel | **工位** |
| GET | `/wms/workstation/get-import-template` | importTemplate | **工位** |
| POST | `/wms/workstation/import` | importExcel | **工位** |
| GET | `/wms/workstation/pageAreaToLocation` | selectAreaTypeToLocation | **Excel 文件** |
| POST | `/wms/workstation/pageAreaToLocationSenior` | selectAreaTypeToLocationSenior | **库区类型** |
| GET | `/wms/workstation/noPage` | getWorkstationNoPage | **工位** |
| GET | `/wms/workstation/listByCodes` | getWorkstationByCodes | **工位** |

