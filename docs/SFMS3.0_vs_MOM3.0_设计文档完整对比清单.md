# MOM3.0 vs SFMS3.0 设计文档完整对比清单

> 生成时间: 2026-04-17
> 说明: 对比MOM3.0和SFMS3.0各模块设计文档，列出缺失内容供补充参考

---

## 一、汇总统计

| 模块 | 缺失DDL表 | 缺失API | 缺失业务功能 | 最紧急优先级 |
|------|-----------|---------|-------------|-------------|
| MES | 40+张 | 200+个 | 30+项 | P1 |
| WMS | 18张 | 100+个 | 10+项 | P1 |
| SCP | 22+张 | 126+个 | 20+项 | P1 |
| QMS | 13张 | 50+个 | 8项 | P1 |
| EAM | 74张 | 200+个 | 29项 | P1 |
| BPM | 1张 | 6个 | 8项 | P2 |

---

## 二、MES模块差异

### 2.1 缺失DDL表结构

| 类别 | 表名 | SFMS3.0说明 | MOM3.0状态 | 优先级 |
|------|------|------------|-----------|--------|
| 生产日计划 | `plan_mes_order_day` | 生产日计划主表 | **未实现** | P1 |
| 生产日计划 | `plan_mes_order_day_bom` | 日计划BOM表 | **未实现** | P1 |
| 生产日计划 | `plan_mes_order_day_route` | 日计划工艺路线表 | **未实现** | P1 |
| 生产日计划 | `plan_mes_order_day_routesub` | 日计划工序明细表 | **未实现** | P1 |
| 生产日计划 | `plan_mes_order_day_equipment` | 日计划设备配置表 | **未实现** | P2 |
| 生产日计划 | `plan_mes_order_day_worker` | 日计划人员配置表 | **未实现** | P2 |
| 生产日计划 | `plan_mes_order_day_workstation` | 日计划工位配置表 | **未实现** | P2 |
| 工单排程 | `plan_mes_work_scheduling` | 工单排程主表 | **未实现** | P1 |
| 工单排程 | `plan_mes_work_scheduling_detail` | 工序任务明细表 | **未实现** | P1 |
| 质检管理 | `qms_qualityclass` | 质检类别表 | **未实现** | P2 |
| 质检管理 | `qms_qualitygroup` | 质检分组表 | **未实现** | P2 |
| 质检管理 | `qms_item` | 质检项目定义表 | **未实现** | P2 |
| 质检管理 | `qms_qualityform` | 质检表单表 | **未实现** | P2 |
| 齐套检查 | `mes_complete_inspect` | 齐套检查配置表 | **未实现** | P1 |
| 报工管理 | `mes_job_report_log` | 工序报工日志表 | **未实现** | P1 |
| 返工返修 | `mes_rework_batch` | 返工登记批量表 | **未实现** | P2 |
| 返工返修 | `mes_rework_single` | 返工登记单件表 | **未实现** | P2 |
| 返工返修 | `mes_product_offline` | 产品离线登记表 | **未实现** | P2 |
| 模具管理 | `mes_pattern` | 模具基本信息表 | **未实现** | P2 |
| 叫料管理 | `mes_item_request_main` | 叫料申请单主表 | **未实现** | P2 |
| 能力矩阵 | `mes_ability_info` | 能力矩阵信息表 | **未实现** | P2 |
| 能力矩阵 | `mes_hr_person_ability` | 人员能力关联表 | **未实现** | P2 |
| 订单月计划 | `plan_mes_order_month_main` | 订单月计划主表 | **未实现** | P2 |
| 订单月计划 | `plan_mes_order_month_sub` | 订单月计划子表 | **未实现** | P2 |

### 2.2 缺失API接口

