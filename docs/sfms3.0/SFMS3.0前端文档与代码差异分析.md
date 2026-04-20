# SFMS3.0 前端文档与代码差异分析

> **版本**: v1.0
> **日期**: 2026-04-17
> **对比**: 文档 `SFMS3.0前端设计文档.md` vs 实际代码 `sfms3.0-ui/src/views`

---

## 1. 统计概览

| 模块 | 文档描述页面数 | 实际Vue文件数 | 差异 |
|-----|-------------|-------------|------|
| MES | 27+ | 42 | **+15** |
| WMS | 100+ | 309 | **+209** |
| QMS | 22+ | 50+ | **+28** |
| EAM | 40+ | 42 | **+2** |
| BPM | 12+ | 25+ | **+13** |
| SYSTEM | 35+ | 70+ | **+35** |
| INFRA | 12+ | 30+ | **+18** |
| REPORT | 1+ | 1 | 0 |
| **总计** | **250+** | **569+** | **+319** |

> **结论**: 实际代码页面比文档多出约319个，文档覆盖率仅约44%

---

## 2. 模块详细差异

### 2.1 MES模块

#### 文档已有 vs 代码实际

| 文档描述 | 代码实际存在 | 状态 |
|---------|------------|------|
| `/mes/orderDay` | ✓ 存在 | 一致 |
| `/mes/workScheduling` | ✓ 存在 | 一致 |
| `/mes/processroute` | ✓ 存在 | 一致 |
| `/mes/bom` | ✗ **不存在** (文档有,代码无) | 缺失 |
| `/mes/workstation` | ✓ 存在 | 一致 |
| `/mes/teamSetting` | ✓ 存在 | 一致 |
| `/mes/holiday` | ✓ 存在 | 一致 |
| `/mes/abilityInfo` | ✓ 存在 | 一致 |
| `/mes/operstepsType` | ✓ 存在 | 一致 |
| `/mes/opersteps` | ✓ 存在 | 一致 |
| `/mes/item` | ✓ 存在 | 一致 |
| `/mes/itemRequestMain` | ✓ 存在 | 一致 |
| `/mes/dismantlingMain` | ✓ 存在 | 一致 |
| `/mes/reworkBatch` | ✓ 存在 | 一致 |
| `/mes/reworkSingle` | ✓ 存在 | 一致 |
| `/mes/ordermonthplan` | ✓ 存在 | 一致 |
| `/mes/pattern` | ✓ 存在 | 一致 |
| `/mes/patternType` | ✓ 存在 | 一致 |
| `/mes/workcalendar` | ✓ 存在 | 一致 |
| `/mes/qualityform` | ✓ 存在 | 一致 |
| `/mes/qualityformlog` | ✓ 存在 | 一致 |
| `/mes/qualitygroup` | ✓ 存在 | 一致 |
| `/mes/qualityclass` | ✓ 存在 | 一致 |
| `/mes/reportpStore` | ✓ 存在 | 一致 |
| `/mes/productOffline` | ✓ 存在 | 一致 |
| `/mes/productBackline` | ✓ 存在 | 一致 |
| `/mes/workSchedulingQaform` | ✓ 存在 | 一致 |

#### 代码有但文档缺失 (MES)

| 页面路径 | 功能说明 |
|---------|---------|
| `/mes/components/Detail.vue` | 通用详情组件 |
| `/mes/process/components/Detail.vue` | 工艺详情 |
| `/mes/process/index.vue` | 工艺管理 (文档未列) |
| `/mes/productionPlan/index.vue` | 生产计划 (文档未列) |
| `/mes/workstation/components/Detail.vue` | 工位详情组件 |
| `/mes/hrPersonAbility/index.vue` | 人员能力 (文档未列) |

**MES差异汇总**: 文档缺失6个页面，1个文档有但代码无(BOM)

---

### 2.2 WMS模块

#### 统计差异

