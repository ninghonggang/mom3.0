# MOM3.0 质量模块设计文档

**版本**: V2.0 | **所属模块**: M05质量管理 | **基于**: [MOM3.0_主设计文档](./MOM3.0_主设计文档.md)

---

## 1. 模块概述

### 1.1 功能定位

质量管理模块覆盖来料、过程、成品、出货全流程质量检验，配合NCR处理、QRCI改进、LPA分层审核等质量工具，实现全面的质量管控与追溯。

### 1.2 核心功能

| 功能 | 说明 |
|------|------|
| IQC来料检验 | 采购物料入库质量检验 |
| IPQC过程检验 | 生产过程质量巡检 |
| FQC成品检验 | 完工产品质量检验 |
| OQC出货检验 | 出货前质量确认 |
| 缺陷代码管理 | 不良缺陷定义与分类 |
| 缺陷记录 | 不良品记录管理 |
| NCR处理 | 不良品处理流程 |
| SPC统计 | 统计过程控制 |
| QRCI质量闭环 | 质量改进闭环管理 |
| LPA分层审核 | 分层审核管理 |
| 首末件检验 | 首件/末件质量确认 |

---

## 2. 页面清单

| 页面 | 路由路径 | 核心功能 |
|------|----------|----------|
| IQC来料检验 | `/quality/iqc` | 来料检验、判定、合格入库 |
| IPQC过程检验 | `/quality/ipqc` | 过程巡检、异常记录 |
| FQC成品检验 | `/quality/fqc` | 成品检验、出库确认 |
| OQC出货检验 | `/quality/oqc` | 出货检验、装箱确认 |
| 缺陷代码 | `/quality/defect-code` | 缺陷分类定义 |
| 缺陷记录 | `/quality/defect-record` | 不良记录管理 |
| NCR处理 | `/quality/ncr` | 不良品处理流程 |
| SPC控制图 | `/quality/spc` | 统计过程控制图 |
| QRCI管理 | `/quality/qrci` | 质量改进闭环 |
| LPA审核 | `/quality/lpa` | 分层审核管理 |
| 首末件检验 | `/quality/first-last` | 首件/末件检验 |

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

**NCR状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| OPEN | warning | 待处理 |
| INVESTIGATING | primary | 处理中 |
| RESOLVED | success | 已解决 |
| CLOSED | info | 已关闭 |
| CANCELLED | info | 已取消 |

**QRCI状态**

| 状态值 | 标签类型 | 显示文本 |
|--------|----------|----------|
| OPEN | warning | 待改进 |
| IN_PROGRESS | primary | 改进中 |
| VERIFIED | success | 已验证 |
| CLOSED | info | 已关闭 |

### 3.3 NCR处理流程

```vue
<template #footer>
  <el-button @click="detailVisible = false">关闭</el-button>
  <el-button type="primary" v-if="detailData.status === 'OPEN'" @click="handleAssign">
    指派处理人
  </el-button>
  <el-button type="success" v-if="detailData.status === 'INVESTIGATING'" @click="handleResolve">
    标记已解决
  </el-button>
  <el-button type="danger" v-if="['OPEN', 'INVESTIGATING'].includes(detailData.status)" @click="handleCancel">
    取消
  </el-button>
</template>
```

---

## 4. 业务流程

### 4.1 来料检验流程

```
采购收货 → IQC检验 → 合格品入库 / 不良品NCR处理
```

### 4.2 过程检验流程

```
工单开工 → IPQC巡检 → 异常记录 → NCR处理 → 继续生产
```

### 4.3 NCR处理流程

```
NCR创建 → 原因分析(5Why) → 纠正措施 → 预防措施 → 验证关闭
```

---

## 5. 数据模型

### 5.1 检验单

| 字段 | 类型 | 说明 |
|------|------|------|
| inspect_no | VARCHAR(50) | 检验单号 |
| inspect_type | VARCHAR(20) | IQC/IPQC/FQC/OQC |
| source_type | VARCHAR(20) | 来源类型 |
| source_id | BIGINT | 来源ID |
| material_id | BIGINT | 物料ID |
| batch_no | VARCHAR(50) | 批次号 |
| inspect_qty | DECIMAL(18,3) | 检验数量 |
| pass_qty | DECIMAL(18,3) | 合格数量 |
| fail_qty | DECIMAL(18,3) | 不合格数量 |
| result | VARCHAR(20) | 检验结果 |
| inspector_id | BIGINT | 检验员ID |
| inspect_time | TIMESTAMP | 检验时间 |
| status | VARCHAR(20) | 状态 |

### 5.2 NCR

| 字段 | 类型 | 说明 |
|------|------|------|
| ncr_no | VARCHAR(50) | NCR编号 |
| source_type | VARCHAR(20) | 来源类型 |
| source_no | VARCHAR(50) | 来源单号 |
| defect_desc | VARCHAR(500) | 缺陷描述 |
| defect_level | VARCHAR(20) | 缺陷等级 |
| quantity | DECIMAL(18,3) | 数量 |
| status | VARCHAR(20) | 状态 |
| root_cause | TEXT | 根本原因 |
| corrective_action | TEXT | 纠正措施 |
| preventive_action | TEXT | 预防措施 |
| assignee_id | BIGINT | 处理人ID |

---

## 6. API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /quality/iqc/list | IQC列表 |
| POST | /quality/iqc | 创建IQC检验单 |
| PUT | /quality/iqc/:id/inspect | 检验判定 |
| GET | /quality/ipqc/list | IPQC列表 |
| GET | /quality/fqc/list | FQC列表 |
| GET | /quality/oqc/list | OQC列表 |
| GET | /quality/ncr/list | NCR列表 |
| POST | /quality/ncr | 创建NCR |
| PUT | /quality/ncr/:id/resolve | NCR解决 |
| GET | /quality/spc/list | SPC数据列表 |
| GET | /quality/spc/capability/:configId | CP/CPK分析 |

---

## 7. 关联文档

- [MOM3.0_主设计文档](./MOM3.0_主设计文档.md) - 系统总览
- [MOM3.0_UI设计规范](./MOM3.0_UI设计规范.md) - UI规范详情

---

## 8. AQL与抽样方案表结构

### 8.1 qms_aql - AQL标准配置表

```sql
CREATE TABLE IF NOT EXISTS qms_aql (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    code VARCHAR(50) NOT NULL COMMENT 'AQL代码',
    name VARCHAR(200) NOT NULL COMMENT 'AQL名称',
    level VARCHAR(20) NOT NULL COMMENT '检验水平(I/II/III/S-1/S-2/S-3/S-4)',
    value DECIMAL(10,4) NOT NULL COMMENT 'AQL值',
    sampling_amount INT NOT NULL COMMENT '抽样数量',
    status VARCHAR(20) DEFAULT 'ACTIVE' COMMENT '状态(ACTIVE/INACTIVE)',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_aql_code ON qms_aql(tenant_id, code) WHERE deleted_at IS NULL;
```

### 8.2 qms_aql_detail - AQL明细表

```sql
CREATE TABLE IF NOT EXISTS qms_aql_detail (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    aql_id BIGINT NOT NULL COMMENT 'AQL标准ID',
    lot_size_min DECIMAL(18,0) NOT NULL COMMENT '批量范围下限',
    lot_size_max DECIMAL(18,0) NOT NULL COMMENT '批量范围上限',
    sampling_type VARCHAR(20) NOT NULL COMMENT '抽样类型(NORMAL/TIGHTENED/RELAXED)',
    sampling_amount INT NOT NULL COMMENT '抽样数量',
    acceptance_num INT NOT NULL COMMENT '接收数',
    rejection_num INT NOT NULL COMMENT '拒收数',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_aql_detail_aql_id ON qms_aql_detail(aql_id) WHERE deleted_at IS NULL;
```

