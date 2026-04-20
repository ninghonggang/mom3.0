# MOM3.0前端_SYSTEM系统管理模块设计文档

> **版本**: V1.0
> **日期**: 2026-04-17
> **项目**: 闻荫科技MOM3.0智能制造执行系统
> **对比版本**: SFMS3.0 SYSTEM系统管理模块 (22页缺失)

---

## 1. 模块概述

### 1.1 功能定位

SYSTEM系统管理模块是MOM3.0的基础模块，负责系统配置、用户权限、安全管理等核心基础设施功能。

### 1.2 当前状态 vs SFMS3.0

| 类别 | MOM3.0现状 | SFMS3.0完整功能 |
|------|-----------|----------------|
| 页面数量 | 13页 | 35+页 |
| 租户套餐 | 缺失 | 套餐管理、套餐配置、套餐权限 |
| 区域管理 | 缺失 | 行政区域、区域配置 |
| 短信管理 | 缺失 | 短信渠道、短信模板 |
| 邮件管理 | 缺失 | 邮件账号、邮件模板 |
| 消息通知 | 缺失 | 消息类型、消息模板 |
| OAuth2 | 缺失 | 客户端配置、授权管理 |
| Token管理 | 缺失 | Token配置、Token刷新 |
| 敏感词 | 缺失 | 敏感词配置、过滤规则 |
| 错误码 | 缺失 | 错误码配置、错误处理 |
| 序列号 | 缺失 | 序列号规则、序列号生成 |
| 密码规则 | 缺失 | 密码策略、复杂度配置 |
| 安全配置 | 缺失 | 安全策略、加密配置 |
| IP白名单 | 缺失 | IP配置、访问控制 |
| 操作限制 | 缺失 | 限制规则、限流配置 |
| 数据权限 | 缺失 | 权限配置、权限组 |
| 字段权限 | 缺失 | 字段配置、字段隐藏 |
| API权限 | 缺失 | 接口配置、权限校验 |

---

## 2. 缺失页面详细设计

### 2.1 租户套餐 (`/system/tenantPackage`)

**路径**: `/system/tenantPackage`
**组件**: `TenantPackageList.vue`
**功能**: 租户套餐管理、套餐配置、套餐权限

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| packageId | String | 套餐ID |
| packageName | String | 套餐名称 |
| packageCode | String | 套餐编码 |
| packageType | String | 套餐类型(BASIC标准/PRO专业/ENTERPRISE企业) |
| price | Decimal | 价格 |
| duration | Integer | 时长(天) |
| userLimit | Integer | 用户数限制 |
| storageLimit | Integer | 存储空间限制(G) |
| apiLimit | Integer | API调用限制(次/天) |
| features | JSON | 功能列表 |
| status | String | 状态(ENABLE/DISABLE) |
| remark | String | 备注 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /system/tenantPackage/create | 创建套餐 |
| PUT | /system/tenantPackage/update | 更新套餐 |
| DELETE | /system/tenantPackage/delete | 删除套餐 |
| GET | /system/tenantPackage/get | 获取套餐详情 |
| GET | /system/tenantPackage/page | 套餐分页查询 |
| PUT | /system/tenantPackage/enable | 启用套餐 |
| PUT | /system/tenantPackage/disable | 禁用套餐 |

---

### 2.2 区域管理 (`/system/region`)

**路径**: `/system/region`
**组件**: `RegionList.vue`
**功能**: 行政区域管理、区域配置

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| regionId | String | 区域ID |
| parentId | String | 父区域ID |
| regionCode | String | 区域编码 |
| regionName | String | 区域名称 |
| regionType | String | 区域类型(PROVINCE省/CITY市/DISTRICT区) |
| longitude | Decimal | 经度 |
| latitude | Decimal | 纬度 |
| sort | Integer | 排序 |
| status | String | 状态 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /system/region/create | 创建区域 |
| PUT | /system/region/update | 更新区域 |
| DELETE | /system/region/delete | 删除区域 |
| GET | /system/region/get | 获取区域详情 |
| GET | /system/region/tree | 获取区域树 |
| GET | /system/region/children | 获取子区域 |
| GET | /system/region/page | 区域分页查询 |

