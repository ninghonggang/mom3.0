# MOM3.0 设备管理模块设计文档

**版本**: V2.0 | **所属模块**: M06设备管理 | **基于**: [MOM3.0_主设计文档](./MOM3.0_主设计文档.md)

---

## 1. 模块概述

### 1.1 功能定位

设备管理模块覆盖设备台账、点检、保养、维修、OEE分析全生命周期管理，实现设备资产的可控、可追溯、可优化。

### 1.2 核心功能

| 功能 | 说明 |
|------|------|
| 设备台账 | 设备主数据、分类、位置 |
| 设备点检 | 日常点检、定期点检、专项点检 |
| 设备保养 | 保养计划、执行记录 |
| 设备维修 | 维修工单、故障记录 |
| OEE分析 | 设备综合效率计算 |
| TEEP分析 | 全员生产效率分析 |
| 模具管理 | 模具寿命管理 |
| 量检具管理 | 计量器具校准管理 |

---

## 2. 页面清单

| 页面 | 路由路径 | 核心功能 |
|------|----------|----------|
| 设备台账 | `/equipment/archive` | 设备列表、新增/编辑/删除 |
| 设备详情 | `/equipment/archive/:id` | 设备详细信息查看 |
| 点检标准 | `/equipment/inspection-template-list` | 点检标准定义 |
| 点检计划 | `/equipment/inspection-plan-list` | 点检计划管理 |
| 今日点检 | `/equipment/inspection-today` | 今日待点检任务 |
| 点检执行 | `/equipment/inspection-execute/:planId` | PDA点检执行 |
| 异常记录 | `/equipment/inspection-defect-list` | 点检异常处理 |
| 点检统计 | `/equipment/inspection-statistics` | 点检完成率统计 |
| 保养计划 | `/equipment/maintenance-plan` | 保养计划管理 |
| 维修工单 | `/equipment/repair-order` | 维修单管理 |
| OEE报表 | `/equipment/oee` | 设备OEE分析 |
| 模具管理 | `/equipment/mold` | 模具寿命管理 |

---

## 3. UI设计规范

### 3.1 页面基本结构

同MES模块标准布局：搜索+工具栏+表格+详情弹窗。

### 3.2 状态映射

**设备状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| ACTIVE | success | 正常运行 |
| IDLE | warning | 闲置 |
| MAINTENANCE | warning | 保养中 |
| REPAIR | danger | 维修中 |
| SCRAPPED | info | 已报废 |

**点检结果状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| OK | success | 正常 |
| NG | danger | 异常 |
| NA | info | 不适用 |

**维修工单状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| PENDING | warning | 待派工 |
| ASSIGNED | primary | 已派工 |
| IN_PROGRESS | primary | 维修中 |
| COMPLETED | success | 已完成 |
| CANCELLED | info | 已取消 |

### 3.3 点检标准详情页

```vue
<template #footer>
  <el-button @click="detailVisible = false">关闭</el-button>
  <el-button type="primary" v-if="detailData.is_draft" @click="handlePublish">
    发布标准
  </el-button>
  <el-button type="danger" v-if="!detailData.is_current" @click="handleArchive">
    归档
  </el-button>
</template>
```

---

## 4. 业务流程

### 4.1 设备点检流程

```
创建点检标准 → 发布标准 → 生成点检计划 → 执行点检 → 记录结果 → 异常处理
```

### 4.2 设备维修流程

```
故障报修 → 维修派工 → 维修接单 → 维修处理 → 完工确认 → 评价归档
```

### 4.3 OEE计算

```
OEE = 时间稼动率 × 性能稼动率 × 良品率
时间稼动率 = (负荷时间 - 停机时间) / 负荷时间
性能稼动率 = 理论节拍 × 产出数量 / 实际运行时间
良品率 = 合格数量 / 产出数量
```

---

## 5. 数据模型

### 5.1 设备台账

| 字段 | 类型 | 说明 |
|------|------|------|
| equipment_code | VARCHAR(50) | 设备编码 |
| equipment_name | VARCHAR(100) | 设备名称 |
| equipment_type | VARCHAR(30) | 设备类型 |
| model | VARCHAR(100) | 型号 |
| manufacturer | VARCHAR(100) | 制造商 |
| serial_no | VARCHAR(100) | 序列号 |
| purchase_date | DATE | 购置日期 |
| status | VARCHAR(20) | 状态 |
| location_id | BIGINT | 位置ID |
| workshop_id | BIGINT | 所属车间 |

