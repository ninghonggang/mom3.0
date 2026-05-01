# MES制造执行模块设计文档

## 1. 模块概述

MES（Manufacturing Execution System）制造执行系统模块是SFMS3.0的核心模块，负责车间级生产执行管理，涵盖生产计划、工序调度、质量检验、设备人员管理等核心制造流程。

**路径**: `win-module-mes`

**子模块结构**:
- `win-module-mes-api` - API接口定义
- `win-module-mes-biz` - 业务实现

---

## 2. 模块职责

| 子模块 | 职责 |
|--------|------|
| 生产日计划 (mesorderday) | 管理日生产计划的创建、发布、终止、恢复 |
| 生产排程 (mesworkscheduling) | 工单生成、工序任务调度执行 |
| 工艺路线 (processroute) | 产品工艺路线定义与工序管理 |
| BOM管理 (mesorderdaybom) | 生产物料清单管理 |
| 设备管理 (mesorderdayequipment) | 车间设备与工序绑定 |
| 人员班组 (mesorderdayworker) | 工序人员与班组配置 |
| 工位管理 (mesorderdayworkstation) | 生产工位配置 |
| 质检管理 (qms) | 质量检验表单与项目 |
| Holiday假期 | 班组排班日历 |
| 能力信息 (abilityinfo) | 人员技能能力档案 |
| 拆解管理 (mesdismantling) | 产品拆解任务 |
| 物料申请 (mesitemrequest) | 生产物料申请 |
| 报工管理 (mesjobreportlog) | 工序报工记录 |
| 返工返修 (mesrework) | 不良品返工返修流程 |

---

## 3. 核心类/接口

### 3.1 生产日计划核心类

**Controller**: `MesOrderDayController`
- 路径: `controller/mesorderday/MesOrderDayController.java`
- 职责: 生产日计划的CRUD、导入导出、发布、终止

**Service接口**: `MesOrderDayService`
```java
public interface MesOrderDayService {
    Long createOrderDay(MesOrderDayCreateReqVO createReqVO);
    Integer updateOrderDay(MesOrderDayUpdateReqVO updateReqVO);
    Integer deleteOrderDay(Long id);
    MesOrderDayDO getOrderDay(Long id);
    List<MesOrderDayDO> getOrderDayList(Collection<Long> ids);
    PageResult<MesOrderDayDO> getOrderDayPage(MesOrderDayPageReqVO pageReqVO);
    Integer publish(MesOrderDayCreateReqVO createReqVO);
    Integer stopPlan(Long id);
    Integer resumPlan(Long id);
}
```

**Service实现**: `MesOrderDayServiceImpl`
- 核心逻辑: 创建日计划时自动创建工艺路线实例、BOM实例、设备实例、工位实例、班组人员
- 发布计划时校验工序齐套性，生成工单

### 3.2 排程调度核心类

**Service**: `MesWorkSchedulingService` / `MesWorkSchedulingServiceImpl`
- 工单状态管理
- 工序任务生成
- 工单流转控制

**DO**: `MesWorkSchedulingDO`
- 关联日计划单号 `planMasterCode`
- 工单编码 `schedulingCode`
- 工序编码 `workingNode`

**DO**: `MesWorkSchedulingDetailDO`
- 工序级任务明细
- 状态机管理 (待领工/待报工/已完成)

---

## 4. 核心数据结构

### 4.1 生产日计划 DO

```java
@TableName("plan_mes_order_day")
public class MesOrderDayDO extends BaseDO {
    private Long id;              // 主键
    private String status;        // 状态: WAITSECHUDLE/待排产, PUBLISHED/已发布, PROCESSING/进行中, TERMINALE/已终止
    private String planNoDay;      // 计划单号
    private String productCode;    // 产品图号
    private String workroomCode;   // 车间编码
    private String lineCode;       // 产线编码
    private Integer planCount;     // 计划数量
    private String processrouteCode; // 工艺路线编码
    private String tempProcessroute; // 是否临时工艺
    private String standardBom;     // 标准BOM编码
    private String tempBom;        // 替代BOM标识
    private String workMode;       // 工单模式: BATCH批量/单件
    private String taskMode;       // 生产模式: ASSIGN派工/RECIEVE领工
    private LocalDateTime planDate; // 计划日期
    private String batchCode;      // 批次码
    private String completeInspect; // 齐套检查状态
    private Integer leftTaskCount; // 待完成工单数
}
```

### 4.2 枚举常量

```java
public interface DictTypeConstants {
    String MES_DO_STATUS = "mes_do_status";      // 工单执行状态
    String MES_WORKBILL_MODEL = "mes_workbill_model"; // 工单模式
    String MES_TASK_MODE = "mes_task_mode";      // 生产模式
    String MES_PLANDO_STATUS = "mes_plando_status"; // 计划执行状态
    String PLAN_TYPE_MES = "plan_type_mes";       // 计划类型
    String MES_PROCESS_STATUS = "mes_process_status"; // 工序状态
    String MES_REQUEST_TYPE = "mes_request_type"; // 物料申请类型
    String OPERSTEP_COLLECT_TYPE = "operstep_collect_type"; // 采集方式
}
```

---

## 5. 数据流向

### 5.1 日计划创建流程

```
用户提交日计划创建请求
        ↓
MesOrderDayController.createOrderDay()
        ↓
MesOrderDayService.createOrderDay()
        ↓
┌───────────────────────────────────────┐
│ 自动创建关联模型(多对一):              │
│   - mesOrderDayRouteService           │
│     → 创建工艺路线实例                  │
│   - mesOrderDayBomService             │
│     → 创建BOM实例                       │
│   - mesOrderDayEquipmentService        │
│     → 创建设备信息实例                 │
│   - mesOrderDayWorkstationService     │
│     → 创建工位信息实例                 │
│   - mesOrderDayWorkerService          │
│     → 派工模式下创建班组人员           │
└───────────────────────────────────────┘
        ↓
返回计划ID
```

### 5.2 日计划发布流程

```
用户调用 /publishPlan 接口
        ↓
MesOrderDayService.publish()
        ↓
┌─────────────────────────────────────┐
│ 前置校验:                            │
│   - 工序齐套检查                     │
│   - 设备配置校验                     │
│   - 人员配置校验                     │
│   - 工位配置校验                     │
└─────────────────────────────────────┘
        ↓
┌─────────────────────────────────────┐
│ 工单生成(getSchedulingSingle):       │
│   - 计算工单数量(批量/单件模式)      │
│   - 按工艺路线生成工序任务           │
│   - 绑定设备/工位/人员验证规则       │
└─────────────────────────────────────┘
        ↓
mesWorkSchedulingService.inertBacthTrans()
        ↓
更新日计划状态为已发布
```

### 5.3 工单执行状态流转

```
READYTODO(待报工)
    ↓
WAITBEGIN(待开工) --[齐套检查通过后]-->
PROCESSING(进行中)
    ↓
FINISH(已完成) / TERMINALE(已终止)
```

---

## 6. 关键技术实现

### 6.1 技术栈

- **持久层**: MyBatis-Plus + MyBatis
- **数据库**: MySQL
- **工具库**: Hutool(日期/JSON)、FastJSON
- **权限**: `@PreAuthorize` + `@ss.hasPermission`

### 6.2 核心实现模式

**事务管理**: `@Transactional` 注解控制

**批量操作**:
```java
mesWorkSchedulingService.inertBacthTrans(masterList, subList);
```

**配置获取**:
```java
InfaConfigUtil infaConfigUtil;
boolean preCheck = infaConfigUtil.getBooleanValue(MesConfigKeys.PRE_START_CHECK_ON, false);
```

**编号生成**:
```java
serialNumberApi.generateCode(RuleCodeEnum.PURCHASE_ORDER.getCode());
```

### 6.3 关键配置

- `MesConfigKeys.PRE_START_CHECK_ON` - 齐套检查开关
- `PlanBillStatusEnum` - 计划单状态枚举
- `WorkingScheduleEnum` - 工单状态枚举
- `TaskModeEnum` - 生产模式枚举

---

## 7. API接口

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| 创建日计划 | POST | /mes/orderday/create | 创建生产日计划 |
| 更新日计划 | PUT | /mes/orderday/update | 更新生产日计划 |
| 删除日计划 | DELETE | /mes/orderday/delete | 删除生产日计划 |
| 获取日计划 | GET | /mes/orderday/get | 获取单个日计划 |
| 日计划分页 | GET | /mes/orderday/page | 分页查询日计划 |
| 发布计划 | POST | /mes/orderday/publishPlan | 发布日计划生成工单 |
| 终止计划 | POST | /mes/orderday/stopPlan/{id} | 终止已发布计划 |
| 获取工艺信息 | GET | /mes/orderday/getProcessInfo | 获取工艺路线及工序 |
| 获取BOM信息 | GET | /mes/orderday/getBomInfo | 获取配置的BOM信息 |

---

## 8. 数据库表

| 表名 | 说明 |
|------|------|
| plan_mes_order_day | 生产日计划主表 |
| plan_mes_order_day_bom | 日计划BOM表 |
| plan_mes_order_day_route | 日计划工艺路线表 |
| plan_mes_order_day_routesub | 日计划工序明细表 |
| plan_mes_order_day_equipment | 日计划设备表 |
| plan_mes_order_day_worker | 日计划人员表 |
| plan_mes_order_day_workstation | 日计划工位表 |
| plan_mes_work_scheduling | 工单排程主表 |
| plan_mes_work_scheduling_detail | 工单排程明细表 |

---

## 9. 与其他模块的交互

- **BPM模块**: 工单审批流程
- **QMS模块**: 质检表单关联
- **System模块**: 用户权限、编号生成
- **EAM模块**: 设备基础数据

