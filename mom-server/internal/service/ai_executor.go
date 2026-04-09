package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"mom-server/internal/dto"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// OperationPermission 操作权限
type OperationPermission struct {
	Module string   // 模块名
	Action string   // 操作名
	Types  []string // 允许的操作类型
}

// AIExecutor AI操作执行器
type AIExecutor struct {
	// 仓储依赖
	productionOrderRepo *repository.ProductionOrderRepository
	materialRepo       *repository.MaterialRepository
	warehouseRepo      *repository.WarehouseRepository
	inventoryRepo      *repository.InventoryRepository
	iqcRepo            *repository.IQCRepository
	ipqcRepo           *repository.IPQCRepository
	fqcRepo            *repository.FQCRepository
	oqcRepo            *repository.OQCRepository
	equipmentRepo      *repository.EquipmentRepository
	mpsRepo            *repository.MPSRepository
	scheduleRepo       *repository.ScheduleRepository
	userRepo           *repository.UserRepository
	deptRepo           *repository.DeptRepository
	roleRepo           *repository.RoleRepository

	// 权限白名单
	whitelist []OperationPermission

	// 限流
	rateLimiter *RateLimiter
}

// RateLimiter 限流器
type RateLimiter struct {
	mu       sync.Mutex
	requests map[int64][]time.Time
	limit    int
	window   time.Duration
}

// NewRateLimiter 创建限流器
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[int64][]time.Time),
		limit:    limit,
		window:   window,
	}
}

// Allow 检查是否允许请求
func (rl *RateLimiter) Allow(userID int64) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-rl.window)

	// 清理过期记录
	var valid []time.Time
	for _, t := range rl.requests[userID] {
		if t.After(windowStart) {
			valid = append(valid, t)
		}
	}

	if len(valid) >= rl.limit {
		rl.requests[userID] = valid
		return false
	}

	rl.requests[userID] = append(valid, now)
	return true
}

// NewAIExecutor 创建AI执行器
func NewAIExecutor(
	productionOrderRepo *repository.ProductionOrderRepository,
	materialRepo *repository.MaterialRepository,
	warehouseRepo *repository.WarehouseRepository,
	inventoryRepo *repository.InventoryRepository,
	iqcRepo *repository.IQCRepository,
	ipqcRepo *repository.IPQCRepository,
	fqcRepo *repository.FQCRepository,
	oqcRepo *repository.OQCRepository,
	equipmentRepo *repository.EquipmentRepository,
	mpsRepo *repository.MPSRepository,
	scheduleRepo *repository.ScheduleRepository,
	userRepo *repository.UserRepository,
	deptRepo *repository.DeptRepository,
	roleRepo *repository.RoleRepository,
) *AIExecutor {
	return &AIExecutor{
		productionOrderRepo: productionOrderRepo,
		materialRepo:        materialRepo,
		warehouseRepo:       warehouseRepo,
		inventoryRepo:       inventoryRepo,
		iqcRepo:             iqcRepo,
		ipqcRepo:            ipqcRepo,
		fqcRepo:             fqcRepo,
		oqcRepo:             oqcRepo,
		equipmentRepo:       equipmentRepo,
		mpsRepo:             mpsRepo,
		scheduleRepo:        scheduleRepo,
		userRepo:            userRepo,
		deptRepo:            deptRepo,
		roleRepo:            roleRepo,
		whitelist:           buildWhitelist(),
		rateLimiter:         NewRateLimiter(5, time.Minute), // 5次/分钟
	}
}

// ExecutionResult 执行结果
type ExecutionResult struct {
	Success     bool                   `json:"success"`
	Data        interface{}            `json:"data,omitempty"`
	Error       string                 `json:"error,omitempty"`
	ResultType  string                 `json:"result_type"` // table | chart | text | error
	Operation   *AIIntent              `json:"operation,omitempty"`
	NeedsConfirm bool                   `json:"needs_confirmation"`
}

