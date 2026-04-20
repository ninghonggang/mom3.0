# MOM3.0 BPM流程模块设计文档

> **版本**: V1.0
> **日期**: 2026-04-17
> **项目**: 闻荫科技MOM3.0智能制造执行系统
> **流程引擎**: Flowable 6.x

---

## 1. 模块概述

BPM（Business Process Management）流程管理模块基于Flowable 6.x实现流程定义、部署、执行、审批等核心功能，提供可视化的流程设计和管理能力。

**模块路径**: `mom-server/internal/bpm/`

**子模块结构**:
- `mom-server/internal/bpm/api` - API接口定义
- `mom-server/internal/bpm/biz` - 业务实现
- `mom-server/internal/bpm/controller` - 控制器
- `mom-server/internal/bpm/repository` - 数据访问层

---

## 2. 技术架构

### 2.1 技术栈

| 组件 | 技术选型 | 版本 |
|------|----------|------|
| 流程引擎 | Flowable | 6.x |
| 持久层 | MyBatis-Plus | 3.5.x |
| 数据库 | MySQL | 8.0+ |
| 安全 | Spring Security | 6.x |
| 前端框架 | Vue 3 + TypeScript | - |

### 2.2 Flowable核心服务

```java
// 注入Flowable核心服务
@Resource
private RepositoryService repositoryService;    // 流程定义管理
@Resource
private RuntimeService runtimeService;          // 流程运行时
@Resource
private TaskService taskService;                 // 任务管理
@Resource
private HistoryService historyService;           // 历史记录
@Resource
private FormService formService;                 // 表单服务
@Resource
private ManagementService managementService;     // 管理服务
@Resource
private IdentityService identityService;         // 身份服务
```

### 2.3 Flowable数据库表结构

Flowable 6.x数据库表前缀说明：

| 前缀 | 说明 | 示例 |
|------|------|------|
| ACT_RE_* | Repository，流程定义存储 | ACT_RE_DEPLOYMENT, ACT_RE_PROCDEF |
| ACT_RU_* | Runtime，运行时数据 | ACT_RU_EXECUTION, ACT_RU_TASK |
| ACT_HI_* | History，历史数据 | ACT_HI_PROCINST, ACT_HI_TASKINST |
| ACT_GE_* | General，通用数据 | ACT_GE_BYTEARRAY |
| ACT_ID_* | Identity，身份数据 | ACT_ID_USER, ACT_ID_GROUP |
| ACT_CM_* | Case Management | - |

**核心表结构**:

```sql
-- 流程部署表
CREATE TABLE ACT_RE_DEPLOYMENT (
    ID_ VARCHAR(64) NOT NULL PRIMARY KEY,
    NAME_ VARCHAR(255),
    CATEGORY_ VARCHAR(255),
    KEY_ VARCHAR(255),
    TENANT_ID_ VARCHAR(255),
    DEPLOY_TIME_ TIMESTAMP(3),
    LAST_UPDATED_TIME_ TIMESTAMP(3),
    ENGINE_VERSION_ VARCHAR(255)
);

-- 流程定义表
CREATE TABLE ACT_RE_PROCDEF (
    ID_ VARCHAR(64) NOT NULL PRIMARY KEY,
    CATEGORY_ VARCHAR(255),
    NAME_ VARCHAR(255),
    KEY_ VARCHAR(255) NOT NULL,
    VERSION_ INT,
    DEPLOYMENT_ID_ VARCHAR(64),
    RESOURCE_NAME_ VARCHAR(4000),
    DGRM_RESOURCE_NAME_ VARCHAR(4000),
    DESCRIPTION_ VARCHAR(4000),
    HAS_GRAPHICAL_NOTATION_ TINYINT,
    HAS_START_FORM_KEY_ TINYINT,
    SUSPENSION_STATE_ INT,
    ENGINE_VERSION_ VARCHAR(255),
    TENANT_ID_ VARCHAR(255)
);

-- 运行时执行实例表
CREATE TABLE ACT_RU_EXECUTION (
    ID_ VARCHAR(64) NOT NULL PRIMARY KEY,
    REV_ INT,
    PROC_INST_ID_ VARCHAR(64),
    BUSINESS_KEY_ VARCHAR(255),
    PARENT_ID_ VARCHAR(64),
    PROC_DEF_ID_ VARCHAR(64),
    SUPER_EXEC_ID_ VARCHAR(64),
    ACT_ID_ VARCHAR(255),
    IS_ACTIVE_ TINYINT,
    IS_SCOPE_ TINYINT,
    PARENT_SCOPE_ID_ VARCHAR(64),
    NAME_ VARCHAR(255),
    START_TIME_ TIMESTAMP(3),
    START_USER_ID_ VARCHAR(255),
    PRIORITY_ INT,
    TENANT_ID_ VARCHAR(255)
);

-- 运行时任务表
CREATE TABLE ACT_RU_TASK (
    ID_ VARCHAR(64) NOT NULL PRIMARY KEY,
    REV_ INT,
    EXECUTION_ID_ VARCHAR(64),
    PROC_INST_ID_ VARCHAR(64),
    PROC_DEF_ID_ VARCHAR(64),
    NAME_ VARCHAR(255),
    PARENT_TASK_ID_ VARCHAR(64),
    DESCRIPTION_ VARCHAR(4000),
    TASK_DEF_KEY_ VARCHAR(255),
    OWNER_ VARCHAR(255),
    ASSIGNEE_ VARCHAR(255),
    DELEGATION_ VARCHAR(255),
    PRIORITY_ INT,
    CREATE_TIME_ TIMESTAMP(3),
    LAST_UPDATED_TIME_ TIMESTAMP(3),
    DUE_DATE_ TIMESTAMP(3),
    SUSPENSION_STATE_ INT,
    TENANT_ID_ VARCHAR(255)
);

-- 运行时变量表
CREATE TABLE ACT_RU_VARIABLE (
    ID_ VARCHAR(64) NOT NULL PRIMARY KEY,
    REV_ INT,
    TYPE_ VARCHAR(255),
    NAME_ VARCHAR(255),
    EXECUTION_ID_ VARCHAR(64),
    PROC_INST_ID_ VARCHAR(64),
    TASK_ID_ VARCHAR(64),
    BYTEARRAY_ID_ VARCHAR(64),
    DOUBLE_ DOUBLE,
    LONG_ BIGINT,
    TEXT_ VARCHAR(4000),
    TEXT2_ VARCHAR(4000),
    CREATE_TIME_ TIMESTAMP(3),
    LAST_UPDATED_TIME_ TIMESTAMP(3),
    TENANT_ID_ VARCHAR(255)
);

-- 历史流程实例表
CREATE TABLE ACT_HI_PROCINST (
    ID_ VARCHAR(64) NOT NULL PRIMARY KEY,
    PROC_INST_ID_ VARCHAR(64),
    BUSINESS_KEY_ VARCHAR(255),
    PROC_DEF_ID_ VARCHAR(64),
    START_TIME_ TIMESTAMP(3),
    END_TIME_ TIMESTAMP(3),
    START_ACT_ID_ VARCHAR(255),
    END_ACT_ID_ VARCHAR(255),
    START_USER_ID_ VARCHAR(255),
    START_ASSIGNEE_ VARCHAR(255),
    CALLBACK_ID_ VARCHAR(255),
    CALLBACK_TYPE_ VARCHAR(255),
    TENANT_ID_ VARCHAR(255)
);

-- 历史任务表
CREATE TABLE ACT_HI_TASKINST (
    ID_ VARCHAR(64) NOT NULL PRIMARY KEY,
    PROC_DEF_ID_ VARCHAR(64),
    PROC_INST_ID_ VARCHAR(64),
    EXECUTION_ID_ VARCHAR(64),
    NAME_ VARCHAR(255),
    PARENT_TASK_ID_ VARCHAR(64),
    DESCRIPTION_ VARCHAR(4000),
    TASK_DEF_KEY_ VARCHAR(255),
    OWNER_ VARCHAR(255),
    ASSIGNEE_ VARCHAR(255),
    START_TIME_ TIMESTAMP(3),
    CLAIM_TIME_ TIMESTAMP(3),
    END_TIME_ TIMESTAMP(3),
    DURATION_ BIGINT,
    DELETE_REASON_ VARCHAR(4000),
    PRIORITY_ INT,
    DUE_DATE_ TIMESTAMP(3),
    TENANT_ID_ VARCHAR(255)
);

-- 流程资源表(BPMN XML等)
CREATE TABLE ACT_GE_BYTEARRAY (
    ID_ VARCHAR(64) NOT NULL PRIMARY KEY,
    REV_ INT,
    NAME_ VARCHAR(255),
    DESCRIPTION_ VARCHAR(4000),
    TYPE_ INT,
    FILE_ID_ VARCHAR(64),
    FILE_NAME_ VARCHAR(255),
    CREATE_TIME_ TIMESTAMP(3),
    LAST_UPDATED_TIME_ TIMESTAMP(3),
    TENANT_ID_ VARCHAR(255)
);

-- 用户表
CREATE TABLE ACT_ID_USER (
    ID_ VARCHAR(64) NOT NULL PRIMARY KEY,
    REV_ INT,
    TYPE_ VARCHAR(255),
    FIRST_ VARCHAR(255),
    LAST_ VARCHAR(255),
    EMAIL_ VARCHAR(255),
    PASSWORD_ VARCHAR(255),
    EMAIL_ VARCHAR(255),
    PICTURE_ID_ VARCHAR(64),
    TENANT_ID_ VARCHAR(255)
);

-- 用户组表
CREATE TABLE ACT_ID_GROUP (
    ID_ VARCHAR(64) NOT NULL PRIMARY KEY,
    REV_ INT,
    NAME_ VARCHAR(255),
    TYPE_ VARCHAR(255),
    TENANT_ID_ VARCHAR(255)
);

-- 用户组关系表
CREATE TABLE ACT_ID_MEMBERSHIP (
    ID_ VARCHAR(64) NOT NULL PRIMARY KEY,
    USER_ID_ VARCHAR(64),
    GROUP_ID_ VARCHAR(64)
);
```

