# EAM设备管理模块设计文档

## 1. 模块概述

### 1.1 模块职责
EAM（Enterprise Asset Management）设备管理模块是SFMS3.0系统中的核心业务模块，负责企业设备的全生命周期管理，包括设备台账管理、设备巡检、设备保养、设备维修、设备停机、设备转移、备件管理等核心功能。

### 1.2 业务价值
- 实现设备信息的集中管理，建立完整的设备档案
- 规范设备巡检、保养、维修流程，提高设备可用率
- 追踪设备运行状态，降低停机时间和故障率
- 管理备件库存，优化备件使用效率

## 2. 技术架构

### 2.1 模块结构
```
win-module-eam/
├── win-module-eam-api/          # API接口层
│   └── src/main/java/com/win/module/eam/
│       ├── api/                 # 对外暴露的API接口
│       └── enums/               # 枚举常量定义
└── win-module-eam-biz/          # 业务实现层
    └── src/main/java/com/win/module/eam/
        ├── controller/          # 控制器层
        ├── service/             # 服务层
        ├── dal/                 # 数据访问层
        │   ├── dataobject/      # 数据对象(DO)
        │   ├── mysql/           # MyBatis Mapper接口
        │   └── redis/           # Redis操作(如有)
        ├── convert/             # 对象转换
        ├── enums/               # 业务枚举
        └── util/                # 工具类
```

### 2.2 技术栈
- **框架**: Spring Boot + MyBatis Plus
- **数据库**: MySQL（表前缀: `basic_`, `eam_`）
- **API风格**: RESTful API，使用Swagger/OpenAPI 3.0文档化
- **权限控制**: `@PreAuthorize("@ss.hasPermission(...)")` 注解方式
- **日志记录**: `@OperateLog` 注解记录操作日志
- **数据校验**: `@Valid` + Bean Validation

## 3. 核心领域模型

### 3.1 设备台账 (EquipmentAccounts)
设备台账是EAM模块的核心实体，记录设备的基础信息。

| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 主键 |
| code | String | 设备编号 |
| name | String | 设备名称 |
| specification | String | 设备型号 |
| factoryCode | String | 出厂编号 |
| type | String | 设备类型(枚举) |
| power | String | 功率(kw) |
| status | String | 设备状态 |
| supplierCode | String | 供应商编号 |
| manufactureCode | String | 生产厂商编码 |
| purchasePrice | BigDecimal | 采购价格 |
| totalRunningTime | BigDecimal | 总运行时长 |
| outageRate | String | 设备停机率 |
| breakdownRecover | BigDecimal | 平均故障恢复时间 |
| factoryAreaCode | String | 所属厂区编号 |
| workshopCode | String | 车间编号 |
| lineCode | String | 产线编号 |
| processCode | String | 工序编号 |
| workstationCode | String | 工位编号 |

### 3.2 设备检验记录 (EquipmentInspectionRecord)
记录设备巡检的结果，包括主表和明细表。

- **主表(EquipmentInspectionRecordMain)**: 巡检任务基本信息
- **明细表(EquipmentInspectionRecordDetail)**: 巡检项的具体检查结果

### 3.3 设备保养记录 (EquipmentMaintenanceRecord)
记录设备保养信息。

- **主表(EquipmentMaintenanceRecordMain)**: 保养任务信息
- **明细表(EquipmentMaintenanceRecordDetail)**: 保养项目明细

### 3.4 设备维修记录 (EquipmentRepairRecord)
记录设备故障维修信息。

- **主表(EquipmentRepairRecordMain)**: 维修工单信息
- **明细表(EquipmentRepairRecordDetail)**: 维修明细

### 3.5 设备停机记录 (EquipmentShutdown)
记录设备非计划停机事件。

### 3.6 设备转移记录 (EquipmentTransferRecord)
记录设备在不同产线/车间之间的转移。

### 3.7 设备点检记录 (EquipmentSpotCheckRecord)
记录设备日常点检结果。

### 3.8 备件管理 (EquipmentToolSparePart)
管理设备备件的工具和库存。

## 4. 核心功能分析

### 4.1 设备台账管理
```
功能列表:
- 创建设备台账
- 更新设备台账
- 删除设备台账
- 查询设备台账(分页/不分页)
- 高级搜索
- Excel导入/导出
- 启用/停用设备
```

**关键接口**:
```java
// EquipmentAccountsController
POST   /eam/device/equipment-accounts/create    # 创建设备
PUT    /eam/device/equipment-accounts/update    # 更新设备
DELETE /eam/device/equipment-accounts/delete    # 删除设备
GET    /eam/device/equipment-accounts/get       # 获取单个设备
GET    /eam/device/equipment-accounts/page      # 分页查询
POST   /eam/device/equipment-accounts/senior     # 高级搜索
GET    /eam/device/equipment-accounts/export-excel # 导出Excel
POST   /eam/device/equipment-accounts/import     # 导入Excel
POST   /eam/device/equipment-accounts/ables      # 启用/停用
```

### 4.2 设备巡检管理
```
功能列表:
- 创建巡检任务
- 执行巡检记录
- 查询巡检历史
```

### 4.3 设备保养管理
```
功能列表:
- 创建保养计划
- 执行保养记录
- 查询保养历史
```

### 4.4 设备维修管理
```
功能列表:
- 报修登记
- 维修派工
- 维修执行
- 维修完成确认
```

### 4.5 设备点检管理
```
功能列表:
- 创建点检任务
- 执行点检
- 查询点检结果
```

## 5. 业务枚举

### 5.1 设备状态枚举 (DeviceStatusEnum)
- NORMAL: 正常
- FAULT: 故障
- MAINTENANCE: 保养中
- IDLE: 闲置
- SCRAPPED: 报废

### 5.2 设备类型枚举 (DeviceTypeEnum)
- 生产设备
- 检测设备
- 辅助设备

### 5.3 维修级别枚举 (RepairLevelEnum)
- 一级维修(紧急)
- 二级维修(重要)
- 三级维修(一般)

### 5.4 保养级别枚举 (MaintainLevelEnum)
- 一级保养
- 二级保养
- 三级保养

## 6. 数据字典

模块使用系统级数据字典，主要包括：

| 字典类型 | 字典键 | 说明 |
|----------|--------|------|
| DEVICE_STATUS | 设备状态 | 正常/故障/保养中/闲置/报废 |
| DEVICE_CLASS | 设备分类 | 生产设备/检测设备/辅助设备 |
| TRUE_FALSE | 是否 | 是/否 |
| FAULT_TYPE | 故障类型 | 按故障原因分类 |
| MAINTENANCE_TYPE | 保养类型 | 按保养级别分类 |

## 7. 权限设计

EAM模块采用基于注解的权限控制：

| 权限标识 | 说明 |
|----------|------|
| eam:equipment-accounts:create | 创建设备台账 |
| eam:equipment-accounts:update | 更新设备台账 |
| eam:equipment-accounts:delete | 删除设备台账 |
| eam:equipment-accounts:query | 查询设备台账 |
| eam:equipment-accounts:export | 导出设备台账 |
| eam:equipment-accounts:import | 导入设备台账 |

## 8. 数据流分析

### 8.1 设备创建流程
```
1. 用户提交设备信息(Controller)
2. 数据校验和转换(Service)
3. 保存到设备台账表(Mapper)
4. 返回新设备ID
```

### 8.2 巡检执行流程
```
1. 创建巡检任务 -> 设备巡检主表
2. 执行巡检项 -> 设备巡解明细表
3. 记录异常情况 -> 故障记录表
4. 更新设备状态 -> 设备台账表
```

### 8.3 维修处理流程
```
1. 报修登记 -> 维修请求表
2. 工单派发 -> 维修工单主表
3. 维修执行 -> 维修工单明细表
4. 完工确认 -> 更新设备状态
```

## 9. 关键技术实现

### 9.1 对象转换
使用MapStruct进行DO/VO/DTO之间的转换：
```java
// EquipmentAccountsConvert
INSTANCE.convert(EquipmentAccountsDO);      // DO -> RespVO
INSTANCE.convertPage(PageResult<DO>);        // 分页转换
INSTANCE.convertList(List<DO>);              // 列表转换
```

### 9.2 Excel导入导出
使用框架级Excel工具，支持：
- 模板下载
- 数据导入(支持更新/追加/覆盖模式)
- 错误数据回写
- 下拉选项数据验证

### 9.3 高级搜索
支持动态条件组合查询：
```java
PageResult<EquipmentAccountsDO> getEquipmentAccountsSenior(CustomConditions conditions);
```

## 10. 外部依赖

### 10.1 依赖模块
- `win-module-system`: 依赖系统模块获取用户信息、字典数据
- `win-module-infra`: 依赖基础设施模块获取文件存储服务

### 10.2 API暴露
EAM模块向其他模块提供以下API：
- 设备信息查询
- 设备状态更新
- 备件库存查询

## 11. 表结构汇总

| 表名 | 说明 |
|------|------|
| basic_equipment_accounts | 设备台账表 |
| basic_equipment_inspection_record_main | 设备巡检记录主表 |
| basic_equipment_inspection_record_detail | 设备巡检记录明细表 |
| basic_equipment_maintenance_main | 设备保养记录主表 |
| basic_equipment_maintenance_detail | 设备保养记录明细表 |
| basic_equipment_repair_record_main | 设备维修记录主表 |
| basic_equipment_repair_record_detail | 设备维修记录明细表 |
| basic_equipment_shutdown | 设备停机记录表 |
| basic_equipment_transfer_record | 设备转移记录表 |
| basic_equipment_spot_check_record_main | 设备点检记录主表 |
| basic_equipment_spot_check_record_detail | 设备点检记录明细表 |
| basic_equipment_tool_spare_part | 设备备件表 |
| basic_equipment_manufacturer | 设备厂商表 |
| basic_equipment_supplier | 设备供应商表 |
| basic_equipment_main_part | 设备主要部件表 |
| basic_fault_type | 故障类型表 |
| basic_fault_cause | 故障原因表 |

## 12. 完整API清单

### 12.1 基础数据

#### 12.1.1 EAM车间管理 (BasicEamWorkshop)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic-eam-workshop/create | 创建EAM车间 |
| PUT | /eam/basic-eam-workshop/update | 更新EAM车间 |
| DELETE | /eam/basic-eam-workshop/delete | 删除EAM车间 |
| GET | /eam/basic-eam-workshop/get | 获得EAM车间 |
| GET | /eam/basic-eam-workshop/page | 获得EAM车间分页 |
| POST | /eam/basic-eam-workshop/senior | 高级搜索获得EAM车间分页 |
| GET | /eam/basic-eam-workshop/export-excel | 导出EAM车间 Excel |
| POST | /eam/basic-eam-workshop/export-excel-senior | 高级搜索导出EAM车间 Excel |
| GET | /eam/basic-eam-workshop/get-import-template | 获得导入EAM车间模板 |
| POST | /eam/basic-eam-workshop/import | 导入EAM车间基本信息 |
| POST | /eam/basic-eam-workshop/ables | 启用/禁用 |
| GET | /eam/basic-eam-workshop/noPage | 获得EAM车间不分页 |

