# MOM3.0 开发设计文档 — 架构总览

**版本**: V1.0 | **基准文档**: 闻荫科技V6.0 | **项目甲方**: 上海闻荫科技系统有限公司

---

> **说明**: 本文档为MOM3.0项目的架构总览文档，定义统一的技术架构、代码规范，并索引所有模块详细设计文档。所有开发工作均需遵循本文档定义的标准。

---

## 模块设计文档索引

| 文档名称 | 涵盖模块 | 说明 |
|---------|---------|------|
| `MOM3.0_基础数据设计文档.md` | M01系统管理、M02主数据管理 | 用户、角色、权限、菜单、物料、BOM、工艺路线等 |
| `MOM3.0_MES生产执行设计文档.md` | M03生产执行、M04 APS、M06设备管理 | 工单、报工、APS排程、甘特图、设备OEE等 |
| `MOM3.0_eSOP电子作业指导书设计.md` | M03生产执行扩展 | 工艺文件电子化、SOP推送、工位终端、电子签名 |
| `MOM3.0_WMS仓储设计文档.md` | M07仓库与物料、M08数据采集 | 入库、出库、库存、拉动、数据采集等 |
| `MOM3.0_QMS质量设计文档.md` | M05质量管理、M10追溯 | IQC/IPQC/FQC/OQC、追溯等 |
| `MOM3.0_高级质量模块设计.md` | M05扩展（QRCI/LPA）、M11实验室、M12量检具 | QRCI质量闭环、LPA分层审核、检测申请、仪器校准 |
| `MOM3.0_M09安灯系统设计文档.md` | M09安灯系统 | 呼叫、响应、升级机制、广播、大屏、统计分析完整设计 |
| `MOM3.0_Alert告警中心设计文档.md` | 全局告警模块 | 统一告警规则、通知渠道、升级机制、统计分析 |
| `MOM3.0_系统集成设计文档.md` | M14系统集成 | ERP对接、AGV调度、飞书/企微通知、供应商Portal |
| `MOM3.0_报表与看板设计文档.md` | M15报表与大屏 | 标准报表、低代码配置、车间大屏、WebSocket实时推送 |
| `MOM3.0_UI设计规范.md` | 全局UI规范 | 色彩、字体、间距、组件、页面布局、状态设计等 |

---

## 目录

