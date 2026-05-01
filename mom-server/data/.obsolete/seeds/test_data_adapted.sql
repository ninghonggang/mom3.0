-- MOM3.0 测试数据 (适配实际表结构 v2)

BEGIN;

-- ============================================
-- 0. 租户数据 (sys_tenant)
-- ============================================
INSERT INTO sys_tenant (id, created_at, updated_at, tenant_name, tenant_key, province, city, address, manager, contact_name, contact_phone, employee_count, status) VALUES
(1, NOW(), NOW(), '富维电子科技有限公司', 'FUWEI', '广东省', '深圳市', '宝安区福永街道', '张总', '张总', '13800000001', 200, 1);

-- ============================================
-- 1. 部门数据 (sys_dept)
-- ============================================
INSERT INTO sys_dept (id, created_at, updated_at, tenant_id, parent_id, dept_name, dept_code, dept_sort, leader, phone, email, status) VALUES
(1, NOW(), NOW(), 1, 0, '总经理办公室', 'D001', 1, '张总', '13800000001', 'zhang@fuwei.com', 1),
(2, NOW(), NOW(), 1, 0, '研发部', 'D002', 2, '李总监', '13800000002', 'li@fuwei.com', 1),
(3, NOW(), NOW(), 1, 0, '生产部', 'D003', 3, '王总监', '13800000003', 'wang@fuwei.com', 1),
(4, NOW(), NOW(), 1, 0, '质量部', 'D004', 4, '刘总监', '13800000004', 'liu@fuwei.com', 1),
(5, NOW(), NOW(), 1, 0, '采购部', 'D005', 5, '陈总监', '13800000005', 'chen@fuwei.com', 1),
(6, NOW(), NOW(), 1, 0, '仓储部', 'D006', 6, '赵主管', '13800000006', 'zhao@fuwei.com', 1),
(7, NOW(), NOW(), 1, 0, '设备部', 'D007', 7, '孙主管', '13800000007', 'sun@fuwei.com', 1),
(8, NOW(), NOW(), 1, 0, '销售部', 'D008', 8, '周经理', '13800000008', 'zhou@fuwei.com', 1),
(9, NOW(), NOW(), 1, 0, '财务部', 'D009', 9, '吴经理', '13800000009', 'wu@fuwei.com', 1),
(10, NOW(), NOW(), 1, 0, '人力资源部', 'D010', 10, '郑经理', '13800000010', 'zheng@fuwei.com', 1),
(11, NOW(), NOW(), 1, 3, '生产一车间', 'D011', 11, '钱主任', '13800000011', 'qian@fuwei.com', 1),
(12, NOW(), NOW(), 1, 3, '生产二车间', 'D012', 12, '沈主任', '13800000012', 'shen@fuwei.com', 1),
(13, NOW(), NOW(), 1, 3, '生产三车间', 'D013', 13, '韩主任', '13800000013', 'han@fuwei.com', 1),
(14, NOW(), NOW(), 1, 3, '装配车间', 'D014', 14, '唐主任', '13800000014', 'tang@fuwei.com', 1),
(15, NOW(), NOW(), 1, 6, '原材料仓库', 'D015', 15, '林主管', '13800000015', 'lin@fuwei.com', 1),
(16, NOW(), NOW(), 1, 6, '成品仓库', 'D016', 16, '柳主管', '13800000016', 'liu2@fuwei.com', 1);

-- ============================================
-- 2. 岗位数据 (sys_post)
-- ============================================
INSERT INTO sys_post (id, created_at, updated_at, tenant_id, post_code, post_name, post_sort, status) VALUES
(1, NOW(), NOW(), 1, 'CEO', '首席执行官', 1, 1),
(2, NOW(), NOW(), 1, 'CTO', '首席技术官', 2, 1),
(3, NOW(), NOW(), 1, 'COO', '首席运营官', 3, 1),
(4, NOW(), NOW(), 1, 'MGR', '部门经理', 4, 1),
(5, NOW(), NOW(), 1, 'SPV', '主管', 5, 1),
(6, NOW(), NOW(), 1, 'ENG', '工程师', 6, 1),
(7, NOW(), NOW(), 1, 'TEC', '技术员', 7, 1),
(8, NOW(), NOW(), 1, 'OPR', '操作员', 8, 1),
(9, NOW(), NOW(), 1, 'QC', '质检员', 9, 1),
(10, NOW(), NOW(), 1, 'WH', '仓库管理员', 10, 1),
(11, NOW(), NOW(), 1, 'PUR', '采购员', 11, 1),
(12, NOW(), NOW(), 1, 'SAL', '销售员', 12, 1),
(13, NOW(), NOW(), 1, 'FIN', '财务', 13, 1),
(14, NOW(), NOW(), 1, 'HR', '人事', 14, 1);

