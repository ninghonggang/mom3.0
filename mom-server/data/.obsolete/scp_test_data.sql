-- MOM3.0 供应链管理测试数据
-- 在 PostgreSQL 中执行: psql -U postgres -d mom3_db -f scp_test_data.sql

-- =====================================================
-- 采购订单 (Purchase Order) - 10条
-- =====================================================
INSERT INTO scp_purchase_order (tenant_id, po_no, po_type, supplier_id, supplier_code, supplier_name, contact_person, contact_phone, order_date, promised_date, currency, payment_terms, tax_rate, total_amount, total_qty, approval_status, status, source_type, created_at, updated_at) VALUES
(1, 'PO-2026-04001', 'STANDARD', 1, 'SUP-TEST01', '华东精密机械有限公司', '张经理', '13812345601', '2026-04-01', '2026-04-15', 'CNY', '月结30天', 13.00, 125500.00, 100, 'APPROVED', 'ISSUED', 'MANUAL', NOW(), NOW()),
(1, 'PO-2026-04002', 'STANDARD', 2, 'SUP-TEST02', '深圳创新电子科技有限公司', '李总监', '13812345602', '2026-04-02', '2026-04-16', 'CNY', '月结30天', 13.00, 89000.00, 200, 'APPROVED', 'ISSUED', 'MANUAL', NOW(), NOW()),
(1, 'PO-2026-04003', 'URGENT', 3, 'SUP-TEST03', '苏州工业园区华鑫塑胶制品厂', '王厂长', '13812345603', '2026-04-03', '2026-04-12', 'CNY', '月结30天', 13.00, 45600.00, 500, 'APPROVED', 'PARTIAL', 'MANUAL', NOW(), NOW()),
(1, 'PO-2026-04004', 'STANDARD', 4, 'SUP-TEST04', '杭州中达金属表面处理有限公司', '陈工', '13812345604', '2026-04-04', '2026-04-18', 'CNY', '月结30天', 13.00, 32000.00, 50, 'APPROVED', 'ISSUED', 'MANUAL', NOW(), NOW()),
(1, 'PO-2026-04005', 'STANDARD', 5, 'SUP-TEST05', '广州岭南包装材料有限公司', '刘经理', '13812345605', '2026-04-05', '2026-04-20', 'CNY', '月结45天', 13.00, 28500.00, 300, 'PENDING', 'PENDING', 'MANUAL', NOW(), NOW()),
(1, 'PO-2026-04006', 'LONG_TERM', 6, 'SUP-TEST06', '北京中关村软件股份有限公司', '赵总', '13812345606', '2026-04-06', '2026-05-06', 'CNY', '月结60天', 13.00, 150000.00, 10, 'APPROVED', 'ISSUED', 'MANUAL', NOW(), NOW()),
(1, 'PO-2026-04007', 'STANDARD', 7, 'SUP-TEST07', '天津港保税区润通物流有限公司', '孙经理', '13812345607', '2026-04-07', '2026-04-22', 'CNY', '月结30天', 13.00, 18000.00, 20, 'APPROVED', 'ISSUED', 'MANUAL', NOW(), NOW()),
(1, 'PO-2026-04008', 'STANDARD', 8, 'SUP-TEST08', '成都西南航空材料有限公司', '周经理', '13812345608', '2026-04-08', '2026-04-23', 'CNY', '月结30天', 13.00, 67500.00, 80, 'PENDING', 'PENDING', 'MANUAL', NOW(), NOW()),
(1, 'PO-2026-04009', 'URGENT', 9, 'SUP-TEST09', '武汉光谷激光设备有限公司', '吴工', '13812345609', '2026-04-09', '2026-04-14', 'CNY', '月结30天', 13.00, 98000.00, 5, 'APPROVED', 'PARTIAL', 'MANUAL', NOW(), NOW()),
(1, 'PO-2026-04010', 'STANDARD', 10, 'SUP-TEST10', '南京金陵化工有限公司', '郑经理', '13812345610', '2026-04-10', '2026-04-25', 'CNY', '月结30天', 13.00, 54200.00, 150, 'APPROVED', 'ISSUED', 'MANUAL', NOW(), NOW());

