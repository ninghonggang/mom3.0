# MOM3.0前端 EAM设备管理模块设计文档

**版本**: V2.0 | **所属模块**: M06设备管理 | **基于**: [MOM3.0_主设计文档](./MOM3.0_主设计文档.md)

---

## 1. 模块概述

### 1.1 功能定位

EAM设备管理模块覆盖设备台账、点检、保养、维修、备件、工具、产线车间全生命周期管理，实现设备资产的可控、可追溯、可优化。

### 1.2 核心功能

| 功能 | 说明 |
|------|------|
| 设备台账 | 设备主数据、分类、位置、供应商、制造商 |
| 设备部件 | 设备主要部件管理 |
| 设备图片 | 设备图片管理 |
| 设备文档 | 设备文档管理 |
| 设备证书 | 设备证书管理 |
| 点检管理 | 点检选项、选择集、标准、路线、周期、配置 |
| 保养管理 | 保养选项、选择集、标准、类型、周期、配置 |
| 维修管理 | 维修备件、经验库、标准、类型、流程、配置 |
| 设备运行 | 设备签到、转移、变更、运行记录 |
| 备件管理 | 备件申请、入库、出库、查询、盘点 |
| 工具管理 | 工具台账、型号、签到、出入库、库存、查询 |
| 产线车间 | 产线、车间、工位基础数据管理 |
| 基础数据 | 设备状态、故障类型、故障原因、维修级别 |

---

## 2. 页面清单

| 页面 | 路由路径 | 核心功能 |
|------|----------|----------|
| 设备台账 | `/eam/equipment` | 设备主数据、供应商、制造商 |
| 设备部件 | `/eam/equipmentPart` | 主要部件管理 |
| 设备供应商 | `/eam/equipmentSupplier` | 供应商管理 |
| 设备制造商 | `/eam/equipmentManufacturer` | 制造商管理 |
| 设备分类 | `/eam/equipmentClass` | 设备分类 |
| 设备图片 | `/eam/equipmentImage` | 设备图片 |
| 设备文档 | `/eam/equipmentDoc` | 设备文档 |
| 设备证书 | `/eam/equipmentCertificate` | 设备证书 |
| 点检选项 | `/eam/checkOption` | 点检选项定义 |
| 点检选择集 | `/eam/checkSet` | 点检选择集 |
| 点检标准 | `/eam/checkStd` | 点检标准 |
| 点检路线 | `/eam/checkRoute` | 点检路线 |
| 点检周期 | `/eam/checkCycle` | 点检周期 |
| 点检配置 | `/eam/checkConfig` | 点检配置 |
| 保养选项 | `/eam/maintenanceOption` | 保养选项定义 |
| 保养选择集 | `/eam/maintenanceSet` | 保养选择集 |
| 保养标准 | `/eam/maintenanceStd` | 保养标准 |
| 保养类型 | `/eam/maintenanceType` | 保养类型 |
| 保养周期 | `/eam/maintenanceCycle` | 保养周期 |
| 保养配置 | `/eam/maintenanceConfig` | 保养配置 |
| 维修备件 | `/eam/repairSpare` | 维修备件 |
| 维修经验 | `/eam/repairExp` | 维修经验库 |
| 维修标准 | `/eam/repairStd` | 维修标准 |
| 维修类型 | `/eam/repairType` | 维修类型 |
| 维修流程 | `/eam/repairFlow` | 维修流程 |
| 维修配置 | `/eam/repairConfig` | 维修配置 |
| 设备签到 | `/eam/checkin` | 设备签到 |
| 设备转移 | `/eam/transfer` | 设备转移 |
| 设备变更 | `/eam/change` | 设备变更 |
| 设备运行记录 | `/eam/runRecord` | 设备运行记录 |
| 备件申请 | `/eam/spareApply` | 备件申请 |
| 备件入库 | `/eam/spareInput` | 备件入库 |
| 备件出库 | `/eam/spareOutput` | 备件出库 |
| 备件查询 | `/eam/spareQuery` | 备件查询 |
| 备件盘点 | `/eam/spareCheck` | 备件盘点 |
| 工具台账 | `/eam/tool` | 工具台账 |
| 工具型号 | `/eam/toolType` | 工具型号 |
| 工具签到 | `/eam/toolCheckin` | 工具签到 |
| 工具出入库 | `/eam/toolInOut` | 工具出入库 |
| 工具库存 | `/eam/toolInventory` | 工具库存 |
| 工具查询 | `/eam/toolQuery` | 工具查询 |
| 产线基础 | `/eam/productionLine` | 产线档案 |
| 车间基础 | `/eam/workshop` | 车间档案 |
| 工位基础 | `/eam/workstation` | 工位档案 |
| 设备状态 | `/eam/equipmentStatus` | 设备状态 |
| 故障类型 | `/eam/faultType` | 故障类型 |
| 故障原因 | `/eam/faultReason` | 故障原因 |
| 维修级别 | `/eam/repairLevel` | 维修级别 |
| 设备分类 | `/eam/equipmentType` | 设备分类 |
| 设备图片 | `/eam/equipmentPhoto` | 设备图片 |
| 设备文档 | `/eam/equipmentDocument` | 设备文档 |

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

**备件状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| AVAILABLE | success | 可用 |
| RESERVED | warning | 已预约 |
| USED | info | 已使用 |
| SCRAPPED | danger | 已报废 |

---

## 4. 业务流程

### 4.1 设备点检流程

```
创建点检标准 → 发布标准 → 生成点检计划 → 执行点检 → 记录结果 → 异常处理
```

### 4.2 设备保养流程

```
创建保养标准 → 生成保养计划 → 执行保养 → 记录结果 → 完工确认
```

### 4.3 设备维修流程

```
故障报修 → 维修派工 → 维修接单 → 维修处理 → 完工确认 → 评价归档
```

### 4.4 备件管理流程

```
备件申请 → 审批 → 入库 → 出库 → 使用 → 盘点
```

---

## 5. 数据模型

### 5.1 设备台账 (eam_equipment)

| 字段 | 类型 | 说明 |
|------|------|------|
| equipment_id | BIGINT | 设备ID |
| equipment_code | VARCHAR(50) | 设备编码 |
| equipment_name | VARCHAR(200) | 设备名称 |
| equipment_type | VARCHAR(30) | 设备类型 |
| equipment_class_id | BIGINT | 设备分类ID |
| model | VARCHAR(100) | 型号 |
| specification | VARCHAR(200) | 规格 |
| manufacturer_id | BIGINT | 制造商ID |
| manufacturer_name | VARCHAR(200) | 制造商名称 |
| supplier_id | BIGINT | 供应商ID |
| supplier_name | VARCHAR(200) | 供应商名称 |
| serial_no | VARCHAR(100) | 序列号 |
| purchase_date | DATE | 购置日期 |
| purchase_amount | DECIMAL(18,2) | 购置金额 |
| service_life | INT | 使用寿命(年) |
| status | VARCHAR(20) | 状态 |
| location_id | BIGINT | 位置ID |
| location_name | VARCHAR(200) | 位置名称 |
| workshop_id | BIGINT | 所属车间 |
| workshop_name | VARCHAR(100) | 车间名称 |
| line_id | BIGINT | 所属产线 |
| line_name | VARCHAR(100) | 产线名称 |

### 5.2 设备部件 (eam_equipment_part)

| 字段 | 类型 | 说明 |
|------|------|------|
| part_id | BIGINT | 部件ID |
| equipment_id | BIGINT | 设备ID |
| part_code | VARCHAR(50) | 部件编码 |
| part_name | VARCHAR(200) | 部件名称 |
| part_type | VARCHAR(50) | 部件类型 |
| specification | VARCHAR(100) | 规格型号 |
| serial_no | VARCHAR(100) | 序列号 |
| lifecycle_count | INT | 设计寿命(次数) |
| lifecycle_hours | DECIMAL(10,2) | 设计寿命(小时) |
| used_count | INT | 已使用次数 |
| used_hours | DECIMAL(10,2) | 已使用小时 |
| replacement_cycle | INT | 更换周期(天) |
| last_replacement_time | DATETIME | 上次更换时间 |
| next_replacement_time | DATETIME | 下次更换时间 |
| manufacturer_id | BIGINT | 生产厂商ID |
| unit_price | DECIMAL(18,2) | 单价 |

### 5.3 设备供应商 (eam_equipment_supplier)

| 字段 | 类型 | 说明 |
|------|------|------|
| supplier_id | BIGINT | 供应商ID |
| equipment_id | BIGINT | 设备ID |
| supplier_code | VARCHAR(50) | 供应商编码 |
| supplier_name | VARCHAR(200) | 供应商名称 |
| supplier_type | VARCHAR(50) | 供应商类型 |
| contact_person | VARCHAR(100) | 联系人 |
| contact_phone | VARCHAR(50) | 联系电话 |
| contact_email | VARCHAR(100) | 联系邮箱 |
| address | VARCHAR(500) | 地址 |

### 5.4 设备制造商 (eam_equipment_manufacturer)

| 字段 | 类型 | 说明 |
|------|------|------|
| manufacturer_id | BIGINT | 制造商ID |
| equipment_id | BIGINT | 设备ID |
| manufacturer_code | VARCHAR(50) | 制造商编码 |
| manufacturer_name | VARCHAR(200) | 制造商名称 |
| manufacturer_type | VARCHAR(50) | 制造商类型 |
| contact_person | VARCHAR(100) | 联系人 |
| contact_phone | VARCHAR(50) | 联系电话 |
| website | VARCHAR(200) | 网址 |

### 5.5 设备分类 (eam_equipment_class)

| 字段 | 类型 | 说明 |
|------|------|------|
| class_id | BIGINT | 分类ID |
| class_code | VARCHAR(50) | 分类编码 |
| class_name | VARCHAR(200) | 分类名称 |
| parent_id | BIGINT | 上级分类ID |
| level | INT | 层级 |
| sort | INT | 排序 |

### 5.6 设备图片 (eam_equipment_image)

| 字段 | 类型 | 说明 |
|------|------|------|
| image_id | BIGINT | 图片ID |
| equipment_id | BIGINT | 设备ID |
| image_url | VARCHAR(500) | 图片URL |
| image_type | VARCHAR(20) | 图片类型 |
| is_primary | BIT | 是否主图 |
| sort | INT | 排序 |

### 5.7 设备文档 (eam_equipment_doc)

| 字段 | 类型 | 说明 |
|------|------|------|
| doc_id | BIGINT | 文档ID |
| equipment_id | BIGINT | 设备ID |
| doc_name | VARCHAR(200) | 文档名称 |
| doc_type | VARCHAR(50) | 文档类型 |
| doc_url | VARCHAR(500) | 文档URL |
| file_size | BIGINT | 文件大小 |
| upload_time | DATETIME | 上传时间 |

### 5.8 设备证书 (eam_equipment_certificate)

| 字段 | 类型 | 说明 |
|------|------|------|
| cert_id | BIGINT | 证书ID |
| equipment_id | BIGINT | 设备ID |
| cert_no | VARCHAR(50) | 证书编号 |
| cert_name | VARCHAR(200) | 证书名称 |
| cert_type | VARCHAR(50) | 证书类型 |
| issue_date | DATE | 发证日期 |
| expiry_date | DATE | 到期日期 |
| cert_url | VARCHAR(500) | 证书附件 |

### 5.9 点检选项 (eam_check_option)

