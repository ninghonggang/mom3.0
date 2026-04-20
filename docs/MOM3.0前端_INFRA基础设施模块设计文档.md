# MOM3.0前端_INFRA基础设施模块设计文档

> **版本**: V1.0
> **日期**: 2026-04-17
> **项目**: 闻荫科技MOM3.0智能制造执行系统
> **对比版本**: SFMS3.0 INFRA基础设施模块 (12+页完全缺失)

---

## 1. 模块概述

### 1.1 功能定位

INFRA基础设施模块是MOM3.0的核心支撑模块，负责文件管理、代码生成、任务调度、系统监控等基础设施功能。本模块在MOM3.0中完全缺失。

### 1.2 当前状态 vs SFMS3.0

| 类别 | MOM3.0现状 | SFMS3.0完整功能 |
|------|-----------|----------------|
| 页面数量 | 0页 | 12+页 |
| 文件管理 | 缺失 | 文件上传、下载、预览 |
| 文件配置 | 缺失 | 存储配置、路径配置 |
| 代码生成器 | 缺失 | 模板配置、生成规则、代码预览 |
| 数据源配置 | 缺失 | 数据源管理、连接测试 |
| 数据库文档 | 缺失 | 文档生成、文档管理 |
| 定时任务 | 缺失 | 任务配置、任务调度、任务监控 |
| 定时日志 | 缺失 | 日志查看、日志分析 |
| Redis管理 | 缺失 | 缓存配置、缓存监控、缓存清理 |
| API日志 | 缺失 | 日志查看、日志统计 |
| 错误日志 | 缺失 | 错误记录、错误分析 |
| Druid监控 | 缺失 | 连接监控、SQL监控 |
| Swagger文档 | 缺失 | 文档配置、文档分组 |

---

## 2. 缺失页面详细设计

### 2.1 文件管理 (`/infra/file`)

**路径**: `/infra/file`
**组件**: `FileList.vue`
**功能**: 文件上传、下载、预览管理

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| fileId | String | 文件ID |
| fileName | String | 文件名称 |
| filePath | String | 文件路径 |
| fileSize | Long | 文件大小(字节) |
| fileType | String | 文件类型 |
| mimeType | String | MIME类型 |
| storageType | String | 存储类型(LOCAL/OSS/S3) |
| bucket | String | 存储桶 |
| uploadTime | DateTime | 上传时间 |
| uploadUserId | String | 上传用户ID |
| uploadUserName | String | 上传用户名 |
| downloadCount | Integer | 下载次数 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /infra/file/upload | 上传文件 |
| POST | /infra/file/uploadChunk | 分片上传 |
| POST | /infra/file/mergeChunk | 合并分片 |
| GET | /infra/file/download/{id} | 下载文件 |
| GET | /infra/file/preview/{id} | 预览文件 |
| DELETE | /infra/file/delete | 删除文件 |
| GET | /infra/file/get | 获取文件详情 |
| GET | /infra/file/page | 文件分页查询 |

---

### 2.2 文件配置 (`/infra/fileConfig`)

**路径**: `/infra/fileConfig`
**组件**: `FileConfigList.vue`
**功能**: 文件存储配置、路径配置

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| configId | String | 配置ID |
| configKey | String | 配置键 |
| configName | String | 配置名称 |
| storageType | String | 存储类型(LOCAL/OSS/S3/MINIO) |
| endpoint | String | 访问端点 |
| accessKey | String | AccessKey |
| secretKey | String | SecretKey |
| bucket | String | 存储桶 |
| basePath | String | 基础路径 |
| cdnUrl | String | CDN地址 |
| maxFileSize | Long | 最大文件大小 |
| allowedExtensions | String | 允许的扩展名 |
| status | String | 状态(ENABLE/DISABLE) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /infra/fileConfig/create | 创建配置 |
| PUT | /infra/fileConfig/update | 更新配置 |
| DELETE | /infra/fileConfig/delete | 删除配置 |
| GET | /infra/fileConfig/get | 获取配置详情 |
| GET | /infra/fileConfig/page | 配置分页查询 |
| POST | /infra/fileConfig/test | 测试连接 |

