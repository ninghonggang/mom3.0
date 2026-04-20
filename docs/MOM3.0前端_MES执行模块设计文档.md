# MOM3.0前端_MES执行模块设计文档

> **版本**: V1.1
> **日期**: 2026-04-17
> **项目**: 闻荫科技MOM3.0智能制造执行系统
> **对比版本**: SFMS3.0 MES执行模块 (31页缺失)

---

## 1. 模块概述

### 1.1 功能定位

MES执行模块是MOM3.0的核心模块，负责将生产计划转化为具体的生产任务，完成工单管理、派工报工、完工入库等核心业务流程。

### 1.2 当前状态 vs SFMS3.0

| 类别 | MOM3.0现状 | SFMS3.0完整功能 |
|------|-----------|----------------|
| 页面数量 | 不完整 | 31页完整实现 |
| 生产日计划 | 有基础 | 完整BOM配置、工艺路线、人员设备配置 |
| 工单排程 | 有基础 | 高级排程、工序明细、报工 |
| BOM管理 | 缺失 | 完整BOM配置、版本管理 |
| 工位管理 | 缺失 | 工位档案、能力配置 |
| 工序管理 | 缺失 | 工序详细配置、类型管理 |
| 产品档案 | 缺失 | 产品BOM、配置管理 |
| 物料申请 | 缺失 | 生产物料申请、物料需求 |
| 拆解管理 | 缺失 | 拆解任务、拆解记录 |
| 返工批次 | 缺失 | 返工管理、返工记录 |
| 月计划 | 缺失 | 月度计划、月计划维护 |
| 工作日历 | 缺失 | 日历配置、工作时间 |
| 产品报交 | 缺失 | 报交管理、报交记录 |
| 产品返线 | 缺失 | 返线管理、返线记录 |
| 排程详情 | 缺失 | 排程明细、优先级 |
| 能力档案 | 缺失 | 能力配置、能力计算 |

---

## 2. 缺失页面详细设计

### 2.1 生产日计划 (`/mes/orderDay`)

**路径**: `/mes/orderDay`
**组件**: `OrderDayList.vue`
**功能**: 日计划创建/发布/终止/恢复

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| planNoDay | String | 计划单号 |
| productCode | String | 产品图号 |
| productName | String | 产品名称 |
| workroomCode | String | 车间编码 |
| lineCode | String | 产线编码 |
| planCount | Integer | 计划数量 |
| planDate | Date | 计划日期 |
| processrouteCode | String | 工艺路线编码 |
| workMode | String | 工单模式(BATCH批量/SINGLE单件) |
| taskMode | String | 生产模式(ASSIGN派工/RECIEVE领工) |
| status | String | 状态(WAITSECHUDLE待排产/PUBLISHED已发布/PROCESSING进行中/TERMINALE已终止) |
| completeInspect | String | 齐套检查状态(COMPLETE/INCOMPLETE/PENDING) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /mes/orderday/create | 创建生产日计划 |
| PUT | /mes/orderday/update | 更新生产日计划 |
| DELETE | /mes/orderday/delete | 删除生产日计划 |
| GET | /mes/orderday/get | 获取单个日计划 |
| GET | /mes/orderday/page | 日计划分页查询 |
| POST | /mes/orderday/publishPlan | 发布日计划生成工单 |
| POST | /mes/orderday/stopPlan/{id} | 终止已发布计划 |
| POST | /mes/orderday/resumePlan/{id} | 恢复已终止计划 |
| GET | /mes/orderday/getProcessInfo | 获取工艺路线及工序信息 |
| GET | /mes/orderday/getBomInfo | 获取配置的BOM信息 |
| GET | /mes/orderday/export-excel | 导出日计划Excel |
| POST | /mes/orderday/import | 导入日计划 |

---

### 2.2 工单排程 (`/mes/workScheduling`)