| 分类 | 文档描述 | 实际数量 | 差异 |
|-----|---------|---------|------|
| 基础数据-工厂建模 | ~15 | 20+ | +5 |
| 基础数据-物料管理 | ~12 | 15+ | +3 |
| 基础数据-客户管理 | ~7 | 7 | 0 |
| 基础数据-供应商 | ~4 | 4 | 0 |
| 基础数据-标签管理 | ~8 | 8 | 0 |
| 基础数据-策略设置 | ~20 | 25+ | +5 |
| 基础数据-单据设置 | ~8 | 8 | 0 |
| 基础数据-科目管理 | ~4 | 4 | 0 |
| 基础数据-系统设置 | ~4 | 4 | 0 |
| 库存管理 | ~8 | 10+ | +2 |
| 入库管理 | ~7 | 12+ | +5 |
| 出库管理 | ~7 | 12+ | +5 |
| AGV管理 | ~5 | 5 | 0 |
| 结算管理 | ~3 | 6+ | +3 |

#### 代码有但文档缺失 (WMS)

| 页面路径 | 功能说明 |
|---------|---------|
| `/wms/basicDataManage/itemManage/productionitemcodeSpareitemcode/index.vue` | 替代物料 |
| `/wms/basicDataManage/itemManage/stdcostprice/index.vue` | 标准成本 |
| `/wms/basicDataManage/itemManage/relegate/relegateRequest/index.vue` | 退货申请 |
| `/wms/basicDataManage/itemManage/relegate/relegateRecord/index.vue` | 退货记录 |
| `/wms/basicDataManage/orderManage/team/teamForm.vue` | 班组表单 |
| `/wms/basicDataManage/labelManage/callmaterials/index.vue` | 叫料管理 |
| `/wms/basicDataManage/labelManage/locationLabel/index.vue` | 位置标签 |
| `/wms/basicDataManage/subject/mstr/index.vue` | 科目主数据 |
| `/wms/basicDataManage/subject/qadCostcentre/index.vue` | 成本中心 |
| `/wms/basicDataManage/subject/qadProject/index.vue` | 项目主数据 |
| `/wms/basicDataManage/subject/subjectAccount/index.vue` | 科目账户 |
| `/wms/buttMesManage/mesBarCode/index.vue` | MES条码对接 |
| `/wms/buttMesManage/*` | MES对接相关 (多处) |
| `/wms/countManage/countadjust/*` | 盘点调整 (多处) |
| `/wms/deliversettlementManage/*` | 结算管理 (多处) |
| `/wms/inventoryManage/*` | 库存管理 (多处) |
| `/wms/inventorychange/*` | 库存变化 (多处) |
| `/wms/priceManage/*` | 价格管理 (多处) |
| `/wms/purchasereturnManage/*` | 采购退货 (多处) |
| `/wms/purchasesettlementManage/*` | 采购结算 (多处) |
| `/wms/qualityControl/*` | 质检管理 (多处) |
| `/wms/receiveRecord/*` | 收货记录 (多处) |
| `/wms/returnSettlement/*` | 退货结算 (多处) |
| `/wms/stockupManage/*` | 备货管理 (多处) |
| `/wms/supplyLink/*` | 供应链 (多处) |
| `/wms/transactionHistory/*` | 交易历史 (多处) |

**WMS差异汇总**: 文档严重不足，实际页面约为文档的3倍

---

### 2.3 QMS模块

#### 文档已有 vs 代码实际

