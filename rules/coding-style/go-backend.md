# Go 后端编码规范

> 本文件继承 [common/coding-style.md](../../common/coding-style.md) 并针对 MOM3.0 Go后端项目进行扩展。

## 1. 项目结构

```
mom-server/
├── cmd/
│   └── main.go                    # 程序入口
├── internal/
│   ├── config/                    # 配置加载
│   ├── middleware/                # JWT、RBAC、租户、限流、日志中间件
│   ├── router/                    # 路由注册（按模块拆分）
│   ├── handler/                   # HTTP Handler层（薄层，只做参数校验和响应）
│   │   ├── system/               # 系统管理
│   │   ├── mdm/                  # 主数据
│   │   ├── production/            # 生产执行
│   │   ├── aps/                  # 计划排程
│   │   ├── quality/              # 质量管理
│   │   ├── equipment/             # 设备管理
│   │   ├── wms/                  # 仓储管理
│   │   ├── trace/                # 追溯管理
│   │   ├── andon/                # 安东系统
│   │   ├── datacollect/          # 数据采集
│   │   ├── report/               # 报表
│   │   ├── lowcode/              # 低代码
│   │   ├── ai/                   # AI服务
│   │   └── energy/               # 能源管理
│   ├── service/                   # 业务逻辑层（核心）
│   ├── repository/                # 数据访问层（GORM）
│   ├── model/                     # 数据模型（GORM结构体）
│   ├── dto/                      # 数据传输对象
│   ├── websocket/                 # WebSocket Hub
│   └── pkg/                      # 公共工具包
│       ├── response/              # 统一响应封装
│       ├── jwt/                   # JWT工具
│       ├── casbin/                # RBAC权限引擎
│       ├── upload/                # 文件上传
│       └── excel/                 # Excel导入导出
├── migrations/                    # 数据库迁移文件
├── docker-compose.yml
├── Dockerfile
└── config.yaml
```

### 分层原则

- **Handler (控制器)**: 负责请求/响应处理、参数校验、调用Service
- **Service (服务)**: 核心业务逻辑、事务管理
- **Repository (仓库)**: 数据库操作、缓存读写
- **Model (模型)**: 数据结构定义

## 2. 命名规范

### 文件命名
- 使用下划线分隔: `user_handler.go`, `permission_service.go`
- 测试文件: `user_handler_test.go`

### 结构体命名
- 使用驼峰命名: `User`, `ProductionOrder`, `MaterialInfo`
- 接口命名: `IUserService`, `IMaterialRepository` (以 I 开头)
- 枚举类型: `OrderStatus`, `TaskType` (Status/Type 后缀)

### 函数命名
- 公开函数: 驼峰 `GetUserByID`
- 私有函数: 驼峰 `calculateOrderPriority`
- getter/setter: `GetName()`, `SetName()`

### 变量命名
- 短变量在循环和错误检查中使用: `i`, `err`, `v`
- 函数参数使用有意义的名称: `userID`, `orderStatus`
- 避免使用无意义的 `tmp`, `temp` (除非临时交换)

## 3. 函数设计

### 函数长度 (CRITICAL)
- 单个函数不超过 **50行**
- 理想情况: **20-30行**
- 超过50行必须拆分

### 参数数量 (HIGH)
- 函数参数不超过 **4个**
- 超过4个使用结构体封装

### 错误处理模式 (CRITICAL)

```go
// ✅ 正确: 明确处理每个错误
func (s *UserService) GetUser(ctx context.Context, id int64) (*User, error) {
    user, err := s.repo.FindByID(ctx, id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrUserNotFound
        }
        s.logger.Error("failed to get user", "id", id, "error", err)
        return nil, ErrInternalServer
    }
    return user, nil
}

// ❌ 错误: 忽略错误
func GetUser(id int64) *User {
    user, _ := repo.FindByID(id)  // 禁止!
    return user
}
```

### Context 使用 (CRITICAL)

```go
// ✅ 正确: 所有数据库/外部调用必须传递 context
func (s *UserService) GetUser(ctx context.Context, id int64) (*User, error) {
    return s.repo.FindByID(ctx, id)
}

// ❌ 错误: 忽略 context
func GetUser(id int64) (*User, error) {
    return repo.FindByID(context.Background(), id)  // 禁止!
}
```

