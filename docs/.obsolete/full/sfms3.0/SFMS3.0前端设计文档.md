# SFMS3.0 前端设计文档

> **版本**: v1.0
> **日期**: 2026-04-17
> **项目**: 闻聪制造执行系统 SFMS3.0 前端

---

## 1. 前端技术架构

### 1.1 技术栈

| 技术 | 版本 | 说明 |
|-----|-----|------|
| Vue | 3.3.4 | 渐进式前端框架 |
| Vite | 4.4.9 |下一代前端构建工具 |
| TypeScript | 5.2.2 | JavaScript超集，强类型 |
| Element Plus | 2.3.14 | Vue3 UI组件库 |
| Pinia | 2.1.6 | Vue3状态管理 |
| Vue Router | 4.2.5 | Vue3路由管理 |
| Axios | 1.5.0 | HTTP客户端 |
| GoView | - | 数据可视化大屏 |
| @antv/x6 | 2.18.1 | Graph图编辑引擎 |
| bpmn-js | 8.9.0 | BPMN流程编辑器 |
| form-create | 3.1.3 | 表单设计器 |
| ECharts | 5.4.3 | 图表库 |

### 1.2 项目结构

```
sfms3.0-ui/
├── src/
│   ├── api/                    # API接口层
│   │   ├── bpm/               # BPM流程管理
│   │   ├── eam/               # EAM设备管理
│   │   ├── infra/             # INFRA基础设施
│   │   ├── login/             # 登录认证
│   │   ├── mes/               # MES制造执行
│   │   ├── qms/               # QMS质量管理
│   │   ├── redis/             # Redis缓存
│   │   ├── scp/               # SCP供应链
│   │   ├── system/            # SYSTEM系统管理
│   │   └── wms/               # WMS仓库管理
│   ├── views/                  # 页面视图层
│   │   ├── bpm/               # 流程管理页面
│   │   ├── eam/               # 设备管理页面
│   │   ├── home/              # 首页
│   │   ├── infra/             # 基础设施页面
│   │   ├── login/             # 登录页面
│   │   ├── mes/               # MES页面
│   │   ├── profile/           # 个人中心
│   │   ├── qms/               # QMS页面
│   │   ├── report/            # 报表页面
│   │   ├── system/            # 系统管理页面
│   │   └── wms/               # WMS页面
│   ├── router/                # 路由配置
│   ├── store/                 # 状态管理
│   ├── components/            # 公共组件
│   ├── config/                # 配置文件
│   ├── utils/                 # 工具函数
│   └── assets/                # 静态资源
├── package.json
└── vite.config.ts
```

### 1.3 状态管理 (Pinia Store)

| Store模块 | 职责 |
|----------|------|
| user | 用户信息、Token、权限 |
| permission | 菜单权限、路由权限 |
| dict | 数据字典缓存 |
| tagsView | 标签页管理 |
| app | 应用配置 |
| locale | 国际化 |

---

## 2. 前端页面清单

### 2.1 首页 (Home)

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 首页 | `/home` | 系统首页Dashboard |
| 首页V2 | `/home/index2` | 系统首页V2版本 |
| 首页副本 | `/home/indexCopy` | 首页功能副本 |
| 物料组件 | `/home/components/material` | 首页物料数据展示组件 |
| 生产组件 | `/home/components/produce` | 首页生产数据展示组件 |
| 产品组件 | `/home/components/product` | 首页产品数据展示组件 |
| 供应商组件 | `/home/components/supplierIndex` | 首页供应商数据展示组件 |

### 2.2 登录模块 (Login)

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 登录页 | `/login` | 用户登录、OAuth2登录、SSO登录 |
| 忘记密码 | `/login/forgetPassword` | 忘记密码页面 |
| 更新密码 | `/login/updatePassword` | 更新密码页面 |
| 新密码提示 | `/login/updatePasswordNewTips` | 新密码设置提示页面 |
| 登录表单组件 | `/login/components/LoginForm` | 用户名密码登录表单组件 |
| 手机登录表单 | `/login/components/MobileForm` | 手机验证码登录表单组件 |
| 二维码PDA | `/login/components/QRCodePDA` | PDA设备二维码登录组件 |
| SSO登录组件 | `/login/components/SSOLogin` | SSO单点登录组件 |
| 注册表单组件 | `/login/components/RegisterForm` | 用户注册表单组件 |
| 二维码表单 | `/login/components/QrCodeForm` | 二维码扫码表单组件 |
| 登录标题组件 | `/login/components/LoginFormTitle` | 登录页标题组件 |

### 2.3 MES制造执行模块

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 生产日计划 | `/mes/orderDay` | 日计划创建、发布、终止、恢复 |
| 生产排程 | `/mes/workScheduling` | 工单生成、工序调度、报工 |
| 工艺路线 | `/mes/processroute` | 工艺定义、工序管理 |
| BOM管理 | `/mes/bom` | 物料清单配置 |
| 工位管理 | `/mes/workstation` | 生产工位配置 |
| 班组设置 | `/mes/teamSetting` | 班组人员配置 |
| 假期管理 | `/mes/holiday` | 假期日历配置 |
| 能力信息 | `/mes/abilityInfo` | 人员技能档案 |
| 工序类型 | `/mes/operstepsType` | 工序类型定义 |
| 工序管理 | `/mes/opersteps` | 工序详细配置 |
| 产品档案 | `/mes/item` | 产品BOM物料档案 |
| 物料申请 | `/mes/itemRequestMain` | 生产物料申请 |
| 拆解管理 | `/mes/dismantlingMain` | 产品拆解任务 |
| 返工批次 | `/mes/reworkBatch` | 返工批次管理 |
| 返工单 | `/mes/reworkSingle` | 返工单管理 |
| 月计划 | `/mes/ordermonthplan` | 月度生产计划 |
| 花字品种 | `/mes/pattern` | 花字品种管理 |
| 花字类型 | `/mes/patternType` | 花字类型管理 |
| 工作日历 | `/mes/workcalendar` | 工作日历配置 |
| 质量表单 | `/mes/qualityform` | 质量检验表单 |
| 质量日志 | `/mes/qualityformlog` | 表单操作日志 |
| 质量分组 | `/mes/qualitygroup` | 质量分组配置 |
| 质量等级 | `/mes/qualityclass` | 质量等级定义 |
| 产品报交 | `/mes/reportpStore` | 产品报交记录 |
| 产品下线 | `/mes/productOffline` | 产品下线管理 |
| 产品返线 | `/mes/productBackline` | 产品返线管理 |
| 排程详情 | `/mes/workSchedulingQaform` | 排程明细表单 |
| 工艺管理 | `/mes/process` | 工艺详细配置 |
| 生产计划 | `/mes/productionPlan` | 生产计划管理 |
| 人员能力 | `/mes/hrPersonAbility` | 人员能力矩阵 |

### 2.4 WMS仓库管理模块

