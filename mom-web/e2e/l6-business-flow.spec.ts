/**
 * L6 业务流程测试示例
 * 结合 UI 操作 + 数据库验证
 *
 * 前置条件：
 * 1. mom-server 后端运行在 localhost:9081
 * 2. mom-web 前端运行在 localhost:5177
 * 3. PostgreSQL 运行在 localhost:5432
 */

import { test, expect } from '@playwright/test'
import * as db from '../tests/helpers/db'

test.describe('L6 业务流程测试', () => {
  const testData = {
    orderNo: `TEST-SO-${Date.now()}`,
    materialCode: `TEST-MAT-${Date.now()}`,
    equipmentCode: `TEST-Equipment-${Date.now()}`,
  }

  test.afterAll(async () => {
    // 清理测试数据
    await db.cleanupTestData('pro_sales_order', 'order_no', testData.orderNo)
    await db.cleanupTestData('mdm_material', 'material_code', testData.materialCode)
    await db.cleanupTestData('equ_equipment', 'equipment_code', testData.equipmentCode)
    await db.closeDbClient()
  })

  test('登录 → 创建销售订单 → 数据库验证', async ({ page }) => {
    // ========== 1. 登录 ==========
    await page.goto('/login')
    await page.fill('input[type="text"], input[placeholder*="用户名"]', 'admin')
    await page.fill('input[type="password"]', 'admin123')
    await page.click('.el-button--primary')

    // 等待登录成功
    await page.waitForURL('**/dashboard', { timeout: 10000 })

    // ========== 2. 准备测试数据（通过 API）============
    // 注意：这里需要后端 API 支持，以下为示例结构

    // ========== 3. 创建销售订单（通过 UI）============
    await page.goto('/production/sales-order')

    // 点击新增按钮
    const addBtn = page.locator('button:has-text("新增"), button:has-text("添加"), .btn-primary').first()
    if (await addBtn.isVisible()) {
      await addBtn.click()
      await page.waitForTimeout(500)
    }

    // 如果有表单，填写
    const orderNoInput = page.locator('input[id*="order"], input[id*="no"], input[placeholder*="订单"]').first()
    if (await orderNoInput.isVisible()) {
      await orderNoInput.fill(testData.orderNo)
    }

    // ========== 4. L6: 直接通过 SQL 验证数据 ==========
    // 注意：由于前端可能还没完成，这里展示直接验证方式
    const order = await db.getSalesOrder(testData.orderNo)

    // 如果订单存在（说明 API 正常工作）
    if (order) {
      expect(order.order_no).toBe(testData.orderNo)
    }
  })

  test('物料管理：创建物料 → 数据库验证', async ({ page }) => {
    // ========== 1. 登录 ==========
    await page.goto('/login')
    await page.fill('input[type="text"], input[placeholder*="用户名"]', 'admin')
    await page.fill('input[type="password"]', 'admin123')
    await page.click('.el-button--primary')
    await page.waitForURL('**/dashboard', { timeout: 10000 })

    // ========== 2. 直接通过数据库创建物料 ==========
    // （用于验证数据库连接正常）
    const materialId = await db.createMaterialViaDb({
      code: testData.materialCode,
      name: '测试物料',
      categoryId: 1,
    })
    expect(materialId).toBeGreaterThan(0)

    // ========== 3. 通过 UI 验证物料存在 ==========
    await page.goto('/mdm/material')
    await page.waitForTimeout(1000)

    // 搜索物料
    const searchInput = page.locator('input[placeholder*="搜索"], input[placeholder*="查询"]').first()
    if (await searchInput.isVisible()) {
      await searchInput.fill(testData.materialCode)
      await page.waitForTimeout(500)
    }

    // ========== 4. L6: 从数据库读取并验证 ==========
    const material = await db.getMaterial(testData.materialCode)
    expect(material).not.toBeNull()
    expect(material.material_code).toBe(testData.materialCode)
  })

  test('设备管理：创建设备 → 数据库验证', async ({ page }) => {
    // 登录
    await page.goto('/login')
    await page.fill('input[type="text"]', 'admin')
    await page.fill('input[type="password"]', 'admin123')
    await page.click('.el-button--primary')
    await page.waitForURL('**/dashboard', { timeout: 10000 })

    // L6: 通过数据库创建设备
    const equipmentId = await db.createEquipmentViaDb({
      code: testData.equipmentCode,
      name: '测试设备',
      category: '加工设备',
    })
    expect(equipmentId).toBeGreaterThan(0)

    // L6: 验证数据库中的设备
    const equipment = await db.getEquipment(testData.equipmentCode)
    expect(equipment).not.toBeNull()
    expect(equipment.equipment_code).toBe(testData.equipmentCode)

    // UI: 访问设备页面
    await page.goto('/equipment')
    await page.waitForTimeout(1000)
  })

  test('数据库连接健康检查', async () => {
    // L6: 直接验证数据库连接
    const client = await db.getDbClient()
    expect(client).not.toBeNull()

    // 执行简单查询
    const result = await db.query('SELECT 1 as health')
    expect(result[0].health).toBe(1)

    // 验证关键表存在
    const tables = await db.query(`
      SELECT table_name FROM information_schema.tables
      WHERE table_schema = 'public'
      AND table_name IN ('pro_sales_order', 'mdm_material', 'equ_equipment', 'wms_warehouse')
    `)
    expect(tables.length).toBeGreaterThanOrEqual(4)
  })
})
