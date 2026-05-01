# MOM3.0 基础数据设计文档

**版本**: V1.0 | **所属模块**: M01系统管理 + M02主数据管理 | **项目**: 闻荫科技MOM3.0

> **前置阅读**: 请先阅读 [MOM3.0_架构与主文档.md](./MOM3.0_架构与主文档.md) 了解统一架构和技术规范。

说明 - 本文档是MOM3.0架构总览的模块补充，详细规定M01系统管理和M02主数据管理的功能范围、API接口、数据库设计。

## 目录

- [4. M01 系统管理](#4-m01-系统管理)
  - [4.1 当前状态](#41-当前状态)
  - [4.2 新增：多车间管理](#42-新增多车间管理)
  - [4.3 新增：打印模板管理](#43-新增打印模板管理)
- [5. M02 主数据管理](#5-m02-主数据管理)
  - [5.1 当前状态](#51-当前状态)
  - [5.2 BOM扩展（金蝶同步）](#52-bom扩展金蝶同步)
  - [5.3 批量导入设计](#53-批量导入设计)

---

---

## 4. M01 系统管理

### 4.1 当前状态

| 功能 | 状态 | Handler | 前端 |
|------|------|---------|------|
| 用户管理 | ✅ 完成 | user.go | UserList.vue |
| 角色管理 | ✅ 完成 | role.go | RoleList.vue |
| 菜单管理 | ✅ 完成 | menu.go | MenuList.vue |
| 部门管理 | ✅ 完成 | dept.go | DeptList.vue |
| 岗位管理 | ✅ 完成 | post.go | PostList.vue |
| 字典管理 | ✅ 完成 | dict.go | DictList.vue |
| 系统参数配置 | ✅ 完成 | config.go | SystemConfig.vue |
| 操作日志 | ✅ 完成 | oper_log.go | OperLogList.vue |
| 登录日志 | ✅ 完成 | login_log.go | LoginLogList.vue |
| 按钮权限配置 | ✅ 完成 | button_permission.go | RoleList.vue |
| 租户管理 | ✅ 完成 | tenant.go | TenantList.vue |
| **多车间管理** | ❌ 需开发 | workshop_config.go | WorkshopConfig.vue |
| **打印模板管理** | ❌ 需开发 | print_template.go | PrintTemplate.vue |
| **通知公告** | ❌ 需开发 | notice.go | NoticeList.vue |

### 4.2 新增：多车间管理

闻荫科技有3个车间独立管理，数据需隔离但可汇总分析。

**数据模型**:
```sql
-- 车间配置表（已有workshop表，扩展字段）
ALTER TABLE sys_workshop ADD COLUMN workshop_type VARCHAR(20);  -- MACHINING/BATTERY/TRIAL
ALTER TABLE sys_workshop ADD COLUMN max_devices INTEGER DEFAULT 0;
ALTER TABLE sys_workshop ADD COLUMN erp_plant_code VARCHAR(50);  -- 金蝶工厂编码
```

**业务规则**:
- 所有业务表均有 `workshop_id` 字段用于数据隔离
- 用户可绑定多个车间，Token中携带当前车间上下文
- 超级管理员可跨车间查看汇总数据

### 4.3 新增：打印模板管理

支持以下标签类型的模板配置:

| 模板类型 | 用途 | 打印时机 |
|---------|------|---------|
| RAW_MATERIAL | 原材料标签 | IQC入库 |
| SEMI_PRODUCT | 半成品标签 | 工序完工 |
| FINISHED_GOODS | 成品标签 | FQC合格 |
| PACKAGE_BOX | 包装箱标签 | 装箱 |
| DEVICE | 设备标签 | 设备台账 |
| CONTAINER | 器具标签 | 器具入库 |
| MOLD | 模具标签 | 模具台账 |

```sql
CREATE TABLE sys_print_template (
    id BIGSERIAL PRIMARY KEY,
    template_code VARCHAR(50) UNIQUE NOT NULL,
    template_name VARCHAR(100) NOT NULL,
    template_type VARCHAR(30) NOT NULL,     -- 见上表
    template_content TEXT,                   -- ZPL/TSPL模板内容
    template_width INTEGER DEFAULT 100,      -- 标签宽度(mm)
    template_height INTEGER DEFAULT 50,      -- 标签高度(mm)
    preview_image VARCHAR(500),              -- 预览图片路径
    printer_type VARCHAR(20),               -- ZEBRA/TSC/BROTHER
    is_default SMALLINT DEFAULT 0,
    is_enabled SMALLINT DEFAULT 1,
    remark VARCHAR(500),
    tenant_id BIGINT NOT NULL,
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

**API**:
```
GET    /api/v1/system/print-templates          # 列表
POST   /api/v1/system/print-templates          # 新增
PUT    /api/v1/system/print-templates/:id      # 更新
DELETE /api/v1/system/print-templates/:id      # 删除
POST   /api/v1/system/print-templates/:id/preview  # 预览
POST   /api/v1/system/print/execute            # 执行打印
```

---

## 5. M02 主数据管理

### 5.1 当前状态

| 功能 | 状态 | 备注 |
|------|------|------|
| 物料管理 | ✅ 完成 | |
| 物料分类 | ✅ 完成 | |
| BOM管理 | ✅ 完成 | 需增加金蝶同步字段 |
| 工艺路线 | ✅ 完成 | |
| 工序管理 | ✅ 完成 | |
| 车间/产线/工位 | ✅ 完成 | |
| 班次管理 | ✅ 完成 | |
| 供应商管理 | ✅ 完成 | |
| 客户管理 | ❌ 需开发 | |
| 物料批量导入 | ❌ 需开发 | |
| BOM批量导入 | ❌ 需开发 | |

### 5.2 BOM扩展（金蝶同步）

```sql
ALTER TABLE mdm_bom ADD COLUMN erp_bom_code VARCHAR(50);       -- 金蝶BOM编码
ALTER TABLE mdm_bom ADD COLUMN erp_sync_time TIMESTAMP;
ALTER TABLE mdm_bom ADD COLUMN erp_sync_status VARCHAR(20);    -- SYNCED/PENDING/FAILED
ALTER TABLE mdm_bom ADD COLUMN version VARCHAR(20);            -- 版本号
ALTER TABLE mdm_bom ADD COLUMN is_current SMALLINT DEFAULT 1;  -- 是否当前版本
```

### 5.3 批量导入设计

**物料导入Excel模板字段**:

| 字段 | 必填 | 说明 | 示例 |
|------|------|------|------|
| material_code | 是 | 物料编码 | BZZ-001 |
| material_name | 是 | 物料名称 | 平衡轴总成 |
| material_type | 是 | 类型: RAW/SEMI/FINISHED | FINISHED |
| spec | 否 | 规格型号 | 6缸 |
| unit | 是 | 单位 | 件 |
| purchase_price | 否 | 采购单价 | 125.00 |
| safety_stock | 否 | 安全库存 | 100 |
| supplier_code | 否 | 供应商编码 | SUP-001 |
| erp_code | 否 | ERP物料编码 | |

**导入流程**:
```
上传Excel → 解析预览（显示错误行） → 用户确认 → 异步导入 → 返回结果报告
```

**API**:
```
POST /api/v1/mdm/materials/import          # 上传并解析
POST /api/v1/mdm/materials/import/confirm  # 确认导入
GET  /api/v1/mdm/materials/import/template # 下载模板
GET  /api/v1/system/import-tasks/:id       # 查询导入进度
```

---

## 10. 通知公告管理

### 10.1 功能概述

通知公告模块用于向用户发布系统公告、运维通知、活动信息等，支持按部门/角色/用户定向推送。

### 10.2 数据模型

```sql
-- 通知公告表
CREATE TABLE sys_notice (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    title VARCHAR(200) NOT NULL,                    -- 标题
    content TEXT,                                   -- 内容
    notice_type VARCHAR(20),                        -- 公告类型: SYSTEM/OPERATION/MAINTENANCE/OTHER
    priority INTEGER DEFAULT 1,                    -- 优先级: 1普通/2重要/3紧急
    publish_dept VARCHAR(100),                     -- 发布部门
    publisher_id BIGINT,                            -- 发布人ID
    publisher_name VARCHAR(50),                     -- 发布人姓名
    publish_time TIMESTAMP,                         -- 发布时间
    effect_time TIMESTAMP,                          -- 生效时间
    expire_time TIMESTAMP,                          -- 失效时间
    target_type VARCHAR(20),                        -- 发布范围: ALL/DEPT/ROLE/USER
    target_ids VARCHAR(500),                        -- 目标ID列表
    is_top INTEGER DEFAULT 0,                       -- 是否置顶
    status INTEGER DEFAULT 1,                        -- 1草稿/2已发布/3已撤回
    view_count INTEGER DEFAULT 0,                   -- 阅读次数
    remark VARCHAR(500),
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 公告阅读记录表
CREATE TABLE sys_notice_read_record (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    notice_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    user_name VARCHAR(50),
    read_time VARCHAR(30),
    created_at TIMESTAMP DEFAULT NOW()
);
```

### 10.3 API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/system/notices | 公告列表（分页） |
| GET | /api/v1/system/notices/:id | 获取公告详情 |
| POST | /api/v1/system/notices | 创建公告 |
| PUT | /api/v1/system/notices/:id | 更新公告 |
| DELETE | /api/v1/system/notices/:id | 删除公告 |
| POST | /api/v1/system/notices/:id/publish | 发布公告 |
| GET | /api/v1/system/notices/my | 获取我的公告 |
| POST | /api/v1/system/notices/:id/read | 标记已读 |

### 10.4 业务规则

- 公告发布后不可直接修改，需先撤回再修改
- 支持定时发布（设置effect_time）
- 支持过期自动下架（设置expire_time）
- 用户阅读后自动记录阅读时间和阅读人
- 紧急公告支持短信/邮件推送（需配置短信邮件服务）

---

## 11. 多租户管理

### 11.1 功能概述

多租户模块支持多个独立企业/工厂共享系统实例，每个租户数据完全隔离。支持SaaS化部署模式。

### 11.2 数据模型

```sql
-- 租户表
CREATE TABLE sys_tenant (
    id BIGSERIAL PRIMARY KEY,
    tenant_name VARCHAR(100) NOT NULL,              -- 租户名称
    tenant_key VARCHAR(50) NOT NULL UNIQUE,          -- 租户标识
    province VARCHAR(50),                           -- 省份
    city VARCHAR(50),                               -- 城市
    district VARCHAR(50),                           -- 区县
    address VARCHAR(200),                           -- 详细地址
    manager VARCHAR(50),                            -- 负责人
    contact_name VARCHAR(50),                       -- 联系人
    contact_phone VARCHAR(20),                      -- 联系电话
    contact_email VARCHAR(100),                     -- 联系邮箱
    factory_type VARCHAR(50),                       -- 工厂类型
    employee_count INTEGER,                         -- 员工人数
    area_size DECIMAL(10,2),                        -- 占地面积(平方米)
    annual_capacity DECIMAL(15,2),                  -- 年产能
    status INTEGER DEFAULT 1,                       -- 1正常/2禁用
    expire_time TIMESTAMP,                          -- 到期时间
    remark VARCHAR(500),
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 租户套餐表
CREATE TABLE sys_tenant_package (
    id BIGSERIAL PRIMARY KEY,
    package_name VARCHAR(100) NOT NULL,             -- 套餐名称
    package_type VARCHAR(20),                      -- 套餐类型: BASIC/PROFESSIONAL/ENTERPRISE
    module_codes VARCHAR(500),                     -- 授权模块编码列表
    max_users INTEGER,                              -- 最大用户数
    max_storage_gb INTEGER,                         -- 最大存储空间(GB)
    price DECIMAL(10,2),                            -- 价格
    status INTEGER DEFAULT 1,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 租户套餐关联表
CREATE TABLE sys_tenant_package_relation (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL,
    package_id BIGINT NOT NULL,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);
```

### 11.3 API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/system/tenants | 租户列表（分页） |
| GET | /api/v1/system/tenants/:id | 获取租户详情 |
| POST | /api/v1/system/tenants | 创建租户 |
| PUT | /api/v1/system/tenants/:id | 更新租户 |
| DELETE | /api/v1/system/tenants/:id | 删除租户 |
| PUT | /api/v1/system/tenants/:id/status | 更新租户状态 |
| GET | /api/v1/system/tenant-packages | 套餐列表 |
| POST | /api/v1/system/tenant-packages | 创建套餐 |

### 11.4 业务规则

- 租户隔离方式：每个表通过tenant_id字段隔离
- 租户创建时自动初始化默认角色、菜单权限
- 租户到期前7天/3天/1天发送提醒通知
- 租户禁用后用户无法登录，但数据保留
- 超级管理员(admin)不受租户限制

---

## 12. 短信邮件服务

### 12.1 功能概述

短信邮件服务提供统一的第三方通知通道，支持验证码发送、通知提醒等场景。

### 12.2 数据模型

```sql
-- 短信配置表
CREATE TABLE sys_sms_config (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    config_name VARCHAR(100) NOT NULL,              -- 配置名称
    provider VARCHAR(30),                           -- 提供商: ALIYUN/TENCENT/HUAWEI
    app_id VARCHAR(100),                           -- 应用ID
    app_secret VARCHAR(200),                        -- 应用密钥(加密存储)
    sign_name VARCHAR(50),                         -- 签名
    status INTEGER DEFAULT 1,                       -- 1启用/0禁用
    remark VARCHAR(500),
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 短信发送记录表
CREATE TABLE sys_sms_send_log (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    config_id BIGINT,                               -- 短信配置ID
    mobile VARCHAR(20) NOT NULL,                   -- 手机号
    template_code VARCHAR(50),                      -- 模板编码
    template_params TEXT,                            -- 模板参数JSON
    content VARCHAR(500),                           -- 短信内容
    send_status VARCHAR(20),                        -- SEND_SUCCESS/SEND_FAIL/VERIFY_SUCCESS/VERIFY_FAIL
    send_time TIMESTAMP DEFAULT NOW(),
    verify_code VARCHAR(10),                        -- 验证码
    verify_expire_time TIMESTAMP,                   -- 验证码过期时间
    error_msg VARCHAR(200),                         -- 错误信息
    created_at TIMESTAMP DEFAULT NOW()
);

-- 邮件配置表
CREATE TABLE sys_mail_config (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    config_name VARCHAR(100) NOT NULL,              -- 配置名称
    smtp_host VARCHAR(100),                        -- SMTP服务器
    smtp_port INTEGER DEFAULT 465,                  -- 端口
    smtp_username VARCHAR(100),                     -- 用户名
    smtp_password VARCHAR(200),                     -- 密码(加密存储)
    from_address VARCHAR(100),                      -- 发件人地址
    from_name VARCHAR(50),                          -- 发件人昵称
    use_tls BOOLEAN DEFAULT true,                   -- 是否使用TLS
    status INTEGER DEFAULT 1,                       -- 1启用/0禁用
    remark VARCHAR(500),
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 邮件发送记录表
CREATE TABLE sys_mail_send_log (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    config_id BIGINT,                               -- 邮件配置ID
    to_address VARCHAR(200) NOT NULL,               -- 收件人地址
    cc_address VARCHAR(500),                        -- 抄送地址
    subject VARCHAR(200),                           -- 邮件主题
    content TEXT,                                   -- 邮件内容
    template_code VARCHAR(50),                     -- 模板编码
    template_params TEXT,                           -- 模板参数JSON
    send_status VARCHAR(20),                       -- SEND_SUCCESS/SEND_FAIL
    send_time TIMESTAMP DEFAULT NOW(),
    error_msg VARCHAR(200),                         -- 错误信息
    created_at TIMESTAMP DEFAULT NOW()
);

-- 消息模板表
CREATE TABLE sys_message_template (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    template_code VARCHAR(50) UNIQUE NOT NULL,     -- 模板编码
    template_name VARCHAR(100) NOT NULL,           -- 模板名称
    channel_type VARCHAR(20),                      -- 渠道类型: SMS/MAIL
    template_content TEXT,                          -- 模板内容（支持变量${name}）
    template_params TEXT,                           -- 参数字段定义JSON
    status INTEGER DEFAULT 1,                       -- 1启用/0禁用
    remark VARCHAR(500),
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### 12.3 API接口

**短信服务API**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/system/sms/configs | 短信配置列表 |
| POST | /api/v1/system/sms/configs | 创建短信配置 |
| PUT | /api/v1/system/sms/configs/:id | 更新短信配置 |
| DELETE | /api/v1/system/sms/configs/:id | 删除短信配置 |
| GET | /api/v1/system/sms/logs | 短信发送记录 |
| POST | /api/v1/system/sms/send | 发送短信 |
| POST | /api/v1/system/sms/send-code | 发送验证码 |
| POST | /api/v1/system/sms/verify-code | 校验验证码 |

**邮件服务API**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/system/mail/configs | 邮件配置列表 |
| POST | /api/v1/system/mail/configs | 创建邮件配置 |
| PUT | /api/v1/system/mail/configs/:id | 更新邮件配置 |
| DELETE | /api/v1/system/mail/configs/:id | 删除邮件配置 |
| GET | /api/v1/system/mail/logs | 邮件发送记录 |
| POST | /api/v1/system/mail/send | 发送邮件 |
| GET | /api/v1/system/mail/templates | 消息模板列表 |
| POST | /api/v1/system/mail/templates | 创建消息模板 |

### 12.4 业务规则

- 短信验证码有效期默认5分钟，验证失败3次后需重新获取
- 短信/邮件发送失败自动重试3次，间隔1分钟
- 敏感信息（密码、密钥）加密存储
- 发送频率限制：单个手机号每分钟最多1条，每小时最多10条
- 邮件支持HTML格式和纯文本格式

---

## 13. OAuth2认证

### 13.1 功能概述

OAuth2认证模块支持第三方应用通过标准OAuth2协议接入系统。

### 13.2 数据模型

```sql
-- OAuth2客户端表
CREATE TABLE sys_oauth2_client (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    client_id VARCHAR(50) UNIQUE NOT NULL,          -- 客户端ID
    client_secret VARCHAR(200) NOT NULL,           -- 客户端密钥(加密存储)
    client_name VARCHAR(100) NOT NULL,             -- 客户端名称
    client_type VARCHAR(20),                        -- 客户端类型: WEB/APP/PC/SERVICE
    grant_types VARCHAR(200),                        -- 授权类型: authorization_code,client_credentials,refresh_token
    redirect_uri VARCHAR(500),                      -- 回调地址
    scopes VARCHAR(500),                            -- 授权范围
    logo_url VARCHAR(500),                         -- logo地址
    status INTEGER DEFAULT 1,                       -- 1启用/0禁用
    expire_time TIMESTAMP,                          -- 过期时间
    remark VARCHAR(500),
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- OAuth2 Token表
CREATE TABLE sys_oauth2_token (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    client_id VARCHAR(50) NOT NULL,                 -- 客户端ID
    user_id BIGINT,                                -- 用户ID
    user_type VARCHAR(20),                         -- 用户类型
    access_token VARCHAR(500) NOT NULL,            -- 访问令牌
    refresh_token VARCHAR(500),                    -- 刷新令牌
    token_type VARCHAR(20) DEFAULT 'Bearer',       -- Token类型
    grant_type VARCHAR(30),                        -- 授权类型
    scopes VARCHAR(500),                           -- 授权范围
    expires_at TIMESTAMP,                          -- 过期时间
    refresh_expires_at TIMESTAMP,                  -- 刷新令牌过期时间
    ip_address VARCHAR(50),                        -- IP地址
    user_agent VARCHAR(500),                        -- User-Agent
    created_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT uk_access_token UNIQUE (access_token),
    CONSTRAINT uk_refresh_token UNIQUE (refresh_token)
);

-- OAuth2授权记录表
CREATE TABLE sys_oauth2_authorize (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    client_id VARCHAR(50) NOT NULL,
    user_id BIGINT NOT NULL,
    response_type VARCHAR(20),
    redirect_uri VARCHAR(500),
    scopes VARCHAR(500),
    state VARCHAR(100),
    code VARCHAR(100),                             -- 授权码
    code_expires_at TIMESTAMP,                      -- 授权码过期时间
    created_at TIMESTAMP DEFAULT NOW()
);
```

### 13.3 API接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/system/oauth2/clients | 客户端列表 |
| POST | /api/v1/system/oauth2/clients | 创建客户端 |
| PUT | /api/v1/system/oauth2/clients/:id | 更新客户端 |
| DELETE | /api/v1/system/oauth2/clients/:id | 删除客户端 |
| GET | /api/v1/system/oauth2/tokens | Token列表 |
| DELETE | /api/v1/system/oauth2/tokens/:id | 删除Token |
| DELETE | /api/v1/system/oauth2/tokens/user/:userId | 删除用户所有Token |
| POST | /oauth2/token | 获取AccessToken |
| POST | /oauth2/refresh | 刷新Token |
| DELETE | /oauth2/logout | 登出 |

### 13.4 业务规则

- 支持的授权模式：authorization_code（授权码）、client_credentials（客户端凭证）、refresh_token（刷新令牌）
- AccessToken默认有效期2小时，RefreshToken默认有效期7天
- 授权码有效期10分钟，只能使用一次
- 支持Token revoke（撤销）

---

## 14. 多车间管理

### 14.1 功能概述

多车间管理模块支持闻荫科技3个独立车间的数据隔离管理，包括机加工车间、电池车间、试制车间。

### 14.2 数据模型

```sql
-- 车间表
CREATE TABLE mdm_workshop (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    workshop_code VARCHAR(50) UNIQUE NOT NULL,      -- 车间编码
    workshop_name VARCHAR(100) NOT NULL,           -- 车间名称
    workshop_type VARCHAR(20) NOT NULL,            -- 车间类型: MACHINING/BATTERY/TRIAL
    erp_plant_code VARCHAR(50),                     -- ERP工厂编码
    max_devices INTEGER DEFAULT 0,                 -- 最大设备数
    max_workers INTEGER DEFAULT 0,                  -- 最大工人数
    max_capacity_per_day DECIMAL(10,2),            -- 最大日产能
    time_zone VARCHAR(30) DEFAULT 'Asia/Shanghai', -- 时区
    manager_id BIGINT,                             -- 负责人ID
    manager_name VARCHAR(50),                      -- 负责人姓名
    contact_phone VARCHAR(20),                     -- 联系电话
    address VARCHAR(200),                          -- 车间地址
    is_default INTEGER DEFAULT 0,                  -- 是否默认车间
    status INTEGER DEFAULT 1,                      -- 1启用/0禁用
    remark VARCHAR(500),
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 产线表
CREATE TABLE mdm_production_line (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    workshop_id BIGINT NOT NULL,                   -- 所属车间ID
    line_code VARCHAR(50) NOT NULL,                -- 产线编码
    line_name VARCHAR(100) NOT NULL,               -- 产线名称
    line_type VARCHAR(30),                         -- 产线类型: ASSEMBLY/PACKAGING/TEST
    max_stations INTEGER DEFAULT 0,                  -- 最大工位数
    status INTEGER DEFAULT 1,
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 工位表
CREATE TABLE mdm_station (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    workshop_id BIGINT NOT NULL,                   -- 所属车间ID
    line_id BIGINT NOT NULL,                       -- 所属产线ID
    station_code VARCHAR(50) NOT NULL,              -- 工位编码
    station_name VARCHAR(100) NOT NULL,            -- 工位名称
    station_type VARCHAR(30),                      -- 工位类型: PROCESS/QC/PACK
    process_code VARCHAR(50),                       -- 工序编码
    status INTEGER DEFAULT 1,
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 工厂日历表
CREATE TABLE aps_working_calendar (
    id BIGSERIAL PRIMARY KEY,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    workshop_id BIGINT NOT NULL,                   -- 车间ID
    calendar_name VARCHAR(100) NOT NULL,          -- 日历名称
    work_days VARCHAR(20),                        -- 工作日，如"1,2,3,4,5"
    shifts VARCHAR(100),                           -- 班次JSON
    holiday_dates TEXT,                            -- 节假日日期列表JSON
    special_work_dates TEXT,                       -- 特殊工作日期JSON
    effective_from DATE,                           -- 生效日期
    effective_to DATE,                             -- 失效日期
    status INTEGER DEFAULT 1,
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### 14.3 API接口

**车间管理API**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/mdm/workshops | 车间列表 |
| GET | /api/v1/mdm/workshops/:id | 获取车间详情 |
| POST | /api/v1/mdm/workshops | 创建车间 |
| PUT | /api/v1/mdm/workshops/:id | 更新车间 |
| DELETE | /api/v1/mdm/workshops/:id | 删除车间 |
| GET | /api/v1/mdm/workshops/:id/config | 获取车间配置 |
| PUT | /api/v1/mdm/workshops/:id/config | 更新车间配置 |

**产线管理API**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/mdm/lines | 产线列表 |
| GET | /api/v1/mdm/lines/:id | 获取产线详情 |
| POST | /api/v1/mdm/lines | 创建产线 |
| PUT | /api/v1/mdm/lines/:id | 更新产线 |
| DELETE | /api/v1/mdm/lines/:id | 删除产线 |

**工位管理API**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/mdm/stations | 工位列表 |
| GET | /api/v1/mdm/stations/:id | 获取工位详情 |
| POST | /api/v1/mdm/stations | 创建工位 |
| PUT | /api/v1/mdm/stations/:id | 更新工位 |
| DELETE | /api/v1/mdm/stations/:id | 删除工位 |

**工厂日历API**:

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/aps/calendar | 获取日历列表 |
| POST | /api/v1/aps/calendar | 创建日历 |
| PUT | /api/v1/aps/calendar/:id | 更新日历 |
| DELETE | /api/v1/aps/calendar/:id | 删除日历 |

### 14.4 业务规则

- 车间类型枚举：MACHINING（机加工）、BATTERY（电池）、TRIAL（试制）
- 所有业务表通过workshop_id字段关联到车间，实现数据隔离
- 用户可绑定多个车间，登录后可切换当前车间
- Token中携带当前车间上下文（workshop_id）
- 超级管理员可跨车间查看汇总数据
- 车间禁用后，该车间下的产线、工位自动禁用

---
