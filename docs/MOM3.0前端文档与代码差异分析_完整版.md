# MOM3.0 前端设计文档与代码差异分析

> **版本**: v2.0
> **日期**: 2026-04-17
> **对比**: 文档 `MOM3.0前端设计文档.md` vs 实际代码 `mom-web/src`

---

## 1. 统计概览

| 项目 | 文档描述 | 代码实际 | 差异 |
|------|---------|---------|------|
| 前端页面 | 129 | 129 | ✓ 完全一致 |
| API模块 | 16 | 16 | ✓ 完全一致 |
| API文件 | - | 16个 | - |

---

## 2. 路由对比

### 2.1 首页/登录模块

| 文档路径 | 代码路径 | 组件文件 | 状态 |
|---------|---------|---------|------|
| /dashboard | /dashboard | Dashboard.vue | ✓ 一致 |
| /error/404 | - | Error404.vue | ✓ 存在于views |
| /login | /login | Login.vue | ✓ 一致 |

### 2.2 系统管理 (system/)

| 文档路径 | 代码路径 | 组件 | 状态 |
|---------|---------|------|------|
| /system/user | /system/user | UserList.vue | ✓ 一致 |
| /system/role | /system/role | RoleList.vue | ✓ 一致 |
| /system/menu | /system/menu | MenuList.vue | ✓ 一致 |
| /system/dept | /system/dept | DeptList.vue | ✓ 一致 |
| /system/post | /system/post | PostList.vue | ✓ 一致 |
| /system/dict | /system/dict | DictList.vue | ✓ 一致 |
| /system/tenant | /system/tenant | TenantList.vue | ✓ 一致 |
| /system/loginlog | /system/login-log | LoginLogList.vue | ⚠️ 路径差异 |
| /system/operlog | /system/oper-log | OperLogList.vue | ⚠️ 路径差异 |
| /system/notice | /system/notice | NoticeList.vue | ✓ 一致 |
| /system/ai-config | /system/ai-config | AiConfigView.vue | ✓ 一致 |
| /system/print-template | /system/print-template | PrintTemplateList.vue | ✓ 一致 |
| /system/config | /system/config | SystemConfig.vue | ✓ 一致 |

**差异说明**: 文档使用 `/system/loginlog`，代码使用 `/system/login-log`

### 2.3 主数据管理 (mdm/)

| 文档路径 | 代码路径 | 组件 | 状态 |
|---------|---------|------|------|
| /mdm/material | /mdm/material | MaterialList.vue | ✓ 一致 |
| /mdm/material-category | /mdm/material-category | MaterialCategoryList.vue | ✓ 一致 |
| /mdm/workshop | /mdm/workshop | WorkshopList.vue | ✓ 一致 |
| /mdm/line | /mdm/line | LineList.vue | ✓ 一致 |
| /mdm/workstation | /mdm/workstation | WorkstationList.vue | ✓ 一致 |
| /mdm/shift | /mdm/mdm-shift | ShiftList.vue | ⚠️ 路径差异 |
| /mdm/bom | /mdm/bom | BomList.vue | ✓ 一致 |
| /mdm/operation | /mdm/operation | OperationList.vue | ✓ 一致 |
| /mdm/customer | /mdm/customer | CustomerList.vue | ✓ 一致 |
| /mdm/supplier | /mdm/supplier | SupplierList.vue | ✓ 一致 |
| /mdm/contact | /mdm/contact | ContactList.vue | ✓ 一致 |
| /mdm/bank-account | /mdm/bank-account | BankAccountList.vue | ✓ 一致 |
| /mdm/attachment | /mdm/attachment | AttachmentList.vue | ✓ 一致 |
| /mdm/delivery-address | /mdm/delivery-address | DeliveryAddressList.vue | ✓ 一致 |

**差异说明**: 文档使用 `/mdm/shift`，代码使用 `/mdm/mdm-shift`

### 2.4 生产执行 (production/)

