-- MOM3.0 演示测试数据 (修正版)
-- 基于实际表结构编写
SET client_encoding = 'UTF8';
SET ON_ERROR_STOP on;

-- ============================================
-- 1. BOM数据 (boms + bom_items)
-- ============================================
INSERT INTO boms (id, created_at, updated_at, tenant_id, bom_code, product_id, product_code, product_name, version, bom_type, status) VALUES
(1, NOW(), NOW(), 1, 'BOM-2024-001', 10, 'MAT-010', '成品A', 'V1.0', 'production', 1),
(2, NOW(), NOW(), 1, 'BOM-2024-002', 11, 'MAT-011', '成品B', 'V1.0', 'production', 1),
(3, NOW(), NOW(), 1, 'BOM-2024-003', 12, 'MAT-012', '成品C', 'V1.0', 'production', 1);

INSERT INTO bom_items (id, created_at, updated_at, tenant_id, bom_id, material_id, material_code, material_name, quantity, unit, scrap_rate) VALUES
(1, NOW(), NOW(), 1, 1, 1, 'MAT-001', '钢板A3', 2.00, 'PCS', 0),
(2, NOW(), NOW(), 1, 1, 5, 'MAT-005', '塑料粒子PP', 0.50, 'KG', 0),
(3, NOW(), NOW(), 1, 1, 8, 'MAT-008', '螺丝M8', 4.00, 'PCS', 0),
(4, NOW(), NOW(), 1, 2, 2, 'MAT-002', '钢板A4', 1.50, 'PCS', 0),
(5, NOW(), NOW(), 1, 2, 3, 'MAT-003', '铝合金板', 1.00, 'PCS', 0),
(6, NOW(), NOW(), 1, 2, 9, 'MAT-009', '包装材料', 2.00, 'PCS', 0),
(7, NOW(), NOW(), 1, 3, 4, 'MAT-004', '不锈钢板', 1.20, 'PCS', 0),
(8, NOW(), NOW(), 1, 3, 6, 'MAT-006', '电线', 0.30, 'M', 0);

-- ============================================
-- 2. 缺陷代码
-- ============================================
INSERT INTO defect_codes (id, created_at, updated_at, tenant_id, defect_code, defect_name, defect_type, severity, is_key) VALUES
(1, NOW(), NOW(), 1, 'DC-001', '划伤', 'surface', 2, true),
(2, NOW(), NOW(), 1, 'DC-002', '尺寸超差', 'dimension', 3, true),
(3, NOW(), NOW(), 1, 'DC-003', '缺件', 'missing', 3, true),
(4, NOW(), NOW(), 1, 'DC-004', '变形', 'deformation', 2, false),
(5, NOW(), NOW(), 1, 'DC-005', '污渍', 'stain', 1, false),
(6, NOW(), NOW(), 1, 'DC-006', '色差', 'color', 1, false),
(7, NOW(), NOW(), 1, 'DC-007', '毛刺', 'burr', 1, false),
(8, NOW(), NOW(), 1, 'DC-008', '裂纹', 'crack', 3, true);

-- ============================================
-- 3. 生产工单
-- ============================================
INSERT INTO production_orders (id, created_at, updated_at, tenant_id, order_no, material_id, material_code, material_name, unit, quantity, completed_qty, rejected_qty, workshop_id, workshop_name, line_id, line_name, plan_start_date, plan_end_date, order_status, priority, remark) VALUES
(1, NOW(), NOW(), 1, 'PO-2024-0001', 10, 'MAT-010', '成品A', 'PCS', 1000, 650, 5, 1, '一车间', 1, '生产线A', '2024-04-01', '2024-04-10', 'in_progress', 2, '优先紧急订单'),
(2, NOW(), NOW(), 1, 'PO-2024-0002', 11, 'MAT-011', '成品B', 'PCS', 500, 500, 8, 1, '一车间', 1, '生产线A', '2024-03-20', '2024-03-28', 'completed', 1, '已完成'),
(3, NOW(), NOW(), 1, 'PO-2024-0003', 12, 'MAT-012', '成品C', 'PCS', 800, 200, 2, 2, '二车间', 2, '生产线B', '2024-04-05', '2024-04-15', 'in_progress', 2, '连续生产中'),
(4, NOW(), NOW(), 1, 'PO-2024-0004', 13, 'MAT-013', '成品D', 'PCS', 300, 0, 0, 2, '二车间', 3, '生产线C', '2024-04-18', '2024-04-25', 'pending', 1, '待排产'),
(5, NOW(), NOW(), 1, 'PO-2024-0005', 14, 'MAT-014', '成品E', 'PCS', 600, 0, 0, 3, '三车间', NULL, NULL, '2024-04-20', '2024-04-30', 'pending', 2, '新品试制');

