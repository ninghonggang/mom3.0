# QMS质量模块设计文档

## 1. 模块概述

### 1.1 模块职责
QMS（Quality Management System）质量管理模块是SFMS3.0系统中的核心业务模块，负责企业产品质量检验与控制的全流程管理。模块遵循ISO/TS质量管理体系标准，提供从原材料入库到成品出厂的全方位质量检验解决方案。

### 1.2 业务价值
- 实现质量检验的标准化、流程化管理
- 依据AQL标准确定抽样方案，保证检验科学性
- 支持定性与定量检验，灵活适应不同产品需求
- 完整记录检验数据，支持质量追溯与统计分析
- 与生产、仓库等模块协同，形成完整的质量管控闭环

## 2. 技术架构

### 2.1 模块结构
```
win-module-qms/
├── win-module-qms-api/          # API接口层
│   └── src/main/java/com/win/module/qms/
│       ├── api/                 # 对外暴露的API接口
│       │   └── InspectionRequest/ # 检验请求API
│       └── enums/               # 枚举常量定义
└── win-module-qms-biz/          # 业务实现层
    └── src/main/java/com/win/module/qms/
        ├── controller/          # 控制器层
        │   ├── aql/            # AQL标准
        │   ├── counter/        # 计数器
        │   ├── dynamicRule/    # 动态规则
        │   ├── inspectionCharacteristics/ # 检验特性
        │   ├── inspectionJob/  # 检验任务
        │   ├── inspectionMethod/ # 检验方法
        │   ├── inspectionProcess/ # 检验工序
        │   ├── inspectionRecord/ # 检验记录
        │   ├── inspectionRequest/ # 检验请求
        │   ├── inspectionScheme/ # 检验方案
        │   ├── inspectionStage/ # 检验阶段
        │   ├── programmeTemplate/ # 方案模板
        │   ├── q1/q2/q3/       # 质量检验类型
        │   ├── sampleCode/     # 样本编码
        │   ├── samplingProcess/ # 采样工序
        │   ├── samplingScheme/ # 采样方案
        │   ├── selectedProject/ # 选中项目
        │   └── selectedSet/    # 选中集合
        ├── service/             # 服务层
        ├── dal/                 # 数据访问层
        │   ├── dataobject/      # 数据对象(DO)
        │   └── mysql/           # MyBatis Mapper接口
        ├── convert/             # 对象转换
        └── enums/               # 业务枚举
```

### 2.2 技术栈
- **框架**: Spring Boot + MyBatis Plus
- **数据库**: MySQL
- **API风格**: RESTful API，Swagger/OpenAPI 3.0文档化
- **权限控制**: `@PreAuthorize` 注解方式
- **日志记录**: `@OperateLog` 注解
- **数据校验**: `@Valid` + Bean Validation

## 3. 核心领域模型

### 3.1 AQL标准 (AqlDO)
Acceptable Quality Level，接收质量限标准。

| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 主键 |
| code | String | AQL代码 |
| name | String | AQL名称 |
| level | String | 检验水平 |
| value | BigDecimal | AQL值 |
| samplingAmount | Integer | 抽样数量 |

### 3.2 检验特性 (InspectionCharacteristicsDO)
定义产品需要检验的质量特性。

| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 主键 |
| code | String | 特性编码 |
| name | String | 特性名称 |
| inspectionMethod | String | 检验方法 |
| inspectionTool | String | 检验工具 |
| qualificationType | String | 定性/定量类型 |
| usl | BigDecimal | 上规格限 |
| lsl | BigDecimal | 下规格限 |
| criticalDefect | BigDecimal | 严重缺陷 |
| majorDefect | BigDecimal | 主要缺陷 |
| minorDefect | BigDecimal | 轻微缺陷 |

### 3.3 检验方法 (InspectionMethodDO)
定义检验所使用的方法标准。

| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 主键 |
| code | String | 方法编码 |
| name | String | 方法名称 |
| type | String | 方法类型 |
| description | String | 方法描述 |

### 3.4 检验工序 (InspectionProcessDO)
定义检验的工序流程。

| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 主键 |
| code | String | 工序编码 |
| name | String | 工序名称 |
| orderNum | Integer | 工序顺序 |
| stageId | Long | 所属阶段ID |

### 3.5 检验阶段 (InspectionStageDO)
定义检验的不同阶段。

| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 主键 |
| code | String | 阶段编码 |
| name | String | 阶段名称 |
| type | String | 阶段类型 |
| orderNum | Integer | 顺序号 |

### 3.6 检验方案 (InspectionSchemeDO)
检验方案主表，定义检验的整体方案。

| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 主键 |
| code | String | 方案编码 |
| name | String | 方案名称 |
| productCode | String | 产品编码 |
| aqlId | Long | AQL标准ID |
| inspectionLevel | String | 检验水平 |
| samplingPlan | String | 抽样方案 |
| status | String | 状态 |
| effectiveDate | LocalDateTime | 生效日期 |
| expiryDate | LocalDateTime | 失效日期 |