#### 12.1.2 生产线管理 (BasicEamProductionline)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic-eam-productionline/create | 创建生产线 |
| PUT | /eam/basic-eam-productionline/update | 更新生产线 |
| DELETE | /eam/basic-eam-productionline/delete | 删除生产线 |
| GET | /eam/basic-eam-productionline/get | 获得生产线 |
| GET | /eam/basic-eam-productionline/page | 获得生产线分页 |
| POST | /eam/basic-eam-productionline/senior | 高级搜索获得生产线分页 |
| GET | /eam/basic-eam-productionline/export-excel | 导出生产线 Excel |
| POST | /eam/basic-eam-productionline/export-excel-senior | 高级搜索导出生产线 Excel |
| GET | /eam/basic-eam-productionline/get-import-template | 获得导入生产线模板 |
| POST | /eam/basic-eam-productionline/import | 导入生产线基本信息 |

#### 12.1.3 故障类型管理 (BasicFaultType)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic-fault-type/create | 创建故障类型 |
| PUT | /eam/basic-fault-type/update | 更新故障类型 |
| DELETE | /eam/basic-fault-type/delete | 删除故障类型 |
| GET | /eam/basic-fault-type/get | 获得故障类型 |
| GET | /eam/basic-fault-type/page | 获得故障类型分页 |
| POST | /eam/basic-fault-type/senior | 高级搜索获得故障类型分页 |
| GET | /eam/basic-fault-type/export-excel | 导出故障类型 Excel |
| POST | /eam/basic-fault-type/export-excel-senior | 高级搜索导出故障类型 Excel |
| GET | /eam/basic-fault-type/get-import-template | 获得导入故障类型模板 |
| POST | /eam/basic-fault-type/import | 导入故障类型基本信息 |
| POST | /eam/basic-fault-type/ables | 启用/禁用 |
| GET | /eam/basic-fault-type/noPage | 获得故障类型不分页 |

#### 12.1.4 故障原因管理 (BasicFaultCause)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic-fault-cause/create | 创建故障原因 |
| PUT | /eam/basic-fault-cause/update | 更新故障原因 |
| DELETE | /eam/basic-fault-cause/delete | 删除故障原因 |
| GET | /eam/basic-fault-cause/get | 获得故障原因 |
| GET | /eam/basic-fault-cause/page | 获得故障原因分页 |
| POST | /eam/basic-fault-cause/senior | 高级搜索获得故障原因分页 |
| GET | /eam/basic-fault-cause/export-excel | 导出故障原因 Excel |
| POST | /eam/basic-fault-cause/export-excel-senior | 高级搜索导出故障原因 Excel |
| GET | /eam/basic-fault-cause/get-import-template | 获得导入故障原因模板 |
| POST | /eam/basic-fault-cause/import | 导入故障原因基本信息 |
| POST | /eam/basic-fault-cause/ables | 启用/禁用 |

#### 12.1.5 厂区班组角色维护 (ClassTypeRole)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic/class-type-role/create | 创建厂区班组角色维护 |
| PUT | /eam/basic/class-type-role/update | 更新厂区班组角色维护 |
| DELETE | /eam/basic/class-type-role/delete | 删除厂区班组角色维护 |
| GET | /eam/basic/class-type-role/get | 获得厂区班组角色维护 |
| GET | /eam/basic/class-type-role/page | 获得厂区班组角色维护分页 |
| POST | /eam/basic/class-type-role/senior | 高级搜索获得厂区班组角色维护分页 |
| GET | /eam/basic/class-type-role/export-excel | 导出厂区班组角色维护 Excel |
| POST | /eam/basic/class-type-role/export-excel-senior | 高级搜索导出厂区班组角色维护 Excel |
| GET | /eam/basic/class-type-role/get-import-template | 获得导入厂区班组角色维护模板 |
| POST | /eam/basic/class-type-role/import | 导入厂区班组角色维护基本信息 |
| POST | /eam/basic/class-type-role/ables | 启用/禁用 |

#### 12.1.6 文档类型管理 (DocumentType)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic/document-type/create | 创建文档类型 |
| PUT | /eam/basic/document-type/update | 更新文档类型 |
| DELETE | /eam/basic/document-type/delete | 删除文档类型 |
| GET | /eam/basic/document-type/get | 获得文档类型 |
| GET | /eam/basic/document-type/page | 获得文档类型分页 |
| POST | /eam/basic/document-type/senior | 高级搜索获得文档类型分页 |
| GET | /eam/basic/document-type/export-excel | 导出文档类型 Excel |
| POST | /eam/basic/document-type/export-excel-senior | 高级搜索导出文档类型 Excel |
| GET | /eam/basic/document-type/get-import-template | 获得导入文档类型模板 |
| POST | /eam/basic/document-type/import | 导入文档类型基本信息 |
| POST | /eam/basic/document-type/ables | 启用/禁用 |

#### 12.1.7 设备厂商管理 (EquipmentManufacturer)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic/equipment-manufacturer/create | 创建设备厂商 |
| PUT | /eam/basic/equipment-manufacturer/update | 更新设备厂商 |
| DELETE | /eam/basic/equipment-manufacturer/delete | 删除设备厂商 |
| GET | /eam/basic/equipment-manufacturer/get | 获得设备厂商 |
| GET | /eam/basic/equipment-manufacturer/page | 获得设备厂商分页 |
| POST | /eam/basic/equipment-manufacturer/senior | 高级搜索获得设备厂商分页 |
| GET | /eam/basic/equipment-manufacturer/export-excel | 导出设备厂商 Excel |
| POST | /eam/basic/equipment-manufacturer/export-excel-senior | 高级搜索导出设备厂商 Excel |
| GET | /eam/basic/equipment-manufacturer/get-import-template | 获得导入设备厂商模板 |
| POST | /eam/basic/equipment-manufacturer/import | 导入设备厂商基本信息 |
| POST | /eam/basic/equipment-manufacturer/ables | 启用/禁用 |
| GET | /eam/basic/equipment-manufacturer/noPage | 获得设备厂商不分页 |

#### 12.1.8 设备供应商管理 (EquipmentSupplier)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic/equipment-supplier/create | 创建供应商 |
| PUT | /eam/basic/equipment-supplier/update | 更新供应商 |
| DELETE | /eam/basic/equipment-supplier/delete | 删除供应商 |
| GET | /eam/basic/equipment-supplier/get | 获得供应商 |
| GET | /eam/basic/equipment-supplier/page | 获得供应商分页 |
| POST | /eam/basic/equipment-supplier/senior | 高级搜索获得供应商分页 |
| GET | /eam/basic/equipment-supplier/export-excel | 导出供应商 Excel |
| POST | /eam/basic/equipment-supplier/export-excel-senior | 高级搜索导出供应商 Excel |
| GET | /eam/basic/equipment-supplier/get-import-template | 获得导入供应商模板 |
| POST | /eam/basic/equipment-supplier/import | 导入供应商基本信息 |
| GET | /eam/basic/equipment-supplier/noPage | 获得供应商不分页 |

#### 12.1.9 设备主要部件管理 (EquipmentMainPart)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-main-part/create | 创建设备主要部件 |
| PUT | /eam/equipment-main-part/update | 更新设备主要部件 |
| DELETE | /eam/equipment-main-part/delete | 删除设备主要部件 |
| GET | /eam/equipment-main-part/get | 获得设备主要部件 |
| GET | /eam/equipment-main-part/page | 获得设备主要部件分页 |
| POST | /eam/equipment-main-part/senior | 高级搜索获得设备主要部件分页 |
| GET | /eam/equipment-main-part/export-excel | 导出设备主要部件 Excel |
| GET | /eam/equipment-main-part/get-import-template | 获得导入设备主要部件模板 |
| POST | /eam/equipment-main-part/import | 导入设备主要部件基本信息 |
| POST | /eam/equipment-main-part/ables | 启用/禁用 |

#### 12.1.10 巡检方案管理 (BasicInspectionOption)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic-inspection-option/create | 创建巡检方案 |
| PUT | /eam/basic-inspection-option/update | 更新巡检方案 |
| DELETE | /eam/basic-inspection-option/delete | 删除巡检方案 |
| GET | /eam/basic-inspection-option/get | 获得巡检方案 |
| GET | /eam/basic-inspection-option/page | 获得巡检方案分页 |
| POST | /eam/basic-inspection-option/senior | 高级搜索获得巡检方案分页 |
| GET | /eam/basic-inspection-option/export-excel | 导出巡检方案 Excel |
| GET | /eam/basic-inspection-option/get-import-template | 获得导入巡检方案模板 |
| POST | /eam/basic-inspection-option/ables | 启用/禁用 |

#### 12.1.11 保养方案管理 (BasicMaintenanceOption)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic-maintenance-option/create | 创建保养方案 |
| PUT | /eam/basic-maintenance-option/update | 更新保养方案 |
| DELETE | /eam/basic-maintenance-option/delete | 删除保养方案 |
| GET | /eam/basic-maintenance-option/get | 获得保养方案 |
| GET | /eam/basic-maintenance-option/page | 获得保养方案分页 |
| POST | /eam/basic-maintenance-option/senior | 高级搜索获得保养方案分页 |
| GET | /eam/basic-maintenance-option/export-excel | 导出保养方案 Excel |
| GET | /eam/basic-maintenance-option/get-import-template | 获得导入保养方案模板 |
| POST | /eam/basic-maintenance-option/ables | 启用/禁用 |

#### 12.1.12 点检方案管理 (BasicSpotCheckOption)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic-spot-check-option/create | 创建点检方案 |
| PUT | /eam/basic-spot-check-option/update | 更新点检方案 |
| DELETE | /eam/basic-spot-check-option/delete | 删除点检方案 |
| GET | /eam/basic-spot-check-option/get | 获得点检方案 |
| GET | /eam/basic-spot-check-option/page | 获得点检方案分页 |
| POST | /eam/basic-spot-check-option/senior | 高级搜索获得点检方案分页 |
| GET | /eam/basic-spot-check-option/export-excel | 导出点检方案 Excel |
| GET | /eam/basic-spot-check-option/get-import-template | 获得导入点检方案模板 |
| POST | /eam/basic-spot-check-option/import | 导入点检方案基本信息 |
| POST | /eam/basic-spot-check-option/ables | 启用/禁用 |

#### 12.1.13 巡检项管理 (InspectionItem)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic/inspection-item/create | 创建巡检项 |
| PUT | /eam/basic/inspection-item/update | 更新巡检项 |
| DELETE | /eam/basic/inspection-item/delete | 删除巡检项 |
| GET | /eam/basic/inspection-item/get | 获得巡检项 |
| GET | /eam/basic/inspection-item/getList | 获得巡检项列表 |
| GET | /eam/basic/inspection-item/page | 获得巡检项分页 |
| POST | /eam/basic/inspection-item/senior | 高级搜索获得巡检项分页 |
| GET | /eam/basic/inspection-item/export-excel | 导出巡检项 Excel |
| POST | /eam/basic/inspection-item/export-excel-senior | 高级搜索导出巡检项 Excel |
| GET | /eam/basic/inspection-item/get-import-template | 获得导入巡检项模板 |
| POST | /eam/basic/inspection-item/import | 导入巡检项基本信息 |
| POST | /eam/basic/inspection-item/ables | 启用/禁用 |

