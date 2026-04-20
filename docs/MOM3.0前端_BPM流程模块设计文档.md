# MOM3.0前端_BPM流程模块设计文档

> **版本**: V1.1
> **日期**: 2026-04-17
> **项目**: 闻荫科技MOM3.0智能制造执行系统
> **对比版本**: SFMS3.0 BPM模块 (27页缺失)

---

## 1. 模块概述

### 1.1 功能定位

BPM（Business Process Management）流程管理模块基于Flowable实现流程定义、部署、执行、审批等核心功能。

### 1.2 当前状态 vs SFMS3.0

| 类别 | MOM3.0现状 | SFMS3.0完整功能 |
|------|-----------|----------------|
| 页面数量 | 3页 | 27页完整实现 |
| 流程模型设计 | 缺失 | 完整BPMN设计器 |
| 表单设计 | 缺失 | 可视化表单设计器 |
| 用户组管理 | 缺失 | 用户组CRUD |
| 任务分配规则 | 缺失 | 分配规则配置 |
| 任务转派 | 缺失 | 转派申请、转派审批 |
| 任务退回 | 缺失 | 退回申请、退回审批 |
| 审批管理 | 缺失 | 审批意见、审批历史 |
| 驳回管理 | 缺失 | 驳回理由、驳回记录 |
| 委托管理 | 缺失 | 委托设置、委托任务 |
| 流程监控 | 缺失 | 流程状态、流程跟踪 |
| 流程日志 | 缺失 | 操作日志、执行记录 |
| 流程版本 | 缺失 | 版本管理、版本对比 |
| 流程分类 | 缺失 | 分类配置、分类管理 |
| 节点配置 | 缺失 | 节点属性、事件配置 |
| 表达式配置 | 缺失 | 条件表达式、脚本配置 |
| 消息配置 | 缺失 | 消息事件、通知配置 |
| 定时配置 | 缺失 | 定时触发、调度配置 |
| 异常处理 | 缺失 | 异常捕获、处理规则 |
| 流程仿真 | 缺失 | 仿真测试、流程验证 |
| 流程导出 | 缺失 | 导出管理 |
| 流程导入 | 缺失 | 导入管理 |
| OA请假示例 | 缺失 | 请假流程 |
| 采购审批示例 | 缺失 | 采购审批 |
| 报销审批示例 | 缺失 | 报销审批 |
| 驳回管理 | 缺失 | 驳回管理 |
| 委托管理 | 缺失 | 委托管理 |
| 流程版本 | 缺失 | 版本管理 |

---

## 2. 缺失页面详细设计

### 2.1 流程模型设计器 (`/bpm/modeler`)

**路径**: `/bpm/modeler`
**组件**: `Modeler.vue`
**功能**: BPMN可视化设计

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| modelId | number | 模型ID |
| name | string | 模型名称 |
| key | string | 模型Key |
| category | string | 分类 |
| version | number | 版本号 |
| bpmnXml | string | BPMN XML内容 |
| svgXml | string | SVG图形内容 |
| status | number | 状态 0-草稿 1-已部署 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/models | 获取模型列表 |
| GET | /bpm/models/{id} | 获取模型详情 |
| POST | /bpm/models | 创建模型 |
| PUT | /bpm/models/{id} | 更新模型 |
| DELETE | /bpm/models/{id} | 删除模型 |
| POST | /bpm/models/{id}/publish | 部署模型 |
| GET | /bpm/models/{id}/bpmn | 获取BPMN XML |
| PUT | /bpm/models/{id}/bpmn | 保存BPMN XML |
| GET | /bpm/models/{id}/svg | 获取SVG图形 |

---

### 2.2 表单设计器 (`/bpm/formDesigner`)

**路径**: `/bpm/formDesigner`
**组件**: `FormDesigner.vue`
**功能**: 表单设计、表单组件

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| formId | number | 表单ID |
| formName | string | 表单名称 |
| formKey | string | 表单标识 |
| category | string | 分类 |
| fields | array | 字段列表 |
| layout | object | 布局配置 |
| validation | object | 校验规则 |
| status | number | 状态 0-草稿 1-已发布 |

**支持的组件类型**:
| 组件 | 说明 |
|------|------|
| input | 单行文本 |
| textarea | 多行文本 |
| number | 数字输入 |
| select | 下拉选择 |
| radio | 单选按钮 |
| checkbox | 多选框 |
| date | 日期选择 |
| datetime | 日期时间选择 |
| time | 时间选择 |
| file | 文件上传 |
| image | 图片上传 |
| editor | 富文本编辑器 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/form/list | 获取表单列表 |
| GET | /bpm/form/{id} | 获取表单详情 |
| POST | /bpm/form | 创建表单 |
| PUT | /bpm/form/{id} | 更新表单 |
| DELETE | /bpm/form/{id} | 删除表单 |
| POST | /bpm/form/{id}/publish | 发布表单 |
| GET | /bpm/form/{id}/render | 渲染表单 |

