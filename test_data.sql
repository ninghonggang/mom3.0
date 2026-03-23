-- MOM3.0 测试数据脚本
-- 执行: docker exec -i mom-postgres psql -U mes -d mom3_db < test_data.sql

BEGIN;

-- ============================================
-- 1. 部门数据 (20条)
-- ============================================
INSERT INTO depts (created_at, updated_at, tenant_id, name, parent_id, leader, phone, sort) VALUES
(NOW(), NOW(), 1, '总经理办公室', 0, '张总', '13800000001', 1),
(NOW(), NOW(), 1, '研发部', 0, '李总监', '13800000002', 2),
(NOW(), NOW(), 1, '生产部', 0, '王总监', '13800000003', 3),
(NOW(), NOW(), 1, '质量部', 0, '刘总监', '13800000004', 4),
(NOW(), NOW(), 1, '采购部', 0, '陈总监', '13800000005', 5),
(NOW(), NOW(), 1, '仓储部', 0, '赵总监', '13800000006', 6),
(NOW(), NOW(), 1, '设备部', 0, '孙总监', '13800000007', 7),
(NOW(), NOW(), 1, '销售部', 0, '周总监', '13800000008', 8),
(NOW(), NOW(), 1, '财务部', 0, '吴总监', '13800000009', 9),
(NOW(), NOW(), 1, '人力资源部', 0, '郑总监', '13800000010', 10),
(NOW(), NOW(), 1, '生产一车间', 3, '钱主任', '13800000011', 11),
(NOW(), NOW(), 1, '生产二车间', 3, '沈主任', '13800000012', 12),
(NOW(), NOW(), 1, '生产三车间', 3, '韩主任', '13800000013', 13),
(NOW(), NOW(), 1, '装配车间', 3, '唐主任', '13800000014', 14),
(NOW(), NOW(), 1, '测试车间', 3, '冯主任', '13800000015', 15),
(NOW(), NOW(), 1, '包装车间', 3, '于主任', '13800000016', 16),
(NOW(), NOW(), 1, '模具车间', 3, '董主任', '13800000017', 17),
(NOW(), NOW(), 1, '原材料仓库', 6, '林主管', '13800000018', 18),
(NOW(), NOW(), 1, '成品仓库', 6, '柳主管', '13800000019', 19),
(NOW(), NOW(), 1, '半成品仓库', 6, '潘主管', '13800000020', 20);

-- ============================================
-- 2. 岗位数据 (20条)
-- ============================================
INSERT INTO posts (created_at, updated_at, tenant_id, post_code, post_name, sort, status) VALUES
(NOW(), NOW(), 1, 'CEO', '首席执行官', 1, 1),
(NOW(), NOW(), 1, 'CTO', '首席技术官', 2, 1),
(NOW(), NOW(), 1, 'COO', '首席运营官', 3, 1),
(NOW(), NOW(), 1, 'MGR', '部门经理', 4, 1),
(NOW(), NOW(), 1, 'SPV', '主管', 5, 1),
(NOW(), NOW(), 1, 'ENG', '工程师', 6, 1),
(NOW(), NOW(), 1, 'TEC', '技术员', 7, 1),
(NOW(), NOW(), 1, 'OPR', '操作员', 8, 1),
(NOW(), NOW(), 1, 'QC', '质检员', 9, 1),
(NOW(), NOW(), 1, 'WH', '仓库管理员', 10, 1),
(NOW(), NOW(), 1, 'PUR', '采购员', 11, 1),
(NOW(), NOW(), 1, 'SAL', '销售员', 12, 1),
(NOW(), NOW(), 1, 'FIN', '财务', 13, 1),
(NOW(), NOW(), 1, 'HR', '人事', 14, 1),
(NOW(), NOW(), 1, 'ASST', '助理', 15, 1),
(NOW(), NOW(), 1, 'INTERN', '实习生', 16, 1),
(NOW(), NOW(), 1, 'CONSULT', '顾问', 17, 1),
(NOW(), NOW(), 1, 'AUDIT', '审计', 18, 1),
(NOW(), NOW(), 1, 'SEC', '秘书', 19, 1),
(NOW(), NOW(), 1, 'DRV', '司机', 20, 1);