#### 12.1.14 巡检项选择集管理 (InspectionItemSelectSet / RelationInspectionItemSelectSet)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic/inspection-item-select-set/create | 创建巡检项选择集 |
| PUT | /eam/basic/inspection-item-select-set/update | 更新巡检项选择集 |
| DELETE | /eam/basic/inspection-item-select-set/delete | 删除巡检项选择集 |
| GET | /eam/basic/inspection-item-select-set/get | 获得巡检项选择集 |
| GET | /eam/basic/inspection-item-select-set/page | 获得巡检项选择集分页 |
| POST | /eam/basic/inspection-item-select-set/senior | 高级搜索获得巡检项选择集分页 |
| POST | /eam/basic/inspection-item-select-set/ables | 启用/禁用 |
| POST | /eam/basic/relation-inspection-item-select-set/create | 创建巡检项选择集与项关联 |
| PUT | /eam/basic/relation-inspection-item-select-set/update | 更新巡检项选择集与项关联 |
| DELETE | /eam/basic/relation-inspection-item-select-set/delete | 删除巡检项选择集与项关联 |
| GET | /eam/basic/relation-inspection-item-select-set/get | 获得巡检项选择集与项关联 |
| GET | /eam/basic/relation-inspection-item-select-set/page | 获得巡检项选择集与项关联分页 |
| POST | /eam/basic/relation-inspection-item-select-set/senior | 高级搜索获得巡检项选择集与项关联分页 |

#### 12.1.15 保养项管理 (MaintenanceItem)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic/maintenance-item/create | 创建保养项 |
| PUT | /eam/basic/maintenance-item/update | 更新保养项 |
| DELETE | /eam/basic/maintenance-item/delete | 删除保养项 |
| GET | /eam/basic/maintenance-item/get | 获得保养项 |
| GET | /eam/basic/maintenance-item/page | 获得保养项分页 |
| POST | /eam/basic/maintenance-item/senior | 高级搜索获得保养项分页 |
| GET | /eam/basic/maintenance-item/export-excel | 导出保养项 Excel |
| POST | /eam/basic/maintenance-item/export-excel-senior | 高级搜索导出保养项 Excel |
| GET | /eam/basic/maintenance-item/get-import-template | 获得导入保养项模板 |
| POST | /eam/basic/maintenance-item/import | 导入保养项基本信息 |
| POST | /eam/basic/maintenance-item/ables | 启用/禁用 |

#### 12.1.16 保养项选择集管理 (MaintenanceItemSelectSet / RelationMaintenanceItemSelectSet)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic/maintenance-item-select-set/create | 创建保养项选择集 |
| PUT | /eam/basic/maintenance-item-select-set/update | 更新保养项选择集 |
| DELETE | /eam/basic/maintenance-item-select-set/delete | 删除保养项选择集 |
| GET | /eam/basic/maintenance-item-select-set/get | 获得保养项选择集 |
| GET | /eam/basic/maintenance-item-select-set/page | 获得保养项选择集分页 |
| POST | /eam/basic/maintenance-item-select-set/senior | 高级搜索获得保养项选择集分页 |
| POST | /eam/basic/maintenance-item-select-set/ables | 启用/禁用 |
| POST | /eam/basic/relation-maintenance-item-select-set/create | 创建保养项选择集与项关联 |
| PUT | /eam/basic/relation-maintenance-item-select-set/update | 更新保养项选择集与项关联 |
| DELETE | /eam/basic/relation-maintenance-item-select-set/delete | 删除保养项选择集与项关联 |
| GET | /eam/basic/relation-maintenance-item-select-set/get | 获得保养项选择集与项关联 |
| GET | /eam/basic/relation-maintenance-item-select-set/page | 获得保养项选择集与项关联分页 |
| POST | /eam/basic/relation-maintenance-item-select-set/senior | 高级搜索获得保养项选择集与项关联分页 |

#### 12.1.17 点检项管理 (SpotCheckItem)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic/spot-check-item/create | 创建点检项 |
| PUT | /eam/basic/spot-check-item/update | 更新点检项 |
| DELETE | /eam/basic/spot-check-item/delete | 删除点检项 |
| GET | /eam/basic/spot-check-item/get | 获得点检项 |
| GET | /eam/basic/spot-check-item/getList | 获得点检项列表 |
| GET | /eam/basic/spot-check-item/page | 获得点检项分页 |
| POST | /eam/basic/spot-check-item/senior | 高级搜索获得点检项分页 |
| GET | /eam/basic/spot-check-item/export-excel | 导出点检项 Excel |
| POST | /eam/basic/spot-check-item/export-excel-senior | 高级搜索导出点检项 Excel |
| GET | /eam/basic/spot-check-item/get-import-template | 获得导入点检项模板 |
| POST | /eam/basic/spot-check-item/import | 导入点检项基本信息 |
| POST | /eam/basic/spot-check-item/ables | 启用/禁用 |

#### 12.1.18 点检选择集管理 (SpotCheckSelectSet / RelationSpotCheckItemSelectSet)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic-spot-check-select-set/create | 创建点检选择集维护 |
| PUT | /eam/basic-spot-check-select-set/update | 更新点检选择集维护 |
| DELETE | /eam/basic-spot-check-select-set/delete | 删除点检选择集维护 |
| GET | /eam/basic-spot-check-select-set/get | 获得点检选择集维护 |
| GET | /eam/basic-spot-check-select-set/page | 获得点检选择集维护分页 |
| POST | /eam/basic-spot-check-select-set/senior | 高级搜索获得点检选择集维护分页 |
| POST | /eam/basic-spot-check-select-set/ables | 启用/禁用 |
| POST | /eam/spot-check-select-set/create | 创建点检项选择集与项关联 |
| PUT | /eam/spot-check-select-set/update | 更新点检项选择集与项关联 |
| DELETE | /eam/spot-check-select-set/delete | 删除点检项选择集与项关联 |
| GET | /eam/spot-check-select-set/get | 获得点检项选择集与项关联 |
| GET | /eam/spot-check-select-set/page | 获得点检项选择集与项关联分页 |
| POST | /eam/spot-check-select-set/senior | 高级搜索获得点检项选择集与项关联分页 |

#### 12.1.19 文档类型选择集管理 (DocumentTypeSelectSet / RelationDocumentTypeSelectSet)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic/document-type-select-set/create | 创建文档类型选择集 |
| PUT | /eam/basic/document-type-select-set/update | 更新文档类型选择集 |
| DELETE | /eam/basic/document-type-select-set/delete | 删除文档类型选择集 |
| GET | /eam/basic/document-type-select-set/get | 获得文档类型选择集 |
| GET | /eam/basic/document-type-select-set/page | 获得文档类型选择集分页 |
| POST | /eam/basic/document-type-select-set/senior | 高级搜索获得文档类型选择集分页 |
| GET | /eam/basic/document-type-select-set/export-excel | 导出文档类型选择集 Excel |
| GET | /eam/basic/document-type-select-set/get-import-template | 获得导入文档类型选择集模板 |
| POST | /eam/basic/document-type-select-set/import | 导入文档类型选择集基本信息 |
| POST | /eam/basic/document-type-select-set/ables | 启用/禁用 |
| POST | /eam/basic/relation-document-type-select-set/create | 创建文档类型选择集与项关联 |
| PUT | /eam/basic/relation-document-type-select-set/update | 更新文档类型选择集与项关联 |
| DELETE | /eam/basic/relation-document-type-select-set/delete | 删除文档类型选择集与项关联 |
| GET | /eam/basic/relation-document-type-select-set/get | 获得文档类型选择集与项关联 |
| GET | /eam/basic/relation-document-type-select-set/page | 获得文档类型选择集与项关联分页 |
| POST | /eam/basic/relation-document-type-select-set/senior | 高级搜索获得文档类型选择集与项关联分页 |

#### 12.1.20 工装物料关联管理 (ToolMod)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/basic/tool-mod/create | 创建工装物料关联 |
| PUT | /eam/basic/tool-mod/update | 更新工装物料关联 |
| DELETE | /eam/basic/tool-mod/delete | 删除工装物料关联 |
| GET | /eam/basic/tool-mod/get | 获得工装物料关联 |
| GET | /eam/basic/tool-mod/page | 获得工装物料关联分页 |
| POST | /eam/basic/tool-mod/senior | 高级搜索获得工装物料关联分页 |
| GET | /eam/basic/tool-mod/export-excel | 导出工装物料关联 Excel |
| GET | /eam/basic/tool-mod/get-import-template | 获得导入工装物料关联模板 |
| POST | /eam/basic/tool-mod/import | 导入工装物料关联基本信息 |
| POST | /eam/basic/tool-mod/createBatch | 批量创建工装物料关联 |

#### 12.1.21 表数据扩展属性管理 (TableDataExtendedAttribute)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/table-data-extended-attribute/create | 创建表数据扩展属性字段 |
| PUT | /eam/table-data-extended-attribute/update | 更新表数据扩展属性字段 |
| DELETE | /eam/table-data-extended-attribute/delete | 删除表数据扩展属性字段 |
| GET | /eam/table-data-extended-attribute/get | 获得表数据扩展属性字段 |
| GET | /eam/table-data-extended-attribute/page | 获得表数据扩展属性字段分页 |
| POST | /eam/table-data-extended-attribute/senior | 高级搜索获得表数据扩展属性字段分页 |
| GET | /eam/table-data-extended-attribute/export-excel | 导出表数据扩展属性字段 Excel |
| GET | /eam/table-data-extended-attribute/get-import-template | 获得导入表数据扩展属性字段模板 |
| POST | /eam/table-data-extended-attribute/import | 导入表数据扩展属性字段基本信息 |
| POST | /eam/table-data-extended-attribute/createBatch | 批量创建表数据扩展属性字段 |
| GET | /eam/table-data-extended-attribute/noPage | 获得表数据扩展属性字段不分页 |

#### 12.1.22 设备台账附件文件管理 (AttachmentFile)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/attachment-file/create | 创建设备模具附件文件 |
| DELETE | /eam/attachment-file/delete | 删除设备模具附件文件 |
| GET | /eam/attachment-file/get | 获得设备模具附件文件 |
| POST | /eam/attachment-file/upload | 上传文件 |
| POST | /eam/attachment-file/listNoPage | 获得设备模具附件文件列表 |

### 12.2 设备台账管理 (EquipmentAccounts)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/device/equipment-accounts/create | 创建设备台账 |
| PUT | /eam/device/equipment-accounts/update | 更新设备台账 |
| DELETE | /eam/device/equipment-accounts/delete | 删除设备台账 |
| GET | /eam/device/equipment-accounts/get | 获得设备台账 |
| GET | /eam/device/equipment-accounts/detail | 获得设备台账详情 |
| GET | /eam/device/equipment-accounts/page | 获得设备台账分页 |
| POST | /eam/device/equipment-accounts/senior | 高级搜索获得设备台账分页 |
| GET | /eam/device/equipment-accounts/export-excel | 导出设备台账 Excel |
| POST | /eam/device/equipment-accounts/export-excel-senior | 高级搜索导出设备台账 Excel |
| GET | /eam/device/equipment-accounts/get-import-template | 获得导入设备台账模板 |
| POST | /eam/device/equipment-accounts/import | 导入设备台账基本信息 |
| GET | /eam/device/equipment-accounts/noPage | 获得设备台账不分页 |
| POST | /eam/device/equipment-accounts/ables | 启用/停用 |