---

### 2.3 用户组管理 (`/bpm/userGroup`)

**路径**: `/bpm/userGroup`
**组件**: `UserGroupList.vue`
**功能**: 用户组CRUD、组成员

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| groupId | number | 组ID |
| groupName | string | 组名称 |
| groupCode | string | 组编码 |
| groupType | string | 组类型(角色/部门/自定义) |
| status | number | 状态 0-禁用 1-启用 |
| memberIds | array | 成员ID列表 |
| memberNames | string | 成员名称列表 |
| remark | string | 备注 |
| createTime | datetime | 创建时间 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/user-group/list | 获取用户组列表 |
| GET | /bpm/user-group/{id} | 获取用户组详情 |
| POST | /bpm/user-group | 创建用户组 |
| PUT | /bpm/user-group/{id} | 更新用户组 |
| DELETE | /bpm/user-group/{id} | 删除用户组 |
| GET | /bpm/user-group/{id}/members | 获取组成员 |
| PUT | /bpm/user-group/{id}/members | 更新组成员 |

---

### 2.4 任务分配规则 (`/bpm/assignRule`)

**路径**: `/bpm/assignRule`
**组件**: `AssignRuleList.vue`
**功能**: 分配规则、分配条件

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| ruleId | number | 规则ID |
| ruleName | string | 规则名称 |
| modelId | number | 流程模型ID |
| taskDefineKey | string | 任务定义Key |
| ruleType | number | 规则类型 1-指定用户 2-指定角色 3-用户组 4-脚本计算 |
| assigneeType | string | 处理人类型 |
| assigneeValue | string | 处理人值 |
| priority | number | 优先级 |
| conditionExpr | string | 条件表达式 |
| status | number | 状态 0-禁用 1-启用 |

**规则类型说明**:
| 类型值 | 说明 | 配置 |
|--------|------|------|
| 1 | 指定用户 | assigneeValue为用户ID |
| 2 | 指定角色 | assigneeValue为角色编码 |
| 3 | 用户组 | assigneeValue为用户组ID |
| 4 | 脚本计算 | assigneeValue为脚本内容 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/assign-rule/list | 获取规则列表 |
| GET | /bpm/assign-rule/{id} | 获取规则详情 |
| POST | /bpm/assign-rule | 创建规则 |
| PUT | /bpm/assign-rule/{id} | 更新规则 |
| DELETE | /bpm/assign-rule/{id} | 删除规则 |
| GET | /bpm/assign-rule/model/{modelId} | 按模型获取规则 |

---

### 2.5 任务转派 (`/bpm/taskTransfer`)

**路径**: `/bpm/taskTransfer`
**组件**: `TaskTransferList.vue`
**功能**: 转派申请、转派审批

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| transferId | number | 转派ID |
| taskId | string | 任务ID |
| taskName | string | 任务名称 |
| processInstanceId | string | 流程实例ID |
| originalAssignee | string | 原执行人ID |
| originalAssigneeName | string | 原执行人名称 |
| newAssignee | string | 新执行人ID |
| newAssigneeName | string | 新执行人名称 |
| transferReason | string | 转派原因 |
| transferType | number | 转派类型 1-正常转派 2-委托转派 |
| status | number | 状态 0-待审批 1-已批准 2-已拒绝 |
| applicant | string | 申请人 |
| applyTime | datetime | 申请时间 |
| approveTime | datetime | 审批时间 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/task-transfer/list | 获取转派列表 |
| GET | /bpm/task-transfer/{id} | 获取转派详情 |
| POST | /bpm/task-transfer | 创建转派申请 |
| PUT | /bpm/task-transfer/{id}/approve | 批准转派 |
| PUT | /bpm/task-transfer/{id}/reject | 拒绝转派 |
| GET | /bpm/task-transfer/todo | 获取待审批转派 |
| GET | /bpm/task-transfer/done | 获取已办转派 |

---

### 2.6 任务退回 (`/bpm/taskReturn`)

**路径**: `/bpm/taskReturn`
**组件**: `TaskReturnList.vue`
**功能**: 退回申请、退回审批

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| returnId | number | 退回ID |
| taskId | string | 任务ID |
| taskName | string | 任务名称 |
| processInstanceId | string | 流程实例ID |
| currentNode | string | 当前节点 |
| returnNode | string | 退回节点 |
| returnReason | string | 退回原因 |
| status | number | 状态 0-待审批 1-已批准 2-已拒绝 |
| applicant | string | 申请人 |
| applyTime | datetime | 申请时间 |
| approveTime | datetime | 审批时间 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/task-return/list | 获取退回列表 |
| GET | /bpm/task-return/{id} | 获取退回详情 |
| POST | /bpm/task-return | 创建退回申请 |
| POST | /bpm/task-return/{id}/approve | 批准退回 |
| POST | /bpm/task-return/{id}/reject | 拒绝退回 |
| GET | /bpm/task-return/nodes/{taskId} | 获取可退回节点 |
| GET | /bpm/task-return/todo | 获取待审批退回 |
| GET | /bpm/task-return/done | 获取已办退回 |

