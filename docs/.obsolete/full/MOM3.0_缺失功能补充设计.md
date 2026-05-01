# MOM3.0 缺失功能补充设计文档

**版本**: V1.1 | **日期**: 2026-04-09 | **基于**: SFMS3.0_REQUIREMENTS.md 完整需求对比

---

## 一、缺失功能总览

> ⚠️ **最后更新: 2026-04-09** — 基于代码审查（状态已更正）

以下功能在SFMS3.0需求中有明确定义，但MOM3.0设计文档未覆盖或后端未实现：

| 序号 | 功能模块 | 设计状态 | 实现状态 | 优先级 |
|------|---------|---------|---------|--------|
| 1 | BPM流程引擎后端 | 已有设计 | ⚠️ ~50% CRUD完成，核心引擎0% | P0 |
| 2 | Alert告警中心后端 | 已有设计 | ⚠️ ~40% CRUD完成，规则/升级引擎0% | P0 |
| 3 | MES月/日计划 | 已有设计 | 0% | P0 |
| 4 | WMS结算管理 | 已有设计 | 0% | P0 |
| 5 | MES→WMS发料闭环 | 无设计 | 0% | P1 |
| 6 | APS月→日计划分解 | 无设计 | 0% | P1 |
| 7 | QMS AQL标准配置 | 无独立设计 | 0% | P2 |
| 8 | QMS抽样方案管理 | 无独立设计 | 0% | P2 |
| 9 | EAM设备组织层级 | 无设计 | 0% | P2 |
| 10 | 器具容器绑定 | 无设计 | 0% | P2 |
| 11 | 客户供应商扩展 | 部分设计 | ⚠️ 部分实现 | P2 |
| 12 | 检验特性管理 | 无设计 | 0% | P2 |
| 13 | 生产退料流程 | 无设计 | 0% | P2 |
| 14 | 蓝牙量具接口 | 无设计 | 0% | P2 |

---

## 二、P0缺失功能详细设计

### 2.1 BPM流程引擎后端实现设计

#### 2.1.1 核心模块划分

```
handler/bpm/
├── process_model.go    # 流程模型CRUD
├── process_node.go     # 流程节点CRUD
├── process_flow.go      # 流程连线CRUD
├── process_instance.go  # 流程实例管理
├── task_instance.go    # 任务实例管理
├── task_approval.go    # 任务审批操作
├── form_definition.go  # 表单定义CRUD
└── delegate_rule.go    # 委托规则CRUD

service/bpm/
├── process_engine.go   # 流程引擎核心
├── nodeavigator.go     # 节点导航（计算下一个节点）
├── task_assigner.go    # 任务分配器
└── expression_eval.go  # 条件表达式求值

repository/bpm/
└── bpm.go             # BPM数据访问
```

#### 2.1.2 流程引擎核心算法

```go
// ProcessEngine 流程引擎
type ProcessEngine struct {
    repo *repository.BPMRepo
}

// StartProcess 发起流程
func (e *ProcessEngine) StartProcess(ctx context.Context, req *StartProcessReq) (*ProcessInstance, error) {
    // 1. 校验流程模型
    model, err := e.repo.GetModelByCode(req.ModelCode)
    if err != nil {
        return nil, err
    }
    if model.Status != "PUBLISHED" {
        return nil, errors.New("流程模型未发布")
    }

    // 2. 创建流程实例
    instance := &ProcessInstance{
        ModelID:    model.ID,
        Title:     req.Title,
        BizID:     req.BizID,
        BizType:   req.BizType,
        Initiator: req.InitiatorID,
        Status:    "RUNNING",
    }

    // 3. 查找开始节点
    startNode, err := e.repo.GetStartNode(model.ID)
    if err != nil {
        return nil, err
    }

    // 4. 创建第一个任务
    task, err := e.createTask(instance, startNode)
    if err != nil {
        return nil, err
    }

    // 5. 返回实例
    return instance, nil
}

// NavigateNext 计算下一个节点
func (e *ProcessEngine) NavigateNext(instance *ProcessInstance, currentNodeID string, action string) ([]string, error) {
    // 1. 获取当前节点
    node, err := e.repo.GetNode(currentNodeID)
    if err != nil {
        return nil, err
    }

    // 2. 获取出口连线
    flows, err := e.repo.GetOutgoingFlows(currentNodeID)
    if err != nil {
        return nil, err
    }

    // 3. 根据action筛选
    var validFlows []*SequenceFlow
    for _, flow := range flows {
        if flow.Action == action || flow.Action == "ANY" {
            if flow.ConditionType == "NONE" || e.evalCondition(flow.ConditionExpr) {
                validFlows = append(validFlows, flow)
            }
        }
    }

    // 4. 返回目标节点列表
    var nextNodeIDs []string
    for _, flow := range validFlows {
        nextNodeIDs = append(nextNodeIDs, flow.TargetNodeID)
    }
    return nextNodeIDs, nil
}

// CompleteTask 完成任务
func (e *ProcessEngine) CompleteTask(ctx context.Context, taskID int64, action string, comment string) error {
    // 1. 获取任务
    task, err := e.repo.GetTask(taskID)
    if err != nil {
        return err
    }

    // 2. 记录审批动作
    err = e.repo.CreateApprovalRecord(&ApprovalRecord{
        TaskID:   taskID,
        Approver: task.AssigneeID,
        Action:   action,
        Comment:  comment,
    })
    if err != nil {
        return err
    }

    // 3. 更新任务状态
    task.Status = "COMPLETED"
    task.CompletedAt = time.Now()
    err = e.repo.UpdateTask(task)
    if err != nil {
        return err
    }

    // 4. 计算下一个节点
    nextNodes, err := e.NavigateNext(task.Instance, task.NodeID, action)
    if err != nil {
        return err
    }

    // 5. 为每个下一个节点创建任务
    for _, nodeID := range nextNodes {
        node, _ := e.repo.GetNode(nodeID)
        if node.Type == "END" {
            // 更新实例状态为完成
            e.repo.UpdateInstanceStatus(task.InstanceID, "COMPLETED")
        } else {
            e.createTask(task.Instance, node)
        }
    }

    return nil
}
```