// ExecuteIntent 执行AI意图
func (e *AIExecutor) ExecuteIntent(ctx context.Context, intent *AIIntent, tenantID, userID int64) (*ExecutionResult, error) {
	// 1. 限流检查
	if !e.rateLimiter.Allow(userID) {
		return &ExecutionResult{
			Success:    false,
			Error:      "请求过于频繁，请稍后再试",
			ResultType: "error",
		}, nil
	}

	// 2. 验证权限
	if !e.ValidatePermission(intent, userID) {
		return &ExecutionResult{
			Success:    false,
			Error:      "您没有权限执行此操作",
			ResultType: "error",
		}, nil
	}

	// 3. 分类处理
	switch intent.OperationType {
	case "query":
		return e.executeQuery(ctx, intent, tenantID)
	case "write":
		return &ExecutionResult{
			Success:     true,
			NeedsConfirm: true,
			ResultType:  "confirmation",
			Operation:   intent,
		}, nil
	case "analysis":
		return e.executeAnalysis(ctx, intent, tenantID)
	default:
		return &ExecutionResult{
			Success:    false,
			Error:      "无法识别的操作类型",
			ResultType: "error",
		}, nil
	}
}

// ValidatePermission 验证操作权限
func (e *AIExecutor) ValidatePermission(intent *AIIntent, userID int64) bool {
	// 检查白名单
	for _, perm := range e.whitelist {
		if perm.Module == intent.Module && perm.Action == intent.Action {
			for _, t := range perm.Types {
				if t == intent.OperationType {
					return true
				}
			}
		}
	}
	return false
}

// executeQuery 执行查询操作
func (e *AIExecutor) executeQuery(ctx context.Context, intent *AIIntent, tenantID int64) (*ExecutionResult, error) {
	switch intent.Module {
	case "production":
		return e.executeProductionQuery(ctx, intent, tenantID)
	case "material":
		return e.executeMaterialQuery(ctx, intent, tenantID)
	case "wms":
		return e.executeWMSQuery(ctx, intent, tenantID)
	case "quality":
		return e.executeQualityQuery(ctx, intent, tenantID)
	case "equipment":
		return e.executeEquipmentQuery(ctx, intent, tenantID)
	case "aps":
		return e.executeAPSQuery(ctx, intent, tenantID)
	case "system":
		return e.executeSystemQuery(ctx, intent, tenantID)
	default:
		return &ExecutionResult{
			Success:    false,
			Error:      fmt.Sprintf("未知模块: %s", intent.Module),
			ResultType: "error",
		}, nil
	}
}

// executeProductionQuery 执行生产模块查询
func (e *AIExecutor) executeProductionQuery(ctx context.Context, intent *AIIntent, tenantID int64) (*ExecutionResult, error) {
	switch intent.Action {
	case "list_orders":
		orders, total, err := e.productionOrderRepo.List(ctx, 0)
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"list": orders, "total": total},
			ResultType: "table",
		}, nil

	case "get_order":
		id := getInt64Param(intent.Parameters, "id")
		if id == 0 {
			return &ExecutionResult{Success: false, Error: "缺少参数: id", ResultType: "error"}, nil
		}
		order, err := e.productionOrderRepo.GetByID(ctx, uint(id))
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      order,
			ResultType: "text",
		}, nil

	default:
		return &ExecutionResult{
			Success:    false,
			Error:      fmt.Sprintf("不支持的生产操作: %s", intent.Action),
			ResultType: "error",
		}, nil
	}
}

// executeMaterialQuery 执行物料模块查询
func (e *AIExecutor) executeMaterialQuery(ctx context.Context, intent *AIIntent, tenantID int64) (*ExecutionResult, error) {
	switch intent.Action {
	case "list_materials":
		materials, _, err := e.materialRepo.List(ctx, tenantID)
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"list": materials},
			ResultType: "table",
		}, nil

	case "get_material":
		id := getInt64Param(intent.Parameters, "id")
		if id == 0 {
			return &ExecutionResult{Success: false, Error: "缺少参数: id", ResultType: "error"}, nil
		}
		material, err := e.materialRepo.GetByID(ctx, uint(id))
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      material,
			ResultType: "text",
		}, nil

	default:
		return &ExecutionResult{
			Success:    false,
			Error:      fmt.Sprintf("不支持的物料操作: %s", intent.Action),
			ResultType: "error",
		}, nil
	}
}