---

### 2.7 审批管理 (`/bpm/approve`)

**路径**: `/bpm/approve`
**组件**: `ApproveList.vue`
**功能**: 审批意见、审批历史

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| approvalId | number | 审批ID |
| taskId | string | 任务ID |
| taskName | string | 任务名称 |
| processInstanceId | string | 流程实例ID |
| processDefinitionName | string | 流程名称 |
| nodeName | string | 节点名称 |
| approver | string | 审批人ID |
| approverName | string | 审批人名称 |
| approvalOpinion | string | 审批意见 |
| approvalResult | number | 审批结果 1-通过 2-驳回 3-退回 |
| approvalTime | datetime | 审批时间 |
| duration | number | 处理时长(分钟) |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/approval/list | 获取审批列表 |
| GET | /bpm/approval/{id} | 获取审批详情 |
| GET | /bpm/approval/history/{processInstanceId} | 获取审批历史 |
| GET | /bpm/approval/todo | 获取待审批列表 |
| GET | /bpm/approval/done | 获取已审批列表 |
| POST | /bpm/approval/{taskId}/approve | 审批通过 |
| POST | /bpm/approval/{taskId}/reject | 审批驳回 |

---

### 2.8 驳回管理 (`/bpm/reject`)

**路径**: `/bpm/reject`
**组件**: `RejectList.vue`
**功能**: 驳回理由、驳回记录

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| rejectId | number | 驳回ID |
| taskId | string | 任务ID |
| taskName | string | 任务名称 |
| processInstanceId | string | 流程实例ID |
| rejectReason | string | 驳回理由 |
| rejectType | number | 驳回类型 1-审批驳回 2-退回驳回 |
| rejecter | string | 驳回人ID |
| rejecterName | string | 驳回人名称 |
| rejectTime | datetime | 驳回时间 |
| status | number | 状态 0-待处理 1-已处理 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/reject/list | 获取驳回列表 |
| GET | /bpm/reject/{id} | 获取驳回详情 |
| POST | /bpm/reject/{id}/handle | 处理驳回 |
| GET | /bpm/reject/todo | 获取待处理驳回 |
| GET | /bpm/reject/history | 获取驳回历史 |

---

### 2.9 委托管理 (`/bpm/delegate`)

**路径**: `/bpm/delegate`
**组件**: `DelegateList.vue`
**功能**: 委托设置、委托任务

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| delegateId | number | 委托ID |
| delegator | string | 委托人ID |
| delegatorName | string | 委托人名称 |
| delegatee | string | 被委托人ID |
| delegateeName | string | 被委托人名称 |
| processDefinitionKey | string | 流程定义Key(空表示全部) |
| startTime | datetime | 委托开始时间 |
| endTime | datetime | 委托结束时间 |
| reason | string | 委托原因 |
| status | number | 状态 0-禁用 1-启用 |
| createTime | datetime | 创建时间 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/delegate/list | 获取委托列表 |
| GET | /bpm/delegate/{id} | 获取委托详情 |
| POST | /bpm/delegate | 创建委托 |
| PUT | /bpm/delegate/{id} | 更新委托 |
| DELETE | /bpm/delegate/{id} | 删除委托 |
| PUT | /bpm/delegate/{id}/enable | 启用委托 |
| PUT | /bpm/delegate/{id}/disable | 禁用委托 |
| GET | /bpm/delegate/my | 获取我的委托 |

---

### 2.10 流程监控 (`/bpm/monitor`)

**路径**: `/bpm/monitor`
**组件**: `MonitorList.vue`
**功能**: 流程状态、流程跟踪

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| instanceId | string | 实例ID |
| processDefinitionId | string | 流程定义ID |
| processDefinitionName | string | 流程名称 |
| businessKey | string | 业务Key |
| startTime | datetime | 开始时间 |
| endTime | datetime | 结束时间 |
| status | number | 状态 1-审批中 2-已完成 3-已取消 |
| result | number | 结果 1-通过 2-不通过 3-取消 |
| currentNodes | array | 当前节点列表 |
| startUserName | string | 发起人名称 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/monitor/instance/list | 获取监控实例列表 |
| GET | /bpm/monitor/instance/{id} | 获取实例详情 |
| GET | /bpm/monitor/instance/{id}/bpmn | 获取实例BPMN |
| GET | /bpm/monitor/instance/{id}/tasks | 获取实例任务列表 |
| GET | /bpm/monitor/instance/{id}/history | 获取实例历史 |
| POST | /bpm/monitor/instance/{id}/terminate | 终止实例 |
| GET | /bpm/monitor/statistics | 获取监控统计 |

---

### 2.11 流程日志 (`/bpm/log`)