-- 采购订单明细
INSERT INTO scp_purchase_order_item (po_id, line_no, material_code, material_name, specification, unit, unit_price, order_qty, delivered_qty, received_qty, tax_amount, line_amount, promised_date, status) VALUES
(1, 1, 'MAT-RAW-001', '钢板A3', '1200*2400mm', 'PCS', 850.00, 100, 100, 100, 11050.00, 96050.00, '2026-04-15', 'COMPLETED'),
(1, 2, 'MAT-RAW-002', '钢板A4', '1200*1800mm', 'PCS', 295.00, 100, 0, 0, 3835.00, 29500.00, '2026-04-15', 'PENDING'),
(2, 1, 'MAT-RAW-003', '铝合金板', '1000*2000mm', 'PCS', 380.00, 150, 150, 150, 7410.00, 57000.00, '2026-04-16', 'COMPLETED'),
(2, 2, 'MAT-RAW-004', '不锈钢板', '1500*3000mm', 'PCS', 320.00, 100, 50, 50, 4160.00, 32000.00, '2026-04-16', 'PARTIAL'),
(3, 1, 'MAT-RAW-005', '塑料粒子PP', '通用级', 'KG', 45.00, 500, 300, 300, 1755.00, 13500.00, '2026-04-12', 'PARTIAL'),
(3, 2, 'MAT-RAW-005', '塑料粒子PP', '增强级', 'KG', 52.00, 400, 0, 0, 2704.00, 20800.00, '2026-04-12', 'PENDING'),
(4, 1, 'MAT-RAW-001', '钢板A3', '1200*2400mm', 'PCS', 860.00, 50, 50, 50, 5590.00, 43000.00, '2026-04-18', 'COMPLETED'),
(5, 1, 'MAT-RAW-002', '钢板A4', '1200*1800mm', 'PCS', 285.00, 100, 0, 0, 3705.00, 28500.00, '2026-04-20', 'PENDING'),
(6, 1, 'MAT-RAW-003', '铝合金板', '1000*2000mm', 'PCS', 15000.00, 10, 10, 0, 19500.00, 150000.00, '2026-05-06', 'PARTIAL'),
(7, 1, 'MAT-RAW-001', '钢板A3', '1200*2400mm', 'PCS', 900.00, 20, 20, 20, 2340.00, 18000.00, '2026-04-22', 'COMPLETED');

-- =====================================================
-- 询价单 (RFQ) - 10条
-- =====================================================
INSERT INTO scp_rfq (tenant_id, rfq_no, rfq_name, rfq_type, inquiry_date, deadline_date, currency, payment_terms, delivery_terms, quality_standard, status, total_bids, is_evaluated, created_by, created_at, updated_at) VALUES
(1, 'RFQ-2026-04001', '钢板年度采购询价', 'ANNUAL', '2026-03-25', '2026-04-10', 'CNY', '月结30天', 'FOB上海港', '国标GB/T3274-2017', 'CLOSED', 5, 1, 'admin', NOW(), NOW()),
(1, 'RFQ-2026-04002', '电子元器件紧急询价', 'QUICK', '2026-04-01', '2026-04-08', 'CNY', '月结30天', 'EXW深圳', '行业标准', 'AWARDED', 3, 1, 'admin', NOW(), NOW()),
(1, 'RFQ-2026-04003', '塑料粒子季度询价', 'STANDARD', '2026-04-02', '2026-04-18', 'CNY', '月结30天', 'CIF广州', 'ISO标准', 'PUBLISHED', 2, 0, 'admin', NOW(), NOW()),
(1, 'RFQ-2026-04004', '表面处理服务询价', 'STANDARD', '2026-04-03', '2026-04-20', 'CNY', '月结30天', 'DDP杭州', '工艺标准', 'DRAFT', 0, 0, 'admin', NOW(), NOW()),
(1, 'RFQ-2026-04005', '包装材料询价', 'STANDARD', '2026-04-04', '2026-04-22', 'CNY', '月结45天', 'FOB上海港', '包装标准', 'PUBLISHED', 1, 0, 'admin', NOW(), NOW()),
(1, 'RFQ-2026-04006', '软件系统采购询价', 'STANDARD', '2026-04-05', '2026-04-25', 'CNY', '年结', 'SAAS服务', '甲方要求', 'DRAFT', 0, 0, 'admin', NOW(), NOW()),
(1, 'RFQ-2026-04007', '物流运输服务询价', 'STANDARD', '2026-04-06', '2026-04-23', 'CNY', '月结30天', '门到门', '物流标准', 'PUBLISHED', 2, 0, 'admin', NOW(), NOW()),
(1, 'RFQ-2026-04008', '复合材料采购询价', 'ANNUAL', '2026-04-07', '2026-04-24', 'CNY', '月结30天', 'FOB成都', '国标', 'CLOSED', 4, 1, 'admin', NOW(), NOW()),
(1, 'RFQ-2026-04009', '激光设备询价', 'STANDARD', '2026-04-08', '2026-04-26', 'CNY', '月结60天', 'CIF武汉', '设备标准', 'PUBLISHED', 1, 0, 'admin', NOW(), NOW()),
(1, 'RFQ-2026-04010', '化工原料询价', 'STANDARD', '2026-04-09', '2026-04-27', 'CNY', '月结30天', 'FOB南京', '危化品标准', 'DRAFT', 0, 0, 'admin', NOW(), NOW());