---

## 3. 核心数据结构

### 3.1 流程表单配置表 bpm_form

```sql
CREATE TABLE bpm_form (
    id              BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '表单ID',
    name            VARCHAR(100) NOT NULL COMMENT '表单名称',
    key             VARCHAR(100) NOT NULL COMMENT '表单Key',
    status          TINYINT NOT NULL DEFAULT 1 COMMENT '状态 0-禁用 1-启用',
    form_type       TINYINT NOT NULL DEFAULT 1 COMMENT '表单类型 1-流程表单 2-业务表单',
    remark          VARCHAR(500) DEFAULT NULL COMMENT '备注',
    form_config     LONGTEXT COMMENT '表单配置JSON',
    create_user     BIGINT UNSIGNED NOT NULL COMMENT '创建用户',
    create_time     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted         TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除',
    INDEX idx_key (key),
    INDEX idx_status (status)
) COMMENT '流程表单配置表';
```

**FormConfig JSON结构**:
```json
{
  "fields": [
    {
      "id": "field_1",
      "type": "input",
      "label": "员工姓名",
      "name": "employeeName",
      "required": true,
      "placeholder": "请输入姓名"
    },
    {
      "id": "field_2",
      "type": "select",
      "label": "请假类型",
      "name": "leaveType",
      "options": [
        {"label": "年假", "value": "ANNUAL"},
        {"label": "病假", "value": "SICK"},
        {"label": "事假", "value": "PERSONAL"}
      ]
    },
    {
      "id": "field_3",
      "type": "dateRange",
      "label": "请假时间",
      "name": "dateRange"
    }
  ],
  "layout": {
    "columns": 1
  }
}
```

### 3.2 任务分配规则表 bpm_task_assign_rule

```sql
CREATE TABLE bpm_task_assign_rule (
    id              BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '规则ID',
    model_id        BIGINT UNSIGNED NOT NULL COMMENT '流程模型ID',
    task_define_key VARCHAR(100) NOT NULL COMMENT '任务定义Key',
    name            VARCHAR(100) NOT NULL COMMENT '规则名称',
    rule_type       TINYINT NOT NULL COMMENT '规则类型 1-指定用户 2-指定角色 3-脚本计算 4-用户组',
    rule_options    LONGTEXT COMMENT '规则选项JSON',
    script          TEXT COMMENT '脚本内容',
    priority        INT NOT NULL DEFAULT 0 COMMENT '优先级',
    status          TINYINT NOT NULL DEFAULT 1 COMMENT '状态 0-禁用 1-启用',
    create_user     BIGINT UNSIGNED NOT NULL COMMENT '创建用户',
    create_time     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted         TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除',
    INDEX idx_model_id (model_id),
    INDEX idx_task_key (task_define_key)
) COMMENT '任务分配规则表';
```

**RuleOptions JSON结构**:
```json
{
  "userIds": [1, 2, 3],
  "roleIds": [10, 20],
  "userGroupId": 100,
  "excludeUserIds": []
}
```

### 3.3 用户组表 bpm_user_group

```sql
CREATE TABLE bpm_user_group (
    id              BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '用户组ID',
    name            VARCHAR(100) NOT NULL COMMENT '用户组名称',
    code            VARCHAR(100) NOT NULL COMMENT '用户组编码',
    status          TINYINT NOT NULL DEFAULT 1 COMMENT '状态 0-禁用 1-启用',
    remark          VARCHAR(500) DEFAULT NULL COMMENT '备注',
    create_user     BIGINT UNSIGNED NOT NULL COMMENT '创建用户',
    create_time     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted         TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除',
    UNIQUE INDEX uk_code (code)
) COMMENT '用户组表';

CREATE TABLE bpm_user_group_member (
    id              BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    group_id        BIGINT UNSIGNED NOT NULL COMMENT '用户组ID',
    user_id         BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    create_time     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    INDEX idx_group_id (group_id),
    INDEX idx_user_id (user_id),
    UNIQUE INDEX uk_group_user (group_id, user_id)
) COMMENT '用户组成员表';
```

### 3.4 OA请假表 bpm_oa_leave

```sql
CREATE TABLE bpm_oa_leave (
    id              BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '请假ID',
    user_id         BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    leave_type      VARCHAR(50) NOT NULL COMMENT '请假类型 ANNUAL-年假 SICK-病假 PERSONAL-事假 MARRIAGE-婚假 MATERNITY-产假 BEREAVEMENT-丧假',
    duration         DECIMAL(10,2) NOT NULL COMMENT '时长(小时)',
    start_time      DATETIME NOT NULL COMMENT '开始时间',
    end_time        DATETIME NOT NULL COMMENT '结束时间',
    reason          VARCHAR(500) DEFAULT NULL COMMENT '请假原因',
    instance_id     VARCHAR(64) DEFAULT NULL COMMENT '流程实例ID',
    result          TINYINT DEFAULT NULL COMMENT '审批结果 1-通过 2-不通过 3-取消',
    process_definition_key VARCHAR(100) DEFAULT NULL COMMENT '流程定义Key',
    form_id         BIGINT UNSIGNED DEFAULT NULL COMMENT '表单ID',
    status          TINYINT NOT NULL DEFAULT 1 COMMENT '状态 0-草稿 1-审批中 2-已完成 3-已取消',
    create_user     BIGINT UNSIGNED NOT NULL COMMENT '创建用户',
    create_time     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted         TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除',
    INDEX idx_user_id (user_id),
    INDEX idx_instance_id (instance_id),
    INDEX idx_create_time (create_time)
) COMMENT 'OA请假记录表';
```