**路径**: `/bpm/log`
**组件**: `LogList.vue`
**功能**: 操作日志、执行记录

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| logId | number | 日志ID |
| processInstanceId | string | 流程实例ID |
| taskId | string | 任务ID |
| taskName | string | 任务名称 |
| operator | string | 操作人ID |
| operatorName | string | 操作人名称 |
| operationType | string | 操作类型 |
| operationDetail | string | 操作详情 |
| operationTime | datetime | 操作时间 |
| ipAddress | string | IP地址 |

**操作类型枚举**:
| 类型值 | 说明 |
|--------|------|
| START | 启动流程 |
| SUBMIT | 提交任务 |
| APPROVE | 审批 |
| REJECT | 驳回 |
| RETURN | 退回 |
| TRANSFER | 转派 |
| DELEGATE | 委托 |
| TERMINATE | 终止 |
| COMPLETE | 完成 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/log/list | 获取日志列表 |
| GET | /bpm/log/{id} | 获取日志详情 |
| GET | /bpm/log/instance/{instanceId} | 获取实例日志 |
| GET | /bpm/log/task/{taskId} | 获取任务日志 |
| GET | /bpm/log/export | 导出日志 |

---

### 2.12 流程版本 (`/bpm/version`)

**路径**: `/bpm/version`
**组件**: `VersionList.vue`
**功能**: 版本管理、版本对比

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| versionId | number | 版本ID |
| processDefinitionId | string | 流程定义ID |
| version | number | 版本号 |
| deploymentTime | datetime | 部署时间 |
| deployer | string | 部署人 |
| deploymentId | string | 部署ID |
| status | number | 状态 0-禁用 1-启用 |
| changeLog | string | 变更日志 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/version/list | 获取版本列表 |
| GET | /bpm/version/{id} | 获取版本详情 |
| GET | /bpm/version/compare | 对比两个版本 |
| POST | /bpm/version/{id}/enable | 启用版本 |
| POST | /bpm/version/{id}/disable | 禁用版本 |
| GET | /bpm/version/{id}/bpmn | 获取版本BPMN |

---

### 2.13 流程分类 (`/bpm/category`)

**路径**: `/bpm/category`
**组件**: `CategoryList.vue`
**功能**: 分类配置、分类管理

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| categoryId | number | 分类ID |
| categoryName | string | 分类名称 |
| categoryCode | string | 分类编码 |
| parentId | number | 父分类ID |
| sortOrder | number | 排序 |
| status | number | 状态 0-禁用 1-启用 |
| remark | string | 备注 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/category/list | 获取分类列表 |
| GET | /bpm/category/tree | 获取分类树 |
| GET | /bpm/category/{id} | 获取分类详情 |
| POST | /bpm/category | 创建分类 |
| PUT | /bpm/category/{id} | 更新分类 |
| DELETE | /bpm/category/{id} | 删除分类 |

---

### 2.14 节点配置 (`/bpm/nodeConfig`)

**路径**: `/bpm/nodeConfig`
**组件**: `NodeConfigList.vue`
**功能**: 节点属性、事件配置

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| configId | number | 配置ID |
| modelId | number | 模型ID |
| nodeId | string | 节点ID |
| nodeName | string | 节点名称 |
| nodeType | string | 节点类型(start/end/task/gateway) |
| properties | JSON | 属性配置 |
| listeners | JSON | 事件监听器 |
| forms | JSON | 表单配置 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/node-config/list | 获取节点配置列表 |
| GET | /bpm/node-config/{id} | 获取节点配置详情 |
| PUT | /bpm/node-config/{id} | 更新节点配置 |
| GET | /bpm/node-config/model/{modelId} | 获取模型所有节点配置 |

---

### 2.15 表达式配置 (`/bpm/expressionConfig`)

**路径**: `/bpm/expressionConfig`
**组件**: `ExpressionConfigList.vue`
**功能**: 条件表达式、脚本配置

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| expressionId | number | 表达式ID |
| expressionName | string | 表达式名称 |
| expressionType | string | 表达式类型(条件/脚本) |
| expressionContent | string | 表达式内容 |
| expressionLang | string | 脚本语言(groovy/javascript) |
| description | string | 描述 |
| status | number | 状态 0-禁用 1-启用 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/expression/list | 获取表达式列表 |
| GET | /bpm/expression/{id} | 获取表达式详情 |
| POST | /bpm/expression | 创建表达式 |
| PUT | /bpm/expression/{id} | 更新表达式 |
| DELETE | /bpm/expression/{id} | 删除表达式 |
| POST | /bpm/expression/validate | 验证表达式 |

---

### 2.16 消息配置 (`/bpm/messageConfig`)

