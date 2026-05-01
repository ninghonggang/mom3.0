-- MOM3.0 数据库初始化脚本
-- 创建扩展
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ============================================
-- 系统管理模块
-- ============================================

-- 租户表
CREATE TABLE IF NOT EXISTS sys_tenant (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_name VARCHAR(100) NOT NULL,
    tenant_key VARCHAR(100) NOT NULL UNIQUE,
    contact VARCHAR(50),
    phone VARCHAR(20),
    email VARCHAR(100),
    status INTEGER DEFAULT 1,
    expire_time TIMESTAMP,
    package_id BIGINT
);

-- 用户表
CREATE TABLE IF NOT EXISTS sys_user (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    username VARCHAR(50) NOT NULL,
    nickname VARCHAR(50),
    password VARCHAR(200) NOT NULL,
    email VARCHAR(100),
    phone VARCHAR(20),
    avatar VARCHAR(500),
    dept_id BIGINT,
    status INTEGER DEFAULT 1,
    login_ip VARCHAR(128),
    login_date TIMESTAMP,
    UNIQUE(tenant_id, username)
);

-- 角色表
CREATE TABLE IF NOT EXISTS sys_role (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    role_name VARCHAR(50) NOT NULL,
    role_key VARCHAR(100) NOT NULL,
    role_sort INTEGER DEFAULT 0,
    data_scope INTEGER DEFAULT 1,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500)
);

-- 菜单表
CREATE TABLE IF NOT EXISTS sys_menu (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    parent_id BIGINT DEFAULT 0,
    menu_name VARCHAR(50) NOT NULL,
    menu_type VARCHAR(1),
    path VARCHAR(200),
    component VARCHAR(200),
    perms VARCHAR(200),
    icon VARCHAR(100),
    sort INTEGER DEFAULT 0,
    visible INTEGER DEFAULT 1,
    status INTEGER DEFAULT 1,
    is_frame INTEGER DEFAULT 1,
    is_cache INTEGER DEFAULT 0
);

-- 部门表
CREATE TABLE IF NOT EXISTS sys_dept (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    parent_id BIGINT DEFAULT 0,
    dept_name VARCHAR(50) NOT NULL,
    dept_code VARCHAR(50),
    dept_sort INTEGER DEFAULT 0,
    leader VARCHAR(50),
    phone VARCHAR(20),
    email VARCHAR(100),
    status INTEGER DEFAULT 1
);

-- 岗位表
CREATE TABLE IF NOT EXISTS sys_post (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    post_code VARCHAR(50) NOT NULL,
    post_name VARCHAR(100) NOT NULL,
    post_sort INTEGER DEFAULT 0,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500)
);

-- 字典类型表
CREATE TABLE IF NOT EXISTS sys_dict_type (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    dict_name VARCHAR(100) NOT NULL,
    dict_type VARCHAR(100) NOT NULL UNIQUE,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500)
);

-- 字典数据表
CREATE TABLE IF NOT EXISTS sys_dict_data (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    dict_sort INTEGER DEFAULT 0,
    dict_label VARCHAR(100) NOT NULL,
    dict_value VARCHAR(100) NOT NULL,
    dict_type VARCHAR(100) NOT NULL,
    dict_key VARCHAR(100),
    css_class VARCHAR(100),
    list_class VARCHAR(100),
    is_default INTEGER DEFAULT 0,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500)
);

-- 角色菜单关联表
CREATE TABLE IF NOT EXISTS sys_role_menu (
    role_id BIGINT NOT NULL,
    menu_id BIGINT NOT NULL,
    PRIMARY KEY (role_id, menu_id)
);

-- 用户角色关联表
CREATE TABLE IF NOT EXISTS sys_user_role (
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,
    PRIMARY KEY (user_id, role_id)
);

