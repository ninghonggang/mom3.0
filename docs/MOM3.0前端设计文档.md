# MOM3.0 前端设计文档

> **版本**: V1.1
> **日期**: 2026-04-17
> **项目**: 闻荫科技MOM3.0智能制造执行系统
> **前端框架**: Vue3 + Element Plus + TypeScript

---

## 1. 前端技术架构

### 1.1 技术栈

| 技术 | 版本 | 说明 |
|-----|-----|------|
| Vue | 3.3.4 | 渐进式前端框架 |
| Vite | 4.4.9 | 新一代前端构建工具 |
| TypeScript | 5.2.2 | JavaScript超集，强类型 |
| Element Plus | 2.3.14 | Vue3 UI组件库 |
| Pinia | 2.1.6 | Vue3状态管理 |
| Vue Router | 4.2.5 | Vue3路由管理 |
| Axios | 1.5.0 | HTTP客户端 |
| ECharts | 5.4.3 | 图表库 |

### 1.2 项目结构

```
mom-web/src/
├── api/                    # API接口层
│   ├── ai-chat.ts         # AI聊天
│   ├── alert.ts           # 安灯告警
│   ├── aps.ts             # APS计划
│   ├── auth.ts            # 认证
│   ├── bpm.ts             # BPM流程
│   ├── equipment.ts       # 设备点检
│   ├── event.ts           # 事件
│   ├── mdm.ts             # 主数据
│   ├── mes.ts             # MES执行
│   ├── production.ts       # 生产执行
│   ├── production_issue.ts # 生产发料
│   ├── quality.ts         # 质量管理
│   ├── scp.ts             # 供应链
│   ├── system.ts          # 系统管理
│   ├── trace.ts           # 追溯
│   └── wms.ts            # 仓储
├── views/                  # 页面视图层
│   ├── agv/              # AGV调度 (3页)
│   ├── alert/            # 安灯告警 (5页)
│   ├── aps/              # APS计划 (10页)
│   ├── bpm/              # BPM流程 (3页)
│   ├── eam/              # 设备管理 (3页)
│   ├── energy/           # 能源管理 (1页)
│   ├── equipment/        # 设备点检 (11页)
│   ├── fin/              # 结算 (3页)
│   ├── integration/      # 系统集成 (2页)
│   ├── mdm/              # 主数据 (15页)
│   ├── mes/              # MES执行 (5页)
│   ├── production/       # 生产执行 (13页)
│   ├── quality/          # 质量管理 (16页)
│   ├── report/           # 报表 (5页)
│   ├── scp/              # 供应链 (6页)
│   ├── supplier/         # 供应商 (1页)
│   ├── system/           # 系统管理 (13页)
│   ├── trace/            # 追溯 (2页)
│   ├── wms/              # 仓储 (9页)
│   ├── Dashboard.vue      # 首页看板
│   ├── Login.vue         # 登录页
│   └── Error404.vue      # 错误页
├── router/                 # 路由配置
├── store/                  # 状态管理
├── components/            # 公共组件
├── utils/                 # 工具函数
│   └── request.ts        # Axios封装
└── styles/                # 样式文件
```

### 1.3 状态管理 (Pinia Store)

| Store模块 | 职责 | 主要状态 |
|----------|------|---------|
| user | 用户信息、Token、权限 | token, userInfo, roles |
| permission | 菜单权限、路由权限 | routes, permissions |
| dict | 数据字典缓存 | dictData |
| tagsView | 标签页管理 | visitedViews |
| app | 应用配置 | sidebar, device |

### 1.4 Axios请求封装

```typescript
// utils/request.ts
import axios from 'axios'
import { ElMessage } from 'element-plus'

const service = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: 30000
})

// 请求拦截器 - 添加Token
service.interceptors.request.use(
  config => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`
    }
    return config
  }
)

// 响应拦截器 - 统一错误处理
service.interceptors.response.use(
  response => response.data,
  error => {
    ElMessage.error(error.response?.data?.message || '请求失败')
    return Promise.reject(error)
  }
)