#### 2.4.1 基础数据管理

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 仓库管理 | `/wms/basicDataManage/factoryModeling/warehouse` | 仓库档案 |
| 库区管理 | `/wms/basicDataManage/factoryModeling/areabasic` | 库区配置 |
| 库位管理 | `/wms/basicDataManage/factoryModeling/location` | 库位定义 |
| 库位组 | `/wms/basicDataManage/factoryModeling/locationgroup` | 库位分组 |
| 产线管理 | `/wms/basicDataManage/factoryModeling/productionline` | 产线配置 |
| 车间管理 | `/wms/basicDataManage/factoryModeling/workshop` | 车间配置 |
| 工位管理 | `/wms/basicDataManage/factoryModeling/workstation` | 工位配置 |
| 工序管理 | `/wms/basicDataManage/factoryModeling/process` | 工序配置 |
| 区域基础 | `/wms/basicDataManage/factoryModeling/areabasic` | 区域基础 |
|  dock管理 | `/wms/basicDataManage/factoryModeling/dock` | dock管理 |
| 企业管理 | `/wms/basicDataManage/factoryModeling/enterprise` | 企业档案 |
| 物料管理 | `/wms/basicDataManage/itemManage/itembasic` | 物料档案 |
| 物料BOM | `/wms/basicDataManage/itemManage/bom` | 物料BOM配置 |
| 物料仓库 | `/wms/basicDataManage/itemManage/itemwarehouse` | 物料仓库关系 |
| 物料包装 | `/wms/basicDataManage/itemManage/itempackage` | 物料包装配置 |
| 物料区域 | `/wms/basicDataManage/itemManage/itemarea` | 物料区域关系 |
| 产线物料 | `/wms/basicDataManage/itemManage/productionlineitem` | 产线物料配置 |
| 替代物料 | `/wms/basicDataManage/itemManage/productionitemcodeSpareitemcode` | 替代物料关系 |
| 包装单位 | `/wms/basicDataManage/itemManage/packageunit` | 包装单位配置 |
| 退货申请 | `/wms/basicDataManage/itemManage/relegate/relegateRequest` | 物料退货申请 |
| 退货记录 | `/wms/basicDataManage/itemManage/relegate/relegateRecord` | 物料退货记录 |
| 标准成本 | `/wms/basicDataManage/itemManage/stdcostprice` | 标准成本配置 |
| 客户管理 | `/wms/basicDataManage/customerManage/customer` | 客户档案 |
| 客户物料 | `/wms/basicDataManage/customerManage/customeritem` | 客户物料关系 |
| 客户dock | `/wms/basicDataManage/customerManage/customerdock` | 客户dock配置 |
| 客户项目 | `/wms/basicDataManage/customerManage/project` | 客户项目 |
| 销售价格 | `/wms/basicDataManage/customerManage/saleprice` | 销售价格 |
| 交货预测 | `/wms/basicDataManage/customerManage/customerDeliveryForecast` | 交货预测 |
| 供应商管理 | `/wms/basicDataManage/supplierManage/supplier` | 供应商档案 |
| 供应商物料 | `/wms/basicDataManage/supplierManage/supplieritem` | 供应商物料关系 |
| 供应商周期 | `/wms/basicDataManage/supplierManage/supplierCycle` | 供应商交货周期 |
| 采购价格 | `/wms/basicDataManage/supplierManage/purchaseprice` | 采购价格 |
| 承运商管理 | `/wms/basicDataManage/orderManage/carrier` | 承运商档案 |
| 货主管理 | `/wms/basicDataManage/orderManage/owner` | 货主档案 |
| 班组管理 | `/wms/basicDataManage/orderManage/team` | 班组配置 |
| 班组表单 | `/wms/basicDataManage/orderManage/team/teamForm` | 班组表单 |
| 班次管理 | `/wms/basicDataManage/orderManage/shift` | 班次配置 |
| 标签类型 | `/wms/basicDataManage/labelManage/labeltype` | 标签类型 |
| 标签模板 | `/wms/basicDataManage/labelManage/barbasic` | 标签模板配置 |
| 条码规则 | `/wms/basicDataManage/labelManage/barcode` | 条码规则 |
| 叫料管理 | `/wms/basicDataManage/labelManage/callmaterials` | 叫料管理 |
| 位置标签 | `/wms/basicDataManage/labelManage/locationLabel` | 库位标签 |
| 生产标签 | `/wms/basicDataManage/labelManage/manufacturePackage` | 生产包装标签 |
| 采购标签 | `/wms/basicDataManage/labelManage/purchasePackage` | 采购包装标签 |
| 器具标签 | `/wms/basicDataManage/labelManage/utensilPackage` | 器具标签 |
| 业务类型 | `/wms/basicDataManage/documentSetting/businesstype` | 业务类型配置 |
| 单据设置 | `/wms/basicDataManage/documentSetting/documentsetting` | 单据编号规则 |
| 单据类型 | `/wms/basicDataManage/documentSetting/transactiontype` | 交易类型配置 |
| 计划设置 | `/wms/basicDataManage/documentSetting/plansetting` | 计划设置 |
| 作业设置 | `/wms/basicDataManage/documentSetting/jobsetting` | 作业设置 |
| 记录设置 | `/wms/basicDataManage/documentSetting/recordsetting` | 记录设置 |
| 请求设置 | `/wms/basicDataManage/documentSetting/requestsetting` | 请求设置 |
| 单据开关 | `/wms/basicDataManage/documentSetting/switch` | 单据开关配置 |
| 策略规则 | `/wms/basicDataManage/strategySetting/rule` | 策略规则配置 |
| 条件配置 | `/wms/basicDataManage/strategySetting/condition` | 策略条件配置 |
| 参数设置 | `/wms/basicDataManage/strategySetting/paramsetting` | 策略参数设置 |
| 配置设置 | `/wms/basicDataManage/strategySetting/configuration` | 配置管理 |
| 配置项 | `/wms/basicDataManage/strategySetting/configurationsetting` | 配置项管理 |
| 备件库位 | `/wms/basicDataManage/strategySetting/spareitemLocation` | 备件库位配置 |
| 上架策略 | `/wms/basicDataManage/strategySetting/strategy/upShelfStrategy` | 上架策略 |
| 下架策略 | `/wms/basicDataManage/strategySetting/strategy/downShelfStrategy` | 下架策略 |
| 批策略 | `/wms/basicDataManage/strategySetting/strategy/batchStrategy` | 批策略 |
| 库容策略 | `/wms/basicDataManage/strategySetting/strategy/storageCapacityStrategy` | 库容策略 |
| 器具容量 | `/wms/basicDataManage/strategySetting/strategy/utensilCapacityStrategy` | 器具容量策略 |
| 供应商交货 | `/wms/basicDataManage/strategySetting/strategy/supplieDeliveryStrategy` | 供应商交货策略 |
| 到货检验策略 | `/wms/basicDataManage/strategySetting/strategy/arrivalInspectionStrategy` | 到货检验策略 |
| 检验策略 | `/wms/basicDataManage/strategySetting/strategy/inspectStrategy` | 检验策略 |
| 采购收货策略 | `/wms/basicDataManage/strategySetting/strategy/purchaseReceiptStrategy` | 采购收货策略 |
| 仓库存储策略 | `/wms/basicDataManage/strategySetting/strategy/warehouseStorageStrategy` | 仓库存储策略 |
| 修管理策略 | `/wms/basicDataManage/strategySetting/strategy/repairMaterialStrategy` | 修管理策略 |
| 库存精度策略 | `/wms/basicDataManage/strategySetting/strategy/manageAccuracyStrategy` | 库存精度策略 |
| 科目管理 | `/wms/basicDataManage/subject/mstr` | 科目主数据 |
| 成本中心 | `/wms/basicDataManage/subject/qadCostcentre` | 成本中心 |
| 项目管理 | `/wms/basicDataManage/subject/qadProject` | 项目主数据 |
| 科目账户 | `/wms/basicDataManage/subject/subjectAccount` | 科目账户 |
| 账户日历 | `/wms/basicDataManage/systemSetting/accountcalendar` | 账户日历 |
| 汇率管理 | `/wms/basicDataManage/systemSetting/currencyexchange` | 货币汇率 |
| 供应商用户 | `/wms/basicDataManage/systemSetting/supplierUser` | 供应商用户 |
| 系统日历 | `/wms/basicDataManage/systemSetting/systemcalendar` | 系统日历 |

#### 2.4.2 库存管理

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 库存台账 | `/wms/inventoryManage/balance` | 实时库存台账 |
| 库存变化历史 | `/wms/inventoryManage/balanceChangeHistory` | 库存异动历史 |
| 库存容器 | `/wms/inventoryManage/balanceContainer` | 容器库存查询 |
| 库存汇总 | `/wms/inventoryManage/balanceSummary` | 库存汇总报表 |
| 容器管理 | `/wms/inventoryManage/container` | 容器档案 |
| 容器绑定记录 | `/wms/inventoryManage/containerinit/containerBindRecord` | 容器绑定记录 |
| 容器维修记录 | `/wms/inventoryManage/containerinit/containerRepair` | 容器维修记录 |
| 容器解绑记录 | `/wms/inventoryManage/containerinit/containerUnbindRecord` | 容器解绑记录 |
| 容器初始化添加 | `/wms/inventoryManage/containerinit/containerinitadd` | 容器初始化添加 |
| 容器初始化记录 | `/wms/inventoryManage/containerinit/containerinitrecord` | 容器初始化记录 |
| 预期入库 | `/wms/inventoryManage/expectin` | 预期入库查询 |
| 预期出库 | `/wms/inventoryManage/expectout` | 预期出库查询 |
| 库位容量 | `/wms/inventoryManage/locationcapacity` | 库位容量配置 |
| 包装管理 | `/wms/inventoryManage/package` | 包装管理 |
| 交易记录 | `/wms/inventoryManage/transaction` | 库存交易记录 |
| 转移日志 | `/wms/inventoryManage/transferlog` | 库存转移日志 |
| 盘点计划 | `/wms/countManage/count/countPlanMain` | 盘点计划制定 |
| 盘点任务 | `/wms/countManage/count/countJobMain` | 盘点任务执行 |
| 盘点记录 | `/wms/countManage/count/countRecordMain` | 盘点结果记录 |
| 盘点请求 | `/wms/countManage/count/countRequestMain` | 盘点申请 |
| 调整记录 | `/wms/countManage/countadjust/countadjustRecordMain` | 库存调整记录 |
| 调整请求 | `/wms/countManage/countadjust/countadjustRequestMain` | 库存调整申请 |

