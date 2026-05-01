# Infra基础设施和System系统模块设计文档

## 1. 模块概述

### 1.1 Infra模块职责
Infra（Infrastructure）基础设施模块是SFMS3.0系统的基础支撑模块，提供企业级的通用基础设施服务，包括：
- 文件存储管理
- 定时任务调度
- API访问日志与错误日志
- 数据库文档生成
- 代码生成器
- 配置管理
- Redis缓存管理
- 测试演示功能

### 1.2 System模块职责
System系统模块是SFMS3.0系统的核心系统模块，提供企业级应用的基础系统功能，包括：
- 用户认证与授权（登录/登出/Token管理）
- 权限管理（菜单/角色/权限）
- 组织架构管理（部门/岗位）
- 字典管理
- 短信/邮件服务
- 通知公告
- OAuth2第三方登录
- 日志管理（登录日志/操作日志）
- 敏感词过滤
- 序列号生成器
- 错误码管理

## 2. 技术架构

### 2.1 模块结构对比

#### Infra模块结构
```
win-module-infra/
├── win-module-infra-api/          # API接口层
│   └── src/main/java/com/win/module/infra/
│       ├── api/                   # 对外API接口
│       │   └── logger/           # 日志API
│       ├── enums/                # 枚举常量
│       └── InfraApiApplication.java
└── win-module-infra-biz/         # 业务实现层
    └── src/main/java/com/win/module/infra/
        ├── controller/           # 控制器
        │   ├── codegen/         # 代码生成
        │   ├── config/          # 配置管理
        │   ├── db/              # 数据库
        │   ├── file/            # 文件存储
        │   ├── job/             # 定时任务
        │   ├── logger/          # 日志
        │   ├── redis/           # Redis
        │   ├── remark/          # 备注
        │   ├── test/            # 测试
        │   └── trends/          # 动态
        ├── service/             # 服务层
        ├── dal/                 # 数据访问层
        │   ├── dataobject/      # DO对象
        │   └── mysql/           # Mapper
        ├── convert/             # 对象转换
        └── enums/               # 业务枚举
```

#### System模块结构
```
win-module-system/
├── win-module-system-api/        # API接口层
│   └── src/main/java/com/win/module/system/api/
│       ├── dept/                # 部门API
│       ├── dict/                # 字典API
│       ├── errorcode/           # 错误码API
│       ├── logger/              # 日志API
│       ├── notify/              # 通知API
│       ├── oauth2/              # OAuth2 API
│       ├── permission/          # 权限API
│       ├── sensitiveword/        # 敏感词API
│       ├── serialnumber/         # 序列号API
│       ├── sms/                  # 短信API
│       └── user/                 # 用户API
└── win-module-system-biz/       # 业务实现层
    └── src/main/java/com/win/module/system/
        ├── controller/           # 控制器
        │   ├── auth/            # 认证
        │   ├── dept/             # 部门
        │   ├── dict/             # 字典
        │   ├── errorcode/        # 错误码
        │   ├── ip/               # IP地址
        │   ├── logger/           # 日志
        │   ├── mail/             # 邮件
        │   ├── notice/            # 公告
        │   ├── notify/            # 通知
        │   ├── oauth2/            # OAuth2
        │   └── permission/        # 权限
        ├── service/              # 服务层
        ├── dal/                   # 数据访问层
        ├── convert/               # 对象转换
        └── enums/                 # 业务枚举
```

## 3. Infra核心功能分析

### 3.1 文件存储服务 (File)
提供统一的文件上传、下载、管理功能。

**核心功能**:
- 文件上传（单文件/多文件）
- 文件下载（附件/在线显示）
- 文件删除（按ID/按关联表）
- 文件分页查询
- 文件配置管理

**关键API**:
```java
POST /infra/file/upload           # 上传文件
POST /infra/file/uploadFile       # 上传文件到硬盘
POST /infra/file/uploadFileData   # 上传文件数据
POST /infra/file/uploads          # 批量上传
DELETE /infra/file/delete         # 删除文件
GET  /infra/file/{configId}/get/** # 下载文件
GET  /infra/file/{configId}/show/** # 在线显示
GET  /infra/file/page             # 分页查询
GET  /infra/file/list             # 文件列表
```

**FileConfigDO数据结构**:
| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 主键 |
| name | String | 配置名称 |
| filePath | String | 存储路径 |
| fileType | String | 文件类型 |
| storageType | Integer | 存储类型 |

**FileDO数据结构**:
| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 主键 |
| configId | Long | 配置ID |
| name | String | 文件名 |
| path | String | 文件路径 |
| content | byte[] | 文件内容 |
| size | Long | 文件大小 |
| tableName | String | 关联表名 |
| tableId | Long | 关联表ID |

