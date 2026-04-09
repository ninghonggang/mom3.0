package ai_schema

import (
	"strings"
	"sync"
)

// SchemaGenerator 系统Schema生成器
type SchemaGenerator struct {
	mu      sync.RWMutex
	schema  string
	modules map[string]ModuleSchema
}

// ModuleSchema 模块Schema
type ModuleSchema struct {
	Name        string                 // 模块名称
	Description string                 // 模块描述
	Operations  map[string]OperationSchema // 操作列表
}

// OperationSchema 操作Schema
type OperationSchema struct {
	Action         string   // 操作名称
	Method         string   // HTTP方法
	Endpoint       string   // API端点
	Description    string   // 操作描述
	OperationType  string   // 操作类型: query/write/analysis
	Parameters     []ParamSchema // 参数列表
	ResultType     string   // 返回结果类型
}

// ParamSchema 参数Schema
type ParamSchema struct {
	Name        string // 参数名称
	Type        string // 参数类型
	Required    bool   // 是否必填
	Description string // 参数描述
}

// NewSchemaGenerator 创建Schema生成器
func NewSchemaGenerator() *SchemaGenerator {
	sg := &SchemaGenerator{
		modules: make(map[string]ModuleSchema),
	}
	sg.buildSchema()
	return sg
}

// Generate 生成Schema字符串
func (sg *SchemaGenerator) Generate() string {
	sg.mu.RLock()
	defer sg.mu.RUnlock()

	if sg.schema != "" {
		return sg.schema
	}

	sg.mu.RUnlock()
	sg.mu.Lock()
	defer sg.mu.Unlock()

	sg.buildSchema()
	return sg.schema
}

// Refresh 重新生成Schema
func (sg *SchemaGenerator) Refresh() {
	sg.mu.Lock()
	defer sg.mu.Unlock()
	sg.buildSchema()
}

// buildSchema 构建Schema
func (sg *SchemaGenerator) buildSchema() {
	sg.buildModules()

	var sb strings.Builder

	// 生成模块列表
	for _, module := range sg.modules {
		sb.WriteString("\n### " + module.Name + " (" + module.Description + ")\n")
		for _, op := range module.Operations {
			sb.WriteString("- " + op.Action + ": " + op.Method + " " + op.Endpoint + " - " + op.Description + "\n")
		}
	}

	sg.schema = sb.String()
}

