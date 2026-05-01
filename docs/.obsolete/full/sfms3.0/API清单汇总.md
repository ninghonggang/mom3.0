# SFMS3.0 完整API清单

> 生成时间: 2026-04-16
> 模块: win-module-bpm | win-module-infra | win-module-system | win-module-report

---

## 1. BPM模块 (BPM流程管理)

### 1.1 用户组管理 (BpmUserGroupController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /bpm/user-group/create | 创建用户组 |
| PUT | /bpm/user-group/update | 更新用户组 |
| DELETE | /bpm/user-group/delete | 删除用户组 |
| GET | /bpm/user-group/get | 获得用户组 |
| GET | /bpm/user-group/page | 获得用户组分页 |
| GET | /bpm/user-group/list-all-simple | 获取用户组精简信息列表 |

### 1.2 任务分配规则 (BpmTaskAssignRuleController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /bpm/task-assign-rule/list | 获得任务分配规则列表 |
| POST | /bpm/task-assign-rule/create | 创建任务分配规则 |
| PUT | /bpm/task-assign-rule/update | 更新任务分配规则 |

### 1.3 流程模型 (BpmModelController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /bpm/model/page | 获得模型分页 |
| GET | /bpm/model/get | 获得模型 |
| POST | /bpm/model/create | 新建模型 |
| PUT | /bpm/model/update | 修改模型 |
| POST | /bpm/model/import | 导入模型 |
| POST | /bpm/model/deploy | 部署模型 |
| PUT | /bpm/model/update-state | 修改模型的状态 |
| DELETE | /bpm/model/delete | 删除模型 |

### 1.4 动态表单 (BpmFormController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /bpm/form/create | 创建动态表单 |
| PUT | /bpm/form/update | 更新动态表单 |
| DELETE | /bpm/form/delete | 删除动态表单 |
| GET | /bpm/form/get | 获得动态表单 |
| GET | /bpm/form/list-all-simple | 获得动态表单的精简列表 |
| GET | /bpm/form/page | 获得动态表单分页 |

### 1.5 流程定义 (BpmProcessDefinitionController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /bpm/process-definition/page | 获得流程定义分页 |
| GET | /bpm/process-definition/list | 获得流程定义列表 |
| GET | /bpm/process-definition/get-bpmn-xml | 获得流程定义的BPMN XML |

### 1.6 OA请假申请 (BpmOALeaveController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /bpm/oa/leave/create | 创建请假申请 |
| GET | /bpm/oa/leave/get | 获得请假申请 |
| GET | /bpm/oa/leave/page | 获得请假申请分页 |

### 1.7 流程活动实例 (BpmActivityController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /bpm/activity/list | 生成指定流程实例的高亮流程图 |

### 1.8 流程实例 (BpmProcessInstanceController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /bpm/process-instance/my-page | 获得我的实例分页列表 |
| POST | /bpm/process-instance/create | 新建流程实例 |
| GET | /bpm/process-instance/get | 获得指定流程实例 |
| DELETE | /bpm/process-instance/cancel | 取消流程实例(撤回) |

### 1.9 流程任务实例 (BpmTaskController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /bpm/task/todo-page | 获取Todo待办任务分页 |
| GET | /bpm/task/done-page | 获取Done已办任务分页 |
| GET | /bpm/task/list-by-process-instance-id | 获得指定流程实例的任务列表 |
| PUT | /bpm/task/approve | 通过任务 |
| PUT | /bpm/task/reject | 不通过任务 |
| PUT | /bpm/task/update-assignee | 更新任务的负责人(转派) |

---

## 2. INFRA模块 (基础设施)