| 字段 | 类型 | 说明 |
|------|------|------|
| option_id | BIGINT | 选项ID |
| option_code | VARCHAR(50) | 选项编码 |
| option_name | VARCHAR(200) | 选项名称 |
| option_type | VARCHAR(20) | 选项类型 |
| description | VARCHAR(500) | 描述 |
| is_enabled | BIT | 是否启用 |

### 5.10 点检选择集 (eam_check_set)

| 字段 | 类型 | 说明 |
|------|------|------|
| set_id | BIGINT | 选择集ID |
| set_code | VARCHAR(50) | 选择集编码 |
| set_name | VARCHAR(200) | 选择集名称 |
| option_list | TEXT | 选项列表(JSON) |
| description | VARCHAR(500) | 描述 |

### 5.11 点检标准 (eam_check_std)

| 字段 | 类型 | 说明 |
|------|------|------|
| std_id | BIGINT | 标准ID |
| std_code | VARCHAR(50) | 标准编码 |
| std_name | VARCHAR(200) | 标准名称 |
| std_type | VARCHAR(20) | 标准类型 |
| check_items | TEXT | 点检项(JSON) |
| is_required_signature | BIT | 需要签名 |
| is_require_photo | BIT | 需要拍照 |

### 5.12 点检路线 (eam_check_route)

| 字段 | 类型 | 说明 |
|------|------|------|
| route_id | BIGINT | 路线ID |
| route_code | VARCHAR(50) | 路线编码 |
| route_name | VARCHAR(200) | 路线名称 |
| equipment_ids | TEXT | 设备ID列表 |
| sequence | TEXT | 顺序 |
| total_points | INT | 总点数 |

### 5.13 点检周期 (eam_check_cycle)

| 字段 | 类型 | 说明 |
|------|------|------|
| cycle_id | BIGINT | 周期ID |
| cycle_code | VARCHAR(50) | 周期编码 |
| cycle_name | VARCHAR(200) | 周期名称 |
| cycle_type | VARCHAR(20) | 周期类型 |
| cycle_days | INT | 周期天数 |
| description | VARCHAR(500) | 描述 |

### 5.14 点检配置 (eam_check_config)

| 字段 | 类型 | 说明 |
|------|------|------|
| config_id | BIGINT | 配置ID |
| config_code | VARCHAR(50) | 配置编码 |
| config_name | VARCHAR(200) | 配置名称 |
| config_items | TEXT | 配置项(JSON) |
| is_enabled | BIT | 是否启用 |

### 5.15 保养选项 (eam_maintenance_option)

| 字段 | 类型 | 说明 |
|------|------|------|
| option_id | BIGINT | 选项ID |
| option_code | VARCHAR(50) | 选项编码 |
| option_name | VARCHAR(200) | 选项名称 |
| option_type | VARCHAR(20) | 选项类型 |
| work_content | VARCHAR(500) | 工作内容 |
| is_enabled | BIT | 是否启用 |

### 5.16 保养选择集 (eam_maintenance_set)

| 字段 | 类型 | 说明 |
|------|------|------|
| set_id | BIGINT | 选择集ID |
| set_code | VARCHAR(50) | 选择集编码 |
| set_name | VARCHAR(200) | 选择集名称 |
| option_list | TEXT | 选项列表(JSON) |
| description | VARCHAR(500) | 描述 |

### 5.17 保养标准 (eam_maintenance_std)

| 字段 | 类型 | 说明 |
|------|------|------|
| std_id | BIGINT | 标准ID |
| std_code | VARCHAR(50) | 标准编码 |
| std_name | VARCHAR(200) | 标准名称 |
| std_type | VARCHAR(20) | 标准类型 |
| maintenance_items | TEXT | 保养项(JSON) |
| duration_hours | DECIMAL(10,2) | 预计时长 |

### 5.18 保养类型 (eam_maintenance_type)

| 字段 | 类型 | 说明 |
|------|------|------|
| type_id | BIGINT | 类型ID |
| type_code | VARCHAR(50) | 类型编码 |
| type_name | VARCHAR(200) | 类型名称 |
| description | VARCHAR(500) | 描述 |
| is_enabled | BIT | 是否启用 |

### 5.19 保养周期 (eam_maintenance_cycle)

| 字段 | 类型 | 说明 |
|------|------|------|
| cycle_id | BIGINT | 周期ID |
| cycle_code | VARCHAR(50) | 周期编码 |
| cycle_name | VARCHAR(200) | 周期名称 |
| cycle_type | VARCHAR(20) | 周期类型 |
| cycle_days | INT | 周期天数 |
| description | VARCHAR(500) | 描述 |

### 5.20 保养配置 (eam_maintenance_config)

| 字段 | 类型 | 说明 |
|------|------|------|
| config_id | BIGINT | 配置ID |
| config_code | VARCHAR(50) | 配置编码 |
| config_name | VARCHAR(200) | 配置名称 |
| config_items | TEXT | 配置项(JSON) |
| is_enabled | BIT | 是否启用 |

### 5.21 维修备件 (eam_repair_spare)

| 字段 | 类型 | 说明 |
|------|------|------|
| spare_id | BIGINT | 备件ID |
| repair_id | BIGINT | 维修ID |
| spare_code | VARCHAR(50) | 备件编码 |
| spare_name | VARCHAR(200) | 备件名称 |
| quantity | INT | 数量 |
| unit_price | DECIMAL(18,2) | 单价 |

### 5.22 维修经验 (eam_repair_exp)

| 字段 | 类型 | 说明 |
|------|------|------|
| exp_id | BIGINT | 经验ID |
| fault_type_id | BIGINT | 故障类型ID |
| fault_desc | VARCHAR(500) | 故障描述 |
| fault_cause | TEXT | 故障原因 |
| solution | TEXT | 解决方案 |
| solution_steps | TEXT | 操作步骤 |
| duration_hours | DECIMAL(10,2) | 维修时长 |

### 5.23 维修标准 (eam_repair_std)

| 字段 | 类型 | 说明 |
|------|------|------|
| std_id | BIGINT | 标准ID |
| std_code | VARCHAR(50) | 标准编码 |
| std_name | VARCHAR(200) | 标准名称 |
| repair_type_id | BIGINT | 维修类型ID |
| std_content | TEXT | 标准内容 |
| duration_hours | DECIMAL(10,2) | 标准时长 |

### 5.24 维修类型 (eam_repair_type)

| 字段 | 类型 | 说明 |
|------|------|------|
| type_id | BIGINT | 类型ID |
| type_code | VARCHAR(50) | 类型编码 |
| type_name | VARCHAR(200) | 类型名称 |
| description | VARCHAR(500) | 描述 |
| is_enabled | BIT | 是否启用 |

### 5.25 维修流程 (eam_repair_flow)

| 字段 | 类型 | 说明 |
|------|------|------|
| flow_id | BIGINT | 流程ID |
| flow_code | VARCHAR(50) | 流程编码 |
| flow_name | VARCHAR(200) | 流程名称 |
| repair_type_id | BIGINT | 维修类型ID |
| flow_steps | TEXT | 流程步骤(JSON) |
| total_steps | INT | 总步骤数 |

### 5.26 维修配置 (eam_repair_config)

| 字段 | 类型 | 说明 |
|------|------|------|
| config_id | BIGINT | 配置ID |
| config_code | VARCHAR(50) | 配置编码 |
| config_name | VARCHAR(200) | 配置名称 |
| config_items | TEXT | 配置项(JSON) |
| is_enabled | BIT | 是否启用 |

### 5.27 设备签到 (eam_checkin)

| 字段 | 类型 | 说明 |
|------|------|------|
| checkin_id | BIGINT | 签到ID |
| equipment_id | BIGINT | 设备ID |
| checkin_time | DATETIME | 签到时间 |
| checkin_type | VARCHAR(20) | 签到类型 |
| operator_id | BIGINT | 操作人ID |
| operator_name | VARCHAR(100) | 操作人姓名 |
| remark | VARCHAR(500) | 备注 |

### 5.28 设备转移 (eam_transfer)

| 字段 | 类型 | 说明 |
|------|------|------|
| transfer_id | BIGINT | 转移ID |
| transfer_no | VARCHAR(50) | 转移单号 |
| equipment_id | BIGINT | 设备ID |
| from_location_id | BIGINT | 源位置ID |
| from_location_name | VARCHAR(200) | 源位置名称 |
| to_location_id | BIGINT | 目标位置ID |
| to_location_name | VARCHAR(200) | 目标位置名称 |
| transfer_date | DATETIME | 转移日期 |
| transfer_reason | VARCHAR(500) | 转移原因 |
| applicant_id | BIGINT | 申请人ID |
| status | VARCHAR(20) | 状态 |

### 5.29 设备变更 (eam_change)

| 字段 | 类型 | 说明 |
|------|------|------|
| change_id | BIGINT | 变更ID |
| change_no | VARCHAR(50) | 变更单号 |
| equipment_id | BIGINT | 设备ID |
| change_type | VARCHAR(20) | 变更类型 |
| change_content | TEXT | 变更内容 |
| change_date | DATETIME | 变更日期 |
| handler_id | BIGINT | 经办人ID |
| remark | VARCHAR(500) | 备注 |

### 5.30 设备运行记录 (eam_run_record)

| 字段 | 类型 | 说明 |
|------|------|------|
| record_id | BIGINT | 记录ID |
| equipment_id | BIGINT | 设备ID |
| start_time | DATETIME | 开始时间 |
| end_time | DATETIME | 结束时间 |
| run_hours | DECIMAL(10,2) | 运行时长 |
| idle_hours | DECIMAL(10,2) | 闲置时长 |
| output_qty | DECIMAL(18,3) | 产出数量 |
| pass_qty | DECIMAL(18,3) | 合格数量 |

### 5.31 备件申请 (eam_spare_apply)

| 字段 | 类型 | 说明 |
|------|------|------|
| apply_id | BIGINT | 申请ID |
| apply_no | VARCHAR(50) | 申请单号 |
| spare_code | VARCHAR(50) | 备件编码 |
| spare_name | VARCHAR(200) | 备件名称 |
| apply_qty | INT | 申请数量 |
| apply_reason | VARCHAR(500) | 申请原因 |
| applicant_id | BIGINT | 申请人ID |
| apply_time | DATETIME | 申请时间 |
| status | VARCHAR(20) | 状态 |

### 5.32 备件入库 (eam_spare_input)

| 字段 | 类型 | 说明 |
|------|------|------|
| input_id | BIGINT | 入库ID |
| input_no | VARCHAR(50) | 入库单号 |
| spare_code | VARCHAR(50) | 备件编码 |
| spare_name | VARCHAR(200) | 备件名称 |
| input_qty | INT | 入库数量 |
| unit_price | DECIMAL(18,2) | 单价 |
| supplier_id | BIGINT | 供应商ID |
| input_time | DATETIME | 入库时间 |
| handler_id | BIGINT | 经办人ID |

### 5.33 备件出库 (eam_spare_output)

| 字段 | 类型 | 说明 |
|------|------|------|
| output_id | BIGINT | 出库ID |
| output_no | VARCHAR(50) | 出库单号 |
| spare_code | VARCHAR(50) | 备件编码 |
| spare_name | VARCHAR(200) | 备件名称 |
| output_qty | INT | 出库数量 |
| use_type | VARCHAR(20) | 使用类型 |
| equipment_id | BIGINT | 设备ID |
| output_time | DATETIME | 出库时间 |
| handler_id | BIGINT | 经办人ID |