### 8.3 qms_material_aql - 物料AQL配置表

```sql
CREATE TABLE IF NOT EXISTS qms_material_aql (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    material_id BIGINT NOT NULL COMMENT '物料ID',
    material_code VARCHAR(50) NOT NULL COMMENT '物料编码',
    material_name VARCHAR(200) COMMENT '物料名称',
    inspection_level VARCHAR(20) NOT NULL COMMENT '检验水平',
    aql_id BIGINT NOT NULL COMMENT 'AQL标准ID',
    aql_code VARCHAR(50) NOT NULL COMMENT 'AQL代码',
    aql_value DECIMAL(10,4) NOT NULL COMMENT 'AQL值',
    apply_type VARCHAR(20) DEFAULT 'ALL' COMMENT '适用类型(ALL/IQC/PQC/OQC)',
    status VARCHAR(20) DEFAULT 'ACTIVE' COMMENT '状态',
    effective_date DATE COMMENT '生效日期',
    expiry_date DATE COMMENT '失效日期',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_material_aql_material ON qms_material_aql(material_id) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_material_type ON qms_material_aql(tenant_id, material_id, apply_type) WHERE deleted_at IS NULL;
```

### 8.4 qms_sampling_plan - 抽样计划表

```sql
CREATE TABLE IF NOT EXISTS qms_sampling_plan (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    code VARCHAR(50) NOT NULL COMMENT '抽样计划代码',
    name VARCHAR(200) NOT NULL COMMENT '抽样计划名称',
    plan_type VARCHAR(20) NOT NULL COMMENT '计划类型(GB/T2828.1/ISO2859/ANSI/自定义)',
    inspection_level VARCHAR(20) NOT NULL COMMENT '检验水平',
    sampling_type VARCHAR(20) NOT NULL COMMENT '抽样方案类型(NORMAL/SINGLE/DOUBLE/MULTIPLE)',
    sample_size_code VARCHAR(20) COMMENT '样本量字码',
    sample_size INT COMMENT '样本量',
    acceptance_num INT COMMENT '接收数Ac',
    rejection_num INT COMMENT '拒收数Re',
    status VARCHAR(20) DEFAULT 'ACTIVE' COMMENT '状态',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_sampling_plan_code ON qms_sampling_plan(tenant_id, code) WHERE deleted_at IS NULL;
```

### 8.5 qms_sampling_rule - 抽样规则表

```sql
CREATE TABLE IF NOT EXISTS qms_sampling_rule (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    code VARCHAR(50) NOT NULL COMMENT '规则代码',
    name VARCHAR(200) NOT NULL COMMENT '规则名称',
    plan_id BIGINT NOT NULL COMMENT '抽样计划ID',
    trigger_condition VARCHAR(500) NOT NULL COMMENT '触发条件(JSON)',
    sampling_action VARCHAR(100) NOT NULL COMMENT '抽样动作',
    priority INT DEFAULT 0 COMMENT '优先级',
    status VARCHAR(20) DEFAULT 'ACTIVE' COMMENT '状态',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_sampling_rule_plan ON qms_sampling_rule(plan_id) WHERE deleted_at IS NULL;
```

### 8.6 qms_sampling_record - 抽样记录表

```sql
CREATE TABLE IF NOT EXISTS qms_sampling_record (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    record_no VARCHAR(50) NOT NULL COMMENT '记录编号',
    source_type VARCHAR(20) NOT NULL COMMENT '来源类型(IQC/PQC/OQC)',
    source_no VARCHAR(50) NOT NULL COMMENT '来源单号',
    lot_no VARCHAR(50) NOT NULL COMMENT '批次号',
    lot_size INT NOT NULL COMMENT '批量大小',
    sample_size INT NOT NULL COMMENT '抽样数量',
    sample_code VARCHAR(50) COMMENT '样本代码',
    sampling_time TIMESTAMP NOT NULL COMMENT '抽样时间',
    sampler_id BIGINT COMMENT '抽样人ID',
    sampler_name VARCHAR(100) COMMENT '抽样人姓名',
    sampling_location VARCHAR(200) COMMENT '抽样地点',
    result VARCHAR(20) COMMENT '抽样结果(PASS/FAIL)',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_sampling_record_source ON qms_sampling_record(source_type, source_no) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_record_no ON qms_sampling_record(tenant_id, record_no) WHERE deleted_at IS NULL;
```

### 8.7 qms_dynamic_rule - 动态检验规则表

```sql
CREATE TABLE IF NOT EXISTS qms_dynamic_rule (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    code VARCHAR(50) NOT NULL COMMENT '规则代码',
    name VARCHAR(200) NOT NULL COMMENT '规则名称',
    rule_type VARCHAR(20) NOT NULL COMMENT '规则类型(INSPECTION_LEVEL/SAMPLING/AQL)',
    condition_expression TEXT NOT NULL COMMENT '条件表达式',
    action_type VARCHAR(50) NOT NULL COMMENT '动作类型',
    action_value VARCHAR(200) COMMENT '动作值',
    priority INT DEFAULT 0 COMMENT '优先级',
    effective_date DATE COMMENT '生效日期',
    expiry_date DATE COMMENT '失效日期',
    status VARCHAR(20) DEFAULT 'ACTIVE' COMMENT '状态',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_dynamic_rule_code ON qms_dynamic_rule(tenant_id, code) WHERE deleted_at IS NULL;
```

### 8.8 qms_sample_code - 样本编码表

```sql
CREATE TABLE IF NOT EXISTS qms_sample_code (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    code VARCHAR(50) NOT NULL COMMENT '样本编码',
    code_rule VARCHAR(200) NOT NULL COMMENT '编码规则',
    current_value BIGINT NOT NULL DEFAULT 0 COMMENT '当前值',
    prefix VARCHAR(20) NOT NULL COMMENT '前缀',
    suffix VARCHAR(20) COMMENT '后缀',
    code_length INT DEFAULT 8 COMMENT '编码长度',
    reset_type VARCHAR(20) DEFAULT 'DAILY' COMMENT '重置类型(DAILY/MONTHLY/YEARLY/NONE)',
    last_reset_date DATE COMMENT '最后重置日期',
    status VARCHAR(20) DEFAULT 'ACTIVE' COMMENT '状态',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_sample_code ON qms_sample_code(tenant_id, code) WHERE deleted_at IS NULL;
```

---

## 9. 检验特性管理表结构

### 9.1 qms_inspection_characteristics - 检验特性表

```sql
CREATE TABLE IF NOT EXISTS qms_inspection_characteristics (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    code VARCHAR(50) NOT NULL COMMENT '特性编码',
    name VARCHAR(200) NOT NULL COMMENT '特性名称',
    characteristics_type VARCHAR(20) NOT NULL COMMENT '特性类型(QUANTITATIVE/QUALITATIVE)',
    inspection_method_id BIGINT COMMENT '检验方法ID',
    inspection_method_code VARCHAR(50) COMMENT '检验方法代码',
    inspection_method_name VARCHAR(200) COMMENT '检验方法名称',
    inspection_tool VARCHAR(100) COMMENT '检验工具',
    qualification_type VARCHAR(20) COMMENT '合格类型(PASS_FAIL/RANGE/VALUE)',
    usl DECIMAL(18,6) COMMENT '上规格限',
    lsl DECIMAL(18,6) COMMENT '下规格限',
    target_value DECIMAL(18,6) COMMENT '目标值',
    unit VARCHAR(20) COMMENT '计量单位',
    critical_defect DECIMAL(10,4) COMMENT '严重缺陷率',
    major_defect DECIMAL(10,4) COMMENT '主要缺陷率',
    minor_defect DECIMAL(10,4) COMMENT '轻微缺陷率',
    sampling_type VARCHAR(20) DEFAULT 'NORMAL' COMMENT '抽样类型',
    status VARCHAR(20) DEFAULT 'ACTIVE' COMMENT '状态',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_characteristics_code ON qms_inspection_characteristics(tenant_id, code) WHERE deleted_at IS NULL;
```