### 3.5 流程定义扩展表 bpm_process_definition_ext

```sql
CREATE TABLE bpm_process_definition_ext (
    id                  BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    process_definition_id VARCHAR(64) NOT NULL COMMENT '流程定义ID',
    description         VARCHAR(500) DEFAULT NULL COMMENT '描述',
    form_type           TINYINT DEFAULT NULL COMMENT '表单类型 1-流程表单 2-自定义表单',
    form_id             BIGINT UNSIGNED DEFAULT NULL COMMENT '表单编号',
    form_custom_create_path VARCHAR(255) DEFAULT NULL COMMENT '自定义创建路径',
    form_custom_view_path VARCHAR(255) DEFAULT NULL COMMENT '自定义查看路径',
    icon                VARCHAR(50) DEFAULT NULL COMMENT '图标',
    color               VARCHAR(50) DEFAULT NULL COMMENT '颜色',
    version             INT DEFAULT NULL COMMENT '版本号',
    create_user         BIGINT UNSIGNED NOT NULL COMMENT '创建用户',
    create_time         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted             TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除',
    UNIQUE INDEX uk_proc_def_id (process_definition_id)
) COMMENT '流程定义扩展表';
```

### 3.6 流程实例扩展表 bpm_process_instance_ext

```sql
CREATE TABLE bpm_process_instance_ext (
    id                  BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    instance_id         VARCHAR(64) NOT NULL COMMENT '流程实例ID',
    business_key        VARCHAR(255) DEFAULT NULL COMMENT '业务Key',
    process_definition_id VARCHAR(64) NOT NULL COMMENT '流程定义ID',
    name                VARCHAR(255) DEFAULT NULL COMMENT '流程名称',
    status              TINYINT NOT NULL DEFAULT 1 COMMENT '状态 1-审批中 2-已完成 3-已取消',
    result              TINYINT DEFAULT NULL COMMENT '结果 1-通过 2-不通过 3-取消',
    start_user_id       BIGINT UNSIGNED DEFAULT NULL COMMENT '发起人ID',
    start_time          DATETIME DEFAULT NULL COMMENT '开始时间',
    end_time            DATETIME DEFAULT NULL COMMENT '结束时间',
    form_id             BIGINT UNSIGNED DEFAULT NULL COMMENT '表单ID',
    biz_no              VARCHAR(100) DEFAULT NULL COMMENT '业务编号',
    create_user         BIGINT UNSIGNED NOT NULL COMMENT '创建用户',
    create_time         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted             TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除',
    UNIQUE INDEX uk_instance_id (instance_id),
    INDEX idx_business_key (business_key),
    INDEX idx_status (status)
) COMMENT '流程实例扩展表';
```

### 3.7 任务扩展表 bpm_task_ext

```sql
CREATE TABLE bpm_task_ext (
    id                  BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    task_id             VARCHAR(64) NOT NULL COMMENT '任务ID',
    instance_id         VARCHAR(64) NOT NULL COMMENT '流程实例ID',
    task_name           VARCHAR(255) DEFAULT NULL COMMENT '任务名称',
    task_def_key        VARCHAR(100) DEFAULT NULL COMMENT '任务定义Key',
    assignee            VARCHAR(64) DEFAULT NULL COMMENT '处理人',
    owner               VARCHAR(64) DEFAULT NULL COMMENT '发起人',
    priority            INT DEFAULT 50 COMMENT '优先级',
    due_date            DATETIME DEFAULT NULL COMMENT '到期时间',
    status              TINYINT NOT NULL DEFAULT 1 COMMENT '状态 1-待处理 2-已处理',
    result              TINYINT DEFAULT NULL COMMENT '处理结果 1-通过 2-驳回 3-转派 4-退回',
    remark              VARCHAR(500) DEFAULT NULL COMMENT '备注',
    create_user         BIGINT UNSIGNED NOT NULL COMMENT '创建用户',
    create_time         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted             TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除',
    UNIQUE INDEX uk_task_id (task_id),
    INDEX idx_instance_id (instance_id)
) COMMENT '任务扩展表';
```

---

## 4. 枚举常量

### 4.1 流程实例状态 BpmProcessInstanceStatusEnum

```java
public enum BpmProcessInstanceStatusEnum {
    RUNNING(1, "审批中"),
    COMPLETED(2, "已完成"),
    CANCELLED(3, "已取消");
}
```

### 4.2 流程实例结果 BpmProcessInstanceResultEnum

```java
public enum BpmProcessInstanceResultEnum {
    APPROVE(1, "通过"),
    REJECT(2, "不通过"),
    CANCEL(3, "取消");
}
```

### 4.3 流程定义状态 BpmProcessDefinitionStatusEnum

```java
public enum BpmProcessDefinitionStatusEnum {
    DISABLED(0, "禁用"),
    ENABLED(1, "启用");
}
```

### 4.4 表单类型 BpmModelFormTypeEnum

```java
public enum BpmModelFormTypeEnum {
    NORMAL(1, "正常表单"),
    CUSTOM(2, "自定义表单");
}
```

### 4.5 任务分配规则类型 BpmTaskAssignRuleTypeEnum

```java
public enum BpmTaskAssignRuleTypeEnum {
    USER(1, "指定用户"),
    ROLE(2, "指定角色"),
    SCRIPT(3, "脚本计算"),
    USER_GROUP(4, "用户组");
}
```

### 4.6 任务结果 BpmTaskResultEnum

```java
public enum BpmTaskResultEnum {
    APPROVE(1, "通过"),
    REJECT(2, "驳回"),
    DELEGATE(3, "转派"),
    RETURN(4, "退回");
}
```

### 4.7 请假类型 LeaveTypeEnum

```java
public enum LeaveTypeEnum {
    ANNUAL("ANNUAL", "年假"),
    SICK("SICK", "病假"),
    PERSONAL("PERSONAL", "事假"),
    MARRIAGE("MARRIAGE", "婚假"),
    MATERNITY("MATERNITY", "产假"),
    BEREAVEMENT("BEREAVEMENT", "丧假");
}
```

---

## 5. API接口设计

### 5.1 流程模型管理 API `/bpm/models`

| 方法 | 路径 | 说明 | 请求体 | 响应 |
|------|------|------|--------|------|
| GET | /bpm/models | 分页查询模型 | Query: pageNo, pageSize, name, key, categoryId, status | PageResult<BpmModelRespVO> |
| GET | /bpm/models/{id} | 获取模型详情 | - | BpmModelRespVO |
| POST | /bpm/models | 创建模型 | BpmModelCreateReqVO | String(id) |
| PUT | /bpm/models/{id} | 更新模型 | BpmModelUpdateReqVO | void |
| DELETE | /bpm/models/{id} | 删除模型 | - | void |
| POST | /bpm/models/{id}/publish | 部署模型 | - | void |
| GET | /bpm/models/{id}/bpmn | 获取BPMN XML | - | String(bpmnXml) |
| PUT | /bpm/models/{id}/bpmn | 更新BPMN XML | String(bpmnXml) | void |

**BpmModelCreateReqVO**:
```java
public class BpmModelCreateReqVO {
    private String name;           // 模型名称
    private String key;           // 模型Key
    private String category;      // 分类
    private String description;   // 描述
    private String bpmnXml;      // BPMN XML
}
```