-- 操作日志表
CREATE TABLE IF NOT EXISTS sys_oper_log (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT,
    title VARCHAR(200),
    business_type VARCHAR(20),
    method VARCHAR(100),
    request_method VARCHAR(10),
    operator_type INTEGER DEFAULT 1,
    oper_name VARCHAR(50),
    dept_name VARCHAR(100),
    oper_url VARCHAR(255),
    oper_ip VARCHAR(50),
    oper_location VARCHAR(255),
    oper_param TEXT,
    json_result TEXT,
    status INTEGER DEFAULT 0,
    error_msg TEXT,
    oper_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 登录日志表
CREATE TABLE IF NOT EXISTS sys_login_log (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT,
    username VARCHAR(50),
    ip VARCHAR(50),
    login_location VARCHAR(255),
    browser VARCHAR(50),
    os VARCHAR(50),
    status INTEGER DEFAULT 0,
    msg VARCHAR(255),
    login_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- 主数据模块
-- ============================================

-- 物料表
CREATE TABLE IF NOT EXISTS mdm_material (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    material_code VARCHAR(50) NOT NULL,
    material_name VARCHAR(100) NOT NULL,
    material_type VARCHAR(20),
    spec VARCHAR(100),
    unit VARCHAR(20),
    unit_name VARCHAR(20),
    weight DECIMAL(18,4),
    length DECIMAL(18,4),
    width DECIMAL(18,4),
    height DECIMAL(18,4),
    category_id BIGINT,
    status INTEGER DEFAULT 1,
    UNIQUE(tenant_id, material_code)
);

-- 车间表
CREATE TABLE IF NOT EXISTS mdm_workshop (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    workshop_code VARCHAR(50) NOT NULL,
    workshop_name VARCHAR(100) NOT NULL,
    workshop_type VARCHAR(20),
    manager VARCHAR(50),
    phone VARCHAR(20),
    address VARCHAR(200),
    status INTEGER DEFAULT 1,
    UNIQUE(tenant_id, workshop_code)
);

-- 生产线表
CREATE TABLE IF NOT EXISTS mdm_production_line (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    line_code VARCHAR(50) NOT NULL,
    line_name VARCHAR(100) NOT NULL,
    workshop_id BIGINT NOT NULL,
    line_type VARCHAR(20),
    status INTEGER DEFAULT 1,
    UNIQUE(tenant_id, line_code)
);

-- 工位表
CREATE TABLE IF NOT EXISTS mdm_workstation (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    station_code VARCHAR(50) NOT NULL,
    station_name VARCHAR(100) NOT NULL,
    line_id BIGINT NOT NULL,
    station_type VARCHAR(20),
    status INTEGER DEFAULT 1,
    UNIQUE(tenant_id, station_code)
);

-- 班次表
CREATE TABLE IF NOT EXISTS mdm_shift (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    shift_code VARCHAR(50) NOT NULL,
    shift_name VARCHAR(100) NOT NULL,
    start_time VARCHAR(10),
    end_time VARCHAR(10),
    break_start VARCHAR(10),
    break_end VARCHAR(10),
    status INTEGER DEFAULT 1,
    UNIQUE(tenant_id, shift_code)
);

-- ============================================
-- 生产执行模块
-- ============================================

-- 销售订单表
CREATE TABLE IF NOT EXISTS pro_sales_order (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    order_no VARCHAR(50) NOT NULL,
    customer_id BIGINT,
    customer_name VARCHAR(100),
    order_date DATE,
    delivery_date DATE,
    order_type VARCHAR(20),
    priority INTEGER DEFAULT 1,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500),
    UNIQUE(tenant_id, order_no)
);

-- 生产工单表
CREATE TABLE IF NOT EXISTS pro_production_order (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    order_no VARCHAR(50) NOT NULL,
    sales_order_no VARCHAR(50),
    material_id BIGINT,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    material_spec VARCHAR(100),
    unit VARCHAR(20),
    quantity DECIMAL(18,4),
    completed_qty DECIMAL(18,4) DEFAULT 0,
    rejected_qty DECIMAL(18,4) DEFAULT 0,
    workshop_id BIGINT,
    workshop_name VARCHAR(100),
    line_id BIGINT,
    line_name VARCHAR(100),
    route_id BIGINT,
    bom_id BIGINT,
    plan_start_date DATE,
    plan_end_date DATE,
    actual_start_date DATE,
    actual_end_date DATE,
    priority INTEGER DEFAULT 1,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500),
    UNIQUE(tenant_id, order_no)
);

-- 生产报工表
CREATE TABLE IF NOT EXISTS pro_production_report (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    order_id BIGINT NOT NULL,
    order_no VARCHAR(50),
    process_id BIGINT,
    process_name VARCHAR(100),
    station_id BIGINT,
    station_name VARCHAR(100),
    report_user_id BIGINT,
    report_user_name VARCHAR(50),
    report_date DATE,
    quantity DECIMAL(18,4),
    qualified_qty DECIMAL(18,4),
    rejected_qty DECIMAL(18,4) DEFAULT 0,
    work_time INTEGER,
    remark VARCHAR(500)
);

-- ============================================
-- 设备管理模块
-- ============================================