// executeWMSQuery 执行仓储模块查询
func (e *AIExecutor) executeWMSQuery(ctx context.Context, intent *AIIntent, tenantID int64) (*ExecutionResult, error) {
	switch intent.Action {
	case "list_warehouse":
		warehouses, _, err := e.warehouseRepo.List(ctx, tenantID)
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"list": warehouses},
			ResultType: "table",
		}, nil

	case "list_inventory":
		// 简化实现，实际应该有关键字过滤等
		// 这里不做实际查询，返回提示信息
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"message": "请使用具体的物料ID查询库存"},
			ResultType: "text",
		}, nil

	default:
		return &ExecutionResult{
			Success:    false,
			Error:      fmt.Sprintf("不支持的仓储操作: %s", intent.Action),
			ResultType: "error",
		}, nil
	}
}

// executeQualityQuery 执行质量模块查询
func (e *AIExecutor) executeQualityQuery(ctx context.Context, intent *AIIntent, tenantID int64) (*ExecutionResult, error) {
	switch intent.Action {
	case "list_iqc":
		records, _, err := e.iqcRepo.List(ctx, tenantID)
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"list": records},
			ResultType: "table",
		}, nil

	case "list_ipqc":
		records, _, err := e.ipqcRepo.List(ctx, tenantID)
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"list": records},
			ResultType: "table",
		}, nil

	case "list_fqc":
		records, _, err := e.fqcRepo.List(ctx, tenantID)
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"list": records},
			ResultType: "table",
		}, nil

	case "list_oqc":
		records, _, err := e.oqcRepo.List(ctx, tenantID)
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"list": records},
			ResultType: "table",
		}, nil

	default:
		return &ExecutionResult{
			Success:    false,
			Error:      fmt.Sprintf("不支持的质量操作: %s", intent.Action),
			ResultType: "error",
		}, nil
	}
}

// executeEquipmentQuery 执行设备模块查询
func (e *AIExecutor) executeEquipmentQuery(ctx context.Context, intent *AIIntent, tenantID int64) (*ExecutionResult, error) {
	switch intent.Action {
	case "list_equipment":
		equipment, _, err := e.equipmentRepo.List(ctx, tenantID)
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"list": equipment},
			ResultType: "table",
		}, nil

	case "get_equipment":
		id := getInt64Param(intent.Parameters, "id")
		if id == 0 {
			return &ExecutionResult{Success: false, Error: "缺少参数: id", ResultType: "error"}, nil
		}
		equip, err := e.equipmentRepo.GetByID(ctx, uint(id))
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      equip,
			ResultType: "text",
		}, nil

	default:
		return &ExecutionResult{
			Success:    false,
			Error:      fmt.Sprintf("不支持的设备操作: %s", intent.Action),
			ResultType: "error",
		}, nil
	}
}

// executeAPSQuery 执行APS模块查询
func (e *AIExecutor) executeAPSQuery(ctx context.Context, intent *AIIntent, tenantID int64) (*ExecutionResult, error) {
	switch intent.Action {
	case "list_mps":
		mpsList, _, err := e.mpsRepo.List(ctx, tenantID)
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"list": mpsList},
			ResultType: "table",
		}, nil

	case "get_mps":
		id := getInt64Param(intent.Parameters, "id")
		if id == 0 {
			return &ExecutionResult{Success: false, Error: "缺少参数: id", ResultType: "error"}, nil
		}
		mps, err := e.mpsRepo.GetByID(ctx, uint(id))
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      mps,
			ResultType: "text",
		}, nil

	case "list_schedule":
		schedules, _, err := e.scheduleRepo.List(ctx, tenantID)
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"list": schedules},
			ResultType: "table",
		}, nil

	default:
		return &ExecutionResult{
			Success:    false,
			Error:      fmt.Sprintf("不支持的APS操作: %s", intent.Action),
			ResultType: "error",
		}, nil
	}
}