---

## 完整API清单

> 统计范围：`win-module-mes-biz/src/main/java/com/win/module/mes/controller/`
> 共计 **52个Controller**，约 **400+个API端点**

### 一、齐套检查

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| GET | /mes/complete-inspect/get | 获得齐套检查标识 | paramCode (String) | MesConfigInfoDO |
| POST | /mes/complete-inspect/get-orderDay-bom | 获取日计划Bom信息 | MesWorkSchedulingBaseVO | List\<MesOrderDayBomRespVO\> |
| POST | /mes/complete-inspect/get-orderDay-bom-page | 获取日计划Bom信息(分页) | MesWorkSchedulingPageReqVO | PageResult\<MesOrderDayBomRespVO\> |
| POST | /mes/complete-inspect/get-orderDay-worker-page | 获取日计划Worker信息(分页) | MesWorkSchedulingPageReqVO | PageResult\<MesOrderDayWorkerRespVO\> |
| POST | /mes/complete-inspect/get-orderDay-equipment-page | 获取日计划Equipment信息(分页) | MesWorkSchedulingPageReqVO | PageResult\<MesOrderDayEquipmentRespVO\> |
| POST | /mes/complete-inspect/get-orderDay-equipment | 获取日计划设备信息 | MesWorkSchedulingBaseVO | List\<MesOrderDayEquipmentRespVO\> |
| POST | /mes/complete-inspect/get-orderDay-worker | 获取日计划人员信息 | MesWorkSchedulingBaseVO | List\<MesOrderDayWorkerRespVO\> |
| POST | /mes/complete-inspect/update | 更新生产日工单 | MesWorkSchedulingUpdateReqVO | Boolean |

### 二、设备基本信息

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/device-info/create | 创建设备基本信息 | DeviceInfoCreateReqVO | Long(id) |
| PUT | /mes/device-info/update | 更新设备基本信息 | DeviceInfoUpdateReqVO | Boolean |
| DELETE | /mes/device-info/delete | 删除设备基本信息 | id (Long) | Boolean |
| GET | /mes/device-info/get | 获得设备基本信息 | id (Long) | DeviceInfoRespVO |
| GET | /mes/device-info/page | 获得设备基本信息分页 | DeviceInfoPageReqVO | PageResult\<DeviceInfoRespVO\> |
| POST | /mes/device-info/senior | 高级搜索设备基本信息 | CustomConditions | PageResult\<DeviceInfoRespVO\> |
| GET | /mes/device-info/export-excel | 导出设备基本信息Excel | DeviceInfoExportReqVO | Excel文件流 |
| GET | /mes/device-info/get-import-template | 获得导入设备基本信息模板 | - | Excel文件流 |
| POST | /mes/device-info/import | 导入设备基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |

### 三、节假日设置

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/holiday/create | 创建节假日设置 | HolidayCreateReqVO | Integer(id) |
| PUT | /mes/holiday/update | 更新节假日设置 | HolidayUpdateReqVO | Boolean |
| DELETE | /mes/holiday/delete | 删除节假日设置 | id (Integer) | Boolean |
| GET | /mes/holiday/get | 获得节假日设置 | id (Integer) | HolidayRespVO |
| GET | /mes/holiday/listByYear | 获得指定年份节假日数据(按keyDate分组) | year (String) | Map\<String,List\<HolidayDO\>\> |
| GET | /mes/holiday/list | 获得指定年份节假日数据列表 | year (String) | List\<HolidayRespVO\> |
| GET | /mes/holiday/page | 获得节假日设置分页 | HolidayPageReqVO | PageResult\<HolidayRespVO\> |
| POST | /mes/holiday/senior | 高级搜索节假日信息 | CustomConditions | PageResult\<HolidayRespVO\> |
| GET | /mes/holiday/export-excel | 导出节假日设置Excel | HolidayExportReqVO | Excel文件流 |
| GET | /mes/holiday/get-import-template | 获得导入节假日设置模板 | - | Excel文件流 |
| POST | /mes/holiday/import | 导入节假日设置基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |

### 四、能力矩阵

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/ability-info/create | 创建能力矩阵信息 | AbilityInfoCreateReqVO | Integer(id) |
| PUT | /mes/ability-info/update | 更新能力矩阵信息 | AbilityInfoUpdateReqVO | Boolean |
| DELETE | /mes/ability-info/delete | 删除能力矩阵信息 | id (Integer) | Boolean |
| GET | /mes/ability-info/get | 获得能力矩阵信息 | id (Integer) | AbilityInfoRespVO |
| GET | /mes/ability-info/list | 获得能力矩阵信息列表 | ids (Collection\<Integer\>) | List\<AbilityInfoRespVO\> |
| GET | /mes/ability-info/page | 获得能力矩阵信息分页 | AbilityInfoPageReqVO | PageResult\<AbilityInfoRespVO\> |
| POST | /mes/ability-info/senior | 高级搜索能力矩阵信息 | CustomConditions | PageResult\<AbilityInfoRespVO\> |
| GET | /mes/ability-info/export-excel | 导出能力矩阵信息Excel | AbilityInfoExportReqVO | Excel文件流 |
| GET | /mes/ability-info/get-import-template | 获得导入能力矩阵信息模板 | - | Excel文件流 |
| POST | /mes/ability-info/import | 导入能力矩阵信息基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |
| POST | /mes/hr-person-ability/create | 创建人员能力矩阵配置 | HrPersonAbilityCreateReqVO | Integer(id) |
| PUT | /mes/hr-person-ability/update | 更新人员能力矩阵配置 | HrPersonAbilityUpdateReqVO | Boolean |
| DELETE | /mes/hr-person-ability/delete | 删除人员能力矩阵配置 | id (Integer) | Boolean |
| GET | /mes/hr-person-ability/get | 获得人员能力矩阵配置 | id (Integer) | HrPersonAbilityRespVO |
| GET | /mes/hr-person-ability/list | 获得人员能力矩阵配置列表 | ids (Collection\<Integer\>) | List\<HrPersonAbilityRespVO\> |
| GET | /mes/hr-person-ability/page | 获得人员能力矩阵配置分页 | HrPersonAbilityPageReqVO | PageResult\<HrPersonAbilityRespVO\> |
| POST | /mes/hr-person-ability/senior | 高级搜索人员能力信息 | CustomConditions | PageResult\<HrPersonAbilityRespVO\> |
| GET | /mes/hr-person-ability/export-excel | 导出人员能力矩阵配置Excel | HrPersonAbilityExportReqVO | Excel文件流 |
| GET | /mes/hr-person-ability/get-import-template | 获得导入人员能力矩阵配置模板 | - | Excel文件流 |
| POST | /mes/hr-person-ability/import | 导入人员能力矩阵配置 | MultipartFile, mode, updatePart | Map\<String,Object\> |

### 五、叫料申请

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/item-request-main/create | 创建叫料申请单主 | MesItemRequestMainCreateReqVO | Long(id) |
| POST | /mes/item-request-main/create-call-material | 创建叫料申请单 | MesItemRequestMainCreateReqVO | Long(id) |
| PUT | /mes/item-request-main/update | 更新叫料申请单主 | MesItemRequestMainUpdateReqVO | Boolean |
| DELETE | /mes/item-request-main/delete | 删除叫料申请单主 | id (Long) | Boolean |
| GET | /mes/item-request-main/get | 获得叫料申请单主 | id (Long) | MesItemRequestMainRespVO |
| GET | /mes/item-request-main/page | 获得叫料申请单主分页 | MesItemRequestMainPageReqVO | PageResult\<MesItemRequestMainRespVO\> |
| POST | /mes/item-request-main/senior | 高级搜索叫料申请单主 | CustomConditions | PageResult\<MesItemRequestMainRespVO\> |
| GET | /mes/item-request-main/export-excel | 导出叫料申请单主Excel | MesItemRequestMainExportReqVO | Excel文件流 |
| GET | /mes/item-request-main/get-import-template | 获得导入叫料申请单主模板 | - | Excel文件流 |
| POST | /mes/item-request-main/addBasicItem | 创建补料申请 | JSONObject | Boolean |
| POST | /mes/item-request-main/receiveBasicItem | 领料 | MesItemRequestMainUpdateReqVO | Boolean |
| POST | /mes/item-request-main/receiveItem | 领料 | JSONObject | Boolean |
| POST | /mes/item-request-detail/create | 创建叫料申请明细 | MesItemRequestDetailCreateReqVO | Long(id) |
| PUT | /mes/item-request-detail/update | 更新叫料申请明细 | MesItemRequestDetailUpdateReqVO | Boolean |
| DELETE | /mes/item-request-detail/delete | 删除叫料申请明细 | id (Long) | Boolean |
| GET | /mes/item-request-detail/get | 获得叫料申请明细 | id (Long) | MesItemRequestDetailRespVO |
| GET | /mes/item-request-detail/page | 获得叫料申请明细分页 | MesItemRequestDetailPageReqVO | PageResult\<MesItemRequestDetailRespVO\> |
| POST | /mes/item-request-detail/senior | 高级搜索叫料申请明细 | CustomConditions | PageResult\<MesItemRequestDetailRespVO\> |
| GET | /mes/item-request-detail/export-excel | 导出叫料申请明细Excel | MesItemRequestDetailExportReqVO | Excel文件流 |
| GET | /mes/item-request-detail/get-import-template | 获得导入叫料申请明细模板 | - | Excel文件流 |