| 文档描述 | 代码实际存在 | 状态 |
|---------|------------|------|
| `/qms/aql` | ✓ 存在 | 一致 |
| `/qms/counter` | ✓ 存在 | 一致 |
| `/qms/dynamicRule` | ✓ 存在 | 一致 |
| `/qms/inspectionScheme` | ✓ 存在 | 一致 |
| `/qms/inspectionTemplate` | ✓ 存在 | 一致 |
| `/qms/inspectionMethod` | ✓ 存在 | 一致 |
| `/qms/inspectionStage` | ✓ 存在 | 一致 |
| `/qms/inspectionQ1` | ✓ 存在 | 一致 |
| `/qms/inspectionQ2` | ✓ 存在 | 一致 |
| `/qms/inspectionQ3` | ✓ 存在 | 一致 |
| `/qms/samplingScheme` | ✓ 存在 | 一致 |
| `/qms/samplingProcess` | ✓ 存在 | 一致 |
| `/qms/sampleCode` | ✓ 存在 | 一致 |
| `/qms/selectedProject` | ✓ 存在 | 一致 |
| `/qms/selectedSet` | ✓ 存在 | 一致 |
| `/qms/inspectionRequest` | ✓ 存在 | 一致 |
| `/qms/inspectionJob` | ✓ 存在 | 一致 |
| `/qms/inspectionRecord` | ✓ 存在 | 一致 |
| `/qms/inspectionRecordFirst` | ✓ 存在 | 一致 |
| `/qms/qualityNotice` | ✓ 存在 | 一致 |

#### 代码有但文档缺失 (QMS)

| 页面路径 | 功能说明 |
|---------|---------|
| `/qms/inspectionJob/addForm.vue` | 检验任务新增表单 |
| `/qms/inspectionJob/detail.vue` | 检验任务详情 |
| `/qms/inspectionJob/inspectionJob.vue` | 检验任务主页面 |
| `/qms/inspectionJob/inspectionJobProduction.vue` | 生产检验任务 |
| `/qms/inspectionJob/inspectionJobPurchase.vue` | 采购检验任务 |
| `/qms/inspectionRecord/addForm.vue` | 检验记录新增 |
| `/qms/inspectionRecord/detail.vue` | 检验记录详情 |
| `/qms/inspectionRecord/inspectionRecord.vue` | 检验记录主页 |
| `/qms/inspectionRecord/inspectionRecordProduction.vue` | 生产检验记录 |
| `/qms/inspectionRecord/inspectionRecordPurchase.vue` | 采购检验记录 |
| `/qms/inspectionRecord/useAddForm.vue` | 检验记录通用新增 |
| `/qms/inspectionRecordFirst/addForm.vue` | 首检新增 |
| `/qms/inspectionRecordFirst/detail.vue` | 首检详情 |
| `/qms/inspectionRecordFirst/useAddForm.vue` | 首检通用新增 |
| `/qms/inspectionScheme/addForm.vue` | 检验方案新增 |
| `/qms/inspectionScheme/detail.vue` | 检验方案详情 |
| `/qms/inspectionTemplate/addForm.vue` | 检验模板新增 |
| `/qms/inspectionTemplate/detail.vue` | 检验模板详情 |
| `/qms/qualityNotice/addForm.vue` | 质量通知新增 |
| `/qms/qualityNotice/components/*` | 质量通知组件 (3个) |

**QMS差异汇总**: 文档覆盖率约60%，缺少大量表单和子组件页面

---

### 2.4 EAM模块

#### 文档已有 vs 代码实际