| 类别 | 路径 | SFMS3.0说明 | 优先级 |
|------|------|------------|--------|
| 日计划核心 | `POST /mes/orderday/create` | 创建日计划 | P1 |
| 日计划核心 | `POST /mes/orderday/publishPlan` | 发布日计划 | P1 |
| 日计划核心 | `POST /mes/orderday/stopPlan/{id}` | 终止计划 | P1 |
| 工单排程 | `POST /mes/workScheduling/create` | 创建生产任务排产 | P1 |
| 工单排程 | `POST /mes/workScheduling/update-status` | 更新工单状态 | P1 |
| 工单排程 | `POST /mes/workScheduling/completeHandle` | 完工处理 | P1 |
| 工单排程 | `POST /mes/workScheduling/reportForAll` | 批量报工 | P1 |
| 工序报工 | `POST /mes/work-scheduling-detail/reportWorkByProcess` | 工序报工 | P1 |
| 工序报工 | `POST /mes/work-scheduling-detail/processFinished` | 工序完工 | P1 |
| 齐套检查 | `POST /mes/complete-inspect/get-orderDay-bom` | 获取日计划Bom | P1 |
| 齐套检查 | `POST /mes/complete-inspect/get-orderDay-equipment` | 获取设备信息 | P1 |
| 公共查询 | `GET /mes/common/getProcessRouteByProductCode` | 获取工艺路线 | P1 |
| 公共查询 | `GET /mes/common/geBomTreeByProductCode` | 获取BOM树 | P1 |
| PDA端 | `GET /mes/workScheduling/PDA-page` | PDA工单列表 | P1 |
| 质检管理 | `POST /mes/qualityclass/*` | 质检类别CRUD | P2 |
| 质检管理 | `POST /mes/item/*` | 质检项目CRUD | P2 |
| 班组管理 | `POST /mes/teamSetting/*` | 班组完整CRUD | P2 |
| 节假日 | `POST /mes/holiday/*` | 节假日CRUD | P2 |
| 生产日历 | `POST /mes/schedulingcalendar/*` | 生产日历CRUD | P2 |
| 返工返修 | `POST /mes/product-offline/*` | 产品离线CRUD | P2 |
| 返工返修 | `POST /mes/rework-batch/*` | 批量返工CRUD | P2 |
| 模具管理 | `POST /mes/pattern/*` | 模具CRUD | P2 |
| 能力矩阵 | `POST /mes/ability-info/*` | 能力信息CRUD | P2 |
| 叫料申请 | `POST /mes/item-request-main/*` | 叫料申请CRUD | P2 |
| 月计划 | `POST /plan/mes-order-month-main/*` | 月计划CRUD | P2 |

### 2.3 缺失业务功能

| 功能 | 说明 | 优先级 |
|------|------|--------|
| 日计划创建 | 自动创建工艺/BOM/设备/工位/人员实例 | P1 |
| 计划发布 | 齐套检查+生成工单 | P1 |
| 工序级报工 | 按工序报工、批量报工 | P1 |
| 工序开工/暂停/恢复 | 工序状态机 | P1 |
| 完工处理 | 工单完工自动处理 | P1 |
| PDA端支持 | 移动端工单列表和报工 | P1 |
| 返工返修流程 | 离线/返线/批量/单件处理 | P2 |
| 质检管理 | 类别/分组/项目/表单 | P2 |
| 班组+日历 | 人员排班基础 | P2 |
| 能力矩阵 | 人员技能管理 | P2 |
| 模具管理 | 模具台账 | P2 |
| 月计划拆解 | breakdown接口 | P2 |

**MES完成度评估: 约5-8%**

---

## 三、WMS模块差异

### 3.1 缺失DDL表结构

| 表名 | 说明 | 优先级 |
|------|------|--------|
| `wms_areabasic` | 库区基础表（仓库-库区-库位三级结构） | P1 |
| `wms_itembasic` | 货品基础表 | P1 |
| `wms_container_main` | 容器主表 | P1 |
| `wms_container_detail` | 容器明细表 | P1 |
| `wms_pick_job` | 拣货作业表 | P1 |
| `wms_pick_record` | 拣货记录表 | P1 |
| `wms_putaway_job` | 上架作业表 | P1 |
| `wms_putaway_record` | 上架记录表 | P1 |
| `wms_count_plan` | 盘点计划表 | P1 |
| `wms_count_job` | 盘点任务表 | P1 |
| `wms_count_record` | 盘点记录表 | P1 |
| `wms_inventorychange_record` | 库存异动记录表 | P2 |
| `wms_agv_task` | AGV任务表 | P2 |
| `wms_agv_job` | AGV作业表 | P2 |
| `wms_label_template` | 标签模板表 | P2 |
| `wms_label_print_record` | 标签打印记录表 | P2 |
| `wms_businesstype` | 业务类型表 | P2 |
| `wms_strategy` | 策略配置表 | P3 |