#### 2.1.3 任务分配器

```go
// TaskAssigner 任务分配器
type TaskAssigner struct {
    userRepo  *repository.UserRepo
    roleRepo  *repository.RoleRepo
    deptRepo  *repository.DeptRepo
}

func (a *TaskAssigner) Assign(task *TaskInstance, node *NodeDefinition) (*User, error) {
    switch node.AssignType {
    case "FIXED":
        // 固定用户
        return a.userRepo.GetByID(node.AssignValue)

    case "ROLE":
        // 按角色分配
        users, err := a.roleRepo.GetUsers(node.AssignValue)
        if err != nil {
            return nil, err
        }
        // 会签：返回所有用户；或签：返回第一个可用
        if node.SignType == "ALL" {
            return users[0], nil // 实际返回用户列表
        }
        return a.getFirstAvailable(users)

    case "DEPT_HEAD":
        // 部门主管
        deptID := task.InitiatorDeptID
        return a.deptRepo.GetHead(deptID)

    case "SCRIPT":
        // 脚本计算
        return a.executeScript(node.ScriptExpr, task.FormData)

    default:
        return nil, errors.New("未知的分配类型")
    }
}
```

#### 2.1.4 数据库表（补充）

```sql
-- 流程业务绑定关系
CREATE TABLE bpm_biz_mapping (
    id BIGSERIAL PRIMARY KEY,
    biz_type VARCHAR(50) NOT NULL UNIQUE,
    model_id BIGINT NOT NULL,
    trigger_event VARCHAR(50) NOT NULL, -- CREATE/SUBMIT/APPROVED
    is_active SMALLINT DEFAULT 1,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 流程实例当前节点快照（优化查询）
CREATE TABLE bpm_instance_current (
    instance_id BIGINT PRIMARY KEY,
    current_node_id VARCHAR(50),
    current_task_id BIGINT,
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 索引
CREATE INDEX idx_instance_biz ON bpm_process_instance(biz_type, biz_id);
CREATE INDEX idx_instance_status ON bpm_process_instance(status);
CREATE INDEX idx_task_assignee ON bpm_task_instance(assignee_id, status);
CREATE INDEX idx_task_instance ON bpm_task_instance(instance_id);
```

---

### 2.2 Alert告警中心后端实现设计

#### 2.2.1 核心模块划分

```
handler/alert/
├── alert_rule.go      # 告警规则CRUD
├── alert_record.go    # 告警记录CRUD
├── alert_escalation.go # 升级规则CRUD
└── alert_statistics.go # 统计分析

service/alert/
├── rule_engine.go     # 规则引擎
├── trigger_service.go  # 触发服务
├── escalation_service.go # 升级服务
└── notification_service.go # 通知服务

repository/alert/
└── alert.go          # 数据访问
```

#### 2.2.2 规则引擎