### 5.2 点检标准

| 字段 | 类型 | 说明 |
|------|------|------|
| template_code | VARCHAR(50) | 标准编码 |
| template_name | VARCHAR(200) | 标准名称 |
| template_type | VARCHAR(20) | DAILY/WEEKLY/MONTHLY |
| frequency_type | VARCHAR(20) | 执行频率 |
| is_required_signature | SMALLINT | 需要签名 |
| is_require_photo | SMALLINT | 需要拍照 |

### 5.3 点检计划

| 字段 | 类型 | 说明 |
|------|------|------|
| plan_no | VARCHAR(50) | 计划编号 |
| template_id | BIGINT | 点检标准ID |
| equipment_id | BIGINT | 设备ID |
| plan_date | DATE | 计划日期 |
| assigned_to | BIGINT | 指派人 |
| status | VARCHAR(20) | 状态 |

### 5.4 点检记录

| 字段 | 类型 | 说明 |
|------|------|------|
| record_no | VARCHAR(50) | 记录编号 |
| plan_id | BIGINT | 计划ID |
| equipment_id | BIGINT | 设备ID |
| inspector_id | BIGINT | 点检人 |
| inspection_start_time | TIMESTAMP | 开始时间 |
| overall_result | VARCHAR(20) | 总体结果 |
| ok_count | INTEGER | 合格项数 |
| ng_count | INTEGER | 不合格项数 |

### 5.5 维修工单

| 字段 | 类型 | 说明 |
|------|------|------|
| repair_no | VARCHAR(50) | 维修单号 |
| equipment_id | BIGINT | 设备ID |
| fault_type | VARCHAR(30) | 故障类型 |
| fault_desc | VARCHAR(500) | 故障描述 |
| reporter_id | BIGINT | 报修人 |
| assign_to | BIGINT | 维修人 |
| status | VARCHAR(20) | 状态 |
| repair_start_time | TIMESTAMP | 开始时间 |
| repair_end_time | TIMESTAMP | 结束时间 |

---

## 6. API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /equipment/archive/list | 设备台账列表 |
| GET | /equipment/archive/:id | 设备详情 |
| GET | /equipment/inspection-templates | 点检标准列表 |
| POST | /equipment/inspection-templates | 创建点检标准 |
| GET | /equipment/inspection-plans | 点检计划列表 |
| POST | /equipment/inspection-plans/generate | 生成点检计划 |
| GET | /equipment/inspection-records | 点检记录列表 |
| POST | /equipment/inspection-records | 开始点检 |
| PUT | /equipment/inspection-records/:id | 提交点检结果 |
| GET | /equipment/inspection-defects | 异常列表 |
| POST | /equipment/inspection-defects/:id/resolve | 处理异常 |
| GET | /equipment/maintenance/plans | 保养计划列表 |
| GET | /equipment/repair-orders | 维修工单列表 |
| POST | /equipment/repair-orders | 创建维修工单 |
| PUT | /equipment/repair-orders/:id/assign | 派工 |
| PUT | /equipment/repair-orders/:id/complete | 完工 |
| GET | /equipment/oee | OEE数据 |
| GET | /equipment/oee/statistics | OEE统计 |

---

## 6. 设备巡检管理

### 6.1 巡检记录主表 (basic_equipment_inspection_record_main)

设备巡检记录主表，记录每次巡检任务的基本信息。

