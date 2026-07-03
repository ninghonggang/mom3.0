# MOM 3.0 模块依赖矩阵

> 版本：V1.0 | 最后更新：2026-07-03 16:18 | 维护人：架构组 / 小二
> 用途：自动化校验跨 module 引用双向一致性 + 上下游关系可视化
> 关联：[MOM3.0_V2.0_整体audit报告.md § 3.3](./MOM3.0_V2.0_整体audit报告.md)

---

## 0. 概述

MOM 3.0 共 16 个业务模块（含主数据/集成），本矩阵列出所有模块的**上下游依赖关系**，作为架构可视化和未来 module V2.0 重写的参考。

**当前 V2.0 覆盖**:**9/16** module(MES/APS/WMS/BPM/SCP/QMS/EAM/MDM/INT)。

---

## 1. 模块代号 + 名称 + V2.0 状态

| 代号 | 模块名 | V2.0 状态 | Commit | 行数 |
|------|--------|----------|--------|------|
| M01 | 系统管理 | ⏳ V1.x | — | — |
| **M02** | **MDM 主数据** | ✅ V2.0 | `2807f51` | 750 |
| **M03** | **MES 生产执行** | ✅ V2.0 | `0f09e79` | 730 |
| **M04** | **APS 计划** | ✅ V2.0 | `0f84028` | 937 |
| M05 | BPM 流程 | ✅ V2.0 | `80a4937` | 768 |
| M05 | 质量 QMS | ✅ V2.0 | `4721e1c` | 712 |
| **M06** | **EAM 设备** | ✅ V2.0 | `78ceda2` | 694 |
| **M07** | **WMS 仓储** | ✅ V2.0 | `874f517` | 892 |
| M08 | 数据采集 | ⏳ V1.x | — | — |
| M09 | 安灯系统 | ⏳ V1.x | — | — |
| M10 | 追溯管理 | ⏳ V1.x | — | — |
| M11 | 实验室 | ⏳ V1.x | — | — |
| M12 | 器具容器 | ⏳ V1.x | — | — |
| M13 | AI 质检 | ⏳ V1.x | — | — |
| **M14** | **INT 系统集成** | ✅ V2.0 | `63fbf71` | 750 |
| M15 | 报表 | ⏳ V1.x | — | — |
| **M16** | **SCP 供应链** | ✅ V2.0 | `a9bacc5` | 724 |

> 加粗 = 已 V2.0(9/16)

---

## 2. 上游矩阵（谁给我数据）

| Module | ERP | MDM | MES | WMS | APS | SCP | EAM | SCADA | AGV | BPM | QMS |
|--------|-----|-----|-----|-----|-----|-----|-----|-------|-----|-----|-----|
| **MDM** | ✅ | — | — | — | — | — | — | — | — | — | — |
| **MES** | — | ✅ | — | — | ✅ | ✅ | — | — | — | — | — |
| **APS** | — | ✅ | — | — | — | ✅ | — | — | — | — | — |
| **WMS** | — | ✅ | ✅ | — | — | ✅ | — | — | — | — | — |
| **BPM** | — | — | — | — | — | — | — | — | — | — | — |
| **SCP** | ✅ | ✅ | — | — | — | — | — | — | — | — | — |
| **QMS** | — | ✅ | ✅ | ✅ | — | — | — | — | — | — | — |
| **EAM** | ✅ | ✅ | — | — | — | — | — | ✅ | — | — | — |
| **INT** | ✅ | — | — | — | — | — | — | ✅ | ✅ | — | — |

**说明**：✅ = 该 module 是当前 module 的上游(给数据)。

---

## 3. 下游矩阵（我给谁数据）

| Module | ERP | MDM | MES | WMS | APS | SCP | EAM | 报表 | 追溯 | QMS | AGV |
|--------|-----|-----|-----|-----|-----|-----|-----|------|------|-----|-----|
| **MDM** | ✅ | — | ✅ | ✅ | ✅ | ✅ | ✅ | — | — | ✅ | — |
| **MES** | — | — | — | ✅ | — | — | — | ✅ | ✅ | — | — |
| **APS** | — | — | ✅ | — | — | — | — | ✅ | — | — | — |
| **WMS** | — | — | — | — | — | — | — | ✅ | ✅ | — | ✅ |
| **BPM** | — | — | — | — | — | — | — | ✅ | — | — | — |
| **SCP** | ✅ | — | ✅ | ✅ | ✅ | — | — | ✅ | — | — | — |
| **QMS** | — | — | ✅ | ✅ | — | — | — | ✅ | ✅ | — | — |
| **EAM** | — | — | ✅ | — | — | — | — | ✅ | — | ✅ | — |
| **INT** | ✅ | ✅ | — | ✅ | — | ✅ | ✅ | — | — | — | — |

**说明**：✅ = 当前 module 是该行的下游(给它数据)。

---

## 4. 双向一致性校验

