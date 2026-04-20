# MOM3.0前端 QMS质量管理模块设计文档

**版本**: V2.0 | **所属模块**: M05质量管理 | **基于**: [MOM3.0_主设计文档](./MOM3.0_主设计文档.md)

---

## 1. 模块概述

### 1.1 功能定位

QMS质量管理模块覆盖来料、过程、成品、出货全流程质量检验，配合NCR处理、质量改进、供应商评审等质量工具，实现全面的质量管控与追溯。

### 1.2 核心功能

| 功能 | 说明 |
|------|------|
| 检验计数 | 检验数据统计与计数管理 |
| 检验方案 | 检验方案定义与模板管理 |
| 检验方法 | 检验方法库与方法定义 |
| 检验阶段 | 检验阶段配置管理 |
| 定性/定量特性 | 检验特性定义与配置 |
| 抽样方案 | 抽样规则与抽样标准管理 |
| 检验请求 | 检验申请单创建与审批 |
| IQC/PQC/OQC | 来料/过程/出货检验记录 |
| 首件/巡检 | 首件检验与巡检管理 |
| 质量通知 | 质量通知发布与分发 |
| 质量标准 | 质量标准库维护 |
| 质量目标 | 质量目标设定与跟踪 |
| 质量分析 | 质量分析与报表 |
| 退货/索赔 | 退货与索赔管理 |
| 供应商评审 | 供应商质量评审 |
| 纠正/预防措施 | CAPA管理 |
| 质量培训 | 培训计划与记录 |
| 质量审核 | 审核计划与执行 |
| 质量证书 | 证书到期管理 |
| 计量/校准 | 计量器具与校准记录 |
| 质量追溯 | 追溯查询与明细 |

---

## 2. 页面清单

| 页面 | 路由路径 | 核心功能 |
|------|----------|----------|
| 计数器 | `/quality/counter` | 检验计数、计数配置 |
| 检验方案 | `/quality/inspectScheme` | 方案创建、方案模板、方案复制 |
| 检验方法 | `/quality/inspectMethod` | 方法定义、方法库 |
| 检验阶段 | `/quality/inspectStage` | 阶段配置、阶段类型 |
| 定性特性 | `/quality/qualitativeFeature` | 定性特性定义、特性模板 |
| 定量特性 | `/quality/quantitativeFeature` | 定量特性定义、特性范围 |
| 抽样方案 | `/quality/samplingScheme` | 抽样规则、抽样标准 |
| 抽样过程 | `/quality/samplingProcess` | 抽样记录、抽样执行 |
| 抽样代码 | `/quality/samplingCode` | 代码规则、代码生成 |
| 检验请求 | `/quality/inspectRequest` | 申请单创建、申请单审批 |
| 检验任务 | `/quality/inspectTask` | 任务分配、任务执行 |
| IQC记录 | `/quality/iqcRecord` | 来料检验记录详情 |
| PQC记录 | `/quality/pqcRecord` | 过程检验记录详情 |
| OQC记录 | `/quality/oqcRecord` | 出货检验记录详情 |
| 首件检验 | `/quality/firstInspect` | 首件记录、首件确认 |
| 巡检记录 | `/quality/patrolInspect` | 巡检计划、巡检执行 |
| 质量通知 | `/quality/qualityNotice` | 通知创建、通知分发 |
| 质量标准 | `/quality/std` | 标准库、标准维护 |
| 质量目标 | `/quality/target` | 目标设定、目标跟踪 |
| 质量分析 | `/quality/analysis` | 分析报表、趋势分析 |
| 质量报表 | `/quality/report` | 报表配置、报表生成 |
| 退货管理 | `/quality/return` | 退货申请、退货处理 |
| 索赔管理 | `/quality/claim` | 索赔登记、索赔处理 |
| 供应商评审 | `/quality/supplierAudit` | 评审计划、评审记录 |
| 纠正措施 | `/quality/corrective` | 纠正申请、纠正跟踪 |
| 预防措施 | `/quality/preventive` | 预防申请、预防跟踪 |
| 质量培训 | `/quality/training` | 培训计划、培训记录 |
| 质量审核 | `/quality/audit` | 审核计划、审核执行 |
| 质量证书 | `/quality/certificate` | 证书管理、证书到期 |
| 计量管理 | `/quality/measure` | 计量器具、检定计划 |
| 校准记录 | `/quality/calibration` | 校准记录、校准历史 |
| 质量追溯 | `/quality/trace` | 追溯查询、追溯明细 |

---

## 3. UI设计规范

### 3.1 页面基本结构

