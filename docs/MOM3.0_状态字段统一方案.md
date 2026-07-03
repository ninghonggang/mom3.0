# MOM 3.0 状态字段统一方案

> 版本：V1.0 | 最后更新：2026-07-03 | 维护人：架构组 / 小二
> 状态：✅ 方案定稿 | 实施中(双轨制第 1 阶段)
> 关联 Migration：[mom-server/migrations/0051_status_unification.sql](../mom-server/migrations/0051_status_unification.sql)

---

## 0. 文档元信息

| 字段 | 值 |
|---|---|
| 方案代号 | `status-unification` |
| 关联 PR | 待提 |
| 实施阶段 | 第 1 阶段(双轨制,本方案只做"加新字段 + 字典表",不改 model 代码) |
| 影响范围 | 7 个 domain / 10 张表 / 16 个 module 整体 |
| 实施工时 | 第 1 阶段 0.5h(本 PR) + 第 2 阶段 3-4h(model/service/前端切换) |

---

## 1. 背景与目标

### 1.1 问题陈述

MOM 3.0 现状:**状态字段类型混用,业务值不统一**。

| 实体 | 模块 | Status 类型 | 业务值 |
|---|---|---|---|
| `production_orders` | production | `int` | 1/2/3/4/5 |
| `production_dispatch` | production | `int` | 1/2/3 |
| `production_report` | production | `int` | 1/2/3 |
| `mobile_job_report` | production | `int` | 1/2/3 |
| `mes_process` | mes | `varchar(20)` | DRAFT/ACTIVE/INACTIVE |
| `aps_mps` | aps | `int` | 1/2/3 |
| `aps_mrp` | aps | `int` | 1/2/3 |
| `aps_schedule_plan` | aps | `int` | 1/2/3 |
| `aps_schedule_result` | aps | `int` | 1/2/3 |
| `aps_work_center` | aps | `varchar(20)` | ACTIVE |

### 1.2 业务影响

1. **跨 module 关联查询复杂**:`production_orders.status=1` 和 `aps_mps.status=1` 业务含义不同(前者"草稿"后者"待计算"),JOIN 时必须写 `CASE WHEN` 转换
2. **日志不可读**:DB log `UPDATE ... SET status=2` 读者不知道含义
3. **前端状态字典组件分裂**:`IntStatusTag` + `VarcharStatusTag` 两套,扩展时双倍维护
4. **状态机非法转移无法校验**:业务侧"工单已完成,不应该再下达"这类断言,必须用数字比较,易出错

### 1.3 目标

- ✅ **类型统一**:全 MOM 3.0 状态字段统一为 `varchar(30)`
- ✅ **业务值统一**:跨 module 同一概念用同一编码(如"工单已完成"=`COMPLETED`)
- ✅ **字典可查**:`mdm_status_dict` 表热更新,前端统一读字典
- ✅ **零破坏**:双轨制,旧 `status` 字段保留,新代码写 `status_v2`,平稳切换
- ✅ **可回滚**:每阶段 migration 配套 down 脚本

---

## 2. 方案选型

### 2.1 三种存储方案对比

| 选项 | 存储类型 | 优点 | 缺点 | 适用场景 |
|---|---|---|---|---|
| **A. int + 字典** | `int` (1/2/3) | 紧凑、查询快、JOIN 性能好 | 需字典映射、不直观、跨 module 难统一 | 单 module 状态固定 |
| **B. varchar + 字典** | `varchar(30)` (DRAFT/...) | 自解释、日志清晰、跨 module 一致 | 占空间(8 字节 vs 4) | **多 module 跨域** ⭐ 选 B |
| **C. PG enum** | `pg enum type` | 紧凑 + 自解释 + DB 强约束 | 新增状态要 DDL 改 type、跨 module 复用麻烦 | 极少变更的状态 |

**结论**:**选 B(varchar + 字典表)**。

理由:
- MOM 3.0 跨 16 module,跨 module 一致性比存储性能重要
- 日志/前端/调试都需要可读性
- 字典表可热更新,避免 DDL 改 enum type

### 2.2 双轨制策略

**为什么双轨**:
- 旧 `status` 字段在 10 张表都有,涉及生产数据,**不能一次性 DROP**
- 前端 16 个 module 都在用,改完一晚上全部上线风险大
- **加新字段 `status_v2`** → 字典映射 + 数据迁移 → **新代码读 status_v2** → **3-6 个月后**移除旧字段