### 六、报废拆解

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/dismantling-main/create | 创建报废拆解登记主 | MesDismantlingMainCreateReqVO | Integer(id) |
| PUT | /mes/dismantling-main/update | 更新报废拆解登记主 | MesDismantlingMainUpdateReqVO | Boolean |
| DELETE | /mes/dismantling-main/delete | 删除报废拆解登记主 | id (Integer) | Boolean |
| GET | /mes/dismantling-main/get | 获得报废拆解登记主 | id (Integer) | MesDismantlingMainRespVO |
| GET | /mes/dismantling-main/page | 获得报废拆解登记主分页 | MesDismantlingMainPageReqVO | PageResult\<MesDismantlingMainRespVO\> |
| POST | /mes/dismantling-main/senior | 高级搜索报废拆解登记主 | CustomConditions | PageResult\<MesDismantlingMainRespVO\> |
| GET | /mes/dismantling-main/export-excel | 导出报废拆解登记主Excel | MesDismantlingMainExportReqVO | Excel文件流 |
| GET | /mes/dismantling-main/get-import-template | 获得导入报废拆解登记主模板 | - | Excel文件流 |
| POST | /mes/dismantling-detail/create | 创建报废拆解明细 | MesDismantlingDetailCreateReqVO | Integer(id) |
| PUT | /mes/dismantling-detail/update | 更新报废拆解明细 | MesDismantlingDetailUpdateReqVO | Boolean |
| DELETE | /mes/dismantling-detail/delete | 删除报废拆解明细 | id (Integer) | Boolean |
| GET | /mes/dismantling-detail/get | 获得报废拆解明细 | id (Integer) | MesDismantlingDetailRespVO |
| GET | /mes/dismantling-detail/page | 获得报废拆解明细分页 | MesDismantlingDetailPageReqVO | PageResult\<MesDismantlingDetailRespVO\> |
| POST | /mes/dismantling-detail/senior | 高级搜索报废拆解明细 | CustomConditions | PageResult\<MesDismantlingDetailRespVO\> |
| GET | /mes/dismantling-detail/export-excel | 导出报废拆解明细Excel | MesDismantlingDetailExportReqVO | Excel文件流 |
| GET | /mes/dismantling-detail/get-import-template | 获得导入报废拆解明细模板 | - | Excel文件流 |

### 七、报工管理

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/mes-job-report-log/create | 创建工单报工日志 | MesJobReportLogCreateReqVO | Long(id) |
| GET | /mes/mes-job-report-log/get | 获得工单报工日志 | id (Long) | MesJobReportLogRespVO |
| GET | /mes/mes-job-report-log/page | 获得工单报工日志分页 | MesJobReportLogPageReqVO | PageResult\<MesJobReportLogRespVO\> |
| POST | /mes/mes-job-report-log/senior | 高级搜索工单报工日志 | CustomConditions | PageResult\<MesJobReportLogRespVO\> |

### 八、生产日计划

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/orderday/create | 创建生产日计划 | MesOrderDayCreateReqVO | Long(id) |
| PUT | /mes/orderday/update | 更新生产日计划 | MesOrderDayUpdateReqVO | Boolean |
| DELETE | /mes/orderday/delete | 删除生产日计划 | id (Long) | Boolean |
| GET | /mes/orderday/get | 获得生产日计划 | id (Long) | MesOrderDayRespVO |
| GET | /mes/orderday/list | 获得生产日计划列表 | ids (Collection\<Long\>) | List\<MesOrderDayRespVO\> |
| GET | /mes/orderday/page | 获得生产日计划分页 | MesOrderDayPageReqVO | PageResult\<MesOrderDayRespVO\> |
| POST | /mes/orderday/publishPlan | 发布日生产计划 | MesOrderDayCreateReqVO | String |
| POST | /mes/orderday/stopPlan/{id} | 终止已发布计划 | id (Long, path) | String |
| GET | /mes/orderday/getProcessInfo | 获得工艺路由及其工序 | code (String) | JSONObject |
| GET | /mes/orderday/getBomInfo | 获得配置的BOM信息 | MesOrderDayBomExportReqVO | List\<MesOrderDayBomDO\> |
| GET | /mes/orderday/getDeviceInfo | 获取设备信息 | workCode, lineCode | List\<JSONObject\> |
| GET | /mes/orderday/getWorkGroup | 获取工作人员信息 | workCode, lineCode, planDate | List\<WorkerQueryVo\> |
| GET | /mes/orderday/export-excel | 导出生产日计划Excel | MesOrderDayExportReqVO | Excel文件流 |
| GET | /mes/orderday/get-import-template | 获得导入生产日计划模板 | - | Excel文件流 |
| POST | /mes/orderday/import | 导入生产日计划基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |

### 九、日计划BOM

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/orderDayBom/create | 创建日计划BOM信息 | MesOrderDayBomCreateReqVO | Long(id) |
| PUT | /mes/orderDayBom/update | 更新日计划BOM信息 | MesOrderDayBomUpdateReqVO | Boolean |
| DELETE | /mes/orderDayBom/delete | 删除日计划BOM信息 | id (Long) | Boolean |
| GET | /mes/orderDayBom/get | 获得日计划BOM信息 | id (Long) | MesOrderDayBomRespVO |
| GET | /mes/orderDayBom/getByOrder | 获得计划产品工艺路线配置 | OrderDayQueryParamVo | List\<MesOrderDayBomRespVO\> |
| GET | /mes/orderDayBom/list | 获得日计划BOM信息列表 | ids (Collection\<Long\>) | List\<MesOrderDayBomRespVO\> |
| GET | /mes/orderDayBom/page | 获得日计划BOM信息分页 | MesOrderDayBomPageReqVO | PageResult\<MesOrderDayBomRespVO\> |

### 十、计划工单设备/人员/工位

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/orderDayequipment/create | 创建计划工单设备配置 | MesOrderDayEquipmentCreateReqVO | Long(id) |
| POST | /mes/orderDayequipment/batchCreate | 批量创建计划工单设备配置 | List\<MesOrderDayEquipmentCreateReqVO\> | null |
| PUT | /mes/orderDayequipment/update | 更新计划工单设备配置 | MesOrderDayEquipmentUpdateReqVO | Boolean |
| DELETE | /mes/orderDayequipment/delete | 删除计划工单设备配置 | id (Long) | Boolean |
| GET | /mes/orderDayequipment/get | 获得计划工单设备配置 | id (Long) | MesOrderDayEquipmentRespVO |
| GET | /mes/orderDayequipment/page | 获得计划工单设备配置分页 | MesOrderDayEquipmentPageReqVO | PageResult\<MesOrderDayEquipmentRespVO\> |
| GET | /mes/orderDayequipment/getByOrder | 获得计划产品工艺路线的工序配置 | OrderDayQueryParamVo | List\<MesOrderDayEquipmentRespVO\> |
| PUT | /mes/orderDayRoute/update | 更新计划产品工艺路线配置 | MesOrderDayRouteUpdateReqVO | Boolean |
| GET | /mes/orderDayRoute/get | 获得计划产品工艺路线配置 | id (Long) | MesOrderDayRouteRespVO |
| GET | /mes/orderDayRoute/getByOrder | 获得计划产品工艺路线配置 | OrderDayQueryParamVo | MesOrderDayRouteRespVO |
| GET | /mes/orderDayRoute/page | 获得计划产品工艺路线配置分页 | MesOrderDayRoutePageReqVO | PageResult\<MesOrderDayRouteRespVO\> |
| POST | /mes/orderDayRoutesub/create | 创建计划工单产品工艺路线工序 | MesOrderDayRoutesubCreateReqVO | Long(id) |
| PUT | /mes/orderDayRoutesub/update | 更新计划工单产品工艺路线工序 | MesOrderDayRoutesubUpdateReqVO | Boolean |
| DELETE | /mes/orderDayRoutesub/delete | 删除计划工单产品工艺路线工序 | id (Long) | Boolean |
| GET | /mes/orderDayRoutesub/get | 获得计划工单产品工艺路线工序 | id (Long) | MesOrderDayRoutesubRespVO |
| GET | /mes/orderDayRoutesub/getByOrder | 获得计划产品工艺路线的工序配置 | OrderDayQueryParamVo | List\<ProcessRespVO\> |
| GET | /mes/orderDayRoutesub/page | 获得计划工单产品工艺路线工序分页 | MesOrderDayRoutesubPageReqVO | PageResult\<MesOrderDayRoutesubRespVO\> |
| POST | /mes/orderDayWorker/create | 创建计划工单人员配置 | MesOrderDayWorkerCreateReqVO | Long(id) |
| POST | /mes/orderDayWorker/batchCreate | 批量创建计划工单人员配置 | List\<MesOrderDayWorkerCreateReqVO\> | null |
| PUT | /mes/orderDayWorker/update | 更新计划工单人员配置 | MesOrderDayWorkerUpdateReqVO | Boolean |
| DELETE | /mes/orderDayWorker/delete | 删除计划工单人员配置 | id (Long) | Boolean |
| GET | /mes/orderDayWorker/get | 获得计划工单人员配置 | id (Long) | MesOrderDayWorkerRespVO |
| GET | /mes/orderDayWorker/getByOrder | 获得计划产品工艺路线的工序配置 | OrderDayQueryParamVo | List\<MesOrderDayWorkerRespVO\> |
| GET | /mes/orderDayWorker/page | 获得计划工单人员配置分页 | MesOrderDayWorkerPageReqVO | PageResult\<MesOrderDayWorkerRespVO\> |
| POST | /mes/orderDayWorkstation/create | 创建计划工单工位配置 | MesOrderDayWorkstationCreateReqVO | Long(id) |
| PUT | /mes/orderDayWorkstation/update | 更新计划工单工位配置 | MesOrderDayWorkstationUpdateReqVO | Boolean |
| DELETE | /mes/orderDayWorkstation/delete | 删除计划工单工位配置 | id (Long) | Boolean |
| GET | /mes/orderDayWorkstation/get | 获得计划工单工位配置 | id (Long) | MesOrderDayWorkstationRespVO |
| GET | /mes/orderDayWorkstation/getByOrder | 获得计划产品工艺路线的工序配置 | OrderDayQueryParamVo | List\<MesOrderDayWorkstationRespVO\> |
| GET | /mes/orderDayWorkstation/page | 获得计划工单工位配置分页 | MesOrderDayWorkstationPageReqVO | PageResult\<MesOrderDayWorkstationRespVO\> |
| GET | /mes/orderDayWorkstation/export-excel | 导出计划工单工位配置Excel | MesOrderDayWorkstationExportReqVO | Excel文件流 |
| GET | /mes/orderDayWorkstation/get-import-template | 获得导入计划工单工位配置模板 | - | Excel文件流 |