同MES模块标准布局：搜索+工具栏+表格+详情弹窗。

### 3.2 状态映射

**检验单状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| PENDING | warning | 待检验 |
| INSPECTING | primary | 检验中 |
| PASSED | success | 合格 |
| FAILED | danger | 不合格 |
| ACCEPTED | success | 已接受 |
| REJECTED | danger | 已拒绝 |

**检验结果状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| OK | success | 正常 |
| NG | danger | 异常 |
| NA | info | 不适用 |

**NCR状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| OPEN | warning | 待处理 |
| INVESTIGATING | primary | 处理中 |
| RESOLVED | success | 已解决 |
| CLOSED | info | 已关闭 |
| CANCELLED | info | 已取消 |

**质量通知状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| DRAFT | info | 草稿 |
| PUBLISHED | primary | 已发布 |
| URGENT | danger | 紧急 |
| ARCHIVED | info | 已归档 |

---

## 4. 业务流程

### 4.1 检验流程

```
检验请求 → 检验任务分配 → 执行检验 → 记录结果 → 合格/不合格处理
```

### 4.2 NCR处理流程

```
NCR创建 → 原因分析(5Why) → 纠正措施 → 预防措施 → 验证关闭
```

### 4.3 供应商评审流程

```
评审计划 → 执行评审 → 记录结果 → 评审结论 → 持续改进
```

---

## 5. 数据模型

### 5.1 计数器 (quality_counter)

| 字段 | 类型 | 说明 |
|------|------|------|
| counter_id | BIGINT | 计数器ID |
| counter_code | VARCHAR(50) | 计数器编码 |
| counter_name | VARCHAR(200) | 计数器名称 |
| counter_type | VARCHAR(20) | 计数类型 |
| count_value | INT | 计数值 |
| unit | VARCHAR(20) | 单位 |
| inspection_id | BIGINT | 关联检验ID |

### 5.2 检验方案 (quality_inspect_scheme)

| 字段 | 类型 | 说明 |
|------|------|------|
| scheme_id | BIGINT | 方案ID |
| scheme_code | VARCHAR(50) | 方案编码 |
| scheme_name | VARCHAR(200) | 方案名称 |
| scheme_type | VARCHAR(20) | 方案类型 |
| template_id | BIGINT | 模板ID |
| status | VARCHAR(20) | 状态 |

### 5.3 检验方法 (quality_inspect_method)

| 字段 | 类型 | 说明 |
|------|------|------|
| method_id | BIGINT | 方法ID |
| method_code | VARCHAR(50) | 方法编码 |
| method_name | VARCHAR(200) | 方法名称 |
| method_type | VARCHAR(20) | 方法类型 |
| description | TEXT | 方法描述 |

### 5.4 检验阶段 (quality_inspect_stage)

| 字段 | 类型 | 说明 |
|------|------|------|
| stage_id | BIGINT | 阶段ID |
| stage_code | VARCHAR(50) | 阶段编码 |
| stage_name | VARCHAR(200) | 阶段名称 |
| stage_type | VARCHAR(20) | 阶段类型 |
| sequence | INT | 顺序 |

### 5.5 定性特性 (quality_qualitative_feature)

| 字段 | 类型 | 说明 |
|------|------|------|
| feature_id | BIGINT | 特性ID |
| feature_code | VARCHAR(50) | 特性编码 |
| feature_name | VARCHAR(200) | 特性名称 |
| standard | VARCHAR(500) | 合格标准 |
| defect_class | VARCHAR(20) | 缺陷分类 |
| inspection_tool | VARCHAR(100) | 检验工具 |

### 5.6 定量特性 (quality_quantitative_feature)

| 字段 | 类型 | 说明 |
|------|------|------|
| feature_id | BIGINT | 特性ID |
| feature_code | VARCHAR(50) | 特性编码 |
| feature_name | VARCHAR(200) | 特性名称 |
| usl | DECIMAL(18,6) | 上限 |
| lsl | DECIMAL(18,6) | 下限 |
| target | DECIMAL(18,6) | 目标值 |
| unit | VARCHAR(20) | 单位 |

### 5.7 抽样方案 (quality_sampling_scheme)

| 字段 | 类型 | 说明 |
|------|------|------|
| scheme_id | BIGINT | 方案ID |
| scheme_code | VARCHAR(50) | 方案编码 |
| scheme_name | VARCHAR(200) | 方案名称 |
| sampling_type | VARCHAR(20) | 抽样类型 |
| sample_size | INT | 抽样数量 |
| acceptance_num | INT | 接收数 |

### 5.8 抽样过程 (quality_sampling_process)