export default service
```

---

## 2. API接口详细清单

### 2.1 System API (`/api/system.ts`)

**用户管理** `/system/user`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /system/user/list | 获取用户列表 |
| GET | /system/user/{id} | 获取用户详情 |
| POST | /system/user | 创建用户 |
| PUT | /system/user/{id} | 更新用户 |
| DELETE | /system/user/{id} | 删除用户 |
| PUT | /system/user/{id}/password | 重置密码 |
| PUT | /system/user/{id}/roles | 分配角色 |

**角色管理** `/system/role`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /system/role/list | 获取角色列表 |
| GET | /system/role/{id} | 获取角色详情 |
| POST | /system/role | 创建角色 |
| PUT | /system/role/{id} | 更新角色 |
| DELETE | /system/role/{id} | 删除角色 |
| GET | /system/role/{id}/menus | 获取角色菜单 |
| PUT | /system/role/{id}/menus | 分配菜单 |
| GET | /system/role/{id}/perms | 获取角色权限 |
| PUT | /system/role/{id}/perms | 分配权限 |

**菜单管理** `/system/menu`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /system/menu/list | 获取菜单列表 |
| GET | /system/menu/tree | 获取菜单树 |
| GET | /system/menu/{id} | 获取菜单详情 |
| POST | /system/menu | 创建菜单 |
| PUT | /system/menu/{id} | 更新菜单 |
| DELETE | /system/menu/{id} | 删除菜单 |

**部门管理** `/system/dept`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /system/dept/list | 获取部门列表 |
| GET | /system/dept/tree | 获取部门树 |
| GET | /system/dept/{id} | 获取部门详情 |
| POST | /system/dept | 创建部门 |
| PUT | /system/dept/{id} | 更新部门 |
| DELETE | /system/dept/{id} | 删除部门 |

**字典管理** `/system/dict`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /system/dict/type/list | 获取字典类型列表 |
| GET | /system/dict/{type}/data | 获取字典数据 |
| POST | /system/dict/type | 创建字典类型 |
| PUT | /system/dict/type/{id} | 更新字典类型 |
| DELETE | /system/dict/type/{id} | 删除字典类型 |

**其他系统接口**

| 模块 | 接口 | 说明 |
|------|------|------|
| 岗位 | /system/post/list, POST, PUT, DELETE | 岗位CRUD |
| 租户 | /system/tenant/list, {id}, POST, PUT, DELETE | 租户CRUD |
| 登录日志 | /system/loginlog/list, DELETE(clean), GET(export) | 登录日志 |
| 操作日志 | /system/operlog/list, DELETE(clean), GET(export) | 操作日志 |
| 通知公告 | /system/notice/list, {id}, POST, PUT, DELETE, PUT/{id}/publish | 公告管理 |
| 打印模板 | /system/print-template/list, {id}, POST, PUT, DELETE | 打印模板 |

---

### 2.2 MDM API (`/api/mdm.ts`)

**物料管理** `/mdm/material`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mdm/material/list | 获取物料列表 |
| GET | /mdm/material/{id} | 获取物料详情 |
| POST | /mdm/material | 创建物料 |
| PUT | /mdm/material/{id} | 更新物料 |
| DELETE | /mdm/material/{id} | 删除物料 |

**BOM管理** `/mdm/bom`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mdm/bom/list | 获取BOM列表 |
| GET | /mdm/bom/{productId}/tree | 获取BOM树 |
| POST | /mdm/bom | 创建BOM |
| PUT | /mdm/bom/{id} | 更新BOM |
| DELETE | /mdm/bom/{id} | 删除BOM |

**工艺路线** `/mdm/process`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mdm/process/list | 获取工艺路线列表 |
| POST | /mdm/process | 创建工艺路线 |
| PUT | /mdm/process/{id} | 更新工艺路线 |
| DELETE | /mdm/process/{id} | 删除工艺路线 |

**车间** `/mdm/workshop`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mdm/workshop/list | 获取车间列表 |
| POST | /mdm/workshop | 创建车间 |
| PUT | /mdm/workshop/{id} | 更新车间 |
| DELETE | /mdm/workshop/{id} | 删除车间 |

**生产线** `/mdm/line`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mdm/line/list | 获取生产线列表 |
| POST | /mdm/line | 创建生产线 |
| PUT | /mdm/line/{id} | 更新生产线 |
| DELETE | /mdm/line/{id} | 删除生产线 |

**工位** `/mdm/workstation`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mdm/workstation/list | 获取工位列表 |
| POST | /mdm/workstation | 创建工位 |
| PUT | /mdm/workstation/{id} | 更新工位 |
| DELETE | /mdm/workstation/{id} | 删除工位 |

**客户** `/mdm/customer`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mdm/customer/list | 获取客户列表 |
| POST | /mdm/customer | 创建客户 |
| PUT | /mdm/customer/{id} | 更新客户 |
| DELETE | /mdm/customer/{id} | 删除客户 |

**供应商** `/mdm/supplier`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /mdm/supplier/list | 获取供应商列表 |
| POST | /mdm/supplier | 创建供应商 |
| PUT | /mdm/supplier/{id} | 更新供应商 |
| DELETE | /mdm/supplier/{id} | 删除供应商 |

---

### 2.3 Production API (`/api/production.ts`)

**生产工单** `/production/order`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /production/order/list | 获取工单列表 |
| GET | /production/order/{id} | 获取工单详情 |
| POST | /production/order | 创建工单 |
| PUT | /production/order/{id} | 更新工单 |
| DELETE | /production/order/{id} | 删除工单 |
| PUT | /production/order/{id}/start | 开始生产 |
| PUT | /production/order/{id}/complete | 完工 |
| PUT | /production/order/{id}/cancel | 取消工单 |

**派工管理** `/production/dispatch`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /production/dispatch/list | 获取派工列表 |
| POST | /production/dispatch | 创建派工 |
| PUT | /production/dispatch/{id} | 更新派工 |
| PUT | /production/dispatch/{id}/start | 开始作业 |
| PUT | /production/dispatch/{id}/complete | 完工 |

**生产报工** `/production/report`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /production/report/list | 获取报工列表 |
| POST | /production/report | 创建报工 |

**销售订单** `/production/sales-order`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /production/sales-order/list | 获取销售订单列表 |
| POST | /production/sales-order | 创建销售订单 |
| PUT | /production/sales-order/{id} | 更新销售订单 |
| DELETE | /production/sales-order/{id} | 删除销售订单 |
| PUT | /production/sales-order/{id}/confirm | 确认订单 |

**首末件检验** `/production/first-last-inspect`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /production/first-last-inspect/list | 获取首末件列表 |
| GET | /production/first-last-inspect/{id} | 获取详情 |
| POST | /production/first-last-inspect | 创建记录 |
| PUT | /production/first-last-inspect/{id} | 更新记录 |
| DELETE | /production/first-last-inspect/{id} | 删除记录 |
| GET | /production/first-last-inspect/overdue | 获取逾期列表 |

**电子SOP** `/production/electronic-sop`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /production/electronic-sop/list | 获取SOP列表 |
| GET | /production/electronic-sop/{id} | 获取SOP详情 |
| POST | /production/electronic-sop | 创建SOP |
| PUT | /production/electronic-sop/{id} | 更新SOP |
| DELETE | /production/electronic-sop/{id} | 删除SOP |

**编码规则** `/production/code-rule`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /production/code-rule/list | 获取规则列表 |
| GET | /production/code-rule/{id} | 获取规则详情 |
| POST | /production/code-rule | 创建规则 |
| PUT | /production/code-rule/{id} | 更新规则 |
| DELETE | /production/code-rule/{id} | 删除规则 |
| GET | /production/code-rule/generate | 生成编码 |

**工序流转卡** `/production/flow-card`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /production/flow-card/list | 获取流转卡列表 |
| GET | /production/flow-card/{id} | 获取流转卡详情 |
| POST | /production/flow-card | 创建流转卡 |
| PUT | /production/flow-card/{id} | 更新流转卡 |
| DELETE | /production/flow-card/{id} | 删除流转卡 |

**包装条码** `/production/packages`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /production/packages/list | 获取包装列表 |
| GET | /production/packages/{id} | 获取包装详情 |
| POST | /production/packages/create | 创建包装 |
| POST | /production/packages/add-item | 添加包装项 |
| POST | /production/packages/seal | 封箱 |
| DELETE | /production/packages/{id} | 删除包装 |

---

### 2.4 Quality API (`/api/quality.ts`)

**IQC来料检验** `/quality/iqc`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /quality/iqc/list | 获取IQC列表 |
| POST | /quality/iqc | 创建IQC记录 |
| PUT | /quality/iqc/{id} | 更新IQC |
| DELETE | /quality/iqc/{id} | 删除IQC |

**IPQC过程检验** `/quality/ipqc`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /quality/ipqc/list | 获取IPQC列表 |
| POST | /quality/ipqc | 创建IPQC记录 |
| PUT | /quality/ipqc/{id} | 更新IPQC |
| DELETE | /quality/ipqc/{id} | 删除IPQC |

**FQC出货检验** `/quality/fqc`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /quality/fqc/list | 获取FQC列表 |
| POST | /quality/fqc | 创建FQC记录 |
| PUT | /quality/fqc/{id} | 更新FQC |
| DELETE | /quality/fqc/{id} | 删除FQC |

**OQC出货检验** `/quality/oqc`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /quality/oqc/list | 获取OQC列表 |
| POST | /quality/oqc | 创建OQC记录 |
| PUT | /quality/oqc/{id} | 更新OQC |
| DELETE | /quality/oqc/{id} | 删除OQC |

**不良品记录** `/quality/defect`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /quality/defect/list | 获取不良品列表 |
| POST | /quality/defect | 创建不良品记录 |
| PUT | /quality/defect/{id} | 更新不良品 |
| PUT | /quality/defect/{id}/handle | 处理不良品 |
| DELETE | /quality/defect/{id} | 删除记录 |

**NCR不合格品** `/quality/ncr`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /quality/ncr/list | 获取NCR列表 |
| POST | /quality/ncr | 创建NCR |
| PUT | /quality/ncr/{id} | 更新NCR |
| DELETE | /quality/ncr/{id} | 删除NCR |

**SPC统计过程控制** `/quality/spc`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /quality/spc/list | 获取SPC数据 |
| GET | /quality/spc/chart | 获取SPC图表 |

**QRCI质量闭环** `/quality/qrci`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /quality/qrci/list | 获取QRCI列表 |
| GET | /quality/qrci/{id} | 获取QRCI详情 |
| POST | /quality/qrci | 创建QRCI |
| PUT | /quality/qrci/{id} | 更新QRCI |
| PUT | /quality/qrci/{id}/close | 关闭QRCI |
| DELETE | /quality/qrci/{id} | 删除QRCI |
| GET | /quality/qrci/{id}/5why | 获取5Why分析 |
| POST | /quality/qrci/{id}/5why | 添加5Why分析 |
| GET | /quality/qrci/{id}/actions | 获取改善行动 |
| POST | /quality/qrci/{id}/actions | 添加改善行动 |
| PUT | /quality/qrci/{id}/actions/{actionId} | 更新行动 |
| POST | /quality/qrci/{id}/verification | 添加验证 |

**LPA分层审核** `/quality/lpa`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /quality/lpa/standards/list | 获取LPA标准列表 |
| GET | /quality/lpa/standards/{id} | 获取标准详情 |
| POST | /quality/lpa/standards | 创建LPA标准 |
| PUT | /quality/lpa/standards/{id} | 更新LPA标准 |
| DELETE | /quality/lpa/standards/{id} | 删除LPA标准 |
| GET | /quality/lpa/standards/{id}/questions | 获取问题项 |
| POST | /quality/lpa/standards/{id}/questions | 添加问题项 |
| GET | /quality/lpa/records/list | 获取LPA审核记录 |
| GET | /quality/lpa/records/{id} | 获取审核记录详情 |
| POST | /quality/lpa/records | 创建审核记录 |
| PUT | /quality/lpa/records/{id}/verify | 审核验证 |

**动态规则** `/quality/dynamic-rule`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /quality/dynamic-rule/list | 获取规则列表 |
| GET | /quality/dynamic-rule/{id} | 获取规则详情 |
| POST | /quality/dynamic-rule | 创建规则 |
| PUT | /quality/dynamic-rule/{id} | 更新规则 |
| DELETE | /quality/dynamic-rule/{id} | 删除规则 |
| POST | /quality/dynamic-rule/evaluate | 评估规则 |
| GET | /quality/dynamic-rule/logs | 获取规则日志 |

---

### 2.5 Equipment API (`/api/equipment.ts`)

**设备台账** `/equipment`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /equipment/list | 获取设备列表 |
| GET | /equipment/{id} | 获取设备详情 |
| POST | /equipment | 创建设备 |
| PUT | /equipment/{id} | 更新设备 |
| DELETE | /equipment/{id} | 删除设备 |
| GET | /equipment/status | 获取设备状态 |

**设备点检** `/equipment/check`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /equipment/check/list | 获取点检列表 |
| POST | /equipment/check | 创建点检 |
| PUT | /equipment/check/{id} | 更新点检 |
| DELETE | /equipment/check/{id} | 删除点检 |

**设备保养** `/equipment/maintenance`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /equipment/maintenance/list | 获取保养列表 |
| POST | /equipment/maintenance | 创建保养 |
| PUT | /equipment/maintenance/{id} | 更新保养 |
| DELETE | /equipment/maintenance/{id} | 删除保养 |

**设备维修** `/equipment/repair`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /equipment/repair/list | 获取维修列表 |
| POST | /equipment/repair | 创建维修 |
| PUT | /equipment/repair/{id}/start | 开始维修 |
| PUT | /equipment/repair/{id}/complete | 完工 |
| PUT | /equipment/repair/{id} | 更新维修 |
| DELETE | /equipment/repair/{id} | 删除维修 |

**OEE设备综合效率** `/equipment/oee`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /equipment/oee/list | 获取OEE列表 |
| GET | /equipment/oee/{id} | 获取OEE详情 |
| POST | /equipment/oee/calculate | 计算OEE |
| GET | /equipment/oee/chart | 获取OEE图表 |

---

### 2.6 WMS API (`/api/wms.ts`)

**仓库管理** `/wms/warehouse`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /wms/warehouse/list | 获取仓库列表 |
| POST | /wms/warehouse | 创建仓库 |
| PUT | /wms/warehouse/{id} | 更新仓库 |
| DELETE | /wms/warehouse/{id} | 删除仓库 |

**库位管理** `/wms/location`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /wms/location/list | 获取库位列表 |
| POST | /wms/location | 创建库位 |
| PUT | /wms/location/{id} | 更新库位 |
| DELETE | /wms/location/{id} | 删除库位 |

**库存管理** `/wms/inventory`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /wms/inventory/list | 获取库存列表 |
| GET | /wms/inventory/material/{materialId} | 按物料查库存 |
| POST | /wms/inventory/adjust | 调整库存 |
| DELETE | /wms/inventory/{id} | 删除库存记录 |

**收货单** `/wms/receive`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /wms/receive/list | 获取收货列表 |
| GET | /wms/receive/{id} | 获取收货单详情 |
| POST | /wms/receive | 创建收货单 |
| PUT | /wms/receive/{id} | 更新收货单 |
| DELETE | /wms/receive/{id} | 删除收货单 |
| PUT | /wms/receive/{id}/confirm | 确认收货 |

**发货单** `/wms/delivery`

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /wms/delivery/list | 获取发货列表 |
| GET | /wms/delivery/{id} | 获取发货单详情 |
| POST | /wms/delivery | 创建发货单 |
| PUT | /wms/delivery/{id} | 更新发货单 |
| DELETE | /wms/delivery/{id} | 删除发货单 |
| PUT | /wms/delivery/{id}/confirm | 确认发货 |

---

### 2.7 APS API (`/api/aps.ts`)

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /aps/mps/list | MPS主计划列表 |
| GET | /aps/mrp/list | MRP物料需求列表 |
| GET | /aps/schedule/list | 排程计划列表 |
| GET | /aps/delivery-analysis | 交付分析 |
| GET | /aps/material-shortage | 物料短缺分析 |
| GET | /aps/work-center/list | 工作中心列表 |
| GET | /aps/product-family/list | 产品族列表 |

---

### 2.8 SCP API (`/api/scp.ts`)

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /scp/purchase/list | 采购订单列表 |
| GET | /scp/purchase/{id} | 采购订单详情 |
| POST | /scp/purchase | 创建采购订单 |
| PUT | /scp/purchase/{id} | 更新采购订单 |
| DELETE | /scp/purchase/{id} | 删除采购订单 |
| GET | /scp/rfq/list | RFQ询价列表 |
| GET | /scp/supplier-quote/list | 供应商报价列表 |
| GET | /scp/supplier-kpi/list | 供应商KPI列表 |
| GET | /scp/customer-inquiry/list | 客户询价列表 |

---

### 2.9 Alert API (`/api/alert.ts`)

| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /alert/rule/list | 告警规则列表 |
| POST | /alert/rule | 创建告警规则 |
| PUT | /alert/rule/{id} | 更新告警规则 |
| DELETE | /alert/rule/{id} | 删除告警规则 |
| GET | /alert/record/list | 告警记录列表 |
| PUT | /alert/record/{id}/ack | 确认告警 |
| GET | /alert/statistics | 告警统计 |

---

## 3. 前端页面清单

### 3.1 首页模块

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| Dashboard | /dashboard | Dashboard.vue | 首页看板,展示核心生产数据 |
| Error404 | /error/404 | Error404.vue | 404错误页面 |

### 3.2 系统登录

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| Login | /login | Login.vue | 用户登录页面 |

### 3.3 M01 系统管理 (`/system/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| 用户管理 | /system/user | UserList.vue | 用户CRUD、导入导出、角色分配 |
| 角色管理 | /system/role | RoleList.vue | 角色权限配置、菜单分配 |
| 菜单管理 | /system/menu | MenuList.vue | 菜单路由配置、图标设置 |
| 部门管理 | /system/dept | DeptList.vue | 组织架构管理、部门树 |
| 岗位管理 | /system/post | PostList.vue | 岗位配置 |
| 数据字典 | /system/dict | DictList.vue | 字典类型管理 |
| 租户管理 | /system/tenant | TenantList.vue | 多租户配置 |
| 登录日志 | /system/loginlog | LoginLogList.vue | 用户登录记录 |
| 操作日志 | /system/operlog | OperLogList.vue | 操作审计日志 |
| 通知公告 | /system/notice | NoticeList.vue | 系统公告发布 |
| AI配置 | /system/ai-config | AiConfigView.vue | AI功能配置 |
| 打印模板 | /system/print-template | PrintTemplateList.vue | 打印模板管理 |
| 系统配置 | /system/config | SystemConfig.vue | 系统参数配置 |