-- ============================================
-- 3. 用户数据 (sys_user) - 密码是 123456
-- ============================================
INSERT INTO sys_user (id, created_at, updated_at, tenant_id, username, nickname, password, email, phone, dept_id, status) VALUES
(1, NOW(), NOW(), 1, 'admin', '系统管理员', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'admin@fuwei.com', '13900000001', 1, 1),
(2, NOW(), NOW(), 1, 'zhangsan', '张三', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'zhang@fuwei.com', '13900000002', 1, 1),
(3, NOW(), NOW(), 1, 'lisi', '李四', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'li@fuwei.com', '13900000003', 2, 1),
(4, NOW(), NOW(), 1, 'wangwu', '王五', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'wang@fuwei.com', '13900000004', 3, 1),
(5, NOW(), NOW(), 1, 'zhaoliu', '赵六', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'liu@fuwei.com', '13900000005', 4, 1),
(6, NOW(), NOW(), 1, 'sunqi', '孙七', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'sun@fuwei.com', '13900000006', 5, 1),
(7, NOW(), NOW(), 1, 'zhouba', '周八', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'zhou@fuwei.com', '13900000007', 6, 1),
(8, NOW(), NOW(), 1, 'wujiu', '吴九', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'wu@fuwei.com', '13900000008', 7, 1),
(9, NOW(), NOW(), 1, 'zhengshi', '郑十', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'zheng@fuwei.com', '13900000009', 8, 1),
(10, NOW(), NOW(), 1, 'chenyi', '陈一', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'cheny@fuwei.com', '13900000010', 9, 1),
(11, NOW(), NOW(), 1, 'qianer', '钱二', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'qian@fuwei.com', '13900000011', 11, 1),
(12, NOW(), NOW(), 1, 'shensan', '沈三', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'shen@fuwei.com', '13900000012', 11, 1),
(13, NOW(), NOW(), 1, 'hansi', '韩四', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'han@fuwei.com', '13900000013', 11, 1),
(14, NOW(), NOW(), 1, 'tangwu', '唐五', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'tang@fuwei.com', '13900000014', 12, 1),
(15, NOW(), NOW(), 1, 'linliu', '林六', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'lin@fuwei.com', '13900000015', 15, 1),
(16, NOW(), NOW(), 1, 'liuqi', '柳七', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'liu2@fuwei.com', '13900000016', 16, 1);

-- ============================================
-- 4. 角色数据 (sys_role)
-- ============================================
INSERT INTO sys_role (id, created_at, updated_at, tenant_id, role_name, role_key, role_sort, data_scope, status) VALUES
(1, NOW(), NOW(), 1, '超级管理员', 'superadmin', 1, 1, 1),
(2, NOW(), NOW(), 1, '生产经理', 'production_manager', 2, 2, 1),
(3, NOW(), NOW(), 1, '质量经理', 'quality_manager', 3, 2, 1),
(4, NOW(), NOW(), 1, '仓库管理员', 'warehouse_manager', 4, 2, 1),
(5, NOW(), NOW(), 1, '操作员', 'operator', 5, 3, 1);

-- ============================================
-- 5. 用户角色关系 (sys_user_role)
-- ============================================
INSERT INTO sys_user_role (user_id, role_id) VALUES
(1, 1), (2, 2), (3, 3), (4, 4), (5, 5);

