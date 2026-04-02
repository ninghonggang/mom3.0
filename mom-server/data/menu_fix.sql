-- 补充缺失的菜单数据
INSERT INTO sys_menu (id, created_at, updated_at, tenant_id, parent_id, menu_name, menu_type, path, component, perms, icon, sort, visible, status, is_frame, is_cache) VALUES
-- 仓储管理 (id=5)
(5, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 0, '仓储管理', 'M', '/wms', NULL, '', 'House', 5, 0, 1, 0, 0),
-- 质量管理 (id=6)
(6, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 0, '质量管理', 'M', '/quality', NULL, '', 'CircleCheck', 6, 0, 1, 0, 0),
-- APS计划 (id=7)
(7, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 0, 'APS计划', 'M', '/aps', NULL, '', 'Calendar', 7, 0, 1, 0, 0),
-- 追溯管理 (id=8)
(8, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 0, '追溯管理', 'M', '/trace', NULL, '', 'Search', 8, 0, 1, 0, 0),

-- 系统管理子菜单 (id=101-110)
(101, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 1, '用户管理', 'C', 'user', 'system/UserList.vue', 'system:user:list', 'User', 1, 0, 1, 0, 0),
(102, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 1, '角色管理', 'C', 'role', 'system/RoleList.vue', 'system:role:list', 'Key', 2, 0, 1, 0, 0),
(103, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 1, '菜单管理', 'C', 'menu', 'system/MenuList.vue', 'system:menu:list', 'Menu', 3, 0, 1, 0, 0),
(104, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 1, '部门管理', 'C', 'dept', 'system/DeptList.vue', 'system:dept:list', 'OfficeBuilding', 4, 0, 1, 0, 0),
(105, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 1, '字典管理', 'C', 'dict', 'system/DictList.vue', 'system:dict:list', 'Document', 5, 0, 1, 0, 0),
(106, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 1, '岗位管理', 'C', 'post', 'system/PostList.vue', 'system:post:list', 'Postcard', 6, 0, 1, 0, 0),
(107, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 1, '租户管理', 'C', 'tenant', 'system/TenantList.vue', 'system:tenant:list', 'Building', 7, 0, 1, 0, 0),
(108, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 1, '登录日志', 'C', 'login-log', 'system/LoginLogList.vue', 'system:loginlog:list', 'Key', 8, 0, 1, 0, 0),
(109, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 1, '操作日志', 'C', 'oper-log', 'system/OperLogList.vue', 'system:operlog:list', 'Document', 9, 0, 1, 0, 0),
(110, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 1, '系统配置', 'C', 'config', 'system/SystemConfig.vue', 'system:config:list', 'Setting', 10, 0, 1, 0, 0),

-- 主数据子菜单 (id=201-208)
(201, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, '物料管理', 'C', 'material', 'mdm/MaterialList.vue', 'mdm:material:list', 'Box', 1, 0, 1, 0, 0),
(202, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, '车间管理', 'C', 'workshop', 'mdm/WorkshopList.vue', 'mdm:workshop:list', 'OfficeBuilding', 2, 0, 1, 0, 0),
(203, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, '生产线管理', 'C', 'line', 'mdm/LineList.vue', 'mdm:line:list', 'Connection', 3, 0, 1, 0, 0),
(204, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, '工位管理', 'C', 'workstation', 'mdm/WorkstationList.vue', 'mdm:workstation:list', 'Grid', 4, 0, 1, 0, 0),
(205, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, '班次管理', 'C', 'shift', 'mdm/ShiftList.vue', 'mdm:shift:list', 'Clock', 5, 0, 1, 0, 0),
(206, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, 'BOM管理', 'C', 'bom', 'mdm/BomList.vue', 'mdm:bom:list', 'Files', 6, 0, 1, 0, 0),
(207, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, '工序管理', 'C', 'operation', 'mdm/OperationList.vue', 'mdm:operation:list', 'Operation', 7, 0, 1, 0, 0),
(208, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 2, '班次定义', 'C', 'mdm-shift', 'mdm/ShiftList.vue', 'mdm:mdmshift:list', 'Clock', 8, 0, 1, 0, 0),

-- 生产执行子菜单 (id=301-304)
(301, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 3, '生产工单', 'C', 'order', 'production/ProductionOrderList.vue', 'production:order:list', 'List', 1, 0, 1, 0, 0),
(302, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 3, '销售订单', 'C', 'sales-order', 'production/SalesOrderList.vue', 'production:salesorder:list', 'Document', 2, 0, 1, 0, 0),
(303, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 3, '生产报工', 'C', 'report', 'production/ReportList.vue', 'production:report:list', 'DocumentCheck', 3, 0, 1, 0, 0),
(304, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 3, '派工', 'C', 'dispatch', 'production/DispatchList.vue', 'production:dispatch:list', 'Tickets', 4, 0, 1, 0, 0),

