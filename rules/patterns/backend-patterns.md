# 后端设计模式

> 本文件适用于 MOM3.0 Go 后端项目。

## 1. 分层架构

### 1.1 分层结构

```
┌─────────────────────────────────────┐
│         Handler (控制器层)           │
│  - 请求/响应处理                    │
│  - 参数校验                         │
│  - 权限检查                         │
└─────────────────────────────────────┘
                ↓
┌─────────────────────────────────────┐
│         Service (服务层)            │
│  - 业务逻辑                         │
│  - 事务管理                         │
│  - 领域模型转换                     │
└─────────────────────────────────────┘
                ↓
┌─────────────────────────────────────┐
│        Repository (仓库层)           │
│  - 数据 CRUD                        │
│  - 缓存管理                         │
│  - 复杂查询                         │
└─────────────────────────────────────┘
                ↓
┌─────────────────────────────────────┐
│         Model (模型层)               │
│  - 数据结构                         │
│  - 数据库映射                       │
└─────────────────────────────────────┘
```

### 1.2 代码示例

```go
// Handler - 处理请求
type UserHandler struct {
    userService IService
}

func (h *UserHandler) GetUser(c *gin.Context) {
    id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

    user, err := h.userService.GetUserByID(c.Request.Context(), id)
    if err != nil {
        c.JSON(404, Response{Code: 404, Message: "用户不存在"})
        return
    }
    c.JSON(200, Response{Code: 0, Data: user})
}

// Service - 业务逻辑
type UserService struct {
    repo      IUserRepository
    cache     ICache
    logger    Logger
}

func (s *UserService) GetUserByID(ctx context.Context, id int64) (*User, error) {
    // 1. 先从缓存获取
    cached, err := s.cache.Get(ctx, fmt.Sprintf("user:%d", id))
    if err == nil && cached != "" {
        var user User
        json.Unmarshal([]byte(cached), &user)
        return &user, nil
    }

    // 2. 缓存未命中，从数据库获取
    user, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }

    // 3. 写入缓存
    if data, err := json.Marshal(user); err == nil {
        s.cache.Set(ctx, fmt.Sprintf("user:%d", id), string(data), time.Hour)
    }

    return user, nil
}

// Repository - 数据访问
type UserRepository struct {
    db *gorm.DB
}

func (r *UserRepository) FindByID(ctx context.Context, id int64) (*User, error) {
    var user User
    err := r.db.WithContext(ctx).First(&user, id).Error
    return &user, err
}
```

## 2. Repository 模式

### 2.1 接口定义

```go
// 通用仓储接口
type IBaseRepository interface {
    Create(ctx context.Context, model interface{}) error
    Update(ctx context.Context, id int64, updates map[string]interface{}) error
    Delete(ctx context.Context, id int64) error
    FindByID(ctx context.Context, id int64, model interface{}) error
    FindAll(ctx context.Context, query Query, result interface{}) error
}

// 业务仓储接口
type IUserRepository interface {
    IBaseRepository
    FindByEmail(ctx context.Context, email string) (*User, error)
    FindByRole(ctx context.Context, role string, pagination Pagination) ([]User, int64, error)
    CountByDepartment(ctx context.Context, deptID int64) (int64, error)
}
```

### 2.2 实现原则

- 每个聚合根一个 Repository
- Repository 只做数据访问，不包含业务逻辑
- 使用接口定义，便于 Mock 测试

## 3. 工厂模式

### 3.1 业务工厂

```go
// 订单工厂
type OrderFactory struct {
    repo         IOrderRepository
    materialRepo IMaterialRepository
}

func (f *OrderFactory) CreateProductionOrder(ctx context.Context, req *CreateOrderRequest) (*Order, error) {
    // 1. 验证物料
    material, err := f.materialRepo.FindByID(ctx, req.MaterialID)
    if err != nil {
        return nil, ErrMaterialNotFound
    }

    // 2. 生成订单号
    orderNo := generateOrderNo()

    // 3. 创建订单实体
    order := &Order{
        OrderNo:    orderNo,
        MaterialID: req.MaterialID,
        Quantity:   req.Quantity,
        Status:     OrderStatusPending,
        // ... 其他字段
    }

    // 4. 保存
    if err := f.repo.Create(ctx, order); err != nil {
        return nil, err
    }

    return order, nil
}
```

## 4. 策略模式

### 4.1 调度策略

```go
// 调度策略接口
type SchedulingStrategy interface {
    Schedule(ctx context.Context, tasks []ProductionTask, resources []Resource) ([]ScheduleResult, error)
}

// 遗传算法策略
type GeneticAlgorithmStrategy struct{}

func (s *GeneticAlgorithmStrategy) Schedule(ctx context.Context, tasks []ProductionTask, resources []Resource) ([]ScheduleResult, error) {
    // 遗传算法实现
}

// 粒子群优化策略
type PSOStrategy struct{}

func (s *PSOStrategy) Schedule(ctx context.Context, tasks []ProductionTask, resources []Resource) ([]ScheduleResult, error) {
    // 粒子群优化实现
}

// 策略选择
func GetSchedulingStrategy(algorithm string) SchedulingStrategy {
    switch algorithm {
    case "genetic":
        return &GeneticAlgorithmStrategy{}
    case "pso":
        return &PSOStrategy{}
    default:
        return &GeneticAlgorithmStrategy{}
    }
}
```

## 5. 观察者模式

