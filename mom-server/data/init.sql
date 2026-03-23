-- =====================================================
-- MOM3.0 数据库初始化脚本
-- PostgreSQL 16+
-- =====================================================

-- 创建数据库(如有需要)
-- CREATE DATABASE mom3 OWNER postgres;

-- =====================================================
-- 系统管理模块
-- =====================================================

-- 租户表
CREATE TABLE IF NOT EXISTS sys_tenant (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_name VARCHAR(100) NOT NULL,
    tenant_key VARCHAR(100) NOT NULL UNIQUE,
    contact VARCHAR(50),
    phone VARCHAR(20),
    email VARCHAR(100),
    status INTEGER DEFAULT 1,
    expire_time TIMESTAMP,
    package_id BIGINT
);
CREATE INDEX IF NOT EXISTS idx_sys_tenant_key ON sys_tenant(tenant_key);

-- 用户表
CREATE TABLE IF NOT EXISTS sys_user (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    username VARCHAR(50) NOT NULL,
    nickname VARCHAR(50),
    password VARCHAR(200) NOT NULL,
    email VARCHAR(100),
    phone VARCHAR(20),
    avatar VARCHAR(500),
    dept_id BIGINT,
    status INTEGER DEFAULT 1,
    login_ip VARCHAR(128),
    login_date TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_sys_user_tenant ON sys_user(tenant_id);
CREATE INDEX IF NOT EXISTS idx_sys_user_dept ON sys_user(dept_id);

-- 角色表
CREATE TABLE IF NOT EXISTS sys_role (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    role_name VARCHAR(50) NOT NULL,
    role_key VARCHAR(100) NOT NULL,
    role_sort INTEGER DEFAULT 0,
    data_scope INTEGER DEFAULT 1,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500)
);
CREATE INDEX IF NOT EXISTS idx_sys_role_tenant ON sys_role(tenant_id);

-- 菜单表
CREATE TABLE IF NOT EXISTS sys_menu (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
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
CREATE INDEX IF NOT EXISTS idx_sys_menu_parent ON sys_menu(parent_id);
CREATE INDEX IF NOT EXISTS idx_sys_menu_tenant ON sys_menu(tenant_id);

-- 部门表
CREATE TABLE IF NOT EXISTS sys_dept (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    parent_id BIGINT DEFAULT 0,
    dept_name VARCHAR(50) NOT NULL,
    dept_code VARCHAR(50),
    dept_sort INTEGER DEFAULT 0,
    leader VARCHAR(50),
    phone VARCHAR(20),
    email VARCHAR(100),
    status INTEGER DEFAULT 1
);
CREATE INDEX IF NOT EXISTS idx_sys_dept_parent ON sys_dept(parent_id);
CREATE INDEX IF NOT EXISTS idx_sys_dept_tenant ON sys_dept(tenant_id);

-- 岗位表
CREATE TABLE IF NOT EXISTS sys_post (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    post_code VARCHAR(50) NOT NULL,
    post_name VARCHAR(100) NOT NULL,
    post_sort INTEGER DEFAULT 0,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500)
);
CREATE INDEX IF NOT EXISTS idx_sys_post_tenant ON sys_post(tenant_id);

-- 字典类型表
CREATE TABLE IF NOT EXISTS sys_dict_type (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    dict_name VARCHAR(100) NOT NULL,
    dict_type VARCHAR(100) NOT NULL UNIQUE,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500)
);