-- 询价单明细
INSERT INTO scp_rfq_item (rfq_id, line_no, material_code, material_name, specification, unit, required_qty, target_price, market_price, requested_date) VALUES
(1, 1, 'MAT-RAW-001', '钢板A3', '1200*2400mm', 'PCS', 500, 800.00, 850.00, '2026-05-01'),
(1, 2, 'MAT-RAW-002', '钢板A4', '1200*1800mm', 'PCS', 300, 280.00, 295.00, '2026-05-01'),
(2, 1, 'MAT-RAW-003', '铝合金板', '1000*2000mm', 'PCS', 100, 350.00, 380.00, '2026-04-15'),
(3, 1, 'MAT-RAW-005', '塑料粒子PP', '通用级', 'KG', 2000, 42.00, 45.00, '2026-05-01'),
(3, 2, 'MAT-RAW-005', '塑料粒子PP', '增强级', 'KG', 1000, 50.00, 52.00, '2026-05-01'),
(4, 1, 'MAT-RAW-001', '钢板A3', '1200*2400mm', 'PCS', 200, 850.00, NULL, '2026-05-15'),
(5, 1, 'MAT-RAW-005', '包装材料', '标准包装', 'PCS', 5000, 5.00, 5.50, '2026-05-01'),
(6, 1, 'MAT-SW-001', 'ERP软件', '标准版', '套', 1, 80000.00, 95000.00, '2026-06-01'),
(7, 1, 'MAT-SVC-001', '物流服务', '国内整车', '次', 12, 5000.00, 5500.00, '2026-05-01'),
(8, 1, 'MAT-RAW-004', '不锈钢板', '1500*3000mm', 'PCS', 100, 280.00, 320.00, '2026-05-10');

-- 询价邀请供应商
INSERT INTO scp_rfq_invite (tenant_id, rfq_id, supplier_id, supplier_code, supplier_name, contact_person, contact_email, invite_date, response_status) VALUES
(1, 1, 1, 'SUP-TEST01', '华东精密机械有限公司', '张经理', 'zhang@example.com', '2026-03-25', 'QUOTED'),
(1, 1, 2, 'SUP-TEST02', '深圳创新电子科技有限公司', '李总监', 'li@example.com', '2026-03-25', 'QUOTED'),
(1, 1, 8, 'SUP-TEST08', '成都西南航空材料有限公司', '周经理', 'zhou@example.com', '2026-03-25', 'DECLINED'),
(1, 2, 2, 'SUP-TEST02', '深圳创新电子科技有限公司', '李总监', 'li@example.com', '2026-04-01', 'QUOTED'),
(1, 2, 9, 'SUP-TEST09', '武汉光谷激光设备有限公司', '吴工', 'wu@example.com', '2026-04-01', 'QUOTED'),
(1, 3, 3, 'SUP-TEST03', '苏州工业园区华鑫塑胶制品厂', '王厂长', 'wang@example.com', '2026-04-02', 'QUOTED'),
(1, 3, 5, 'SUP-TEST05', '广州岭南包装材料有限公司', '刘经理', 'liu@example.com', '2026-04-02', 'QUOTED'),
(1, 4, 4, 'SUP-TEST04', '杭州中达金属表面处理有限公司', '陈工', 'chen@example.com', '2026-04-03', 'PENDING'),
(1, 5, 5, 'SUP-TEST05', '广州岭南包装材料有限公司', '刘经理', 'liu@example.com', '2026-04-04', 'QUOTED'),
(1, 7, 7, 'SUP-TEST07', '天津港保税区润通物流有限公司', '孙经理', 'sun@example.com', '2026-04-06', 'QUOTED'),
(1, 8, 8, 'SUP-TEST08', '成都西南航空材料有限公司', '周经理', 'zhou@example.com', '2026-04-07', 'QUOTED'),
(1, 8, 1, 'SUP-TEST01', '华东精密机械有限公司', '张经理', 'zhang@example.com', '2026-04-07', 'QUOTED'),
(1, 9, 9, 'SUP-TEST09', '武汉光谷激光设备有限公司', '吴工', 'wu@example.com', '2026-04-08', 'QUOTED');