### 3.4 M02 主数据管理 (`/mdm/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| 物料管理 | /mdm/material | MaterialList.vue | 物料主数据CRUD、分类 |
| BOM管理 | /mdm/bom | BomList.vue | 物料清单配置 |
| BOM编辑 | /mdm/bom-editor | BomItemEditor.vue | BOM明细编辑 |
| 工艺路线 | /mdm/operation | OperationList.vue | 工序定义、工时配置 |
| 客户管理 | /mdm/customer | CustomerList.vue | 客户档案、联系人 |
| 供应商管理 | /mdm/supplier | SupplierList.vue | 供应商档案、联系人 |
| 银行账户 | /mdm/bank-account | BankAccountList.vue | 银行账户管理 |
| 联系人 | /mdm/contact | ContactList.vue | 联系人管理 |
| 交货地址 | /mdm/delivery-address | DeliveryAddressList.vue | 交货地址配置 |
| 物料分类 | /mdm/material-category | MaterialCategoryList.vue | 物料分类管理 |
| 产线管理 | /mdm/line | LineList.vue | 产线配置 |
| 车间管理 | /mdm/workshop | WorkshopList.vue | 车间配置 |
| 工位管理 | /mdm/workstation | WorkstationList.vue | 工位配置 |
| 班次管理 | /mdm/shift | ShiftList.vue | 工作班次定义 |
| 附件管理 | /mdm/attachment | AttachmentList.vue | 附件上传管理 |

