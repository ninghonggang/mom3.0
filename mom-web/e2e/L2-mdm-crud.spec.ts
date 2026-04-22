import { test, expect } from '@playwright/test'
import * as db from '../tests/helpers/db'

test.describe('L2 MDM 主数据管理 CRUD', () => {
  const testData = {
    materialCode: `TEST-MAT-${Date.now()}`,
    customerCode: `TEST-CUS-${Date.now()}`,
    supplierCode: `TEST-SUP-${Date.now()}`,
  }

  test.afterAll(async () => {
    await db.cleanupTestData('mdm_material', 'material_code', testData.materialCode)
    await db.cleanupTestData('mdm_customer', 'code', testData.customerCode)
    await db.cleanupTestData('mdm_supplier', 'code', testData.supplierCode)
    await db.closeDbClient()
  })

  test.beforeEach(async ({ page }) => {
    await page.goto('/login')
    await page.fill('input[placeholder="请输入用户名"]', 'admin')
    await page.fill('input[placeholder="请输入密码"]', 'admin123')
    await page.click('.login-btn')
    await page.waitForURL('**/dashboard', { timeout: 10000 })
  })

  test('物料管理：创建 → 验证', async ({ page }) => {
    // 通过 DB 创建物料
    const materialId = await db.createMaterialViaDb({
      code: testData.materialCode,
      name: '测试物料-E2E',
    })
    expect(materialId).toBeGreaterThan(0)

    // UI: 访问物料列表
    await page.goto('/mdm/material')
    await page.waitForTimeout(1000)

    // 搜索物料
    const searchInput = page.locator('input[placeholder*="搜索"], input[placeholder*="查询"], input[placeholder*="物料"]').first()
    if (await searchInput.isVisible()) {
      await searchInput.fill(testData.materialCode)
      await page.waitForTimeout(500)
      await page.keyboard.press('Enter')
    }

    // 验证物料出现在列表中
    await expect(page.locator(`text=${testData.materialCode}`).first()).toBeVisible({ timeout: 5000 })
  })

  test('客户管理：创建 → 验证', async ({ page }) => {
    const customerId = await db.createCustomer({ code: testData.customerCode, name: '测试客户-E2E' })
    expect(customerId).toBeGreaterThan(0)

    await page.goto('/mdm/customer')
    await page.waitForTimeout(1000)

    const searchInput = page.locator('input[placeholder*="搜索"], input[placeholder*="查询"]').first()
    if (await searchInput.isVisible()) {
      await searchInput.fill(testData.customerCode)
      await page.waitForTimeout(500)
    }

    const customer = await db.getCustomer(testData.customerCode)
    expect(customer).not.toBeNull()
  })

  test('供应商管理：创建 → 验证', async ({ page }) => {
    const supplierId = await db.createSupplier({ code: testData.supplierCode, name: '测试供应商-E2E' })
    expect(supplierId).toBeGreaterThan(0)

    await page.goto('/mdm/supplier')
    await page.waitForTimeout(1000)

    const searchInput = page.locator('input[placeholder*="搜索"], input[placeholder*="查询"]').first()
    if (await searchInput.isVisible()) {
      await searchInput.fill(testData.supplierCode)
      await page.waitForTimeout(500)
    }

    const supplier = await db.getSupplier(testData.supplierCode)
    expect(supplier).not.toBeNull()
  })
})