-- =====================================================
-- 供应商报价 (Supplier Quote) - 10条
-- =====================================================
INSERT INTO scp_supplier_quote (tenant_id, rfq_id, rfq_no, supplier_id, supplier_code, supplier_name, quote_no, quote_date, valid_until, currency, payment_terms, delivery_days, total_amount, is_accepted, is_lowest, rank_position, quote_status, created_at, updated_at) VALUES
(1, 1, 'RFQ-2026-04001', 1, 'SUP-TEST01', '华东精密机械有限公司', 'QUO-2026-04001', '2026-04-08', '2026-05-08', 'CNY', '月结30天', 15, 425000.00, 1, 1, 1, 'SUBMITTED', NOW(), NOW()),
(1, 1, 'RFQ-2026-04001', 2, 'SUP-TEST02', '深圳创新电子科技有限公司', 'QUO-2026-04002', '2026-04-08', '2026-05-08', 'CNY', '月结30天', 20, 438000.00, 0, 0, 2, 'SUBMITTED', NOW(), NOW()),
(1, 2, 'RFQ-2026-04002', 2, 'SUP-TEST02', '深圳创新电子科技有限公司', 'QUO-2026-04003', '2026-04-06', '2026-05-06', 'CNY', '月结30天', 10, 38000.00, 1, 1, 1, 'SUBMITTED', NOW(), NOW()),
(1, 3, 'RFQ-2026-04003', 3, 'SUP-TEST03', '苏州工业园区华鑫塑胶制品厂', 'QUO-2026-04004', '2026-04-15', '2026-05-15', 'CNY', '月结30天', 7, 94000.00, 0, 1, 1, 'SUBMITTED', NOW(), NOW()),
(1, 5, 'RFQ-2026-04005', 5, 'SUP-TEST05', '广州岭南包装材料有限公司', 'QUO-2026-04005', '2026-04-18', '2026-05-18', 'CNY', '月结45天', 5, 27500.00, 1, 1, 1, 'SUBMITTED', NOW(), NOW()),
(1, 7, 'RFQ-2026-04007', 7, 'SUP-TEST07', '天津港保税区润通物流有限公司', 'QUO-2026-04006', '2026-04-20', '2026-05-20', 'CNY', '月结30天', 3, 66000.00, 0, 1, 1, 'SUBMITTED', NOW(), NOW()),
(1, 8, 'RFQ-2026-04008', 8, 'SUP-TEST08', '成都西南航空材料有限公司', 'QUO-2026-04007', '2026-04-20', '2026-05-20', 'CNY', '月结30天', 12, 28500.00, 1, 1, 1, 'SUBMITTED', NOW(), NOW()),
(1, 8, 'RFQ-2026-04008', 1, 'SUP-TEST01', '华东精密机械有限公司', 'QUO-2026-04008', '2026-04-20', '2026-05-20', 'CNY', '月结30天', 15, 29500.00, 0, 0, 2, 'SUBMITTED', NOW(), NOW()),
(1, 9, 'RFQ-2026-04009', 9, 'SUP-TEST09', '武汉光谷激光设备有限公司', 'QUO-2026-04009', '2026-04-22', '2026-05-22', 'CNY', '月结60天', 30, 98000.00, 1, 1, 1, 'SUBMITTED', NOW(), NOW()),
(1, 3, 'RFQ-2026-04003', 5, 'SUP-TEST05', '广州岭南包装材料有限公司', 'QUO-2026-04010', '2026-04-16', '2026-05-16', 'CNY', '月结45天', 5, 98000.00, 0, 0, 2, 'REVISED', NOW(), NOW());

