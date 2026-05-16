# MOM3.0 准时化生产与发运(JIT/JIS)设计

**版本**: V1.0 | **日期**: 2026-05-16 | **基于**: SAP S/4HANA Next Generation JIT

---

## 一、核心概念对照

| SAP概念 | 说明 | MOM对应 |
|--------|------|---------|
| JIT Control Cycle | 控制循环，定义物料补充策略 | 物料供应策略 |
| PSA (Production Supply Area) | 生产供应区，线边暂存区 | 线边仓/供应点 |
| JIT Call | 物料拉动请求 | 拉动单 |
| JIS Call | 按序配送请求 | 排序发货单 |
| Sequencing | 车辆/零件排序 | 排序管理 |
| EDI Parser | 报文解析 | 准时化报文解析 |

---

## 二、功能模块设计

### 2.1 JIT控制循环管理

```
JIT控制循环 (JIT Control Cycle)
├── 控制循环主数据
│   ├── 物料编码/名称
│   ├── 供应策略 (JIT/JIS/VMI)
│   ├── 最小补充量
│   ├── 最大库存量
│   ├── 目标库存量
│   ├── 提前期 (Lead Time)
│   └── 供应间隔
├── 供应点配置
│   ├── PSA编号/名称
│   ├── 库位编码
│   ├── 接收工位
│   └── 供应商/物流商
└── 触发规则
    ├── 安全库存触发
    ├── 节拍触发 (Kanban)
    ├── 工单触发
    └── 手工触发
```

### 2.2 JIT Call管理

```
JIT Call管理
├── JIT Call创建
│   ├── 基于控制循环自动创建
│   ├── 基于工单手动创建
│   └── 基于日计划自动创建
├── JIT Call状态
│   ├── 待确认 (PENDING)
│   ├── 已确认 (CONFIRMED)
│   ├── 发货中 (IN_TRANSIT)
│   ├── 已收货 (RECEIVED)
│   └── 已完成 (COMPLETED)
├── JIT Call内容
│   ├── 物料编码/数量
│   ├── 交货时间窗口
│   ├── 目的地PSA
│   ├── 供应商/物流商
│   └── 参考工单/日计划
└── 优先级管理
    ├── 紧急/加急标记
    ├── 排序号
    └── 取消/变更
```

### 2.3 JIS Call管理

```
JIS Call管理 (Just-In-Sequence)
├── JIS Call创建
│   ├── 基于车辆序列号创建
│   ├── 基于生产顺序创建
│   └── 上游转发创建
├── 序列信息
│   ├── 车辆VIN/序列号
│   ├── 工位顺序号
│   ├── 车型/配置
│   └── 颜色/选装
├── 零件明细
│   ├── 零件编号
│   ├── 排序位置
│   ├── 配送顺序
│   └── 配套关系
└── 转发管理
    ├── 转发给供应商
    ├── 转发给物流商
    └── 重新排序
```

### 2.4 生产供应区(PSA)管理

```
PSA管理
├── PSA主数据
│   ├── PSA编码/名称
│   ├── 类型 (线边仓/排序区)
│   ├── 对应产线/工位
│   └── 容积限制
├── 库位配置
│   ├── 库位列表
│   ├── 库位类型
│   └── 库存上限
├── 补货策略
│   ├── 补货方式 (JIT/JIS/KIT)
│   ├── 触发点
│   └── 补充量计算
└── 看板配置
    ├── 看板号
    ├── 看板容量
    └── 补充路线
```

### 2.5 发货排序管理

```
发货排序管理 (Dispatch Sequencing)
├── 车辆排序
│   ├── VIN/序列号
│   ├── 下线时间
│   ├── 颜色/配置
│   └── 目的仓库/经销商
├── 排序监控
│   ├── 当前排序状态
│   ├── 完成率
│   ├── 异常车辆
│   └── 时间窗口监控
├── 发运计划
│   ├── 发运批次
│   ├── 车辆清单
│   ├── 预计发运时间
│   └── 物流窗口
└── 排序工厂布局
    ├── 工位映射
    ├── 产线监控
    └── 完成进度
```

### 2.6 物料监控预警

```
物料监控预警
├── JIT Call监控
│   ├── 逾期预警
│   ├── 数量偏差预警
│   ├── 质量异常预警
│   └── 供应商绩效
├── PSA库存监控
│   ├── 最高库存预警
│   ├── 安全库存预警
│   ├── 呆滞预警
│   └── 库存准确率
├── 发货监控
│   ├── 延误预警
│   ├── 错漏件预警
│   └── 车辆滞留预警
└── 绩效报表
    ├── 供应及时率
    ├── 供应准确率
    ├── 供应商KPI
    └── 产线停线分析
```

### 2.7 通信组管理