**路径**: `/bpm/messageConfig`
**组件**: `MessageConfigList.vue`
**功能**: 消息事件、通知配置

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| messageId | number | 消息ID |
| messageName | string | 消息名称 |
| messageType | string | 消息类型(信号/消息/通知) |
| messageCode | string | 消息编码 |
| content | string | 消息内容 |
| receivers | array | 接收者配置 |
| channel | string | 通知渠道(站内信/邮件/短信) |
| triggerEvent | string | 触发事件 |
| status | number | 状态 0-禁用 1-启用 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/message-config/list | 获取消息配置列表 |
| GET | /bpm/message-config/{id} | 获取消息配置详情 |
| POST | /bpm/message-config | 创建消息配置 |
| PUT | /bpm/message-config/{id} | 更新消息配置 |
| DELETE | /bpm/message-config/{id} | 删除消息配置 |
| PUT | /bpm/message-config/{id}/enable | 启用消息配置 |
| PUT | /bpm/message-config/{id}/disable | 禁用消息配置 |

---

### 2.17 定时配置 (`/bpm/timerConfig`)

**路径**: `/bpm/timerConfig`
**组件**: `TimerConfigList.vue`
**功能**: 定时触发、调度配置

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| timerId | number | 定时ID |
| timerName | string | 定时名称 |
| cronExpression | string | Cron表达式 |
| processDefinitionKey | string | 流程定义Key |
| startTime | datetime | 开始时间 |
| endTime | datetime | 结束时间 |
| triggerType | string | 触发类型(一次/循环/周期) |
| description | string | 描述 |
| status | number | 状态 0-禁用 1-启用 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/timer-config/list | 获取定时配置列表 |
| GET | /bpm/timer-config/{id} | 获取定时配置详情 |
| POST | /bpm/timer-config | 创建定时配置 |
| PUT | /bpm/timer-config/{id} | 更新定时配置 |
| DELETE | /bpm/timer-config/{id} | 删除定时配置 |
| PUT | /bpm/timer-config/{id}/enable | 启用定时配置 |
| PUT | /bpm/timer-config/{id}/disable | 禁用定时配置 |
| GET | /bpm/timer-config/cron/validate | 验证Cron表达式 |

---

### 2.18 异常处理 (`/bpm/exceptionHandle`)

**路径**: `/bpm/exceptionHandle`
**组件**: `ExceptionHandleList.vue`
**功能**: 异常捕获、处理规则

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| exceptionId | number | 异常ID |
| exceptionName | string | 异常名称 |
| exceptionType | string | 异常类型 |
| processDefinitionKey | string | 流程定义Key |
| nodeId | string | 节点ID |
| handleType | string | 处理方式(重试/跳过/终止/人工处理) |
| handleRule | JSON | 处理规则配置 |
| maxRetries | number | 最大重试次数 |
| description | string | 描述 |
| status | number | 状态 0-禁用 1-启用 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/exception-handle/list | 获取异常处理列表 |
| GET | /bpm/exception-handle/{id} | 获取异常处理详情 |
| POST | /bpm/exception-handle | 创建异常处理 |
| PUT | /bpm/exception-handle/{id} | 更新异常处理 |
| DELETE | /bpm/exception-handle/{id} | 删除异常处理 |
| GET | /bpm/exception-handle/history | 获取异常历史 |

---

### 2.19 流程仿真 (`/bpm/simulation`)

**路径**: `/bpm/simulation`
**组件**: `SimulationList.vue`
**功能**: 仿真测试、流程验证

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| simulationId | number | 仿真ID |
| modelId | number | 模型ID |
| simulationName | string | 仿真名称 |
| testData | JSON | 测试数据 |
| testScenario | string | 测试场景 |
| startTime | datetime | 开始时间 |
| endTime | datetime | 结束时间 |
| result | string | 仿真结果 |
| reportUrl | string | 报告URL |
| status | number | 状态 0-进行中 1-完成 2-失败 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/simulation/list | 获取仿真列表 |
| GET | /bpm/simulation/{id} | 获取仿真详情 |
| POST | /bpm/simulation | 创建仿真 |
| POST | /bpm/simulation/{id}/start | 启动仿真 |
| POST | /bpm/simulation/{id}/stop | 停止仿真 |
| GET | /bpm/simulation/{id}/report | 获取仿真报告 |

---

### 2.20 流程导出 (`/bpm/export`)

**路径**: `/bpm/export`
**组件**: `ExportList.vue`
**功能**: 导出管理

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| exportId | number | 导出ID |
| exportName | string | 导出名称 |
| modelId | number | 模型ID |
| exportFormat | string | 导出格式(bpmn/xml/json/zip) |
| exportPath | string | 导出路径 |
| exportTime | datetime | 导出时间 |
| exporter | string | 导出人 |
| status | number | 状态 0-处理中 1-完成 2-失败 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/export/list | 获取导出列表 |
| GET | /bpm/export/{id} | 获取导出详情 |
| POST | /bpm/export | 创建导出任务 |
| GET | /bpm/export/{id}/download | 下载导出文件 |
| DELETE | /bpm/export/{id} | 删除导出记录 |

---

### 2.21 流程导入 (`/bpm/import`)