-- 供应商报价明细
INSERT INTO scp_quote_item (quote_id, rfq_line_id, material_code, material_name, unit, quoted_qty, unit_price, line_amount, delivery_date, lead_time_days) VALUES
(1, 1, 'MAT-RAW-001', '钢板A3', 'PCS', 500, 820.00, 410000.00, '2026-05-01', 15),
(1, 2, 'MAT-RAW-002', '钢板A4', 'PCS', 300, 50.00, 15000.00, '2026-05-01', 15),
(2, 1, 'MAT-RAW-001', '钢板A3', 'PCS', 500, 850.00, 425000.00, '2026-05-01', 20),
(2, 2, 'MAT-RAW-002', '钢板A4', 'PCS', 300, 43.33, 13000.00, '2026-05-01', 20),
(3, 3, 'MAT-RAW-003', '铝合金板', 'PCS', 100, 380.00, 38000.00, '2026-04-15', 10),
(4, 4, 'MAT-RAW-005', '塑料粒子PP', 'KG', 2000, 43.00, 86000.00, '2026-05-01', 7),
(4, 5, 'MAT-RAW-005', '塑料粒子PP', 'KG', 1000, 50.00, 5000.00, '2026-05-01', 7),
(5, 7, 'MAT-RAW-005', '包装材料', 'PCS', 5000, 5.50, 27500.00, '2026-05-01', 5),
(6, 9, 'MAT-SVC-001', '物流服务', '次', 12, 5500.00, 66000.00, '2026-05-01', 3),
(7, 10, 'MAT-RAW-004', '不锈钢板', 'PCS', 100, 285.00, 28500.00, '2026-05-10', 12);

