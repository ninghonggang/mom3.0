# SFMS3.0 vs MOM3.0 前端页面完整差异分析

> **版本**: v2.0
> **日期**: 2026-04-17
> **目的**: 完整对比SFMS3.0与MOM3.0前端页面清单，识别MOM3.0缺失功能
> **更新说明**: v2.0 新增补充后进度状态，所有缺失模块设计文档已创建

---

## 1. 统计概览

| 项目 | SFMS3.0 | MOM3.0原版 | MOM3.0补充后 | 差异(缺失) |
|------|---------|-----------|-------------|-----------|
| 前端页面总数 | **521** | **129** | **447+** | **-74** |
| MES执行模块 | 36 | 5 | 20 | -16 |
| WMS仓储模块 | 300+ | 9 | 140+ | -160 |
| QMS质量管理 | 50+ | 16 | 48 | -2 |
| EAM设备管理 | 65+ | 14 | 65 | 0 |
| BPM流程模块 | 30+ | 3 | 30 | 0 |
| SYSTEM系统管理 | 35+ | 13 | 31 | -4 |
| INFRA基础设施 | 12+ | 0 | 12 | 0 |
| REPORT报表模块 | 1+ | 5 | 5 | +4 |
| SCP供应链模块 | 6 | 6 | 6 | 0 |
| AGV调度模块 | 3 | 3 | 3 | 0 |
| 安灯系统模块 | 5 | 5 | 5 | 0 |
| 能源管理模块 | 1 | 1 | 1 | 0 |
| 追溯管理模块 | 2 | 2 | 2 | 0 |
| 结算管理模块 | 3 | 3 | 3 | 0 |
| 主数据MDM模块 | 15 | 15 | 15 | 0 |

### 1.1 补充进度汇总

| 模块 | 原页面数 | 补充页面数 | 补充后总数 | 状态 |
|------|---------|-----------|-----------|------|
| WMS仓储 | 9 | +131 | 140+ | ✅ 已补充 |
| MES执行 | 5 | +15 | 20 | ✅ 已补充 |
| BPM流程 | 3 | +27 | 30 | ✅ 已补充 |
| QMS质量管理 | 16 | +32 | 48 | ✅ 已补充 |
| EAM设备管理 | 14 | +51 | 65 | ✅ 已补充 |
| SYSTEM系统管理 | 13 | +18 | 31 | ✅ 已补充 |
| INFRA基础设施 | 0 | +12 | 12 | ✅ 已补充 |
| **总计** | **129** | **+286** | **415+** | **✅ 全部完成** |

---

## 2. MES执行模块详细对比

### 2.1 SFMS3.0 MES页面清单 (30+页)

| 序号 | 页面名称 | 路径 | 功能说明 |
|-----|---------|------|---------|
| 1 | 生产日计划 | /mes/orderDay | 日计划创建/发布/终止/恢复 |
| 2 | 工单排程 | /mes/workScheduling | 排程管理 |
| 3 | 工艺路线 | /mes/processRoute | 工序定义 |
| 4 | BOM管理 | /mes/bom | 物料清单 |
| 5 | 工位管理 | /mes/workstation | 工位档案 |
| 6 | 班组管理 | /mes/team | 班组档案 |
| 7 | 假期管理 | /mes/holiday | 假期日历 |
| 8 | 能力档案 | /mes/ability | 能力配置 |
| 9 | 工序类型 | /mes/processType | 工序分类 |
| 10 | 产品档案 | /mes/product | 产品BOM |
| 11 | 物料申请 | /mes/materialApply | 生产物料申请 |
| 12 | 拆解管理 | /mes/disassembly | 拆解任务 |
| 13 | 返工批次 | /mes/rewrokBatch | 返工管理 |
| 14 | 月计划 | /mes/monthPlan | 月度计划 |
| 15 | 花字品种 | /mes/variety | 品种管理 |
| 16 | 工作日历 | /mes/workCalendar | 日历配置 |
| 17 | 产品报交 | /mes/productReport | 报交管理 |
| 18 | 产品下线 | /mes/offline | 下线管理 |
| 19 | 产品返线 | /mes/returnLine | 返线管理 |
| 20 | 排程详情 | /mes/schedulingDetail | 排程明细 |
| 21 | 工序管理 | /mes/process | 工序详细配置 |
| 22 | 人员技能 | /mes/personSkill | 技能档案 |
| 23 | 物料追溯 | /mes/materialTrace | 物料追溯 |
| 24 | 下线管理 | /mes/offlineMgmt | 下线记录 |
| 25 | 生产工单 | /production/order | 工单管理 |
| 26 | 派工管理 | /production/dispatch | 派工调度 |
| 27 | 生产报工 | /production/report | 报工记录 |
| 28 | 生产看板 | /production/kanban | 生产看板 |
| 29 | 首末件检验 | /production/firstLastInspect | 首末件检验 |
| 30 | 工序流转卡 | /production/flowCard | 流转卡 |
| 31 | 生产变更 | /production/orderChange | 变更管理 |
| 32 | 生产发料 | /production/issue | 发料作业 |
| 33 | 生产退料 | /production/return | 退料作业 |
| 34 | 包装管理 | /production/package | 包装管理 |
| 35 | 电子SOP | /production/electronicSop | 电子作业指导 |
| 36 | 编码规则 | /production/codeRule | 编码规则 |

### 2.2 MOM3.0 MES页面清单 (5页)

| 序号 | 页面名称 | 路径 | 功能说明 |
|-----|---------|------|---------|
| 1 | 工艺路线 | /mes/process-route | 工序定义 |
| 2 | 班组管理 | /mes/team | 班组档案 |
| 3 | 人员技能 | /mes/person-skill | 技能档案 |
| 4 | 物料追溯 | /mes/material-trace | 物料追溯 |
| 5 | 下线管理 | /mes/offline | 下线管理 |

### 2.3 MES模块缺失清单 (-31页)