### 12.3 库位管理

#### 12.3.1 库位管理 (LocationEAM)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/location/create | 创建库位 |
| PUT | /eam/location/update | 更新库位 |
| DELETE | /eam/location/delete | 删除库位 |
| GET | /eam/location/get | 获得库位 |
| GET | /eam/location/page | 获得库位分页 |
| POST | /eam/location/senior | 高级搜索获得库位分页 |
| GET | /eam/location/export-excel | 导出库位 Excel |
| GET | /eam/location/get-import-template | 获得导入库位模板 |
| POST | /eam/location/import | 导入库位基本信息 |
| GET | /eam/location/noPage | 获得库位不分页 |
| POST | /eam/location/ables | 启用/禁用 |

#### 12.3.2 库区管理 (LocationAreaEAM)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/location-area/create | 创建库区 |
| PUT | /eam/location-area/update | 更新库区 |
| DELETE | /eam/location-area/delete | 删除库区 |
| GET | /eam/location-area/get | 获得库区 |
| GET | /eam/location-area/page | 获得库区分页 |
| POST | /eam/location-area/senior | 高级搜索获得库区分页 |
| GET | /eam/location-area/export-excel | 导出库区 Excel |
| GET | /eam/location-area/get-import-template | 获得导入库区模板 |
| POST | /eam/location-area/import | 导入库区基本信息 |
| POST | /eam/location-area/ables | 启用/禁用 |

### 12.4 备件管理

#### 12.4.1 备件物料管理 (ItemEAM)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/item/create | 创建备件物料 |
| PUT | /eam/item/update | 更新备件物料 |
| DELETE | /eam/item/delete | 删除备件物料 |
| GET | /eam/item/get | 获得备件物料 |
| GET | /eam/item/page | 获得备件物料分页 |
| POST | /eam/item/senior | 高级搜索获得备件物料分页 |
| GET | /eam/item/export-excel | 导出备件物料 Excel |
| POST | /eam/item/export-excel-senior | 高级搜索导出备件物料 Excel |
| GET | /eam/item/get-import-template | 获得导入备件物料模板 |
| POST | /eam/item/import | 导入备件物料基本信息 |
| POST | /eam/item/ables | 启用/禁用 |
| GET | /eam/item/noPage | 获得备件物料不分页 |

#### 12.4.2 备件台账管理 (ItemAccounts)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/item-accounts/create | 创建备件台账 |
| PUT | /eam/item-accounts/update | 更新备件台账 |
| DELETE | /eam/item-accounts/delete | 删除备件台账 |
| GET | /eam/item-accounts/get | 获得备件台账 |
| GET | /eam/item-accounts/page | 获得备件台账分页 |
| POST | /eam/item-accounts/senior | 高级搜索获得备件台账分页 |
| GET | /eam/item-accounts/export-excel | 导出备件台账 Excel |
| POST | /eam/item-accounts/export-excel-senior | 高级搜索导出备件台账 Excel |
| GET | /eam/item-accounts/get-import-template | 获得导入备件台账模板 |
| POST | /eam/item-accounts/import | 导入备件台账基本信息 |
| GET | /eam/item-accounts/noPage | 获得备件台账不分页 |

#### 12.4.3 库存事务管理 (TransactionEAM)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/transaction/create | 创建库存事务 |
| PUT | /eam/transaction/update | 更新库存事务 |
| DELETE | /eam/transaction/delete | 删除库存事务 |
| GET | /eam/transaction/get | 获得库存事务 |
| GET | /eam/transaction/list | 获得库存事务列表 |
| GET | /eam/transaction/page | 获得库存事务分页 |
| POST | /eam/transaction/senior | 高级搜索获得库存事务分页 |
| GET | /eam/transaction/export-excel | 导出库存事务 Excel |
| GET | /eam/transaction/get-import-template | 获得导入库存事务模板 |

### 12.5 巡检管理

#### 12.5.1 巡检工单管理 (EquipmentInspectionMain)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-inspection-main/create | 创建巡检工单主 |
| PUT | /eam/equipment-inspection-main/update | 更新巡检工单主 |
| POST | /eam/equipment-inspection-main/updateOrder | 执行设备巡检工单状态 |
| POST | /eam/equipment-inspection-main/verifyOrder | 验证设备巡检工单 |
| POST | /eam/equipment-inspection-main/execute | 执行设备巡检工单 |
| POST | /eam/equipment-inspection-main/fallback | 打回设备巡检工单 |
| DELETE | /eam/equipment-inspection-main/delete | 删除巡检工单主 |
| GET | /eam/equipment-inspection-main/get | 获得巡检工单主 |
| GET | /eam/equipment-inspection-main/page | 获得巡检工单主分页 |
| POST | /eam/equipment-inspection-main/senior | 高级搜索获得巡检工单主分页 |
| GET | /eam/equipment-inspection-main/export-excel | 导出巡检工单主 Excel |
| GET | /eam/equipment-inspection-main/get-import-template | 获得导入巡检工单主模板 |

#### 12.5.2 巡检工单明细管理 (EquipmentInspectionDetail)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-inspection-detail/create | 创建巡检工单子 |
| PUT | /eam/equipment-inspection-detail/update | 更新巡检工单子 |
| DELETE | /eam/equipment-inspection-detail/delete | 删除巡检工单子 |
| GET | /eam/equipment-inspection-detail/get | 获得巡检工单子 |
| GET | /eam/equipment-inspection-detail/getList | 获得巡检工单子列表 |
| GET | /eam/equipment-inspection-detail/page | 获得巡检工单子分页 |
| POST | /eam/equipment-inspection-detail/senior | 高级搜索获得巡检工单子分页 |
| GET | /eam/equipment-inspection-detail/export-excel | 导出巡检工单子 Excel |
| GET | /eam/equipment-inspection-detail/get-import-template | 获得导入巡检工单子模板 |

#### 12.5.3 巡检记录管理 (EquipmentInspectionRecordMain)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-inspection-record-main/create | 创建巡检记录主 |
| PUT | /eam/equipment-inspection-record-main/update | 更新巡检记录主 |
| DELETE | /eam/equipment-inspection-record-main/delete | 删除巡检记录主 |
| GET | /eam/equipment-inspection-record-main/get | 获得巡检记录主 |
| GET | /eam/equipment-inspection-record-main/page | 获得巡检记录主分页 |
| POST | /eam/equipment-inspection-record-main/senior | 高级搜索获得巡检记录主分页 |
| GET | /eam/equipment-inspection-record-main/export-excel | 导出巡检记录主 Excel |
| GET | /eam/equipment-inspection-record-main/get-import-template | 获得导入巡检记录主模板 |
| POST | /eam/equipment-inspection-record-main/import | 导入巡检记录主基本信息 |

#### 12.5.4 巡检记录明细管理 (EquipmentInspectionRecordDetail)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-inspection-record-detail/create | 创建巡检记录子 |
| PUT | /eam/equipment-inspection-record-detail/update | 更新巡检记录子 |
| DELETE | /eam/equipment-inspection-record-detail/delete | 删除巡检记录子 |
| GET | /eam/equipment-inspection-record-detail/get | 获得巡检记录子 |
| GET | /eam/equipment-inspection-record-detail/page | 获得巡检记录子分页 |
| POST | /eam/equipment-inspection-record-detail/senior | 高级搜索获得巡检记录子分页 |
| GET | /eam/equipment-inspection-record-detail/export-excel | 导出巡检记录子 Excel |
| GET | /eam/equipment-inspection-record-detail/get-import-template | 获得导入巡检记录子模板 |
| POST | /eam/equipment-inspection-record-detail/import | 导入巡检记录子基本信息 |

#### 12.5.5 巡检计划管理 (PlanInspection)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/plan-inspection/create | 创建巡检计划 |
| PUT | /eam/plan-inspection/update | 更新巡检计划 |
| DELETE | /eam/plan-inspection/delete | 删除巡检计划 |
| GET | /eam/plan-inspection/get | 获得巡检计划 |
| GET | /eam/plan-inspection/page | 获得巡检计划分页 |
| POST | /eam/plan-inspection/senior | 高级搜索获得巡检计划分页 |
| GET | /eam/plan-inspection/export-excel | 导出巡检计划 Excel |
| GET | /eam/plan-inspection/get-import-template | 获得导入巡检计划模板 |
| POST | /eam/plan-inspection/import | 导入巡检计划基本信息 |
| POST | /eam/plan-inspection/ables | 启用/禁用 |

### 12.6 保养管理

#### 12.6.1 保养工单管理 (EquipmentMaintenanceMain)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-maintenance-main/create | 创建设备保养工单主 |
| PUT | /eam/equipment-maintenance-main/update | 更新设备保养工单主 |
| POST | /eam/equipment-maintenance-main/updateOrder | 执行设备保养工单状态 |
| POST | /eam/equipment-maintenance-main/verifyOrder | 验证设备保养工单状态 |
| POST | /eam/equipment-maintenance-main/execute | 执行设备保养工单 |
| POST | /eam/equipment-maintenance-main/fallback | 打回保养工单 |
| DELETE | /eam/equipment-maintenance-main/delete | 删除设备保养工单主 |
| GET | /eam/equipment-maintenance-main/get | 获得设备保养工单主 |
| GET | /eam/equipment-maintenance-main/get-order-config | 获得设备保养工单系统配置 |
| GET | /eam/equipment-maintenance-main/page | 获得设备保养工单主分页 |
| POST | /eam/equipment-maintenance-main/senior | 高级搜索获得设备保养工单主分页 |
| GET | /eam/equipment-maintenance-main/export-excel | 导出设备保养工单主 Excel |
| GET | /eam/equipment-maintenance-main/get-import-template | 获得导入设备保养工单主模板 |
| POST | /eam/equipment-maintenance-main/ables | 启用/禁用 |

#### 12.6.2 保养工单明细管理 (EquipmentMaintenanceDetail)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-maintenance-detail/create | 创建保养工单子 |
| PUT | /eam/equipment-maintenance-detail/update | 更新保养工单子 |
| DELETE | /eam/equipment-maintenance-detail/delete | 删除保养工单子 |
| GET | /eam/equipment-maintenance-detail/get | 获得保养工单子 |
| GET | /eam/equipment-maintenance-detail/getList | 获得保养工单子列表 |
| GET | /eam/equipment-maintenance-detail/page | 获得保养工单子分页 |
| POST | /eam/equipment-maintenance-detail/senior | 高级搜索获得保养工单子分页 |
| GET | /eam/equipment-maintenance-detail/export-excel | 导出保养工单子 Excel |
| GET | /eam/equipment-maintenance-detail/get-import-template | 获得导入保养工单子模板 |