**BpmModelRespVO**:
```java
public class BpmModelRespVO {
    private String id;
    private String name;
    private String key;
    private String category;
    private Integer version;
    private Integer status;        // 0-草稿 1-已部署
    private String description;
    private String bpmnXml;
    private LocalDateTime createTime;
    private LocalDateTime updateTime;
}
```

### 5.2 表单管理 API `/bpm/forms`

| 方法 | 路径 | 说明 | 请求体 | 响应 |
|------|------|------|--------|------|
| GET | /bpm/forms | 分页查询表单 | Query: pageNo, pageSize, name, key, status | PageResult<BpmFormRespVO> |
| GET | /bpm/forms/{id} | 获取表单详情 | - | BpmFormRespVO |
| POST | /bpm/forms | 创建表单 | BpmFormCreateReqVO | Long(id) |
| PUT | /bpm/forms/{id} | 更新表单 | BpmFormUpdateReqVO | void |
| DELETE | /bpm/forms/{id} | 删除表单 | - | void |

**BpmFormCreateReqVO**:
```java
public class BpmFormCreateReqVO {
    private String name;           // 表单名称
    private String key;            // 表单Key
    private Integer status;        // 状态
    private Integer formType;     // 表单类型
    private String remark;         // 备注
    private String formConfig;    // 表单配置JSON
}
```

**BpmFormRespVO**:
```java
public class BpmFormRespVO {
    private Long id;
    private String name;
    private String key;
    private Integer status;
    private Integer formType;
    private String remark;
    private String formConfig;
    private LocalDateTime createTime;
}
```

### 5.3 用户组管理 API `/bpm/user-group`

| 方法 | 路径 | 说明 | 请求体 | 响应 |
|------|------|------|--------|------|
| GET | /bpm/user-group/list | 列表查询用户组 | Query: name, code, status | List<BpmUserGroupRespVO> |
| GET | /bpm/user-group/{id} | 获取用户组详情 | - | BpmUserGroupRespVO |
| POST | /bpm/user-group | 创建用户组 | BpmUserGroupCreateReqVO | Long(id) |
| PUT | /bpm/user-group/{id} | 更新用户组 | BpmUserGroupUpdateReqVO | void |
| DELETE | /bpm/user-group/{id} | 删除用户组 | - | void |
| GET | /bpm/user-group/{id}/members | 获取组成员 | - | List<BpmUserGroupMemberRespVO> |
| PUT | /bpm/user-group/{id}/members | 更新组成员 | List<Long> userIds | void |

**BpmUserGroupCreateReqVO**:
```java
public class BpmUserGroupCreateReqVO {
    private String name;           // 用户组名称
    private String code;           // 用户组编码
    private Integer status;       // 状态
    private String remark;         // 备注
    private List<Long> memberIds; // 成员ID列表
}
```

**BpmUserGroupMemberRespVO**:
```java
public class BpmUserGroupMemberRespVO {
    private Long userId;
    private String userName;
    private String nickName;
}
```

### 5.4 任务分配规则 API `/bpm/task-assign-rule`

| 方法 | 路径 | 说明 | 请求体 | 响应 |
|------|------|------|--------|------|
| GET | /bpm/task-assign-rule/list | 列表查询规则 | Query: modelId, taskDefineKey | List<BpmTaskAssignRuleRespVO> |
| GET | /bpm/task-assign-rule/{id} | 获取规则详情 | - | BpmTaskAssignRuleRespVO |
| POST | /bpm/task-assign-rule | 创建规则 | BpmTaskAssignRuleCreateReqVO | Long(id) |
| PUT | /bpm/task-assign-rule/{id} | 更新规则 | BpmTaskAssignRuleUpdateReqVO | void |
| DELETE | /bpm/task-assign-rule/{id} | 删除规则 | - | void |

**BpmTaskAssignRuleCreateReqVO**:
```java
public class BpmTaskAssignRuleCreateReqVO {
    private Long modelId;          // 流程模型ID
    private String taskDefineKey;  // 任务定义Key
    private String name;           // 规则名称
    private Integer ruleType;      // 规则类型
    private String ruleOptions;   // 规则选项JSON
    private String script;         // 脚本内容
    private Integer priority;      // 优先级
    private Integer status;        // 状态
}
```

### 5.5 流程实例 API `/bpm/instances`

| 方法 | 路径 | 说明 | 请求体 | 响应 |
|------|------|------|--------|------|
| GET | /bpm/instances | 分页查询实例 | Query: pageNo, pageSize, userId, status | PageResult<BpmProcessInstanceRespVO> |
| GET | /bpm/instances/my-page | 我的申请 | Query: pageNo, pageSize | PageResult<BpmProcessInstanceRespVO> |
| GET | /bpm/instances/{id} | 获取实例详情 | - | BpmProcessInstanceRespVO |
| POST | /bpm/instances | 创建实例 | BpmProcessInstanceCreateReqVO | String(instanceId) |
| POST | /bpm/instances/{id}/cancel | 取消实例 | - | void |
| GET | /bpm/instances/{id}/tasks | 获取实例任务列表 | - | List<BpmTaskRespVO> |
| GET | /bpm/instances/{id}/history | 获取审批历史 | - | List<BpmProcessInstanceHistoryRespVO> |

**BpmProcessInstanceCreateReqVO**:
```java
public class BpmProcessInstanceCreateReqVO {
    private String processDefinitionKey;  // 流程定义Key
    private String businessKey;          // 业务Key
    private Long formId;                 // 表单ID
    private Map<String, Object> variables; // 流程变量
    private Map<String, Object> formData;  // 表单数据
}
```

**BpmProcessInstanceRespVO**:
```java
public class BpmProcessInstanceRespVO {
    private String id;
    private String businessKey;
    private String processDefinitionId;
    private String processDefinitionName;
    private String name;
    private Integer status;
    private Integer result;
    private Long startUserId;
    private String startUserName;
    private LocalDateTime startTime;
    private LocalDateTime endTime;
    private String bizNo;
    private List<String> currentNodeNames;
}
```

### 5.6 任务 API `/bpm/tasks`

| 方法 | 路径 | 说明 | 请求体 | 响应 |
|------|------|------|--------|------|
| GET | /bpm/tasks/todo-page | 待办任务分页 | Query: pageNo, pageSize, userId | PageResult<BpmTaskRespVO> |
| GET | /bpm/tasks/done-page | 已办任务分页 | Query: pageNo, pageSize, userId | PageResult<BpmTaskRespVO> |
| GET | /bpm/tasks/{id} | 获取任务详情 | - | BpmTaskRespVO |
| POST | /bpm/tasks/{id}/approve | 审批通过 | BpmTaskApproveReqVO | void |
| POST | /bpm/tasks/{id}/reject | 审批驳回 | BpmTaskRejectReqVO | void |
| POST | /bpm/tasks/{id}/return | 任务退回 | BpmTaskReturnReqVO | void |
| POST | /bpm/tasks/{id}/delegate | 任务转派 | BpmTaskDelegateReqVO | void |

**BpmTaskApproveReqVO**:
```java
public class BpmTaskApproveReqVO {
    private String comment;           // 审批意见
    private Map<String, Object> variables; // 流程变量
}
```

**BpmTaskReturnReqVO**:
```java
public class BpmTaskReturnReqVO {
    private String targetTaskId;      // 目标任务ID
    private String reason;            // 退回原因
}
```

**BpmTaskDelegateReqVO**:
```java
public class BpmTaskDelegateReqVO {
    private String assignee;          // 新处理人ID
    private String reason;            // 转派原因
}
```

**BpmTaskRespVO**:
```java
public class BpmTaskRespVO {
    private String id;
    private String name;
    private String taskDefinitionKey;
    private String assignee;
    private String assigneeName;
    private String owner;
    private String ownerName;
    private String instanceId;
    private String processDefinitionName;
    private Integer priority;
    private LocalDateTime createTime;
    private LocalDateTime dueDate;
    private String comment;
}
```