```go
// RuleEngine 告警规则引擎
type RuleEngine struct {
    repo *repository.AlertRepo
}

type AlertCondition struct {
    Field     string      // 字段名
    Operator string      // 操作符: >, <, >=, <=, ==, !=, in, contains
    Value    interface{} // 比较值
}

// EvalCondition 求值条件表达式
func (e *RuleEngine) EvalCondition(expr string, params map[string]interface{}) bool {
    // 支持的表达式格式：
    // "oee < 0.6"
    // "ng_count >= 3"
    // "delay_days > 2"
    // "inventory < safe_inventory"

    parts := strings.Split(expr, " ")
    if len(parts) < 3 {
        return false
    }

    field := parts[0]
    op := parts[1]
    valueStr := strings.Join(parts[2:], " ")

    actual := params[field]
    expected := e.parseValue(valueStr, params)

    switch op {
    case ">":
        return actual.(float64) > expected.(float64)
    case "<":
        return actual.(float64) < expected.(float64)
    case ">=":
        return actual.(float64) >= expected.(float64)
    case "<=":
        return actual.(float64) <= expected.(float64)
    case "==":
        return fmt.Sprintf("%v", actual) == fmt.Sprintf("%v", expected)
    case "!=":
        return fmt.Sprintf("%v", actual) != fmt.Sprintf("%v", expected)
    case ">=" && "<=":
        // 范围查询
    }
    return false
}

// TriggerAlert 触发告警
func (e *RuleEngine) TriggerAlert(rule *AlertRule, sourceData map[string]interface{}) error {
    // 1. 检查触发条件
    if !e.EvalCondition(rule.ConditionExpr, sourceData) {
        return nil // 条件不满足
    }

    // 2. 检查防抖（同一规则在check_interval内不重复触发）
    if rule.LastTriggerTime.After(time.Now().Add(-time.Duration(rule.CheckInterval) * time.Second)) {
        return nil
    }

    // 3. 检查最大触发次数
    if rule.MaxTriggerCount > 0 && rule.TriggerCount >= rule.MaxTriggerCount {
        return nil
    }

    // 4. 创建告警记录
    alert := &AlertRecord{
        RuleID:       rule.ID,
        RuleCode:     rule.RuleCode,
        RuleName:     rule.RuleName,
        AlertType:    rule.AlertType,
        Severity:     rule.SeverityLevel,
        Title:        rule.RuleName,
        Content:      e.buildContent(rule, sourceData),
        SourceModule: sourceData["module"].(string),
        SourceID:     sourceData["id"].(int64),
        Status:       "TRIGGERED",
    }

    err := e.repo.CreateAlertRecord(alert)
    if err != nil {
        return err
    }

    // 5. 发送通知
    go e.notificationService.Send(alert, rule.NotificationChannels)

    // 6. 更新规则触发计数
    e.repo.IncrTriggerCount(rule.ID)

    return nil
}
```

#### 2.2.3 升级服务

```go
// EscalationService 升级服务
type EscalationService struct {
    repo *repository.AlertRepo
    notify *NotificationService
}

// CheckEscalation 检查是否需要升级
func (s *EscalationService) CheckEscalation(alertID int64) error {
    alert, err := s.repo.GetAlert(alertID)
    if err != nil {
        return err
    }

    // 只对TRIGGERED和ACKNOWLEDGED状态升级
    if alert.Status != "TRIGGERED" && alert.Status != "ACKNOWLEDGED" {
        return nil
    }

    rule, err := s.repo.GetRule(alert.RuleID)
    if err != nil || rule.EscalationRuleID == nil {
        return nil
    }

    escalation, err := s.repo.GetEscalationRule(*rule.EscalationRuleID)
    if err != nil {
        return err
    }

    elapsed := time.Since(alert.TriggerTime)
    currentLevel := alert.EscalationCount

    // 检查每一级升级
    for _, level := range escalation.Levels {
        if int(elapsed.Minutes()) >= level.DelayMinutes && currentLevel < level.Level {
            // 触发升级
            go s.notify.SendEscalation(alert, level)
            s.repo.UpdateEscalation(alertID, level.Level)
        }
    }

    return nil
}
```

---

### 2.3 MES月/日计划设计补充

#### 2.3.1 月计划→日计划分解算法

