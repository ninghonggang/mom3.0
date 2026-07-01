# MOM 3.0 文档维护规约

> 最后更新：2026-07-01
> 适用范围：`/data/mom3.0/docs/` 全部活跃文档（不含 `archive/`、`mom-server/`、`mom-web/` 源代码内嵌注释）
> 借鉴来源：[faway_mes docs/](../home/nhg/faway_mes/docs/DOCUMENTATION_GUIDE.md)、Arc42 模板、IEEE 1016

---

## 1. 核心原则（5 条）

1. **改代码必改文档**：修改 API、路由、模型、状态机、状态值 → **必须**同步对应模块的设计文档；CI 检查
2. **文档与代码同源同 PR**：文档 PR 与代码 PR 必须联动 review，不能"代码先合文档后补"
3. **图表用 Mermaid/PlantUML**：禁止 ASCII art、截图替代、Word 框图（截图只能用于仪表盘/报表样例）
4. **废弃文档 mv 到 archive/**：不要堆在活跃目录；归档后更新 `archive/README.md` 索引
5. **每篇文档都有修订日期**：放在文档第 1 行 `> 最后更新：YYYY-MM-DD`

---

## 2. 目录与文件命名

### 2.1 目录约定

| 路径 | 内容 | 维护策略 |
|------|------|---------|
| `docs/README.md` | 文档中心入口 | 文档结构变更时更新 |
| `docs/DOCUMENTATION_GUIDE.md` | 本规约 | 规约变更时更新 |
| `docs/MOM3.0_主设计文档.md` | 系统总览 | 业务范围、模块、术语变更时更新 |
| `docs/MOM3.0_UI设计规范.md` | UI 规范 | UI 风格变更时更新 |
| `docs/MOM3.0_技术架构文档.md` | 技术架构 | 架构变更时更新 |
| `docs/MOM3.0_模块设计模板.md` | 模块文档模板 | 模板章节调整时更新 |
| `docs/MOM3.0_附录.md` | 通用附录（错误码字典、术语表、角色、字段命名）| 各模块共享内容变更时更新 |
| `docs/MOM3.0_<模块名>模块设计文档.md` | 后端模块设计 | 模块 owner 维护 |
| `docs/MOM3.0前端_<模块名>模块设计文档.md` | 前端模块设计 | 前端 owner 维护 |
| `docs/MOM3.0_SAP_Gap_Analysis_and_Development_Plan.md` | SAP 对照整改 | 架构组维护 |
| `docs/DEVELOPMENT_STATUS.md` | 模块实现状态 | 测试组维护（每次发版前更新）|
| `docs/TODO.md` | P0/P1/P2 问题清单 | 全员（bug 修复时勾掉）|
| `docs/rules/` | 项目开发规范（Vue/后端/工作流/组件）| 架构组维护 |
| `docs/superpowers/plans/` | 开发计划（按周）| 项目经理维护 |
| `docs/research/` | 调研与审计报告 | 调研人维护 |
| `docs/archive/` | ⛔ 废弃归档（**只读**）| 任何人 mv 进来后由 owner 更新 README 索引 |

### 2.2 文件命名规范

| 类型 | 格式 | 示例 |
|------|------|------|
| 模块设计文档 | `MOM3.0_<模块名>模块设计文档.md` | `MOM3.0_质量模块设计文档.md` |
| 前端设计文档 | `MOM3.0前端_<模块名>模块设计文档.md` | `MOM3.0前端_QMS质量管理模块设计文档.md` |
| 模板 | `MOM3.0_<文档类型>模板.md` | `MOM3.0_模块设计模板.md` |
| 总览 | `MOM3.0_<文档类型>文档.md` | `MOM3.0_主设计文档.md` |
| 调研报告 | `<主题>-<类型>-<日期>.md` | `MOM3.0-design-doc-audit-2026-07-01.md` |
| 开发计划 | `YYYY-MM-DD-<项目>-<计划主题>.md` | `2026-04-17-MOM3.0-开发计划.md` |

**禁止**：

- 文件名带空格（用下划线或连字符代替）
- 中文标点混入文件名
- 同名不同后缀重复文件

---

## 3. 文档结构规范

### 3.1 章节顺序（所有模块设计文档统一）

```
1.  模块概述（业务定位、Top 3 干系人、Top 3 质量目标）
2.  依赖关系（上游模块、下游模块、外部系统）
3.  功能清单（功能 / 描述 / 状态 / 优先级）
4.  页面与交互（路由 / 标题 / 关键按钮 / 表格列 / 表单字段）
5.  业务流程（★ 必须是 BPMN/活动图/时序图，禁止文字流）
6.  状态机（★ 必须有 Mermaid stateDiagram）
7.  数据模型（★ 必须有 Mermaid erDiagram + 字段表）
8.  API 规范（路由表 + 请求/响应 schema + 错误码 + 幂等/限流）
9.  角色与权限（角色 × 操作矩阵）
10. 集成与事件（上下游消息流、事件发布/订阅、消息格式）
11. 可观测性（关键指标、告警规则、日志样例）
12. 非功能需求（性能 P95、吞吐量、可用性、数据量、保留期）
13. 附录（CHANGELOG、相关链接）
```

> 章节 5/6/7/8/9/10/12 是 MOM 3.0 现有文档**严重缺失**的，**必须**补。其它章节如无可省略。

完整模板见 [MOM3.0_模块设计模板.md](./MOM3.0_模块设计模板.md)。

### 3.2 章节编号规范

- 用一级 `#`、二级 `##`、三级 `###` 标题
- 章节序号必须连续（不要跳号）
- 附录用字母（A.1, A.2）
- 子章节不超过 4 级（避免过深）

---

## 4. 命名与字段规范

### 4.1 数据库字段（统一）

| 类别 | 命名 | 备注 |
|------|------|------|
| 主键 | `id` | 系统自增 |
| 外键 | `<ref_table_singular>_id` | 例：`workshop_id` |
| 审计 | `created_at` / `created_by` / `updated_at` / `updated_by` / `deleted_at` | 每张业务表必含 |
| 租户隔离 | `tenant_id` | 多租户表必含 |
| 布尔字段 | `is_xxx` / `has_xxx` | 例：`is_enabled` |
| 时间字段 | `xxx_at`（时刻）/ `xxx_date`（日期）/ `xxx_interval`（时段）| |
| 枚举字段 | `<field>_code` | 例：`status_code` |
| 多语言字段 | `<field>_i18n`（JSON 格式）| 推荐用 JSON，不推荐多列 |
| 货币字段 | `amount` + `currency_code` | 分离存储 |
| 计量字段 | `value` + `unit_code` | 分离存储 |

完整规范见 [MOM3.0_附录.md § 字段命名规范](./MOM3.0_附录.md)。

### 4.2 API 路由规范

| 类别 | 命名 | 示例 |
|------|------|------|
| 资源集合 | `GET/POST /<resource>` | `GET /production/orders` |
| 单资源 | `GET/PUT/DELETE /<resource>/:id` | `GET /production/orders/:id` |
| 子资源 | `GET /<parent>/:pid/<child>` | `GET /production/orders/:pid/reports` |
| 动作 | `POST /<resource>/:id/<action>` | `POST /production/orders/:id/release` |
| 高级搜索 | `POST /<resource>/search` | `POST /production/orders/search` |
| 导出/导入 | `GET/POST /<resource>/export-import` | `GET /production/orders/export-excel` |
| 版本 | `/api/v<MAJOR>/...` | `/api/v1/...` |

完整规范见 [MOM3.0_附录.md § API 路由规范](./MOM3.0_附录.md)。

### 4.3 错误码规范

**4 段式编码**：`模块段(2位) + 类别段(2位) + 序号段(2位) = 6位`

```
01 = 系统模块
02 = 主数据
03 = 生产执行
04 = APS
05 = 质量
06 = 设备
07 = WMS
08 = 数据采集
09 = 告警
10 = 追溯
14 = 集成
20 = 通用错误
21 = 鉴权
22 = 限流

类别段（2 位）：
00 = 成功
01 = 通用校验
02 = 业务校验
03 = 资源不存在
04 = 资源冲突
05 = 状态错误
06 = 权限错误
07 = 集成错误

示例：
01-01-0001 = 系统模块-通用校验-0001
05-03-0001 = 质量模块-资源不存在-0001
```

完整字典见 [MOM3.0_附录.md § 错误码字典](./MOM3.0_附录.md)。

### 4.4 状态值规范

- 状态字段：`status`（主状态）或 `status_code`（字典枚举）
- 状态值命名：`UPPER_SNAKE_CASE`，语义清晰，避免缩写
- 状态值必含 `DELETED` / `CANCELLED` 终态
- 状态转移文档化（章节 6）

---

## 5. 图表规范（强制）

### 5.1 必须用 Mermaid/PlantUML

| 场景 | 工具 | 说明 |
|------|------|------|
| 业务流程 | Mermaid `flowchart` | GitHub/GitLab 原生渲染 |
| 状态机 | Mermaid `stateDiagram-v2` | 每个核心实体 1 张 |
| ER 图 | Mermaid `erDiagram` | 轻量、与文档同源 |
| 时序图（简单）| Mermaid `sequenceDiagram` | 同上 |
| 复杂时序 | PlantUML | 需要 VSCode 插件或 server |
| C4 架构 | PlantUML + C4-PlantUML | 放 `docs/architecture/` |
| 部署图 | PlantUML `deployment` | 同上 |

### 5.2 禁止使用

- ❌ ASCII art 框图（`┌─┐│└┘`）
- ❌ 截图替代流程图（截图只能用于仪表盘/报表样例）
- ❌ Word 框图（除非在 .docx 设计文档中）
- ❌ 文字流（「创建X → 编辑X → 保存X」）

### 5.3 Mermaid 渲染验证

提交前在 GitHub Preview 或 VSCode Mermaid Preview 插件验证渲染正常。

---

## 6. 何时更新（强制约束）

| 触发事件 | 必须更新 | 建议更新 |
|---------|---------|---------|
| 新增/修改 API 路由 | 模块设计文档 § 8 | - |
| 新增/修改状态值 | 模块设计文档 § 6 状态机 | 附录错误码字典 |
| 新增/修改数据库表 | 模块设计文档 § 7 ER 图 + 字段表 | 附录字段命名 |
| 新增/修改业务流程 | 模块设计文档 § 5 流程图 | - |
| 新增/修改模块边界 | 主设计文档 § 2 模块地图 | 技术架构文档 § C4 |
| 新增 UI 规范 | UI 设计规范文档 | 模块前端文档 |
| 新增安全合规要求 | 技术架构文档 § 安全分区 | - |
| 发版（版本号变化）| DEVELOPMENT_STATUS.md | - |
| 修 bug | TODO.md | - |
| 调研/审计 | research/ 目录 | - |

### 6.1 CI 检查（建议接入）

- 文档修改日期超过 90 天的模块 → 标记 stale
- 文档引用的代码路径不存在 → 警告
- API 路由在文档里有但代码里没有 → 警告（反过来也行）

---

## 7. Review 流程

### 7.1 PR 要求

| 类型 | 必须 Reviewer |
|------|--------------|
| 模块设计文档 | 模块 owner + 架构组 |
| 总览/架构文档 | 架构组 + 模块受影响 owner |
| 模板/规约 | 架构组 + 全员 review |
| 调研/审计 | 调研人 + 1 名架构 |

### 7.2 Review 检查清单

- [ ] 章节顺序符合 [§ 3.1](#31-章节顺序所有模块设计文档统一)
- [ ] 图表全部使用 Mermaid/PlantUML（[§ 5](#5-图表规范强制)）
- [ ] 状态机、ER 图、流程图齐全（适用于模块设计）
- [ ] 错误码、字段命名符合规范（[§ 4](#4-命名与字段规范)）
- [ ] 修订日期更新（第 1 行）
- [ ] 链接全部有效（无死链）
- [ ] 与代码保持同步（API 路由、字段名、状态值一致）
- [ ] 不含重复内容（不与主设计/附录重复）

---

## 8. 归档流程

### 8.1 何时归档

- 旧版本设计（如 MOM 3.0 V1、SFMS3.0）
- 已被新文档替代的设计
- 调研结束后不再需要的对比文档

### 8.2 操作步骤

1. 确认新文档已覆盖旧文档内容
2. `mv 旧文件 docs/archive/<分类>/`
3. 更新 `docs/archive/README.md` 索引
4. 删除活跃区对旧文件的链接（搜索 + 替换）

### 8.3 归档目录结构

```
docs/archive/
├── README.md                          # 归档索引
├── 2026-04-full-rewrite/             # SFMS3.0 时代大批文档
├── 2026-04-v1/                        # MOM 3.0 V1 设计
└── 2026-07-清理-nextjs-模板/          # Next.js 模板清理批次
```

### 8.4 归档规则

- ⛔ **禁止修改** 归档目录中的文件（除非做错误修正 + 记录在归档 README）
- ⛔ **禁止** 在活跃区链接归档文件（用 `archive/README.md` 索引）
- ✅ 可以 `mv` 新文件进归档，但要在 `archive/README.md` 加一行记录

---

## 9. 工具链

### 9.1 编辑工具

| 工具 | 用途 | 备注 |
|------|------|------|
| VSCode + Markdown All in One | 编辑 | 推荐 |
| VSCode + Mermaid Preview | 预览 Mermaid | 必装 |
| VSCode + PlantUML | 预览 PlantUML | 必装（架构/部署图用）|
| GitHub/GitLab | 渲染 | Mermaid 原生 |

### 9.2 长期项目（规划中）

- **VitePress 静态站**：搜索、版本、导航、SEO 强，但维护成本
- **OpenAPI 自动生成**：用 swag/oapi-codegen，文档与代码同步
- **CODEMAP.md**：借鉴 faway_mes 的 53KB 版本

---

## 10. 例外与豁免

- **历史快照**：superpowers/plans/、archive/ 不强制更新
- **调研文档**：research/ 不强制规范章节顺序（按需）
- **临时文档**：可在 docs/ 根目录临时放 `.todo.md` 或 `.draft.md` 前缀文件，7 天后归档或删除

---

## 11. 修订记录

| 版本 | 日期 | 修订人 | 说明 |
|------|------|--------|------|
| V1.0 | 2026-07-01 | 架构组 | 初版，基于 faway_mes DOCUMENTATION_GUIDE 简化 |