### 3.2 定时任务调度 (Job)
基于Quartz的分布式定时任务管理。

**核心功能**:
- 创建定时任务
- 更新定时任务
- 删除定时任务
- 触发任务执行
- 更新任务状态
- Cron表达式预览
- 任务执行日志

**关键API**:
```java
POST /infra/job/create            # 创建任务
PUT  /infra/job/update            # 更新任务
PUT  /infra/job/update-status     # 更新状态
DELETE /infra/job/delete          # 删除任务
PUT  /infra/job/trigger           # 触发任务
GET  /infra/job/get               # 获取任务
GET  /infra/job/list              # 任务列表
GET  /infra/job/page              # 分页查询
GET  /infra/job/export-excel      # 导出Excel
GET  /infra/job/get_next_times     # 下N次执行时间
```

**JobDO数据结构**:
| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 任务编号 |
| name | String | 任务名称 |
| status | Integer | 任务状态(枚举) |
| handlerName | String | 处理器名称 |
| handlerParam | String | 处理器参数 |
| cronExpression | String | CRON表达式 |
| retryCount | Integer | 重试次数 |
| retryInterval | Integer | 重试间隔(ms) |
| monitorTimeout | Integer | 监控超时(ms) |

### 3.3 API访问日志 (ApiAccessLog)
记录所有API请求的访问日志。

**关键API**:
```java
POST /infra/api-access-log/create  # 创建访问日志(内部)
GET  /infra/api-access-log/page   # 分页查询
```

**ApiAccessLogDO数据结构**:
| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 主键 |
| userId | Long | 用户编号 |
| userType | Integer | 用户类型 |
| module | String | 模块 |
| name | String | 操作名字 |
| reqMethod | String | 请求方法 |
| reqUrl | String | 请求地址 |
| reqParams | String | 请求参数 |
| ip | String | IP地址 |
| ua | String | 浏览器UA |
| startTime | LocalDateTime | 开始时间 |
| duration | Integer | 耗时(ms) |
| resultCode | Integer | 结果码 |
| resultMsg | String | 结果消息 |

### 3.4 API错误日志 (ApiErrorLog)
记录API请求中的异常信息。

### 3.5 代码生成器 (Codegen)
根据数据库表结构自动生成CRUD代码。

**核心功能**:
- 表结构读取
- 代码预览
- 代码生成
- 列信息配置

### 3.6 数据库管理 (DataSourceConfig)
管理多数据源配置。

### 3.7 配置管理 (Config)
系统配置参数管理。

### 3.8 Redis管理 (Redis)
提供Redis缓存的监控和管理。

### 3.9 数据库文档 (DatabaseDoc)
自动生成数据库文档。

## 4. System核心功能分析

### 4.1 认证授权 (Auth)
提供用户登录、登出、Token刷新等功能。

**关键API**:
```java
POST /system/auth/login           # 账号密码登录
POST /system/auth/loginNoCode     # 无验证码登录
POST /system/auth/logout          # 登出
POST /system/auth/refresh-token   # 刷新令牌
GET  /system/auth/get-permission-info # 获取权限信息
GET  /system/auth/public-key      # 获取RSA公钥
```

**登录流程**:
1. 获取RSA公钥
2. 使用公钥加密密码
3. 调用登录接口
4. 获取Token
5. 使用Token访问其他接口

### 4.2 权限管理 (Permission)
基于RBAC的权限控制模型。

**核心功能**:
- 菜单管理
- 角色管理
- 用户权限分配
- 数据权限控制

**PermissionApi接口**:
```java
Set<Long> getUserRoleIdListByRoleIds(Collection<Long> roleIds);
boolean hasAnyPermissions(Long userId, String... permissions);
boolean hasAnyRoles(Long userId, String... roles);
DeptDataPermissionRespDTO getDeptDataPermission(Long userId);
Object getPermissionIdentifications();
```

**MenuDO数据结构**:
| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 主键 |
| name | String | 菜单名称 |
| permission | String | 权限标识 |
| menuType | Integer | 菜单类型 |
| parentId | Long | 父菜单ID |
| path | String | 路由地址 |
| icon | String | 图标 |
| orderNum | Integer | 排序 |
| status | Integer | 状态 |

**RoleDO数据结构**:
| 字段 | 类型 | 说明 |
|------|------|------|
| id | Long | 主键 |
| name | String | 角色名称 |
| code | String | 角色标识 |
| sort | Integer | 排序 |
| dataScope | Integer | 数据范围 |
| status | Integer | 状态 |

### 4.3 部门管理 (Dept)
组织架构管理，支持树形结构。