### 3.5 M03 生产执行 (`/production/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| 销售订单 | /production/sales-order | SalesOrderList.vue | 销售订单列表、确认 |
| 生产工单 | /production/order | ProductionOrderList.vue | 工单CRUD、状态管理 |
| 派工管理 | /production/dispatch | DispatchList.vue | 派工单、报工管理 |
| 生产报工 | /production/report | ReportList.vue | 工序报工、完工入库 |
| 生产看板 | /production/kanban | KanbanBoard.vue | 车间看板可视化 |
| 首末件检验 | /production/first-last-inspect | FirstLastInspectList.vue | 首件/末件检验 |
| 工序流转卡 | /production/flow-card | FlowCardList.vue | 工序流转卡管理 |
| 生产变更 | /production/order-change | OrderChangeList.vue | 工单变更记录 |
| 生产发料 | /production/issue | ProductionIssueList.vue | 生产发料管理 |
| 生产退料 | /production/return | ProductionReturnList.vue | 生产退料管理 |
| 包装管理 | /production/package | PackageList.vue | 产品包装管理 |
| 电子SOP | /production/electronic-sop | ElectronicSOPList.vue | 电子作业指导书 |
| 编码规则 | /production/code-rule | CodeRuleList.vue | 产品序列号编码规则 |