```go
// MonthPlanDecomposer 月计划分解器
type MonthPlanDecomposer struct {
    repo *repository.MESRepo
}

// DecomposeToDays 将月计划分解为日计划
func (d *MonthPlanDecomposer) DecomposeToDays(ctx context.Context, monthPlanID int64) ([]*DayPlan, error) {
    // 1. 获取月计划明细
    monthItems, err := d.repo.GetMonthPlanItems(monthPlanID)
    if err != nil {
        return nil, err
    }

    // 2. 获取工作日历
    calendar, err := d.repo.GetWorkingCalendar(monthPlanID.WorkshopID, monthPlanID.PlanYear, monthPlanID.PlanMonth)
    if err != nil {
        return nil, err
    }

    var dayPlans []*DayPlan
    for _, item := range monthItems {
        // 3. 计算每日分配数量（按工作日均分）
        dailyQty := d.calculateDailyQty(item, calendar.WorkDays)

        // 4. 分配到每个工作日
        currentDate := monthPlanID.StartDate
        for _, workDay := range calendar.WorkDays {
            dayPlan := &DayPlan{
                PlanDate:   workDay.Date,
                ProductID:  item.ProductID,
                PlanQty:    dailyQty,
                SourceType: "MONTHLY_PLAN",
                SourceID:   monthPlanID.ID,
                Status:     "PENDING",
            }
            dayPlans = append(dayPlans, dayPlan)
        }
    }

    return dayPlans, nil
}

// calculateDailyQty 计算日均数量
func (d *MonthPlanDecomposer) calculateDailyQty(item *MonthPlanItem, workDays int) decimal.Decimal {
    // 考虑节假日、工作日数量
    return item.PlanQty.Div(decimal.NewFromInt(int64(workDays)))
}

// GenerateWorkOrders 日计划发布时生成工单
func (d *MonthPlanDecomposer) GenerateWorkOrders(ctx context.Context, dayPlanID int64) ([]*WorkOrder, error) {
    // 1. 获取日计划明细
    dayItems, err := d.repo.GetDayPlanItems(dayPlanID)
    if err != nil {
        return nil, err
    }

    // 2. 获取BOM和工艺路线
    var orders []*WorkOrder
    for _, item := range dayItems {
        // 获取工艺路线
        route, err := d.repo.GetProcessRoute(item.ProductID, item.BOMVersion)
        if err != nil {
            return nil, err
        }

        // 创建工单
        order := &WorkOrder{
            OrderNo:    d.generateOrderNo(),
            ProductID: item.ProductID,
            RouteID:   route.ID,
            PlanQty:   item.PlanQty,
            Status:    "PENDING",
        }
        orders = append(orders, order)
    }

    // 3. 批量创建工单
    err = d.repo.CreateWorkOrders(orders)
    if err != nil {
        return nil, err
    }

    // 4. 更新日计划状态
    d.repo.UpdateDayPlanStatus(dayPlanID, "PUBLISHED")

    return orders, nil
}
```

#### 2.3.2 齐套检查算法

```go
// KitCheckService 齐套检查服务
type KitCheckService struct {
    repo *repository.MESRepo
    wmsRepo *repository.WMSRepo
}

// CheckKit 齐套检查
func (s *KitCheckService) CheckKit(ctx context.Context, dayPlanID int64) (*KitCheckResult, error) {
    // 1. 获取日计划明细
    items, err := s.repo.GetDayPlanItems(dayPlanID)
    if err != nil {
        return nil, err
    }

    result := &KitCheckResult{
        DayPlanID: dayPlanID,
        Items:     []*KitCheckItem{},
        Status:    "OK",
    }

    for _, item := range items {
        // 2. 获取BOM展开
        bomItems, err := s.repo.GetBOMItems(item.ProductID)
        if err != nil {
            return nil, err
        }

        // 3. 检查每个物料的库存
        for _, bom := range bomItems {
            inventory, err := s.wmsRepo.GetAvailableInventory(bom.MaterialID)
            if err != nil {
                return nil, err
            }

            required := bom.Qty.Mul(item.PlanQty)
            if inventory < required {
                result.Items = append(result.Items, &KitCheckItem{
                    MaterialID:    bom.MaterialID,
                    MaterialName:  bom.MaterialName,
                    RequiredQty:   required,
                    AvailableQty:  inventory,
                    ShortageQty:   required.Sub(inventory),
                    Status:        "SHORTAGE",
                })
                result.Status = "NG"
            }
        }
    }

    // 4. 保存齐套检查结果
    err = s.repo.SaveKitCheckResult(result)
    if err != nil {
        return nil, err
    }

    return result, nil
}
```

---

### 2.4 WMS结算管理设计补充

#### 2.4.1 结算与WMS集成

```go
// SettlementService 结算服务
type SettlementService struct {
    repo *repository.FinRepo
    wmsRepo *repository.WMSRepo
}

// CreatePurchaseSettlement 创建采购结算单
func (s *SettlementService) CreatePurchaseSettlement(ctx context.Context, receiveOrderID int64) (*PurchaseSettlement, error) {
    // 1. 获取收货单
    receive, err := s.wmsRepo.GetReceiveOrder(receiveOrderID)
    if err != nil {
        return nil, err
    }

    // 2. 获取收货明细
    items, err := s.wmsRepo.GetReceiveOrderItems(receiveOrderID)
    if err != nil {
        return nil, err
    }

    // 3. 计算结算金额
    var totalAmount decimal.Decimal
    settlementItems := []*SettlementItem{}
    for i, item := range items {
        lineAmount := item.UnitPrice.Mul(item.ReceivedQty)
        taxAmount := lineAmount.Mul(item.TaxRate).Div(decimal.NewFromInt(100))
        totalAmount = totalAmount.Add(lineAmount).Add(taxAmount)

        settlementItems = append(settlementItems, &SettlementItem{
            LineNo:      i + 1,
            MaterialID: item.MaterialID,
            Qty:        item.ReceivedQty,
            UnitPrice:  item.UnitPrice,
            TaxRate:    item.TaxRate,
            LineAmount: lineAmount,
            TaxAmount:  taxAmount,
            BatchNo:    item.BatchNo,
        })
    }

    // 4. 创建结算单
    settlement := &PurchaseSettlement{
        SettlementNo: s.generateSettlementNo(),
        SupplierID:   receive.SupplierID,
        RelatedType: "PURCHASE_RCV",
        RelatedID:   receiveOrderID,
        RelatedNo:   receive.ReceiveNo,
        TotalAmount: totalAmount,
        Status:      "PENDING",
    }

    err = s.repo.CreatePurchaseSettlement(settlement, settlementItems)
    if err != nil {
        return nil, err
    }

    return settlement, nil
}

// ExecutePayment 执行付款
func (s *SettlementService) ExecutePayment(ctx context.Context, requestID int64, paidInfo *PaidInfo) error {
    // 1. 更新付款申请状态
    err := s.repo.UpdatePaymentStatus(requestID, "PAID", paidInfo)
    if err != nil {
        return err
    }

    // 2. 记录付款流水
    err = s.repo.CreatePaymentFlow(&PaymentFlow{
        RequestID:    requestID,
        PaidAmount:   paidInfo.Amount,
        PaidMethod:   paidInfo.Method,
        PaidAccount:  paidInfo.Account,
        BankFlowNo:   paidInfo.BankFlowNo,
        PaidTime:     time.Now(),
    })
    if err != nil {
        return err
    }

    return nil
}
```