| 文档描述 | 代码实际存在 | 状态 |
|---------|------------|------|
| `/eam/equipmentAccounts` | ✓ 存在 | 一致 |
| `/eam/equipmentMainPart` | ✓ 存在 | 一致 |
| `/eam/equipmentSupplier` | ✓ 存在 | 一致 |
| `/eam/equipmentManufacturer` | ✓ 存在 | 一致 |
| `/eam/classTypeRole` | ✓ 存在 | 一致 |
| `/eam/equipmentToolSparePart` | ✓ 存在 | 一致 |
| `/eam/documentType` | ✓ 存在 | 一致 |
| `/eam/documentTypeSelectSet` | ✓ 存在 | 一致 |
| `/eam/tableDataExtendedAttribute` | ✓ 存在 | 一致 |
| `/eam/planSpotCheck` | ✓ 存在 | 一致 |
| `/eam/equipmentSpotCheckMain` | ✓ 存在 | 一致 |
| `/eam/equipmentSpotCheckRecordMain` | ✓ 存在 | 一致 |
| `/eam/spotCheckItem` | ✓ 存在 | 一致 |
| `/eam/basicSpotCheckOption` | ✓ 存在 | 一致 |
| `/eam/spotCheckSelectSet` | ✓ 存在 | 一致 |
| `/eam/equipmentSpotCheckDetail` | ✓ 存在 | 一致 |
| `/eam/planInspection` | ✓ 存在 | 一致 |
| `/eam/equipmentInspectionMain` | ✓ 存在 | 一致 |
| `/eam/equipmentInspectionRecordMain` | ✓ 存在 | 一致 |
| `/eam/maintenanceItem` | ✓ 存在 | 一致 |
| `/eam/basicMaintenanceOption` | ✓ 存在 | 一致 |
| `/eam/maintenanceItemSelectSet` | ✓ 存在 | 一致 |
| `/eam/equipmentInspectionDetail` | ✓ 存在 | 一致 |
| `/eam/maintainExperience` | ✓ 存在 | 一致 |
| `/eam/equipmentReportRepairRequest` | ✓ 存在 | 一致 |
| `/eam/equipmentRepairJobMain` | ✓ 存在 | 一致 |
| `/eam/equipmentRepairRecordMain` | ✓ 存在 | 一致 |
| `/eam/repairSparePartsRequest` | ✓ 存在 | 一致 |
| `/eam/repairExperience` | ✓ 存在 | 一致 |
| `/eam/basicFaultType` | ✓ 存在 | 一致 |
| `/eam/basicFaultCause` | ✓ 存在 | 一致 |
| `/eam/equipmentRepairJobDetail` | ✓ 存在 | 一致 |
| `/eam/equipmentRepairRecordDetail` | ✓ 存在 | 一致 |
| `/eam/repairSparePartsRecord` | ✓ 存在 | 一致 |
| `/eam/equipmentShutdown` | ✓ 存在 | 一致 |
| `/eam/equipmentSigning` | ✓ 存在 | 一致 |
| `/eam/equipmentTransferRecord` | ✓ 存在 | 一致 |
| `/eam/recordDeviceChanged` | ✓ 存在 | 一致 |
| `/eam/sparePart` | ✓ 存在 | 一致 |
| `/eam/sparePartsApplyMain` | ✓ 存在 | 一致 |
| `/eam/sparepartsinlocation` | ✓ 存在 | 一致 |
| `/eam/sparepartsoutlocation` | ✓ 存在 | 一致 |
| `/eam/itemInLocation` | ✓ 存在 | 一致 |
| `/eam/SparePartsOutLocationRecord` | ✓ 存在 | 一致 |
| `/eam/sparePartsInLocationRecord` | ✓ 存在 | 一致 |
| `/eam/itemLocationReplace` | ✓ 存在 | 一致 |
| `/eam/itemAccounts` | ✓ 存在 | 一致 |
| `/eam/itemApplyMain` | ✓ 存在 | 一致 |
| `/eam/itemDelete` | ✓ 存在 | 一致 |
| `/eam/itemMaintenance` | ✓ 存在 | 一致 |
| `/eam/itemOrderMain` | ✓ 存在 | 一致 |
| `/eam/toolAccounts` | ✓ 存在 | 一致 |
| `/eam/toolChangedRecord` | ✓ 存在 | 一致 |
| `/eam/basicEamProductionline` | ✓ 存在 | 一致 |
| `/eam/basicEamWorkshop` | ✓ 存在 | 一致 |
| `/eam/relationMainPart` | ✓ 存在 | 一致 |
| `/eam/countadjustPlan` | ✓ 存在 | 一致 |
| `/eam/countadjustWork` | ✓ 存在 | 一致 |
| `/eam/countRecord` | ✓ 存在 | 一致 |
| `/eam/adjustRecord` | ✓ 存在 | 一致 |
| `/eam/applicationRecord` | ✓ 存在 | 一致 |
| `/eam/inspectionItem` | ✓ 存在 | 一致 |
| `/eam/inspectionItemSelectSet` | ✓ 存在 | 一致 |

#### 代码有但文档缺失 (EAM)