#### 2.4.3 入库管理

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 采购收货 | `/wms/purchasereceiptManage/purchasereceipt/purchasereceiptJobMain` | 采购收货作业 |
| 采购收货记录 | `/wms/purchasereceiptManage/purchasereceipt/purchasereceiptRecordMain` | 采购收货记录 |
| 采购收货请求 | `/wms/purchasereceiptManage/purchasereceipt/purchasereceiptRequestMain` | 采购收货请求 |
| 采购收货报表 | `/wms/purchasereceiptManage/purchasereceipt/purchasereceiptReport` | 采购收货报表 |
| 采购收货记录-类型 | `/wms/purchasereceiptManage/purchasereceipt/purchasereceiptRecordMTypeMain` | 按类型收货记录 |
| 采购收货记录-拒绝 | `/wms/purchasereceiptManage/purchasereceipt/purchasereceiptRecordRefuseMain` | 收货拒绝记录 |
| 采购收货请求-订单 | `/wms/purchasereceiptManage/purchasereceipt/purchasereceiptRequestOrderMTypeMain` | 按订单收货请求 |
| IQC检验 | `/wms/purchasereceiptManage/inspect/inspectJobMain` | 来料检验任务 |
| IQC检验记录 | `/wms/purchasereceiptManage/inspect/inspectRecordMain` | 来料检验记录 |
| IQC检验请求 | `/wms/purchasereceiptManage/inspect/inspectRequestMain` | 来料检验请求 |
| 采购上架 | `/wms/purchasereceiptManage/putaway/putawayJobMain` | 采购上架作业 |
| 采购上架记录 | `/wms/purchasereceiptManage/putaway/putawayRecordMain` | 采购上架记录 |
| 采购上架请求 | `/wms/purchasereceiptManage/putaway/putawayRequestMain` | 采购上架请求 |
| 采购退货 | `/wms/purchasereceiptManage/purchasereturn/purchasereturnJobMain` | 采购退货作业 |
| 采购退货记录 | `/wms/purchasereceiptManage/purchasereturn/purchasereturnRecordMain` | 采购退货记录 |
| 采购退货请求 | `/wms/purchasereceiptManage/purchasereturn/purchasereturnRequestMain` | 采购退货请求 |
| 采购退货-按订单 | `/wms/purchasereceiptManage/purchasereturn/purchasereturnRecordMOrderTypeMain` | 按订单退货记录 |
| 采购退货请求-按订单 | `/wms/purchasereceiptManage/purchasereturn/purchasereturnRequestMOrderTypeMain` | 按订单退货请求 |
| 采购退货新请求 | `/wms/purchasereceiptManage/purchasereturn/purchasereturnRequestMainNew` | 采购退货新请求 |
| 采购退备件 | `/wms/purchasereceiptManage/purchasereturnspare/purchasereturnRecordSpareMain` | 备件退货记录 |
| 采购退备件请求 | `/wms/purchasereceiptManage/purchasereturnspare/purchasereturnRequestSpareMain` | 备件退货请求 |
| 备件收货 | `/wms/purchasereceiptManage/sparereceipt/sparereceiptJobMain` | 备件收货作业 |
| 备件收货记录 | `/wms/purchasereceiptManage/sparereceipt/sparereceiptRecordMain` | 备件收货记录 |
| 备件收货请求 | `/wms/purchasereceiptManage/sparereceipt/sparereceiptRequestMain` | 备件收货请求 |
| 供应商交货 | `/wms/purchasereceiptManage/supplierdeliver/supplierdeliverRequestMain` | 供应商交货请求 |
| 供应商交货记录 | `/wms/purchasereceiptManage/supplierdeliver/supplierdeliverRecordMain` | 供应商交货记录 |
| 供应商交货检验 | `/wms/purchasereceiptManage/supplierdeliver/supplierdeliverInspectionDetail` | 供应商交货检验 |
| 供应商交货标签 | `/wms/purchasereceiptManage/supplierdeliver/supplierdeliverRequestMain/labelForm` | 供应商交货标签表单 |
| 供应商交货基础表单 | `/wms/purchasereceiptManage/supplierdeliver/supplierdeliverRequestMain/supplierdeliverBasicForm` | 供应商交货基础表单 |
| 供应商包装 | `/wms/purchasereceiptManage/supplierdeliver/supplierPackage` | 供应商包装配置 |
| 供应商恢复 | `/wms/purchasereceiptManage/supplierdeliver/supplierResume` | 供应商恢复记录 |
| 采购主数据 | `/wms/purchasereceiptManage/supplierdeliver/purchaseMain` | 采购主数据 |
| 采购主数据WMS | `/wms/purchasereceiptManage/supplierdeliver/purchaseMainWms` | 采购主数据WMS版 |
| 采购计划 | `/wms/purchasereceiptManage/supplierdeliver/purchasePlanMain` | 采购计划主数据 |
| 需求预测-海诺F | `/wms/purchasereceiptManage/supplierdeliver/demandforecastingHainuoF` | 需求预测F |
| 需求预测-海诺P | `/wms/purchasereceiptManage/supplierdeliver/demandforecastingHainuoP` | 需求预测P |
| 需求预测主数据 | `/wms/purchasereceiptManage/supplierdeliver/demandforecastingMain` | 需求预测主数据 |
| 供应商需求预测F | `/wms/purchasereceiptManage/supplierdeliver/demandforecastingSupplierHainouF` | 供应商需求预测F |
| 供应商需求预测P | `/wms/purchasereceiptManage/supplierdeliver/demandforecastingSupplierHainouP` | 供应商需求预测P |
| 供应商需求预测主数据 | `/wms/purchasereceiptManage/supplierdeliver/demandforecastingSupplierMain` | 供应商需求预测主数据 |
| 生产入库 | `/wms/productionManage/productreceipt/productreceiptJobMain` | 生产入库作业 |
| 生产入库记录 | `/wms/productionManage/productreceipt/productreceiptRecordMain` | 生产入库记录 |
| 生产入库请求 | `/wms/productionManage/productreceipt/productreceiptRequestMain` | 生产入库请求 |
| 生产入库组装 | `/wms/productionManage/productreceiptAssemble/productreceiptAssembleJobMain` | 生产入库组装作业 |
| 生产入库组装记录 | `/wms/productionManage/productreceiptAssemble/productreceiptAssembleRecordMain` | 生产入库组装记录 |
| 生产入库组装请求 | `/wms/productionManage/productreceiptAssemble/productreceiptAssembleRequestMain` | 生产入库组装请求 |
| MES原材料消耗 | `/wms/productionManage/productreceiptAssemble/mesRawMaterialConsumptionInfo` | MES原材料消耗信息 |
| 原材料消耗信息 | `/wms/productionManage/productreceiptAssemble/rawMaterialConsumptionInfo` | 原材料消耗信息 |
| 生产上架 | `/wms/productionManage/productputaway/productputawayJobMain` | 生产上架作业 |
| 生产上架记录 | `/wms/productionManage/productputaway/productputawayRecordMain` | 生产上架记录 |
| 生产上架请求 | `/wms/productionManage/productputaway/productputawayRequestMain` | 生产上架请求 |
| 生产上架组装 | `/wms/productionManage/productputawayAssemble/productputawayAssembleJobMain` | 生产上架组装作业 |
| 生产上架组装记录 | `/wms/productionManage/productputawayAssemble/productputawayAssembleRecordMain` | 生产上架组装记录 |
| 生产上架组装请求 | `/wms/productionManage/productputawayAssemble/productputawayAssembleRequestMain` | 生产上架组装请求 |
| 生产退货 | `/wms/productionManage/productionreturn/productionreturnJobMain` | 生产退料作业 |
| 生产退货记录 | `/wms/productionManage/productionreturn/productionreturnRecordMain` | 生产退料记录 |
| 生产退货记录-HOLD | `/wms/productionManage/productionreturn/productionreturnRecordMainHold` | 生产退料HOLD记录 |
| 生产退货请求 | `/wms/productionManage/productionreturn/productionreturnRequestMain` | 生产退料请求 |
| 生产退货请求-无 | `/wms/productionManage/productionreturn/productionreturnRequestMainNo` | 生产退料请求(无) |
| 生产拆解 | `/wms/productionManage/productdismantle/productdismantleJobMain` | 产品拆解作业 |
| 生产拆解记录 | `/wms/productionManage/productdismantle/productdismantleRecordMain` | 产品拆解记录 |
| 生产拆解请求 | `/wms/productionManage/productdismantle/productdismantleRequestMain` | 产品拆解请求 |
| 生产在制品 | `/wms/productionManage/processproduction/processproductionRecord` | 在制品生产记录 |
| 生产在制品请求 | `/wms/productionManage/processproduction/processproductionRequest` | 在制品生产请求 |
| 生产计划 | `/wms/productionManage/productionplan/productionMain` | 生产计划主数据 |
| 生产计划-组装 | `/wms/productionManage/productionplan/productionMainAssemble` | 生产计划组装 |
| 生产计划-组装备件 | `/wms/productionManage/productionplan/productionMainAssembleSparePart` | 生产计划组装备件 |
| 生产计划-预测备件 | `/wms/productionManage/productionplan/productionMainPredictSparePart` | 生产计划预测备件 |
| QAD生产计划 | `/wms/productionManage/productionplan/qadproductionplan` | QAD生产计划 |
| 生产工单 | `/wms/productionManage/productionplan/workMain` | 生产工单 |
| 产品报废 | `/wms/productionManage/productscrap/productscrapJobMain` | 产品报废作业 |
| 产品报废记录 | `/wms/productionManage/productscrap/productscrapRecordMain` | 产品报废记录 |
| 产品报废请求 | `/wms/productionManage/productscrap/productscrapRequestMain` | 产品报废请求 |
| 生产入库报废 | `/wms/productionManage/productreceiptscrap/productreceiptscrapJobMain` | 生产入库报废作业 |
| 生产入库报废记录 | `/wms/productionManage/productreceiptscrap/productreceiptscrapRecordMain` | 生产入库报废记录 |
| 生产入库报废请求 | `/wms/productionManage/productreceiptscrap/productreceiptscrapRequestMain` | 生产入库报废请求 |
| 产品纠正 | `/wms/productionManage/productredress/productredressJobMain` | 产品纠正作业 |
| 产品纠正记录 | `/wms/productionManage/productredress/productredressRecordMain` | 产品纠正记录 |
| 产品纠正请求 | `/wms/productionManage/productredress/productredressRequestMain` | 产品纠正请求 |
| 产品维修 | `/wms/productionManage/productrepair/productrepairRequestMain` | 产品维修请求 |
| 产品维修记录 | `/wms/productionManage/productrepair/productrepairRecordMain` | 产品维修记录 |
| 生产离线结算 | `/wms/productionManage/offlinesettlement/offlinesettlementRecordMain` | 生产离线结算记录 |
| 生产离线结算请求 | `/wms/productionManage/offlinesettlement/offlinesettlementRequestMain` | 生产离线结算请求 |

#### 2.4.4 出库管理

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 销售发货请求 | `/wms/deliversettlementManage/saleShipmentMainRequest` | 销售发货请求 |
| 销售发货记录 | `/wms/deliversettlementManage/saleShipmentMainRecord` | 销售发货记录 |
| 领料作业 | `/wms/issueManage/issue/issueJobMain` | 生产领料作业 |
| 领料记录 | `/wms/issueManage/issue/issueRecordMain` | 领料记录 |
| 领料请求 | `/wms/issueManage/issue/issueRequestMain` | 领料请求 |
| 拣货作业 | `/wms/issueManage/pick/pickJobMain` | 拣货作业 |
| 拣货记录 | `/wms/issueManage/pick/pickRecordMain` | 拣货记录 |
| 拣货请求 | `/wms/issueManage/pick/pickRequestMain` | 拣货请求 |
| 备料计划 | `/wms/issueManage/preparetoissueplan/preparetoMain` | 备料计划 |
| 生产补料 | `/wms/issueManage/repleinsh/repleinshJobMain` | 生产补料作业 |
| 生产补料记录 | `/wms/issueManage/repleinsh/repleinshRecordMain` | 生产补料记录 |
| 生产补料请求 | `/wms/issueManage/repleinsh/repleinshRequestMain` | 生产补料请求 |
| 非计划出库作业 | `/wms/issueManage/unplannedissue/unplannedissueJobMain` | 非计划出库作业 |
| 非计划出库记录 | `/wms/issueManage/unplannedissue/unplannedissueRecordMain` | 非计划出库记录 |
| 非计划出库请求 | `/wms/issueManage/unplannedissue/unplannedissueRequestMain` | 非计划出库请求 |
| 生产退料作业 | `/wms/issueManage/productionreturn/productionreturnJobMain` | 生产退料作业 |
| 生产退料记录 | `/wms/issueManage/productionreturn/productionreturnRecordMain` | 生产退料记录 |
| 生产退料请求 | `/wms/issueManage/productionreturn/productionreturnRequestMain` | 生产退料请求 |
| 生产领料作业 | `/wms/issueManage/productionreceipt/productionreceiptJobMain` | 生产领料作业 |
| 生产领料记录 | `/wms/issueManage/productionreceipt/productionreceiptRecordMain` | 生产领料记录 |
| 生产报废记录 | `/wms/issueManage/productionscrap/productionscrapRecordMain` | 生产报废记录 |
| 生产报废请求 | `/wms/issueManage/productionscrap/productionscrapRequestMain` | 生产报废请求 |
| 在线结算信息 | `/wms/issueManage/onlinesettlement/onlinesettlementInfo` | 在线结算信息 |
| 在线结算明细 | `/wms/issueManage/onlinesettlement/onlinesettlementInfoDetail` | 在线结算明细 |
| 在线结算记录 | `/wms/issueManage/onlinesettlement/onlinesettlementRecordMain` | 在线结算记录 |
| 在线结算请求 | `/wms/issueManage/onlinesettlement/onlinesettlementRequestMain` | 在线结算请求 |
| 结算配置 | `/wms/issueManage/onlinesettlement/settlementConfiguration` | 结算配置 |
| 结算报表 | `/wms/issueManage/onlinesettlement/settlementReport` | 结算报表 |
| 结算报表SCP | `/wms/issueManage/onlinesettlement/settlementReportScp` | SCP结算报表 |
| 发货作业 | `/wms/deliversettlementManage/deliver/deliverJobMain` | 发货作业 |
| 发货记录 | `/wms/deliversettlementManage/deliver/deliverRecordMain` | 发货记录 |
| 发货请求 | `/wms/deliversettlementManage/deliver/deliverRequestMain` | 发货请求 |
| 发货计划 | `/wms/deliversettlementManage/deliverplan/deliverPlanMain` | 发货计划主数据 |
| 销售主数据 | `/wms/deliversettlementManage/deliverplan/saleMain` | 销售主数据 |
| 客户退货作业 | `/wms/deliversettlementManage/customerreturn/customerreturnJobMain` | 客户退货作业 |
| 客户退货记录 | `/wms/deliversettlementManage/customerreturn/customerreturnRecordMain` | 客户退货记录 |
| 客户退货请求 | `/wms/deliversettlementManage/customerreturn/customerreturnRequestMain` | 客户退货请求 |
| 客户收款请求 | `/wms/deliversettlementManage/customerreceipt/customerreceiptRequestMain` | 客户收款请求 |
| 客户结算请求 | `/wms/deliversettlementManage/customersettle/customersettleRequestMain` | 客户结算请求 |