-- 字典数据表
CREATE TABLE IF NOT EXISTS sys_dict_data (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
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
CREATE INDEX IF NOT EXISTS idx_sys_dict_data_type ON sys_dict_data(dict_type);

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
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
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
    oper_time TIMESTAMP NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_sys_oper_log_tenant ON sys_oper_log(tenant_id);

-- 登录日志表
CREATE TABLE IF NOT EXISTS sys_login_log (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT,
    username VARCHAR(50),
    ip VARCHAR(50),
    login_location VARCHAR(255),
    browser VARCHAR(50),
    os VARCHAR(50),
    status INTEGER DEFAULT 0,
    msg VARCHAR(255),
    login_time TIMESTAMP NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_sys_login_log_tenant ON sys_login_log(tenant_id);

-- =====================================================
-- 主数据模块
-- =====================================================

-- 物料表
CREATE TABLE IF NOT EXISTS mdm_material (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
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
    UNIQUE (tenant_id, material_code)
);
CREATE INDEX IF NOT EXISTS idx_mdm_material_tenant ON mdm_material(tenant_id);
CREATE INDEX IF NOT EXISTS idx_mdm_material_code ON mdm_material(material_code);

-- 物料分类表
CREATE TABLE IF NOT EXISTS mdm_material_category (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    parent_id BIGINT DEFAULT 0,
    category_name VARCHAR(100) NOT NULL,
    category_code VARCHAR(50),
    sort INTEGER DEFAULT 0,
    status INTEGER DEFAULT 1
);
CREATE INDEX IF NOT EXISTS idx_mdm_material_category_tenant ON mdm_material_category(tenant_id);

-- 车间表
CREATE TABLE IF NOT EXISTS mdm_workshop (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    workshop_code VARCHAR(50) NOT NULL,
    workshop_name VARCHAR(100) NOT NULL,
    workshop_type VARCHAR(20),
    manager VARCHAR(50),
    phone VARCHAR(20),
    address VARCHAR(200),
    status INTEGER DEFAULT 1,
    UNIQUE (tenant_id, workshop_code)
);
CREATE INDEX IF NOT EXISTS idx_mdm_workshop_tenant ON mdm_workshop(tenant_id);

-- 生产线表
CREATE TABLE IF NOT EXISTS mdm_production_line (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    line_code VARCHAR(50) NOT NULL,
    line_name VARCHAR(100) NOT NULL,
    workshop_id BIGINT NOT NULL,
    line_type VARCHAR(20),
    status INTEGER DEFAULT 1,
    UNIQUE (tenant_id, line_code)
);
CREATE INDEX IF NOT EXISTS idx_mdm_production_line_tenant ON mdm_production_line(tenant_id);
CREATE INDEX IF NOT EXISTS idx_mdm_production_line_workshop ON mdm_production_line(workshop_id);

-- 工位表
CREATE TABLE IF NOT EXISTS mdm_workstation (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    station_code VARCHAR(50) NOT NULL,
    station_name VARCHAR(100) NOT NULL,
    line_id BIGINT NOT NULL,
    station_type VARCHAR(20),
    status INTEGER DEFAULT 1,
    UNIQUE (tenant_id, station_code)
);
CREATE INDEX IF NOT EXISTS idx_mdm_workstation_tenant ON mdm_workstation(tenant_id);
CREATE INDEX IF NOT EXISTS idx_mdm_workstation_line ON mdm_workstation(line_id);

-- 班次表
CREATE TABLE IF NOT EXISTS mdm_shift (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    shift_code VARCHAR(50) NOT NULL,
    shift_name VARCHAR(100) NOT NULL,
    start_time VARCHAR(10),
    end_time VARCHAR(10),
    break_start VARCHAR(10),
    break_end VARCHAR(10),
    status INTEGER DEFAULT 1,
    UNIQUE (tenant_id, shift_code)
);
CREATE INDEX IF NOT EXISTS idx_mdm_shift_tenant ON mdm_shift(tenant_id);

-- BOM表
CREATE TABLE IF NOT EXISTS mdm_bom (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    bom_code VARCHAR(50) NOT NULL,
    bom_name VARCHAR(200) NOT NULL,
    material_id BIGINT NOT NULL,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    version VARCHAR(20),
    status VARCHAR(20) DEFAULT 'DRAFT',
    eff_date DATE,
    exp_date DATE,
    remark VARCHAR(500),
    UNIQUE (tenant_id, bom_code)
);
CREATE INDEX IF NOT EXISTS idx_mdm_bom_tenant ON mdm_bom(tenant_id);
CREATE INDEX IF NOT EXISTS idx_mdm_bom_code ON mdm_bom(bom_code);

-- BOM明细表
CREATE TABLE IF NOT EXISTS mdm_bom_item (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    bom_id BIGINT NOT NULL,
    line_no INTEGER DEFAULT 0,
    material_id BIGINT NOT NULL,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    quantity DECIMAL(18,4) DEFAULT 0,
    unit VARCHAR(20),
    scrap_rate DECIMAL(10,4) DEFAULT 0,
    substitute_group VARCHAR(50),
    is_alternative INTEGER DEFAULT 0
);
CREATE INDEX IF NOT EXISTS idx_mdm_bom_item_tenant ON mdm_bom_item(tenant_id);
CREATE INDEX IF NOT EXISTS idx_mdm_bom_item_bom ON mdm_bom_item(bom_id);

-- 工序表
CREATE TABLE IF NOT EXISTS mdm_operation (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    operation_code VARCHAR(50) NOT NULL,
    operation_name VARCHAR(100) NOT NULL,
    workcenter_id BIGINT,
    workcenter_name VARCHAR(100),
    standard_worktime INTEGER DEFAULT 0,
    quality_std VARCHAR(500),
    is_key_process INTEGER DEFAULT 0,
    is_qc_point INTEGER DEFAULT 0,
    sequence INTEGER DEFAULT 0,
    remark VARCHAR(500),
    UNIQUE (tenant_id, operation_code)
);
CREATE INDEX IF NOT EXISTS idx_mdm_operation_tenant ON mdm_operation(tenant_id);

-- =====================================================
-- 设备管理模块
-- =====================================================

-- 设备台账表
CREATE TABLE IF NOT EXISTS equ_equipment (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
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
    purchase_date TIMESTAMP,
    purchase_price DECIMAL(18,2),
    warranty_end_date TIMESTAMP,
    status INTEGER DEFAULT 1,
    UNIQUE (tenant_id, equipment_code)
);
CREATE INDEX IF NOT EXISTS idx_equ_equipment_tenant ON equ_equipment(tenant_id);
CREATE INDEX IF NOT EXISTS idx_equ_equipment_code ON equ_equipment(equipment_code);

-- 设备点检表
CREATE TABLE IF NOT EXISTS equ_equipment_check (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    equipment_id BIGINT,
    equipment_code VARCHAR(50),
    equipment_name VARCHAR(100),
    check_plan_id BIGINT,
    check_user_id BIGINT,
    check_user_name VARCHAR(50),
    check_date TIMESTAMP,
    check_result INTEGER,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500)
);
CREATE INDEX IF NOT EXISTS idx_equ_equipment_check_tenant ON equ_equipment_check(tenant_id);
CREATE INDEX IF NOT EXISTS idx_equ_equipment_check_equipment ON equ_equipment_check(equipment_id);

-- 设备保养表
CREATE TABLE IF NOT EXISTS equ_equipment_maintenance (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    equipment_id BIGINT,
    equipment_code VARCHAR(50),
    equipment_name VARCHAR(100),
    maint_type VARCHAR(20),
    maint_plan_id BIGINT,
    maint_user_id BIGINT,
    maint_user_name VARCHAR(50),
    maint_date TIMESTAMP,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    duration INTEGER,
    content TEXT,
    cost DECIMAL(18,2),
    status INTEGER DEFAULT 1
);
CREATE INDEX IF NOT EXISTS idx_equ_equipment_maintenance_tenant ON equ_equipment_maintenance(tenant_id);
CREATE INDEX IF NOT EXISTS idx_equ_equipment_maintenance_equipment ON equ_equipment_maintenance(equipment_id);

-- 设备维修表
CREATE TABLE IF NOT EXISTS equ_equipment_repair (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    equipment_id BIGINT,
    equipment_code VARCHAR(50),
    equipment_name VARCHAR(100),
    fault_desc TEXT,
    fault_time TIMESTAMP,
    report_user_id BIGINT,
    report_user_name VARCHAR(50),
    repair_user_id BIGINT,
    repair_user_name VARCHAR(50),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    duration INTEGER,
    repair_content TEXT,
    cost DECIMAL(18,2),
    status INTEGER DEFAULT 1
);
CREATE INDEX IF NOT EXISTS idx_equ_equipment_repair_tenant ON equ_equipment_repair(tenant_id);
CREATE INDEX IF NOT EXISTS idx_equ_equipment_repair_equipment ON equ_equipment_repair(equipment_id);

-- 备件表
CREATE TABLE IF NOT EXISTS equ_spare_part (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    spare_part_code VARCHAR(50) NOT NULL,
    spare_part_name VARCHAR(100) NOT NULL,
    spec VARCHAR(100),
    unit VARCHAR(20),
    quantity DECIMAL(18,2) DEFAULT 0,
    min_quantity DECIMAL(18,2),
    price DECIMAL(18,2),
    supplier VARCHAR(100),
    status INTEGER DEFAULT 1,
    UNIQUE (tenant_id, spare_part_code)
);
CREATE INDEX IF NOT EXISTS idx_equ_spare_part_tenant ON equ_spare_part(tenant_id);

-- =====================================================
-- 生产执行模块
-- =====================================================

-- 销售订单表
CREATE TABLE IF NOT EXISTS pro_sales_order (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    order_no VARCHAR(50) NOT NULL,
    customer_id BIGINT,
    customer_name VARCHAR(100),
    order_date TIMESTAMP,
    delivery_date TIMESTAMP,
    order_type VARCHAR(20),
    priority INTEGER DEFAULT 1,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500),
    UNIQUE (tenant_id, order_no)
);
CREATE INDEX IF NOT EXISTS idx_pro_sales_order_tenant ON pro_sales_order(tenant_id);
CREATE INDEX IF NOT EXISTS idx_pro_sales_order_no ON pro_sales_order(order_no);

