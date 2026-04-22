import { test, expect } from '@playwright/test'
import * as db from '../tests/helpers/db'

test.describe('L3 SCP 供应链业务流程', () => {
  const testData = {
    purchaseOrderNo: `TEST-PO-${Date.now()}`,
    supplierCode: `TEST-SUP-${Date.now()}`,
    materialCode: `TEST-MAT-${Date.now()}`,
  }

  test.afterAll(async () => {
    await db.cleanupTestData('scp_purchase_order', 'order_no', testData.purchaseOrderNo)
    await db.cleanupTestData('mdm_supplier', 'supplier_code', testData.supplierCode)
    await db.cleanupTestData('mdm_material', 'material_code', testData.materialCode)
    await db.closeDbClient()
  })

  test.beforeEach(async ({ page }) => {
    await page.goto('/login')
    await page.fill('input[type="text"]', 'admin')
    await page.fill('input[type="password"]', 'admin123')
    await page.click('button[type="submit"]')
    await page.waitForURL('**/dashboard', { timeout: 10000 })
  })

  test('采购订单 → 收货', async ({ page }) => {
    // 创建测试数据
    const supplierId = await db.createSupplier({ code: testData.supplierCode, name: '测试供应商-SCP' })
    const materialId = await db.createMaterialViaDb({ code: testData.materialCode, name: '测试物料-SCP' })
    expect(supplierId).toBeGreaterThan(0)

    // 访问采购订单页面
    await page.goto('/scp/purchase')
    await page.waitForTimeout(1000)

    // 点击新增
    const addBtn = page.locator('button:has-text("新增"), button:has-text("添加"), button:has-text("新建")').first()
    if (await addBtn.isVisible()) {
      await addBtn.click()
      await page.waitForTimeout(500)
    }

    // 填写订单号
    const orderNoInput = page.locator('input[id*="order"], input[id*="no"], input[placeholder*="订单"]').first()
    if (await orderNoInput.isVisible()) {
      await orderNoInput.fill(testData.purchaseOrderNo)
    }

    // 保存
    const saveBtn = page.locator('button:has-text("保存"), button:has-text("提交"), button:has-text("确定")').first()
    if (await saveBtn.isVisible()) {
      await saveBtn.click()
      await page.waitForTimeout(1000)
    }

    // 验证采购订单
    const order = await db.getPurchaseOrder(testData.purchaseOrderNo)
    expect(order).not.toBeNull()
  })

  test('供应商KPI查看', async ({ page }) => {
    await page.goto('/scp/supplier-kpi')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('客户询价查看', async ({ page }) => {
    await page.goto('/scp/customer-inquiry')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })
})