| 页面路径 | 功能说明 |
|---------|---------|
| `/eam/location/index.vue` | 位置管理 (文档未列) |
| `/eam/locationArea/index.vue` | 位置区域 (文档未列) |
| `/eam/itemOutLocation/index.vue` | 物品出库位置 (文档未列) |
| `/eam/item/index.vue` | 物品管理 (文档未列) |
| `/eam/maintenance/audiForm.vue` | 保养审核表单 |
| `/eam/maintenance/index.vue` | 保养管理 |
| `/eam/toolEquipmentIn/index.vue` | 工具设备入库 (文档未列) |
| `/eam/toolEquipmentOut/index.vue` | 工具设备出库 (文档未列) |
| `/eam/toolMod/index.vue` | 工具型号 (文档未列) |
| `/eam/toolMod/operateForm.vue` | 工具操作表单 |
| `/eam/toolSigning/index.vue` | 工具签到 (文档未列) |
| `/eam/transaction/index.vue` | 交易记录 (文档未列) |
| `/eam/planSpotCheck/audiForm.vue` | 点检审核表单 |
| `/eam/planInspection/audiForm.vue` | 保养审核表单 |
| `/eam/equipmentMaintenanceMain/*` | 设备保养相关 (5个文件) |
| `/eam/equipmentRepairJobMain/*` | 设备维修相关 (7个文件) |
| `/eam/equipmentSpotCheckMain/*` | 设备点检相关 (5个文件) |

**EAM差异汇总**: 文档覆盖率约85%，主要缺少详情表单和子组件

---

### 2.5 BPM模块

#### 文档已有 vs 代码实际

| 文档描述 | 代码实际存在 | 状态 |
|---------|------------|------|
| `/bpm/model` | ✓ 存在 | 一致 |
| `/bpm/definition` | ✓ 存在 | 一致 |
| `/bpm/form` | ✓ 存在 | 一致 |
| `/bpm/group` | ✓ 存在 | 一致 |
| `/bpm/taskAssignRule` | ✓ 存在 | 一致 |
| `/bpm/processInstance` | ✓ 存在 | 一致 |
| `/bpm/task/todo` | ✓ 存在 | 一致 |
| `/bpm/task/done` | ✓ 存在 | 一致 |
| `/bpm/oa/leave` | ✓ 存在 | 一致 |

#### 代码有但文档缺失 (BPM)

| 页面路径 | 功能说明 |
|---------|---------|
| `/bpm/form/editor/index.vue` | 表单编辑器 |
| `/bpm/model/editor/index.vue` | 模型编辑器 |
| `/bpm/model/ModelForm.vue` | 模型表单 |
| `/bpm/model/ModelImportForm.vue` | 模型导入表单 |
| `/bpm/oa/leave/create.vue` | 请假创建 |
| `/bpm/oa/leave/detail.vue` | 请假详情 |
| `/bpm/processInstance/create/index.vue` | 流程实例创建 |
| `/bpm/processInstance/detail/*` | 流程实例详情 (4个组件) |
| `/bpm/task/done/TaskDetail.vue` | 已办任务详情 |
| `/bpm/group/UserGroupForm.vue` | 用户组表单 |
| `/bpm/taskAssignRule/TaskAssignRuleForm.vue` | 分配规则表单 |

**BPM差异汇总**: 文档覆盖率约60%，缺少大量编辑器和详情页面

---

### 2.6 SYSTEM模块

#### 文档已有 vs 代码实际