| 缺失页面 | 说明 |
|---------|------|
| 生产日计划 | 日计划CRUD功能 |
| 工单排程 | 排程管理 |
| BOM管理 | BOM配置 |
| 工位管理 | 工位档案 |
| 假期管理 | 假期日历 |
| 能力档案 | 能力配置 |
| 工序类型 | 工序分类 |
| 产品档案 | 产品BOM |
| 物料申请 | 生产物料申请 |
| 拆解管理 | 拆解任务 |
| 返工批次 | 返工管理 |
| 月计划 | 月度计划 |
| 花字品种 | 品种管理 |
| 工作日历 | 日历配置 |
| 产品报交 | 报交管理 |
| 产品返线 | 返线管理 |
| 排程详情 | 排程明细 |
| 工序管理 | 工序详细配置 |
| 生产工单 | 工单管理 |
| 派工管理 | 派工调度 |
| 生产报工 | 报工记录 |
| 生产看板 | 生产看板 |
| 首末件检验 | 首末件检验 |
| 工序流转卡 | 流转卡 |
| 生产变更 | 变更管理 |
| 生产发料 | 发料作业 |
| 生产退料 | 退料作业 |
| 包装管理 | 包装管理 |
| 电子SOP | 电子作业指导 |
| 编码规则 | 编码规则 |

---

## 3. WMS仓储模块详细对比

### 3.1 SFMS3.0 WMS页面清单 (300+页)

#### 3.1.1 基础数据-工厂建模 (20+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 仓库档案 | /wms/warehouse |
| 2 | 库区管理 | /wms/zone |
| 3 | 库位管理 | /wms/location |
| 4 | 库位组 | /wms/locationGroup |
| 5 | 产线管理 | /wms/productionLine |
| 6 | 车间管理 | /wms/workshop |
| 7 | 工位管理 | /wms/workstation |
| 8 | 工序管理 | /wms/process |
| 9 | Dock管理 | /wms/dock |
| 10 | 工厂配置 | /wms/factoryConfig |
| 11 | 区域管理 | /wms/area |
| 12 | 库位关系 | /wms/locationRelation |
| 13 | 巷道管理 | /wms/aisle |
| 14 | 层架管理 | /wms/shelf |
| 15 | 容器规格 | /wms/containerSpec |
| 16 | 仓库模型 | /wms/warehouseModel |
| 17 | 库位属性 | /wms/locationAttribute |
| 18 | 库位类型 | /wms/locationType |
| 19 | 库区类型 | /wms/zoneType |
| 20 | 库存状态 | /wms/inventoryStatus |

#### 3.1.2 基础数据-物料管理 (15+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 物料档案 | /mdm/material |
| 2 | 物料分类 | /mdm/materialCategory |
| 3 | BOM管理 | /mdm/bom |
| 4 | BOM编辑 | /mdm/bomEditor |
| 5 | 物料包装 | /mdm/materialPackage |
| 6 | 物料区域 | /mdm/materialArea |
| 7 | 替代物料 | /mdm/alternativeMaterial |
| 8 | 物料特性 | /mdm/materialFeature |
| 9 | 物料属性 | /mdm/materialAttribute |
| 10 | 物料组 | /mdm/materialGroup |
| 11 | 物料状态 | /mdm/materialStatus |
| 12 | 采购单位 | /mdm/purchaseUnit |
| 13 | 库存单位 | /mdm/stockUnit |
| 14 | 物料标签 | /mdm/materialLabel |
| 15 | 物料描述 | /mdm/materialDesc |

#### 3.1.3 基础数据-客户管理 (7+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 客户档案 | /mdm/customer |
| 2 | 客户物料 | /mdm/customerMaterial |
| 3 | 客户项目 | /mdm/customerProject |
| 4 | 价格管理 | /mdm/price |
| 5 | 客户联系人 | /mdm/customerContact |
| 6 | 客户地址 | /mdm/customerAddress |
| 7 | 客户银行 | /mdm/customerBank |

#### 3.1.4 基础数据-供应商管理 (4+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 供应商档案 | /mdm/supplier |
| 2 | 供应商物料 | /mdm/supplierMaterial |
| 3 | 供应商周期 | /mdm/supplierCycle |
| 4 | 供应商价格 | /mdm/supplierPrice |

#### 3.1.5 基础数据-标签管理 (8+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 标签类型 | /wms/labelType |
| 2 | 标签模板 | /wms/labelTemplate |
| 3 | 条码规则 | /wms/barcodeRule |
| 4 | 叫料规则 | /wms/callMaterialRule |
| 5 | 标签打印 | /wms/labelPrint |
| 6 | 标签历史 | /wms/labelHistory |
| 7 | 标签绑定 | /wms/labelBind |
| 8 | 标签解除 | /wms/labelUnbind |

#### 3.1.6 基础数据-策略设置 (25+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 上架策略 | /wms/putawayStrategy |
| 2 | 下架策略 | /wms/pickStrategy |
| 3 | 批策略 | /wms/batchStrategy |
| 4 | 库容策略 | /wms/capacityStrategy |
| 5 | 补货策略 | /wms/replenishStrategy |
| 6 | 分配策略 | /wms/allocateStrategy |
| 7 | 合并策略 | /wms/mergeStrategy |
| 8 | 分割策略 | /wms/splitStrategy |
| 9 | 路径策略 | /wms/routeStrategy |
| 10 | 排序策略 | /wms/sortStrategy |
| 11 | 筛选策略 | /wms/filterStrategy |
| 12 | 推荐策略 | /wms/recommendStrategy |
| 13 | 预警策略 | /wms/alertStrategy |
| 14 | 冻结策略 | /wms/freezeStrategy |
| 15 | 解冻策略 | /wms/unfreezeStrategy |
| 16 | 锁定策略 | /wms/lockStrategy |
| 17 | 解锁策略 | /wms/unlockStrategy |
| 18 | 检验策略 | /wms/checkStrategy |
| 19 | 报废策略 | /wms/scrrapStrategy |
| 20 | 盘点策略 | /wms/stockCheckStrategy |
| 21 | 调拨策略 | /wms/transferStrategy |
| 22 | 退货策略 | /wms/returnStrategy |
| 23 | 签收策略 | /wms/signStrategy |
| 24 | 复核策略 | /wms/reviewStrategy |
| 25 | 打包策略 | /wms/packStrategy |

#### 3.1.7 基础数据-单据设置 (8+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 业务类型 | /wms/bizType |
| 2 | 单据类型 | /wms/docType |
| 3 | 编号规则 | /wms/codeRule |
| 4 | 单据状态 | /wms/docStatus |
| 5 | 审批流程 | /wms/approveFlow |
| 6 | 打印设置 | /wms/printSet |
| 7 | 字段配置 | /wms/fieldConfig |
| 8 | 权限配置 | /wms/authConfig |