-- ============================================
-- 4. 质检记录 (IQC)
-- ============================================
INSERT INTO iqcs (id, created_at, updated_at, tenant_id, iqc_no, supplier_id, supplier_name, material_id, material_code, material_name, quantity, unit, check_user_id, check_user_name, check_date, result, remark) VALUES
(1, NOW(), NOW(), 1, 'IQC-2024-0001', 1, '供应商A', 1, 'MAT-001', '钢板A3', 100, 'PCS', 1, 'admin', NOW(), 3, '来料检验合格'),
(2, NOW(), NOW(), 1, 'IQC-2024-0002', 2, '供应商B', 2, 'MAT-002', '钢板A4', 50, 'PCS', 1, 'admin', NOW(), 3, '来料检验合格'),
(3, NOW(), NOW(), 1, 'IQC-2024-0003', 1, '供应商A', 5, 'MAT-005', '塑料粒子PP', 200, 'KG', 2, 'nhg', NOW(), 3, '来料检验'),
(4, NOW(), NOW(), 1, 'IQC-2024-0004', 3, '供应商C', 3, 'MAT-003', '铝合金板', 80, 'PCS', 1, 'admin', NOW(), 3, '免检批次'),
(5, NOW(), NOW(), 1, 'IQC-2024-0005', 2, '供应商B', 8, 'MAT-008', '螺丝M8', 150, 'PCS', 1, 'admin', NOW(), 4, '来料不合格');

-- ============================================
-- 5. IPQC记录
-- ============================================
INSERT INTO ip_qcs (id, created_at, updated_at, tenant_id, ipqc_no, production_order_id, production_order_no, workstation_id, workstation_name, inspection_qty, defect_qty, defect_rate, inspection_result, inspector, inspection_time, result, remark) VALUES
(1, NOW(), NOW(), 1, 'IPQC-2024-0001', 1, 'PO-2024-0001', 1, '工作站A', 50, 2, 4.0, 'qualified', 'admin', NOW(), 3, '工序检验'),
(2, NOW(), NOW(), 1, 'IPQC-2024-0002', 1, 'PO-2024-0001', 2, '工作站B', 80, 1, 1.25, 'qualified', 'nhg', NOW(), 3, '装配工序'),
(3, NOW(), NOW(), 1, 'IPQC-2024-0003', 3, 'PO-2024-0003', 3, '工作站C', 100, 3, 3.0, 'qualified', 'admin', NOW(), 3, '加工工序'),
(4, NOW(), NOW(), 1, 'IPQC-2024-0004', 2, 'PO-2024-0002', 1, '工作站A', 60, 8, 13.3, 'rejected', 'admin', NOW(), 4, '批量不合格');

-- ============================================
-- 6. FQC记录
-- ============================================
INSERT INTO fqcs (id, created_at, updated_at, tenant_id, fqc_no, production_order_id, production_order_no, final_qty, accepted_qty, rejected_qty, inspection_result, inspector, inspection_time, result, remark) VALUES
(1, NOW(), NOW(), 1, 'FQC-2024-0001', 2, 'PO-2024-0002', 500, 492, 8, 'qualified', 'admin', NOW(), 3, '最终检验合格'),
(2, NOW(), NOW(), 1, 'FQC-2024-0002', 1, 'PO-2024-0001', 200, 195, 5, 'qualified', 'nhg', NOW(), 3, '首批检验');