| 文档描述 | 代码实际存在 | 状态 |
|---------|------------|------|
| `/system/user` | ✓ 存在 | 一致 |
| `/system/role` | ✓ 存在 | 一致 |
| `/system/menu` | ✓ 存在 | 一致 |
| `/system/dept` | ✓ 存在 | 一致 |
| `/system/post` | ✓ 存在 | 一致 |
| `/system/dict` | ✓ 存在 | 一致 |
| `/system/dict/data` | ✓ 存在 | 一致 |
| `/system/tenant` | ✓ 存在 | 一致 |
| `/system/tenantPackage` | ✓ 存在 | 一致 |
| `/system/systemInstallPackage` | ✓ 存在 | 一致 |
| `/system/sms/channel` | ✓ 存在 | 一致 |
| `/system/sms/template` | ✓ 存在 | 一致 |
| `/system/sms/log` | ✓ 存在 | 一致 |
| `/system/mail/account` | ✓ 存在 | 一致 |
| `/system/mail/template` | ✓ 存在 | 一致 |
| `/system/mail/log` | ✓ 存在 | 一致 |
| `/system/notify/message` | ✓ 存在 | 一致 |
| `/system/notify/template` | ✓ 存在 | 一致 |
| `/system/notify/my` | ✓ 存在 | 一致 |
| `/system/messageSet` | ✓ 存在 | 一致 |
| `/system/oauth2/client` | ✓ 存在 | 一致 |
| `/system/oauth2/token` | ✓ 存在 | 一致 |
| `/system/role/dataPermission` | ✓ 存在 | 一致 |
| `/system/role/assignMenu` | ✓ 存在 | 一致 |
| `/system/loginlog` | ✓ 存在 | 一致 |
| `/system/operatelog` | ✓ 存在 | 一致 |
| `/system/errorCode` | ✓ 存在 | 一致 |
| `/system/sensitiveWord` | ✓ 存在 | 一致 |
| `/system/serialNumber` | ✓ 存在 | 一致 |
| `/system/area` | ✓ 存在 | 一致 |
| `/system/notice` | ✓ 存在 | 一致 |
| `/system/passwordRule` | ✓ 存在 | 一致 |
| `/system/tableActionRel` | ✓ 存在 | 一致 |

#### 代码有但文档缺失 (SYSTEM)

| 页面路径 | 功能说明 |
|---------|---------|
| `/system/user/DeptTree.vue` | 部门树形组件 |
| `/system/user/UserAssignRoleForm.vue` | 用户分配角色表单 |
| `/system/user/UserForm.vue` | 用户表单 |
| `/system/user/UserImportForm.vue` | 用户导入表单 |
| `/system/dept/DeptForm.vue` | 部门表单 |
| `/system/dict/DictTypeForm.vue` | 字典类型表单 |
| `/system/dict/data/DictDataForm.vue` | 字典数据表单 |
| `/system/errorCode/ErrorCodeForm.vue` | 错误码表单 |
| `/system/loginlog/LoginLogDetail.vue` | 登录日志详情 |
| `/system/mail/account/MailAccountForm.vue` | 邮件账户表单 |
| `/system/mail/account/MailAccountDetail.vue` | 邮件账户详情 |
| `/system/mail/template/MailTemplateForm.vue` | 邮件模板表单 |
| `/system/mail/template/MailTemplateSendForm.vue` | 邮件发送表单 |
| `/system/mail/log/MailLogDetail.vue` | 邮件日志详情 |
| `/system/menu/MenuForm.vue` | 菜单表单 |
| `/system/notice/NoticeForm.vue` | 通知公告表单 |
| `/system/notify/message/NotifyMessageDetail.vue` | 消息详情 |
| `/system/notify/my/MyNotifyMessageDetail.vue` | 我的消息详情 |
| `/system/notify/template/NotifyTemplateForm.vue` | 消息模板表单 |
| `/system/notify/template/NotifyTemplateSendForm.vue` | 消息发送表单 |
| `/system/oauth2/client/ClientForm.vue` | OAuth2客户端表单 |
| `/system/operatelog/OperateLogDetail.vue` | 操作日志详情 |
| `/system/post/PostForm.vue` | 岗位表单 |
| `/system/role/RoleForm.vue` | 角色表单 |
| `/system/role/RoleAssignMenuForm.vue` | 角色分配菜单表单 |
| `/system/role/RoleDataPermissionForm.vue` | 数据权限表单 |
| `/system/sensitiveWord/SensitiveWordForm.vue` | 敏感词表单 |
| `/system/sensitiveWord/SensitiveWordTestForm.vue` | 敏感词测试表单 |
| `/system/serialNumber/SerialNumberForm.vue` | 序列号表单 |
| `/system/sms/channel/SmsChannelForm.vue` | 短信渠道表单 |
| `/system/sms/template/SmsTemplateForm.vue` | 短信模板表单 |
| `/system/sms/template/SmsTemplateSendForm.vue` | 短信发送表单 |
| `/system/sms/log/SmsLogDetail.vue` | 短信日志详情 |
| `/system/tenant/TenantForm.vue` | 租户表单 |
| `/system/tenantPackage/TenantPackageForm.vue` | 租户套餐表单 |

