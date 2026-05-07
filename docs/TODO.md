# MOM3.0 未完成工作清单

**版本**: V1.1 | **更新日期**: 2026-05-05 | **基于**: 代码审查 + API测试

---

## P0 紧急问题（影响核心功能）✅ 已全部修复

| # | 问题 | 文件 | 状态 |
|---|------|------|------|
| P0-1 | BOM详情返回"invalid id" | `handler/mdm/bom.go` | ✅ 已修复 |
| P0-2 | 工艺路线详情返回"无效的ID" | `handler/mes/process.go` | ✅ 已修复 |
| P0-3 | 通知公告GetByID返回"record not found" | `handler/system/notice.go` | ✅ 已修复 |

---

## P1 重要问题（影响业务流程）

| # | 问题 | 文件 | 说明 | 状态 |
|---|------|------|------|------|
| P1-1 | 齐套检查逻辑简化 | `service/mes.go:803` | 简化处理直接设READY，需实际逻辑 | 待实现 |
| P1-2 | ERP同步调用未实现 | `service/erp_sync.go` | TODO注释，需对接ERP | 待实现 |
| P1-3 | 告警邮件发送未实现 | `service/alert.go:307` | TODO注释，需SMTP配置 | 待实现 |
| P1-4 | 首末件检验超期提醒未实现 | `service/first_last_inspect.go:50` | TODO注释 | 待实现 |
| P1-5 | QAD系统接口调用未实现 | `service/scp_qad.go` | TODO注释，需对接QAD | 待实现 |
| P1-6 | 退出登录逻辑 | `handler/system/auth.go:106` | JWT无状态，客户端删除token即可 | ✅ 设计如此 |

---

## P2 待优化（体验/性能问题）

| # | 问题 | 文件 | 说明 | 状态 |
|---|------|------|------|------|
| P2-1 | BOM页面操作按钮未实现 | `views/mdm/BomList.vue:95-97` | 前端需对接后端API | 待前端修复 |
| P2-2 | 物料追溯API调用未实现 | `views/mes/MaterialTrace.vue:254` | TODO注释 | 待前端修复 |
| P2-3 | 工单列表API路径预留 | `views/aps/WorkOrderList.vue:181` | TODO注释 | 待前端修复 |
| P2-4 | AGV回调地址未配置 | `service/agv.go:137` | 需在config.yaml添加配置 | 待配置 |
| P2-5 | 多处"invalid id"错误响应 | 多个handler文件 | 部分handler仍用旧格式 | 部分修复 |

---

## 功能缺失清单

| # | 功能 | 优先级 | 说明 |
|---|------|--------|------|
| F1 | 智能排程算法 | P1 | APS自动排程引擎未实现 |
| F2 | 甘特图可视化 | P2 | 排程结果可视化展示 |
| F3 | 移动端H5页面 | P1 | 完整的移动端UI |
| F4 | 数字孪生 | P2 | 3D工厂可视化 |
| F5 | 低代码平台 | P2 | 表单/流程设计器 |

---

## 已验证正常的功能

| # | 功能 | 验证结果 |
|---|------|---------|
| V1 | 移动报工product_code/product_name | ✅ 已修复 |
| V2 | IDOC接口 | ✅ 正常 |
| V3 | 移动报工CRUD | ✅ 正常 |
| V4 | 租户管理API | ✅ 正常 |
| V5 | 设备台账API | ✅ 正常 |
| V6 | 设备维修API | ✅ 正常 |
| V7 | WMS库存API | ✅ 正常 |
| V8 | IQC检验API | ✅ 正常 |
| V9 | 工单管理API | ✅ 正常 |
| V10 | 采购订单API | ✅ 正常 |
| V11 | 销售订单API | ✅ 正常 |
| V12 | BOM详情API | ✅ 已修复 |
| V13 | 工艺路线API | ✅ 已修复 |
| V14 | 通知公告API | ✅ 已修复 |

---

## P0修复详情 (2026-05-05)

### P0-1: BOM详情错误处理
- 文件: `service/bom.go`
- 修复: 添加gorm.ErrRecordNotFound处理，返回" BOM不存在"

### P0-2: 工艺路线详情错误处理
- 文件: `service/mes_process.go`, `handler/mes/process.go`
- 修复: 添加gorm.ErrRecordNotFound处理，统一使用response包

### P0-3: 通知公告错误处理
- 文件: `service/system_ext.go`
- 修复: 添加gorm.ErrRecordNotFound处理，返回"通知公告不存在"

---

**文档状态**: 本文档为工作清单，持续更新