### 3.7 检验任务 (InspectionJob)
#### InspectionJobMainDO - 检验任务主表
| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 主键 |
| number | String | 任务单据号 |
| schemeId | Long | 检验方案ID |
| productCode | String | 产品编码 |
| quantity | BigDecimal | 检验数量 |
| status | String | 任务状态 |
| inspector | String | 检验员 |
| inspectionDate | LocalDateTime | 检验日期 |

#### InspectionJobDetailDO - 检验任务明细
| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 主键 |
| masterId | Long | 主表ID |
| characteristicsId | Long | 检验特性ID |
| result | String | 检验结果 |
| value | BigDecimal | 检验值 |

#### InspectionJobPackageDO - 检验任务包装
| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 主键 |
| number | String | 单据号 |
| masterId | Long | 主表ID |
| packageCode | String | 包装号 |
| packageSpecificationCode | String | 包装规格 |
| amount | BigDecimal | 数量 |
| measuringUnit | String | 计量单位 |
| sampleAmount | BigDecimal | 采样数量 |

#### InspectionJobCharacteristicsDO - 检验任务特性
记录每个检验特性的执行情况。

### 3.8 检验记录 (InspectionRecord)
#### InspectionRecordMainDO - 检验记录主表
| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 主键 |
| number | String | 记录单据号 |
| jobId | Long | 检验任务ID |
| inspector | String | 检验员 |
| inspectionTime | LocalDateTime | 检验时间 |
| result | String | 总体结果 |
| qualifiedCount | Integer | 合格数量 |
| unqualifiedCount | Integer | 不合格数量 |

#### InspectionRecordDetailDO - 检验记录明细
记录每项检验的具体数据。

#### InspectionRecordPackageDO - 检验记录包装
记录包装级别的检验结果。

#### InspectionRecordQualitativeDO - 定性检验记录
针对定性指标的检验记录（合格/不合格）。

#### InspectionRecordQuantifyDO - 定量检验记录
针对定量指标（数值）的检验记录。

### 3.9 检验请求 (InspectionRequest)
#### InspectionRequestMainDO - 检验请求主表
来自其他模块（如仓库）的检验申请。

#### InspectionRequestPackageDO - 检验请求包装
包装级别的检验请求。

### 3.10 动态规则 (DynamicRuleDO)
支持动态配置的检验规则。

### 3.11 采样方案 (SamplingSchemeDO, SamplingProcessDO)
定义采样规则和采样工序。

### 3.12 样本编码 (SampleCodeDO)
管理检验样本的唯一编码。

### 3.13 计数器 (CounterDO)
提供各种业务单据的序号生成。

## 4. 核心功能分析

### 4.1 AQL标准管理
```
功能列表:
- 创建AQL标准
- 更新AQL标准
- 删除AQL标准
- 查询AQL标准列表
- 验证AQL值
```

**关键API**:
```java
GET /qms/aql/list           # AQL标准列表
POST /qms/aql/create        # 创建AQL
PUT  /qms/aql/update        # 更新AQL
DELETE /qms/aql/delete     # 删除AQL
```

### 4.2 检验特性管理
```
功能列表:
- 创建检验特性
- 更新检验特性
- 删除检验特性
- 查询检验特性
- 导出/导入
```

### 4.3 检验方案管理
```
功能列表:
- 创建检验方案
- 配置检验项目和抽样规则
- 关联AQL标准
- 方案启用/停用
- 方案导入/导出
```

### 4.4 检验任务管理
```
功能列表:
- 创建检验任务
- 分配检验特性
- 执行检验记录
- 检验结果判定
- 任务完成确认
```

**关键API**:
```java
GET  /qms/inspection-job-package/list    # 检验任务包装列表
GET  /qms/inspection-job-package/page     # 检验任务包装分页
```

### 4.5 检验记录管理
```
功能列表:
- 记录定性检验结果
- 记录定量检验结果
- 包装级别检验记录
- 特性级别检验记录
- 检验数据导出
```

**关键API**:
```java
POST /qms/inspection-record/qualitative/create  # 定性检验记录
POST /qms/inspection-record/quantify/create    # 定量检验记录
```

### 4.6 检验请求管理
接收来自其他模块的检验申请。

```
功能列表:
- 创建检验请求
- 更新请求状态
- 查询请求列表
- 关联检验任务
```

### 4.7 动态规则引擎
支持灵活配置的质量检验规则。

```
功能列表:
- 创建动态规则
- 配置规则条件
- 配置规则动作
- 规则启用/停用
```

## 5. 业务流程

### 5.1 来料检验流程 (IQC)
```
1. 仓库创建检验请求 (InspectionRequest)
2. QMS创建检验任务 (InspectionJob)
3. 检验员执行检验 (InspectionRecord)
4. 根据AQL判定合格数量
5. 合格 → 入库; 不合格 → 不良品处理
```

### 5.2 过程检验流程 (PQC)
```
1. 生产工单触发检验
2. 按检验方案创建检验任务
3. 执行各工序检验
4. 记录定性/定量结果
5. 判定是否通过
```

### 5.3 出货检验流程 (OQC)
```
1. 销售订单触发出货检验
2. 按照检验方案执行检验
3. 全数检验或抽样检验
4. 生成出货检验报告
5. 允许出货
```