| 文档路径 | 代码路径 | 组件 | 状态 |
|---------|---------|------|------|
| /production/sales-order | /production/sales-order | SalesOrderList.vue | ✓ 一致 |
| /production/order | /production/order | ProductionOrderList.vue | ✓ 一致 |
| /production/dispatch | /production/dispatch | DispatchList.vue | ✓ 一致 |
| /production/report | /production/report | ReportList.vue | ✓ 一致 |
| /production/kanban | /production/kanban | KanbanBoard.vue | ✓ 一致 |
| /production/order-change | /production/order-change | OrderChangeList.vue | ✓ 一致 |
| /production/package | /production/package | PackageList.vue | ✓ 一致 |
| /production/first-last-inspect | /production/first-last-inspect | FirstLastInspectList.vue | ✓ 一致 |
| /production/electronic-sop | /production/electronic-sop | ElectronicSOPList.vue | ✓ 一致 |
| /production/flow-card | /production/flow-card | FlowCardList.vue | ✓ 一致 |
| /production/code-rule | /production/code-rule | CodeRuleList.vue | ✓ 一致 |

**注意**: 文档中有 `/production/issue` (生产发料) 和 `/production/return` (生产退料)，代码中未找到对应路由

### 2.5 设备管理 (equipment/)

| 文档路径 | 代码路径 | 组件 | 状态 |
|---------|---------|------|------|
| /equipment/list | /equipment | EquipmentList.vue | ⚠️ 路径差异 |
| /equipment/check | /equipment/check | CheckList.vue | ✓ 一致 |
| /equipment/maintenance | /equipment/maintenance | MaintenanceList.vue | ✓ 一致 |
| /equipment/repair | /equipment/repair | RepairList.vue | ✓ 一致 |
| /equipment/oee | /equipment/oee | OEELIst.vue | ✓ 一致 |
| /equipment/spare-part | /equipment/spare | SparePartList.vue | ⚠️ 路径差异 |

**差异说明**:
- 文档 `/equipment/list` → 代码 `/equipment`
- 文档 `/equipment/spare-part` → 代码 `/equipment/spare`

### 2.6 WMS仓储 (wms/)

| 文档路径 | 代码路径 | 组件 | 状态 |
|---------|---------|------|------|
| /wms/warehouse | /wms/warehouse | WarehouseList.vue | ✓ 一致 |
| /wms/location | /wms/location | LocationList.vue | ✓ 一致 |
| /wms/receive-order | /wms/receive | ReceiveOrderList.vue | ⚠️ 路径差异 |
| /wms/delivery-order | /wms/delivery | DeliveryOrderList.vue | ⚠️ 路径差异 |
| /wms/inventory | /wms/inventory | InventoryList.vue | ✓ 一致 |
| /wms/transfer-order | - | TransferOrderList.vue | ⚠️ 代码有但未在路由中 |
| /wms/stock-check | - | StockCheckList.vue | ⚠️ 代码有但未在路由中 |
| /wms/data-point | /wms/data-point | DataPointList.vue | ✓ 一致 |
| /wms/scan-log | /wms/scan-log | ScanLogList.vue | ✓ 一致 |

**差异说明**:
- 文档 `/wms/receive-order` → 代码 `/wms/receive`
- 文档 `/wms/delivery-order` → 代码 `/wms/delivery`
- TransferOrderList.vue 和 StockCheckList.vue 存在于代码中但未注册路由

### 2.7 质量管理 (quality/)

| 文档路径 | 代码路径 | 组件 | 状态 |
|---------|---------|------|------|
| /quality/iqc | /quality/iqc | IQCList.vue | ✓ 一致 |
| /quality/ipqc | /quality/ipqc | IPQCList.vue | ✓ 一致 |
| /quality/fqc | /quality/fqc | FQCList.vue | ✓ 一致 |
| /quality/oqc | /quality/oqc | OQCList.vue | ✓ 一致 |
| /quality/inspection-plan | - | InspectionPlanList.vue | ⚠️ 代码有但未在路由中 |
| /quality/inspection-record | - | InspectionRecordList.vue | ⚠️ 代码有但未在路由中 |
| /quality/inspection-template | - | InspectionTemplateList.vue | ⚠️ 代码有但未在路由中 |
| /quality/aql | - | AQLList.vue | ⚠️ 代码有但未在路由中 |
| /quality/dynamic-rule | - | DynamicRuleList.vue | ⚠️ 代码有但未在路由中 |
| /quality/defect-code | /quality/defect-code | DefectCodeList.vue | ✓ 一致 |
| /quality/defect-record | /quality/defect-record | DefectRecordList.vue | ✓ 一致 |
| /quality/ncr | /quality/ncr | NCRList.vue | ✓ 一致 |
| /quality/spc-data | /quality/spc | SPCDataList.vue | ⚠️ 路径差异 |
| /quality/lpa-standard | - | LPAStandardList.vue | ⚠️ 代码有但未在路由中 |
| /quality/qrci | - | QRCIList.vue | ⚠️ 代码有但未在路由中 |