### 十一、计划工单操作流水

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/order-oper-log/create | 创建计划工单操作流水日志 | MesOrderOperLogCreateReqVO | Long(id) |
| PUT | /mes/order-oper-log/update | 更新计划工单操作流水日志 | MesOrderOperLogUpdateReqVO | Boolean |
| DELETE | /mes/order-oper-log/delete | 删除计划工单操作流水日志 | id (Long) | Boolean |
| GET | /mes/order-oper-log/get | 获得计划工单操作流水日志 | id (Long) | MesOrderOperLogRespVO |
| GET | /mes/order-oper-log/list | 获得计划工单操作流水日志列表 | ids (Collection\<Long\>) | List\<MesOrderOperLogRespVO\> |
| GET | /mes/order-oper-log/page | 获得计划工单操作流水日志分页 | MesOrderOperLogPageReqVO | PageResult\<MesOrderOperLogRespVO\> |
| GET | /mes/order-oper-log/export-excel | 导出计划工单操作流水日志Excel | MesOrderOperLogExportReqVO | Excel文件流 |
| GET | /mes/order-oper-log/get-import-template | 获得导入计划工单操作流水日志模板 | - | Excel文件流 |
| POST | /mes/order-oper-log/import | 导入计划工单操作流水日志基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |

### 十二、工序与资源关联

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/mes-process-itembasic/createRelation | 创建工序与物料关联 | ProcessItembasicRelationReqVO | Long(id) |
| POST | /mes/mes-process-itembasic/deleteRelation | 删除工位物料关联 | ProcessItembasicRelationReqVO | Object |
| GET | /mes/mes-process-itembasic/pageByProcessCode | 获得工序与物料关联分页 | MesProcessItembasicPageReqVO | PageResult\<ItembasicDTO\> |
| POST | /mes/mes-process-itembasic/create | 创建工序与物料关联 | MesProcessItembasicCreateReqVO | Integer(id) |
| PUT | /mes/mes-process-itembasic/update | 更新工序与物料关联 | MesProcessItembasicUpdateReqVO | Boolean |
| DELETE | /mes/mes-process-itembasic/delete | 删除工序与物料关联 | id (Integer) | Boolean |
| GET | /mes/mes-process-itembasic/get | 获得工序与物料关联 | id (Integer) | MesProcessItembasicRespVO |
| GET | /mes/mes-process-itembasic/page | 获得工序与物料关联分页 | MesProcessItembasicPageReqVO | PageResult\<MesProcessItembasicRespVO\> |
| POST | /mes/mes-process-itembasic/senior | 高级搜索工序与物料关联 | CustomConditions | PageResult\<MesProcessItembasicRespVO\> |
| GET | /mes/mes-process-itembasic/export-excel | 导出工序与物料关联Excel | MesProcessItembasicExportReqVO | Excel文件流 |
| GET | /mes/mes-process-itembasic/get-import-template | 获得导入工序与物料关联模板 | - | Excel文件流 |
| POST | /mes/mes-process-itembasic/import | 导入工序与物料关联基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |
| POST | /mes/mes-process-pattern/createRelation | 创建工序与模具关联 | ProcessPatternRelationReqVO | Long(id) |
| POST | /mes/mes-process-pattern/create | 创建工序与模具关联 | MesProcessPatternCreateReqVO | Long(id) |
| GET | /mes/mes-process-pattern/pagePatternByProcessCode | 获得工序与产线关联分页 | MesProcessPatternPageReqVO | PageResult\<ProductionlineDTO\> |
| POST | /mes/mes-process-pattern/deleteRelation | 删除工位设备关联 | ProcessPatternRelationReqVO | Object |
| PUT | /mes/mes-process-pattern/update | 更新工序与模具关联 | MesProcessPatternUpdateReqVO | Boolean |
| DELETE | /mes/mes-process-pattern/delete | 删除工序与模具关联 | id (Long) | Boolean |
| GET | /mes/mes-process-pattern/get | 获得工序与模具关联 | id (Long) | MesProcessPatternRespVO |
| GET | /mes/mes-process-pattern/page | 获得工序与模具关联分页 | MesProcessPatternPageReqVO | PageResult\<MesProcessPatternRespVO\> |
| POST | /mes/mes-process-pattern/senior | 高级搜索工序与模具关联 | CustomConditions | PageResult\<MesProcessPatternRespVO\> |
| GET | /mes/mes-process-pattern/export-excel | 导出工序与模具关联Excel | MesProcessPatternExportReqVO | Excel文件流 |
| GET | /mes/mes-process-pattern/get-import-template | 获得导入工序与模具关联模板 | - | Excel文件流 |
| POST | /mes/mes-process-pattern/import | 导入工序与模具关联基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |
| POST | /mes/mes-process-productionline/createRelation | 创建工序与产线关联 | ProcessProductionineRelationReqVO | Object |
| GET | /mes/mes-process-productionline/pageByProcessCode | 获得工序与产线关联分页 | MesProcessProductionlinePageReqVO | PageResult\<ProductionlineDTO\> |
| POST | /mes/mes-process-productionline/deleteRelation | 删除工位设备关联 | ProcessProductionineRelationReqVO | Object |
| POST | /mes/mes-process-productionline/create | 创建工序与产线关联 | MesProcessProductionlineCreateReqVO | Long(id) |
| PUT | /mes/mes-process-productionline/update | 更新工序与产线关联 | MesProcessProductionlineUpdateReqVO | Boolean |
| DELETE | /mes/mes-process-productionline/delete | 删除工序与产线关联 | id (Long) | Boolean |
| GET | /mes/mes-process-productionline/get | 获得工序与产线关联 | id (Long) | MesProcessProductionlineRespVO |
| GET | /mes/mes-process-productionline/page | 获得工序与产线关联分页 | MesProcessProductionlinePageReqVO | PageResult\<MesProcessProductionlineRespVO\> |
| POST | /mes/mes-process-productionline/senior | 高级搜索工序与产线关联 | CustomConditions | PageResult\<MesProcessProductionlineRespVO\> |
| GET | /mes/mes-process-productionline/export-excel | 导出工序与产线关联Excel | MesProcessProductionlineExportReqVO | Excel文件流 |
| GET | /mes/mes-process-productionline/get-import-template | 获得导入工序与产线关联模板 | - | Excel文件流 |
| POST | /mes/mes-process-productionline/import | 导入工序与产线关联基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |

