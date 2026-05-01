# SFMS3.0 vs MOM3.0 设计文档对比清单

> 生成时间: 2026-04-17
> 说明: 本文档对比SFMS3.0完整设计与当前MOM3.0设计文档的差异，列出缺失内容供参考补充。

---

## 1. MES模块差异

### 1.1 缺失表结构

| 表名 | 说明 | 优先级 |
|------|------|--------|
| plan_mes_order_day_bom | 日计划BOM实例表 | P1 |
| plan_mes_order_day_route | 日计划工艺路线实例表 | P1 |
| plan_mes_order_day_routesub | 日计划工序明细表 | P1 |
| plan_mes_order_day_equipment | 日计划设备配置表 | P1 |
| plan_mes_order_day_worker | 日计划人员配置表 | P1 |
| plan_mes_order_day_workstation | 日计划工位配置表 | P1 |
| plan_mes_work_scheduling_detail | 工序排程明细表 | P1 |

### 1.2 缺失API接口

| API路径 | 说明 | 优先级 |
|--------|------|--------|
| /mes/orderday/bom/* | BOM管理API | P1 |
| /mes/orderday/route/* | 工艺路线API | P1 |
| /mes/orderday/check/stuffing | 齐套检查API | P1 |
| /mes/workstation/* | 工位管理API | P1 |
| /mes/worker/* | 人员班组API | P1 |
| /mes/equipment/bind/* | 设备绑定API | P1 |

### 1.3 缺失业务功能

| 功能 | 说明 | 优先级 |
|------|------|--------|
| 工艺路线管理 | 工艺路线定义、工序管理 | P1 |
| BOM管理 | 物料清单管理、版本控制 | P1 |
| 工位配置 | 工位定义、产线关联 | P1 |
| 人员班组 | 班组管理、人员排班 | P1 |
| 设备绑定 | 设备与工位/订单绑定 | P1 |
| 齐套检查 | 物料齐套验证 | P1 |
| 工序级报工 | 按工序报工而非工单报工 | P2 |

---

## 2. WMS模块差异

### 2.1 缺失表结构

| 表名 | 说明 | 优先级 |
|------|------|--------|
| wms_areabasic | 库区基础表 | P1 |
| wms_itembasic | 货品基础表 | P1 |
| wms_container_main | 容器主表 | P1 |
| wms_container_detail | 容器明细表 | P1 |
| wms_pick_job | 拣货作业表 | P1 |
| wms_pick_record | 拣货记录表 | P1 |
| wms_putaway_job | 上架作业表 | P1 |
| wms_putaway_record | 上架记录表 | P1 |
| wms_count_plan | 盘点计划表 | P1 |
| wms_count_job | 盘点任务表 | P1 |
| wms_count_record | 盘点记录表 | P1 |
| wms_inventorychange_record | 库存异动记录表 | P1 |
| wms_agv_task | AGV任务表 | P1 |
| wms_agv_job | AGV作业表 | P1 |
| wms_label_template | 标签模板表 | P1 |
| wms_label_print_record | 标签打印记录表 | P1 |
| wms_businesstype | 业务类型表 | P1 |
| wms_strategy | 策略配置表 | P1 |

### 2.2 缺失API接口

| API路径 | 说明 | 优先级 |
|--------|------|--------|
| /wms/area/* | 库区管理API | P1 |
| /wms/item/* | 货品管理API | P1 |
| /wms/container/* | 容器管理API | P1 |
| /wms/pick/* | 拣货作业API | P1 |
| /wms/putaway/* | 上架作业API | P1 |
| /wms/count/* | 盘点管理API | P1 |
| /wms/agv/* | AGV调度API | P1 |
| /wms/label/* | 标签打印API | P1 |
| /wms/strategy/* | 策略配置API | P1 |

### 2.3 缺失业务功能

| 功能 | 说明 | 优先级 |
|------|------|--------|
| 库区管理 | 仓库-库区-库位三级结构 | P1 |
| 货品管理 | 物料与货品对应关系 | P1 |
| 容器管理 | 容器绑定/解绑/维修 | P1 |
| AGV调度 | AGV任务下发与跟踪 | P1 |
| 标签打印 | 标签模板设计、条码规则 | P1 |
| 拣货下架 | 销售发货/领料拣货 | P1 |
| 上架策略 | 入库上架规则配置 | P1 |
| 盘点管理 | 盘点计划/任务/执行 | P1 |

---

## 3. SCP模块差异

### 3.1 缺失表结构

| 表名 | 说明 | 优先级 |
|------|------|--------|
| scp_purchase_plan | 采购计划表(MPS/MRP) | P1 |
| scp_supplier | 供应商主表 | P1 |
| scp_supplier_contact | 供应商联系人表 | P1 |
| scp_supplier_bank | 供应商银行账户表 | P1 |
| scp_customer | 客户主表 | P1 |
| scp_customer_contact | 客户联系人表 | P1 |
| scp_customer_bank | 客户银行账户表 | P1 |
| scp_customer_credit | 客户信用表 | P1 |

### 3.2 缺失API接口

| API路径 | 说明 | 优先级 |
|--------|------|--------|
| /scp/plan/purchase/* | 采购计划API | P1 |
| /scp/supplier/* | 供应商管理API | P1 |
| /scp/customer/* | 客户管理API | P1 |
| /scp/supplier/confirm | 供应商确认API | P2 |
| /scp/plan/confirm | 计划员确认API | P2 |

### 3.3 缺失业务功能

| 功能 | 说明 | 优先级 |
|------|------|--------|
| MPS/MRP采购计划 | 生产计划驱动的采购 | P1 |
| 供应商档案管理 | 完整供应商信息管理 | P1 |
| 供应商评价 | 供应商绩效考核 | P1 |
| 客户档案管理 | 客户信息管理 | P1 |
| 客户信用管理 | 信用额度控制 | P1 |
| 采购统计 | 采购分析报表 | P2 |
| 销售统计 | 销售分析报表 | P2 |

---

## 4. QMS模块差异

### 4.1 缺失表结构

| 表名 | 说明 | 优先级 |
|------|------|--------|
| qms_aql | AQL标准配置表 | P1 |
| qms_aql_detail | AQL明细表 | P1 |
| qms_material_aql | 物料AQL配置表 | P1 |
| qms_inspection_characteristics | 检验特性表 | P1 |
| qms_char_quantitative | 定量特性表 | P1 |
| qms_char_qualitative | 定性特性表 | P1 |
| qms_inspection_method | 检验方法表 | P1 |
| qms_inspection_scheme | 检验方案表 | P1 |
| qms_sampling_plan | 抽样计划表 | P1 |
| qms_sampling_rule | 抽样规则表 | P1 |
| qms_sampling_record | 抽样记录表 | P1 |
| qms_dynamic_rule | 动态检验规则表 | P1 |
| qms_sample_code | 样本编码表 | P1 |

### 4.2 缺失API接口

| API路径 | 说明 | 优先级 |
|--------|------|--------|
| /qms/aql/* | AQL标准管理API | P1 |
| /qms/characteristics/* | 检验特性管理API | P1 |
| /qms/method/* | 检验方法管理API | P1 |
| /qms/scheme/* | 检验方案管理API | P1 |
| /qms/sampling/* | 抽样方案管理API | P1 |
| /qms/dynamic/* | 动态规则API | P1 |

### 4.3 缺失业务功能

| 功能 | 说明 | 优先级 |
|------|------|--------|
| AQL标准管理 | 接收质量限标准配置 | P1 |
| 检验特性管理 | 定性/定量特性定义 | P1 |
| 检验方案配置 | 检验方案模板 | P1 |
| 抽样方案 | 抽样规则定义 | P1 |
| 动态检验规则 | 按条件自动调整检验级别 | P1 |
| 样本编码管理 | 样本追踪编码 | P2 |

---

## 5. EAM模块差异

### 5.1 缺失表结构

| 表名 | 说明 | 优先级 |
|------|------|--------|
| basic_equipment_inspection_record_main | 巡检记录主表 | P1 |
| basic_equipment_inspection_record_detail | 巡检记录明细表 | P1 |
| basic_equipment_maintenance_main | 保养记录主表 | P1 |
| basic_equipment_maintenance_detail | 保养记录明细表 | P1 |
| basic_equipment_shutdown | 设备停机记录表 | P1 |
| basic_equipment_transfer_record | 设备转移记录表 | P1 |
| basic_equipment_tool_spare_part | 备件表 | P1 |
| basic_equipment_manufacturer | 设备厂商表 | P1 |
| basic_equipment_supplier | 设备供应商表 | P1 |
| basic_equipment_main_part | 设备主要部件表 | P1 |
| basic_fault_type | 故障类型表 | P1 |
| basic_fault_cause | 故障原因表 | P1 |

### 5.2 缺失API接口

| API路径 | 说明 | 优先级 |
|--------|------|--------|
| /eam/inspection/* | 设备巡检API | P1 |
| /eam/maintenance/* | 设备保养API | P1 |
| /eam/shutdown/* | 设备停机API | P1 |
| /eam/transfer/* | 设备转移API | P1 |
| /eam/spare/* | 备件管理API | P1 |
| /eam/manufacturer/* | 厂商管理API | P1 |

### 5.3 缺失业务功能

| 功能 | 说明 | 优先级 |
|------|------|--------|
| 设备巡检管理 | 巡检计划/执行/记录 | P1 |
| 设备保养管理 | 保养计划/执行/级别 | P1 |
| 设备停机记录 | 非计划停机追踪 | P1 |
| 设备转移记录 | 设备调拨追踪 | P1 |
| 备件管理 | 备件库存/使用记录 | P1 |
| 故障分析 | 故障类型/原因统计 | P2 |

---

## 6. BPM模块差异

### 6.1 缺失表结构

| 表名 | 说明 | 优先级 |
|------|------|--------|
| bpm_process_definition_ext | 流程定义扩展表 | P1 |
| bpm_task_ext | 任务扩展表 | P1 |
| bpm_form_config | 表单配置表 | P1 |
| bpm_task_assign_rule | 任务分配规则表 | P1 |

### 6.2 缺失API接口

| API路径 | 说明 | 优先级 |
|--------|------|--------|
| /bpm/process/* | 流程定义API | P1 |
| /bpm/task/* | 任务管理API | P1 |
| /bpm/form/* | 表单配置API | P1 |
| /bpm/oa/* | OA集成API | P2 |

### 6.3 缺失业务功能

| 功能 | 说明 | 优先级 |
|------|------|--------|
| BPMN流程设计 | 流程定义与部署 | P1 |
| 流程实例管理 | 发起/审批/取消 | P1 |
| 任务分配规则 | 自动分配策略 | P1 |
| OA集成 | 请假等审批流 | P2 |

---

## 7. 优先级汇总

### P1 - 必须补充

| 模块 | 缺失数量 | 主要缺失 |
|------|---------|---------|
| MES | 7表+6API | 工艺路线、BOM、工位、齐套检查 |
| WMS | 18表+9API | 库区、容器、AGV、标签、盘点 |
| SCP | 8表+5API | 采购计划、供应商、客户 |
| QMS | 13表+6API | AQL、检验特性、抽样方案 |
| EAM | 12表+6API | 巡检、保养、停机、备件 |
| BPM | 4表+4API | 流程定义、表单配置 |

### P2 - 建议补充

| 模块 | 缺失内容 |
|------|---------|
| MES | 工序级报工、临时工艺 |
| WMS | 验货流程、上架策略 |
| SCP | 供应商/计划员审批、采购统计 |
| QMS | 样本编码管理 |
| EAM | 故障分析统计 |
| BPM | OA集成 |

---

## 8. 补充建议

1. **优先补充P1缺失**：这些是核心业务功能，当前设计文档中完全缺失

2. **MES模块最紧急**：缺少工艺路线和BOM实例表，这是生产执行的核心

3. **WMS模块工作量大**：表结构差异达18张，需要分批补充

4. **保持文档同步**：补充后应及时更新设计文档状态表
