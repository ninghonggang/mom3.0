# MOM3.0 开发进度记录

**版本**: V1.0 | **更新日期**: 2026-04-06 | **项目**: 峰梅动力MOM3.0

> 本文档记录所有待开发功能清单，按优先级分组。开发完成后在此更新状态。

---

## P0 紧急 - 待开发

| 功能 | 模块 | 说明 | 状态 | 开发日期 |
|------|------|------|------|----------|
| BOM批量导入完整实现 | M02 | 当前返回"开发中"，需实现ParseBOMExcel + ImportBOMs | ❌ 未开发 | - |
| 质量模块前端连接API | M05 | IQC/IPQC/FQC/OQC/NCR/DefectRecord的增删改需连接真实API | ❌ 未开发 | - |

---

## P1 重要 - 待开发

### M01 系统管理
| 功能 | 说明 | 状态 | 开发日期 |
|------|------|------|----------|
| 多车间管理 | WorkshopConfig + sys_workshop扩展字段 | ❌ 未开发 | - |
| 打印模板管理 | PrintTemplate + sys_print_template表 | ❌ 未开发 | - |
| 通知公告 | Notice + sys_notice表 | ❌ 未开发 | - |

### M02 主数据管理
| 功能 | 说明 | 状态 | 开发日期 |
|------|------|------|----------|
| BOM金蝶同步逻辑 | ErpBomCode等字段已有，需实现同步Worker | ⚠️ 字段已加 | - |

### M03 生产执行
| 功能 | 说明 | 状态 | 开发日期 |
|------|------|------|----------|
| 首末件检验对话框 | 当前handleAdd/Edit是ElMessage.info | ⚠️ 部分完成 | - |
| 电子SOP | ElectronicSOP | ❌ 未开发 | - |
| 批次号编码规则 | mes_code_rule | ❌ 未开发 | - |
| 生产指示单(流程卡) | FlowCard | ❌ 未开发 | - |

### M04 APS
| 功能 | 说明 | 状态 | 开发日期 |
|------|------|------|----------|
| APS-008 产能分析看板 | Capacity Analysis | ❌ 未开发 | - |
| APS-010 交付率/达成率 | Delivery Rate | ❌ 未开发 | - |
| APS-011 换型矩阵管理 | Changeover Matrix | ❌ 未开发 | - |
| APS-012 滚动排程 | Rolling Schedule | ❌ 未开发 | - |
| APS-013 JIT/JIS接入 | JIT/JIS Demand | ❌ 未开发 | - |
| AI智能排程完善 | OR-Tools/GA集成 | ⚠️ 仅基础规则 | - |

### M06 设备管理
| 功能 | 说明 | 状态 | 开发日期 |
|------|------|------|----------|
| TEEP分析 | TEEP Calculation | ❌ 未开发 | - |
| 模具管理 | Mold Management | ❌ 未开发 | - |
| 量检具管理 | Gauge Management | ❌ 未开发 | - |

### M07 WMS仓储
| 功能 | 说明 | 状态 | 开发日期 |
|------|------|------|----------|
| 调拨管理 | Transfer Order | ❌ 未开发 | - |
| 盘点管理 | Stock Check | ❌ 未开发 | - |
| JIT物料拉动 | JIT Pull | ❌ 未开发 | - |
| 电子看板拉动 | Kanban Pull | ❌ 未开发 | - |
| 线边库位管理 | Side Location | ❌ 未开发 | - |
| InventoryList.handleDelete | 当前是stub | ⚠️ 需修复 | - |

---

## P2 常规 - 待开发

### M05 质量管理
| 功能 | 说明 | 状态 | 开发日期 |
|------|------|------|----------|
| SPC CP/CPK计算 | SPC Lab Capability Indices | ❌ 未开发 | - |
| 蓝牙量具接口 | Bluetooth Gauge | ❌ 未开发 | - |
| 检验标准管理 | Inspect Standard | ❌ 未开发 | - |
| Andon升级机制 | Escalation + 多级通知 | ❌ 未开发 | - |
| 追溯API完善 | /trace/forward/backward路由 | ⚠️ 方法存在无路由 | - |

### M12 器具管理
| 功能 | 说明 | 状态 | 开发日期 |
|------|------|------|----------|
| 器具管理 | Container/Gauge | ❌ 未开发 | - |

### M13 AI质检
| 功能 | 说明 | 状态 | 开发日期 |
|------|------|------|----------|
| 视觉检测 | Vision Detection | ❌ 未开发 | - |

### M14 系统集成
| 功能 | 说明 | 状态 | 开发日期 |
|------|------|------|----------|
| 金蝶ERP同步 | Kingdee ERP Sync | ❌ 未开发 | - |
| AGV调度接入 | AGV Integration | ❌ 未开发 | - |
| 消息推送 | Feishu/Wechat/Audio | ❌ 未开发 | - |
| 供应商Portal | Supplier Portal | ❌ 未开发 | - |

### M15 报表
| 功能 | 说明 | 状态 | 开发日期 |
|------|------|------|----------|
| 全部8个报表 | 生产日报/质量周报/OEE/交付率等 | ❌ 未开发 | - |

---

## 已完成功能清单

### P0 已完成 (2026-04-06)
| 功能 | 模块 | 说明 |
|------|------|------|
| ✅ M08数据采集 | M08 | dc_data_point/scan_log/collect_record三表+CRUD |
| ✅ 首末件检验 | M03 | mes_first_last_inspect表+API+前端页面 |
| ✅ 包装条码管理 | M03 | mes_package表+create/add-item/seal+前端页面 |
| ✅ BOM金蝶同步字段 | M02 | ErpBomCode/ErpSyncTime/ErpSyncStatus/IsCurrent |
| ✅ BOM批量导入骨架 | M02 | 模板下载/上传接口(解析逻辑TODO) |

### 更早已完成
| 功能 | 模块 |
|------|------|
| ✅ 用户/角色/菜单/部门/岗位/字典管理 | M01 |
| ✅ 物料/BOM/工艺/工序/车间/产线/工位/班次 | M02 |
| ✅ 客户管理 | M02 |
| ✅ 物料批量导入(完整) | M02 |
| ✅ 工单/派工/报工 | M03 |
| ✅ APS甘特图/工作中心/缺料分析 | M04 |
| ✅ OEE分析 | M06 |
| ✅ 仓库/库位/库存/收货/发货 | M07 |
| ✅ IQC/IPQC/FQC/OQC/NCR/SPC列表API | M05 |
| ✅ 缺陷代码管理 | M05 |
| ✅ 供应商管理 | M02 |
| ✅ 按钮级权限控制 | M01 |

---

## 更新日志

| 日期 | 版本 | 更新内容 |
|------|------|---------|
| 2026-04-06 | V1.0 | 初始化开发状态文档，记录所有P0/P1/P2待开发功能 |