---

### 2.3 代码生成器 (`/infra/codegen`)

**路径**: `/infra/codegen`
**组件**: `CodegenList.vue`
**功能**: CRUD代码生成、模板配置

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| tableId | String | 表ID |
| tableName | String | 表名称 |
| tableComment | String | 表注释 |
| entityName | String | 实体名称 |
| packageName | String | 包名 |
| moduleName | String | 模块名 |
| businessName | String | 业务名 |
| functionName | String | 功能名称 |
| tablePrefix | String | 表前缀 |
| generateType | String | 生成类型(SINGLE单表/TREE树表) |
| template | String | 模板 |
| status | String | 状态(GENERATED未生成/GENERATING生成中/DONE已生成) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /infra/codegen/importTable | 导入表 |
| POST | /infra/codegen/batchImport | 批量导入表 |
| GET | /infra/codegen/tableList | 表列表 |
| GET | /infra/codegen/get/{tableId} | 获取表详情 |
| PUT | /infra/codegen/update | 更新表配置 |
| DELETE | /infra/codegen/delete/{tableId} | 删除表 |
| POST | /infra/codegen/generate | 生成代码 |
| GET | /infra/codegen/preview | 预览代码 |
| GET | /infra/codegen/download/{tableId} | 下载代码 |

---

### 2.4 代码模板 (`/infra/codegenTemplate`)

**路径**: `/infra/codegenTemplate`
**组件**: `CodegenTemplateList.vue`
**功能**: 代码生成模板配置

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| templateId | String | 模板ID |
| templateName | String | 模板名称 |
| templateCode | String | 模板编码 |
| templateType | String | 模板类型(ENTITY/VIEW/SERVICE/CONTROLLER) |
| templateContent | Text | 模板内容 |
| outputPath | String | 输出路径 |
| fileName | String | 文件名模板 |
| sort | Integer | 排序 |
| status | String | 状态(ENABLE/DISABLE) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /infra/codegenTemplate/create | 创建模板 |
| PUT | /infra/codegenTemplate/update | 更新模板 |
| DELETE | /infra/codegenTemplate/delete | 删除模板 |
| GET | /infra/codegenTemplate/get | 获取模板详情 |
| GET | /infra/codegenTemplate/page | 模板分页查询 |
| GET | /infra/codegenTemplate/listByType | 根据类型获取模板 |

---

### 2.5 数据源配置 (`/infra/datasource`)

**路径**: `/infra/datasource`
**组件**: `DatasourceList.vue`
**功能**: 多数据源配置、连接管理

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| dsId | String | 数据源ID |
| dsName | String | 数据源名称 |
| dsCode | String | 数据源编码 |
| dsType | String | 数据源类型(MYSQL/POSTGRESQL/ORACLE/SQLSERVER) |
| host | String | 主机地址 |
| port | Integer | 端口 |
| databaseName | String | 数据库名 |
| username | String | 用户名 |
| password | String | 密码 |
| url | String | JDBC URL |
| minPoolSize | Integer | 最小连接数 |
| maxPoolSize | Integer | 最大连接数 |
| timeout | Integer | 连接超时(毫秒) |
| status | String | 状态(ENABLE/DISABLE) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /infra/datasource/create | 创建数据源 |
| PUT | /infra/datasource/update | 更新数据源 |
| DELETE | /infra/datasource/delete | 删除数据源 |
| GET | /infra/datasource/get | 获取数据源详情 |
| GET | /infra/datasource/page | 数据源分页查询 |
| POST | /infra/datasource/test | 测试连接 |
| GET | /infra/datasource/tables | 获取数据表列表 |

---

### 2.6 数据库文档 (`/infra/dbdoc`)