**路径**: `/mes/workScheduling`
**组件**: `WorkSchedulingList.vue`
**功能**: 排程管理、工单调度

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| schedulingCode | String | 工单编码 |
| planMasterCode | String | 计划单号(来源日计划) |
| productCode | String | 产品图号 |
| productName | String | 产品名称 |
| workroomCode | String | 车间编码 |
| lineCode | String | 产线编码 |
| planQty | Decimal | 计划数量 |
| doneQty | Decimal | 已完成数量 |
| qualifiedQty | Decimal | 合格数量 |
| workMode | String | 工单模式(BATCH批量/SINGLE单件) |
| taskMode | String | 生产模式(ASSIGN派工/RECIEVE领工) |
| processrouteCode | String | 工艺路线编码 |
| bomCode | String | BOM编码 |
| planStartDate | Date | 计划开始日期 |
| planEndDate | Date | 计划结束日期 |
| actualStartDate | DateTime | 实际开始时间 |
| actualEndDate | DateTime | 实际结束时间 |
| workingStatus | String | 作业状态(READYTODO待报工/WAITBEGIN待开工/PROCESSING进行中/FINISH已完成/TERMINALE已终止) |
| priority | Integer | 优先级(1-10) |
| batchCode | String | 批次码 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /mes/workScheduling/create | 创建生产任务排产 |
| PUT | /mes/workScheduling/update | 更新生产任务排产 |
| POST | /mes/workScheduling/update-status | 更新生产任务排产状态 |
| DELETE | /mes/workScheduling/delete | 删除生产任务排产 |
| GET | /mes/workScheduling/get | 获取生产任务排产 |
| GET | /mes/workScheduling/page | 获取生产任务排产分页 |
| GET | /mes/workScheduling/PDA-page | 获取生产任务排产分页(PDA) |
| POST | /mes/workScheduling/senior | 高级搜索生产任务排产 |
| POST | /mes/workScheduling/completeHandle | 完工处理 |
| POST | /mes/workScheduling/reportForAll | 批量报工处理 |
| GET | /mes/workScheduling/getNodeInfo | 获取当前工序基本信息 |
| GET | /mes/workScheduling/getProcessList | 获取工单所有工序信息 |
| GET | /mes/workScheduling/get-PDF | SOP-获取PDF |

---

### 2.3 BOM管理 (`/mes/bomMgmt`)

**路径**: `/mes/bomMgmt`
**组件**: `BomMgmtList.vue`
**功能**: BOM配置、BOM版本

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| bomCode | String | BOM编码 |
| bomVersion | String | BOM版本 |
| productCode | String | 产品图号 |
| itemCode | String | 物料编码 |
| itemName | String | 物料名称 |
| itemSpec | String | 物料规格 |
| itemUnit | String | 计量单位 |
| itemType | String | 物料类型(原材料/成品/半成品) |
| bomLevel | Integer | BOM层级 |
| parentItemCode | String | 父物料编码 |
| standardQty | Decimal | 标准用量 |
| lossRate | Decimal | 损耗率 |
| isKey | Boolean | 是否关键件 |
| status | String | 状态(ACTIVE激活/INACTIVE未激活) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mes/bom/page | BOM分页查询 |
| GET | /mes/bom/get | 获取BOM详情 |
| POST | /mes/bom/create | 创建BOM |
| PUT | /mes/bom/update | 更新BOM |
| DELETE | /mes/bom/delete | 删除BOM |
| GET | /mes/bom/list | BOM列表 |
| GET | /mes/bom/version/{bomCode} | 获取BOM版本列表 |
| POST | /mes/bom/copy | 复制BOM |
| GET | /mes/bom/tree | 获取BOM树形结构 |
| GET | /mes/common/geBomTreeByProductCode | 根据产品编码获取BOM树形结构 |
| GET | /mes/common/getBomListByProductCode | 根据产品编码获取BOM列表 |
| GET | /mes/common/getBomListByProductAndProcess | 根据产品编码获取工序物料 |

---

### 2.4 工位管理 (`/mes/workstationMgmt`)