-- ============================================
-- 7. 调拨单
-- ============================================
INSERT INTO wms_transfer_order (id, created_at, updated_at, tenant_id, transfer_no, from_warehouse_id, to_warehouse_id, transfer_date, transfer_user_id, status, transfer_type, requester_name, request_time, remark) VALUES
(1, NOW(), NOW(), 1, 'TR-2024-0001', 1, 2, '2024-04-11', 1, 'COMPLETED', 'transfer', 'admin', '2024-04-10 09:00:00', '原材料转成品'),
(2, NOW(), NOW(), 1, 'TR-2024-0002', 4, 3, '2024-04-13', 2, 'IN_TRANSIT', 'transfer', 'nhg', '2024-04-12 09:00:00', '线边仓备料'),
(3, NOW(), NOW(), 1, 'TR-2024-0003', 2, 5, '2024-04-16', 1, 'PENDING', 'transfer', 'admin', '2024-04-15 09:00:00', '等待审批');

INSERT INTO wms_transfer_order_item (id, created_at, updated_at, tenant_id, transfer_order_id, material_id, material_code, material_name, quantity, shipped_qty, received_qty, remark) VALUES
(1, NOW(), NOW(), 1, 1, 10, 'MAT-010', '钢板A3', 100, 100, 100, '已收货'),
(2, NOW(), NOW(), 1, 1, 11, 'MAT-011', '钢板A4', 50, 50, 50, '已收货'),
(3, NOW(), NOW(), 1, 2, 3, 'MAT-003', '铝合金板', 80, 80, 0, '发货中'),
(4, NOW(), NOW(), 1, 3, 12, 'MAT-012', '不锈钢板', 30, 0, 0, '待发货');

-- ============================================
-- 8. 盘点单
-- ============================================
INSERT INTO wms_stock_check (id, created_at, updated_at, tenant_id, check_no, warehouse_id, check_type, status, check_user_id, check_date, plan_start_date, plan_end_date, checker_name, remark) VALUES
(1, NOW(), NOW(), 1, 'SC-2024-0001', 1, 'full', 'COMPLETED', 1, '2024-04-10', '2024-04-10', '2024-04-10', 'admin', '原材料仓月度盘点'),
(2, NOW(), NOW(), 1, 'SC-2024-0002', 2, 'full', 'IN_PROGRESS', 2, '2024-04-14', '2024-04-14', '2024-04-14', 'nhg', '成品仓盘点');

INSERT INTO wms_stock_check_item (id, created_at, updated_at, tenant_id, stock_check_id, material_id, material_code, material_name, system_qty, counted_qty, variance_qty, handle_status, handle_remark, remark) VALUES
(1, NOW(), NOW(), 1, 1, 1, 'MAT-001', '钢板A3', 1000, 995, -5, 3, '已调整', '钢板A3盘点差异'),
(2, NOW(), NOW(), 1, 1, 2, 'MAT-002', '钢板A4', 500, 502, 2, 3, '已调整', '钢板A4多出2件'),
(3, NOW(), NOW(), 1, 2, 10, 'MAT-010', '钢板A3', 200, 0, -200, 1, '待处理', '成品未盘点');

-- ============================================
-- 9. 库存
-- ============================================
INSERT INTO inventories (id, created_at, updated_at, tenant_id, warehouse_id, location_id, material_id, material_code, material_name, quantity, available_qty, allocated_qty, locked_qty) VALUES
(1, NOW(), NOW(), 1, 1, 1, 1, 'MAT-001', '钢板A3', 1000, 900, 100, 0),
(2, NOW(), NOW(), 1, 1, 2, 2, 'MAT-002', '钢板A4', 500, 500, 0, 0),
(3, NOW(), NOW(), 1, 1, 1, 3, 'MAT-003', '铝合金板', 300, 300, 0, 0),
(4, NOW(), NOW(), 1, 2, 3, 10, 'MAT-010', '钢板A3', 200, 150, 50, 0),
(5, NOW(), NOW(), 1, 2, 4, 11, 'MAT-011', '钢板A4', 150, 150, 0, 0),
(6, NOW(), NOW(), 1, 3, 5, 5, 'MAT-005', '塑料粒子PP', 80, 60, 20, 0),
(7, NOW(), NOW(), 1, 4, 6, 8, 'MAT-008', '螺丝M8', 500, 500, 0, 0),
(8, NOW(), NOW(), 1, 5, 7, 12, 'MAT-012', '不锈钢板', 100, 100, 0, 0);