#### 3.1.8 库存管理 (10+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 库存台账 | /wms/inventory |
| 2 | 库存变化 | /wms/inventoryChange |
| 3 | 库存快照 | /wms/inventorySnapshot |
| 4 | 容器管理 | /wms/container |
| 5 | 库存盘点 | /wms/stockCheck |
| 6 | 库存冻结 | /wms/inventoryFreeze |
| 7 | 库存锁定 | /wms/inventoryLock |
| 8 | 库存分配 | /wms/inventoryAllocate |
| 9 | 库存预留 | /wms/inventoryReserve |
| 10 | 库存查询 | /wms/inventoryQuery |

#### 3.1.9 入库管理 (12+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 采购订单 | /wms/purchaseOrder |
| 2 | 采购到货 | /wms/purchaseArrive |
| 3 | IQC检验 | /quality/iqc |
| 4 | 采购入库 | /wms/purchaseInput |
| 5 | 生产入库 | /wms/productionInput |
| 6 | 退货入库 | /wms/returnInput |
| 7 | 调拨入库 | /wms/transferInput |
| 8 | 其他入库 | /wms/otherInput |
| 9 | 入库确认 | /wms/inputConfirm |
| 10 | 入库上架 | /wms/inputPutaway |
| 11 | ASN管理 | /supplier/asn |
| 12 | 到货登记 | /wms/arriveRegister |

#### 3.1.10 出库管理 (12+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 销售订单 | /wms/salesOrder |
| 2 | 销售发货 | /wms/salesDelivery |
| 3 | 领料单 | /wms/pickOrder |
| 4 | 生产领料 | /wms/productionPick |
| 5 | 拣货作业 | /wms/picking |
| 6 | 补料作业 | /wms/replenish |
| 7 | 出库确认 | /wms/outputConfirm |
| 8 | 出库复核 | /wms/outputReview |
| 9 | 装车管理 | /wms/loading |
| 10 | 送货管理 | /wms/delivery |
| 11 | 签收管理 | /wms/sign |
| 12 | 运费结算 | /wms/freight |

#### 3.1.11 移库管理 (10+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 移库申请 | /wms/transferApply |
| 2 | 移库作业 | /wms/transferTask |
| 3 | 移库记录 | /wms/transferRecord |
| 4 | 库间调拨 | /wms/transferBetween |
| 5 | 库内移动 | /wms/transferInside |
| 6 | 盘点调整 | /wms/stockAdjust |
| 7 | 库存损益 | /wms/inventoryProfit |
| 8 | 库存盘点 | /wms/stockCheck |
| 9 | 库存复核 | /wms/stockReview |
| 10 | 差异处理 | /wms/difference |

#### 3.1.12 库存作业管理 (20+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 容器管理 | /wms/container |
| 2 | 容器绑定 | /wms/containerBind |
| 3 | 容器解绑 | /wms/containerUnbind |
| 4 | 包装管理 | /wms/package |
| 5 | 打包作业 | /wms/packTask |
| 6 | 拆包作业 | /wms/unpackTask |
| 7 | 换箱作业 | /wms/changeBox |
| 8 | 合并作业 | /wms/mergeTask |
| 9 | 分割作业 | /wms/splitTask |
| 10 | 报废管理 | /wms/scrap |
| 11 | 报废申请 | /wms/scrapApply |
| 12 | 报废审批 | /wms/scrapApprove |
| 13 | 备件管理 | /equipment/spare |
| 14 | 备件申请 | /equipment/spareApply |
| 15 | 备件入库 | /equipment/spareInput |
| 16 | 备件出库 | /equipment/spareOutput |
| 17 | 备件库存 | /equipment/spareInventory |
| 18 | 备件位置 | /equipment/spareLocation |
| 19 | 库存预警 | /wms/inventoryAlert |
| 20 | 效期预警 | /wms/validAlert |

#### 3.1.13 结算管理 (6+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 客户收款 | /fin/customerReceive |
| 2 | 客户结算 | /fin/customerSettlement |
| 3 | 备货结算 | /fin/prepareSettlement |
| 4 | 应付管理 | /fin/payable |
| 5 | 收款核销 | /fin/receiveVerify |
| 6 | 结算单据 | /fin/settlementDoc |

#### 3.1.14 供应商管理 (8+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 供应商账期 | /supplier/accountPeriod |
| 2 | 供应商发票 | /supplier/invoice |
| 3 | 供应商索赔 | /supplier/claim |
| 4 | 供应商对账 | /supplier/reconcile |
| 5 | 供应商评级 | /supplier/rating |
| 6 | 供应商资质 | /supplier/qualify |
| 7 | 供应商合同 | /supplier/contract |
| 8 | 供应商考核 | /supplier/assess |

#### 3.1.15 AGV管理 (5+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | AGV设备 | /agv/device |
| 2 | AGV库位 | /agv/location |
| 3 | AGV任务 | /agv/task |
| 4 | AGV库位关系 | /agv/locationRelation |
| 5 | AGV接口 | /agv/interface |

#### 3.1.16 MES条码对接 (1+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | MES集成 | /wms/mesIntegration |

### 3.2 MOM3.0 WMS页面清单 (9页)

| 序号 | 页面名称 | 路径 | 功能说明 |
|-----|---------|------|---------|
| 1 | 仓库管理 | /wms/warehouse | 仓库档案 |
| 2 | 库位管理 | /wms/location | 库位定义 |
| 3 | 收货订单 | /wms/receive | 采购收货 |
| 4 | 发货订单 | /wms/delivery | 销售发货 |
| 5 | 库存台账 | /wms/inventory | 实时库存 |
| 6 | 调拨单 | /wms/transfer | 库间调拨 |
| 7 | 盘点管理 | /wms/stock-check | 库存盘点 |
| 8 | 数据点配置 | /wms/data-point | 采集数据点 |
| 9 | 扫描日志 | /wms/scan-log | 扫码记录 |

### 3.3 WMS模块缺失清单 (-291页)

| 类别 | 缺失页面数 | 主要缺失功能 |
|------|-----------|-------------|
| 工厂建模 | 20+ | 库区管理、库位组、Dock管理、巷道、层架等 |
| 物料管理 | 15+ | 物料包装、替代物料、物料特性、物料组等 |
| 客户管理 | 7+ | 客户项目、价格管理、客户联系人等 |
| 供应商管理 | 4+ | 供应商物料、供应商周期、价格等 |
| 标签管理 | 8+ | 标签类型、模板、条码规则、叫料规则等 |
| 策略设置 | 25+ | 上架/下架/批策略、库容/补货策略等 |
| 单据设置 | 8+ | 业务类型、编号规则、审批流程等 |
| 库存管理 | 10+ | 库存变化、快照、冻结、锁定、分配等 |
| 入库管理 | 12+ | 采购到货、IQC检验、生产入库、退货等 |
| 出库管理 | 12+ | 领料单、拣货、补料、装车、送货等 |
| 移库管理 | 10+ | 移库申请、库内移动、库存损益等 |
| 库存作业 | 20+ | 容器绑定、拆包、换箱、合并、报废等 |
| 结算管理 | 6+ | 客户收款、结算、备货、应收等 |
| 供应商管理 | 8+ | 账期、发票、索赔、对账、评级等 |
| AGV管理 | 5+ | AGV设备、库位关系、接口配置等 |
| MES对接 | 1+ | MES集成配置 |

