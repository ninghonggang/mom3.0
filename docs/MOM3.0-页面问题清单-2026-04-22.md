# MOM3.0 页面问题清单

> 生成时间: 2026-04-22
> 测试方法: Playwright E2E 自动化测试
> 测试结果: 84 passed / 6 failed (71 页面加载成功，6个页面有后端问题)

---

## 一、后端 API 未实现（需要开发后端 Handler）

### 1.1 BPM 流程模块

| 菜单路径 | 前端组件 | 后端状态 | 优先级 |
|---------|---------|---------|--------|
| /bpm/process | bpm/ProcessList.vue | ❌ 后端未实现 | P1 |
| /bpm/instance | bpm/InstanceList.vue | ❌ 后端未实现 | P1 |
| /bpm/task | bpm/TaskList.vue | ❌ 后端未实现 | P1 |

### 1.2 APS 高级计划排程

| 菜单路径 | 前端组件 | 后端状态 | 优先级 |
|---------|---------|---------|--------|
| /aps/rolling-config | aps/RollingConfigList.vue | ❌ 后端未实现 | P2 |
| /aps/delivery-analysis | aps/DeliveryAnalysisList.vue | ❌ 后端未实现 | P2 |
| /aps/material-shortage | aps/MaterialShortageList.vue | ❌ 后端未实现 | P2 |
| /aps/shortage-rule | aps/ShortageRuleList.vue | ❌ 后端未实现 | P2 |
| /aps/changeover-matrix | aps/ChangeoverMatrixList.vue | ❌ 后端未实现 | P2 |
| /aps/product-family | aps/ProductFamilyList.vue | ❌ 后端未实现 | P2 |

### 1.3 SCP 供应链管理

| 菜单路径 | 前端组件 | 后端状态 | 优先级 |
|---------|---------|---------|--------|
| /scp/customer-inquiry | scp/CustomerInquiryList.vue | ❌ 后端未实现 | P1 |

### 1.4 Alert 统一告警

| 菜单路径 | 前端组件 | 后端状态 | 优先级 |
|---------|---------|---------|--------|
| /alert/escalation | alert/AlertEscalationList.vue | ❌ 后端未实现 | P2 |
| /alert/notification | alert/AlertNotification.vue | ❌ 后端未实现 | P2 |
| /alert/statistics | alert/AlertStatistics.vue | ❌ 后端未实现 | P2 |

### 1.5 WMS 仓储管理

| 菜单路径 | 前端组件 | 后端状态 | 优先级 |
|---------|---------|---------|--------|
| /wms/transfer | wms/TransferOrderList.vue | ❌ 后端未实现 | P1 |
| /wms/stock-check | wms/StockCheckList.vue | ❌ 后端未实现 | P1 |

### 1.6 EAM 设备管理

| 菜单路径 | 前端组件 | 后端状态 | 优先级 |
|---------|---------|---------|--------|
| /eam/factory | eam/FactoryList.vue | ❌ 后端未实现 | P2 |
| /eam/equipment-org | eam/EquipmentOrgList.vue | ❌ 后端未实现 | P2 |
| /eam/downtime | eam/DowntimeList.vue | ❌ 后端未实现 | P2 |

### 1.7 MES 生产执行

| 菜单路径 | 前端组件 | 后端状态 | 优先级 |
|---------|---------|---------|--------|
| /mes/team | mes/TeamList.vue | ❌ 后端未实现 | P2 |
| /mes/process-routes | mes/ProcessRouteList.vue | ❌ 后端未实现 | P2 |
| /mes/offline | mes/OfflineList.vue | ❌ 后端未实现 | P2 |
| /mes/person-skill | mes/PersonSkillList.vue | ❌ 后端未实现 | P2 |
| /mes/material-trace | mes/MaterialTrace.vue | ❌ 后端未实现 | P2 |

### 1.8 Fin 财务管理

| 菜单路径 | 前端组件 | 后端状态 | 优先级 |
|---------|---------|---------|--------|
| /fin/payment-request | fin/PaymentRequestList.vue | ❌ 后端未实现 | P1 |
| /fin/purchase-settlement | fin/PurchaseSettlementList.vue | ❌ 后端未实现 | P1 |
| /fin/sales-settlement | fin/SalesSettlementList.vue | ❌ 后端未实现 | P1 |

### 1.9 Report 生产报表

| 菜单路径 | 前端组件 | 后端状态 | 优先级 |
|---------|---------|---------|--------|
| /report/production-daily | report/ProductionDailyReport.vue | ❌ 后端未实现 | P2 |
| /report/quality-weekly | report/QualityWeeklyReport.vue | ❌ 后端未实现 | P2 |
| /report/oee | report/OEEReport.vue | ❌ 后端未实现 | P2 |
| /report/delivery | report/DeliveryReport.vue | ❌ 后端未实现 | P2 |
| /report/andon | report/AndonReport.vue | ❌ 后端未实现 | P2 |

### 1.10 Integration 集成管理

| 菜单路径 | 前端组件 | 后端状态 | 优先级 |
|---------|---------|---------|--------|
| /integration/interface-config | integration/InterfaceConfigList.vue | ❌ 后端未实现 | P2 |
| /integration/execution-log | integration/ExecutionLogList.vue | ❌ 后端未实现 | P2 |

### 1.11 AGV 智能物流