### 5.34 备件查询 (eam_spare_query)

| 字段 | 类型 | 说明 |
|------|------|------|
| query_type | VARCHAR(20) | 查询类型 |
| spare_code | VARCHAR(50) | 备件编码 |
| spare_name | VARCHAR(200) | 备件名称 |
| warehouse | VARCHAR(100) | 仓库 |
| min_stock | INT | 最小库存 |
| max_stock | INT | 最大库存 |

### 5.35 备件盘点 (eam_spare_check)

| 字段 | 类型 | 说明 |
|------|------|------|
| check_id | BIGINT | 盘点ID |
| check_no | VARCHAR(50) | 盘点单号 |
| check_time | DATETIME | 盘点时间 |
| checker_id | BIGINT | 盘点人ID |
| spare_list | TEXT | 备件列表(JSON) |
| result | TEXT | 盘点结果 |
| status | VARCHAR(20) | 状态 |

### 5.36 工具台账 (eam_tool)

| 字段 | 类型 | 说明 |
|------|------|------|
| tool_id | BIGINT | 工具ID |
| tool_code | VARCHAR(50) | 工具编码 |
| tool_name | VARCHAR(200) | 工具名称 |
| tool_type_id | BIGINT | 工具型号ID |
| specification | VARCHAR(100) | 规格 |
| unit | VARCHAR(20) | 单位 |
| current_stock | INT | 当前库存 |
| warehouse_location | VARCHAR(100) | 仓库位置 |
| status | VARCHAR(20) | 状态 |

### 5.37 工具型号 (eam_tool_type)

| 字段 | 类型 | 说明 |
|------|------|------|
| type_id | BIGINT | 型号ID |
| type_code | VARCHAR(50) | 型号编码 |
| type_name | VARCHAR(200) | 型号名称 |
| tool_category | VARCHAR(50) | 工具类别 |
| specification | VARCHAR(100) | 规格 |
| manufacturer | VARCHAR(100) | 制造商 |

### 5.38 工具签到 (eam_tool_checkin)

| 字段 | 类型 | 说明 |
|------|------|------|
| checkin_id | BIGINT | 签到ID |
| tool_id | BIGINT | 工具ID |
| checkin_time | DATETIME | 签到时间 |
| checkin_type | VARCHAR(20) | 签到类型 |
| operator_id | BIGINT | 操作人ID |
| operator_name | VARCHAR(100) | 操作人姓名 |

### 5.39 工具出入库 (eam_tool_inout)

| 字段 | 类型 | 说明 |
|------|------|------|
| record_id | BIGINT | 记录ID |
| tool_id | BIGINT | 工具ID |
| inout_type | VARCHAR(20) | 出入类型 |
| quantity | INT | 数量 |
| inout_time | DATETIME | 出入时间 |
| handler_id | BIGINT | 经办人ID |
| remark | VARCHAR(500) | 备注 |

### 5.40 工具库存 (eam_tool_inventory)

| 字段 | 类型 | 说明 |
|------|------|------|
| inventory_id | BIGINT | 库存ID |
| tool_id | BIGINT | 工具ID |
| warehouse_id | BIGINT | 仓库ID |
| warehouse_name | VARCHAR(100) | 仓库名称 |
| stock_qty | INT | 库存数量 |
| min_stock | INT | 最小库存 |
| max_stock | INT | 最大库存 |

### 5.41 工具查询 (eam_tool_query)

| 字段 | 类型 | 说明 |
|------|------|------|
| query_type | VARCHAR(20) | 查询类型 |
| tool_code | VARCHAR(50) | 工具编码 |
| tool_name | VARCHAR(200) | 工具名称 |
| tool_type | VARCHAR(50) | 工具类型 |
| warehouse | VARCHAR(100) | 仓库 |

### 5.42 产线基础 (eam_production_line)

| 字段 | 类型 | 说明 |
|------|------|------|
| line_id | BIGINT | 产线ID |
| line_code | VARCHAR(50) | 产线编码 |
| line_name | VARCHAR(200) | 产线名称 |
| workshop_id | BIGINT | 车间ID |
| line_type | VARCHAR(50) | 产线类型 |
| capacity | DECIMAL(18,3) | 产能 |
| status | VARCHAR(20) | 状态 |

### 5.43 车间基础 (eam_workshop)

| 字段 | 类型 | 说明 |
|------|------|------|
| workshop_id | BIGINT | 车间ID |
| workshop_code | VARCHAR(50) | 车间编码 |
| workshop_name | VARCHAR(200) | 车间名称 |
| factory_id | BIGINT | 工厂ID |
| area | DECIMAL(10,2) | 面积 |
| manager_id | BIGINT | 负责人ID |
| status | VARCHAR(20) | 状态 |

### 5.44 工位基础 (eam_workstation)

| 字段 | 类型 | 说明 |
|------|------|------|
| workstation_id | BIGINT | 工位ID |
| workstation_code | VARCHAR(50) | 工位编码 |
| workstation_name | VARCHAR(200) | 工位名称 |
| line_id | BIGINT | 产线ID |
| workshop_id | BIGINT | 车间ID |
| workstation_type | VARCHAR(50) | 工位类型 |
| status | VARCHAR(20) | 状态 |

### 5.45 设备状态 (eam_equipment_status)

| 字段 | 类型 | 说明 |
|------|------|------|
| status_id | BIGINT | 状态ID |
| status_code | VARCHAR(50) | 状态编码 |
| status_name | VARCHAR(100) | 状态名称 |
| status_color | VARCHAR(20) | 状态颜色 |
| is_running | BIT | 是否运行 |
| description | VARCHAR(500) | 描述 |

### 5.46 故障类型 (eam_fault_type)

| 字段 | 类型 | 说明 |
|------|------|------|
| type_id | BIGINT | 类型ID |
| fault_type_code | VARCHAR(50) | 故障类型编码 |
| fault_type_name | VARCHAR(100) | 故障类型名称 |
| fault_category | VARCHAR(50) | 故障类别 |
| solution_hint | VARCHAR(500) | 解决方案提示 |
| is_enabled | BIT | 是否启用 |

### 5.47 故障原因 (eam_fault_reason)

| 字段 | 类型 | 说明 |
|------|------|------|
| reason_id | BIGINT | 原因ID |
| fault_type_id | BIGINT | 故障类型ID |
| reason_code | VARCHAR(50) | 原因编码 |
| reason_desc | VARCHAR(500) | 原因描述 |
| is_enabled | BIT | 是否启用 |

### 5.48 维修级别 (eam_repair_level)

| 字段 | 类型 | 说明 |
|------|------|------|
| level_id | BIGINT | 级别ID |
| level_code | VARCHAR(50) | 级别编码 |
| level_name | VARCHAR(100) | 级别名称 |
| response_time | INT | 响应时间(小时) |
| repair_time | INT | 维修时间(小时) |
| is_enabled | BIT | 是否启用 |

### 5.49 设备分类 (eam_equipment_type)

| 字段 | 类型 | 说明 |
|------|------|------|
| type_id | BIGINT | 分类ID |
| type_code | VARCHAR(50) | 分类编码 |
| type_name | VARCHAR(200) | 分类名称 |
| parent_id | BIGINT | 上级分类ID |
| level | INT | 层级 |

### 5.50 设备图片 (eam_equipment_photo)

| 字段 | 类型 | 说明 |
|------|------|------|
| photo_id | BIGINT | 图片ID |
| equipment_id | BIGINT | 设备ID |
| photo_url | VARCHAR(500) | 图片URL |
| photo_type | VARCHAR(20) | 图片类型 |
| is_primary | BIT | 是否主图 |

### 5.51 设备文档 (eam_equipment_document)

| 字段 | 类型 | 说明 |
|------|------|------|
| document_id | BIGINT | 文档ID |
| equipment_id | BIGINT | 设备ID |
| doc_name | VARCHAR(200) | 文档名称 |
| doc_type | VARCHAR(50) | 文档类型 |
| doc_url | VARCHAR(500) | 文档URL |
| file_size | BIGINT | 文件大小 |

---

## 6. API接口

### 6.1 设备台账 (/eam/equipment)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/equipment/list | 设备列表 |
| GET | /eam/equipment/:id | 设备详情 |
| POST | /eam/equipment | 创建设备 |
| PUT | /eam/equipment/:id | 更新设备 |
| DELETE | /eam/equipment/:id | 删除设备 |
| GET | /eam/equipment/part | 主要部件 |
| GET | /eam/equipment/supplier | 供应商 |
| GET | /eam/equipment/manufacturer | 制造商 |

### 6.2 设备部件 (/eam/equipmentPart)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/equipmentPart/list | 部件列表 |
| GET | /eam/equipmentPart/:id | 部件详情 |
| POST | /eam/equipmentPart | 创建部件 |
| PUT | /eam/equipmentPart/:id | 更新部件 |
| DELETE | /eam/equipmentPart/:id | 删除部件 |

### 6.3 设备供应商 (/eam/equipmentSupplier)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/equipmentSupplier/list | 供应商列表 |
| GET | /eam/equipmentSupplier/:id | 供应商详情 |
| POST | /eam/equipmentSupplier | 创建供应商 |
| PUT | /eam/equipmentSupplier/:id | 更新供应商 |
| DELETE | /eam/equipmentSupplier/:id | 删除供应商 |

### 6.4 设备制造商 (/eam/equipmentManufacturer)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/equipmentManufacturer/list | 制造商列表 |
| GET | /eam/equipmentManufacturer/:id | 制造商详情 |
| POST | /eam/equipmentManufacturer | 创建制造商 |
| PUT | /eam/equipmentManufacturer/:id | 更新制造商 |
| DELETE | /eam/equipmentManufacturer/:id | 删除制造商 |

### 6.5 设备分类 (/eam/equipmentClass)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/equipmentClass/list | 分类列表 |
| GET | /eam/equipmentClass/:id | 分类详情 |
| POST | /eam/equipmentClass | 创建分类 |
| PUT | /eam/equipmentClass/:id | 更新分类 |
| DELETE | /eam/equipmentClass/:id | 删除分类 |
| GET | /eam/equipmentClass/tree | 分类树 |

### 6.6 设备图片 (/eam/equipmentImage)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/equipmentImage/list | 图片列表 |
| GET | /eam/equipmentImage/:id | 图片详情 |
| POST | /eam/equipmentImage | 上传图片 |
| PUT | /eam/equipmentImage/:id | 更新图片 |
| DELETE | /eam/equipmentImage/:id | 删除图片 |

### 6.7 设备文档 (/eam/equipmentDoc)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/equipmentDoc/list | 文档列表 |
| GET | /eam/equipmentDoc/:id | 文档详情 |
| POST | /eam/equipmentDoc | 上传文档 |
| PUT | /eam/equipmentDoc/:id | 更新文档 |
| DELETE | /eam/equipmentDoc/:id | 删除文档 |

### 6.8 设备证书 (/eam/equipmentCertificate)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/equipmentCertificate/list | 证书列表 |
| GET | /eam/equipmentCertificate/:id | 证书详情 |
| POST | /eam/equipmentCertificate | 创建证书 |
| PUT | /eam/equipmentCertificate/:id | 更新证书 |
| DELETE | /eam/equipmentCertificate/:id | 删除证书 |
| GET | /eam/equipmentCertificate/expiry | 证书到期提醒 |