### 2.1 代码生成器 (CodegenController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /infra/codegen/db/table/list | 获得数据库自带的表定义列表 |
| GET | /infra/codegen/table/page | 获得表定义分页 |
| GET | /infra/codegen/detail | 获得表和字段的明细 |
| POST | /infra/codegen/create-list | 基于数据库的表结构创建代码生成器的表和字段定义 |
| PUT | /infra/codegen/update | 更新数据库的表和字段定义 |
| PUT | /infra/codegen/sync-from-db | 同步数据库的表和字段定义 |
| DELETE | /infra/codegen/delete | 删除数据库的表和字段定义 |
| GET | /infra/codegen/preview | 预览生成代码 |
| GET | /infra/codegen/download | 下载生成代码 |

### 2.2 参数配置 (ConfigController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /infra/config/create | 创建参数配置 |
| PUT | /infra/config/update | 修改参数配置 |
| DELETE | /infra/config/delete | 删除参数配置 |
| GET | /infra/config/get | 获得参数配置 |
| GET | /infra/config/get-value-by-key | 根据参数键名查询参数值 |
| GET | /infra/config/queryByKey | 根据参数键名查询参数值(详细) |
| GET | /infra/config/page | 获取参数配置分页 |
| GET | /infra/config/export | 导出参数配置 |

### 2.3 数据库文档 (DatabaseDocController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /infra/db-doc/export-html | 导出HTML格式的数据文档 |
| GET | /infra/db-doc/export-word | 导出Word格式的数据文档 |
| GET | /infra/db-doc/export-markdown | 导出Markdown格式的数据文档 |

### 2.4 数据源配置 (DataSourceConfigController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /infra/data-source-config/create | 创建数据源配置 |
| PUT | /infra/data-source-config/update | 更新数据源配置 |
| DELETE | /infra/data-source-config/delete | 删除数据源配置 |
| GET | /infra/data-source-config/get | 获得数据源配置 |
| GET | /infra/data-source-config/list | 获得数据源配置列表 |

### 2.5 文件配置 (FileConfigController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /infra/file-config/create | 创建文件配置 |
| PUT | /infra/file-config/update | 更新文件配置 |
| PUT | /infra/file-config/update-master | 更新文件配置为Master |
| DELETE | /infra/file-config/delete | 删除文件配置 |
| GET | /infra/file-config/get | 获得文件配置 |
| GET | /infra/file-config/page | 获得文件配置分页 |
| GET | /infra/file-config/test | 测试文件配置是否正确 |