-- 销售订单明细表
CREATE TABLE IF NOT EXISTS pro_sales_order_item (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    order_id BIGINT NOT NULL,
    material_id BIGINT,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    quantity DECIMAL(18,4),
    unit VARCHAR(20),
    price DECIMAL(18,2),
    amount DECIMAL(18,2),
    shipped_qty DECIMAL(18,4) DEFAULT 0
);
CREATE INDEX IF NOT EXISTS idx_pro_sales_order_item_order ON pro_sales_order_item(order_id);

-- 生产工单表
CREATE TABLE IF NOT EXISTS pro_production_order (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
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
    plan_start_date TIMESTAMP,
    plan_end_date TIMESTAMP,
    actual_start_date TIMESTAMP,
    actual_end_date TIMESTAMP,
    priority INTEGER DEFAULT 1,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500),
    UNIQUE (tenant_id, order_no)
);
CREATE INDEX IF NOT EXISTS idx_pro_production_order_tenant ON pro_production_order(tenant_id);
CREATE INDEX IF NOT EXISTS idx_pro_production_order_no ON pro_production_order(order_no);

-- 生产报工表
CREATE TABLE IF NOT EXISTS pro_production_report (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    order_id BIGINT NOT NULL,
    order_no VARCHAR(50),
    process_id BIGINT,
    process_name VARCHAR(100),
    station_id BIGINT,
    station_name VARCHAR(100),
    report_user_id BIGINT,
    report_user_name VARCHAR(50),
    report_date TIMESTAMP,
    quantity DECIMAL(18,4),
    qualified_qty DECIMAL(18,4),
    rejected_qty DECIMAL(18,4) DEFAULT 0,
    work_time INTEGER,
    remark VARCHAR(500)
);
CREATE INDEX IF NOT EXISTS idx_pro_production_report_tenant ON pro_production_report(tenant_id);
CREATE INDEX IF NOT EXISTS idx_pro_production_report_order ON pro_production_report(order_id);

-- 派工表
CREATE TABLE IF NOT EXISTS pro_dispatch (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    order_id BIGINT NOT NULL,
    order_no VARCHAR(50),
    process_id BIGINT,
    process_name VARCHAR(100),
    station_id BIGINT,
    station_name VARCHAR(100),
    assign_user_id BIGINT,
    assign_user_name VARCHAR(50),
    quantity DECIMAL(18,4),
    status INTEGER DEFAULT 1
);
CREATE INDEX IF NOT EXISTS idx_pro_dispatch_tenant ON pro_dispatch(tenant_id);
CREATE INDEX IF NOT EXISTS idx_pro_dispatch_order ON pro_dispatch(order_id);

-- =====================================================
-- 仓储管理模块
-- =====================================================

-- 仓库表
CREATE TABLE IF NOT EXISTS wms_warehouse (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    warehouse_code VARCHAR(50) NOT NULL,
    warehouse_name VARCHAR(100) NOT NULL,
    warehouse_type VARCHAR(20),
    address VARCHAR(200),
    manager VARCHAR(50),
    phone VARCHAR(20),
    status INTEGER DEFAULT 1,
    UNIQUE (tenant_id, warehouse_code)
);
CREATE INDEX IF NOT EXISTS idx_wms_warehouse_tenant ON wms_warehouse(tenant_id);