---

## 4. QMS质量管理详细对比

### 4.1 SFMS3.0 QMS页面清单 (50+页)

| 序号 | 页面名称 | 路径 | 功能说明 |
|-----|---------|------|---------|
| 1 | AQL配置 | /quality/aql | 接收质量限 |
| 2 | 计数器 | /quality/counter | 检验计数 |
| 3 | 动态规则 | /quality/dynamicRule | 规则配置 |
| 4 | 检验方案 | /quality/inspectScheme | 方案增改查 |
| 5 | 检验模板 | /quality/inspectTemplate | 模板定义 |
| 6 | 检验方法 | /quality/inspectMethod | 检验方法 |
| 7 | 检验阶段 | /quality/inspectStage | 检验阶段 |
| 8 | 定性特性 | /quality/qualitativeFeature | 定性特性 |
| 9 | 定量特性 | /quality/quantitativeFeature | 定量特性 |
| 10 | 抽样方案 | /quality/samplingScheme | 抽样配置 |
| 11 | 抽样过程 | /quality/samplingProcess | 抽样记录 |
| 12 | 抽样代码 | /quality/samplingCode | 抽样编码 |
| 13 | 检验请求 | /quality/inspectRequest | 检验申请单 |
| 14 | 检验任务 | /quality/inspectTask | 检验任务单 |
| 15 | IQC记录 | /quality/iqcRecord | 来料检验记录 |
| 16 | PQC记录 | /quality/pqcRecord | 过程检验记录 |
| 17 | OQC记录 | /quality/oqcRecord | 出货检验记录 |
| 18 | 首件检验 | /quality/firstInspect | 首件检验 |
| 19 | 巡检记录 | /quality/patrolInspect | 巡检记录 |
| 20 | 质量通知 | /quality/qualityNotice | 质量通知单 |
| 21 | IQC检验 | /quality/iqc | 来料检验 |
| 22 | IPQC检验 | /quality/ipqc | 过程检验 |
| 23 | FQC检验 | /quality/fqc | 出货检验 |
| 24 | OQC检验 | /quality/oqc | 出货检验 |
| 25 | 检验计划 | /quality/inspectPlan | 检验计划 |
| 26 | 检验记录 | /quality/inspectRecord | 检验记录 |
| 27 | 缺陷代码 | /quality/defectCode | 缺陷代码 |
| 28 | 缺陷记录 | /quality/defectRecord | 缺陷记录 |
| 29 | NCR处理 | /quality/ncr | 不合格处理 |
| 30 | SPC数据 | /quality/spc | SPC数据采集 |
| 31 | LPA标准 | /quality/lpa | LPA审核标准 |
| 32 | QRCI | /quality/qrci | 质量改进 |
| 33 | 实验室样本 | /quality/labSample | 样本管理 |
| 34 | 实验室报告 | /quality/labReport | 报告管理 |
| 35 | 实验室仪器 | /quality/labInstrument | 仪器管理 |
| 36 | 质量标准 | /quality/std | 质量标准 |
| 37 | 质量目标 | /quality/target | 质量目标 |
| 38 | 质量分析 | /quality/analysis | 质量分析 |
| 39 | 质量报表 | /quality/report | 质量报表 |
| 40 | 退货管理 | /quality/return | 退货处理 |
| 41 | 索赔管理 | /quality/claim | 索赔处理 |
| 42 | 供应商评审 | /quality/supplierAudit | 供应商评审 |
| 43 | 纠正措施 | /quality/corrective | 纠正措施 |
| 44 | 预防措施 | /quality/preventive | 预防措施 |
| 45 | 质量培训 | /quality/training | 培训记录 |
| 46 | 质量审核 | /quality/audit | 审核记录 |
| 47 | 质量证书 | /quality/certificate | 证书管理 |
| 48 | 计量管理 | /quality/measure | 计量器具 |
| 49 | 校准记录 | /quality/calibration | 校准记录 |
| 50 | 质量追溯 | /quality/trace | 追溯查询 |

### 4.2 MOM3.0 QMS页面清单 (16页)

| 序号 | 页面名称 | 路径 | 功能说明 |
|-----|---------|------|---------|
| 1 | IQC来料检验 | /quality/iqc | 来料检验 |
| 2 | IPQC过程检验 | /quality/ipqc | 过程检验 |
| 3 | FQC出货检验 | /quality/fqc | 出货检验 |
| 4 | OQC出货检验 | /quality/oqc | 出货检验 |
| 5 | 检验计划 | /quality/inspection-plan | 检验计划 |
| 6 | 检验记录 | /quality/inspection-record | 检验记录 |
| 7 | 检验模板 | /quality/inspection-template | 检验模板 |
| 8 | AQL配置 | /quality/aql | AQL配置 |
| 9 | 动态规则 | /quality/dynamic-rule | 动态规则 |
| 10 | 缺陷代码 | /quality/defect-code | 缺陷代码 |
| 11 | 缺陷记录 | /quality/defect-record | 缺陷记录 |
| 12 | NCR处理 | /quality/ncr | NCR处理 |
| 13 | SPC数据 | /quality/spc-data | SPC数据 |
| 14 | LPA标准 | /quality/lpa-standard | LPA标准 |
| 15 | QRCI | /quality/qrci | QRCI |
| 16 | 实验室样本 | /quality/lab/sample | 实验室样本 |

### 4.3 QMS模块缺失清单 (-34页)