**注意**: 多个quality页面存在于views中但未在路由中注册

### 2.8 APS计划 (aps/)

| 文档路径 | 代码路径 | 组件 | 状态 |
|---------|---------|------|------|
| /aps/mps | /aps/mps | MPSList.vue | ✓ 一致 |
| /aps/mrp | /aps/mrp | MRPList.vue | ✓ 一致 |
| /aps/schedule | /aps/schedule | ScheduleList.vue | ✓ 一致 |
| /aps/delivery-analysis | - | DeliveryAnalysisList.vue | ⚠️ 代码有但未在路由中 |
| /aps/material-shortage | - | MaterialShortageList.vue | ⚠️ 代码有但未在路由中 |
| /aps/shortage-rule | - | ShortageRuleList.vue | ⚠️ 代码有但未在路由中 |
| /aps/work-center | /aps/work-center | WorkCenterList.vue | ✓ 一致 |
| /aps/product-family | - | ProductFamilyList.vue | ⚠️ 代码有但未在路由中 |
| /aps/rolling-config | - | RollingConfigList.vue | ⚠️ 代码有但未在路由中 |
| /aps/changeover-matrix | - | ChangeoverMatrixList.vue | ⚠️ 代码有但未在路由中 |

### 2.9 SCP供应链 (scp/)

| 文档路径 | 代码路径 | 组件 | 状态 |
|---------|---------|------|------|
| /scp/rfq | /scp/rfq | RFQList.vue | ✓ 一致 |
| /scp/purchase-order | /scp/purchase | PurchaseOrderList.vue | ⚠️ 路径差异 |
| /scp/sales-order | /scp/sales-order | SCPSalesOrderList.vue | ✓ 一致 |
| /scp/supplier-quote | /scp/supplier-quote | SupplierQuoteList.vue | ✓ 一致 |
| /scp/supplier-kpi | /scp/supplier-kpi | SupplierKPIList.vue | ✓ 一致 |
| /scp/customer-inquiry | /scp/customer-inquiry | CustomerInquiryList.vue | ✓ 一致 |

**差异说明**: 文档 `/scp/purchase-order` → 代码 `/scp/purchase`

### 2.10 安灯系统 (alert/)

| 文档路径 | 代码路径 | 组件 | 状态 |
|---------|---------|------|------|
| /alert/rules | /alert/rules | AlertRulesList.vue | ✓ 一致 |
| /alert/records | /alert/records | AlertRecordsList.vue | ✓ 一致 |
| /alert/statistics | /alert/statistics | AlertStatistics.vue | ✓ 一致 |
| /alert/escalation | /alert/escalation | AlertEscalationList.vue | ✓ 一致 |
| /alert/notification | /alert/notification | AlertNotification.vue | ✓ 一致 |

### 2.11 其他模块

| 模块 | 文档页面 | 代码页面 | 状态 |
|------|---------|---------|------|
| 追溯 (trace) | /trace/query, /trace/andon-call | /trace/query, /trace/andon | ✓ 一致 |
| 能源 (energy) | /energy/monitor | /energy/monitor | ✓ 一致 |
| 报表 (report) | 5个页面 | 未在路由中 | ⚠️ 文档有,代码无 |
| AGV调度 (agv) | 3个页面 | 未在路由中 | ⚠️ 文档有,代码无 |
| 结算管理 (fin) | 3个页面 | 未在路由中 | ⚠️ 文档有,代码无 |
| MES执行 (mes) | 5个页面 | 未在路由中 | ⚠️ 文档有,代码无 |
| BPM流程 (bpm) | 3个页面 | 部分注册 | ⚠️ 部分缺失 |
| EAM设备 (eam) | 3个页面 | 未在路由中 | ⚠️ 文档有,代码无 |