| 字段 | 类型 | 说明 |
|------|------|------|
| record_id | BIGINT | 记录ID |
| record_no | VARCHAR(50) | 记录编号 |
| sampling_time | TIMESTAMP | 抽样时间 |
| sampler_id | BIGINT | 抽样人ID |
| sampler_name | VARCHAR(100) | 抽样人姓名 |
| sampling_location | VARCHAR(200) | 抽样地点 |
| result | VARCHAR(20) | 抽样结果 |

### 5.9 抽样代码 (quality_sampling_code)

| 字段 | 类型 | 说明 |
|------|------|------|
| code_id | BIGINT | 代码ID |
| code | VARCHAR(50) | 样本编码 |
| code_rule | VARCHAR(200) | 编码规则 |
| prefix | VARCHAR(20) | 前缀 |
| suffix | VARCHAR(20) | 后缀 |
| current_value | BIGINT | 当前值 |

### 5.10 检验请求 (quality_inspect_request)

| 字段 | 类型 | 说明 |
|------|------|------|
| request_id | BIGINT | 申请单ID |
| request_no | VARCHAR(50) | 申请单号 |
| request_type | VARCHAR(20) | 申请类型 |
| source_type | VARCHAR(20) | 来源类型 |
| source_no | VARCHAR(50) | 来源单号 |
| status | VARCHAR(20) | 状态 |
| applicant_id | BIGINT | 申请人ID |
| apply_time | TIMESTAMP | 申请时间 |

### 5.11 检验任务 (quality_inspect_task)

| 字段 | 类型 | 说明 |
|------|------|------|
| task_id | BIGINT | 任务ID |
| task_no | VARCHAR(50) | 任务编号 |
| request_id | BIGINT | 关联申请ID |
| inspector_id | BIGINT | 检验员ID |
| inspector_name | VARCHAR(100) | 检验员姓名 |
| assign_time | TIMESTAMP | 分派时间 |
| status | VARCHAR(20) | 状态 |

### 5.12 IQC记录 (quality_iqc_record)

| 字段 | 类型 | 说明 |
|------|------|------|
| record_id | BIGINT | 记录ID |
| record_no | VARCHAR(50) | 记录编号 |
| inspect_no | VARCHAR(50) | 检验单号 |
| material_id | BIGINT | 物料ID |
| material_code | VARCHAR(50) | 物料编码 |
| batch_no | VARCHAR(50) | 批次号 |
| inspect_qty | DECIMAL(18,3) | 检验数量 |
| pass_qty | DECIMAL(18,3) | 合格数量 |
| fail_qty | DECIMAL(18,3) | 不合格数量 |
| result | VARCHAR(20) | 检验结果 |
| inspector_id | BIGINT | 检验员ID |
| inspect_time | TIMESTAMP | 检验时间 |

### 5.13 PQC记录 (quality_pqc_record)

| 字段 | 类型 | 说明 |
|------|------|------|
| record_id | BIGINT | 记录ID |
| record_no | VARCHAR(50) | 记录编号 |
| workorder_id | BIGINT | 工单ID |
| workorder_no | VARCHAR(50) | 工单编号 |
| process_id | BIGINT | 工序ID |
| inspect_qty | DECIMAL(18,3) | 检验数量 |
| pass_qty | DECIMAL(18,3) | 合格数量 |
| fail_qty | DECIMAL(18,3) | 不合格数量 |
| result | VARCHAR(20) | 检验结果 |
| inspector_id | BIGINT | 检验员ID |
| inspect_time | TIMESTAMP | 检验时间 |

### 5.14 OQC记录 (quality_oqc_record)

| 字段 | 类型 | 说明 |
|------|------|------|
| record_id | BIGINT | 记录ID |
| record_no | VARCHAR(50) | 记录编号 |
| delivery_no | VARCHAR(50) | 出货单号 |
| inspect_qty | DECIMAL(18,3) | 检验数量 |
| pass_qty | DECIMAL(18,3) | 合格数量 |
| fail_qty | DECIMAL(18,3) | 不合格数量 |
| result | VARCHAR(20) | 检验结果 |
| inspector_id | BIGINT | 检验员ID |
| inspect_time | TIMESTAMP | 检验时间 |

### 5.15 首件检验 (quality_first_inspect)

