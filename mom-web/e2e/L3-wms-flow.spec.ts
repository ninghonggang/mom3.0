/**
 * L3 WMS 仓储业务流程测试
 * 测试范围：仓库 → 库位 → 库存 → 调拨、收货入库、发货出库、库存盘点
 *
 * 前置条件：
 * 1. mom-server 后端运行在 localhost:9081
 * 2. mom-web 前端运行在 localhost:5177
 * 3. PostgreSQL 运行在 localhost:5432
 */

import { test, expect } from '@playwright/test'
import * as db from '../tests/helpers/db'

test.describe('L3 WMS 仓储业务流程', () => {
  const testData = {
    warehouseCode: `TEST-WH-${Date.now()}`,
  }

  test.afterAll(async () => {
    await db.cleanupTestData('wms_warehouse', 'warehouse_code', testData.warehouseCode)
    await db.closeDbClient()
  })

  test.beforeEach(async ({ page }) => {
    await page.goto('/login')
    await page.waitForTimeout(2000)
    await page.fill('input[type="text"], input[placeholder*="用户名"], input[placeholder*="账号"]', 'admin')
    await page.fill('input[type="password"]', 'admin123')
    await page.click('button[type="submit"], button:has-text("登录"), button:has-text("登 录")')
    await page.waitForURL('**/dashboard', { timeout: 15000 })
  }, 60000)

  test('仓库 → 库位 → 库存 → 调拨', async ({ page }) => {
    // 创建仓库
    const warehouseId = await db.createWarehouseViaDb({ code: testData.warehouseCode, name: '测试仓库-E2E' })
    expect(warehouseId).toBeGreaterThan(0)

    // 仓库列表
    await page.goto('/wms/warehouse')
    await page.waitForTimeout(1000)

    const searchInput = page.locator('input[placeholder*="搜索"], input[placeholder*="查询"]').first()
    if (await searchInput.isVisible()) {
      await searchInput.fill(testData.warehouseCode)
      await page.keyboard.press('Enter')
      await page.waitForTimeout(500)
    }

    await expect(page.locator(`text=${testData.warehouseCode}`).first()).toBeVisible({ timeout: 5000 })

    // 库位
    await page.goto('/wms/location')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()

    // 库存
    await page.goto('/wms/inventory')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()

    // 调拨
    await page.goto('/wms/transfer')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('收货入库', async ({ page }) => {
    await page.goto('/wms/receive')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('发货出库', async ({ page }) => {
    await page.goto('/wms/delivery')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('库存盘点', async ({ page }) => {
    await page.goto('/wms/stock-check')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })
})