-- ============================================
-- 6. 字典类型 (sys_dict_type)
-- ============================================
INSERT INTO sys_dict_type (id, created_at, updated_at, dict_name, dict_type, status) VALUES
(1, NOW(), NOW(), '物料类型', 'material_type', 1),
(2, NOW(), NOW(), '仓库类型', 'warehouse_type', 1),
(3, NOW(), NOW(), '工单状态', 'order_status', 1),
(4, NOW(), NOW(), '质检类型', 'qc_type', 1),
(5, NOW(), NOW(), '优先级', 'priority', 1),
(6, NOW(), NOW(), '班次', 'shift_type', 1),
(7, NOW(), NOW(), '区域', 'zone', 1),
(8, NOW(), NOW(), '单位', 'unit', 1),
(9, NOW(), NOW(), '审批状态', 'approval_status', 1);

-- ============================================
-- 7. 字典数据 (sys_dict_data)
-- ============================================
INSERT INTO sys_dict_data (id, created_at, updated_at, tenant_id, dict_type, dict_label, dict_value, sort, status) VALUES
-- 物料类型
(1, NOW(), NOW(), 1, 'material_type', '原材料', 'raw', 1, 1),
(2, NOW(), NOW(), 1, 'material_type', '半成品', 'semi', 2, 1),
(3, NOW(), NOW(), 1, 'material_type', '成品', 'finished', 3, 1),
(4, NOW(), NOW(), 1, 'material_type', '包材', 'packaging', 4, 1),
(5, NOW(), NOW(), 1, 'material_type', '辅料', 'auxiliary', 5, 1),
-- 仓库类型
(10, NOW(), NOW(), 1, 'warehouse_type', '原材料仓', 'raw', 1, 1),
(11, NOW(), NOW(), 1, 'warehouse_type', '成品仓', 'finished', 2, 1),
(12, NOW(), NOW(), 1, 'warehouse_type', '半成品仓', 'semi', 3, 1),
-- 工单状态
(20, NOW(), NOW(), 1, 'order_status', '待下达', 'pending', 1, 1),
(21, NOW(), NOW(), 1, 'order_status', '已下达', 'released', 2, 1),
(22, NOW(), NOW(), 1, 'order_status', '生产中', 'in_production', 3, 1),
(23, NOW(), NOW(), 1, 'order_status', '已完成', 'completed', 4, 1),
(24, NOW(), NOW(), 1, 'order_status', '已关闭', 'closed', 5, 1),
-- 质检类型
(30, NOW(), NOW(), 1, 'qc_type', 'IQC来料检', 'iqc', 1, 1),
(31, NOW(), NOW(), 1, 'qc_type', 'IPQC过程检', 'ipqc', 2, 1),
(32, NOW(), NOW(), 1, 'qc_type', 'FQC成品检', 'fqc', 3, 1),
(33, NOW(), NOW(), 1, 'qc_type', 'OQC出货检', 'oqc', 4, 1),
-- 优先级
(40, NOW(), NOW(), 1, 'priority', '紧急', 'urgent', 1, 1),
(41, NOW(), NOW(), 1, 'priority', '高', 'high', 2, 1),
(42, NOW(), NOW(), 1, 'priority', '普通', 'normal', 3, 1),
(43, NOW(), NOW(), 1, 'priority', '低', 'low', 4, 1),
-- 班次
(50, NOW(), NOW(), 1, 'shift_type', '白班', 'day', 1, 1),
(51, NOW(), NOW(), 1, 'shift_type', '中班', 'swing', 2, 1),
(52, NOW(), NOW(), 1, 'shift_type', '夜班', 'night', 3, 1),
-- 区域
(60, NOW(), NOW(), 1, 'zone', 'A区', 'A', 1, 1),
(61, NOW(), NOW(), 1, 'zone', 'B区', 'B', 2, 1),
(62, NOW(), NOW(), 1, 'zone', 'C区', 'C', 3, 1),
(63, NOW(), NOW(), 1, 'zone', 'D区', 'D', 4, 1),
-- 单位
(70, NOW(), NOW(), 1, 'unit', '个', 'PCS', 1, 1),
(71, NOW(), NOW(), 1, 'unit', '台', 'SET', 2, 1),
(72, NOW(), NOW(), 1, 'unit', '箱', 'BOX', 3, 1),
(73, NOW(), NOW(), 1, 'unit', '千克', 'KG', 4, 1),
(74, NOW(), NOW(), 1, 'unit', '米', 'M', 5, 1),
-- 审批状态
(80, NOW(), NOW(), 1, 'approval_status', '待审批', 'pending', 1, 1),
(81, NOW(), NOW(), 1, 'approval_status', '已批准', 'approved', 2, 1),
(82, NOW(), NOW(), 1, 'approval_status', '已拒绝', 'rejected', 3, 1),
(83, NOW(), NOW(), 1, 'approval_status', '已撤回', 'withdrawn', 4, 1);