### 3.2 缺失API接口

| 路径 | 说明 | 优先级 |
|------|------|--------|
| `/wms/container/create` | 创建容器档案 | P1 |
| `/wms/container/bind` | 容器绑定货品 | P1 |
| `/wms/container/unbind` | 容器解绑 | P1 |
| `/wms/pick/create` | 创建拣货作业 | P1 |
| `/wms/pick/assign` | 拣货分配 | P1 |
| `/wms/pick/execute` | 执行拣货 | P1 |
| `/wms/putaway/create` | 创建上架作业 | P1 |
| `/wms/putaway/execute` | 执行上架 | P1 |
| `/wms/count/plan/submit` | 盘点计划提交 | P1 |
| `/wms/count/job/execute` | 执行盘点 | P1 |
| `/wms/agv/task/dispatch` | AGV任务下发 | P2 |
| `/wms/agv/task/cancel` | AGV任务取消 | P2 |
| `/wms/label/print` | 标签打印 | P2 |
| `/wms/strategy/create` | 策略创建 | P2 |

### 3.3 缺失业务功能

| 功能 | 说明 | 优先级 |
|------|------|--------|
| 库区三级结构 | 仓库-库区-库位层级管理 | P1 |
| 货品档案 | MDM物料与WMS货品映射 | P1 |
| 容器全生命周期 | 绑定/解绑/维修/冻结 | P1 |
| 拣货下架 | 销售发货/领料拣货流程 | P1 |
| 上架策略 | 入库上架规则引擎 | P1 |
| 盘点完整流程 | 计划→任务→执行→审批→调账 | P1 |
| AGV调度 | 任务下发与跟踪 | P2 |
| 标签打印 | 模板设计与批量打印 | P2 |
| 策略规则引擎 | 上架/拣货/库位分配策略 | P2 |
| 库存异动追溯 | 完整库存流水链 | P2 |

**WMS完成度评估: 约15-20%**

---

## 四、SCP模块差异

### 4.1 缺失DDL表结构

| 表名 | 说明 | 优先级 |
|------|------|--------|
| `scp_purchase_main` | 采购订单主表 | P1 |
| `scp_purchase_detail` | 采购订单明细表 | P1 |
| `scp_purchase_plan_main` | 要货计划主表 | P1 |
| `scp_purchase_plan_detail` | 要货计划明细表 | P1 |
| `scp_supplierdeliver_request_main` | 供应商发货申请主表 | P1 |
| `scp_supplierdeliver_request_detail` | 供应商发货申请明细表 | P1 |
| `scp_supplierdeliver_record_main` | 供应商发货记录主表 | P1 |
| `scp_purchasereceipt_job_main` | 采购收货任务表 | P1 |
| `scp_purchasereturn_job_main` | 采购退货任务表 | P1 |
| `scp_supplier_item` | 供应商物料表 | P1 |
| `scp_purchasereturn_record_main` | 采购退货记录表 | P2 |
| `scp_supplier_user` | 供应商用户关联表 | P2 |
| `scp_supplier_apbalance_main` | 供应商应付款余额表 | P2 |
| `scp_purchase_mrs_statistics` | MRS汇总统计表 | P2 |
| `scp_customer_return_job_main` | 客户退货任务表 | P2 |

### 4.2 缺失API接口

| 路径 | 说明 | 优先级 |
|------|------|--------|
| `/scp/purchase-main/*` | 采购订单(40+接口) | P1 |
| `/scp/purchase-detail/*` | 采购订单明细 | P1 |
| `/scp/purchase-plan-main/*` | 要货计划(20+接口) | P1 |
| `/scp/supplierdeliver-request-main/*` | 供应商发货申请 | P1 |
| `/scp/supplierdeliver-record-main/*` | 供应商发货记录 | P1 |
| `/scp/purchasereceipt-job-main/*` | 采购收货任务 | P1 |
| `/scp/supplieritem/*` | 供应商物料 | P1 |
| `/scp/mrs/*` | MRS外部接口 | P1 |
| `/scp/purchasereturn-job-main/*` | 采购退货任务 | P2 |
| `/scp/supplier-user/*` | 供应商用户关联 | P2 |
| `/scp/supplier-apbalance-main/*` | 供应商应付款 | P2 |

