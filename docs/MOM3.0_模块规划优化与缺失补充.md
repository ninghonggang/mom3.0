# MOM3.0 模块规划优化与缺失补充

**版本**: V1.1 | **日期**: 2026-04-09 | **项目**: 闻荫科技MOM3.0

---

## 一、现有文档分析

### 1.1 已覆盖模块

| 文档 | 覆盖模块 | 完成度 |
|------|---------|--------|
| MOM3.0_架构与主文档.md | 全局架构、技术栈、模块索引 | 70% |
| MOM3.0_MES生产执行设计文档.md | M03/M04/M06 | 65% |
| MOM3.0_QMS质量设计文档.md | M05/M09/M10/M11/M12/M13/M14/M15 | 55% |
| MOM3.0_WMS仓储设计文档.md | M07/M08 | 50% |
| MOM3.0_基础数据设计文档.md | M01/M02 | 60% |
| MOM3.0_M09安灯系统设计文档.md | M09 | 75% |

### 1.2 设计缺失清单

#### A. 完全缺失的模块/功能

| 序号 | 模块/功能 | 影响 | 说明 |
|------|---------|------|------|
| 1 | **电子SOP (e-SOP)** | M03生产执行 | 工艺文件电子化，推送到工位终端 |
| 2 | **实验室管理系统** | M11 | 只有SPC设计，缺少样品管理、仪器管理、检测申请 |
| 3 | **量检具管理系统** | M12 | 测量设备(游标卡尺/千分表等)管理、计量校准提醒 |
| 4 | **低代码报表平台** | M15 | 可视化报表配置平台，用户自定义报表 |
| 5 | **QRCI质量闭环** | M05 | 不良品根本原因分析→整改→验证闭环 |
| 6 | **分层审核(LPA)** | M05 | 标准化过程审核，逐层审核表单 |
| 7 | **调拨管理** | M07 | 仓库之间调拨 |
| 8 | **盘点管理** | M07 | 周期盘点/盲盘 |
| 9 | **Alert告警中心** | 全局 | 统一告警通知(质量/设备/库存/APS) |

#### B. 设计不完整的模块

| 序号 | 模块 | 缺失内容 |
|------|------|---------|
| 1 | APS | 滚动排程(每日自动触发)、交付率分析、缺料分析、APS工作日历(班次) |
| 2 | WMS | 线边库位配置、仓库库区(Zone)划分、物料批次规则 |
| 3 | QMS | 检验标准与物料/工序关联、AQL抽样方案 |
| 4 | EAM | 备件消耗成本、关键设备TEEP与运行状态关联 |
| 5 | Andon | 升级规则L1/L2/L3通知对象配置、升级历史记录 |
| 6 | M03生产 | 派工单(Dispatch)、工单变更(工艺/数量变更) |
| 7 | M01系统 | 多车间配置(ERP编码/产能)、打印模板绑定规则 |

---

## 二、新增模块详细设计

### 2.1 电子SOP (e-SOP)

**定位**: 工艺文件电子化管理，推送至工位终端显示