| 字段 | 类型 | 说明 |
|------|------|------|
| record_id | BIGINT | 记录ID |
| record_no | VARCHAR(50) | 记录编号 |
| workorder_id | BIGINT | 工单ID |
| inspect_type | VARCHAR(20) | 检验类型(首件/末件) |
| inspect_qty | DECIMAL(18,3) | 检验数量 |
| pass_qty | DECIMAL(18,3) | 合格数量 |
| result | VARCHAR(20) | 检验结果 |
| confirm_status | VARCHAR(20) | 确认状态 |
| confirm_time | TIMESTAMP | 确认时间 |

### 5.16 巡检记录 (quality_patrol_inspect)

| 字段 | 类型 | 说明 |
|------|------|------|
| record_id | BIGINT | 记录ID |
| record_no | VARCHAR(50) | 记录编号 |
| plan_id | BIGINT | 巡检计划ID |
| patrol_type | VARCHAR(20) | 巡检类型 |
| inspector_id | BIGINT | 巡检员ID |
| inspect_time | TIMESTAMP | 巡检时间 |
| result | VARCHAR(20) | 巡检结果 |
| ok_count | INT | 正常项数 |
| ng_count | INT | 异常项数 |

### 5.17 质量通知 (quality_notice)

| 字段 | 类型 | 说明 |
|------|------|------|
| notice_id | BIGINT | 通知ID |
| notice_no | VARCHAR(50) | 通知编号 |
| title | VARCHAR(200) | 通知标题 |
| content | TEXT | 通知内容 |
| notice_type | VARCHAR(20) | 通知类型 |
| priority | VARCHAR(20) | 优先级 |
| status | VARCHAR(20) | 状态 |
| publish_time | TIMESTAMP | 发布时间 |

### 5.18 质量标准 (quality_standard)

| 字段 | 类型 | 说明 |
|------|------|------|
| std_id | BIGINT | 标准ID |
| std_code | VARCHAR(50) | 标准编码 |
| std_name | VARCHAR(200) | 标准名称 |
| std_type | VARCHAR(20) | 标准类型 |
| version | VARCHAR(20) | 版本 |
| content | TEXT | 标准内容 |
| status | VARCHAR(20) | 状态 |

### 5.19 质量目标 (quality_target)

| 字段 | 类型 | 说明 |
|------|------|------|
| target_id | BIGINT | 目标ID |
| target_code | VARCHAR(50) | 目标编码 |
| target_name | VARCHAR(200) | 目标名称 |
| target_value | DECIMAL(18,4) | 目标值 |
| current_value | DECIMAL(18,4) | 当前值 |
| measure_unit | VARCHAR(20) | 计量单位 |
| start_date | DATE | 开始日期 |
| end_date | DATE | 结束日期 |
| status | VARCHAR(20) | 状态 |

### 5.20 质量分析 (quality_analysis)

| 字段 | 类型 | 说明 |
|------|------|------|
| analysis_id | BIGINT | 分析ID |
| analysis_no | VARCHAR(50) | 分析编号 |
| analysis_type | VARCHAR(20) | 分析类型 |
| analysis_name | VARCHAR(200) | 分析名称 |
| date_range | VARCHAR(50) | 日期范围 |
| data_source | VARCHAR(100) | 数据来源 |
| result | TEXT | 分析结果 |

### 5.21 质量报表 (quality_report)

| 字段 | 类型 | 说明 |
|------|------|------|
| report_id | BIGINT | 报表ID |
| report_code | VARCHAR(50) | 报表编码 |
| report_name | VARCHAR(200) | 报表名称 |
| report_type | VARCHAR(20) | 报表类型 |
| config | TEXT | 报表配置 |
| generate_time | TIMESTAMP | 生成时间 |

### 5.22 退货管理 (quality_return)

| 字段 | 类型 | 说明 |
|------|------|------|
| return_id | BIGINT | 退货ID |
| return_no | VARCHAR(50) | 退货单号 |
| source_type | VARCHAR(20) | 来源类型 |
| source_no | VARCHAR(50) | 来源单号 |
| return_reason | VARCHAR(500) | 退货原因 |
| return_qty | DECIMAL(18,3) | 退货数量 |
| status | VARCHAR(20) | 状态 |
| applicant_id | BIGINT | 申请人ID |
| apply_time | TIMESTAMP | 申请时间 |

### 5.23 索赔管理 (quality_claim)

| 字段 | 类型 | 说明 |
|------|------|------|
| claim_id | BIGINT | 索赔ID |
| claim_no | VARCHAR(50) | 索赔单号 |
| supplier_id | BIGINT | 供应商ID |
| claim_type | VARCHAR(20) | 索赔类型 |
| claim_amount | DECIMAL(18,2) | 索赔金额 |
| claim_reason | VARCHAR(500) | 索赔原因 |
| status | VARCHAR(20) | 状态 |