### 2.6 文件存储 (FileController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /infra/file/upload | 上传文件 |
| POST | /infra/file/uploadFile | 上传文件到本地磁盘 |
| POST | /infra/file/uploadFileData | 上传文件数据 |
| DELETE | /infra/file/delete | 删除文件 |
| DELETE | /infra/file/deleteByTable | 根据table删除文件 |
| GET | /infra/file/{configId}/get/** | 下载文件 |
| GET | /infra/file/{configId}/show/** | 在线展示文件 |
| GET | /infra/file/page | 获得文件分页 |
| GET | /infra/file/list | 获得文件列表 |
| POST | /infra/file/uploads | 批量上传文件 |

### 2.7 定时任务 (JobController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /infra/job/create | 创建定时任务 |
| PUT | /infra/job/update | 更新定时任务 |
| PUT | /infra/job/update-status | 更新定时任务的状态 |
| DELETE | /infra/job/delete | 删除定时任务 |
| PUT | /infra/job/trigger | 触发定时任务 |
| GET | /infra/job/get | 获得定时任务 |
| GET | /infra/job/list | 获得定时任务列表 |
| GET | /infra/job/page | 获得定时任务分页 |
| GET | /infra/job/export-excel | 导出定时任务Excel |
| GET | /infra/job/get_next_times | 获得定时任务的下n次执行时间 |

### 2.8 定时任务日志 (JobLogController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /infra/job-log/get | 获得定时任务日志 |
| GET | /infra/job-log/list | 获得定时任务日志列表 |
| GET | /infra/job-log/page | 获得定时任务日志分页 |
| GET | /infra/job-log/export-excel | 导出定时任务日志Excel |

### 2.9 API访问日志 (ApiAccessLogController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /infra/api-access-log/page | 获得API访问日志分页 |
| GET | /infra/api-access-log/export-excel | 导出API访问日志Excel |

### 2.10 API错误日志 (ApiErrorLogController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| PUT | /infra/api-error-log/update-status | 更新API错误日志的状态 |
| GET | /infra/api-error-log/page | 获得API错误日志分页 |
| GET | /infra/api-error-log/export-excel | 导出API错误日志Excel |

### 2.11 Redis监控 (RedisController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /infra/redis/get-monitor-info | 获得Redis监控信息 |
| POST | /infra/redis/set | 加入缓存 |
| GET | /infra/redis/get | 获取缓存 |
| DELETE | /infra/redis/delete | 删除缓存 |

### 2.12 备注 (RemarkController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /infra/remark/create | 创建备注 |
| GET | /infra/remark/get | 获得备注 |
| GET | /infra/remark/list | 获得备注列表 |

### 2.13 测试示例 (TestDemoController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /infra/test-demo/create | 创建字典类型 |
| PUT | /infra/test-demo/update | 更新字典类型 |
| DELETE | /infra/test-demo/delete | 删除字典类型 |
| GET | /infra/test-demo/get | 获得字典类型 |
| GET | /infra/test-demo/list | 获得字典类型列表 |
| GET | /infra/test-demo/page | 获得字典类型分页 |
| GET | /infra/test-demo/export-excel | 导出字典类型Excel |

### 2.14 动态记录 (TrendsController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /infra/trends/get | 获得动态记录 |
| GET | /infra/trends/list | 获得动态记录列表 |
| GET | /infra/trends/page | 获得动态记录分页 |
| GET | /infra/trends/export-excel | 导出动态记录Excel |

---

## 3. SYSTEM模块 (系统模块)

### 3.1 认证 (AuthController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/auth/login | 使用账号密码登录 |
| POST | /system/auth/loginNoCode | 使用账号密码登录(无验证码) |
| POST | /system/auth/logout | 登出系统 |
| POST | /system/auth/refresh-token | 刷新令牌 |
| GET | /system/auth/get-permission-info | 获取登录用户的权限信息 |
| GET | /system/auth/public-key | 获取公钥 |

### 3.2 验证码 (CaptchaController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/captcha/check | 校验验证码 |
| GET | /system/captcha/captchaImage | 生成验证码 |

### 3.3 部门 (DeptController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/dept/create | 创建部门 |
| PUT | /system/dept/update | 更新部门 |
| DELETE | /system/dept/delete | 删除部门 |
| GET | /system/dept/list | 获取部门列表 |
| GET | /system/dept/list-all-simple | 获取部门精简信息列表 |
| GET | /system/dept/get | 获得部门信息 |

### 3.4 岗位 (PostController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/post/create | 创建岗位 |
| PUT | /system/post/update | 修改岗位 |
| DELETE | /system/post/delete | 删除岗位 |
| GET | /system/post/get | 获得岗位信息 |
| GET | /system/post/list-all-simple | 获取岗位精简信息列表 |
| GET | /system/post/page | 获得岗位分页列表 |
| GET | /system/post/export | 导出岗位Excel |

### 3.5 字典类型 (DictTypeController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/dict-type/create | 创建字典类型 |
| PUT | /system/dict-type/update | 修改字典类型 |
| DELETE | /system/dict-type/delete | 删除字典类型 |
| GET | /system/dict-type/page | 获得字典类型的分页列表 |
| GET | /system/dict-type/get | 查询字典类型详细 |
| GET | /system/dict-type/list-all-simple | 获得全部字典类型列表 |
| GET | /system/dict-type/export | 导出数据类型 |
| POST | /system/dict-type/getDictByTypes | 根据类型查询展示类型下所有详细列表 |
| GET | /system/dict-type/list-all-data-all | 获得全部字典类型和字典项列表 |

### 3.6 字典数据 (DictDataController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/dict-data/create | 新增字典数据 |
| PUT | /system/dict-data/update | 修改字典数据 |
| DELETE | /system/dict-data/delete | 删除字典数据 |
| GET | /system/dict-data/list-all-simple | 获得全部字典数据列表 |
| GET | /system/dict-data/page | 获得字典类型的分页列表 |
| GET | /system/dict-data/get | 查询字典数据详细 |
| GET | /system/dict-data/queryByDictType | 按字典类型查询数据 |
| GET | /system/dict-data/export | 导出字典数据 |

### 3.7 错误码 (ErrorCodeController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/error-code/create | 创建错误码 |
| PUT | /system/error-code/update | 更新错误码 |
| DELETE | /system/error-code/delete | 删除错误码 |
| GET | /system/error-code/get | 获得错误码 |
| GET | /system/error-code/page | 获得错误码分页 |
| GET | /system/error-code/export-excel | 导出错误码Excel |

### 3.8 地区 (AreaController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /system/area/tree | 获得地区树 |
| GET | /system/area/get-children | 获得地区的下级区域 |
| GET | /system/area/get-by-ids | 通过区域ids获得地区列表 |
| GET | /system/area/get-by-ip | 获得IP对应的地区名 |

### 3.9 登录日志 (LoginLogController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /system/login-log/page | 获得登录日志分页列表 |
| GET | /system/login-log/export | 导出登录日志Excel |

### 3.10 操作日志 (OperateLogController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /system/operate-log/page | 查看操作日志分页列表 |
| GET | /system/operate-log/export | 导出操作日志 |

### 3.11 邮箱账号 (MailAccountController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/mail-account/create | 创建邮箱账号 |
| PUT | /system/mail-account/update | 修改邮箱账号 |
| DELETE | /system/mail-account/delete | 删除邮箱账号 |
| GET | /system/mail-account/get | 获得邮箱账号 |
| GET | /system/mail-account/page | 获得邮箱账号分页 |
| GET | /system/mail-account/list-all-simple | 获得邮箱账号精简列表 |

### 3.12 邮件日志 (MailLogController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /system/mail-log/page | 获得邮箱日志分页 |
| GET | /system/mail-log/get | 获得邮箱日志 |

### 3.13 邮件模板 (MailTemplateController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/mail-template/create | 创建邮件模版 |
| PUT | /system/mail-template/update | 修改邮件模版 |
| DELETE | /system/mail-template/delete | 删除邮件模版 |
| GET | /system/mail-template/get | 获得邮件模版 |
| GET | /system/mail-template/page | 获得邮件模版分页 |
| GET | /system/mail-template/list-all-simple | 获得邮件模版精简列表 |
| POST | /system/mail-template/send-mail | 发送邮件 |
| GET | /system/mail-template/noPage | 获得邮件模版不分页列表 |

### 3.14 消息设置 (MessageSetController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/message-set/create | 创建消息设置 |
| PUT | /system/message-set/update | 更新消息设置 |
| DELETE | /system/message-set/delete | 删除消息设置 |
| GET | /system/message-set/get | 获得消息设置 |
| GET | /system/message-set/page | 获得消息设置分页 |
| POST | /system/message-set/senior | 高级搜索获得消息设置分页 |
| GET | /system/message-set/export-excel | 导出消息设置Excel |
| GET | /system/message-set/get-import-template | 获得导入消息设置模板 |
| POST | /system/message-set/import | 导入消息设置 |
| GET | /system/message-set/noPage | 获得消息设置不分页 |

### 3.15 通知公告 (NoticeController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/notice/create | 创建通知公告 |
| PUT | /system/notice/update | 修改通知公告 |
| DELETE | /system/notice/delete | 删除通知公告 |
| GET | /system/notice/page | 获取通知公告列表 |
| GET | /system/notice/get | 获得通知公告 |

### 3.16 站内信模板 (NotifyTemplateController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/notify-template/create | 创建站内信模版 |
| PUT | /system/notify-template/update | 更新站内信模版 |
| DELETE | /system/notify-template/delete | 删除站内信模版 |
| GET | /system/notify-template/get | 获得站内信模版 |
| GET | /system/notify-template/page | 获得站内信模版分页 |
| POST | /system/notify-template/send-notify | 发送站内信 |
| GET | /system/notify-template/noPage | 获得站内信模版不分页列表 |

### 3.17 站内信 (NotifyMessageController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /system/notify-message/get | 获得站内信 |
| GET | /system/notify-message/page | 获得站内信分页(管理) |
| GET | /system/notify-message/my-page | 获得我的站内信分页 |
| PUT | /system/notify-message/update-read | 标记站内信为已读 |
| PUT | /system/notify-message/update-all-read | 标记所有站内信为已读 |
| GET | /system/notify-message/get-unread-list | 获取当前用户的最新站内信列表 |
| GET | /system/notify-message/get-unread-count | 获得当前用户的未读站内信数量 |

### 3.18 OAuth2.0授权 (OAuth2OpenController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/oauth2/token | 获得访问令牌 |
| DELETE | /system/oauth2/token | 删除访问令牌 |
| POST | /system/oauth2/check-token | 校验访问令牌 |
| GET | /system/oauth2/authorize | 获得授权信息 |
| POST | /system/oauth2/authorize | 申请授权 |

### 3.19 OAuth2.0令牌 (OAuth2TokenController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /system/oauth2-token/page | 获得访问令牌分页 |
| DELETE | /system/oauth2-token/delete | 删除访问令牌 |

### 3.20 OAuth2.0用户 (OAuth2UserController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /system/oauth2/user/get | 获得用户基本信息 |
| PUT | /system/oauth2/user/update | 更新用户基本信息 |

### 3.21 OAuth2客户端 (OAuth2ClientController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/oauth2-client/create | 创建OAuth2客户端 |
| PUT | /system/oauth2-client/update | 更新OAuth2客户端 |
| DELETE | /system/oauth2-client/delete | 删除OAuth2客户端 |
| GET | /system/oauth2-client/get | 获得OAuth2客户端 |
| GET | /system/oauth2-client/page | 获得OAuth2客户端分页 |

### 3.22 密码策略 (PassWordController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /system/password/getConfig | 获取密码策略 |
| POST | /system/password/setConfig | 设置密码策略 |
| GET | /system/password/getRuleList | 获取密码复杂度 |
| GET | /system/password/validateResetTime | 验证密码是否过期 |

### 3.23 菜单 (MenuController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/menu/create | 创建菜单 |
| PUT | /system/menu/update | 修改菜单 |
| DELETE | /system/menu/delete | 删除菜单 |
| GET | /system/menu/list | 获取菜单列表 |
| GET | /system/menu/list-all-simple | 获取菜单精简信息列表 |
| GET | /system/menu/get | 获取菜单信息 |

### 3.24 权限 (PermissionController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /system/permission/list-role-menus | 获得角色拥有的菜单编号 |
| POST | /system/permission/assign-role-menu | 赋予角色菜单 |
| POST | /system/permission/assign-role-data-scope | 赋予角色数据权限 |
| GET | /system/permission/list-user-roles | 获得管理员拥有的角色编号列表 |
| POST | /system/permission/assign-user-role | 赋予用户角色 |

### 3.25 角色 (RoleController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/role/create | 创建角色 |
| PUT | /system/role/update | 修改角色 |
| PUT | /system/role/update-status | 修改角色状态 |
| DELETE | /system/role/delete | 删除角色 |
| GET | /system/role/get | 获得角色信息 |
| GET | /system/role/page | 获得角色分页 |
| GET | /system/role/list-all-simple | 获取角色精简信息列表 |
| GET | /system/role/export | 导出角色Excel |
| GET | /system/role/noPage | 获得角色不分页列表 |

### 3.26 敏感词 (SensitiveWordController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/sensitive-word/create | 创建敏感词 |
| PUT | /system/sensitive-word/update | 更新敏感词 |
| DELETE | /system/sensitive-word/delete | 删除敏感词 |
| GET | /system/sensitive-word/get | 获得敏感词 |
| GET | /system/sensitive-word/page | 获得敏感词分页 |
| GET | /system/sensitive-word/export-excel | 导出敏感词Excel |
| GET | /system/sensitive-word/get-tags | 获取所有敏感词的标签数组 |
| GET | /system/sensitive-word/validate-text | 获得文本所包含的不合法的敏感词数组 |

### 3.27 流水号规则 (SerialNumberController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/serial-number/create | 创建流水号规则 |
| PUT | /system/serial-number/update | 更新流水号规则 |
| DELETE | /system/serial-number/delete | 删除流水号规则 |
| GET | /system/serial-number/get | 获得流水号规则 |
| GET | /system/serial-number/list | 获得流水号规则列表 |
| GET | /system/serial-number/page | 获得流水号规则分页 |
| GET | /system/serial-number/export-excel | 导出流水号规则Excel |

### 3.28 短信回调 (SmsCallbackController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/sms/callback/aliyun | 阿里云短信的回调 |
| POST | /system/sms/callback/tencent | 腾讯云短信的回调 |

### 3.29 短信渠道 (SmsChannelController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/sms-channel/create | 创建短信渠道 |
| PUT | /system/sms-channel/update | 更新短信渠道 |
| DELETE | /system/sms-channel/delete | 删除短信渠道 |
| GET | /system/sms-channel/get | 获得短信渠道 |
| GET | /system/sms-channel/page | 获得短信渠道分页 |
| GET | /system/sms-channel/list-all-simple | 获得短信渠道精简列表 |

### 3.30 短信日志 (SmsLogController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /system/sms-log/page | 获得短信日志分页 |
| GET | /system/sms-log/export-excel | 导出短信日志Excel |

### 3.31 短信模板 (SmsTemplateController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/sms-template/create | 创建短信模板 |
| PUT | /system/sms-template/update | 更新短信模板 |
| DELETE | /system/sms-template/delete | 删除短信模板 |
| GET | /system/sms-template/get | 获得短信模板 |
| GET | /system/sms-template/page | 获得短信模板分页 |
| GET | /system/sms-template/export-excel | 导出短信模板Excel |
| POST | /system/sms-template/send-sms | 发送短信 |
| GET | /system/sms-template/noPage | 获得短信模板不分页列表 |

### 3.32 表名动作关系 (TableActionRelController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/table-action-rel/create | 创建表名动作关系 |
| PUT | /system/table-action-rel/update | 更新表名动作关系 |
| DELETE | /system/table-action-rel/delete | 删除表名动作关系 |
| GET | /system/table-action-rel/get | 获得表名动作关系 |
| GET | /system/table-action-rel/page | 获得表名动作关系分页 |
| POST | /system/table-action-rel/senior | 高级搜索表名动作关系 |
| GET | /system/table-action-rel/export-excel | 导出表名动作关系Excel |
| GET | /system/table-action-rel/get-import-template | 获得导入模板 |
| POST | /system/table-action-rel/import | 导入表名动作关系 |
| GET | /system/table-action-rel/noPage | 获得表名动作关系不分页 |

### 3.33 模板关系 (TemplateRelController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/template-rel/create | 创建模板关系 |
| PUT | /system/template-rel/update | 更新模板关系 |
| DELETE | /system/template-rel/delete | 删除模板关系 |
| GET | /system/template-rel/get | 获得模板关系 |
| GET | /system/template-rel/page | 获得模板关系分页 |
| POST | /system/template-rel/senior | 高级搜索模板关系 |
| GET | /system/template-rel/export-excel | 导出模板关系Excel |
| GET | /system/template-rel/get-import-template | 获得导入模板 |
| POST | /system/template-rel/import | 导入模板关系 |

### 3.34 租户套餐 (TenantPackageController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/tenant-package/create | 创建租户套餐 |
| PUT | /system/tenant-package/update | 更新租户套餐 |
| DELETE | /system/tenant-package/delete | 删除租户套餐 |
| GET | /system/tenant-package/get | 获得租户套餐 |
| GET | /system/tenant-package/page | 获得租户套餐分页 |
| GET | /system/tenant-package/get-simple-list | 获取租户套餐精简信息列表 |

### 3.35 租户 (TenantController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /system/tenant/get-id-by-name | 使用租户名获得租户编号 |
| POST | /system/tenant/senior | 数据筛选 |
| POST | /system/tenant/create | 创建租户 |
| PUT | /system/tenant/update | 更新租户 |
| DELETE | /system/tenant/delete | 删除租户 |
| GET | /system/tenant/get | 获得租户 |
| GET | /system/tenant/page | 获得租户分页 |
| GET | /system/tenant/export-excel | 导出租户Excel |

### 3.36 用户 (UserController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /system/user/create | 新增用户 |
| PUT | /system/user/update | 修改用户 |
| DELETE | /system/user/delete | 删除用户 |
| PUT | /system/user/update-password | 重置用户密码 |
| PUT | /system/user/update-password-by-mailKey | 通过邮件链接重置密码 |
| PUT | /system/user/update-status | 修改用户状态 |
| GET | /system/user/page | 获得用户分页列表 |
| POST | /system/user/senior | 高级搜索用户信息分页 |
| GET | /system/user/list-all-simple | 获取用户精简信息列表 |
| GET | /system/user/get | 获得用户详情 |
| GET | /system/user/export | 导出用户 |
| GET | /system/user/get-import-template | 获得导入用户模板 |
| POST | /system/user/import | 导入用户 |
| PUT | /system/user/forgetPassword | 忘记密码 |
| PUT | /system/user/updatePassword | 修改密码 |
| POST | /system/user/getUserListByDeptIds | 获得用户列表(按部门) |
| GET | /system/user/listAll | 查询所有用户信息列表 |
| GET | /system/user/unLockUser | 解除冻结 |

### 3.37 用户个人中心 (UserProfileController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| GET | /system/user/profile/get | 获得登录用户信息 |
| PUT | /system/user/profile/update | 修改用户个人信息 |
| PUT | /system/user/profile/update-password | 修改用户个人密码 |
| POST | /system/user/profile/update-avatar | 上传用户个人头像 |

---

## 4. REPORT模块 (报表模块)

### 4.1 GoView数据 (GoViewDataController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| REQUEST | /report/go-view/data/get-by-sql | 使用SQL查询数据 |
| REQUEST | /report/go-view/data/get-by-http | 使用HTTP查询数据 |

### 4.2 GoView项目 (GoViewProjectController)

| HTTP方法 | URL路径 | 接口说明 |
|---------|--------|---------|
| POST | /report/go-view/project/create | 创建项目 |
| PUT | /report/go-view/project/update | 更新项目 |
| DELETE | /report/go-view/project/delete | 删除GoView项目 |
| GET | /report/go-view/project/get | 获得项目 |
| GET | /report/go-view/project/my-page | 获得我的项目分页 |

---

## API统计汇总

| 模块 | Controller数量 | API数量 |
|------|---------------|---------|
| BPM | 9 | 32 |
| INFRA | 14 | 53 |
| SYSTEM | 37 | 192 |
| REPORT | 2 | 7 |
| **总计** | **62** | **284** |

---

*本文档由自动化工具生成，如需更新请重新运行脚本*