1. [系统总览](#1-系统总览)
2. [技术架构](#2-技术架构)
3. [统一代码规范](#3-统一代码规范)
4. [模块开发状态与优先级](#4-模块开发状态与优先级)
5. [数据库完整DDL](#5-数据库完整ddl)
6. [API接口全清单](#6-api接口全清单)
7. [前端页面路由清单](#7-前端页面路由清单)
8. [开发排期与里程碑](#8-开发排期与里程碑)
9. [非功能性需求](#9-非功能性需求)

---

## 1. 系统总览

### 1.1 产品定位

MOM3.0（Manufacturing Operations Management）是面向**汽车零部件离散制造业**的生产运营管理系统，集成MES执行层与APS计划层，为闻荫科技三个车间提供统一管控平台。

**三个车间**:
- 平衡轴车间（2条产线，机加工+磨削+焊接+铆接）
- 电池包车间（2条产线，装配+气密+清洗+涂胶）
- 试制车间（≤10台设备，小批量试制）

### 1.2 核心工艺流程

```
原材料入库(IQC) → 上料(AGV/人工) → 机加工/磨削 → 焊接/铆接/涂胶
→ 清洗 → 视觉检测 → 装配 → 气密测试 → 打印条码 → 装箱 → OQC → 出库
```

### 1.3 系统架构图

```
┌─────────────────────────────────────────────────────────────────┐
│                    用户交互层                                      │
│  Web浏览器  │  移动端APP  │  车间大屏  │  工位终端  │  供应商Portal │
└──────────────────────────┬──────────────────────────────────────┘
                           │ HTTPS / WebSocket
┌──────────────────────────▼──────────────────────────────────────┐
│                    API网关 (Nginx)                                 │
│  认证鉴权 │ 路由分发 │ 限流熔断 │ 日志审计                          │
└──────┬───────────┬───────────┬───────────┬────────────┬──────────┘
       │           │           │           │            │
  ┌────▼───┐  ┌───▼────┐  ┌──▼─────┐  ┌──▼─────┐  ┌──▼──────┐
  │MES服务 │  │APS服务 │  │QMS服务 │  │EAM服务 │  │AI质检服务│
  │:8080   │  │:8081   │  │:8082   │  │:8083   │  │:8084    │
  └────┬───┘  └───┬────┘  └──┬─────┘  └──┬─────┘  └──┬──────┘
       └──────────┴───────────┴───────────┴────────────┘
                              │
┌─────────────────────────────▼──────────────────────────────────┐
│                    数据存储层                                      │
│  PostgreSQL 15  │  Redis 7  │  Kafka 3  │  MinIO  │  Milvus    │
└─────────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────▼──────────────────────────────────┐
│                    外部系统集成层                                  │
│  金蝶ERP  │  AGV系统  │  视觉检测  │  激光打码  │  飞书/企微      │
│  供应商Portal │  主机厂JIT  │  蓝牙量具  │  OPC-UA设备          │
└─────────────────────────────────────────────────────────────────┘
```

### 1.4 模块总览与当前完成度

| 模块ID | 模块名称 | 当前完成度 | P0目标完成度 | 新增功能 |
|--------|---------|-----------|------------|---------|
| M01 | 系统管理 | 80% | 100% | 多车间管理、打印模板、系统参数 |
| M02 | 主数据管理 | 75% | 100% | 批量导入、客户管理 |
| M03 | 生产执行 | 60% | 90% | +e-SOP电子作业指导书、首末件、流程卡、序列号、批次码 |
| M04 | AI-APS排程 | 70% | 85% | +滚动排程、交付分析、缺料分析、换型矩阵感知 |
| M05 | 质量管理 | 85% | 90% | +QRCI质量闭环、LPA分层审核、检验标准关联 |
| M06 | 设备管理 | 85% | 85% | TEEP、模具、量检具管理 |
| M07 | 仓库管理 | 85% | 90% | JIT拉动、调拨管理、盘点管理 |
| M08 | 数据采集 | 60% | 80% | OPC-UA/MQTT数据采集核心功能 |
| M09 | Andon系统 | 75% | 90% | 升级规则、通知日志、广播集成 |
| M10 | 追溯管理 | 50% | 85% | 精确批次、激光打码对接 |
| M11 | 实验室/SPC | 70% | 70% | 检测申请、样品管理、仪器校准 |
| M12 | 器具/量检具 | 80% | 80% | 量检具台账、借用、校准记录 |
| M13 | AI质检 | 50% | 60% | 视觉检测对接方案 |
| M14 | 系统集成 | 70% | 70% | ERP对接、AGV调度、飞书/企微通知 |
| M15 | 报表看板 | 30% | 75% | 低代码报表平台、车间大屏 |

---

## 2. 技术架构

### 2.1 技术栈

#### 后端

| 组件 | 版本 | 用途 |
|------|------|------|
| Go | 1.23+ | 主开发语言 |
| Gin | 1.9+ | HTTP Web框架 |
| GORM | 1.25+ | ORM框架 |
| PostgreSQL | 15.x | 主数据库 |
| Redis | 7.x | 缓存 / 分布式锁 / 消息队列 |
| Kafka | 3.x | 异步消息 / 设备数据流 |
| gorilla/websocket | 1.5+ | WebSocket推送 |
| go-cron | - | 定时任务 |
| zap | - | 结构化日志 |

#### 前端

| 组件 | 版本 | 用途 |
|------|------|------|
| Vue 3 | 3.x | 前端框架 |
| Element Plus | 2.x | UI组件库 |
| Pinia | 2.x | 状态管理 |
| Vite | 5.x | 构建工具 |
| ECharts | 6.x | 图表/看板 |
| dhtmlx-gantt | 9.x | APS甘特图 |
| socket.io-client | 4.x | WebSocket |

#### AI服务（M13 AI质检 / M04 AI-APS）

| 组件 | 版本 | 用途 |
|------|------|------|
| Python | 3.10+ | AI服务语言 |
| PyTorch | 2.0+ | 深度学习框架 |
| TensorRT | 10.x | GPU推理加速 |
| OR-Tools | - | APS约束规划 |
| scikit-learn | - | 质量预测 |
| OpenCV | 4.8+ | 图像处理 |

### 2.2 项目目录结构

```
mom3.0/
├── mom-api/                          # Go后端
│   ├── cmd/main.go
│   └── internal/
│       ├── config/                   # 配置（多环境）
│       ├── middleware/               # JWT认证/RBAC/按钮权限/日志
│       ├── handlers/                 # HTTP处理器（按模块分目录）
│       │   ├── system/               # M01
│       │   ├── mdm/                  # M02
│       │   ├── production/           # M03
│       │   ├── aps/                  # M04 AI-APS
│       │   ├── quality/              # M05
│       │   ├── equipment/            # M06
│       │   ├── warehouse/            # M07
│       │   ├── datacollect/          # M08
│       │   ├── andon/                # M09
│       │   ├── trace/                # M10
│       │   ├── laboratory/           # M11 新增
│       │   ├── container/            # M12 新增
│       │   ├── erp/                  # M14 ERP对接
│       │   ├── agv/                  # M14 AGV对接
│       │   └── report/               # M15
│       ├── models/                   # GORM数据模型
│       ├── service/                  # 业务逻辑
│       │   ├── aps/
│       │   │   ├── scheduler.go      # 排程引擎
│       │   │   ├── algorithm/        # 算法库
│       │   │   └── ai_optimizer.go   # AI优化器
│       │   ├── spc/                  # SPC计算
│       │   ├── erp_sync.go           # ERP同步
│       │   └── agv.go                # AGV调度
│       ├── repository/               # 数据访问层
│       ├── router/                   # 路由注册
│       ├── scheduler/                # 定时任务
│       ├── cache/                    # Redis封装
│       └── websocket/                # WS Hub
│
├── mom-web/                          # Vue前端
│   └── src/
│       ├── api/                      # axios接口层
│       ├── components/               # 公共组件
│       │   ├── PermissionButton.vue  # 按钮权限
│       │   ├── GanttChart.vue        # 甘特图
│       │   └── SpcChart.vue          # SPC控制图
│       ├── router/                   # vue-router
│       ├── store/                    # Pinia stores
│       └── views/                    # 页面
│           ├── system/
│           ├── mdm/
│           ├── production/
│           ├── aps/                  # AI-APS页面
│           ├── quality/
│           ├── equipment/
│           ├── warehouse/
│           ├── andon/
│           ├── trace/
│           ├── laboratory/           # 新增
│           ├── container/            # 新增
│           └── report/
│
└── mom-ai/                           # Python AI服务
    ├── aps_optimizer/                # APS AI优化
    ├── vision_inspect/               # 视觉质检
    └── spc_analyzer/                 # SPC分析
```

### 2.3 通用响应格式

```json
{
  "code": 200,
  "message": "success",
  "data": {},
  "timestamp": "2026-04-01T10:00:00Z"
}
```

分页响应:
```json
{
  "code": 200,
  "data": {
    "list": [],
    "total": 100,
    "page": 1,
    "page_size": 20
  }
}
```

### 2.4 通用BaseModel（所有表继承）

```go
// models/base.go
type BaseModel struct {
    ID          int64          `gorm:"primaryKey;autoIncrement" json:"id"`
    TenantID    int64          `gorm:"not null;index" json:"tenant_id"`         // 多车间隔离
    WorkshopID  int64          `gorm:"index" json:"workshop_id"`                // 车间
    CreatedBy   string         `gorm:"size:50" json:"created_by"`
    UpdatedBy   string         `gorm:"size:50" json:"updated_by"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`                          // 软删除
}
```

---

## 3. 统一代码规范

### 3.1 后端 (Go/Gin)

- **分层架构**: Handler → Service → Repository
- **错误处理**: 所有API handler必须用panic-recover捕获异常，service层返回error
- **日志**: 使用zap日志库，记录操作前后状态
- **权限**: JWT + 租户隔离 + 按钮级权限(perms字段)
- **API响应格式**: 统一{code, data, message}格式
- **参数校验**: 所有输入参数必须校验，使用validator库
- **事务**: 多表操作必须使用事务
- **SQL**: 必须使用GORM参数化查询，防止SQL注入

### 3.2 前端 (Vue3/Element Plus)

- **按钮级权限控制**: 所有页面必须添加 `v-if="hasPermission('module:page:action')"`
- **状态管理**: Pinia stores
- **async/await + try-catch**: 所有API调用必须包在try-catch中
- **表单验证**: el-form + FormRules
- **分页**: el-pagination + pagination对象
- **弹窗编辑**: el-dialog + formData对象
- **组件命名**: PascalCase组件，camelCase工具函数
- **样式**: 优先使用Tailwind CSS或Element Plus主题变量，避免内联样式

---

## 4. 模块开发状态与优先级

### 4.1 P0 - 7月15日前必须完成（上线阻塞项）

| 序号 | 功能 | 模块 | 预估工时 | 负责方向 |
|------|------|------|---------|---------|
| 1 | 生产执行闭环（工单/报工/序列号/批次号） | M03 | 15天 | 后端+前端 |
| 2 | AI-APS智能排程（核心算法+甘特图） | M04 | 20天 | 后端+前端+AI |
| 3 | IQC/IPQC/FQC完整质量闭环 | M05 | 10天 | 后端+前端 |
| 4 | 首末件检验 | M05 | 5天 | 后端+前端 |
| 5 | 金蝶ERP对接（BOM/订单/报工） | M14 | 10天 | 后端 |
| 6 | 设备数据采集（OPC-UA/MQTT） | M08 | 10天 | 后端 |
| 7 | Andon系统完善（广播/升级） | M09 | 5天 | 后端+前端 |
| 8 | 多车间管理 | M01 | 3天 | 后端+前端 |
| 9 | 打印模板管理 | M01 | 3天 | 后端+前端 |
| 10 | 批次号/序列号编码规则 | M03 | 3天 | 后端 |
| 11 | 包装条码管理 | M03 | 5天 | 后端+前端 |
| 12 | 生产指示单（流程卡） | M03 | 3天 | 前端 |

### 4.2 P1 - 8月底前完成

| 序号 | 功能 | 模块 | 预估工时 |
|------|------|------|---------|
| 1 | 实验室管理/SPC | M11 | 15天 |
| 2 | 量检具管理 | M12 | 8天 |
| 3 | JIT/JIS物料拉动 | M07 | 10天 |
| 4 | AGV调度集成 | M14 | 8天 |
| 5 | 视觉检测对接 | M14 | 5天 |
| 6 | TEEP分析 | M06 | 5天 |
| 7 | 供应商Portal | M14 | 8天 |
| 8 | Alert告警中心 | 全局 | 10天 |
| 9 | e-SOP电子作业指导书 | M03 | 8天 |

### 4.3 P2 - 后续迭代

| 序号 | 功能 | 模块 |
|------|------|------|
| 1 | AI质检模块 | M13 |
| 2 | 移动端APP | 全局 |
| 3 | 低代码报表 | M15 |

---

## 5. 数据库完整DDL

> 以下为新增/需修改的核心表，已有表请参考当前代码库

### 5.1 DDL汇总（新增表）

```sql
-- ============================================
-- M01 系统管理 - 新增
-- ============================================
CREATE TABLE sys_print_template (
    id BIGSERIAL PRIMARY KEY,
    template_code VARCHAR(50) UNIQUE NOT NULL,
    template_name VARCHAR(100) NOT NULL,
    template_type VARCHAR(30) NOT NULL,
    template_content TEXT,
    template_width INTEGER DEFAULT 100,
    template_height INTEGER DEFAULT 50,
    preview_image VARCHAR(500),
    printer_type VARCHAR(20),
    is_default SMALLINT DEFAULT 0,
    is_enabled SMALLINT DEFAULT 1,
    remark VARCHAR(500),
    tenant_id BIGINT NOT NULL,
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- ============================================
-- M03 生产执行 - 新增
-- ============================================
CREATE TABLE mes_code_rule (
    id BIGSERIAL PRIMARY KEY,
    rule_code VARCHAR(50) UNIQUE NOT NULL,
    rule_name VARCHAR(100) NOT NULL,
    rule_type VARCHAR(20) NOT NULL,
    prefix VARCHAR(20),
    date_format VARCHAR(20),
    serial_length INTEGER DEFAULT 4,
    reset_cycle VARCHAR(20),
    current_serial INTEGER DEFAULT 0,
    product_id BIGINT,
    workshop_id BIGINT,
    is_enabled SMALLINT DEFAULT 1,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE mes_first_last_inspect (
    id BIGSERIAL PRIMARY KEY,
    inspect_no VARCHAR(50) UNIQUE NOT NULL,
    inspect_type VARCHAR(10) NOT NULL,
    production_order_id BIGINT NOT NULL,
    process_id BIGINT NOT NULL,
    workstation_id BIGINT NOT NULL,
    shift_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    serial_no VARCHAR(100),
    inspect_items JSONB,
    overall_result VARCHAR(10),
    inspector_id BIGINT,
    inspector_name VARCHAR(50),
    inspect_time TIMESTAMP,
    bluetooth_device_id VARCHAR(100),
    remark VARCHAR(500),
    status VARCHAR(20) DEFAULT 'PENDING',
    tenant_id BIGINT NOT NULL,
    workshop_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE mes_package (
    id BIGSERIAL PRIMARY KEY,
    package_no VARCHAR(100) UNIQUE NOT NULL,
    package_type VARCHAR(20),
    production_order_id BIGINT,
    product_id BIGINT NOT NULL,
    product_code VARCHAR(50),
    qty INTEGER NOT NULL,
    serial_nos JSONB,
    status VARCHAR(20) DEFAULT 'OPEN',
    seal_time TIMESTAMP,
    seal_by VARCHAR(50),
    ship_time TIMESTAMP,
    customer_id BIGINT,
    container_id BIGINT,
    tenant_id BIGINT NOT NULL,
    workshop_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- ============================================
-- M04 AI-APS - 新增
-- ============================================
CREATE TABLE aps_work_center (
    id BIGSERIAL PRIMARY KEY,
    work_center_code VARCHAR(50) UNIQUE NOT NULL,
    work_center_name VARCHAR(100) NOT NULL,
    work_center_type VARCHAR(20) NOT NULL,
    workshop_id BIGINT,
    capacity_unit VARCHAR(20) DEFAULT 'HOUR',
    standard_capacity DECIMAL(18,3) DEFAULT 1,
    max_capacity DECIMAL(18,3) DEFAULT 1,
    efficiency_factor DECIMAL(5,2) DEFAULT 100,
    utilization_target DECIMAL(5,2) DEFAULT 85,
    setup_time INTEGER DEFAULT 0,
    calendar_template_id BIGINT,
    parallel_jobs INTEGER DEFAULT 1,
    status VARCHAR(20) DEFAULT 'ACTIVE',
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE aps_changeover_matrix (
    id BIGSERIAL PRIMARY KEY,
    work_center_id BIGINT NOT NULL,
    from_product_family VARCHAR(50),
    to_product_family VARCHAR(50),
    from_product_id BIGINT,
    to_product_id BIGINT,
    changeover_time INTEGER NOT NULL,
    changeover_type VARCHAR(20),
    notes VARCHAR(200),
    tenant_id BIGINT NOT NULL,
    UNIQUE(work_center_id, from_product_family, to_product_family)
);

CREATE TABLE aps_calendar_template (
    id BIGSERIAL PRIMARY KEY,
    template_code VARCHAR(50) UNIQUE NOT NULL,
    template_name VARCHAR(100) NOT NULL,
    work_days JSONB NOT NULL,
    shifts JSONB NOT NULL,
    holiday_dates JSONB,
    special_work_dates JSONB,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE aps_schedule_plan (
    id BIGSERIAL PRIMARY KEY,
    plan_no VARCHAR(50) UNIQUE NOT NULL,
    plan_name VARCHAR(100) NOT NULL,
    plan_type VARCHAR(20) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    schedule_mode VARCHAR(20) NOT NULL,
    algorithm_type VARCHAR(20) NOT NULL,
    optimize_target VARCHAR(50),
    horizon_days INTEGER DEFAULT 14,
    status VARCHAR(20) DEFAULT 'DRAFT',
    total_orders INTEGER DEFAULT 0,
    scheduled_orders INTEGER DEFAULT 0,
    delayed_orders INTEGER DEFAULT 0,
    on_time_rate DECIMAL(5,2),
    avg_utilization DECIMAL(5,2),
    total_changeover_time INTEGER,
    execute_start_time TIMESTAMP,
    execute_end_time TIMESTAMP,
    execute_duration INTEGER,
    result_summary JSONB,
    ai_suggestions JSONB,
    tenant_id BIGINT NOT NULL,
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE aps_schedule_task (
    id BIGSERIAL PRIMARY KEY,
    plan_id BIGINT NOT NULL REFERENCES aps_schedule_plan(id),
    task_no VARCHAR(50) NOT NULL,
    order_id BIGINT,
    production_order_id BIGINT,
    product_id BIGINT,
    product_code VARCHAR(50),
    product_name VARCHAR(100),
    product_family VARCHAR(50),
    process_id BIGINT,
    process_seq INTEGER,
    work_center_id BIGINT,
    task_type VARCHAR(20) DEFAULT 'PRODUCTION',
    quantity DECIMAL(18,3),
    processing_time INTEGER,
    setup_time INTEGER DEFAULT 0,
    changeover_time INTEGER DEFAULT 0,
    plan_start_time TIMESTAMP,
    plan_end_time TIMESTAMP,
    actual_start_time TIMESTAMP,
    actual_end_time TIMESTAMP,
    earliest_start_time TIMESTAMP,
    latest_end_time TIMESTAMP,
    jit_demand_time TIMESTAMP,
    slack_time INTEGER,
    priority INTEGER DEFAULT 0,
    status VARCHAR(20) DEFAULT 'PLANNED',
    locked SMALLINT DEFAULT 0,
    lock_reason VARCHAR(200),
    predecessor_task_ids JSONB,
    constraint_violations JSONB,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE aps_jit_demand (
    id BIGSERIAL PRIMARY KEY,
    demand_no VARCHAR(50) UNIQUE NOT NULL,
    customer_id BIGINT NOT NULL,
    customer_name VARCHAR(100),
    product_id BIGINT NOT NULL,
    product_code VARCHAR(50),
    demand_qty INTEGER NOT NULL,
    need_time TIMESTAMP NOT NULL,
    delivery_point VARCHAR(100),
    demand_type VARCHAR(20),
    jis_sequence INTEGER,
    status VARCHAR(20) DEFAULT 'RECEIVED',
    production_order_id BIGINT,
    schedule_task_id BIGINT,
    source_message JSONB,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- ============================================
-- M06 设备管理 - 新增
-- ============================================
CREATE TABLE eam_oee_record (
    id BIGSERIAL PRIMARY KEY,
    device_id BIGINT NOT NULL,
    workshop_id BIGINT NOT NULL,
    shift_id BIGINT,
    record_date DATE NOT NULL,
    planned_time INTEGER,
    downtime_planned INTEGER DEFAULT 0,
    downtime_unplanned INTEGER DEFAULT 0,
    downtime_idle INTEGER DEFAULT 0,
    actual_run_time INTEGER,
    standard_cycle_time DECIMAL(10,2),
    theoretical_output INTEGER,
    actual_output INTEGER,
    qualified_output INTEGER,
    availability DECIMAL(5,2),
    performance DECIMAL(5,2),
    quality_rate DECIMAL(5,2),
    oee DECIMAL(5,2),
    teep DECIMAL(5,2),
    downtime_details JSONB,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(device_id, record_date, shift_id)
);

CREATE TABLE eam_mold (
    id BIGSERIAL PRIMARY KEY,
    mold_code VARCHAR(50) UNIQUE NOT NULL,
    mold_name VARCHAR(100) NOT NULL,
    mold_type VARCHAR(30),
    applicable_products JSONB,
    standard_life INTEGER,
    current_count INTEGER DEFAULT 0,
    remaining_life INTEGER,
    next_maintain_count INTEGER,
    status VARCHAR(20) DEFAULT 'ACTIVE',
    location VARCHAR(100),
    supplier_id BIGINT,
    purchase_date DATE,
    seal_date DATE,
    seal_reason VARCHAR(200),
    attachments JSONB,
    tenant_id BIGINT NOT NULL,
    workshop_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- ============================================
-- M07 仓库 - 新增
-- ============================================
CREATE TABLE wms_material_pull (
    id BIGSERIAL PRIMARY KEY,
    pull_no VARCHAR(50) UNIQUE NOT NULL,
    pull_type VARCHAR(20) NOT NULL,
    workstation_id BIGINT NOT NULL,
    material_id BIGINT NOT NULL,
    material_code VARCHAR(50),
    qty_requested DECIMAL(18,3) NOT NULL,
    qty_delivered DECIMAL(18,3) DEFAULT 0,
    request_time TIMESTAMP NOT NULL,
    need_time TIMESTAMP,
    actual_deliver_time TIMESTAMP,
    agv_task_id VARCHAR(100),
    handler_id BIGINT,
    status VARCHAR(20) DEFAULT 'PENDING',
    trigger_source VARCHAR(20),
    production_order_id BIGINT,
    tenant_id BIGINT NOT NULL,
    workshop_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- ============================================
-- M08 数据采集 - 新增
-- ============================================
CREATE TABLE dc_data_point (
    id BIGSERIAL PRIMARY KEY,
    point_code VARCHAR(50) UNIQUE NOT NULL,
    point_name VARCHAR(100) NOT NULL,
    device_id BIGINT NOT NULL,
    data_type VARCHAR(20),
    protocol VARCHAR(20) NOT NULL,
    address VARCHAR(200),
    unit VARCHAR(20),
    scan_rate INTEGER DEFAULT 1000,
    deadband DECIMAL(10,4) DEFAULT 0,
    store_policy VARCHAR(20) DEFAULT 'ALL',
    alarm_enabled SMALLINT DEFAULT 0,
    alarm_high DECIMAL(18,4),
    alarm_low DECIMAL(18,4),
    map_to_field VARCHAR(100),
    status VARCHAR(20) DEFAULT 'ACTIVE',
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE dc_scan_log (
    id BIGSERIAL PRIMARY KEY,
    scan_type VARCHAR(20) NOT NULL,
    scan_code VARCHAR(200) NOT NULL,
    scan_device VARCHAR(50),
    workstation_id BIGINT,
    scan_user_id BIGINT,
    scan_time TIMESTAMP DEFAULT NOW(),
    parse_result JSONB,
    business_type VARCHAR(50),
    related_id BIGINT,
    status VARCHAR(20) DEFAULT 'SUCCESS',
    fail_reason VARCHAR(200),
    tenant_id BIGINT NOT NULL,
    workshop_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- ============================================
-- M09 Andon - 新增
-- ============================================
CREATE TABLE andon_escalation_rule (
    id BIGSERIAL PRIMARY KEY,
    rule_name VARCHAR(100) NOT NULL,
    andon_type VARCHAR(30),
    workshop_id BIGINT,
    level1_timeout INTEGER NOT NULL,
    level1_notify JSONB,
    level2_timeout INTEGER,
    level2_notify JSONB,
    level3_timeout INTEGER,
    level3_notify JSONB,
    notify_channels JSONB,
    is_enabled SMALLINT DEFAULT 1,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- ============================================
-- M10 追溯 - 新增
-- ============================================
CREATE TABLE mes_traceability (
    id BIGSERIAL PRIMARY KEY,
    trace_no VARCHAR(64) UNIQUE NOT NULL,
    serial_no VARCHAR(100),
    batch_no VARCHAR(100),
    product_id BIGINT NOT NULL,
    product_code VARCHAR(50),
    production_order_id BIGINT,
    production_date DATE,
    production_line_id BIGINT,
    status VARCHAR(20) DEFAULT 'ACTIVE',
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE mes_trace_material (
    id BIGSERIAL PRIMARY KEY,
    trace_no VARCHAR(64) NOT NULL,
    material_id BIGINT,
    material_code VARCHAR(50),
    material_batch VARCHAR(100),
    supplier_id BIGINT,
    qty_used DECIMAL(18,3),
    iqc_record_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE mes_trace_process (
    id BIGSERIAL PRIMARY KEY,
    trace_no VARCHAR(64) NOT NULL,
    process_seq INTEGER,
    process_name VARCHAR(100),
    workstation_id BIGINT,
    operator_id BIGINT,
    operator_name VARCHAR(50),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    parameters JSONB,
    device_id BIGINT,
    device_parameters JSONB,
    quality_result VARCHAR(20),
    created_at TIMESTAMP DEFAULT NOW()
);

-- ============================================
-- M11 实验室/SPC - 新增
-- ============================================
CREATE TABLE lab_spc_chart_config (
    id BIGSERIAL PRIMARY KEY,
    config_code VARCHAR(50) UNIQUE NOT NULL,
    config_name VARCHAR(100) NOT NULL,
    chart_type VARCHAR(20) NOT NULL,
    product_id BIGINT,
    process_id BIGINT,
    inspect_item VARCHAR(100) NOT NULL,
    sample_size INTEGER DEFAULT 5,
    sample_frequency INTEGER DEFAULT 25,
    ucl DECIMAL(18,6),
    lcl DECIMAL(18,6),
    cl DECIMAL(18,6),
    usl DECIMAL(18,6),
    lsl DECIMAL(18,6),
    target_value DECIMAL(18,6),
    auto_calculate_limits SMALLINT DEFAULT 1,
    is_enabled SMALLINT DEFAULT 1,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE lab_spc_data (
    id BIGSERIAL PRIMARY KEY,
    config_id BIGINT NOT NULL,
    sample_group INTEGER NOT NULL,
    sample_seq INTEGER,
    measured_value DECIMAL(18,6) NOT NULL,
    batch_no VARCHAR(100),
    serial_no VARCHAR(100),
    production_order_id BIGINT,
    workstation_id BIGINT,
    operator_id BIGINT,
    measure_time TIMESTAMP NOT NULL,
    gauge_id BIGINT,
    x_bar DECIMAL(18,6),
    r_value DECIMAL(18,6),
    s_value DECIMAL(18,6),
    is_out_of_control SMALLINT DEFAULT 0,
    rule_violations JSONB,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- ============================================
-- M12 器具管理 - 新增
-- ============================================
CREATE TABLE container_master (
    id BIGSERIAL PRIMARY KEY,
    container_code VARCHAR(50) UNIQUE NOT NULL,
    container_name VARCHAR(100) NOT NULL,
    container_type VARCHAR(30) NOT NULL,
    standard_qty INTEGER,
    applicable_products JSONB,
    status VARCHAR(20) DEFAULT 'IN_STOCK',
    location_type VARCHAR(20),
    current_location VARCHAR(100),
    customer_id BIGINT,
    barcode VARCHAR(100) UNIQUE,
    total_trips INTEGER DEFAULT 0,
    last_clean_date DATE,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE container_movement (
    id BIGSERIAL PRIMARY KEY,
    container_id BIGINT NOT NULL,
    movement_type VARCHAR(20) NOT NULL,
    from_location VARCHAR(100),
    to_location VARCHAR(100),
    qty INTEGER DEFAULT 1,
    related_order_no VARCHAR(50),
    operator_id BIGINT,
    movement_time TIMESTAMP NOT NULL,
    remark VARCHAR(200),
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- ============================================
-- M13 AI质检 - 视觉对接
-- ============================================
CREATE TABLE qms_vision_result (
    id BIGSERIAL PRIMARY KEY,
    result_no VARCHAR(50) UNIQUE NOT NULL,
    source_system VARCHAR(50),
    workstation_id BIGINT,
    production_order_id BIGINT,
    product_id BIGINT,
    serial_no VARCHAR(100),
    image_urls JSONB,
    ai_judgment VARCHAR(10) NOT NULL,
    defect_types JSONB,
    confidence DECIMAL(5,4),
    inspection_time TIMESTAMP NOT NULL,
    final_judgment VARCHAR(10),
    confirmed_by BIGINT,
    confirmed_time TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- ============================================
-- M14 集成接口
-- ============================================
CREATE TABLE integration_erp_sync_log (
    id BIGSERIAL PRIMARY KEY,
    sync_type VARCHAR(50) NOT NULL,
    direction VARCHAR(10) NOT NULL,
    erp_bill_no VARCHAR(100),
    mes_bill_no VARCHAR(100),
    request_body TEXT,
    response_body TEXT,
    status VARCHAR(20),
    error_msg TEXT,
    retry_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 创建索引
CREATE INDEX idx_aps_task_plan_id ON aps_schedule_task(plan_id);
CREATE INDEX idx_aps_task_work_center ON aps_schedule_task(work_center_id);
CREATE INDEX idx_aps_task_plan_time ON aps_schedule_task(plan_start_time, plan_end_time);
CREATE INDEX idx_dc_scan_time ON dc_scan_log(scan_time);
CREATE INDEX idx_spc_data_config ON lab_spc_data(config_id, measure_time);
CREATE INDEX idx_trace_serial ON mes_traceability(serial_no);
CREATE INDEX idx_trace_batch ON mes_traceability(batch_no);
```

---

## 6. API接口全清单

### 6.1 M01 系统管理（新增部分）

```
# 打印模板
GET    /api/v1/system/print-templates
POST   /api/v1/system/print-templates
PUT    /api/v1/system/print-templates/:id
DELETE /api/v1/system/print-templates/:id
POST   /api/v1/system/print/execute

# 系统参数（已有，补充峰梅参数键值）
GET    /api/v1/system/configs
PUT    /api/v1/system/configs
```

### 6.2 M03 生产执行（新增）

```
# 编码规则
GET    /api/v1/production/code-rules
POST   /api/v1/production/code-rules
PUT    /api/v1/production/code-rules/:id
POST   /api/v1/production/code-rules/generate  # 生成编码

# 首末件检验
GET    /api/v1/production/first-last-inspect
POST   /api/v1/production/first-last-inspect
PUT    /api/v1/production/first-last-inspect/:id
GET    /api/v1/production/first-last-inspect/overdue

# 包装条码
GET    /api/v1/production/packages
POST   /api/v1/production/packages/create
POST   /api/v1/production/packages/add-item
POST   /api/v1/production/packages/seal
GET    /api/v1/production/packages/:id

# 流程卡
GET    /api/v1/production/orders/:id/flow-card
POST   /api/v1/production/orders/:id/flow-card/print
```

### 6.3 M04 AI-APS

```
# 工作中心
GET    /api/v1/aps/work-centers
POST   /api/v1/aps/work-centers
PUT    /api/v1/aps/work-centers/:id
DELETE /api/v1/aps/work-centers/:id

# 换型矩阵
GET    /api/v1/aps/changeover-matrix
POST   /api/v1/aps/changeover-matrix
PUT    /api/v1/aps/changeover-matrix/:id

# 日历模板
GET    /api/v1/aps/calendars
POST   /api/v1/aps/calendars
PUT    /api/v1/aps/calendars/:id

# 排程计划
GET    /api/v1/aps/plans
POST   /api/v1/aps/plans
GET    /api/v1/aps/plans/:id
DELETE /api/v1/aps/plans/:id
POST   /api/v1/aps/plans/:id/execute     # 执行排程
POST   /api/v1/aps/rolling-schedule      # 滚动排程（定时任务触发）

# 甘特图
GET    /api/v1/aps/gantt/data
PUT    /api/v1/aps/gantt/tasks/:id/time
PUT    /api/v1/aps/gantt/tasks/:id/center
POST   /api/v1/aps/gantt/tasks/:id/lock
POST   /api/v1/aps/gantt/tasks/:id/unlock
GET    /api/v1/aps/gantt/conflict-check

# JIT需求
GET    /api/v1/aps/jit-demands
POST   /api/v1/aps/jit-demands           # 接收主机厂JIT
GET    /api/v1/aps/jit-demands/:id

# 分析
GET    /api/v1/aps/analysis/capacity          # 产能分析
GET    /api/v1/aps/analysis/on-time-rate      # 准时交付率
GET    /api/v1/aps/analysis/material-shortage # 缺料分析
```

### 6.4 M05 质量管理（新增）

```
# 检验标准
GET    /api/v1/quality/inspect-standards
POST   /api/v1/quality/inspect-standards
PUT    /api/v1/quality/inspect-standards/:id

# 蓝牙量具（前端直接通过Web Bluetooth，此为数据保存接口）
POST   /api/v1/quality/gauge-readings
```

### 6.5 M06 设备管理（新增）

```
# OEE/TEEP
GET    /api/v1/equipment/oee                 # OEE报表
GET    /api/v1/equipment/teep                # TEEP分析
POST   /api/v1/equipment/oee/manual-input   # 手动录入

# 模具
GET    /api/v1/equipment/molds
POST   /api/v1/equipment/molds
PUT    /api/v1/equipment/molds/:id
POST   /api/v1/equipment/molds/:id/seal
GET    /api/v1/equipment/molds/:id/usage-history
```

### 6.6 M07 仓库（新增）

```
# 物料拉动
GET    /api/v1/warehouse/material-pulls
POST   /api/v1/warehouse/material-pulls
PUT    /api/v1/warehouse/material-pulls/:id/accept
PUT    /api/v1/warehouse/material-pulls/:id/complete

# 线边看板
GET    /api/v1/warehouse/line-side-stock     # 线边库存看板数据
```

### 6.7 M08 数据采集（新增）

```
GET    /api/v1/dc/data-points
POST   /api/v1/dc/data-points
PUT    /api/v1/dc/data-points/:id
GET    /api/v1/dc/realtime/:deviceId         # 实时数据
GET    /api/v1/dc/scan-logs                  # 扫描记录
POST   /api/v1/dc/scan                       # 扫码
```

### 6.8 M09 Andon（补充）

```
POST   /api/v1/andon/calls
PUT    /api/v1/andon/calls/:id/respond
PUT    /api/v1/andon/calls/:id/handle
PUT    /api/v1/andon/calls/:id/escalate
GET    /api/v1/andon/calls
GET    /api/v1/andon/statistics
GET    /api/v1/andon/escalation-rules
POST   /api/v1/andon/escalation-rules
PUT    /api/v1/andon/escalation-rules/:id
```

### 6.9 M10 追溯（新增）

```
GET    /api/v1/trace/forward/:serialNo
GET    /api/v1/trace/backward/:batchNo
GET    /api/v1/trace/detail/:traceNo
```

### 6.10 M11 实验室SPC（新增）

```
GET    /api/v1/lab/spc/configs
POST   /api/v1/lab/spc/configs
PUT    /api/v1/lab/spc/configs/:id
POST   /api/v1/lab/spc/data
GET    /api/v1/lab/spc/chart/:configId
GET    /api/v1/lab/spc/capability/:configId
GET    /api/v1/lab/spc/pareto
```

### 6.11 M12 器具管理（新增）

```
GET    /api/v1/containers
POST   /api/v1/containers
PUT    /api/v1/containers/:id
POST   /api/v1/containers/:id/in          # 入库
POST   /api/v1/containers/:id/out         # 出库
POST   /api/v1/containers/:id/return     # 退回
GET    /api/v1/containers/:id/movements  # 流转记录
```

### 6.12 M13 AI质检（新增）

```
POST   /api/v1/vision/results             # 视觉系统推送
GET    /api/v1/vision/results
PUT    /api/v1/vision/results/:id/confirm
```

### 6.13 M14 集成接口（新增）

```
# ERP（内部定时任务调用）
POST   /api/v1/erp/sync/bom              # 手动触发BOM同步
POST   /api/v1/erp/sync/orders           # 手动触发订单同步
GET    /api/v1/erp/sync/logs             # 同步日志

# AGV回调（AGV系统调用）
POST   /api/v1/agv/callback/task-complete
POST   /api/v1/agv/callback/task-exception

# 供应商Portal
POST   /api/v1/supplier/asn
GET    /api/v1/supplier/delivery-orders

# JIT（主机厂推送）
POST   /api/v1/jit/demands
```

### 6.14 WebSocket频道

```
ws://host/api/v1/ws?token=xxx

# 订阅频道（发送消息）
{"action": "subscribe", "channel": "andon", "workshop_id": 1}
{"action": "subscribe", "channel": "production", "line_id": 1}
{"action": "subscribe", "channel": "inventory", "workshop_id": 1}
{"action": "subscribe", "channel": "equipment", "device_id": 1}
{"action": "subscribe", "channel": "spc_alarm"}

# 服务端推送消息类型
ANDON_CALL         # 安灯呼叫
ANDON_ESCALATE     # 安灯升级
PRODUCTION_UPDATE  # 产量更新
INVENTORY_ALERT    # 库存预警
EQUIPMENT_ALARM    # 设备报警
SPC_VIOLATION      # SPC失控
QUALITY_ALARM      # 质量报警
APS_COMPLETE       # 排程完成
```

---

## 7. 前端页面路由清单

```javascript
// router/index.js 路由配置

const routes = [
  // M01 系统管理
  { path: '/system/user',            component: UserList },
  { path: '/system/role',            component: RoleList },
  { path: '/system/menu',            component: MenuList },
  { path: '/system/dept',            component: DeptList },
  { path: '/system/dict',            component: DictList },
  { path: '/system/config',          component: SystemConfig },
  { path: '/system/print-template',  component: PrintTemplate },    // 新增
  { path: '/system/workshop-config', component: WorkshopConfig },   // 新增
  { path: '/system/log/oper',        component: OperLog },
  { path: '/system/log/login',       component: LoginLog },

  // M02 主数据
  { path: '/mdm/material',           component: MaterialList },
  { path: '/mdm/bom',                component: BomList },
  { path: '/mdm/route',              component: RoutingList },
  { path: '/mdm/operation',          component: OperationList },
  { path: '/mdm/workshop',           component: WorkshopList },
  { path: '/mdm/line',              component: LineList },
  { path: '/mdm/workstation',        component: WorkstationList },
  { path: '/mdm/shift',              component: ShiftList },
  { path: '/mdm/supplier',           component: SupplierList },
  { path: '/mdm/customer',           component: CustomerList },     // 新增
  { path: '/mdm/import',             component: BatchImport },

  // M03 生产执行
  { path: '/production/order',       component: ProductionOrderList },
  { path: '/production/order/:id',   component: ProductionOrderDetail },
  { path: '/production/dispatch',   component: DispatchList },
  { path: '/production/report',      component: ReportList },
  { path: '/production/code-rule',   component: CodeRuleList },     // 新增
  { path: '/production/first-last',  component: FirstLastInspect }, // 新增
  { path: '/production/package',     component: PackageManage },    // 新增
  { path: '/production/progress',    component: ProductionProgress },

  // M04 APS
  { path: '/aps/sales-order',        component: SalesOrderList },
  { path: '/aps/mps',               component: MPSList },
  { path: '/aps/mrp',               component: MRPList },
  { path: '/aps/work-center',        component: WorkCenterList },   // 新增
  { path: '/aps/changeover',         component: ChangeoverMatrix }, // 新增
  { path: '/aps/calendar',          component: CalendarTemplate },  // 新增
  { path: '/aps/plan',              component: SchedulePlan },      // 新增
  { path: '/aps/gantt',             component: GanttChart },
  { path: '/aps/jit',               component: JITDemand },         // 新增
  { path: '/aps/analysis',           component: APSAnalysis },     // 新增

  // M05 质量
  { path: '/quality/iqc',           component: IQCList },
  { path: '/quality/ipqc',          component: IPQCList },
  { path: '/quality/fqc',           component: FQCList },
  { path: '/quality/oqc',           component: OQCList },
  { path: '/quality/defect',         component: DefectRecordList },
  { path: '/quality/ncr',           component: NCRList },
  { path: '/quality/standard',      component: InspectStandard },  // 新增

  // M06 设备
  { path: '/equipment/list',         component: EquipmentList },
  { path: '/equipment/check',        component: CheckList },
  { path: '/equipment/maintenance',  component: MaintenanceList },
  { path: '/equipment/repair',       component: RepairList },
  { path: '/equipment/spare',        component: SparePartList },
  { path: '/equipment/oee',          component: OEEAnalysis },      // 新增
  { path: '/equipment/teep',         component: TEEPAnalysis },     // 新增
  { path: '/equipment/mold',         component: MoldList },          // 新增

  // M07 仓库
  { path: '/warehouse/list',         component: WarehouseList },
  { path: '/warehouse/location',     component: LocationList },
  { path: '/warehouse/inventory',    component: InventoryList },
  { path: '/warehouse/receive',      component: ReceiveOrderList },
  { path: '/warehouse/delivery',     component: DeliveryOrderList },
  { path: '/warehouse/pull',         component: MaterialPull },     // 新增
  { path: '/warehouse/line-side',    component: LineSideStock },   // 新增（大屏）

  // M08 数据采集
  { path: '/dc/data-point',          component: DataPointList },    // 新增
  { path: '/dc/monitor',             component: DataMonitor },      // 新增
  { path: '/dc/scan',               component: ScanLog },           // 新增

  // M09 Andon
  { path: '/andon/call',             component: AndonCall },
  { path: '/andon/rule',            component: EscalationRule },   // 新增
  { path: '/andon/stats',           component: AndonStats },        // 新增

  // M10 追溯
  { path: '/trace/query',            component: TraceQuery },

  // M11 实验室/SPC
  { path: '/lab/spc-config',         component: SPCConfig },        // 新增
  { path: '/lab/spc-chart',          component: SPCChart },         // 新增
  { path: '/lab/capability',         component: CapabilityAnalysis },// 新增
  { path: '/lab/pareto',            component: ParetoChart },       // 新增

  // M12 器具
  { path: '/container/list',        component: ContainerList },    // 新增
  { path: '/container/movement',    component: ContainerMovement }, // 新增

  // M13 AI质检
  { path: '/vision/result',         component: VisionResult },     // 新增

  // M15 报表
  { path: '/report/production',      component: ProductionReport },
  { path: '/report/quality',         component: QualityReport },
  { path: '/report/equipment',       component: EquipmentReport },
  { path: '/report/delivery',        component: DeliveryReport },

  // 大屏
  { path: '/screen/balance-shaft',   component: BalanceShaftScreen },
  { path: '/screen/battery',         component: BatteryScreen },
  { path: '/screen/trial',          component: TrialScreen },
]
```

---

## 8. 开发排期与里程碑

### 8.1 P0冲刺计划（目标: 2026年7月15日）

> 开发周期约11周（2026年5月1日 - 7月15日）

| 周次 | 时间 | 开发重点 | 交付物 |
|------|------|---------|-------|
| W1-W2 | 5/1-5/15 | 多车间管理、批次码规则、打印模板、ERP对接框架 | 基础数据就绪 |
| W3-W4 | 5/16-5/31 | 首末件检验、包装条码、APS工作中心+换型矩阵 | 生产执行完善 |
| W5-W6 | 6/1-6/15 | APS排程引擎（规则排程+CP求解器）、甘特图 | APS核心完成 |
| W7-W8 | 6/16-6/30 | AI优化器（OR-Tools）、JIT接入、数据采集OPC-UA | APS+数采 |
| W9-W10 | 7/1-7/10 | Andon升级广播、追溯完善、系统集成联调 | 集成联调 |
| W11 | 7/11-7/15 | 性能测试、数据初始化、用户培训 | 上线 |

### 8.2 P1计划（2026年8月底）

| 功能 | 负责模块 | 预计完成 |
|------|---------|---------|
| SPC实验室管理 | M11 | 7月31日 |
| 器具管理 | M12 | 7月31日 |
| JIT/JIS电子看板 | M07 | 8月15日 |
| AGV集成 | M14 | 8月15日 |
| TEEP分析 | M06 | 8月31日 |
| 供应商Portal | M14 | 8月31日 |

### 8.3 风险识别

| 风险 | 概率 | 影响 | 应对 |
|------|------|------|------|
| AI-APS排程复杂度超预期 | 高 | 高 | 先上规则排程，AI优化器后迭代 |
| 金蝶ERP接口对接延期 | 中 | 高 | 提前获取API文档，并行开发 |
| AGV系统接口不稳定 | 中 | 中 | 支持手动配送作为降级方案 |
| 设备OPC-UA协议兼容性 | 中 | 中 | 提前采集设备信息，现场调试 |
| 峰梅需求变更 | 高 | 中 | 敏捷迭代，周迭代对齐 |

---

## 9. 非功能性需求

### 9.1 性能要求

| 指标 | 目标值 | 验收方式 |
|------|-------|---------|
| 事务响应时间 | ≤3秒 | 接口压测 |
| 列表查询响应 | ≤15秒 | 接口压测 |
| APS排程时间 | ≤120秒（100工单） | 功能测试 |
| 系统并发用户 | ≥100人同时在线 | 压测 |
| 系统可用性 | ≥99.9%（7×24h） | 监控 |
| 数据采集延迟 | ≤1秒 | 实测 |

### 9.2 数据保留策略

| 数据类型 | 保留期限 | 存储方式 |
|---------|---------|---------|
| 基础数据（物料/BOM等） | 永久 | 主库 |
| 业务数据（工单/检验等） | 12个月活跃 | 主库 |
| 历史数据 | 10年 | 历史库归档 |
| 设备采集数据 | 1年 | 分区表 + 定期归档 |

**自动归档**: 每日凌晨2点将12个月前的业务数据迁移至历史库，主库保持高性能。

### 9.3 安全要求

- JWT认证，Token有效期2小时，支持RefreshToken
- 按钮级RBAC权限控制
- 所有API日志记录（操作人、时间、参数、响应）
- SQL注入防护（使用GORM参数化查询）
- 数据隔离：TenantID + WorkshopID 双重隔离
- 敏感配置（数据库密码、API Key）不入代码，使用环境变量

### 9.4 部署要求

| 容器 | 基础镜像 | 端口 | 说明 |
|------|---------|------|------|
| mom-frontend | nginx:alpine | 80 | Vue静态资源 |
| mom-backend | golang:1.23-alpine | 8080 | 主API服务 |
| mom-aps-ai | python:3.10-slim | 9000 | AI排程服务 |
| mom-websocket | golang:1.23-alpine | 8083 | WS推送服务 |
| postgresql | postgres:15 | 5432 | 主数据库 |
| redis | redis:7-alpine | 6379 | 缓存 |
| kafka | confluentinc/kafka | 9092 | 消息队列 |

### 9.5 知识产权

- 源码知识产权双方共有（闻荫科技有二次开发权）
- 架构支持SOA，B/S框架，Linux/Windows双平台
- 支持未来多工厂复制部署

---

*文档版本: V1.0 | 生成日期: 2026年4月6日 | 适用项目: 闻荫科技MOM3.0*
