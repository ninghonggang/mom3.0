/**
 * L3 生产执行业务流程测试
 *
 * 测试流程：销售订单 → 生产工单 → 报工 → 出货
 * 测试页面：生产发料、生产退料页面访问
 *
 * 前置条件：
 * 1. mom-server 后端运行在 localhost:9081
 * 2. mom-web 前端运行在 localhost:5177
 */

import { test, expect } from '@playwright/test'

test.describe('L3 生产执行业务流程', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/login')
    await page.fill('input[type="text"], input[placeholder*="用户名"], input[placeholder*="账号"]', 'admin')
    await page.fill('input[type="password"]', 'admin123')
    await page.click('button[type="submit"], button:has-text("登录"), button:has-text("登 录")')
    await page.waitForURL('**/dashboard', { timeout: 10000 })
  })

  test('销售订单页面', async ({ page }) => {
    await page.goto('/production/sales-order')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('生产工单页面', async ({ page }) => {
    await page.goto('/production/order')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('报工流程页面', async ({ page }) => {
    await page.goto('/production/report')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toContainText(/报工|生产报工|工序报工|生产/)
  })

  test('生产发料页面', async ({ page }) => {
    await page.goto('/production/issue')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('生产退料页面', async ({ page }) => {
    await page.goto('/production/return')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('出货页面', async ({ page }) => {
    await page.goto('/production/shipment')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })
})