-- ============================================
-- 10. 月计划
-- ============================================
INSERT INTO mes_order_month (id, created_at, updated_at, tenant_id, month_plan_no, plan_month, title, workshop_id, workshop_name, total_product_count, total_plan_qty, total_completed_qty, approval_status, created_by) VALUES
(1, NOW(), NOW(), 1, 'MP-2024-04', '2024-04', '4月生产计划', 1, '一车间', 3, 5000, 1500, 'RELEASED', 'admin'),
(2, NOW(), NOW(), 1, 'MP-2024-05', '2024-05', '5月生产计划', 1, '一车间', 2, 3000, 0, 'APPROVED', 'admin');

INSERT INTO mes_order_month_item (id, created_at, updated_at, tenant_id, month_plan_id, material_id, material_code, material_name, plan_qty, completed_qty, completed, remark) VALUES
(1, NOW(), NOW(), 1, 1, 10, 'MAT-010', '成品A', 2000, 650, 32.5, '进行中'),
(2, NOW(), NOW(), 1, 1, 11, 'MAT-011', '成品B', 1500, 500, 33.3, '进行中'),
(3, NOW(), NOW(), 1, 1, 12, 'MAT-012', '成品C', 1500, 200, 13.3, '进行中'),
(4, NOW(), NOW(), 1, 2, 13, 'MAT-013', '成品D', 1500, 0, 0, '待排产'),
(5, NOW(), NOW(), 1, 2, 14, 'MAT-014', '成品E', 1500, 0, 0, '待排产');

-- ============================================
-- 11. 日计划
-- ============================================
INSERT INTO mes_order_day (id, created_at, updated_at, tenant_id, day_plan_no, plan_date, workshop_id, month_plan_item_id, material_id, material_code, material_name, plan_qty, completed_qty, qualified_qty, status, remark) VALUES
(1, NOW(), NOW(), 1, 'DP-2024-04-01', '2024-04-01', 1, 1, 10, 'MAT-010', '钢板A3', 500, 480, 475, 'COMPLETED', '日计划-1'),
(2, NOW(), NOW(), 1, 'DP-2024-04-02', '2024-04-02', 1, 1, 10, 'MAT-010', '钢板A3', 500, 520, 515, 'COMPLETED', '日计划-2'),
(3, NOW(), NOW(), 1, 'DP-2024-04-03', '2024-04-03', 1, 2, 11, 'MAT-011', '钢板A4', 300, 200, 198, 'IN_PROGRESS', '日计划-进行中'),
(4, NOW(), NOW(), 1, 'DP-2024-04-08', '2024-04-08', 1, 1, 10, 'MAT-010', '钢板A3', 400, 0, 0, 'PENDING', '日计划-待执行');

INSERT INTO mes_order_day_item (id, created_at, updated_at, tenant_id, order_day_id, material_id, material_code, material_name, plan_qty, completed_qty, status, remark) VALUES
(1, NOW(), NOW(), 1, 1, 10, 'MAT-010', '钢板A3', 500, 480, 3, '工序1'),
(2, NOW(), NOW(), 1, 2, 10, 'MAT-010', '钢板A3', 500, 520, 3, '工序1-超额完成'),
(3, NOW(), NOW(), 1, 3, 11, 'MAT-011', '钢板A4', 300, 200, 2, '工序2');