// executeSystemQuery 执行系统模块查询
func (e *AIExecutor) executeSystemQuery(ctx context.Context, intent *AIIntent, tenantID int64) (*ExecutionResult, error) {
	switch intent.Action {
	case "list_users":
		users, _, err := e.userRepo.FindByPage(ctx, tenantID, dto.PageRequest{Page: 1, PageSize: 100}, "", "")
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"list": users},
			ResultType: "table",
		}, nil

	case "list_depts":
		depts, err := e.deptRepo.List(ctx, tenantID)
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"list": depts},
			ResultType: "table",
		}, nil

	case "list_roles":
		roles, err := e.roleRepo.FindAll(ctx, tenantID)
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"list": roles},
			ResultType: "table",
		}, nil

	default:
		return &ExecutionResult{
			Success:    false,
			Error:      fmt.Sprintf("不支持的系统操作: %s", intent.Action),
			ResultType: "error",
		}, nil
	}
}

// executeAnalysis 执行分析操作
func (e *AIExecutor) executeAnalysis(ctx context.Context, intent *AIIntent, tenantID int64) (*ExecutionResult, error) {
	switch intent.Module {
	case "production":
		return e.executeProductionAnalysis(ctx, intent, tenantID)
	case "equipment":
		return e.executeEquipmentAnalysis(ctx, intent, tenantID)
	case "aps":
		return e.executeAPSAnalysis(ctx, intent, tenantID)
	default:
		return &ExecutionResult{
			Success:    false,
			Error:      fmt.Sprintf("不支持的分析模块: %s", intent.Module),
			ResultType: "error",
		}, nil
	}
}

// executeProductionAnalysis 执行生产分析
func (e *AIExecutor) executeProductionAnalysis(ctx context.Context, intent *AIIntent, tenantID int64) (*ExecutionResult, error) {
	switch intent.Action {
	case "order_stats":
		// 简单的工单统计
		orders, _, err := e.productionOrderRepo.List(ctx, 0)
		if err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}

		var total, pending, inProgress, completed int
		for _, o := range orders {
			total++
			switch o.Status {
			case 1:
				pending++
			case 2:
				inProgress++
			case 3:
				completed++
			}
		}

		return &ExecutionResult{
			Success:   true,
			Data:      map[string]interface{}{"total": total, "pending": pending, "in_progress": inProgress, "completed": completed},
			ResultType: "text",
		}, nil

	default:
		return &ExecutionResult{
			Success:    false,
			Error:      fmt.Sprintf("不支持的生产分析: %s", intent.Action),
			ResultType: "error",
		}, nil
	}
}

// executeEquipmentAnalysis 执行设备分析
func (e *AIExecutor) executeEquipmentAnalysis(ctx context.Context, intent *AIIntent, tenantID int64) (*ExecutionResult, error) {
	switch intent.Action {
	case "oee_analysis":
		// OEE分析
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"message": "OEE分析需要设备运行数据，请使用设备OEE功能"},
			ResultType: "text",
		}, nil

	default:
		return &ExecutionResult{
			Success:    false,
			Error:      fmt.Sprintf("不支持的设备分析: %s", intent.Action),
			ResultType: "error",
		}, nil
	}
}

// executeAPSAnalysis 执行APS分析
func (e *AIExecutor) executeAPSAnalysis(ctx context.Context, intent *AIIntent, tenantID int64) (*ExecutionResult, error) {
	switch intent.Action {
	case "capacity_analysis":
		// 产能分析
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"message": "产能分析请使用APS产能分析功能"},
			ResultType: "text",
		}, nil

	default:
		return &ExecutionResult{
			Success:    false,
			Error:      fmt.Sprintf("不支持的APS分析: %s", intent.Action),
			ResultType: "error",
		}, nil
	}
}