### 9.2 qms_char_quantitative - 定量特性表

```sql
CREATE TABLE IF NOT EXISTS qms_char_quantitative (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    characteristics_id BIGINT NOT NULL COMMENT '检验特性ID',
    characteristics_code VARCHAR(50) NOT NULL COMMENT '特性编码',
    usl DECIMAL(18,6) COMMENT '上规格限',
    lsl DECIMAL(18,6) COMMENT '下规格限',
    target DECIMAL(18,6) COMMENT '目标值',
    tolerance DECIMAL(18,6) COMMENT '公差',
    control_ucl DECIMAL(18,6) COMMENT '控制上限',
    control_lcl DECIMAL(18,6) COMMENT '控制下限',
    sample_mean DECIMAL(18,6) COMMENT '样本均值',
    sample_std DECIMAL(18,6) COMMENT '样本标准差',
    cpk DECIMAL(10,4) COMMENT '过程能力指数',
    measuring_range VARCHAR(100) COMMENT '测量范围',
    precision_level VARCHAR(20) COMMENT '精度等级',
    calibration_cycle INT COMMENT '校准周期(天)',
    status VARCHAR(20) DEFAULT 'ACTIVE' COMMENT '状态',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_quantitative_char ON qms_char_quantitative(characteristics_id) WHERE deleted_at IS NULL;
```

### 9.3 qms_char_qualitative - 定性特性表

```sql
CREATE TABLE IF NOT EXISTS qms_char_qualitative (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    characteristics_id BIGINT NOT NULL COMMENT '检验特性ID',
    characteristics_code VARCHAR(50) NOT NULL COMMENT '特性编码',
    standard VARCHAR(500) NOT NULL COMMENT '合格标准',
    defect_classification VARCHAR(20) COMMENT '缺陷分类(CR/MA/MI)',
    inspection_tool VARCHAR(100) COMMENT '检验工具',
    reference_sample VARCHAR(200) COMMENT '参考样本',
    judgement_criteria TEXT COMMENT '判定标准详情',
    status VARCHAR(20) DEFAULT 'ACTIVE' COMMENT '状态',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_qualitative_char ON qms_char_qualitative(characteristics_id) WHERE deleted_at IS NULL;
```

### 9.4 qms_inspection_method - 检验方法表

```sql
CREATE TABLE IF NOT EXISTS qms_inspection_method (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    code VARCHAR(50) NOT NULL COMMENT '方法编码',
    name VARCHAR(200) NOT NULL COMMENT '方法名称',
    type VARCHAR(20) NOT NULL COMMENT '方法类型(VISUAL/MEASURING/TESTING)',
    standard VARCHAR(50) COMMENT '执行标准',
    description TEXT COMMENT '方法描述',
    procedure TEXT COMMENT '操作步骤',
    equipment_required VARCHAR(500) COMMENT '所需设备',
    environment_required VARCHAR(200) COMMENT '环境要求',
    duration INT COMMENT '检验时长(分钟)',
    status VARCHAR(20) DEFAULT 'ACTIVE' COMMENT '状态',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_method_code ON qms_inspection_method(tenant_id, code) WHERE deleted_at IS NULL;
```

### 9.5 qms_inspection_scheme - 检验方案表

```sql
CREATE TABLE IF NOT EXISTS qms_inspection_scheme (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    code VARCHAR(50) NOT NULL COMMENT '方案编码',
    name VARCHAR(200) NOT NULL COMMENT '方案名称',
    scheme_type VARCHAR(20) NOT NULL COMMENT '方案类型(IQC/PQC/FQC/OQC)',
    product_code VARCHAR(50) COMMENT '产品编码',
    product_name VARCHAR(200) COMMENT '产品名称',
    aql_id BIGINT COMMENT 'AQL标准ID',
    aql_code VARCHAR(50) COMMENT 'AQL代码',
    aql_value DECIMAL(10,4) COMMENT 'AQL值',
    inspection_level VARCHAR(20) NOT NULL COMMENT '检验水平',
    sampling_plan_id BIGINT COMMENT '抽样计划ID',
    sampling_type VARCHAR(20) DEFAULT 'NORMAL' COMMENT '抽样类型',
    critical_limit INT DEFAULT 0 COMMENT '严重缺陷限',
    major_limit INT DEFAULT 0 COMMENT '主要缺陷限',
    minor_limit INT DEFAULT 0 COMMENT '轻微缺陷限',
    status VARCHAR(20) DEFAULT 'DRAFT' COMMENT '状态(DRAFT/ACTIVE/INACTIVE)',
    effective_date DATE COMMENT '生效日期',
    expiry_date DATE COMMENT '失效日期',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_scheme_code ON qms_inspection_scheme(tenant_id, code) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_scheme_product ON qms_inspection_scheme(product_code) WHERE deleted_at IS NULL;
```

### 9.6 qms_inspection_scheme_detail - 检验方案明细表

```sql
CREATE TABLE IF NOT EXISTS qms_inspection_scheme_detail (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    scheme_id BIGINT NOT NULL COMMENT '检验方案ID',
    characteristics_id BIGINT NOT NULL COMMENT '检验特性ID',
    characteristics_code VARCHAR(50) NOT NULL COMMENT '特性编码',
    characteristics_name VARCHAR(200) NOT NULL COMMENT '特性名称',
    characteristics_type VARCHAR(20) NOT NULL COMMENT '特性类型',
    sequence INT DEFAULT 0 COMMENT '顺序号',
    is_mandatory INT DEFAULT 1 COMMENT '是否必检(0否/1是)',
    sampling_ratio DECIMAL(10,4) COMMENT '抽样比例',
    status VARCHAR(20) DEFAULT 'ACTIVE' COMMENT '状态',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_scheme_detail_scheme ON qms_inspection_scheme_detail(scheme_id) WHERE deleted_at IS NULL;
```

---

## 10. AQL与抽样方案API接口

### 10.1 AQL标准管理 (/qms/aql)

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| POST | /qms/aql/create | 创建AQL标准 | AqlCreateReqVO |
| PUT | /qms/aql/update | 更新AQL标准 | AqlUpdateReqVO |
| DELETE | /qms/aql/delete | 删除AQL标准 | Long id |
| GET | /qms/aql/get | 获取AQL标准 | Long id |
| GET | /qms/aql/list | 获取AQL标准列表 | AqlPageReqVO |
| GET | /qms/aql/page | 获取AQL标准分页 | AqlPageReqVO |
| POST | /qms/aql/senior | 高级搜索 | AqlSeniorReqVO |
| GET | /qms/aql/export-excel | 导出Excel | AqlPageReqVO |
| POST | /qms/aql/import | 导入AQL | MultipartFile |
| POST | /qms/aql/enable | 启用AQL | Long id |
| POST | /qms/aql/disable | 禁用AQL | Long id |

**AqlCreateReqVO**:
```json
{
  "code": "AQL-I-0.65",
  "name": "一般检验水平I级AQL0.65",
  "level": "I",
  "value": 0.65,
  "samplingAmount": 80,
  "remark": ""
}
```

### 10.2 物料AQL配置 (/qms/material-aql)

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/material-aql/create | 创建物料AQL配置 |
| PUT | /qms/material-aql/update | 更新物料AQL配置 |
| DELETE | /qms/material-aql/delete | 删除物料AQL配置 |
| GET | /qms/material-aql/get | 获取物料AQL配置 |
| GET | /qms/material-aql/list | 获取物料AQL配置列表 |
| GET | /qms/material-aql/page | 获取物料AQL配置分页 |
| POST | /qms/material-aql/bind | 绑定物料与AQL |
| POST | /qms/material-aql/unbind | 解绑物料与AQL |