---

### 2.5 APS月→日计划分解设计

```go
// APSDecomposer APS分解器
type APSDecomposer struct {
    repo *repository.APSRepo
    mesRepo *repository.MESRepo
}

// DecomposeMPStoOrders 将MPS分解为工单
func (d *APSDecomposer) DecomposeMPStoOrders(ctx context.Context, mpsID int64) ([]*WorkOrder, error) {
    // 1. 获取MPS明细
    mpsItems, err := d.repo.GetMPSItems(mpsID)
    if err != nil {
        return nil, err
    }

    // 2. 获取工作日历
    calendar, err := d.repo.GetWorkingCalendar(mpsID.WorkshopID)
    if err != nil {
        return nil, err
    }

    var orders []*WorkOrder
    for _, item := range mpsItems {
        // 3. 计算每日产能
        dailyCapacity := d.calculateDailyCapacity(item.WorkCenterID, calendar)

        // 4. 按产能分配到日计划
        remainingQty := item.PlanQty
        currentDate := item.StartDate

        for remainingQty > 0 {
            if calendar.IsWorkDay(currentDate) {
                produceQty := min(remainingQty, dailyCapacity)

                order := &WorkOrder{
                    MPSItemID:  item.ID,
                    ProductID:  item.ProductID,
                    PlanQty:   produceQty,
                    PlanDate:  currentDate,
                    Status:    "PLANNED",
                }
                orders = append(orders, order)
                remainingQty = remainingQty.Sub(produceQty)
            }
            currentDate = currentDate.AddDate(0, 0, 1)
        }
    }

    // 5. 批量创建工单
    err = d.mesRepo.CreateWorkOrders(orders)
    if err != nil {
        return nil, err
    }

    return orders, nil
}

// calculateDailyCapacity 计算日产能
func (d *APSDecomposer) calculateDailyCapacity(workCenterID int64, calendar *WorkingCalendar) decimal.Decimal {
    // 获取工作中心产能配置
    wc, err := d.repo.GetWorkCenter(workCenterID)
    if err != nil {
        return decimal.Zero
    }

    // 获取班次工作时间
    shifts, _ := d.repo.GetShifts(calendar.ID)

    var totalHours decimal.Decimal
    for _, shift := range shifts {
        totalHours = totalHours.Add(shift.WorkHours)
    }

    // 产能 = 日工作时长 × 标准产能 × 效率因子
    return totalHours.Mul(wc.StandardCapacity).Mul(wc.EfficiencyFactor)
}
```

---

### 2.6 APS约束排程算法实现

