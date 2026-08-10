// ==================== MES 制造执行系统 ====================

export interface ProductionOrder {
  id: string;
  orderNo: string;
  materialCode: string;
  materialName: string;
  spec: string;
  quantity: number;
  completedQty: number;
  unit: string;
  status: OrderStatus;
  planStartDate: string;
  planEndDate: string;
  actualStartDate?: string;
  actualEndDate?: string;
  workCenter: string;
  priority: number;
  bomVersion: string;
  routingVersion: string;
  createdAt: string;
}

export type OrderStatus =
  | 'DRAFT'
  | 'RELEASED'
  | 'IN_PROGRESS'
  | 'COMPLETED'
  | 'CLOSED'
  | 'ON_HOLD';

export interface DispatchRecord {
  id: string;
  orderId: string;
  operationSeq: number;
  operationName: string;
  workCenter: string;
  machineId?: string;
  workerName: string;
  planStart: string;
  planEnd: string;
  actualStart?: string;
  actualEnd?: string;
  quantity: number;
  goodQty: number;
  defectQty: number;
  status: string;
}

export interface JobReport {
  id: string;
  dispatchId: string;
  operationName: string;
  workerName: string;
  machineId?: string;
  startTime: string;
  endTime?: string;
  goodQty: number;
  defectQty: number;
  defectCodes: string[];
  remark?: string;
}

export interface OrderCompletion {
  id: string;
  orderId: string;
  totalQty: number;
  goodQty: number;
  defectQty: number;
  scrapQty: number;
  passRate: number;
  completedAt: string;
  inspectorName: string;
}

export interface OrderTimeline {
  id: string;
  orderId: string;
  event: string;
  description: string;
  operator: string;
  timestamp: string;
}

// ==================== QMS 质量管理系统 ====================

export interface InspectionSheet {
  id: string;
  sheetNo: string;
  type: InspectionType;
  orderNo: string;
  materialCode: string;
  materialName: string;
  supplier?: string;
  lotNo: string;
  batchQty: number;
  sampleQty: number;
  status: InspectionStatus;
  inspector?: string;
  startTime?: string;
  endTime?: string;
  result: InspectionResult;
  aqlLevel?: string;
  createdAt: string;
}

export type InspectionType = 'IQC' | 'IPQC' | 'FQC' | 'OQC';
export type InspectionStatus = 'PENDING' | 'IN_PROGRESS' | 'COMPLETED' | 'CLOSED';
export type InspectionResult = 'PASS' | 'FAIL' | 'PENDING';

export interface InspectionCharacteristic {
  id: string;
  sheetId: string;
  seq: number;
  itemName: string;
  spec: string;
  upperLimit?: number;
  lowerLimit?: number;
  unit: string;
  inspectionMethod: string;
  sampleSize: number;
  result: string;
  isQualified: boolean;
}

export interface InspectionRecord {
  id: string;
  characteristicId: string;
  seq: number;
  measuredValue: number | string;
  isQualified: boolean;
  remark?: string;
}

export interface Ncr {
  id: string;
  ncrNo: string;
  source: string;
  orderNo: string;
  materialCode: string;
  materialName: string;
  defectDescription: string;
  severity: NcrSeverity;
  status: NcrStatus;
  quantity: number;
  inspector: string;
  foundAt: string;
  disposition: string;
  closedAt?: string;
}

export type NcrSeverity = 'CRITICAL' | 'MAJOR' | 'MINOR';
export type NcrStatus = 'OPEN' | 'DISPOSITION' | 'IN_REWORK' | 'VERIFIED' | 'CLOSED';

export interface NcrAction {
  id: string;
  ncrId: string;
  seq: number;
  action: string;
  operator: string;
  result?: string;
  createdAt: string;
}

// ==================== EAM 设备管理系统 ====================

export interface Equipment {
  id: string;
  equipmentNo: string;
  name: string;
  model: string;
  type: string;
  workCenter: string;
  status: EquipmentStatus;
  oee: number;
  availability: number;
  performance: number;
  quality: number;
  totalRunTime: number;
  totalDowntime: number;
  lastMaintenance?: string;
  nextMaintenance?: string;
  installDate: string;
  manufacturer: string;
}

export type EquipmentStatus = 'RUNNING' | 'IDLE' | 'MAINTENANCE' | 'REPAIR' | 'STOPPED';

export interface MaintenancePlan {
  id: string;
  planNo: string;
  equipmentId: string;
  type: MaintenanceType;
  description: string;
  plannedDate: string;
  actualDate?: string;
  status: MaintenanceStatus;
  technician?: string;
  nextPlanDate?: string;
}