-- 库位表
CREATE TABLE IF NOT EXISTS wms_location (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    location_code VARCHAR(50) NOT NULL,
    location_name VARCHAR(100),
    warehouse_id BIGINT,
    zone_code VARCHAR(20),
    row INTEGER,
    col INTEGER,
    layer INTEGER,
    location_type VARCHAR(20),
    capacity INTEGER,
    status INTEGER DEFAULT 1,
    UNIQUE (tenant_id, location_code)
);
CREATE INDEX IF NOT EXISTS idx_wms_location_tenant ON wms_location(tenant_id);
CREATE INDEX IF NOT EXISTS idx_wms_location_warehouse ON wms_location(warehouse_id);

-- 库存表
CREATE TABLE IF NOT EXISTS wms_inventory (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    material_id BIGINT,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    warehouse_id BIGINT,
    location_id BIGINT,
    quantity DECIMAL(18,4) DEFAULT 0,
    available_qty DECIMAL(18,4) DEFAULT 0,
    allocated_qty DECIMAL(18,4) DEFAULT 0,
    locked_qty DECIMAL(18,4) DEFAULT 0,
    batch_no VARCHAR(50)
);
CREATE INDEX IF NOT EXISTS idx_wms_inventory_tenant ON wms_inventory(tenant_id);
CREATE INDEX IF NOT EXISTS idx_wms_inventory_material ON wms_inventory(material_id);
CREATE INDEX IF NOT EXISTS idx_wms_inventory_warehouse ON wms_inventory(warehouse_id);
CREATE INDEX IF NOT EXISTS idx_wms_inventory_location ON wms_inventory(location_id);

-- 库存记录表
CREATE TABLE IF NOT EXISTS wms_inventory_record (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    record_no VARCHAR(50) NOT NULL,
    record_type VARCHAR(20),
    material_id BIGINT,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    warehouse_id BIGINT,
    location_id BIGINT,
    quantity DECIMAL(18,4),
    batch_no VARCHAR(50),
    source_no VARCHAR(50),
    operator_id BIGINT,
    operator_name VARCHAR(50),
    operate_time TIMESTAMP,
    remark VARCHAR(500),
    UNIQUE (tenant_id, record_no)
);
CREATE INDEX IF NOT EXISTS idx_wms_inventory_record_tenant ON wms_inventory_record(tenant_id);

-- 收货单表
CREATE TABLE IF NOT EXISTS wms_receive_order (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    receive_no VARCHAR(50) NOT NULL,
    supplier_id BIGINT,
    supplier_name VARCHAR(100),
    warehouse_id BIGINT,
    receive_date TIMESTAMP,
    receive_user_id BIGINT,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500),
    UNIQUE (tenant_id, receive_no)
);
CREATE INDEX IF NOT EXISTS idx_wms_receive_order_tenant ON wms_receive_order(tenant_id);

-- 收货单明细表
CREATE TABLE IF NOT EXISTS wms_receive_order_item (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    receive_id BIGINT,
    material_id BIGINT,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    quantity DECIMAL(18,4),
    received_qty DECIMAL(18,4) DEFAULT 0,
    unit VARCHAR(20),
    batch_no VARCHAR(50)
);
CREATE INDEX IF NOT EXISTS idx_wms_receive_order_item_receive ON wms_receive_order_item(receive_id);

-- 发货单表
CREATE TABLE IF NOT EXISTS wms_delivery_order (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    delivery_no VARCHAR(50) NOT NULL,
    customer_id BIGINT,
    customer_name VARCHAR(100),
    warehouse_id BIGINT,
    delivery_date TIMESTAMP,
    delivery_user_id BIGINT,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500),
    UNIQUE (tenant_id, delivery_no)
);
CREATE INDEX IF NOT EXISTS idx_wms_delivery_order_tenant ON wms_delivery_order(tenant_id);

-- 发货单明细表
CREATE TABLE IF NOT EXISTS wms_delivery_order_item (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    delivery_id BIGINT,
    material_id BIGINT,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    quantity DECIMAL(18,4),
    shipped_qty DECIMAL(18,4) DEFAULT 0,
    unit VARCHAR(20),
    batch_no VARCHAR(50)
);
CREATE INDEX IF NOT EXISTS idx_wms_delivery_order_item_delivery ON wms_delivery_order_item(delivery_id);

-- 调拨单表
CREATE TABLE IF NOT EXISTS wms_transfer_order (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    transfer_no VARCHAR(50) NOT NULL,
    from_warehouse_id BIGINT,
    to_warehouse_id BIGINT,
    transfer_date TIMESTAMP,
    transfer_user_id BIGINT,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500),
    UNIQUE (tenant_id, transfer_no)
);
CREATE INDEX IF NOT EXISTS idx_wms_transfer_order_tenant ON wms_transfer_order(tenant_id);

-- 盘点单表
CREATE TABLE IF NOT EXISTS wms_stock_check (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    check_no VARCHAR(50) NOT NULL,
    warehouse_id BIGINT,
    check_date TIMESTAMP,
    check_user_id BIGINT,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500),
    UNIQUE (tenant_id, check_no)
);
CREATE INDEX IF NOT EXISTS idx_wms_stock_check_tenant ON wms_stock_check(tenant_id);

-- =====================================================
-- APS模块
-- =====================================================

-- MRP表
CREATE TABLE IF NOT EXISTS aps_mrp (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    mrp_no VARCHAR(50) NOT NULL,
    mrp_type VARCHAR(20),
    plan_date TIMESTAMP,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500),
    UNIQUE (tenant_id, mrp_no)
);
CREATE INDEX IF NOT EXISTS idx_aps_mrp_tenant ON aps_mrp(tenant_id);

-- MRP明细表
CREATE TABLE IF NOT EXISTS aps_mrp_item (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    mrp_id BIGINT,
    material_id BIGINT,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    quantity DECIMAL(18,4),
    stock_qty DECIMAL(18,4),
    allocated_qty DECIMAL(18,4),
    net_qty DECIMAL(18,4),
    source_type VARCHAR(20),
    source_no VARCHAR(50)
);
CREATE INDEX IF NOT EXISTS idx_aps_mrp_item_mrp ON aps_mrp_item(mrp_id);