#### 12.6.3 保养记录管理 (EquipmentMaintenanceRecordMain)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-maintenance-record-main/create | 创建设备保养记录主 |
| PUT | /eam/equipment-maintenance-record-main/update | 更新设备保养记录主 |
| DELETE | /eam/equipment-maintenance-record-main/delete | 删除设备保养记录主 |
| GET | /eam/equipment-maintenance-record-main/get | 获得设备保养记录主 |
| GET | /eam/equipment-maintenance-record-main/page | 获得设备保养记录主分页 |
| POST | /eam/equipment-maintenance-record-main/senior | 高级搜索获得设备保养记录主分页 |
| GET | /eam/equipment-maintenance-record-main/export-excel | 导出设备保养记录主 Excel |
| GET | /eam/equipment-maintenance-record-main/get-import-template | 获得导入设备保养记录主模板 |
| POST | /eam/equipment-maintenance-record-main/import | 导入设备保养记录主基本信息 |

#### 12.6.4 保养记录明细管理 (EquipmentMaintenanceRecordDetail)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-maintenance-record-detail/create | 创建保养记录子 |
| PUT | /eam/equipment-maintenance-record-detail/update | 更新保养记录子 |
| DELETE | /eam/equipment-maintenance-record-detail/delete | 删除保养记录子 |
| GET | /eam/equipment-maintenance-record-detail/get | 获得保养记录子 |
| GET | /eam/equipment-maintenance-record-detail/page | 获得保养记录子分页 |
| POST | /eam/equipment-maintenance-record-detail/senior | 高级搜索获得保养记录子分页 |
| GET | /eam/equipment-maintenance-record-detail/export-excel | 导出保养记录子 Excel |
| GET | /eam/equipment-maintenance-record-detail/get-import-template | 获得导入保养记录子模板 |
| POST | /eam/equipment-maintenance-record-detail/import | 导入保养记录子基本信息 |

#### 12.6.5 保养计划管理 (Maintenance)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/maintenance/create | 创建保养计划 |
| PUT | /eam/maintenance/update | 更新保养计划 |
| DELETE | /eam/maintenance/delete | 删除保养计划 |
| GET | /eam/maintenance/get | 获得保养计划 |
| GET | /eam/maintenance/page | 获得保养计划分页 |
| POST | /eam/maintenance/senior | 高级搜索获得保养计划分页 |
| GET | /eam/maintenance/export-excel | 导出保养计划 Excel |
| GET | /eam/maintenance/get-import-template | 获得导入保养计划模板 |
| POST | /eam/maintenance/import | 导入保养计划基本信息 |
| POST | /eam/maintenance/ables | 启用/禁用 |

### 12.7 点检管理

#### 12.7.1 点检工单管理 (EquipmentSpotCheckMain)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-spot-check-main/create | 创建点检工单主 |
| PUT | /eam/equipment-spot-check-main/update | 更新点检工单主 |
| POST | /eam/equipment-spot-check-main/updateOrder | 更新点检工单主状态 |
| POST | /eam/equipment-spot-check-main/verifyOrder | 验证点检工单主 |
| DELETE | /eam/equipment-spot-check-main/delete | 删除点检工单主 |
| POST | /eam/equipment-spot-check-main/execute | 执行设备点检工单 |
| POST | /eam/equipment-spot-check-main/fallback | 打回设备点检工单 |
| GET | /eam/equipment-spot-check-main/get | 获得点检工单主 |
| GET | /eam/equipment-spot-check-main/page | 获得点检工单主分页 |
| POST | /eam/equipment-spot-check-main/senior | 高级搜索获得点检工单主分页 |
| GET | /eam/equipment-spot-check-main/export-excel | 导出点检工单主 Excel |
| GET | /eam/equipment-spot-check-main/get-import-template | 获得导入点检工单主模板 |

#### 12.7.2 点检工单明细管理 (EquipmentSpotCheckDetail)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-spot-check-detail/create | 创建点检工单子 |
| PUT | /eam/equipment-spot-check-detail/update | 更新点检工单子 |
| DELETE | /eam/equipment-spot-check-detail/delete | 删除点检工单子 |
| GET | /eam/equipment-spot-check-detail/get | 获得点检工单子 |
| GET | /eam/equipment-spot-check-detail/getList | 获得点检工单子列表 |
| GET | /eam/equipment-spot-check-detail/page | 获得点检工单子分页 |
| POST | /eam/equipment-spot-check-detail/senior | 高级搜索获得点检工单子分页 |
| GET | /eam/equipment-spot-check-detail/export-excel | 导出点检工单子 Excel |
| GET | /eam/equipment-spot-check-detail/get-import-template | 获得导入点检工单子模板 |

#### 12.7.3 点检记录管理 (EquipmentSpotCheckRecordMain)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-spot-check-record-main/create | 创建点检记录主 |
| PUT | /eam/equipment-spot-check-record-main/update | 更新点检记录主 |
| DELETE | /eam/equipment-spot-check-record-main/delete | 删除点检记录主 |
| GET | /eam/equipment-spot-check-record-main/get | 获得点检记录主 |
| GET | /eam/equipment-spot-check-record-main/page | 获得点检记录主分页 |
| POST | /eam/equipment-spot-check-record-main/senior | 高级搜索获得点检记录主分页 |
| GET | /eam/equipment-spot-check-record-main/export-excel | 导出点检记录主 Excel |
| GET | /eam/equipment-spot-check-record-main/get-import-template | 获得导入点检记录主模板 |
| POST | /eam/equipment-spot-check-record-main/import | 导入点检记录主基本信息 |

#### 12.7.4 点检记录明细管理 (EquipmentSpotCheckRecordDetail)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-spot-check-record-detail/create | 创建点检记录子 |
| PUT | /eam/equipment-spot-check-record-detail/update | 更新点检记录子 |
| DELETE | /eam/equipment-spot-check-record-detail/delete | 删除点检记录子 |
| GET | /eam/equipment-spot-check-record-detail/get | 获得点检记录子 |
| GET | /eam/equipment-spot-check-record-detail/page | 获得点检记录子分页 |
| POST | /eam/equipment-spot-check-record-detail/senior | 高级搜索获得点检记录子分页 |
| GET | /eam/equipment-spot-check-record-detail/export-excel | 导出点检记录子 Excel |
| GET | /eam/equipment-spot-check-record-detail/get-import-template | 获得导入点检记录子模板 |
| POST | /eam/equipment-spot-check-record-detail/import | 导入点检记录子基本信息 |

#### 12.7.5 点检计划管理 (PlanSpotCheck)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/plan-spot-check/create | 创建点检计划 |
| PUT | /eam/plan-spot-check/update | 更新点检计划 |
| DELETE | /eam/plan-spot-check/delete | 删除点检计划 |
| GET | /eam/plan-spot-check/get | 获得点检计划 |
| GET | /eam/plan-spot-check/page | 获得点检计划分页 |
| POST | /eam/plan-spot-check/senior | 高级搜索获得点检计划分页 |
| GET | /eam/plan-spot-check/export-excel | 导出点检计划 Excel |
| GET | /eam/plan-spot-check/get-import-template | 获得导入点检计划模板 |
| POST | /eam/plan-spot-check/import | 导入点检计划基本信息 |
| POST | /eam/plan-spot-check/ables | 启用/禁用 |

### 12.8 维修管理

#### 12.8.1 报修申请管理 (EquipmentReportRepairRequest)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-report-repair-request/create | 创建设报修申请 |
| PUT | /eam/equipment-report-repair-request/update | 更新设报修申请 |
| POST | /eam/equipment-report-repair-request/updateOrder | 更新设备报修工单状态 |
| POST | /eam/equipment-report-repair-request/audiOrder | 审核更新设备报修工单状态 |
| POST | /eam/equipment-report-repair-request/updateCreateOrder | 更新设备报修工单状态并创建维修工单 |
| DELETE | /eam/equipment-report-repair-request/delete | 删除设报修申请 |
| GET | /eam/equipment-report-repair-request/get | 获得设报修申请 |
| GET | /eam/equipment-report-repair-request/page | 获得设报修申请分页 |
| POST | /eam/equipment-report-repair-request/senior | 高级搜索获得设报修申请分页 |
| POST | /eam/equipment-report-repair-request/fileListInfo | 获得报修设备模具附件文件列表 |
| GET | /eam/equipment-report-repair-request/export-excel | 导出设报修申请 Excel |
| GET | /eam/equipment-report-repair-request/get-import-template | 获得导入设报修申请模板 |
| POST | /eam/equipment-report-repair-request/import | 导入设报修申请基本信息 |

#### 12.8.2 维修工单管理 (EquipmentRepairJobMain)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-repair-job-main/create | 创建维修工单主 |
| POST | /eam/equipment-repair-job-main/createNew | 报修验证通过/不通过,创建维修工单主 |
| PUT | /eam/equipment-repair-job-main/update | 更新维修工单主 |
| POST | /eam/equipment-repair-job-main/updateRepair | 更新维修工单主 |
| POST | /eam/equipment-repair-job-main/execute | 完成维修工单主 |
| POST | /eam/equipment-repair-job-main/updateOrder | 更新维修工单主状态 |
| POST | /eam/equipment-repair-job-main/wxVerify | 维修验证更新工单主 |
| POST | /eam/equipment-repair-job-main/fallback | 更新维修工单主 |
| DELETE | /eam/equipment-repair-job-main/delete | 删除维修工单主 |
| GET | /eam/equipment-repair-job-main/get | 获得维修工单主 |
| GET | /eam/equipment-repair-job-main/page | 获得维修工单主分页 |
| POST | /eam/equipment-repair-job-main/senior | 高级搜索获得维修工单主分页 |
| GET | /eam/equipment-repair-job-main/export-excel | 导出维修工单主 Excel |
| GET | /eam/equipment-repair-job-main/get-import-template | 获得导入维修工单主模板 |
| POST | /eam/equipment-repair-job-main/import | 导入维修工单主基本信息 |

#### 12.8.3 维修工单明细管理 (EquipmentRepairJobDetail)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-repair-job-detail/create | 创建维修工单子 |
| PUT | /eam/equipment-repair-job-detail/update | 更新维修工单子 |
| DELETE | /eam/equipment-repair-job-detail/delete | 删除维修工单子 |
| GET | /eam/equipment-repair-job-detail/get | 获得维修工单子 |
| GET | /eam/equipment-repair-job-detail/getList | 获得维修工单子列表 |
| GET | /eam/equipment-repair-job-detail/page | 获得维修工单子分页 |
| POST | /eam/equipment-repair-job-detail/senior | 高级搜索获得维修工单子分页 |
| GET | /eam/equipment-repair-job-detail/export-excel | 导出维修工单子 Excel |
| GET | /eam/equipment-repair-job-detail/get-import-template | 获得导入维修工单子模板 |
| POST | /eam/equipment-repair-job-detail/ables | 启用/禁用 |