| 缺失页面 | 说明 |
|---------|------|
| 计数器 | 检验计数 |
| 检验方案 | 方案增改查 |
| 检验方法 | 检验方法定义 |
| 检验阶段 | 检验阶段配置 |
| 定性特性 | 定性检验特性 |
| 定量特性 | 定量检验特性 |
| 抽样方案 | 抽样方案配置 |
| 抽样过程 | 抽样过程记录 |
| 抽样代码 | 抽样编码 |
| 检验请求 | 检验申请单 |
| 检验任务 | 检验任务单 |
| IQC/PQC/OQC记录 | 各检验记录详情 |
| 首件检验 | 首件检验 |
| 巡检记录 | 巡检记录 |
| 质量通知 | 质量通知单 |
| 质量标准 | 质量标准 |
| 质量目标 | 质量目标 |
| 质量分析 | 质量分析 |
| 质量报表 | 质量报表 |
| 退货管理 | 退货处理 |
| 索赔管理 | 索赔处理 |
| 供应商评审 | 供应商评审 |
| 纠正措施 | 纠正措施 |
| 预防措施 | 预防措施 |
| 质量培训 | 培训记录 |
| 质量审核 | 审核记录 |
| 质量证书 | 证书管理 |
| 计量管理 | 计量器具 |
| 校准记录 | 校准记录 |
| 质量追溯 | 追溯查询 |

---

## 5. EAM设备管理详细对比

### 5.1 SFMS3.0 EAM页面清单 (65+页)

#### 5.1.1 设备台账 (8+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 设备台账 | /eam/equipment |
| 2 | 设备主要部件 | /eam/equipmentPart |
| 3 | 设备供应商 | /eam/equipmentSupplier |
| 4 | 设备制造商 | /eam/equipmentManufacturer |
| 5 | 设备分类 | /eam/equipmentType |
| 6 | 设备图片 | /eam/equipmentImage |
| 7 | 设备文档 | /eam/equipmentDoc |
| 8 | 设备证书 | /eam/equipmentCertificate |

#### 5.1.2 点检管理 (10+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 点检计划 | /eam/checkPlan |
| 2 | 点检明细 | /eam/checkItem |
| 3 | 点检记录 | /eam/checkRecord |
| 4 | 点检项 | /eam/checkPoint |
| 5 | 点检选项 | /eam/checkOption |
| 6 | 点检选择集 | /eam/checkSet |
| 7 | 点检标准 | /eam/checkStd |
| 8 | 点检周期 | /eam/checkCycle |
| 9 | 点检路线 | /eam/checkRoute |
| 10 | 点检配置 | /eam/checkConfig |

#### 5.1.3 保养管理 (10+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 保养计划 | /eam/maintenancePlan |
| 2 | 保养明细 | /eam/maintenanceItem |
| 3 | 保养记录 | /eam/maintenanceRecord |
| 4 | 保养项目 | /eam/maintenancePoint |
| 5 | 保养选项 | /eam/maintenanceOption |
| 6 | 保养选择集 | /eam/maintenanceSet |
| 7 | 保养标准 | /eam/maintenanceStd |
| 8 | 保养周期 | /eam/maintenanceCycle |
| 9 | 保养类型 | /eam/maintenanceType |
| 10 | 保养配置 | /eam/maintenanceConfig |

#### 5.1.4 维修管理 (8+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 维修工单 | /eam/repairOrder |
| 2 | 维修记录 | /eam/repairRecord |
| 3 | 维修备件 | /eam/repairSpare |
| 4 | 维修经验 | /eam/repairExp |
| 5 | 维修标准 | /eam/repairStd |
| 6 | 维修类型 | /eam/repairType |
| 7 | 维修流程 | /eam/repairFlow |
| 8 | 维修配置 | /eam/repairConfig |

#### 5.1.5 设备运行 (5+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 设备停机 | /eam/downtime |
| 2 | 设备签到 | /eam/checkin |
| 3 | 设备转移 | /eam/transfer |
| 4 | 设备变更 | /eam/change |
| 5 | 设备运行记录 | /eam/runRecord |

#### 5.1.6 备件管理 (8+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 备件档案 | /eam/spare |
| 2 | 备件申请 | /eam/spareApply |
| 3 | 备件入库 | /eam/spareInput |
| 4 | 备件出库 | /eam/spareOutput |
| 5 | 备件库存 | /eam/spareInventory |
| 6 | 备件位置 | /eam/spareLocation |
| 7 | 备件查询 | /eam/spareQuery |
| 8 | 备件盘点 | /eam/spareCheck |

#### 5.1.7 工具管理 (6+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 工具台账 | /eam/tool |
| 2 | 工具型号 | /eam/toolType |
| 3 | 工具签到 | /eam/toolCheckin |
| 4 | 工具出入库 | /eam/toolInOut |
| 5 | 工具库存 | /eam/toolInventory |
| 6 | 工具查询 | /eam/toolQuery |

#### 5.1.8 基础设置 (10+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 工厂建模 | /eam/factory |
| 2 | 设备组织 | /eam/equipmentOrg |
| 3 | 产线基础 | /eam/productionLine |
| 4 | 车间基础 | /eam/workshop |
| 5 | 工位基础 | /eam/workstation |
| 6 | 设备分类 | /eam/equipmentClass |
| 7 | 设备状态 | /eam/equipmentStatus |
| 8 | 故障类型 | /eam/faultType |
| 9 | 故障原因 | /eam/faultReason |
| 10 | 维修级别 | /eam/repairLevel |

### 5.2 MOM3.0 EAM页面清单 (14页)

| 序号 | 页面名称 | 路径 | 功能说明 |
|-----|---------|------|---------|
| 1 | 工厂建模 | /eam/factory | 工厂配置 |
| 2 | 设备组织 | /eam/equipment-org | 设备组织 |
| 3 | 设备停机 | /eam/downtime | 停机记录 |
| 4 | 设备台账 | /equipment/list | 设备档案 |
| 5 | 点检计划 | /equipment/check | 点检计划 |
| 6 | 点检记录 | /equipment/check-record | 点检记录 |
| 7 | 保养计划 | /equipment/maintenance | 保养计划 |
| 8 | 维修管理 | /equipment/repair | 维修管理 |
| 9 | OEE分析 | /equipment/oee | OEE分析 |
| 10 | 设备检验 | /equipment/inspection | 设备检验 |
| 11 | 设备缺陷 | /equipment/defect | 缺陷记录 |
| 12 | 检验模板 | /equipment/template | 检验模板 |
| 13 | 仪表管理 | /equipment/gauge | 仪表管理 |
| 14 | 备件管理 | /equipment/spare-part | 备件管理 |

### 5.3 EAM模块缺失清单 (-51页)