-- MPS表
CREATE TABLE IF NOT EXISTS aps_mps (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    mps_no VARCHAR(50) NOT NULL,
    plan_month VARCHAR(10),
    material_id BIGINT,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    quantity DECIMAL(18,4),
    status INTEGER DEFAULT 1,
    UNIQUE (tenant_id, mps_no)
);
CREATE INDEX IF NOT EXISTS idx_aps_mps_tenant ON aps_mps(tenant_id);

-- 排程计划表
CREATE TABLE IF NOT EXISTS aps_schedule_plan (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    plan_no VARCHAR(50) NOT NULL,
    plan_type VARCHAR(20),
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    algorithm VARCHAR(20),
    status INTEGER DEFAULT 1,
    remark VARCHAR(500),
    UNIQUE (tenant_id, plan_no)
);
CREATE INDEX IF NOT EXISTS idx_aps_schedule_plan_tenant ON aps_schedule_plan(tenant_id);

-- 排程结果表
CREATE TABLE IF NOT EXISTS aps_schedule_result (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    plan_id BIGINT,
    order_id BIGINT,
    order_no VARCHAR(50),
    sequence INTEGER,
    line_id BIGINT,
    line_name VARCHAR(100),
    station_id BIGINT,
    station_name VARCHAR(100),
    plan_start_time TIMESTAMP,
    plan_end_time TIMESTAMP,
    actual_start_time TIMESTAMP,
    actual_end_time TIMESTAMP,
    status INTEGER DEFAULT 1
);
CREATE INDEX IF NOT EXISTS idx_aps_schedule_result_plan ON aps_schedule_result(plan_id);
CREATE INDEX IF NOT EXISTS idx_aps_schedule_result_order ON aps_schedule_result(order_id);

-- 资源表
CREATE TABLE IF NOT EXISTS aps_resource (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    resource_code VARCHAR(50) NOT NULL,
    resource_name VARCHAR(100) NOT NULL,
    resource_type VARCHAR(20),
    workshop_id BIGINT,
    capacity DECIMAL(18,2),
    unit VARCHAR(20),
    efficiency DECIMAL(10,2) DEFAULT 100,
    status INTEGER DEFAULT 1,
    UNIQUE (tenant_id, resource_code)
);
CREATE INDEX IF NOT EXISTS idx_aps_resource_tenant ON aps_resource(tenant_id);

-- 工作中心表
CREATE TABLE IF NOT EXISTS aps_work_center (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    work_center_code VARCHAR(50) NOT NULL,
    work_center_name VARCHAR(100) NOT NULL,
    workshop_id BIGINT,
    capacity DECIMAL(18,2),
    status INTEGER DEFAULT 1,
    UNIQUE (tenant_id, work_center_code)
);
CREATE INDEX IF NOT EXISTS idx_aps_work_center_tenant ON aps_work_center(tenant_id);

-- =====================================================
-- 质量管理模块
-- =====================================================

-- IQC检验表
CREATE TABLE IF NOT EXISTS qc_iqc (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    iqc_no VARCHAR(50) NOT NULL,
    supplier_id BIGINT,
    supplier_name VARCHAR(100),
    material_id BIGINT,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    quantity DECIMAL(18,4),
    unit VARCHAR(20),
    check_user_id BIGINT,
    check_user_name VARCHAR(50),
    check_date TIMESTAMP,
    result INTEGER DEFAULT 1,
    remark VARCHAR(500),
    UNIQUE (tenant_id, iqc_no)
);
CREATE INDEX IF NOT EXISTS idx_qc_iqc_tenant ON qc_iqc(tenant_id);

-- IQC检验明细表
CREATE TABLE IF NOT EXISTS qc_iqc_item (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    iqc_id BIGINT,
    check_item_id BIGINT,
    check_item VARCHAR(100),
    check_standard VARCHAR(200),
    check_method VARCHAR(100),
    result INTEGER DEFAULT 1,
    remark VARCHAR(200)
);
CREATE INDEX IF NOT EXISTS idx_qc_iqc_item_iqc ON qc_iqc_item(iqc_id);

-- IPQC过程检验表
CREATE TABLE IF NOT EXISTS qc_ipqc (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    ipqc_no VARCHAR(50) NOT NULL,
    order_id BIGINT,
    order_no VARCHAR(50),
    process_id BIGINT,
    process_name VARCHAR(100),
    quantity DECIMAL(18,4),
    sample_size INTEGER,
    check_user_id BIGINT,
    check_user_name VARCHAR(50),
    check_date TIMESTAMP,
    result INTEGER DEFAULT 1,
    remark VARCHAR(500),
    UNIQUE (tenant_id, ipqc_no)
);
CREATE INDEX IF NOT EXISTS idx_qc_ipqc_tenant ON qc_ipqc(tenant_id);

-- FQC最终检验表
CREATE TABLE IF NOT EXISTS qc_fqc (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    fqc_no VARCHAR(50) NOT NULL,
    order_id BIGINT,
    order_no VARCHAR(50),
    quantity DECIMAL(18,4),
    sample_size INTEGER,
    qualified_qty DECIMAL(18,4),
    rejected_qty DECIMAL(18,4) DEFAULT 0,
    check_user_id BIGINT,
    check_user_name VARCHAR(50),
    check_date TIMESTAMP,
    result INTEGER DEFAULT 1,
    remark VARCHAR(500),
    UNIQUE (tenant_id, fqc_no)
);
CREATE INDEX IF NOT EXISTS idx_qc_fqc_tenant ON qc_fqc(tenant_id);