## 6. 关键技术实现

### 6.1 AQL抽样标准
基于GB/T 2828.1/ISO 2859.1标准实现：
- 正常检验抽样方案
- 加严检验抽样方案
- 放宽检验抽样方案
- 特殊检验水平

### 6.2 检验判定逻辑
```java
// 定量检验判定
result = (value >= lsl && value <= usl) ? PASS : FAIL;

// 定性检验判定
result = (observed == expected) ? PASS : FAIL;

// 基于缺陷等级的综合判定
if (criticalDefectCount > 0) → REJECT;
if (majorDefectCount > aql) → REJECT;
if (minorDefectCount > aql) → REJECT;
```

### 6.3 与其他模块集成
- **仓库模块(WMS)**: 接收入库检验请求，反馈检验结果
- **生产模块(MES)**: 接收过程检验任务，反馈检验结果
- **系统模块(System)**: 获取用户信息、字典数据

## 7. 权限设计

QMS模块采用基于注解的权限控制：

| 权限标识 | 说明 |
|----------|------|
| qms:aql:create | 创建AQL标准 |
| qms:aql:update | 更新AQL标准 |
| qms:aql:delete | 删除AQL标准 |
| qms:aql:query | 查询AQL标准 |
| qms:aql:export | 导出AQL标准 |
| qms:inspection-characteristics:create | 创建检验特性 |
| qms:inspection-characteristics:update | 更新检验特性 |
| qms:inspection-characteristics:delete | 删除检验特性 |
| qms:inspection-characteristics:query | 查询检验特性 |
| qms:inspection-scheme:create | 创建检验方案 |
| qms:inspection-scheme:update | 更新检验方案 |
| qms:inspection-scheme:delete | 删除检验方案 |
| qms:inspection-scheme:query | 查询检验方案 |
| qms:inspection-job:create | 创建检验任务 |
| qms:inspection-job:execute | 执行检验任务 |
| qms:inspection-job:query | 查询检验任务 |
| qms:inspection-record:create | 创建检验记录 |
| qms:inspection-record:query | 查询检验记录 |

## 8. 表结构汇总

| 表名 | 说明 |
|------|------|
| qms_aql | AQL标准表 |
| qms_inspection_characteristics | 检验特性表 |
| qms_inspection_method | 检验方法表 |
| qms_inspection_process | 检验工序表 |
| qms_inspection_stage | 检验阶段表 |
| qms_inspection_scheme | 检验方案表 |
| qms_inspection_scheme_detail | 检验方案明细表 |
| qms_job_inspection_main | 检验任务主表 |
| qms_job_inspection_detail | 检验任务明细表 |
| qms_job_inspection_package | 检验任务包装表 |
| qms_job_inspection_characteristics | 检验任务特性表 |
| qms_record_inspection_main | 检验记录主表 |
| qms_record_inspection_detail | 检验记录明细表 |
| qms_record_inspection_package | 检验记录包装表 |
| qms_record_inspection_qualitative | 定性检验记录表 |
| qms_record_inspection_quantify | 定量检验记录表 |
| qms_request_inspection_main | 检验请求主表 |
| qms_request_inspection_package | 检验请求包装表 |
| qms_dynamic_rule | 动态规则表 |
| qms_sampling_scheme | 采样方案表 |
| qms_sampling_process | 采样工序表 |
| qms_sample_code | 样本编码表 |
| qms_counter | 计数器表 |

## 9. 开发规范

### 9.1 命名规范
遵循项目统一命名规范：
- Controller: `XxxController.java`
- Service: `XxxService.java` / `XxxServiceImpl.java`
- DO: `XxxDO.java`
- VO请求: `XxxCreateReqVO.java` / `XxxUpdateReqVO.java` / `XxxPageReqVO.java`
- VO响应: `XxxRespVO.java`

### 9.2 状态流转
```
检验任务状态:
  待检验 → 检验中 → 已完成/已取消

检验记录结果:
  待判定 → 合格 → 不合格

检验请求状态:
  待受理 → 已受理 → 检验中 → 已完成
```

### 9.3 异常处理
检验相关异常：
- 检验特性未找到
- 检验方案未生效
- 检验数量超过库存
- AQL标准不匹配

## 10. 完整API清单

### 10.1 AQL标准 (/qms/aql)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/aql/create | 创建AQL标准 |
| PUT | /qms/aql/update | 更新AQL标准 |
| DELETE | /qms/aql/delete | 删除AQL标准 |
| GET | /qms/aql/get | 获得AQL标准 |
| GET | /qms/aql/list | 获得AQL标准列表 |
| GET | /qms/aql/page | 获得AQL标准分页 |
| POST | /qms/aql/senior | 高级搜索获得AQL分页 |
| GET | /qms/aql/export-excel | 导出AQL Excel |
| POST | /qms/aql/export-excel-senior | 导出AQL Excel（高级搜索） |
| GET | /qms/aql/get-import-template | 获得导入AQL模板 |
| POST | /qms/aql/import | 导入AQL基本信息 |
| POST | /qms/aql/enable | 启用AQL |
| POST | /qms/aql/disable | 禁用AQL |