**路径**: `/bpm/import`
**组件**: `ImportList.vue`
**功能**: 导入管理

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| importId | number | 导入ID |
| importName | string | 导入名称 |
| importFile | string | 导入文件 |
| importFormat | string | 导入格式(bpmn/xml/json/zip) |
| importMode | string | 导入模式(新建/覆盖/合并) |
| importTime | datetime | 导入时间 |
| importer | string | 导入人 |
| result | string | 导入结果 |
| status | number | 状态 0-处理中 1-成功 2-失败 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/import/list | 获取导入列表 |
| GET | /bpm/import/{id} | 获取导入详情 |
| POST | /bpm/import | 创建导入任务 |
| GET | /bpm/import/template | 下载导入模板 |
| GET | /bpm/import/{id}/preview | 预览导入内容 |

---

### 2.22 OA请假示例 (`/bpm/leaveExample`)

**路径**: `/bpm/leaveExample`
**组件**: `LeaveExampleList.vue`
**功能**: 请假流程

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| leaveId | number | 请假ID |
| leaveNo | string | 请假单号 |
| leaveType | string | 请假类型(年假/病假/事假/婚假/产假/丧假) |
| startTime | datetime | 开始时间 |
| endTime | datetime | 结束时间 |
| duration | decimal | 时长(小时) |
| reason | string | 请假原因 |
| instanceId | string | 流程实例ID |
| status | number | 状态 0-草稿 1-审批中 2-已通过 3-已驳回 |
| applicant | string | 申请人 |
| applyTime | datetime | 申请时间 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/leave/list | 获取请假列表 |
| GET | /bpm/leave/{id} | 获取请假详情 |
| POST | /bpm/leave | 创建请假 |
| PUT | /bpm/leave/{id} | 更新请假 |
| DELETE | /bpm/leave/{id} | 删除请假 |
| POST | /bpm/leave/{id}/submit | 提交请假 |
| POST | /bpm/leave/{id}/cancel | 取消请假 |

---

### 2.23 采购审批示例 (`/bpm/purchaseExample`)

**路径**: `/bpm/purchaseExample`
**组件**: `PurchaseExampleList.vue`
**功能**: 采购审批

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| purchaseId | number | 采购ID |
| purchaseNo | string | 采购单号 |
| purchaseType | string | 采购类型(物资/设备/服务) |
| supplier | string | 供应商 |
| amount | decimal | 采购金额 |
| currency | string | 币种 |
| reason | string | 采购原因 |
| instanceId | string | 流程实例ID |
| status | number | 状态 0-草稿 1-审批中 2-已通过 3-已驳回 |
| applicant | string | 申请人 |
| applyTime | datetime | 申请时间 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/purchase/list | 获取采购列表 |
| GET | /bpm/purchase/{id} | 获取采购详情 |
| POST | /bpm/purchase | 创建采购 |
| PUT | /bpm/purchase/{id} | 更新采购 |
| DELETE | /bpm/purchase/{id} | 删除采购 |
| POST | /bpm/purchase/{id}/submit | 提交采购 |
| POST | /bpm/purchase/{id}/cancel | 取消采购 |

---

### 2.24 报销审批示例 (`/bpm/expenseExample`)

**路径**: `/bpm/expenseExample`
**组件**: `ExpenseExampleList.vue`
**功能**: 报销审批

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| expenseId | number | 报销ID |
| expenseNo | string | 报销单号 |
| expenseType | string | 报销类型(差旅/交通/餐饮/办公/其他) |
| amount | decimal | 报销金额 |
| currency | string | 币种 |
| projectCode | string | 项目编码 |
| reason | string | 报销原因 |
| receipts | array | 票据列表 |
| instanceId | string | 流程实例ID |
| status | number | 状态 0-草稿 1-审批中 2-已通过 3-已驳回 |
| applicant | string | 申请人 |
| applyTime | datetime | 申请时间 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/expense/list | 获取报销列表 |
| GET | /bpm/expense/{id} | 获取报销详情 |
| POST | /bpm/expense | 创建报销 |
| PUT | /bpm/expense/{id} | 更新报销 |
| DELETE | /bpm/expense/{id} | 删除报销 |
| POST | /bpm/expense/{id}/submit | 提交报销 |
| POST | /bpm/expense/{id}/cancel | 取消报销 |

---

### 2.25 驳回管理 (`/bpm/rejectMgmt`)

**路径**: `/bpm/rejectMgmt`
**组件**: `RejectMgmtList.vue`
**功能**: 驳回管理

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| rejectId | number | 驳回ID |
| rejectNo | string | 驳回单号 |
| taskId | string | 任务ID |
| taskName | string | 任务名称 |
| processInstanceId | string | 流程实例ID |
| rejectReason | string | 驳回原因 |
| rejectType | string | 驳回类型 |
| rejecter | string | 驳回人 |
| rejectTime | datetime | 驳回时间 |
| status | number | 状态 0-待处理 1-已处理 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/reject-mgmt/list | 获取驳回列表 |
| GET | /bpm/reject-mgmt/{id} | 获取驳回详情 |
| POST | /bpm/reject-mgmt/{id}/handle | 处理驳回 |
| GET | /bpm/reject-mgmt/todo | 获取待处理驳回 |
| GET | /bpm/reject-mgmt/history | 获取驳回历史 |