-- ============================================
-- 8. 物料数据 (mdm_material)
-- ============================================
INSERT INTO mdm_material (id, created_at, updated_at, tenant_id, material_code, material_name, material_type, spec, unit, unit_name, status) VALUES
(1, NOW(), NOW(), 1, 'MAT-001', '钢板', 'raw', '1000*2000*2mm', 'PCS', '张', 1),
(2, NOW(), NOW(), 1, 'MAT-002', '铝板', 'raw', '1000*2000*1.5mm', 'PCS', '张', 1),
(3, NOW(), NOW(), 1, 'MAT-003', '铜板', 'raw', '500*1000*1mm', 'PCS', '张', 1),
(4, NOW(), NOW(), 1, 'MAT-004', 'ABS塑料粒子', 'raw', 'ABS粒料', 'KG', '千克', 1),
(5, NOW(), NOW(), 1, 'MAT-005', 'PC塑料粒子', 'raw', 'PC粒料', 'KG', '千克', 1),
(6, NOW(), NOW(), 1, 'MAT-006', '螺丝M3*10', 'auxiliary', '不锈钢', 'PCS', '个', 1),
(7, NOW(), NOW(), 1, 'MAT-007', '螺丝M4*15', 'auxiliary', '不锈钢', 'PCS', '个', 1),
(8, NOW(), NOW(), 1, 'MAT-008', '轴承608', 'auxiliary', '内径8mm', 'PCS', '个', 1),
(9, NOW(), NOW(), 1, 'MAT-009', '轴承6200', 'auxiliary', '内径10mm', 'PCS', '个', 1),
(10, NOW(), NOW(), 1, 'MAT-010', '电机 DC220V 50W', 'auxiliary', '直流无刷', 'PCS', '个', 1),
(11, NOW(), NOW(), 1, 'MAT-011', '电机 DC220V 100W', 'auxiliary', '直流无刷', 'PCS', '个', 1),
(12, NOW(), NOW(), 1, 'MAT-012', 'PCB板 4层', 'auxiliary', '4层板', 'PCS', '个', 1),
(13, NOW(), NOW(), 1, 'MAT-013', 'PCB板 6层', 'auxiliary', '6层板', 'PCS', '个', 1),
(14, NOW(), NOW(), 1, 'MAT-014', '显示屏 4寸', 'auxiliary', '480*320', 'PCS', '个', 1),
(15, NOW(), NOW(), 1, 'MAT-015', '触摸屏 5寸', 'auxiliary', '电容式', 'PCS', '个', 1),
(16, NOW(), NOW(), 1, 'MAT-016', '连接线 20cm', 'auxiliary', 'AWG22', 'PCS', '条', 1),
(17, NOW(), NOW(), 1, 'MAT-017', '散热器 40*40', 'auxiliary', '铝合金', 'PCS', '个', 1),
(18, NOW(), NOW(), 1, 'MAT-018', '包装盒 L', 'packaging', '200*150*80', 'PCS', '个', 1),
(19, NOW(), NOW(), 1, 'MAT-019', '说明书 A5', 'packaging', '210*148', 'PCS', '张', 1),
(20, NOW(), NOW(), 1, 'MAT-020', '保修卡', 'packaging', 'A6', 'PCS', '张', 1),
-- 半成品
(21, NOW(), NOW(), 1, 'SEM-001', '半成品主板A', 'semi', '已贴片未测试', 'PCS', '个', 1),
(22, NOW(), NOW(), 1, 'SEM-002', '半成品主板B', 'semi', '已测试未组装', 'PCS', '个', 1),
(23, NOW(), NOW(), 1, 'SEM-003', '显示模块', 'semi', '已组装', 'PCS', '个', 1),
-- 成品
(31, NOW(), NOW(), 1, 'PRD-001', '智能控制器 V1', 'finished', 'DC220V输入', 'PCS', '台', 1),
(32, NOW(), NOW(), 1, 'PRD-002', '智能控制器 V2', 'finished', 'DC220V输入', 'PCS', '台', 1),
(33, NOW(), NOW(), 1, 'PRD-003', '触摸屏模块', 'finished', '5寸电容', 'PCS', '台', 1),
(34, NOW(), NOW(), 1, 'PRD-004', '传感器模块', 'finished', '温度湿度', 'PCS', '台', 1);