### 10.2 物料检验计数器 (/qms/counter)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/counter/create | 创建物料检验计数器 |
| PUT | /qms/counter/update | 更新物料检验计数器 |
| DELETE | /qms/counter/delete | 删除物料检验计数器 |
| GET | /qms/counter/get | 获得物料检验计数器 |
| GET | /qms/counter/list | 获得物料检验计数器列表 |
| GET | /qms/counter/page | 获得物料检验计数器分页 |
| POST | /qms/counter/senior | 高级搜索获得物料检验计数器分页 |
| GET | /qms/counter/export-excel | 导出物料检验计数器 Excel |
| POST | /qms/counter/export-excel-senior | 导出物料检验计数器 Excel（高级搜索） |
| GET | /qms/counter/get-import-template | 获得导入物料检验计数器模板 |
| POST | /qms/counter/import | 导入物料检验计数器基本信息 |
| POST | /qms/counter/enable | 启用物料检验计数器 |
| POST | /qms/counter/disable | 禁用物料检验计数器 |
| GET | /qms/counter/getNextStage | 获取下一个阶段 |

### 10.3 动态修改规则 (/qms/dynamic-rule)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/dynamic-rule/create | 创建动态修改规则 |
| PUT | /qms/dynamic-rule/update | 更新动态修改规则 |
| DELETE | /qms/dynamic-rule/delete | 删除动态修改规则 |
| GET | /qms/dynamic-rule/get | 获得动态修改规则 |
| GET | /qms/dynamic-rule/list | 获得动态修改规则列表 |
| GET | /qms/dynamic-rule/page | 获得动态修改规则分页 |
| POST | /qms/dynamic-rule/senior | 高级搜索获得动态修改规则分页 |
| GET | /qms/dynamic-rule/export-excel | 导出动态修改规则 Excel |
| POST | /qms/dynamic-rule/export-excel-senior | 导出动态修改规则 Excel（高级搜索） |
| GET | /qms/dynamic-rule/get-import-template | 获得导入动态修改规则模板 |
| POST | /qms/dynamic-rule/import | 导入动态修改规则基本信息 |
| POST | /qms/dynamic-rule/enable | 启用动态修改规则 |
| POST | /qms/dynamic-rule/disable | 禁用动态修改规则 |

### 10.4 检验特性 (/basic/inspection-characteristics)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /basic/inspection-characteristics/create | 创建检验特性 |
| PUT | /basic/inspection-characteristics/update | 更新检验特性 |
| DELETE | /basic/inspection-characteristics/delete | 删除检验特性 |
| GET | /basic/inspection-characteristics/get | 获得检验特性 |
| GET | /basic/inspection-characteristics/list | 获得检验特性列表 |
| GET | /basic/inspection-characteristics/page | 获得检验特性分页 |
| GET | /basic/inspection-characteristics/export-excel | 导出检验特性 Excel |
| GET | /basic/inspection-characteristics/get-import-template | 获得导入检验特性模板 |
| POST | /basic/inspection-characteristics/import | 导入检验特性基本信息 |

### 10.5 检验方法 (/qms/inspection-method)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/inspection-method/create | 创建检验方法 |
| PUT | /qms/inspection-method/update | 更新检验方法 |
| DELETE | /qms/inspection-method/delete | 删除检验方法 |
| GET | /qms/inspection-method/get | 获得检验方法 |
| GET | /qms/inspection-method/list | 获得检验方法列表 |
| GET | /qms/inspection-method/page | 获得检验方法分页 |
| POST | /qms/inspection-method/senior | 高级搜索获得检验方法分页 |
| GET | /qms/inspection-method/export-excel | 导出检验方法 Excel |
| POST | /qms/inspection-method/export-excel-senior | 导出检验方法 Excel（高级搜索） |
| GET | /qms/inspection-method/get-import-template | 获得导入检验方法模板 |
| POST | /qms/inspection-method/import | 导入检验方法基本信息 |
| POST | /qms/inspection-method/enable | 启用检验方法 |
| POST | /qms/inspection-method/disable | 禁用检验方法 |

### 10.6 检验工序 (/qms/inspection-process)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/inspection-process/create | 创建检验工序 |
| PUT | /qms/inspection-process/update | 更新检验工序 |
| DELETE | /qms/inspection-process/delete | 删除检验工序 |
| GET | /qms/inspection-process/get | 获得检验工序 |
| GET | /qms/inspection-process/list | 获得检验工序列表 |
| GET | /qms/inspection-process/page | 获得检验工序分页 |
| GET | /qms/inspection-process/export-excel | 导出检验工序 Excel |
| GET | /qms/inspection-process/get-import-template | 获得导入检验工序模板 |
| POST | /qms/inspection-process/import | 导入检验工序基本信息 |
| GET | /qms/inspection-process/getListByTempleteCode | 根据模版code获取工序列表以及检验特性 |