-- ============================================
-- 12. 设备台账
-- ============================================
INSERT INTO equ_equipment (id, created_at, updated_at, tenant_id, equipment_code, equipment_name, equipment_type, brand, model, serial_number, workshop_id, workshop_name, purchase_date, purchase_price, status) VALUES
(1, NOW(), NOW(), 1, 'EQ-001', '加工中心A1', '加工中心', '沈阳机床', 'VMC-850', 'SN-2020-001', 1, '一车间', '2020-01-15', 500000.00, 1),
(2, NOW(), NOW(), 1, 'EQ-002', '数控车床B1', '数控车床', '大连机床', 'CK-6150', 'SN-2020-002', 1, '一车间', '2020-03-20', 300000.00, 1),
(3, NOW(), NOW(), 1, 'EQ-003', '激光切割机C1', '激光切割', '大族激光', 'LCT-3015', 'SN-2021-005', 2, '二车间', '2021-06-10', 800000.00, 1),
(4, NOW(), NOW(), 1, 'EQ-004', '冲压机D1', '冲压设备', '济南二机床', 'J21-100', 'SN-2019-008', 2, '二车间', '2019-11-30', 200000.00, 1),
(5, NOW(), NOW(), 1, 'EQ-005', '焊接机器人E1', '焊接机器人', 'ABB', 'ARCO-2000', 'SN-2022-003', 3, '三车间', '2022-05-15', 600000.00, 1);

-- ============================================
-- 13. 设备维修
-- ============================================
INSERT INTO equ_equipment_repair (id, created_at, updated_at, tenant_id, repair_no, equipment_id, equipment_name, repair_type, fault_desc, fault_time, reporter, assignee, scheduled_start, scheduled_end, actual_start, actual_end, status, actual_cost, evaluation, remark) VALUES
(1, NOW(), NOW(), 1, 'REP-2024-0001', 1, '加工中心A1', 'breakdown', '主轴异响', '2024-04-05 10:00:00', 'admin', 'nhg', '2024-04-05 14:00:00', '2024-04-05 18:00:00', '2024-04-05 14:00:00', '2024-04-05 17:30:00', 4, 800.00, 5, '维修及时'),
(2, NOW(), NOW(), 1, 'REP-2024-0002', 3, '激光切割机C1', 'preventive', '激光器保养', '2024-04-10 08:00:00', 'nhg', 'admin', '2024-04-12 09:00:00', '2024-04-12 12:00:00', '2024-04-12 09:00:00', '2024-04-12 11:30:00', 4, 500.00, 5, '定期保养完成'),
(3, NOW(), NOW(), 1, 'REP-2024-0003', 2, '数控车床B1', 'breakdown', '刀架故障', '2024-04-15 09:30:00', 'admin', 'nhg', '2024-04-15 13:00:00', '2024-04-15 16:00:00', NULL, NULL, 2, NULL, NULL, '待维修');

-- ============================================
-- 14. 告警规则
-- ============================================
INSERT INTO alert_rule (id, created_at, rule_code, rule_name, alert_type, biz_module, condition_expression, condition_params, severity_level, notification_channels, is_enabled, check_interval, tenant_id) VALUES
(1, NOW(), 'ALERT-001', '设备OEE低于70%', 'EQUIPMENT', 'equipment', 'oee_value < 70', '{"oee_field":"oee_value"}', 'HIGH', '["feishu"]', 1, 60, 1),
(2, NOW(), 'ALERT-002', '库存低于安全库存', 'INVENTORY', 'wms', 'quantity < safe_stock', '{"qty_field":"quantity"}', 'MEDIUM', '["feishu"]', 1, 60, 1),
(3, NOW(), 'ALERT-003', '质量不合格率超5%', 'QUALITY', 'quality', 'reject_rate > 5', '{"rate_field":"reject_rate"}', 'HIGH', '["feishu"]', 1, 60, 1),
(4, NOW(), 'ALERT-004', '工单超期', 'PRODUCTION', 'production', 'planned_end_date < NOW()', '{}', 'MEDIUM', '["feishu"]', 0, 60, 1);

INSERT INTO alert_record (id, created_at, rule_id, alert_no, alert_type, biz_module, source_id, source_name, alert_content, trigger_value, status, handler, handle_time, handle_result, handle_remark, tenant_id) VALUES
(1, NOW(), 1, 'ALT-2024-0001', 'EQUIPMENT', 'equipment', 3, '激光切割机C1', '设备OEE低于70%告警', '68.5', 3, 'admin', NOW(), '已处理', '调整切割参数', 1),
(2, NOW(), 3, 'ALT-2024-0002', 'QUALITY', 'quality', 4, 'PO-2024-0004', '工单超期未完工', '2024-04-25', 1, NULL, NULL, NULL, NULL, 1),
(3, NOW(), 2, 'ALT-2024-0003', 'INVENTORY', 'wms', 6, '线边仓-塑料粒子PP', '库存低于安全库存', '60', 2, 'nhg', NOW(), '已确认', '已通知采购', 1);