-- ============================================
-- 9. 物料分类 (mdm_material_category)
-- ============================================
INSERT INTO mdm_material_category (id, created_at, updated_at, tenant_id, category_code, category_name, parent_id, sort, status) VALUES
(1, NOW(), NOW(), 1, 'C01', '原材料', 0, 1, 1),
(2, NOW(), NOW(), 1, 'C02', '半成品', 0, 2, 1),
(3, NOW(), NOW(), 1, 'C03', '成品', 0, 3, 1),
(4, NOW(), NOW(), 1, 'C04', '包材', 0, 4, 1),
(5, NOW(), NOW(), 1, 'C05', '辅料', 0, 5, 1),
(11, NOW(), NOW(), 1, 'C011', '金属材料', 1, 1, 1),
(12, NOW(), NOW(), 1, 'C012', '塑料粒子', 1, 2, 1),
(13, NOW(), NOW(), 1, 'C013', '标准件', 5, 3, 1),
(14, NOW(), NOW(), 1, 'C014', '电气件', 5, 4, 1);

-- ============================================
-- 10. 车间数据 (mdm_workshop)
-- ============================================
INSERT INTO mdm_workshop (id, created_at, updated_at, tenant_id, workshop_code, workshop_name, workshop_type, manager, phone, address, status) VALUES
(1, NOW(), NOW(), 1, 'WS01', '生产一车间', 'main', '钱主任', '13800000011', 'A栋1楼', 1),
(2, NOW(), NOW(), 1, 'WS02', '生产二车间', 'main', '沈主任', '13800000012', 'A栋2楼', 1),
(3, NOW(), NOW(), 1, 'WS03', '生产三车间', 'main', '韩主任', '13800000013', 'B栋1楼', 1),
(4, NOW(), NOW(), 1, 'WS04', '装配车间', 'assembly', '唐主任', '13800000014', 'B栋2楼', 1);

-- ============================================
-- 11. 产线数据 (mdm_production_line)
-- ============================================
INSERT INTO mdm_production_line (id, created_at, updated_at, tenant_id, line_code, line_name, workshop_id, line_type, status) VALUES
(1, NOW(), NOW(), 1, 'LINE01', '生产线A1', 1, 'assembly', 1),
(2, NOW(), NOW(), 1, 'LINE02', '生产线A2', 1, 'assembly', 1),
(3, NOW(), NOW(), 1, 'LINE03', '生产线B1', 2, 'assembly', 1),
(4, NOW(), NOW(), 1, 'LINE04', '生产线B2', 2, 'assembly', 1),
(5, NOW(), NOW(), 1, 'LINE05', '生产线C1', 3, 'machining', 1),
(6, NOW(), NOW(), 1, 'LINE06', '生产线C2', 3, 'machining', 1),
(7, NOW(), NOW(), 1, 'LINE07', '装配线D1', 4, 'assembly', 1),
(8, NOW(), NOW(), 1, 'LINE08', '装配线D2', 4, 'assembly', 1);

-- ============================================
-- 12. 班次数据 (mdm_shift)
-- ============================================
INSERT INTO mdm_shift (id, created_at, updated_at, tenant_id, shift_code, shift_name, start_time, end_time, work_hours, is_night, status) VALUES
(1, NOW(), NOW(), 1, 'D', '白班', '08:00', '16:00', 8, 0, 1),
(2, NOW(), NOW(), 1, 'S', '中班', '16:00', '24:00', 8, 0, 1),
(3, NOW(), NOW(), 1, 'N', '夜班', '00:00', '08:00', 8, 1, 1);