### 十三、返工返修

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/product-offline/create | 创建产品离线登记记录 | MesProductOfflineCreateReqVO | Integer(id) |
| PUT | /mes/product-offline/update | 更新产品离线登记记录 | MesProductOfflineUpdateReqVO | Boolean |
| DELETE | /mes/product-offline/delete | 删除产品离线登记记录 | id (Integer) | Boolean |
| GET | /mes/product-offline/get | 获得产品离线登记记录 | id (Integer) | MesProductOfflineRespVO |
| GET | /mes/product-offline/list | 获得产品离线登记记录列表 | ids (Collection\<Integer\>) | List\<MesProductOfflineRespVO\> |
| GET | /mes/product-offline/page | 获得产品离线登记记录分页 | MesProductOfflinePageReqVO | PageResult\<MesProductOfflineRespVO\> |
| POST | /mes/product-offline/senior | 高级搜索产品离线登记记录 | CustomConditions | PageResult\<MesProductOfflineRespVO\> |
| GET | /mes/product-offline/export-excel | 导出产品离线登记记录Excel | MesProductOfflineExportReqVO | Excel文件流 |
| GET | /mes/product-offline/get-import-template | 获得导入产品离线登记记录模板 | - | Excel文件流 |
| POST | /mes/product-offline/import | 导入产品离线登记记录基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |
| POST | /mes/product-backline/create | 创建产品返线登记 | MesProductBacklineCreateReqVO | Integer(id) |
| PUT | /mes/product-backline/update | 更新产品返线登记 | MesProductBacklineUpdateReqVO | Boolean |
| DELETE | /mes/product-backline/delete | 删除产品返线登记 | id (Integer) | Boolean |
| GET | /mes/product-backline/get | 获得产品返线登记 | id (Integer) | MesProductBacklineRespVO |
| GET | /mes/product-backline/list | 获得产品返线登记列表 | ids (Collection\<Integer\>) | List\<MesProductBacklineRespVO\> |
| GET | /mes/product-backline/page | 获得产品返线登记分页 | MesProductBacklinePageReqVO | PageResult\<MesProductBacklineRespVO\> |
| POST | /mes/product-backline/senior | 高级搜索产品返线登记 | CustomConditions | PageResult\<MesProductBacklineRespVO\> |
| GET | /mes/product-backline/export-excel | 导出产品返线登记Excel | MesProductBacklineExportReqVO | Excel文件流 |
| GET | /mes/product-backline/get-import-template | 获得导入产品返线登记模板 | - | Excel文件流 |
| POST | /mes/product-backline/import | 导入产品返线登记基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |
| POST | /mes/reportfinish-store/create | 创建完工库存中间 | MesReportfinishStoreCreateReqVO | Long(id) |
| PUT | /mes/reportfinish-store/update | 更新完工库存中间 | MesReportfinishStoreUpdateReqVO | Boolean |
| DELETE | /mes/reportfinish-store/delete | 删除完工库存中间 | id (Long) | Boolean |
| GET | /mes/reportfinish-store/get | 获得完工库存中间 | id (Long) | MesReportfinishStoreRespVO |
| GET | /mes/reportfinish-store/page | 获得完工库存中间分页 | MesReportfinishStorePageReqVO | PageResult\<MesReportfinishStoreRespVO\> |
| POST | /mes/reportfinish-store/senior | 高级搜索完工库存中间 | CustomConditions | PageResult\<MesReportfinishStoreRespVO\> |
| GET | /mes/reportfinish-store/export-excel | 导出完工库存中间Excel | MesReportfinishStoreExportReqVO | Excel文件流 |
| GET | /mes/reportfinish-store/get-import-template | 获得导入完工库存中间模板 | - | Excel文件流 |
| POST | /mes/reportfinish-store/import | 导入完工库存中间基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |
| POST | /mes/reportp-store/create | 创建工序报工物料明细 | MesReportpStoreCreateReqVO | Long(id) |
| PUT | /mes/reportp-store/update | 更新工序报工物料明细 | MesReportpStoreUpdateReqVO | Boolean |
| DELETE | /mes/reportp-store/delete | 删除工序报工物料明细 | id (Long) | Boolean |
| GET | /mes/reportp-store/get | 获得工序报工物料明细 | id (Long) | MesReportpStoreRespVO |
| GET | /mes/reportp-store/page | 获得工序报工物料明细分页 | MesReportpStorePageReqVO | PageResult\<MesReportpStoreRespVO\> |
| POST | /mes/reportp-store/senior | 高级搜索工序报工物料明细 | CustomConditions | PageResult\<MesReportpStoreRespVO\> |
| GET | /mes/reportp-store/export-excel | 导出工序报工物料明细Excel | MesReportpStoreExportReqVO | Excel文件流 |
| POST | /mes/rework-batch/create | 创建返工登记批量 | MesReworkBatchCreateReqVO | Integer(id) |
| PUT | /mes/rework-batch/update | 更新返工登记批量 | MesReworkBatchUpdateReqVO | Boolean |
| DELETE | /mes/rework-batch/delete | 删除返工登记批量 | id (Integer) | Boolean |
| DELETE | /mes/rework-batch/suspend | 中止返工登记批量任务 | id (Integer) | Boolean |
| PUT | /mes/rework-batch/receive | 领取返工登记批量任务 | id (Integer) | Boolean |
| PUT | /mes/rework-batch/finish | 完成返工登记批量任务 | id (Integer) | Boolean |
| GET | /mes/rework-batch/get | 获得返工登记批量 | id (Integer) | MesReworkBatchRespVO |
| GET | /mes/rework-batch/list | 获得返工登记批量列表 | ids (Collection\<Integer\>) | List\<MesReworkBatchRespVO\> |
| GET | /mes/rework-batch/page | 获得返工登记批量分页 | MesReworkBatchPageReqVO | PageResult\<MesReworkBatchRespVO\> |
| POST | /mes/rework-batch/senior | 高级搜索返工登记批量 | CustomConditions | PageResult\<MesReworkBatchRespVO\> |
| GET | /mes/rework-batch/export-excel | 导出返工返修登记Excel | MesReworkBatchExportReqVO | Excel文件流 |
| GET | /mes/rework-batch/get-import-template | 获得导入返工返修登记批量模板 | - | Excel文件流 |
| POST | /mes/rework-single/create | 创建返工登记单件 | MesReworkSingleCreateReqVO | Integer(id) |
| PUT | /mes/rework-single/update | 更新返工登记单件 | MesReworkSingleUpdateReqVO | Boolean |
| DELETE | /mes/rework-single/delete | 删除返工登记单件 | id (Integer) | Boolean |
| DELETE | /mes/rework-single/suspend | 中止返工登记单件任务 | id (Integer) | Boolean |
| PUT | /mes/rework-single/receive | 领取返工登记单件任务 | id (Integer) | Boolean |
| PUT | /mes/rework-single/finish | 完成返工登记单件任务 | id (Integer) | Boolean |
| GET | /mes/rework-single/get | 获得返工登记单件 | id (Integer) | MesReworkSingleRespVO |
| GET | /mes/rework-single/list | 获得返工登记单件列表 | ids (Collection\<Integer\>) | List\<MesReworkSingleRespVO\> |
| GET | /mes/rework-single/page | 获得返工登记单件分页 | MesReworkSinglePageReqVO | PageResult\<MesReworkSingleRespVO\> |
| POST | /mes/rework-single/senior | 高级搜索返工登记单件 | CustomConditions | PageResult\<MesReworkSingleRespVO\> |
| GET | /mes/rework-single/export-excel | 导出返工登记单件Excel | MesReworkSingleExportReqVO | Excel文件流 |
| GET | /mes/rework-single/get-import-template | 获得导入返工登记单件模板 | - | Excel文件流 |

### 十四、生产任务排产

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/workScheduling/create | 创建生产任务排产 | MesWorkSchedulingCreateReqVO | Integer(id) |
| PUT | /mes/workScheduling/update | 更新生产任务排产 | MesWorkSchedulingUpdateReqVO | Boolean |
| POST | /mes/workScheduling/update-status | 更新生产任务排产状态 | MesWorkSchedulingUpdateReqVO | Boolean |
| PUT | /mes/workScheduling/updateStatus | 更新生产任务状态 | JSONObject | Boolean |
| DELETE | /mes/workScheduling/delete | 删除生产任务排产 | id (Integer) | Boolean |
| GET | /mes/workScheduling/get | 获得生产任务排产 | id (Integer) | MesWorkSchedulingRespVO |
| GET | /mes/workScheduling/get-PDF | SOP-获取PDF | planMasterCode (String) | List\<FileDO\> |
| GET | /mes/workScheduling/getNodeInfo | 获取对应工单的当前工序基本信息 | planDayCode, processCode | MesOrderDayRoutesubRespVO |
| GET | /mes/workScheduling/getCurrentWorkerList | 获取对应工单的当前工序人员列表 | planDayCode, processCode | List\<MesOrderDayWorkerRespVO\> |
| GET | /mes/workScheduling/getProcessList | 获取对应工单的所有工序信息 | planDayCode, schedulingCode | List\<MesOrderDayRoutesubRespVO\> |
| GET | /mes/workScheduling/getProcessListByStaus | 获取对应工单指定状态的工序信息 | planDayCode, schedulingCode, status | List\<MesOrderDayRoutesubRespVO\> |
| GET | /mes/workScheduling/getNodePosition | 获取当前工序的位置 | planDayCode, processCode | String |
| GET | /mes/workScheduling/list | 获得生产任务排产列表 | ids (Collection\<Integer\>) | List\<MesWorkSchedulingRespVO\> |
| GET | /mes/workScheduling/page | 获得生产任务排产分页 | MesWorkSchedulingPageReqVO | PageResult\<MesWorkSchedulingRespVO\> |
| GET | /mes/workScheduling/PDA-page | 获得生产任务排产分页(PDA) | MesWorkSchedulingPageReqVO | PageResult\<JSONObject\> |
| POST | /mes/workScheduling/senior | 高级搜索生产任务排产 | CustomConditions | PageResult\<MesWorkSchedulingRespVO\> |
| POST | /mes/workScheduling/completeHandle | 完工处理 | MesWorkSchedulingUpdateReqVO | Integer |
| POST | /mes/workScheduling/reportForAll | 批量报工处理 | MesReportBatchVo | Integer |

### 十五、工单任务明细

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/work-scheduling-detail/create | 创建工单任务明细 | MesWorkSchedulingDetailCreateReqVO | Long(id) |
| PUT | /mes/work-scheduling-detail/update | 更新工单任务明细 | MesWorkSchedulingDetailUpdateReqVO | Boolean |
| DELETE | /mes/work-scheduling-detail/delete | 删除工单任务明细 | id (Long) | Boolean |
| GET | /mes/work-scheduling-detail/get | 获得工单任务明细 | id (Long) | MesWorkSchedulingDetailRespVO |
| POST | /mes/work-scheduling-detail/get-info | 获得工单任务明细 | MesWorkSchedulingDetailBaseVO | List\<MesWorkSchedulingDetailRespVO\> |
| GET | /mes/work-scheduling-detail/page | 获得工单任务明细分页 | MesWorkSchedulingDetailPageReqVO | PageResult\<MesWorkSchedulingDetailRespVO\> |
| POST | /mes/work-scheduling-detail/senior | 高级搜索工单任务明细 | CustomConditions | PageResult\<MesWorkSchedulingDetailRespVO\> |
| GET | /mes/work-scheduling-detail/export-excel | 导出工单任务明细Excel | MesWorkSchedulingDetailExportReqVO | Excel文件流 |
| GET | /mes/work-scheduling-detail/get-import-template | 获得导入工单任务明细模板 | - | Excel文件流 |
| POST | /mes/work-scheduling-detail/reportWorkByProcess | 工序报工 | ReportWorkByProcessReqVO | Boolean |
| GET | /mes/work-scheduling-detail/getPeopleReportList | 获取人员报工列表 | GetPeopleReportByOrderReqVO | GetPeopleReportByOrderResVO |
| GET | /mes/work-scheduling-detail/processFinished | 工序完工 | id (Long) | Integer |
| POST | /mes/work-scheduling-detail/processQualified | 工序质检 | ProcessQualifiedVo | Integer |

