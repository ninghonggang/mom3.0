# MOM3.0 项目开发规范

MOM3.0（制造运营管理系统）项目开发规范，基于产品详细设计文档 V1.1。

## 技术栈

| 层级 | 技术 | 版本 |
|------|------|------|
| 前端框架 | Vue 3 + Composition API | 3.4+ |
| UI组件库 | Element Plus | 2.x |
| 状态管理 | Pinia | 2.x |
| 图表库 | ECharts | 5.x |
| 打包工具 | Vite | 5.x |
| 后端语言 | Go | 1.22+ |
| Web框架 | Gin | 1.9+ |
| ORM | GORM | 2.x |
| 主数据库 | PostgreSQL | 16+ |
| 缓存 | Redis | 7.x |
| 文件存储 | MinIO | latest |

## 目录结构

```
rules/
├── README.md                    # 本文件
├── coding-style/               # 编码风格规范
│   ├── README.md              # 编码风格总览
│   ├── go-backend.md          # Go后端规范
│   ├── vue-frontend.md        # Vue3前端规范
│   └── ui-design.md           # UI设计规范
├── patterns/                   # 设计模式与架构
│   ├── README.md
│   ├── backend-patterns.md
│   └── frontend-patterns.md
├── components/                 # 组件规范
│   ├── README.md
│   ├── form-components.md
│   ├── table-components.md
│   └── chart-components.md
└── workflow/                   # 开发流程
    ├── README.md
    └── git-workflow.md
```

## 快速开始

1. **编码前必读**: 先阅读 `coding-style/` 目录了解整体规范
2. **UI设计**: 参考 `coding-style/ui-design.md` 确保界面一致性
3. **组件开发**: 参考 `components/` 目录的组件规范
4. **提交规范**: 提交前参考 `workflow/git-workflow.md`

## 规范优先级

- **CRITICAL**: 必须遵守，违反会导致构建失败或严重bug
- **HIGH**: 强烈建议遵守，提升代码质量
- **MEDIUM**: 推荐遵守，提升可维护性
- **LOW**: 可选遵守，视情况而定

## 核心架构规范

### 后端分层
- **Handler**: HTTP 请求处理、参数校验、响应封装
- **Service**: 业务逻辑、事务管理
- **Repository**: 数据访问、GORM 操作

### 统一 API 响应格式
```json
// 成功
{ "code": 200, "message": "success", "data": {...} }

// 分页
{ "code": 200, "message": "success", "data": { "list": [], "total": 100, "page": 1, "pageSize": 20 } }

// 错误
{ "code": 40001, "message": "错误信息", "data": null }
```

### 认证与权限
- JWT 双令牌机制（Access Token 2小时 + Refresh Token 7天）
- Casbin RBAC 权限模型
- 多租户隔离（tenant_id）

## 与全局规则的关系

本项目规则继承全局规则 (`~/.claude/rules/`)，当两者冲突时，**本项目规则优先**。