-- OQC出货检验表
CREATE TABLE IF NOT EXISTS qc_oqc (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    oqc_no VARCHAR(50) NOT NULL,
    shipping_no VARCHAR(50),
    customer_id BIGINT,
    customer_name VARCHAR(100),
    quantity DECIMAL(18,4),
    check_user_id BIGINT,
    check_user_name VARCHAR(50),
    check_date TIMESTAMP,
    result INTEGER DEFAULT 1,
    remark VARCHAR(500),
    UNIQUE (tenant_id, oqc_no)
);
CREATE INDEX IF NOT EXISTS idx_qc_oqc_tenant ON qc_oqc(tenant_id);

-- 不良品代码表
CREATE TABLE IF NOT EXISTS qc_defect_code (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    defect_code VARCHAR(20) NOT NULL,
    defect_name VARCHAR(100) NOT NULL,
    defect_type VARCHAR(20),
    severity INTEGER DEFAULT 1,
    status INTEGER DEFAULT 1,
    UNIQUE (tenant_id, defect_code)
);
CREATE INDEX IF NOT EXISTS idx_qc_defect_code_tenant ON qc_defect_code(tenant_id);

-- 不良品记录表
CREATE TABLE IF NOT EXISTS qc_defect_record (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    record_no VARCHAR(50) NOT NULL,
    order_id BIGINT,
    order_no VARCHAR(50),
    process_id BIGINT,
    process_name VARCHAR(100),
    defect_code_id BIGINT,
    defect_code VARCHAR(20),
    defect_name VARCHAR(100),
    quantity DECIMAL(18,4),
    handle_method INTEGER DEFAULT 1,
    handle_user_id BIGINT,
    handle_date TIMESTAMP,
    status INTEGER DEFAULT 1,
    UNIQUE (tenant_id, record_no)
);
CREATE INDEX IF NOT EXISTS idx_qc_defect_record_tenant ON qc_defect_record(tenant_id);

-- NCR不良品处理表
CREATE TABLE IF NOT EXISTS qc_ncr (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    ncr_no VARCHAR(50) NOT NULL,
    defect_id BIGINT,
    source_type VARCHAR(20),
    issue_desc TEXT,
    root_cause TEXT,
    corrective_action TEXT,
    preventive_action TEXT,
    verify_result VARCHAR(200),
    verify_user_id BIGINT,
    verify_date TIMESTAMP,
    status INTEGER DEFAULT 1,
    UNIQUE (tenant_id, ncr_no)
);
CREATE INDEX IF NOT EXISTS idx_qc_ncr_tenant ON qc_ncr(tenant_id);

