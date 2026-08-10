# MOM Platform 3.0

制造运营管理系统 (Manufacturing Operations Management) 微服务平台。

## 架构概览

```
┌─────────────────────────────────────────────────────────┐
│                    前端 (Next.js)                        │
│                  http://localhost:3001                   │
└──────────────────────┬──────────────────────────────────┘
                       │ REST API
┌──────────────────────▼──────────────────────────────────┐
│                API 网关 (NestJS)                         │
│              http://localhost:3000                       │
│              Swagger: /api/docs                          │
└──────┬───────┬───────┬───────┬──────┬──────┬──────┬─────┘
       │gRPC   │gRPC   │gRPC   │gRPC  │gRPC  │gRPC  │gRPC
  ┌────▼──┐┌──▼───┐┌──▼───┐┌──▼──┐┌─▼───┐┌─▼───┐┌─▼────┐
  │ MDM  ││ MES  ││ QMS  ││ EAM ││ WMS ││ APS ││Andon │
  │:50051││:50052││:50053││:50054││:50055│:50056│:50058│
  └──────┘└──────┘└──────┘└─────┘└─────┘└─────┘└──────┘
       │       │       │       │       │       │      │
  ┌────▼───────▼───────▼───────▼───────▼───────▼──────▼──┐
  │           PostgreSQL 16  (多租户/多库)                 │
  │           NATS 2.10      (消息/事件)                   │
  │           Redis 7        (缓存/锁)                     │
  │           MinIO          (文件/图纸)                   │
  └───────────────────────────────────────────────────────┘
```

## 技术栈

| 层 | 技术 | 端口 |
|---|------|------|
| 前端 | Next.js 14 + Tailwind + shadcn/ui | 3001 |
| 网关 | NestJS + Swagger | 3000 |
| 微服务 | Go + gRPC + GORM + PostgreSQL | 50051-50058 |
| 消息 | NATS (JetStream) | 4222 |
| 缓存 | Redis | 6379 |
| 存储 | MinIO (S3) | 9000 |

## 微服务列表

| 服务 | 端口 | 数据库 | 职责 |
|------|------|--------|------|
| mdm-service | 50051 | mom_mdm | 主数据：物料/BOM/车间/产线/工位/客户/供应商 |
| mes-service | 50052 | mom_mes | 生产执行：工单/派工/报工/完工/流转卡 |
| qms-service | 50053 | mom_qms | 质量管理：检验单/NCR/SPC/缺陷代码/AQL |
| eam-service | 50054 | mom_eam | 设备管理：台账/维修/OEE/保养/点检/停机 |
| wms-service | 50055 | mom_wms | 仓储管理：库存/入库/出库/盘点/调拨 |
| aps-service | 50056 | mom_aps | 计划排程：MPS/MRP/排程/工作中心/换型 |
| trace-service | 50057 | mom_trace | 追溯与采集：追溯链/序列号/数据采集/扫码 |
| andon-service | 50058 | mom_andon | 安灯与告警：安灯呼叫/告警配置/升级规则 |

## 快速开始

### 1. 启动基础设施

```bash
docker-compose up -d postgres nats redis minio
```

### 2. 启动微服务

```bash
# 全部启动
make run

# 或单个启动
cd services/mes-service && GOWORK=off go run ./cmd/main.go
```

### 3. 启动网关

```bash
cd gateway && npm install && npm run start:dev
```

### 4. 启动前端

```bash
cd web && npm install && npm run dev
```

### 一键启动 (Docker)

```bash
make docker-up-all
```

### 访问地址

| 服务 | URL |
|------|-----|
| 前端 | http://localhost:3001 |
| API 文档 | http://localhost:3000/api/docs |
| NATS 监控 | http://localhost:8222 |
| MinIO 控制台 | http://localhost:9001 |

## 项目结构

```
mom-platform/
├── proto/                    # Protobuf 契约 (25 个 .proto)
│   ├── common/               #   通用类型 + 分页 + 状态码
│   ├── mdm/                  #   物料/BOM/车间
│   ├── mes/                  #   工单/派工/报工
│   ├── qms/                  #   检验/NCR/SPC
│   ├── eam/                  #   设备/维修/OEE
│   ├── wms/                  #   库存/入库/出库
│   ├── aps/                  #   计划/排程/约束
│   ├── trace/                #   追溯/采集/序列号
│   ├── andon/                #   安灯/告警/升级
│   └── events/               #   领域事件定义
├── services/                 # Go 微服务 (7 个, 64 个 .go)
│   ├── mdm-service/
│   ├── mes-service/
│   ├── qms-service/
│   ├── eam-service/
│   ├── wms-service/
│   ├── aps-service/
│   ├── trace-service/
│   └── andon-service/
├── gateway/                  # NestJS API 网关 (40 个 .ts)
├── web/                      # Next.js 前端 (41 个 .tsx/.ts)
├── scripts/                  # 数据库初始化脚本
├── docker-compose.yml        # 12 个容器编排
├── Makefile                  # 开发工具链
└── buf.yaml / buf.gen.yaml   # Proto 编译配置
```

## API 端点

网关提供 **69 个 REST 端点**，跨 9 个业务模块：

| 模块 | 端点数 | 前缀 |
|------|:------:|------|
| Dashboard | 4 | `/api/dashboard` |
| MES | 7 | `/api/mes` |
| QMS | 9 | `/api/qms` |
| EAM | 9 | `/api/eam` |
| WMS | 12 | `/api/wms` |
| MDM | 9 | `/api/mdm` |
| APS | 5 | `/api/aps` |
| Trace | 8 | `/api/trace` |
| Andon | 6 | `/api/andon` |

## 核心业务规则

### 状态机
- **检验单**: PENDING → IN_PROGRESS → PASSED/FAILED/WAIVED
- **NCR**: OPEN → INVESTIGATING → DISPOSITIONED → VERIFIED → CLOSED
- **维修单**: REPORTED → ASSIGNED → IN_PROGRESS → COMPLETED → VERIFIED
- **入库单**: DRAFT → RECEIVING → RECEIVED → PUTAWAY → COMPLETED
- **出库单**: DRAFT → PICKING → PICKED → PACKING → SHIPPED
- **安灯**: TRIGGERED → ACKNOWLEDGED → IN_PROGRESS → RESOLVED → CLOSED

### 告警规则
- P0 告警 → 全渠道通知 (钉钉+短信+邮件+声光)
- 5 分钟未响应 → 升级 L1 (车间主任)
- 15 分钟未响应 → 升级 L2 (厂长)
- 30 分钟未响应 → 升级 L3 (生产副总)
- 设备 OEE < 70% → 1 周持续告警

### 库存不变量
`available_qty = quantity - locked_qty`

## License

Internal Use Only