### 十六、任务质检单

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/work-scheduling-qaform/create | 创建任务质检单 | MesWorkSchedulingQaformCreateReqVO | Integer(id) |
| PUT | /mes/work-scheduling-qaform/update | 更新任务质检单 | MesWorkSchedulingQaformUpdateReqVO | Boolean |
| DELETE | /mes/work-scheduling-qaform/delete | 删除任务质检单 | id (Integer) | Boolean |
| GET | /mes/work-scheduling-qaform/get | 获得任务质检单 | id (Integer) | MesWorkSchedulingQaformRespVO |
| GET | /mes/work-scheduling-qaform/page | 获得任务质检单分页 | MesWorkSchedulingQaformPageReqVO | PageResult\<MesWorkSchedulingQaformRespVO\> |
| POST | /mes/work-scheduling-qaform/senior | 高级搜索任务质检单 | CustomConditions | PageResult\<MesWorkSchedulingQaformRespVO\> |
| GET | /mes/work-scheduling-qaform/export-excel | 导出任务质检单Excel | MesWorkSchedulingQaformExportReqVO | Excel文件流 |
| GET | /mes/work-scheduling-qaform/get-import-template | 获得导入任务质检单模板 | - | Excel文件流 |

### 十七、班组管理

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/teamSetting/create | 创建班组人员管理 | MesTeamSettingCreateReqVO | Long(id) |
| PUT | /mes/teamSetting/update | 更新班组人员管理 | MesTeamSettingUpdateReqVO | Boolean |
| DELETE | /mes/teamSetting/delete | 删除班组人员管理 | id (Long) | Boolean |
| GET | /mes/teamSetting/get | 获得班组人员管理 | id (Long) | MesTeamSettingRespVO |
| GET | /mes/teamSetting/list | 获得班组人员管理列表 | ids (Collection\<Long\>) | List\<MesTeamSettingRespVO\> |
| GET | /mes/teamSetting/page | 获得班组人员管理分页 | MesTeamSettingPageReqVO | PageResult\<MesTeamSettingRespVO\> |
| GET | /mes/teamSetting/export-excel | 导出班组人员管理Excel | MesTeamSettingExportReqVO | Excel文件流 |
| GET | /mes/teamSetting/get-import-template | 获得导入班组人员管理模板 | - | Excel文件流 |
| POST | /mes/teamSetting/import | 导入班组人员管理基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |

### 十八、MES操作步骤

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/opersteps/create | 创建MES操作步骤信息 | OperstepsCreateReqVO | Integer(id) |
| PUT | /mes/opersteps/update | 更新MES操作步骤信息 | OperstepsUpdateReqVO | Boolean |
| DELETE | /mes/opersteps/delete | 删除MES操作步骤信息 | id (Integer) | Boolean |
| GET | /mes/opersteps/get | 获得MES操作步骤信息 | id (Integer) | OperstepsRespVO |
| GET | /mes/opersteps/list | 获得MES操作步骤信息列表 | ids (Collection\<Integer\>) | List\<OperstepsRespVO\> |
| GET | /mes/opersteps/page | 获得MES操作步骤信息分页 | OperstepsPageReqVO | PageResult\<OperstepsRespVO\> |
| POST | /mes/opersteps/senior | 高级搜索操作步骤类型配置 | CustomConditions | PageResult\<OperstepsRespVO\> |
| GET | /mes/opersteps/export-excel | 导出MES操作步骤信息Excel | OperstepsExportReqVO | Excel文件流 |
| GET | /mes/opersteps/get-import-template | 获得导入MES操作步骤信息模板 | - | Excel文件流 |
| POST | /mes/opersteps/import | 导入MES操作步骤信息基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |
| POST | /mes/opersteps-type/create | 创建操作步骤类型配置 | OperstepsTypeCreateReqVO | Integer(id) |
| PUT | /mes/opersteps-type/update | 更新操作步骤类型配置 | OperstepsTypeUpdateReqVO | Boolean |
| DELETE | /mes/opersteps-type/delete | 删除操作步骤类型配置 | id (Integer) | Boolean |
| GET | /mes/opersteps-type/get | 获得操作步骤类型配置 | id (Integer) | OperstepsTypeRespVO |
| GET | /mes/opersteps-type/list | 获得操作步骤类型配置列表 | ids (Collection\<Integer\>) | List\<OperstepsTypeRespVO\> |
| GET | /mes/opersteps-type/page | 获得操作步骤类型配置分页 | OperstepsTypePageReqVO | PageResult\<OperstepsTypeRespVO\> |
| POST | /mes/opersteps-type/senior | 高级搜索操作步骤类型配置 | CustomConditions | PageResult\<OperstepsTypeRespVO\> |
| GET | /mes/opersteps-type/export-excel | 导出操作步骤类型配置Excel | OperstepsTypeExportReqVO | Excel文件流 |
| GET | /mes/opersteps-type/get-import-template | 获得导入操作步骤类型配置模板 | - | Excel文件流 |
| POST | /mes/opersteps-type/import | 导入操作步骤类型配置基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |

### 十九、订单月计划

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /plan/mes-order-month-main/create | 创建订单计划月主 | MesOrderMonthMainCreateReqVO | Integer(id) |
| PUT | /plan/mes-order-month-main/update | 更新订单计划月主 | MesOrderMonthMainUpdateReqVO | Boolean |
| DELETE | /plan/mes-order-month-main/delete | 删除订单计划月主 | id (Integer) | Boolean |
| GET | /plan/mes-order-month-main/get | 获得订单计划月主 | id (Integer) | MesOrderMonthMainRespVO |
| GET | /plan/mes-order-month-main/list | 获得订单计划月主列表 | ids (Collection\<Integer\>) | List\<MesOrderMonthMainRespVO\> |
| GET | /plan/mes-order-month-main/page | 获得订单计划月主分页 | MesOrderMonthMainPageReqVO | PageResult\<MesOrderMonthMainRespVO\> |
| GET | /plan/mes-order-month-main/export-excel | 导出订单计划Excel | MesOrderMonthMainExportReqVO | Excel文件流 |
| GET | /plan/mes-order-month-main/get-import-template | 获得导入订单计划模板 | - | Excel文件流 |
| POST | /plan/mes-order-month-main/import | 导入订单计划基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |
| POST | /plan/mes-order-month-sub/create | 创建订单月计划子 | MesOrderMonthSubCreateReqVO | Integer(id) |
| PUT | /plan/mes-order-month-sub/update | 更新订单月计划子 | MesOrderMonthSubUpdateReqVO | Boolean |
| DELETE | /plan/mes-order-month-sub/delete | 删除订单月计划子 | id (Integer) | Boolean |
| GET | /plan/mes-order-month-sub/get | 获得订单月计划子 | id (Integer) | MesOrderMonthSubRespVO |
| GET | /plan/mes-order-month-sub/list | 获得订单月计划子列表 | ids (Collection\<Integer\>) | List\<MesOrderMonthSubRespVO\> |
| GET | /plan/mes-order-month-sub/page | 获得订单月计划子分页 | MesOrderMonthSubPageReqVO | PageResult\<MesOrderMonthSubRespVO\> |
| POST | /plan/mes-order-month-sub/breakdown | 拆解为日计划 | MesOrderMonthSubBreakdownReqVO | null |

### 二十、工艺路线

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/process/create | 创建工序 | MesProcessCreateReqVO | Long(id) |
| PUT | /mes/process/update | 更新工序 | MesProcessUpdateReqVO | Boolean |
| DELETE | /mes/process/delete | 删除工序 | id (Long) | Boolean |
| GET | /mes/process/get | 获得工序 | id (Long) | MesProcessRespVO |
| GET | /mes/process/getByCode | 获得工序基本信息 | code (String) | MesProcessRespVO |
| GET | /mes/process/page | 获得工序分页 | MesProcessPageReqVO | PageResult\<MesProcessRespVO\> |
| POST | /mes/process/senior | 高级搜索工序 | CustomConditions | PageResult\<MesProcessRespVO\> |
| GET | /mes/process/export-excel | 导出工序Excel | MesProcessExportReqVO | Excel文件流 |
| POST | /mes/process/export-excel-senior | 高级搜索导出工序Excel | CustomConditions | Excel文件流 |
| GET | /mes/process/get-import-template | 获得导入工序模板 | - | Excel文件流 |
| POST | /mes/process/import | 导入工序基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |
| POST | /mes/processroute/create | 创建工艺路线定义 | ProcessrouteCreateReqVO | Long(id) |
| PUT | /mes/processroute/update | 更新工艺路线定义 | ProcessrouteUpdateReqVO | Boolean |
| POST | /mes/processroute/updategraph | 更新工艺路线的图形定义 | ProcessrouteUpdateGraphReqVO | Boolean |
| DELETE | /mes/processroute/delete | 删除工艺路线定义 | id (Long) | Boolean |
| GET | /mes/processroute/get | 获得工艺路线定义 | id (Long) | ProcessrouteRespVO |
| GET | /mes/processroute/list | 获得工艺路线定义列表 | ids (Collection\<Long\>) | List\<ProcessrouteRespVO\> |
| GET | /mes/processroute/page | 获得工艺路线定义分页 | ProcessroutePageReqVO | PageResult\<ProcessrouteRespVO\> |
| POST | /mes/processroute/senior | 高级搜索工艺路线定义 | CustomConditions | PageResult\<ProcessrouteRespVO\> |
| GET | /mes/processroute/export-excel | 导出工艺路线定义Excel | ProcessrouteExportReqVO | Excel文件流 |
| GET | /mes/processroute/get-import-template | 获得导入工艺路线定义模板 | - | Excel文件流 |
| POST | /mes/processroute/import | 导入工艺路线定义基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |
| POST | /mes/processrouteNodeDetail/create | 创建工艺路由工序节点配置明细 | ProcessrouteNodeDetailCreateReqVO | Long(id) |
| PUT | /mes/processrouteNodeDetail/update | 更新工艺路由工序节点配置明细 | ProcessrouteNodeDetailUpdateReqVO | Boolean |
| DELETE | /mes/processrouteNodeDetail/delete | 删除工艺路由工序节点配置明细 | id (Long) | Boolean |
| GET | /mes/processrouteNodeDetail/get | 获得工艺路由工序节点配置明细 | id (Long) | ProcessrouteNodeDetailRespVO |
| GET | /mes/processrouteNodeDetail/list | 获得工艺路由工序节点配置明细列表 | ids (Collection\<Long\>) | List\<ProcessrouteNodeDetailRespVO\> |
| GET | /mes/processrouteNodeDetail/page | 获得工艺路由工序节点配置明细分页 | ProcessrouteNodeDetailPageReqVO | PageResult\<ProcessrouteNodeDetailRespVO\> |
| GET | /mes/processrouteNodeDetail/export-excel | 导出工艺路由工序节点配置明细Excel | ProcessrouteNodeDetailExportReqVO | Excel文件流 |
| GET | /mes/processrouteNodeDetail/get-import-template | 获得导入工艺路由工序节点配置明细模板 | - | Excel文件流 |
| POST | /mes/processrouteNodeDetail/import | 导入工艺路由工序节点配置明细基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |
| GET | /mes/processrouteNodeDetail/getRouteNodes | 获得计划产品工艺路线已配置的工序列表 | code (String) | List\<ProcessRespVO\> |