### 5.7 OA请假 API `/bpm/oa/leave`

| 方法 | 路径 | 说明 | 请求体 | 响应 |
|------|------|------|--------|------|
| GET | /bpm/oa/leave | 分页查询请假 | Query: pageNo, pageSize, userId, leaveType, status | PageResult<BpmOALeaveRespVO> |
| GET | /bpm/oa/leave/{id} | 获取请假详情 | - | BpmOALeaveRespVO |
| POST | /bpm/oa/leave | 创建请假 | BpmOALeaveCreateReqVO | Long(id) |
| PUT | /bpm/oa/leave/{id} | 更新请假 | BpmOALeaveUpdateReqVO | void |
| DELETE | /bpm/oa/leave/{id} | 删除请假 | - | void |
| POST | /bpm/oa/leave/{id}/cancel | 取消请假 | - | void |

**BpmOALeaveCreateReqVO**:
```java
public class BpmOALeaveCreateReqVO {
    private String leaveType;         // 请假类型
    private BigDecimal duration;      // 时长(小时)
    private LocalDateTime startTime;  // 开始时间
    private LocalDateTime endTime;   // 结束时间
    private String reason;            // 请假原因
}
```

**BpmOALeaveRespVO**:
```java
public class BpmOALeaveRespVO {
    private Long id;
    private Long userId;
    private String userName;
    private String leaveType;
    private String leaveTypeName;
    private BigDecimal duration;
    private LocalDateTime startTime;
    private LocalDateTime endTime;
    private String reason;
    private String instanceId;
    private Integer result;
    private String resultName;
    private Integer status;
    private LocalDateTime createTime;
}
```

---

## 6. 核心服务类设计

### 6.1 BpmModelService

```java
public interface BpmModelService {
    /**
     * 分页查询流程模型
     */
    PageResult<BpmModelRespVO> getModelPage(BpmModelPageReqVO pageReqVO);

    /**
     * 获取流程模型详情
     */
    BpmModelRespVO getModel(String id);

    /**
     * 创建流程模型
     */
    String createModel(BpmModelCreateReqVO createReqVO);

    /**
     * 更新流程模型
     */
    void updateModel(BpmModelUpdateReqVO updateReqVO);

    /**
     * 删除流程模型
     */
    void deleteModel(String id);

    /**
     * 部署流程模型
     */
    void deployModel(String id);

    /**
     * 获取BPMN模型
     */
    BpmnModel getBpmnModel(String id);

    /**
     * 更新BPMN XML
     */
    void updateModelBpmnXml(String id, String bpmnXml);
}
```

### 6.2 BpmProcessInstanceService

```java
public interface BpmProcessInstanceService {
    /**
     * 获取流程实例
     */
    ProcessInstance getProcessInstance(String id);

    /**
     * 分页查询流程实例
     */
    PageResult<BpmProcessInstanceRespVO> getProcessInstancePage(BpmProcessInstancePageReqVO pageReqVO);

    /**
     * 获取我的申请分页
     */
    PageResult<BpmProcessInstanceRespVO> getMyProcessInstancePage(Long userId, BpmProcessInstancePageReqVO pageReqVO);

    /**
     * 创建流程实例
     */
    String createProcessInstance(Long userId, BpmProcessInstanceCreateReqVO createReqVO);

    /**
     * 获取流程实例详情
     */
    BpmProcessInstanceRespVO getProcessInstanceVO(String id);

    /**
     * 取消流程实例
     */
    void cancelProcessInstance(Long userId, String id);

    /**
     * 获取实例任务列表
     */
    List<BpmTaskRespVO> getInstanceTasks(String id);

    /**
     * 获取审批历史
     */
    List<BpmProcessInstanceHistoryRespVO> getInstanceHistory(String id);
}
```

### 6.3 BpmTaskService

```java
public interface BpmTaskService {
    /**
     * 获取待办任务分页
     */
    PageResult<BpmTaskRespVO> getTodoTaskPage(Long userId, BpmTaskPageReqVO pageReqVO);

    /**
     * 获取已办任务分页
     */
    PageResult<BpmTaskRespVO> getDoneTaskPage(Long userId, BpmTaskPageReqVO pageReqVO);

    /**
     * 获取任务详情
     */
    BpmTaskRespVO getTask(String id);

    /**
     * 审批通过
     */
    void approveTask(Long userId, String taskId, BpmTaskApproveReqVO approveReqVO);

    /**
     * 审批驳回
     */
    void rejectTask(Long userId, String taskId, BpmTaskRejectReqVO rejectReqVO);

    /**
     * 任务退回
     */
    void returnTask(Long userId, String taskId, BpmTaskReturnReqVO returnReqVO);

    /**
     * 任务转派
     */
    void delegateTask(Long userId, String taskId, BpmTaskDelegateReqVO delegateReqVO);
}
```

### 6.4 BpmFormService

```java
public interface BpmFormService {
    /**
     * 分页查询表单
     */
    PageResult<BpmFormRespVO> getFormPage(BpmFormPageReqVO pageReqVO);

    /**
     * 获取表单详情
     */
    BpmFormRespVO getForm(Long id);

    /**
     * 创建表单
     */
    Long createForm(BpmFormCreateReqVO createReqVO, Long userId);

    /**
     * 更新表单
     */
    void updateForm(BpmFormUpdateReqVO updateReqVO, Long userId);

    /**
     * 删除表单
     */
    void deleteForm(Long id);

    /**
     * 渲染表单
     */
    String renderForm(Long id);
}
```

---

## 7. 数据流向

### 7.1 流程创建流程

```
用户发起流程申请
        ↓
BpmProcessInstanceController.createProcessInstance()
        ↓
BpmProcessInstanceService.createProcessInstance()
        ↓
Flowable RuntimeService.startProcessInstanceByKey()
        ↓
BpmProcessInstanceService.createProcessInstanceExt()  // 创建扩展记录
        ↓
BpmMessageService.sendMessageWhenProcessInstanceCreated()  // 发送消息
        ↓
返回流程实例ID
```

### 7.2 流程审批流程

```
用户提交审批
        ↓
BpmTaskController.approveTask()
        ↓
Flowable TaskService.complete()
        ↓
┌─────────────────────────────────────┐
│ 事件监听:                            │
│   BpmProcessInstanceResultEvent      │
│   → 审批通过/驳回后更新扩展记录      │
└─────────────────────────────────────┘
        ↓
BpmMessageService 发送审批结果通知
```

### 7.3 任务分配计算流程

```
任务创建时
        ↓
BpmTaskAssignRuleCalculator.calculate(taskDefinitionKey, variables)
        ↓
遍历规则优先级
        ↓
┌─────────────────────────────────────┐
│ 规则类型判断:                        │
│   USER → 直接返回用户ID              │
│   ROLE → 查询角色对应用户            │
│   SCRIPT → 执行脚本计算              │
│   USER_GROUP → 查询用户组成员        │
└─────────────────────────────────────┘
        ↓
返回最终处理人列表
```

---

## 8. 前端组件设计

### 8.1 文件结构