### 10.7 检验阶段 (/qms/inspection-stage)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/inspection-stage/create | 创建检验阶段 |
| PUT | /qms/inspection-stage/update | 更新检验阶段 |
| DELETE | /qms/inspection-stage/delete | 删除检验阶段 |
| GET | /qms/inspection-stage/get | 获得检验阶段 |
| GET | /qms/inspection-stage/list | 获得检验阶段列表 |
| GET | /qms/inspection-stage/page | 获得检验阶段分页 |
| POST | /qms/inspection-stage/senior | 高级搜索获得检验阶段分页 |
| GET | /qms/inspection-stage/export-excel | 导出检验阶段 Excel |
| GET | /qms/inspection-stage/get-import-template | 获得导入检验阶段模板 |
| POST | /qms/inspection-stage/import | 导入检验阶段基本信息 |
| GET | /qms/inspection-stage/noPage | 获得检验阶段不分页 |

### 10.8 检验方案 (/qms/inspection-scheme)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/inspection-scheme/create | 创建检验方案 |
| PUT | /qms/inspection-scheme/update | 更新检验方案 |
| DELETE | /qms/inspection-scheme/delete | 删除检验方案 |
| GET | /qms/inspection-scheme/get | 获得检验方案 |
| GET | /qms/inspection-scheme/list | 获得检验方案列表 |
| GET | /qms/inspection-scheme/page | 获得检验方案分页 |
| POST | /qms/inspection-scheme/senior | 高级搜索获得检验方案分页 |
| GET | /qms/inspection-scheme/export-excel | 导出检验方案 Excel |
| POST | /qms/inspection-scheme/export-excel-senior | 导出检验方案 Excel（高级搜索） |
| GET | /qms/inspection-scheme/get-import-template | 获得导入检验方案模板 |
| POST | /qms/inspection-scheme/import | 导入检验方案基本信息 |
| POST | /qms/inspection-scheme/enable | 启用检验方案 |
| POST | /qms/inspection-scheme/disable | 禁用检验方案 |

### 10.9 检验方案模板 (/qms/programme-template)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/programme-template/create | 创建检验方案模板 |
| PUT | /qms/programme-template/update | 更新检验方案模板 |
| DELETE | /qms/programme-template/delete | 删除检验方案模板 |
| GET | /qms/programme-template/get | 获得检验方案模板 |
| GET | /qms/programme-template/list | 获得检验方案模板列表 |
| GET | /qms/programme-template/page | 获得检验方案模板分页 |
| POST | /qms/programme-template/senior | 高级搜索获得检验方案模板分页 |
| GET | /qms/programme-template/export-excel | 导出检验方案模板 Excel |
| POST | /qms/programme-template/export-excel-senior | 导出检验方案模板 Excel（高级搜索） |
| GET | /qms/programme-template/get-import-template | 获得导入检验方案模板模板 |
| POST | /qms/programme-template/import | 导入检验方案模板基本信息 |
| POST | /qms/programme-template/enable | 启用检验方案模板 |
| POST | /qms/programme-template/disable | 禁用检验方案模板 |

### 10.10 Q1通知单 (/qms/inspectionQ1)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/inspectionQ1/create | 创建Q1通知单 |
| PUT | /qms/inspectionQ1/update | 更新Q1通知单 |
| DELETE | /qms/inspectionQ1/delete | 删除Q1通知单 |
| GET | /qms/inspectionQ1/get | 获得Q1通知单 |
| GET | /qms/inspectionQ1/page | 获得Q1通知单分页 |
| POST | /qms/inspectionQ1/senior | 高级搜索获得Q1通知单分页 |
| GET | /qms/inspectionQ1/export-excel | 导出Q1通知单 Excel |
| GET | /qms/inspectionQ1/get-import-template | 获得导入Q1通知单模板 |
| GET | /qms/inspectionQ1/finish | 完成Q1通知单 |

### 10.11 Q2通知单 (/qms/inspectionQ2)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| GET | /qms/inspectionQ2/send | Q2通知单发送邮件 |
| POST | /qms/inspectionQ2/create | 创建Q2通知单 |
| PUT | /qms/inspectionQ2/update | 更新Q2通知单 |
| DELETE | /qms/inspectionQ2/delete | 删除Q2通知单 |
| GET | /qms/inspectionQ2/get | 获得Q2通知单 |
| GET | /qms/inspectionQ2/page | 获得Q2通知单分页 |
| POST | /qms/inspectionQ2/senior | 高级搜索获得Q2通知单分页 |
| GET | /qms/inspectionQ2/export-excel | 导出Q2通知单 Excel |
| GET | /qms/inspectionQ2/get-import-template | 获得导入Q2通知单模板 |
| GET | /qms/inspectionQ2/finish | 完成Q2通知单 |
| GET | /qms/inspectionQ2/getEmail | 获取系统中的email地址 |

### 10.12 Q3通知单 (/qms/inspectionQ3)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/inspectionQ3/create | 创建Q3通知单 |
| PUT | /qms/inspectionQ3/update | 更新Q3通知单 |
| DELETE | /qms/inspectionQ3/delete | 删除Q3通知单 |
| GET | /qms/inspectionQ3/get | 获得Q3通知单 |
| GET | /qms/inspectionQ3/page | 获得Q3通知单分页 |
| POST | /qms/inspectionQ3/senior | 高级搜索获得Q3通知单分页 |
| GET | /qms/inspectionQ3/export-excel | 导出Q3通知单 Excel |
| GET | /qms/inspectionQ3/get-import-template | 获得导入Q3通知单模板 |
| GET | /qms/inspectionQ3/finish | 完成Q3通知单 |