#### 2.4.5 移库管理

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 移库作业 | `/wms/moveManage/inventorymove/inventorymoveJobMain` | 移库作业 |
| 移库记录 | `/wms/moveManage/inventorymove/inventorymoveRecordMain` | 移库记录 |
| 移库记录-新版 | `/wms/moveManage/inventorymove/inventorymoveRecordMainNew` | 移库记录新版 |
| 移库记录-备件 | `/wms/moveManage/inventorymove/inventorymoveRecordMainNewSparePart` | 移库记录备件版 |
| 移库记录-OKHOLD | `/wms/moveManage/inventorymove/inventorymoveRecordMainOKHOLD` | 移库OKHOLD记录 |
| 移库记录-结算 | `/wms/moveManage/inventorymove/inventorymoveRecordMainSettlement` | 移库结算记录 |
| 移库请求 | `/wms/moveManage/inventorymove/inventorymoveRequestMain` | 移库请求 |
| 移库请求-HOLD WIP | `/wms/moveManage/inventorymove/inventorymoveRequestMainHOLDWIP` | 移库HOLD WIP请求 |
| 移库请求-MOVE | `/wms/moveManage/inventorymove/inventorymoveRequestMainMOVE` | 移库MOVE请求 |
| 移库请求-OKHOLD | `/wms/moveManage/inventorymove/inventorymoveRequestMainOKHOLD` | 移库OKHOLD请求 |
| 库存变化记录 | `/wms/moveManage/inventorychange/inventorychangeRecordMain` | 库存变化记录 |
| 库存变化请求 | `/wms/moveManage/inventorychange/inventorychangeRequestMain` | 库存变化请求 |
| 物料变更 | `/wms/moveManage/itemChange` | 物料变更 |

#### 2.4.6 库存作业管理

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 容器主数据请求 | `/wms/inventoryjobManage/containermanage/containerMainRequest` | 容器主数据请求 |
| 容器记录主数据 | `/wms/inventoryjobManage/containermanage/containerRecordMain` | 容器记录主数据 |
| 创建容器请求 | `/wms/inventoryjobManage/containermanage/createContainerMainRequest` | 创建容器请求 |
| 创建容器记录 | `/wms/inventoryjobManage/containermanage/createContainerRecordMain` | 创建容器记录 |
| 交付容器记录 | `/wms/inventoryjobManage/containermanage/deliverContainerRecordMain` | 交付容器记录 |
| 初始化容器请求 | `/wms/inventoryjobManage/containermanage/initialContainerMainRequest` | 初始化容器请求 |
| 初始化容器记录 | `/wms/inventoryjobManage/containermanage/initialContainerRecordMain` | 初始化容器记录 |
| 移动容器请求 | `/wms/inventoryjobManage/containermanage/moveContainerMainRequest` | 移动容器请求 |
| 移动容器记录 | `/wms/inventoryjobManage/containermanage/moveContainerRecordMain` | 移动容器记录 |
| 归还容器请求 | `/wms/inventoryjobManage/containermanage/returnContainerMainRequest` | 归还容器请求 |
| 归还容器记录 | `/wms/inventoryjobManage/containermanage/returnContainerRecordMain` | 归还容器记录 |
| 报废容器请求 | `/wms/inventoryjobManage/containermanage/scrapContainerMainRequest` | 报废容器请求 |
| 报废容器记录 | `/wms/inventoryjobManage/containermanage/scrapContainerRecordMain` | 报废容器记录 |
| 库存初始化记录 | `/wms/inventoryjobManage/inventoryinitial/inventoryinitRecordMain` | 库存初始化记录 |
| 库存初始化请求 | `/wms/inventoryjobManage/inventoryinitial/inventoryinitRequestMain` | 库存初始化请求 |
| 包装合并 | `/wms/inventoryjobManage/packageManage/packagemergeMain` | 包装合并主数据 |
| 包装拆分 | `/wms/inventoryjobManage/packageManage/packagesplitMain` | 包装拆分主数据 |
| 包装合并TU | `/wms/inventoryjobManage/packageManage/packagetuomergeMain` | TU包装合并主数据 |
| 包装超限主数据 | `/wms/inventoryjobManage/packageManage/packageoverMain/packageoverJobMain` | 包装超限作业 |
| 包装超限记录 | `/wms/inventoryjobManage/packageManage/packageoverMain/packageoverRecordMain` | 包装超限记录 |
| 包装超限请求 | `/wms/inventoryjobManage/packageManage/packageoverMain/packageoverRequestMain` | 包装超限请求 |
| 包装超限追溯 | `/wms/inventoryjobManage/packageManage/packageoverMain/packageoverRetrospect` | 包装超限追溯 |
| 报废作业 | `/wms/inventoryjobManage/scrap/scrapJobMain` | 报废作业主数据 |
| 报废记录 | `/wms/inventoryjobManage/scrap/scrapRecordMain` | 报废记录主数据 |
| 报废请求 | `/wms/inventoryjobManage/scrap/scrapRequestMain` | 报废请求主数据 |
| 备件退料记录 | `/wms/inventoryjobManage/sparepartReturn/sparepartReturnRecordMain` | 备件退料记录 |
| 备件退料请求 | `/wms/inventoryjobManage/sparepartReturn/sparepartReturnRequestMain` | 备件退料请求 |
| 备件申请作业 | `/wms/inventoryjobManage/sparepartsrequisition/sparepartsrequisitionJobMain` | 备件申请作业 |
| 备件申请记录 | `/wms/inventoryjobManage/sparepartsrequisition/sparepartsrequisitionRecordMain` | 备件申请记录 |
| 备件申请请求 | `/wms/inventoryjobManage/sparepartsrequisition/sparepartsrequisitionRequestMain` | 备件申请请求 |
| 转移出库作业 | `/wms/inventoryjobManage/transferissue/transferissueJobMain` | 转移出库作业 |
| 转移出库记录 | `/wms/inventoryjobManage/transferissue/transferissueRecordMain` | 转移出库记录 |
| 转移出库请求 | `/wms/inventoryjobManage/transferissue/transferissueRequestMain` | 转移出库请求 |
| 转移入库作业 | `/wms/inventoryjobManage/transferreceipt/transferreceiptJobMain` | 转移入库作业 |
| 转移入库记录 | `/wms/inventoryjobManage/transferreceipt/transferreceiptRecordMain` | 转移入库记录 |
| 转移入库请求 | `/wms/inventoryjobManage/transferreceipt/transferreceiptRequestMain` | 转移入库请求 |
| 非计划入库作业 | `/wms/inventoryjobManage/unplannedreceipt/unplannedreceiptJobMain` | 非计划入库作业 |
| 非计划入库记录 | `/wms/inventoryjobManage/unplannedreceipt/unplannedreceiptRecordMain` | 非计划入库记录 |
| 非计划入库请求 | `/wms/inventoryjobManage/unplannedreceipt/unplannedreceiptRequestMain` | 非计划入库请求 |

#### 2.4.7 结算管理

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 客户收款记录 | `/wms/deliversettlementManage/customerreceipt/customerreceiptRecordMain` | 客户收款记录 |
| 客户收款请求 | `/wms/deliversettlementManage/customerreceipt/customerreceiptRequestMain` | 客户收款请求 |
| 客户结算记录 | `/wms/deliversettlementManage/customersettle/customersettleRecordMain` | 客户结算记录 |
| 客户结算请求 | `/wms/deliversettlementManage/customersettle/customersettleRequestMain` | 客户结算请求 |
| 发货结算作业 | `/wms/deliversettlementManage/deliver/deliverJobMain` | 发货结算作业 |
| 发货结算记录 | `/wms/deliversettlementManage/deliver/deliverRecordMain` | 发货结算记录 |
| 发货结算请求 | `/wms/deliversettlementManage/deliver/deliverRequestMain` | 发货结算请求 |
| 离线结算记录 | `/wms/deliversettlementManage/offlinesettlement/offlinesettlementRecordMain` | 离线结算记录 |
| 离线结算请求 | `/wms/deliversettlementManage/offlinesettlement/offlinesettlementRequestMain` | 离线结算请求 |
| 备货作业 | `/wms/deliversettlementManage/stockup/stockupMainJob` | 备货作业主数据 |
| 备货记录 | `/wms/deliversettlementManage/stockup/stockupMainRecord` | 备货记录主数据 |
| 备货请求 | `/wms/deliversettlementManage/stockup/stockupMainRequest` | 备货请求主数据 |