### 3.6 M04 APS计划 (`/aps/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| MPS主计划 | /aps/mps | MPSList.vue | 主生产计划管理 |
| MRP物料需求 | /aps/mrp | MRPList.vue | 物料需求计划 |
| 排程计划 | /aps/schedule | ScheduleList.vue | 甘特图排程 |
| 交付分析 | /aps/delivery-analysis | DeliveryAnalysisList.vue | 交付能力分析 |
| 物料短缺 | /aps/material-shortage | MaterialShortageList.vue | 物料短缺分析 |
| 短缺规则 | /aps/shortage-rule | ShortageRuleList.vue | 短缺处理规则 |
| 工作中心 | /aps/work-center | WorkCenterList.vue | 工作中心配置 |
| 产品族 | /aps/product-family | ProductFamilyList.vue | 产品族管理 |
| 滚动配置 | /aps/rolling-config | RollingConfigList.vue | 滚动计划配置 |
| 换型矩阵 | /aps/changeover-matrix | ChangeoverMatrixList.vue | 换型时间矩阵 |

### 3.7 M05 质量管理 (`/quality/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| IQC来料检验 | /quality/iqc | IQCList.vue | 来料质量控制 |
| IPQC过程检验 | /quality/ipqc | IPQCList.vue | 制程质量控制 |
| FQC出货检验 | /quality/fqc | FQCList.vue | 出货质量控制 |
| OQC出货检验 | /quality/oqc | OQCList.vue | 出货质量检验 |
| 检验计划 | /quality/inspection-plan | InspectionPlanList.vue | 检验方案配置 |
| 检验记录 | /quality/inspection-record | InspectionRecordList.vue | 检验结果记录 |
| 检验模板 | /quality/inspection-template | InspectionTemplateList.vue | 检验模板管理 |
| AQL配置 | /quality/aql | AQLList.vue | AQL抽样标准 |
| 动态规则 | /quality/dynamic-rule | DynamicRuleList.vue | 动态检验规则 |
| 缺陷代码 | /quality/defect-code | DefectCodeList.vue | 缺陷代码管理 |
| 缺陷记录 | /quality/defect-record | DefectRecordList.vue | 缺陷记录管理 |
| NCR处理 | /quality/ncr | NCRList.vue | 不合格品处理 |
| SPC数据 | /quality/spc-data | SPCDataList.vue | 统计过程控制 |
| LPA标准 | /quality/lpa-standard | LPAStandardList.vue | LPA审核标准 |
| QRCI | /quality/qrci | QRCIList.vue | 质量持续改进 |