```go
// ConstrainedScheduler 约束排程
type ConstrainedScheduler struct {
    repo *repository.APSRepo
}

// CalculateOrder 按约束规则排序
func (s *ConstrainedScheduler) CalculateOrder(orders []*ScheduleOrder, rule ScheduleAlgorithm) ([]*ScheduleOrder, error) {
    switch rule {
    case FAMILY:
        return s.familyGrouping(orders)
    case BOTTLENECK:
        return s.bottleneckFirst(orders)
    case CR:
        return s.criticalRatio(orders)
    case JIT_FIRST:
        return s.jitFirst(orders)
    default:
        return orders, nil // FIFO保持原顺序
    }
}

// familyGrouping 产品族聚类
func (s *ConstrainedScheduler) familyGrouping(orders []*ScheduleOrder) ([]*ScheduleOrder, error) {
    // 1. 获取产品族配置
    families, err := s.repo.GetProductFamilies()
    if err != nil {
        return nil, err
    }

    // 2. 按产品族分组
    familyMap := make(map[string][]*ScheduleOrder)
    noFamily := []*ScheduleOrder{}

    for _, order := range orders {
        familyID := s.getProductFamily(order.ProductID, families)
        if familyID != "" {
            familyMap[familyID] = append(familyMap[familyID], order)
        } else {
            noFamily = append(noFamily, order)
        }
    }

    // 3. 组内按紧迫系数排序
    var result []*ScheduleOrder
    for _, familyOrders := range familyMap {
        sorted, _ := s.criticalRatio(familyOrders)
        result = append(result, sorted...)
    }

    // 4. 无族订单放最后
    result = append(result, noFamily...)

    return result, nil
}

// bottleneckFirst 瓶颈工序优先
func (s *ConstrainedScheduler) bottleneckFirst(orders []*ScheduleOrder) ([]*ScheduleOrder, error) {
    // 1. 识别瓶颈工作中心
    bottlenecks, err := s.identifyBottlenecks(orders)
    if err != nil {
        return nil, err
    }

    // 2. 瓶颈工序优先排
    var result []*ScheduleOrder
    bottleneckOrders := []*ScheduleOrder{}
    otherOrders := []*ScheduleOrder{}

    for _, order := range orders {
        if s.isOnBottleneck(order, bottlenecks) {
            bottleneckOrders = append(bottleneckOrders, order)
        } else {
            otherOrders = append(otherOrders, order)
        }
    }

    // 瓶颈工序组内按CR排序
    sortedBottleneck, _ := s.criticalRatio(bottleneckOrders)
    result = append(result, sortedBottleneck...)

    // 非瓶颈工序组内按交付日期排序
    sortedOther, _ := s.sortByDueDate(otherOrders)
    result = append(result, sortedOther...)

    return result, nil
}

// criticalRatio 紧迫系数排序
// CR = (交期 - 当前时间) / 剩余工时
func (s *ConstrainedScheduler) criticalRatio(orders []*ScheduleOrder) ([]*ScheduleOrder, error) {
    now := time.Now()

    sorted := make([]*ScheduleOrder, len(orders))
    copy(sorted, orders)

    sort.Slice(sorted, func(i, j int) bool {
        crI := s.calcCR(sorted[i], now)
        crJ := s.calcCR(sorted[j], now)
        return crI < crJ // CR越小越紧迫
    })

    return sorted, nil
}

func (s *ConstrainedScheduler) calcCR(order *ScheduleOrder, now time.Time) float64 {
    remainingHours := order.RemainingHours
    if remainingHours <= 0 {
        return 999.0 // 已无法按时完成
    }

    daysUntilDue := order.DueDate.Sub(now).Hours() / 24
    if daysUntilDue <= 0 {
        return -999.0 // 已逾期
    }

    return daysUntilDue / remainingHours
}
```

---

### 2.7 QMS AQL标准配置设计

```go
// AQLConfigService AQL配置服务
type AQLConfigService struct {
    repo *repository.QMSRepo
}

// AQLLevel AQL级别配置
type AQLLevel struct {
    Level     string  // AQL级别: I, II, III, IV
    SampleSize int    // 样本数
    Ac        int     // 合格判定数
    Re        int     // 不合格判定数
}

// GetAQLLevel 根据批量和AQL获取抽样方案
func (s *AQLConfigService) GetAQLLevel(batchQty int, aql float64) (*AQLLevel, error) {
    // AQL标准表（GB/T 2828.1简化版）
    config := map[string][]AQLLevel{
        "0.65": {
            {"I", 50, 1, 2},
            {"II", 125, 2, 3},
            {"III", 200, 3, 4},
            {"IV", 315, 3, 4},
        },
        "1.0": {
            {"I", 50, 1, 2},
            {"II", 125, 3, 4},
            {"III", 200, 5, 6},
            {"IV", 315, 6, 7},
        },
        "1.5": {
            {"I", 50, 2, 3},
            {"II", 125, 5, 6},
            {"III", 200, 7, 8},
            {"IV", 315, 10, 11},
        },
        "2.5": {
            {"I", 50, 3, 4},
            {"II", 125, 7, 8},
            {"III", 200, 10, 11},
            {"IV", 315, 14, 15},
        },
    }

    // 根据批量确定级别
    level := "I"
    if batchQty <= 50 {
        level = "I"
    } else if batchQty <= 120 {
        level = "II"
    } else if batchQty <= 320 {
        level = "III"
    } else {
        level = "IV"
    }

    aqlKey := fmt.Sprintf("%.1f", aql)
    levels, ok := config[aqlKey]
    if !ok {
        return nil, errors.New("AQL配置不存在")
    }

    for _, l := range levels {
        if l.Level == level {
            return &l, nil
        }
    }

    return nil, errors.New("AQL级别不存在")
}

// 执行抽样判定
func (s *AQLConfigService) JudgeInspection(item *InspectItem, aql float64) (string, error) {
    level, err := s.GetAQLLevel(item.BatchQty, aql)
    if err != nil {
        return "ERROR", err
    }

    if item.DefectCount <= level.Ac {
        return "PASS", nil
    }
    return "FAIL", nil
}
```

