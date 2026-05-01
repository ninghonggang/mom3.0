# 峰梅动力 MOM3.0 - Architecture Context Hub

> Last updated: 2026-04-21
> AI Assistant: Hermes Agent

---

## 项目概览

**项目名称**: 峰梅动力 MOM 3.0 (制造运营管理系统)
**技术栈**: React (前端 mom-web) + Go (后端 mom-server)
**测试框架**: Playwright (UI) + Go test (后端)
**工作目录**: `/data/mom3.0`

---

## Architecture Skills (已加载)

| 技能 | 路径 | 用途 |
|------|------|------|
| Python Backend Architecture Review | `~/.hermes/skills/development/python-architecture-review/` | 后端架构审查 |
| UI Design Review | `~/.hermes/skills/design/ui-design-review/` | 前端/UI 审查 |

---

## Context Hub 工作流程

### 1. Feature 完成后的代码审查

当用户完成一个 Feature 后，AI Assistant 通过 Claude Code 调用 gstack/review：

```bash
cd /data/mom3.0 && claude -p "使用 /review 技能审查最近的代码变更，输出结构化审查报告" --max-turns 20
```

**审查流程：**
1. Hermes 调用 Claude Code（使用已安装的 gstack/review 技能）
2. Claude Code 执行完整代码审查
3. 输出审查报告

### 2. 第三方库文档获取

当遇到不熟悉的第三方库时：

1. 使用 `fetch` MCP 工具抓取官方文档
2. 总结接入方案，包括：
   - 库用途和适用场景
   - 安装/依赖方式
   - 核心 API 速查
   - 代码示例
   - 注意事项和已知问题

---

## gstack/review 技能（通过 Claude Code 调用）

| 属性 | 值 |
|------|------|
| 路径 | `~/.claude/skills/review/` |
| 来源 | gstack (Claude Code native) |
| 用途 | PR/代码审查 |
| 调用方式 | 通过 `claude -p` 在终端调用 |

**注意:** 此技能通过 Claude Code 隔离调用，安全性由 Claude Code 沙盒保障。

---

## 项目规范 (来自 rules/)

### 目录结构

```
mom3.0/
├── ARCHITECTURE.md          # 本文件 - Context Hub
├── mom-web/                 # React 前端
│   ├── src/
│   ├── components/          # 组件 (参考 rules/components/)
│   └── ...
├── mom-server/              # Go 后端
│   ├── cmd/
│   ├── internal/
│   └── ...
├── tests/                   # Playwright 测试
├── docs/                    # 文档
└── rules/                   # 项目规范
    ├── coding-style/        # 编码风格
    ├── components/          # 组件规范
    ├── patterns/            # 设计模式
    └── workflow/            # 工作流
```

### 设计原则

- **前端**: React + TypeScript，遵循 components/ 规范
- **后端**: Go clean architecture，遵循 patterns/ 规范
- **测试**: Playwright E2E + Go unit test

---

## 第三方库接入记录

> 记录已接入的第三方库和使用方案

| 库名 | 版本 | 用途 | 接入日期 | 备注 |
|------|------|------|----------|------|
| @playwright/mcp | latest | 前端 E2E 测试 | 2026-04-21 | MCP server |
| @modelcontextprotocol/server-github | latest | GitHub API 集成 | 2026-04-21 | MCP server |

---

## Notes

- Project Memory Skill 未安装，相关上下文通过本文件维护
- gstack/review — 通过 Claude Code 调用（`~/.claude/skills/review/`），安全性由 Claude Code 沙盒保障
- External Doc Searcher 未安装，文档获取通过 `fetch` MCP 或 `web_search` 工具

---

*本文件由 Hermes Agent 维护，每次架构变更后更新*
