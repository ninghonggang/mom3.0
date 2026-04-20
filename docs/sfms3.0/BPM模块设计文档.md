# BPM流程管理模块设计文档

## 1. 模块概述

BPM（Business Process Management）流程管理模块是SFMS3.0的工作流引擎模块，基于Flowable实现流程定义、部署、执行、审批等核心功能，提供可视化的流程设计和管理能力。

**路径**: `win-module-bpm`

**子模块结构**:
- `win-module-bpm-api` - API接口定义
- `win-module-bpm-biz` - 业务实现

---

## 2. 模块职责

| 子模块 | 职责 |
|--------|------|
| 流程定义 (definition) | 流程模型管理、表单定义、流程定义、任务分配规则、用户组 |
| 流程任务 (task) | 流程实例管理、任务查询与审批、活动节点查询 |
| OA业务 (oa) | 集成示例：OA请假流程 |
| 消息通知 (message) | 流程审批消息推送 |

---

## 3. 核心类/接口

### 3.1 流程实例服务

**Service接口**: `BpmProcessInstanceService`
```java
public interface BpmProcessInstanceService {
    ProcessInstance getProcessInstance(String id);
    List<ProcessInstance> getProcessInstances(Set<String> ids);
    PageResult<BpmProcessInstancePageItemRespVO> getMyProcessInstancePage(Long userId, BpmProcessInstanceMyPageReqVO pageReqVO);
    String createProcessInstance(Long userId, BpmProcessInstanceCreateReqVO createReqVO);
    String createProcessInstance(Long userId, BpmProcessInstanceCreateReqDTO createReqDTO);
    BpmProcessInstanceRespVO getProcessInstanceVO(String id);
    void cancelProcessInstance(Long userId, BpmProcessInstanceCancelReqVO cancelReqVO);
    void updateProcessInstanceExtComplete(ProcessInstance instance);
    void updateProcessInstanceExtReject(String id, String reason);
}
```

**API接口**: `BpmProcessInstanceApi`
```java
public interface BpmProcessInstanceApi {
    String createProcessInstance(Long userId, BpmProcessInstanceCreateReqDTO reqDTO);
}
```

### 3.2 流程模型服务

**Service接口**: `BpmModelService`
```java
public interface BpmModelService {
    PageResult<BpmModelPageItemRespVO> getModelPage(BpmModelPageReqVO pageVO);
    String createModel(BpmModelCreateReqVO modelVO, String bpmnXml);
    BpmModelRespVO getModel(String id);
    void updateModel(BpmModelUpdateReqVO updateReqVO);
    void deployModel(String id);  // 部署流程模型
    void deleteModel(String id);
    void updateModelState(String id, Integer state);
    BpmnModel getBpmnModel(String id);
}
```

### 3.3 任务服务

**Service接口**: `BpmTaskService`
- 待办任务查询
- 已办任务查询
- 任务审批/驳回
- 任务转派

---

## 4. 核心数据结构

### 4.1 流程模型 DO

```java
// Flowable原生BpmnModel
import org.flowable.bpmn.model.BpmnModel;
```

### 4.2 流程定义扩展 DO

```java
@TableName("bpm_process_definition_ext")
public class BpmProcessDefinitionExtDO {
    private String processDefinitionId;  // 流程定义ID
    private String description;           // 描述
    private String formType;              // 表单类型
    private String formId;                // 表单编号
    private String formCustomCreatePath;  // 自定义创建路径
    private String formCustomViewPath;    // 自定义查看路径
}
```

### 4.3 任务分配规则 DO

```java
@TableName("bpm_task_assign_rule")
public class BpmTaskAssignRuleDO {
    private String modelId;          // 流程模型ID
    private String taskDefineKey;    // 任务定义Key
    private Integer ruleType;        // 规则类型
    private String options;          // 选项(JSON)
    private String script;           // 脚本
}
```

### 4.4 枚举常量

```java
// 流程实例状态
public enum BpmProcessInstanceStatusEnum {
    RUNNING(1, "审批中"),
    COMPLETED(2, "已完成"),
    CANCELLED(3, "已取消");
}

// 流程实例结果
public enum BpmProcessInstanceResultEnum {
    APPROVE(1, "通过"),
    REJECT(2, "不通过"),
    CANCEL(3, "取消");
}

// 表单类型
public enum BpmModelFormTypeEnum {
    NORMAL(1, "正常表单"),
    CUSTOM(2, "自定义表单");
}

// 任务分配规则类型
public enum BpmTaskAssignRuleTypeEnum {
    USER(1, "指定用户"),
    ROLE(2, "指定角色"),
    SCRIPT(3, "脚本计算");
}
```

