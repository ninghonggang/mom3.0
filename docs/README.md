# MOM 3.0 文档中心

> 最后更新：2026-07-03
> 适用范围：`/data/mom3.0/` 全部项目文档（不含 `mom-server/`、`mom-web/` 源代码内的内嵌注释）
> **V2.0 文档升级状态**：9/16 module 已 V2.0,详见 [§ 4.5 V2.0 升级进度](#45-v20-升级进度)

---

## 1. 是什么

MOM 3.0（Manufacturing Operations Management，制造运营管理系统）是面向**离散制造多品种小批量**场景自研的工业软件，覆盖主数据、计划排程、生产执行、质量、设备、仓储、追溯、供应链等核心业务。遵循 ISA-95 / IEC 62264 标准，支持多工厂、多车间部署。

文档维护规约见 [DOCUMENTATION_GUIDE.md](./DOCUMENTATION_GUIDE.md)。

---

## 2. 文档结构

```
docs/
├── README.md                         ← 本文件（入口）
├── DOCUMENTATION_GUIDE.md            ← 文档维护规约
│
├── MOM3.0_主设计文档.md              ← ★ 系统总览（业务范围、模块地图、术语表）
├── MOM3.0_UI设计规范.md              ← ★ UI 设计规范（布局、状态色、表单）
├── MOM3.0_技术架构文档.md            ← ★ 技术架构（C4 图、部署图、数据流）
│
├── MOM3.0_<模块名>模块设计文档.md    ← 16 个模块设计文档（每个模块 1 篇）
│
├── MOM3.0_模块设计模板.md            ← 模块设计文档的统一模板
├── MOM3.0_附录.md                    ← 通用附录：错误码字典 / 术语表 / 角色定义 / 字段命名
│
├── MOM3.0_SAP_Gap_Analysis_and_Development_Plan.md   ← SAP 对照 + 整改计划
├── DEVELOPMENT_STATUS.md             ← 模块实现状态（API 覆盖率）
├── TODO.md                           ← P0/P1/P2 问题清单
│
├── research/                         ← 调研与审计
├── superpowers/plans/                ← 开发计划（按周）
├── rules/                            ← 项目开发规范（Vue/后端/工作流）
└── archive/                          ← ⛔ 废弃归档（只读）
```

---

## 3. 快速导航

### 3.1 按角色

| 角色 | 推荐阅读路径 |
|------|--------------|
| **新员工入职** | [MOM3.0_主设计文档.md](./MOM3.0_主设计文档.md) → 模块设计文档（按需）→ [DEVELOPMENT_STATUS.md](./DEVELOPMENT_STATUS.md) |
| **后端开发** | [MOM3.0_技术架构文档.md](./MOM3.0_技术架构文档.md) → [MOM3.0_模块设计模板.md](./MOM3.0_模块设计模板.md) → 负责模块的 `MOM3.0_<模块名>模块设计文档.md` → [TODO.md](./TODO.md) |
| **前端开发** | [MOM3.0_UI设计规范.md](./MOM3.0_UI设计规范.md) → [MOM3.0_模块设计模板.md](./MOM3.0_模块设计模板.md) → 负责模块的 `MOM3.0_<模块名>前端设计文档.md` → `rules/前端/` |
| **产品 / 业务** | [MOM3.0_主设计文档.md](./MOM3.0_主设计文档.md) → 各模块设计文档的「模块概述」「业务流程」节 → [MOM3.0_SAP_Gap_Analysis_and_Development_Plan.md](./MOM3.0_SAP_Gap_Analysis_and_Development_Plan.md) |
| **实施 / 客户演示** | [MOM3.0_主设计文档.md](./MOM3.0_主设计文档.md) → [DEVELOPMENT_STATUS.md](./DEVELOPMENT_STATUS.md) → 演示模块的 `MOM3.0_<模块名>模块设计文档.md` |
| **运维 / 部署** | [MOM3.0_技术架构文档.md](./MOM3.0_技术架构文档.md) 的「部署视图」+「安全分区」节 |
| **架构评审** | [MOM3.0_技术架构文档.md](./MOM3.0_技术架构文档.md) + `research/MOM3.0-design-doc-improvement-2026-07-01.md` |

### 3.2 按任务

| 我想… | 去这里 |
|-------|--------|
| 了解系统全貌 | [MOM3.0_主设计文档.md](./MOM3.0_主设计文档.md) § 2 模块地图 |
| 知道一个模块的状态 | [DEVELOPMENT_STATUS.md](./DEVELOPMENT_STATUS.md) |
| 找一个 bug 在哪 | [TODO.md](./TODO.md) + GitHub Issues |
| 知道下一步要做什么 | [MOM3.0_SAP_Gap_Analysis_and_Development_Plan.md](./MOM3.0_SAP_Gap_Analysis_and_Development_Plan.md) |
| 看架构图 | [MOM3.0_技术架构文档.md](./MOM3.0_技术架构文档.md) |
| 看 UI 规范 | [MOM3.0_UI设计规范.md](./MOM3.0_UI设计规范.md) |
| 写一个新模块的设计文档 | [MOM3.0_模块设计模板.md](./MOM3.0_模块设计模板.md)（如果不存在请先建模板） |
| 看 V2.0 重写前内容保留清单 | [MOM3.0_WMS_V2.0_重写前内容保留清单.md](./MOM3.0_WMS_V2.0_重写前内容保留清单.md)（其他 module 保留清单待补）|
| 看 V2.0 升级批量计划 | [MOM3.0_V2.0_批量推广计划.md](./MOM3.0_V2.0_批量推广计划.md) |
| 看 V2.0 整体 audit 报告 | [MOM3.0_V2.0_整体audit报告.md](./MOM3.0_V2.0_整体audit报告.md) |
| 看 module 依赖矩阵 | [MOM3.0_模块依赖矩阵.md](./MOM3.0_模块依赖矩阵.md) |
| 看状态字段统一方案 | [MOM3.0_状态字段统一方案.md](./MOM3.0_状态字段统一方案.md) |
| 看旧设计（已废弃）| [archive/](./archive/)（只读） |

---

## 4. 文档状态

| 文档类型 | 状态 | 维护人 |
|---------|------|--------|
| MOM3.0_主设计文档.md | ✅ 活跃 | 架构组 |
| MOM3.0_UI设计规范.md | ✅ 活跃 | 前端架构 |
| MOM3.0_技术架构文档.md | ✅ 活跃 | 后端架构 |
| MOM3.0_模块设计模板.md | ✅ 活跃 | 架构组 |
| MOM3.0_附录.md | ✅ 活跃 | 架构组 |
| **MOM3.0_状态字段统一方案.md** | ✅ 活跃（V2.0 新增）| 架构组 |
| **MOM3.0_V2.0_批量推广计划.md** | ✅ 活跃（V2.0 新增）| 架构组 |
| **MOM3.0_V2.0_整体audit报告.md** | ✅ 活跃（V2.0 新增）| 架构组 |
| **MOM3.0_模块依赖矩阵.md** | ✅ 活跃（V2.0 新增）| 架构组 |
| **MOM3.0_WMS_V2.0_重写前内容保留清单.md** | ✅ 活跃（V2.0 模板）| 架构组 |
| MOM3.0_<模块名>模块设计文档.md × 16 | 🟡 9 V2.0 + 7 V1.x | 各模块 owner |
| MOM3.0_<模块名>前端设计文档.md × 9 | ✅ 活跃 | 前端 owner |
| MOM3.0_SAP_Gap_Analysis_and_Development_Plan.md | ✅ 活跃 | 架构组 |
| DEVELOPMENT_STATUS.md | ✅ 活跃 | 测试组 |
| TODO.md | ✅ 活跃 | 全员 |
| rules/ | ✅ 活跃 | 架构组 |
| superpowers/plans/ | ✅ 活跃（开发计划快照） | 项目经理 |
| archive/ | ⛔ 只读 | - |

### 4.5 V2.0 升级进度

**进度**：14 / 14 = 100% module 已 V2.0（2026-07-02 起开始批量升级,**2026-07-03 16:25 1 下午完成**,原计划 3 周）

| Module | 旧版行数 | V2.0 行数 | Commit | 状态 |
|--------|---------|---------|--------|------|
| M03 MES 生产执行 | 2229 (旧) | 730 | `0f09e79` | ✅ V2.0 |
| M04 APS 计划 | 217 | 937 | `0f84028` | ✅ V2.0 |
| M07 WMS 仓储 | 6352 | 892 | `874f517` | ✅ V2.0 |
| M05 BPM 流程 | 1737 | 768 | `80a4937` | ✅ V2.0 |
| M16 SCP 供应链 | 1594 | 724 | `a9bacc5` | ✅ V2.0 |
| M05 质量 QMS | 1421 | 712 | `4721e1c` | ✅ V2.0 |
| M06 EAM 设备 | 737 | 694 | `78ceda2` | ✅ V2.0 |
| M02 MDM 主数据 | 159 | 750 | `2807f51` | ✅ V2.0 |
| M14 INT 系统集成 | 160 | 750 | `63fbf71` | ✅ V2.0 |
| M08/M10 追溯与数据采集 | 143 | 750 | `f27f4c1` | ✅ V2.0 |
| M09 安灯系统 | 182 | 750 | `3e24b8a` | ✅ V2.0 |
| M11/M12 实验室 + 量检具 | 191 | 750 | `f8b5ded` | ✅ V2.0 |
| M07 结算 | 192 | 750 | `9786e8b` | ✅ V2.0 |
| M15 报表 + 大屏 | 139 | 750 | `2a431c5` | ✅ V2.0 |
| M01 系统管理 | — | — | — | ⏳ 横切模块(无独立文档) |
| M13 AI 质检 | — | — | — | ⏳ V1.x(占位/无独立设计文档) |
| M03 扩展 eSOP | — | — | — | ✅ 合并到 MES V2.0 § 1.2 |
| M12 器具容器 | — | — | — | ✅ 合并到 LAB V2.0 § 12.3 |

**V2.0 模板特点**：
- 13 章节 Arc42 简化版 + MESA 11 项 + ISA-95/88 + IATF 16949
- 4 类 Mermaid 图( flowchart / sequenceDiagram / stateDiagram-v2 / erDiagram)
- § 6.1.4 字段类型说明强制引用 `mdm_status_dict`
- § 13.1 CHANGELOG 强制记录 V1→V2
- 模板验证：通过 MES + APS 双 sample 验证可推广

**计划节奏**：详见 [MOM3.0_V2.0_批量推广计划.md](./MOM3.0_V2.0_批量推广计划.md) — 第 1 批 P0 5/5 完成, P1 进行中, 3 周内完成 16/16。

---

## 5. 命名约定

| 类型 | 格式 | 示例 |
|------|------|------|
| 模块设计文档 | `MOM3.0_<模块名>模块设计文档.md` | `MOM3.0_质量模块设计文档.md` |
| 前端设计文档 | `MOM3.0前端_<模块名>模块设计文档.md` | `MOM3.0前端_QMS质量管理模块设计文档.md` |
| 模板 | `MOM3.0_<文档类型>模板.md` | `MOM3.0_模块设计模板.md` |
| 总览 | `MOM3.0_<文档类型>文档.md` | `MOM3.0_主设计文档.md` |
| 调研/审计 | `<主题>-<类型>-<日期>.md` | `MOM3.0-design-doc-audit-2026-07-01.md` |

详见 [DOCUMENTATION_GUIDE.md § 4](./DOCUMENTATION_GUIDE.md)。

---

## 6. 文档维护规约（精简）

**核心 5 条**：

1. **改代码必改文档**：修改任何 API、路由、模型、状态机 → 必须同步对应模块的设计文档；CI 检查
2. **文档与代码同源同 PR**：文档 PR 与代码 PR 必须联动 review
3. **图表用 Mermaid/PlantUML**：禁止 ASCII art、截图替代、Word 框图
4. **废弃文档 mv 到 archive/**：不要堆在活跃目录
5. **每篇文档都有修订日期**：放在文档第 1 行

完整规约见 [DOCUMENTATION_GUIDE.md](./DOCUMENTATION_GUIDE.md)。

---

## 7. 相关资源

- 项目仓库：<https://github.com/ninghonggang/mom3.0>
- 操作手册：暂未拆分（计划放 `manual/` 目录）
- 历史归档：[archive/](./archive/)（SFMS3.0 时代 / MOM 3.0 V1 设计）
- 调研报告：[research/](./research/)（行业最佳实践、审计与改进建议）

---

**维护人**：MOM 3.0 架构组  
**联系**：通过 GitHub Issues