-- 设备管理子菜单 (id=401-405)
(401, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 4, '设备台账', 'C', '', 'equipment/EquipmentList.vue', 'equipment:list:list', 'Monitor', 1, 0, 1, 0, 0),
(402, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 4, '设备点检', 'C', 'check', 'equipment/CheckList.vue', 'equipment:check:list', 'Check', 2, 0, 1, 0, 0),
(403, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 4, '设备保养', 'C', 'maintenance', 'equipment/MaintenanceList.vue', 'equipment:maintenance:list', 'Tools', 3, 0, 1, 0, 0),
(404, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 4, '设备维修', 'C', 'repair', 'equipment/RepairList.vue', 'equipment:repair:list', 'Tool', 4, 0, 1, 0, 0),
(405, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 4, '备件管理', 'C', 'spare', 'equipment/SparePartList.vue', 'equipment:spare:list', 'Box', 5, 0, 1, 0, 0),

-- 仓储管理子菜单 (id=501-503)
(501, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 5, '仓库管理', 'C', 'warehouse', 'wms/WarehouseList.vue', 'wms:warehouse:list', 'House', 1, 0, 1, 0, 0),
(502, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 5, '库位管理', 'C', 'location', 'wms/LocationList.vue', 'wms:location:list', 'Location', 2, 0, 1, 0, 0),
(503, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 5, '库存管理', 'C', 'inventory', 'wms/InventoryList.vue', 'wms:inventory:list', 'Box', 3, 0, 1, 0, 0),

-- 质量管理子菜单 (id=601-608)
(601, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 6, 'IQC检验', 'C', 'iqc', 'quality/IQCList.vue', 'quality:iqc:list', 'CircleCheck', 1, 0, 1, 0, 0),
(602, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 6, 'IPQC检验', 'C', 'ipqc', 'quality/IPQCList.vue', 'quality:ipqc:list', 'CircleCheck', 2, 0, 1, 0, 0),
(603, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 6, 'FQC检验', 'C', 'fqc', 'quality/FQCList.vue', 'quality:fqc:list', 'CircleCheck', 3, 0, 1, 0, 0),
(604, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 6, 'OQC检验', 'C', 'oqc', 'quality/OQCList.vue', 'quality:oqc:list', 'CircleCheck', 4, 0, 1, 0, 0),
(605, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 6, '缺陷代码', 'C', 'defect-code', 'quality/DefectCodeList.vue', 'quality:defectcode:list', 'Warning', 5, 0, 1, 0, 0),
(606, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 6, '缺陷记录', 'C', 'defect-record', 'quality/DefectRecordList.vue', 'quality:defectrecord:list', 'Document', 6, 0, 1, 0, 0),
(607, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 6, 'NCR处理', 'C', 'ncr', 'quality/NCRList.vue', 'quality:ncr:list', 'Close', 7, 0, 1, 0, 0),
(608, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 6, 'SPC数据', 'C', 'spc', 'quality/SPCList.vue', 'quality:spc:list', 'DataLine', 8, 0, 1, 0, 0),

-- APS计划子菜单 (id=701-703)
(701, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 7, 'MPS计划', 'C', 'mps', 'aps/MPSList.vue', 'aps:mps:list', 'Calendar', 1, 0, 1, 0, 0),
(702, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 7, 'MRP计划', 'C', 'mrp', 'aps/MRPList.vue', 'aps:mrp:list', 'Grid', 2, 0, 1, 0, 0),
(703, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 7, '排程计划', 'C', 'schedule', 'aps/ScheduleList.vue', 'aps:schedule:list', 'List', 3, 0, 1, 0, 0),

-- 追溯管理子菜单 (id=801-802)
(801, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 8, '追溯查询', 'C', 'query', 'trace/TraceQuery.vue', 'trace:query:list', 'Search', 1, 0, 1, 0, 0),
(802, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 8, '安东呼叫', 'C', 'andon', 'trace/AndonCall.vue', 'trace:andon:list', 'Bell', 2, 0, 1, 0, 0),

-- 能源管理子菜单 (id=901)
(901, '2026-01-01 08:00:00', '2026-01-01 08:00:00', 1, 9, '能源监控', 'C', 'monitor', 'energy/EnergyMonitor.vue', 'energy:monitor:list', 'Lightning', 1, 0, 1, 0, 0)

ON CONFLICT (id) DO NOTHING;