### 4.3 缺失业务功能

| 功能 | 说明 | 优先级 |
|------|------|--------|
| M型采购订单 | 特殊订单类型 | P1 |
| 供应商发货申请 | 供应商主动发起发货 | P1 |
| 标签生成打印 | ASN标签生成 | P1 |
| 采购计划策略S001 | 策略驱动采购计划 | P1 |
| 收货任务承接/拒收 | 任务制收货流程 | P1 |
| 供应商物料关联 | 物料属性筛选 | P1 |
| MRS外部接口 | 与MRS系统对接 | P1 |
| QAD要货预测对接 | QAD数据同步 | P1 |
| 采购退货完整功能 | 多种退货类型 | P2 |
| 供应商用户/应付款 | 供应商门户 | P2 |
| 客户退货任务 | 客户退货流程 | P2 |

**SCP完成度评估: 约10-15%**

---

## 五、QMS模块差异

### 5.1 缺失DDL表结构

| 表名 | 说明 | 优先级 |
|------|------|--------|
| `qms_aql` | AQL标准配置表 | P1 |
| `qms_aql_detail` | AQL明细表 | P1 |
| `qms_material_aql` | 物料AQL配置表 | P1 |
| `qms_inspection_characteristics` | 检验特性表 | P1 |
| `qms_char_quantitative` | 定量特性表 | P1 |
| `qms_char_qualitative` | 定性特性表 | P1 |
| `qms_inspection_method` | 检验方法表 | P1 |
| `qms_inspection_scheme` | 检验方案表 | P1 |
| `qms_sampling_plan` | 抽样计划表 | P1 |
| `qms_sampling_rule` | 抽样规则表 | P1 |
| `qms_sampling_record` | 抽样记录表 | P1 |
| `qms_dynamic_rule` | 动态检验规则表 | P1 |
| `qms_sample_code` | 样本编码表 | P2 |

### 5.2 缺失API接口

| 路径 | 说明 | 优先级 |
|------|------|--------|
| `/qms/aql/*` | AQL标准管理 | P1 |
| `/qms/characteristics/*` | 检验特性管理 | P1 |
| `/qms/method/*` | 检验方法管理 | P1 |
| `/qms/scheme/*` | 检验方案管理 | P1 |
| `/qms/sampling/*` | 抽样方案管理 | P1 |
| `/qms/dynamic/*` | 动态规则 | P1 |

### 5.3 缺失业务功能

| 功能 | 说明 | 优先级 |
|------|------|--------|
| AQL标准管理 | 接收质量限标准配置 | P1 |
| 检验特性管理 | 定性/定量特性定义 | P1 |
| 检验方案配置 | 检验方案模板 | P1 |
| 抽样方案 | 抽样规则定义 | P1 |
| 动态检验规则 | 按条件自动调整检验级别 | P1 |
| 样本编码管理 | 样本追踪编码 | P2 |

**QMS完成度评估: 约60-70%** (AQL/抽样/特性已补充)

---

## 六、EAM模块差异

### 6.1 缺失DDL表结构