### 10.3 抽样计划管理 (/qms/sampling-plan)

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/sampling-plan/create | 创建抽样计划 |
| PUT | /qms/sampling-plan/update | 更新抽样计划 |
| DELETE | /qms/sampling-plan/delete | 删除抽样计划 |
| GET | /qms/sampling-plan/get | 获取抽样计划 |
| GET | /qms/sampling-plan/list | 获取抽样计划列表 |
| GET | /qms/sampling-plan/page | 获取抽样计划分页 |
| POST | /qms/sampling-plan/enable | 启用抽样计划 |
| POST | /qms/sampling-plan/disable | 禁用抽样计划 |
| GET | /qms/sampling-plan/get-sample-size | 根据批量获取样本量 |

### 10.4 抽样规则管理 (/qms/sampling-rule)

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/sampling-rule/create | 创建抽样规则 |
| PUT | /qms/sampling-rule/update | 更新抽样规则 |
| DELETE | /qms/sampling-rule/delete | 删除抽样规则 |
| GET | /qms/sampling-rule/get | 获取抽样规则 |
| GET | /qms/sampling-rule/list | 获取抽样规则列表 |
| GET | /qms/sampling-rule/page | 获取抽样规则分页 |
| POST | /qms/sampling-rule/enable | 启用抽样规则 |
| POST | /qms/sampling-rule/disable | 禁用抽样规则 |

### 10.5 动态检验规则 (/qms/dynamic-rule)

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/dynamic-rule/create | 创建动态规则 |
| PUT | /qms/dynamic-rule/update | 更新动态规则 |
| DELETE | /qms/dynamic-rule/delete | 删除动态规则 |
| GET | /qms/dynamic-rule/get | 获取动态规则 |
| GET | /qms/dynamic-rule/list | 获取动态规则列表 |
| GET | /qms/dynamic-rule/page | 获取动态规则分页 |
| POST | /qms/dynamic-rule/enable | 启用动态规则 |
| POST | /qms/dynamic-rule/disable | 禁用动态规则 |
| POST | /qms/dynamic-rule/test | 测试规则表达式 |

### 10.6 样本编码管理 (/qms/sample-code)

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/sample-code/create | 创建样本编码规则 |
| PUT | /qms/sample-code/update | 更新样本编码规则 |
| DELETE | /qms/sample-code/delete | 删除样本编码规则 |
| GET | /qms/sample-code/get | 获取样本编码规则 |
| GET | /qms/sample-code/list | 获取样本编码规则列表 |
| GET | /qms/sample-code/page | 获取样本编码规则分页 |
| GET | /qms/sample-code/generate | 生成样本编码 | Long ruleId |
| POST | /qms/sample-code/enable | 启用编码规则 |
| POST | /qms/sample-code/disable | 禁用编码规则 |

---

## 11. 检验特性管理API接口

### 11.1 检验特性管理 (/qms/characteristics)

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/characteristics/create | 创建检验特性 |
| PUT | /qms/characteristics/update | 更新检验特性 |
| DELETE | /qms/characteristics/delete | 删除检验特性 |
| GET | /qms/characteristics/get | 获取检验特性 |
| GET | /qms/characteristics/list | 获取检验特性列表 |
| GET | /qms/characteristics/page | 获取检验特性分页 |
| POST | /qms/characteristics/import | 导入检验特性 |
| GET | /qms/characteristics/export | 导出检验特性 |

**CharacteristicsCreateReqVO**:
```json
{
  "code": "CHAR-001",
  "name": "外观检验",
  "characteristicsType": "QUALITATIVE",
  "inspectionMethodId": 1,
  "inspectionTool": "目视",
  "qualificationType": "PASS_FAIL",
  "criticalDefect": 0,
  "majorDefect": 0.5,
  "minorDefect": 1.5
}
```

### 11.2 定量特性管理 (/qms/char-quantitative)

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/char-quantitative/create | 创建定量特性 |
| PUT | /qms/char-quantitative/update | 更新定量特性 |
| DELETE | /qms/char-quantitative/delete | 删除定量特性 |
| GET | /qms/char-quantitative/get | 获取定量特性 |
| GET | /qms/char-quantitative/list | 获取定量特性列表 |
| GET | /qms/char-quantitative/page | 获取定量特性分页 |
| GET | /qms/char-quantitative/get-by-char | 根据特性ID获取 |

### 11.3 定性特性管理 (/qms/char-qualitative)

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/char-qualitative/create | 创建定性特性 |
| PUT | /qms/char-qualitative/update | 更新定性特性 |
| DELETE | /qms/char-qualitative/delete | 删除定性特性 |
| GET | /qms/char-qualitative/get | 获取定性特性 |
| GET | /qms/char-qualitative/list | 获取定性特性列表 |
| GET | /qms/char-qualitative/page | 获取定性特性分页 |
| GET | /qms/char-qualitative/get-by-char | 根据特性ID获取 |

### 11.4 检验方法管理 (/qms/method)

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/method/create | 创建检验方法 |
| PUT | /qms/method/update | 更新检验方法 |
| DELETE | /qms/method/delete | 删除检验方法 |
| GET | /qms/method/get | 获取检验方法 |
| GET | /qms/method/list | 获取检验方法列表 |
| GET | /qms/method/page | 获取检验方法分页 |
| POST | /qms/method/enable | 启用检验方法 |
| POST | /qms/method/disable | 禁用检验方法 |

### 11.5 检验方案管理 (/qms/scheme)

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/scheme/create | 创建检验方案 |
| PUT | /qms/scheme/update | 更新检验方案 |
| DELETE | /qms/scheme/delete | 删除检验方案 |
| GET | /qms/scheme/get | 获取检验方案 |
| GET | /qms/scheme/list | 获取检验方案列表 |
| GET | /qms/scheme/page | 获取检验方案分页 |
| POST | /qms/scheme/detail/add | 添加方案明细 |
| PUT | /qms/scheme/detail/update | 更新方案明细 |
| DELETE | /qms/scheme/detail/delete | 删除方案明细 |
| GET | /qms/scheme/detail/list | 获取方案明细列表 |
| POST | /qms/scheme/enable | 启用检验方案 |
| POST | /qms/scheme/disable | 禁用检验方案 |
| GET | /qms/scheme/get-by-product | 根据产品获取方案 |

**SchemeCreateReqVO**:
```json
{
  "code": "SCHEME-IQC-001",
  "name": "来料检验方案A",
  "schemeType": "IQC",
  "productCode": "MAT-001",
  "aqlId": 1,
  "inspectionLevel": "II",
  "samplingPlanId": 1,
  "criticalLimit": 0,
  "majorLimit": 1,
  "minorLimit": 3
}
```

---

## 12. 补充表结构清单

以下为QMS模块补充的完整表结构汇总：

| 序号 | 表名 | 说明 |
|------|------|------|
| 1 | qms_aql | AQL标准配置表 |
| 2 | qms_aql_detail | AQL明细表 |
| 3 | qms_material_aql | 物料AQL配置表 |
| 4 | qms_sampling_plan | 抽样计划表 |
| 5 | qms_sampling_rule | 抽样规则表 |
| 6 | qms_sampling_record | 抽样记录表 |
| 7 | qms_dynamic_rule | 动态检验规则表 |
| 8 | qms_sample_code | 样本编码表 |
| 9 | qms_inspection_characteristics | 检验特性表 |
| 10 | qms_char_quantitative | 定量特性表 |
| 11 | qms_char_qualitative | 定性特性表 |
| 12 | qms_inspection_method | 检验方法表 |
| 13 | qms_inspection_scheme | 检验方案表 |
| 14 | qms_inspection_scheme_detail | 检验方案明细表 |

---

## 13. 补充API接口汇总