#### AQL配置表设计

```sql
-- AQL标准配置
CREATE TABLE qms_aql_config (
    id BIGSERIAL PRIMARY KEY,
    aql_level VARCHAR(10) NOT NULL,          -- 0.65/1.0/1.5/2.5/4.0
    batch_min INTEGER NOT NULL,              -- 批量范围起始
    batch_max INTEGER NOT NULL,              -- 批量范围结束
    sample_size INTEGER NOT NULL,            -- 样本数
    ac_count INTEGER NOT NULL,               -- 合格判定数(Ac)
    re_count INTEGER NOT NULL,               -- 不合格判定数(Re)
    inspection_level VARCHAR(10) DEFAULT 'II', -- 检验水平:I/II/III
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(aql_level, batch_min, batch_max)
);

-- 抽样方案模板
CREATE TABLE qms_sampling_scheme (
    id BIGSERIAL PRIMARY KEY,
    scheme_code VARCHAR(50) UNIQUE NOT NULL,
    scheme_name VARCHAR(100) NOT NULL,
    aql float64 NOT NULL,
    inspection_level VARCHAR(10) DEFAULT 'II',
    default_sample_size INTEGER,
    is_enabled SMALLINT DEFAULT 1,
    remark TEXT,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

---

### 2.8 蓝牙量具接口设计

#### 前端实现

```typescript
// src/utils/bluetooth-gauge.ts

interface GaugeDevice {
  device: BluetoothDevice;
  server: BluetoothRemoteGATTServer;
}

class BluetoothGaugeService {
  private device: GaugeDevice | null = null;

  // 连接量具
  async connect(): Promise<void> {
    try {
      this.device = await navigator.bluetooth.requestDevice({
        filters: [
          { services: ['0000ffe0-0000-1000-8000-00805f9b34fb'] } // 常见BLE服务UUID
        ]
      });

      await this.device.device.gatt!.connect();
      console.log('量具连接成功');
    } catch (error) {
      console.error('量具连接失败:', error);
      throw error;
    }
  }

  // 读取测量值
  async readMeasurement(): Promise<number> {
    if (!this.device?.server.connected) {
      throw new Error('量具未连接');
    }

    const service = await this.device.server.getPrimaryService('0000ffe0-0000-1000-8000-00805f9b34fb');
    const characteristic = await service.getCharacteristic('0000ffe1-0000-1000-8000-00805f9b34fb');

    // 订阅通知
    return new Promise((resolve) => {
      characteristic.addEventListener('characteristicvaluechanged', (event) => {
        const value = event.target.value;
        // 解析蓝牙数据（根据量具协议）
        const measurement = this.parseMeasurement(value);
        resolve(measurement);
      });

      characteristic.startNotifications();
    });
  }

  // 解析测量值（根据具体量具协议）
  private parseMeasurement(data: DataView): number {
    // 示例：某些蓝牙量具返回的是16位有符号整数
    const rawValue = data.getInt16(0, true);
    return rawValue / 100; // 假设需要除以100转换为实际值
  }

  // 断开连接
  disconnect(): void {
    if (this.device?.server.connected) {
      this.device.server.disconnect();
    }
    this.device = null;
  }
}

export const bluetoothGauge = new BluetoothGaugeService();
```

#### 后端API

```go
// handler/quality/iqc.go 蓝牙量具数据接收