### 10.13 Q3通知单主 (/qms/inspection-Q3-main)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/inspection-Q3-main/create | 创建Q3通知单主 |
| PUT | /qms/inspection-Q3-main/update | 更新Q3通知单主 |
| DELETE | /qms/inspection-Q3-main/delete | 删除Q3通知单主 |
| GET | /qms/inspection-Q3-main/get | 获得Q3通知单主 |
| GET | /qms/inspection-Q3-main/page | 获得Q3通知单主分页 |
| POST | /qms/inspection-Q3-main/senior | 高级搜索获得Q3通知单主分页 |
| GET | /qms/inspection-Q3-main/export-excel | 导出Q3通知单主 Excel |
| POST | /qms/inspection-Q3-main/export-excel-senior | 导出Q3通知单主 Excel（高级搜索） |
| GET | /qms/inspection-Q3-main/finish | 完成Q3通知单主 |

### 10.14 Q3通知单子 (/qms/inspection-Q3-detail)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/inspection-Q3-detail/create | 创建Q3通知单子 |
| PUT | /qms/inspection-Q3-detail/update | 更新Q3通知单子 |
| DELETE | /qms/inspection-Q3-detail/delete | 删除Q3通知单子 |
| GET | /qms/inspection-Q3-detail/get | 获得Q3通知单子 |
| GET | /qms/inspection-Q3-detail/page | 获得Q3通知单子分页 |
| POST | /qms/inspection-Q3-detail/senior | 高级搜索获得Q3通知单子分页 |
| GET | /qms/inspection-Q3-detail/export-excel | 导出Q3通知单子 Excel |
| GET | /qms/inspection-Q3-detail/get-import-template | 获得导入Q3通知单子模板 |

### 10.15 检验任务主 (/qms/inspection-job-main)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| GET | /qms/inspection-job-main/get | 获得检验任务主 |
| GET | /qms/inspection-job-main/page | 获得检验任务主分页 |
| POST | /qms/inspection-job-main/senior | 高级搜索获得检验任务主分页 |
| GET | /qms/inspection-job-main/export-excel | 导出检验任务 Excel |
| POST | /qms/inspection-job-main/export-excel-senior | 导出检验任务主 Excel（高级搜索） |
| PUT | /qms/inspection-job-main/accept | 承接检验任务 |
| PUT | /qms/inspection-job-main/abandon | 放弃检验任务 |
| PUT | /qms/inspection-job-main/close | 关闭检验任务 |
| PUT | /qms/inspection-job-main/execute | 执行检验任务 |
| PUT | /qms/inspection-job-main/release | 发布检验任务 |
| POST | /qms/inspection-job-main/staging | 暂存检验任务 |

### 10.16 检验任务子(工序) (/qms/inspection-job-detail)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| GET | /qms/inspection-job-detail/page | 获得检验任务子（工序）分页 |
| POST | /qms/inspection-job-detail/senior | 高级搜索获得检验任务子分页 |
| GET | /qms/inspection-job-detail/list | 获得检验任务子列表 |
| GET | /qms/inspection-job-detail/getBySchemeCode | 首件检验根据检验编码获取检验任务子（工序） |

### 10.17 检验任务包装 (/qms/inspection-job-package)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| GET | /qms/inspection-job-package/list | 获得检验任务包装列表 |

### 10.18 检验任务检验特性 (/qms/inspection-job-characteristics)

> 注：该Controller仅注入Service，无独立API端点，作为子模块供其他Controller调用。

### 10.19 检验记录主 (/qms/inspection-record-main)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/inspection-record-main/create | 创建首件检验记录 |
| PUT | /qms/inspection-record-main/update | 更新检验记录 |
| PUT | /qms/inspection-record-main/publish | 发布检验记录 |
| DELETE | /qms/inspection-record-main/delete | 删除检验记录 |
| GET | /qms/inspection-record-main/get | 获得检验记录 |
| GET | /qms/inspection-record-main/page | 获得检验记录分页 |
| POST | /qms/inspection-record-main/senior | 高级搜索获得检验记录分页 |
| GET | /qms/inspection-record-main/export-excel | 导出检验记录 Excel |
| POST | /qms/inspection-record-main/export-excel-senior | 导出检验记录 Excel（高级搜索） |
| GET | /qms/inspection-record-main/export-excel-first | 导出首件检验记录 Excel |
| POST | /qms/inspection-record-main/export-excel-first-senior | 导出首件检验记录 Excel（高级搜索） |
| PUT | /qms/inspection-record-main/firstInspectionUpdate | 更新首件检验记录 |
| PUT | /qms/inspection-record-main/execute | 执行检验记录任务 |