-- SPC数据表
CREATE TABLE IF NOT EXISTS qc_spc_data (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    equipment_id BIGINT,
    station_id BIGINT,
    process_id BIGINT,
    process_name VARCHAR(100),
    check_item VARCHAR(100),
    check_value DECIMAL(18,4),
    usl DECIMAL(18,4),
    lsl DECIMAL(18,4),
    cl DECIMAL(18,4),
    ucl DECIMAL(18,4),
    lcl DECIMAL(18,4),
    check_time TIMESTAMP NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_qc_spc_data_tenant ON qc_spc_data(tenant_id);

-- =====================================================
-- 追溯模块
-- =====================================================

-- 序列号表
CREATE TABLE IF NOT EXISTS tra_serial_number (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    serial_number VARCHAR(50) NOT NULL,
    material_id BIGINT,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    batch_no VARCHAR(50),
    line_id BIGINT,
    line_name VARCHAR(100),
    order_id BIGINT,
    order_no VARCHAR(50),
    production_date TIMESTAMP,
    status INTEGER DEFAULT 1,
    UNIQUE (tenant_id, serial_number)
);
CREATE INDEX IF NOT EXISTS idx_tra_serial_number_tenant ON tra_serial_number(tenant_id);

-- 追溯记录表
CREATE TABLE IF NOT EXISTS tra_trace_record (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    trace_no VARCHAR(50) NOT NULL,
    material_id BIGINT,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    serial_number VARCHAR(50),
    batch_no VARCHAR(50),
    process_id BIGINT,
    process_name VARCHAR(100),
    station_id BIGINT,
    station_name VARCHAR(100),
    operator_id BIGINT,
    operator_name VARCHAR(50),
    operate_time TIMESTAMP NOT NULL,
    operate_type VARCHAR(20),
    input_qty DECIMAL(18,4),
    output_qty DECIMAL(18,4),
    reject_qty DECIMAL(18,4) DEFAULT 0,
    UNIQUE (tenant_id, trace_no)
);
CREATE INDEX IF NOT EXISTS idx_tra_trace_record_tenant ON tra_trace_record(tenant_id);

-- 安东呼叫表
CREATE TABLE IF NOT EXISTS tra_andon_call (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    call_no VARCHAR(50) NOT NULL,
    line_id BIGINT,
    line_name VARCHAR(100),
    station_id BIGINT,
    station_name VARCHAR(100),
    call_type VARCHAR(20),
    call_level INTEGER DEFAULT 1,
    call_desc VARCHAR(500),
    call_user_id BIGINT,
    call_user_name VARCHAR(50),
    call_time TIMESTAMP NOT NULL,
    response_user_id BIGINT,
    response_time TIMESTAMP,
    resolve_time TIMESTAMP,
    status INTEGER DEFAULT 1,
    remark VARCHAR(500),
    UNIQUE (tenant_id, call_no)
);
CREATE INDEX IF NOT EXISTS idx_tra_andon_call_tenant ON tra_andon_call(tenant_id);

-- 数据采集表
CREATE TABLE IF NOT EXISTS tra_data_collection (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    equipment_id BIGINT,
    equipment_code VARCHAR(50),
    station_id BIGINT,
    data_type VARCHAR(20),
    data_key VARCHAR(50),
    data_value VARCHAR(200),
    unit VARCHAR(20),
    collect_time TIMESTAMP NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_tra_data_collection_tenant ON tra_data_collection(tenant_id);

-- =====================================================
-- 能源管理模块
-- =====================================================

-- 能源记录表
CREATE TABLE IF NOT EXISTS ene_energy_record (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tenant_id BIGINT NOT NULL,
    energy_type VARCHAR(20),
    workshop_id BIGINT,
    workshop_name VARCHAR(100),
    equipment_id BIGINT,
    equipment_name VARCHAR(100),
    meter_no VARCHAR(50),
    quantity DECIMAL(18,2),
    unit VARCHAR(20),
    unit_price DECIMAL(18,4),
    amount DECIMAL(18,2),
    record_date TIMESTAMP NOT NULL,
    remark VARCHAR(500)
);
CREATE INDEX IF NOT EXISTS idx_ene_energy_record_tenant ON ene_energy_record(tenant_id);
CREATE INDEX IF NOT EXISTS idx_ene_energy_record_date ON ene_energy_record(record_date);

-- =====================================================
-- 初始化数据
-- =====================================================

-- 插入默认租户
INSERT INTO sys_tenant (tenant_name, tenant_key, status) VALUES ('默认租户', 'default', 1);

-- 插入默认管理员用户 (密码: admin123)
INSERT INTO sys_user (tenant_id, username, nickname, password, status) VALUES
(1, 'admin', '管理员', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt6Z5E', 1);

-- 插入默认角色
INSERT INTO sys_role (tenant_id, role_name, role_key, role_sort, status) VALUES
(1, '超级管理员', 'admin', 1, 1),
(1, '普通用户', 'user', 2, 1);

-- 插入管理员菜单
INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, icon, sort, status) VALUES
(1, '首页', 'C', '/dashboard', 'HomeFilled', 1, 1),
(1, '系统管理', 'M', '/system', 'Setting', 2, 1),
(1, '用户管理', 'C', '/system/user', 'User', 1, 1),
(1, '角色管理', 'C', '/system/role', 'Key', 2, 1),
(1, '菜单管理', 'C', '/system/menu', 'Menu', 3, 1),
(1, '部门管理', 'C', '/system/dept', 'OfficeBuilding', 4, 1),
(1, '字典管理', 'C', '/system/dict', 'Document', 5, 1),
(1, '岗位管理', 'C', '/system/post', 'Postcard', 6, 1),
(1, '主数据', 'M', '/mdm', 'Box', 3, 1),
(1, '物料管理', 'C', '/mdm/material', 'Box', 1, 1),
(1, '车间管理', 'C', '/mdm/workshop', 'OfficeBuilding', 2, 1),
(1, '生产线管理', 'C', '/mdm/line', 'Connection', 3, 1),
(1, '工位管理', 'C', '/mdm/workstation', 'Grid', 4, 1),
(1, 'BOM管理', 'C', '/mdm/bom', 'Files', 5, 1),
(1, '工序管理', 'C', '/mdm/operation', 'Operation', 6, 1),
(1, '生产执行', 'M', '/production', 'List', 4, 1),
(1, '生产工单', 'C', '/production/order', 'List', 1, 1),
(1, '销售订单', 'C', '/production/sales-order', 'Document', 2, 1),
(1, '生产报工', 'C', '/production/report', 'DocumentCheck', 3, 1),
(1, '派工管理', 'C', '/production/dispatch', 'Tickets', 4, 1),
(1, '设备管理', 'M', '/equipment', 'Monitor', 5, 1),
(1, '设备台账', 'C', '/equipment', 'Monitor', 1, 1),
(1, '设备点检', 'C', '/equipment/check', 'Check', 2, 1),
(1, '设备保养', 'C', '/equipment/maintenance', 'Tools', 3, 1),
(1, '设备维修', 'C', '/equipment/repair', 'Tool', 4, 1),
(1, '备件管理', 'C', '/equipment/spare', 'Box', 5, 1),
(1, '仓储管理', 'M', '/wms', 'House', 6, 1),
(1, '仓库管理', 'C', '/wms/warehouse', 'House', 1, 1),
(1, '库位管理', 'C', '/wms/location', 'Location', 2, 1),
(1, '库存管理', 'C', '/wms/inventory', 'Box', 3, 1),
(1, '质量管理', 'M', '/quality', 'CircleCheck', 7, 1),
(1, 'IQC检验', 'C', '/quality/iqc', 'CircleCheck', 1, 1),
(1, 'APS计划', 'M', '/aps', 'Calendar', 8, 1),
(1, 'MPS计划', 'C', '/aps/mps', 'Calendar', 1, 1),
(1, 'MRP计划', 'C', '/aps/mrp', 'Grid', 2, 1),
(1, '排程计划', 'C', '/aps/schedule', 'List', 3, 1),
(1, '追溯管理', 'M', '/trace', 'Search', 9, 1),
(1, '追溯查询', 'C', '/trace/query', 'Search', 1, 1),
(1, '安东呼叫', 'C', '/trace/andon', 'Bell', 2, 1),
(1, '能源监控', 'C', '/energy/monitor', 'Lightning', 10, 1);

-- 插入字典类型
INSERT INTO sys_dict_type (dict_name, dict_type, status) VALUES
('用户状态', 'sys_user_status', 1),
('角色状态', 'sys_role_status', 1),
('菜单状态', 'sys_menu_status', 1),
('部门状态', 'sys_dept_status', 1),
('岗位状态', 'sys_post_status', 1),
('设备状态', 'equ_status', 1),
('设备类型', 'equ_type', 1),
('点检结果', 'check_result', 1),
('保养类型', 'maint_type', 1),
('维修状态', 'repair_status', 1),
('生产订单状态', 'pro_order_status', 1),
('销售订单状态', 'sales_order_status', 1),
('派工状态', 'dispatch_status', 1),
('仓库状态', 'warehouse_status', 1),
('仓库类型', 'warehouse_type', 1),
('库位状态', 'location_status', 1),
('库存状态', 'inventory_status', 1),
('单据状态', 'doc_status', 1),
('检验结果', 'qc_result', 1),
('能源类型', 'energy_type', 1);

-- 插入字典数据
INSERT INTO sys_dict_data (dict_sort, dict_label, dict_value, dict_type, status) VALUES
-- 用户状态
(1, '正常', '1', 'sys_user_status', 1),
(2, '停用', '0', 'sys_user_status', 1),
-- 角色状态
(1, '正常', '1', 'sys_role_status', 1),
(2, '停用', '0', 'sys_role_status', 1),
-- 设备状态
(1, '运行', '1', 'equ_status', 1),
(2, '待机', '2', 'equ_status', 1),
(3, '故障', '3', 'equ_status', 1),
(4, '维修', '4', 'equ_status', 1),
(5, '报废', '5', 'equ_status', 1),
-- 设备类型
(1, '生产设备', 'production', 'equ_type', 1),
(2, '检测设备', 'inspection', 'equ_type', 1),
(3, '辅助设备', 'auxiliary', 'equ_type', 1),
-- 点检结果
(1, '正常', '1', 'check_result', 1),
(2, '异常', '2', 'check_result', 1),
-- 保养类型
(1, '日常保养', 'daily', 'maint_type', 1),
(2, '一级保养', 'level1', 'maint_type', 1),
(3, '二级保养', 'level2', 'maint_type', 1),
-- 维修状态
(1, '待维修', '1', 'repair_status', 1),
(2, '维修中', '2', 'repair_status', 1),
(3, '已完成', '3', 'repair_status', 1),
-- 生产订单状态
(1, '待生产', '1', 'pro_order_status', 1),
(2, '生产中', '2', 'pro_order_status', 1),
(3, '已完成', '3', 'pro_order_status', 1),
(4, '已取消', '4', 'pro_order_status', 1),
-- 销售订单状态
(1, '待确认', '1', 'sales_order_status', 1),
(2, '已确认', '2', 'sales_order_status', 1),
(3, '生产中', '3', 'sales_order_status', 1),
(4, '已完成', '4', 'sales_order_status', 1),
(5, '已关闭', '5', 'sales_order_status', 1),
-- 派工状态
(1, '待开始', '1', 'dispatch_status', 1),
(2, '进行中', '2', 'dispatch_status', 1),
(3, '已完成', '3', 'dispatch_status', 1),
-- 仓库状态
(1, '正常', '1', 'warehouse_status', 1),
(2, '停用', '0', 'warehouse_status', 1),
-- 仓库类型
(1, '原料仓', 'raw', 'warehouse_type', 1),
(2, '成品仓', 'finished', 'warehouse_type', 1),
(3, '线边仓', 'inline', 'warehouse_type', 1),
(4, '工具仓', 'tool', 'warehouse_type', 1),
-- 能源类型
(1, '电', 'electricity', 'energy_type', 1),
(2, '水', 'water', 'energy_type', 1),
(3, '气', 'gas', 'energy_type', 1),
(4, '蒸汽', 'steam', 'energy_type', 1);

-- 给管理员分配所有菜单
INSERT INTO sys_role_menu (role_id, menu_id) SELECT 1, id FROM sys_menu WHERE tenant_id = 1;

-- 给管理员分配用户角色
INSERT INTO sys_user_role (user_id, role_id) VALUES (1, 1);

-- =====================================================
-- 初始化示例数据
-- =====================================================

-- 插入示例车间
INSERT INTO mdm_workshop (tenant_id, workshop_code, workshop_name, workshop_type, manager, phone, status) VALUES
(1, 'WS001', '一车间', 'assembly', '张三', '13800138001', 1),
(1, 'WS002', '二车间', 'machining', '李四', '13800138002', 1),
(1, 'WS003', '三车间', 'inspection', '王五', '13800138003', 1);

-- 插入示例生产线
INSERT INTO mdm_production_line (tenant_id, line_code, line_name, workshop_id, line_type, status) VALUES
(1, 'LINE001', '组装线A', 1, 'automation', 1),
(1, 'LINE002', '组装线B', 1, 'automation', 1),
(1, 'LINE003', '加工线C', 2, 'semiautomation', 1);

-- 插入示例工位
INSERT INTO mdm_workstation (tenant_id, station_code, station_name, line_id, station_type, status) VALUES
(1, 'ST001', '组装工位1', 1, 'assembly', 1),
(1, 'ST002', '组装工位2', 1, 'assembly', 1),
(1, 'ST003', '检测工位', 1, 'inspection', 1);

-- 插入示例仓库
INSERT INTO wms_warehouse (tenant_id, warehouse_code, warehouse_name, warehouse_type, manager, phone, status) VALUES
(1, 'WH001', '原材料仓', 'raw', '仓管A', '13900139001', 1),
(1, 'WH002', '成品仓', 'finished', '仓管B', '13900139002', 1),
(1, 'WH003', '线边仓', 'inline', '仓管C', '13900139003', 1);

-- 插入示例设备
INSERT INTO equ_equipment (tenant_id, equipment_code, equipment_name, equipment_type, brand, model, workshop_id, line_id, station_id, status) VALUES
(1, 'EQ001', '装配机器人A1', 'production', 'ABB', 'IRB120', 1, 1, 1, 1),
(1, 'EQ002', '检测设备C1', 'inspection', 'KEYENCE', 'CV-X300', 3, NULL, 3, 1),
(1, 'EQ003', '加工中心M1', 'production', 'DMG', 'CMX50U', 2, 3, NULL, 1);

-- 完成
COMMENT ON DATABASE mom3 IS 'MOM3.0 Manufacturing Operation Management System';