**路径**: `/mes/workstationMgmt`
**组件**: `WorkstationMgmtList.vue`
**功能**: 工位档案、工位能力

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| workstationCode | String | 工位编码 |
| workstationName | String | 工位名称 |
| workstationType | String | 工位类型 |
| workroomCode | String | 车间编码 |
| lineCode | String | 产线编码 |
| capacity | Integer | 产能(件/日) |
| abilityValue | Decimal | 能力值 |
| abilityUnit | String | 能力单位 |
| status | String | 状态(ACTIVE激活/INACTIVE未激活) |
| location | String | 位置描述 |
| remark | String | 备注 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mes/workstation/page | 工位分页查询 |
| GET | /mes/workstation/get | 获取工位详情 |
| POST | /mes/workstation/create | 创建工位 |
| PUT | /mes/workstation/update | 更新工位 |
| DELETE | /mes/workstation/delete | 删除工位 |
| GET | /mes/workstation/list | 工位列表 |
| PUT | /mes/workstation/ability | 更新工位能力 |
| GET | /mes/workstation/ability/{code} | 获取工位能力 |

---

### 2.5 工序管理 (`/mes/processMgmt`)

**路径**: `/mes/processMgmt`
**组件**: `ProcessMgmtList.vue`
**功能**: 工序详细配置、工序类型

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| processCode | String | 工序编码 |
| processName | String | 工序名称 |
| processType | String | 工序类型 |
| stdWorkTime | Decimal | 标准工时(分钟) |
| stdCapacity | Integer | 标准产能(件/日) |
| qualityControlType | String | 质检类型(免检/自检/互检/专检) |
| requireInspect | Boolean | 是否需要质检 |
| inspectFormCode | String | 质检表单编码 |
| toolRequired | String | 需要的工具/模具 |
| materialFeedType | String | 投料方式(齐套/顺序) |
| collectType | String | 采集方式 |
| isKeyProcess | Boolean | 是否关键工序 |
| isAutoComplete | Boolean | 是否自动完工 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mes/process/page | 工序分页查询 |
| GET | /mes/process/get | 获取工序详情 |
| POST | /mes/process/create | 创建工序 |
| PUT | /mes/process/update | 更新工序 |
| DELETE | /mes/process/delete | 删除工序 |
| GET | /mes/process/list | 工序列表 |
| POST | /mes/process/senior | 高级搜索工序 |

---

### 2.6 产品档案 (`/mes/productArchive`)

**路径**: `/mes/productArchive`
**组件**: `ProductArchiveList.vue`
**功能**: 产品BOM、产品配置

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| productCode | String | 产品图号 |
| productName | String | 产品名称 |
| itemCode | String | 物料编码 |
| itemName | String | 物料名称 |
| itemSpec | String | 物料规格 |
| itemUnit | String | 计量单位 |
| bomCode | String | BOM编码 |
| bomVersion | String | BOM版本 |
| configItems | JSON | 配置项JSON |
| status | String | 状态 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mes/product/page | 产品分页查询 |
| GET | /mes/product/get | 获取产品详情 |
| POST | /mes/product/create | 创建产品 |
| PUT | /mes/product/update | 更新产品 |
| DELETE | /mes/product/delete | 删除产品 |
| GET | /mes/product/list | 产品列表 |
| GET | /mes/product/bom/{code} | 获取产品BOM |
| GET | /mes/product/config/{code} | 获取产品配置 |
| PUT | /mes/product/config | 更新产品配置 |

---

### 2.7 物料申请 (`/mes/materialApply`)

**路径**: `/mes/materialApply`
**组件**: `MaterialApplyList.vue`
**功能**: 生产物料申请、物料需求

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| requestNo | String | 申请单号 |
| requestType | String | 申请类型(叫料/补料) |
| workOrderCode | String | 工单编号 |
| productCode | String | 产品图号 |
| workroomCode | String | 车间编码 |
| lineCode | String | 产线编码 |
| planDate | Date | 计划日期 |
| status | String | 状态 |
| applicant | String | 申请人 |
| applyTime | DateTime | 申请时间 |