#### 12.8.4 维修记录管理 (EquipmentRepairRecordMain)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-repair-record-main/create | 创建设备维修记录主 |
| PUT | /eam/equipment-repair-record-main/update | 更新设备维修记录主 |
| DELETE | /eam/equipment-repair-record-main/delete | 删除设备维修记录主 |
| GET | /eam/equipment-repair-record-main/get | 获得设备维修记录主 |
| GET | /eam/equipment-repair-record-main/page | 获得设备维修记录主分页 |
| POST | /eam/equipment-repair-record-main/senior | 高级搜索获得设备维修记录主分页 |
| GET | /eam/equipment-repair-record-main/export-excel | 导出设备维修记录主 Excel |
| GET | /eam/equipment-repair-record-main/get-import-template | 获得导入设备维修记录主模板 |
| POST | /eam/equipment-repair-record-main/import | 导入设备维修记录主基本信息 |

#### 12.8.5 维修记录明细管理 (EquipmentRepairRecordDetail)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-repair-record-detail/create | 创建设备维修记录子 |
| PUT | /eam/equipment-repair-record-detail/update | 更新设备维修记录子 |
| DELETE | /eam/equipment-repair-record-detail/delete | 删除设备维修记录子 |
| GET | /eam/equipment-repair-record-detail/get | 获得设备维修记录子 |
| GET | /eam/equipment-repair-record-detail/page | 获得设备维修记录子分页 |
| POST | /eam/equipment-repair-record-detail/senior | 高级搜索获得设备维修记录子分页 |
| GET | /eam/equipment-repair-record-detail/export-excel | 导出设备维修记录子 Excel |
| GET | /eam/equipment-repair-record-detail/get-import-template | 获得导入设备维修记录子模板 |
| POST | /eam/equipment-repair-record-detail/import | 导入设备维修记录子基本信息 |

### 12.9 设备管理

#### 12.9.1 设备停机管理 (EquipmentShutdown)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-shutdown/create | 创建设备停机记录 |
| PUT | /eam/equipment-shutdown/update | 更新设备停机记录 |
| DELETE | /eam/equipment-shutdown/delete | 删除设备停机记录 |
| GET | /eam/equipment-shutdown/get | 获得设备停机记录 |
| GET | /eam/equipment-shutdown/page | 获得设备停机记录分页 |
| POST | /eam/equipment-shutdown/senior | 高级搜索获得设备停机记录分页 |
| GET | /eam/equipment-shutdown/export-excel | 导出设备停机记录 Excel |
| GET | /eam/equipment-shutdown/get-import-template | 获得导入设备停机记录模板 |
| POST | /eam/equipment-shutdown/import | 导入设备停机记录基本信息 |

#### 12.9.2 设备到货签收管理 (EquipmentSigning)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-signing/create | 创建设备到货签收记录 |
| POST | /eam/equipment-signing/creates | 批量创建设备到货签收记录 |
| PUT | /eam/equipment-signing/update | 更新设备到货签收记录 |
| DELETE | /eam/equipment-signing/delete | 删除设备到货签收记录 |
| GET | /eam/equipment-signing/get | 获得设备到货签收记录 |
| GET | /eam/equipment-signing/page | 获得设备到货签收记录分页 |
| POST | /eam/equipment-signing/senior | 高级搜索获得设备到货签收记录分页 |
| GET | /eam/equipment-signing/export-excel | 导出设备到货签收记录 Excel |
| GET | /eam/equipment-signing/get-import-template | 获得导入设备到货签收记录模板 |
| POST | /eam/equipment-signing/import | 导入设备到货签收记录基本信息 |

#### 12.9.3 设备转移记录管理 (EquipmentTransferRecord)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/equipment-transfer-record/create | 创建设备移动记录 |
| PUT | /eam/equipment-transfer-record/update | 更新设备移动记录 |
| DELETE | /eam/equipment-transfer-record/delete | 删除设备移动记录 |
| GET | /eam/equipment-transfer-record/get | 获得设备移动记录 |
| GET | /eam/equipment-transfer-record/page | 获得设备移动记录分页 |
| POST | /eam/equipment-transfer-record/senior | 高级搜索获得设备移动记录分页 |
| GET | /eam/equipment-transfer-record/export-excel | 导出设备移动记录 Excel |
| GET | /eam/equipment-transfer-record/get-import-template | 获得导入设备移动记录模板 |
| POST | /eam/equipment-transfer-record/import | 导入设备移动记录基本信息 |

### 12.10 备件库存管理

#### 12.10.1 备件申请管理 (SparePartsApplyMain/Detail)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/spare-parts-apply-main/create | 创建备件申请主 |
| PUT | /eam/spare-parts-apply-main/update | 更新备件申请主 |
| DELETE | /eam/spare-parts-apply-main/delete | 删除备件申请主 |
| GET | /eam/spare-parts-apply-main/get | 获得备件申请主 |
| GET | /eam/spare-parts-apply-main/list | 获得备件申请主列表 |
| GET | /eam/spare-parts-apply-main/page | 获得备件申请主分页 |
| POST | /eam/spare-parts-apply-main/senior | 高级搜索获得备件申请主分页 |
| GET | /eam/spare-parts-apply-main/export-excel | 导出备件申请主 Excel |
| GET | /eam/spare-parts-apply-main/get-import-template | 获得导入备件申请主模板 |
| POST | /eam/spare-parts-apply-detail/create | 创建备件申请子 |
| PUT | /eam/spare-parts-apply-detail/update | 更新备件申请子 |
| DELETE | /eam/spare-parts-apply-detail/delete | 删除备件申请子 |
| GET | /eam/spare-parts-apply-detail/get | 获得备件申请子 |
| GET | /eam/spare-parts-apply-detail/list | 获得备件申请子列表 |
| GET | /eam/spare-parts-apply-detail/page | 获得备件申请子分页 |
| GET | /eam/spare-parts-apply-detail/export-excel | 导出备件申请子 Excel |
| GET | /eam/spare-parts-apply-detail/get-import-template | 获得导入备件申请子模板 |

#### 12.10.2 备件入库管理 (SparePartsInLocationMain/Detail)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/spare-parts-in-location-main/create | 创建备件入库主 |
| PUT | /eam/spare-parts-in-location-main/update | 更新备件入库主 |
| DELETE | /eam/spare-parts-in-location-main/delete | 删除备件入库主 |
| GET | /eam/spare-parts-in-location-main/get | 获得备件入库主 |
| GET | /eam/spare-parts-in-location-main/list | 获得备件入库主列表 |
| GET | /eam/spare-parts-in-location-main/page | 获得备件入库主分页 |
| POST | /eam/spare-parts-in-location-main/senior | 高级搜索获得备件入库主分页 |
| GET | /eam/spare-parts-in-location-main/export-excel | 导出备件入库主 Excel |
| GET | /eam/spare-parts-in-location-main/get-import-template | 获得导入备件入库主模板 |
| POST | /eam/spare-parts-in-location-detail/create | 创建备件入库子 |
| PUT | /eam/spare-parts-in-location-detail/update | 更新备件入库子 |
| DELETE | /eam/spare-parts-in-location-detail/delete | 删除备件入库子 |
| GET | /eam/spare-parts-in-location-detail/get | 获得备件入库子 |
| GET | /eam/spare-parts-in-location-detail/list | 获得备件入库子列表 |
| GET | /eam/spare-parts-in-location-detail/page | 获得备件入库子分页 |
| GET | /eam/spare-parts-in-location-detail/export-excel | 导出备件入库子 Excel |
| GET | /eam/spare-parts-in-location-detail/get-import-template | 获得导入备件入库子模板 |

#### 12.10.3 备件出库管理 (SparePartsOutLocationMain/Detail)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/spare-parts-out-location-main/create | 创建领用出库主 |
| PUT | /eam/spare-parts-out-location-main/update | 更新领用出库主 |
| DELETE | /eam/spare-parts-out-location-main/delete | 删除领用出库主 |
| GET | /eam/spare-parts-out-location-main/get | 获得领用出库主 |
| GET | /eam/spare-parts-out-location-main/list | 获得领用出库主列表 |
| GET | /eam/spare-parts-out-location-main/page | 获得领用出库主分页 |
| POST | /eam/spare-parts-out-location-main/senior | 高级搜索获得领用出库主分页 |
| GET | /eam/spare-parts-out-location-main/export-excel | 导出领用出库主 Excel |
| GET | /eam/spare-parts-out-location-main/get-import-template | 获得导入领用出库主模板 |
| POST | /eam/spare-parts-out-location-detail/create | 创建领用出库子 |
| PUT | /eam/spare-parts-out-location-detail/update | 更新领用出库子 |
| DELETE | /eam/spare-out出库子/delete | 删除领用出库子 |
| GET | /eam/spare-parts-out-location-detail/get | 获得领用出库子 |
| GET | /eam/spare-parts-out-location-detail/list | 获得领用出库子列表 |
| GET | /eam/spare-parts-out-location-detail/page | 获得领用出库子分页 |
| GET | /eam/spare-parts-out-location-detail/export-excel | 导出领用出库子 Excel |
| GET | /eam/spare-parts-out-location-detail/get-import-template | 获得导入领用出库子模板 |

#### 12.10.4 备件入库记录管理 (SparePartsInLocationRecordMain/Detail)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/spare-parts-in-location-main-record/create | 创建备件入库记录主 |
| PUT | /eam/spare-parts-in-location-main-record/update | 更新备件入库记录主 |
| DELETE | /eam/spare-parts-in-location-main-record/delete | 删除备件入库记录主 |
| GET | /eam/spare-parts-in-location-main-record/get | 获得备件入库记录主 |
| GET | /eam/spare-parts-in-location-main-record/list | 获得备件入库记录主列表 |
| GET | /eam/spare-parts-in-location-main-record/page | 获得备件入库记录主分页 |
| POST | /eam/spare-parts-in-location-main-record/senior | 高级搜索获得备件入库记录主分页 |
| GET | /eam/spare-parts-in-location-main-record/export-excel | 导出备件入库记录主 Excel |
| GET | /eam/spare-parts-in-location-main-record/get-import-template | 获得导入备件入库记录主模板 |
| POST | /eam/spare-parts-in-location-detail-record/create | 创建备件入库记录子 |
| PUT | /eam/spare-parts-in-location-detail-record/update | 更新备件入库记录子 |
| DELETE | /eam/spare-parts-in-location-detail-record/delete | 删除备件入库记录子 |
| GET | /eam/spare-parts-in-location-detail-record/get | 获得备件入库记录子 |
| GET | /eam/spare-parts-in-location-detail-record/list | 获得备件入库记录子列表 |
| GET | /eam/spare-parts-in-location-detail-record/page | 获得备件入库记录子分页 |
| GET | /eam/spare-parts-in-location-detail-record/export-excel | 导出备件入库记录子 Excel |
| GET | /eam/spare-parts-in-location-detail-record/get-import-template | 获得导入备件入库记录子模板 |

