# MOM3.0 设计文档提升建议

**日期**: 2026-07-01
**配套阅读**:
- [MOM3.0-design-doc-audit-2026-07-01.md](./MOM3.0-design-doc-audit-2026-07-01.md) — 现状审计
- [mes-design-best-practices-2026-07-01.md](./mes-design-best-practices-2026-07-01.md) — 行业最佳实践调研（686 行）

**目标**: 把 MOM 3.0 设计文档从「3/10 工程级质量」提到「7/10」，给出**可执行、分阶段、带工作量**的方案。

---

## 0. 30 秒结论

- **现状一句话**：30K 行 markdown 覆盖了 16 个模块，但结构雷同、缺图、缺总览、3 个文档指向不存在的总览文件、2 个文档还停留在 SFMS3.0 时代（BPM 写 MySQL、设备/WMS 写 Java 目录）。
- **建议一句话**：**先补 4 个总览文档 + 1 个统一模板 + 1 套图示工具链；再补 10 个缺失章节；最后清理旧项目残留**。分 3 阶段、共约 22-28 人天。
- **不在本建议范围内**：代码重构、OpenAPI 自动生成、引入 BPMN 工具等需要长期投入的事项（只在「长期」节简述）。

---

## 1. 整体升级路线图

```
阶段 1（本周，3-4 天）: 补「骨架」 — 让总览不再缺席
  ├─ 1.1 新建 MOM3.0_主设计文档.md （系统总览、模块地图、技术架构、术语表）
  ├─ 1.2 新建 MOM3.0_UI设计规范.md （布局、状态色、表格、表单弹窗、错误展示）
  ├─ 1.3 新建 MOM3.0_技术架构文档.md （C4 组件图、部署图、数据流图）
  ├─ 1.4 新建 docs/README.md （文档地图 + 阅读路径）
  └─ 1.5 建立文档维护规约 docs/DOCUMENTATION_GUIDE.md

阶段 2（1-2 周，8-10 天）: 立「规范」 — 统一模板 + 图表工具链
  ├─ 2.1 写 MOM3.0_模块设计模板.md （13 章节 Arc42 简化版 + MES 必备小节）
  ├─ 2.2 用 Mermaid 重画所有「业务流程」段（删掉文字流、ASCII art）
  ├─ 2.3 用 Mermaid stateDiagram 给核心实体画状态机（工单/批次/设备/检验单/库存/工位）
  ├─ 2.4 用 Mermaid erDiagram 给核心域画 ER 图（主数据/生产/质量/库存）
  ├─ 2.5 补全 5 个标准附录：角色权限矩阵 / 错误码字典 / 状态机一览 / 核心指标公式 / 术语表
  └─ 2.6 写 MOM3.0_API规范.md （URL 风格、错误体、幂等、限流、鉴权、OpenAPI 生成）

阶段 3（1-2 月，10-14 天）: 清「死角」 — 补薄文档 + 修技术栈偏差 + 删残留
  ├─ 3.1 给 5 个「< 200 行」的薄模块（ASP 217、追溯 143、报表 139、主数据 159、集成 160）扩到 500-800 行
  ├─ 3.2 修 BPM 文档的技术栈描述（Go + GORM + PostgreSQL 18 + pgvector）
  ├─ 3.3 修设备模块的 MySQL DDL（转成 PostgreSQL 18 DDL 或删掉）
  ├─ 3.4 修 WMS 模块的「win-module-wms-*」目录残留（改名 + 改技术栈）
  ├─ 3.5 把 SAP Gap Analysis 的 P0 待开发项写进对应模块的「版本历史」和「TODO」节
  ├─ 3.6 增补 SAP Gap Analysis 没覆盖的：可观测性、安全合规、多租户、国际化
  └─ 3.7 把 docs/.obsolete/ 旧文档全部 mv 到 docs/archive/，加 archive/README.md
```

**总工作量估算**:
- 阶段 1: 3-4 人天
- 阶段 2: 8-10 人天
- 阶段 3: 10-14 人天
- **合计: 22-28 人天（约 4-5 周单人全职）**