---

### 2.3 短信管理 (`/system/sms`)

**路径**: `/system/sms`
**组件**: `SmsList.vue`
**功能**: 短信渠道配置、短信发送管理

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| smsId | String | 短信ID |
| provider | String | 运营商(ALIYUN/TENCENT/HUAWEI) |
| appId | String | 应用ID |
| appKey | String | 应用密钥 |
| signature | String | 签名 |
| status | String | 状态(ENABLE/DISABLE) |
| remark | String | 备注 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /system/sms/create | 创建短信配置 |
| PUT | /system/sms/update | 更新短信配置 |
| DELETE | /system/sms/delete | 删除短信配置 |
| GET | /system/sms/get | 获取短信配置 |
| GET | /system/sms/page | 短信配置分页查询 |
| POST | /system/sms/send | 发送短信 |
| POST | /system/sms/test | 测试短信发送 |

---

### 2.4 短信模板 (`/system/smsTemplate`)

**路径**: `/system/smsTemplate`
**组件**: `SmsTemplateList.vue`
**功能**: 短信模板管理、变量配置

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| templateId | String | 模板ID |
| templateCode | String | 模板编码 |
| templateName | String | 模板名称 |
| content | String | 模板内容 |
| variables | JSON | 变量列表 |
| status | String | 状态(ENABLE/DISABLE) |
| remark | String | 备注 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /system/smsTemplate/create | 创建模板 |
| PUT | /system/smsTemplate/update | 更新模板 |
| DELETE | /system/smsTemplate/delete | 删除模板 |
| GET | /system/smsTemplate/get | 获取模板详情 |
| GET | /system/smsTemplate/page | 模板分页查询 |

---

### 2.5 邮件管理 (`/system/mail`)

**路径**: `/system/mail`
**组件**: `MailList.vue`
**功能**: 邮件账号配置、邮件发送管理

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| mailId | String | 邮件ID |
| provider | String | 运营商(SMTP/EXCHANGE) |
| host | String | 服务器地址 |
| port | Integer | 端口 |
| username | String | 用户名 |
| password | String | 密码 |
| fromAddress | String | 发件人地址 |
| fromName | String | 发件人名称 |
| useSsl | Boolean | 是否使用SSL |
| status | String | 状态(ENABLE/DISABLE) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /system/mail/create | 创建邮件配置 |
| PUT | /system/mail/update | 更新邮件配置 |
| DELETE | /system/mail/delete | 删除邮件配置 |
| GET | /system/mail/get | 获取邮件配置 |
| GET | /system/mail/page | 邮件配置分页查询 |
| POST | /system/mail/send | 发送邮件 |
| POST | /system/mail/test | 测试邮件发送 |

---

### 2.6 邮件模板 (`/system/mailTemplate`)

**路径**: `/system/mailTemplate`
**组件**: `MailTemplateList.vue`
**功能**: 邮件模板管理、变量配置

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| templateId | String | 模板ID |
| templateCode | String | 模板编码 |
| templateName | String | 模板名称 |
| subject | String | 邮件主题 |
| content | String | 模板内容(HTML) |
| variables | JSON | 变量列表 |
| status | String | 状态(ENABLE/DISABLE) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /system/mailTemplate/create | 创建模板 |
| PUT | /system/mailTemplate/update | 更新模板 |
| DELETE | /system/mailTemplate/delete | 删除模板 |
| GET | /system/mailTemplate/get | 获取模板详情 |
| GET | /system/mailTemplate/page | 模板分页查询 |

---

### 2.7 消息通知 (`/system/message`)

**路径**: `/system/message`
**组件**: `MessageList.vue`
**功能**: 消息类型配置、消息模板管理

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| messageId | String | 消息ID |
| messageType | String | 消息类型(SYSTEM订单/告警/审批) |
| title | String | 消息标题 |
| content | String | 消息内容 |
| templateId | String | 模板ID |
| channel | String | 发送渠道(SMS/MAIL/WEBSOCKET) |
| status | String | 状态(ENABLE/DISABLE) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /system/message/create | 创建消息 |
| PUT | /system/message/update | 更新消息 |
| DELETE | /system/message/delete | 删除消息 |
| GET | /system/message/get | 获取消息详情 |
| GET | /system/message/page | 消息分页查询 |
| POST | /system/message/send | 发送消息 |