-- =====================================================
-- 销售订单 (Sales Order) - 10条
-- =====================================================
INSERT INTO scp_sales_order (tenant_id, so_no, so_type, customer_id, customer_code, customer_name, contact_person, contact_phone, contact_email, sales_person_name, order_date, promised_date, currency, payment_terms, tax_rate, total_amount, total_qty, delivered_amount, delivered_qty, approval_status, status, source_type, delivery_address, created_at, updated_at) VALUES
(1, 'SO-2026-04001', 'STANDARD', 1, 'CUS001', '杭州汽车配件有限公司', '赵经理', '13900010001', 'zhao@hzauto.com', '张三', '2026-04-01', '2026-04-20', 'CNY', '月结30天', 13.00, 226000.00, 200, 0.00, 0, 'APPROVED', 'CONFIRMED', 'MANUAL', '杭州市萧山区', NOW(), NOW()),
(1, 'SO-2026-04002', 'STANDARD', 2, 'CUS002', '上海机电设备有限公司', '钱经理', '13900010002', 'qian@shjd.com', '张三', '2026-04-02', '2026-04-22', 'CNY', '月结30天', 13.00, 169000.00, 150, 0.00, 0, 'APPROVED', 'CONFIRMED', 'MANUAL', '上海市闵行区', NOW(), NOW()),
(1, 'SO-2026-04003', 'URGENT', 3, 'CUS003', '宁波电子科技有限公司', '孙经理', '13900010003', 'sun@nbec.com', '李四', '2026-04-03', '2026-04-12', 'CNY', '预付30%', 13.00, 85000.00, 100, 85000.00, 100, 'APPROVED', 'SHIPPED', 'MANUAL', '宁波市鄞州区', NOW(), NOW()),
(1, 'SO-2026-04004', 'STANDARD', 4, 'CUS004', '温州电器股份有限公司', '李经理', '13900010004', 'li@wzelectric.com', '李四', '2026-04-04', '2026-04-25', 'CNY', '月结45天', 13.00, 135000.00, 120, 45000.00, 40, 'APPROVED', 'PARTIAL', 'MANUAL', '温州市乐清市', NOW(), NOW()),
(1, 'SO-2026-04005', 'DISTRIBUTION', 1, 'CUS001', '杭州汽车配件有限公司', '赵经理', '13900010001', 'zhao@hzauto.com', '王五', '2026-04-05', '2026-04-28', 'CNY', '月结30天', 13.00, 95000.00, 500, 0.00, 0, 'PENDING', 'PENDING', 'MANUAL', '杭州市萧山区', NOW(), NOW()),
(1, 'SO-2026-04006', 'STANDARD', 2, 'CUS002', '上海机电设备有限公司', '钱经理', '13900010002', 'qian@shjd.com', '王五', '2026-04-06', '2026-05-01', 'CNY', '月结30天', 13.00, 280000.00, 80, 0.00, 0, 'APPROVED', 'CONFIRMED', 'MANUAL', '上海市闵行区', NOW(), NOW()),
(1, 'SO-2026-04007', 'STANDARD', 3, 'CUS003', '宁波电子科技有限公司', '孙经理', '13900010003', 'sun@nbec.com', '张三', '2026-04-07', '2026-05-05', 'CNY', '月结30天', 13.00, 78000.00, 60, 0.00, 0, 'PENDING', 'PENDING', 'MANUAL', '宁波市鄞州区', NOW(), NOW()),
(1, 'SO-2026-04008', 'URGENT', 4, 'CUS004', '温州电器股份有限公司', '李经理', '13900010004', 'li@wzelectric.com', '李四', '2026-04-08', '2026-04-18', 'CNY', '预付50%', 13.00, 156000.00, 45, 156000.00, 45, 'APPROVED', 'SHIPPED', 'MANUAL', '温州市乐清市', NOW(), NOW()),
(1, 'SO-2026-04009', 'STANDARD', 1, 'CUS001', '杭州汽车配件有限公司', '赵经理', '13900010001', 'zhao@hzauto.com', '王五', '2026-04-09', '2026-05-10', 'CNY', '月结30天', 13.00, 198000.00, 180, 0.00, 0, 'APPROVED', 'CONFIRMED', 'MANUAL', '杭州市萧山区', NOW(), NOW()),
(1, 'SO-2026-04010', 'STANDARD', 2, 'CUS002', '上海机电设备有限公司', '钱经理', '13900010002', 'qian@shjd.com', '张三', '2026-04-10', '2026-05-15', 'CNY', '月结30天', 13.00, 245000.00, 220, 0.00, 0, 'PENDING', 'PENDING', 'MANUAL', '上海市闵行区', NOW(), NOW());

-- 销售订单明细
INSERT INTO scp_sales_order_item (so_id, line_no, material_code, material_name, specification, unit, unit_price, order_qty, delivered_qty, shipped_qty, tax_amount, line_amount, promised_date, status) VALUES
(1, 1, 'MAT-RAW-001', '钢板A3', '1200*2400mm', 'PCS', 980.00, 150, 0, 0, 19110.00, 147000.00, '2026-04-20', 'PENDING'),
(1, 2, 'MAT-RAW-002', '钢板A4', '1200*1800mm', 'PCS', 580.00, 50, 0, 0, 3770.00, 29000.00, '2026-04-20', 'PENDING'),
(1, 3, 'MAT-RAW-003', '铝合金板', '1000*2000mm', 'PCS', 1000.00, 50, 0, 0, 6500.00, 50000.00, '2026-04-20', 'PENDING'),
(2, 1, 'MAT-RAW-001', '钢板A3', '1200*2400mm', 'PCS', 990.00, 100, 0, 0, 12870.00, 99000.00, '2026-04-22', 'PENDING'),
(2, 2, 'MAT-RAW-004', '不锈钢板', '1500*3000mm', 'PCS', 700.00, 100, 0, 0, 9100.00, 70000.00, '2026-04-22', 'PENDING'),
(3, 1, 'MAT-RAW-003', '铝合金板', '1000*2000mm', 'PCS', 850.00, 100, 100, 100, 11050.00, 85000.00, '2026-04-12', 'COMPLETED'),
(4, 1, 'MAT-RAW-001', '钢板A3', '1200*2400mm', 'PCS', 1000.00, 80, 40, 40, 10400.00, 80000.00, '2026-04-25', 'PARTIAL'),
(4, 2, 'MAT-RAW-002', '钢板A4', '1200*1800mm', 'PCS', 550.00, 40, 0, 0, 2860.00, 22000.00, '2026-04-25', 'PENDING'),
(4, 3, 'MAT-RAW-003', '铝合金板', '1000*2000mm', 'PCS', 825.00, 33, 0, 0, 3543.75, 27272.73, '2026-04-25', 'PENDING'),
(5, 1, 'MAT-RAW-005', '塑料粒子PP', '通用级', 'KG', 190.00, 500, 0, 0, 12350.00, 95000.00, '2026-04-28', 'PENDING');