---

### 2.26 委托管理 (`/bpm/delegateMgmt`)

**路径**: `/bpm/delegateMgmt`
**组件**: `DelegateMgmtList.vue`
**功能**: 委托管理

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| delegateId | number | 委托ID |
| delegateNo | string | 委托单号 |
| delegator | string | 委托人 |
| delegatorName | string | 委托人名称 |
| delegatee | string | 被委托人 |
| delegateeName | string | 被委托人名称 |
| processDefinitionKey | string | 流程定义Key |
| delegateType | string | 委托类型 |
| startTime | datetime | 开始时间 |
| endTime | datetime | 结束时间 |
| status | number | 状态 0-禁用 1-启用 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/delegate-mgmt/list | 获取委托列表 |
| GET | /bpm/delegate-mgmt/{id} | 获取委托详情 |
| POST | /bpm/delegate-mgmt | 创建委托 |
| PUT | /bpm/delegate-mgmt/{id} | 更新委托 |
| DELETE | /bpm/delegate-mgmt/{id} | 删除委托 |
| PUT | /bpm/delegate-mgmt/{id}/enable | 启用委托 |
| PUT | /bpm/delegate-mgmt/{id}/disable | 禁用委托 |

---

### 2.27 流程版本 (`/bpm/versionMgmt`)

**路径**: `/bpm/versionMgmt`
**组件**: `VersionMgmtList.vue`
**功能**: 版本管理

**核心字段**:
| 字段 | 类型 | 说明 |
|------|------|------|
| versionId | number | 版本ID |
| versionNo | string | 版本号 |
| processDefinitionId | string | 流程定义ID |
| processDefinitionName | string | 流程名称 |
| version | number | 版本号 |
| deploymentId | string | 部署ID |
| deploymentTime | datetime | 部署时间 |
| deployer | string | 部署人 |
| isActive | boolean | 是否当前版本 |
| status | number | 状态 0-禁用 1-启用 |

**API接口**:
| 方法 | 接口 | 说明 |
|------|------|------|
| GET | /bpm/version-mgmt/list | 获取版本列表 |
| GET | /bpm/version-mgmt/{id} | 获取版本详情 |
| POST | /bpm/version-mgmt/{id}/set-active | 设为当前版本 |
| POST | /bpm/version-mgmt/{id}/enable | 启用版本 |
| POST | /bpm/version-mgmt/{id}/disable | 禁用版本 |
| GET | /bpm/version-mgmt/{id}/bpmn | 获取版本BPMN |

---

## 3. 文件结构

```
mom-web/src/views/bpm/
├── modeler/
│   ├── Modeler.vue            # BPMN设计器
│   └── ModelerToolbar.vue     # 设计器工具栏
├── formDesigner/
│   ├── FormDesigner.vue      # 表单设计器
│   ├── FormPalette.vue       # 表单组件面板
│   └── FormProperties.vue     # 表单属性面板
├── userGroup/
│   ├── UserGroupList.vue     # 用户组列表
│   └── UserGroupEdit.vue     # 用户组编辑
├── assignRule/
│   ├── AssignRuleList.vue    # 分配规则列表
│   └── AssignRuleEdit.vue    # 分配规则编辑
├── taskTransfer/
│   ├── TaskTransferList.vue  # 任务转派列表
│   └── TaskTransferDetail.vue # 任务转派详情
├── taskReturn/
│   ├── TaskReturnList.vue    # 任务退回列表
│   └── TaskReturnDetail.vue  # 任务退回详情
├── approve/
│   ├── ApproveList.vue       # 审批列表
│   └── ApproveDetail.vue     # 审批详情
├── reject/
│   ├── RejectList.vue        # 驳回列表
│   └── RejectDetail.vue      # 驳回详情
├── delegate/
│   ├── DelegateList.vue     # 委托列表
│   └── DelegateEdit.vue      # 委托编辑
├── monitor/
│   ├── MonitorList.vue       # 监控列表
│   └── MonitorDetail.vue     # 监控详情
├── log/
│   ├── LogList.vue           # 日志列表
│   └── LogDetail.vue         # 日志详情
├── version/
│   ├── VersionList.vue       # 版本列表
│   └── VersionCompare.vue    # 版本对比
├── category/
│   ├── CategoryList.vue     # 分类列表
│   └── CategoryEdit.vue     # 分类编辑
├── nodeConfig/
│   ├── NodeConfigList.vue    # 节点配置列表
│   └── NodeConfigEdit.vue    # 节点配置编辑
├── expressionConfig/
│   ├── ExpressionConfigList.vue # 表达式配置列表
│   └── ExpressionConfigEdit.vue # 表达式配置编辑
├── messageConfig/
│   ├── MessageConfigList.vue # 消息配置列表
│   └── MessageConfigEdit.vue # 消息配置编辑
├── timerConfig/
│   ├── TimerConfigList.vue   # 定时配置列表
│   └── TimerConfigEdit.vue   # 定时配置编辑
├── exceptionHandle/
│   ├── ExceptionHandleList.vue # 异常处理列表
│   └── ExceptionHandleEdit.vue # 异常处理编辑
├── simulation/
│   ├── SimulationList.vue   # 仿真列表
│   └── SimulationDetail.vue  # 仿真详情
├── export/
│   ├── ExportList.vue        # 导出列表
│   └── ExportDetail.vue      # 导出详情
├── import/
│   ├── ImportList.vue        # 导入列表
│   └── ImportDetail.vue      # 导入详情
└── examples/
    ├── LeaveExample.vue      # 请假示例
    ├── PurchaseExample.vue   # 采购示例
    └── ExpenseExample.vue    # 报销示例
```