| 模块 | API前缀 | 接口数量 | 说明 |
|------|--------|---------|------|
| AQL标准 | /qms/aql | 12 | AQL标准CRUD+导入导出+启禁用 |
| 物料AQL | /qms/material-aql | 8 | 物料与AQL绑定管理 |
| 抽样计划 | /qms/sampling-plan | 9 | 抽样计划管理 |
| 抽样规则 | /qms/sampling-rule | 8 | 抽样规则管理 |
| 动态规则 | /qms/dynamic-rule | 10 | 动态检验规则 |
| 样本编码 | /qms/sample-code | 9 | 样本编码生成管理 |
| 检验特性 | /qms/characteristics | 8 | 特性定义管理 |
| 定量特性 | /qms/char-quantitative | 6 | 定量特性配置 |
| 定性特性 | /qms/char-qualitative | 6 | 定性特性配置 |
| 检验方法 | /qms/method | 9 | 检验方法管理 |
| 检验方案 | /qms/scheme | 13 | 检验方案+明细管理 |

**补充API合计**: 98个

---

## 14. Q1/Q2/Q3质量通知单

### 14.1 qms_inspection_q1 - Q1通知单表

**业务说明**: Q1质量通知单，用于工序首件检验不合格通知。

```sql
CREATE TABLE IF NOT EXISTS qms_inspection_q1 (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    notice_no VARCHAR(50) NOT NULL COMMENT '通知单号',
    source_type VARCHAR(20) COMMENT '来源类型(工单/报工)',
    source_no VARCHAR(50) COMMENT '来源单号',
    process_id BIGINT COMMENT '工序ID',
    process_code VARCHAR(50) COMMENT '工序编码',
    process_name VARCHAR(200) COMMENT '工序名称',
    workstation_id BIGINT COMMENT '工位ID',
    workstation_code VARCHAR(50) COMMENT '工位编码',
    workstation_name VARCHAR(200) COMMENT '工位名称',
    product_id BIGINT COMMENT '产品ID',
    product_code VARCHAR(50) COMMENT '产品编码',
    product_name VARCHAR(200) COMMENT '产品名称',
    work_order_id BIGINT COMMENT '工单ID',
    work_order_no VARCHAR(50) COMMENT '工单号',
    batch_no VARCHAR(50) COMMENT '批次号',
    quantity DECIMAL(18,3) COMMENT '数量',
    defect_count INT DEFAULT 0 COMMENT '缺陷数量',
    inspector_id BIGINT COMMENT '检验员ID',
    inspector_name VARCHAR(100) COMMENT '检验员姓名',
    inspection_time TIMESTAMP COMMENT '检验时间',
    result VARCHAR(20) COMMENT '检验结果(PASS/FAIL)',
    status VARCHAR(20) DEFAULT 'PENDING' COMMENT '状态(PENDING/SEND/FINISH)',
    is_send INT DEFAULT 0 COMMENT '是否发送(0否/1是)',
    send_time TIMESTAMP COMMENT '发送时间',
    finish_time TIMESTAMP COMMENT '完成时间',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_q1_notice_no ON qms_inspection_q1(tenant_id, notice_no) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_q1_source ON qms_inspection_q1(source_type, source_no) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_q1_status ON qms_inspection_q1(status) WHERE deleted_at IS NULL;
```

**Q1通知单API (/qms/inspectionQ1)**:

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/inspectionQ1/create | 创建Q1通知单 |
| PUT | /qms/inspectionQ1/update | 更新Q1通知单 |
| DELETE | /qms/inspectionQ1/delete | 删除Q1通知单 |
| GET | /qms/inspectionQ1/get | 获取Q1通知单 |
| GET | /qms/inspectionQ1/page | 获取Q1通知单分页 |
| POST | /qms/inspectionQ1/senior | 高级搜索Q1通知单 |
| GET | /qms/inspectionQ1/export-excel | 导出Q1通知单Excel |
| GET | /qms/inspectionQ1/get-import-template | 获取导入Q1通知单模板 |
| GET | /qms/inspectionQ1/finish | 完成Q1通知单 |

### 14.2 qms_inspection_q2 - Q2通知单表

**业务说明**: Q2质量通知单，用于过程检验/巡检不合格通知，支持邮件发送。

```sql
CREATE TABLE IF NOT EXISTS qms_inspection_q2 (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    notice_no VARCHAR(50) NOT NULL COMMENT '通知单号',
    source_type VARCHAR(20) COMMENT '来源类型(IPQC/FQC)',
    source_no VARCHAR(50) COMMENT '来源单号',
    inspection_type VARCHAR(20) COMMENT '检验类型(巡检/专检)',
    process_id BIGINT COMMENT '工序ID',
    process_code VARCHAR(50) COMMENT '工序编码',
    process_name VARCHAR(200) COMMENT '工序名称',
    workstation_id BIGINT COMMENT '工位ID',
    workstation_code VARCHAR(50) COMMENT '工位编码',
    workstation_name VARCHAR(200) COMMENT '工位名称',
    product_id BIGINT COMMENT '产品ID',
    product_code VARCHAR(50) COMMENT '产品编码',
    product_name VARCHAR(200) COMMENT '产品名称',
    work_order_id BIGINT COMMENT '工单ID',
    work_order_no VARCHAR(50) COMMENT '工单号',
    batch_no VARCHAR(50) COMMENT '批次号',
    quantity DECIMAL(18,3) COMMENT '数量',
    defect_count INT DEFAULT 0 COMMENT '缺陷数量',
    defect_level VARCHAR(20) COMMENT '缺陷等级(CR/MA/MI)',
    inspector_id BIGINT COMMENT '检验员ID',
    inspector_name VARCHAR(100) COMMENT '检验员姓名',
    inspection_time TIMESTAMP COMMENT '检验时间',
    result VARCHAR(20) COMMENT '检验结果(PASS/FAIL)',
    status VARCHAR(20) DEFAULT 'PENDING' COMMENT '状态(PENDING/SEND/FINISH)',
    is_send INT DEFAULT 0 COMMENT '是否发送(0否/1是)',
    send_time TIMESTAMP COMMENT '发送时间',
    email VARCHAR(200) COMMENT '通知邮箱',
    finish_time TIMESTAMP COMMENT '完成时间',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_q2_notice_no ON qms_inspection_q2(tenant_id, notice_no) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_q2_source ON qms_inspection_q2(source_type, source_no) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_q2_defect_level ON qms_inspection_q2(defect_level) WHERE deleted_at IS NULL;
```

**Q2通知单API (/qms/inspectionQ2)**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /qms/inspectionQ2/send | Q2通知单发送邮件 |
| POST | /qms/inspectionQ2/create | 创建Q2通知单 |
| PUT | /qms/inspectionQ2/update | 更新Q2通知单 |
| DELETE | /qms/inspectionQ2/delete | 删除Q2通知单 |
| GET | /qms/inspectionQ2/get | 获取Q2通知单 |
| GET | /qms/inspectionQ2/page | 获取Q2通知单分页 |
| POST | /qms/inspectionQ2/senior | 高级搜索Q2通知单 |
| GET | /qms/inspectionQ2/export-excel | 导出Q2通知单Excel |
| GET | /qms/inspectionQ2/get-import-template | 获取导入Q2通知单模板 |
| GET | /qms/inspectionQ2/finish | 完成Q2通知单 |
| GET | /qms/inspectionQ2/getEmail | 获取系统中的email地址 |

### 14.3 qms_inspection_q3_main - Q3通知单主表

**业务说明**: Q3质量通知单，用于成品检验/出货检验不合格通知，支持多项目明细。