// ExecuteConfirmedOperation 执行确认的写操作
func (e *AIExecutor) ExecuteConfirmedOperation(ctx context.Context, intent *AIIntent, tenantID, userID int64) (*ExecutionResult, error) {
	// 再次验证权限
	if !e.ValidatePermission(intent, userID) {
		return &ExecutionResult{
			Success:    false,
			Error:      "您没有权限执行此操作",
			ResultType: "error",
		}, nil
	}

	switch intent.Module {
	case "production":
		return e.executeProductionWrite(ctx, intent, tenantID)
	case "material":
		return e.executeMaterialWrite(ctx, intent, tenantID)
	default:
		return &ExecutionResult{
			Success:    false,
			Error:      fmt.Sprintf("暂不支持模块 %s 的写操作", intent.Module),
			ResultType: "error",
		}, nil
	}
}

// executeProductionWrite 执行生产写操作
func (e *AIExecutor) executeProductionWrite(ctx context.Context, intent *AIIntent, tenantID int64) (*ExecutionResult, error) {
	switch intent.Action {
	case "create_order":
		// 创建立即工单
		data, err := json.Marshal(intent.Parameters)
		if err != nil {
			return &ExecutionResult{Success: false, Error: "参数错误", ResultType: "error"}, nil
		}
		var order model.ProductionOrder
		if err := json.Unmarshal(data, &order); err != nil {
			return &ExecutionResult{Success: false, Error: "参数错误", ResultType: "error"}, nil
		}
		order.TenantID = tenantID
		if err := e.productionOrderRepo.Create(ctx, &order); err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"id": order.ID, "order_no": order.OrderNo},
			ResultType: "text",
		}, nil

	default:
		return &ExecutionResult{
			Success:    false,
			Error:      fmt.Sprintf("不支持的生产写操作: %s", intent.Action),
			ResultType: "error",
		}, nil
	}
}

// executeMaterialWrite 执行物料写操作
func (e *AIExecutor) executeMaterialWrite(ctx context.Context, intent *AIIntent, tenantID int64) (*ExecutionResult, error) {
	switch intent.Action {
	case "create_material":
		data, err := json.Marshal(intent.Parameters)
		if err != nil {
			return &ExecutionResult{Success: false, Error: "参数错误", ResultType: "error"}, nil
		}
		var material model.Material
		if err := json.Unmarshal(data, &material); err != nil {
			return &ExecutionResult{Success: false, Error: "参数错误", ResultType: "error"}, nil
		}
		material.TenantID = tenantID
		if err := e.materialRepo.Create(ctx, &material); err != nil {
			return &ExecutionResult{Success: false, Error: err.Error(), ResultType: "error"}, nil
		}
		return &ExecutionResult{
			Success:    true,
			Data:      map[string]interface{}{"id": material.ID, "material_code": material.MaterialCode},
			ResultType: "text",
		}, nil

	default:
		return &ExecutionResult{
			Success:    false,
			Error:      fmt.Sprintf("不支持的物料写操作: %s", intent.Action),
			ResultType: "error",
		}, nil
	}
}