---

## 5. 数据流向

### 5.1 流程创建流程

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

### 5.2 流程审批流程

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

### 5.3 流程模型部署

```
用户上传BPMN XML或可视化设计
        ↓
BpmModelController.createModel()
        ↓
BpmModelService.createModel()
        ↓
BpmModelService.deployModel()  // 部署
        ↓
创建ProcessDefinition记录
```

---

## 6. 关键技术实现

### 6.1 技术栈

- **流程引擎**: Flowable 6.x
- **持久层**: MyBatis-Plus + MyBatis
- **数据库**: MySQL
- **安全**: Spring Security + 自定义权限表达式

### 6.2 Flowable核心服务注入

```java
@Resource
private RepositoryService repositoryService;    // 流程定义
@Resource
private RuntimeService runtimeService;         // 流程运行
@Resource
private TaskService taskService;               // 任务管理
@Resource
private HistoryService historyService;         // 历史记录
@Resource
private ManagementService managementService;   // 管理服务
```

### 6.3 事件机制

```java
// 流程实例结果事件
public class BpmProcessInstanceResultEvent extends ApplicationEvent {
    private String processInstanceId;
    private Integer result;  // 审批结果
}

// 事件监听器
@Component
public class BpmProcessInstanceResultEventListener {
    @EventListener
    public void onResult(BpmProcessInstanceResultEvent event) {
        // 更新流程实例扩展记录
    }
}
```

### 6.4 安全配置

```java
@Configuration
public class BpmSecurityConfiguration extends WebSecurityConfigurerAdapter {
    // BPM专用安全配置
}
```

---

## 7. API接口

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| 创建流程实例 | POST | /bpm/process-instance/create | 发起新流程 |
| 取消流程实例 | PUT | /bpm/process-instance/cancel | 取消流程 |
| 我的流程实例 | GET | /bpm/process-instance/my-page | 我的申请分页 |
| 获取流程详情 | GET | /bpm/process-instance/get | 获取流程VO |
| 创建流程模型 | POST | /bpm/model/create | 创建模型 |
| 部署流程模型 | PUT | /bpm/model/deploy/{id} | 部署模型 |
| 删除流程模型 | DELETE | /bpm/model/delete/{id} | 删除模型 |
| 获取流程模型 | GET | /bpm/model/get/{id} | 获取模型详情 |
| 分页查询模型 | GET | /bpm/model/page | 模型分页 |
| 获取待办任务 | GET | /bpm/task/todo-page | 待办任务分页 |
| 获取已办任务 | GET | /bpm/task/done-page | 已办任务分页 |
| 审批任务 | PUT | /bpm/task/approve | 审批任务 |
| 驳回任务 | PUT | /bpm/task/reject | 驳回任务 |

---

## 8. 数据库表

| 表名 | 说明 |
|------|------|
| bpm_form | 流程表单配置 |
| bpm_process_definition_ext | 流程定义扩展 |
| bpm_task_assign_rule | 任务分配规则 |
| bpm_task_message_rule | 任务消息规则 |
| bpm_user_group | 用户组 |
| bpm_process_instance_ext | 流程实例扩展 |
| bpm_task_ext | 任务扩展 |
| bpm_oa_leave | OA请假记录 |

---

## 9. 与其他模块的交互

- **System模块**: 用户信息、权限验证
- **业务模块**: 通过`BpmProcessInstanceApi`被其他模块调用创建流程
- **消息模块**: 流程审批通知推送

---

## 10. OA集成示例 - 请假流程

**Controller**: `BpmOALeaveController`
- `/bpm/oa/leave/create` - 创建请假
- `/bpm/oa/leave/page` - 请假记录分页

**DO**: `BpmOALeaveDO`
```java
@TableName("bpm_oa_leave")
public class BpmOALeaveDO extends BaseDO {
    private Long userId;           // 用户ID
    private String leaveType;      // 请假类型
    private BigDecimal duration;   // 时长
    private LocalDateTime startTime;
    private LocalDateTime endTime;
    private String reason;         // 理由
    private String instanceId;      // 流程实例ID
    private Integer result;        // 审批结果
}
```

**监听器**: `BpmOALeaveResultListener`
- 实现流程审批结果回调
- 更新请假记录的审批结果