```sql
CREATE TABLE IF NOT EXISTS qms_inspection_q3_main (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    notice_no VARCHAR(50) NOT NULL COMMENT '通知单号',
    notice_type VARCHAR(20) DEFAULT 'OQC' COMMENT '通知单类型(OQC/FQC)',
    source_type VARCHAR(20) COMMENT '来源类型',
    source_no VARCHAR(50) COMMENT '来源单号',
    customer_id BIGINT COMMENT '客户ID',
    customer_code VARCHAR(50) COMMENT '客户编码',
    customer_name VARCHAR(200) COMMENT '客户名称',
    product_id BIGINT COMMENT '产品ID',
    product_code VARCHAR(50) COMMENT '产品编码',
    product_name VARCHAR(200) COMMENT '产品名称',
    sales_order_no VARCHAR(50) COMMENT '销售订单号',
    delivery_no VARCHAR(50) COMMENT '发货单号',
    batch_no VARCHAR(50) COMMENT '批次号',
    quantity DECIMAL(18,3) COMMENT '数量',
    defect_count INT DEFAULT 0 COMMENT '缺陷数量',
    inspector_id BIGINT COMMENT '检验员ID',
    inspector_name VARCHAR(100) COMMENT '检验员姓名',
    inspection_time TIMESTAMP COMMENT '检验时间',
    result VARCHAR(20) COMMENT '检验结果(PASS/FAIL)',
    status VARCHAR(20) DEFAULT 'PENDING' COMMENT '状态(PENDING/FINISH)',
    finish_time TIMESTAMP COMMENT '完成时间',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_q3_main_notice_no ON qms_inspection_q3_main(tenant_id, notice_no) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_q3_main_source ON qms_inspection_q3_main(source_type, source_no) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_q3_main_customer ON qms_inspection_q3_main(customer_id) WHERE deleted_at IS NULL;
```

**Q3通知单主API (/qms/inspection-Q3-main)**:

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/inspection-Q3-main/create | 创建Q3通知单主 |
| PUT | /qms/inspection-Q3-main/update | 更新Q3通知单主 |
| DELETE | /qms/inspection-Q3-main/delete | 删除Q3通知单主 |
| GET | /qms/inspection-Q3-main/get | 获取Q3通知单主 |
| GET | /qms/inspection-Q3-main/page | 获取Q3通知单主分页 |
| POST | /qms/inspection-Q3-main/senior | 高级搜索Q3通知单主 |
| GET | /qms/inspection-Q3-main/export-excel | 导出Q3通知单主Excel |
| POST | /qms/inspection-Q3-main/export-excel-senior | 导出Q3通知单主Excel(高级搜索) |
| GET | /qms/inspection-Q3-main/finish | 完成Q3通知单主 |

### 14.4 qms_inspection_q3_detail - Q3通知单子表

**业务说明**: Q3通知单明细，记录每个不合格项目的详细信息。

```sql
CREATE TABLE IF NOT EXISTS qms_inspection_q3_detail (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    main_id BIGINT NOT NULL COMMENT '主表ID',
    line_no INT DEFAULT 1 COMMENT '行号',
    defect_code VARCHAR(50) COMMENT '缺陷代码',
    defect_name VARCHAR(200) COMMENT '缺陷名称',
    defect_classification VARCHAR(20) COMMENT '缺陷分类(CR/MA/MI)',
    defect_count INT DEFAULT 0 COMMENT '缺陷数量',
    defect_rate DECIMAL(10,4) COMMENT '缺陷率',
    sample_size INT COMMENT '样本数量',
    description TEXT COMMENT '缺陷描述',
    photos VARCHAR(1000) COMMENT '缺陷照片(JSON数组)',
    handle_suggestion VARCHAR(500) COMMENT '处理建议',
    handle_result VARCHAR(200) COMMENT '处理结果',
    handle_time TIMESTAMP COMMENT '处理时间',
    handle_user_id BIGINT COMMENT '处理人ID',
    handle_user_name VARCHAR(100) COMMENT '处理人姓名',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_q3_detail_main_id ON qms_inspection_q3_detail(main_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_q3_detail_classification ON qms_inspection_q3_detail(defect_classification) WHERE deleted_at IS NULL;
```

**Q3通知单子API (/qms/inspection-Q3-detail)**:

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/inspection-Q3-detail/create | 创建Q3通知单子 |
| PUT | /qms/inspection-Q3-detail/update | 更新Q3通知单子 |
| DELETE | /qms/inspection-Q3-detail/delete | 删除Q3通知单子 |
| GET | /qms/inspection-Q3-detail/get | 获取Q3通知单子 |
| GET | /qms/inspection-Q3-detail/page | 获取Q3通知单子分页 |
| POST | /qms/inspection-Q3-detail/senior | 高级搜索Q3通知单子 |
| GET | /qms/inspection-Q3-detail/export-excel | 导出Q3通知单子Excel |
| GET | /qms/inspection-Q3-detail/get-import-template | 获取导入Q3通知单子模板 |

---

## 15. 检验申请管理

### 15.1 qms_request_inspection_main - 检验申请主表

**业务说明**: 接收来自仓库、生产等模块的检验申请，支持审批流程。

```sql
CREATE TABLE IF NOT EXISTS qms_request_inspection_main (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    request_no VARCHAR(50) NOT NULL COMMENT '申请单号',
    request_type VARCHAR(20) NOT NULL COMMENT '申请类型(IQC/FQC/OQC)',
    source_type VARCHAR(20) COMMENT '来源类型(采购/工单/发货)',
    source_no VARCHAR(50) COMMENT '来源单号',
    warehouse_id BIGINT COMMENT '仓库ID',
    warehouse_code VARCHAR(50) COMMENT '仓库编码',
    warehouse_name VARCHAR(200) COMMENT '仓库名称',
    applicant_id BIGINT COMMENT '申请人ID',
    applicant_name VARCHAR(100) COMMENT '申请人姓名',
    apply_time TIMESTAMP COMMENT '申请时间',
    product_id BIGINT COMMENT '产品/物料ID',
    product_code VARCHAR(50) COMMENT '产品/物料编码',
    product_name VARCHAR(200) COMMENT '产品/物料名称',
    batch_no VARCHAR(50) COMMENT '批次号',
    quantity DECIMAL(18,3) COMMENT '申请数量',
    sample_quantity INT COMMENT '抽样数量',
    inspection_scheme_id BIGINT COMMENT '检验方案ID',
    inspection_scheme_code VARCHAR(50) COMMENT '检验方案编码',
    inspection_scheme_name VARCHAR(200) COMMENT '检验方案名称',
    priority INT DEFAULT 0 COMMENT '优先级(0普通/1紧急)',
    expected_time TIMESTAMP COMMENT '期望检验时间',
    status VARCHAR(20) DEFAULT 'PENDING' COMMENT '状态(PENDING/SUBMITTED/APPROVED/REJECTED/HANDLED/CLOSED)',
    approval_id BIGINT COMMENT '审批人ID',
    approval_name VARCHAR(100) COMMENT '审批人姓名',
    approval_time TIMESTAMP COMMENT '审批时间',
    approval_remark VARCHAR(500) COMMENT '审批备注',
    handler_id BIGINT COMMENT '处理人ID',
    handler_name VARCHAR(100) COMMENT '处理人姓名',
    handle_time TIMESTAMP COMMENT '处理时间',
    handle_remark VARCHAR(500) COMMENT '处理备注',
    finish_time TIMESTAMP COMMENT '完成时间',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_request_main_no ON qms_request_inspection_main(tenant_id, request_no) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_request_main_source ON qms_request_inspection_main(source_type, source_no) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_request_main_status ON qms_request_inspection_main(status) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_request_main_applicant ON qms_request_inspection_main(applicant_id) WHERE deleted_at IS NULL;
```