-- 设备台账表
CREATE TABLE IF NOT EXISTS eqp_equipment (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    equipment_code VARCHAR(50) NOT NULL,
    equipment_name VARCHAR(100) NOT NULL,
    equipment_type VARCHAR(20),
    brand VARCHAR(50),
    model VARCHAR(50),
    serial_number VARCHAR(100),
    workshop_id BIGINT,
    workshop_name VARCHAR(100),
    line_id BIGINT,
    line_name VARCHAR(100),
    station_id BIGINT,
    station_name VARCHAR(100),
    supplier VARCHAR(100),
    purchase_date DATE,
    purchase_price DECIMAL(18,2),
    warranty_end_date DATE,
    status INTEGER DEFAULT 1,
    UNIQUE(tenant_id, equipment_code)
);

-- ============================================
-- 仓储管理模块
-- ============================================

-- 仓库表
CREATE TABLE IF NOT EXISTS wms_warehouse (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    warehouse_code VARCHAR(50) NOT NULL,
    warehouse_name VARCHAR(100) NOT NULL,
    warehouse_type VARCHAR(20),
    address VARCHAR(200),
    manager VARCHAR(50),
    phone VARCHAR(20),
    status INTEGER DEFAULT 1,
    UNIQUE(tenant_id, warehouse_code)
);

-- 库位表
CREATE TABLE IF NOT EXISTS wms_location (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    warehouse_id BIGINT NOT NULL,
    location_code VARCHAR(50) NOT NULL,
    location_name VARCHAR(100),
    location_type VARCHAR(20),
    status INTEGER DEFAULT 1,
    UNIQUE(tenant_id, location_code)
);

-- 库存表
CREATE TABLE IF NOT EXISTS wms_inventory (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    warehouse_id BIGINT NOT NULL,
    location_id BIGINT,
    material_id BIGINT NOT NULL,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    batch_no VARCHAR(50),
    quantity DECIMAL(18,4) DEFAULT 0,
    frozen_qty DECIMAL(18,4) DEFAULT 0,
    available_qty DECIMAL(18,4) DEFAULT 0,
    unit VARCHAR(20)
);

-- ============================================
-- 质量管理模块
-- ============================================

-- IQC检验表
CREATE TABLE IF NOT EXISTS qc_iqc (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    iqc_no VARCHAR(50) NOT NULL,
    material_id BIGINT,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    supplier_id BIGINT,
    supplier_name VARCHAR(100),
    quantity DECIMAL(18,4),
    qualified_qty DECIMAL(18,4),
    unqualified_qty DECIMAL(18,4),
    check_user VARCHAR(50),
    check_date DATE,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500),
    UNIQUE(tenant_id, iqc_no)
);

-- ============================================
-- 追溯模块
-- ============================================

-- 序列号表
CREATE TABLE IF NOT EXISTS trace_serial_number (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    serial_number VARCHAR(100) NOT NULL,
    batch_no VARCHAR(50),
    material_id BIGINT,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    order_id BIGINT,
    order_no VARCHAR(50),
    workshop_id BIGINT,
    line_id BIGINT,
    status INTEGER DEFAULT 1,
    UNIQUE(tenant_id, serial_number)
);

-- 安东呼叫表
CREATE TABLE IF NOT EXISTS andon_call (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    tenant_id BIGINT NOT NULL DEFAULT 1,
    call_no VARCHAR(50) NOT NULL,
    call_type VARCHAR(20),
    equipment_id BIGINT,
    equipment_code VARCHAR(50),
    equipment_name VARCHAR(100),
    station_id BIGINT,
    station_name VARCHAR(100),
    call_user_id BIGINT,
    call_user VARCHAR(50),
    call_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    response_user_id BIGINT,
    response_user VARCHAR(50),
    response_time TIMESTAMP,
    resolve_time TIMESTAMP,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500),
    UNIQUE(tenant_id, call_no)
);

-- ============================================
-- 初始化数据
-- ============================================

-- 初始化租户
INSERT INTO sys_tenant (id, tenant_name, tenant_key, status) VALUES (1, '默认租户', 'default', 1) ON CONFLICT DO NOTHING;