```
通信组管理
├── 供应商管理
│   ├── 供应商编码/名称
│   ├── 联系人
│   ├── EDI配置
│   └── 供应范围
├── 物流商管理
│   ├── 物流商编码/名称
│   ├── 联系人
│   ├── 车辆信息
│   └── 窗口时间配置
└── 通知配置
    ├── JIT Call发送
    ├── 变更通知
    ├── 逾期通知
    └── 完成确认
```

### 2.8 准时化报文解析模块

#### 2.8.1 报文格式支持

| 格式 | 说明 | 应用场景 |
|------|------|---------|
| **VDA 4905** | 材料拉动请求 | 大众集团 |
| **VDA 4906** | 交货计划 | 大众集团 |
| **VDA 4915** | 发货通知 | 大众集团 |
| **IDOC ORDERS** | 采购订单 | SAP标准 |
| **IDOC DESADV** | 发货通知 | SAP标准 |
| **ANSI X12 868** | 计划交互 | 北美 |
| **XML (VDA Set)** | 现代EDI | 通用 |

#### 2.8.2 报文处理流程

```
EDI报文接收
    │
    ▼
┌─────────────────┐
│  报文安全验证   │ ← 签名/加密校验
└─────────────────┘
    │
    ▼
┌─────────────────┐
│  报文格式识别   │ ← 自动识别VDA/IDOC/XML
└─────────────────┘
    │
    ▼
┌─────────────────┐
│  报文解析引擎   │ ← 字段提取、格式转换
└─────────────────┘
    │
    ▼
┌─────────────────┐
│  业务数据转换   │ ← 转换为内部数据模型
└─────────────────┘
    │
    ▼
┌─────────────────┐
│  数据校验       │ ← 必填/格式/业务规则
└─────────────────┘
    │
    ▼
┌─────────────────┐
│  创建JIT/JIS    │ ← 生成系统单据
└─────────────────┘
    │
    ▼
┌─────────────────┐
│  回执发送       │ ← 确认ACK/错误NACK
└─────────────────┘
```

#### 2.8.3 VDA 4905 解析规格

```
VDA 4905 大众材料拉动报文结构：
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
段号   内容                    字段
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
KNVGT  发货方                   供应商编号
KNDT   接收方                   主机厂编号
VKOND  合同条件                 协议号
ABLFD  开始日期
ABLFD  结束日期
ABLNR  顺序号
POSNR  行项目号
MATNR  物料编号
MENGE  数量
MEINS  单位
ENDDA  截止日期
ENDEU  截止时间
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```

#### 2.8.4 IDOC ORDERS 解析规格

```
IDOC ORDERS 采购订单结构：
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
段号          内容               字段
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
E1EDK010     Header             订单编号/日期
E1EDK14      合作伙伴           供应商信息
E1EDKT1      描述段             描述信息
E1EDP01      行项目            行项目号
E1EDP19      数量              订单数量
E1EDPA1      交货日期          计划日期
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```

---

## 三、数据模型

### 3.1 JIT核心表

