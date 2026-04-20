-- 清空并重置菜单表
TRUNCATE TABLE sys_menu RESTART IDENTITY;

-- 重新插入完整菜单数据
INSERT INTO sys_menu (id, created_at, updated_at, tenant_id, parent_id, menu_name, menu_type, path, component, perms, icon, sort, visible, status, is_frame, is_cache) VALUES
-- 一级菜单
(1, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 0, '首页', 'C', '/dashboard', NULL, '', 'HomeFilled', 1, 1, 1, 0, 0),
(2, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 0, '系统管理', 'M', '/system', NULL, '', 'Setting', 2, 1, 1, 0, 0),
(3, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 0, '主数据', 'M', '/mdm', NULL, '', 'Box', 3, 1, 1, 0, 0),
(4, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 0, '生产执行', 'M', '/production', NULL, '', 'List', 4, 1, 1, 0, 0),
(5, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 0, '设备管理', 'M', '/equipment', NULL, '', 'Monitor', 5, 1, 1, 0, 0),
(6, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 0, '仓储管理', 'M', '/wms', NULL, '', 'House', 6, 1, 1, 0, 0),
(7, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 0, '质量管理', 'M', '/quality', NULL, '', 'CircleCheck', 7, 1, 1, 0, 0),
(8, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 0, 'APS计划', 'M', '/aps', NULL, '', 'Calendar', 8, 1, 1, 0, 0),
(9, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 0, '追溯管理', 'M', '/trace', NULL, '', 'Search', 9, 1, 1, 0, 0),
(10, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 0, '能源管理', 'M', '/energy', NULL, '', 'Lightning', 10, 1, 1, 0, 0),

-- 系统管理子菜单 (parent_id=2)
(101, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, '用户管理', 'C', '/system/user', 'system/UserList.vue', 'system:user:list', 'User', 1, 1, 1, 0, 0),
(102, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, '角色管理', 'C', '/system/role', 'system/RoleList.vue', 'system:role:list', 'Key', 2, 1, 1, 0, 0),
(103, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, '菜单管理', 'C', '/system/menu', 'system/MenuList.vue', 'system:menu:list', 'Menu', 3, 1, 1, 0, 0),
(104, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, '部门管理', 'C', '/system/dept', 'system/DeptList.vue', 'system:dept:list', 'OfficeBuilding', 4, 1, 1, 0, 0),
(105, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, '字典管理', 'C', '/system/dict', 'system/DictList.vue', 'system:dict:list', 'Document', 5, 1, 1, 0, 0),
(106, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, '岗位管理', 'C', '/system/post', 'system/PostList.vue', 'system:post:list', 'Postcard', 6, 1, 1, 0, 0),
(107, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, '租户管理', 'C', '/system/tenant', 'system/TenantList.vue', 'system:tenant:list', 'Building', 7, 1, 1, 0, 0),
(108, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, '登录日志', 'C', '/system/login-log', 'system/LoginLogList.vue', 'system:loginlog:list', 'Key', 8, 1, 1, 0, 0),
(109, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, '操作日志', 'C', '/system/oper-log', 'system/OperLogList.vue', 'system:operlog:list', 'Document', 9, 1, 1, 0, 0),
(110, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, '系统配置', 'C', '/system/config', 'system/SystemConfig.vue', 'system:config:list', 'Setting', 10, 1, 1, 0, 0),

-- 主数据子菜单 (parent_id=3)
(201, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 3, '物料管理', 'C', '/mdm/material', 'mdm/MaterialList.vue', 'mdm:material:list', 'Box', 1, 1, 1, 0, 0),
(202, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 3, '车间管理', 'C', '/mdm/workshop', 'mdm/WorkshopList.vue', 'mdm:workshop:list', 'OfficeBuilding', 2, 1, 1, 0, 0),
(203, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 3, '生产线管理', 'C', '/mdm/line', 'mdm/LineList.vue', 'mdm:line:list', 'Connection', 3, 1, 1, 0, 0),
(204, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 3, '工位管理', 'C', '/mdm/workstation', 'mdm/WorkstationList.vue', 'mdm:workstation:list', 'Grid', 4, 1, 1, 0, 0),
(205, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 3, '班次管理', 'C', '/mdm/shift', 'mdm/ShiftList.vue', 'mdm:shift:list', 'Clock', 5, 1, 1, 0, 0),
(206, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 3, 'BOM管理', 'C', '/mdm/bom', 'mdm/BomList.vue', 'mdm:bom:list', 'Files', 6, 1, 1, 0, 0),
(207, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 3, '工序管理', 'C', '/mdm/operation', 'mdm/OperationList.vue', 'mdm:operation:list', 'Operation', 7, 1, 1, 0, 0),
(208, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 3, '班次定义', 'C', '/mdm/mdm-shift', 'mdm/ShiftList.vue', 'mdm:mdmshift:list', 'Clock', 8, 1, 1, 0, 0),