### 二十一、模具管理

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/pattern/create | 创建模具基本信息 | PatternCreateReqVO | Integer(id) |
| PUT | /mes/pattern/update | 更新模具基本信息 | PatternUpdateReqVO | Boolean |
| DELETE | /mes/pattern/delete | 删除模具基本信息 | id (Integer) | Boolean |
| GET | /mes/pattern/get | 获得模具基本信息 | id (Integer) | PatternRespVO |
| GET | /mes/pattern/page | 获得模具基本信息分页 | PatternPageReqVO | PageResult\<PatternRespVO\> |
| POST | /mes/pattern/senior | 高级搜索模具基本信息 | CustomConditions | PageResult\<PatternRespVO\> |
| GET | /mes/pattern/export-excel | 导出模具基本信息Excel | PatternExportReqVO | Excel文件流 |
| GET | /mes/pattern/get-import-template | 获得导入模具基本信息模板 | - | Excel文件流 |
| POST | /mes/pattern/import | 导入模具基本信息基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |
| POST | /mes/pattern-type/create | 创建模具类型 | PatternTypeCreateReqVO | Integer(id) |
| PUT | /mes/pattern-type/update | 更新模具类型 | PatternTypeUpdateReqVO | Boolean |
| DELETE | /mes/pattern-type/delete | 删除模具类型 | id (Integer) | Boolean |
| GET | /mes/pattern-type/get | 获得模具类型 | id (Integer) | PatternTypeRespVO |
| GET | /mes/pattern-type/page | 获得模具类型分页 | PatternTypePageReqVO | PageResult\<PatternTypeRespVO\> |
| POST | /mes/pattern-type/senior | 高级搜索模具类型 | CustomConditions | PageResult\<PatternTypeRespVO\> |
| GET | /mes/pattern-type/export-excel | 导出模具类型Excel | PatternTypeExportReqVO | Excel文件流 |
| GET | /mes/pattern-type/get-import-template | 获得导入模具类型模板 | - | Excel文件流 |
| POST | /mes/pattern-type/import | 导入模具类型基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |

### 二十二、质检管理

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/qualityclass/create | 创建质检类别 | QmsQualityclassCreateReqVO | Integer(id) |
| PUT | /mes/qualityclass/update | 更新质检类别 | QmsQualityclassUpdateReqVO | Boolean |
| DELETE | /mes/qualityclass/delete | 删除质检类别 | id (Integer) | Boolean |
| GET | /mes/qualityclass/get | 获得质检类别 | id (Integer) | QmsQualityclassRespVO |
| GET | /mes/qualityclass/list | 获得质检类别列表 | ids (Collection\<Integer\>) | List\<QmsQualityclassRespVO\> |
| GET | /mes/qualityclass/page | 获得质检类别分页 | QmsQualityclassPageReqVO | PageResult\<QmsQualityclassRespVO\> |
| POST | /mes/qualityclass/senior | 高级搜索质检类别信息 | CustomConditions | PageResult\<QmsQualityclassRespVO\> |
| GET | /mes/qualityclass/export-excel | 导出质检类别Excel | QmsQualityclassExportReqVO | Excel文件流 |
| GET | /mes/qualityclass/get-import-template | 获得导入质检类别模板 | - | Excel文件流 |
| POST | /mes/qualityclass/import | 导入质检类别基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |
| POST | /mes/qualitygroup/create | 创建质检分组 | QmsQualitygroupCreateReqVO | Integer(id) |
| PUT | /mes/qualitygroup/update | 更新质检分组 | QmsQualitygroupUpdateReqVO | Boolean |
| DELETE | /mes/qualitygroup/delete | 删除质检分组 | id (Integer) | Boolean |
| GET | /mes/qualitygroup/get | 获得质检分组 | id (Integer) | QmsQualitygroupRespVO |
| GET | /mes/qualitygroup/list | 获得质检分组列表 | ids (Collection\<Integer\>) | List\<QmsQualitygroupRespVO\> |
| GET | /mes/qualitygroup/page | 获得质检分组分页 | QmsQualitygroupPageReqVO | PageResult\<QmsQualitygroupRespVO\> |
| POST | /mes/qualitygroup/senior | 高级搜索质检分组信息 | CustomConditions | PageResult\<QmsQualitygroupRespVO\> |
| GET | /mes/qualitygroup/export-excel | 导出质检分组Excel | QmsQualitygroupExportReqVO | Excel文件流 |
| GET | /mes/qualitygroup/get-import-template | 获得导入质检分组模板 | - | Excel文件流 |
| POST | /mes/qualitygroup/import | 导入质检分组基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |
| POST | /mes/item/create | 创建质检项目定义 | QmsItemCreateReqVO | Integer(id) |
| PUT | /mes/item/update | 更新质检项目定义 | QmsItemUpdateReqVO | Boolean |
| DELETE | /mes/item/delete | 删除质检项目定义 | id (Integer) | Boolean |
| GET | /mes/item/get | 获得质检项目定义 | id (Integer) | QmsItemRespVO |
| GET | /mes/item/list | 获得质检项目定义列表 | ids (Collection\<Integer\>) | List\<QmsItemRespVO\> |
| GET | /mes/item/page | 获得质检项目定义分页 | QmsItemPageReqVO | PageResult\<QmsItemRespVO\> |
| POST | /mes/item/senior | 高级搜索质检项目信息 | CustomConditions | PageResult\<QmsItemRespVO\> |
| POST | /mes/item/senior-add | 高级搜索添加质检项目信息 | CustomConditions | PageResult\<QmsItemRespVO\> |
| GET | /mes/item/export-excel | 导出质检项目定义Excel | QmsItemExportReqVO | Excel文件流 |
| GET | /mes/item/get-import-template | 获得导入质检项目定义模板 | - | Excel文件流 |
| POST | /mes/item/import | 导入质检项目定义基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |
| POST | /mes/qualityform/create | 创建质检表单 | QmsQualityformCreateReqVO | Integer(id) |
| PUT | /mes/qualityform/update | 更新质检表单 | QmsQualityformUpdateReqVO | Boolean |
| DELETE | /mes/qualityform/delete | 删除质检表单 | id (Integer) | Boolean |
| GET | /mes/qualityform/get | 获得质检表单 | id (Integer) | QmsQualityformRespVO |
| GET | /mes/qualityform/getQualityform | 根据单号获得质检表单 | fromNo (String) | QmsQualityformRespVO |
| GET | /mes/qualityform/list | 获得质检表单列表 | ids (Collection\<Integer\>) | List\<QmsQualityformRespVO\> |
| GET | /mes/qualityform/page | 获得质检表单分页 | QmsQualityformPageReqVO | PageResult\<QmsQualityformRespVO\> |
| POST | /mes/qualityform/senior | 高级搜索质检表单信息 | CustomConditions | PageResult\<QmsQualityformRespVO\> |
| GET | /mes/qualityform/export-excel | 导出质检表单Excel | QmsQualityformExportReqVO | Excel文件流 |
| POST | /mes/qualityformdetail/create | 创建质检表单子表 | QmsQualityformdetailCreateReqVO | Integer(id) |
| PUT | /mes/qualityformdetail/update | 更新质检表单子表 | QmsQualityformdetailUpdateReqVO | Boolean |
| DELETE | /mes/qualityformdetail/delete | 删除质检表单子表 | id (Integer) | Boolean |
| GET | /mes/qualityformdetail/get | 获得质检表单子表 | id (Integer) | QmsQualityformdetailRespVO |
| GET | /mes/qualityformdetail/list | 获得质检表单子表列表 | ids (Collection\<Integer\>) | List\<QmsQualityformdetailRespVO\> |
| GET | /mes/qualityformdetail/page | 获得质检表单子表分页 | QmsQualityformdetailPageReqVO | PageResult\<QmsQualityformdetailRespVO\> |
| POST | /mes/qualityformdetail/senior | 高级搜索质检表单子表信息 | CustomConditions | PageResult\<QmsQualityformdetailRespVO\> |
| GET | /mes/qualityformdetail/export-excel | 导出质检表单子表Excel | QmsQualityformdetailExportReqVO | Excel文件流 |
| PUT | /mes/qualityformlog/update | 更新质检表单日志 | QmsQualityformlogUpdateReqVO | Boolean |
| DELETE | /mes/qualityformlog/delete | 删除质检表单日志 | id (Integer) | Boolean |
| GET | /mes/qualityformlog/get | 获得质检表单日志 | id (Integer) | QmsQualityformlogRespVO |
| GET | /mes/qualityformlog/list | 获得质检表单日志列表 | ids (Collection\<Integer\>) | List\<QmsQualityformlogRespVO\> |
| GET | /mes/qualityformlog/page | 获得质检表单日志分页 | QmsQualityformlogPageReqVO | PageResult\<QmsQualityformlogRespVO\> |
| POST | /mes/qualityformlog/senior | 高级搜索质检表单日志信息 | CustomConditions | PageResult\<QmsQualityformlogRespVO\> |
| GET | /mes/qualityformlog/export-excel | 导出质检表单日志Excel | QmsQualityformlogExportReqVO | Excel文件流 |