**路径**: `/infra/dbdoc`
**组件**: `DbdocList.vue`
**功能**: 数据库文档生成、文档管理

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| docId | String | 文档ID |
| docName | String | 文档名称 |
| dsId | String | 数据源ID |
| content | Text | 文档内容 |
| version | String | 版本号 |
| createTime | DateTime | 创建时间 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /infra/dbdoc/generate | 生成文档 |
| GET | /infra/dbdoc/get/{docId} | 获取文档详情 |
| GET | /infra/dbdoc/page | 文档分页查询 |
| DELETE | /infra/dbdoc/delete | 删除文档 |
| GET | /infra/dbdoc/export/{docId} | 导出文档 |

---

### 2.7 定时任务 (`/infra/job`)

**路径**: `/infra/job`
**组件**: `JobList.vue`
**功能**: 任务调度配置、任务监控

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| jobId | String | 任务ID |
| jobName | String | 任务名称 |
| jobGroup | String | 任务组 |
| jobType | String | 任务类型(HTTP/DUBBO/SPRING_BEAN/SHELL) |
| cronExpression | String | Cron表达式 |
| url | String | 调用URL |
| method | String | 调用方法 |
| param | String | 调用参数 |
| description | String | 描述 |
| misfirePolicy | String | 失火策略(DO_NOTHING/FIRE_NOW) |
| status | String | 状态(NORMAL暂停/RUNNING运行) |
| lastRunTime | DateTime | 上次运行时间 |
| nextRunTime | DateTime | 下次运行时间 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| POST | /infra/job/create | 创建任务 |
| PUT | /infra/job/update | 更新任务 |
| DELETE | /infra/job/delete | 删除任务 |
| GET | /infra/job/get | 获取任务详情 |
| GET | /infra/job/page | 任务分页查询 |
| POST | /infra/job/start | 启动任务 |
| POST | /infra/job/stop | 停止任务 |
| POST | /infra/job/runOnce | 立即执行一次 |
| GET | /infra/job/listByGroup | 根据任务组获取任务 |

---

### 2.8 定时日志 (`/infra/jobLog`)

**路径**: `/infra/jobLog`
**组件**: `JobLogList.vue`
**功能**: 任务执行日志查看

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| logId | String | 日志ID |
| jobId | String | 任务ID |
| jobName | String | 任务名称 |
| jobGroup | String | 任务组 |
| invokeTarget | String | 调用目标 |
| startTime | DateTime | 开始时间 |
| endTime | DateTime | 结束时间 |
| duration | Long | 执行时长(毫秒) |
| status | String | 执行状态(SUCCESS成功/FAILURE失败) |
| exception | Text | 异常信息 |
| remark | String | 备注 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /infra/jobLog/page | 日志分页查询 |
| GET | /infra/jobLog/get/{logId} | 获取日志详情 |
| DELETE | /infra/jobLog/delete | 删除日志 |
| DELETE | /infra/jobLog/deleteByJobId/{jobId} | 删除任务日志 |
| GET | /infra/jobLog/statistics | 日志统计 |

---

### 2.9 Redis管理 (`/infra/redis`)

**路径**: `/infra/redis`
**组件**: `RedisList.vue`
**功能**: Redis缓存配置、监控、清理

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| key | String | 缓存键 |
| value | String | 缓存值 |
| type | String | 数据类型(STRING/LIST/HASH/SET/ZSET) |
| ttl | Long | 过期时间(秒) |
| size | Long | 大小(字节) |
| hitRate | Decimal | 命中率 |
| createTime | DateTime | 创建时间 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /infra/redis/keys | 获取键列表 |
| GET | /infra/redis/get | 获取值 |
| PUT | /infra/redis/set | 设置值 |
| DELETE | /infra/redis/delete | 删除键 |
| DELETE | /infra/redis/deleteByPattern | 按模式删除 |
| POST | /infra/redis/expire | 设置过期时间 |
| GET | /infra/redis/info | 获取Redis信息 |
| GET | /infra/redis/stats | 获取统计信息 |
| POST | /infra/redis/flushdb | 清空当前数据库 |

---

### 2.10 API日志 (`/infra/apiLog`)