-- 初始化用户 (密码: admin123)
INSERT INTO sys_user (id, tenant_id, username, nickname, password, status) VALUES
(1, 1, 'admin', '管理员', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt6Z5EH', 1),
(2, 1, 'user', '普通用户', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt6Z5EH', 1) ON CONFLICT DO NOTHING;

-- 初始化角色
INSERT INTO sys_role (id, tenant_id, role_name, role_key, role_sort, status) VALUES
(1, 1, '超级管理员', 'admin', 0, 1),
(2, 1, '普通用户', 'user', 1, 1) ON CONFLICT DO NOTHING;

-- 初始化用户角色关联
INSERT INTO sys_user_role (user_id, role_id) VALUES (1, 1), (2, 2) ON CONFLICT DO NOTHING;

-- 初始化菜单
INSERT INTO sys_menu (id, tenant_id, parent_id, menu_name, menu_type, path, component, sort, status) VALUES
(1, 1, 0, '系统管理', 'M', '/system', NULL, 1, 1),
(2, 1, 1, '用户管理', 'C', 'system/user', 'system/UserList', 1, 1),
(3, 1, 1, '角色管理', 'C', 'system/role', 'system/RoleList', 2, 1),
(4, 1, 1, '菜单管理', 'C', 'system/menu', 'system/MenuList', 3, 1),
(5, 1, 1, '部门管理', 'C', 'system/dept', 'system/DeptList', 4, 1),
(6, 1, 1, '字典管理', 'C', 'system/dict', 'system/DictList', 5, 1),
(7, 1, 1, '岗位管理', 'C', 'system/post', 'system/PostList', 6, 1),
(10, 1, 0, '主数据', 'M', '/mdm', NULL, 2, 1),
(11, 1, 10, '物料管理', 'C', 'mdm/material', 'mdm/MaterialList', 1, 1),
(12, 1, 10, '车间管理', 'C', 'mdm/workshop', 'mdm/WorkshopList', 2, 1),
(20, 1, 0, '生产执行', 'M', '/production', NULL, 3, 1),
(21, 1, 20, '生产工单', 'C', 'production/order', 'production/ProductionOrderList', 1, 1),
(30, 1, 0, '设备管理', 'M', '/equipment', NULL, 4, 1),
(31, 1, 30, '设备台账', 'C', 'equipment', 'equipment/EquipmentList', 1, 1),
(40, 1, 0, '仓储管理', 'M', '/wms', NULL, 5, 1),
(41, 1, 40, '仓库管理', 'C', 'wms/warehouse', 'wms/WarehouseList', 1, 1),
(50, 1, 0, '质量管理', 'M', '/quality', NULL, 6, 1),
(51, 1, 50, 'IQC检验', 'C', 'quality/iqc', 'quality/IQCList', 1, 1),
(60, 1, 0, 'APS计划', 'M', '/aps', NULL, 7, 1),
(61, 1, 60, 'MPS计划', 'C', 'aps/mps', 'aps/MPSList', 1, 1),
(70, 1, 0, '追溯管理', 'M', '/trace', NULL, 8, 1),
(71, 1, 70, '追溯查询', 'C', 'trace/query', 'trace/TraceQuery', 1, 1),
(72, 1, 70, '安东呼叫', 'C', 'trace/andon', 'trace/AndonCall', 2, 1),
(80, 1, 0, '能源监控', 'C', '/energy/monitor', 'energy/EnergyMonitor', 9, 1) ON CONFLICT DO NOTHING;

-- 初始化角色菜单关联 (admin拥有所有权限)
INSERT INTO sys_role_menu (role_id, menu_id)
SELECT 1, id FROM sys_menu ON CONFLICT DO NOTHING;

-- 初始化部门
INSERT INTO sys_dept (id, tenant_id, parent_id, dept_name, dept_code, dept_sort, status) VALUES
(1, 1, 0, '总公司', 'HQ', 0, 1),
(2, 1, 1, '研发部', 'RD', 1, 1),
(3, 1, 1, '生产部', 'PROD', 2, 1),
(4, 1, 1, '质量部', 'QC', 3, 1),
(5, 1, 1, '仓储部', 'WH', 4, 1) ON CONFLICT DO NOTHING;

-- 初始化岗位
INSERT INTO sys_post (id, tenant_id, post_code, post_name, post_sort, status) VALUES
(1, 1, 'CEO', '总经理', 0, 1),
(2, 1, 'MGR', '部门经理', 1, 1),
(3, 1, 'LEADER', '组长', 2, 1),
(4, 1, 'STAFF', '员工', 3, 1) ON CONFLICT DO NOTHING;

-- 初始化字典类型
INSERT INTO sys_dict_type (id, dict_name, dict_type, status) VALUES
(1, '用户状态', 'sys_user_status', 1),
(2, '角色状态', 'sys_role_status', 1),
(3, '菜单类型', 'sys_menu_type', 1),
(4, '物料类型', 'material_type', 1),
(5, '仓库类型', 'warehouse_type', 1),
(6, '设备状态', 'equipment_status', 1),
(7, '生产工单状态', 'production_status', 1) ON CONFLICT DO NOTHING;

-- 初始化字典数据
INSERT INTO sys_dict_data (dict_sort, dict_label, dict_value, dict_type, status) VALUES
(1, '正常', '1', 'sys_user_status', 1),
(2, '禁用', '0', 'sys_user_status', 1),
(1, '正常', '1', 'sys_role_status', 1),
(2, '禁用', '0', 'sys_role_status', 1),
(1, '目录', 'M', 'sys_menu_type', 1),
(2, '菜单', 'C', 'sys_menu_type', 1),
(3, '按钮', 'F', 'sys_menu_type', 1),
(1, '原材料', 'raw', 'material_type', 1),
(2, '半成品', 'semi', 'material_type', 1),
(3, '成品', 'finished', 'material_type', 1),
(1, '原料仓', 'raw', 'warehouse_type', 1),
(2, '成品仓', 'finished', 'warehouse_type', 1),
(3, '半成品仓', 'semi', 'warehouse_type', 1),
(1, '运行', '1', 'equipment_status', 1),
(2, '待机', '2', 'equipment_status', 1),
(3, '故障', '3', 'equipment_status', 1),
(4, '维修', '4', 'equipment_status', 1),
(5, '报废', '5', 'equipment_status', 1),
(1, '待生产', '1', 'production_status', 1),
(2, '生产中', '2', 'production_status', 1),
(3, '已完成', '3', 'production_status', 1),
(4, '已取消', '4', 'production_status', 1) ON CONFLICT DO NOTHING;

-- 初始化示例数据
INSERT INTO mdm_material (tenant_id, material_code, material_name, material_type, spec, unit, status) VALUES
(1, 'MAT-001', '钢材A', 'raw', '10mm*100mm', 'kg', 1),
(1, 'MAT-002', '塑料粒子B', 'raw', '5mm颗粒', 'kg', 1),
(1, 'PRD-001', '产品X', 'finished', '标准规格', '个', 1) ON CONFLICT DO NOTHING;

INSERT INTO mdm_workshop (tenant_id, workshop_code, workshop_name, workshop_type, manager, status) VALUES
(1, 'WS-001', '一车间', '加工', '张三', 1),
(1, 'WS-002', '二车间', '装配', '李四', 1) ON CONFLICT DO NOTHING;

INSERT INTO wms_warehouse (tenant_id, warehouse_code, warehouse_name, warehouse_type, manager, status) VALUES
(1, 'WH-001', '原料仓', 'raw', '王五', 1),
(1, 'WH-002', '成品仓', 'finished', '赵六', 1) ON CONFLICT DO NOTHING;

-- 更新序列
SELECT setval('sys_tenant_id_seq', (SELECT MAX(id) FROM sys_tenant));
SELECT setval('sys_user_id_seq', (SELECT MAX(id) FROM sys_user));
SELECT setval('sys_role_id_seq', (SELECT MAX(id) FROM sys_role));
SELECT setval('sys_menu_id_seq', (SELECT MAX(id) FROM sys_menu));
SELECT setval('sys_dept_id_seq', (SELECT MAX(id) FROM sys_dept));
SELECT setval('sys_post_id_seq', (SELECT MAX(id) FROM sys_post));
SELECT setval('mdm_material_id_seq', (SELECT MAX(id) FROM mdm_material));
SELECT setval('mdm_workshop_id_seq', (SELECT MAX(id) FROM mdm_workshop));
SELECT setval('wms_warehouse_id_seq', (SELECT MAX(id) FROM wms_warehouse));

-- 创建索引
CREATE INDEX IF NOT EXISTS idx_sys_user_tenant ON sys_user(tenant_id);
CREATE INDEX IF NOT EXISTS idx_sys_role_tenant ON sys_role(tenant_id);
CREATE INDEX IF NOT EXISTS idx_sys_menu_tenant ON sys_menu(tenant_id);
CREATE INDEX IF NOT EXISTS idx_sys_dept_tenant ON sys_dept(tenant_id);
CREATE INDEX IF NOT EXISTS idx_mdm_material_tenant ON mdm_material(tenant_id);
CREATE INDEX IF NOT EXISTS idx_pro_production_order_tenant ON pro_production_order(tenant_id);
CREATE INDEX IF NOT EXISTS idx_eqp_equipment_tenant ON eqp_equipment(tenant_id);
CREATE INDEX IF NOT EXISTS idx_wms_warehouse_tenant ON wms_warehouse(tenant_id);

-- 提交
COMMIT;