**关键API**:
```java
POST /system/dept/create          # 创建部门
PUT  /system/dept/update         # 更新部门
DELETE /system/dept/delete       # 删除部门
GET  /system/dept/get            # 获取部门
GET  /system/dept/list           # 部门列表
GET  /system/dept/page           # 分页查询
GET  /system/dept/excel          # 导出Excel
```

### 4.4 岗位管理 (Post)
用户岗位管理。

### 4.5 字典管理 (Dict)
系统级数据字典。

**关键API**:
```java
# 字典类型
POST /system/dict/type/create    # 创建字典类型
PUT  /system/dict/type/update     # 更新字典类型
DELETE /system/dict/type/delete  # 删除字典类型
GET  /system/dict/type/get       # 获取字典类型
GET  /system/dict/type/page      # 分页查询
GET  /system/dict/type/export    # 导出
GET  /system/dict/type/get-all   # 获取所有类型

# 字典数据
POST /system/dict/data/create    # 创建字典数据
PUT  /system/dict/data/update    # 更新字典数据
DELETE /system/dict/data/delete  # 删除字典数据
GET  /system/dict/data/get       # 获取字典数据
GET  /system/dict/data/page      # 分页查询
GET  /system/dict/data/export    # 导出
GET  /system/dict/data/get-by-type # 根据类型获取
```

### 4.6 短信服务 (Sms)
短信发送和验证码管理。

**关键API**:
```java
POST /system/sms/send             # 发送短信
POST /system/sms/send-code       # 发送验证码
POST /system/sms/validate-code   # 校验验证码
```

### 4.7 邮件服务 (Mail)
邮件发送和管理。

**关键API**:
```java
# 邮件账号
POST /system/mail/account/create  # 创建账号
PUT  /system/mail/account/update # 更新账号
DELETE /system/mail/account/delete # 删除账号
GET  /system/mail/account/get    # 获取账号
GET  /system/mail/account/page   # 分页查询

# 邮件模板
POST /system/mail/template/create # 创建模板
PUT  /system/mail/template/update # 更新模板
DELETE /system/mail/template/delete # 删除模板
GET  /system/mail/template/get    # 获取模板
GET  /system/mail/template/page   # 分页查询
POST /system/mail/template/send  # 发送邮件
```

### 4.8 通知公告 (Notice/Notify)
系统通知和公告管理。

**NotifyMessageSendApi接口**:
```java
void sendSingleToUser(NotifySendSingleToUserReqDTO dto);
void sendSingleToRole(NotifySendSingleToRoleReqDTO dto);
void sendSingleToRoles(NotifySendSingleToRolesReqDTO dto);
```

### 4.9 OAuth2认证 (OAuth2)
支持第三方应用接入。

**关键API**:
```java
# Token管理
POST /oauth2/token               # 获取Token
DELETE /oauth2/token             # 删除Token
GET  /oauth2/token/page          # Token列表

# 客户端管理
POST /oauth2/client/create       # 创建客户端
PUT  /oauth2/client/update       # 更新客户端
DELETE /oauth2/client/delete     # 删除客户端
GET  /oauth2/client/get          # 获取客户端
GET  /oauth2/client/page         # 分页查询
```

### 4.10 日志管理 (Logger)
**登录日志 (LoginLog)**:
```java
POST /system/login-log/create     # 创建登录日志
GET  /system/login-log/page      # 分页查询
GET  /system/login-log/export    # 导出
```

**操作日志 (OperateLog)**:
记录用户的关键操作。

### 4.11 错误码管理 (ErrorCode)
统一管理业务错误码。

### 4.12 敏感词管理 (SensitiveWord)
文字敏感词过滤。

### 4.13 序列号生成器 (SerialNumber)
提供统一的序列号生成服务。

## 5. 权限设计

### 5.1 Infra模块权限
| 权限标识 | 说明 |
|----------|------|
| infra:file:create | 上传文件 |
| infra:file:update | 更新文件 |
| infra:file:delete | 删除文件 |
| infra:file:query | 查询文件 |
| infra:file:export | 导出文件 |
| infra:job:create | 创建任务 |
| infra:job:update | 更新任务 |
| infra:job:delete | 删除任务 |
| infra:job:query | 查询任务 |
| infra:job:trigger | 触发任务 |
| infra:job:export | 导出任务 |
| infra:codegen:query | 代码生成查询 |
| infra:codegen:preview | 代码预览 |
| infra:codegen:create | 生成代码 |
| infra:config:create | 创建配置 |
| infra:config:update | 更新配置 |
| infra:config:delete | 删除配置 |
| infra:config:query | 查询配置 |
| infra:config:export | 导出配置 |