-- ============================================
-- 15. BPM流程模型
-- ============================================
INSERT INTO bpm_process_model (id, created_at, model_code, model_name, model_type, version, category, description, form_type, is_published, is_active, tenant_id) VALUES
(1, NOW(), 'BPM-PURCHASE', '采购审批流程', 'PURCHASE', '1.0', 'procurement', '采购申请多级审批', 'custom', 1, 1, 1),
(2, NOW(), 'BPM-QC', '质量异常处理流程', 'QUALITY', '1.0', 'quality', '质量异常-NCR处理流程', 'custom', 1, 1, 1);

INSERT INTO bpm_process_instance (id, created_at, model_id, instance_key, business_key, initiator, current_node_name, status, start_time, tenant_id) VALUES
(1, NOW(), 1, 'PI-2024-0001', 'PO-2024-0001', 'admin', '部门经理审批', 'RUNNING', '2024-04-10 09:00:00', 1),
(2, NOW(), 1, 'PI-2024-0002', 'PO-2024-0002', 'nhg', '部门经理审批', 'COMPLETED', '2024-04-01 10:00:00', 1);

-- ============================================
-- 16. 安灯呼叫
-- ============================================
INSERT INTO andon_call (id, created_at, tenant_id, call_no, workshop_id, workshop_name, production_line_id, production_line_name, workstation_id, workstation_name, andon_type, andon_type_name, call_level, description, call_by, call_time, response_by, response_time, resolve_time, status, remark) VALUES
(1, NOW(), 1, 'ANDON-2024-0001', 1, '一车间', 1, '生产线A', 1, '工作站A', 'quality', '质量呼叫', 2, '发现质量问题-尺寸超差', 'admin', NOW(), 'nhg', NOW(), NOW(), 3, '质量呼叫-已解决'),
(2, NOW(), 1, 'ANDON-2024-0002', 2, '二车间', 2, '生产线B', 3, '工作站C', 'equipment', '设备呼叫', 1, '设备故障-无法启动', 'nhg', NOW(), 'admin', NOW(), NOW(), 3, '设备呼叫-已解决'),
(3, NOW(), 1, 'ANDON-2024-0003', 1, '一车间', 1, '生产线A', 2, '工作站B', 'material', '物料呼叫', 3, '物料短缺-待料停线', 'admin', NOW(), NULL, NULL, NULL, 1, '物料呼叫-待处理');

-- ============================================
-- 17. 采购结算
-- ============================================
INSERT INTO fin_purchase_settlement (id, created_at, settlement_no, settlement_type, related_type, related_id, related_no, supplier_id, supplier_code, supplier_name, invoice_no, invoice_date, goods_amount, tax_amount, total_amount, paid_amount, status, tenant_id) VALUES
(1, NOW(), 'PS-2024-0001', 'PURCHASE', 'PO', 1, 'PO-2024-0001', 1, 'SUP-001', '供应商A', 'INV-2024-001', '2024-04-01', 100000.00, 16000.00, 116000.00, 116000.00, 'PAID', 1),
(2, NOW(), 'PS-2024-0002', 'PURCHASE', 'PO', 2, 'PO-2024-0002', 2, 'SUP-002', '供应商B', 'INV-2024-002', '2024-04-15', 50000.00, 8000.00, 58000.00, 0, 'PENDING', 1);

INSERT INTO fin_sales_settlement (id, created_at, settlement_no, settlement_type, related_type, customer_id, customer_code, customer_name, invoice_no, invoice_date, goods_amount, tax_amount, total_amount, paid_amount, status, tenant_id) VALUES
(1, NOW(), 'SS-2024-0001', 'SALES', 'SO', 1, 'CUST-001', '客户A', 'INV-S-2024-001', '2024-04-05', 200000.00, 32000.00, 232000.00, 232000.00, 'PAID', 1);