// buildModules 构建所有模块的Schema
func (sg *SchemaGenerator) buildModules() {
	// 清空现有数据
	sg.modules = make(map[string]ModuleSchema)

	// 生产执行模块
	sg.modules["production"] = ModuleSchema{
		Name:        "production",
		Description: "生产执行管理",
		Operations: map[string]OperationSchema{
			"list_orders": {
				Action:        "list_orders",
				Method:        "GET",
				Endpoint:      "/api/v1/production/order/list",
				Description:   "查询生产工单列表",
				OperationType: "query",
				Parameters: []ParamSchema{
					{Name: "status", Type: "integer", Required: false, Description: "工单状态: 1=待生产, 2=生产中, 3=已完成, 4=已取消"},
					{Name: "line_id", Type: "integer", Required: false, Description: "产线ID"},
					{Name: "start_date", Type: "string", Required: false, Description: "计划开始日期"},
					{Name: "end_date", Type: "string", Required: false, Description: "计划结束日期"},
				},
				ResultType: "table",
			},
			"get_order": {
				Action:        "get_order",
				Method:        "GET",
				Endpoint:      "/api/v1/production/order/:id",
				Description:   "获取工单详情",
				OperationType: "query",
				Parameters: []ParamSchema{
					{Name: "id", Type: "integer", Required: true, Description: "工单ID"},
				},
				ResultType: "form",
			},
			"create_order": {
				Action:        "create_order",
				Method:        "POST",
				Endpoint:      "/api/v1/production/order",
				Description:   "创建生产工单",
				OperationType: "write",
				Parameters: []ParamSchema{
					{Name: "order_no", Type: "string", Required: true, Description: "工单编号"},
					{Name: "material_id", Type: "integer", Required: true, Description: "物料ID"},
					{Name: "material_code", Type: "string", Required: true, Description: "物料编码"},
					{Name: "material_name", Type: "string", Required: true, Description: "物料名称"},
					{Name: "quantity", Type: "number", Required: true, Description: "计划数量"},
					{Name: "line_id", Type: "integer", Required: false, Description: "产线ID"},
					{Name: "workshop_id", Type: "integer", Required: false, Description: "车间ID"},
					{Name: "plan_start_date", Type: "string", Required: false, Description: "计划开始日期"},
					{Name: "plan_end_date", Type: "string", Required: false, Description: "计划结束日期"},
					{Name: "priority", Type: "integer", Required: false, Description: "优先级: 1=低, 2=中, 3=高"},
				},
				ResultType: "text",
			},
			"update_order": {
				Action:        "update_order",
				Method:        "PUT",
				Endpoint:      "/api/v1/production/order/:id",
				Description:   "更新生产工单",
				OperationType: "write",
				Parameters: []ParamSchema{
					{Name: "id", Type: "integer", Required: true, Description: "工单ID"},
					{Name: "quantity", Type: "number", Required: false, Description: "计划数量"},
					{Name: "priority", Type: "integer", Required: false, Description: "优先级"},
					{Name: "plan_start_date", Type: "string", Required: false, Description: "计划开始日期"},
					{Name: "plan_end_date", Type: "string", Required: false, Description: "计划结束日期"},
				},
				ResultType: "text",
			},
			"delete_order": {
				Action:        "delete_order",
				Method:        "DELETE",
				Endpoint:      "/api/v1/production/order/:id",
				Description:   "删除生产工单",
				OperationType: "write",
				Parameters: []ParamSchema{
					{Name: "id", Type: "integer", Required: true, Description: "工单ID"},
				},
				ResultType: "text",
			},
			"start_order": {
				Action:        "start_order",
				Method:        "PUT",
				Endpoint:      "/api/v1/production/order/:id/start",
				Description:   "开始生产",
				OperationType: "write",
				Parameters: []ParamSchema{
					{Name: "id", Type: "integer", Required: true, Description: "工单ID"},
				},
				ResultType: "text",
			},
			"complete_order": {
				Action:        "complete_order",
				Method:        "PUT",
				Endpoint:      "/api/v1/production/order/:id/complete",
				Description:   "完成生产",
				OperationType: "write",
				Parameters: []ParamSchema{
					{Name: "id", Type: "integer", Required: true, Description: "工单ID"},
				},
				ResultType: "text",
			},
			"create_report": {
				Action:        "create_report",
				Method:        "POST",
				Endpoint:      "/api/v1/production/report",
				Description:   "创建生产报工",
				OperationType: "write",
				Parameters: []ParamSchema{
					{Name: "order_id", Type: "integer", Required: true, Description: "工单ID"},
					{Name: "quantity", Type: "number", Required: true, Description: "报工数量"},
					{Name: "qualified_quantity", Type: "number", Required: false, Description: "合格数量"},
					{Name: "workstation_id", Type: "integer", Required: false, Description: "工位ID"},
					{Name: "remarks", Type: "string", Required: false, Description: "备注"},
				},
				ResultType: "text",
			},
			"order_stats": {
				Action:        "order_stats",
				Method:        "GET",
				Endpoint:      "/api/v1/production/order/stats",
				Description:   "工单统计",
				OperationType: "analysis",
				Parameters:     []ParamSchema{},
				ResultType:     "text",
			},
		},
	}

	// 物料管理模块
	sg.modules["material"] = ModuleSchema{
		Name:        "material",
		Description: "物料主数据管理",
		Operations: map[string]OperationSchema{
			"list_materials": {
				Action:        "list_materials",
				Method:        "GET",
				Endpoint:      "/api/v1/mdm/material/list",
				Description:   "查询物料列表",
				OperationType: "query",
				Parameters: []ParamSchema{
					{Name: "keyword", Type: "string", Required: false, Description: "搜索关键字(编码/名称)"},
					{Name: "category_id", Type: "integer", Required: false, Description: "物料分类ID"},
					{Name: "status", Type: "integer", Required: false, Description: "状态: 1=启用, 0=禁用"},
				},
				ResultType: "table",
			},
			"get_material": {
				Action:        "get_material",
				Method:        "GET",
				Endpoint:      "/api/v1/mdm/material/:id",
				Description:   "获取物料详情",
				OperationType: "query",
				Parameters: []ParamSchema{
					{Name: "id", Type: "integer", Required: true, Description: "物料ID"},
				},
				ResultType: "form",
			},
			"create_material": {
				Action:        "create_material",
				Method:        "POST",
				Endpoint:      "/api/v1/mdm/material",
				Description:   "创建物料",
				OperationType: "write",
				Parameters: []ParamSchema{
					{Name: "material_code", Type: "string", Required: true, Description: "物料编码"},
					{Name: "material_name", Type: "string", Required: true, Description: "物料名称"},
					{Name: "material_type", Type: "string", Required: false, Description: "物料类型"},
					{Name: "unit", Type: "string", Required: false, Description: "单位"},
					{Name: "specification", Type: "string", Required: false, Description: "规格"},
					{Name: "category_id", Type: "integer", Required: false, Description: "分类ID"},
					{Name: "status", Type: "integer", Required: false, Description: "状态"},
				},
				ResultType: "text",
			},
			"update_material": {
				Action:        "update_material",
				Method:        "PUT",
				Endpoint:      "/api/v1/mdm/material/:id",
				Description:   "更新物料",
				OperationType: "write",
				Parameters: []ParamSchema{
					{Name: "id", Type: "integer", Required: true, Description: "物料ID"},
					{Name: "material_name", Type: "string", Required: false, Description: "物料名称"},
					{Name: "specification", Type: "string", Required: false, Description: "规格"},
					{Name: "status", Type: "integer", Required: false, Description: "状态"},
				},
				ResultType: "text",
			},
		},
	}

	// 仓储管理模块
	sg.modules["wms"] = ModuleSchema{
		Name:        "wms",
		Description: "仓储管理",
		Operations: map[string]OperationSchema{
			"list_warehouse": {
				Action:        "list_warehouse",
				Method:        "GET",
				Endpoint:      "/api/v1/wms/warehouse/list",
				Description:   "查询仓库列表",
				OperationType: "query",
				Parameters:     []ParamSchema{},
				ResultType:     "table",
			},
			"list_location": {
				Action:        "list_location",
				Method:        "GET",
				Endpoint:      "/api/v1/wms/location/list",
				Description:   "查询库位列表",
				OperationType: "query",
				Parameters: []ParamSchema{
					{Name: "warehouse_id", Type: "integer", Required: false, Description: "仓库ID"},
				},
				ResultType: "table",
			},
			"list_inventory": {
				Action:        "list_inventory",
				Method:        "GET",
				Endpoint:      "/api/v1/wms/inventory/list",
				Description:   "查询库存列表",
				OperationType: "query",
				Parameters: []ParamSchema{
					{Name: "material_id", Type: "integer", Required: false, Description: "物料ID"},
					{Name: "warehouse_id", Type: "integer", Required: false, Description: "仓库ID"},
					{Name: "location_id", Type: "integer", Required: false, Description: "库位ID"},
				},
				ResultType: "table",
			},
			"get_inventory": {
				Action:        "get_inventory",
				Method:        "GET",
				Endpoint:      "/api/v1/wms/inventory/:id",
				Description:   "获取库存详情",
				OperationType: "query",
				Parameters: []ParamSchema{
					{Name: "id", Type: "integer", Required: true, Description: "库存ID"},
				},
				ResultType: "form",
			},
		},
	}

	// 质量管理模块
	sg.modules["quality"] = ModuleSchema{
		Name:        "quality",
		Description: "质量管理",
		Operations: map[string]OperationSchema{
			"list_iqc": {
				Action:        "list_iqc",
				Method:        "GET",
				Endpoint:      "/api/v1/quality/iqc/list",
				Description:   "查询IQC检验列表",
				OperationType: "query",
				Parameters: []ParamSchema{
					{Name: "status", Type: "integer", Required: false, Description: "状态"},
					{Name: "supplier_id", Type: "integer", Required: false, Description: "供应商ID"},
				},
				ResultType: "table",
			},
			"create_iqc": {
				Action:        "create_iqc",
				Method:        "POST",
				Endpoint:      "/api/v1/quality/iqc",
				Description:   "创建IQC检验单",
				OperationType: "write",
				Parameters: []ParamSchema{
					{Name: "supplier_id", Type: "integer", Required: true, Description: "供应商ID"},
					{Name: "material_id", Type: "integer", Required: true, Description: "物料ID"},
					{Name: "quantity", Type: "number", Required: true, Description: "检验数量"},
					{Name: "batch_no", Type: "string", Required: false, Description: "批次号"},
				},
				ResultType: "text",
			},
			"list_ipqc": {
				Action:        "list_ipqc",
				Method:        "GET",
				Endpoint:      "/api/v1/quality/ipqc/list",
				Description:   "查询IPQC检验列表",
				OperationType: "query",
				Parameters: []ParamSchema{
					{Name: "status", Type: "integer", Required: false, Description: "状态"},
				},
				ResultType: "table",
			},
			"create_ipqc": {
				Action:        "create_ipqc",
				Method:        "POST",
				Endpoint:      "/api/v1/quality/ipqc",
				Description:   "创建IPQC检验单",
				OperationType: "write",
				Parameters: []ParamSchema{
					{Name: "order_id", Type: "integer", Required: true, Description: "工单ID"},
					{Name: "workstation_id", Type: "integer", Required: false, Description: "工位ID"},
					{Name: "quantity", Type: "number", Required: true, Description: "检验数量"},
				},
				ResultType: "text",
			},
			"list_fqc": {
				Action:        "list_fqc",
				Method:        "GET",
				Endpoint:      "/api/v1/quality/fqc/list",
				Description:   "查询FQC检验列表",
				OperationType: "query",
				Parameters:     []ParamSchema{},
				ResultType:     "table",
			},
			"create_fqc": {
				Action:        "create_fqc",
				Method:        "POST",
				Endpoint:      "/api/v1/quality/fqc",
				Description:   "创建FQC检验单",
				OperationType: "write",
				Parameters: []ParamSchema{
					{Name: "order_id", Type: "integer", Required: true, Description: "工单ID"},
					{Name: "quantity", Type: "number", Required: true, Description: "检验数量"},
					{Name: "qualified_quantity", Type: "number", Required: false, Description: "合格数量"},
				},
				ResultType: "text",
			},
			"list_oqc": {
				Action:        "list_oqc",
				Method:        "GET",
				Endpoint:      "/api/v1/quality/oqc/list",
				Description:   "查询OQC检验列表",
				OperationType: "query",
				Parameters:     []ParamSchema{},
				ResultType:     "table",
			},
			"create_oqc": {
				Action:        "create_oqc",
				Method:        "POST",
				Endpoint:      "/api/v1/quality/oqc",
				Description:   "创建OQC检验单",
				OperationType: "write",
				Parameters: []ParamSchema{
					{Name: "order_id", Type: "integer", Required: true, Description: "工单ID"},
					{Name: "quantity", Type: "number", Required: true, Description: "检验数量"},
				},
				ResultType: "text",
			},
		},
	}

	// 设备管理模块
	sg.modules["equipment"] = ModuleSchema{
		Name:        "equipment",
		Description: "设备管理",
		Operations: map[string]OperationSchema{
			"list_equipment": {
				Action:        "list_equipment",
				Method:        "GET",
				Endpoint:      "/api/v1/equipment/list",
				Description:   "查询设备列表",
				OperationType: "query",
				Parameters: []ParamSchema{
					{Name: "status", Type: "string", Required: false, Description: "设备状态"},
					{Name: "line_id", Type: "integer", Required: false, Description: "产线ID"},
				},
				ResultType: "table",
			},
			"get_equipment": {
				Action:        "get_equipment",
				Method:        "GET",
				Endpoint:      "/api/v1/equipment/:id",
				Description:   "获取设备详情",
				OperationType: "query",
				Parameters: []ParamSchema{
					{Name: "id", Type: "integer", Required: true, Description: "设备ID"},
				},
				ResultType: "form",
			},
			"oee_analysis": {
				Action:        "oee_analysis",
				Method:        "GET",
				Endpoint:      "/api/v1/equipment/oee/list",
				Description:   "OEE分析",
				OperationType: "analysis",
				Parameters: []ParamSchema{
					{Name: "equipment_id", Type: "integer", Required: false, Description: "设备ID"},
					{Name: "start_date", Type: "string", Required: false, Description: "开始日期"},
					{Name: "end_date", Type: "string", Required: false, Description: "结束日期"},
				},
				ResultType: "chart",
			},
		},
	}

	// APS模块
	sg.modules["aps"] = ModuleSchema{
		Name:        "aps",
		Description: "高级计划排程",
		Operations: map[string]OperationSchema{
			"list_mps": {
				Action:        "list_mps",
				Method:        "GET",
				Endpoint:      "/api/v1/aps/mps/list",
				Description:   "查询主生产计划列表",
				OperationType: "query",
				Parameters: []ParamSchema{
					{Name: "status", Type: "integer", Required: false, Description: "状态"},
					{Name: "plan_month", Type: "string", Required: false, Description: "计划月份"},
				},
				ResultType: "table",
			},
			"get_mps": {
				Action:        "get_mps",
				Method:        "GET",
				Endpoint:      "/api/v1/aps/mps/:id",
				Description:   "获取MPS详情",
				OperationType: "query",
				Parameters: []ParamSchema{
					{Name: "id", Type: "integer", Required: true, Description: "MPS ID"},
				},
				ResultType: "form",
			},
			"create_mps": {
				Action:        "create_mps",
				Method:        "POST",
				Endpoint:      "/api/v1/aps/mps",
				Description:   "创建MPS",
				OperationType: "write",
				Parameters: []ParamSchema{
					{Name: "plan_month", Type: "string", Required: true, Description: "计划月份"},
					{Name: "material_id", Type: "integer", Required: true, Description: "物料ID"},
					{Name: "quantity", Type: "number", Required: true, Description: "数量"},
				},
				ResultType: "text",
			},
			"list_schedule": {
				Action:        "list_schedule",
				Method:        "GET",
				Endpoint:      "/api/v1/aps/schedule/list",
				Description:   "查询排程计划列表",
				OperationType: "query",
				Parameters:     []ParamSchema{},
				ResultType:     "table",
			},
			"get_schedule_results": {
				Action:        "get_schedule_results",
				Method:        "GET",
				Endpoint:      "/api/v1/aps/schedule/:id/results",
				Description:   "获取排程结果",
				OperationType: "query",
				Parameters: []ParamSchema{
					{Name: "id", Type: "integer", Required: true, Description: "排程计划ID"},
				},
				ResultType: "table",
			},
			"capacity_analysis": {
				Action:        "capacity_analysis",
				Method:        "GET",
				Endpoint:      "/api/v1/aps/capacity/list",
				Description:   "产能分析",
				OperationType: "analysis",
				Parameters: []ParamSchema{
					{Name: "workshop_id", Type: "integer", Required: false, Description: "车间ID"},
					{Name: "line_id", Type: "integer", Required: false, Description: "产线ID"},
				},
				ResultType: "chart",
			},
		},
	}

	// 系统管理模块
	sg.modules["system"] = ModuleSchema{
		Name:        "system",
		Description: "系统管理",
		Operations: map[string]OperationSchema{
			"list_users": {
				Action:        "list_users",
				Method:        "GET",
				Endpoint:      "/api/v1/system/user/list",
				Description:   "查询用户列表",
				OperationType: "query",
				Parameters: []ParamSchema{
					{Name: "keyword", Type: "string", Required: false, Description: "搜索关键字"},
					{Name: "dept_id", Type: "integer", Required: false, Description: "部门ID"},
				},
				ResultType: "table",
			},
			"list_depts": {
				Action:        "list_depts",
				Method:        "GET",
				Endpoint:      "/api/v1/system/dept/list",
				Description:   "查询部门列表",
				OperationType: "query",
				Parameters:     []ParamSchema{},
				ResultType:     "table",
			},
			"list_roles": {
				Action:        "list_roles",
				Method:        "GET",
				Endpoint:      "/api/v1/system/role/list",
				Description:   "查询角色列表",
				OperationType: "query",
				Parameters:     []ParamSchema{},
				ResultType:     "table",
			},
		},
	}
}

// GetModule 获取指定模块的Schema
func (sg *SchemaGenerator) GetModule(moduleName string) *ModuleSchema {
	sg.mu.RLock()
	defer sg.mu.RUnlock()

	if m, ok := sg.modules[moduleName]; ok {
		return &m
	}
	return nil
}

// GetOperation 获取指定模块和操作的Schema
func (sg *SchemaGenerator) GetOperation(moduleName, actionName string) *OperationSchema {
	sg.mu.RLock()
	defer sg.mu.RUnlock()

	if m, ok := sg.modules[moduleName]; ok {
		if op, ok := m.Operations[actionName]; ok {
			return &op
		}
	}
	return nil
}

// GetAllModules 获取所有模块名称
func (sg *SchemaGenerator) GetAllModules() []string {
	sg.mu.RLock()
	defer sg.mu.RUnlock()

	modules := make([]string, 0, len(sg.modules))
	for name := range sg.modules {
		modules = append(modules, name)
	}
	return modules
}