| 关系 | A → B | B 确认收到? | 状态 |
|------|-------|-----------|------|
| MDM → MES | MDM 下游 = MES | MES 上游 = MDM | ✅ |
| MDM → WMS | MDM 下游 = WMS | WMS 上游 = MDM | ✅ |
| MDM → APS | MDM 下游 = APS | APS 上游 = MDM | ✅ |
| MDM → SCP | MDM 下游 = SCP | SCP 上游 = MDM | ✅ |
| MDM → EAM | MDM 下游 = EAM | EAM 上游 = MDM | ✅ |
| MDM → QMS | MDM 下游 = QMS | QMS 上游 = MDM | ✅ |
| MES → APS | APS 上游 = MES | MES 下游 = APS | ✅ |
| MES → WMS | MES 下游 = WMS | WMS 上游 = MES | ✅ |
| WMS → MES | WMS 下游 = MES | MES 上游 = WMS | ✅ |
| APS → MES | APS 下游 = MES | MES 上游 = APS | ✅ |
| SCP → APS | SCP 下游 = APS | APS 上游 = SCP | ✅ |
| SCP → WMS | SCP 下游 = WMS | WMS 上游 = SCP | ✅ |
| EAM → MES | EAM 下游 = MES | MES 上游 = EAM | ⚠️ 缺 |
| INT → MDM | INT 下游 = MDM | MDM 上游 = INT | ⚠️ 缺 |
| INT → SCP | INT 下游 = SCP | SCP 上游 = INT | ⚠️ 缺 |
| INT → EAM | INT 下游 = EAM | EAM 上游 = INT | ⚠️ 缺 |
| INT → WMS | INT 下游 = WMS | WMS 上游 = INT | ⚠️ 缺 |
| QMS → MES | QMS 下游 = MES | MES 上游 = QMS | ⚠️ 缺 |
| BPM → MES | BPM 下游 = MES(经审批) | MES 上游 = BPM | ⚠️ 缺 |

**统计**:**12/19 一致,7/19 需修订**。

---

## 5. 端到端业务流（核心 4 条）

### 5.1 销售订单到发货

```
SCP.SO.confirmed
  → APS.MPS.released
  → APS.Schedule.published
  → MES.WorkOrder.created
  → MES.Report.audited
  → QMS.FQC.passed
  → WMS.Putaway
  → WMS.Delivery.shipped
  → SCP.SO.shipped
```

**涉及 module**:**SCP → APS → MES → QMS → WMS → SCP**(6 跳,9 module 全部 5 跳参与)

### 5.2 采购订单到入库

```
SCP.PO.approved
  → SCP.PO.sent
  → SCP.ASN.received
  → WMS.ReceiveOrder
  → QMS.IQC.passed
  → WMS.Balance 更新
  → MES 物料就绪
```

**涉及**:**SCP → WMS → QMS → MES**(4 跳)

### 5.3 设备故障到工单挂起

```
EAM.Equipment.fault
  → MES.WorkOrder.hold
  → MES.Notification
  → QMS 可能触发 NCR
  → EAM.RepairOrder
  → EAM.Equipment.resume
  → MES.WorkOrder.resume
```

**涉及**:**EAM → MES → QMS → EAM**(循环 4 跳)

### 5.4 ERP 主数据同步

```
ERP.IDOC.MATMAS
  → INT.IDOC.received
  → INT.FieldMap.apply
  → MDM.Material.created/updated
  → APS/MES/WMS/SCP/QMS 引用
```

**涉及**:**ERP → INT → MDM → 5 module**(1 进 5 出)

---

## 6. 待修订事项

### 6.1 P1 修订(下次 module V2.0 同步)

| # | 修订 | 优先级 |
|---|------|--------|
| 1 | MES § 2.1 上游加 EAM(设备故障触发工单挂起) | P1 |
| 2 | MDM § 2.1 上游加 INT(ERP 主数据同步) | P1 |
| 3 | SCP § 2.1 上游加 INT(ERP 订单同步) | P1 |
| 4 | EAM § 2.1 上游加 INT(SCADA/PLC 数据) | P1 |
| 5 | WMS § 2.1 上游加 INT(AGV 任务) | P1 |
| 6 | MES § 2.1 上游加 BPM(工单变更审批) | P1 |
| 7 | MES § 2.1 上游加 QMS(质量问题触发返工工单) | P1 |

### 6.2 P2 修订(V2.1 实施时)

| # | 修订 | 优先级 |
|---|------|--------|
| 1 | QMS § 9 emoji 风格统一 ✅/❌ | P2 |
| 2 | APS § 6.1.4 字段类型说明统一引用 migration 0051 | P2 |
| 3 | 各 module production_orders 等字段名确认 | P2 |
| 4 | 错误码字典化(mdm_errno 表) | P2 |

---

## 7. CHANGELOG

| 版本 | 日期 | 修订人 | 说明 |
|------|------|--------|------|
| V1.0 | 2026-07-03 | 架构组 / 小二 | 初版,基于 9 module V2.0 整体 audit |