#### 实验室 (`/quality/lab/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| 实验室样本 | /quality/lab/sample | LabSampleList.vue | 检测样本管理 |
| 实验室报告 | /quality/lab/report | LabReportList.vue | 检测报告管理 |
| 实验室仪器 | /quality/lab/instrument | LabInstrumentList.vue | 仪器校准管理 |

### 3.8 M06 设备管理 (`/eam/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| 工厂建模 | /eam/factory | FactoryList.vue | 工厂建模基础数据 |
| 设备组织 | /eam/equipment-org | EquipmentOrgList.vue | 设备组织架构 |
| 设备停机 | /eam/downtime | DowntimeList.vue | 设备停机记录 |

### 3.9 设备点检 (`/equipment/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| 设备台账 | /equipment/list | EquipmentList.vue | 设备档案管理 |
| 点检计划 | /equipment/check | CheckList.vue | 设备点检计划 |
| 点检记录 | /equipment/check-record | CheckRecordList.vue | 点检执行记录 |
| 保养计划 | /equipment/maintenance | MaintenanceList.vue | 设备保养计划 |
| 维修管理 | /equipment/repair | RepairList.vue | 设备维修管理 |
| OEE分析 | /equipment/oee | OEELIst.vue | 设备OEE分析 |
| 设备检验 | /equipment/inspection | InspectionRecordList.vue | 设备检验记录 |
| 设备缺陷 | /equipment/defect | InspectionDefectList.vue | 设备缺陷记录 |
| 检验模板 | /equipment/template | InspectionTemplateList.vue | 检验模板管理 |
| 仪表管理 | /equipment/gauge | GaugeList.vue | 仪表量具管理 |
| 备件管理 | /equipment/spare-part | SparePartList.vue | 备件库存管理 |

### 3.10 M07 WMS仓储 (`/wms/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| 仓库管理 | /wms/warehouse | WarehouseList.vue | 仓库档案管理 |
| 库位管理 | /wms/location | LocationList.vue | 库位定义、库区 |
| 收货订单 | /wms/receive-order | ReceiveOrderList.vue | 采购收货管理 |
| 发货订单 | /wms/delivery-order | DeliveryOrderList.vue | 销售发货管理 |
| 库存台账 | /wms/inventory | InventoryList.vue | 实时库存查询 |
| 调拨单 | /wms/transfer-order | TransferOrderList.vue | 库间调拨管理 |
| 盘点管理 | /wms/stock-check | StockCheckList.vue | 库存盘点 |
| 数据点配置 | /wms/data-point | DataPointList.vue | 采集数据点配置 |
| 扫描日志 | /wms/scan-log | ScanLogList.vue | 扫码记录查询 |