export type MaintenanceType = 'ROUTINE' | 'INSPECTION' | 'OVERHAUL' | 'PREDICTIVE';
export type MaintenanceStatus = 'PLANNED' | 'IN_PROGRESS' | 'COMPLETED' | 'OVERDUE';

export interface RepairOrder {
  id: string;
  repairNo: string;
  equipmentId: string;
  equipmentName: string;
  faultDesc: string;
  faultType: string;
  severity: 'CRITICAL' | 'MAJOR' | 'MINOR';
  status: RepairStatus;
  reporter: string;
  reportedAt: string;
  technician?: string;
  startTime?: string;
  endTime?: string;
  rootCause?: string;
  solution?: string;
  downtime: number;
}

export type RepairStatus =
  | 'REPORTED'
  | 'ASSIGNED'
  | 'IN_PROGRESS'
  | 'COMPLETED'
  | 'CONFIRMED';

export interface DowntimeRecord {
  id: string;
  equipmentId: string;
  startTime: string;
  endTime?: string;
  duration?: number;
  reason: string;
  category: string;
  reportedBy: string;
}

// ==================== WMS 仓储管理系统 ====================

export interface Warehouse {
  id: string;
  code: string;
  name: string;
  location: string;
  type: string;
  capacity: number;
  usedCapacity: number;
  status: string;
}

export interface InventoryBalance {
  id: string;
  materialCode: string;
  materialName: string;
  spec: string;
  unit: string;
  warehouse: string;
  location: string;
  batchNo: string;
  quantity: number;
  lockedQty: number;
  availableQty: number;
  status: InventoryStatus;
  lastUpdate: string;
}

export type InventoryStatus = 'NORMAL' | 'LOCKED' | 'ON_HOLD' | 'SCRAP';

export interface ReceiveOrder {
  id: string;
  receiveNo: string;
  orderNo: string;
  supplier: string;
  materialCode: string;
  materialName: string;
  expectedQty: number;
  receivedQty: number;
  unit: string;
  status: string;
  receiveDate: string;
  inspector?: string;
  inspectionResult?: string;
}

export interface DeliveryOrder {
  id: string;
  deliveryNo: string;
  orderNo: string;
  customer: string;
  materialCode: string;
  materialName: string;
  quantity: number;
  unit: string;
  status: string;
  planDate: string;
  actualDate?: string;
}

// ==================== APS 高级排程 ====================

export interface MpsPlan {
  id: string;
  planNo: string;
  materialCode: string;
  materialName: string;
  quantity: number;
  startDate: string;
  endDate: string;
  status: string;
  frozen: boolean;
}

export interface ScheduleTask {
  id: string;
  orderNo: string;
  operationName: string;
  workCenter: string;
  startTime: string;
  endTime: string;
  status: string;
  progress: number;
}

// ==================== 追溯 ====================

export interface TraceNode {
  id: string;
  name: string;
  type: 'MATERIAL' | 'ORDER' | 'BATCH' | 'SUPPLIER' | 'CUSTOMER';
  label: string;
  children?: TraceNode[];
}

export interface TraceResult {
  serialNo?: string;
  batchNo?: string;
  material: string;
  forward: TraceNode;
  backward: TraceNode;
}

// ==================== Andon 安灯 ====================

export interface AndonCall {
  id: string;
  callNo: string;
  workCenter: string;
  machineId?: string;
  type: AndonType;
  severity: 'CRITICAL' | 'WARNING' | 'INFO';
  description: string;
  status: AndonStatus;
  caller: string;
  calledAt: string;
  respondedAt?: string;
  resolvedAt?: string;
  responder?: string;
}

export type AndonType = 'MATERIAL' | 'EQUIPMENT' | 'QUALITY' | 'SAFETY' | 'PROCESS';
export type AndonStatus = 'ACTIVE' | 'RESPONDED' | 'RESOLVED' | 'TIMEOUT';

export interface Alert {
  id: string;
  title: string;
  message: string;
  severity: 'CRITICAL' | 'WARNING' | 'INFO';
  source: string;
  status: 'ACTIVE' | 'ACKNOWLEDGED' | 'RESOLVED';
  createdAt: string;
  acknowledgedAt?: string;
}

// ==================== 通用 ====================

export interface PaginatedResponse<T> {
  data: T[];
  total: number;
  page: number;
  pageSize: number;
  totalPages: number;
}

export interface ApiResponse<T> {
  success: boolean;
  data: T;
  message?: string;
}

export interface DashboardStats {
  totalOrders: number;
  qualityPassRate: number;
  oeeAverage: number;
  activeAlarms: number;
  orderStatusDistribution: { status: string; count: number }[];
  recentAlarms: Alert[];
  productionTrend: { date: string; planned: number; actual: number }[];
}