**路径**: `/infra/apiLog`
**组件**: `ApiLogList.vue`
**功能**: API访问日志查看、统计

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| logId | String | 日志ID |
| traceId | String | 跟踪ID |
| userId | String | 用户ID |
| userName | String | 用户名 |
| method | String | 请求方法 |
| url | String | 请求URL |
| queryParams | String | 查询参数 |
| requestBody | Text | 请求体 |
| responseBody | Text | 响应体 |
| statusCode | Integer | 状态码 |
| duration | Long | 执行时长(毫秒) |
| ip | String | IP地址 |
| userAgent | String | UserAgent |
| createTime | DateTime | 创建时间 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /infra/apiLog/page | 日志分页查询 |
| GET | /infra/apiLog/get/{logId} | 获取日志详情 |
| DELETE | /infra/apiLog/delete | 删除日志 |
| GET | /infra/apiLog/statistics | 日志统计 |
| GET | /infra/apiLog/topApi | 热门API统计 |

---

### 2.11 错误日志 (`/infra/errorLog`)

**路径**: `/infra/errorLog`
**组件**: `ErrorLogList.vue`
**功能**: 错误日志查看、错误分析

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| logId | String | 日志ID |
| errorType | String | 错误类型 |
| errorMessage | String | 错误信息 |
| stackTrace | Text | 堆栈信息 |
| source | String | 来源 |
| url | String | 请求URL |
| method | String | 请求方法 |
| params | String | 请求参数 |
| userId | String | 用户ID |
| userName | String | 用户名 |
| ip | String | IP地址 |
| userAgent | String | UserAgent |
| createTime | DateTime | 创建时间 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /infra/errorLog/page | 日志分页查询 |
| GET | /infra/errorLog/get/{logId} | 获取日志详情 |
| DELETE | /infra/errorLog/delete | 删除日志 |
| GET | /infra/errorLog/statistics | 错误统计 |
| GET | /infra/errorLog/export | 导出错误日志 |

---

### 2.12 Druid监控 (`/infra/druid`)

**路径**: `/infra/druid`
**组件**: `DruidStatList.vue`
**功能**: 数据库连接池监控、SQL监控

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| name | String | 数据源名称 |
| activeCount | Integer | 活跃连接数 |
| poolingCount | Integer | 池中连接数 |
| poolingPeak | Integer | 池中峰值 |
| waitThreadCount | Integer | 等待线程数 |
| waitThreadPeak | Integer | 等待峰值 |
| executeCount | Integer | 执行次数 |
| executePeak | Integer | 执行峰值 |
| commitCount | Integer | 提交次数 |
| rollbackCount | Integer | 回滚次数 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /infra/druid/datasources | 数据源列表 |
| GET | /infra/druid/datasource/{name} | 数据源详情 |
| GET | /infra/druid/sqlStat | SQL统计 |
| GET | /infra/druid/sqlDetail/{id} | SQL详情 |
| GET | /infra/druid/connectionStat | 连接池统计 |
| POST | /infra/druid/reset | 重置统计 |
| GET | /infra/druid/basicInfo | Druid基本信息 |

---

## 3. 页面清单汇总

| 序号 | 页面名称 | 路径 | 功能说明 |
|-----|---------|------|---------|
| 1 | 文件管理 | /infra/file | 文件上传下载 |
| 2 | 文件配置 | /infra/fileConfig | 存储配置 |
| 3 | 代码生成器 | /infra/codegen | CRUD生成 |
| 4 | 代码模板 | /infra/codegenTemplate | 模板配置 |
| 5 | 数据源配置 | /infra/datasource | 多数据源 |
| 6 | 数据库文档 | /infra/dbdoc | 文档生成 |
| 7 | 定时任务 | /infra/job | 任务调度 |
| 8 | 定时日志 | /infra/jobLog | 任务日志 |
| 9 | Redis管理 | /infra/redis | 缓存管理 |
| 10 | API日志 | /infra/apiLog | API日志 |
| 11 | 错误日志 | /infra/errorLog | 错误日志 |
| 12 | Druid监控 | /infra/druid | 连接监控 |

---

*文档版本: V1.0 | 生成日期: 2026-04-17*