### 5.24 供应商评审 (quality_supplier_audit)

| 字段 | 类型 | 说明 |
|------|------|------|
| audit_id | BIGINT | 评审ID |
| audit_no | VARCHAR(50) | 评审编号 |
| supplier_id | BIGINT | 供应商ID |
| supplier_name | VARCHAR(200) | 供应商名称 |
| audit_type | VARCHAR(20) | 评审类型 |
| plan_date | DATE | 计划日期 |
| audit_result | VARCHAR(20) | 评审结果 |
| auditor_id | BIGINT | 评审员ID |
| audit_time | TIMESTAMP | 评审时间 |

### 5.25 纠正措施 (quality_corrective)

| 字段 | 类型 | 说明 |
|------|------|------|
| corrective_id | BIGINT | 纠正ID |
| corrective_no | VARCHAR(50) | 纠正编号 |
| source_type | VARCHAR(20) | 来源类型 |
| source_no | VARCHAR(50) | 来源单号 |
| problem_desc | TEXT | 问题描述 |
| root_cause | TEXT | 根本原因 |
| corrective_action | TEXT | 纠正措施 |
| responsible_id | BIGINT | 负责人ID |
| due_date | DATE | 截止日期 |
| status | VARCHAR(20) | 状态 |

### 5.26 预防措施 (quality_preventive)

| 字段 | 类型 | 说明 |
|------|------|------|
| preventive_id | BIGINT | 预防ID |
| preventive_no | VARCHAR(50) | 预防编号 |
| risk_desc | TEXT | 风险描述 |
| preventive_action | TEXT | 预防措施 |
| responsible_id | BIGINT | 负责人ID |
| due_date | DATE | 截止日期 |
| status | VARCHAR(20) | 状态 |

### 5.27 质量培训 (quality_training)

| 字段 | 类型 | 说明 |
|------|------|------|
| training_id | BIGINT | 培训ID |
| training_no | VARCHAR(50) | 培训编号 |
| training_title | VARCHAR(200) | 培训标题 |
| training_content | TEXT | 培训内容 |
| trainer_id | BIGINT | 培训师ID |
| training_date | DATE | 培训日期 |
| training_hours | DECIMAL(10,2) | 培训时长 |
| attendee_count | INT | 参加人数 |
| status | VARCHAR(20) | 状态 |

### 5.28 质量审核 (quality_audit)

| 字段 | 类型 | 说明 |
|------|------|------|
| audit_id | BIGINT | 审核ID |
| audit_no | VARCHAR(50) | 审核编号 |
| audit_type | VARCHAR(20) | 审核类型 |
| audit_scope | VARCHAR(200) | 审核范围 |
| auditor_id | BIGINT | 审核员ID |
| plan_date | DATE | 计划日期 |
| audit_result | VARCHAR(20) | 审核结果 |
| finding_count | INT | 不符合项数 |

### 5.29 质量证书 (quality_certificate)

| 字段 | 类型 | 说明 |
|------|------|------|
| cert_id | BIGINT | 证书ID |
| cert_no | VARCHAR(50) | 证书编号 |
| cert_name | VARCHAR(200) | 证书名称 |
| cert_type | VARCHAR(20) | 证书类型 |
| holder_id | BIGINT | 持有人ID |
| issue_date | DATE | 发证日期 |
| expiry_date | DATE | 到期日期 |
| status | VARCHAR(20) | 状态 |

### 5.30 计量管理 (quality_measure)

| 字段 | 类型 | 说明 |
|------|------|------|
| measure_id | BIGINT | 器具ID |
| measure_code | VARCHAR(50) | 器具编码 |
| measure_name | VARCHAR(200) | 器具名称 |
| measure_type | VARCHAR(20) | 器具类型 |
| measure_range | VARCHAR(100) | 测量范围 |
| precision_level | VARCHAR(20) | 精度等级 |
| calibration_cycle | INT | 校准周期(天) |
| last_cal_date | DATE | 上次检定日期 |
| next_cal_date | DATE | 下次检定日期 |
| status | VARCHAR(20) | 状态 |

### 5.31 校准记录 (quality_calibration)

| 字段 | 类型 | 说明 |
|------|------|------|
| calibration_id | BIGINT | 校准ID |
| calibration_no | VARCHAR(50) | 校准编号 |
| measure_id | BIGINT | 器具ID |
| calibration_date | DATE | 校准日期 |
| calibration_result | VARCHAR(20) | 校准结果 |
| calibrator_id | BIGINT | 校准员ID |
| calibrator_name | VARCHAR(100) | 校准员姓名 |
| deviation | DECIMAL(18,6) | 偏差值 |
| conclusion | TEXT | 校准结论 |