-- 生产执行子菜单 (parent_id=4)
(301, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 4, '生产工单', 'C', '/production/order', 'production/ProductionOrderList.vue', 'production:order:list', 'List', 1, 1, 1, 0, 0),
(302, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 4, '销售订单', 'C', '/production/sales-order', 'production/SalesOrderList.vue', 'production:salesorder:list', 'Document', 2, 1, 1, 0, 0),
(303, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 4, '生产报工', 'C', '/production/report', 'production/ReportList.vue', 'production:report:list', 'DocumentCheck', 3, 1, 1, 0, 0),
(304, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 4, '派工', 'C', '/production/dispatch', 'production/DispatchList.vue', 'production:dispatch:list', 'Tickets', 4, 1, 1, 0, 0),

-- 设备管理子菜单 (parent_id=5)
(401, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 5, '设备台账', 'C', '/equipment', 'equipment/EquipmentList.vue', 'equipment:list:list', 'Monitor', 1, 1, 1, 0, 0),
(402, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 5, '设备点检', 'C', '/equipment/check', 'equipment/CheckList.vue', 'equipment:check:list', 'Check', 2, 1, 1, 0, 0),
(403, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 5, '设备保养', 'C', '/equipment/maintenance', 'equipment/MaintenanceList.vue', 'equipment:maintenance:list', 'Tools', 3, 1, 1, 0, 0),
(404, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 5, '设备维修', 'C', '/equipment/repair', 'equipment/RepairList.vue', 'equipment:repair:list', 'Tool', 4, 1, 1, 0, 0),
(405, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 5, '备件管理', 'C', '/equipment/spare', 'equipment/SparePartList.vue', 'equipment:spare:list', 'Box', 5, 1, 1, 0, 0),

-- 仓储管理子菜单 (parent_id=6)
(501, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 6, '仓库管理', 'C', '/wms/warehouse', 'wms/WarehouseList.vue', 'wms:warehouse:list', 'House', 1, 1, 1, 0, 0),
(502, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 6, '库位管理', 'C', '/wms/location', 'wms/LocationList.vue', 'wms:location:list', 'Location', 2, 1, 1, 0, 0),
(503, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 6, '库存管理', 'C', '/wms/inventory', 'wms/InventoryList.vue', 'wms:inventory:list', 'Box', 3, 1, 1, 0, 0),

-- 质量管理子菜单 (parent_id=7)
(601, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 7, 'IQC检验', 'C', '/quality/iqc', 'quality/IQCList.vue', 'quality:iqc:list', 'CircleCheck', 1, 1, 1, 0, 0),
(602, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 7, 'IPQC检验', 'C', '/quality/ipqc', 'quality/IPQCList.vue', 'quality:ipqc:list', 'CircleCheck', 2, 1, 1, 0, 0),
(603, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 7, 'FQC检验', 'C', '/quality/fqc', 'quality/FQCList.vue', 'quality:fqc:list', 'CircleCheck', 3, 1, 1, 0, 0),
(604, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 7, 'OQC检验', 'C', '/quality/oqc', 'quality/OQCList.vue', 'quality:oqc:list', 'CircleCheck', 4, 1, 1, 0, 0),
(605, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 7, '缺陷代码', 'C', '/quality/defect-code', 'quality/DefectCodeList.vue', 'quality:defectcode:list', 'Warning', 5, 1, 1, 0, 0),
(606, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 7, '缺陷记录', 'C', '/quality/defect-record', 'quality/DefectRecordList.vue', 'quality:defectrecord:list', 'Document', 6, 1, 1, 0, 0),
(607, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 7, 'NCR处理', 'C', '/quality/ncr', 'quality/NCRList.vue', 'quality:ncr:list', 'Close', 7, 1, 1, 0, 0),
(608, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 7, 'SPC数据', 'C', '/quality/spc', 'quality/SPCList.vue', 'quality:spc:list', 'DataLine', 8, 1, 1, 0, 0),

-- APS计划子菜单 (parent_id=8)
(701, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 8, 'MPS计划', 'C', '/aps/mps', 'aps/MPSList.vue', 'aps:mps:list', 'Calendar', 1, 1, 1, 0, 0),
(702, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 8, 'MRP计划', 'C', '/aps/mrp', 'aps/MRPList.vue', 'aps:mrp:list', 'Grid', 2, 1, 1, 0, 0),
(703, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 8, '排程计划', 'C', '/aps/schedule', 'aps/ScheduleList.vue', 'aps:schedule:list', 'List', 3, 1, 1, 0, 0),

-- 追溯管理子菜单 (parent_id=9)
(801, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 9, '追溯查询', 'C', '/trace/query', 'trace/TraceQuery.vue', 'trace:query:list', 'Search', 1, 1, 1, 0, 0),
(802, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 9, '安东呼叫', 'C', '/trace/andon', 'trace/AndonCall.vue', 'trace:andon:list', 'Bell', 2, 1, 1, 0, 0),

-- 能源管理子菜单 (parent_id=10)
(901, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 10, '能源监控', 'C', '/energy/monitor', 'energy/EnergyMonitor.vue', 'energy:monitor:list', 'Lightning', 1, 1, 1, 0, 0);
