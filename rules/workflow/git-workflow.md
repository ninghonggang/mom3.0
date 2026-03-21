# Git 工作流规范

## 1. 分支策略

### 1.1 分支类型

| 分支 | 用途 | 命名规范 | 生命周期 |
|------|------|----------|----------|
| `master` | 生产环境 | - | 长期 |
| `develop` | 开发环境 | - | 长期 |
| `feature/*` | 新功能开发 | feature/功能名 | 临时 |
| `fix/*` | Bug 修复 | fix/问题描述 | 临时 |
| `refactor/*` | 代码重构 | refactor/范围 | 临时 |
| `hotfix/*` | 紧急修复 | hotfix/问题描述 | 临时 |

### 1.2 分支命名示例

```
feature/user-management
feature/material-import
fix/order-status-error
refactor/api-handlers
hotfix/login-timeout
```

## 2. 提交规范

### 2.1 提交信息格式

```
<type>(<scope>): <subject>

<body>

<footer>
```

### 2.2 Type 类型

| Type | 说明 |
|------|------|
| `feat` | 新功能 |
| `fix` | Bug 修复 |
| `docs` | 文档更新 |
| `style` | 代码格式（不影响功能） |
| `refactor` | 重构（既不是新功能也不是修复） |
| `perf` | 性能优化 |
| `test` | 测试相关 |
| `chore` | 构建过程或辅助工具变动 |
| `ci` | CI 配置文件和脚本变动 |

### 2.3 提交示例

```
feat(user): 添加用户导入功能

- 支持 CSV/Excel 格式批量导入
- 支持模板下载
- 导入结果展示错误详情

Closes #123
```

```
fix(order): 修复工单状态更新失败问题

工单完成后状态未正确更新，导致统计异常。
已添加状态校验和事务保证。

Fixes #456
```

```
refactor(api): 统一 API 响应格式

- 抽取 Response 结构体
- 统一错误处理中间件
- 移除冗余代码
```

## 3. 工作流程

### 3.1 功能开发流程

```bash
# 1. 从 develop 创建功能分支
git checkout develop
git pull origin develop
git checkout -b feature/user-import

# 2. 开发功能，定期 rebase develop
git add .
git commit -m "feat(user): 添加用户导入页面"
git rebase develop

# 3. 功能完成，合并到 develop
git checkout develop
git merge --no-ff feature/user-import
git push origin develop

# 4. 删除功能分支
git branch -d feature/user-import
```

### 3.2 Bug 修复流程

```bash
# 1. 从 master 创建修复分支
git checkout master
git pull origin master
git checkout -b fix/order-status

# 2. 修复问题
git add .
git commit -m "fix(order): 修复状态更新问题"

# 3. 修复完成
git checkout master
git merge --no-ff fix/order-status
git push origin master

# 4. 同步到 develop
git checkout develop
git merge master
git push origin develop
```

## 4. 提交检查

### 4.1 提交前检查清单

- [ ] 代码符合项目编码规范
- [ ] 没有硬编码的敏感信息
- [ ] 已添加必要的注释
- [ ] 已更新相关文档（如果需要）
- [ ] 提交信息清晰描述了变更内容

### 4.2 禁止的提交

```bash
# ❌ 禁止: 无意义的提交信息
git commit -m "asdf"
git commit -m "update"
git commit -m "fix"

# ✅ 正确: 描述性的提交信息
git commit -m "feat(user): 添加用户搜索功能"
git commit -m "fix(order): 修复删除订单时的空指针异常"
```

## 5. 代码审查

### 5.1 审查要点

1. **功能正确性**: 代码是否实现了预期功能
2. **代码质量**: 是否符合编码规范，是否清晰易读
3. **安全性**: 是否有安全漏洞，是否正确处理用户输入
4. **性能**: 是否有性能问题，是否有优化的可能
5. **测试**: 是否有必要的测试用例

### 5.2 审查评论格式

```
# 严重问题 - 必须修复
[suggestion] 这里可以简化...

# 建议 - 可选改进
[nitpick] 这里的命名可以更清晰...

# 提问 - 需要澄清
[question] 这里为什么要这样处理？
```

## 6. 标签使用

### 6.1 版本标签

```bash
# 打标签
git tag -a v1.0.0 -m "Version 1.0.0"

# 推送标签
git push origin v1.0.0

# 查看标签
git tag -l
```

### 6.2 标签命名

| 类型 | 格式 | 示例 |
|------|------|------|
| 版本 | v主版本.次版本.修订 | v1.0.0, v1.1.0 |
| 预发布 | v版本-标识 | v2.0.0-alpha, v2.0.0-beta |
| 构建 | v版本+构建号 | v1.0.0+20230320 |

## 7. 常见操作

### 7.1 撤销操作

```bash
# 撤销未提交的修改
git checkout -- .

# 撤销已暂存的修改
git reset HEAD .

# 撤销最近一次提交（保留修改）
git reset --soft HEAD~1

# 撤销最近一次提交（不保留修改）
git reset --hard HEAD~1
```

### 7.2 合并与变基

```bash
# 合并（保留完整历史）
git merge feature/xxx

# 变基（保持线性历史）
git rebase develop

# 交互式变基（整理提交）
git rebase -i HEAD~3
```

### 7.3 储藏操作

```bash
# 储藏当前修改
git stash

# 储藏时添加说明
git stash save "work in progress"

# 查看储藏列表
git stash list

# 恢复储藏
git stash pop

# 删除储藏
git stash drop
```

## 8. 冲突解决

### 8.1 解决冲突步骤

```bash
# 1. 拉取最新代码
git fetch origin

# 2. 尝试合并
git merge origin/develop

# 3. 解决冲突后
git add .
git commit -m "merge: 解决合并冲突"
```

### 8.2 冲突标记

```
<<<<<<< HEAD
当前分支的代码
=======
要合并分支的代码
>>>>>>> feature/xxx
```

## 9. 工作区管理

### 9.1 多任务切换

```bash
# 储藏当前工作
git stash

# 切换分支
git checkout other-branch

# 恢复工作
git stash pop
```

### 9.2 清理工作区

```bash
# 删除已合并的分支
git branch --merged | grep -v "\*" | xargs -n 1 git branch -d

# 清理远程已删除的本地分支
git fetch --prune
```
