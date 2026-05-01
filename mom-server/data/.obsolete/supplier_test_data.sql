-- MOM3.0 供应商测试数据
-- 在 PostgreSQL 中执行: psql -U postgres -d mom3_db -f supplier_test_data.sql

INSERT INTO mdm_supplier (tenant_id, code, name, type, contact, phone, email, address, category, level, status, created_at, updated_at) VALUES
(1, 'SUP-TEST01', '华东精密机械有限公司', '原材料', '张经理', '13812345601', 'zhang@example.com', '上海市浦东新区', '金属材料', 1, 1, NOW(), NOW()),
(1, 'SUP-TEST02', '深圳创新电子科技有限公司', '原材料', '李总监', '13812345602', 'li@example.com', '深圳市南山区', '电子元器件', 1, 1, NOW(), NOW()),
(1, 'SUP-TEST03', '苏州工业园区华鑫塑胶制品厂', '辅料', '王厂长', '13812345603', 'wang@example.com', '苏州市工业园区', '塑料粒子', 2, 1, NOW(), NOW()),
(1, 'SUP-TEST04', '杭州中达金属表面处理有限公司', '服务', '陈工', '13812345604', 'chen@example.com', '杭州市余杭区', '表面处理', 2, 1, NOW(), NOW()),
(1, 'SUP-TEST05', '广州岭南包装材料有限公司', '辅料', '刘经理', '13812345605', 'liu@example.com', '广州市白云区', '包装材料', 3, 1, NOW(), NOW()),
(1, 'SUP-TEST06', '北京中关村软件股份有限公司', '设备', '赵总', '13812345606', 'zhao@example.com', '北京市海淀区', '软件系统', 1, 1, NOW(), NOW()),
(1, 'SUP-TEST07', '天津港保税区润通物流有限公司', '服务', '孙经理', '13812345607', 'sun@example.com', '天津港保税区', '物流运输', 2, 1, NOW(), NOW()),
(1, 'SUP-TEST08', '成都西南航空材料有限公司', '原材料', '周经理', '13812345608', 'zhou@example.com', '成都市双流区', '复合材料', 1, 1, NOW(), NOW()),
(1, 'SUP-TEST09', '武汉光谷激光设备有限公司', '设备', '吴工', '13812345609', 'wu@example.com', '武汉市东湖新区', '激光设备', 2, 1, NOW(), NOW()),
(1, 'SUP-TEST10', '南京金陵化工有限公司', '原材料', '郑经理', '13812345610', 'zheng@example.com', '南京市化学工业园区', '化工原料', 2, 1, NOW(), NOW()),
(1, 'SUP-TEST11', '西安古都标准件有限公司', '辅料', '冯经理', '13812345611', 'feng@example.com', '西安市经开区', '标准件', 3, 1, NOW(), NOW()),
(1, 'SUP-TEST12', '青岛海尔智造装备有限公司', '设备', '卫总', '13812345612', 'wei@example.com', '青岛市黄岛区', '生产设备', 1, 1, NOW(), NOW()),
(1, 'SUP-TEST13', '厦门特区进出口贸易有限公司', '原材料', '蒋经理', '13812345613', 'jiang@example.com', '厦门市湖里区', '进口材料', 2, 1, NOW(), NOW()),
(1, 'SUP-TEST14', '长沙中沙电缆桥架有限公司', '辅料', '沈经理', '13812345614', 'shen@example.com', '长沙市雨花区', '电缆桥架', 3, 1, NOW(), NOW()),
(1, 'SUP-TEST15', '东莞松山湖智能装备有限公司', '设备', '楚经理', '13812345615', 'chu@example.com', '东莞市松山湖', '智能装备', 1, 1, NOW(), NOW());