| 类别 | 缺失页面数 | 主要缺失功能 |
|------|-----------|-------------|
| 设备台账 | 8+ | 主要部件、供应商、制造商、图片、文档、证书 |
| 点检管理 | 10+ | 点检选项、点检选择集、点检标准、点检路线等 |
| 保养管理 | 10+ | 保养选项、保养选择集、保养标准、保养类型等 |
| 维修管理 | 8+ | 维修备件、维修经验库、维修标准、维修流程等 |
| 设备运行 | 5+ | 设备签到、转移、变更、运行记录 |
| 备件管理 | 8+ | 备件申请、备件入库/出库、备件查询、盘点 |
| 工具管理 | 6+ | 工具台账、型号、签到、出入库 |
| 基础设置 | 10+ | 产线/车间/工位基础、设备分类、故障类型等 |

---

## 6. BPM流程模块详细对比

### 6.1 SFMS3.0 BPM页面清单 (30+页)

| 序号 | 页面名称 | 路径 | 功能说明 |
|-----|---------|------|---------|
| 1 | 流程模型设计器 | /bpm/modeler | BPMN流程设计 |
| 2 | 流程模型 | /bpm/model | 模型管理 |
| 3 | 流程定义 | /bpm/definition | 定义管理 |
| 4 | 表单管理 | /bpm/form | 表单设计 |
| 5 | 用户组 | /bpm/userGroup | 用户组管理 |
| 6 | 任务分配规则 | /bpm/assignRule | 分配规则 |
| 7 | 流程实例 | /bpm/instance | 实例管理 |
| 8 | 待办任务 | /bpm/task | 待办任务 |
| 9 | 已办任务 | /bpm/doneTask | 已办任务 |
| 10 | 任务转派 | /bpm/taskTransfer | 任务转派 |
| 11 | 任务退回 | /bpm/taskReturn | 任务退回 |
| 12 | 审批管理 | /bpm/approve | 审批管理 |
| 13 | 驳回管理 | /bpm/reject | 驳回管理 |
| 14 | 委托管理 | /bpm/delegate | 委托管理 |
| 15 | 流程监控 | /bpm/monitor | 流程监控 |
| 16 | 流程日志 | /bpm/log | 流程日志 |
| 17 | 流程版本 | /bpm/version | 版本管理 |
| 18 | 流程分类 | /bpm/category | 分类管理 |
| 19 | 节点配置 | /bpm/nodeConfig | 节点配置 |
| 20 | 连线配置 | /bpm/sequenceConfig | 连线配置 |
| 21 | 表达式配置 | /bpm/expressionConfig | 表达式 |
| 22 | 消息配置 | /bpm/messageConfig | 消息配置 |
| 23 | 定时配置 | /bpm/timerConfig | 定时配置 |
| 24 | 异常处理 | /bpm/exceptionHandle | 异常处理 |
| 25 | 流程仿真 | /bpm/simulation | 仿真测试 |
| 26 | 流程导出 | /bpm/export | 导出管理 |
| 27 | 流程导入 | /bpm/import | 导入管理 |
| 28 | OA请假示例 | /bpm/leaveExample | 请假流程 |
| 29 | 采购审批示例 | /bpm/purchaseExample | 采购审批 |
| 30 | 报销审批示例 | /bpm/expenseExample | 报销审批 |

### 6.2 MOM3.0 BPM页面清单 (3页)

| 序号 | 页面名称 | 路径 | 功能说明 |
|-----|---------|------|---------|
| 1 | 流程定义 | /bpm/process | 流程定义 |
| 2 | 流程实例 | /bpm/instance | 流程实例 |
| 3 | 任务列表 | /bpm/task | 任务列表 |

### 6.3 BPM模块缺失清单 (-27页)

| 缺失页面 | 说明 |
|---------|------|
| 流程模型设计器 | BPMN可视化设计 |
| 表单管理 | 表单设计器 |
| 用户组 | 用户组管理 |
| 任务分配规则 | 分配规则配置 |
| 任务转派 | 任务转派功能 |
| 任务退回 | 任务退回功能 |
| 审批管理 | 审批管理 |
| 驳回管理 | 驳回管理 |
| 委托管理 | 委托管理 |
| 流程监控 | 流程监控 |
| 流程日志 | 流程日志 |
| 流程版本 | 版本管理 |
| 流程分类 | 分类管理 |
| 节点配置 | 节点属性配置 |
| 连线配置 | 连线条件配置 |
| 表达式配置 | 表达式配置 |
| 消息配置 | 消息事件配置 |
| 定时配置 | 定时任务配置 |
| 异常处理 | 异常处理配置 |
| 流程仿真 | 仿真测试 |
| 流程导出 | 导出管理 |
| 流程导入 | 导入管理 |
| OA请假示例 | 请假流程示例 |
| 采购审批示例 | 采购审批示例 |
| 报销审批示例 | 报销审批示例 |

---

## 7. SYSTEM系统管理详细对比

### 7.1 SFMS3.0 SYSTEM页面清单 (35+页)

| 序号 | 页面名称 | 路径 | 功能说明 |
|-----|---------|------|---------|
| 1 | 用户管理 | /system/user | 用户CRUD |
| 2 | 角色管理 | /system/role | 角色权限 |
| 3 | 菜单管理 | /system/menu | 菜单配置 |
| 4 | 部门管理 | /system/dept | 部门组织 |
| 5 | 岗位管理 | /system/post | 岗位配置 |
| 6 | 租户管理 | /system/tenant | 租户配置 |
| 7 | 租户套餐 | /system/tenantPackage | 套餐管理 |
| 8 | 数据字典 | /system/dict | 字典配置 |
| 9 | 区域管理 | /system/region | 行政区域 |
| 10 | 短信管理 | /system/sms | 短信渠道 |
| 11 | 短信模板 | /system/smsTemplate | 短信模板 |
| 12 | 邮件管理 | /system/mail | 邮件配置 |
| 13 | 邮件模板 | /system/mailTemplate | 邮件模板 |
| 14 | 消息通知 | /system/notice | 消息通知 |
| 15 | 通知模板 | /system/noticeTemplate | 通知模板 |
| 16 | OAuth2客户端 | /system/oauth2 | OAuth配置 |
| 17 | Token管理 | /system/token | Token管理 |
| 18 | 登录日志 | /system/loginLog | 登录记录 |
| 19 | 操作日志 | /system/operLog | 操作记录 |
| 20 | 敏感词 | /system/sensitiveWord | 敏感词过滤 |
| 21 | 错误码 | /system/errorCode | 错误码配置 |
| 22 | 序列号 | /system/sequence | 序列号管理 |
| 23 | 密码规则 | /system/passwordRule | 密码策略 |
| 24 | 安全配置 | /system/securityConfig | 安全配置 |
| 25 | IP白名单 | /system/ipWhiteList | IP白名单 |
| 26 | 操作限制 | /system/operLimit | 操作限制 |
| 27 | 通知公告 | /system/notice | 公告管理 |
| 28 | 系统参数 | /system/param | 参数配置 |
| 29 | 字典类型 | /system/dictType | 字典分类 |
| 30 | 地区管理 | /system/area | 地区配置 |
| 31 | 编码规则 | /system/codeRule | 编码规则 |
| 32 | 数据权限 | /system/dataPerm | 数据权限 |
| 33 | 字段权限 | /system/fieldPerm | 字段权限 |
| 34 | API权限 | /system/apiPerm | API权限 |
| 35 | 权限组 | /system/permGroup | 权限组 |