```
mom-web/src/views/bpm/
├── ProcessList.vue           # 流程定义列表 (现有)
├── InstanceList.vue          # 流程实例列表 (现有)
├── TaskList.vue              # 任务列表 (现有)
├── ModelList.vue             # [新增] 流程模型列表
├── ModelEditor.vue           # [新增] 模型编辑器
├── FormList.vue              # [新增] 表单管理列表
├── FormEditor.vue            # [新增] 表单设计器
├── UserGroupList.vue         # [新增] 用户组管理
├── TaskAssignRuleList.vue    # [新增] 任务分配规则
├── InstanceCreate.vue        # [新增] 流程实例创建
├── InstanceDetail.vue        # [新增] 流程实例详情
├── InstanceBpmnViewer.vue    # [新增] BPMN图形查看器
├── TaskReturnDialogForm.vue  # [新增] 任务退回表单
├── TaskUpdateAssigneeForm.vue # [新增] 任务转派表单
├── TaskDoneList.vue          # [新增] 已办任务列表
└── oa/
    └── OALeaveList.vue       # [新增] OA请假申请
```

### 8.2 BPMN可视化设计器

**技术选型**: bpmn-js

**核心功能**:
- 拖拽式流程节点设计
- 节点属性配置（名称、 assignee、candidateUsers等）
- 连线配置（条件表达式、流转方向）
- 表单绑定
- 边界事件配置
- 会签/或签支持

**组件引用**:
```bash
npm install bpmn-js bpmn-js-properties-panel
```

**ModelEditor.vue 核心实现**:
```vue
<template>
  <div class="model-editor">
    <div class="model-editor__palette">
      <!-- 左侧工具栏 -->
    </div>
    <div class="model-editor__canvas" ref="canvasRef">
      <!-- BPMN画布 -->
    </div>
    <div class="model-editor__properties" v-if="selectedElement">
      <!-- 属性面板 -->
    </div>
  </div>
</template>

<script setup>
import BpmnModeler from 'bpmn-js/lib/Modeler';
import PropertiesPanel from 'bpmn-js-properties-panel';
import 'bpmn-js/dist/assets/diagram-js.css';
import 'bpmn-js/dist/assets/bpmn-js.css';
import 'bpmn-js-properties-panel/dist/assets/properties-panel.css';
</script>
```

---

## 9. 与其他模块的交互

| 模块 | 交互方式 | 说明 |
|------|----------|------|
| System模块 | 本地调用 | 用户信息、权限验证 |
| 业务模块 | BpmProcessInstanceApi | 被其他模块调用创建流程 |
| 消息模块 | 事件监听 | 流程审批通知推送 |
| 前端 | REST API | 前后端分离交互 |

---

## 10. 实现优先级

### P0 核心内容 (MVP必备)

| 序号 | 功能 | 说明 |
|------|------|------|
| 1 | Flowable引擎集成 | 配置、数据表、核心服务 |
| 2 | 流程模型管理 | CRUD、部署 |
| 3 | 流程实例管理 | 创建、查询、取消 |
| 4 | 任务管理 | 待办、已办、审批 |
| 5 | 表单配置表 bpm_form | DDL + API |
| 6 | 任务分配规则表 bpm_task_assign_rule | DDL + API |
| 7 | BPMN可视化设计器 | 前端组件 |

### P1 重要内容

| 序号 | 功能 | 说明 |
|------|------|------|
| 8 | 用户组表 bpm_user_group | DDL + API |
| 9 | OA请假表 bpm_oa_leave | DDL + API |
| 10 | 任务退回API | 退回至指定节点 |
| 11 | 任务转派API | 转派至其他人 |
| 12 | 表单管理API | 表单CRUD |
| 13 | 流程模型管理API | 完整CRUD+部署 |

---

---

## 11. 补充数据表结构

### 11.1 任务消息规则表 bpm_task_message_rule

用于配置任务催办、审批结果通知等消息推送规则。

```sql
CREATE TABLE bpm_task_message_rule (
    id                  BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '规则ID',
    task_id             VARCHAR(64) NOT NULL COMMENT '任务定义Key或*表示全部任务',
    message_type        TINYINT NOT NULL COMMENT '消息类型 1-任务催办 2-审批结果通知 3-会签提醒 4-转派通知 5-退回通知',
    trigger_condition   VARCHAR(50) DEFAULT NULL COMMENT '触发条件如 BEFORE_DUE(到期前) AFTER_COMPLETE(完成之后) DURING(进行中)',
    template_code       VARCHAR(100) DEFAULT NULL COMMENT '消息模板编码',
    enabled             TINYINT NOT NULL DEFAULT 1 COMMENT '是否启用 0-禁用 1-启用',
    creator             BIGINT UNSIGNED DEFAULT NULL COMMENT '创建者',
    create_time         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updater             BIGINT UNSIGNED DEFAULT NULL COMMENT '更新者',
    update_time         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    tenant_id           BIGINT UNSIGNED DEFAULT NULL COMMENT '租户ID',
    deleted             TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除',
    INDEX idx_task_id (task_id),
    INDEX idx_message_type (message_type),
    INDEX idx_enabled (enabled)
) COMMENT '任务消息规则表';
```

---

## 12. 补充API接口设计

### 12.1 BpmProcessInstanceApi - 跨模块API接口

供其他业务模块（如WMS、Quality等）调用，用于在业务操作时发起流程审批。

**接口路径**: `mom-server/internal/bpm/api/BpmProcessInstanceApi.java`

```java
public interface BpmProcessInstanceApi {

    /**
     * 创建流程实例
     * 供其他模块调用，在业务数据提交后触发流程审批
     *
     * @param userId        发起人用户ID
     * @param createReqDTO   创建请求参数
     * @return 流程实例ID
     */
    String createProcessInstance(Long userId, BpmProcessInstanceCreateReqDTO createReqDTO);

    /**
     * 获取流程实例详情
     *
     * @param id 流程实例ID
     * @return 流程实例响应DTO
     */
    BpmProcessInstanceRespDTO getProcessInstance(String id);

    /**
     * 取消流程实例
     *
     * @param id     流程实例ID
     * @param userId 操作用户ID
     */
    void cancelProcessInstance(String id, Long userId);

    /**
     * 获取流程实例状态
     *
     * @param instanceId 流程实例ID
     * @return 实例状态
     */
    BpmProcessInstanceStatusEnum getProcessInstanceStatus(String instanceId);
}
```

**BpmProcessInstanceCreateReqDTO**:
```java
public class BpmProcessInstanceCreateReqDTO {
    private String processDefinitionKey;   // 流程定义Key（必填）
    private String businessKey;            // 业务Key（可选，用于关联业务表）
    private Long formId;                    // 表单ID
    private Map<String, Object> variables; // 流程变量
    private Map<String, Object> formData;   // 表单数据
    private String bizNo;                  // 业务编号
}
```

**BpmProcessInstanceRespDTO**:
```java
public class BpmProcessInstanceRespDTO {
    private String id;                      // 流程实例ID
    private String businessKey;              // 业务Key
    private String processDefinitionId;      // 流程定义ID
    private String processDefinitionName;   // 流程定义名称
    private Integer status;                  // 状态 1-审批中 2-已完成 3-已取消
    private Integer result;                  // 结果 1-通过 2-不通过 3-取消
    private Long startUserId;                // 发起人ID
    private String startUserName;            // 发起人名称
    private LocalDateTime startTime;         // 开始时间
    private LocalDateTime endTime;           // 结束时间
    private String bizNo;                    // 业务编号
}
```

---

### 12.2 BpmMessageService - 流程消息通知服务

负责流程各环节的消息推送，包括任务催办、审批结果、会签提醒等。

**接口路径**: `mom-server/internal/bpm/biz/message/BpmMessageService.java`