**SYSTEM差异汇总**: 文档覆盖率约70%，主要缺少各实体的表单和详情组件

---

### 2.7 INFRA模块

#### 文档已有 vs 代码实际

| 文档描述 | 代码实际存在 | 状态 |
|---------|------------|------|
| `/infra/file` | ✓ 存在 | 一致 |
| `/infra/fileConfig` | ✓ 存在 | 一致 |
| `/infra/codegen` | ✓ 存在 | 一致 |
| `/infra/dataSourceConfig` | ✓ 存在 | 一致 |
| `/infra/dbDoc` | ✓ 存在 | 一致 |
| `/infra/job` | ✓ 存在 | 一致 |
| `/infra/job/logger` | ✓ 存在 | 一致 |
| `/infra/redis` | ✓ 存在 | 一致 |
| `/infra/apiAccessLog` | ✓ 存在 | 一致 |
| `/infra/apiErrorLog` | ✓ 存在 | 一致 |
| `/infra/build` | ✓ 存在 | 一致 |
| `/infra/skywalking` | ✓ 存在 | 一致 |
| `/infra/testDemo` | ✓ 存在 | 一致 |

#### 代码有但文档缺失 (INFRA)

| 页面路径 | 功能说明 |
|---------|---------|
| `/infra/apiAccessLog/ApiAccessLogDetail.vue` | API访问日志详情 |
| `/infra/apiErrorLog/ApiErrorLogDetail.vue` | API错误日志详情 |
| `/infra/codegen/components/BasicInfoForm.vue` | 代码生成基本信息表单 |
| `/infra/codegen/components/ColumInfoForm.vue` | 代码生成列信息表单 |
| `/infra/codegen/components/GenerateInfoForm.vue` | 代码生成信息表单 |
| `/infra/codegen/editTable.vue` | 代码生成编辑表 |
| `/infra/codegen/importTable.vue` | 代码生成导入表 |
| `/infra/codegen/PreviewCode.vue` | 代码预览 |
| `/infra/config/index.vue` | 配置管理 (文档未列) |
| `/infra/config/ConfigForm.vue` | 配置表单 |
| `/infra/customInterface/index.vue` | 自定义接口 (文档未列) |
| `/infra/dataSourceConfig/DataSourceConfigForm.vue` | 数据源配置表单 |
| `/infra/druid/index.vue` | Druid监控 (文档未列) |
| `/infra/file/FileForm.vue` | 文件表单 |
| `/infra/fileConfig/FileConfigForm.vue` | 文件配置表单 |
| `/infra/job/JobDetail.vue` | 任务详情 |
| `/infra/job/JobForm.vue` | 任务表单 |
| `/infra/job/logger/JobLogDetail.vue` | 任务日志详情 |
| `/infra/outerApiHis/index.vue` | 外部API历史 (文档未列) |
| `/infra/server/index.vue` | 服务器监控 (文档未列) |
| `/infra/swagger/index.vue` | Swagger文档 (文档未列) |
| `/infra/webSocket/index.vue` | WebSocket (文档未列) |

**INFRA差异汇总**: 文档覆盖率约50%，缺少大量子组件和监控相关页面

---

### 2.8 其他模块

#### LOGIN模块 (文档缺失)