#### 2.4.8 供应商管理

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 供应商账期日历 | `/wms/supplierManage/supplierApbalanceCalendar` | 供应商账期日历 |
| 供应商AP账期主数据 | `/wms/supplierManage/supplierApbalance/supplierApbalanceMain` | 供应商AP账期主数据 |
| 供应商AP账期明细 | `/wms/supplierManage/supplierApbalance/supplierApbalanceDetail` | 供应商AP账期明细 |
| 供应商发票记录 | `/wms/supplierManage/supplierinvoice/supplierinvoiceRecordMain` | 供应商发票记录 |
| 供应商发票请求 | `/wms/supplierManage/supplierinvoice/supplierinvoiceRequestMain` | 供应商发票请求 |
| 供应商发票差异 | `/wms/supplierManage/supplierinvoice/supplierinvoiceRequestMainDifference` | 供应商发票差异 |
| 供应商发票已开票离散 | `/wms/supplierManage/supplierinvoiceInvoicedDiscrete` | 已开票离散 |
| 供应商发票已排程 | `/wms/supplierManage/supplierinvoiceInvoicedSchedule` | 已排程发票 |
| 供应商发票已删除 | `/wms/supplierManage/supplierinvoiceInvoicedScheduleDeleted` | 已删除排程发票 |
| 供应商采购索赔记录 | `/wms/supplierManage/purchaseclaim/purchaseclaimRecordMain` | 采购索赔记录 |
| 供应商采购索赔请求 | `/wms/supplierManage/purchaseclaim/purchaseclaimRequestMain` | 采购索赔请求 |
| 采购离散订单 | `/wms/supplierManage/purchaseDiscreteOrder/purchaseDiscreteOrderMain` | 采购离散订单主数据 |
| 供应商开票日历 | `/wms/supplierManage/invoicingcalendar` | 供应商开票日历 |

#### 2.4.9 MES条码对接

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| MES条码 | `/wms/buttMesManage/mesBarCode` | MES条码对接 |

#### 2.4.10 对账管理

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 非包装交易余额 | `/wms/reconciliation/notPackageTransactionBalance` | 非包装交易余额 |
| 包装交易余额 | `/wms/reconciliation/transactionBalancePackage` | 包装交易余额 |

#### 2.4.11 报表中心

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 报表清单 | `/wms/reportList` | WMS报表清单 |

#### 2.4.12 AGV管理

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| AGV库位关系 | `/wms/agvManage/agvLocationrelation` | AGV与库位映射 |
| 反冲明细 | `/wms/agvManage/backflushDetailbQad` | 反冲明细 |
| AGV接口信息 | `/wms/agvManage/interfaceInfo` | AGV接口配置 |
| AGV推荐库位历史 | `/wms/agvManage/recommendLocationHistory` | 推荐库位历史 |

### 2.5 QMS质量管理模块

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| AQL配置 | `/qms/aql` | AQL抽样标准配置 |
| 计数器 | `/qms/counter` | 质量计数管理 |
| 动态规则 | `/qms/dynamicRule` | 动态检验规则 |
| 检验方案 | `/qms/inspectionScheme` | 检验方案配置 |
| 检验方案新增 | `/qms/inspectionScheme/addForm` | 检验方案新增表单 |
| 检验方案详情 | `/qms/inspectionScheme/detail` | 检验方案详情页 |
| 检验模板 | `/qms/inspectionTemplate` | 检验模板管理 |
| 检验模板新增 | `/qms/inspectionTemplate/addForm` | 检验模板新增表单 |
| 检验模板详情 | `/qms/inspectionTemplate/detail` | 检验模板详情页 |
| 检验方法 | `/qms/inspectionMethod` | 检验方法定义 |
| 检验阶段 | `/qms/inspectionStage` | 检验阶段配置 |
| 检验特性 | `/qms/inspectionQ1` | 定性检验特性 |
| 检验特性 | `/qms/inspectionQ2` | 定量检验特性 |
| 检验特性 | `/qms/inspectionQ3` | 检验特性Q3 |
| 抽样方案 | `/qms/samplingScheme` | 抽样方案配置 |
| 抽样过程 | `/qms/samplingProcess` | 抽样过程管理 |
| 抽样代码 | `/qms/sampleCode` | 抽样代码管理 |
| 选中的项目 | `/qms/selectedProject` | 选中项目配置 |
| 选中的集合 | `/qms/selectedSet` | 选中集合管理 |
| 检验请求 | `/qms/inspectionRequest` | 检验申请单 |
| 检验任务 | `/qms/inspectionJob` | 检验任务分配 |
| 检验任务新增 | `/qms/inspectionJob/addForm` | 检验任务新增表单 |
| 检验任务详情 | `/qms/inspectionJob/detail` | 检验任务详情页 |
| 生产检验任务 | `/qms/inspectionJob/inspectionJobProduction` | 生产检验任务页 |
| 采购检验任务 | `/qms/inspectionJob/inspectionJobPurchase` | 采购检验任务页 |
| 检验记录 | `/qms/inspectionRecord` | 检验结果记录 |
| 检验记录新增 | `/qms/inspectionRecord/addForm` | 检验记录新增表单 |
| 检验记录详情 | `/qms/inspectionRecord/detail` | 检验记录详情页 |
| 检验记录通用新增 | `/qms/inspectionRecord/useAddForm` | 检验记录通用新增表单 |
| 生产检验记录 | `/qms/inspectionRecord/inspectionRecordProduction` | 生产检验记录页 |
| 采购检验记录 | `/qms/inspectionRecord/inspectionRecordPurchase` | 采购检验记录页 |
| 首件检验 | `/qms/inspectionRecordFirst` | 首件检验记录 |
| 首件检验新增 | `/qms/inspectionRecordFirst/addForm` | 首件检验新增表单 |
| 首件检验详情 | `/qms/inspectionRecordFirst/detail` | 首件检验详情页 |
| 首件检验通用新增 | `/qms/inspectionRecordFirst/useAddForm` | 首件检验通用新增表单 |
| 质量通知 | `/qms/qualityNotice` | 质量通知单 |
| 质量通知新增 | `/qms/qualityNotice/addForm` | 质量通知新增表单 |
| 质量通知类型组件 | `/qms/qualityNotice/components/notaicType` | 质量通知类型组件 |
| 质量批次组件 | `/qms/qualityNotice/components/qualityBatch` | 质量批次组件 |
| 质量通知组件 | `/qms/qualityNotice/components/qualityNotice` | 质量通知组件 |

### 2.6 EAM设备管理模块

#### 2.6.1 设备基础档案

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 设备台账 | `/eam/equipmentAccounts` | 设备档案管理 |
| 设备台账新增 | `/eam/equipmentAccounts/ablesForm` | 设备台账新增表单 |
| 设备主要部件 | `/eam/equipmentMainPart` | 设备主要部件 |
| 设备供应商 | `/eam/equipmentSupplier` | 设备供应商 |
| 设备制造商 | `/eam/equipmentManufacturer` | 设备制造商 |
| 设备分类 | `/eam/classTypeRole` | 设备分类角色 |
| 设备工具备件 | `/eam/equipmentToolSparePart` | 工具备件关联 |
| 设备文档类型 | `/eam/documentType` | 文档类型配置 |
| 文档选择集 | `/eam/documentTypeSelectSet` | 文档选择集 |
| 文档选择集表单 | `/eam/documentTypeSelectSet/itemSelectSetForm` | 文档选择集表单 |
| 表格数据扩展 | `/eam/tableDataExtendedAttribute` | 扩展属性配置 |

#### 2.6.2 设备巡检

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 点检计划 | `/eam/planSpotCheck` | 设备点检计划 |
| 点检明细 | `/eam/equipmentSpotCheckMain` | 点检任务明细 |
| 点检记录 | `/eam/equipmentSpotCheckRecordMain` | 点检记录 |
| 点检项 | `/eam/spotCheckItem` | 点检项目配置 |
| 点检选项 | `/eam/basicSpotCheckOption` | 点检选项配置 |
| 点检选项选择集 | `/eam/basicSpotCheckOption/itemSelectSetForm` | 点检选项选择集表单 |
| 点检选择集 | `/eam/spotCheckSelectSet` | 点检选择集 |
| 点检选择集表单 | `/eam/spotCheckSelectSet/itemSelectSetForm` | 点检选择集表单 |
| 点检明细记录 | `/eam/equipmentSpotCheckDetail` | 点检明细记录 |
| 点检审核表单 | `/eam/planSpotCheck/audiForm` | 点检审核表单 |
| 点检订单详情 | `/eam/equipmentSpotCheckMain/SpotCheckOrderDetail` | 点检订单详情页 |
| 点检完工表单1 | `/eam/equipmentSpotCheckMain/finishForm1` | 点检完工表单一 |
| 点检完工表单2 | `/eam/equipmentSpotCheckMain/finishForm2` | 点检完工表单二 |
| 点检记录主页 | `/eam/equipmentSpotCheckRecordMain/index` | 点检记录主页 |
| 点检记录明细 | `/eam/equipmentSpotCheckRecordDetail/index` | 点检记录明细页 |

#### 2.6.3 设备保养

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 保养计划 | `/eam/planInspection` | 设备保养计划 |
| 保养明细 | `/eam/equipmentInspectionMain` | 保养任务明细 |
| 保养记录 | `/eam/equipmentInspectionRecordMain` | 保养记录 |
| 保养项目 | `/eam/maintenanceItem` | 保养项目配置 |
| 保养选项 | `/eam/basicMaintenanceOption` | 保养选项配置 |
| 保养选项选择集 | `/eam/basicMaintenanceOption/itemSelectSetForm` | 保养选项选择集表单 |
| 保养选择集 | `/eam/maintenanceItemSelectSet` | 保养选择集 |
| 保养选择集表单 | `/eam/maintenanceItemSelectSet/itemSelectSetForm` | 保养选择集表单 |
| 保养明细记录 | `/eam/equipmentInspectionDetail` | 保养明细记录 |
| 保养审核表单 | `/eam/planInspection/audiForm` | 保养审核表单 |
| 保养管理 | `/eam/maintenance/index` | 保养管理主页 |
| 设备保养主页 | `/eam/equipmentMaintenanceMain/index` | 设备保养主页 |
| 设备保养订单详情 | `/eam/equipmentMaintenanceMain/MaintenanceOrderDetail` | 设备保养订单详情 |
| 设备保养完工表单1 | `/eam/equipmentMaintenanceMain/finishForm1` | 设备保养完工表单一 |
| 设备保养完工表单2 | `/eam/equipmentMaintenanceMain/finishForm2` | 设备保养完工表单二 |
| 设备保养记录主页 | `/eam/equipmentMaintenanceRecordMain/index` | 设备保养记录主页 |
| 设备保养记录明细 | `/eam/equipmentMaintenanceRecordDetail/index` | 设备保养记录明细页 |
| 保养审核 | `/eam/maintenance/audiForm` | 保养审核表单 |
| 保养订单详情 | `/eam/equipmentInspectionMain/InspectionOrderDetail` | 保养订单详情页 |
| 保养完工表单1 | `/eam/equipmentInspectionMain/finishForm1` | 保养完工表单一 |
| 保养完工表单2 | `/eam/equipmentInspectionMain/finishForm2` | 保养完工表单二 |
| 保养记录主页 | `/eam/equipmentInspectionRecordMain/index` | 保养记录主页 |
| 保养记录明细 | `/eam/equipmentInspectionRecordDetail/index` | 保养记录明细页 |
| 保养经验 | `/eam/maintainExperience` | 保养经验库 |