### 3.11 M09 安灯系统 (`/alert/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| 告警规则 | /alert/rules | AlertRulesList.vue | 告警规则配置 |
| 告警记录 | /alert/records | AlertRecordsList.vue | 告警事件记录 |
| 告警统计 | /alert/statistics | AlertStatistics.vue | 告警统计分析 |
| 告警升级 | /alert/escalation | AlertEscalationList.vue | 告警升级配置 |
| 告警通知 | /alert/notification | AlertNotification.vue | 告警通知管理 |

### 3.12 M10 追溯管理 (`/trace/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| 追溯查询 | /trace/query | TraceQuery.vue | 序列号/批次追溯 |
| Andon呼叫 | /trace/andon-call | AndonCall.vue | Andon呼叫记录 |

### 3.13 M14 系统集成 (`/integration/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| 接口配置 | /integration/config | InterfaceConfigList.vue | 系统接口配置 |
| 执行日志 | /integration/execution-log | ExecutionLogList.vue | 接口调用日志 |

### 3.14 M15 报表 (`/report/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| 生产日报 | /report/production-daily | ProductionDailyReport.vue | 生产日报表 |
| OEE报表 | /report/oee | OEEReport.vue | 设备OEE报表 |
| 质量周报 | /report/quality-weekly | QualityWeeklyReport.vue | 质量周报 |
| 交付报表 | /report/delivery | DeliveryReport.vue | 交付率报表 |
| Andon报表 | /report/andon | AndonReport.vue | 安灯统计报表 |

### 3.15 M16 SCP供应链 (`/scp/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| 询价管理 | /scp/rfq | RFQList.vue | RFQ询价单管理 |
| 采购订单 | /scp/purchase-order | PurchaseOrderList.vue | 采购订单管理 |
| 销售订单 | /scp/sales-order | SCPSalesOrderList.vue | SCP销售订单 |
| 供应商报价 | /scp/supplier-quote | SupplierQuoteList.vue | 供应商报价管理 |
| 供应商KPI | /scp/supplier-kpi | SupplierKPIList.vue | 供应商绩效评分 |
| 客户询价 | /scp/customer-inquiry | CustomerInquiryList.vue | 客户询价管理 |

### 3.16 AGV调度 (`/agv/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| AGV设备 | /agv/device | DeviceList.vue | AGV设备管理 |
| AGV库位 | /agv/location | LocationList.vue | AGV库位映射 |
| AGV任务 | /agv/task | TaskList.vue | AGV任务调度 |

### 3.17 能源管理 (`/energy/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| 能源监控 | /energy/monitor | EnergyMonitor.vue | 能源消耗监控 |

### 3.18 结算管理 (`/fin/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| 采购结算 | /fin/purchase-settlement | PurchaseSettlementList.vue | 采购结算管理 |
| 销售结算 | /fin/sales-settlement | SalesSettlementList.vue | 销售结算管理 |
| 付款申请 | /fin/payment-request | PaymentRequestList.vue | 付款申请管理 |

### 3.19 供应商模块 (`/supplier/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| ASN管理 | /supplier/asn | ASNList.vue | ASN提前发货通知 |

### 3.20 MES执行 (`/mes/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| 工艺路线 | /mes/process-route | ProcessRouteList.vue | 工艺路线管理 |
| 班组管理 | /mes/team | TeamList.vue | 班组配置 |
| 人员技能 | /mes/person-skill | PersonSkillList.vue | 人员技能档案 |
| 物料追溯 | /mes/material-trace | MaterialTrace.vue | 物料追溯查询 |
| 下线管理 | /mes/offline | OfflineList.vue | 产品下线管理 |

### 3.21 BPM流程 (`/bpm/`)

| 页面 | 路径 | 组件 | 功能说明 |
|-----|------|------|---------|
| 流程定义 | /bpm/process | ProcessList.vue | BPMN流程定义 |
| 流程实例 | /bpm/instance | InstanceList.vue | 流程实例管理 |
| 任务列表 | /bpm/task | TaskList.vue | 待办任务列表 |

---

## 4. 页面统计

### 4.1 按模块统计

| 模块 | 目录 | 页面数 | API接口数 |
|------|------|--------|----------|
| 系统管理 | system/ | 13 | 40+ |
| 主数据MDM | mdm/ | 15 | 50+ |
| 生产执行 | production/ | 13 | 45+ |
| APS计划 | aps/ | 10 | 15+ |
| 质量管理 | quality/ | 16 | 60+ |
| 设备管理 | eam/ | 3 | 10+ |
| 设备点检 | equipment/ | 11 | 35+ |
| WMS仓储 | wms/ | 9 | 30+ |
| 安灯系统 | alert/ | 5 | 15+ |
| 追溯管理 | trace/ | 2 | 5+ |
| 系统集成 | integration/ | 2 | 5+ |
| 报表 | report/ | 5 | 10+ |
| SCP供应链 | scp/ | 6 | 20+ |
| AGV调度 | agv/ | 3 | 10+ |
| 能源管理 | energy/ | 1 | 5+ |
| 结算管理 | fin/ | 3 | 10+ |
| 供应商 | supplier/ | 1 | 5+ |
| MES执行 | mes/ | 5 | 15+ |
| BPM流程 | bpm/ | 3 | 10+ |
| 首页/登录 | / | 3 | 5+ |
| **总计** | | **129** | **350+** |