### 10.20 检验记录子(工序) (/qms/inspection-record-detail)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| GET | /qms/inspection-record-detail/page | 获得检验记录子（工序）分页 |
| POST | /qms/inspection-record-detail/senior | 高级搜索获得检验记录子分页 |
| GET | /qms/inspection-record-detail/list | 根据masterId获得检验记录子列表 |

### 10.21 检验记录包装 (/qms/inspection-record-package)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| GET | /qms/inspection-record-package/list | 获得检验结果包装列表 |

### 10.22 检验记录检验特性 (/qms/inspection-record-characteristics)

> 注：该Controller仅注入Service，无独立API端点，作为子模块供其他Controller调用。

### 10.23 定性结果 (/qms/inspection-record-qualitative)

> 注：该Controller仅注入Service，无独立API端点，作为子模块供其他Controller调用。

### 10.24 定量结果 (/qms/inspection-record-quantify)

> 注：该Controller仅注入Service，无独立API端点，作为子模块供其他Controller调用。

### 10.25 检验申请主 (/qms/inspection-request-main)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/inspection-request-main/create | 创建检验申请 |
| PUT | /qms/inspection-request-main/update | 更新检验申请 |
| DELETE | /qms/inspection-request-main/delete | 删除检验申请 |
| GET | /qms/inspection-request-main/get | 获得检验申请 |
| GET | /qms/inspection-request-main/page | 获得检验申请分页 |
| POST | /qms/inspection-request-main/senior | 高级搜索获得检验申请分页 |
| GET | /qms/inspection-request-main/export-excel | 导出检验申请 Excel |
| GET | /qms/inspection-request-main/get-import-template | 获得导入检验申请模板 |
| PUT | /qms/inspection-request-main/close | 关闭检验申请 |
| PUT | /qms/inspection-request-main/submit | 提交检验申请 |
| PUT | /qms/inspection-request-main/reAdd | 重新添加检验申请 |
| PUT | /qms/inspection-request-main/agree | 审批通过检验申请 |
| PUT | /qms/inspection-request-main/handle | 执行检验申请 |
| PUT | /qms/inspection-request-main/refused | 审批拒绝检验申请 |

### 10.26 检验申请包装 (/qms/inspection-request-package)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/inspection-request-package/create | 创建检验申请包装 |
| PUT | /qms/inspection-request-package/update | 更新检验申请包装 |
| DELETE | /qms/inspection-request-package/delete | 删除检验申请包装 |
| GET | /qms/inspection-request-package/get | 获得检验申请包装 |
| GET | /qms/inspection-request-package/list | 获得检验申请包装列表 |

### 10.27 采样方案 (/qms/sampling-scheme)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/sampling-scheme/create | 创建采样方案 |
| PUT | /qms/sampling-scheme/update | 更新采样方案 |
| DELETE | /qms/sampling-scheme/delete | 删除采样方案 |
| GET | /qms/sampling-scheme/get | 获得采样方案 |
| GET | /qms/sampling-scheme/list | 获得采样方案列表 |
| GET | /qms/sampling-scheme/page | 获得采样方案分页 |
| POST | /qms/sampling-scheme/senior | 高级搜索获得采样方案分页 |
| GET | /qms/sampling-scheme/export-excel | 导出采样方案 Excel |
| POST | /qms/sampling-scheme/export-excel-senior | 导出采样方案 Excel（高级搜索） |
| GET | /qms/sampling-scheme/get-import-template | 获得导入采样方案模板 |
| POST | /qms/sampling-scheme/import | 导入采样方案基本信息 |
| GET | /qms/sampling-scheme/get-available-list | 查询可用的采样方案 |
| POST | /qms/sampling-scheme/enable | 启用采样方案 |
| POST | /qms/sampling-scheme/disable | 禁用采样方案 |

### 10.28 采样过程 (/qms/sampling-process)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/sampling-process/create | 创建采样过程 |
| PUT | /qms/sampling-process/update | 更新采样过程 |
| DELETE | /qms/sampling-process/delete | 删除采样过程 |
| GET | /qms/sampling-process/get | 获得采样过程 |
| GET | /qms/sampling-process/page | 获得采样过程分页 |
| POST | /qms/sampling-process/senior | 高级搜索获得采样过程分页 |
| GET | /qms/sampling-process/export-excel | 导出采样过程 Excel |
| POST | /qms/sampling-process/export-excel-senior | 导出采样过程 Excel（高级搜索） |
| GET | /qms/sampling-process/get-import-template | 获得导入采样过程模板 |
| POST | /qms/sampling-process/import | 导入采样过程基本信息 |
| POST | /qms/sampling-process/enable | 启用采样过程 |
| POST | /qms/sampling-process/disable | 禁用采样过程 |

### 10.29 选定集 (/qms/selected-set)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/selected-set/create | 创建选定集 |
| PUT | /qms/selected-set/update | 更新选定集 |
| DELETE | /qms/selected-set/delete | 删除选定集 |
| GET | /qms/selected-set/get | 获得选定集 |
| GET | /qms/selected-set/list | 获得选定集列表 |
| GET | /qms/selected-set/page | 获得选定集分页 |
| POST | /qms/selected-set/senior | 高级搜索获得选定集分页 |
| GET | /qms/selected-set/export-excel | 导出选定集 Excel |
| POST | /qms/selected-set/export-excel-senior | 导出选定集 Excel（高级搜索） |
| GET | /qms/selected-set/get-import-template | 获得导入选定集模板 |
| POST | /qms/selected-set/import | 导入选定集基本信息 |
| POST | /qms/selected-set/enable | 启用选定集 |
| POST | /qms/selected-set/disable | 禁用选定集 |