```sql
-- 工艺文件表
CREATE TABLE mes_process_document (
    id BIGSERIAL PRIMARY KEY,
    doc_code VARCHAR(50) UNIQUE NOT NULL,
    doc_name VARCHAR(200) NOT NULL,
    doc_type VARCHAR(20) NOT NULL,      -- SOP/作业指导书/检验标准/安全须知
    version VARCHAR(20) NOT NULL,
    product_id BIGINT,
    process_id BIGINT,
    file_urls JSONB,                    -- [{name, url, type}]
    thumbnail_url VARCHAR(500),
    effective_date DATE,
    expiry_date DATE,
    status VARCHAR(20) DEFAULT 'ACTIVE', -- DRAFT/ACTIVE/ARCHIVED
    is_mandatory SMALLINT DEFAULT 1,     -- 是否必读
    tenant_id BIGINT NOT NULL,
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW()
);

-- SOP推送到工位
CREATE TABLE mes_sop_push_log (
    id BIGSERIAL PRIMARY KEY,
    document_id BIGINT NOT NULL,
    workstation_id BIGINT NOT NULL,
    production_order_id BIGINT,
    push_time TIMESTAMP DEFAULT NOW(),
    read_status VARCHAR(20) DEFAULT 'UNREAD', -- UNREAD/READ/ACKNOWLEDGED
    read_time TIMESTAMP,
    operator_id BIGINT,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 电子签名配置
CREATE TABLE mes_esignature_config (
    id BIGSIAL PRIMARY KEY,
    biz_type VARCHAR(50) NOT NULL,      -- FIRST_INSPECT/LAST_INSPECT/QUALITY_SIGN/OPS_SIGN
    require_password SMALLINT DEFAULT 1,
    allow_pin替代 SMALLINT DEFAULT 0,
    record_location SMALLINT DEFAULT 1,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

**API设计**:
```
GET    /api/v1/production/sop/documents         # 工艺文件列表
POST   /api/v1/production/sop/documents         # 上传工艺文件
GET    /api/v1/production/sop/:productId/:processId  # 获取有效SOP
POST   /api/v1/production/sop/push              # 推送到工位
PUT    /api/v1/production/sop/:id/acknowledge   # 确认阅读(电子签名)
GET    /api/v1/production/sop/push-log/:workstationId  # 工位SOP推送记录
```

---

### 2.2 实验室管理系统 (M11扩展)

**定位**: 样品管理、仪器管理、检测委托

```sql
-- 实验室检测申请
CREATE TABLE lab_inspect_request (
    id BIGSERIAL PRIMARY KEY,
    request_no VARCHAR(50) UNIQUE NOT NULL,
    request_type VARCHAR(20) NOT NULL,   -- INTERNAL/EXTERNAL
    applicant_id BIGINT NOT NULL,
    applicant_name VARCHAR(50),
    department VARCHAR(100),
    sample_name VARCHAR(200) NOT NULL,
    sample_code VARCHAR(100),
    sample_batch VARCHAR(100),
    sample_qty DECIMAL(18,3),
    sample_source VARCHAR(100),           -- 来料/制程/客户退回
    inspect_items JSONB NOT NULL,        -- [{item_name, standard, method}]
    required_date DATE,
    priority VARCHAR(10) DEFAULT 'NORMAL', -- URGENT/NORMAL/LOW
    status VARCHAR(20) DEFAULT 'SUBMITTED', -- SUBMITTED/ACCEPTED/TESTING/COMPLETED/CANCELLED
    assign_to BIGINT,
    remark VARCHAR(500),
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 实验室检测结果
CREATE TABLE lab_inspect_result (
    id BIGSERIAL PRIMARY KEY,
    request_id BIGINT NOT NULL,
    result_no VARCHAR(50) UNIQUE NOT NULL,
    inspect_time TIMESTAMP NOT NULL,
    inspector_id BIGINT,
    inspector_name VARCHAR(50),
    result_conclusion VARCHAR(20),       -- PASS/FAIL/PENDING
    inspect_data JSONB,                  -- [{item, value, unit, conclusion}]
    is_out_of_standard SMALLINT DEFAULT 0,
    attachments JSONB,                    -- 报告文件URL
    raw_data_url VARCHAR(500),            -- 原始数据文件
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 实验室仪器管理
CREATE TABLE lab_instrument (
    id BIGSERIAL PRIMARY KEY,
    instrument_code VARCHAR(50) UNIQUE NOT NULL,
    instrument_name VARCHAR(100) NOT NULL,
    instrument_type VARCHAR(30),         -- 三坐标/投影仪/硬度计/拉力机
    manufacturer VARCHAR(100),
    model_no VARCHAR(100),
    serial_no VARCHAR(100),               -- 出厂编号
    calibration_cycle INTEGER,          -- 校准周期(月)
    last_calibration_date DATE,
    next_calibration_date DATE,
    calibration_status VARCHAR(20) DEFAULT 'VALID', -- VALID/EXPIRING/EXPIRED
    location VARCHAR(100),
    status VARCHAR(20) DEFAULT 'ACTIVE', -- ACTIVE/MAINTENANCE/SCRAPPED
    remark VARCHAR(500),
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 仪器校准记录
CREATE TABLE lab_calibration_record (
    id BIGSERIAL PRIMARY KEY,
    instrument_id BIGINT NOT NULL,
    calibration_date DATE NOT NULL,
    next_calibration_date DATE,
    calibration_result VARCHAR(20),       -- PASS/FAIL
    calibration_org VARCHAR(100),       -- 校准机构
    certificate_no VARCHAR(100),
    certificate_url VARCHAR(500),
    remark VARCHAR(500),
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

**API设计**:
```
# 检测申请
GET    /api/v1/lab/requests
POST   /api/v1/lab/requests
PUT    /api/v1/lab/requests/:id
PUT    /api/v1/lab/requests/:id/accept      # 受理
PUT    /api/v1/lab/requests/:id/complete    # 完成检测

# 检测结果
POST   /api/v1/lab/results
GET    /api/v1/lab/results/:requestId

# 仪器管理
GET    /api/v1/lab/instruments
POST   /api/v1/lab/instruments
PUT    /api/v1/lab/instruments/:id
GET    /api/v1/lab/instruments/:id/calibration-history
GET    /api/v1/lab/instruments/expiring      # 即将到期校准
```

---

### 2.3 量检具管理系统

**定位**: 测量设备(游标卡尺/千分表/高度尺等)管理、计量校准提醒

```sql
-- 量检具台账
CREATE TABLE eam_gauge (
    id BIGSERIAL PRIMARY KEY,
    gauge_code VARCHAR(50) UNIQUE NOT NULL,
    gauge_name VARCHAR(100) NOT NULL,
    gauge_type VARCHAR(30) NOT NULL,     -- 卡尺/千分尺/高度尺/塞规/螺纹规/其他
    specification VARCHAR(100),          -- 规格型号
    precision_level VARCHAR(20),        -- 精度等级(0.01mm/0.001mm)
    measurement_range VARCHAR(50),     -- 测量范围
    manufacturer VARCHAR(100),
    purchase_date DATE,
    purchase_cost DECIMAL(18,2),
    calibration_cycle INTEGER,          -- 校准周期(月)
    last_calibration_date DATE,
    next_calibration_date DATE,
    calibration_status VARCHAR(20) DEFAULT 'VALID', -- VALID/EXPIRING/EXPIRED/OVERDUE
    current_location VARCHAR(100),      -- 当前位置(工位/仓库)
    current_holder BIGINT,              -- 当前持有人
    status VARCHAR(20) DEFAULT 'ACTIVE', -- ACTIVE/MAINTENANCE/RETIRED
    remark VARCHAR(500),
    tenant_id BIGINT NOT NULL,
    workshop_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 量检具校准记录
CREATE TABLE eam_gauge_calibration (
    id BIGSERIAL PRIMARY KEY,
    gauge_id BIGINT NOT NULL,
    calibration_date DATE NOT NULL,
    calibration_type VARCHAR(20),        -- INTERNAL/EXTERNAL
    calibration_result VARCHAR(20),     -- PASS/FAIL
    calibrator VARCHAR(100),             -- 校准员/校准机构
    certificate_no VARCHAR(100),
    certificate_url VARCHAR(500),
    deviation_values JSONB,              -- [{point, standard, actual, deviation}]
    is_within_tolerance SMALLINT DEFAULT 1,
    next_calibration_date DATE,
    remark VARCHAR(500),
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 量检具借用记录
CREATE TABLE eam_gauge_borrow (
    id BIGSERIAL PRIMARY KEY,
    gauge_id BIGINT NOT NULL,
    borrower_id BIGINT NOT NULL,
    borrower_name VARCHAR(50),
    borrow_time TIMESTAMP NOT NULL,
    expected_return_time TIMESTAMP,
    actual_return_time TIMESTAMP,
    purpose VARCHAR(200),
    status VARCHAR(20) DEFAULT 'BORROWED', -- BORROWED/RETURNED/OVERDUE
    remark VARCHAR(500),
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

**API设计**:
```
GET    /api/v1/equipment/gauges
POST   /api/v1/equipment/gauges
PUT    /api/v1/equipment/gauges/:id
GET    /api/v1/equipment/gauges/expiring          # 校准到期提醒
GET    /api/v1/equipment/gauges/:id/history        # 校准历史

# 借用管理
POST   /api/v1/equipment/gauges/:id/borrow
PUT    /api/v1/equipment/gauges/:id/return
GET    /api/v1/equipment/gauges/:id/borrow-log    # 借用记录
```

---

### 2.4 QRCI质量闭环

**定位**: 不良品根本原因分析→纠正措施→验证闭环

```sql
-- QRCI头表
CREATE TABLE qms_qrci (
    id BIGSERIAL PRIMARY KEY,
    qrci_no VARCHAR(50) UNIQUE NOT NULL,
    source_type VARCHAR(20),            -- NCR/客户投诉/内审/巡检
    source_id BIGINT,                   -- 来源单据ID
    defect_description TEXT NOT NULL,
    severity_level VARCHAR(10),         -- 严重/重要/一般
    discovery_date DATE NOT NULL,
    discovery_location VARCHAR(100),
    responsible_dept_id BIGINT,
    responsible_dept_name VARCHAR(100),
    owner_id BIGINT,                     -- 责任人
    owner_name VARCHAR(50),
    target_close_date DATE,
    actual_close_date DATE,
    status VARCHAR(20) DEFAULT 'OPEN',  -- OPEN/IN_PROGRESS/VERIFIED/CLOSED/CANCELLED
    verification_result VARCHAR(20),
    verification_by BIGINT,
    verification_time TIMESTAMP,
    remark VARCHAR(500),
    tenant_id BIGINT NOT NULL,
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW()
);

-- QRCI 5Why分析
CREATE TABLE qms_qrci_5why (
    id BIGSERIAL PRIMARY KEY,
    qrci_id BIGINT NOT NULL,
    why_level INTEGER NOT NULL,         -- 1-5
    question VARCHAR(500) NOT NULL,
    answer VARCHAR(500) NOT NULL,
    is_root_cause SMALLINT DEFAULT 0,   -- 是否为根本原因
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- QRCI 纠正措施
CREATE TABLE qms_qrci_action (
    id BIGSERIAL PRIMARY KEY,
    qrci_id BIGINT NOT NULL,
    action_type VARCHAR(20),            -- IMMEDIATE/CORRECTIVE/PREVENTIVE
    action_description VARCHAR(500) NOT NULL,
    responsible_id BIGINT NOT NULL,
    responsible_name VARCHAR(50),
    due_date DATE NOT NULL,
    completed_date DATE,
    evidence_urls JSONB,                -- 证据文件
    status VARCHAR(20) DEFAULT 'PENDING', -- PENDING/IN_PROGRESS/COMPLETED/VERIFIED
    verification_result VARCHAR(20),
    verification_by BIGINT,
    verification_time TIMESTAMP,
    remark VARCHAR(500),
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- QRCI 效果确认
CREATE TABLE qms_qrci_verification (
    id BIGSERIAL PRIMARY KEY,
    qrci_id BIGINT NOT NULL,
    verification_date DATE NOT NULL,
    verifier_id BIGINT NOT NULL,
    verifier_name VARCHAR(50),
    effectiveness VARCHAR(20),          -- EFFECTIVE/INEFFECTIVE/NEED_IMPROVEMENT
    evidence_description VARCHAR(500),
    evidence_urls JSONB,
    follow_up_required SMALLINT DEFAULT 0,
    follow_up_remark VARCHAR(500),
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

**API设计**:
```
GET    /api/v1/quality/qrci
POST   /api/v1/quality/qrci
GET    /api/v1/quality/qrci/:id
PUT    /api/v1/quality/qrci/:id
PUT    /api/v1/quality/qrci/:id/close         # 关闭QRCI

# 5Why分析
GET    /api/v1/quality/qrci/:id/5why
POST   /api/v1/quality/qrci/:id/5why         # 添加5Why记录

# 纠正措施
GET    /api/v1/quality/qrci/:id/actions
POST   /api/v1/quality/qrci/:id/actions
PUT    /api/v1/quality/qrci/:id/actions/:actionId

# 效果确认
POST   /api/v1/quality/qrci/:id/verification
```

---

### 2.5 分层审核 (Layered Process Audit / LPA)

**定位**: 标准化过程审核，逐层审核表单

```sql
-- LPA审核标准
CREATE TABLE qms_lpa_standard (
    id BIGSERIAL PRIMARY KEY,
    standard_code VARCHAR(50) UNIQUE NOT NULL,
    standard_name VARCHAR(200) NOT NULL,
    version VARCHAR(20) DEFAULT '1.0',
    dept_id BIGINT,
    dept_name VARCHAR(100),
    audit_frequency VARCHAR(20),        -- DAILY/WEEKLY/MONTHLY
    auditor_levels JSONB,               -- [{level, title, min_count}]
    question_count INTEGER,
    passing_score DECIMAL(5,2),
    is_active SMALLINT DEFAULT 1,
    effective_date DATE,
    remark VARCHAR(500),
    tenant_id BIGINT NOT NULL,
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW()
);

-- LPA审核问题项
CREATE TABLE qms_lpa_question (
    id BIGSERIAL PRIMARY KEY,
    standard_id BIGINT NOT NULL,
    question_no VARCHAR(10) NOT NULL,
    question_text VARCHAR(500) NOT NULL,
    audit_point VARCHAR(200),           -- 审核要点
    severity VARCHAR(10),               -- 严重/重要/一般
    is_critical_point SMALLINT DEFAULT 0, -- 是否为关键控制点
    sort_order INTEGER DEFAULT 0,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- LPA审核记录
CREATE TABLE qms_lpa_record (
    id BIGSERIAL PRIMARY KEY,
    record_no VARCHAR(50) UNIQUE NOT NULL,
    standard_id BIGINT NOT NULL,
    standard_name VARCHAR(200),
    auditor_id BIGINT NOT NULL,
    auditor_name VARCHAR(50),
    auditor_level VARCHAR(20),
    audit_dept_id BIGINT,
    audit_dept_name VARCHAR(100),
    audit_location VARCHAR(100),
    audit_date DATE NOT NULL,
    audit_time TIME NOT NULL,
    shift VARCHAR(20),                  -- 班次
    total_questions INTEGER,
    yes_count INTEGER DEFAULT 0,
    no_count INTEGER DEFAULT 0,
    na_count INTEGER DEFAULT 0,         -- Not Applicable
    score DECIMAL(5,2),
    result VARCHAR(10),                  -- PASS/FAIL
    findings JSONB,                     -- [{question_no, answer, remark}]
    next_action VARCHAR(500),
    status VARCHAR(20) DEFAULT 'SUBMITTED', -- SUBMITTED/VERIFIED/CLOSED
    verified_by BIGINT,
    verified_time TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

---

### 2.6 Alert告警中心

**定位**: 统一告警通知平台，整合所有子系统的告警

```sql
-- 告警规则配置
CREATE TABLE alert_rule (
    id BIGSERIAL PRIMARY KEY,
    rule_code VARCHAR(50) UNIQUE NOT NULL,
    rule_name VARCHAR(200) NOT NULL,
    alert_type VARCHAR(30) NOT NULL,     -- QUALITY/EQUIPMENT/INVENTORY/APS/SYSTEM
    biz_module VARCHAR(30),             -- 来源模块
    condition_expression VARCHAR(500) NOT NULL, -- 条件表达式
    -- 示例: oee < 60 OR defect_rate > 5
    severity_level VARCHAR(10) DEFAULT 'MEDIUM', -- HIGH/MEDIUM/LOW/INFO
    notification_channels JSONB,        -- [{channel: WEBSOCKET, config: {}}]
    notify_templates JSONB,             -- [{channel, title_template, content_template}]
    escalation_rule_id BIGINT,          -- 升级规则ID
    is_enabled SMALLINT DEFAULT 1,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 告警记录
CREATE TABLE alert_record (
    id BIGSERIAL PRIMARY KEY,
    alert_no VARCHAR(50) UNIQUE NOT NULL,
    rule_id BIGINT,
    alert_type VARCHAR(30) NOT NULL,
    severity_level VARCHAR(10) NOT NULL,
    title VARCHAR(200) NOT NULL,
    content TEXT NOT NULL,
    trigger_time TIMESTAMP NOT NULL,
    source_module VARCHAR(30),
    source_id BIGINT,                   -- 来源单据ID
    source_data JSONB,                  -- 触发时的数据快照
    status VARCHAR(20) DEFAULT 'TRIGGERED', -- TRIGGERED/ACKNOWLEDGED/RESOLVED/CLOSED
    acknowledged_by BIGINT,
    acknowledged_time TIMESTAMP,
    resolved_by BIGINT,
    resolved_time TIMESTAMP,
    resolution_remark VARCHAR(500),
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 告警通知日志
CREATE TABLE alert_notification_log (
    id BIGSERIAL PRIMARY KEY,
    alert_id BIGINT NOT NULL,
    channel VARCHAR(20) NOT NULL,      -- WEBSOCKET/FEISHU/WECHAT/SMS/EMAIL
    receiver_id BIGINT,
    receiver_name VARCHAR(50),
    receiver_value VARCHAR(200),        -- 实际联系方式
    notification_status VARCHAR(20),  -- SENT/FAILED/READ
    sent_time TIMESTAMP,
    read_time TIMESTAMP,
    error_msg VARCHAR(500),
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

**API设计**:
```
# 告警规则
GET    /api/v1/alert/rules
POST   /api/v1/alert/rules
PUT    /api/v1/alert/rules/:id
DELETE /api/v1/alert/rules/:id

# 告警记录
GET    /api/v1/alert/records
GET    /api/v1/alert/records/:id
PUT    /api/v1/alert/records/:id/acknowledge  # 确认告警
PUT    /api/v1/alert/records/:id/resolve     # 解决告警

# 告警统计
GET    /api/v1/alert/statistics               # 告警统计
GET    /api/v1/alert/statistics/by-type      # 按类型统计
GET    /api/v1/alert/statistics/trend        # 趋势分析

# WebSocket实时推送
ws://host/api/v1/ws/alert?token=xxx
```

---

### 2.7 WMS扩展: 调拨与盘点

```sql
-- 仓库调拨单
CREATE TABLE wms_transfer_order (
    id BIGSERIAL PRIMARY KEY,
    transfer_no VARCHAR(50) UNIQUE NOT NULL,
    transfer_type VARCHAR(20) NOT NULL, -- TRANSFER/ADJUSTMENT
    from_warehouse_id BIGINT NOT NULL,
    from_warehouse_name VARCHAR(100),
    to_warehouse_id BIGINT NOT NULL,
    to_warehouse_name VARCHAR(100),
    status VARCHAR(20) DEFAULT 'DRAFT',  -- DRAFT/SUBMITTED/APPROVED/IN_TRANSIT/COMPLETED/CANCELLED
    requester_id BIGINT,
    requester_name VARCHAR(50),
    approver_id BIGINT,
    approved_time TIMESTAMP,
    actual_transfer_time TIMESTAMP,
    remark VARCHAR(500),
    tenant_id BIGINT NOT NULL,
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW()
);

-- 调拨明细
CREATE TABLE wms_transfer_item (
    id BIGSERIAL PRIMARY KEY,
    transfer_order_id BIGINT NOT NULL,
    material_id BIGINT NOT NULL,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    from_location_id BIGINT,
    from_location_name VARCHAR(100),
    to_location_id BIGINT,
    to_location_name VARCHAR(100),
    transfer_qty DECIMAL(18,3) NOT NULL,
    transferred_qty DECIMAL(18,3) DEFAULT 0,
    batch_no VARCHAR(100),
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 盘点单
CREATE TABLE wms_stocktake (
    id BIGSERIAL PRIMARY KEY,
    stocktake_no VARCHAR(50) UNIQUE NOT NULL,
    stocktake_type VARCHAR(20) NOT NULL, -- FULL/CYCLE/BLIND
    warehouse_id BIGINT NOT NULL,
    warehouse_name VARCHAR(100),
    location_ids JSONB,                   -- 盘点库位列表
    status VARCHAR(20) DEFAULT 'DRAFT',   -- DRAFT/IN_PROGRESS/COMPLETED/CANCELLED
    plan_start_date DATE,
    plan_end_date DATE,
    actual_start_date DATE,
    actual_end_date DATE,
    checker_id BIGINT,
    remark VARCHAR(500),
    tenant_id BIGINT NOT NULL,
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW()
);

-- 盘点明细
CREATE TABLE wms_stocktake_item (
    id BIGSERIAL PRIMARY KEY,
    stocktake_id BIGINT NOT NULL,
    material_id BIGINT NOT NULL,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    location_id BIGINT,
    location_name VARCHAR(100),
    batch_no VARCHAR(100),
    system_qty DECIMAL(18,3),            -- 系统库存
    counted_qty DECIMAL(18,3),           -- 盘点数量
    variance_qty DECIMAL(18,3),          -- 差异数量
    variance_reason VARCHAR(200),       -- 差异原因
    is_locked SMALLINT DEFAULT 0,       -- 是否锁定
    count_time TIMESTAMP,
    counter_id BIGINT,
    counter_name VARCHAR(50),
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

---

### 2.8 APS扩展: 滚动排程与交付分析

```sql
-- APS滚动排程配置
CREATE TABLE aps_rolling_config (
    id BIGSERIAL PRIMARY KEY,
    config_code VARCHAR(50) UNIQUE NOT NULL,
    config_name VARCHAR(100) NOT NULL,
    trigger_type VARCHAR(20) NOT NULL,  -- DAILY/HOURLY/EVENT
    trigger_cron VARCHAR(50),            -- Cron表达式
    horizon_days INTEGER DEFAULT 7,       -- 排程视野
    scheduling_algorithm VARCHAR(20),    -- 使用的算法
    locked_tasks_handling VARCHAR(20),  -- IGNORE/ADJUST/RETAIN
    auto_execute SMALLINT DEFAULT 0,    -- 是否自动执行
    is_enabled SMALLINT DEFAULT 1,
    remark VARCHAR(500),
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- APS交付性能分析
CREATE TABLE aps_delivery_analysis (
    id BIGSERIAL PRIMARY KEY,
    analysis_date DATE NOT NULL,
    workshop_id BIGINT,
    total_orders INTEGER DEFAULT 0,
    on_time_orders INTEGER DEFAULT 0,
    late_orders INTEGER DEFAULT 0,
    early_orders INTEGER DEFAULT 0,
    on_time_rate DECIMAL(5,2),
    avg_delay_days DECIMAL(5,2),
    max_delay_days INTEGER,
    critical_late_orders INTEGER,       -- 紧急延期
    otd_target DECIMAL(5,2),            -- 目标准时交付率
    otd_gap DECIMAL(5,2),               -- 与目标差距
    analysis_summary JSONB,             -- 详细分析
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- APS缺料分析
CREATE TABLE aps_material_shortage (
    id BIGSERIAL PRIMARY KEY,
    analysis_no VARCHAR(50) UNIQUE NOT NULL,
    analysis_date DATE NOT NULL,
    plan_id BIGINT,
    material_id BIGINT NOT NULL,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    required_qty DECIMAL(18,3) NOT NULL,
    available_qty DECIMAL(18,3),
    shortage_qty DECIMAL(18,3),
    shortage_percentage DECIMAL(5,2),
    shortage_level VARCHAR(10),        -- CRITICAL/HIGH/MEDIUM/LOW
    alternative_sources JSONB,          -- 替代来源 [{supplier, qty, lead_time}]
    suggested_action VARCHAR(200),
    affected_orders JSONB,              -- 受影响订单
    status VARCHAR(20) DEFAULT 'ANALYZED', -- ANALYZED/ORDERED/PARTIAL/PROJECTED
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

---

### 2.9 班次管理与工厂日历(APS基础)

```sql
-- 班次模板
CREATE TABLE mdm_shift_template (
    id BIGSERIAL PRIMARY KEY,
    template_code VARCHAR(50) UNIQUE NOT NULL,
    template_name VARCHAR(100) NOT NULL,
    work_days JSONB NOT NULL,          -- [1,2,3,4,5] 周一到周五
    shifts JSONB NOT NULL,             -- [{name, start, end, break_start, break_end}]
    holiday_dates JSONB,                -- ["2026-10-01", ...]
    special_work_dates JSONB,           -- 特殊工作日
    effective_from DATE,
    effective_to DATE,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 班次明细
CREATE TABLE mdm_shift (
    id BIGSERIAL PRIMARY KEY,
    shift_code VARCHAR(50) UNIQUE NOT NULL,
    shift_name VARCHAR(100) NOT NULL,
    shift_type VARCHAR(20) NOT NULL,    -- REGULAR/OVERTIME/HOLIDAY
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    break_start TIME,
    break_end TIME,
    work_hours DECIMAL(5,2),            -- 净工作时间
    sort_order INTEGER DEFAULT 0,
    color VARCHAR(20),                  -- 甘特图颜色
    is_night_shift SMALLINT DEFAULT 0,
    remark VARCHAR(200),
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 班次日历(车间班次排班)
CREATE TABLE mdm_workshop_calendar (
    id BIGSERIAL PRIMARY KEY,
    workshop_id BIGINT NOT NULL,
    work_date DATE NOT NULL,
    shift_id BIGINT NOT NULL,
    shift_name VARCHAR(100),
    is_working_day SMALLINT DEFAULT 1,  -- 1=工作日, 0=休息日
    is_holiday SMALLINT DEFAULT 0,
    holiday_name VARCHAR(100),
    overtime_hours DECIMAL(5,2) DEFAULT 0,
    remark VARCHAR(200),
    tenant_id BIGINT NOT NULL,
    UNIQUE(workshop_id, work_date, shift_id),
    created_at TIMESTAMP DEFAULT NOW()
);
```

---

## 三、模块规划优化

### 3.1 模块依赖关系优化

```
                    ┌─────────────────────────────────────────┐
                    │           M01 系统管理                    │
                    │  (租户/用户/权限/菜单/租户配置)          │
                    └────────────────┬────────────────────────┘
                                     │
        ┌────────────────────────────┼────────────────────────────┐
        │                            │                            │
        ▼                            ▼                            ▼
┌───────────────┐         ┌──────────────────┐         ┌───────────────────┐
│  M02 主数据   │         │  M03 生产执行    │         │  M06 设备管理     │
│ (物料/BOM/    │────────▶│ (工单/报工/      │◀────────│ (设备台账/        │
│  工艺/工序)   │         │  首末件/条码)    │         │  维修/保养/点检)  │
└───────┬───────┘         └────────┬─────────┘         └─────────┬─────────┘
        │                          │                            │
        │                          ▼                            │
        │                ┌──────────────────┐                  │
        │                │  M04 APS排程     │                  │
        │                │ (工作中心/日历/   │◀─────────────────┘
        │                │  换型矩阵/JIT)   │
        │                └────────┬─────────┘
        │                         │
        │         ┌───────────────┼───────────────┐
        │         │               │               │
        │         ▼               ▼               ▼
        │  ┌────────────┐  ┌────────────┐  ┌─────────────┐
        │  │ M07 仓库   │  │ M05 质量   │  │ M08 数据采集│
        │  │ (库存/拉动)│  │ (IQC/IPQC/ │  │ (OPC-UA/   │
        │  │           │  │  FQC/OQC)  │  │  MQTT/条码)│
        │  └─────┬──────┘  └─────┬──────┘  └──────┬───────┘
        │        │               │               │
        │        └───────────────┼───────────────┘
        │                         │
        │                         ▼
        │               ┌──────────────────┐
        │               │  M09 Andon       │
        │               │  (告警/升级/广播)│
        │               └────────┬─────────┘
        │                          │
        │         ┌────────────────┼────────────────┐
        │         │                │                │
        │         ▼                ▼                ▼
        │  ┌────────────┐  ┌────────────┐  ┌─────────────┐
        │  │ M10 追溯   │  │M11 实验室  │  │M12 器具/量检│
        │  │(批次追溯/  │  │(SPC/检测申 │  │具管理       │
        │  │ 工序追溯)  │  │请/仪器校准)│  │             │
        │  └────────────┘  └────────────┘  └─────────────┘
        │                          │
        │                          ▼
        │               ┌──────────────────┐
        │               │ M13 AI质检       │
        │               │ (视觉检测对接)    │
        │               └────────┬─────────┘
        │                          │
        └──────────────────────────┼──────────────────────────┐
                                   ▼                           │
                        ┌──────────────────┐                   │
                        │  M15 报表与看板  │                   │
                        │ (生产/质量/设备/ │                   │
                        │  交付/Andon)     │                   │
                        └──────────────────┘                   │
                                                               │
                        ┌───────────────────────────────────────┘
                        │
                        ▼
               ┌──────────────────┐
               │ M14 系统集成     │
               │ (ERP/AGV/视觉/  │
               │  飞书/供应商Portal)
               └──────────────────┘
```

### 3.2 新增全局模块

| 模块 | 说明 | 优先级 | 依赖 |
|------|------|--------|------|
| **Alert Center** | 统一告警中心 | P1 | M05/M06/M07/M09 |
| **e-SOP** | 电子作业指导书 | P1 | M02/M03 |
| **QRCI** | 质量闭环管理 | P2 | M05 |
| **LPA** | 分层过程审核 | P2 | M05 |

### 3.3 模块完成度更新

| 模块ID | 模块名称 | 原完成度 | 新完成度 | 需补充内容 |
|--------|---------|---------|---------|-----------|
| M01 | 系统管理 | 75% | 80% | 多车间配置、打印模板 |
| M02 | 主数据管理 | 70% | 75% | 客户管理、批量导入 |
| M03 | 生产执行 | 35% | 60% | +e-SOP、派工单、工单变更 |
| M04 | AI-APS排程 | 25% | 70% | +滚动排程、交付分析、缺料分析 |
| M05 | 质量管理 | 75% | 85% | +QRCI、LPA、检验标准关联 |
| M06 | 设备管理 | 70% | 85% | +TEEP、量检具管理 |
| M07 | 仓库管理 | 70% | 85% | +调拨、盘点、线边库位 |
| M08 | 数据采集 | 0% | 60% | 核心功能已设计 |
| M09 | Andon系统 | 40% | 75% | +升级规则、通知日志 |
| M10 | 追溯管理 | 30% | 50% | 待与生产执行联动 |
| M11 | 实验室/SPC | 0% | 70% | +检测申请、仪器管理 |
| M12 | 器具管理 | 0% | 80% | 量检具+器具 |
| M13 | AI质检 | 0% | 50% | 视觉对接方案已有 |
| M14 | 系统集成 | 0% | 40% | ERP对接方案已有 |
| M15 | 报表看板 | 20% | 30% | 需补充低代码平台设计 |

---

## 四、文档维护建议

### 4.1 文档结构优化

```
docs/
├── MOM3.0_架构与主文档.md          # 索引 + 全局规范
├── MOM3.0_基础数据设计文档.md      # M01+M02
├── MOM3.0_MES生产执行设计文档.md    # M03+M04(APS)+M06(设备OEE)
├── MOM3.0_WMS仓储设计文档.md       # M07+M08
├── MOM3.0_QMS质量设计文档.md       # M05+M09(Andon)+M10(追溯)
├── MOM3.0_高级质量模块设计.md       # M11(实验室)+M12(器具)+QRCI+LPA (新增)
├── MOM3.0_AI质检与APS设计文档.md   # M13(AI)+M04高级(AI-APS) (新增)
├── MOM3.0_系统集成设计文档.md       # M14(ERP/AGV/飞书/供应商) (新增)
├── MOM3.0_报表与看板设计文档.md     # M15 (新增，原内容打散)
├── MOM3.0_Alert告警中心设计.md      # 统一告警 (新增)
├── MOM3.0_UI设计规范.md
└── MOM3.0_安灯系统设计文档.md      # M09 (独立)
```

### 4.2 需更新的文档

1. **MOM3.0_架构与主文档.md** - 更新模块完成度、更新API清单
2. **MOM3.0_MES生产执行设计文档.md** - 补充e-SOP设计、派工单设计
3. **MOM3.0_QMS质量设计文档.md** - 补充QRCI、LPA、Alert集成
4. **新增文档** - 按上表结构新建6个文档

---

## 五、总结

### 5.1 需立即补充的设计

| 优先级 | 内容 |
|-------|------|
| **P0** | e-SOP电子作业指导书(影响M03生产执行闭环) |
| **P0** | APS滚动排程+交付分析+缺料分析 |
| **P1** | Alert统一告警中心 |
| **P1** | 量检具管理(设备模块扩展) |
| **P1** | 调拨+盘点(WMS扩展) |
| **P1** | M11实验室(检测申请+仪器管理) |

### 5.2 文档工作量估算

| 文档 | 工时 |
|------|------|
| 新增e-SOP设计 | 0.5天 |
| 新增实验室+仪器+量检具设计 | 1天 |
| 新增QRCI+LPA设计 | 0.5天 |
| 新增Alert告警中心设计 | 0.5天 |
| 新增系统集成设计 | 1天 |
| 新增报表平台设计 | 1天 |
| 文档结构重组+更新主文档 | 0.5天 |
| **合计** | **6天** |

---

*文档版本: V1.1 | 更新日期: 2026-04-09 | 更新内容: 模块规划优化与缺失功能补充*