---

## 5. UI设计规范

### 5.1 页面布局结构

```vue
<template>
  <div class="page-container">
    <!-- 搜索区域 -->
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" clearable>
            <el-option label="启用" value="ENABLED" />
            <el-option label="禁用" value="DISABLED" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">
            <el-icon><Search /></el-icon>查询
          </el-button>
          <el-button @click="handleReset">
            <el-icon><Refresh /></el-icon>重置
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 工具栏 -->
    <el-card class="toolbar-card">
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
      <el-button :disabled="!selected.length" @click="handleBatchDelete">
        <el-icon><Delete /></el-icon>批量删除
      </el-button>
      <el-button @click="handleExport">
        <el-icon><Download /></el-icon>导出
      </el-button>
    </el-card>

    <!-- 数据表格 -->
    <el-card class="table-card">
      <el-table
        :data="tableData"
        v-loading="loading"
        stripe
        border
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="code" label="编码" min-width="120" />
        <el-table-column prop="name" label="名称" min-width="150" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 'ENABLED' ? 'success' : 'info'">
              {{ row.status === 'ENABLED' ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" min-width="150" show-overflow-tooltip />
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <el-pagination
        v-model:current-page="pagination.page"
        v-model:page-size="pagination.size"
        :total="pagination.total"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        class="pagination"
      />
    </el-card>
  </div>
</template>
```

### 5.2 表单对话框

```vue
<template>
  <el-dialog
    v-model="dialogVisible"
    :title="dialogTitle"
    width="600px"
    @close="handleClose"
  >
    <el-form ref="formRef" :model="form" :rules="rules" label-width="100px">
      <el-form-item label="编码" prop="code">
        <el-input v-model="form.code" placeholder="请输入编码" />
      </el-form-item>
      <el-form-item label="名称" prop="name">
        <el-input v-model="form.name" placeholder="请输入名称" />
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-radio-group v-model="form.status">
          <el-radio label="ENABLED">启用</el-radio>
          <el-radio label="DISABLED">禁用</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="备注">
        <el-input v-model="form.remark" type="textarea" rows="3" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="handleClose">取消</el-button>
      <el-button type="primary" @click="handleSubmit" :loading="submitting">
        确定
      </el-button>
    </template>
  </el-dialog>
</template>
```

### 5.3 状态色彩规范

| 状态 | 色值 | Hex | El-Tag类型 | 用途 |
|------|-----|-----|-----------|------|
| 正常/启用 | 绿色 | #10b981 | success | 正常、合格、完成 |
| 警告/待处理 | 橙色 | #f59e0b | warning | 待处理、待审批 |
| 异常/禁用 | 红色 | #ef4444 | danger | 异常、紧急、不合格 |
| 进行中 | 蓝色 | #2563eb | primary | 进行中、启用 |
| 已完成 | 灰色 | #6b7280 | info | 已完成、禁用 |

### 5.4 通用CRUD流程

```typescript
// 列表查询
const loadData = async () => {
  loading.value = true
  try {
    const params = { ...searchForm, page: pagination.page, size: pagination.size }
    const res = await getList(params)
    tableData.value = res.data
    pagination.total = res.total
  } finally {
    loading.value = false
  }
}

// 新增
const handleAdd = () => {
  dialogTitle.value = '新增'
  form.value = { status: 'ENABLED' }
  dialogVisible.value = true
}

// 编辑
const handleEdit = (row: any) => {
  dialogTitle.value = '编辑'
  form.value = { ...row }
  dialogVisible.value = true
}

// 删除
const handleDelete = async (row: any) => {
  await ElMessageBox.confirm('确定删除该数据吗?', '提示')
  await deleteById(row.id)
  ElMessage.success('删除成功')
  loadData()
}
```

---

## 6. 路由配置

### 6.1 路由守卫流程

```
1. 白名单检查 (/login, /error/*)
2. Token检查 (localStorage.getItem('token'))
3. 用户信息获取 (/system/user/info)
4. 权限路由生成 (根据用户角色)
5. 动态路由添加
```

### 6.2 路由模块映射

```
/system/*   → M01系统管理
/mdm/*      → M02主数据管理
/production/* → M03生产执行
/aps/*      → M04 APS计划
/quality/*  → M05质量管理
/equipment/* → M06设备管理
/eam/*      → M06设备管理
/wms/*      → M07 WMS仓储
/alert/*    → M09安灯系统
/trace/*    → M10追溯管理
/integration/* → M14系统集成
/report/*   → M15报表
/scp/*      → M16 SCP供应链
/agv/*      → AGV调度
/fin/*      → 结算管理
/supplier/* → 供应商管理
/mes/*      → MES执行
/bpm/*      → BPM流程
```

---

*文档版本: V1.1 | 最后更新: 2026-04-17*