```sql
CREATE TABLE `basic_equipment_inspection_record_main` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `record_no` varchar(50) NOT NULL COMMENT '巡检记录编号',
  `inspection_no` varchar(50) NOT NULL COMMENT '巡检单编号',
  `equipment_id` bigint NOT NULL COMMENT '设备ID',
  `equipment_code` varchar(50) DEFAULT NULL COMMENT '设备编码',
  `equipment_name` varchar(200) DEFAULT NULL COMMENT '设备名称',
  `inspector_id` bigint DEFAULT NULL COMMENT '巡检人ID',
  `inspector_name` varchar(100) DEFAULT NULL COMMENT '巡检人名称',
  `inspection_start_time` datetime DEFAULT NULL COMMENT '巡检开始时间',
  `inspection_end_time` datetime DEFAULT NULL COMMENT '巡检结束时间',
  `overall_result` varchar(20) DEFAULT NULL COMMENT '总体结果(OK/NG/PENDING)',
  `ok_count` int DEFAULT 0 COMMENT '合格项数',
  `ng_count` int DEFAULT 0 COMMENT '不合格项数',
  `na_count` int DEFAULT 0 COMMENT '不适用项数',
  `status` varchar(20) DEFAULT 'PENDING' COMMENT '状态(PENDING/IN_PROGRESS/COMPLETED/CANCELLED)',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_record_no` (`record_no`),
  KEY `idx_equipment_id` (`equipment_id`),
  KEY `idx_inspection_date` (`inspection_start_time`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备巡检记录主表';
```

### 6.2 巡检记录明细表 (basic_equipment_inspection_record_detail)

```sql
CREATE TABLE `basic_equipment_inspection_record_detail` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `record_id` bigint NOT NULL COMMENT '巡检记录主表ID',
  `inspection_item_id` bigint DEFAULT NULL COMMENT '巡检项ID',
  `item_code` varchar(50) DEFAULT NULL COMMENT '巡检项编码',
  `item_name` varchar(200) NOT NULL COMMENT '巡检项名称',
  `item_type` varchar(20) DEFAULT NULL COMMENT '巡检项类型(TEXT/NUMBER/SELECT/PHOTO)',
  `check_standard` varchar(500) DEFAULT NULL COMMENT '检查标准',
  `result_type` varchar(20) DEFAULT NULL COMMENT '结果类型(OK/NG/NA)',
  `result_value` varchar(200) DEFAULT NULL COMMENT '检查结果值',
  `result_remark` varchar(500) DEFAULT NULL COMMENT '结果备注',
  `photo_urls` text DEFAULT NULL COMMENT '拍照图片URLs(JSON数组)',
  `is_normal` bit(1) DEFAULT b'1' COMMENT '是否正常',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  KEY `idx_record_id` (`record_id`),
  KEY `idx_item_id` (`inspection_item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备巡检记录明细表';
```

### 6.3 巡检管理API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /equipment/inspection-record-main/create | 创建巡检记录主 |
| PUT | /equipment/inspection-record-main/update | 更新巡检记录主 |
| DELETE | /equipment/inspection-record-main/delete | 删除巡检记录主 |
| GET | /equipment/inspection-record-main/get | 获得巡检记录主 |
| GET | /equipment/inspection-record-main/page | 获得巡检记录主分页 |
| POST | /equipment/inspection-record-main/senior | 高级搜索获得巡检记录主分页 |
| GET | /equipment/inspection-record-main/export-excel | 导出巡检记录主Excel |
| POST | /equipment/inspection-record-main/import | 导入巡检记录主 |
| POST | /equipment/inspection-record-detail/create | 创建巡检记录子 |
| PUT | /equipment/inspection-record-detail/update | 更新巡检记录子 |
| DELETE | /equipment/inspection-record-detail/delete | 删除巡检记录子 |
| GET | /equipment/inspection-record-detail/get | 获得巡检记录子 |
| GET | /equipment/inspection-record-detail/list | 获得巡检记录子列表 |
| GET | /equipment/inspection-record-detail/page | 获得巡检记录子分页 |
| POST | /equipment/inspection-record-detail/senior | 高级搜索获得巡检记录子分页 |

---

## 7. 设备保养管理

### 7.1 保养记录主表 (basic_equipment_maintenance_main)

```sql
CREATE TABLE `basic_equipment_maintenance_main` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `record_no` varchar(50) NOT NULL COMMENT '保养记录编号',
  `maintenance_no` varchar(50) NOT NULL COMMENT '保养单编号',
  `equipment_id` bigint NOT NULL COMMENT '设备ID',
  `equipment_code` varchar(50) DEFAULT NULL COMMENT '设备编码',
  `equipment_name` varchar(200) DEFAULT NULL COMMENT '设备名称',
  `maintainer_id` bigint DEFAULT NULL COMMENT '保养人ID',
  `maintainer_name` varchar(100) DEFAULT NULL COMMENT '保养人名称',
  `maintenance_start_time` datetime DEFAULT NULL COMMENT '保养开始时间',
  `maintenance_end_time` datetime DEFAULT NULL COMMENT '保养结束时间',
  `maintenance_level` varchar(20) DEFAULT NULL COMMENT '保养级别(L1/L2/L3)',
  `overall_result` varchar(20) DEFAULT NULL COMMENT '总体结果(OK/NG/PENDING)',
  `ok_count` int DEFAULT 0 COMMENT '完成项数',
  `ng_count` int DEFAULT 0 COMMENT '异常项数',
  `status` varchar(20) DEFAULT 'PENDING' COMMENT '状态(PENDING/IN_PROGRESS/COMPLETED/CANCELLED)',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_record_no` (`record_no`),
  KEY `idx_equipment_id` (`equipment_id`),
  KEY `idx_maintenance_date` (`maintenance_start_time`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备保养记录主表';
```

### 7.2 保养记录明细表 (basic_equipment_maintenance_detail)

```sql
CREATE TABLE `basic_equipment_maintenance_detail` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `record_id` bigint NOT NULL COMMENT '保养记录主表ID',
  `maintenance_item_id` bigint DEFAULT NULL COMMENT '保养项ID',
  `item_code` varchar(50) DEFAULT NULL COMMENT '保养项编码',
  `item_name` varchar(200) NOT NULL COMMENT '保养项名称',
  `item_type` varchar(20) DEFAULT NULL COMMENT '保养项类型',
  `work_content` varchar(500) DEFAULT NULL COMMENT '工作内容',
  `result_type` varchar(20) DEFAULT NULL COMMENT '结果类型(OK/NG/NA)',
  `result_value` varchar(200) DEFAULT NULL COMMENT '保养结果值',
  `result_remark` varchar(500) DEFAULT NULL COMMENT '结果备注',
  `photo_urls` text DEFAULT NULL COMMENT '拍照图片URLs(JSON数组)',
  `is_completed` bit(1) DEFAULT b'0' COMMENT '是否完成',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  KEY `idx_record_id` (`record_id`),
  KEY `idx_item_id` (`maintenance_item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备保养记录明细表';
```

### 7.3 保养管理API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /equipment/maintenance-record-main/create | 创建保养记录主 |
| PUT | /equipment/maintenance-record-main/update | 更新保养记录主 |
| DELETE | /equipment/maintenance-record-main/delete | 删除保养记录主 |
| GET | /equipment/maintenance-record-main/get | 获得保养记录主 |
| GET | /equipment/maintenance-record-main/page | 获得保养记录主分页 |
| POST | /equipment/maintenance-record-main/senior | 高级搜索获得保养记录主分页 |
| GET | /equipment/maintenance-record-main/export-excel | 导出保养记录主Excel |
| POST | /equipment/maintenance-record-main/import | 导入保养记录主 |
| POST | /equipment/maintenance-record-detail/create | 创建保养记录子 |
| PUT | /equipment/maintenance-record-detail/update | 更新保养记录子 |
| DELETE | /equipment/maintenance-record-detail/delete | 删除保养记录子 |
| GET | /equipment/maintenance-record-detail/get | 获得保养记录子 |
| GET | /equipment/maintenance-record-detail/list | 获得保养记录子列表 |
| GET | /equipment/maintenance-record-detail/page | 获得保养记录子分页 |
| POST | /equipment/maintenance-record-detail/senior | 高级搜索获得保养记录子分页 |

---

## 8. 设备停机与故障管理

### 8.1 设备停机记录表 (basic_equipment_shutdown)

```sql
CREATE TABLE `basic_equipment_shutdown` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `shutdown_no` varchar(50) NOT NULL COMMENT '停机编号',
  `equipment_id` bigint NOT NULL COMMENT '设备ID',
  `equipment_code` varchar(50) DEFAULT NULL COMMENT '设备编码',
  `equipment_name` varchar(200) DEFAULT NULL COMMENT '设备名称',
  `shutdown_type` varchar(20) DEFAULT NULL COMMENT '停机类型(PLANNED/UNPLANNED)',
  `shutdown_reason` varchar(500) DEFAULT NULL COMMENT '停机原因',
  `fault_type_id` bigint DEFAULT NULL COMMENT '故障类型ID',
  `fault_type_name` varchar(100) DEFAULT NULL COMMENT '故障类型名称',
  `fault_cause_id` bigint DEFAULT NULL COMMENT '故障原因ID',
  `fault_cause_name` varchar(100) DEFAULT NULL COMMENT '故障原因名称',
  `shutdown_time` datetime NOT NULL COMMENT '停机开始时间',
  `restart_time` datetime DEFAULT NULL COMMENT '重启时间',
  `duration_minutes` int DEFAULT 0 COMMENT '停机时长(分钟)',
  `reporter_id` bigint DEFAULT NULL COMMENT '报告人ID',
  `reporter_name` varchar(100) DEFAULT NULL COMMENT '报告人名称',
  `handler_id` bigint DEFAULT NULL COMMENT '处理人ID',
  `handler_name` varchar(100) DEFAULT NULL COMMENT '处理人名称',
  `status` varchar(20) DEFAULT 'PENDING' COMMENT '状态(PENDING/PROCESSING/COMPLETED)',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_shutdown_no` (`shutdown_no`),
  KEY `idx_equipment_id` (`equipment_id`),
  KEY `idx_shutdown_time` (`shutdown_time`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备停机记录表';
```

### 8.2 设备转移记录表 (basic_equipment_transfer_record)

```sql
CREATE TABLE `basic_equipment_transfer_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `transfer_no` varchar(50) NOT NULL COMMENT '转移编号',
  `equipment_id` bigint NOT NULL COMMENT '设备ID',
  `equipment_code` varchar(50) DEFAULT NULL COMMENT '设备编码',
  `equipment_name` varchar(200) DEFAULT NULL COMMENT '设备名称',
  `from_workshop_id` bigint DEFAULT NULL COMMENT '源车间ID',
  `from_workshop_name` varchar(100) DEFAULT NULL COMMENT '源车间名称',
  `from_line_id` bigint DEFAULT NULL COMMENT '源产线ID',
  `from_line_name` varchar(100) DEFAULT NULL COMMENT '源产线名称',
  `from_location_id` bigint DEFAULT NULL COMMENT '源位置ID',
  `from_location_name` varchar(100) DEFAULT NULL COMMENT '源位置名称',
  `to_workshop_id` bigint DEFAULT NULL COMMENT '目标车间ID',
  `to_workshop_name` varchar(100) DEFAULT NULL COMMENT '目标车间名称',
  `to_line_id` bigint DEFAULT NULL COMMENT '目标产线ID',
  `to_line_name` varchar(100) DEFAULT NULL COMMENT '目标产线名称',
  `to_location_id` bigint DEFAULT NULL COMMENT '目标位置ID',
  `to_location_name` varchar(100) DEFAULT NULL COMMENT '目标位置名称',
  `transfer_date` datetime NOT NULL COMMENT '转移日期',
  `transfer_reason` varchar(500) DEFAULT NULL COMMENT '转移原因',
  `applicant_id` bigint DEFAULT NULL COMMENT '申请人ID',
  `applicant_name` varchar(100) DEFAULT NULL COMMENT '申请人名称',
  `handler_id` bigint DEFAULT NULL COMMENT '处理人ID',
  `handler_name` varchar(100) DEFAULT NULL COMMENT '处理人名称',
  `status` varchar(20) DEFAULT 'PENDING' COMMENT '状态(PENDING/APPROVED/COMPLETED/REJECTED)',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_transfer_no` (`transfer_no`),
  KEY `idx_equipment_id` (`equipment_id`),
  KEY `idx_transfer_date` (`transfer_date`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备转移记录表';
```

### 8.3 故障类型表 (basic_fault_type)

```sql
CREATE TABLE `basic_fault_type` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `fault_type_code` varchar(50) NOT NULL COMMENT '故障类型编码',
  `fault_type_name` varchar(100) NOT NULL COMMENT '故障类型名称',
  `fault_category` varchar(50) DEFAULT NULL COMMENT '故障类别',
  `description` varchar(500) DEFAULT NULL COMMENT '描述',
  `solution_hint` varchar(500) DEFAULT NULL COMMENT '解决方案提示',
  `is_enabled` bit(1) DEFAULT b'1' COMMENT '是否启用',
  `sort` int DEFAULT 0 COMMENT '排序',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_fault_type_code` (`fault_type_code`),
  KEY `idx_is_enabled` (`is_enabled`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='故障类型表';
```

### 8.4 故障原因表 (basic_fault_cause)

```sql
CREATE TABLE `basic_fault_cause` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `fault_cause_code` varchar(50) NOT NULL COMMENT '故障原因编码',
  `fault_cause_name` varchar(100) NOT NULL COMMENT '故障原因名称',
  `fault_type_id` bigint DEFAULT NULL COMMENT '故障类型ID',
  `fault_type_name` varchar(100) DEFAULT NULL COMMENT '故障类型名称',
  `description` varchar(500) DEFAULT NULL COMMENT '描述',
  `is_enabled` bit(1) DEFAULT b'1' COMMENT '是否启用',
  `sort` int DEFAULT 0 COMMENT '排序',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_fault_cause_code` (`fault_cause_code`),
  KEY `idx_fault_type_id` (`fault_type_id`),
  KEY `idx_is_enabled` (`is_enabled`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='故障原因表';
```

### 8.5 停机与故障管理API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /equipment/shutdown/create | 创建设备停机记录 |
| PUT | /equipment/shutdown/update | 更新设备停机记录 |
| DELETE | /equipment/shutdown/delete | 删除设备停机记录 |
| GET | /equipment/shutdown/get | 获得设备停机记录 |
| GET | /equipment/shutdown/page | 获得设备停机记录分页 |
| POST | /equipment/shutdown/senior | 高级搜索获得设备停机记录分页 |
| GET | /equipment/shutdown/export-excel | 导出设备停机记录Excel |
| POST | /equipment/shutdown/import | 导入设备停机记录 |
| POST | /equipment/transfer-record/create | 创建设备转移记录 |
| PUT | /equipment/transfer-record/update | 更新设备转移记录 |
| DELETE | /equipment/transfer-record/delete | 删除设备转移记录 |
| GET | /equipment/transfer-record/get | 获得设备转移记录 |
| GET | /equipment/transfer-record/page | 获得设备转移记录分页 |
| POST | /equipment/transfer-record/senior | 高级搜索获得设备转移记录分页 |
| GET | /equipment/transfer-record/export-excel | 导出设备转移记录Excel |
| POST | /equipment/transfer-record/import | 导入设备转移记录 |
| POST | /equipment/fault-type/create | 创建故障类型 |
| PUT | /equipment/fault-type/update | 更新故障类型 |
| DELETE | /equipment/fault-type/delete | 删除故障类型 |
| GET | /equipment/fault-type/get | 获得故障类型 |
| GET | /equipment/fault-type/page | 获得故障类型分页 |
| POST | /equipment/fault-type/senior | 高级搜索获得故障类型分页 |
| GET | /equipment/fault-type/export-excel | 导出故障类型Excel |
| POST | /equipment/fault-type/ables | 启用/禁用 |
| GET | /equipment/fault-type/noPage | 获得故障类型不分页 |
| POST | /equipment/fault-cause/create | 创建故障原因 |
| PUT | /equipment/fault-cause/update | 更新故障原因 |
| DELETE | /equipment/fault-cause/delete | 删除故障原因 |
| GET | /equipment/fault-cause/get | 获得故障原因 |
| GET | /equipment/fault-cause/page | 获得故障原因分页 |
| POST | /equipment/fault-cause/senior | 高级搜索获得故障原因分页 |
| GET | /equipment/fault-cause/export-excel | 导出故障原因Excel |
| POST | /equipment/fault-cause/ables | 启用/禁用 |

---

## 9. 备件与厂商管理

### 9.1 备件表 (basic_equipment_tool_spare_part)

```sql
CREATE TABLE `basic_equipment_tool_spare_part` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `spare_part_code` varchar(50) NOT NULL COMMENT '备件编码',
  `spare_part_name` varchar(200) NOT NULL COMMENT '备件名称',
  `spare_part_type` varchar(50) DEFAULT NULL COMMENT '备件类型',
  `specification` varchar(100) DEFAULT NULL COMMENT '规格型号',
  `unit` varchar(20) DEFAULT NULL COMMENT '单位',
  `min_stock` decimal(10,2) DEFAULT 0 COMMENT '最小库存',
  `max_stock` decimal(10,2) DEFAULT 0 COMMENT '最大库存',
  `current_stock` decimal(10,2) DEFAULT 0 COMMENT '当前库存',
  `warehouse_location` varchar(100) DEFAULT NULL COMMENT '仓库位置',
  `manufacturer_id` bigint DEFAULT NULL COMMENT '生产厂商ID',
  `manufacturer_name` varchar(100) DEFAULT NULL COMMENT '生产厂商名称',
  `supplier_id` bigint DEFAULT NULL COMMENT '供应商ID',
  `supplier_name` varchar(100) DEFAULT NULL COMMENT '供应商名称',
  `unit_price` decimal(10,2) DEFAULT 0 COMMENT '单价',
  `image_urls` text DEFAULT NULL COMMENT '图片URLs(JSON数组)',
  `description` varchar(500) DEFAULT NULL COMMENT '描述',
  `is_enabled` bit(1) DEFAULT b'1' COMMENT '是否启用',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_spare_part_code` (`spare_part_code`),
  KEY `idx_spare_part_type` (`spare_part_type`),
  KEY `idx_is_enabled` (`is_enabled`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备备件表';
```

### 9.2 设备厂商表 (basic_equipment_manufacturer)

```sql
CREATE TABLE `basic_equipment_manufacturer` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `manufacturer_code` varchar(50) NOT NULL COMMENT '厂商编码',
  `manufacturer_name` varchar(200) NOT NULL COMMENT '厂商名称',
  `manufacturer_type` varchar(50) DEFAULT NULL COMMENT '厂商类型(ORIGINAL/THIRD PARTY)',
  `contact_person` varchar(100) DEFAULT NULL COMMENT '联系人',
  `contact_phone` varchar(50) DEFAULT NULL COMMENT '联系电话',
  `contact_email` varchar(100) DEFAULT NULL COMMENT '联系邮箱',
  `address` varchar(500) DEFAULT NULL COMMENT '地址',
  `website` varchar(200) DEFAULT NULL COMMENT '网址',
  `description` varchar(500) DEFAULT NULL COMMENT '描述',
  `is_enabled` bit(1) DEFAULT b'1' COMMENT '是否启用',
  `sort` int DEFAULT 0 COMMENT '排序',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_manufacturer_code` (`manufacturer_code`),
  KEY `idx_is_enabled` (`is_enabled`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备厂商表';
```

### 9.3 设备供应商表 (basic_equipment_supplier)

```sql
CREATE TABLE `basic_equipment_supplier` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `supplier_code` varchar(50) NOT NULL COMMENT '供应商编码',
  `supplier_name` varchar(200) NOT NULL COMMENT '供应商名称',
  `supplier_type` varchar(50) DEFAULT NULL COMMENT '供应商类型',
  `contact_person` varchar(100) DEFAULT NULL COMMENT '联系人',
  `contact_phone` varchar(50) DEFAULT NULL COMMENT '联系电话',
  `contact_email` varchar(100) DEFAULT NULL COMMENT '联系邮箱',
  `address` varchar(500) DEFAULT NULL COMMENT '地址',
  `bank_name` varchar(100) DEFAULT NULL COMMENT '开户银行',
  `bank_account` varchar(100) DEFAULT NULL COMMENT '银行账号',
  `tax_number` varchar(100) DEFAULT NULL COMMENT '税号',
  `description` varchar(500) DEFAULT NULL COMMENT '描述',
  `is_enabled` bit(1) DEFAULT b'1' COMMENT '是否启用',
  `sort` int DEFAULT 0 COMMENT '排序',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_supplier_code` (`supplier_code`),
  KEY `idx_is_enabled` (`is_enabled`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备供应商表';
```

### 9.4 设备主要部件表 (basic_equipment_main_part)

```sql
CREATE TABLE `basic_equipment_main_part` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `part_code` varchar(50) NOT NULL COMMENT '部件编码',
  `part_name` varchar(200) NOT NULL COMMENT '部件名称',
  `equipment_id` bigint DEFAULT NULL COMMENT '关联设备ID',
  `equipment_code` varchar(50) DEFAULT NULL COMMENT '设备编码',
  `equipment_name` varchar(200) DEFAULT NULL COMMENT '设备名称',
  `part_type` varchar(50) DEFAULT NULL COMMENT '部件类型',
  `specification` varchar(100) DEFAULT NULL COMMENT '规格型号',
  `serial_no` varchar(100) DEFAULT NULL COMMENT '序列号',
  `lifecycle_count` int DEFAULT 0 COMMENT '设计寿命(次数)',
  `lifecycle_hours` decimal(10,2) DEFAULT 0 COMMENT '设计寿命(小时)',
  `used_count` int DEFAULT 0 COMMENT '已使用次数',
  `used_hours` decimal(10,2) DEFAULT 0 COMMENT '已使用小时',
  `replacement_cycle` int DEFAULT 0 COMMENT '更换周期(天)',
  `last_replacement_time` datetime DEFAULT NULL COMMENT '上次更换时间',
  `next_replacement_time` datetime DEFAULT NULL COMMENT '下次更换时间',
  `manufacturer_id` bigint DEFAULT NULL COMMENT '生产厂商ID',
  `manufacturer_name` varchar(100) DEFAULT NULL COMMENT '生产厂商名称',
  `unit_price` decimal(10,2) DEFAULT 0 COMMENT '单价',
  `is_enabled` bit(1) DEFAULT b'1' COMMENT '是否启用',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `deleted` bit(1) DEFAULT b'0' COMMENT '删除标志',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_part_code` (`part_code`),
  KEY `idx_equipment_id` (`equipment_id`),
  KEY `idx_is_enabled` (`is_enabled`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备主要部件表';
```

### 9.5 备件与厂商管理API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /equipment/spare-part/create | 创建备件 |
| PUT | /equipment/spare-part/update | 更新备件 |
| DELETE | /equipment/spare-part/delete | 删除备件 |
| GET | /equipment/spare-part/get | 获得备件 |
| GET | /equipment/spare-part/page | 获得备件分页 |
| POST | /equipment/spare-part/senior | 高级搜索获得备件分页 |
| GET | /equipment/spare-part/export-excel | 导出备件Excel |
| POST | /equipment/spare-part/ables | 启用/禁用 |
| GET | /equipment/spare-part/noPage | 获得备件不分页 |
| POST | /equipment/manufacturer/create | 创建设备厂商 |
| PUT | /equipment/manufacturer/update | 更新设备厂商 |
| DELETE | /equipment/manufacturer/delete | 删除设备厂商 |
| GET | /equipment/manufacturer/get | 获得设备厂商 |
| GET | /equipment/manufacturer/page | 获得设备厂商分页 |
| POST | /equipment/manufacturer/senior | 高级搜索获得设备厂商分页 |
| GET | /equipment/manufacturer/export-excel | 导出设备厂商Excel |
| POST | /equipment/manufacturer/ables | 启用/禁用 |
| GET | /equipment/manufacturer/noPage | 获得设备厂商不分页 |
| POST | /equipment/supplier/create | 创建供应商 |
| PUT | /equipment/supplier/update | 更新供应商 |
| DELETE | /equipment/supplier/delete | 删除供应商 |
| GET | /equipment/supplier/get | 获得供应商 |
| GET | /equipment/supplier/page | 获得供应商分页 |
| POST | /equipment/supplier/senior | 高级搜索获得供应商分页 |
| GET | /equipment/supplier/export-excel | 导出供应商Excel |
| GET | /equipment/supplier/noPage | 获得供应商不分页 |
| POST | /equipment/main-part/create | 创建设备主要部件 |
| PUT | /equipment/main-part/update | 更新设备主要部件 |
| DELETE | /equipment/main-part/delete | 删除设备主要部件 |
| GET | /equipment/main-part/get | 获得设备主要部件 |
| GET | /equipment/main-part/page | 获得设备主要部件分页 |
| POST | /equipment/main-part/senior | 高级搜索获得设备主要部件分页 |
| GET | /equipment/main-part/export-excel | 导出设备主要部件Excel |
| POST | /equipment/main-part/ables | 启用/禁用 |

---

## 10. 关联文档

- [MOM3.0_主设计文档](./MOM3.0_主设计文档.md) - 系统总览
- [MOM3.0_UI设计规范](./MOM3.0_UI设计规范.md) - UI规范详情