```sql
-- JIT控制循环
CREATE TABLE jit_control_cycle (
    id BIGSERIAL PRIMARY KEY,
    cycle_code VARCHAR(50) UNIQUE NOT NULL,
    cycle_name VARCHAR(100),
    material_id BIGINT,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    supply_strategy VARCHAR(20),  -- JIT/JIS/VMI
    psa_id BIGINT,
    supplier_id BIGINT,
    min_qty DECIMAL(18,6),
    max_qty DECIMAL(18,6),
    target_qty DECIMAL(18,6),
    lead_time_minutes INT,
    trigger_type VARCHAR(20),  -- AUTO/MANUAL/KANBAN
    status VARCHAR(20),
    tenant_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- PSA生产供应区
CREATE TABLE psa_master (
    id BIGSERIAL PRIMARY KEY,
    psa_code VARCHAR(50) UNIQUE NOT NULL,
    psa_name VARCHAR(100),
    psa_type VARCHAR(20),  -- LINE_SIDE/SORTING/KIT
    line_id BIGINT,
    location_id BIGINT,
    max_capacity INT,
    status VARCHAR(20),
    tenant_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- JIT Call
CREATE TABLE jit_call (
    id BIGSERIAL PRIMARY KEY,
    call_code VARCHAR(50) UNIQUE NOT NULL,
    call_type VARCHAR(20),  -- JIT/JIS
    cycle_id BIGINT,
    material_id BIGINT,
    request_qty DECIMAL(18,6),
    delivered_qty DECIMAL(18,6),
    request_time TIMESTAMP,
    delivery_window_start TIMESTAMP,
    delivery_window_end TIMESTAMP,
    psa_id BIGINT,
    supplier_id BIGINT,
    logistics_id BIGINT,
    priority INT DEFAULT 0,
    status VARCHAR(20),  -- PENDING/CONFIRMED/IN_TRANSIT/RECEIVED/COMPLETED/CANCELLED
    reference_order_id BIGINT,
    reference_day_plan_id BIGINT,
    tenant_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- JIS Call
CREATE TABLE jis_call (
    id BIGSERIAL PRIMARY KEY,
    call_code VARCHAR(50) UNIQUE NOT NULL,
    cycle_id BIGINT,
    vehicle_vin VARCHAR(50),
    sequence_no INT,
    model_code VARCHAR(50),
    color_code VARCHAR(20),
    production_date DATE,
    psa_id BIGINT,
    supplier_id BIGINT,
    status VARCHAR(20),  -- PENDING/SENT/CONFIRMED/DELIVERED
    source_call_id BIGINT,  -- 上游转发
    tenant_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 车辆排序
CREATE TABLE vehicle_sequence (
    id BIGSERIAL PRIMARY KEY,
    vin VARCHAR(50) UNIQUE NOT NULL,
    sequence_no INT,
    line_id BIGINT,
    plan_date DATE,
    model_code VARCHAR(50),
    color_code VARCHAR(20),
    production_time TIMESTAMP,
    dispatch_status VARCHAR(20),  -- WAITING/IN_PRODUCTION/COMPLETED/DISPATCHED
    target_warehouse VARCHAR(50),
    estimated_dispatch_time TIMESTAMP,
    actual_dispatch_time TIMESTAMP,
    tenant_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 发运计划
CREATE TABLE dispatch_plan (
    id BIGSERIAL PRIMARY KEY,
    plan_code VARCHAR(50) UNIQUE NOT NULL,
    plan_date DATE,
    batch_no VARCHAR(20),
    warehouse_id BIGINT,
    vehicle_count INT,
    planned_time TIMESTAMP,
    status VARCHAR(20),
    tenant_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 通信组
CREATE TABLE communication_group (
    id BIGSERIAL PRIMARY KEY,
    group_code VARCHAR(50) UNIQUE NOT NULL,
    group_name VARCHAR(100),
    group_type VARCHAR(20),  -- SUPPLIER/LOGISTICS
    partner_id BIGINT,
    contact_person VARCHAR(50),
    contact_phone VARCHAR(20),
    contact_email VARCHAR(100),
    edi_config JSONB,
    tenant_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW()
);
```

### 3.2 报文解析表

```sql
-- EDI报文日志
CREATE TABLE edi_message_log (
    id BIGSERIAL PRIMARY KEY,
    message_id VARCHAR(100) UNIQUE,
    message_type VARCHAR(30),      -- VDA4905/VDA4906/VDA4915/ORDERS/DESADV/X12
    direction VARCHAR(10),           -- IN/OUT
    sender VARCHAR(50),
    receiver VARCHAR(50),
    raw_content TEXT,               -- 原始报文
    parse_status VARCHAR(20),      -- SUCCESS/PARSING_ERROR/VALIDATION_ERROR
    error_message TEXT,
    processed_data JSONB,
    tenant_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 报文解析配置
CREATE TABLE edi_parser_config (
    id BIGSERIAL PRIMARY KEY,
    config_code VARCHAR(50) UNIQUE,
    config_name VARCHAR(100),
    message_type VARCHAR(30),
    partner_id BIGINT,             -- 供应商/主机厂
    version VARCHAR(20),           -- 报文版本
    field_mapping JSONB,            -- 字段映射
    validation_rules JSONB,        -- 校验规则
    transformation_rules JSONB,     -- 转换规则
    status VARCHAR(20),
    tenant_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- JIT Call映射表
CREATE TABLE jit_call_mapping (
    id BIGSERIAL PRIMARY KEY,
    edi_message_id BIGINT,
    jit_call_id BIGINT,
    mapping_status VARCHAR(20),
    tenant_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- EDI合作伙伴
CREATE TABLE edi_partner (
    id BIGSERIAL PRIMARY KEY,
    partner_code VARCHAR(50) UNIQUE,
    partner_name VARCHAR(100),
    partner_type VARCHAR(20),      -- SUPPLIER/CUSTOMER/LOGISTICS
    industry VARCHAR(20),          -- AUTOMOTIVE/OTHER
    edi_standard VARCHAR(20),       -- VDA/IDOC/X12/XML
    connection_config JSONB,        -- 连接配置
    security_config JSONB,         -- 安全配置
    status VARCHAR(20),
    tenant_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW()
);
```

---

## 四、API接口设计

### 4.1 JIT管理接口