-- =====================================================
-- 供应商绩效 (Supplier KPI) - 10条
-- =====================================================
INSERT INTO scp_supplier_kpi (tenant_id, supplier_id, supplier_code, supplier_name, evaluation_month, evaluation_date, evaluated_by, evaluated_by_name, on_time_delivery_rate, total_delivery_orders, on_time_delivery_count, avg_delay_days, quality_pass_rate, total_iqc_orders, passed_iqc_orders, defect_parts_count, defect_rate, price_competitiveness, last_purchase_price, market_avg_price, total_score, grade, rank_position) VALUES
(1, 1, 'SUP-TEST01', '华东精密机械有限公司', '2026-03', '2026-04-05', 1, '系统管理员', 95.50, 20, 19, 0.5, 98.00, 15, 15, 5, 0.0025, 92.00, 850.00, 870.00, 95.20, 'A', 1),
(1, 2, 'SUP-TEST02', '深圳创新电子科技有限公司', '2026-03', '2026-04-05', 1, '系统管理员', 88.00, 15, 13, 1.2, 95.50, 10, 10, 8, 0.0060, 88.00, 380.00, 395.00, 89.80, 'B', 3),
(1, 3, 'SUP-TEST03', '苏州工业园区华鑫塑胶制品厂', '2026-03', '2026-04-05', 1, '系统管理员', 92.00, 12, 11, 0.8, 96.00, 8, 8, 3, 0.0030, 90.00, 45.00, 47.00, 92.50, 'A', 2),
(1, 4, 'SUP-TEST04', '杭州中达金属表面处理有限公司', '2026-03', '2026-04-05', 1, '系统管理员', 85.00, 8, 7, 1.5, 93.00, 6, 6, 4, 0.0055, 85.00, NULL, NULL, 86.30, 'B', 5),
(1, 5, 'SUP-TEST05', '广州岭南包装材料有限公司', '2026-03', '2026-04-05', 1, '系统管理员', 90.00, 10, 9, 1.0, 94.00, 7, 7, 4, 0.0045, 88.00, 5.50, 5.80, 90.50, 'B', 4),
(1, 6, 'SUP-TEST06', '北京中关村软件股份有限公司', '2026-03', '2026-04-05', 1, '系统管理员', 100.00, 3, 3, 0.0, 100.00, 2, 2, 0, 0.0000, 95.00, 95000.00, 98000.00, 98.00, 'A', 1),
(1, 7, 'SUP-TEST07', '天津港保税区润通物流有限公司', '2026-03', '2026-04-05', 1, '系统管理员', 93.00, 6, 6, 0.3, NULL, NULL, NULL, NULL, NULL, 90.00, 5500.00, 5800.00, 91.50, 'B', 2),
(1, 8, 'SUP-TEST08', '成都西南航空材料有限公司', '2026-03', '2026-04-05', 1, '系统管理员', 87.00, 9, 8, 1.8, 91.00, 5, 5, 6, 0.0080, 82.00, 320.00, 350.00, 85.20, 'C', 6),
(1, 9, 'SUP-TEST09', '武汉光谷激光设备有限公司', '2026-03', '2026-04-05', 1, '系统管理员', 80.00, 5, 4, 2.5, 88.00, 4, 4, 5, 0.0100, 78.00, 98000.00, 105000.00, 81.00, 'C', 7),
(1, 10, 'SUP-TEST10', '南京金陵化工有限公司', '2026-03', '2026-04-05', 1, '系统管理员', 91.00, 11, 10, 0.7, 95.00, 9, 9, 5, 0.0045, 89.00, NULL, NULL, 91.00, 'B', 3);