### 二十三、生产日历

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/schedulingcalendar/create | 创建生产日历 | SchedulingcalendarCreateReqVO | Integer(id) |
| POST | /mes/schedulingcalendar/createBatch | 批量创建生产日历 | List\<SchedulingcalendarCreateReqVO\> | Integer(0) |
| POST | /mes/schedulingcalendar/createObj | 根据json对象创建工作日历 | JSONObject | Integer(0) |
| PUT | /mes/schedulingcalendar/update | 更新生产日历 | SchedulingcalendarUpdateReqVO | Boolean |
| DELETE | /mes/schedulingcalendar/delete | 删除生产日历 | id (Integer) | Boolean |
| POST | /mes/schedulingcalendar/deleteTeam | 删除班组日历 | JSONObject | Integer |
| GET | /mes/schedulingcalendar/get | 获得生产日历 | id (Integer) | SchedulingcalendarRespVO |
| GET | /mes/schedulingcalendar/getList | 获得班组的生产日历(按keyDate分组) | code, startTime, endTime | Map\<String,List\<SchedulingcalendarDO\>\> |
| GET | /mes/schedulingcalendar/page | 获得生产日历分页 | SchedulingcalendarPageReqVO | PageResult\<SchedulingcalendarRespVO\> |
| POST | /mes/schedulingcalendar/senior | 高级搜索生产日历 | CustomConditions | PageResult\<SchedulingcalendarRespVO\> |
| GET | /mes/schedulingcalendar/export-excel | 导出生产日历Excel | SchedulingcalendarExportReqVO | Excel文件流 |
| GET | /mes/schedulingcalendar/get-import-template | 获得导入生产日历模板 | - | Excel文件流 |
| POST | /mes/schedulingcalendar/import | 导入生产日历基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |

### 二十四、工位管理

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/workstation/create | 创建工位 | MesWorkstationCreateReqVO | Long(id) |
| PUT | /mes/workstation/update | 更新工位 | MesWorkstationUpdateReqVO | Boolean |
| DELETE | /mes/workstation/delete | 删除工位 | id (Long) | Boolean |
| GET | /mes/workstation/get | 获得工位 | id (Long) | MesWorkstationRespVO |
| GET | /mes/workstation/list | 获得工位列表 | ids (Collection\<Long\>) | List\<MesWorkstationRespVO\> |
| GET | /mes/workstation/page | 获得工位分页 | MesWorkstationPageReqVO | PageResult\<MesWorkstationRespVO\> |
| POST | /mes/workstation-ability/create | 创建工位能力关联 | WorkstationAbilityRelationReqVO | Void |
| POST | /mes/workstation-ability/delete | 删除工位能力关联 | WorkstationAbilityRelationReqVO | Boolean |
| GET | /mes/workstation-ability/getWorkstationAbilityPage | 获得工位能力关联分页 | WorkstationAbilityPageReqVO | PageResult\<WorkstationAbilityRespVO\> |
| POST | /mes/workstation-ability/senior | 高级搜索工位能力关联 | CustomConditions | PageResult\<WorkstationAbilityRespVO\> |
| GET | /mes/equipment/getEquipmentPage | 获得工位设备关联分页 | WorkstationEquipmentPageReqVO | PageResult\<WorkstationEquipmentRespVO\> |
| GET | /mes/equipment/pageCheckList | 根据工位code获得可关联设备列表 | WorkstationEquipmentPageReqVO | PageResult\<WorkstationEquipmentRespVO\> |
| POST | /mes/equipment/createRelation | 关联工位设备 | WorkstationEquipmentRelationReqVO | Object |
| POST | /mes/equipment/deleteRelation | 删除工位设备关联 | WorkstationEquipmentRelationReqVO | Object |
| GET | /mes/workstation-opersteps/getWorkstationOperstepsPage | 根据工位code获得操作步骤 | WorkstationOperstepsPageReqVO | PageResult\<WorkstationOperstepsRespVO\> |
| POST | /mes/workstation-opersteps/create | 关联工位操作步骤 | WorkstationOperstepsRelationReqVO | Object |
| POST | /mes/workstation-opersteps/delete | 删除工位操作步骤关联 | WorkstationOperstepsRelationReqVO | Object |
| POST | /mes/workstation-post/create | 创建工位岗位关联 | WorkstationPostRelationReqVO | null |
| POST | /mes/workstation-post/delete | 删除工位岗位关联 | WorkstationPostRelationReqVO | Boolean |
| GET | /mes/workstation-post/getWorkstationPostPage | 获得工位岗位关联分页 | WorkstationPostPageReqVO | PageResult\<WorkstationPostRespVO\> |
| POST | /mes/workstation-post/senior | 高级搜索工位岗位关联 | CustomConditions | PageResult\<WorkstationPostRespVO\> |
| POST | /mes/workstation-order/create | 创建工位工单实时 | WorkstationOrderCreateReqVO | Long(id) |
| PUT | /mes/workstation-order/update | 更新工位工单实时 | WorkstationOrderUpdateReqVO | Boolean |
| DELETE | /mes/workstation-order/delete | 删除工位工单实时 | id (Long) | Boolean |
| GET | /mes/workstation-order/get | 获得工位工单实时 | id (Long) | WorkstationOrderRespVO |
| GET | /mes/workstation-order/list | 获得工位工单实时列表 | ids (Collection\<Long\>) | List\<WorkstationOrderRespVO\> |
| GET | /mes/workstation-order/page | 获得工位工单实时分页 | WorkstationOrderPageReqVO | PageResult\<WorkstationOrderRespVO\> |
| GET | /mes/workstation-order/export-excel | 导出工位工单实时Excel | WorkstationOrderExportReqVO | Excel文件流 |
| GET | /mes/workstation-order/get-import-template | 获得导入工位工单实时模板 | - | Excel文件流 |
| POST | /mes/workstation-order/import | 导入工位工单实时基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |
| POST | /mes/workstation-order-history/create | 创建工位工单历史 | WorkstationOrderHistoryCreateReqVO | Long(id) |
| PUT | /mes/workstation-order-history/update | 更新工位工单历史 | WorkstationOrderHistoryUpdateReqVO | Boolean |
| DELETE | /mes/workstation-order-history/delete | 删除工位工单历史 | id (Long) | Boolean |
| GET | /mes/workstation-order-history/get | 获得工位工单历史 | id (Long) | WorkstationOrderHistoryRespVO |
| GET | /mes/workstation-order-history/list | 获得工位工单历史列表 | ids (Collection\<Long\>) | List\<WorkstationOrderHistoryRespVO\> |
| GET | /mes/workstation-order-history/page | 获得工位工单历史分页 | WorkstationOrderHistoryPageReqVO | PageResult\<WorkstationOrderHistoryRespVO\> |
| GET | /mes/workstation-order-history/export-excel | 导出工位工单历史Excel | WorkstationOrderHistoryExportReqVO | Excel文件流 |
| GET | /mes/workstation-order-history/get-import-template | 获得导入工位工单历史模板 | - | Excel文件流 |
| POST | /mes/workstation-order-history/import | 导入工位工单历史基本信息 | MultipartFile, mode, updatePart | Map\<String,Object\> |

### 二十五、关联基础信息查询

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| GET | /mes/common/getProcessRouteByProductCode | 根据产品编码获得工艺路线的编码 | code (String) | List\<SelectRespVo\> |
| GET | /mes/common/geBomTreeByProductCode | 根据产品编码获得BOM树形结构 | code, version (String) | List\<ElBOMTreeVo\> |
| GET | /mes/common/getBomListByProductCode | 根据产品编码获得BOM列表 | code, version (String) | List\<BomDTO\> |
| GET | /mes/common/getBomListByProductBomAndProcess | 根据产品编码获得工序BOM物料 | code, version, processCode | List\<BomDTO\> |
| GET | /mes/common/getBomListByProductAndProcess | 根据产品编码获得工序物料 | code, processCode | List\<BomDTO\> |
| GET | /mes/common/getLinesByWorkRoomCode | 根据车间编码获得产线列表 | code (String) | List\<SelectRespVo\> |

---

> 文档生成时间: 2026-04-16
> 统计口径: 所有标注 `@Tag` 的Controller类，统计其中所有标注 `@GetMapping/@PostMapping/@PutMapping/@DeleteMapping` 的方法