| 接口 | 方法 | 说明 |
|------|------|------|
| `/api/v1/jit/cycles` | GET | 查询控制循环列表 |
| `/api/v1/jit/cycles` | POST | 创建控制循环 |
| `/api/v1/jit/cycles/:id` | PUT | 更新控制循环 |
| `/api/v1/jit/cycles/:id` | DELETE | 删除控制循环 |
| `/api/v1/jit/calls` | GET | 查询JIT Call列表 |
| `/api/v1/jit/calls` | POST | 创建JIT Call |
| `/api/v1/jit/calls/:id` | PUT | 更新JIT Call |
| `/api/v1/jit/calls/:id/cancel` | POST | 取消JIT Call |
| `/api/v1/jit/calls/:id/confirm` | POST | 确认JIT Call |

### 4.2 JIS管理接口

| 接口 | 方法 | 说明 |
|------|------|------|
| `/api/v1/jis/calls` | GET | 查询JIS Call列表 |
| `/api/v1/jis/calls` | POST | 创建JIS Call |
| `/api/v1/jis/calls/:id` | PUT | 更新JIS Call |
| `/api/v1/jis/calls/:id/forward` | POST | 转发JIS Call |
| `/api/v1/jis/calls/:id/reorder` | POST | 重新排序 |

### 4.3 PSA管理接口

| 接口 | 方法 | 说明 |
|------|------|------|
| `/api/v1/psa` | GET | 查询PSA列表 |
| `/api/v1/psa` | POST | 创建PSA |
| `/api/v1/psa/:id` | PUT | 更新PSA |
| `/api/v1/psa/:id/stock` | GET | 查询PSA库存 |

### 4.4 排序管理接口

| 接口 | 方法 | 说明 |
|------|------|------|
| `/api/v1/sequence/vehicles` | GET | 查询车辆排序 |
| `/api/v1/sequence/vehicles` | POST | 创建车辆排序 |
| `/api/v1/sequence/monitor` | GET | 产线排序监控 |
| `/api/v1/dispatch/plans` | GET | 查询发运计划 |
| `/api/v1/dispatch/plans` | POST | 创建发运计划 |

### 4.5 EDI报文接口

| 接口 | 方法 | 说明 |
|------|------|------|
| `/api/v1/edi/inbound` | POST | 接收EDI报文 |
| `/api/v1/edi/outbound` | POST | 发送EDI报文 |
| `/api/v1/edi/logs` | GET | 报文日志查询 |
| `/api/v1/edi/config` | GET/POST | 报文配置 |
| `/api/v1/edi/config/:id/test` | POST | 配置测试 |
| `/api/v1/edi/convert` | POST | 格式转换 |
| `/api/v1/edi/partners` | GET/POST | 合作伙伴管理 |

---

## 五、菜单规划

```
生产执行 (Production Execution)
├── 生产订单 (已有)
├── 日计划 (已有)
├── 物料供应
│   ├── JIT控制循环
│   ├── JIT Call管理
│   ├── JIS Call管理
│   └── PSA管理
├── 排序管理
│   ├── 车辆排序
│   ├── 产线监控
│   └── 发运计划
├── 物料拉动
│   ├── 拉动请求
│   ├── 拉动确认
│   └── 库存监控
├── 通信组
│   ├── 供应商管理
│   └── 物流商管理
└── EDI报文
    ├── 报文配置
    ├── 报文日志
    └── 合作伙伴
```

---

## 六、错误处理

| 错误码 | 说明 | 处理 |
|--------|------|------|
| E001 | 报文格式错误 | 记录日志，返NACK |
| E002 | 字段缺失 | 记录日志，返NACK |
| E003 | 物料不存在 | 自动创建或报错 |
| E004 | 数量超限 | 警告+人工确认 |
| E005 | 供应商未知 | 记录日志，返NACK |
| E006 | 解析超时 | 记录日志，返NACK |
| E007 | 校验失败 | 记录日志，返NACK |
| E008 | 转换失败 | 记录日志，返NACK |

---

## 七、参考资料

| 来源 | 内容 |
|------|------|
| [SAP Help Portal - Next Generation JIT](https://help.sap.com/docs/SAP_S4HANA_ON-PREMISE/9832125c23154a179bfa1784cdc9577a/2e7f076b73554959a185b8adba16cb3e.html) | SAP S/4HANA Next Generation JIT/JIS完整功能列表 |
| [SAP Community - JIT Processing](https://community.sap.com/t5/supply-chain-management-blog-posts-by-sap/jit-processing-with-the-next-generation-just-in-time-in-s-4hana-part-1/ba-p/13611311) | JIT处理流程说明 |
| [VDA - German Automotive Association](https://www.vda.de/) | VDA报文标准 |

---

**文档状态**: 初稿，待评审

**使用的Skill**: `sap-design`