#### 2.6.4 设备维修

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 报修申请 | `/eam/equipmentReportRepairRequest` | 设备报修申请 |
| 报修审核表单 | `/eam/equipmentReportRepairRequest/audiForm` | 报修审核表单 |
| 维修工单主页 | `/eam/equipmentRepairJobMain/index` | 维修工单主页 |
| 维修工单详情 | `/eam/equipmentRepairJobMain/EquipmentRepairJobDetail` | 维修工单详情页 |
| 维修完工表单1 | `/eam/equipmentRepairJobMain/finishForm1` | 维修完工表单一 |
| 维修完工表单2 | `/eam/equipmentRepairJobMain/finishForm2` | 维修完工表单二 |
| 维修完工表单3 | `/eam/equipmentRepairJobMain/finishForm3` | 维修完工表单三 |
| 维修转移表单 | `/eam/equipmentRepairJobMain/transferForm` | 维修转移表单 |
| 维修工单明细 | `/eam/equipmentRepairJobDetail/index` | 维修工单明细页 |
| 维修记录主页 | `/eam/equipmentRepairRecordMain/index` | 维修记录主页 |
| 维修记录明细 | `/eam/equipmentRepairRecordDetail/index` | 维修记录明细页 |
| 维修工单 | `/eam/equipmentRepairJobMain` | 维修工单管理 |
| 维修记录 | `/eam/equipmentRepairRecordMain` | 维修执行记录 |
| 维修备件 | `/eam/repairSparePartsRequest` | 维修备件申请 |
| 维修经验 | `/eam/repairExperience` | 维修经验库 |
| 故障类型 | `/eam/basicFaultType` | 故障类型定义 |
| 故障原因 | `/eam/basicFaultCause` | 故障原因分析 |
| 维修工单明细 | `/eam/equipmentRepairJobDetail` | 维修工单明细 |
| 维修记录明细 | `/eam/equipmentRepairRecordDetail` | 维修记录明细 |
| 备件出库记录 | `/eam/repairSparePartsRecord` | 备件出库记录 |

#### 2.6.5 设备停机

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 设备停机 | `/eam/equipmentShutdown` | 设备停机记录 |
| 设备签到 | `/eam/equipmentSigning` | 设备签到管理 |

#### 2.6.6 设备转移

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 设备转移记录 | `/eam/equipmentTransferRecord` | 设备转移记录 |
| 设备变更记录 | `/eam/recordDeviceChanged` | 设备变更记录 |

#### 2.6.7 备件管理

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 备件档案 | `/eam/sparePart` | 备件库存档案 |
| 备件申请 | `/eam/sparePartsApplyMain` | 备件申请单 |
| 备件入库 | `/eam/sparepartsinlocation` | 备件入库记录 |
| 备件出库 | `/eam/sparepartsoutlocation` | 备件出库记录 |
| 备件位置 | `/eam/itemInLocation` | 备件位置管理 |
| 位置管理 | `/eam/location/index` | 位置管理主页 |
| 位置区域 | `/eam/locationArea/index` | 位置区域管理 |
| 物品管理 | `/eam/item/index` | 物品管理主页 |
| 物品出库位置 | `/eam/itemOutLocation/index` | 物品出库位置管理 |
| 备件出库位置 | `/eam/SparePartsOutLocationRecord` | 备件出库位置记录 |
| 备件入库位置 | `/eam/sparePartsInLocationRecord` | 备件入库位置记录 |
| 位置替换 | `/eam/itemLocationReplace` | 备件位置替换 |
| 物品台账 | `/eam/itemAccounts` | 物品台账管理 |
| 物品申请 | `/eam/itemApplyMain` | 物品申请单 |
| 物品删除 | `/eam/itemDelete` | 物品删除申请 |
| 物品保养 | `/eam/itemMaintenance` | 物品保养记录 |
| 物品订单 | `/eam/itemOrderMain` | 物品订单管理 |

#### 2.6.8 工装管理

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 工具台账 | `/eam/toolAccounts` | 工具资产台账 |
| 工具变更记录 | `/eam/toolChangedRecord` | 工具变更记录 |
| 工具型号 | `/eam/toolMod/index` | 工具型号管理 |
| 工具操作表单 | `/eam/toolMod/operateForm` | 工具操作表单 |
| 工具签到 | `/eam/toolSigning/index` | 工具签到管理 |
| 工具入库 | `/eam/toolEquipmentIn/index` | 工具设备入库 |
| 工具出库 | `/eam/toolEquipmentOut/index` | 工具设备出库 |
| 交易记录 | `/eam/transaction/index` | 交易记录管理 |

#### 2.6.9 产线管理

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 产线基础 | `/eam/basicEamProductionline` | 产线基础档案 |
| 车间基础 | `/eam/basicEamWorkshop` | 车间基础档案 |
| 关系主要部件 | `/eam/relationMainPart` | 关系主要部件 |

#### 2.6.10 其他

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 盘点调整计划 | `/eam/countadjustPlan` | 盘点调整计划 |
| 盘点调整工单 | `/eam/countadjustWork` | 盘点调整工单 |
| 盘点记录 | `/eam/countRecord` | 盘点执行记录 |
| 调整记录 | `/eam/adjustRecord` | 调整记录管理 |
| 申请记录 | `/eam/applicationRecord` | 申请记录管理 |
| 检验项目 | `/eam/inspectionItem` | 检验项目配置 |
| 检验项选择集 | `/eam/inspectionItemSelectSet` | 检验项选择集 |
| 检验项选择集表单 | `/eam/inspectionItemSelectSet/itemSelectSetForm` | 检验项选择集表单 |
| 备件出库位置记录 | `/eam/SparePartsOutLocationRecord/index` | 备件出库位置记录 |
| 调整记录 | `/eam/adjustRecord/index` | 调整记录管理 |
| 申请记录 | `/eam/applicationRecord/index` | 申请记录管理 |

### 2.7 BPM流程管理模块

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 流程模型 | `/bpm/model` | BPMN流程设计 |
| 模型编辑器 | `/bpm/model/editor/index` | BPMN模型编辑器 |
| 模型表单 | `/bpm/model/ModelForm` | 模型表单页 |
| 模型导入表单 | `/bpm/model/ModelImportForm` | 模型导入表单 |
| 流程定义 | `/bpm/definition` | 流程定义管理 |
| 表单管理 | `/bpm/form` | 流程表单配置 |
| 表单编辑器 | `/bpm/form/editor/index` | 表单设计器 |
| 用户组 | `/bpm/group` | 用户组管理 |
| 用户组表单 | `/bpm/group/UserGroupForm` | 用户组表单 |
| 任务分配规则 | `/bpm/taskAssignRule` | 任务分配规则 |
| 分配规则表单 | `/bpm/taskAssignRule/TaskAssignRuleForm` | 分配规则表单 |
| 流程实例 | `/bpm/processInstance` | 流程实例查询 |
| 流程实例创建 | `/bpm/processInstance/create/index` | 流程实例创建 |
| 流程实例详情 | `/bpm/processInstance/detail/index` | 流程实例详情页 |
| 流程实例BPMN查看器 | `/bpm/processInstance/detail/ProcessInstanceBpmnViewer` | 流程BPMN查看组件 |
| 流程实例任务列表 | `/bpm/processInstance/detail/ProcessInstanceTaskList` | 流程任务列表组件 |
| 任务退回表单 | `/bpm/processInstance/detail/TaskReturnDialogForm` | 任务退回表单组件 |
| 任务更新处理人表单 | `/bpm/processInstance/detail/TaskUpdateAssigneeForm` | 任务更新处理人表单 |
| 待办任务 | `/bpm/task/todo` | 待办任务列表 |
| 已办任务 | `/bpm/task/done` | 已办任务列表 |
| 已办任务详情 | `/bpm/task/done/TaskDetail` | 已办任务详情页 |
| OA请假 | `/bpm/oa/leave` | OA请假申请 |

### 2.8 SYSTEM系统管理模块

#### 2.8.1 认证授权

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 用户管理 | `/system/user` | 用户CRUD、导入导出 |
| 用户表单 | `/system/user/UserForm` | 用户新增/编辑表单 |
| 用户导入表单 | `/system/user/UserImportForm` | 用户批量导入表单 |
| 用户分配角色表单 | `/system/user/UserAssignRoleForm` | 用户分配角色表单 |
| 部门树形组件 | `/system/user/DeptTree` | 部门树形选择组件 |
| 角色管理 | `/system/role` | 角色权限配置 |
| 角色表单 | `/system/role/RoleForm` | 角色新增/编辑表单 |
| 角色分配菜单表单 | `/system/role/RoleAssignMenuForm` | 角色分配菜单表单 |
| 角色数据权限表单 | `/system/role/RoleDataPermissionForm` | 角色数据权限表单 |
| 菜单管理 | `/system/menu` | 菜单路由配置 |
| 菜单表单 | `/system/menu/MenuForm` | 菜单新增/编辑表单 |
| 部门管理 | `/system/dept` | 组织架构管理 |
| 部门表单 | `/system/dept/DeptForm` | 部门新增/编辑表单 |
| 岗位管理 | `/system/post` | 岗位配置 |
| 岗位表单 | `/system/post/PostForm` | 岗位新增/编辑表单 |
| 数据字典 | `/system/dict` | 字典类型管理 |
| 字典类型表单 | `/system/dict/DictTypeForm` | 字典类型新增/编辑表单 |
| 字典数据 | `/system/dict/data` | 字典值管理 |
| 字典数据表单 | `/system/dict/data/DictDataForm` | 字典数据新增/编辑表单 |

