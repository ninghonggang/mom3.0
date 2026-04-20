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

## 11. 质检管理

### 11.1 质检类别表 (qms_qualityclass)

```sql
CREATE TABLE `qms_qualityclass` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `class_code` varchar(50) NOT NULL COMMENT '质检类别编码',
  `class_name` varchar(100) NOT NULL COMMENT '质检类别名称',
  `class_type` varchar(20) DEFAULT NULL COMMENT '类别类型(IQC/IPQC/FQC/OQC)',
  `parent_id` bigint DEFAULT NULL COMMENT '父级ID',
  `sort` int DEFAULT 0 COMMENT '排序',
  `status` int DEFAULT 1 COMMENT '状态(1启用/0禁用)',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_class_code` (`class_code`),
  KEY `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='质检类别表';
```

### 11.2 质检分组表 (qms_qualitygroup)

```sql
CREATE TABLE `qms_qualitygroup` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `group_code` varchar(50) NOT NULL COMMENT '质检分组编码',
  `group_name` varchar(100) NOT NULL COMMENT '质检分组名称',
  `group_type` varchar(20) DEFAULT NULL COMMENT '分组类型',
  `workshop_id` bigint DEFAULT NULL COMMENT '车间ID',
  `workshop_name` varchar(100) DEFAULT NULL COMMENT '车间名称',
  `line_id` bigint DEFAULT NULL COMMENT '产线ID',
  `line_name` varchar(100) DEFAULT NULL COMMENT '产线名称',
  `leader_id` bigint DEFAULT NULL COMMENT '组长ID',
  `leader_name` varchar(50) DEFAULT NULL COMMENT '组长名称',
  `member_count` int DEFAULT 0 COMMENT '成员数量',
  `status` int DEFAULT 1 COMMENT '状态(1启用/0禁用)',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_group_code` (`group_code`),
  KEY `idx_workshop_id` (`workshop_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='质检分组表';
```

### 11.3 质检项目表 (qms_item)

```sql
CREATE TABLE `qms_item` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `item_code` varchar(50) NOT NULL COMMENT '质检项目编码',
  `item_name` varchar(200) NOT NULL COMMENT '质检项目名称',
  `item_type` varchar(20) DEFAULT NULL COMMENT '项目类型(定量/定性)',
  `class_id` bigint DEFAULT NULL COMMENT '质检类别ID',
  `spec_lower` decimal(18,4) DEFAULT NULL COMMENT '规格下限',
  `spec_upper` decimal(18,4) DEFAULT NULL COMMENT '规格上限',
  `target_value` decimal(18,4) DEFAULT NULL COMMENT '目标值',
  `unit` varchar(20) DEFAULT NULL COMMENT '单位',
  `inspection_method` varchar(100) DEFAULT NULL COMMENT '检测方法',
  `aql` decimal(5,2) DEFAULT NULL COMMENT 'AQL值',
  `is_key_item` bit(1) DEFAULT b'0' COMMENT '是否关键项目',
  `status` int DEFAULT 1 COMMENT '状态(1启用/0禁用)',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_item_code` (`item_code`),
  KEY `idx_class_id` (`class_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='质检项目表';
```

### 11.4 质检表单表 (qms_qualityform)


```sql
CREATE TABLE `qms_qualityform` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `form_no` varchar(50) NOT NULL COMMENT '质检表单编号',
  `form_name` varchar(200) NOT NULL COMMENT '质检表单名称',
  `form_type` varchar(20) DEFAULT NULL COMMENT '表单类型(IQC/IPQC/FQC/OQC)',
  `source_type` varchar(20) DEFAULT NULL COMMENT '来源类型(工单/发货单/来料单)',
  `source_id` bigint DEFAULT NULL COMMENT '来源ID',
  `source_no` varchar(50) DEFAULT NULL COMMENT '来源单号',
  `product_id` bigint DEFAULT NULL COMMENT '产品ID',
  `product_code` varchar(50) DEFAULT NULL COMMENT '产品编码',
  `product_name` varchar(100) DEFAULT NULL COMMENT '产品名称',
  `quantity` decimal(18,4) DEFAULT NULL COMMENT '送检数量',
  `sample_size` int DEFAULT 0 COMMENT '抽样数量',
  `qualified_qty` decimal(18,4) DEFAULT 0 COMMENT '合格数量',
  `rejected_qty` decimal(18,4) DEFAULT 0 COMMENT '不合格数量',
  `inspector_id` bigint DEFAULT NULL COMMENT '质检员ID',
  `inspector_name` varchar(50) DEFAULT NULL COMMENT '质检员名称',
  `inspect_time` datetime DEFAULT NULL COMMENT '质检时间',
  `result` int DEFAULT 1 COMMENT '结果(1待检/2合格/3不合格)',
  `status` varchar(20) DEFAULT 'PENDING' COMMENT '状态(PENDING/IN_PROGRESS/COMPLETED)',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_form_no` (`form_no`),
  KEY `idx_source_id` (`source_id`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_inspector_id` (`inspector_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='质检表单表';
```

### 11.5 质检表单明细表 (qms_qualityformdetail)


```sql
CREATE TABLE `qms_qualityformdetail` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `form_id` bigint NOT NULL COMMENT '质检表单ID',
  `item_id` bigint NOT NULL COMMENT '质检项目ID',
  `item_code` varchar(50) DEFAULT NULL COMMENT '项目编码',
  `item_name` varchar(200) DEFAULT NULL COMMENT '项目名称',
  `item_type` varchar(20) DEFAULT NULL COMMENT '项目类型',
  `spec_lower` decimal(18,4) DEFAULT NULL COMMENT '规格下限',
  `spec_upper` decimal(18,4) DEFAULT NULL COMMENT '规格上限',
  `target_value` decimal(18,4) DEFAULT NULL COMMENT '目标值',
  `unit` varchar(20) DEFAULT NULL COMMENT '单位',
  `detect_value` decimal(18,4) DEFAULT NULL COMMENT '检测值',
  `is_normal` bit(1) DEFAULT b'1' COMMENT '是否正常',
  `result` int DEFAULT 1 COMMENT '结果(1合格/2不合格)',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_form_id` (`form_id`),
  KEY `idx_item_id` (`item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='质检表单明细表';
```

### 11.6 质检表单日志表 (qms_qualityformlog)

```sql
CREATE TABLE `qms_qualityformlog` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `form_id` bigint NOT NULL COMMENT '质检表单ID',
  `form_no` varchar(50) DEFAULT NULL COMMENT '表单编号',
  `operation_type` varchar(20) NOT NULL COMMENT '操作类型(CREATE/UPDATE/INSPECT/APPROVE)',
  `operator_id` bigint DEFAULT NULL COMMENT '操作人ID',
  `operator_name` varchar(50) DEFAULT NULL COMMENT '操作人名称',
  `operation_time` datetime DEFAULT NULL COMMENT '操作时间',
  `operation_content` varchar(500) DEFAULT NULL COMMENT '操作内容',
  `before_value` text DEFAULT NULL COMMENT '变更前值',
  `after_value` text DEFAULT NULL COMMENT '变更后值',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_form_id` (`form_id`),
  KEY `idx_operation_time` (`operation_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='质检表单日志表';
```

### 11.7 质检管理API接口

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/qualityclass/create | 创建质检类别 | QmsQualityclassCreateReqVO | Integer(id) |
| PUT | /mes/qualityclass/update | 更新质检类别 | QmsQualityclassUpdateReqVO | Boolean |
| DELETE | /mes/qualityclass/delete | 删除质检类别 | id (Integer) | Boolean |
| GET | /mes/qualityclass/get | 获得质检类别 | id (Integer) | QmsQualityclassRespVO |
| GET | /mes/qualityclass/list | 获得质检类别列表 | ids (Collection\<Integer\>) | List\<QmsQualityclassRespVO\> |
| GET | /mes/qualityclass/page | 获得质检类别分页 | QmsQualityclassPageReqVO | PageResult\<QmsQualityclassRespVO\> |
| POST | /mes/qualityclass/senior | 高级搜索质检类别 | CustomConditions | PageResult\<QmsQualityclassRespVO\> |
| GET | /mes/qualityclass/export-excel | 导出质检类别Excel | QmsQualityclassExportReqVO | Excel文件流 |
| POST | /mes/qualitygroup/create | 创建质检分组 | QmsQualitygroupCreateReqVO | Integer(id) |
| PUT | /mes/qualitygroup/update | 更新质检分组 | QmsQualitygroupUpdateReqVO | Boolean |
| DELETE | /mes/qualitygroup/delete | 删除质检分组 | id (Integer) | Boolean |
| GET | /mes/qualitygroup/get | 获得质检分组 | id (Integer) | QmsQualitygroupRespVO |
| GET | /mes/qualitygroup/list | 获得质检分组列表 | ids (Collection\<Integer\>) | List\<QmsQualitygroupRespVO\> |
| GET | /mes/qualitygroup/page | 获得质检分组分页 | QmsQualitygroupPageReqVO | PageResult\<QmsQualitygroupRespVO\> |
| POST | /mes/qualitygroup/senior | 高级搜索质检分组 | CustomConditions | PageResult\<QmsQualitygroupRespVO\> |
| GET | /mes/qualitygroup/export-excel | 导出质检分组Excel | QmsQualitygroupExportReqVO | Excel文件流 |
| POST | /mes/item/create | 创建质检项目 | QmsItemCreateReqVO | Integer(id) |
| PUT | /mes/item/update | 更新质检项目 | QmsItemUpdateReqVO | Boolean |
| DELETE | /mes/item/delete | 删除质检项目 | id (Integer) | Boolean |
| GET | /mes/item/get | 获得质检项目 | id (Integer) | QmsItemRespVO |
| GET | /mes/item/list | 获得质检项目列表 | ids (Collection\<Integer\>) | List\<QmsItemRespVO\> |
| GET | /mes/item/page | 获得质检项目分页 | QmsItemPageReqVO | PageResult\<QmsItemRespVO\> |
| POST | /mes/item/senior | 高级搜索质检项目 | CustomConditions | PageResult\<QmsItemRespVO\> |
| GET | /mes/item/export-excel | 导出质检项目Excel | QmsItemExportReqVO | Excel文件流 |
| POST | /mes/qualityform/create | 创建质检表单 | QmsQualityformCreateReqVO | Integer(id) |
| PUT | /mes/qualityform/update | 更新质检表单 | QmsQualityformUpdateReqVO | Boolean |
| DELETE | /mes/qualityform/delete | 删除质检表单 | id (Integer) | Boolean |
| GET | /mes/qualityform/get | 获得质检表单 | id (Integer) | QmsQualityformRespVO |
| GET | /mes/qualityform/list | 获得质检表单列表 | ids (Collection\<Integer\>) | List\<QmsQualityformRespVO\> |
| GET | /mes/qualityform/page | 获得质检表单分页 | QmsQualityformPageReqVO | PageResult\<QmsQualityformRespVO\> |
| POST | /mes/qualityform/senior | 高级搜索质检表单 | CustomConditions | PageResult\<QmsQualityformRespVO\> |
| GET | /mes/qualityform/export-excel | 导出质检表单Excel | QmsQualityformExportReqVO | Excel文件流 |
| POST | /mes/qualityformdetail/create | 创建质检表单明细 | QmsQualityformdetailCreateReqVO | Integer(id) |
| PUT | /mes/qualityformdetail/update | 更新质检表单明细 | QmsQualityformdetailUpdateReqVO | Boolean |
| DELETE | /mes/qualityformdetail/delete | 删除质检表单明细 | id (Integer) | Boolean |
| GET | /mes/qualityformdetail/get | 获得质检表单明细 | id (Integer) | QmsQualityformdetailRespVO |
| GET | /mes/qualityformdetail/list | 获得质检表单明细列表 | ids (Collection\<Integer\>) | List\<QmsQualityformdetailRespVO\> |
| GET | /mes/qualityformdetail/page | 获得质检表单明细分页 | QmsQualityformdetailPageReqVO | PageResult\<QmsQualityformdetailRespVO\> |
| POST | /mes/qualityformdetail/senior | 高级搜索质检表单明细 | CustomConditions | PageResult\<QmsQualityformdetailRespVO\> |
| GET | /mes/qualityformdetail/export-excel | 导出质检表单明细Excel | QmsQualityformdetailExportReqVO | Excel文件流 |
| PUT | /mes/qualityformlog/update | 更新质检表单日志 | QmsQualityformlogUpdateReqVO | Boolean |
| DELETE | /mes/qualityformlog/delete | 删除质检表单日志 | id (Integer) | Boolean |
| GET | /mes/qualityformlog/get | 获得质检表单日志 | id (Integer) | QmsQualityformlogRespVO |
| GET | /mes/qualityformlog/list | 获得质检表单日志列表 | ids (Collection\<Integer\>) | List\<QmsQualityformlogRespVO\> |
| GET | /mes/qualityformlog/page | 获得质检表单日志分页 | QmsQualityformlogPageReqVO | PageResult\<QmsQualityformlogRespVO\> |
| POST | /mes/qualityformlog/senior | 高级搜索质检表单日志 | CustomConditions | PageResult\<QmsQualityformlogRespVO\> |
| GET | /mes/qualityformlog/export-excel | 导出质检表单日志Excel | QmsQualityformlogExportReqVO | Excel文件流 |

---

## 12. 工位管理

### 12.1 工位台账表 (mes_workstation)

```sql
CREATE TABLE `mes_workstation` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `workshop_id` bigint NOT NULL COMMENT '车间ID',
  `workshop_name` varchar(100) DEFAULT NULL COMMENT '车间名称',
  `line_id` bigint DEFAULT NULL COMMENT '产线ID',
  `line_name` varchar(100) DEFAULT NULL COMMENT '产线名称',
  `station_code` varchar(50) NOT NULL COMMENT '工位编码',
  `station_name` varchar(100) NOT NULL COMMENT '工位名称',
  `station_type` varchar(20) DEFAULT NULL COMMENT '工位类型(加工/装配/检测/物流)',
  `status` varchar(20) DEFAULT 'IDLE' COMMENT '状态(IDLE运行/IDLE待机/MAINTENANCE保养/REPAIR维修)',
  `capacity` decimal(10,2) DEFAULT 0 COMMENT '产能(件/小时)',
  `standard_time` int DEFAULT 0 COMMENT '标准节拍时间(秒)',
  `worker_count` int DEFAULT 0 COMMENT '配置人数',
  `device_count` int DEFAULT 0 COMMENT '设备数量',
  `is_enabled` bit(1) DEFAULT b'1' COMMENT '是否启用',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_station_code` (`station_code`),
  KEY `idx_workshop_id` (`workshop_id`),
  KEY `idx_line_id` (`line_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='工位台账表';
```

### 12.2 工位能力关联表 (mes_workstation_ability)


```sql
CREATE TABLE `mes_workstation_ability` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `station_id` bigint NOT NULL COMMENT '工位ID',
  `station_code` varchar(50) DEFAULT NULL COMMENT '工位编码',
  `station_name` varchar(100) DEFAULT NULL COMMENT '工位名称',
  `ability_id` bigint NOT NULL COMMENT '能力ID',
  `ability_code` varchar(50) DEFAULT NULL COMMENT '能力编码',
  `ability_name` varchar(100) DEFAULT NULL COMMENT '能力名称',
  `ability_level` varchar(20) DEFAULT NULL COMMENT '能力等级(A/B/C)',
  `status` int DEFAULT 1 COMMENT '状态(1启用/0禁用)',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_station_id` (`station_id`),
  KEY `idx_ability_id` (`ability_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='工位能力关联表';
```

### 12.3 工位设备关联表 (mes_workstation_equipment)

```sql
CREATE TABLE `mes_workstation_equipment` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `station_id` bigint NOT NULL COMMENT '工位ID',
  `station_code` varchar(50) DEFAULT NULL COMMENT '工位编码',
  `station_name` varchar(100) DEFAULT NULL COMMENT '工位名称',
  `equipment_id` bigint NOT NULL COMMENT '设备ID',
  `equipment_code` varchar(50) DEFAULT NULL COMMENT '设备编码',
  `equipment_name` varchar(100) DEFAULT NULL COMMENT '设备名称',
  `equipment_type` varchar(20) DEFAULT NULL COMMENT '设备类型',
  `is_primary` bit(1) DEFAULT b'0' COMMENT '是否主设备',
  `status` int DEFAULT 1 COMMENT '状态(1启用/0禁用)',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_station_id` (`station_id`),
  KEY `idx_equipment_id` (`equipment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='工位设备关联表';
```

### 12.4 工位操作步骤表 (mes_workstation_opersteps)

```sql
CREATE TABLE `mes_workstation_opersteps` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `station_id` bigint NOT NULL COMMENT '工位ID',
  `station_code` varchar(50) DEFAULT NULL COMMENT '工位编码',
  `station_name` varchar(100) DEFAULT NULL COMMENT '工位名称',
  `step_code` varchar(50) DEFAULT NULL COMMENT '步骤编码',
  `step_name` varchar(200) NOT NULL COMMENT '步骤名称',
  `step_sequence` int NOT NULL COMMENT '步骤顺序',
  `operation_id` bigint DEFAULT NULL COMMENT '工序ID',
  `operation_code` varchar(50) DEFAULT NULL COMMENT '工序编码',
  `operation_name` varchar(100) DEFAULT NULL COMMENT '工序名称',
  `standard_time` int DEFAULT 0 COMMENT '标准工时(秒)',
  `quality_std` varchar(500) DEFAULT NULL COMMENT '质量标准',
  `is_key_step` bit(1) DEFAULT b'0' COMMENT '是否关键步骤',
  `status` int DEFAULT 1 COMMENT '状态(1启用/0禁用)',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_station_id` (`station_id`),
  KEY `idx_operation_id` (`operation_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='工位操作步骤表';
```

### 12.5 工位岗位关联表 (mes_workstation_post)

```sql
CREATE TABLE `mes_workstation_post` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `station_id` bigint NOT NULL COMMENT '工位ID',
  `station_code` varchar(50) DEFAULT NULL COMMENT '工位编码',
  `station_name` varchar(100) DEFAULT NULL COMMENT '工位名称',
  `post_id` bigint NOT NULL COMMENT '岗位ID',
  `post_code` varchar(50) DEFAULT NULL COMMENT '岗位编码',
  `post_name` varchar(100) DEFAULT NULL COMMENT '岗位名称',
  `required_count` int DEFAULT 1 COMMENT '需求人数',
  `status` int DEFAULT 1 COMMENT '状态(1启用/0禁用)',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_station_id` (`station_id`),
  KEY `idx_post_id` (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='工位岗位关联表';
```

### 12.6 工位管理API接口

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/workstation/create | 创建工位 | MesWorkstationCreateReqVO | Long(id) |
| PUT | /mes/workstation/update | 更新工位 | MesWorkstationUpdateReqVO | Boolean |
| DELETE | /mes/workstation/delete | 删除工位 | id (Long) | Boolean |
| GET | /mes/workstation/get | 获得工位 | id (Long) | MesWorkstationRespVO |
| GET | /mes/workstation/list | 获得工位列表 | ids (Collection\<Long\>) | List\<MesWorkstationRespVO\> |
| GET | /mes/workstation/page | 获得工位分页 | MesWorkstationPageReqVO | PageResult\<MesWorkstationRespVO\> |
| POST | /mes/workstation/senior | 高级搜索工位 | CustomConditions | PageResult\<MesWorkstationRespVO\> |
| POST | /mes/workstation-ability/create | 创建工位能力关联 | WorkstationAbilityRelationReqVO | Void |
| PUT | /mes/workstation-ability/update | 更新工位能力关联 | WorkstationAbilityRelationReqVO | Boolean |
| DELETE | /mes/workstation-ability/delete | 删除工位能力关联 | WorkstationAbilityRelationReqVO | Boolean |
| GET | /mes/workstation-ability/get | 获得工位能力关联 | id (Long) | WorkstationAbilityRespVO |
| GET | /mes/workstation-ability/list | 获得工位能力关联列表 | ids (Collection\<Long\>) | List\<WorkstationAbilityRespVO\> |
| GET | /mes/workstation-ability/page | 获得工位能力关联分页 | WorkstationAbilityPageReqVO | PageResult\<WorkstationAbilityRespVO\> |
| POST | /mes/workstation-ability/senior | 高级搜索工位能力关联 | CustomConditions | PageResult\<WorkstationAbilityRespVO\> |
| POST | /mes/workstation-equipment/create | 创建工位设备关联 | WorkstationEquipmentRelationReqVO | Long(id) |
| PUT | /mes/workstation-equipment/update | 更新工位设备关联 | WorkstationEquipmentRelationReqVO | Boolean |
| DELETE | /mes/workstation-equipment/delete | 删除工位设备关联 | id (Long) | Boolean |
| GET | /mes/workstation-equipment/get | 获得工位设备关联 | id (Long) | WorkstationEquipmentRespVO |
| GET | /mes/workstation-equipment/list | 获得工位设备关联列表 | ids (Collection\<Long\>) | List\<WorkstationEquipmentRespVO\> |
| GET | /mes/workstation-equipment/page | 获得工位设备关联分页 | WorkstationEquipmentPageReqVO | PageResult\<WorkstationEquipmentRespVO\> |
| POST | /mes/workstation-equipment/senior | 高级搜索工位设备关联 | CustomConditions | PageResult\<WorkstationEquipmentRespVO\> |
| POST | /mes/workstation-opersteps/create | 创建工位操作步骤 | WorkstationOperstepsRelationReqVO | Long(id) |
| PUT | /mes/workstation-opersteps/update | 更新工位操作步骤 | WorkstationOperstepsRelationReqVO | Boolean |
| DELETE | /mes/workstation-opersteps/delete | 删除工位操作步骤 | id (Long) | Boolean |
| GET | /mes/workstation-opersteps/get | 获得工位操作步骤 | id (Long) | WorkstationOperstepsRespVO |
| GET | /mes/workstation-opersteps/list | 获得工位操作步骤列表 | ids (Collection\<Long\>) | List\<WorkstationOperstepsRespVO\> |
| GET | /mes/workstation-opersteps/page | 获得工位操作步骤分页 | WorkstationOperstepsPageReqVO | PageResult\<WorkstationOperstepsRespVO\> |
| POST | /mes/workstation-opersteps/senior | 高级搜索工位操作步骤 | CustomConditions | PageResult\<WorkstationOperstepsRespVO\> |
| POST | /mes/workstation-post/create | 创建工位岗位关联 | WorkstationPostRelationReqVO | Long(id) |
| PUT | /mes/workstation-post/update | 更新工位岗位关联 | WorkstationPostRelationReqVO | Boolean |
| DELETE | /mes/workstation-post/delete | 删除工位岗位关联 | id (Long) | Boolean |
| GET | /mes/workstation-post/get | 获得工位岗位关联 | id (Long) | WorkstationPostRespVO |
| GET | /mes/workstation-post/list | 获得工位岗位关联列表 | ids (Collection\<Long\>) | List\<WorkstationPostRespVO\> |
| GET | /mes/workstation-post/page | 获得工位岗位关联分页 | WorkstationPostPageReqVO | PageResult\<WorkstationPostRespVO\> |
| POST | /mes/workstation-post/senior | 高级搜索工位岗位关联 | CustomConditions | PageResult\<WorkstationPostRespVO\> |

---

## 13. 设备基础信息

### 13.1 设备基础信息表 (mes_device_info)

```sql
CREATE TABLE `mes_device_info` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `device_code` varchar(50) NOT NULL COMMENT '设备编码',
  `device_name` varchar(200) NOT NULL COMMENT '设备名称',
  `device_type` varchar(50) DEFAULT NULL COMMENT '设备类型',
  `brand` varchar(100) DEFAULT NULL COMMENT '品牌',
  `model` varchar(100) DEFAULT NULL COMMENT '型号',
  `serial_number` varchar(100) DEFAULT NULL COMMENT '序列号',
  `manufacturer` varchar(200) DEFAULT NULL COMMENT '制造商',
  `workshop_id` bigint DEFAULT NULL COMMENT '所属车间ID',
  `workshop_name` varchar(100) DEFAULT NULL COMMENT '所属车间名称',
  `line_id` bigint DEFAULT NULL COMMENT '所属产线ID',
  `line_name` varchar(100) DEFAULT NULL COMMENT '所属产线名称',
  `station_id` bigint DEFAULT NULL COMMENT '所属工位ID',
  `station_name` varchar(100) DEFAULT NULL COMMENT '所属工位名称',
  `purchase_date` date DEFAULT NULL COMMENT '采购日期',
  `purchase_price` decimal(18,2) DEFAULT NULL COMMENT '采购价格',
  `warranty_end_date` date DEFAULT NULL COMMENT '保修期结束日期',
  `supplier_id` bigint DEFAULT NULL COMMENT '供应商ID',
  `supplier_name` varchar(100) DEFAULT NULL COMMENT '供应商名称',
  `status` varchar(20) DEFAULT 'RUNNING' COMMENT '状态(RUNNING运行/IDLE待机/DOWNTIME故障/MAINTENANCE保养/SCRAPPED报废)',
  `running_time` int DEFAULT 0 COMMENT '累计运行时间(小时)',
  `maintenance_count` int DEFAULT 0 COMMENT '保养次数',
  `repair_count` int DEFAULT 0 COMMENT '维修次数',
  `fault_count` int DEFAULT 0 COMMENT '故障次数',
  `oee` decimal(5,2) DEFAULT 0 COMMENT 'OEE值(%)',
  `image_urls` text DEFAULT NULL COMMENT '设备图片URLs(JSON数组)',
  `specification` text DEFAULT NULL COMMENT '设备规格',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_device_code` (`device_code`),
  KEY `idx_workshop_id` (`workshop_id`),
  KEY `idx_line_id` (`line_id`),
  KEY `idx_station_id` (`station_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备基础信息表';
```

### 13.2 设备基础信息API接口

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/device/create | 创建设备 | MesDeviceCreateReqVO | Long(id) |
| PUT | /mes/device/update | 更新设备 | MesDeviceUpdateReqVO | Boolean |
| DELETE | /mes/device/delete | 删除设备 | id (Long) | Boolean |
| GET | /mes/device/get | 获得设备 | id (Long) | MesDeviceRespVO |
| GET | /mes/device/list | 获得设备列表 | ids (Collection\<Long\>) | List\<MesDeviceRespVO\> |
| GET | /mes/device/page | 获得设备分页 | MesDevicePageReqVO | PageResult\<MesDeviceRespVO\> |
| POST | /mes/device/senior | 高级搜索设备 | CustomConditions | PageResult\<MesDeviceRespVO\> |
| GET | /mes/device/export-excel | 导出设备Excel | MesDeviceExportReqVO | Excel文件流 |
| GET | /mes/device/get-import-template | 获得导入设备模板 | - | Excel文件流 |
| POST | /mes/device/import | 导入设备 | MultipartFile, mode, updatePart | Map\<String,Object\> |
| PUT | /mes/device/update-status | 更新设备状态 | MesDeviceStatusUpdateReqVO | Boolean |
| GET | /mes/device/statistics | 获得设备统计 | - | MesDeviceStatisticsRespVO |

---

## 14. 日计划BOM管理（补充）

### 14.1 日计划BOM实例表 (plan_mes_order_day_bom)

```sql
CREATE TABLE `plan_mes_order_day_bom` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `order_id` bigint NOT NULL COMMENT '日计划ID',
  `product_id` bigint DEFAULT NULL COMMENT '产品ID',
  `product_code` varchar(50) DEFAULT NULL COMMENT '产品编码',
  `product_name` varchar(200) DEFAULT NULL COMMENT '产品名称',
  `bom_id` bigint DEFAULT NULL COMMENT 'BOM配置ID',
  `bom_code` varchar(50) DEFAULT NULL COMMENT 'BOM编码',
  `version` varchar(20) DEFAULT NULL COMMENT 'BOM版本',
  `quantity` decimal(18,4) DEFAULT 0 COMMENT '需求数量',
  `status` varchar(20) DEFAULT 'PENDING' COMMENT '状态(PENDING待配料/ALLOCATED已配料/SUFFICIENT已齐套/INSUFFICIENT不足)',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_bom_id` (`bom_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日计划BOM实例表';
```

### 14.2 日计划BOM物料明细表 (plan_mes_order_day_bom_detail)

```sql
CREATE TABLE `plan_mes_order_day_bom_detail` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `order_day_bom_id` bigint NOT NULL COMMENT '日计划BOM实例ID',
  `material_id` bigint DEFAULT NULL COMMENT '物料ID',
  `material_code` varchar(50) DEFAULT NULL COMMENT '物料编码',
  `material_name` varchar(200) DEFAULT NULL COMMENT '物料名称',
  `spec` varchar(100) DEFAULT NULL COMMENT '规格型号',
  `unit` varchar(20) DEFAULT NULL COMMENT '单位',
  `bom_quantity` decimal(18,4) DEFAULT 0 COMMENT 'BOM需求量',
  `actual_quantity` decimal(18,4) DEFAULT 0 COMMENT '实际库存量',
  `required_quantity` decimal(18,4) DEFAULT 0 COMMENT '需求数量',
  `allocated_quantity` decimal(18,4) DEFAULT 0 COMMENT '已分配数量',
  `status` varchar(20) DEFAULT 'PENDING' COMMENT '状态(PENDING待配料/ALLOCATED已配料/SHORTAGE短缺)',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  PRIMARY KEY (`id`),
  KEY `idx_order_day_bom_id` (`order_day_bom_id`),
  KEY `idx_material_id` (`material_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日计划BOM物料明细表';
```

### 14.3 日计划BOM管理API接口

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/orderday/bom/create | 创建日计划BOM实例 | MesOrderDayBomCreateReqVO | Long(id) |
| PUT | /mes/orderday/bom/update | 更新日计划BOM实例 | MesOrderDayBomUpdateReqVO | Boolean |
| DELETE | /mes/orderday/bom/delete | 删除日计划BOM实例 | id (Long) | Boolean |
| GET | /mes/orderday/bom/get | 获得日计划BOM实例 | id (Long) | MesOrderDayBomRespVO |
| GET | /mes/orderday/bom/list | 获得日计划BOM实例列表 | ids (Collection\<Long\>) | List\<MesOrderDayBomRespVO\> |
| GET | /mes/orderday/bom/page | 获得日计划BOM实例分页 | MesOrderDayBomPageReqVO | PageResult\<MesOrderDayBomRespVO\> |
| POST | /mes/orderday/bom/senior | 高级搜索日计划BOM | CustomConditions | PageResult\<MesOrderDayBomRespVO\> |
| POST | /mes/orderday/bom/batchCreate | 批量创建日计划BOM | List\<MesOrderDayBomCreateReqVO\> | Boolean |
| PUT | /mes/orderday/bom/allocate | 配料分配 | MesOrderDayBomAllocateReqVO | Boolean |
| GET | /mes/orderday/bom/export-excel | 导出日计划BOM Excel | MesOrderDayBomExportReqVO | Excel文件流 |

---

## 15. 日计划工艺路线管理（补充）

### 15.1 日计划工艺路线实例表 (plan_mes_order_day_route)

```sql
CREATE TABLE `plan_mes_order_day_route` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `order_id` bigint NOT NULL COMMENT '日计划ID',
  `order_day_bom_id` bigint DEFAULT NULL COMMENT '日计划BOM实例ID',
  `route_id` bigint DEFAULT NULL COMMENT '工艺路线定义ID',
  `route_code` varchar(50) DEFAULT NULL COMMENT '工艺路线编码',
  `route_name` varchar(200) DEFAULT NULL COMMENT '工艺路线名称',
  `version` varchar(20) DEFAULT NULL COMMENT '版本号',
  `status` varchar(20) DEFAULT 'PENDING' COMMENT '状态(PENDING待配置/CONFIGURED已配置/PUBLISHED已发布)',
  `effective_date` date DEFAULT NULL COMMENT '生效日期',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_route_id` (`route_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日计划工艺路线实例表';
```

### 15.2 日计划工序明细表 (plan_mes_order_day_routesub)

```sql
CREATE TABLE `plan_mes_order_day_routesub` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `order_day_route_id` bigint NOT NULL COMMENT '日计划工艺路线实例ID',
  `sequence` int DEFAULT 0 COMMENT '工序序号',
  `process_code` varchar(50) DEFAULT NULL COMMENT '工序编码',
  `process_name` varchar(200) DEFAULT NULL COMMENT '工序名称',
  `workstation_id` bigint DEFAULT NULL COMMENT '工位ID',
  `workstation_code` varchar(50) DEFAULT NULL COMMENT '工位编码',
  `workstation_name` varchar(100) DEFAULT NULL COMMENT '工位名称',
  `standard_time` int DEFAULT 0 COMMENT '标准工时(秒)',
  `actual_time` int DEFAULT 0 COMMENT '实际工时(秒)',
  `status` varchar(20) DEFAULT 'PENDING' COMMENT '状态(PENDING待生产/RUNNING进行中/COMPLETED已完成/SUSPENDED已挂起)',
  `start_time` datetime DEFAULT NULL COMMENT '实际开始时间',
  `end_time` datetime DEFAULT NULL COMMENT '实际结束时间',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  PRIMARY KEY (`id`),
  KEY `idx_order_day_route_id` (`order_day_route_id`),
  KEY `idx_workstation_id` (`workstation_id`),
  KEY `idx_sequence` (`sequence`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日计划工序明细表';
```

### 15.3 日计划工艺路线API接口

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/orderday/route/create | 创建日计划工艺路线实例 | MesOrderDayRouteCreateReqVO | Long(id) |
| PUT | /mes/orderday/route/update | 更新日计划工艺路线实例 | MesOrderDayRouteUpdateReqVO | Boolean |
| DELETE | /mes/orderday/route/delete | 删除日计划工艺路线实例 | id (Long) | Boolean |
| GET | /mes/orderday/route/get | 获得日计划工艺路线实例 | id (Long) | MesOrderDayRouteRespVO |
| GET | /mes/orderday/route/getByOrder | 获得计划产品工艺路线配置 | orderId (Long) | MesOrderDayRouteRespVO |
| GET | /mes/orderday/route/list | 获得日计划工艺路线实例列表 | ids (Collection\<Long\>) | List\<MesOrderDayRouteRespVO\> |
| GET | /mes/orderday/route/page | 获得日计划工艺路线实例分页 | MesOrderDayRoutePageReqVO | PageResult\<MesOrderDayRouteRespVO\> |
| POST | /mes/orderday/route/senior | 高级搜索日计划工艺路线 | CustomConditions | PageResult\<MesOrderDayRouteRespVO\> |
| POST | /mes/orderday/routesub/create | 创建日计划工序明细 | MesOrderDayRoutesubCreateReqVO | Long(id) |
| PUT | /mes/orderday/routesub/update | 更新日计划工序明细 | MesOrderDayRoutesubUpdateReqVO | Boolean |
| DELETE | /mes/orderday/routesub/delete | 删除日计划工序明细 | id (Long) | Boolean |
| GET | /mes/orderday/routesub/get | 获得日计划工序明细 | id (Long) | MesOrderDayRoutesubRespVO |
| GET | /mes/orderday/routesub/list | 获得日计划工序明细列表 | ids (Collection\<Long\>) | List\<MesOrderDayRoutesubRespVO\> |
| GET | /mes/orderday/routesub/page | 获得日计划工序明细分页 | MesOrderDayRoutesubPageReqVO | PageResult\<MesOrderDayRoutesubRespVO\> |
| POST | /mes/orderday/routesub/senior | 高级搜索日计划工序明细 | CustomConditions | PageResult\<MesOrderDayRoutesubRespVO\> |
| PUT | /mes/orderday/routesub/bindWorkstation | 绑定工位到工序 | RoutesubWorkstationBindReqVO | Boolean |
| PUT | /mes/orderday/routesub/bindEquipment | 绑定设备到工序 | RoutesubEquipmentBindReqVO | Boolean |
| PUT | /mes/orderday/routesub/bindWorker | 绑定人员到工序 | RoutesubWorkerBindReqVO | Boolean |

---

## 16. 日计划资源配置管理（补充）

### 16.1 日计划设备配置表 (plan_mes_order_day_equipment)

```sql
CREATE TABLE `plan_mes_order_day_equipment` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `order_id` bigint NOT NULL COMMENT '日计划ID',
  `order_day_routesub_id` bigint DEFAULT NULL COMMENT '日计划工序明细ID',
  `equipment_id` bigint DEFAULT NULL COMMENT '设备ID',
  `equipment_code` varchar(50) DEFAULT NULL COMMENT '设备编码',
  `equipment_name` varchar(200) DEFAULT NULL COMMENT '设备名称',
  `equipment_type` varchar(50) DEFAULT NULL COMMENT '设备类型',
  `is_primary` bit(1) DEFAULT b'0' COMMENT '是否主设备',
  `status` varchar(20) DEFAULT 'IDLE' COMMENT '状态(IDLE待机/ALLOCATED已分配/RUNNING运行中/MAINTENANCE保养中)',
  `bind_time` datetime DEFAULT NULL COMMENT '绑定时间',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_order_day_routesub_id` (`order_day_routesub_id`),
  KEY `idx_equipment_id` (`equipment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日计划设备配置表';
```

### 16.2 日计划人员配置表 (plan_mes_order_day_worker)

```sql
CREATE TABLE `plan_mes_order_day_worker` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `order_id` bigint NOT NULL COMMENT '日计划ID',
  `order_day_routesub_id` bigint DEFAULT NULL COMMENT '日计划工序明细ID',
  `worker_id` bigint DEFAULT NULL COMMENT '人员ID',
  `worker_code` varchar(50) DEFAULT NULL COMMENT '工号',
  `worker_name` varchar(100) DEFAULT NULL COMMENT '人员姓名',
  `team_id` bigint DEFAULT NULL COMMENT '班组ID',
  `team_name` varchar(100) DEFAULT NULL COMMENT '班组名称',
  `post_id` bigint DEFAULT NULL COMMENT '岗位ID',
  `post_name` varchar(100) DEFAULT NULL COMMENT '岗位名称',
  `role_type` varchar(20) DEFAULT NULL COMMENT '角色类型(OPERATOR操作员/LEADER组长/QC质检员)',
  `status` varchar(20) DEFAULT 'AVAILABLE' COMMENT '状态(AVAILABLE可用/BIND已绑定/ABSENT缺勤)',
  `bind_time` datetime DEFAULT NULL COMMENT '绑定时间',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_order_day_routesub_id` (`order_day_routesub_id`),
  KEY `idx_worker_id` (`worker_id`),
  KEY `idx_team_id` (`team_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日计划人员配置表';
```

### 16.3 日计划工位配置表 (plan_mes_order_day_workstation)

```sql
CREATE TABLE `plan_mes_order_day_workstation` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `order_id` bigint NOT NULL COMMENT '日计划ID',
  `order_day_routesub_id` bigint DEFAULT NULL COMMENT '日计划工序明细ID',
  `workstation_id` bigint DEFAULT NULL COMMENT '工位ID',
  `workstation_code` varchar(50) DEFAULT NULL COMMENT '工位编码',
  `workstation_name` varchar(200) DEFAULT NULL COMMENT '工位名称',
  `workshop_id` bigint DEFAULT NULL COMMENT '车间ID',
  `workshop_name` varchar(100) DEFAULT NULL COMMENT '车间名称',
  `line_id` bigint DEFAULT NULL COMMENT '产线ID',
  `line_name` varchar(100) DEFAULT NULL COMMENT '产线名称',
  `status` varchar(20) DEFAULT 'IDLE' COMMENT '状态(IDLE待机/ALLOCATED已分配/PRODUCTING生产中/MAINTENANCE保养中)',
  `capacity` decimal(10,2) DEFAULT 0 COMMENT '产能(件/小时)',
  `used_capacity` decimal(10,2) DEFAULT 0 COMMENT '已用产能',
  `bind_time` datetime DEFAULT NULL COMMENT '绑定时间',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_order_day_routesub_id` (`order_day_routesub_id`),
  KEY `idx_workstation_id` (`workstation_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日计划工位配置表';
```

### 16.4 日计划资源配置API接口

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/orderday/equipment/create | 创建设备配置 | MesOrderDayEquipmentCreateReqVO | Long(id) |
| POST | /mes/orderday/equipment/batchCreate | 批量创建设备配置 | List\<MesOrderDayEquipmentCreateReqVO\> | Boolean |
| PUT | /mes/orderday/equipment/update | 更新设备配置 | MesOrderDayEquipmentUpdateReqVO | Boolean |
| DELETE | /mes/orderday/equipment/delete | 删除设备配置 | id (Long) | Boolean |
| GET | /mes/orderday/equipment/get | 获得设备配置 | id (Long) | MesOrderDayEquipmentRespVO |
| GET | /mes/orderday/equipment/list | 获得设备配置列表 | ids (Collection\<Long\>) | List\<MesOrderDayEquipmentRespVO\> |
| GET | /mes/orderday/equipment/page | 获得设备配置分页 | MesOrderDayEquipmentPageReqVO | PageResult\<MesOrderDayEquipmentRespVO\> |
| POST | /mes/orderday/equipment/senior | 高级搜索设备配置 | CustomConditions | PageResult\<MesOrderDayEquipmentRespVO\> |
| POST | /mes/orderday/worker/create | 创建人员配置 | MesOrderDayWorkerCreateReqVO | Long(id) |
| POST | /mes/orderday/worker/batchCreate | 批量创建人员配置 | List\<MesOrderDayWorkerCreateReqVO\> | Boolean |
| PUT | /mes/orderday/worker/update | 更新人员配置 | MesOrderDayWorkerUpdateReqVO | Boolean |
| DELETE | /mes/orderday/worker/delete | 删除人员配置 | id (Long) | Boolean |
| GET | /mes/orderday/worker/get | 获得人员配置 | id (Long) | MesOrderDayWorkerRespVO |
| GET | /mes/orderday/worker/list | 获得人员配置列表 | ids (Collection\<Long\>) | List\<MesOrderDayWorkerRespVO\> |
| GET | /mes/orderday/worker/page | 获得人员配置分页 | MesOrderDayWorkerPageReqVO | PageResult\<MesOrderDayWorkerRespVO\> |
| POST | /mes/orderday/worker/senior | 高级搜索人员配置 | CustomConditions | PageResult\<MesOrderDayWorkerRespVO\> |
| POST | /mes/orderday/workstation/create | 创建工位配置 | MesOrderDayWorkstationCreateReqVO | Long(id) |
| POST | /mes/orderday/workstation/batchCreate | 批量创建工位配置 | List\<MesOrderDayWorkstationCreateReqVO\> | Boolean |
| PUT | /mes/orderday/workstation/update | 更新工位配置 | MesOrderDayWorkstationUpdateReqVO | Boolean |
| DELETE | /mes/orderday/workstation/delete | 删除工位配置 | id (Long) | Boolean |
| GET | /mes/orderday/workstation/get | 获得工位配置 | id (Long) | MesOrderDayWorkstationRespVO |
| GET | /mes/orderday/workstation/list | 获得工位配置列表 | ids (Collection\<Long\>) | List\<MesOrderDayWorkstationRespVO\> |
| GET | /mes/orderday/workstation/page | 获得工位配置分页 | MesOrderDayWorkstationPageReqVO | PageResult\<MesOrderDayWorkstationRespVO\> |
| POST | /mes/orderday/workstation/senior | 高级搜索工位配置 | CustomConditions | PageResult\<MesOrderDayWorkstationRespVO\> |

---

## 17. 工序排程明细管理（补充）

### 17.1 工序排程明细表 (plan_mes_work_scheduling_detail)

```sql
CREATE TABLE `plan_mes_work_scheduling_detail` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `scheduling_id` bigint NOT NULL COMMENT '排程主表ID',
  `scheduling_code` varchar(50) DEFAULT NULL COMMENT '排程单号',
  `order_day_routesub_id` bigint DEFAULT NULL COMMENT '日计划工序明细ID',
  `sequence` int DEFAULT 0 COMMENT '工序序号',
  `process_code` varchar(50) DEFAULT NULL COMMENT '工序编码',
  `process_name` varchar(200) DEFAULT NULL COMMENT '工序名称',
  `workstation_id` bigint DEFAULT NULL COMMENT '工位ID',
  `workstation_code` varchar(50) DEFAULT NULL COMMENT '工位编码',
  `workstation_name` varchar(100) DEFAULT NULL COMMENT '工位名称',
  `equipment_ids` varchar(500) DEFAULT NULL COMMENT '设备ID列表(JSON)',
  `worker_ids` varchar(500) DEFAULT NULL COMMENT '人员ID列表(JSON)',
  `plan_quantity` decimal(18,4) DEFAULT 0 COMMENT '计划数量',
  `completed_quantity` decimal(18,4) DEFAULT 0 COMMENT '已完成数量',
  `qualified_quantity` decimal(18,4) DEFAULT 0 COMMENT '合格数量',
  `rejected_quantity` decimal(18,4) DEFAULT 0 COMMENT '不合格数量',
  `standard_time` int DEFAULT 0 COMMENT '标准工时(秒)',
  `actual_start_time` datetime DEFAULT NULL COMMENT '实际开始时间',
  `actual_end_time` datetime DEFAULT NULL COMMENT '实际结束时间',
  `status` varchar(20) DEFAULT 'PENDING' COMMENT '状态(PENDING待生产/READY待开工/RUNNING进行中/COMPLETED已完成/QUALIFIED已质检/SUSPENDED已挂起)',
  `quality_status` varchar(20) DEFAULT 'PENDING' COMMENT '质检状态(PENDING待质检/PASSED合格/FAILED不合格)',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  PRIMARY KEY (`id`),
  KEY `idx_scheduling_id` (`scheduling_id`),
  KEY `idx_order_day_routesub_id` (`order_day_routesub_id`),
  KEY `idx_workstation_id` (`workstation_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='工序排程明细表';
```

### 17.2 工序排程明细API接口

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/work-scheduling-detail/create | 创建工序排程明细 | MesWorkSchedulingDetailCreateReqVO | Long(id) |
| POST | /mes/work-scheduling-detail/batchCreate | 批量创建工序排程明细 | List\<MesWorkSchedulingDetailCreateReqVO\> | Boolean |
| PUT | /mes/work-scheduling-detail/update | 更新工序排程明细 | MesWorkSchedulingDetailUpdateReqVO | Boolean |
| DELETE | /mes/work-scheduling-detail/delete | 删除工序排程明细 | id (Long) | Boolean |
| GET | /mes/work-scheduling-detail/get | 获得工序排程明细 | id (Long) | MesWorkSchedulingDetailRespVO |
| GET | /mes/work-scheduling-detail/list | 获得工序排程明细列表 | ids (Collection\<Long\>) | List\<MesWorkSchedulingDetailRespVO\> |
| GET | /mes/work-scheduling-detail/page | 获得工序排程明细分页 | MesWorkSchedulingDetailPageReqVO | PageResult\<MesWorkSchedulingDetailRespVO\> |
| POST | /mes/work-scheduling-detail/senior | 高级搜索工序排程明细 | CustomConditions | PageResult\<MesWorkSchedulingDetailRespVO\> |
| PUT | /mes/work-scheduling-detail/start | 工序开工 | MesWorkSchedulingDetailStartReqVO | Boolean |
| PUT | /mes/work-scheduling-detail/pause | 工序暂停 | MesWorkSchedulingDetailPauseReqVO | Boolean |
| PUT | /mes/work-scheduling-detail/resume | 工序恢复 | MesWorkSchedulingDetailResumeReqVO | Boolean |
| PUT | /mes/work-scheduling-detail/complete | 工序完工 | MesWorkSchedulingDetailCompleteReqVO | Boolean |
| POST | /mes/work-scheduling-detail/report | 工序报工 | MesWorkSchedulingDetailReportReqVO | Boolean |
| PUT | /mes/work-scheduling-detail/bindEquipment | 绑定设备 | MesWorkSchedulingDetailBindEquipReqVO | Boolean |
| PUT | /mes/work-scheduling-detail/bindWorker | 绑定人员 | MesWorkSchedulingDetailBindWorkerReqVO | Boolean |

---

## 18. 齐套检查API（补充）

### 18.1 齐套检查业务说明

齐套检查是日计划发布前的关键校验环节，确保生产所需的物料、设备、人员、工位等资源均已就绪。

**检查维度**：
- 物料齐套：校验BOM物料是否充足
- 设备齐套：校验工序所需设备是否可用
- 人员齐套：校验工序所需人员是否配置
- 工位齐套：校验工序所需工位是否可用

### 18.2 齐套检查API接口

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/orderday/check/stuffing | 齐套检查 | MesOrderDayStuffingCheckReqVO | MesOrderDayStuffingCheckRespVO |
| POST | /mes/orderday/check/stuffing/bom | 物料齐套检查 | MesOrderDayStuffingCheckReqVO | MesOrderDayBomStuffingRespVO |
| POST | /mes/orderday/check/stuffing/equipment | 设备齐套检查 | MesOrderDayStuffingCheckReqVO | MesOrderDayEquipmentStuffingRespVO |
| POST | /mes/orderday/check/stuffing/worker | 人员齐套检查 | MesOrderDayStuffingCheckReqVO | MesOrderDayWorkerStuffingRespVO |
| POST | /mes/orderday/check/stuffing/workstation | 工位齐套检查 | MesOrderDayStuffingCheckReqVO | MesOrderDayWorkstationStuffingRespVO |
| GET | /mes/orderday/check/stuffing/report | 获取齐套检查报告 | orderId (Long) | MesOrderDayStuffingReportRespVO |

---

## 19. 工位管理API（补充）

### 19.1 工位管理业务说明

工位是生产执行的基本单元，关联车间、产线、设备、人员等资源。

**核心功能**：
- 工位台账维护
- 工位能力配置
- 工位与设备关联
- 工位与工序关联
- 工位状态监控

### 19.2 工位管理API接口

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/workstation/create | 创建工位 | MesWorkstationCreateReqVO | Long(id) |
| PUT | /mes/workstation/update | 更新工位 | MesWorkstationUpdateReqVO | Boolean |
| DELETE | /mes/workstation/delete | 删除工位 | id (Long) | Boolean |
| GET | /mes/workstation/get | 获得工位 | id (Long) | MesWorkstationRespVO |
| GET | /mes/workstation/list | 获得工位列表 | ids (Collection\<Long\>) | List\<MesWorkstationRespVO\> |
| GET | /mes/workstation/page | 获得工位分页 | MesWorkstationPageReqVO | PageResult\<MesWorkstationRespVO\> |
| POST | /mes/workstation/senior | 高级搜索工位 | CustomConditions | PageResult\<MesWorkstationRespVO\> |
| GET | /mes/workstation/export-excel | 导出工位Excel | MesWorkstationExportReqVO | Excel文件流 |
| POST | /mes/workstation/import | 导入工位 | MultipartFile | Map\<String,Object\> |
| PUT | /mes/workstation/update-status | 更新工位状态 | MesWorkstationStatusUpdateReqVO | Boolean |
| GET | /mes/workstation/statistics | 获取工位统计 | - | MesWorkstationStatisticsRespVO |
| POST | /mes/workstation/bindLine | 绑定产线 | WorkstationBindLineReqVO | Boolean |
| POST | /mes/workstation/unbindLine | 解绑产线 | id (Long) | Boolean |

---

## 20. 人员班组API（补充）

### 20.1 人员班组业务说明

人员班组管理支持生产现场的人员组织、班组划分和排班管理。

**核心功能**：
- 班组定义与维护
- 班组人员管理
- 人员技能配置
- 人员排班管理
- 考勤管理

### 20.2 人员班组API接口

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/worker/team/create | 创建班组 | MesWorkerTeamCreateReqVO | Long(id) |
| PUT | /mes/worker/team/update | 更新班组 | MesWorkerTeamUpdateReqVO | Boolean |
| DELETE | /mes/worker/team/delete | 删除班组 | id (Long) | Boolean |
| GET | /mes/worker/team/get | 获得班组 | id (Long) | MesWorkerTeamRespVO |
| GET | /mes/worker/team/list | 获得班组列表 | ids (Collection\<Long\>) | List\<MesWorkerTeamRespVO\> |
| GET | /mes/worker/team/page | 获得班组分页 | MesWorkerTeamPageReqVO | PageResult\<MesWorkerTeamRespVO\> |
| POST | /mes/worker/team/senior | 高级搜索班组 | CustomConditions | PageResult\<MesWorkerTeamRespVO\> |
| POST | /mes/worker/member/create | 创建班组成员 | MesWorkerMemberCreateReqVO | Long(id) |
| PUT | /mes/worker/member/update | 更新班组成员 | MesWorkerMemberUpdateReqVO | Boolean |
| DELETE | /mes/worker/member/delete | 删除班组成员 | id (Long) | Boolean |
| GET | /mes/worker/member/get | 获得班组成员 | id (Long) | MesWorkerMemberRespVO |
| GET | /mes/worker/member/list | 获得班组成员列表 | ids (Collection\<Long\>) | List\<MesWorkerMemberRespVO\> |
| GET | /mes/worker/member/page | 获得班组成员分页 | MesWorkerMemberPageReqVO | PageResult\<MesWorkerMemberRespVO\> |
| POST | /mes/worker/schedule/create | 创建人员排班 | MesWorkerScheduleCreateReqVO | Long(id) |
| PUT | /mes/worker/schedule/update | 更新人员排班 | MesWorkerScheduleUpdateReqVO | Boolean |
| DELETE | /mes/worker/schedule/delete | 删除人员排班 | id (Long) | Boolean |
| GET | /mes/worker/schedule/get | 获得人员排班 | id (Long) | MesWorkerScheduleRespVO |
| GET | /mes/worker/schedule/page | 获得人员排班分页 | MesWorkerSchedulePageReqVO | PageResult\<MesWorkerScheduleRespVO\> |
| GET | /mes/worker/attendance/export | 导出考勤记录 | MesWorkerAttendanceExportReqVO | Excel文件流 |

---

## 21. 设备绑定API（补充）

### 21.1 设备绑定业务说明

设备绑定实现设备与工位、订单的关联管理，支持设备资源调度和状态跟踪。

**核心功能**：
- 设备与工位绑定
- 设备与订单绑定
- 设备状态管理
- 设备运行监控

### 21.2 设备绑定API接口

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/equipment/bind/workstation | 设备绑定工位 | EquipmentBindWorkstationReqVO | Boolean |
| POST | /mes/equipment/bind/workstation/batch | 批量设备绑定工位 | List\<EquipmentBindWorkstationReqVO\> | Boolean |
| POST | /mes/equipment/bind/order | 设备绑定订单 | EquipmentBindOrderReqVO | Boolean |
| POST | /mes/equipment/bind/order/batch | 批量设备绑定订单 | List\<EquipmentBindOrderReqVO\> | Boolean |
| POST | /mes/equipment/unbind | 设备解绑 | EquipmentUnbindReqVO | Boolean |
| GET | /mes/equipment/bind/list | 获取设备绑定列表 | equipmentId (Long) | List\<EquipmentBindRespVO\> |
| GET | /mes/equipment/bind/workstation/list | 获取工位设备列表 | workstationId (Long) | List\<EquipmentBindRespVO\> |
| GET | /mes/equipment/bind/order/list | 获取订单设备列表 | orderId (Long) | List\<EquipmentBindRespVO\> |
| PUT | /mes/equipment/bind/status | 更新绑定状态 | EquipmentBindStatusUpdateReqVO | Boolean |
| GET | /mes/equipment/bind/history | 获取绑定历史 | equipmentId (Long) | List\<EquipmentBindHistoryRespVO\> |

---

## 22. 业务功能说明（补充）

### 22.1 工艺路线管理

**功能描述**：
工艺路线管理用于定义产品生产所经过的加工工序序列，是生产执行的基础数据。

**主要功能**：
- 工艺路线定义：创建、编辑、复制工艺路线
- 工序管理：定义工序顺序、工时、设备需求
- 版本控制：支持工艺路线多版本管理
- 图形化配置：可视化工艺路线流程配置

**流程说明**：
```
工艺路线定义 → 工序配置 → 版本发布 → 日计划引用 → 工单生成
```

### 22.2 BOM管理

**功能描述**：
BOM（Bill of Materials）物料清单管理，定义产品生产所需物料及其用量。

**主要功能**：
- BOM定义：创建产品物料清单
- 版本管理：BOM多版本控制与生效日期
- 物料查询：按产品/工序查询物料需求
- 齐套检查：物料库存与需求对比

**流程说明**：
```
BOM定义 → 版本配置 → 日计划引用 → 物料需求计算 → 齐套检查
```

### 22.3 工位配置

**功能描述**：
工位配置管理生产现场的工作站点，包括设备、人员、产能等资源配置。

**主要功能**：
- 工位定义：创建工位并关联产线、车间
- 能力配置：配置工位加工能力
- 资源关联：绑定设备、人员到工位
- 状态监控：实时工位生产状态

### 22.4 人员班组

**功能描述**：
人员班组管理生产现场的人员组织结构，支持班组划分、排班管理。

**主要功能**：
- 班组管理：创建、编辑班组
- 人员配置：配置班组成员及角色
- 排班管理：制定人员工作排班
- 考勤记录：记录人员出勤情况

### 22.5 设备绑定

**功能描述**：
设备绑定管理设备与工位、订单的关联关系，支持设备资源调度。

**主要功能**：
- 工位绑定：设备与工位关联
- 订单绑定：设备与生产订单关联
- 状态管理：设备使用状态跟踪
- 占用查询：查询设备占用情况

### 22.6 齐套检查

**功能描述**：
齐套检查验证日计划发布前所需的物料、设备、人员、工位等资源是否就绪。

**检查维度**：
| 维度 | 检查内容 | 异常处理 |
|------|----------|----------|
| 物料齐套 | BOM物料库存是否满足需求 | 提示缺料明细 |
| 设备齐套 | 工序设备是否可用 | 提示设备冲突 |
| 人员齐套 | 工序人员是否配置 | 提示人员缺口 |
| 工位齐套 | 工序工位是否可用 | 提示工位占用 |

**检查时机**：
- 日计划发布前强制检查
- 可手动触发齐套检查
- 生成齐套检查报告

---

## 23. 文档状态更新

根据本补充文档，以下功能的设计已补充完成：

| 功能模块 | 功能点 | 状态 | 备注 |
|----------|--------|------|------|
| MES生产执行 | 日计划BOM管理 | 📋设计已补充 | 新增14章 |
| MES生产执行 | 日计划工艺路线管理 | 📋设计已补充 | 新增15章 |
| MES生产执行 | 日计划资源配置管理 | 📋设计已补充 | 新增16章 |
| MES生产执行 | 工序排程明细管理 | 📋设计已补充 | 新增17章 |
| MES生产执行 | 齐套检查API | 📋设计已补充 | 新增18章 |
| MES生产执行 | 工位管理API | 📋设计已补充 | 新增19章 |
| MES生产执行 | 人员班组API | 📋设计已补充 | 新增20章 |
| MES生产执行 | 设备绑定API | 📋设计已补充 | 新增21章 |
| MES生产执行 | 工艺路线管理说明 | 📋设计已补充 | 新增22.1章 |
| MES生产执行 | BOM管理说明 | 📋设计已补充 | 新增22.2章 |
| MES生产执行 | 工位配置说明 | 📋设计已补充 | 新增22.3章 |
| MES生产执行 | 人员班组说明 | 📋设计已补充 | 新增22.4章 |
| MES生产执行 | 设备绑定说明 | 📋设计已补充 | 新增22.5章 |
| MES生产执行 | 齐套检查说明 | 📋设计已补充 | 新增22.6章 |

---

> 文档生成时间: 2026-04-17
> 补充内容: DDL表结构、API接口设计、业务功能说明
> 统计口径: 所有标注 `@Tag` 的Controller类，统计其中所有标注 `@GetMapping/@PostMapping/@PutMapping/@DeleteMapping` 的方法

---

## 24. 补充内容核对确认（2026-04-17）

本章节确认以下用户要求的补充内容均已完整收录于本文档中。

### 24.1 DDL表结构核对

| 序号 | 表名 | 所在章节 | 状态 |
|------|------|---------|------|
| 1 | plan_mes_order_day_bom | 第14.1节 | ✅ 已收录 |
| 2 | plan_mes_order_day_bom_detail | 第14.2节 | ✅ 已收录 |
| 3 | plan_mes_order_day_route | 第15.1节 | ✅ 已收录 |
| 4 | plan_mes_order_day_routesub | 第15.2节 | ✅ 已收录 |
| 5 | plan_mes_order_day_equipment | 第16.1节 | ✅ 已收录 |
| 6 | plan_mes_order_day_worker | 第16.2节 | ✅ 已收录 |
| 7 | plan_mes_order_day_workstation | 第16.3节 | ✅ 已收录 |
| 8 | plan_mes_work_scheduling_detail | 第17.1节 | ✅ 已收录 |

### 24.2 API设计核对

| 序号 | API前缀 | 所在章节 | 状态 |
|------|---------|---------|------|
| 1 | /mes/orderday/bom/* | 第14.3节 | ✅ 已收录 |
| 2 | /mes/orderday/route/* | 第15.3节 | ✅ 已收录（含routesub） |
| 3 | /mes/orderday/check/stuffing | 第18.2节 | ✅ 已收录 |
| 4 | /mes/workstation/* | 第19.2节 | ✅ 已收录 |
| 5 | /mes/worker/* | 第16.4节、第20.2节 | ✅ 已收录 |
| 6 | /mes/equipment/bind/* | 第21.2节 | ✅ 已收录 |

### 24.3 业务功能说明核对

| 序号 | 业务功能 | 所在章节 | 状态 |
|------|---------|---------|------|
| 1 | 工艺路线管理 | 第22.1节 | ✅ 已收录 |
| 2 | BOM管理 | 第22.2节 | ✅ 已收录 |
| 3 | 工位配置 | 第22.3节 | ✅ 已收录 |
| 4 | 人员班组 | 第22.4节 | ✅ 已收录 |
| 5 | 设备绑定 | 第22.5节 | ✅ 已收录 |
| 6 | 齐套检查 | 第22.6节 | ✅ 已收录 |

### 24.4 状态表核对（第23章）

文档第23章"文档状态更新"已将上述14项功能全部标记为 📋设计已补充，核对结果：14/14 完成。

**补充核对结论**: 用户要求补充的全部内容已在本文档中完整收录，无需追加新内容。第23章状态表准确反映了所有功能的设计补充状态。

---

## 25. 临时工艺路线管理

### 25.1 plan_mes_order_day_temp_route - 日计划临时工艺路线表

**表用途**：存储日计划的临时工艺路线信息，用于应对紧急工艺调整场景。

**DDL定义**：
```sql
CREATE TABLE `plan_mes_order_day_temp_route` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `order_day_id` bigint NOT NULL COMMENT '日计划ID',
  `temp_route_name` varchar(200) NOT NULL COMMENT '临时工艺路线名称',
  `route_content` longtext COMMENT '工艺路线内容(JSON格式)',
  `reason` varchar(500) DEFAULT NULL COMMENT '变更原因',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态: 0-待审核 1-已审核 2-已驳回',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `tenant_id` bigint DEFAULT '0' COMMENT '租户编号',
  `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `idx_order_day_id` (`order_day_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日计划临时工艺路线表';
```

### 25.2 临时工艺路线API接口

**基础路径**：`/mes/orderday/temp-route`

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/orderday/temp-route/create | 创建临时工艺路线 | TempRouteCreateReqVO | Long |
| PUT | /mes/orderday/temp-route/update | 更新临时工艺路线 | TempRouteUpdateReqVO | Boolean |
| DELETE | /mes/orderday/temp-route/{id} | 删除临时工艺路线 | id (Long) | Boolean |
| GET | /mes/orderday/temp-route/listByOrderDay | 按日计划查询临时工艺路线 | orderDayId (Long) | List\<TempRouteRespVO\> |
| GET | /mes/orderday/temp-route/{id} | 获取临时工艺路线详情 | id (Long) | TempRouteRespVO |
| PUT | /mes/orderday/temp-route/approve | 审核临时工艺路线 | TempRouteApproveReqVO | Boolean |

**TempRouteCreateReqVO**：
```go
type TempRouteCreateReqVO struct {
    OrderDayId    int64  `json:"orderDayId" validate:"required"`
    TempRouteName string `json:"tempRouteName" validate:"required,max=200"`
    RouteContent  string `json:"routeContent"` // JSON格式工艺路线内容
    Reason        string `json:"reason" validate:"max=500"`
}
```

**TempRouteRespVO**：
```go
type TempRouteRespVO struct {
    Id            int64  `json:"id"`
    OrderDayId    int64  `json:"orderDayId"`
    TempRouteName string `json:"tempRouteName"`
    RouteContent  string `json:"routeContent"`
    Reason        string `json:"reason"`
    Status        int    `json:"status"`
    StatusText     string `json:"statusText"`
    Creator       string `json:"creator"`
    CreateTime    string `json:"createTime"`
}
```

---

## 26. 替代BOM管理

### 26.1 plan_mes_order_day_temp_bom - 日计划替代BOM表

**表用途**：存储日计划的替代物料清单信息，用于应对物料短缺等紧急场景。

**DDL定义**：
```sql
CREATE TABLE `plan_mes_order_day_temp_bom` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `order_day_id` bigint NOT NULL COMMENT '日计划ID',
  `temp_bom_name` varchar(200) NOT NULL COMMENT '替代BOM名称',
  `bom_content` longtext COMMENT 'BOM内容(JSON格式)',
  `reason` varchar(500) DEFAULT NULL COMMENT '变更原因',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态: 0-待审核 1-已审核 2-已驳回',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `tenant_id` bigint DEFAULT '0' COMMENT '租户编号',
  `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `idx_order_day_id` (`order_day_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日计划替代BOM表';
```

### 26.2 替代BOM API接口

**基础路径**：`/mes/orderday/temp-bom`

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/orderday/temp-bom/create | 创建替代BOM | TempBomCreateReqVO | Long |
| PUT | /mes/orderday/temp-bom/update | 更新替代BOM | TempBomUpdateReqVO | Boolean |
| DELETE | /mes/orderday/temp-bom/{id} | 删除替代BOM | id (Long) | Boolean |
| GET | /mes/orderday/temp-bom/listByOrderDay | 按日计划查询替代BOM | orderDayId (Long) | List\<TempBomRespVO\> |
| GET | /mes/orderday/temp-bom/{id} | 获取替代BOM详情 | id (Long) | TempBomRespVO |
| PUT | /mes/orderday/temp-bom/approve | 审核替代BOM | TempBomApproveReqVO | Boolean |

**TempBomCreateReqVO**：
```go
type TempBomCreateReqVO struct {
    OrderDayId  int64  `json:"orderDayId" validate:"required"`
    TempBomName string `json:"tempBomName" validate:"required,max=200"`
    BomContent  string `json:"bomContent"` // JSON格式BOM内容
    Reason      string `json:"reason" validate:"max=500"`
}
```

**TempBomRespVO**：
```go
type TempBomRespVO struct {
    Id           int64  `json:"id"`
    OrderDayId   int64  `json:"orderDayId"`
    TempBomName  string `json:"tempBomName"`
    BomContent   string `json:"bomContent"`
    Reason       string `json:"reason"`
    Status       int    `json:"status"`
    StatusText   string `json:"statusText"`
    Creator      string `json:"creator"`
    CreateTime   string `json:"createTime"`
}
```

---

## 27. SOP-PDF文档管理

### 27.1 SOP-PDF管理概述

**功能描述**：
SOP（Standard Operating Procedure）标准作业规程文档管理，用于关联工单与SOP操作文档，支持PDF格式文件的上传、存储和获取。

**主要功能**：
- SOP文档上传：支持PDF格式标准作业文档上传
- 工单关联：SOP文档与工艺路线、工序关联
- 文档获取：按工艺路线或工序获取对应SOP
- 版本管理：支持SOP文档多版本管理

**流程说明**：
```
SOP文档上传 → 工艺路线关联 → 工单生成 → SOP文档获取 → 作业执行
```

### 27.2 SOP-PDF API接口

**基础路径**：`/mes/sop`

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/sop/upload | 上传SOP文档 | SopUploadReqVO (multipart/form-data) | SopRespVO |
| GET | /mes/sop/getPDF | 获取工单关联SOP文档 | orderId (Long) | File/byte[] |
| GET | /mes/sop/listByProcessRoute | 按工艺路线获取SOP列表 | processRouteId (Long) | List\<SopRespVO\> |
| GET | /mes/sop/{id} | 获取SOP详情 | id (Long) | SopRespVO |
| DELETE | /mes/sop/{id} | 删除SOP文档 | id (Long) | Boolean |
| GET | /mes/sop/download/{id} | 下载SOP文档 | id (Long) | File/byte[] |

**SopUploadReqVO**：
```go
type SopUploadReqVO struct {
    File           *multipart.FileHeader `json:"-" form:"file"`        // SOP PDF文件
    SopName        string                `json:"sopName"`               // SOP文档名称
    ProcessRouteId int64                 `json:"processRouteId"`        // 关联工艺路线ID
    WorkOrderId    int64                 `json:"workOrderId"`            // 关联工单ID（可选）
    Version        string                `json:"version"`              // 文档版本
}
```

**SopRespVO**：
```go
type SopRespVO struct {
    Id              int64  `json:"id"`
    SopName         string `json:"sopName"`
    FileName        string `json:"fileName"`
    FileSize        int64  `json:"fileSize"`
    FileUrl         string `json:"fileUrl"`
    ProcessRouteId  int64  `json:"processRouteId"`
    WorkOrderId     int64  `json:"workOrderId"`
    Version         string `json:"version"`
    UploadTime      string `json:"uploadTime"`
    Uploader        string `json:"uploader"`
}
```

---

## 28. 文档状态更新（第二轮）

根据本补充文档，以下功能的设计已补充完成：

| 功能模块 | 功能点 | 状态 | 备注 |
|----------|--------|------|------|
| MES生产执行 | 临时工艺路线管理 | 📋设计已补充 | 新增25章 |
| MES生产执行 | 替代BOM管理 | 📋设计已补充 | 新增26章 |
| MES生产执行 | SOP-PDF文档管理 | 📋设计已补充 | 新增27章 |

---

## 29. 补充内容核对确认（第二轮）

本章节确认以下用户要求的补充内容均已完整收录于本文档中。

### 29.1 DDL表结构核对

| 序号 | 表名 | 所在章节 | 状态 |
|------|------|---------|------|
| 1 | plan_mes_order_day_temp_route | 第25.1节 | ✅ 已收录 |
| 2 | plan_mes_order_day_temp_bom | 第26.1节 | ✅ 已收录 |

### 29.2 API设计核对

| 序号 | API前缀 | 所在章节 | 状态 |
|------|---------|---------|------|
| 1 | /mes/orderday/temp-route/* | 第25.2节 | ✅ 已收录 |
| 2 | /mes/orderday/temp-bom/* | 第26.2节 | ✅ 已收录 |
| 3 | /mes/sop/* | 第27.2节 | ✅ 已收录 |

### 29.3 业务功能说明核对

| 序号 | 业务功能 | 所在章节 | 状态 |
|------|---------|---------|------|
| 1 | 临时工艺路线 | 第25章、第27.1节 | ✅ 已收录 |
| 2 | 替代BOM | 第26章、第27.1节 | ✅ 已收录 |
| 3 | SOP-PDF关联 | 第27章 | ✅ 已收录 |

### 29.4 状态表核对（第28章）

文档第28章已将上述3项功能全部标记为 📋设计已补充，核对结果：3/3 完成。

**补充核对结论**: 用户要求补充的全部内容已在本文档中完整收录（第25-27章），包括2张DDL表、3组API设计、3项业务功能说明。第28章状态表准确反映了所有新增功能的设计补充状态。

---

> 文档更新时间: 2026-04-17（第二轮补充）
> 补充内容: 临时工艺路线DDL+API、替代BOM DDL+API、SOP-PDF API+说明
> 统计口径: 所有标注 `@Tag` 的Controller类，统计其中所有标注 `@GetMapping/@PostMapping/@PutMapping/@DeleteMapping` 的方法