-- ============================================
-- 13. 仓库数据 (wms_warehouse)
-- ============================================
INSERT INTO wms_warehouse (id, created_at, updated_at, tenant_id, warehouse_code, warehouse_name, warehouse_type, address, manager, phone, status) VALUES
(1, NOW(), NOW(), 1, 'WH01', '原材料仓库', 'raw', 'C栋1楼', '林主管', '13800000015', 1),
(2, NOW(), NOW(), 1, 'WH02', '成品仓库', 'finished', 'C栋2楼', '柳主管', '13800000016', 1),
(3, NOW(), NOW(), 1, 'WH03', '半成品仓库', 'semi', 'C栋3楼', '林主管', '13800000015', 1),
(4, NOW(), NOW(), 1, 'WH04', '包材仓库', 'packaging', 'C栋4楼', '林主管', '13800000015', 1);

-- ============================================
-- 14. 库位数据 (wms_location)
-- ============================================
INSERT INTO wms_location (id, created_at, updated_at, tenant_id, location_code, location_name, warehouse_id, zone_code, row, col, layer, location_type, capacity, status) VALUES
(1, NOW(), NOW(), 1, 'R-A01-01-01', 'R-A01-01-01', 1, 'A', 1, 1, 1, 'storage', 100, 1),
(2, NOW(), NOW(), 1, 'R-A01-01-02', 'R-A01-01-02', 1, 'A', 1, 1, 2, 'storage', 100, 1),
(3, NOW(), NOW(), 1, 'R-A01-02-01', 'R-A01-02-01', 1, 'A', 1, 2, 1, 'storage', 100, 1),
(4, NOW(), NOW(), 1, 'R-B01-01-01', 'R-B01-01-01', 1, 'B', 1, 1, 1, 'storage', 100, 1),
(5, NOW(), NOW(), 1, 'R-B02-01-01', 'R-B02-01-01', 1, 'B', 2, 1, 1, 'storage', 100, 1),
(10, NOW(), NOW(), 1, 'F-A01-01-01', 'F-A01-01-01', 2, 'A', 1, 1, 1, 'storage', 50, 1),
(11, NOW(), NOW(), 1, 'F-A01-01-02', 'F-A01-01-02', 2, 'A', 1, 1, 2, 'storage', 50, 1),
(12, NOW(), NOW(), 1, 'F-B01-01-01', 'F-B01-01-01', 2, 'B', 1, 1, 1, 'storage', 50, 1),
(20, NOW(), NOW(), 1, 'S-A01-01-01', 'S-A01-01-01', 3, 'A', 1, 1, 1, 'storage', 80, 1),
(21, NOW(), NOW(), 1, 'S-A02-01-01', 'S-A02-01-01', 3, 'A', 2, 1, 1, 'storage', 80, 1);

-- ============================================
-- 15. 销售订单 (pro_sales_order)
-- ============================================
INSERT INTO pro_sales_order (id, created_at, updated_at, tenant_id, order_no, customer_name, order_date, delivery_date, order_type, priority, status, remark) VALUES
(1, NOW(), NOW(), 1, 'SO202604001', '华为技术有限公司', '2026-04-20', '2026-05-15', 'standard', 2, 2, '优先交付'),
(2, NOW(), NOW(), 1, 'SO202604002', '比亚迪股份有限公司', '2026-04-21', '2026-05-20', 'standard', 3, 2, ''),
(3, NOW(), NOW(), 1, 'SO202604003', '中兴通讯', '2026-04-22', '2026-05-25', 'standard', 2, 1, '加急订单'),
(4, NOW(), NOW(), 1, 'SO202604004', '海尔智家', '2026-04-23', '2026-06-01', 'standard', 3, 1, ''),
(5, NOW(), NOW(), 1, 'SO202604005', '格力电器', '2026-04-24', '2026-06-05', 'standard', 3, 1, '');

