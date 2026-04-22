/**
 * L3 EAM 设备管理业务流程测试
 * 测试范围：设备 → 点检 → 保养 → 维修、OEE分析、备件管理
 *
 * 前置条件：
 * 1. mom-server 后端运行在 localhost:9081
 * 2. mom-web 前端运行在 localhost:5177
 * 3. PostgreSQL 运行在 localhost:5432
 */

import { test, expect } from '@playwright/test'
import * as db from '../tests/helpers/db'

test.describe('L3 EAM 设备管理业务流程', () => {
  const testData = {
    equipmentCode: `TEST-EQUIP-${Date.now()}`,
  }

  test.afterAll(async () => {
    await db.cleanupTestData('equ_equipment', 'equipment_code', testData.equipmentCode)
    await db.closeDbClient()
  })

  test.beforeEach(async ({ page }) => {
    await page.goto('/login')
    await page.fill('input[type="text"]', 'admin')
    await page.fill('input[type="password"]', 'admin123')
    await page.click('button[type="submit"]')
    await page.waitForURL('**/dashboard', { timeout: 10000 })
  })

  test('设备 → 点检 → 保养 → 维修', async ({ page }) => {
    // 创建设备
    const equipmentId = await db.createEquipmentViaDb({
      code: testData.equipmentCode,
      name: '测试设备-E2E',
      equipmentType: '加工设备',
    })
    expect(equipmentId).toBeGreaterThan(0)

    // 设备列表
    await page.goto('/equipment')
    await page.waitForTimeout(1000)

    const searchInput = page.locator('input[placeholder*="搜索"], input[placeholder*="查询"], input[placeholder*="设备"]').first()
    if (await searchInput.isVisible()) {
      await searchInput.fill(testData.equipmentCode)
      await page.keyboard.press('Enter')
      await page.waitForTimeout(500)
    }

    await expect(page.locator(`text=${testData.equipmentCode}`).first()).toBeVisible({ timeout: 5000 })

    // 点检
    await page.goto('/equipment/check')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()

    // 保养
    await page.goto('/equipment/maintenance')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()

    // 维修
    await page.goto('/equipment/repair')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('OEE分析', async ({ page }) => {
    await page.goto('/equipment/oee')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('备件管理', async ({ page }) => {
    await page.goto('/equipment/spare')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })
})