### 5.2 System模块权限
| 权限标识 | 说明 |
|----------|------|
| system:user:create | 创建用户 |
| system:user:update | 更新用户 |
| system:user:delete | 删除用户 |
| system:user:query | 查询用户 |
| system:user:export | 导出用户 |
| system:user:update-password | 修改密码 |
| system:role:create | 创建角色 |
| system:role:update | 更新角色 |
| system:role:delete | 删除角色 |
| system:role:query | 查询角色 |
| system:role:permission | 角色权限 |
| system:menu:create | 创建菜单 |
| system:menu:update | 更新菜单 |
| system:menu:delete | 删除菜单 |
| system:menu:query | 查询菜单 |
| system:dept:create | 创建部门 |
| system:dept:update | 更新部门 |
| system:dept:delete | 删除部门 |
| system:dept:query | 查询部门 |

## 6. 数据流分析

### 6.1 用户登录流程
```
1. 前端获取公钥 (GET /system/auth/public-key)
2. 前端加密密码 (使用RSA公钥)
3. 提交登录 (POST /system/auth/login)
4. 后端验证密码
5. 生成Token和RefreshToken
6. 记录登录日志
7. 返回Token给前端
8. 前端携带Token访问受保护资源
```

### 6.2 权限校验流程
```
1. 请求携带Token
2. Security Filter解析Token
3. 获取用户ID和角色
4. 权限注解 @PreAuthorize("@ss.hasPermission('xxx')")
5. 调用PermissionService校验权限
6. 返回校验结果
```

### 6.3 文件上传流程
```
1. 前端调用上传接口 (POST /infra/file/upload)
2. Controller接收MultipartFile
3. FileService保存文件到配置路径
4. 保存FileDO元数据
5. 返回文件访问路径
```

### 6.4 定时任务执行流程
```
1. CronTrigger触发任务
2. JobExecutionContext传递上下文
3. 执行Handler处理器
4. 记录执行日志(JobLogDO)
5. 失败时重试(retryCount)
6. 超时监控(monitorTimeout)
```

## 7. 关键技术实现

### 7.1 OAuth2实现
使用Spring Security OAuth2实现，支持：
- password模式（账号密码）
- refresh_token模式（刷新令牌）
- client_credentials模式（客户端凭证）

### 7.2 Redis Token存储
使用Redis存储Token，支持：
- Token自动续期
- 强制登出
- 并发登录控制

### 7.3 Cron表达式
使用Quartz CronExpression，支持：
- 标准Cron格式
- 下N次执行时间预览
- 动态修改执行周期

### 7.4 RSA密码加密
前端使用RSA公钥加密密码传输，保护用户密码安全。

## 8. 表结构汇总

### 8.1 Infra模块表
| 表名 | 说明 |
|------|------|
| infra_file_config | 文件配置表 |
| infra_file | 文件表 |
| infra_job | 定时任务表 |
| infra_job_log | 任务执行日志表 |
| infra_api_access_log | API访问日志表 |
| infra_api_error_log | API错误日志表 |
| infra_codegen_table | 代码生成表定义表 |
| infra_codegen_column | 代码生成列定义表 |
| infra_db_data_source_config | 数据源配置表 |
| infra_config | 系统配置表 |

### 8.2 System模块表
| 表名 | 说明 |
|------|------|
| system_user | 用户表 |
| system_dept | 部门表 |
| system_post | 岗位表 |
| system_menu | 菜单表 |
| system_role | 角色表 |
| system_user_role | 用户角色关联表 |
| system_role_menu | 角色菜单关联表 |
| system_dict_type | 字典类型表 |
| system_dict_data | 字典数据表 |
| system_oauth2_token | OAuth2 Token表 |
| system_oauth2_client | OAuth2客户端表 |
| system_sms_code | 短信验证码表 |
| system_sms_log | 短信日志表 |
| system_mail_account | 邮件账号表 |
| system_mail_template | 邮件模板表 |
| system_mail_log | 邮件日志表 |
| system_notice | 通知公告表 |
| system_notify_message | 通知消息表 |
| system_login_log | 登录日志表 |
| system_operation_log | 操作日志表 |
| system_sensitive_word | 敏感词表 |
| system_error_code | 错误码表 |
| system_serial_number | 序列号表 |
| system_area | 地区表 |

## 9. 外部依赖关系

### 9.1 Infra依赖
- MySQL: 文件存储、任务日志
- Redis: 分布式锁（任务调度）
- 文件系统: 本地存储/S3兼容存储

### 9.2 System依赖
- Infra: 日志服务
- Redis: Token存储、缓存
- 第三方服务: 短信网关、邮件SMTP

### 9.3 被依赖关系
- 其他业务模块都依赖System进行认证授权
- 其他业务模块可使用Infra的文件服务、定时任务