#### 12.10.5 备件出库记录管理 (SparePartsOutLocationRecordMain/Detail)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/spare-parts-out-location-main-record/create | 创建领用出库记录主 |
| PUT | /eam/spare-parts-out-location-main-record/update | 更新领用出库记录主 |
| DELETE | /eam/spare-parts-out-location-main-record/delete | 删除领用出库记录主 |
| GET | /eam/spare-parts-out-location-main-record/get | 获得领用出库记录主 |
| GET | /eam/spare-parts-out-location-main-record/list | 获得领用出库记录主列表 |
| GET | /eam/spare-parts-out-location-main-record/page | 获得领用出库记录主分页 |
| POST | /eam/spare-parts-out-location-main-record/senior | 高级搜索获得领用出库记录主分页 |
| GET | /eam/spare-parts-out-location-main-record/export-excel | 导出领用出库记录主 Excel |
| GET | /eam/spare-parts-out-location-main-record/get-import-template | 获得导入领用出库记录主模板 |
| POST | /eam/spare-parts-out-location-detail-record/create | 创建领用出库记录子 |
| PUT | /eam/spare-parts-out-location-detail-record/update | 更新领用出库记录子 |
| DELETE | /eam/spare-parts-out-location-detail-record/delete | 删除领用出库记录子 |
| GET | /eam/spare-parts-out-location-detail-record/get | 获得领用出库记录子 |
| GET | /eam/spare-parts-out-location-detail-record/list | 获得领用出库记录子列表 |
| GET | /eam/spare-parts-out-location-detail-record/page | 获得领用出库记录子分页 |
| GET | /eam/spare-parts-out-location-detail-record/export-excel | 导出领用出库记录子 Excel |
| GET | /eam/spare-parts-out-location-detail-record/get-import-template | 获得导入领用出库记录子模板 |

### 12.11 盘点管理

#### 12.11.1 备件盘点任务管理 (CountJobMain/Detail)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/countJobMain/create | 创建备件盘点任务主 |
| PUT | /eam/countJobMain/update | 更新备件盘点任务主 |
| DELETE | /eam/countJobMain/delete | 删除备件盘点任务主 |
| GET | /eam/countJobMain/handleMainExport | 导出备件盘点工单 Excel |
| GET | /eam/countJobMain/get | 获得备件盘点任务主 |
| GET | /eam/countJobMain/list | 获得备件盘点任务主列表 |
| GET | /eam/countJobMain/page | 获得备件盘点任务主分页 |
| POST | /eam/countJobMain/senior | 高级搜索获得备件盘点任务主分页 |
| GET | /eam/countJobMain/export-excel | 导出备件盘点任务主 Excel |
| GET | /eam/countJobMain/get-import-template | 获得导入备件盘点任务主模板 |
| POST | /eam/countJobDetail/create | 创建备件盘点任务子 |
| PUT | /eam/countJobDetail/update | 更新备件盘点任务子 |
| DELETE | /eam/countJobDetail/delete | 删除备件盘点任务子 |
| GET | /eam/countJobDetail/get | 获得备件盘点任务子 |
| GET | /eam/countJobDetail/list | 获得备件盘点任务子列表 |
| GET | /eam/countJobDetail/page | 获得备件盘点任务子分页 |
| GET | /eam/countJobDetail/export-excel | 导出备件盘点任务子 Excel |
| GET | /eam/countJobDetail/get-import-template | 获得导入备件盘点任务子模板 |
| POST | /eam/countJobDetail/import | 导入备件盘点任务子基本信息 |

#### 12.11.2 备件盘点记录管理 (CountRecordMain/Detail)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/countRecordMain/create | 创建备件盘点记录主 |
| PUT | /eam/countRecordMain/update | 更新备件盘点记录主 |
| DELETE | /eam/countRecordMain/delete | 删除备件盘点记录主 |
| GET | /eam/countRecordMain/get | 获得备件盘点记录主 |
| GET | /eam/countRecordMain/list | 获得备件盘点记录主列表 |
| GET | /eam/countRecordMain/page | 获得备件盘点记录主分页 |
| POST | /eam/countRecordMain/senior | 高级搜索获得备件盘点记录主分页 |
| GET | /eam/countRecordMain/adjust | 盘点调整 |
| GET | /eam/countRecordMain/export-excel | 导出备件盘点记录主 Excel |
| GET | /eam/countRecordMain/get-import-template | 获得导入备件盘点记录主模板 |
| POST | /eam/countRecordDetail/create | 创建备件盘点记录子 |
| PUT | /eam/countRecordDetail/update | 更新备件盘点记录子 |
| DELETE | /eam/countRecordDetail/delete | 删除备件盘点记录子 |
| GET | /eam/countRecordDetail/get | 获得备件盘点记录子 |
| GET | /eam/countRecordDetail/list | 获得备件盘点记录子列表 |
| GET | /eam/countRecordDetail/page | 获得备件盘点记录子分页 |
| GET | /eam/countRecordDetail/export-excel | 导出备件盘点记录子 Excel |
| GET | /eam/countRecordDetail/get-import-template | 获得导入备件盘点记录子模板 |

#### 12.11.3 备件盘点调整记录管理 (AdjustRecordMain/Detail)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/adjustRecordMain/create | 创建备件盘点调整记录主 |
| PUT | /eam/adjustRecordMain/update | 更新备件盘点调整记录主 |
| DELETE | /eam/adjustRecordMain/delete | 删除备件盘点调整记录主 |
| GET | /eam/adjustRecordMain/get | 获得备件盘点调整记录主 |
| GET | /eam/adjustRecordMain/list | 获得备件盘点调整记录主列表 |
| GET | /eam/adjustRecordMain/page | 获得备件盘点调整记录主分页 |
| POST | /eam/adjustRecordMain/senior | 高级搜索获得备件盘点调整记录主分页 |
| GET | /eam/adjustRecordMain/export-excel | 导出备件盘点调整记录主 Excel |
| GET | /eam/adjustRecordMain/get-import-template | 获得导入备件盘点调整记录主模板 |
| POST | /eam/adjustRecordDetail/create | 创建备件盘点调整记录子 |
| PUT | /eam/adjustRecordDetail/update | 更新备件盘点调整记录子 |
| DELETE | /eam/adjustRecordDetail/delete | 删除备件盘点调整记录子 |
| GET | /eam/adjustRecordDetail/get | 获得备件盘点调整记录子 |
| GET | /eam/adjustRecordDetail/list | 获得备件盘点调整记录子列表 |
| GET | /eam/adjustRecordDetail/page | 获得备件盘点调整记录子分页 |
| GET | /eam/adjustRecordDetail/export-excel | 导出备件盘点调整记录子 Excel |
| GET | /eam/adjustRecordDetail/get-import-template | 获得导入备件盘点调整记录子模板 |

#### 12.11.4 备件盘点计划管理 (CountadjustPlan)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/countadjust-plan/create | 创建备件盘点计划 |
| PUT | /eam/countadjust-plan/update | 更新备件盘点计划 |
| DELETE | /eam/countadjust-plan/delete | 删除备件盘点计划 |
| GET | /eam/countadjust-plan/get | 获得备件盘点计划 |
| GET | /eam/countadjust-plan/list | 获得备件盘点计划列表 |
| GET | /eam/countadjust-plan/page | 获得备件盘点计划分页 |
| POST | /eam/countadjust-plan/senior | 高级搜索获得备件盘点计划分页 |
| GET | /eam/countadjust-plan/export-excel | 导出备件盘点计划 Excel |
| GET | /eam/countadjust-plan/get-import-template | 获得导入备件盘点计划模板 |

### 12.12 工装模具管理

#### 12.12.1 工装台账管理 (ToolAccounts)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/tool/tool-accounts/create | 创建工装台账 |
| PUT | /eam/tool/tool-accounts/update | 更新工装台账 |
| DELETE | /eam/tool/tool-accounts/delete | 删除工装台账 |
| GET | /eam/tool/tool-accounts/get | 获得工装台账 |
| GET | /eam/tool/tool-accounts/page | 获得工装台账分页 |
| POST | /eam/tool/tool-accounts/senior | 高级搜索获得工装台账分页 |
| GET | /eam/tool/tool-accounts/export-excel | 导出工装台账 Excel |
| POST | /eam/tool/tool-accounts/export-excel-senior | 高级搜索导出巡检项 Excel |
| GET | /eam/tool/tool-accounts/get-import-template | 获得导入工装台账模板 |
| POST | /eam/tool/tool-accounts/import | 导入工装台账基本信息 |
| GET | /eam/tool/tool-accounts/noPage | 获得工装台账不分页 |

#### 12.12.2 工装入库管理 (ToolEquipmentIn)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/tool/tool-equipment-in/create | 创建工装入库记录 |
| PUT | /eam/tool/tool-equipment-in/update | 更新工装入库记录 |
| DELETE | /eam/tool/tool-equipment-in/delete | 删除工装入库记录 |
| GET | /eam/tool/tool-equipment-in/get | 获得工装入库记录 |
| GET | /eam/tool/tool-equipment-in/page | 获得工装入库记录分页 |
| POST | /eam/tool/tool-equipment-in/senior | 高级搜索获得工装入库记录分页 |
| GET | /eam/tool/tool-equipment-in/export-excel | 导出工装入库记录 Excel |
| GET | /eam/tool/tool-equipment-in/get-import-template | 获得导入工装入库记录模板 |
| POST | /eam/tool/tool-equipment-in/import | 导入工装入库记录基本信息 |

#### 12.12.3 工装出库管理 (ToolEquipmentOut)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/tool/tool-equipment-out/create | 创建工装出库记录 |
| PUT | /eam/tool/tool-equipment-out/update | 更新工装出库记录 |
| DELETE | /eam/tool/tool-equipment-out/delete | 删除工装出库记录 |
| GET | /eam/tool/tool-equipment-out/get | 获得工装出库记录 |
| GET | /eam/tool/tool-equipment-out/page | 获得工装出库记录分页 |
| POST | /eam/tool/tool-equipment-out/senior | 高级搜索获得工装出库记录分页 |
| GET | /eam/tool/tool-equipment-out/export-excel | 导出工装出库记录 Excel |
| GET | /eam/tool/tool-equipment-out/get-import-template | 获得导入工装出库记录模板 |
| POST | /eam/tool/tool-equipment-out/import | 导入工装出库记录基本信息 |

#### 12.12.4 工装/设备到货签收管理 (ToolSigning)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/tool-signing/create | 创建工装到货签收记录 |
| POST | /eam/tool-signing/createNew | 创建设备到货签收记录 |
| PUT | /eam/tool-signing/update | 更新工装到货签收记录 |
| DELETE | /eam/tool-signing/delete | 删除工装到货签收记录 |
| GET | /eam/tool-signing/get | 获得工装到货签收记录 |
| GET | /eam/tool-signing/page | 获得工装到货签收记录分页 |
| POST | /eam/tool-signing/senior | 高级搜索获得工装到货签收记录分页 |
| GET | /eam/tool-signing/export-excel | 导出工装到货签收记录 Excel |
| GET | /eam/tool-signing/get-import-template | 获得导入工装到货签收记录模板 |
| POST | /eam/tool-signing/import | 导入工装到货签收记录基本信息 |