> 假设：1 人 = 文档 owner 1 名 + 1 名开发 review 协助。多人分担可压缩到 2-3 周。

---

## 2. 阶段 1 详细：补 4 个总览文档 + 1 个规约

### 2.1 `docs/MOM3.0_主设计文档.md`（新建，~400 行）

**目的**: 解决「所有模块文档都引用但文件不存在」的问题。

**章节**:
1. 系统定位与目标（1 段话 + 3 个核心目标 + Top 3 质量目标）
2. 业务背景（1 段话 + 干系人清单）
3. 业务范围（做什么/不做什么）
4. **模块地图**（16 模块 × 1 行：`M01 系统管理 | M02 主数据 | M03 生产执行 | ... | M16 SCP 供应链`）
5. 术语表（10-30 个 MOM 行业术语 + 缩略语）
6. 文档阅读路径（新开发 / 实施 / 客户演示 / 维护 4 类角色怎么读）
7. 关联文档清单（所有模块文档 + 配套资产）

**模板来源**: 借鉴 [faway_mes overview/README.md](../../../../../home/nhg/faway_mes/docs/overview/README.md) + [Arc42 Section 1-3](https://docs.arc42.org/)。

### 2.2 `docs/MOM3.0_UI设计规范.md`（新建，~300 行）

**目的**: 解决「同 MES 模块标准布局」= 0 内容的循环引用。

**章节**:
1. 整体布局（左侧菜单 + 顶栏 + 内容区）
2. 列表页模板（搜索区 / 工具栏 / 表格 / 分页）
3. 表单弹窗模板（增/改/详情 3 种弹窗的字段布局）
4. 状态色规范（success/warning/danger/info 各自对应什么业务场景）
5. 按钮规范（主/次/危险/链接 4 类的位置和颜色）
6. 空状态/加载/错误展示（图片占位 + 文案规范）
7. 移动端/平板适配规则
8. 国际化文案规范（i18n key 命名、占位符、复数、日期格式）

**模板来源**: faway_mes 的 web_layout_enhance 风格 + Element Plus 设计指南。

### 2.3 `docs/MOM3.0_技术架构文档.md`（新建，~500 行）

**目的**: 给人看的架构图 + 部署图 + 数据流图（不是给 AI 的 Context Hub）。

**章节**（Arc42 Section 3-7 风格）:
1. **C4 Level 1 上下文图**（PlantUML）— MOM 在工厂 IT/OT 网络里的位置
2. **C4 Level 2 容器图** — 前后端 + DB + 缓存 + 消息 + 网关
3. **C4 Level 3 组件图** — 后端 handler/service/repo 三层 + 前端 api/view/store
4. **部署视图**（PlantUML deployment）— dev/test/uat/prod 4 套环境 + 节点清单
5. **数据流图** — 关键场景的时序图（工单下发 / 报工回传 / 批次放行 / 设备数据采集）
6. **集成拓扑** — 上下游系统 + 协议（ERP、QMS、WMS、AGV、视觉检测、SCADA、PLC）
7. **安全分区**（Purdue 模型 Layer 0-4）— IT/OT 隔离、DMZ、跳板

### 2.4 `docs/README.md`（新建，~100 行）

**章节**:
- 文档目录结构
- 快速导航表（按角色）
- 文档状态（活跃 / 已废弃 → archive）
- 文档维护规约链接

### 2.5 `docs/DOCUMENTATION_GUIDE.md`（新建，~200 行）

**借鉴** [faway_mes 的版本](../../../../../home/nhg/faway_mes/docs/DOCUMENTATION_GUIDE.md)。

**核心内容**:
- 活跃区 vs 归档区（活跃区 = overview / modules / appendix；归档区 = archive/）
- 何时更新（改代码 → 改对应模块 DESIGN.md + CODEMAP.md + CHANGELOG）
- 章节顺序规范（每类文档固定的章节顺序）
- 命名规范（文件命名 / 章节编号 / 锚点命名）
- 图表规范（必须用 Mermaid/PlantUML，禁止 ASCII art + 截图替代）
- 文档 review 流程（PR 必须有 doc owner review）
- 归档流程（废弃文档 mv 到 archive/，更新 archive/README.md）

---

## 3. 阶段 2 详细：模板 + 图表工具链

### 3.1 `docs/MOM3.0_模块设计模板.md`（新建，~250 行）

**模板主干（13 章节 Arc42 简化版 + MES 必备小节）**:

| # | 章节 | 必含内容 | 工具/格式 |
|---|------|---------|----------|
| 1 | **模块概述** | 业务目标、价值流位置、Top 3 干系人、Top 3 质量目标 | 文字 + 1 张系统上下文子图 |
| 2 | **依赖关系** | 上游模块、下游模块、外部系统、对应的 ISA-95 段 | 表格 + PlantUML 依赖图 |
| 3 | **功能清单** | 完整功能列表 + 每项 1 行说明 | 表格（功能 / 描述 / 状态 / 优先级）|
| 4 | **页面与交互** | 页面清单（路由 / 标题 / 关键按钮 / 表格列 / 表单字段）| 表格 |
| 5 | **业务流程** | 关键流程用 BPMN/活动图，禁用文字流 | Mermaid flowchart / PlantUML activity |
| 6 | **状态机** | 本模块核心实体的生命周期 | Mermaid stateDiagram-v2 |
| 7 | **数据模型** | ER 图（必含）+ 字段表 + 索引/约束/审计字段 | Mermaid erDiagram + 表格 |
| 8 | **API 规范** | 路由表 + 请求/响应 schema + 错误码 + 幂等/限流 | 表格 + 链接到 OpenAPI 文档 |
| 9 | **角色与权限** | 角色 × 操作矩阵（谁能做什么）| 矩阵表 |
| 10 | **集成与事件** | 上下游消息流、事件发布/订阅、消息格式 | 时序图 + AsyncAPI 片段 |
| 11 | **可观测性** | 关键指标、告警规则、日志样例 | 表格 + Prometheus metric 清单 |
| 12 | **非功能需求** | 性能（P95 / 吞吐量）、可用性、数据量、保留期 | 表格 + 量化值 |
| 13 | **附录** | 术语表、错误码字典、变更日志（CHANGELOG）| 表格 |

> 章节 5/6/7/8/9/10 是 MOM 现有文档**严重缺失**的，**必须**补。其它章节如无可省略。

**配套: 1 个 `docs/MOM3.0_附录.md`**（新建）:
- 错误码字典（所有模块共享）
- 术语表（公司 + ISA-95 + ISA-88 + MESA）
- 角色定义（系统管理员 / 计划员 / 车间主任 / 班组长 / 操作工 / 质检员 / 设备工程师）
- 通用字段命名规范（created_at / created_by / tenant_id / deleted_at 等）
- CHANGELOG 模板

### 3.2 图表工具链统一

| 场景 | 工具 | 落地方式 |
|------|------|---------|
| 业务流程 | **Mermaid flowchart** | GitHub/GitLab 原生渲染，零成本 |
| 状态机 | **Mermaid stateDiagram-v2** | 同上 |
| ER 图 | **Mermaid erDiagram** | 同上（轻量、足够） |
| 复杂时序图 | **Mermaid sequenceDiagram** 或 **PlantUML** | 复杂场景用 PlantUML |
| 架构图（C4）| **PlantUML + C4-PlantUML** | 单独 arch 仓库或 docs/architecture/ |
| 部署图 | **PlantUML deployment** | 同上 |
| 仪表盘 / 报表截图 | 保留 PNG | design-references/ 目录 |
| **禁止** | ASCII art、流程图截图、Word 框图 | 全部替换 |

**清理目标**:
- 删除所有"创建X → 编辑X → 保存X"形式的文字流程
- 删除追溯文档的 ASCII art 框图 → 换 Mermaid
- 删除 UI 章节里"同 X 模块标准布局"的循环引用 → 用统一的 UI 设计规范

### 3.3 MOM 3.0 必备状态机清单（按业务域）

> 这 12 个状态机是行业最低集合（基于 ISA-95 + MESA 11 项 + 行业共识）。MOM 3.0 已实现功能里**多数**没画出来。

| 状态机 | 当前文档状态 | 优先级 |
|--------|------------|--------|
| 生产工单（RELEASED → IN_PROGRESS → COMPLETED → CLOSED）| 仅在 ASP 文档列了 5 个状态值，无图 | P0 |
| 批次（CREATED → IN_PROCESS → QA_HOLD → RELEASED）| 完全没图 | P0 |
| 设备（IDLE → RUNNING → DOWN → MAINTENANCE → RETIRED）| 设备文档列了 5 个状态值 | P1 |
| 质量检验（PENDING → IN_PROGRESS → PASS / FAIL）| 质量文档列了 6 个状态值 | P1 |
| 库存（AVAILABLE → LOCKED → ALLOCATED → ISSUED）| 完全没图 | P1 |
| 销售订单（DRAFT → CONFIRMED → DELIVERED → INVOICED）| SCP 文档列了 7 个状态值 | P2 |
| 采购订单（DRAFT → PENDING → APPROVED → RECEIVED → CLOSED）| SCP 文档列了 8 个状态值 | P2 |
| 容器（ACTIVE → IN_USE → MAINTENANCE → SCRAPPED）| 其他文档有 ASCII art | P2 |
| 安灯（RAISED → ACKNOWLEDGED → RESOLVED）| 完全没图 | P2 |
| NCR（OPEN → INVESTIGATING → RESOLVED → CLOSED）| 质量文档列了 5 个状态值 | P1 |
| 维修工单（PENDING → ASSIGNED → IN_PROGRESS → COMPLETED）| 设备文档列了 5 个状态值 | P2 |
| 异常/告警（NEW → ACK → RESOLVED → CLOSED）| 完全没图 | P2 |

### 3.4 MOM 3.0 必备 ER 图清单

> 这是 MESA 11 项 + ISA-95 Common Object Model 的最小集合。

| ER 图 | 涉及表 | 建议位置 |
|-------|--------|---------|
| **生产主数据 ER** | 物料 / BOM / 工艺路线 / 工作中心 / 资源 / 班次 | 阶段 2 写入「主数据」文档 |
| **生产执行 ER** | 工单 / 派工 / 报工 / 流转卡 / 首末件 / 完工入库 | 阶段 2 写入「MES 生产」文档 |
| **质量 ER** | 检验单 / 检验明细 / 不良代码 / NCR / SPC / 8D | 阶段 2 写入「质量」文档 |
| **设备 ER** | 设备台账 / 点检 / 保养 / 维修 / 备件 / OEE | 阶段 2 写入「设备」文档 |
| **库存 ER** | 仓库 / 库位 / 库存 / 收货 / 发货 / 调拨 / 盘点 | 阶段 2 写入「WMS」文档（注意：当前 WMS 文档要先清理 SFMS3.0 残留）|
| **供应链 ER** | 采购 / 销售 / 询价 / 报价 / 供应商 / 客户 | 阶段 2 写入「SCP」文档 |
| **追溯 ER** | 序列号 / 批次 / 工序 / 投料 / 设备 | 阶段 2 写入「追溯」文档 |

---

## 4. 阶段 3 详细：清死角

### 4.1 修 4 个与代码不符的文档（高优先）

| 文档 | 问题 | 修复方式 | 估时 |
|------|------|---------|------|
| BPM 流程模块设计文档 | 技术栈写 MyBatis-Plus + MySQL + Spring Security，实际 Go + GORM + PostgreSQL | 重写「技术架构」章节 | 0.5d |
| 设备管理模块设计文档 | 大段 MySQL DDL（`datetime DEFAULT CURRENT_TIMESTAMP` 等），实际 PostgreSQL 18 | 整段删或转 PostgreSQL DDL | 1d |
| WMS 仓储模块设计文档 | 6 千多行里大半是 SFMS3.0 时代的 Java 内容（`win-module-wms-biz` 目录）| 删 Java 部分，按 MOM 3.0 实际重写 | 1.5d |
| 主设计文档（不存在）| 所有模块都引用 | 阶段 1 解决 | 0.5d（已计阶段 1）|

### 4.2 扩 5 个「< 200 行」的薄模块文档

| 文档 | 当前 | 目标 | 估时 |
|------|------|------|------|
| APS 计划模块设计文档 | 217 行 | 800 行（补齐状态机 + ER + 时序 + 排程算法）| 1.5d |
| 追溯与数据采集模块设计文档 | 143 行 | 600 行（补齐谱系模型 + 追溯查询性能 + ISA-95 段映射）| 1d |
| 报表模块设计文档 | 139 行 | 500 行（补齐指标计算公式 + 数据来源 + 刷新策略）| 0.5d |
| 主数据管理模块设计文档 | 159 行 | 600 行（补齐 BOM 多层结构 + 工艺路线 ER）| 0.5d |
| 系统集成模块设计文档 | 160 行 | 600 行（补齐集成模式 + 消息流 + 重试幂等）| 0.5d |

### 4.3 清理旧项目残留

```
/data/mom3.0/docs/
├── .obsolete/                      # 旧 SFMS3.0 文档
│   ├── full/                       # 14 个旧设计文档
│   ├── MOM3.0_主设计文档_v1_wrong.md
│   └── ...其他旧版
├── ──→ 全部 mv 到 /data/mom3.0/docs/archive/
├── archive/                        # 新建
│   ├── README.md                   # 归档索引
│   ├── 2026-04-full-rewrite/       # 旧 SFMS3.0 大批文档
│   └── 2026-04-v1/                # 旧 MOM 3.0 V1 设计
└── superpowers/plans/              # 保留（是开发计划，不是设计文档）
```

### 4.4 把 SAP Gap Analysis 的待开发项落到模块文档

SAP Gap Analysis 列出了 P0/P1/P2 待开发项和工时。这些**应该写进对应模块文档的「功能清单」章节**（标记"待开发"），否则下次复盘找不到对应位置。

---

## 5. 可选/远期（不在本次范围）

> 这些是 MOM 3.0 走向工业级**应该做但工作量很大**的事，建议独立成项目。

| 事项 | 工具/方法 | 估时 | 价值 |
|------|----------|------|------|
| **OpenAPI 自动生成** | 用 swag / oapi-codegen 从 Go 注解生成 OpenAPI 3.1 | 2-3d | API 文档与代码同步 |
| **AsyncAPI 文档** | 给消息流（AGV、ERP、视觉检测、SCADA）写 AsyncAPI | 3-5d | 集成可对接 |
| **PlantUML 架构图独立仓库** | docs/architecture/，CI 渲染 PNG | 1-2d | 架构评审 |
| **CODEMAP.md**（借鉴 faway_mes 53KB 版本）| 自动化生成（每张表/每个 API 列行号）| 2-3d | 找代码快 |
| **BPMN 流程图设计器集成** | bpmn.io / Camunda Modeler 嵌入 BPM 模块 | 10+d | 流程可配置 |
| **状态机可视化（每个实体可点击看机）**| 用 stately.ai / XState 描述 | 5+d | 测试覆盖 |
| **追溯谱系图（基因图）**| Neo4j 落谱系 + GraphQL 查询 | 10+d | 复杂召回 |
| **多租户架构** | 重新设计 schema（tenant → site → area → line → cell）| 30+d | 多工厂 |

---

## 6. 关键决策点（要老板拍板）

1. **文档要 Markdown 还是 DocFX/VitePress 站？**
   - 现成 MD + GitHub 渲染：**零成本**，但搜索/导航弱
   - DocFX/VitePress 静态站：**+2-3 天建站**，但搜索/版本/导航/SEO 全有
   - **建议**: 阶段 1 用 MD，阶段 3 末再考虑 VitePress 站

2. **图表是 Mermaid 还是 PlantUML？**
   - Mermaid: 轻量、GitHub 原生、足够覆盖 80% 场景
   - PlantUML: 表达力强（C4/部署/复杂时序）但需要 PlantUML server
   - **建议**: 流程/状态机/ER 用 Mermaid，复杂架构/部署用 PlantUML（VSCode 插件 + 单独 docs/architecture/）

3. **要不要把 SAP Gap Analysis 的 P0/P1/P2 直接并入模块文档？**
   - 并入：方便开发按模块找 TODO，但文档体积会膨胀 30%
   - 不并入：保持文档清洁，但要靠 docs/TODO.md 作为索引
   - **建议**: 并入（每个模块文档加一节「未完成功能」+ 链接到 SAP Gap Analysis）

4. **faway_mes 风格要不要照搬？**
   - 借鉴 faway 的**结构**（overview/modules/manual/archive + CODEMAP + GAP_ANALYSIS + GUIDELINE）有极大价值
   - 但 faway 是 Odoo，模块化方式与 MOM 3.0 不同；照搬形式不照搬内容
   - **建议**: 借鉴 faway 的「分层 + 维护规约」，不照搬 Odoo 特定章节

5. **OpenAPI/AsyncAPI 自动生成什么时候做？**
   - 现在做（阶段 3 中插入）— 一开始就文档驱动，但会拖慢开发
   - 模块全做完再做 — 但到时候补很痛苦
   - **建议**: 阶段 3 中插入，**新写的 API 立刻生成**，老 API 边维护边补

6. **docx 的《MOM3.0_完整系统设计规范 V1.0.docx》怎么办？**
   - 转 markdown 整合进 docs/
   - 留在根目录作为"对外交付物"
   - 归档到 archive/
   - **建议**: 转 markdown 进 docs/，根目录留 .docx 作为"客户演示版"

---

## 7. 风险与依赖

| 风险 | 影响 | 缓解 |
|------|------|------|
| 文档 owner 没时间维护 | 阶段 2/3 烂尾 | 强制 PR 流程：改 API 必须改对应文档，CI 检查 |
| 团队对 Mermaid/PlantUML 不熟 | 阶段 2 慢 | 阶段 1 末做 2 小时工具培训，文档模板里给范例 |
| 跟代码漂移 | 又回 3/10 状态 | 阶段 3 末加 `docs/CODEMAP.md` + CI 检查 |
| 老员工觉得「文档够多了」 | 推动难 | 先用 1 张对比图（现状 vs 目标）说服老板 |
| 老板希望「文档先不动，把功能补完」 | 优先级冲突 | SAP Gap Analysis 已经在管功能；文档不冲突，二者可并行 |

---

## 8. 下一步行动（建议立刻做）

1. **今天/明天**: 拍板上面 6 个决策（特别是 1、3、5）
2. **本周内**: 完成阶段 1（4 个总览 + 1 个规约 = 5 个新文档）
3. **下周**: 启动阶段 2（先做模块设计模板 + 1 个示范模块，比如「生产执行」重写作为样板）
4. **本月**: 阶段 2 完成，阶段 3 启动
5. **下月**: 阶段 3 完成，开始长期项目（OpenAPI/AsyncAPI/CODEMAP/VitePress）

**需要老板给的资源**:
- 1 名文档 owner（专职或 30% 时间）
- 1 名开发协助 review（10% 时间）
- 工具培训时间：2 小时（Mermaid + PlantUML）

---

## 9. 验证标准（阶段 3 完成后怎么评）

| 维度 | 阶段 1 前 | 阶段 3 后目标 |
|------|----------|--------------|
| 覆盖广度 | 8 | 9 |
| 结构规范 | 4 | 8 |
| 图表表达 | 1 | 8 |
| 与代码同步 | 3 | 7 |
| 可维护性 | 3 | 8 |
| 总览性 | 2 | 9 |
| API 文档化 | 3 | 7 |
| 数据建模文档 | 2 | 7 |
| 非功能需求 | 1 | 6 |
| **总体** | **3 / 10** | **7 / 10** |

> 8/10 以上需要 OpenAPI + AsyncAPI + CODEMAP + VitePress 站等长期项目（见 §5）。

---

**报告完结**。需要看现状细节 → [audit 报告](./MOM3.0-design-doc-audit-2026-07-01.md)；
需要看行业最佳实践 → [调研报告](./mes-design-best-practices-2026-07-01.md)。