-- ============================================
-- 16. 生产工单 (pro_production_order)
-- ============================================
INSERT INTO pro_production_order (id, created_at, updated_at, tenant_id, order_no, sales_order_no, material_id, material_code, material_name, material_spec, unit, quantity, workshop_id, workshop_name, line_id, line_name, plan_start_date, plan_end_date, priority, status, remark) VALUES
(1, NOW(), NOW(), 1, 'WO202604001', 'SO202604001', 31, 'PRD-001', '智能控制器 V1', 'DC220V输入', 'PCS', 500, 1, '生产一车间', 1, '生产线A1', '2026-04-26', '2026-05-10', 2, 2, ''),
(2, NOW(), NOW(), 1, 'WO202604002', 'SO202604001', 32, 'PRD-002', '智能控制器 V2', 'DC220V输入', 'PCS', 300, 1, '生产一车间', 2, '生产线A2', '2026-04-26', '2026-05-12', 2, 2, ''),
(3, NOW(), NOW(), 1, 'WO202604003', 'SO202604002', 33, 'PRD-003', '触摸屏模块', '5寸电容', 'PCS', 200, 4, '装配车间', 7, '装配线D1', '2026-04-27', '2026-05-15', 3, 1, ''),
(4, NOW(), NOW(), 1, 'WO202604004', 'SO202604003', 34, 'PRD-004', '传感器模块', '温度湿度', 'PCS', 400, 2, '生产二车间', 3, '生产线B1', '2026-04-28', '2026-05-20', 1, 1, '加急'),
(5, NOW(), NOW(), 1, 'WO202604005', 'SO202604004', 31, 'PRD-001', '智能控制器 V1', 'DC220V输入', 'PCS', 150, 3, '生产三车间', 5, '生产线C1', '2026-04-29', '2026-05-18', 3, 1, '');

-- ============================================
-- 17. 库存数据 (wms_inventory)
-- ============================================
INSERT INTO wms_inventory (id, created_at, updated_at, tenant_id, material_id, location_id, batch_no, quantity, status) VALUES
(1, NOW(), NOW(), 1, 1, 1, 'BATCH20260401', 500, 1),
(2, NOW(), NOW(), 1, 2, 2, 'BATCH20260402', 300, 1),
(3, NOW(), NOW(), 1, 4, 3, 'BATCH20260403', 1000, 1),
(4, NOW(), NOW(), 1, 6, 4, 'BATCH20260404', 5000, 1),
(5, NOW(), NOW(), 1, 8, 5, 'BATCH20260405', 2000, 1),
(6, NOW(), NOW(), 1, 31, 10, 'BATCH20260410', 50, 1),
(7, NOW(), NOW(), 1, 32, 11, 'BATCH20260411', 30, 1);

-- ============================================
-- 18. 采购订单 (scp_purchase_order)
-- ============================================
INSERT INTO scp_purchase_order (id, created_at, updated_at, tenant_id, po_no, po_type, supplier_name, order_date, promised_date, total_amount, total_qty, approved_by, approved_time, approval_status, status, remark) VALUES
(1, NOW(), NOW(), 1, 'PO202604001', 'standard', '上海金属材料有限公司', '2026-04-15', '2026-04-25', 50000.00, 500, 2, '2026-04-16', 'approved', '4', '已入库'),
(2, NOW(), NOW(), 1, 'PO202604002', 'standard', '深圳塑料科技有限公司', '2026-04-18', '2026-04-28', 30000.00, 1000, 2, '2026-04-19', 'approved', '3', '运输中'),
(3, NOW(), NOW(), 1, 'PO202604003', 'standard', '东莞标准件厂', '2026-04-20', '2026-04-30', 20000.00, 5000, 2, NULL, 'pending', '2', ''),
(4, NOW(), NOW(), 1, 'PO202604004', 'standard', '广州电子股份有限公司', '2026-04-22', '2026-05-02', 80000.00, 200, 2, NULL, 'pending', '1', '');