---

### 2.8 OAuth2客户端 (`/system/oauth2`)

**路径**: `/system/oauth2`
**组件**: `OAuth2List.vue`
**功能**: OAuth2客户端配置、授权管理

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| clientId | String | 客户端ID |
| clientSecret | String | 客户端密钥 |
| clientName | String | 客户端名称 |
| grantTypes | String | 授权类型(AUTHORIZATION_CODE/REFRESH_TOKEN/CLIENT_CREDENTIALS) |
| redirectUri | String | 回调地址 |
| scopes | String | 权限范围 |
| accessTokenValidity | Integer | AccessToken有效期(秒) |
| refreshTokenValidity | Integer | RefreshToken有效期(秒) |
| status | String | 状态(ENABLE/DISABLE) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /system/oauth2/create | 创建客户端 |
| PUT | /system/oauth2/update | 更新客户端 |
| DELETE | /system/oauth2/delete | 删除客户端 |
| GET | /system/oauth2/get | 获取客户端详情 |
| GET | /system/oauth2/page | 客户端分页查询 |
| POST | /system/oauth2/resetSecret | 重置密钥 |

---

### 2.9 Token管理 (`/system/token`)

**路径**: `/system/token`
**组件**: `TokenList.vue`
**功能**: Token配置、Token刷新管理

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| tokenId | String | TokenID |
| userId | String | 用户ID |
| userName | String | 用户名 |
| token | String | Token值 |
| tokenType | String | Token类型(ACCESS/REFRESH) |
| deviceId | String | 设备ID |
| deviceType | String | 设备类型(WEB/APP/IOS/ANDROID) |
| expireTime | DateTime | 过期时间 |
| createTime | DateTime | 创建时间 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /system/token/page | Token分页查询 |
| DELETE | /system/token/revoke | 撤销Token |
| DELETE | /system/token/revokeAll | 撤销所有Token |
| POST | /system/token/refresh | 刷新Token |

---

### 2.10 敏感词 (`/system/sensitiveWord`)

**路径**: `/system/sensitiveWord`
**组件**: `SensitiveWordList.vue`
**功能**: 敏感词配置、过滤规则

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| wordId | String | 词汇ID |
| word | String | 敏感词 |
| wordType | String | 词汇类型(政治/色情/暴力/广告) |
| replaceWord | String | 替换词 |
| level | Integer | 级别(1轻度/2中度/3重度) |
| status | String | 状态(ENABLE/DISABLE) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /system/sensitiveWord/create | 创建敏感词 |
| PUT | /system/sensitiveWord/update | 更新敏感词 |
| DELETE | /system/sensitiveWord/delete | 删除敏感词 |
| GET | /system/sensitiveWord/get | 获取敏感词详情 |
| GET | /system/sensitiveWord/page | 敏感词分页查询 |
| POST | /system/sensitiveWord/check | 检查敏感词 |
| POST | /system/sensitiveWord/batchImport | 批量导入 |

---

### 2.11 错误码 (`/system/errorCode`)

**路径**: `/system/errorCode`
**组件**: `ErrorCodeList.vue`
**功能**: 错误码配置、错误处理

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| codeId | String | 错误码ID |
| errorCode | String | 错误码 |
| errorType | String | 错误类型(SYSTEM/BUSINESS) |
| module | String | 模块 |
| message | String | 错误信息 |
| solution | String | 解决方案 |
| httpStatus | Integer | HTTP状态码 |
| status | String | 状态(ENABLE/DISABLE) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /system/errorCode/create | 创建错误码 |
| PUT | /system/errorCode/update | 更新错误码 |
| DELETE | /system/errorCode/delete | 删除错误码 |
| GET | /system/errorCode/get | 获取错误码详情 |
| GET | /system/errorCode/page | 错误码分页查询 |
| GET | /system/errorCode/getByCode | 根据错误码查询 |

---

### 2.12 序列号 (`/system/sequence`)