**检验申请主API (/qms/inspection-request-main)**:

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/inspection-request-main/create | 创建检验申请 |
| PUT | /qms/inspection-request-main/update | 更新检验申请 |
| DELETE | /qms/inspection-request-main/delete | 删除检验申请 |
| GET | /qms/inspection-request-main/get | 获取检验申请 |
| GET | /qms/inspection-request-main/page | 获取检验申请分页 |
| POST | /qms/inspection-request-main/senior | 高级搜索检验申请 |
| GET | /qms/inspection-request-main/export-excel | 导出检验申请Excel |
| GET | /qms/inspection-request-main/get-import-template | 获取导入检验申请模板 |
| PUT | /qms/inspection-request-main/close | 关闭检验申请 |
| PUT | /qms/inspection-request-main/submit | 提交检验申请 |
| PUT | /qms/inspection-request-main/reAdd | 重新添加检验申请 |
| PUT | /qms/inspection-request-main/agree | 审批通过检验申请 |
| PUT | /qms/inspection-request-main/handle | 执行检验申请 |
| PUT | /qms/inspection-request-main/refused | 审批拒绝检验申请 |

### 15.2 qms_request_inspection_package - 检验申请包装表

**业务说明**: 检验申请包装信息，记录每个包装的检验详情。

```sql
CREATE TABLE IF NOT EXISTS qms_request_inspection_package (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    main_id BIGINT NOT NULL COMMENT '主表ID',
    package_no VARCHAR(50) COMMENT '包装编号',
    package_spec VARCHAR(100) COMMENT '包装规格',
    quantity DECIMAL(18,3) COMMENT '数量',
    sample_quantity INT COMMENT '抽样数量',
    unit VARCHAR(20) COMMENT '单位',
    batch_no VARCHAR(50) COMMENT '批次号',
    location VARCHAR(200) COMMENT '库位',
    status VARCHAR(20) DEFAULT 'PENDING' COMMENT '状态',
    inspection_result VARCHAR(20) COMMENT '检验结果(PASS/FAIL)',
    qualified_quantity DECIMAL(18,3) COMMENT '合格数量',
    unqualified_quantity DECIMAL(18,3) COMMENT '不合格数量',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_request_package_main_id ON qms_request_inspection_package(main_id) WHERE deleted_at IS NULL;
```

**检验申请包装API (/qms/inspection-request-package)**:

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/inspection-request-package/create | 创建检验申请包装 |
| PUT | /qms/inspection-request-package/update | 更新检验申请包装 |
| DELETE | /qms/inspection-request-package/delete | 删除检验申请包装 |
| GET | /qms/inspection-request-package/get | 获取检验申请包装 |
| GET | /qms/inspection-request-package/list | 获取检验申请包装列表 |

---

## 16. 其他补充表

### 16.1 qms_inspection_process - 检验工序表

**业务说明**: 定义检验方案中的检验工序及顺序。

```sql
CREATE TABLE IF NOT EXISTS qms_inspection_process (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    code VARCHAR(50) NOT NULL COMMENT '工序编码',
    name VARCHAR(200) NOT NULL COMMENT '工序名称',
    template_code VARCHAR(50) COMMENT '模板编码',
    template_name VARCHAR(200) COMMENT '模板名称',
    scheme_id BIGINT COMMENT '检验方案ID',
    scheme_code VARCHAR(50) COMMENT '检验方案编码',
    sequence INT DEFAULT 0 COMMENT '顺序号',
    inspection_type VARCHAR(20) COMMENT '检验类型(首件/巡检/末件)',
    is_mandatory INT DEFAULT 1 COMMENT '是否必检(0否/1是)',
    timeout_minutes INT COMMENT '标准时长(分钟)',
    status VARCHAR(20) DEFAULT 'ACTIVE' COMMENT '状态(ACTIVE/INACTIVE)',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_process_code ON qms_inspection_process(tenant_id, code) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_process_scheme ON qms_inspection_process(scheme_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_process_sequence ON qms_inspection_process(sequence) WHERE deleted_at IS NULL;
```

**检验工序API (/qms/inspection-process)**:

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/inspection-process/create | 创建检验工序 |
| PUT | /qms/inspection-process/update | 更新检验工序 |
| DELETE | /qms/inspection-process/delete | 删除检验工序 |
| GET | /qms/inspection-process/get | 获取检验工序 |
| GET | /qms/inspection-process/list | 获取检验工序列表 |
| GET | /qms/inspection-process/page | 获取检验工序分页 |
| GET | /qms/inspection-process/export-excel | 导出检验工序Excel |
| GET | /qms/inspection-process/get-import-template | 获取导入检验工序模板 |
| POST | /qms/inspection-process/import | 导入检验工序 |
| GET | /qms/inspection-process/getListByTempleteCode | 根据模板code获取工序列表及检验特性 |

### 16.2 qms_inspection_stage - 检验阶段表

**业务说明**: 定义检验的不同阶段(如IQC/PQC/FQC/OQC)。

```sql
CREATE TABLE IF NOT EXISTS qms_inspection_stage (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    code VARCHAR(50) NOT NULL COMMENT '阶段编码',
    name VARCHAR(200) NOT NULL COMMENT '阶段名称',
    type VARCHAR(20) NOT NULL COMMENT '阶段类型(IQC/PQC/FQC/OQC)',
    sequence INT DEFAULT 0 COMMENT '顺序号',
    color VARCHAR(20) COMMENT '标签颜色',
    icon VARCHAR(50) COMMENT '图标',
    is_enabled INT DEFAULT 1 COMMENT '是否启用(0否/1是)',
    status VARCHAR(20) DEFAULT 'ACTIVE' COMMENT '状态',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_stage_code ON qms_inspection_stage(tenant_id, code) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_stage_type ON qms_inspection_stage(type) WHERE deleted_at IS NULL;
```

**检验阶段API (/qms/inspection-stage)**:

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/inspection-stage/create | 创建检验阶段 |
| PUT | /qms/inspection-stage/update | 更新检验阶段 |
| DELETE | /qms/inspection-stage/delete | 删除检验阶段 |
| GET | /qms/inspection-stage/get | 获取检验阶段 |
| GET | /qms/inspection-stage/list | 获取检验阶段列表 |
| GET | /qms/inspection-stage/page | 获取检验阶段分页 |
| POST | /qms/inspection-stage/senior | 高级搜索检验阶段 |
| GET | /qms/inspection-stage/export-excel | 导出检验阶段Excel |
| GET | /qms/inspection-stage/get-import-template | 获取导入检验阶段模板 |
| POST | /qms/inspection-stage/import | 导入检验阶段 |
| GET | /qms/inspection-stage/noPage | 获取检验阶段不分页列表 |

### 16.3 qms_programme_template - 检验方案模板表

**业务说明**: 定义可复用的检验方案模板。

```sql
CREATE TABLE IF NOT EXISTS qms_programme_template (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    code VARCHAR(50) NOT NULL COMMENT '模板编码',
    name VARCHAR(200) NOT NULL COMMENT '模板名称',
    type VARCHAR(20) NOT NULL COMMENT '模板类型(IQC/PQC/FQC/OQC)',
    version VARCHAR(20) COMMENT '版本号',
    description TEXT COMMENT '模板描述',
    product_type VARCHAR(50) COMMENT '适用产品类型',
    inspection_level VARCHAR(20) COMMENT '检验水平',
    aql_id BIGINT COMMENT 'AQL标准ID',
    aql_code VARCHAR(50) COMMENT 'AQL代码',
    sampling_type VARCHAR(20) DEFAULT 'NORMAL' COMMENT '抽样类型',
    process_codes TEXT COMMENT '工序编码列表(JSON数组)',
    characteristics_codes TEXT COMMENT '检验特性编码列表(JSON数组)',
    status VARCHAR(20) DEFAULT 'DRAFT' COMMENT '状态(DRAFT/ACTIVE/INACTIVE)',
    effective_date DATE COMMENT '生效日期',
    expiry_date DATE COMMENT '失效日期',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_template_code ON qms_programme_template(tenant_id, code) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_template_type ON qms_programme_template(type) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_template_status ON qms_programme_template(status) WHERE deleted_at IS NULL;
```