**物料明细字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| requestNo | String | 申请单号 |
| itemCode | String | 物料编码 |
| itemName | String | 物料名称 |
| itemSpec | String | 物料规格 |
| requestQty | Decimal | 申请数量 |
| deliveredQty | Decimal | 已配送数量 |
| receivedQty | Decimal | 已领用数量 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mes/material-apply/page | 物料申请分页 |
| GET | /mes/material-apply/get | 获取物料申请详情 |
| POST | /mes/material-apply/create | 创建物料申请 |
| PUT | /mes/material-apply/update | 更新物料申请 |
| DELETE | /mes/material-apply/delete | 删除物料申请 |
| POST | /mes/material-apply/submit | 提交申请 |
| POST | /mes/material-apply/approve | 审批申请 |
| POST | /mes/material-apply/reject | 驳回申请 |
| GET | /mes/material-apply/detail/{requestNo} | 获取申请明细 |

---

### 2.8 拆解管理 (`/mes/disassembly`)

**路径**: `/mes/disassembly`
**组件**: `DisassemblyList.vue`
**功能**: 拆解任务、拆解记录

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| dismantlingNo | String | 拆解单号 |
| productCode | String | 产品图号 |
| productName | String | 产品名称 |
| dismantlingType | String | 拆解类型(报废/回收) |
| quantity | Integer | 拆解数量 |
| workroomCode | String | 车间编码 |
| reason | String | 拆解原因 |
| status | String | 状态 |
| operator | String | 操作人 |
| operateTime | DateTime | 操作时间 |

**拆解明细字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| dismantlingNo | String | 拆解单号 |
| itemCode | String | 物料编码 |
| itemName | String | 物料名称 |
| itemSpec | String | 物料规格 |
| quantity | Decimal | 拆解数量 |
|回收数量 | Decimal | 回收数量 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mes/disassembly/page | 拆解分页查询 |
| GET | /mes/disassembly/get | 获取拆解详情 |
| POST | /mes/disassembly/create | 创建拆解单 |
| PUT | /mes/disassembly/update | 更新拆解单 |
| DELETE | /mes/disassembly/delete | 删除拆解单 |
| POST | /mes/disassembly/submit | 提交拆解 |
| POST | /mes/disassembly/complete | 完成拆解 |
| GET | /mes/disassembly/detail/{dismantlingNo} | 获取拆解明细 |

---

### 2.9 返工批次 (`/mes/reworkBatch`)

**路径**: `/mes/reworkBatch`
**组件**: `ReworkBatchList.vue`
**功能**: 返工管理、返工记录

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| reworkNo | String | 返工单号 |
| batchCode | String | 批次码 |
| productCode | String | 产品图号 |
| productName | String | 产品名称 |
| quantity | Integer | 返工数量 |
| reworkType | String | 返工类型 |
| reason | String | 返工原因 |
| workroomCode | String | 车间编码 |
| lineCode | String | 产线编码 |
| status | String | 状态(PENDING待领取/RECEIVED已领取/PROCESSING进行中/COMPLETED已完成/SUSPENDED已中止) |
| reworkProcessCode | String | 返工工序编码 |
| reworkProcessName | String | 返工工序名称 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mes/rework-batch/page | 返工批次分页 |
| GET | /mes/rework-batch/get | 获取返工批次详情 |
| POST | /mes/rework-batch/create | 创建返工批次 |
| PUT | /mes/rework-batch/update | 更新返工批次 |
| DELETE | /mes/rework-batch/delete | 删除返工批次 |
| DELETE | /mes/rework-batch/suspend | 中止返工批次 |
| PUT | /mes/rework-batch/receive | 领取返工批次 |
| PUT | /mes/rework-batch/finish | 完成返工批次 |
| POST | /mes/rework-batch/senior | 高级搜索返工批次 |

---

### 2.10 月计划 (`/mes/monthPlan`)

**路径**: `/mes/monthPlan`
**组件**: `MonthPlanList.vue`
**功能**: 月度计划、月计划维护

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| monthPlanNo | String | 月计划单号 |
| month | String | 计划月份(格式: yyyy-MM) |
| productCode | String | 产品图号 |
| productName | String | 产品名称 |
| totalQty | Integer | 计划总量 |
| deliveredQty | Integer | 已排产量 |
| status | String | 状态 |

