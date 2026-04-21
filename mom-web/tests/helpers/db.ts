/**
 * 数据库验证辅助工具 - L6 测试用
 * 连接 mom3_db 验证业务流程数据
 */

import { Client } from 'pg'

const dbConfig = {
  host: 'localhost',
  port: 5432,
  user: 'postgres',
  password: 'postgres',
  database: 'mom3_db',
}

let client: Client | null = null

export async function getDbClient(): Promise<Client> {
  if (!client) {
    client = new Client(dbConfig)
    await client.connect()
  }
  return client
}

export async function closeDbClient(): Promise<void> {
  if (client) {
    await client.end()
    client = null
  }
}

// ============ 基础操作 ============

export async function query(sql: string, params?: any[]): Promise<any[]> {
  const c = await getDbClient()
  const result = await c.query(sql, params)
  return result.rows
}

// ============ 销售订单 (pro_sales_order) ============

export async function createSalesOrderViaDb(data: {
  orderNo: string
  customerId: number
  quantity: number
  deliveryDate?: string
}): Promise<number> {
  const sql = `
    INSERT INTO pro_sales_order (tenant_id, order_no, customer_id, quantity, delivery_date, status, created_at)
    VALUES (1, $1, $2, $3, $4, 1, NOW())
    RETURNING id
  `
  const result = await query(sql, [
    data.orderNo,
    data.customerId,
    data.quantity,
    data.deliveryDate || new Date().toISOString().split('T')[0],
  ])
  return Number(result[0].id)
}

export async function getSalesOrder(orderNo: string): Promise<any> {
  const rows = await query(
    'SELECT * FROM pro_sales_order WHERE order_no = $1',
    [orderNo]
  )
  return rows[0] || null
}

export async function updateSalesOrderStatus(orderNo: string, status: string): Promise<void> {
  await query(
    'UPDATE pro_sales_order SET status = $1, updated_at = NOW() WHERE order_no = $2',
    [status, orderNo]
  )
}

// ============ 生产工单 (pro_production_order) ============

export async function createProductionOrderViaDb(data: {
  orderNo: string
  salesOrderNo?: string
  quantity: number
}): Promise<number> {
  const sql = `
    INSERT INTO pro_production_order (order_no, sales_order_no, quantity, status, created_at)
    VALUES ($1, $2, $3, 'planned', NOW())
    RETURNING id
  `
  const result = await query(sql, [
    data.orderNo,
    data.salesOrderNo,
    data.quantity,
  ])
  return Number(result[0].id)
}

export async function getProductionOrder(orderNo: string): Promise<any> {
  const rows = await query(
    'SELECT * FROM pro_production_order WHERE order_no = $1',
    [orderNo]
  )
  return rows[0] || null
}

// ============ 物料 (mdm_material) ============

export async function createMaterialViaDb(data: {
  code: string
  name: string
  categoryId?: number
}): Promise<number> {
  const sql = `
    INSERT INTO mdm_material (tenant_id, material_code, material_name, category_id, status, created_at)
    VALUES (1, $1, $2, $3, 1, NOW())
    RETURNING id
  `
  const result = await query(sql, [data.code, data.name, data.categoryId || null])
  return Number(result[0].id)
}

export async function getMaterial(code: string): Promise<any> {
  const rows = await query(
    'SELECT * FROM mdm_material WHERE material_code = $1',
    [code]
  )
  return rows[0] || null
}

// ============ 设备 (equ_equipment) ============

export async function createEquipmentViaDb(data: {
  code: string
  name: string
  equipmentType?: string
}): Promise<number> {
  const sql = `
    INSERT INTO equ_equipment (tenant_id, equipment_code, equipment_name, equipment_type, status, created_at)
    VALUES (1, $1, $2, $3, 1, NOW())
    RETURNING id
  `
  const result = await query(sql, [data.code, data.name, data.equipmentType || null])
  return Number(result[0].id)
}

export async function getEquipment(code: string): Promise<any> {
  const rows = await query(
    'SELECT * FROM equ_equipment WHERE equipment_code = $1',
    [code]
  )
  return rows[0] || null
}

// ============ 仓库 (wms_warehouse) ============

export async function createWarehouseViaDb(data: {
  code: string
  name: string
}): Promise<number> {
  const sql = `
    INSERT INTO wms_warehouse (tenant_id, warehouse_code, warehouse_name, status, created_at)
    VALUES (1, $1, $2, 1, NOW())
    RETURNING id
  `
  const result = await query(sql, [data.code, data.name])
  return Number(result[0].id)
}

export async function getWarehouse(code: string): Promise<any> {
  const rows = await query(
    'SELECT * FROM wms_warehouse WHERE warehouse_code = $1',
    [code]
  )
  return rows[0] || null
}

// ============ IQC 检验 (iqc) ============

export async function createIqcViaDb(data: {
  iqcNo: string
  supplierId: number
}): Promise<number> {
  const sql = `
    INSERT INTO iqcs (iqc_no, supplier_id, status, inspection_type, created_at)
    VALUES ($1, $2, 'pending', 'normal', NOW())
    RETURNING id
  `
  const result = await query(sql, [data.iqcNo, data.supplierId])
  return Number(result[0].id)
}

// ============ 清理测试数据 ============

export async function cleanupTestData(table: string, column: string, value: string): Promise<void> {
  await query(`DELETE FROM ${table} WHERE ${column} = $1`, [value])
}