-- ============================================
-- 18. APS MPS
-- ============================================
INSERT INTO aps_mps (id, created_at, updated_at, tenant_id, mps_no, plan_month, material_id, material_code, material_name, quantity, status) VALUES
(1, NOW(), NOW(), 1, 'MPS-2024-04-001', '2024-04', 10, 'MAT-010', '钢板A3', 5000, 3),
(2, NOW(), NOW(), 1, 'MPS-2024-04-002', '2024-04', 11, 'MAT-011', '钢板A4', 3000, 3),
(3, NOW(), NOW(), 1, 'MPS-2024-05-001', '2024-05', 12, 'MAT-012', '不锈钢板', 4000, 2);

-- ============================================
-- 19. 班组
-- ============================================
INSERT INTO mes_team (id, created_at, updated_at, tenant_id, team_code, team_name, workshop_id, leader_id, status) VALUES
(1, NOW(), NOW(), 1, 'TEAM-001', '一车间甲班', 1, 1, 1),
(2, NOW(), NOW(), 1, 'TEAM-002', '一车间乙班', 1, 2, 1),
(3, NOW(), NOW(), 1, 'TEAM-003', '二车间甲班', 2, 1, 1);

INSERT INTO mes_team_member (id, created_at, updated_at, tenant_id, team_id, user_id, user_name, role) VALUES
(1, NOW(), NOW(), 1, 1, 1, 'admin', 'leader'),
(2, NOW(), NOW(), 1, 1, 2, 'nhg', 'member'),
(3, NOW(), NOW(), 1, 2, 1, 'admin', 'member');

-- ============================================
-- 20. 生产报工
-- ============================================
INSERT INTO production_reports (id, created_at, updated_at, tenant_id, report_no, production_order_id, production_order_no, workstation_id, workstation_name, report_type, reported_by, reported_by_name, report_time, plan_qty, qualified_qty, rejected_qty, working_hours, efficiency, remark) VALUES
(1, NOW(), NOW(), 1, 'RPT-2024-0001', 1, 'PO-2024-0001', 1, '工作站A', 'normal', 1, 'admin', '2024-04-01 17:00:00', 500, 480, 5, 8.0, 98.5, '正常报工'),
(2, NOW(), NOW(), 1, 'RPT-2024-0002', 1, 'PO-2024-0001', 1, '工作站A', 'normal', 2, 'nhg', '2024-04-02 17:00:00', 500, 520, 3, 8.0, 105.0, '超额完成'),
(3, NOW(), NOW(), 1, 'RPT-2024-0003', 2, 'PO-2024-0002', 2, '工作站B', 'normal', 1, 'admin', '2024-04-03 17:00:00', 300, 200, 8, 8.0, 85.0, '部分报工');

-- ============================================
-- 验证查询
-- ============================================
DO $$ 
BEGIN 
  RAISE NOTICE '========== 数据插入完成 ==========';
  RAISE NOTICE '生产工单: %', (SELECT COUNT(*) FROM production_orders);
  RAISE NOTICE 'BOM: %', (SELECT COUNT(*) FROM boms);
  RAISE NOTICE 'BOM组件: %', (SELECT COUNT(*) FROM bom_items);
  RAISE NOTICE '质检IQC: %', (SELECT COUNT(*) FROM iqcs);
  RAISE NOTICE '调拨单: %', (SELECT COUNT(*) FROM wms_transfer_order);
  RAISE NOTICE '库存: %', (SELECT COUNT(*) FROM inventories);
  RAISE NOTICE '月计划: %', (SELECT COUNT(*) FROM mes_order_month);
  RAISE NOTICE '日计划: %', (SELECT COUNT(*) FROM mes_order_day);
  RAISE NOTICE '设备: %', (SELECT COUNT(*) FROM equ_equipment);
  RAISE NOTICE '告警规则: %', (SELECT COUNT(*) FROM alert_rule);
  RAISE NOTICE 'BPM模型: %', (SELECT COUNT(*) FROM bpm_process_model);
  RAISE NOTICE '安灯呼叫: %', (SELECT COUNT(*) FROM andon_call);
END $$;