-- ============================================
-- 19. Andon 呼叫记录 (andon_call)
-- ============================================
INSERT INTO andon_call (id, created_at, updated_at, tenant_id, call_no, workshop_id, workshop_name, production_line_id, production_line_name, workstation_id, workstation_name, andon_type, andon_type_name, call_level, priority, description, call_by, call_time, status) VALUES
(1, NOW(), NOW(), 1, 'AND202604001', 1, '生产一车间', 1, '生产线A1', 1, 'A1-001', 'equipment', '设备故障', 2, 2, '贴片机报警E01', '钱二', NOW(), 'processing'),
(2, NOW(), NOW(), 1, 'AND202604002', 1, '生产一车间', 1, '生产线A1', 2, 'A1-002', 'quality', '质量问题', 1, 1, '发现来料异常', '钱二', NOW(), 'open'),
(3, NOW(), NOW(), 1, 'AND202604003', 2, '生产二车间', 3, '生产线B1', 1, 'B1-001', 'material', '物料短缺', 3, 3, '缺少MAT-003物料', '沈三', NOW(), 'open'),
(4, NOW(), NOW(), 1, 'AND202604004', 4, '装配车间', 7, '装配线D1', 1, 'D1-001', 'equipment', '设备故障', 2, 2, '组装工位治具损坏', '唐五', NOW(), 'open');

-- ============================================
-- 20. 质检单 (qc_iqc / qc_ipqc / qc_fqc / qc_oqc)
-- ============================================
INSERT INTO qc_iqc (id, created_at, updated_at, tenant_id, iqc_no, supplier_id, supplier_name, material_id, material_code, material_name, quantity, unit, check_user_id, check_user_name, check_date, result, remark) VALUES
(1, NOW(), NOW(), 1, 'IQC202604001', 1, '上海金属材料有限公司', 1, 'MAT-001', '钢板', 500, 'PCS', 3, '李四', NOW(), 1, '合格'),
(2, NOW(), NOW(), 1, 'IQC202604002', 2, '深圳塑料科技有限公司', 4, 'MAT-004', 'ABS塑料粒子', 1000, 'KG', 3, '李四', NOW(), 2, '部分不良');

INSERT INTO qc_ipqc (id, created_at, updated_at, tenant_id, ip_qc_no, order_id, order_no, process_id, process_name, quantity, sample_size, check_user_id, check_user_name, check_date, result, remark) VALUES
(1, NOW(), NOW(), 1, 'IPQC202604001', 1, 'WO202604001', 1, 'SMT贴装', 100, 20, 3, '李四', NOW(), 1, '首件合格'),
(2, NOW(), NOW(), 1, 'IPQC202604002', 2, 'WO202604002', 2, 'DIP焊接', 50, 10, 3, '李四', NOW(), 1, '合格');

INSERT INTO qc_fqc (id, created_at, updated_at, tenant_id, fqc_no, order_id, order_no, workshop_id, workshop_name, line_id, line_name, quantity, sample_size, check_user_id, check_user_name, check_date, result, remark) VALUES
(1, NOW(), NOW(), 1, 'FQC202604001', 1, 'WO202604001', 1, '生产一车间', 1, '生产线A1', 100, 20, 3, '李四', NOW(), 1, '抽检合格');

INSERT INTO qc_oqc (id, created_at, updated_at, tenant_id, oqc_no, order_id, order_no, sales_order_no, quantity, sample_size, check_user_id, check_user_name, check_date, result, remark) VALUES
(1, NOW(), NOW(), 1, 'OQC202604001', 1, 'WO202604001', 'SO202604001', 50, 10, 3, '李四', NOW(), 1, '出货抽检合格');

-- ============================================
-- 21. 追溯数据 (tra_serial_number)
-- ============================================
INSERT INTO tra_serial_number (id, created_at, updated_at, tenant_id, serial_no, product_code, batch_no, workshop_id, line_id, worker_id, status) VALUES
(1, NOW(), NOW(), 1, 'SN20260400001', 'PRD-001', 'BATCH20260410', 4, 7, 11, 1),
(2, NOW(), NOW(), 1, 'SN20260400002', 'PRD-001', 'BATCH20260410', 4, 7, 11, 1),
(3, NOW(), NOW(), 1, 'SN20260400003', 'PRD-002', 'BATCH20260411', 4, 8, 11, 1),
(4, NOW(), NOW(), 1, 'SN20260400004', 'PRD-001', 'BATCH20260410', 1, 1, 12, 1),
(5, NOW(), NOW(), 1, 'SN20260400005', 'PRD-003', 'BATCH20260412', 4, 7, 11, 1);

COMMIT;