### 10.30 选定集项目 (/qms/selected-project)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/selected-project/create | 创建选定集项目 |
| PUT | /qms/selected-project/update | 更新选定集项目 |
| DELETE | /qms/selected-project/delete | 删除选定集项目 |
| GET | /qms/selected-project/get | 获得选定集项目 |
| GET | /qms/selected-project/list | 获得选定集项目列表 |
| GET | /qms/selected-project/page | 获得选定集项目分页 |
| POST | /qms/selected-project/senior | 高级搜索获得选定集项目分页 |
| GET | /qms/selected-project/export-excel | 导出选定集项目 Excel |
| GET | /qms/selected-project/get-import-template | 获得导入选定集项目模板 |
| POST | /qms/selected-project/import | 导入选定集项目基本信息 |
| GET | /qms/selected-project/noPage | 获得选定集项目不分页 |

### 10.31 样本字码 (/qms/sample-code)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /qms/sample-code/create | 创建样本字码 |
| PUT | /qms/sample-code/update | 更新样本字码 |
| DELETE | /qms/sample-code/delete | 删除样本字码 |
| GET | /qms/sample-code/get | 获得样本字码 |
| GET | /qms/sample-code/list | 获得样本字码列表 |
| GET | /qms/sample-code/page | 获得样本字码分页 |
| POST | /qms/sample-code/senior | 高级搜索获得样本字码分页 |
| GET | /qms/sample-code/export-excel | 导出样本字码 Excel |
| POST | /qms/sample-code/export-excel-senior | 导出样本字码 Excel（高级搜索） |
| GET | /qms/sample-code/get-import-template | 获得导入样本字码模板 |
| POST | /qms/sample-code/import | 导入样本字码基本信息 |
| POST | /qms/sample-code/enable | 启用样本字码 |
| POST | /qms/sample-code/disable | 禁用样本字码 |

### 10.32 暂存JSON数据 (/job/staging-json-job)

| HTTP方法 | URL路径 | 接口说明 |
|----------|---------|---------|
| POST | /job/staging-json-job/create | 创建暂存JSON数据 |
| PUT | /job/staging-json-job/update | 更新暂存JSON数据 |
| DELETE | /job/staging-json-job/delete | 删除暂存JSON数据 |
| GET | /job/staging-json-job/get | 获得暂存JSON数据 |
| GET | /job/staging-json-job/page | 获得暂存JSON数据分页 |
| POST | /job/staging-json-job/senior | 高级搜索获得暂存JSON数据分页 |

### 10.33 API统计汇总

| 子模块 | API数量 | 说明 |
|--------|---------|------|
| aql | 14 | AQL标准管理 |
| counter | 14 | 物料检验计数器 |
| dynamicRule | 14 | 动态修改规则 |
| inspectionCharacteristics | 9 | 检验特性 |
| inspectionMethod | 14 | 检验方法 |
| inspectionProcess | 10 | 检验工序 |
| inspectionStage | 11 | 检验阶段 |
| inspectionScheme | 14 | 检验方案 |
| programmeTemplate | 14 | 检验方案模板 |
| inspectionQ1 | 9 | Q1通知单 |
| inspectionQ2 | 11 | Q2通知单 |
| inspectionQ3 | 9 | Q3通知单 |
| inspectionQ3Main | 10 | Q3通知单主 |
| inspectionQ3Detail | 8 | Q3通知单子 |
| inspectionJobMain | 11 | 检验任务主 |
| inspectionJobDetail | 4 | 检验任务子(工序) |
| inspectionJobPackage | 1 | 检验任务包装 |
| inspectionJobCharacteristics | 0 | 检验任务检验特性（无独立API） |
| inspectionRecordMain | 14 | 检验记录主 |
| inspectionRecordDetail | 3 | 检验记录子(工序) |
| inspectionRecordPackage | 1 | 检验记录包装 |
| inspectionRecordCharacteristics | 0 | 检验记录检验特性（无独立API） |
| inspectionRecordQualitative | 0 | 定性结果（无独立API） |
| inspectionRecordQuantify | 0 | 定量结果（无独立API） |
| inspectionRequestMain | 14 | 检验申请主 |
| inspectionRequestPackage | 5 | 检验申请包装 |
| samplingScheme | 14 | 采样方案 |
| samplingProcess | 14 | 采样过程 |
| selectedSet | 14 | 选定集 |
| selectedProject | 11 | 选定集项目 |
| sampleCode | 14 | 样本字码 |
| stagingJsonJob | 6 | 暂存JSON数据 |

> **说明**: API数量统计包含所有HTTP方法（POST/GET/PUT/DELETE）。部分Controller仅作为子模块提供服务，无独立REST端点。