#### 2.8.2 租户管理

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 租户管理 | `/system/tenant` | 多租户配置 |
| 租户表单 | `/system/tenant/TenantForm` | 租户新增/编辑表单 |
| 租户套餐 | `/system/tenantPackage` | 租户权限套餐 |
| 租户套餐表单 | `/system/tenantPackage/TenantPackageForm` | 租户套餐新增/编辑表单 |
| 系统安装包 | `/system/systemInstallPackage` | 系统安装包 |

#### 2.8.3 消息管理

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 短信渠道 | `/system/sms/channel` | 短信通道配置 |
| 短信渠道表单 | `/system/sms/channel/SmsChannelForm` | 短信渠道新增/编辑表单 |
| 短信模板 | `/system/sms/template` | 短信模板管理 |
| 短信模板表单 | `/system/sms/template/SmsTemplateForm` | 短信模板新增/编辑表单 |
| 短信发送表单 | `/system/sms/template/SmsTemplateSendForm` | 短信发送表单 |
| 短信日志 | `/system/sms/log` | 短信发送记录 |
| 短信日志详情 | `/system/sms/log/SmsLogDetail` | 短信日志详情 |
| 邮件账号 | `/system/mail/account` | 邮件账户配置 |
| 邮件账号表单 | `/system/mail/account/MailAccountForm` | 邮件账号新增/编辑表单 |
| 邮件账号详情 | `/system/mail/account/MailAccountDetail` | 邮件账号详情 |
| 邮件模板 | `/system/mail/template` | 邮件模板管理 |
| 邮件模板表单 | `/system/mail/template/MailTemplateForm` | 邮件模板新增/编辑表单 |
| 邮件发送表单 | `/system/mail/template/MailTemplateSendForm` | 邮件发送表单 |
| 邮件日志 | `/system/mail/log` | 邮件发送记录 |
| 邮件日志详情 | `/system/mail/log/MailLogDetail` | 邮件日志详情 |
| 站内消息 | `/system/notify/message` | 系统消息推送 |
| 消息详情 | `/system/notify/message/NotifyMessageDetail` | 消息推送详情 |
| 消息模板 | `/system/notify/template` | 消息模板配置 |
| 消息模板表单 | `/system/notify/template/NotifyTemplateForm` | 消息模板新增/编辑表单 |
| 消息发送表单 | `/system/notify/template/NotifyTemplateSendForm` | 消息发送表单 |
| 我的消息 | `/system/notify/my` | 个人消息查看 |
| 我的消息详情 | `/system/notify/my/MyNotifyMessageDetail` | 我的消息详情 |
| 消息设置 | `/system/messageSet` | 消息推送设置 |

#### 2.8.4 权限配置

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| OAuth2客户端 | `/system/oauth2/client` | OAuth2应用配置 |
| OAuth2客户端表单 | `/system/oauth2/client/ClientForm` | OAuth2客户端新增/编辑表单 |
| OAuth2 Token | `/system/oauth2/token` | Token管理 |
| 数据权限 | `/system/role/dataPermission` | 数据权限配置 |
| 菜单权限 | `/system/role/assignMenu` | 角色菜单分配 |

#### 2.8.5 日志审计

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 登录日志 | `/system/loginlog` | 用户登录记录 |
| 登录日志详情 | `/system/loginlog/LoginLogDetail` | 登录日志详情 |
| 操作日志 | `/system/operatelog` | 操作审计日志 |
| 操作日志详情 | `/system/operatelog/OperateLogDetail` | 操作日志详情 |
| 错误码管理 | `/system/errorCode` | 错误码配置 |
| 错误码表单 | `/system/errorCode/ErrorCodeForm` | 错误码新增/编辑表单 |
| 敏感词 | `/system/sensitiveWord` | 敏感词过滤 |
| 敏感词表单 | `/system/sensitiveWord/SensitiveWordForm` | 敏感词新增/编辑表单 |
| 敏感词测试表单 | `/system/sensitiveWord/SensitiveWordTestForm` | 敏感词测试表单 |
| 序列号规则 | `/system/serialNumber` | 编码规则配置 |
| 序列号表单 | `/system/serialNumber/SerialNumberForm` | 序列号规则新增/编辑表单 |

#### 2.8.6 系统配置

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 区域管理 | `/system/area` | 行政区域管理 |
| 区域表单 | `/system/area/AreaForm` | 区域新增/编辑表单 |
| 通知公告 | `/system/notice` | 系统公告发布 |
| 通知公告表单 | `/system/notice/NoticeForm` | 通知公告新增/编辑表单 |
| 密码规则 | `/system/passwordRule` | 密码安全策略 |
| 表单关联 | `/system/tableActionRel` | 表单动作关联 |

### 2.9 INFRA基础设施模块

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 文件管理 | `/infra/file` | 文件上传下载 |
| 文件表单 | `/infra/file/FileForm` | 文件上传表单 |
| 文件配置 | `/infra/fileConfig` | 文件存储配置 |
| 文件配置表单 | `/infra/fileConfig/FileConfigForm` | 文件配置新增/编辑表单 |
| 代码生成 | `/infra/codegen` | CRUD代码生成 |
| 代码生成-基本信息表单 | `/infra/codegen/components/BasicInfoForm` | 代码生成基本信息表单 |
| 代码生成-列信息表单 | `/infra/codegen/components/ColumInfoForm` | 代码生成列信息配置表单 |
| 代码生成-生成信息表单 | `/infra/codegen/components/GenerateInfoForm` | 代码生成信息配置表单 |
| 代码生成-编辑表 | `/infra/codegen/editTable` | 代码生成编辑表页面 |
| 代码生成-导入表 | `/infra/codegen/importTable` | 代码生成导入表页面 |
| 代码生成-预览 | `/infra/codegen/PreviewCode` | 代码预览页面 |
| 数据源配置 | `/infra/dataSourceConfig` | 多数据源配置 |
| 数据源配置表单 | `/infra/dataSourceConfig/DataSourceConfigForm` | 数据源配置新增/编辑表单 |
| 数据库文档 | `/infra/dbDoc` | 数据库文档 |
| 定时任务 | `/infra/job` | Quartz任务配置 |
| 任务详情 | `/infra/job/JobDetail` | 定时任务详情 |
| 任务表单 | `/infra/job/JobForm` | 定时任务新增/编辑表单 |
| 任务日志 | `/infra/job/logger` | 任务执行日志 |
| 任务日志详情 | `/infra/job/logger/JobLogDetail` | 任务日志详情 |
| Redis管理 | `/infra/redis` | Redis缓存监控 |
| API访问日志 | `/infra/apiAccessLog` | API访问日志 |
| API访问日志详情 | `/infra/apiAccessLog/ApiAccessLogDetail` | API访问日志详情 |
| API错误日志 | `/infra/apiErrorLog` | API错误日志 |
| API错误日志详情 | `/infra/apiErrorLog/ApiErrorLogDetail` | API错误日志详情 |
| 配置管理 | `/infra/config` | 系统配置管理 |
| 配置表单 | `/infra/config/ConfigForm` | 配置新增/编辑表单 |
| Druid监控 | `/infra/druid` | Druid数据库监控 |
| Swagger文档 | `/infra/swagger` | Swagger API文档 |
| WebSocket | `/infra/webSocket` | WebSocket配置 |
| 自定义接口 | `/infra/customInterface` | 自定义接口管理 |
| 外部API历史 | `/infra/outerApiHis` | 外部API调用历史 |
| 服务器监控 | `/infra/server` | 服务器状态监控 |
| 构建配置 | `/infra/build` | 前端构建配置 |
| SkyWalking | `/infra/skywalking` | 链路追踪配置 |
| 测试示例 | `/infra/testDemo` | 开发测试页面 |

### 2.10 REPORT报表模块

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| GoView大屏 | `/report/goview` | 数据可视化大屏 |

### 2.11 个人中心

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 个人中心 | `/profile` | 用户信息展示 |
| 基本信息 | `/profile/basic` | 基础信息修改 |
| 头像上传 | `/profile/avatar` | 头像设置 |
| 密码修改 | `/profile/resetPwd` | 密码修改 |
| 社交账号 | `/profile/social` | 第三方绑定 |

### 2.12 错误页面

| 页面 | 路径 | 功能说明 |
|-----|------|---------|
| 403 | `/error/403` | 无权限页面 |
| 404 | `/error/404` | 资源不存在 |
| 500 | `/error/500` | 服务器错误 |

---

## 3. 前端API接口清单

### 3.1 MES模块API

| API文件 | 接口数 | 功能说明 |
|--------|-------|---------|
| abilityInfo | 7 | 能力矩阵管理 |
| hrPersonAbility | - | 人员能力 |
| opersteps | - | 工序管理 |
| operstepsType | - | 工序类型 |

### 3.2 WMS模块API