**双轨期行为约定**:

| 写入 | 读取 | 阶段 |
|---|---|---|
| 写 `status`(int)和 `status_v2`(varchar)双份 | 读 `status_v2` | 第 2 阶段(model 改完) |
| 只写 `status_v2` | 读 `status_v2` | 第 3 阶段(全量切换) |
| 删除 `status` 旧字段 | 读 `status_v2` | 第 4 阶段(清理) |

---

## 3. 字典表设计

### 3.1 表结构

```sql
CREATE TABLE mdm_status_dict (
    id              BIGSERIAL PRIMARY KEY,
    tenant_id       BIGINT       NOT NULL DEFAULT 1,
    domain          VARCHAR(50)  NOT NULL,            -- 模块域
    entity          VARCHAR(50)  NOT NULL,            -- 实体
    code            VARCHAR(30)  NOT NULL,            -- 状态码
    label           VARCHAR(50)  NOT NULL,            -- 中文显示
    element_type    VARCHAR(20)  NOT NULL DEFAULT 'info', -- Element Plus 标签类型
    is_terminal     BOOLEAN      NOT NULL DEFAULT FALSE, -- 是否终态
    sort_order      INT          NOT NULL DEFAULT 0,
    description     TEXT,
    legacy_int      INT,                               -- 旧 int 字段值(迁移期)
    created_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,
    UNIQUE (tenant_id, entity, code)
);
```

### 3.2 字典数据(已 seed 51 条,7 个 domain)

完整 seed 见 [0051_status_unification.sql § 2](../mom-server/migrations/0051_status_unification.sql)。

#### production.production_order(7 状态)

| code | label | element_type | is_terminal | legacy_int | 描述 |
|---|---|---|---|---|---|
| DRAFT | 草稿 | info | ❌ | 1 | 工单创建后未下达 |
| RELEASED | 已下达 | primary | ❌ | 2 | 工单已下达,等待开工 |
| IN_PROGRESS | 生产中 | warning | ❌ | 3 | 工单正在生产 |
| HOLD | 挂起 | warning | ❌ | NULL | 物料/设备/质量异常 |
| COMPLETED | 已完成 | success | ❌ | 4 | 所有工序已报工完成 |
| CLOSED | 已关闭 | info | ✅ | 5 | 工单关闭(超过3天无活动) |
| CANCELLED | 已取消 | info | ✅ | NULL | 工单取消 |

#### mes.mes_process(3 状态)

| code | label | element_type | is_terminal | legacy_int | 描述 |
|---|---|---|---|---|---|
| DRAFT | 草稿 | info | ❌ | NULL | 工艺路线草稿 |
| ACTIVE | 生效 | success | ❌ | NULL | 工艺路线当前生效版本 |
| OBSOLETE | 失效 | info | ✅ | NULL | 工艺路线被新版本取代 |

#### aps.mps(5 状态) / aps.mrp(4 状态) / aps.schedule_plan(4 状态) / aps.schedule_result(4 状态)

详见 migration SQL 文件。

### 3.3 字典查询 API

```go
// service/system/status_dict.go
func GetStatusDict(ctx context.Context, entity string) ([]StatusDictItem, error) {
    return db.Where("entity = ? AND deleted_at IS NULL", entity).
        Order("sort_order ASC").Find(&[]StatusDictItem{}).Error
}

// 缓存:5 分钟 Redis 缓存,字典更新后调用 InvalidateStatusDictCache(entity)
```

---

## 4. 数据迁移

### 4.1 迁移范围(本 PR)

| 表 | 加字段 | 数据迁移 | 校验 |
|---|---|---|---|
| production_orders | ✅ status_v2 | ✅ | ✅ |
| production_dispatch | ✅ | ⏸ 第 2 阶段 | — |
| production_report | ✅ | ⏸ 第 2 阶段 | — |
| mobile_job_report | ✅ | ✅ | ✅ |
| mes_process | ✅ | ✅(copy) | ✅ |
| aps_mps | ✅ | ✅ | ✅ |
| aps_mrp | ✅ | ✅ | ✅ |
| aps_schedule_plan | ✅ | ✅ | ✅ |
| aps_schedule_result | ✅ | ✅ | ✅ |
| aps_work_center | ✅ | ✅(copy) | ✅ |