### 5.32 质量追溯 (quality_trace)

| 字段 | 类型 | 说明 |
|------|------|------|
| trace_id | BIGINT | 追溯ID |
| trace_no | VARCHAR(50) | 追溯编号 |
| batch_no | VARCHAR(50) | 批次号 |
| product_id | BIGINT | 产品ID |
| product_code | VARCHAR(50) | 产品编码 |
| trace_type | VARCHAR(20) | 追溯类型 |
| trace_data | TEXT | 追溯数据(JSON) |
| trace_time | TIMESTAMP | 追溯时间 |

---

## 6. API接口

### 6.1 计数器 (/quality/counter)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/counter/list | 计数器列表 |
| GET | /quality/counter/:id | 计数器详情 |
| POST | /quality/counter | 创建计数器 |
| PUT | /quality/counter/:id | 更新计数器 |
| DELETE | /quality/counter/:id | 删除计数器 |
| POST | /quality/counter/increment | 计数增加 |
| POST | /quality/counter/reset | 计数重置 |

### 6.2 检验方案 (/quality/inspectScheme)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/inspectScheme/list | 方案列表 |
| GET | /quality/inspectScheme/:id | 方案详情 |
| POST | /quality/inspectScheme | 创建方案 |
| PUT | /quality/inspectScheme/:id | 更新方案 |
| DELETE | /quality/inspectScheme/:id | 删除方案 |
| POST | /quality/inspectScheme/copy | 方案复制 |
| GET | /quality/inspectScheme/template | 方案模板 |

### 6.3 检验方法 (/quality/inspectMethod)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/inspectMethod/list | 方法列表 |
| GET | /quality/inspectMethod/:id | 方法详情 |
| POST | /quality/inspectMethod | 创建方法 |
| PUT | /quality/inspectMethod/:id | 更新方法 |
| DELETE | /quality/inspectMethod/:id | 删除方法 |
| GET | /quality/inspectMethod/library | 方法库 |

### 6.4 检验阶段 (/quality/inspectStage)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/inspectStage/list | 阶段列表 |
| GET | /quality/inspectStage/:id | 阶段详情 |
| POST | /quality/inspectStage | 创建阶段 |
| PUT | /quality/inspectStage/:id | 更新阶段 |
| DELETE | /quality/inspectStage/:id | 删除阶段 |

### 6.5 定性特性 (/quality/qualitativeFeature)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/qualitativeFeature/list | 特性列表 |
| GET | /quality/qualitativeFeature/:id | 特性详情 |
| POST | /quality/qualitativeFeature | 创建特性 |
| PUT | /quality/qualitativeFeature/:id | 更新特性 |
| DELETE | /quality/qualitativeFeature/:id | 删除特性 |
| GET | /quality/qualitativeFeature/template | 特性模板 |

### 6.6 定量特性 (/quality/quantitativeFeature)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/quantitativeFeature/list | 特性列表 |
| GET | /quality/quantitativeFeature/:id | 特性详情 |
| POST | /quality/quantitativeFeature | 创建特性 |
| PUT | /quality/quantitativeFeature/:id | 更新特性 |
| DELETE | /quality/quantitativeFeature/:id | 删除特性 |
| GET | /quality/quantitativeFeature/range | 特性范围 |

### 6.7 抽样方案 (/quality/samplingScheme)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/samplingScheme/list | 方案列表 |
| GET | /quality/samplingScheme/:id | 方案详情 |
| POST | /quality/samplingScheme | 创建方案 |
| PUT | /quality/samplingScheme/:id | 更新方案 |
| DELETE | /quality/samplingScheme/:id | 删除方案 |
| GET | /quality/samplingScheme/rule | 抽样规则 |

### 6.8 抽样过程 (/quality/samplingProcess)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/samplingProcess/list | 过程列表 |
| GET | /quality/samplingProcess/:id | 过程详情 |
| POST | /quality/samplingProcess | 创建记录 |
| PUT | /quality/samplingProcess/:id | 更新记录 |
| DELETE | /quality/samplingProcess/:id | 删除记录 |
| POST | /quality/samplingProcess/execute | 执行抽样 |

### 6.9 抽样代码 (/quality/samplingCode)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/samplingCode/list | 代码列表 |
| GET | /quality/samplingCode/:id | 代码详情 |
| POST | /quality/samplingCode | 创建代码规则 |
| PUT | /quality/samplingCode/:id | 更新代码规则 |
| DELETE | /quality/samplingCode/:id | 删除代码规则 |
| GET | /quality/samplingCode/generate | 生成代码 |