| 菜单路径 | 前端组件 | 后端状态 | 优先级 |
|---------|---------|---------|--------|
| /agv/task | agv/TaskList.vue | ❌ 后端未实现 | P2 |
| /agv/device | agv/DeviceList.vue | ❌ 后端未实现 | P2 |
| /agv/location | agv/LocationList.vue | ❌ 后端未实现 | P2 |

### 1.12 Equipment 设备扩展

| 菜单路径 | 前端组件 | 后端状态 | 优先级 |
|---------|---------|---------|--------|
| /equipment/gauge | equipment/GaugeList.vue | ❌ 后端未实现 | P2 |
| /equipment/inspection/plans | equipment/InspectionPlanList.vue | ❌ 后端未实现 | P2 |
| /equipment/inspection/records | equipment/InspectionRecordList.vue | ❌ 后端未实现 | P2 |
| /equipment/inspection/templates | equipment/InspectionTemplateList.vue | ❌ 后端未实现 | P2 |
| /equipment/inspection/defects | equipment/InspectionDefectList.vue | ❌ 后端未实现 | P2 |

### 1.13 Quality 质量扩展

| 菜单路径 | 前端组件 | 后端状态 | 优先级 |
|---------|---------|---------|--------|
| /quality/qrci | quality/QRCIList.vue | ❌ 后端未实现 | P2 |
| /quality/lpa | quality/LPAStandardList.vue | ❌ 后端未实现 | P2 |

---

## 二、组件名不匹配问题

### 2.1 组件名称大小写不匹配

| 数据库菜单 component | 实际 Vue 文件 | 问题 |
|---------------------|--------------|------|
| quality/LPAList.vue | quality/LPAStandardList.vue | 文件名不匹配 |

**解决方案**: 重命名 `LPAStandardList.vue` 为 `LPAList.vue` 或更新数据库

---

## 三、前端已实现但需验证的页面（71个通过）

以下页面前端组件已存在且可正常加载：

### 系统管理 (14)
- /system/user, /system/role, /system/menu, /system/dept
- /system/dict, /system/post, /system/tenant
- /system/login-log, /system/oper-log, /system/config
- /system/ai-config

### MDM 主数据 (11)
- /mdm/material, /mdm/material-category, /mdm/workshop
- /mdm/line, /mdm/workstation, /mdm/mdm-shift
- /mdm/bom, /mdm/operation, /mdm/customer
- /mdm/supplier

### 生产执行 (9)
- /production/sales-order, /production/report, /production/dispatch
- /production/order, /production/issue, /production/return
- /production/code-rule, /production/electronic-sop, /production/flow-card

### 设备管理 (5)
- /equipment, /equipment/check, /equipment/maintenance
- /equipment/repair, /equipment/spare

### WMS 仓储 (3)
- /wms/warehouse, /wms/location, /wms/inventory

### 质量检验 (6)
- /quality/iqc, /quality/ipqc, /quality/fqc
- /quality/oqc, /quality/ncr, /quality/spc

### APS 基础 (4)
- /aps/mps, /aps/mrp, /aps/schedule, /aps/work-center

### 其他 (19)
- /dashboard, /trace/query, /trace/andon, /energy/monitor
- /scp/purchase, /scp/rfq, /scp/supplier-quote, /scp/sales-order
- /alert/rules, /alert/records, /supplier/asn
- /wms/data-point, /wms/scan-log, /wms/receive, /wms/delivery
- /production/kanban, /production/package, /production/order-change
- /production/first-last-inspect, /quality/defect-code, /quality/defect-record
- /quality/inspection-plans, /quality/aql, /quality/dynamic-rule

---

## 四、数据库菜单清理记录

### 4.1 已删除的重复菜单

| 路径 | 删除的ID |
|------|---------|
| /equipment/gauge | 45 |
| /equipment/inspection/defects | 43 |
| /equipment/inspection/plans | 39 |
| /equipment/inspection/records | 41 |
| /equipment/inspection/templates | 37 |
| /mdm/mdm-shift | 208 |

---

## 五、待处理问题汇总

| 优先级 | 模块 | 问题数 | 说明 |
|-------|------|-------|------|
| P0 | 组件名不匹配 | 1 | LPAStandardList.vue 需重命名 |
| P1 | SCP/BPM/WMS/Fin | 9 | 核心业务流程，必须实现 |
| P2 | APS/MES/EAM/Alert/Report | 39 | 重要功能，规划实现 |
| P3 | Equipment/Quality扩展 | 7 | 辅助功能，后续实现 |

**总计待开发 Handler**: 约 55+ 个

---

## 六、开发建议

### Phase 1: 核心流程优先 (P1)
1. **BPM**: Process/Instance/Task - 3个
2. **SCP**: CustomerInquiry - 1个
3. **WMS**: Transfer/StockCheck - 2个
4. **Fin**: PaymentRequest/PurchaseSettlement/SalesSettlement - 3个

### Phase 2: 重要功能 (P2)
1. **APS**: 6个滚动排程相关
2. **MES**: 5个生产执行相关
3. **EAM**: 3个设备管理相关
4. **Alert**: 3个告警相关
5. **Report**: 5个报表相关
6. **Integration**: 2个集成相关
7. **AGV**: 3个物流相关

### Phase 3: 辅助功能 (P3)
1. **Equipment扩展**: 5个量具检验相关
2. **Quality扩展**: 2个QRCI/LPA相关