**月计划子表字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| monthPlanNo | String | 月计划单号 |
| productCode | String | 产品图号 |
| day01-day31 | Integer | 每日分配数量 |
| actualDay01-actualDay31 | Integer | 每日实际完成数量 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mes/month-plan/page | 月计划分页 |
| GET | /mes/month-plan/get | 获取月计划详情 |
| POST | /mes/month-plan/create | 创建月计划 |
| PUT | /mes/month-plan/update | 更新月计划 |
| DELETE | /mes/month-plan/delete | 删除月计划 |
| POST | /mes/month-plan/breakdown | 拆解为日计划 |
| GET | /mes/month-plan/sub/{monthPlanNo} | 获取月计划子表 |
| PUT | /mes/month-plan/sub/update | 更新月计划子表 |

---

### 2.11 工作日历 (`/mes/workCalendar`)

**路径**: `/mes/workCalendar`
**组件**: `WorkCalendarList.vue`
**功能**: 日历配置、工作时间

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| calendarCode | String | 日历编码 |
| teamCode | String | 班组编码 |
| teamName | String | 班组名称 |
| workDate | Date | 工作日期 |
| workStartTime | Time | 工作开始时间 |
| workEndTime | Time | 工作结束时间 |
| isWorkDay | Boolean | 是否工作日 |
| shiftType | String | 班次类型(白班/夜班) |
| year | String | 年份 |
| keyDate | String | 日期Key(格式: yyyy-MM-dd) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mes/work-calendar/page | 工作日历分页 |
| GET | /mes/work-calendar/get | 获取工作日历详情 |
| POST | /mes/work-calendar/create | 创建工作日历 |
| POST | /mes/work-calendar/createBatch | 批量创建工作日历 |
| PUT | /mes/work-calendar/update | 更新工作日历 |
| DELETE | /mes/work-calendar/delete | 删除工作日历 |
| POST | /mes/work-calendar/deleteTeam | 删除班组日历 |
| GET | /mes/work-calendar/list | 获取班组日历列表 |
| GET | /mes/work-calendar/export-excel | 导出工作日历Excel |
| POST | /mes/work-calendar/import | 导入工作日历 |

---

### 2.12 产品报交 (`/mes/productReport`)

**路径**: `/mes/productReport`
**组件**: `ProductReportList.vue`
**功能**: 报交管理、报交记录

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| reportNo | String | 报交单号 |
| workOrderCode | String | 工单编号 |
| productCode | String | 产品图号 |
| productName | String | 产品名称 |
| reportQty | Integer | 报交数量 |
| qualifiedQty | Integer | 合格数量 |
| rejectQty | Integer | 不合格数量 |
| reportType | String | 报交类型(首报/续报/末报) |
| reportTime | DateTime | 报交时间 |
| inspector | String | 检验员 |
| status | String | 状态 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mes/product-report/page | 产品报交分页 |
| GET | /mes/product-report/get | 获取产品报交详情 |
| POST | /mes/product-report/create | 创建产品报交 |
| PUT | /mes/product-report/update | 更新产品报交 |
| DELETE | /mes/product-report/delete | 删除产品报交 |
| POST | /mes/product-report/submit | 提交报交 |
| GET | /mes/product-report/history/{workOrderCode} | 获取报交历史 |

---

### 2.13 产品返线 (`/mes/returnLine`)