### 6.9 点检选项 (/eam/checkOption)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/checkOption/list | 选项列表 |
| GET | /eam/checkOption/:id | 选项详情 |
| POST | /eam/checkOption | 创建选项 |
| PUT | /eam/checkOption/:id | 更新选项 |
| DELETE | /eam/checkOption/:id | 删除选项 |

### 6.10 点检选择集 (/eam/checkSet)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/checkSet/list | 选择集列表 |
| GET | /eam/checkSet/:id | 选择集详情 |
| POST | /eam/checkSet | 创建选择集 |
| PUT | /eam/checkSet/:id | 更新选择集 |
| DELETE | /eam/checkSet/:id | 删除选择集 |

### 6.11 点检标准 (/eam/checkStd)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/checkStd/list | 标准列表 |
| GET | /eam/checkStd/:id | 标准详情 |
| POST | /eam/checkStd | 创建标准 |
| PUT | /eam/checkStd/:id | 更新标准 |
| DELETE | /eam/checkStd/:id | 删除标准 |
| POST | /eam/checkStd/publish | 发布标准 |

### 6.12 点检路线 (/eam/checkRoute)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/checkRoute/list | 路线列表 |
| GET | /eam/checkRoute/:id | 路线详情 |
| POST | /eam/checkRoute | 创建路线 |
| PUT | /eam/checkRoute/:id | 更新路线 |
| DELETE | /eam/checkRoute/:id | 删除路线 |

### 6.13 点检周期 (/eam/checkCycle)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/checkCycle/list | 周期列表 |
| GET | /eam/checkCycle/:id | 周期详情 |
| POST | /eam/checkCycle | 创建周期 |
| PUT | /eam/checkCycle/:id | 更新周期 |
| DELETE | /eam/checkCycle/:id | 删除周期 |

### 6.14 点检配置 (/eam/checkConfig)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/checkConfig/list | 配置列表 |
| GET | /eam/checkConfig/:id | 配置详情 |
| POST | /eam/checkConfig | 创建配置 |
| PUT | /eam/checkConfig/:id | 更新配置 |
| DELETE | /eam/checkConfig/:id | 删除配置 |

### 6.15 保养选项 (/eam/maintenanceOption)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/maintenanceOption/list | 选项列表 |
| GET | /eam/maintenanceOption/:id | 选项详情 |
| POST | /eam/maintenanceOption | 创建选项 |
| PUT | /eam/maintenanceOption/:id | 更新选项 |
| DELETE | /eam/maintenanceOption/:id | 删除选项 |

### 6.16 保养选择集 (/eam/maintenanceSet)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/maintenanceSet/list | 选择集列表 |
| GET | /eam/maintenanceSet/:id | 选择集详情 |
| POST | /eam/maintenanceSet | 创建选择集 |
| PUT | /eam/maintenanceSet/:id | 更新选择集 |
| DELETE | /eam/maintenanceSet/:id | 删除选择集 |

### 6.17 保养标准 (/eam/maintenanceStd)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/maintenanceStd/list | 标准列表 |
| GET | /eam/maintenanceStd/:id | 标准详情 |
| POST | /eam/maintenanceStd | 创建标准 |
| PUT | /eam/maintenanceStd/:id | 更新标准 |
| DELETE | /eam/maintenanceStd/:id | 删除标准 |

### 6.18 保养类型 (/eam/maintenanceType)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/maintenanceType/list | 类型列表 |
| GET | /eam/maintenanceType/:id | 类型详情 |
| POST | /eam/maintenanceType | 创建类型 |
| PUT | /eam/maintenanceType/:id | 更新类型 |
| DELETE | /eam/maintenanceType/:id | 删除类型 |

### 6.19 保养周期 (/eam/maintenanceCycle)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/maintenanceCycle/list | 周期列表 |
| GET | /eam/maintenanceCycle/:id | 周期详情 |
| POST | /eam/maintenanceCycle | 创建周期 |
| PUT | /eam/maintenanceCycle/:id | 更新周期 |
| DELETE | /eam/maintenanceCycle/:id | 删除周期 |

### 6.20 保养配置 (/eam/maintenanceConfig)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/maintenanceConfig/list | 配置列表 |
| GET | /eam/maintenanceConfig/:id | 配置详情 |
| POST | /eam/maintenanceConfig | 创建配置 |
| PUT | /eam/maintenanceConfig/:id | 更新配置 |
| DELETE | /eam/maintenanceConfig/:id | 删除配置 |

### 6.21 维修备件 (/eam/repairSpare)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/repairSpare/list | 备件列表 |
| GET | /eam/repairSpare/:id | 备件详情 |
| POST | /eam/repairSpare | 创建备件 |
| PUT | /eam/repairSpare/:id | 更新备件 |
| DELETE | /eam/repairSpare/:id | 删除备件 |

### 6.22 维修经验 (/eam/repairExp)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/repairExp/list | 经验列表 |
| GET | /eam/repairExp/:id | 经验详情 |
| POST | /eam/repairExp | 创建经验 |
| PUT | /eam/repairExp/:id | 更新经验 |
| DELETE | /eam/repairExp/:id | 删除经验 |
| GET | /eam/repairExp/search | 经验搜索 |

### 6.23 维修标准 (/eam/repairStd)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/repairStd/list | 标准列表 |
| GET | /eam/repairStd/:id | 标准详情 |
| POST | /eam/repairStd | 创建标准 |
| PUT | /eam/repairStd/:id | 更新标准 |
| DELETE | /eam/repairStd/:id | 删除标准 |

### 6.24 维修类型 (/eam/repairType)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/repairType/list | 类型列表 |
| GET | /eam/repairType/:id | 类型详情 |
| POST | /eam/repairType | 创建类型 |
| PUT | /eam/repairType/:id | 更新类型 |
| DELETE | /eam/repairType/:id | 删除类型 |

### 6.25 维修流程 (/eam/repairFlow)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/repairFlow/list | 流程列表 |
| GET | /eam/repairFlow/:id | 流程详情 |
| POST | /eam/repairFlow | 创建流程 |
| PUT | /eam/repairFlow/:id | 更新流程 |
| DELETE | /eam/repairFlow/:id | 删除流程 |

### 6.26 维修配置 (/eam/repairConfig)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/repairConfig/list | 配置列表 |
| GET | /eam/repairConfig/:id | 配置详情 |
| POST | /eam/repairConfig | 创建配置 |
| PUT | /eam/repairConfig/:id | 更新配置 |
| DELETE | /eam/repairConfig/:id | 删除配置 |

### 6.27 设备签到 (/eam/checkin)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/checkin/list | 签到列表 |
| GET | /eam/checkin/:id | 签到详情 |
| POST | /eam/checkin | 创建签到 |
| PUT | /eam/checkin/:id | 更新签到 |
| DELETE | /eam/checkin/:id | 删除签到 |

### 6.28 设备转移 (/eam/transfer)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/transfer/list | 转移列表 |
| GET | /eam/transfer/:id | 转移详情 |
| POST | /eam/transfer | 创建转移 |
| PUT | /eam/transfer/:id | 更新转移 |
| DELETE | /eam/transfer/:id | 删除转移 |
| POST | /eam/transfer/approve | 审批转移 |

### 6.29 设备变更 (/eam/change)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/change/list | 变更列表 |
| GET | /eam/change/:id | 变更详情 |
| POST | /eam/change | 创建变更 |
| PUT | /eam/change/:id | 更新变更 |
| DELETE | /eam/change/:id | 删除变更 |

### 6.30 设备运行记录 (/eam/runRecord)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/runRecord/list | 记录列表 |
| GET | /eam/runRecord/:id | 记录详情 |
| POST | /eam/runRecord | 创建记录 |
| PUT | /eam/runRecord/:id | 更新记录 |
| DELETE | /eam/runRecord/:id | 删除记录 |
| GET | /eam/runRecord/statistics | 运行统计 |

### 6.31 备件申请 (/eam/spareApply)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/spareApply/list | 申请列表 |
| GET | /eam/spareApply/:id | 申请详情 |
| POST | /eam/spareApply | 创建申请 |
| PUT | /eam/spareApply/:id | 更新申请 |
| DELETE | /eam/spareApply/:id | 删除申请 |
| POST | /eam/spareApply/approve | 审批申请 |

### 6.32 备件入库 (/eam/spareInput)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/spareInput/list | 入库列表 |
| GET | /eam/spareInput/:id | 入库详情 |
| POST | /eam/spareInput | 创建入库 |
| PUT | /eam/spareInput/:id | 更新入库 |
| DELETE | /eam/spareInput/:id | 删除入库 |

### 6.33 备件出库 (/eam/spareOutput)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/spareOutput/list | 出库列表 |
| GET | /eam/spareOutput/:id | 出库详情 |
| POST | /eam/spareOutput | 创建出库 |
| PUT | /eam/spareOutput/:id | 更新出库 |
| DELETE | /eam/spareOutput/:id | 删除出库 |

### 6.34 备件查询 (/eam/spareQuery)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/spareQuery/list | 查询列表 |
| GET | /eam/spareQuery/:id | 查询详情 |
| GET | /eam/spareQuery/stock | 库存查询 |
| GET | /eam/spareQuery/track | 跟踪查询 |

### 6.35 备件盘点 (/eam/spareCheck)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/spareCheck/list | 盘点列表 |
| GET | /eam/spareCheck/:id | 盘点详情 |
| POST | /eam/spareCheck | 创建盘点 |
| PUT | /eam/spareCheck/:id | 更新盘点 |
| DELETE | /eam/spareCheck/:id | 删除盘点 |
| POST | /eam/spareCheck/result | 盘点结果 |

### 6.36 工具台账 (/eam/tool)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/tool/list | 工具列表 |
| GET | /eam/tool/:id | 工具详情 |
| POST | /eam/tool | 创建工具 |
| PUT | /eam/tool/:id | 更新工具 |
| DELETE | /eam/tool/:id | 删除工具 |

### 6.37 工具型号 (/eam/toolType)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/toolType/list | 型号列表 |
| GET | /eam/toolType/:id | 型号详情 |
| POST | /eam/toolType | 创建型号 |
| PUT | /eam/toolType/:id | 更新型号 |
| DELETE | /eam/toolType/:id | 删除型号 |

### 6.38 工具签到 (/eam/toolCheckin)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/toolCheckin/list | 签到列表 |
| GET | /eam/toolCheckin/:id | 签到详情 |
| POST | /eam/toolCheckin | 创建签到 |
| PUT | /eam/toolCheckin/:id | 更新签到 |
| DELETE | /eam/toolCheckin/:id | 删除签到 |

### 6.39 工具出入库 (/eam/toolInOut)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/toolInOut/list | 记录列表 |
| GET | /eam/toolInOut/:id | 记录详情 |
| POST | /eam/toolInOut | 创建记录 |
| PUT | /eam/toolInOut/:id | 更新记录 |
| DELETE | /eam/toolInOut/:id | 删除记录 |

### 6.40 工具库存 (/eam/toolInventory)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/toolInventory/list | 库存列表 |
| GET | /eam/toolInventory/:id | 库存详情 |
| POST | /eam/toolInventory | 创建库存 |
| PUT | /eam/toolInventory/:id | 更新库存 |
| DELETE | /eam/toolInventory/:id | 删除库存 |
| GET | /eam/toolInventory/warning | 库存预警 |