### 6.10 检验请求 (/quality/inspectRequest)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/inspectRequest/list | 请求列表 |
| GET | /quality/inspectRequest/:id | 请求详情 |
| POST | /quality/inspectRequest | 创建请求 |
| PUT | /quality/inspectRequest/:id | 更新请求 |
| DELETE | /quality/inspectRequest/:id | 删除请求 |
| POST | /quality/inspectRequest/approve | 审批请求 |
| POST | /quality/inspectRequest/reject | 拒绝请求 |

### 6.11 检验任务 (/quality/inspectTask)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/inspectTask/list | 任务列表 |
| GET | /quality/inspectTask/:id | 任务详情 |
| POST | /quality/inspectTask/assign | 分配任务 |
| PUT | /quality/inspectTask/:id/execute | 执行任务 |
| PUT | /quality/inspectTask/:id/complete | 完成任务 |

### 6.12 IQC记录 (/quality/iqcRecord)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/iqcRecord/list | 记录列表 |
| GET | /quality/iqcRecord/:id | 记录详情 |
| POST | /quality/iqcRecord | 创建记录 |
| PUT | /quality/iqcRecord/:id | 更新记录 |
| GET | /quality/iqcRecord/export | 导出记录 |

### 6.13 PQC记录 (/quality/pqcRecord)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/pqcRecord/list | 记录列表 |
| GET | /quality/pqcRecord/:id | 记录详情 |
| POST | /quality/pqcRecord | 创建记录 |
| PUT | /quality/pqcRecord/:id | 更新记录 |
| GET | /quality/pqcRecord/export | 导出记录 |

### 6.14 OQC记录 (/quality/oqcRecord)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/oqcRecord/list | 记录列表 |
| GET | /quality/oqcRecord/:id | 记录详情 |
| POST | /quality/oqcRecord | 创建记录 |
| PUT | /quality/oqcRecord/:id | 更新记录 |
| GET | /quality/oqcRecord/export | 导出记录 |

### 6.15 首件检验 (/quality/firstInspect)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/firstInspect/list | 记录列表 |
| GET | /quality/firstInspect/:id | 记录详情 |
| POST | /quality/firstInspect | 创建记录 |
| PUT | /quality/firstInspect/:id | 更新记录 |
| POST | /quality/firstInspect/:id/confirm | 确认首件 |

### 6.16 巡检记录 (/quality/patrolInspect)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/patrolInspect/list | 记录列表 |
| GET | /quality/patrolInspect/:id | 记录详情 |
| POST | /quality/patrolInspect | 创建记录 |
| PUT | /quality/patrolInspect/:id | 更新记录 |
| GET | /quality/patrolInspect/plan | 巡检计划 |

### 6.17 质量通知 (/quality/qualityNotice)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/qualityNotice/list | 通知列表 |
| GET | /quality/qualityNotice/:id | 通知详情 |
| POST | /quality/qualityNotice | 创建通知 |
| PUT | /quality/qualityNotice/:id | 更新通知 |
| DELETE | /quality/qualityNotice/:id | 删除通知 |
| POST | /quality/qualityNotice/publish | 发布通知 |
| POST | /quality/qualityNotice/distribute | 分发通知 |

### 6.18 质量标准 (/quality/std)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/std/list | 标准列表 |
| GET | /quality/std/:id | 标准详情 |
| POST | /quality/std | 创建标准 |
| PUT | /quality/std/:id | 更新标准 |
| DELETE | /quality/std/:id | 删除标准 |
| GET | /quality/std/library | 标准库 |

### 6.19 质量目标 (/quality/target)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/target/list | 目标列表 |
| GET | /quality/target/:id | 目标详情 |
| POST | /quality/target | 创建目标 |
| PUT | /quality/target/:id | 更新目标 |
| DELETE | /quality/target/:id | 删除目标 |
| GET | /quality/target/track | 目标跟踪 |

### 6.20 质量分析 (/quality/analysis)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/analysis/list | 分析列表 |
| GET | /quality/analysis/:id | 分析详情 |
| POST | /quality/analysis | 创建分析 |
| GET | /quality/analysis/report | 分析报表 |
| GET | /quality/analysis/trend | 趋势分析 |

### 6.21 质量报表 (/quality/report)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/report/list | 报表列表 |
| GET | /quality/report/:id | 报表详情 |
| POST | /quality/report | 创建报表 |
| PUT | /quality/report/:id | 更新报表 |
| DELETE | /quality/report/:id | 删除报表 |
| POST | /quality/report/generate | 生成报表 |