**路径**: `/mes/returnLine`
**组件**: `ReturnLineList.vue`
**功能**: 返线管理、返线记录

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| returnNo | String | 返线单号 |
| workOrderCode | String | 工单编号 |
| productCode | String | 产品图号 |
| productName | String | 产品名称 |
| returnQty | Integer | 返线数量 |
| returnReason | String | 返线原因 |
| returnType | String | 返线类型(质量问题/设备故障/其他) |
| fromProcessCode | String | 来源工序编码 |
| toProcessCode | String | 目标工序编码 |
| status | String | 状态 |
| applicant | String | 申请人 |
| applyTime | DateTime | 申请时间 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mes/return-line/page | 产品返线分页 |
| GET | /mes/return-line/get | 获取产品返线详情 |
| POST | /mes/return-line/create | 创建产品返线 |
| PUT | /mes/return-line/update | 更新产品返线 |
| DELETE | /mes/return-line/delete | 删除产品返线 |
| POST | /mes/return-line/submit | 提交返线申请 |
| POST | /mes/return-line/approve | 审批返线申请 |
| POST | /mes/return-line/reject | 驳回返线申请 |

---

### 2.14 排程详情 (`/mes/schedulingDetail`)

**路径**: `/mes/schedulingDetail`
**组件**: `SchedulingDetailList.vue`
**功能**: 排程明细、优先级

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| detailId | String | 明细ID |
| schedulingCode | String | 工单编码 |
| processCode | String | 工序编码 |
| processName | String | 工序名称 |
| seqNo | Integer | 工序顺序 |
| stationCode | String | 工位编码 |
| workstationCode | String | 工位作业地 |
| workGroupCode | String | 班组编码 |
| workerCode | String | 作业人员编码 |
| equipmentCode | String | 设备编码 |
| planQty | Decimal | 计划数量 |
| taskQty | Decimal | 任务数量 |
| doneQty | Decimal | 完成数量 |
| priority | Integer | 优先级(1-10) |
| taskStatus | String | 任务状态(WAIT待领/PENDING待报/PROCESSING进行中/COMPLETED已完成) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mes/scheduling-detail/page | 排程明细分页 |
| GET | /mes/scheduling-detail/get | 获取排程明细详情 |
| POST | /mes/scheduling-detail/create | 创建排程明细 |
| PUT | /mes/scheduling-detail/update | 更新排程明细 |
| DELETE | /mes/scheduling-detail/delete | 删除排程明细 |
| PUT | /mes/scheduling-detail/priority | 更新优先级 |
| PUT | /mes/scheduling-detail/status | 更新状态 |
| GET | /mes/scheduling-detail/list/{schedulingCode} | 获取工单所有明细 |

---

### 2.15 能力档案 (`/mes/ability`)

**路径**: `/mes/ability`
**组件**: `AbilityList.vue`
**功能**: 能力配置、能力计算

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| abilityCode | String | 能力编号 |
| abilityName | String | 能力名称 |
| abilityType | String | 能力类型(设备/人员/工位) |
| ownerCode | String | 所属者编码 |
| ownerName | String | 所属者名称 |
| abilityValue | Decimal | 能力值 |
| abilityUnit | String | 能力单位 |
| workTime | Decimal | 日工作时长(小时) |
| workDays | Integer | 月工作天数 |
| efficiency | Decimal | 效率系数 |
| status | String | 状态 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mes/ability/page | 能力分页查询 |
| GET | /mes/ability/get | 获取能力详情 |
| POST | /mes/ability/create | 创建能力 |
| PUT | /mes/ability/update | 更新能力 |
| DELETE | /mes/ability/delete | 删除能力 |
| GET | /mes/ability/list | 能力列表 |
| GET | /mes/ability/calculate | 计算能力 |
| PUT | /mes/ability/calibrate | 校准能力 |

---

## 3. 文件结构