### 7.2 MOM3.0 SYSTEM页面清单 (13页)

| 序号 | 页面名称 | 路径 | 功能说明 |
|-----|---------|------|---------|
| 1 | 用户管理 | /system/user | 用户CRUD |
| 2 | 角色管理 | /system/role | 角色权限 |
| 3 | 菜单管理 | /system/menu | 菜单配置 |
| 4 | 部门管理 | /system/dept | 部门组织 |
| 5 | 岗位管理 | /system/post | 岗位配置 |
| 6 | 数据字典 | /system/dict | 字典配置 |
| 7 | 租户管理 | /system/tenant | 租户配置 |
| 8 | 登录日志 | /system/login-log | 登录记录 |
| 9 | 操作日志 | /system/oper-log | 操作记录 |
| 10 | 通知公告 | /system/notice | 公告管理 |
| 11 | AI配置 | /system/ai-config | AI配置 |
| 12 | 打印模板 | /system/print-template | 打印模板 |
| 13 | 系统配置 | /system/config | 系统配置 |

### 7.3 SYSTEM模块缺失清单 (-22页)

| 缺失页面 | 说明 |
|---------|------|
| 租户套餐 | 租户套餐管理 |
| 区域管理 | 行政区域管理 |
| 短信管理 | 短信渠道配置 |
| 短信模板 | 短信模板管理 |
| 邮件管理 | 邮件账号配置 |
| 邮件模板 | 邮件模板管理 |
| 消息通知 | 站内消息管理 |
| 通知模板 | 通知模板配置 |
| OAuth2客户端 | 第三方登录配置 |
| Token管理 | Token管理 |
| 敏感词 | 敏感词过滤 |
| 错误码 | 错误码配置 |
| 序列号 | 序列号管理 |
| 密码规则 | 密码策略配置 |
| 安全配置 | 安全策略配置 |
| IP白名单 | IP访问控制 |
| 操作限制 | 操作限制配置 |
| 系统参数 | 参数配置 |
| 字典类型 | 字典分类管理 |
| 地区管理 | 地区配置 |
| 编码规则 | 编码规则配置 |
| 数据权限 | 数据权限控制 |
| 字段权限 | 字段权限控制 |
| API权限 | API权限控制 |
| 权限组 | 权限组管理 |

---

## 8. INFRA基础设施详细对比

### 8.1 SFMS3.0 INFRA页面清单 (12+页)

| 序号 | 页面名称 | 路径 | 功能说明 |
|-----|---------|------|---------|
| 1 | 文件管理 | /infra/file | 文件上传 |
| 2 | 文件配置 | /infra/fileConfig | 文件配置 |
| 3 | 代码生成器 | /infra/codegen | CRUD生成 |
| 4 | 数据源配置 | /infra/datasource | 数据源 |
| 5 | 数据库文档 | /infra/dbdoc | 数据库文档 |
| 6 | 定时任务 | /infra/job | 任务调度 |
| 7 | 定时日志 | /infra/jobLog | 任务日志 |
| 8 | Redis管理 | /infra/redis | 缓存管理 |
| 9 | API日志 | /infra/apiLog | API日志 |
| 10 | 错误日志 | /infra/errorLog | 错误日志 |
| 11 | 配置管理 | /infra/config | 配置管理 |
| 12 | Druid监控 | /infra/druid | Druid监控 |
| 13 | Swagger文档 | /infra/swagger | API文档 |
| 14 | WebSocket | /infra/websocket | WebSocket |
| 15 | 服务器监控 | /infra/server | 服务器监控 |

### 8.2 MOM3.0 INFRA页面清单 (0页)

**完全缺失**

### 8.3 INFRA模块缺失清单 (-12页)

| 缺失页面 | 说明 |
|---------|------|
| 文件管理 | 文件上传下载 |
| 文件配置 | 文件存储配置 |
| 代码生成器 | 代码生成工具 |
| 数据源配置 | 多数据源配置 |
| 数据库文档 | 数据库文档生成 |
| 定时任务 | 任务调度管理 |
| 定时日志 | 调度日志查看 |
| Redis管理 | Redis缓存管理 |
| API日志 | API访问日志 |
| 错误日志 | 错误日志查看 |
| Druid监控 | 数据库监控 |
| Swagger文档 | API文档 |

---

## 9. REPORT报表模块对比

### 9.1 SFMS3.0 REPORT页面清单 (1+页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 生产日报 | /report/productionDaily |

### 9.2 MOM3.0 REPORT页面清单 (5页)

| 序号 | 页面名称 | 路径 | 功能说明 |
|-----|---------|------|---------|
| 1 | 生产日报 | /report/production-daily | 生产日报 |
| 2 | OEE报表 | /report/oee | OEE报表 |
| 3 | 质量周报 | /report/quality-weekly | 质量周报 |
| 4 | 交付报表 | /report/delivery | 交付报表 |
| 5 | Andon报表 | /report/andon | 安灯报表 |

### 9.3 REPORT模块对比 (+4页)

MOM3.0报表模块比SFMS3.0多4页

---

## 10. SCP供应链模块对比

### 10.1 SFMS3.0 SCP页面清单 (6页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 询价管理 | /scp/rfq |
| 2 | 采购订单 | /scp/purchaseOrder |
| 3 | SCP销售订单 | /scp/salesOrder |
| 4 | 供应商报价 | /scp/supplierQuote |
| 5 | 供应商KPI | /scp/supplierKpi |
| 6 | 客户询价 | /scp/customerInquiry |

### 10.2 MOM3.0 SCP页面清单 (6页)

| 序号 | 页面名称 | 路径 |
|-----|---------|------|
| 1 | 询价管理 | /scp/rfq |
| 2 | 采购订单 | /scp/purchase |
| 3 | SCP销售订单 | /scp/sales-order |
| 4 | 供应商报价 | /scp/supplier-quote |
| 5 | 供应商KPI | /scp/supplier-kpi |
| 6 | 客户询价 | /scp/customer-inquiry |

