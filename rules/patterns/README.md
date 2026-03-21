# 设计模式与架构

本目录定义 MOM3.0 项目中使用的设计模式和架构规范。

## 文件清单

- **[backend-patterns.md](./backend-patterns.md)** - 后端设计模式
- **[frontend-patterns.md](./frontend-patterns.md)** - 前端设计模式

## 核心模式

### 后端

1. **分层架构**: Handler → Service → Repository
2. **DDD 领域驱动**: 聚合根、领域事件、限界上下文
3. **事务脚本**: 简单业务使用事务脚本模式

### 前端

1. **组合式 API**: Vue 3 Composition API
2. **Store 模式**: Pinia 状态管理
3. **表驱动开发**: 配置化的列表/表单