```java
public interface BpmMessageService {

    /**
     * 发送任务催办提醒
     * 定时任务扫描到期前未完成的任务，发送催办通知
     *
     * @param taskId 任务ID
     */
    void sendTaskReminder(String taskId);

    /**
     * 发送审批结果通知
     * 流程审批完成后，通知申请人审批结果
     *
     * @param processInstanceId 流程实例ID
     * @param result            审批结果（1-通过 2-驳回 3-取消）
     */
    void sendApprovalResult(String processInstanceId, Integer result);

    /**
     * 发送会签催办提醒
     * 会签任务中各签署人尚未完成时，发送提醒
     *
     * @param taskId 任务ID
     */
    void sendCounterSignReminder(String taskId);

    /**
     * 发送任务转移通知
     * 任务转派后通知原处理人和新处理人
     *
     * @param taskId        任务ID
     * @param fromUserId    原处理人ID
     * @param toUserId      新处理人ID
     * @param reason        转派原因
     */
    void sendTaskTransferNotification(String taskId, Long fromUserId, Long toUserId, String reason);

    /**
     * 发送任务退回通知
     * 任务被退回后通知发起人
     *
     * @param taskId            任务ID
     * @param processInstanceId 流程实例ID
     * @param returnReason      退回原因
     */
    void sendTaskReturnNotification(String taskId, String processInstanceId, String returnReason);

    /**
     * 流程实例创建后发送通知
     * 通知下一节点处理人有新的待办任务
     *
     * @param processInstanceId 流程实例ID
     * @param taskId            任务ID
     */
    void sendTaskCreatedNotification(String processInstanceId, String taskId);
}
```

---

### 12.3 BpmProcessInstanceResultEventListener - 审批结果事件监听器

监听流程实例审批完成事件，更新业务表状态（由Flowable的事件监听机制触发）。

**接口路径**: `mom-server/internal/bpm/biz/listener/BpmProcessInstanceResultEventListener.java`

```java
@Component
public class BpmProcessInstanceResultEventListener {

    @Resource
    private BpmProcessInstanceResultInvoker resultInvoker;

    /**
     * 流程审批通过后处理
     * 更新业务表状态为"已通过"，触发后续业务逻辑
     *
     * @param processInstanceId 流程实例ID
     */
    @EventListener
    public void onProcessInstanceApproved(String processInstanceId) {
        // 1. 更新 bpm_process_instance_ext 表中实例状态
        // 2. 根据 business_key 查找关联业务表，调用业务模块更新状态
        // 3. 发送审批通过通知
        resultInvoker.updateBusinessStatus(processInstanceId, BpmProcessInstanceResultEnum.APPROVE);
    }

    /**
     * 流程审批驳回后处理
     * 更新业务表状态为"已驳回"，记录驳回原因
     *
     * @param processInstanceId 流程实例ID
     */
    @EventListener
    public void onProcessInstanceRejected(String processInstanceId) {
        // 1. 更新 bpm_process_instance_ext 表中实例状态
        // 2. 根据 business_key 查找关联业务表，调用业务模块更新状态
        // 3. 发送审批驳回通知
        resultInvoker.updateBusinessStatus(processInstanceId, BpmProcessInstanceResultEnum.REJECT);
    }

    /**
     * 流程取消后处理
     *
     * @param processInstanceId 流程实例ID
     */
    @EventListener
    public void onProcessInstanceCancelled(String processInstanceId) {
        resultInvoker.updateBusinessStatus(processInstanceId, BpmProcessInstanceResultEnum.CANCEL);
    }
}
```

**BpmProcessInstanceResultInvoker**（结果处理调用器）:
```java
@Component
public class BpmProcessInstanceResultInvoker {

    /**
     * 更新业务表状态
     * 根据流程实例关联的 business_key 类型，调用对应业务模块更新状态
     *
     * @param processInstanceId 流程实例ID
     * @param result            审批结果
     */
    public void updateBusinessStatus(String processInstanceId, BpmProcessInstanceResultEnum result) {
        // 1. 根据 processInstanceId 查询 bpm_process_instance_ext
        // 2. 获取 business_key，判断业务类型前缀（如 OA_LEAVE_、WMS_TRANSFER_）
        // 3. 根据业务类型调用对应业务Service更新状态
        //    OA_LEAVE_*   → BpmOALeaveResultListener.onLeaveApproved/Rejected
        //    WMS_*        → WMS模块监听器
        //    QUALITY_*    → Quality模块监听器
    }
}
```

---

### 12.4 BpmProcessDefinitionExtService - 流程定义扩展服务

管理流程定义的扩展属性（表单、图标、描述等）。

**接口路径**: `mom-server/internal/bpm/biz/BpmProcessDefinitionExtService.java`

```java
public interface BpmProcessDefinitionExtService {

    /**
     * 获取流程定义及扩展信息
     *
     * @param id 流程定义ID
     * @return 包含扩展信息的流程定义响应DTO
     */
    BpmProcessDefinitionRespVO getProcessDefinitionWithExt(String id);

    /**
     * 更新流程定义扩展信息
     *
     * @param id     流程定义ID
     * @param extDTO 扩展信息更新请求DTO
     */
    void updateProcessDefinitionExt(String id, BpmProcessDefinitionExtUpdateReqDTO extDTO);

    /**
     * 获取流程定义扩展信息
     *
     * @param processDefinitionId 流程定义ID
     * @return 扩展信息DTO
     */
    BpmProcessDefinitionExtRespDTO getProcessDefinitionExt(String processDefinitionId);

    /**
     * 创建流程定义扩展记录
     *
     * @param processDefinitionId 流程定义ID
     * @param createDTO            创建参数
     * @return 扩展记录ID
     */
    Long createProcessDefinitionExt(String processDefinitionId, BpmProcessDefinitionExtCreateReqDTO createDTO);
}
```

**BpmProcessDefinitionExtUpdateReqDTO**:
```java
public class BpmProcessDefinitionExtUpdateReqDTO {
    private String description;              // 描述
    private Integer formType;                // 表单类型 1-流程表单 2-自定义表单
    private Long formId;                     // 表单ID
    private String formCustomCreatePath;    // 自定义创建路径
    private String formCustomViewPath;      // 自定义查看路径
    private String icon;                     // 图标
    private String color;                    // 颜色
}
```

---

### 12.5 任务转移/移交 API `/bpm/task/transfer`

将任务从当前处理人转移至其他用户，常用于岗位调整、代理审批等场景。

**接口路径**: `mom-server/internal/bpm/controller/BpmTaskTransferController.java`

| 方法 | 路径 | 说明 | 请求体 | 响应 |
|------|------|------|--------|------|
| POST | /bpm/task/transfer | 任务转移/移交 | BpmTaskTransferReqVO | void |

**BpmTaskTransferReqVO**:
```java
public class BpmTaskTransferReqVO {
    private String taskId;         // 任务ID
    private String targetUserId;  // 目标处理人ID
    private String reason;        // 转移原因
    private String comment;       // 备注说明
}
```

**接口说明**:
```
1. 校验任务是否存在且属于当前用户
2. 校验目标处理人是否为有效用户
3. 调用 Flowable TaskService.delegateTask() 进行任务转派
4. 创建任务转移记录 bpm_task_transfer_log
5. 发送转派通知给原处理人和目标处理人
```

---

### 12.6 脚本计算任务分配 API `/bpm/task/calculateAssignee`

用于脚本动态计算任务的处理人，支持在流程执行过程中根据业务数据确定审批人。

**接口路径**: `mom-server/internal/bpm/controller/BpmTaskAssigneeCalculateController.java`

| 方法 | 路径 | 说明 | 请求体 | 响应 |
|------|------|------|--------|------|
| POST | /bpm/task/calculateAssignee | 脚本计算任务处理人 | BpmTaskCalculateAssigneeReqVO | BpmTaskCalculateAssigneeRespVO |
| POST | /bpm/task/calculateAssignee/preview | 预览脚本计算结果 | BpmTaskCalculateAssigneeReqVO | List<Long> userIds |