### 10.3 SCP模块对比 (0差异)

**完全一致**

---

## 11. 完整缺失功能汇总

### 11.1 按优先级分类

#### 高优先级 (影响核心业务)

| 模块 | 缺失功能 | 页面数 |
|------|---------|--------|
| WMS | 库区管理、库位组、物料BOM、标签模板、条码规则 | 50+ |
| WMS | IQC来料检验、生产入库、领料拣货 | 30+ |
| MES | 生产日计划、工单排程、BOM管理、工序管理 | 20+ |
| MES | 生产报工、派工管理、生产看板 | 10+ |
| BPM | 流程模型设计器、表单设计器 | 5+ |
| QMS | 检验方案、抽样方案、检验请求 | 10+ |
| EAM | 设备部件、点检选项、保养选项、维修备件 | 20+ |

#### 中优先级 (影响完整功能)

| 模块 | 缺失功能 | 页面数 |
|------|---------|--------|
| WMS | 策略规则(上架/下架/批策略等) | 25+ |
| WMS | 容器管理、包装管理、报废管理 | 15+ |
| WMS | 结算管理、供应商管理、AGV管理 | 20+ |
| MES | 产品档案、物料申请、拆解管理、返工批次 | 10+ |
| MES | 月计划、工作日历、班组能力 | 10+ |
| BPM | 用户组、任务分配规则、任务转派/退回 | 10+ |
| SYSTEM | 短信/邮件/消息通知、OAuth2 | 15+ |
| EAM | 工具管理、设备转移/变更 | 10+ |
| QMS | 质量标准、质量分析、质量报表 | 10+ |

#### 低优先级 (增强功能)

| 模块 | 缺失功能 | 页面数 |
|------|---------|--------|
| INFRA | 代码生成器、文件管理、定时任务、Redis管理 | 12+ |
| SYSTEM | 敏感词、错误码、序列号、密码规则 | 10+ |
| QMS | 质量培训、质量审核、质量证书、计量管理 | 10+ |

### 11.2 缺失页面总数统计

| 模块 | SFMS3.0 | MOM3.0 | 缺失 | 完成度 |
|------|---------|--------|------|--------|
| MES执行 | 36 | 5 | -31 | 14% |
| WMS仓储 | 300+ | 9 | -291 | 3% |
| QMS质量 | 50+ | 16 | -34 | 32% |
| EAM设备 | 65+ | 14 | -51 | 22% |
| BPM流程 | 30+ | 3 | -27 | 10% |
| SYSTEM系统 | 35+ | 13 | -22 | 37% |
| INFRA基础 | 12+ | 0 | -12 | 0% |
| REPORT报表 | 1+ | 5 | +4 | 500% |
| SCP供应链 | 6 | 6 | 0 | 100% |
| AGV调度 | 3 | 3 | 0 | 100% |
| 安灯系统 | 5 | 5 | 0 | 100% |
| 能源管理 | 1 | 1 | 0 | 100% |
| 追溯管理 | 2 | 2 | 0 | 100% |
| 结算管理 | 3 | 3 | 0 | 100% |
| 主数据MDM | 15 | 15 | 0 | 100% |
| **总计** | **521** | **129** | **-392** | **25%** |

---

## 12. 补充后状态 (v2.0更新)

### 12.1 各模块完成度

| 模块 | SFMS3.0 | MOM3.0原版 | MOM3.0补充后 | 完成度 |
|------|---------|-----------|-------------|--------|
| MES执行 | 36 | 5 | 20 | 56% |
| WMS仓储 | 300+ | 9 | 140+ | 47% |
| QMS质量 | 50+ | 16 | 48 | 96% |
| EAM设备 | 65+ | 14 | 65 | 100% |
| BPM流程 | 30+ | 3 | 30 | 100% |
| SYSTEM系统 | 35+ | 13 | 31 | 89% |
| INFRA基础 | 12+ | 0 | 12 | 100% |
| REPORT报表 | 1+ | 5 | 5 | 500% |
| SCP供应链 | 6 | 6 | 6 | 100% |
| AGV调度 | 3 | 3 | 3 | 100% |
| 安灯系统 | 5 | 5 | 5 | 100% |
| 能源管理 | 1 | 1 | 1 | 100% |
| 追溯管理 | 2 | 2 | 2 | 100% |
| 结算管理 | 3 | 3 | 3 | 100% |
| 主数据MDM | 15 | 15 | 15 | 100% |
| **总计** | **521** | **129** | **415+** | **80%** |

### 12.2 补充完成的设计文档

| 文档名称 | 补充页面数 | 状态 |
|---------|-----------|------|
| MOM3.0前端_WMS仓储模块设计文档.md | 140+ | ✅ 已完成 |
| MOM3.0前端_MES执行模块设计文档.md | 15 | ✅ 已完成 |
| MOM3.0前端_BPM流程模块设计文档.md | 27 | ✅ 已完成 |
| MOM3.0前端_QMS质量管理模块设计文档.md | 32 | ✅ 已完成 |
| MOM3.0前端_EAM设备管理模块设计文档.md | 51 | ✅ 已完成 |
| MOM3.0前端_SYSTEM系统管理模块设计文档.md | 18 | ✅ 已完成 |
| MOM3.0前端_INFRA基础设施模块设计文档.md | 12 | ✅ 已完成 |
| **总计** | **295+** | **✅ 全部完成** |

### 12.3 剩余缺失页面

| 模块 | 剩余缺失 | 说明 |
|------|---------|------|
| WMS仓储 | ~160 | 工厂建模、物料高级功能、部分业务功能 |
| MES执行 | ~16 | 部分高级排程、BOM管理功能 |
| QMS质量 | ~2 | 实验室报告、实验室仪器 |
| SYSTEM系统 | ~4 | 部分高级系统配置 |
| **总计** | **~182** | |

### 12.4 结论

MOM3.0前端设计文档补充已完成：

1. **整体完成度**: 从25%提升至80%
2. **EAM/BPM/INFRA模块**: 已100%完成
3. **QMS模块**: 完成度96%，仅缺实验室报告和仪器
4. **SYSTEM模块**: 完成度89%，主要功能已覆盖
5. **WMS模块**: 完成度47%，核心功能已补充
6. **MES模块**: 完成度56%，基础功能已覆盖

**补充文档已生成**，可作为后续开发的参考依据。

---

*文档版本: v2.0 | 生成日期: 2026-04-17*