### 6.22 退货管理 (/quality/return)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/return/list | 退货列表 |
| GET | /quality/return/:id | 退货详情 |
| POST | /quality/return | 创建退货 |
| PUT | /quality/return/:id | 更新退货 |
| DELETE | /quality/return/:id | 删除退货 |
| POST | /quality/return/:id/process | 处理退货 |

### 6.23 索赔管理 (/quality/claim)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/claim/list | 索赔列表 |
| GET | /quality/claim/:id | 索赔详情 |
| POST | /quality/claim | 创建索赔 |
| PUT | /quality/claim/:id | 更新索赔 |
| DELETE | /quality/claim/:id | 删除索赔 |
| POST | /quality/claim/:id/settle | 处理索赔 |

### 6.24 供应商评审 (/quality/supplierAudit)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/supplierAudit/list | 评审列表 |
| GET | /quality/supplierAudit/:id | 评审详情 |
| POST | /quality/supplierAudit | 创建评审 |
| PUT | /quality/supplierAudit/:id | 更新评审 |
| DELETE | /quality/supplierAudit/:id | 删除评审 |
| GET | /quality/supplierAudit/record | 评审记录 |

### 6.25 纠正措施 (/quality/corrective)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/corrective/list | 纠正列表 |
| GET | /quality/corrective/:id | 纠正详情 |
| POST | /quality/corrective | 创建纠正 |
| PUT | /quality/corrective/:id | 更新纠正 |
| DELETE | /quality/corrective/:id | 删除纠正 |
| POST | /quality/corrective/:id/track | 跟踪纠正 |

### 6.26 预防措施 (/quality/preventive)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/preventive/list | 预防列表 |
| GET | /quality/preventive/:id | 预防详情 |
| POST | /quality/preventive | 创建预防 |
| PUT | /quality/preventive/:id | 更新预防 |
| DELETE | /quality/preventive/:id | 删除预防 |
| POST | /quality/preventive/:id/track | 跟踪预防 |

### 6.27 质量培训 (/quality/training)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/training/list | 培训列表 |
| GET | /quality/training/:id | 培训详情 |
| POST | /quality/training | 创建培训 |
| PUT | /quality/training/:id | 更新培训 |
| DELETE | /quality/training/:id | 删除培训 |
| GET | /quality/training/record | 培训记录 |

### 6.28 质量审核 (/quality/audit)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/audit/list | 审核列表 |
| GET | /quality/audit/:id | 审核详情 |
| POST | /quality/audit | 创建审核 |
| PUT | /quality/audit/:id | 更新审核 |
| DELETE | /quality/audit/:id | 删除审核 |
| POST | /quality/audit/:id/execute | 执行审核 |

### 6.29 质量证书 (/quality/certificate)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/certificate/list | 证书列表 |
| GET | /quality/certificate/:id | 证书详情 |
| POST | /quality/certificate | 创建证书 |
| PUT | /quality/certificate/:id | 更新证书 |
| DELETE | /quality/certificate/:id | 删除证书 |
| GET | /quality/certificate/expiry | 证书到期提醒 |

### 6.30 计量管理 (/quality/measure)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/measure/list | 器具列表 |
| GET | /quality/measure/:id | 器具详情 |
| POST | /quality/measure | 创建器具 |
| PUT | /quality/measure/:id | 更新器具 |
| DELETE | /quality/measure/:id | 删除器具 |
| GET | /quality/measure/plan | 检定计划 |

### 6.31 校准记录 (/quality/calibration)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/calibration/list | 记录列表 |
| GET | /quality/calibration/:id | 记录详情 |
| POST | /quality/calibration | 创建记录 |
| PUT | /quality/calibration/:id | 更新记录 |
| DELETE | /quality/calibration/:id | 删除记录 |
| GET | /quality/calibration/history | 校准历史 |

### 6.32 质量追溯 (/quality/trace)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/trace/list | 追溯列表 |
| GET | /quality/trace/:id | 追溯详情 |
| GET | /quality/trace/query | 追溯查询 |
| GET | /quality/trace/batch/:batchNo | 按批次追溯 |
| GET | /quality/trace/detail | 追溯明细 |

---

## 7. 关联文档

- [MOM3.0_主设计文档](./MOM3.0_主设计文档.md) - 系统总览
- [MOM3.0_UI设计规范](./MOM3.0_UI设计规范.md) - UI规范详情
- [MOM3.0_质量模块设计文档](./MOM3.0_质量模块设计文档.md) - 后端设计详情