**路径**: `/system/sequence`
**组件**: `SequenceList.vue`
**功能**: 序列号规则配置、序列号生成

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| sequenceId | String | 序列号ID |
| sequenceCode | String | 序列号编码 |
| sequenceName | String | 序列号名称 |
| prefix | String | 前缀 |
| dateFormat | String | 日期格式(YYYYMMDD) |
| startValue | Integer | 起始值 |
| currentValue | Integer | 当前值 |
| step | Integer | 步长 |
| length | Integer | 序号长度 |
| suffix | String | 后缀 |
| status | String | 状态(ENABLE/DISABLE) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /system/sequence/create | 创建序列号 |
| PUT | /system/sequence/update | 更新序列号 |
| DELETE | /system/sequence/delete | 删除序列号 |
| GET | /system/sequence/get | 获取序列号详情 |
| GET | /system/sequence/page | 序列号分页查询 |
| GET | /system/sequence/generate | 生成序列号 |

---

### 2.13 密码规则 (`/system/passwordRule`)

**路径**: `/system/passwordRule`
**组件**: `PasswordRuleList.vue`
**功能**: 密码策略配置、复杂度规则

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| ruleId | String | 规则ID |
| ruleName | String | 规则名称 |
| minLength | Integer | 最小长度 |
| maxLength | Integer | 最大长度 |
| requireUppercase | Boolean | 必须包含大写字母 |
| requireLowercase | Boolean | 必须包含小写字母 |
| requireDigit | Boolean | 必须包含数字 |
| requireSpecial | Boolean | 必须包含特殊字符 |
| notAllowUsername | Boolean | 不允许包含用户名 |
| notAllowRepeat | Boolean | 不允许连续重复 |
| passwordHistory | Integer | 密码历史记录数 |
| expireDays | Integer | 密码有效期(天) |
| status | String | 状态(ENABLE/DISABLE) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /system/passwordRule/create | 创建规则 |
| PUT | /system/passwordRule/update | 更新规则 |
| DELETE | /system/passwordRule/delete | 删除规则 |
| GET | /system/passwordRule/get | 获取规则详情 |
| GET | /system/passwordRule/page | 规则分页查询 |
| POST | /system/passwordRule/validate | 验证密码 |

---

### 2.14 IP白名单 (`/system/ipWhitelist`)

**路径**: `/system/ipWhitelist`
**组件**: `IpWhitelistList.vue`
**功能**: IP配置、访问控制

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| ipId | String | IP ID |
| ipAddress | String | IP地址 |
| ipType | String | IP类型(SINGLE单IP/RANGE范围/CIDR网段) |
| startIp | String | 起始IP |
| endIp | String | 结束IP |
| description | String | 描述 |
| status | String | 状态(ENABLE/DISABLE) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /system/ipWhitelist/create | 创建白名单 |
| PUT | /system/ipWhitelist/update | 更新白名单 |
| DELETE | /system/ipWhitelist/delete | 删除白名单 |
| GET | /system/ipWhitelist/get | 获取白名单详情 |
| GET | /system/ipWhitelist/page | 白名单分页查询 |
| POST | /system/ipWhitelist/check | 检查IP |

---

### 2.15 操作限制 (`/system/operLimit`)

**路径**: `/system/operLimit`
**组件**: `OperLimitList.vue`
**功能**: 限制规则配置、限流控制

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| limitId | String | 限制ID |
| limitType | String | 限制类型(API接口/登录/操作) |
| limitKey | String | 限制键 |
| limitValue | Integer | 限制值 |
| windowSeconds | Integer | 时间窗口(秒) |
| blockSeconds | Integer | 封禁时长(秒) |
| description | String | 描述 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /system/operLimit/create | 创建限制规则 |
| PUT | /system/operLimit/update | 更新限制规则 |
| DELETE | /system/operLimit/delete | 删除限制规则 |
| GET | /system/operLimit/get | 获取限制规则详情 |
| GET | /system/operLimit/page | 限制规则分页查询 |

---

### 2.16 数据权限 (`/system/dataPerm`)