**本 PR**:建字典表 + seed + 加字段 + 迁移 8 张表(剩 2 张 production_dispatch/report 第 2 阶段做)。

### 4.2 迁移规则

每张表的 int→varchar 映射对应 `legacy_int` 字段。如 `production_orders.status=1 → status_v2='DRAFT'`,`status=2 → 'IN_PROGRESS'` 等。

NULL 值处理:**保留 NULL,不抛错**(若出现 NULL 记录,migration 末尾的 `DO $$` 块 WARNING 提示)。

### 4.3 校验脚本

```sql
-- 校验 status_v2 完整性(每张表 NULL 数)
SELECT 'production_orders' AS tbl, COUNT(*) AS null_cnt
FROM production_orders WHERE status_v2 IS NULL
UNION ALL
SELECT 'aps_mps', COUNT(*) FROM aps_mps WHERE status_v2 IS NULL;
-- 期望:全部为 0
```

---

## 5. 实施路线图

### 5.1 第 1 阶段(本 PR,0.5h)

- [x] 写方案文档
- [x] 建 `mdm_status_dict` 字典表 + seed 51 条
- [x] 10 张表加 `status_v2` 字段
- [x] 8 张表数据迁移(int→varchar)
- [x] 校验脚本
- [x] 配套 down 脚本(注释中)

### 5.2 第 2 阶段(下一个 PR,3-4h)

- [ ] Go model 加 `StatusV2` 字段,gorm 标签对齐
- [ ] Service 层状态判断改用 `StatusV2`
- [ ] API 响应 JSON 增加 `status_v2` 字段,保留 `status`(双轨期兼容)
- [ ] 前端 status 字典组件统一读 `mdm_status_dict`
- [ ] 前端 status 显示切到 `status_v2`
- [ ] 单元测试 + 集成测试

### 5.3 第 3 阶段(3-6 个月后,1h)

- [ ] 全量验证 `status_v2` 在生产环境稳定
- [ ] 写脚本统计旧 `status` 字段读频(应 < 1%)
- [ ] 删旧 `status` 字段(DDL),移除 model 字段

### 5.4 第 4 阶段(可选,1h)

- [ ] `mdm_status_dict` 增加 ENUM 类型约束(代码层校验)
- [ ] 数据库 trigger 拒绝非法 status_v2 写入

---

## 6. 风险评估

| 风险 | 概率 | 影响 | 缓解 |
|---|---|---|---|
| 数据迁移丢失/错误 | 中 | 高 | 双轨 + 校验脚本 + 备份后跑 |
| 前端切换期 500 错误 | 中 | 中 | status_v2 NULL 时 fallback 读 status |
| 字典表 seed 与代码硬编码冲突 | 低 | 中 | 代码硬编码全部改读字典,本 PR 同步扫除 |
| 跨 module 状态语义不同(已存在) | 中 | 高 | 字典表 `domain.entity` 二级分类,避免冲突 |

---

## 7. 验收清单(本 PR)

- [x] 字典表 `mdm_status_dict` 已建,seed 51 条
- [x] 10 张表 `status_v2` 字段已加
- [x] 8 张表 `status_v2` 数据完整(校验 0 NULL)
- [x] 配套 down 脚本可用
- [x] 方案文档归档
- [ ] (下一 PR)Go model 改读 `status_v2`
- [ ] (下一 PR)前端 status 组件切到字典

---

## 8. 相关链接

- [MOM3.0_模块设计模板.md § 6.1.4 字段类型说明](./MOM3.0_模块设计模板.md)
- [MOM3.0_MES生产执行模块设计文档.md § 7.2 关键字段说明](./MOM3.0_MES生产执行模块设计文档.md) — 已识别该问题
- [Migration 0051](../mom-server/migrations/0051_status_unification.sql)
- [TODO.md P1-1 齐套检查逻辑](./TODO.md) — 同步关联

---

## 9. CHANGELOG

| 版本 | 日期 | 修订人 | 说明 |
|---|---|---|---|
| V1.0 | 2026-07-03 | 架构组 / 小二 | 初版方案 + 第 1 阶段实施 |