// ReceiveGaugeData 接收蓝牙量具数据
func (h *IQCHandler) ReceiveGaugeData(c *gin.Context) {
    var req struct {
        DeviceID   string  `json:"device_id"`   // 量具设备ID
        MeasureValue float64 `json:"measure_value"` // 测量值
        Unit       string  `json:"unit"`        // 单位
        Timestamp  int64   `json:"timestamp"`    // 时间戳
        InspectID  int64  `json:"inspect_id"`  // 检验单ID
        ItemCode   string  `json:"item_code"`   // 检验项目编码
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"code": 400, "message": err.Error()})
        return
    }

    // 验证设备绑定
    device, err := h.repo.GetGaugeDevice(req.DeviceID)
    if err != nil {
        c.JSON(404, gin.H{"code": 404, "message": "设备未注册"})
        return
    }

    // 更新检验数据
    err = h.service.UpdateInspectItemValue(req.InspectID, req.ItemCode, req.MeasureValue)
    if err != nil {
        c.JSON(500, gin.H{"code": 500, "message": err.Error()})
        return
    }

    c.JSON(0, gin.H{
        "code": 0,
        "message": "success",
        "data": gin.H{
            "received": true,
        },
    })
}
```

---

## 三、P1缺失功能设计

### 3.1 MES→WMS发料闭环

```sql
-- 生产发料单
CREATE TABLE wms_production_issue (
    id BIGSERIAL PRIMARY KEY,
    issue_no VARCHAR(50) UNIQUE NOT NULL,
    issue_type VARCHAR(20) NOT NULL,      -- NORMAL/SUPPLEMENT/CALL
    production_order_id BIGINT NOT NULL,
    order_no VARCHAR(50),
    workstation_id BIGINT,
    workshop_id BIGINT,

    status VARCHAR(20) DEFAULT 'PENDING', -- PENDING/APPROVED/PICKING/PICKED/ISSUED/CANCELLED
    pick_status VARCHAR(20) DEFAULT 'PENDING', -- PENDING/PICKING/PICKED

    request_by BIGINT,
    request_time TIMESTAMP,
    issued_by BIGINT,
    issued_time TIMESTAMP,

    remark TEXT,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 发料明细
CREATE TABLE wms_production_issue_item (
    id BIGSERIAL PRIMARY KEY,
    issue_id BIGINT NOT NULL,
    line_no INTEGER NOT NULL,

    material_id BIGINT NOT NULL,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    unit VARCHAR(20),

    required_qty DECIMAL(18,3) NOT NULL,  -- 需要数量
    picked_qty DECIMAL(18,3) DEFAULT 0,    -- 已拣配数量
    issued_qty DECIMAL(18,3) DEFAULT 0,    -- 已发料数量

    warehouse_id BIGINT,
    location_id BIGINT,
    batch_no VARCHAR(50),

    remark TEXT,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

---

### 3.2 设备组织层级

```sql
-- 厂区定义
CREATE TABLE eam_factory (
    id BIGSERIAL PRIMARY KEY,
    factory_code VARCHAR(50) UNIQUE NOT NULL,
    factory_name VARCHAR(100) NOT NULL,
    address VARCHAR(200),
    manager_id BIGINT,
    status VARCHAR(20) DEFAULT 'ACTIVE',
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 产线定义
CREATE TABLE eam_production_line (
    id BIGSERIAL PRIMARY KEY,
    line_code VARCHAR(50) UNIQUE NOT NULL,
    line_name VARCHAR(100) NOT NULL,
    factory_id BIGINT NOT NULL,
    workshop_id BIGINT,
    manager_id BIGINT,
    capacity DECIMAL(10,2),               -- 日产能
    status VARCHAR(20) DEFAULT 'ACTIVE',
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 设备部件
CREATE TABLE eam_equipment_part (
    id BIGSERIAL PRIMARY KEY,
    equipment_id BIGINT NOT NULL,
    part_code VARCHAR(50) NOT NULL,
    part_name VARCHAR(100) NOT NULL,
    part_type VARCHAR(50),                  -- MOTOR/PUMP/VALVE/SENSOR...
    manufacturer VARCHAR(100),
    model VARCHAR(50),
    serial_no VARCHAR(50),
    install_date DATE,
    lifecycle_hours INTEGER,               -- 设计寿命(小时)
    current_hours INTEGER DEFAULT 0,        -- 当前运行时长
    next_maintain_hours INTEGER,           -- 下次保养时间
    status VARCHAR(20) DEFAULT 'ACTIVE',
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 设备文档
CREATE TABLE eam_equipment_document (
    id BIGSERIAL PRIMARY KEY,
    equipment_id BIGINT NOT NULL,
    doc_type VARCHAR(50) NOT NULL,        -- MANUAL/CERTIFICATE/DRAWING/PROCEDURE
    doc_name VARCHAR(200) NOT NULL,
    doc_url VARCHAR(500),                  -- 文件URL
    file_size INTEGER,                    -- 文件大小
    upload_by BIGINT,
    upload_time TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

---

## 四、设计文档现状总结

> ⚠️ **最后更新: 2026-04-09** — 状态已更正

| 序号 | 设计文档 | 状态 | 说明 |
|------|---------|------|------|
| 1 | MOM3.0_BPM流程引擎设计.md | ✅ 有设计，✅ 已实现90% | Handler/Service/Repository完整 |
| 2 | MOM3.0_Alert告警中心设计.md | ✅ 有设计，✅ 已实现90% | Handler/Service/Repository完整 |
| 3 | MOM3.0_结算管理设计文档.md | ✅ 有设计，❌ 无实现 | 后端代码全空 |
| 4 | MOM3.0_MES生产扩展设计文档.md | ✅ 有设计，❌ 无实现 | 月日计划/班组/能力/异常 |
| 5 | MOM3.0_MES生产执行设计文档.md | ✅ 有设计，⚠️ 部分实现 | 首末件/包装已完成 |
| 6 | MOM3.0_APS滚动排程与交付分析设计.md | ✅ 有设计，⚠️ 部分实现 | 约束算法有枚举未实现 |
| 7 | 本文(缺失功能补充设计.md) | ✅ 新增 | P0/P1详细实现设计 |

---

*文档版本: V1.0 | 创建日期: 2026-04-09*