---

## 4. 页面清单

| 序号 | 页面 | 路由路径 | 功能说明 |
|------|------|----------|----------|
| 1 | 流程模型设计器 | `/bpm/modeler` | BPMN可视化设计 |
| 2 | 表单设计器 | `/bpm/formDesigner` | 表单设计、表单组件 |
| 3 | 用户组管理 | `/bpm/userGroup` | 用户组CRUD、组成员 |
| 4 | 任务分配规则 | `/bpm/assignRule` | 分配规则、分配条件 |
| 5 | 任务转派 | `/bpm/taskTransfer` | 转派申请、转派审批 |
| 6 | 任务退回 | `/bpm/taskReturn` | 退回申请、退回审批 |
| 7 | 审批管理 | `/bpm/approve` | 审批意见、审批历史 |
| 8 | 驳回管理 | `/bpm/reject` | 驳回理由、驳回记录 |
| 9 | 委托管理 | `/bpm/delegate` | 委托设置、委托任务 |
| 10 | 流程监控 | `/bpm/monitor` | 流程状态、流程跟踪 |
| 11 | 流程日志 | `/bpm/log` | 操作日志、执行记录 |
| 12 | 流程版本 | `/bpm/version` | 版本管理、版本对比 |
| 13 | 流程分类 | `/bpm/category` | 分类配置、分类管理 |
| 14 | 节点配置 | `/bpm/nodeConfig` | 节点属性、事件配置 |
| 15 | 表达式配置 | `/bpm/expressionConfig` | 条件表达式、脚本配置 |
| 16 | 消息配置 | `/bpm/messageConfig` | 消息事件、通知配置 |
| 17 | 定时配置 | `/bpm/timerConfig` | 定时触发、调度配置 |
| 18 | 异常处理 | `/bpm/exceptionHandle` | 异常捕获、处理规则 |
| 19 | 流程仿真 | `/bpm/simulation` | 仿真测试、流程验证 |
| 20 | 流程导出 | `/bpm/export` | 导出管理 |
| 21 | 流程导入 | `/bpm/import` | 导入管理 |
| 22 | OA请假示例 | `/bpm/leaveExample` | 请假流程 |
| 23 | 采购审批示例 | `/bpm/purchaseExample` | 采购审批 |
| 24 | 报销审批示例 | `/bpm/expenseExample` | 报销审批 |
| 25 | 驳回管理 | `/bpm/rejectMgmt` | 驳回管理 |
| 26 | 委托管理 | `/bpm/delegateMgmt` | 委托管理 |
| 27 | 流程版本 | `/bpm/versionMgmt` | 版本管理 |

---

## 5. 实现优先级

### 高优先级 (MVP必备)
1. 流程模型设计器 - 核心功能
2. 表单设计器 - 流程配置必备
3. 用户组管理 - 任务分配基础
4. 任务分配规则 - 自动化分配
5. 流程监控 - 流程状态跟踪

### 中优先级 (完整流程)
6. 任务转派 - 转派功能
7. 任务退回 - 退回功能
8. 审批管理 - 审批历史
9. 流程日志 - 操作记录
10. 流程分类 - 分类管理

### 低优先级 (增强功能)
11. 流程版本 - 版本管理
12. 节点配置 - 节点属性
13. 表达式配置 - 条件表达式
14. 消息配置 - 通知配置
15. 定时配置 - 定时触发
16. 异常处理 - 异常规则
17. 流程仿真 - 仿真测试
18. 流程导出 - 导出功能
19. 流程导入 - 导入功能
20. OA请假示例 - 业务示例
21. 采购审批示例 - 业务示例
22. 报销审批示例 - 业务示例
23. 驳回管理 - 驳回管理
24. 委托管理 - 委托管理
25. 流程版本 - 版本管理

---

*文档版本: V1.1 | 对比版本: SFMS3.0 | 最后更新: 2026-04-17*