## 4. 数据库操作

### GORM 约定 (CRITICAL)

```go
// 使用指针处理 nullable 字段
type User struct {
    ID        int64        `json:"id" gorm:"primaryKey"`
    TenantID  int64        `json:"tenant_id" gorm:"index;not null"`
    Username  string       `json:"username" gorm:"size:50;not null"`
    Nickname  string       `json:"nickname" gorm:"size:50"`
    Password  string       `json:"-" gorm:"size:200;not null"`
    Email     *string      `json:"email" gorm:"size:100"`      // nullable
    Phone     *string      `json:"phone" gorm:"size:20"`        // nullable
    Avatar    *string      `json:"avatar" gorm:"size:500"`      // nullable
    DeptID    *int64       `json:"dept_id" gorm:"index"`        // nullable
    Status    int          `json:"status" gorm:"default:1"`     // 1正常 0停用
    LoginIP   *string      `json:"login_ip" gorm:"size:128"`    // nullable
    LoginDate *time.Time   `json:"login_date"`                  // nullable
    CreatedAt time.Time   `json:"created_at"`
    UpdatedAt time.Time   `json:"updated_at"`
}

// 软删除
type User struct {
    // ...
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// 联合唯一索引
type User struct {
    // ...
    UNIQUE(tenant_id, username)
}
```

### 租户隔离 (CRITICAL)

```go
// 所有查询必须包含 tenant_id（通过中间件自动注入）
func (r *UserRepository) FindByID(ctx context.Context, userID int64) (*User, error) {
    var user User
    err := r.db.WithContext(ctx).
        Where("id = ?", userID).
        First(&user).Error
    return &user, err
}

// 跨租户操作（管理员）
func (r *UserRepository) FindByIDWithTenant(ctx context.Context, tenantID, userID int64) (*User, error) {
    var user User
    err := r.db.WithContext(ctx).
        Where("id = ? AND tenant_id = ?", userID, tenantID).
        First(&user).Error
    return &user, err
}
```

### 预加载关系 (HIGH)

```go
// 使用 Preload 避免 N+1 查询
func (s *OrderService) GetOrderWithItems(ctx context.Context, id int64) (*Order, error) {
    return s.repo.FindWithItems(ctx, id)
}

// Repository 实现
func (r *OrderRepository) FindWithItems(ctx context.Context, id int64) (*Order, error) {
    var order Order
    err := r.db.WithContext(ctx).
        Preload("Items").
        Preload("Items.Product").
        First(&order, id).Error
    return &order, err
}
```

## 5. API 设计

### 统一响应格式 (CRITICAL)

```go
// 成功响应
type Response struct {
    Code    int         `json:"code"`      // 200=成功
    Message string      `json:"message"`   // "success"
    Data    interface{} `json:"data,omitempty"`
}

// 分页响应
type PageResponse struct {
    List     interface{} `json:"list"`
    Total    int64      `json:"total"`
    Page     int        `json:"page"`
    PageSize int        `json:"page_size"`
}

// 错误响应
type ErrorResponse struct {
    Code    int         `json:"code"`      // 非200=失败，如 40001
    Message string      `json:"message"`    // 错误信息
    Data    interface{} `json:"data"`       // null
}

// 使用封装函数
func Success(c *gin.Context, data interface{}) {
    c.JSON(200, gin.H{
        "code":    200,
        "message": "success",
        "data":    data,
    })
}

func SuccessPage(c *gin.Context, list interface{}, total int64, page, pageSize int) {
    Success(c, gin.H{
        "list":      list,
        "total":     total,
        "page":      page,
        "pageSize":  pageSize,
    })
}

func Error(c *gin.Context, code int, message string) {
    c.JSON(200, gin.H{
        "code":    code,
        "message": message,
        "data":    nil,
    })
}
```

### 路由命名 (CRITICAL)

- 资源使用复数: `/api/v1/users`, `/api/v1/orders`, `/api/v1/materials`
- 嵌套资源: `/api/v1/orders/{orderID}/items`
- 操作使用 REST 动词: `GET/POST/PUT/DELETE`

### 状态码 (HIGH)

| Code | 含义 |
|------|------|
| 200 | 成功 |
| 40001 | 请求参数错误 |
| 40101 | 未授权 |
| 40301 | 无权限 |
| 40401 | 资源不存在 |
| 50001 | 服务器内部错误 |