### 5.1 领域事件

```go
// 事件定义
type DomainEvent interface {
    EventType() string
    OccurredAt() time.Time
}

// 工单完成事件
type OrderCompletedEvent struct {
    OrderID    int64     `json:"order_id"`
    CompletedAt time.Time `json:"completed_at"`
}

func (e *OrderCompletedEvent) EventType() string {
    return "order.completed"
}

func (e *OrderCompletedEvent) OccurredAt() time.Time {
    return e.CompletedAt
}

// 事件处理器
type EventHandler interface {
    Handle(ctx context.Context, event DomainEvent) error
}

// 事件发布器
type EventDispatcher struct {
    handlers map[string][]EventHandler
}

func (d *EventDispatcher) Publish(ctx context.Context, event DomainEvent) error {
    handlers := d.handlers[event.EventType()]
    for _, handler := range handlers {
        if err := handler.Handle(ctx, event); err != nil {
            return err
        }
    }
    return nil
}
```

## 6. Builder 模式

### 6.1 复杂对象构建

```go
// 查询 Builder
type QueryBuilder struct {
    db         *gorm.DB
    conditions []string
    params      []interface{}
    orderBy    string
    limitVal   int
    offsetVal  int
}

func (b *QueryBuilder) Where(condition string, params ...interface{}) *QueryBuilder {
    b.conditions = append(b.conditions, condition)
    b.params = append(b.params, params...)
    return b
}

func (b *QueryBuilder) Order(order string) *QueryBuilder {
    b.orderBy = order
    return b
}

func (b *QueryBuilder) Limit(limit int) *QueryBuilder {
    b.limitVal = limit
    return b
}

func (b *QueryBuilder) Offset(offset int) *QueryBuilder {
    b.offsetVal = offset
    return b
}

func (b *QueryBuilder) Find(result interface{}) error {
    query := b.db.Where(b.conditions[0], b.params[0:]...)
    if b.orderBy != "" {
        query = query.Order(b.orderBy)
    }
    if b.limitVal > 0 {
        query = query.Limit(b.limitVal)
    }
    if b.offsetVal > 0 {
        query = query.Offset(b.offsetVal)
    }
    return query.Find(result).Error
}

// 使用
users := &[]User{}
QueryBuilder{db: db}.
    Where("status = ?", 1).
    Where("department_id = ?", deptID).
    Order("created_at DESC").
    Limit(20).
    Offset(0).
    Find(users)
```

## 7. DDD 限界上下文

### 7.1 上下文划分

```
┌─────────────────────────────────────────────────────┐
│                   MES3.0 限界上下文                  │
├───────────────┬───────────────┬─────────────────────┤
│  生产执行上下文  │  质量上下文    │   设备上下文        │
│  - 生产工单    │  - 检验单     │   - 设备台账        │
│  - 生产报工    │  - 不良品     │   - 维保计划        │
│  - 工序流转    │  - SPC控制    │   - OEE计算         │
├───────────────┼───────────────┼─────────────────────┤
│  APS上下文      │  库存上下文    │   主数据上下文      │
│  - 排程计划    │  - 物料库存    │   - 物料           │
│  - 资源分配    │  - 出入库     │   - BOM            │
│  - MPS/MRP     │  - 库位管理    │   - 工艺路线        │
└───────────────┴───────────────┴─────────────────────┘
```

### 7.2 聚合根示例

```go
// 生产工单聚合根
type ProductionOrder struct {
    ID           int64
    OrderNo      string
    MaterialID   int64
    Quantity     int
    Status       OrderStatus
    Items        []OrderItem      // 聚合内实体
    Operations   []OrderOperation // 聚合内实体
    CreatedAt    time.Time
    UpdatedAt    time.Time
}

// 聚合根不变量
func (o *ProductionOrder) Start() error {
    if o.Status != OrderStatusReady {
        return ErrInvalidOrderStatus
    }
    if len(o.Operations) == 0 {
        return ErrNoOperations
    }
    o.Status = OrderStatusInProgress
    // 发布领域事件
    return nil
}
```

## 8. 缓存模式

### 8.1 多级缓存

```go
type CacheManager struct {
    localCache  *ccache.Cache    // 进程内缓存
    redisCache  redis.Redis      // 分布式缓存
}

func (m *CacheManager) Get(ctx context.Context, key string) (interface{}, error) {
    // 1. 先从本地缓存获取
    if item := m.localCache.Get(key); item != nil {
        return item.Value(), nil
    }

    // 2. 本地缓存未命中，从 Redis 获取
    val, err := m.redisCache.Get(ctx, key).Result()
    if err != nil {
        return nil, err
    }

    // 3. 回填本地缓存
    var result interface{}
    json.Unmarshal([]byte(val), &result)
    m.localCache.Set(key, result, time.Minute*10)

    return result, nil
}
```

## 9. 事务管理

### 9.1 编程式事务

```go
func (s *OrderService) CreateOrderWithItems(ctx context.Context, order *Order, items []OrderItem) error {
    return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        // 1. 创建订单
        if err := tx.Create(order).Error; err != nil {
            return err
        }

        // 2. 创建订单项
        for i := range items {
            items[i].OrderID = order.ID
        }
        if err := tx.Create(&items).Error; err != nil {
            return err
        }

        // 3. 更新库存
        if err := s.updateInventory(tx, order.MaterialID, order.Quantity); err != nil {
            return err
        }

        return nil
    })
}
```