| 表名 | 说明 | 优先级 |
|------|------|--------|
| `basic_equipment_accounts` | 设备台账主表 | P1 |
| `basic_equipment_repair_record_main` | 维修记录主表 | P1 |
| `basic_equipment_repair_record_detail` | 维修记录明细表 | P1 |
| `equipment_report_repair_request` | 报修申请表 | P1 |
| `equipment_repair_job_main` | 维修工单主表 | P1 |
| `equipment_repair_job_detail` | 维修工单明细表 | P1 |
| `basic_inspection_option` | 巡检方案表 | P2 |
| `basic_inspection_item` | 巡检项表 | P2 |
| `plan_inspection` | 巡检计划表 | P2 |
| `basic_maintenance_option` | 保养方案表 | P2 |
| `maintenance_item` | 保养项表 | P2 |
| `maintenance` | 保养计划表 | P2 |
| `basic_spot_check_option` | 点检方案表 | P2 |
| `spot_check_item` | 点检项表 | P2 |
| `plan_spot_check` | 点检计划表 | P2 |
| `equipment_spot_check_record_main` | 点检记录主表 | P1 |
| `equipment_spot_check_record_detail` | 点检记录明细表 | P1 |
| `equipment_spot_check_main` | 点检工单主表 | P1 |
| `equipment_spot_check_detail` | 点检工单明细表 | P1 |
| `basic_equipment_tool_spare_part` | 备件表 | P2 |
| `transaction_eam` | 库存事务表 | P2 |
| `spare_parts_apply_main` | 备件申请主表 | P2 |
| `spare_parts_in_location_main` | 备件入库主表 | P2 |
| `spare_parts_out_location_main` | 备件出库主表 | P2 |
| `count_job_main` | 备件盘点任务主表 | P2 |
| `item_eam` | 备件物料表 | P2 |
| `item_accounts` | 备件台账表 | P2 |
| `tool_accounts` | 工装台账表 | P2 |
| `tool_equipment_in` | 工装入库表 | P2 |
| `tool_equipment_out` | 工装出库表 | P2 |
| `repair_experience` | 维修经验记录表 | P2 |
| `basic_eam_workshop` | EAM车间表 | P2 |
| `basic_eam_productionline` | 生产线表 | P2 |

### 6.2 缺失API接口

| 路径 | 说明 | 优先级 |
|------|------|--------|
| `/eam/equipment-accounts/*` | 设备台账完整CRUD | P1 |
| `/eam/equipment-report-repair-request/*` | 报修申请 | P1 |
| `/eam/equipment-repair-job-main/*` | 维修工单 | P1 |
| `/eam/equipment-inspection-main/*` | 巡检工单 | P1 |
| `/eam/equipment-maintenance-main/*` | 保养工单 | P1 |
| `/eam/equipment-spot-check-main/*` | 点检工单 | P1 |
| `/eam/equipment-spot-check-record-main/*` | 点检记录 | P1 |
| `/eam/basic-eam-workshop/*` | EAM车间管理 | P2 |
| `/eam/basic-eam-productionline/*` | 生产线管理 | P2 |
| `/eam/basic-inspection-option/*` | 巡检方案 | P2 |
| `/eam/basic-maintenance-option/*` | 保养方案 | P2 |
| `/eam/basic-spot-check-option/*` | 点检方案 | P2 |
| `/eam/plan-inspection/*` | 巡检计划 | P2 |
| `/eam/maintenance/*` | 保养计划 | P2 |
| `/eam/plan-spot-check/*` | 点检计划 | P2 |
| `/eam/item/*` | 备件物料 | P2 |
| `/eam/item-accounts/*` | 备件台账 | P2 |
| `/eam/transaction/*` | 库存事务 | P2 |
| `/eam/spare-parts-apply-main/*` | 备件申请 | P2 |
| `/eam/spare-parts-in-location-main/*` | 备件入库 | P2 |
| `/eam/spare-parts-out-location-main/*` | 备件出库 | P2 |
| `/eam/countJobMain/*` | 备件盘点任务 | P2 |
| `/eam/tool/tool-accounts/*` | 工装台账 | P2 |
| `/eam/record/repair-experience/*` | 维修经验 | P2 |
| `/eam/record/maintain-experience/*` | 保养经验 | P2 |

### 6.3 缺失业务功能