**路径**: `/system/dataPerm`
**组件**: `DataPermList.vue`
**功能**: 数据权限配置、权限组

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| permId | String | 权限ID |
| permName | String | 权限名称 |
| permCode | String | 权限编码 |
| resourceType | String | 资源类型(ORG部门/WAREHOUSE仓库/LINE产线) |
| permRules | JSON | 权限规则 |
| status | String | 状态(ENABLE/DISABLE) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /system/dataPerm/create | 创建数据权限 |
| PUT | /system/dataPerm/update | 更新数据权限 |
| DELETE | /system/dataPerm/delete | 删除数据权限 |
| GET | /system/dataPerm/get | 获取数据权限详情 |
| GET | /system/dataPerm/page | 数据权限分页查询 |
| GET | /system/dataPerm/tree | 获取权限树 |

---

### 2.17 字段权限 (`/system/fieldPerm`)

**路径**: `/system/fieldPerm`
**组件**: `FieldPermList.vue`
**功能**: 字段权限配置、字段隐藏

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| fieldPermId | String | 字段权限ID |
| menuId | String | 菜单ID |
| fieldName | String | 字段名称 |
| fieldLabel | String | 字段标签 |
| visible | Boolean | 是否可见 |
| editable | Boolean | 是否可编辑 |
| required | Boolean | 是否必填 |
| description | String | 描述 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /system/fieldPerm/create | 创建字段权限 |
| PUT | /system/fieldPerm/update | 更新字段权限 |
| DELETE | /system/fieldPerm/delete | 删除字段权限 |
| GET | /system/fieldPerm/get | 获取字段权限详情 |
| GET | /system/fieldPerm/page | 字段权限分页查询 |
| GET | /system/fieldPerm/listByMenu | 根据菜单获取字段权限 |

---

### 2.18 API权限 (`/system/apiPerm`)

**路径**: `/system/apiPerm`
**组件**: `ApiPermList.vue`
**功能**: 接口权限配置、权限校验

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| apiPermId | String | API权限ID |
| apiPath | String | API路径 |
| apiMethod | String | 请求方法(GET/POST/PUT/DELETE) |
| apiName | String | API名称 |
| requirePerm | Boolean | 是否需要权限 |
| permCode | String | 权限编码 |
| description | String | 描述 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /system/apiPerm/create | 创建API权限 |
| PUT | /system/apiPerm/update | 更新API权限 |
| DELETE | /system/apiPerm/delete | 删除API权限 |
| GET | /system/apiPerm/get | 获取API权限详情 |
| GET | /system/apiPerm/page | API权限分页查询 |
| GET | /system/apiPerm/listByRole | 根据角色获取API权限 |

---

## 3. 页面清单汇总

| 序号 | 页面名称 | 路径 | 功能说明 |
|-----|---------|------|---------|
| 1 | 租户套餐 | /system/tenantPackage | 套餐管理 |
| 2 | 区域管理 | /system/region | 行政区域 |
| 3 | 短信管理 | /system/sms | 短信渠道 |
| 4 | 短信模板 | /system/smsTemplate | 短信模板 |
| 5 | 邮件管理 | /system/mail | 邮件配置 |
| 6 | 邮件模板 | /system/mailTemplate | 邮件模板 |
| 7 | 消息通知 | /system/message | 消息类型 |
| 8 | OAuth2客户端 | /system/oauth2 | OAuth配置 |
| 9 | Token管理 | /system/token | Token管理 |
| 10 | 敏感词 | /system/sensitiveWord | 敏感词过滤 |
| 11 | 错误码 | /system/errorCode | 错误码配置 |
| 12 | 序列号 | /system/sequence | 序列号管理 |
| 13 | 密码规则 | /system/passwordRule | 密码策略 |
| 14 | IP白名单 | /system/ipWhitelist | IP访问控制 |
| 15 | 操作限制 | /system/operLimit | 限流配置 |
| 16 | 数据权限 | /system/dataPerm | 数据权限 |
| 17 | 字段权限 | /system/fieldPerm | 字段权限 |
| 18 | API权限 | /system/apiPerm | API权限 |

---

*文档版本: V1.0 | 生成日期: 2026-04-17*