**BpmTaskCalculateAssigneeReqVO**:
```java
public class BpmTaskCalculateAssigneeReqVO {
    private String processDefinitionId;      // 流程定义ID
    private String taskDefinitionKey;        // 任务定义Key
    private String script;                   // 脚本内容（支持Groovy）
    private Map<String, Object> variables;   // 流程变量
    private Map<String, Object> bizData;     // 业务数据
}
```

**BpmTaskCalculateAssigneeRespVO**:
```java
public class BpmTaskCalculateAssigneeRespVO {
    private List<Long> assigneeIds;          // 计算出的处理人ID列表
    private String assigneeNames;            // 处理人名称（逗号分隔）
    private String executedScript;           // 实际执行的脚本
    private Long executeTimeMs;             // 执行耗时（毫秒）
}
```

**脚本计算示例**（Groovy）:
```groovy
// 根据部门经理和金额范围计算审批人
def amount = variables.get("amount") as BigDecimal
def deptId = bizData.get("deptId") as Long

if (amount > 100000) {
    // 金额大于10万，需要总监审批
    return userService.getUsersByRole("director")
} else {
    // 其他由部门经理审批
    return userService.getDeptManager(deptId)
}
```

---

### 12.7 候选人/候选组 API `/bpm/task/candidate`

获取任务当前候选人（candidateUsers）和候选组（candidateGroups），支持多人会签/或签场景。

**接口路径**: `mom-server/internal/bpm/controller/BpmTaskCandidateController.java`

| 方法 | 路径 | 说明 | 请求体 | 响应 |
|------|------|------|--------|------|
| GET | /bpm/task/candidateUsers | 获取任务候选人 | Query: taskId | List<BpmTaskCandidateUserRespVO> |
| GET | /bpm/task/candidateGroups | 获取任务候选组 | Query: taskId | List<BpmTaskCandidateGroupRespVO> |

**BpmTaskCandidateUserRespVO**:
```java
public class BpmTaskCandidateUserRespVO {
    private Long userId;        // 用户ID
    private String userName;    // 用户名
    private String nickName;   // 昵称
    private String deptName;   // 部门名称
    private String roleNames;  // 角色名称
}
```

**BpmTaskCandidateGroupRespVO**:
```java
public class BpmTaskCandidateGroupRespVO {
    private Long groupId;       // 用户组ID
    private String groupName;  // 用户组名称
    private String groupCode;  // 用户组编码
    private Integer memberCount; // 组成员数量
}
```

---

### 12.8 BpmOALeaveResultListener - OA请假结果监听器

监听OA请假流程的审批结果，更新请假记录状态。

**接口路径**: `mom-server/internal/bpm/biz/listener/BpmOALeaveResultListener.java`

```java
@Component
public class BpmOALeaveResultListener implements BpmProcessInstanceResultListener {

    @Resource
    private BpmOALeaveRepository leaveRepository;

    @Resource
    private MessageService messageService;

    /**
     * 请假审批通过后处理
     * 更新请假记录状态和审批结果
     *
     * @param leaveId  请假记录ID
     * @param result   审批结果（1-通过 2-驳回 3-取消）
     */
    @Override
    public void onLeaveApproved(Long leaveId, Integer result) {
        BpmOALeave leave = leaveRepository.getById(leaveId);
        if (leave == null) {
            return;
        }

        // 更新请假记录状态
        leave.setStatus(BpmProcessInstanceStatusEnum.COMPLETED.getStatus());
        leave.setResult(result);
        leave.setUpdateTime(LocalDateTime.now());
        leaveRepository.update(leave);

        // 发送通知给申请人
        messageService.sendMessage(MessageEnum.BPM_LEAVE_APPROVED, leave.getUserId(),
            buildTemplateData(leave));

        // 记录审批日志
        log.info("请假审批完成: leaveId={}, result={}", leaveId, result);
    }

    /**
     * 请假审批驳回后处理
     */
    @Override
    public void onLeaveRejected(Long leaveId, String rejectReason) {
        BpmOALeave leave = leaveRepository.getById(leaveId);
        if (leave == null) {
            return;
        }

        leave.setStatus(BpmProcessInstanceStatusEnum.COMPLETED.getStatus());
        leave.setResult(BpmProcessInstanceResultEnum.REJECT.getResult());
        leave.setRemark(rejectReason);
        leave.setUpdateTime(LocalDateTime.now());
        leaveRepository.update(leave);

        // 发送驳回通知
        messageService.sendMessage(MessageEnum.BPM_LEAVE_REJECTED, leave.getUserId(),
            buildTemplateData(leave));
    }

    /**
     * 请假取消后处理
     */
    @Override
    public void onLeaveCancelled(Long leaveId) {
        BpmOALeave leave = leaveRepository.getById(leaveId);
        if (leave == null) {
            return;
        }

        leave.setStatus(BpmProcessInstanceStatusEnum.CANCELLED.getStatus());
        leave.setResult(BpmProcessInstanceResultEnum.CANCEL.getResult());
        leave.setUpdateTime(LocalDateTime.now());
        leaveRepository.update(leave);
    }

    private Map<String, Object> buildTemplateData(BpmOALeave leave) {
        Map<String, Object> data = new HashMap<>();
        data.put("leaveType", leave.getLeaveType());
        data.put("startTime", leave.getStartTime());
        data.put("endTime", leave.getEndTime());
        data.put("duration", leave.getDuration());
        data.put("reason", leave.getReason());
        return data;
    }
}
```

---

### 12.9 任务扩展API补充

| 方法 | 路径 | 说明 | 请求体 | 响应 |
|------|------|------|--------|------|
| GET | /bpm/task/candidate | 获取任务候选人列表 | Query: taskId | List<BpmTaskCandidateUserRespVO> |
| POST | /bpm/task/calculateAssignee | 脚本计算任务处理人 | BpmTaskCalculateAssigneeReqVO | BpmTaskCalculateAssigneeRespVO |
| POST | /bpm/task/calculateAssignee/preview | 预览脚本计算结果 | BpmTaskCalculateAssigneeReqVO | List<Long> userIds |
| POST | /bpm/task/transfer | 任务转移/移交 | BpmTaskTransferReqVO | void |
| GET | /bpm/task/assignRule/{taskDefKey} | 获取任务分配规则 | Query: processDefinitionId | BpmTaskAssignRuleRespVO |

---

## 13. 补充枚举常量

### 13.1 消息类型 BpmMessageTypeEnum

```java
public enum BpmMessageTypeEnum {
    TASK_REMINDER(1, "任务催办"),
    APPROVAL_RESULT(2, "审批结果通知"),
    COUNTERSIGN_REMINDER(3, "会签提醒"),
    TRANSFER_NOTIFY(4, "转派通知"),
    RETURN_NOTIFY(5, "退回通知"),
    TASK_CREATED(6, "新任务通知");
}
```

### 13.2 消息触发条件 BpmMessageTriggerConditionEnum

```java
public enum BpmMessageTriggerConditionEnum {
    BEFORE_DUE("BEFORE_DUE", "到期前"),
    AFTER_COMPLETE("AFTER_COMPLETE", "完成后"),
    DURING("DURING", "进行中"),
    ON_CREATE("ON_CREATE", "创建时"),
    ON_ASSIGN("ON_ASSIGN", "分配时");
}
```

### 13.3 流程实例状态 BpmProcessInstanceStatusEnum（完整版）

```java
public enum BpmProcessInstanceStatusEnum {
    RUNNING(1, "审批中"),
    COMPLETED(2, "已完成"),
    CANCELLED(3, "已取消");

    private final Integer status;
    private final String desc;

    public Integer getStatus() { return status; }
    public String getDesc() { return desc; }
}
```

---

*文档版本: V1.1 | 最后更新: 2026-04-17*