-- ============================================
-- 3. 用户数据 (20条)
-- ============================================
INSERT INTO users (created_at, updated_at, tenant_id, username, password, nickname, email, phone, dept_id, post_id, status) VALUES
(NOW(), NOW(), 1, 'zhangsan', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '张三', 'zhangsan@company.com', '13900000001', 1, 1, 1),
(NOW(), NOW(), 1, 'lisi', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '李四', 'lisi@company.com', '13900000002', 2, 2, 1),
(NOW(), NOW(), 1, 'wangwu', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '王五', 'wangwu@company.com', '13900000003', 3, 3, 1),
(NOW(), NOW(), 1, 'zhaoliu', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '赵六', 'zhaoliu@company.com', '13900000004', 4, 4, 1),
(NOW(), NOW(), 1, 'sunqi', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '孙七', 'sunqi@company.com', '13900000005', 5, 5, 1),
(NOW(), NOW(), 1, 'zhouba', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '周八', 'zhouba@company.com', '13900000006', 6, 6, 1),
(NOW(), NOW(), 1, 'wujiu', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '吴九', 'wujiu@company.com', '13900000007', 7, 7, 1),
(NOW(), NOW(), 1, 'zhengshi', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '郑十', 'zhengshi@company.com', '13900000008', 8, 8, 1),
(NOW(), NOW(), 1, 'chenyiyi', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '陈一一', 'chenyiyi@company.com', '13900000009', 9, 9, 1),
(NOW(), NOW(), 1, 'qianersh', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '钱二', 'qianersh@company.com', '13900000010', 10, 10, 1),
(NOW(), NOW(), 1, 'sunser', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '孙三', 'sunser@company.com', '13900000011', 11, 11, 1),
(NOW(), NOW(), 1, 'liusi', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '刘四', 'liusi@company.com', '13900000012', 12, 12, 1),
(NOW(), NOW(), 1, 'wuwu', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '吴五', 'wuwu@company.com', '13900000013', 13, 13, 1),
(NOW(), NOW(), 1, 'zhengliu', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '郑六', 'zhengliu@company.com', '13900000014', 14, 14, 1),
(NOW(), NOW(), 1, 'chendaping', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '陈七', 'chendaping@company.com', '13900000015', 15, 15, 1),
(NOW(), NOW(), 1, 'zhoubaba', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '周八', 'zhoubaba@company.com', '13900000016', 16, 16, 1),
(NOW(), NOW(), 1, 'wujiujiu', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '吴九', 'wujiujiu@company.com', '13900000017', 17, 17, 1),
(NOW(), NOW(), 1, 'zhengshi1', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '郑十', 'zhengshi1@company.com', '13900000018', 1, 18, 1),
(NOW(), NOW(), 1, 'chenershi', '$2a$10$N9qo8uLOickgx2ZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '陈十一', 'chenershi@company.com', '13900000019', 2, 19, 1),
(NOW(), NOW(), 1, 'qianshier', '$2a$10$N9qo8uLOickgx2ZZMRZoMye0X4tY1Q0x1Y1Y1Y1Y1Y1Y1Y1Y1Y1', '钱十二', 'qianshier@company.com', '13900000020', 3, 20, 1);

-- ============================================
-- 4. 物料数据 (20条)
-- ============================================
INSERT INTO mdm_material (created_at, updated_at, tenant_id, material_code, material_name, spec, unit, type, status) VALUES
(NOW(), NOW(), 1, 'MAT-001', '钢板', '1000*2000*2mm', '张', '原材料', 1),
(NOW(), NOW(), 1, 'MAT-002', '铝板', '1000*2000*1.5mm', '张', '原材料', 1),
(NOW(), NOW(), 1, 'MAT-003', '铜板', '500*1000*1mm', '张', '原材料', 1),
(NOW(), NOW(), 1, 'MAT-004', '塑料粒子A', 'ABS粒料', 'kg', '原材料', 1),
(NOW(), NOW(), 1, 'MAT-005', '塑料粒子B', 'PC粒料', 'kg', '原材料', 1),
(NOW(), NOW(), 1, 'MAT-006', '螺丝M3', '不锈钢', '个', '标准件', 1),
(NOW(), NOW(), 1, 'MAT-007', '螺丝M4', '不锈钢', '个', '标准件', 1),
(NOW(), NOW(), 1, 'MAT-008', '螺丝M5', '不锈钢', '个', '标准件', 1),
(NOW(), NOW(), 1, 'MAT-009', '轴承608', '内径8mm', '个', '标准件', 1),
(NOW(), NOW(), 1, 'MAT-010', '轴承6200', '内径10mm', '个', '标准件', 1),
(NOW(), NOW(), 1, 'MAT-011', '电机220V50W', '直流无刷', '个', '电气件', 1),
(NOW(), NOW(), 1, 'MAT-012', '电机220V100W', '直流无刷', '个', '电气件', 1),
(NOW(), NOW(), 1, 'MAT-013', 'PCB板A型', '4层板', '个', '电气件', 1),
(NOW(), NOW(), 1, 'MAT-014', 'PCB板B型', '6层板', '个', '电气件', 1),
(NOW(), NOW(), 1, 'MAT-015', '显示屏4寸', '480*320', '个', '电气件', 1),
(NOW(), NOW(), 1, 'MAT-016', '触摸屏5寸', '电容式', '个', '电气件', 1),
(NOW(), NOW(), 1, 'MAT-017', '电池18650', '3.7V 2000mAh', '个', '电气件', 1),
(NOW(), NOW(), 1, 'MAT-018', '电池26650', '3.2V 3000mAh', '个', '电气件', 1),
(NOW(), NOW(), 1, 'MAT-019', '线材AWG22', '22AWG 100米/卷', '卷', '电气件', 1),
(NOW(), NOW(), 1, 'MAT-020', '线材AWG18', '18AWG 100米/卷', '卷', '电气件', 1);

-- ============================================
-- 5. 车间数据 (20条)
-- ============================================
INSERT INTO mdm_workshop (created_at, updated_at, tenant_id, workshop_code, workshop_name, dept_id, area, status) VALUES
(NOW(), NOW(), 1, 'WS-001', '一车间-A线', 11, 500, 1),
(NOW(), NOW(), 1, 'WS-002', '一车间-B线', 11, 500, 1),
(NOW(), NOW(), 1, 'WS-003', '一车间-C线', 11, 500, 1),
(NOW(), NOW(), 1, 'WS-004', '二车间-A线', 12, 600, 1),
(NOW(), NOW(), 1, 'WS-005', '二车间-B线', 12, 600, 1),
(NOW(), NOW(), 1, 'WS-006', '三车间-A线', 13, 450, 1),
(NOW(), NOW(), 1, 'WS-007', '三车间-B线', 13, 450, 1),
(NOW(), NOW(), 1, 'WS-008', '装配车间-A区', 14, 800, 1),
(NOW(), NOW(), 1, 'WS-009', '装配车间-B区', 14, 800, 1),
(NOW(), NOW(), 1, 'WS-010', '装配车间-C区', 14, 800, 1),
(NOW(), NOW(), 1, 'WS-011', '测试车间', 15, 300, 1),
(NOW(), NOW(), 1, 'WS-012', '包装车间', 16, 400, 1),
(NOW(), NOW(), 1, 'WS-013', '模具车间', 17, 350, 1),
(NOW(), NOW(), 1, 'WS-014', '喷涂车间', 3, 400, 1),
(NOW(), NOW(), 1, 'WS-015', '焊接车间', 3, 300, 1),
(NOW(), NOW(), 1, 'WS-016', '电镀车间', 3, 350, 1),
(NOW(), NOW(), 1, 'WS-017', '热处理车间', 3, 250, 1),
(NOW(), NOW(), 1, 'WS-018', 'CNC加工车间', 3, 500, 1),
(NOW(), NOW(), 1, 'WS-019', '钣金车间', 3, 450, 1),
(NOW(), NOW(), 1, 'WS-020', '仓库作业区', 6, 1000, 1);

-- ============================================
-- 6. 生产线数据 (20条)
-- ============================================
INSERT INTO mdm_production_line (created_at, updated_at, tenant_id, line_code, line_name, workshop_id, type, capacity, status) VALUES
(NOW(), NOW(), 1, 'LINE-001', 'A线-1号', 1, '装配线', 100, 1),
(NOW(), NOW(), 1, 'LINE-002', 'A线-2号', 1, '装配线', 100, 1),
(NOW(), NOW(), 1, 'LINE-003', 'B线-1号', 2, '装配线', 120, 1),
(NOW(), NOW(), 1, 'LINE-004', 'B线-2号', 2, '装配线', 120, 1),
(NOW(), NOW(), 1, 'LINE-005', 'C线-1号', 3, '装配线', 80, 1),
(NOW(), NOW(), 1, 'LINE-006', 'D线-1号', 4, 'SMT线', 200, 1),
(NOW(), NOW(), 1, 'LINE-007', 'D线-2号', 4, 'SMT线', 200, 1),
(NOW(), NOW(), 1, 'LINE-008', 'E线-1号', 5, '测试线', 50, 1),
(NOW(), NOW(), 1, 'LINE-009', 'F线-1号', 6, '包装线', 150, 1),
(NOW(), NOW(), 1, 'LINE-010', 'G线-1号', 7, '焊接线', 60, 1),
(NOW(), NOW(), 1, 'LINE-011', 'H线-1号', 8, 'CNC线', 30, 1),
(NOW(), NOW(), 1, 'LINE-012', 'I线-1号', 9, '钣金线', 40, 1),
(NOW(), NOW(), 1, 'LINE-013', 'J线-1号', 10, '喷涂线', 35, 1),
(NOW(), NOW(), 1, 'LINE-014', 'K线-1号', 11, '电镀线', 25, 1),
(NOW(), NOW(), 1, 'LINE-015', 'L线-1号', 12, '组装线', 90, 1),
(NOW(), NOW(), 1, 'LINE-016', 'M线-1号', 13, '组装线', 90, 1),
(NOW(), NOW(), 1, 'LINE-017', 'N线-1号', 14, '老化线', 40, 1),
(NOW(), NOW(), 1, 'LINE-018', 'O线-1号', 15, '包装线', 100, 1),
(NOW(), NOW(), 1, 'LINE-019', 'P线-1号', 16, '成品线', 80, 1),
(NOW(), NOW(), 1, 'LINE-020', 'Q线-1号', 17, '检验线', 50, 1);

-- ============================================
-- 7. 销售订单数据 (20条)
-- ============================================
INSERT INTO pro_sales_order (created_at, updated_at, tenant_id, order_no, customer_name, contact, phone, order_date, delivery_date, total_amount, status) VALUES
(NOW(), NOW(), 1, 'SO-2024-0001', '华为技术有限公司', '张经理', '18600001001', NOW()- INTERVAL '30 days', NOW() + INTERVAL '15 days', 500000.00, 2),
(NOW(), NOW(), 1, 'SO-2024-0002', '小米科技有限公司', '李经理', '18600001002', NOW()- INTERVAL '28 days', NOW() + INTERVAL '12 days', 350000.00, 2),
(NOW(), NOW(), 1, 'SO-2024-0003', 'OPPO广东移动通信', '王经理', '18600001003', NOW()- INTERVAL '25 days', NOW() + INTERVAL '10 days', 420000.00, 2),
(NOW(), NOW(), 1, 'SO-2024-0004', 'VIVO移动通信', '刘经理', '18600001004', NOW()- INTERVAL '22 days', NOW() + INTERVAL '8 days', 380000.00, 1),
(NOW(), NOW(), 1, 'SO-2024-0005', '中兴通讯股份', '陈经理', '18600001005', NOW()- INTERVAL '20 days', NOW() + INTERVAL '20 days', 280000.00, 1),
(NOW(), NOW(), 1, 'SO-2024-0006', '联想集团', '赵经理', '18600001006', NOW()- INTERVAL '18 days', NOW() + INTERVAL '18 days', 520000.00, 1),
(NOW(), NOW(), 1, 'SO-2024-0007', 'Dell中国', '孙经理', '18600001007', NOW()- INTERVAL '15 days', NOW() + INTERVAL '25 days', 450000.00, 1),
(NOW(), NOW(), 1, 'SO-2024-0008', 'HP中国', '周经理', '18600001008', NOW()- INTERVAL '12 days', NOW() + INTERVAL '22 days', 320000.00, 1),
(NOW(), NOW(), 1, 'SO-2024-0009', 'Cisco中国', '吴经理', '18600001009', NOW()- INTERVAL '10 days', NOW() + INTERVAL '30 days', 680000.00, 1),
(NOW(), NOW(), 1, 'SO-2024-0010', '海康威视', '郑经理', '18600001010', NOW()- INTERVAL '8 days', NOW() + INTERVAL '7 days', 250000.00, 1),
(NOW(), NOW(), 1, 'SO-2024-0011', '大华技术', '冯经理', '18600001011', NOW()- INTERVAL '6 days', NOW() + INTERVAL '5 days', 220000.00, 1),
(NOW(), NOW(), 1, 'SO-2024-0012', '宇视科技', '于经理', '18600001012', NOW()- INTERVAL '5 days', NOW() + INTERVAL '3 days', 180000.00, 1),
(NOW(), NOW(), 1, 'SO-2024-0013', '比亚迪电子', '董经理', '18600001013', NOW()- INTERVAL '4 days', NOW() + INTERVAL '2 days', 420000.00, 1),
(NOW(), NOW(), 1, 'SO-2024-0014', '富士康科技', '潘经理', '18600001014', NOW()- INTERVAL '3 days', NOW() + INTERVAL '1 days', 750000.00, 1),
(NOW(), NOW(), 1, 'SO-2024-0015', '歌尔股份', '林经理', '18600001015', NOW()- INTERVAL '2 days', NOW() + INTERVAL '0 days', 310000.00, 1),
(NOW(), NOW(), 1, 'SO-2024-0016', '立讯精密', '柳经理', '18600001016', NOW()- INTERVAL '1 days', NOW() + INTERVAL '4 days', 280000.00, 1),
(NOW(), NOW(), 1, 'SO-2024-0017', '领益智造', '温经理', '18600001017', NOW(), NOW() + INTERVAL '6 days', 350000.00, 0),
(NOW(), NOW(), 1, 'SO-2024-0018', '华勤技术', '尤经理', '18600001018', NOW(), NOW() + INTERVAL '8 days', 420000.00, 0),
(NOW(), NOW(), 1, 'SO-2024-0019', '闻泰科技', '韦经理', '18600001019', NOW(), NOW() + INTERVAL '10 days', 380000.00, 0),
(NOW(), NOW(), 1, 'SO-2024-0020', '信维通信', '昌经理', '18600001020', NOW(), NOW() + INTERVAL '12 days', 290000.00, 0);

-- ============================================
-- 8. MPS主生产计划数据 (20条)
-- ============================================
INSERT INTO mps (created_at, updated_at, tenant_id, plan_no, sales_order_no, product_name, quantity, start_date, end_date, workshop_id, line_id, status) VALUES
(NOW(), NOW(), 1, 'MPS-2024-0001', 'SO-2024-0001', '智能音箱A1', 1000, NOW(), NOW() + INTERVAL '10 days', 1, 1, 2),
(NOW(), NOW(), 1, 'MPS-2024-0002', 'SO-2024-0002', '智能音箱A2', 800, NOW(), NOW() + INTERVAL '8 days', 2, 3, 2),
(NOW(), NOW(), 1, 'MPS-2024-0003', 'SO-2024-0003', '智能手表B1', 1200, NOW(), NOW() + INTERVAL '12 days', 3, 5, 2),
(NOW(), NOW(), 1, 'MPS-2024-0004', 'SO-2024-0004', '智能手表B2', 900, NOW(), NOW() + INTERVAL '10 days', 4, 6, 1),
(NOW(), NOW(), 1, 'MPS-2024-0005', 'SO-2024-0005', '蓝牙耳机C1', 2000, NOW(), NOW() + INTERVAL '7 days', 5, 8, 1),
(NOW(), NOW(), 1, 'MPS-2024-0006', 'SO-2024-0006', '蓝牙耳机C2', 1500, NOW(), NOW() + INTERVAL '6 days', 6, 10, 1),
(NOW(), NOW(), 1, 'MPS-2024-0007', 'SO-2024-0007', '路由器D1', 600, NOW(), NOW() + INTERVAL '15 days', 7, 12, 1),
(NOW(), NOW(), 1, 'MPS-2024-0008', 'SO-2024-0008', '路由器D2', 500, NOW(), NOW() + INTERVAL '14 days', 8, 15, 1),
(NOW(), NOW(), 1, 'MPS-2024-0009', 'SO-2024-0009', '摄像头E1', 800, NOW(), NOW() + INTERVAL '9 days', 9, 17, 1),
(NOW(), NOW(), 1, 'MPS-2024-0010', 'SO-2024-0010', '摄像头E2', 700, NOW(), NOW() + INTERVAL '8 days', 10, 1, 1),
(NOW(), NOW(), 1, 'MPS-2024-0011', 'SO-2024-0011', '平板电脑F1', 400, NOW(), NOW() + INTERVAL '20 days', 11, 2, 0),
(NOW(), NOW(), 1, 'MPS-2024-0012', 'SO-2024-0012', '平板电脑F2', 350, NOW(), NOW() + INTERVAL '18 days', 12, 4, 0),
(NOW(), NOW(), 1, 'MPS-2024-0013', 'SO-2024-0013', '笔记本G1', 200, NOW(), NOW() + INTERVAL '25 days', 13, 6, 0),
(NOW(), NOW(), 1, 'MPS-2024-0014', 'SO-2024-0014', '笔记本G2', 180, NOW(), NOW() + INTERVAL '22 days', 14, 7, 0),
(NOW(), NOW(), 1, 'MPS-2024-0015', 'SO-2024-0015', '手机H1', 3000, NOW(), NOW() + INTERVAL '30 days', 1, 1, 0),
(NOW(), NOW(), 1, 'MPS-2024-0016', 'SO-2024-0016', '手机H2', 2500, NOW(), NOW() + INTERVAL '28 days', 2, 3, 0),
(NOW(), NOW(), 1, 'MPS-2024-0017', 'SO-2024-0017', '手环I1', 5000, NOW(), NOW() + INTERVAL '14 days', 3, 5, 0),
(NOW(), NOW(), 1, 'MPS-2024-0018', 'SO-2024-0018', '手环I2', 4500, NOW(), NOW() + INTERVAL '12 days', 4, 6, 0),
(NOW(), NOW(), 1, 'MPS-2024-0019', 'SO-2024-0019', '台灯J1', 2000, NOW(), NOW() + INTERVAL '10 days', 5, 8, 0),
(NOW(), NOW(), 1, 'MPS-2024-0020', 'SO-2024-0020', '台灯J2', 1800, NOW(), NOW() + INTERVAL '8 days', 6, 10, 0);

-- ============================================
-- 9. 设备台账数据 (20条)
-- ============================================
INSERT INTO eqp_equipment (created_at, updated_at, tenant_id, equipment_code, equipment_name, model, manufacturer, purchase_date, workshop_id, line_id, status) VALUES
(NOW(), NOW(), 1, 'EQP-001', '贴片机SM471', 'SM471+', '三星', '2023-01-15', 1, 1, 1),
(NOW(), NOW(), 1, 'EQP-002', '贴片机SM482', 'SM482', '三星', '2023-02-20', 1, 2, 1),
(NOW(), NOW(), 1, 'EQP-003', '回流焊IR10', 'IR10', ' Heller', '2023-03-10', 2, 3, 1),
(NOW(), NOW(), 1, 'EQP-004', '回流焊BTU', 'VIP-70N', 'BTU', '2023-04-15', 2, 4, 1),
(NOW(), NOW(), 1, 'EQP-005', 'AOI检测仪', 'SP-500', '韩华', '2023-05-20', 3, 5, 1),
(NOW(), NOW(), 1, 'EQP-006', 'AOI检测仪', 'VT-808N2', '德律', '2023-06-10', 3, 6, 1),
(NOW(), NOW(), 1, 'EQP-007', 'SPI锡膏检测', 'S3000', 'Koh Young', '2023-07-15', 4, 7, 1),
(NOW(), NOW(), 1, 'EQP-008', 'SPI锡膏检测', 'Proton', 'Technotron', '2023-08-20', 4, 8, 1),
(NOW(), NOW(), 1, 'EQP-009', 'CNC加工中心', 'VM-500', '斗山', '2022-09-15', 5, 11, 1),
(NOW(), NOW(), 1, 'EQP-010', 'CNC加工中心', 'TV-500', '斗山', '2022-10-20', 5, 12, 1),
(NOW(), NOW(), 1, 'EQP-011', '激光切割机', 'LCM-3015', '宏山', '2022-11-10', 6, 13, 1),
(NOW(), NOW(), 1, 'EQP-012', '激光切割机', 'HTS4020', '海目星', '2022-12-15', 6, 14, 1),
(NOW(), NOW(), 1, 'EQP-013', '注塑机', 'JSW-350', '日精', '2023-01-20', 7, 15, 1),
(NOW(), NOW(), 1, 'EQP-014', '注塑机', 'JSW-450', '日精', '2023-02-25', 7, 16, 1),
(NOW(), NOW(), 1, 'EQP-015', '冲床', 'TPS-200', '金丰', '2023-03-10', 8, 17, 1),
(NOW(), NOW(), 1, 'EQP-016', '冲床', 'TPS-300', '金丰', '2023-04-15', 8, 18, 1),
(NOW(), NOW(), 1, 'EQP-017', '老化测试柜', 'AGT-1000', '中颖', '2023-05-20', 9, 19, 1),
(NOW(), NOW(), 1, 'EQP-018', '老化测试柜', 'AGT-2000', '中颖', '2023-06-25', 9, 20, 1),
(NOW(), NOW(), 1, 'EQP-019', '包装机', 'PVS-500', '松川', '2023-07-30', 10, 1, 1),
(NOW(), NOW(), 1, 'EQP-020', '包装机', 'PVS-800', '松川', '2023-08-10', 10, 2, 1);

-- ============================================
-- 10. 仓库数据 (20条)
-- ============================================
INSERT INTO wms_warehouse (created_at, updated_at, tenant_id, warehouse_code, warehouse_name, type, area, address, status) VALUES
(NOW(), NOW(), 1, 'WH-001', '原材料仓库A区', '原材料', 2000, 'A栋1楼', 1),
(NOW(), NOW(), 1, 'WH-002', '原材料仓库B区', '原材料', 1800, 'A栋2楼', 1),
(NOW(), NOW(), 1, 'WH-003', '原材料仓库C区', '原材料', 1500, 'A栋3楼', 1),
(NOW(), NOW(), 1, 'WH-004', '半成品仓库A区', '半成品', 1200, 'B栋1楼', 1),
(NOW(), NOW(), 1, 'WH-005', '半成品仓库B区', '半成品', 1000, 'B栋2楼', 1),
(NOW(), NOW(), 1, 'WH-006', '成品仓库A区', '成品', 2500, 'C栋1楼', 1),
(NOW(), NOW(), 1, 'WH-007', '成品仓库B区', '成品', 2200, 'C栋2楼', 1),
(NOW(), NOW(), 1, 'WH-008', '成品仓库C区', '成品', 2000, 'C栋3楼', 1),
(NOW(), NOW(), 1, 'WH-009', '包材仓库', '包材', 800, 'D栋1楼', 1),
(NOW(), NOW(), 1, 'WH-010', '危险品仓库', '危险品', 300, 'E栋1楼', 1),
(NOW(), NOW(), 1, 'WH-011', '退货仓库', '退货', 400, 'F栋1楼', 1),
(NOW(), NOW(), 1, 'WH-012', '待检仓库', '待检', 500, 'G栋1楼', 1),
(NOW(), NOW(), 1, 'WH-013', '不良品仓库', '不良品', 200, 'H栋1楼', 1),
(NOW(), NOW(), 1, 'WH-014', '工装仓库', '工装', 350, 'I栋1楼', 1),
(NOW(), NOW(), 1, 'WH-015', '模具仓库', '模具', 400, 'J栋1楼', 1),
(NOW(), NOW(), 1, 'WH-016', '备件仓库', '备件', 300, 'K栋1楼', 1),
(NOW(), NOW(), 1, 'WH-017', '办公用品仓库', '办公', 150, 'L栋1楼', 1),
(NOW(), NOW(), 1, 'WH-018', '劳保用品仓库', '劳保', 120, 'M栋1楼', 1),
(NOW(), NOW(), 1, 'WH-019', '在制品暂存区', '在制品', 600, 'N栋1楼', 1),
(NOW(), NOW(), 1, 'WH-020', '出货暂存区', '出货', 800, 'O栋1楼', 1);

-- ============================================
-- 11. 库位数据 (20条)
-- ============================================
INSERT INTO wms_location (created_at, updated_at, tenant_id, location_code, location_name, warehouse_id, zone, row, column, status) VALUES
(NOW(), NOW(), 1, 'LOC-A01-01-01', 'A01区01排01列', 1, 'A', 1, 1, 1),
(NOW(), NOW(), 1, 'LOC-A01-01-02', 'A01区01排02列', 1, 'A', 1, 2, 1),
(NOW(), NOW(), 1, 'LOC-A01-02-01', 'A01区02排01列', 1, 'A', 2, 1, 1),
(NOW(), NOW(), 1, 'LOC-A01-02-02', 'A01区02排02列', 1, 'A', 2, 2, 1),
(NOW(), NOW(), 1, 'LOC-B01-01-01', 'B01区01排01列', 2, 'B', 1, 1, 1),
(NOW(), NOW(), 1, 'LOC-B01-01-02', 'B01区01排02列', 2, 'B', 1, 2, 1),
(NOW(), NOW(), 1, 'LOC-B01-02-01', 'B01区02排01列', 2, 'B', 2, 1, 1),
(NOW(), NOW(), 1, 'LOC-B01-02-02', 'B01区02排02列', 2, 'B', 2, 2, 1),
(NOW(), NOW(), 1, 'LOC-C01-01-01', 'C01区01排01列', 3, 'C', 1, 1, 1),
(NOW(), NOW(), 1, 'LOC-C01-01-02', 'C01区01排02列', 3, 'C', 1, 2, 1),
(NOW(), NOW(), 1, 'LOC-C01-02-01', 'C01区02排01列', 3, 'C', 2, 1, 1),
(NOW(), NOW(), 1, 'LOC-C01-02-02', 'C01区02排02列', 3, 'C', 2, 2, 1),
(NOW(), NOW(), 1, 'LOC-D01-01-01', 'D01区01排01列', 4, 'D', 1, 1, 1),
(NOW(), NOW(), 1, 'LOC-D01-01-02', 'D01区01排02列', 4, 'D', 1, 2, 1),
(NOW(), NOW(), 1, 'LOC-D01-02-01', 'D01区02排01列', 4, 'D', 2, 1, 1),
(NOW(), NOW(), 1, 'LOC-D01-02-02', 'D01区02排02列', 4, 'D', 2, 2, 1),
(NOW(), NOW(), 1, 'LOC-E01-01-01', 'E01区01排01列', 5, 'E', 1, 1, 1),
(NOW(), NOW(), 1, 'LOC-E01-01-02', 'E01区01排02列', 5, 'E', 1, 2, 1),
(NOW(), NOW(), 1, 'LOC-E01-02-01', 'E01区02排01列', 5, 'E', 2, 1, 1),
(NOW(), NOW(), 1, 'LOC-E01-02-02', 'E01区02排02列', 5, 'E', 2, 2, 1);

-- ============================================
-- 12. 库存数据 (20条)
-- ============================================
INSERT INTO wms_inventory (created_at, updated_at, tenant_id, material_id, location_id, batch_no, quantity, unit, status) VALUES
(NOW(), NOW(), 1, 1, 1, 'BATCH-2024-001', 5000, '张', 1),
(NOW(), NOW(), 1, 2, 2, 'BATCH-2024-002', 3000, '张', 1),
(NOW(), NOW(), 1, 3, 3, 'BATCH-2024-003', 2000, '张', 1),
(NOW(), NOW(), 1, 4, 4, 'BATCH-2024-004', 10000, 'kg', 1),
(NOW(), NOW(), 1, 5, 5, 'BATCH-2024-005', 8000, 'kg', 1),
(NOW(), NOW(), 1, 6, 6, 'BATCH-2024-006', 50000, '个', 1),
(NOW(), NOW(), 1, 7, 7, 'BATCH-2024-007', 40000, '个', 1),
(NOW(), NOW(), 1, 8, 8, 'BATCH-2024-008', 30000, '个', 1),
(NOW(), NOW(), 1, 9, 9, 'BATCH-2024-009', 20000, '个', 1),
(NOW(), NOW(), 1, 10, 10, 'BATCH-2024-010', 15000, '个', 1),
(NOW(), NOW(), 1, 11, 11, 'BATCH-2024-011', 5000, '个', 1),
(NOW(), NOW(), 1, 12, 12, 'BATCH-2024-012', 4000, '个', 1),
(NOW(), NOW(), 1, 13, 13, 'BATCH-2024-013', 3000, '个', 1),
(NOW(), NOW(), 1, 14, 14, 'BATCH-2024-014', 2000, '个', 1),
(NOW(), NOW(), 1, 15, 15, 'BATCH-2024-015', 1000, '个', 1),
(NOW(), NOW(), 1, 16, 16, 'BATCH-2024-016', 800, '个', 1),
(NOW(), NOW(), 1, 17, 17, 'BATCH-2024-017', 20000, '个', 1),
(NOW(), NOW(), 1, 18, 18, 'BATCH-2024-018', 15000, '个', 1),
(NOW(), NOW(), 1, 19, 19, 'BATCH-2024-019', 500, '卷', 1),
(NOW(), NOW(), 1, 20, 20, 'BATCH-2024-020', 400, '卷', 1);

-- ============================================
-- 13. 能源记录数据 (20条)
-- ============================================
INSERT INTO energy_records (created_at, tenant_id, workshop_id, line_id, record_date, electricity, water, gas, status) VALUES
(NOW(), 1, 1, 1, NOW() - INTERVAL '1 day', 1500.5, 200.3, 50.2, 1),
(NOW(), 1, 2, 3, NOW() - INTERVAL '1 day', 1400.3, 180.5, 45.1, 1),
(NOW(), 1, 3, 5, NOW() - INTERVAL '1 day', 1200.8, 160.2, 40.5, 1),
(NOW(), 1, 4, 6, NOW() - INTERVAL '2 days', 1100.2, 150.8, 35.3, 1),
(NOW(), 1, 5, 8, NOW() - INTERVAL '2 days', 1000.5, 140.6, 30.2, 1),
(NOW(), 1, 6, 10, NOW() - INTERVAL '3 days', 900.3, 130.4, 28.5, 1),
(NOW(), 1, 7, 12, NOW() - INTERVAL '3 days', 850.8, 120.3, 25.1, 1),
(NOW(), 1, 8, 15, NOW() - INTERVAL '4 days', 780.5, 110.7, 22.3, 1),
(NOW(), 1, 9, 17, NOW() - INTERVAL '4 days', 700.2, 100.5, 20.8, 1),
(NOW(), 1, 10, 1, NOW() - INTERVAL '5 days', 650.8, 95.3, 18.5, 1),
(NOW(), 1, 11, 2, NOW() - INTERVAL '5 days', 600.3, 88.6, 15.2, 1),
(NOW(), 1, 12, 4, NOW() - INTERVAL '6 days', 550.5, 80.4, 12.8, 1),
(NOW(), 1, 13, 6, NOW() - INTERVAL '6 days', 500.2, 75.5, 10.5, 1),
(NOW(), 1, 14, 7, NOW() - INTERVAL '7 days', 450.8, 68.3, 8.2, 1),
(NOW(), 1, 15, 8, NOW() - INTERVAL '7 days', 400.3, 60.5, 5.8, 1),
(NOW(), 1, 16, 9, NOW() - INTERVAL '8 days', 350.5, 55.2, 3.5, 1),
(NOW(), 1, 17, 10, NOW() - INTERVAL '8 days', 300.8, 48.6, 2.2, 1),
(NOW(), 1, 1, 1, NOW() - INTERVAL '9 days', 1450.2, 195.3, 48.5, 1),
(NOW(), 1, 2, 3, NOW() - INTERVAL '9 days', 1350.5, 175.8, 42.3, 1),
(NOW(), 1, 3, 5, NOW() - INTERVAL '10 days', 1150.3, 155.6, 38.2, 1);

-- ============================================
-- 14.andon呼叫记录 (20条)
-- ============================================
INSERT INTO andon_call (created_at, updated_at, tenant_id, call_no, workshop_id, line_id, station, type, level, status) VALUES
(NOW(), NOW(), 1, 'ANDON-001', 1, 1, 'A工位', '质量问题', 2, 1),
(NOW(), NOW(), 1, 'ANDON-002', 2, 3, 'B工位', '物料短缺', 3, 1),
(NOW(), NOW(), 1, 'ANDON-003', 3, 5, 'C工位', '设备故障', 1, 2),
(NOW(), NOW(), 1, 'ANDON-004', 4, 6, 'D工位', '质量问题', 2, 1),
(NOW(), NOW(), 1, 'ANDON-005', 5, 8, 'E工位', '其他', 3, 1),
(NOW(), NOW(), 1, 'ANDON-006', 6, 10, 'F工位', '物料短缺', 2, 1),
(NOW(), NOW(), 1, 'ANDON-007', 7, 12, 'G工位', '设备故障', 1, 2),
(NOW(), NOW(), 1, 'ANDON-008', 8, 15, 'H工位', '质量问题', 2, 1),
(NOW(), NOW(), 1, 'ANDON-009', 9, 17, 'I工位', '其他', 3, 1),
(NOW(), NOW(), 1, 'ANDON-010', 10, 1, 'J工位', '物料短缺', 2, 1),
(NOW(), NOW(), 1, 'ANDON-011', 11, 2, 'K工位', '质量问题', 2, 1),
(NOW(), NOW(), 1, 'ANDON-012', 12, 4, 'L工位', '设备故障', 1, 1),
(NOW(), NOW(), 1, 'ANDON-013', 13, 6, 'M工位', '其他', 3, 1),
(NOW(), NOW(), 1, 'ANDON-014', 14, 7, 'N工位', '物料短缺', 2, 1),
(NOW(), NOW(), 1, 'ANDON-015', 15, 8, 'O工位', '质量问题', 2, 1),
(NOW(), NOW(), 1, 'ANDON-016', 1, 1, 'P工位', '设备故障', 1, 1),
(NOW(), NOW(), 1, 'ANDON-017', 2, 3, 'Q工位', '其他', 3, 1),
(NOW(), NOW(), 1, 'ANDON-018', 3, 5, 'R工位', '物料短缺', 2, 1),
(NOW(), NOW(), 1, 'ANDON-019', 4, 6, 'S工位', '质量问题', 2, 1),
(NOW(), NOW(), 1, 'ANDON-020', 5, 8, 'T工位', '设备故障', 1, 1);

-- ============================================
-- 15. 追溯记录 (20条)
-- ============================================
INSERT INTO trace_serial_number (created_at, updated_at, tenant_id, serial_no, product_code, batch_no, workshop_id, line_id, worker_id, status) VALUES
(NOW(), NOW(), 1, 'SN-2024-000001', 'PRD-A1-001', 'BATCH-2024-001', 1, 1, 1, 1),
(NOW(), NOW(), 1, 'SN-2024-000002', 'PRD-A1-002', 'BATCH-2024-001', 1, 1, 2, 1),
(NOW(), NOW(), 1, 'SN-2024-000003', 'PRD-A2-001', 'BATCH-2024-002', 2, 3, 3, 1),
(NOW(), NOW(), 1, 'SN-2024-000004', 'PRD-A2-002', 'BATCH-2024-002', 2, 3, 4, 1),
(NOW(), NOW(), 1, 'SN-2024-000005', 'PRD-B1-001', 'BATCH-2024-003', 3, 5, 5, 1),
(NOW(), NOW(), 1, 'SN-2024-000006', 'PRD-B1-002', 'BATCH-2024-003', 3, 5, 6, 1),
(NOW(), NOW(), 1, 'SN-2024-000007', 'PRD-B2-001', 'BATCH-2024-004', 4, 6, 7, 1),
(NOW(), NOW(), 1, 'SN-2024-000008', 'PRD-B2-002', 'BATCH-2024-004', 4, 6, 8, 1),
(NOW(), NOW(), 1, 'SN-2024-000009', 'PRD-C1-001', 'BATCH-2024-005', 5, 8, 9, 1),
(NOW(), NOW(), 1, 'SN-2024-000010', 'PRD-C1-002', 'BATCH-2024-005', 5, 8, 10, 1),
(NOW(), NOW(), 1, 'SN-2024-000011', 'PRD-C2-001', 'BATCH-2024-006', 6, 10, 11, 1),
(NOW(), NOW(), 1, 'SN-2024-000012', 'PRD-C2-002', 'BATCH-2024-006', 6, 10, 12, 1),
(NOW(), NOW(), 1, 'SN-2024-000013', 'PRD-D1-001', 'BATCH-2024-007', 7, 12, 13, 1),
(NOW(), NOW(), 1, 'SN-2024-000014', 'PRD-D1-002', 'BATCH-2024-007', 7, 12, 14, 1),
(NOW(), NOW(), 1, 'SN-2024-000015', 'PRD-D2-001', 'BATCH-2024-008', 8, 15, 15, 1),
(NOW(), NOW(), 1, 'SN-2024-000016', 'PRD-D2-002', 'BATCH-2024-008', 8, 15, 16, 1),
(NOW(), NOW(), 1, 'SN-2024-000017', 'PRD-E1-001', 'BATCH-2024-009', 9, 17, 17, 1),
(NOW(), NOW(), 1, 'SN-2024-000018', 'PRD-E1-002', 'BATCH-2024-009', 9, 17, 18, 1),
(NOW(), NOW(), 1, 'SN-2024-000019', 'PRD-E2-001', 'BATCH-2024-010', 10, 1, 19, 1),
(NOW(), NOW(), 1, 'SN-2024-000020', 'PRD-E2-002', 'BATCH-2024-010', 10, 1, 20, 1);

-- ============================================
-- 16. 字典数据 (20条)
-- ============================================
INSERT INTO dict_data (created_at, updated_at, tenant_id, dict_type, dict_label, dict_value, sort, status) VALUES
(NOW(), NOW(), 1, 'sys_user_sex', '男', '0', 1, 1),
(NOW(), NOW(), 1, 'sys_user_sex', '女', '1', 2, 1),
(NOW(), NOW(), 1, 'sys_user_sex', '未知', '2', 3, 1),
(NOW(), NOW(), 1, 'sys_normal_disable', '正常', '0', 1, 1),
(NOW(), NOW(), 1, 'sys_normal_disable', '停用', '1', 2, 1),
(NOW(), NOW(), 1, 'sys_customer_type', '终端客户', '1', 1, 1),
(NOW(), NOW(), 1, 'sys_customer_type', '经销商', '2', 2, 1),
(NOW(), NOW(), 1, 'sys_customer_type', '代理商', '3', 3, 1),
(NOW(), NOW(), 1, 'sys_order_status', '待确认', '0', 1, 1),
(NOW(), NOW(), 1, 'sys_order_status', '已确认', '1', 2, 1),
(NOW(), NOW(), 1, 'sys_order_status', '生产中', '2', 3, 1),
(NOW(), NOW(), 1, 'sys_order_status', '已完成', '3', 4, 1),
(NOW(), NOW(), 1, 'sys_order_status', '已取消', '4', 5, 1),
(NOW(), NOW(), 1, 'sys_quality_level', 'A级', 'A', 1, 1),
(NOW(), NOW(), 1, 'sys_quality_level', 'B级', 'B', 2, 1),
(NOW(), NOW(), 1, 'sys_quality_level', 'C级', 'C', 3, 1),
(NOW(), NOW(), 1, 'sys_quality_level', '不合格', 'D', 4, 1),
(NOW(), NOW(), 1, 'sys_andon_level', '一级-紧急', '1', 1, 1),
(NOW(), NOW(), 1, 'sys_andon_level', '二级-重要', '2', 2, 1),
(NOW(), NOW(), 1, 'sys_andon_level', '三级-一般', '3', 3, 1);

-- ============================================
-- 17. 班次数据 (20条)
-- ============================================
INSERT INTO mdm_shift (created_at, updated_at, tenant_id, shift_code, shift_name, start_time, end_time, duration_hours, status) VALUES
(NOW(), NOW(), 1, 'SHIFT-A1', 'A班早班', '08:00', '16:00', 8, 1),
(NOW(), NOW(), 1, 'SHIFT-A2', 'A班中班', '16:00', '00:00', 8, 1),
(NOW(), NOW(), 1, 'SHIFT-A3', 'A班晚班', '00:00', '08:00', 8, 1),
(NOW(), NOW(), 1, 'SHIFT-B1', 'B班早班', '07:00', '15:00', 8, 1),
(NOW(), NOW(), 1, 'SHIFT-B2', 'B班中班', '15:00', '23:00', 8, 1),
(NOW(), NOW(), 1, 'SHIFT-B3', 'B班晚班', '23:00', '07:00', 8, 1),
(NOW(), NOW(), 1, 'SHIFT-C1', 'C班早班', '06:00', '14:00', 8, 1),
(NOW(), NOW(), 1, 'SHIFT-C2', 'C班中班', '14:00', '22:00', 8, 1),
(NOW(), NOW(), 1, 'SHIFT-C3', 'C班晚班', '22:00', '06:00', 8, 1),
(NOW(), NOW(), 1, 'SHIFT-D1', 'D班早班', '09:00', '17:00', 8, 1),
(NOW(), NOW(), 1, 'SHIFT-D2', 'D班中班', '17:00', '01:00', 8, 1),
(NOW(), NOW(), 1, 'SHIFT-D3', 'D班晚班', '01:00', '09:00', 8, 1),
(NOW(), NOW(), 1, 'SHIFT-E1', '白班', '08:30', '17:30', 9, 1),
(NOW(), NOW(), 1, 'SHIFT-E2', '中班', '17:30', '02:30', 9, 1),
(NOW(), NOW(), 1, 'SHIFT-E3', '夜班', '02:30', '11:30', 9, 1),
(NOW(), NOW(), 1, 'SHIFT-F1', '早班2H', '08:00', '10:00', 2, 1),
(NOW(), NOW(), 1, 'SHIFT-F2', '中班2H', '14:00', '16:00', 2, 1),
(NOW(), NOW(), 1, 'SHIFT-F3', '晚班2H', '20:00', '22:00', 2, 1),
(NOW(), NOW(), 1, 'SHIFT-G1', '加班班', '18:00', '22:00', 4, 1),
(NOW(), NOW(), 1, 'SHIFT-G2', '周末班', '09:00', '18:00', 9, 1);

COMMIT;