### 6.41 工具查询 (/eam/toolQuery)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/toolQuery/list | 查询列表 |
| GET | /eam/toolQuery/stock | 库存查询 |
| GET | /eam/toolQuery/track | 跟踪查询 |

### 6.42 产线基础 (/eam/productionLine)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/productionLine/list | 产线列表 |
| GET | /eam/productionLine/:id | 产线详情 |
| POST | /eam/productionLine | 创建产线 |
| PUT | /eam/productionLine/:id | 更新产线 |
| DELETE | /eam/productionLine/:id | 删除产线 |

### 6.43 车间基础 (/eam/workshop)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/workshop/list | 车间列表 |
| GET | /eam/workshop/:id | 车间详情 |
| POST | /eam/workshop | 创建车间 |
| PUT | /eam/workshop/:id | 更新车间 |
| DELETE | /eam/workshop/:id | 删除车间 |

### 6.44 工位基础 (/eam/workstation)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/workstation/list | 工位列表 |
| GET | /eam/workstation/:id | 工位详情 |
| POST | /eam/workstation | 创建工位 |
| PUT | /eam/workstation/:id | 更新工位 |
| DELETE | /eam/workstation/:id | 删除工位 |

### 6.45 设备状态 (/eam/equipmentStatus)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/equipmentStatus/list | 状态列表 |
| GET | /eam/equipmentStatus/:id | 状态详情 |
| POST | /eam/equipmentStatus | 创建状态 |
| PUT | /eam/equipmentStatus/:id | 更新状态 |
| DELETE | /eam/equipmentStatus/:id | 删除状态 |

### 6.46 故障类型 (/eam/faultType)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/faultType/list | 类型列表 |
| GET | /eam/faultType/:id | 类型详情 |
| POST | /eam/faultType | 创建类型 |
| PUT | /eam/faultType/:id | 更新类型 |
| DELETE | /eam/faultType/:id | 删除类型 |
| POST | /eam/faultType/enable | 启用/禁用 |

### 6.47 故障原因 (/eam/faultReason)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/faultReason/list | 原因列表 |
| GET | /eam/faultReason/:id | 原因详情 |
| POST | /eam/faultReason | 创建原因 |
| PUT | /eam/faultReason/:id | 更新原因 |
| DELETE | /eam/faultReason/:id | 删除原因 |
| POST | /eam/faultReason/enable | 启用/禁用 |

### 6.48 维修级别 (/eam/repairLevel)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/repairLevel/list | 级别列表 |
| GET | /eam/repairLevel/:id | 级别详情 |
| POST | /eam/repairLevel | 创建级别 |
| PUT | /eam/repairLevel/:id | 更新级别 |
| DELETE | /eam/repairLevel/:id | 删除级别 |

### 6.49 设备分类 (/eam/equipmentType)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/equipmentType/list | 分类列表 |
| GET | /eam/equipmentType/:id | 分类详情 |
| POST | /eam/equipmentType | 创建分类 |
| PUT | /eam/equipmentType/:id | 更新分类 |
| DELETE | /eam/equipmentType/:id | 删除分类 |
| GET | /eam/equipmentType/tree | 分类树 |

### 6.50 设备图片 (/eam/equipmentPhoto)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/equipmentPhoto/list | 图片列表 |
| GET | /eam/equipmentPhoto/:id | 图片详情 |
| POST | /eam/equipmentPhoto | 上传图片 |
| PUT | /eam/equipmentPhoto/:id | 更新图片 |
| DELETE | /eam/equipmentPhoto/:id | 删除图片 |

### 6.51 设备文档 (/eam/equipmentDocument)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/equipmentDocument/list | 文档列表 |
| GET | /eam/equipmentDocument/:id | 文档详情 |
| POST | /eam/equipmentDocument | 上传文档 |
| PUT | /eam/equipmentDocument/:id | 更新文档 |
| DELETE | /eam/equipmentDocument/:id | 删除文档 |

### 6.52 设备巡检 (/eam/inspection)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/inspection/plan/list | 巡检计划列表 |
| GET | /eam/inspection/plan/:id | 巡检计划详情 |
| POST | /eam/inspection/plan | 创建巡检计划 |
| PUT | /eam/inspection/plan/:id | 更新巡检计划 |
| DELETE | /eam/inspection/plan/:id | 删除巡检计划 |
| POST | /eam/inspection/plan/publish | 发布巡检计划 |
| GET | /eam/inspection/record/list | 巡检记录列表 |
| GET | /eam/inspection/record/:id | 巡检记录详情 |
| POST | /eam/inspection/record | 创建巡检记录 |
| PUT | /eam/inspection/record/:id | 更新巡检记录 |
| GET | /eam/inspection/record/detail/:id | 巡检记录明细 |
| GET | /eam/inspection/option/list | 巡检选项列表 |
| GET | /eam/inspection/option/:id | 巡检选项详情 |
| POST | /eam/inspection/option | 创建巡检选项 |
| PUT | /eam/inspection/option/:id | 更新巡检选项 |
| DELETE | /eam/inspection/option/:id | 删除巡检选项 |
| GET | /eam/inspection/set/list | 巡检选择集列表 |
| GET | /eam/inspection/set/:id | 巡检选择集详情 |
| POST | /eam/inspection/set | 创建巡检选择集 |
| PUT | /eam/inspection/set/:id | 更新巡检选择集 |
| DELETE | /eam/inspection/set/:id | 删除巡检选择集 |
| GET | /eam/inspection/item/list | 巡检项目列表 |
| GET | /eam/inspection/item/:id | 巡检项目详情 |
| POST | /eam/inspection/item | 创建巡检项目 |
| PUT | /eam/inspection/item/:id | 更新巡检项目 |
| DELETE | /eam/inspection/item/:id | 删除巡检项目 |

### 6.53 设备保养 (/eam/maintenance)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/maintenance/plan/list | 保养计划列表 |
| GET | /eam/maintenance/plan/:id | 保养计划详情 |
| POST | /eam/maintenance/plan | 创建保养计划 |
| PUT | /eam/maintenance/plan/:id | 更新保养计划 |
| DELETE | /eam/maintenance/plan/:id | 删除保养计划 |
| POST | /eam/maintenance/plan/publish | 发布保养计划 |
| GET | /eam/maintenance/record/list | 保养记录列表 |
| GET | /eam/maintenance/record/:id | 保养记录详情 |
| POST | /eam/maintenance/record | 创建保养记录 |
| PUT | /eam/maintenance/record/:id | 更新保养记录 |
| GET | /eam/maintenance/record/detail/:id | 保养记录明细 |
| GET | /eam/maintenance/option/list | 保养选项列表 |
| GET | /eam/maintenance/option/:id | 保养选项详情 |
| POST | /eam/maintenance/option | 创建保养选项 |
| PUT | /eam/maintenance/option/:id | 更新保养选项 |
| DELETE | /eam/maintenance/option/:id | 删除保养选项 |
| GET | /eam/maintenance/set/list | 保养选择集列表 |
| GET | /eam/maintenance/set/:id | 保养选择集详情 |
| POST | /eam/maintenance/set | 创建保养选择集 |
| PUT | /eam/maintenance/set/:id | 更新保养选择集 |
| DELETE | /eam/maintenance/set/:id | 删除保养选择集 |
| GET | /eam/maintenance/item/list | 保养项目列表 |
| GET | /eam/maintenance/item/:id | 保养项目详情 |
| POST | /eam/maintenance/item | 创建保养项目 |
| PUT | /eam/maintenance/item/:id | 更新保养项目 |
| DELETE | /eam/maintenance/item/:id | 删除保养项目 |
| GET | /eam/maintenance/experience/list | 保养经验列表 |
| GET | /eam/maintenance/experience/:id | 保养经验详情 |
| POST | /eam/maintenance/experience | 创建保养经验 |
| PUT | /eam/maintenance/experience/:id | 更新保养经验 |
| DELETE | /eam/maintenance/experience/:id | 删除保养经验 |

### 6.54 设备停机 (/eam/shutdown)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/shutdown/list | 停机记录列表 |
| GET | /eam/shutdown/:id | 停机记录详情 |
| POST | /eam/shutdown | 创建停机记录 |
| PUT | /eam/shutdown/:id | 更新停机记录 |
| DELETE | /eam/shutdown/:id | 删除停机记录 |
| POST | /eam/shutdown/start | 开始停机 |
| POST | /eam/shutdown/end | 结束停机 |
| GET | /eam/shutdown/statistics | 停机统计 |
| GET | /eam/shutdown/analysis | 停机分析 |

### 6.55 设备转移 (/eam/transfer)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/transfer/list | 转移记录列表 |
| GET | /eam/transfer/:id | 转移记录详情 |
| POST | /eam/transfer | 创建转移记录 |
| PUT | /eam/transfer/:id | 更新转移记录 |
| DELETE | /eam/transfer/:id | 删除转移记录 |
| POST | /eam/transfer/approve | 审批转移 |
| POST | /eam/transfer/execute | 执行转移 |
| GET | /eam/transfer/track | 转移追踪 |

### 6.56 备件管理 (/eam/spare)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/spare/list | 备件列表 |
| GET | /eam/spare/:id | 备件详情 |
| POST | /eam/spare | 创建备件 |
| PUT | /eam/spare/:id | 更新备件 |
| DELETE | /eam/spare/:id | 删除备件 |
| GET | /eam/spare/stock | 库存查询 |
| POST | /eam/spare/apply | 备件申请 |
| POST | /eam/spare/input | 备件入库 |
| POST | /eam/spare/output | 备件出库 |
| GET | /eam/spare/track | 使用追踪 |
| GET | /eam/spare/location/list | 备件位置列表 |
| GET | /eam/spare/location/:id | 备件位置详情 |
| POST | /eam/spare/location | 创建备件位置 |
| PUT | /eam/spare/location/:id | 更新备件位置 |
| DELETE | /eam/spare/location/:id | 删除备件位置 |
| GET | /eam/spare/count/list | 备件盘点列表 |
| GET | /eam/spare/count/:id | 备件盘点详情 |
| POST | /eam/spare/count | 创建备件盘点 |
| PUT | /eam/spare/count/:id | 更新备件盘点 |
| POST | /eam/spare/count/result | 盘点结果确认 |

### 6.57 设备厂商 (/eam/manufacturer)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/manufacturer/list | 厂商列表 |
| GET | /eam/manufacturer/:id | 厂商详情 |
| POST | /eam/manufacturer | 创建厂商 |
| PUT | /eam/manufacturer/:id | 更新厂商 |
| DELETE | /eam/manufacturer/:id | 删除厂商 |
| POST | /eam/manufacturer/enable | 启用/禁用 |
| GET | /eam/manufacturer/noPage | 厂商不分页列表 |

### 6.58 设备供应商 (/eam/supplier)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/supplier/list | 供应商列表 |
| GET | /eam/supplier/:id | 供应商详情 |
| POST | /eam/supplier | 创建供应商 |
| PUT | /eam/supplier/:id | 更新供应商 |
| DELETE | /eam/supplier/:id | 删除供应商 |
| POST | /eam/supplier/enable | 启用/禁用 |
| GET | /eam/supplier/noPage | 供应商不分页列表 |