#### 12.12.5 设备/工装变更记录管理 (ToolChangedRecord / RecordEquipmentChanged)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/record-equipment-changed/create | 创建设备变更记录 |
| PUT | /eam/record-equipment-changed/update | 更新设备变更记录 |
| DELETE | /eam/record-equipment-changed/delete | 删除设备变更记录 |
| GET | /eam/record-equipment-changed/get | 获得设备变更记录 |
| GET | /eam/record-equipment-changed/page | 获得设备变更记录分页 |
| POST | /eam/record-equipment-changed/senior | 高级搜索获得设备变更记录分页 |
| GET | /eam/record-equipment-changed/export-excel | 导出设备变更记录 Excel |
| GET | /eam/record-equipment-changed/get-import-template | 获得导入设备变更记录模板 |
| POST | /eam/record-equipment-changed/import | 导入设备变更记录基本信息 |
| POST | /eam/tool-changed-record/create | 创建设备变更记录 |
| PUT | /eam/tool-changed-record/update | 更新设备变更记录 |
| DELETE | /eam/tool-changed-record/delete | 删除设备变更记录 |
| GET | /eam/tool-changed-record/get | 获得设备变更记录 |
| GET | /eam/tool-changed-record/page | 获得设备变更记录分页 |
| POST | /eam/tool-changed-record/senior | 高级搜索获得设备变更记录分页 |
| GET | /eam/tool-changed-record/export-excel | 导出设备变更记录 Excel |
| GET | /eam/tool-changed-record/get-import-template | 获得导入设备变更记录模板 |
| POST | /eam/tool-changed-record/import | 导入设备变更记录基本信息 |

#### 12.12.6 设备/工装与备件关联管理 (EquipmentToolSparePart)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/relation/equipment-tool-spare-part/create | 创建设备或工装与备件关联 |
| PUT | /eam/relation/equipment-tool-spare-part/update | 更新设备或工装与备件关联 |
| DELETE | /eam/relation/equipment-tool-spare-part/delete | 删除设备或工装与备件关联 |
| GET | /eam/relation/equipment-tool-spare-part/get | 获得设备或工装与备件关联 |
| GET | /eam/relation/equipment-tool-spare-part/page | 获得设备或工装与备件关联分页 |
| POST | /eam/relation/equipment-tool-spare-part/senior | 高级搜索获得设备或工装与备件关联分页 |
| GET | /eam/relation/equipment-tool-spare-part/export-excel | 导出设备或工装与备件关联 Excel |
| GET | /eam/relation/equipment-tool-spare-part/get-import-template | 获得导入设备或工装与备件关联模板 |
| POST | /eam/relation/equipment-tool-spare-part/import | 导入设备或工装与备件关联基本信息 |
| POST | /eam/relation/equipment-tool-spare-part/createBatch | 批量创建设备备件关系 |
| GET | /eam/relation/equipment-tool-spare-part/noPage | 获得设备或工装与备件关联不分页 |

### 12.13 经验记录管理

#### 12.13.1 保养经验记录 (MaintainExperience)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /record/maintain-experience/create | 创建保养经验记录 |
| PUT | /record/maintain-experience/update | 更新保养经验记录 |
| DELETE | /record/maintain-experience/delete | 删除保养经验记录 |
| GET | /record/maintain-experience/get | 获得保养经验记录 |
| GET | /record/maintain-experience/page | 获得保养经验记录分页 |
| POST | /record/maintain-experience/senior | 高级搜索获得保养经验记录分页 |
| GET | /record/maintain-experience/export-excel | 导出保养经验记录 Excel |
| GET | /record/maintain-experience/get-import-template | 获得导入保养经验记录模板 |
| POST | /record/maintain-experience/import | 导入保养经验记录基本信息 |

#### 12.13.2 维修经验记录 (RepairExperience)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /record/repair-experience/create | 创建维修经验记录 |
| PUT | /record/repair-experience/update | 更新维修经验记录 |
| DELETE | /record/repair-experience/delete | 删除维修经验记录 |
| GET | /record/repair-experience/get | 获得维修经验记录 |
| GET | /record/repair-experience/page | 获得维修经验记录分页 |
| POST | /record/repair-experience/senior | 高级搜索获得维修经验记录分页 |
| GET | /record/repair-experience/export-excel | 导出维修经验记录 Excel |
| GET | /record/repair-experience/get-import-template | 获得导入维修经验记录模板 |
| POST | /record/repair-experience/import | 导入维修经验记录基本信息 |
| POST | /record/repair-experience/ables | 启用/禁用 |

### 12.14 备件维修管理

#### 12.14.1 备件维修申请管理 (RepairSparePartsRequest)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/repair-spare-parts-request/create | 创建备件维修申请 |
| PUT | /eam/repair-spare-parts-request/update | 更新备件维修申请 |
| DELETE | /eam/repair-spare-parts-request/delete | 删除备件维修申请 |
| GET | /eam/repair-spare-parts-request/get | 获得备件维修申请 |
| GET | /eam/repair-spare-parts-request/page | 获得备件维修申请分页 |
| POST | /eam/repair-spare-parts-request/senior | 高级搜索获得备件维修申请分页 |
| GET | /eam/repair-spare-parts-request/export-excel | 导出备件维修申请 Excel |
| GET | /eam/repair-spare-parts-request/get-import-template | 获得导入备件维修申请模板 |
| POST | /eam/repair-spare-parts-request/ables | 启用/禁用 |

#### 12.14.2 备件维修记录管理 (RepairSparePartsRecord)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/repair-spare-parts-record/create | 创建备件维修记录 |
| PUT | /eam/repair-spare-parts-record/update | 更新备件维修记录 |
| DELETE | /eam/repair-spare-parts-record/delete | 删除备件维修记录 |
| GET | /eam/repair-spare-parts-record/get | 获得备件维修记录 |
| GET | /eam/repair-spare-parts-record/page | 获得备件维修记录分页 |
| POST | /eam/repair-spare-parts-record/senior | 高级搜索获得备件维修记录分页 |
| GET | /eam/repair-spare-parts-record/export-excel | 导出备件维修记录 Excel |
| GET | /eam/repair-spare-parts-record/get-import-template | 获得导入备件维修记录模板 |
| POST | /eam/repair-spare-parts-record/import | 导入备件维修记录基本信息 |

### 12.15 关系管理

#### 12.15.1 设备主要部件关联 (RelationMainPart)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/relation-main-part/create | 创建设备主要部件关联 |
| PUT | /eam/relation-main-part/update | 更新设备主要部件关联 |
| DELETE | /eam/relation-main-part/delete | 删除设备主要部件关联 |
| GET | /eam/relation-main-part/get | 获得设备主要部件关联 |
| GET | /eam/relation-main-part/page | 获得设备主要部件关联分页 |
| POST | /eam/relation-main-part/senior | 高级搜索获得设备主要部件关联分页 |
| GET | /eam/relation-main-part/export-excel | 导出设备主要部件关联 Excel |
| GET | /eam/relation-main-part/get-import-template | 获得导入设备主要部件关联模板 |
| POST | /eam/relation-main-part/import | 导入设备主要部件关联基本信息 |

#### 12.15.2 巡检方案与项关联 (RelationInspectionOptionItem / RelationInspectionPlanItem)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/relation-inspection-option-item/create | 创建巡检方案与项关联 |
| PUT | /eam/relation-inspection-option-item/update | 更新巡检方案与项关联 |
| DELETE | /eam/relation-inspection-option-item/delete | 删除巡检方案与项关联 |
| GET | /eam/relation-inspection-option-item/get | 获得巡检方案与项关联 |
| GET | /eam/relation-inspection-option-item/page | 获得巡检方案与项关联分页 |
| POST | /eam/relation-inspection-option-item/senior | 高级搜索获得巡检方案与项关联分页 |
| POST | /eam/relation-inspection-plan-item/create | 创建巡检计划与项关联 |
| PUT | /eam/relation-inspection-plan-item/update | 更新巡检计划与项关联 |
| DELETE | /eam/relation-inspection-plan-item/delete | 删除巡检计划与项关联 |
| GET | /eam/relation-inspection-plan-item/get | 获得巡检计划与项关联 |
| GET | /eam/relation-inspection-plan-item/page | 获得巡检计划与项关联分页 |
| POST | /eam/relation-inspection-plan-item/senior | 高级搜索获得巡检计划与项关联分页 |

#### 12.15.3 保养方案/计划与项关联 (RelationMaintenanceOptionItem / RelationMaintenancePlanItem)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/relation-maintenance-option-item/create | 创建保养方案与项关联 |
| PUT | /eam/relation-maintenance-option-item/update | 更新保养方案与项关联 |
| DELETE | /eam/relation-maintenance-option-item/delete | 删除保养方案与项关联 |
| GET | /eam/relation-maintenance-option-item/get | 获得保养方案与项关联 |
| GET | /eam/relation-maintenance-option-item/page | 获得保养方案与项关联分页 |
| POST | /eam/relation-maintenance-option-item/senior | 高级搜索获得保养方案与项关联分页 |
| POST | /eam/relation-maintenance-plan-item/create | 创建保养计划与项关联 |
| PUT | /eam/relation-maintenance-plan-item/update | 更新保养计划与项关联 |
| DELETE | /eam/relation-maintenance-plan-item/delete | 删除保养计划与项关联 |
| GET | /eam/relation-maintenance-plan-item/get | 获得保养计划与项关联 |
| GET | /eam/relation-maintenance-plan-item/page | 获得保养计划与项关联分页 |
| POST | /eam/relation-maintenance-plan-item/senior | 高级搜索获得保养计划与项关联分页 |

#### 12.15.4 点检方案/计划与项关联 (RelationSpotCheckOptionItem / RelationSpotCheckPlanItem)

| HTTP方法 | API路径 | 接口说明 |
|----------|---------|----------|
| POST | /eam/relation-spot-check-option-item/create | 创建点检方案与项关联 |
| PUT | /eam/relation-spot-check-option-item/update | 更新点检方案与项关联 |
| DELETE | /eam/relation-spot-check-option-item/delete | 删除点检方案与项关联 |
| GET | /eam/relation-spot-check-option-item/get | 获得点检方案与项关联 |
| GET | /eam/relation-spot-check-option-item/page | 获得点检方案与项关联分页 |
| POST | /eam/relation-spot-check-option-item/senior | 高级搜索获得点检方案与项关联分页 |
| POST | /eam/relation-spot-check-plan-item/create | 创建点检计划与项关联 |
| PUT | /eam/relation-spot-check-plan-item/update | 更新点检计划与项关联 |
| DELETE | /eam/relation-spot-check-plan-item/delete | 删除点检计划与项关联 |
| GET | /eam/relation-spot-check-plan-item/get | 获得点检计划与项关联 |
| GET | /eam/relation-spot-check-plan-item/page | 获得点检计划与项关联分页 |
| POST | /eam/relation-spot-check-plan-item/senior | 高级搜索获得点检计划与项关联分页 |

### 12.16 开发规范

### 12.16.1 命名规范
- Controller: `XxxController.java`
- Service接口: `XxxService.java`
- Service实现: `XxxServiceImpl.java`
- Mapper: `XxxMapper.java`
- DO: `XxxDO.java`
- VO请求: `XxxCreateReqVO.java` / `XxxUpdateReqVO.java` / `XxxPageReqVO.java`
- VO响应: `XxxRespVO.java`

### 12.16.2 分页规范
- 请求: `XxxPageReqVO` 继承分页请求基类
- 响应: `PageResult<XxxRespVO>`
- 默认页大小: 10/20/50/100

### 12.16.3 状态流转
设备状态机遵循: `正常 -> 故障/保养/闲置 -> 正常/报废`
