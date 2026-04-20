# MOM3.0 MES生产执行模块设计文档

**版本**: V2.0 | **所属模块**: M03生产执行 | **基于**: [MOM3.0_主设计文档](./MOM3.0_主设计文档.md)

---

## 1. 模块概述

### 1.1 功能定位

MES生产执行模块是MOM3.0的核心模块，负责将生产计划转化为具体的生产任务，完成工单管理、派工报工、完工入库等核心业务流程。

### 1.2 核心功能

| 功能 | 说明 |
|------|------|
| 销售订单 | 接收销售订单，转化为生产依据 |
| 月计划 | 月度生产计划分解 |
| 日计划 | 日生产工单生成 |
| 派工管理 | 生产派工单管理 |
| 生产报工 | 工序报工、完工入库 |
| 完工入库 | 生产入库闭环 |

---

## 2. 页面清单

| 页面 | 路由路径 | 核心功能 |
|------|----------|----------|
| 销售订单 | `/production/sales-order` | 订单列表、确认、发货 |
| 生产工单 | `/production/order` | 工单CRUD、状态管理 |
| 派工管理 | `/production/dispatch` | 派工单、报工 |
| 生产报工 | `/production/report` | 报工、完工入库 |
| 生产看板 | `/production/kanban` | 车间看板可视化 |
| 月计划 | - | 月度计划分解（无独立页面） |
| 日计划 | - | 日计划生成（无独立页面） |

---

## 3. UI设计规范

### 3.1 页面基本结构

```vue
<template>
  <div class="page-container">
    <!-- 搜索区域 -->
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="工单号">
          <el-input v-model="searchForm.orderNo" placeholder="请输入" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待生产" value="PENDING" />
            <el-option label="生产中" value="RUNNING" />
            <el-option label="已完成" value="COMPLETED" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">
            <el-icon><Search /></el-icon>查询
          </el-button>
          <el-button @click="handleReset">
            <el-icon><Refresh /></el-icon>重置
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 工具栏 -->
    <el-card class="toolbar-card">
      <div class="toolbar-left">
        <el-button type="primary" @click="handleAdd">
          <el-icon><Plus /></el-icon>新增
        </el-button>
      </div>
      <div class="toolbar-right">
        <el-button text @click="searchCollapsed = !searchCollapsed">
          {{ searchCollapsed ? '展开搜索' : '收起搜索' }}
        </el-button>
      </div>
    </el-card>

    <!-- 表格 -->
    <el-card class="table-card">
      <el-table :data="tableData" v-loading="loading" stripe border>
        <el-table-column type="selection" width="55" />
        <el-table-column prop="order_no" label="工单号" min-width="120" />
        <el-table-column prop="material_name" label="产品" min-width="150" show-overflow-tooltip />
        <el-table-column prop="quantity" label="数量" width="100" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ row.statusText }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="plan_start_date" label="计划开始" width="120" />
        <el-table-column prop="plan_end_date" label="计划结束" width="120" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.pageSize"
          :total="pagination.total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
        />
      </div>
    </el-card>

    <!-- 编辑弹窗 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="工单号" prop="orderNo">
              <el-input v-model="formData.orderNo" :disabled="isEdit" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="产品" prop="materialId">
              <el-select v-model="formData.materialId" placeholder="请选择">
                <el-option label="产品A" :value="1" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>
```

### 3.2 状态映射

```js
const statusMap = {
  'PENDING': { type: 'warning', text: '待生产' },
  'RUNNING': { type: 'primary', text: '生产中' },
  'COMPLETED': { type: 'success', text: '已完成' },
  'CANCELLED': { type: 'info', text: '已取消' },
}

const getStatusType = (status) => statusMap[status]?.type || 'info'
```

---

## 4. 业务流程

### 4.1 工单完整流程

```
销售订单 → 月计划分解 → 日计划生成 → 工单创建 → 派工 → 报工 → 完工入库
```

### 4.2 状态机

| 状态 | 可执行操作 | 下一状态 |
|------|-----------|----------|
| PENDING（待生产） | 开始生产、取消 | RUNNING, CANCELLED |
| RUNNING（生产中） | 暂停、完工 | PAUSED, COMPLETED |
| PAUSED（已暂停） | 恢复生产、取消 | RUNNING, CANCELLED |
| COMPLETED（已完成） | - | - |
| CANCELLED（已取消） | - | - |

---

## 5. 数据模型

### 5.1 工单

| 字段 | 类型 | 说明 |
|------|------|------|
| order_no | VARCHAR(50) | 工单号 |
| material_id | BIGINT | 产品ID |
| quantity | DECIMAL(18,3) | 数量 |
| plan_start_date | DATE | 计划开始日期 |
| plan_end_date | DATE | 计划结束日期 |
| status | VARCHAR(20) | 状态 |
| workshop_id | BIGINT | 车间ID |
| line_id | BIGINT | 产线ID |

---

## 6. API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /production/order/list | 工单列表 |
| GET | /production/order/:id | 工单详情 |
| POST | /production/order | 创建工单 |
| PUT | /production/order/:id | 更新工单 |
| DELETE | /production/order/:id | 删除工单 |
| PUT | /production/order/:id/start | 开始生产 |
| PUT | /production/order/:id/complete | 完工 |

---

## 7. 关联文档

- [MOM3.0_主设计文档](./MOM3.0_主设计文档.md) - 系统总览
- [MOM3.0_UI设计规范](./MOM3.0_UI设计规范.md) - UI规范详情

---

## §11 生产日计划管理 (orderDay)

### 11.1 功能说明

生产日计划是MES系统的核心业务入口，负责日生产计划的创建、发布、终止、恢复等全生命周期管理。

### 11.2 页面路径

| 页面 | 路由路径 | 说明 |
|------|----------|------|
| 生产日计划 | `/mes/orderDay` | 日计划主页面 |

### 11.3 核心字段

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

### 11.4 API接口

| 方法 | 路径 | 说明 |
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

## §12 工单排程管理 (workScheduling)

### 12.1 功能说明

工单排程是生产任务的核心载体，承载工单生成、工序调度、报工等核心生产执行流程。

### 12.2 页面路径

| 页面 | 路由路径 | 说明 |
|------|----------|------|
| 工单排程 | `/mes/workScheduling` | 工单主页面 |
| 工单任务明细 | `/mes/workScheduling/detail` | 工序级任务明细 |

### 12.3 核心字段

**工单主表(MesWorkScheduling)**

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
| completeStatus | String | 完工状态 |
| priority | Integer | 优先级(1-10) |
| batchCode | String | 批次码 |

**工序任务明细(MesWorkSchedulingDetail)**