### 6.59 设备主要部件 (/eam/mainPart)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /eam/mainPart/list | 主要部件列表 |
| GET | /eam/mainPart/:id | 主要部件详情 |
| POST | /eam/mainPart | 创建主要部件 |
| PUT | /eam/mainPart/:id | 更新主要部件 |
| DELETE | /eam/mainPart/:id | 删除主要部件 |
| GET | /eam/mainPart/byEquipment/:equipmentId | 按设备查询部件 |
| POST | /eam/mainPart/replacement | 记录部件更换 |

---

## 7. 业务功能说明

### 7.1 设备巡检管理

#### 7.1.1 巡检计划管理
- **功能描述**: 根据设备类型、位置、重要性等维度制定周期性巡检计划
- **流程**: 创建巡检计划 -> 关联设备/路线 -> 设置巡检项 -> 发布计划 -> 自动生成巡检任务
- **配置项**: 巡检路线、巡检周期、巡检人员、巡检项选择集

#### 7.1.2 巡检执行管理
- **功能描述**: 巡检人员按计划执行巡检任务，记录设备状态
- **流程**: 接收巡检任务 -> 到达现场 -> 逐项点检 -> 记录结果(正常/异常/不适用) -> 上传照片 -> 签名确认
- **异常处理**: 发现异常自动生成报修工单

#### 7.1.3 巡检记录查询
- **功能描述**: 查询历史巡检记录，支持按设备、时间、人员等维度筛选
- **统计报表**: 巡检完成率、异常率、设备健康度评分

### 7.2 设备保养管理

#### 7.2.1 保养计划管理
- **功能描述**: 根据设备手册、运行小时数等制定预防性保养计划
- **流程**: 创建保养计划 -> 关联设备 -> 设置保养级别/项目 -> 发布计划 -> 自动生成保养任务
- **保养级别**: 一级保养(日常)、二级保养(周/月)、三级保养(季度/年)

#### 7.2.2 保养执行管理
- **功能描述**: 维修人员按计划执行保养任务，记录保养内容
- **流程**: 接收保养任务 -> 准备工作 -> 执行保养项 -> 记录更换配件 -> 完工确认 -> 签名
- **经验积累**: 保养完成后可记录保养经验，供后续参考

#### 7.2.3 保养记录查询
- **功能描述**: 查询历史保养记录，支持按设备、时间、人员等维度筛选
- **统计报表**: 保养完成率、配件消耗统计、设备MTBF分析

### 7.3 设备停机记录

#### 7.3.1 停机登记
- **功能描述**: 记录设备非计划停机事件
- **登记项**: 设备信息、停机时间、停机原因、影响产线、紧急程度
- **关联**: 自动关联报修工单、维修记录

#### 7.3.2 停机分析
- **功能描述**: 分析停机原因、停机时长，为设备改进提供数据支持
- **分析维度**: 停机次数、停机时长、MTBF、MTTR、停机原因分布

### 7.4 设备转移记录

#### 7.4.1 转移申请
- **功能描述**: 记录设备在不同产线、车间、工位之间的调拨
- **申请项**: 设备信息、源位置、目标位置、转移原因、预计时间

#### 7.4.2 转移审批与执行
- **流程**: 提交转移申请 -> 审批确认 -> 执行转移 -> 更新设备位置
- **追踪**: 全程记录转移历史，支持追溯

### 7.5 备件管理

#### 7.5.1 备件档案
- **功能描述**: 建立备件库存档案，记录备件基本信息、关联设备、安全库存
- **管理项**: 备件编码、名称、规格、供应商、最低/最高库存、库位

#### 7.5.2 备件出入库
- **入库流程**: 采购入库 -> 检验 -> 登记入库单 -> 入库确认 -> 更新库存
- **出库流程**: 申请出库 -> 审批 -> 登记出库单 -> 出库确认 -> 更新库存
- **追溯**: 记录每笔出入库的来源、去向、审批人

#### 7.5.3 备件盘点
- **功能描述**: 定期盘点备件库存，核对系统库存与实际库存
- **流程**: 创建盘点计划 -> 执行盘点 -> 录入实盘数 -> 差异分析 -> 调整库存

### 7.6 故障分析

#### 7.6.1 故障类型管理
- **功能描述**: 定义设备故障分类体系
- **分类维度**: 故障部位、故障现象、紧急程度

#### 7.6.2 故障原因管理
- **功能描述**: 关联故障类型，记录可能原因
- **原因层次**: 直接原因、间接原因、根因分析

#### 7.6.3 故障统计分析
- **分析维度**: 故障类型分布、故障原因占比、MTBF趋势、设备故障排名

---

## 8. DDL表结构

### 8.1 巡检记录主表 (basic_equipment_inspection_record_main)