// buildWhitelist 构建权限白名单
func buildWhitelist() []OperationPermission {
	return []OperationPermission{
		// 生产模块
		{Module: "production", Action: "list_orders", Types: []string{"query"}},
		{Module: "production", Action: "get_order", Types: []string{"query"}},
		{Module: "production", Action: "create_order", Types: []string{"write"}},
		{Module: "production", Action: "update_order", Types: []string{"write"}},
		{Module: "production", Action: "delete_order", Types: []string{"write"}},
		{Module: "production", Action: "create_report", Types: []string{"write"}},
		{Module: "production", Action: "order_stats", Types: []string{"analysis"}},

		// 物料模块
		{Module: "material", Action: "list_materials", Types: []string{"query"}},
		{Module: "material", Action: "get_material", Types: []string{"query"}},
		{Module: "material", Action: "create_material", Types: []string{"write"}},
		{Module: "material", Action: "update_material", Types: []string{"write"}},

		// 仓储模块
		{Module: "wms", Action: "list_inventory", Types: []string{"query"}},
		{Module: "wms", Action: "get_inventory", Types: []string{"query"}},
		{Module: "wms", Action: "list_warehouse", Types: []string{"query"}},
		{Module: "wms", Action: "list_location", Types: []string{"query"}},

		// 质量模块
		{Module: "quality", Action: "list_iqc", Types: []string{"query"}},
		{Module: "quality", Action: "create_iqc", Types: []string{"write"}},
		{Module: "quality", Action: "list_ipqc", Types: []string{"query"}},
		{Module: "quality", Action: "create_ipqc", Types: []string{"write"}},
		{Module: "quality", Action: "list_fqc", Types: []string{"query"}},
		{Module: "quality", Action: "create_fqc", Types: []string{"write"}},
		{Module: "quality", Action: "list_oqc", Types: []string{"query"}},
		{Module: "quality", Action: "create_oqc", Types: []string{"write"}},

		// 设备模块
		{Module: "equipment", Action: "list_equipment", Types: []string{"query"}},
		{Module: "equipment", Action: "get_equipment", Types: []string{"query"}},
		{Module: "equipment", Action: "oee_analysis", Types: []string{"analysis"}},

		// APS模块
		{Module: "aps", Action: "list_mps", Types: []string{"query"}},
		{Module: "aps", Action: "get_mps", Types: []string{"query"}},
		{Module: "aps", Action: "list_schedule", Types: []string{"query"}},
		{Module: "aps", Action: "capacity_analysis", Types: []string{"analysis"}},

		// 系统模块
		{Module: "system", Action: "list_users", Types: []string{"query"}},
		{Module: "system", Action: "list_depts", Types: []string{"query"}},
		{Module: "system", Action: "list_roles", Types: []string{"query"}},
	}
}

// getInt64Param 从参数map中获取int64参数
func getInt64Param(params map[string]interface{}, key string) int64 {
	if v, ok := params[key]; ok {
		switch val := v.(type) {
		case float64:
			return int64(val)
		case int:
			return int64(val)
		case int64:
			return val
		case string:
			var i int64
			fmt.Sscanf(val, "%d", &i)
			return i
		}
	}
	return 0
}

// ValidateOperationType 验证操作类型是否在白名单中
func (e *AIExecutor) ValidateOperationType(module, action, opType string) bool {
	for _, perm := range e.whitelist {
		if perm.Module == module && perm.Action == action {
			for _, t := range perm.Types {
				if t == opType {
					return true
				}
			}
		}
	}
	return false
}

// IsModuleAllowed 检查模块是否允许AI调用
func (e *AIExecutor) IsModuleAllowed(module string) bool {
	allowedModules := map[string]bool{
		"production": true,
		"material":   true,
		"wms":       true,
		"quality":   true,
		"equipment": true,
		"aps":       true,
		"system":    false, // 系统模块需要特殊权限
	}
	return allowedModules[module]
}

// ParseNaturalDate 解析自然语言日期
func ParseNaturalDate(text string) *time.Time {
	lower := strings.ToLower(text)
	now := time.Now()

	if strings.Contains(lower, "今天") {
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		return &today
	}
	if strings.Contains(lower, "昨天") {
		yesterday := now.AddDate(0, 0, -1)
		t := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, yesterday.Location())
		return &t
	}
	if strings.Contains(lower, "明天") {
		tomorrow := now.AddDate(0, 0, 1)
		t := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, tomorrow.Location())
		return &t
	}
	if strings.Contains(lower, "本周") {
		weekday := int(now.Weekday())
		if weekday == 0 {
			weekday = 7
		}
		weekStart := now.AddDate(0, 0, -(weekday - 1))
		t := time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, weekStart.Location())
		return &t
	}
	if strings.Contains(lower, "本月") {
		t := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		return &t
	}

	return nil
}