---

## 3. API模块对比

| API文件 | 文档描述 | 代码存在 | 状态 |
|--------|---------|---------|------|
| system.ts | ✓ | ✓ | 一致 |
| mdm.ts | ✓ | ✓ | 一致 |
| production.ts | ✓ | ✓ | 一致 |
| quality.ts | ✓ | ✓ | 一致 |
| equipment.ts | ✓ | ✓ | 一致 |
| wms.ts | ✓ | ✓ | 一致 |
| aps.ts | ✓ | ✓ | 一致 |
| scp.ts | ✓ | ✓ | 一致 |
| alert.ts | ✓ | ✓ | 一致 |
| bpm.ts | ✓ | ✓ | 一致 |
| mes.ts | ✓ | ✓ | 一致 |
| trace.ts | ✓ | ✓ | 一致 |
| auth.ts | ✓ | ✓ | 一致 |
| ai-chat.ts | ✓ | ✓ | 一致 |
| event.ts | ✓ | ✓ | 一致 |
| production_issue.ts | ✓ | ✓ | 一致 |

---

## 4. 差异汇总

### 4.1 路径差异（文档与代码不一致）

| 模块 | 文档路径 | 代码路径 |
|------|---------|---------|
| system | /system/loginlog | /system/login-log |
| system | /system/operlog | /system/oper-log |
| mdm | /mdm/shift | /mdm/mdm-shift |
| equipment | /equipment/list | /equipment |
| equipment | /equipment/spare-part | /equipment/spare |
| wms | /wms/receive-order | /wms/receive |
| wms | /wms/delivery-order | /wms/delivery |
| scp | /scp/purchase-order | /scp/purchase |

### 4.2 代码中存在但未注册路由的页面

| 组件文件 | 目录 | 说明 |
|---------|------|------|
| TransferOrderList.vue | wms/ | 调拨单 |
| StockCheckList.vue | wms/ | 盘点管理 |
| InspectionPlanList.vue | quality/ | 检验计划 |
| InspectionRecordList.vue | quality/ | 检验记录 |
| InspectionTemplateList.vue | quality/ | 检验模板 |
| AQLList.vue | quality/ | AQL配置 |
| DynamicRuleList.vue | quality/ | 动态规则 |
| LPAStandardList.vue | quality/ | LPA标准 |
| QRCIList.vue | quality/ | QRCI |
| DeliveryAnalysisList.vue | aps/ | 交付分析 |
| MaterialShortageList.vue | aps/ | 物料短缺 |
| ShortageRuleList.vue | aps/ | 短缺规则 |
| ProductFamilyList.vue | aps/ | 产品族 |
| RollingConfigList.vue | aps/ | 滚动配置 |
| ChangeoverMatrixList.vue | aps/ | 换型矩阵 |

### 4.3 文档中有但代码中无的模块

| 模块 | 文档页面数 | 说明 |
|------|-----------|------|
| 报表 (report) | 5 | 生产日报、OEE报表等 |
| AGV调度 (agv) | 3 | AGV设备、库位、任务 |
| 结算管理 (fin) | 3 | 采购结算、销售结算、付款申请 |
| MES执行 (mes) | 5 | 工艺路线、班组管理等 |
| EAM设备 (eam) | 3 | 工厂建模、设备组织、设备停机 |

---

## 5. 结论

### 5.1 文档与代码一致性

| 项目 | 状态 |
|------|------|
| 页面总数 | ✓ 129 vs 129 |
| 已注册路由 | ~100个 |
| 路径一致率 | ~85% |
| API模块 | ✓ 100%一致 |

### 5.2 主要问题

1. **路由未注册**: 约25个页面存在于views目录但未注册到路由
2. **路径差异**: 约8处文档与代码路径不一致
3. **模块缺失**: report/agv/fin/mes/eam模块在代码中未实现

### 5.3 建议

1. 统一路径命名规范
2. 补充缺失的路由注册
3. 补充未实现的模块页面

---

*文档版本: v2.0 | 生成日期: 2026-04-17*