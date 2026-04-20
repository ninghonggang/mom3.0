# SFMS3.0 vs MOM3.0 完整对比清单

> 生成时间: 2026-04-17
> 说明: 完整对比sfms3.0目录下所有设计文档与MOM3.0当前设计文档的差异

---

## 一、MES模块对比

### 1.1 缺失表结构（约40+张）

| 序号 | 表名 | 说明 | 优先级 |
|------|------|------|--------|
| 1 | plan_mes_order_day | 日计划主表 | P1 |
| 2 | plan_mes_work_scheduling | 工单排程主表 | P1 |
| 3 | plan_mes_work_scheduling_detail | 工单排程明细表 | P1 |
| 4 | mes_device_info | 设备基础信息表 | P1 |
| 5 | mes_ability_info | 能力信息表 | P1 |
| 6 | hr_person_ability | 人员能力矩阵表 | P1 |
| 7 | mes_process_itembasic | 工序物料关联表 | P1 |
| 8 | mes_process_pattern | 工序模具关联表 | P1 |
| 9 | mes_process_productionline | 工序产线关联表 | P1 |
| 10 | mes_pattern | 模具台账表 | P2 |
| 11 | mes_pattern_type | 模具类型表 | P2 |
| 12 | mes_team_setting | 班组管理表 | P2 |
| 13 | mes_job_report_log | 报工日志表 | P2 |
| 14 | mes_workstation | 工位台账表 | P1 |
| 15 | mes_workstation_ability | 工位能力关联表 | P1 |
| 16 | mes_workstation_equipment | 工位设备关联表 | P1 |
| 17 | mes_workstation_opersteps | 工位操作步骤表 | P1 |
| 18 | mes_workstation_post | 工位岗位关联表 | P1 |
| 19 | mes_workstation_order | 工位工单实时表 | P2 |
| 20 | mes_workstation_order_history | 工位工单历史表 | P2 |
| 21 | mes_order_oper_log | 操作流水日志表 | P2 |
| 22 | plan_mes_order_month_main | 月计划主表 | P2 |
| 23 | plan_mes_order_month_sub | 月计划子表 | P2 |
| 24 | qms_qualityclass | 质检类别表 | P1 |
| 25 | qms_qualitygroup | 质检分组表 | P1 |
| 26 | qms_item | 质检项目表 | P1 |
| 27 | qms_qualityform | 质检表单表 | P1 |
| 28 | qms_qualityformdetail | 质检表单明细表 | P1 |
| 29 | qms_qualityformlog | 质检表单日志表 | P1 |
| 30 | mes_product_offline | 产品离线表 | P2 |
| 31 | mes_product_backline | 产品返线表 | P2 |
| 32 | mes_reportfinish_store | 完工库存中间表 | P2 |
| 33 | mes_reportp_store | 报工物料明细表 | P2 |

### 1.2 缺失API模块（约200+ API）