| 字段 | 类型 | 说明 |
|------|------|------|
| schedulingCode | String | 工单编码 |
| processCode | String | 工序编码 |
| processName | String | 工序名称 |
| seqNo | Integer | 工序顺序 |
| workroomCode | String | 车间编码 |
| lineCode | String | 产线编码 |
| stationCode | String | 工位编码 |
| workstationCode | String | 工位作业地 |
| workGroupCode | String | 班组编码 |
| workerCode | String | 作业人员编码 |
| equipmentCode | String | 设备编码 |
| planQty | Decimal | 计划数量 |
| taskQty | Decimal | 任务数量 |
| doneQty | Decimal | 完成数量 |
| taskStatus | String | 任务状态(WAIT待领/PENDING待报/PROCESSING进行中/COMPLETED已完成) |
| stdWorkTime | Decimal | 标准工时(分钟) |
| planStartTime | DateTime | 计划开始时间 |
| planEndTime | DateTime | 计划结束时间 |

### 12.4 API接口

**工单排程 (/mes/workScheduling/*)**

| 方法 | 路径 | 说明 |
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

**工单任务明细 (/mes/work-scheduling-detail/*)**

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /mes/work-scheduling-detail/create | 创建工单任务明细 |
| PUT | /mes/work-scheduling-detail/update | 更新工单任务明细 |
| DELETE | /mes/work-scheduling-detail/delete | 删除工单任务明细 |
| GET | /mes/work-scheduling-detail/get | 获取工单任务明细 |
| GET | /mes/work-scheduling-detail/page | 获取工单任务明细分页 |
| POST | /mes/work-scheduling-detail/senior | 高级搜索工单任务明细 |
| POST | /mes/work-scheduling-detail/reportWorkByProcess | 工序报工 |
| GET | /mes/work-scheduling-detail/processFinished | 工序完工 |
| POST | /mes/work-scheduling-detail/processQualified | 工序质检 |
| GET | /mes/work-scheduling-detail/export-excel | 导出工单任务明细Excel |

---

## §13 BOM管理 (bom)

### 13.1 功能说明

BOM管理用于配置产品物料清单，支持多级BOM结构，维护物料用量、损耗率等核心数据。

### 13.2 页面路径

| 页面 | 路由路径 | 说明 |
|------|----------|------|
| BOM管理 | `/mes/bom` | 物料清单配置 |
| BOM版本 | `/mes/bom/version` | BOM版本管理 |

### 13.3 核心字段

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

### 13.4 API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /mes/common/geBomTreeByProductCode | 根据产品编码获取BOM树形结构 |
| GET | /mes/common/getBomListByProductCode | 根据产品编码获取BOM列表 |
| GET | /mes/common/getBomListByProductAndProcess | 根据产品编码获取工序物料 |

---

## §14 工序管理 (opersteps)

### 14.1 功能说明

工序管理用于配置生产工序的详细参数，包括标准工时、质检要求、设备工具需求等。

### 14.2 页面路径

| 页面 | 路由路径 | 说明 |
|------|----------|------|
| 工序管理 | `/mes/opersteps` | 工序详细配置 |
| 工序类型 | `/mes/operstepsType` | 工序类型定义 |

### 14.3 核心字段

**工序(Opersteps)**

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

### 14.4 API接口

**工序管理 (/mes/opersteps/*)**

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /mes/opersteps/create | 创建MES操作步骤信息 |
| PUT | /mes/opersteps/update | 更新MES操作步骤信息 |
| DELETE | /mes/opersteps/delete | 删除MES操作步骤信息 |
| GET | /mes/opersteps/get | 获取MES操作步骤信息 |
| GET | /mes/opersteps/list | 获取MES操作步骤信息列表 |
| GET | /mes/opersteps/page | 获取MES操作步骤信息分页 |
| POST | /mes/opersteps/senior | 高级搜索操作步骤类型配置 |
| GET | /mes/opersteps/export-excel | 导出MES操作步骤信息Excel |
| POST | /mes/opersteps/import | 导入MES操作步骤信息 |

**工序类型 (/mes/opersteps-type/*)**

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /mes/opersteps-type/create | 创建操作步骤类型配置 |
| PUT | /mes/opersteps-type/update | 更新操作步骤类型配置 |
| DELETE | /mes/opersteps-type/delete | 删除操作步骤类型配置 |
| GET | /mes/opersteps-type/get | 获取操作步骤类型配置 |
| GET | /mes/opersteps-type/list | 获取操作步骤类型配置列表 |
| GET | /mes/opersteps-type/page | 获取操作步骤类型配置分页 |
| POST | /mes/opersteps-type/senior | 高级搜索操作步骤类型配置 |
| GET | /mes/opersteps-type/export-excel | 导出操作步骤类型配置Excel |
| POST | /mes/opersteps-type/import | 导入操作步骤类型配置 |

---

## §15 假期管理 (holiday)

### 15.1 功能说明

假期管理用于配置工厂日历中的节假日信息，支持按年查看和管理假期数据。

### 15.2 页面路径

| 页面 | 路由路径 | 说明 |
|------|----------|------|
| 假期管理 | `/mes/holiday` | 假期日历配置 |

### 15.3 核心字段

| 字段 | 类型 | 说明 |
|------|------|------|
| holidayName | String | 假期名称 |
| holidayDate | Date | 假期日期 |
| holidayType | String | 假期类型(法定节假日/公司福利假/调休工作日) |
| year | String | 年份 |
| keyDate | String | 日期Key(格式: yyyy-MM-dd) |

### 15.4 API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /mes/holiday/create | 创建节假日设置 |
| PUT | /mes/holiday/update | 更新节假日设置 |
| DELETE | /mes/holiday/delete | 删除节假日设置 |
| GET | /mes/holiday/get | 获取节假日设置 |
| GET | /mes/holiday/list | 获取指定年份节假日数据列表 |
| GET | /mes/holiday/listByYear | 获取指定年份节假日数据(按keyDate分组) |
| GET | /mes/holiday/page | 获取节假日设置分页 |
| POST | /mes/holiday/senior | 高级搜索节假日信息 |
| GET | /mes/holiday/export-excel | 导出节假日设置Excel |
| POST | /mes/holiday/import | 导入节假日设置 |

---

## §16 产品档案 (item)

### 16.1 功能说明

产品档案管理产品的BOM物料信息，建立产品与物料的关联关系，支持生产物料追溯。

### 16.2 页面路径

| 页面 | 路由路径 | 说明 |
|------|----------|------|
| 产品档案 | `/mes/item` | 产品BOM物料档案 |

### 16.3 核心字段

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

### 16.4 API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /mes/common/getBomTreeByProductCode | 根据产品编码获取BOM树形结构 |
| GET | /mes/common/getBomListByProductCode | 根据产品编码获取BOM列表 |
| GET | /mes/common/getBomListByProductAndProcess | 根据产品编码获取工序物料 |

---

## §17 物料申请 (itemRequestMain)

### 17.1 功能说明

物料申请用于生产物料的领用申请管理，支持叫料申请、补料申请、领料等业务流程。

### 17.2 页面路径

| 页面 | 路由路径 | 说明 |
|------|----------|------|
| 物料申请 | `/mes/itemRequestMain` | 生产物料申请 |
| 物料申请明细 | `/mes/itemRequestMain/detail` | 物料申请明细 |

### 17.3 核心字段

**物料申请主表**

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

**物料申请明细**

| 字段 | 类型 | 说明 |
|------|------|------|
| requestNo | String | 申请单号 |
| itemCode | String | 物料编码 |
| itemName | String | 物料名称 |
| itemSpec | String | 物料规格 |
| requestQty | Decimal | 申请数量 |
| deliveredQty | Decimal | 已配送数量 |
| receivedQty | Decimal | 已领用数量 |

### 17.4 API接口

**物料申请主表 (/mes/item-request-main/*)**

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /mes/item-request-main/create | 创建叫料申请单主 |
| POST | /mes/item-request-main/create-call-material | 创建叫料申请单 |
| PUT | /mes/item-request-main/update | 更新叫料申请单主 |
| DELETE | /mes/item-request-main/delete | 删除叫料申请单主 |
| GET | /mes/item-request-main/get | 获取叫料申请单主 |
| GET | /mes/item-request-main/page | 获取叫料申请单主分页 |
| POST | /mes/item-request-main/senior | 高级搜索叫料申请单主 |
| POST | /mes/item-request-main/addBasicItem | 创建补料申请 |
| POST | /mes/item-request-main/receiveBasicItem | 领料 |
| POST | /mes/item-request-main/receiveItem | 领料 |

**物料申请明细 (/mes/item-request-detail/*)**

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /mes/item-request-detail/create | 创建叫料申请明细 |
| PUT | /mes/item-request-detail/update | 更新叫料申请明细 |
| DELETE | /mes/item-request-detail/delete | 删除叫料申请明细 |
| GET | /mes/item-request-detail/get | 获取叫料申请明细 |
| GET | /mes/item-request-detail/page | 获取叫料申请明细分页 |
| POST | /mes/item-request-detail/senior | 高级搜索叫料申请明细 |

---

## §18 产品拆解 (dismantlingMain)

### 18.1 功能说明

产品拆解用于处理报废拆解任务，记录拆解过程中物料的回收和损耗情况。

### 18.2 页面路径

| 页面 | 路由路径 | 说明 |
|------|----------|------|
| 产品拆解 | `/mes/dismantlingMain` | 产品拆解任务 |
| 拆解明细 | `/mes/dismantlingMain/detail` | 拆解物料明细 |

### 18.3 核心字段

**拆解主表**

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

**拆解明细**

| 字段 | 类型 | 说明 |
|------|------|------|
| dismantlingNo | String | 拆解单号 |
| itemCode | String | 物料编码 |
| itemName | String | 物料名称 |
| itemSpec | String | 物料规格 |
| quantity | Decimal | 拆解数量 |
|回收数量 | Decimal | 回收数量 |

### 18.4 API接口

**拆解主表 (/mes/dismantling-main/*)**

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /mes/dismantling-main/create | 创建报废拆解登记主 |
| PUT | /mes/dismantling-main/update | 更新报废拆解登记主 |
| DELETE | /mes/dismantling-main/delete | 删除报废拆解登记主 |
| GET | /mes/dismantling-main/get | 获取报废拆解登记主 |
| GET | /mes/dismantling-main/page | 获取报废拆解登记主分页 |
| POST | /mes/dismantling-main/senior | 高级搜索报废拆解登记主 |

**拆解明细 (/mes/dismantling-detail/*)**

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /mes/dismantling-detail/create | 创建报废拆解明细 |
| PUT | /mes/dismantling-detail/update | 更新报废拆解明细 |
| DELETE | /mes/dismantling-detail/delete | 删除报废拆解明细 |
| GET | /mes/dismantling-detail/get | 获取报废拆解明细 |
| GET | /mes/dismantling-detail/page | 获取报废拆解明细分页 |
| POST | /mes/dismantling-detail/senior | 高级搜索报废拆解明细 |

---

## §19 返工批次 (reworkBatch)

### 19.1 功能说明

返工批次管理用于处理不良品的返工返修流程，支持批次级返工任务的分发、领取、完成等。

### 19.2 页面路径

| 页面 | 路由路径 | 说明 |
|------|----------|------|
| 返工批次 | `/mes/reworkBatch` | 返工批次管理 |
| 返工单件 | `/mes/reworkSingle` | 返工单件管理 |

### 19.3 核心字段

**返工批次主表**

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

### 19.4 API接口

**返工批次 (/mes/rework-batch/*)**

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /mes/rework-batch/create | 创建返工登记批量 |
| PUT | /mes/rework-batch/update | 更新返工登记批量 |
| DELETE | /mes/rework-batch/delete | 删除返工登记批量 |
| DELETE | /mes/rework-batch/suspend | 中止返工登记批量任务 |
| PUT | /mes/rework-batch/receive | 领取返工登记批量任务 |
| PUT | /mes/rework-batch/finish | 完成返工登记批量任务 |
| GET | /mes/rework-batch/get | 获取返工登记批量 |
| GET | /mes/rework-batch/list | 获取返工登记批量列表 |
| GET | /mes/rework-batch/page | 获取返工登记批量分页 |
| POST | /mes/rework-batch/senior | 高级搜索返工登记批量 |

**返工单件 (/mes/rework-single/*)**

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /mes/rework-single/create | 创建返工登记单件 |
| PUT | /mes/rework-single/update | 更新返工登记单件 |
| DELETE | /mes/rework-single/delete | 删除返工登记单件 |
| DELETE | /mes/rework-single/suspend | 中止返工登记单件任务 |
| PUT | /mes/rework-single/receive | 领取返工登记单件任务 |
| PUT | /mes/rework-single/finish | 完成返工登记单件任务 |
| GET | /mes/rework-single/get | 获取返工登记单件 |
| GET | /mes/rework-single/list | 获取返工登记单件列表 |
| GET | /mes/rework-single/page | 获取返工登记单件分页 |
| POST | /mes/rework-single/senior | 高级搜索返工登记单件 |

---

## §20 月计划管理 (ordermonthplan)

### 20.1 功能说明

月计划管理用于月度生产计划的制定和分解，支持将月计划拆解为日计划。

### 20.2 页面路径

| 页面 | 路由路径 | 说明 |
|------|----------|------|
| 月计划 | `/mes/ordermonthplan` | 月度生产计划 |
| 月计划明细 | `/mes/ordermonthplan/detail` | 月计划子项明细 |

### 20.3 核心字段

**月计划主表**

| 字段 | 类型 | 说明 |
|------|------|------|
| monthPlanNo | String | 月计划单号 |
| month | String | 计划月份(格式: yyyy-MM) |
| productCode | String | 产品图号 |
| productName | String | 产品名称 |
| totalQty | Integer | 计划总量 |
| deliveredQty | Integer | 已排产量 |
| status | String | 状态 |

**月计划子表**

| 字段 | 类型 | 说明 |
|------|------|------|
| monthPlanNo | String | 月计划单号 |
| productCode | String | 产品图号 |
| day01-day31 | Integer | 每日分配数量 |
| actualDay01-actualDay31 | Integer | 每日实际完成数量 |

### 20.4 API接口

**月计划主表 (/plan/mes-order-month-main/*)**

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /plan/mes-order-month-main/create | 创建订单计划月主 |
| PUT | /plan/mes-order-month-main/update | 更新订单计划月主 |
| DELETE | /plan/mes-order-month-main/delete | 删除订单计划月主 |
| GET | /plan/mes-order-month-main/get | 获取订单计划月主 |
| GET | /plan/mes-order-month-main/list | 获取订单计划月主列表 |
| GET | /plan/mes-order-month-main/page | 获取订单计划月主分页 |

**月计划子表 (/plan/mes-order-month-sub/*)**

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /plan/mes-order-month-sub/create | 创建订单月计划子 |
| PUT | /plan/mes-order-month-sub/update | 更新订单月计划子 |
| DELETE | /plan/mes-order-month-sub/delete | 删除订单月计划子 |
| GET | /plan/mes-order-month-sub/get | 获取订单月计划子 |
| GET | /plan/mes-order-month-sub/list | 获取订单月计划子列表 |
| GET | /plan/mes-order-month-sub/page | 获取订单月计划子分页 |
| POST | /plan/mes-order-month-sub/breakdown | 拆解为日计划 |

---

## §21 工作日历 (workcalendar)

### 21.1 功能说明

工作日历用于配置班组的生产工作日历，支持按班组设置工作时段和休息日。

### 21.2 页面路径

| 页面 | 路由路径 | 说明 |
|------|----------|------|
| 工作日历 | `/mes/workcalendar` | 工作日历配置 |

### 21.3 核心字段

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

### 21.4 API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /mes/schedulingcalendar/create | 创建生产日历 |
| POST | /mes/schedulingcalendar/createBatch | 批量创建生产日历 |
| POST | /mes/schedulingcalendar/createObj | 根据json对象创建工作日历 |
| PUT | /mes/schedulingcalendar/update | 更新生产日历 |
| DELETE | /mes/schedulingcalendar/delete | 删除生产日历 |
| POST | /mes/schedulingcalendar/deleteTeam | 删除班组日历 |
| GET | /mes/schedulingcalendar/get | 获取生产日历 |
| GET | /mes/schedulingcalendar/getList | 获取班组的生产日历(按keyDate分组) |
| GET | /mes/schedulingcalendar/page | 获取生产日历分页 |
| POST | /mes/schedulingcalendar/senior | 高级搜索生产日历 |
| GET | /mes/schedulingcalendar/export-excel | 导出生产日历Excel |
| POST | /mes/schedulingcalendar/import | 导入生产日历 |

---

## §22 工艺路线管理 (processroute)

### 22.1 功能说明

工艺路线管理用于定义产品的生产加工路线，配置各工序的先后顺序、工艺参数和资源需求。

### 22.2 页面路径

| 页面 | 路由路径 | 说明 |
|------|----------|------|
| 工艺路线 | `/mes/processroute` | 工艺路线定义 |
| 工艺路线工序 | `/mes/processrouteNodeDetail` | 工艺路线工序明细 |

### 22.3 核心字段

**工艺路线主表**

| 字段 | 类型 | 说明 |
|------|------|------|
| processrouteCode | String | 工艺路线编码 |
| processrouteName | String | 工艺路线名称 |
| processrouteVersion | String | 工艺路线版本 |
| productCode | String | 适用产品图号 |
| workroomCode | String | 车间编码 |
| lineCode | String | 产线编码 |
| routeGraph | JSON | 工艺路线图形定义 |
| status | String | 状态(EDITING编辑/PUBLISHED已发布) |

**工艺路线工序明细**

| 字段 | 类型 | 说明 |
|------|------|------|
| processrouteCode | String | 工艺路线编码 |
| processCode | String | 工序编码 |
| processName | String | 工序名称 |
| seqNo | Integer | 工序顺序号 |
| preProcessCode | String | 前工序编码 |
| nextProcessCode | String | 后工序编码 |
| isKeyProcess | Boolean | 是否关键工序 |
| stdWorkTime | Decimal | 标准工时(分钟) |
| stdCapacity | Integer | 标准产能(件/日) |

### 22.4 API接口

**工艺路线 (/mes/processroute/*)**

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /mes/processroute/create | 创建工艺路线定义 |
| PUT | /mes/processroute/update | 更新工艺路线定义 |
| POST | /mes/processroute/updategraph | 更新工艺路线的图形定义 |
| DELETE | /mes/processroute/delete | 删除工艺路线定义 |
| GET | /mes/processroute/get | 获取工艺路线定义 |
| GET | /mes/processroute/list | 获取工艺路线定义列表 |
| GET | /mes/processroute/page | 获取工艺路线定义分页 |
| POST | /mes/processroute/senior | 高级搜索工艺路线定义 |

**工艺路线工序 (/mes/processrouteNodeDetail/*)**

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /mes/processrouteNodeDetail/create | 创建工艺路由工序节点配置明细 |
| PUT | /mes/processrouteNodeDetail/update | 更新工艺路由工序节点配置明细 |
| DELETE | /mes/processrouteNodeDetail/delete | 删除工艺路由工序节点配置明细 |
| GET | /mes/processrouteNodeDetail/get | 获取工艺路由工序节点配置明细 |
| GET | /mes/processrouteNodeDetail/list | 获取工艺路由工序节点配置明细列表 |
| GET | /mes/processrouteNodeDetail/page | 获取工艺路由工序节点配置明细分页 |
| GET | /mes/processrouteNodeDetail/getRouteNodes | 获取工艺路线已配置的工序列表 |

---

## §23 状态机汇总

### 23.1 日计划状态

| 状态 | 值 | 可执行操作 | 下一状态 |
|------|-----|-----------|----------|
| 待排产 | WAITSECHUDLE | 编辑、发布、删除 | PUBLISHED |
| 已发布 | PUBLISHED | 终止 | TERMINALE |
| 进行中 | PROCESSING | 终止 | TERMINALE |
| 已终止 | TERMINALE | 恢复 | PUBLISHED |

### 23.2 工单作业状态

| 状态 | 值 | 可执行操作 | 下一状态 |
|------|-----|-----------|----------|
| 待报工 | READYTODO | 开始 | WAITBEGIN |
| 待开工 | WAITBEGIN | 开始 | PROCESSING |
| 进行中 | PROCESSING | 完工、终止 | FINISH, TERMINALE |
| 已完成 | FINISH | - | - |
| 已终止 | TERMINALE | - | - |

### 23.3 返工批次状态

| 状态 | 值 | 可执行操作 | 下一状态 |
|------|-----|-----------|----------|
| 待领取 | PENDING | 领取 | RECEIVED |
| 已领取 | RECEIVED | 开始 | PROCESSING |
| 进行中 | PROCESSING | 完成、中止 | COMPLETED, SUSPENDED |
| 已完成 | COMPLETED | - | - |
| 已中止 | SUSPENDED | - | - |

---

> 文档补充时间: 2026-04-17
> 补充内容: 生产日计划、工单排程、BOM管理、工序管理、假期管理、产品档案、物料申请、产品拆解、返工批次、月计划、工作日历、工艺路线等完整页面设计

### 10.1 数据库表结构

#### 10.1.1 plan_mes_order_day_bom - 日计划BOM表

```sql
CREATE TABLE `plan_mes_order_day_bom` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `plan_no_day` varchar(50) DEFAULT NULL COMMENT '日计划单号',
  `product_code` varchar(50) DEFAULT NULL COMMENT '产品图号',
  `bom_code` varchar(50) DEFAULT NULL COMMENT 'BOM编码',
  `bom_version` varchar(20) DEFAULT NULL COMMENT 'BOM版本',
  `item_code` varchar(50) DEFAULT NULL COMMENT '物料编码',
  `item_name` varchar(200) DEFAULT NULL COMMENT '物料名称',
  `item_spec` varchar(100) DEFAULT NULL COMMENT '物料规格',
  `item_unit` varchar(20) DEFAULT NULL COMMENT '计量单位',
  `item_type` varchar(20) DEFAULT NULL COMMENT '物料类型(原材料/成品/半成品)',
  `bom_level` int DEFAULT NULL COMMENT 'BOM层级',
  `parent_item_code` varchar(50) DEFAULT NULL COMMENT '父物料编码',
  `standard_qty` decimal(18,6) DEFAULT NULL COMMENT '标准用量',
  `actual_qty` decimal(18,6) DEFAULT NULL COMMENT '实际用量',
  `loss_rate` decimal(6,4) DEFAULT NULL COMMENT '损耗率',
  `warehouse_code` varchar(50) DEFAULT NULL COMMENT '仓库编码',
  `location_code` varchar(50) DEFAULT NULL COMMENT '库位编码',
  `process_route_code` varchar(50) DEFAULT NULL COMMENT '工艺路线编码',
  `process_code` varchar(50) DEFAULT NULL COMMENT '工序编码',
  `seq_no` int DEFAULT NULL COMMENT '顺序号',
  `is_key` tinyint(1) DEFAULT '0' COMMENT '是否关键件',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  KEY `idx_plan_no_day` (`plan_no_day`),
  KEY `idx_product_code` (`product_code`),
  KEY `idx_item_code` (`item_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日计划BOM表';
```

#### 10.1.2 plan_mes_order_day_route - 日计划工艺路线实例表

```sql
CREATE TABLE `plan_mes_order_day_route` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `plan_no_day` varchar(50) DEFAULT NULL COMMENT '日计划单号',
  `product_code` varchar(50) DEFAULT NULL COMMENT '产品图号',
  `processroute_code` varchar(50) DEFAULT NULL COMMENT '工艺路线编码',
  `processroute_name` varchar(200) DEFAULT NULL COMMENT '工艺路线名称',
  `processroute_version` varchar(20) DEFAULT NULL COMMENT '工艺路线版本',
  `workroom_code` varchar(50) DEFAULT NULL COMMENT '车间编码',
  `line_code` varchar(50) DEFAULT NULL COMMENT '产线编码',
  `is_temp` tinyint(1) DEFAULT '0' COMMENT '是否临时工艺(0否/1是)',
  `route_graph` longtext COMMENT '工艺路线图形定义(JSON)',
  `status` varchar(20) DEFAULT NULL COMMENT '状态(EDITING编辑/PUBLISHED已发布)',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  KEY `idx_plan_no_day` (`plan_no_day`),
  KEY `idx_product_code` (`product_code`),
  KEY `idx_processroute_code` (`processroute_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日计划工艺路线实例表';
```

#### 10.1.3 plan_mes_order_day_routesub - 日计划工序明细表

```sql
CREATE TABLE `plan_mes_order_day_routesub` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `plan_no_day` varchar(50) DEFAULT NULL COMMENT '日计划单号',
  `product_code` varchar(50) DEFAULT NULL COMMENT '产品图号',
  `processroute_code` varchar(50) DEFAULT NULL COMMENT '工艺路线编码',
  `process_code` varchar(50) DEFAULT NULL COMMENT '工序编码',
  `process_name` varchar(200) DEFAULT NULL COMMENT '工序名称',
  `seq_no` int DEFAULT NULL COMMENT '工序顺序号',
  `workroom_code` varchar(50) DEFAULT NULL COMMENT '车间编码',
  `line_code` varchar(50) DEFAULT NULL COMMENT '产线编码',
  `station_code` varchar(50) DEFAULT NULL COMMENT '工位编码',
  `workstation_code` varchar(50) DEFAULT NULL COMMENT '工位作业地',
  `std_work_time` decimal(10,2) DEFAULT NULL COMMENT '标准工时(分钟)',
  `std_capacity` int DEFAULT NULL COMMENT '标准产能(件/日)',
  `pre_process_code` varchar(50) DEFAULT NULL COMMENT '前工序编码',
  `next_process_code` varchar(50) DEFAULT NULL COMMENT '后工序编码',
  `is_key_process` tinyint(1) DEFAULT '0' COMMENT '是否关键工序',
  `is_auto_complete` tinyint(1) DEFAULT '0' COMMENT '是否自动完工',
  `quality_control_type` varchar(20) DEFAULT NULL COMMENT '质检类型(免检/自检/互检/专检)',
  `require_inspect` tinyint(1) DEFAULT '0' COMMENT '是否需要质检',
  `inspect_form_code` varchar(50) DEFAULT NULL COMMENT '质检表单编码',
  `tool_required` varchar(500) DEFAULT NULL COMMENT '需要的工具/模具',
  `material_feed_type` varchar(20) DEFAULT NULL COMMENT '投料方式(齐套/顺序)',
  `status` varchar(20) DEFAULT NULL COMMENT '状态',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  KEY `idx_plan_no_day` (`plan_no_day`),
  KEY `idx_process_code` (`process_code`),
  KEY `idx_seq_no` (`seq_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日计划工序明细表';
```

#### 10.1.4 plan_mes_work_scheduling - 工单排程主表

```sql
CREATE TABLE `plan_mes_work_scheduling` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `scheduling_code` varchar(50) DEFAULT NULL COMMENT '工单编码',
  `plan_master_code` varchar(50) DEFAULT NULL COMMENT '计划单号(来源日计划)',
  `product_code` varchar(50) DEFAULT NULL COMMENT '产品图号',
  `product_name` varchar(200) DEFAULT NULL COMMENT '产品名称',
  `workroom_code` varchar(50) DEFAULT NULL COMMENT '车间编码',
  `line_code` varchar(50) DEFAULT NULL COMMENT '产线编码',
  `plan_qty` decimal(18,3) DEFAULT NULL COMMENT '计划数量',
  `done_qty` decimal(18,3) DEFAULT NULL COMMENT '已完成数量',
  `qualified_qty` decimal(18,3) DEFAULT NULL COMMENT '合格数量',
  `rejected_qty` decimal(18,3) DEFAULT NULL COMMENT '不合格数量',
  `work_mode` varchar(20) DEFAULT NULL COMMENT '工单模式(BATCH批量/SINGLE单件)',
  `task_mode` varchar(20) DEFAULT NULL COMMENT '生产模式(ASSIGN派工/RECIEVE领工)',
  `processroute_code` varchar(50) DEFAULT NULL COMMENT '工艺路线编码',
  `bom_code` varchar(50) DEFAULT NULL COMMENT 'BOM编码',
  `plan_start_date` date DEFAULT NULL COMMENT '计划开始日期',
  `plan_end_date` date DEFAULT NULL COMMENT '计划结束日期',
  `actual_start_date` datetime DEFAULT NULL COMMENT '实际开始时间',
  `actual_end_date` datetime DEFAULT NULL COMMENT '实际结束时间',
  `working_status` varchar(20) DEFAULT NULL COMMENT '作业状态',
  `complete_status` varchar(20) DEFAULT NULL COMMENT '完工状态',
  `priority` int DEFAULT '5' COMMENT '优先级(1-10)',
  `batch_code` varchar(50) DEFAULT NULL COMMENT '批次码',
  `sales_order_no` varchar(50) DEFAULT NULL COMMENT '销售订单号',
  `customer_name` varchar(200) DEFAULT NULL COMMENT '客户名称',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_scheduling_code` (`scheduling_code`),
  KEY `idx_plan_master_code` (`plan_master_code`),
  KEY `idx_product_code` (`product_code`),
  KEY `idx_working_status` (`working_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='工单排程主表';
```

#### 10.1.5 plan_mes_work_scheduling_detail - 工单排程明细表

```sql
CREATE TABLE `plan_mes_work_scheduling_detail` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `scheduling_id` bigint DEFAULT NULL COMMENT '工单主表ID',
  `scheduling_code` varchar(50) DEFAULT NULL COMMENT '工单编码',
  `plan_no_day` varchar(50) DEFAULT NULL COMMENT '日计划单号',
  `product_code` varchar(50) DEFAULT NULL COMMENT '产品图号',
  `process_code` varchar(50) DEFAULT NULL COMMENT '工序编码',
  `process_name` varchar(200) DEFAULT NULL COMMENT '工序名称',
  `seq_no` int DEFAULT NULL COMMENT '工序顺序',
  `workroom_code` varchar(50) DEFAULT NULL COMMENT '车间编码',
  `line_code` varchar(50) DEFAULT NULL COMMENT '产线编码',
  `station_code` varchar(50) DEFAULT NULL COMMENT '工位编码',
  `workstation_code` varchar(50) DEFAULT NULL COMMENT '工位作业地',
  `work_group_code` varchar(50) DEFAULT NULL COMMENT '班组编码',
  `work_group_name` varchar(100) DEFAULT NULL COMMENT '班组名称',
  `worker_code` varchar(50) DEFAULT NULL COMMENT '作业人员编码',
  `worker_name` varchar(100) DEFAULT NULL COMMENT '作业人员姓名',
  `equipment_code` varchar(50) DEFAULT NULL COMMENT '设备编码',
  `equipment_name` varchar(200) DEFAULT NULL COMMENT '设备名称',
  `plan_qty` decimal(18,3) DEFAULT NULL COMMENT '计划数量',
  `task_qty` decimal(18,3) DEFAULT NULL COMMENT '任务数量',
  `done_qty` decimal(18,3) DEFAULT NULL COMMENT '完成数量',
  `qualified_qty` decimal(18,3) DEFAULT NULL COMMENT '合格数量',
  `rejected_qty` decimal(18,3) DEFAULT NULL COMMENT '不合格数量',
  `std_work_time` decimal(10,2) DEFAULT NULL COMMENT '标准工时(分钟)',
  `actual_work_time` decimal(10,2) DEFAULT NULL COMMENT '实际工时(分钟)',
  `plan_start_time` datetime DEFAULT NULL COMMENT '计划开始时间',
  `plan_end_time` datetime DEFAULT NULL COMMENT '计划结束时间',
  `actual_start_time` datetime DEFAULT NULL COMMENT '实际开始时间',
  `actual_end_time` datetime DEFAULT NULL COMMENT '实际结束时间',
  `task_status` varchar(20) DEFAULT NULL COMMENT '任务状态(WAIT待领/PENDING待报/PROCESSING进行中/COMPLETED已完成)',
  `report_status` varchar(20) DEFAULT NULL COMMENT '报工状态',
  `inspect_status` varchar(20) DEFAULT NULL COMMENT '质检状态',
  `is_current_process` tinyint(1) DEFAULT '0' COMMENT '是否当前工序',
  `pre_task_id` bigint DEFAULT NULL COMMENT '前道工序任务ID',
  `next_task_id` bigint DEFAULT NULL COMMENT '后道工序任务ID',
  `complete_inspect` varchar(20) DEFAULT NULL COMMENT '齐套检查状态',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  KEY `idx_scheduling_id` (`scheduling_id`),
  KEY `idx_scheduling_code` (`scheduling_code`),
  KEY `idx_process_code` (`process_code`),
  KEY `idx_task_status` (`task_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='工单排程明细表';
```

#### 10.1.6 plan_mes_order_day_equipment - 日计划设备表

```sql
CREATE TABLE `plan_mes_order_day_equipment` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `plan_no_day` varchar(50) DEFAULT NULL COMMENT '日计划单号',
  `product_code` varchar(50) DEFAULT NULL COMMENT '产品图号',
  `processroute_code` varchar(50) DEFAULT NULL COMMENT '工艺路线编码',
  `process_code` varchar(50) DEFAULT NULL COMMENT '工序编码',
  `process_name` varchar(200) DEFAULT NULL COMMENT '工序名称',
  `seq_no` int DEFAULT NULL COMMENT '工序顺序号',
  `workroom_code` varchar(50) DEFAULT NULL COMMENT '车间编码',
  `line_code` varchar(50) DEFAULT NULL COMMENT '产线编码',
  `equipment_code` varchar(50) DEFAULT NULL COMMENT '设备编码',
  `equipment_name` varchar(200) DEFAULT NULL COMMENT '设备名称',
  `equipment_type` varchar(50) DEFAULT NULL COMMENT '设备类型',
  `equipment_spec` varchar(200) DEFAULT NULL COMMENT '设备规格',
  `is_main` tinyint(1) DEFAULT '0' COMMENT '是否主设备',
  `status` varchar(20) DEFAULT NULL COMMENT '状态',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  KEY `idx_plan_no_day` (`plan_no_day`),
  KEY `idx_process_code` (`process_code`),
  KEY `idx_equipment_code` (`equipment_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日计划设备表';
```

#### 10.1.7 plan_mes_order_day_worker - 日计划人员表

```sql
CREATE TABLE `plan_mes_order_day_worker` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `plan_no_day` varchar(50) DEFAULT NULL COMMENT '日计划单号',
  `product_code` varchar(50) DEFAULT NULL COMMENT '产品图号',
  `processroute_code` varchar(50) DEFAULT NULL COMMENT '工艺路线编码',
  `process_code` varchar(50) DEFAULT NULL COMMENT '工序编码',
  `process_name` varchar(200) DEFAULT NULL COMMENT '工序名称',
  `seq_no` int DEFAULT NULL COMMENT '工序顺序号',
  `workroom_code` varchar(50) DEFAULT NULL COMMENT '车间编码',
  `line_code` varchar(50) DEFAULT NULL COMMENT '产线编码',
  `work_group_code` varchar(50) DEFAULT NULL COMMENT '班组编码',
  `work_group_name` varchar(100) DEFAULT NULL COMMENT '班组名称',
  `worker_code` varchar(50) DEFAULT NULL COMMENT '人员编码',
  `worker_name` varchar(100) DEFAULT NULL COMMENT '人员姓名',
  `post_code` varchar(50) DEFAULT NULL COMMENT '岗位编码',
  `post_name` varchar(100) DEFAULT NULL COMMENT '岗位名称',
  `skill_level` varchar(20) DEFAULT NULL COMMENT '技能等级',
  `is_leader` tinyint(1) DEFAULT '0' COMMENT '是否班组长',
  `work_ratio` decimal(5,2) DEFAULT '100.00' COMMENT '工作分配比例(%)',
  `status` varchar(20) DEFAULT NULL COMMENT '状态',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  KEY `idx_plan_no_day` (`plan_no_day`),
  KEY `idx_process_code` (`process_code`),
  KEY `idx_worker_code` (`worker_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日计划人员表';
```

#### 10.1.8 plan_mes_order_day_workstation - 日计划工位表

```sql
CREATE TABLE `plan_mes_order_day_workstation` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `plan_no_day` varchar(50) DEFAULT NULL COMMENT '日计划单号',
  `product_code` varchar(50) DEFAULT NULL COMMENT '产品图号',
  `processroute_code` varchar(50) DEFAULT NULL COMMENT '工艺路线编码',
  `process_code` varchar(50) DEFAULT NULL COMMENT '工序编码',
  `process_name` varchar(200) DEFAULT NULL COMMENT '工序名称',
  `seq_no` int DEFAULT NULL COMMENT '工序顺序号',
  `workroom_code` varchar(50) DEFAULT NULL COMMENT '车间编码',
  `line_code` varchar(50) DEFAULT NULL COMMENT '产线编码',
  `workstation_code` varchar(50) DEFAULT NULL COMMENT '工位编码',
  `workstation_name` varchar(200) DEFAULT NULL COMMENT '工位名称',
  `workstation_type` varchar(50) DEFAULT NULL COMMENT '工位类型',
  `capacity` int DEFAULT NULL COMMENT '产能(件/日)',
  `status` varchar(20) DEFAULT NULL COMMENT '状态',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  KEY `idx_plan_no_day` (`plan_no_day`),
  KEY `idx_process_code` (`process_code`),
  KEY `idx_workstation_code` (`workstation_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日计划工位表';
```

---

### 10.2 齐套检查API (/mes/complete-inspect/*)

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| GET | /mes/complete-inspect/get | 获得齐套检查标识 | paramCode (String) | MesConfigInfoDO |
| POST | /mes/complete-inspect/get-orderDay-bom | 获取日计划Bom信息 | MesWorkSchedulingBaseVO | List\<MesOrderDayBomRespVO\> |
| POST | /mes/complete-inspect/get-orderDay-bom-page | 获取日计划Bom信息(分页) | MesWorkSchedulingPageReqVO | PageResult\<MesOrderDayBomRespVO\> |
| POST | /mes/complete-inspect/get-orderDay-worker | 获取日计划人员信息 | MesWorkSchedulingBaseVO | List\<MesOrderDayWorkerRespVO\> |
| POST | /mes/complete-inspect/get-orderDay-worker-page | 获取日计划Worker信息(分页) | MesWorkSchedulingPageReqVO | PageResult\<MesOrderDayWorkerRespVO\> |
| POST | /mes/complete-inspect/get-orderDay-equipment | 获取日计划设备信息 | MesWorkSchedulingBaseVO | List\<MesOrderDayEquipmentRespVO\> |
| POST | /mes/complete-inspect/get-orderDay-equipment-page | 获取日计划Equipment信息(分页) | MesWorkSchedulingPageReqVO | PageResult\<MesOrderDayEquipmentRespVO\> |
| POST | /mes/complete-inspect/update | 更新生产日工单 | MesWorkSchedulingUpdateReqVO | Boolean |

---

### 10.3 工艺路线API (/mes/orderday/route/*)

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| GET | /mes/orderDayRoute/get | 获得计划产品工艺路线配置 | id (Long) | MesOrderDayRouteRespVO |
| GET | /mes/orderDayRoute/getByOrder | 获得计划产品工艺路线配置 | OrderDayQueryParamVo | MesOrderDayRouteRespVO |
| PUT | /mes/orderDayRoute/update | 更新计划产品工艺路线配置 | MesOrderDayRouteUpdateReqVO | Boolean |
| GET | /mes/orderDayRoute/page | 获得计划产品工艺路线配置分页 | MesOrderDayRoutePageReqVO | PageResult\<MesOrderDayRouteRespVO\> |

#### 工序明细API (/mes/orderday/routesub/*)

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/orderDayRoutesub/create | 创建计划工单产品工艺路线工序 | MesOrderDayRoutesubCreateReqVO | Long(id) |
| PUT | /mes/orderDayRoutesub/update | 更新计划工单产品工艺路线工序 | MesOrderDayRoutesubUpdateReqVO | Boolean |
| DELETE | /mes/orderDayRoutesub/delete | 删除计划工单产品工艺路线工序 | id (Long) | Boolean |
| GET | /mes/orderDayRoutesub/get | 获得计划工单产品工艺路线工序 | id (Long) | MesOrderDayRoutesubRespVO |
| GET | /mes/orderDayRoutesub/getByOrder | 获得计划产品工艺路线的工序配置 | OrderDayQueryParamVo | List\<ProcessRespVO\> |
| GET | /mes/orderDayRoutesub/page | 获得计划工单产品工艺路线工序分页 | MesOrderDayRoutesubPageReqVO | PageResult\<MesOrderDayRoutesubRespVO\> |

---

### 10.4 工位/人员/设备配置API

#### 设备配置 (/mes/orderDayequipment/*)

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/orderDayequipment/create | 创建计划工单设备配置 | MesOrderDayEquipmentCreateReqVO | Long(id) |
| POST | /mes/orderDayequipment/batchCreate | 批量创建计划工单设备配置 | List\<MesOrderDayEquipmentCreateReqVO\> | null |
| PUT | /mes/orderDayequipment/update | 更新计划工单设备配置 | MesOrderDayEquipmentUpdateReqVO | Boolean |
| DELETE | /mes/orderDayequipment/delete | 删除计划工单设备配置 | id (Long) | Boolean |
| GET | /mes/orderDayequipment/get | 获得计划工单设备配置 | id (Long) | MesOrderDayEquipmentRespVO |
| GET | /mes/orderDayequipment/getByOrder | 获得计划产品工艺路线的工序配置 | OrderDayQueryParamVo | List\<MesOrderDayEquipmentRespVO\> |
| GET | /mes/orderDayequipment/page | 获得计划工单设备配置分页 | MesOrderDayEquipmentPageReqVO | PageResult\<MesOrderDayEquipmentRespVO\> |

#### 人员配置 (/mes/orderDayWorker/*)

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/orderDayWorker/create | 创建计划工单人员配置 | MesOrderDayWorkerCreateReqVO | Long(id) |
| POST | /mes/orderDayWorker/batchCreate | 批量创建计划工单人员配置 | List\<MesOrderDayWorkerCreateReqVO\> | null |
| PUT | /mes/orderDayWorker/update | 更新计划工单人员配置 | MesOrderDayWorkerUpdateReqVO | Boolean |
| DELETE | /mes/orderDayWorker/delete | 删除计划工单人员配置 | id (Long) | Boolean |
| GET | /mes/orderDayWorker/get | 获得计划工单人员配置 | id (Long) | MesOrderDayWorkerRespVO |
| GET | /mes/orderDayWorker/getByOrder | 获得计划产品工艺路线的工序配置 | OrderDayQueryParamVo | List\<MesOrderDayWorkerRespVO\> |
| GET | /mes/orderDayWorker/page | 获得计划工单人员配置分页 | MesOrderDayWorkerPageReqVO | PageResult\<MesOrderDayWorkerRespVO\> |

#### 工位配置 (/mes/orderDayWorkstation/*)

| 方法 | 路径 | 说明 | 请求参数 | 响应 |
|------|------|------|---------|------|
| POST | /mes/orderDayWorkstation/create | 创建计划工单工位配置 | MesOrderDayWorkstationCreateReqVO | Long(id) |
| PUT | /mes/orderDayWorkstation/update | 更新计划工单工位配置 | MesOrderDayWorkstationUpdateReqVO | Boolean |
| DELETE | /mes/orderDayWorkstation/delete | 删除计划工单工位配置 | id (Long) | Boolean |
| GET | /mes/orderDayWorkstation/get | 获得计划工单工位配置 | id (Long) | MesOrderDayWorkstationRespVO |
| GET | /mes/orderDayWorkstation/getByOrder | 获得计划产品工艺路线的工序配置 | OrderDayQueryParamVo | List\<MesOrderDayWorkstationRespVO\> |
| GET | /mes/orderDayWorkstation/page | 获得计划工单工位配置分页 | MesOrderDayWorkstationPageReqVO | PageResult\<MesOrderDayWorkstationRespVO\> |
| GET | /mes/orderDayWorkstation/export-excel | 导出计划工单工位配置Excel | MesOrderDayWorkstationExportReqVO | Excel文件流 |
| GET | /mes/orderDayWorkstation/get-import-template | 获得导入计划工单工位配置模板 | - | Excel文件流 |

---

### 10.5 工单排程API (/mes/workScheduling/*)

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

#### 工单任务明细 (/mes/work-scheduling-detail/*)

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

---

### 10.6 状态枚举说明

| 枚举 | 值 | 说明 |
|------|-----|------|
| 任务状态 | WAIT | 待领工 |
| 任务状态 | PENDING | 待报工 |
| 任务状态 | PROCESSING | 进行中 |
| 任务状态 | COMPLETED | 已完成 |
| 齐套检查状态 | COMPLETE | 齐套 |
| 齐套检查状态 | INCOMPLETE | 不齐套 |
| 齐套检查状态 | PENDING | 待检查 |
| 作业状态 | READYTODO | 待报工 |
| 作业状态 | WAITBEGIN | 待开工 |
| 作业状态 | PROCESSING | 进行中 |
| 作业状态 | FINISH | 已完成 |
| 作业状态 | TERMINALE | 已终止 |

---

> 文档补充时间: 2026-04-17
> 补充内容: 工艺路线与齐套管理完整表结构及API