| API文件 | 接口数 | 功能说明 |
|--------|-------|---------|
| accountcalendar | - | 账户日历 |
| backflushRecordDetailb | - | 反冲记录 |
| bom | - | BOM管理 |
| carrier | - | 承运商 |
| condition | - | 条件配置 |
| configuration | - | 配置管理 |
| configurationsetting | - | 配置项 |
| consumeRecordDetailb | - | 消耗记录 |
| consumereRequestDetailb | - | 消耗请求 |
| containerBindRecordDetail | - | 容器绑定记录 |
| containerDetail | - | 容器明细 |
| containerInitRecordDetail | - | 容器初始化 |
| containerMain | - | 容器主数据 |
| containerRepairRecordDetail | - | 容器维修 |
| containerUnbindRecordDetail | - | 容器解绑 |
| countJobDetail | - | 盘点任务 |
| countPlanDetail | - | 盘点计划 |
| countRecordDetail | - | 盘点记录 |
| countRequestDetail | - | 盘点请求 |
| countadjustRecordDetail | - | 调整记录 |
| countadjustRequestDetail | - | 调整请求 |
| currencyexchange | - | 汇率 |
| customerreceiptRecordDetail | - | 客户收款 |
| customerreceiptRequestDetail | - | 客户收款请求 |
| customerreturnJobDetail | - | 客户退货任务 |
| customerreturnRecordDetail | - | 客户退货记录 |
| customerreturnRequestDetail | - | 客户退货请求 |
| customersettleRecordDetail | - | 客户结算 |
| customersettleRequestDetail | - | 客户结算请求 |
| deliverJobDetail | - | 发货任务 |
| deliverRequestDetail | - | 发货请求 |
| detail | - | 详情查询 |
| dismantleRecordDetailb | - | 拆解记录 |
| dismantleRequestDetailb | - | 拆解请求 |
| dock | - | Dock管理 |
| documentsetting | - | 单据设置 |
| enterprise | - | 企业管理 |
| file | - | 文件管理 |
| inspectJobDetail | - | 检验任务 |
| inspectRecordDetail | - | 检验记录 |
| inspectRequestDetail | - | 检验请求 |
| inventorychangeRecordDetail | - | 库存变化记录 |
| inventorychangeRequestDetail | - | 库存变化请求 |
| inventoryinitRecordMain | - | 库存初始化记录 |
| inventoryinitRequestMain | - | 库存初始化请求 |
| inventorymoveJobDetail | - | 移库任务 |
| issueRecordDetail | - | 领料记录 |
| issueRequestDetail | - | 领料请求 |
| jobsetting | - | 作业设置 |
| labeltype | - | 标签类型 |
| offlinesettlementRecordDetail | - | 离线结算记录 |
| offlinesettlementRequestDetail | - | 离线结算请求 |
| putawayJobDetail | - | 上架任务 |
| purchasereturnJobDetail | - | 采购退货任务 |
| purchasereturnRecordDetail | - | 采购退货记录 |
| purchasereturnRequestDetail | - | 采购退货请求 |
| purchasesettlementRecordDetail | - | 采购结算记录 |
| purchasesettlementRequestDetail | - | 采购结算请求 |
| qualityControlDetail | - | 质检明细 |
| qualityControlRecordDetail | - | 质检记录 |
| qualityControlRequestDetail | - | 质检请求 |
| receiveRecordDetail | - | 收货记录 |
| receiveRequestDetail | - | 收货请求 |
| reportDetail | - | 报表明细 |
| returnSettlementDetail | - | 退货结算 |
| stationDetail | - | 工位明细 |
| supplierdeliverinspectiondetail | - | 供应商交货检验 |
| supplyLinkDetail | - | 供应链明细 |

### 3.3 BPM模块API

| API文件 | 功能说明 |
|--------|---------|
| activity | 活动节点管理 |
| definition | 流程定义 |
| form | 表单管理 |
| leave | 请假管理 |
| model | 流程模型 |
| processInstance | 流程实例 |
| task | 任务管理 |
| taskAssignRule | 分配规则 |
| userGroup | 用户组 |

### 3.4 SYSTEM模块API

| API文件 | 功能说明 |
|--------|---------|
| area | 区域管理 |
| dept | 部门管理 |
| errorCode | 错误码 |
| loginLog | 登录日志 |
| mail/account | 邮件账号 |
| menu | 菜单管理 |
| notify/message | 消息通知 |
| oauth2/client | OAuth2客户端 |
| oauth2/token | Token管理 |
| operatelog | 操作日志 |
| permission | 权限管理 |
| post | 岗位管理 |
| sensitiveWord | 敏感词 |
| serialNumber | 序列号 |
| sms/smsChannel | 短信渠道 |
| sms/smsLog | 短信日志 |
| tenantPackage | 租户套餐 |
| user | 用户管理 |

### 3.5 INFRA模块API

| API文件 | 功能说明 |
|--------|---------|
| apiAccessLog | API访问日志 |
| apiErrorLog | API错误日志 |
| codegen | 代码生成 |
| dataSourceConfig | 数据源配置 |
| dbDoc | 数据库文档 |
| file | 文件管理 |
| fileConfig | 文件配置 |
| job | 定时任务 |
| jobLog | 任务日志 |
| redis | Redis管理 |

---

## 4. 前后端接口对应关系

### 4.1 MES模块

| 前端页面 | 前端API | 后端Controller |
|---------|--------|---------------|
| 生产日计划 | abilityInfo | MesAbilityInfoController |
| 工序管理 | opersteps | MesOperstepsController |
| 工序类型 | operstepsType | MesOperstepsTypeController |

### 4.2 WMS模块

| 前端子模块 | 前端API路径 | 后端Controller包 |
|----------|-----------|----------------|
| 基础数据-工厂建模 | basicDataManage/factoryModeling/* | wms/areabasic, wms/warehouse, wms/location等 |
| 基础数据-物料管理 | basicDataManage/itemManage/* | wms/itembasic, wms/itempackage等 |
| 基础数据-客户管理 | basicDataManage/customerManage/* | wms/customer等 |
| 基础数据-供应商 | basicDataManage/supplierManage/* | wms/supplier等 |
| 基础数据-标签 | basicDataManage/labelManage/* | wms/labeltype, wms/barbasic等 |
| 基础数据-策略 | basicDataManage/strategySetting/* | wms/strategy, wms/rule等 |
| 库存-盘点 | countManage/* | wms/countPlan, wms/countJob等 |
| 入库-采购 | purchasereceiptManage/* | wms/purchasereceiptJob |
| AGV | agvManage/* | wms/agvService |

### 4.3 SYSTEM模块

| 前端页面 | 前端API | 后端Controller |
|---------|--------|---------------|
| 用户管理 | system/user | SystemUserController |
| 角色管理 | system/role | SystemRoleController |
| 菜单管理 | system/menu | SystemMenuController |
| 部门管理 | system/dept | SystemDeptController |
| 字典管理 | system/dict | SystemDictTypeController |
| 租户管理 | system/tenant | SystemTenantController |
| 短信管理 | system/sms/* | SystemSmsChannelController |
| 邮件管理 | system/mail/* | SystemMailAccountController |
| 消息管理 | system/notify/* | SystemNotifyMessageController |

---

## 5. 前端功能清单汇总

### 5.1 按模块统计

| 模块 | 页面数 | API接口数 | 说明 |
|-----|-------|----------|------|
| MES | 30+ | 4+ | 制造执行核心模块 |
| WMS | 300+ | 80+ | 仓库管理核心模块 |
| QMS | 50+ | 10+ | 质量检验模块 |
| EAM | 65+ | 20+ | 设备管理模块 |
| BPM | 30+ | 9 | 流程管理模块 |
| SYSTEM | 35+ | 25+ | 系统管理模块 |
| INFRA | 12+ | 10+ | 基础设施模块 |
| REPORT | 1+ | 1 | 报表模块 |

**前端页面总计: 450+ 个**

### 5.2 核心功能清单

#### MES制造执行
- [x] 生产日计划 CRUD、发布、终止、恢复
- [x] 工单排程与工序调度
- [x] 工艺路线定义与工序管理
- [x] BOM物料清单配置
- [x] 设备产线工位配置
- [x] 班组人员与能力档案
- [x] 工序报工与报交
- [x] 产品拆解与返工
- [x] 月度生产计划
- [x] 质量表单与检验

#### WMS仓库管理
- [x] 仓库/库区/库位/库位组 基础档案
- [x] 物料/产品/容器 档案管理
- [x] 客户/供应商/承运商 档案
- [x] 标签模板与条码规则
- [x] 业务策略规则配置
- [x] 采购入库/生产入库/退货入库
- [x] 销售发货/领料/拣货
- [x] 库存盘点与调整
- [x] AGV任务调度
- [x] 客户结算与付款

#### QMS质量管理
- [x] AQL抽样标准配置
- [x] 检验方案与模板
- [x] 检验方法与阶段
- [x] 定性/定量检验特性
- [x] 抽样方案与代码
- [x] IQC/PQC/OQC 检验流程
- [x] 质量通知单

#### EAM设备管理
- [x] 设备台账与分类
- [x] 设备点检计划与记录
- [x] 设备保养计划与记录
- [x] 设备维修与报修
- [x] 设备停机与转移
- [x] 备件库存与使用
- [x] 工具资产

#### BPM流程管理
- [x] BPMN流程设计器
- [x] 流程定义与部署
- [x] 流程实例管理
- [x] 待办/已办任务
- [x] 任务转派与审批
- [x] 自定义表单
- [x] OA请假示例

#### SYSTEM系统管理
- [x] 用户/角色/菜单/权限
- [x] 部门与岗位管理
- [x] 多租户管理
- [x] 数据字典
- [x] 短信/邮件/站内消息
- [x] OAuth2第三方登录
- [x] 登录日志与操作日志
- [x] 敏感词与错误码

#### INFRA基础设施
- [x] 文件上传与下载
- [x] CRUD代码生成器
- [x] 多数据源配置
- [x] 定时任务调度
- [x] Redis缓存管理
- [x] API访问日志

#### REPORT报表
- [x] GoView数据大屏

---

## 6. 前端路由配置

### 6.1 路由守卫

```typescript
// 路由鉴权流程
1. 白名单校验 (/login, /register)
2. Token有效性校验
3. 权限码校验 (后端返回)
4. 动态路由生成
```

### 6.2 动态路由示例

```typescript
// 路由模块划分
- 静态路由: /login, /error, /redirect
- 动态路由: /mes/*, /wms/*, /qms/*, /eam/*, /bpm/*, /system/*, /infra/*, /report/*
```

---

## 7. 前端组件规范

### 7.1 组件命名

| 类型 | 命名规范 | 示例 |
|-----|---------|------|
| 页面组件 | XxxPage.vue | OrderDayPage.vue |
| 业务组件 | Xxx/index.vue | Detail/index.vue |
| 表单组件 | XxxForm.vue | UserForm.vue |
| 通用组件 | ElXxx (扩展) | ElTableColumn |

### 7.2 目录结构

```
views/
  ├── module/                   # 按业务模块划分
  │   ├── page/                # 页面
  │   │   ├── index.vue        # 列表页
  │   │   ├── Detail.vue        # 详情页
  │   │   └── Form.vue          # 表单页
  │   └── components/           # 业务组件
  └── common/                   # 公共页面
      ├── login/
      └── error/
```

---

*文档版本: v1.0 | 最后更新: 2026-04-17*