| 序号 | API模块 | 说明 | 优先级 |
|------|---------|------|--------|
| 1 | /mes/device-info/* | 设备基础信息管理(9个API) | P1 |
| 2 | /mes/ability-info/* | 能力信息管理 | P1 |
| 3 | /mes/hr-person-ability/* | 人员能力矩阵 | P1 |
| 4 | /mes/mes-process-itembasic/* | 工序物料关联 | P1 |
| 5 | /mes/mes-process-pattern/* | 工序模具关联 | P1 |
| 6 | /mes/mes-process-productionline/* | 工序产线关联 | P1 |
| 7 | /mes/qualityclass/* | 质检类别管理 | P1 |
| 8 | /mes/qualitygroup/* | 质检分组管理 | P1 |
| 9 | /mes/item/* | 质检项目管理 | P1 |
| 10 | /mes/qualityform/* | 质检表单管理 | P1 |
| 11 | /mes/qualityformdetail/* | 质检表单明细 | P1 |
| 12 | /mes/qualityformlog/* | 质检表单日志 | P1 |
| 13 | /mes/teamSetting/* | 班组管理 | P2 |
| 14 | /mes/workstation/* | 工位管理(完整CRUD) | P1 |
| 15 | /mes/workstation-ability/* | 工位能力关联 | P1 |
| 16 | /mes/workstation-opersteps/* | 工位操作步骤 | P1 |
| 17 | /mes/workstation-post/* | 工位岗位关联 | P1 |
| 18 | /mes/workstation-order/* | 工位工单实时 | P2 |
| 19 | /mes/product-offline/* | 产品离线 | P2 |
| 20 | /mes/product-backline/* | 产品返线 | P2 |
| 21 | /mes/reportfinish-store/* | 完工库存中间 | P2 |
| 22 | /mes/reportp-store/* | 报工物料明细 | P2 |
| 23 | /mes/mes-job-report-log/* | 报工日志管理 | P2 |
| 24 | /mes/pattern/* | 模具管理 | P2 |
| 25 | /mes/pattern-type/* | 模具类型管理 | P2 |
| 26 | /mes/order-month-main/* | 月计划主管理 | P2 |
| 27 | /mes/order-month-sub/* | 月计划子管理 | P2 |

---

## 二、SCP模块对比

### 2.1 缺失表结构（约15张）

| 序号 | 表名 | 说明 | 优先级 |
|------|------|------|--------|
| 1 | scp_purchase | 采购订单主表 | P1 |
| 2 | scp_purchase_detail | 采购订单子表 | P1 |
| 3 | scp_purchase_plan_item | 要货计划明细表 | P1 |
| 4 | scp_supplierdeliver_request_main | 供应商发货申请主表 | P1 |
| 5 | scp_supplierdeliver_request_detail | 供应商发货申请子表 | P1 |
| 6 | scp_supplierdeliver_record_main | 供应商发货记录主表 | P1 |
| 7 | scp_purchasereceipt_job_main | 采购收货任务主表 | P1 |
| 8 | scp_purchasereturn_job_main | 采购退货任务主表 | P1 |
| 9 | scp_purchasereturn_record_main | 采购退货记录主表 | P1 |
| 10 | scp_supplier_item | 供应商物料表 | P2 |
| 11 | scp_supplier_user | 供应商用户关联表 | P2 |
| 12 | scp_supplier_apbalance_main | 供应商应付款余额表 | P2 |
| 13 | scp_customer_delivery_forecast | 客户发货预测表 | P2 |
| 14 | scp_demandforecasting_detail | 要货预测子表 | P2 |
| 15 | scp_customerreturn_job_main | 客户退货任务主表 | P2 |
| 16 | scp_purchase_mrs_statistics | MRS统计表 | P2 |

### 2.2 缺失API接口（约180+个）

| 序号 | API模块 | 说明 | 优先级 |
|------|---------|------|--------|
| 1 | /wms/purchase/* | 采购订单主(20个API) | P1 |
| 2 | /wms/purchase-detail/* | 采购订单子(25个API) | P1 |
| 3 | /wms/purchase-plan-main/* | 要货计划主(25个API) | P1 |
| 4 | /wms/supplierdeliver-request-main/* | 供应商发货申请(25个API) | P1 |
| 5 | /wms/supplierdeliver-record-main/* | 供应商发货记录(15个API) | P1 |
| 6 | /wms/purchasereceipt-job-main/* | 采购收货任务(20个API) | P1 |
| 7 | /wms/purchasereturn-job-main/* | 采购退货任务(15个API) | P1 |
| 8 | /scp/supplier-item/* | 供应商物料(25个API) | P2 |
| 9 | /scp/supplier-user/* | 供应商用户关联(12个API) | P2 |
| 10 | /scp/mrs-statistics/* | MRS统计(8个API) | P2 |
| 11 | /scp/customerreturn-job-main/* | 客户退货任务(15个API) | P2 |

### 2.3 缺失业务流程

| 序号 | 流程 | 说明 | 优先级 |
|------|------|------|--------|
| 1 | 采购执行流程 | 收货任务→执行→质检→入库 | P1 |
| 2 | 供应商发货管理 | 发货申请→自检报告→标签→发货记录→ASN | P1 |
| 3 | 采购退货流程 | 退货任务→执行→退货记录 | P1 |
| 4 | MRS统计流程 | 工单/预测/安全库存→MRS→采购计划 | P2 |
| 5 | 客户退货流程 | 退货任务→承接→执行→关闭 | P2 |

---

## 三、QMS模块对比

### 3.1 缺失表结构

| 序号 | 表名 | 说明 | 优先级 |
|------|------|------|--------|
| 1 | qms_counter | 计数器表 | P1 |
| 2 | qms_inspection_process | 检验工序表 | P1 |
| 3 | qms_inspection_stage | 检验阶段表 | P1 |
| 4 | qms_programme_template | 检验方案模板表 | P1 |
| 5 | qms_inspection_q1 | Q1通知单表 | P1 |
| 6 | qms_inspection_q2 | Q2通知单表 | P1 |
| 7 | qms_inspection_q3_main/detail | Q3通知单主/子表 | P1 |
| 8 | qms_job_inspection_main | 检验任务主表 | P2 |
| 9 | qms_job_inspection_detail | 检验任务明细表 | P2 |
| 10 | qms_job_inspection_package | 检验任务包装表 | P2 |
| 11 | qms_job_inspection_characteristics | 检验任务特性表 | P2 |
| 12 | qms_record_inspection_main | 检验记录主表 | P2 |
| 13 | qms_record_inspection_detail | 检验记录明细表 | P2 |
| 14 | qms_request_inspection_main | 检验请求主表 | P1 |
| 15 | qms_request_inspection_package | 检验请求包装表 | P1 |
| 16 | qms_selected_set | 选定集表 | P2 |
| 17 | qms_selected_project | 选定集项目表 | P2 |
| 18 | qms_staging_json_job | 暂存JSON表 | P2 |

### 3.2 缺失API接口

| 序号 | API模块 | 说明 | 优先级 |
|------|---------|------|--------|
| 1 | /qms/counter/* | 物料检验计数器(14个API) | P1 |
| 2 | /qms/inspection-process/* | 检验工序(10个API) | P1 |
| 3 | /qms/inspection-stage/* | 检验阶段(11个API) | P1 |
| 4 | /qms/programme-template/* | 检验方案模板(14个API) | P1 |
| 5 | /qms/inspection-q1/* | Q1通知单(9个API) | P1 |
| 6 | /qms/inspection-q2/* | Q2通知单(11个API) | P1 |
| 7 | /qms/inspection-q3/* | Q3通知单(9+10+8个API) | P1 |
| 8 | /qms/request-inspection-main/* | 检验请求主(14个API) | P1 |
| 9 | /qms/selected-set/* | 选定集(14个API) | P2 |
| 10 | /qms/sampling-process/* | 采样过程(14个API) | P2 |

### 3.3 缺失业务功能

| 序号 | 功能 | 说明 | 优先级 |
|------|------|------|--------|
| 1 | Q1/Q2/Q3质量通知单 | 完整质量异常通知流程 | P1 |
| 2 | 检验申请管理 | 与WMS/MES集成接口 | P1 |
| 3 | 检验工序/阶段管理 | 检验方案基础数据 | P1 |
| 4 | 检验方案模板 | 检验方案复用机制 | P1 |
| 5 | OQC出货检验完整流程 | 销售订单触发→出货报告 | P1 |
| 6 | AQL抽样标准技术实现 | GB/T 2828.1标准 | P2 |
| 7 | 选定集管理 | 检验项目批量选择 | P2 |

---

## 四、EAM模块对比

### 4.1 缺失表结构

| 序号 | 表名 | 说明 | 优先级 |
|------|------|------|--------|
| 1 | eam_spare_parts_apply_main/detail | 备件申请主/子表 | P2 |
| 2 | eam_spare_parts_in_location_main/detail | 备件入库主/子表 | P2 |
| 3 | eam_spare_parts_out_location_main/detail | 备件出库主/子表 | P2 |
| 4 | eam_count_job_main/detail | 盘点任务主/子表 | P2 |
| 5 | eam_count_record_main/detail | 盘点记录主/子表 | P2 |
| 6 | eam_tool_tool_accounts | 工装台账表 | P2 |
| 7 | eam_tool_equipment_in | 工装入库表 | P2 |
| 8 | eam_tool_equipment_out | 工装出库表 | P2 |
| 9 | eam_repair_spare_parts_request | 备件维修申请表 | P3 |
| 10 | eam_record_maintain_experience | 保养经验记录表 | P3 |
| 11 | eam_record_repair_experience | 维修经验记录表 | P3 |

### 4.2 缺失API接口

| 序号 | API模块 | 说明 | 优先级 |
|------|---------|------|--------|
| 1 | /eam/spare-parts-apply/* | 备件申请(完整CRUD) | P2 |
| 2 | /eam/spare-parts-in-location/* | 备件入库(完整CRUD) | P2 |
| 3 | /eam/spare-parts-out-location/* | 备件出库(完整CRUD) | P2 |
| 4 | /eam/count-job/* | 备件盘点任务(完整CRUD) | P2 |
| 5 | /eam/tool-account/* | 工装台账(完整CRUD) | P2 |
| 6 | /eam/tool-equipment-in/* | 工装入库(完整CRUD) | P2 |
| 7 | /eam/tool-equipment-out/* | 工装出库(完整CRUD) | P2 |

---

## 五、BPM模块对比

### 5.1 缺失表结构

| 序号 | 表名 | 说明 | 优先级 |
|------|------|------|--------|
| 1 | bpm_form | 表单配置表 | P1 |
| 2 | bpm_task_assign_rule | 任务分配规则表 | P1 |
| 3 | bpm_user_group | 用户组表 | P1 |
| 4 | bpm_oa_leave | OA请假表 | P3 |

### 5.2 缺失API接口

| 序号 | API模块 | 说明 | 优先级 |
|------|---------|------|--------|
| 1 | /bpm/models/* | 流程模型CRUD | P1 |
| 2 | /bpm/forms/* | 表单配置CRUD | P1 |
| 3 | /bpm/user-group/* | 用户组CRUD | P1 |
| 4 | /bpm/task-assign-rule/* | 任务分配规则CRUD | P1 |
| 5 | /bpm/tasks/{id}/return | 任务退回 | P2 |
| 6 | /bpm/tasks/{id}/update-assignee | 任务转派 | P2 |

### 5.3 缺失技术实现

| 序号 | 技术 | 说明 | 优先级 |
|------|------|------|--------|
| 1 | Flowable流程引擎集成 | BPMN2.0流程引擎 | P1 |
| 2 | BPMN可视化设计器 | 流程模型图形化设计 | P1 |
| 3 | 表单设计器 | 可视化表单配置 | P1 |
| 4 | 事件监听机制 | 流程事件处理 | P2 |

---

## 六、Infra_System模块对比

### 6.1 缺失表结构

| 序号 | 表名 | 说明 | 优先级 |
|------|------|------|--------|
| 1 | system_notice | 通知公告表 | P1 |
| 2 | system_sms_* | 短信相关3表 | P2 |
| 3 | system_mail_* | 邮件相关3表 | P2 |
| 4 | system_oauth2_* | OAuth2相关2表 | P2 |
| 5 | system_sensitive_word | 敏感词表 | P3 |
| 6 | system_serial_number | 序列号生成表 | P3 |
| 7 | system_area | 地区表 | P3 |
| 8 | infra_job | 定时任务表 | P2 |

### 6.2 缺失API接口

| 序号 | API模块 | 说明 | 优先级 |
|------|---------|------|--------|
| 1 | /system/notice/* | 通知公告管理 | P1 |
| 2 | /system/tenant/* | 多租户管理 | P1 |
| 3 | /system/sms/* | 短信服务 | P2 |
| 4 | /system/mail/* | 邮件服务 | P2 |
| 5 | /system/oauth2/* | OAuth2认证 | P2 |
| 6 | /infra/codegen/* | 代码生成器 | P3 |

### 6.3 缺失业务功能

| 序号 | 功能 | 说明 | 优先级 |
|------|------|------|--------|
| 1 | 通知公告 | 系统公告发布管理 | P1 |
| 2 | 多车间管理 | 车间配置管理 | P1 |
| 3 | 打印模板管理 | 标签/单据模板 | P2 |
| 4 | 短信服务 | 短信发送管理 | P2 |
| 5 | 邮件服务 | 邮件发送管理 | P2 |
| 6 | OAuth2认证 | 第三方登录 | P2 |

---

## 七、Report模块对比

### 7.1 缺失表结构

| 序号 | 表名 | 说明 | 优先级 |
|------|------|------|--------|
| 1 | report_go_view_project | GoView项目配置表 | P1 |
| 2 | report_aj_data_source | AJ-Report数据源表 | P2 |

### 7.2 缺失API接口

| 序号 | API模块 | 说明 | 优先级 |
|------|---------|------|--------|
| 1 | /report/go-view/project/* | GoView项目管理(CRUD) | P1 |
| 2 | /report/go-view/data/* | 大屏数据接口 | P1 |
| 3 | /report/aj/* | AJ-Report集成接口 | P2 |

### 7.3 缺失业务功能

| 序号 | 功能 | 说明 | 优先级 |
|------|------|------|--------|
| 1 | GoView大屏项目管理 | 大屏配置管理 | P1 |
| 2 | AJ-Report积木报表 | 复杂报表设计 | P2 |
| 3 | 自定义报表设计器 | 报表可视化设计 | P2 |
| 4 | 产能利用率报表 | 生产分析报表 | P2 |
| 5 | 设备故障分析报表 | 设备分析报表 | P2 |

---

## 八、优先级汇总

### P0 - 核心缺失（严重影响业务闭环）

| 模块 | 缺失内容 |
|------|---------|
| MES | 质检管理(qms_*)、工位管理(mes_workstation_*)、设备基础信息 |
| SCP | 采购执行流程、供应商发货管理、收货任务 |
| QMS | Q1/Q2/Q3通知单、检验申请管理 |
| BPM | Flowable引擎、表单设计器 |

### P1 - 重要缺失（影响完整功能）

| 模块 | 缺失内容 |
|------|---------|
| MES | 能力矩阵、班组管理、月计划、工艺资源关联 |
| SCP | 退货流程、MRS统计 |
| QMS | 检验工序/阶段、方案模板、OQC |
| EAM | 备件/工装完整业务链路 |
| BPM | 流程模型设计、用户组、任务分配规则 |
| Infra | 通知公告、多车间、短信邮件 |
| Report | GoView项目管理、AJ-Report集成 |

### P2 - 一般缺失（增强功能）

| 模块 | 缺失内容 |
|------|---------|
| MES | 报工日志、产品离线返线、模具管理 |
| SCP | 供应商物料关联、MRS外部接口 |
| QMS | 选定集、采样过程 |
| EAM | 备件维修经验记录 |
| Report | 自定义报表设计器 |

---

## 九、总体评估

| 模块 | 完整度 | 主要差距 |
|------|--------|---------|
| **MES** | ~40% | 质检管理、工位管理、设备基础信息大量缺失 |
| **WMS** | ~100% | 从sfms3.0完整复制 |
| **SCP** | ~30% | 采购执行流程、供应商发货管理缺失 |
| **QMS** | ~50% | Q1/Q2/Q3通知单、检验申请、OQC缺失 |
| **EAM** | ~60% | 备件/工装业务链路缺失 |
| **BPM** | ~15% | Flowable、表单设计器、模型设计器完全缺失 |
| **Infra** | ~50% | 通知公告、多车间、短信邮件缺失 |
| **Report** | ~40% | GoView项目、AJ-Report集成缺失 |

**建议优先级**：
1. **第一阶段**：MES质检管理、BPM核心引擎、SCP采购执行流程
2. **第二阶段**：QMS通知单、EAM备件链路、Infra通知公告
3. **第三阶段**：Report大屏管理、能力矩阵、班组管理等增强功能
