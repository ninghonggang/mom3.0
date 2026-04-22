# 项目测试规范

## 菜单配置说明
- 本项目菜单是**动态配置**的，不是硬编码
- 菜单配置来源：[API接口地址 / 配置文件路径]
- 菜单数据包含以下字段：[字段说明]
- 项目加载的菜单存在数据库表[sys_menu]中，所有功能页面都应该配置到其中，没有的应该补充进去
## E2E 测试要求
1. 编写测试前，必须先获取完整的菜单配置数据
2. 测试必须覆盖菜单配置中返回的**所有**菜单项
3. 每个菜单项对应的页面都需要执行基础访问测试

## 已上线/已变更的页面
- [页面A] - 路由: /page-a
- [页面B] - 路由: /page-b
- [新页面C] - 路由: /page-c （**本次新增**）

## 待补充的页面（请开发人员填写）
- [ ] 页面D - 路由: /page-d
- [ ] 页面E - 路由: /page-e

<operating_principles>
- **强制思考协议：在输出最终答案前，必须先在 <thinking> 标签内展示完整的推理过程。**
- **严谨性：如果信息不足或存在歧义，直接说明，严禁猜测。**
- Delegate specialized work to the most appropriate agent.
- Prefer evidence over assumptions: verify outcomes before final claims.
- Choose the lightest-weight path that preserves quality.
- Consult official docs before implementing with SDKs/frameworks/APIs.
</operating_principles>
@AGENTS.md