```sql
CREATE TABLE `basic_equipment_inspection_record_main` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `inspection_no` varchar(50) NOT NULL COMMENT '巡检单号',
  `equipment_id` bigint NOT NULL COMMENT '设备ID',
  `equipment_code` varchar(50) DEFAULT NULL COMMENT '设备编码',
  `equipment_name` varchar(200) DEFAULT NULL COMMENT '设备名称',
  `plan_id` bigint DEFAULT NULL COMMENT '巡检计划ID',
  `plan_no` varchar(50) DEFAULT NULL COMMENT '巡检计划单号',
  `inspection_type` varchar(20) DEFAULT NULL COMMENT '巡检类型',
  `inspection_level` varchar(20) DEFAULT NULL COMMENT '巡检级别',
  `inspector_id` bigint DEFAULT NULL COMMENT '巡检人员ID',
  `inspector_name` varchar(100) DEFAULT NULL COMMENT '巡检人员姓名',
  `inspection_time` datetime DEFAULT NULL COMMENT '巡检时间',
  `start_time` datetime DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime DEFAULT NULL COMMENT '结束时间',
  `duration_minutes` int DEFAULT NULL COMMENT '巡检时长(分钟)',
  `location_id` bigint DEFAULT NULL COMMENT '位置ID',
  `location_name` varchar(200) DEFAULT NULL COMMENT '位置名称',
  `workshop_id` bigint DEFAULT NULL COMMENT '车间ID',
  `workshop_name` varchar(100) DEFAULT NULL COMMENT '车间名称',
  `line_id` bigint DEFAULT NULL COMMENT '产线ID',
  `line_name` varchar(100) DEFAULT NULL COMMENT '产线名称',
  `total_items` int DEFAULT '0' COMMENT '总项数',
  `checked_items` int DEFAULT '0' COMMENT '已检项数',
  `normal_items` int DEFAULT '0' COMMENT '正常项数',
  `abnormal_items` int DEFAULT '0' COMMENT '异常项数',
  `na_items` int DEFAULT '0' COMMENT '不适用项数',
  `has_photo` bit(1) DEFAULT b'0' COMMENT '是否有照片',
  `has_signature` bit(1) DEFAULT b'0' COMMENT '是否有签名',
  `status` varchar(20) DEFAULT NULL COMMENT '状态',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户ID',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `version` int DEFAULT '0' COMMENT '乐观锁版本号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_inspection_no` (`inspection_no`),
  KEY `idx_equipment_id` (`equipment_id`),
  KEY `idx_plan_id` (`plan_id`),
  KEY `idx_inspector_id` (`inspector_id`),
  KEY `idx_inspection_time` (`inspection_time`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备巡检记录主表';
```

### 8.2 巡检记录明细表 (basic_equipment_inspection_record_detail)

```sql
CREATE TABLE `basic_equipment_inspection_record_detail` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `main_id` bigint NOT NULL COMMENT '主表ID',
  `inspection_no` varchar(50) DEFAULT NULL COMMENT '巡检单号',
  `item_id` bigint NOT NULL COMMENT '巡检项目ID',
  `item_code` varchar(50) DEFAULT NULL COMMENT '巡检项目编码',
  `item_name` varchar(200) DEFAULT NULL COMMENT '巡检项目名称',
  `item_type` varchar(20) DEFAULT NULL COMMENT '巡检项目类型',
  `standard_value` varchar(500) DEFAULT NULL COMMENT '标准值/判定标准',
  `result` varchar(20) DEFAULT NULL COMMENT '检查结果(OK/NG/NA)',
  `actual_value` varchar(500) DEFAULT NULL COMMENT '实际值',
  `is_normal` bit(1) DEFAULT NULL COMMENT '是否正常',
  `description` varchar(500) DEFAULT NULL COMMENT '描述/异常说明',
  `photo_urls` text COMMENT '照片URL列表(JSON)',
  `check_time` datetime DEFAULT NULL COMMENT '检查时间',
  `sequence` int DEFAULT '0' COMMENT '顺序',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户ID',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `version` int DEFAULT '0' COMMENT '乐观锁版本号',
  PRIMARY KEY (`id`),
  KEY `idx_main_id` (`main_id`),
  KEY `idx_item_id` (`item_id`),
  KEY `idx_result` (`result`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备巡检记录明细表';
```

### 8.3 保养记录主表 (basic_equipment_maintenance_main)

```sql
CREATE TABLE `basic_equipment_maintenance_main` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `maintenance_no` varchar(50) NOT NULL COMMENT '保养单号',
  `equipment_id` bigint NOT NULL COMMENT '设备ID',
  `equipment_code` varchar(50) DEFAULT NULL COMMENT '设备编码',
  `equipment_name` varchar(200) DEFAULT NULL COMMENT '设备名称',
  `plan_id` bigint DEFAULT NULL COMMENT '保养计划ID',
  `plan_no` varchar(50) DEFAULT NULL COMMENT '保养计划单号',
  `maintenance_type` varchar(20) DEFAULT NULL COMMENT '保养类型',
  `maintenance_level` varchar(20) DEFAULT NULL COMMENT '保养级别',
  `maintainer_id` bigint DEFAULT NULL COMMENT '保养人员ID',
  `maintainer_name` varchar(100) DEFAULT NULL COMMENT '保养人员姓名',
  `maintenance_time` datetime DEFAULT NULL COMMENT '保养时间',
  `start_time` datetime DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime DEFAULT NULL COMMENT '结束时间',
  `duration_hours` decimal(10,2) DEFAULT NULL COMMENT '保养时长(小时)',
  `location_id` bigint DEFAULT NULL COMMENT '位置ID',
  `location_name` varchar(200) DEFAULT NULL COMMENT '位置名称',
  `workshop_id` bigint DEFAULT NULL COMMENT '车间ID',
  `workshop_name` varchar(100) DEFAULT NULL COMMENT '车间名称',
  `line_id` bigint DEFAULT NULL COMMENT '产线ID',
  `line_name` varchar(100) DEFAULT NULL COMMENT '产线名称',
  `total_items` int DEFAULT '0' COMMENT '总项目数',
  `completed_items` int DEFAULT '0' COMMENT '已完成项目数',
  `total_parts` int DEFAULT '0' COMMENT '总配件数',
  `used_parts` int DEFAULT '0' COMMENT '使用配件数',
  `has_photo` bit(1) DEFAULT b'0' COMMENT '是否有照片',
  `has_signature` bit(1) DEFAULT b'0' COMMENT '是否有签名',
  `quality_check` bit(1) DEFAULT b'0' COMMENT '是否需要质检',
  `quality_result` varchar(20) DEFAULT NULL COMMENT '质检结果',
  `status` varchar(20) DEFAULT NULL COMMENT '状态',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户ID',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `version` int DEFAULT '0' COMMENT '乐观锁版本号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_maintenance_no` (`maintenance_no`),
  KEY `idx_equipment_id` (`equipment_id`),
  KEY `idx_plan_id` (`plan_id`),
  KEY `idx_maintainer_id` (`maintainer_id`),
  KEY `idx_maintenance_time` (`maintenance_time`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备保养记录主表';
```

### 8.4 保养记录明细表 (basic_equipment_maintenance_record_detail)

```sql
CREATE TABLE `basic_equipment_maintenance_record_detail` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `main_id` bigint NOT NULL COMMENT '主表ID',
  `maintenance_no` varchar(50) DEFAULT NULL COMMENT '保养单号',
  `item_id` bigint NOT NULL COMMENT '保养项目ID',
  `item_code` varchar(50) DEFAULT NULL COMMENT '保养项目编码',
  `item_name` varchar(200) DEFAULT NULL COMMENT '保养项目名称',
  `item_type` varchar(20) DEFAULT NULL COMMENT '保养项目类型',
  `work_content` varchar(500) DEFAULT NULL COMMENT '工作内容',
  `is_completed` bit(1) DEFAULT b'0' COMMENT '是否完成',
  `completion_time` datetime DEFAULT NULL COMMENT '完成时间',
  `completion_rate` decimal(5,2) DEFAULT NULL COMMENT '完成率',
  `description` varchar(500) DEFAULT NULL COMMENT '描述/说明',
  `photo_urls` text COMMENT '照片URL列表(JSON)',
  `spare_usage` text COMMENT '配件使用情况(JSON)',
  `sequence` int DEFAULT '0' COMMENT '顺序',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户ID',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `version` int DEFAULT '0' COMMENT '乐观锁版本号',
  PRIMARY KEY (`id`),
  KEY `idx_main_id` (`main_id`),
  KEY `idx_item_id` (`item_id`),
  KEY `idx_is_completed` (`is_completed`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备保养记录明细表';
```

### 8.5 设备停机记录表 (basic_equipment_shutdown)

```sql
CREATE TABLE `basic_equipment_shutdown` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `shutdown_no` varchar(50) NOT NULL COMMENT '停机单号',
  `equipment_id` bigint NOT NULL COMMENT '设备ID',
  `equipment_code` varchar(50) DEFAULT NULL COMMENT '设备编码',
  `equipment_name` varchar(200) DEFAULT NULL COMMENT '设备名称',
  `equipment_type` varchar(50) DEFAULT NULL COMMENT '设备类型',
  `shutdown_type` varchar(20) DEFAULT NULL COMMENT '停机类型(计划/非计划)',
  `shutdown_reason` varchar(50) DEFAULT NULL COMMENT '停机原因',
  `fault_type_id` bigint DEFAULT NULL COMMENT '故障类型ID',
  `fault_type_name` varchar(100) DEFAULT NULL COMMENT '故障类型名称',
  `fault_cause_id` bigint DEFAULT NULL COMMENT '故障原因ID',
  `fault_cause_name` varchar(200) DEFAULT NULL COMMENT '故障原因描述',
  `fault_description` varchar(500) DEFAULT NULL COMMENT '故障描述',
  `shutdown_time` datetime NOT NULL COMMENT '停机时间',
  `restart_time` datetime DEFAULT NULL COMMENT '重启时间',
  `duration_minutes` int DEFAULT NULL COMMENT '停机时长(分钟)',
  `affected_line_id` bigint DEFAULT NULL COMMENT '影响产线ID',
  `affected_line_name` varchar(100) DEFAULT NULL COMMENT '影响产线名称',
  `affected_quantity` decimal(18,3) DEFAULT NULL COMMENT '影响产量',
  `loss_amount` decimal(18,2) DEFAULT NULL COMMENT '损失金额',
  `reporter_id` bigint DEFAULT NULL COMMENT '上报人ID',
  `reporter_name` varchar(100) DEFAULT NULL COMMENT '上报人姓名',
  `reporter_time` datetime DEFAULT NULL COMMENT '上报时间',
  `handler_id` bigint DEFAULT NULL COMMENT '处理人ID',
  `handler_name` varchar(100) DEFAULT NULL COMMENT '处理人姓名',
  `handle_time` datetime DEFAULT NULL COMMENT '处理时间',
  `repair_ticket_no` varchar(50) DEFAULT NULL COMMENT '关联维修工单号',
  `status` varchar(20) DEFAULT NULL COMMENT '状态',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户ID',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `version` int DEFAULT '0' COMMENT '乐观锁版本号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_shutdown_no` (`shutdown_no`),
  KEY `idx_equipment_id` (`equipment_id`),
  KEY `idx_shutdown_time` (`shutdown_time`),
  KEY `idx_shutdown_type` (`shutdown_type`),
  KEY `idx_fault_type_id` (`fault_type_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备停机记录表';
```

### 8.6 设备转移记录表 (basic_equipment_transfer_record)

```sql
CREATE TABLE `basic_equipment_transfer_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `transfer_no` varchar(50) NOT NULL COMMENT '转移单号',
  `equipment_id` bigint NOT NULL COMMENT '设备ID',
  `equipment_code` varchar(50) DEFAULT NULL COMMENT '设备编码',
  `equipment_name` varchar(200) DEFAULT NULL COMMENT '设备名称',
  `from_workshop_id` bigint DEFAULT NULL COMMENT '源车间ID',
  `from_workshop_name` varchar(100) DEFAULT NULL COMMENT '源车间名称',
  `from_line_id` bigint DEFAULT NULL COMMENT '源产线ID',
  `from_line_name` varchar(100) DEFAULT NULL COMMENT '源产线名称',
  `from_location_id` bigint DEFAULT NULL COMMENT '源位置ID',
  `from_location_name` varchar(200) DEFAULT NULL COMMENT '源位置名称',
  `from_workstation_id` bigint DEFAULT NULL COMMENT '源工位ID',
  `from_workstation_name` varchar(100) DEFAULT NULL COMMENT '源工位名称',
  `to_workshop_id` bigint DEFAULT NULL COMMENT '目标车间ID',
  `to_workshop_name` varchar(100) DEFAULT NULL COMMENT '目标车间名称',
  `to_line_id` bigint DEFAULT NULL COMMENT '目标产线ID',
  `to_line_name` varchar(100) DEFAULT NULL COMMENT '目标产线名称',
  `to_location_id` bigint DEFAULT NULL COMMENT '目标位置ID',
  `to_location_name` varchar(200) DEFAULT NULL COMMENT '目标位置名称',
  `to_workstation_id` bigint DEFAULT NULL COMMENT '目标工位ID',
  `to_workstation_name` varchar(100) DEFAULT NULL COMMENT '目标工位名称',
  `transfer_type` varchar(20) DEFAULT NULL COMMENT '转移类型(调拨/借用/归还)',
  `transfer_reason` varchar(500) DEFAULT NULL COMMENT '转移原因',
  `plan_date` date DEFAULT NULL COMMENT '计划日期',
  `actual_date` date DEFAULT NULL COMMENT '实际日期',
  `applicant_id` bigint DEFAULT NULL COMMENT '申请人ID',
  `applicant_name` varchar(100) DEFAULT NULL COMMENT '申请人姓名',
  `apply_time` datetime DEFAULT NULL COMMENT '申请时间',
  `approver_id` bigint DEFAULT NULL COMMENT '审批人ID',
  `approver_name` varchar(100) DEFAULT NULL COMMENT '审批人姓名',
  `approve_time` datetime DEFAULT NULL COMMENT '审批时间',
  `approve_result` varchar(20) DEFAULT NULL COMMENT '审批结果',
  `approve_remark` varchar(500) DEFAULT NULL COMMENT '审批备注',
  `executor_id` bigint DEFAULT NULL COMMENT '执行人ID',
  `executor_name` varchar(100) DEFAULT NULL COMMENT '执行人姓名',
  `execute_time` datetime DEFAULT NULL COMMENT '执行时间',
  `status` varchar(20) DEFAULT NULL COMMENT '状态',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户ID',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `version` int DEFAULT '0' COMMENT '乐观锁版本号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_transfer_no` (`transfer_no`),
  KEY `idx_equipment_id` (`equipment_id`),
  KEY `idx_from_location_id` (`from_location_id`),
  KEY `idx_to_location_id` (`to_location_id`),
  KEY `idx_transfer_date` (`actual_date`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备转移记录表';
```

### 8.7 备件表 (basic_equipment_tool_spare_part)

```sql
CREATE TABLE `basic_equipment_tool_spare_part` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `spare_code` varchar(50) NOT NULL COMMENT '备件编码',
  `spare_name` varchar(200) NOT NULL COMMENT '备件名称',
  `spare_type` varchar(50) DEFAULT NULL COMMENT '备件类型',
  `specification` varchar(200) DEFAULT NULL COMMENT '规格型号',
  `unit` varchar(20) DEFAULT NULL COMMENT '单位',
  `material` varchar(100) DEFAULT NULL COMMENT '材质',
  `weight` decimal(10,2) DEFAULT NULL COMMENT '重量',
  `supplier_id` bigint DEFAULT NULL COMMENT '供应商ID',
  `supplier_name` varchar(200) DEFAULT NULL COMMENT '供应商名称',
  `manufacturer_id` bigint DEFAULT NULL COMMENT '生产厂商ID',
  `manufacturer_name` varchar(200) DEFAULT NULL COMMENT '生产厂商名称',
  `purchase_price` decimal(18,2) DEFAULT NULL COMMENT '采购价格',
  `standard_price` decimal(18,2) DEFAULT NULL COMMENT '标准价格',
  `min_stock` int DEFAULT '0' COMMENT '最小库存',
  `max_stock` int DEFAULT '0' COMMENT '最大库存',
  `current_stock` int DEFAULT '0' COMMENT '当前库存',
  `safe_stock` int DEFAULT '0' COMMENT '安全库存',
  `stock_status` varchar(20) DEFAULT NULL COMMENT '库存状态',
  `warehouse_id` bigint DEFAULT NULL COMMENT '仓库ID',
  `warehouse_name` varchar(100) DEFAULT NULL COMMENT '仓库名称',
  `location_id` bigint DEFAULT NULL COMMENT '库位ID',
  `location_name` varchar(100) DEFAULT NULL COMMENT '库位名称',
  `image_url` varchar(500) DEFAULT NULL COMMENT '图片URL',
  `description` varchar(500) DEFAULT NULL COMMENT '描述',
  `is_key_spare` bit(1) DEFAULT b'0' COMMENT '是否关键备件',
  `lead_time` int DEFAULT NULL COMMENT '采购周期(天)',
  `life_cycle` int DEFAULT NULL COMMENT '使用寿命',
  `replacement_cycle` int DEFAULT NULL COMMENT '更换周期(天)',
  `status` varchar(20) DEFAULT NULL COMMENT '状态',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户ID',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `version` int DEFAULT '0' COMMENT '乐观锁版本号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_spare_code` (`spare_code`),
  KEY `idx_spare_name` (`spare_name`),
  KEY `idx_spare_type` (`spare_type`),
  KEY `idx_supplier_id` (`supplier_id`),
  KEY `idx_manufacturer_id` (`manufacturer_id`),
  KEY `idx_stock_status` (`stock_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备备件表';
```

### 8.8 设备厂商表 (basic_equipment_manufacturer)

```sql
CREATE TABLE `basic_equipment_manufacturer` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `manufacturer_code` varchar(50) NOT NULL COMMENT '厂商编码',
  `manufacturer_name` varchar(200) NOT NULL COMMENT '厂商名称',
  `manufacturer_type` varchar(50) DEFAULT NULL COMMENT '厂商类型',
  `country` varchar(50) DEFAULT NULL COMMENT '国家/地区',
  `province` varchar(50) DEFAULT NULL COMMENT '省份',
  `city` varchar(50) DEFAULT NULL COMMENT '城市',
  `address` varchar(500) DEFAULT NULL COMMENT '详细地址',
  `contact_person` varchar(100) DEFAULT NULL COMMENT '联系人',
  `contact_phone` varchar(50) DEFAULT NULL COMMENT '联系电话',
  `contact_email` varchar(100) DEFAULT NULL COMMENT '联系邮箱',
  `website` varchar(200) DEFAULT NULL COMMENT '网址',
  `business_scope` varchar(500) DEFAULT NULL COMMENT '经营范围',
  `main_products` varchar(500) DEFAULT NULL COMMENT '主要产品',
  `qualification_cert` varchar(200) DEFAULT NULL COMMENT '资质证书',
  `cert_expiry_date` date DEFAULT NULL COMMENT '证书到期日期',
  `credit_rating` varchar(20) DEFAULT NULL COMMENT '信用评级',
  `cooperate_level` varchar(20) DEFAULT NULL COMMENT '合作级别',
  `cooperate_start_date` date DEFAULT NULL COMMENT '合作开始日期',
  `payment_terms` varchar(100) DEFAULT NULL COMMENT '付款条款',
  `bank_name` varchar(100) DEFAULT NULL COMMENT '开户银行',
  `bank_account` varchar(100) DEFAULT NULL COMMENT '银行账号',
  `tax_no` varchar(50) DEFAULT NULL COMMENT '税号',
  `status` varchar(20) DEFAULT NULL COMMENT '状态',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户ID',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `version` int DEFAULT '0' COMMENT '乐观锁版本号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_manufacturer_code` (`manufacturer_code`),
  KEY `idx_manufacturer_name` (`manufacturer_name`),
  KEY `idx_manufacturer_type` (`manufacturer_type`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备厂商表';
```

### 8.9 设备供应商表 (basic_equipment_supplier)

```sql
CREATE TABLE `basic_equipment_supplier` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `supplier_code` varchar(50) NOT NULL COMMENT '供应商编码',
  `supplier_name` varchar(200) NOT NULL COMMENT '供应商名称',
  `supplier_type` varchar(50) DEFAULT NULL COMMENT '供应商类型',
  `category` varchar(50) DEFAULT NULL COMMENT '供应商类别',
  `country` varchar(50) DEFAULT NULL COMMENT '国家/地区',
  `province` varchar(50) DEFAULT NULL COMMENT '省份',
  `city` varchar(50) DEFAULT NULL COMMENT '城市',
  `address` varchar(500) DEFAULT NULL COMMENT '详细地址',
  `contact_person` varchar(100) DEFAULT NULL COMMENT '联系人',
  `contact_phone` varchar(50) DEFAULT NULL COMMENT '联系电话',
  `contact_fax` varchar(50) DEFAULT NULL COMMENT '传真',
  `contact_email` varchar(100) DEFAULT NULL COMMENT '联系邮箱',
  `website` varchar(200) DEFAULT NULL COMMENT '网址',
  `business_scope` varchar(500) DEFAULT NULL COMMENT '经营范围',
  `main_products` varchar(500) DEFAULT NULL COMMENT '主要产品',
  `registered_capital` decimal(18,2) DEFAULT NULL COMMENT '注册资本',
  `annual_revenue` decimal(18,2) DEFAULT NULL COMMENT '年营业额',
  `employee_count` int DEFAULT NULL COMMENT '员工人数',
  `qualification_cert` varchar(200) DEFAULT NULL COMMENT '资质证书',
  `cert_expiry_date` date DEFAULT NULL COMMENT '证书到期日期',
  `credit_rating` varchar(20) DEFAULT NULL COMMENT '信用评级',
  `cooperate_level` varchar(20) DEFAULT NULL COMMENT '合作级别',
  `cooperate_start_date` date DEFAULT NULL COMMENT '合作开始日期',
  `payment_terms` varchar(100) DEFAULT NULL COMMENT '付款条款',
  `delivery_cycle` int DEFAULT NULL COMMENT '交货周期(天)',
  `min_order_qty` int DEFAULT NULL COMMENT '最小订购量',
  `bank_name` varchar(100) DEFAULT NULL COMMENT '开户银行',
  `bank_account` varchar(100) DEFAULT NULL COMMENT '银行账号',
  `tax_no` varchar(50) DEFAULT NULL COMMENT '税号',
  `status` varchar(20) DEFAULT NULL COMMENT '状态',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户ID',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `version` int DEFAULT '0' COMMENT '乐观锁版本号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_supplier_code` (`supplier_code`),
  KEY `idx_supplier_name` (`supplier_name`),
  KEY `idx_supplier_type` (`supplier_type`),
  KEY `idx_category` (`category`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备供应商表';
```

### 8.10 设备主要部件表 (basic_equipment_main_part)

```sql
CREATE TABLE `basic_equipment_main_part` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `equipment_id` bigint NOT NULL COMMENT '设备ID',
  `equipment_code` varchar(50) DEFAULT NULL COMMENT '设备编码',
  `equipment_name` varchar(200) DEFAULT NULL COMMENT '设备名称',
  `part_code` varchar(50) NOT NULL COMMENT '部件编码',
  `part_name` varchar(200) NOT NULL COMMENT '部件名称',
  `part_type` varchar(50) DEFAULT NULL COMMENT '部件类型',
  `specification` varchar(200) DEFAULT NULL COMMENT '规格型号',
  `serial_no` varchar(100) DEFAULT NULL COMMENT '序列号',
  `material` varchar(100) DEFAULT NULL COMMENT '材质',
  `weight` decimal(10,2) DEFAULT NULL COMMENT '重量',
  `manufacturer_id` bigint DEFAULT NULL COMMENT '生产厂商ID',
  `manufacturer_name` varchar(200) DEFAULT NULL COMMENT '生产厂商名称',
  `purchase_date` date DEFAULT NULL COMMENT '购置日期',
  `unit_price` decimal(18,2) DEFAULT NULL COMMENT '单价',
  `lifecycle_count` int DEFAULT NULL COMMENT '设计寿命(次数)',
  `lifecycle_hours` decimal(10,2) DEFAULT NULL COMMENT '设计寿命(小时)',
  `used_count` int DEFAULT '0' COMMENT '已使用次数',
  `used_hours` decimal(10,2) DEFAULT '0' COMMENT '已使用小时',
  `replacement_cycle` int DEFAULT NULL COMMENT '更换周期(天)',
  `last_replacement_time` datetime DEFAULT NULL COMMENT '上次更换时间',
  `next_replacement_time` datetime DEFAULT NULL COMMENT '下次更换时间',
  `current_status` varchar(20) DEFAULT NULL COMMENT '当前状态',
  `wear_level` varchar(20) DEFAULT NULL COMMENT '磨损程度',
  `position` varchar(100) DEFAULT NULL COMMENT '安装位置',
  `image_url` varchar(500) DEFAULT NULL COMMENT '图片URL',
  `description` varchar(500) DEFAULT NULL COMMENT '描述',
  `is_critical` bit(1) DEFAULT b'0' COMMENT '是否关键部件',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户ID',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `version` int DEFAULT '0' COMMENT '乐观锁版本号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_part_code` (`part_code`),
  KEY `idx_equipment_id` (`equipment_id`),
  KEY `idx_part_type` (`part_type`),
  KEY `idx_current_status` (`current_status`),
  KEY `idx_next_replacement_time` (`next_replacement_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备主要部件表';
```

### 8.11 故障类型表 (basic_fault_type)

```sql
CREATE TABLE `basic_fault_type` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `fault_type_code` varchar(50) NOT NULL COMMENT '故障类型编码',
  `fault_type_name` varchar(100) NOT NULL COMMENT '故障类型名称',
  `fault_category` varchar(50) DEFAULT NULL COMMENT '故障类别',
  `fault_severity` varchar(20) DEFAULT NULL COMMENT '故障严重程度',
  `fault_position` varchar(100) DEFAULT NULL COMMENT '故障部位',
  `fault_phenomenon` varchar(200) DEFAULT NULL COMMENT '故障现象',
  `solution_hint` varchar(500) DEFAULT NULL COMMENT '解决方案提示',
  `repair_level` varchar(20) DEFAULT NULL COMMENT '维修级别',
  `estimated_time` int DEFAULT NULL COMMENT '预计维修时间(分钟)',
  `spare_hints` varchar(500) DEFAULT NULL COMMENT '所需备件提示',
  `skill_requirements` varchar(500) DEFAULT NULL COMMENT '技能要求',
  `is_enabled` bit(1) DEFAULT b'1' COMMENT '是否启用',
  `sort` int DEFAULT '0' COMMENT '排序',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户ID',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `version` int DEFAULT '0' COMMENT '乐观锁版本号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_fault_type_code` (`fault_type_code`),
  KEY `idx_fault_type_name` (`fault_type_name`),
  KEY `idx_fault_category` (`fault_category`),
  KEY `idx_is_enabled` (`is_enabled`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='故障类型表';
```

### 8.12 故障原因表 (basic_fault_cause)

```sql
CREATE TABLE `basic_fault_cause` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `fault_type_id` bigint NOT NULL COMMENT '故障类型ID',
  `fault_type_name` varchar(100) DEFAULT NULL COMMENT '故障类型名称',
  `reason_code` varchar(50) NOT NULL COMMENT '原因编码',
  `reason_name` varchar(200) NOT NULL COMMENT '原因名称',
  `reason_level` varchar(20) DEFAULT NULL COMMENT '原因层级(根因/间因/直接)',
  `reason_category` varchar(50) DEFAULT NULL COMMENT '原因类别',
  `description` varchar(500) DEFAULT NULL COMMENT '原因描述',
  `solution` varchar(500) DEFAULT NULL COMMENT '解决方案',
  `prevention` varchar(500) DEFAULT NULL COMMENT '预防措施',
  `occurrence_prob` decimal(5,2) DEFAULT NULL COMMENT '发生概率(%)',
  `is_enabled` bit(1) DEFAULT b'1' COMMENT '是否启用',
  `sort` int DEFAULT '0' COMMENT '排序',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint DEFAULT NULL COMMENT '租户ID',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) DEFAULT b'0' COMMENT '是否删除',
  `version` int DEFAULT '0' COMMENT '乐观锁版本号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_reason_code` (`reason_code`),
  KEY `idx_fault_type_id` (`fault_type_id`),
  KEY `idx_reason_name` (`reason_name`),
  KEY `idx_reason_level` (`reason_level`),
  KEY `idx_is_enabled` (`is_enabled`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='故障原因表';
```

---

## 9. 关联文档

- [MOM3.0_主设计文档](./MOM3.0_主设计文档.md) - 系统总览
- [MOM3.0_UI设计规范](./MOM3.0_UI设计规范.md) - UI规范详情
- [MOM3.0_设备管理模块设计文档](./MOM3.0_设备管理模块设计文档.md) - 后端设计详情
- [sfms3.0/EAM模块设计文档](./sfms3.0/EAM模块设计文档.md) - SFMS3.0 EAM参考