## 6. 认证与权限

### JWT 双令牌

```go
// Token Claims
type Claims struct {
    UserID   int64   `json:"user_id"`
    TenantID int64   `json:"tenant_id"`
    Username string  `json:"username"`
    Roles    []string `json:"roles"`
    ExpiresAt int64  `json:"expires_at"`
    jwt.RegisteredClaims
}

// Access Token: 2小时有效期
// Refresh Token: 7天有效期
```

### Casbin 权限模型

```go
// 权限模型 conf
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

// 策略示例
p, admin, /api/v1/system/user, *
p, user, /api/v1/system/user, GET
```

## 7. WebSocket 规范

```go
// Topic 定义
const (
    TopicProductionReport = "production.report"  // 工序报工完成
    TopicQualityAlert     = "quality.alert"       // 质量预警触发
    TopicEquipmentStatus  = "equipment.status"     // 设备状态变化
    TopicAndonCall       = "andon.call"          // 安东呼叫创建
    TopicWmsStockLow    = "wms.stock.low"        // 库存低于预警线
    TopicOrderStatus    = "order.status"          // 工单状态变更
)

// 消息格式
type WSMessage struct {
    Type    string      `json:"type"`
    Topic   string      `json:"topic"`
    Payload interface{} `json:"payload"`
}
```

## 8. 错误定义

```go
// 业务错误定义
var (
    ErrUserNotFound     = errors.New("user not found")
    ErrInvalidParameter = errors.New("invalid parameter")
    ErrUnauthorized     = errors.New("unauthorized")
    ErrForbidden        = errors.New("forbidden")
)

// 错误码
const (
    CodeSuccess        = 200
    CodeParamError     = 40001
    CodeUnauthorized   = 40101
    CodeForbidden      = 40301
    CodeNotFound       = 40401
    CodeInternalError  = 50001
)

// 错误包装
func (s *Service) DoSomething() error {
    if err := do(); err != nil {
        return fmt.Errorf("do something failed: %w", err)
    }
}
```

## 9. 日志规范

```go
// 使用 Zap 结构化日志
logger.Info("user created",
    zap.Int64("user_id", user.ID),
    zap.Int64("tenant_id", user.TenantID),
    zap.String("email", *user.Email),
)

// 错误日志必须包含上下文
logger.Error("failed to create order",
    zap.Int64("tenant_id", tenantID),
    zap.String("order_no", orderNo),
    zap.Error(err),
)
```

## 10. 并发安全

```go
// 使用 sync.Map 或 sync.RWMutex 保护共享资源
type Cache struct {
    mu    sync.RWMutex
    items map[string]interface{}
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    v, ok := c.items[key]
    return v, ok
}
```

## 11. 单元测试

```go
func TestUserService_GetUserByID(t *testing.T) {
    // Arrange
    mockRepo := &mocks.MockUserRepository{}
    service := NewUserService(mockRepo, logger)

    mockRepo.On("FindByID", mock.Anything, int64(1)).Return(&User{ID: 1, Name: "test"}, nil)

    // Act
    user, err := service.GetUserByID(context.Background(), 1)

    // Assert
    assert.NoError(t, err)
    assert.Equal(t, "test", user.Name)
    mockRepo.AssertExpectations(t)
}
```

## 12. 导入顺序 (HIGH)

```go
import (
    // 标准库
    "context"
    "fmt"
    "time"

    // 第三方库
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    "gorm.io/gorm"

    // 项目内部包
    "mom-server/internal/config"
    "mom-server/internal/models"
)
```

## 13. 禁止的模式

```go
// ❌ 禁止: 硬编码敏感信息
apiKey := "sk-xxx"  // 禁止!

// ✅ 正确: 从配置读取
apiKey := viper.GetString("openai.api_key")

// ❌ 禁止: 裸 return
if err != nil {
    return err
}
return nil  // 缺少日志

// ❌ 禁止: 暴露内部错误
return nil, err  // 直接暴露原始错误

// ✅ 正确: 包装错误
return nil, fmt.Errorf("get user failed: %w", err)

// ❌ 禁止: 全局变量
var globalCache map[string]interface{}  // 禁止!

// ✅ 正确: 依赖注入
type Service struct {
    cache Cache  // 接口
}
```