```
mom-web/src/views/mes/
├── orderDay/
│   ├── OrderDayList.vue      # 日计划列表
│   └── OrderDayDetail.vue    # 日计划详情
├── workScheduling/
│   ├── WorkSchedulingList.vue # 工单排程列表
│   └── WorkSchedulingDetail.vue # 工单排程详情
├── bomMgmt/
│   ├── BomMgmtList.vue        # BOM管理列表
│   ├── BomMgmtEdit.vue        # BOM编辑
│   └── BomVersion.vue         # BOM版本
├── workstationMgmt/
│   ├── WorkstationList.vue    # 工位列表
│   └── WorkstationEdit.vue    # 工位编辑
├── processMgmt/
│   ├── ProcessMgmtList.vue    # 工序列表
│   └── ProcessMgmtEdit.vue    # 工序编辑
├── productArchive/
│   ├── ProductArchiveList.vue # 产品档案列表
│   └── ProductArchiveEdit.vue # 产品档案编辑
├── materialApply/
│   ├── MaterialApplyList.vue  # 物料申请列表
│   └── MaterialApplyDetail.vue # 物料申请详情
├── disassembly/
│   ├── DisassemblyList.vue    # 拆解列表
│   └── DisassemblyDetail.vue  # 拆解详情
├── reworkBatch/
│   ├── ReworkBatchList.vue    # 返工批次列表
│   └── ReworkBatchDetail.vue  # 返工批次详情
├── monthPlan/
│   ├── MonthPlanList.vue      # 月计划列表
│   └── MonthPlanDetail.vue    # 月计划详情
├── workCalendar/
│   ├── WorkCalendarList.vue   # 工作日历列表
│   └── WorkCalendarEdit.vue   # 工作日历编辑
├── productReport/
│   ├── ProductReportList.vue  # 产品报交列表
│   └── ProductReportDetail.vue # 产品报交详情
├── returnLine/
│   ├── ReturnLineList.vue     # 产品返线列表
│   └── ReturnLineDetail.vue  # 产品返线详情
├── schedulingDetail/
│   ├── SchedulingDetailList.vue # 排程明细列表
│   └── SchedulingDetailEdit.vue # 排程明细编辑
└── ability/
    ├── AbilityList.vue        # 能力档案列表
    └── AbilityEdit.vue        # 能力档案编辑
```

---

## 4. 页面清单

| 序号 | 页面 | 路由路径 | 功能说明 |
|------|------|----------|----------|
| 1 | 生产日计划 | `/mes/orderDay` | 日计划创建/发布/终止/恢复 |
| 2 | 工单排程 | `/mes/workScheduling` | 排程管理、工单调度 |
| 3 | BOM管理 | `/mes/bomMgmt` | BOM配置、BOM版本 |
| 4 | 工位管理 | `/mes/workstationMgmt` | 工位档案、工位能力 |
| 5 | 工序管理 | `/mes/processMgmt` | 工序详细配置、工序类型 |
| 6 | 产品档案 | `/mes/productArchive` | 产品BOM、产品配置 |
| 7 | 物料申请 | `/mes/materialApply` | 生产物料申请、物料需求 |
| 8 | 拆解管理 | `/mes/disassembly` | 拆解任务、拆解记录 |
| 9 | 返工批次 | `/mes/reworkBatch` | 返工管理、返工记录 |
| 10 | 月计划 | `/mes/monthPlan` | 月度计划、月计划维护 |
| 11 | 工作日历 | `/mes/workCalendar` | 日历配置、工作时间 |
| 12 | 产品报交 | `/mes/productReport` | 报交管理、报交记录 |
| 13 | 产品返线 | `/mes/returnLine` | 返线管理、返线记录 |
| 14 | 排程详情 | `/mes/schedulingDetail` | 排程明细、优先级 |
| 15 | 能力档案 | `/mes/ability` | 能力配置、能力计算 |

---

## 5. 实现优先级

### 高优先级 (MVP必备)
1. 生产日计划 - 核心业务入口
2. 工单排程 - 生产任务核心
3. BOM管理 - 物料清单基础
4. 工序管理 - 生产流程定义

### 中优先级 (完整流程)
5. 工位管理 - 资源分配
6. 物料申请 - 物料管理
7. 排程详情 - 工序级管理
8. 产品报交 - 报交记录

### 低优先级 (增强功能)
9. 产品档案 - 产品配置
10. 拆解管理 - 拆解记录
11. 返工批次 - 返工管理
12. 月计划 - 计划分解
13. 工作日历 - 日历配置
14. 产品返线 - 返线管理
15. 能力档案 - 能力配置

---

*文档版本: V1.1 | 对比版本: SFMS3.0 | 最后更新: 2026-04-17*