**检验方案模板API (/qms/programme-template)**:

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/programme-template/create | 创建检验方案模板 |
| PUT | /qms/programme-template/update | 更新检验方案模板 |
| DELETE | /qms/programme-template/delete | 删除检验方案模板 |
| GET | /qms/programme-template/get | 获取检验方案模板 |
| GET | /qms/programme-template/list | 获取检验方案模板列表 |
| GET | /qms/programme-template/page | 获取检验方案模板分页 |
| POST | /qms/programme-template/senior | 高级搜索检验方案模板 |
| GET | /qms/programme-template/export-excel | 导出检验方案模板Excel |
| POST | /qms/programme-template/export-excel-senior | 导出检验方案模板Excel(高级搜索) |
| GET | /qms/programme-template/get-import-template | 获取导入检验方案模板模板 |
| POST | /qms/programme-template/import | 导入检验方案模板 |
| POST | /qms/programme-template/enable | 启用检验方案模板 |
| POST | /qms/programme-template/disable | 禁用检验方案模板 |

### 16.4 qms_selected_set - 选定集表

**业务说明**: 定义质量检验的检验项选定集。

```sql
CREATE TABLE IF NOT EXISTS qms_selected_set (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    code VARCHAR(50) NOT NULL COMMENT '选定集编码',
    name VARCHAR(200) NOT NULL COMMENT '选定集名称',
    set_type VARCHAR(20) NOT NULL COMMENT '选定集类型(特性/缺陷/其他)',
    description TEXT COMMENT '描述',
    applicable_object VARCHAR(50) COMMENT '适用对象(产品/物料/工单)',
    applicable_object_codes TEXT COMMENT '适用对象编码列表(JSON)',
    is_system INT DEFAULT 0 COMMENT '是否系统预设(0否/1是)',
    status VARCHAR(20) DEFAULT 'ACTIVE' COMMENT '状态',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_set_code ON qms_selected_set(tenant_id, code) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_set_type ON qms_selected_set(set_type) WHERE deleted_at IS NULL;
```

**选定集API (/qms/selected-set)**:

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/selected-set/create | 创建选定集 |
| PUT | /qms/selected-set/update | 更新选定集 |
| DELETE | /qms/selected-set/delete | 删除选定集 |
| GET | /qms/selected-set/get | 获取选定集 |
| GET | /qms/selected-set/list | 获取选定集列表 |
| GET | /qms/selected-set/page | 获取选定集分页 |
| POST | /qms/selected-set/senior | 高级搜索选定集 |
| GET | /qms/selected-set/export-excel | 导出选定集Excel |
| POST | /qms/selected-set/export-excel-senior | 导出选定集Excel(高级搜索) |
| GET | /qms/selected-set/get-import-template | 获取导入选定集模板 |
| POST | /qms/selected-set/import | 导入选定集 |
| POST | /qms/selected-set/enable | 启用选定集 |
| POST | /qms/selected-set/disable | 禁用选定集 |

### 16.5 qms_selected_project - 选定集项目表

**业务说明**: 选定集中的具体检验项目明细。

```sql
CREATE TABLE IF NOT EXISTS qms_selected_project (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 0,
    set_id BIGINT NOT NULL COMMENT '选定集ID',
    set_code VARCHAR(50) NOT NULL COMMENT '选定集编码',
    code VARCHAR(50) NOT NULL COMMENT '项目编码',
    name VARCHAR(200) NOT NULL COMMENT '项目名称',
    project_type VARCHAR(20) COMMENT '项目类型',
    characteristics_id BIGINT COMMENT '检验特性ID',
    characteristics_code VARCHAR(50) COMMENT '检验特性编码',
    is_mandatory INT DEFAULT 0 COMMENT '是否必检(0否/1是)',
    inspection_method_id BIGINT COMMENT '检验方法ID',
    inspection_method_code VARCHAR(50) COMMENT '检验方法代码',
    inspection_method_name VARCHAR(200) COMMENT '检验方法名称',
    qualification_type VARCHAR(20) COMMENT '合格类型(PASS_FAIL/RANGE/VALUE)',
    usl DECIMAL(18,6) COMMENT '上规格限',
    lsl DECIMAL(18,6) COMMENT '下规格限',
    target_value DECIMAL(18,6) COMMENT '目标值',
    unit VARCHAR(20) COMMENT '单位',
    sequence INT DEFAULT 0 COMMENT '顺序号',
    status VARCHAR(20) DEFAULT 'ACTIVE' COMMENT '状态',
    remark VARCHAR(500) COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_selected_project_set ON qms_selected_project(set_id) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS idx_tenant_set_project_code ON qms_selected_project(tenant_id, set_id, code) WHERE deleted_at IS NULL;
```

**选定集项目API (/qms/selected-project)**:

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /qms/selected-project/create | 创建选定集项目 |
| PUT | /qms/selected-project/update | 更新选定集项目 |
| DELETE | /qms/selected-project/delete | 删除选定集项目 |
| GET | /qms/selected-project/get | 获取选定集项目 |
| GET | /qms/selected-project/list | 获取选定集项目列表 |
| GET | /qms/selected-project/page | 获取选定集项目分页 |
| POST | /qms/selected-project/senior | 高级搜索选定集项目 |
| GET | /qms/selected-project/export-excel | 导出选定集项目Excel |
| GET | /qms/selected-project/get-import-template | 获取导入选定集项目模板 |
| POST | /qms/selected-project/import | 导入选定集项目 |
| GET | /qms/selected-project/noPage | 获取选定集项目不分页 |

---

## 17. 补充表结构汇总

| 序号 | 表名 | 说明 |
|------|------|------|
| 1 | qms_inspection_q1 | Q1通知单表 |
| 2 | qms_inspection_q2 | Q2通知单表 |
| 3 | qms_inspection_q3_main | Q3通知单主表 |
| 4 | qms_inspection_q3_detail | Q3通知单子表 |
| 5 | qms_request_inspection_main | 检验申请主表 |
| 6 | qms_request_inspection_package | 检验申请包装表 |
| 7 | qms_inspection_process | 检验工序表 |
| 8 | qms_inspection_stage | 检验阶段表 |
| 9 | qms_programme_template | 检验方案模板表 |
| 10 | qms_selected_set | 选定集表 |
| 11 | qms_selected_project | 选定集项目表 |

---

## 18. 补充API接口汇总

| 模块 | API前缀 | 接口数量 | 说明 |
|------|--------|---------|------|
| Q1通知单 | /qms/inspectionQ1 | 9 | Q1工序首件检验通知单 |
| Q2通知单 | /qms/inspectionQ2 | 11 | Q2过程检验通知单(支持邮件) |
| Q3通知单主 | /qms/inspection-Q3-main | 10 | Q3成品/出货通知单主 |
| Q3通知单子 | /qms/inspection-Q3-detail | 8 | Q3通知单明细 |
| 检验申请主 | /qms/inspection-request-main | 14 | 检验申请管理(审批流程) |
| 检验申请包装 | /qms/inspection-request-package | 5 | 检验申请包装 |
| 检验工序 | /qms/inspection-process | 10 | 检验工序管理 |
| 检验阶段 | /qms/inspection-stage | 11 | 检验阶段管理 |
| 检验方案模板 | /qms/programme-template | 14 | 方案模板管理 |
| 选定集 | /qms/selected-set | 14 | 选定集管理 |
| 选定集项目 | /qms/selected-project | 11 | 选定集项目明细 |

**本节补充API合计**: 117个

---

## 19. 与现有章节API合计

| 章节 | API数量 |
|------|--------|
| 第10章 AQL与抽样方案API | 98 |
| 第18章 补充API | 117 |
| **总计** | **215** |