| 功能 | 说明 | 优先级 |
|------|------|--------|
| 设备维修完整流程 | 报修→审核→工单→执行→完工→记录 | P1 |
| 设备点检完整流程 | 计划→工单→执行→记录 | P1 |
| 设备巡检完整流程 | 计划→工单→执行→记录 | P1 |
| 设备保养完整流程 | 计划→工单→执行→完工→经验 | P1 |
| 设备台账全字段 | 含功率/运行时长/停机率/OEE | P1 |
| EAM车间管理 | 厂区-车间-产线-工位四级 | P2 |
| 设备厂商/供应商管理 | 完整档案字段 | P2 |
| 故障类型/原因体系 | 故障分析统计 | P2 |
| 备件台账与出入库 | 备件全流程管理 | P2 |
| 工装台账管理 | 工装出入库/库存 | P2 |
| 巡检/保养/点检方案 | 方案→项→选择集 | P2 |
| 保养/维修经验库 | 经验积累复用 | P2 |

**EAM完成度评估: 约20-25%**

---

## 七、BPM模块差异

### 7.1 缺失DDL表结构

| 表名 | 说明 | 优先级 |
|------|------|--------|
| `bpm_task_message_rule` | 任务消息规则表 | P2 |

### 7.2 缺失API接口

| 路径 | 说明 | 优先级 |
|------|------|--------|
| `BpmProcessInstanceApi` | 跨模块API接口供其他模块调用 | P1 |
| `/bpm/process-definition/get/{id}` | 获取流程定义详情 | P3 |
| `/bpm/process-definition/page` | 流程定义分页查询 | P2 |

### 7.3 缺失业务功能

| 功能 | 说明 | 优先级 |
|------|------|--------|
| BpmMessageService | 流程消息通知服务 | P1 |
| BpmProcessInstanceResultEventListener | 审批结果事件监听 | P1 |
| 流程定义扩展服务 | BpmProcessDefinitionExtService | P2 |
| 任务转移/移交 | Transfer Task功能 | P2 |
| 脚本计算任务分配 | 动态脚本计算处理人 | P2 |
| 任务候选人与候选组 | Candidate Users/Groups | P2 |
| BpmOALeaveResultListener | OA请假结果监听 | P2 |

**BPM完成度评估: 约70-80%**

---

## 八、优先级汇总

### P1 - 核心功能（必须实现）

| 模块 | 缺失项 |
|------|--------|
| MES | 日计划CRUD/publish/stop、工单排程、工序报工/完工/质检、齐套检查 |
| WMS | 库区三级结构、货品档案、容器生命周期、拣货下架、上架策略、盘点流程 |
| SCP | 采购订单完整API、供应商发货申请、收货任务、供应商物料关联、MRS接口 |
| QMS | AQL/检验特性/抽样方案/动态规则（已补充大部分） |
| EAM | 设备维修/点检/巡检/保养完整流程、设备台账 |
| BPM | 跨模块API接口、消息通知服务、审批结果监听 |

### P2 - 重要功能（短期内应实现）

| 模块 | 缺失项 |
|------|--------|
| MES | 日计划子表、返工返修、质检管理、班组日历、能力矩阵、模具管理、月计划 |
| WMS | AGV调度、标签打印、策略引擎、库存异动追溯 |
| SCP | 采购退货、供应商用户/应付款、客户退货、MRS统计 |
| EAM | 车间/产线、厂商/供应商/部件、故障分析、备件/工装管理、方案/项/计划 |
| BPM | 任务转移、脚本分配、候选人/组、流程定义查询 |

### P3 - 增强功能（中长期规划）

| 模块 | 缺失项 |
|------|--------|
| MES | 工位能力关联、工位工单实时/历史、工序资源关联、完工库存 |
| WMS | 看板、首页数据 |
| EAM | 设备到货签收、变更记录、文档类型、附件管理 |
| BPM | 流程定义直接查询 |

---

## 九、实施建议

1. **第一阶段(P1)**：补齐各模块核心业务闭环
   - MES: 日计划→工单→工序→报工→完工
   - WMS: 库区→货品→容器→拣货→上架→盘点
   - SCP: 采购订单→发货申请→收货→退货
   - EAM: 设备台账→维修/点检/巡检/保养

2. **第二阶段(P2)**：完善支撑功能和配置
   - 班组日历、能力矩阵、质检配置
   - AGV调度、标签打印、策略引擎
   - 供应商门户、备件管理

3. **第三阶段(P3)**：增强功能和优化
   - 工位精细化管理
   - 看板/报表
   - 高级查询和导出