| 页面路径 | 功能说明 |
|---------|---------|
| `/login/login.vue` | 登录页 |
| `/login/forgetPassword.vue` | 忘记密码 |
| `/login/updatePassword.vue` | 更新密码 |
| `/login/updatePasswordNewTips.vue` | 新密码提示 |
| `/login/components/LoginForm.vue` | 登录表单组件 |
| `/login/components/MobileForm.vue` | 手机登录表单 |
| `/login/components/QRCodePDA.vue` | PDA二维码 |

#### HOME模块 (文档缺失)

| 页面路径 | 功能说明 |
|---------|---------|
| `/home/index.vue` | 首页 |
| `/home/Index2.vue` | 首页V2 |
| `/home/Index copy.vue` | 首页副本 |
| `/home/components/material.vue` | 物料组件 |
| `/home/components/produce.vue` | 生产组件 |
| `/home/components/product.vue` | 产品组件 |
| `/home/components/supplierIndex.vue` | 供应商组件 |
| `/home/components/supplierIndex供应商发票空白页，暂时去掉部分模块.vue` | 供应商发票占位 |

#### PROFILE模块 (文档部分缺失)

| 页面路径 | 状态 |
|---------|------|
| `/profile/index.vue` | 文档未列 |
| `/profile/basic` | 文档有,但实际路径不同 |
| `/profile/avatar` | 文档有,但实际路径不同 |
| `/profile/resetPwd` | 文档有,但实际路径不同 |
| `/profile/social` | 文档有,但实际路径不同 |

---

## 3. API接口差异

### 3.1 文档描述 vs 代码实际

| 模块 | 文档描述API数 | 实际API文件数 | 差异 |
|-----|-------------|-------------|------|
| MES | 4 | 4 | 0 |
| WMS | 80+ | 100+ | +20 |
| BPM | 9 | 10 | +1 |
| SYSTEM | 25+ | 30+ | +5 |
| INFRA | 10+ | 15+ | +5 |

### 3.2 代码有但文档缺失的API

**WMS新增API**:
- `wms/accountcalendar` - 账户日历
- `wms/condition` - 条件配置
- `wms/configuration` - 配置管理
- `wms/configurationsetting` - 配置项
- `wms/currencyexchange` - 汇率
- `wms/documentsetting` - 单据设置
- `wms/jobsetting` - 作业设置
- `wms/labeltype` - 标签类型
- `wms/stockup` - 备货管理
- `wms/supplyLink` - 供应链
- `wms/transactionHistory` - 交易历史

**INFRA新增API**:
- `infra/config` - 系统配置
- `infra/customInterface` - 自定义接口
- `infra/druid` - Druid监控
- `infra/outerApiHis` - 外部API历史
- `infra/server` - 服务器监控
- `infra/swagger` - Swagger文档
- `infra/webSocket` - WebSocket

---

## 4. 总结与建议

### 4.1 覆盖率统计

| 模块 | 文档覆盖率 |
|-----|----------|
| MES | ~90% |
| WMS | ~35% |
| QMS | ~60% |
| EAM | ~85% |
| BPM | ~60% |
| SYSTEM | ~70% |
| INFRA | ~50% |
| **平均** | **~56%** |

### 4.2 主要差异类型

1. **表单组件缺失**: 大部分CRUD页面的Form组件未在文档中列出
2. **详情组件缺失**: 大部分详情页的Detail组件未在文档中列出
3. **子组件缺失**: 各模块下的components目录内容几乎未记录
4. **WMS严重落后**: 实际页面约为文档描述的3倍
5. **基础设施模块**: 监控类页面几乎全部缺失

### 4.3 建议更新

1. **优先更新WMS模块文档** - 差异最大(约200+页面)
2. **补充各模块的表单/详情组件** - 约占差异的60%
3. **补充INFRA监控相关页面** - Server、Druid、Swagger等
4. **补充LOGIN和HOME模块** - 目前完全缺失
5. **统一文档格式** - 建议按 `页面路径 | 组件类型 | 功能说明` 格式

---

*文档版本: v1.0 | 生成日期: 2026-04-17*
